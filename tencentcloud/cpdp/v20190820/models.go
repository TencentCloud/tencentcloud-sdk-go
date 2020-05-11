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

func (r *ApplyApplicationMaterialRequest) FromJsonString(s string) error {
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

	// 收款人类型
	PayeeType *string `json:"PayeeType,omitempty" name:"PayeeType"`

	// 收款人账号
	PayeeAccount *string `json:"PayeeAccount,omitempty" name:"PayeeAccount"`

	// 源币种金额
	SourceAmount *float64 `json:"SourceAmount,omitempty" name:"SourceAmount"`

	// 目的金额
	TargetAmount *float64 `json:"TargetAmount,omitempty" name:"TargetAmount"`

	// 收款人姓名
	PayeeName *string `json:"PayeeName,omitempty" name:"PayeeName"`

	// 收款人地址
	PayeeAddress *string `json:"PayeeAddress,omitempty" name:"PayeeAddress"`

	// 收款人银行账号类型
	PayeeBankAccountType *string `json:"PayeeBankAccountType,omitempty" name:"PayeeBankAccountType"`

	// 收款人国家或地区编码
	PayeeCountryCode *string `json:"PayeeCountryCode,omitempty" name:"PayeeCountryCode"`

	// 收款人开户银行名称
	PayeeBankName *string `json:"PayeeBankName,omitempty" name:"PayeeBankName"`

	// 收款人开户银行地址
	PayeeBankAddress *string `json:"PayeeBankAddress,omitempty" name:"PayeeBankAddress"`

	// 收款人开户银行所在国家或地区编码
	PayeeBankDistrict *string `json:"PayeeBankDistrict,omitempty" name:"PayeeBankDistrict"`

	// 收款银行SwiftCode
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

func (r *ApplyOutwardOrderRequest) FromJsonString(s string) error {
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

	// 付款人联系电话 (PayerType=CORPORATE 必填)
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

func (r *ApplyPayerInfoRequest) FromJsonString(s string) error {
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
}

func (r *ApplyReWithdrawalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ApplyReWithdrawalRequest) FromJsonString(s string) error {
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

	// 付款人姓名
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

func (r *ApplyTradeRequest) FromJsonString(s string) error {
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
	// <敏感信息>加密详见《商户端接口敏感信息加密说明》
	SettleAcctNo *string `json:"SettleAcctNo,omitempty" name:"SettleAcctNo"`

	// 结算账户户名
	// <敏感信息>加密详见《商户端接口敏感信息加密说明》
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
	// <敏感信息>加密详见《商户端接口敏感信息加密说明》
	IdCode *string `json:"IdCode,omitempty" name:"IdCode"`

	// 聚鑫分配的安全ID
	MidasSecretId *string `json:"MidasSecretId,omitempty" name:"MidasSecretId"`

	// 按照聚鑫安全密钥计算的签名
	MidasSignature *string `json:"MidasSignature,omitempty" name:"MidasSignature"`
}

func (r *ApplyWithdrawalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ApplyWithdrawalRequest) FromJsonString(s string) error {
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

type BindAcctRequest struct {
	*tchttp.BaseRequest

	// 聚鑫分配的支付主MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 聚鑫计费SubAppId，代表子商户
	SubAppId *string `json:"SubAppId,omitempty" name:"SubAppId"`

	// 1 – 小额转账验证
	// 2 – 短信验证
	// 每个结算账户每天只能使用一次小额转账验证
	BindType *int64 `json:"BindType,omitempty" name:"BindType"`

	// 用于提现
	// <敏感信息>加密详见《商户端接口敏感信息加密说明》
	SettleAcctNo *string `json:"SettleAcctNo,omitempty" name:"SettleAcctNo"`

	// 结算账户户名
	// <敏感信息>加密详见《商户端接口敏感信息加密说明》
	SettleAcctName *string `json:"SettleAcctName,omitempty" name:"SettleAcctName"`

	// 1 – 本行账户
	// 2 – 他行账户
	SettleAcctType *int64 `json:"SettleAcctType,omitempty" name:"SettleAcctType"`

	// 证件类型，见《证件类型》表
	IdType *string `json:"IdType,omitempty" name:"IdType"`

	// 证件号码
	// <敏感信息>加密详见《商户端接口敏感信息加密说明》
	IdCode *string `json:"IdCode,omitempty" name:"IdCode"`

	// 开户行名称
	AcctBranchName *string `json:"AcctBranchName,omitempty" name:"AcctBranchName"`

	// 聚鑫分配的安全ID
	MidasSecretId *string `json:"MidasSecretId,omitempty" name:"MidasSecretId"`

	// 按照聚鑫安全密钥计算的签名
	MidasSignature *string `json:"MidasSignature,omitempty" name:"MidasSignature"`

	// 用于短信验证
	// BindType==2时必填
	// <敏感信息>加密详见《商户端接口敏感信息加密说明》
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 大小额行号，超级网银行号和大小额行号
	// 二选一
	CnapsBranchId *string `json:"CnapsBranchId,omitempty" name:"CnapsBranchId"`

	// 超级网银行号，超级网银行号和大小额行号
	// 二选一
	EiconBankBranchId *string `json:"EiconBankBranchId,omitempty" name:"EiconBankBranchId"`
}

func (r *BindAcctRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindAcctRequest) FromJsonString(s string) error {
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
}

func (r *BindRelateAccReUnionPayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindRelateAccReUnionPayRequest) FromJsonString(s string) error {
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
}

func (r *BindRelateAcctSmallAmountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindRelateAcctSmallAmountRequest) FromJsonString(s string) error {
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
}

func (r *BindRelateAcctUnionPayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BindRelateAcctUnionPayRequest) FromJsonString(s string) error {
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

func (r *BindRelateAcctUnionPayResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckAcctRequest struct {
	*tchttp.BaseRequest

	// 聚鑫分配的支付主MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 聚鑫计费SubAppId，代表子商户
	SubAppId *string `json:"SubAppId,omitempty" name:"SubAppId"`

	// 1：小额鉴权
	// 2：短信校验鉴权
	BindType *int64 `json:"BindType,omitempty" name:"BindType"`

	// 结算账户账号
	// <敏感信息>加密详见《商户端接口敏感信息加密说明》
	SettleAcctNo *string `json:"SettleAcctNo,omitempty" name:"SettleAcctNo"`

	// 聚鑫分配的安全ID
	MidasSecretId *string `json:"MidasSecretId,omitempty" name:"MidasSecretId"`

	// 按照聚鑫安全密钥计算的签名
	MidasSignature *string `json:"MidasSignature,omitempty" name:"MidasSignature"`

	// 短信验证码
	// BindType==2必填
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
}

func (r *CheckAcctRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckAcctRequest) FromJsonString(s string) error {
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
}

func (r *CheckAmountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckAmountRequest) FromJsonString(s string) error {
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
}

func (r *CloseOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CloseOrderRequest) FromJsonString(s string) error {
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

func (r *CloseOrderResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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
	// <敏感信息>加密详见《商户端接口敏感信息加密说明》
	Contact *string `json:"Contact,omitempty" name:"Contact"`

	// 联系人手机号
	// <敏感信息>加密详见《商户端接口敏感信息加密说明》
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 邮箱 
	// <敏感信息>加密详见《商户端接口敏感信息加密说明》
	Email *string `json:"Email,omitempty" name:"Email"`

	// 聚鑫分配的安全ID
	MidasSecretId *string `json:"MidasSecretId,omitempty" name:"MidasSecretId"`

	// 按照聚鑫安全密钥计算的签名
	MidasSignature *string `json:"MidasSignature,omitempty" name:"MidasSignature"`

	// 子商户类型：
	// 个人: personal
	// 企业：enterprise
	// 缺省： enterprise
	SubMchType *string `json:"SubMchType,omitempty" name:"SubMchType"`

	// 不填则默认子商户名称
	ShortName *string `json:"ShortName,omitempty" name:"ShortName"`

	// 子商户会员类型：
	// general:普通子账户
	// merchant:商户子账户
	// 缺省： general
	SubMerchantMemberType *string `json:"SubMerchantMemberType,omitempty" name:"SubMerchantMemberType"`
}

func (r *CreateAcctRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAcctRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAcctResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 聚鑫计费SubAppId，代表子商户
		SubAppId *string `json:"SubAppId,omitempty" name:"SubAppId"`

		// 平安银行生成的子商户账户
		SubAcctNo *string `json:"SubAcctNo,omitempty" name:"SubAcctNo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAcctResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

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
	AgentTaxPaymentInfos []*AgentTaxPayment `json:"AgentTaxPaymentInfos,omitempty" name:"AgentTaxPaymentInfos" list`

	// 接入环境。沙箱环境填sandbox
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *CreateAgentTaxPaymentInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAgentTaxPaymentInfosRequest) FromJsonString(s string) error {
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
}

func (r *CreateCustAcctIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateCustAcctIdRequest) FromJsonString(s string) error {
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

func (r *CreateCustAcctIdResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateInvoiceItem struct {

	// 商品名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 税收商品编码
	TaxCode *string `json:"TaxCode,omitempty" name:"TaxCode"`

	// 不含税商品总价（商品含税价总额/（1+税率））。单位为分
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

	// 不含税商品单价
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

	// 开票平台ID。0：高灯
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

	// 不含税总金额（单位为分）
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
	Items []*CreateInvoiceItem `json:"Items,omitempty" name:"Items" list`

	// 接入环境。沙箱环境填sandbox。
	Profile *string `json:"Profile,omitempty" name:"Profile"`

	// 撤销部分商品。0-不撤销，1-撤销
	UndoPart *int64 `json:"UndoPart,omitempty" name:"UndoPart"`
}

func (r *CreateInvoiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateInvoiceRequest) FromJsonString(s string) error {
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

func (r *CreateMerchantRequest) FromJsonString(s string) error {
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

type CreateRedInvoiceItem struct {

	// 订单号
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 发票结果回传地址
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 业务开票号
	OrderSn *string `json:"OrderSn,omitempty" name:"OrderSn"`

	// 红字信息表编码
	RedSerialNo *string `json:"RedSerialNo,omitempty" name:"RedSerialNo"`
}

type CreateRedInvoiceRequest struct {
	*tchttp.BaseRequest

	// 开票平台ID
	InvoicePlatformId *int64 `json:"InvoicePlatformId,omitempty" name:"InvoicePlatformId"`

	// 红冲明细
	Invoices []*CreateRedInvoiceItem `json:"Invoices,omitempty" name:"Invoices" list`

	// 接入环境。沙箱环境填 sandbox。
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *CreateRedInvoiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRedInvoiceRequest) FromJsonString(s string) error {
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
	Data []*CreateRedInvoiceResultData `json:"Data,omitempty" name:"Data" list`
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

func (r *DeleteAgentTaxPaymentInfoRequest) FromJsonString(s string) error {
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

func (r *DeleteAgentTaxPaymentInfosRequest) FromJsonString(s string) error {
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

func (r *DeleteAgentTaxPaymentInfosResponse) FromJsonString(s string) error {
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
}

func (r *DownloadBillRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DownloadBillRequest) FromJsonString(s string) error {
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

func (r *DownloadBillResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *ModifyAgentTaxPaymentInfoRequest) FromJsonString(s string) error {
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

func (r *ModifyAgentTaxPaymentInfoResponse) FromJsonString(s string) error {
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
}

func (r *ModifyMntMbrBindRelateAcctBankCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMntMbrBindRelateAcctBankCodeRequest) FromJsonString(s string) error {
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

func (r *ModifyMntMbrBindRelateAcctBankCodeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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
}

func (r *QueryAcctBindingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryAcctBindingRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryAcctBindingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总行数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 银行卡信息列表
		BankCardItems []*BankCardItem `json:"BankCardItems,omitempty" name:"BankCardItems" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryAcctBindingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

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
}

func (r *QueryAcctInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryAcctInfoListRequest) FromJsonString(s string) error {
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
		QueryAcctItems []*QueryAcctItem `json:"QueryAcctItems,omitempty" name:"QueryAcctItems" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryAcctInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

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
}

func (r *QueryAcctInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryAcctInfoRequest) FromJsonString(s string) error {
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

func (r *QueryAgentTaxPaymentBatchRequest) FromJsonString(s string) error {
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

func (r *QueryAgentTaxPaymentBatchResponse) FromJsonString(s string) error {
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

func (r *QueryApplicationMaterialRequest) FromJsonString(s string) error {
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
}

func (r *QueryBalanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryBalanceRequest) FromJsonString(s string) error {
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
		QueryItems []*QueryItem `json:"QueryItems,omitempty" name:"QueryItems" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryBalanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

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
}

func (r *QueryBankClearRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryBankClearRequest) FromJsonString(s string) error {
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
		TranItemArray []*ClearItem `json:"TranItemArray,omitempty" name:"TranItemArray" list`

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
}

func (r *QueryBankTransactionDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryBankTransactionDetailsRequest) FromJsonString(s string) error {
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
		TranItemArray []*TransactionItem `json:"TranItemArray,omitempty" name:"TranItemArray" list`

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
}

func (r *QueryBankWithdrawCashDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryBankWithdrawCashDetailsRequest) FromJsonString(s string) error {
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
		TranItemArray []*WithdrawItem `json:"TranItemArray,omitempty" name:"TranItemArray" list`

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

func (r *QueryBankWithdrawCashDetailsResponse) FromJsonString(s string) error {
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
}

func (r *QueryCommonTransferRechargeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryCommonTransferRechargeRequest) FromJsonString(s string) error {
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
		TranItemArray []*TransferItem `json:"TranItemArray,omitempty" name:"TranItemArray" list`

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

func (r *QueryCommonTransferRechargeResponse) FromJsonString(s string) error {
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
}

func (r *QueryCustAcctIdBalanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryCustAcctIdBalanceRequest) FromJsonString(s string) error {
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
		AcctArray []*Acct `json:"AcctArray,omitempty" name:"AcctArray" list`

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

func (r *QueryExchangeRateRequest) FromJsonString(s string) error {
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
	Data []*QueryExchangerateData `json:"Data,omitempty" name:"Data" list`
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
}

func (r *QueryInvoiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryInvoiceRequest) FromJsonString(s string) error {
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
}

func (r *QueryMemberBindRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryMemberBindRequest) FromJsonString(s string) error {
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
		TranItemArray []*TranItem `json:"TranItemArray,omitempty" name:"TranItemArray" list`

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
}

func (r *QueryMemberTransactionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryMemberTransactionRequest) FromJsonString(s string) error {
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

func (r *QueryMerchantBalanceRequest) FromJsonString(s string) error {
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

func (r *QueryMerchantBalanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryMerchantBalanceResult struct {

	// 错误码
	Code *string `json:"Code,omitempty" name:"Code"`

	// 对接账户余额查询数据
	Data *QueryMerchantBalanceData `json:"Data,omitempty" name:"Data"`
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
	SubOrderList []*QueryOrderOutSubOrderList `json:"SubOrderList,omitempty" name:"SubOrderList" list`

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
}

func (r *QueryOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryOrderRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryOrderResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回订单数
		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 查询结果的订单列表
		OrderList []*QueryOrderOutOrderList `json:"OrderList,omitempty" name:"OrderList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

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

func (r *QueryOutwardOrderRequest) FromJsonString(s string) error {
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

func (r *QueryPayerInfoRequest) FromJsonString(s string) error {
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
}

func (r *QueryReconciliationDocumentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryReconciliationDocumentRequest) FromJsonString(s string) error {
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
		TranItemArray []*FileItem `json:"TranItemArray,omitempty" name:"TranItemArray" list`

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
}

func (r *QueryRefundRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryRefundRequest) FromJsonString(s string) error {
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

func (r *QueryRefundResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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
}

func (r *QuerySingleTransactionStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QuerySingleTransactionStatusRequest) FromJsonString(s string) error {
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
}

func (r *QuerySmallAmountTransferRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QuerySmallAmountTransferRequest) FromJsonString(s string) error {
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

func (r *QueryTradeRequest) FromJsonString(s string) error {
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

func (r *QueryTradeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryTradeResult struct {

	// 贸易材料明细查询数据
	Data *QueryTradeData `json:"Data,omitempty" name:"Data"`

	// 错误码
	Code *string `json:"Code,omitempty" name:"Code"`
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
}

func (r *RechargeMemberThirdPayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RechargeMemberThirdPayRequest) FromJsonString(s string) error {
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

func (r *RechargeMemberThirdPayResponse) FromJsonString(s string) error {
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
	SubOrderRefundList []*RefundOutSubOrderRefundList `json:"SubOrderRefundList,omitempty" name:"SubOrderRefundList" list`
}

func (r *RefundRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RefundRequest) FromJsonString(s string) error {
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

func (r *RefundResponse) FromJsonString(s string) error {
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
}

func (r *RegisterBillSupportWithdrawRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RegisterBillSupportWithdrawRequest) FromJsonString(s string) error {
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

func (r *RegisterBillSupportWithdrawResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RevRegisterBillSupportWithdrawRequest struct {
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
}

func (r *RevRegisterBillSupportWithdrawRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RevRegisterBillSupportWithdrawRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RevRegisterBillSupportWithdrawResponse struct {
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

func (r *RevRegisterBillSupportWithdrawResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RevRegisterBillSupportWithdrawResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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
}

func (r *RevResigterBillSupportWithdrawRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RevResigterBillSupportWithdrawRequest) FromJsonString(s string) error {
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
}

func (r *ReviseMbrPropertyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReviseMbrPropertyRequest) FromJsonString(s string) error {
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
}

func (r *RevokeMemberRechargeThirdPayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RevokeMemberRechargeThirdPayRequest) FromJsonString(s string) error {
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

func (r *RevokeMemberRechargeThirdPayResponse) FromJsonString(s string) error {
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

type UnBindAcctRequest struct {
	*tchttp.BaseRequest

	// 聚鑫分配的支付主MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 聚鑫计费SubAppId，代表子商户
	SubAppId *string `json:"SubAppId,omitempty" name:"SubAppId"`

	// 用于提现
	// <敏感信息>加密详见《商户端接口敏感信息加密说明》
	SettleAcctNo *string `json:"SettleAcctNo,omitempty" name:"SettleAcctNo"`

	// 聚鑫分配的安全ID
	MidasSecretId *string `json:"MidasSecretId,omitempty" name:"MidasSecretId"`

	// 按照聚鑫安全密钥计算的签名
	MidasSignature *string `json:"MidasSignature,omitempty" name:"MidasSignature"`
}

func (r *UnBindAcctRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnBindAcctRequest) FromJsonString(s string) error {
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
}

func (r *UnbindRelateAcctRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnbindRelateAcctRequest) FromJsonString(s string) error {
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
	SubOrderList []*UnifiedOrderInSubOrderList `json:"SubOrderList,omitempty" name:"SubOrderList" list`

	// 结算应收金额，单位：分
	TotalMchIncome *int64 `json:"TotalMchIncome,omitempty" name:"TotalMchIncome"`

	// 平台应收金额，单位：分
	TotalPlatformIncome *int64 `json:"TotalPlatformIncome,omitempty" name:"TotalPlatformIncome"`

	// 微信公众号/小程序支付时为必选，需要传微信下的openid
	WxOpenId *string `json:"WxOpenId,omitempty" name:"WxOpenId"`

	// 在服务商模式下，微信公众号/小程序支付时wx_sub_openid和wx_openid二选一
	WxSubOpenId *string `json:"WxSubOpenId,omitempty" name:"WxSubOpenId"`
}

func (r *UnifiedOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnifiedOrderRequest) FromJsonString(s string) error {
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

func (r *UnifiedOrderResponse) FromJsonString(s string) error {
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
}

func (r *WithdrawCashMembershipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *WithdrawCashMembershipRequest) FromJsonString(s string) error {
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
