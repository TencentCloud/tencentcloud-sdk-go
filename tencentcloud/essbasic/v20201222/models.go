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

package v20201222

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Address struct {
	// 省份
	Province *string `json:"Province,omitempty" name:"Province"`

	// 城市
	City *string `json:"City,omitempty" name:"City"`

	// 区县
	County *string `json:"County,omitempty" name:"County"`

	// 详细地址
	Details *string `json:"Details,omitempty" name:"Details"`

	// 国家，默认中国
	Country *string `json:"Country,omitempty" name:"Country"`
}

// Predefined struct for user
type ArchiveFlowRequestParams struct {
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 流程ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
}

type ArchiveFlowRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 流程ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
}

func (r *ArchiveFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ArchiveFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "FlowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ArchiveFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ArchiveFlowResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ArchiveFlowResponse struct {
	*tchttp.BaseResponse
	Response *ArchiveFlowResponseParams `json:"Response"`
}

func (r *ArchiveFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ArchiveFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Caller struct {
	// 应用号
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 下属机构ID
	SubOrganizationId *string `json:"SubOrganizationId,omitempty" name:"SubOrganizationId"`

	// 经办人的用户ID
	OperatorId *string `json:"OperatorId,omitempty" name:"OperatorId"`
}

// Predefined struct for user
type CancelFlowRequestParams struct {
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 流程ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 撤销原因
	CancelMessage *string `json:"CancelMessage,omitempty" name:"CancelMessage"`
}

type CancelFlowRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 流程ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 撤销原因
	CancelMessage *string `json:"CancelMessage,omitempty" name:"CancelMessage"`
}

func (r *CancelFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "FlowId")
	delete(f, "CancelMessage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelFlowResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CancelFlowResponse struct {
	*tchttp.BaseResponse
	Response *CancelFlowResponseParams `json:"Response"`
}

func (r *CancelFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CatalogApprovers struct {
	// 流程ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 参与者列表
	Approvers []*FlowApproverInfo `json:"Approvers,omitempty" name:"Approvers"`
}

type CatalogComponents struct {
	// 流程ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 签署区列表
	SignComponents []*Component `json:"SignComponents,omitempty" name:"SignComponents"`

	// 签署任务ID
	SignId *string `json:"SignId,omitempty" name:"SignId"`
}

// Predefined struct for user
type CheckBankCard2EVerificationRequestParams struct {
	// 调用方信息; 必选
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 银行卡号
	BankCard *string `json:"BankCard,omitempty" name:"BankCard"`

	// 姓名
	Name *string `json:"Name,omitempty" name:"Name"`
}

type CheckBankCard2EVerificationRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息; 必选
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 银行卡号
	BankCard *string `json:"BankCard,omitempty" name:"BankCard"`

	// 姓名
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *CheckBankCard2EVerificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckBankCard2EVerificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "BankCard")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckBankCard2EVerificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckBankCard2EVerificationResponseParams struct {
	// 检测结果
	// 计费结果码：
	//   0:  认证通过
	//   1:  认证未通过
	//   2:  持卡人信息有误
	//   3:  未开通无卡支付
	//   4:  此卡被没收
	//   5:  无效卡号
	//   6:  此卡无对应发卡行
	//   7:  该卡未初始化或睡眠卡
	//   8:  作弊卡、吞卡
	//   9:  此卡已挂失
	//   10: 该卡已过期
	//   11: 受限制的卡
	//   12: 密码错误次数超限
	//   13: 发卡行不支持此交易
	// 不收费结果码:
	//   101: 姓名校验不通过
	//   102: 银行卡号码有误
	//   103: 验证中心服务繁忙
	//   104: 身份证号码有误
	//   105: 手机号码不合法
	Result *int64 `json:"Result,omitempty" name:"Result"`

	// 结果描述; 未通过时必选
	Description *string `json:"Description,omitempty" name:"Description"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CheckBankCard2EVerificationResponse struct {
	*tchttp.BaseResponse
	Response *CheckBankCard2EVerificationResponseParams `json:"Response"`
}

func (r *CheckBankCard2EVerificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckBankCard2EVerificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckBankCard3EVerificationRequestParams struct {
	// 调用方信息; 必选
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 银行卡号
	BankCard *string `json:"BankCard,omitempty" name:"BankCard"`

	// 姓名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 身份证件号码
	IdCardNumber *string `json:"IdCardNumber,omitempty" name:"IdCardNumber"`

	// 身份证件类型; ID_CARD
	IdCardType *string `json:"IdCardType,omitempty" name:"IdCardType"`
}

type CheckBankCard3EVerificationRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息; 必选
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 银行卡号
	BankCard *string `json:"BankCard,omitempty" name:"BankCard"`

	// 姓名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 身份证件号码
	IdCardNumber *string `json:"IdCardNumber,omitempty" name:"IdCardNumber"`

	// 身份证件类型; ID_CARD
	IdCardType *string `json:"IdCardType,omitempty" name:"IdCardType"`
}

func (r *CheckBankCard3EVerificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckBankCard3EVerificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "BankCard")
	delete(f, "Name")
	delete(f, "IdCardNumber")
	delete(f, "IdCardType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckBankCard3EVerificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckBankCard3EVerificationResponseParams struct {
	// 检测结果
	// 计费结果码：
	//   0:  认证通过
	//   1:  认证未通过
	//   2:  持卡人信息有误
	//   3:  未开通无卡支付
	//   4:  此卡被没收
	//   5:  无效卡号
	//   6:  此卡无对应发卡行
	//   7:  该卡未初始化或睡眠卡
	//   8:  作弊卡、吞卡
	//   9:  此卡已挂失
	//   10: 该卡已过期
	//   11: 受限制的卡
	//   12: 密码错误次数超限
	//   13: 发卡行不支持此交易
	// 不收费结果码:
	//   101: 姓名校验不通过
	//   102: 银行卡号码有误
	//   103: 验证中心服务繁忙
	//   104: 身份证号码有误
	//   105: 手机号码不合法
	Result *int64 `json:"Result,omitempty" name:"Result"`

	// 结果描述; 未通过时必选
	Description *string `json:"Description,omitempty" name:"Description"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CheckBankCard3EVerificationResponse struct {
	*tchttp.BaseResponse
	Response *CheckBankCard3EVerificationResponseParams `json:"Response"`
}

func (r *CheckBankCard3EVerificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckBankCard3EVerificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckBankCard4EVerificationRequestParams struct {
	// 调用方信息; 必选
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 银行卡号
	BankCard *string `json:"BankCard,omitempty" name:"BankCard"`

	// 姓名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 身份证件号码
	IdCardNumber *string `json:"IdCardNumber,omitempty" name:"IdCardNumber"`

	// 手机号
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 身份证件类型; ID_CARD
	IdCardType *string `json:"IdCardType,omitempty" name:"IdCardType"`
}

type CheckBankCard4EVerificationRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息; 必选
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 银行卡号
	BankCard *string `json:"BankCard,omitempty" name:"BankCard"`

	// 姓名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 身份证件号码
	IdCardNumber *string `json:"IdCardNumber,omitempty" name:"IdCardNumber"`

	// 手机号
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 身份证件类型; ID_CARD
	IdCardType *string `json:"IdCardType,omitempty" name:"IdCardType"`
}

func (r *CheckBankCard4EVerificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckBankCard4EVerificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "BankCard")
	delete(f, "Name")
	delete(f, "IdCardNumber")
	delete(f, "Mobile")
	delete(f, "IdCardType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckBankCard4EVerificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckBankCard4EVerificationResponseParams struct {
	// 检测结果
	// 计费结果码：
	//   0:  认证通过
	//   1:  认证未通过
	//   2:  持卡人信息有误
	//   3:  未开通无卡支付
	//   4:  此卡被没收
	//   5:  无效卡号
	//   6:  此卡无对应发卡行
	//   7:  该卡未初始化或睡眠卡
	//   8:  作弊卡、吞卡
	//   9:  此卡已挂失
	//   10: 该卡已过期
	//   11: 受限制的卡
	//   12: 密码错误次数超限
	//   13: 发卡行不支持此交易
	// 不收费结果码:
	//   101: 姓名校验不通过
	//   102: 银行卡号码有误
	//   103: 验证中心服务繁忙
	//   104: 身份证号码有误
	//   105: 手机号码不合法
	Result *int64 `json:"Result,omitempty" name:"Result"`

	// 结果描述; 未通过时必选
	Description *string `json:"Description,omitempty" name:"Description"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CheckBankCard4EVerificationResponse struct {
	*tchttp.BaseResponse
	Response *CheckBankCard4EVerificationResponseParams `json:"Response"`
}

func (r *CheckBankCard4EVerificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckBankCard4EVerificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckBankCardVerificationRequestParams struct {
	// 调用方信息; 必选
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 银行卡号
	BankCard *string `json:"BankCard,omitempty" name:"BankCard"`

	// 姓名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 身份证件号码
	IdCardNumber *string `json:"IdCardNumber,omitempty" name:"IdCardNumber"`

	// 手机号
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 身份证件类型; ID_CARD
	IdCardType *string `json:"IdCardType,omitempty" name:"IdCardType"`
}

type CheckBankCardVerificationRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息; 必选
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 银行卡号
	BankCard *string `json:"BankCard,omitempty" name:"BankCard"`

	// 姓名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 身份证件号码
	IdCardNumber *string `json:"IdCardNumber,omitempty" name:"IdCardNumber"`

	// 手机号
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 身份证件类型; ID_CARD
	IdCardType *string `json:"IdCardType,omitempty" name:"IdCardType"`
}

func (r *CheckBankCardVerificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckBankCardVerificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "BankCard")
	delete(f, "Name")
	delete(f, "IdCardNumber")
	delete(f, "Mobile")
	delete(f, "IdCardType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckBankCardVerificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckBankCardVerificationResponseParams struct {
	// 检测结果
	// 计费结果码：
	//   0:  认证通过
	//   1:  认证未通过
	//   2:  持卡人信息有误
	//   3:  未开通无卡支付
	//   4:  此卡被没收
	//   5:  无效卡号
	//   6:  此卡无对应发卡行
	//   7:  该卡未初始化或睡眠卡
	//   8:  作弊卡、吞卡
	//   9:  此卡已挂失
	//   10: 该卡已过期
	//   11: 受限制的卡
	//   12: 密码错误次数超限
	//   13: 发卡行不支持此交易
	// 不收费结果码:
	//   101: 姓名校验不通过
	//   102: 银行卡号码有误
	//   103: 验证中心服务繁忙
	//   104: 身份证号码有误
	//   105: 手机号码不合法
	Result *int64 `json:"Result,omitempty" name:"Result"`

	// 结果描述; 未通过时必选
	Description *string `json:"Description,omitempty" name:"Description"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CheckBankCardVerificationResponse struct {
	*tchttp.BaseResponse
	Response *CheckBankCardVerificationResponseParams `json:"Response"`
}

func (r *CheckBankCardVerificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckBankCardVerificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckFaceIdentifyRequestParams struct {
	// 调用方信息; 必选
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 人脸核身渠道; 必选; WEIXINAPP:腾讯电子签小程序,FACEID:腾讯电子签慧眼,None:白名单中的客户直接通过
	VerifyChannel *string `json:"VerifyChannel,omitempty" name:"VerifyChannel"`

	// 核身订单号; 必选; 对于WEIXINAPP,直接取响应的{VerifyResult};对于FACEID,使用{WbAppId}:{OrderNo}拼接
	VerifyResult *string `json:"VerifyResult,omitempty" name:"VerifyResult"`

	// 要对比的姓名; 可选; 未填写时对比caller.OperatorId的实名信息
	Name *string `json:"Name,omitempty" name:"Name"`

	// 要对比的身份证号码; 可选; 未填写时对比caller.OperatorId的实名信息
	IdCardNumber *string `json:"IdCardNumber,omitempty" name:"IdCardNumber"`

	// 是否取认证时的照片
	GetPhoto *bool `json:"GetPhoto,omitempty" name:"GetPhoto"`
}

type CheckFaceIdentifyRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息; 必选
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 人脸核身渠道; 必选; WEIXINAPP:腾讯电子签小程序,FACEID:腾讯电子签慧眼,None:白名单中的客户直接通过
	VerifyChannel *string `json:"VerifyChannel,omitempty" name:"VerifyChannel"`

	// 核身订单号; 必选; 对于WEIXINAPP,直接取响应的{VerifyResult};对于FACEID,使用{WbAppId}:{OrderNo}拼接
	VerifyResult *string `json:"VerifyResult,omitempty" name:"VerifyResult"`

	// 要对比的姓名; 可选; 未填写时对比caller.OperatorId的实名信息
	Name *string `json:"Name,omitempty" name:"Name"`

	// 要对比的身份证号码; 可选; 未填写时对比caller.OperatorId的实名信息
	IdCardNumber *string `json:"IdCardNumber,omitempty" name:"IdCardNumber"`

	// 是否取认证时的照片
	GetPhoto *bool `json:"GetPhoto,omitempty" name:"GetPhoto"`
}

func (r *CheckFaceIdentifyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckFaceIdentifyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "VerifyChannel")
	delete(f, "VerifyResult")
	delete(f, "Name")
	delete(f, "IdCardNumber")
	delete(f, "GetPhoto")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckFaceIdentifyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckFaceIdentifyResponseParams struct {
	// 核身结果; 0:通过,1:不通过
	Result *int64 `json:"Result,omitempty" name:"Result"`

	// 核身结果描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 渠道名
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 认证通过时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerifiedOn *int64 `json:"VerifiedOn,omitempty" name:"VerifiedOn"`

	// 核身流水号
	SerialNumber *string `json:"SerialNumber,omitempty" name:"SerialNumber"`

	// 渠道核身服务器IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerifyServerIp *string `json:"VerifyServerIp,omitempty" name:"VerifyServerIp"`

	// 核身照片文件名
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhotoFileName *string `json:"PhotoFileName,omitempty" name:"PhotoFileName"`

	// 核身照片内容base64(文件格式见文件名后缀,一般为jpg)
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhotoFileData *string `json:"PhotoFileData,omitempty" name:"PhotoFileData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CheckFaceIdentifyResponse struct {
	*tchttp.BaseResponse
	Response *CheckFaceIdentifyResponseParams `json:"Response"`
}

func (r *CheckFaceIdentifyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckFaceIdentifyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckIdCardVerificationRequestParams struct {
	// 调用方信息; 必选
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 姓名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 身份证件号码
	IdCardNumber *string `json:"IdCardNumber,omitempty" name:"IdCardNumber"`

	// 身份证件类型; ID_CARD
	IdCardType *string `json:"IdCardType,omitempty" name:"IdCardType"`
}

type CheckIdCardVerificationRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息; 必选
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 姓名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 身份证件号码
	IdCardNumber *string `json:"IdCardNumber,omitempty" name:"IdCardNumber"`

	// 身份证件类型; ID_CARD
	IdCardType *string `json:"IdCardType,omitempty" name:"IdCardType"`
}

func (r *CheckIdCardVerificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckIdCardVerificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "Name")
	delete(f, "IdCardNumber")
	delete(f, "IdCardType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckIdCardVerificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckIdCardVerificationResponseParams struct {
	// 检测结果; 
	// 收费错误码:
	//   0: 通过,
	//   1: 姓名和身份证号不一致,
	// 免费错误码:
	//   101: 非法身份证号(长度,格式等不正确),
	//   102: 非法姓名(长度,格式等不正确),
	//   103: 验证平台异常,
	//   104: 证件库中无此身份证记录
	Result *int64 `json:"Result,omitempty" name:"Result"`

	// 结果描述; 未通过时必选
	Description *string `json:"Description,omitempty" name:"Description"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CheckIdCardVerificationResponse struct {
	*tchttp.BaseResponse
	Response *CheckIdCardVerificationResponseParams `json:"Response"`
}

func (r *CheckIdCardVerificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckIdCardVerificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckMobileAndNameRequestParams struct {
	// 调用方信息; 必选
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 手机号
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 姓名
	Name *string `json:"Name,omitempty" name:"Name"`
}

type CheckMobileAndNameRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息; 必选
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 手机号
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 姓名
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *CheckMobileAndNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckMobileAndNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "Mobile")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckMobileAndNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckMobileAndNameResponseParams struct {
	// 检测结果
	// 计费结果码：
	//   0:  验证结果一致
	//   1:  手机号未实名
	//   2:  姓名和手机号不一致
	//   3:  信息不一致(手机号已实名,但姓名和身份证号与实名信息不一致)
	// 不收费结果码:
	//   101: 查无记录
	//   102: 非法姓名(长度,格式等不正确)
	//   103: 非法手机号(长度,格式等不正确)
	//   104: 非法身份证号(长度,校验位等不正确)
	//   105: 认证未通过
	//   106: 验证平台异常
	Result *int64 `json:"Result,omitempty" name:"Result"`

	// 结果描述; 未通过时必选
	Description *string `json:"Description,omitempty" name:"Description"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CheckMobileAndNameResponse struct {
	*tchttp.BaseResponse
	Response *CheckMobileAndNameResponseParams `json:"Response"`
}

func (r *CheckMobileAndNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckMobileAndNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckMobileVerificationRequestParams struct {
	// 调用方信息; 必选
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 手机号
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 姓名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 身份证件号码
	IdCardNumber *string `json:"IdCardNumber,omitempty" name:"IdCardNumber"`

	// 身份证件类型; ID_CARD
	IdCardType *string `json:"IdCardType,omitempty" name:"IdCardType"`
}

type CheckMobileVerificationRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息; 必选
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 手机号
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 姓名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 身份证件号码
	IdCardNumber *string `json:"IdCardNumber,omitempty" name:"IdCardNumber"`

	// 身份证件类型; ID_CARD
	IdCardType *string `json:"IdCardType,omitempty" name:"IdCardType"`
}

func (r *CheckMobileVerificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckMobileVerificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "Mobile")
	delete(f, "Name")
	delete(f, "IdCardNumber")
	delete(f, "IdCardType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckMobileVerificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckMobileVerificationResponseParams struct {
	// 检测结果
	// 计费结果码：
	//   0:  验证结果一致
	//   1:  手机号未实名
	//   2:  姓名和手机号不一致
	//   3:  信息不一致(手机号已实名,但姓名和身份证号与实名信息不一致)
	// 不收费结果码:
	//   101: 查无记录
	//   102: 非法姓名(长度,格式等不正确)
	//   103: 非法手机号(长度,格式等不正确)
	//   104: 非法身份证号(长度,校验位等不正确)
	//   105: 认证未通过
	//   106: 验证平台异常
	Result *int64 `json:"Result,omitempty" name:"Result"`

	// 结果描述; 未通过时必选
	Description *string `json:"Description,omitempty" name:"Description"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CheckMobileVerificationResponse struct {
	*tchttp.BaseResponse
	Response *CheckMobileVerificationResponseParams `json:"Response"`
}

func (r *CheckMobileVerificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckMobileVerificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckVerifyCodeMatchFlowIdRequestParams struct {
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 手机号
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 验证码
	VerifyCode *string `json:"VerifyCode,omitempty" name:"VerifyCode"`

	// 流程(目录) id
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
}

type CheckVerifyCodeMatchFlowIdRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 手机号
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 验证码
	VerifyCode *string `json:"VerifyCode,omitempty" name:"VerifyCode"`

	// 流程(目录) id
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
}

func (r *CheckVerifyCodeMatchFlowIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckVerifyCodeMatchFlowIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "Mobile")
	delete(f, "VerifyCode")
	delete(f, "FlowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckVerifyCodeMatchFlowIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckVerifyCodeMatchFlowIdResponseParams struct {
	// true: 验证码正确，false: 验证码错误
	Success *bool `json:"Success,omitempty" name:"Success"`

	// 0: 验证码正确 1:验证码错误或过期 2:验证码错误 3:验证码和流程不匹配 4:验证码输入错误超过次数 5:内部错误
	// 6:参数错误
	Result *int64 `json:"Result,omitempty" name:"Result"`

	// 结果描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CheckVerifyCodeMatchFlowIdResponse struct {
	*tchttp.BaseResponse
	Response *CheckVerifyCodeMatchFlowIdResponseParams `json:"Response"`
}

func (r *CheckVerifyCodeMatchFlowIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckVerifyCodeMatchFlowIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Component struct {
	// 控件编号
	// 
	// 注：
	// 当GenerateMode=3时，通过"^"来决定是否使用关键字整词匹配能力。
	// 例：
	// 当GenerateMode=3时，如果传入关键字"^甲方签署^"，则会在PDF文件中有且仅有"甲方签署"关键字的地方进行对应操作。
	// 如传入的关键字为"甲方签署"，则PDF文件中每个出现关键字的位置都会执行相应操作。
	ComponentId *string `json:"ComponentId,omitempty" name:"ComponentId"`

	// 如果是Component控件类型，则可选的字段为：
	// TEXT - 普通文本控件；
	// DATE - 普通日期控件；
	// SELECT- 勾选框控件；
	// 如果是SignComponent控件类型，则可选的字段为
	// SIGN_SEAL- 签署印章控件；
	// SIGN_DATE- 签署日期控件；
	// SIGN_SIGNATURE - 用户签名控件；
	ComponentType *string `json:"ComponentType,omitempty" name:"ComponentType"`

	// 控件名称
	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`

	// 定义控件是否为必填项，默认为false
	ComponentRequired *bool `json:"ComponentRequired,omitempty" name:"ComponentRequired"`

	// 控件所属文件的序号 (模板中的resourceId排列序号)
	FileIndex *int64 `json:"FileIndex,omitempty" name:"FileIndex"`

	// 控件生成的方式：
	// 0 - 普通控件
	// 1 - 表单域
	// 2 - html 控件
	// 3 - 关键字
	GenerateMode *int64 `json:"GenerateMode,omitempty" name:"GenerateMode"`

	// 参数控件宽度，单位px
	ComponentWidth *float64 `json:"ComponentWidth,omitempty" name:"ComponentWidth"`

	// 参数控件高度，单位px
	ComponentHeight *float64 `json:"ComponentHeight,omitempty" name:"ComponentHeight"`

	// 参数控件所在页码
	ComponentPage *int64 `json:"ComponentPage,omitempty" name:"ComponentPage"`

	// 参数控件X位置，单位px
	ComponentPosX *float64 `json:"ComponentPosX,omitempty" name:"ComponentPosX"`

	// 参数控件Y位置，单位px
	ComponentPosY *float64 `json:"ComponentPosY,omitempty" name:"ComponentPosY"`

	// 参数控件样式
	ComponentExtra *string `json:"ComponentExtra,omitempty" name:"ComponentExtra"`

	// 印章ID，如果是手写签名则为jpg或png格式的base64图片
	// 
	// SIGN_SEAL控件,可以用ORG_DEFAULT_SEAL表示主企业的默认印章
	// SIGN_SEAL控件,可以用SUBORG_DEFAULT_SEAL表示子企业的默认印章
	// SIGN_SEAL控件,可以用USER_DEFAULT_SEAL表示个人默认印章
	ComponentValue *string `json:"ComponentValue,omitempty" name:"ComponentValue"`

	// 如果是SIGN_SEAL类型的签署控件, 参数标识H5签署界面是否在该签署区上进行放置展示, 1为放置,其他为不放置
	SealOperate *int64 `json:"SealOperate,omitempty" name:"SealOperate"`

	// 不同GenerateMode对应的额外信息
	GenerateExtra *string `json:"GenerateExtra,omitempty" name:"GenerateExtra"`
}

type ComponentSeal struct {
	// 签署区ID
	ComponentId *string `json:"ComponentId,omitempty" name:"ComponentId"`

	// 印章ID
	SealId *string `json:"SealId,omitempty" name:"SealId"`
}

// Predefined struct for user
type CreateFaceIdSignRequestParams struct {
	// 调用方信息; 必选
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 除api_ticket之外的其它要参与签名的参数值,包括UserId
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type CreateFaceIdSignRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息; 必选
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 除api_ticket之外的其它要参与签名的参数值,包括UserId
	Values []*string `json:"Values,omitempty" name:"Values"`
}

func (r *CreateFaceIdSignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFaceIdSignRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "Values")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFaceIdSignRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFaceIdSignResponseParams struct {
	// 慧眼API签名
	Sign *string `json:"Sign,omitempty" name:"Sign"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateFaceIdSignResponse struct {
	*tchttp.BaseResponse
	Response *CreateFaceIdSignResponseParams `json:"Response"`
}

func (r *CreateFaceIdSignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFaceIdSignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowByFilesRequestParams struct {
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 流程创建信息
	FlowInfo *FlowInfo `json:"FlowInfo,omitempty" name:"FlowInfo"`

	// 文件资源列表 (支持多文件)
	FileIds []*string `json:"FileIds,omitempty" name:"FileIds"`

	// 自定义流程id
	CustomId *string `json:"CustomId,omitempty" name:"CustomId"`
}

type CreateFlowByFilesRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 流程创建信息
	FlowInfo *FlowInfo `json:"FlowInfo,omitempty" name:"FlowInfo"`

	// 文件资源列表 (支持多文件)
	FileIds []*string `json:"FileIds,omitempty" name:"FileIds"`

	// 自定义流程id
	CustomId *string `json:"CustomId,omitempty" name:"CustomId"`
}

func (r *CreateFlowByFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowByFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "FlowInfo")
	delete(f, "FileIds")
	delete(f, "CustomId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFlowByFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowByFilesResponseParams struct {
	// 流程ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateFlowByFilesResponse struct {
	*tchttp.BaseResponse
	Response *CreateFlowByFilesResponseParams `json:"Response"`
}

func (r *CreateFlowByFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowByFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateH5FaceIdUrlRequestParams struct {
	// 调用方信息; 必选
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 慧眼业务ID; 不填写时后台使用Caller反查
	WbAppId *string `json:"WbAppId,omitempty" name:"WbAppId"`

	// 姓名; 可选(未通过实名认证的用户必选)
	Name *string `json:"Name,omitempty" name:"Name"`

	// 用户证件类型; 可选; 默认ID_CARD:中国居民身份证
	IdCardType *string `json:"IdCardType,omitempty" name:"IdCardType"`

	// 用户证件号; 可选(未通过实名认证的用户必选)
	IdCardNumber *string `json:"IdCardNumber,omitempty" name:"IdCardNumber"`

	// H5人脸核身完成后回调的第三方Url; 可选; 不需要做Encode, 跳转的参数: ?code=XX&orderNo=XX&liveRate=xx, code=0表示成功,orderNo为订单号,liveRate为百分制活体检测得分
	JumpUrl *string `json:"JumpUrl,omitempty" name:"JumpUrl"`

	// 参数值为"1":直接跳转到url回调地址; 可选; 其他值:跳转提供的结果页面
	JumpType *string `json:"JumpType,omitempty" name:"JumpType"`

	// browser:表示在浏览器启动刷脸, app:表示在App里启动刷脸,默认值为browser; 可选
	OpenFrom *string `json:"OpenFrom,omitempty" name:"OpenFrom"`

	// 跳转类型; 可选; 参数值为"1"时,刷脸页面使用replace方式跳转,不在浏览器history中留下记录;不传或其他值则正常跳转
	RedirectType *string `json:"RedirectType,omitempty" name:"RedirectType"`
}

type CreateH5FaceIdUrlRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息; 必选
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 慧眼业务ID; 不填写时后台使用Caller反查
	WbAppId *string `json:"WbAppId,omitempty" name:"WbAppId"`

	// 姓名; 可选(未通过实名认证的用户必选)
	Name *string `json:"Name,omitempty" name:"Name"`

	// 用户证件类型; 可选; 默认ID_CARD:中国居民身份证
	IdCardType *string `json:"IdCardType,omitempty" name:"IdCardType"`

	// 用户证件号; 可选(未通过实名认证的用户必选)
	IdCardNumber *string `json:"IdCardNumber,omitempty" name:"IdCardNumber"`

	// H5人脸核身完成后回调的第三方Url; 可选; 不需要做Encode, 跳转的参数: ?code=XX&orderNo=XX&liveRate=xx, code=0表示成功,orderNo为订单号,liveRate为百分制活体检测得分
	JumpUrl *string `json:"JumpUrl,omitempty" name:"JumpUrl"`

	// 参数值为"1":直接跳转到url回调地址; 可选; 其他值:跳转提供的结果页面
	JumpType *string `json:"JumpType,omitempty" name:"JumpType"`

	// browser:表示在浏览器启动刷脸, app:表示在App里启动刷脸,默认值为browser; 可选
	OpenFrom *string `json:"OpenFrom,omitempty" name:"OpenFrom"`

	// 跳转类型; 可选; 参数值为"1"时,刷脸页面使用replace方式跳转,不在浏览器history中留下记录;不传或其他值则正常跳转
	RedirectType *string `json:"RedirectType,omitempty" name:"RedirectType"`
}

func (r *CreateH5FaceIdUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateH5FaceIdUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "WbAppId")
	delete(f, "Name")
	delete(f, "IdCardType")
	delete(f, "IdCardNumber")
	delete(f, "JumpUrl")
	delete(f, "JumpType")
	delete(f, "OpenFrom")
	delete(f, "RedirectType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateH5FaceIdUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateH5FaceIdUrlResponseParams struct {
	// 跳转到人脸核身页面的链接
	Url *string `json:"Url,omitempty" name:"Url"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateH5FaceIdUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreateH5FaceIdUrlResponseParams `json:"Response"`
}

func (r *CreateH5FaceIdUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateH5FaceIdUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePreviewSignUrlRequestParams struct {
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// URL过期时间戳
	Deadline *int64 `json:"Deadline,omitempty" name:"Deadline"`

	// 目录ID。当 SignUrlType 为 CATALOG 时必填
	CatalogId *string `json:"CatalogId,omitempty" name:"CatalogId"`

	// 流程ID。当 SignUrlType 为 FLOW 时必填
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 签署链接类型：
	// 1. FLOW - 单流程签署 (默认) 
	// 2. CATALOG - 目录签署
	SignUrlType *string `json:"SignUrlType,omitempty" name:"SignUrlType"`
}

type CreatePreviewSignUrlRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// URL过期时间戳
	Deadline *int64 `json:"Deadline,omitempty" name:"Deadline"`

	// 目录ID。当 SignUrlType 为 CATALOG 时必填
	CatalogId *string `json:"CatalogId,omitempty" name:"CatalogId"`

	// 流程ID。当 SignUrlType 为 FLOW 时必填
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 签署链接类型：
	// 1. FLOW - 单流程签署 (默认) 
	// 2. CATALOG - 目录签署
	SignUrlType *string `json:"SignUrlType,omitempty" name:"SignUrlType"`
}

func (r *CreatePreviewSignUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePreviewSignUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "Deadline")
	delete(f, "CatalogId")
	delete(f, "FlowId")
	delete(f, "SignUrlType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePreviewSignUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePreviewSignUrlResponseParams struct {
	// 合同预览URL
	PreviewSignUrl *string `json:"PreviewSignUrl,omitempty" name:"PreviewSignUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePreviewSignUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreatePreviewSignUrlResponseParams `json:"Response"`
}

func (r *CreatePreviewSignUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePreviewSignUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSealRequestParams struct {
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 印章类型：
	// 1. PERSONAL - 个人私章
	// 2. OFFICIAL - 公章
	// 3. SPECIAL_FINANCIAL - 财务专用章
	// 4. CONTRACT - 合同专用章
	// 5. LEGAL_REPRESENTATIVE - 法定代表人章
	// 6. SPECIAL_NATIONWIDE_INVOICE - 发票专用章
	// 7. OTHER-其他
	SealType *string `json:"SealType,omitempty" name:"SealType"`

	// 印章名称
	SealName *string `json:"SealName,omitempty" name:"SealName"`

	// 请求创建印章的客户端IP
	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`

	// 印章图片，base64编码（与FileId参数二选一，同时传入参数时优先使用Image参数）
	Image *string `json:"Image,omitempty" name:"Image"`

	// 印章文件图片ID（与Image参数二选一，同时传入参数时优先使用Image参数）
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 需要创建印章的用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 是否是默认印章 true：是，false：否
	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`
}

type CreateSealRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 印章类型：
	// 1. PERSONAL - 个人私章
	// 2. OFFICIAL - 公章
	// 3. SPECIAL_FINANCIAL - 财务专用章
	// 4. CONTRACT - 合同专用章
	// 5. LEGAL_REPRESENTATIVE - 法定代表人章
	// 6. SPECIAL_NATIONWIDE_INVOICE - 发票专用章
	// 7. OTHER-其他
	SealType *string `json:"SealType,omitempty" name:"SealType"`

	// 印章名称
	SealName *string `json:"SealName,omitempty" name:"SealName"`

	// 请求创建印章的客户端IP
	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`

	// 印章图片，base64编码（与FileId参数二选一，同时传入参数时优先使用Image参数）
	Image *string `json:"Image,omitempty" name:"Image"`

	// 印章文件图片ID（与Image参数二选一，同时传入参数时优先使用Image参数）
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 需要创建印章的用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 是否是默认印章 true：是，false：否
	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`
}

func (r *CreateSealRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSealRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "SealType")
	delete(f, "SealName")
	delete(f, "SourceIp")
	delete(f, "Image")
	delete(f, "FileId")
	delete(f, "UserId")
	delete(f, "IsDefault")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSealRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSealResponseParams struct {
	// 电子印章Id
	SealId *string `json:"SealId,omitempty" name:"SealId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSealResponse struct {
	*tchttp.BaseResponse
	Response *CreateSealResponseParams `json:"Response"`
}

func (r *CreateSealResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSealResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateServerFlowSignRequestParams struct {
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 流程ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 签署区域信息
	SignComponents []*Component `json:"SignComponents,omitempty" name:"SignComponents"`

	// 客户端IP
	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`
}

type CreateServerFlowSignRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 流程ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 签署区域信息
	SignComponents []*Component `json:"SignComponents,omitempty" name:"SignComponents"`

	// 客户端IP
	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`
}

func (r *CreateServerFlowSignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServerFlowSignRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "FlowId")
	delete(f, "SignComponents")
	delete(f, "SourceIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateServerFlowSignRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateServerFlowSignResponseParams struct {
	// 任务状态：
	// 0：失败
	// 1：成功
	SignStatus *int64 `json:"SignStatus,omitempty" name:"SignStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateServerFlowSignResponse struct {
	*tchttp.BaseResponse
	Response *CreateServerFlowSignResponseParams `json:"Response"`
}

func (r *CreateServerFlowSignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServerFlowSignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSignUrlRequestParams struct {
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 签署人ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 文件签署截止时间戳
	Deadline *int64 `json:"Deadline,omitempty" name:"Deadline"`

	// 目录ID。当 SignUrlType 为 CATALOG 时必填
	CatalogId *string `json:"CatalogId,omitempty" name:"CatalogId"`

	// 流程ID。当 SignUrlType 为 FLOW 时必填
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 签署链接类型：
	// 1. FLOW - 单流程签署 (默认) 
	// 2. CATALOG - 目录签署
	SignUrlType *string `json:"SignUrlType,omitempty" name:"SignUrlType"`

	// 发送流程或目录时生成的签署任务ID
	SignId *string `json:"SignId,omitempty" name:"SignId"`
}

type CreateSignUrlRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 签署人ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 文件签署截止时间戳
	Deadline *int64 `json:"Deadline,omitempty" name:"Deadline"`

	// 目录ID。当 SignUrlType 为 CATALOG 时必填
	CatalogId *string `json:"CatalogId,omitempty" name:"CatalogId"`

	// 流程ID。当 SignUrlType 为 FLOW 时必填
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 签署链接类型：
	// 1. FLOW - 单流程签署 (默认) 
	// 2. CATALOG - 目录签署
	SignUrlType *string `json:"SignUrlType,omitempty" name:"SignUrlType"`

	// 发送流程或目录时生成的签署任务ID
	SignId *string `json:"SignId,omitempty" name:"SignId"`
}

func (r *CreateSignUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSignUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "UserId")
	delete(f, "Deadline")
	delete(f, "CatalogId")
	delete(f, "FlowId")
	delete(f, "SignUrlType")
	delete(f, "SignId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSignUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSignUrlResponseParams struct {
	// 合同签署链接
	SignUrl *string `json:"SignUrl,omitempty" name:"SignUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSignUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreateSignUrlResponseParams `json:"Response"`
}

func (r *CreateSignUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSignUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSubOrganizationAndSealRequestParams struct {
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 机构名称全称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 机构证件号码类型可选值：
	// 1. USCC - 统一社会信用代码
	// 2. BIZREGISTNO - 营业执照注册号
	IdCardType *string `json:"IdCardType,omitempty" name:"IdCardType"`

	// 机构证件号码
	IdCardNumber *string `json:"IdCardNumber,omitempty" name:"IdCardNumber"`

	// 机构类型可选值：
	// 1. ENTERPRISE - 企业
	// 2. INDIVIDUALBIZ - 个体工商户
	// 3. PUBLICINSTITUTION - 政府/事业单位
	// 4. OTHERS - 其他组织
	OrganizationType *string `json:"OrganizationType,omitempty" name:"OrganizationType"`

	// 机构法人/经营者姓名
	LegalName *string `json:"LegalName,omitempty" name:"LegalName"`

	// 机构法人/经营者证件类型可选值：
	// 1. ID_CARD - 居民身份证
	// 2. PASSPORT - 护照
	// 3. MAINLAND_TRAVEL_PERMIT_FOR_HONGKONG_AND_MACAO_RESIDENTS - 港澳居民来往内地通行证
	// 4. MAINLAND_TRAVEL_PERMIT_FOR_TAIWAN_RESIDENTS - 台湾居民来往大陆通行证
	// 5. HOUSEHOLD_REGISTER - 户口本
	// 6. TEMP_ID_CARD - 临时居民身份证
	LegalIdCardType *string `json:"LegalIdCardType,omitempty" name:"LegalIdCardType"`

	// 机构法人/经营者证件号码；
	// OrganizationType 为 ENTERPRISE时，INDIVIDUALBIZ 时必填，其他情况选填
	LegalIdCardNumber *string `json:"LegalIdCardNumber,omitempty" name:"LegalIdCardNumber"`

	// 实名认证的客户端IP/请求生成企业印章的客户端Ip
	VerifyClientIp *string `json:"VerifyClientIp,omitempty" name:"VerifyClientIp"`

	// 机构电子邮箱
	Email *string `json:"Email,omitempty" name:"Email"`

	// 机构证件文件类型可选值：
	// 1. USCCFILE - 统一社会信用代码证书
	// 2. LICENSEFILE - 营业执照
	IdCardFileType *string `json:"IdCardFileType,omitempty" name:"IdCardFileType"`

	// 机构证件照片文件，base64编码，支持jpg、jpeg、png格式
	BizLicenseFile *string `json:"BizLicenseFile,omitempty" name:"BizLicenseFile"`

	// 机构证件照片文件名
	BizLicenseFileName *string `json:"BizLicenseFileName,omitempty" name:"BizLicenseFileName"`

	// 机构法人/经营者/联系人手机号码
	LegalMobile *string `json:"LegalMobile,omitempty" name:"LegalMobile"`

	// 组织联系人姓名
	ContactName *string `json:"ContactName,omitempty" name:"ContactName"`

	// 实名认证的服务器IP
	VerifyServerIp *string `json:"VerifyServerIp,omitempty" name:"VerifyServerIp"`

	// 企业联系地址
	ContactAddress *Address `json:"ContactAddress,omitempty" name:"ContactAddress"`

	// 电子印章名称
	SealName *string `json:"SealName,omitempty" name:"SealName"`

	// 印章类型：默认: CONTRACT
	// 1. OFFICIAL-公章
	// 2. SPECIAL_FINANCIAL-财务专用章
	// 3. CONTRACT-合同专用章
	// 4. LEGAL_REPRESENTATIVE-法定代表人章
	// 5. SPECIAL_NATIONWIDE_INVOICE-发票专用章
	// 6. OTHER-其他
	SealType *string `json:"SealType,omitempty" name:"SealType"`

	// 企业印章横向文字，最多可填8个汉字（可为空，默认为"电子签名专用章"）
	SealHorizontalText *string `json:"SealHorizontalText,omitempty" name:"SealHorizontalText"`

	// 机构在第三方的唯一标识，32位以内标识符
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 是否使用OpenId作为数据主键，如果为true，请确保OpenId在当前应用号唯一
	UseOpenId *bool `json:"UseOpenId,omitempty" name:"UseOpenId"`
}

type CreateSubOrganizationAndSealRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 机构名称全称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 机构证件号码类型可选值：
	// 1. USCC - 统一社会信用代码
	// 2. BIZREGISTNO - 营业执照注册号
	IdCardType *string `json:"IdCardType,omitempty" name:"IdCardType"`

	// 机构证件号码
	IdCardNumber *string `json:"IdCardNumber,omitempty" name:"IdCardNumber"`

	// 机构类型可选值：
	// 1. ENTERPRISE - 企业
	// 2. INDIVIDUALBIZ - 个体工商户
	// 3. PUBLICINSTITUTION - 政府/事业单位
	// 4. OTHERS - 其他组织
	OrganizationType *string `json:"OrganizationType,omitempty" name:"OrganizationType"`

	// 机构法人/经营者姓名
	LegalName *string `json:"LegalName,omitempty" name:"LegalName"`

	// 机构法人/经营者证件类型可选值：
	// 1. ID_CARD - 居民身份证
	// 2. PASSPORT - 护照
	// 3. MAINLAND_TRAVEL_PERMIT_FOR_HONGKONG_AND_MACAO_RESIDENTS - 港澳居民来往内地通行证
	// 4. MAINLAND_TRAVEL_PERMIT_FOR_TAIWAN_RESIDENTS - 台湾居民来往大陆通行证
	// 5. HOUSEHOLD_REGISTER - 户口本
	// 6. TEMP_ID_CARD - 临时居民身份证
	LegalIdCardType *string `json:"LegalIdCardType,omitempty" name:"LegalIdCardType"`

	// 机构法人/经营者证件号码；
	// OrganizationType 为 ENTERPRISE时，INDIVIDUALBIZ 时必填，其他情况选填
	LegalIdCardNumber *string `json:"LegalIdCardNumber,omitempty" name:"LegalIdCardNumber"`

	// 实名认证的客户端IP/请求生成企业印章的客户端Ip
	VerifyClientIp *string `json:"VerifyClientIp,omitempty" name:"VerifyClientIp"`

	// 机构电子邮箱
	Email *string `json:"Email,omitempty" name:"Email"`

	// 机构证件文件类型可选值：
	// 1. USCCFILE - 统一社会信用代码证书
	// 2. LICENSEFILE - 营业执照
	IdCardFileType *string `json:"IdCardFileType,omitempty" name:"IdCardFileType"`

	// 机构证件照片文件，base64编码，支持jpg、jpeg、png格式
	BizLicenseFile *string `json:"BizLicenseFile,omitempty" name:"BizLicenseFile"`

	// 机构证件照片文件名
	BizLicenseFileName *string `json:"BizLicenseFileName,omitempty" name:"BizLicenseFileName"`

	// 机构法人/经营者/联系人手机号码
	LegalMobile *string `json:"LegalMobile,omitempty" name:"LegalMobile"`

	// 组织联系人姓名
	ContactName *string `json:"ContactName,omitempty" name:"ContactName"`

	// 实名认证的服务器IP
	VerifyServerIp *string `json:"VerifyServerIp,omitempty" name:"VerifyServerIp"`

	// 企业联系地址
	ContactAddress *Address `json:"ContactAddress,omitempty" name:"ContactAddress"`

	// 电子印章名称
	SealName *string `json:"SealName,omitempty" name:"SealName"`

	// 印章类型：默认: CONTRACT
	// 1. OFFICIAL-公章
	// 2. SPECIAL_FINANCIAL-财务专用章
	// 3. CONTRACT-合同专用章
	// 4. LEGAL_REPRESENTATIVE-法定代表人章
	// 5. SPECIAL_NATIONWIDE_INVOICE-发票专用章
	// 6. OTHER-其他
	SealType *string `json:"SealType,omitempty" name:"SealType"`

	// 企业印章横向文字，最多可填8个汉字（可为空，默认为"电子签名专用章"）
	SealHorizontalText *string `json:"SealHorizontalText,omitempty" name:"SealHorizontalText"`

	// 机构在第三方的唯一标识，32位以内标识符
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 是否使用OpenId作为数据主键，如果为true，请确保OpenId在当前应用号唯一
	UseOpenId *bool `json:"UseOpenId,omitempty" name:"UseOpenId"`
}

func (r *CreateSubOrganizationAndSealRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSubOrganizationAndSealRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "Name")
	delete(f, "IdCardType")
	delete(f, "IdCardNumber")
	delete(f, "OrganizationType")
	delete(f, "LegalName")
	delete(f, "LegalIdCardType")
	delete(f, "LegalIdCardNumber")
	delete(f, "VerifyClientIp")
	delete(f, "Email")
	delete(f, "IdCardFileType")
	delete(f, "BizLicenseFile")
	delete(f, "BizLicenseFileName")
	delete(f, "LegalMobile")
	delete(f, "ContactName")
	delete(f, "VerifyServerIp")
	delete(f, "ContactAddress")
	delete(f, "SealName")
	delete(f, "SealType")
	delete(f, "SealHorizontalText")
	delete(f, "OpenId")
	delete(f, "UseOpenId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSubOrganizationAndSealRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSubOrganizationAndSealResponseParams struct {
	// 子机构在电子文件签署平台唯一标识
	SubOrganizationId *string `json:"SubOrganizationId,omitempty" name:"SubOrganizationId"`

	// 电子印章ID
	SealId *string `json:"SealId,omitempty" name:"SealId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSubOrganizationAndSealResponse struct {
	*tchttp.BaseResponse
	Response *CreateSubOrganizationAndSealResponseParams `json:"Response"`
}

func (r *CreateSubOrganizationAndSealResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSubOrganizationAndSealResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSubOrganizationRequestParams struct {
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 机构证件号码类型可选值：
	// 1. USCC - 统一社会信用代码
	// 2. BIZREGISTNO - 营业执照注册号
	IdCardType *string `json:"IdCardType,omitempty" name:"IdCardType"`

	// 机构证件号码
	IdCardNumber *string `json:"IdCardNumber,omitempty" name:"IdCardNumber"`

	// 机构类型可选值：
	// 1. ENTERPRISE - 企业
	// 2. INDIVIDUALBIZ - 个体工商户
	// 3. PUBLICINSTITUTION - 政府/事业单位
	// 4. OTHERS - 其他组织
	OrganizationType *string `json:"OrganizationType,omitempty" name:"OrganizationType"`

	// 机构法人/经营者姓名
	LegalName *string `json:"LegalName,omitempty" name:"LegalName"`

	// 机构法人/经营者证件类型可选值：
	// 1. ID_CARD - 居民身份证
	// 2. PASSPORT - 护照
	// 3. MAINLAND_TRAVEL_PERMIT_FOR_HONGKONG_AND_MACAO_RESIDENTS - 港澳居民来往内地通行证
	// 4. MAINLAND_TRAVEL_PERMIT_FOR_TAIWAN_RESIDENTS - 台湾居民来往大陆通行证
	// 5. HOUSEHOLD_REGISTER - 户口本
	// 6. TEMP_ID_CARD - 临时居民身份证
	LegalIdCardType *string `json:"LegalIdCardType,omitempty" name:"LegalIdCardType"`

	// 机构法人/经营者证件号码；
	// OrganizationType 为 ENTERPRISE时，INDIVIDUALBIZ 时必填，其他情况选填
	LegalIdCardNumber *string `json:"LegalIdCardNumber,omitempty" name:"LegalIdCardNumber"`

	// 机构名称全称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 机构在第三方的唯一标识，32位以内标识符
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 是否使用OpenId作为数据主键，如果为true，请确保OpenId在当前应用号唯一
	UseOpenId *bool `json:"UseOpenId,omitempty" name:"UseOpenId"`

	// 机构证件文件类型可选值：
	// 1. USCCFILE - 统一社会信用代码证书
	// 2. LICENSEFILE - 营业执照
	IdCardFileType *string `json:"IdCardFileType,omitempty" name:"IdCardFileType"`

	// 机构证件照片文件，base64编码，支持jpg、jpeg、png格式
	BizLicenseFile *string `json:"BizLicenseFile,omitempty" name:"BizLicenseFile"`

	// 机构证件照片文件名
	BizLicenseFileName *string `json:"BizLicenseFileName,omitempty" name:"BizLicenseFileName"`

	// 机构法人/经营者/联系人手机号码
	LegalMobile *string `json:"LegalMobile,omitempty" name:"LegalMobile"`

	// 组织联系人姓名
	ContactName *string `json:"ContactName,omitempty" name:"ContactName"`

	// 实名认证的客户端IP
	VerifyClientIp *string `json:"VerifyClientIp,omitempty" name:"VerifyClientIp"`

	// 实名认证的服务器IP
	VerifyServerIp *string `json:"VerifyServerIp,omitempty" name:"VerifyServerIp"`

	// 企业联系地址
	ContactAddress *Address `json:"ContactAddress,omitempty" name:"ContactAddress"`

	// 机构电子邮箱
	Email *string `json:"Email,omitempty" name:"Email"`
}

type CreateSubOrganizationRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 机构证件号码类型可选值：
	// 1. USCC - 统一社会信用代码
	// 2. BIZREGISTNO - 营业执照注册号
	IdCardType *string `json:"IdCardType,omitempty" name:"IdCardType"`

	// 机构证件号码
	IdCardNumber *string `json:"IdCardNumber,omitempty" name:"IdCardNumber"`

	// 机构类型可选值：
	// 1. ENTERPRISE - 企业
	// 2. INDIVIDUALBIZ - 个体工商户
	// 3. PUBLICINSTITUTION - 政府/事业单位
	// 4. OTHERS - 其他组织
	OrganizationType *string `json:"OrganizationType,omitempty" name:"OrganizationType"`

	// 机构法人/经营者姓名
	LegalName *string `json:"LegalName,omitempty" name:"LegalName"`

	// 机构法人/经营者证件类型可选值：
	// 1. ID_CARD - 居民身份证
	// 2. PASSPORT - 护照
	// 3. MAINLAND_TRAVEL_PERMIT_FOR_HONGKONG_AND_MACAO_RESIDENTS - 港澳居民来往内地通行证
	// 4. MAINLAND_TRAVEL_PERMIT_FOR_TAIWAN_RESIDENTS - 台湾居民来往大陆通行证
	// 5. HOUSEHOLD_REGISTER - 户口本
	// 6. TEMP_ID_CARD - 临时居民身份证
	LegalIdCardType *string `json:"LegalIdCardType,omitempty" name:"LegalIdCardType"`

	// 机构法人/经营者证件号码；
	// OrganizationType 为 ENTERPRISE时，INDIVIDUALBIZ 时必填，其他情况选填
	LegalIdCardNumber *string `json:"LegalIdCardNumber,omitempty" name:"LegalIdCardNumber"`

	// 机构名称全称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 机构在第三方的唯一标识，32位以内标识符
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 是否使用OpenId作为数据主键，如果为true，请确保OpenId在当前应用号唯一
	UseOpenId *bool `json:"UseOpenId,omitempty" name:"UseOpenId"`

	// 机构证件文件类型可选值：
	// 1. USCCFILE - 统一社会信用代码证书
	// 2. LICENSEFILE - 营业执照
	IdCardFileType *string `json:"IdCardFileType,omitempty" name:"IdCardFileType"`

	// 机构证件照片文件，base64编码，支持jpg、jpeg、png格式
	BizLicenseFile *string `json:"BizLicenseFile,omitempty" name:"BizLicenseFile"`

	// 机构证件照片文件名
	BizLicenseFileName *string `json:"BizLicenseFileName,omitempty" name:"BizLicenseFileName"`

	// 机构法人/经营者/联系人手机号码
	LegalMobile *string `json:"LegalMobile,omitempty" name:"LegalMobile"`

	// 组织联系人姓名
	ContactName *string `json:"ContactName,omitempty" name:"ContactName"`

	// 实名认证的客户端IP
	VerifyClientIp *string `json:"VerifyClientIp,omitempty" name:"VerifyClientIp"`

	// 实名认证的服务器IP
	VerifyServerIp *string `json:"VerifyServerIp,omitempty" name:"VerifyServerIp"`

	// 企业联系地址
	ContactAddress *Address `json:"ContactAddress,omitempty" name:"ContactAddress"`

	// 机构电子邮箱
	Email *string `json:"Email,omitempty" name:"Email"`
}

func (r *CreateSubOrganizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSubOrganizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "IdCardType")
	delete(f, "IdCardNumber")
	delete(f, "OrganizationType")
	delete(f, "LegalName")
	delete(f, "LegalIdCardType")
	delete(f, "LegalIdCardNumber")
	delete(f, "Name")
	delete(f, "OpenId")
	delete(f, "UseOpenId")
	delete(f, "IdCardFileType")
	delete(f, "BizLicenseFile")
	delete(f, "BizLicenseFileName")
	delete(f, "LegalMobile")
	delete(f, "ContactName")
	delete(f, "VerifyClientIp")
	delete(f, "VerifyServerIp")
	delete(f, "ContactAddress")
	delete(f, "Email")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSubOrganizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSubOrganizationResponseParams struct {
	// 子机构ID
	SubOrganizationId *string `json:"SubOrganizationId,omitempty" name:"SubOrganizationId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSubOrganizationResponse struct {
	*tchttp.BaseResponse
	Response *CreateSubOrganizationResponseParams `json:"Response"`
}

func (r *CreateSubOrganizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSubOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserAndSealRequestParams struct {
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 第三方平台唯一标识，要求应用内OpenId唯一
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 用户姓名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 用户证件类型：
	// 1. ID_CARD - 居民身份证
	// 5. HOUSEHOLD_REGISTER - 户口本
	// 6. TEMP_ID_CARD - 临时居民身份证
	IdCardType *string `json:"IdCardType,omitempty" name:"IdCardType"`

	// 用户证件号
	IdCardNumber *string `json:"IdCardNumber,omitempty" name:"IdCardNumber"`

	// 请求生成个人印章的客户端IP
	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`

	// 用户手机号码，不要求唯一
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 用户邮箱，不要求唯一
	Email *string `json:"Email,omitempty" name:"Email"`

	// 默认印章名称
	SealName *string `json:"SealName,omitempty" name:"SealName"`

	// 是否以OpenId作为UserId (为true时将直接以OpenId生成腾讯电子签平台的UserId)
	UseOpenId *bool `json:"UseOpenId,omitempty" name:"UseOpenId"`
}

type CreateUserAndSealRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 第三方平台唯一标识，要求应用内OpenId唯一
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 用户姓名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 用户证件类型：
	// 1. ID_CARD - 居民身份证
	// 5. HOUSEHOLD_REGISTER - 户口本
	// 6. TEMP_ID_CARD - 临时居民身份证
	IdCardType *string `json:"IdCardType,omitempty" name:"IdCardType"`

	// 用户证件号
	IdCardNumber *string `json:"IdCardNumber,omitempty" name:"IdCardNumber"`

	// 请求生成个人印章的客户端IP
	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`

	// 用户手机号码，不要求唯一
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 用户邮箱，不要求唯一
	Email *string `json:"Email,omitempty" name:"Email"`

	// 默认印章名称
	SealName *string `json:"SealName,omitempty" name:"SealName"`

	// 是否以OpenId作为UserId (为true时将直接以OpenId生成腾讯电子签平台的UserId)
	UseOpenId *bool `json:"UseOpenId,omitempty" name:"UseOpenId"`
}

func (r *CreateUserAndSealRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserAndSealRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "OpenId")
	delete(f, "Name")
	delete(f, "IdCardType")
	delete(f, "IdCardNumber")
	delete(f, "SourceIp")
	delete(f, "Mobile")
	delete(f, "Email")
	delete(f, "SealName")
	delete(f, "UseOpenId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserAndSealRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserAndSealResponseParams struct {
	// 用户唯一标识，按应用号隔离
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 默认印章ID
	SealId *string `json:"SealId,omitempty" name:"SealId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateUserAndSealResponse struct {
	*tchttp.BaseResponse
	Response *CreateUserAndSealResponseParams `json:"Response"`
}

func (r *CreateUserAndSealResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserAndSealResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserRequestParams struct {
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 第三方平台唯一标识；要求应用内OpenId唯一; len<=32
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 用户姓名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 用户证件类型：
	// 1. ID_CARD - 居民身份证
	// 2. PASSPORT - 护照
	// 3. MAINLAND_TRAVEL_PERMIT_FOR_HONGKONG_AND_MACAO_RESIDENTS - 港澳居民来往内地通行证
	// 4. MAINLAND_TRAVEL_PERMIT_FOR_TAIWAN_RESIDENTS - 台湾居民来往大陆通行证
	// 5. HOUSEHOLD_REGISTER - 户口本
	// 6. TEMP_ID_CARD - 临时居民身份证
	IdCardType *string `json:"IdCardType,omitempty" name:"IdCardType"`

	// 用户证件号
	IdCardNumber *string `json:"IdCardNumber,omitempty" name:"IdCardNumber"`

	// 是否以OpenId作为UserId (为true时将直接以OpenId生成腾讯电子签平台的UserId)
	UseOpenId *bool `json:"UseOpenId,omitempty" name:"UseOpenId"`

	// 用户邮箱，不要求唯一
	Email *string `json:"Email,omitempty" name:"Email"`

	// 用户手机号码，不要求唯一
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`
}

type CreateUserRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 第三方平台唯一标识；要求应用内OpenId唯一; len<=32
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 用户姓名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 用户证件类型：
	// 1. ID_CARD - 居民身份证
	// 2. PASSPORT - 护照
	// 3. MAINLAND_TRAVEL_PERMIT_FOR_HONGKONG_AND_MACAO_RESIDENTS - 港澳居民来往内地通行证
	// 4. MAINLAND_TRAVEL_PERMIT_FOR_TAIWAN_RESIDENTS - 台湾居民来往大陆通行证
	// 5. HOUSEHOLD_REGISTER - 户口本
	// 6. TEMP_ID_CARD - 临时居民身份证
	IdCardType *string `json:"IdCardType,omitempty" name:"IdCardType"`

	// 用户证件号
	IdCardNumber *string `json:"IdCardNumber,omitempty" name:"IdCardNumber"`

	// 是否以OpenId作为UserId (为true时将直接以OpenId生成腾讯电子签平台的UserId)
	UseOpenId *bool `json:"UseOpenId,omitempty" name:"UseOpenId"`

	// 用户邮箱，不要求唯一
	Email *string `json:"Email,omitempty" name:"Email"`

	// 用户手机号码，不要求唯一
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`
}

func (r *CreateUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "OpenId")
	delete(f, "Name")
	delete(f, "IdCardType")
	delete(f, "IdCardNumber")
	delete(f, "UseOpenId")
	delete(f, "Email")
	delete(f, "Mobile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserResponseParams struct {
	// 用户ID，按应用号隔离
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateUserResponse struct {
	*tchttp.BaseResponse
	Response *CreateUserResponseParams `json:"Response"`
}

func (r *CreateUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomFileIdMap struct {
	// 用户自定义ID
	CustomId *string `json:"CustomId,omitempty" name:"CustomId"`

	// 文件id
	FileId *string `json:"FileId,omitempty" name:"FileId"`
}

type CustomFlowIdMap struct {
	// 自定义id
	CustomId *string `json:"CustomId,omitempty" name:"CustomId"`

	// 流程id
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
}

// Predefined struct for user
type DeleteSealRequestParams struct {
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 印章ID
	SealId *string `json:"SealId,omitempty" name:"SealId"`

	// 请求删除印章的客户端IP
	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`

	// 用户唯一标识，默认为空时删除企业印章，如非空则删除个人印章
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

type DeleteSealRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 印章ID
	SealId *string `json:"SealId,omitempty" name:"SealId"`

	// 请求删除印章的客户端IP
	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`

	// 用户唯一标识，默认为空时删除企业印章，如非空则删除个人印章
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

func (r *DeleteSealRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSealRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "SealId")
	delete(f, "SourceIp")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSealRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSealResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteSealResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSealResponseParams `json:"Response"`
}

func (r *DeleteSealResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSealResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCatalogApproversRequestParams struct {
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 目录ID
	CatalogId *string `json:"CatalogId,omitempty" name:"CatalogId"`

	// 查询指定用户是否为参与者,为空表示查询所有参与者
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

type DescribeCatalogApproversRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 目录ID
	CatalogId *string `json:"CatalogId,omitempty" name:"CatalogId"`

	// 查询指定用户是否为参与者,为空表示查询所有参与者
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

func (r *DescribeCatalogApproversRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCatalogApproversRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "CatalogId")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCatalogApproversRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCatalogApproversResponseParams struct {
	// 参与者列表
	Approvers []*CatalogApprovers `json:"Approvers,omitempty" name:"Approvers"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCatalogApproversResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCatalogApproversResponseParams `json:"Response"`
}

func (r *DescribeCatalogApproversResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCatalogApproversResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCatalogSignComponentsRequestParams struct {
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 目录ID
	CatalogId *string `json:"CatalogId,omitempty" name:"CatalogId"`
}

type DescribeCatalogSignComponentsRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 目录ID
	CatalogId *string `json:"CatalogId,omitempty" name:"CatalogId"`
}

func (r *DescribeCatalogSignComponentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCatalogSignComponentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "CatalogId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCatalogSignComponentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCatalogSignComponentsResponseParams struct {
	// 签署区列表
	SignComponents []*CatalogComponents `json:"SignComponents,omitempty" name:"SignComponents"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCatalogSignComponentsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCatalogSignComponentsResponseParams `json:"Response"`
}

func (r *DescribeCatalogSignComponentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCatalogSignComponentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomFlowIdsByFlowIdRequestParams struct {
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 流程 id 列表，最多同时查询 10 个流程 id
	FlowIds []*string `json:"FlowIds,omitempty" name:"FlowIds"`
}

type DescribeCustomFlowIdsByFlowIdRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 流程 id 列表，最多同时查询 10 个流程 id
	FlowIds []*string `json:"FlowIds,omitempty" name:"FlowIds"`
}

func (r *DescribeCustomFlowIdsByFlowIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomFlowIdsByFlowIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "FlowIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomFlowIdsByFlowIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomFlowIdsByFlowIdResponseParams struct {
	// 自定义流程 id 映射列表
	CustomIdList []*CustomFlowIdMap `json:"CustomIdList,omitempty" name:"CustomIdList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCustomFlowIdsByFlowIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomFlowIdsByFlowIdResponseParams `json:"Response"`
}

func (r *DescribeCustomFlowIdsByFlowIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomFlowIdsByFlowIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomFlowIdsRequestParams struct {
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 自定义 id 列表，最多同时查询 10 个自定义 id
	CustomIds []*string `json:"CustomIds,omitempty" name:"CustomIds"`
}

type DescribeCustomFlowIdsRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 自定义 id 列表，最多同时查询 10 个自定义 id
	CustomIds []*string `json:"CustomIds,omitempty" name:"CustomIds"`
}

func (r *DescribeCustomFlowIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomFlowIdsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "CustomIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomFlowIdsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomFlowIdsResponseParams struct {
	// 自定义流程 id 映射列表
	CustomIdList []*CustomFlowIdMap `json:"CustomIdList,omitempty" name:"CustomIdList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCustomFlowIdsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomFlowIdsResponseParams `json:"Response"`
}

func (r *DescribeCustomFlowIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomFlowIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFaceIdPhotosRequestParams struct {
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 慧眼业务ID
	WbAppId *string `json:"WbAppId,omitempty" name:"WbAppId"`

	// 订单号(orderNo); 限制在3个或以内
	OrderNumbers []*string `json:"OrderNumbers,omitempty" name:"OrderNumbers"`
}

type DescribeFaceIdPhotosRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 慧眼业务ID
	WbAppId *string `json:"WbAppId,omitempty" name:"WbAppId"`

	// 订单号(orderNo); 限制在3个或以内
	OrderNumbers []*string `json:"OrderNumbers,omitempty" name:"OrderNumbers"`
}

func (r *DescribeFaceIdPhotosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFaceIdPhotosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "WbAppId")
	delete(f, "OrderNumbers")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFaceIdPhotosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFaceIdPhotosResponseParams struct {
	// 照片信息列表
	Photos []*FaceIdPhoto `json:"Photos,omitempty" name:"Photos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFaceIdPhotosResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFaceIdPhotosResponseParams `json:"Response"`
}

func (r *DescribeFaceIdPhotosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFaceIdPhotosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFaceIdResultsRequestParams struct {
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 慧眼业务ID
	WbAppId *string `json:"WbAppId,omitempty" name:"WbAppId"`

	// 订单号(orderNo); 限制在3个或以内
	OrderNumbers []*string `json:"OrderNumbers,omitempty" name:"OrderNumbers"`

	// 1:视频+照片,2:照片,3:视频,0（或其他数字）:无; 可选
	FileType *int64 `json:"FileType,omitempty" name:"FileType"`
}

type DescribeFaceIdResultsRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 慧眼业务ID
	WbAppId *string `json:"WbAppId,omitempty" name:"WbAppId"`

	// 订单号(orderNo); 限制在3个或以内
	OrderNumbers []*string `json:"OrderNumbers,omitempty" name:"OrderNumbers"`

	// 1:视频+照片,2:照片,3:视频,0（或其他数字）:无; 可选
	FileType *int64 `json:"FileType,omitempty" name:"FileType"`
}

func (r *DescribeFaceIdResultsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFaceIdResultsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "WbAppId")
	delete(f, "OrderNumbers")
	delete(f, "FileType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFaceIdResultsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFaceIdResultsResponseParams struct {
	// 核身结果列表
	Results []*FaceIdResult `json:"Results,omitempty" name:"Results"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFaceIdResultsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFaceIdResultsResponseParams `json:"Response"`
}

func (r *DescribeFaceIdResultsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFaceIdResultsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileIdsByCustomIdsRequestParams struct {
	// 调用方信息, OrganizationId必填
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 用户自定义ID
	CustomIds []*string `json:"CustomIds,omitempty" name:"CustomIds"`
}

type DescribeFileIdsByCustomIdsRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息, OrganizationId必填
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 用户自定义ID
	CustomIds []*string `json:"CustomIds,omitempty" name:"CustomIds"`
}

func (r *DescribeFileIdsByCustomIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileIdsByCustomIdsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "CustomIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFileIdsByCustomIdsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileIdsByCustomIdsResponseParams struct {
	// <自定义Id,文件id>数组
	CustomIdList []*CustomFileIdMap `json:"CustomIdList,omitempty" name:"CustomIdList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFileIdsByCustomIdsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFileIdsByCustomIdsResponseParams `json:"Response"`
}

func (r *DescribeFileIdsByCustomIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileIdsByCustomIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileUrlsRequestParams struct {
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 业务编号数组，如模板编号、文档编号、印章编号、流程编号、目录编号
	BusinessIds []*string `json:"BusinessIds,omitempty" name:"BusinessIds"`

	// 业务类型：
	// 1. TEMPLATE - 模板
	// 2. SEAL - 印章
	// 3. FLOW - 流程
	// 4.CATALOG - 目录
	BusinessType *string `json:"BusinessType,omitempty" name:"BusinessType"`

	// 下载后的文件命名，只有FileType为“ZIP”时生效
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 单个业务ID多个资源情况下，指定资源起始偏移量
	ResourceOffset *int64 `json:"ResourceOffset,omitempty" name:"ResourceOffset"`

	// 单个业务ID多个资源情况下，指定资源数量
	ResourceLimit *int64 `json:"ResourceLimit,omitempty" name:"ResourceLimit"`

	// 文件类型，支持"JPG", "PDF","ZIP"等，默认为上传的文件类型
	FileType *string `json:"FileType,omitempty" name:"FileType"`
}

type DescribeFileUrlsRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 业务编号数组，如模板编号、文档编号、印章编号、流程编号、目录编号
	BusinessIds []*string `json:"BusinessIds,omitempty" name:"BusinessIds"`

	// 业务类型：
	// 1. TEMPLATE - 模板
	// 2. SEAL - 印章
	// 3. FLOW - 流程
	// 4.CATALOG - 目录
	BusinessType *string `json:"BusinessType,omitempty" name:"BusinessType"`

	// 下载后的文件命名，只有FileType为“ZIP”时生效
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 单个业务ID多个资源情况下，指定资源起始偏移量
	ResourceOffset *int64 `json:"ResourceOffset,omitempty" name:"ResourceOffset"`

	// 单个业务ID多个资源情况下，指定资源数量
	ResourceLimit *int64 `json:"ResourceLimit,omitempty" name:"ResourceLimit"`

	// 文件类型，支持"JPG", "PDF","ZIP"等，默认为上传的文件类型
	FileType *string `json:"FileType,omitempty" name:"FileType"`
}

func (r *DescribeFileUrlsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileUrlsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "BusinessIds")
	delete(f, "BusinessType")
	delete(f, "FileName")
	delete(f, "ResourceOffset")
	delete(f, "ResourceLimit")
	delete(f, "FileType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFileUrlsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileUrlsResponseParams struct {
	// 文件下载URL数组
	FileUrls []*FileUrl `json:"FileUrls,omitempty" name:"FileUrls"`

	// URL数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFileUrlsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFileUrlsResponseParams `json:"Response"`
}

func (r *DescribeFileUrlsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileUrlsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowApproversRequestParams struct {
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 需要查询的流程ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 需要查询的用户ID，为空则默认查询所有用户信息
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 需要查询的签署ID，为空则不按签署ID过滤
	SignId *string `json:"SignId,omitempty" name:"SignId"`
}

type DescribeFlowApproversRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 需要查询的流程ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 需要查询的用户ID，为空则默认查询所有用户信息
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 需要查询的签署ID，为空则不按签署ID过滤
	SignId *string `json:"SignId,omitempty" name:"SignId"`
}

func (r *DescribeFlowApproversRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowApproversRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "FlowId")
	delete(f, "UserId")
	delete(f, "SignId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFlowApproversRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowApproversResponseParams struct {
	// 流程编号
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 流程参与者信息
	Approvers []*FlowApproverInfo `json:"Approvers,omitempty" name:"Approvers"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFlowApproversResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFlowApproversResponseParams `json:"Response"`
}

func (r *DescribeFlowApproversResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowApproversResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowFilesRequestParams struct {
	// 调用方信息; 必选
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 需要查询的流程ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
}

type DescribeFlowFilesRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息; 必选
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 需要查询的流程ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
}

func (r *DescribeFlowFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "FlowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFlowFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowFilesResponseParams struct {
	// 流程编号
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 流程文件列表
	FlowFileInfos []*FlowFileInfo `json:"FlowFileInfos,omitempty" name:"FlowFileInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFlowFilesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFlowFilesResponseParams `json:"Response"`
}

func (r *DescribeFlowFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowRequestParams struct {
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 需要查询的流程ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
}

type DescribeFlowRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 需要查询的流程ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
}

func (r *DescribeFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "FlowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowResponseParams struct {
	// 流程创建者信息
	Creator *Caller `json:"Creator,omitempty" name:"Creator"`

	// 流程编号
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 流程名称
	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`

	// 流程描述
	FlowDescription *string `json:"FlowDescription,omitempty" name:"FlowDescription"`

	// 流程的类型: ”劳务合同“,”租赁合同“,”销售合同“,”其他“
	FlowType *string `json:"FlowType,omitempty" name:"FlowType"`

	// 流程状态：
	// 0-创建；
	// 1-签署中；
	// 2-拒签；
	// 3-撤回；
	// 4-签完存档完成；
	// 5-已过期；
	// 6-已销毁
	// 7-签署完成未归档
	FlowStatus *int64 `json:"FlowStatus,omitempty" name:"FlowStatus"`

	// 流程创建时间
	CreatedOn *int64 `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// 流程完成时间
	UpdatedOn *int64 `json:"UpdatedOn,omitempty" name:"UpdatedOn"`

	// 流程截止日期
	Deadline *int64 `json:"Deadline,omitempty" name:"Deadline"`

	// 回调地址
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 流程中止原因
	FlowMessage *string `json:"FlowMessage,omitempty" name:"FlowMessage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFlowResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFlowResponseParams `json:"Response"`
}

func (r *DescribeFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSealsRequestParams struct {
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 印章ID列表
	SealIds []*string `json:"SealIds,omitempty" name:"SealIds"`

	// 用户唯一标识
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

type DescribeSealsRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 印章ID列表
	SealIds []*string `json:"SealIds,omitempty" name:"SealIds"`

	// 用户唯一标识
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

func (r *DescribeSealsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSealsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "SealIds")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSealsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSealsResponseParams struct {
	// 印章信息
	Seals []*Seal `json:"Seals,omitempty" name:"Seals"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSealsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSealsResponseParams `json:"Response"`
}

func (r *DescribeSealsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSealsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubOrganizationsRequestParams struct {
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 子机构ID数组
	SubOrganizationIds []*string `json:"SubOrganizationIds,omitempty" name:"SubOrganizationIds"`
}

type DescribeSubOrganizationsRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 子机构ID数组
	SubOrganizationIds []*string `json:"SubOrganizationIds,omitempty" name:"SubOrganizationIds"`
}

func (r *DescribeSubOrganizationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubOrganizationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "SubOrganizationIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSubOrganizationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubOrganizationsResponseParams struct {
	// 子机构信息列表
	SubOrganizationInfos []*SubOrganizationDetail `json:"SubOrganizationInfos,omitempty" name:"SubOrganizationInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSubOrganizationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSubOrganizationsResponseParams `json:"Response"`
}

func (r *DescribeSubOrganizationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubOrganizationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsersRequestParams struct {
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// UserId列表，最多支持100个UserId
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`
}

type DescribeUsersRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// UserId列表，最多支持100个UserId
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`
}

func (r *DescribeUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "UserIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsersResponseParams struct {
	// 用户信息查询结果
	Users []*UserDescribe `json:"Users,omitempty" name:"Users"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUsersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUsersResponseParams `json:"Response"`
}

func (r *DescribeUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyFlowFileRequestParams struct {
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 流程ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
}

type DestroyFlowFileRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 流程ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
}

func (r *DestroyFlowFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyFlowFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "FlowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyFlowFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyFlowFileResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DestroyFlowFileResponse struct {
	*tchttp.BaseResponse
	Response *DestroyFlowFileResponseParams `json:"Response"`
}

func (r *DestroyFlowFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyFlowFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FaceIdPhoto struct {
	// 核身结果：
	// 0 - 通过；
	// 1 - 未通过
	Result *int64 `json:"Result,omitempty" name:"Result"`

	// 核身失败描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 照片数据 (base64编码, 一般为JPG或PNG)
	Photo *string `json:"Photo,omitempty" name:"Photo"`

	// 订单号 (orderNo)
	OrderNumber *string `json:"OrderNumber,omitempty" name:"OrderNumber"`
}

type FaceIdResult struct {
	// 核身结果：
	// 0 - 通过；
	// 1 - 未通过
	Result *int64 `json:"Result,omitempty" name:"Result"`

	// 核身失败描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 订单号 (orderNo)
	OrderNumber *string `json:"OrderNumber,omitempty" name:"OrderNumber"`

	// 姓名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 身份证件类型： 
	// ID_CARD - 居民身份证
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdCardType *string `json:"IdCardType,omitempty" name:"IdCardType"`

	// 身份证件号码
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdCardNumber *string `json:"IdCardNumber,omitempty" name:"IdCardNumber"`

	// 活体检测得分 (百分制)
	// 注意：此字段可能返回 null，表示取不到有效值。
	LiveRate *int64 `json:"LiveRate,omitempty" name:"LiveRate"`

	// 人脸检测得分 (百分制)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Similarity *float64 `json:"Similarity,omitempty" name:"Similarity"`

	// 刷脸时间 (UNIX时间戳)
	// 注意：此字段可能返回 null，表示取不到有效值。
	OccurredTime *int64 `json:"OccurredTime,omitempty" name:"OccurredTime"`

	// 照片数据 (base64编码, 一般为JPG或PNG)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Photo *string `json:"Photo,omitempty" name:"Photo"`

	// 视频数据 (base64编码, 一般为MP4)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Video *string `json:"Video,omitempty" name:"Video"`
}

type FileUrl struct {
	// 下载文件的URL
	Url *string `json:"Url,omitempty" name:"Url"`

	// 下载文件的附加信息
	Option *string `json:"Option,omitempty" name:"Option"`

	// 下载文件所属的资源序号
	Index *int64 `json:"Index,omitempty" name:"Index"`

	// 目录业务下，文件对应的流程
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
}

type FlowApproverInfo struct {
	// 用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 认证方式：
	// WEIXINAPP - 微信小程序；
	// FACEID - 慧眼 (默认)；
	// VERIFYCODE - 验证码；
	// THIRD - 第三方 (暂不支持)
	VerifyChannel []*string `json:"VerifyChannel,omitempty" name:"VerifyChannel"`

	// 签署状态：
	// 0 - 待签署；
	// 1- 已签署；
	// 2 - 拒绝；
	// 3 - 过期未处理；
	// 4 - 流程已撤回,
	// 12-审核中,
	// 13-审核驳回
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproveStatus *int64 `json:"ApproveStatus,omitempty" name:"ApproveStatus"`

	// 拒签/签署/审核驳回原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproveMessage *string `json:"ApproveMessage,omitempty" name:"ApproveMessage"`

	// 签约时间的时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproveTime *int64 `json:"ApproveTime,omitempty" name:"ApproveTime"`

	// 签署企业ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubOrganizationId *string `json:"SubOrganizationId,omitempty" name:"SubOrganizationId"`

	// 签署完成后跳转的URL
	// 注意：此字段可能返回 null，表示取不到有效值。
	JumpUrl *string `json:"JumpUrl,omitempty" name:"JumpUrl"`

	// 用户签署区ID到印章ID的映射集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentSeals []*ComponentSeal `json:"ComponentSeals,omitempty" name:"ComponentSeals"`

	// 签署前置条件：是否强制用户全文阅读，即阅读到待签署文档的最后一页。默认FALSE
	IsFullText *bool `json:"IsFullText,omitempty" name:"IsFullText"`

	// 签署前置条件：强制阅读时长，页面停留时长不足则不允许签署。默认不限制
	PreReadTime *int64 `json:"PreReadTime,omitempty" name:"PreReadTime"`

	// 签署人手机号，脱敏显示
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 签署链接截止时间，默认签署流程发起后7天失效
	Deadline *int64 `json:"Deadline,omitempty" name:"Deadline"`

	// 是否为最后一个签署人, 若为最后一人，则其签署完成后自动归档
	IsLastApprover *bool `json:"IsLastApprover,omitempty" name:"IsLastApprover"`

	// 短信模板
	// 注意：此字段可能返回 null，表示取不到有效值。
	SmsTemplate *SmsTemplate `json:"SmsTemplate,omitempty" name:"SmsTemplate"`

	// 身份证号，脱敏显示
	IdCardNumber *string `json:"IdCardNumber,omitempty" name:"IdCardNumber"`

	// 用户姓名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否支持线下核身
	CanOffLine *bool `json:"CanOffLine,omitempty" name:"CanOffLine"`

	// 证件号码类型：ID_CARD - 身份证，PASSPORT - 护照，MAINLAND_TRAVEL_PERMIT_FOR_HONGKONG_AND_MACAO_RESIDENTS - 港澳居民来往内地通行证; 暂不支持用于电子签自有平台实名认证，MAINLAND_TRAVEL_PERMIT_FOR_TAIWAN_RESIDENTS - 台湾居民来往大陆通行证; 暂不支持用于电子签自有平台实名认证，HOUSEHOLD_REGISTER - 户口本; 暂不支持用于电子签自有平台实名认证，TEMP_ID_CARD - 临时居民身份证; 暂不支持用于电子签自有平台实名认证
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdCardType *string `json:"IdCardType,omitempty" name:"IdCardType"`

	// 签署回调地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 签署任务ID，标识每一次的流程发送
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignId *string `json:"SignId,omitempty" name:"SignId"`
}

type FlowFileInfo struct {
	// 文件序号
	FileIndex *int64 `json:"FileIndex,omitempty" name:"FileIndex"`

	// 文件类型
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 文件的MD5码
	FileMd5 *string `json:"FileMd5,omitempty" name:"FileMd5"`

	// 文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 文件大小，单位为Byte
	FileSize *int64 `json:"FileSize,omitempty" name:"FileSize"`

	// 文件创建时间戳
	CreatedOn *int64 `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// 文件的下载地址
	Url *string `json:"Url,omitempty" name:"Url"`
}

type FlowInfo struct {
	// 合同名字
	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`

	// 签署截止时间戳，超过有效签署时间则该签署流程失败
	Deadline *int64 `json:"Deadline,omitempty" name:"Deadline"`

	// 合同描述
	FlowDescription *string `json:"FlowDescription,omitempty" name:"FlowDescription"`

	// 合同类型：
	// 1. “劳务”
	// 2. “销售”
	// 3. “租赁”
	// 4. “其他”
	FlowType *string `json:"FlowType,omitempty" name:"FlowType"`

	// 回调地址
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 用户自定义数据
	UserData *string `json:"UserData,omitempty" name:"UserData"`
}

// Predefined struct for user
type GenerateOrganizationSealRequestParams struct {
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 印章类型：
	// OFFICIAL-公章
	// SPECIAL_FINANCIAL-财务专用章
	// CONTRACT-合同专用章
	// LEGAL_REPRESENTATIVE-法定代表人章
	// SPECIAL_NATIONWIDE_INVOICE-发票专用章
	// OTHER-其他
	SealType *string `json:"SealType,omitempty" name:"SealType"`

	// 请求生成企业印章的客户端Ip
	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`

	// 电子印章名称
	SealName *string `json:"SealName,omitempty" name:"SealName"`

	// 企业印章横向文字，最多可填8个汉字（可不填，默认为"电子签名专用章"）
	SealHorizontalText *string `json:"SealHorizontalText,omitempty" name:"SealHorizontalText"`

	// 是否是默认印章 true：是，false：否
	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`
}

type GenerateOrganizationSealRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 印章类型：
	// OFFICIAL-公章
	// SPECIAL_FINANCIAL-财务专用章
	// CONTRACT-合同专用章
	// LEGAL_REPRESENTATIVE-法定代表人章
	// SPECIAL_NATIONWIDE_INVOICE-发票专用章
	// OTHER-其他
	SealType *string `json:"SealType,omitempty" name:"SealType"`

	// 请求生成企业印章的客户端Ip
	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`

	// 电子印章名称
	SealName *string `json:"SealName,omitempty" name:"SealName"`

	// 企业印章横向文字，最多可填8个汉字（可不填，默认为"电子签名专用章"）
	SealHorizontalText *string `json:"SealHorizontalText,omitempty" name:"SealHorizontalText"`

	// 是否是默认印章 true：是，false：否
	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`
}

func (r *GenerateOrganizationSealRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateOrganizationSealRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "SealType")
	delete(f, "SourceIp")
	delete(f, "SealName")
	delete(f, "SealHorizontalText")
	delete(f, "IsDefault")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GenerateOrganizationSealRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenerateOrganizationSealResponseParams struct {
	// 电子印章Id
	SealId *string `json:"SealId,omitempty" name:"SealId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GenerateOrganizationSealResponse struct {
	*tchttp.BaseResponse
	Response *GenerateOrganizationSealResponseParams `json:"Response"`
}

func (r *GenerateOrganizationSealResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateOrganizationSealResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenerateUserSealRequestParams struct {
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 请求生成个人印章的客户端IP
	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`

	// 电子印章名称
	SealName *string `json:"SealName,omitempty" name:"SealName"`

	// 是否是默认印章 true：是，false：否
	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`
}

type GenerateUserSealRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 请求生成个人印章的客户端IP
	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`

	// 电子印章名称
	SealName *string `json:"SealName,omitempty" name:"SealName"`

	// 是否是默认印章 true：是，false：否
	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`
}

func (r *GenerateUserSealRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateUserSealRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "UserId")
	delete(f, "SourceIp")
	delete(f, "SealName")
	delete(f, "IsDefault")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GenerateUserSealRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenerateUserSealResponseParams struct {
	// 电子印章Id
	SealId *string `json:"SealId,omitempty" name:"SealId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GenerateUserSealResponse struct {
	*tchttp.BaseResponse
	Response *GenerateUserSealResponseParams `json:"Response"`
}

func (r *GenerateUserSealResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateUserSealResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyOrganizationDefaultSealRequestParams struct {
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 重新指定的默认印章ID
	SealId *string `json:"SealId,omitempty" name:"SealId"`

	// 请求重新指定企业默认印章的客户端IP
	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`
}

type ModifyOrganizationDefaultSealRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 重新指定的默认印章ID
	SealId *string `json:"SealId,omitempty" name:"SealId"`

	// 请求重新指定企业默认印章的客户端IP
	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`
}

func (r *ModifyOrganizationDefaultSealRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyOrganizationDefaultSealRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "SealId")
	delete(f, "SourceIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyOrganizationDefaultSealRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyOrganizationDefaultSealResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyOrganizationDefaultSealResponse struct {
	*tchttp.BaseResponse
	Response *ModifyOrganizationDefaultSealResponseParams `json:"Response"`
}

func (r *ModifyOrganizationDefaultSealResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyOrganizationDefaultSealResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySealRequestParams struct {
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 请求更新印章的客户端IP
	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`

	// 电子印章ID。若为空，则修改个人/机构的默认印章。
	SealId *string `json:"SealId,omitempty" name:"SealId"`

	// 电子印章名称
	SealName *string `json:"SealName,omitempty" name:"SealName"`

	// 印章图片，base64编码（与FileId参数二选一，同时传入参数时优先使用Image参数）
	Image *string `json:"Image,omitempty" name:"Image"`

	// 印章图片文件ID（与Image参数二选一，同时传入参数时优先使用Image参数）
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 需要更新印章的用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

type ModifySealRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 请求更新印章的客户端IP
	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`

	// 电子印章ID。若为空，则修改个人/机构的默认印章。
	SealId *string `json:"SealId,omitempty" name:"SealId"`

	// 电子印章名称
	SealName *string `json:"SealName,omitempty" name:"SealName"`

	// 印章图片，base64编码（与FileId参数二选一，同时传入参数时优先使用Image参数）
	Image *string `json:"Image,omitempty" name:"Image"`

	// 印章图片文件ID（与Image参数二选一，同时传入参数时优先使用Image参数）
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 需要更新印章的用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

func (r *ModifySealRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySealRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "SourceIp")
	delete(f, "SealId")
	delete(f, "SealName")
	delete(f, "Image")
	delete(f, "FileId")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySealRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySealResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySealResponse struct {
	*tchttp.BaseResponse
	Response *ModifySealResponseParams `json:"Response"`
}

func (r *ModifySealResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySealResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubOrganizationInfoRequestParams struct {
	// 调用方信息，该接口 SubOrganizationId 字段与 OpenId 字段二者至少需要传入一个，全部传入时则使用 SubOrganizationId 信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 机构在第三方的唯一标识，32位定长字符串，与 Caller 中 SubOrgnizationId 二者至少需要传入一个，全部传入时则使用 SubOrganizationId 信息
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 机构名称全称，修改后机构状态将变为未实名，需要调用实名接口重新实名。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 机构类型可选值：
	// 1. ENTERPRISE - 企业；
	// 2. INDIVIDUALBIZ - 个体工商户；
	// 3. PUBLICINSTITUTION - 政府/事业单位
	// 4. OTHERS - 其他组织
	OrganizationType *string `json:"OrganizationType,omitempty" name:"OrganizationType"`

	// 机构证件照片文件，base64编码。支持jpg，jpeg，png格式；如果传值，则重新上传文件后，机构状态将变为未实名，需要调用实名接口重新实名。
	BizLicenseFile *string `json:"BizLicenseFile,omitempty" name:"BizLicenseFile"`

	// 机构证件照片文件名
	BizLicenseFileName *string `json:"BizLicenseFileName,omitempty" name:"BizLicenseFileName"`

	// 机构法人/经营者姓名
	LegalName *string `json:"LegalName,omitempty" name:"LegalName"`

	// 机构法人/经营者证件类型，可选值：ID_CARD - 居民身份证。OrganizationType 为 ENTERPRISE、INDIVIDUALBIZ 时，此项必填，其他情况选填。
	LegalIdCardType *string `json:"LegalIdCardType,omitempty" name:"LegalIdCardType"`

	// 机构法人/经营者证件号码。OrganizationType 为 ENTERPRISE、INDIVIDUALBIZ 时，此项必填，其他情况选填
	LegalIdCardNumber *string `json:"LegalIdCardNumber,omitempty" name:"LegalIdCardNumber"`

	// 机构法人/经营者/联系人手机号码
	LegalMobile *string `json:"LegalMobile,omitempty" name:"LegalMobile"`

	// 组织联系人姓名
	ContactName *string `json:"ContactName,omitempty" name:"ContactName"`

	// 企业联系地址
	ContactAddress *Address `json:"ContactAddress,omitempty" name:"ContactAddress"`

	// 机构电子邮箱
	Email *string `json:"Email,omitempty" name:"Email"`
}

type ModifySubOrganizationInfoRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息，该接口 SubOrganizationId 字段与 OpenId 字段二者至少需要传入一个，全部传入时则使用 SubOrganizationId 信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 机构在第三方的唯一标识，32位定长字符串，与 Caller 中 SubOrgnizationId 二者至少需要传入一个，全部传入时则使用 SubOrganizationId 信息
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 机构名称全称，修改后机构状态将变为未实名，需要调用实名接口重新实名。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 机构类型可选值：
	// 1. ENTERPRISE - 企业；
	// 2. INDIVIDUALBIZ - 个体工商户；
	// 3. PUBLICINSTITUTION - 政府/事业单位
	// 4. OTHERS - 其他组织
	OrganizationType *string `json:"OrganizationType,omitempty" name:"OrganizationType"`

	// 机构证件照片文件，base64编码。支持jpg，jpeg，png格式；如果传值，则重新上传文件后，机构状态将变为未实名，需要调用实名接口重新实名。
	BizLicenseFile *string `json:"BizLicenseFile,omitempty" name:"BizLicenseFile"`

	// 机构证件照片文件名
	BizLicenseFileName *string `json:"BizLicenseFileName,omitempty" name:"BizLicenseFileName"`

	// 机构法人/经营者姓名
	LegalName *string `json:"LegalName,omitempty" name:"LegalName"`

	// 机构法人/经营者证件类型，可选值：ID_CARD - 居民身份证。OrganizationType 为 ENTERPRISE、INDIVIDUALBIZ 时，此项必填，其他情况选填。
	LegalIdCardType *string `json:"LegalIdCardType,omitempty" name:"LegalIdCardType"`

	// 机构法人/经营者证件号码。OrganizationType 为 ENTERPRISE、INDIVIDUALBIZ 时，此项必填，其他情况选填
	LegalIdCardNumber *string `json:"LegalIdCardNumber,omitempty" name:"LegalIdCardNumber"`

	// 机构法人/经营者/联系人手机号码
	LegalMobile *string `json:"LegalMobile,omitempty" name:"LegalMobile"`

	// 组织联系人姓名
	ContactName *string `json:"ContactName,omitempty" name:"ContactName"`

	// 企业联系地址
	ContactAddress *Address `json:"ContactAddress,omitempty" name:"ContactAddress"`

	// 机构电子邮箱
	Email *string `json:"Email,omitempty" name:"Email"`
}

func (r *ModifySubOrganizationInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubOrganizationInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "OpenId")
	delete(f, "Name")
	delete(f, "OrganizationType")
	delete(f, "BizLicenseFile")
	delete(f, "BizLicenseFileName")
	delete(f, "LegalName")
	delete(f, "LegalIdCardType")
	delete(f, "LegalIdCardNumber")
	delete(f, "LegalMobile")
	delete(f, "ContactName")
	delete(f, "ContactAddress")
	delete(f, "Email")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySubOrganizationInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubOrganizationInfoResponseParams struct {
	// 子机构ID
	SubOrganizationId *string `json:"SubOrganizationId,omitempty" name:"SubOrganizationId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySubOrganizationInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifySubOrganizationInfoResponseParams `json:"Response"`
}

func (r *ModifySubOrganizationInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubOrganizationInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserDefaultSealRequestParams struct {
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 用户唯一标识，需要重新指定默认印章的用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 重新指定的默认印章ID
	SealId *string `json:"SealId,omitempty" name:"SealId"`

	// 请求重新指定个人默认印章的客户端IP
	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`
}

type ModifyUserDefaultSealRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 用户唯一标识，需要重新指定默认印章的用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 重新指定的默认印章ID
	SealId *string `json:"SealId,omitempty" name:"SealId"`

	// 请求重新指定个人默认印章的客户端IP
	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`
}

func (r *ModifyUserDefaultSealRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserDefaultSealRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "UserId")
	delete(f, "SealId")
	delete(f, "SourceIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserDefaultSealRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserDefaultSealResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyUserDefaultSealResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUserDefaultSealResponseParams `json:"Response"`
}

func (r *ModifyUserDefaultSealResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserDefaultSealResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserRequestParams struct {
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 第三方平台用户唯一标识; OpenId 和 UserId 二选一填写, 两个都不为空则优先使用UserId
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 腾讯电子签平台用户唯一标识; OpenId 和 UserId 二选一填写, 两个都不为空则优先使用UserId
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户手机号码，不要求唯一
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 用户邮箱，不要求唯一
	Email *string `json:"Email,omitempty" name:"Email"`

	// 用户姓名
	Name *string `json:"Name,omitempty" name:"Name"`
}

type ModifyUserRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 第三方平台用户唯一标识; OpenId 和 UserId 二选一填写, 两个都不为空则优先使用UserId
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 腾讯电子签平台用户唯一标识; OpenId 和 UserId 二选一填写, 两个都不为空则优先使用UserId
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户手机号码，不要求唯一
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 用户邮箱，不要求唯一
	Email *string `json:"Email,omitempty" name:"Email"`

	// 用户姓名
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *ModifyUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "OpenId")
	delete(f, "UserId")
	delete(f, "Mobile")
	delete(f, "Email")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserResponseParams struct {
	// 腾讯电子签平台用户唯一标识
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyUserResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUserResponseParams `json:"Response"`
}

func (r *ModifyUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RejectFlowRequestParams struct {
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 流程编号
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 意愿确认票据。
	// 1. VerifyChannel 为 WEIXINAPP，使用响应的VerifyResult；
	// 2. VerifyChannel 为 FACEID时，使用OrderNo；
	// 3. VerifyChannel 为 VERIFYCODE，使用短信验证码
	// 4. VerifyChannel 为 NONE，传空值
	// （注：普通情况下，VerifyResult不能为None，如您不希望腾讯电子签对用户签署意愿做校验，请提前与客户经理或邮件至e-contract@tencent.com与我们联系）
	VerifyResult *string `json:"VerifyResult,omitempty" name:"VerifyResult"`

	// 意愿确认渠道：
	// 1. WEIXINAPP - 微信小程序
	// 2. FACEID - 慧眼 (默认) 
	// 3. VERIFYCODE - 验证码
	// 4. THIRD - 第三方 (暂不支持)
	// 5. NONE - 无需电子签系统验证
	// （注：普通情况下，VerifyChannel不能为None，如您不希望腾讯电子签对用户签署意愿做校验，请提前与客户经理或邮件至e-contract@tencent.com与我们联系）
	VerifyChannel *string `json:"VerifyChannel,omitempty" name:"VerifyChannel"`

	// 客户端来源IP
	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`

	// 拒签原因
	RejectMessage *string `json:"RejectMessage,omitempty" name:"RejectMessage"`

	// 签署参与者编号
	SignId *string `json:"SignId,omitempty" name:"SignId"`
}

type RejectFlowRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 流程编号
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 意愿确认票据。
	// 1. VerifyChannel 为 WEIXINAPP，使用响应的VerifyResult；
	// 2. VerifyChannel 为 FACEID时，使用OrderNo；
	// 3. VerifyChannel 为 VERIFYCODE，使用短信验证码
	// 4. VerifyChannel 为 NONE，传空值
	// （注：普通情况下，VerifyResult不能为None，如您不希望腾讯电子签对用户签署意愿做校验，请提前与客户经理或邮件至e-contract@tencent.com与我们联系）
	VerifyResult *string `json:"VerifyResult,omitempty" name:"VerifyResult"`

	// 意愿确认渠道：
	// 1. WEIXINAPP - 微信小程序
	// 2. FACEID - 慧眼 (默认) 
	// 3. VERIFYCODE - 验证码
	// 4. THIRD - 第三方 (暂不支持)
	// 5. NONE - 无需电子签系统验证
	// （注：普通情况下，VerifyChannel不能为None，如您不希望腾讯电子签对用户签署意愿做校验，请提前与客户经理或邮件至e-contract@tencent.com与我们联系）
	VerifyChannel *string `json:"VerifyChannel,omitempty" name:"VerifyChannel"`

	// 客户端来源IP
	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`

	// 拒签原因
	RejectMessage *string `json:"RejectMessage,omitempty" name:"RejectMessage"`

	// 签署参与者编号
	SignId *string `json:"SignId,omitempty" name:"SignId"`
}

func (r *RejectFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RejectFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "FlowId")
	delete(f, "VerifyResult")
	delete(f, "VerifyChannel")
	delete(f, "SourceIp")
	delete(f, "RejectMessage")
	delete(f, "SignId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RejectFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RejectFlowResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RejectFlowResponse struct {
	*tchttp.BaseResponse
	Response *RejectFlowResponseParams `json:"Response"`
}

func (r *RejectFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RejectFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Seal struct {
	// 电子印章ID
	SealId *string `json:"SealId,omitempty" name:"SealId"`

	// 电子印章名称
	SealName *string `json:"SealName,omitempty" name:"SealName"`

	// 电子印章类型
	SealType *string `json:"SealType,omitempty" name:"SealType"`

	// 电子印章来源：
	// CREATE - 通过图片上传
	// GENERATE - 通过文字生成
	SealSource *string `json:"SealSource,omitempty" name:"SealSource"`

	// 电子印章创建者
	Creator *Caller `json:"Creator,omitempty" name:"Creator"`

	// 电子印章创建时间戳
	CreatedOn *int64 `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// 电子印章所有人
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 电子印章URL
	FileUrl *FileUrl `json:"FileUrl,omitempty" name:"FileUrl"`

	// 是否为默认印章，false-非默认，true-默认
	DefaultSeal *bool `json:"DefaultSeal,omitempty" name:"DefaultSeal"`
}

// Predefined struct for user
type SendFlowRequestParams struct {
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 需要推送合同的流程ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 签署人用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 签署控件信息 (支持添加多个控件)
	SignComponents []*Component `json:"SignComponents,omitempty" name:"SignComponents"`

	// 签署人手机号 (如果选择短信验证码签署，则此字段必填)
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 签署人对应的子机构ID，个人签署者此字段不填
	SubOrganizationId *string `json:"SubOrganizationId,omitempty" name:"SubOrganizationId"`

	// 签名后校验方式：
	// 1. WEIXINAPP - 微信小程序；
	// 2. FACEID - 慧眼 (默认) ；
	// 3. VERIFYCODE - 验证码；
	// 4. NONE - 无。此选项为白名单参数，暂不支持公开调用。如需开通权限，请通过客户经理或邮件至e-contract@tencent.com与我们联系；
	// 5. THIRD - 第三方 (暂不支持)
	VerifyChannel []*string `json:"VerifyChannel,omitempty" name:"VerifyChannel"`

	// 签署链接失效截止时间，默认为7天
	Deadline *int64 `json:"Deadline,omitempty" name:"Deadline"`

	// 是否为最后一个签署人。若为最后一人，本次签署完成以后自动归档。
	IsLastApprover *bool `json:"IsLastApprover,omitempty" name:"IsLastApprover"`

	// 签署完成后，前端跳转的URL
	JumpUrl *string `json:"JumpUrl,omitempty" name:"JumpUrl"`

	// 短信模板。默认使用腾讯电子签官方短信模板，如有自定义需求，请通过客户经理或邮件至e-contract@tencent.com与我们联系。
	SmsTemplate *SmsTemplate `json:"SmsTemplate,omitempty" name:"SmsTemplate"`

	// 签署前置条件：是否要全文阅读，默认否
	IsFullText *bool `json:"IsFullText,omitempty" name:"IsFullText"`

	// 签署前置条件：强制用户阅读待签署文件时长，默认不限制
	PreReadTime *int64 `json:"PreReadTime,omitempty" name:"PreReadTime"`

	// 当前参与者是否支持线下核身,默认为不支持
	CanOffLine *bool `json:"CanOffLine,omitempty" name:"CanOffLine"`

	// 签署任务的回调地址
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`
}

type SendFlowRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 需要推送合同的流程ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 签署人用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 签署控件信息 (支持添加多个控件)
	SignComponents []*Component `json:"SignComponents,omitempty" name:"SignComponents"`

	// 签署人手机号 (如果选择短信验证码签署，则此字段必填)
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 签署人对应的子机构ID，个人签署者此字段不填
	SubOrganizationId *string `json:"SubOrganizationId,omitempty" name:"SubOrganizationId"`

	// 签名后校验方式：
	// 1. WEIXINAPP - 微信小程序；
	// 2. FACEID - 慧眼 (默认) ；
	// 3. VERIFYCODE - 验证码；
	// 4. NONE - 无。此选项为白名单参数，暂不支持公开调用。如需开通权限，请通过客户经理或邮件至e-contract@tencent.com与我们联系；
	// 5. THIRD - 第三方 (暂不支持)
	VerifyChannel []*string `json:"VerifyChannel,omitempty" name:"VerifyChannel"`

	// 签署链接失效截止时间，默认为7天
	Deadline *int64 `json:"Deadline,omitempty" name:"Deadline"`

	// 是否为最后一个签署人。若为最后一人，本次签署完成以后自动归档。
	IsLastApprover *bool `json:"IsLastApprover,omitempty" name:"IsLastApprover"`

	// 签署完成后，前端跳转的URL
	JumpUrl *string `json:"JumpUrl,omitempty" name:"JumpUrl"`

	// 短信模板。默认使用腾讯电子签官方短信模板，如有自定义需求，请通过客户经理或邮件至e-contract@tencent.com与我们联系。
	SmsTemplate *SmsTemplate `json:"SmsTemplate,omitempty" name:"SmsTemplate"`

	// 签署前置条件：是否要全文阅读，默认否
	IsFullText *bool `json:"IsFullText,omitempty" name:"IsFullText"`

	// 签署前置条件：强制用户阅读待签署文件时长，默认不限制
	PreReadTime *int64 `json:"PreReadTime,omitempty" name:"PreReadTime"`

	// 当前参与者是否支持线下核身,默认为不支持
	CanOffLine *bool `json:"CanOffLine,omitempty" name:"CanOffLine"`

	// 签署任务的回调地址
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`
}

func (r *SendFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "FlowId")
	delete(f, "UserId")
	delete(f, "SignComponents")
	delete(f, "Mobile")
	delete(f, "SubOrganizationId")
	delete(f, "VerifyChannel")
	delete(f, "Deadline")
	delete(f, "IsLastApprover")
	delete(f, "JumpUrl")
	delete(f, "SmsTemplate")
	delete(f, "IsFullText")
	delete(f, "PreReadTime")
	delete(f, "CanOffLine")
	delete(f, "CallbackUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendFlowResponseParams struct {
	// 签署任务ID，标识每一次的流程发送
	SignId *string `json:"SignId,omitempty" name:"SignId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SendFlowResponse struct {
	*tchttp.BaseResponse
	Response *SendFlowResponseParams `json:"Response"`
}

func (r *SendFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendFlowUrlRequestParams struct {
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 需要推送合同的流程ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 签署人ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 签署控件信息 (支持添加多个控件)
	SignComponents []*Component `json:"SignComponents,omitempty" name:"SignComponents"`

	// 签署人手机号 (如果选择短信验证码签署，则此字段必填)
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 签署人对应的子机构ID，个人签署者此字段不填
	SubOrganizationId *string `json:"SubOrganizationId,omitempty" name:"SubOrganizationId"`

	// 签名后校验方式：
	// 1. WEIXINAPP - 微信小程序；
	// 2. FACEID - 慧眼 (默认) ；
	// 3. VERIFYCODE - 验证码；
	// 4. NONE - 无。此选项为白名单参数，暂不支持公开调用。如需开通权限，请通过客户经理或邮件至e-contract@tencent.com与我们联系；
	// 5. THIRD - 第三方 (暂不支持)
	// 6. OFFLINE - 线下人工审核
	VerifyChannel []*string `json:"VerifyChannel,omitempty" name:"VerifyChannel"`

	// 签署链接失效截止时间，默认为7天
	Deadline *int64 `json:"Deadline,omitempty" name:"Deadline"`

	// 是否为最后一个签署人。若为最后一人，本次签署完成以后自动归档
	IsLastApprover *bool `json:"IsLastApprover,omitempty" name:"IsLastApprover"`

	// 签署完成后，前端跳转的url
	JumpUrl *string `json:"JumpUrl,omitempty" name:"JumpUrl"`

	// 短信模板
	// 默认使用腾讯电子签官方短信模板，如有自定义需求，请通过客户经理或邮件至e-contract@tencent.com与我们联系。
	SmsTemplate *SmsTemplate `json:"SmsTemplate,omitempty" name:"SmsTemplate"`

	// 签署前置条件：是否要全文阅读，默认否
	IsFullText *bool `json:"IsFullText,omitempty" name:"IsFullText"`

	// 签署前置条件：强制用户阅读待签署文件时长，默认不限制
	PreReadTime *int64 `json:"PreReadTime,omitempty" name:"PreReadTime"`

	// 当前参与者是否支持线下核身,默认为不支持
	CanOffLine *bool `json:"CanOffLine,omitempty" name:"CanOffLine"`

	// 签署任务的回调地址
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`
}

type SendFlowUrlRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 需要推送合同的流程ID
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 签署人ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 签署控件信息 (支持添加多个控件)
	SignComponents []*Component `json:"SignComponents,omitempty" name:"SignComponents"`

	// 签署人手机号 (如果选择短信验证码签署，则此字段必填)
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 签署人对应的子机构ID，个人签署者此字段不填
	SubOrganizationId *string `json:"SubOrganizationId,omitempty" name:"SubOrganizationId"`

	// 签名后校验方式：
	// 1. WEIXINAPP - 微信小程序；
	// 2. FACEID - 慧眼 (默认) ；
	// 3. VERIFYCODE - 验证码；
	// 4. NONE - 无。此选项为白名单参数，暂不支持公开调用。如需开通权限，请通过客户经理或邮件至e-contract@tencent.com与我们联系；
	// 5. THIRD - 第三方 (暂不支持)
	// 6. OFFLINE - 线下人工审核
	VerifyChannel []*string `json:"VerifyChannel,omitempty" name:"VerifyChannel"`

	// 签署链接失效截止时间，默认为7天
	Deadline *int64 `json:"Deadline,omitempty" name:"Deadline"`

	// 是否为最后一个签署人。若为最后一人，本次签署完成以后自动归档
	IsLastApprover *bool `json:"IsLastApprover,omitempty" name:"IsLastApprover"`

	// 签署完成后，前端跳转的url
	JumpUrl *string `json:"JumpUrl,omitempty" name:"JumpUrl"`

	// 短信模板
	// 默认使用腾讯电子签官方短信模板，如有自定义需求，请通过客户经理或邮件至e-contract@tencent.com与我们联系。
	SmsTemplate *SmsTemplate `json:"SmsTemplate,omitempty" name:"SmsTemplate"`

	// 签署前置条件：是否要全文阅读，默认否
	IsFullText *bool `json:"IsFullText,omitempty" name:"IsFullText"`

	// 签署前置条件：强制用户阅读待签署文件时长，默认不限制
	PreReadTime *int64 `json:"PreReadTime,omitempty" name:"PreReadTime"`

	// 当前参与者是否支持线下核身,默认为不支持
	CanOffLine *bool `json:"CanOffLine,omitempty" name:"CanOffLine"`

	// 签署任务的回调地址
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`
}

func (r *SendFlowUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendFlowUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "FlowId")
	delete(f, "UserId")
	delete(f, "SignComponents")
	delete(f, "Mobile")
	delete(f, "SubOrganizationId")
	delete(f, "VerifyChannel")
	delete(f, "Deadline")
	delete(f, "IsLastApprover")
	delete(f, "JumpUrl")
	delete(f, "SmsTemplate")
	delete(f, "IsFullText")
	delete(f, "PreReadTime")
	delete(f, "CanOffLine")
	delete(f, "CallbackUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendFlowUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendFlowUrlResponseParams struct {
	// 签署任务ID，标识每一次的流程发送
	SignId *string `json:"SignId,omitempty" name:"SignId"`

	// 签署链接
	SignUrl *string `json:"SignUrl,omitempty" name:"SignUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SendFlowUrlResponse struct {
	*tchttp.BaseResponse
	Response *SendFlowUrlResponseParams `json:"Response"`
}

func (r *SendFlowUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendFlowUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendSignInnerVerifyCodeRequestParams struct {
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 手机号
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 验证码类型，取值(SIGN)
	VerifyType *string `json:"VerifyType,omitempty" name:"VerifyType"`

	// 用户 id
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 模板 id
	VerifyTemplateId *string `json:"VerifyTemplateId,omitempty" name:"VerifyTemplateId"`

	// 签名
	VerifySign *string `json:"VerifySign,omitempty" name:"VerifySign"`

	// 流程(目录) id
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 三要素检测结果
	CheckThreeElementResult *int64 `json:"CheckThreeElementResult,omitempty" name:"CheckThreeElementResult"`
}

type SendSignInnerVerifyCodeRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 手机号
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 验证码类型，取值(SIGN)
	VerifyType *string `json:"VerifyType,omitempty" name:"VerifyType"`

	// 用户 id
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 模板 id
	VerifyTemplateId *string `json:"VerifyTemplateId,omitempty" name:"VerifyTemplateId"`

	// 签名
	VerifySign *string `json:"VerifySign,omitempty" name:"VerifySign"`

	// 流程(目录) id
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 三要素检测结果
	CheckThreeElementResult *int64 `json:"CheckThreeElementResult,omitempty" name:"CheckThreeElementResult"`
}

func (r *SendSignInnerVerifyCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendSignInnerVerifyCodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "Mobile")
	delete(f, "VerifyType")
	delete(f, "UserId")
	delete(f, "VerifyTemplateId")
	delete(f, "VerifySign")
	delete(f, "FlowId")
	delete(f, "CheckThreeElementResult")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendSignInnerVerifyCodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendSignInnerVerifyCodeResponseParams struct {
	// true: 验证码正确，false: 验证码错误
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SendSignInnerVerifyCodeResponse struct {
	*tchttp.BaseResponse
	Response *SendSignInnerVerifyCodeResponseParams `json:"Response"`
}

func (r *SendSignInnerVerifyCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendSignInnerVerifyCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SignFlowRequestParams struct {
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 流程编号
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 意愿确认票据。
	// 1. VerifyChannel 为 WEIXINAPP，使用响应的VerifyResult；
	// 2. VerifyChannel 为 FACEID时，使用OrderNo；
	// 3. VerifyChannel 为 VERIFYCODE，使用短信验证码
	// 4. VerifyChannel 为 NONE，传空值
	// （注：普通情况下，VerifyResult不能为None，如您不希望腾讯电子签对用户签署意愿做校验，请提前与客户经理或邮件至e-contract@tencent.com与我们联系）
	VerifyResult *string `json:"VerifyResult,omitempty" name:"VerifyResult"`

	// 意愿确认渠道：
	// 1. WEIXINAPP - 微信小程序
	// 2. FACEID - 慧眼 (默认) 
	// 3. VERIFYCODE - 验证码
	// 4. THIRD - 第三方 (暂不支持)
	// 5. NONE - 无需电子签系统验证
	// （注：普通情况下，VerifyChannel不能为None，如您不希望腾讯电子签对用户签署意愿做校验，请提前与客户经理或邮件至e-contract@tencent.com与我们联系）
	VerifyChannel *string `json:"VerifyChannel,omitempty" name:"VerifyChannel"`

	// 客户端来源IP
	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`

	// 签署内容
	SignSeals []*SignSeal `json:"SignSeals,omitempty" name:"SignSeals"`

	// 签署备注
	ApproveMessage *string `json:"ApproveMessage,omitempty" name:"ApproveMessage"`

	// 签署参与者编号
	SignId *string `json:"SignId,omitempty" name:"SignId"`
}

type SignFlowRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 流程编号
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 意愿确认票据。
	// 1. VerifyChannel 为 WEIXINAPP，使用响应的VerifyResult；
	// 2. VerifyChannel 为 FACEID时，使用OrderNo；
	// 3. VerifyChannel 为 VERIFYCODE，使用短信验证码
	// 4. VerifyChannel 为 NONE，传空值
	// （注：普通情况下，VerifyResult不能为None，如您不希望腾讯电子签对用户签署意愿做校验，请提前与客户经理或邮件至e-contract@tencent.com与我们联系）
	VerifyResult *string `json:"VerifyResult,omitempty" name:"VerifyResult"`

	// 意愿确认渠道：
	// 1. WEIXINAPP - 微信小程序
	// 2. FACEID - 慧眼 (默认) 
	// 3. VERIFYCODE - 验证码
	// 4. THIRD - 第三方 (暂不支持)
	// 5. NONE - 无需电子签系统验证
	// （注：普通情况下，VerifyChannel不能为None，如您不希望腾讯电子签对用户签署意愿做校验，请提前与客户经理或邮件至e-contract@tencent.com与我们联系）
	VerifyChannel *string `json:"VerifyChannel,omitempty" name:"VerifyChannel"`

	// 客户端来源IP
	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`

	// 签署内容
	SignSeals []*SignSeal `json:"SignSeals,omitempty" name:"SignSeals"`

	// 签署备注
	ApproveMessage *string `json:"ApproveMessage,omitempty" name:"ApproveMessage"`

	// 签署参与者编号
	SignId *string `json:"SignId,omitempty" name:"SignId"`
}

func (r *SignFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SignFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "FlowId")
	delete(f, "VerifyResult")
	delete(f, "VerifyChannel")
	delete(f, "SourceIp")
	delete(f, "SignSeals")
	delete(f, "ApproveMessage")
	delete(f, "SignId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SignFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SignFlowResponseParams struct {
	// 签署任务状态。签署成功 - SUCCESS、提交审核 - REVIEW
	Status *string `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SignFlowResponse struct {
	*tchttp.BaseResponse
	Response *SignFlowResponseParams `json:"Response"`
}

func (r *SignFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SignFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SignSeal struct {
	// 签署控件ID
	ComponentId *string `json:"ComponentId,omitempty" name:"ComponentId"`

	// 签署印章类型:
	// SIGN_SIGNATURE - 签名
	// SIGN_SEAL - 印章
	// SIGN_DATE - 日期
	// SIGN_IMAGE - 图片
	SignType *string `json:"SignType,omitempty" name:"SignType"`

	// 合同文件ID
	FileIndex *int64 `json:"FileIndex,omitempty" name:"FileIndex"`

	// 印章ID，仅当 SignType 为 SIGN_SEAL 时必填
	SealId *string `json:"SealId,omitempty" name:"SealId"`

	// 签名内容，仅当 SignType 为SIGN_SIGNATURE或SIGN_IMAGE 时必填，base64编码
	SealContent *string `json:"SealContent,omitempty" name:"SealContent"`
}

type SmsTemplate struct {
	// 模板ID，必须填写已审核通过的模板ID。模板ID可登录短信控制台查看。
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 短信签名内容，使用UTF-8编码，必须填写已审核通过的签名，签名信息可登录短信控制台查看。
	Sign *string `json:"Sign,omitempty" name:"Sign"`
}

type SubOrganizationDetail struct {
	// 组织ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 机构名称全称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 机构电子邮箱
	Email *string `json:"Email,omitempty" name:"Email"`

	// 机构证件号码类型
	IdCardType *string `json:"IdCardType,omitempty" name:"IdCardType"`

	// 机构证件号码
	IdCardNumber *string `json:"IdCardNumber,omitempty" name:"IdCardNumber"`

	// 机构类型
	OrganizationType *string `json:"OrganizationType,omitempty" name:"OrganizationType"`

	// 机构证件文件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdCardFileType *string `json:"IdCardFileType,omitempty" name:"IdCardFileType"`

	// 机构证件照片文件，base64编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	BizLicenseFile *string `json:"BizLicenseFile,omitempty" name:"BizLicenseFile"`

	// 机构证件照片文件名
	BizLicenseFileName *string `json:"BizLicenseFileName,omitempty" name:"BizLicenseFileName"`

	// 机构法人/经营者姓名
	LegalName *string `json:"LegalName,omitempty" name:"LegalName"`

	// 机构法人/经营者证件类型
	LegalIdCardType *string `json:"LegalIdCardType,omitempty" name:"LegalIdCardType"`

	// 机构法人/经营者证件号码
	LegalIdCardNumber *string `json:"LegalIdCardNumber,omitempty" name:"LegalIdCardNumber"`

	// 机构法人/经营者/联系人手机号码
	LegalMobile *string `json:"LegalMobile,omitempty" name:"LegalMobile"`

	// 组织联系人姓名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContactName *string `json:"ContactName,omitempty" name:"ContactName"`

	// 机构实名状态
	VerifyStatus *string `json:"VerifyStatus,omitempty" name:"VerifyStatus"`

	// 机构通过实名时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerifiedOn *int64 `json:"VerifiedOn,omitempty" name:"VerifiedOn"`

	// 机构创建时间
	CreatedOn *int64 `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// 机构更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedOn *int64 `json:"UpdatedOn,omitempty" name:"UpdatedOn"`

	// 实名认证的客户端IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerifyClientIp *string `json:"VerifyClientIp,omitempty" name:"VerifyClientIp"`

	// 实名认证的服务器IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerifyServerIp *string `json:"VerifyServerIp,omitempty" name:"VerifyServerIp"`

	// 企业联系地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContactAddress *Address `json:"ContactAddress,omitempty" name:"ContactAddress"`
}

type UploadFile struct {
	// Base64编码后的文件内容
	FileBody *string `json:"FileBody,omitempty" name:"FileBody"`

	// 文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`
}

// Predefined struct for user
type UploadFilesRequestParams struct {
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 文件对应业务类型，用于区分文件存储路径：
	// 1. TEMPLATE - 模版； 文件类型：.pdf/.html
	// 2. DOCUMENT - 签署过程及签署后的合同文档 文件类型：.pdf/.html
	// 3. FLOW - 签署过程 文件类型：.pdf/.html
	// 4. SEAL - 印章； 文件类型：.jpg/.jpeg/.png
	// 5. BUSINESSLICENSE - 营业执照 文件类型：.jpg/.jpeg/.png
	// 6. IDCARD - 身份证 文件类型：.jpg/.jpeg/.png
	BusinessType *string `json:"BusinessType,omitempty" name:"BusinessType"`

	// 上传文件内容数组，最多支持20个文件
	FileInfos []*UploadFile `json:"FileInfos,omitempty" name:"FileInfos"`

	// 上传文件链接数组，最多支持20个URL
	FileUrls []*string `json:"FileUrls,omitempty" name:"FileUrls"`

	// 是否将pdf灰色矩阵置白
	// true--是，处理置白
	// false--否，不处理
	CoverRect *bool `json:"CoverRect,omitempty" name:"CoverRect"`

	// 特殊文件类型需要指定文件类型：
	// HTML-- .html文件
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 用户自定义ID数组，与上传文件一一对应
	CustomIds []*string `json:"CustomIds,omitempty" name:"CustomIds"`
}

type UploadFilesRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 文件对应业务类型，用于区分文件存储路径：
	// 1. TEMPLATE - 模版； 文件类型：.pdf/.html
	// 2. DOCUMENT - 签署过程及签署后的合同文档 文件类型：.pdf/.html
	// 3. FLOW - 签署过程 文件类型：.pdf/.html
	// 4. SEAL - 印章； 文件类型：.jpg/.jpeg/.png
	// 5. BUSINESSLICENSE - 营业执照 文件类型：.jpg/.jpeg/.png
	// 6. IDCARD - 身份证 文件类型：.jpg/.jpeg/.png
	BusinessType *string `json:"BusinessType,omitempty" name:"BusinessType"`

	// 上传文件内容数组，最多支持20个文件
	FileInfos []*UploadFile `json:"FileInfos,omitempty" name:"FileInfos"`

	// 上传文件链接数组，最多支持20个URL
	FileUrls []*string `json:"FileUrls,omitempty" name:"FileUrls"`

	// 是否将pdf灰色矩阵置白
	// true--是，处理置白
	// false--否，不处理
	CoverRect *bool `json:"CoverRect,omitempty" name:"CoverRect"`

	// 特殊文件类型需要指定文件类型：
	// HTML-- .html文件
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 用户自定义ID数组，与上传文件一一对应
	CustomIds []*string `json:"CustomIds,omitempty" name:"CustomIds"`
}

func (r *UploadFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "BusinessType")
	delete(f, "FileInfos")
	delete(f, "FileUrls")
	delete(f, "CoverRect")
	delete(f, "FileType")
	delete(f, "CustomIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadFilesResponseParams struct {
	// 文件id数组
	FileIds []*string `json:"FileIds,omitempty" name:"FileIds"`

	// 上传成功文件数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UploadFilesResponse struct {
	*tchttp.BaseResponse
	Response *UploadFilesResponseParams `json:"Response"`
}

func (r *UploadFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserDescribe struct {
	// 用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 手机号，隐藏中间4位数字，用*代替
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 注册时间点 (UNIX时间戳)
	CreatedOn *int64 `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// 实名认证状态：
	// 0 - 未实名；
	// 1 - 通过实名
	VerifyStatus *int64 `json:"VerifyStatus,omitempty" name:"VerifyStatus"`

	// 真实姓名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 实名认证通过时间 (UNIX时间戳)
	VerifiedOn *int64 `json:"VerifiedOn,omitempty" name:"VerifiedOn"`

	// 身份证件类型; 
	// ID_CARD - 居民身份证；
	// PASSPORT - 护照；
	// MAINLAND_TRAVEL_PERMIT_FOR_HONGKONG_AND_MACAO_RESIDENTS - 港澳居民来往内地通行证；
	// MAINLAND_TRAVEL_PERMIT_FOR_TAIWAN_RESIDENTS - 台湾居民来往大陆通行证；
	// HOUSEHOLD_REGISTER - 户口本；
	// TEMP_ID_CARD - 临时居民身份证
	IdCardType *string `json:"IdCardType,omitempty" name:"IdCardType"`

	// 身份证件号码 (脱敏)
	IdCardNumber *string `json:"IdCardNumber,omitempty" name:"IdCardNumber"`
}

// Predefined struct for user
type VerifySubOrganizationRequestParams struct {
	// 调用方信息，该接口SubOrganizationId必填
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 机构在第三方的唯一标识，32位定长字符串，与 Caller 中 SubOrgnizationId 二者至少需要传入一个，全部传入时则使用 SubOrganizationId 信息
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`
}

type VerifySubOrganizationRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息，该接口SubOrganizationId必填
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 机构在第三方的唯一标识，32位定长字符串，与 Caller 中 SubOrgnizationId 二者至少需要传入一个，全部传入时则使用 SubOrganizationId 信息
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`
}

func (r *VerifySubOrganizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifySubOrganizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "OpenId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifySubOrganizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifySubOrganizationResponseParams struct {
	// 子机构ID
	SubOrganizationId *string `json:"SubOrganizationId,omitempty" name:"SubOrganizationId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type VerifySubOrganizationResponse struct {
	*tchttp.BaseResponse
	Response *VerifySubOrganizationResponseParams `json:"Response"`
}

func (r *VerifySubOrganizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifySubOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyUserRequestParams struct {
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 电子签平台用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 是否需要下发个人长效证书，默认为false
	// 注：如您有下发个人长效证书需求，请提前邮件至e-contract@oa.com进行申请。
	CertificateRequired *bool `json:"CertificateRequired,omitempty" name:"CertificateRequired"`
}

type VerifyUserRequest struct {
	*tchttp.BaseRequest
	
	// 调用方信息
	Caller *Caller `json:"Caller,omitempty" name:"Caller"`

	// 电子签平台用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 是否需要下发个人长效证书，默认为false
	// 注：如您有下发个人长效证书需求，请提前邮件至e-contract@oa.com进行申请。
	CertificateRequired *bool `json:"CertificateRequired,omitempty" name:"CertificateRequired"`
}

func (r *VerifyUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Caller")
	delete(f, "UserId")
	delete(f, "CertificateRequired")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyUserResponseParams struct {
	// 电子签平台用户ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type VerifyUserResponse struct {
	*tchttp.BaseResponse
	Response *VerifyUserResponseParams `json:"Response"`
}

func (r *VerifyUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}