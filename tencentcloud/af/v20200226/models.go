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

package v20200226

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type DescribeAntiFraudRequest struct {
	*tchttp.BaseRequest

	// 原始业务入参(二选一）
	BusinessSecurityData *FinanceAntiFraudFilter `json:"BusinessSecurityData,omitempty" name:"BusinessSecurityData"`

	// 密文业务入参(二选一）
	BusinessCryptoData *FinanceAntiFraudCryptoFilter `json:"BusinessCryptoData,omitempty" name:"BusinessCryptoData"`
}

func (r *DescribeAntiFraudRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAntiFraudRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BusinessSecurityData")
	delete(f, "BusinessCryptoData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAntiFraudRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAntiFraudResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回结果
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data *FinanceAntiFraudRecord `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAntiFraudResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAntiFraudResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FinanceAntiFraudCryptoFilter struct {

	// 值1定义：AES加密方式[加密模式ECB；填充格式pkcs7padding；秘钥16字节即128位
	CryptoType *string `json:"CryptoType,omitempty" name:"CryptoType"`

	// 业务字段BusinessSecurityData的json数据格式，采用CryptoType相应的加密方式计算后得到的bash64编码内容。比如对{"PhoneNumber":"13430420001","IdNumber":"420115199501010001","BankCardNumber":"6214000100010001"}包体做加密。
	CryptoContent *string `json:"CryptoContent,omitempty" name:"CryptoContent"`
}

type FinanceAntiFraudFilter struct {

	// 电话号码(五选二)
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// 身份证Id(五选二) 参数详细定义请加微信：TYXGJ-01
	IdNumber *string `json:"IdNumber,omitempty" name:"IdNumber"`

	// 银行卡号(五选二)
	BankCardNumber *string `json:"BankCardNumber,omitempty" name:"BankCardNumber"`

	// 用户请求来源 IP(五选二)
	UserIp *string `json:"UserIp,omitempty" name:"UserIp"`

	// 国际移动设备识别码(五选二)
	Imei *string `json:"Imei,omitempty" name:"Imei"`

	// ios 系统广告标示符(五选二)
	Idfa *string `json:"Idfa,omitempty" name:"Idfa"`

	// 业务场景ID，需要找技术对接
	Scene *string `json:"Scene,omitempty" name:"Scene"`

	// 姓名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 用户邮箱地址
	EmailAddress *string `json:"EmailAddress,omitempty" name:"EmailAddress"`

	// 用户住址
	Address *string `json:"Address,omitempty" name:"Address"`

	// MAC 地址
	Mac *string `json:"Mac,omitempty" name:"Mac"`

	// 国际移动用户识别码
	Imsi *string `json:"Imsi,omitempty" name:"Imsi"`

	// 1：关联的腾讯帐号QQ；2：开放帐号微信
	AccountType *string `json:"AccountType,omitempty" name:"AccountType"`

	// 可选的 QQ 或微信 openid
	Uid *string `json:"Uid,omitempty" name:"Uid"`

	// qq 或微信分配给网站或应用的 appid，用来唯一标识网站或应用
	AppIdU *string `json:"AppIdU,omitempty" name:"AppIdU"`

	// WIFI MAC
	WifiMac *string `json:"WifiMac,omitempty" name:"WifiMac"`

	// WIFI 服务集标识
	WifiSSID *string `json:"WifiSSID,omitempty" name:"WifiSSID"`

	// WIFI-BSSID
	WifiBSSID *string `json:"WifiBSSID,omitempty" name:"WifiBSSID"`

	// 业务 ID，在多个业务中使用此服务，通过此ID 区分统计数据
	BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

	// 手机号加密类型 0：不加密（默认值）；1：md5；2：sha256；3：SM3
	PhoneCryptoType *string `json:"PhoneCryptoType,omitempty" name:"PhoneCryptoType"`

	// 身份证Id加密类型 0：不加密（默认值）；1：md5；2：sha256；3：SM3
	IdCryptoType *string `json:"IdCryptoType,omitempty" name:"IdCryptoType"`

	// 姓名加密类型 0：不加密（默认值）；1：md5；2：sha256；3：SM3
	NameCryptoType *string `json:"NameCryptoType,omitempty" name:"NameCryptoType"`
}

type FinanceAntiFraudRecord struct {

	// 表示该条记录能否查到：1为能查到，-1为查不到
	// 注意：此字段可能返回 null，表示取不到有效值。
	Found *string `json:"Found,omitempty" name:"Found"`

	// 表示该条Id能否查到：1为能查到，-1为查不到
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdFound *string `json:"IdFound,omitempty" name:"IdFound"`

	// 评分0~100;值越高 欺诈可能性越大
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskScore *string `json:"RiskScore,omitempty" name:"RiskScore"`

	// 扩展字段，对风险类型的说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskInfo []*RiskDetailInfo `json:"RiskInfo,omitempty" name:"RiskInfo"`

	// 多模型返回结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	OtherModelScores []*FinanceOtherModelScores `json:"OtherModelScores,omitempty" name:"OtherModelScores"`

	// 业务侧错误码。成功时返回0，错误时返回非0值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *string `json:"Code,omitempty" name:"Code"`

	// 业务侧错误信息。成功时返回Success，错误时返回具体业务错误原因。
	// 注意：此字段可能返回 null，表示取不到有效值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`
}

type FinanceOtherModelScores struct {

	// 模型ID序号
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 模型ID序号对应的评分结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelScore *string `json:"ModelScore,omitempty" name:"ModelScore"`
}

type QueryAntiFraudRequest struct {
	*tchttp.BaseRequest

	// 电话号码(五选二)
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// Id(五选二)
	IdNumber *string `json:"IdNumber,omitempty" name:"IdNumber"`

	// 银行卡号(五选二)
	BankCardNumber *string `json:"BankCardNumber,omitempty" name:"BankCardNumber"`

	// 用户请求来源 IP(五选二)
	UserIp *string `json:"UserIp,omitempty" name:"UserIp"`

	// 国际移动设备识别码(五选二)
	Imei *string `json:"Imei,omitempty" name:"Imei"`

	// ios 系统广告标示符(五选二)
	Idfa *string `json:"Idfa,omitempty" name:"Idfa"`

	// 业务场景 ID，需要找技术对接
	Scene *string `json:"Scene,omitempty" name:"Scene"`

	// 姓名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 用户邮箱地址
	EmailAddress *string `json:"EmailAddress,omitempty" name:"EmailAddress"`

	// 用户住址
	Address *string `json:"Address,omitempty" name:"Address"`

	// MAC 地址
	Mac *string `json:"Mac,omitempty" name:"Mac"`

	// 国际移动用户识别码
	Imsi *string `json:"Imsi,omitempty" name:"Imsi"`

	// 关联的腾讯帐号 QQ：1；
	// 开放帐号微信： 2；
	AccountType *string `json:"AccountType,omitempty" name:"AccountType"`

	// 可选的 QQ 或微信 openid
	Uid *string `json:"Uid,omitempty" name:"Uid"`

	// qq 或微信分配给网站或应用的 appid，用来
	// 唯一标识网站或应用
	AppIdU *string `json:"AppIdU,omitempty" name:"AppIdU"`

	// WIFI MAC
	WifiMac *string `json:"WifiMac,omitempty" name:"WifiMac"`

	// WIFI 服务集标识
	WifiSSID *string `json:"WifiSSID,omitempty" name:"WifiSSID"`

	// WIFI-BSSID
	WifiBSSID *string `json:"WifiBSSID,omitempty" name:"WifiBSSID"`

	// 业务 ID，在多个业务中使用此服务，通过此
	// ID 区分统计数据
	BusinessId *string `json:"BusinessId,omitempty" name:"BusinessId"`

	// Id加密类型
	// 0：不加密（默认值）
	// 1：md5
	// 2：sha256
	// 3：SM3
	IdCryptoType *string `json:"IdCryptoType,omitempty" name:"IdCryptoType"`

	// 手机号加密类型
	// 0：不加密（默认值）
	// 1：md5, 2：sha256
	// 3：SM3
	PhoneCryptoType *string `json:"PhoneCryptoType,omitempty" name:"PhoneCryptoType"`

	// 姓名加密类型
	// 0：不加密（默认值）
	// 1：md5
	// 2：sha256
	// 3：SM3
	NameCryptoType *string `json:"NameCryptoType,omitempty" name:"NameCryptoType"`
}

func (r *QueryAntiFraudRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryAntiFraudRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PhoneNumber")
	delete(f, "IdNumber")
	delete(f, "BankCardNumber")
	delete(f, "UserIp")
	delete(f, "Imei")
	delete(f, "Idfa")
	delete(f, "Scene")
	delete(f, "Name")
	delete(f, "EmailAddress")
	delete(f, "Address")
	delete(f, "Mac")
	delete(f, "Imsi")
	delete(f, "AccountType")
	delete(f, "Uid")
	delete(f, "AppIdU")
	delete(f, "WifiMac")
	delete(f, "WifiSSID")
	delete(f, "WifiBSSID")
	delete(f, "BusinessId")
	delete(f, "IdCryptoType")
	delete(f, "PhoneCryptoType")
	delete(f, "NameCryptoType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryAntiFraudRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryAntiFraudResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 表示该条记录能否查到：1为能查到，-1为查不到
		Found *int64 `json:"Found,omitempty" name:"Found"`

		// 表示该条Id能否查到：1为能查到，-1为查不到
		IdFound *int64 `json:"IdFound,omitempty" name:"IdFound"`

		// 0~100;值越高 欺诈可能性越大
		RiskScore *uint64 `json:"RiskScore,omitempty" name:"RiskScore"`

		// 扩展字段，对风险类型的说明
		RiskInfo []*RiskDetail `json:"RiskInfo,omitempty" name:"RiskInfo"`

		// 业务侧错误码。成功时返回Success，错误时返回具体业务错误原因。
	// 注意：此字段可能返回 null，表示取不到有效值。
		CodeDesc *string `json:"CodeDesc,omitempty" name:"CodeDesc"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryAntiFraudResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryAntiFraudResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RiskDetail struct {

	// 风险码 参数详细定义请加微信：TYXGJ-01
	RiskCode *uint64 `json:"RiskCode,omitempty" name:"RiskCode"`
}

type RiskDetailInfo struct {

	// 风险码
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskCode *string `json:"RiskCode,omitempty" name:"RiskCode"`

	// 风险码对应的风险值
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskValue *string `json:"RiskValue,omitempty" name:"RiskValue"`
}
