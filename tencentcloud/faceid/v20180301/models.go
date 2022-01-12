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
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type BankCard2EVerificationRequest struct {
	*tchttp.BaseRequest

	// 姓名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 银行卡
	BankCard *string `json:"BankCard,omitempty" name:"BankCard"`

	// 敏感数据加密信息。对传入信息（姓名、银行卡号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitempty" name:"Encryption"`
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

type BankCard2EVerificationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
		Result *string `json:"Result,omitempty" name:"Result"`

		// 业务结果描述。
		Description *string `json:"Description,omitempty" name:"Description"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type BankCard4EVerificationRequest struct {
	*tchttp.BaseRequest

	// 姓名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 银行卡
	BankCard *string `json:"BankCard,omitempty" name:"BankCard"`

	// 手机号码
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 开户证件号，与CertType参数的证件类型一致，如：身份证，则传入身份证号。
	IdCard *string `json:"IdCard,omitempty" name:"IdCard"`

	// 证件类型，请确认该证件为开户时使用的证件类型，未用于开户的证件信息不支持验证。
	// 目前默认为0：身份证，其他证件类型暂不支持。
	CertType *int64 `json:"CertType,omitempty" name:"CertType"`

	// 敏感数据加密信息。对传入信息（姓名、身份证号、手机号、银行卡号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitempty" name:"Encryption"`
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

type BankCard4EVerificationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
		Result *string `json:"Result,omitempty" name:"Result"`

		// 业务结果描述。
		Description *string `json:"Description,omitempty" name:"Description"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type BankCardVerificationRequest struct {
	*tchttp.BaseRequest

	// 开户证件号，与CertType参数的证件类型一致，如：身份证，则传入身份证号。
	IdCard *string `json:"IdCard,omitempty" name:"IdCard"`

	// 姓名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 银行卡
	BankCard *string `json:"BankCard,omitempty" name:"BankCard"`

	// 证件类型，请确认该证件为开户时使用的证件类型，未用于开户的证件信息不支持验证。
	// 目前默认：0 身份证，其他证件类型需求可以添加[腾讯云人脸核身小助手](https://cloud.tencent.com/document/product/1007/56130)进行确认。
	CertType *int64 `json:"CertType,omitempty" name:"CertType"`

	// 敏感数据加密信息。对传入信息（姓名、身份证号、银行卡号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitempty" name:"Encryption"`
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

type BankCardVerificationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
		Result *string `json:"Result,omitempty" name:"Result"`

		// 业务结果描述。
		Description *string `json:"Description,omitempty" name:"Description"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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
	ReqTime *string `json:"ReqTime,omitempty" name:"ReqTime"`

	// 一比一请求的唯一标记。
	Seq *string `json:"Seq,omitempty" name:"Seq"`

	// 一比一时使用的、脱敏后的身份证号。
	Idcard *string `json:"Idcard,omitempty" name:"Idcard"`

	// 一比一时使用的、脱敏后的姓名。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 一比一的相似度。0-100，保留2位小数。
	Sim *string `json:"Sim,omitempty" name:"Sim"`

	// 本次详情是否收费。
	IsNeedCharge *bool `json:"IsNeedCharge,omitempty" name:"IsNeedCharge"`

	// 收费类型，比对、核身、混合部署。
	ChargeType *string `json:"ChargeType,omitempty" name:"ChargeType"`

	// 本次活体一比一最终结果。
	ErrorCode *string `json:"ErrorCode,omitempty" name:"ErrorCode"`

	// 本次活体一比一最终结果描述。
	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`
}

type CheckBankCardInformationRequest struct {
	*tchttp.BaseRequest

	// 银行卡号。
	BankCard *string `json:"BankCard,omitempty" name:"BankCard"`

	// 敏感数据加密信息。对传入信息（银行卡号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitempty" name:"Encryption"`
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

type CheckBankCardInformationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 认证结果码，收费情况如下。
	// 收费结果码：
	// 0: 查询成功
	// -1: 未查到信息
	// 不收费结果码
	// -2：验证中心服务繁忙
	// -3：银行卡不存在
		Result *string `json:"Result,omitempty" name:"Result"`

		// 业务结果描述
		Description *string `json:"Description,omitempty" name:"Description"`

		// 开户行
		AccountBank *string `json:"AccountBank,omitempty" name:"AccountBank"`

		// 卡性质：1. 借记卡；2. 贷记卡
		AccountType *int64 `json:"AccountType,omitempty" name:"AccountType"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CheckEidTokenStatusRequest struct {
	*tchttp.BaseRequest

	// E证通流程的唯一标识，调用GetEidToken接口时生成。
	EidToken *string `json:"EidToken,omitempty" name:"EidToken"`
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

type CheckEidTokenStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 枚举：
	// init：token未验证
	// doing: 验证中
	// finished: 验证完成
	// timeout: token已超时
		Status *string `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CheckIdCardInformationRequest struct {
	*tchttp.BaseRequest

	// 身份证人像面的 Base64 值
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 7M。
	// 请使用标准的Base64编码方式(带=补位)，编码规范参考RFC4648。
	// ImageBase64、ImageUrl二者必须提供其中之一。若都提供了，则按照ImageUrl>ImageBase64的优先级使用参数。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 身份证人像面的 Url 地址
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 3M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 以下可选字段均为bool 类型，默认false：
	// CopyWarn，复印件告警
	// BorderCheckWarn，边框和框内遮挡告警
	// ReshootWarn，翻拍告警
	// DetectPsWarn，PS检测告警
	// TempIdWarn，临时身份证告警
	// Quality，图片质量告警（评价图片模糊程度）
	// 
	// SDK 设置方式参考：
	// Config = Json.stringify({"CopyWarn":true,"ReshootWarn":true})
	// API 3.0 Explorer 设置方式参考：
	// Config = {"CopyWarn":true,"ReshootWarn":true}
	Config *string `json:"Config,omitempty" name:"Config"`

	// 是否需要对返回中的敏感信息进行加密。默认false。
	// 其中敏感信息包括：Response.IdNum、Response.Name
	IsEncrypt *bool `json:"IsEncrypt,omitempty" name:"IsEncrypt"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckIdCardInformationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CheckIdCardInformationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 相似度，取值范围 [0.00, 100.00]。推荐相似度大于等于70时可判断为同一人，可根据具体场景自行调整阈值（阈值70的误通过率为千分之一，阈值80的误通过率是万分之一）
		Sim *float64 `json:"Sim,omitempty" name:"Sim"`

		// 业务错误码，成功情况返回Success, 错误情况请参考下方错误码 列表中FailedOperation部分
		Result *string `json:"Result,omitempty" name:"Result"`

		// 业务结果描述。
		Description *string `json:"Description,omitempty" name:"Description"`

		// 姓名
		Name *string `json:"Name,omitempty" name:"Name"`

		// 性别
		Sex *string `json:"Sex,omitempty" name:"Sex"`

		// 民族
		Nation *string `json:"Nation,omitempty" name:"Nation"`

		// 出生日期
		Birth *string `json:"Birth,omitempty" name:"Birth"`

		// 地址
		Address *string `json:"Address,omitempty" name:"Address"`

		// 身份证号
		IdNum *string `json:"IdNum,omitempty" name:"IdNum"`

		// 身份证头像照片的base64编码，如果抠图失败会拿整张身份证做比对并返回空。
		Portrait *string `json:"Portrait,omitempty" name:"Portrait"`

		// 告警信息，当在Config中配置了告警信息会停止人像比对，Result返回错误（FailedOperation.OcrWarningOccurred）并有此告警信息，Code 告警码列表和释义：
	// 
	// -9101 身份证边框不完整告警，
	// -9102 身份证复印件告警，
	// -9103 身份证翻拍告警，
	// -9105 身份证框内遮挡告警，
	// -9104 临时身份证告警，
	// -9106 身份证 PS 告警。
	// -8001 图片模糊告警
	// 多个会 |  隔开如 "-9101|-9106|-9104"
		Warnings *string `json:"Warnings,omitempty" name:"Warnings"`

		// 图片质量分数，当请求Config中配置图片模糊告警该参数才有意义，取值范围（0～100），目前默认阈值是50分，低于50分会触发模糊告警。
		Quality *float64 `json:"Quality,omitempty" name:"Quality"`

		// 敏感数据加密信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Encryption *Encryption `json:"Encryption,omitempty" name:"Encryption"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CheckIdNameDateRequest struct {
	*tchttp.BaseRequest

	// 姓名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 身份证号
	IdCard *string `json:"IdCard,omitempty" name:"IdCard"`

	// 身份证有效期开始时间，格式：YYYYMMDD。如：20210701
	ValidityBegin *string `json:"ValidityBegin,omitempty" name:"ValidityBegin"`

	// 身份证有效期到期时间，格式：YYYYMMDD，长期用“00000000”代替；如：20210701
	ValidityEnd *string `json:"ValidityEnd,omitempty" name:"ValidityEnd"`

	// 敏感数据加密信息。对传入信息（姓名、身份证号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitempty" name:"Encryption"`
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

type CheckIdNameDateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
		Result *string `json:"Result,omitempty" name:"Result"`

		// 业务结果描述。
		Description *string `json:"Description,omitempty" name:"Description"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CheckPhoneAndNameRequest struct {
	*tchttp.BaseRequest

	// ⼿机号
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 姓名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 敏感数据加密信息。对传入信息（姓名、手机号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitempty" name:"Encryption"`
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

type CheckPhoneAndNameResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 认证结果码，收费情况如下。
	// 收费结果码：
	// 0: 验证结果一致
	// 1: 验证结果不一致
	// 不收费结果码：
	// -1:查无记录
	// -2:引擎未知错误
	// -3:引擎服务异常
		Result *string `json:"Result,omitempty" name:"Result"`

		// 业务结果描述
		Description *string `json:"Description,omitempty" name:"Description"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DetectAuthRequest struct {
	*tchttp.BaseRequest

	// 用于细分客户使用场景，申请开通服务后，可以在腾讯云慧眼人脸核身控制台（https://console.cloud.tencent.com/faceid） 自助接入里面创建，审核通过后即可调用。如有疑问，请添加[腾讯云人脸核身小助手](https://cloud.tencent.com/document/product/1007/56130)进行咨询。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 本接口不需要传递此参数。
	TerminalType *string `json:"TerminalType,omitempty" name:"TerminalType"`

	// 身份标识（未使用OCR服务时，必须传入）。
	// 规则：a-zA-Z0-9组合。最长长度32位。
	IdCard *string `json:"IdCard,omitempty" name:"IdCard"`

	// 姓名。（未使用OCR服务时，必须传入）最长长度32位。中文请使用UTF-8编码。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 认证结束后重定向的回调链接地址。最长长度1024位。
	RedirectUrl *string `json:"RedirectUrl,omitempty" name:"RedirectUrl"`

	// 透传字段，在获取验证结果时返回。
	Extra *string `json:"Extra,omitempty" name:"Extra"`

	// 用于人脸比对的照片，图片的Base64值；
	// Base64编码后的图片数据大小不超过3M，仅支持jpg、png格式。请使用标准的Base64编码方式(带=补位)，编码规范参考RFC4648。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 敏感数据加密信息。对传入信息（姓名、身份证号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitempty" name:"Encryption"`

	// 意愿核身使用的文案，若未使用意愿核身功能，该字段无需传入。默认为空，最长可接受100的字符串长度。
	IntentionVerifyText *string `json:"IntentionVerifyText,omitempty" name:"IntentionVerifyText"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetectAuthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DetectAuthResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用于发起核身流程的URL，仅微信H5场景使用。
		Url *string `json:"Url,omitempty" name:"Url"`

		// 一次核身流程的标识，有效时间为7,200秒；
	// 完成核身后，可用该标识获取验证结果信息。
		BizToken *string `json:"BizToken,omitempty" name:"BizToken"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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
	ReqTime *string `json:"ReqTime,omitempty" name:"ReqTime"`

	// 本次活体一比一请求的唯一标记。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Seq *string `json:"Seq,omitempty" name:"Seq"`

	// 参与本次活体一比一的身份证号。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Idcard *string `json:"Idcard,omitempty" name:"Idcard"`

	// 参与本次活体一比一的姓名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 本次活体一比一的相似度。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sim *string `json:"Sim,omitempty" name:"Sim"`

	// 本次活体一比一是否收费
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsNeedCharge *bool `json:"IsNeedCharge,omitempty" name:"IsNeedCharge"`

	// 本次活体一比一最终结果。0为成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Errcode *int64 `json:"Errcode,omitempty" name:"Errcode"`

	// 本次活体一比一最终结果描述。（仅描述用，文案更新时不会通知。）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Errmsg *string `json:"Errmsg,omitempty" name:"Errmsg"`

	// 本次活体结果。0为成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Livestatus *int64 `json:"Livestatus,omitempty" name:"Livestatus"`

	// 本次活体结果描述。（仅描述用，文案更新时不会通知。）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Livemsg *string `json:"Livemsg,omitempty" name:"Livemsg"`

	// 本次一比一结果。0为成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comparestatus *int64 `json:"Comparestatus,omitempty" name:"Comparestatus"`

	// 本次一比一结果描述。（仅描述用，文案更新时不会通知。）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comparemsg *string `json:"Comparemsg,omitempty" name:"Comparemsg"`

	// 比对库源类型。包括：
	// 公安商业库；
	// 业务方自有库（用户上传照片、客户的混合库、混合部署库）；
	// 二次验证库；
	// 人工审核库；
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompareLibType *string `json:"CompareLibType,omitempty" name:"CompareLibType"`
}

type DetectInfoBestFrame struct {

	// 活体比对最佳帧Base64编码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BestFrame *string `json:"BestFrame,omitempty" name:"BestFrame"`

	// 自截帧Base64编码数组。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BestFrames []*string `json:"BestFrames,omitempty" name:"BestFrames"`
}

type DetectInfoIdCardData struct {

	// OCR正面照片的base64编码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrFront *string `json:"OcrFront,omitempty" name:"OcrFront"`

	// OCR反面照片的base64编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrBack *string `json:"OcrBack,omitempty" name:"OcrBack"`

	// 旋转裁边后的正面照片base64编码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessedFrontImage *string `json:"ProcessedFrontImage,omitempty" name:"ProcessedFrontImage"`

	// 旋转裁边后的背面照片base64编码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessedBackImage *string `json:"ProcessedBackImage,omitempty" name:"ProcessedBackImage"`

	// 身份证正面人像图base64编码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Avatar *string `json:"Avatar,omitempty" name:"Avatar"`

	// 开启身份证防翻拍告警功能后才会返回，返回数组中可能出现的告警码如下：
	// -9102 身份证复印件告警。
	// -9103 身份证翻拍告警。
	// -9106 身份证 PS 告警。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WarnInfos []*int64 `json:"WarnInfos,omitempty" name:"WarnInfos"`
}

type DetectInfoText struct {

	// 本次流程最终验证结果。0为成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 本次流程最终验证结果描述。（仅描述用，文案更新时不会通知。）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 本次验证使用的身份证号。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdCard *string `json:"IdCard,omitempty" name:"IdCard"`

	// 本次验证使用的姓名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// Ocr识别结果。民族。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrNation *string `json:"OcrNation,omitempty" name:"OcrNation"`

	// Ocr识别结果。家庭住址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrAddress *string `json:"OcrAddress,omitempty" name:"OcrAddress"`

	// Ocr识别结果。生日。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrBirth *string `json:"OcrBirth,omitempty" name:"OcrBirth"`

	// Ocr识别结果。签发机关。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrAuthority *string `json:"OcrAuthority,omitempty" name:"OcrAuthority"`

	// Ocr识别结果。有效日期。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrValidDate *string `json:"OcrValidDate,omitempty" name:"OcrValidDate"`

	// Ocr识别结果。姓名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrName *string `json:"OcrName,omitempty" name:"OcrName"`

	// Ocr识别结果。身份证号。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrIdCard *string `json:"OcrIdCard,omitempty" name:"OcrIdCard"`

	// Ocr识别结果。性别。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrGender *string `json:"OcrGender,omitempty" name:"OcrGender"`

	// 本次流程最终活体结果。0为成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	LiveStatus *int64 `json:"LiveStatus,omitempty" name:"LiveStatus"`

	// 本次流程最终活体结果描述。（仅描述用，文案更新时不会通知。）
	// 注意：此字段可能返回 null，表示取不到有效值。
	LiveMsg *string `json:"LiveMsg,omitempty" name:"LiveMsg"`

	// 本次流程最终一比一结果。0为成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comparestatus *int64 `json:"Comparestatus,omitempty" name:"Comparestatus"`

	// 本次流程最终一比一结果描述。（仅描述用，文案更新时不会通知。）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comparemsg *string `json:"Comparemsg,omitempty" name:"Comparemsg"`

	// 本次流程活体一比一的分数，取值范围 [0.00, 100.00]。相似度大于等于70时才判断为同一人，也可根据具体场景自行调整阈值（阈值70的误通过率为千分之一，阈值80的误通过率是万分之一）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sim *string `json:"Sim,omitempty" name:"Sim"`

	// 地理位置经纬度。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Location *string `json:"Location,omitempty" name:"Location"`

	// Auth接口带入额外信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extra *string `json:"Extra,omitempty" name:"Extra"`

	// 本次流程进行的活体一比一流水。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LivenessDetail []*DetectDetail `json:"LivenessDetail,omitempty" name:"LivenessDetail"`

	// 手机号码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 本次流程最终比对库源类型。包括：
	// 权威库；
	// 业务方自有库（用户上传照片、客户的混合库、混合部署库）；
	// 二次验证库；
	// 人工审核库；
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompareLibType *string `json:"CompareLibType,omitempty" name:"CompareLibType"`
}

type DetectInfoVideoData struct {

	// 活体视频的base64编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	LivenessVideo *string `json:"LivenessVideo,omitempty" name:"LivenessVideo"`
}

type DetectReflectLivenessAndCompareRequest struct {
	*tchttp.BaseRequest

	// SDK生成的活体检测数据包的资源地址。
	LiveDataUrl *string `json:"LiveDataUrl,omitempty" name:"LiveDataUrl"`

	// SDK生成的活体检测数据包的资源内容MD5（32位，用于校验LiveData的一致性）。
	LiveDataMd5 *string `json:"LiveDataMd5,omitempty" name:"LiveDataMd5"`

	// 用于比对的目标图片的资源地址。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 用于比对的目标图片的资源MD5（32位，用于校验Image的一致性）。
	ImageMd5 *string `json:"ImageMd5,omitempty" name:"ImageMd5"`
}

func (r *DetectReflectLivenessAndCompareRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectReflectLivenessAndCompareRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LiveDataUrl")
	delete(f, "LiveDataMd5")
	delete(f, "ImageUrl")
	delete(f, "ImageMd5")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetectReflectLivenessAndCompareRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DetectReflectLivenessAndCompareResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 验证通过后的视频最佳截图资源临时地址，jpg格式，资源和链接有效期2小时，务必在有效期内下载。
		BestFrameUrl *string `json:"BestFrameUrl,omitempty" name:"BestFrameUrl"`

		// 验证通过后的视频最佳截图资源MD5（32位，用于校验BestFrame的一致性）。
		BestFrameMd5 *string `json:"BestFrameMd5,omitempty" name:"BestFrameMd5"`

		// 业务错误码，成功情况返回Success，错误情况请参考下方错误码 列表中FailedOperation部分。
		Result *string `json:"Result,omitempty" name:"Result"`

		// 业务结果描述。
		Description *string `json:"Description,omitempty" name:"Description"`

		// 相似度，取值范围 [0.00, 100.00]。推荐相似度大于等于70时可判断为同一人，可根据具体场景自行调整阈值（阈值70的误通过率为千分之一，阈值80的误通过率是万分之一）。
		Sim *float64 `json:"Sim,omitempty" name:"Sim"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetectReflectLivenessAndCompareResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectReflectLivenessAndCompareResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EidInfo struct {

	// 商户方 appeIDcode 的数字证书
	EidCode *string `json:"EidCode,omitempty" name:"EidCode"`

	// Eid中心针对商户方EidCode的电子签名
	EidSign *string `json:"EidSign,omitempty" name:"EidSign"`

	// 商户方公钥加密的会话密钥的base64字符串，[指引详见](https://cloud.tencent.com/document/product/1007/63370)
	DesKey *string `json:"DesKey,omitempty" name:"DesKey"`

	// 会话密钥sm2加密后的base64字符串，[指引详见](https://cloud.tencent.com/document/product/1007/63370)
	UserInfo *string `json:"UserInfo,omitempty" name:"UserInfo"`
}

type EncryptedPhoneVerificationRequest struct {
	*tchttp.BaseRequest

	// 身份证号，加密方式以EncryptionMode为准
	IdCard *string `json:"IdCard,omitempty" name:"IdCard"`

	// 姓名，加密方式以EncryptionMode为准
	Name *string `json:"Name,omitempty" name:"Name"`

	// 手机号，加密方式以EncryptionMode为准
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 敏感信息的加密方式，目前只支持MD5加密传输，参数取值：
	// 
	// 0：明文，不加密
	// 1：使用MD5加密
	EncryptionMode *string `json:"EncryptionMode,omitempty" name:"EncryptionMode"`
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

type EncryptedPhoneVerificationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 认证结果码:
	// 【收费结果码】
	// 0: 认证通过
	// -4: 信息不一致
	// 
	// 【不收费结果码】
	// -7: 身份证号码有误
	// -8: 参数错误
	// -9: 没有记录
	// -11: 验证中心服务繁忙
		Result *string `json:"Result,omitempty" name:"Result"`

		// 业务结果描述。
		Description *string `json:"Description,omitempty" name:"Description"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

	// 有加密需求的用户，接入传入kms的CiphertextBlob，关于数据加密可查阅<a href="https://cloud.tencent.com/document/product/1007/47180">数据加密</a> 文档。
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" name:"CiphertextBlob"`

	// 在使用加密服务时，填入要被加密的字段。本接口中可填入加密后的一个或多个字段
	EncryptList []*string `json:"EncryptList,omitempty" name:"EncryptList"`

	// 有加密需求的用户，传入CBC加密的初始向量
	Iv *string `json:"Iv,omitempty" name:"Iv"`
}

type GetActionSequenceRequest struct {
	*tchttp.BaseRequest

	// 默认不需要使用
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`
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

type GetActionSequenceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 动作顺序(2,1 or 1,2) 。1代表张嘴，2代表闭眼。
		ActionSequence *string `json:"ActionSequence,omitempty" name:"ActionSequence"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type GetDetectInfoEnhancedRequest struct {
	*tchttp.BaseRequest

	// 人脸核身流程的标识，调用DetectAuth接口时生成。
	BizToken *string `json:"BizToken,omitempty" name:"BizToken"`

	// 用于细分客户使用场景，由腾讯侧在线下对接时分配。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 指定拉取的结果信息，取值（0：全部；1：文本类；2：身份证信息；3：视频最佳截图信息）。
	// 如 13表示拉取文本类、视频最佳截图信息。
	// 默认值：0
	InfoType *string `json:"InfoType,omitempty" name:"InfoType"`

	// 从活体视频中截取一定张数的最佳帧（仅部分服务支持，若需使用请与慧眼小助手沟通）。默认为0，最大为10，超出10的最多只给10张。（InfoType需要包含3）
	BestFramesCount *uint64 `json:"BestFramesCount,omitempty" name:"BestFramesCount"`

	// 是否对身份证照片进行裁边。默认为false。（InfoType需要包含2）
	IsCutIdCardImage *bool `json:"IsCutIdCardImage,omitempty" name:"IsCutIdCardImage"`

	// 是否需要从身份证中抠出头像。默认为false。（InfoType需要包含2）
	IsNeedIdCardAvatar *bool `json:"IsNeedIdCardAvatar,omitempty" name:"IsNeedIdCardAvatar"`

	// 是否需要对返回中的敏感信息进行加密。其中敏感信息包括：Response.Text.IdCard、Response.Text.Name、Response.Text.OcrIdCard、Response.Text.OcrName
	IsEncrypt *bool `json:"IsEncrypt,omitempty" name:"IsEncrypt"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDetectInfoEnhancedRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetDetectInfoEnhancedResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 文本类信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Text *DetectInfoText `json:"Text,omitempty" name:"Text"`

		// 身份证照片信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
		IdCardData *DetectInfoIdCardData `json:"IdCardData,omitempty" name:"IdCardData"`

		// 最佳帧信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
		BestFrame *DetectInfoBestFrame `json:"BestFrame,omitempty" name:"BestFrame"`

		// 视频信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
		VideoData *DetectInfoVideoData `json:"VideoData,omitempty" name:"VideoData"`

		// 敏感数据加密信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Encryption *Encryption `json:"Encryption,omitempty" name:"Encryption"`

		// 意愿核身相关信息。若未使用意愿核身功能，该字段返回值可以不处理。
	// 注意：此字段可能返回 null，表示取不到有效值。
		IntentionVerifyData *IntentionVerifyData `json:"IntentionVerifyData,omitempty" name:"IntentionVerifyData"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type GetDetectInfoRequest struct {
	*tchttp.BaseRequest

	// 人脸核身流程的标识，调用DetectAuth接口时生成。
	BizToken *string `json:"BizToken,omitempty" name:"BizToken"`

	// 用于细分客户使用场景，申请开通服务后，可以在腾讯云慧眼人脸核身控制台（https://console.cloud.tencent.com/faceid） 自助接入里面创建，审核通过后即可调用。如有疑问，请加慧眼小助手微信（faceid001）进行咨询。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 指定拉取的结果信息，取值（0：全部；1：文本类；2：身份证正反面；3：视频最佳截图照片；4：视频）。
	// 如 134表示拉取文本类、视频最佳截图照片、视频。
	// 默认值：0
	InfoType *string `json:"InfoType,omitempty" name:"InfoType"`
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

type GetDetectInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
		DetectInfo *string `json:"DetectInfo,omitempty" name:"DetectInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type GetEidResultRequest struct {
	*tchttp.BaseRequest

	// E证通流程的唯一标识，调用GetEidToken接口时生成。
	EidToken *string `json:"EidToken,omitempty" name:"EidToken"`

	// 指定拉取的结果信息，取值（0：全部；1：文本类；2：身份证信息；3：最佳截图信息）。
	// 如 13表示拉取文本类、最佳截图信息。
	// 默认值：0
	InfoType *string `json:"InfoType,omitempty" name:"InfoType"`

	// 从活体视频中截取一定张数的最佳帧。默认为0，最大为3，超出3的最多只给3张。（InfoType需要包含3）
	BestFramesCount *uint64 `json:"BestFramesCount,omitempty" name:"BestFramesCount"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetEidResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetEidResultResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 文本类信息。（基于对敏感信息的保护，验证使用的姓名和身份证号统一通过加密后从Eidinfo参数中返回，如需获取请在控制台申请返回身份信息，详见[E证通获取实名信息指引](https://cloud.tencent.com/document/product/1007/63370)）
	// 注意：此字段可能返回 null，表示取不到有效值。
		Text *DetectInfoText `json:"Text,omitempty" name:"Text"`

		// 身份证照片信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
		IdCardData *DetectInfoIdCardData `json:"IdCardData,omitempty" name:"IdCardData"`

		// 最佳帧信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
		BestFrame *DetectInfoBestFrame `json:"BestFrame,omitempty" name:"BestFrame"`

		// Eid信息。（包括商户下用户唯一标识以及加密后的姓名、身份证号信息。解密方式详见[E证通获取实名信息指引](https://cloud.tencent.com/document/product/1007/63370)）
	// 注意：此字段可能返回 null，表示取不到有效值。
		EidInfo *EidInfo `json:"EidInfo,omitempty" name:"EidInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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
	InputType *string `json:"InputType,omitempty" name:"InputType"`
}

type GetEidTokenRequest struct {
	*tchttp.BaseRequest

	// EID商户id，字段长度最长50位。
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 身份标识（未使用OCR服务时，必须传入）。
	// 规则：a-zA-Z0-9组合。最长长度32位。
	IdCard *string `json:"IdCard,omitempty" name:"IdCard"`

	// 姓名。（未使用OCR服务时，必须传入）最长长度32位。中文请使用UTF-8编码。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 透传字段，在获取验证结果时返回。最长长度1024位。
	Extra *string `json:"Extra,omitempty" name:"Extra"`

	// 小程序模式配置，包括如何传入姓名身份证的配置。
	Config *GetEidTokenConfig `json:"Config,omitempty" name:"Config"`

	// 最长长度1024位。用户从Url中进入核身认证结束后重定向的回调链接地址。EidToken会在该链接的query参数中。
	RedirectUrl *string `json:"RedirectUrl,omitempty" name:"RedirectUrl"`

	// 敏感数据加密信息。对传入信息（姓名、身份证号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitempty" name:"Encryption"`
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

type GetEidTokenResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 一次核身流程的标识，有效时间为600秒；
	// 完成核身后，可用该标识获取验证结果信息。
		EidToken *string `json:"EidToken,omitempty" name:"EidToken"`

		// 发起核身流程的URL，用于H5场景核身。
		Url *string `json:"Url,omitempty" name:"Url"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type GetFaceIdResultRequest struct {
	*tchttp.BaseRequest

	// SDK人脸核身流程的标识，调用GetFaceIdToken接口时生成。
	FaceIdToken *string `json:"FaceIdToken,omitempty" name:"FaceIdToken"`

	// 是否需要拉取视频，默认false不需要
	IsNeedVideo *bool `json:"IsNeedVideo,omitempty" name:"IsNeedVideo"`

	// 是否需要拉取截帧，默认false不需要
	IsNeedBestFrame *bool `json:"IsNeedBestFrame,omitempty" name:"IsNeedBestFrame"`
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

type GetFaceIdResultResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 身份证
		IdCard *string `json:"IdCard,omitempty" name:"IdCard"`

		// 姓名
		Name *string `json:"Name,omitempty" name:"Name"`

		// 业务核验结果，参考https://cloud.tencent.com/document/product/1007/47912
		Result *string `json:"Result,omitempty" name:"Result"`

		// 业务核验描述
		Description *string `json:"Description,omitempty" name:"Description"`

		// 相似度，0-100，数值越大相似度越高
		Similarity *float64 `json:"Similarity,omitempty" name:"Similarity"`

		// 用户核验的视频base64，如果选择了使用cos，返回完整cos地址如https://bucket.cos.ap-guangzhou.myqcloud.com/objectKey
	// 注意：此字段可能返回 null，表示取不到有效值。
		VideoBase64 *string `json:"VideoBase64,omitempty" name:"VideoBase64"`

		// 用户核验视频的截帧base64，如果选择了使用cos，返回完整cos地址如https://bucket.cos.ap-guangzhou.myqcloud.com/objectKey
	// 注意：此字段可能返回 null，表示取不到有效值。
		BestFrameBase64 *string `json:"BestFrameBase64,omitempty" name:"BestFrameBase64"`

		// 获取token时透传的信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Extra *string `json:"Extra,omitempty" name:"Extra"`

		// 设备风险标签，仅错误码返回1007（设备疑似被劫持）时返回风险标签。标签说明：
	// 202、5001：设备疑似被Root
	// 203、5004：设备疑似被注入
	// 205：设备疑似被Hook
	// 206：设备疑似虚拟运行环境
	// 5007、1005：设备疑似摄像头被劫持
	// 8000：设备疑似存在异常篡改行为
	// 注意：此字段可能返回 null，表示取不到有效值。
		DeviceInfoTag *string `json:"DeviceInfoTag,omitempty" name:"DeviceInfoTag"`

		// 行为风险标签，仅错误码返回1007（设备疑似被劫持）时返回风险标签。标签说明：
	// 02：攻击风险
	// 注意：此字段可能返回 null，表示取不到有效值。
		RiskInfoTag *string `json:"RiskInfoTag,omitempty" name:"RiskInfoTag"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type GetFaceIdTokenRequest struct {
	*tchttp.BaseRequest

	// 本地上传照片(LOCAL)、商业库(BUSINESS)
	CompareLib *string `json:"CompareLib,omitempty" name:"CompareLib"`

	// CompareLib为商业库时必传。
	IdCard *string `json:"IdCard,omitempty" name:"IdCard"`

	// CompareLib为商业库库时必传。
	Name *string `json:"Name,omitempty" name:"Name"`

	// CompareLib为上传照片比对时必传，Base64后图片最大8MB。
	// 请使用标准的Base64编码方式(带=补位)，编码规范参考RFC4648。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// SDK中生成的Meta字符串
	Meta *string `json:"Meta,omitempty" name:"Meta"`

	// 透传参数 1000长度字符串
	Extra *string `json:"Extra,omitempty" name:"Extra"`

	// 默认为false，设置该参数为true后，核身过程中的视频图片将会存储在人脸核身控制台授权cos的bucket中，拉取结果时会返回对应资源完整cos地址。开通地址见https://console.cloud.tencent.com/faceid/cos
	// 【注意】选择该参数为true后将不返回base64数据，请根据接入情况谨慎修改。
	UseCos *bool `json:"UseCos,omitempty" name:"UseCos"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetFaceIdTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetFaceIdTokenResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 有效期 10分钟。只能完成1次核身。
		FaceIdToken *string `json:"FaceIdToken,omitempty" name:"FaceIdToken"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type GetLiveCodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数字验证码，如：1234
		LiveCode *string `json:"LiveCode,omitempty" name:"LiveCode"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type GetRealNameAuthResultRequest struct {
	*tchttp.BaseRequest

	// 实名认证凭证
	AuthToken *string `json:"AuthToken,omitempty" name:"AuthToken"`
}

func (r *GetRealNameAuthResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRealNameAuthResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AuthToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetRealNameAuthResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetRealNameAuthResultResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 认证结果码，收费情况如下：
	// 
	// 收费码：
	// 0:  姓名和身份证号一致
	// -1: 姓名和身份证号不一致
	// -2: 姓名和微信实名姓名不一致
	// 
	// 不收费码：
	// -3: 微信号未实名
		ResultType *string `json:"ResultType,omitempty" name:"ResultType"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRealNameAuthResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRealNameAuthResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRealNameAuthTokenRequest struct {
	*tchttp.BaseRequest

	// 姓名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 身份证号
	IDCard *string `json:"IDCard,omitempty" name:"IDCard"`

	// 回调地址。实名认证完成后，将会重定向到这个地址通知认证发起方。仅支持http或https协议。
	CallbackURL *string `json:"CallbackURL,omitempty" name:"CallbackURL"`
}

func (r *GetRealNameAuthTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRealNameAuthTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "IDCard")
	delete(f, "CallbackURL")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetRealNameAuthTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetRealNameAuthTokenResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询实名认证结果的唯一凭证
		AuthToken *string `json:"AuthToken,omitempty" name:"AuthToken"`

		// 实名认证授权地址，认证发起方需要重定向到这个地址获取认证用户的授权，仅能在微信环境下打开。
		RedirectURL *string `json:"RedirectURL,omitempty" name:"RedirectURL"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRealNameAuthTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRealNameAuthTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetWeChatBillDetailsRequest struct {
	*tchttp.BaseRequest

	// 拉取的日期（YYYY-MM-DD）。最大可追溯到365天前。当天6点后才能拉取前一天的数据。
	Date *string `json:"Date,omitempty" name:"Date"`

	// 游标。用于分页，取第一页时传0，取后续页面时，传入本接口响应中返回的NextCursor字段的值。
	Cursor *uint64 `json:"Cursor,omitempty" name:"Cursor"`

	// 需要拉取账单详情业务对应的RuleId。不传会返回所有RuleId数据。默认为空字符串。
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
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

type GetWeChatBillDetailsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 是否还有下一页。该字段为true时，需要将NextCursor的值作为入参Cursor继续调用本接口。
		HasNextPage *bool `json:"HasNextPage,omitempty" name:"HasNextPage"`

		// 下一页的游标。用于分页。
		NextCursor *uint64 `json:"NextCursor,omitempty" name:"NextCursor"`

		// 数据
		WeChatBillDetails []*WeChatBillDetail `json:"WeChatBillDetails,omitempty" name:"WeChatBillDetails"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type IdCardOCRVerificationRequest struct {
	*tchttp.BaseRequest

	// 身份证号
	// 姓名和身份证号、ImageBase64、ImageUrl三者必须提供其中之一。若都提供了，则按照姓名和身份证号>ImageBase64>ImageUrl的优先级使用参数。
	IdCard *string `json:"IdCard,omitempty" name:"IdCard"`

	// 姓名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 身份证人像面的 Base64 值
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过 3M。请使用标准的Base64编码方式(带=补位)，编码规范参考RFC4648。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 身份证人像面的 Url 地址
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持 GIF 格式。
	// 支持的图片大小：所下载图片经 Base64 编码后不超过 3M。图片下载时间不超过 3 秒。
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。
	// 非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 敏感数据加密信息。对传入信息（姓名、身份证号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitempty" name:"Encryption"`
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

type IdCardOCRVerificationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 认证结果码，收费情况如下。
	// 收费结果码：
	// 0: 姓名和身份证号一致
	// -1: 姓名和身份证号不一致
	// 不收费结果码：
	// -2: 非法身份证号（长度、校验位等不正确）
	// -3: 非法姓名（长度、格式等不正确）
	// -4: 证件库服务异常
	// -5: 证件库中无此身份证记录
		Result *string `json:"Result,omitempty" name:"Result"`

		// 业务结果描述。
		Description *string `json:"Description,omitempty" name:"Description"`

		// 用于验证的姓名
		Name *string `json:"Name,omitempty" name:"Name"`

		// 用于验证的身份证号
		IdCard *string `json:"IdCard,omitempty" name:"IdCard"`

		// OCR得到的性别
	// 注意：此字段可能返回 null，表示取不到有效值。
		Sex *string `json:"Sex,omitempty" name:"Sex"`

		// OCR得到的民族
	// 注意：此字段可能返回 null，表示取不到有效值。
		Nation *string `json:"Nation,omitempty" name:"Nation"`

		// OCR得到的生日
	// 注意：此字段可能返回 null，表示取不到有效值。
		Birth *string `json:"Birth,omitempty" name:"Birth"`

		// OCR得到的地址
	// 注意：此字段可能返回 null，表示取不到有效值。
		Address *string `json:"Address,omitempty" name:"Address"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type IdCardVerificationRequest struct {
	*tchttp.BaseRequest

	// 身份证号
	IdCard *string `json:"IdCard,omitempty" name:"IdCard"`

	// 姓名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 敏感数据加密信息。对传入信息（姓名、身份证号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitempty" name:"Encryption"`
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

type IdCardVerificationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 认证结果码，收费情况如下。
	// 收费结果码：
	// 0: 姓名和身份证号一致
	// -1: 姓名和身份证号不一致
	// 不收费结果码：
	// -2: 非法身份证号（长度、校验位等不正确）
	// -3: 非法姓名（长度、格式等不正确）
	// -4: 证件库服务异常
	// -5: 证件库中无此身份证记录
		Result *string `json:"Result,omitempty" name:"Result"`

		// 业务结果描述。
		Description *string `json:"Description,omitempty" name:"Description"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ImageRecognitionRequest struct {
	*tchttp.BaseRequest

	// 身份证号
	IdCard *string `json:"IdCard,omitempty" name:"IdCard"`

	// 姓名。中文请使用UTF-8编码。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 用于人脸比对的照片，图片的Base64值；
	// Base64编码后的图片数据大小不超过3M，仅支持jpg、png格式。
	// 请使用标准的Base64编码方式(带=补位)，编码规范参考RFC4648。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 本接口不需要传递此参数。
	Optional *string `json:"Optional,omitempty" name:"Optional"`

	// 敏感数据加密信息。对传入信息（姓名、身份证号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitempty" name:"Encryption"`
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

type ImageRecognitionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 相似度，取值范围 [0.00, 100.00]。推荐相似度大于等于70时可判断为同一人，可根据具体场景自行调整阈值（阈值70的误通过率为千分之一，阈值80的误通过率是万分之一）
		Sim *float64 `json:"Sim,omitempty" name:"Sim"`

		// 业务错误码，成功情况返回Success, 错误情况请参考下方错误码 列表中FailedOperation部分
		Result *string `json:"Result,omitempty" name:"Result"`

		// 业务结果描述。
		Description *string `json:"Description,omitempty" name:"Description"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type IntentionVerifyData struct {

	// 意愿确认环节中录制的视频（base64）。若不存在则为空字符串。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntentionVerifyVideo *string `json:"IntentionVerifyVideo,omitempty" name:"IntentionVerifyVideo"`

	// 意愿确认环节中用户语音转文字的识别结果。若不存在则为空字符串。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsrResult *string `json:"AsrResult,omitempty" name:"AsrResult"`

	// 意愿确认环节的结果码。当该结果码为0时，语音朗读的视频与语音识别结果才会返回。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`

	// 意愿确认环节的结果信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`

	// 意愿确认环节中录制视频的最佳帧（base64）。若不存在则为空字符串。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntentionVerifyBestFrame *string `json:"IntentionVerifyBestFrame,omitempty" name:"IntentionVerifyBestFrame"`
}

type LivenessCompareRequest struct {
	*tchttp.BaseRequest

	// 活体检测类型，取值：LIP/ACTION/SILENT。
	// LIP为数字模式，ACTION为动作模式，SILENT为静默模式，三种模式选择一种传入。
	LivenessType *string `json:"LivenessType,omitempty" name:"LivenessType"`

	// 用于人脸比对的照片的Base64值；
	// Base64编码后的图片数据大小不超过3M，仅支持jpg、png格式。
	// 请使用标准的Base64编码方式(带=补位)，编码规范参考RFC4648。
	// 
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageBase64。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 用于人脸比对照片的URL地址；图片下载后经Base64编码后的数据大小不超过3M，仅支持jpg、png格式。
	// 
	// 图片的 ImageUrl、ImageBase64 必须提供一个，如果都提供，只使用 ImageBase64。
	// 
	// 图片存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议图片存储于腾讯云。非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 数字模式传参：传数字验证码，验证码需先调用<a href="https://cloud.tencent.com/document/product/1007/31821">获取数字验证码接口</a>得到；
	// 动作模式传参：传动作顺序，动作顺序需先调用<a href="https://cloud.tencent.com/document/product/1007/31822">获取动作顺序接口</a>得到；
	// 静默模式传参：空。
	ValidateData *string `json:"ValidateData,omitempty" name:"ValidateData"`

	// 额外配置，传入JSON字符串。
	// {
	// "BestFrameNum": 2  //需要返回多张最佳截图，取值范围2-10
	// }
	Optional *string `json:"Optional,omitempty" name:"Optional"`

	// 用于活体检测的视频，视频的Base64值；
	// Base64编码后的大小不超过8M，支持mp4、avi、flv格式。
	// 请使用标准的Base64编码方式(带=补位)，编码规范参考RFC4648。
	// 
	// 视频的 VideoUrl、VideoBase64 必须提供一个，如果都提供，只使用 VideoBase64。
	VideoBase64 *string `json:"VideoBase64,omitempty" name:"VideoBase64"`

	// 用于活体检测的视频Url 地址。视频下载后经Base64编码后不超过 8M，视频下载耗时不超过4S，支持mp4、avi、flv格式。
	// 
	// 视频的 VideoUrl、VideoBase64 必须提供一个，如果都提供，只使用 VideoBase64。
	// 
	// 建议视频存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议视频存储于腾讯云。非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	VideoUrl *string `json:"VideoUrl,omitempty" name:"VideoUrl"`
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

type LivenessCompareResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 验证通过后的视频最佳截图照片，照片为BASE64编码后的值，jpg格式。
	// 注意：此字段可能返回 null，表示取不到有效值。
		BestFrameBase64 *string `json:"BestFrameBase64,omitempty" name:"BestFrameBase64"`

		// 相似度，取值范围 [0.00, 100.00]。推荐相似度大于等于70时可判断为同一人，可根据具体场景自行调整阈值（阈值70的误通过率为千分之一，阈值80的误通过率是万分之一）。
		Sim *float64 `json:"Sim,omitempty" name:"Sim"`

		// 业务错误码，成功情况返回Success, 错误情况请参考下方错误码 列表中FailedOperation部分
		Result *string `json:"Result,omitempty" name:"Result"`

		// 业务结果描述。
		Description *string `json:"Description,omitempty" name:"Description"`

		// 最佳截图列表，仅在配置了返回多张最佳截图时返回。
	// 注意：此字段可能返回 null，表示取不到有效值。
		BestFrameList []*string `json:"BestFrameList,omitempty" name:"BestFrameList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type LivenessRecognitionRequest struct {
	*tchttp.BaseRequest

	// 身份证号
	IdCard *string `json:"IdCard,omitempty" name:"IdCard"`

	// 姓名。中文请使用UTF-8编码。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 活体检测类型，取值：LIP/ACTION/SILENT。
	// LIP为数字模式，ACTION为动作模式，SILENT为静默模式，三种模式选择一种传入。
	LivenessType *string `json:"LivenessType,omitempty" name:"LivenessType"`

	// 用于活体检测的视频，视频的BASE64值；
	// BASE64编码后的大小不超过8M，支持mp4、avi、flv格式。
	VideoBase64 *string `json:"VideoBase64,omitempty" name:"VideoBase64"`

	// 用于活体检测的视频Url 地址。视频下载后经Base64编码不超过 8M，视频下载耗时不超过4S，支持mp4、avi、flv格式。
	// 
	// 视频的 VideoUrl、VideoBase64 必须提供一个，如果都提供，只使用 VideoBase64。
	// 
	// 建议视频存储于腾讯云的 Url 可保障更高的下载速度和稳定性，建议视频存储于腾讯云。非腾讯云存储的 Url 速度和稳定性可能受一定影响。
	VideoUrl *string `json:"VideoUrl,omitempty" name:"VideoUrl"`

	// 数字模式传参：传数字验证码，验证码需先调用<a href="https://cloud.tencent.com/document/product/1007/31821">获取数字验证码接口</a>得到；
	// 动作模式传参：传动作顺序，动作顺序需先调用<a href="https://cloud.tencent.com/document/product/1007/31822">获取动作顺序接口</a>得到；
	// 静默模式传参：空。
	ValidateData *string `json:"ValidateData,omitempty" name:"ValidateData"`

	// 额外配置，传入JSON字符串。
	// {
	// "BestFrameNum": 2  //需要返回多张最佳截图，取值范围2-10
	// }
	Optional *string `json:"Optional,omitempty" name:"Optional"`

	// 敏感数据加密信息。对传入信息（姓名、身份证号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitempty" name:"Encryption"`
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

type LivenessRecognitionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 验证通过后的视频最佳截图照片，照片为BASE64编码后的值，jpg格式。
	// 注意：此字段可能返回 null，表示取不到有效值。
		BestFrameBase64 *string `json:"BestFrameBase64,omitempty" name:"BestFrameBase64"`

		// 相似度，取值范围 [0.00, 100.00]。推荐相似度大于等于70时可判断为同一人，可根据具体场景自行调整阈值（阈值70的误通过率为千分之一，阈值80的误通过率是万分之一）
		Sim *float64 `json:"Sim,omitempty" name:"Sim"`

		// 业务错误码，成功情况返回Success, 错误情况请参考下方错误码 列表中FailedOperation部分
		Result *string `json:"Result,omitempty" name:"Result"`

		// 业务结果描述。
		Description *string `json:"Description,omitempty" name:"Description"`

		// 最佳截图列表，仅在配置了返回多张最佳截图时返回。
	// 注意：此字段可能返回 null，表示取不到有效值。
		BestFrameList []*string `json:"BestFrameList,omitempty" name:"BestFrameList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type LivenessRequest struct {
	*tchttp.BaseRequest

	// 用于活体检测的视频，视频的BASE64值；
	// BASE64编码后的大小不超过8M，支持mp4、avi、flv格式。
	VideoBase64 *string `json:"VideoBase64,omitempty" name:"VideoBase64"`

	// 活体检测类型，取值：LIP/ACTION/SILENT。
	// LIP为数字模式，ACTION为动作模式，SILENT为静默模式，三种模式选择一种传入。
	LivenessType *string `json:"LivenessType,omitempty" name:"LivenessType"`

	// 数字模式传参：数字验证码(1234)，需先调用接口获取数字验证码；
	// 动作模式传参：传动作顺序(2,1 or 1,2)，需先调用接口获取动作顺序；
	// 静默模式传参：不需要传递此参数。
	ValidateData *string `json:"ValidateData,omitempty" name:"ValidateData"`

	// 额外配置，传入JSON字符串。
	// {
	// "BestFrameNum": 2  //需要返回多张最佳截图，取值范围1-10
	// }
	Optional *string `json:"Optional,omitempty" name:"Optional"`
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

type LivenessResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 验证通过后的视频最佳截图照片，照片为BASE64编码后的值，jpg格式。
	// 注意：此字段可能返回 null，表示取不到有效值。
		BestFrameBase64 *string `json:"BestFrameBase64,omitempty" name:"BestFrameBase64"`

		// 业务错误码，成功情况返回Success, 错误情况请参考下方错误码 列表中FailedOperation部分
		Result *string `json:"Result,omitempty" name:"Result"`

		// 业务结果描述。
		Description *string `json:"Description,omitempty" name:"Description"`

		// 最佳最佳截图列表，仅在配置了返回多张最佳截图时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
		BestFrameList []*string `json:"BestFrameList,omitempty" name:"BestFrameList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type MinorsVerificationRequest struct {
	*tchttp.BaseRequest

	// 参与校验的参数类型。
	// 0：使用手机号进行校验；
	// 1：使用姓名与身份证号进行校验。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 手机号，11位数字，
	// 特别提示：
	// 手机号验证只限制在腾讯健康守护可信模型覆盖的数据范围内，与手机号本身在运营商是否实名无关联，不在范围会提示“手机号未实名”，建议客户与传入姓名和身份证号信息组合使用。
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 身份证号码。
	IdCard *string `json:"IdCard,omitempty" name:"IdCard"`

	// 姓名。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 敏感数据加密信息。对传入信息（姓名、身份证号、手机号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitempty" name:"Encryption"`
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

type MinorsVerificationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
		Result *string `json:"Result,omitempty" name:"Result"`

		// 业务结果描述。
		Description *string `json:"Description,omitempty" name:"Description"`

		// 该字段的值为年龄区间。格式为[a,b)，
	// [0,8)表示年龄小于8周岁区间，不包括8岁；
	// [8,16)表示年龄8-16周岁区间，不包括16岁；
	// [16,18)表示年龄16-18周岁区间，不包括18岁；
	// [18,+)表示年龄大于18周岁。
		AgeRange *string `json:"AgeRange,omitempty" name:"AgeRange"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type MobileNetworkTimeVerificationRequest struct {
	*tchttp.BaseRequest

	// 手机号码
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 敏感数据加密信息。对传入信息（手机号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitempty" name:"Encryption"`
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

type MobileNetworkTimeVerificationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 认证结果码，收费情况如下。
	// 收费结果码：
	// 0: 成功
	// -2: 手机号不存在
	// -3: 手机号存在，但无法查询到在网时长
	// 不收费结果码：
	// -1: 手机号格式不正确
	// -4: 验证中心服务繁忙
		Result *string `json:"Result,omitempty" name:"Result"`

		// 业务结果描述。
		Description *string `json:"Description,omitempty" name:"Description"`

		// 在网时长区间。
	// 格式为(a,b]，表示在网时长在a个月以上，b个月以下。若b为+时表示没有上限。
		Range *string `json:"Range,omitempty" name:"Range"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type MobileStatusRequest struct {
	*tchttp.BaseRequest

	// 手机号码
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 敏感数据加密信息。对传入信息（手机号）有加密需求的用户可使用此参数，详情请点击左侧链接。
	Encryption *Encryption `json:"Encryption,omitempty" name:"Encryption"`
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

type MobileStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 认证结果码，收费情况如下。
	// 收费结果码：
	// 0：成功
	// 不收费结果码：
	// -1：未查询到结果
	// -2：手机号格式不正确
	// -3：验证中心服务繁忙
		Result *string `json:"Result,omitempty" name:"Result"`

		// 业务结果描述。
		Description *string `json:"Description,omitempty" name:"Description"`

		// 状态码：
	// 0：正常
	// 1：停机
	// 2：销号
	// 3：空号
	// 4：不在网
	// 99：未知状态
		StatusCode *int64 `json:"StatusCode,omitempty" name:"StatusCode"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type PhoneVerificationRequest struct {
	*tchttp.BaseRequest

	// 身份证号
	IdCard *string `json:"IdCard,omitempty" name:"IdCard"`

	// 姓名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 手机号
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 有加密需求的用户，传入kms的CiphertextBlob，关于数据加密可查阅 <a href="https://cloud.tencent.com/document/product/1007/47180">数据加密</a> 文档。
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" name:"CiphertextBlob"`

	// 在使用加密服务时，填入要被加密的字段。本接口中可填入加密后的IdCard，Name，Phone中的一个或多个。
	EncryptList []*string `json:"EncryptList,omitempty" name:"EncryptList"`

	// 有加密需求的用户，传入CBC加密的初始向量。
	Iv *string `json:"Iv,omitempty" name:"Iv"`
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
	delete(f, "CiphertextBlob")
	delete(f, "EncryptList")
	delete(f, "Iv")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PhoneVerificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type PhoneVerificationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 认证结果码:
	// 收费结果码
	// 0: 认证通过
	// -4: 信息不一致（手机号已实名，但姓名和身份证号与实名信息不一致）
	// -5: 手机号未实名
	// 不收费结果码
	// -6: 手机号码不合法
	// -7: 身份证号码有误
	// -8: 姓名校验不通过
	// -9: 没有记录
	// -10: 认证未通过
	// -11: 验证中心服务繁忙
		Result *string `json:"Result,omitempty" name:"Result"`

		// 业务结果描述。
		Description *string `json:"Description,omitempty" name:"Description"`

		// 运营商名称。
	// 取值范围为["","移动","电信","联通"]
		Isp *string `json:"Isp,omitempty" name:"Isp"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type WeChatBillDetail struct {

	// token
	BizToken *string `json:"BizToken,omitempty" name:"BizToken"`

	// 本token收费次数
	ChargeCount *uint64 `json:"ChargeCount,omitempty" name:"ChargeCount"`

	// 本token计费详情
	ChargeDetails []*ChargeDetail `json:"ChargeDetails,omitempty" name:"ChargeDetails"`

	// 业务RuleId
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}
