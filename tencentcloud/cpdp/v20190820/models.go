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

package v20190820

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Acct struct {

	// STRING(50)，见证子账户的账号（可重复）
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubAcctNo *string `json:"SubAcctNo,omitempty" name:"SubAcctNo"`

	// STRING(10)，见证子账户的属性（可重复。1: 普通会员子账号; 2: 挂账子账号; 3: 手续费子账号; 4: 利息子账号; 5: 平台担保子账号）
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubAcctProperty *string `json:"SubAcctProperty,omitempty" name:"SubAcctProperty"`

	// STRING(32)，交易网会员代码（可重复）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranNetMemberCode *string `json:"TranNetMemberCode,omitempty" name:"TranNetMemberCode"`

	// STRING(150)，见证子账户的名称（可重复）
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubAcctName *string `json:"SubAcctName,omitempty" name:"SubAcctName"`

	// STRING(20)，见证子账户可用余额（可重复）
	// 注意：此字段可能返回 null，表示取不到有效值。
	AcctAvailBal *string `json:"AcctAvailBal,omitempty" name:"AcctAvailBal"`

	// STRING(20)，见证子账户可提现金额（可重复。开户日期或修改日期）
	// 注意：此字段可能返回 null，表示取不到有效值。
	CashAmt *string `json:"CashAmt,omitempty" name:"CashAmt"`

	// STRING(8)，维护日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaintenanceDate *string `json:"MaintenanceDate,omitempty" name:"MaintenanceDate"`
}

type AgencyClientInfo struct {

	// 经办人姓名，存在经办人必输
	AgencyClientName *string `json:"AgencyClientName,omitempty" name:"AgencyClientName"`

	// 经办人证件类型，存在经办人必输
	AgencyClientGlobalType *string `json:"AgencyClientGlobalType,omitempty" name:"AgencyClientGlobalType"`

	// 经办人证件号，存在经办人必输
	AgencyClientGlobalId *string `json:"AgencyClientGlobalId,omitempty" name:"AgencyClientGlobalId"`

	// 经办人手机号，存在经办人必输
	AgencyClientMobile *string `json:"AgencyClientMobile,omitempty" name:"AgencyClientMobile"`
}

type AgentTaxPayment struct {

	// 主播银行账号
	AnchorId *string `json:"AnchorId,omitempty" name:"AnchorId"`

	// 主播姓名
	AnchorName *string `json:"AnchorName,omitempty" name:"AnchorName"`

	// 主播身份证
	AnchorIDCard *string `json:"AnchorIDCard,omitempty" name:"AnchorIDCard"`

	// 纳税的开始时间，格式yyyy-MM-dd
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 纳税的结束时间，格式yyyy-MM-dd
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 流水金额。以“分”为单位
	Amount *int64 `json:"Amount,omitempty" name:"Amount"`

	// 应缴税款。以“分”为单位
	Tax *int64 `json:"Tax,omitempty" name:"Tax"`
}

type AgentTaxPaymentBatch struct {

	// 状态消息
	StatusMsg *string `json:"StatusMsg,omitempty" name:"StatusMsg"`

	// 批次号
	BatchNum *int64 `json:"BatchNum,omitempty" name:"BatchNum"`

	// 录入记录的条数
	InfoNum *int64 `json:"InfoNum,omitempty" name:"InfoNum"`

	// 源电子凭证下载地址
	RawElectronicCertUrl *string `json:"RawElectronicCertUrl,omitempty" name:"RawElectronicCertUrl"`

	// 代理商账号
	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`

	// 文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 状态码。0表示下载成功
	StatusCode *int64 `json:"StatusCode,omitempty" name:"StatusCode"`

	// 渠道号
	Channel *int64 `json:"Channel,omitempty" name:"Channel"`

	// 0-视同，1-个体工商户
	Type *int64 `json:"Type,omitempty" name:"Type"`
}

type AnchorContractInfo struct {

	// 主播ID
	AnchorId *string `json:"AnchorId,omitempty" name:"AnchorId"`

	// 主播名称
	AnchorName *string `json:"AnchorName,omitempty" name:"AnchorName"`

	// 代理商ID
	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`

	// 代理商名称
	AgentName *string `json:"AgentName,omitempty" name:"AgentName"`

	// 主播身份证号
	IdNo *string `json:"IdNo,omitempty" name:"IdNo"`
}

type ApplyApplicationMaterialRequest struct {
	*tchttp.BaseRequest

	// 对接方汇出指令编号
	TransactionId *string `json:"TransactionId,omitempty" name:"TransactionId"`

	// 申报流水号
	DeclareId *string `json:"DeclareId,omitempty" name:"DeclareId"`

	// 付款人ID
	PayerId *string `json:"PayerId,omitempty" name:"PayerId"`

	// 源币种
	SourceCurrency *string `json:"SourceCurrency,omitempty" name:"SourceCurrency"`

	// 目的币种
	TargetCurrency *string `json:"TargetCurrency,omitempty" name:"TargetCurrency"`

	// 贸易编码
	TradeCode *string `json:"TradeCode,omitempty" name:"TradeCode"`

	// 原申报流水号
	OriginalDeclareId *string `json:"OriginalDeclareId,omitempty" name:"OriginalDeclareId"`

	// 源金额
	SourceAmount *int64 `json:"SourceAmount,omitempty" name:"SourceAmount"`

	// 目的金额
	TargetAmount *int64 `json:"TargetAmount,omitempty" name:"TargetAmount"`

	// 接入环境。沙箱环境填sandbox
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *ApplyApplicationMaterialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyApplicationMaterialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TransactionId")
	delete(f, "DeclareId")
	delete(f, "PayerId")
	delete(f, "SourceCurrency")
	delete(f, "TargetCurrency")
	delete(f, "TradeCode")
	delete(f, "OriginalDeclareId")
	delete(f, "SourceAmount")
	delete(f, "TargetAmount")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyApplicationMaterialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ApplyApplicationMaterialResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 提交申报材料结果
		Result *ApplyDeclareResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApplyApplicationMaterialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyApplicationMaterialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplyDeclareData struct {

	// 商户号
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 第三方指令编号
	TransactionId *string `json:"TransactionId,omitempty" name:"TransactionId"`

	// 受理状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 申报流水号
	DeclareId *string `json:"DeclareId,omitempty" name:"DeclareId"`

	// 原申报流水号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginalDeclareId *string `json:"OriginalDeclareId,omitempty" name:"OriginalDeclareId"`

	// 付款人ID
	PayerId *string `json:"PayerId,omitempty" name:"PayerId"`
}

type ApplyDeclareResult struct {

	// 错误码
	Code *string `json:"Code,omitempty" name:"Code"`

	// 提交申报材料数据
	Data *ApplyDeclareData `json:"Data,omitempty" name:"Data"`
}

type ApplyOutwardOrderData struct {

	// 商户号
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 对接方汇出指令编号
	TransactionId *string `json:"TransactionId,omitempty" name:"TransactionId"`

	// 受理状态
	Status *string `json:"Status,omitempty" name:"Status"`
}

type ApplyOutwardOrderRequest struct {
	*tchttp.BaseRequest

	// 对接方汇出指令编号
	TransactionId *string `json:"TransactionId,omitempty" name:"TransactionId"`

	// 定价币种
	PricingCurrency *string `json:"PricingCurrency,omitempty" name:"PricingCurrency"`

	// 源币种
	SourceCurrency *string `json:"SourceCurrency,omitempty" name:"SourceCurrency"`

	// 目的币种
	TargetCurrency *string `json:"TargetCurrency,omitempty" name:"TargetCurrency"`

	// 收款人类型（银行卡填"BANK_ACCOUNT"）
	PayeeType *string `json:"PayeeType,omitempty" name:"PayeeType"`

	// 收款人账号
	PayeeAccount *string `json:"PayeeAccount,omitempty" name:"PayeeAccount"`

	// 源币种金额
	SourceAmount *float64 `json:"SourceAmount,omitempty" name:"SourceAmount"`

	// 目的金额
	TargetAmount *float64 `json:"TargetAmount,omitempty" name:"TargetAmount"`

	// 收款人姓名（PayeeType为"BANK_COUNT"时必填）
	PayeeName *string `json:"PayeeName,omitempty" name:"PayeeName"`

	// 收款人地址（PayeeType为"BANK_COUNT"时必填）
	PayeeAddress *string `json:"PayeeAddress,omitempty" name:"PayeeAddress"`

	// 收款人银行账号类型（PayeeType为"BANK_COUNT"时必填）
	// 个人填"INDIVIDUAL"
	// 企业填"CORPORATE"
	PayeeBankAccountType *string `json:"PayeeBankAccountType,omitempty" name:"PayeeBankAccountType"`

	// 收款人国家或地区编码（PayeeType为"BANK_COUNT"时必填）
	PayeeCountryCode *string `json:"PayeeCountryCode,omitempty" name:"PayeeCountryCode"`

	// 收款人开户银行名称（PayeeType为"BANK_COUNT"时必填）
	PayeeBankName *string `json:"PayeeBankName,omitempty" name:"PayeeBankName"`

	// 收款人开户银行地址（PayeeType为"BANK_COUNT"时必填）
	PayeeBankAddress *string `json:"PayeeBankAddress,omitempty" name:"PayeeBankAddress"`

	// 收款人开户银行所在国家或地区编码（PayeeType为"BANK_COUNT"时必填）
	PayeeBankDistrict *string `json:"PayeeBankDistrict,omitempty" name:"PayeeBankDistrict"`

	// 收款银行SwiftCode（PayeeType为"BANK_COUNT"时必填）
	PayeeBankSwiftCode *string `json:"PayeeBankSwiftCode,omitempty" name:"PayeeBankSwiftCode"`

	// 收款银行国际编码类型
	PayeeBankType *string `json:"PayeeBankType,omitempty" name:"PayeeBankType"`

	// 收款银行国际编码
	PayeeBankCode *string `json:"PayeeBankCode,omitempty" name:"PayeeBankCode"`

	// 收款人附言
	ReferenceForBeneficiary *string `json:"ReferenceForBeneficiary,omitempty" name:"ReferenceForBeneficiary"`

	// 接入环境。沙箱环境填sandbox
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *ApplyOutwardOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyOutwardOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TransactionId")
	delete(f, "PricingCurrency")
	delete(f, "SourceCurrency")
	delete(f, "TargetCurrency")
	delete(f, "PayeeType")
	delete(f, "PayeeAccount")
	delete(f, "SourceAmount")
	delete(f, "TargetAmount")
	delete(f, "PayeeName")
	delete(f, "PayeeAddress")
	delete(f, "PayeeBankAccountType")
	delete(f, "PayeeCountryCode")
	delete(f, "PayeeBankName")
	delete(f, "PayeeBankAddress")
	delete(f, "PayeeBankDistrict")
	delete(f, "PayeeBankSwiftCode")
	delete(f, "PayeeBankType")
	delete(f, "PayeeBankCode")
	delete(f, "ReferenceForBeneficiary")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyOutwardOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ApplyOutwardOrderResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 汇出指令申请
		Result *ApplyOutwardOrderResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApplyOutwardOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyOutwardOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplyOutwardOrderResult struct {

	// 汇出指令申请数据
	Data *ApplyOutwardOrderData `json:"Data,omitempty" name:"Data"`

	// 错误码
	Code *string `json:"Code,omitempty" name:"Code"`
}

type ApplyPayerInfoRequest struct {
	*tchttp.BaseRequest

	// 付款人ID
	PayerId *string `json:"PayerId,omitempty" name:"PayerId"`

	// 付款人类型 (个人: INDIVIDUAL, 企业: CORPORATE)
	PayerType *string `json:"PayerType,omitempty" name:"PayerType"`

	// 付款人姓名
	PayerName *string `json:"PayerName,omitempty" name:"PayerName"`

	// 付款人证件类型 (身份证: ID_CARD, 统一社会信用代码: UNIFIED_CREDIT_CODE)
	PayerIdType *string `json:"PayerIdType,omitempty" name:"PayerIdType"`

	// 付款人证件号
	PayerIdNo *string `json:"PayerIdNo,omitempty" name:"PayerIdNo"`

	// 付款人常驻国家或地区编码 (见常见问题-国家/地区编码)
	PayerCountryCode *string `json:"PayerCountryCode,omitempty" name:"PayerCountryCode"`

	// 付款人联系人名称
	PayerContactName *string `json:"PayerContactName,omitempty" name:"PayerContactName"`

	// 付款人联系电话
	PayerContactNumber *string `json:"PayerContactNumber,omitempty" name:"PayerContactNumber"`

	// 付款人联系邮箱
	PayerEmailAddress *string `json:"PayerEmailAddress,omitempty" name:"PayerEmailAddress"`

	// 接入环境。沙箱环境填sandbox
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *ApplyPayerInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyPayerInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PayerId")
	delete(f, "PayerType")
	delete(f, "PayerName")
	delete(f, "PayerIdType")
	delete(f, "PayerIdNo")
	delete(f, "PayerCountryCode")
	delete(f, "PayerContactName")
	delete(f, "PayerContactNumber")
	delete(f, "PayerEmailAddress")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyPayerInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ApplyPayerInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 付款人申请结果
		Result *ApplyPayerinfoResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApplyPayerInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyPayerInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplyPayerinfoData struct {

	// 商户号
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 付款人ID
	PayerId *string `json:"PayerId,omitempty" name:"PayerId"`

	// 状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailReason *string `json:"FailReason,omitempty" name:"FailReason"`
}

type ApplyPayerinfoResult struct {

	// 错误码
	Code *string `json:"Code,omitempty" name:"Code"`

	// 数据
	Data *ApplyPayerinfoData `json:"Data,omitempty" name:"Data"`
}

type ApplyReWithdrawalRequest struct {
	*tchttp.BaseRequest

	// 聚鑫业务类型
	BusinessType *uint64 `json:"BusinessType,omitempty" name:"BusinessType"`

	// 由平台客服提供的计费密钥Id
	MidasSecretId *string `json:"MidasSecretId,omitempty" name:"MidasSecretId"`

	// 计费签名
	MidasSignature *string `json:"MidasSignature,omitempty" name:"MidasSignature"`

	// 提现信息
	Body *WithdrawBill `json:"Body,omitempty" name:"Body"`

	// 聚鑫业务ID
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 环境名:
	// release: 现网环境
	// sandbox: 沙箱环境
	// development: 开发环境
	// 缺省: release
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`
}

func (r *ApplyReWithdrawalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyReWithdrawalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BusinessType")
	delete(f, "MidasSecretId")
	delete(f, "MidasSignature")
	delete(f, "Body")
	delete(f, "MidasAppId")
	delete(f, "MidasEnvironment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyReWithdrawalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ApplyReWithdrawalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 重新提现业务订单号
		WithdrawOrderId *string `json:"WithdrawOrderId,omitempty" name:"WithdrawOrderId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApplyReWithdrawalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyReWithdrawalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplyTradeData struct {

	// 商户号
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 贸易材料流水号
	TradeFileId *string `json:"TradeFileId,omitempty" name:"TradeFileId"`

	// 交易币种
	TradeCurrency *string `json:"TradeCurrency,omitempty" name:"TradeCurrency"`

	// 交易金额
	TradeAmount *string `json:"TradeAmount,omitempty" name:"TradeAmount"`

	// 付款人ID
	PayerId *string `json:"PayerId,omitempty" name:"PayerId"`

	// 状态
	Status *string `json:"Status,omitempty" name:"Status"`
}

type ApplyTradeRequest struct {
	*tchttp.BaseRequest

	// 贸易材料流水号
	TradeFileId *string `json:"TradeFileId,omitempty" name:"TradeFileId"`

	// 贸易材料订单号
	TradeOrderId *string `json:"TradeOrderId,omitempty" name:"TradeOrderId"`

	// 付款人ID
	PayerId *string `json:"PayerId,omitempty" name:"PayerId"`

	// 收款人姓名
	PayeeName *string `json:"PayeeName,omitempty" name:"PayeeName"`

	// 收款人常驻国家或地区编码 (见常见问题)
	PayeeCountryCode *string `json:"PayeeCountryCode,omitempty" name:"PayeeCountryCode"`

	// 贸易类型 (GOODS: 商品, SERVICE: 服务)
	TradeType *string `json:"TradeType,omitempty" name:"TradeType"`

	// 交易时间 (格式: yyyyMMdd)
	TradeTime *string `json:"TradeTime,omitempty" name:"TradeTime"`

	// 交易币种
	TradeCurrency *string `json:"TradeCurrency,omitempty" name:"TradeCurrency"`

	// 交易金额
	TradeAmount *float64 `json:"TradeAmount,omitempty" name:"TradeAmount"`

	// 交易名称 
	// (TradeType=GOODS时填写物品名称，可填写多个，格式无要求；
	// TradeType=SERVICE时填写贸易类别，见常见问题-贸易类别)
	TradeName *string `json:"TradeName,omitempty" name:"TradeName"`

	// 交易数量 (TradeType=GOODS 填写物品数量, TradeType=SERVICE填写服务次数)
	TradeCount *int64 `json:"TradeCount,omitempty" name:"TradeCount"`

	// 货贸承运人 (TradeType=GOODS 必填)
	GoodsCarrier *string `json:"GoodsCarrier,omitempty" name:"GoodsCarrier"`

	// 服贸交易细节 (TradeType=GOODS 必填, 见常见问题-交易细节)
	ServiceDetail *string `json:"ServiceDetail,omitempty" name:"ServiceDetail"`

	// 服贸服务时间 (TradeType=GOODS 必填, 见常见问题-服务时间)
	ServiceTime *string `json:"ServiceTime,omitempty" name:"ServiceTime"`

	// 接入环境。沙箱环境填sandbox
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *ApplyTradeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyTradeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TradeFileId")
	delete(f, "TradeOrderId")
	delete(f, "PayerId")
	delete(f, "PayeeName")
	delete(f, "PayeeCountryCode")
	delete(f, "TradeType")
	delete(f, "TradeTime")
	delete(f, "TradeCurrency")
	delete(f, "TradeAmount")
	delete(f, "TradeName")
	delete(f, "TradeCount")
	delete(f, "GoodsCarrier")
	delete(f, "ServiceDetail")
	delete(f, "ServiceTime")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyTradeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ApplyTradeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 提交贸易材料结果
		Result *ApplyTradeResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApplyTradeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyTradeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplyTradeResult struct {

	// 错误码
	Code *string `json:"Code,omitempty" name:"Code"`

	// 提交贸易材料数据
	Data *ApplyTradeData `json:"Data,omitempty" name:"Data"`
}

type ApplyWithdrawalRequest struct {
	*tchttp.BaseRequest

	// 聚鑫分配的支付主MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 聚鑫计费SubAppId，代表子商户
	SubAppId *string `json:"SubAppId,omitempty" name:"SubAppId"`

	// 用于提现
	// <敏感信息>加密详见<a href="https://cloud.tencent.com/document/product/1122/48979" target="_blank">《商户端接口敏感信息加密说明》</a>
	SettleAcctNo *string `json:"SettleAcctNo,omitempty" name:"SettleAcctNo"`

	// 结算账户户名
	// <敏感信息>加密详见<a href="https://cloud.tencent.com/document/product/1122/48979" target="_blank">《商户端接口敏感信息加密说明》</a>
	SettleAcctName *string `json:"SettleAcctName,omitempty" name:"SettleAcctName"`

	// 币种 RMB
	CurrencyType *string `json:"CurrencyType,omitempty" name:"CurrencyType"`

	// 单位，1：元，2：角，3：分
	CurrencyUnit *int64 `json:"CurrencyUnit,omitempty" name:"CurrencyUnit"`

	// 金额
	CurrencyAmt *string `json:"CurrencyAmt,omitempty" name:"CurrencyAmt"`

	// 交易网名称
	TranWebName *string `json:"TranWebName,omitempty" name:"TranWebName"`

	// 会员证件类型
	IdType *string `json:"IdType,omitempty" name:"IdType"`

	// 会员证件号码
	// <敏感信息>加密详见<a href="https://cloud.tencent.com/document/product/1122/48979" target="_blank">《商户端接口敏感信息加密说明》</a>
	IdCode *string `json:"IdCode,omitempty" name:"IdCode"`

	// 聚鑫分配的安全ID
	MidasSecretId *string `json:"MidasSecretId,omitempty" name:"MidasSecretId"`

	// 按照聚鑫安全密钥计算的签名
	MidasSignature *string `json:"MidasSignature,omitempty" name:"MidasSignature"`

	// 敏感信息加密类型:
	// RSA: rsa非对称加密，使用RSA-PKCS1-v1_5
	// AES: aes对称加密，使用AES256-CBC-PCKS7padding
	// 缺省: RSA
	EncryptType *string `json:"EncryptType,omitempty" name:"EncryptType"`

	// 环境名:
	// release: 现网环境
	// sandbox: 沙箱环境
	// development: 开发环境
	// 缺省: release
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`

	// 手续费金额
	CommissionAmount *string `json:"CommissionAmount,omitempty" name:"CommissionAmount"`

	// 提现单号，长度32字节
	WithdrawOrderId *string `json:"WithdrawOrderId,omitempty" name:"WithdrawOrderId"`
}

func (r *ApplyWithdrawalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyWithdrawalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MidasAppId")
	delete(f, "SubAppId")
	delete(f, "SettleAcctNo")
	delete(f, "SettleAcctName")
	delete(f, "CurrencyType")
	delete(f, "CurrencyUnit")
	delete(f, "CurrencyAmt")
	delete(f, "TranWebName")
	delete(f, "IdType")
	delete(f, "IdCode")
	delete(f, "MidasSecretId")
	delete(f, "MidasSignature")
	delete(f, "EncryptType")
	delete(f, "MidasEnvironment")
	delete(f, "CommissionAmount")
	delete(f, "WithdrawOrderId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyWithdrawalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ApplyWithdrawalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApplyWithdrawalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyWithdrawalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BankCardItem struct {

	// 超级网银行号
	EiconBankBranchId *string `json:"EiconBankBranchId,omitempty" name:"EiconBankBranchId"`

	// 大小额行号
	CnapsBranchId *string `json:"CnapsBranchId,omitempty" name:"CnapsBranchId"`

	// 结算账户类型
	// 1 – 本行账户
	// 2 – 他行账户
	SettleAcctType *int64 `json:"SettleAcctType,omitempty" name:"SettleAcctType"`

	// 结算账户户名
	// <敏感信息>
	SettleAcctName *string `json:"SettleAcctName,omitempty" name:"SettleAcctName"`

	// 开户行名称
	AcctBranchName *string `json:"AcctBranchName,omitempty" name:"AcctBranchName"`

	// 用于提现
	// <敏感信息>
	SettleAcctNo *string `json:"SettleAcctNo,omitempty" name:"SettleAcctNo"`

	// 聚鑫计费SubAppId，代表子商户
	SubAppId *string `json:"SubAppId,omitempty" name:"SubAppId"`

	// 验证类型
	// 1 – 小额转账验证
	// 2 – 短信验证
	BindType *int64 `json:"BindType,omitempty" name:"BindType"`

	// 用于短信验证
	// BindType==2时必填
	// <敏感信息>
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 证件类型
	IdType *string `json:"IdType,omitempty" name:"IdType"`

	// 证件号码
	// <敏感信息>
	IdCode *string `json:"IdCode,omitempty" name:"IdCode"`
}

type BindAccountRequest struct {
	*tchttp.BaseRequest

	// 主播Id
	AnchorId *string `json:"AnchorId,omitempty" name:"AnchorId"`

	// 1 微信企业付款 
	// 2 支付宝转账 
	// 3 平安银企直连代发转账
	TransferType *int64 `json:"TransferType,omitempty" name:"TransferType"`

	// 收款方标识。
	// 微信为open_id；
	// 支付宝为会员alipay_user_id;
	// 平安为收款方银行账号;
	AccountNo *string `json:"AccountNo,omitempty" name:"AccountNo"`

	// 手机号
	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`
}

func (r *BindAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AnchorId")
	delete(f, "TransferType")
	delete(f, "AccountNo")
	delete(f, "PhoneNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type BindAccountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 错误码。响应成功："SUCCESS"，其他为不成功。
		ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

		// 响应消息。
		ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

		// 该字段为null。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindAcctRequest struct {
	*tchttp.BaseRequest

	// 聚鑫分配的支付主MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 聚鑫计费SubAppId，代表子商户
	SubAppId *string `json:"SubAppId,omitempty" name:"SubAppId"`

	// 1 – 小额转账验证
	// 2 – 短信验证
	// 3 - 一分钱转账验证，无需再调CheckAcct验证绑卡
	// 4 - 银行四要素验证，无需再调CheckAcct验证绑卡
	// 每个结算账户每天只能使用一次小额转账验证
	BindType *int64 `json:"BindType,omitempty" name:"BindType"`

	// 用于提现
	// <敏感信息>加密详见<a href="https://cloud.tencent.com/document/product/1122/48979" target="_blank">《商户端接口敏感信息加密说明》</a>
	SettleAcctNo *string `json:"SettleAcctNo,omitempty" name:"SettleAcctNo"`

	// 结算账户户名
	// <敏感信息>加密详见<a href="https://cloud.tencent.com/document/product/1122/48979" target="_blank">《商户端接口敏感信息加密说明》</a>
	SettleAcctName *string `json:"SettleAcctName,omitempty" name:"SettleAcctName"`

	// 1 – 本行账户
	// 2 – 他行账户
	SettleAcctType *int64 `json:"SettleAcctType,omitempty" name:"SettleAcctType"`

	// 证件类型，见《证件类型》表
	IdType *string `json:"IdType,omitempty" name:"IdType"`

	// 证件号码
	// <敏感信息>加密详见<a href="https://cloud.tencent.com/document/product/1122/48979" target="_blank">《商户端接口敏感信息加密说明》</a>
	IdCode *string `json:"IdCode,omitempty" name:"IdCode"`

	// 开户行名称
	AcctBranchName *string `json:"AcctBranchName,omitempty" name:"AcctBranchName"`

	// 聚鑫分配的安全ID
	MidasSecretId *string `json:"MidasSecretId,omitempty" name:"MidasSecretId"`

	// 按照聚鑫安全密钥计算的签名
	MidasSignature *string `json:"MidasSignature,omitempty" name:"MidasSignature"`

	// 用于短信验证
	// BindType==2时必填
	// <敏感信息>加密详见<a href="https://cloud.tencent.com/document/product/1122/48979" target="_blank">《商户端接口敏感信息加密说明》</a>
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 大小额行号，超级网银行号和大小额行号
	// 二选一
	CnapsBranchId *string `json:"CnapsBranchId,omitempty" name:"CnapsBranchId"`

	// 超级网银行号，超级网银行号和大小额行号
	// 二选一
	EiconBankBranchId *string `json:"EiconBankBranchId,omitempty" name:"EiconBankBranchId"`

	// 敏感信息加密类型:
	// RSA: rsa非对称加密，使用RSA-PKCS1-v1_5
	// AES: aes对称加密，使用AES256-CBC-PCKS7padding
	// 缺省: RSA
	EncryptType *string `json:"EncryptType,omitempty" name:"EncryptType"`

	// 环境名:
	// release: 现网环境
	// sandbox: 沙箱环境
	// development: 开发环境
	// 缺省: release
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`

	// 经办人信息
	AgencyClientInfo *AgencyClientInfo `json:"AgencyClientInfo,omitempty" name:"AgencyClientInfo"`
}

func (r *BindAcctRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindAcctRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MidasAppId")
	delete(f, "SubAppId")
	delete(f, "BindType")
	delete(f, "SettleAcctNo")
	delete(f, "SettleAcctName")
	delete(f, "SettleAcctType")
	delete(f, "IdType")
	delete(f, "IdCode")
	delete(f, "AcctBranchName")
	delete(f, "MidasSecretId")
	delete(f, "MidasSignature")
	delete(f, "Mobile")
	delete(f, "CnapsBranchId")
	delete(f, "EiconBankBranchId")
	delete(f, "EncryptType")
	delete(f, "MidasEnvironment")
	delete(f, "AgencyClientInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindAcctRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type BindAcctResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindAcctResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindAcctResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindRelateAccReUnionPayRequest struct {
	*tchttp.BaseRequest

	// String(22)，商户号（签约客户号）
	MrchCode *string `json:"MrchCode,omitempty" name:"MrchCode"`

	// STRING(32)，交易网会员代码（若需要把一个待绑定账户关联到两个会员名下，此字段可上送两个会员的交易网代码，并且须用“|::|”（右侧）进行分隔）
	TranNetMemberCode *string `json:"TranNetMemberCode,omitempty" name:"TranNetMemberCode"`

	// STRING(50)，会员的待绑定账户的账号（即 BindRelateAcctUnionPay接口中的“会员的待绑定账户的账号”）
	MemberAcctNo *string `json:"MemberAcctNo,omitempty" name:"MemberAcctNo"`

	// STRING(20)，短信验证码（即 BindRelateAcctUnionPay接口中的手机所接收到的短信验证码）
	MessageCheckCode *string `json:"MessageCheckCode,omitempty" name:"MessageCheckCode"`

	// STRING(1027)，保留域
	ReservedMsg *string `json:"ReservedMsg,omitempty" name:"ReservedMsg"`

	// STRING(12)，接入环境，默认接入沙箱环境。接入正式环境填"prod"
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *BindRelateAccReUnionPayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindRelateAccReUnionPayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MrchCode")
	delete(f, "TranNetMemberCode")
	delete(f, "MemberAcctNo")
	delete(f, "MessageCheckCode")
	delete(f, "ReservedMsg")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindRelateAccReUnionPayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type BindRelateAccReUnionPayResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// STRING(52)，见证系统流水号
	// 注意：此字段可能返回 null，表示取不到有效值。
		FrontSeqNo *string `json:"FrontSeqNo,omitempty" name:"FrontSeqNo"`

		// STRING(1027)，保留域
	// 注意：此字段可能返回 null，表示取不到有效值。
		ReservedMsg *string `json:"ReservedMsg,omitempty" name:"ReservedMsg"`

		// String(20)，返回码
		TxnReturnCode *string `json:"TxnReturnCode,omitempty" name:"TxnReturnCode"`

		// String(100)，返回信息
		TxnReturnMsg *string `json:"TxnReturnMsg,omitempty" name:"TxnReturnMsg"`

		// String(22)，交易流水号
		CnsmrSeqNo *string `json:"CnsmrSeqNo,omitempty" name:"CnsmrSeqNo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindRelateAccReUnionPayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindRelateAccReUnionPayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindRelateAcctSmallAmountRequest struct {
	*tchttp.BaseRequest

	// String(22)，商户号（签约客户号）
	MrchCode *string `json:"MrchCode,omitempty" name:"MrchCode"`

	// STRING(32)，交易网会员代码（若需要把一个待绑定账户关联到两个会员名下，此字段可上送两个会员的交易网代码，并且须用“|::|”(右侧)进行分隔）
	TranNetMemberCode *string `json:"TranNetMemberCode,omitempty" name:"TranNetMemberCode"`

	// STRING(150)，见证子账户的户名（首次绑定的情况下，此字段即为待绑定的提现账户的户名。非首次绑定的情况下，须注意带绑定的提现账户的户名须与留存在后台系统的会员户名一致）
	MemberName *string `json:"MemberName,omitempty" name:"MemberName"`

	// STRING(5)，会员证件类型（详情见“常见问题”）
	MemberGlobalType *string `json:"MemberGlobalType,omitempty" name:"MemberGlobalType"`

	// STRING(32)，会员证件号码
	MemberGlobalId *string `json:"MemberGlobalId,omitempty" name:"MemberGlobalId"`

	// STRING(50)，会员的待绑定账户的账号（提现的银行卡）
	MemberAcctNo *string `json:"MemberAcctNo,omitempty" name:"MemberAcctNo"`

	// STRING(10)，会员的待绑定账户的本他行类型（1: 本行; 2: 他行）
	BankType *string `json:"BankType,omitempty" name:"BankType"`

	// STRING(150)，会员的待绑定账户的开户行名称
	AcctOpenBranchName *string `json:"AcctOpenBranchName,omitempty" name:"AcctOpenBranchName"`

	// STRING(30)，会员的手机号（手机号须由长度为11位的数字构成）
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// STRING(20)，会员的待绑定账户的开户行的联行号（本他行类型为他行的情况下，此字段和下一个字段至少一个不为空）
	CnapsBranchId *string `json:"CnapsBranchId,omitempty" name:"CnapsBranchId"`

	// STRING(20)，会员的待绑定账户的开户行的超级网银行号（本他行类型为他行的情况下，此字段和上一个字段至少一个不为空）
	EiconBankBranchId *string `json:"EiconBankBranchId,omitempty" name:"EiconBankBranchId"`

	// STRING(1027)，转账方式（1: 往账鉴权(默认值); 2: 来账鉴权）
	ReservedMsg *string `json:"ReservedMsg,omitempty" name:"ReservedMsg"`

	// STRING(12)，接入环境，默认接入沙箱环境。接入正式环境填"prod"
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *BindRelateAcctSmallAmountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindRelateAcctSmallAmountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MrchCode")
	delete(f, "TranNetMemberCode")
	delete(f, "MemberName")
	delete(f, "MemberGlobalType")
	delete(f, "MemberGlobalId")
	delete(f, "MemberAcctNo")
	delete(f, "BankType")
	delete(f, "AcctOpenBranchName")
	delete(f, "Mobile")
	delete(f, "CnapsBranchId")
	delete(f, "EiconBankBranchId")
	delete(f, "ReservedMsg")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindRelateAcctSmallAmountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type BindRelateAcctSmallAmountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// String(20)，返回码
		TxnReturnCode *string `json:"TxnReturnCode,omitempty" name:"TxnReturnCode"`

		// String(100)，返回信息
		TxnReturnMsg *string `json:"TxnReturnMsg,omitempty" name:"TxnReturnMsg"`

		// String(22)，交易流水号
		CnsmrSeqNo *string `json:"CnsmrSeqNo,omitempty" name:"CnsmrSeqNo"`

		// STRING(1027)，保留域（来账鉴权的方式下，此字段的值为客户需往监管账户转入的金额。在同名子账户绑定的场景下，若返回""VERIFIED""则说明无需验证直接绑定成功）
	// 注意：此字段可能返回 null，表示取不到有效值。
		ReservedMsg *string `json:"ReservedMsg,omitempty" name:"ReservedMsg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindRelateAcctSmallAmountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindRelateAcctSmallAmountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindRelateAcctUnionPayRequest struct {
	*tchttp.BaseRequest

	// STRING(32)，交易网会员代码（若需要把一个待绑定账户关联到两个会员名下，此字段可上送两个会员的交易网代码，并且须用“|::|”（右侧）进行分隔）
	TranNetMemberCode *string `json:"TranNetMemberCode,omitempty" name:"TranNetMemberCode"`

	// STRING(150)，见证子账户的户名（首次绑定的情况下，此字段即为待绑定的提现账户的户名。非首次绑定的情况下，须注意带绑定的提现账户的户名须与留存在后台系统的会员户名一致）
	MemberName *string `json:"MemberName,omitempty" name:"MemberName"`

	// STRING(5)，会员证件类型（详情见“常见问题”）
	MemberGlobalType *string `json:"MemberGlobalType,omitempty" name:"MemberGlobalType"`

	// STRING(32)，会员证件号码
	MemberGlobalId *string `json:"MemberGlobalId,omitempty" name:"MemberGlobalId"`

	// STRING(50)，会员的待绑定账户的账号（提现的银行卡）
	MemberAcctNo *string `json:"MemberAcctNo,omitempty" name:"MemberAcctNo"`

	// STRING(10)，会员的待绑定账户的本他行类型（1: 本行; 2: 他行）
	BankType *string `json:"BankType,omitempty" name:"BankType"`

	// STRING(150)，会员的待绑定账户的开户行名称（若大小额行号不填则送超级网银号对应的银行名称，若填大小额行号则送大小额行号对应的银行名称）
	AcctOpenBranchName *string `json:"AcctOpenBranchName,omitempty" name:"AcctOpenBranchName"`

	// STRING(30)，会员的手机号（手机号须由长度为11位的数字构成）
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// String(22)，商户号（签约客户号）
	MrchCode *string `json:"MrchCode,omitempty" name:"MrchCode"`

	// STRING(20)，会员的待绑定账户的开户行的联行号（本他行类型为他行的情况下，此字段和下一个字段至少一个不为空）
	CnapsBranchId *string `json:"CnapsBranchId,omitempty" name:"CnapsBranchId"`

	// STRING(20)，会员的待绑定账户的开户行的超级网银行号（本他行类型为他行的情况下，此字段和上一个字段至少一个不为空）
	EiconBankBranchId *string `json:"EiconBankBranchId,omitempty" name:"EiconBankBranchId"`

	// STRING(1027)，保留域
	ReservedMsg *string `json:"ReservedMsg,omitempty" name:"ReservedMsg"`

	// STRING(12)，接入环境，默认接入沙箱环境。接入正式环境填"prod"
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *BindRelateAcctUnionPayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindRelateAcctUnionPayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TranNetMemberCode")
	delete(f, "MemberName")
	delete(f, "MemberGlobalType")
	delete(f, "MemberGlobalId")
	delete(f, "MemberAcctNo")
	delete(f, "BankType")
	delete(f, "AcctOpenBranchName")
	delete(f, "Mobile")
	delete(f, "MrchCode")
	delete(f, "CnapsBranchId")
	delete(f, "EiconBankBranchId")
	delete(f, "ReservedMsg")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindRelateAcctUnionPayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type BindRelateAcctUnionPayResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// STRING(1027)，保留域（在同名子账户绑定的场景下，若返回"VERIFIED"则说明无需验证直接绑定成功）
	// 注意：此字段可能返回 null，表示取不到有效值。
		ReservedMsg *string `json:"ReservedMsg,omitempty" name:"ReservedMsg"`

		// String(20)，返回码
		TxnReturnCode *string `json:"TxnReturnCode,omitempty" name:"TxnReturnCode"`

		// String(100)，返回信息
		TxnReturnMsg *string `json:"TxnReturnMsg,omitempty" name:"TxnReturnMsg"`

		// String(22)，交易流水号
		CnsmrSeqNo *string `json:"CnsmrSeqNo,omitempty" name:"CnsmrSeqNo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindRelateAcctUnionPayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindRelateAcctUnionPayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChannelContractInfo struct {

	// 外部合约协议号
	OutContractCode *string `json:"OutContractCode,omitempty" name:"OutContractCode"`

	// 米大师内部生成的合约协议号
	ChannelContractCode *string `json:"ChannelContractCode,omitempty" name:"ChannelContractCode"`
}

type ChannelReturnContractInfo struct {

	// 平台合约状态
	// 协议状态，枚举值：
	// CONTRACT_STATUS_SIGNED：已签约
	// CONTRACT_STATUS_TERMINATED：未签约
	// CONTRACT_STATUS_PENDING：签约进行中
	ContractStatus *string `json:"ContractStatus,omitempty" name:"ContractStatus"`

	// 米大师内部存放的合约信息
	ChannelContractInfo *ChannelContractInfo `json:"ChannelContractInfo,omitempty" name:"ChannelContractInfo"`
}

type CheckAcctRequest struct {
	*tchttp.BaseRequest

	// 聚鑫分配的支付主MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 聚鑫计费SubAppId，代表子商户
	SubAppId *string `json:"SubAppId,omitempty" name:"SubAppId"`

	// 1 – 小额转账验证
	// 2 – 短信验证
	// 每个结算账户每天只能使用一次小额转账验证
	BindType *int64 `json:"BindType,omitempty" name:"BindType"`

	// 结算账户账号
	// <敏感信息>加密详见<a href="https://cloud.tencent.com/document/product/1122/48979" target="_blank">《商户端接口敏感信息加密说明》</a>
	SettleAcctNo *string `json:"SettleAcctNo,omitempty" name:"SettleAcctNo"`

	// 聚鑫分配的安全ID
	MidasSecretId *string `json:"MidasSecretId,omitempty" name:"MidasSecretId"`

	// 按照聚鑫安全密钥计算的签名
	MidasSignature *string `json:"MidasSignature,omitempty" name:"MidasSignature"`

	// 短信验证码或指令号
	// BindType==2必填，平安渠道必填
	CheckCode *string `json:"CheckCode,omitempty" name:"CheckCode"`

	// 币种 RMB
	// BindType==1必填
	CurrencyType *string `json:"CurrencyType,omitempty" name:"CurrencyType"`

	// 单位
	// 1：元，2：角，3：分
	// BindType==1必填
	CurrencyUnit *int64 `json:"CurrencyUnit,omitempty" name:"CurrencyUnit"`

	// 金额
	// BindType==1必填
	CurrencyAmt *string `json:"CurrencyAmt,omitempty" name:"CurrencyAmt"`

	// 敏感信息加密类型:
	// RSA: rsa非对称加密，使用RSA-PKCS1-v1_5
	// AES: aes对称加密，使用AES256-CBC-PCKS7padding
	// 缺省: RSA
	EncryptType *string `json:"EncryptType,omitempty" name:"EncryptType"`

	// 环境名:
	// release: 现网环境
	// sandbox: 沙箱环境
	// development: 开发环境
	// 缺省: release
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`
}

func (r *CheckAcctRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckAcctRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MidasAppId")
	delete(f, "SubAppId")
	delete(f, "BindType")
	delete(f, "SettleAcctNo")
	delete(f, "MidasSecretId")
	delete(f, "MidasSignature")
	delete(f, "CheckCode")
	delete(f, "CurrencyType")
	delete(f, "CurrencyUnit")
	delete(f, "CurrencyAmt")
	delete(f, "EncryptType")
	delete(f, "MidasEnvironment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckAcctRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CheckAcctResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 前置流水号，请保存
		FrontSeqNo *string `json:"FrontSeqNo,omitempty" name:"FrontSeqNo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckAcctResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckAcctResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckAmountRequest struct {
	*tchttp.BaseRequest

	// String(22)，商户号（签约客户号）
	MrchCode *string `json:"MrchCode,omitempty" name:"MrchCode"`

	// STRING(32)，交易网会员代码（若需要把一个待绑定账户关联到两个会员名下，此字段可上送两个会员的交易网代码，并且须用“|::|”(右侧)进行分隔）
	TranNetMemberCode *string `json:"TranNetMemberCode,omitempty" name:"TranNetMemberCode"`

	// STRING(50)，会员的待绑定账户的账号（即 BindRelateAcctSmallAmount接口中的“会员的待绑定账户的账号”）
	TakeCashAcctNo *string `json:"TakeCashAcctNo,omitempty" name:"TakeCashAcctNo"`

	// STRING(20)，鉴权验证金额（即 BindRelateAcctSmallAmount接口中的“会员的待绑定账户收到的验证金额。原小额转账鉴权方式为来账鉴权的情况下此字段须赋值为0.00）
	AuthAmt *string `json:"AuthAmt,omitempty" name:"AuthAmt"`

	// STRING(3)，币种（默认为RMB）
	Ccy *string `json:"Ccy,omitempty" name:"Ccy"`

	// STRING(1027)，原小额转账方式（1: 往账鉴权，此为默认值; 2: 来账鉴权）
	ReservedMsg *string `json:"ReservedMsg,omitempty" name:"ReservedMsg"`

	// STRING(12)，接入环境，默认接入沙箱环境。接入正式环境填"prod"
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *CheckAmountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckAmountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MrchCode")
	delete(f, "TranNetMemberCode")
	delete(f, "TakeCashAcctNo")
	delete(f, "AuthAmt")
	delete(f, "Ccy")
	delete(f, "ReservedMsg")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckAmountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CheckAmountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// String(20)，返回码
		TxnReturnCode *string `json:"TxnReturnCode,omitempty" name:"TxnReturnCode"`

		// String(100)，返回信息
		TxnReturnMsg *string `json:"TxnReturnMsg,omitempty" name:"TxnReturnMsg"`

		// String(22)，交易流水号
		CnsmrSeqNo *string `json:"CnsmrSeqNo,omitempty" name:"CnsmrSeqNo"`

		// STRING(52)，见证系统流水号（即电商见证宝系统生成的流水号，可关联具体一笔请求）
		FrontSeqNo *string `json:"FrontSeqNo,omitempty" name:"FrontSeqNo"`

		// STRING(1027)，保留域
		ReservedMsg *string `json:"ReservedMsg,omitempty" name:"ReservedMsg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckAmountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckAmountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClearItem struct {

	// STRING(8)，日期（格式: 20190101）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Date *string `json:"Date,omitempty" name:"Date"`

	// STRING(40)，子账号类型（子帐号类型。1: 普通会员子账号; 2: 挂账子账号; 3: 手续费子账号; 4: 利息子账号; 5: 平台担保子账号; 7: 在途; 8: 理财购买子帐号; 9: 理财赎回子帐号; 10: 平台子拥有结算子帐号）
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubAcctType *string `json:"SubAcctType,omitempty" name:"SubAcctType"`

	// STRING(3)，对账状态（0: 成功; 1: 失败）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReconcileStatus *string `json:"ReconcileStatus,omitempty" name:"ReconcileStatus"`

	// STRING(300)，对账返回信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReconcileReturnMsg *string `json:"ReconcileReturnMsg,omitempty" name:"ReconcileReturnMsg"`

	// STRING(20)，清算状态（0: 成功; 1: 失败; 2: 异常; 3: 待处理）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClearingStatus *string `json:"ClearingStatus,omitempty" name:"ClearingStatus"`

	// STRING(2)，清算返回信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClearingReturnMsg *string `json:"ClearingReturnMsg,omitempty" name:"ClearingReturnMsg"`

	// STRING(300)，待清算总金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalAmt *string `json:"TotalAmt,omitempty" name:"TotalAmt"`
}

type CloseOrderRequest struct {
	*tchttp.BaseRequest

	// 聚鑫分配的支付主MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 用户ID，长度不小于5位， 仅支持字母和数字的组合
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 聚鑫分配的安全ID
	MidasSecretId *string `json:"MidasSecretId,omitempty" name:"MidasSecretId"`

	// 按照聚鑫安全密钥计算的签名
	MidasSignature *string `json:"MidasSignature,omitempty" name:"MidasSignature"`

	// 业务订单号，OutTradeNo ， TransactionId二选一，不能都为空,优先使用 OutTradeNo
	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`

	// 聚鑫订单号，OutTradeNo ， TransactionId二选一，不能都为空,优先使用 OutTradeNo
	TransactionId *string `json:"TransactionId,omitempty" name:"TransactionId"`

	// 环境名:
	// release: 现网环境
	// sandbox: 沙箱环境
	// development: 开发环境
	// 缺省: release
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`
}

func (r *CloseOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MidasAppId")
	delete(f, "UserId")
	delete(f, "MidasSecretId")
	delete(f, "MidasSignature")
	delete(f, "OutTradeNo")
	delete(f, "TransactionId")
	delete(f, "MidasEnvironment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CloseOrderResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CloseOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfirmOrderRequest struct {
	*tchttp.BaseRequest

	// 分配给商户的AppId
	MerchantAppId *string `json:"MerchantAppId,omitempty" name:"MerchantAppId"`

	// 平台流水号。消费订单发起成功后，返回的平台唯一订单号。
	OrderNo *string `json:"OrderNo,omitempty" name:"OrderNo"`
}

func (r *ConfirmOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConfirmOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MerchantAppId")
	delete(f, "OrderNo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ConfirmOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ConfirmOrderResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分配给商户的AppId
		MerchantAppId *string `json:"MerchantAppId,omitempty" name:"MerchantAppId"`

		// 平台流水号。消费订单发起成功后，返回的平台唯一订单号。
		OrderNo *string `json:"OrderNo,omitempty" name:"OrderNo"`

		// 订单确认状态。0-确认失败
	// 1-确认成功 
	// 2-可疑状态
		Status *string `json:"Status,omitempty" name:"Status"`

		// 订单确认状态描述
	// 注意：此字段可能返回 null，表示取不到有效值。
		Description *string `json:"Description,omitempty" name:"Description"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ConfirmOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConfirmOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ContractInfo struct {

	// 米大师内部签约商户号
	ChannelContractMerchantId *string `json:"ChannelContractMerchantId,omitempty" name:"ChannelContractMerchantId"`

	// 米大师内部签约子商户号
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelContractSubMerchantId *string `json:"ChannelContractSubMerchantId,omitempty" name:"ChannelContractSubMerchantId"`

	// 米大师内部签约应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelContractAppId *string `json:"ChannelContractAppId,omitempty" name:"ChannelContractAppId"`

	// 米大师内部签约子应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelContractSubAppId *string `json:"ChannelContractSubAppId,omitempty" name:"ChannelContractSubAppId"`

	// 业务合约协议号
	OutContractCode *string `json:"OutContractCode,omitempty" name:"OutContractCode"`

	// 第三方渠道用户信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalContractUserInfoList []*ExternalContractUserInfo `json:"ExternalContractUserInfoList,omitempty" name:"ExternalContractUserInfoList"`

	// 签约方式，如 wechat_app ，使用app方式下的微信签
	ContractMethod *string `json:"ContractMethod,omitempty" name:"ContractMethod"`

	// 合约场景id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContractSceneId *string `json:"ContractSceneId,omitempty" name:"ContractSceneId"`

	// 用户信息
	UserInfo *ContractUserInfo `json:"UserInfo,omitempty" name:"UserInfo"`

	// 第三方渠道签约数据
	ExternalContractData *string `json:"ExternalContractData,omitempty" name:"ExternalContractData"`
}

type ContractOrderInSubOrder struct {

	// 子订单结算应收金额，单位： 分
	SubMchIncome *int64 `json:"SubMchIncome,omitempty" name:"SubMchIncome"`

	// 子订单平台应收金额，单位：分
	PlatformIncome *int64 `json:"PlatformIncome,omitempty" name:"PlatformIncome"`

	// 子订单商品详情
	ProductDetail *string `json:"ProductDetail,omitempty" name:"ProductDetail"`

	// 子订单商品名称
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 聚鑫计费SubAppId，代表子商户
	SubAppId *string `json:"SubAppId,omitempty" name:"SubAppId"`

	// 子订单号
	SubOutTradeNo *string `json:"SubOutTradeNo,omitempty" name:"SubOutTradeNo"`

	// 子订单支付金额
	Amt *int64 `json:"Amt,omitempty" name:"Amt"`

	// 子订单原始金额
	OriginalAmt *int64 `json:"OriginalAmt,omitempty" name:"OriginalAmt"`

	// 发货标识，由业务在调用聚鑫下单接口的 时候下发
	Metadata *string `json:"Metadata,omitempty" name:"Metadata"`
}

type ContractOrderRequest struct {
	*tchttp.BaseRequest

	// ISO 货币代码，CNY
	CurrencyType *string `json:"CurrencyType,omitempty" name:"CurrencyType"`

	// 聚鑫分配的支付主MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 支付订单号，仅支持数字、字母、下划线（_）、横杠字符（-）、点（.）的组合
	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`

	// 商品详情，需要URL编码
	ProductDetail *string `json:"ProductDetail,omitempty" name:"ProductDetail"`

	// 商品ID，仅支持数字、字母、下划线（_）、横杠字符（-）、点（.）的组合
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 商品名称，需要URL编码
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 支付金额，单位： 分
	TotalAmt *int64 `json:"TotalAmt,omitempty" name:"TotalAmt"`

	// 用户ID，长度不小于5位，仅支持字母和数字的组合
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 银行真实渠道.如:bank_pingan
	RealChannel *string `json:"RealChannel,omitempty" name:"RealChannel"`

	// 原始金额
	OriginalAmt *int64 `json:"OriginalAmt,omitempty" name:"OriginalAmt"`

	// 聚鑫分配的安全ID
	MidasSecretId *string `json:"MidasSecretId,omitempty" name:"MidasSecretId"`

	// 按照聚鑫安全密钥计算的签名
	MidasSignature *string `json:"MidasSignature,omitempty" name:"MidasSignature"`

	// 签约通知地址
	ContractNotifyUrl *string `json:"ContractNotifyUrl,omitempty" name:"ContractNotifyUrl"`

	// Web端回调地址
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 指定支付渠道：  wechat：微信支付  qqwallet：QQ钱包 
	//  bank：网银支付  只有一个渠道时需要指定
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// 透传字段，支付成功回调透传给应用，用于业务透传自定义内容
	Metadata *string `json:"Metadata,omitempty" name:"Metadata"`

	// 购买数量，不传默认为1
	Quantity *int64 `json:"Quantity,omitempty" name:"Quantity"`

	// 聚鑫计费SubAppId，代表子商户
	SubAppId *string `json:"SubAppId,omitempty" name:"SubAppId"`

	// 子订单信息列表，格式：子订单号、子应用ID、金额。 压缩后最长不可超过65535字节(去除空格，换行，制表符等无意义字符)
	// 注：接入银行或其他支付渠道服务商模式下，必传
	SubOrderList []*ContractOrderInSubOrder `json:"SubOrderList,omitempty" name:"SubOrderList"`

	// 结算应收金额，单位：分
	TotalMchIncome *int64 `json:"TotalMchIncome,omitempty" name:"TotalMchIncome"`

	// 平台应收金额，单位：分
	TotalPlatformIncome *int64 `json:"TotalPlatformIncome,omitempty" name:"TotalPlatformIncome"`

	// 微信公众号/小程序支付时为必选，需要传微信下的openid
	WxOpenId *string `json:"WxOpenId,omitempty" name:"WxOpenId"`

	// 在服务商模式下，微信公众号/小程序支付时wx_sub_openid和wx_openid二选一
	WxSubOpenId *string `json:"WxSubOpenId,omitempty" name:"WxSubOpenId"`

	// 环境名:
	// release: 现网环境
	// sandbox: 沙箱环境
	// development: 开发环境
	// 缺省: release
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`

	// 微信商户应用ID
	WxAppId *string `json:"WxAppId,omitempty" name:"WxAppId"`

	// 微信商户子应用ID
	WxSubAppId *string `json:"WxSubAppId,omitempty" name:"WxSubAppId"`

	// 支付通知地址
	PaymentNotifyUrl *string `json:"PaymentNotifyUrl,omitempty" name:"PaymentNotifyUrl"`

	// 传入调用方在Midas注册签约信息时获得的ContractSceneId。若未在Midas注册签约信息，则传入ExternalContractData。注意：ContractSceneId与ExternalContractData必须二选一传入其中一个
	ContractSceneId *string `json:"ContractSceneId,omitempty" name:"ContractSceneId"`

	// 需要按照各个渠道的扩展签约信息规范组装好该字段。若未在Midas注册签约信息，则传入该字段。注意：ContractSceneId与ExternalContractData必须二选一传入其中一个
	ExternalContractData *string `json:"ExternalContractData,omitempty" name:"ExternalContractData"`

	// 外部签约协议号，唯一标记一个签约关系。仅支持数字、字母、下划线（_）、横杠字符（-）、点（.）的组合
	OutContractCode *string `json:"OutContractCode,omitempty" name:"OutContractCode"`

	// 透传给第三方渠道的附加数据
	AttachData *string `json:"AttachData,omitempty" name:"AttachData"`

	// 展示用的签约用户名称，若不传入时，默认取UserId
	ContractDisplayName *string `json:"ContractDisplayName,omitempty" name:"ContractDisplayName"`
}

func (r *ContractOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ContractOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CurrencyType")
	delete(f, "MidasAppId")
	delete(f, "OutTradeNo")
	delete(f, "ProductDetail")
	delete(f, "ProductId")
	delete(f, "ProductName")
	delete(f, "TotalAmt")
	delete(f, "UserId")
	delete(f, "RealChannel")
	delete(f, "OriginalAmt")
	delete(f, "MidasSecretId")
	delete(f, "MidasSignature")
	delete(f, "ContractNotifyUrl")
	delete(f, "CallbackUrl")
	delete(f, "Channel")
	delete(f, "Metadata")
	delete(f, "Quantity")
	delete(f, "SubAppId")
	delete(f, "SubOrderList")
	delete(f, "TotalMchIncome")
	delete(f, "TotalPlatformIncome")
	delete(f, "WxOpenId")
	delete(f, "WxSubOpenId")
	delete(f, "MidasEnvironment")
	delete(f, "WxAppId")
	delete(f, "WxSubAppId")
	delete(f, "PaymentNotifyUrl")
	delete(f, "ContractSceneId")
	delete(f, "ExternalContractData")
	delete(f, "OutContractCode")
	delete(f, "AttachData")
	delete(f, "ContractDisplayName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ContractOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ContractOrderResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 支付金额，单位： 分
		TotalAmt *int64 `json:"TotalAmt,omitempty" name:"TotalAmt"`

		// 应用支付订单号
		OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`

		// 支付参数透传给聚鑫SDK（原文透传给SDK即可，不需要解码）
		PayInfo *string `json:"PayInfo,omitempty" name:"PayInfo"`

		// 聚鑫的交易订单号
		TransactionId *string `json:"TransactionId,omitempty" name:"TransactionId"`

		// 外部签约协议号
		OutContractCode *string `json:"OutContractCode,omitempty" name:"OutContractCode"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ContractOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ContractOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ContractSyncInfo struct {

	// 第三方渠道合约信息
	ExternalReturnContractInfo *ExternalReturnContractInfo `json:"ExternalReturnContractInfo,omitempty" name:"ExternalReturnContractInfo"`

	// 第三方渠道用户信息
	ExternalContractUserInfo []*ExternalContractUserInfo `json:"ExternalContractUserInfo,omitempty" name:"ExternalContractUserInfo"`

	// 签约方式，枚举值，
	// <br/>CONTRACT_METHOD_WECHAT_INVALID: 无效
	// CONTRACT_METHOD_WECHAT_APP: 微信APP
	// CONTRACT_METHOD_WECHAT_PUBLIC: 微信公众号
	// CONTRACT_METHOD_WECHAT_MINIPROGRAM: 微信小程序
	// CONTRACT_METHOD_WECHAT_H5: 微信H5
	ContractMethod *string `json:"ContractMethod,omitempty" name:"ContractMethod"`

	// 在米大师侧分配的场景id
	ContractSceneId *string `json:"ContractSceneId,omitempty" name:"ContractSceneId"`

	// 调用方从第三方渠道查询到的签约数据，由各个渠道定义
	ExternalReturnContractData *string `json:"ExternalReturnContractData,omitempty" name:"ExternalReturnContractData"`
}

type ContractUserInfo struct {

	// USER_ID: 用户ID
	// ANONYMOUS: 匿名类型用户ID
	UserType *string `json:"UserType,omitempty" name:"UserType"`

	// 用户类型
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

type CreateAcctRequest struct {
	*tchttp.BaseRequest

	// 聚鑫平台分配的支付MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 业务平台的子商户ID，唯一
	SubMchId *string `json:"SubMchId,omitempty" name:"SubMchId"`

	// 子商户名称
	SubMchName *string `json:"SubMchName,omitempty" name:"SubMchName"`

	// 子商户地址
	Address *string `json:"Address,omitempty" name:"Address"`

	// 子商户联系人
	// <敏感信息>加密详见<a href="https://cloud.tencent.com/document/product/1122/48979" target="_blank">《商户端接口敏感信息加密说明》</a>
	Contact *string `json:"Contact,omitempty" name:"Contact"`

	// 联系人手机号
	// <敏感信息>加密详见<a href="https://cloud.tencent.com/document/product/1122/48979" target="_blank">《商户端接口敏感信息加密说明》</a>
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 邮箱 
	// <敏感信息>加密详见<a href="https://cloud.tencent.com/document/product/1122/48979" target="_blank">《商户端接口敏感信息加密说明》</a>
	Email *string `json:"Email,omitempty" name:"Email"`

	// 聚鑫分配的安全ID
	MidasSecretId *string `json:"MidasSecretId,omitempty" name:"MidasSecretId"`

	// 按照聚鑫安全密钥计算的签名
	MidasSignature *string `json:"MidasSignature,omitempty" name:"MidasSignature"`

	// 子商户类型：
	// 个人: personal
	// 企业: enterprise
	// 个体工商户: individual
	// 缺省: enterprise
	SubMchType *string `json:"SubMchType,omitempty" name:"SubMchType"`

	// 不填则默认子商户名称
	ShortName *string `json:"ShortName,omitempty" name:"ShortName"`

	// 子商户会员类型：
	// general: 普通子账户
	// merchant: 商户子账户
	// 缺省: general
	SubMerchantMemberType *string `json:"SubMerchantMemberType,omitempty" name:"SubMerchantMemberType"`

	// 子商户密钥
	// <敏感信息>加密详见<a href="https://cloud.tencent.com/document/product/1122/48979" target="_blank">《商户端接口敏感信息加密说明》</a>
	SubMerchantKey *string `json:"SubMerchantKey,omitempty" name:"SubMerchantKey"`

	// 子商户私钥
	// <敏感信息>加密详见<a href="https://cloud.tencent.com/document/product/1122/48979" target="_blank">《商户端接口敏感信息加密说明》</a>
	SubMerchantPrivateKey *string `json:"SubMerchantPrivateKey,omitempty" name:"SubMerchantPrivateKey"`

	// 敏感信息加密类型:
	// RSA: rsa非对称加密，使用RSA-PKCS1-v1_5
	// AES: aes对称加密，使用AES256-CBC-PCKS7padding
	// 缺省: RSA
	EncryptType *string `json:"EncryptType,omitempty" name:"EncryptType"`

	// 银行生成的子商户账户，已开户的场景需要录入
	SubAcctNo *string `json:"SubAcctNo,omitempty" name:"SubAcctNo"`

	// 环境名:
	// release: 现网环境
	// sandbox: 沙箱环境
	// development: 开发环境
	// 缺省: release
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`

	// 店铺名称
	// 企业、个体工商户必输
	SubMerchantStoreName *string `json:"SubMerchantStoreName,omitempty" name:"SubMerchantStoreName"`

	// 公司信息
	OrganizationInfo *OrganizationInfo `json:"OrganizationInfo,omitempty" name:"OrganizationInfo"`
}

func (r *CreateAcctRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAcctRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MidasAppId")
	delete(f, "SubMchId")
	delete(f, "SubMchName")
	delete(f, "Address")
	delete(f, "Contact")
	delete(f, "Mobile")
	delete(f, "Email")
	delete(f, "MidasSecretId")
	delete(f, "MidasSignature")
	delete(f, "SubMchType")
	delete(f, "ShortName")
	delete(f, "SubMerchantMemberType")
	delete(f, "SubMerchantKey")
	delete(f, "SubMerchantPrivateKey")
	delete(f, "EncryptType")
	delete(f, "SubAcctNo")
	delete(f, "MidasEnvironment")
	delete(f, "SubMerchantStoreName")
	delete(f, "OrganizationInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAcctRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAcctResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 聚鑫计费SubAppId，代表子商户
		SubAppId *string `json:"SubAppId,omitempty" name:"SubAppId"`

		// 银行生成的子商户账户
		SubAcctNo *string `json:"SubAcctNo,omitempty" name:"SubAcctNo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAcctResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAcctResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAgentTaxPaymentInfosRequest struct {
	*tchttp.BaseRequest

	// 代理商ID
	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`

	// 平台渠道
	Channel *int64 `json:"Channel,omitempty" name:"Channel"`

	// 类型。0-视同，1-个体工商户
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 源电子凭证下载地址
	RawElectronicCertUrl *string `json:"RawElectronicCertUrl,omitempty" name:"RawElectronicCertUrl"`

	// 文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 完税信息
	AgentTaxPaymentInfos []*AgentTaxPayment `json:"AgentTaxPaymentInfos,omitempty" name:"AgentTaxPaymentInfos"`

	// 接入环境。沙箱环境填sandbox
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *CreateAgentTaxPaymentInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAgentTaxPaymentInfosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AgentId")
	delete(f, "Channel")
	delete(f, "Type")
	delete(f, "RawElectronicCertUrl")
	delete(f, "FileName")
	delete(f, "AgentTaxPaymentInfos")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAgentTaxPaymentInfosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAgentTaxPaymentInfosResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 代理商完税证明批次信息
		AgentTaxPaymentBatch *AgentTaxPaymentBatch `json:"AgentTaxPaymentBatch,omitempty" name:"AgentTaxPaymentBatch"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAgentTaxPaymentInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAgentTaxPaymentInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCustAcctIdRequest struct {
	*tchttp.BaseRequest

	// STRING(2)，功能标志（1: 开户; 3: 销户）
	FunctionFlag *string `json:"FunctionFlag,omitempty" name:"FunctionFlag"`

	// STRING(50)，资金汇总账号（即收单资金归集入账的账号）
	FundSummaryAcctNo *string `json:"FundSummaryAcctNo,omitempty" name:"FundSummaryAcctNo"`

	// STRING(32)，交易网会员代码（平台端的用户ID，需要保证唯一性，可数字字母混合，如HY_120）
	TranNetMemberCode *string `json:"TranNetMemberCode,omitempty" name:"TranNetMemberCode"`

	// STRING(10)，会员属性（00-普通子账户(默认); SH-商户子账户）
	MemberProperty *string `json:"MemberProperty,omitempty" name:"MemberProperty"`

	// STRING(30)，手机号码
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// String(22)，商户号（签约客户号）
	MrchCode *string `json:"MrchCode,omitempty" name:"MrchCode"`

	// String(2)，是否为自营业务（0位非自营，1为自营）
	SelfBusiness *bool `json:"SelfBusiness,omitempty" name:"SelfBusiness"`

	// String(64)，联系人
	ContactName *string `json:"ContactName,omitempty" name:"ContactName"`

	// String(64)，子账户名称
	SubAcctName *string `json:"SubAcctName,omitempty" name:"SubAcctName"`

	// String(64)，子账户简称
	SubAcctShortName *string `json:"SubAcctShortName,omitempty" name:"SubAcctShortName"`

	// String(4)，子账户类型（0: 个人子账户; 1: 企业子账户）
	SubAcctType *int64 `json:"SubAcctType,omitempty" name:"SubAcctType"`

	// STRING(150)，用户昵称
	UserNickname *string `json:"UserNickname,omitempty" name:"UserNickname"`

	// STRING(150)，邮箱
	Email *string `json:"Email,omitempty" name:"Email"`

	// STRING(1027)，保留域
	ReservedMsg *string `json:"ReservedMsg,omitempty" name:"ReservedMsg"`

	// STRING(12)，接入环境，默认接入沙箱环境。接入正式环境填"prod"
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *CreateCustAcctIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustAcctIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FunctionFlag")
	delete(f, "FundSummaryAcctNo")
	delete(f, "TranNetMemberCode")
	delete(f, "MemberProperty")
	delete(f, "Mobile")
	delete(f, "MrchCode")
	delete(f, "SelfBusiness")
	delete(f, "ContactName")
	delete(f, "SubAcctName")
	delete(f, "SubAcctShortName")
	delete(f, "SubAcctType")
	delete(f, "UserNickname")
	delete(f, "Email")
	delete(f, "ReservedMsg")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCustAcctIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateCustAcctIdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// STRING(50)，见证子账户的账号（平台需要记录下来，后续所有接口交互都会用到）
	// 注意：此字段可能返回 null，表示取不到有效值。
		SubAcctNo *string `json:"SubAcctNo,omitempty" name:"SubAcctNo"`

		// STRING(1027)，保留域（需要开通智能收款，此处返回智能收款账号，正常情况下返回空）
	// 注意：此字段可能返回 null，表示取不到有效值。
		ReservedMsg *string `json:"ReservedMsg,omitempty" name:"ReservedMsg"`

		// String(20)，返回码
		TxnReturnCode *string `json:"TxnReturnCode,omitempty" name:"TxnReturnCode"`

		// String(100)，返回信息
		TxnReturnMsg *string `json:"TxnReturnMsg,omitempty" name:"TxnReturnMsg"`

		// String(22)，交易流水号
		CnsmrSeqNo *string `json:"CnsmrSeqNo,omitempty" name:"CnsmrSeqNo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCustAcctIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustAcctIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateExternalAnchorData struct {

	// 主播Id
	AnchorId *string `json:"AnchorId,omitempty" name:"AnchorId"`
}

type CreateExternalAnchorRequest struct {
	*tchttp.BaseRequest

	// 平台业务系统唯一标示的主播id
	Uid *string `json:"Uid,omitempty" name:"Uid"`

	// 主播名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 身份证号
	IdNo *string `json:"IdNo,omitempty" name:"IdNo"`

	// 身份证正面图片下载链接
	IdCardFront *string `json:"IdCardFront,omitempty" name:"IdCardFront"`

	// 身份证反面图片下载链接
	IdCardReverse *string `json:"IdCardReverse,omitempty" name:"IdCardReverse"`
}

func (r *CreateExternalAnchorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateExternalAnchorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uid")
	delete(f, "Name")
	delete(f, "IdNo")
	delete(f, "IdCardFront")
	delete(f, "IdCardReverse")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateExternalAnchorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateExternalAnchorResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 错误码。响应成功："SUCCESS"，其他为不成功。
		ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

		// 响应消息。
		ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

		// 返回响应
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *CreateExternalAnchorData `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateExternalAnchorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateExternalAnchorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInvoiceItem struct {

	// 商品名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 税收商品编码
	TaxCode *string `json:"TaxCode,omitempty" name:"TaxCode"`

	// 不含税商品总价（商品含税价总额/（1+税率））。InvoicePlatformId 为1时，该金额为含税总金额。单位为分。
	TotalPrice *int64 `json:"TotalPrice,omitempty" name:"TotalPrice"`

	// 商品税率
	TaxRate *int64 `json:"TaxRate,omitempty" name:"TaxRate"`

	// 商品税额（不含税商品总价*税率）。单位为分
	TaxAmount *int64 `json:"TaxAmount,omitempty" name:"TaxAmount"`

	// 税收商品类别
	TaxType *string `json:"TaxType,omitempty" name:"TaxType"`

	// 商品规格
	Models *string `json:"Models,omitempty" name:"Models"`

	// 商品单位
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// 商品数量
	Total *string `json:"Total,omitempty" name:"Total"`

	// 不含税商品单价。InvoicePlatformId 为1时，该金额为含税单价。
	Price *string `json:"Price,omitempty" name:"Price"`

	// 含税折扣总额。单位为分
	Discount *int64 `json:"Discount,omitempty" name:"Discount"`

	// 优惠政策标志。0：不使用优惠政策；1：使用优惠政策。
	PreferentialPolicyFlag *string `json:"PreferentialPolicyFlag,omitempty" name:"PreferentialPolicyFlag"`

	// 零税率标识：
	// 空：非零税率；
	// 0：出口零税率；
	// 1：免税；
	// 2：不征税；
	// 3：普通零税率。
	ZeroTaxFlag *string `json:"ZeroTaxFlag,omitempty" name:"ZeroTaxFlag"`

	// 增值税特殊管理。PreferentialPolicyFlag字段为1时，必填。目前仅支持5%(3%，2%，1.5%)简易征税、免税、不征税。
	VatSpecialManagement *string `json:"VatSpecialManagement,omitempty" name:"VatSpecialManagement"`
}

type CreateInvoiceRequest struct {
	*tchttp.BaseRequest

	// 开票平台ID。0：高灯，1：票易通
	InvoicePlatformId *int64 `json:"InvoicePlatformId,omitempty" name:"InvoicePlatformId"`

	// 抬头类型：1：个人/政府事业单位；2：企业
	TitleType *int64 `json:"TitleType,omitempty" name:"TitleType"`

	// 购方名称
	BuyerTitle *string `json:"BuyerTitle,omitempty" name:"BuyerTitle"`

	// 业务开票号
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 含税总金额（单位为分）
	AmountHasTax *int64 `json:"AmountHasTax,omitempty" name:"AmountHasTax"`

	// 总税额（单位为分）
	TaxAmount *int64 `json:"TaxAmount,omitempty" name:"TaxAmount"`

	// 不含税总金额（单位为分）。InvoicePlatformId 为1时，传默认值-1
	AmountWithoutTax *int64 `json:"AmountWithoutTax,omitempty" name:"AmountWithoutTax"`

	// 销方纳税人识别号
	SellerTaxpayerNum *string `json:"SellerTaxpayerNum,omitempty" name:"SellerTaxpayerNum"`

	// 销方名称。（不填默认读取商户注册时输入的信息）
	SellerName *string `json:"SellerName,omitempty" name:"SellerName"`

	// 销方地址。（不填默认读取商户注册时输入的信息）
	SellerAddress *string `json:"SellerAddress,omitempty" name:"SellerAddress"`

	// 销方电话。（不填默认读取商户注册时输入的信息）
	SellerPhone *string `json:"SellerPhone,omitempty" name:"SellerPhone"`

	// 销方银行名称。（不填默认读取商户注册时输入的信息）
	SellerBankName *string `json:"SellerBankName,omitempty" name:"SellerBankName"`

	// 销方银行账号。（不填默认读取商户注册时输入的信息）
	SellerBankAccount *string `json:"SellerBankAccount,omitempty" name:"SellerBankAccount"`

	// 购方纳税人识别号（购方票面信息）,若抬头类型为2时，必传
	BuyerTaxpayerNum *string `json:"BuyerTaxpayerNum,omitempty" name:"BuyerTaxpayerNum"`

	// 购方地址。开具专用发票时必填
	BuyerAddress *string `json:"BuyerAddress,omitempty" name:"BuyerAddress"`

	// 购方银行名称。开具专用发票时必填
	BuyerBankName *string `json:"BuyerBankName,omitempty" name:"BuyerBankName"`

	// 购方银行账号。开具专用发票时必填
	BuyerBankAccount *string `json:"BuyerBankAccount,omitempty" name:"BuyerBankAccount"`

	// 购方电话。开具专用发票时必填
	BuyerPhone *string `json:"BuyerPhone,omitempty" name:"BuyerPhone"`

	// 收票人邮箱。若填入，会收到发票推送邮件
	BuyerEmail *string `json:"BuyerEmail,omitempty" name:"BuyerEmail"`

	// 收票人手机号。若填入，会收到发票推送短信
	TakerPhone *string `json:"TakerPhone,omitempty" name:"TakerPhone"`

	// 开票类型：
	// 1：增值税专用发票；
	// 2：增值税普通发票；
	// 3：增值税电子发票；
	// 4：增值税卷式发票；
	// 5：区块链电子发票。
	// 若该字段不填，或值不为1-5，则认为开具”增值税电子发票”
	InvoiceType *int64 `json:"InvoiceType,omitempty" name:"InvoiceType"`

	// 发票结果回传地址
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 开票人姓名。（不填默认读取商户注册时输入的信息）
	Drawer *string `json:"Drawer,omitempty" name:"Drawer"`

	// 收款人姓名。（不填默认读取商户注册时输入的信息）
	Payee *string `json:"Payee,omitempty" name:"Payee"`

	// 复核人姓名。（不填默认读取商户注册时输入的信息）
	Checker *string `json:"Checker,omitempty" name:"Checker"`

	// 税盘号
	TerminalCode *string `json:"TerminalCode,omitempty" name:"TerminalCode"`

	// 征收方式。开具差额征税发票时必填2。开具普通征税发票时为空
	LevyMethod *string `json:"LevyMethod,omitempty" name:"LevyMethod"`

	// 差额征税扣除额（单位为分）
	Deduction *int64 `json:"Deduction,omitempty" name:"Deduction"`

	// 备注（票面信息）
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 项目商品明细
	Items []*CreateInvoiceItem `json:"Items,omitempty" name:"Items"`

	// 接入环境。沙箱环境填sandbox。
	Profile *string `json:"Profile,omitempty" name:"Profile"`

	// 撤销部分商品。0-不撤销，1-撤销
	UndoPart *int64 `json:"UndoPart,omitempty" name:"UndoPart"`

	// 订单下单时间（格式 YYYYMMDD）
	OrderDate *string `json:"OrderDate,omitempty" name:"OrderDate"`

	// 订单级别折扣（单位为分）
	Discount *int64 `json:"Discount,omitempty" name:"Discount"`

	// 门店编码
	StoreNo *string `json:"StoreNo,omitempty" name:"StoreNo"`

	// 开票渠道。0：APP渠道，1：线下渠道，2：小程序渠道。不填默认为APP渠道
	InvoiceChannel *int64 `json:"InvoiceChannel,omitempty" name:"InvoiceChannel"`
}

func (r *CreateInvoiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInvoiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InvoicePlatformId")
	delete(f, "TitleType")
	delete(f, "BuyerTitle")
	delete(f, "OrderId")
	delete(f, "AmountHasTax")
	delete(f, "TaxAmount")
	delete(f, "AmountWithoutTax")
	delete(f, "SellerTaxpayerNum")
	delete(f, "SellerName")
	delete(f, "SellerAddress")
	delete(f, "SellerPhone")
	delete(f, "SellerBankName")
	delete(f, "SellerBankAccount")
	delete(f, "BuyerTaxpayerNum")
	delete(f, "BuyerAddress")
	delete(f, "BuyerBankName")
	delete(f, "BuyerBankAccount")
	delete(f, "BuyerPhone")
	delete(f, "BuyerEmail")
	delete(f, "TakerPhone")
	delete(f, "InvoiceType")
	delete(f, "CallbackUrl")
	delete(f, "Drawer")
	delete(f, "Payee")
	delete(f, "Checker")
	delete(f, "TerminalCode")
	delete(f, "LevyMethod")
	delete(f, "Deduction")
	delete(f, "Remark")
	delete(f, "Items")
	delete(f, "Profile")
	delete(f, "UndoPart")
	delete(f, "OrderDate")
	delete(f, "Discount")
	delete(f, "StoreNo")
	delete(f, "InvoiceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInvoiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateInvoiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 发票开具结果
		Result *CreateInvoiceResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateInvoiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInvoiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInvoiceResult struct {

	// 错误消息
	Message *string `json:"Message,omitempty" name:"Message"`

	// 错误码
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// 数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *CreateInvoiceResultData `json:"Data,omitempty" name:"Data"`
}

type CreateInvoiceResultData struct {

	// 开票状态
	State *int64 `json:"State,omitempty" name:"State"`

	// 发票ID
	InvoiceId *string `json:"InvoiceId,omitempty" name:"InvoiceId"`

	// 业务开票号
	OrderSn *string `json:"OrderSn,omitempty" name:"OrderSn"`
}

type CreateInvoiceResultV2 struct {

	// 发票ID
	InvoiceId *string `json:"InvoiceId,omitempty" name:"InvoiceId"`
}

type CreateInvoiceV2Request struct {
	*tchttp.BaseRequest

	// 开票平台ID。0：高灯，1：票易通
	InvoicePlatformId *int64 `json:"InvoicePlatformId,omitempty" name:"InvoicePlatformId"`

	// 抬头类型：1：个人/政府事业单位；2：企业
	TitleType *int64 `json:"TitleType,omitempty" name:"TitleType"`

	// 购方名称
	BuyerTitle *string `json:"BuyerTitle,omitempty" name:"BuyerTitle"`

	// 业务开票号
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 含税总金额（单位为分）
	AmountHasTax *int64 `json:"AmountHasTax,omitempty" name:"AmountHasTax"`

	// 总税额（单位为分）
	TaxAmount *int64 `json:"TaxAmount,omitempty" name:"TaxAmount"`

	// 不含税总金额（单位为分）。InvoicePlatformId 为1时，传默认值-1
	AmountWithoutTax *int64 `json:"AmountWithoutTax,omitempty" name:"AmountWithoutTax"`

	// 销方纳税人识别号
	SellerTaxpayerNum *string `json:"SellerTaxpayerNum,omitempty" name:"SellerTaxpayerNum"`

	// 销方名称。（不填默认读取商户注册时输入的信息）
	SellerName *string `json:"SellerName,omitempty" name:"SellerName"`

	// 销方地址。（不填默认读取商户注册时输入的信息）
	SellerAddress *string `json:"SellerAddress,omitempty" name:"SellerAddress"`

	// 销方电话。（不填默认读取商户注册时输入的信息）
	SellerPhone *string `json:"SellerPhone,omitempty" name:"SellerPhone"`

	// 销方银行名称。（不填默认读取商户注册时输入的信息）
	SellerBankName *string `json:"SellerBankName,omitempty" name:"SellerBankName"`

	// 销方银行账号。（不填默认读取商户注册时输入的信息）
	SellerBankAccount *string `json:"SellerBankAccount,omitempty" name:"SellerBankAccount"`

	// 购方纳税人识别号（购方票面信息）,若抬头类型为2时，必传
	BuyerTaxpayerNum *string `json:"BuyerTaxpayerNum,omitempty" name:"BuyerTaxpayerNum"`

	// 购方地址。开具专用发票时必填
	BuyerAddress *string `json:"BuyerAddress,omitempty" name:"BuyerAddress"`

	// 购方银行名称。开具专用发票时必填
	BuyerBankName *string `json:"BuyerBankName,omitempty" name:"BuyerBankName"`

	// 购方银行账号。开具专用发票时必填
	BuyerBankAccount *string `json:"BuyerBankAccount,omitempty" name:"BuyerBankAccount"`

	// 购方电话。开具专用发票时必填
	BuyerPhone *string `json:"BuyerPhone,omitempty" name:"BuyerPhone"`

	// 收票人邮箱。若填入，会收到发票推送邮件
	BuyerEmail *string `json:"BuyerEmail,omitempty" name:"BuyerEmail"`

	// 收票人手机号。若填入，会收到发票推送短信
	TakerPhone *string `json:"TakerPhone,omitempty" name:"TakerPhone"`

	// 开票类型：
	// 1：增值税专用发票；
	// 2：增值税普通发票；
	// 3：增值税电子发票；
	// 4：增值税卷式发票；
	// 5：区块链电子发票。
	// 若该字段不填，或值不为1-5，则认为开具”增值税电子发票”
	InvoiceType *int64 `json:"InvoiceType,omitempty" name:"InvoiceType"`

	// 发票结果回传地址
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 开票人姓名。（不填默认读取商户注册时输入的信息）
	Drawer *string `json:"Drawer,omitempty" name:"Drawer"`

	// 收款人姓名。（不填默认读取商户注册时输入的信息）
	Payee *string `json:"Payee,omitempty" name:"Payee"`

	// 复核人姓名。（不填默认读取商户注册时输入的信息）
	Checker *string `json:"Checker,omitempty" name:"Checker"`

	// 税盘号
	TerminalCode *string `json:"TerminalCode,omitempty" name:"TerminalCode"`

	// 征收方式。开具差额征税发票时必填2。开具普通征税发票时为空
	LevyMethod *string `json:"LevyMethod,omitempty" name:"LevyMethod"`

	// 差额征税扣除额（单位为分）
	Deduction *int64 `json:"Deduction,omitempty" name:"Deduction"`

	// 备注（票面信息）
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 项目商品明细
	Items []*CreateInvoiceItem `json:"Items,omitempty" name:"Items"`

	// 接入环境。沙箱环境填sandbox。
	Profile *string `json:"Profile,omitempty" name:"Profile"`

	// 撤销部分商品。0-不撤销，1-撤销
	UndoPart *int64 `json:"UndoPart,omitempty" name:"UndoPart"`

	// 订单下单时间（格式 YYYYMMDD）
	OrderDate *string `json:"OrderDate,omitempty" name:"OrderDate"`

	// 订单级别折扣（单位为分）
	Discount *int64 `json:"Discount,omitempty" name:"Discount"`

	// 门店编码
	StoreNo *string `json:"StoreNo,omitempty" name:"StoreNo"`

	// 开票渠道。0：APP渠道，1：线下渠道，2：小程序渠道。不填默认为APP渠道
	InvoiceChannel *int64 `json:"InvoiceChannel,omitempty" name:"InvoiceChannel"`
}

func (r *CreateInvoiceV2Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInvoiceV2Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InvoicePlatformId")
	delete(f, "TitleType")
	delete(f, "BuyerTitle")
	delete(f, "OrderId")
	delete(f, "AmountHasTax")
	delete(f, "TaxAmount")
	delete(f, "AmountWithoutTax")
	delete(f, "SellerTaxpayerNum")
	delete(f, "SellerName")
	delete(f, "SellerAddress")
	delete(f, "SellerPhone")
	delete(f, "SellerBankName")
	delete(f, "SellerBankAccount")
	delete(f, "BuyerTaxpayerNum")
	delete(f, "BuyerAddress")
	delete(f, "BuyerBankName")
	delete(f, "BuyerBankAccount")
	delete(f, "BuyerPhone")
	delete(f, "BuyerEmail")
	delete(f, "TakerPhone")
	delete(f, "InvoiceType")
	delete(f, "CallbackUrl")
	delete(f, "Drawer")
	delete(f, "Payee")
	delete(f, "Checker")
	delete(f, "TerminalCode")
	delete(f, "LevyMethod")
	delete(f, "Deduction")
	delete(f, "Remark")
	delete(f, "Items")
	delete(f, "Profile")
	delete(f, "UndoPart")
	delete(f, "OrderDate")
	delete(f, "Discount")
	delete(f, "StoreNo")
	delete(f, "InvoiceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInvoiceV2Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateInvoiceV2Response struct {
	*tchttp.BaseResponse
	Response *struct {

		// 发票开具结果
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *CreateInvoiceResultV2 `json:"Result,omitempty" name:"Result"`

		// 错误码
		ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

		// 错误消息
		ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateInvoiceV2Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInvoiceV2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMerchantRequest struct {
	*tchttp.BaseRequest

	// 开票平台ID
	InvoicePlatformId *int64 `json:"InvoicePlatformId,omitempty" name:"InvoicePlatformId"`

	// 企业名称
	TaxpayerName *string `json:"TaxpayerName,omitempty" name:"TaxpayerName"`

	// 销方纳税人识别号
	TaxpayerNum *string `json:"TaxpayerNum,omitempty" name:"TaxpayerNum"`

	// 注册企业法定代表人名称
	LegalPersonName *string `json:"LegalPersonName,omitempty" name:"LegalPersonName"`

	// 联系人
	ContactsName *string `json:"ContactsName,omitempty" name:"ContactsName"`

	// 联系人手机号
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 不包含省市名称的地址
	Address *string `json:"Address,omitempty" name:"Address"`

	// 地区编码
	RegionCode *int64 `json:"RegionCode,omitempty" name:"RegionCode"`

	// 市（地区）名称
	CityName *string `json:"CityName,omitempty" name:"CityName"`

	// 开票人
	Drawer *string `json:"Drawer,omitempty" name:"Drawer"`

	// 税务登记证图片（Base64）字符串，需小于 3M
	TaxRegistrationCertificate *string `json:"TaxRegistrationCertificate,omitempty" name:"TaxRegistrationCertificate"`

	// 联系人邮箱地址
	Email *string `json:"Email,omitempty" name:"Email"`

	// 企业电话
	BusinessMobile *string `json:"BusinessMobile,omitempty" name:"BusinessMobile"`

	// 银行名称
	BankName *string `json:"BankName,omitempty" name:"BankName"`

	// 银行账号
	BankAccount *string `json:"BankAccount,omitempty" name:"BankAccount"`

	// 复核人
	Reviewer *string `json:"Reviewer,omitempty" name:"Reviewer"`

	// 收款人
	Payee *string `json:"Payee,omitempty" name:"Payee"`

	// 注册邀请码
	RegisterCode *string `json:"RegisterCode,omitempty" name:"RegisterCode"`

	// 不填默认为1，有效状态
	// 0：表示无效；
	// 1:表示有效；
	// 2:表示禁止开蓝票；
	// 3:表示禁止冲红。
	State *string `json:"State,omitempty" name:"State"`

	// 接收推送的消息地址
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 接入环境。沙箱环境填 sandbox。
	Profile *string `json:"Profile,omitempty" name:"Profile"`
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
	delete(f, "InvoicePlatformId")
	delete(f, "TaxpayerName")
	delete(f, "TaxpayerNum")
	delete(f, "LegalPersonName")
	delete(f, "ContactsName")
	delete(f, "Phone")
	delete(f, "Address")
	delete(f, "RegionCode")
	delete(f, "CityName")
	delete(f, "Drawer")
	delete(f, "TaxRegistrationCertificate")
	delete(f, "Email")
	delete(f, "BusinessMobile")
	delete(f, "BankName")
	delete(f, "BankAccount")
	delete(f, "Reviewer")
	delete(f, "Payee")
	delete(f, "RegisterCode")
	delete(f, "State")
	delete(f, "CallbackUrl")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMerchantRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateMerchantResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 商户注册结果
		Result *CreateMerchantResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateMerchantResult struct {

	// 状态码
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// 响应消息
	Message *string `json:"Message,omitempty" name:"Message"`

	// 创建商户结果数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *CreateMerchantResultData `json:"Data,omitempty" name:"Data"`
}

type CreateMerchantResultData struct {

	// 企业名称
	TaxpayerName *string `json:"TaxpayerName,omitempty" name:"TaxpayerName"`

	// 请求流水号
	SerialNo *string `json:"SerialNo,omitempty" name:"SerialNo"`

	// 纳税号
	TaxpayerNum *string `json:"TaxpayerNum,omitempty" name:"TaxpayerNum"`
}

type CreateOrderRequest struct {
	*tchttp.BaseRequest

	// 渠道编号。ZSB2B：招商银行B2B。
	ChannelCode *string `json:"ChannelCode,omitempty" name:"ChannelCode"`

	// 进件成功后返给商户方的 AppId。
	MerchantAppId *string `json:"MerchantAppId,omitempty" name:"MerchantAppId"`

	// 交易金额。单位：元
	Amount *string `json:"Amount,omitempty" name:"Amount"`

	// 商户流水号。商户唯一订单号由字母或数字组成。
	TraceNo *string `json:"TraceNo,omitempty" name:"TraceNo"`

	// 通知地址。商户接收交易结果的通知地址。
	NotifyUrl *string `json:"NotifyUrl,omitempty" name:"NotifyUrl"`

	// 返回地址。支付成功后，页面将跳 转返回到商户的该地址。
	ReturnUrl *string `json:"ReturnUrl,omitempty" name:"ReturnUrl"`
}

func (r *CreateOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelCode")
	delete(f, "MerchantAppId")
	delete(f, "Amount")
	delete(f, "TraceNo")
	delete(f, "NotifyUrl")
	delete(f, "ReturnUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrderResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 进件成功后返给商户方的AppId。
		MerchantAppId *string `json:"MerchantAppId,omitempty" name:"MerchantAppId"`

		// 商户流水号，商户唯一订单号由字母或数字组成。
		TraceNo *string `json:"TraceNo,omitempty" name:"TraceNo"`

		// 平台流水号，若下单成功则返回。
		OrderNo *string `json:"OrderNo,omitempty" name:"OrderNo"`

		// 支付页面跳转地址，若下单成功则返回。
		PayUrl *string `json:"PayUrl,omitempty" name:"PayUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePayMerchantRequest struct {
	*tchttp.BaseRequest

	// 平台编号
	PlatformCode *string `json:"PlatformCode,omitempty" name:"PlatformCode"`

	// 渠道方收款商户编号，由渠道方(银行)提 供。
	ChannelMerchantNo *string `json:"ChannelMerchantNo,omitempty" name:"ChannelMerchantNo"`

	// 是否需要向渠道进行 商户信息验证 1:验证
	// 0:不验证
	ChannelCheckFlag *string `json:"ChannelCheckFlag,omitempty" name:"ChannelCheckFlag"`

	// 收款商户名称
	MerchantName *string `json:"MerchantName,omitempty" name:"MerchantName"`

	// 是否开通 B2B 支付 1:开通 0:不开通 缺省:1
	BusinessPayFlag *string `json:"BusinessPayFlag,omitempty" name:"BusinessPayFlag"`
}

func (r *CreatePayMerchantRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePayMerchantRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PlatformCode")
	delete(f, "ChannelMerchantNo")
	delete(f, "ChannelCheckFlag")
	delete(f, "MerchantName")
	delete(f, "BusinessPayFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePayMerchantRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreatePayMerchantResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分配给商户的 AppId。该 AppId 为后续各项 交易的商户标识。
		MerchantAppId *string `json:"MerchantAppId,omitempty" name:"MerchantAppId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePayMerchantResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePayMerchantResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRedInvoiceItem struct {

	// 订单号
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 发票结果回传地址
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 业务开票号
	OrderSn *string `json:"OrderSn,omitempty" name:"OrderSn"`

	// 红字信息表编码
	RedSerialNo *string `json:"RedSerialNo,omitempty" name:"RedSerialNo"`

	// 门店编号
	StoreNo *string `json:"StoreNo,omitempty" name:"StoreNo"`
}

type CreateRedInvoiceRequest struct {
	*tchttp.BaseRequest

	// 开票平台ID
	InvoicePlatformId *int64 `json:"InvoicePlatformId,omitempty" name:"InvoicePlatformId"`

	// 红冲明细
	Invoices []*CreateRedInvoiceItem `json:"Invoices,omitempty" name:"Invoices"`

	// 接入环境。沙箱环境填 sandbox。
	Profile *string `json:"Profile,omitempty" name:"Profile"`

	// 开票渠道。0：线上渠道，1：线下渠道。不填默认为线上渠道
	InvoiceChannel *int64 `json:"InvoiceChannel,omitempty" name:"InvoiceChannel"`
}

func (r *CreateRedInvoiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRedInvoiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InvoicePlatformId")
	delete(f, "Invoices")
	delete(f, "Profile")
	delete(f, "InvoiceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRedInvoiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateRedInvoiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 红冲结果
		Result *CreateRedInvoiceResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRedInvoiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRedInvoiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRedInvoiceResult struct {

	// 错误消息
	Message *string `json:"Message,omitempty" name:"Message"`

	// 错误码
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// 红票数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*CreateRedInvoiceResultData `json:"Data,omitempty" name:"Data"`
}

type CreateRedInvoiceResultData struct {

	// 红冲状态码
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// 红冲状态消息
	Message *string `json:"Message,omitempty" name:"Message"`

	// 发票ID
	InvoiceId *string `json:"InvoiceId,omitempty" name:"InvoiceId"`

	// 业务开票号
	OrderSn *string `json:"OrderSn,omitempty" name:"OrderSn"`
}

type CreateRedInvoiceResultV2 struct {

	// 红票ID
	InvoiceId *string `json:"InvoiceId,omitempty" name:"InvoiceId"`
}

type CreateRedInvoiceV2Request struct {
	*tchttp.BaseRequest

	// 开票平台ID
	InvoicePlatformId *int64 `json:"InvoicePlatformId,omitempty" name:"InvoicePlatformId"`

	// 订单号
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 接入环境。沙箱环境填 sandbox。
	Profile *string `json:"Profile,omitempty" name:"Profile"`

	// 开票渠道。0：线上渠道，1：线下渠道。不填默认为线上渠道
	InvoiceChannel *int64 `json:"InvoiceChannel,omitempty" name:"InvoiceChannel"`
}

func (r *CreateRedInvoiceV2Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRedInvoiceV2Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InvoicePlatformId")
	delete(f, "OrderId")
	delete(f, "Profile")
	delete(f, "InvoiceChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRedInvoiceV2Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateRedInvoiceV2Response struct {
	*tchttp.BaseResponse
	Response *struct {

		// 红冲结果
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *CreateRedInvoiceResultV2 `json:"Result,omitempty" name:"Result"`

		// 错误码
		ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

		// 错误消息
		ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRedInvoiceV2Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRedInvoiceV2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSinglePayRequest struct {
	*tchttp.BaseRequest

	// 业务流水号，历史唯一
	SerialNumber *string `json:"SerialNumber,omitempty" name:"SerialNumber"`

	// 付方账户号
	PayAccountNumber *string `json:"PayAccountNumber,omitempty" name:"PayAccountNumber"`

	// 付方账户名称
	PayAccountName *string `json:"PayAccountName,omitempty" name:"PayAccountName"`

	// 金额
	Amount *int64 `json:"Amount,omitempty" name:"Amount"`

	// 收方账户号
	RecvAccountNumber *string `json:"RecvAccountNumber,omitempty" name:"RecvAccountNumber"`

	// 收方账户名称
	RecvAccountName *string `json:"RecvAccountName,omitempty" name:"RecvAccountName"`

	// 付方账户CNAPS号
	PayBankCnaps *string `json:"PayBankCnaps,omitempty" name:"PayBankCnaps"`

	// 付方账户银行大类，PayBankCnaps为空时必传（见常见问题-银企直连银行类型）
	PayBankType *string `json:"PayBankType,omitempty" name:"PayBankType"`

	// 付方账户银行所在省，PayBankCnaps为空时必传（见常见问题-银企直连省份枚举信息）
	PayBankProvince *string `json:"PayBankProvince,omitempty" name:"PayBankProvince"`

	// 付方账户银行所在地区，PayBankCnaps为空时必传（见常见问题-银企直连城市枚举信息）
	PayBankCity *string `json:"PayBankCity,omitempty" name:"PayBankCity"`

	// 收方账户CNAPS号
	RecvBankCnaps *string `json:"RecvBankCnaps,omitempty" name:"RecvBankCnaps"`

	// 收方账户银行大类，RecvBankCnaps为空时必传（见常见问题-银企直连银行类型）
	RecvBankType *string `json:"RecvBankType,omitempty" name:"RecvBankType"`

	// 收方账户银行所在省，RecvBankCnaps为空时必传（见常见问题-银企直连省份枚举信息）
	RecvBankProvince *string `json:"RecvBankProvince,omitempty" name:"RecvBankProvince"`

	// 收方账户银行所在地区，RecvBankCnaps为空时必传（见常见问题-银企直连城市枚举信息）
	RecvBankCity *string `json:"RecvBankCity,omitempty" name:"RecvBankCity"`

	// 收款方证件类型（见常见问题-银企直连证件类型枚举信息）
	RecvCertType *string `json:"RecvCertType,omitempty" name:"RecvCertType"`

	// 收款方证件号码
	RecvCertNo *string `json:"RecvCertNo,omitempty" name:"RecvCertNo"`

	// 摘要信息
	Summary *string `json:"Summary,omitempty" name:"Summary"`

	// 接入环境。沙箱环境填sandbox
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *CreateSinglePayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSinglePayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SerialNumber")
	delete(f, "PayAccountNumber")
	delete(f, "PayAccountName")
	delete(f, "Amount")
	delete(f, "RecvAccountNumber")
	delete(f, "RecvAccountName")
	delete(f, "PayBankCnaps")
	delete(f, "PayBankType")
	delete(f, "PayBankProvince")
	delete(f, "PayBankCity")
	delete(f, "RecvBankCnaps")
	delete(f, "RecvBankType")
	delete(f, "RecvBankProvince")
	delete(f, "RecvBankCity")
	delete(f, "RecvCertType")
	delete(f, "RecvCertNo")
	delete(f, "Summary")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSinglePayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateSinglePayResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回结果
		Result *CreateSinglePayResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSinglePayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSinglePayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSinglePayResult struct {

	// 受理状态（S：处理成功；F：处理失败）
	HandleStatus *string `json:"HandleStatus,omitempty" name:"HandleStatus"`

	// 受理状态描述
	HandleMsg *string `json:"HandleMsg,omitempty" name:"HandleMsg"`

	// 业务流水号，历史唯一
	// 注意：此字段可能返回 null，表示取不到有效值。
	SerialNo *string `json:"SerialNo,omitempty" name:"SerialNo"`

	// 银行指令流水
	// 注意：此字段可能返回 null，表示取不到有效值。
	BankSerialNo *string `json:"BankSerialNo,omitempty" name:"BankSerialNo"`

	// 付款状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayStatus *string `json:"PayStatus,omitempty" name:"PayStatus"`

	// 银行原始返回码
	// 注意：此字段可能返回 null，表示取不到有效值。
	BankRetCode *string `json:"BankRetCode,omitempty" name:"BankRetCode"`

	// 银行原始返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	BankRetMsg *string `json:"BankRetMsg,omitempty" name:"BankRetMsg"`
}

type CreateTransferBatchRequest struct {
	*tchttp.BaseRequest

	// 商户号。
	// 示例值：129284394
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 转账明细列表。
	// 发起批量转账的明细列表，最多三千笔
	TransferDetails []*TransferDetailRequest `json:"TransferDetails,omitempty" name:"TransferDetails"`

	// 直连商户appId。
	// 即商户号绑定的appid。
	// 示例值：wxf636efh567hg4356
	MerchantAppId *string `json:"MerchantAppId,omitempty" name:"MerchantAppId"`

	// 商家批次单号。
	// 商户系统内部的商家批次单号，此参数只能由数字、字母组成，商户系统内部唯一，UTF8编码，最多32个字符。
	// 示例值：plfk2020042013
	MerchantBatchNo *string `json:"MerchantBatchNo,omitempty" name:"MerchantBatchNo"`

	// 批次名称。
	// 批量转账的名称。
	// 示例值：2019年1月深圳分部报销单
	BatchName *string `json:"BatchName,omitempty" name:"BatchName"`

	// 转账说明。
	// UTF8编码，最多32个字符。
	// 示例值：2019年深圳分部报销单
	BatchRemark *string `json:"BatchRemark,omitempty" name:"BatchRemark"`

	// 转账总金额。
	// 转账金额，单位为分。
	// 示例值：4000000
	TotalAmount *uint64 `json:"TotalAmount,omitempty" name:"TotalAmount"`

	// 转账总笔数。
	// 一个转账批次最多允许发起三千笔转账。
	// 示例值：200
	TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

	// 环境名。
	// release: 现网环境
	// sandbox: 沙箱环境
	// development: 开发环境
	// 缺省: release
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *CreateTransferBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTransferBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MerchantId")
	delete(f, "TransferDetails")
	delete(f, "MerchantAppId")
	delete(f, "MerchantBatchNo")
	delete(f, "BatchName")
	delete(f, "BatchRemark")
	delete(f, "TotalAmount")
	delete(f, "TotalNum")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTransferBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateTransferBatchResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 商家批次单号。
	// 商户系统内部的商家批次单号，此参数只能由数字、字母组成，商户系统内部唯一，UTF8编码，最多32个字符。
	// 示例值：plfk2020042013
		MerchantBatchNo *string `json:"MerchantBatchNo,omitempty" name:"MerchantBatchNo"`

		// 微信批次单号。
	// 微信商家转账系统返回的唯一标识。
	// 示例值：1030000071100999991182020050700019480001
		BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

		// 批次受理成功时返回，遵循rfc3339标准格式。格式为YYYY-MM-DDTHH:mm:ss.sss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss.sss表示时分秒毫秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35.120+08:00表示北京时间2015年05月20日13点29分35秒。
	// 示例值：2015-05-20T13:29:35.120+08:00
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTransferBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTransferBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAgentTaxPaymentInfoRequest struct {
	*tchttp.BaseRequest

	// 批次号
	BatchNum *int64 `json:"BatchNum,omitempty" name:"BatchNum"`

	// 接入环境。沙箱环境填sandbox
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *DeleteAgentTaxPaymentInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAgentTaxPaymentInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BatchNum")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAgentTaxPaymentInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAgentTaxPaymentInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAgentTaxPaymentInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAgentTaxPaymentInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAgentTaxPaymentInfosRequest struct {
	*tchttp.BaseRequest

	// 批次号
	BatchNum *int64 `json:"BatchNum,omitempty" name:"BatchNum"`
}

func (r *DeleteAgentTaxPaymentInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAgentTaxPaymentInfosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BatchNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAgentTaxPaymentInfosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAgentTaxPaymentInfosResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAgentTaxPaymentInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAgentTaxPaymentInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeChargeDetailRequest struct {
	*tchttp.BaseRequest

	// 请求类型
	RequestType *string `json:"RequestType,omitempty" name:"RequestType"`

	// 商户号
	MerchantCode *string `json:"MerchantCode,omitempty" name:"MerchantCode"`

	// 支付渠道
	PayChannel *string `json:"PayChannel,omitempty" name:"PayChannel"`

	// 子渠道
	PayChannelSubId *int64 `json:"PayChannelSubId,omitempty" name:"PayChannelSubId"`

	// 原始交易订单号或者流水号
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 父账户账号，资金汇总账号
	BankAccountNumber *string `json:"BankAccountNumber,omitempty" name:"BankAccountNumber"`

	// 收单渠道类型
	AcquiringChannelType *string `json:"AcquiringChannelType,omitempty" name:"AcquiringChannelType"`

	// 平台短号(银行分配)
	PlatformShortNumber *string `json:"PlatformShortNumber,omitempty" name:"PlatformShortNumber"`

	// 聚鑫分配的安全ID
	MidasSecretId *string `json:"MidasSecretId,omitempty" name:"MidasSecretId"`

	// 聚鑫分配的支付主MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 计费签名
	MidasSignature *string `json:"MidasSignature,omitempty" name:"MidasSignature"`

	// 交易流水号
	TransSequenceNumber *string `json:"TransSequenceNumber,omitempty" name:"TransSequenceNumber"`

	// Midas环境参数
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`

	// 保留域
	ReservedMessage *string `json:"ReservedMessage,omitempty" name:"ReservedMessage"`
}

func (r *DescribeChargeDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChargeDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RequestType")
	delete(f, "MerchantCode")
	delete(f, "PayChannel")
	delete(f, "PayChannelSubId")
	delete(f, "OrderId")
	delete(f, "BankAccountNumber")
	delete(f, "AcquiringChannelType")
	delete(f, "PlatformShortNumber")
	delete(f, "MidasSecretId")
	delete(f, "MidasAppId")
	delete(f, "MidasSignature")
	delete(f, "TransSequenceNumber")
	delete(f, "MidasEnvironment")
	delete(f, "ReservedMessage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeChargeDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeChargeDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 交易状态 （0：成功，1：失败，2：异常,3:冲正，5：待处理）
		OrderStatus *string `json:"OrderStatus,omitempty" name:"OrderStatus"`

		// 交易金额
		OrderAmount *string `json:"OrderAmount,omitempty" name:"OrderAmount"`

		// 佣金费
		CommissionAmount *string `json:"CommissionAmount,omitempty" name:"CommissionAmount"`

		// 支付方式  0-冻结支付 1-普通支付
		PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

		// 交易日期
		OrderDate *string `json:"OrderDate,omitempty" name:"OrderDate"`

		// 交易时间
		OrderTime *string `json:"OrderTime,omitempty" name:"OrderTime"`

		// 订单实际转入见证子账户的名称
		OrderActualInSubAccountName *string `json:"OrderActualInSubAccountName,omitempty" name:"OrderActualInSubAccountName"`

		// 订单实际转入见证子账户的帐号
		OrderActualInSubAccountNumber *string `json:"OrderActualInSubAccountNumber,omitempty" name:"OrderActualInSubAccountNumber"`

		// 订单实际转入见证子账户的帐号
		OrderInSubAccountName *string `json:"OrderInSubAccountName,omitempty" name:"OrderInSubAccountName"`

		// 订单转入见证子账户的帐号
		OrderInSubAccountNumber *string `json:"OrderInSubAccountNumber,omitempty" name:"OrderInSubAccountNumber"`

		// 银行流水号
		FrontSequenceNumber *string `json:"FrontSequenceNumber,omitempty" name:"FrontSequenceNumber"`

		// 当充值失败时，返回交易失败原因
		FailMessage *string `json:"FailMessage,omitempty" name:"FailMessage"`

		// 请求类型
		RequestType *string `json:"RequestType,omitempty" name:"RequestType"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeChargeDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChargeDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrderStatusRequest struct {
	*tchttp.BaseRequest

	// 请求类型，此接口固定填：QueryOrderStatusReq
	RequestType *string `json:"RequestType,omitempty" name:"RequestType"`

	// 商户号
	MerchantCode *string `json:"MerchantCode,omitempty" name:"MerchantCode"`

	// 支付渠道
	PayChannel *string `json:"PayChannel,omitempty" name:"PayChannel"`

	// 子渠道
	PayChannelSubId *int64 `json:"PayChannelSubId,omitempty" name:"PayChannelSubId"`

	// 交易订单号或流水号，提现，充值或会员交易请求时的CnsmrSeqNo值
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 父账户账号，资金汇总账号
	BankAccountNumber *string `json:"BankAccountNumber,omitempty" name:"BankAccountNumber"`

	// 平台短号(银行分配)
	PlatformShortNumber *string `json:"PlatformShortNumber,omitempty" name:"PlatformShortNumber"`

	// 功能标志 0：会员间交易 1：提现 2：充值
	QueryType *string `json:"QueryType,omitempty" name:"QueryType"`

	// 银行流水号
	TransSequenceNumber *string `json:"TransSequenceNumber,omitempty" name:"TransSequenceNumber"`

	// 计费签名
	MidasSignature *string `json:"MidasSignature,omitempty" name:"MidasSignature"`

	// 聚鑫分配的支付主MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 聚鑫分配的安全ID
	MidasSecretId *string `json:"MidasSecretId,omitempty" name:"MidasSecretId"`

	// Midas环境参数
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`

	// 保留字段
	ReservedMessage *string `json:"ReservedMessage,omitempty" name:"ReservedMessage"`

	// 子账户账号 暂未使用
	BankSubAccountNumber *string `json:"BankSubAccountNumber,omitempty" name:"BankSubAccountNumber"`

	// 交易日期 暂未使用
	TransDate *string `json:"TransDate,omitempty" name:"TransDate"`
}

func (r *DescribeOrderStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrderStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RequestType")
	delete(f, "MerchantCode")
	delete(f, "PayChannel")
	delete(f, "PayChannelSubId")
	delete(f, "OrderId")
	delete(f, "BankAccountNumber")
	delete(f, "PlatformShortNumber")
	delete(f, "QueryType")
	delete(f, "TransSequenceNumber")
	delete(f, "MidasSignature")
	delete(f, "MidasAppId")
	delete(f, "MidasSecretId")
	delete(f, "MidasEnvironment")
	delete(f, "ReservedMessage")
	delete(f, "BankSubAccountNumber")
	delete(f, "TransDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrderStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrderStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 交易状态 （0：成功，1：失败，2：待确认, 5：待处理，6：处理中）
		OrderStatus *string `json:"OrderStatus,omitempty" name:"OrderStatus"`

		// 交易金额
		OrderAmount *string `json:"OrderAmount,omitempty" name:"OrderAmount"`

		// 交易日期
		OrderDate *string `json:"OrderDate,omitempty" name:"OrderDate"`

		// 交易时间
		OrderTime *string `json:"OrderTime,omitempty" name:"OrderTime"`

		// 转出子账户账号
		OutSubAccountNumber *string `json:"OutSubAccountNumber,omitempty" name:"OutSubAccountNumber"`

		// 转入子账户账号
		InSubAccountNumber *string `json:"InSubAccountNumber,omitempty" name:"InSubAccountNumber"`

		// 记账标志（1：登记挂账 2：支付 3：提现 4：清分5:下单预支付 6：确认并付款 7：退款 8：支付到平台 N:其他）
		BookingFlag *string `json:"BookingFlag,omitempty" name:"BookingFlag"`

		// 当交易失败时，返回交易失败原因
		FailMessage *string `json:"FailMessage,omitempty" name:"FailMessage"`

		// 请求类型
		RequestType *string `json:"RequestType,omitempty" name:"RequestType"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrderStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrderStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadBillRequest struct {
	*tchttp.BaseRequest

	// 请求下载对账单日期
	StateDate *string `json:"StateDate,omitempty" name:"StateDate"`

	// 聚鑫分配的MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 聚鑫分配的SecretId
	MidasSecretId *string `json:"MidasSecretId,omitempty" name:"MidasSecretId"`

	// 使用聚鑫安全密钥计算的签名
	MidasSignature *string `json:"MidasSignature,omitempty" name:"MidasSignature"`

	// 环境名:
	// release: 现网环境
	// sandbox: 沙箱环境
	// development: 开发环境
	// 缺省: release
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`
}

func (r *DownloadBillRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadBillRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StateDate")
	delete(f, "MidasAppId")
	delete(f, "MidasSecretId")
	delete(f, "MidasSignature")
	delete(f, "MidasEnvironment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DownloadBillRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DownloadBillResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 账单文件名
		FileName *string `json:"FileName,omitempty" name:"FileName"`

		// 账单文件的MD5值
		FileMD5 *string `json:"FileMD5,omitempty" name:"FileMD5"`

		// 账单文件的真实下载地址
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadBillResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadBillResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExecuteMemberTransactionRequest struct {
	*tchttp.BaseRequest

	// 请求类型此接口固定填：MemberTransactionReq
	RequestType *string `json:"RequestType,omitempty" name:"RequestType"`

	// 银行注册商户号
	MerchantCode *string `json:"MerchantCode,omitempty" name:"MerchantCode"`

	// 支付渠道
	PayChannel *string `json:"PayChannel,omitempty" name:"PayChannel"`

	// 子渠道
	PayChannelSubId *int64 `json:"PayChannelSubId,omitempty" name:"PayChannelSubId"`

	// 转出交易网会员代码
	OutTransNetMemberCode *string `json:"OutTransNetMemberCode,omitempty" name:"OutTransNetMemberCode"`

	// 转出见证子账户的户名
	OutSubAccountName *string `json:"OutSubAccountName,omitempty" name:"OutSubAccountName"`

	// 转入见证子账户的户名
	InSubAccountName *string `json:"InSubAccountName,omitempty" name:"InSubAccountName"`

	// 转出子账户账号
	OutSubAccountNumber *string `json:"OutSubAccountNumber,omitempty" name:"OutSubAccountNumber"`

	// 转入子账户账号
	InSubAccountNumber *string `json:"InSubAccountNumber,omitempty" name:"InSubAccountNumber"`

	// 父账户账号，资金汇总账号
	BankAccountNumber *string `json:"BankAccountNumber,omitempty" name:"BankAccountNumber"`

	// 货币单位 单位，1：元，2：角，3：分
	CurrencyUnit *string `json:"CurrencyUnit,omitempty" name:"CurrencyUnit"`

	// 币种
	CurrencyType *string `json:"CurrencyType,omitempty" name:"CurrencyType"`

	// 交易金额
	CurrencyAmount *string `json:"CurrencyAmount,omitempty" name:"CurrencyAmount"`

	// 订单号
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 聚鑫分配的支付主MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 聚鑫分配的安全ID
	MidasSecretId *string `json:"MidasSecretId,omitempty" name:"MidasSecretId"`

	// 计费签名
	MidasSignature *string `json:"MidasSignature,omitempty" name:"MidasSignature"`

	// 交易流水号 
	// 生成方式：用户短号+日期（6位）+ 随机编号（10位）例如：F088722005120904930798
	// 短号：F08872  日期： 200512   随机编号：0904930798
	TransSequenceNumber *string `json:"TransSequenceNumber,omitempty" name:"TransSequenceNumber"`

	// 转入交易网会员代码
	InTransNetMemberCode *string `json:"InTransNetMemberCode,omitempty" name:"InTransNetMemberCode"`

	// Midas环境标识 release 现网环境 sandbox 沙箱环境
	// development 开发环境
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`

	// 平台短号(银行分配)
	PlatformShortNumber *string `json:"PlatformShortNumber,omitempty" name:"PlatformShortNumber"`

	// 1：下单预支付 
	// 2：确认并付款
	// 3：退款
	// 6：直接支付T+1
	// 9：直接支付T+0
	TransType *string `json:"TransType,omitempty" name:"TransType"`

	// 交易手续费
	TransFee *string `json:"TransFee,omitempty" name:"TransFee"`

	// 保留域
	ReservedMessage *string `json:"ReservedMessage,omitempty" name:"ReservedMessage"`
}

func (r *ExecuteMemberTransactionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExecuteMemberTransactionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RequestType")
	delete(f, "MerchantCode")
	delete(f, "PayChannel")
	delete(f, "PayChannelSubId")
	delete(f, "OutTransNetMemberCode")
	delete(f, "OutSubAccountName")
	delete(f, "InSubAccountName")
	delete(f, "OutSubAccountNumber")
	delete(f, "InSubAccountNumber")
	delete(f, "BankAccountNumber")
	delete(f, "CurrencyUnit")
	delete(f, "CurrencyType")
	delete(f, "CurrencyAmount")
	delete(f, "OrderId")
	delete(f, "MidasAppId")
	delete(f, "MidasSecretId")
	delete(f, "MidasSignature")
	delete(f, "TransSequenceNumber")
	delete(f, "InTransNetMemberCode")
	delete(f, "MidasEnvironment")
	delete(f, "PlatformShortNumber")
	delete(f, "TransType")
	delete(f, "TransFee")
	delete(f, "ReservedMessage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExecuteMemberTransactionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ExecuteMemberTransactionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 请求类型
		RequestType *string `json:"RequestType,omitempty" name:"RequestType"`

		// 银行流水号
		FrontSequenceNumber *string `json:"FrontSequenceNumber,omitempty" name:"FrontSequenceNumber"`

		// 保留域
	// 注意：此字段可能返回 null，表示取不到有效值。
		ReservedMessage *string `json:"ReservedMessage,omitempty" name:"ReservedMessage"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExecuteMemberTransactionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExecuteMemberTransactionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExternalContractUserInfo struct {

	// 第三方用户类型，例如:  WX_OPENID, WX_SUB_OPENID,WX_PAYER_OPENID
	ExternalUserType *string `json:"ExternalUserType,omitempty" name:"ExternalUserType"`

	// 第三方用户ID
	ExternalUserId *string `json:"ExternalUserId,omitempty" name:"ExternalUserId"`
}

type ExternalReturnContractInfo struct {

	// 第三方渠道协议id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalReturnAgreementId *string `json:"ExternalReturnAgreementId,omitempty" name:"ExternalReturnAgreementId"`

	// 第三方渠道协议生效时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalReturnContractEffectiveTimestamp *string `json:"ExternalReturnContractEffectiveTimestamp,omitempty" name:"ExternalReturnContractEffectiveTimestamp"`

	// 第三方渠道协议解约时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalReturnContractTerminationTimestamp *string `json:"ExternalReturnContractTerminationTimestamp,omitempty" name:"ExternalReturnContractTerminationTimestamp"`

	// 平台合约状态
	// 协议状态，枚举值：
	// CONTRACT_STATUS_SIGNED：已签约
	// CONTRACT_STATUS_TERMINATED：未签约
	// CONTRACT_STATUS_PENDING：签约进行中
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalReturnContractStatus *string `json:"ExternalReturnContractStatus,omitempty" name:"ExternalReturnContractStatus"`

	// 第三方渠道请求序列号
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalReturnRequestId *string `json:"ExternalReturnRequestId,omitempty" name:"ExternalReturnRequestId"`

	// 第三方渠道协议签署时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalReturnContractSignedTimestamp *string `json:"ExternalReturnContractSignedTimestamp,omitempty" name:"ExternalReturnContractSignedTimestamp"`

	// 第三方渠道协议到期时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalReturnContractExpiredTimestamp *string `json:"ExternalReturnContractExpiredTimestamp,omitempty" name:"ExternalReturnContractExpiredTimestamp"`

	// 第三方渠道返回的合约数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalReturnContractData *string `json:"ExternalReturnContractData,omitempty" name:"ExternalReturnContractData"`

	// 第三方渠道解约备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalReturnContractTerminationRemark *string `json:"ExternalReturnContractTerminationRemark,omitempty" name:"ExternalReturnContractTerminationRemark"`

	// 第三方渠道协议解约方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalReturnContractTerminationMode *string `json:"ExternalReturnContractTerminationMode,omitempty" name:"ExternalReturnContractTerminationMode"`
}

type FileItem struct {

	// STRING(256)，文件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// STRING(120)，随机密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	RandomPassword *string `json:"RandomPassword,omitempty" name:"RandomPassword"`

	// STRING(512)，文件路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`

	// STRING(64)，提取码
	// 注意：此字段可能返回 null，表示取不到有效值。
	DrawCode *string `json:"DrawCode,omitempty" name:"DrawCode"`
}

type MerchantManagementList struct {

	// 企业名称。
	TaxpayerName *string `json:"TaxpayerName,omitempty" name:"TaxpayerName"`

	// 纳税人识别号(税号)	。
	TaxpayerNum *string `json:"TaxpayerNum,omitempty" name:"TaxpayerNum"`

	// 请求流水号。
	SerialNo *string `json:"SerialNo,omitempty" name:"SerialNo"`

	// 开票系统ID
	InvoicePlatformId *int64 `json:"InvoicePlatformId,omitempty" name:"InvoicePlatformId"`
}

type MerchantManagementResult struct {

	// 总数。
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 商户列表。
	List []*MerchantManagementList `json:"List,omitempty" name:"List"`
}

type MigrateOrderRefundQueryRequest struct {
	*tchttp.BaseRequest

	// 商户号
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 支付渠道，ALIPAY对应支付宝渠道；UNIONPAY对应银联渠道
	PayChannel *string `json:"PayChannel,omitempty" name:"PayChannel"`

	// 退款订单号，最长64位，仅支持数字、 字母
	RefundOrderId *string `json:"RefundOrderId,omitempty" name:"RefundOrderId"`

	// 退款流水号
	TradeSerialNo *string `json:"TradeSerialNo,omitempty" name:"TradeSerialNo"`

	// 接入环境。沙箱环境填 sandbox。
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *MigrateOrderRefundQueryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MigrateOrderRefundQueryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MerchantId")
	delete(f, "PayChannel")
	delete(f, "RefundOrderId")
	delete(f, "TradeSerialNo")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MigrateOrderRefundQueryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type MigrateOrderRefundQueryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 请求成功状态
		IsSuccess *bool `json:"IsSuccess,omitempty" name:"IsSuccess"`

		// 交易流水号
		TradeSerialNo *string `json:"TradeSerialNo,omitempty" name:"TradeSerialNo"`

		// 交易备注
		TradeMsg *string `json:"TradeMsg,omitempty" name:"TradeMsg"`

		// 交易状态：0=交易待处理；1=交易处理中；2=交易处理成功；3=交易失败；4=状态未知
		TradeStatus *int64 `json:"TradeStatus,omitempty" name:"TradeStatus"`

		// 第三方支付机构支付交易号
	// 注意：此字段可能返回 null，表示取不到有效值。
		ThirdChannelOrderId *string `json:"ThirdChannelOrderId,omitempty" name:"ThirdChannelOrderId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MigrateOrderRefundQueryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MigrateOrderRefundQueryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MigrateOrderRefundRequest struct {
	*tchttp.BaseRequest

	// 商户代码
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 支付渠道，ALIPAY对应支付宝渠道；UNIONPAY对应银联渠道
	PayChannel *string `json:"PayChannel,omitempty" name:"PayChannel"`

	// 正向支付商户订单号
	PayOrderId *string `json:"PayOrderId,omitempty" name:"PayOrderId"`

	// 退款订单号，最长64位，仅支持数字、 字母
	RefundOrderId *string `json:"RefundOrderId,omitempty" name:"RefundOrderId"`

	// 退款金额，单位：分。备注：改字段必须大于0 和小于10000000000的整数。
	RefundAmt *uint64 `json:"RefundAmt,omitempty" name:"RefundAmt"`

	// 第三方支付机构支付交易号
	ThirdChannelOrderId *string `json:"ThirdChannelOrderId,omitempty" name:"ThirdChannelOrderId"`

	// 原始支付金额，单位：分。备注：当该字段为空或者为0 时，系统会默认使用订单当 实付金额作为退款金额
	PayAmt *uint64 `json:"PayAmt,omitempty" name:"PayAmt"`

	// 接入环境。沙箱环境填 sandbox。
	Profile *string `json:"Profile,omitempty" name:"Profile"`

	// 退款原因
	RefundReason *string `json:"RefundReason,omitempty" name:"RefundReason"`
}

func (r *MigrateOrderRefundRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MigrateOrderRefundRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MerchantId")
	delete(f, "PayChannel")
	delete(f, "PayOrderId")
	delete(f, "RefundOrderId")
	delete(f, "RefundAmt")
	delete(f, "ThirdChannelOrderId")
	delete(f, "PayAmt")
	delete(f, "Profile")
	delete(f, "RefundReason")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MigrateOrderRefundRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type MigrateOrderRefundResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 请求成功状态
		IsSuccess *bool `json:"IsSuccess,omitempty" name:"IsSuccess"`

		// 退款流水号
		TradeSerialNo *string `json:"TradeSerialNo,omitempty" name:"TradeSerialNo"`

		// 交易备注
		TradeMsg *string `json:"TradeMsg,omitempty" name:"TradeMsg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MigrateOrderRefundResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MigrateOrderRefundResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAgentTaxPaymentInfoRequest struct {
	*tchttp.BaseRequest

	// 批次号
	BatchNum *int64 `json:"BatchNum,omitempty" name:"BatchNum"`

	// 新源电子凭证地址
	RawElectronicCertUrl *string `json:"RawElectronicCertUrl,omitempty" name:"RawElectronicCertUrl"`

	// 新的文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 接入环境。沙箱环境填sandbox
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *ModifyAgentTaxPaymentInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAgentTaxPaymentInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BatchNum")
	delete(f, "RawElectronicCertUrl")
	delete(f, "FileName")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAgentTaxPaymentInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAgentTaxPaymentInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 代理商完税证明批次信息
		AgentTaxPaymentBatch *AgentTaxPaymentBatch `json:"AgentTaxPaymentBatch,omitempty" name:"AgentTaxPaymentBatch"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAgentTaxPaymentInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAgentTaxPaymentInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBindedAccountRequest struct {
	*tchttp.BaseRequest

	// 主播Id
	AnchorId *string `json:"AnchorId,omitempty" name:"AnchorId"`

	// 1 微信企业付款 
	// 2 支付宝转账 
	// 3 平安银企直连代发转账
	TransferType *int64 `json:"TransferType,omitempty" name:"TransferType"`

	// 收款方标识。
	// 微信为open_id；
	// 支付宝为会员alipay_user_id;
	// 平安为收款方银行账号;
	AccountNo *string `json:"AccountNo,omitempty" name:"AccountNo"`

	// 手机号
	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`
}

func (r *ModifyBindedAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBindedAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AnchorId")
	delete(f, "TransferType")
	delete(f, "AccountNo")
	delete(f, "PhoneNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBindedAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBindedAccountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 错误码。响应成功："SUCCESS"，其他为不成功。
		ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

		// 响应消息。
		ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

		// 该字段为null。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBindedAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBindedAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMerchantRequest struct {
	*tchttp.BaseRequest

	// 进件成功后返给商户的AppId
	MerchantAppId *string `json:"MerchantAppId,omitempty" name:"MerchantAppId"`

	// 收款商户名称
	MerchantName *string `json:"MerchantName,omitempty" name:"MerchantName"`

	// B2B 支付标志。是否开通 B2B支付， 1:开通 0:不开通。
	BusinessPayFlag *string `json:"BusinessPayFlag,omitempty" name:"BusinessPayFlag"`
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
	delete(f, "MerchantAppId")
	delete(f, "MerchantName")
	delete(f, "BusinessPayFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMerchantRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMerchantResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ModifyMntMbrBindRelateAcctBankCodeRequest struct {
	*tchttp.BaseRequest

	// String(22)，商户号（签约客户号）
	MrchCode *string `json:"MrchCode,omitempty" name:"MrchCode"`

	// STRING(50)，见证子账户的账号
	SubAcctNo *string `json:"SubAcctNo,omitempty" name:"SubAcctNo"`

	// STRING(50)，会员绑定账号
	MemberBindAcctNo *string `json:"MemberBindAcctNo,omitempty" name:"MemberBindAcctNo"`

	// STRING(150)，开户行名称（若大小额行号不填则送超级网银号对应的银行名称，若填大小额行号则送大小额行号对应的银行名称）
	AcctOpenBranchName *string `json:"AcctOpenBranchName,omitempty" name:"AcctOpenBranchName"`

	// STRING(20)，大小额行号（CnapsBranchId和EiconBankBranchId两者二选一必填）
	CnapsBranchId *string `json:"CnapsBranchId,omitempty" name:"CnapsBranchId"`

	// STRING(20)，超级网银行号
	EiconBankBranchId *string `json:"EiconBankBranchId,omitempty" name:"EiconBankBranchId"`

	// STRING(1027)，保留域
	ReservedMsg *string `json:"ReservedMsg,omitempty" name:"ReservedMsg"`

	// STRING(12)，接入环境，默认接入沙箱环境。接入正式环境填"prod"
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *ModifyMntMbrBindRelateAcctBankCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMntMbrBindRelateAcctBankCodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MrchCode")
	delete(f, "SubAcctNo")
	delete(f, "MemberBindAcctNo")
	delete(f, "AcctOpenBranchName")
	delete(f, "CnapsBranchId")
	delete(f, "EiconBankBranchId")
	delete(f, "ReservedMsg")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMntMbrBindRelateAcctBankCodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMntMbrBindRelateAcctBankCodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// String(20)，返回码
		TxnReturnCode *string `json:"TxnReturnCode,omitempty" name:"TxnReturnCode"`

		// String(100)，返回信息
		TxnReturnMsg *string `json:"TxnReturnMsg,omitempty" name:"TxnReturnMsg"`

		// String(22)，交易流水号
		CnsmrSeqNo *string `json:"CnsmrSeqNo,omitempty" name:"CnsmrSeqNo"`

		// STRING(1027)，保留域
	// 注意：此字段可能返回 null，表示取不到有效值。
		ReservedMsg *string `json:"ReservedMsg,omitempty" name:"ReservedMsg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMntMbrBindRelateAcctBankCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMntMbrBindRelateAcctBankCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Order struct {

	// 含税金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	AmountHasTax *float64 `json:"AmountHasTax,omitempty" name:"AmountHasTax"`

	// 优惠金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	Discount *float64 `json:"Discount,omitempty" name:"Discount"`

	// 销方名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SellerName *string `json:"SellerName,omitempty" name:"SellerName"`

	// 发票类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvoiceType *int64 `json:"InvoiceType,omitempty" name:"InvoiceType"`

	// 默认“”
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 支付金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	Amount *float64 `json:"Amount,omitempty" name:"Amount"`

	// 下单日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderDate *string `json:"OrderDate,omitempty" name:"OrderDate"`

	// 订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 门店号
	// 注意：此字段可能返回 null，表示取不到有效值。
	StoreNo *string `json:"StoreNo,omitempty" name:"StoreNo"`

	// 明细
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*OrderItem `json:"Items,omitempty" name:"Items"`
}

type OrderItem struct {

	// 明细金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	AmountHasTax *float64 `json:"AmountHasTax,omitempty" name:"AmountHasTax"`

	// 优惠金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	Discount *float64 `json:"Discount,omitempty" name:"Discount"`

	// 商品名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 型号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Models *string `json:"Models,omitempty" name:"Models"`

	// 数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 数量单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// 默认“0”
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 单价
	// 注意：此字段可能返回 null，表示取不到有效值。
	Price *float64 `json:"Price,omitempty" name:"Price"`

	// 商品编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaxCode *string `json:"TaxCode,omitempty" name:"TaxCode"`
}

type OrganizationInfo struct {

	// 公司名称，个体工商户必输
	OrganizationName *string `json:"OrganizationName,omitempty" name:"OrganizationName"`

	// 公司证件类型，个体工商户必输，证件类型仅支持73
	OrganizationType *string `json:"OrganizationType,omitempty" name:"OrganizationType"`

	// 公司证件号码，个体工商户必输
	OrganizationCode *string `json:"OrganizationCode,omitempty" name:"OrganizationCode"`

	// 法人名称，如果SubMchName不是法人，需要另外送入法人信息（企业必输）
	// <敏感信息>加密详见<a href="https://cloud.tencent.com/document/product/1122/48979" target="_blank">《商户端接口敏感信息加密说明》</a>
	LegalPersonName *string `json:"LegalPersonName,omitempty" name:"LegalPersonName"`

	// 法人证件类型，如果SubMchName不是法人，需要另外送入法人信息（企业必输）
	LegalPersonIdType *string `json:"LegalPersonIdType,omitempty" name:"LegalPersonIdType"`

	// 法人证件号码，如果SubMchName不是法人，需要另外送入法人信息（企业必输）
	// <敏感信息>加密详见<a href="https://cloud.tencent.com/document/product/1122/48979" target="_blank">《商户端接口敏感信息加密说明》</a>
	LegalPersonIdCode *string `json:"LegalPersonIdCode,omitempty" name:"LegalPersonIdCode"`
}

type QueryAcctBindingRequest struct {
	*tchttp.BaseRequest

	// 聚鑫分配的支付主MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 聚鑫计费SubAppId，代表子商户
	SubAppId *string `json:"SubAppId,omitempty" name:"SubAppId"`

	// 由平台客服提供的计费密钥Id
	MidasSecretId *string `json:"MidasSecretId,omitempty" name:"MidasSecretId"`

	// 计费签名
	MidasSignature *string `json:"MidasSignature,omitempty" name:"MidasSignature"`

	// 敏感信息加密类型:
	// RSA: rsa非对称加密，使用RSA-PKCS1-v1_5
	// AES: aes对称加密，使用AES256-CBC-PCKS7padding
	// 缺省: RSA
	EncryptType *string `json:"EncryptType,omitempty" name:"EncryptType"`

	// 环境名:
	// release: 现网环境
	// sandbox: 沙箱环境
	// development: 开发环境
	// 缺省: release
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`
}

func (r *QueryAcctBindingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryAcctBindingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MidasAppId")
	delete(f, "SubAppId")
	delete(f, "MidasSecretId")
	delete(f, "MidasSignature")
	delete(f, "EncryptType")
	delete(f, "MidasEnvironment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryAcctBindingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryAcctBindingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总行数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 银行卡信息列表
		BankCardItems []*BankCardItem `json:"BankCardItems,omitempty" name:"BankCardItems"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryAcctBindingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryAcctBindingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryAcctInfoListRequest struct {
	*tchttp.BaseRequest

	// 聚鑫分配的支付主MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 查询开始时间(以开户时间为准)
	QueryAcctBeginTime *string `json:"QueryAcctBeginTime,omitempty" name:"QueryAcctBeginTime"`

	// 查询结束时间(以开户时间为准)
	QueryAcctEndTime *string `json:"QueryAcctEndTime,omitempty" name:"QueryAcctEndTime"`

	// 分页号,  起始值为1，每次最多返回20条记录，第二页返回的记录数为第21至40条记录，第三页为41至60条记录，顺序均按照开户时间的先后
	PageOffset *string `json:"PageOffset,omitempty" name:"PageOffset"`

	// 由平台客服提供的计费密钥Id
	MidasSecretId *string `json:"MidasSecretId,omitempty" name:"MidasSecretId"`

	// 计费签名
	MidasSignature *string `json:"MidasSignature,omitempty" name:"MidasSignature"`

	// 敏感信息加密类型:
	// RSA: rsa非对称加密，使用RSA-PKCS1-v1_5
	// AES: aes对称加密，使用AES256-CBC-PCKS7padding
	// 缺省: RSA
	EncryptType *string `json:"EncryptType,omitempty" name:"EncryptType"`

	// 环境名:
	// release: 现网环境
	// sandbox: 沙箱环境
	// development: 开发环境
	// 缺省: release
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`
}

func (r *QueryAcctInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryAcctInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MidasAppId")
	delete(f, "QueryAcctBeginTime")
	delete(f, "QueryAcctEndTime")
	delete(f, "PageOffset")
	delete(f, "MidasSecretId")
	delete(f, "MidasSignature")
	delete(f, "EncryptType")
	delete(f, "MidasEnvironment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryAcctInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryAcctInfoListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 本次交易返回查询结果记录数
		ResultCount *int64 `json:"ResultCount,omitempty" name:"ResultCount"`

		// 符合业务查询条件的记录总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 查询结果项 [object,object]
		QueryAcctItems []*QueryAcctItem `json:"QueryAcctItems,omitempty" name:"QueryAcctItems"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryAcctInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryAcctInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryAcctInfoRequest struct {
	*tchttp.BaseRequest

	// 聚鑫平台分配的支付MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 业务平台的子商户Id，唯一
	SubMchId *string `json:"SubMchId,omitempty" name:"SubMchId"`

	// 由平台客服提供的计费密钥Id
	MidasSecretId *string `json:"MidasSecretId,omitempty" name:"MidasSecretId"`

	// 计费签名
	MidasSignature *string `json:"MidasSignature,omitempty" name:"MidasSignature"`

	// 敏感信息加密类型:
	// RSA: rsa非对称加密，使用RSA-PKCS1-v1_5
	// AES: aes对称加密，使用AES256-CBC-PCKS7padding
	// 缺省: RSA
	EncryptType *string `json:"EncryptType,omitempty" name:"EncryptType"`

	// 环境名:
	// release: 现网环境
	// sandbox: 沙箱环境
	// development: 开发环境
	// 缺省: release
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`
}

func (r *QueryAcctInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryAcctInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MidasAppId")
	delete(f, "SubMchId")
	delete(f, "MidasSecretId")
	delete(f, "MidasSignature")
	delete(f, "EncryptType")
	delete(f, "MidasEnvironment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryAcctInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryAcctInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 聚鑫计费SubAppId，代表子商户
		SubAppId *string `json:"SubAppId,omitempty" name:"SubAppId"`

		// 子商户名称
		SubMchName *string `json:"SubMchName,omitempty" name:"SubMchName"`

		// 子商户类型：
	// 个人: personal
	// 企业：enterprise
	// 缺省： enterprise
		SubMchType *string `json:"SubMchType,omitempty" name:"SubMchType"`

		// 不填则默认子商户名称
		ShortName *string `json:"ShortName,omitempty" name:"ShortName"`

		// 子商户地址
		Address *string `json:"Address,omitempty" name:"Address"`

		// 子商户联系人子商户联系人
	// <敏感信息>
		Contact *string `json:"Contact,omitempty" name:"Contact"`

		// 联系人手机号
	// <敏感信息>
		Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

		// 邮箱 
	// <敏感信息>
		Email *string `json:"Email,omitempty" name:"Email"`

		// 子商户id
		SubMchId *string `json:"SubMchId,omitempty" name:"SubMchId"`

		// 子账户
		SubAcctNo *string `json:"SubAcctNo,omitempty" name:"SubAcctNo"`

		// 子商户会员类型：
	// general:普通子账户
	// merchant:商户子账户
	// 缺省： general
	// 注意：此字段可能返回 null，表示取不到有效值。
		SubMerchantMemberType *string `json:"SubMerchantMemberType,omitempty" name:"SubMerchantMemberType"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryAcctInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryAcctInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryAcctItem struct {

	// 子商户类型：
	// 个人: personal
	// 企业：enterprise
	// 缺省： enterprise
	SubMchType *string `json:"SubMchType,omitempty" name:"SubMchType"`

	// 子商户名称
	SubMchName *string `json:"SubMchName,omitempty" name:"SubMchName"`

	// 子账号号
	SubAcctNo *string `json:"SubAcctNo,omitempty" name:"SubAcctNo"`

	// 不填则默认子商户名称
	ShortName *string `json:"ShortName,omitempty" name:"ShortName"`

	// 业务平台的子商户Id，唯一
	SubMchId *string `json:"SubMchId,omitempty" name:"SubMchId"`

	// 聚鑫计费SubAppId，代表子商户
	SubAppId *string `json:"SubAppId,omitempty" name:"SubAppId"`

	// 子商户联系人
	// <敏感信息>
	Contact *string `json:"Contact,omitempty" name:"Contact"`

	// 子商户地址
	Address *string `json:"Address,omitempty" name:"Address"`

	// 联系人手机号
	// <敏感信息>
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 邮箱 
	// <敏感信息>
	Email *string `json:"Email,omitempty" name:"Email"`

	// 子商户会员类型：
	// general:普通子账户
	// merchant:商户子账户
	// 缺省： general
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubMerchantMemberType *string `json:"SubMerchantMemberType,omitempty" name:"SubMerchantMemberType"`
}

type QueryAgentStatementsRequest struct {
	*tchttp.BaseRequest

	// 结算单日期，月结算单填每月1日
	Date *string `json:"Date,omitempty" name:"Date"`

	// 日结算单:daily
	// 月结算单:monthly
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *QueryAgentStatementsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryAgentStatementsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Date")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryAgentStatementsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryAgentStatementsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 文件下载链接
	// 注意：此字段可能返回 null，表示取不到有效值。
		FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryAgentStatementsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryAgentStatementsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryAgentTaxPaymentBatchRequest struct {
	*tchttp.BaseRequest

	// 批次号
	BatchNum *int64 `json:"BatchNum,omitempty" name:"BatchNum"`

	// 接入环境。沙箱环境填sandbox
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *QueryAgentTaxPaymentBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryAgentTaxPaymentBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BatchNum")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryAgentTaxPaymentBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryAgentTaxPaymentBatchResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 代理商完税证明批次信息
		AgentTaxPaymentBatch *AgentTaxPaymentBatch `json:"AgentTaxPaymentBatch,omitempty" name:"AgentTaxPaymentBatch"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryAgentTaxPaymentBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryAgentTaxPaymentBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryAnchorContractInfoRequest struct {
	*tchttp.BaseRequest

	// 起始时间，格式为yyyy-MM-dd
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 起始时间，格式为yyyy-MM-dd
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *QueryAnchorContractInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryAnchorContractInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryAnchorContractInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryAnchorContractInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 签约主播数据
		AnchorContractInfoList []*AnchorContractInfo `json:"AnchorContractInfoList,omitempty" name:"AnchorContractInfoList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryAnchorContractInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryAnchorContractInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryApplicationMaterialRequest struct {
	*tchttp.BaseRequest

	// 申报流水号
	DeclareId *string `json:"DeclareId,omitempty" name:"DeclareId"`

	// 接入环境。沙箱环境填sandbox
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *QueryApplicationMaterialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryApplicationMaterialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeclareId")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryApplicationMaterialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryApplicationMaterialResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 成功申报材料查询结果
		Result *QueryDeclareResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryApplicationMaterialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryApplicationMaterialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryBalanceRequest struct {
	*tchttp.BaseRequest

	// 聚鑫分配的支付主MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 聚鑫计费SubAppId，代表子商户
	SubAppId *string `json:"SubAppId,omitempty" name:"SubAppId"`

	// 2：普通会员子账号
	// 3：功能子账号
	QueryFlag *string `json:"QueryFlag,omitempty" name:"QueryFlag"`

	// 起始值为1，每次最多返回20条记录，第二页返回的记录数为第21至40条记录，第三页为41至60条记录，顺序均按照建立时间的先后
	PageOffset *string `json:"PageOffset,omitempty" name:"PageOffset"`

	// 聚鑫分配的安全ID
	MidasSecretId *string `json:"MidasSecretId,omitempty" name:"MidasSecretId"`

	// 按照聚鑫安全密钥计算的签名
	MidasSignature *string `json:"MidasSignature,omitempty" name:"MidasSignature"`

	// 环境名:
	// release: 现网环境
	// sandbox: 沙箱环境
	// development: 开发环境
	// 缺省: release
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`
}

func (r *QueryBalanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryBalanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MidasAppId")
	delete(f, "SubAppId")
	delete(f, "QueryFlag")
	delete(f, "PageOffset")
	delete(f, "MidasSecretId")
	delete(f, "MidasSignature")
	delete(f, "MidasEnvironment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryBalanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryBalanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 本次交易返回查询结果记录数
		ResultCount *string `json:"ResultCount,omitempty" name:"ResultCount"`

		// 起始记录号
		StartRecordOffset *string `json:"StartRecordOffset,omitempty" name:"StartRecordOffset"`

		// 结束标志
		EndFlag *string `json:"EndFlag,omitempty" name:"EndFlag"`

		// 符合业务查询条件的记录总数
		TotalCount *string `json:"TotalCount,omitempty" name:"TotalCount"`

		// 查询结果项
		QueryItems []*QueryItem `json:"QueryItems,omitempty" name:"QueryItems"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryBalanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryBalanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryBankClearRequest struct {
	*tchttp.BaseRequest

	// String(22)，商户号（签约客户号）
	MrchCode *string `json:"MrchCode,omitempty" name:"MrchCode"`

	// STRING(2)，功能标志（1: 全部; 2: 指定时间段）
	FunctionFlag *string `json:"FunctionFlag,omitempty" name:"FunctionFlag"`

	// STRING (10)，页码（起始值为1，每次最多返回20条记录，第二页返回的记录数为第21至40条记录，第三页为41至60条记录，顺序均按照建立时间的先后）
	PageNum *string `json:"PageNum,omitempty" name:"PageNum"`

	// STRING(8)，开始日期（若是指定时间段查询，则必输，当查询全部时，不起作用。格式: 20190101）
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// STRING(8)，终止日期（若是指定时间段查询，则必输，当查询全部时，不起作用。格式：20190101）
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// STRING(1027)，保留域
	ReservedMsg *string `json:"ReservedMsg,omitempty" name:"ReservedMsg"`

	// STRING(12)，接入环境，默认接入沙箱环境。接入正式环境填"prod"
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *QueryBankClearRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryBankClearRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MrchCode")
	delete(f, "FunctionFlag")
	delete(f, "PageNum")
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "ReservedMsg")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryBankClearRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryBankClearResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// String(20)，返回码
		TxnReturnCode *string `json:"TxnReturnCode,omitempty" name:"TxnReturnCode"`

		// String(100)，返回信息
		TxnReturnMsg *string `json:"TxnReturnMsg,omitempty" name:"TxnReturnMsg"`

		// String(22)，交易流水号
		CnsmrSeqNo *string `json:"CnsmrSeqNo,omitempty" name:"CnsmrSeqNo"`

		// STRING (10)，本次交易返回查询结果记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
		ResultNum *string `json:"ResultNum,omitempty" name:"ResultNum"`

		// STRING(30)，起始记录号
	// 注意：此字段可能返回 null，表示取不到有效值。
		StartRecordNo *string `json:"StartRecordNo,omitempty" name:"StartRecordNo"`

		// STRING(2)，结束标志（0: 否; 1: 是）
	// 注意：此字段可能返回 null，表示取不到有效值。
		EndFlag *string `json:"EndFlag,omitempty" name:"EndFlag"`

		// STRING (10)，符合业务查询条件的记录总数（重复次数, 一次最多返回20条记录）
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalNum *string `json:"TotalNum,omitempty" name:"TotalNum"`

		// 交易信息数组
	// 注意：此字段可能返回 null，表示取不到有效值。
		TranItemArray []*ClearItem `json:"TranItemArray,omitempty" name:"TranItemArray"`

		// STRING(1027)，保留域
	// 注意：此字段可能返回 null，表示取不到有效值。
		ReservedMsg *string `json:"ReservedMsg,omitempty" name:"ReservedMsg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryBankClearResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryBankClearResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryBankTransactionDetailsRequest struct {
	*tchttp.BaseRequest

	// String(22)，商户号（签约客户号）
	MrchCode *string `json:"MrchCode,omitempty" name:"MrchCode"`

	// STRING(2)，功能标志（1: 当日; 2: 历史）
	FunctionFlag *string `json:"FunctionFlag,omitempty" name:"FunctionFlag"`

	// STRING(50)，见证子帐户的帐号
	SubAcctNo *string `json:"SubAcctNo,omitempty" name:"SubAcctNo"`

	// STRING(4)，查询标志（1: 全部; 2: 转出; 3: 转入 ）
	QueryFlag *string `json:"QueryFlag,omitempty" name:"QueryFlag"`

	// STRING(10)，页码（起始值为1，每次最多返回20条记录，第二页返回的记录数为第21至40条记录，第三页为41至60条记录，顺序均按照建立时间的先后）
	PageNum *string `json:"PageNum,omitempty" name:"PageNum"`

	// STRING(8)，开始日期（若是历史查询，则必输，当日查询时，不起作用。格式：20190101）
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// STRING(8)，终止日期（若是历史查询，则必输，当日查询时，不起作用。格式：20190101）
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// STRING(1027)，保留域
	ReservedMsg *string `json:"ReservedMsg,omitempty" name:"ReservedMsg"`

	// STRING(12)，接入环境，默认接入沙箱环境。接入正式环境填"prod"
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *QueryBankTransactionDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryBankTransactionDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MrchCode")
	delete(f, "FunctionFlag")
	delete(f, "SubAcctNo")
	delete(f, "QueryFlag")
	delete(f, "PageNum")
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "ReservedMsg")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryBankTransactionDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryBankTransactionDetailsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// String(20)，返回码
		TxnReturnCode *string `json:"TxnReturnCode,omitempty" name:"TxnReturnCode"`

		// String(100)，返回信息
		TxnReturnMsg *string `json:"TxnReturnMsg,omitempty" name:"TxnReturnMsg"`

		// String(22)，交易流水号
		CnsmrSeqNo *string `json:"CnsmrSeqNo,omitempty" name:"CnsmrSeqNo"`

		// STRING(10)，本次交易返回查询结果记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
		ResultNum *string `json:"ResultNum,omitempty" name:"ResultNum"`

		// STRING(30)，起始记录号
	// 注意：此字段可能返回 null，表示取不到有效值。
		StartRecordNo *string `json:"StartRecordNo,omitempty" name:"StartRecordNo"`

		// STRING(2)，结束标志（0: 否; 1: 是）
	// 注意：此字段可能返回 null，表示取不到有效值。
		EndFlag *string `json:"EndFlag,omitempty" name:"EndFlag"`

		// STRING(10)，符合业务查询条件的记录总数（重复次数，一次最多返回20条记录）
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalNum *string `json:"TotalNum,omitempty" name:"TotalNum"`

		// 交易信息数组
	// 注意：此字段可能返回 null，表示取不到有效值。
		TranItemArray []*TransactionItem `json:"TranItemArray,omitempty" name:"TranItemArray"`

		// STRING(1027)，保留域
	// 注意：此字段可能返回 null，表示取不到有效值。
		ReservedMsg *string `json:"ReservedMsg,omitempty" name:"ReservedMsg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryBankTransactionDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryBankTransactionDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryBankWithdrawCashDetailsRequest struct {
	*tchttp.BaseRequest

	// String(22)，商户号（签约客户号）
	MrchCode *string `json:"MrchCode,omitempty" name:"MrchCode"`

	// STRING(2)，功能标志（1: 当日; 2: 历史）
	FunctionFlag *string `json:"FunctionFlag,omitempty" name:"FunctionFlag"`

	// STRING(50)，见证子帐户的帐号
	SubAcctNo *string `json:"SubAcctNo,omitempty" name:"SubAcctNo"`

	// STRING(4)，查询标志（2: 提现; 3: 清分）
	QueryFlag *string `json:"QueryFlag,omitempty" name:"QueryFlag"`

	// STRING(10)，页码（起始值为1，每次最多返回20条记录，第二页返回的记录数为第21至40条记录，第三页为41至60条记录，顺序均按照建立时间的先后）
	PageNum *string `json:"PageNum,omitempty" name:"PageNum"`

	// STRING(8)，开始日期（若是历史查询，则必输，当日查询时，不起作用。格式：20190101）
	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`

	// STRING(8)，结束日期（若是历史查询，则必输，当日查询时，不起作用。格式：20190101）
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// STRING(1027)，保留域
	ReservedMsg *string `json:"ReservedMsg,omitempty" name:"ReservedMsg"`

	// STRING(12)，接入环境，默认接入沙箱环境。接入正式环境填"prod"
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *QueryBankWithdrawCashDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryBankWithdrawCashDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MrchCode")
	delete(f, "FunctionFlag")
	delete(f, "SubAcctNo")
	delete(f, "QueryFlag")
	delete(f, "PageNum")
	delete(f, "BeginDate")
	delete(f, "EndDate")
	delete(f, "ReservedMsg")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryBankWithdrawCashDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryBankWithdrawCashDetailsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// String(20)，返回码
		TxnReturnCode *string `json:"TxnReturnCode,omitempty" name:"TxnReturnCode"`

		// String(100)，返回信息
		TxnReturnMsg *string `json:"TxnReturnMsg,omitempty" name:"TxnReturnMsg"`

		// String(22)，交易流水号
		CnsmrSeqNo *string `json:"CnsmrSeqNo,omitempty" name:"CnsmrSeqNo"`

		// STRING(10)，本次交易返回查询结果记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
		ResultNum *string `json:"ResultNum,omitempty" name:"ResultNum"`

		// STRING(30)，起始记录号
	// 注意：此字段可能返回 null，表示取不到有效值。
		StartRecordNo *string `json:"StartRecordNo,omitempty" name:"StartRecordNo"`

		// STRING(2)，结束标志（0:否; 1:是）
	// 注意：此字段可能返回 null，表示取不到有效值。
		EndFlag *string `json:"EndFlag,omitempty" name:"EndFlag"`

		// STRING(10)，符合业务查询条件的记录总数（重复次数，一次最多返回20条记录）
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalNum *string `json:"TotalNum,omitempty" name:"TotalNum"`

		// 交易信息数组
	// 注意：此字段可能返回 null，表示取不到有效值。
		TranItemArray []*WithdrawItem `json:"TranItemArray,omitempty" name:"TranItemArray"`

		// STRING(1027)，保留域
	// 注意：此字段可能返回 null，表示取不到有效值。
		ReservedMsg *string `json:"ReservedMsg,omitempty" name:"ReservedMsg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryBankWithdrawCashDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryBankWithdrawCashDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryBillDownloadURLData struct {

	// 统一对账单下载链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillDownloadURL *string `json:"BillDownloadURL,omitempty" name:"BillDownloadURL"`

	// 渠道原始对账单下载链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginalBillDownloadURL *string `json:"OriginalBillDownloadURL,omitempty" name:"OriginalBillDownloadURL"`
}

type QueryBillDownloadURLRequest struct {
	*tchttp.BaseRequest

	// 商户号
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 代发类型：
	// 1、 微信企业付款 
	// 2、 支付宝转账 
	// 3、 平安银企直联代发转账
	TransferType *int64 `json:"TransferType,omitempty" name:"TransferType"`

	// 账单日期，格式yyyy-MM-dd
	BillDate *string `json:"BillDate,omitempty" name:"BillDate"`
}

func (r *QueryBillDownloadURLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryBillDownloadURLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MerchantId")
	delete(f, "TransferType")
	delete(f, "BillDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryBillDownloadURLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryBillDownloadURLResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 错误码。响应成功："SUCCESS"，其他为不成功
		ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

		// 响应消息
		ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

		// 返回结果
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *QueryBillDownloadURLData `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryBillDownloadURLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryBillDownloadURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCommonTransferRechargeRequest struct {
	*tchttp.BaseRequest

	// String(22)，商户号（签约客户号）
	MrchCode *string `json:"MrchCode,omitempty" name:"MrchCode"`

	// STRING(2)，功能标志（1为查询当日数据，0查询历史数据）
	FunctionFlag *string `json:"FunctionFlag,omitempty" name:"FunctionFlag"`

	// STRING(8)，开始日期（格式：20190101）
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// STRING(8)，终止日期（格式：20190101）
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// STRING(10)，页码（起始值为1，每次最多返回20条记录，第二页返回的记录数为第21至40条记录，第三页为41至60条记录，顺序均按照建立时间的先后）
	PageNum *string `json:"PageNum,omitempty" name:"PageNum"`

	// STRING(1027)，保留域
	ReservedMsg *string `json:"ReservedMsg,omitempty" name:"ReservedMsg"`

	// STRING(12)，接入环境，默认接入沙箱环境。接入正式环境填"prod"
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *QueryCommonTransferRechargeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCommonTransferRechargeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MrchCode")
	delete(f, "FunctionFlag")
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "PageNum")
	delete(f, "ReservedMsg")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryCommonTransferRechargeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryCommonTransferRechargeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// String(20)，返回码
		TxnReturnCode *string `json:"TxnReturnCode,omitempty" name:"TxnReturnCode"`

		// String(100)，返回信息
		TxnReturnMsg *string `json:"TxnReturnMsg,omitempty" name:"TxnReturnMsg"`

		// String(22)，交易流水号
		CnsmrSeqNo *string `json:"CnsmrSeqNo,omitempty" name:"CnsmrSeqNo"`

		// STRING(10)，本次交易返回查询结果记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
		ResultNum *string `json:"ResultNum,omitempty" name:"ResultNum"`

		// STRING(30)，起始记录号
	// 注意：此字段可能返回 null，表示取不到有效值。
		StartRecordNo *string `json:"StartRecordNo,omitempty" name:"StartRecordNo"`

		// STRING(2)，结束标志（0: 否; 1: 是）
	// 注意：此字段可能返回 null，表示取不到有效值。
		EndFlag *string `json:"EndFlag,omitempty" name:"EndFlag"`

		// STRING(10)，符合业务查询条件的记录总数（重复次数，一次最多返回20条记录）
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalNum *string `json:"TotalNum,omitempty" name:"TotalNum"`

		// 交易信息数组
	// 注意：此字段可能返回 null，表示取不到有效值。
		TranItemArray []*TransferItem `json:"TranItemArray,omitempty" name:"TranItemArray"`

		// STRING(1027)，保留域
	// 注意：此字段可能返回 null，表示取不到有效值。
		ReservedMsg *string `json:"ReservedMsg,omitempty" name:"ReservedMsg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryCommonTransferRechargeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCommonTransferRechargeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryContractRequest struct {
	*tchttp.BaseRequest

	// 聚鑫分配的支付主MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 用户ID，长度不小于5位，仅支持字母和数字的组合
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 指定渠道：  wechat：微信支付  qqwallet：QQ钱包 
	//  bank：网银支付  只有一个渠道时需要指定
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// 枚举值：
	// CONTRACT_QUERY_MODE_BY_OUT_CONTRACT_CODE：按 OutContractCode + ContractSceneId 查询
	// CONTRACT_QUERY_MODE_BY_CHANNEL_CONTRACT_CODE：按ChannelContractCode查询
	ContractQueryMode *string `json:"ContractQueryMode,omitempty" name:"ContractQueryMode"`

	// 按照聚鑫安全密钥计算的签名
	MidasSignature *string `json:"MidasSignature,omitempty" name:"MidasSignature"`

	// 聚鑫分配的安全ID
	MidasSecretId *string `json:"MidasSecretId,omitempty" name:"MidasSecretId"`

	// 聚鑫计费SubAppId，代表子商户
	SubAppId *string `json:"SubAppId,omitempty" name:"SubAppId"`

	// 业务签约合同协议号 当 ContractQueryMode=CONTRACT_QUERY_MODE_BY_OUT_CONTRACT_CODE 时 ，必填
	OutContractCode *string `json:"OutContractCode,omitempty" name:"OutContractCode"`

	// 签约场景ID，当 ContractQueryMode=CONTRACT_QUERY_MODE_BY_OUT_CONTRACT_CODE 时 必填，在米大师侧托管后生成
	ContractSceneId *string `json:"ContractSceneId,omitempty" name:"ContractSceneId"`

	// 米大师生成的协议号 ，当 ContractQueryMode=CONTRACT_QUERY_MODE_BY_CHANNEL_CONTRACT_CODE 时必填
	ChannelContractCode *string `json:"ChannelContractCode,omitempty" name:"ChannelContractCode"`

	// 第三方渠道合约数据，为json字符串，与特定渠道有关
	ExternalContractData *string `json:"ExternalContractData,omitempty" name:"ExternalContractData"`

	// 环境名:
	// release: 现网环境
	// sandbox: 沙箱环境
	// development: 开发环境
	// 缺省: release
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`

	// USER_ID: 用户ID
	// ANONYMOUS: 匿名类型 USER_ID
	// 默认值为 USER_ID
	UserType *string `json:"UserType,omitempty" name:"UserType"`

	// 签约代扣穿透查询存量数据迁移模式
	MigrateMode *string `json:"MigrateMode,omitempty" name:"MigrateMode"`

	// 签约方式
	ContractMethod *string `json:"ContractMethod,omitempty" name:"ContractMethod"`
}

func (r *QueryContractRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryContractRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MidasAppId")
	delete(f, "UserId")
	delete(f, "Channel")
	delete(f, "ContractQueryMode")
	delete(f, "MidasSignature")
	delete(f, "MidasSecretId")
	delete(f, "SubAppId")
	delete(f, "OutContractCode")
	delete(f, "ContractSceneId")
	delete(f, "ChannelContractCode")
	delete(f, "ExternalContractData")
	delete(f, "MidasEnvironment")
	delete(f, "UserType")
	delete(f, "MigrateMode")
	delete(f, "ContractMethod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryContractRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryContractResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 签约数据
		ContractData *ResponseQueryContract `json:"ContractData,omitempty" name:"ContractData"`

		// 请求处理信息
		Msg *string `json:"Msg,omitempty" name:"Msg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryContractResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryContractResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryCustAcctIdBalanceRequest struct {
	*tchttp.BaseRequest

	// String(22)，商户号（签约客户号）
	MrchCode *string `json:"MrchCode,omitempty" name:"MrchCode"`

	// STRING(4)，查询标志（2: 普通会员子账号; 3: 功能子账号）
	QueryFlag *string `json:"QueryFlag,omitempty" name:"QueryFlag"`

	// STRING(10)，页码（起始值为1，每次最多返回20条记录，第二页返回的记录数为第21至40条记录，第三页为41至60条记录，顺序均按照建立时间的先后）
	PageNum *string `json:"PageNum,omitempty" name:"PageNum"`

	// STRING(50)，见证子账户的账号（若SelectFlag为2时，子账号必输）
	SubAcctNo *string `json:"SubAcctNo,omitempty" name:"SubAcctNo"`

	// STRING(1027)，保留域
	ReservedMsg *string `json:"ReservedMsg,omitempty" name:"ReservedMsg"`

	// STRING(12)，接入环境，默认接入沙箱环境。接入正式环境填"prod"
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *QueryCustAcctIdBalanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCustAcctIdBalanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MrchCode")
	delete(f, "QueryFlag")
	delete(f, "PageNum")
	delete(f, "SubAcctNo")
	delete(f, "ReservedMsg")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryCustAcctIdBalanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryCustAcctIdBalanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// String(20)，返回码
		TxnReturnCode *string `json:"TxnReturnCode,omitempty" name:"TxnReturnCode"`

		// String(100)，返回信息
		TxnReturnMsg *string `json:"TxnReturnMsg,omitempty" name:"TxnReturnMsg"`

		// String(22)，交易流水号
		CnsmrSeqNo *string `json:"CnsmrSeqNo,omitempty" name:"CnsmrSeqNo"`

		// STRING(10)，本次交易返回查询结果记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
		ResultNum *string `json:"ResultNum,omitempty" name:"ResultNum"`

		// STRING(30)，起始记录号
	// 注意：此字段可能返回 null，表示取不到有效值。
		StartRecordNo *string `json:"StartRecordNo,omitempty" name:"StartRecordNo"`

		// STRING(2)，结束标志（0: 否; 1: 是）
	// 注意：此字段可能返回 null，表示取不到有效值。
		EndFlag *string `json:"EndFlag,omitempty" name:"EndFlag"`

		// STRING(10)，符合业务查询条件的记录总数（重复次数，一次最多返回20条记录）
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalNum *string `json:"TotalNum,omitempty" name:"TotalNum"`

		// 账户信息数组
	// 注意：此字段可能返回 null，表示取不到有效值。
		AcctArray []*Acct `json:"AcctArray,omitempty" name:"AcctArray"`

		// STRING(1027)，保留域
	// 注意：此字段可能返回 null，表示取不到有效值。
		ReservedMsg *string `json:"ReservedMsg,omitempty" name:"ReservedMsg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryCustAcctIdBalanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCustAcctIdBalanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryDeclareData struct {

	// 商户号
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 对接方汇出指令编号
	TransactionId *string `json:"TransactionId,omitempty" name:"TransactionId"`

	// 申报流水号
	DeclareId *string `json:"DeclareId,omitempty" name:"DeclareId"`

	// 原申报流水号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginalDeclareId *string `json:"OriginalDeclareId,omitempty" name:"OriginalDeclareId"`

	// 付款人ID
	PayerId *string `json:"PayerId,omitempty" name:"PayerId"`

	// 源币种
	SourceCurrency *string `json:"SourceCurrency,omitempty" name:"SourceCurrency"`

	// 源金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceAmount *string `json:"SourceAmount,omitempty" name:"SourceAmount"`

	// 目的币种
	TargetCurrency *string `json:"TargetCurrency,omitempty" name:"TargetCurrency"`

	// 目的金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetAmount *string `json:"TargetAmount,omitempty" name:"TargetAmount"`

	// 交易编码
	TradeCode *string `json:"TradeCode,omitempty" name:"TradeCode"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`
}

type QueryDeclareResult struct {

	// 成功申报材料查询数据
	Data *QueryDeclareData `json:"Data,omitempty" name:"Data"`

	// 错误码
	Code *string `json:"Code,omitempty" name:"Code"`
}

type QueryDownloadBillURLRequest struct {
	*tchttp.BaseRequest

	// 分配给商户的AppId。进件成功后返给商户方的AppId。
	MerchantAppId *string `json:"MerchantAppId,omitempty" name:"MerchantAppId"`

	// 渠道编号。固定值：ZSB2B
	ChannelCode *string `json:"ChannelCode,omitempty" name:"ChannelCode"`

	// 对账单日期，格式yyyyMMdd
	BillDate *string `json:"BillDate,omitempty" name:"BillDate"`
}

func (r *QueryDownloadBillURLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryDownloadBillURLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MerchantAppId")
	delete(f, "ChannelCode")
	delete(f, "BillDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryDownloadBillURLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryDownloadBillURLResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分配给商户的AppId。进件成功后返给商户方的AppId。
		MerchantAppId *string `json:"MerchantAppId,omitempty" name:"MerchantAppId"`

		// 对账单下载地址。
		DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryDownloadBillURLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryDownloadBillURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryExchangeRateRequest struct {
	*tchttp.BaseRequest

	// 源币种 (默认CNY)
	SourceCurrency *string `json:"SourceCurrency,omitempty" name:"SourceCurrency"`

	// 目的币种 (见常见问题-汇出币种)
	TargetCurrency *string `json:"TargetCurrency,omitempty" name:"TargetCurrency"`

	// 接入环境。沙箱环境填sandbox
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *QueryExchangeRateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryExchangeRateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SourceCurrency")
	delete(f, "TargetCurrency")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryExchangeRateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryExchangeRateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询汇率结果
		Result *QueryExchangerateResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryExchangeRateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryExchangeRateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryExchangerateData struct {

	// 汇率
	Rate *string `json:"Rate,omitempty" name:"Rate"`

	// 源币种
	SourceCurrency *string `json:"SourceCurrency,omitempty" name:"SourceCurrency"`

	// 目的币种
	TargetCurrency *string `json:"TargetCurrency,omitempty" name:"TargetCurrency"`

	// 汇率时间
	RateTime *string `json:"RateTime,omitempty" name:"RateTime"`

	// 基准币种
	BaseCurrency *string `json:"BaseCurrency,omitempty" name:"BaseCurrency"`
}

type QueryExchangerateResult struct {

	// 错误码
	Code *string `json:"Code,omitempty" name:"Code"`

	// 查询汇率数据数组
	Data []*QueryExchangerateData `json:"Data,omitempty" name:"Data"`
}

type QueryInvoiceRequest struct {
	*tchttp.BaseRequest

	// 开票平台ID
	InvoicePlatformId *int64 `json:"InvoicePlatformId,omitempty" name:"InvoicePlatformId"`

	// 订单号
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 业务开票号
	OrderSn *string `json:"OrderSn,omitempty" name:"OrderSn"`

	// 发票种类：
	// 0：蓝票
	// 1：红票【该字段默认为0， 如果需要查询红票信息，本字段必须传1，否则可能查询不到需要的发票信息】。
	IsRed *int64 `json:"IsRed,omitempty" name:"IsRed"`

	// 接入环境。沙箱环境填sandbox。
	Profile *string `json:"Profile,omitempty" name:"Profile"`

	// 开票渠道。0：线上渠道，1：线下渠道。不填默认为线上渠道
	InvoiceChannel *int64 `json:"InvoiceChannel,omitempty" name:"InvoiceChannel"`

	// 当渠道为线下渠道时，必填
	SellerTaxpayerNum *string `json:"SellerTaxpayerNum,omitempty" name:"SellerTaxpayerNum"`
}

func (r *QueryInvoiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryInvoiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InvoicePlatformId")
	delete(f, "OrderId")
	delete(f, "OrderSn")
	delete(f, "IsRed")
	delete(f, "Profile")
	delete(f, "InvoiceChannel")
	delete(f, "SellerTaxpayerNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryInvoiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryInvoiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 发票查询结果
		Result *QueryInvoiceResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryInvoiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryInvoiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryInvoiceResult struct {

	// 错误消息
	Message *string `json:"Message,omitempty" name:"Message"`

	// 错误码
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// 查询发票数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *QueryInvoiceResultData `json:"Data,omitempty" name:"Data"`

	// 订单数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Order *Order `json:"Order,omitempty" name:"Order"`
}

type QueryInvoiceResultData struct {

	// 订单号
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 业务开票号
	OrderSn *string `json:"OrderSn,omitempty" name:"OrderSn"`

	// 发票状态
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 开票描述
	Message *string `json:"Message,omitempty" name:"Message"`

	// 开票日期
	TicketDate *string `json:"TicketDate,omitempty" name:"TicketDate"`

	// 发票号码
	TicketSn *string `json:"TicketSn,omitempty" name:"TicketSn"`

	// 发票代码
	TicketCode *string `json:"TicketCode,omitempty" name:"TicketCode"`

	// 检验码
	CheckCode *string `json:"CheckCode,omitempty" name:"CheckCode"`

	// 含税金额(元)
	AmountWithTax *string `json:"AmountWithTax,omitempty" name:"AmountWithTax"`

	// 不含税金额(元)
	AmountWithoutTax *string `json:"AmountWithoutTax,omitempty" name:"AmountWithoutTax"`

	// 税额(元)
	TaxAmount *string `json:"TaxAmount,omitempty" name:"TaxAmount"`

	// 是否被红冲
	IsRedWashed *int64 `json:"IsRedWashed,omitempty" name:"IsRedWashed"`

	// pdf地址
	PdfUrl *string `json:"PdfUrl,omitempty" name:"PdfUrl"`
}

type QueryInvoiceV2Request struct {
	*tchttp.BaseRequest

	// 开票平台ID
	InvoicePlatformId *int64 `json:"InvoicePlatformId,omitempty" name:"InvoicePlatformId"`

	// 订单号
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 发票种类：
	// 0：蓝票
	// 1：红票【该字段默认为0， 如果需要查询红票信息，本字段必须传1，否则可能查询不到需要的发票信息】。
	IsRed *int64 `json:"IsRed,omitempty" name:"IsRed"`

	// 接入环境。沙箱环境填sandbox。
	Profile *string `json:"Profile,omitempty" name:"Profile"`

	// 开票渠道。0：线上渠道，1：线下渠道。不填默认为线上渠道
	InvoiceChannel *int64 `json:"InvoiceChannel,omitempty" name:"InvoiceChannel"`

	// 当渠道为线下渠道时，必填
	SellerTaxpayerNum *string `json:"SellerTaxpayerNum,omitempty" name:"SellerTaxpayerNum"`
}

func (r *QueryInvoiceV2Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryInvoiceV2Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InvoicePlatformId")
	delete(f, "OrderId")
	delete(f, "IsRed")
	delete(f, "Profile")
	delete(f, "InvoiceChannel")
	delete(f, "SellerTaxpayerNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryInvoiceV2Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryInvoiceV2Response struct {
	*tchttp.BaseResponse
	Response *struct {

		// 发票查询结果
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *QueryInvoiceResultData `json:"Result,omitempty" name:"Result"`

		// 错误码
		ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

		// 错误消息
		ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryInvoiceV2Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryInvoiceV2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryItem struct {

	// 子商户账户
	SubAcctNo *string `json:"SubAcctNo,omitempty" name:"SubAcctNo"`

	// 子账户属性 
	// 1：普通会员子账号 
	// 2：挂账子账号 
	// 3：手续费子账号 
	// 4：利息子账号
	// 5：平台担保子账号
	SubAcctProperty *string `json:"SubAcctProperty,omitempty" name:"SubAcctProperty"`

	// 业务平台的子商户Id，唯一
	SubMchId *string `json:"SubMchId,omitempty" name:"SubMchId"`

	// 子账户名称
	SubAcctName *string `json:"SubAcctName,omitempty" name:"SubAcctName"`

	// 账户可用余额
	AcctAvailBal *string `json:"AcctAvailBal,omitempty" name:"AcctAvailBal"`

	// 可提现金额
	CashAmt *string `json:"CashAmt,omitempty" name:"CashAmt"`

	// 维护日期 开户日期或修改日期
	MaintenanceDate *string `json:"MaintenanceDate,omitempty" name:"MaintenanceDate"`
}

type QueryMemberBindRequest struct {
	*tchttp.BaseRequest

	// String(22)，商户号（签约客户号）
	MrchCode *string `json:"MrchCode,omitempty" name:"MrchCode"`

	// STRING(4)，查询标志（1: 全部会员; 2: 单个会员; 3: 单个会员的证件信息）
	QueryFlag *string `json:"QueryFlag,omitempty" name:"QueryFlag"`

	// STRING (10)，页码（起始值为1，每次最多返回20条记录，第二页返回的记录数为第21至40条记录，第三页为41至60条记录，顺序均按照建立时间的先后）
	PageNum *string `json:"PageNum,omitempty" name:"PageNum"`

	// STRING(50)，见证子账户的账号（若SelectFlag为2或3时，子账户账号必输）
	SubAcctNo *string `json:"SubAcctNo,omitempty" name:"SubAcctNo"`

	// STRING(1027)，保留域
	ReservedMsg *string `json:"ReservedMsg,omitempty" name:"ReservedMsg"`

	// STRING(12)，接入环境，默认接入沙箱环境。接入正式环境填"prod"
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *QueryMemberBindRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryMemberBindRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MrchCode")
	delete(f, "QueryFlag")
	delete(f, "PageNum")
	delete(f, "SubAcctNo")
	delete(f, "ReservedMsg")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryMemberBindRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryMemberBindResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// STRING (10)，本次交易返回查询结果记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
		ResultNum *string `json:"ResultNum,omitempty" name:"ResultNum"`

		// STRING(30)，起始记录号
	// 注意：此字段可能返回 null，表示取不到有效值。
		StartRecordNo *string `json:"StartRecordNo,omitempty" name:"StartRecordNo"`

		// STRING(2)，结束标志（0: 否; 1: 是）
	// 注意：此字段可能返回 null，表示取不到有效值。
		EndFlag *string `json:"EndFlag,omitempty" name:"EndFlag"`

		// STRING (10)，符合业务查询条件的记录总数（重复次数，一次最多返回20条记录）
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalNum *string `json:"TotalNum,omitempty" name:"TotalNum"`

		// 交易信息数组
	// 注意：此字段可能返回 null，表示取不到有效值。
		TranItemArray []*TranItem `json:"TranItemArray,omitempty" name:"TranItemArray"`

		// STRING(1027)，保留域
	// 注意：此字段可能返回 null，表示取不到有效值。
		ReservedMsg *string `json:"ReservedMsg,omitempty" name:"ReservedMsg"`

		// String(20)，返回码
		TxnReturnCode *string `json:"TxnReturnCode,omitempty" name:"TxnReturnCode"`

		// String(100)，返回信息
		TxnReturnMsg *string `json:"TxnReturnMsg,omitempty" name:"TxnReturnMsg"`

		// String(22)，交易流水号
		CnsmrSeqNo *string `json:"CnsmrSeqNo,omitempty" name:"CnsmrSeqNo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryMemberBindResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryMemberBindResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryMemberTransactionRequest struct {
	*tchttp.BaseRequest

	// String(22)，商户号（签约客户号）
	MrchCode *string `json:"MrchCode,omitempty" name:"MrchCode"`

	// STRING(2)，功能标志（1: 下单预支付; 2: 确认并付款; 3: 退款; 6: 直接支付T+1; 9: 直接支付T+0）
	FunctionFlag *string `json:"FunctionFlag,omitempty" name:"FunctionFlag"`

	// STRING(50)，转出方的见证子账户的账号（付款方）
	OutSubAcctNo *string `json:"OutSubAcctNo,omitempty" name:"OutSubAcctNo"`

	// STRING(32)，转出方的交易网会员代码
	OutMemberCode *string `json:"OutMemberCode,omitempty" name:"OutMemberCode"`

	// STRING(150)，转出方的见证子账户的户名（户名是绑卡时上送的账户名称，如果未绑卡，就送OpenCustAcctId接口上送的用户昵称UserNickname）
	OutSubAcctName *string `json:"OutSubAcctName,omitempty" name:"OutSubAcctName"`

	// STRING(50)，转入方的见证子账户的账号（收款方）
	InSubAcctNo *string `json:"InSubAcctNo,omitempty" name:"InSubAcctNo"`

	// STRING(32)，转入方的交易网会员代码
	InMemberCode *string `json:"InMemberCode,omitempty" name:"InMemberCode"`

	// STRING(150)，转入方的见证子账户的户名（户名是绑卡时上送的账户名称，如果未绑卡，就送OpenCustAcctId接口上送的用户昵称UserNickname）
	InSubAcctName *string `json:"InSubAcctName,omitempty" name:"InSubAcctName"`

	// STRING(20)，交易金额
	TranAmt *string `json:"TranAmt,omitempty" name:"TranAmt"`

	// STRING(20)，交易费用（平台收取交易费用）
	TranFee *string `json:"TranFee,omitempty" name:"TranFee"`

	// STRING(20)，交易类型（01: 普通交易）
	TranType *string `json:"TranType,omitempty" name:"TranType"`

	// STRING(3)，币种（默认: RMB）
	Ccy *string `json:"Ccy,omitempty" name:"Ccy"`

	// STRING(50)，订单号（功能标志为1,2,3时必输）
	OrderNo *string `json:"OrderNo,omitempty" name:"OrderNo"`

	// STRING(500)，订单内容
	OrderContent *string `json:"OrderContent,omitempty" name:"OrderContent"`

	// STRING(300)，备注（建议可送订单号，可在对账文件的备注字段获取到）
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// STRING(1027)，保留域（若需短信验证码则此项必输短信指令号）
	ReservedMsg *string `json:"ReservedMsg,omitempty" name:"ReservedMsg"`

	// STRING(300)，网银签名（若需短信验证码则此项必输）
	WebSign *string `json:"WebSign,omitempty" name:"WebSign"`

	// STRING(12)，接入环境，默认接入沙箱环境。接入正式环境填"prod"
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *QueryMemberTransactionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryMemberTransactionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MrchCode")
	delete(f, "FunctionFlag")
	delete(f, "OutSubAcctNo")
	delete(f, "OutMemberCode")
	delete(f, "OutSubAcctName")
	delete(f, "InSubAcctNo")
	delete(f, "InMemberCode")
	delete(f, "InSubAcctName")
	delete(f, "TranAmt")
	delete(f, "TranFee")
	delete(f, "TranType")
	delete(f, "Ccy")
	delete(f, "OrderNo")
	delete(f, "OrderContent")
	delete(f, "Remark")
	delete(f, "ReservedMsg")
	delete(f, "WebSign")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryMemberTransactionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryMemberTransactionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// String(20)，返回码
		TxnReturnCode *string `json:"TxnReturnCode,omitempty" name:"TxnReturnCode"`

		// String(100)，返回信息
		TxnReturnMsg *string `json:"TxnReturnMsg,omitempty" name:"TxnReturnMsg"`

		// String(22)，交易流水号
		CnsmrSeqNo *string `json:"CnsmrSeqNo,omitempty" name:"CnsmrSeqNo"`

		// STRING(52)，见证系统流水号（即电商见证宝系统生成的流水号，可关联具体一笔请求）
	// 注意：此字段可能返回 null，表示取不到有效值。
		FrontSeqNo *string `json:"FrontSeqNo,omitempty" name:"FrontSeqNo"`

		// STRING(1027)，保留域
	// 注意：此字段可能返回 null，表示取不到有效值。
		ReservedMsg *string `json:"ReservedMsg,omitempty" name:"ReservedMsg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryMemberTransactionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryMemberTransactionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryMerchantBalanceData struct {

	// 余额币种
	Currency *string `json:"Currency,omitempty" name:"Currency"`

	// 账户余额
	Balance *string `json:"Balance,omitempty" name:"Balance"`

	// 商户ID
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`
}

type QueryMerchantBalanceRequest struct {
	*tchttp.BaseRequest

	// 余额币种
	Currency *string `json:"Currency,omitempty" name:"Currency"`

	// 接入环境。沙箱环境填sandbox
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *QueryMerchantBalanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryMerchantBalanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Currency")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryMerchantBalanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryMerchantBalanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 对接方账户余额查询结果
		Result *QueryMerchantBalanceResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryMerchantBalanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryMerchantBalanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryMerchantBalanceResult struct {

	// 错误码
	Code *string `json:"Code,omitempty" name:"Code"`

	// 对接账户余额查询数据
	Data *QueryMerchantBalanceData `json:"Data,omitempty" name:"Data"`
}

type QueryMerchantInfoForManagementRequest struct {
	*tchttp.BaseRequest

	// 开票平台ID
	InvoicePlatformId *int64 `json:"InvoicePlatformId,omitempty" name:"InvoicePlatformId"`

	// 页码
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 页大小
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 接入环境。沙箱环境填sandbox。
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *QueryMerchantInfoForManagementRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryMerchantInfoForManagementRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InvoicePlatformId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryMerchantInfoForManagementRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryMerchantInfoForManagementResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 商户结果
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *MerchantManagementResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryMerchantInfoForManagementResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryMerchantInfoForManagementResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryMerchantOrderRequest struct {
	*tchttp.BaseRequest

	// 进件成功后返给商户方的AppId。
	MerchantAppId *string `json:"MerchantAppId,omitempty" name:"MerchantAppId"`

	// 平台流水号。平台唯一订单号。
	OrderNo *string `json:"OrderNo,omitempty" name:"OrderNo"`
}

func (r *QueryMerchantOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryMerchantOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MerchantAppId")
	delete(f, "OrderNo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryMerchantOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryMerchantOrderResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 进件成功后返给商户方的AppId。
		MerchantAppId *string `json:"MerchantAppId,omitempty" name:"MerchantAppId"`

		// 平台流水号。平台唯一订单号。
		OrderNo *string `json:"OrderNo,omitempty" name:"OrderNo"`

		// 订单支付状态。0-下单失败 1-下单成功未支付 2-支付成功 3-支付失败 4-退款中 5-退款成功 6-退款失败 7-待付款 8-待确认。
		Status *string `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryMerchantOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryMerchantOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryMerchantRequest struct {
	*tchttp.BaseRequest

	// 进件成功后返给商户方的 AppId
	MerchantAppId *string `json:"MerchantAppId,omitempty" name:"MerchantAppId"`
}

func (r *QueryMerchantRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryMerchantRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MerchantAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryMerchantRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryMerchantResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分配给商户的 AppId，该 AppId 为后续各项 交易的商户标识。
		MerchantAppId *string `json:"MerchantAppId,omitempty" name:"MerchantAppId"`

		// 收款商户名称。
		MerchantName *string `json:"MerchantName,omitempty" name:"MerchantName"`

		// B2B 支付标志。是否开通 B2B 支付， 1:开通 0:不开通。
		BusinessPayFlag *string `json:"BusinessPayFlag,omitempty" name:"BusinessPayFlag"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryMerchantResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryMerchantResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryOrderOutOrderList struct {

	// 聚鑫分配的支付主MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 支付金额，单位：分
	Amt *int64 `json:"Amt,omitempty" name:"Amt"`

	// 用户Id
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 现金支付金额
	CashAmt *string `json:"CashAmt,omitempty" name:"CashAmt"`

	// 发货标识，由业务在调用聚鑫下单 接口的时候下发
	Metadata *string `json:"Metadata,omitempty" name:"Metadata"`

	// 支付时间unix时间戳
	PayTime *string `json:"PayTime,omitempty" name:"PayTime"`

	// 抵扣券金额
	CouponAmt *string `json:"CouponAmt,omitempty" name:"CouponAmt"`

	// 下单时间unix时间戳
	OrderTime *string `json:"OrderTime,omitempty" name:"OrderTime"`

	// 物品id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 高速场景信息
	SceneInfo *string `json:"SceneInfo,omitempty" name:"SceneInfo"`

	// 当前订单的订单状态 
	// 0：初始状态，获取聚鑫交易订单成功；  
	// 1：拉起聚鑫支付页面成功，用户未 支付；
	// 2：用户支付成功，正在发货；
	// 3：用户支付成功，发货失败；
	// 4：用户支付成功，发货成功；
	// 5：聚鑫支付页面正在失效中；
	// 6：聚鑫支付页面已经失效；
	OrderState *string `json:"OrderState,omitempty" name:"OrderState"`

	// 支付渠道：wechat：微信支付;
	// qqwallet：QQ钱包;
	// bank：网银
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// 是否曾退款
	RefundFlag *string `json:"RefundFlag,omitempty" name:"RefundFlag"`

	// 务支付订单号
	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`

	// 商品名称
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 支付回调时间，unix时间戳
	CallBackTime *string `json:"CallBackTime,omitempty" name:"CallBackTime"`

	// ISO 货币代码，CNY
	CurrencyType *string `json:"CurrencyType,omitempty" name:"CurrencyType"`

	// 微校场景账户Id
	AcctSubAppId *string `json:"AcctSubAppId,omitempty" name:"AcctSubAppId"`

	// 调用下单接口获取的聚鑫交易订单
	TransactionId *string `json:"TransactionId,omitempty" name:"TransactionId"`

	// 聚鑫内部渠道订单号
	ChannelOrderId *string `json:"ChannelOrderId,omitempty" name:"ChannelOrderId"`

	// 调用下单接口传进来的 SubOutTradeNoList
	SubOrderList []*QueryOrderOutSubOrderList `json:"SubOrderList,omitempty" name:"SubOrderList"`

	// 支付机构订单号
	ChannelExternalOrderId *string `json:"ChannelExternalOrderId,omitempty" name:"ChannelExternalOrderId"`
}

type QueryOrderOutSubOrderList struct {

	// 子订单支付金额
	Amt *int64 `json:"Amt,omitempty" name:"Amt"`

	// 子订单结算应收金额，单位：分
	SubMchIncome *int64 `json:"SubMchIncome,omitempty" name:"SubMchIncome"`

	// 发货标识，由业务在调用Midas下单接口的时候下发。
	Metadata *string `json:"Metadata,omitempty" name:"Metadata"`

	// 子订单原始金额
	OriginalAmt *int64 `json:"OriginalAmt,omitempty" name:"OriginalAmt"`

	// 子订单平台应收金额，单位：分
	PlatformIncome *int64 `json:"PlatformIncome,omitempty" name:"PlatformIncome"`

	// 子订单商品详情
	ProductDetail *string `json:"ProductDetail,omitempty" name:"ProductDetail"`

	// 子订单商品名称
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 核销状态，1表示核销，0表示未核销
	SettleCheck *int64 `json:"SettleCheck,omitempty" name:"SettleCheck"`

	// 聚鑫计费SubAppId，代表子商户
	SubAppId *string `json:"SubAppId,omitempty" name:"SubAppId"`

	// 子订单号
	SubOutTradeNo *string `json:"SubOutTradeNo,omitempty" name:"SubOutTradeNo"`
}

type QueryOrderRequest struct {
	*tchttp.BaseRequest

	// 聚鑫分配的支付主 MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 用户ID，长度不小于5位， 仅支持字母和数字的组合
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// type=by_order根据订单号 查订单；
	// type=by_user根据用户id 查订单 。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 聚鑫分配的安全ID
	MidasSecretId *string `json:"MidasSecretId,omitempty" name:"MidasSecretId"`

	// 按照聚鑫安全密钥计算的签名
	MidasSignature *string `json:"MidasSignature,omitempty" name:"MidasSignature"`

	// 每页返回的记录数。根据用户 号码查询订单列表时需要传。 用于分页展示。Type=by_order时必填
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 记录数偏移量，默认从0开 始。根据用户号码查询订单列 表时需要传。用于分页展示。Type=by_order时必填
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询开始时间，Unix时间戳。Type=by_order时必填
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，Unix时间戳。Type=by_order时必填
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 业务订单号，OutTradeNo与 TransactionId不能同时为 空，都传优先使用 OutTradeNo
	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`

	// 聚鑫订单号，OutTradeNo与 TransactionId不能同时为 空，都传优先使用 OutTradeNo
	TransactionId *string `json:"TransactionId,omitempty" name:"TransactionId"`

	// 环境名:
	// release: 现网环境
	// sandbox: 沙箱环境
	// development: 开发环境
	// 缺省: release
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`
}

func (r *QueryOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MidasAppId")
	delete(f, "UserId")
	delete(f, "Type")
	delete(f, "MidasSecretId")
	delete(f, "MidasSignature")
	delete(f, "Count")
	delete(f, "Offset")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "OutTradeNo")
	delete(f, "TransactionId")
	delete(f, "MidasEnvironment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryOrderResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回订单数
		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 查询结果的订单列表
		OrderList []*QueryOrderOutOrderList `json:"OrderList,omitempty" name:"OrderList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryOutwardOrderData struct {

	// 商户号
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 对接方汇出指令编号
	TransactionId *string `json:"TransactionId,omitempty" name:"TransactionId"`

	// 财务日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	AcctDate *string `json:"AcctDate,omitempty" name:"AcctDate"`

	// 定价币种
	// 注意：此字段可能返回 null，表示取不到有效值。
	PricingCurrency *string `json:"PricingCurrency,omitempty" name:"PricingCurrency"`

	// 源币种
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceCurrency *string `json:"SourceCurrency,omitempty" name:"SourceCurrency"`

	// 源金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceAmount *string `json:"SourceAmount,omitempty" name:"SourceAmount"`

	// 目的币种
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetCurrency *string `json:"TargetCurrency,omitempty" name:"TargetCurrency"`

	// 目的金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetAmount *string `json:"TargetAmount,omitempty" name:"TargetAmount"`

	// 汇率
	// 注意：此字段可能返回 null，表示取不到有效值。
	FxRate *string `json:"FxRate,omitempty" name:"FxRate"`

	// 指令状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailReason *string `json:"FailReason,omitempty" name:"FailReason"`

	// 退汇金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	RefundAmount *string `json:"RefundAmount,omitempty" name:"RefundAmount"`

	// 退汇币种
	// 注意：此字段可能返回 null，表示取不到有效值。
	RefundCurrency *string `json:"RefundCurrency,omitempty" name:"RefundCurrency"`
}

type QueryOutwardOrderRequest struct {
	*tchttp.BaseRequest

	// 对接方汇出指令编号
	TransactionId *string `json:"TransactionId,omitempty" name:"TransactionId"`

	// 接入环境。沙箱环境填sandbox
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *QueryOutwardOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryOutwardOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TransactionId")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryOutwardOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryOutwardOrderResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询汇出结果
		Result *QueryOutwardOrderResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryOutwardOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryOutwardOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryOutwardOrderResult struct {

	// 错误码
	Code *string `json:"Code,omitempty" name:"Code"`

	// 查询汇出数据
	Data *QueryOutwardOrderData `json:"Data,omitempty" name:"Data"`
}

type QueryPayerInfoRequest struct {
	*tchttp.BaseRequest

	// 付款人ID
	PayerId *string `json:"PayerId,omitempty" name:"PayerId"`

	// 接入环境。沙箱环境填sandbox
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *QueryPayerInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryPayerInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PayerId")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryPayerInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryPayerInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 付款人查询结果
		Result *QueryPayerinfoResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryPayerInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryPayerInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryPayerinfoData struct {

	// 商户号
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 付款人ID
	PayerId *string `json:"PayerId,omitempty" name:"PayerId"`

	// 审核状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailReason *string `json:"FailReason,omitempty" name:"FailReason"`

	// 付款人类型
	PayerType *string `json:"PayerType,omitempty" name:"PayerType"`

	// 付款人姓名
	PayerName *string `json:"PayerName,omitempty" name:"PayerName"`

	// 付款人证件类型
	PayerIdType *string `json:"PayerIdType,omitempty" name:"PayerIdType"`

	// 付款人证件号
	PayerIdNo *string `json:"PayerIdNo,omitempty" name:"PayerIdNo"`

	// 付款人联系电话
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayerContactNumber *string `json:"PayerContactNumber,omitempty" name:"PayerContactNumber"`

	// 付款人联系邮箱
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayerEmailAddress *string `json:"PayerEmailAddress,omitempty" name:"PayerEmailAddress"`

	// 付款人常驻国家或地区编码
	PayerCountryCode *string `json:"PayerCountryCode,omitempty" name:"PayerCountryCode"`

	// 付款人联系名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayerContactName *string `json:"PayerContactName,omitempty" name:"PayerContactName"`
}

type QueryPayerinfoResult struct {

	// 错误码
	Code *string `json:"Code,omitempty" name:"Code"`

	// 付款人查询数据
	Data *QueryPayerinfoData `json:"Data,omitempty" name:"Data"`
}

type QueryReconciliationDocumentRequest struct {
	*tchttp.BaseRequest

	// String(22)，商户号
	MrchCode *string `json:"MrchCode,omitempty" name:"MrchCode"`

	// STRING(10)，文件类型（充值文件-CZ; 提现文件-TX; 交易文件-JY; 余额文件-YE; 合约文件-HY）
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// STRING(8)，文件日期（格式：20190101）
	FileDate *string `json:"FileDate,omitempty" name:"FileDate"`

	// STRING(1027)，保留域
	ReservedMsg *string `json:"ReservedMsg,omitempty" name:"ReservedMsg"`

	// STRING(12)，接入环境，默认接入沙箱环境。接入正式环境填"prod"
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *QueryReconciliationDocumentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryReconciliationDocumentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MrchCode")
	delete(f, "FileType")
	delete(f, "FileDate")
	delete(f, "ReservedMsg")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryReconciliationDocumentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryReconciliationDocumentResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// String(20)，返回码
		TxnReturnCode *string `json:"TxnReturnCode,omitempty" name:"TxnReturnCode"`

		// String(100)，返回信息
		TxnReturnMsg *string `json:"TxnReturnMsg,omitempty" name:"TxnReturnMsg"`

		// String(22)，交易流水号
		CnsmrSeqNo *string `json:"CnsmrSeqNo,omitempty" name:"CnsmrSeqNo"`

		// STRING(10)，本次交易返回查询结果记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
		ResultNum *string `json:"ResultNum,omitempty" name:"ResultNum"`

		// 交易信息数组
	// 注意：此字段可能返回 null，表示取不到有效值。
		TranItemArray []*FileItem `json:"TranItemArray,omitempty" name:"TranItemArray"`

		// STRING(1027)，保留域
	// 注意：此字段可能返回 null，表示取不到有效值。
		ReservedMsg *string `json:"ReservedMsg,omitempty" name:"ReservedMsg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryReconciliationDocumentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryReconciliationDocumentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryRefundRequest struct {
	*tchttp.BaseRequest

	// 用户ID，长度不小于5位，仅支持字母和数字的组合。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 退款订单号，仅支持数字、字母、下划线（_）、横杠字符（-）、点（.）的组合。
	RefundId *string `json:"RefundId,omitempty" name:"RefundId"`

	// 聚鑫分配的支付主MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 聚鑫分配的安全ID
	MidasSecretId *string `json:"MidasSecretId,omitempty" name:"MidasSecretId"`

	// 按照聚鑫安全密钥计算的签名
	MidasSignature *string `json:"MidasSignature,omitempty" name:"MidasSignature"`

	// 环境名:
	// release: 现网环境
	// sandbox: 沙箱环境
	// development: 开发环境
	// 缺省: release
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`
}

func (r *QueryRefundRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryRefundRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "RefundId")
	delete(f, "MidasAppId")
	delete(f, "MidasSecretId")
	delete(f, "MidasSignature")
	delete(f, "MidasEnvironment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryRefundRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryRefundResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 退款状态码，退款提交成功后返回  1：退款中；  2：退款成功；  3：退款失败。
		State *string `json:"State,omitempty" name:"State"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryRefundResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryRefundResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QuerySinglePayItem struct {

	// 付款状态（S：支付成功；P：支付处理中；F：支付失败）
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayStatus *string `json:"PayStatus,omitempty" name:"PayStatus"`

	// 平台信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlatformMsg *string `json:"PlatformMsg,omitempty" name:"PlatformMsg"`

	// 银行原始返回码
	// 注意：此字段可能返回 null，表示取不到有效值。
	BankRetCode *string `json:"BankRetCode,omitempty" name:"BankRetCode"`

	// 银行原始返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	BankRetMsg *string `json:"BankRetMsg,omitempty" name:"BankRetMsg"`
}

type QuerySinglePayRequest struct {
	*tchttp.BaseRequest

	// 业务流水号
	SerialNumber *string `json:"SerialNumber,omitempty" name:"SerialNumber"`

	// 接入环境。沙箱环境填sandbox
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *QuerySinglePayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QuerySinglePayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SerialNumber")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QuerySinglePayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QuerySinglePayResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回结果
		Result *QuerySinglePayResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QuerySinglePayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QuerySinglePayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QuerySinglePayResult struct {

	// 受理状态（S：处理成功；F：处理失败）
	HandleStatus *string `json:"HandleStatus,omitempty" name:"HandleStatus"`

	// 受理状态描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	HandleMsg *string `json:"HandleMsg,omitempty" name:"HandleMsg"`

	// 业务流水号
	SerialNo *string `json:"SerialNo,omitempty" name:"SerialNo"`

	// 支付明细
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*QuerySinglePayItem `json:"Items,omitempty" name:"Items"`
}

type QuerySingleTransactionStatusRequest struct {
	*tchttp.BaseRequest

	// String(22)，商户号（签约客户号）
	MrchCode *string `json:"MrchCode,omitempty" name:"MrchCode"`

	// STRING(2)，功能标志（2: 会员间交易; 3: 提现; 4: 充值）
	FunctionFlag *string `json:"FunctionFlag,omitempty" name:"FunctionFlag"`

	// STRING(52)，交易网流水号（提现，充值或会员交易请求时的CnsmrSeqNo值）
	TranNetSeqNo *string `json:"TranNetSeqNo,omitempty" name:"TranNetSeqNo"`

	// STRING(50)，见证子帐户的帐号（未启用）
	SubAcctNo *string `json:"SubAcctNo,omitempty" name:"SubAcctNo"`

	// STRING(8)，交易日期（未启用）
	TranDate *string `json:"TranDate,omitempty" name:"TranDate"`

	// STRING(1027)，保留域
	ReservedMsg *string `json:"ReservedMsg,omitempty" name:"ReservedMsg"`

	// STRING(12)，接入环境，默认接入沙箱环境。接入正式环境填"prod"
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *QuerySingleTransactionStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QuerySingleTransactionStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MrchCode")
	delete(f, "FunctionFlag")
	delete(f, "TranNetSeqNo")
	delete(f, "SubAcctNo")
	delete(f, "TranDate")
	delete(f, "ReservedMsg")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QuerySingleTransactionStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QuerySingleTransactionStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// String(20)，返回码
		TxnReturnCode *string `json:"TxnReturnCode,omitempty" name:"TxnReturnCode"`

		// String(100)，返回信息
		TxnReturnMsg *string `json:"TxnReturnMsg,omitempty" name:"TxnReturnMsg"`

		// String(22)，交易流水号
		CnsmrSeqNo *string `json:"CnsmrSeqNo,omitempty" name:"CnsmrSeqNo"`

		// STRING(2)，记账标志（记账标志。1: 登记挂账; 2: 支付; 3: 提现; 4: 清分; 5: 下单预支付; 6: 确认并付款; 7: 退款; 8: 支付到平台; N: 其他）
	// 注意：此字段可能返回 null，表示取不到有效值。
		BookingFlag *string `json:"BookingFlag,omitempty" name:"BookingFlag"`

		// STRING(32)，交易状态（0: 成功; 1: 失败; 2: 待确认; 5: 待处理; 6: 处理中。0和1是终态，2、5、6是中间态，其中2是特指提现后待确认提现是否成功，5是银行收到交易等待处理，6是交易正在处理）
	// 注意：此字段可能返回 null，表示取不到有效值。
		TranStatus *string `json:"TranStatus,omitempty" name:"TranStatus"`

		// STRING(20)，交易金额
	// 注意：此字段可能返回 null，表示取不到有效值。
		TranAmt *string `json:"TranAmt,omitempty" name:"TranAmt"`

		// STRING(8)，交易日期
	// 注意：此字段可能返回 null，表示取不到有效值。
		TranDate *string `json:"TranDate,omitempty" name:"TranDate"`

		// STRING(20)，交易时间
	// 注意：此字段可能返回 null，表示取不到有效值。
		TranTime *string `json:"TranTime,omitempty" name:"TranTime"`

		// STRING(50)，转入子账户账号
	// 注意：此字段可能返回 null，表示取不到有效值。
		InSubAcctNo *string `json:"InSubAcctNo,omitempty" name:"InSubAcctNo"`

		// STRING(50)，转出子账户账号
	// 注意：此字段可能返回 null，表示取不到有效值。
		OutSubAcctNo *string `json:"OutSubAcctNo,omitempty" name:"OutSubAcctNo"`

		// STRING(300)，失败信息（当提现失败时，返回交易失败原因）
	// 注意：此字段可能返回 null，表示取不到有效值。
		FailMsg *string `json:"FailMsg,omitempty" name:"FailMsg"`

		// STRING(50)，原前置流水号
	// 注意：此字段可能返回 null，表示取不到有效值。
		OldTranFrontSeqNo *string `json:"OldTranFrontSeqNo,omitempty" name:"OldTranFrontSeqNo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QuerySingleTransactionStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QuerySingleTransactionStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QuerySmallAmountTransferRequest struct {
	*tchttp.BaseRequest

	// String(22)，商户号（签约客户号）
	MrchCode *string `json:"MrchCode,omitempty" name:"MrchCode"`

	// STRING(52)，原交易流水号（小额鉴权交易请求时的CnsmrSeqNo值）
	OldTranSeqNo *string `json:"OldTranSeqNo,omitempty" name:"OldTranSeqNo"`

	// STRING(8)，交易日期（格式：20190101）
	TranDate *string `json:"TranDate,omitempty" name:"TranDate"`

	// STRING(1027)，保留域
	ReservedMsg *string `json:"ReservedMsg,omitempty" name:"ReservedMsg"`

	// STRING(12)，接入环境，默认接入沙箱环境。接入正式环境填"prod"
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *QuerySmallAmountTransferRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QuerySmallAmountTransferRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MrchCode")
	delete(f, "OldTranSeqNo")
	delete(f, "TranDate")
	delete(f, "ReservedMsg")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QuerySmallAmountTransferRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QuerySmallAmountTransferResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// String(20)，返回码
		TxnReturnCode *string `json:"TxnReturnCode,omitempty" name:"TxnReturnCode"`

		// String(100)，返回信息
		TxnReturnMsg *string `json:"TxnReturnMsg,omitempty" name:"TxnReturnMsg"`

		// String(22)，交易流水号
		CnsmrSeqNo *string `json:"CnsmrSeqNo,omitempty" name:"CnsmrSeqNo"`

		// STRING(10)，返回状态（0: 成功; 1: 失败; 2: 待确认）
	// 注意：此字段可能返回 null，表示取不到有效值。
		ReturnStatus *string `json:"ReturnStatus,omitempty" name:"ReturnStatus"`

		// STRING(512)，返回信息（失败返回具体信息）
	// 注意：此字段可能返回 null，表示取不到有效值。
		ReturnMsg *string `json:"ReturnMsg,omitempty" name:"ReturnMsg"`

		// STRING(1027)，保留域
	// 注意：此字段可能返回 null，表示取不到有效值。
		ReservedMsg *string `json:"ReservedMsg,omitempty" name:"ReservedMsg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QuerySmallAmountTransferResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QuerySmallAmountTransferResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryTradeData struct {

	// 商户号
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 贸易材料流水号
	TradeFileId *string `json:"TradeFileId,omitempty" name:"TradeFileId"`

	// 贸易材料订单号
	TradeOrderId *string `json:"TradeOrderId,omitempty" name:"TradeOrderId"`

	// 审核状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailReason *string `json:"FailReason,omitempty" name:"FailReason"`

	// 付款人ID
	PayerId *string `json:"PayerId,omitempty" name:"PayerId"`

	// 收款人姓名
	PayeeName *string `json:"PayeeName,omitempty" name:"PayeeName"`

	// 收款人常驻国家或地区编码
	PayeeCountryCode *string `json:"PayeeCountryCode,omitempty" name:"PayeeCountryCode"`

	// 交易类型
	TradeType *string `json:"TradeType,omitempty" name:"TradeType"`

	// 交易日期
	TradeTime *string `json:"TradeTime,omitempty" name:"TradeTime"`

	// 交易币种
	TradeCurrency *string `json:"TradeCurrency,omitempty" name:"TradeCurrency"`

	// 交易金额
	TradeAmount *string `json:"TradeAmount,omitempty" name:"TradeAmount"`

	// 交易名称
	TradeName *string `json:"TradeName,omitempty" name:"TradeName"`

	// 交易数量
	TradeCount *int64 `json:"TradeCount,omitempty" name:"TradeCount"`

	// 货贸承运人
	// 注意：此字段可能返回 null，表示取不到有效值。
	GoodsCarrier *string `json:"GoodsCarrier,omitempty" name:"GoodsCarrier"`

	// 服贸交易细节
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceDetail *string `json:"ServiceDetail,omitempty" name:"ServiceDetail"`

	// 服贸服务时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceTime *string `json:"ServiceTime,omitempty" name:"ServiceTime"`
}

type QueryTradeRequest struct {
	*tchttp.BaseRequest

	// 贸易材料流水号
	TradeFileId *string `json:"TradeFileId,omitempty" name:"TradeFileId"`

	// 接入环境。沙箱环境填sandbox
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *QueryTradeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryTradeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TradeFileId")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryTradeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryTradeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 贸易材料明细查询结果
		Result *QueryTradeResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryTradeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryTradeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryTradeResult struct {

	// 贸易材料明细查询数据
	Data *QueryTradeData `json:"Data,omitempty" name:"Data"`

	// 错误码
	Code *string `json:"Code,omitempty" name:"Code"`
}

type QueryTransferBatchRequest struct {
	*tchttp.BaseRequest

	// 商户号。
	// 示例值：129284394
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 微信明细单号。
	// 微信区分明细单返回的唯一标识。
	// 示例值：1030000071100999991182020050700019480101
	NeedQueryDetail *bool `json:"NeedQueryDetail,omitempty" name:"NeedQueryDetail"`

	// 商家批次单号。
	// 商户系统内部的商家批次单号，此参数只能由数字、字母组成，商户系统内部唯一，UTF8编码，最多32个字符。
	// 示例值：plfk2020042013
	MerchantBatchNo *string `json:"MerchantBatchNo,omitempty" name:"MerchantBatchNo"`

	// 是否查询账单明细。
	// true-是；
	// false-否，默认否。
	// 商户可选择是否查询指定状态的转账明细单，当转账批次单状态为“FINISHED”（已完成）时，才会返回满足条件的转账明细单。
	// 示例值：true
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// 环境名:
	// release: 现网环境
	// sandbox: 沙箱环境
	// development: 开发环境
	// 缺省: release
	Profile *string `json:"Profile,omitempty" name:"Profile"`

	// 请求资源起始位置。
	// 从0开始，默认值为0。
	// 示例值：20
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 最大资源条数。
	// 该次请求可返回的最大资源（转账明细单）条数，最小20条，最大100条，不传则默认20条。不足20条按实际条数返回
	// 示例值：20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 明细状态。
	// ALL：全部，需要同时查询转账成功喝失败的明细单；
	// SUCCESS：转账成功，只查询成功的明细单；
	// FAIL：转账失败，只查询转账失败的明细单。
	// 示例值：FAIL
	DetailStatus *string `json:"DetailStatus,omitempty" name:"DetailStatus"`
}

func (r *QueryTransferBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryTransferBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MerchantId")
	delete(f, "NeedQueryDetail")
	delete(f, "MerchantBatchNo")
	delete(f, "BatchId")
	delete(f, "Profile")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "DetailStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryTransferBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryTransferBatchResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 商户号。
	// 示例值：19300009329
		MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

		// 商家批次单号。
	// 商户系统内部的商家批次单号，此参数只能由数字、字母组成，商户系统内部唯一，UTF8编码，最多32个字符。
	// 示例值：plfk2020042013
		MerchantBatchNo *string `json:"MerchantBatchNo,omitempty" name:"MerchantBatchNo"`

		// 微信批次单号。
	// 微信商家转账系统返回的唯一标识。
	// 示例值：1030000071100999991182020050700019480001
		BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

		// 直连商户appId。
	// 商户号绑定的appid。
	// 示例值：wxf636efh567hg4356
		MerchantAppId *string `json:"MerchantAppId,omitempty" name:"MerchantAppId"`

		// 批次状态。
	// ACCEPTED:已受理，批次已受理成功，若发起批量转账的30分钟后，转账批次单仍处于该状态，可能原因是商户账户余额不足等。商户可查询账户资金流水，若该笔转账批次单的扣款已经发生，则表示批次已经进入转账中，请再次查单确认；
	// PROCESSING:转账中，已开始处理批次内的转账明细单；
	// FINISHED:已完成，批次内的所有转账明细单都已处理完成；
	// CLOSED:已关闭，可查询具体的批次关闭原因确认；
	// 示例值：ACCEPTED
		BatchStatus *string `json:"BatchStatus,omitempty" name:"BatchStatus"`

		// 批次关闭原因。
	// 如果批次单状态为“CLOSED”（已关闭），则有关闭原因；
	// MERCHANT_REVOCATION：商户主动撤销；
	// OVERDUE_CLOSE：系统超时关闭。
	// 示例值：OVERDUE_CLOSE
	// 注意：此字段可能返回 null，表示取不到有效值。
		CloseReason *string `json:"CloseReason,omitempty" name:"CloseReason"`

		// 转账总金额。
	// 转账金额，单位为分。
	// 示例值：4000000
		TotalAmount *uint64 `json:"TotalAmount,omitempty" name:"TotalAmount"`

		// 转账总笔数。
	// 一个转账批次最多允许发起三千笔转账。
	// 示例值：200
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 批次受理成功时返回，遵循rfc3339标准格式。格式为YYYY-MM-DDTHH:mm:ss.sss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss.sss表示时分秒毫秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35.120+08:00表示北京时间2015年05月20日13点29分35秒。
	// 示例值：2015-05-20T13:29:35.120+08:00
	// 注意：此字段可能返回 null，表示取不到有效值。
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// 批次最近一次更新时间，遵循rfc3339标准格式。格式为YYYY-MM-DDTHH:mm:ss.sss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss.sss表示时分秒毫秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35.120+08:00表示北京时间2015年05月20日13点29分35秒。
	// 示例值：2015-05-20T13:29:35.120+08:00
	// 注意：此字段可能返回 null，表示取不到有效值。
		UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

		// 转账成功金额。
	// 转账成功的金额，单位为分，可能随时变化。
	// 示例值：4000000
	// 注意：此字段可能返回 null，表示取不到有效值。
		SuccessAmount *uint64 `json:"SuccessAmount,omitempty" name:"SuccessAmount"`

		// 转账成功的笔数。
	// 可能随时变化。
	// 示例值：200
	// 注意：此字段可能返回 null，表示取不到有效值。
		SuccessNum *uint64 `json:"SuccessNum,omitempty" name:"SuccessNum"`

		// 转账失败金额。
	// 转账失败的金额，单位为分，可能随时变化。
	// 示例值：4000000
	// 注意：此字段可能返回 null，表示取不到有效值。
		FailAmount *uint64 `json:"FailAmount,omitempty" name:"FailAmount"`

		// 转账失败笔数。
	// 可能随时变化。
	// 示例值：200
	// 注意：此字段可能返回 null，表示取不到有效值。
		FailNum *uint64 `json:"FailNum,omitempty" name:"FailNum"`

		// 转账明细列表。
	// 返回明细详情
	// 注意：此字段可能返回 null，表示取不到有效值。
		TransferDetails []*TransferDetailResponse `json:"TransferDetails,omitempty" name:"TransferDetails"`

		// 批次类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
		BatchType *string `json:"BatchType,omitempty" name:"BatchType"`

		// 批次名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
		BatchName *string `json:"BatchName,omitempty" name:"BatchName"`

		// 批次标注。
	// 注意：此字段可能返回 null，表示取不到有效值。
		BatchRemark *string `json:"BatchRemark,omitempty" name:"BatchRemark"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryTransferBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryTransferBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryTransferDetailRequest struct {
	*tchttp.BaseRequest

	// 商户号。
	// 示例值：129284394
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 商家批次单号。
	// 商户系统内部的商家批次单号，此参数只能由数字、字母组成，商户系统内部唯一，UTF8编码，最多32个字符。
	// 示例值：plfk2020042013
	MerchantBatchNo *string `json:"MerchantBatchNo,omitempty" name:"MerchantBatchNo"`

	// 商家明细单号。
	// 商户系统内部的商家明细单号
	// 示例值：plfk2020042013
	MerchantDetailNo *string `json:"MerchantDetailNo,omitempty" name:"MerchantDetailNo"`

	// 微信批次单号。
	// 微信商家转账系统返回的唯一标识。
	// 商家单号（包含批次号和明细单号）和微信单号（包含批次号和明细单号）二者必填其一。
	// 示例值：1030000071100999991182020050700019480001
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// 微信明细单号。
	// 微信区分明细单返回的唯一标识。
	// 示例值：1030000071100999991182020050700019480001
	DetailId *string `json:"DetailId,omitempty" name:"DetailId"`

	// 环境名:
	// release: 现网环境
	// sandbox: 沙箱环境
	// development: 开发环境
	// 缺省: release
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *QueryTransferDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryTransferDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MerchantId")
	delete(f, "MerchantBatchNo")
	delete(f, "MerchantDetailNo")
	delete(f, "BatchId")
	delete(f, "DetailId")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryTransferDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryTransferDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 商户号。
	// 示例值：19300009329
		MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

		// 商家批次单号。
	// 商户系统内部的商家批次单号，此参数只能由数字、字母组成，商户系统内部唯一，UTF8编码，最多32个字符。
	// 示例值：plfk2020042013
		MerchantBatchNo *string `json:"MerchantBatchNo,omitempty" name:"MerchantBatchNo"`

		// 微信批次单号。
	// 微信商家转账系统返回的唯一标识。
	// 示例值：1030000071100999991182020050700019480001
		BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

		// 商家明细单号。
	// 商户系统内部的商家明细单号
	// 示例值：plfk2020042013
		MerchantDetailNo *string `json:"MerchantDetailNo,omitempty" name:"MerchantDetailNo"`

		// 微信明细单号。
	// 微信区分明细单返回的唯一标识。
	// 示例值：1030000071100999991182020050700019480001
		DetailId *string `json:"DetailId,omitempty" name:"DetailId"`

		// 明细状态。
	// PROCESSING：转账中，正在处理，结果未明；
	// SUCCESS：转账成功；
	// FAIL：转账失败，需要确认失败原因以后，再决定是否重新发起地该笔明细的转账。
	// 示例值：SUCCESS
		DetailStatus *string `json:"DetailStatus,omitempty" name:"DetailStatus"`

		// 转账金额。
	// 单位为分。
	// 示例值：200
		TransferAmount *uint64 `json:"TransferAmount,omitempty" name:"TransferAmount"`

		// 失败原因。
	// 如果转账失败则有失败原因
	// ACCOUNT_FROZEN：账户冻结
	// REAL_NAME_CHECK_FAIL：用户未实名
	// NAME_NOT_CORRECT：用户姓名校验失败
	// OPENID_INVALID：Openid校验失败
	// TRANSFER_QUOTA_EXCEED：超过用户单笔收款额度
	// DAY_RECEIVED_QUOTA_EXCEED：超过用户单日收款额度
	// MONTH_RECEIVED_QUOTA_EXCEED：超过用户单月收款额度
	// DAY_RECEIVED_COUNT_EXCEED：超过用户单日收款次数
	// PRODUCT_AUTH_CHECK_FAIL：产品权限校验失败
	// OVERDUE_CLOSE：转账关闭
	// ID_CARD_NOT_CORRECT：用户身份证校验失败
	// ACCOUNT_NOT_EXIST：用户账户不存在
	// TRANSFER_RISK：转账存在风险
	// 示例值：ACCOUNT_FROZEN
	// 注意：此字段可能返回 null，表示取不到有效值。
		FailReason *string `json:"FailReason,omitempty" name:"FailReason"`

		// 转账发起时间。
	// 遵循rfc3339标准格式。格式为YYYY-MM-DDTHH:mm:ss.sss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss.sss表示时分秒毫秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35.120+08:00表示北京时间2015年05月20日13点29分35秒。
	// 示例值：2015-05-20T13:29:35.120+08:00
	// 注意：此字段可能返回 null，表示取不到有效值。
		InitiateTime *string `json:"InitiateTime,omitempty" name:"InitiateTime"`

		// 转账更新时间。
	// 遵循rfc3339标准格式。格式为YYYY-MM-DDTHH:mm:ss.sss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss.sss表示时分秒毫秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35.120+08:00表示北京时间2015年05月20日13点29分35秒。
	// 示例值：2015-05-20T13:29:35.120+08:00
	// 注意：此字段可能返回 null，表示取不到有效值。
		UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

		// 用户名。
	// 示例值：张三
	// 注意：此字段可能返回 null，表示取不到有效值。
		UserName *string `json:"UserName,omitempty" name:"UserName"`

		// 转账备注。
	// 单条转账备注（微信用户会收到该备注）。UTF8编码，最多32字符。
	// 示例值：2020年4月报销
	// 注意：此字段可能返回 null，表示取不到有效值。
		TransferRemark *string `json:"TransferRemark,omitempty" name:"TransferRemark"`

		// 商家绑定公众号APPID。
	// 注意：此字段可能返回 null，表示取不到有效值。
		MerchantAppId *string `json:"MerchantAppId,omitempty" name:"MerchantAppId"`

		// 用户openId。
	// 注意：此字段可能返回 null，表示取不到有效值。
		OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryTransferDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryTransferDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryTransferResultData struct {

	// 平台交易流水号
	TradeSerialNo *string `json:"TradeSerialNo,omitempty" name:"TradeSerialNo"`

	// 订单号
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 交易状态。
	// 0 处理中  
	// 1 提交成功 
	// 2 交易成功 
	// 3 交易失败 
	// 4 未知渠道异常 
	// 99 未知系统异常
	TradeStatus *int64 `json:"TradeStatus,omitempty" name:"TradeStatus"`

	// 业务备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type QueryTransferResultRequest struct {
	*tchttp.BaseRequest

	// 商户号
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 申请商户号的appid或者商户号绑定的appid
	MerchantAppId *string `json:"MerchantAppId,omitempty" name:"MerchantAppId"`

	// 1、 微信企业付款 
	// 2、 支付宝转账 
	// 3、 平安银企直联代发转账
	TransferType *int64 `json:"TransferType,omitempty" name:"TransferType"`

	// 交易流水流水号（与 OrderId 不能同时为空）
	TradeSerialNo *string `json:"TradeSerialNo,omitempty" name:"TradeSerialNo"`

	// 订单号（与 TradeSerialNo 不能同时为空）
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 接入环境。沙箱环境填sandbox。
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *QueryTransferResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryTransferResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MerchantId")
	delete(f, "MerchantAppId")
	delete(f, "TransferType")
	delete(f, "TradeSerialNo")
	delete(f, "OrderId")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryTransferResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryTransferResultResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 错误码。响应成功："SUCCESS"，其他为不成功
		ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

		// 响应消息
		ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

		// 返回结果
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *QueryTransferResultData `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryTransferResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryTransferResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RechargeByThirdPayRequest struct {
	*tchttp.BaseRequest

	// 请求类型 此接口固定填：MemberRechargeThirdPayReq
	RequestType *string `json:"RequestType,omitempty" name:"RequestType"`

	// 商户号
	MerchantCode *string `json:"MerchantCode,omitempty" name:"MerchantCode"`

	// 支付渠道
	PayChannel *string `json:"PayChannel,omitempty" name:"PayChannel"`

	// 子渠道
	PayChannelSubId *int64 `json:"PayChannelSubId,omitempty" name:"PayChannelSubId"`

	// 交易订单号
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 父账户账号，资金汇总账号
	BankAccountNumber *string `json:"BankAccountNumber,omitempty" name:"BankAccountNumber"`

	// 平台短号(银行分配)
	PlatformShortNumber *string `json:"PlatformShortNumber,omitempty" name:"PlatformShortNumber"`

	// 聚鑫分配的安全ID
	MidasSecretId *string `json:"MidasSecretId,omitempty" name:"MidasSecretId"`

	// 聚鑫分配的支付主MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 计费签名
	MidasSignature *string `json:"MidasSignature,omitempty" name:"MidasSignature"`

	// 交易流水号
	TransSequenceNumber *string `json:"TransSequenceNumber,omitempty" name:"TransSequenceNumber"`

	// 子账户账号
	BankSubAccountNumber *string `json:"BankSubAccountNumber,omitempty" name:"BankSubAccountNumber"`

	// 交易手续费，以元为单位
	TransFee *string `json:"TransFee,omitempty" name:"TransFee"`

	// 第三方支付渠道类型 0001-微信 0002-支付宝 0003-京东支付
	ThirdPayChannel *string `json:"ThirdPayChannel,omitempty" name:"ThirdPayChannel"`

	// 第三方渠道商户号
	ThirdPayChannelMerchantCode *string `json:"ThirdPayChannelMerchantCode,omitempty" name:"ThirdPayChannelMerchantCode"`

	// 第三方渠道订单号或流水号
	ThirdPayChannelOrderId *string `json:"ThirdPayChannelOrderId,omitempty" name:"ThirdPayChannelOrderId"`

	// 交易金额
	CurrencyAmount *string `json:"CurrencyAmount,omitempty" name:"CurrencyAmount"`

	// 单位，1：元，2：角，3：分
	CurrencyUnit *string `json:"CurrencyUnit,omitempty" name:"CurrencyUnit"`

	// 币种
	CurrencyType *string `json:"CurrencyType,omitempty" name:"CurrencyType"`

	// 交易网会员代码
	TransNetMemberCode *string `json:"TransNetMemberCode,omitempty" name:"TransNetMemberCode"`

	// midas环境参数
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`

	// 保留域
	ReservedMessage *string `json:"ReservedMessage,omitempty" name:"ReservedMessage"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *RechargeByThirdPayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RechargeByThirdPayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RequestType")
	delete(f, "MerchantCode")
	delete(f, "PayChannel")
	delete(f, "PayChannelSubId")
	delete(f, "OrderId")
	delete(f, "BankAccountNumber")
	delete(f, "PlatformShortNumber")
	delete(f, "MidasSecretId")
	delete(f, "MidasAppId")
	delete(f, "MidasSignature")
	delete(f, "TransSequenceNumber")
	delete(f, "BankSubAccountNumber")
	delete(f, "TransFee")
	delete(f, "ThirdPayChannel")
	delete(f, "ThirdPayChannelMerchantCode")
	delete(f, "ThirdPayChannelOrderId")
	delete(f, "CurrencyAmount")
	delete(f, "CurrencyUnit")
	delete(f, "CurrencyType")
	delete(f, "TransNetMemberCode")
	delete(f, "MidasEnvironment")
	delete(f, "ReservedMessage")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RechargeByThirdPayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RechargeByThirdPayResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 保留字段
	// 注意：此字段可能返回 null，表示取不到有效值。
		ReservedMessage *string `json:"ReservedMessage,omitempty" name:"ReservedMessage"`

		// 银行流水号
	// 注意：此字段可能返回 null，表示取不到有效值。
		FrontSequenceNumber *string `json:"FrontSequenceNumber,omitempty" name:"FrontSequenceNumber"`

		// 请求类型
		RequestType *string `json:"RequestType,omitempty" name:"RequestType"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RechargeByThirdPayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RechargeByThirdPayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RechargeMemberThirdPayRequest struct {
	*tchttp.BaseRequest

	// STRING(32)，交易网会代码
	TranNetMemberCode *string `json:"TranNetMemberCode,omitempty" name:"TranNetMemberCode"`

	// STRING(20)，会员充值金额
	MemberFillAmt *string `json:"MemberFillAmt,omitempty" name:"MemberFillAmt"`

	// STRING(20)，手续费金额
	Commission *string `json:"Commission,omitempty" name:"Commission"`

	// STRING(3)，币种。如RMB
	Ccy *string `json:"Ccy,omitempty" name:"Ccy"`

	// STRING(20)，支付渠道类型。
	// 0001-微信
	// 0002-支付宝
	// 0003-京东支付
	PayChannelType *string `json:"PayChannelType,omitempty" name:"PayChannelType"`

	// STRING(50)，支付渠道所分配的商户号
	PayChannelAssignMerNo *string `json:"PayChannelAssignMerNo,omitempty" name:"PayChannelAssignMerNo"`

	// STRING(52)，支付渠道交易流水号
	PayChannelTranSeqNo *string `json:"PayChannelTranSeqNo,omitempty" name:"PayChannelTranSeqNo"`

	// STRING(52)，电商见证宝订单号
	EjzbOrderNo *string `json:"EjzbOrderNo,omitempty" name:"EjzbOrderNo"`

	// String(22)，商户号
	MrchCode *string `json:"MrchCode,omitempty" name:"MrchCode"`

	// STRING(500)，电商见证宝订单内容
	EjzbOrderContent *string `json:"EjzbOrderContent,omitempty" name:"EjzbOrderContent"`

	// STRING(300)，备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// STRING(300)，保留域1
	ReservedMsgOne *string `json:"ReservedMsgOne,omitempty" name:"ReservedMsgOne"`

	// STRING(300)，保留域2
	ReservedMsgTwo *string `json:"ReservedMsgTwo,omitempty" name:"ReservedMsgTwo"`

	// STRING(300)，保留域3
	ReservedMsgThree *string `json:"ReservedMsgThree,omitempty" name:"ReservedMsgThree"`

	// STRING(12)，接入环境，默认接入沙箱环境。接入正式环境填"prod"
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *RechargeMemberThirdPayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RechargeMemberThirdPayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TranNetMemberCode")
	delete(f, "MemberFillAmt")
	delete(f, "Commission")
	delete(f, "Ccy")
	delete(f, "PayChannelType")
	delete(f, "PayChannelAssignMerNo")
	delete(f, "PayChannelTranSeqNo")
	delete(f, "EjzbOrderNo")
	delete(f, "MrchCode")
	delete(f, "EjzbOrderContent")
	delete(f, "Remark")
	delete(f, "ReservedMsgOne")
	delete(f, "ReservedMsgTwo")
	delete(f, "ReservedMsgThree")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RechargeMemberThirdPayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RechargeMemberThirdPayResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// String(20)，返回码
		TxnReturnCode *string `json:"TxnReturnCode,omitempty" name:"TxnReturnCode"`

		// String(100)，返回信息
		TxnReturnMsg *string `json:"TxnReturnMsg,omitempty" name:"TxnReturnMsg"`

		// String(22)，交易流水号
		CnsmrSeqNo *string `json:"CnsmrSeqNo,omitempty" name:"CnsmrSeqNo"`

		// STRING(52)，前置流水号
	// 注意：此字段可能返回 null，表示取不到有效值。
		FrontSeqNo *string `json:"FrontSeqNo,omitempty" name:"FrontSeqNo"`

		// STRING(20)，会员子账户交易前可用余额
	// 注意：此字段可能返回 null，表示取不到有效值。
		MemberSubAcctPreAvailBal *string `json:"MemberSubAcctPreAvailBal,omitempty" name:"MemberSubAcctPreAvailBal"`

		// STRING(300)，保留域1
	// 注意：此字段可能返回 null，表示取不到有效值。
		ReservedMsgOne *string `json:"ReservedMsgOne,omitempty" name:"ReservedMsgOne"`

		// STRING(300)，保留域2
	// 注意：此字段可能返回 null，表示取不到有效值。
		ReservedMsgTwo *string `json:"ReservedMsgTwo,omitempty" name:"ReservedMsgTwo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RechargeMemberThirdPayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RechargeMemberThirdPayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RefundMemberTransactionRequest struct {
	*tchttp.BaseRequest

	// 转出见证子账户的户名
	OutSubAccountName *string `json:"OutSubAccountName,omitempty" name:"OutSubAccountName"`

	// 转入见证子账户的户名
	InSubAccountName *string `json:"InSubAccountName,omitempty" name:"InSubAccountName"`

	// 子渠道
	PayChannelSubId *int64 `json:"PayChannelSubId,omitempty" name:"PayChannelSubId"`

	// 转出见证子账户账号
	OutSubAccountNumber *string `json:"OutSubAccountNumber,omitempty" name:"OutSubAccountNumber"`

	// 计费签名
	MidasSignature *string `json:"MidasSignature,omitempty" name:"MidasSignature"`

	// 转入见证子账户账号
	InSubAccountNumber *string `json:"InSubAccountNumber,omitempty" name:"InSubAccountNumber"`

	// 聚鑫分配的安全ID
	MidasSecretId *string `json:"MidasSecretId,omitempty" name:"MidasSecretId"`

	// 父账户账号，资金汇总账号
	BankAccountNumber *string `json:"BankAccountNumber,omitempty" name:"BankAccountNumber"`

	// 原老订单流水号
	OldTransSequenceNumber *string `json:"OldTransSequenceNumber,omitempty" name:"OldTransSequenceNumber"`

	// 银行注册商户号
	MerchantCode *string `json:"MerchantCode,omitempty" name:"MerchantCode"`

	// 请求类型，固定为MemberTransactionRefundReq
	RequestType *string `json:"RequestType,omitempty" name:"RequestType"`

	// 交易金额
	CurrencyAmount *string `json:"CurrencyAmount,omitempty" name:"CurrencyAmount"`

	// 交易流水号
	TransSequenceNumber *string `json:"TransSequenceNumber,omitempty" name:"TransSequenceNumber"`

	// 渠道
	PayChannel *string `json:"PayChannel,omitempty" name:"PayChannel"`

	// 原订单号
	OldOrderId *string `json:"OldOrderId,omitempty" name:"OldOrderId"`

	// 聚鑫分配的支付主MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 订单号
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// Midas环境标识 release 现网环境 sandbox 沙箱环境
	// development 开发环境
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`

	// 转出子账户交易网会员代码
	OutTransNetMemberCode *string `json:"OutTransNetMemberCode,omitempty" name:"OutTransNetMemberCode"`

	// 转入子账户交易网会员代码
	InTransNetMemberCode *string `json:"InTransNetMemberCode,omitempty" name:"InTransNetMemberCode"`

	// 保留域
	ReservedMessage *string `json:"ReservedMessage,omitempty" name:"ReservedMessage"`

	// 平台短号(银行分配)
	PlatformShortNumber *string `json:"PlatformShortNumber,omitempty" name:"PlatformShortNumber"`

	// 0-登记挂账，1-撤销挂账
	TransType *string `json:"TransType,omitempty" name:"TransType"`

	// 交易手续费
	TransFee *string `json:"TransFee,omitempty" name:"TransFee"`
}

func (r *RefundMemberTransactionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefundMemberTransactionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OutSubAccountName")
	delete(f, "InSubAccountName")
	delete(f, "PayChannelSubId")
	delete(f, "OutSubAccountNumber")
	delete(f, "MidasSignature")
	delete(f, "InSubAccountNumber")
	delete(f, "MidasSecretId")
	delete(f, "BankAccountNumber")
	delete(f, "OldTransSequenceNumber")
	delete(f, "MerchantCode")
	delete(f, "RequestType")
	delete(f, "CurrencyAmount")
	delete(f, "TransSequenceNumber")
	delete(f, "PayChannel")
	delete(f, "OldOrderId")
	delete(f, "MidasAppId")
	delete(f, "OrderId")
	delete(f, "MidasEnvironment")
	delete(f, "OutTransNetMemberCode")
	delete(f, "InTransNetMemberCode")
	delete(f, "ReservedMessage")
	delete(f, "PlatformShortNumber")
	delete(f, "TransType")
	delete(f, "TransFee")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RefundMemberTransactionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RefundMemberTransactionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 请求类型
		RequestType *string `json:"RequestType,omitempty" name:"RequestType"`

		// 银行流水号
		FrontSequenceNumber *string `json:"FrontSequenceNumber,omitempty" name:"FrontSequenceNumber"`

		// 保留域
		ReservedMessage *string `json:"ReservedMessage,omitempty" name:"ReservedMessage"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RefundMemberTransactionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefundMemberTransactionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RefundOrderRequest struct {
	*tchttp.BaseRequest

	// 进件成功后返给商户方的AppId
	MerchantAppId *string `json:"MerchantAppId,omitempty" name:"MerchantAppId"`

	// 平台流水号。消费订单发起成功后，返回的平台唯一订单号。
	OrderNo *string `json:"OrderNo,omitempty" name:"OrderNo"`
}

func (r *RefundOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefundOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MerchantAppId")
	delete(f, "OrderNo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RefundOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RefundOrderResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 进件成功后返给商户方的AppId
		MerchantAppId *string `json:"MerchantAppId,omitempty" name:"MerchantAppId"`

		// 平台流水号。消费订单发起成功后，返回的平台唯一订单号。
		OrderNo *string `json:"OrderNo,omitempty" name:"OrderNo"`

		// 订单退款状态。0-退款失败
	// 1-退款成功 
	// 2-可疑状态
		Status *string `json:"Status,omitempty" name:"Status"`

		// 订单退款状态描述
	// 注意：此字段可能返回 null，表示取不到有效值。
		Description *string `json:"Description,omitempty" name:"Description"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RefundOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefundOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RefundOutSubOrderRefundList struct {

	// 平台应退金额
	PlatformRefundAmt *int64 `json:"PlatformRefundAmt,omitempty" name:"PlatformRefundAmt"`

	// 子订单退款金额
	RefundAmt *int64 `json:"RefundAmt,omitempty" name:"RefundAmt"`

	// 商家应退金额
	SubMchRefundAmt *int64 `json:"SubMchRefundAmt,omitempty" name:"SubMchRefundAmt"`

	// 子订单号
	SubOutTradeNo *string `json:"SubOutTradeNo,omitempty" name:"SubOutTradeNo"`

	// 子退款单号，调用方需要保证 全局唯一性
	SubRefundId *string `json:"SubRefundId,omitempty" name:"SubRefundId"`
}

type RefundRequest struct {
	*tchttp.BaseRequest

	// 用户ID，长度不小于5位， 仅支持字母和数字的组合
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 退款订单号，仅支持数字、 字母、下划线（_）、横杠字 符（-）、点（.）的组合
	RefundId *string `json:"RefundId,omitempty" name:"RefundId"`

	// 聚鑫分配的支付主MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 退款金额，单位：分。备注：当该字段为空或者为0 时，系统会默认使用订单当 实付金额作为退款金额
	TotalRefundAmt *int64 `json:"TotalRefundAmt,omitempty" name:"TotalRefundAmt"`

	// 聚鑫分配的安全ID
	MidasSecretId *string `json:"MidasSecretId,omitempty" name:"MidasSecretId"`

	// 按照聚鑫安全密钥计算的签名
	MidasSignature *string `json:"MidasSignature,omitempty" name:"MidasSignature"`

	// 商品订单，仅支持数字、字 母、下划线（_）、横杠字符 （-）、点（.）的组合。  OutTradeNo ,TransactionId 二选一,不能都为空,优先使用 OutTradeNo
	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`

	// 结算应收金额，单位：分
	MchRefundAmt *int64 `json:"MchRefundAmt,omitempty" name:"MchRefundAmt"`

	// 调用下单接口获取的聚鑫交 易订单。  OutTradeNo ,TransactionId 二选一,不能都为空,优先使用 OutTradeNo
	TransactionId *string `json:"TransactionId,omitempty" name:"TransactionId"`

	// 平台应收金额，单位：分
	PlatformRefundAmt *int64 `json:"PlatformRefundAmt,omitempty" name:"PlatformRefundAmt"`

	// 支持多个子订单批量退款单 个子订单退款支持传 SubOutTradeNo ，也支持传 SubOutTradeNoList ，都传的时候以 SubOutTradeNoList 为准。  如果传了子单退款细节，外 部不需要再传退款金额，平 台应退，商户应退金额，我 们可以直接根据子单退款算出来总和。
	SubOrderRefundList []*RefundOutSubOrderRefundList `json:"SubOrderRefundList,omitempty" name:"SubOrderRefundList"`

	// 环境名:
	// release: 现网环境
	// sandbox: 沙箱环境
	// development: 开发环境
	// 缺省: release
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`
}

func (r *RefundRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefundRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "RefundId")
	delete(f, "MidasAppId")
	delete(f, "TotalRefundAmt")
	delete(f, "MidasSecretId")
	delete(f, "MidasSignature")
	delete(f, "OutTradeNo")
	delete(f, "MchRefundAmt")
	delete(f, "TransactionId")
	delete(f, "PlatformRefundAmt")
	delete(f, "SubOrderRefundList")
	delete(f, "MidasEnvironment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RefundRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RefundResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RefundResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefundResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegisterBehaviorRequest struct {
	*tchttp.BaseRequest

	// 聚鑫分配的支付主MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 聚鑫计费SubAppId，代表子商户
	SubAppId *string `json:"SubAppId,omitempty" name:"SubAppId"`

	// 聚鑫分配的安全ID
	MidasSecretId *string `json:"MidasSecretId,omitempty" name:"MidasSecretId"`

	// 按照聚鑫安全密钥计算的签名
	MidasSignature *string `json:"MidasSignature,omitempty" name:"MidasSignature"`

	// 功能标志
	// 1：登记行为记录信息
	// 2：查询补录信息
	FunctionFlag *int64 `json:"FunctionFlag,omitempty" name:"FunctionFlag"`

	// 环境名:
	// release: 现网环境
	// sandbox: 沙箱环境
	// development: 开发环境
	// 缺省: release
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`

	// 操作点击时间
	// yyyyMMddHHmmss
	// 功能标志FunctionFlag=1时必输
	OperationClickTime *string `json:"OperationClickTime,omitempty" name:"OperationClickTime"`

	// IP地址
	// 功能标志FunctionFlag=1时必输
	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`

	// MAC地址
	// 功能标志FunctionFlag=1时必输
	MacAddress *string `json:"MacAddress,omitempty" name:"MacAddress"`

	// 签约渠道
	// 1:  App
	// 2:  平台H5网页
	// 3：公众号
	// 4：小程序
	// 功能标志FunctionFlag=1时必输
	SignChannel *int64 `json:"SignChannel,omitempty" name:"SignChannel"`
}

func (r *RegisterBehaviorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterBehaviorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MidasAppId")
	delete(f, "SubAppId")
	delete(f, "MidasSecretId")
	delete(f, "MidasSignature")
	delete(f, "FunctionFlag")
	delete(f, "MidasEnvironment")
	delete(f, "OperationClickTime")
	delete(f, "IpAddress")
	delete(f, "MacAddress")
	delete(f, "SignChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RegisterBehaviorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RegisterBehaviorResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 补录是否成功标志
	// 功能标志为2时存在。
	// S：成功
	// F：失败
	// 注意：此字段可能返回 null，表示取不到有效值。
		ReplenishSuccessFlag *string `json:"ReplenishSuccessFlag,omitempty" name:"ReplenishSuccessFlag"`

		// 签约信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		RegisterInfo *RegisterInfo `json:"RegisterInfo,omitempty" name:"RegisterInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RegisterBehaviorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterBehaviorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegisterBillRequest struct {
	*tchttp.BaseRequest

	// 请求类型此接口固定填：RegBillSupportWithdrawReq
	RequestType *string `json:"RequestType,omitempty" name:"RequestType"`

	// 商户号
	MerchantCode *string `json:"MerchantCode,omitempty" name:"MerchantCode"`

	// 支付渠道
	PayChannel *string `json:"PayChannel,omitempty" name:"PayChannel"`

	// 子渠道
	PayChannelSubId *int64 `json:"PayChannelSubId,omitempty" name:"PayChannelSubId"`

	// 交易订单号
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 父账户账号，资金汇总账号
	BankAccountNo *string `json:"BankAccountNo,omitempty" name:"BankAccountNo"`

	// 平台短号(银行分配)
	PlatformShortNo *string `json:"PlatformShortNo,omitempty" name:"PlatformShortNo"`

	// 聚鑫分配的安全ID
	MidasSecretId *string `json:"MidasSecretId,omitempty" name:"MidasSecretId"`

	// 聚鑫分配的支付主MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 计费签名
	MidasSignature *string `json:"MidasSignature,omitempty" name:"MidasSignature"`

	// 交易流水号
	TransSeqNo *string `json:"TransSeqNo,omitempty" name:"TransSeqNo"`

	// 暂未使用，默认传0.0
	TranFee *string `json:"TranFee,omitempty" name:"TranFee"`

	// 挂账金额，以元为单位
	OrderAmt *string `json:"OrderAmt,omitempty" name:"OrderAmt"`

	// 子账户账号
	BankSubAccountNo *string `json:"BankSubAccountNo,omitempty" name:"BankSubAccountNo"`

	// 交易网会员代码
	TranNetMemberCode *string `json:"TranNetMemberCode,omitempty" name:"TranNetMemberCode"`

	// 0,登记挂账，1，撤销挂账
	TranType *string `json:"TranType,omitempty" name:"TranType"`

	// 保留域
	ReservedMessage *string `json:"ReservedMessage,omitempty" name:"ReservedMessage"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Midas环境参数
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`
}

func (r *RegisterBillRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterBillRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RequestType")
	delete(f, "MerchantCode")
	delete(f, "PayChannel")
	delete(f, "PayChannelSubId")
	delete(f, "OrderId")
	delete(f, "BankAccountNo")
	delete(f, "PlatformShortNo")
	delete(f, "MidasSecretId")
	delete(f, "MidasAppId")
	delete(f, "MidasSignature")
	delete(f, "TransSeqNo")
	delete(f, "TranFee")
	delete(f, "OrderAmt")
	delete(f, "BankSubAccountNo")
	delete(f, "TranNetMemberCode")
	delete(f, "TranType")
	delete(f, "ReservedMessage")
	delete(f, "Remark")
	delete(f, "MidasEnvironment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RegisterBillRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RegisterBillResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 银行流水号
		FrontSeqNo *string `json:"FrontSeqNo,omitempty" name:"FrontSeqNo"`

		// 保留字段
		ReservedMessage *string `json:"ReservedMessage,omitempty" name:"ReservedMessage"`

		// 请求类型
		RequestType *string `json:"RequestType,omitempty" name:"RequestType"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RegisterBillResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterBillResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegisterBillSupportWithdrawRequest struct {
	*tchttp.BaseRequest

	// STRING(32)，交易网会员代码
	TranNetMemberCode *string `json:"TranNetMemberCode,omitempty" name:"TranNetMemberCode"`

	// STRING(50)，订单号
	OrderNo *string `json:"OrderNo,omitempty" name:"OrderNo"`

	// STRING(20)，挂账金额（包含交易费用）
	SuspendAmt *string `json:"SuspendAmt,omitempty" name:"SuspendAmt"`

	// STRING(20)，交易费用（暂未使用，默认传0.0）
	TranFee *string `json:"TranFee,omitempty" name:"TranFee"`

	// String(22)，商户号（签约客户号）
	MrchCode *string `json:"MrchCode,omitempty" name:"MrchCode"`

	// STRING(300)，备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// STRING(300)，保留域1
	ReservedMsgOne *string `json:"ReservedMsgOne,omitempty" name:"ReservedMsgOne"`

	// STRING(300)，保留域2
	ReservedMsgTwo *string `json:"ReservedMsgTwo,omitempty" name:"ReservedMsgTwo"`

	// STRING(300)，保留域3
	ReservedMsgThree *string `json:"ReservedMsgThree,omitempty" name:"ReservedMsgThree"`

	// STRING(12)，接入环境，默认接入沙箱环境。接入正式环境填"prod"
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *RegisterBillSupportWithdrawRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterBillSupportWithdrawRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TranNetMemberCode")
	delete(f, "OrderNo")
	delete(f, "SuspendAmt")
	delete(f, "TranFee")
	delete(f, "MrchCode")
	delete(f, "Remark")
	delete(f, "ReservedMsgOne")
	delete(f, "ReservedMsgTwo")
	delete(f, "ReservedMsgThree")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RegisterBillSupportWithdrawRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RegisterBillSupportWithdrawResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// String(20)，返回码
		TxnReturnCode *string `json:"TxnReturnCode,omitempty" name:"TxnReturnCode"`

		// String(100)，返回信息
		TxnReturnMsg *string `json:"TxnReturnMsg,omitempty" name:"TxnReturnMsg"`

		// STRING(52)，见证系统流水号
	// 注意：此字段可能返回 null，表示取不到有效值。
		FrontSeqNo *string `json:"FrontSeqNo,omitempty" name:"FrontSeqNo"`

		// String(22)，交易流水号
	// 注意：此字段可能返回 null，表示取不到有效值。
		CnsmrSeqNo *string `json:"CnsmrSeqNo,omitempty" name:"CnsmrSeqNo"`

		// STRING(1027)，保留域
	// 注意：此字段可能返回 null，表示取不到有效值。
		ReservedMsg *string `json:"ReservedMsg,omitempty" name:"ReservedMsg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RegisterBillSupportWithdrawResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterBillSupportWithdrawResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegisterInfo struct {

	// 法人证件号码
	// 注意：此字段可能返回 null，表示取不到有效值。
	LegalPersonIdCode *string `json:"LegalPersonIdCode,omitempty" name:"LegalPersonIdCode"`

	// 法人证件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	LegalPersonIdType *string `json:"LegalPersonIdType,omitempty" name:"LegalPersonIdType"`

	// 法人名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LegalPersonName *string `json:"LegalPersonName,omitempty" name:"LegalPersonName"`

	// 公司证件号码
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationCode *string `json:"OrganizationCode,omitempty" name:"OrganizationCode"`

	// 公司名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationName *string `json:"OrganizationName,omitempty" name:"OrganizationName"`

	// 公司证件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationType *string `json:"OrganizationType,omitempty" name:"OrganizationType"`
}

type ResponseQueryContract struct {

	// 第三方渠道错误码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalReturnCode *string `json:"ExternalReturnCode,omitempty" name:"ExternalReturnCode"`

	// 第三方渠道错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalReturnMessage *string `json:"ExternalReturnMessage,omitempty" name:"ExternalReturnMessage"`

	// 第三方渠道返回的原始数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalReturnData *string `json:"ExternalReturnData,omitempty" name:"ExternalReturnData"`

	// 米大师内部商户号
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 米大师内部子商户号
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelSubMerchantId *string `json:"ChannelSubMerchantId,omitempty" name:"ChannelSubMerchantId"`

	// 米大师内部应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelAppId *string `json:"ChannelAppId,omitempty" name:"ChannelAppId"`

	// 米大师内部子应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelSubAppId *string `json:"ChannelSubAppId,omitempty" name:"ChannelSubAppId"`

	// 渠道名称
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 返回的合约信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReturnContractInfo *ReturnContractInfo `json:"ReturnContractInfo,omitempty" name:"ReturnContractInfo"`

	// 签约通知地址
	NotifyUrl *string `json:"NotifyUrl,omitempty" name:"NotifyUrl"`
}

type ResponseTerminateContract struct {

	// 第三方渠道错误码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalReturnCode *string `json:"ExternalReturnCode,omitempty" name:"ExternalReturnCode"`

	// 第三方渠道错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalReturnMessage *string `json:"ExternalReturnMessage,omitempty" name:"ExternalReturnMessage"`

	// 第三方渠道返回的原始数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalReturnData *string `json:"ExternalReturnData,omitempty" name:"ExternalReturnData"`
}

type ReturnContractInfo struct {

	// 合约信息
	ContractInfo *ContractInfo `json:"ContractInfo,omitempty" name:"ContractInfo"`

	// 米大师内部生成的合约信息
	ChannelReturnContractInfo *ChannelReturnContractInfo `json:"ChannelReturnContractInfo,omitempty" name:"ChannelReturnContractInfo"`

	// 第三方渠道合约信息
	ExternalReturnContractInfo *ExternalReturnContractInfo `json:"ExternalReturnContractInfo,omitempty" name:"ExternalReturnContractInfo"`
}

type RevResigterBillSupportWithdrawRequest struct {
	*tchttp.BaseRequest

	// String(22)，商户号（签约客户号）
	MrchCode *string `json:"MrchCode,omitempty" name:"MrchCode"`

	// STRING(32)，交易网会员代码
	TranNetMemberCode *string `json:"TranNetMemberCode,omitempty" name:"TranNetMemberCode"`

	// STRING(30)，原订单号（RegisterBillSupportWithdraw接口中的订单号）
	OldOrderNo *string `json:"OldOrderNo,omitempty" name:"OldOrderNo"`

	// STRING(20)，撤销金额（支持部分撤销，不能大于原订单可用金额，包含交易费用）
	CancelAmt *string `json:"CancelAmt,omitempty" name:"CancelAmt"`

	// STRING(20)，交易费用（暂未使用，默认传0.0）
	TranFee *string `json:"TranFee,omitempty" name:"TranFee"`

	// STRING(300)，备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// STRING(300)，保留域1
	ReservedMsgOne *string `json:"ReservedMsgOne,omitempty" name:"ReservedMsgOne"`

	// STRING(300)，保留域2
	ReservedMsgTwo *string `json:"ReservedMsgTwo,omitempty" name:"ReservedMsgTwo"`

	// STRING(300)，保留域3
	ReservedMsgThree *string `json:"ReservedMsgThree,omitempty" name:"ReservedMsgThree"`

	// STRING(12)，接入环境，默认接入沙箱环境。接入正式环境填"prod"
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *RevResigterBillSupportWithdrawRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RevResigterBillSupportWithdrawRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MrchCode")
	delete(f, "TranNetMemberCode")
	delete(f, "OldOrderNo")
	delete(f, "CancelAmt")
	delete(f, "TranFee")
	delete(f, "Remark")
	delete(f, "ReservedMsgOne")
	delete(f, "ReservedMsgTwo")
	delete(f, "ReservedMsgThree")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RevResigterBillSupportWithdrawRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RevResigterBillSupportWithdrawResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// String(20)，返回码
		TxnReturnCode *string `json:"TxnReturnCode,omitempty" name:"TxnReturnCode"`

		// String(100)，返回信息
		TxnReturnMsg *string `json:"TxnReturnMsg,omitempty" name:"TxnReturnMsg"`

		// String(22)，交易流水号
		CnsmrSeqNo *string `json:"CnsmrSeqNo,omitempty" name:"CnsmrSeqNo"`

		// STRING(52)，见证系统流水号
	// 注意：此字段可能返回 null，表示取不到有效值。
		FrontSeqNo *string `json:"FrontSeqNo,omitempty" name:"FrontSeqNo"`

		// STRING(1027)，保留域
	// 注意：此字段可能返回 null，表示取不到有效值。
		ReservedMsg *string `json:"ReservedMsg,omitempty" name:"ReservedMsg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RevResigterBillSupportWithdrawResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RevResigterBillSupportWithdrawResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReviseMbrPropertyRequest struct {
	*tchttp.BaseRequest

	// String(22)，商户号（签约客户号）
	MrchCode *string `json:"MrchCode,omitempty" name:"MrchCode"`

	// STRING(50)，见证子账户的账号
	SubAcctNo *string `json:"SubAcctNo,omitempty" name:"SubAcctNo"`

	// STRING(10)，会员属性（00-普通子账号; SH-商户子账户。暂时只支持00-普通子账号改为SH-商户子账户）
	MemberProperty *string `json:"MemberProperty,omitempty" name:"MemberProperty"`

	// STRING(1027)，保留域
	ReservedMsg *string `json:"ReservedMsg,omitempty" name:"ReservedMsg"`

	// STRING(12)，接入环境，默认接入沙箱环境。接入正式环境填"prod"
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *ReviseMbrPropertyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReviseMbrPropertyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MrchCode")
	delete(f, "SubAcctNo")
	delete(f, "MemberProperty")
	delete(f, "ReservedMsg")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReviseMbrPropertyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ReviseMbrPropertyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// String(20)，返回码
		TxnReturnCode *string `json:"TxnReturnCode,omitempty" name:"TxnReturnCode"`

		// String(100)，返回信息
		TxnReturnMsg *string `json:"TxnReturnMsg,omitempty" name:"TxnReturnMsg"`

		// String(22)，交易流水号
		CnsmrSeqNo *string `json:"CnsmrSeqNo,omitempty" name:"CnsmrSeqNo"`

		// STRING(1027)，保留域
	// 注意：此字段可能返回 null，表示取不到有效值。
		ReservedMsg *string `json:"ReservedMsg,omitempty" name:"ReservedMsg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReviseMbrPropertyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReviseMbrPropertyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RevokeMemberRechargeThirdPayRequest struct {
	*tchttp.BaseRequest

	// STRING(52)，原充值的前置流水号
	OldFillFrontSeqNo *string `json:"OldFillFrontSeqNo,omitempty" name:"OldFillFrontSeqNo"`

	// STRING(20)，原充值的支付渠道类型
	OldFillPayChannelType *string `json:"OldFillPayChannelType,omitempty" name:"OldFillPayChannelType"`

	// STRING(52)，原充值的支付渠道交易流水号
	OldPayChannelTranSeqNo *string `json:"OldPayChannelTranSeqNo,omitempty" name:"OldPayChannelTranSeqNo"`

	// STRING(52)，原充值的电商见证宝订单号
	OldFillEjzbOrderNo *string `json:"OldFillEjzbOrderNo,omitempty" name:"OldFillEjzbOrderNo"`

	// STRING(20)，申请撤销的会员金额
	ApplyCancelMemberAmt *string `json:"ApplyCancelMemberAmt,omitempty" name:"ApplyCancelMemberAmt"`

	// STRING(20)，申请撤销的手续费金额
	ApplyCancelCommission *string `json:"ApplyCancelCommission,omitempty" name:"ApplyCancelCommission"`

	// String(22)，商户号
	MrchCode *string `json:"MrchCode,omitempty" name:"MrchCode"`

	// STRING(300)，备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// STRING(300)，保留域1
	ReservedMsgOne *string `json:"ReservedMsgOne,omitempty" name:"ReservedMsgOne"`

	// STRING(300)，保留域2
	ReservedMsgTwo *string `json:"ReservedMsgTwo,omitempty" name:"ReservedMsgTwo"`

	// STRING(300)，保留域3
	ReservedMsgThree *string `json:"ReservedMsgThree,omitempty" name:"ReservedMsgThree"`

	// STRING(12)，接入环境，默认接入沙箱环境。接入正式环境填"prod"
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *RevokeMemberRechargeThirdPayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RevokeMemberRechargeThirdPayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OldFillFrontSeqNo")
	delete(f, "OldFillPayChannelType")
	delete(f, "OldPayChannelTranSeqNo")
	delete(f, "OldFillEjzbOrderNo")
	delete(f, "ApplyCancelMemberAmt")
	delete(f, "ApplyCancelCommission")
	delete(f, "MrchCode")
	delete(f, "Remark")
	delete(f, "ReservedMsgOne")
	delete(f, "ReservedMsgTwo")
	delete(f, "ReservedMsgThree")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RevokeMemberRechargeThirdPayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RevokeMemberRechargeThirdPayResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// String(20)，返回码
		TxnReturnCode *string `json:"TxnReturnCode,omitempty" name:"TxnReturnCode"`

		// String(100)，返回信息
		TxnReturnMsg *string `json:"TxnReturnMsg,omitempty" name:"TxnReturnMsg"`

		// String(22)，交易流水号
		CnsmrSeqNo *string `json:"CnsmrSeqNo,omitempty" name:"CnsmrSeqNo"`

		// STRING(52)，前置流水号
	// 注意：此字段可能返回 null，表示取不到有效值。
		FrontSeqNo *string `json:"FrontSeqNo,omitempty" name:"FrontSeqNo"`

		// STRING(300)，保留域1
	// 注意：此字段可能返回 null，表示取不到有效值。
		ReservedMsgOne *string `json:"ReservedMsgOne,omitempty" name:"ReservedMsgOne"`

		// STRING(300)，保留域2
	// 注意：此字段可能返回 null，表示取不到有效值。
		ReservedMsgTwo *string `json:"ReservedMsgTwo,omitempty" name:"ReservedMsgTwo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RevokeMemberRechargeThirdPayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RevokeMemberRechargeThirdPayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RevokeRechargeByThirdPayRequest struct {
	*tchttp.BaseRequest

	// 请求类型此接口固定填：RevokeMemberRechargeThirdPayReq
	RequestType *string `json:"RequestType,omitempty" name:"RequestType"`

	// 商户号
	MerchantCode *string `json:"MerchantCode,omitempty" name:"MerchantCode"`

	// 支付渠道
	PayChannel *string `json:"PayChannel,omitempty" name:"PayChannel"`

	// 子渠道
	PayChannelSubId *int64 `json:"PayChannelSubId,omitempty" name:"PayChannelSubId"`

	// 原始充值交易订单号
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 父账户账号，资金汇总账号
	BankAccountNumber *string `json:"BankAccountNumber,omitempty" name:"BankAccountNumber"`

	// 平台短号(银行分配)
	PlatformShortNumber *string `json:"PlatformShortNumber,omitempty" name:"PlatformShortNumber"`

	// 聚鑫分配的安全ID
	MidasSecretId *string `json:"MidasSecretId,omitempty" name:"MidasSecretId"`

	// 聚鑫分配的支付主MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 计费签名
	MidasSignature *string `json:"MidasSignature,omitempty" name:"MidasSignature"`

	// 交易流水号
	TransSequenceNumber *string `json:"TransSequenceNumber,omitempty" name:"TransSequenceNumber"`

	// 申请撤销的手续费金额,以元为单位
	TransFee *string `json:"TransFee,omitempty" name:"TransFee"`

	// 第三方支付渠道类型 0001-微信 0002-支付宝 0003-京东支付
	ThirdPayChannel *string `json:"ThirdPayChannel,omitempty" name:"ThirdPayChannel"`

	// 第三方渠道订单号或流水号
	ThirdPayChannelOrderId *string `json:"ThirdPayChannelOrderId,omitempty" name:"ThirdPayChannelOrderId"`

	// 充值接口银行返回的流水号(FrontSeqNo)
	OldFrontSequenceNumber *string `json:"OldFrontSequenceNumber,omitempty" name:"OldFrontSequenceNumber"`

	// 申请撤销的金额
	CurrencyAmount *string `json:"CurrencyAmount,omitempty" name:"CurrencyAmount"`

	// 单位，1：元，2：角，3：分 目前固定填1
	CurrencyUnit *string `json:"CurrencyUnit,omitempty" name:"CurrencyUnit"`

	// 币种 目前固定填RMB
	CurrencyType *string `json:"CurrencyType,omitempty" name:"CurrencyType"`

	// Midas环境标识
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`

	// 保留域
	ReservedMessage *string `json:"ReservedMessage,omitempty" name:"ReservedMessage"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *RevokeRechargeByThirdPayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RevokeRechargeByThirdPayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RequestType")
	delete(f, "MerchantCode")
	delete(f, "PayChannel")
	delete(f, "PayChannelSubId")
	delete(f, "OrderId")
	delete(f, "BankAccountNumber")
	delete(f, "PlatformShortNumber")
	delete(f, "MidasSecretId")
	delete(f, "MidasAppId")
	delete(f, "MidasSignature")
	delete(f, "TransSequenceNumber")
	delete(f, "TransFee")
	delete(f, "ThirdPayChannel")
	delete(f, "ThirdPayChannelOrderId")
	delete(f, "OldFrontSequenceNumber")
	delete(f, "CurrencyAmount")
	delete(f, "CurrencyUnit")
	delete(f, "CurrencyType")
	delete(f, "MidasEnvironment")
	delete(f, "ReservedMessage")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RevokeRechargeByThirdPayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RevokeRechargeByThirdPayResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 请求类型
		RequestType *string `json:"RequestType,omitempty" name:"RequestType"`

		// 保留域
	// 注意：此字段可能返回 null，表示取不到有效值。
		ReservedMessage *string `json:"ReservedMessage,omitempty" name:"ReservedMessage"`

		// 银行流水号
		FrontSequenceNumber *string `json:"FrontSequenceNumber,omitempty" name:"FrontSequenceNumber"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RevokeRechargeByThirdPayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RevokeRechargeByThirdPayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SceneInfo struct {

	// 语言代码
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocaleCode *string `json:"LocaleCode,omitempty" name:"LocaleCode"`

	// 地区代码
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionCode *string `json:"RegionCode,omitempty" name:"RegionCode"`

	// 用户IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserClientIp *string `json:"UserClientIp,omitempty" name:"UserClientIp"`
}

type SyncContractDataRequest struct {
	*tchttp.BaseRequest

	// 聚鑫分配的支付主MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 用户ID，长度不小于5位，仅支持字母和数字的组合
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 签约使用的渠道
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// 业务签约合同协议号
	OutContractCode *string `json:"OutContractCode,omitempty" name:"OutContractCode"`

	// 签约状态，枚举值
	// CONTRACT_STATUS_INVALID=无效状态
	// CONTRACT_STATUS_SIGNED=已签约
	// CONTRACT_STATUS_TERMINATED=已解约
	// CONTRACT_STATUS_PENDING=签约进行中
	ContractStatus *string `json:"ContractStatus,omitempty" name:"ContractStatus"`

	// 签约同步信息
	ContractSyncInfo *ContractSyncInfo `json:"ContractSyncInfo,omitempty" name:"ContractSyncInfo"`

	// 按照聚鑫安全密钥计算的签名
	MidasSignature *string `json:"MidasSignature,omitempty" name:"MidasSignature"`

	// 聚鑫分配的安全ID
	MidasSecretId *string `json:"MidasSecretId,omitempty" name:"MidasSecretId"`

	// 聚鑫计费SubAppId，代表子商户
	SubAppId *string `json:"SubAppId,omitempty" name:"SubAppId"`

	// 用户类型，枚举值
	// USER_ID: 用户ID
	// ANONYMOUS: 匿名类型 USER_ID
	// 默认值为 USER_ID
	UserType *string `json:"UserType,omitempty" name:"UserType"`

	// 场景信息
	SceneInfo *SceneInfo `json:"SceneInfo,omitempty" name:"SceneInfo"`

	// 环境名:
	// release: 现网环境
	// sandbox: 沙箱环境
	// development: 开发环境
	// 缺省: release
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`
}

func (r *SyncContractDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncContractDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MidasAppId")
	delete(f, "UserId")
	delete(f, "Channel")
	delete(f, "OutContractCode")
	delete(f, "ContractStatus")
	delete(f, "ContractSyncInfo")
	delete(f, "MidasSignature")
	delete(f, "MidasSecretId")
	delete(f, "SubAppId")
	delete(f, "UserType")
	delete(f, "SceneInfo")
	delete(f, "MidasEnvironment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SyncContractDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SyncContractDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 请求处理信息
		Msg *string `json:"Msg,omitempty" name:"Msg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SyncContractDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncContractDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TerminateContractRequest struct {
	*tchttp.BaseRequest

	// 聚鑫分配的支付主MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 用户ID，长度不小于5位，仅支持字母和数字的组合
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 指定渠道：  wechat：微信支付  qqwallet：QQ钱包 
	//  bank：网银支付  只有一个渠道时需要指定
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// 枚举值：
	// CONTRACT_TERMINATION_MODE_BY_OUT_CONTRACT_CODE: 按OutContractCode+ContractSceneId解约
	// CONTRACT_TERMINATION_MODE_BY_CHANNEL_CONTRACT_CODE：按ChannelContractCode解约
	TerminateMode *string `json:"TerminateMode,omitempty" name:"TerminateMode"`

	// 聚鑫分配的安全ID
	MidasSecretId *string `json:"MidasSecretId,omitempty" name:"MidasSecretId"`

	// 按照聚鑫安全密钥计算的签名
	MidasSignature *string `json:"MidasSignature,omitempty" name:"MidasSignature"`

	// 聚鑫计费SubAppId，代表子商户
	SubAppId *string `json:"SubAppId,omitempty" name:"SubAppId"`

	// 业务签约合同协议号 当TerminateMode=CONTRACT_TERMINATION_MODE_BY_OUT_CONTRACT_CODE 时 必填
	OutContractCode *string `json:"OutContractCode,omitempty" name:"OutContractCode"`

	// 签约场景ID，当 TerminateMode=CONTRACT_TERMINATION_MODE_BY_OUT_CONTRACT_CODE 时 必填，在米大师侧托管后生成
	ContractSceneId *string `json:"ContractSceneId,omitempty" name:"ContractSceneId"`

	// 米大师生成的协议号 当 TerminateMode=CONTRACT_TERMINATION_MODE_BY_CHANNEL_CONTRACT_CODE 时 必填
	ChannelContractCode *string `json:"ChannelContractCode,omitempty" name:"ChannelContractCode"`

	// 第三方渠道合约数据，json字符串，与特定渠道有关
	ExternalContractData *string `json:"ExternalContractData,omitempty" name:"ExternalContractData"`

	// 终止合约原因
	TerminationReason *string `json:"TerminationReason,omitempty" name:"TerminationReason"`

	// 环境名:
	// release: 现网环境
	// sandbox: 沙箱环境
	// development: 开发环境
	// 缺省: release
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`

	// USER_ID: 用户ID
	// ANONYMOUS: 匿名类型 USER_ID
	// 默认值为 USER_ID
	UserType *string `json:"UserType,omitempty" name:"UserType"`

	// 签约方式
	ContractMethod *string `json:"ContractMethod,omitempty" name:"ContractMethod"`

	// 签约代扣穿透查询存量数据迁移模式
	MigrateMode *string `json:"MigrateMode,omitempty" name:"MigrateMode"`
}

func (r *TerminateContractRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateContractRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MidasAppId")
	delete(f, "UserId")
	delete(f, "Channel")
	delete(f, "TerminateMode")
	delete(f, "MidasSecretId")
	delete(f, "MidasSignature")
	delete(f, "SubAppId")
	delete(f, "OutContractCode")
	delete(f, "ContractSceneId")
	delete(f, "ChannelContractCode")
	delete(f, "ExternalContractData")
	delete(f, "TerminationReason")
	delete(f, "MidasEnvironment")
	delete(f, "UserType")
	delete(f, "ContractMethod")
	delete(f, "MigrateMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateContractRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type TerminateContractResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 解约数据
		ContractTerminateData *ResponseTerminateContract `json:"ContractTerminateData,omitempty" name:"ContractTerminateData"`

		// 请求处理信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Msg *string `json:"Msg,omitempty" name:"Msg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TerminateContractResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateContractResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TranItem struct {

	// STRING(50)，资金汇总账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	FundSummaryAcctNo *string `json:"FundSummaryAcctNo,omitempty" name:"FundSummaryAcctNo"`

	// STRING(50)，见证子账户的账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubAcctNo *string `json:"SubAcctNo,omitempty" name:"SubAcctNo"`

	// STRING(32)，交易网会员代码
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranNetMemberCode *string `json:"TranNetMemberCode,omitempty" name:"TranNetMemberCode"`

	// STRING(150)，会员名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemberName *string `json:"MemberName,omitempty" name:"MemberName"`

	// STRING(5)，会员证件类型（详情见“常见问题”）
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemberGlobalType *string `json:"MemberGlobalType,omitempty" name:"MemberGlobalType"`

	// STRING(32)，会员证件号码
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemberGlobalId *string `json:"MemberGlobalId,omitempty" name:"MemberGlobalId"`

	// STRING(50)，会员绑定账户的账号（提现的银行卡）
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemberAcctNo *string `json:"MemberAcctNo,omitempty" name:"MemberAcctNo"`

	// STRING(10)，会员绑定账户的本他行类型（1: 本行; 2: 他行）
	// 注意：此字段可能返回 null，表示取不到有效值。
	BankType *string `json:"BankType,omitempty" name:"BankType"`

	// STRING(150)，会员绑定账户的开户行名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AcctOpenBranchName *string `json:"AcctOpenBranchName,omitempty" name:"AcctOpenBranchName"`

	// STRING(20)，会员绑定账户的开户行的联行号
	// 注意：此字段可能返回 null，表示取不到有效值。
	CnapsBranchId *string `json:"CnapsBranchId,omitempty" name:"CnapsBranchId"`

	// STRING(20)，会员绑定账户的开户行的超级网银行号
	// 注意：此字段可能返回 null，表示取不到有效值。
	EiconBankBranchId *string `json:"EiconBankBranchId,omitempty" name:"EiconBankBranchId"`

	// STRING(30)，会员的手机号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`
}

type TransactionItem struct {

	// STRING(2)，记账标志（1: 转出; 2: 转入）
	// 注意：此字段可能返回 null，表示取不到有效值。
	BookingFlag *string `json:"BookingFlag,omitempty" name:"BookingFlag"`

	// STRING(32)，交易状态（0: 成功）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranStatus *string `json:"TranStatus,omitempty" name:"TranStatus"`

	// STRING(20)，交易金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranAmt *string `json:"TranAmt,omitempty" name:"TranAmt"`

	// STRING(8)，交易日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranDate *string `json:"TranDate,omitempty" name:"TranDate"`

	// STRING(20)，交易时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranTime *string `json:"TranTime,omitempty" name:"TranTime"`

	// STRING(52)，见证系统流水号
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrontSeqNo *string `json:"FrontSeqNo,omitempty" name:"FrontSeqNo"`

	// STRING(20)，记账类型（详情见“常见问题”）
	// 注意：此字段可能返回 null，表示取不到有效值。
	BookingType *string `json:"BookingType,omitempty" name:"BookingType"`

	// STRING(50)，转入见证子账户的帐号
	// 注意：此字段可能返回 null，表示取不到有效值。
	InSubAcctNo *string `json:"InSubAcctNo,omitempty" name:"InSubAcctNo"`

	// STRING(50)，转出见证子账户的帐号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutSubAcctNo *string `json:"OutSubAcctNo,omitempty" name:"OutSubAcctNo"`

	// STRING(300)，备注（返回交易订单号）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type TransferDetailRequest struct {

	// 商家明细单号。
	// 商户系统内部区分转账批次单下不同转账明细单的唯一标识，要求此参数只能由数字、大小写字母组成。
	// 示例值：x23zy545Bd5436
	MerchantDetailNo *string `json:"MerchantDetailNo,omitempty" name:"MerchantDetailNo"`

	// 转账金额。
	// 转账金额单位为分。
	// 示例值：200000
	TransferAmount *uint64 `json:"TransferAmount,omitempty" name:"TransferAmount"`

	// 转账备注。
	// 单条转账备注（微信用户会收到该备注）。UTF8编码，最多32字符。
	// 示例值：2020年4月报销
	TransferRemark *string `json:"TransferRemark,omitempty" name:"TransferRemark"`

	// 用户在直连商户下的唯一标识。
	// 示例值：o-MYE42l80oelYMDE34nYD456Xoy
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 收款用户姓名。
	// 收款方姓名。
	// 示例值：张三
	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

type TransferDetailResponse struct {

	// 商家明细单号。
	// 商户系统内部的商家明细单号
	// 示例值：plfk2020042013
	MerchantDetailNo *string `json:"MerchantDetailNo,omitempty" name:"MerchantDetailNo"`

	// 微信明细单号。
	// 微信区分明细单返回的唯一标识。
	// 示例值：1030000071100999991182020050700019480001
	DetailId *string `json:"DetailId,omitempty" name:"DetailId"`

	// 明细状态。
	// PROCESSING：转账中，正在处理，结果未明；
	// SUCCESS：转账成功；
	// FAIL：转账失败，需要确认失败原因以后，再决定是否重新发起地该笔明细的转账。
	// 示例值：SUCCESS
	DetailStatus *string `json:"DetailStatus,omitempty" name:"DetailStatus"`
}

type TransferItem struct {

	// STRING(10)，入账类型（02: 会员充值; 03: 资金挂账）
	// 注意：此字段可能返回 null，表示取不到有效值。
	InAcctType *string `json:"InAcctType,omitempty" name:"InAcctType"`

	// STRING(32)，交易网会员代码
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranNetMemberCode *string `json:"TranNetMemberCode,omitempty" name:"TranNetMemberCode"`

	// STRING(50)，见证子帐户的帐号
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubAcctNo *string `json:"SubAcctNo,omitempty" name:"SubAcctNo"`

	// STRING(20)，入金金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranAmt *string `json:"TranAmt,omitempty" name:"TranAmt"`

	// STRING(50)，入金账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	InAcctNo *string `json:"InAcctNo,omitempty" name:"InAcctNo"`

	// STRING(150)，入金账户名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InAcctName *string `json:"InAcctName,omitempty" name:"InAcctName"`

	// STRING(3)，币种
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ccy *string `json:"Ccy,omitempty" name:"Ccy"`

	// STRING(8)，会计日期（即银行主机记账日期）
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountingDate *string `json:"AccountingDate,omitempty" name:"AccountingDate"`

	// STRING(150)，银行名称（付款账户银行名称）
	// 注意：此字段可能返回 null，表示取不到有效值。
	BankName *string `json:"BankName,omitempty" name:"BankName"`

	// STRING(300)，转账备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// STRING(52)，见证系统流水号
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrontSeqNo *string `json:"FrontSeqNo,omitempty" name:"FrontSeqNo"`
}

type TransferSinglePayData struct {

	// 平台交易流水号，唯一
	TradeSerialNo *string `json:"TradeSerialNo,omitempty" name:"TradeSerialNo"`
}

type TransferSinglePayRequest struct {
	*tchttp.BaseRequest

	// 商户号
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 微信申请商户号的appid或者商户号绑定的appid
	// 支付宝、平安填入MerchantId
	MerchantAppId *string `json:"MerchantAppId,omitempty" name:"MerchantAppId"`

	// 1、 微信企业付款 
	// 2、 支付宝转账 
	// 3、 平安银企直联代发转账
	TransferType *int64 `json:"TransferType,omitempty" name:"TransferType"`

	// 订单流水号，唯一，不能包含特殊字符，长度最大限制64位，推荐使用字母，数字组合，"_","-"组合
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 转账金额，单位分
	TransferAmount *int64 `json:"TransferAmount,omitempty" name:"TransferAmount"`

	// 收款方标识。
	// 微信为open_id；
	// 支付宝为会员alipay_user_id;
	// 平安为收款方银行账号
	PayeeId *string `json:"PayeeId,omitempty" name:"PayeeId"`

	// 收款方姓名。支付宝可选；微信，平安模式下必传
	PayeeName *string `json:"PayeeName,omitempty" name:"PayeeName"`

	// 收款方附加信息，平安接入使用。需要以JSON格式提供以下字段：
	// PayeeBankName：收款人开户行名称
	//  CcyCode：货币类型（RMB-人民币）
	//  UnionFlag：行内跨行标志（1：行内转账，0：跨行转账）。
	PayeeExtends *string `json:"PayeeExtends,omitempty" name:"PayeeExtends"`

	// 请求预留字段，原样透传返回
	ReqReserved *string `json:"ReqReserved,omitempty" name:"ReqReserved"`

	// 业务备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 转账结果回调通知URL。若不填，则不进行回调。
	NotifyUrl *string `json:"NotifyUrl,omitempty" name:"NotifyUrl"`

	// 接入环境。沙箱环境填sandbox。
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *TransferSinglePayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TransferSinglePayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MerchantId")
	delete(f, "MerchantAppId")
	delete(f, "TransferType")
	delete(f, "OrderId")
	delete(f, "TransferAmount")
	delete(f, "PayeeId")
	delete(f, "PayeeName")
	delete(f, "PayeeExtends")
	delete(f, "ReqReserved")
	delete(f, "Remark")
	delete(f, "NotifyUrl")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TransferSinglePayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type TransferSinglePayResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 错误码。响应成功："SUCCESS"，其他为不成功
		ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

		// 响应消息
		ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

		// 返回结果
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *TransferSinglePayData `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TransferSinglePayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TransferSinglePayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnBindAcctRequest struct {
	*tchttp.BaseRequest

	// 聚鑫分配的支付主MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 聚鑫计费SubAppId，代表子商户
	SubAppId *string `json:"SubAppId,omitempty" name:"SubAppId"`

	// 用于提现
	// <敏感信息>加密详见<a href="https://cloud.tencent.com/document/product/1122/48979" target="_blank">《商户端接口敏感信息加密说明》</a>
	SettleAcctNo *string `json:"SettleAcctNo,omitempty" name:"SettleAcctNo"`

	// 聚鑫分配的安全ID
	MidasSecretId *string `json:"MidasSecretId,omitempty" name:"MidasSecretId"`

	// 按照聚鑫安全密钥计算的签名
	MidasSignature *string `json:"MidasSignature,omitempty" name:"MidasSignature"`

	// 敏感信息加密类型:
	// RSA: rsa非对称加密，使用RSA-PKCS1-v1_5
	// AES: aes对称加密，使用AES256-CBC-PCKS7padding
	// 缺省: RSA
	EncryptType *string `json:"EncryptType,omitempty" name:"EncryptType"`

	// 环境名:
	// release: 现网环境
	// sandbox: 沙箱环境
	// development: 开发环境
	// 缺省: release
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`
}

func (r *UnBindAcctRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnBindAcctRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MidasAppId")
	delete(f, "SubAppId")
	delete(f, "SettleAcctNo")
	delete(f, "MidasSecretId")
	delete(f, "MidasSignature")
	delete(f, "EncryptType")
	delete(f, "MidasEnvironment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnBindAcctRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UnBindAcctResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnBindAcctResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnBindAcctResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnbindRelateAcctRequest struct {
	*tchttp.BaseRequest

	// String(22)，商户号（签约客户号）
	MrchCode *string `json:"MrchCode,omitempty" name:"MrchCode"`

	// STRING(2)，功能标志（1: 解绑）
	FunctionFlag *string `json:"FunctionFlag,omitempty" name:"FunctionFlag"`

	// STRING(32)，交易网会员代码（若需要把一个待绑定账户关联到两个会员名下，此字段可上送两个会员的交易网代码，并且须用“|::|”(右侧)进行分隔）
	TranNetMemberCode *string `json:"TranNetMemberCode,omitempty" name:"TranNetMemberCode"`

	// STRING(50)，待解绑的提现账户的账号（提现账号）
	MemberAcctNo *string `json:"MemberAcctNo,omitempty" name:"MemberAcctNo"`

	// STRING(1027)，保留域
	ReservedMsg *string `json:"ReservedMsg,omitempty" name:"ReservedMsg"`

	// STRING(12)，接入环境，默认接入沙箱环境。接入正式环境填"prod"
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *UnbindRelateAcctRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindRelateAcctRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MrchCode")
	delete(f, "FunctionFlag")
	delete(f, "TranNetMemberCode")
	delete(f, "MemberAcctNo")
	delete(f, "ReservedMsg")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindRelateAcctRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UnbindRelateAcctResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// String(20)，返回码
		TxnReturnCode *string `json:"TxnReturnCode,omitempty" name:"TxnReturnCode"`

		// String(100)，返回信息
		TxnReturnMsg *string `json:"TxnReturnMsg,omitempty" name:"TxnReturnMsg"`

		// String(22)，交易流水号
		CnsmrSeqNo *string `json:"CnsmrSeqNo,omitempty" name:"CnsmrSeqNo"`

		// STRING(52)，见证系统流水号（即电商见证宝系统生成的流水号，可关联具体一笔请求）
	// 注意：此字段可能返回 null，表示取不到有效值。
		FrontSeqNo *string `json:"FrontSeqNo,omitempty" name:"FrontSeqNo"`

		// STRING(1027)，保留域
	// 注意：此字段可能返回 null，表示取不到有效值。
		ReservedMsg *string `json:"ReservedMsg,omitempty" name:"ReservedMsg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnbindRelateAcctResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindRelateAcctResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnifiedOrderInSubOrderList struct {

	// 子订单结算应收金额，单位： 分
	SubMchIncome *int64 `json:"SubMchIncome,omitempty" name:"SubMchIncome"`

	// 子订单平台应收金额，单位：分
	PlatformIncome *int64 `json:"PlatformIncome,omitempty" name:"PlatformIncome"`

	// 子订单商品详情
	ProductDetail *string `json:"ProductDetail,omitempty" name:"ProductDetail"`

	// 子订单商品名称
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 聚鑫计费SubAppId，代表子商户
	SubAppId *string `json:"SubAppId,omitempty" name:"SubAppId"`

	// 子订单号
	SubOutTradeNo *string `json:"SubOutTradeNo,omitempty" name:"SubOutTradeNo"`

	// 子订单支付金额
	Amt *int64 `json:"Amt,omitempty" name:"Amt"`

	// 发货标识，由业务在调用聚鑫下单接口的 时候下发
	Metadata *string `json:"Metadata,omitempty" name:"Metadata"`

	// 子订单原始金额
	OriginalAmt *int64 `json:"OriginalAmt,omitempty" name:"OriginalAmt"`
}

type UnifiedOrderRequest struct {
	*tchttp.BaseRequest

	// ISO 货币代码，CNY
	CurrencyType *string `json:"CurrencyType,omitempty" name:"CurrencyType"`

	// 聚鑫分配的支付主MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 支付订单号，仅支持数字、字母、下划线（_）、横杠字符（-）、点（.）的组合
	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`

	// 商品详情，需要URL编码
	ProductDetail *string `json:"ProductDetail,omitempty" name:"ProductDetail"`

	// 商品ID，仅支持数字、字母、下划线（_）、横杠字符（-）、点（.）的组合
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 商品名称，需要URL编码
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 支付金额，单位： 分
	TotalAmt *int64 `json:"TotalAmt,omitempty" name:"TotalAmt"`

	// 用户ID，长度不小于5位，仅支持字母和数字的组合
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 银行真实渠道.如:bank_pingan
	RealChannel *string `json:"RealChannel,omitempty" name:"RealChannel"`

	// 原始金额
	OriginalAmt *int64 `json:"OriginalAmt,omitempty" name:"OriginalAmt"`

	// 聚鑫分配的安全ID
	MidasSecretId *string `json:"MidasSecretId,omitempty" name:"MidasSecretId"`

	// 按照聚鑫安全密钥计算的签名
	MidasSignature *string `json:"MidasSignature,omitempty" name:"MidasSignature"`

	// Web端回调地址
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 指定支付渠道：  wechat：微信支付  qqwallet：QQ钱包 
	//  bank：网银支付  只有一个渠道时需要指定
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// 透传字段，支付成功回调透传给应用，用于业务透传自定义内容
	Metadata *string `json:"Metadata,omitempty" name:"Metadata"`

	// 购买数量，不传默认为1
	Quantity *int64 `json:"Quantity,omitempty" name:"Quantity"`

	// 聚鑫计费SubAppId，代表子商户
	SubAppId *string `json:"SubAppId,omitempty" name:"SubAppId"`

	// 子订单信息列表，格式：子订单号、子应用ID、金额。 压缩后最长不可超过65535字节(去除空格，换行，制表符等无意义字符)
	// 注：接入银行或其他支付渠道服务商模式下，必传
	SubOrderList []*UnifiedOrderInSubOrderList `json:"SubOrderList,omitempty" name:"SubOrderList"`

	// 结算应收金额，单位：分
	TotalMchIncome *int64 `json:"TotalMchIncome,omitempty" name:"TotalMchIncome"`

	// 平台应收金额，单位：分
	TotalPlatformIncome *int64 `json:"TotalPlatformIncome,omitempty" name:"TotalPlatformIncome"`

	// 微信公众号/小程序支付时为必选，需要传微信下的openid
	WxOpenId *string `json:"WxOpenId,omitempty" name:"WxOpenId"`

	// 在服务商模式下，微信公众号/小程序支付时wx_sub_openid和wx_openid二选一
	WxSubOpenId *string `json:"WxSubOpenId,omitempty" name:"WxSubOpenId"`

	// 环境名:
	// release: 现网环境
	// sandbox: 沙箱环境
	// development: 开发环境
	// 缺省: release
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`

	// 微信商户应用ID
	WxAppId *string `json:"WxAppId,omitempty" name:"WxAppId"`

	// 微信商户子应用ID
	WxSubAppId *string `json:"WxSubAppId,omitempty" name:"WxSubAppId"`

	// 支付通知地址
	PaymentNotifyUrl *string `json:"PaymentNotifyUrl,omitempty" name:"PaymentNotifyUrl"`
}

func (r *UnifiedOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnifiedOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CurrencyType")
	delete(f, "MidasAppId")
	delete(f, "OutTradeNo")
	delete(f, "ProductDetail")
	delete(f, "ProductId")
	delete(f, "ProductName")
	delete(f, "TotalAmt")
	delete(f, "UserId")
	delete(f, "RealChannel")
	delete(f, "OriginalAmt")
	delete(f, "MidasSecretId")
	delete(f, "MidasSignature")
	delete(f, "CallbackUrl")
	delete(f, "Channel")
	delete(f, "Metadata")
	delete(f, "Quantity")
	delete(f, "SubAppId")
	delete(f, "SubOrderList")
	delete(f, "TotalMchIncome")
	delete(f, "TotalPlatformIncome")
	delete(f, "WxOpenId")
	delete(f, "WxSubOpenId")
	delete(f, "MidasEnvironment")
	delete(f, "WxAppId")
	delete(f, "WxSubAppId")
	delete(f, "PaymentNotifyUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnifiedOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UnifiedOrderResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 支付金额，单位： 分
		TotalAmt *int64 `json:"TotalAmt,omitempty" name:"TotalAmt"`

		// 应用支付订单号
		OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`

		// 支付参数透传给聚鑫SDK（原文透传给SDK即可，不需要解码）
		PayInfo *string `json:"PayInfo,omitempty" name:"PayInfo"`

		// 聚鑫的交易订单
		TransactionId *string `json:"TransactionId,omitempty" name:"TransactionId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnifiedOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnifiedOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UploadExternalAnchorInfoRequest struct {
	*tchttp.BaseRequest

	// 主播Id
	AnchorId *string `json:"AnchorId,omitempty" name:"AnchorId"`

	// 身份证正面图片下载链接
	IdCardFront *string `json:"IdCardFront,omitempty" name:"IdCardFront"`

	// 身份证反面图片下载链接
	IdCardReverse *string `json:"IdCardReverse,omitempty" name:"IdCardReverse"`
}

func (r *UploadExternalAnchorInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadExternalAnchorInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AnchorId")
	delete(f, "IdCardFront")
	delete(f, "IdCardReverse")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadExternalAnchorInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UploadExternalAnchorInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 错误码。响应成功："SUCCESS"，其他为不成功。
		ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

		// 响应消息。
		ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

		// 该字段为null。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UploadExternalAnchorInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadExternalAnchorInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UploadTaxListRequest struct {
	*tchttp.BaseRequest

	// 平台渠道
	Channel *int64 `json:"Channel,omitempty" name:"Channel"`

	// 起始月份，YYYY-MM
	BeginMonth *string `json:"BeginMonth,omitempty" name:"BeginMonth"`

	// 结束月份。如果只上传一个月，结束月份等于起始月份
	EndMonth *string `json:"EndMonth,omitempty" name:"EndMonth"`

	// 完税列表下载地址
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`
}

func (r *UploadTaxListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadTaxListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Channel")
	delete(f, "BeginMonth")
	delete(f, "EndMonth")
	delete(f, "FileUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadTaxListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UploadTaxListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 完税ID
		TaxId *string `json:"TaxId,omitempty" name:"TaxId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UploadTaxListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadTaxListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UploadTaxPaymentRequest struct {
	*tchttp.BaseRequest

	// 平台渠道
	Channel *int64 `json:"Channel,omitempty" name:"Channel"`

	// 完税ID
	TaxId *string `json:"TaxId,omitempty" name:"TaxId"`

	// 完税列表下载地址
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`
}

func (r *UploadTaxPaymentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadTaxPaymentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Channel")
	delete(f, "TaxId")
	delete(f, "FileUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadTaxPaymentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UploadTaxPaymentResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UploadTaxPaymentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadTaxPaymentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WithdrawBill struct {

	// 业务提现订单号
	WithdrawOrderId *string `json:"WithdrawOrderId,omitempty" name:"WithdrawOrderId"`

	// 提现日期
	Date *string `json:"Date,omitempty" name:"Date"`

	// 提现金额，单位： 分
	PayAmt *string `json:"PayAmt,omitempty" name:"PayAmt"`

	// 聚鑫分配转入账户appid
	InSubAppId *string `json:"InSubAppId,omitempty" name:"InSubAppId"`

	// 聚鑫分配转出账户appid
	OutSubAppId *string `json:"OutSubAppId,omitempty" name:"OutSubAppId"`

	// ISO货币代码
	CurrencyType *string `json:"CurrencyType,omitempty" name:"CurrencyType"`

	// 透传字段
	MetaData *string `json:"MetaData,omitempty" name:"MetaData"`

	// 扩展字段
	ExtendFieldData *string `json:"ExtendFieldData,omitempty" name:"ExtendFieldData"`
}

type WithdrawCashMembershipRequest struct {
	*tchttp.BaseRequest

	// String(22)，商户号（签约客户号）
	MrchCode *string `json:"MrchCode,omitempty" name:"MrchCode"`

	// STRING(150)，交易网名称（市场名称）
	TranWebName *string `json:"TranWebName,omitempty" name:"TranWebName"`

	// STRING(5)，会员证件类型（详情见“常见问题”）
	MemberGlobalType *string `json:"MemberGlobalType,omitempty" name:"MemberGlobalType"`

	// STRING(32)，会员证件号码
	MemberGlobalId *string `json:"MemberGlobalId,omitempty" name:"MemberGlobalId"`

	// STRING(32)，交易网会员代码
	TranNetMemberCode *string `json:"TranNetMemberCode,omitempty" name:"TranNetMemberCode"`

	// STRING(150)，会员名称
	MemberName *string `json:"MemberName,omitempty" name:"MemberName"`

	// STRING(50)，提现账号（银行卡）
	TakeCashAcctNo *string `json:"TakeCashAcctNo,omitempty" name:"TakeCashAcctNo"`

	// STRING(150)，出金账户名称（银行卡户名）
	OutAmtAcctName *string `json:"OutAmtAcctName,omitempty" name:"OutAmtAcctName"`

	// STRING(3)，币种（默认为RMB）
	Ccy *string `json:"Ccy,omitempty" name:"Ccy"`

	// STRING(20)，可提现金额
	CashAmt *string `json:"CashAmt,omitempty" name:"CashAmt"`

	// STRING(300)，备注（建议可送订单号，可在对账文件的备注字段获取到）
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// STRING(1027)，保留域
	ReservedMsg *string `json:"ReservedMsg,omitempty" name:"ReservedMsg"`

	// STRING(300)，网银签名
	WebSign *string `json:"WebSign,omitempty" name:"WebSign"`

	// STRING(12)，接入环境，默认接入沙箱环境。接入正式环境填"prod"
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *WithdrawCashMembershipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *WithdrawCashMembershipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MrchCode")
	delete(f, "TranWebName")
	delete(f, "MemberGlobalType")
	delete(f, "MemberGlobalId")
	delete(f, "TranNetMemberCode")
	delete(f, "MemberName")
	delete(f, "TakeCashAcctNo")
	delete(f, "OutAmtAcctName")
	delete(f, "Ccy")
	delete(f, "CashAmt")
	delete(f, "Remark")
	delete(f, "ReservedMsg")
	delete(f, "WebSign")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "WithdrawCashMembershipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type WithdrawCashMembershipResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// String(20)，返回码
		TxnReturnCode *string `json:"TxnReturnCode,omitempty" name:"TxnReturnCode"`

		// String(100)，返回信息
		TxnReturnMsg *string `json:"TxnReturnMsg,omitempty" name:"TxnReturnMsg"`

		// String(22)，交易流水号
		CnsmrSeqNo *string `json:"CnsmrSeqNo,omitempty" name:"CnsmrSeqNo"`

		// STRING(52)，见证系统流水号
	// 注意：此字段可能返回 null，表示取不到有效值。
		FrontSeqNo *string `json:"FrontSeqNo,omitempty" name:"FrontSeqNo"`

		// STRING(20)，转账手续费（固定返回0.00）
	// 注意：此字段可能返回 null，表示取不到有效值。
		TransferFee *string `json:"TransferFee,omitempty" name:"TransferFee"`

		// STRING(1027)，保留域
	// 注意：此字段可能返回 null，表示取不到有效值。
		ReservedMsg *string `json:"ReservedMsg,omitempty" name:"ReservedMsg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WithdrawCashMembershipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *WithdrawCashMembershipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WithdrawItem struct {

	// STRING(2)，记账标志（01: 提现; 02: 清分 ）
	// 注意：此字段可能返回 null，表示取不到有效值。
	BookingFlag *string `json:"BookingFlag,omitempty" name:"BookingFlag"`

	// STRING(32)，交易状态（0: 成功）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranStatus *string `json:"TranStatus,omitempty" name:"TranStatus"`

	// STRING(200)，记账说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	BookingMsg *string `json:"BookingMsg,omitempty" name:"BookingMsg"`

	// STRING(32)，交易网会员代码
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranNetMemberCode *string `json:"TranNetMemberCode,omitempty" name:"TranNetMemberCode"`

	// STRING(50)，见证子帐户的帐号
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubAcctNo *string `json:"SubAcctNo,omitempty" name:"SubAcctNo"`

	// STRING(150)，见证子账户的名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubAcctName *string `json:"SubAcctName,omitempty" name:"SubAcctName"`

	// STRING(20)，交易金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranAmt *string `json:"TranAmt,omitempty" name:"TranAmt"`

	// STRING(20)，手续费
	// 注意：此字段可能返回 null，表示取不到有效值。
	Commission *string `json:"Commission,omitempty" name:"Commission"`

	// STRING(8)，交易日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranDate *string `json:"TranDate,omitempty" name:"TranDate"`

	// STRING(20)，交易时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranTime *string `json:"TranTime,omitempty" name:"TranTime"`

	// STRING(52)，见证系统流水号
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrontSeqNo *string `json:"FrontSeqNo,omitempty" name:"FrontSeqNo"`

	// STRING(300)，备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}
