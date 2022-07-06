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

// Predefined struct for user
type AddContractRequestParams struct {
	// 收单系统分配的开放ID
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 收单系统分配的密钥
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 机构合同主键（系统有唯一性校验），建议使用合同表的主键ID，防止重复添加合同
	OutContractId *string `json:"OutContractId,omitempty" name:"OutContractId"`

	// 合同编号（系统有唯一性校验）
	Code *string `json:"Code,omitempty" name:"Code"`

	// 支付方式编号
	PaymentId *string `json:"PaymentId,omitempty" name:"PaymentId"`

	// 支付方式行业分类编号
	PaymentClassificationId *string `json:"PaymentClassificationId,omitempty" name:"PaymentClassificationId"`

	// 封顶值（分为单位，无封顶填0）
	PaymentClassificationLimit *string `json:"PaymentClassificationLimit,omitempty" name:"PaymentClassificationLimit"`

	// 商户编号
	MerchantNo *string `json:"MerchantNo,omitempty" name:"MerchantNo"`

	// 签约扣率百分比（如：0.32）
	Fee *string `json:"Fee,omitempty" name:"Fee"`

	// 合同生效日期（yyyy-mm-dd）
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 合同过期日期（yyyy-mm-dd）
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 合同签约人
	SignMan *string `json:"SignMan,omitempty" name:"SignMan"`

	// 签购单名称，建议使用商户招牌名称
	SignName *string `json:"SignName,omitempty" name:"SignName"`

	// 合同签署日期（yyyy-mm-dd）
	SignDate *string `json:"SignDate,omitempty" name:"SignDate"`

	// 是否自动续签（1是，0否）
	AutoSign *string `json:"AutoSign,omitempty" name:"AutoSign"`

	// 联系人
	Contact *string `json:"Contact,omitempty" name:"Contact"`

	// 联系人电话
	ContactTelephone *string `json:"ContactTelephone,omitempty" name:"ContactTelephone"`

	// 合同照片【私密区】
	PictureOne *string `json:"PictureOne,omitempty" name:"PictureOne"`

	// 合同照片【私密区】
	PictureTwo *string `json:"PictureTwo,omitempty" name:"PictureTwo"`

	// 渠道扩展字段，json格式
	ChannelExtJson *string `json:"ChannelExtJson,omitempty" name:"ChannelExtJson"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`

	// 合同选项1（不同支付方式规则不一样，请以支付方式规定的格式传值）
	PaymentOptionOne *string `json:"PaymentOptionOne,omitempty" name:"PaymentOptionOne"`

	// 合同选项2（不同支付方式规则不一样，请以支付方式规定的格式传值）
	PaymentOptionTwo *string `json:"PaymentOptionTwo,omitempty" name:"PaymentOptionTwo"`

	// 合同选项3（不同支付方式规则不一样，请以支付方式规定的格式传值）
	PaymentOptionThree *string `json:"PaymentOptionThree,omitempty" name:"PaymentOptionThree"`

	// 合同选项4（不同支付方式规则不一样，请以支付方式规定的格式传值）
	PaymentOptionFour *string `json:"PaymentOptionFour,omitempty" name:"PaymentOptionFour"`

	// 合同证书选项1（不同支付方式规则不一样，请以支付方式规定的格式传值）
	PaymentOptionFive *string `json:"PaymentOptionFive,omitempty" name:"PaymentOptionFive"`

	// 合同证书选项2（不同支付方式规则不一样，请以支付方式规定的格式传值）
	PaymentOptionSix *string `json:"PaymentOptionSix,omitempty" name:"PaymentOptionSix"`

	// 合同选项5（不同支付方式规则不一样，请以支付方式规定的格式传值）
	PaymentOptionSeven *string `json:"PaymentOptionSeven,omitempty" name:"PaymentOptionSeven"`

	// 合同选项6（不同支付方式规则不一样，请以支付方式规定的格式传值）
	PaymentOptionOther *string `json:"PaymentOptionOther,omitempty" name:"PaymentOptionOther"`

	// 合同选项8
	PaymentOptionTen *string `json:"PaymentOptionTen,omitempty" name:"PaymentOptionTen"`

	// 合同选项7（不同支付方式规则不一样，请以支付方式规定的格式传值）
	PaymentOptionNine *string `json:"PaymentOptionNine,omitempty" name:"PaymentOptionNine"`
}

type AddContractRequest struct {
	*tchttp.BaseRequest
	
	// 收单系统分配的开放ID
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 收单系统分配的密钥
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 机构合同主键（系统有唯一性校验），建议使用合同表的主键ID，防止重复添加合同
	OutContractId *string `json:"OutContractId,omitempty" name:"OutContractId"`

	// 合同编号（系统有唯一性校验）
	Code *string `json:"Code,omitempty" name:"Code"`

	// 支付方式编号
	PaymentId *string `json:"PaymentId,omitempty" name:"PaymentId"`

	// 支付方式行业分类编号
	PaymentClassificationId *string `json:"PaymentClassificationId,omitempty" name:"PaymentClassificationId"`

	// 封顶值（分为单位，无封顶填0）
	PaymentClassificationLimit *string `json:"PaymentClassificationLimit,omitempty" name:"PaymentClassificationLimit"`

	// 商户编号
	MerchantNo *string `json:"MerchantNo,omitempty" name:"MerchantNo"`

	// 签约扣率百分比（如：0.32）
	Fee *string `json:"Fee,omitempty" name:"Fee"`

	// 合同生效日期（yyyy-mm-dd）
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 合同过期日期（yyyy-mm-dd）
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 合同签约人
	SignMan *string `json:"SignMan,omitempty" name:"SignMan"`

	// 签购单名称，建议使用商户招牌名称
	SignName *string `json:"SignName,omitempty" name:"SignName"`

	// 合同签署日期（yyyy-mm-dd）
	SignDate *string `json:"SignDate,omitempty" name:"SignDate"`

	// 是否自动续签（1是，0否）
	AutoSign *string `json:"AutoSign,omitempty" name:"AutoSign"`

	// 联系人
	Contact *string `json:"Contact,omitempty" name:"Contact"`

	// 联系人电话
	ContactTelephone *string `json:"ContactTelephone,omitempty" name:"ContactTelephone"`

	// 合同照片【私密区】
	PictureOne *string `json:"PictureOne,omitempty" name:"PictureOne"`

	// 合同照片【私密区】
	PictureTwo *string `json:"PictureTwo,omitempty" name:"PictureTwo"`

	// 渠道扩展字段，json格式
	ChannelExtJson *string `json:"ChannelExtJson,omitempty" name:"ChannelExtJson"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`

	// 合同选项1（不同支付方式规则不一样，请以支付方式规定的格式传值）
	PaymentOptionOne *string `json:"PaymentOptionOne,omitempty" name:"PaymentOptionOne"`

	// 合同选项2（不同支付方式规则不一样，请以支付方式规定的格式传值）
	PaymentOptionTwo *string `json:"PaymentOptionTwo,omitempty" name:"PaymentOptionTwo"`

	// 合同选项3（不同支付方式规则不一样，请以支付方式规定的格式传值）
	PaymentOptionThree *string `json:"PaymentOptionThree,omitempty" name:"PaymentOptionThree"`

	// 合同选项4（不同支付方式规则不一样，请以支付方式规定的格式传值）
	PaymentOptionFour *string `json:"PaymentOptionFour,omitempty" name:"PaymentOptionFour"`

	// 合同证书选项1（不同支付方式规则不一样，请以支付方式规定的格式传值）
	PaymentOptionFive *string `json:"PaymentOptionFive,omitempty" name:"PaymentOptionFive"`

	// 合同证书选项2（不同支付方式规则不一样，请以支付方式规定的格式传值）
	PaymentOptionSix *string `json:"PaymentOptionSix,omitempty" name:"PaymentOptionSix"`

	// 合同选项5（不同支付方式规则不一样，请以支付方式规定的格式传值）
	PaymentOptionSeven *string `json:"PaymentOptionSeven,omitempty" name:"PaymentOptionSeven"`

	// 合同选项6（不同支付方式规则不一样，请以支付方式规定的格式传值）
	PaymentOptionOther *string `json:"PaymentOptionOther,omitempty" name:"PaymentOptionOther"`

	// 合同选项8
	PaymentOptionTen *string `json:"PaymentOptionTen,omitempty" name:"PaymentOptionTen"`

	// 合同选项7（不同支付方式规则不一样，请以支付方式规定的格式传值）
	PaymentOptionNine *string `json:"PaymentOptionNine,omitempty" name:"PaymentOptionNine"`
}

func (r *AddContractRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddContractRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OpenId")
	delete(f, "OpenKey")
	delete(f, "OutContractId")
	delete(f, "Code")
	delete(f, "PaymentId")
	delete(f, "PaymentClassificationId")
	delete(f, "PaymentClassificationLimit")
	delete(f, "MerchantNo")
	delete(f, "Fee")
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "SignMan")
	delete(f, "SignName")
	delete(f, "SignDate")
	delete(f, "AutoSign")
	delete(f, "Contact")
	delete(f, "ContactTelephone")
	delete(f, "PictureOne")
	delete(f, "PictureTwo")
	delete(f, "ChannelExtJson")
	delete(f, "Profile")
	delete(f, "PaymentOptionOne")
	delete(f, "PaymentOptionTwo")
	delete(f, "PaymentOptionThree")
	delete(f, "PaymentOptionFour")
	delete(f, "PaymentOptionFive")
	delete(f, "PaymentOptionSix")
	delete(f, "PaymentOptionSeven")
	delete(f, "PaymentOptionOther")
	delete(f, "PaymentOptionTen")
	delete(f, "PaymentOptionNine")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddContractRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddContractResponseParams struct {
	// 业务系统返回消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 业务系统返回码
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 添加合同响应对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *AddContractResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddContractResponse struct {
	*tchttp.BaseResponse
	Response *AddContractResponseParams `json:"Response"`
}

func (r *AddContractResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddContractResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddContractResult struct {
	// 合同主键
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContractId *string `json:"ContractId,omitempty" name:"ContractId"`
}

// Predefined struct for user
type AddFlexIdInfoRequestParams struct {
	// 证件类型
	// 0:身份证
	// 1:社会信用代码
	IdType *int64 `json:"IdType,omitempty" name:"IdType"`

	// 证件号
	IdNo *string `json:"IdNo,omitempty" name:"IdNo"`

	// 收款用户ID
	PayeeId *string `json:"PayeeId,omitempty" name:"PayeeId"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// __test__:测试环境
	// 缺省默认为生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type AddFlexIdInfoRequest struct {
	*tchttp.BaseRequest
	
	// 证件类型
	// 0:身份证
	// 1:社会信用代码
	IdType *int64 `json:"IdType,omitempty" name:"IdType"`

	// 证件号
	IdNo *string `json:"IdNo,omitempty" name:"IdNo"`

	// 收款用户ID
	PayeeId *string `json:"PayeeId,omitempty" name:"PayeeId"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// __test__:测试环境
	// 缺省默认为生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *AddFlexIdInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddFlexIdInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdType")
	delete(f, "IdNo")
	delete(f, "PayeeId")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddFlexIdInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddFlexIdInfoResponseParams struct {
	// 错误码。SUCCESS为成功，其他为失败
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddFlexIdInfoResponse struct {
	*tchttp.BaseResponse
	Response *AddFlexIdInfoResponseParams `json:"Response"`
}

func (r *AddFlexIdInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddFlexIdInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddFlexPhoneNoRequestParams struct {
	// 手机号
	PhoneNo *string `json:"PhoneNo,omitempty" name:"PhoneNo"`

	// 收款用户ID
	PayeeId *string `json:"PayeeId,omitempty" name:"PayeeId"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// __test__:测试环境
	// 缺省默认为生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type AddFlexPhoneNoRequest struct {
	*tchttp.BaseRequest
	
	// 手机号
	PhoneNo *string `json:"PhoneNo,omitempty" name:"PhoneNo"`

	// 收款用户ID
	PayeeId *string `json:"PayeeId,omitempty" name:"PayeeId"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// __test__:测试环境
	// 缺省默认为生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *AddFlexPhoneNoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddFlexPhoneNoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PhoneNo")
	delete(f, "PayeeId")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddFlexPhoneNoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddFlexPhoneNoResponseParams struct {
	// 错误码。SUCCESS为成功，其他为失败
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddFlexPhoneNoResponse struct {
	*tchttp.BaseResponse
	Response *AddFlexPhoneNoResponseParams `json:"Response"`
}

func (r *AddFlexPhoneNoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddFlexPhoneNoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddMerchantRequestParams struct {
	// 收单系统分配的开放ID
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 收单系统分配的密钥
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 机构商户主键（系统有唯一性校验），建议使用商户表的主键ID，防止重复添加商户
	OutMerchantId *string `json:"OutMerchantId,omitempty" name:"OutMerchantId"`

	// 商户名称，小微商户命名要符合“”商户_名字” （例如：商户_张三）
	MerchantName *string `json:"MerchantName,omitempty" name:"MerchantName"`

	// 营业执照类型（1三证合一，2非三证合一）
	BusinessLicenseType *string `json:"BusinessLicenseType,omitempty" name:"BusinessLicenseType"`

	// 营业执照编号（系统有唯一性校验），（小微商户不效验，随意传要有值，公司/个体户必传）
	BusinessLicenseNo *string `json:"BusinessLicenseNo,omitempty" name:"BusinessLicenseNo"`

	// 营业执照图片【私密区】（系统返回的图片路径），（小微商户不效验，随意传要有值，公司/个体户必传）
	BusinessLicensePicture *string `json:"BusinessLicensePicture,omitempty" name:"BusinessLicensePicture"`

	// 营业执照生效时间（yyyy-mm-dd），（小微商户不效验，随意传要有值，公司/个体户必传）
	BusinessLicenseStartDate *string `json:"BusinessLicenseStartDate,omitempty" name:"BusinessLicenseStartDate"`

	// 营业执照过期时间（yyyy-mm-dd），（小微商户不效验，随意传要有值，公司/个体户必传）
	BusinessLicenseEndDate *string `json:"BusinessLicenseEndDate,omitempty" name:"BusinessLicenseEndDate"`

	// 行业分类编号列表（第一个分类编号为主分类，后面的为二级分类）
	ClassificationIds []*string `json:"ClassificationIds,omitempty" name:"ClassificationIds"`

	// 招牌名称
	BrandName *string `json:"BrandName,omitempty" name:"BrandName"`

	// 联系电话
	Telephone *string `json:"Telephone,omitempty" name:"Telephone"`

	// 城市编号
	CityId *string `json:"CityId,omitempty" name:"CityId"`

	// 详细地址，不含省市区县名称，长度需要大于5。
	Address *string `json:"Address,omitempty" name:"Address"`

	// 营业时间，多个以小写逗号分开(9:00-12:00,13:00-18:00)
	OpenHours *string `json:"OpenHours,omitempty" name:"OpenHours"`

	// 结算账户类型（2对私，1对公）
	AccountType *string `json:"AccountType,omitempty" name:"AccountType"`

	// 清算联行号，开户行行号
	BankNo *string `json:"BankNo,omitempty" name:"BankNo"`

	// 开户行名称
	BankName *string `json:"BankName,omitempty" name:"BankName"`

	// 银行账号
	AccountNo *string `json:"AccountNo,omitempty" name:"AccountNo"`

	// 银行户名
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// 法人证件类型（1居民身份证,2临时居民身份证,3居民户口簿,4护照,5港澳居民来往内地通行证,6回乡证,7军人证,8武警身份证,9其他法定文件）
	BossIdType *string `json:"BossIdType,omitempty" name:"BossIdType"`

	// 法人证件号码
	BossIdNo *string `json:"BossIdNo,omitempty" name:"BossIdNo"`

	// 法人姓名
	BossName *string `json:"BossName,omitempty" name:"BossName"`

	// 法人性别（1男，2女）
	BossSex *string `json:"BossSex,omitempty" name:"BossSex"`

	// 法人证件国别/地区（中国CHN，香港HKG，澳门MAC，台湾CTN）
	BossIdCountry *string `json:"BossIdCountry,omitempty" name:"BossIdCountry"`

	// 法人身份证正面【私密区】（系统返回的图片路径）
	BossPositive *string `json:"BossPositive,omitempty" name:"BossPositive"`

	// 法人身份证背面【私密区】（系统返回的图片路径）
	BossBack *string `json:"BossBack,omitempty" name:"BossBack"`

	// 法人证件生效时间（yyyy-mm-dd）
	BossStartDate *string `json:"BossStartDate,omitempty" name:"BossStartDate"`

	// 法人证件过期时间（yyyy-mm-dd）
	BossEndDate *string `json:"BossEndDate,omitempty" name:"BossEndDate"`

	// 许可证图片【私密区】，（小微商户不效验，随意传要有值，公司/个体户必传）
	LicencePicture *string `json:"LicencePicture,omitempty" name:"LicencePicture"`

	// 商户类型：1-个体，2-小微，3-企业。不传默认为2-小微商户。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 组织机构代码证号
	OrganizationNo *string `json:"OrganizationNo,omitempty" name:"OrganizationNo"`

	// 组织机构代码证生效时间（yyyy-mm-dd）
	OrganizationStartDate *string `json:"OrganizationStartDate,omitempty" name:"OrganizationStartDate"`

	// 组织机构代码证图片【私密区】
	OrganizationPicture *string `json:"OrganizationPicture,omitempty" name:"OrganizationPicture"`

	// 组织机构代码证过期时间（yyyy-mm-dd）
	OrganizationEndDate *string `json:"OrganizationEndDate,omitempty" name:"OrganizationEndDate"`

	// 商户简介
	Intro *string `json:"Intro,omitempty" name:"Intro"`

	// 商户logo【公共区】
	Logo *string `json:"Logo,omitempty" name:"Logo"`

	// 商户标记，自定义参数
	Tag *string `json:"Tag,omitempty" name:"Tag"`

	// 财务联系人电话
	FinancialTelephone *string `json:"FinancialTelephone,omitempty" name:"FinancialTelephone"`

	// 财务联系人
	FinancialContact *string `json:"FinancialContact,omitempty" name:"FinancialContact"`

	// 税务登记证号
	TaxRegistrationNo *string `json:"TaxRegistrationNo,omitempty" name:"TaxRegistrationNo"`

	// 税务登记证图片【私密区】
	TaxRegistrationPicture *string `json:"TaxRegistrationPicture,omitempty" name:"TaxRegistrationPicture"`

	// 税务登记证生效时间（yyyy-mm-dd）
	TaxRegistrationStartDate *string `json:"TaxRegistrationStartDate,omitempty" name:"TaxRegistrationStartDate"`

	// 税务登记证过期时间（yyyy-mm-dd）
	TaxRegistrationEndDate *string `json:"TaxRegistrationEndDate,omitempty" name:"TaxRegistrationEndDate"`

	// 结算账户人身份（1法人，2法人亲属），结算帐户为对私时必填
	AccountBoss *string `json:"AccountBoss,omitempty" name:"AccountBoss"`

	// 客户经理姓名，必须为系统后台的管理员真实姓名
	AccountManagerName *string `json:"AccountManagerName,omitempty" name:"AccountManagerName"`

	// 法人电话
	BossTelephone *string `json:"BossTelephone,omitempty" name:"BossTelephone"`

	// 法人职业
	BossJob *string `json:"BossJob,omitempty" name:"BossJob"`

	// 法人邮箱
	BossEmail *string `json:"BossEmail,omitempty" name:"BossEmail"`

	// 法人住址
	BossAddress *string `json:"BossAddress,omitempty" name:"BossAddress"`

	// 法人亲属证件类型（1居民身份证,2临时居民身份证,3居民户口簿,4护照,5港澳居民来往内地通行证,6回乡证,7军人证,8武警身份证,9其他法定文件）结算账户人身份为法人亲属时必填
	AccountIdType *string `json:"AccountIdType,omitempty" name:"AccountIdType"`

	// 法人亲属证件号码
	AccountIdNo *string `json:"AccountIdNo,omitempty" name:"AccountIdNo"`

	// 授权文件【私密区】
	LicencePictureTwo *string `json:"LicencePictureTwo,omitempty" name:"LicencePictureTwo"`

	// 其他资料1
	OtherPictureOne *string `json:"OtherPictureOne,omitempty" name:"OtherPictureOne"`

	// 其他资料2
	OtherPictureTwo *string `json:"OtherPictureTwo,omitempty" name:"OtherPictureTwo"`

	// 其他资料3
	OtherPictureThree *string `json:"OtherPictureThree,omitempty" name:"OtherPictureThree"`

	// 其他资料4
	OtherPictureFour *string `json:"OtherPictureFour,omitempty" name:"OtherPictureFour"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

type AddMerchantRequest struct {
	*tchttp.BaseRequest
	
	// 收单系统分配的开放ID
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 收单系统分配的密钥
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 机构商户主键（系统有唯一性校验），建议使用商户表的主键ID，防止重复添加商户
	OutMerchantId *string `json:"OutMerchantId,omitempty" name:"OutMerchantId"`

	// 商户名称，小微商户命名要符合“”商户_名字” （例如：商户_张三）
	MerchantName *string `json:"MerchantName,omitempty" name:"MerchantName"`

	// 营业执照类型（1三证合一，2非三证合一）
	BusinessLicenseType *string `json:"BusinessLicenseType,omitempty" name:"BusinessLicenseType"`

	// 营业执照编号（系统有唯一性校验），（小微商户不效验，随意传要有值，公司/个体户必传）
	BusinessLicenseNo *string `json:"BusinessLicenseNo,omitempty" name:"BusinessLicenseNo"`

	// 营业执照图片【私密区】（系统返回的图片路径），（小微商户不效验，随意传要有值，公司/个体户必传）
	BusinessLicensePicture *string `json:"BusinessLicensePicture,omitempty" name:"BusinessLicensePicture"`

	// 营业执照生效时间（yyyy-mm-dd），（小微商户不效验，随意传要有值，公司/个体户必传）
	BusinessLicenseStartDate *string `json:"BusinessLicenseStartDate,omitempty" name:"BusinessLicenseStartDate"`

	// 营业执照过期时间（yyyy-mm-dd），（小微商户不效验，随意传要有值，公司/个体户必传）
	BusinessLicenseEndDate *string `json:"BusinessLicenseEndDate,omitempty" name:"BusinessLicenseEndDate"`

	// 行业分类编号列表（第一个分类编号为主分类，后面的为二级分类）
	ClassificationIds []*string `json:"ClassificationIds,omitempty" name:"ClassificationIds"`

	// 招牌名称
	BrandName *string `json:"BrandName,omitempty" name:"BrandName"`

	// 联系电话
	Telephone *string `json:"Telephone,omitempty" name:"Telephone"`

	// 城市编号
	CityId *string `json:"CityId,omitempty" name:"CityId"`

	// 详细地址，不含省市区县名称，长度需要大于5。
	Address *string `json:"Address,omitempty" name:"Address"`

	// 营业时间，多个以小写逗号分开(9:00-12:00,13:00-18:00)
	OpenHours *string `json:"OpenHours,omitempty" name:"OpenHours"`

	// 结算账户类型（2对私，1对公）
	AccountType *string `json:"AccountType,omitempty" name:"AccountType"`

	// 清算联行号，开户行行号
	BankNo *string `json:"BankNo,omitempty" name:"BankNo"`

	// 开户行名称
	BankName *string `json:"BankName,omitempty" name:"BankName"`

	// 银行账号
	AccountNo *string `json:"AccountNo,omitempty" name:"AccountNo"`

	// 银行户名
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// 法人证件类型（1居民身份证,2临时居民身份证,3居民户口簿,4护照,5港澳居民来往内地通行证,6回乡证,7军人证,8武警身份证,9其他法定文件）
	BossIdType *string `json:"BossIdType,omitempty" name:"BossIdType"`

	// 法人证件号码
	BossIdNo *string `json:"BossIdNo,omitempty" name:"BossIdNo"`

	// 法人姓名
	BossName *string `json:"BossName,omitempty" name:"BossName"`

	// 法人性别（1男，2女）
	BossSex *string `json:"BossSex,omitempty" name:"BossSex"`

	// 法人证件国别/地区（中国CHN，香港HKG，澳门MAC，台湾CTN）
	BossIdCountry *string `json:"BossIdCountry,omitempty" name:"BossIdCountry"`

	// 法人身份证正面【私密区】（系统返回的图片路径）
	BossPositive *string `json:"BossPositive,omitempty" name:"BossPositive"`

	// 法人身份证背面【私密区】（系统返回的图片路径）
	BossBack *string `json:"BossBack,omitempty" name:"BossBack"`

	// 法人证件生效时间（yyyy-mm-dd）
	BossStartDate *string `json:"BossStartDate,omitempty" name:"BossStartDate"`

	// 法人证件过期时间（yyyy-mm-dd）
	BossEndDate *string `json:"BossEndDate,omitempty" name:"BossEndDate"`

	// 许可证图片【私密区】，（小微商户不效验，随意传要有值，公司/个体户必传）
	LicencePicture *string `json:"LicencePicture,omitempty" name:"LicencePicture"`

	// 商户类型：1-个体，2-小微，3-企业。不传默认为2-小微商户。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 组织机构代码证号
	OrganizationNo *string `json:"OrganizationNo,omitempty" name:"OrganizationNo"`

	// 组织机构代码证生效时间（yyyy-mm-dd）
	OrganizationStartDate *string `json:"OrganizationStartDate,omitempty" name:"OrganizationStartDate"`

	// 组织机构代码证图片【私密区】
	OrganizationPicture *string `json:"OrganizationPicture,omitempty" name:"OrganizationPicture"`

	// 组织机构代码证过期时间（yyyy-mm-dd）
	OrganizationEndDate *string `json:"OrganizationEndDate,omitempty" name:"OrganizationEndDate"`

	// 商户简介
	Intro *string `json:"Intro,omitempty" name:"Intro"`

	// 商户logo【公共区】
	Logo *string `json:"Logo,omitempty" name:"Logo"`

	// 商户标记，自定义参数
	Tag *string `json:"Tag,omitempty" name:"Tag"`

	// 财务联系人电话
	FinancialTelephone *string `json:"FinancialTelephone,omitempty" name:"FinancialTelephone"`

	// 财务联系人
	FinancialContact *string `json:"FinancialContact,omitempty" name:"FinancialContact"`

	// 税务登记证号
	TaxRegistrationNo *string `json:"TaxRegistrationNo,omitempty" name:"TaxRegistrationNo"`

	// 税务登记证图片【私密区】
	TaxRegistrationPicture *string `json:"TaxRegistrationPicture,omitempty" name:"TaxRegistrationPicture"`

	// 税务登记证生效时间（yyyy-mm-dd）
	TaxRegistrationStartDate *string `json:"TaxRegistrationStartDate,omitempty" name:"TaxRegistrationStartDate"`

	// 税务登记证过期时间（yyyy-mm-dd）
	TaxRegistrationEndDate *string `json:"TaxRegistrationEndDate,omitempty" name:"TaxRegistrationEndDate"`

	// 结算账户人身份（1法人，2法人亲属），结算帐户为对私时必填
	AccountBoss *string `json:"AccountBoss,omitempty" name:"AccountBoss"`

	// 客户经理姓名，必须为系统后台的管理员真实姓名
	AccountManagerName *string `json:"AccountManagerName,omitempty" name:"AccountManagerName"`

	// 法人电话
	BossTelephone *string `json:"BossTelephone,omitempty" name:"BossTelephone"`

	// 法人职业
	BossJob *string `json:"BossJob,omitempty" name:"BossJob"`

	// 法人邮箱
	BossEmail *string `json:"BossEmail,omitempty" name:"BossEmail"`

	// 法人住址
	BossAddress *string `json:"BossAddress,omitempty" name:"BossAddress"`

	// 法人亲属证件类型（1居民身份证,2临时居民身份证,3居民户口簿,4护照,5港澳居民来往内地通行证,6回乡证,7军人证,8武警身份证,9其他法定文件）结算账户人身份为法人亲属时必填
	AccountIdType *string `json:"AccountIdType,omitempty" name:"AccountIdType"`

	// 法人亲属证件号码
	AccountIdNo *string `json:"AccountIdNo,omitempty" name:"AccountIdNo"`

	// 授权文件【私密区】
	LicencePictureTwo *string `json:"LicencePictureTwo,omitempty" name:"LicencePictureTwo"`

	// 其他资料1
	OtherPictureOne *string `json:"OtherPictureOne,omitempty" name:"OtherPictureOne"`

	// 其他资料2
	OtherPictureTwo *string `json:"OtherPictureTwo,omitempty" name:"OtherPictureTwo"`

	// 其他资料3
	OtherPictureThree *string `json:"OtherPictureThree,omitempty" name:"OtherPictureThree"`

	// 其他资料4
	OtherPictureFour *string `json:"OtherPictureFour,omitempty" name:"OtherPictureFour"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *AddMerchantRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddMerchantRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OpenId")
	delete(f, "OpenKey")
	delete(f, "OutMerchantId")
	delete(f, "MerchantName")
	delete(f, "BusinessLicenseType")
	delete(f, "BusinessLicenseNo")
	delete(f, "BusinessLicensePicture")
	delete(f, "BusinessLicenseStartDate")
	delete(f, "BusinessLicenseEndDate")
	delete(f, "ClassificationIds")
	delete(f, "BrandName")
	delete(f, "Telephone")
	delete(f, "CityId")
	delete(f, "Address")
	delete(f, "OpenHours")
	delete(f, "AccountType")
	delete(f, "BankNo")
	delete(f, "BankName")
	delete(f, "AccountNo")
	delete(f, "AccountName")
	delete(f, "BossIdType")
	delete(f, "BossIdNo")
	delete(f, "BossName")
	delete(f, "BossSex")
	delete(f, "BossIdCountry")
	delete(f, "BossPositive")
	delete(f, "BossBack")
	delete(f, "BossStartDate")
	delete(f, "BossEndDate")
	delete(f, "LicencePicture")
	delete(f, "Type")
	delete(f, "OrganizationNo")
	delete(f, "OrganizationStartDate")
	delete(f, "OrganizationPicture")
	delete(f, "OrganizationEndDate")
	delete(f, "Intro")
	delete(f, "Logo")
	delete(f, "Tag")
	delete(f, "FinancialTelephone")
	delete(f, "FinancialContact")
	delete(f, "TaxRegistrationNo")
	delete(f, "TaxRegistrationPicture")
	delete(f, "TaxRegistrationStartDate")
	delete(f, "TaxRegistrationEndDate")
	delete(f, "AccountBoss")
	delete(f, "AccountManagerName")
	delete(f, "BossTelephone")
	delete(f, "BossJob")
	delete(f, "BossEmail")
	delete(f, "BossAddress")
	delete(f, "AccountIdType")
	delete(f, "AccountIdNo")
	delete(f, "LicencePictureTwo")
	delete(f, "OtherPictureOne")
	delete(f, "OtherPictureTwo")
	delete(f, "OtherPictureThree")
	delete(f, "OtherPictureFour")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddMerchantRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddMerchantResponseParams struct {
	// 业务系统返回消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 业务系统返回码，0表示成功，其他表示失败。
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 添加商户响应对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *AddMerchantResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddMerchantResponse struct {
	*tchttp.BaseResponse
	Response *AddMerchantResponseParams `json:"Response"`
}

func (r *AddMerchantResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddMerchantResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddMerchantResult struct {
	// 系统商户号
	// 注意：此字段可能返回 null，表示取不到有效值。
	MerchantNo *string `json:"MerchantNo,omitempty" name:"MerchantNo"`
}

// Predefined struct for user
type AddShopRequestParams struct {
	// 收单系统分配的开放ID
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 收单系统分配的密钥
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 机构门店主键（系统有唯一性校验），建议使用门店表的主键ID，防止重复添加门店
	OutShopId *string `json:"OutShopId,omitempty" name:"OutShopId"`

	// 门店简称（例如：南山店）
	ShopName *string `json:"ShopName,omitempty" name:"ShopName"`

	// 门店全称（例如：江山小厨（南山店））
	ShopFullName *string `json:"ShopFullName,omitempty" name:"ShopFullName"`

	// 商户编号
	MerchantNo *string `json:"MerchantNo,omitempty" name:"MerchantNo"`

	// 门店电话
	Telephone *string `json:"Telephone,omitempty" name:"Telephone"`

	// 营业时间，多个以小写逗号分开(9:00-12:00,13:00-18:00)
	OpenHours *string `json:"OpenHours,omitempty" name:"OpenHours"`

	// 门店所在的城市编码
	CityId *string `json:"CityId,omitempty" name:"CityId"`

	// 门店详细地址，不含省市区县名称
	Address *string `json:"Address,omitempty" name:"Address"`

	// 整体门面（含招牌）图片【公共区】
	PictureOne *string `json:"PictureOne,omitempty" name:"PictureOne"`

	// 整体门面（含招牌）图片【公共区】
	PictureTwo *string `json:"PictureTwo,omitempty" name:"PictureTwo"`

	// 店内环境图片【公共区】
	PictureThree *string `json:"PictureThree,omitempty" name:"PictureThree"`

	// 负责人手机号码
	FinancialTelephone *string `json:"FinancialTelephone,omitempty" name:"FinancialTelephone"`

	// 门店负责人
	Contact *string `json:"Contact,omitempty" name:"Contact"`

	// 百度地图纬度
	Latitude *string `json:"Latitude,omitempty" name:"Latitude"`

	// 高德地图纬度
	LatitudeTwo *string `json:"LatitudeTwo,omitempty" name:"LatitudeTwo"`

	// 百度地图经度
	Longitude *string `json:"Longitude,omitempty" name:"Longitude"`

	// 高德地图经度
	LongitudeTwo *string `json:"LongitudeTwo,omitempty" name:"LongitudeTwo"`

	// 其他照片【公共区】
	OtherPicture *string `json:"OtherPicture,omitempty" name:"OtherPicture"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

type AddShopRequest struct {
	*tchttp.BaseRequest
	
	// 收单系统分配的开放ID
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 收单系统分配的密钥
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 机构门店主键（系统有唯一性校验），建议使用门店表的主键ID，防止重复添加门店
	OutShopId *string `json:"OutShopId,omitempty" name:"OutShopId"`

	// 门店简称（例如：南山店）
	ShopName *string `json:"ShopName,omitempty" name:"ShopName"`

	// 门店全称（例如：江山小厨（南山店））
	ShopFullName *string `json:"ShopFullName,omitempty" name:"ShopFullName"`

	// 商户编号
	MerchantNo *string `json:"MerchantNo,omitempty" name:"MerchantNo"`

	// 门店电话
	Telephone *string `json:"Telephone,omitempty" name:"Telephone"`

	// 营业时间，多个以小写逗号分开(9:00-12:00,13:00-18:00)
	OpenHours *string `json:"OpenHours,omitempty" name:"OpenHours"`

	// 门店所在的城市编码
	CityId *string `json:"CityId,omitempty" name:"CityId"`

	// 门店详细地址，不含省市区县名称
	Address *string `json:"Address,omitempty" name:"Address"`

	// 整体门面（含招牌）图片【公共区】
	PictureOne *string `json:"PictureOne,omitempty" name:"PictureOne"`

	// 整体门面（含招牌）图片【公共区】
	PictureTwo *string `json:"PictureTwo,omitempty" name:"PictureTwo"`

	// 店内环境图片【公共区】
	PictureThree *string `json:"PictureThree,omitempty" name:"PictureThree"`

	// 负责人手机号码
	FinancialTelephone *string `json:"FinancialTelephone,omitempty" name:"FinancialTelephone"`

	// 门店负责人
	Contact *string `json:"Contact,omitempty" name:"Contact"`

	// 百度地图纬度
	Latitude *string `json:"Latitude,omitempty" name:"Latitude"`

	// 高德地图纬度
	LatitudeTwo *string `json:"LatitudeTwo,omitempty" name:"LatitudeTwo"`

	// 百度地图经度
	Longitude *string `json:"Longitude,omitempty" name:"Longitude"`

	// 高德地图经度
	LongitudeTwo *string `json:"LongitudeTwo,omitempty" name:"LongitudeTwo"`

	// 其他照片【公共区】
	OtherPicture *string `json:"OtherPicture,omitempty" name:"OtherPicture"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *AddShopRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddShopRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OpenId")
	delete(f, "OpenKey")
	delete(f, "OutShopId")
	delete(f, "ShopName")
	delete(f, "ShopFullName")
	delete(f, "MerchantNo")
	delete(f, "Telephone")
	delete(f, "OpenHours")
	delete(f, "CityId")
	delete(f, "Address")
	delete(f, "PictureOne")
	delete(f, "PictureTwo")
	delete(f, "PictureThree")
	delete(f, "FinancialTelephone")
	delete(f, "Contact")
	delete(f, "Latitude")
	delete(f, "LatitudeTwo")
	delete(f, "Longitude")
	delete(f, "LongitudeTwo")
	delete(f, "OtherPicture")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddShopRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddShopResponseParams struct {
	// 业务系统返回消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 业务系统返回码
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 添加申请响应对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *AddShopResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddShopResponse struct {
	*tchttp.BaseResponse
	Response *AddShopResponseParams `json:"Response"`
}

func (r *AddShopResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddShopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddShopResult struct {
	// 门店编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShopNo *string `json:"ShopNo,omitempty" name:"ShopNo"`
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

type AmountBeforeTaxResult struct {
	// 税前金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	AmountBeforeTax *string `json:"AmountBeforeTax,omitempty" name:"AmountBeforeTax"`
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

type AnchorExtendInfo struct {
	// 扩展信息类型
	// __id_card_no__:身份证号码
	// __id_card_name__:身份证姓名
	// __id_card_front__:身份证图片正面
	// __id_card_back__:身份证图片反面
	// __tax_type__:完税类型:0-自然人,1-个体工商户
	// __channel_account__:渠道账号(_敏感信息_ 使用 __AES128-CBC-PKCS#7__ 加密)
	Type *string `json:"Type,omitempty" name:"Type"`

	// 扩展信息
	Value *string `json:"Value,omitempty" name:"Value"`
}

// Predefined struct for user
type ApplyApplicationMaterialRequestParams struct {
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

// Predefined struct for user
type ApplyApplicationMaterialResponseParams struct {
	// 提交申报材料结果
	Result *ApplyDeclareResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ApplyApplicationMaterialResponse struct {
	*tchttp.BaseResponse
	Response *ApplyApplicationMaterialResponseParams `json:"Response"`
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

// Predefined struct for user
type ApplyFlexPaymentRequestParams struct {
	// 收款用户ID
	PayeeId *string `json:"PayeeId,omitempty" name:"PayeeId"`

	// 收入类型
	// LABOR:劳务所得
	// OCCASION:偶然所得
	IncomeType *string `json:"IncomeType,omitempty" name:"IncomeType"`

	// 税前金额
	AmountBeforeTax *string `json:"AmountBeforeTax,omitempty" name:"AmountBeforeTax"`

	// 外部订单ID
	OutOrderId *string `json:"OutOrderId,omitempty" name:"OutOrderId"`

	// 资金账户信息
	FundingAccountInfo *FlexFundingAccountInfo `json:"FundingAccountInfo,omitempty" name:"FundingAccountInfo"`

	// 提现备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// __test__:测试环境
	// 缺省默认为生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type ApplyFlexPaymentRequest struct {
	*tchttp.BaseRequest
	
	// 收款用户ID
	PayeeId *string `json:"PayeeId,omitempty" name:"PayeeId"`

	// 收入类型
	// LABOR:劳务所得
	// OCCASION:偶然所得
	IncomeType *string `json:"IncomeType,omitempty" name:"IncomeType"`

	// 税前金额
	AmountBeforeTax *string `json:"AmountBeforeTax,omitempty" name:"AmountBeforeTax"`

	// 外部订单ID
	OutOrderId *string `json:"OutOrderId,omitempty" name:"OutOrderId"`

	// 资金账户信息
	FundingAccountInfo *FlexFundingAccountInfo `json:"FundingAccountInfo,omitempty" name:"FundingAccountInfo"`

	// 提现备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// __test__:测试环境
	// 缺省默认为生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *ApplyFlexPaymentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyFlexPaymentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PayeeId")
	delete(f, "IncomeType")
	delete(f, "AmountBeforeTax")
	delete(f, "OutOrderId")
	delete(f, "FundingAccountInfo")
	delete(f, "Remark")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyFlexPaymentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyFlexPaymentResponseParams struct {
	// 错误码。SUCCESS为成功，其他为失败
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *ApplyFlexPaymentResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ApplyFlexPaymentResponse struct {
	*tchttp.BaseResponse
	Response *ApplyFlexPaymentResponseParams `json:"Response"`
}

func (r *ApplyFlexPaymentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyFlexPaymentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplyFlexPaymentResult struct {
	// 订单ID
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 税前金额
	AmountBeforeTax *string `json:"AmountBeforeTax,omitempty" name:"AmountBeforeTax"`

	// 税后金额
	AmountAfterTax *string `json:"AmountAfterTax,omitempty" name:"AmountAfterTax"`

	// 税金
	Tax *string `json:"Tax,omitempty" name:"Tax"`
}

// Predefined struct for user
type ApplyFlexSettlementRequestParams struct {
	// 收款用户ID
	PayeeId *string `json:"PayeeId,omitempty" name:"PayeeId"`

	// 收入类型
	// LABOR:劳务所得
	// OCCASION:偶然所得
	IncomeType *string `json:"IncomeType,omitempty" name:"IncomeType"`

	// 税前金额
	AmountBeforeTax *string `json:"AmountBeforeTax,omitempty" name:"AmountBeforeTax"`

	// 外部订单ID
	OutOrderId *string `json:"OutOrderId,omitempty" name:"OutOrderId"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// __test__:测试环境
	// 缺省默认为生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type ApplyFlexSettlementRequest struct {
	*tchttp.BaseRequest
	
	// 收款用户ID
	PayeeId *string `json:"PayeeId,omitempty" name:"PayeeId"`

	// 收入类型
	// LABOR:劳务所得
	// OCCASION:偶然所得
	IncomeType *string `json:"IncomeType,omitempty" name:"IncomeType"`

	// 税前金额
	AmountBeforeTax *string `json:"AmountBeforeTax,omitempty" name:"AmountBeforeTax"`

	// 外部订单ID
	OutOrderId *string `json:"OutOrderId,omitempty" name:"OutOrderId"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// __test__:测试环境
	// 缺省默认为生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *ApplyFlexSettlementRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyFlexSettlementRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PayeeId")
	delete(f, "IncomeType")
	delete(f, "AmountBeforeTax")
	delete(f, "OutOrderId")
	delete(f, "Remark")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyFlexSettlementRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyFlexSettlementResponseParams struct {
	// 错误码。SUCCESS为成功，其他为失败
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *ApplyFlexSettlementResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ApplyFlexSettlementResponse struct {
	*tchttp.BaseResponse
	Response *ApplyFlexSettlementResponseParams `json:"Response"`
}

func (r *ApplyFlexSettlementResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyFlexSettlementResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplyFlexSettlementResult struct {
	// 订单ID
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 税前金额
	AmountBeforeTax *string `json:"AmountBeforeTax,omitempty" name:"AmountBeforeTax"`

	// 税后金额
	AmountAfterTax *string `json:"AmountAfterTax,omitempty" name:"AmountAfterTax"`

	// 税金
	Tax *string `json:"Tax,omitempty" name:"Tax"`
}

// Predefined struct for user
type ApplyOpenBankOrderDetailReceiptRequestParams struct {
	// 外部回单申请编号
	OutApplyId *string `json:"OutApplyId,omitempty" name:"OutApplyId"`

	// 渠道商户ID
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 渠道子商户ID
	ChannelSubMerchantId *string `json:"ChannelSubMerchantId,omitempty" name:"ChannelSubMerchantId"`

	// 渠道名称，目前只支持ALIPAY
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 支付方式，目前只支持SAFT_ISV
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`

	// 云企付平台订单号
	ChannelOrderId *string `json:"ChannelOrderId,omitempty" name:"ChannelOrderId"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type ApplyOpenBankOrderDetailReceiptRequest struct {
	*tchttp.BaseRequest
	
	// 外部回单申请编号
	OutApplyId *string `json:"OutApplyId,omitempty" name:"OutApplyId"`

	// 渠道商户ID
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 渠道子商户ID
	ChannelSubMerchantId *string `json:"ChannelSubMerchantId,omitempty" name:"ChannelSubMerchantId"`

	// 渠道名称，目前只支持ALIPAY
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 支付方式，目前只支持SAFT_ISV
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`

	// 云企付平台订单号
	ChannelOrderId *string `json:"ChannelOrderId,omitempty" name:"ChannelOrderId"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *ApplyOpenBankOrderDetailReceiptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyOpenBankOrderDetailReceiptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OutApplyId")
	delete(f, "ChannelMerchantId")
	delete(f, "ChannelSubMerchantId")
	delete(f, "ChannelName")
	delete(f, "PaymentMethod")
	delete(f, "ChannelOrderId")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyOpenBankOrderDetailReceiptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyOpenBankOrderDetailReceiptResponseParams struct {
	// 错误码。
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *ApplyOpenBankOrderDetailReceiptResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ApplyOpenBankOrderDetailReceiptResponse struct {
	*tchttp.BaseResponse
	Response *ApplyOpenBankOrderDetailReceiptResponseParams `json:"Response"`
}

func (r *ApplyOpenBankOrderDetailReceiptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyOpenBankOrderDetailReceiptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplyOpenBankOrderDetailReceiptResult struct {
	// 渠道回单申请ID
	ChannelApplyId *string `json:"ChannelApplyId,omitempty" name:"ChannelApplyId"`

	// 申请状态。
	// SUCCESS：申请成功；
	// FAILED：申请失败；
	// PROCESSING：申请中。
	// 注意：若返回申请中，需要再次调用回单申请结果查询接口，查询结果。
	ReceiptStatus *string `json:"ReceiptStatus,omitempty" name:"ReceiptStatus"`

	// 申请返回描述，例如失败原因等。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReceiptMessage *string `json:"ReceiptMessage,omitempty" name:"ReceiptMessage"`

	// 回单下载链接，申请成功时返回。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 过期时间，yyyy-MM-dd HH:mm:ss格式。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

type ApplyOutwardOrderData struct {
	// 商户号
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 对接方汇出指令编号
	TransactionId *string `json:"TransactionId,omitempty" name:"TransactionId"`

	// 受理状态
	Status *string `json:"Status,omitempty" name:"Status"`
}

// Predefined struct for user
type ApplyOutwardOrderRequestParams struct {
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

// Predefined struct for user
type ApplyOutwardOrderResponseParams struct {
	// 汇出指令申请
	Result *ApplyOutwardOrderResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ApplyOutwardOrderResponse struct {
	*tchttp.BaseResponse
	Response *ApplyOutwardOrderResponseParams `json:"Response"`
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

// Predefined struct for user
type ApplyPayerInfoRequestParams struct {
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

// Predefined struct for user
type ApplyPayerInfoResponseParams struct {
	// 付款人申请结果
	Result *ApplyPayerinfoResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ApplyPayerInfoResponse struct {
	*tchttp.BaseResponse
	Response *ApplyPayerInfoResponseParams `json:"Response"`
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

// Predefined struct for user
type ApplyReWithdrawalRequestParams struct {
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

// Predefined struct for user
type ApplyReWithdrawalResponseParams struct {
	// 重新提现业务订单号
	WithdrawOrderId *string `json:"WithdrawOrderId,omitempty" name:"WithdrawOrderId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ApplyReWithdrawalResponse struct {
	*tchttp.BaseResponse
	Response *ApplyReWithdrawalResponseParams `json:"Response"`
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

// Predefined struct for user
type ApplyReconciliationFileRequestParams struct {
	// 申请的文件类型。
	// __CZ__：充值文件
	// __TX__：提现文件
	// __JY__：交易文件
	// __YE__：余额文件
	ApplyFileType *string `json:"ApplyFileType,omitempty" name:"ApplyFileType"`

	// 申请的对账文件日期，格式：yyyyMMdd。
	ApplyFileDate *string `json:"ApplyFileDate,omitempty" name:"ApplyFileDate"`

	// 父账户账号。
	// _平安渠道为资金汇总账号_
	BankAccountNumber *string `json:"BankAccountNumber,omitempty" name:"BankAccountNumber"`

	// 环境名。
	// __release__: 现网环境
	// __sandbox__: 沙箱环境
	// __development__: 开发环境
	// _缺省: release_
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`
}

type ApplyReconciliationFileRequest struct {
	*tchttp.BaseRequest
	
	// 申请的文件类型。
	// __CZ__：充值文件
	// __TX__：提现文件
	// __JY__：交易文件
	// __YE__：余额文件
	ApplyFileType *string `json:"ApplyFileType,omitempty" name:"ApplyFileType"`

	// 申请的对账文件日期，格式：yyyyMMdd。
	ApplyFileDate *string `json:"ApplyFileDate,omitempty" name:"ApplyFileDate"`

	// 父账户账号。
	// _平安渠道为资金汇总账号_
	BankAccountNumber *string `json:"BankAccountNumber,omitempty" name:"BankAccountNumber"`

	// 环境名。
	// __release__: 现网环境
	// __sandbox__: 沙箱环境
	// __development__: 开发环境
	// _缺省: release_
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`
}

func (r *ApplyReconciliationFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyReconciliationFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplyFileType")
	delete(f, "ApplyFileDate")
	delete(f, "BankAccountNumber")
	delete(f, "MidasEnvironment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyReconciliationFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyReconciliationFileResponseParams struct {
	// 错误码。
	// __SUCCESS__: 成功
	// __其他__: 见附录-错误码表
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *ApplyReconciliationFileResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ApplyReconciliationFileResponse struct {
	*tchttp.BaseResponse
	Response *ApplyReconciliationFileResponseParams `json:"Response"`
}

func (r *ApplyReconciliationFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyReconciliationFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplyReconciliationFileResult struct {
	// 申请对账文件的任务ID。
	ApplyFileId *string `json:"ApplyFileId,omitempty" name:"ApplyFileId"`

	// 对账文件申请状态。
	// __I__：申请中
	// __S__：申请成功
	// __F__：申请失败
	ApplyStatus *string `json:"ApplyStatus,omitempty" name:"ApplyStatus"`

	// 申请结果描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplyMessage *string `json:"ApplyMessage,omitempty" name:"ApplyMessage"`
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

// Predefined struct for user
type ApplyTradeRequestParams struct {
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

// Predefined struct for user
type ApplyTradeResponseParams struct {
	// 提交贸易材料结果
	Result *ApplyTradeResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ApplyTradeResponse struct {
	*tchttp.BaseResponse
	Response *ApplyTradeResponseParams `json:"Response"`
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

// Predefined struct for user
type ApplyWithdrawalRequestParams struct {
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

// Predefined struct for user
type ApplyWithdrawalResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ApplyWithdrawalResponse struct {
	*tchttp.BaseResponse
	Response *ApplyWithdrawalResponseParams `json:"Response"`
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

type AssignmentData struct {
	// 主播ID
	AnchorId *string `json:"AnchorId,omitempty" name:"AnchorId"`

	// 主播名称
	AnchorName *string `json:"AnchorName,omitempty" name:"AnchorName"`

	// 代理商ID
	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`

	// 代理商名称
	AgentName *string `json:"AgentName,omitempty" name:"AgentName"`
}

type BankBranchInfo struct {
	// 银行名称。
	BankName *string `json:"BankName,omitempty" name:"BankName"`

	// 银行简称。
	BankAbbreviation *string `json:"BankAbbreviation,omitempty" name:"BankAbbreviation"`

	// 支行名。
	BankBranchName *string `json:"BankBranchName,omitempty" name:"BankBranchName"`

	// 联行号。
	BankBranchId *string `json:"BankBranchId,omitempty" name:"BankBranchId"`
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

type BillDownloadUrlResult struct {
	// 对账单下载地址。GET方式访问，返回zip包，解压后为csv格式文件。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
}

// Predefined struct for user
type BindAccountRequestParams struct {
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

// Predefined struct for user
type BindAccountResponseParams struct {
	// 错误码。响应成功："SUCCESS"，其他为不成功。
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 响应消息。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 该字段为null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindAccountResponse struct {
	*tchttp.BaseResponse
	Response *BindAccountResponseParams `json:"Response"`
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

// Predefined struct for user
type BindAcctRequestParams struct {
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

// Predefined struct for user
type BindAcctResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindAcctResponse struct {
	*tchttp.BaseResponse
	Response *BindAcctResponseParams `json:"Response"`
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

// Predefined struct for user
type BindOpenBankExternalSubMerchantBankAccountRequestParams struct {
	// 渠道商户ID。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 渠道子商户ID。
	ChannelSubMerchantId *string `json:"ChannelSubMerchantId,omitempty" name:"ChannelSubMerchantId"`

	// 渠道名称。
	// __TENPAY__: 商企付
	// __WECHAT__: 微信支付
	// __ALIPAY__: 支付宝
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 支付方式。
	// __EBANK_PAYMENT__: ebank支付
	// __OPENBANK_PAYMENT__: openbank支付
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`

	// 第三方渠道子商户收款方银行卡信息, 为JSON格式字符串。详情见附录-复杂类型。
	ExternalSubMerchantBindBankAccountData *string `json:"ExternalSubMerchantBindBankAccountData,omitempty" name:"ExternalSubMerchantBindBankAccountData"`

	// 外部申请编号。
	OutApplyId *string `json:"OutApplyId,omitempty" name:"OutApplyId"`

	// 通知地址。
	NotifyUrl *string `json:"NotifyUrl,omitempty" name:"NotifyUrl"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type BindOpenBankExternalSubMerchantBankAccountRequest struct {
	*tchttp.BaseRequest
	
	// 渠道商户ID。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 渠道子商户ID。
	ChannelSubMerchantId *string `json:"ChannelSubMerchantId,omitempty" name:"ChannelSubMerchantId"`

	// 渠道名称。
	// __TENPAY__: 商企付
	// __WECHAT__: 微信支付
	// __ALIPAY__: 支付宝
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 支付方式。
	// __EBANK_PAYMENT__: ebank支付
	// __OPENBANK_PAYMENT__: openbank支付
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`

	// 第三方渠道子商户收款方银行卡信息, 为JSON格式字符串。详情见附录-复杂类型。
	ExternalSubMerchantBindBankAccountData *string `json:"ExternalSubMerchantBindBankAccountData,omitempty" name:"ExternalSubMerchantBindBankAccountData"`

	// 外部申请编号。
	OutApplyId *string `json:"OutApplyId,omitempty" name:"OutApplyId"`

	// 通知地址。
	NotifyUrl *string `json:"NotifyUrl,omitempty" name:"NotifyUrl"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *BindOpenBankExternalSubMerchantBankAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindOpenBankExternalSubMerchantBankAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelMerchantId")
	delete(f, "ChannelSubMerchantId")
	delete(f, "ChannelName")
	delete(f, "PaymentMethod")
	delete(f, "ExternalSubMerchantBindBankAccountData")
	delete(f, "OutApplyId")
	delete(f, "NotifyUrl")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindOpenBankExternalSubMerchantBankAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindOpenBankExternalSubMerchantBankAccountResponseParams struct {
	// 错误码。
	// __SUCCESS__: 成功
	// __其他__: 见附录-错误码表
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *BindOpenBankExternalSubMerchantBankAccountResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindOpenBankExternalSubMerchantBankAccountResponse struct {
	*tchttp.BaseResponse
	Response *BindOpenBankExternalSubMerchantBankAccountResponseParams `json:"Response"`
}

func (r *BindOpenBankExternalSubMerchantBankAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindOpenBankExternalSubMerchantBankAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindOpenBankExternalSubMerchantBankAccountResult struct {
	// 渠道申请编号。
	ChannelApplyId *string `json:"ChannelApplyId,omitempty" name:"ChannelApplyId"`

	// 绑定状态。
	// __SUCCESS__: 绑定成功
	// __FAILED__: 绑定失败
	// __PROCESSING__: 绑定中。
	// 注意：若返回绑定中，需要再次调用绑定结果查询接口,查询结果。
	BindStatus *string `json:"BindStatus,omitempty" name:"BindStatus"`

	// 绑定返回描述, 例如失败原因等。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindMessage *string `json:"BindMessage,omitempty" name:"BindMessage"`

	// 渠道子商户银行账户信息, 为JSON格式字符串（绑定成功状态下返回）。详情见附录-复杂类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalSubMerchantBankAccountReturnData *string `json:"ExternalSubMerchantBankAccountReturnData,omitempty" name:"ExternalSubMerchantBankAccountReturnData"`

	// 绑卡序列号。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindSerialNo *string `json:"BindSerialNo,omitempty" name:"BindSerialNo"`
}

// Predefined struct for user
type BindRelateAccReUnionPayRequestParams struct {
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

// Predefined struct for user
type BindRelateAccReUnionPayResponseParams struct {
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
}

type BindRelateAccReUnionPayResponse struct {
	*tchttp.BaseResponse
	Response *BindRelateAccReUnionPayResponseParams `json:"Response"`
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

// Predefined struct for user
type BindRelateAcctSmallAmountRequestParams struct {
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

// Predefined struct for user
type BindRelateAcctSmallAmountResponseParams struct {
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
}

type BindRelateAcctSmallAmountResponse struct {
	*tchttp.BaseResponse
	Response *BindRelateAcctSmallAmountResponseParams `json:"Response"`
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

// Predefined struct for user
type BindRelateAcctUnionPayRequestParams struct {
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

// Predefined struct for user
type BindRelateAcctUnionPayResponseParams struct {
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
}

type BindRelateAcctUnionPayResponse struct {
	*tchttp.BaseResponse
	Response *BindRelateAcctUnionPayResponseParams `json:"Response"`
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

type BusinessLicenseInfo struct {
	// 营业证件类型
	//  IDCARD：身份证
	//  CREDITCODE：统一社会信用代码
	BusinessLicenseType *string `json:"BusinessLicenseType,omitempty" name:"BusinessLicenseType"`

	// 营业证件号码 非个人商户上送统一社会信用代码，个人商户上送身份证号码
	BusinessLicenseNumber *string `json:"BusinessLicenseNumber,omitempty" name:"BusinessLicenseNumber"`

	// 营业证件有效期类型 
	// LONGTERM：长期有效
	// OTHER：非长期有效
	BusinessLicenseValidityType *string `json:"BusinessLicenseValidityType,omitempty" name:"BusinessLicenseValidityType"`

	// 营业证件生效日期，yyyy-MM-dd
	BusinessLicenseEffectiveDate *string `json:"BusinessLicenseEffectiveDate,omitempty" name:"BusinessLicenseEffectiveDate"`

	// 营业证件失效日期，yyyy-MM-dd
	BusinessLicenseExpireDate *string `json:"BusinessLicenseExpireDate,omitempty" name:"BusinessLicenseExpireDate"`
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

// Predefined struct for user
type CheckAcctRequestParams struct {
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

// Predefined struct for user
type CheckAcctResponseParams struct {
	// 前置流水号，请保存
	FrontSeqNo *string `json:"FrontSeqNo,omitempty" name:"FrontSeqNo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CheckAcctResponse struct {
	*tchttp.BaseResponse
	Response *CheckAcctResponseParams `json:"Response"`
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

// Predefined struct for user
type CheckAmountRequestParams struct {
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

// Predefined struct for user
type CheckAmountResponseParams struct {
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
}

type CheckAmountResponse struct {
	*tchttp.BaseResponse
	Response *CheckAmountResponseParams `json:"Response"`
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

type CityCodeResult struct {
	// 城市编码cityid，数字与字母的结合
	// 注意：此字段可能返回 null，表示取不到有效值。
	CityId *string `json:"CityId,omitempty" name:"CityId"`

	// 省份
	// 注意：此字段可能返回 null，表示取不到有效值。
	Province *string `json:"Province,omitempty" name:"Province"`

	// 县区
	// 注意：此字段可能返回 null，表示取不到有效值。
	District *string `json:"District,omitempty" name:"District"`

	// 城市
	// 注意：此字段可能返回 null，表示取不到有效值。
	City *string `json:"City,omitempty" name:"City"`
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

// Predefined struct for user
type CloseCloudOrderRequestParams struct {
	// 米大师分配的支付主MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 用户Id，长度不小于5位，仅支持字母和数字的组合
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 开发者订单号
	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`
}

type CloseCloudOrderRequest struct {
	*tchttp.BaseRequest
	
	// 米大师分配的支付主MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 用户Id，长度不小于5位，仅支持字母和数字的组合
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 开发者订单号
	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`
}

func (r *CloseCloudOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseCloudOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MidasAppId")
	delete(f, "UserId")
	delete(f, "OutTradeNo")
	delete(f, "MidasEnvironment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseCloudOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseCloudOrderResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CloseCloudOrderResponse struct {
	*tchttp.BaseResponse
	Response *CloseCloudOrderResponseParams `json:"Response"`
}

func (r *CloseCloudOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseCloudOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseOpenBankPaymentOrderRequestParams struct {
	// 渠道商户ID，云企付平台下发给外部接入平台。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 外部商户订单号，与ChannelOrderId不能同时为空
	OutOrderId *string `json:"OutOrderId,omitempty" name:"OutOrderId"`

	// 云企付平台订单号，与OutOrderId不能同时为空
	ChannelOrderId *string `json:"ChannelOrderId,omitempty" name:"ChannelOrderId"`

	// 接入环境。沙箱环境填 sandbox。缺省默认调用生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type CloseOpenBankPaymentOrderRequest struct {
	*tchttp.BaseRequest
	
	// 渠道商户ID，云企付平台下发给外部接入平台。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 外部商户订单号，与ChannelOrderId不能同时为空
	OutOrderId *string `json:"OutOrderId,omitempty" name:"OutOrderId"`

	// 云企付平台订单号，与OutOrderId不能同时为空
	ChannelOrderId *string `json:"ChannelOrderId,omitempty" name:"ChannelOrderId"`

	// 接入环境。沙箱环境填 sandbox。缺省默认调用生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *CloseOpenBankPaymentOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseOpenBankPaymentOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelMerchantId")
	delete(f, "OutOrderId")
	delete(f, "ChannelOrderId")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseOpenBankPaymentOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseOpenBankPaymentOrderResponseParams struct {
	// 业务系统返回码，SUCCESS表示成功，其他表示失败。
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 业务系统返回消息
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 关单响应对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *CloseOpenBankPaymentOrderResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CloseOpenBankPaymentOrderResponse struct {
	*tchttp.BaseResponse
	Response *CloseOpenBankPaymentOrderResponseParams `json:"Response"`
}

func (r *CloseOpenBankPaymentOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseOpenBankPaymentOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloseOpenBankPaymentOrderResult struct {
	// 外部商户订单号
	OutOrderId *string `json:"OutOrderId,omitempty" name:"OutOrderId"`

	// 云企付平台订单号
	ChannelOrderId *string `json:"ChannelOrderId,omitempty" name:"ChannelOrderId"`

	// 订单状态。关单成功CLOSED
	OrderStatus *string `json:"OrderStatus,omitempty" name:"OrderStatus"`
}

// Predefined struct for user
type CloseOrderRequestParams struct {
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

// Predefined struct for user
type CloseOrderResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CloseOrderResponse struct {
	*tchttp.BaseResponse
	Response *CloseOrderResponseParams `json:"Response"`
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

type CloudAttachmentInfo struct {
	// 附加项金额。
	// 附加项的金额（必须是正数，单位：分），代表积分的数量、抵扣的金额、溢价的金额、补贴的金额
	AttachmentAmount *int64 `json:"AttachmentAmount,omitempty" name:"AttachmentAmount"`

	// 附加项类型。
	// Add：加项；
	// Sub：减项；
	// Point：积分项；
	// Subsidy：补贴项。
	AttachmentType *string `json:"AttachmentType,omitempty" name:"AttachmentType"`

	// 附加项名称。
	// 当银行作为收单机构可能会对该字段有要求，请向米大师确认。
	AttachmentName *string `json:"AttachmentName,omitempty" name:"AttachmentName"`

	// 附加项编号。
	// 当银行作为收单机构可能会对该字段有要求，请向米大师确认。
	AttachmentCode *string `json:"AttachmentCode,omitempty" name:"AttachmentCode"`
}

type CloudChannelExternalUserInfo struct {
	// 渠道方用户类型，枚举值:
	// WX_OPENID 微信支付类型
	// ALIPAY_BUYERID 支付宝支付类型
	ChannelExternalUserType *string `json:"ChannelExternalUserType,omitempty" name:"ChannelExternalUserType"`

	// 渠道方用户Id
	ChannelExternalUserId *string `json:"ChannelExternalUserId,omitempty" name:"ChannelExternalUserId"`
}

type CloudClientInfo struct {
	// 场景类型。
	// wechat_ecommerce渠道 - h5支付方式，此字段必填；
	// 枚举值：
	// CLIENT_TYPE_UNKNOWN 未知;
	// CLIENT_TYPE_IOS ios系统;
	// CLIENT_TYPE_ANDROID 安卓系统;
	// CLIENT_TYPE_WAP WAP场景;
	// CLIENT_TYPE_H5 H5场景;
	ClientType *string `json:"ClientType,omitempty" name:"ClientType"`

	// 应用名称。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 网站URL。
	AppUrl *string `json:"AppUrl,omitempty" name:"AppUrl"`

	// IOS平台BundleID。
	BundleId *string `json:"BundleId,omitempty" name:"BundleId"`

	// Android平台PackageName
	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`
}

type CloudExternalChannelData struct {
	// 第三方渠道数据名。
	// PAYMENT_ORDER_EXTERNAL_REQUEST_DATA: 支付下单请求数据
	// PAYMENT_ORDER_EXTERNAL_RETURN_DATA: 支付下单返回数据
	// PAYMENT_ORDER_EXTERNAL_NOTIFY_DATA: 支付通知数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalChannelDataName *string `json:"ExternalChannelDataName,omitempty" name:"ExternalChannelDataName"`

	// 第三方渠道数据值。
	// 当ExternalChannelDataType=PAYMENT时，反序列化格式请参考[ExternalChannelPaymentDataValue](https://dev.tke.midas.qq.com/juxin-doc-next/apidocs/external-channel-data/QueryExternalChannelData.html#externalchannelpaymentdatavalue)
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalChannelDataValue *string `json:"ExternalChannelDataValue,omitempty" name:"ExternalChannelDataValue"`
}

type CloudExternalPromptGroup struct {
	// 渠道名。
	// 为米大师定义的枚举值：
	// wechat 微信渠道
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 渠道扩展促销信息列表，由各个渠道自行定义。
	// ChannelName为wechat时，组成为 <Wechat-ExternalPromptInfo>
	ExternalPromptInfoList []*CloudExternalPromptInfo `json:"ExternalPromptInfoList,omitempty" name:"ExternalPromptInfoList"`
}

type CloudExternalPromptInfo struct {
	// 优惠商品信息类型。
	ExternalPromptType *string `json:"ExternalPromptType,omitempty" name:"ExternalPromptType"`

	// 优惠商品信息数据。
	ExternalPromptValue *string `json:"ExternalPromptValue,omitempty" name:"ExternalPromptValue"`

	// 优惠商品名称。
	ExternalPromptName *string `json:"ExternalPromptName,omitempty" name:"ExternalPromptName"`
}

type CloudGlobalPayTimeInfo struct {
	// 订单开始时间。
	// 不指定时默认为当前时间。
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" name:"StartTimestamp"`

	// 订单结束时间。
	// 逾期将会拒绝下单。不指定时默认为当前时间的7天后结束。
	ExpireTimestamp *int64 `json:"ExpireTimestamp,omitempty" name:"ExpireTimestamp"`

	// 时区。
	// 不指定时默认为28800，表示北京时间（东八区）。
	TimeOffset *int64 `json:"TimeOffset,omitempty" name:"TimeOffset"`
}

type CloudOrderReturn struct {
	// 米大师分配的支付主MidasAppId
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// 开发者支付订单号
	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`

	// 调用下单接口传进来的子单列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubOrderList []*CloudSubOrderReturn `json:"SubOrderList,omitempty" name:"SubOrderList"`

	// 调用下单接口获取的米大师交易订单号
	TransactionId *string `json:"TransactionId,omitempty" name:"TransactionId"`

	// 用户Id
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 支付渠道
	// wechat:微信支付
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// 物品Id
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 发货标识，由开发者在调用下单接口的时候传入
	Metadata *string `json:"Metadata,omitempty" name:"Metadata"`

	// ISO货币代码
	CurrencyType *string `json:"CurrencyType,omitempty" name:"CurrencyType"`

	// 支付金额，单位：分
	Amt *int64 `json:"Amt,omitempty" name:"Amt"`

	// 订单状态
	// 0:初始状态，获取米大师交易订单成功
	// 1:拉起米大师支付页面成功，用户未支付
	// 2:用户支付成功，正在发货
	// 3:用户支付成功，发货失败
	// 4:用户支付成功，发货成功
	// 5:关单中
	// 6:已关单
	OrderState *string `json:"OrderState,omitempty" name:"OrderState"`

	// 下单时间，unix时间戳
	OrderTime *string `json:"OrderTime,omitempty" name:"OrderTime"`

	// 支付时间，unix时间戳
	PayTime *string `json:"PayTime,omitempty" name:"PayTime"`

	// 支付回调时间，unix时间戳
	CallBackTime *string `json:"CallBackTime,omitempty" name:"CallBackTime"`

	// 支付机构订单号
	ChannelExternalOrderId *string `json:"ChannelExternalOrderId,omitempty" name:"ChannelExternalOrderId"`

	// 米大师内部渠道订单号
	ChannelOrderId *string `json:"ChannelOrderId,omitempty" name:"ChannelOrderId"`

	// 是否曾退款
	RefundFlag *string `json:"RefundFlag,omitempty" name:"RefundFlag"`

	// 用户支付金额
	CashAmt *string `json:"CashAmt,omitempty" name:"CashAmt"`

	// 抵扣券金额
	CouponAmt *string `json:"CouponAmt,omitempty" name:"CouponAmt"`

	// 商品名称
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 结算信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SettleInfo *CloudSettleInfo `json:"SettleInfo,omitempty" name:"SettleInfo"`

	// 附加项信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttachmentInfoList []*CloudAttachmentInfo `json:"AttachmentInfoList,omitempty" name:"AttachmentInfoList"`

	// 渠道方返回的用户信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelExternalUserInfoList []*CloudChannelExternalUserInfo `json:"ChannelExternalUserInfoList,omitempty" name:"ChannelExternalUserInfoList"`

	// 渠道扩展促销列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalReturnPromptGroupList []*CloudExternalPromptGroup `json:"ExternalReturnPromptGroupList,omitempty" name:"ExternalReturnPromptGroupList"`

	// 场景扩展信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SceneInfo *string `json:"SceneInfo,omitempty" name:"SceneInfo"`
}

type CloudSettleInfo struct {
	// 是否需要支付确认。
	// 0: 不需要支付确认
	// 1: 需要支付确认
	// 传1时，需要在支付完成后成功调用了《支付确认》接口，该笔订单才会被清分出去
	NeedToBeConfirmed *int64 `json:"NeedToBeConfirmed,omitempty" name:"NeedToBeConfirmed"`

	// 是否指定分账。
	// 0: 不指定分账
	// 1: 指定分账
	ProfitSharing *int64 `json:"ProfitSharing,omitempty" name:"ProfitSharing"`
}

type CloudStoreInfo struct {
	// 门店ID。
	StoreId *string `json:"StoreId,omitempty" name:"StoreId"`

	// 门店名称。
	StoreName *string `json:"StoreName,omitempty" name:"StoreName"`

	// 门店地址。
	StoreAddress *string `json:"StoreAddress,omitempty" name:"StoreAddress"`

	// 门店地区代码。
	StoreAreaCode *string `json:"StoreAreaCode,omitempty" name:"StoreAreaCode"`

	// 设备ID。
	// wechat_ecommerce渠道 - h5支付方式，此字段必填。
	StoreDeviceId *string `json:"StoreDeviceId,omitempty" name:"StoreDeviceId"`
}

type CloudSubOrder struct {
	// 子订单号。
	// 长度32个字符供参考，部分渠道存在长度更短的情况接入时请联系开发咨询。
	SubOutTradeNo *string `json:"SubOutTradeNo,omitempty" name:"SubOutTradeNo"`

	// 支付子商户ID。
	// 米大师计费SubAppId，代表子商户。
	SubAppId *string `json:"SubAppId,omitempty" name:"SubAppId"`

	// 商品名称。
	// 业务自定义的子订单商品名称，无需URL编码，长度限制以具体所接入渠道为准。
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 商品详情。
	// 业务自定义的子订单商品详情，无需URL编码，长度限制以具体所接入渠道为准。
	ProductDetail *string `json:"ProductDetail,omitempty" name:"ProductDetail"`

	// 平台应收。
	// 子订单平台应收金额，单位：分，需要注意的是Amt = PlatformIncome+SubMchIncome。
	PlatformIncome *int64 `json:"PlatformIncome,omitempty" name:"PlatformIncome"`

	// 商户应收。
	// 子订单结算应收金额，单位：分，需要注意的是Amt = PlatformIncome+SubMchIncome。
	SubMchIncome *int64 `json:"SubMchIncome,omitempty" name:"SubMchIncome"`

	// 透传字段。
	// 发货标识，由开发者在调用米大师下单接口的 时候下发。
	Metadata *string `json:"Metadata,omitempty" name:"Metadata"`

	// 支付金额。
	// 子订单支付金额，需要注意的是Amt = PlatformIncome+SubMchIncome。
	Amt *int64 `json:"Amt,omitempty" name:"Amt"`

	// 原始金额。
	// 子订单原始金额，OriginalAmt>=Amt。
	OriginalAmt *int64 `json:"OriginalAmt,omitempty" name:"OriginalAmt"`

	// 微信子商户号。
	WxSubMchId *string `json:"WxSubMchId,omitempty" name:"WxSubMchId"`

	// 结算信息。
	// 例如是否需要分账、是否需要支付确认等。
	SettleInfo *CloudSettleInfo `json:"SettleInfo,omitempty" name:"SettleInfo"`

	// 附加项信息列表。
	// 例如溢价信息、抵扣信息、积分信息、补贴信息
	// 通过该字段可以实现渠道方的优惠抵扣补贴等营销功能。
	AttachmentInfoList []*CloudAttachmentInfo `json:"AttachmentInfoList,omitempty" name:"AttachmentInfoList"`
}

type CloudSubOrderRefund struct {
	// 子订单退款金额
	RefundAmt *int64 `json:"RefundAmt,omitempty" name:"RefundAmt"`

	// 平台应退金额
	PlatformRefundAmt *int64 `json:"PlatformRefundAmt,omitempty" name:"PlatformRefundAmt"`

	// 商家应退金额
	SubMchRefundAmt *int64 `json:"SubMchRefundAmt,omitempty" name:"SubMchRefundAmt"`

	// 子订单号
	SubOutTradeNo *string `json:"SubOutTradeNo,omitempty" name:"SubOutTradeNo"`

	// 子退款单号，调用方需要保证全局唯一性
	SubRefundId *string `json:"SubRefundId,omitempty" name:"SubRefundId"`
}

type CloudSubOrderReturn struct {
	// 子订单号
	SubOutTradeNo *string `json:"SubOutTradeNo,omitempty" name:"SubOutTradeNo"`

	// 米大师计费SubAppId，代表子商户
	SubAppId *string `json:"SubAppId,omitempty" name:"SubAppId"`

	// 子订单商品名称
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 子订单商品详情
	ProductDetail *string `json:"ProductDetail,omitempty" name:"ProductDetail"`

	// 子订单平台应收金额，单位：分
	PlatformIncome *int64 `json:"PlatformIncome,omitempty" name:"PlatformIncome"`

	// 子订单结算应收金额，单位：分
	SubMchIncome *int64 `json:"SubMchIncome,omitempty" name:"SubMchIncome"`

	// 子订单支付金额
	Amt *int64 `json:"Amt,omitempty" name:"Amt"`

	// 子订单原始金额
	OriginalAmt *int64 `json:"OriginalAmt,omitempty" name:"OriginalAmt"`

	// 核销状态，1表示核销，0表示未核销
	SettleCheck *int64 `json:"SettleCheck,omitempty" name:"SettleCheck"`

	// 结算信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SettleInfo *CloudSettleInfo `json:"SettleInfo,omitempty" name:"SettleInfo"`

	// 透传字段，由开发者在调用米大师下单接口的时候下发
	Metadata *string `json:"Metadata,omitempty" name:"Metadata"`

	// 附加项信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttachmentInfoList *CloudAttachmentInfo `json:"AttachmentInfoList,omitempty" name:"AttachmentInfoList"`

	// 渠道方应答的订单号，透传处理
	ChannelExternalSubOrderId *string `json:"ChannelExternalSubOrderId,omitempty" name:"ChannelExternalSubOrderId"`

	// 微信子商户号
	WxSubMchId *string `json:"WxSubMchId,omitempty" name:"WxSubMchId"`
}

type CloudSubRefundItem struct {
	// 渠道方应答的退款ID，透传处理
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelExternalRefundId *string `json:"ChannelExternalRefundId,omitempty" name:"ChannelExternalRefundId"`

	// 渠道方应答的订单号，透传处理
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelExternalOrderId *string `json:"ChannelExternalOrderId,omitempty" name:"ChannelExternalOrderId"`

	// 子单退款金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	RefundAmt *int64 `json:"RefundAmt,omitempty" name:"RefundAmt"`

	// 子单订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubOutTradeNo *string `json:"SubOutTradeNo,omitempty" name:"SubOutTradeNo"`

	// 子单退款id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubRefundId *string `json:"SubRefundId,omitempty" name:"SubRefundId"`

	// 子应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubAppId *string `json:"SubAppId,omitempty" name:"SubAppId"`

	// 渠道子单支付订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelSubOrderId *string `json:"ChannelSubOrderId,omitempty" name:"ChannelSubOrderId"`

	// 渠道子退款订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelSubRefundId *string `json:"ChannelSubRefundId,omitempty" name:"ChannelSubRefundId"`
}

// Predefined struct for user
type ConfirmOrderRequestParams struct {
	// 分配给商户的AppId
	MerchantAppId *string `json:"MerchantAppId,omitempty" name:"MerchantAppId"`

	// 平台流水号。消费订单发起成功后，返回的平台唯一订单号。
	OrderNo *string `json:"OrderNo,omitempty" name:"OrderNo"`
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

// Predefined struct for user
type ConfirmOrderResponseParams struct {
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
}

type ConfirmOrderResponse struct {
	*tchttp.BaseResponse
	Response *ConfirmOrderResponseParams `json:"Response"`
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

// Predefined struct for user
type ContractOrderRequestParams struct {
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

// Predefined struct for user
type ContractOrderResponseParams struct {
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
}

type ContractOrderResponse struct {
	*tchttp.BaseResponse
	Response *ContractOrderResponseParams `json:"Response"`
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

type ContractPayListResult struct {
	// 支付方式编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	PaymentId *string `json:"PaymentId,omitempty" name:"PaymentId"`

	// 支持的交易类型（多个以小写逗号分开，0现金，1刷卡，2主扫，3被扫，4JSPAY，5预授权）
	// 注意：此字段可能返回 null，表示取不到有效值。
	PaymentType *string `json:"PaymentType,omitempty" name:"PaymentType"`

	// 支付标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	PaymentTag *string `json:"PaymentTag,omitempty" name:"PaymentTag"`

	// 支付方式图片url路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	PaymentIcon *string `json:"PaymentIcon,omitempty" name:"PaymentIcon"`

	// 付款方式名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PaymentName *string `json:"PaymentName,omitempty" name:"PaymentName"`

	// 付款方式名称（内部名称）
	// 注意：此字段可能返回 null，表示取不到有效值。
	PaymentInternalName *string `json:"PaymentInternalName,omitempty" name:"PaymentInternalName"`

	// 支付方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	PaymentOptionOne *string `json:"PaymentOptionOne,omitempty" name:"PaymentOptionOne"`

	// 支付方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	PaymentOptionTwo *string `json:"PaymentOptionTwo,omitempty" name:"PaymentOptionTwo"`

	// 支付方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	PaymentOptionThree *string `json:"PaymentOptionThree,omitempty" name:"PaymentOptionThree"`

	// 支付方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	PaymentOptionFour *string `json:"PaymentOptionFour,omitempty" name:"PaymentOptionFour"`

	// 支付方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	PaymentOptionFive *string `json:"PaymentOptionFive,omitempty" name:"PaymentOptionFive"`

	// 支付方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	PaymentOptionSix *string `json:"PaymentOptionSix,omitempty" name:"PaymentOptionSix"`

	// 支付方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	PaymentOptionSeven *string `json:"PaymentOptionSeven,omitempty" name:"PaymentOptionSeven"`

	// 支付方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	PaymentOptionOther *string `json:"PaymentOptionOther,omitempty" name:"PaymentOptionOther"`

	// 支付方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	PaymentOptionNine *string `json:"PaymentOptionNine,omitempty" name:"PaymentOptionNine"`

	// 支付方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	PaymentOptionTen *string `json:"PaymentOptionTen,omitempty" name:"PaymentOptionTen"`
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

// Predefined struct for user
type CreateAcctRequestParams struct {
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

// Predefined struct for user
type CreateAcctResponseParams struct {
	// 聚鑫计费SubAppId，代表子商户
	SubAppId *string `json:"SubAppId,omitempty" name:"SubAppId"`

	// 银行生成的子商户账户
	SubAcctNo *string `json:"SubAcctNo,omitempty" name:"SubAcctNo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAcctResponse struct {
	*tchttp.BaseResponse
	Response *CreateAcctResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateAgentTaxPaymentInfosRequestParams struct {
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

// Predefined struct for user
type CreateAgentTaxPaymentInfosResponseParams struct {
	// 代理商完税证明批次信息
	AgentTaxPaymentBatch *AgentTaxPaymentBatch `json:"AgentTaxPaymentBatch,omitempty" name:"AgentTaxPaymentBatch"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAgentTaxPaymentInfosResponse struct {
	*tchttp.BaseResponse
	Response *CreateAgentTaxPaymentInfosResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateAnchorRequestParams struct {
	// 主播业务ID，唯一
	AnchorUid *string `json:"AnchorUid,omitempty" name:"AnchorUid"`

	// 主播姓名
	AnchorName *string `json:"AnchorName,omitempty" name:"AnchorName"`

	// 主播电话
	// _敏感信息_ 使用 __AES128-CBC-PKCS#7__ 加密
	AnchorPhone *string `json:"AnchorPhone,omitempty" name:"AnchorPhone"`

	// 主播邮箱
	// _敏感信息_ 使用 __AES128-CBC-PKCS#7__ 加密
	AnchorEmail *string `json:"AnchorEmail,omitempty" name:"AnchorEmail"`

	// 主播地址
	// _敏感信息_ 使用 __AES128-CBC-PKCS#7__ 加密
	AnchorAddress *string `json:"AnchorAddress,omitempty" name:"AnchorAddress"`

	// 主播身份证号
	// _敏感信息_ 使用 __AES128-CBC-PKCS#7__ 加密
	AnchorIdNo *string `json:"AnchorIdNo,omitempty" name:"AnchorIdNo"`

	// 主播类型
	// __KMusic__:全民K歌
	// __QMusic__:QQ音乐
	// __WeChat__:微信视频号
	AnchorType *string `json:"AnchorType,omitempty" name:"AnchorType"`

	// 主播扩展信息
	AnchorExtendInfo []*AnchorExtendInfo `json:"AnchorExtendInfo,omitempty" name:"AnchorExtendInfo"`
}

type CreateAnchorRequest struct {
	*tchttp.BaseRequest
	
	// 主播业务ID，唯一
	AnchorUid *string `json:"AnchorUid,omitempty" name:"AnchorUid"`

	// 主播姓名
	AnchorName *string `json:"AnchorName,omitempty" name:"AnchorName"`

	// 主播电话
	// _敏感信息_ 使用 __AES128-CBC-PKCS#7__ 加密
	AnchorPhone *string `json:"AnchorPhone,omitempty" name:"AnchorPhone"`

	// 主播邮箱
	// _敏感信息_ 使用 __AES128-CBC-PKCS#7__ 加密
	AnchorEmail *string `json:"AnchorEmail,omitempty" name:"AnchorEmail"`

	// 主播地址
	// _敏感信息_ 使用 __AES128-CBC-PKCS#7__ 加密
	AnchorAddress *string `json:"AnchorAddress,omitempty" name:"AnchorAddress"`

	// 主播身份证号
	// _敏感信息_ 使用 __AES128-CBC-PKCS#7__ 加密
	AnchorIdNo *string `json:"AnchorIdNo,omitempty" name:"AnchorIdNo"`

	// 主播类型
	// __KMusic__:全民K歌
	// __QMusic__:QQ音乐
	// __WeChat__:微信视频号
	AnchorType *string `json:"AnchorType,omitempty" name:"AnchorType"`

	// 主播扩展信息
	AnchorExtendInfo []*AnchorExtendInfo `json:"AnchorExtendInfo,omitempty" name:"AnchorExtendInfo"`
}

func (r *CreateAnchorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAnchorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AnchorUid")
	delete(f, "AnchorName")
	delete(f, "AnchorPhone")
	delete(f, "AnchorEmail")
	delete(f, "AnchorAddress")
	delete(f, "AnchorIdNo")
	delete(f, "AnchorType")
	delete(f, "AnchorExtendInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAnchorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAnchorResponseParams struct {
	// 主播ID
	AnchorId *string `json:"AnchorId,omitempty" name:"AnchorId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAnchorResponse struct {
	*tchttp.BaseResponse
	Response *CreateAnchorResponseParams `json:"Response"`
}

func (r *CreateAnchorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAnchorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateBatchPaymentBatchData struct {
	// 订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 交易流水号
	// 注意：此字段可能返回 null，表示取不到有效值。
	TradeSerialNo *string `json:"TradeSerialNo,omitempty" name:"TradeSerialNo"`

	// 交易状态。
	// 0 处理中  
	// 1 预占成功 
	// 2 交易成功 
	// 3 交易失败 
	// 4 未知渠道异常 
	// 5 预占额度失败
	// 6 提交成功
	// 7 提交失败
	// 8 订单重复提交
	// 99 未知系统异常
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 状态描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

	// 代理商ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`

	// 代理商名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentName *string `json:"AgentName,omitempty" name:"AgentName"`
}

type CreateBatchPaymentData struct {
	// 批次号
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// 批次列表详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchInfoList []*CreateBatchPaymentBatchData `json:"BatchInfoList,omitempty" name:"BatchInfoList"`
}

type CreateBatchPaymentRecipient struct {
	// 转账金额
	TransferAmount *int64 `json:"TransferAmount,omitempty" name:"TransferAmount"`

	// 订单号
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 主播ID（与主播业务ID不能同时为空，两者都填取主播ID）
	AnchorId *string `json:"AnchorId,omitempty" name:"AnchorId"`

	// 主播业务ID（与主播业务ID不能同时为空，两者都填取主播ID）
	Uid *string `json:"Uid,omitempty" name:"Uid"`

	// 主播名称。如果该字段填入，则会对AnchorName和AnchorId/Uid进行校验。
	AnchorName *string `json:"AnchorName,omitempty" name:"AnchorName"`

	// 业务备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 子单请求预留字段
	ReqReserved *string `json:"ReqReserved,omitempty" name:"ReqReserved"`
}

// Predefined struct for user
type CreateBatchPaymentRequestParams struct {
	// 1 微信企业付款 
	// 2 支付宝转账 
	// 3 平安银企直连代发转账
	TransferType *int64 `json:"TransferType,omitempty" name:"TransferType"`

	// 转账详情
	RecipientList []*CreateBatchPaymentRecipient `json:"RecipientList,omitempty" name:"RecipientList"`

	// 请求预留字段
	ReqReserved *string `json:"ReqReserved,omitempty" name:"ReqReserved"`

	// 回调Url
	NotifyUrl *string `json:"NotifyUrl,omitempty" name:"NotifyUrl"`
}

type CreateBatchPaymentRequest struct {
	*tchttp.BaseRequest
	
	// 1 微信企业付款 
	// 2 支付宝转账 
	// 3 平安银企直连代发转账
	TransferType *int64 `json:"TransferType,omitempty" name:"TransferType"`

	// 转账详情
	RecipientList []*CreateBatchPaymentRecipient `json:"RecipientList,omitempty" name:"RecipientList"`

	// 请求预留字段
	ReqReserved *string `json:"ReqReserved,omitempty" name:"ReqReserved"`

	// 回调Url
	NotifyUrl *string `json:"NotifyUrl,omitempty" name:"NotifyUrl"`
}

func (r *CreateBatchPaymentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBatchPaymentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TransferType")
	delete(f, "RecipientList")
	delete(f, "ReqReserved")
	delete(f, "NotifyUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBatchPaymentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBatchPaymentResponseParams struct {
	// 错误码。响应成功："SUCCESS"，其他为不成功。
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 响应消息。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回响应
	Result *CreateBatchPaymentData `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateBatchPaymentResponse struct {
	*tchttp.BaseResponse
	Response *CreateBatchPaymentResponseParams `json:"Response"`
}

func (r *CreateBatchPaymentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBatchPaymentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudSubMerchantRequestParams struct {
	// 米大师分配的支付主MidasAppId，根应用Id。
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 父应用Id。
	ParentAppId *string `json:"ParentAppId,omitempty" name:"ParentAppId"`

	// 子商户名。
	SubMchName *string `json:"SubMchName,omitempty" name:"SubMchName"`

	// 子商户描述。
	SubMchDescription *string `json:"SubMchDescription,omitempty" name:"SubMchDescription"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`

	// 子应用Id，为空则自动创建子应用id。
	SubAppId *string `json:"SubAppId,omitempty" name:"SubAppId"`

	// 子商户名缩写。
	SubMchShortName *string `json:"SubMchShortName,omitempty" name:"SubMchShortName"`

	// 业务平台自定义的子商户Id，唯一。
	OutSubMerchantId *string `json:"OutSubMerchantId,omitempty" name:"OutSubMerchantId"`
}

type CreateCloudSubMerchantRequest struct {
	*tchttp.BaseRequest
	
	// 米大师分配的支付主MidasAppId，根应用Id。
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 父应用Id。
	ParentAppId *string `json:"ParentAppId,omitempty" name:"ParentAppId"`

	// 子商户名。
	SubMchName *string `json:"SubMchName,omitempty" name:"SubMchName"`

	// 子商户描述。
	SubMchDescription *string `json:"SubMchDescription,omitempty" name:"SubMchDescription"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`

	// 子应用Id，为空则自动创建子应用id。
	SubAppId *string `json:"SubAppId,omitempty" name:"SubAppId"`

	// 子商户名缩写。
	SubMchShortName *string `json:"SubMchShortName,omitempty" name:"SubMchShortName"`

	// 业务平台自定义的子商户Id，唯一。
	OutSubMerchantId *string `json:"OutSubMerchantId,omitempty" name:"OutSubMerchantId"`
}

func (r *CreateCloudSubMerchantRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudSubMerchantRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MidasAppId")
	delete(f, "ParentAppId")
	delete(f, "SubMchName")
	delete(f, "SubMchDescription")
	delete(f, "MidasEnvironment")
	delete(f, "SubAppId")
	delete(f, "SubMchShortName")
	delete(f, "OutSubMerchantId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloudSubMerchantRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudSubMerchantResponseParams struct {
	// 子应用Id。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubAppId *string `json:"SubAppId,omitempty" name:"SubAppId"`

	// 渠道子商户Id。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelSubMerchantId *string `json:"ChannelSubMerchantId,omitempty" name:"ChannelSubMerchantId"`

	// 层级，从0开始。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Level *int64 `json:"Level,omitempty" name:"Level"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCloudSubMerchantResponse struct {
	*tchttp.BaseResponse
	Response *CreateCloudSubMerchantResponseParams `json:"Response"`
}

func (r *CreateCloudSubMerchantResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudSubMerchantResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCustAcctIdRequestParams struct {
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

// Predefined struct for user
type CreateCustAcctIdResponseParams struct {
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
}

type CreateCustAcctIdResponse struct {
	*tchttp.BaseResponse
	Response *CreateCustAcctIdResponseParams `json:"Response"`
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

type CreateExternalAccountBookResult struct {
	// 处理状态。
	// __SUCCESS__: 成功
	// __FAILED__: 失败
	// __PROCESSING__: 进行中。
	DealStatus *string `json:"DealStatus,omitempty" name:"DealStatus"`

	// 处理返回描述，例如失败原因等
	// 注意：此字段可能返回 null，表示取不到有效值。
	DealMessage *string `json:"DealMessage,omitempty" name:"DealMessage"`

	// 渠道电子记账本ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelAccountBookId *string `json:"ChannelAccountBookId,omitempty" name:"ChannelAccountBookId"`

	// 电子记账本对外收款的账户信息。为JSON格式字符串（成功状态下返回）。详情见附录-复杂类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CollectMoneyAccountInfo *string `json:"CollectMoneyAccountInfo,omitempty" name:"CollectMoneyAccountInfo"`
}

type CreateExternalAnchorData struct {
	// 主播Id
	AnchorId *string `json:"AnchorId,omitempty" name:"AnchorId"`
}

// Predefined struct for user
type CreateExternalAnchorRequestParams struct {
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

	// 指定分配的代理商ID
	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`
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

	// 指定分配的代理商ID
	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`
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
	delete(f, "AgentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateExternalAnchorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateExternalAnchorResponseParams struct {
	// 错误码。响应成功："SUCCESS"，其他为不成功。
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 响应消息。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回响应
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *CreateExternalAnchorData `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateExternalAnchorResponse struct {
	*tchttp.BaseResponse
	Response *CreateExternalAnchorResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateFlexPayeeRequestParams struct {
	// 用户外部业务ID
	OutUserId *string `json:"OutUserId,omitempty" name:"OutUserId"`

	// 姓名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 证件号
	IdNo *string `json:"IdNo,omitempty" name:"IdNo"`

	// 账户名称
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// 服务商ID
	ServiceProviderId *string `json:"ServiceProviderId,omitempty" name:"ServiceProviderId"`

	// 计税信息
	TaxInfo *PayeeTaxInfo `json:"TaxInfo,omitempty" name:"TaxInfo"`

	// 证件类型
	// 0:身份证
	// 1:社会信用代码
	IdType *int64 `json:"IdType,omitempty" name:"IdType"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 手机号码
	PhoneNo *string `json:"PhoneNo,omitempty" name:"PhoneNo"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// __test__:测试环境
	// 缺省默认为生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type CreateFlexPayeeRequest struct {
	*tchttp.BaseRequest
	
	// 用户外部业务ID
	OutUserId *string `json:"OutUserId,omitempty" name:"OutUserId"`

	// 姓名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 证件号
	IdNo *string `json:"IdNo,omitempty" name:"IdNo"`

	// 账户名称
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// 服务商ID
	ServiceProviderId *string `json:"ServiceProviderId,omitempty" name:"ServiceProviderId"`

	// 计税信息
	TaxInfo *PayeeTaxInfo `json:"TaxInfo,omitempty" name:"TaxInfo"`

	// 证件类型
	// 0:身份证
	// 1:社会信用代码
	IdType *int64 `json:"IdType,omitempty" name:"IdType"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 手机号码
	PhoneNo *string `json:"PhoneNo,omitempty" name:"PhoneNo"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// __test__:测试环境
	// 缺省默认为生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *CreateFlexPayeeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlexPayeeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OutUserId")
	delete(f, "Name")
	delete(f, "IdNo")
	delete(f, "AccountName")
	delete(f, "ServiceProviderId")
	delete(f, "TaxInfo")
	delete(f, "IdType")
	delete(f, "Remark")
	delete(f, "PhoneNo")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFlexPayeeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlexPayeeResponseParams struct {
	// 错误码。SUCCESS为成功，其他为失败
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *CreateFlexPayeeResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateFlexPayeeResponse struct {
	*tchttp.BaseResponse
	Response *CreateFlexPayeeResponseParams `json:"Response"`
}

func (r *CreateFlexPayeeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlexPayeeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateFlexPayeeResult struct {
	// 收款用户ID
	PayeeId *string `json:"PayeeId,omitempty" name:"PayeeId"`
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

// Predefined struct for user
type CreateInvoiceRequestParams struct {
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

// Predefined struct for user
type CreateInvoiceResponseParams struct {
	// 发票开具结果
	Result *CreateInvoiceResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateInvoiceResponse struct {
	*tchttp.BaseResponse
	Response *CreateInvoiceResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateInvoiceV2RequestParams struct {
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

// Predefined struct for user
type CreateInvoiceV2ResponseParams struct {
	// 发票开具结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *CreateInvoiceResultV2 `json:"Result,omitempty" name:"Result"`

	// 错误码
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateInvoiceV2Response struct {
	*tchttp.BaseResponse
	Response *CreateInvoiceV2ResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateMerchantRequestParams struct {
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

// Predefined struct for user
type CreateMerchantResponseParams struct {
	// 商户注册结果
	Result *CreateMerchantResult `json:"Result,omitempty" name:"Result"`

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

// Predefined struct for user
type CreateOpenBankAggregatedSubMerchantRegistrationRequestParams struct {
	// 外部进件序列号。
	OutRegistrationNo *string `json:"OutRegistrationNo,omitempty" name:"OutRegistrationNo"`

	// 渠道商户ID。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 外部子商户ID。
	OutSubMerchantId *string `json:"OutSubMerchantId,omitempty" name:"OutSubMerchantId"`

	// 渠道名称。详见附录-云企付枚举类说明-ChannelName。
	// TENPAY: 商企付
	// WECHAT: 微信支付
	// ALIPAY: 支付宝
	// HELIPAY:合利宝
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 外部子商户类型。
	// ENTERPRISE：企业商户 
	// INSTITUTION：事业单位商户 
	// INDIVIDUALBISS：个体工商户 
	// PERSON：个人商户(小微商户) 
	// SUBJECT_TYPE_OTHERS：其他组织
	OutSubMerchantType *string `json:"OutSubMerchantType,omitempty" name:"OutSubMerchantType"`

	// 外部子商户名称。
	// HELIPAY渠道(长度不能小于5大于150)。
	OutSubMerchantName *string `json:"OutSubMerchantName,omitempty" name:"OutSubMerchantName"`

	// 商户法人代表信息。
	LegalPersonInfo *LegalPersonInfo `json:"LegalPersonInfo,omitempty" name:"LegalPersonInfo"`

	// 营业证件信息。
	BusinessLicenseInfo *BusinessLicenseInfo `json:"BusinessLicenseInfo,omitempty" name:"BusinessLicenseInfo"`

	// 支付渠道子商户进件信息。
	// json字符串，详情见附录-复杂类型-InterConnectionSubMerchantData。
	InterConnectionSubMerchantData *string `json:"InterConnectionSubMerchantData,omitempty" name:"InterConnectionSubMerchantData"`

	// 支付方式。详见附录-云企付枚举类说明-PaymentMethod。
	// 合利宝渠道不需要传。
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`

	// 外部子商户简称。
	// HELIPAY渠道必传(长度不能小于2大于20)。
	OutSubMerchantShortName *string `json:"OutSubMerchantShortName,omitempty" name:"OutSubMerchantShortName"`

	// 外部子商户描述。
	OutSubMerchantDescription *string `json:"OutSubMerchantDescription,omitempty" name:"OutSubMerchantDescription"`

	// 通知地址。
	NotifyUrl *string `json:"NotifyUrl,omitempty" name:"NotifyUrl"`

	// 相关自然人信息列表。
	// HELIPAY渠道必传业务联系人。
	NaturalPersonList []*NaturalPersonInfo `json:"NaturalPersonList,omitempty" name:"NaturalPersonList"`

	// 商户结算信息。
	// HELIPAY渠道必传。
	SettleInfo *SettleInfo `json:"SettleInfo,omitempty" name:"SettleInfo"`

	// 外部子商户其他公用扩展信息。
	// HELIPAY渠道必传。
	OutSubMerchantExtensionInfo *OutSubMerchantExtensionInfo `json:"OutSubMerchantExtensionInfo,omitempty" name:"OutSubMerchantExtensionInfo"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type CreateOpenBankAggregatedSubMerchantRegistrationRequest struct {
	*tchttp.BaseRequest
	
	// 外部进件序列号。
	OutRegistrationNo *string `json:"OutRegistrationNo,omitempty" name:"OutRegistrationNo"`

	// 渠道商户ID。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 外部子商户ID。
	OutSubMerchantId *string `json:"OutSubMerchantId,omitempty" name:"OutSubMerchantId"`

	// 渠道名称。详见附录-云企付枚举类说明-ChannelName。
	// TENPAY: 商企付
	// WECHAT: 微信支付
	// ALIPAY: 支付宝
	// HELIPAY:合利宝
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 外部子商户类型。
	// ENTERPRISE：企业商户 
	// INSTITUTION：事业单位商户 
	// INDIVIDUALBISS：个体工商户 
	// PERSON：个人商户(小微商户) 
	// SUBJECT_TYPE_OTHERS：其他组织
	OutSubMerchantType *string `json:"OutSubMerchantType,omitempty" name:"OutSubMerchantType"`

	// 外部子商户名称。
	// HELIPAY渠道(长度不能小于5大于150)。
	OutSubMerchantName *string `json:"OutSubMerchantName,omitempty" name:"OutSubMerchantName"`

	// 商户法人代表信息。
	LegalPersonInfo *LegalPersonInfo `json:"LegalPersonInfo,omitempty" name:"LegalPersonInfo"`

	// 营业证件信息。
	BusinessLicenseInfo *BusinessLicenseInfo `json:"BusinessLicenseInfo,omitempty" name:"BusinessLicenseInfo"`

	// 支付渠道子商户进件信息。
	// json字符串，详情见附录-复杂类型-InterConnectionSubMerchantData。
	InterConnectionSubMerchantData *string `json:"InterConnectionSubMerchantData,omitempty" name:"InterConnectionSubMerchantData"`

	// 支付方式。详见附录-云企付枚举类说明-PaymentMethod。
	// 合利宝渠道不需要传。
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`

	// 外部子商户简称。
	// HELIPAY渠道必传(长度不能小于2大于20)。
	OutSubMerchantShortName *string `json:"OutSubMerchantShortName,omitempty" name:"OutSubMerchantShortName"`

	// 外部子商户描述。
	OutSubMerchantDescription *string `json:"OutSubMerchantDescription,omitempty" name:"OutSubMerchantDescription"`

	// 通知地址。
	NotifyUrl *string `json:"NotifyUrl,omitempty" name:"NotifyUrl"`

	// 相关自然人信息列表。
	// HELIPAY渠道必传业务联系人。
	NaturalPersonList []*NaturalPersonInfo `json:"NaturalPersonList,omitempty" name:"NaturalPersonList"`

	// 商户结算信息。
	// HELIPAY渠道必传。
	SettleInfo *SettleInfo `json:"SettleInfo,omitempty" name:"SettleInfo"`

	// 外部子商户其他公用扩展信息。
	// HELIPAY渠道必传。
	OutSubMerchantExtensionInfo *OutSubMerchantExtensionInfo `json:"OutSubMerchantExtensionInfo,omitempty" name:"OutSubMerchantExtensionInfo"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *CreateOpenBankAggregatedSubMerchantRegistrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOpenBankAggregatedSubMerchantRegistrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OutRegistrationNo")
	delete(f, "ChannelMerchantId")
	delete(f, "OutSubMerchantId")
	delete(f, "ChannelName")
	delete(f, "OutSubMerchantType")
	delete(f, "OutSubMerchantName")
	delete(f, "LegalPersonInfo")
	delete(f, "BusinessLicenseInfo")
	delete(f, "InterConnectionSubMerchantData")
	delete(f, "PaymentMethod")
	delete(f, "OutSubMerchantShortName")
	delete(f, "OutSubMerchantDescription")
	delete(f, "NotifyUrl")
	delete(f, "NaturalPersonList")
	delete(f, "SettleInfo")
	delete(f, "OutSubMerchantExtensionInfo")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOpenBankAggregatedSubMerchantRegistrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOpenBankAggregatedSubMerchantRegistrationResponseParams struct {
	// 错误码。
	// __SUCCESS__: 成功
	// __其他__: 见附录-错误码表
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *CreateOpenBankExternalAggregatedSubMerchantRegistrationResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateOpenBankAggregatedSubMerchantRegistrationResponse struct {
	*tchttp.BaseResponse
	Response *CreateOpenBankAggregatedSubMerchantRegistrationResponseParams `json:"Response"`
}

func (r *CreateOpenBankAggregatedSubMerchantRegistrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOpenBankAggregatedSubMerchantRegistrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOpenBankExternalAggregatedSubMerchantRegistrationResult struct {
	// 进件状态 
	// SUCCESS: 进件成功 
	// FAILED: 进件失败
	// PROCESSING: 进件中 
	// 注意：若返回进件中，需要再次调用进件结果查询接口，查询结果。
	RegistrationStatus *string `json:"RegistrationStatus,omitempty" name:"RegistrationStatus"`

	// 进件返回描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegistrationMessage *string `json:"RegistrationMessage,omitempty" name:"RegistrationMessage"`

	// 渠道进件序列号
	ChannelRegistrationNo *string `json:"ChannelRegistrationNo,omitempty" name:"ChannelRegistrationNo"`

	// 渠道子商户ID
	ChannelSubMerchantId *string `json:"ChannelSubMerchantId,omitempty" name:"ChannelSubMerchantId"`
}

// Predefined struct for user
type CreateOpenBankExternalSubMerchantAccountBookRequestParams struct {
	// 外部账本ID
	OutAccountBookId *string `json:"OutAccountBookId,omitempty" name:"OutAccountBookId"`

	// 渠道商户ID
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 渠道子商户ID
	ChannelSubMerchantId *string `json:"ChannelSubMerchantId,omitempty" name:"ChannelSubMerchantId"`

	// 渠道名称。目前只支持支付宝
	// __TENPAY__: 商企付
	// __WECHAT__: 微信支付
	// __ALIPAY__: 支付宝
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 支付方式。目前只支持安心发支付
	// __EBANK_PAYMENT__: ebank支付
	// __OPENBANK_PAYMENT__: openbank支付
	// __SAFT_ISV__: 安心发支付
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type CreateOpenBankExternalSubMerchantAccountBookRequest struct {
	*tchttp.BaseRequest
	
	// 外部账本ID
	OutAccountBookId *string `json:"OutAccountBookId,omitempty" name:"OutAccountBookId"`

	// 渠道商户ID
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 渠道子商户ID
	ChannelSubMerchantId *string `json:"ChannelSubMerchantId,omitempty" name:"ChannelSubMerchantId"`

	// 渠道名称。目前只支持支付宝
	// __TENPAY__: 商企付
	// __WECHAT__: 微信支付
	// __ALIPAY__: 支付宝
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 支付方式。目前只支持安心发支付
	// __EBANK_PAYMENT__: ebank支付
	// __OPENBANK_PAYMENT__: openbank支付
	// __SAFT_ISV__: 安心发支付
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *CreateOpenBankExternalSubMerchantAccountBookRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOpenBankExternalSubMerchantAccountBookRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OutAccountBookId")
	delete(f, "ChannelMerchantId")
	delete(f, "ChannelSubMerchantId")
	delete(f, "ChannelName")
	delete(f, "PaymentMethod")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOpenBankExternalSubMerchantAccountBookRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOpenBankExternalSubMerchantAccountBookResponseParams struct {
	// 错误码。
	// __SUCCESS__: 成功
	// __其他__: 见附录-错误码表
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *CreateExternalAccountBookResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateOpenBankExternalSubMerchantAccountBookResponse struct {
	*tchttp.BaseResponse
	Response *CreateOpenBankExternalSubMerchantAccountBookResponseParams `json:"Response"`
}

func (r *CreateOpenBankExternalSubMerchantAccountBookResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOpenBankExternalSubMerchantAccountBookResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOpenBankExternalSubMerchantRegistrationRequestParams struct {
	// 外部进件序列号。
	OutRegistrationNo *string `json:"OutRegistrationNo,omitempty" name:"OutRegistrationNo"`

	// 渠道商户ID。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 外部子商户ID,平台侧商户唯一ID。
	OutSubMerchantId *string `json:"OutSubMerchantId,omitempty" name:"OutSubMerchantId"`

	// 渠道名称。详见附录-云企付枚举类说明-ChannelName。
	// __TENPAY__: 商企付
	// __WECHAT__: 微信支付
	// __ALIPAY__: 支付宝
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 支付方式。详见附录-云企付枚举类说明-PaymentMethod。
	// __EBANK_PAYMENT__: ebank支付
	// __OPENBANK_PAYMENT__: openbank支付
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`

	// 社会信用代码。
	BusinessLicenseNumber *string `json:"BusinessLicenseNumber,omitempty" name:"BusinessLicenseNumber"`

	// 外部子商户名称。
	OutSubMerchantName *string `json:"OutSubMerchantName,omitempty" name:"OutSubMerchantName"`

	// 法人姓名。
	LegalName *string `json:"LegalName,omitempty" name:"LegalName"`

	// 外部子商户简称。
	OutSubMerchantShortName *string `json:"OutSubMerchantShortName,omitempty" name:"OutSubMerchantShortName"`

	// 外部子商户描述。
	OutSubMerchantDescription *string `json:"OutSubMerchantDescription,omitempty" name:"OutSubMerchantDescription"`

	// 第三方子商户进件信息，为JSON格式字符串。详情见附录-复杂类型。
	ExternalSubMerchantRegistrationData *string `json:"ExternalSubMerchantRegistrationData,omitempty" name:"ExternalSubMerchantRegistrationData"`

	// 通知地址。
	NotifyUrl *string `json:"NotifyUrl,omitempty" name:"NotifyUrl"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type CreateOpenBankExternalSubMerchantRegistrationRequest struct {
	*tchttp.BaseRequest
	
	// 外部进件序列号。
	OutRegistrationNo *string `json:"OutRegistrationNo,omitempty" name:"OutRegistrationNo"`

	// 渠道商户ID。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 外部子商户ID,平台侧商户唯一ID。
	OutSubMerchantId *string `json:"OutSubMerchantId,omitempty" name:"OutSubMerchantId"`

	// 渠道名称。详见附录-云企付枚举类说明-ChannelName。
	// __TENPAY__: 商企付
	// __WECHAT__: 微信支付
	// __ALIPAY__: 支付宝
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 支付方式。详见附录-云企付枚举类说明-PaymentMethod。
	// __EBANK_PAYMENT__: ebank支付
	// __OPENBANK_PAYMENT__: openbank支付
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`

	// 社会信用代码。
	BusinessLicenseNumber *string `json:"BusinessLicenseNumber,omitempty" name:"BusinessLicenseNumber"`

	// 外部子商户名称。
	OutSubMerchantName *string `json:"OutSubMerchantName,omitempty" name:"OutSubMerchantName"`

	// 法人姓名。
	LegalName *string `json:"LegalName,omitempty" name:"LegalName"`

	// 外部子商户简称。
	OutSubMerchantShortName *string `json:"OutSubMerchantShortName,omitempty" name:"OutSubMerchantShortName"`

	// 外部子商户描述。
	OutSubMerchantDescription *string `json:"OutSubMerchantDescription,omitempty" name:"OutSubMerchantDescription"`

	// 第三方子商户进件信息，为JSON格式字符串。详情见附录-复杂类型。
	ExternalSubMerchantRegistrationData *string `json:"ExternalSubMerchantRegistrationData,omitempty" name:"ExternalSubMerchantRegistrationData"`

	// 通知地址。
	NotifyUrl *string `json:"NotifyUrl,omitempty" name:"NotifyUrl"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *CreateOpenBankExternalSubMerchantRegistrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOpenBankExternalSubMerchantRegistrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OutRegistrationNo")
	delete(f, "ChannelMerchantId")
	delete(f, "OutSubMerchantId")
	delete(f, "ChannelName")
	delete(f, "PaymentMethod")
	delete(f, "BusinessLicenseNumber")
	delete(f, "OutSubMerchantName")
	delete(f, "LegalName")
	delete(f, "OutSubMerchantShortName")
	delete(f, "OutSubMerchantDescription")
	delete(f, "ExternalSubMerchantRegistrationData")
	delete(f, "NotifyUrl")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOpenBankExternalSubMerchantRegistrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOpenBankExternalSubMerchantRegistrationResponseParams struct {
	// 错误码。
	// __SUCCESS__: 成功
	// __其他__: 见附录-错误码表
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *CreateOpenBankExternalSubMerchantRegistrationResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateOpenBankExternalSubMerchantRegistrationResponse struct {
	*tchttp.BaseResponse
	Response *CreateOpenBankExternalSubMerchantRegistrationResponseParams `json:"Response"`
}

func (r *CreateOpenBankExternalSubMerchantRegistrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOpenBankExternalSubMerchantRegistrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOpenBankExternalSubMerchantRegistrationResult struct {
	// 进件状态。
	// __SUCCESS__: 进件成功
	// __FAILED__: 进件失败
	// __PROCESSING__: 进件中
	// 注意：若返回进件中，需要再次调用进件结果查询接口，查询结果。
	RegistrationStatus *string `json:"RegistrationStatus,omitempty" name:"RegistrationStatus"`

	// 进件返回描述, 例如失败原因等。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegistrationMessage *string `json:"RegistrationMessage,omitempty" name:"RegistrationMessage"`

	// 渠道进件序列号。
	ChannelRegistrationNo *string `json:"ChannelRegistrationNo,omitempty" name:"ChannelRegistrationNo"`

	// 渠道子商户ID。
	ChannelSubMerchantId *string `json:"ChannelSubMerchantId,omitempty" name:"ChannelSubMerchantId"`

	// 第三方渠道返回信息, 为JSON格式字符串。详情见附录-复杂类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalReturnData *string `json:"ExternalReturnData,omitempty" name:"ExternalReturnData"`
}

// Predefined struct for user
type CreateOpenBankMerchantRequestParams struct {
	// 外部商户ID。
	OutMerchantId *string `json:"OutMerchantId,omitempty" name:"OutMerchantId"`

	// 渠道名称。
	// __TENPAY__: 商企付
	// __WECHAT__: 微信支付
	// __ALIPAY__: 支付宝
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 外部商户名称。
	OutMerchantName *string `json:"OutMerchantName,omitempty" name:"OutMerchantName"`

	// 第三方渠道商户信息。详情见附录-复杂类型。
	ExternalMerchantInfo *string `json:"ExternalMerchantInfo,omitempty" name:"ExternalMerchantInfo"`

	// 外部商户简称。
	OutMerchantShortName *string `json:"OutMerchantShortName,omitempty" name:"OutMerchantShortName"`

	// 外部商户描述
	OutMerchantDescription *string `json:"OutMerchantDescription,omitempty" name:"OutMerchantDescription"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type CreateOpenBankMerchantRequest struct {
	*tchttp.BaseRequest
	
	// 外部商户ID。
	OutMerchantId *string `json:"OutMerchantId,omitempty" name:"OutMerchantId"`

	// 渠道名称。
	// __TENPAY__: 商企付
	// __WECHAT__: 微信支付
	// __ALIPAY__: 支付宝
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 外部商户名称。
	OutMerchantName *string `json:"OutMerchantName,omitempty" name:"OutMerchantName"`

	// 第三方渠道商户信息。详情见附录-复杂类型。
	ExternalMerchantInfo *string `json:"ExternalMerchantInfo,omitempty" name:"ExternalMerchantInfo"`

	// 外部商户简称。
	OutMerchantShortName *string `json:"OutMerchantShortName,omitempty" name:"OutMerchantShortName"`

	// 外部商户描述
	OutMerchantDescription *string `json:"OutMerchantDescription,omitempty" name:"OutMerchantDescription"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *CreateOpenBankMerchantRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOpenBankMerchantRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OutMerchantId")
	delete(f, "ChannelName")
	delete(f, "OutMerchantName")
	delete(f, "ExternalMerchantInfo")
	delete(f, "OutMerchantShortName")
	delete(f, "OutMerchantDescription")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOpenBankMerchantRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOpenBankMerchantResponseParams struct {
	// 错误码。
	// __SUCCESS__: 成功
	// __其他__: 见附录-错误码表
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *CreateOpenBankMerchantResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateOpenBankMerchantResponse struct {
	*tchttp.BaseResponse
	Response *CreateOpenBankMerchantResponseParams `json:"Response"`
}

func (r *CreateOpenBankMerchantResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOpenBankMerchantResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOpenBankMerchantResult struct {
	// 渠道商户ID。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`
}

type CreateOpenBankOrderPaymentResult struct {
	// 云企付平台订单号。
	ChannelOrderId *string `json:"ChannelOrderId,omitempty" name:"ChannelOrderId"`

	// 第三方支付平台返回支付订单号。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ThirdPayOrderId *string `json:"ThirdPayOrderId,omitempty" name:"ThirdPayOrderId"`

	// 跳转参数
	// 渠道为TENPAY，付款方式为EBANK_PAYMENT时必选。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RedirectInfo *OpenBankRedirectInfo `json:"RedirectInfo,omitempty" name:"RedirectInfo"`

	// 外部商户订单号，只能是数字、大小写字母，且在同一个接入平台下唯一。
	OutOrderId *string `json:"OutOrderId,omitempty" name:"OutOrderId"`
}

type CreateOpenBankOrderRechargeResult struct {
	// 云企付平台订单号。
	ChannelOrderId *string `json:"ChannelOrderId,omitempty" name:"ChannelOrderId"`

	// 第三方支付平台返回支付订单号。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ThirdPayOrderId *string `json:"ThirdPayOrderId,omitempty" name:"ThirdPayOrderId"`

	// 跳转参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RedirectInfo *OpenBankRechargeRedirectInfo `json:"RedirectInfo,omitempty" name:"RedirectInfo"`

	// 外部商户订单号，只能是数字、大小写字母，且在同一个接入平台下唯一。
	OutOrderId *string `json:"OutOrderId,omitempty" name:"OutOrderId"`
}

// Predefined struct for user
type CreateOpenBankPaymentOrderRequestParams struct {
	// 云企付渠道商户号。外部接入平台入驻云企付平台后下发。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 渠道名称。详见附录-云企付枚举类说明-ChannelName。
	// __TENPAY__: 商企付
	// __WECHAT__: 微信支付
	// __ALIPAY__: 支付宝
	// __WECHAT__: 微信支付
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 付款方式。详见附录-云企付枚举类说明-PaymentMethod。
	// __EBANK_PAYMENT__:B2B EBank付款
	// __OPENBANK_PAYMENT__:B2C  openbank付款
	// __SAFT_ISV__:支付宝安心发
	// __TRANS_TO_CHANGE__: 微信支付转账到零钱v2
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`

	// 付款模式。默认直接支付，如
	// __DIRECT__:直接支付
	// __FREEZE__:担保支付
	PaymentMode *string `json:"PaymentMode,omitempty" name:"PaymentMode"`

	// 外部订单号,只能是数字、大小写字母，且在同一个接入平台下唯一，限定长度40位。
	OutOrderId *string `json:"OutOrderId,omitempty" name:"OutOrderId"`

	// 付款金额，单位分。
	TotalAmount *int64 `json:"TotalAmount,omitempty" name:"TotalAmount"`

	// 固定值CNY。
	Currency *string `json:"Currency,omitempty" name:"Currency"`

	// 付款方信息。
	PayerInfo *OpenBankPayerInfo `json:"PayerInfo,omitempty" name:"PayerInfo"`

	// 收款方信息。
	PayeeInfo *OpenBankPayeeInfo `json:"PayeeInfo,omitempty" name:"PayeeInfo"`

	// 通知地址，如www.test.com。
	NotifyUrl *string `json:"NotifyUrl,omitempty" name:"NotifyUrl"`

	// 订单过期时间，yyyy-MM-dd HH:mm:ss格式。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 前端成功回调URL。条件可选。
	FrontUrl *string `json:"FrontUrl,omitempty" name:"FrontUrl"`

	// 前端刷新 URL。条件可选。
	RefreshUrl *string `json:"RefreshUrl,omitempty" name:"RefreshUrl"`

	// 设备信息，条件可选。
	SceneInfo *OpenBankSceneInfo `json:"SceneInfo,omitempty" name:"SceneInfo"`

	// 商品信息，条件可选。
	GoodsInfo *OpenBankGoodsInfo `json:"GoodsInfo,omitempty" name:"GoodsInfo"`

	// 附加信息，查询时原样返回。
	Attachment *string `json:"Attachment,omitempty" name:"Attachment"`

	// 若不上传，即使用默认值无需分润
	// __NO_NEED_SHARE__：无需分润；
	// __SHARE_BY_INFO__：分润时指定金额，此时如果分润信息 ProfitShareInfo为空，只冻结，不分账给其他商户；需要调用解冻接口。
	// __SHARE_BY_API__：后续调用分润接口决定分润金额
	ProfitShareFlag *string `json:"ProfitShareFlag,omitempty" name:"ProfitShareFlag"`

	// 分润信息，配合ProfitShareFlag使用。
	ProfitShareInfoList []*OpenBankProfitShareInfo `json:"ProfitShareInfoList,omitempty" name:"ProfitShareInfoList"`

	// 备注信息。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// 缺省默认为生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type CreateOpenBankPaymentOrderRequest struct {
	*tchttp.BaseRequest
	
	// 云企付渠道商户号。外部接入平台入驻云企付平台后下发。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 渠道名称。详见附录-云企付枚举类说明-ChannelName。
	// __TENPAY__: 商企付
	// __WECHAT__: 微信支付
	// __ALIPAY__: 支付宝
	// __WECHAT__: 微信支付
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 付款方式。详见附录-云企付枚举类说明-PaymentMethod。
	// __EBANK_PAYMENT__:B2B EBank付款
	// __OPENBANK_PAYMENT__:B2C  openbank付款
	// __SAFT_ISV__:支付宝安心发
	// __TRANS_TO_CHANGE__: 微信支付转账到零钱v2
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`

	// 付款模式。默认直接支付，如
	// __DIRECT__:直接支付
	// __FREEZE__:担保支付
	PaymentMode *string `json:"PaymentMode,omitempty" name:"PaymentMode"`

	// 外部订单号,只能是数字、大小写字母，且在同一个接入平台下唯一，限定长度40位。
	OutOrderId *string `json:"OutOrderId,omitempty" name:"OutOrderId"`

	// 付款金额，单位分。
	TotalAmount *int64 `json:"TotalAmount,omitempty" name:"TotalAmount"`

	// 固定值CNY。
	Currency *string `json:"Currency,omitempty" name:"Currency"`

	// 付款方信息。
	PayerInfo *OpenBankPayerInfo `json:"PayerInfo,omitempty" name:"PayerInfo"`

	// 收款方信息。
	PayeeInfo *OpenBankPayeeInfo `json:"PayeeInfo,omitempty" name:"PayeeInfo"`

	// 通知地址，如www.test.com。
	NotifyUrl *string `json:"NotifyUrl,omitempty" name:"NotifyUrl"`

	// 订单过期时间，yyyy-MM-dd HH:mm:ss格式。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 前端成功回调URL。条件可选。
	FrontUrl *string `json:"FrontUrl,omitempty" name:"FrontUrl"`

	// 前端刷新 URL。条件可选。
	RefreshUrl *string `json:"RefreshUrl,omitempty" name:"RefreshUrl"`

	// 设备信息，条件可选。
	SceneInfo *OpenBankSceneInfo `json:"SceneInfo,omitempty" name:"SceneInfo"`

	// 商品信息，条件可选。
	GoodsInfo *OpenBankGoodsInfo `json:"GoodsInfo,omitempty" name:"GoodsInfo"`

	// 附加信息，查询时原样返回。
	Attachment *string `json:"Attachment,omitempty" name:"Attachment"`

	// 若不上传，即使用默认值无需分润
	// __NO_NEED_SHARE__：无需分润；
	// __SHARE_BY_INFO__：分润时指定金额，此时如果分润信息 ProfitShareInfo为空，只冻结，不分账给其他商户；需要调用解冻接口。
	// __SHARE_BY_API__：后续调用分润接口决定分润金额
	ProfitShareFlag *string `json:"ProfitShareFlag,omitempty" name:"ProfitShareFlag"`

	// 分润信息，配合ProfitShareFlag使用。
	ProfitShareInfoList []*OpenBankProfitShareInfo `json:"ProfitShareInfoList,omitempty" name:"ProfitShareInfoList"`

	// 备注信息。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// 缺省默认为生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *CreateOpenBankPaymentOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOpenBankPaymentOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelMerchantId")
	delete(f, "ChannelName")
	delete(f, "PaymentMethod")
	delete(f, "PaymentMode")
	delete(f, "OutOrderId")
	delete(f, "TotalAmount")
	delete(f, "Currency")
	delete(f, "PayerInfo")
	delete(f, "PayeeInfo")
	delete(f, "NotifyUrl")
	delete(f, "ExpireTime")
	delete(f, "FrontUrl")
	delete(f, "RefreshUrl")
	delete(f, "SceneInfo")
	delete(f, "GoodsInfo")
	delete(f, "Attachment")
	delete(f, "ProfitShareFlag")
	delete(f, "ProfitShareInfoList")
	delete(f, "Remark")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOpenBankPaymentOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOpenBankPaymentOrderResponseParams struct {
	// 业务系统返回码，SUCCESS表示成功，其他表示失败。
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 业务系统返回消息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 统一下单响应对象。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *CreateOpenBankOrderPaymentResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateOpenBankPaymentOrderResponse struct {
	*tchttp.BaseResponse
	Response *CreateOpenBankPaymentOrderResponseParams `json:"Response"`
}

func (r *CreateOpenBankPaymentOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOpenBankPaymentOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOpenBankRechargeOrderRequestParams struct {
	// 云企付渠道商户号。外部接入平台入驻云企付平台后下发。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 外部订单号,只能是数字、大小写字母，且在同一个接入平台下唯一，限定长度40位。
	OutOrderId *string `json:"OutOrderId,omitempty" name:"OutOrderId"`

	// 付款金额，单位分。
	TotalAmount *int64 `json:"TotalAmount,omitempty" name:"TotalAmount"`

	// 固定值CNY。
	Currency *string `json:"Currency,omitempty" name:"Currency"`

	// 订单过期时间，yyyy-MM-dd HH:mm:ss格式。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 渠道名称。
	// __TENPAY__: 商企付
	// __WECHAT__: 微信支付
	// __ALIPAY__: 支付宝
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 渠道名称。
	// __TENPAY__: 商企付
	// __WECHAT__: 微信支付
	// __ALIPAY__: 支付宝
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`

	// 收款方信息。
	PayeeInfo *OpenBankRechargePayeeInfo `json:"PayeeInfo,omitempty" name:"PayeeInfo"`

	// 渠道子商户号
	ChannelSubMerchantId *string `json:"ChannelSubMerchantId,omitempty" name:"ChannelSubMerchantId"`

	// 通知地址，如www.test.com。
	NotifyUrl *string `json:"NotifyUrl,omitempty" name:"NotifyUrl"`

	// 备注信息。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type CreateOpenBankRechargeOrderRequest struct {
	*tchttp.BaseRequest
	
	// 云企付渠道商户号。外部接入平台入驻云企付平台后下发。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 外部订单号,只能是数字、大小写字母，且在同一个接入平台下唯一，限定长度40位。
	OutOrderId *string `json:"OutOrderId,omitempty" name:"OutOrderId"`

	// 付款金额，单位分。
	TotalAmount *int64 `json:"TotalAmount,omitempty" name:"TotalAmount"`

	// 固定值CNY。
	Currency *string `json:"Currency,omitempty" name:"Currency"`

	// 订单过期时间，yyyy-MM-dd HH:mm:ss格式。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 渠道名称。
	// __TENPAY__: 商企付
	// __WECHAT__: 微信支付
	// __ALIPAY__: 支付宝
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 渠道名称。
	// __TENPAY__: 商企付
	// __WECHAT__: 微信支付
	// __ALIPAY__: 支付宝
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`

	// 收款方信息。
	PayeeInfo *OpenBankRechargePayeeInfo `json:"PayeeInfo,omitempty" name:"PayeeInfo"`

	// 渠道子商户号
	ChannelSubMerchantId *string `json:"ChannelSubMerchantId,omitempty" name:"ChannelSubMerchantId"`

	// 通知地址，如www.test.com。
	NotifyUrl *string `json:"NotifyUrl,omitempty" name:"NotifyUrl"`

	// 备注信息。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *CreateOpenBankRechargeOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOpenBankRechargeOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelMerchantId")
	delete(f, "OutOrderId")
	delete(f, "TotalAmount")
	delete(f, "Currency")
	delete(f, "ExpireTime")
	delete(f, "ChannelName")
	delete(f, "PaymentMethod")
	delete(f, "PayeeInfo")
	delete(f, "ChannelSubMerchantId")
	delete(f, "NotifyUrl")
	delete(f, "Remark")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOpenBankRechargeOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOpenBankRechargeOrderResponseParams struct {
	// 业务系统返回码，SUCCESS表示成功，其他表示失败。
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 业务系统返回消息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 充值响应对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *CreateOpenBankOrderRechargeResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateOpenBankRechargeOrderResponse struct {
	*tchttp.BaseResponse
	Response *CreateOpenBankRechargeOrderResponseParams `json:"Response"`
}

func (r *CreateOpenBankRechargeOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOpenBankRechargeOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOpenBankSubMerchantRateConfigureRequestParams struct {
	// 渠道进件序列号。
	// 填写子商户进件返回的渠道进件编号。
	ChannelRegistrationNo *string `json:"ChannelRegistrationNo,omitempty" name:"ChannelRegistrationNo"`

	// 外部产品费率申请序列号。
	OutProductFeeNo *string `json:"OutProductFeeNo,omitempty" name:"OutProductFeeNo"`

	// 渠道商户ID。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 渠道子商户ID。
	ChannelSubMerchantId *string `json:"ChannelSubMerchantId,omitempty" name:"ChannelSubMerchantId"`

	// 渠道名称。详见附录-云企付枚举类说明-ChannelName。
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 支付类型。
	// SWIPE:刷卡
	// SCAN:扫码
	// WAP:WAP
	// PUBLIC:公众号支付
	// SDK:SDK
	// MINI_PROGRAM:小程序
	// 注意：HELIPAY渠道传SDK。
	PayType *string `json:"PayType,omitempty" name:"PayType"`

	// 支付渠道。
	// ALIPAY：支付宝 
	// WXPAY：微信支付 
	// UNIONPAY：银联
	PayChannel *string `json:"PayChannel,omitempty" name:"PayChannel"`

	// 计费模式。
	// SINGLE：按单笔金额计费
	// RATIO：按单笔费率计费 
	// RANGE：按分段区间计费
	FeeMode *string `json:"FeeMode,omitempty" name:"FeeMode"`

	// 费用值，单位（0.01%或分）。
	FeeValue *uint64 `json:"FeeValue,omitempty" name:"FeeValue"`

	// 支付方式。详见附录-云企付枚举类说明-PaymentMethod。
	// HELIPAY渠道不需要传入。
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`

	// 最低收费金额，单位（分）。
	MinFee *uint64 `json:"MinFee,omitempty" name:"MinFee"`

	// 最高收费金额，单位（分）。
	MaxFee *uint64 `json:"MaxFee,omitempty" name:"MaxFee"`

	// 通知地址。
	NotifyUrl *string `json:"NotifyUrl,omitempty" name:"NotifyUrl"`

	// 分段计费区间列表。
	FeeRangeList []*FeeRangInfo `json:"FeeRangeList,omitempty" name:"FeeRangeList"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type CreateOpenBankSubMerchantRateConfigureRequest struct {
	*tchttp.BaseRequest
	
	// 渠道进件序列号。
	// 填写子商户进件返回的渠道进件编号。
	ChannelRegistrationNo *string `json:"ChannelRegistrationNo,omitempty" name:"ChannelRegistrationNo"`

	// 外部产品费率申请序列号。
	OutProductFeeNo *string `json:"OutProductFeeNo,omitempty" name:"OutProductFeeNo"`

	// 渠道商户ID。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 渠道子商户ID。
	ChannelSubMerchantId *string `json:"ChannelSubMerchantId,omitempty" name:"ChannelSubMerchantId"`

	// 渠道名称。详见附录-云企付枚举类说明-ChannelName。
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 支付类型。
	// SWIPE:刷卡
	// SCAN:扫码
	// WAP:WAP
	// PUBLIC:公众号支付
	// SDK:SDK
	// MINI_PROGRAM:小程序
	// 注意：HELIPAY渠道传SDK。
	PayType *string `json:"PayType,omitempty" name:"PayType"`

	// 支付渠道。
	// ALIPAY：支付宝 
	// WXPAY：微信支付 
	// UNIONPAY：银联
	PayChannel *string `json:"PayChannel,omitempty" name:"PayChannel"`

	// 计费模式。
	// SINGLE：按单笔金额计费
	// RATIO：按单笔费率计费 
	// RANGE：按分段区间计费
	FeeMode *string `json:"FeeMode,omitempty" name:"FeeMode"`

	// 费用值，单位（0.01%或分）。
	FeeValue *uint64 `json:"FeeValue,omitempty" name:"FeeValue"`

	// 支付方式。详见附录-云企付枚举类说明-PaymentMethod。
	// HELIPAY渠道不需要传入。
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`

	// 最低收费金额，单位（分）。
	MinFee *uint64 `json:"MinFee,omitempty" name:"MinFee"`

	// 最高收费金额，单位（分）。
	MaxFee *uint64 `json:"MaxFee,omitempty" name:"MaxFee"`

	// 通知地址。
	NotifyUrl *string `json:"NotifyUrl,omitempty" name:"NotifyUrl"`

	// 分段计费区间列表。
	FeeRangeList []*FeeRangInfo `json:"FeeRangeList,omitempty" name:"FeeRangeList"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *CreateOpenBankSubMerchantRateConfigureRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOpenBankSubMerchantRateConfigureRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelRegistrationNo")
	delete(f, "OutProductFeeNo")
	delete(f, "ChannelMerchantId")
	delete(f, "ChannelSubMerchantId")
	delete(f, "ChannelName")
	delete(f, "PayType")
	delete(f, "PayChannel")
	delete(f, "FeeMode")
	delete(f, "FeeValue")
	delete(f, "PaymentMethod")
	delete(f, "MinFee")
	delete(f, "MaxFee")
	delete(f, "NotifyUrl")
	delete(f, "FeeRangeList")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOpenBankSubMerchantRateConfigureRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOpenBankSubMerchantRateConfigureResponseParams struct {
	// 错误码。
	// __SUCCESS__: 成功
	// __其他__: 见附录-错误码表
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *CreateOpenBankSubMerchantRateConfigureResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateOpenBankSubMerchantRateConfigureResponse struct {
	*tchttp.BaseResponse
	Response *CreateOpenBankSubMerchantRateConfigureResponseParams `json:"Response"`
}

func (r *CreateOpenBankSubMerchantRateConfigureResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOpenBankSubMerchantRateConfigureResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOpenBankSubMerchantRateConfigureResult struct {
	// 处理状态 
	// SUCCESS: 开通成功 
	// FAILED: 开通失败
	// PROCESSING: 开通中 
	// 注意：若返回开通中，需要再次调用费率配置结果查询接口，查询结果。
	DealStatus *string `json:"DealStatus,omitempty" name:"DealStatus"`

	// 处理描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	DealMessage *string `json:"DealMessage,omitempty" name:"DealMessage"`

	// 渠道产品费率序列号
	ChannelProductFeeNo *string `json:"ChannelProductFeeNo,omitempty" name:"ChannelProductFeeNo"`
}

// Predefined struct for user
type CreateOpenBankUnifiedOrderRequestParams struct {
	// 渠道商户号。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 渠道名称。
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 支付产品类型。
	// 被扫（扫码）：SWIPE, 主扫（刷卡）：SCAN, 
	// H5：WAP, 公众号：PUBLIC, 
	// APP-SDK：SDK, 小程序：MINI_PROGRAM, 
	// 快捷支付：QUICK, 网银支付：ONLINE_BANK。
	PayType *string `json:"PayType,omitempty" name:"PayType"`

	// 外部商户订单号。
	// 只能是数字、大小写字母，且在同一个接入平台下唯一。
	OutOrderId *string `json:"OutOrderId,omitempty" name:"OutOrderId"`

	// 交易金额，单位分。
	TotalAmount *int64 `json:"TotalAmount,omitempty" name:"TotalAmount"`

	// 币种。固定：CNY。
	Currency *string `json:"Currency,omitempty" name:"Currency"`

	// 渠道子商户号。
	ChannelSubMerchantId *string `json:"ChannelSubMerchantId,omitempty" name:"ChannelSubMerchantId"`

	// 实际支付渠道。没有则无需填写。如
	// 支付宝 ALIPAY
	// 微信支付 WXPAY
	// 银联 UNIONPAY
	// 一般在间连模式下使用。
	PayChannel *string `json:"PayChannel,omitempty" name:"PayChannel"`

	// 设备信息。
	SceneInfo *OpenBankSceneInfo `json:"SceneInfo,omitempty" name:"SceneInfo"`

	// 分账信息列表。
	ProfitShareInfoList []*OpenBankProfitShareInfo `json:"ProfitShareInfoList,omitempty" name:"ProfitShareInfoList"`

	// 订单标题。
	OrderSubject *string `json:"OrderSubject,omitempty" name:"OrderSubject"`

	// 商品信息。
	GoodsDetail *string `json:"GoodsDetail,omitempty" name:"GoodsDetail"`

	// 超时时间。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 支付成功回调地址。
	NotifyUrl *string `json:"NotifyUrl,omitempty" name:"NotifyUrl"`

	// 支付成功前端跳转URL。
	FrontUrl *string `json:"FrontUrl,omitempty" name:"FrontUrl"`

	// 订单附加信息，查询或者回调的时候原样返回。
	Attachment *string `json:"Attachment,omitempty" name:"Attachment"`

	// 第三方渠道扩展字段。见附录-复杂类型。
	// 未作特殊说明，则无需传入。
	ExternalPaymentData *string `json:"ExternalPaymentData,omitempty" name:"ExternalPaymentData"`

	// 备注。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 门店信息。
	StoreInfo *OpenBankStoreInfo `json:"StoreInfo,omitempty" name:"StoreInfo"`

	// 支付限制。
	PayLimitInfo *OpenBankPayLimitInfo `json:"PayLimitInfo,omitempty" name:"PayLimitInfo"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type CreateOpenBankUnifiedOrderRequest struct {
	*tchttp.BaseRequest
	
	// 渠道商户号。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 渠道名称。
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 支付产品类型。
	// 被扫（扫码）：SWIPE, 主扫（刷卡）：SCAN, 
	// H5：WAP, 公众号：PUBLIC, 
	// APP-SDK：SDK, 小程序：MINI_PROGRAM, 
	// 快捷支付：QUICK, 网银支付：ONLINE_BANK。
	PayType *string `json:"PayType,omitempty" name:"PayType"`

	// 外部商户订单号。
	// 只能是数字、大小写字母，且在同一个接入平台下唯一。
	OutOrderId *string `json:"OutOrderId,omitempty" name:"OutOrderId"`

	// 交易金额，单位分。
	TotalAmount *int64 `json:"TotalAmount,omitempty" name:"TotalAmount"`

	// 币种。固定：CNY。
	Currency *string `json:"Currency,omitempty" name:"Currency"`

	// 渠道子商户号。
	ChannelSubMerchantId *string `json:"ChannelSubMerchantId,omitempty" name:"ChannelSubMerchantId"`

	// 实际支付渠道。没有则无需填写。如
	// 支付宝 ALIPAY
	// 微信支付 WXPAY
	// 银联 UNIONPAY
	// 一般在间连模式下使用。
	PayChannel *string `json:"PayChannel,omitempty" name:"PayChannel"`

	// 设备信息。
	SceneInfo *OpenBankSceneInfo `json:"SceneInfo,omitempty" name:"SceneInfo"`

	// 分账信息列表。
	ProfitShareInfoList []*OpenBankProfitShareInfo `json:"ProfitShareInfoList,omitempty" name:"ProfitShareInfoList"`

	// 订单标题。
	OrderSubject *string `json:"OrderSubject,omitempty" name:"OrderSubject"`

	// 商品信息。
	GoodsDetail *string `json:"GoodsDetail,omitempty" name:"GoodsDetail"`

	// 超时时间。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 支付成功回调地址。
	NotifyUrl *string `json:"NotifyUrl,omitempty" name:"NotifyUrl"`

	// 支付成功前端跳转URL。
	FrontUrl *string `json:"FrontUrl,omitempty" name:"FrontUrl"`

	// 订单附加信息，查询或者回调的时候原样返回。
	Attachment *string `json:"Attachment,omitempty" name:"Attachment"`

	// 第三方渠道扩展字段。见附录-复杂类型。
	// 未作特殊说明，则无需传入。
	ExternalPaymentData *string `json:"ExternalPaymentData,omitempty" name:"ExternalPaymentData"`

	// 备注。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 门店信息。
	StoreInfo *OpenBankStoreInfo `json:"StoreInfo,omitempty" name:"StoreInfo"`

	// 支付限制。
	PayLimitInfo *OpenBankPayLimitInfo `json:"PayLimitInfo,omitempty" name:"PayLimitInfo"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *CreateOpenBankUnifiedOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOpenBankUnifiedOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelMerchantId")
	delete(f, "ChannelName")
	delete(f, "PayType")
	delete(f, "OutOrderId")
	delete(f, "TotalAmount")
	delete(f, "Currency")
	delete(f, "ChannelSubMerchantId")
	delete(f, "PayChannel")
	delete(f, "SceneInfo")
	delete(f, "ProfitShareInfoList")
	delete(f, "OrderSubject")
	delete(f, "GoodsDetail")
	delete(f, "ExpireTime")
	delete(f, "NotifyUrl")
	delete(f, "FrontUrl")
	delete(f, "Attachment")
	delete(f, "ExternalPaymentData")
	delete(f, "Remark")
	delete(f, "StoreInfo")
	delete(f, "PayLimitInfo")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOpenBankUnifiedOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOpenBankUnifiedOrderResponseParams struct {
	// 业务系统返回码，SUCCESS表示成功，其他表示失败。
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 业务系统返回消息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 统一下单响应对象。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *CreateOpenBankOrderPaymentResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateOpenBankUnifiedOrderResponse struct {
	*tchttp.BaseResponse
	Response *CreateOpenBankUnifiedOrderResponseParams `json:"Response"`
}

func (r *CreateOpenBankUnifiedOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOpenBankUnifiedOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrderRequestParams struct {
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

// Predefined struct for user
type CreateOrderResponseParams struct {
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
}

type CreateOrderResponse struct {
	*tchttp.BaseResponse
	Response *CreateOrderResponseParams `json:"Response"`
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

// Predefined struct for user
type CreatePayMerchantRequestParams struct {
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

// Predefined struct for user
type CreatePayMerchantResponseParams struct {
	// 分配给商户的 AppId。该 AppId 为后续各项 交易的商户标识。
	MerchantAppId *string `json:"MerchantAppId,omitempty" name:"MerchantAppId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePayMerchantResponse struct {
	*tchttp.BaseResponse
	Response *CreatePayMerchantResponseParams `json:"Response"`
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

// Predefined struct for user
type CreatePayRollPreOrderRequestParams struct {
	// 用户在商户对应appid下的唯一标识
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 微信服务商下特约商户的商户号，由微信支付生成并下发
	SubMerchantId *string `json:"SubMerchantId,omitempty" name:"SubMerchantId"`

	// 商户系统内部的商家核身单号，要求此参数只能由数字、大小写字母组成，在服务商内部唯一
	AuthNumber *string `json:"AuthNumber,omitempty" name:"AuthNumber"`

	// 该劳务活动的项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 该工人所属的用工企业
	CompanyName *string `json:"CompanyName,omitempty" name:"CompanyName"`

	// 是服务商在微信申请公众号/小程序或移动应用成功后分配的账号ID（与服务商主体一致）
	// 当输入服务商Appid时，会校验其与服务商商户号的绑定关系。服务商APPID和与特约商户APPID至少输入一个，且必须要有拉起领薪卡小程序时使用的APPID
	WechatAppId *string `json:"WechatAppId,omitempty" name:"WechatAppId"`

	// 特约商户在微信申请公众号/小程序或移动应用成功后分配的账号ID（与特约商户主体一致）
	// 当输入特约商户Appid时，会校验其与特约商户号的绑定关系。服务商APPID和与特约商户APPID至少输入一个，且必须要有拉起领薪卡小程序时使用的APPID
	WechatSubAppId *string `json:"WechatSubAppId,omitempty" name:"WechatSubAppId"`
}

type CreatePayRollPreOrderRequest struct {
	*tchttp.BaseRequest
	
	// 用户在商户对应appid下的唯一标识
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 微信服务商下特约商户的商户号，由微信支付生成并下发
	SubMerchantId *string `json:"SubMerchantId,omitempty" name:"SubMerchantId"`

	// 商户系统内部的商家核身单号，要求此参数只能由数字、大小写字母组成，在服务商内部唯一
	AuthNumber *string `json:"AuthNumber,omitempty" name:"AuthNumber"`

	// 该劳务活动的项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 该工人所属的用工企业
	CompanyName *string `json:"CompanyName,omitempty" name:"CompanyName"`

	// 是服务商在微信申请公众号/小程序或移动应用成功后分配的账号ID（与服务商主体一致）
	// 当输入服务商Appid时，会校验其与服务商商户号的绑定关系。服务商APPID和与特约商户APPID至少输入一个，且必须要有拉起领薪卡小程序时使用的APPID
	WechatAppId *string `json:"WechatAppId,omitempty" name:"WechatAppId"`

	// 特约商户在微信申请公众号/小程序或移动应用成功后分配的账号ID（与特约商户主体一致）
	// 当输入特约商户Appid时，会校验其与特约商户号的绑定关系。服务商APPID和与特约商户APPID至少输入一个，且必须要有拉起领薪卡小程序时使用的APPID
	WechatSubAppId *string `json:"WechatSubAppId,omitempty" name:"WechatSubAppId"`
}

func (r *CreatePayRollPreOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePayRollPreOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OpenId")
	delete(f, "SubMerchantId")
	delete(f, "AuthNumber")
	delete(f, "ProjectName")
	delete(f, "CompanyName")
	delete(f, "WechatAppId")
	delete(f, "WechatSubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePayRollPreOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePayRollPreOrderResponseParams struct {
	// 商户系统内部的商家核身单号，要求此参数只能由数字、大小写字母组成，在服务商内部唯一
	AuthNumber *string `json:"AuthNumber,omitempty" name:"AuthNumber"`

	// Token有效时间，单位秒
	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 微信服务商商户的商户号，由微信支付生成并下发
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 用户在商户对应appid下的唯一标识
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 微信服务商下特约商户的商户号，由微信支付生成并下发
	SubMerchantId *string `json:"SubMerchantId,omitempty" name:"SubMerchantId"`

	// Token值
	Token *string `json:"Token,omitempty" name:"Token"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePayRollPreOrderResponse struct {
	*tchttp.BaseResponse
	Response *CreatePayRollPreOrderResponseParams `json:"Response"`
}

func (r *CreatePayRollPreOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePayRollPreOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePayRollPreOrderWithAuthRequestParams struct {
	// 用户在商户对应appid下的唯一标识
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 微信服务商下特约商户的商户号，由微信支付生成并下发
	SubMerchantId *string `json:"SubMerchantId,omitempty" name:"SubMerchantId"`

	// 商户系统内部的商家核身单号，要求此参数只能由数字、大小写字母组成，在服务商内部唯一
	AuthNumber *string `json:"AuthNumber,omitempty" name:"AuthNumber"`

	// 该劳务活动的项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 该工人所属的用工企业
	CompanyName *string `json:"CompanyName,omitempty" name:"CompanyName"`

	// 用户实名信息，该字段需进行加密处理，加密方法详见[敏感信息加密说明](https://pay.weixin.qq.com/wiki/doc/apiv3_partner/wechatpay/wechatpay4_3.shtml)
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 用户证件号，该字段需进行加密处理，加密方法详见[敏感信息加密说明](https://pay.weixin.qq.com/wiki/doc/apiv3_partner/wechatpay/wechatpay4_3.shtml)
	IdNo *string `json:"IdNo,omitempty" name:"IdNo"`

	// 微工卡服务仅支持用于与商户有用工关系的用户，需明确用工类型；参考值：
	// LONG_TERM_EMPLOYMENT：长期用工，
	// SHORT_TERM_EMPLOYMENT： 短期用工，
	// COOPERATION_EMPLOYMENT：合作关系
	EmploymentType *string `json:"EmploymentType,omitempty" name:"EmploymentType"`

	// 是服务商在微信申请公众号/小程序或移动应用成功后分配的账号ID（与服务商主体一致）
	// 当输入服务商Appid时，会校验其与服务商商户号的绑定关系。服务商APPID和与特约商户APPID至少输入一个，且必须要有拉起领薪卡小程序时使用的APPID
	WechatAppId *string `json:"WechatAppId,omitempty" name:"WechatAppId"`

	// 特约商户在微信申请公众号/小程序或移动应用成功后分配的账号ID（与特约商户主体一致）
	// 当输入特约商户Appid时，会校验其与特约商户号的绑定关系。服务商APPID和与特约商户APPID至少输入一个，且必须要有拉起领薪卡小程序时使用的APPID
	WechatSubAppId *string `json:"WechatSubAppId,omitempty" name:"WechatSubAppId"`
}

type CreatePayRollPreOrderWithAuthRequest struct {
	*tchttp.BaseRequest
	
	// 用户在商户对应appid下的唯一标识
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 微信服务商下特约商户的商户号，由微信支付生成并下发
	SubMerchantId *string `json:"SubMerchantId,omitempty" name:"SubMerchantId"`

	// 商户系统内部的商家核身单号，要求此参数只能由数字、大小写字母组成，在服务商内部唯一
	AuthNumber *string `json:"AuthNumber,omitempty" name:"AuthNumber"`

	// 该劳务活动的项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 该工人所属的用工企业
	CompanyName *string `json:"CompanyName,omitempty" name:"CompanyName"`

	// 用户实名信息，该字段需进行加密处理，加密方法详见[敏感信息加密说明](https://pay.weixin.qq.com/wiki/doc/apiv3_partner/wechatpay/wechatpay4_3.shtml)
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 用户证件号，该字段需进行加密处理，加密方法详见[敏感信息加密说明](https://pay.weixin.qq.com/wiki/doc/apiv3_partner/wechatpay/wechatpay4_3.shtml)
	IdNo *string `json:"IdNo,omitempty" name:"IdNo"`

	// 微工卡服务仅支持用于与商户有用工关系的用户，需明确用工类型；参考值：
	// LONG_TERM_EMPLOYMENT：长期用工，
	// SHORT_TERM_EMPLOYMENT： 短期用工，
	// COOPERATION_EMPLOYMENT：合作关系
	EmploymentType *string `json:"EmploymentType,omitempty" name:"EmploymentType"`

	// 是服务商在微信申请公众号/小程序或移动应用成功后分配的账号ID（与服务商主体一致）
	// 当输入服务商Appid时，会校验其与服务商商户号的绑定关系。服务商APPID和与特约商户APPID至少输入一个，且必须要有拉起领薪卡小程序时使用的APPID
	WechatAppId *string `json:"WechatAppId,omitempty" name:"WechatAppId"`

	// 特约商户在微信申请公众号/小程序或移动应用成功后分配的账号ID（与特约商户主体一致）
	// 当输入特约商户Appid时，会校验其与特约商户号的绑定关系。服务商APPID和与特约商户APPID至少输入一个，且必须要有拉起领薪卡小程序时使用的APPID
	WechatSubAppId *string `json:"WechatSubAppId,omitempty" name:"WechatSubAppId"`
}

func (r *CreatePayRollPreOrderWithAuthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePayRollPreOrderWithAuthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OpenId")
	delete(f, "SubMerchantId")
	delete(f, "AuthNumber")
	delete(f, "ProjectName")
	delete(f, "CompanyName")
	delete(f, "UserName")
	delete(f, "IdNo")
	delete(f, "EmploymentType")
	delete(f, "WechatAppId")
	delete(f, "WechatSubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePayRollPreOrderWithAuthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePayRollPreOrderWithAuthResponseParams struct {
	// 商户系统内部的商家核身单号，要求此参数只能由数字、大小写字母组成，在服务商内部唯一
	AuthNumber *string `json:"AuthNumber,omitempty" name:"AuthNumber"`

	// Token有效时间，单位秒
	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 微信服务商商户的商户号，由微信支付生成并下发
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 用户在商户对应appid下的唯一标识
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 微信服务商下特约商户的商户号，由微信支付生成并下发
	SubMerchantId *string `json:"SubMerchantId,omitempty" name:"SubMerchantId"`

	// Token值
	Token *string `json:"Token,omitempty" name:"Token"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePayRollPreOrderWithAuthResponse struct {
	*tchttp.BaseResponse
	Response *CreatePayRollPreOrderWithAuthResponseParams `json:"Response"`
}

func (r *CreatePayRollPreOrderWithAuthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePayRollPreOrderWithAuthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePayRollTokenRequestParams struct {
	// 用户在商户对应appid下的唯一标识
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 微信服务商下特约商户的商户号，由微信支付生成并下发
	SubMerchantId *string `json:"SubMerchantId,omitempty" name:"SubMerchantId"`

	// 用户实名信息，该字段需进行加密处理，加密方法详见[敏感信息加密说明](https://pay.weixin.qq.com/wiki/doc/apiv3_partner/wechatpay/wechatpay4_3.shtml)
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 用户证件号，该字段需进行加密处理，加密方法详见[敏感信息加密说明](https://pay.weixin.qq.com/wiki/doc/apiv3_partner/wechatpay/wechatpay4_3.shtml)
	IdNo *string `json:"IdNo,omitempty" name:"IdNo"`

	// 微工卡服务仅支持用于与商户有用工关系的用户，需明确用工类型；参考值：
	// LONG_TERM_EMPLOYMENT：长期用工，
	// SHORT_TERM_EMPLOYMENT： 短期用工，
	// COOPERATION_EMPLOYMENT：合作关系
	EmploymentType *string `json:"EmploymentType,omitempty" name:"EmploymentType"`

	// 是服务商在微信申请公众号/小程序或移动应用成功后分配的账号ID（与服务商主体一致）
	// 当输入服务商Appid时，会校验其与服务商商户号的绑定关系。服务商APPID和与特约商户APPID至少输入一个，且必须要有拉起领薪卡小程序时使用的APPID
	WechatAppId *string `json:"WechatAppId,omitempty" name:"WechatAppId"`

	// 特约商户在微信申请公众号/小程序或移动应用成功后分配的账号ID（与特约商户主体一致）
	// 当输入特约商户Appid时，会校验其与特约商户号的绑定关系。服务商APPID和与特约商户APPID至少输入一个，且必须要有拉起领薪卡小程序时使用的APPID
	WechatSubAppId *string `json:"WechatSubAppId,omitempty" name:"WechatSubAppId"`
}

type CreatePayRollTokenRequest struct {
	*tchttp.BaseRequest
	
	// 用户在商户对应appid下的唯一标识
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 微信服务商下特约商户的商户号，由微信支付生成并下发
	SubMerchantId *string `json:"SubMerchantId,omitempty" name:"SubMerchantId"`

	// 用户实名信息，该字段需进行加密处理，加密方法详见[敏感信息加密说明](https://pay.weixin.qq.com/wiki/doc/apiv3_partner/wechatpay/wechatpay4_3.shtml)
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 用户证件号，该字段需进行加密处理，加密方法详见[敏感信息加密说明](https://pay.weixin.qq.com/wiki/doc/apiv3_partner/wechatpay/wechatpay4_3.shtml)
	IdNo *string `json:"IdNo,omitempty" name:"IdNo"`

	// 微工卡服务仅支持用于与商户有用工关系的用户，需明确用工类型；参考值：
	// LONG_TERM_EMPLOYMENT：长期用工，
	// SHORT_TERM_EMPLOYMENT： 短期用工，
	// COOPERATION_EMPLOYMENT：合作关系
	EmploymentType *string `json:"EmploymentType,omitempty" name:"EmploymentType"`

	// 是服务商在微信申请公众号/小程序或移动应用成功后分配的账号ID（与服务商主体一致）
	// 当输入服务商Appid时，会校验其与服务商商户号的绑定关系。服务商APPID和与特约商户APPID至少输入一个，且必须要有拉起领薪卡小程序时使用的APPID
	WechatAppId *string `json:"WechatAppId,omitempty" name:"WechatAppId"`

	// 特约商户在微信申请公众号/小程序或移动应用成功后分配的账号ID（与特约商户主体一致）
	// 当输入特约商户Appid时，会校验其与特约商户号的绑定关系。服务商APPID和与特约商户APPID至少输入一个，且必须要有拉起领薪卡小程序时使用的APPID
	WechatSubAppId *string `json:"WechatSubAppId,omitempty" name:"WechatSubAppId"`
}

func (r *CreatePayRollTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePayRollTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OpenId")
	delete(f, "SubMerchantId")
	delete(f, "UserName")
	delete(f, "IdNo")
	delete(f, "EmploymentType")
	delete(f, "WechatAppId")
	delete(f, "WechatSubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePayRollTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePayRollTokenResponseParams struct {
	// Token有效时间，单位秒
	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 微信服务商商户的商户号，由微信支付生成并下发
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 用户在商户对应appid下的唯一标识
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 微信服务商下特约商户的商户号，由微信支付生成并下发
	SubMerchantId *string `json:"SubMerchantId,omitempty" name:"SubMerchantId"`

	// Token值
	Token *string `json:"Token,omitempty" name:"Token"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePayRollTokenResponse struct {
	*tchttp.BaseResponse
	Response *CreatePayRollTokenResponseParams `json:"Response"`
}

func (r *CreatePayRollTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePayRollTokenResponse) FromJsonString(s string) error {
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

// Predefined struct for user
type CreateRedInvoiceRequestParams struct {
	// 开票平台ID
	// 0 : 高灯
	// 1 : 票易通
	InvoicePlatformId *int64 `json:"InvoicePlatformId,omitempty" name:"InvoicePlatformId"`

	// 红冲明细
	Invoices []*CreateRedInvoiceItem `json:"Invoices,omitempty" name:"Invoices"`

	// 接入环境。沙箱环境填 sandbox。
	Profile *string `json:"Profile,omitempty" name:"Profile"`

	// 开票渠道。0：线上渠道，1：线下渠道。不填默认为线上渠道
	InvoiceChannel *int64 `json:"InvoiceChannel,omitempty" name:"InvoiceChannel"`
}

type CreateRedInvoiceRequest struct {
	*tchttp.BaseRequest
	
	// 开票平台ID
	// 0 : 高灯
	// 1 : 票易通
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

// Predefined struct for user
type CreateRedInvoiceResponseParams struct {
	// 红冲结果
	Result *CreateRedInvoiceResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateRedInvoiceResponse struct {
	*tchttp.BaseResponse
	Response *CreateRedInvoiceResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateRedInvoiceV2RequestParams struct {
	// 开票平台ID
	// 0 : 高灯
	// 1 : 票易通
	InvoicePlatformId *int64 `json:"InvoicePlatformId,omitempty" name:"InvoicePlatformId"`

	// 订单号
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 接入环境。沙箱环境填 sandbox。
	Profile *string `json:"Profile,omitempty" name:"Profile"`

	// 开票渠道。0：线上渠道，1：线下渠道。不填默认为线上渠道
	InvoiceChannel *int64 `json:"InvoiceChannel,omitempty" name:"InvoiceChannel"`
}

type CreateRedInvoiceV2Request struct {
	*tchttp.BaseRequest
	
	// 开票平台ID
	// 0 : 高灯
	// 1 : 票易通
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

// Predefined struct for user
type CreateRedInvoiceV2ResponseParams struct {
	// 红冲结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *CreateRedInvoiceResultV2 `json:"Result,omitempty" name:"Result"`

	// 错误码
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateRedInvoiceV2Response struct {
	*tchttp.BaseResponse
	Response *CreateRedInvoiceV2ResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateSinglePayRequestParams struct {
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

// Predefined struct for user
type CreateSinglePayResponseParams struct {
	// 返回结果
	Result *CreateSinglePayResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSinglePayResponse struct {
	*tchttp.BaseResponse
	Response *CreateSinglePayResponseParams `json:"Response"`
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

type CreateSinglePaymentData struct {
	// 平台交易流水号，唯一
	TradeSerialNo *string `json:"TradeSerialNo,omitempty" name:"TradeSerialNo"`
}

// Predefined struct for user
type CreateSinglePaymentRequestParams struct {
	// 转账类型
	TransferType *uint64 `json:"TransferType,omitempty" name:"TransferType"`

	// 订单流水号
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 转账金额
	TransferAmount *uint64 `json:"TransferAmount,omitempty" name:"TransferAmount"`

	// 主播ID（与主播业务ID不能同时为空，两者都填取主播ID）
	AnchorId *string `json:"AnchorId,omitempty" name:"AnchorId"`

	// 请求预留字段，原样透传返回
	ReqReserved *string `json:"ReqReserved,omitempty" name:"ReqReserved"`

	// 业务备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 主播名称。如果该字段填入，则会对AnchorName和AnchorId/Uid进行校验。
	AnchorName *string `json:"AnchorName,omitempty" name:"AnchorName"`

	// 主播业务ID（与主播ID不能同时为空，两者都填取主播ID）
	Uid *string `json:"Uid,omitempty" name:"Uid"`

	// 转账结果回调通知URL。若不填，则不进行回调。
	NotifyUrl *string `json:"NotifyUrl,omitempty" name:"NotifyUrl"`
}

type CreateSinglePaymentRequest struct {
	*tchttp.BaseRequest
	
	// 转账类型
	TransferType *uint64 `json:"TransferType,omitempty" name:"TransferType"`

	// 订单流水号
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 转账金额
	TransferAmount *uint64 `json:"TransferAmount,omitempty" name:"TransferAmount"`

	// 主播ID（与主播业务ID不能同时为空，两者都填取主播ID）
	AnchorId *string `json:"AnchorId,omitempty" name:"AnchorId"`

	// 请求预留字段，原样透传返回
	ReqReserved *string `json:"ReqReserved,omitempty" name:"ReqReserved"`

	// 业务备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 主播名称。如果该字段填入，则会对AnchorName和AnchorId/Uid进行校验。
	AnchorName *string `json:"AnchorName,omitempty" name:"AnchorName"`

	// 主播业务ID（与主播ID不能同时为空，两者都填取主播ID）
	Uid *string `json:"Uid,omitempty" name:"Uid"`

	// 转账结果回调通知URL。若不填，则不进行回调。
	NotifyUrl *string `json:"NotifyUrl,omitempty" name:"NotifyUrl"`
}

func (r *CreateSinglePaymentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSinglePaymentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TransferType")
	delete(f, "OrderId")
	delete(f, "TransferAmount")
	delete(f, "AnchorId")
	delete(f, "ReqReserved")
	delete(f, "Remark")
	delete(f, "AnchorName")
	delete(f, "Uid")
	delete(f, "NotifyUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSinglePaymentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSinglePaymentResponseParams struct {
	// 错误码
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 响应消息
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *CreateSinglePaymentData `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSinglePaymentResponse struct {
	*tchttp.BaseResponse
	Response *CreateSinglePaymentResponseParams `json:"Response"`
}

func (r *CreateSinglePaymentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSinglePaymentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTransferBatchRequestParams struct {
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

// Predefined struct for user
type CreateTransferBatchResponseParams struct {
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
}

type CreateTransferBatchResponse struct {
	*tchttp.BaseResponse
	Response *CreateTransferBatchResponseParams `json:"Response"`
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

// Predefined struct for user
type DeduceQuotaRequestParams struct {
	// 主播ID
	AnchorId *string `json:"AnchorId,omitempty" name:"AnchorId"`

	// 提现金额，单位为"分"
	Amount *int64 `json:"Amount,omitempty" name:"Amount"`

	// 外部业务订单号
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`
}

type DeduceQuotaRequest struct {
	*tchttp.BaseRequest
	
	// 主播ID
	AnchorId *string `json:"AnchorId,omitempty" name:"AnchorId"`

	// 提现金额，单位为"分"
	Amount *int64 `json:"Amount,omitempty" name:"Amount"`

	// 外部业务订单号
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`
}

func (r *DeduceQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeduceQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AnchorId")
	delete(f, "Amount")
	delete(f, "OrderId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeduceQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeduceQuotaResponseParams struct {
	// 错误码。响应成功："SUCCESS"，其他为不成功。
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 响应消息
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 返回响应
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *AssignmentData `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeduceQuotaResponse struct {
	*tchttp.BaseResponse
	Response *DeduceQuotaResponseParams `json:"Response"`
}

func (r *DeduceQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeduceQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAgentTaxPaymentInfoRequestParams struct {
	// 批次号
	BatchNum *int64 `json:"BatchNum,omitempty" name:"BatchNum"`

	// 接入环境。沙箱环境填sandbox
	Profile *string `json:"Profile,omitempty" name:"Profile"`
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

// Predefined struct for user
type DeleteAgentTaxPaymentInfoResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteAgentTaxPaymentInfoResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAgentTaxPaymentInfoResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteAgentTaxPaymentInfosRequestParams struct {
	// 批次号
	BatchNum *int64 `json:"BatchNum,omitempty" name:"BatchNum"`
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

// Predefined struct for user
type DeleteAgentTaxPaymentInfosResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteAgentTaxPaymentInfosResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAgentTaxPaymentInfosResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeChargeDetailRequestParams struct {
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

// Predefined struct for user
type DescribeChargeDetailResponseParams struct {
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
}

type DescribeChargeDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeChargeDetailResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeOrderStatusRequestParams struct {
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

// Predefined struct for user
type DescribeOrderStatusResponseParams struct {
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
}

type DescribeOrderStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrderStatusResponseParams `json:"Response"`
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

// Predefined struct for user
type DistributeAccreditQueryRequestParams struct {
	// 使用门店OpenId
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 使用门店OpenKey
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

type DistributeAccreditQueryRequest struct {
	*tchttp.BaseRequest
	
	// 使用门店OpenId
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 使用门店OpenKey
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *DistributeAccreditQueryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DistributeAccreditQueryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OpenId")
	delete(f, "OpenKey")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DistributeAccreditQueryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DistributeAccreditQueryResponseParams struct {
	// 业务系统返回消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 业务系统返回码
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 查询授权申请结果响应对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *DistributeAccreditQueryResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DistributeAccreditQueryResponse struct {
	*tchttp.BaseResponse
	Response *DistributeAccreditQueryResponseParams `json:"Response"`
}

func (r *DistributeAccreditQueryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DistributeAccreditQueryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DistributeAccreditQueryResult struct {
	// 状态（0-未开通，1-已开通，2-商户主动关闭，3-待审核，4-冻结，5-注销，6-待签合同）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 合同h5地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContractUrl *string `json:"ContractUrl,omitempty" name:"ContractUrl"`

	// 说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type DistributeAccreditResult struct {
	// 合同h5地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContractUrl *string `json:"ContractUrl,omitempty" name:"ContractUrl"`

	// 系统商户号
	// 注意：此字段可能返回 null，表示取不到有效值。
	MerchantNo *string `json:"MerchantNo,omitempty" name:"MerchantNo"`
}

// Predefined struct for user
type DistributeAccreditTlinxRequestParams struct {
	// 使用门店OpenId
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 使用门店OpenKey
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 验证方式，传1手机验证(验证码时效60S)传2结算卡验证(时效6小时)，多种方式用逗号隔开
	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`

	// 分账比例（500=5%）不传默认百分之10
	Percent *string `json:"Percent,omitempty" name:"Percent"`

	// 营业执照商户全称
	FullName *string `json:"FullName,omitempty" name:"FullName"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

type DistributeAccreditTlinxRequest struct {
	*tchttp.BaseRequest
	
	// 使用门店OpenId
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 使用门店OpenKey
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 验证方式，传1手机验证(验证码时效60S)传2结算卡验证(时效6小时)，多种方式用逗号隔开
	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`

	// 分账比例（500=5%）不传默认百分之10
	Percent *string `json:"Percent,omitempty" name:"Percent"`

	// 营业执照商户全称
	FullName *string `json:"FullName,omitempty" name:"FullName"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *DistributeAccreditTlinxRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DistributeAccreditTlinxRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OpenId")
	delete(f, "OpenKey")
	delete(f, "AuthType")
	delete(f, "Percent")
	delete(f, "FullName")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DistributeAccreditTlinxRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DistributeAccreditTlinxResponseParams struct {
	// 业务系统返回消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 业务系统返回码
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 授权申请响应对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *DistributeAccreditResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DistributeAccreditTlinxResponse struct {
	*tchttp.BaseResponse
	Response *DistributeAccreditTlinxResponseParams `json:"Response"`
}

func (r *DistributeAccreditTlinxResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DistributeAccreditTlinxResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DistributeAddReceiverRequestParams struct {
	// 使用门店OpenId
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 使用门店OpenKey
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 商户编号
	MerchantNo *string `json:"MerchantNo,omitempty" name:"MerchantNo"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

type DistributeAddReceiverRequest struct {
	*tchttp.BaseRequest
	
	// 使用门店OpenId
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 使用门店OpenKey
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 商户编号
	MerchantNo *string `json:"MerchantNo,omitempty" name:"MerchantNo"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *DistributeAddReceiverRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DistributeAddReceiverRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OpenId")
	delete(f, "OpenKey")
	delete(f, "MerchantNo")
	delete(f, "Remark")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DistributeAddReceiverRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DistributeAddReceiverResponseParams struct {
	// 业务系统返回消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 业务系统返回码
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 添加分账接收方响应对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *DistributeReceiverResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DistributeAddReceiverResponse struct {
	*tchttp.BaseResponse
	Response *DistributeAddReceiverResponseParams `json:"Response"`
}

func (r *DistributeAddReceiverResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DistributeAddReceiverResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DistributeApplyRequestParams struct {
	// 使用门店OpenId
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 使用门店OpenKey
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 商户分账单号
	OutDistributeNo *string `json:"OutDistributeNo,omitempty" name:"OutDistributeNo"`

	// 分账明细
	Details []*MultiApplyDetail `json:"Details,omitempty" name:"Details"`

	// 商户交易订单号，和OrderNo二者传其一
	DeveloperNo *string `json:"DeveloperNo,omitempty" name:"DeveloperNo"`

	// 平台交易订单号，和DeveloperNo二者传其一
	OrderNo *string `json:"OrderNo,omitempty" name:"OrderNo"`

	// 说明
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

type DistributeApplyRequest struct {
	*tchttp.BaseRequest
	
	// 使用门店OpenId
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 使用门店OpenKey
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 商户分账单号
	OutDistributeNo *string `json:"OutDistributeNo,omitempty" name:"OutDistributeNo"`

	// 分账明细
	Details []*MultiApplyDetail `json:"Details,omitempty" name:"Details"`

	// 商户交易订单号，和OrderNo二者传其一
	DeveloperNo *string `json:"DeveloperNo,omitempty" name:"DeveloperNo"`

	// 平台交易订单号，和DeveloperNo二者传其一
	OrderNo *string `json:"OrderNo,omitempty" name:"OrderNo"`

	// 说明
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *DistributeApplyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DistributeApplyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OpenId")
	delete(f, "OpenKey")
	delete(f, "OutDistributeNo")
	delete(f, "Details")
	delete(f, "DeveloperNo")
	delete(f, "OrderNo")
	delete(f, "Remark")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DistributeApplyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DistributeApplyResponseParams struct {
	// 业务系统返回消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 业务系统返回码
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 分账申请响应对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *DistributeMultiApplyResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DistributeApplyResponse struct {
	*tchttp.BaseResponse
	Response *DistributeApplyResponseParams `json:"Response"`
}

func (r *DistributeApplyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DistributeApplyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DistributeCancelRequestParams struct {
	// 使用门店OpenId
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 使用门店OpenKey
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 平台交易订单号
	OrderNo *string `json:"OrderNo,omitempty" name:"OrderNo"`

	// 商户分账单号，type为2时，和DistributeNo二者传其一
	OutDistributeNo *string `json:"OutDistributeNo,omitempty" name:"OutDistributeNo"`

	// 平台分账单号，type为2时，和OutDistributeNo二者传其一
	DistributeNo *string `json:"DistributeNo,omitempty" name:"DistributeNo"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

type DistributeCancelRequest struct {
	*tchttp.BaseRequest
	
	// 使用门店OpenId
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 使用门店OpenKey
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 平台交易订单号
	OrderNo *string `json:"OrderNo,omitempty" name:"OrderNo"`

	// 商户分账单号，type为2时，和DistributeNo二者传其一
	OutDistributeNo *string `json:"OutDistributeNo,omitempty" name:"OutDistributeNo"`

	// 平台分账单号，type为2时，和OutDistributeNo二者传其一
	DistributeNo *string `json:"DistributeNo,omitempty" name:"DistributeNo"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *DistributeCancelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DistributeCancelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OpenId")
	delete(f, "OpenKey")
	delete(f, "OrderNo")
	delete(f, "OutDistributeNo")
	delete(f, "DistributeNo")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DistributeCancelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DistributeCancelResponseParams struct {
	// 业务系统返回消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 业务系统返回码
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 分账撤销响应对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *DistributeCancelResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DistributeCancelResponse struct {
	*tchttp.BaseResponse
	Response *DistributeCancelResponseParams `json:"Response"`
}

func (r *DistributeCancelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DistributeCancelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DistributeCancelResult struct {
	// 分账订单状态（0初始1成功2失败3撤销）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 平台交易订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderNo *string `json:"OrderNo,omitempty" name:"OrderNo"`

	// 商户分账单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutDistributeNo *string `json:"OutDistributeNo,omitempty" name:"OutDistributeNo"`

	// 平台分账单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	DistributeNo *string `json:"DistributeNo,omitempty" name:"DistributeNo"`
}

type DistributeMultiApplyResult struct {
	// 分账状态（0分账初始 1分账成功 2分账失败）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 平台分账单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	DistributeNo *string `json:"DistributeNo,omitempty" name:"DistributeNo"`

	// 入账日期，yyyy-MM-dd格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	InDate *string `json:"InDate,omitempty" name:"InDate"`

	// 分账金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	Amount *string `json:"Amount,omitempty" name:"Amount"`

	// 商户分账单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutDistributeNo *string `json:"OutDistributeNo,omitempty" name:"OutDistributeNo"`

	// 平台支付单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderNo *string `json:"OrderNo,omitempty" name:"OrderNo"`
}

// Predefined struct for user
type DistributeQueryReceiverRequestParams struct {
	// 使用门店OpenId
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 使用门店OpenKey
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

type DistributeQueryReceiverRequest struct {
	*tchttp.BaseRequest
	
	// 使用门店OpenId
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 使用门店OpenKey
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *DistributeQueryReceiverRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DistributeQueryReceiverRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OpenId")
	delete(f, "OpenKey")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DistributeQueryReceiverRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DistributeQueryReceiverResponseParams struct {
	// 业务系统返回消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 业务系统返回码
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 查询询已添加分账接收方响应对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *DistributeReceiverResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DistributeQueryReceiverResponse struct {
	*tchttp.BaseResponse
	Response *DistributeQueryReceiverResponseParams `json:"Response"`
}

func (r *DistributeQueryReceiverResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DistributeQueryReceiverResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DistributeQueryRequestParams struct {
	// 使用门店OpenId
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 使用门店OpenKey
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 查询类型（1-全部，2-单笔）
	Type *string `json:"Type,omitempty" name:"Type"`

	// 商户分账单号，type为2时，和DistributeNo二者传其一
	OutDistributeNo *string `json:"OutDistributeNo,omitempty" name:"OutDistributeNo"`

	// 平台分账单号，type为2时，和OutDistributeNo二者传其一
	DistributeNo *string `json:"DistributeNo,omitempty" name:"DistributeNo"`

	// 平台交易订单号
	OrderNo *string `json:"OrderNo,omitempty" name:"OrderNo"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

type DistributeQueryRequest struct {
	*tchttp.BaseRequest
	
	// 使用门店OpenId
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 使用门店OpenKey
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 查询类型（1-全部，2-单笔）
	Type *string `json:"Type,omitempty" name:"Type"`

	// 商户分账单号，type为2时，和DistributeNo二者传其一
	OutDistributeNo *string `json:"OutDistributeNo,omitempty" name:"OutDistributeNo"`

	// 平台分账单号，type为2时，和OutDistributeNo二者传其一
	DistributeNo *string `json:"DistributeNo,omitempty" name:"DistributeNo"`

	// 平台交易订单号
	OrderNo *string `json:"OrderNo,omitempty" name:"OrderNo"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *DistributeQueryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DistributeQueryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OpenId")
	delete(f, "OpenKey")
	delete(f, "Type")
	delete(f, "OutDistributeNo")
	delete(f, "DistributeNo")
	delete(f, "OrderNo")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DistributeQueryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DistributeQueryResponseParams struct {
	// 业务系统返回消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 业务系统返回码
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 分账结果响应对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *DistributeQueryResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DistributeQueryResponse struct {
	*tchttp.BaseResponse
	Response *DistributeQueryResponseParams `json:"Response"`
}

func (r *DistributeQueryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DistributeQueryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DistributeQueryResult struct {
	// 分账订单列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Orders []*MultiApplyOrder `json:"Orders,omitempty" name:"Orders"`
}

type DistributeReceiverResult struct {
	// 商户编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	MerchantNo *string `json:"MerchantNo,omitempty" name:"MerchantNo"`
}

// Predefined struct for user
type DistributeRemoveReceiverRequestParams struct {
	// 使用门店OpenId
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 使用门店OpenKey
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 商户编号
	MerchantNo *string `json:"MerchantNo,omitempty" name:"MerchantNo"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

type DistributeRemoveReceiverRequest struct {
	*tchttp.BaseRequest
	
	// 使用门店OpenId
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 使用门店OpenKey
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 商户编号
	MerchantNo *string `json:"MerchantNo,omitempty" name:"MerchantNo"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *DistributeRemoveReceiverRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DistributeRemoveReceiverRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OpenId")
	delete(f, "OpenKey")
	delete(f, "MerchantNo")
	delete(f, "Remark")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DistributeRemoveReceiverRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DistributeRemoveReceiverResponseParams struct {
	// 业务系统返回消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 业务系统返回码
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 解除分账接收方响应对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *DistributeReceiverResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DistributeRemoveReceiverResponse struct {
	*tchttp.BaseResponse
	Response *DistributeRemoveReceiverResponseParams `json:"Response"`
}

func (r *DistributeRemoveReceiverResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DistributeRemoveReceiverResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadBillRequestParams struct {
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

// Predefined struct for user
type DownloadBillResponseParams struct {
	// 账单文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 账单文件的MD5值
	FileMD5 *string `json:"FileMD5,omitempty" name:"FileMD5"`

	// 账单文件的真实下载地址
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DownloadBillResponse struct {
	*tchttp.BaseResponse
	Response *DownloadBillResponseParams `json:"Response"`
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

type DownloadFileResult struct {
	// 文件内容（base64加密的二进制内容）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitempty" name:"Content"`

	// 存储区域（0私密区，1公共区），请严格按文件要求，上传到不同的区域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Storage *string `json:"Storage,omitempty" name:"Storage"`
}

// Predefined struct for user
type DownloadOrgFileRequestParams struct {
	// 收单系统分配的开放ID
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 收单系统分配的密钥
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 存储区域（0私密区，1公共区），请严格按文件要求，上传到不同的区域
	Storage *string `json:"Storage,omitempty" name:"Storage"`

	// 文件路径
	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

type DownloadOrgFileRequest struct {
	*tchttp.BaseRequest
	
	// 收单系统分配的开放ID
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 收单系统分配的密钥
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 存储区域（0私密区，1公共区），请严格按文件要求，上传到不同的区域
	Storage *string `json:"Storage,omitempty" name:"Storage"`

	// 文件路径
	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *DownloadOrgFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadOrgFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OpenId")
	delete(f, "OpenKey")
	delete(f, "Storage")
	delete(f, "FilePath")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DownloadOrgFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadOrgFileResponseParams struct {
	// 业务系统返回消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 业务系统返回码
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 下载机构文件响应对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *DownloadFileResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DownloadOrgFileResponse struct {
	*tchttp.BaseResponse
	Response *DownloadOrgFileResponseParams `json:"Response"`
}

func (r *DownloadOrgFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadOrgFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadReconciliationUrlRequestParams struct {
	// 平台应用ID
	MainAppId *string `json:"MainAppId,omitempty" name:"MainAppId"`

	// 平台代码
	AppCode *string `json:"AppCode,omitempty" name:"AppCode"`

	// 账单日期，yyyy-MM-dd
	BillDate *string `json:"BillDate,omitempty" name:"BillDate"`

	// 商户或者代理商ID
	SubAppId *string `json:"SubAppId,omitempty" name:"SubAppId"`
}

type DownloadReconciliationUrlRequest struct {
	*tchttp.BaseRequest
	
	// 平台应用ID
	MainAppId *string `json:"MainAppId,omitempty" name:"MainAppId"`

	// 平台代码
	AppCode *string `json:"AppCode,omitempty" name:"AppCode"`

	// 账单日期，yyyy-MM-dd
	BillDate *string `json:"BillDate,omitempty" name:"BillDate"`

	// 商户或者代理商ID
	SubAppId *string `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DownloadReconciliationUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadReconciliationUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MainAppId")
	delete(f, "AppCode")
	delete(f, "BillDate")
	delete(f, "SubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DownloadReconciliationUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadReconciliationUrlResponseParams struct {
	// 下载地址
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// hash类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	HashType *string `json:"HashType,omitempty" name:"HashType"`

	// hash值
	HashValue *string `json:"HashValue,omitempty" name:"HashValue"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DownloadReconciliationUrlResponse struct {
	*tchttp.BaseResponse
	Response *DownloadReconciliationUrlResponseParams `json:"Response"`
}

func (r *DownloadReconciliationUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadReconciliationUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExecuteMemberTransactionRequestParams struct {
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

// Predefined struct for user
type ExecuteMemberTransactionResponseParams struct {
	// 请求类型
	RequestType *string `json:"RequestType,omitempty" name:"RequestType"`

	// 银行流水号
	FrontSequenceNumber *string `json:"FrontSequenceNumber,omitempty" name:"FrontSequenceNumber"`

	// 保留域
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReservedMessage *string `json:"ReservedMessage,omitempty" name:"ReservedMessage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ExecuteMemberTransactionResponse struct {
	*tchttp.BaseResponse
	Response *ExecuteMemberTransactionResponseParams `json:"Response"`
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

type FeeRangInfo struct {
	// 卡类型，银联产品使用 
	// DEBIT：借记卡 
	// CREDIT：贷记卡
	CardType *string `json:"CardType,omitempty" name:"CardType"`

	// 区间起始金额，单位（分）
	RangeStartValue *uint64 `json:"RangeStartValue,omitempty" name:"RangeStartValue"`

	// 区间结束金额，单位（分）
	RangeEndValue *uint64 `json:"RangeEndValue,omitempty" name:"RangeEndValue"`

	// 分段计费模式 
	// SINGLE：按金额计费 
	// RATIO：按费率计费
	RangeFeeMode *string `json:"RangeFeeMode,omitempty" name:"RangeFeeMode"`

	// 费用值，单位（0.01%或分）
	FeeValue *uint64 `json:"FeeValue,omitempty" name:"FeeValue"`

	// 最低收费金额，单位（分）
	MinFee *uint64 `json:"MinFee,omitempty" name:"MinFee"`

	// 最高收费金额，单位（分）
	MaxFee *uint64 `json:"MaxFee,omitempty" name:"MaxFee"`
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

type FlexBillDownloadUrlResult struct {
	// 对账单文件下载链接
	Url *string `json:"Url,omitempty" name:"Url"`

	// 下载链接过期时间
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

type FlexFundingAccountInfo struct {
	// 资金账户号
	FundingAccountNo *string `json:"FundingAccountNo,omitempty" name:"FundingAccountNo"`

	// 资金账户类型
	FundingAccountType *string `json:"FundingAccountType,omitempty" name:"FundingAccountType"`

	// 资金账户绑定序列号
	FundingAccountBindSerialNo *string `json:"FundingAccountBindSerialNo,omitempty" name:"FundingAccountBindSerialNo"`
}

// Predefined struct for user
type FreezeFlexBalanceRequestParams struct {
	// 收款用户ID
	PayeeId *string `json:"PayeeId,omitempty" name:"PayeeId"`

	// 税前金额
	AmountBeforeTax *string `json:"AmountBeforeTax,omitempty" name:"AmountBeforeTax"`

	// 收入类型
	// LABOR:劳务所得
	// OCCASION:偶然所得
	IncomeType *string `json:"IncomeType,omitempty" name:"IncomeType"`

	// 外部订单ID
	OutOrderId *string `json:"OutOrderId,omitempty" name:"OutOrderId"`

	// 操作类型
	// FREEZE:冻结
	// UNFREEZE:解冻
	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`

	// 冻结备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// __test__:测试环境
	// 缺省默认为生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type FreezeFlexBalanceRequest struct {
	*tchttp.BaseRequest
	
	// 收款用户ID
	PayeeId *string `json:"PayeeId,omitempty" name:"PayeeId"`

	// 税前金额
	AmountBeforeTax *string `json:"AmountBeforeTax,omitempty" name:"AmountBeforeTax"`

	// 收入类型
	// LABOR:劳务所得
	// OCCASION:偶然所得
	IncomeType *string `json:"IncomeType,omitempty" name:"IncomeType"`

	// 外部订单ID
	OutOrderId *string `json:"OutOrderId,omitempty" name:"OutOrderId"`

	// 操作类型
	// FREEZE:冻结
	// UNFREEZE:解冻
	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`

	// 冻结备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// __test__:测试环境
	// 缺省默认为生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *FreezeFlexBalanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FreezeFlexBalanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PayeeId")
	delete(f, "AmountBeforeTax")
	delete(f, "IncomeType")
	delete(f, "OutOrderId")
	delete(f, "OperationType")
	delete(f, "Remark")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FreezeFlexBalanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FreezeFlexBalanceResponseParams struct {
	// 错误码。SUCCESS为成功，其他为失败
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *FreezeFlexBalanceResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type FreezeFlexBalanceResponse struct {
	*tchttp.BaseResponse
	Response *FreezeFlexBalanceResponseParams `json:"Response"`
}

func (r *FreezeFlexBalanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FreezeFlexBalanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FreezeFlexBalanceResult struct {
	// 冻结订单ID
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`
}

type FreezeOrderResult struct {
	// 税前金额
	AmountBeforeTax *string `json:"AmountBeforeTax,omitempty" name:"AmountBeforeTax"`

	// 收入类型
	// LABOR:劳务所得
	// OCCASION:偶然所得
	IncomeType *string `json:"IncomeType,omitempty" name:"IncomeType"`

	// 外部订单ID
	OutOrderId *string `json:"OutOrderId,omitempty" name:"OutOrderId"`

	// 订单ID
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 操作类型
	// FREEZE:冻结
	// UNFREEZE:解冻
	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`

	// 发起时间
	InitiateTime *string `json:"InitiateTime,omitempty" name:"InitiateTime"`

	// 完成时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`

	// 状态
	// ACCEPTED:已受理
	// ACCOUNTED:已记账
	// SUCCEED:已成功
	// FAILED:已失败
	Status *string `json:"Status,omitempty" name:"Status"`

	// 状态描述
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

	// 冻结备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type FreezeOrders struct {
	// 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*FreezeOrderResult `json:"List,omitempty" name:"List"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type FundsTransactionItem struct {
	// 资金交易类型。
	// __1__：提现/退款
	// __2__：清分/充值
	TransType *string `json:"TransType,omitempty" name:"TransType"`

	// 银行记账说明。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BankBookingMessage *string `json:"BankBookingMessage,omitempty" name:"BankBookingMessage"`

	// 交易状态。
	// __0__：成功
	TranStatus *string `json:"TranStatus,omitempty" name:"TranStatus"`

	// 业务方会员标识。
	// _平安渠道为交易网会员代码_
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransNetMemberCode *string `json:"TransNetMemberCode,omitempty" name:"TransNetMemberCode"`

	// 子账户账号。
	// _平安渠道为见证子账户的账号_
	SubAccountNumber *string `json:"SubAccountNumber,omitempty" name:"SubAccountNumber"`

	// 子账户名称。
	// _平安渠道为见证子账户的户名_
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubAccountName *string `json:"SubAccountName,omitempty" name:"SubAccountName"`

	// 交易金额。
	TransAmount *string `json:"TransAmount,omitempty" name:"TransAmount"`

	// 交易手续费。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransFee *string `json:"TransFee,omitempty" name:"TransFee"`

	// 交易日期，格式：yyyyMMdd。
	TransDate *string `json:"TransDate,omitempty" name:"TransDate"`

	// 交易时间，格式：HHmmss。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransTime *string `json:"TransTime,omitempty" name:"TransTime"`

	// 银行系统流水号。
	// _平安渠道为见证系统流水号_
	BankSequenceNumber *string `json:"BankSequenceNumber,omitempty" name:"BankSequenceNumber"`

	// 备注。
	// _平安渠道，如果是见证+收单的交易，返回交易订单号_
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

// Predefined struct for user
type GetBillDownloadUrlRequestParams struct {
	// 收单系统分配的开放ID
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 收单系统分配的密钥
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 清算日期（YYYYMMDD，今天传昨天的日期，每日下午1点后出前一日的账单）
	Day *string `json:"Day,omitempty" name:"Day"`
}

type GetBillDownloadUrlRequest struct {
	*tchttp.BaseRequest
	
	// 收单系统分配的开放ID
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 收单系统分配的密钥
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 清算日期（YYYYMMDD，今天传昨天的日期，每日下午1点后出前一日的账单）
	Day *string `json:"Day,omitempty" name:"Day"`
}

func (r *GetBillDownloadUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetBillDownloadUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OpenId")
	delete(f, "OpenKey")
	delete(f, "Day")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetBillDownloadUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetBillDownloadUrlResponseParams struct {
	// 业务系统返回码
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 业务系统返回消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 账单文件下载地址响应对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *BillDownloadUrlResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetBillDownloadUrlResponse struct {
	*tchttp.BaseResponse
	Response *GetBillDownloadUrlResponseParams `json:"Response"`
}

func (r *GetBillDownloadUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetBillDownloadUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDistributeBillDownloadUrlRequestParams struct {
	// 收单系统分配的开放ID
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 收单系统分配的密钥
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 分账日期（YYYYMMDD，今天传昨天的日期）
	Day *string `json:"Day,omitempty" name:"Day"`
}

type GetDistributeBillDownloadUrlRequest struct {
	*tchttp.BaseRequest
	
	// 收单系统分配的开放ID
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 收单系统分配的密钥
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 分账日期（YYYYMMDD，今天传昨天的日期）
	Day *string `json:"Day,omitempty" name:"Day"`
}

func (r *GetDistributeBillDownloadUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDistributeBillDownloadUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OpenId")
	delete(f, "OpenKey")
	delete(f, "Day")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDistributeBillDownloadUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDistributeBillDownloadUrlResponseParams struct {
	// 业务系统返回码
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 业务系统返回消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 账单文件下载地址响应对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *BillDownloadUrlResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetDistributeBillDownloadUrlResponse struct {
	*tchttp.BaseResponse
	Response *GetDistributeBillDownloadUrlResponseParams `json:"Response"`
}

func (r *GetDistributeBillDownloadUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDistributeBillDownloadUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPayRollAuthListRequestParams struct {
	// 用户在商户对应appid下的唯一标识
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 微信服务商下特约商户的商户号，由微信支付生成并下发
	SubMerchantId *string `json:"SubMerchantId,omitempty" name:"SubMerchantId"`

	// 核身日期，一次只能查询一天，最久可查询90天内的记录，格式为YYYY-MM-DD
	AuthDate *string `json:"AuthDate,omitempty" name:"AuthDate"`

	// 非负整数，表示该次请求资源的起始位置，从0开始计数
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 非0非负的整数，该次请求可返回的最大资源条数，默认值为10，最大支持10条
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 是服务商在微信申请公众号/小程序或移动应用成功后分配的账号ID（与服务商主体一致）
	// 当输入服务商Appid时，会校验其与服务商商户号的绑定关系。服务商APPID和与特约商户APPID至少输入一个，且必须要有拉起领薪卡小程序时使用的APPID
	WechatAppId *string `json:"WechatAppId,omitempty" name:"WechatAppId"`

	// 特约商户在微信申请公众号/小程序或移动应用成功后分配的账号ID（与特约商户主体一致）
	// 当输入特约商户Appid时，会校验其与特约商户号的绑定关系。服务商APPID和与特约商户APPID至少输入一个，且必须要有拉起领薪卡小程序时使用的APPID
	WechatSubAppId *string `json:"WechatSubAppId,omitempty" name:"WechatSubAppId"`

	// 核身状态，列表查询仅提供成功状态的核身记录查询，故此字段固定AUTHENTICATE_SUCCESS即可
	AuthStatus *string `json:"AuthStatus,omitempty" name:"AuthStatus"`
}

type GetPayRollAuthListRequest struct {
	*tchttp.BaseRequest
	
	// 用户在商户对应appid下的唯一标识
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 微信服务商下特约商户的商户号，由微信支付生成并下发
	SubMerchantId *string `json:"SubMerchantId,omitempty" name:"SubMerchantId"`

	// 核身日期，一次只能查询一天，最久可查询90天内的记录，格式为YYYY-MM-DD
	AuthDate *string `json:"AuthDate,omitempty" name:"AuthDate"`

	// 非负整数，表示该次请求资源的起始位置，从0开始计数
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 非0非负的整数，该次请求可返回的最大资源条数，默认值为10，最大支持10条
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 是服务商在微信申请公众号/小程序或移动应用成功后分配的账号ID（与服务商主体一致）
	// 当输入服务商Appid时，会校验其与服务商商户号的绑定关系。服务商APPID和与特约商户APPID至少输入一个，且必须要有拉起领薪卡小程序时使用的APPID
	WechatAppId *string `json:"WechatAppId,omitempty" name:"WechatAppId"`

	// 特约商户在微信申请公众号/小程序或移动应用成功后分配的账号ID（与特约商户主体一致）
	// 当输入特约商户Appid时，会校验其与特约商户号的绑定关系。服务商APPID和与特约商户APPID至少输入一个，且必须要有拉起领薪卡小程序时使用的APPID
	WechatSubAppId *string `json:"WechatSubAppId,omitempty" name:"WechatSubAppId"`

	// 核身状态，列表查询仅提供成功状态的核身记录查询，故此字段固定AUTHENTICATE_SUCCESS即可
	AuthStatus *string `json:"AuthStatus,omitempty" name:"AuthStatus"`
}

func (r *GetPayRollAuthListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPayRollAuthListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OpenId")
	delete(f, "SubMerchantId")
	delete(f, "AuthDate")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "WechatAppId")
	delete(f, "WechatSubAppId")
	delete(f, "AuthStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetPayRollAuthListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPayRollAuthListResponseParams struct {
	// 核身结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Results []*PayRollAuthResult `json:"Results,omitempty" name:"Results"`

	// 总记录条数
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 记录起始位置，该次请求资源的起始位置，请求中包含偏移量时应答消息返回相同偏移量，否则返回默认值0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 本次返回条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetPayRollAuthListResponse struct {
	*tchttp.BaseResponse
	Response *GetPayRollAuthListResponseParams `json:"Response"`
}

func (r *GetPayRollAuthListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPayRollAuthListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPayRollAuthRequestParams struct {
	// 用户在商户对应appid下的唯一标识
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 微信服务商下特约商户的商户号，由微信支付生成并下发
	SubMerchantId *string `json:"SubMerchantId,omitempty" name:"SubMerchantId"`

	// 是服务商在微信申请公众号/小程序或移动应用成功后分配的账号ID（与服务商主体一致）
	// 当输入服务商Appid时，会校验其与服务商商户号的绑定关系。服务商APPID和与特约商户APPID至少输入一个，且必须要有拉起领薪卡小程序时使用的APPID
	WechatAppId *string `json:"WechatAppId,omitempty" name:"WechatAppId"`

	// 特约商户在微信申请公众号/小程序或移动应用成功后分配的账号ID（与特约商户主体一致）
	// 当输入特约商户Appid时，会校验其与特约商户号的绑定关系。服务商APPID和与特约商户APPID至少输入一个，且必须要有拉起领薪卡小程序时使用的APPID
	WechatSubAppId *string `json:"WechatSubAppId,omitempty" name:"WechatSubAppId"`
}

type GetPayRollAuthRequest struct {
	*tchttp.BaseRequest
	
	// 用户在商户对应appid下的唯一标识
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 微信服务商下特约商户的商户号，由微信支付生成并下发
	SubMerchantId *string `json:"SubMerchantId,omitempty" name:"SubMerchantId"`

	// 是服务商在微信申请公众号/小程序或移动应用成功后分配的账号ID（与服务商主体一致）
	// 当输入服务商Appid时，会校验其与服务商商户号的绑定关系。服务商APPID和与特约商户APPID至少输入一个，且必须要有拉起领薪卡小程序时使用的APPID
	WechatAppId *string `json:"WechatAppId,omitempty" name:"WechatAppId"`

	// 特约商户在微信申请公众号/小程序或移动应用成功后分配的账号ID（与特约商户主体一致）
	// 当输入特约商户Appid时，会校验其与特约商户号的绑定关系。服务商APPID和与特约商户APPID至少输入一个，且必须要有拉起领薪卡小程序时使用的APPID
	WechatSubAppId *string `json:"WechatSubAppId,omitempty" name:"WechatSubAppId"`
}

func (r *GetPayRollAuthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPayRollAuthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OpenId")
	delete(f, "SubMerchantId")
	delete(f, "WechatAppId")
	delete(f, "WechatSubAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetPayRollAuthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPayRollAuthResponseParams struct {
	// 授权状态：
	// UNAUTHORIZED：未授权
	// AUTHORIZED：已授权
	// DEAUTHORIZED：已取消授权
	AuthStatus *string `json:"AuthStatus,omitempty" name:"AuthStatus"`

	// 授权时间，遵循[rfc3339](https://datatracker.ietf.org/doc/html/rfc3339)标准格式，格式为YYYY-MM-DDTHH:mm:ss.sss+TIMEZONE，空字符串等同null
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthTime *string `json:"AuthTime,omitempty" name:"AuthTime"`

	// 授权时间，遵循[rfc3339](https://datatracker.ietf.org/doc/html/rfc3339)标准格式，格式为YYYY-MM-DDTHH:mm:ss.sss+TIMEZONE，空字符串等同null
	// 注意：此字段可能返回 null，表示取不到有效值。
	CancelAuthTime *string `json:"CancelAuthTime,omitempty" name:"CancelAuthTime"`

	// 微信服务商商户的商户号，由微信支付生成并下发
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 用户在商户对应appid下的唯一标识
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 微信服务商下特约商户的商户号，由微信支付生成并下发
	SubMerchantId *string `json:"SubMerchantId,omitempty" name:"SubMerchantId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetPayRollAuthResponse struct {
	*tchttp.BaseResponse
	Response *GetPayRollAuthResponseParams `json:"Response"`
}

func (r *GetPayRollAuthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPayRollAuthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPayRollAuthResultRequestParams struct {
	// 商户系统内部的商家核身单号，要求此参数只能由数字、大小写字母组成，在服务商内部唯一
	AuthNumber *string `json:"AuthNumber,omitempty" name:"AuthNumber"`

	// 微信服务商下特约商户的商户号，由微信支付生成并下发
	SubMerchantId *string `json:"SubMerchantId,omitempty" name:"SubMerchantId"`
}

type GetPayRollAuthResultRequest struct {
	*tchttp.BaseRequest
	
	// 商户系统内部的商家核身单号，要求此参数只能由数字、大小写字母组成，在服务商内部唯一
	AuthNumber *string `json:"AuthNumber,omitempty" name:"AuthNumber"`

	// 微信服务商下特约商户的商户号，由微信支付生成并下发
	SubMerchantId *string `json:"SubMerchantId,omitempty" name:"SubMerchantId"`
}

func (r *GetPayRollAuthResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPayRollAuthResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AuthNumber")
	delete(f, "SubMerchantId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetPayRollAuthResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPayRollAuthResultResponseParams struct {
	// 核身结果
	Result *PayRollAuthResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetPayRollAuthResultResponse struct {
	*tchttp.BaseResponse
	Response *GetPayRollAuthResultResponseParams `json:"Response"`
}

func (r *GetPayRollAuthResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPayRollAuthResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LegalPersonInfo struct {
	// 证件类型 
	// IDCARD：身份证 
	// PASSPORT：护照 SOLDIERSCERTIFICATE：士兵证 OFFICERSCERTIFICATE：军官证 GATXCERTIFICATE：香港居民来往内地通行证 TWNDCERTIFICATE：台湾同胞来往内地通行证 
	// MACAOCERTIFICATE：澳门来往内地通行证
	IdType *string `json:"IdType,omitempty" name:"IdType"`

	// 证件号码
	IdNumber *string `json:"IdNumber,omitempty" name:"IdNumber"`

	// 姓名
	PersonName *string `json:"PersonName,omitempty" name:"PersonName"`

	// 证件有效期类型 
	// LONGTERM：长期有效 
	// OTHER：非长期有效
	IdValidityType *string `json:"IdValidityType,omitempty" name:"IdValidityType"`

	// 证件生效日期，yyyy-MM-dd
	IdEffectiveDate *string `json:"IdEffectiveDate,omitempty" name:"IdEffectiveDate"`

	// 联系电话
	ContactPhone *string `json:"ContactPhone,omitempty" name:"ContactPhone"`

	// 证件失效日期，yyyy-MM-dd
	IdExpireDate *string `json:"IdExpireDate,omitempty" name:"IdExpireDate"`

	// 联系地址
	ContactAddress *string `json:"ContactAddress,omitempty" name:"ContactAddress"`

	// 邮箱地址
	EmailAddress *string `json:"EmailAddress,omitempty" name:"EmailAddress"`
}

type MemberTransactionItem struct {
	// 交易类型。
	// __1__：转出
	// __2__：转入
	TransType *string `json:"TransType,omitempty" name:"TransType"`

	// 交易状态。
	// __0__：成功
	TranStatus *string `json:"TranStatus,omitempty" name:"TranStatus"`

	// 交易金额。
	TransAmount *string `json:"TransAmount,omitempty" name:"TransAmount"`

	// 交易日期，格式：yyyyMMdd。
	TransDate *string `json:"TransDate,omitempty" name:"TransDate"`

	// 交易时间，格式：HHmmss。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransTime *string `json:"TransTime,omitempty" name:"TransTime"`

	// 银行系统流水号。
	// _平安渠道为见证系统流水号_
	BankSequenceNumber *string `json:"BankSequenceNumber,omitempty" name:"BankSequenceNumber"`

	// 银行记账类型。
	// _平安渠道为：_
	// _1：会员支付_
	// _2：会员冻结_
	// _3：会员解冻_
	// _4：登记挂账_
	// _6：下单预支付_
	// _7：确认并付款_
	// _8：会员退款_
	// _22：见证+收单平台调账_
	// _23：见证+收单资金冻结_
	// _24：见证+收单资金解冻_
	// _25：会员间交易退款_
	// 注意：此字段可能返回 null，表示取不到有效值。
	BankBookingType *string `json:"BankBookingType,omitempty" name:"BankBookingType"`

	// 转入方子账户账号。
	// _平安渠道为转入见证子账户的账号_
	InSubAccountNumber *string `json:"InSubAccountNumber,omitempty" name:"InSubAccountNumber"`

	// 转出方子账户账号。
	// _平安渠道为转出见证子账户的账号_
	OutSubAccountNumber *string `json:"OutSubAccountNumber,omitempty" name:"OutSubAccountNumber"`

	// 备注。
	// _平安渠道，如果是见证+收单的交易，返回交易订单号_
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type MerchantClassificationId struct {
	// 分类编号
	Code *string `json:"Code,omitempty" name:"Code"`

	// 分类名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 父级编号（0为一级编号，大于0为父级分类编号）
	Parent *string `json:"Parent,omitempty" name:"Parent"`
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

type MerchantPayWayData struct {
	// 支付币种
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayCurrency *string `json:"PayCurrency,omitempty" name:"PayCurrency"`

	// 支付图标
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayIcon *string `json:"PayIcon,omitempty" name:"PayIcon"`

	// 支付名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayInternalName *string `json:"PayInternalName,omitempty" name:"PayInternalName"`

	// 支付简称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayName *string `json:"PayName,omitempty" name:"PayName"`

	// 是否支持退款
	// 注意：此字段可能返回 null，表示取不到有效值。
	PaySplitRefund *string `json:"PaySplitRefund,omitempty" name:"PaySplitRefund"`

	// 支付标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayTag *string `json:"PayTag,omitempty" name:"PayTag"`

	// 支付凭证图标
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayTicketIcon *string `json:"PayTicketIcon,omitempty" name:"PayTicketIcon"`

	// 支付类型，逗号分隔
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayType *string `json:"PayType,omitempty" name:"PayType"`

	// 凭证名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TicketName *string `json:"TicketName,omitempty" name:"TicketName"`
}

type MerchantRiskInfo struct {
	// 恶意注册等级，0-9级，9级最高
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskLevel *int64 `json:"RiskLevel,omitempty" name:"RiskLevel"`

	// 恶意注册代码，代码以|分割，如"G001|T002"
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskTypes *string `json:"RiskTypes,omitempty" name:"RiskTypes"`
}

// Predefined struct for user
type MigrateOrderRefundQueryRequestParams struct {
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

// Predefined struct for user
type MigrateOrderRefundQueryResponseParams struct {
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
}

type MigrateOrderRefundQueryResponse struct {
	*tchttp.BaseResponse
	Response *MigrateOrderRefundQueryResponseParams `json:"Response"`
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

// Predefined struct for user
type MigrateOrderRefundRequestParams struct {
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

// Predefined struct for user
type MigrateOrderRefundResponseParams struct {
	// 请求成功状态
	IsSuccess *bool `json:"IsSuccess,omitempty" name:"IsSuccess"`

	// 退款流水号
	TradeSerialNo *string `json:"TradeSerialNo,omitempty" name:"TradeSerialNo"`

	// 交易备注
	TradeMsg *string `json:"TradeMsg,omitempty" name:"TradeMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type MigrateOrderRefundResponse struct {
	*tchttp.BaseResponse
	Response *MigrateOrderRefundResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyAgentTaxPaymentInfoRequestParams struct {
	// 批次号
	BatchNum *int64 `json:"BatchNum,omitempty" name:"BatchNum"`

	// 新源电子凭证地址
	RawElectronicCertUrl *string `json:"RawElectronicCertUrl,omitempty" name:"RawElectronicCertUrl"`

	// 新的文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 接入环境。沙箱环境填sandbox
	Profile *string `json:"Profile,omitempty" name:"Profile"`
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

// Predefined struct for user
type ModifyAgentTaxPaymentInfoResponseParams struct {
	// 代理商完税证明批次信息
	AgentTaxPaymentBatch *AgentTaxPaymentBatch `json:"AgentTaxPaymentBatch,omitempty" name:"AgentTaxPaymentBatch"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAgentTaxPaymentInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAgentTaxPaymentInfoResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyBindedAccountRequestParams struct {
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

// Predefined struct for user
type ModifyBindedAccountResponseParams struct {
	// 错误码。响应成功："SUCCESS"，其他为不成功。
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 响应消息。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 该字段为null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyBindedAccountResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBindedAccountResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyFlexPayeeAccountRightStatusRequestParams struct {
	// 收款用户ID
	PayeeId *string `json:"PayeeId,omitempty" name:"PayeeId"`

	// 账户权益类型
	// SETTLEMENT:结算权益
	// PAYMENT:付款权益
	AccountRightType *string `json:"AccountRightType,omitempty" name:"AccountRightType"`

	// 账户权益状态
	// ENABLE:启用
	// DISABLE:停用
	AccountRightStatus *string `json:"AccountRightStatus,omitempty" name:"AccountRightStatus"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// __test__:测试环境
	// 缺省默认为生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type ModifyFlexPayeeAccountRightStatusRequest struct {
	*tchttp.BaseRequest
	
	// 收款用户ID
	PayeeId *string `json:"PayeeId,omitempty" name:"PayeeId"`

	// 账户权益类型
	// SETTLEMENT:结算权益
	// PAYMENT:付款权益
	AccountRightType *string `json:"AccountRightType,omitempty" name:"AccountRightType"`

	// 账户权益状态
	// ENABLE:启用
	// DISABLE:停用
	AccountRightStatus *string `json:"AccountRightStatus,omitempty" name:"AccountRightStatus"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// __test__:测试环境
	// 缺省默认为生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *ModifyFlexPayeeAccountRightStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFlexPayeeAccountRightStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PayeeId")
	delete(f, "AccountRightType")
	delete(f, "AccountRightStatus")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyFlexPayeeAccountRightStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFlexPayeeAccountRightStatusResponseParams struct {
	// 错误码。SUCCESS为成功，其他为失败
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果。默认为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyFlexPayeeAccountRightStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyFlexPayeeAccountRightStatusResponseParams `json:"Response"`
}

func (r *ModifyFlexPayeeAccountRightStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFlexPayeeAccountRightStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMerchantRequestParams struct {
	// 进件成功后返给商户的AppId
	MerchantAppId *string `json:"MerchantAppId,omitempty" name:"MerchantAppId"`

	// 收款商户名称
	MerchantName *string `json:"MerchantName,omitempty" name:"MerchantName"`

	// B2B 支付标志。是否开通 B2B支付， 1:开通 0:不开通。
	BusinessPayFlag *string `json:"BusinessPayFlag,omitempty" name:"BusinessPayFlag"`
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

// Predefined struct for user
type ModifyMerchantResponseParams struct {
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
type ModifyMntMbrBindRelateAcctBankCodeRequestParams struct {
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

// Predefined struct for user
type ModifyMntMbrBindRelateAcctBankCodeResponseParams struct {
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
}

type ModifyMntMbrBindRelateAcctBankCodeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMntMbrBindRelateAcctBankCodeResponseParams `json:"Response"`
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

type MultiApplyDetail struct {
	// 商户编号
	MerchantNo *string `json:"MerchantNo,omitempty" name:"MerchantNo"`

	// 分账金额
	Amount *string `json:"Amount,omitempty" name:"Amount"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type MultiApplyOrder struct {
	// 商户分账单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutDistributeNo *string `json:"OutDistributeNo,omitempty" name:"OutDistributeNo"`

	// 平台分账单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	DistributeNo *string `json:"DistributeNo,omitempty" name:"DistributeNo"`

	// 平台交易订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderNo *string `json:"OrderNo,omitempty" name:"OrderNo"`

	// 分账订单状态（0初始1成功2失败3撤销）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 入账日期，格式yyyy-MM-dd
	// 注意：此字段可能返回 null，表示取不到有效值。
	InDate *string `json:"InDate,omitempty" name:"InDate"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 分账明细
	// 注意：此字段可能返回 null，表示取不到有效值。
	Details []*MultiApplyDetail `json:"Details,omitempty" name:"Details"`
}

type NaturalPersonInfo struct {
	// 自然人类型 
	// 2：商户负责人 
	// 3：授权经办人
	// 4：业务联系人 
	// 5：实际控制人 
	// 6：控股股东 
	// 7：受益人 
	// 8：结算人
	// 注意：HELIPAY渠道必传业务联系人
	PersonType *string `json:"PersonType,omitempty" name:"PersonType"`

	// 证件类型 
	// IDCARD：身份证 
	// PASSPORT：护照 SOLDIERSCERTIFICATE：士兵证 OFFICERSCERTIFICATE：军官证 GATXCERTIFICATE：香港居民来往内地通行证 TWNDCERTIFICATE：台湾同胞来往内地通行证 MACAOCERTIFICATE：澳门来往内地通行证
	IdType *string `json:"IdType,omitempty" name:"IdType"`

	// 证件号码
	IdNumber *string `json:"IdNumber,omitempty" name:"IdNumber"`

	// 姓名
	PersonName *string `json:"PersonName,omitempty" name:"PersonName"`

	// 证件有效期类型 
	// LONGTERM：长期有效 
	// OTHER：非长期有效
	IdValidityType *string `json:"IdValidityType,omitempty" name:"IdValidityType"`

	// 证件生效日期，yyyy-MM-dd
	IdEffectiveDate *string `json:"IdEffectiveDate,omitempty" name:"IdEffectiveDate"`

	// 证件失效日期，yyyy-MM-dd
	IdExpireDate *string `json:"IdExpireDate,omitempty" name:"IdExpireDate"`

	// 联系电话，HELIPAY渠道业务联系人必传
	ContactPhone *string `json:"ContactPhone,omitempty" name:"ContactPhone"`

	// 联系地址
	ContactAddress *string `json:"ContactAddress,omitempty" name:"ContactAddress"`

	// 邮箱地址
	EmailAddress *string `json:"EmailAddress,omitempty" name:"EmailAddress"`
}

type OpenBankApprovalGuideInfo struct {
	// PC网银指引
	PcGuideUrl *string `json:"PcGuideUrl,omitempty" name:"PcGuideUrl"`

	// 手机网银指引
	// 注意：此字段可能返回 null，表示取不到有效值。
	MobileGuideUrl *string `json:"MobileGuideUrl,omitempty" name:"MobileGuideUrl"`
}

type OpenBankFormInfo struct {
	// 网银页面提交html
	// 注意：此字段可能返回 null，表示取不到有效值。
	FormHtml *string `json:"FormHtml,omitempty" name:"FormHtml"`

	// 网银提交页面字符集
	// 注意：此字段可能返回 null，表示取不到有效值。
	FormEncoding *string `json:"FormEncoding,omitempty" name:"FormEncoding"`
}

type OpenBankGoodsInfo struct {
	// 商品标题。默认值“商品支付”。
	GoodsName *string `json:"GoodsName,omitempty" name:"GoodsName"`

	// 商品详细描述（商品列表）。
	GoodsDetail *string `json:"GoodsDetail,omitempty" name:"GoodsDetail"`

	// 银行附言。不可以有以下字符：<>+{}()%*&';"[]等特殊符号
	GoodsDescription *string `json:"GoodsDescription,omitempty" name:"GoodsDescription"`

	// 业务类型。汇付渠道必填，汇付渠道传入固定值100099。
	GoodsBizType *string `json:"GoodsBizType,omitempty" name:"GoodsBizType"`
}

type OpenBankPayLimitInfo struct {
	// 限制类型
	PayLimitType *string `json:"PayLimitType,omitempty" name:"PayLimitType"`

	// 限制类型值
	PayLimitValue *string `json:"PayLimitValue,omitempty" name:"PayLimitValue"`
}

type OpenBankPayeeInfo struct {
	// 收款方唯一标识。
	// 当渠道为TENPAY，付款方式为EBANK_PAYMENT，必填，上送收款方入驻云企付商户ID；
	// 付款方式为OPENBANK_PAYMENT时，非必填，输入外部收款方的标识ID
	// 渠道为WECHAT，付款方式为TRANS_TO_CHANGE时，上送微信OPEN_ID；
	PayeeId *string `json:"PayeeId,omitempty" name:"PayeeId"`

	// 支行名称。
	BankBranchName *string `json:"BankBranchName,omitempty" name:"BankBranchName"`

	// 银行账号。渠道为TENPAY，付款方式为OPENBANK_PAYMENT时必选
	BankAccountNumber *string `json:"BankAccountNumber,omitempty" name:"BankAccountNumber"`

	// 收款方名称。
	// 当渠道为TENPAY，付款方式为EBANK_PAYMENT时，上送收款方入驻云企付的商户名称；
	// 渠道为TENPAY，付款方式为OPENBANK_PAYMENT时必选，上送收款方账户名称；
	// 渠道为ALIPAY，付款方式为SAFT_ISV时，收款账户标识类型为ALIPAY_LOGON_ID时必传，上送收款方真实姓名。
	// 渠道为WECHAT，付款方式为TRANS_TO_CHANGE时，上送收款人姓名。
	PayeeName *string `json:"PayeeName,omitempty" name:"PayeeName"`

	// 联行号。渠道为TENPAY，付款方式为OPENBANK_PAYMENT时必选
	BankBranchId *string `json:"BankBranchId,omitempty" name:"BankBranchId"`

	// 收款方绑卡序列号。
	// 当渠道为TENPAY，付款方式为EBANK_PAYMENT时，必填，上送收款方入驻云企付平台时，下发的绑卡序列号；当渠道为ALIPAY，付款方式为SAFT_ISV时，必填，根据收款账户标识类型上送。
	BindSerialNo *string `json:"BindSerialNo,omitempty" name:"BindSerialNo"`

	// 收款账户标识类型
	// BANK_ACCOUNT：绑定银行账户
	// ACCOUNT_BOOK_ID：电子记账本ID
	// ALIPAY_USER_ID：支付宝的会员ID
	// ALIPAY_LOGON_ID：支付宝登录号。
	// 付款方式为SAFT_ISV时，必填。
	AccountType *string `json:"AccountType,omitempty" name:"AccountType"`
}

type OpenBankPayerInfo struct {
	// 付款方唯一标识。当TENPAY时，必填上送
	// 付款方入驻云企付商户ID。
	PayerId *string `json:"PayerId,omitempty" name:"PayerId"`

	// 付款方名称。当TENPAY上送付款方入驻云企付的商户名称。
	PayerName *string `json:"PayerName,omitempty" name:"PayerName"`

	// 付款方付款账户标识。
	// 当付款方式为OPENBANK_PAYMENT时，必输表示企业账户ID；当付款方式为SAFT_ISV时，必须上送付款方的渠道电子记账本ID；当付款方式为ONLINEBANK，上送付款方银行编号BankId。
	BindSerialNo *string `json:"BindSerialNo,omitempty" name:"BindSerialNo"`

	// 付款账户标识类型
	// BANK_ACCOUNT：绑定银行账户
	// ACCOUNT_BOOK_ID：电子记账本ID。
	// 当付款方式为SAFT_ISV时，必须上送类型为ACCOUNT_BOOK_ID。
	AccountType *string `json:"AccountType,omitempty" name:"AccountType"`

	// 付款卡类型。汇付渠道必填。
	// DEBIT_CARD：借记卡
	// CREDIT_CARD：信用卡
	BankCardType *string `json:"BankCardType,omitempty" name:"BankCardType"`
}

type OpenBankProfitShareInfo struct {
	// 分润接收方，渠道商户号ID
	RecvId *string `json:"RecvId,omitempty" name:"RecvId"`

	// 分润金额，单位分
	ProfitShareFee *int64 `json:"ProfitShareFee,omitempty" name:"ProfitShareFee"`
}

type OpenBankQueryRefundOrderResult struct {
	// 外部商户退款单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutRefundId *string `json:"OutRefundId,omitempty" name:"OutRefundId"`

	// 渠道退款单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelRefundId *string `json:"ChannelRefundId,omitempty" name:"ChannelRefundId"`

	// 退款原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	RefundReason *string `json:"RefundReason,omitempty" name:"RefundReason"`

	// 退款金额，单位分
	// 注意：此字段可能返回 null，表示取不到有效值。
	RefundAmount *int64 `json:"RefundAmount,omitempty" name:"RefundAmount"`

	// 实际退款金额，单位分
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealRefundAmount *int64 `json:"RealRefundAmount,omitempty" name:"RealRefundAmount"`

	// 原支付订单总金额，单位分
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalAmount *int64 `json:"TotalAmount,omitempty" name:"TotalAmount"`

	// 退款完成时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeFinish *string `json:"TimeFinish,omitempty" name:"TimeFinish"`

	// 退款订单状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	RefundStatus *string `json:"RefundStatus,omitempty" name:"RefundStatus"`

	// 退款明细信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RefundInfo *string `json:"RefundInfo,omitempty" name:"RefundInfo"`

	// 退款手续费金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	FeeAmount *int64 `json:"FeeAmount,omitempty" name:"FeeAmount"`

	// 退款返回描述，比如失败原因等。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RefundMessage *string `json:"RefundMessage,omitempty" name:"RefundMessage"`
}

type OpenBankRechargePayeeInfo struct {
	// 收款方标识
	// 收款方类型为电子记账本时，上送渠道电子记账本ID
	PayeeId *string `json:"PayeeId,omitempty" name:"PayeeId"`

	// 收款方标识类型
	// ACCOUNT_BOOK_ID：电子记账本ID
	PayeeIdType *string `json:"PayeeIdType,omitempty" name:"PayeeIdType"`

	// 收款方名称
	PayeeName *string `json:"PayeeName,omitempty" name:"PayeeName"`
}

type OpenBankRechargeRedirectInfo struct {
	// 跳转URL
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitempty" name:"Url"`
}

type OpenBankRedirectInfo struct {
	// 生成二维码，引导用户扫码
	QRCodeUrl *string `json:"QRCodeUrl,omitempty" name:"QRCodeUrl"`

	// 二维码凭证
	QRCodeKey *string `json:"QRCodeKey,omitempty" name:"QRCodeKey"`

	// 跳转 URL,用于客户端跳转，订单未支付时返回该参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 跳转凭证过期时间,yyyy-MM-dd HH:mm:ss
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 小程序 appid
	MpAppId *string `json:"MpAppId,omitempty" name:"MpAppId"`

	// 小程序路径
	MpPath *string `json:"MpPath,omitempty" name:"MpPath"`

	// 小程序原始 id
	MpUserName *string `json:"MpUserName,omitempty" name:"MpUserName"`

	// 网银支付提交页面信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FormInfo *OpenBankFormInfo `json:"FormInfo,omitempty" name:"FormInfo"`
}

type OpenBankRefundOrderApplyResult struct {
	// 云企付订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelOrderId *string `json:"ChannelOrderId,omitempty" name:"ChannelOrderId"`

	// 云企付退款流水号
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelRefundId *string `json:"ChannelRefundId,omitempty" name:"ChannelRefundId"`

	// 外部商户退款单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutRefundId *string `json:"OutRefundId,omitempty" name:"OutRefundId"`

	// 外部商户订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutOrderId *string `json:"OutOrderId,omitempty" name:"OutOrderId"`

	// 退款返回描述，比如失败原因等。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RefundMessage *string `json:"RefundMessage,omitempty" name:"RefundMessage"`

	// 退款金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	RefundAmount *int64 `json:"RefundAmount,omitempty" name:"RefundAmount"`

	// 退款手续费金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	FeeAmount *int64 `json:"FeeAmount,omitempty" name:"FeeAmount"`

	// 退款状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	RefundStatus *string `json:"RefundStatus,omitempty" name:"RefundStatus"`
}

type OpenBankSceneInfo struct {
	// 用户端实际 ip。汇付渠道必填。
	PayerClientIp *string `json:"PayerClientIp,omitempty" name:"PayerClientIp"`

	// 浏览器 User-Agent。
	PayerUa *string `json:"PayerUa,omitempty" name:"PayerUa"`

	// 用户下单时间。若不上送，服务端默认当前时间。
	OrderTime *string `json:"OrderTime,omitempty" name:"OrderTime"`

	// 终端设备号（门店号或收银设备 ID），示例值：POS1:1。
	DeviceId *string `json:"DeviceId,omitempty" name:"DeviceId"`

	// 终端设备类型。MOBILE_BROWSER:手机浏览器，MOBILE_APP:手机应用程序，TABLET:平板；WATCH:手表，PC:电脑PC，OTHER:其他。
	// 汇付渠道必填。
	DeviceType *string `json:"DeviceType,omitempty" name:"DeviceType"`
}

type OpenBankStoreInfo struct {
	// 门店名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 地区编码
	AreaCode *string `json:"AreaCode,omitempty" name:"AreaCode"`

	// 详细地址
	Address *string `json:"Address,omitempty" name:"Address"`

	// 门店编号
	Id *string `json:"Id,omitempty" name:"Id"`
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

type OrderSummaries struct {
	// 汇总列表
	List []*OrderSummaryResult `json:"List,omitempty" name:"List"`

	// 总数
	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type OrderSummaryResult struct {
	// 汇总ID
	SummaryId *string `json:"SummaryId,omitempty" name:"SummaryId"`

	// 收款账户ID
	PayeeId *string `json:"PayeeId,omitempty" name:"PayeeId"`

	// 收款账户名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 收入类型
	IncomeType *string `json:"IncomeType,omitempty" name:"IncomeType"`

	// 汇总金额
	SummaryAmount *string `json:"SummaryAmount,omitempty" name:"SummaryAmount"`

	// 汇总日期
	SummaryTime *string `json:"SummaryTime,omitempty" name:"SummaryTime"`

	// 汇总记录数量
	SummaryCount *int64 `json:"SummaryCount,omitempty" name:"SummaryCount"`
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

type OutSubMerchantExtensionInfo struct {
	// 地区代码，国标码
	// HELIPAY渠道必传
	RegionCode *string `json:"RegionCode,omitempty" name:"RegionCode"`

	// 注册地址
	RegisterAddress *string `json:"RegisterAddress,omitempty" name:"RegisterAddress"`

	// 通讯地址
	// HELIPAY渠道必传
	MailingAddress *string `json:"MailingAddress,omitempty" name:"MailingAddress"`

	// 营业地址/经营地址
	BusinessAddress *string `json:"BusinessAddress,omitempty" name:"BusinessAddress"`

	// 客服电话
	ServicePhone *string `json:"ServicePhone,omitempty" name:"ServicePhone"`

	// 网站url
	WebSiteUrl *string `json:"WebSiteUrl,omitempty" name:"WebSiteUrl"`

	// 邮箱地址
	EmailAddress *string `json:"EmailAddress,omitempty" name:"EmailAddress"`
}

type Paging struct {
	// 页码
	Index *int64 `json:"Index,omitempty" name:"Index"`

	// 页长
	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type PayDataResult struct {
	// 支付标签（唯一性）
	PaymentTag *string `json:"PaymentTag,omitempty" name:"PaymentTag"`

	// 添加合同时需要隐藏的选项（多个以小写逗号分开）
	PaymentOptionHide *string `json:"PaymentOptionHide,omitempty" name:"PaymentOptionHide"`

	// 支付方式图片url路径
	PaymentIcon *string `json:"PaymentIcon,omitempty" name:"PaymentIcon"`

	// 合同选项名称6
	PaymentOptionSix *string `json:"PaymentOptionSix,omitempty" name:"PaymentOptionSix"`

	// 付款方式名称
	PaymentName *string `json:"PaymentName,omitempty" name:"PaymentName"`

	// 合同选项名称7
	PaymentOptionSeven *string `json:"PaymentOptionSeven,omitempty" name:"PaymentOptionSeven"`

	// 合同选项名称8
	PaymentOptionOther *string `json:"PaymentOptionOther,omitempty" name:"PaymentOptionOther"`

	// 合同选项名称2
	PaymentOptionTwo *string `json:"PaymentOptionTwo,omitempty" name:"PaymentOptionTwo"`

	// 合同选项名称1
	PaymentOptionOne *string `json:"PaymentOptionOne,omitempty" name:"PaymentOptionOne"`

	// 是否开启智能扣率（1是，0否）
	PaymentDiscountFee *string `json:"PaymentDiscountFee,omitempty" name:"PaymentDiscountFee"`

	// 支持的交易类型（多个以小写逗号分开，0现金，1刷卡，2主扫，3被扫，4JSPAY，5预授权）
	PaymentType *string `json:"PaymentType,omitempty" name:"PaymentType"`

	// 合同选项名称5
	PaymentOptionFive *string `json:"PaymentOptionFive,omitempty" name:"PaymentOptionFive"`

	// 合同选项名称9
	PaymentOptionNine *string `json:"PaymentOptionNine,omitempty" name:"PaymentOptionNine"`

	// 支付方式编号
	PaymentId *string `json:"PaymentId,omitempty" name:"PaymentId"`

	// 合同选项名称3
	PaymentOptionThree *string `json:"PaymentOptionThree,omitempty" name:"PaymentOptionThree"`

	// 付款方式名称（内部名称）
	PaymentInternalName *string `json:"PaymentInternalName,omitempty" name:"PaymentInternalName"`

	// 合同选项名称4
	PaymentOptionFour *string `json:"PaymentOptionFour,omitempty" name:"PaymentOptionFour"`

	// 合同选项名称10
	PaymentOptionTen *string `json:"PaymentOptionTen,omitempty" name:"PaymentOptionTen"`
}

type PayFeeDataResult struct {
	// 机构的分佣扣率扣率
	OrganizationFee *string `json:"OrganizationFee,omitempty" name:"OrganizationFee"`

	// 商户手续费封顶值，0为不限封顶
	PaymentClassificationLimit *string `json:"PaymentClassificationLimit,omitempty" name:"PaymentClassificationLimit"`

	// 机构的分佣扣率类型(1按签约扣率，2按收单收益)
	OrganizationFeeType *string `json:"OrganizationFeeType,omitempty" name:"OrganizationFeeType"`

	// 商户扣率最大值
	PaymentClassificationMaxFee *string `json:"PaymentClassificationMaxFee,omitempty" name:"PaymentClassificationMaxFee"`

	// 商户扣率最小值
	PaymentClassificationMinFee *string `json:"PaymentClassificationMinFee,omitempty" name:"PaymentClassificationMinFee"`

	// 行业会类编号
	PaymentClassificationId *string `json:"PaymentClassificationId,omitempty" name:"PaymentClassificationId"`

	// 行业分类名称
	PaymentClassificationName *string `json:"PaymentClassificationName,omitempty" name:"PaymentClassificationName"`
}

type PayOrderResult struct {
	// 付款订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderNo *string `json:"OrderNo,omitempty" name:"OrderNo"`

	// 开发者流水号
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeveloperNo *string `json:"DeveloperNo,omitempty" name:"DeveloperNo"`

	// 交易优惠金额（免充值券）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TradeDiscountAmount *string `json:"TradeDiscountAmount,omitempty" name:"TradeDiscountAmount"`

	// 付款方式名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayName *string `json:"PayName,omitempty" name:"PayName"`

	// 商户流水号（从1开始自增长不重复）
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderMerchantId *string `json:"OrderMerchantId,omitempty" name:"OrderMerchantId"`

	// 交易帐号（银行卡号、支付宝帐号、微信帐号等，某些收单机构没有此数据）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TradeAccount *string `json:"TradeAccount,omitempty" name:"TradeAccount"`

	// 实际交易金额（以分为单位，没有小数点）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TradeAmount *string `json:"TradeAmount,omitempty" name:"TradeAmount"`

	// 币种签名
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrencySign *string `json:"CurrencySign,omitempty" name:"CurrencySign"`

	// 付款完成时间（以收单机构为准）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TradePayTime *string `json:"TradePayTime,omitempty" name:"TradePayTime"`

	// 门店流水号（从1开始自增长不重复）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShopOrderId *string `json:"ShopOrderId,omitempty" name:"ShopOrderId"`

	// 支付标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayTag *string `json:"PayTag,omitempty" name:"PayTag"`

	// 订单状态（1交易成功，2待支付，4已取消，9等待用户输入密码确认
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 币种代码
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderCurrency *string `json:"OrderCurrency,omitempty" name:"OrderCurrency"`

	// 二维码字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	TradeQrcode *string `json:"TradeQrcode,omitempty" name:"TradeQrcode"`

	// 微信返回调起小程序/原生JS支付的appid参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	WechatAppId *string `json:"WechatAppId,omitempty" name:"WechatAppId"`

	// 微信返回调起小程序/原生JS支付的timeStamp参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	WechatTimeStamp *string `json:"WechatTimeStamp,omitempty" name:"WechatTimeStamp"`

	// 微信返回调起小程序/原生JS支付的nonceStr参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	WechatNonceStr *string `json:"WechatNonceStr,omitempty" name:"WechatNonceStr"`

	// 微信返回调起小程序/原生JS支付的signType参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	WechatSignType *string `json:"WechatSignType,omitempty" name:"WechatSignType"`

	// 微信返回调起小程序/原生JS支付的package参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	WechatPackage *string `json:"WechatPackage,omitempty" name:"WechatPackage"`

	// 微信返回调起小程序/原生JS支付的paySign参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	WechatPaySign *string `json:"WechatPaySign,omitempty" name:"WechatPaySign"`
}

type PayRollAuthResult struct {
	// 结果为核身失败时的原因描述，仅在失败记录返回，空字符串等同null
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthFailedReason *string `json:"AuthFailedReason,omitempty" name:"AuthFailedReason"`

	// 商户系统内部的商家核身单号，要求此参数只能由数字、大小写字母组成，在服务商内部唯一
	AuthNumber *string `json:"AuthNumber,omitempty" name:"AuthNumber"`

	// 核身渠道，发起核身时的来源渠道，如通过小程序，硬件设备等
	// FROM_MINI_APP：来自小程序方式核身
	// FROM_HARDWARE：来自硬件设备方式核身
	AuthScene *string `json:"AuthScene,omitempty" name:"AuthScene"`

	// 核身渠道标识，用于定位渠道具体来源，如果是扫码打卡渠道标识就是具体的小程序appid，若是硬件设备，则是设备的序列号等
	AuthSource *string `json:"AuthSource,omitempty" name:"AuthSource"`

	// 核身状态
	// AUTHENTICATE_PROCESSING：核身中
	// AUTHENTICATE_SUCCESS：核身成功
	// AUTHENTICATE_FAILED：核身失败
	AuthStatus *string `json:"AuthStatus,omitempty" name:"AuthStatus"`

	// 核身时间，遵循[rfc3339](https://datatracker.ietf.org/doc/html/rfc3339)标准格式，格式为YYYY-MM-DDTHH:mm:ss.sss+TIMEZONE
	AuthTime *string `json:"AuthTime,omitempty" name:"AuthTime"`

	// 该用户所属的单位名称
	CompanyName *string `json:"CompanyName,omitempty" name:"CompanyName"`

	// 微信服务商商户的商户号，由微信支付生成并下发
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 用户在商户对应appid下的唯一标识
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 该项目的名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 微信服务商下特约商户的商户号，由微信支付生成并下发
	SubMerchantId *string `json:"SubMerchantId,omitempty" name:"SubMerchantId"`
}

type PayeeAccountBalanceResult struct {
	// 账户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountId *string `json:"AccountId,omitempty" name:"AccountId"`

	// 收入类型
	// LABOR:劳务所得
	// OCCASION:偶然所得
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncomeType *int64 `json:"IncomeType,omitempty" name:"IncomeType"`

	// 总余额
	// 注意：此字段可能返回 null，表示取不到有效值。
	Balance *string `json:"Balance,omitempty" name:"Balance"`

	// 系统冻结余额
	// 注意：此字段可能返回 null，表示取不到有效值。
	SystemFreezeBalance *string `json:"SystemFreezeBalance,omitempty" name:"SystemFreezeBalance"`

	// 人工冻结余额
	// 注意：此字段可能返回 null，表示取不到有效值。
	ManualFreezeBalance *string `json:"ManualFreezeBalance,omitempty" name:"ManualFreezeBalance"`

	// 可提现余额
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayableBalance *string `json:"PayableBalance,omitempty" name:"PayableBalance"`

	// 已提现余额
	// 注意：此字段可能返回 null，表示取不到有效值。
	PaidBalance *string `json:"PaidBalance,omitempty" name:"PaidBalance"`

	// 提现中余额
	// 注意：此字段可能返回 null，表示取不到有效值。
	InPayBalance *string `json:"InPayBalance,omitempty" name:"InPayBalance"`
}

type PayeeAccountInfoResult struct {
	// 账户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountId *string `json:"AccountId,omitempty" name:"AccountId"`

	// 账户名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 用户信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserInfo *PayeeAccountUserInfo `json:"UserInfo,omitempty" name:"UserInfo"`

	// 属性信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PropertyInfo *PayeeAccountPropertyInfo `json:"PropertyInfo,omitempty" name:"PropertyInfo"`
}

type PayeeAccountInfos struct {
	// 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*PayeeAccountInfoResult `json:"List,omitempty" name:"List"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type PayeeAccountPropertyInfo struct {
	// 结算权益状态
	// ENABLE:启用
	// DISABLE:停用
	// 注意：此字段可能返回 null，表示取不到有效值。
	SettleRightStatus *string `json:"SettleRightStatus,omitempty" name:"SettleRightStatus"`

	// 付款权益状态
	// ENABLE:启用
	// DISABLE:停用
	// 注意：此字段可能返回 null，表示取不到有效值。
	PaymentRightStatus *string `json:"PaymentRightStatus,omitempty" name:"PaymentRightStatus"`
}

type PayeeAccountUserInfo struct {
	// 外部用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutUserId *string `json:"OutUserId,omitempty" name:"OutUserId"`

	// 用户类型
	// 0:B端用户
	// 1:C端用户
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserType *int64 `json:"UserType,omitempty" name:"UserType"`

	// 证件类型
	// 0:身份证
	// 1:社会信用代码
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdType *int64 `json:"IdType,omitempty" name:"IdType"`

	// 证件号
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdNo *string `json:"IdNo,omitempty" name:"IdNo"`

	// 姓名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`
}

type PayeeInfoResult struct {
	// 收款用户ID
	PayeeId *string `json:"PayeeId,omitempty" name:"PayeeId"`

	// 用户外部业务ID
	OutUserId *string `json:"OutUserId,omitempty" name:"OutUserId"`

	// 姓名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 证件类型
	// 0:身份证
	// 1:社会信用代码
	IdType *int64 `json:"IdType,omitempty" name:"IdType"`

	// 证件号
	IdNo *string `json:"IdNo,omitempty" name:"IdNo"`

	// 服务商ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceProviderId *string `json:"ServiceProviderId,omitempty" name:"ServiceProviderId"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type PayeeTaxInfo struct {
	// 计税模板列表
	TaxTemplateInfoList []*PayeeTaxTemplateInfo `json:"TaxTemplateInfoList,omitempty" name:"TaxTemplateInfoList"`

	// 纳税人识别号
	TaxpayerIdNo *string `json:"TaxpayerIdNo,omitempty" name:"TaxpayerIdNo"`

	// 纳税主体类型
	// NATURAL:自然人
	// NON_NATURAL:非自然人
	TaxEntityType *string `json:"TaxEntityType,omitempty" name:"TaxEntityType"`

	// 财税服务商ID
	TaxServiceProviderId *string `json:"TaxServiceProviderId,omitempty" name:"TaxServiceProviderId"`
}

type PayeeTaxTemplateInfo struct {
	// 收入类型
	// LABOR: 劳务所得
	// OCCASION: 偶然所得
	IncomeType *string `json:"IncomeType,omitempty" name:"IncomeType"`

	// 计税模板ID
	TaxTemplateId *string `json:"TaxTemplateId,omitempty" name:"TaxTemplateId"`
}

type PaymentOrderResult struct {
	// 收入类型
	// LABOR:劳务所得
	// OCCASION:偶然所得
	IncomeType *string `json:"IncomeType,omitempty" name:"IncomeType"`

	// 税前金额
	AmountBeforeTax *string `json:"AmountBeforeTax,omitempty" name:"AmountBeforeTax"`

	// 税后金额
	AmountAfterTax *string `json:"AmountAfterTax,omitempty" name:"AmountAfterTax"`

	// 税金
	Tax *string `json:"Tax,omitempty" name:"Tax"`

	// 外部订单ID
	OutOrderId *string `json:"OutOrderId,omitempty" name:"OutOrderId"`

	// 订单ID
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 发起时间
	InitiateTime *string `json:"InitiateTime,omitempty" name:"InitiateTime"`

	// 完成时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`

	// 状态
	// ACCEPTED:已受理
	// ACCOUNTED:已记账
	// PAYING:付款中
	// PAYED:完成付款渠道调用
	// SUCCEED:已成功
	// FAILED:已失败
	Status *string `json:"Status,omitempty" name:"Status"`

	// 状态描述
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

	// 提现备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 收款用户ID
	PayeeId *string `json:"PayeeId,omitempty" name:"PayeeId"`
}

type PaymentOrderStatusResult struct {
	// 状态
	// ACCEPTED:已受理
	// ACCOUNTED:已记账
	// PAYING:付款中
	// PAYED:完成付款渠道调用
	// SUCCEED:已成功
	// FAILED:已失败
	Status *string `json:"Status,omitempty" name:"Status"`

	// 状态描述
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`
}

type PaymentOrders struct {
	// 列表
	List []*PaymentOrderResult `json:"List,omitempty" name:"List"`

	// 总数
	Count *int64 `json:"Count,omitempty" name:"Count"`
}

// Predefined struct for user
type QueryAcctBindingRequestParams struct {
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

// Predefined struct for user
type QueryAcctBindingResponseParams struct {
	// 总行数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 银行卡信息列表
	BankCardItems []*BankCardItem `json:"BankCardItems,omitempty" name:"BankCardItems"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryAcctBindingResponse struct {
	*tchttp.BaseResponse
	Response *QueryAcctBindingResponseParams `json:"Response"`
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

// Predefined struct for user
type QueryAcctInfoListRequestParams struct {
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

// Predefined struct for user
type QueryAcctInfoListResponseParams struct {
	// 本次交易返回查询结果记录数
	ResultCount *int64 `json:"ResultCount,omitempty" name:"ResultCount"`

	// 符合业务查询条件的记录总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 查询结果项 [object,object]
	QueryAcctItems []*QueryAcctItem `json:"QueryAcctItems,omitempty" name:"QueryAcctItems"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryAcctInfoListResponse struct {
	*tchttp.BaseResponse
	Response *QueryAcctInfoListResponseParams `json:"Response"`
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

// Predefined struct for user
type QueryAcctInfoRequestParams struct {
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

// Predefined struct for user
type QueryAcctInfoResponseParams struct {
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
}

type QueryAcctInfoResponse struct {
	*tchttp.BaseResponse
	Response *QueryAcctInfoResponseParams `json:"Response"`
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

	// 子账号
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

// Predefined struct for user
type QueryAgentStatementsRequestParams struct {
	// 结算单日期，月结算单填每月1日
	Date *string `json:"Date,omitempty" name:"Date"`

	// 日结算单:daily
	// 月结算单:monthly
	Type *string `json:"Type,omitempty" name:"Type"`
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

// Predefined struct for user
type QueryAgentStatementsResponseParams struct {
	// 文件下载链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryAgentStatementsResponse struct {
	*tchttp.BaseResponse
	Response *QueryAgentStatementsResponseParams `json:"Response"`
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

// Predefined struct for user
type QueryAgentTaxPaymentBatchRequestParams struct {
	// 批次号
	BatchNum *int64 `json:"BatchNum,omitempty" name:"BatchNum"`

	// 接入环境。沙箱环境填sandbox
	Profile *string `json:"Profile,omitempty" name:"Profile"`
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

// Predefined struct for user
type QueryAgentTaxPaymentBatchResponseParams struct {
	// 代理商完税证明批次信息
	AgentTaxPaymentBatch *AgentTaxPaymentBatch `json:"AgentTaxPaymentBatch,omitempty" name:"AgentTaxPaymentBatch"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryAgentTaxPaymentBatchResponse struct {
	*tchttp.BaseResponse
	Response *QueryAgentTaxPaymentBatchResponseParams `json:"Response"`
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

// Predefined struct for user
type QueryAnchorContractInfoRequestParams struct {
	// 起始时间，格式为yyyy-MM-dd
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 起始时间，格式为yyyy-MM-dd
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
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

// Predefined struct for user
type QueryAnchorContractInfoResponseParams struct {
	// 签约主播数据
	AnchorContractInfoList []*AnchorContractInfo `json:"AnchorContractInfoList,omitempty" name:"AnchorContractInfoList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryAnchorContractInfoResponse struct {
	*tchttp.BaseResponse
	Response *QueryAnchorContractInfoResponseParams `json:"Response"`
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

// Predefined struct for user
type QueryApplicationMaterialRequestParams struct {
	// 申报流水号
	DeclareId *string `json:"DeclareId,omitempty" name:"DeclareId"`

	// 接入环境。沙箱环境填sandbox
	Profile *string `json:"Profile,omitempty" name:"Profile"`
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

// Predefined struct for user
type QueryApplicationMaterialResponseParams struct {
	// 成功申报材料查询结果
	Result *QueryDeclareResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryApplicationMaterialResponse struct {
	*tchttp.BaseResponse
	Response *QueryApplicationMaterialResponseParams `json:"Response"`
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

// Predefined struct for user
type QueryAssignmentRequestParams struct {
	// 主播ID
	AnchorId *string `json:"AnchorId,omitempty" name:"AnchorId"`
}

type QueryAssignmentRequest struct {
	*tchttp.BaseRequest
	
	// 主播ID
	AnchorId *string `json:"AnchorId,omitempty" name:"AnchorId"`
}

func (r *QueryAssignmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryAssignmentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AnchorId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryAssignmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryAssignmentResponseParams struct {
	// 错误码。响应成功："SUCCESS"，其他为不成功。
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 响应消息
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 返回响应
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *AssignmentData `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryAssignmentResponse struct {
	*tchttp.BaseResponse
	Response *QueryAssignmentResponseParams `json:"Response"`
}

func (r *QueryAssignmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryAssignmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryBalanceRequestParams struct {
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

// Predefined struct for user
type QueryBalanceResponseParams struct {
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
}

type QueryBalanceResponse struct {
	*tchttp.BaseResponse
	Response *QueryBalanceResponseParams `json:"Response"`
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

// Predefined struct for user
type QueryBankClearRequestParams struct {
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

// Predefined struct for user
type QueryBankClearResponseParams struct {
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
}

type QueryBankClearResponse struct {
	*tchttp.BaseResponse
	Response *QueryBankClearResponseParams `json:"Response"`
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

// Predefined struct for user
type QueryBankTransactionDetailsRequestParams struct {
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

// Predefined struct for user
type QueryBankTransactionDetailsResponseParams struct {
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
}

type QueryBankTransactionDetailsResponse struct {
	*tchttp.BaseResponse
	Response *QueryBankTransactionDetailsResponseParams `json:"Response"`
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

// Predefined struct for user
type QueryBankWithdrawCashDetailsRequestParams struct {
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

// Predefined struct for user
type QueryBankWithdrawCashDetailsResponseParams struct {
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
}

type QueryBankWithdrawCashDetailsResponse struct {
	*tchttp.BaseResponse
	Response *QueryBankWithdrawCashDetailsResponseParams `json:"Response"`
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

type QueryBatchPaymentResultData struct {
	// 批次号
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`

	// 批次总额
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalAmount *int64 `json:"TotalAmount,omitempty" name:"TotalAmount"`

	// 批次总量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 批次预留字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReqReserved *string `json:"ReqReserved,omitempty" name:"ReqReserved"`

	// 批次备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 渠道类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferType *int64 `json:"TransferType,omitempty" name:"TransferType"`

	// 转账明细
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferInfoList []*QueryBatchPaymentResultDataInfo `json:"TransferInfoList,omitempty" name:"TransferInfoList"`
}

type QueryBatchPaymentResultDataInfo struct {
	// 订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 代理商ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`

	// 代理商名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentName *string `json:"AgentName,omitempty" name:"AgentName"`

	// 交易状态。
	// 0 处理中  
	// 1 预占成功 
	// 2 交易成功 
	// 3 交易失败 
	// 4 未知渠道异常 
	// 5 预占额度失败
	// 6 提交成功
	// 7 提交失败
	// 8 订单重复提交
	// 99 未知系统异常
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 状态描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

	// 转账金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferAmount *int64 `json:"TransferAmount,omitempty" name:"TransferAmount"`
}

// Predefined struct for user
type QueryBatchPaymentResultRequestParams struct {
	// 批次号
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`
}

type QueryBatchPaymentResultRequest struct {
	*tchttp.BaseRequest
	
	// 批次号
	BatchId *string `json:"BatchId,omitempty" name:"BatchId"`
}

func (r *QueryBatchPaymentResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryBatchPaymentResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BatchId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryBatchPaymentResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryBatchPaymentResultResponseParams struct {
	// 错误码。响应成功："SUCCESS"，其他为不成功。
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 响应消息。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回响应
	Result *QueryBatchPaymentResultData `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryBatchPaymentResultResponse struct {
	*tchttp.BaseResponse
	Response *QueryBatchPaymentResultResponseParams `json:"Response"`
}

func (r *QueryBatchPaymentResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryBatchPaymentResultResponse) FromJsonString(s string) error {
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

// Predefined struct for user
type QueryBillDownloadURLRequestParams struct {
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

// Predefined struct for user
type QueryBillDownloadURLResponseParams struct {
	// 错误码。响应成功："SUCCESS"，其他为不成功
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 响应消息
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *QueryBillDownloadURLData `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryBillDownloadURLResponse struct {
	*tchttp.BaseResponse
	Response *QueryBillDownloadURLResponseParams `json:"Response"`
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

// Predefined struct for user
type QueryCityCodeRequestParams struct {
	// 收单系统分配的开放ID
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 收单系统分配的密钥
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

type QueryCityCodeRequest struct {
	*tchttp.BaseRequest
	
	// 收单系统分配的开放ID
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 收单系统分配的密钥
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *QueryCityCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCityCodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OpenId")
	delete(f, "OpenKey")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryCityCodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryCityCodeResponseParams struct {
	// 业务系统返回消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 业务系统返回码
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 查询城市编码响应对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result []*CityCodeResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryCityCodeResponse struct {
	*tchttp.BaseResponse
	Response *QueryCityCodeResponseParams `json:"Response"`
}

func (r *QueryCityCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCityCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryCloudChannelDataRequestParams struct {
	// 米大师分配的支付主MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 业务订单号，外部订单号
	OutOrderNo *string `json:"OutOrderNo,omitempty" name:"OutOrderNo"`

	// 数据类型
	// PAYMENT:支付
	ExternalChannelDataType *string `json:"ExternalChannelDataType,omitempty" name:"ExternalChannelDataType"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`

	// 子应用ID
	SubAppId *string `json:"SubAppId,omitempty" name:"SubAppId"`

	// 渠道订单号
	ChannelOrderId *string `json:"ChannelOrderId,omitempty" name:"ChannelOrderId"`

	// 渠道名称，指定渠道查询
	// wechat:微信支付
	Channel *string `json:"Channel,omitempty" name:"Channel"`
}

type QueryCloudChannelDataRequest struct {
	*tchttp.BaseRequest
	
	// 米大师分配的支付主MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 业务订单号，外部订单号
	OutOrderNo *string `json:"OutOrderNo,omitempty" name:"OutOrderNo"`

	// 数据类型
	// PAYMENT:支付
	ExternalChannelDataType *string `json:"ExternalChannelDataType,omitempty" name:"ExternalChannelDataType"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`

	// 子应用ID
	SubAppId *string `json:"SubAppId,omitempty" name:"SubAppId"`

	// 渠道订单号
	ChannelOrderId *string `json:"ChannelOrderId,omitempty" name:"ChannelOrderId"`

	// 渠道名称，指定渠道查询
	// wechat:微信支付
	Channel *string `json:"Channel,omitempty" name:"Channel"`
}

func (r *QueryCloudChannelDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCloudChannelDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MidasAppId")
	delete(f, "OutOrderNo")
	delete(f, "ExternalChannelDataType")
	delete(f, "MidasEnvironment")
	delete(f, "SubAppId")
	delete(f, "ChannelOrderId")
	delete(f, "Channel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryCloudChannelDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryCloudChannelDataResponseParams struct {
	// 外部订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutOrderNo *string `json:"OutOrderNo,omitempty" name:"OutOrderNo"`

	// 渠道订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelOrderId *string `json:"ChannelOrderId,omitempty" name:"ChannelOrderId"`

	// 第三方渠道数据类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalChannelDataType *string `json:"ExternalChannelDataType,omitempty" name:"ExternalChannelDataType"`

	// 渠道名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// 第三方渠道数据列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalChannelDataList []*CloudExternalChannelData `json:"ExternalChannelDataList,omitempty" name:"ExternalChannelDataList"`

	// 子应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubAppId *string `json:"SubAppId,omitempty" name:"SubAppId"`

	// 米大师分配的支付主MidasAppId
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryCloudChannelDataResponse struct {
	*tchttp.BaseResponse
	Response *QueryCloudChannelDataResponseParams `json:"Response"`
}

func (r *QueryCloudChannelDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCloudChannelDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryCloudOrderRequestParams struct {
	// 米大师分配的支付主MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 用户Id，长度不小于5位，仅支持字母和数字的组合
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 查询类型
	// by_order:根据订单号查订单
	Type *string `json:"Type,omitempty" name:"Type"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`

	// 开发者的主订单号
	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`
}

type QueryCloudOrderRequest struct {
	*tchttp.BaseRequest
	
	// 米大师分配的支付主MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 用户Id，长度不小于5位，仅支持字母和数字的组合
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 查询类型
	// by_order:根据订单号查订单
	Type *string `json:"Type,omitempty" name:"Type"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`

	// 开发者的主订单号
	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`
}

func (r *QueryCloudOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCloudOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MidasAppId")
	delete(f, "UserId")
	delete(f, "Type")
	delete(f, "MidasEnvironment")
	delete(f, "OutTradeNo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryCloudOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryCloudOrderResponseParams struct {
	// 订单数量
	TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`

	// 订单列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderList []*CloudOrderReturn `json:"OrderList,omitempty" name:"OrderList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryCloudOrderResponse struct {
	*tchttp.BaseResponse
	Response *QueryCloudOrderResponseParams `json:"Response"`
}

func (r *QueryCloudOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCloudOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryCloudRefundOrderRequestParams struct {
	// 米大师分配的支付主MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 用户Id，长度不小于5位，仅支持字母和数字的组合
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 退款订单号，仅支持数字、字母、下划线（_）、横杠字符（-）、点（.）的组合
	RefundId *string `json:"RefundId,omitempty" name:"RefundId"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`
}

type QueryCloudRefundOrderRequest struct {
	*tchttp.BaseRequest
	
	// 米大师分配的支付主MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 用户Id，长度不小于5位，仅支持字母和数字的组合
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 退款订单号，仅支持数字、字母、下划线（_）、横杠字符（-）、点（.）的组合
	RefundId *string `json:"RefundId,omitempty" name:"RefundId"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`
}

func (r *QueryCloudRefundOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCloudRefundOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MidasAppId")
	delete(f, "UserId")
	delete(f, "RefundId")
	delete(f, "MidasEnvironment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryCloudRefundOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryCloudRefundOrderResponseParams struct {
	// 该笔退款订单对应的UnifiedOrder下单时传入的OutTradeNo
	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`

	// 该笔退款订单对应的支付成功后支付机构返回的支付订单号
	ChannelExternalOrderId *string `json:"ChannelExternalOrderId,omitempty" name:"ChannelExternalOrderId"`

	// 该笔退款订单退款后支付机构返回的退款单号
	ChannelExternalRefundId *string `json:"ChannelExternalRefundId,omitempty" name:"ChannelExternalRefundId"`

	// 内部请求微信支付、银行等支付机构的订单号
	ChannelOrderId *string `json:"ChannelOrderId,omitempty" name:"ChannelOrderId"`

	// 请求退款时传的退款ID后查询退款时传的RefundId
	RefundId *string `json:"RefundId,omitempty" name:"RefundId"`

	// 被使用的RefundId，业务可忽略该字段
	UsedRefundId *string `json:"UsedRefundId,omitempty" name:"UsedRefundId"`

	// 退款总金额
	TotalRefundAmt *int64 `json:"TotalRefundAmt,omitempty" name:"TotalRefundAmt"`

	// ISO货币代码
	CurrencyType *string `json:"CurrencyType,omitempty" name:"CurrencyType"`

	// 退款状态码，退款提交成功后返回
	// 1:退款中
	// 2:退款成功
	// 3:退款失败
	State *string `json:"State,omitempty" name:"State"`

	// 子单退款信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubRefundList []*CloudSubRefundItem `json:"SubRefundList,omitempty" name:"SubRefundList"`

	// 透传字段，退款成功回调透传给应用，用于开发者透传自定义内容
	Metadata *string `json:"Metadata,omitempty" name:"Metadata"`

	// 米大师分配的支付主MidasAppId
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// 该笔退款订单退款后内部返回的退款单号
	ChannelRefundId *string `json:"ChannelRefundId,omitempty" name:"ChannelRefundId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryCloudRefundOrderResponse struct {
	*tchttp.BaseResponse
	Response *QueryCloudRefundOrderResponseParams `json:"Response"`
}

func (r *QueryCloudRefundOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryCloudRefundOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryCommonTransferRechargeRequestParams struct {
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

// Predefined struct for user
type QueryCommonTransferRechargeResponseParams struct {
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
}

type QueryCommonTransferRechargeResponse struct {
	*tchttp.BaseResponse
	Response *QueryCommonTransferRechargeResponseParams `json:"Response"`
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

// Predefined struct for user
type QueryContractPayFeeRequestParams struct {
	// 收单系统分配的开放ID
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 收单系统分配的密钥
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 支付方式编号
	PaymentId *string `json:"PaymentId,omitempty" name:"PaymentId"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

type QueryContractPayFeeRequest struct {
	*tchttp.BaseRequest
	
	// 收单系统分配的开放ID
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 收单系统分配的密钥
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 支付方式编号
	PaymentId *string `json:"PaymentId,omitempty" name:"PaymentId"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *QueryContractPayFeeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryContractPayFeeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OpenId")
	delete(f, "OpenKey")
	delete(f, "PaymentId")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryContractPayFeeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryContractPayFeeResponseParams struct {
	// 业务系统返回消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 业务系统返回码，0表示成功，其他表示失败。
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 查询支付方式费率及自定义表单项响应对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *QueryContractPayFeeResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryContractPayFeeResponse struct {
	*tchttp.BaseResponse
	Response *QueryContractPayFeeResponseParams `json:"Response"`
}

func (r *QueryContractPayFeeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryContractPayFeeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryContractPayFeeResult struct {
	// pay支付方式json数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pay *PayDataResult `json:"Pay,omitempty" name:"Pay"`

	// 合同扩展自定义字段
	ExtraInput []*string `json:"ExtraInput,omitempty" name:"ExtraInput"`

	// pay_fee支付方式行业分类费率json数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayFee []*PayFeeDataResult `json:"PayFee,omitempty" name:"PayFee"`
}

// Predefined struct for user
type QueryContractPayWayListRequestParams struct {
	// 收单系统分配的开放ID
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 收单系统分配的密钥
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

type QueryContractPayWayListRequest struct {
	*tchttp.BaseRequest
	
	// 收单系统分配的开放ID
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 收单系统分配的密钥
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *QueryContractPayWayListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryContractPayWayListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OpenId")
	delete(f, "OpenKey")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryContractPayWayListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryContractPayWayListResponseParams struct {
	// 业务系统返回消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 业务系统返回码
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 查询合同支付方式响应对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result []*ContractPayListResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryContractPayWayListResponse struct {
	*tchttp.BaseResponse
	Response *QueryContractPayWayListResponseParams `json:"Response"`
}

func (r *QueryContractPayWayListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryContractPayWayListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryContractRelateShopRequestParams struct {
	// 收单系统分配的开放ID
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 收单系统分配的密钥
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 合同主键
	ContractId *string `json:"ContractId,omitempty" name:"ContractId"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

type QueryContractRelateShopRequest struct {
	*tchttp.BaseRequest
	
	// 收单系统分配的开放ID
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 收单系统分配的密钥
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 合同主键
	ContractId *string `json:"ContractId,omitempty" name:"ContractId"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *QueryContractRelateShopRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryContractRelateShopRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OpenId")
	delete(f, "OpenKey")
	delete(f, "ContractId")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryContractRelateShopRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryContractRelateShopResponseParams struct {
	// 业务系统返回消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 业务系统返回码
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 查询合同可关联门店响应对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result []*QueryContractRelateShopResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryContractRelateShopResponse struct {
	*tchttp.BaseResponse
	Response *QueryContractRelateShopResponseParams `json:"Response"`
}

func (r *QueryContractRelateShopResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryContractRelateShopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryContractRelateShopResult struct {
	// 省份
	// 注意：此字段可能返回 null，表示取不到有效值。
	Province *string `json:"Province,omitempty" name:"Province"`

	// 城市编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	CityId *string `json:"CityId,omitempty" name:"CityId"`

	// 门店简称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShopName *string `json:"ShopName,omitempty" name:"ShopName"`

	// 终端数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TerminalCount *string `json:"TerminalCount,omitempty" name:"TerminalCount"`

	// 城市
	// 注意：此字段可能返回 null，表示取不到有效值。
	City *string `json:"City,omitempty" name:"City"`

	// 门店状态（0未审核，1已审核，2审核未通过，3待审核，4已删除，5初审通过）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShopStatus *string `json:"ShopStatus,omitempty" name:"ShopStatus"`

	// 若是支付宝合同，支付宝上线状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	AliPayOnline *string `json:"AliPayOnline,omitempty" name:"AliPayOnline"`

	// 门店编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShopNo *string `json:"ShopNo,omitempty" name:"ShopNo"`

	// 县/区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Country *string `json:"Country,omitempty" name:"Country"`

	// 若是支付宝合同，支付宝审核状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	AliPayStatus *string `json:"AliPayStatus,omitempty" name:"AliPayStatus"`

	// 为空或者0表示未关联，大于0表示已关联
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsChecked *string `json:"IsChecked,omitempty" name:"IsChecked"`

	// 详细地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Address *string `json:"Address,omitempty" name:"Address"`

	// 若是支付宝合同，支付宝审核描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	AliPayDesc *string `json:"AliPayDesc,omitempty" name:"AliPayDesc"`
}

// Predefined struct for user
type QueryContractRequestParams struct {
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

// Predefined struct for user
type QueryContractResponseParams struct {
	// 签约数据
	ContractData *ResponseQueryContract `json:"ContractData,omitempty" name:"ContractData"`

	// 请求处理信息
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryContractResponse struct {
	*tchttp.BaseResponse
	Response *QueryContractResponseParams `json:"Response"`
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

// Predefined struct for user
type QueryCustAcctIdBalanceRequestParams struct {
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

// Predefined struct for user
type QueryCustAcctIdBalanceResponseParams struct {
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
}

type QueryCustAcctIdBalanceResponse struct {
	*tchttp.BaseResponse
	Response *QueryCustAcctIdBalanceResponseParams `json:"Response"`
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

// Predefined struct for user
type QueryDownloadBillURLRequestParams struct {
	// 分配给商户的AppId。进件成功后返给商户方的AppId。
	MerchantAppId *string `json:"MerchantAppId,omitempty" name:"MerchantAppId"`

	// 渠道编号。固定值：ZSB2B
	ChannelCode *string `json:"ChannelCode,omitempty" name:"ChannelCode"`

	// 对账单日期，格式yyyyMMdd
	BillDate *string `json:"BillDate,omitempty" name:"BillDate"`
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

// Predefined struct for user
type QueryDownloadBillURLResponseParams struct {
	// 分配给商户的AppId。进件成功后返给商户方的AppId。
	MerchantAppId *string `json:"MerchantAppId,omitempty" name:"MerchantAppId"`

	// 对账单下载地址。
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryDownloadBillURLResponse struct {
	*tchttp.BaseResponse
	Response *QueryDownloadBillURLResponseParams `json:"Response"`
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

type QueryExceedingInfoData struct {
	// 代理商ID。
	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`

	// 代理商名称。
	AgentName *string `json:"AgentName,omitempty" name:"AgentName"`

	// 主播ID。当入参Dimension为ANCHOR或ORDER时，该字段才会有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnchorId *string `json:"AnchorId,omitempty" name:"AnchorId"`

	// 主播名称。当入参Dimension为ANCHOR或ORDER时，该字段才会有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnchorName *string `json:"AnchorName,omitempty" name:"AnchorName"`

	// 订单号。当入参Dimension为ORDER时，该字段才会有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 超额类型。目前支持 AGENT_EXCEED_100 和 ANCHOR_EXCEED_100_12 两种类型。
	ExceedingType *string `json:"ExceedingType,omitempty" name:"ExceedingType"`
}

// Predefined struct for user
type QueryExceedingInfoRequestParams struct {
	// 超额日期。格式为yyyy-MM-dd。
	TimeStr *string `json:"TimeStr,omitempty" name:"TimeStr"`

	// 维度。目前支持三个维度: AGENT, ANCHOR, ORDER。不填默认使用AGENT维度。
	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`

	// 分页信息。不填默认Index为1，Count为100。
	PageNumber *Paging `json:"PageNumber,omitempty" name:"PageNumber"`
}

type QueryExceedingInfoRequest struct {
	*tchttp.BaseRequest
	
	// 超额日期。格式为yyyy-MM-dd。
	TimeStr *string `json:"TimeStr,omitempty" name:"TimeStr"`

	// 维度。目前支持三个维度: AGENT, ANCHOR, ORDER。不填默认使用AGENT维度。
	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`

	// 分页信息。不填默认Index为1，Count为100。
	PageNumber *Paging `json:"PageNumber,omitempty" name:"PageNumber"`
}

func (r *QueryExceedingInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryExceedingInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeStr")
	delete(f, "Dimension")
	delete(f, "PageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryExceedingInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryExceedingInfoResponseParams struct {
	// 错误码。
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 超额信息结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *QueryExceedingInfoResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryExceedingInfoResponse struct {
	*tchttp.BaseResponse
	Response *QueryExceedingInfoResponseParams `json:"Response"`
}

func (r *QueryExceedingInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryExceedingInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryExceedingInfoResult struct {
	// 记录总数。
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 超额信息数据。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*QueryExceedingInfoData `json:"Data,omitempty" name:"Data"`
}

// Predefined struct for user
type QueryExchangeRateRequestParams struct {
	// 源币种 (默认CNY)
	SourceCurrency *string `json:"SourceCurrency,omitempty" name:"SourceCurrency"`

	// 目的币种 (见常见问题-汇出币种)
	TargetCurrency *string `json:"TargetCurrency,omitempty" name:"TargetCurrency"`

	// 接入环境。沙箱环境填sandbox
	Profile *string `json:"Profile,omitempty" name:"Profile"`
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

// Predefined struct for user
type QueryExchangeRateResponseParams struct {
	// 查询汇率结果
	Result *QueryExchangerateResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryExchangeRateResponse struct {
	*tchttp.BaseResponse
	Response *QueryExchangeRateResponseParams `json:"Response"`
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

type QueryExternalAccountBookResult struct {
	// 渠道记账本ID
	ChannelAccountBookId *string `json:"ChannelAccountBookId,omitempty" name:"ChannelAccountBookId"`

	// 可用余额。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvailableBalance *string `json:"AvailableBalance,omitempty" name:"AvailableBalance"`

	// 电子记账本对外收款的账户信息。为JSON格式字符串（成功状态下返回）。详情见附录-复杂类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CollectMoneyAccountInfo *string `json:"CollectMoneyAccountInfo,omitempty" name:"CollectMoneyAccountInfo"`
}

// Predefined struct for user
type QueryFlexAmountBeforeTaxRequestParams struct {
	// 收款用户ID
	PayeeId *string `json:"PayeeId,omitempty" name:"PayeeId"`

	// 收入类型
	// LABOR:劳务所得
	// OCCASION:偶然所得
	IncomeType *string `json:"IncomeType,omitempty" name:"IncomeType"`

	// 税后金额
	AmountAfterTax *string `json:"AmountAfterTax,omitempty" name:"AmountAfterTax"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// __test__:测试环境
	// 缺省默认为生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type QueryFlexAmountBeforeTaxRequest struct {
	*tchttp.BaseRequest
	
	// 收款用户ID
	PayeeId *string `json:"PayeeId,omitempty" name:"PayeeId"`

	// 收入类型
	// LABOR:劳务所得
	// OCCASION:偶然所得
	IncomeType *string `json:"IncomeType,omitempty" name:"IncomeType"`

	// 税后金额
	AmountAfterTax *string `json:"AmountAfterTax,omitempty" name:"AmountAfterTax"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// __test__:测试环境
	// 缺省默认为生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *QueryFlexAmountBeforeTaxRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryFlexAmountBeforeTaxRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PayeeId")
	delete(f, "IncomeType")
	delete(f, "AmountAfterTax")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryFlexAmountBeforeTaxRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryFlexAmountBeforeTaxResponseParams struct {
	// 错误码。SUCCESS为成功，其他为失败
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *AmountBeforeTaxResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryFlexAmountBeforeTaxResponse struct {
	*tchttp.BaseResponse
	Response *QueryFlexAmountBeforeTaxResponseParams `json:"Response"`
}

func (r *QueryFlexAmountBeforeTaxResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryFlexAmountBeforeTaxResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryFlexBillDownloadUrlRequestParams struct {
	// 对账单日期
	BillDate *string `json:"BillDate,omitempty" name:"BillDate"`

	// 对账单类型：FREEZE, SETTLEMENT,PAYMENT
	BillType *string `json:"BillType,omitempty" name:"BillType"`

	// 服务商ID，如不填则查询平台级别对账单文件
	ServiceProviderId *string `json:"ServiceProviderId,omitempty" name:"ServiceProviderId"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// __test__:测试环境
	// 缺省默认为生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type QueryFlexBillDownloadUrlRequest struct {
	*tchttp.BaseRequest
	
	// 对账单日期
	BillDate *string `json:"BillDate,omitempty" name:"BillDate"`

	// 对账单类型：FREEZE, SETTLEMENT,PAYMENT
	BillType *string `json:"BillType,omitempty" name:"BillType"`

	// 服务商ID，如不填则查询平台级别对账单文件
	ServiceProviderId *string `json:"ServiceProviderId,omitempty" name:"ServiceProviderId"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// __test__:测试环境
	// 缺省默认为生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *QueryFlexBillDownloadUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryFlexBillDownloadUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BillDate")
	delete(f, "BillType")
	delete(f, "ServiceProviderId")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryFlexBillDownloadUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryFlexBillDownloadUrlResponseParams struct {
	// 错误码。SUCCESS为成功，其他为失败
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *FlexBillDownloadUrlResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryFlexBillDownloadUrlResponse struct {
	*tchttp.BaseResponse
	Response *QueryFlexBillDownloadUrlResponseParams `json:"Response"`
}

func (r *QueryFlexBillDownloadUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryFlexBillDownloadUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryFlexFreezeOrderListRequestParams struct {
	// 收款用户ID
	PayeeId *string `json:"PayeeId,omitempty" name:"PayeeId"`

	// 操作类型
	// FREEZE:冻结
	// UNFREEZE:解冻
	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`

	// 开始时间，格式"yyyy-MM-dd hh:mm:ss"
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，格式"yyyy-MM-dd hh:mm:ss"
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分页
	PageNumber *Paging `json:"PageNumber,omitempty" name:"PageNumber"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// __test__:测试环境
	// 缺省默认为生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type QueryFlexFreezeOrderListRequest struct {
	*tchttp.BaseRequest
	
	// 收款用户ID
	PayeeId *string `json:"PayeeId,omitempty" name:"PayeeId"`

	// 操作类型
	// FREEZE:冻结
	// UNFREEZE:解冻
	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`

	// 开始时间，格式"yyyy-MM-dd hh:mm:ss"
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，格式"yyyy-MM-dd hh:mm:ss"
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分页
	PageNumber *Paging `json:"PageNumber,omitempty" name:"PageNumber"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// __test__:测试环境
	// 缺省默认为生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *QueryFlexFreezeOrderListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryFlexFreezeOrderListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PayeeId")
	delete(f, "OperationType")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageNumber")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryFlexFreezeOrderListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryFlexFreezeOrderListResponseParams struct {
	// 错误码。SUCCESS为成功，其他为失败
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *FreezeOrders `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryFlexFreezeOrderListResponse struct {
	*tchttp.BaseResponse
	Response *QueryFlexFreezeOrderListResponseParams `json:"Response"`
}

func (r *QueryFlexFreezeOrderListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryFlexFreezeOrderListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryFlexOrderSummaryListRequestParams struct {
	// 汇总日期:yyyy-MM-dd
	SummaryDate *string `json:"SummaryDate,omitempty" name:"SummaryDate"`

	// 分页
	PageNumber *Paging `json:"PageNumber,omitempty" name:"PageNumber"`

	// 汇总订单类型:FREEZE, SETTLEMENT,PAYMENT
	OrderType *string `json:"OrderType,omitempty" name:"OrderType"`

	// 收款用户ID
	PayeeId *string `json:"PayeeId,omitempty" name:"PayeeId"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// __test__:测试环境
	// 缺省默认为生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type QueryFlexOrderSummaryListRequest struct {
	*tchttp.BaseRequest
	
	// 汇总日期:yyyy-MM-dd
	SummaryDate *string `json:"SummaryDate,omitempty" name:"SummaryDate"`

	// 分页
	PageNumber *Paging `json:"PageNumber,omitempty" name:"PageNumber"`

	// 汇总订单类型:FREEZE, SETTLEMENT,PAYMENT
	OrderType *string `json:"OrderType,omitempty" name:"OrderType"`

	// 收款用户ID
	PayeeId *string `json:"PayeeId,omitempty" name:"PayeeId"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// __test__:测试环境
	// 缺省默认为生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *QueryFlexOrderSummaryListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryFlexOrderSummaryListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SummaryDate")
	delete(f, "PageNumber")
	delete(f, "OrderType")
	delete(f, "PayeeId")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryFlexOrderSummaryListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryFlexOrderSummaryListResponseParams struct {
	// 错误码。SUCCESS为成功，其他为失败
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *OrderSummaries `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryFlexOrderSummaryListResponse struct {
	*tchttp.BaseResponse
	Response *QueryFlexOrderSummaryListResponseParams `json:"Response"`
}

func (r *QueryFlexOrderSummaryListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryFlexOrderSummaryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryFlexPayeeAccountBalanceRequestParams struct {
	// 收款用户ID
	PayeeId *string `json:"PayeeId,omitempty" name:"PayeeId"`

	// 收入类型
	// LABOR:劳务所得
	// OCCASION:偶然所得
	IncomeType *string `json:"IncomeType,omitempty" name:"IncomeType"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// __test__:测试环境
	// 缺省默认为生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type QueryFlexPayeeAccountBalanceRequest struct {
	*tchttp.BaseRequest
	
	// 收款用户ID
	PayeeId *string `json:"PayeeId,omitempty" name:"PayeeId"`

	// 收入类型
	// LABOR:劳务所得
	// OCCASION:偶然所得
	IncomeType *string `json:"IncomeType,omitempty" name:"IncomeType"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// __test__:测试环境
	// 缺省默认为生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *QueryFlexPayeeAccountBalanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryFlexPayeeAccountBalanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PayeeId")
	delete(f, "IncomeType")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryFlexPayeeAccountBalanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryFlexPayeeAccountBalanceResponseParams struct {
	// 错误码。SUCCESS为成功，其他为失败
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *PayeeAccountBalanceResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryFlexPayeeAccountBalanceResponse struct {
	*tchttp.BaseResponse
	Response *QueryFlexPayeeAccountBalanceResponseParams `json:"Response"`
}

func (r *QueryFlexPayeeAccountBalanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryFlexPayeeAccountBalanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryFlexPayeeAccountInfoRequestParams struct {
	// 收款用户ID
	PayeeId *string `json:"PayeeId,omitempty" name:"PayeeId"`

	// 外部用户ID
	OutUserId *string `json:"OutUserId,omitempty" name:"OutUserId"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// __test__:测试环境
	// 缺省默认为生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type QueryFlexPayeeAccountInfoRequest struct {
	*tchttp.BaseRequest
	
	// 收款用户ID
	PayeeId *string `json:"PayeeId,omitempty" name:"PayeeId"`

	// 外部用户ID
	OutUserId *string `json:"OutUserId,omitempty" name:"OutUserId"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// __test__:测试环境
	// 缺省默认为生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *QueryFlexPayeeAccountInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryFlexPayeeAccountInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PayeeId")
	delete(f, "OutUserId")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryFlexPayeeAccountInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryFlexPayeeAccountInfoResponseParams struct {
	// 错误码。SUCCESS为成功，其他为失败
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *PayeeAccountInfoResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryFlexPayeeAccountInfoResponse struct {
	*tchttp.BaseResponse
	Response *QueryFlexPayeeAccountInfoResponseParams `json:"Response"`
}

func (r *QueryFlexPayeeAccountInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryFlexPayeeAccountInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryFlexPayeeAccountListRequestParams struct {
	// 账户属性信息
	PropertyInfo *PayeeAccountPropertyInfo `json:"PropertyInfo,omitempty" name:"PropertyInfo"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分页
	PageNumber *Paging `json:"PageNumber,omitempty" name:"PageNumber"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// __test__:测试环境
	// 缺省默认为生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type QueryFlexPayeeAccountListRequest struct {
	*tchttp.BaseRequest
	
	// 账户属性信息
	PropertyInfo *PayeeAccountPropertyInfo `json:"PropertyInfo,omitempty" name:"PropertyInfo"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分页
	PageNumber *Paging `json:"PageNumber,omitempty" name:"PageNumber"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// __test__:测试环境
	// 缺省默认为生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *QueryFlexPayeeAccountListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryFlexPayeeAccountListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PropertyInfo")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageNumber")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryFlexPayeeAccountListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryFlexPayeeAccountListResponseParams struct {
	// 错误码。SUCCESS为成功，其他为失败
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *PayeeAccountInfos `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryFlexPayeeAccountListResponse struct {
	*tchttp.BaseResponse
	Response *QueryFlexPayeeAccountListResponseParams `json:"Response"`
}

func (r *QueryFlexPayeeAccountListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryFlexPayeeAccountListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryFlexPayeeInfoRequestParams struct {
	// 收款用户ID
	PayeeId *string `json:"PayeeId,omitempty" name:"PayeeId"`

	// 外部用户ID
	OutUserId *string `json:"OutUserId,omitempty" name:"OutUserId"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// __test__:测试环境
	// 缺省默认为生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type QueryFlexPayeeInfoRequest struct {
	*tchttp.BaseRequest
	
	// 收款用户ID
	PayeeId *string `json:"PayeeId,omitempty" name:"PayeeId"`

	// 外部用户ID
	OutUserId *string `json:"OutUserId,omitempty" name:"OutUserId"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// __test__:测试环境
	// 缺省默认为生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *QueryFlexPayeeInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryFlexPayeeInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PayeeId")
	delete(f, "OutUserId")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryFlexPayeeInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryFlexPayeeInfoResponseParams struct {
	// 错误码。SUCCESS为成功，其他为失败
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *PayeeInfoResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryFlexPayeeInfoResponse struct {
	*tchttp.BaseResponse
	Response *QueryFlexPayeeInfoResponseParams `json:"Response"`
}

func (r *QueryFlexPayeeInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryFlexPayeeInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryFlexPaymentOrderListRequestParams struct {
	// 收款用户ID
	PayeeId *string `json:"PayeeId,omitempty" name:"PayeeId"`

	// 开始时间，格式"yyyy-MM-dd hh:mm:ss"
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，格式"yyyy-MM-dd hh:mm:ss"
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分页
	PageNumber *Paging `json:"PageNumber,omitempty" name:"PageNumber"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// __test__:测试环境
	// 缺省默认为生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type QueryFlexPaymentOrderListRequest struct {
	*tchttp.BaseRequest
	
	// 收款用户ID
	PayeeId *string `json:"PayeeId,omitempty" name:"PayeeId"`

	// 开始时间，格式"yyyy-MM-dd hh:mm:ss"
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，格式"yyyy-MM-dd hh:mm:ss"
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分页
	PageNumber *Paging `json:"PageNumber,omitempty" name:"PageNumber"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// __test__:测试环境
	// 缺省默认为生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *QueryFlexPaymentOrderListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryFlexPaymentOrderListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PayeeId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageNumber")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryFlexPaymentOrderListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryFlexPaymentOrderListResponseParams struct {
	// 错误码。SUCCESS为成功，其他为失败
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *PaymentOrders `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryFlexPaymentOrderListResponse struct {
	*tchttp.BaseResponse
	Response *QueryFlexPaymentOrderListResponseParams `json:"Response"`
}

func (r *QueryFlexPaymentOrderListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryFlexPaymentOrderListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryFlexPaymentOrderStatusRequestParams struct {
	// 外部订单ID
	OutOrderId *string `json:"OutOrderId,omitempty" name:"OutOrderId"`

	// 订单ID
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// __test__:测试环境
	// 缺省默认为生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type QueryFlexPaymentOrderStatusRequest struct {
	*tchttp.BaseRequest
	
	// 外部订单ID
	OutOrderId *string `json:"OutOrderId,omitempty" name:"OutOrderId"`

	// 订单ID
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// __test__:测试环境
	// 缺省默认为生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *QueryFlexPaymentOrderStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryFlexPaymentOrderStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OutOrderId")
	delete(f, "OrderId")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryFlexPaymentOrderStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryFlexPaymentOrderStatusResponseParams struct {
	// 错误码。SUCCESS为成功，其他为失败
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *PaymentOrderStatusResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryFlexPaymentOrderStatusResponse struct {
	*tchttp.BaseResponse
	Response *QueryFlexPaymentOrderStatusResponseParams `json:"Response"`
}

func (r *QueryFlexPaymentOrderStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryFlexPaymentOrderStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryFlexSettlementOrderListRequestParams struct {
	// 收款用户ID
	PayeeId *string `json:"PayeeId,omitempty" name:"PayeeId"`

	// 开始时间，格式"yyyy-MM-dd hh:mm:ss"
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，格式"yyyy-MM-dd hh:mm:ss"
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分页
	PageNumber *Paging `json:"PageNumber,omitempty" name:"PageNumber"`

	// 操作类型。
	// ENABLE_SETTLE: 正常结算
	// DISABLE_SETTLE: 停用结算
	// UNFREEZE_SETTLE: 解冻结算
	// 若需要支持多个操作类型，则以;分隔
	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// __test__:测试环境
	// 缺省默认为生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type QueryFlexSettlementOrderListRequest struct {
	*tchttp.BaseRequest
	
	// 收款用户ID
	PayeeId *string `json:"PayeeId,omitempty" name:"PayeeId"`

	// 开始时间，格式"yyyy-MM-dd hh:mm:ss"
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，格式"yyyy-MM-dd hh:mm:ss"
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分页
	PageNumber *Paging `json:"PageNumber,omitempty" name:"PageNumber"`

	// 操作类型。
	// ENABLE_SETTLE: 正常结算
	// DISABLE_SETTLE: 停用结算
	// UNFREEZE_SETTLE: 解冻结算
	// 若需要支持多个操作类型，则以;分隔
	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// __test__:测试环境
	// 缺省默认为生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *QueryFlexSettlementOrderListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryFlexSettlementOrderListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PayeeId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageNumber")
	delete(f, "OperationType")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryFlexSettlementOrderListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryFlexSettlementOrderListResponseParams struct {
	// 错误码。SUCCESS为成功，其他为失败
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *SettlementOrders `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryFlexSettlementOrderListResponse struct {
	*tchttp.BaseResponse
	Response *QueryFlexSettlementOrderListResponseParams `json:"Response"`
}

func (r *QueryFlexSettlementOrderListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryFlexSettlementOrderListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryFundsTransactionDetailsRequestParams struct {
	// 查询的交易发生时间类型。
	// __1__：当日
	// __2__：历史
	QueryDateType *string `json:"QueryDateType,omitempty" name:"QueryDateType"`

	// 查询的交易类型。
	// __2__：提现/退款
	// __3__：清分/充值
	QueryTranType *string `json:"QueryTranType,omitempty" name:"QueryTranType"`

	// 父账户账号。
	// _平安渠道为资金汇总账号_
	BankAccountNumber *string `json:"BankAccountNumber,omitempty" name:"BankAccountNumber"`

	// 子账户账号。
	// _平安渠道为见证子账户的账号_
	SubAccountNumber *string `json:"SubAccountNumber,omitempty" name:"SubAccountNumber"`

	// 分页号, 起始值为1。
	PageOffSet *string `json:"PageOffSet,omitempty" name:"PageOffSet"`

	// 查询开始日期，格式：yyyyMMdd。
	// __若是历史查询，则必输，当日查询时，不起作用；开始日期不能超过当前日期__
	QueryStartDate *string `json:"QueryStartDate,omitempty" name:"QueryStartDate"`

	// 查询终止日期，格式：yyyyMMdd。
	// __若是历史查询，则必输，当日查询时，不起作用；终止日期不能超过当前日期__
	QueryEndDate *string `json:"QueryEndDate,omitempty" name:"QueryEndDate"`

	// 环境名。
	// __release__: 现网环境
	// __sandbox__: 沙箱环境
	// __development__: 开发环境
	// _缺省: release_
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`
}

type QueryFundsTransactionDetailsRequest struct {
	*tchttp.BaseRequest
	
	// 查询的交易发生时间类型。
	// __1__：当日
	// __2__：历史
	QueryDateType *string `json:"QueryDateType,omitempty" name:"QueryDateType"`

	// 查询的交易类型。
	// __2__：提现/退款
	// __3__：清分/充值
	QueryTranType *string `json:"QueryTranType,omitempty" name:"QueryTranType"`

	// 父账户账号。
	// _平安渠道为资金汇总账号_
	BankAccountNumber *string `json:"BankAccountNumber,omitempty" name:"BankAccountNumber"`

	// 子账户账号。
	// _平安渠道为见证子账户的账号_
	SubAccountNumber *string `json:"SubAccountNumber,omitempty" name:"SubAccountNumber"`

	// 分页号, 起始值为1。
	PageOffSet *string `json:"PageOffSet,omitempty" name:"PageOffSet"`

	// 查询开始日期，格式：yyyyMMdd。
	// __若是历史查询，则必输，当日查询时，不起作用；开始日期不能超过当前日期__
	QueryStartDate *string `json:"QueryStartDate,omitempty" name:"QueryStartDate"`

	// 查询终止日期，格式：yyyyMMdd。
	// __若是历史查询，则必输，当日查询时，不起作用；终止日期不能超过当前日期__
	QueryEndDate *string `json:"QueryEndDate,omitempty" name:"QueryEndDate"`

	// 环境名。
	// __release__: 现网环境
	// __sandbox__: 沙箱环境
	// __development__: 开发环境
	// _缺省: release_
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`
}

func (r *QueryFundsTransactionDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryFundsTransactionDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "QueryDateType")
	delete(f, "QueryTranType")
	delete(f, "BankAccountNumber")
	delete(f, "SubAccountNumber")
	delete(f, "PageOffSet")
	delete(f, "QueryStartDate")
	delete(f, "QueryEndDate")
	delete(f, "MidasEnvironment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryFundsTransactionDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryFundsTransactionDetailsResponseParams struct {
	// 错误码。
	// __SUCCESS__: 成功
	// __其他__: 见附录-错误码表
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *QueryFundsTransactionDetailsResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryFundsTransactionDetailsResponse struct {
	*tchttp.BaseResponse
	Response *QueryFundsTransactionDetailsResponseParams `json:"Response"`
}

func (r *QueryFundsTransactionDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryFundsTransactionDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryFundsTransactionDetailsResult struct {
	// 本次交易返回查询结果记录数。
	ResultCount *uint64 `json:"ResultCount,omitempty" name:"ResultCount"`

	// 符合业务查询条件的记录总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 结束标志。
	// __0__：否
	// __1__：是
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndFlag *string `json:"EndFlag,omitempty" name:"EndFlag"`

	// 会员资金交易信息数组。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranItemArray []*FundsTransactionItem `json:"TranItemArray,omitempty" name:"TranItemArray"`
}

// Predefined struct for user
type QueryInvoiceRequestParams struct {
	// 开票平台ID
	// 0 : 高灯
	// 1 : 票易通
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

type QueryInvoiceRequest struct {
	*tchttp.BaseRequest
	
	// 开票平台ID
	// 0 : 高灯
	// 1 : 票易通
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

// Predefined struct for user
type QueryInvoiceResponseParams struct {
	// 发票查询结果
	Result *QueryInvoiceResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryInvoiceResponse struct {
	*tchttp.BaseResponse
	Response *QueryInvoiceResponseParams `json:"Response"`
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

	// png地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

// Predefined struct for user
type QueryInvoiceV2RequestParams struct {
	// 开票平台ID
	// 0 : 高灯
	// 1 : 票易通
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

type QueryInvoiceV2Request struct {
	*tchttp.BaseRequest
	
	// 开票平台ID
	// 0 : 高灯
	// 1 : 票易通
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

// Predefined struct for user
type QueryInvoiceV2ResponseParams struct {
	// 发票查询结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *QueryInvoiceResultData `json:"Result,omitempty" name:"Result"`

	// 错误码
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryInvoiceV2Response struct {
	*tchttp.BaseResponse
	Response *QueryInvoiceV2ResponseParams `json:"Response"`
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

// Predefined struct for user
type QueryMaliciousRegistrationRequestParams struct {
	// 商户ID，调用方使用的商户号信息，与商户主体一一对应
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 商户名称
	MerchantName *string `json:"MerchantName,omitempty" name:"MerchantName"`

	// 企业工商注册标准名称
	CompanyName *string `json:"CompanyName,omitempty" name:"CompanyName"`

	// 注册地址
	RegAddress *string `json:"RegAddress,omitempty" name:"RegAddress"`

	// 商户进件Unix时间，单位秒（非企业注册工商时间)
	RegTime *uint64 `json:"RegTime,omitempty" name:"RegTime"`

	// 统一社会信用代码
	USCI *string `json:"USCI,omitempty" name:"USCI"`

	// 工商注册码，匹配优先级为Usci>RegNumber>CompanyName
	RegNumber *string `json:"RegNumber,omitempty" name:"RegNumber"`

	// 手机号码32位MD5加密结果，全大写，格式为0086-13812345678
	EncryptedPhoneNumber *string `json:"EncryptedPhoneNumber,omitempty" name:"EncryptedPhoneNumber"`

	// 邮箱32位MD5加密结果，全大写
	EncryptedEmailAddress *string `json:"EncryptedEmailAddress,omitempty" name:"EncryptedEmailAddress"`

	// 身份证MD5加密结果，最后一位x大写
	EncryptedPersonId *string `json:"EncryptedPersonId,omitempty" name:"EncryptedPersonId"`

	// 填写信息设备的IP地址
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 进件渠道号，客户自行编码即可
	Channel *string `json:"Channel,omitempty" name:"Channel"`
}

type QueryMaliciousRegistrationRequest struct {
	*tchttp.BaseRequest
	
	// 商户ID，调用方使用的商户号信息，与商户主体一一对应
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 商户名称
	MerchantName *string `json:"MerchantName,omitempty" name:"MerchantName"`

	// 企业工商注册标准名称
	CompanyName *string `json:"CompanyName,omitempty" name:"CompanyName"`

	// 注册地址
	RegAddress *string `json:"RegAddress,omitempty" name:"RegAddress"`

	// 商户进件Unix时间，单位秒（非企业注册工商时间)
	RegTime *uint64 `json:"RegTime,omitempty" name:"RegTime"`

	// 统一社会信用代码
	USCI *string `json:"USCI,omitempty" name:"USCI"`

	// 工商注册码，匹配优先级为Usci>RegNumber>CompanyName
	RegNumber *string `json:"RegNumber,omitempty" name:"RegNumber"`

	// 手机号码32位MD5加密结果，全大写，格式为0086-13812345678
	EncryptedPhoneNumber *string `json:"EncryptedPhoneNumber,omitempty" name:"EncryptedPhoneNumber"`

	// 邮箱32位MD5加密结果，全大写
	EncryptedEmailAddress *string `json:"EncryptedEmailAddress,omitempty" name:"EncryptedEmailAddress"`

	// 身份证MD5加密结果，最后一位x大写
	EncryptedPersonId *string `json:"EncryptedPersonId,omitempty" name:"EncryptedPersonId"`

	// 填写信息设备的IP地址
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 进件渠道号，客户自行编码即可
	Channel *string `json:"Channel,omitempty" name:"Channel"`
}

func (r *QueryMaliciousRegistrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryMaliciousRegistrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MerchantId")
	delete(f, "MerchantName")
	delete(f, "CompanyName")
	delete(f, "RegAddress")
	delete(f, "RegTime")
	delete(f, "USCI")
	delete(f, "RegNumber")
	delete(f, "EncryptedPhoneNumber")
	delete(f, "EncryptedEmailAddress")
	delete(f, "EncryptedPersonId")
	delete(f, "Ip")
	delete(f, "Channel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryMaliciousRegistrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryMaliciousRegistrationResponseParams struct {
	// 错误码
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 商户风险信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *MerchantRiskInfo `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryMaliciousRegistrationResponse struct {
	*tchttp.BaseResponse
	Response *QueryMaliciousRegistrationResponseParams `json:"Response"`
}

func (r *QueryMaliciousRegistrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryMaliciousRegistrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryMemberBindRequestParams struct {
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

// Predefined struct for user
type QueryMemberBindResponseParams struct {
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
}

type QueryMemberBindResponse struct {
	*tchttp.BaseResponse
	Response *QueryMemberBindResponseParams `json:"Response"`
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

// Predefined struct for user
type QueryMemberTransactionDetailsRequestParams struct {
	// 查询的交易发生时间类型。
	// __1__：当日
	// __2__：历史
	QueryDateType *string `json:"QueryDateType,omitempty" name:"QueryDateType"`

	// 查询的交易类型。
	// __1__：全部
	// __2__：转出
	// __3__：转入
	QueryTranType *string `json:"QueryTranType,omitempty" name:"QueryTranType"`

	// 父账户账号。
	// _平安渠道为资金汇总账号_
	BankAccountNumber *string `json:"BankAccountNumber,omitempty" name:"BankAccountNumber"`

	// 子账户账号。
	// _平安渠道为见证子账户的账号_
	SubAccountNumber *string `json:"SubAccountNumber,omitempty" name:"SubAccountNumber"`

	// 分页号, 起始值为1。
	PageOffSet *string `json:"PageOffSet,omitempty" name:"PageOffSet"`

	// 查询开始日期，格式：yyyyMMdd。
	// __若是历史查询，则必输，当日查询时，不起作用；开始日期不能超过当前日期__
	QueryStartDate *string `json:"QueryStartDate,omitempty" name:"QueryStartDate"`

	// 查询终止日期，格式：yyyyMMdd。
	// __若是历史查询，则必输，当日查询时，不起作用；终止日期不能超过当前日期__
	QueryEndDate *string `json:"QueryEndDate,omitempty" name:"QueryEndDate"`

	// 环境名。
	// __release__: 现网环境
	// __sandbox__: 沙箱环境
	// __development__: 开发环境
	// _缺省: release_
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`
}

type QueryMemberTransactionDetailsRequest struct {
	*tchttp.BaseRequest
	
	// 查询的交易发生时间类型。
	// __1__：当日
	// __2__：历史
	QueryDateType *string `json:"QueryDateType,omitempty" name:"QueryDateType"`

	// 查询的交易类型。
	// __1__：全部
	// __2__：转出
	// __3__：转入
	QueryTranType *string `json:"QueryTranType,omitempty" name:"QueryTranType"`

	// 父账户账号。
	// _平安渠道为资金汇总账号_
	BankAccountNumber *string `json:"BankAccountNumber,omitempty" name:"BankAccountNumber"`

	// 子账户账号。
	// _平安渠道为见证子账户的账号_
	SubAccountNumber *string `json:"SubAccountNumber,omitempty" name:"SubAccountNumber"`

	// 分页号, 起始值为1。
	PageOffSet *string `json:"PageOffSet,omitempty" name:"PageOffSet"`

	// 查询开始日期，格式：yyyyMMdd。
	// __若是历史查询，则必输，当日查询时，不起作用；开始日期不能超过当前日期__
	QueryStartDate *string `json:"QueryStartDate,omitempty" name:"QueryStartDate"`

	// 查询终止日期，格式：yyyyMMdd。
	// __若是历史查询，则必输，当日查询时，不起作用；终止日期不能超过当前日期__
	QueryEndDate *string `json:"QueryEndDate,omitempty" name:"QueryEndDate"`

	// 环境名。
	// __release__: 现网环境
	// __sandbox__: 沙箱环境
	// __development__: 开发环境
	// _缺省: release_
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`
}

func (r *QueryMemberTransactionDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryMemberTransactionDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "QueryDateType")
	delete(f, "QueryTranType")
	delete(f, "BankAccountNumber")
	delete(f, "SubAccountNumber")
	delete(f, "PageOffSet")
	delete(f, "QueryStartDate")
	delete(f, "QueryEndDate")
	delete(f, "MidasEnvironment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryMemberTransactionDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryMemberTransactionDetailsResponseParams struct {
	// 错误码。
	// __SUCCESS__: 成功
	// __其他__: 见附录-错误码表
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *QueryMemberTransactionDetailsResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryMemberTransactionDetailsResponse struct {
	*tchttp.BaseResponse
	Response *QueryMemberTransactionDetailsResponseParams `json:"Response"`
}

func (r *QueryMemberTransactionDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryMemberTransactionDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryMemberTransactionDetailsResult struct {
	// 本次交易返回查询结果记录数。
	ResultCount *uint64 `json:"ResultCount,omitempty" name:"ResultCount"`

	// 符合业务查询条件的记录总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 结束标志。
	// __0__：否
	// __1__：是
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndFlag *string `json:"EndFlag,omitempty" name:"EndFlag"`

	// 会员间交易信息数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranItemArray []*MemberTransactionItem `json:"TranItemArray,omitempty" name:"TranItemArray"`
}

// Predefined struct for user
type QueryMemberTransactionRequestParams struct {
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

// Predefined struct for user
type QueryMemberTransactionResponseParams struct {
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
}

type QueryMemberTransactionResponse struct {
	*tchttp.BaseResponse
	Response *QueryMemberTransactionResponseParams `json:"Response"`
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

// Predefined struct for user
type QueryMerchantBalanceRequestParams struct {
	// 余额币种
	Currency *string `json:"Currency,omitempty" name:"Currency"`

	// 接入环境。沙箱环境填sandbox
	Profile *string `json:"Profile,omitempty" name:"Profile"`
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

// Predefined struct for user
type QueryMerchantBalanceResponseParams struct {
	// 对接方账户余额查询结果
	Result *QueryMerchantBalanceResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryMerchantBalanceResponse struct {
	*tchttp.BaseResponse
	Response *QueryMerchantBalanceResponseParams `json:"Response"`
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

// Predefined struct for user
type QueryMerchantClassificationRequestParams struct {
	// 收单系统分配的开放ID
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 收单系统分配的密钥
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

type QueryMerchantClassificationRequest struct {
	*tchttp.BaseRequest
	
	// 收单系统分配的开放ID
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 收单系统分配的密钥
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *QueryMerchantClassificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryMerchantClassificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OpenId")
	delete(f, "OpenKey")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryMerchantClassificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryMerchantClassificationResponseParams struct {
	// 业务系统返回消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 业务系统返回码
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 查询商户分类响应对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result []*MerchantClassificationId `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryMerchantClassificationResponse struct {
	*tchttp.BaseResponse
	Response *QueryMerchantClassificationResponseParams `json:"Response"`
}

func (r *QueryMerchantClassificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryMerchantClassificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryMerchantInfoForManagementRequestParams struct {
	// 开票平台ID
	InvoicePlatformId *int64 `json:"InvoicePlatformId,omitempty" name:"InvoicePlatformId"`

	// 页码
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 页大小
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 接入环境。沙箱环境填sandbox。
	Profile *string `json:"Profile,omitempty" name:"Profile"`
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

// Predefined struct for user
type QueryMerchantInfoForManagementResponseParams struct {
	// 商户结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *MerchantManagementResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryMerchantInfoForManagementResponse struct {
	*tchttp.BaseResponse
	Response *QueryMerchantInfoForManagementResponseParams `json:"Response"`
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

// Predefined struct for user
type QueryMerchantOrderRequestParams struct {
	// 进件成功后返给商户方的AppId。
	MerchantAppId *string `json:"MerchantAppId,omitempty" name:"MerchantAppId"`

	// 平台流水号。平台唯一订单号。
	OrderNo *string `json:"OrderNo,omitempty" name:"OrderNo"`
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

// Predefined struct for user
type QueryMerchantOrderResponseParams struct {
	// 进件成功后返给商户方的AppId。
	MerchantAppId *string `json:"MerchantAppId,omitempty" name:"MerchantAppId"`

	// 平台流水号。平台唯一订单号。
	OrderNo *string `json:"OrderNo,omitempty" name:"OrderNo"`

	// 订单支付状态。0-下单失败 1-下单成功未支付 2-支付成功 3-支付失败 4-退款中 5-退款成功 6-退款失败 7-待付款 8-待确认。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryMerchantOrderResponse struct {
	*tchttp.BaseResponse
	Response *QueryMerchantOrderResponseParams `json:"Response"`
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

// Predefined struct for user
type QueryMerchantPayWayListRequestParams struct {
	// 使用门店OpenId
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 使用门店OpenKey
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 支付类型，逗号分隔。1-现金，2-主扫，3-被扫，4-JSAPI。
	PayType *string `json:"PayType,omitempty" name:"PayType"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

type QueryMerchantPayWayListRequest struct {
	*tchttp.BaseRequest
	
	// 使用门店OpenId
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 使用门店OpenKey
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 支付类型，逗号分隔。1-现金，2-主扫，3-被扫，4-JSAPI。
	PayType *string `json:"PayType,omitempty" name:"PayType"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *QueryMerchantPayWayListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryMerchantPayWayListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OpenId")
	delete(f, "OpenKey")
	delete(f, "PayType")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryMerchantPayWayListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryMerchantPayWayListResponseParams struct {
	// 业务系统返回码，0表示成功，其他表示失败。
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 业务系统返回消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 查询商户支付方式列表结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result []*MerchantPayWayData `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryMerchantPayWayListResponse struct {
	*tchttp.BaseResponse
	Response *QueryMerchantPayWayListResponseParams `json:"Response"`
}

func (r *QueryMerchantPayWayListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryMerchantPayWayListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryMerchantRequestParams struct {
	// 进件成功后返给商户方的 AppId
	MerchantAppId *string `json:"MerchantAppId,omitempty" name:"MerchantAppId"`
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

// Predefined struct for user
type QueryMerchantResponseParams struct {
	// 分配给商户的 AppId，该 AppId 为后续各项 交易的商户标识。
	MerchantAppId *string `json:"MerchantAppId,omitempty" name:"MerchantAppId"`

	// 收款商户名称。
	MerchantName *string `json:"MerchantName,omitempty" name:"MerchantName"`

	// B2B 支付标志。是否开通 B2B 支付， 1:开通 0:不开通。
	BusinessPayFlag *string `json:"BusinessPayFlag,omitempty" name:"BusinessPayFlag"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryMerchantResponse struct {
	*tchttp.BaseResponse
	Response *QueryMerchantResponseParams `json:"Response"`
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

// Predefined struct for user
type QueryOpenBankBankAccountBalanceRequestParams struct {
	// 云企付渠道商户号。外部接入平台入驻云企付平台后下发。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 云企付渠道子商户号。入驻在渠道商户下的子商户ID，如付款方的商户ID，对应创建支付订单中接口参数中的PayerInfo中的payerId。
	ChannelSubMerchantId *string `json:"ChannelSubMerchantId,omitempty" name:"ChannelSubMerchantId"`

	// 渠道名称。
	// __TENPAY__: 商企付
	// __WECHAT__: 微信支付
	// __ALIPAY__: 支付宝
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 支付方式，如
	// __EBANK_PAYMENT__:ebank付款
	// __OPENBANK_PAYMENT__: openbank付款
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`

	// 绑卡序列号，银行账户唯一ID，区分多卡或多账户的场景
	BindSerialNo *string `json:"BindSerialNo,omitempty" name:"BindSerialNo"`

	// 环境类型
	// release:生产环境
	// sandbox:沙箱环境
	// 缺省默认为生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type QueryOpenBankBankAccountBalanceRequest struct {
	*tchttp.BaseRequest
	
	// 云企付渠道商户号。外部接入平台入驻云企付平台后下发。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 云企付渠道子商户号。入驻在渠道商户下的子商户ID，如付款方的商户ID，对应创建支付订单中接口参数中的PayerInfo中的payerId。
	ChannelSubMerchantId *string `json:"ChannelSubMerchantId,omitempty" name:"ChannelSubMerchantId"`

	// 渠道名称。
	// __TENPAY__: 商企付
	// __WECHAT__: 微信支付
	// __ALIPAY__: 支付宝
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 支付方式，如
	// __EBANK_PAYMENT__:ebank付款
	// __OPENBANK_PAYMENT__: openbank付款
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`

	// 绑卡序列号，银行账户唯一ID，区分多卡或多账户的场景
	BindSerialNo *string `json:"BindSerialNo,omitempty" name:"BindSerialNo"`

	// 环境类型
	// release:生产环境
	// sandbox:沙箱环境
	// 缺省默认为生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *QueryOpenBankBankAccountBalanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryOpenBankBankAccountBalanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelMerchantId")
	delete(f, "ChannelSubMerchantId")
	delete(f, "ChannelName")
	delete(f, "PaymentMethod")
	delete(f, "BindSerialNo")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryOpenBankBankAccountBalanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryOpenBankBankAccountBalanceResponseParams struct {
	// 业务系统返回码，SUCCESS表示成功，其他表示失败。
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 业务系统返回消息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 账户余额查询响应对象。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *QueryOpenBankBankAccountBalanceResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryOpenBankBankAccountBalanceResponse struct {
	*tchttp.BaseResponse
	Response *QueryOpenBankBankAccountBalanceResponseParams `json:"Response"`
}

func (r *QueryOpenBankBankAccountBalanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryOpenBankBankAccountBalanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryOpenBankBankAccountBalanceResult struct {
	// 总余额，单位分
	TotalBalance *string `json:"TotalBalance,omitempty" name:"TotalBalance"`

	// 昨日余额，单位分
	YesterdayBalance *string `json:"YesterdayBalance,omitempty" name:"YesterdayBalance"`
}

// Predefined struct for user
type QueryOpenBankBankBranchListRequestParams struct {
	// 渠道商户ID。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 渠道名称。
	// __TENPAY__: 商企付
	// __WECHAT__: 微信支付
	// __ALIPAY__: 支付宝
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 支付方式。
	// __EBANK_PAYMENT__:ebank付款
	// __OPENBANK_PAYMENT__: openbank付款
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`

	// 支行名称。
	BankBranchName *string `json:"BankBranchName,omitempty" name:"BankBranchName"`

	// 银行简称。
	BankAbbreviation *string `json:"BankAbbreviation,omitempty" name:"BankAbbreviation"`

	// 页码。Index和Count必须大于等于1。
	PageNumber *Paging `json:"PageNumber,omitempty" name:"PageNumber"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type QueryOpenBankBankBranchListRequest struct {
	*tchttp.BaseRequest
	
	// 渠道商户ID。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 渠道名称。
	// __TENPAY__: 商企付
	// __WECHAT__: 微信支付
	// __ALIPAY__: 支付宝
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 支付方式。
	// __EBANK_PAYMENT__:ebank付款
	// __OPENBANK_PAYMENT__: openbank付款
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`

	// 支行名称。
	BankBranchName *string `json:"BankBranchName,omitempty" name:"BankBranchName"`

	// 银行简称。
	BankAbbreviation *string `json:"BankAbbreviation,omitempty" name:"BankAbbreviation"`

	// 页码。Index和Count必须大于等于1。
	PageNumber *Paging `json:"PageNumber,omitempty" name:"PageNumber"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *QueryOpenBankBankBranchListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryOpenBankBankBranchListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelMerchantId")
	delete(f, "ChannelName")
	delete(f, "PaymentMethod")
	delete(f, "BankBranchName")
	delete(f, "BankAbbreviation")
	delete(f, "PageNumber")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryOpenBankBankBranchListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryOpenBankBankBranchListResponseParams struct {
	// 错误码。
	// __SUCCESS__: 成功
	// __其他__: 见附录-错误码表
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *QueryOpenBankBankBranchListResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryOpenBankBankBranchListResponse struct {
	*tchttp.BaseResponse
	Response *QueryOpenBankBankBranchListResponseParams `json:"Response"`
}

func (r *QueryOpenBankBankBranchListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryOpenBankBankBranchListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryOpenBankBankBranchListResult struct {
	// 支行列表。
	BankBranchList []*BankBranchInfo `json:"BankBranchList,omitempty" name:"BankBranchList"`

	// 列表总数。
	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type QueryOpenBankBillData struct {
	// 交易日期
	BillDate *string `json:"BillDate,omitempty" name:"BillDate"`

	// 渠道编码
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// 二级渠道
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubChannel *string `json:"SubChannel,omitempty" name:"SubChannel"`

	// 系统父商户号
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentMerchantId *string `json:"ParentMerchantId,omitempty" name:"ParentMerchantId"`

	// 外部商户号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutMerchantId *string `json:"OutMerchantId,omitempty" name:"OutMerchantId"`

	// 系统商户号
	MerchantId *string `json:"MerchantId,omitempty" name:"MerchantId"`

	// 第三方商户号
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndMerchantId *string `json:"EndMerchantId,omitempty" name:"EndMerchantId"`

	// 外部订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`

	// 系统订单号
	TradeNo *string `json:"TradeNo,omitempty" name:"TradeNo"`

	// 第三方订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTradeNo *string `json:"EndTradeNo,omitempty" name:"EndTradeNo"`

	// 收付类型，PAYMENT:付款，INCOME:收款
	// 注意：此字段可能返回 null，表示取不到有效值。
	PaymentType *string `json:"PaymentType,omitempty" name:"PaymentType"`

	// 业务类型，WITHDRAW:提现，PAY:支付，RECHARGE:充值，TRANSFER:转账，REFUND:退款
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessType *string `json:"BusinessType,omitempty" name:"BusinessType"`

	// 发起交易时间，格式yyyy-MM-dd HH:mm:ss
	TradeTime *string `json:"TradeTime,omitempty" name:"TradeTime"`

	// 交易完成时间，格式yyyy-MM-dd HH:mm:ss
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`

	// 交易状态，0:未知，1:成功，2:失败
	TradeStatus *string `json:"TradeStatus,omitempty" name:"TradeStatus"`

	// 对账状态，1:成功，2:失败 3:长账 4:短账
	CheckStatus *string `json:"CheckStatus,omitempty" name:"CheckStatus"`

	// 对账失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckFailReason *string `json:"CheckFailReason,omitempty" name:"CheckFailReason"`

	// 交易金额（元）
	OrderAmount *string `json:"OrderAmount,omitempty" name:"OrderAmount"`

	// 服务费（元）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceFee *string `json:"ServiceFee,omitempty" name:"ServiceFee"`

	// 收款人账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayeeAccount *string `json:"PayeeAccount,omitempty" name:"PayeeAccount"`

	// 收款人名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayeeName *string `json:"PayeeName,omitempty" name:"PayeeName"`

	// 付款人账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayerAccount *string `json:"PayerAccount,omitempty" name:"PayerAccount"`

	// 付款人名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayerName *string `json:"PayerName,omitempty" name:"PayerName"`

	// 支付信息描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`
}

// Predefined struct for user
type QueryOpenBankBillDataPageRequestParams struct {
	// 渠道商户号，外部接入平台方入驻云企付平台后下发。
	// EBANK_PAYMENT支付方式下，填写渠道商户号；
	// SAFT_ISV支付方式下，填写渠道子商户号;
	// HELIPAY渠道下，填写渠道子商户号。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 账单日期,yyyy-MM-dd。
	BillDate *string `json:"BillDate,omitempty" name:"BillDate"`

	// 渠道名称。详见附录-云企付枚举类说明-ChannelName。
	// __TENPAY__: 商企付
	// __WECHAT__: 微信支付
	// __ALIPAY__: 支付宝
	// HELIPAY：合利宝
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 分页页码。
	PageNo *uint64 `json:"PageNo,omitempty" name:"PageNo"`

	// 分页大小，最大1000。
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 账单类型，默认交易账单。
	BillType *string `json:"BillType,omitempty" name:"BillType"`

	// 支付方式。详见附录-云企付枚举类说明-PaymentMethod。
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type QueryOpenBankBillDataPageRequest struct {
	*tchttp.BaseRequest
	
	// 渠道商户号，外部接入平台方入驻云企付平台后下发。
	// EBANK_PAYMENT支付方式下，填写渠道商户号；
	// SAFT_ISV支付方式下，填写渠道子商户号;
	// HELIPAY渠道下，填写渠道子商户号。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 账单日期,yyyy-MM-dd。
	BillDate *string `json:"BillDate,omitempty" name:"BillDate"`

	// 渠道名称。详见附录-云企付枚举类说明-ChannelName。
	// __TENPAY__: 商企付
	// __WECHAT__: 微信支付
	// __ALIPAY__: 支付宝
	// HELIPAY：合利宝
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 分页页码。
	PageNo *uint64 `json:"PageNo,omitempty" name:"PageNo"`

	// 分页大小，最大1000。
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 账单类型，默认交易账单。
	BillType *string `json:"BillType,omitempty" name:"BillType"`

	// 支付方式。详见附录-云企付枚举类说明-PaymentMethod。
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *QueryOpenBankBillDataPageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryOpenBankBillDataPageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelMerchantId")
	delete(f, "BillDate")
	delete(f, "ChannelName")
	delete(f, "PageNo")
	delete(f, "PageSize")
	delete(f, "BillType")
	delete(f, "PaymentMethod")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryOpenBankBillDataPageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryOpenBankBillDataPageResponseParams struct {
	// 错误码。
	// __SUCCESS__: 成功
	// __其他__: 见附录-错误码表
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *QueryOpenBankBillDataPageResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryOpenBankBillDataPageResponse struct {
	*tchttp.BaseResponse
	Response *QueryOpenBankBillDataPageResponseParams `json:"Response"`
}

func (r *QueryOpenBankBillDataPageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryOpenBankBillDataPageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryOpenBankBillDataPageResult struct {
	// 页码
	PageNo *uint64 `json:"PageNo,omitempty" name:"PageNo"`

	// 分页大小
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 总数
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 账单数据明细
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataList []*QueryOpenBankBillData `json:"DataList,omitempty" name:"DataList"`
}

// Predefined struct for user
type QueryOpenBankBindExternalSubMerchantBankAccountRequestParams struct {
	// 渠道子商户ID。
	ChannelSubMerchantId *string `json:"ChannelSubMerchantId,omitempty" name:"ChannelSubMerchantId"`

	// 渠道商户ID。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 渠道申请编号，与外部申请编号二者选填其一。
	ChannelApplyId *string `json:"ChannelApplyId,omitempty" name:"ChannelApplyId"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// 外部申请编号，与渠道申请编号二者选填其一。
	OutApplyId *string `json:"OutApplyId,omitempty" name:"OutApplyId"`
}

type QueryOpenBankBindExternalSubMerchantBankAccountRequest struct {
	*tchttp.BaseRequest
	
	// 渠道子商户ID。
	ChannelSubMerchantId *string `json:"ChannelSubMerchantId,omitempty" name:"ChannelSubMerchantId"`

	// 渠道商户ID。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 渠道申请编号，与外部申请编号二者选填其一。
	ChannelApplyId *string `json:"ChannelApplyId,omitempty" name:"ChannelApplyId"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// 外部申请编号，与渠道申请编号二者选填其一。
	OutApplyId *string `json:"OutApplyId,omitempty" name:"OutApplyId"`
}

func (r *QueryOpenBankBindExternalSubMerchantBankAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryOpenBankBindExternalSubMerchantBankAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelSubMerchantId")
	delete(f, "ChannelMerchantId")
	delete(f, "ChannelApplyId")
	delete(f, "Environment")
	delete(f, "OutApplyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryOpenBankBindExternalSubMerchantBankAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryOpenBankBindExternalSubMerchantBankAccountResponseParams struct {
	// 错误码。
	// __SUCCESS__: 成功
	// __其他__: 见附录-错误码表
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *QueryOpenBankBindExternalSubMerchantBankAccountResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryOpenBankBindExternalSubMerchantBankAccountResponse struct {
	*tchttp.BaseResponse
	Response *QueryOpenBankBindExternalSubMerchantBankAccountResponseParams `json:"Response"`
}

func (r *QueryOpenBankBindExternalSubMerchantBankAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryOpenBankBindExternalSubMerchantBankAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryOpenBankBindExternalSubMerchantBankAccountResult struct {
	// 渠道子商户收款方银行卡信息, 为JSON格式字符串（绑定成功状态下返回）。详情见附录-复杂类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalSubMerchantBankAccountReturnData *string `json:"ExternalSubMerchantBankAccountReturnData,omitempty" name:"ExternalSubMerchantBankAccountReturnData"`

	// 渠道申请编号。
	ChannelApplyId *string `json:"ChannelApplyId,omitempty" name:"ChannelApplyId"`

	// 绑定状态。
	// __SUCCESS__: 绑定成功
	// __FAILED__: 绑定失败
	// __PROCESSING__: 绑定中
	BindStatus *string `json:"BindStatus,omitempty" name:"BindStatus"`

	// 绑定返回描述, 例如失败原因等。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindMessage *string `json:"BindMessage,omitempty" name:"BindMessage"`

	// 绑卡序列号。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindSerialNo *string `json:"BindSerialNo,omitempty" name:"BindSerialNo"`
}

// Predefined struct for user
type QueryOpenBankDailyReceiptDownloadUrlRequestParams struct {
	// 云企付渠道商户号。外部接入平台入驻云企付平台后下发。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 云企付渠道子商户号。入驻在渠道商户下的子商户ID，如付款方的商户ID，对应创建支付订单中接口参数中的PayerInfo中的payerId。
	ChannelSubMerchantId *string `json:"ChannelSubMerchantId,omitempty" name:"ChannelSubMerchantId"`

	// 渠道名称。
	// __TENPAY__: 商企付
	// __WECHAT__: 微信支付
	// __ALIPAY__: 支付宝
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 付款方式。如
	// __EBANK_PAYMENT__:ebank付款
	// __OPENBANK_PAYMENT__: openbank付款
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`

	// 绑卡序列号，银行卡唯一标记，资金账户ID，用于区分商户绑定多卡或多账户场景
	BindSerialNo *string `json:"BindSerialNo,omitempty" name:"BindSerialNo"`

	// 查询日期，D日查询D-1日的回单文件
	QueryDate *string `json:"QueryDate,omitempty" name:"QueryDate"`

	// 环境类型
	// release:生产环境
	// sandbox:沙箱环境
	// 缺省默认为生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type QueryOpenBankDailyReceiptDownloadUrlRequest struct {
	*tchttp.BaseRequest
	
	// 云企付渠道商户号。外部接入平台入驻云企付平台后下发。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 云企付渠道子商户号。入驻在渠道商户下的子商户ID，如付款方的商户ID，对应创建支付订单中接口参数中的PayerInfo中的payerId。
	ChannelSubMerchantId *string `json:"ChannelSubMerchantId,omitempty" name:"ChannelSubMerchantId"`

	// 渠道名称。
	// __TENPAY__: 商企付
	// __WECHAT__: 微信支付
	// __ALIPAY__: 支付宝
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 付款方式。如
	// __EBANK_PAYMENT__:ebank付款
	// __OPENBANK_PAYMENT__: openbank付款
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`

	// 绑卡序列号，银行卡唯一标记，资金账户ID，用于区分商户绑定多卡或多账户场景
	BindSerialNo *string `json:"BindSerialNo,omitempty" name:"BindSerialNo"`

	// 查询日期，D日查询D-1日的回单文件
	QueryDate *string `json:"QueryDate,omitempty" name:"QueryDate"`

	// 环境类型
	// release:生产环境
	// sandbox:沙箱环境
	// 缺省默认为生产环境
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *QueryOpenBankDailyReceiptDownloadUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryOpenBankDailyReceiptDownloadUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelMerchantId")
	delete(f, "ChannelSubMerchantId")
	delete(f, "ChannelName")
	delete(f, "PaymentMethod")
	delete(f, "BindSerialNo")
	delete(f, "QueryDate")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryOpenBankDailyReceiptDownloadUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryOpenBankDailyReceiptDownloadUrlResponseParams struct {
	// 业务系统返回码，SUCCESS表示成功，其他表示失败。
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 业务系统返回消息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 按日期查询回单下载地址响应对象。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *QueryOpenBankDailyReceiptDownloadUrlResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryOpenBankDailyReceiptDownloadUrlResponse struct {
	*tchttp.BaseResponse
	Response *QueryOpenBankDailyReceiptDownloadUrlResponseParams `json:"Response"`
}

func (r *QueryOpenBankDailyReceiptDownloadUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryOpenBankDailyReceiptDownloadUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryOpenBankDailyReceiptDownloadUrlResult struct {
	// 回单文件下载链接
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 过期时间
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 回单状态
	// PENDING: 处理中
	// READY: 可以下载
	ReceiptStatus *string `json:"ReceiptStatus,omitempty" name:"ReceiptStatus"`
}

// Predefined struct for user
type QueryOpenBankDownLoadUrlRequestParams struct {
	// 渠道商户号，外部接入平台方入驻云企付平台后下发。
	// EBANK_PAYMENT支付方式下，填写渠道商户号；
	// SAFT_ISV支付方式下，填写渠道子商户号。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 账单日期,yyyy-MM-dd。
	BillDate *string `json:"BillDate,omitempty" name:"BillDate"`

	// 账单类型，默认交易账单。
	BillType *string `json:"BillType,omitempty" name:"BillType"`

	// 接入环境。沙箱环境填 sandbox。缺省默认调用生产环境。
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// 渠道名称。不填默认为商企付。
	// __TENPAY__: 商企付
	// __WECHAT__: 微信支付
	// __ALIPAY__: 支付宝
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 支付方式。不填默认为ebank支付。
	// __EBANK_PAYMENT__: ebank支付
	// __OPENBANK_PAYMENT__: openbank支付
	// __SAFT_ISV__: 人资ISV支付
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`
}

type QueryOpenBankDownLoadUrlRequest struct {
	*tchttp.BaseRequest
	
	// 渠道商户号，外部接入平台方入驻云企付平台后下发。
	// EBANK_PAYMENT支付方式下，填写渠道商户号；
	// SAFT_ISV支付方式下，填写渠道子商户号。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 账单日期,yyyy-MM-dd。
	BillDate *string `json:"BillDate,omitempty" name:"BillDate"`

	// 账单类型，默认交易账单。
	BillType *string `json:"BillType,omitempty" name:"BillType"`

	// 接入环境。沙箱环境填 sandbox。缺省默认调用生产环境。
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// 渠道名称。不填默认为商企付。
	// __TENPAY__: 商企付
	// __WECHAT__: 微信支付
	// __ALIPAY__: 支付宝
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 支付方式。不填默认为ebank支付。
	// __EBANK_PAYMENT__: ebank支付
	// __OPENBANK_PAYMENT__: openbank支付
	// __SAFT_ISV__: 人资ISV支付
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`
}

func (r *QueryOpenBankDownLoadUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryOpenBankDownLoadUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelMerchantId")
	delete(f, "BillDate")
	delete(f, "BillType")
	delete(f, "Environment")
	delete(f, "ChannelName")
	delete(f, "PaymentMethod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryOpenBankDownLoadUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryOpenBankDownLoadUrlResponseParams struct {
	// 业务系统返回码，SUCCESS表示成功，其他表示失败。
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 业务系统返回消息。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 查询对账文件下载响应对象。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *QueryOpenBankDownLoadUrlResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryOpenBankDownLoadUrlResponse struct {
	*tchttp.BaseResponse
	Response *QueryOpenBankDownLoadUrlResponseParams `json:"Response"`
}

func (r *QueryOpenBankDownLoadUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryOpenBankDownLoadUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryOpenBankDownLoadUrlResult struct {
	// 供下一步请求账单文件的下载地址。
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 从 download_url 下载的文件的哈希值，用于校验文件的完整性。
	HashValue *string `json:"HashValue,omitempty" name:"HashValue"`

	// 从 download_url 下载的文件的哈希类型，用于校验文件的完整性。
	HashType *string `json:"HashType,omitempty" name:"HashType"`
}

// Predefined struct for user
type QueryOpenBankExternalSubAccountBookBalanceRequestParams struct {
	// 渠道商户ID
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 渠道子商户ID
	ChannelSubMerchantId *string `json:"ChannelSubMerchantId,omitempty" name:"ChannelSubMerchantId"`

	// 渠道名称。目前只支持支付宝
	// __TENPAY__: 商企付
	// __WECHAT__: 微信支付
	// __ALIPAY__: 支付宝
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 支付方式。目前只支持安心发支付
	// __EBANK_PAYMENT__: ebank支付
	// __OPENBANK_PAYMENT__: openbank支付
	// __SAFT_ISV__: 安心发支付
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`

	// 外部账本号ID。与ChannelAccountBookId二者选填其一。
	OutAccountBookId *string `json:"OutAccountBookId,omitempty" name:"OutAccountBookId"`

	// 渠道账本号ID。与OutAccountBookId二者选填其一。
	ChannelAccountBookId *string `json:"ChannelAccountBookId,omitempty" name:"ChannelAccountBookId"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type QueryOpenBankExternalSubAccountBookBalanceRequest struct {
	*tchttp.BaseRequest
	
	// 渠道商户ID
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 渠道子商户ID
	ChannelSubMerchantId *string `json:"ChannelSubMerchantId,omitempty" name:"ChannelSubMerchantId"`

	// 渠道名称。目前只支持支付宝
	// __TENPAY__: 商企付
	// __WECHAT__: 微信支付
	// __ALIPAY__: 支付宝
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 支付方式。目前只支持安心发支付
	// __EBANK_PAYMENT__: ebank支付
	// __OPENBANK_PAYMENT__: openbank支付
	// __SAFT_ISV__: 安心发支付
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`

	// 外部账本号ID。与ChannelAccountBookId二者选填其一。
	OutAccountBookId *string `json:"OutAccountBookId,omitempty" name:"OutAccountBookId"`

	// 渠道账本号ID。与OutAccountBookId二者选填其一。
	ChannelAccountBookId *string `json:"ChannelAccountBookId,omitempty" name:"ChannelAccountBookId"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *QueryOpenBankExternalSubAccountBookBalanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryOpenBankExternalSubAccountBookBalanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelMerchantId")
	delete(f, "ChannelSubMerchantId")
	delete(f, "ChannelName")
	delete(f, "PaymentMethod")
	delete(f, "OutAccountBookId")
	delete(f, "ChannelAccountBookId")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryOpenBankExternalSubAccountBookBalanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryOpenBankExternalSubAccountBookBalanceResponseParams struct {
	// 错误码。
	// __SUCCESS__: 成功
	// __其他__: 见附录-错误码表
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *QueryExternalAccountBookResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryOpenBankExternalSubAccountBookBalanceResponse struct {
	*tchttp.BaseResponse
	Response *QueryOpenBankExternalSubAccountBookBalanceResponseParams `json:"Response"`
}

func (r *QueryOpenBankExternalSubAccountBookBalanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryOpenBankExternalSubAccountBookBalanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryOpenBankExternalSubMerchantBankAccountData struct {
	// 开户银行。
	AccountBank *string `json:"AccountBank,omitempty" name:"AccountBank"`

	// 绑卡序列号。
	BindSerialNo *string `json:"BindSerialNo,omitempty" name:"BindSerialNo"`

	// 账号类型。
	// __COLLECT_MONEY__: 收款卡
	// __PAYMENT__: 付款卡
	AccountType *string `json:"AccountType,omitempty" name:"AccountType"`

	// 支行号。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BankBranchId *string `json:"BankBranchId,omitempty" name:"BankBranchId"`

	// 银行卡卡后四位。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountNumberLastFour *string `json:"AccountNumberLastFour,omitempty" name:"AccountNumberLastFour"`
}

// Predefined struct for user
type QueryOpenBankExternalSubMerchantBankAccountRequestParams struct {
	// 渠道商户ID。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 渠道子商户ID。
	ChannelSubMerchantId *string `json:"ChannelSubMerchantId,omitempty" name:"ChannelSubMerchantId"`

	// 渠道名称。
	// __TENPAY__: 商企付
	// __WECHAT__: 微信支付
	// __ALIPAY__: 支付宝
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 支付方式。
	// __EBANK_PAYMENT__: ebank支付
	// __OPENBANK_PAYMENT__: openbank支付
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type QueryOpenBankExternalSubMerchantBankAccountRequest struct {
	*tchttp.BaseRequest
	
	// 渠道商户ID。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 渠道子商户ID。
	ChannelSubMerchantId *string `json:"ChannelSubMerchantId,omitempty" name:"ChannelSubMerchantId"`

	// 渠道名称。
	// __TENPAY__: 商企付
	// __WECHAT__: 微信支付
	// __ALIPAY__: 支付宝
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 支付方式。
	// __EBANK_PAYMENT__: ebank支付
	// __OPENBANK_PAYMENT__: openbank支付
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *QueryOpenBankExternalSubMerchantBankAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryOpenBankExternalSubMerchantBankAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelMerchantId")
	delete(f, "ChannelSubMerchantId")
	delete(f, "ChannelName")
	delete(f, "PaymentMethod")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryOpenBankExternalSubMerchantBankAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryOpenBankExternalSubMerchantBankAccountResponseParams struct {
	// 错误码。
	// __SUCCESS__: 成功
	// __其他__: 见附录-错误码表
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *QueryOpenBankExternalSubMerchantBankAccountResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryOpenBankExternalSubMerchantBankAccountResponse struct {
	*tchttp.BaseResponse
	Response *QueryOpenBankExternalSubMerchantBankAccountResponseParams `json:"Response"`
}

func (r *QueryOpenBankExternalSubMerchantBankAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryOpenBankExternalSubMerchantBankAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryOpenBankExternalSubMerchantBankAccountResult struct {
	// 第三方渠道子商户查询银行账户返回。
	AccountList []*QueryOpenBankExternalSubMerchantBankAccountData `json:"AccountList,omitempty" name:"AccountList"`
}

// Predefined struct for user
type QueryOpenBankExternalSubMerchantRegistrationRequestParams struct {
	// 渠道商户号。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 渠道进件号，与外部进件号二者选填其一。
	ChannelRegistrationNo *string `json:"ChannelRegistrationNo,omitempty" name:"ChannelRegistrationNo"`

	// 外部进件号，与渠道进件号二者选填其一。
	OutRegistrationNo *string `json:"OutRegistrationNo,omitempty" name:"OutRegistrationNo"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type QueryOpenBankExternalSubMerchantRegistrationRequest struct {
	*tchttp.BaseRequest
	
	// 渠道商户号。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 渠道进件号，与外部进件号二者选填其一。
	ChannelRegistrationNo *string `json:"ChannelRegistrationNo,omitempty" name:"ChannelRegistrationNo"`

	// 外部进件号，与渠道进件号二者选填其一。
	OutRegistrationNo *string `json:"OutRegistrationNo,omitempty" name:"OutRegistrationNo"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *QueryOpenBankExternalSubMerchantRegistrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryOpenBankExternalSubMerchantRegistrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelMerchantId")
	delete(f, "ChannelRegistrationNo")
	delete(f, "OutRegistrationNo")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryOpenBankExternalSubMerchantRegistrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryOpenBankExternalSubMerchantRegistrationResponseParams struct {
	// 错误码。
	// __SUCCESS__: 成功
	// __其他__: 见附录-错误码表
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *QueryOpenBankExternalSubMerchantRegistrationResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryOpenBankExternalSubMerchantRegistrationResponse struct {
	*tchttp.BaseResponse
	Response *QueryOpenBankExternalSubMerchantRegistrationResponseParams `json:"Response"`
}

func (r *QueryOpenBankExternalSubMerchantRegistrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryOpenBankExternalSubMerchantRegistrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryOpenBankExternalSubMerchantRegistrationResult struct {
	// 进件状态。
	// __SUCCESS__: 进件成功
	// __FAILED__: 进件失败
	// __PROCESSING__: 进件中
	RegistrationStatus *string `json:"RegistrationStatus,omitempty" name:"RegistrationStatus"`

	// 进件返回描述, 例如失败原因等。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegistrationMessage *string `json:"RegistrationMessage,omitempty" name:"RegistrationMessage"`

	// 渠道进件号。
	ChannelRegistrationNo *string `json:"ChannelRegistrationNo,omitempty" name:"ChannelRegistrationNo"`

	// 渠道子商户ID（进件成功返回）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelSubMerchantId *string `json:"ChannelSubMerchantId,omitempty" name:"ChannelSubMerchantId"`

	// 外部子商户名称（进件成功返回）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutSubMerchantName *string `json:"OutSubMerchantName,omitempty" name:"OutSubMerchantName"`

	// 渠道名称（进件成功返回）。
	// __TENPAY__: 商企付
	// __WECHAT__: 微信支付
	// __ALIPAY__: 支付宝
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 支付方式（进件成功返回）。
	// __EBANK_PAYMENT__: ebank支付
	// __OPENBANK_PAYMENT__: openbank支付
	// 注意：此字段可能返回 null，表示取不到有效值。
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`

	// 社会信用代码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessLicenseNumber *string `json:"BusinessLicenseNumber,omitempty" name:"BusinessLicenseNumber"`

	// 法人姓名
	// 注意：此字段可能返回 null，表示取不到有效值。
	LegalName *string `json:"LegalName,omitempty" name:"LegalName"`

	// 第三方渠道查询进件返回数据。详情见附录-复杂类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalReturnData *string `json:"ExternalReturnData,omitempty" name:"ExternalReturnData"`
}

// Predefined struct for user
type QueryOpenBankOrderDetailReceiptInfoRequestParams struct {
	// 渠道商户ID
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 渠道子商户ID
	ChannelSubMerchantId *string `json:"ChannelSubMerchantId,omitempty" name:"ChannelSubMerchantId"`

	// 渠道名称，目前只支持ALIPAY
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 支付方式，目前只支持SAFT_ISV
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`

	// 外部回单申请ID，与渠道回单申请ID二者选填其一
	OutApplyId *string `json:"OutApplyId,omitempty" name:"OutApplyId"`

	// 渠道回单申请ID，与外部回单申请ID二者选填其一
	ChannelApplyId *string `json:"ChannelApplyId,omitempty" name:"ChannelApplyId"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type QueryOpenBankOrderDetailReceiptInfoRequest struct {
	*tchttp.BaseRequest
	
	// 渠道商户ID
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 渠道子商户ID
	ChannelSubMerchantId *string `json:"ChannelSubMerchantId,omitempty" name:"ChannelSubMerchantId"`

	// 渠道名称，目前只支持ALIPAY
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 支付方式，目前只支持SAFT_ISV
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`

	// 外部回单申请ID，与渠道回单申请ID二者选填其一
	OutApplyId *string `json:"OutApplyId,omitempty" name:"OutApplyId"`

	// 渠道回单申请ID，与外部回单申请ID二者选填其一
	ChannelApplyId *string `json:"ChannelApplyId,omitempty" name:"ChannelApplyId"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *QueryOpenBankOrderDetailReceiptInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryOpenBankOrderDetailReceiptInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelMerchantId")
	delete(f, "ChannelSubMerchantId")
	delete(f, "ChannelName")
	delete(f, "PaymentMethod")
	delete(f, "OutApplyId")
	delete(f, "ChannelApplyId")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryOpenBankOrderDetailReceiptInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryOpenBankOrderDetailReceiptInfoResponseParams struct {
	// 错误码。
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *QueryOpenBankOrderDetailReceiptInfoResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryOpenBankOrderDetailReceiptInfoResponse struct {
	*tchttp.BaseResponse
	Response *QueryOpenBankOrderDetailReceiptInfoResponseParams `json:"Response"`
}

func (r *QueryOpenBankOrderDetailReceiptInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryOpenBankOrderDetailReceiptInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryOpenBankOrderDetailReceiptInfoResult struct {
	// 渠道回单申请ID
	ChannelApplyId *string `json:"ChannelApplyId,omitempty" name:"ChannelApplyId"`

	// 申请状态。
	// SUCCESS：申请成功；
	// FAILED：申请失败；
	// PROCESSING：申请中。
	// 注意：若返回申请中，需要再次调用回单申请结果查询接口，查询结果。
	ReceiptStatus *string `json:"ReceiptStatus,omitempty" name:"ReceiptStatus"`

	// 申请返回描述，例如失败原因等。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReceiptMessage *string `json:"ReceiptMessage,omitempty" name:"ReceiptMessage"`

	// 回单下载链接，申请成功时返回。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 过期时间，yyyy-MM-dd HH:mm:ss格式。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

// Predefined struct for user
type QueryOpenBankPaymentOrderRequestParams struct {
	// 渠道商户号。外部接入平台入驻云企付平台下发。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 外部商户订单号。与ChannelOrderId不能同时为空。
	OutOrderId *string `json:"OutOrderId,omitempty" name:"OutOrderId"`

	// 云平台订单号。与OutOrderId不能同时为空。
	ChannelOrderId *string `json:"ChannelOrderId,omitempty" name:"ChannelOrderId"`

	// 接入环境。沙箱环境填 sandbox。缺省默认调用生产环境。
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type QueryOpenBankPaymentOrderRequest struct {
	*tchttp.BaseRequest
	
	// 渠道商户号。外部接入平台入驻云企付平台下发。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 外部商户订单号。与ChannelOrderId不能同时为空。
	OutOrderId *string `json:"OutOrderId,omitempty" name:"OutOrderId"`

	// 云平台订单号。与OutOrderId不能同时为空。
	ChannelOrderId *string `json:"ChannelOrderId,omitempty" name:"ChannelOrderId"`

	// 接入环境。沙箱环境填 sandbox。缺省默认调用生产环境。
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *QueryOpenBankPaymentOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryOpenBankPaymentOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelMerchantId")
	delete(f, "OutOrderId")
	delete(f, "ChannelOrderId")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryOpenBankPaymentOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryOpenBankPaymentOrderResponseParams struct {
	// 业务系统返回码，SUCCESS表示成功，其他表示失败。
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 业务系统返回消息。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 查询支付结果响应对象。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *QueryOpenBankPaymentOrderResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryOpenBankPaymentOrderResponse struct {
	*tchttp.BaseResponse
	Response *QueryOpenBankPaymentOrderResponseParams `json:"Response"`
}

func (r *QueryOpenBankPaymentOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryOpenBankPaymentOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryOpenBankPaymentOrderResult struct {
	// 渠道商户号。外部接入平台入驻云企付平台下发
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 外部商户订单号
	OutOrderId *string `json:"OutOrderId,omitempty" name:"OutOrderId"`

	// 云企付平台订单号
	ChannelOrderId *string `json:"ChannelOrderId,omitempty" name:"ChannelOrderId"`

	// 第三方支付平台订单号
	ThirdPayOrderId *string `json:"ThirdPayOrderId,omitempty" name:"ThirdPayOrderId"`

	// 订单状态。
	// INIT：初始化
	// PAYING：支付中
	// ACCEPTED：支付受理成功
	// SUCCESS：支付成功
	// CLOSED：关单
	// PAY_FAIL：支付失败
	// REVOKE：退票
	OrderStatus *string `json:"OrderStatus,omitempty" name:"OrderStatus"`

	// 支付渠道名称，如TENPAY
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 付款方式。如EBANK_PAYMENT
	// OPENBANK_PAYMENT
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`

	// 订单金额。单位分
	TotalAmount *int64 `json:"TotalAmount,omitempty" name:"TotalAmount"`

	// 实际支付金额。单位分，支付成功时返回
	PayAmount *int64 `json:"PayAmount,omitempty" name:"PayAmount"`

	// 失败原因，若失败的返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailReason *string `json:"FailReason,omitempty" name:"FailReason"`

	// 附加信息，查询时原样透传
	// 注意：此字段可能返回 null，表示取不到有效值。
	Attachment *string `json:"Attachment,omitempty" name:"Attachment"`

	// 重定向参数，用于客户端跳转，订单未支付时返回该参数
	// 渠道为TENPAY，付款方式为EBANK_PAYMENT时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	RedirectInfo *OpenBankRedirectInfo `json:"RedirectInfo,omitempty" name:"RedirectInfo"`

	// 第三方渠道返回信息，见渠道特殊说明,详情见附录-复杂类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalReturnData *string `json:"ExternalReturnData,omitempty" name:"ExternalReturnData"`

	// 银行复核指引。当TENPAY下OPENBANT_PAYMENT时，下单受理成功是返回。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BankApprovalGuideInfo *OpenBankApprovalGuideInfo `json:"BankApprovalGuideInfo,omitempty" name:"BankApprovalGuideInfo"`

	// 手续费金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	FeeAmount *int64 `json:"FeeAmount,omitempty" name:"FeeAmount"`

	// 手续费费率
	// 注意：此字段可能返回 null，表示取不到有效值。
	FeeRate *int64 `json:"FeeRate,omitempty" name:"FeeRate"`
}

// Predefined struct for user
type QueryOpenBankRefundOrderRequestParams struct {
	// 渠道商户号。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 外部商户退单号，与渠道退款单号二者选填其一。
	OutRefundId *string `json:"OutRefundId,omitempty" name:"OutRefundId"`

	// 渠道退款订单号，与外部商户退款单号二者选填其一。
	ChannelRefundId *string `json:"ChannelRefundId,omitempty" name:"ChannelRefundId"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type QueryOpenBankRefundOrderRequest struct {
	*tchttp.BaseRequest
	
	// 渠道商户号。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 外部商户退单号，与渠道退款单号二者选填其一。
	OutRefundId *string `json:"OutRefundId,omitempty" name:"OutRefundId"`

	// 渠道退款订单号，与外部商户退款单号二者选填其一。
	ChannelRefundId *string `json:"ChannelRefundId,omitempty" name:"ChannelRefundId"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *QueryOpenBankRefundOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryOpenBankRefundOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelMerchantId")
	delete(f, "OutRefundId")
	delete(f, "ChannelRefundId")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryOpenBankRefundOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryOpenBankRefundOrderResponseParams struct {
	// 错误码
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *OpenBankQueryRefundOrderResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryOpenBankRefundOrderResponse struct {
	*tchttp.BaseResponse
	Response *QueryOpenBankRefundOrderResponseParams `json:"Response"`
}

func (r *QueryOpenBankRefundOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryOpenBankRefundOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryOpenBankSubMerchantCredentialRequestParams struct {
	// 渠道商户ID。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 渠道子商户ID。
	ChannelSubMerchantId *string `json:"ChannelSubMerchantId,omitempty" name:"ChannelSubMerchantId"`

	// 渠道名称。详见附录-枚举类型-ChannelName。
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 支付方式。
	// 合利宝渠道不需要传。
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`

	// 外部申请流水号。
	// 外部申请流水号与渠道申请流水号两者选填其一。
	OutApplyId *string `json:"OutApplyId,omitempty" name:"OutApplyId"`

	// 渠道申请流水号。
	// 外部申请流水号与渠道申请流水号两者选填其一。
	ChannelApplyId *string `json:"ChannelApplyId,omitempty" name:"ChannelApplyId"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type QueryOpenBankSubMerchantCredentialRequest struct {
	*tchttp.BaseRequest
	
	// 渠道商户ID。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 渠道子商户ID。
	ChannelSubMerchantId *string `json:"ChannelSubMerchantId,omitempty" name:"ChannelSubMerchantId"`

	// 渠道名称。详见附录-枚举类型-ChannelName。
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 支付方式。
	// 合利宝渠道不需要传。
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`

	// 外部申请流水号。
	// 外部申请流水号与渠道申请流水号两者选填其一。
	OutApplyId *string `json:"OutApplyId,omitempty" name:"OutApplyId"`

	// 渠道申请流水号。
	// 外部申请流水号与渠道申请流水号两者选填其一。
	ChannelApplyId *string `json:"ChannelApplyId,omitempty" name:"ChannelApplyId"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *QueryOpenBankSubMerchantCredentialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryOpenBankSubMerchantCredentialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelMerchantId")
	delete(f, "ChannelSubMerchantId")
	delete(f, "ChannelName")
	delete(f, "PaymentMethod")
	delete(f, "OutApplyId")
	delete(f, "ChannelApplyId")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryOpenBankSubMerchantCredentialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryOpenBankSubMerchantCredentialResponseParams struct {
	// 错误码。
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *QueryOpenBankSubMerchantCredentialResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryOpenBankSubMerchantCredentialResponse struct {
	*tchttp.BaseResponse
	Response *QueryOpenBankSubMerchantCredentialResponseParams `json:"Response"`
}

func (r *QueryOpenBankSubMerchantCredentialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryOpenBankSubMerchantCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryOpenBankSubMerchantCredentialResult struct {
	// 上传状态
	UploadStatus *string `json:"UploadStatus,omitempty" name:"UploadStatus"`

	// 上传描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	UploadMessage *string `json:"UploadMessage,omitempty" name:"UploadMessage"`
}

// Predefined struct for user
type QueryOpenBankSubMerchantRateConfigureRequestParams struct {
	// 渠道进件序列号。
	ChannelRegistrationNo *string `json:"ChannelRegistrationNo,omitempty" name:"ChannelRegistrationNo"`

	// 渠道商户ID。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 渠道子商户ID。
	ChannelSubMerchantId *string `json:"ChannelSubMerchantId,omitempty" name:"ChannelSubMerchantId"`

	// 渠道名称。详见附录-云企付枚举类说明-ChannelName。
	// TENPAY: 商企付
	// WECHAT: 微信支付
	// ALIPAY: 支付宝
	// HELIPAY:合利宝
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 渠道产品费率序列号。与外部产品费率序列号二者选填其一。
	ChannelProductFeeNo *string `json:"ChannelProductFeeNo,omitempty" name:"ChannelProductFeeNo"`

	// 外部产品费率序列号。与渠道产品费率序列号二者选填其一。
	OutProductFeeNo *string `json:"OutProductFeeNo,omitempty" name:"OutProductFeeNo"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type QueryOpenBankSubMerchantRateConfigureRequest struct {
	*tchttp.BaseRequest
	
	// 渠道进件序列号。
	ChannelRegistrationNo *string `json:"ChannelRegistrationNo,omitempty" name:"ChannelRegistrationNo"`

	// 渠道商户ID。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 渠道子商户ID。
	ChannelSubMerchantId *string `json:"ChannelSubMerchantId,omitempty" name:"ChannelSubMerchantId"`

	// 渠道名称。详见附录-云企付枚举类说明-ChannelName。
	// TENPAY: 商企付
	// WECHAT: 微信支付
	// ALIPAY: 支付宝
	// HELIPAY:合利宝
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 渠道产品费率序列号。与外部产品费率序列号二者选填其一。
	ChannelProductFeeNo *string `json:"ChannelProductFeeNo,omitempty" name:"ChannelProductFeeNo"`

	// 外部产品费率序列号。与渠道产品费率序列号二者选填其一。
	OutProductFeeNo *string `json:"OutProductFeeNo,omitempty" name:"OutProductFeeNo"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *QueryOpenBankSubMerchantRateConfigureRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryOpenBankSubMerchantRateConfigureRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelRegistrationNo")
	delete(f, "ChannelMerchantId")
	delete(f, "ChannelSubMerchantId")
	delete(f, "ChannelName")
	delete(f, "ChannelProductFeeNo")
	delete(f, "OutProductFeeNo")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryOpenBankSubMerchantRateConfigureRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryOpenBankSubMerchantRateConfigureResponseParams struct {
	// 错误码。
	// __SUCCESS__: 成功
	// __其他__: 见附录-错误码表
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *QueryOpenBankSubMerchantRateConfigureResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryOpenBankSubMerchantRateConfigureResponse struct {
	*tchttp.BaseResponse
	Response *QueryOpenBankSubMerchantRateConfigureResponseParams `json:"Response"`
}

func (r *QueryOpenBankSubMerchantRateConfigureResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryOpenBankSubMerchantRateConfigureResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryOpenBankSubMerchantRateConfigureResult struct {
	// 处理状态 
	// SUCCESS: 开通成功 
	// FAILED: 开通失败
	// PROCESSING: 开通中
	DealStatus *string `json:"DealStatus,omitempty" name:"DealStatus"`

	// 处理描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	DealMessage *string `json:"DealMessage,omitempty" name:"DealMessage"`
}

// Predefined struct for user
type QueryOpenBankSupportBankListRequestParams struct {
	// 渠道商户ID。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 渠道名称。
	// __TENPAY__: 商企付
	// __WECHAT__: 微信支付
	// __ALIPAY__: 支付宝
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 支付方式。
	// __EBANK_PAYMENT__:ebank付款
	// __OPENBANK_PAYMENT__: openbank付款
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type QueryOpenBankSupportBankListRequest struct {
	*tchttp.BaseRequest
	
	// 渠道商户ID。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 渠道名称。
	// __TENPAY__: 商企付
	// __WECHAT__: 微信支付
	// __ALIPAY__: 支付宝
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 支付方式。
	// __EBANK_PAYMENT__:ebank付款
	// __OPENBANK_PAYMENT__: openbank付款
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *QueryOpenBankSupportBankListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryOpenBankSupportBankListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelMerchantId")
	delete(f, "ChannelName")
	delete(f, "PaymentMethod")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryOpenBankSupportBankListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryOpenBankSupportBankListResponseParams struct {
	// 错误码。
	// __SUCCESS__: 成功
	// __其他__: 见附录-错误码表
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *QueryOpenBankSupportBankListResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryOpenBankSupportBankListResponse struct {
	*tchttp.BaseResponse
	Response *QueryOpenBankSupportBankListResponseParams `json:"Response"`
}

func (r *QueryOpenBankSupportBankListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryOpenBankSupportBankListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryOpenBankSupportBankListResult struct {
	// 支持的银行列表
	SupportBankList []*SupportBankInfo `json:"SupportBankList,omitempty" name:"SupportBankList"`
}

// Predefined struct for user
type QueryOpenBankUnbindExternalSubMerchantBankAccountRequestParams struct {
	// 渠道子商户ID。
	ChannelSubMerchantId *string `json:"ChannelSubMerchantId,omitempty" name:"ChannelSubMerchantId"`

	// 渠道商户ID。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 渠道申请编号，与外部申请编号二者选填其一。
	ChannelApplyId *string `json:"ChannelApplyId,omitempty" name:"ChannelApplyId"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// 外部申请编号，与渠道申请编号二者选填其一。
	OutApplyId *string `json:"OutApplyId,omitempty" name:"OutApplyId"`
}

type QueryOpenBankUnbindExternalSubMerchantBankAccountRequest struct {
	*tchttp.BaseRequest
	
	// 渠道子商户ID。
	ChannelSubMerchantId *string `json:"ChannelSubMerchantId,omitempty" name:"ChannelSubMerchantId"`

	// 渠道商户ID。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 渠道申请编号，与外部申请编号二者选填其一。
	ChannelApplyId *string `json:"ChannelApplyId,omitempty" name:"ChannelApplyId"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`

	// 外部申请编号，与渠道申请编号二者选填其一。
	OutApplyId *string `json:"OutApplyId,omitempty" name:"OutApplyId"`
}

func (r *QueryOpenBankUnbindExternalSubMerchantBankAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryOpenBankUnbindExternalSubMerchantBankAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelSubMerchantId")
	delete(f, "ChannelMerchantId")
	delete(f, "ChannelApplyId")
	delete(f, "Environment")
	delete(f, "OutApplyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryOpenBankUnbindExternalSubMerchantBankAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryOpenBankUnbindExternalSubMerchantBankAccountResponseParams struct {
	// 错误码。
	// __SUCCESS__: 成功
	// __其他__: 见附录-错误码表
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *QueryOpenBankUnbindExternalSubMerchantBankAccountResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryOpenBankUnbindExternalSubMerchantBankAccountResponse struct {
	*tchttp.BaseResponse
	Response *QueryOpenBankUnbindExternalSubMerchantBankAccountResponseParams `json:"Response"`
}

func (r *QueryOpenBankUnbindExternalSubMerchantBankAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryOpenBankUnbindExternalSubMerchantBankAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryOpenBankUnbindExternalSubMerchantBankAccountResult struct {
	// 渠道申请编号。
	ChannelApplyId *string `json:"ChannelApplyId,omitempty" name:"ChannelApplyId"`

	// 解绑状态。
	// __SUCCESS__: 解绑成功
	// __FAILED__: 解绑失败
	// __PROCESSING__: 解绑中
	UnbindStatus *string `json:"UnbindStatus,omitempty" name:"UnbindStatus"`

	// 解绑返回描述, 例如失败原因等。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnbindMessage *string `json:"UnbindMessage,omitempty" name:"UnbindMessage"`
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

// Predefined struct for user
type QueryOrderRequestParams struct {
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

// Predefined struct for user
type QueryOrderResponseParams struct {
	// 返回订单数
	TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`

	// 查询结果的订单列表
	OrderList []*QueryOrderOutOrderList `json:"OrderList,omitempty" name:"OrderList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryOrderResponse struct {
	*tchttp.BaseResponse
	Response *QueryOrderResponseParams `json:"Response"`
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

// Predefined struct for user
type QueryOrderStatusRequestParams struct {
	// 使用门店OpenId
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 使用门店OpenKey
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 开发者流水号
	DeveloperNo *string `json:"DeveloperNo,omitempty" name:"DeveloperNo"`

	// 付款订单号
	OrderNo *string `json:"OrderNo,omitempty" name:"OrderNo"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

type QueryOrderStatusRequest struct {
	*tchttp.BaseRequest
	
	// 使用门店OpenId
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 使用门店OpenKey
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 开发者流水号
	DeveloperNo *string `json:"DeveloperNo,omitempty" name:"DeveloperNo"`

	// 付款订单号
	OrderNo *string `json:"OrderNo,omitempty" name:"OrderNo"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *QueryOrderStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryOrderStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OpenId")
	delete(f, "OpenKey")
	delete(f, "DeveloperNo")
	delete(f, "OrderNo")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryOrderStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryOrderStatusResponseParams struct {
	// 业务系统返回码，0表示成功，其他表示失败。
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 业务系统返回消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 查询订单付款状态结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *QueryOrderStatusResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryOrderStatusResponse struct {
	*tchttp.BaseResponse
	Response *QueryOrderStatusResponseParams `json:"Response"`
}

func (r *QueryOrderStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryOrderStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryOrderStatusResult struct {
	// 付款订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderNo *string `json:"OrderNo,omitempty" name:"OrderNo"`

	// 开发者流水号
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeveloperNo *string `json:"DeveloperNo,omitempty" name:"DeveloperNo"`

	// 交易优惠金额（免充值券）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TradeDiscountAmount *string `json:"TradeDiscountAmount,omitempty" name:"TradeDiscountAmount"`

	// 付款方式名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayName *string `json:"PayName,omitempty" name:"PayName"`

	// 商户流水号（从1开始自增长不重复）
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderMerchantId *string `json:"OrderMerchantId,omitempty" name:"OrderMerchantId"`

	// 交易帐号（银行卡号、支付宝帐号、微信帐号等，某些收单机构没有此数据）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TradeAccount *string `json:"TradeAccount,omitempty" name:"TradeAccount"`

	// 实际交易金额（以分为单位，没有小数点）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TradeAmount *string `json:"TradeAmount,omitempty" name:"TradeAmount"`

	// 币种签名
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrencySign *string `json:"CurrencySign,omitempty" name:"CurrencySign"`

	// 付款完成时间（以收单机构为准）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TradePayTime *string `json:"TradePayTime,omitempty" name:"TradePayTime"`

	// 门店流水号（从1开始自增长不重复）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShopOrderId *string `json:"ShopOrderId,omitempty" name:"ShopOrderId"`

	// 支付标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayTag *string `json:"PayTag,omitempty" name:"PayTag"`

	// 订单状态（1交易成功，2待支付，4已取消，9等待用户输入密码确认
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 币种代码
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderCurrency *string `json:"OrderCurrency,omitempty" name:"OrderCurrency"`

	// 二维码字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	TradeQrcode *string `json:"TradeQrcode,omitempty" name:"TradeQrcode"`

	// 开始交易时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TradeTime *string `json:"TradeTime,omitempty" name:"TradeTime"`

	// 折扣金额（以分为单位，没有小数点）
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiscountAmount *string `json:"DiscountAmount,omitempty" name:"DiscountAmount"`

	// 商户号
	// 注意：此字段可能返回 null，表示取不到有效值。
	MerchantNo *string `json:"MerchantNo,omitempty" name:"MerchantNo"`

	// 订单备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 订单标题
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderName *string `json:"OrderName,omitempty" name:"OrderName"`

	// 原始金额（以分为单位，没有小数点）
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginalAmount *string `json:"OriginalAmount,omitempty" name:"OriginalAmount"`

	// 门店编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShopNo *string `json:"ShopNo,omitempty" name:"ShopNo"`

	// 收单机构原始交易数据，如果返回非标准json数据，请自行转换
	// 注意：此字段可能返回 null，表示取不到有效值。
	TradeResult *string `json:"TradeResult,omitempty" name:"TradeResult"`

	// 订单流水号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 订单类型（1消费，2辙单）
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderType *string `json:"OrderType,omitempty" name:"OrderType"`

	// 收单机构交易号
	// 注意：此字段可能返回 null，表示取不到有效值。
	TradeNo *string `json:"TradeNo,omitempty" name:"TradeNo"`

	// 原始订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginalOrderNo *string `json:"OriginalOrderNo,omitempty" name:"OriginalOrderNo"`

	// 订单标记，订单附加数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag *string `json:"Tag,omitempty" name:"Tag"`

	// 下单时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

	// 收银员编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	CashierId *string `json:"CashierId,omitempty" name:"CashierId"`

	// 收银员名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	CashierRealName *string `json:"CashierRealName,omitempty" name:"CashierRealName"`

	// 店铺全称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShopFullName *string `json:"ShopFullName,omitempty" name:"ShopFullName"`

	// 店铺名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShopName *string `json:"ShopName,omitempty" name:"ShopName"`
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

// Predefined struct for user
type QueryOutwardOrderRequestParams struct {
	// 对接方汇出指令编号
	TransactionId *string `json:"TransactionId,omitempty" name:"TransactionId"`

	// 接入环境。沙箱环境填sandbox
	Profile *string `json:"Profile,omitempty" name:"Profile"`
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

// Predefined struct for user
type QueryOutwardOrderResponseParams struct {
	// 查询汇出结果
	Result *QueryOutwardOrderResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryOutwardOrderResponse struct {
	*tchttp.BaseResponse
	Response *QueryOutwardOrderResponseParams `json:"Response"`
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

// Predefined struct for user
type QueryPayerInfoRequestParams struct {
	// 付款人ID
	PayerId *string `json:"PayerId,omitempty" name:"PayerId"`

	// 接入环境。沙箱环境填sandbox
	Profile *string `json:"Profile,omitempty" name:"Profile"`
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

// Predefined struct for user
type QueryPayerInfoResponseParams struct {
	// 付款人查询结果
	Result *QueryPayerinfoResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryPayerInfoResponse struct {
	*tchttp.BaseResponse
	Response *QueryPayerInfoResponseParams `json:"Response"`
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

// Predefined struct for user
type QueryReconciliationDocumentRequestParams struct {
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

// Predefined struct for user
type QueryReconciliationDocumentResponseParams struct {
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
}

type QueryReconciliationDocumentResponse struct {
	*tchttp.BaseResponse
	Response *QueryReconciliationDocumentResponseParams `json:"Response"`
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

// Predefined struct for user
type QueryReconciliationFileApplyInfoRequestParams struct {
	// 申请对账文件的任务ID。
	ApplyFileId *string `json:"ApplyFileId,omitempty" name:"ApplyFileId"`

	// 环境名。
	// __release__: 现网环境
	// __sandbox__: 沙箱环境
	// __development__: 开发环境
	// _缺省: release_
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`
}

type QueryReconciliationFileApplyInfoRequest struct {
	*tchttp.BaseRequest
	
	// 申请对账文件的任务ID。
	ApplyFileId *string `json:"ApplyFileId,omitempty" name:"ApplyFileId"`

	// 环境名。
	// __release__: 现网环境
	// __sandbox__: 沙箱环境
	// __development__: 开发环境
	// _缺省: release_
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`
}

func (r *QueryReconciliationFileApplyInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryReconciliationFileApplyInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplyFileId")
	delete(f, "MidasEnvironment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryReconciliationFileApplyInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryReconciliationFileApplyInfoResponseParams struct {
	// 错误码。
	// __SUCCESS__: 成功
	// __其他__: 见附录-错误码表
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *QueryReconciliationFileApplyInfoResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryReconciliationFileApplyInfoResponse struct {
	*tchttp.BaseResponse
	Response *QueryReconciliationFileApplyInfoResponseParams `json:"Response"`
}

func (r *QueryReconciliationFileApplyInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryReconciliationFileApplyInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryReconciliationFileApplyInfoResult struct {
	// 申请对账文件的任务ID。
	ApplyFileId *string `json:"ApplyFileId,omitempty" name:"ApplyFileId"`

	// 对账文件申请状态。
	// __I__：申请中
	// __S__：申请成功
	// __F__：申请失败
	ApplyStatus *string `json:"ApplyStatus,omitempty" name:"ApplyStatus"`

	// 申请结果描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplyMessage *string `json:"ApplyMessage,omitempty" name:"ApplyMessage"`

	// 对账文件下载地址列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileUrlArray []*string `json:"FileUrlArray,omitempty" name:"FileUrlArray"`
}

// Predefined struct for user
type QueryRefundRequestParams struct {
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

// Predefined struct for user
type QueryRefundResponseParams struct {
	// 退款状态码，退款提交成功后返回  1：退款中；  2：退款成功；  3：退款失败。
	State *string `json:"State,omitempty" name:"State"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryRefundResponse struct {
	*tchttp.BaseResponse
	Response *QueryRefundResponseParams `json:"Response"`
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

// Predefined struct for user
type QueryShopOpenIdRequestParams struct {
	// 收单系统分配的开放ID
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 收单系统分配的密钥
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 门店编号
	ShopNo *string `json:"ShopNo,omitempty" name:"ShopNo"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

type QueryShopOpenIdRequest struct {
	*tchttp.BaseRequest
	
	// 收单系统分配的开放ID
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 收单系统分配的密钥
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 门店编号
	ShopNo *string `json:"ShopNo,omitempty" name:"ShopNo"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *QueryShopOpenIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryShopOpenIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OpenId")
	delete(f, "OpenKey")
	delete(f, "ShopNo")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryShopOpenIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryShopOpenIdResponseParams struct {
	// 业务系统返回消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 业务系统返回码
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 获取门店OpenId响应对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *QueryShopOpenIdResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryShopOpenIdResponse struct {
	*tchttp.BaseResponse
	Response *QueryShopOpenIdResponseParams `json:"Response"`
}

func (r *QueryShopOpenIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryShopOpenIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryShopOpenIdResult struct {
	// 省份
	// 注意：此字段可能返回 null，表示取不到有效值。
	Province *string `json:"Province,omitempty" name:"Province"`

	// 城市
	// 注意：此字段可能返回 null，表示取不到有效值。
	City *string `json:"City,omitempty" name:"City"`

	// 门店简称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShopName *string `json:"ShopName,omitempty" name:"ShopName"`

	// 商户编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	MerchantNo *string `json:"MerchantNo,omitempty" name:"MerchantNo"`

	// 城市编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	CityId *string `json:"CityId,omitempty" name:"CityId"`

	// open_id
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 门店电话
	// 注意：此字段可能返回 null，表示取不到有效值。
	Telephone *string `json:"Telephone,omitempty" name:"Telephone"`

	// 门店编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShopNo *string `json:"ShopNo,omitempty" name:"ShopNo"`

	// 县/区
	// 注意：此字段可能返回 null，表示取不到有效值。
	County *string `json:"County,omitempty" name:"County"`

	// 门店全称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShopFullName *string `json:"ShopFullName,omitempty" name:"ShopFullName"`

	// 品牌名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	BrandName *string `json:"BrandName,omitempty" name:"BrandName"`

	// 详细地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Address *string `json:"Address,omitempty" name:"Address"`

	// open_key
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 商户名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	MerchantName *string `json:"MerchantName,omitempty" name:"MerchantName"`
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

// Predefined struct for user
type QuerySinglePayRequestParams struct {
	// 业务流水号
	SerialNumber *string `json:"SerialNumber,omitempty" name:"SerialNumber"`

	// 接入环境。沙箱环境填sandbox
	Profile *string `json:"Profile,omitempty" name:"Profile"`
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

// Predefined struct for user
type QuerySinglePayResponseParams struct {
	// 返回结果
	Result *QuerySinglePayResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QuerySinglePayResponse struct {
	*tchttp.BaseResponse
	Response *QuerySinglePayResponseParams `json:"Response"`
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

type QuerySinglePaymentResultData struct {
	// 平台交易流水号，唯一
	// 注意：此字段可能返回 null，表示取不到有效值。
	TradeSerialNo *string `json:"TradeSerialNo,omitempty" name:"TradeSerialNo"`

	// 订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 交易状态。
	// 0 处理中  
	// 1 预占成功 
	// 2 交易成功 
	// 3 交易失败 
	// 4 未知渠道异常 
	// 5 预占额度失败
	// 6 提交成功
	// 7 提交失败
	// 8 订单重复提交
	// 99 未知系统异常
	// 注意：此字段可能返回 null，表示取不到有效值。
	TradeStatus *int64 `json:"TradeStatus,omitempty" name:"TradeStatus"`

	// 业务备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 代理商ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`

	// 代理商名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentName *string `json:"AgentName,omitempty" name:"AgentName"`

	// 交易状态描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TradeStatusDesc *string `json:"TradeStatusDesc,omitempty" name:"TradeStatusDesc"`
}

// Predefined struct for user
type QuerySinglePaymentResultRequestParams struct {
	// 转账类型
	TransferType *int64 `json:"TransferType,omitempty" name:"TransferType"`

	// 交易流水流水号，唯一
	TradeSerialNo *string `json:"TradeSerialNo,omitempty" name:"TradeSerialNo"`

	// 订单号，与TradeSerialNo不能同时为空
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`
}

type QuerySinglePaymentResultRequest struct {
	*tchttp.BaseRequest
	
	// 转账类型
	TransferType *int64 `json:"TransferType,omitempty" name:"TransferType"`

	// 交易流水流水号，唯一
	TradeSerialNo *string `json:"TradeSerialNo,omitempty" name:"TradeSerialNo"`

	// 订单号，与TradeSerialNo不能同时为空
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`
}

func (r *QuerySinglePaymentResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QuerySinglePaymentResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TransferType")
	delete(f, "TradeSerialNo")
	delete(f, "OrderId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QuerySinglePaymentResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QuerySinglePaymentResultResponseParams struct {
	// 错误码。响应成功："SUCCESS"，其他为不成功
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 响应消息。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回响应
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *QuerySinglePaymentResultData `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QuerySinglePaymentResultResponse struct {
	*tchttp.BaseResponse
	Response *QuerySinglePaymentResultResponseParams `json:"Response"`
}

func (r *QuerySinglePaymentResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QuerySinglePaymentResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QuerySingleTransactionStatusRequestParams struct {
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

// Predefined struct for user
type QuerySingleTransactionStatusResponseParams struct {
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
}

type QuerySingleTransactionStatusResponse struct {
	*tchttp.BaseResponse
	Response *QuerySingleTransactionStatusResponseParams `json:"Response"`
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

// Predefined struct for user
type QuerySmallAmountTransferRequestParams struct {
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

// Predefined struct for user
type QuerySmallAmountTransferResponseParams struct {
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
}

type QuerySmallAmountTransferResponse struct {
	*tchttp.BaseResponse
	Response *QuerySmallAmountTransferResponseParams `json:"Response"`
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

// Predefined struct for user
type QueryTradeRequestParams struct {
	// 贸易材料流水号
	TradeFileId *string `json:"TradeFileId,omitempty" name:"TradeFileId"`

	// 接入环境。沙箱环境填sandbox
	Profile *string `json:"Profile,omitempty" name:"Profile"`
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

// Predefined struct for user
type QueryTradeResponseParams struct {
	// 贸易材料明细查询结果
	Result *QueryTradeResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryTradeResponse struct {
	*tchttp.BaseResponse
	Response *QueryTradeResponseParams `json:"Response"`
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

// Predefined struct for user
type QueryTransferBatchRequestParams struct {
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

// Predefined struct for user
type QueryTransferBatchResponseParams struct {
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
}

type QueryTransferBatchResponse struct {
	*tchttp.BaseResponse
	Response *QueryTransferBatchResponseParams `json:"Response"`
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

// Predefined struct for user
type QueryTransferDetailRequestParams struct {
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

// Predefined struct for user
type QueryTransferDetailResponseParams struct {
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
}

type QueryTransferDetailResponse struct {
	*tchttp.BaseResponse
	Response *QueryTransferDetailResponseParams `json:"Response"`
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

// Predefined struct for user
type QueryTransferResultRequestParams struct {
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

// Predefined struct for user
type QueryTransferResultResponseParams struct {
	// 错误码。响应成功："SUCCESS"，其他为不成功
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 响应消息
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *QueryTransferResultData `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryTransferResultResponse struct {
	*tchttp.BaseResponse
	Response *QueryTransferResultResponseParams `json:"Response"`
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

// Predefined struct for user
type RechargeByThirdPayRequestParams struct {
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

// Predefined struct for user
type RechargeByThirdPayResponseParams struct {
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
}

type RechargeByThirdPayResponse struct {
	*tchttp.BaseResponse
	Response *RechargeByThirdPayResponseParams `json:"Response"`
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

// Predefined struct for user
type RechargeMemberThirdPayRequestParams struct {
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

// Predefined struct for user
type RechargeMemberThirdPayResponseParams struct {
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
}

type RechargeMemberThirdPayResponse struct {
	*tchttp.BaseResponse
	Response *RechargeMemberThirdPayResponseParams `json:"Response"`
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

// Predefined struct for user
type RefundCloudOrderRequestParams struct {
	// 米大师分配的支付主MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 用户Id，长度不小于5位，仅支持字母和数字的组合
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 退款订单号，仅支持数字、字母、下划线（_）、横杠字符（-）、点（.）的组合
	RefundId *string `json:"RefundId,omitempty" name:"RefundId"`

	// 退款金额，单位：分
	// 当该字段为空或者为0时，系统会默认使用订单当实付金额作为退款金额
	TotalRefundAmt *int64 `json:"TotalRefundAmt,omitempty" name:"TotalRefundAmt"`

	// 商品订单，仅支持数字、字母、下划线（_）、横杠字符（-）、点（.）的组合
	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`

	// 平台应收金额，单位：分
	PlatformRefundAmt *int64 `json:"PlatformRefundAmt,omitempty" name:"PlatformRefundAmt"`

	// 结算应收金额，单位：分
	MchRefundAmt *int64 `json:"MchRefundAmt,omitempty" name:"MchRefundAmt"`

	// 支持多个子订单批量退款单个子订单退款支持传SubOutTradeNo
	// 也支持传SubOrderRefundList，都传的时候以SubOrderRefundList为准。
	// 如果传了子单退款细节，外部不需要再传退款金额，平台应退，商户应退金额
	SubOrderRefundList []*CloudSubOrderRefund `json:"SubOrderRefundList,omitempty" name:"SubOrderRefundList"`

	// 渠道订单号，当出现重复支付时，可以将重复支付订单的渠道订单号传入，以进行退款（注意：目前该重复支付订单的渠道订单号仅能通过米大师内部获取），更多重复支付订单退款说明，请参考[重复支付订单退款说明](https://dev.tke.midas.qq.com/juxin-doc-next/apidocs/receive-order/Refund.html#%E9%87%8D%E5%A4%8D%E6%94%AF%E4%BB%98%E8%AE%A2%E5%8D%95%E9%80%80%E6%AC%BE%E8%AF%B4%E6%98%8E)
	ChannelOrderId *string `json:"ChannelOrderId,omitempty" name:"ChannelOrderId"`

	// 通知地址
	RefundNotifyUrl *string `json:"RefundNotifyUrl,omitempty" name:"RefundNotifyUrl"`

	// 透传字段，退款成功回调透传给应用，用于开发者透传自定义内容
	Metadata *string `json:"Metadata,omitempty" name:"Metadata"`

	// 渠道扩展退款促销列表，可将各个渠道的退款促销信息放于该列表
	ExternalRefundPromptGroupList *string `json:"ExternalRefundPromptGroupList,omitempty" name:"ExternalRefundPromptGroupList"`
}

type RefundCloudOrderRequest struct {
	*tchttp.BaseRequest
	
	// 米大师分配的支付主MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 用户Id，长度不小于5位，仅支持字母和数字的组合
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 退款订单号，仅支持数字、字母、下划线（_）、横杠字符（-）、点（.）的组合
	RefundId *string `json:"RefundId,omitempty" name:"RefundId"`

	// 退款金额，单位：分
	// 当该字段为空或者为0时，系统会默认使用订单当实付金额作为退款金额
	TotalRefundAmt *int64 `json:"TotalRefundAmt,omitempty" name:"TotalRefundAmt"`

	// 商品订单，仅支持数字、字母、下划线（_）、横杠字符（-）、点（.）的组合
	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`

	// 平台应收金额，单位：分
	PlatformRefundAmt *int64 `json:"PlatformRefundAmt,omitempty" name:"PlatformRefundAmt"`

	// 结算应收金额，单位：分
	MchRefundAmt *int64 `json:"MchRefundAmt,omitempty" name:"MchRefundAmt"`

	// 支持多个子订单批量退款单个子订单退款支持传SubOutTradeNo
	// 也支持传SubOrderRefundList，都传的时候以SubOrderRefundList为准。
	// 如果传了子单退款细节，外部不需要再传退款金额，平台应退，商户应退金额
	SubOrderRefundList []*CloudSubOrderRefund `json:"SubOrderRefundList,omitempty" name:"SubOrderRefundList"`

	// 渠道订单号，当出现重复支付时，可以将重复支付订单的渠道订单号传入，以进行退款（注意：目前该重复支付订单的渠道订单号仅能通过米大师内部获取），更多重复支付订单退款说明，请参考[重复支付订单退款说明](https://dev.tke.midas.qq.com/juxin-doc-next/apidocs/receive-order/Refund.html#%E9%87%8D%E5%A4%8D%E6%94%AF%E4%BB%98%E8%AE%A2%E5%8D%95%E9%80%80%E6%AC%BE%E8%AF%B4%E6%98%8E)
	ChannelOrderId *string `json:"ChannelOrderId,omitempty" name:"ChannelOrderId"`

	// 通知地址
	RefundNotifyUrl *string `json:"RefundNotifyUrl,omitempty" name:"RefundNotifyUrl"`

	// 透传字段，退款成功回调透传给应用，用于开发者透传自定义内容
	Metadata *string `json:"Metadata,omitempty" name:"Metadata"`

	// 渠道扩展退款促销列表，可将各个渠道的退款促销信息放于该列表
	ExternalRefundPromptGroupList *string `json:"ExternalRefundPromptGroupList,omitempty" name:"ExternalRefundPromptGroupList"`
}

func (r *RefundCloudOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefundCloudOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MidasAppId")
	delete(f, "UserId")
	delete(f, "RefundId")
	delete(f, "TotalRefundAmt")
	delete(f, "OutTradeNo")
	delete(f, "MidasEnvironment")
	delete(f, "PlatformRefundAmt")
	delete(f, "MchRefundAmt")
	delete(f, "SubOrderRefundList")
	delete(f, "ChannelOrderId")
	delete(f, "RefundNotifyUrl")
	delete(f, "Metadata")
	delete(f, "ExternalRefundPromptGroupList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RefundCloudOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RefundCloudOrderResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RefundCloudOrderResponse struct {
	*tchttp.BaseResponse
	Response *RefundCloudOrderResponseParams `json:"Response"`
}

func (r *RefundCloudOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefundCloudOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RefundMemberTransactionRequestParams struct {
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

// Predefined struct for user
type RefundMemberTransactionResponseParams struct {
	// 请求类型
	RequestType *string `json:"RequestType,omitempty" name:"RequestType"`

	// 银行流水号
	FrontSequenceNumber *string `json:"FrontSequenceNumber,omitempty" name:"FrontSequenceNumber"`

	// 保留域
	ReservedMessage *string `json:"ReservedMessage,omitempty" name:"ReservedMessage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RefundMemberTransactionResponse struct {
	*tchttp.BaseResponse
	Response *RefundMemberTransactionResponseParams `json:"Response"`
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

// Predefined struct for user
type RefundOpenBankOrderRequestParams struct {
	// 外部商户退款单号。
	OutRefundId *string `json:"OutRefundId,omitempty" name:"OutRefundId"`

	// 退款金额。单位分。
	RefundAmount *int64 `json:"RefundAmount,omitempty" name:"RefundAmount"`

	// 渠道商户号。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 外部商户订单号，与云企付渠道订单号二者选填其一。
	OutOrderId *string `json:"OutOrderId,omitempty" name:"OutOrderId"`

	// 云企付渠道订单号，与外部订单号二者选填其一。
	ChannelOrderId *string `json:"ChannelOrderId,omitempty" name:"ChannelOrderId"`

	// 退款通知地址。
	NotifyUrl *string `json:"NotifyUrl,omitempty" name:"NotifyUrl"`

	// 退款原因。
	RefundReason *string `json:"RefundReason,omitempty" name:"RefundReason"`

	// 第三方渠道退款附加信息。详见附录-复杂类型。
	// 若未作特殊说明，则无需传入。
	ExternalRefundData *string `json:"ExternalRefundData,omitempty" name:"ExternalRefundData"`

	// 备注信息
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type RefundOpenBankOrderRequest struct {
	*tchttp.BaseRequest
	
	// 外部商户退款单号。
	OutRefundId *string `json:"OutRefundId,omitempty" name:"OutRefundId"`

	// 退款金额。单位分。
	RefundAmount *int64 `json:"RefundAmount,omitempty" name:"RefundAmount"`

	// 渠道商户号。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 外部商户订单号，与云企付渠道订单号二者选填其一。
	OutOrderId *string `json:"OutOrderId,omitempty" name:"OutOrderId"`

	// 云企付渠道订单号，与外部订单号二者选填其一。
	ChannelOrderId *string `json:"ChannelOrderId,omitempty" name:"ChannelOrderId"`

	// 退款通知地址。
	NotifyUrl *string `json:"NotifyUrl,omitempty" name:"NotifyUrl"`

	// 退款原因。
	RefundReason *string `json:"RefundReason,omitempty" name:"RefundReason"`

	// 第三方渠道退款附加信息。详见附录-复杂类型。
	// 若未作特殊说明，则无需传入。
	ExternalRefundData *string `json:"ExternalRefundData,omitempty" name:"ExternalRefundData"`

	// 备注信息
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *RefundOpenBankOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefundOpenBankOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OutRefundId")
	delete(f, "RefundAmount")
	delete(f, "ChannelMerchantId")
	delete(f, "OutOrderId")
	delete(f, "ChannelOrderId")
	delete(f, "NotifyUrl")
	delete(f, "RefundReason")
	delete(f, "ExternalRefundData")
	delete(f, "Remark")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RefundOpenBankOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RefundOpenBankOrderResponseParams struct {
	// 错误码
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *OpenBankRefundOrderApplyResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RefundOpenBankOrderResponse struct {
	*tchttp.BaseResponse
	Response *RefundOpenBankOrderResponseParams `json:"Response"`
}

func (r *RefundOpenBankOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefundOpenBankOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RefundOrderRequestParams struct {
	// 进件成功后返给商户方的AppId
	MerchantAppId *string `json:"MerchantAppId,omitempty" name:"MerchantAppId"`

	// 平台流水号。消费订单发起成功后，返回的平台唯一订单号。
	OrderNo *string `json:"OrderNo,omitempty" name:"OrderNo"`
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

// Predefined struct for user
type RefundOrderResponseParams struct {
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
}

type RefundOrderResponse struct {
	*tchttp.BaseResponse
	Response *RefundOrderResponseParams `json:"Response"`
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

type RefundOrderResult struct {
	// 付款订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderNo *string `json:"OrderNo,omitempty" name:"OrderNo"`

	// 开发者流水号
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeveloperNo *string `json:"DeveloperNo,omitempty" name:"DeveloperNo"`

	// 交易优惠金额（免充值券）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TradeDiscountAmount *string `json:"TradeDiscountAmount,omitempty" name:"TradeDiscountAmount"`

	// 付款方式名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayName *string `json:"PayName,omitempty" name:"PayName"`

	// 商户流水号（从1开始自增长不重复）
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderMerchantId *string `json:"OrderMerchantId,omitempty" name:"OrderMerchantId"`

	// 实际交易金额（以分为单位，没有小数点）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TradeAmount *string `json:"TradeAmount,omitempty" name:"TradeAmount"`

	// 币种签名
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrencySign *string `json:"CurrencySign,omitempty" name:"CurrencySign"`

	// 付款完成时间（以收单机构为准）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TradePayTime *string `json:"TradePayTime,omitempty" name:"TradePayTime"`

	// 门店流水号（从1开始自增长不重复）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShopOrderId *string `json:"ShopOrderId,omitempty" name:"ShopOrderId"`

	// 支付标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayTag *string `json:"PayTag,omitempty" name:"PayTag"`

	// 订单状态（1交易成功，2待支付，4已取消，9等待用户输入密码确认
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 币种代码
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderCurrency *string `json:"OrderCurrency,omitempty" name:"OrderCurrency"`

	// 开始交易时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TradeTime *string `json:"TradeTime,omitempty" name:"TradeTime"`

	// 折扣金额（以分为单位，没有小数点）
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiscountAmount *string `json:"DiscountAmount,omitempty" name:"DiscountAmount"`

	// 原始订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginalOrderNo *string `json:"OriginalOrderNo,omitempty" name:"OriginalOrderNo"`
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

// Predefined struct for user
type RefundRequestParams struct {
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

// Predefined struct for user
type RefundResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RefundResponse struct {
	*tchttp.BaseResponse
	Response *RefundResponseParams `json:"Response"`
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

// Predefined struct for user
type RefundTlinxOrderRequestParams struct {
	// 使用门店OpenId
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 使用门店OpenKey
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 原始订单的开发者交易流水号
	DeveloperNo *string `json:"DeveloperNo,omitempty" name:"DeveloperNo"`

	// 新退款订单的开发者流水号，同一门店内唯一
	RefundOutNo *string `json:"RefundOutNo,omitempty" name:"RefundOutNo"`

	// 退款订单名称，可以为空
	RefundOrderName *string `json:"RefundOrderName,omitempty" name:"RefundOrderName"`

	// 退款金额（以分为单位，没有小数点）
	RefundAmount *string `json:"RefundAmount,omitempty" name:"RefundAmount"`

	// 主管密码，对密码进行SHA-1加密，默认为123456
	ShopPassword *string `json:"ShopPassword,omitempty" name:"ShopPassword"`

	// 退款备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

type RefundTlinxOrderRequest struct {
	*tchttp.BaseRequest
	
	// 使用门店OpenId
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 使用门店OpenKey
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 原始订单的开发者交易流水号
	DeveloperNo *string `json:"DeveloperNo,omitempty" name:"DeveloperNo"`

	// 新退款订单的开发者流水号，同一门店内唯一
	RefundOutNo *string `json:"RefundOutNo,omitempty" name:"RefundOutNo"`

	// 退款订单名称，可以为空
	RefundOrderName *string `json:"RefundOrderName,omitempty" name:"RefundOrderName"`

	// 退款金额（以分为单位，没有小数点）
	RefundAmount *string `json:"RefundAmount,omitempty" name:"RefundAmount"`

	// 主管密码，对密码进行SHA-1加密，默认为123456
	ShopPassword *string `json:"ShopPassword,omitempty" name:"ShopPassword"`

	// 退款备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *RefundTlinxOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefundTlinxOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OpenId")
	delete(f, "OpenKey")
	delete(f, "DeveloperNo")
	delete(f, "RefundOutNo")
	delete(f, "RefundOrderName")
	delete(f, "RefundAmount")
	delete(f, "ShopPassword")
	delete(f, "Remark")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RefundTlinxOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RefundTlinxOrderResponseParams struct {
	// 业务系统返回码，0表示成功，其他表示失败。
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 业务系统返回消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 退款响应对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *RefundOrderResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RefundTlinxOrderResponse struct {
	*tchttp.BaseResponse
	Response *RefundTlinxOrderResponseParams `json:"Response"`
}

func (r *RefundTlinxOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefundTlinxOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterBehaviorRequestParams struct {
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

// Predefined struct for user
type RegisterBehaviorResponseParams struct {
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
}

type RegisterBehaviorResponse struct {
	*tchttp.BaseResponse
	Response *RegisterBehaviorResponseParams `json:"Response"`
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

// Predefined struct for user
type RegisterBillRequestParams struct {
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

// Predefined struct for user
type RegisterBillResponseParams struct {
	// 银行流水号
	FrontSeqNo *string `json:"FrontSeqNo,omitempty" name:"FrontSeqNo"`

	// 保留字段
	ReservedMessage *string `json:"ReservedMessage,omitempty" name:"ReservedMessage"`

	// 请求类型
	RequestType *string `json:"RequestType,omitempty" name:"RequestType"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RegisterBillResponse struct {
	*tchttp.BaseResponse
	Response *RegisterBillResponseParams `json:"Response"`
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

// Predefined struct for user
type RegisterBillSupportWithdrawRequestParams struct {
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

// Predefined struct for user
type RegisterBillSupportWithdrawResponseParams struct {
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
}

type RegisterBillSupportWithdrawResponse struct {
	*tchttp.BaseResponse
	Response *RegisterBillSupportWithdrawResponseParams `json:"Response"`
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

// Predefined struct for user
type RevResigterBillSupportWithdrawRequestParams struct {
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

// Predefined struct for user
type RevResigterBillSupportWithdrawResponseParams struct {
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
}

type RevResigterBillSupportWithdrawResponse struct {
	*tchttp.BaseResponse
	Response *RevResigterBillSupportWithdrawResponseParams `json:"Response"`
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

// Predefined struct for user
type ReviseMbrPropertyRequestParams struct {
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

// Predefined struct for user
type ReviseMbrPropertyResponseParams struct {
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
}

type ReviseMbrPropertyResponse struct {
	*tchttp.BaseResponse
	Response *ReviseMbrPropertyResponseParams `json:"Response"`
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

// Predefined struct for user
type RevokeMemberRechargeThirdPayRequestParams struct {
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

// Predefined struct for user
type RevokeMemberRechargeThirdPayResponseParams struct {
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
}

type RevokeMemberRechargeThirdPayResponse struct {
	*tchttp.BaseResponse
	Response *RevokeMemberRechargeThirdPayResponseParams `json:"Response"`
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

// Predefined struct for user
type RevokeRechargeByThirdPayRequestParams struct {
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

// Predefined struct for user
type RevokeRechargeByThirdPayResponseParams struct {
	// 请求类型
	RequestType *string `json:"RequestType,omitempty" name:"RequestType"`

	// 保留域
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReservedMessage *string `json:"ReservedMessage,omitempty" name:"ReservedMessage"`

	// 银行流水号
	FrontSequenceNumber *string `json:"FrontSequenceNumber,omitempty" name:"FrontSequenceNumber"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RevokeRechargeByThirdPayResponse struct {
	*tchttp.BaseResponse
	Response *RevokeRechargeByThirdPayResponseParams `json:"Response"`
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

type SettleInfo struct {
	// 结算账户类型 
	// PRIVATE：对私 
	// BUSINESS：对公
	// HELIPAY渠道必传
	SettleAccountType *string `json:"SettleAccountType,omitempty" name:"SettleAccountType"`

	// 结算账号
	// HELIPAY渠道必传
	SettleAccountNumber *string `json:"SettleAccountNumber,omitempty" name:"SettleAccountNumber"`

	// 结算账户名称
	// HELIPAY渠道必传
	SettleAccountName *string `json:"SettleAccountName,omitempty" name:"SettleAccountName"`

	// 支行号
	// HELIPAY渠道必传
	BankBranchId *string `json:"BankBranchId,omitempty" name:"BankBranchId"`

	// 支行名称
	BankBranchName *string `json:"BankBranchName,omitempty" name:"BankBranchName"`

	// 结算方式 
	// AUTO：自动结算 
	// SELF：自主结算
	// HELIPAY渠道必传
	SettleMode *string `json:"SettleMode,omitempty" name:"SettleMode"`

	// 结算周期 
	// T1：工作日隔天结算 
	// D1：自然日隔天结算 
	// D0：当日结算
	// HELIPAY渠道必传
	SettlePeriod *string `json:"SettlePeriod,omitempty" name:"SettlePeriod"`
}

type SettlementOrderResult struct {
	// 收入类型
	// LABOR:劳务所得
	// OCCASION:偶然所得
	IncomeType *string `json:"IncomeType,omitempty" name:"IncomeType"`

	// 税前金额
	AmountBeforeTax *string `json:"AmountBeforeTax,omitempty" name:"AmountBeforeTax"`

	// 税后金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	AmountAfterTax *string `json:"AmountAfterTax,omitempty" name:"AmountAfterTax"`

	// 税金
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tax *string `json:"Tax,omitempty" name:"Tax"`

	// 外部订单ID
	OutOrderId *string `json:"OutOrderId,omitempty" name:"OutOrderId"`

	// 订单ID
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 发起时间
	InitiateTime *string `json:"InitiateTime,omitempty" name:"InitiateTime"`

	// 完成时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`

	// 状态
	// ACCEPTED:已受理
	// ACCOUNTED:已记账
	// SUCCEED:已成功
	// FAILED:已失败
	Status *string `json:"Status,omitempty" name:"Status"`

	// 状态描述
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type SettlementOrders struct {
	// 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*SettlementOrderResult `json:"List,omitempty" name:"List"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type SupportBankInfo struct {
	// 银行简称。
	BankCode *string `json:"BankCode,omitempty" name:"BankCode"`

	// 银行名称。
	BankName *string `json:"BankName,omitempty" name:"BankName"`

	// 状态。
	// __MAINTAINING__: 维护中
	// __WORKING__: 正常工作
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaintainStatus *string `json:"MaintainStatus,omitempty" name:"MaintainStatus"`

	// 银行渠道维护公告。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BankNotice *string `json:"BankNotice,omitempty" name:"BankNotice"`

	// 支持银行代码
	// 注意：此字段可能返回 null，表示取不到有效值。
	BankId *string `json:"BankId,omitempty" name:"BankId"`

	// 卡类型。
	// D：借记卡，C：信用卡，Z：借贷合一卡。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CardType *string `json:"CardType,omitempty" name:"CardType"`
}

// Predefined struct for user
type SyncContractDataRequestParams struct {
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

// Predefined struct for user
type SyncContractDataResponseParams struct {
	// 请求处理信息
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SyncContractDataResponse struct {
	*tchttp.BaseResponse
	Response *SyncContractDataResponseParams `json:"Response"`
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

// Predefined struct for user
type TerminateContractRequestParams struct {
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

// Predefined struct for user
type TerminateContractResponseParams struct {
	// 解约数据
	ContractTerminateData *ResponseTerminateContract `json:"ContractTerminateData,omitempty" name:"ContractTerminateData"`

	// 请求处理信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type TerminateContractResponse struct {
	*tchttp.BaseResponse
	Response *TerminateContractResponseParams `json:"Response"`
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

// Predefined struct for user
type TransferSinglePayRequestParams struct {
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

// Predefined struct for user
type TransferSinglePayResponseParams struct {
	// 错误码。响应成功："SUCCESS"，其他为不成功
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 响应消息
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *TransferSinglePayData `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type TransferSinglePayResponse struct {
	*tchttp.BaseResponse
	Response *TransferSinglePayResponseParams `json:"Response"`
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

// Predefined struct for user
type UnBindAcctRequestParams struct {
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

// Predefined struct for user
type UnBindAcctResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnBindAcctResponse struct {
	*tchttp.BaseResponse
	Response *UnBindAcctResponseParams `json:"Response"`
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

// Predefined struct for user
type UnbindOpenBankExternalSubMerchantBankAccountRequestParams struct {
	// 渠道商户ID。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 渠道子商户ID。
	ChannelSubMerchantId *string `json:"ChannelSubMerchantId,omitempty" name:"ChannelSubMerchantId"`

	// 渠道名称。
	// __TENPAY__: 商企付
	// __WECHAT__: 微信支付
	// __ALIPAY__: 支付宝
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 支付方式。
	// __EBANK_PAYMENT__: ebank支付
	// __OPENBANK_PAYMENT__: openbank支付
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`

	// 绑卡序列号。
	BindSerialNo *string `json:"BindSerialNo,omitempty" name:"BindSerialNo"`

	// 外部申请编号。
	OutApplyId *string `json:"OutApplyId,omitempty" name:"OutApplyId"`

	// 通知地址。
	NotifyUrl *string `json:"NotifyUrl,omitempty" name:"NotifyUrl"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type UnbindOpenBankExternalSubMerchantBankAccountRequest struct {
	*tchttp.BaseRequest
	
	// 渠道商户ID。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 渠道子商户ID。
	ChannelSubMerchantId *string `json:"ChannelSubMerchantId,omitempty" name:"ChannelSubMerchantId"`

	// 渠道名称。
	// __TENPAY__: 商企付
	// __WECHAT__: 微信支付
	// __ALIPAY__: 支付宝
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 支付方式。
	// __EBANK_PAYMENT__: ebank支付
	// __OPENBANK_PAYMENT__: openbank支付
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`

	// 绑卡序列号。
	BindSerialNo *string `json:"BindSerialNo,omitempty" name:"BindSerialNo"`

	// 外部申请编号。
	OutApplyId *string `json:"OutApplyId,omitempty" name:"OutApplyId"`

	// 通知地址。
	NotifyUrl *string `json:"NotifyUrl,omitempty" name:"NotifyUrl"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *UnbindOpenBankExternalSubMerchantBankAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindOpenBankExternalSubMerchantBankAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelMerchantId")
	delete(f, "ChannelSubMerchantId")
	delete(f, "ChannelName")
	delete(f, "PaymentMethod")
	delete(f, "BindSerialNo")
	delete(f, "OutApplyId")
	delete(f, "NotifyUrl")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindOpenBankExternalSubMerchantBankAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindOpenBankExternalSubMerchantBankAccountResponseParams struct {
	// 错误码。
	// __SUCCESS__: 成功
	// __其他__: 见附录-错误码表
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *UnbindOpenBankExternalSubMerchantBankAccountResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnbindOpenBankExternalSubMerchantBankAccountResponse struct {
	*tchttp.BaseResponse
	Response *UnbindOpenBankExternalSubMerchantBankAccountResponseParams `json:"Response"`
}

func (r *UnbindOpenBankExternalSubMerchantBankAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindOpenBankExternalSubMerchantBankAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnbindOpenBankExternalSubMerchantBankAccountResult struct {
	// 渠道申请编号。
	ChannelApplyId *string `json:"ChannelApplyId,omitempty" name:"ChannelApplyId"`

	// 解绑状态。
	// __SUCCESS__: 解绑成功
	// __FAILED__: 解绑失败
	// __PROCESSING__: 解绑中
	// 注意：若返回解绑中，需要再次调用解绑结果查询接口查询结果。
	UnbindStatus *string `json:"UnbindStatus,omitempty" name:"UnbindStatus"`

	// 解绑返回描述, 例如失败原因等。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnbindMessage *string `json:"UnbindMessage,omitempty" name:"UnbindMessage"`
}

// Predefined struct for user
type UnbindRelateAcctRequestParams struct {
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

// Predefined struct for user
type UnbindRelateAcctResponseParams struct {
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
}

type UnbindRelateAcctResponse struct {
	*tchttp.BaseResponse
	Response *UnbindRelateAcctResponseParams `json:"Response"`
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

// Predefined struct for user
type UnifiedCloudOrderRequestParams struct {
	// 米大师分配的支付主MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 用户ID
	// 长度不小于5位，仅支持字母和数字的组合，长度限制以具体接入渠道为准
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 开发者主订单号
	// 支付订单号，仅支持数字、字母、下划线（_）、横杠字符（-）、点（.）的组合，长度供参考，部分渠道存在长度更短的情况接入时请联系开发咨询
	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`

	// 货币类型
	// ISO货币代码，CNY
	CurrencyType *string `json:"CurrencyType,omitempty" name:"CurrencyType"`

	// 商品ID
	// 业务自定义的商品id，仅支持数字、字母、下划线（_）、横杠字符（-）、点（.）的组合。
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 商品名称
	// 业务自定义的商品名称，无需URL编码，长度限制以具体所接入渠道为准。
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 商品详情
	// 业务自定义的商品详情，无需URL编码，长度限制以具体所接入渠道为准。
	ProductDetail *string `json:"ProductDetail,omitempty" name:"ProductDetail"`

	// 原始金额
	// 单位：分，需要注意的是，OriginalAmt>=TotalAmt
	OriginalAmt *int64 `json:"OriginalAmt,omitempty" name:"OriginalAmt"`

	// 支付金额
	// 单位：分，需要注意的是，TotalAmt=TotalPlatformIncome+TotalMchIncome。
	TotalAmt *int64 `json:"TotalAmt,omitempty" name:"TotalAmt"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`

	// 支付SubAppId
	// 米大师计费SubAppId，代表子商户。指定使用该商户的商户号下单时必传。
	SubAppId *string `json:"SubAppId,omitempty" name:"SubAppId"`

	// 顶层支付渠道
	// 银行收单:
	// openbank_ccb: 建设银行
	// openbank_icbc: 工商银行
	// openbank_cmb: 招商银行
	// openbank_ping: 平安银行
	// openbank_icbc_jft：工商银行聚付通
	// 非银行收单，可以为空
	RealChannel *string `json:"RealChannel,omitempty" name:"RealChannel"`

	// 支付渠道
	// wechat：微信支付
	// wechat_ecommerce: 微信电商收付通
	// open_alipay: 支付宝
	// open_quickpass: 银联云闪付
	// icbc_epay: 工银e支付
	// foreign_cardpay: 外卡支付
	// icbc_jft_wechat: 工行聚付通-微信
	// icbc_jft_alipay: 工行聚付通-支付宝
	// icbc_jft_epay: 工行聚付通-e支付
	// 指定渠道下单时必传
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// 透传字段
	// 支付成功回调透传给应用，用于开发者透传自定义内容。
	Metadata *string `json:"Metadata,omitempty" name:"Metadata"`

	// 数量
	// 购买数量,不传默认为1。
	Quantity *int64 `json:"Quantity,omitempty" name:"Quantity"`

	// Web端回调地址
	// Web端网页回调地址，仅当Web端SDK使用页面跳转方式时有效。
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 支付取消地址
	CancelUrl *string `json:"CancelUrl,omitempty" name:"CancelUrl"`

	// 微信AppId
	// wechat渠道或wchat_ecommerce渠道可以指定下单时的wxappid。
	WxAppId *string `json:"WxAppId,omitempty" name:"WxAppId"`

	// 微信SubAppId
	// wechat渠道可以指定下单时的sub_appid。
	WxSubAppId *string `json:"WxSubAppId,omitempty" name:"WxSubAppId"`

	// 微信公众号/小程序OpenId
	// 微信公众号/小程序支付时为必选，需要传微信下的openid。
	WxOpenId *string `json:"WxOpenId,omitempty" name:"WxOpenId"`

	// 微信公众号/小程序SubOpenId
	// 在服务商模式下，微信公众号/小程序支付时wx_sub_openid和wx_openid二选一。
	WxSubOpenId *string `json:"WxSubOpenId,omitempty" name:"WxSubOpenId"`

	// 平台应收金额
	// 单位：分，需要注意的是，TotalAmt=TotalPlatformIncome+TotalMchIncome
	TotalPlatformIncome *int64 `json:"TotalPlatformIncome,omitempty" name:"TotalPlatformIncome"`

	// 结算应收金额
	// 单位：分，需要注意的是，TotalAmt=TotalPlatformIncome+TotalMchIncome
	TotalMchIncome *int64 `json:"TotalMchIncome,omitempty" name:"TotalMchIncome"`

	// 子订单列表
	// 格式：子订单号、子应用Id、金额。压缩后最长不可超过32K字节(去除空格，换行，制表符等无意义字符)。
	SubOrderList []*CloudSubOrder `json:"SubOrderList,omitempty" name:"SubOrderList"`

	// 结算信息
	// 例如是否需要分账、是否需要支付确认等，
	// 注意：如果子单列表中传入了SettleInfo，在主单中不可再传入SettleInfo字段。
	SettleInfo *CloudSettleInfo `json:"SettleInfo,omitempty" name:"SettleInfo"`

	// 附加项信息列表
	// 例如溢价信息、抵扣信息、积分信息、补贴信息
	// 通过该字段可以实现渠道方的优惠抵扣补贴等营销功能
	// 注意：当传SubOrderList时，请在子单信息中传附加项信息，不要在主单中传该字段。
	AttachmentInfoList []*CloudAttachmentInfo `json:"AttachmentInfoList,omitempty" name:"AttachmentInfoList"`

	// 支付通知地址
	// 调用方可通过该字段传入自定义支付通知地址。
	PaymentNotifyUrl *string `json:"PaymentNotifyUrl,omitempty" name:"PaymentNotifyUrl"`

	// 支付场景
	// 需要结合 RealChannel和Channel字段使用可选值:
	// wechat-app 微信APP支付方式
	// wechat-mini 微信小程序支付，示例：当 RealChannel=wechat Channel=wechat PayScene=wechat-mini时，内部会直接以小程序方式调用微信统一下单接口。
	PayScene *string `json:"PayScene,omitempty" name:"PayScene"`

	// 语言代码
	// 取值请参考[ISO 639-1代码表](https://zh.wikipedia.org/zh-cn/ISO_639-1%E4%BB%A3%E7%A0%81%E8%A1%A8)
	LocaleCode *string `json:"LocaleCode,omitempty" name:"LocaleCode"`

	// 地区代码
	// 取值请参考[ISO 3166-1二位字母代码表](https://zh.wikipedia.org/zh-cn/ISO_3166-1%E4%BA%8C%E4%BD%8D%E5%AD%97%E6%AF%8D%E4%BB%A3%E7%A0%81#%E6%AD%A3%E5%BC%8F%E5%88%86%E9%85%8D%E4%BB%A3%E7%A0%81)
	RegionCode *string `json:"RegionCode,omitempty" name:"RegionCode"`

	// 用户IP
	// 请求用户的IP地址，特定的渠道或特定的支付方式，此字段为必填
	// wechat_ecommerce渠道 - h5支付方式，此字段必填。
	UserClientIp *string `json:"UserClientIp,omitempty" name:"UserClientIp"`

	// 渠道订单号生成模式
	// 枚举值。决定请求渠道方时的订单号的生成模式，详情请联系米大师沟通。不指定时默认为由米大师自行生成。
	ChannelOrderIdMode *string `json:"ChannelOrderIdMode,omitempty" name:"ChannelOrderIdMode"`

	// 全局支付时间信息
	GlobalPayTimeInfo *CloudGlobalPayTimeInfo `json:"GlobalPayTimeInfo,omitempty" name:"GlobalPayTimeInfo"`

	// 渠道应用ID取用方式
	// USE_APPID 使用渠道应用Id;
	// USE_SUB_APPID 使用子渠道应用Id;
	// USE_APPID_AND_SUB_APPID 既使用渠道应用Id也使用子渠道应用ID。
	ChannelAppIdPolicy *string `json:"ChannelAppIdPolicy,omitempty" name:"ChannelAppIdPolicy"`

	// 门店信息
	// 特定的渠道或特定的支付方式，此字段为必填
	// wechat_ecommerce渠道 - h5支付方式，此字段必填
	StoreInfo *CloudStoreInfo `json:"StoreInfo,omitempty" name:"StoreInfo"`

	// 客户端信息
	// 特定的渠道或特定的支付方式，此字段为必填
	// wechat_ecommerce渠道 - h5支付方式，此字段必填
	ClientInfo *CloudClientInfo `json:"ClientInfo,omitempty" name:"ClientInfo"`

	// 渠道扩展促销列表
	// 可将各个渠道的促销信息放于该列表。
	ExternalPromptGroupList []*CloudExternalPromptGroup `json:"ExternalPromptGroupList,omitempty" name:"ExternalPromptGroupList"`

	// 收单模式
	// ORDER_RECEIVE_MODE_COMMON - 普通支付
	// ORDER_RECEIVE_MODE_COMBINE - 合单支付
	// ORDER_RECEIVE_MODE_V_COMBINE - 虚拟合单支付
	// 若不传入该字段，则会根据是否传入子单来判断是 普通支付 还是 合单支付
	OrderReceiveMode *string `json:"OrderReceiveMode,omitempty" name:"OrderReceiveMode"`
}

type UnifiedCloudOrderRequest struct {
	*tchttp.BaseRequest
	
	// 米大师分配的支付主MidasAppId
	MidasAppId *string `json:"MidasAppId,omitempty" name:"MidasAppId"`

	// 用户ID
	// 长度不小于5位，仅支持字母和数字的组合，长度限制以具体接入渠道为准
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 开发者主订单号
	// 支付订单号，仅支持数字、字母、下划线（_）、横杠字符（-）、点（.）的组合，长度供参考，部分渠道存在长度更短的情况接入时请联系开发咨询
	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`

	// 货币类型
	// ISO货币代码，CNY
	CurrencyType *string `json:"CurrencyType,omitempty" name:"CurrencyType"`

	// 商品ID
	// 业务自定义的商品id，仅支持数字、字母、下划线（_）、横杠字符（-）、点（.）的组合。
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 商品名称
	// 业务自定义的商品名称，无需URL编码，长度限制以具体所接入渠道为准。
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 商品详情
	// 业务自定义的商品详情，无需URL编码，长度限制以具体所接入渠道为准。
	ProductDetail *string `json:"ProductDetail,omitempty" name:"ProductDetail"`

	// 原始金额
	// 单位：分，需要注意的是，OriginalAmt>=TotalAmt
	OriginalAmt *int64 `json:"OriginalAmt,omitempty" name:"OriginalAmt"`

	// 支付金额
	// 单位：分，需要注意的是，TotalAmt=TotalPlatformIncome+TotalMchIncome。
	TotalAmt *int64 `json:"TotalAmt,omitempty" name:"TotalAmt"`

	// 环境类型
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	MidasEnvironment *string `json:"MidasEnvironment,omitempty" name:"MidasEnvironment"`

	// 支付SubAppId
	// 米大师计费SubAppId，代表子商户。指定使用该商户的商户号下单时必传。
	SubAppId *string `json:"SubAppId,omitempty" name:"SubAppId"`

	// 顶层支付渠道
	// 银行收单:
	// openbank_ccb: 建设银行
	// openbank_icbc: 工商银行
	// openbank_cmb: 招商银行
	// openbank_ping: 平安银行
	// openbank_icbc_jft：工商银行聚付通
	// 非银行收单，可以为空
	RealChannel *string `json:"RealChannel,omitempty" name:"RealChannel"`

	// 支付渠道
	// wechat：微信支付
	// wechat_ecommerce: 微信电商收付通
	// open_alipay: 支付宝
	// open_quickpass: 银联云闪付
	// icbc_epay: 工银e支付
	// foreign_cardpay: 外卡支付
	// icbc_jft_wechat: 工行聚付通-微信
	// icbc_jft_alipay: 工行聚付通-支付宝
	// icbc_jft_epay: 工行聚付通-e支付
	// 指定渠道下单时必传
	Channel *string `json:"Channel,omitempty" name:"Channel"`

	// 透传字段
	// 支付成功回调透传给应用，用于开发者透传自定义内容。
	Metadata *string `json:"Metadata,omitempty" name:"Metadata"`

	// 数量
	// 购买数量,不传默认为1。
	Quantity *int64 `json:"Quantity,omitempty" name:"Quantity"`

	// Web端回调地址
	// Web端网页回调地址，仅当Web端SDK使用页面跳转方式时有效。
	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`

	// 支付取消地址
	CancelUrl *string `json:"CancelUrl,omitempty" name:"CancelUrl"`

	// 微信AppId
	// wechat渠道或wchat_ecommerce渠道可以指定下单时的wxappid。
	WxAppId *string `json:"WxAppId,omitempty" name:"WxAppId"`

	// 微信SubAppId
	// wechat渠道可以指定下单时的sub_appid。
	WxSubAppId *string `json:"WxSubAppId,omitempty" name:"WxSubAppId"`

	// 微信公众号/小程序OpenId
	// 微信公众号/小程序支付时为必选，需要传微信下的openid。
	WxOpenId *string `json:"WxOpenId,omitempty" name:"WxOpenId"`

	// 微信公众号/小程序SubOpenId
	// 在服务商模式下，微信公众号/小程序支付时wx_sub_openid和wx_openid二选一。
	WxSubOpenId *string `json:"WxSubOpenId,omitempty" name:"WxSubOpenId"`

	// 平台应收金额
	// 单位：分，需要注意的是，TotalAmt=TotalPlatformIncome+TotalMchIncome
	TotalPlatformIncome *int64 `json:"TotalPlatformIncome,omitempty" name:"TotalPlatformIncome"`

	// 结算应收金额
	// 单位：分，需要注意的是，TotalAmt=TotalPlatformIncome+TotalMchIncome
	TotalMchIncome *int64 `json:"TotalMchIncome,omitempty" name:"TotalMchIncome"`

	// 子订单列表
	// 格式：子订单号、子应用Id、金额。压缩后最长不可超过32K字节(去除空格，换行，制表符等无意义字符)。
	SubOrderList []*CloudSubOrder `json:"SubOrderList,omitempty" name:"SubOrderList"`

	// 结算信息
	// 例如是否需要分账、是否需要支付确认等，
	// 注意：如果子单列表中传入了SettleInfo，在主单中不可再传入SettleInfo字段。
	SettleInfo *CloudSettleInfo `json:"SettleInfo,omitempty" name:"SettleInfo"`

	// 附加项信息列表
	// 例如溢价信息、抵扣信息、积分信息、补贴信息
	// 通过该字段可以实现渠道方的优惠抵扣补贴等营销功能
	// 注意：当传SubOrderList时，请在子单信息中传附加项信息，不要在主单中传该字段。
	AttachmentInfoList []*CloudAttachmentInfo `json:"AttachmentInfoList,omitempty" name:"AttachmentInfoList"`

	// 支付通知地址
	// 调用方可通过该字段传入自定义支付通知地址。
	PaymentNotifyUrl *string `json:"PaymentNotifyUrl,omitempty" name:"PaymentNotifyUrl"`

	// 支付场景
	// 需要结合 RealChannel和Channel字段使用可选值:
	// wechat-app 微信APP支付方式
	// wechat-mini 微信小程序支付，示例：当 RealChannel=wechat Channel=wechat PayScene=wechat-mini时，内部会直接以小程序方式调用微信统一下单接口。
	PayScene *string `json:"PayScene,omitempty" name:"PayScene"`

	// 语言代码
	// 取值请参考[ISO 639-1代码表](https://zh.wikipedia.org/zh-cn/ISO_639-1%E4%BB%A3%E7%A0%81%E8%A1%A8)
	LocaleCode *string `json:"LocaleCode,omitempty" name:"LocaleCode"`

	// 地区代码
	// 取值请参考[ISO 3166-1二位字母代码表](https://zh.wikipedia.org/zh-cn/ISO_3166-1%E4%BA%8C%E4%BD%8D%E5%AD%97%E6%AF%8D%E4%BB%A3%E7%A0%81#%E6%AD%A3%E5%BC%8F%E5%88%86%E9%85%8D%E4%BB%A3%E7%A0%81)
	RegionCode *string `json:"RegionCode,omitempty" name:"RegionCode"`

	// 用户IP
	// 请求用户的IP地址，特定的渠道或特定的支付方式，此字段为必填
	// wechat_ecommerce渠道 - h5支付方式，此字段必填。
	UserClientIp *string `json:"UserClientIp,omitempty" name:"UserClientIp"`

	// 渠道订单号生成模式
	// 枚举值。决定请求渠道方时的订单号的生成模式，详情请联系米大师沟通。不指定时默认为由米大师自行生成。
	ChannelOrderIdMode *string `json:"ChannelOrderIdMode,omitempty" name:"ChannelOrderIdMode"`

	// 全局支付时间信息
	GlobalPayTimeInfo *CloudGlobalPayTimeInfo `json:"GlobalPayTimeInfo,omitempty" name:"GlobalPayTimeInfo"`

	// 渠道应用ID取用方式
	// USE_APPID 使用渠道应用Id;
	// USE_SUB_APPID 使用子渠道应用Id;
	// USE_APPID_AND_SUB_APPID 既使用渠道应用Id也使用子渠道应用ID。
	ChannelAppIdPolicy *string `json:"ChannelAppIdPolicy,omitempty" name:"ChannelAppIdPolicy"`

	// 门店信息
	// 特定的渠道或特定的支付方式，此字段为必填
	// wechat_ecommerce渠道 - h5支付方式，此字段必填
	StoreInfo *CloudStoreInfo `json:"StoreInfo,omitempty" name:"StoreInfo"`

	// 客户端信息
	// 特定的渠道或特定的支付方式，此字段为必填
	// wechat_ecommerce渠道 - h5支付方式，此字段必填
	ClientInfo *CloudClientInfo `json:"ClientInfo,omitempty" name:"ClientInfo"`

	// 渠道扩展促销列表
	// 可将各个渠道的促销信息放于该列表。
	ExternalPromptGroupList []*CloudExternalPromptGroup `json:"ExternalPromptGroupList,omitempty" name:"ExternalPromptGroupList"`

	// 收单模式
	// ORDER_RECEIVE_MODE_COMMON - 普通支付
	// ORDER_RECEIVE_MODE_COMBINE - 合单支付
	// ORDER_RECEIVE_MODE_V_COMBINE - 虚拟合单支付
	// 若不传入该字段，则会根据是否传入子单来判断是 普通支付 还是 合单支付
	OrderReceiveMode *string `json:"OrderReceiveMode,omitempty" name:"OrderReceiveMode"`
}

func (r *UnifiedCloudOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnifiedCloudOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MidasAppId")
	delete(f, "UserId")
	delete(f, "OutTradeNo")
	delete(f, "CurrencyType")
	delete(f, "ProductId")
	delete(f, "ProductName")
	delete(f, "ProductDetail")
	delete(f, "OriginalAmt")
	delete(f, "TotalAmt")
	delete(f, "MidasEnvironment")
	delete(f, "SubAppId")
	delete(f, "RealChannel")
	delete(f, "Channel")
	delete(f, "Metadata")
	delete(f, "Quantity")
	delete(f, "CallbackUrl")
	delete(f, "CancelUrl")
	delete(f, "WxAppId")
	delete(f, "WxSubAppId")
	delete(f, "WxOpenId")
	delete(f, "WxSubOpenId")
	delete(f, "TotalPlatformIncome")
	delete(f, "TotalMchIncome")
	delete(f, "SubOrderList")
	delete(f, "SettleInfo")
	delete(f, "AttachmentInfoList")
	delete(f, "PaymentNotifyUrl")
	delete(f, "PayScene")
	delete(f, "LocaleCode")
	delete(f, "RegionCode")
	delete(f, "UserClientIp")
	delete(f, "ChannelOrderIdMode")
	delete(f, "GlobalPayTimeInfo")
	delete(f, "ChannelAppIdPolicy")
	delete(f, "StoreInfo")
	delete(f, "ClientInfo")
	delete(f, "ExternalPromptGroupList")
	delete(f, "OrderReceiveMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnifiedCloudOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnifiedCloudOrderResponseParams struct {
	// 米大师的交易订单号。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransactionId *string `json:"TransactionId,omitempty" name:"TransactionId"`

	// 开发者的支付订单号。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutTradeNo *string `json:"OutTradeNo,omitempty" name:"OutTradeNo"`

	// SDK的支付参数。
	// 支付参数透传给米大师SDK（原文透传给SDK即可，不需要解码）
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayInfo *string `json:"PayInfo,omitempty" name:"PayInfo"`

	// 支付金额，单位：分。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalAmt *int64 `json:"TotalAmt,omitempty" name:"TotalAmt"`

	// 渠道信息，用于拉起渠道支付。j
	// son字符串，注意此字段仅会在传入正确的PayScene入参时才会有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelInfo *string `json:"ChannelInfo,omitempty" name:"ChannelInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnifiedCloudOrderResponse struct {
	*tchttp.BaseResponse
	Response *UnifiedCloudOrderResponseParams `json:"Response"`
}

func (r *UnifiedCloudOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnifiedCloudOrderResponse) FromJsonString(s string) error {
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

// Predefined struct for user
type UnifiedOrderRequestParams struct {
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

// Predefined struct for user
type UnifiedOrderResponseParams struct {
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
}

type UnifiedOrderResponse struct {
	*tchttp.BaseResponse
	Response *UnifiedOrderResponseParams `json:"Response"`
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

// Predefined struct for user
type UnifiedTlinxOrderRequestParams struct {
	// 使用门店OpenId
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 使用门店OpenKey
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 开发者流水号
	DeveloperNo *string `json:"DeveloperNo,omitempty" name:"DeveloperNo"`

	// 支付标签
	PayTag *string `json:"PayTag,omitempty" name:"PayTag"`

	// 实际交易金额（以分为单位，没有小数点）
	TradeAmount *string `json:"TradeAmount,omitempty" name:"TradeAmount"`

	// 交易结果异步通知url地址
	NotifyUrl *string `json:"NotifyUrl,omitempty" name:"NotifyUrl"`

	// 付款方式名称(当PayTag为Diy时，PayName不能为空)
	PayName *string `json:"PayName,omitempty" name:"PayName"`

	// 公众号支付时，支付成功后跳转url地址
	JumpUrl *string `json:"JumpUrl,omitempty" name:"JumpUrl"`

	// 订单名称（描述）
	OrderName *string `json:"OrderName,omitempty" name:"OrderName"`

	// 原始交易金额（以分为单位，没有小数点）
	OriginalAmount *string `json:"OriginalAmount,omitempty" name:"OriginalAmount"`

	// 抹零金额（以分为单位，没有小数点）
	IgnoreAmount *string `json:"IgnoreAmount,omitempty" name:"IgnoreAmount"`

	// 折扣金额（以分为单位，没有小数点）
	DiscountAmount *string `json:"DiscountAmount,omitempty" name:"DiscountAmount"`

	// 交易帐号（银行卡号）
	TradeAccount *string `json:"TradeAccount,omitempty" name:"TradeAccount"`

	// 交易号（收单机构交易号）
	TradeNo *string `json:"TradeNo,omitempty" name:"TradeNo"`

	// 条码支付的授权码（条码抢扫手机扫到的一串数字）
	AuthCode *string `json:"AuthCode,omitempty" name:"AuthCode"`

	// 订单标记，订单附加数据。
	Tag *string `json:"Tag,omitempty" name:"Tag"`

	// 订单备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 收单机构原始交易报文，请转换为json
	TradeResult *string `json:"TradeResult,omitempty" name:"TradeResult"`

	// 0-不分账，1-需分账。为1时标记为待分账订单，待分账订单不会进行清算。不传默认为不分账。
	Royalty *string `json:"Royalty,omitempty" name:"Royalty"`

	// 小程序支付参数：填默认值 1
	Jsapi *string `json:"Jsapi,omitempty" name:"Jsapi"`

	// 小程序支付参数：
	// 当前调起支付的小程序APPID
	SubAppId *string `json:"SubAppId,omitempty" name:"SubAppId"`

	// 小程序支付参数:
	// 用户在子商户appid下的唯一标识。
	SubOpenId *string `json:"SubOpenId,omitempty" name:"SubOpenId"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

type UnifiedTlinxOrderRequest struct {
	*tchttp.BaseRequest
	
	// 使用门店OpenId
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 使用门店OpenKey
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 开发者流水号
	DeveloperNo *string `json:"DeveloperNo,omitempty" name:"DeveloperNo"`

	// 支付标签
	PayTag *string `json:"PayTag,omitempty" name:"PayTag"`

	// 实际交易金额（以分为单位，没有小数点）
	TradeAmount *string `json:"TradeAmount,omitempty" name:"TradeAmount"`

	// 交易结果异步通知url地址
	NotifyUrl *string `json:"NotifyUrl,omitempty" name:"NotifyUrl"`

	// 付款方式名称(当PayTag为Diy时，PayName不能为空)
	PayName *string `json:"PayName,omitempty" name:"PayName"`

	// 公众号支付时，支付成功后跳转url地址
	JumpUrl *string `json:"JumpUrl,omitempty" name:"JumpUrl"`

	// 订单名称（描述）
	OrderName *string `json:"OrderName,omitempty" name:"OrderName"`

	// 原始交易金额（以分为单位，没有小数点）
	OriginalAmount *string `json:"OriginalAmount,omitempty" name:"OriginalAmount"`

	// 抹零金额（以分为单位，没有小数点）
	IgnoreAmount *string `json:"IgnoreAmount,omitempty" name:"IgnoreAmount"`

	// 折扣金额（以分为单位，没有小数点）
	DiscountAmount *string `json:"DiscountAmount,omitempty" name:"DiscountAmount"`

	// 交易帐号（银行卡号）
	TradeAccount *string `json:"TradeAccount,omitempty" name:"TradeAccount"`

	// 交易号（收单机构交易号）
	TradeNo *string `json:"TradeNo,omitempty" name:"TradeNo"`

	// 条码支付的授权码（条码抢扫手机扫到的一串数字）
	AuthCode *string `json:"AuthCode,omitempty" name:"AuthCode"`

	// 订单标记，订单附加数据。
	Tag *string `json:"Tag,omitempty" name:"Tag"`

	// 订单备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 收单机构原始交易报文，请转换为json
	TradeResult *string `json:"TradeResult,omitempty" name:"TradeResult"`

	// 0-不分账，1-需分账。为1时标记为待分账订单，待分账订单不会进行清算。不传默认为不分账。
	Royalty *string `json:"Royalty,omitempty" name:"Royalty"`

	// 小程序支付参数：填默认值 1
	Jsapi *string `json:"Jsapi,omitempty" name:"Jsapi"`

	// 小程序支付参数：
	// 当前调起支付的小程序APPID
	SubAppId *string `json:"SubAppId,omitempty" name:"SubAppId"`

	// 小程序支付参数:
	// 用户在子商户appid下的唯一标识。
	SubOpenId *string `json:"SubOpenId,omitempty" name:"SubOpenId"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *UnifiedTlinxOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnifiedTlinxOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OpenId")
	delete(f, "OpenKey")
	delete(f, "DeveloperNo")
	delete(f, "PayTag")
	delete(f, "TradeAmount")
	delete(f, "NotifyUrl")
	delete(f, "PayName")
	delete(f, "JumpUrl")
	delete(f, "OrderName")
	delete(f, "OriginalAmount")
	delete(f, "IgnoreAmount")
	delete(f, "DiscountAmount")
	delete(f, "TradeAccount")
	delete(f, "TradeNo")
	delete(f, "AuthCode")
	delete(f, "Tag")
	delete(f, "Remark")
	delete(f, "TradeResult")
	delete(f, "Royalty")
	delete(f, "Jsapi")
	delete(f, "SubAppId")
	delete(f, "SubOpenId")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnifiedTlinxOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnifiedTlinxOrderResponseParams struct {
	// 业务系统返回消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 业务系统返回码，0表示成功，其他表示失败。
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 统一下单响应对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *PayOrderResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnifiedTlinxOrderResponse struct {
	*tchttp.BaseResponse
	Response *UnifiedTlinxOrderResponseParams `json:"Response"`
}

func (r *UnifiedTlinxOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnifiedTlinxOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadExternalAnchorInfoRequestParams struct {
	// 主播Id
	AnchorId *string `json:"AnchorId,omitempty" name:"AnchorId"`

	// 身份证正面图片下载链接
	IdCardFront *string `json:"IdCardFront,omitempty" name:"IdCardFront"`

	// 身份证反面图片下载链接
	IdCardReverse *string `json:"IdCardReverse,omitempty" name:"IdCardReverse"`
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

// Predefined struct for user
type UploadExternalAnchorInfoResponseParams struct {
	// 错误码。响应成功："SUCCESS"，其他为不成功。
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 响应消息。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 该字段为null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UploadExternalAnchorInfoResponse struct {
	*tchttp.BaseResponse
	Response *UploadExternalAnchorInfoResponseParams `json:"Response"`
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

// Predefined struct for user
type UploadFileRequestParams struct {
	// 文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 文件类型
	// __IdCard__:身份证
	// __IdCardCheck__:身份证加验证(只支持人像面)
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 文件链接
	// __FileUrl和FileContent二选一__
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// 文件内容，Base64编码
	// __FileUrl和FileContent二选一__
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`

	// 文件扩展信息
	FileExtendInfo []*AnchorExtendInfo `json:"FileExtendInfo,omitempty" name:"FileExtendInfo"`
}

type UploadFileRequest struct {
	*tchttp.BaseRequest
	
	// 文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 文件类型
	// __IdCard__:身份证
	// __IdCardCheck__:身份证加验证(只支持人像面)
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 文件链接
	// __FileUrl和FileContent二选一__
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// 文件内容，Base64编码
	// __FileUrl和FileContent二选一__
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`

	// 文件扩展信息
	FileExtendInfo []*AnchorExtendInfo `json:"FileExtendInfo,omitempty" name:"FileExtendInfo"`
}

func (r *UploadFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileName")
	delete(f, "FileType")
	delete(f, "FileUrl")
	delete(f, "FileContent")
	delete(f, "FileExtendInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadFileResponseParams struct {
	// 文件ID
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UploadFileResponse struct {
	*tchttp.BaseResponse
	Response *UploadFileResponseParams `json:"Response"`
}

func (r *UploadFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UploadFileResult struct {
	// 存储区域（0私密区，1公共区），请严格按文件要求，上传到不同的区域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Storage *string `json:"Storage,omitempty" name:"Storage"`

	// 文件路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`
}

// Predefined struct for user
type UploadOpenBankSubMerchantCredentialRequestParams struct {
	// 渠道商户ID。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 渠道子商户ID。
	ChannelSubMerchantId *string `json:"ChannelSubMerchantId,omitempty" name:"ChannelSubMerchantId"`

	// 渠道名称。详见附录-枚举类型-ChannelName。
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 外部序列进件号。
	OutApplyId *string `json:"OutApplyId,omitempty" name:"OutApplyId"`

	// 资质类型，详见附录-枚举类型-CredentialType。
	CredentialType *string `json:"CredentialType,omitempty" name:"CredentialType"`

	// 文件类型。
	// 合利宝渠道，文件类型为PNG/JPG格式。
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 支付方式。
	// 合利宝渠道不需要传。
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`

	// 资质文件内容。Base64编码，资质文件内容和链接二选一。
	// 合利宝渠道，文件限制大小5M以内。
	CredentialContent *string `json:"CredentialContent,omitempty" name:"CredentialContent"`

	// 资质文件链接。资质文件内容和链接二选一。
	// 合利宝渠道，文件限制大小5M以内。
	CredentialUrl *string `json:"CredentialUrl,omitempty" name:"CredentialUrl"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

type UploadOpenBankSubMerchantCredentialRequest struct {
	*tchttp.BaseRequest
	
	// 渠道商户ID。
	ChannelMerchantId *string `json:"ChannelMerchantId,omitempty" name:"ChannelMerchantId"`

	// 渠道子商户ID。
	ChannelSubMerchantId *string `json:"ChannelSubMerchantId,omitempty" name:"ChannelSubMerchantId"`

	// 渠道名称。详见附录-枚举类型-ChannelName。
	ChannelName *string `json:"ChannelName,omitempty" name:"ChannelName"`

	// 外部序列进件号。
	OutApplyId *string `json:"OutApplyId,omitempty" name:"OutApplyId"`

	// 资质类型，详见附录-枚举类型-CredentialType。
	CredentialType *string `json:"CredentialType,omitempty" name:"CredentialType"`

	// 文件类型。
	// 合利宝渠道，文件类型为PNG/JPG格式。
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 支付方式。
	// 合利宝渠道不需要传。
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`

	// 资质文件内容。Base64编码，资质文件内容和链接二选一。
	// 合利宝渠道，文件限制大小5M以内。
	CredentialContent *string `json:"CredentialContent,omitempty" name:"CredentialContent"`

	// 资质文件链接。资质文件内容和链接二选一。
	// 合利宝渠道，文件限制大小5M以内。
	CredentialUrl *string `json:"CredentialUrl,omitempty" name:"CredentialUrl"`

	// 环境类型。
	// __release__:生产环境
	// __sandbox__:沙箱环境
	// _不填默认为生产环境_
	Environment *string `json:"Environment,omitempty" name:"Environment"`
}

func (r *UploadOpenBankSubMerchantCredentialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadOpenBankSubMerchantCredentialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChannelMerchantId")
	delete(f, "ChannelSubMerchantId")
	delete(f, "ChannelName")
	delete(f, "OutApplyId")
	delete(f, "CredentialType")
	delete(f, "FileType")
	delete(f, "PaymentMethod")
	delete(f, "CredentialContent")
	delete(f, "CredentialUrl")
	delete(f, "Environment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadOpenBankSubMerchantCredentialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadOpenBankSubMerchantCredentialResponseParams struct {
	// 错误码。
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误消息。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 返回结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *UploadOpenBankSubMerchantCredentialResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UploadOpenBankSubMerchantCredentialResponse struct {
	*tchttp.BaseResponse
	Response *UploadOpenBankSubMerchantCredentialResponseParams `json:"Response"`
}

func (r *UploadOpenBankSubMerchantCredentialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadOpenBankSubMerchantCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UploadOpenBankSubMerchantCredentialResult struct {
	// 上传状态
	// SUCCESS：上传成功
	// FAILED：上传失败
	// PROCESSING:上传中
	UploadStatus *string `json:"UploadStatus,omitempty" name:"UploadStatus"`

	// 上传描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	UploadMessage *string `json:"UploadMessage,omitempty" name:"UploadMessage"`

	// 渠道上传流水号
	ChannelApplyId *string `json:"ChannelApplyId,omitempty" name:"ChannelApplyId"`
}

// Predefined struct for user
type UploadOrgFileRequestParams struct {
	// 收单系统分配的开放ID
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 收单系统分配的密钥
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 存储区域（0私密区，1公共区），请严格按文件要求，上传到不同的区域
	Storage *string `json:"Storage,omitempty" name:"Storage"`

	// 文件的md5值（为防止平台多次上传重复文件，文件名为文件md5,且不会覆盖，重复上传返回第一次上传成功的文件路径）
	FileMd5 *string `json:"FileMd5,omitempty" name:"FileMd5"`

	// 文件内容（先将图片转换成二进制，再进行base64加密）
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`

	// 文件扩展名（png,jpg,gif）
	FileExtension *string `json:"FileExtension,omitempty" name:"FileExtension"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

type UploadOrgFileRequest struct {
	*tchttp.BaseRequest
	
	// 收单系统分配的开放ID
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 收单系统分配的密钥
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 存储区域（0私密区，1公共区），请严格按文件要求，上传到不同的区域
	Storage *string `json:"Storage,omitempty" name:"Storage"`

	// 文件的md5值（为防止平台多次上传重复文件，文件名为文件md5,且不会覆盖，重复上传返回第一次上传成功的文件路径）
	FileMd5 *string `json:"FileMd5,omitempty" name:"FileMd5"`

	// 文件内容（先将图片转换成二进制，再进行base64加密）
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`

	// 文件扩展名（png,jpg,gif）
	FileExtension *string `json:"FileExtension,omitempty" name:"FileExtension"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *UploadOrgFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadOrgFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OpenId")
	delete(f, "OpenKey")
	delete(f, "Storage")
	delete(f, "FileMd5")
	delete(f, "FileContent")
	delete(f, "FileExtension")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadOrgFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadOrgFileResponseParams struct {
	// 业务系统返回消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 业务系统返回码
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 上传机构文件响应对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *UploadFileResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UploadOrgFileResponse struct {
	*tchttp.BaseResponse
	Response *UploadOrgFileResponseParams `json:"Response"`
}

func (r *UploadOrgFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadOrgFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadTaxListRequestParams struct {
	// 平台渠道
	Channel *int64 `json:"Channel,omitempty" name:"Channel"`

	// 起始月份，YYYY-MM
	BeginMonth *string `json:"BeginMonth,omitempty" name:"BeginMonth"`

	// 结束月份。如果只上传一个月，结束月份等于起始月份
	EndMonth *string `json:"EndMonth,omitempty" name:"EndMonth"`

	// 完税列表下载地址
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`
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

// Predefined struct for user
type UploadTaxListResponseParams struct {
	// 完税ID
	TaxId *string `json:"TaxId,omitempty" name:"TaxId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UploadTaxListResponse struct {
	*tchttp.BaseResponse
	Response *UploadTaxListResponseParams `json:"Response"`
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

// Predefined struct for user
type UploadTaxPaymentRequestParams struct {
	// 平台渠道
	Channel *int64 `json:"Channel,omitempty" name:"Channel"`

	// 完税ID
	TaxId *string `json:"TaxId,omitempty" name:"TaxId"`

	// 完税列表下载地址
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`
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

// Predefined struct for user
type UploadTaxPaymentResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UploadTaxPaymentResponse struct {
	*tchttp.BaseResponse
	Response *UploadTaxPaymentResponseParams `json:"Response"`
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

// Predefined struct for user
type ViewContractRequestParams struct {
	// 收单系统分配的开放ID
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 收单系统分配的密钥
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 外部合同主键编号（ContractId或OutContractId必须传一个）
	OutContractId *string `json:"OutContractId,omitempty" name:"OutContractId"`

	// 合同主键（ContractId或OutContractId必须传一个）
	ContractId *string `json:"ContractId,omitempty" name:"ContractId"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

type ViewContractRequest struct {
	*tchttp.BaseRequest
	
	// 收单系统分配的开放ID
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 收单系统分配的密钥
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 外部合同主键编号（ContractId或OutContractId必须传一个）
	OutContractId *string `json:"OutContractId,omitempty" name:"OutContractId"`

	// 合同主键（ContractId或OutContractId必须传一个）
	ContractId *string `json:"ContractId,omitempty" name:"ContractId"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *ViewContractRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ViewContractRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OpenId")
	delete(f, "OpenKey")
	delete(f, "OutContractId")
	delete(f, "ContractId")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ViewContractRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ViewContractResponseParams struct {
	// 业务系统返回消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 业务系统返回码
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 合同明细响应对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *ViewContractResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ViewContractResponse struct {
	*tchttp.BaseResponse
	Response *ViewContractResponseParams `json:"Response"`
}

func (r *ViewContractResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ViewContractResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ViewContractResult struct {
	// 支付标签（唯一性）
	// 注意：此字段可能返回 null，表示取不到有效值。
	PaymentTag *string `json:"PaymentTag,omitempty" name:"PaymentTag"`

	// 城市
	// 注意：此字段可能返回 null，表示取不到有效值。
	City *string `json:"City,omitempty" name:"City"`

	// 机构编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentNo *string `json:"AgentNo,omitempty" name:"AgentNo"`

	// 合同选项值4
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContractOptionFour *string `json:"ContractOptionFour,omitempty" name:"ContractOptionFour"`

	// 合同选项值2
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContractOptionTwo *string `json:"ContractOptionTwo,omitempty" name:"ContractOptionTwo"`

	// 合同状态（0未审核，1已审核，2审核未通过，3待审核，4已删除，5初审通过）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 支付方式编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	PaymentId *string `json:"PaymentId,omitempty" name:"PaymentId"`

	// 商户签约扣率
	// 注意：此字段可能返回 null，表示取不到有效值。
	Fee *string `json:"Fee,omitempty" name:"Fee"`

	// 合同选项名称5
	// 注意：此字段可能返回 null，表示取不到有效值。
	PaymentOptionFive *string `json:"PaymentOptionFive,omitempty" name:"PaymentOptionFive"`

	// 机构合同主键
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutContractId *string `json:"OutContractId,omitempty" name:"OutContractId"`

	// 不同的支付方式对于进件有不同的个性化需求，支付方式字段都是以双下划写开头的字段名称，请以支付方式规定的格式传值
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelExtJson *string `json:"ChannelExtJson,omitempty" name:"ChannelExtJson"`

	// 合同选项值5
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContractOptionFive *string `json:"ContractOptionFive,omitempty" name:"ContractOptionFive"`

	// 省份
	// 注意：此字段可能返回 null，表示取不到有效值。
	Province *string `json:"Province,omitempty" name:"Province"`

	// 生效日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 详细地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Address *string `json:"Address,omitempty" name:"Address"`

	// 过期日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 合同选项值6
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContractOptionSix *string `json:"ContractOptionSix,omitempty" name:"ContractOptionSix"`

	// 合同选项名称7
	// 注意：此字段可能返回 null，表示取不到有效值。
	PaymentOptionSeven *string `json:"PaymentOptionSeven,omitempty" name:"PaymentOptionSeven"`

	// 合同照片补充【私密区】
	// 注意：此字段可能返回 null，表示取不到有效值。
	PictureTwo *string `json:"PictureTwo,omitempty" name:"PictureTwo"`

	// 商户编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	MerchantNo *string `json:"MerchantNo,omitempty" name:"MerchantNo"`

	// 机构名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentName *string `json:"AgentName,omitempty" name:"AgentName"`

	// 合同选项值8
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContractOptionOther *string `json:"ContractOptionOther,omitempty" name:"ContractOptionOther"`

	// 合同选项值3
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContractOptionThree *string `json:"ContractOptionThree,omitempty" name:"ContractOptionThree"`

	// 县/区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Country *string `json:"Country,omitempty" name:"Country"`

	// 合同关联的门店数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShopCount *string `json:"ShopCount,omitempty" name:"ShopCount"`

	// 合同选项名称3
	// 注意：此字段可能返回 null，表示取不到有效值。
	PaymentOptionThree *string `json:"PaymentOptionThree,omitempty" name:"PaymentOptionThree"`

	// 支付方式行业名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PaymentClassificationName *string `json:"PaymentClassificationName,omitempty" name:"PaymentClassificationName"`

	// 合同选项值7
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContractOptionSeven *string `json:"ContractOptionSeven,omitempty" name:"ContractOptionSeven"`

	// 合同选项名称4
	// 注意：此字段可能返回 null，表示取不到有效值。
	PaymentOptionFour *string `json:"PaymentOptionFour,omitempty" name:"PaymentOptionFour"`

	// 商户签约扣率封顶值（分为单位）
	// 注意：此字段可能返回 null，表示取不到有效值。
	PaymentClassificationLimit *string `json:"PaymentClassificationLimit,omitempty" name:"PaymentClassificationLimit"`

	// 审核备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 合同选项名称6
	// 注意：此字段可能返回 null，表示取不到有效值。
	PaymentOptionSix *string `json:"PaymentOptionSix,omitempty" name:"PaymentOptionSix"`

	// 品牌名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	MerchantName *string `json:"MerchantName,omitempty" name:"MerchantName"`

	// 合同选项值1
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContractOptionOne *string `json:"ContractOptionOne,omitempty" name:"ContractOptionOne"`

	// 合同选项名称8
	// 注意：此字段可能返回 null，表示取不到有效值。
	PaymentOptionOther *string `json:"PaymentOptionOther,omitempty" name:"PaymentOptionOther"`

	// 合同选项名称2
	// 注意：此字段可能返回 null，表示取不到有效值。
	PaymentOptionTwo *string `json:"PaymentOptionTwo,omitempty" name:"PaymentOptionTwo"`

	// 合同选项名称1
	// 注意：此字段可能返回 null，表示取不到有效值。
	PaymentOptionOne *string `json:"PaymentOptionOne,omitempty" name:"PaymentOptionOne"`

	// 更新时间（yyyy-mm-dd hh:ii:ss）
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 联系人电话
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContactTelephone *string `json:"ContactTelephone,omitempty" name:"ContactTelephone"`

	// 联系人
	// 注意：此字段可能返回 null，表示取不到有效值。
	Contact *string `json:"Contact,omitempty" name:"Contact"`

	// 签约日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignDate *string `json:"SignDate,omitempty" name:"SignDate"`

	// 合同选项名称9
	// 注意：此字段可能返回 null，表示取不到有效值。
	PaymentOptionNine *string `json:"PaymentOptionNine,omitempty" name:"PaymentOptionNine"`

	// 付款方式名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PaymentName *string `json:"PaymentName,omitempty" name:"PaymentName"`

	// 付款方式名称（内部名称）
	// 注意：此字段可能返回 null，表示取不到有效值。
	PaymentInternalName *string `json:"PaymentInternalName,omitempty" name:"PaymentInternalName"`

	// 合同选项值10
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContractOptionTen *string `json:"ContractOptionTen,omitempty" name:"ContractOptionTen"`

	// 合同编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *string `json:"Code,omitempty" name:"Code"`

	// 合同照片【私密区】
	// 注意：此字段可能返回 null，表示取不到有效值。
	PictureOne *string `json:"PictureOne,omitempty" name:"PictureOne"`

	// 签约人
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignMan *string `json:"SignMan,omitempty" name:"SignMan"`

	// 渠道号
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelNo *string `json:"ChannelNo,omitempty" name:"ChannelNo"`

	// 添加时间（yyyy-mm-dd hh:ii:ss）
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

	// 是否自动续签（1是，0否）
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoSign *string `json:"AutoSign,omitempty" name:"AutoSign"`

	// 合同选项值9
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContractOptionNine *string `json:"ContractOptionNine,omitempty" name:"ContractOptionNine"`

	// 城市编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	CityId *string `json:"CityId,omitempty" name:"CityId"`

	// 交易类型（多个以小写逗号分开，0现金，1刷卡，2主扫，3被扫，4JSPAY，5预授权）
	// 注意：此字段可能返回 null，表示取不到有效值。
	PaymentType *string `json:"PaymentType,omitempty" name:"PaymentType"`

	// 支付方式行业编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	PaymentClassificationId *string `json:"PaymentClassificationId,omitempty" name:"PaymentClassificationId"`

	// 品牌名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	BrandName *string `json:"BrandName,omitempty" name:"BrandName"`

	// 合同选项名称10
	// 注意：此字段可能返回 null，表示取不到有效值。
	PaymentOptionTen *string `json:"PaymentOptionTen,omitempty" name:"PaymentOptionTen"`

	// 合同主键
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContractId *string `json:"ContractId,omitempty" name:"ContractId"`
}

// Predefined struct for user
type ViewMerchantRequestParams struct {
	// 收单系统分配的开放ID
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 收单系统分配的密钥
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 外部商户主键编号（MerchantNo或OutMerchantId必须传一个）
	OutMerchantId *string `json:"OutMerchantId,omitempty" name:"OutMerchantId"`

	// 商户编号（MerchantNo或OutMerchantId必须传一个）
	MerchantNo *string `json:"MerchantNo,omitempty" name:"MerchantNo"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

type ViewMerchantRequest struct {
	*tchttp.BaseRequest
	
	// 收单系统分配的开放ID
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 收单系统分配的密钥
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 外部商户主键编号（MerchantNo或OutMerchantId必须传一个）
	OutMerchantId *string `json:"OutMerchantId,omitempty" name:"OutMerchantId"`

	// 商户编号（MerchantNo或OutMerchantId必须传一个）
	MerchantNo *string `json:"MerchantNo,omitempty" name:"MerchantNo"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *ViewMerchantRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ViewMerchantRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OpenId")
	delete(f, "OpenKey")
	delete(f, "OutMerchantId")
	delete(f, "MerchantNo")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ViewMerchantRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ViewMerchantResponseParams struct {
	// 业务系统返回消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 业务系统返回码
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 商户明细响应对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *ViewMerchantResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ViewMerchantResponse struct {
	*tchttp.BaseResponse
	Response *ViewMerchantResponseParams `json:"Response"`
}

func (r *ViewMerchantResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ViewMerchantResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ViewMerchantResult struct {
	// 城市
	// 注意：此字段可能返回 null，表示取不到有效值。
	City *string `json:"City,omitempty" name:"City"`

	// 税务登记证图片【私密区】
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaxCollectionPicture *string `json:"TaxCollectionPicture,omitempty" name:"TaxCollectionPicture"`

	// 法人证件号码
	// 注意：此字段可能返回 null，表示取不到有效值。
	BossIdNo *string `json:"BossIdNo,omitempty" name:"BossIdNo"`

	// 法人亲属证件号码
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountIdNo *string `json:"AccountIdNo,omitempty" name:"AccountIdNo"`

	// 其他资料3
	// 注意：此字段可能返回 null，表示取不到有效值。
	OtherPictureThree *string `json:"OtherPictureThree,omitempty" name:"OtherPictureThree"`

	// 法人亲属证件类型（1居民身份证,2临时居民身份证,3居民户口簿,4护照,5港澳居民来往内地通行证,6回乡证,7军人证,8武警身份证,9其他法定文件）结算账户人身份为法人亲属时必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountIdType *string `json:"AccountIdType,omitempty" name:"AccountIdType"`

	// 商户状态（0未审核，1已审核，2审核未通过，3待审核，4已删除，5初审通过）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 营业执照图片【私密区】（系统返回的图片路径）
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessLicensePicture *string `json:"BusinessLicensePicture,omitempty" name:"BusinessLicensePicture"`

	// 品牌名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	BrandName *string `json:"BrandName,omitempty" name:"BrandName"`

	// 法人身份证正面【私密区】（系统返回的图片路径）
	// 注意：此字段可能返回 null，表示取不到有效值。
	BossPositive *string `json:"BossPositive,omitempty" name:"BossPositive"`

	// 开通应用数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppCount *string `json:"AppCount,omitempty" name:"AppCount"`

	// 法人身份证背面【私密区】（系统返回的图片路径）
	// 注意：此字段可能返回 null，表示取不到有效值。
	BossBack *string `json:"BossBack,omitempty" name:"BossBack"`

	// 组织机构代码证图片【私密区】
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationCodePicture *string `json:"OrganizationCodePicture,omitempty" name:"OrganizationCodePicture"`

	// 营业执照过期时间（yyyy-mm-dd）
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessLicenseEndDate *string `json:"BusinessLicenseEndDate,omitempty" name:"BusinessLicenseEndDate"`

	// 组织机构代码证号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationCodeNo *string `json:"OrganizationCodeNo,omitempty" name:"OrganizationCodeNo"`

	// 机构编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentNo *string `json:"AgentNo,omitempty" name:"AgentNo"`

	// 省份
	// 注意：此字段可能返回 null，表示取不到有效值。
	Province *string `json:"Province,omitempty" name:"Province"`

	// 法人证件生效时间（yyyy-mm-dd）
	// 注意：此字段可能返回 null，表示取不到有效值。
	BossStartDate *string `json:"BossStartDate,omitempty" name:"BossStartDate"`

	// 更新时间（yyyy-mm-dd hh:ii:ss）
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 清算联行号，开户行行号
	// 注意：此字段可能返回 null，表示取不到有效值。
	BankNo *string `json:"BankNo,omitempty" name:"BankNo"`

	// 税务登记证生效时间（yyyy-mm-dd）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaxCollectionStartDate *string `json:"TaxCollectionStartDate,omitempty" name:"TaxCollectionStartDate"`

	// 营业执照生效时间（yyyy-mm-dd）
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessLicenseStartDate *string `json:"BusinessLicenseStartDate,omitempty" name:"BusinessLicenseStartDate"`

	// 客户经理用户编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountManagerId *string `json:"AccountManagerId,omitempty" name:"AccountManagerId"`

	// 分类编号(多个以小写逗号分开)
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClassificationIds *string `json:"ClassificationIds,omitempty" name:"ClassificationIds"`

	// 营业执照类型（1三证合一，2非三证合一）
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessLicenseType *string `json:"BusinessLicenseType,omitempty" name:"BusinessLicenseType"`

	// 法人证件过期时间（yyyy-mm-dd）
	// 注意：此字段可能返回 null，表示取不到有效值。
	BossEndDate *string `json:"BossEndDate,omitempty" name:"BossEndDate"`

	// 营业执照编号（系统有唯一性校验）
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessLicenseNo *string `json:"BusinessLicenseNo,omitempty" name:"BusinessLicenseNo"`

	// 机构名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentName *string `json:"AgentName,omitempty" name:"AgentName"`

	// 商户简介
	// 注意：此字段可能返回 null，表示取不到有效值。
	Intro *string `json:"Intro,omitempty" name:"Intro"`

	// 法人证件类型（1居民身份证,2临时居民身份证,3居民户口簿,4护照,5港澳居民来往内地通行证,6回乡证,7军人证,8武警身份证,9其他法定文件）
	// 注意：此字段可能返回 null，表示取不到有效值。
	BossIdType *string `json:"BossIdType,omitempty" name:"BossIdType"`

	// 添加时间（yyyy-mm-dd hh:ii:ss）
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

	// 门店数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShopCount *string `json:"ShopCount,omitempty" name:"ShopCount"`

	// 结算账户人身份（1法人，2法人亲属），结算帐户为对私时必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountBoss *string `json:"AccountBoss,omitempty" name:"AccountBoss"`

	// 分类名称(多个以小写逗号分开)
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClassificationNames *string `json:"ClassificationNames,omitempty" name:"ClassificationNames"`

	// 法人电话
	// 注意：此字段可能返回 null，表示取不到有效值。
	BossTelephone *string `json:"BossTelephone,omitempty" name:"BossTelephone"`

	// 客户经理姓名，必须为系统后台的管理员真实姓名
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountManagerName *string `json:"AccountManagerName,omitempty" name:"AccountManagerName"`

	// 终端数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TerminalCount *string `json:"TerminalCount,omitempty" name:"TerminalCount"`

	// 审核备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 财务联系人
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinancialContact *string `json:"FinancialContact,omitempty" name:"FinancialContact"`

	// 商户名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	MerchantName *string `json:"MerchantName,omitempty" name:"MerchantName"`

	// 法人性别（1男，2女）
	// 注意：此字段可能返回 null，表示取不到有效值。
	BossSex *string `json:"BossSex,omitempty" name:"BossSex"`

	// 商户编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	MerchantNo *string `json:"MerchantNo,omitempty" name:"MerchantNo"`

	// 法人住址
	// 注意：此字段可能返回 null，表示取不到有效值。
	BossAddress *string `json:"BossAddress,omitempty" name:"BossAddress"`

	// 县/区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Country *string `json:"Country,omitempty" name:"Country"`

	// 详细地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Address *string `json:"Address,omitempty" name:"Address"`

	// 法人职业
	// 注意：此字段可能返回 null，表示取不到有效值。
	BossJob *string `json:"BossJob,omitempty" name:"BossJob"`

	// 许可证图片【私密区】
	// 注意：此字段可能返回 null，表示取不到有效值。
	LicencePicture *string `json:"LicencePicture,omitempty" name:"LicencePicture"`

	// 组织机构代码证过期时间（yyyy-mm-dd）
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationCodeEndDate *string `json:"OrganizationCodeEndDate,omitempty" name:"OrganizationCodeEndDate"`

	// 营业时间，多个以小写逗号分开(9:00-12:00,13:00-18:00)
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpenHours *string `json:"OpenHours,omitempty" name:"OpenHours"`

	// 其他资料2
	// 注意：此字段可能返回 null，表示取不到有效值。
	OtherPictureTwo *string `json:"OtherPictureTwo,omitempty" name:"OtherPictureTwo"`

	// 其他资料1
	// 注意：此字段可能返回 null，表示取不到有效值。
	OtherPictureOne *string `json:"OtherPictureOne,omitempty" name:"OtherPictureOne"`

	// 银行户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// 合同数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContractCount *string `json:"ContractCount,omitempty" name:"ContractCount"`

	// 授权文件【私密区】（当结算帐户身份为法人亲属时必传）
	// 注意：此字段可能返回 null，表示取不到有效值。
	LicencePictureTwo *string `json:"LicencePictureTwo,omitempty" name:"LicencePictureTwo"`

	// 银行账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountNo *string `json:"AccountNo,omitempty" name:"AccountNo"`

	// 法人邮箱
	// 注意：此字段可能返回 null，表示取不到有效值。
	BossEmail *string `json:"BossEmail,omitempty" name:"BossEmail"`

	// 结算账户类型（2对私，1对公）
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountType *string `json:"AccountType,omitempty" name:"AccountType"`

	// 税务登记证过期时间（yyyy-mm-dd）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaxCollectionEndDate *string `json:"TaxCollectionEndDate,omitempty" name:"TaxCollectionEndDate"`

	// 开户行名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	BankName *string `json:"BankName,omitempty" name:"BankName"`

	// 联系电话
	// 注意：此字段可能返回 null，表示取不到有效值。
	Telephone *string `json:"Telephone,omitempty" name:"Telephone"`

	// 外部商户主键编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutMerchantId *string `json:"OutMerchantId,omitempty" name:"OutMerchantId"`

	// 城市编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	CityId *string `json:"CityId,omitempty" name:"CityId"`

	// 法人证件国别/地区（中国CHN，香港HKG，澳门MAC，台湾CTN）
	// 注意：此字段可能返回 null，表示取不到有效值。
	BossIdCount *string `json:"BossIdCount,omitempty" name:"BossIdCount"`

	// 商户标记，自定义参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag *string `json:"Tag,omitempty" name:"Tag"`

	// 财务联系人电话
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinancialTelephone *string `json:"FinancialTelephone,omitempty" name:"FinancialTelephone"`

	// 法人姓名
	// 注意：此字段可能返回 null，表示取不到有效值。
	BossName *string `json:"BossName,omitempty" name:"BossName"`

	// 组织机构代码证生效时间（yyyy-mm-dd）
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationCodeStartDate *string `json:"OrganizationCodeStartDate,omitempty" name:"OrganizationCodeStartDate"`

	// 商户logo【公共区】
	// 注意：此字段可能返回 null，表示取不到有效值。
	Logo *string `json:"Logo,omitempty" name:"Logo"`

	// 其他资料4
	// 注意：此字段可能返回 null，表示取不到有效值。
	OtherPictureFour *string `json:"OtherPictureFour,omitempty" name:"OtherPictureFour"`

	// 税务登记证号
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaxCollectionNo *string `json:"TaxCollectionNo,omitempty" name:"TaxCollectionNo"`
}

// Predefined struct for user
type ViewShopRequestParams struct {
	// 收单系统分配的开放ID
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 收单系统分配的密钥
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 外部商户主键编号（ShopNo或OutShopId必须传一个）
	OutShopId *string `json:"OutShopId,omitempty" name:"OutShopId"`

	// 门店编号（ShopNo或OutShopId必须传一个）
	ShopNo *string `json:"ShopNo,omitempty" name:"ShopNo"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

type ViewShopRequest struct {
	*tchttp.BaseRequest
	
	// 收单系统分配的开放ID
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 收单系统分配的密钥
	OpenKey *string `json:"OpenKey,omitempty" name:"OpenKey"`

	// 外部商户主键编号（ShopNo或OutShopId必须传一个）
	OutShopId *string `json:"OutShopId,omitempty" name:"OutShopId"`

	// 门店编号（ShopNo或OutShopId必须传一个）
	ShopNo *string `json:"ShopNo,omitempty" name:"ShopNo"`

	// 沙箱环境填sandbox，正式环境不填
	Profile *string `json:"Profile,omitempty" name:"Profile"`
}

func (r *ViewShopRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ViewShopRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OpenId")
	delete(f, "OpenKey")
	delete(f, "OutShopId")
	delete(f, "ShopNo")
	delete(f, "Profile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ViewShopRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ViewShopResponseParams struct {
	// 业务系统返回消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMessage *string `json:"ErrMessage,omitempty" name:"ErrMessage"`

	// 业务系统返回码
	ErrCode *string `json:"ErrCode,omitempty" name:"ErrCode"`

	// 门店明细响应对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *ViewShopResult `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ViewShopResponse struct {
	*tchttp.BaseResponse
	Response *ViewShopResponseParams `json:"Response"`
}

func (r *ViewShopResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ViewShopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ViewShopResult struct {
	// 城市
	// 注意：此字段可能返回 null，表示取不到有效值。
	City *string `json:"City,omitempty" name:"City"`

	// 门店简称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShopName *string `json:"ShopName,omitempty" name:"ShopName"`

	// 百度地图纬度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Latitude *string `json:"Latitude,omitempty" name:"Latitude"`

	// 品牌名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	BrandName *string `json:"BrandName,omitempty" name:"BrandName"`

	// 开通应用数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppCount *string `json:"AppCount,omitempty" name:"AppCount"`

	// 负责人手机号码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContactTelephone *string `json:"ContactTelephone,omitempty" name:"ContactTelephone"`

	// 商户名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	MerchantName *string `json:"MerchantName,omitempty" name:"MerchantName"`

	// 省份
	// 注意：此字段可能返回 null，表示取不到有效值。
	Province *string `json:"Province,omitempty" name:"Province"`

	// 县/区
	// 注意：此字段可能返回 null，表示取不到有效值。
	County *string `json:"County,omitempty" name:"County"`

	// 更新时间（yyyy-mm-dd hh:ii:ss）
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 终端数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TerminalCount *string `json:"TerminalCount,omitempty" name:"TerminalCount"`

	// 收银台图片【公共区】
	// 注意：此字段可能返回 null，表示取不到有效值。
	PictureTwo *string `json:"PictureTwo,omitempty" name:"PictureTwo"`

	// 高德地图纬度
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatitudeTwo *string `json:"LatitudeTwo,omitempty" name:"LatitudeTwo"`

	// 机构名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentName *string `json:"AgentName,omitempty" name:"AgentName"`

	// 其他照片【公共区】
	// 注意：此字段可能返回 null，表示取不到有效值。
	PictureFour *string `json:"PictureFour,omitempty" name:"PictureFour"`

	// 高德地图经度
	// 注意：此字段可能返回 null，表示取不到有效值。
	LongitudeTwo *string `json:"LongitudeTwo,omitempty" name:"LongitudeTwo"`

	// 门店状态（0未审核，1已审核，2审核未通过，3待审核，4已删除，5初审通过）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 审核备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 机构编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentNo *string `json:"AgentNo,omitempty" name:"AgentNo"`

	// 商户编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	MerchantNo *string `json:"MerchantNo,omitempty" name:"MerchantNo"`

	// 添加时间（yyyy-mm-dd hh:ii:ss）
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

	// 详细地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Address *string `json:"Address,omitempty" name:"Address"`

	// 百度地图经度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Longitude *string `json:"Longitude,omitempty" name:"Longitude"`

	// 门店编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShopNo *string `json:"ShopNo,omitempty" name:"ShopNo"`

	// 门店全称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShopFullName *string `json:"ShopFullName,omitempty" name:"ShopFullName"`

	// 门店负责人
	// 注意：此字段可能返回 null，表示取不到有效值。
	Contact *string `json:"Contact,omitempty" name:"Contact"`

	// 店内环境图片【公共区】
	// 注意：此字段可能返回 null，表示取不到有效值。
	PictureThree *string `json:"PictureThree,omitempty" name:"PictureThree"`

	// 整体门面（含招牌）图片【公共区】
	// 注意：此字段可能返回 null，表示取不到有效值。
	PictureOne *string `json:"PictureOne,omitempty" name:"PictureOne"`

	// 门店电话
	// 注意：此字段可能返回 null，表示取不到有效值。
	Telephone *string `json:"Telephone,omitempty" name:"Telephone"`

	// 机构门店主键
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutShopId *string `json:"OutShopId,omitempty" name:"OutShopId"`

	// 城市编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	CityId *string `json:"CityId,omitempty" name:"CityId"`
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

// Predefined struct for user
type WithdrawCashMembershipRequestParams struct {
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

// Predefined struct for user
type WithdrawCashMembershipResponseParams struct {
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
}

type WithdrawCashMembershipResponse struct {
	*tchttp.BaseResponse
	Response *WithdrawCashMembershipResponseParams `json:"Response"`
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