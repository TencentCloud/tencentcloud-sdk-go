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

package v20180301

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AttackRiskDetail struct {
	// 疑似的攻击痕迹类型
	// SuspectedSpoofingAttack：翻拍攻击
	// SuspectedSynthesisImage：疑似合成图片
	// SuspectedSynthesisVideo：疑似合成视频
	// SuspectedeAnomalyAttack：人脸特征疑似非真人
	// SuspectedAdversarialAttack：疑似对抗样本攻击
	// SuspectedBlackIndustry：疑似黑产批量模版攻击
	// SuspectedWatermark：疑似存在水印
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

// Predefined struct for user
type BankCard2EVerificationRequestParams struct {
	// 姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 银行卡
	BankCard *string `json:"BankCard,omitnil,omitempty" name:"BankCard"`

	// 敏感数据加密信息。对传入信息（姓名、银行卡号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

type BankCard2EVerificationRequest struct {
	*tchttp.BaseRequest
	
	// 姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 银行卡
	BankCard *string `json:"BankCard,omitnil,omitempty" name:"BankCard"`

	// 敏感数据加密信息。对传入信息（姓名、银行卡号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

func (r *BankCard2EVerificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BankCard2EVerificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "BankCard")
	delete(f, "Encryption")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BankCard2EVerificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BankCard2EVerificationResponseParams struct {
	// 认证结果码
	// 计费结果码：
	//   '0': '认证通过'
	//   '-1': '认证未通过'
	//  '-4': '持卡人信息有误'
	//   '-5': '未开通无卡支付'
	//   '-6': '此卡被没收'
	//   '-7': '无效卡号'
	//   '-8': '此卡无对应发卡行'
	//   '-9': '该卡未初始化或睡眠卡'
	//   '-10': '作弊卡、吞卡'
	//   '-11': '此卡已挂失'
	//   '-12': '该卡已过期'
	//   '-13': '受限制的卡'
	//   '-14': '密码错误次数超限'
	//   '-15': '发卡行不支持此交易'
	// 不计费结果码：
	//   '-2': '姓名校验不通过'
	//   '-3': '银行卡号码有误'
	//   '-16': '验证中心服务繁忙'
	//   '-17': '验证次数超限，请次日重试'
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 业务结果描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BankCard2EVerificationResponse struct {
	*tchttp.BaseResponse
	Response *BankCard2EVerificationResponseParams `json:"Response"`
}

func (r *BankCard2EVerificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BankCard2EVerificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BankCard4EVerificationRequestParams struct {
	// 姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 银行卡
	BankCard *string `json:"BankCard,omitnil,omitempty" name:"BankCard"`

	// 手机号码
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 开户证件号，与CertType参数的证件类型一致，如：身份证，则传入身份证号。
	IdCard *string `json:"IdCard,omitnil,omitempty" name:"IdCard"`

	// 证件类型，请确认该证件为开户时使用的证件类型，未用于开户的证件信息不支持验证。
	// 目前默认为0：身份证，其他证件类型暂不支持。
	CertType *int64 `json:"CertType,omitnil,omitempty" name:"CertType"`

	// 敏感数据加密信息。对传入信息（姓名、身份证号、手机号、银行卡号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

type BankCard4EVerificationRequest struct {
	*tchttp.BaseRequest
	
	// 姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 银行卡
	BankCard *string `json:"BankCard,omitnil,omitempty" name:"BankCard"`

	// 手机号码
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 开户证件号，与CertType参数的证件类型一致，如：身份证，则传入身份证号。
	IdCard *string `json:"IdCard,omitnil,omitempty" name:"IdCard"`

	// 证件类型，请确认该证件为开户时使用的证件类型，未用于开户的证件信息不支持验证。
	// 目前默认为0：身份证，其他证件类型暂不支持。
	CertType *int64 `json:"CertType,omitnil,omitempty" name:"CertType"`

	// 敏感数据加密信息。对传入信息（姓名、身份证号、手机号、银行卡号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

func (r *BankCard4EVerificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BankCard4EVerificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "BankCard")
	delete(f, "Phone")
	delete(f, "IdCard")
	delete(f, "CertType")
	delete(f, "Encryption")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BankCard4EVerificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BankCard4EVerificationResponseParams struct {
	// 认证结果码
	// 收费结果码：
	// '0': '认证通过'
	// '-1': '认证未通过'
	// '-6': '持卡人信息有误'
	// '-7': '未开通无卡支付'
	// '-8': '此卡被没收'
	// '-9': '无效卡号'
	// '-10': '此卡无对应发卡行'
	// '-11': '该卡未初始化或睡眠卡'
	// '-12': '作弊卡、吞卡'
	// '-13': '此卡已挂失'
	// '-14': '该卡已过期'
	// '-15': '受限制的卡'
	// '-16': '密码错误次数超限'
	// '-17': '发卡行不支持此交易'
	// 不收费结果码：
	// '-2': '姓名校验不通过'
	// '-3': '身份证号码有误'
	// '-4': '银行卡号码有误'
	// '-5': '手机号码不合法'
	// '-18': '验证中心服务繁忙'
	// '-19': '验证次数超限，请次日重试'
	// '-20': '该证件号暂不支持核验，当前仅支持二代身份证'
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 业务结果描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BankCard4EVerificationResponse struct {
	*tchttp.BaseResponse
	Response *BankCard4EVerificationResponseParams `json:"Response"`
}

func (r *BankCard4EVerificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BankCard4EVerificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BankCardVerificationRequestParams struct {
	// 开户证件号，与CertType参数的证件类型一致，如：身份证，则传入身份证号。
	IdCard *string `json:"IdCard,omitnil,omitempty" name:"IdCard"`

	// 姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 银行卡
	BankCard *string `json:"BankCard,omitnil,omitempty" name:"BankCard"`

	// 证件类型，请确认该证件为开户时使用的证件类型，未用于开户的证件信息不支持验证。
	// 目前默认：0 身份证，其他证件类型暂不支持。
	CertType *int64 `json:"CertType,omitnil,omitempty" name:"CertType"`

	// 敏感数据加密信息。对传入信息（姓名、身份证号、银行卡号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

type BankCardVerificationRequest struct {
	*tchttp.BaseRequest
	
	// 开户证件号，与CertType参数的证件类型一致，如：身份证，则传入身份证号。
	IdCard *string `json:"IdCard,omitnil,omitempty" name:"IdCard"`

	// 姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 银行卡
	BankCard *string `json:"BankCard,omitnil,omitempty" name:"BankCard"`

	// 证件类型，请确认该证件为开户时使用的证件类型，未用于开户的证件信息不支持验证。
	// 目前默认：0 身份证，其他证件类型暂不支持。
	CertType *int64 `json:"CertType,omitnil,omitempty" name:"CertType"`

	// 敏感数据加密信息。对传入信息（姓名、身份证号、银行卡号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

func (r *BankCardVerificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BankCardVerificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdCard")
	delete(f, "Name")
	delete(f, "BankCard")
	delete(f, "CertType")
	delete(f, "Encryption")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BankCardVerificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BankCardVerificationResponseParams struct {
	// 认证结果码
	// 收费结果码：
	// '0': '认证通过'
	// '-1': '认证未通过'
	// '-5': '持卡人信息有误'
	// '-6': '未开通无卡支付'
	// '-7': '此卡被没收'
	// '-8': '无效卡号'
	// '-9': '此卡无对应发卡行'
	// '-10': '该卡未初始化或睡眠卡'
	// '-11': '作弊卡、吞卡'
	// '-12': '此卡已挂失'
	// '-13': '该卡已过期'
	// '-14': '受限制的卡'
	// '-15': '密码错误次数超限'
	// '-16': '发卡行不支持此交易'
	// 不收费结果码：
	// '-2': '姓名校验不通过'
	// '-3': '身份证号码有误'
	// '-4': '银行卡号码有误'
	// '-17': '验证中心服务繁忙'
	// '-18': '验证次数超限，请次日重试'
	// '-19': '该证件号暂不支持核验，当前仅支持二代身份证'	
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 业务结果描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BankCardVerificationResponse struct {
	*tchttp.BaseResponse
	Response *BankCardVerificationResponseParams `json:"Response"`
}

func (r *BankCardVerificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BankCardVerificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChargeDetail struct {
	// 一比一时间时间戳，13位。
	ReqTime *string `json:"ReqTime,omitnil,omitempty" name:"ReqTime"`

	// 一比一请求的唯一标记。
	Seq *string `json:"Seq,omitnil,omitempty" name:"Seq"`

	// 一比一时使用的、脱敏后的身份证号。
	IdCard *string `json:"IdCard,omitnil,omitempty" name:"IdCard"`

	// 已废弃。请使用“IdCard”字段
	//
	// Deprecated: Idcard is deprecated.
	Idcard *string `json:"Idcard,omitnil,omitempty" name:"Idcard"`

	// 一比一时使用的、脱敏后的姓名。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 一比一的相似度。0-100，保留2位小数。
	Sim *string `json:"Sim,omitnil,omitempty" name:"Sim"`

	// 本次详情是否收费。
	IsNeedCharge *bool `json:"IsNeedCharge,omitnil,omitempty" name:"IsNeedCharge"`

	// 收费类型，比对、核身、混合部署。
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 本次活体一比一最终结果。
	ErrorCode *string `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// 本次活体一比一最终结果描述。
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`
}

// Predefined struct for user
type CheckBankCardInformationRequestParams struct {
	// 银行卡号。
	BankCard *string `json:"BankCard,omitnil,omitempty" name:"BankCard"`

	// 敏感数据加密信息。对传入信息（银行卡号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

type CheckBankCardInformationRequest struct {
	*tchttp.BaseRequest
	
	// 银行卡号。
	BankCard *string `json:"BankCard,omitnil,omitempty" name:"BankCard"`

	// 敏感数据加密信息。对传入信息（银行卡号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

func (r *CheckBankCardInformationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckBankCardInformationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BankCard")
	delete(f, "Encryption")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckBankCardInformationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckBankCardInformationResponseParams struct {
	// 认证结果码，收费情况如下。
	// 收费结果码：
	// 0: 查询成功
	// -1: 未查到信息
	// 不收费结果码：
	// -2：验证中心服务繁忙
	// -3：银行卡不存在
	// -4：认证次数超过当日限制，请次日重试
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 业务结果描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 开户行
	AccountBank *string `json:"AccountBank,omitnil,omitempty" name:"AccountBank"`

	// 卡性质：1. 借记卡；2. 贷记卡；3. 预付费卡；4. 准贷记卡
	AccountType *int64 `json:"AccountType,omitnil,omitempty" name:"AccountType"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckBankCardInformationResponse struct {
	*tchttp.BaseResponse
	Response *CheckBankCardInformationResponseParams `json:"Response"`
}

func (r *CheckBankCardInformationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckBankCardInformationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckEidTokenStatusRequestParams struct {
	// E证通流程的唯一标识，调用GetEidToken接口时生成。
	EidToken *string `json:"EidToken,omitnil,omitempty" name:"EidToken"`
}

type CheckEidTokenStatusRequest struct {
	*tchttp.BaseRequest
	
	// E证通流程的唯一标识，调用GetEidToken接口时生成。
	EidToken *string `json:"EidToken,omitnil,omitempty" name:"EidToken"`
}

func (r *CheckEidTokenStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckEidTokenStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EidToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckEidTokenStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckEidTokenStatusResponseParams struct {
	// 枚举：
	// init：token未验证
	// doing: 验证中
	// finished: 验证完成
	// timeout: token已超时
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckEidTokenStatusResponse struct {
	*tchttp.BaseResponse
	Response *CheckEidTokenStatusResponseParams `json:"Response"`
}

func (r *CheckEidTokenStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckEidTokenStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckIdCardInformationRequestParams struct {
	// 身份证人像面的 Base64 值
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。
	// 请使用标准的Base64编码方式(带=补位)，编码规范参考RFC4648。
	// ImageBase64、ImageUrl二者必须提供其中之一。若都提供了，则按照ImageUrl>ImageBase64的优先级使用参数。
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// 身份证人像面的 Url 地址
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 3M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 以下可选字段均为bool 类型，默认false：
	// CopyWarn，复印件告警
	// BorderCheckWarn，边框和框内遮挡告警
	// ReshootWarn，翻拍告警
	// DetectPsWarn，PS检测告警（疑似存在PS痕迹）
	// TempIdWarn，临时身份证告警
	// Quality，图片质量告警（评价图片模糊程度）
	// 
	// SDK 设置方式参考：
	// Config = Json.stringify({"CopyWarn":true,"ReshootWarn":true})
	// API 3.0 Explorer 设置方式参考：
	// Config = {"CopyWarn":true,"ReshootWarn":true}
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`

	// 是否需要对返回中的敏感信息进行加密。默认false。
	// 其中敏感信息包括：Response.IdNum、Response.Name
	IsEncrypt *bool `json:"IsEncrypt,omitnil,omitempty" name:"IsEncrypt"`

	// 是否需要对响应体加密
	IsEncryptResponse *bool `json:"IsEncryptResponse,omitnil,omitempty" name:"IsEncryptResponse"`

	// 是否需要对返回中的敏感信息进行加密,需指定加密算法Algorithm、CBC加密的初始向量、加密后的对称密钥。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

type CheckIdCardInformationRequest struct {
	*tchttp.BaseRequest
	
	// 身份证人像面的 Base64 值
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。
	// 请使用标准的Base64编码方式(带=补位)，编码规范参考RFC4648。
	// ImageBase64、ImageUrl二者必须提供其中之一。若都提供了，则按照ImageUrl>ImageBase64的优先级使用参数。
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// 身份证人像面的 Url 地址
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 3M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 以下可选字段均为bool 类型，默认false：
	// CopyWarn，复印件告警
	// BorderCheckWarn，边框和框内遮挡告警
	// ReshootWarn，翻拍告警
	// DetectPsWarn，PS检测告警（疑似存在PS痕迹）
	// TempIdWarn，临时身份证告警
	// Quality，图片质量告警（评价图片模糊程度）
	// 
	// SDK 设置方式参考：
	// Config = Json.stringify({"CopyWarn":true,"ReshootWarn":true})
	// API 3.0 Explorer 设置方式参考：
	// Config = {"CopyWarn":true,"ReshootWarn":true}
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`

	// 是否需要对返回中的敏感信息进行加密。默认false。
	// 其中敏感信息包括：Response.IdNum、Response.Name
	IsEncrypt *bool `json:"IsEncrypt,omitnil,omitempty" name:"IsEncrypt"`

	// 是否需要对响应体加密
	IsEncryptResponse *bool `json:"IsEncryptResponse,omitnil,omitempty" name:"IsEncryptResponse"`

	// 是否需要对返回中的敏感信息进行加密,需指定加密算法Algorithm、CBC加密的初始向量、加密后的对称密钥。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

func (r *CheckIdCardInformationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckIdCardInformationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "Config")
	delete(f, "IsEncrypt")
	delete(f, "IsEncryptResponse")
	delete(f, "Encryption")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckIdCardInformationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckIdCardInformationResponseParams struct {
	// 相似度，取值范围 [0.00, 100.00]。推荐相似度大于等于70时可判断为同一人，可根据具体场景自行调整阈值（阈值70的误通过率为千分之一，阈值80的误通过率是万分之一）
	Sim *float64 `json:"Sim,omitnil,omitempty" name:"Sim"`

	// 业务错误码，成功情况返回Success, 错误情况请参考下方错误码 列表中FailedOperation部分
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 业务结果描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 性别
	Sex *string `json:"Sex,omitnil,omitempty" name:"Sex"`

	// 民族
	Nation *string `json:"Nation,omitnil,omitempty" name:"Nation"`

	// 出生日期
	Birth *string `json:"Birth,omitnil,omitempty" name:"Birth"`

	// 地址
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// 身份证号
	IdNum *string `json:"IdNum,omitnil,omitempty" name:"IdNum"`

	// 身份证头像照片的base64编码，如果抠图失败会拿整张身份证做比对并返回空。
	Portrait *string `json:"Portrait,omitnil,omitempty" name:"Portrait"`

	// 告警信息，当在Config中配置了告警信息会停止人像比对，Result返回错误（FailedOperation.OcrWarningOccurred）并有此告警信息，Code 告警码列表和释义：
	// 
	// -9101 身份证边框不完整告警，
	// -9102 身份证复印件告警，
	// -9103 身份证翻拍告警，
	// -9105 身份证框内遮挡告警，
	// -9104 临时身份证告警，
	// -9106 身份证 PS 告警（疑似存在PS痕迹）。
	// -8001 图片模糊告警
	// 多个会 |  隔开如 "-9101|-9106|-9104"
	Warnings *string `json:"Warnings,omitnil,omitempty" name:"Warnings"`

	// 图片质量分数，当请求Config中配置图片模糊告警该参数才有意义，取值范围（0～100），目前默认阈值是50分，低于50分会触发模糊告警。
	Quality *float64 `json:"Quality,omitnil,omitempty" name:"Quality"`

	// 敏感数据加密信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`

	// 加密后的数据
	EncryptedBody *string `json:"EncryptedBody,omitnil,omitempty" name:"EncryptedBody"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckIdCardInformationResponse struct {
	*tchttp.BaseResponse
	Response *CheckIdCardInformationResponseParams `json:"Response"`
}

func (r *CheckIdCardInformationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckIdCardInformationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckIdNameDateRequestParams struct {
	// 姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 身份证号
	IdCard *string `json:"IdCard,omitnil,omitempty" name:"IdCard"`

	// 身份证有效期开始时间，格式：YYYYMMDD。如：20210701
	ValidityBegin *string `json:"ValidityBegin,omitnil,omitempty" name:"ValidityBegin"`

	// 身份证有效期到期时间，格式：YYYYMMDD，长期用“00000000”代替；如：20210701
	ValidityEnd *string `json:"ValidityEnd,omitnil,omitempty" name:"ValidityEnd"`

	// 敏感数据加密信息。对传入信息（姓名、身份证号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

type CheckIdNameDateRequest struct {
	*tchttp.BaseRequest
	
	// 姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 身份证号
	IdCard *string `json:"IdCard,omitnil,omitempty" name:"IdCard"`

	// 身份证有效期开始时间，格式：YYYYMMDD。如：20210701
	ValidityBegin *string `json:"ValidityBegin,omitnil,omitempty" name:"ValidityBegin"`

	// 身份证有效期到期时间，格式：YYYYMMDD，长期用“00000000”代替；如：20210701
	ValidityEnd *string `json:"ValidityEnd,omitnil,omitempty" name:"ValidityEnd"`

	// 敏感数据加密信息。对传入信息（姓名、身份证号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

func (r *CheckIdNameDateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckIdNameDateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "IdCard")
	delete(f, "ValidityBegin")
	delete(f, "ValidityEnd")
	delete(f, "Encryption")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckIdNameDateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckIdNameDateResponseParams struct {
	// 认证结果码，收费情况如下。
	// 收费结果码：
	// 0: 一致
	// -1: 不一致
	// 不收费结果码：
	// -2: 非法身份证号（长度、校验位等不正确）
	// -3: 非法姓名（长度、格式等不正确）
	// -4: 非法有效期（长度、格式等不正确）
	// -5: 身份信息无效
	// -6: 证件库服务异常
	// -7: 证件库中无此身份证记录
	// -8: 认证次数超过当日限制，请次日重试
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 业务结果描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckIdNameDateResponse struct {
	*tchttp.BaseResponse
	Response *CheckIdNameDateResponseParams `json:"Response"`
}

func (r *CheckIdNameDateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckIdNameDateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckPhoneAndNameRequestParams struct {
	// ⼿机号
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 敏感数据加密信息。对传入信息（姓名、手机号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

type CheckPhoneAndNameRequest struct {
	*tchttp.BaseRequest
	
	// ⼿机号
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 敏感数据加密信息。对传入信息（姓名、手机号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

func (r *CheckPhoneAndNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckPhoneAndNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Mobile")
	delete(f, "Name")
	delete(f, "Encryption")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckPhoneAndNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckPhoneAndNameResponseParams struct {
	// 认证结果码，收费情况如下。
	// 收费结果码：
	// 0: 验证结果一致
	// 1: 验证结果不一致
	// 不收费结果码：
	// -1:查无记录
	// -2:引擎未知错误
	// -3:引擎服务异常
	// -4:姓名校验不通过
	// -5:手机号码不合法
	// -6: 认证次数超过当日限制，请次日重试
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 业务结果描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckPhoneAndNameResponse struct {
	*tchttp.BaseResponse
	Response *CheckPhoneAndNameResponseParams `json:"Response"`
}

func (r *CheckPhoneAndNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckPhoneAndNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectAIFakeFacesRequestParams struct {
	// 传入需要进行检测的带有人脸的图片或视频，使用base64编码的形式。
	// 
	// 图片的Base64值：
	// 建议整体图像480x640的分辨率，脸部 大小 100X100 以上；
	// Base64编码后的图片数据大小不超过3M，仅支持jpg、png格式。请使用标准的Base64编码方式(带=补位)，编码规范参考RFC4648。
	// 
	// 视频的Base64值：
	// Base64编码后的大小不超过8M，支持mp4、avi、flv格式。
	// 请使用标准的Base64编码方式(带=补位)，编码规范参考RFC4648。
	// 视频时长最大支持20s，建议时长2～5s。
	// 建议视频分辨率为480x640，帧率在25fps~30fps之间。
	FaceInput *string `json:"FaceInput,omitnil,omitempty" name:"FaceInput"`

	// 传入的类型
	// 1- 传入的是图片类型
	// 2- 传入的是视频类型
	// 其他 - 返回错误码InvalidParameter
	FaceInputType *int64 `json:"FaceInputType,omitnil,omitempty" name:"FaceInputType"`

	// 是否需要对请求信息进行全包体加密； 支持的加密算法:AES-256-CBC、SM4-GCM； 有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`

	// 加密后的密文； 加密前的数据格式如下:{"FaceInput":"AAAAA","FaceInputType":1}
	EncryptedBody *string `json:"EncryptedBody,omitnil,omitempty" name:"EncryptedBody"`
}

type DetectAIFakeFacesRequest struct {
	*tchttp.BaseRequest
	
	// 传入需要进行检测的带有人脸的图片或视频，使用base64编码的形式。
	// 
	// 图片的Base64值：
	// 建议整体图像480x640的分辨率，脸部 大小 100X100 以上；
	// Base64编码后的图片数据大小不超过3M，仅支持jpg、png格式。请使用标准的Base64编码方式(带=补位)，编码规范参考RFC4648。
	// 
	// 视频的Base64值：
	// Base64编码后的大小不超过8M，支持mp4、avi、flv格式。
	// 请使用标准的Base64编码方式(带=补位)，编码规范参考RFC4648。
	// 视频时长最大支持20s，建议时长2～5s。
	// 建议视频分辨率为480x640，帧率在25fps~30fps之间。
	FaceInput *string `json:"FaceInput,omitnil,omitempty" name:"FaceInput"`

	// 传入的类型
	// 1- 传入的是图片类型
	// 2- 传入的是视频类型
	// 其他 - 返回错误码InvalidParameter
	FaceInputType *int64 `json:"FaceInputType,omitnil,omitempty" name:"FaceInputType"`

	// 是否需要对请求信息进行全包体加密； 支持的加密算法:AES-256-CBC、SM4-GCM； 有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`

	// 加密后的密文； 加密前的数据格式如下:{"FaceInput":"AAAAA","FaceInputType":1}
	EncryptedBody *string `json:"EncryptedBody,omitnil,omitempty" name:"EncryptedBody"`
}

func (r *DetectAIFakeFacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectAIFakeFacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FaceInput")
	delete(f, "FaceInputType")
	delete(f, "Encryption")
	delete(f, "EncryptedBody")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetectAIFakeFacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectAIFakeFacesResponseParams struct {
	// 检测到的图片是否存在攻击：
	// Low：无攻击风险
	// Mid：中度疑似攻击
	// High：高度疑似攻击
	AttackRiskLevel *string `json:"AttackRiskLevel,omitnil,omitempty" name:"AttackRiskLevel"`

	// 检测到疑似的攻击痕迹列表
	// 说明：未检测到攻击痕迹时，返回空数组
	// 此出参仅作为结果判断的参考，实际应用仍建议使用AttackRiskLevel的结果。
	AttackRiskDetailList []*AttackRiskDetail `json:"AttackRiskDetailList,omitnil,omitempty" name:"AttackRiskDetailList"`

	// 额外信息
	ExtraInfo *ExtraInfo `json:"ExtraInfo,omitnil,omitempty" name:"ExtraInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DetectAIFakeFacesResponse struct {
	*tchttp.BaseResponse
	Response *DetectAIFakeFacesResponseParams `json:"Response"`
}

func (r *DetectAIFakeFacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectAIFakeFacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectAuthRequestParams struct {
	// 用于细分客户使用场景，申请开通服务后，可以在腾讯云慧眼人脸核身控制台（https://console.cloud.tencent.com/faceid） 自助接入里面创建，审核通过后即可调用。如有疑问，请添加[腾讯云人脸核身小助手](https://cloud.tencent.com/document/product/1007/56130)进行咨询。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 本接口不需要传递此参数。
	TerminalType *string `json:"TerminalType,omitnil,omitempty" name:"TerminalType"`

	// 身份标识（是否必传基于[控制台](https://console.cloud.tencent.com/faceid/access)申请业务流程时配置的提示）。
	// 规则：a-z，A-Z，0-9组合。最长长度32位。
	IdCard *string `json:"IdCard,omitnil,omitempty" name:"IdCard"`

	// 姓名。（是否必传基于[控制台](https://console.cloud.tencent.com/faceid/access)申请业务流程时配置的提示）。
	// 最长长度32位。中文请使用UTF-8编码。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 认证结束后重定向的回调链接地址。最长长度1024位。
	RedirectUrl *string `json:"RedirectUrl,omitnil,omitempty" name:"RedirectUrl"`

	// 透传字段，在获取验证结果时返回。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 用于人脸比对的照片，图片的Base64值；
	// Base64编码后的图片数据大小不超过3M，仅支持jpg、png格式。请使用标准的Base64编码方式(带=补位)，编码规范参考RFC4648。
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// 敏感数据加密信息。对传入信息（姓名、身份证号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`

	// 意愿核身（朗读模式）使用的文案，若未使用意愿核身（朗读模式），则该字段无需传入。默认为空，最长可接受120的字符串长度。
	IntentionVerifyText *string `json:"IntentionVerifyText,omitnil,omitempty" name:"IntentionVerifyText"`

	// 意愿核身语音问答模式（即语音播报+语音回答）使用的文案，包括：系统语音播报的文本、需要核验的标准文本。当前仅支持1轮问答。
	IntentionQuestions []*IntentionQuestion `json:"IntentionQuestions,omitnil,omitempty" name:"IntentionQuestions"`

	// RuleId相关配置
	Config *RuleIdConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// 意愿核身（点头确认模式）使用的文案，若未使用意愿核身（点头确认模式），则该字段无需传入。当前仅支持一个提示文本。
	IntentionActions []*IntentionActionConfig `json:"IntentionActions,omitnil,omitempty" name:"IntentionActions"`
}

type DetectAuthRequest struct {
	*tchttp.BaseRequest
	
	// 用于细分客户使用场景，申请开通服务后，可以在腾讯云慧眼人脸核身控制台（https://console.cloud.tencent.com/faceid） 自助接入里面创建，审核通过后即可调用。如有疑问，请添加[腾讯云人脸核身小助手](https://cloud.tencent.com/document/product/1007/56130)进行咨询。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 本接口不需要传递此参数。
	TerminalType *string `json:"TerminalType,omitnil,omitempty" name:"TerminalType"`

	// 身份标识（是否必传基于[控制台](https://console.cloud.tencent.com/faceid/access)申请业务流程时配置的提示）。
	// 规则：a-z，A-Z，0-9组合。最长长度32位。
	IdCard *string `json:"IdCard,omitnil,omitempty" name:"IdCard"`

	// 姓名。（是否必传基于[控制台](https://console.cloud.tencent.com/faceid/access)申请业务流程时配置的提示）。
	// 最长长度32位。中文请使用UTF-8编码。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 认证结束后重定向的回调链接地址。最长长度1024位。
	RedirectUrl *string `json:"RedirectUrl,omitnil,omitempty" name:"RedirectUrl"`

	// 透传字段，在获取验证结果时返回。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 用于人脸比对的照片，图片的Base64值；
	// Base64编码后的图片数据大小不超过3M，仅支持jpg、png格式。请使用标准的Base64编码方式(带=补位)，编码规范参考RFC4648。
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// 敏感数据加密信息。对传入信息（姓名、身份证号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`

	// 意愿核身（朗读模式）使用的文案，若未使用意愿核身（朗读模式），则该字段无需传入。默认为空，最长可接受120的字符串长度。
	IntentionVerifyText *string `json:"IntentionVerifyText,omitnil,omitempty" name:"IntentionVerifyText"`

	// 意愿核身语音问答模式（即语音播报+语音回答）使用的文案，包括：系统语音播报的文本、需要核验的标准文本。当前仅支持1轮问答。
	IntentionQuestions []*IntentionQuestion `json:"IntentionQuestions,omitnil,omitempty" name:"IntentionQuestions"`

	// RuleId相关配置
	Config *RuleIdConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// 意愿核身（点头确认模式）使用的文案，若未使用意愿核身（点头确认模式），则该字段无需传入。当前仅支持一个提示文本。
	IntentionActions []*IntentionActionConfig `json:"IntentionActions,omitnil,omitempty" name:"IntentionActions"`
}

func (r *DetectAuthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectAuthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	delete(f, "TerminalType")
	delete(f, "IdCard")
	delete(f, "Name")
	delete(f, "RedirectUrl")
	delete(f, "Extra")
	delete(f, "ImageBase64")
	delete(f, "Encryption")
	delete(f, "IntentionVerifyText")
	delete(f, "IntentionQuestions")
	delete(f, "Config")
	delete(f, "IntentionActions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetectAuthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectAuthResponseParams struct {
	// 用于发起核身流程的URL，仅微信H5场景使用。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 一次核身流程的标识，有效时间为7,200秒；
	// 完成核身后，可用该标识获取验证结果信息。
	BizToken *string `json:"BizToken,omitnil,omitempty" name:"BizToken"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DetectAuthResponse struct {
	*tchttp.BaseResponse
	Response *DetectAuthResponseParams `json:"Response"`
}

func (r *DetectAuthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectAuthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetectDetail struct {
	// 请求时间戳。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReqTime *string `json:"ReqTime,omitnil,omitempty" name:"ReqTime"`

	// 本次活体一比一请求的唯一标记。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Seq *string `json:"Seq,omitnil,omitempty" name:"Seq"`

	// 参与本次活体一比一的身份证号。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Idcard *string `json:"Idcard,omitnil,omitempty" name:"Idcard"`

	// 参与本次活体一比一的姓名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 本次活体一比一的相似度。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sim *string `json:"Sim,omitnil,omitempty" name:"Sim"`

	// 本次活体一比一是否收费
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsNeedCharge *bool `json:"IsNeedCharge,omitnil,omitempty" name:"IsNeedCharge"`

	// 本次活体一比一最终结果。0为成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Errcode *int64 `json:"Errcode,omitnil,omitempty" name:"Errcode"`

	// 本次活体一比一最终结果描述。（仅描述用，文案更新时不会通知。）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Errmsg *string `json:"Errmsg,omitnil,omitempty" name:"Errmsg"`

	// 本次活体结果。0为成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Livestatus *int64 `json:"Livestatus,omitnil,omitempty" name:"Livestatus"`

	// 本次活体结果描述。（仅描述用，文案更新时不会通知。）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Livemsg *string `json:"Livemsg,omitnil,omitempty" name:"Livemsg"`

	// 本次一比一结果。0为成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comparestatus *int64 `json:"Comparestatus,omitnil,omitempty" name:"Comparestatus"`

	// 本次一比一结果描述。（仅描述用，文案更新时不会通知。）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comparemsg *string `json:"Comparemsg,omitnil,omitempty" name:"Comparemsg"`

	// 比对库源类型。包括：
	// 公安商业库；
	// 业务方自有库（用户上传照片、客户的混合库、混合部署库）；
	// 二次验证库；
	// 人工审核库；
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompareLibType *string `json:"CompareLibType,omitnil,omitempty" name:"CompareLibType"`

	// 枚举活体检测类型：
	// 0：未知
	// 1：数字活体
	// 2：动作活体
	// 3：静默活体
	// 4：一闪活体（动作+光线）
	// 注意：此字段可能返回 null，表示取不到有效值。
	LivenessMode *uint64 `json:"LivenessMode,omitnil,omitempty" name:"LivenessMode"`
}

type DetectInfoBestFrame struct {
	// 活体比对最佳帧Base64编码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BestFrame *string `json:"BestFrame,omitnil,omitempty" name:"BestFrame"`

	// 自截帧Base64编码数组。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BestFrames []*string `json:"BestFrames,omitnil,omitempty" name:"BestFrames"`
}

type DetectInfoIdCardData struct {
	// OCR正面照片的base64编码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrFront *string `json:"OcrFront,omitnil,omitempty" name:"OcrFront"`

	// OCR反面照片的base64编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrBack *string `json:"OcrBack,omitnil,omitempty" name:"OcrBack"`

	// 旋转裁边后的正面照片base64编码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessedFrontImage *string `json:"ProcessedFrontImage,omitnil,omitempty" name:"ProcessedFrontImage"`

	// 旋转裁边后的背面照片base64编码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessedBackImage *string `json:"ProcessedBackImage,omitnil,omitempty" name:"ProcessedBackImage"`

	// 身份证正面人像图base64编码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`

	// 身份证人像面告警码，开启身份证告警功能后才会返回，返回数组中可能出现的告警码如下：
	// -9100 身份证有效日期不合法告警，
	// -9101 身份证边框不完整告警，
	// -9102 身份证复印件告警，
	// -9103 身份证翻拍告警，
	// -9105 身份证框内遮挡告警，
	// -9104 临时身份证告警，
	// -9106 身份证 PS 告警（疑似存在PS痕迹），
	// -9107 身份证反光告警。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WarnInfos []*int64 `json:"WarnInfos,omitnil,omitempty" name:"WarnInfos"`

	// 身份证国徽面告警码，开启身份证告警功能后才会返回，返回数组中可能出现的告警码如下：
	// -9100 身份证有效日期不合法告警，
	// -9101 身份证边框不完整告警，
	// -9102 身份证复印件告警，
	// -9103 身份证翻拍告警，
	// -9105 身份证框内遮挡告警，
	// -9104 临时身份证告警，
	// -9106 身份证 PS 告警（疑似存在PS痕迹），
	// -9107 身份证反光告警。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackWarnInfos []*int64 `json:"BackWarnInfos,omitnil,omitempty" name:"BackWarnInfos"`
}

type DetectInfoText struct {
	// 本次流程最终验证结果。0为成功（仅包含活体人脸核身结果，不包含意愿核身结果）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrCode *int64 `json:"ErrCode,omitnil,omitempty" name:"ErrCode"`

	// 本次流程最终验证结果描述。（仅描述用，文案更新时不会通知。）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMsg *string `json:"ErrMsg,omitnil,omitempty" name:"ErrMsg"`

	// 本次验证使用的身份证号。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdCard *string `json:"IdCard,omitnil,omitempty" name:"IdCard"`

	// 用户认证时使用的证件号码类型：
	// 0：二代身份证的证件号码
	// 1：港澳台居住证的证件号码
	// 2：其他（核验使用的证件号码非合法身份号码）
	// 注意：此字段可能返回 null，表示取不到有效值。
	UseIDType *uint64 `json:"UseIDType,omitnil,omitempty" name:"UseIDType"`

	// 本次验证使用的姓名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 身份校验环节识别结果：民族。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrNation *string `json:"OcrNation,omitnil,omitempty" name:"OcrNation"`

	// 身份校验环节识别结果：家庭住址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrAddress *string `json:"OcrAddress,omitnil,omitempty" name:"OcrAddress"`

	// 身份校验环节识别结果：生日。格式为：YYYY/M/D
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrBirth *string `json:"OcrBirth,omitnil,omitempty" name:"OcrBirth"`

	// 身份校验环节识别结果：签发机关。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrAuthority *string `json:"OcrAuthority,omitnil,omitempty" name:"OcrAuthority"`

	// 身份校验环节识别结果：有效日期。格式为：YYYY.MM.DD-YYYY.MM.DD
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrValidDate *string `json:"OcrValidDate,omitnil,omitempty" name:"OcrValidDate"`

	// 身份校验环节识别结果：姓名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrName *string `json:"OcrName,omitnil,omitempty" name:"OcrName"`

	// 身份校验环节识别结果：身份证号。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrIdCard *string `json:"OcrIdCard,omitnil,omitempty" name:"OcrIdCard"`

	// 身份校验环节识别结果：性别。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrGender *string `json:"OcrGender,omitnil,omitempty" name:"OcrGender"`

	// 身份校验环节采用的信息上传方式。
	// 取值有"NFC"、"OCR"、"手动输入"、"其他"
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdInfoFrom *string `json:"IdInfoFrom,omitnil,omitempty" name:"IdInfoFrom"`

	// 本次流程最终活体结果。0为成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	LiveStatus *int64 `json:"LiveStatus,omitnil,omitempty" name:"LiveStatus"`

	// 本次流程最终活体结果描述。（仅描述用，文案更新时不会通知。）
	// 注意：此字段可能返回 null，表示取不到有效值。
	LiveMsg *string `json:"LiveMsg,omitnil,omitempty" name:"LiveMsg"`

	// 本次流程最终一比一结果。0为成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comparestatus *int64 `json:"Comparestatus,omitnil,omitempty" name:"Comparestatus"`

	// 本次流程最终一比一结果描述。（仅描述用，文案更新时不会通知。）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comparemsg *string `json:"Comparemsg,omitnil,omitempty" name:"Comparemsg"`

	// 本次流程活体一比一的分数，取值范围 [0.00, 100.00]。相似度大于等于70时才判断为同一人，也可根据具体场景自行调整阈值（阈值70的误通过率为千分之一，阈值80的误通过率是万分之一）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sim *string `json:"Sim,omitnil,omitempty" name:"Sim"`

	// 地理位置经纬度。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Location *string `json:"Location,omitnil,omitempty" name:"Location"`

	// Auth接口带入额外信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 本次流程进行的活体一比一流水。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LivenessDetail []*DetectDetail `json:"LivenessDetail,omitnil,omitempty" name:"LivenessDetail"`

	// 描述当前请求活体阶段被拒绝的详细原因，详情如下：
	// 01-用户全程闭眼
	// 02-用户未完成指定动作
	// 03-疑似翻拍攻击
	// 04-疑似合成攻击
	// 05-疑似黑产模版
	// 06-疑似存在水印
	// 07-反光校验未通过
	// 08-疑似中途换人
	// 09-人脸质量过差
	// 10-距离校验不通过
	// 11-疑似对抗样本攻击
	// 12-嘴巴区域疑似存在攻击痕迹
	// 13-眼睛区域疑似存在攻击痕迹
	// 14-眼睛或嘴巴被遮挡
	// 注意：此字段可能返回 null，表示取不到有效值。
	LivenessInfoTag []*string `json:"LivenessInfoTag,omitnil,omitempty" name:"LivenessInfoTag"`

	// 手机号码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 本次流程最终比对库源类型。包括：
	// 权威库；
	// 业务方自有库（用户上传照片、客户的混合库、混合部署库）；
	// 二次验证库；
	// 人工审核库；
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompareLibType *string `json:"CompareLibType,omitnil,omitempty" name:"CompareLibType"`

	// 本次流程最终活体类型。包括：
	// 0：未知
	// 1：数字活体
	// 2：动作活体
	// 3：静默活体
	// 4：一闪活体（动作+光线）
	// 注意：此字段可能返回 null，表示取不到有效值。
	LivenessMode *uint64 `json:"LivenessMode,omitnil,omitempty" name:"LivenessMode"`

	// nfc重复计费requestId列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NFCRequestIds []*string `json:"NFCRequestIds,omitnil,omitempty" name:"NFCRequestIds"`

	// nfc重复计费计数
	// 注意：此字段可能返回 null，表示取不到有效值。
	NFCBillingCounts *int64 `json:"NFCBillingCounts,omitnil,omitempty" name:"NFCBillingCounts"`

	// 港澳台居住证通行证号码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PassNo *string `json:"PassNo,omitnil,omitempty" name:"PassNo"`

	// 港澳台居住证签发次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	VisaNum *string `json:"VisaNum,omitnil,omitempty" name:"VisaNum"`
}

type DetectInfoVideoData struct {
	// 活体视频的base64编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	LivenessVideo *string `json:"LivenessVideo,omitnil,omitempty" name:"LivenessVideo"`
}

type EidInfo struct {
	// 商户方 appeIDcode 的数字证书
	EidCode *string `json:"EidCode,omitnil,omitempty" name:"EidCode"`

	// Eid中心针对商户方EidCode的电子签名
	EidSign *string `json:"EidSign,omitnil,omitempty" name:"EidSign"`

	// 商户方公钥加密的会话密钥的base64字符串，[指引详见](https://cloud.tencent.com/document/product/1007/63370)
	DesKey *string `json:"DesKey,omitnil,omitempty" name:"DesKey"`

	// 会话密钥sm2加密后的base64字符串，[指引详见](https://cloud.tencent.com/document/product/1007/63370)
	UserInfo *string `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`
}

// Predefined struct for user
type EncryptedPhoneVerificationRequestParams struct {
	// 身份证号，加密方式以EncryptionMode为准
	IdCard *string `json:"IdCard,omitnil,omitempty" name:"IdCard"`

	// 姓名，加密方式以EncryptionMode为准
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 手机号，加密方式以EncryptionMode为准
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 敏感信息的加密方式，目前支持明文、MD5和SHA256加密传输，参数取值：
	// 
	// 0：明文，不加密
	// 1：使用MD5加密
	// 2：使用SHA256
	// 3：使用SM3加密
	EncryptionMode *string `json:"EncryptionMode,omitnil,omitempty" name:"EncryptionMode"`
}

type EncryptedPhoneVerificationRequest struct {
	*tchttp.BaseRequest
	
	// 身份证号，加密方式以EncryptionMode为准
	IdCard *string `json:"IdCard,omitnil,omitempty" name:"IdCard"`

	// 姓名，加密方式以EncryptionMode为准
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 手机号，加密方式以EncryptionMode为准
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 敏感信息的加密方式，目前支持明文、MD5和SHA256加密传输，参数取值：
	// 
	// 0：明文，不加密
	// 1：使用MD5加密
	// 2：使用SHA256
	// 3：使用SM3加密
	EncryptionMode *string `json:"EncryptionMode,omitnil,omitempty" name:"EncryptionMode"`
}

func (r *EncryptedPhoneVerificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EncryptedPhoneVerificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdCard")
	delete(f, "Name")
	delete(f, "Phone")
	delete(f, "EncryptionMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EncryptedPhoneVerificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EncryptedPhoneVerificationResponseParams struct {
	// 认证结果码:
	// 【收费结果码】
	// 0:   三要素信息一致
	// -4:  三要素信息不一致
	// 
	// 【不收费结果码】
	// -7: 身份证号码有误
	// -8: 参数错误
	// -9: 没有记录
	// -11: 验证中心服务繁忙
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 业务结果描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 运营商名称。
	// 取值范围为["移动","联通","电信",""]
	ISP *string `json:"ISP,omitnil,omitempty" name:"ISP"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EncryptedPhoneVerificationResponse struct {
	*tchttp.BaseResponse
	Response *EncryptedPhoneVerificationResponseParams `json:"Response"`
}

func (r *EncryptedPhoneVerificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EncryptedPhoneVerificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Encryption struct {
	// 在使用加密服务时，填入要被加密的字段。本接口中可填入加密后的一个或多个字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	EncryptList []*string `json:"EncryptList,omitnil,omitempty" name:"EncryptList"`

	// 加密后的对称密钥，关于密钥的生成和使用请查阅<a href="https://cloud.tencent.com/document/product/1007/47180">数据加密</a> 文档。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CiphertextBlob *string `json:"CiphertextBlob,omitnil,omitempty" name:"CiphertextBlob"`

	// 有加密需求的用户，传入CBC加密的初始向量（客户自定义字符串，长度16字符）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Iv *string `json:"Iv,omitnil,omitempty" name:"Iv"`

	// 加密使用的算法（支持'AES-256-CBC'、'SM4-GCM'），不传默认为'AES-256-CBC'
	// 注意：此字段可能返回 null，表示取不到有效值。
	Algorithm *string `json:"Algorithm,omitnil,omitempty" name:"Algorithm"`

	// SM4-GCM算法生成的消息摘要（校验消息完整性时使用）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagList []*string `json:"TagList,omitnil,omitempty" name:"TagList"`
}

type ExtraInfo struct {
	// 命中模板的详细信息，仅返回命中的相似度最高的模板信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetrievalLivenessExtraInfo []*RetrievalLivenessExtraInfo `json:"RetrievalLivenessExtraInfo,omitnil,omitempty" name:"RetrievalLivenessExtraInfo"`
}

// Predefined struct for user
type GetActionSequenceRequestParams struct {
	// 默认不需要使用
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`
}

type GetActionSequenceRequest struct {
	*tchttp.BaseRequest
	
	// 默认不需要使用
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`
}

func (r *GetActionSequenceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetActionSequenceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ActionType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetActionSequenceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetActionSequenceResponseParams struct {
	// 动作顺序(2,1 or 1,2) 。1代表张嘴，2代表闭眼。
	ActionSequence *string `json:"ActionSequence,omitnil,omitempty" name:"ActionSequence"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetActionSequenceResponse struct {
	*tchttp.BaseResponse
	Response *GetActionSequenceResponseParams `json:"Response"`
}

func (r *GetActionSequenceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetActionSequenceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDetectInfoEnhancedRequestParams struct {
	// 人脸核身流程的标识，调用DetectAuth接口时生成。
	BizToken *string `json:"BizToken,omitnil,omitempty" name:"BizToken"`

	// 用于细分客户使用场景，由腾讯侧在线下对接时分配。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 指定拉取的结果信息，取值（0：全部；1：文本类；2：身份证信息；3：视频最佳截图信息）。
	// 如 13表示拉取文本类、视频最佳截图信息。
	// 默认值：0
	InfoType *string `json:"InfoType,omitnil,omitempty" name:"InfoType"`

	// 从活体视频中截取一定张数的最佳帧（仅部分服务支持，若需使用请与慧眼小助手沟通）。默认为0，最大为10，超出10的最多只给10张。（InfoType需要包含3）
	BestFramesCount *uint64 `json:"BestFramesCount,omitnil,omitempty" name:"BestFramesCount"`

	// 是否对身份证照片进行裁边。默认为false。（InfoType需要包含2）
	IsCutIdCardImage *bool `json:"IsCutIdCardImage,omitnil,omitempty" name:"IsCutIdCardImage"`

	// 是否需要从身份证中抠出头像。默认为false。（InfoType需要包含2）
	IsNeedIdCardAvatar *bool `json:"IsNeedIdCardAvatar,omitnil,omitempty" name:"IsNeedIdCardAvatar"`

	// 已弃用。
	IsEncrypt *bool `json:"IsEncrypt,omitnil,omitempty" name:"IsEncrypt"`

	// 是否需要对返回中的敏感信息进行加密。仅指定加密算法Algorithm即可，其余字段传入默认值。其中敏感信息包括：Response.Text.IdCard、Response.Text.Name、Response.Text.OcrIdCard、Response.Text.OcrName
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`

	// 是否对回包整体进行加密
	IsEncryptResponse *bool `json:"IsEncryptResponse,omitnil,omitempty" name:"IsEncryptResponse"`
}

type GetDetectInfoEnhancedRequest struct {
	*tchttp.BaseRequest
	
	// 人脸核身流程的标识，调用DetectAuth接口时生成。
	BizToken *string `json:"BizToken,omitnil,omitempty" name:"BizToken"`

	// 用于细分客户使用场景，由腾讯侧在线下对接时分配。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 指定拉取的结果信息，取值（0：全部；1：文本类；2：身份证信息；3：视频最佳截图信息）。
	// 如 13表示拉取文本类、视频最佳截图信息。
	// 默认值：0
	InfoType *string `json:"InfoType,omitnil,omitempty" name:"InfoType"`

	// 从活体视频中截取一定张数的最佳帧（仅部分服务支持，若需使用请与慧眼小助手沟通）。默认为0，最大为10，超出10的最多只给10张。（InfoType需要包含3）
	BestFramesCount *uint64 `json:"BestFramesCount,omitnil,omitempty" name:"BestFramesCount"`

	// 是否对身份证照片进行裁边。默认为false。（InfoType需要包含2）
	IsCutIdCardImage *bool `json:"IsCutIdCardImage,omitnil,omitempty" name:"IsCutIdCardImage"`

	// 是否需要从身份证中抠出头像。默认为false。（InfoType需要包含2）
	IsNeedIdCardAvatar *bool `json:"IsNeedIdCardAvatar,omitnil,omitempty" name:"IsNeedIdCardAvatar"`

	// 已弃用。
	IsEncrypt *bool `json:"IsEncrypt,omitnil,omitempty" name:"IsEncrypt"`

	// 是否需要对返回中的敏感信息进行加密。仅指定加密算法Algorithm即可，其余字段传入默认值。其中敏感信息包括：Response.Text.IdCard、Response.Text.Name、Response.Text.OcrIdCard、Response.Text.OcrName
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`

	// 是否对回包整体进行加密
	IsEncryptResponse *bool `json:"IsEncryptResponse,omitnil,omitempty" name:"IsEncryptResponse"`
}

func (r *GetDetectInfoEnhancedRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDetectInfoEnhancedRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizToken")
	delete(f, "RuleId")
	delete(f, "InfoType")
	delete(f, "BestFramesCount")
	delete(f, "IsCutIdCardImage")
	delete(f, "IsNeedIdCardAvatar")
	delete(f, "IsEncrypt")
	delete(f, "Encryption")
	delete(f, "IsEncryptResponse")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDetectInfoEnhancedRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDetectInfoEnhancedResponseParams struct {
	// 文本类信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *DetectInfoText `json:"Text,omitnil,omitempty" name:"Text"`

	// 身份证照片信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdCardData *DetectInfoIdCardData `json:"IdCardData,omitnil,omitempty" name:"IdCardData"`

	// 最佳帧信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BestFrame *DetectInfoBestFrame `json:"BestFrame,omitnil,omitempty" name:"BestFrame"`

	// 视频信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoData *DetectInfoVideoData `json:"VideoData,omitnil,omitempty" name:"VideoData"`

	// 敏感数据加密信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`

	// 意愿核身朗读模式结果信息。若未使用意愿核身功能，该字段返回值可以不处理。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntentionVerifyData *IntentionVerifyData `json:"IntentionVerifyData,omitnil,omitempty" name:"IntentionVerifyData"`

	// 意愿核身问答模式结果。若未使用该意愿核身功能，该字段返回值可以不处理。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntentionQuestionResult *IntentionQuestionResult `json:"IntentionQuestionResult,omitnil,omitempty" name:"IntentionQuestionResult"`

	// 意愿核身点头确认模式的结果信息，若未使用该意愿核身功能，该字段返回值可以不处理。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntentionActionResult *IntentionActionResult `json:"IntentionActionResult,omitnil,omitempty" name:"IntentionActionResult"`

	// 加密后的数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	EncryptedBody *string `json:"EncryptedBody,omitnil,omitempty" name:"EncryptedBody"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetDetectInfoEnhancedResponse struct {
	*tchttp.BaseResponse
	Response *GetDetectInfoEnhancedResponseParams `json:"Response"`
}

func (r *GetDetectInfoEnhancedResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDetectInfoEnhancedResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDetectInfoRequestParams struct {
	// 人脸核身流程的标识，调用DetectAuth接口时生成。
	BizToken *string `json:"BizToken,omitnil,omitempty" name:"BizToken"`

	// 用于细分客户使用场景，申请开通服务后，可以在腾讯云慧眼人脸核身控制台（https://console.cloud.tencent.com/faceid） 自助接入里面创建，审核通过后即可调用。如有疑问，请加慧眼小助手微信（faceid001）进行咨询。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 指定拉取的结果信息，取值（0：全部；1：文本类；2：身份证正反面；3：视频最佳截图照片；4：视频）。
	// 如 134表示拉取文本类、视频最佳截图照片、视频。
	// 默认值：0
	InfoType *string `json:"InfoType,omitnil,omitempty" name:"InfoType"`
}

type GetDetectInfoRequest struct {
	*tchttp.BaseRequest
	
	// 人脸核身流程的标识，调用DetectAuth接口时生成。
	BizToken *string `json:"BizToken,omitnil,omitempty" name:"BizToken"`

	// 用于细分客户使用场景，申请开通服务后，可以在腾讯云慧眼人脸核身控制台（https://console.cloud.tencent.com/faceid） 自助接入里面创建，审核通过后即可调用。如有疑问，请加慧眼小助手微信（faceid001）进行咨询。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 指定拉取的结果信息，取值（0：全部；1：文本类；2：身份证正反面；3：视频最佳截图照片；4：视频）。
	// 如 134表示拉取文本类、视频最佳截图照片、视频。
	// 默认值：0
	InfoType *string `json:"InfoType,omitnil,omitempty" name:"InfoType"`
}

func (r *GetDetectInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDetectInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizToken")
	delete(f, "RuleId")
	delete(f, "InfoType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDetectInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDetectInfoResponseParams struct {
	// JSON字符串。
	// {
	//   // 文本类信息
	//   "Text": {
	//     "ErrCode": null,      // 本次核身最终结果。0为成功
	//     "ErrMsg": null,       // 本次核身最终结果信息描述。
	//     "IdCard": "",         // 本次核身最终获得的身份证号。
	//     "Name": "",           // 本次核身最终获得的姓名。
	//     "OcrNation": null,    // ocr阶段获取的民族
	//     "OcrAddress": null,   // ocr阶段获取的地址
	//     "OcrBirth": null,     // ocr阶段获取的出生信息
	//     "OcrAuthority": null, // ocr阶段获取的证件签发机关
	//     "OcrValidDate": null, // ocr阶段获取的证件有效期
	//     "OcrName": null,      // ocr阶段获取的姓名
	//     "OcrIdCard": null,    // ocr阶段获取的身份证号
	//     "OcrGender": null,    // ocr阶段获取的性别
	//     "LiveStatus": null,   // 活体检测阶段的错误码。0为成功
	//     "LiveMsg": null,      // 活体检测阶段的错误信息
	//     "Comparestatus": null,// 一比一阶段的错误码。0为成功
	//     "Comparemsg": null,   // 一比一阶段的错误信息
	//     "Sim": null, // 比对相似度
	//     "Location": null, // 地理位置信息
	//     "Extra": "",          // DetectAuth结果传进来的Extra信息
	//     "Detail": {           // 活体一比一信息详情
	//       "LivenessData": [
	//             {
	//               ErrCode: null, // 活体比对验证错误码
	//               ErrMsg: null, // 活体比对验证错误描述
	//               ReqTime: null, // 活体验证时间戳
	//               IdCard: null, // 验证身份证号
	//               Name: null // 验证姓名
	//             }
	//       ]
	//     }
	//   },
	//   // 身份证正反面照片Base64
	//   "IdCardData": {
	//     "OcrFront": null,
	//     "OcrBack": null
	//   },
	//   // 视频最佳帧截图Base64
	//   "BestFrame": {
	//     "BestFrame": null
	//   },
	//   // 活体视频Base64
	//   "VideoData": {
	//     "LivenessVideo": null
	//   }
	// }
	DetectInfo *string `json:"DetectInfo,omitnil,omitempty" name:"DetectInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetDetectInfoResponse struct {
	*tchttp.BaseResponse
	Response *GetDetectInfoResponseParams `json:"Response"`
}

func (r *GetDetectInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDetectInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetEidResultRequestParams struct {
	// E证通流程的唯一标识，调用GetEidToken接口时生成。
	EidToken *string `json:"EidToken,omitnil,omitempty" name:"EidToken"`

	// 指定拉取的结果信息，取值（0：全部；1：文本类；2：身份证信息；3：最佳截图信息；5：意愿核身朗读模式相关结果；6：意愿核身问答模式相关结果）。
	// 如 13表示拉取文本类、最佳截图信息。
	// 默认值：0
	InfoType *string `json:"InfoType,omitnil,omitempty" name:"InfoType"`

	// 从活体视频中截取一定张数的最佳帧。默认为0，最大为3，超出3的最多只给3张。（InfoType需要包含3）
	BestFramesCount *uint64 `json:"BestFramesCount,omitnil,omitempty" name:"BestFramesCount"`

	// 是否对身份证照片进行裁边。默认为false。（InfoType需要包含2）
	IsCutIdCardImage *bool `json:"IsCutIdCardImage,omitnil,omitempty" name:"IsCutIdCardImage"`

	// 是否需要从身份证中抠出头像。默认为false。（InfoType需要包含2）
	IsNeedIdCardAvatar *bool `json:"IsNeedIdCardAvatar,omitnil,omitempty" name:"IsNeedIdCardAvatar"`
}

type GetEidResultRequest struct {
	*tchttp.BaseRequest
	
	// E证通流程的唯一标识，调用GetEidToken接口时生成。
	EidToken *string `json:"EidToken,omitnil,omitempty" name:"EidToken"`

	// 指定拉取的结果信息，取值（0：全部；1：文本类；2：身份证信息；3：最佳截图信息；5：意愿核身朗读模式相关结果；6：意愿核身问答模式相关结果）。
	// 如 13表示拉取文本类、最佳截图信息。
	// 默认值：0
	InfoType *string `json:"InfoType,omitnil,omitempty" name:"InfoType"`

	// 从活体视频中截取一定张数的最佳帧。默认为0，最大为3，超出3的最多只给3张。（InfoType需要包含3）
	BestFramesCount *uint64 `json:"BestFramesCount,omitnil,omitempty" name:"BestFramesCount"`

	// 是否对身份证照片进行裁边。默认为false。（InfoType需要包含2）
	IsCutIdCardImage *bool `json:"IsCutIdCardImage,omitnil,omitempty" name:"IsCutIdCardImage"`

	// 是否需要从身份证中抠出头像。默认为false。（InfoType需要包含2）
	IsNeedIdCardAvatar *bool `json:"IsNeedIdCardAvatar,omitnil,omitempty" name:"IsNeedIdCardAvatar"`
}

func (r *GetEidResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetEidResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EidToken")
	delete(f, "InfoType")
	delete(f, "BestFramesCount")
	delete(f, "IsCutIdCardImage")
	delete(f, "IsNeedIdCardAvatar")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetEidResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetEidResultResponseParams struct {
	// 文本类信息。（基于对敏感信息的保护，验证使用的姓名和身份证号统一通过加密后从Eidinfo参数中返回，如需获取请在控制台申请返回身份信息，详见[E证通获取实名信息指引](https://cloud.tencent.com/document/product/1007/63370)）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *DetectInfoText `json:"Text,omitnil,omitempty" name:"Text"`

	// 身份证照片信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdCardData *DetectInfoIdCardData `json:"IdCardData,omitnil,omitempty" name:"IdCardData"`

	// 最佳帧信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BestFrame *DetectInfoBestFrame `json:"BestFrame,omitnil,omitempty" name:"BestFrame"`

	// Eid信息。（包括商户下用户唯一标识以及加密后的姓名、身份证号信息。解密方式详见[E证通获取实名信息指引](https://cloud.tencent.com/document/product/1007/63370)）
	// 注意：此字段可能返回 null，表示取不到有效值。
	EidInfo *EidInfo `json:"EidInfo,omitnil,omitempty" name:"EidInfo"`

	// 意愿核身朗读模式相关信息。若未使用意愿核身朗读功能，该字段返回值可以不处理。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntentionVerifyData *IntentionVerifyData `json:"IntentionVerifyData,omitnil,omitempty" name:"IntentionVerifyData"`

	// 意愿核身问答模式相关信息。若未使用意愿核身问答模式功能，该字段返回值可以不处理。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntentionQuestionResult *IntentionQuestionResult `json:"IntentionQuestionResult,omitnil,omitempty" name:"IntentionQuestionResult"`

	// 意愿核身点头确认模式的结果信息，若未使用该意愿核身功能，该字段返回值可以不处理。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntentionActionResult *IntentionActionResult `json:"IntentionActionResult,omitnil,omitempty" name:"IntentionActionResult"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetEidResultResponse struct {
	*tchttp.BaseResponse
	Response *GetEidResultResponseParams `json:"Response"`
}

func (r *GetEidResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetEidResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetEidTokenConfig struct {
	// 姓名身份证输入方式。
	// 1：传身份证正反面OCR   
	// 2：传身份证正面OCR  
	// 3：用户手动输入  
	// 4：客户后台传入  
	// 默认1
	// 注：使用OCR时仅支持用户修改结果中的姓名
	InputType *string `json:"InputType,omitnil,omitempty" name:"InputType"`

	// 是否使用意愿核身，默认不使用。注意：如开启使用，则计费标签按【意愿核身】计费标签计价；如不开启，则计费标签按【E证通】计费标签计价，价格详见：[价格说明](https://cloud.tencent.com/document/product/1007/56804)。
	UseIntentionVerify *bool `json:"UseIntentionVerify,omitnil,omitempty" name:"UseIntentionVerify"`

	// 意愿核身模式。枚举值：1( 语音朗读模式)，2（语音问答模式） ，3（点头确认模式）。默认值为1。
	IntentionMode *string `json:"IntentionMode,omitnil,omitempty" name:"IntentionMode"`

	// 意愿核身朗读模式使用的文案，若未使用意愿核身朗读功能，该字段无需传入。默认为空，最长可接受120的字符串长度。
	IntentionVerifyText *string `json:"IntentionVerifyText,omitnil,omitempty" name:"IntentionVerifyText"`

	// 意愿核身问答模式的配置列表。当前仅支持一个问答。
	IntentionQuestions []*IntentionQuestion `json:"IntentionQuestions,omitnil,omitempty" name:"IntentionQuestions"`

	// 意愿核身（点头确认模式）使用的文案，若未使用意愿核身（点头确认模式），则该字段无需传入。默认为空，最长可接受150的字符串长度。
	IntentionActions []*IntentionActionConfig `json:"IntentionActions,omitnil,omitempty" name:"IntentionActions"`

	// 意愿核身过程中识别用户的回答意图，开启后除了IntentionQuestions的Answers列表中的标准回答会通过，近似意图的回答也会通过，默认开启。
	IntentionRecognition *bool `json:"IntentionRecognition,omitnil,omitempty" name:"IntentionRecognition"`

	// 是否支持港澳台居住证识别
	IsSupportHMTResidentPermitOCR *bool `json:"IsSupportHMTResidentPermitOCR,omitnil,omitempty" name:"IsSupportHMTResidentPermitOCR"`

	// 用户语音回答过程中是否开启张嘴识别检测，默认不开启，仅在意愿核身问答模式中使用。
	MouthOpenRecognition *bool `json:"MouthOpenRecognition,omitnil,omitempty" name:"MouthOpenRecognition"`

	// 意愿核身语音播报速度，配置后问答模式和点头模式的语音播报环节都会生效，默认值为0： 0：智能语速（根据播报文案的长度自动调整语音播报速度） 1：固定1倍速 2：固定1.2倍速 3：固定1.5倍速
	Speed *uint64 `json:"Speed,omitnil,omitempty" name:"Speed"`
}

// Predefined struct for user
type GetEidTokenRequestParams struct {
	// EID商户id，字段长度最长50位。
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`

	// 身份标识（未使用OCR服务时，必须传入）。
	// 规则：a-z，A-Z，0-9组合。最长长度32位。
	IdCard *string `json:"IdCard,omitnil,omitempty" name:"IdCard"`

	// 姓名。（未使用OCR服务时，必须传入）最长长度32位。中文请使用UTF-8编码。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 透传字段，在获取验证结果时返回。最长长度1024位。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 小程序模式配置，包括如何传入姓名身份证的配置，以及是否使用意愿核身。
	Config *GetEidTokenConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// 最长长度1024位。用户从Url中进入核身认证结束后重定向的回调链接地址。EidToken会在该链接的query参数中。
	RedirectUrl *string `json:"RedirectUrl,omitnil,omitempty" name:"RedirectUrl"`

	// 敏感数据加密信息。对传入信息（姓名、身份证号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

type GetEidTokenRequest struct {
	*tchttp.BaseRequest
	
	// EID商户id，字段长度最长50位。
	MerchantId *string `json:"MerchantId,omitnil,omitempty" name:"MerchantId"`

	// 身份标识（未使用OCR服务时，必须传入）。
	// 规则：a-z，A-Z，0-9组合。最长长度32位。
	IdCard *string `json:"IdCard,omitnil,omitempty" name:"IdCard"`

	// 姓名。（未使用OCR服务时，必须传入）最长长度32位。中文请使用UTF-8编码。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 透传字段，在获取验证结果时返回。最长长度1024位。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 小程序模式配置，包括如何传入姓名身份证的配置，以及是否使用意愿核身。
	Config *GetEidTokenConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// 最长长度1024位。用户从Url中进入核身认证结束后重定向的回调链接地址。EidToken会在该链接的query参数中。
	RedirectUrl *string `json:"RedirectUrl,omitnil,omitempty" name:"RedirectUrl"`

	// 敏感数据加密信息。对传入信息（姓名、身份证号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

func (r *GetEidTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetEidTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MerchantId")
	delete(f, "IdCard")
	delete(f, "Name")
	delete(f, "Extra")
	delete(f, "Config")
	delete(f, "RedirectUrl")
	delete(f, "Encryption")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetEidTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetEidTokenResponseParams struct {
	// 一次核身流程的标识，有效时间为600秒；
	// 完成核身后，可用该标识获取验证结果信息。
	EidToken *string `json:"EidToken,omitnil,omitempty" name:"EidToken"`

	// 发起核身流程的URL，用于H5场景核身。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetEidTokenResponse struct {
	*tchttp.BaseResponse
	Response *GetEidTokenResponseParams `json:"Response"`
}

func (r *GetEidTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetEidTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetFaceIdResultRequestParams struct {
	// SDK人脸核身流程的标识，调用GetFaceIdToken接口时生成。
	FaceIdToken *string `json:"FaceIdToken,omitnil,omitempty" name:"FaceIdToken"`

	// 是否需要拉取视频，默认false不需要
	IsNeedVideo *bool `json:"IsNeedVideo,omitnil,omitempty" name:"IsNeedVideo"`

	// 是否需要拉取截帧，默认false不需要
	IsNeedBestFrame *bool `json:"IsNeedBestFrame,omitnil,omitempty" name:"IsNeedBestFrame"`
}

type GetFaceIdResultRequest struct {
	*tchttp.BaseRequest
	
	// SDK人脸核身流程的标识，调用GetFaceIdToken接口时生成。
	FaceIdToken *string `json:"FaceIdToken,omitnil,omitempty" name:"FaceIdToken"`

	// 是否需要拉取视频，默认false不需要
	IsNeedVideo *bool `json:"IsNeedVideo,omitnil,omitempty" name:"IsNeedVideo"`

	// 是否需要拉取截帧，默认false不需要
	IsNeedBestFrame *bool `json:"IsNeedBestFrame,omitnil,omitempty" name:"IsNeedBestFrame"`
}

func (r *GetFaceIdResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFaceIdResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FaceIdToken")
	delete(f, "IsNeedVideo")
	delete(f, "IsNeedBestFrame")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetFaceIdResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetFaceIdResultResponseParams struct {
	// 身份证
	IdCard *string `json:"IdCard,omitnil,omitempty" name:"IdCard"`

	// 姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 业务核验结果，参考https://cloud.tencent.com/document/product/1007/47912
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 业务核验描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 相似度，0-100，数值越大相似度越高
	Similarity *float64 `json:"Similarity,omitnil,omitempty" name:"Similarity"`

	// 用户核验的视频base64，如果选择了使用cos，返回完整cos地址如https://bucket.cos.ap-guangzhou.myqcloud.com/objectKey
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoBase64 *string `json:"VideoBase64,omitnil,omitempty" name:"VideoBase64"`

	// 用户核验视频的截帧base64，如果选择了使用cos，返回完整cos地址如https://bucket.cos.ap-guangzhou.myqcloud.com/objectKey
	// 注意：此字段可能返回 null，表示取不到有效值。
	BestFrameBase64 *string `json:"BestFrameBase64,omitnil,omitempty" name:"BestFrameBase64"`

	// 获取token时透传的信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// plus版：描述当前请求所在设备的风险标签，详情如下：
	// 01-设备疑似被Root/设备疑似越狱
	// 02-设备疑似被注入
	// 03-设备疑似为模拟器
	// 04-设备疑似存在风险操作
	// 05-摄像头疑似被劫持
	// 06-疑似黑产设备
	// null-无设备风险
	// 增强版：此字段不生效，默认为null
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceInfoTag *string `json:"DeviceInfoTag,omitnil,omitempty" name:"DeviceInfoTag"`

	// 行为风险标签，仅错误码返回1007（设备疑似被劫持）时返回风险标签。标签说明：
	// 02：攻击风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskInfoTag *string `json:"RiskInfoTag,omitnil,omitempty" name:"RiskInfoTag"`

	// plus版：描述当前请求活体阶段被拒绝的详细原因，详情如下：01-用户全程闭眼02-用户未完成指定动作03-疑似翻拍攻击04-疑似合成图片05-疑似合成视频06-疑似合成动作07-疑似黑产模板08-疑似存在水印09-反光校验未通过10-最佳帧校验未通过11-人脸质量过差12-人脸距离不匹配13-疑似对抗样本攻击null-无增强版：此字段不生效，默认为null
	// 注意：此字段可能返回 null，表示取不到有效值。
	LivenessInfoTag *string `json:"LivenessInfoTag,omitnil,omitempty" name:"LivenessInfoTag"`

	// plus版：描述当前请求所在设备的风险等级，共4级，详情如下：
	// 1 - 安全
	// 2 - 低风险
	// 3 - 中风险
	// 4 - 高危
	// null - 未获取到风险等级
	// 增强版：此字段不生效，默认为null
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceInfoLevel *string `json:"DeviceInfoLevel,omitnil,omitempty" name:"DeviceInfoLevel"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetFaceIdResultResponse struct {
	*tchttp.BaseResponse
	Response *GetFaceIdResultResponseParams `json:"Response"`
}

func (r *GetFaceIdResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFaceIdResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetFaceIdRiskInfoRequestParams struct {
	// SDK人脸核身流程的标识，调用GetFaceidRiskInfoToken接口时生成。
	FaceIdToken *string `json:"FaceIdToken,omitnil,omitempty" name:"FaceIdToken"`
}

type GetFaceIdRiskInfoRequest struct {
	*tchttp.BaseRequest
	
	// SDK人脸核身流程的标识，调用GetFaceidRiskInfoToken接口时生成。
	FaceIdToken *string `json:"FaceIdToken,omitnil,omitempty" name:"FaceIdToken"`
}

func (r *GetFaceIdRiskInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFaceIdRiskInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FaceIdToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetFaceIdRiskInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetFaceIdRiskInfoResponseParams struct {
	// 描述当前请求所在设备的风险标签，详情如下： 01-设备疑似被Root/设备疑似越狱 02-设备疑似被注入 03-设备疑似为模拟器 04-设备疑似存在风险操作 05-摄像头疑似被劫持 06-疑似黑产设备
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceInfoTag *string `json:"DeviceInfoTag,omitnil,omitempty" name:"DeviceInfoTag"`

	// 描述当前请求所在设备的风险等级，共4级，详情如下： 1 - 低风险 2 - 中风险 3 - 高风险 4 - 攻击 ，-1表示未获取到风险等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceInfoLevel *int64 `json:"DeviceInfoLevel,omitnil,omitempty" name:"DeviceInfoLevel"`

	// 设备id标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpenId *string `json:"OpenId,omitnil,omitempty" name:"OpenId"`

	// 描述当前请求所在设备的相机指纹风险等级，共4级，详情如下： 1 - 低风险 2 - 中风险 3 - 高风险 4 - 攻击 ，-1表示未获取到风险等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	CameraInfoLevel *int64 `json:"CameraInfoLevel,omitnil,omitempty" name:"CameraInfoLevel"`

	// 描述当前请求所在设备的相机指纹风险标签，详情如下： 01-设备疑似被Root/设备疑似越狱 02-设备疑似被注入 03-设备疑似为模拟器 04-设备疑似存在风险操作 05-摄像头疑似被劫持 06-疑似黑产设备，空表示没有相机指纹风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	CameraInfoTag *string `json:"CameraInfoTag,omitnil,omitempty" name:"CameraInfoTag"`

	// 获取token时透传的信息	
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetFaceIdRiskInfoResponse struct {
	*tchttp.BaseResponse
	Response *GetFaceIdRiskInfoResponseParams `json:"Response"`
}

func (r *GetFaceIdRiskInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFaceIdRiskInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetFaceIdTokenRequestParams struct {
	// 本地上传照片(LOCAL)、商业库(BUSINESS)
	CompareLib *string `json:"CompareLib,omitnil,omitempty" name:"CompareLib"`

	// CompareLib为商业库时必传。
	IdCard *string `json:"IdCard,omitnil,omitempty" name:"IdCard"`

	// CompareLib为商业库时必传。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// CompareLib为上传照片比对时必传，Base64后图片最大8MB。
	// 请使用标准的Base64编码方式(带=补位)，编码规范参考RFC4648。
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// SDK中生成的Meta字符串
	Meta *string `json:"Meta,omitnil,omitempty" name:"Meta"`

	// 透传参数 1000长度字符串
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 默认为false，设置该参数为true后，核身过程中的视频图片将会存储在人脸核身控制台授权cos的bucket中，拉取结果时会返回对应资源完整cos地址。开通地址见https://console.cloud.tencent.com/faceid/cos
	// 【注意】选择该参数为true后将不返回base64数据，请根据接入情况谨慎修改。
	UseCos *bool `json:"UseCos,omitnil,omitempty" name:"UseCos"`

	// 敏感数据加密信息。对传入信息（姓名、身份证号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`

	// 用于细分客户使用场景，申请开通服务后，可以在腾讯云慧眼人脸核身控制台（https://console.cloud.tencent.com/faceid） 自助接入里面创建，审核通过后即可调用。如有疑问，请添加腾讯云人脸核身小助手进行咨询。
	// 示例值：1
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

type GetFaceIdTokenRequest struct {
	*tchttp.BaseRequest
	
	// 本地上传照片(LOCAL)、商业库(BUSINESS)
	CompareLib *string `json:"CompareLib,omitnil,omitempty" name:"CompareLib"`

	// CompareLib为商业库时必传。
	IdCard *string `json:"IdCard,omitnil,omitempty" name:"IdCard"`

	// CompareLib为商业库时必传。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// CompareLib为上传照片比对时必传，Base64后图片最大8MB。
	// 请使用标准的Base64编码方式(带=补位)，编码规范参考RFC4648。
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// SDK中生成的Meta字符串
	Meta *string `json:"Meta,omitnil,omitempty" name:"Meta"`

	// 透传参数 1000长度字符串
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// 默认为false，设置该参数为true后，核身过程中的视频图片将会存储在人脸核身控制台授权cos的bucket中，拉取结果时会返回对应资源完整cos地址。开通地址见https://console.cloud.tencent.com/faceid/cos
	// 【注意】选择该参数为true后将不返回base64数据，请根据接入情况谨慎修改。
	UseCos *bool `json:"UseCos,omitnil,omitempty" name:"UseCos"`

	// 敏感数据加密信息。对传入信息（姓名、身份证号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`

	// 用于细分客户使用场景，申请开通服务后，可以在腾讯云慧眼人脸核身控制台（https://console.cloud.tencent.com/faceid） 自助接入里面创建，审核通过后即可调用。如有疑问，请添加腾讯云人脸核身小助手进行咨询。
	// 示例值：1
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

func (r *GetFaceIdTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFaceIdTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CompareLib")
	delete(f, "IdCard")
	delete(f, "Name")
	delete(f, "ImageBase64")
	delete(f, "Meta")
	delete(f, "Extra")
	delete(f, "UseCos")
	delete(f, "Encryption")
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetFaceIdTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetFaceIdTokenResponseParams struct {
	// 有效期 10分钟。只能完成1次核身。
	FaceIdToken *string `json:"FaceIdToken,omitnil,omitempty" name:"FaceIdToken"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetFaceIdTokenResponse struct {
	*tchttp.BaseResponse
	Response *GetFaceIdTokenResponseParams `json:"Response"`
}

func (r *GetFaceIdTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFaceIdTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetFaceidRiskInfoTokenRequestParams struct {
	// SDK中生成的Meta字符串
	Meta *string `json:"Meta,omitnil,omitempty" name:"Meta"`

	// 透传参数 1000长度字符串
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`
}

type GetFaceidRiskInfoTokenRequest struct {
	*tchttp.BaseRequest
	
	// SDK中生成的Meta字符串
	Meta *string `json:"Meta,omitnil,omitempty" name:"Meta"`

	// 透传参数 1000长度字符串
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`
}

func (r *GetFaceidRiskInfoTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFaceidRiskInfoTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Meta")
	delete(f, "Extra")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetFaceidRiskInfoTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetFaceidRiskInfoTokenResponseParams struct {
	// 有效期 10分钟。只能完成1次核身。
	FaceIdToken *string `json:"FaceIdToken,omitnil,omitempty" name:"FaceIdToken"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetFaceidRiskInfoTokenResponse struct {
	*tchttp.BaseResponse
	Response *GetFaceidRiskInfoTokenResponseParams `json:"Response"`
}

func (r *GetFaceidRiskInfoTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFaceidRiskInfoTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetLiveCodeRequestParams struct {

}

type GetLiveCodeRequest struct {
	*tchttp.BaseRequest
	
}

func (r *GetLiveCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetLiveCodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetLiveCodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetLiveCodeResponseParams struct {
	// 数字验证码，如：1234
	LiveCode *string `json:"LiveCode,omitnil,omitempty" name:"LiveCode"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetLiveCodeResponse struct {
	*tchttp.BaseResponse
	Response *GetLiveCodeResponseParams `json:"Response"`
}

func (r *GetLiveCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetLiveCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetWeChatBillDetailsRequestParams struct {
	// 拉取的日期（YYYY-MM-DD）。最大可追溯到365天前。当天6点后才能拉取前一天的数据。
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// 游标。用于分页，取第一页时传0，取后续页面时，传入本接口响应中返回的NextCursor字段的值。
	Cursor *uint64 `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 需要拉取账单详情业务对应的RuleId。不传会返回所有RuleId数据。默认为空字符串。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

type GetWeChatBillDetailsRequest struct {
	*tchttp.BaseRequest
	
	// 拉取的日期（YYYY-MM-DD）。最大可追溯到365天前。当天6点后才能拉取前一天的数据。
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// 游标。用于分页，取第一页时传0，取后续页面时，传入本接口响应中返回的NextCursor字段的值。
	Cursor *uint64 `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// 需要拉取账单详情业务对应的RuleId。不传会返回所有RuleId数据。默认为空字符串。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

func (r *GetWeChatBillDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetWeChatBillDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Date")
	delete(f, "Cursor")
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetWeChatBillDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetWeChatBillDetailsResponseParams struct {
	// 是否还有下一页。该字段为true时，需要将NextCursor的值作为入参Cursor继续调用本接口。
	HasNextPage *bool `json:"HasNextPage,omitnil,omitempty" name:"HasNextPage"`

	// 下一页的游标。用于分页。
	NextCursor *uint64 `json:"NextCursor,omitnil,omitempty" name:"NextCursor"`

	// 数据
	WeChatBillDetails []*WeChatBillDetail `json:"WeChatBillDetails,omitnil,omitempty" name:"WeChatBillDetails"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetWeChatBillDetailsResponse struct {
	*tchttp.BaseResponse
	Response *GetWeChatBillDetailsResponseParams `json:"Response"`
}

func (r *GetWeChatBillDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetWeChatBillDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IdCardOCRVerificationRequestParams struct {
	// 身份证号
	// 姓名和身份证号、ImageBase64、ImageUrl三者必须提供其中之一。若都提供了，则按照姓名和身份证号>ImageBase64>ImageUrl的优先级使用参数。
	IdCard *string `json:"IdCard,omitnil,omitempty" name:"IdCard"`

	// 姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 身份证人像面的 Base64 值
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 3M。请使用标准的Base64编码方式(带=补位)，编码规范参考RFC4648。
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// 身份证人像面的 Url 地址
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 3M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 敏感数据加密信息。对传入信息（姓名、身份证号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

type IdCardOCRVerificationRequest struct {
	*tchttp.BaseRequest
	
	// 身份证号
	// 姓名和身份证号、ImageBase64、ImageUrl三者必须提供其中之一。若都提供了，则按照姓名和身份证号>ImageBase64>ImageUrl的优先级使用参数。
	IdCard *string `json:"IdCard,omitnil,omitempty" name:"IdCard"`

	// 姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 身份证人像面的 Base64 值
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 3M。请使用标准的Base64编码方式(带=补位)，编码规范参考RFC4648。
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// 身份证人像面的 Url 地址
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 3M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 敏感数据加密信息。对传入信息（姓名、身份证号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

func (r *IdCardOCRVerificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IdCardOCRVerificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdCard")
	delete(f, "Name")
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "Encryption")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IdCardOCRVerificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IdCardOCRVerificationResponseParams struct {
	// 认证结果码，收费情况如下。
	// 收费结果码：
	// 0: 姓名和身份证号一致
	// -1: 姓名和身份证号不一致
	// 不收费结果码：
	// -2: 非法身份证号（长度、校验位等不正确）
	// -3: 非法姓名（长度、格式等不正确）
	// -4: 证件库服务异常
	// -5: 证件库中无此身份证记录
	// -6: 权威比对系统升级中，请稍后再试
	// -7: 认证次数超过当日限制
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 业务结果描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 用于验证的姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用于验证的身份证号
	IdCard *string `json:"IdCard,omitnil,omitempty" name:"IdCard"`

	// OCR得到的性别
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sex *string `json:"Sex,omitnil,omitempty" name:"Sex"`

	// OCR得到的民族
	// 注意：此字段可能返回 null，表示取不到有效值。
	Nation *string `json:"Nation,omitnil,omitempty" name:"Nation"`

	// OCR得到的生日
	// 注意：此字段可能返回 null，表示取不到有效值。
	Birth *string `json:"Birth,omitnil,omitempty" name:"Birth"`

	// OCR得到的地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type IdCardOCRVerificationResponse struct {
	*tchttp.BaseResponse
	Response *IdCardOCRVerificationResponseParams `json:"Response"`
}

func (r *IdCardOCRVerificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IdCardOCRVerificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IdCardVerificationRequestParams struct {
	// 身份证号
	IdCard *string `json:"IdCard,omitnil,omitempty" name:"IdCard"`

	// 姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 敏感数据加密信息。对传入信息（姓名、身份证号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

type IdCardVerificationRequest struct {
	*tchttp.BaseRequest
	
	// 身份证号
	IdCard *string `json:"IdCard,omitnil,omitempty" name:"IdCard"`

	// 姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 敏感数据加密信息。对传入信息（姓名、身份证号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

func (r *IdCardVerificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IdCardVerificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdCard")
	delete(f, "Name")
	delete(f, "Encryption")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IdCardVerificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IdCardVerificationResponseParams struct {
	// 认证结果码，收费情况如下。
	// 收费结果码：
	// 0: 姓名和身份证号一致
	// -1: 姓名和身份证号不一致
	// 不收费结果码：
	// -2: 非法身份证号（长度、校验位等不正确）
	// -3: 非法姓名（长度、格式等不正确）
	// -4: 证件库服务异常
	// -5: 证件库中无此身份证记录
	// -6: 权威比对系统升级中，请稍后再试
	// -7: 认证次数超过当日限制
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 业务结果描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type IdCardVerificationResponse struct {
	*tchttp.BaseResponse
	Response *IdCardVerificationResponseParams `json:"Response"`
}

func (r *IdCardVerificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IdCardVerificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImageRecognitionRequestParams struct {
	// 身份证号
	IdCard *string `json:"IdCard,omitnil,omitempty" name:"IdCard"`

	// 姓名。中文请使用UTF-8编码。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用于人脸比对的照片，图片的Base64值；
	// Base64编码后的图片数据大小不超过3M，仅支持jpg、png格式。
	// 请使用标准的Base64编码方式(带=补位)，编码规范参考RFC4648。
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// 本接口不需要传递此参数。
	Optional *string `json:"Optional,omitnil,omitempty" name:"Optional"`

	// 敏感数据加密信息。对传入信息（姓名、身份证号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

type ImageRecognitionRequest struct {
	*tchttp.BaseRequest
	
	// 身份证号
	IdCard *string `json:"IdCard,omitnil,omitempty" name:"IdCard"`

	// 姓名。中文请使用UTF-8编码。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用于人脸比对的照片，图片的Base64值；
	// Base64编码后的图片数据大小不超过3M，仅支持jpg、png格式。
	// 请使用标准的Base64编码方式(带=补位)，编码规范参考RFC4648。
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// 本接口不需要传递此参数。
	Optional *string `json:"Optional,omitnil,omitempty" name:"Optional"`

	// 敏感数据加密信息。对传入信息（姓名、身份证号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

func (r *ImageRecognitionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageRecognitionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdCard")
	delete(f, "Name")
	delete(f, "ImageBase64")
	delete(f, "Optional")
	delete(f, "Encryption")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImageRecognitionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImageRecognitionResponseParams struct {
	// 相似度，取值范围 [0.00, 100.00]。推荐相似度大于等于70时可判断为同一人，可根据具体场景自行调整阈值（阈值70的误通过率为千分之一，阈值80的误通过率是万分之一）
	Sim *float64 `json:"Sim,omitnil,omitempty" name:"Sim"`

	// 业务错误码，成功情况返回Success, 错误情况请参考下方错误码 列表中FailedOperation部分
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 业务结果描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ImageRecognitionResponse struct {
	*tchttp.BaseResponse
	Response *ImageRecognitionResponseParams `json:"Response"`
}

func (r *ImageRecognitionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageRecognitionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImageRecognitionV2RequestParams struct {
	// 身份证号
	IdCard *string `json:"IdCard,omitnil,omitempty" name:"IdCard"`

	// 姓名。中文请使用UTF-8编码。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用于人脸比对的照片，图片的Base64值；
	// Base64编码后的图片数据大小不超过3M，仅支持jpg、png格式。
	// 请使用标准的Base64编码方式(带=补位)，编码规范参考RFC4648。
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// 本接口不需要传递此参数。
	Optional *string `json:"Optional,omitnil,omitempty" name:"Optional"`

	// 敏感数据加密信息。对传入信息（姓名、身份证号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

type ImageRecognitionV2Request struct {
	*tchttp.BaseRequest
	
	// 身份证号
	IdCard *string `json:"IdCard,omitnil,omitempty" name:"IdCard"`

	// 姓名。中文请使用UTF-8编码。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用于人脸比对的照片，图片的Base64值；
	// Base64编码后的图片数据大小不超过3M，仅支持jpg、png格式。
	// 请使用标准的Base64编码方式(带=补位)，编码规范参考RFC4648。
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// 本接口不需要传递此参数。
	Optional *string `json:"Optional,omitnil,omitempty" name:"Optional"`

	// 敏感数据加密信息。对传入信息（姓名、身份证号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

func (r *ImageRecognitionV2Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageRecognitionV2Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdCard")
	delete(f, "Name")
	delete(f, "ImageBase64")
	delete(f, "Optional")
	delete(f, "Encryption")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImageRecognitionV2Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImageRecognitionV2ResponseParams struct {
	// 相似度，取值范围 [0.00, 100.00]。推荐相似度大于等于70时可判断为同一人，可根据具体场景自行调整阈值（阈值70的误通过率为千分之一，阈值80的误通过率是万分之一）
	Sim *float64 `json:"Sim,omitnil,omitempty" name:"Sim"`

	// 业务错误码，成功情况返回Success, 错误情况请参考下方错误码 列表中FailedOperation部分
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 业务结果描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ImageRecognitionV2Response struct {
	*tchttp.BaseResponse
	Response *ImageRecognitionV2ResponseParams `json:"Response"`
}

func (r *ImageRecognitionV2Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageRecognitionV2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IntentionActionConfig struct {
	// 点头确认模式下，系统语音播报使用的问题文本，问题最大长度为150个字符。
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type IntentionActionResult struct {
	// 意愿核身错误码：
	// 0: "成功"       
	// -1: "参数错误"    
	// -2: "系统异常"    
	// -101: "请保持人脸在框内"    
	// -102: "检测到多张人脸"   
	// -103: "人脸检测失败"   
	// -104: "人脸检测不完整"   
	// -105: "请勿遮挡眼睛"    
	// -106: "请勿遮挡嘴巴"     
	// -107: "请勿遮挡鼻子"     
	// -201: "人脸比对相似度低"    
	// -202: "人脸比对失败"    
	// -301: "意愿核验不通过"   
	// -800: "前端不兼容错误"    
	// -801: "用户未授权摄像头和麦克风权限"   
	// -802: "核验流程异常中断，请勿切屏或进行其他操作"   
	// -803: "用户主动关闭链接/异常断开链接"   
	// -998: "系统数据异常"   
	// -999: "系统未知错误，请联系人工核实"   
	// 若在人脸核身过程失败、未进入意愿确认过程，则该参数返回为空，请参考人脸核身错误码结果（DetectInfoText.ErrCode)
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinalResultDetailCode *int64 `json:"FinalResultDetailCode,omitnil,omitempty" name:"FinalResultDetailCode"`

	// 意愿核身错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinalResultMessage *string `json:"FinalResultMessage,omitnil,omitempty" name:"FinalResultMessage"`

	// 意愿核身结果详细数据，与每段点头确认过程一一对应
	// 注意：此字段可能返回 null，表示取不到有效值。
	Details []*IntentionActionResultDetail `json:"Details,omitnil,omitempty" name:"Details"`
}

type IntentionActionResultDetail struct {
	// 视频base64编码（其中包含全程提示文本和点头音频，mp4格式）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Video *string `json:"Video,omitnil,omitempty" name:"Video"`

	// 屏幕截图base64编码列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScreenShot []*string `json:"ScreenShot,omitnil,omitempty" name:"ScreenShot"`
}

type IntentionQuestion struct {
	// 当选择语音问答模式时，系统自动播报的问题文本，最大长度为150个字符。
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 当选择语音问答模式时，用于判断用户回答是否通过的标准答案列表，传入后可自动判断用户回答文本是否在标准文本列表中。列表长度最大为50，单个答案长度限制10个字符。
	Answers []*string `json:"Answers,omitnil,omitempty" name:"Answers"`
}

type IntentionQuestionResult struct {
	// 意愿核身错误码：
	// 0: "成功"       
	// -1: "参数错误"    
	// -2: "系统异常"    
	// -101: "请保持人脸在框内"    
	// -102: "检测到多张人脸"   
	// -103: "人脸检测失败"   
	// -104: "人脸检测不完整"   
	// -105: "请勿遮挡眼睛"    
	// -106: "请勿遮挡嘴巴"     
	// -107: "请勿遮挡鼻子"     
	// -201: "人脸比对相似度低"    
	// -202: "人脸比对失败"    
	// -301: "意愿核验不通过"       
	// -302: "用户回答阶段未检测到张嘴动作"  
	// -800: "前端不兼容错误"    
	// -801: "用户未授权摄像头和麦克风权限"   
	// -802: "核验流程异常中断，请勿切屏或进行其他操作"   
	// -803: "用户主动关闭链接/异常断开链接"   
	// -998: "系统数据异常"   
	// -999: "系统未知错误，请联系人工核实"   
	// 若在人脸核身过程失败、未进入意愿确认过程，则该参数返回为空，请参考人脸核身错误码结果（DetectInfoText.ErrCode)
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinalResultDetailCode *int64 `json:"FinalResultDetailCode,omitnil,omitempty" name:"FinalResultDetailCode"`

	// 意愿核身错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinalResultMessage *string `json:"FinalResultMessage,omitnil,omitempty" name:"FinalResultMessage"`

	// 视频base64（其中包含全程问题和回答音频，mp4格式）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Video *string `json:"Video,omitnil,omitempty" name:"Video"`

	// 屏幕截图base64列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScreenShot []*string `json:"ScreenShot,omitnil,omitempty" name:"ScreenShot"`

	// 和答案匹配结果列表
	// 0：成功，-1：不匹配
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultCode []*string `json:"ResultCode,omitnil,omitempty" name:"ResultCode"`

	// 回答问题语音识别结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsrResult []*string `json:"AsrResult,omitnil,omitempty" name:"AsrResult"`

	// 答案录音音频
	// 注意：此字段可能返回 null，表示取不到有效值。
	Audios []*string `json:"Audios,omitnil,omitempty" name:"Audios"`

	// 意愿核身最终结果：
	// 0：认证通过，-1：认证未通过，-2：浏览器内核不兼容，无法进行意愿校验。建议使用“FinalResultDetailCode”参数获取详细的错误码信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinalResultCode *string `json:"FinalResultCode,omitnil,omitempty" name:"FinalResultCode"`
}

type IntentionVerifyData struct {
	// 意愿确认环节中录制的视频（base64）。若不存在则为空字符串。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntentionVerifyVideo *string `json:"IntentionVerifyVideo,omitnil,omitempty" name:"IntentionVerifyVideo"`

	// 意愿确认环节中用户语音转文字的识别结果。若不存在则为空字符串。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsrResult *string `json:"AsrResult,omitnil,omitempty" name:"AsrResult"`

	// 意愿确认环节的结果码。当该结果码为0时，语音朗读的视频与语音识别结果才会返回。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorCode *int64 `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// 意愿确认环节的结果信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 意愿确认环节中录制视频的最佳帧（base64）。若不存在则为空字符串。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntentionVerifyBestFrame *string `json:"IntentionVerifyBestFrame,omitnil,omitempty" name:"IntentionVerifyBestFrame"`

	// 本次流程用户语音与传入文本比对的相似度分值，取值范围 [0.00, 100.00]。只有配置了相似度阈值后才进行语音校验并返回相似度分值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: AsrResultSimilarity is deprecated.
	AsrResultSimilarity *string `json:"AsrResultSimilarity,omitnil,omitempty" name:"AsrResultSimilarity"`
}

// Predefined struct for user
type LivenessCompareRequestParams struct {
	// 活体检测类型，取值：LIP/ACTION/SILENT。
	// LIP为数字模式，ACTION为动作模式，SILENT为静默模式，三种模式选择一种传入。
	LivenessType *string `json:"LivenessType,omitnil,omitempty" name:"LivenessType"`

	// 用于人脸比对的照片的Base64值；
	// Base64编码后的图片数据大小不超过3M，仅支持jpg、png格式。
	// 请使用标准的Base64编码方式(带=补位)，编码规范参考RFC4648。
	// 
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageBase64。
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// 用于人脸比对照片的URL地址；图片下载后经Base64编码后的数据大小不超过3M，仅支持jpg、png格式。
	// 
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageBase64。
	// 
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 数字模式传参：传数字验证码，验证码需先调用<a href="https://cloud.tencent.com/document/product/1007/31821">获取数字验证码接口</a>得到；
	// 动作模式传参：传动作顺序，动作顺序需先调用<a href="https://cloud.tencent.com/document/product/1007/31822">获取动作顺序接口</a>得到；
	// 静默模式传参：空。
	ValidateData *string `json:"ValidateData,omitnil,omitempty" name:"ValidateData"`

	// 额外配置，传入JSON字符串。
	// {
	// "BestFrameNum": 2  //需要返回多张最佳截图，取值范围2-10
	// }
	Optional *string `json:"Optional,omitnil,omitempty" name:"Optional"`

	// 用于活体检测的视频，视频的Base64值；
	// Base64编码后的大小不超过8M，支持mp4、avi、flv格式。
	// 请使用标准的Base64编码方式(带=补位)，编码规范参考RFC4648。
	// 
	// 视频的 VideoUrl、VideoBase64 必须提供一个，如果都提供，只使用 VideoBase64。
	VideoBase64 *string `json:"VideoBase64,omitnil,omitempty" name:"VideoBase64"`

	// 用于活体检测的视频Url 地址。视频下载后经Base64编码后不超过 8M，视频下载耗时不超过4S，支持mp4、avi、flv格式。
	// 
	// 视频的 VideoUrl、VideoBase64 必须提供一个，如果都提供，只使用 VideoBase64。
	// 
	// 建议视频存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议视频存储于腾讯云。非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	VideoUrl *string `json:"VideoUrl,omitnil,omitempty" name:"VideoUrl"`
}

type LivenessCompareRequest struct {
	*tchttp.BaseRequest
	
	// 活体检测类型，取值：LIP/ACTION/SILENT。
	// LIP为数字模式，ACTION为动作模式，SILENT为静默模式，三种模式选择一种传入。
	LivenessType *string `json:"LivenessType,omitnil,omitempty" name:"LivenessType"`

	// 用于人脸比对的照片的Base64值；
	// Base64编码后的图片数据大小不超过3M，仅支持jpg、png格式。
	// 请使用标准的Base64编码方式(带=补位)，编码规范参考RFC4648。
	// 
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageBase64。
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// 用于人脸比对照片的URL地址；图片下载后经Base64编码后的数据大小不超过3M，仅支持jpg、png格式。
	// 
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageBase64。
	// 
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// 数字模式传参：传数字验证码，验证码需先调用<a href="https://cloud.tencent.com/document/product/1007/31821">获取数字验证码接口</a>得到；
	// 动作模式传参：传动作顺序，动作顺序需先调用<a href="https://cloud.tencent.com/document/product/1007/31822">获取动作顺序接口</a>得到；
	// 静默模式传参：空。
	ValidateData *string `json:"ValidateData,omitnil,omitempty" name:"ValidateData"`

	// 额外配置，传入JSON字符串。
	// {
	// "BestFrameNum": 2  //需要返回多张最佳截图，取值范围2-10
	// }
	Optional *string `json:"Optional,omitnil,omitempty" name:"Optional"`

	// 用于活体检测的视频，视频的Base64值；
	// Base64编码后的大小不超过8M，支持mp4、avi、flv格式。
	// 请使用标准的Base64编码方式(带=补位)，编码规范参考RFC4648。
	// 
	// 视频的 VideoUrl、VideoBase64 必须提供一个，如果都提供，只使用 VideoBase64。
	VideoBase64 *string `json:"VideoBase64,omitnil,omitempty" name:"VideoBase64"`

	// 用于活体检测的视频Url 地址。视频下载后经Base64编码后不超过 8M，视频下载耗时不超过4S，支持mp4、avi、flv格式。
	// 
	// 视频的 VideoUrl、VideoBase64 必须提供一个，如果都提供，只使用 VideoBase64。
	// 
	// 建议视频存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议视频存储于腾讯云。非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	VideoUrl *string `json:"VideoUrl,omitnil,omitempty" name:"VideoUrl"`
}

func (r *LivenessCompareRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LivenessCompareRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LivenessType")
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "ValidateData")
	delete(f, "Optional")
	delete(f, "VideoBase64")
	delete(f, "VideoUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "LivenessCompareRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LivenessCompareResponseParams struct {
	// 验证通过后的视频最佳截图照片，照片为BASE64编码后的值，jpg格式。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BestFrameBase64 *string `json:"BestFrameBase64,omitnil,omitempty" name:"BestFrameBase64"`

	// 相似度，取值范围 [0.00, 100.00]。推荐相似度大于等于70时可判断为同一人，可根据具体场景自行调整阈值（阈值70的误通过率为千分之一，阈值80的误通过率是万分之一）。
	Sim *float64 `json:"Sim,omitnil,omitempty" name:"Sim"`

	// 业务错误码，成功情况返回Success, 错误情况请参考下方错误码 列表中FailedOperation部分
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 业务结果描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 最佳截图列表，仅在配置了返回多张最佳截图时返回。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BestFrameList []*string `json:"BestFrameList,omitnil,omitempty" name:"BestFrameList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type LivenessCompareResponse struct {
	*tchttp.BaseResponse
	Response *LivenessCompareResponseParams `json:"Response"`
}

func (r *LivenessCompareResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LivenessCompareResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LivenessRecognitionRequestParams struct {
	// 身份证号
	IdCard *string `json:"IdCard,omitnil,omitempty" name:"IdCard"`

	// 姓名。中文请使用UTF-8编码。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 活体检测类型，取值：LIP/ACTION/SILENT。
	// LIP为数字模式，ACTION为动作模式，SILENT为静默模式，三种模式选择一种传入。
	LivenessType *string `json:"LivenessType,omitnil,omitempty" name:"LivenessType"`

	// 用于活体检测的视频，视频的BASE64值；
	// BASE64编码后的大小不超过8M，支持mp4、avi、flv格式。
	VideoBase64 *string `json:"VideoBase64,omitnil,omitempty" name:"VideoBase64"`

	// 用于活体检测的视频Url 地址。视频下载后经Base64编码不超过 8M，视频下载耗时不超过4S，支持mp4、avi、flv格式。
	// 
	// 视频的 VideoUrl、VideoBase64 必须提供一个，如果都提供，只使用 VideoBase64。
	// 
	// 建议视频存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议视频存储于腾讯云。非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	VideoUrl *string `json:"VideoUrl,omitnil,omitempty" name:"VideoUrl"`

	// 数字模式传参：传数字验证码，验证码需先调用<a href="https://cloud.tencent.com/document/product/1007/31821">获取数字验证码接口</a>得到；
	// 动作模式传参：传动作顺序，动作顺序需先调用<a href="https://cloud.tencent.com/document/product/1007/31822">获取动作顺序接口</a>得到；
	// 静默模式传参：空。
	ValidateData *string `json:"ValidateData,omitnil,omitempty" name:"ValidateData"`

	// 额外配置，传入JSON字符串。
	// {
	// "BestFrameNum": 2  //需要返回多张最佳截图，取值范围2-10
	// }
	Optional *string `json:"Optional,omitnil,omitempty" name:"Optional"`

	// 敏感数据加密信息。对传入信息（姓名、身份证号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

type LivenessRecognitionRequest struct {
	*tchttp.BaseRequest
	
	// 身份证号
	IdCard *string `json:"IdCard,omitnil,omitempty" name:"IdCard"`

	// 姓名。中文请使用UTF-8编码。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 活体检测类型，取值：LIP/ACTION/SILENT。
	// LIP为数字模式，ACTION为动作模式，SILENT为静默模式，三种模式选择一种传入。
	LivenessType *string `json:"LivenessType,omitnil,omitempty" name:"LivenessType"`

	// 用于活体检测的视频，视频的BASE64值；
	// BASE64编码后的大小不超过8M，支持mp4、avi、flv格式。
	VideoBase64 *string `json:"VideoBase64,omitnil,omitempty" name:"VideoBase64"`

	// 用于活体检测的视频Url 地址。视频下载后经Base64编码不超过 8M，视频下载耗时不超过4S，支持mp4、avi、flv格式。
	// 
	// 视频的 VideoUrl、VideoBase64 必须提供一个，如果都提供，只使用 VideoBase64。
	// 
	// 建议视频存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议视频存储于腾讯云。非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	VideoUrl *string `json:"VideoUrl,omitnil,omitempty" name:"VideoUrl"`

	// 数字模式传参：传数字验证码，验证码需先调用<a href="https://cloud.tencent.com/document/product/1007/31821">获取数字验证码接口</a>得到；
	// 动作模式传参：传动作顺序，动作顺序需先调用<a href="https://cloud.tencent.com/document/product/1007/31822">获取动作顺序接口</a>得到；
	// 静默模式传参：空。
	ValidateData *string `json:"ValidateData,omitnil,omitempty" name:"ValidateData"`

	// 额外配置，传入JSON字符串。
	// {
	// "BestFrameNum": 2  //需要返回多张最佳截图，取值范围2-10
	// }
	Optional *string `json:"Optional,omitnil,omitempty" name:"Optional"`

	// 敏感数据加密信息。对传入信息（姓名、身份证号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

func (r *LivenessRecognitionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LivenessRecognitionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdCard")
	delete(f, "Name")
	delete(f, "LivenessType")
	delete(f, "VideoBase64")
	delete(f, "VideoUrl")
	delete(f, "ValidateData")
	delete(f, "Optional")
	delete(f, "Encryption")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "LivenessRecognitionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LivenessRecognitionResponseParams struct {
	// 验证通过后的视频最佳截图照片，照片为BASE64编码后的值，jpg格式。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BestFrameBase64 *string `json:"BestFrameBase64,omitnil,omitempty" name:"BestFrameBase64"`

	// 相似度，取值范围 [0.00, 100.00]。推荐相似度大于等于70时可判断为同一人，可根据具体场景自行调整阈值（阈值70的误通过率为千分之一，阈值80的误通过率是万分之一）
	Sim *float64 `json:"Sim,omitnil,omitempty" name:"Sim"`

	// 业务错误码，成功情况返回Success, 错误情况请参考下方错误码 列表中FailedOperation部分
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 业务结果描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 最佳截图列表，仅在配置了返回多张最佳截图时返回。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BestFrameList []*string `json:"BestFrameList,omitnil,omitempty" name:"BestFrameList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type LivenessRecognitionResponse struct {
	*tchttp.BaseResponse
	Response *LivenessRecognitionResponseParams `json:"Response"`
}

func (r *LivenessRecognitionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LivenessRecognitionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LivenessRequestParams struct {
	// 用于活体检测的视频，视频的BASE64值；
	// BASE64编码后的大小不超过8M，支持mp4、avi、flv格式。
	VideoBase64 *string `json:"VideoBase64,omitnil,omitempty" name:"VideoBase64"`

	// 活体检测类型，取值：LIP/ACTION/SILENT。
	// LIP为数字模式，ACTION为动作模式，SILENT为静默模式，三种模式选择一种传入。
	LivenessType *string `json:"LivenessType,omitnil,omitempty" name:"LivenessType"`

	// 数字模式传参：数字验证码(1234)，需先调用接口获取数字验证码；
	// 动作模式传参：传动作顺序(2,1 or 1,2)，需先调用接口获取动作顺序；
	// 静默模式传参：不需要传递此参数。
	ValidateData *string `json:"ValidateData,omitnil,omitempty" name:"ValidateData"`

	// 额外配置，传入JSON字符串。
	// {
	// "BestFrameNum": 2  //需要返回多张最佳截图，取值范围1-10
	// }
	Optional *string `json:"Optional,omitnil,omitempty" name:"Optional"`
}

type LivenessRequest struct {
	*tchttp.BaseRequest
	
	// 用于活体检测的视频，视频的BASE64值；
	// BASE64编码后的大小不超过8M，支持mp4、avi、flv格式。
	VideoBase64 *string `json:"VideoBase64,omitnil,omitempty" name:"VideoBase64"`

	// 活体检测类型，取值：LIP/ACTION/SILENT。
	// LIP为数字模式，ACTION为动作模式，SILENT为静默模式，三种模式选择一种传入。
	LivenessType *string `json:"LivenessType,omitnil,omitempty" name:"LivenessType"`

	// 数字模式传参：数字验证码(1234)，需先调用接口获取数字验证码；
	// 动作模式传参：传动作顺序(2,1 or 1,2)，需先调用接口获取动作顺序；
	// 静默模式传参：不需要传递此参数。
	ValidateData *string `json:"ValidateData,omitnil,omitempty" name:"ValidateData"`

	// 额外配置，传入JSON字符串。
	// {
	// "BestFrameNum": 2  //需要返回多张最佳截图，取值范围1-10
	// }
	Optional *string `json:"Optional,omitnil,omitempty" name:"Optional"`
}

func (r *LivenessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LivenessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VideoBase64")
	delete(f, "LivenessType")
	delete(f, "ValidateData")
	delete(f, "Optional")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "LivenessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LivenessResponseParams struct {
	// 验证通过后的视频最佳截图照片，照片为BASE64编码后的值，jpg格式。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BestFrameBase64 *string `json:"BestFrameBase64,omitnil,omitempty" name:"BestFrameBase64"`

	// 业务错误码，成功情况返回Success, 错误情况请参考下方错误码 列表中FailedOperation部分
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 业务结果描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 最佳最佳截图列表，仅在配置了返回多张最佳截图时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BestFrameList []*string `json:"BestFrameList,omitnil,omitempty" name:"BestFrameList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type LivenessResponse struct {
	*tchttp.BaseResponse
	Response *LivenessResponseParams `json:"Response"`
}

func (r *LivenessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LivenessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MinorsVerificationRequestParams struct {
	// 参与校验的参数类型。
	// 0：使用手机号进行校验；
	// 1：使用姓名与身份证号进行校验。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 手机号，11位数字，
	// 特别提示：
	// 手机号验证只限制在腾讯健康守护可信模型覆盖的数据范围内，与手机号本身在运营商是否实名无关联，不在范围会提示“手机号未实名”，建议客户与传入姓名和身份证号信息组合使用。
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 身份证号码。
	IdCard *string `json:"IdCard,omitnil,omitempty" name:"IdCard"`

	// 姓名。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 敏感数据加密信息。对传入信息（姓名、身份证号、手机号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

type MinorsVerificationRequest struct {
	*tchttp.BaseRequest
	
	// 参与校验的参数类型。
	// 0：使用手机号进行校验；
	// 1：使用姓名与身份证号进行校验。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 手机号，11位数字，
	// 特别提示：
	// 手机号验证只限制在腾讯健康守护可信模型覆盖的数据范围内，与手机号本身在运营商是否实名无关联，不在范围会提示“手机号未实名”，建议客户与传入姓名和身份证号信息组合使用。
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 身份证号码。
	IdCard *string `json:"IdCard,omitnil,omitempty" name:"IdCard"`

	// 姓名。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 敏感数据加密信息。对传入信息（姓名、身份证号、手机号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

func (r *MinorsVerificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MinorsVerificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "Mobile")
	delete(f, "IdCard")
	delete(f, "Name")
	delete(f, "Encryption")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MinorsVerificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MinorsVerificationResponseParams struct {
	// 结果码，收费情况如下。
	// 收费结果码：
	// 0: 成年
	// -1: 未成年
	// -3: 姓名和身份证号不一致
	// 
	// 不收费结果码：
	// -2: 未查询到手机号信息
	// -4: 非法身份证号（长度、校验位等不正确）
	// -5: 非法姓名（长度、格式等不正确）
	// -6: 权威数据源服务异常
	// -7: 未查询到身份信息
	// -8: 权威数据源升级中，请稍后再试
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 业务结果描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 该字段的值为年龄区间。格式为[a,b)，
	// [0,8)表示年龄小于8周岁区间，不包括8岁；
	// [8,16)表示年龄8-16周岁区间，不包括16岁；
	// [16,18)表示年龄16-18周岁区间，不包括18岁；
	// [18,+)表示年龄大于18周岁。
	AgeRange *string `json:"AgeRange,omitnil,omitempty" name:"AgeRange"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type MinorsVerificationResponse struct {
	*tchttp.BaseResponse
	Response *MinorsVerificationResponseParams `json:"Response"`
}

func (r *MinorsVerificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MinorsVerificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MobileNetworkTimeVerificationRequestParams struct {
	// 手机号码
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 敏感数据加密信息。对传入信息（手机号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

type MobileNetworkTimeVerificationRequest struct {
	*tchttp.BaseRequest
	
	// 手机号码
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 敏感数据加密信息。对传入信息（手机号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

func (r *MobileNetworkTimeVerificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MobileNetworkTimeVerificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Mobile")
	delete(f, "Encryption")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MobileNetworkTimeVerificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MobileNetworkTimeVerificationResponseParams struct {
	// 认证结果码，收费情况如下。
	// 收费结果码：
	// 0: 成功
	// -2: 手机号不存在
	// -3: 手机号存在，但无法查询到在网时长
	// 不收费结果码：
	// -1: 手机号格式不正确
	// -4: 验证中心服务繁忙
	// -5：认证次数超过当日限制，请次日重试
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 业务结果描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 在网时长区间。
	// 格式为(a,b]，表示在网时长在a个月以上，b个月以下。若b为+时表示没有上限。
	Range *string `json:"Range,omitnil,omitempty" name:"Range"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type MobileNetworkTimeVerificationResponse struct {
	*tchttp.BaseResponse
	Response *MobileNetworkTimeVerificationResponseParams `json:"Response"`
}

func (r *MobileNetworkTimeVerificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MobileNetworkTimeVerificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MobileStatusRequestParams struct {
	// 手机号码
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 敏感数据加密信息。对传入信息（手机号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

type MobileStatusRequest struct {
	*tchttp.BaseRequest
	
	// 手机号码
	Mobile *string `json:"Mobile,omitnil,omitempty" name:"Mobile"`

	// 敏感数据加密信息。对传入信息（手机号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

func (r *MobileStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MobileStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Mobile")
	delete(f, "Encryption")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MobileStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MobileStatusResponseParams struct {
	// 认证结果码，收费情况如下。
	// 收费结果码：
	// 0：成功
	// 不收费结果码：
	// -1：未查询到结果
	// -2：手机号格式不正确
	// -3：验证中心服务繁忙
	// -4：认证次数超过当日限制，请次日重试
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 业务结果描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 状态码：
	// 0：正常
	// 1：停机
	// 2：销号
	// 3：空号
	// 4：不在网
	// 99：未知状态
	StatusCode *int64 `json:"StatusCode,omitnil,omitempty" name:"StatusCode"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type MobileStatusResponse struct {
	*tchttp.BaseResponse
	Response *MobileStatusResponseParams `json:"Response"`
}

func (r *MobileStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MobileStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ParseNfcDataRequestParams struct {
	// 前端SDK返回
	ReqId *string `json:"ReqId,omitnil,omitempty" name:"ReqId"`
}

type ParseNfcDataRequest struct {
	*tchttp.BaseRequest
	
	// 前端SDK返回
	ReqId *string `json:"ReqId,omitnil,omitempty" name:"ReqId"`
}

func (r *ParseNfcDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ParseNfcDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReqId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ParseNfcDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ParseNfcDataResponseParams struct {
	// 0为首次查询成功，-1为查询失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultCode *string `json:"ResultCode,omitnil,omitempty" name:"ResultCode"`

	// 身份证号
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdNum *string `json:"IdNum,omitnil,omitempty" name:"IdNum"`

	// 姓名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 照片
	// 注意：此字段可能返回 null，表示取不到有效值。
	Picture *string `json:"Picture,omitnil,omitempty" name:"Picture"`

	// 出生日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	BirthDate *string `json:"BirthDate,omitnil,omitempty" name:"BirthDate"`

	// 有效期起始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 有效期结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 住址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// 民族
	// 注意：此字段可能返回 null，表示取不到有效值。
	Nation *string `json:"Nation,omitnil,omitempty" name:"Nation"`

	// 性别
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sex *string `json:"Sex,omitnil,omitempty" name:"Sex"`

	// 身份证 01 中国护照 03 军官证 04 武警证 05 港澳通行证 06 台胞证 07 外国护照 08 士兵证 09 临时身份证 10 户口本 11 警官证 12 外国人永久居留证 13 港澳台居民居住证 14 回乡证 15 大陆居民来往台湾通行证 16 其他证件 99
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdType *string `json:"IdType,omitnil,omitempty" name:"IdType"`

	// 英文姓名
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnName *string `json:"EnName,omitnil,omitempty" name:"EnName"`

	// 签发机关
	// 注意：此字段可能返回 null，表示取不到有效值。
	SigningOrganization *string `json:"SigningOrganization,omitnil,omitempty" name:"SigningOrganization"`

	// 港澳台居民居住证，通行证号码
	// 注意：此字段可能返回 null，表示取不到有效值。
	OtherIdNum *string `json:"OtherIdNum,omitnil,omitempty" name:"OtherIdNum"`

	// 旅行证件国籍
	// 注意：此字段可能返回 null，表示取不到有效值。
	Nationality *string `json:"Nationality,omitnil,omitempty" name:"Nationality"`

	// 旅行证件机读区第二行 29~42 位
	// 注意：此字段可能返回 null，表示取不到有效值。
	PersonalNumber *string `json:"PersonalNumber,omitnil,omitempty" name:"PersonalNumber"`

	// 旅行证件类的核验结果。JSON格式如下：
	// {"result_issuer ":"签发者证书合法性验证结果 ","result_pape r":"证件安全对象合法性验证 结果 ","result_data" :"防数据篡改验证结果 ","result_chip" :"防证书件芯片被复制验证结果"} 
	//  0:验证通过，1: 验证不通过，2: 未验证，3:部分通过，当4项核验结果都为0时，表示证件为真
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckMRTD *string `json:"CheckMRTD,omitnil,omitempty" name:"CheckMRTD"`

	// 身份证照片面合成图片
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageA *string `json:"ImageA,omitnil,omitempty" name:"ImageA"`

	// 身份证国徽面合成图片
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageB *string `json:"ImageB,omitnil,omitempty" name:"ImageB"`

	// 对result code的结果描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultDescription *string `json:"ResultDescription,omitnil,omitempty" name:"ResultDescription"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ParseNfcDataResponse struct {
	*tchttp.BaseResponse
	Response *ParseNfcDataResponseParams `json:"Response"`
}

func (r *ParseNfcDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ParseNfcDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PhoneVerificationCMCCRequestParams struct {
	// 身份证号
	IdCard *string `json:"IdCard,omitnil,omitempty" name:"IdCard"`

	// 姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 手机号
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 敏感数据加密信息。对传入信息（姓名、身份证号、手机号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

type PhoneVerificationCMCCRequest struct {
	*tchttp.BaseRequest
	
	// 身份证号
	IdCard *string `json:"IdCard,omitnil,omitempty" name:"IdCard"`

	// 姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 手机号
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 敏感数据加密信息。对传入信息（姓名、身份证号、手机号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

func (r *PhoneVerificationCMCCRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PhoneVerificationCMCCRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdCard")
	delete(f, "Name")
	delete(f, "Phone")
	delete(f, "Encryption")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PhoneVerificationCMCCRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PhoneVerificationCMCCResponseParams struct {
	// 认证结果码，收费情况如下。
	// 收费结果码：
	// 0: 认证通过
	// -4: 信息不一致（手机号已实名，但姓名和身份证号与实名信息不一致）
	// 不收费结果码：
	// -6: 手机号码不合法
	// -7: 身份证号码有误
	// -8: 姓名校验不通过
	// -9: 没有记录
	// -11: 验证中心服务繁忙
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 运营商名称。
	// 取值范围为["移动","联通","电信",""]
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 业务结果描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PhoneVerificationCMCCResponse struct {
	*tchttp.BaseResponse
	Response *PhoneVerificationCMCCResponseParams `json:"Response"`
}

func (r *PhoneVerificationCMCCResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PhoneVerificationCMCCResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PhoneVerificationCTCCRequestParams struct {
	// 身份证号
	IdCard *string `json:"IdCard,omitnil,omitempty" name:"IdCard"`

	// 姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 手机号
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 敏感数据加密信息。对传入信息（姓名、身份证号、手机号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

type PhoneVerificationCTCCRequest struct {
	*tchttp.BaseRequest
	
	// 身份证号
	IdCard *string `json:"IdCard,omitnil,omitempty" name:"IdCard"`

	// 姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 手机号
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 敏感数据加密信息。对传入信息（姓名、身份证号、手机号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

func (r *PhoneVerificationCTCCRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PhoneVerificationCTCCRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdCard")
	delete(f, "Name")
	delete(f, "Phone")
	delete(f, "Encryption")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PhoneVerificationCTCCRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PhoneVerificationCTCCResponseParams struct {
	// 认证结果码，收费情况如下。
	// 收费结果码：
	// 0: 认证通过
	// -4: 信息不一致（手机号已实名，但姓名和身份证号与实名信息不一致）
	// 不收费结果码：
	// -6: 手机号码不合法
	// -7: 身份证号码有误
	// -8: 姓名校验不通过
	// -9: 没有记录
	// -11: 验证中心服务繁忙
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 运营商名称。
	// 取值范围为["移动","联通","电信",""]
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 业务结果描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PhoneVerificationCTCCResponse struct {
	*tchttp.BaseResponse
	Response *PhoneVerificationCTCCResponseParams `json:"Response"`
}

func (r *PhoneVerificationCTCCResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PhoneVerificationCTCCResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PhoneVerificationCUCCRequestParams struct {
	// 身份证号
	IdCard *string `json:"IdCard,omitnil,omitempty" name:"IdCard"`

	// 姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 手机号
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 敏感数据加密信息。对传入信息（姓名、身份证号、手机号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

type PhoneVerificationCUCCRequest struct {
	*tchttp.BaseRequest
	
	// 身份证号
	IdCard *string `json:"IdCard,omitnil,omitempty" name:"IdCard"`

	// 姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 手机号
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 敏感数据加密信息。对传入信息（姓名、身份证号、手机号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitnil,omitempty" name:"Encryption"`
}

func (r *PhoneVerificationCUCCRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PhoneVerificationCUCCRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdCard")
	delete(f, "Name")
	delete(f, "Phone")
	delete(f, "Encryption")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PhoneVerificationCUCCRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PhoneVerificationCUCCResponseParams struct {
	// 认证结果码，收费情况如下。
	// 收费结果码：
	// 0: 认证通过
	// -4: 信息不一致（手机号已实名，但姓名和身份证号与实名信息不一致）
	// 不收费结果码：
	// -6: 手机号码不合法
	// -7: 身份证号码有误
	// -8: 姓名校验不通过
	// -9: 没有记录
	// -11: 验证中心服务繁忙
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 运营商名称。
	// 取值范围为["移动","联通","电信",""]
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 业务结果描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PhoneVerificationCUCCResponse struct {
	*tchttp.BaseResponse
	Response *PhoneVerificationCUCCResponseParams `json:"Response"`
}

func (r *PhoneVerificationCUCCResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PhoneVerificationCUCCResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PhoneVerificationRequestParams struct {
	// 身份证号
	IdCard *string `json:"IdCard,omitnil,omitempty" name:"IdCard"`

	// 姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 手机号
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 验证模式（详版/简版）。简版与详版价格不一致，详见[价格说明](https://cloud.tencent.com/document/product/1007/84321)。
	// 
	// 枚举值：0（简版），1（详版）。默认值为0。
	VerifyMode *string `json:"VerifyMode,omitnil,omitempty" name:"VerifyMode"`

	// 有加密需求的用户，传入kms的CiphertextBlob，关于数据加密可查阅 <a href="https://cloud.tencent.com/document/product/1007/47180">数据加密</a> 文档。
	CiphertextBlob *string `json:"CiphertextBlob,omitnil,omitempty" name:"CiphertextBlob"`

	// 在使用加密服务时，填入要被加密的字段。本接口中可填入加密后的IdCard，Name，Phone中的一个或多个。
	EncryptList []*string `json:"EncryptList,omitnil,omitempty" name:"EncryptList"`

	// 有加密需求的用户，传入CBC加密的初始向量。
	Iv *string `json:"Iv,omitnil,omitempty" name:"Iv"`
}

type PhoneVerificationRequest struct {
	*tchttp.BaseRequest
	
	// 身份证号
	IdCard *string `json:"IdCard,omitnil,omitempty" name:"IdCard"`

	// 姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 手机号
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 验证模式（详版/简版）。简版与详版价格不一致，详见[价格说明](https://cloud.tencent.com/document/product/1007/84321)。
	// 
	// 枚举值：0（简版），1（详版）。默认值为0。
	VerifyMode *string `json:"VerifyMode,omitnil,omitempty" name:"VerifyMode"`

	// 有加密需求的用户，传入kms的CiphertextBlob，关于数据加密可查阅 <a href="https://cloud.tencent.com/document/product/1007/47180">数据加密</a> 文档。
	CiphertextBlob *string `json:"CiphertextBlob,omitnil,omitempty" name:"CiphertextBlob"`

	// 在使用加密服务时，填入要被加密的字段。本接口中可填入加密后的IdCard，Name，Phone中的一个或多个。
	EncryptList []*string `json:"EncryptList,omitnil,omitempty" name:"EncryptList"`

	// 有加密需求的用户，传入CBC加密的初始向量。
	Iv *string `json:"Iv,omitnil,omitempty" name:"Iv"`
}

func (r *PhoneVerificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PhoneVerificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdCard")
	delete(f, "Name")
	delete(f, "Phone")
	delete(f, "VerifyMode")
	delete(f, "CiphertextBlob")
	delete(f, "EncryptList")
	delete(f, "Iv")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PhoneVerificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PhoneVerificationResponseParams struct {
	// 认证结果码:
	// 收费结果码
	// 0: 三要素信息一致
	// -4: 三要素信息不一致
	// 不收费结果码
	// -6: 手机号码不合法
	// -7: 身份证号码有误
	// -8: 姓名校验不通过
	// -9: 没有记录
	// -11: 验证中心服务繁忙
	// -12: 认证次数超过当日限制，请次日重试
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 业务结果描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 运营商名称。
	// 取值范围为["","移动","电信","联通"]
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 业务结果详细信息。（当VerifyMode配置"详版"，且Result为"-4: 三要素信息不一致"时返回）
	// 枚举值：
	// PhoneIdCardMismatch：手机号码与姓名一致，与身份证号不一致；
	// PhoneNameMismatch：手机号码身份证号一致，与姓名不一致；
	// PhoneNameIdCardMismatch：手机号码与姓名和身份证号均不一致；
	// NameIdCardMismatch：姓名和身份证号不一致；
	// OtherMismatch：其他不一致；
	ResultDetail *string `json:"ResultDetail,omitnil,omitempty" name:"ResultDetail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PhoneVerificationResponse struct {
	*tchttp.BaseResponse
	Response *PhoneVerificationResponseParams `json:"Response"`
}

func (r *PhoneVerificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PhoneVerificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetrievalLivenessExtraInfo struct {
	// 命中的模版类型，其中Common-公共库；Auto-自动聚类库；Owner-自建模版库
	// 注意：此字段可能返回 null，表示取不到有效值。
	HitGroup *string `json:"HitGroup,omitnil,omitempty" name:"HitGroup"`

	// 命中的相似度
	// 注意：此字段可能返回 null，表示取不到有效值。
	SimilarityScore *float64 `json:"SimilarityScore,omitnil,omitempty" name:"SimilarityScore"`

	// 命中的模板id
	// 注意：此字段可能返回 null，表示取不到有效值。
	HitTemplate *string `json:"HitTemplate,omitnil,omitempty" name:"HitTemplate"`
}

type RuleIdConfig struct {
	// 意愿核身过程中识别用户的回答意图，开启后除了IntentionQuestions的Answers列表中的标准回答会通过，近似意图的回答也会通过，默认开启。
	IntentionRecognition *bool `json:"IntentionRecognition,omitnil,omitempty" name:"IntentionRecognition"`

	// 意愿核身类型，默认为0：
	// 0：问答模式，DetectAuth接口需要传入IntentionQuestions字段；
	// 1：点头模式，DetectAuth接口需要传入IntentionActions字段；
	IntentionType *int64 `json:"IntentionType,omitnil,omitempty" name:"IntentionType"`

	// 用户语音回答过程中是否开启张嘴识别检测，默认不开启，仅在意愿核身问答模式中使用。
	MouthOpenRecognition *bool `json:"MouthOpenRecognition,omitnil,omitempty" name:"MouthOpenRecognition"`

	// 意愿核身语音播报速度，配置后问答模式和点头模式的语音播报环节都会生效，默认值为0：
	// 0：智能语速（根据播报文案的长度自动调整语音播报速度）
	// 1：固定1倍速
	// 2：固定1.2倍速
	// 3：固定1.5倍速
	Speed *uint64 `json:"Speed,omitnil,omitempty" name:"Speed"`
}

type WeChatBillDetail struct {
	// token
	BizToken *string `json:"BizToken,omitnil,omitempty" name:"BizToken"`

	// 本token收费次数
	ChargeCount *uint64 `json:"ChargeCount,omitnil,omitempty" name:"ChargeCount"`

	// 本token计费详情
	ChargeDetails []*ChargeDetail `json:"ChargeDetails,omitnil,omitempty" name:"ChargeDetails"`

	// 业务RuleId
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}