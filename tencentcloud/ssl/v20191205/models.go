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

package v20191205

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type ApiGatewayInstanceDetail struct {
	// 实例ID
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 实例名称
	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 证书ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// 使用协议
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

// Predefined struct for user
type ApplyCertificateRequestParams struct {
	// 验证方式：DNS_AUTO = 自动DNS验证，DNS = 手动DNS验证，FILE = 文件验证。
	DvAuthMethod *string `json:"DvAuthMethod,omitempty" name:"DvAuthMethod"`

	// 域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 项目 ID。
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 证书类型，目前仅支持类型2。2 = TrustAsia TLS RSA CA。
	PackageType *string `json:"PackageType,omitempty" name:"PackageType"`

	// 邮箱。
	ContactEmail *string `json:"ContactEmail,omitempty" name:"ContactEmail"`

	// 手机。
	ContactPhone *string `json:"ContactPhone,omitempty" name:"ContactPhone"`

	// 有效期，默认12个月，目前仅支持12个月。
	ValidityPeriod *string `json:"ValidityPeriod,omitempty" name:"ValidityPeriod"`

	// 加密算法，支持 RSA及ECC。
	CsrEncryptAlgo *string `json:"CsrEncryptAlgo,omitempty" name:"CsrEncryptAlgo"`

	// 密钥对参数，RSA仅支持2048。ECC仅支持prime256v1
	CsrKeyParameter *string `json:"CsrKeyParameter,omitempty" name:"CsrKeyParameter"`

	// CSR 的加密密码。
	CsrKeyPassword *string `json:"CsrKeyPassword,omitempty" name:"CsrKeyPassword"`

	// 备注名称。
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 原证书 ID，用于重新申请。
	OldCertificateId *string `json:"OldCertificateId,omitempty" name:"OldCertificateId"`

	// 权益包ID，用于免费证书扩容包使用
	PackageId *string `json:"PackageId,omitempty" name:"PackageId"`

	// 签发后是否删除自动域名验证记录， 默认为否；仅域名为DNS_AUTO验证类型支持传参
	DeleteDnsAutoRecord *bool `json:"DeleteDnsAutoRecord,omitempty" name:"DeleteDnsAutoRecord"`
}

type ApplyCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 验证方式：DNS_AUTO = 自动DNS验证，DNS = 手动DNS验证，FILE = 文件验证。
	DvAuthMethod *string `json:"DvAuthMethod,omitempty" name:"DvAuthMethod"`

	// 域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 项目 ID。
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 证书类型，目前仅支持类型2。2 = TrustAsia TLS RSA CA。
	PackageType *string `json:"PackageType,omitempty" name:"PackageType"`

	// 邮箱。
	ContactEmail *string `json:"ContactEmail,omitempty" name:"ContactEmail"`

	// 手机。
	ContactPhone *string `json:"ContactPhone,omitempty" name:"ContactPhone"`

	// 有效期，默认12个月，目前仅支持12个月。
	ValidityPeriod *string `json:"ValidityPeriod,omitempty" name:"ValidityPeriod"`

	// 加密算法，支持 RSA及ECC。
	CsrEncryptAlgo *string `json:"CsrEncryptAlgo,omitempty" name:"CsrEncryptAlgo"`

	// 密钥对参数，RSA仅支持2048。ECC仅支持prime256v1
	CsrKeyParameter *string `json:"CsrKeyParameter,omitempty" name:"CsrKeyParameter"`

	// CSR 的加密密码。
	CsrKeyPassword *string `json:"CsrKeyPassword,omitempty" name:"CsrKeyPassword"`

	// 备注名称。
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 原证书 ID，用于重新申请。
	OldCertificateId *string `json:"OldCertificateId,omitempty" name:"OldCertificateId"`

	// 权益包ID，用于免费证书扩容包使用
	PackageId *string `json:"PackageId,omitempty" name:"PackageId"`

	// 签发后是否删除自动域名验证记录， 默认为否；仅域名为DNS_AUTO验证类型支持传参
	DeleteDnsAutoRecord *bool `json:"DeleteDnsAutoRecord,omitempty" name:"DeleteDnsAutoRecord"`
}

func (r *ApplyCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DvAuthMethod")
	delete(f, "DomainName")
	delete(f, "ProjectId")
	delete(f, "PackageType")
	delete(f, "ContactEmail")
	delete(f, "ContactPhone")
	delete(f, "ValidityPeriod")
	delete(f, "CsrEncryptAlgo")
	delete(f, "CsrKeyParameter")
	delete(f, "CsrKeyPassword")
	delete(f, "Alias")
	delete(f, "OldCertificateId")
	delete(f, "PackageId")
	delete(f, "DeleteDnsAutoRecord")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyCertificateResponseParams struct {
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ApplyCertificateResponse struct {
	*tchttp.BaseResponse
	Response *ApplyCertificateResponseParams `json:"Response"`
}

func (r *ApplyCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelCertificateOrderRequestParams struct {
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

type CancelCertificateOrderRequest struct {
	*tchttp.BaseRequest
	
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

func (r *CancelCertificateOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelCertificateOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelCertificateOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelCertificateOrderResponseParams struct {
	// 取消订单成功的证书 ID。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CancelCertificateOrderResponse struct {
	*tchttp.BaseResponse
	Response *CancelCertificateOrderResponseParams `json:"Response"`
}

func (r *CancelCertificateOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelCertificateOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CdnInstanceDetail struct {
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 已部署证书ID
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// 域名状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 域名计费状态
	HttpsBillingSwitch *string `json:"HttpsBillingSwitch,omitempty" name:"HttpsBillingSwitch"`
}

type CertHostingInfo struct {
	// 证书ID
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// 已替换的新证书ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewCertId *string `json:"RenewCertId,omitempty" name:"RenewCertId"`

	// 云资源托管 ，CDN或CLB：部分开启，CDN,CLB：已开启，null：未开启托管
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type Certificate struct {
	// 证书ID
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// 证书绑定的域名
	DnsNames []*string `json:"DnsNames,omitempty" name:"DnsNames"`

	// 根证书ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertCaId *string `json:"CertCaId,omitempty" name:"CertCaId"`

	// 证书认证模式：UNIDIRECTIONAL单向认证，MUTUAL双向认证
	// 注意：此字段可能返回 null，表示取不到有效值。
	SSLMode *string `json:"SSLMode,omitempty" name:"SSLMode"`
}

type CertificateExtra struct {
	// 证书可配置域名数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainNumber *string `json:"DomainNumber,omitempty" name:"DomainNumber"`

	// 原始证书 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginCertificateId *string `json:"OriginCertificateId,omitempty" name:"OriginCertificateId"`

	// 重颁发证书原始 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReplacedBy *string `json:"ReplacedBy,omitempty" name:"ReplacedBy"`

	// 重颁发证书新 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReplacedFor *string `json:"ReplacedFor,omitempty" name:"ReplacedFor"`

	// 新订单证书 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewOrder *string `json:"RenewOrder,omitempty" name:"RenewOrder"`

	// 是否是国密证书
	// 注意：此字段可能返回 null，表示取不到有效值。
	SMCert *int64 `json:"SMCert,omitempty" name:"SMCert"`
}

type Certificates struct {
	// 用户 UIN。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 项目 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 证书来源。
	// 注意：此字段可能返回 null，表示取不到有效值。
	From *string `json:"From,omitempty" name:"From"`

	// 证书套餐类型：
	// null = 用户上传证书（没有套餐类型），
	// 1 = GeoTrust DV SSL CA - G3， 2 = TrustAsia TLS RSA CA， 3 = SecureSite 增强型企业版（EV Pro）， 4 = SecureSite 增强型（EV）， 5 = SecureSite 企业型专业版（OV Pro）， 6 = SecureSite 企业型（OV）， 7 = SecureSite 企业型（OV）通配符， 8 = Geotrust 增强型（EV）， 9 = Geotrust 企业型（OV）， 10 = Geotrust 企业型（OV）通配符， 11 = TrustAsia 域名型多域名 SSL 证书， 12 = TrustAsia 域名型（DV）通配符， 13 = TrustAsia 企业型通配符（OV）SSL 证书（D3）， 14 = TrustAsia 企业型（OV）SSL 证书（D3）， 15 = TrustAsia 企业型多域名 （OV）SSL 证书（D3）， 16 = TrustAsia 增强型 （EV）SSL 证书（D3）， 17 = TrustAsia 增强型多域名（EV）SSL 证书（D3）， 18 = GlobalSign 企业型（OV）SSL 证书， 19 = GlobalSign 企业型通配符 （OV）SSL 证书， 20 = GlobalSign 增强型 （EV）SSL 证书， 21 = TrustAsia 企业型通配符多域名（OV）SSL 证书（D3）， 22 = GlobalSign 企业型多域名（OV）SSL 证书， 23 = GlobalSign 企业型通配符多域名（OV）SSL 证书， 24 = GlobalSign 增强型多域名（EV）SSL 证书。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageType *string `json:"PackageType,omitempty" name:"PackageType"`

	// 证书类型：CA = 客户端证书，SVR = 服务器证书。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertificateType *string `json:"CertificateType,omitempty" name:"CertificateType"`

	// 颁发者。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductZhName *string `json:"ProductZhName,omitempty" name:"ProductZhName"`

	// 主域名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 备注名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 状态。0：审核中，1：已通过，2：审核失败，3：已过期，4：验证方式为 DNS_AUTO 类型的证书， 已添加DNS记录，5：企业证书，待提交，6：订单取消中，7：已取消，8：已提交资料， 待上传确认函，9：证书吊销中，10：已吊销，11：重颁发中，12：待上传吊销确认函，13：免费证书待提交资料状态，14：已退款，
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 证书扩展信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertificateExtra *CertificateExtra `json:"CertificateExtra,omitempty" name:"CertificateExtra"`

	// 漏洞扫描状态：INACTIVE = 未开启，ACTIVE = 已开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulnerabilityStatus *string `json:"VulnerabilityStatus,omitempty" name:"VulnerabilityStatus"`

	// 状态信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusMsg *string `json:"StatusMsg,omitempty" name:"StatusMsg"`

	// 验证类型：DNS_AUTO = 自动DNS验证，DNS = 手动DNS验证，FILE = 文件验证，EMAIL = 邮件验证。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerifyType *string `json:"VerifyType,omitempty" name:"VerifyType"`

	// 证书生效时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertBeginTime *string `json:"CertBeginTime,omitempty" name:"CertBeginTime"`

	// 证书过期时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertEndTime *string `json:"CertEndTime,omitempty" name:"CertEndTime"`

	// 证书有效期，单位（月）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValidityPeriod *string `json:"ValidityPeriod,omitempty" name:"ValidityPeriod"`

	// 创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`

	// 证书 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 证书包含的多个域名（包含主域名）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubjectAltName []*string `json:"SubjectAltName,omitempty" name:"SubjectAltName"`

	// 证书类型名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageTypeName *string `json:"PackageTypeName,omitempty" name:"PackageTypeName"`

	// 状态名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusName *string `json:"StatusName,omitempty" name:"StatusName"`

	// 是否为 VIP 客户。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsVip *bool `json:"IsVip,omitempty" name:"IsVip"`

	// 是否为 DV 版证书。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDv *bool `json:"IsDv,omitempty" name:"IsDv"`

	// 是否为泛域名证书。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsWildcard *bool `json:"IsWildcard,omitempty" name:"IsWildcard"`

	// 是否启用了漏洞扫描功能。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsVulnerability *bool `json:"IsVulnerability,omitempty" name:"IsVulnerability"`

	// 是否可重颁发证书。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewAble *bool `json:"RenewAble,omitempty" name:"RenewAble"`

	// 项目信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectInfo *ProjectInfo `json:"ProjectInfo,omitempty" name:"ProjectInfo"`

	// 关联的云资源，暂不可用
	// 注意：此字段可能返回 null，表示取不到有效值。
	BoundResource []*string `json:"BoundResource,omitempty" name:"BoundResource"`

	// 是否可部署。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Deployable *bool `json:"Deployable,omitempty" name:"Deployable"`

	// 标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tags `json:"Tags,omitempty" name:"Tags"`

	// 是否已忽略到期通知
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsIgnore *bool `json:"IsIgnore,omitempty" name:"IsIgnore"`

	// 是否国密证书
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsSM *bool `json:"IsSM,omitempty" name:"IsSM"`

	// 证书算法
	// 注意：此字段可能返回 null，表示取不到有效值。
	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitempty" name:"EncryptAlgorithm"`

	// 上传CA证书的加密算法
	// 注意：此字段可能返回 null，表示取不到有效值。
	CAEncryptAlgorithms []*string `json:"CAEncryptAlgorithms,omitempty" name:"CAEncryptAlgorithms"`

	// 上传CA证书的过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CAEndTimes []*string `json:"CAEndTimes,omitempty" name:"CAEndTimes"`

	// 上传CA证书的通用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	CACommonNames []*string `json:"CACommonNames,omitempty" name:"CACommonNames"`

	// 证书预审核信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PreAuditInfo *PreAuditInfo `json:"PreAuditInfo,omitempty" name:"PreAuditInfo"`

	// 是否自动续费
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
}

// Predefined struct for user
type CheckCertificateChainRequestParams struct {
	// 待检查的证书链
	CertificateChain *string `json:"CertificateChain,omitempty" name:"CertificateChain"`
}

type CheckCertificateChainRequest struct {
	*tchttp.BaseRequest
	
	// 待检查的证书链
	CertificateChain *string `json:"CertificateChain,omitempty" name:"CertificateChain"`
}

func (r *CheckCertificateChainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckCertificateChainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateChain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckCertificateChainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckCertificateChainResponseParams struct {
	// true为通过检查，false为未通过检查。
	IsValid *bool `json:"IsValid,omitempty" name:"IsValid"`

	// true为可信CA，false为不可信CA。
	IsTrustedCA *bool `json:"IsTrustedCA,omitempty" name:"IsTrustedCA"`

	// 包含证书链中每一段证书的通用名称。
	Chains []*string `json:"Chains,omitempty" name:"Chains"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CheckCertificateChainResponse struct {
	*tchttp.BaseResponse
	Response *CheckCertificateChainResponseParams `json:"Response"`
}

func (r *CheckCertificateChainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckCertificateChainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClbInstanceDetail struct {
	// CLB实例ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// CLB实例名称
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// CLB监听器列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Listeners []*ClbListener `json:"Listeners,omitempty" name:"Listeners"`
}

type ClbListener struct {
	// 监听器ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 监听器名称
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 是否开启SNI，1为开启，0为关闭
	SniSwitch *uint64 `json:"SniSwitch,omitempty" name:"SniSwitch"`

	// 监听器协议类型， HTTPS|TCP_SSL
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 监听器绑定的证书数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Certificate *Certificate `json:"Certificate,omitempty" name:"Certificate"`

	// 监听器规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rules []*ClbListenerRule `json:"Rules,omitempty" name:"Rules"`

	// 不匹配域名列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NoMatchDomains []*string `json:"NoMatchDomains,omitempty" name:"NoMatchDomains"`
}

type ClbListenerRule struct {
	// 规则ID
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// 规则绑定的域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 规则是否匹配待绑定证书的域名
	IsMatch *bool `json:"IsMatch,omitempty" name:"IsMatch"`

	// 规则已绑定的证书数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Certificate *Certificate `json:"Certificate,omitempty" name:"Certificate"`

	// 不匹配域名列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NoMatchDomains []*string `json:"NoMatchDomains,omitempty" name:"NoMatchDomains"`
}

// Predefined struct for user
type CommitCertificateInformationRequestParams struct {
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 域名验证方式
	VerifyType *string `json:"VerifyType,omitempty" name:"VerifyType"`
}

type CommitCertificateInformationRequest struct {
	*tchttp.BaseRequest
	
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 域名验证方式
	VerifyType *string `json:"VerifyType,omitempty" name:"VerifyType"`
}

func (r *CommitCertificateInformationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CommitCertificateInformationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	delete(f, "VerifyType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CommitCertificateInformationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CommitCertificateInformationResponseParams struct {
	// CA机构侧订单号。
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 证书状态：0 = 审核中，1 = 已通过，2 = 审核失败，3 = 已过期，4 = 已添加DNS记录，5 = 企业证书，待提交，6 = 订单取消中，7 = 已取消，8 = 已提交资料， 待上传确认函，9 = 证书吊销中，10 = 已吊销，11 = 重颁发中，12 = 待上传吊销确认函，13 = 免费证书待提交资料。
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CommitCertificateInformationResponse struct {
	*tchttp.BaseResponse
	Response *CommitCertificateInformationResponseParams `json:"Response"`
}

func (r *CommitCertificateInformationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CommitCertificateInformationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CompanyInfo struct {
	// 公司名称
	CompanyName *string `json:"CompanyName,omitempty" name:"CompanyName"`

	// 公司ID
	CompanyId *int64 `json:"CompanyId,omitempty" name:"CompanyId"`

	// 公司所在国家
	CompanyCountry *string `json:"CompanyCountry,omitempty" name:"CompanyCountry"`

	// 公司所在省份
	CompanyProvince *string `json:"CompanyProvince,omitempty" name:"CompanyProvince"`

	// 公司所在城市
	CompanyCity *string `json:"CompanyCity,omitempty" name:"CompanyCity"`

	// 公司所在详细地址
	CompanyAddress *string `json:"CompanyAddress,omitempty" name:"CompanyAddress"`

	// 公司电话
	CompanyPhone *string `json:"CompanyPhone,omitempty" name:"CompanyPhone"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdType *string `json:"IdType,omitempty" name:"IdType"`

	// ID号
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdNumber *string `json:"IdNumber,omitempty" name:"IdNumber"`
}

// Predefined struct for user
type CompleteCertificateRequestParams struct {
	// 证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

type CompleteCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

func (r *CompleteCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CompleteCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CompleteCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CompleteCertificateResponseParams struct {
	// 证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CompleteCertificateResponse struct {
	*tchttp.BaseResponse
	Response *CompleteCertificateResponseParams `json:"Response"`
}

func (r *CompleteCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CompleteCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CosInstanceDetail struct {
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 已绑定的证书ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// ENABLED: 域名上线状态
	// DISABLED:域名下线状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 存储桶名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// 存储桶地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`
}

// Predefined struct for user
type CreateCertificateByPackageRequestParams struct {
	// 证书产品PID。
	ProductPid *uint64 `json:"ProductPid,omitempty" name:"ProductPid"`

	// 要消耗的权益包ID。
	PackageIds []*string `json:"PackageIds,omitempty" name:"PackageIds"`

	// 证书域名数量。
	DomainCount *string `json:"DomainCount,omitempty" name:"DomainCount"`

	// 多年期证书年限。
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 要续费的原证书ID（续费时填写）。
	OldCertificateId *string `json:"OldCertificateId,omitempty" name:"OldCertificateId"`

	// 续费时CSR生成方式（original、upload、online）。
	RenewGenCsrMethod *string `json:"RenewGenCsrMethod,omitempty" name:"RenewGenCsrMethod"`

	// 续费时选择上传CSR时填写CSR。
	RenewCsr *string `json:"RenewCsr,omitempty" name:"RenewCsr"`

	// 续费证书CSR的算法类型。
	RenewAlgorithmType *string `json:"RenewAlgorithmType,omitempty" name:"RenewAlgorithmType"`

	// 续费证书CSR的算法参数。
	RenewAlgorithmParam *string `json:"RenewAlgorithmParam,omitempty" name:"RenewAlgorithmParam"`

	// 项目ID。
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 标签。
	Tags []*Tags `json:"Tags,omitempty" name:"Tags"`

	// 续费证书的私钥密码。
	RenewKeyPass *string `json:"RenewKeyPass,omitempty" name:"RenewKeyPass"`

	// 批量购买证书时预填写的域名。
	DomainNames *string `json:"DomainNames,omitempty" name:"DomainNames"`

	// 批量购买证书数量。
	CertificateCount *uint64 `json:"CertificateCount,omitempty" name:"CertificateCount"`

	// 预填写的管理人ID。
	ManagerId *uint64 `json:"ManagerId,omitempty" name:"ManagerId"`

	// 预填写的公司ID。
	CompanyId *uint64 `json:"CompanyId,omitempty" name:"CompanyId"`

	// 验证方式
	VerifyType *string `json:"VerifyType,omitempty" name:"VerifyType"`
}

type CreateCertificateByPackageRequest struct {
	*tchttp.BaseRequest
	
	// 证书产品PID。
	ProductPid *uint64 `json:"ProductPid,omitempty" name:"ProductPid"`

	// 要消耗的权益包ID。
	PackageIds []*string `json:"PackageIds,omitempty" name:"PackageIds"`

	// 证书域名数量。
	DomainCount *string `json:"DomainCount,omitempty" name:"DomainCount"`

	// 多年期证书年限。
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 要续费的原证书ID（续费时填写）。
	OldCertificateId *string `json:"OldCertificateId,omitempty" name:"OldCertificateId"`

	// 续费时CSR生成方式（original、upload、online）。
	RenewGenCsrMethod *string `json:"RenewGenCsrMethod,omitempty" name:"RenewGenCsrMethod"`

	// 续费时选择上传CSR时填写CSR。
	RenewCsr *string `json:"RenewCsr,omitempty" name:"RenewCsr"`

	// 续费证书CSR的算法类型。
	RenewAlgorithmType *string `json:"RenewAlgorithmType,omitempty" name:"RenewAlgorithmType"`

	// 续费证书CSR的算法参数。
	RenewAlgorithmParam *string `json:"RenewAlgorithmParam,omitempty" name:"RenewAlgorithmParam"`

	// 项目ID。
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 标签。
	Tags []*Tags `json:"Tags,omitempty" name:"Tags"`

	// 续费证书的私钥密码。
	RenewKeyPass *string `json:"RenewKeyPass,omitempty" name:"RenewKeyPass"`

	// 批量购买证书时预填写的域名。
	DomainNames *string `json:"DomainNames,omitempty" name:"DomainNames"`

	// 批量购买证书数量。
	CertificateCount *uint64 `json:"CertificateCount,omitempty" name:"CertificateCount"`

	// 预填写的管理人ID。
	ManagerId *uint64 `json:"ManagerId,omitempty" name:"ManagerId"`

	// 预填写的公司ID。
	CompanyId *uint64 `json:"CompanyId,omitempty" name:"CompanyId"`

	// 验证方式
	VerifyType *string `json:"VerifyType,omitempty" name:"VerifyType"`
}

func (r *CreateCertificateByPackageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCertificateByPackageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductPid")
	delete(f, "PackageIds")
	delete(f, "DomainCount")
	delete(f, "Period")
	delete(f, "OldCertificateId")
	delete(f, "RenewGenCsrMethod")
	delete(f, "RenewCsr")
	delete(f, "RenewAlgorithmType")
	delete(f, "RenewAlgorithmParam")
	delete(f, "ProjectId")
	delete(f, "Tags")
	delete(f, "RenewKeyPass")
	delete(f, "DomainNames")
	delete(f, "CertificateCount")
	delete(f, "ManagerId")
	delete(f, "CompanyId")
	delete(f, "VerifyType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCertificateByPackageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCertificateByPackageResponseParams struct {
	// 证书ID。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 批量购买证书时返回多个证书ID。
	CertificateIds []*string `json:"CertificateIds,omitempty" name:"CertificateIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCertificateByPackageResponse struct {
	*tchttp.BaseResponse
	Response *CreateCertificateByPackageResponseParams `json:"Response"`
}

func (r *CreateCertificateByPackageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCertificateByPackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCertificateRequestParams struct {
	// 证书商品ID，3 = SecureSite 增强型企业版（EV Pro）， 4 = SecureSite 增强型（EV）， 5 = SecureSite 企业型专业版（OV Pro）， 6 = SecureSite 企业型（OV）， 7 = SecureSite 企业型（OV）通配符， 8 = Geotrust 增强型（EV）， 9 = Geotrust 企业型（OV）， 10 = Geotrust 企业型（OV）通配符， 11 = TrustAsia 域名型多域名 SSL 证书， 12 = TrustAsia 域名型（DV）通配符， 13 = TrustAsia 企业型通配符（OV）SSL 证书（D3）， 14 = TrustAsia 企业型（OV）SSL 证书（D3）， 15 = TrustAsia 企业型多域名 （OV）SSL 证书（D3）， 16 = TrustAsia 增强型 （EV）SSL 证书（D3）， 17 = TrustAsia 增强型多域名（EV）SSL 证书（D3）， 18 = GlobalSign 企业型（OV）SSL 证书， 19 = GlobalSign 企业型通配符 （OV）SSL 证书， 20 = GlobalSign 增强型 （EV）SSL 证书， 21 = TrustAsia 企业型通配符多域名（OV）SSL 证书（D3）， 22 = GlobalSign 企业型多域名（OV）SSL 证书， 23 = GlobalSign 企业型通配符多域名（OV）SSL 证书， 24 = GlobalSign 增强型多域名（EV）SSL 证书，25 = Wotrus 域名型证书，26 = Wotrus 域名型多域名证书，27 = Wotrus 域名型通配符证书，28 = Wotrus 企业型证书，29 = Wotrus 企业型多域名证书，30 = Wotrus 企业型通配符证书，31 = Wotrus 增强型证书，32 = Wotrus 增强型多域名证书，33 = DNSPod 国密域名型证书，34 = DNSPod 国密域名型多域名证书，35 = DNSPod 国密域名型通配符证书，37 = DNSPod 国密企业型证书，38 = DNSPod 国密企业型多域名证书，39 = DNSPod 国密企业型通配符证书，40 = DNSPod 国密增强型证书，41 = DNSPod 国密增强型多域名证书，42 = TrustAsia 域名型通配符多域名证书。
	ProductId *int64 `json:"ProductId,omitempty" name:"ProductId"`

	// 证书包含的域名数量
	DomainNum *int64 `json:"DomainNum,omitempty" name:"DomainNum"`

	// 证书年限，当前只支持 1 年证书的购买
	TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`
}

type CreateCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 证书商品ID，3 = SecureSite 增强型企业版（EV Pro）， 4 = SecureSite 增强型（EV）， 5 = SecureSite 企业型专业版（OV Pro）， 6 = SecureSite 企业型（OV）， 7 = SecureSite 企业型（OV）通配符， 8 = Geotrust 增强型（EV）， 9 = Geotrust 企业型（OV）， 10 = Geotrust 企业型（OV）通配符， 11 = TrustAsia 域名型多域名 SSL 证书， 12 = TrustAsia 域名型（DV）通配符， 13 = TrustAsia 企业型通配符（OV）SSL 证书（D3）， 14 = TrustAsia 企业型（OV）SSL 证书（D3）， 15 = TrustAsia 企业型多域名 （OV）SSL 证书（D3）， 16 = TrustAsia 增强型 （EV）SSL 证书（D3）， 17 = TrustAsia 增强型多域名（EV）SSL 证书（D3）， 18 = GlobalSign 企业型（OV）SSL 证书， 19 = GlobalSign 企业型通配符 （OV）SSL 证书， 20 = GlobalSign 增强型 （EV）SSL 证书， 21 = TrustAsia 企业型通配符多域名（OV）SSL 证书（D3）， 22 = GlobalSign 企业型多域名（OV）SSL 证书， 23 = GlobalSign 企业型通配符多域名（OV）SSL 证书， 24 = GlobalSign 增强型多域名（EV）SSL 证书，25 = Wotrus 域名型证书，26 = Wotrus 域名型多域名证书，27 = Wotrus 域名型通配符证书，28 = Wotrus 企业型证书，29 = Wotrus 企业型多域名证书，30 = Wotrus 企业型通配符证书，31 = Wotrus 增强型证书，32 = Wotrus 增强型多域名证书，33 = DNSPod 国密域名型证书，34 = DNSPod 国密域名型多域名证书，35 = DNSPod 国密域名型通配符证书，37 = DNSPod 国密企业型证书，38 = DNSPod 国密企业型多域名证书，39 = DNSPod 国密企业型通配符证书，40 = DNSPod 国密增强型证书，41 = DNSPod 国密增强型多域名证书，42 = TrustAsia 域名型通配符多域名证书。
	ProductId *int64 `json:"ProductId,omitempty" name:"ProductId"`

	// 证书包含的域名数量
	DomainNum *int64 `json:"DomainNum,omitempty" name:"DomainNum"`

	// 证书年限，当前只支持 1 年证书的购买
	TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`
}

func (r *CreateCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductId")
	delete(f, "DomainNum")
	delete(f, "TimeSpan")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCertificateResponseParams struct {
	// 证书ID列表
	CertificateIds []*string `json:"CertificateIds,omitempty" name:"CertificateIds"`

	// 订单号列表
	DealIds []*string `json:"DealIds,omitempty" name:"DealIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCertificateResponse struct {
	*tchttp.BaseResponse
	Response *CreateCertificateResponseParams `json:"Response"`
}

func (r *CreateCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DdosInstanceDetail struct {
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 协议类型
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 证书ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// 转发端口
	VirtualPort *string `json:"VirtualPort,omitempty" name:"VirtualPort"`
}

// Predefined struct for user
type DeleteCertificateRequestParams struct {
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

type DeleteCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

func (r *DeleteCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCertificateResponseParams struct {
	// 删除结果（true：删除成功，false：删除失败）
	DeleteResult *bool `json:"DeleteResult,omitempty" name:"DeleteResult"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteCertificateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCertificateResponseParams `json:"Response"`
}

func (r *DeleteCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteManagerRequestParams struct {
	// 管理人ID
	ManagerId *int64 `json:"ManagerId,omitempty" name:"ManagerId"`
}

type DeleteManagerRequest struct {
	*tchttp.BaseRequest
	
	// 管理人ID
	ManagerId *int64 `json:"ManagerId,omitempty" name:"ManagerId"`
}

func (r *DeleteManagerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteManagerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ManagerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteManagerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteManagerResponseParams struct {
	// 管理人ID
	ManagerId *int64 `json:"ManagerId,omitempty" name:"ManagerId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteManagerResponse struct {
	*tchttp.BaseResponse
	Response *DeleteManagerResponseParams `json:"Response"`
}

func (r *DeleteManagerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteManagerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeployCertificateInstanceRequestParams struct {
	// 待部署的证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 需要部署实例列表
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList"`

	// 部署的云资源类型
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 部署云资源状态：
	// 云直播：
	// -1：域名未关联证书。
	// 1： 域名https已开启。
	// 0： 域名https已关闭。
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type DeployCertificateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 待部署的证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 需要部署实例列表
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList"`

	// 部署的云资源类型
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 部署云资源状态：
	// 云直播：
	// -1：域名未关联证书。
	// 1： 域名https已开启。
	// 0： 域名https已关闭。
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *DeployCertificateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeployCertificateInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	delete(f, "InstanceIdList")
	delete(f, "ResourceType")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeployCertificateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeployCertificateInstanceResponseParams struct {
	// 云资源部署任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployRecordId *uint64 `json:"DeployRecordId,omitempty" name:"DeployRecordId"`

	// 部署状态，1表示部署成功，0表示部署失败
	DeployStatus *int64 `json:"DeployStatus,omitempty" name:"DeployStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeployCertificateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DeployCertificateInstanceResponseParams `json:"Response"`
}

func (r *DeployCertificateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeployCertificateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeployCertificateRecordRetryRequestParams struct {
	// 待重试部署记录ID
	DeployRecordId *int64 `json:"DeployRecordId,omitempty" name:"DeployRecordId"`

	// 待重试部署记录详情ID
	DeployRecordDetailId *int64 `json:"DeployRecordDetailId,omitempty" name:"DeployRecordDetailId"`
}

type DeployCertificateRecordRetryRequest struct {
	*tchttp.BaseRequest
	
	// 待重试部署记录ID
	DeployRecordId *int64 `json:"DeployRecordId,omitempty" name:"DeployRecordId"`

	// 待重试部署记录详情ID
	DeployRecordDetailId *int64 `json:"DeployRecordDetailId,omitempty" name:"DeployRecordDetailId"`
}

func (r *DeployCertificateRecordRetryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeployCertificateRecordRetryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeployRecordId")
	delete(f, "DeployRecordDetailId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeployCertificateRecordRetryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeployCertificateRecordRetryResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeployCertificateRecordRetryResponse struct {
	*tchttp.BaseResponse
	Response *DeployCertificateRecordRetryResponseParams `json:"Response"`
}

func (r *DeployCertificateRecordRetryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeployCertificateRecordRetryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeployCertificateRecordRollbackRequestParams struct {
	// 待重试部署记录ID
	DeployRecordId *int64 `json:"DeployRecordId,omitempty" name:"DeployRecordId"`
}

type DeployCertificateRecordRollbackRequest struct {
	*tchttp.BaseRequest
	
	// 待重试部署记录ID
	DeployRecordId *int64 `json:"DeployRecordId,omitempty" name:"DeployRecordId"`
}

func (r *DeployCertificateRecordRollbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeployCertificateRecordRollbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeployRecordId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeployCertificateRecordRollbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeployCertificateRecordRollbackResponseParams struct {
	// 回滚部署记录ID
	DeployRecordId *int64 `json:"DeployRecordId,omitempty" name:"DeployRecordId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeployCertificateRecordRollbackResponse struct {
	*tchttp.BaseResponse
	Response *DeployCertificateRecordRollbackResponseParams `json:"Response"`
}

func (r *DeployCertificateRecordRollbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeployCertificateRecordRollbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeployRecordDetail struct {
	// 部署记录详情ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 部署证书ID
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// 原绑定证书ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	OldCertId *string `json:"OldCertId,omitempty" name:"OldCertId"`

	// 部署实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 部署实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 部署监听器ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 部署域名列表
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 部署监听器协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 部署状态
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 部署错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`

	// 部署记录详情创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 部署记录详情最后一次更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 部署监听器名称
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 是否开启SNI
	SniSwitch *int64 `json:"SniSwitch,omitempty" name:"SniSwitch"`

	// COS存储桶名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// secret名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// 端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// TCB环境ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 部署的TCB类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TCBType *string `json:"TCBType,omitempty" name:"TCBType"`

	// 部署的TCB地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`
}

type DeployRecordInfo struct {
	// 部署记录ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 部署证书ID
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// 部署资源类型
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 部署地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 部署状态
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 部署时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 最近一次更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type DeployedResources struct {
	// 证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 数量
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 资源标识:clb,cdn,live,waf,antiddos
	Type *string `json:"Type,omitempty" name:"Type"`

	// 不建议使用。字段返回和Resources相同。本字段后续只返回null
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// 关联资源ID或关联域名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resources []*string `json:"Resources,omitempty" name:"Resources"`
}

// Predefined struct for user
type DescribeCertificateDetailRequestParams struct {
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

type DescribeCertificateDetailRequest struct {
	*tchttp.BaseRequest
	
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

func (r *DescribeCertificateDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificateDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCertificateDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertificateDetailResponseParams struct {
	// 用户 UIN。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 项目 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 证书来源：trustasia = 亚洲诚信，upload = 用户上传。
	// 注意：此字段可能返回 null，表示取不到有效值。
	From *string `json:"From,omitempty" name:"From"`

	// 证书类型：CA = 客户端证书，SVR = 服务器证书。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertificateType *string `json:"CertificateType,omitempty" name:"CertificateType"`

	// 证书套餐类型：null = 用户上传证书（没有套餐类型），1 = GeoTrust DV SSL CA - G3， 2 = TrustAsia TLS RSA CA， 3 = SecureSite 增强型企业版（EV Pro）， 4 = SecureSite 增强型（EV）， 5 = SecureSite 企业型专业版（OV Pro）， 6 = SecureSite 企业型（OV）， 7 = SecureSite 企业型（OV）通配符， 8 = Geotrust 增强型（EV）， 9 = Geotrust 企业型（OV）， 10 = Geotrust 企业型（OV）通配符， 11 = TrustAsia 域名型多域名 SSL 证书， 12 = TrustAsia 域名型（DV）通配符， 13 = TrustAsia 企业型通配符（OV）SSL 证书（D3）， 14 = TrustAsia 企业型（OV）SSL 证书（D3）， 15 = TrustAsia 企业型多域名 （OV）SSL 证书（D3）， 16 = TrustAsia 增强型 （EV）SSL 证书（D3）， 17 = TrustAsia 增强型多域名（EV）SSL 证书（D3）， 18 = GlobalSign 企业型（OV）SSL 证书， 19 = GlobalSign 企业型通配符 （OV）SSL 证书， 20 = GlobalSign 增强型 （EV）SSL 证书， 21 = TrustAsia 企业型通配符多域名（OV）SSL 证书（D3）， 22 = GlobalSign 企业型多域名（OV）SSL 证书， 23 = GlobalSign 企业型通配符多域名（OV）SSL 证书， 24 = GlobalSign 增强型多域名（EV）SSL 证书，25 = Wotrus 域名型证书，26 = Wotrus 域名型多域名证书，27 = Wotrus 域名型通配符证书，28 = Wotrus 企业型证书，29 = Wotrus 企业型多域名证书，30 = Wotrus 企业型通配符证书，31 = Wotrus 增强型证书，32 = Wotrus 增强型多域名证书，33 = DNSPod 国密域名型证书，34 = DNSPod 国密域名型多域名证书，35 = DNSPod 国密域名型通配符证书，37 = DNSPod 国密企业型证书，38 = DNSPod 国密企业型多域名证书，39 = DNSPod 国密企业型通配符证书，40 = DNSPod 国密增强型证书，41 = DNSPod 国密增强型多域名证书，42 = TrustAsia 域名型通配符多域名证书。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageType *string `json:"PackageType,omitempty" name:"PackageType"`

	// 颁发者。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductZhName *string `json:"ProductZhName,omitempty" name:"ProductZhName"`

	// 域名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 备注名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 证书状态：0 = 审核中，1 = 已通过，2 = 审核失败，3 = 已过期，4 = 已添加DNS记录，5 = 企业证书，待提交，6 = 订单取消中，7 = 已取消，8 = 已提交资料， 待上传确认函，9 = 证书吊销中，10 = 已吊销，11 = 重颁发中，12 = 待上传吊销确认函，13 = 免费证书待提交资料。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 状态信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusMsg *string `json:"StatusMsg,omitempty" name:"StatusMsg"`

	// 验证类型：DNS_AUTO = 自动DNS验证，DNS = 手动DNS验证，FILE = 文件验证，EMAIL = 邮件验证。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerifyType *string `json:"VerifyType,omitempty" name:"VerifyType"`

	// 漏洞扫描状态。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulnerabilityStatus *string `json:"VulnerabilityStatus,omitempty" name:"VulnerabilityStatus"`

	// 证书生效时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertBeginTime *string `json:"CertBeginTime,omitempty" name:"CertBeginTime"`

	// 证书失效时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertEndTime *string `json:"CertEndTime,omitempty" name:"CertEndTime"`

	// 证书有效期：单位（月）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValidityPeriod *string `json:"ValidityPeriod,omitempty" name:"ValidityPeriod"`

	// 申请时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`

	// 订单 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 证书扩展信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertificateExtra *CertificateExtra `json:"CertificateExtra,omitempty" name:"CertificateExtra"`

	// 证书私钥
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertificatePrivateKey *string `json:"CertificatePrivateKey,omitempty" name:"CertificatePrivateKey"`

	// 证书公钥（即证书内容）
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertificatePublicKey *string `json:"CertificatePublicKey,omitempty" name:"CertificatePublicKey"`

	// DV 认证信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DvAuthDetail *DvAuthDetail `json:"DvAuthDetail,omitempty" name:"DvAuthDetail"`

	// 漏洞扫描评估报告。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulnerabilityReport *string `json:"VulnerabilityReport,omitempty" name:"VulnerabilityReport"`

	// 证书 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 证书类型名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TypeName *string `json:"TypeName,omitempty" name:"TypeName"`

	// 状态描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusName *string `json:"StatusName,omitempty" name:"StatusName"`

	// 证书包含的多个域名（不包含主域名，主域名使用Domain字段）
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubjectAltName []*string `json:"SubjectAltName,omitempty" name:"SubjectAltName"`

	// 是否为付费证书。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsVip *bool `json:"IsVip,omitempty" name:"IsVip"`

	// 是否为泛域名证书。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsWildcard *bool `json:"IsWildcard,omitempty" name:"IsWildcard"`

	// 是否为 DV 版证书。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDv *bool `json:"IsDv,omitempty" name:"IsDv"`

	// 是否启用了漏洞扫描功能。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsVulnerability *bool `json:"IsVulnerability,omitempty" name:"IsVulnerability"`

	// 提交的资料信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubmittedData *SubmittedData `json:"SubmittedData,omitempty" name:"SubmittedData"`

	// 是否可续费。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewAble *bool `json:"RenewAble,omitempty" name:"RenewAble"`

	// 是否可部署。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Deployable *bool `json:"Deployable,omitempty" name:"Deployable"`

	// 关联标签列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tags `json:"Tags,omitempty" name:"Tags"`

	// 根证书。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RootCert *RootCertificates `json:"RootCert,omitempty" name:"RootCert"`

	// 国密加密证书
	// 注意：此字段可能返回 null，表示取不到有效值。
	EncryptCert *string `json:"EncryptCert,omitempty" name:"EncryptCert"`

	// 国密加密私钥
	// 注意：此字段可能返回 null，表示取不到有效值。
	EncryptPrivateKey *string `json:"EncryptPrivateKey,omitempty" name:"EncryptPrivateKey"`

	// 签名证书 SHA1指纹
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertFingerprint *string `json:"CertFingerprint,omitempty" name:"CertFingerprint"`

	// 加密证书 SHA1指纹 （国密证书特有）
	// 注意：此字段可能返回 null，表示取不到有效值。
	EncryptCertFingerprint *string `json:"EncryptCertFingerprint,omitempty" name:"EncryptCertFingerprint"`

	// 证书算法
	// 注意：此字段可能返回 null，表示取不到有效值。
	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitempty" name:"EncryptAlgorithm"`

	// DV证书吊销验证值
	// 注意：此字段可能返回 null，表示取不到有效值。
	DvRevokeAuthDetail []*DvAuths `json:"DvRevokeAuthDetail,omitempty" name:"DvRevokeAuthDetail"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCertificateDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCertificateDetailResponseParams `json:"Response"`
}

func (r *DescribeCertificateDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificateDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertificateOperateLogsRequestParams struct {
	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 请求日志数量，默认为20。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 开始时间，默认15天前。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，默认现在时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeCertificateOperateLogsRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 请求日志数量，默认为20。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 开始时间，默认15天前。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，默认现在时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeCertificateOperateLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificateOperateLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCertificateOperateLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertificateOperateLogsResponseParams struct {
	// 当前查询条件日志总数。
	AllTotal *uint64 `json:"AllTotal,omitempty" name:"AllTotal"`

	// 本次请求返回的日志数量。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 证书操作日志列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperateLogs []*OperationLog `json:"OperateLogs,omitempty" name:"OperateLogs"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCertificateOperateLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCertificateOperateLogsResponseParams `json:"Response"`
}

func (r *DescribeCertificateOperateLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificateOperateLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertificateRequestParams struct {
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

type DescribeCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

func (r *DescribeCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertificateResponseParams struct {
	// 用户 UIN。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 项目 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 证书来源：trustasia = 亚洲诚信，upload = 用户上传。
	// 注意：此字段可能返回 null，表示取不到有效值。
	From *string `json:"From,omitempty" name:"From"`

	// 证书类型：CA = 客户端证书，SVR = 服务器证书。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertificateType *string `json:"CertificateType,omitempty" name:"CertificateType"`

	// 证书套餐类型：1 = GeoTrust DV SSL CA - G3， 2 = TrustAsia TLS RSA CA， 3 = SecureSite 增强型企业版（EV Pro）， 4 = SecureSite 增强型（EV）， 5 = SecureSite 企业型专业版（OV Pro）， 6 = SecureSite 企业型（OV）， 7 = SecureSite 企业型（OV）通配符， 8 = Geotrust 增强型（EV）， 9 = Geotrust 企业型（OV）， 10 = Geotrust 企业型（OV）通配符， 11 = TrustAsia 域名型多域名 SSL 证书， 12 = TrustAsia 域名型（DV）通配符， 13 = TrustAsia 企业型通配符（OV）SSL 证书（D3）， 14 = TrustAsia 企业型（OV）SSL 证书（D3）， 15 = TrustAsia 企业型多域名 （OV）SSL 证书（D3）， 16 = TrustAsia 增强型 （EV）SSL 证书（D3）， 17 = TrustAsia 增强型多域名（EV）SSL 证书（D3）， 18 = GlobalSign 企业型（OV）SSL 证书， 19 = GlobalSign 企业型通配符 （OV）SSL 证书， 20 = GlobalSign 增强型 （EV）SSL 证书， 21 = TrustAsia 企业型通配符多域名（OV）SSL 证书（D3）， 22 = GlobalSign 企业型多域名（OV）SSL 证书， 23 = GlobalSign 企业型通配符多域名（OV）SSL 证书， 24 = GlobalSign 增强型多域名（EV）SSL 证书。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageType *string `json:"PackageType,omitempty" name:"PackageType"`

	// 证书颁发者名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductZhName *string `json:"ProductZhName,omitempty" name:"ProductZhName"`

	// 域名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 备注名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 证书状态：0 = 审核中，1 = 已通过，2 = 审核失败，3 = 已过期，4 = 已添加DNS记录，5 = 企业证书，待提交，6 = 订单取消中，7 = 已取消，8 = 已提交资料， 待上传确认函，9 = 证书吊销中，10 = 已吊销，11 = 重颁发中，12 = 待上传吊销确认函。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 状态信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusMsg *string `json:"StatusMsg,omitempty" name:"StatusMsg"`

	// 验证类型：DNS_AUTO = 自动DNS验证，DNS = 手动DNS验证，FILE = 文件验证，EMAIL = 邮件验证。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerifyType *string `json:"VerifyType,omitempty" name:"VerifyType"`

	// 漏洞扫描状态。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulnerabilityStatus *string `json:"VulnerabilityStatus,omitempty" name:"VulnerabilityStatus"`

	// 证书生效时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertBeginTime *string `json:"CertBeginTime,omitempty" name:"CertBeginTime"`

	// 证书失效时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertEndTime *string `json:"CertEndTime,omitempty" name:"CertEndTime"`

	// 证书有效期：单位(月)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValidityPeriod *string `json:"ValidityPeriod,omitempty" name:"ValidityPeriod"`

	// 申请时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`

	// 订单 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 证书扩展信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertificateExtra *CertificateExtra `json:"CertificateExtra,omitempty" name:"CertificateExtra"`

	// DV 认证信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DvAuthDetail *DvAuthDetail `json:"DvAuthDetail,omitempty" name:"DvAuthDetail"`

	// 漏洞扫描评估报告。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulnerabilityReport *string `json:"VulnerabilityReport,omitempty" name:"VulnerabilityReport"`

	// 证书 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 证书类型名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageTypeName *string `json:"PackageTypeName,omitempty" name:"PackageTypeName"`

	// 状态描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusName *string `json:"StatusName,omitempty" name:"StatusName"`

	// 证书包含的多个域名（包含主域名）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubjectAltName []*string `json:"SubjectAltName,omitempty" name:"SubjectAltName"`

	// 是否为 VIP 客户。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsVip *bool `json:"IsVip,omitempty" name:"IsVip"`

	// 是否为泛域名证书。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsWildcard *bool `json:"IsWildcard,omitempty" name:"IsWildcard"`

	// 是否为 DV 版证书。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDv *bool `json:"IsDv,omitempty" name:"IsDv"`

	// 是否启用了漏洞扫描功能。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsVulnerability *bool `json:"IsVulnerability,omitempty" name:"IsVulnerability"`

	// 是否可重颁发证书。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewAble *bool `json:"RenewAble,omitempty" name:"RenewAble"`

	// 提交的资料信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubmittedData *SubmittedData `json:"SubmittedData,omitempty" name:"SubmittedData"`

	// 是否可部署。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Deployable *bool `json:"Deployable,omitempty" name:"Deployable"`

	// 标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tags `json:"Tags,omitempty" name:"Tags"`

	// CA证书的所有加密方式	
	// 注意：此字段可能返回 null，表示取不到有效值。
	CAEncryptAlgorithms []*string `json:"CAEncryptAlgorithms,omitempty" name:"CAEncryptAlgorithms"`

	// CA证书的所有通用名称	
	// 注意：此字段可能返回 null，表示取不到有效值。
	CACommonNames []*string `json:"CACommonNames,omitempty" name:"CACommonNames"`

	// CA证书所有的到期时间	
	// 注意：此字段可能返回 null，表示取不到有效值。
	CAEndTimes []*string `json:"CAEndTimes,omitempty" name:"CAEndTimes"`

	// DV证书吊销验证值
	// 注意：此字段可能返回 null，表示取不到有效值。
	DvRevokeAuthDetail []*DvAuths `json:"DvRevokeAuthDetail,omitempty" name:"DvRevokeAuthDetail"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCertificateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCertificateResponseParams `json:"Response"`
}

func (r *DescribeCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertificatesRequestParams struct {
	// 分页偏移量，从0开始。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页数量，默认20。最大1000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索关键词，可搜索证书 ID、备注名称、域名。例如： a8xHcaIs。
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// 证书类型：CA = 客户端证书，SVR = 服务器证书。
	CertificateType *string `json:"CertificateType,omitempty" name:"CertificateType"`

	// 项目 ID。
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 按到期时间排序：DESC = 降序， ASC = 升序。
	ExpirationSort *string `json:"ExpirationSort,omitempty" name:"ExpirationSort"`

	// 证书状态：0 = 审核中，1 = 已通过，2 = 审核失败，3 = 已过期，4 = 已添加DNS记录，5 = 企业证书，待提交，6 = 订单取消中，7 = 已取消，8 = 已提交资料， 待上传确认函，9 = 证书吊销中，10 = 已吊销，11 = 重颁发中，12 = 待上传吊销确认函，13 = 免费证书待提交资料。
	CertificateStatus []*uint64 `json:"CertificateStatus,omitempty" name:"CertificateStatus"`

	// 是否可部署，可选值：1 = 可部署，0 =  不可部署。
	Deployable *uint64 `json:"Deployable,omitempty" name:"Deployable"`

	// 是否筛选上传托管的 1筛选，0不筛选
	Upload *int64 `json:"Upload,omitempty" name:"Upload"`

	// 是否筛选可续期证书 1筛选 0不筛选
	Renew *int64 `json:"Renew,omitempty" name:"Renew"`

	// 筛选来源， upload：上传证书， buy：腾讯云证书， 不传默认全部
	FilterSource *string `json:"FilterSource,omitempty" name:"FilterSource"`

	// 是否筛选国密证书。1:筛选  0:不筛选
	IsSM *int64 `json:"IsSM,omitempty" name:"IsSM"`

	// 筛选证书是否即将过期，传1是筛选，0不筛选
	FilterExpiring *uint64 `json:"FilterExpiring,omitempty" name:"FilterExpiring"`

	// 是否可托管，可选值：1 = 可托管，0 =  不可托管。
	Hostable *uint64 `json:"Hostable,omitempty" name:"Hostable"`
}

type DescribeCertificatesRequest struct {
	*tchttp.BaseRequest
	
	// 分页偏移量，从0开始。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页数量，默认20。最大1000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索关键词，可搜索证书 ID、备注名称、域名。例如： a8xHcaIs。
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// 证书类型：CA = 客户端证书，SVR = 服务器证书。
	CertificateType *string `json:"CertificateType,omitempty" name:"CertificateType"`

	// 项目 ID。
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 按到期时间排序：DESC = 降序， ASC = 升序。
	ExpirationSort *string `json:"ExpirationSort,omitempty" name:"ExpirationSort"`

	// 证书状态：0 = 审核中，1 = 已通过，2 = 审核失败，3 = 已过期，4 = 已添加DNS记录，5 = 企业证书，待提交，6 = 订单取消中，7 = 已取消，8 = 已提交资料， 待上传确认函，9 = 证书吊销中，10 = 已吊销，11 = 重颁发中，12 = 待上传吊销确认函，13 = 免费证书待提交资料。
	CertificateStatus []*uint64 `json:"CertificateStatus,omitempty" name:"CertificateStatus"`

	// 是否可部署，可选值：1 = 可部署，0 =  不可部署。
	Deployable *uint64 `json:"Deployable,omitempty" name:"Deployable"`

	// 是否筛选上传托管的 1筛选，0不筛选
	Upload *int64 `json:"Upload,omitempty" name:"Upload"`

	// 是否筛选可续期证书 1筛选 0不筛选
	Renew *int64 `json:"Renew,omitempty" name:"Renew"`

	// 筛选来源， upload：上传证书， buy：腾讯云证书， 不传默认全部
	FilterSource *string `json:"FilterSource,omitempty" name:"FilterSource"`

	// 是否筛选国密证书。1:筛选  0:不筛选
	IsSM *int64 `json:"IsSM,omitempty" name:"IsSM"`

	// 筛选证书是否即将过期，传1是筛选，0不筛选
	FilterExpiring *uint64 `json:"FilterExpiring,omitempty" name:"FilterExpiring"`

	// 是否可托管，可选值：1 = 可托管，0 =  不可托管。
	Hostable *uint64 `json:"Hostable,omitempty" name:"Hostable"`
}

func (r *DescribeCertificatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchKey")
	delete(f, "CertificateType")
	delete(f, "ProjectId")
	delete(f, "ExpirationSort")
	delete(f, "CertificateStatus")
	delete(f, "Deployable")
	delete(f, "Upload")
	delete(f, "Renew")
	delete(f, "FilterSource")
	delete(f, "IsSM")
	delete(f, "FilterExpiring")
	delete(f, "Hostable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCertificatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertificatesResponseParams struct {
	// 总数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Certificates []*Certificates `json:"Certificates,omitempty" name:"Certificates"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCertificatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCertificatesResponseParams `json:"Response"`
}

func (r *DescribeCertificatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCompaniesRequestParams struct {
	// 分页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页每页限制数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 公司ID
	CompanyId *int64 `json:"CompanyId,omitempty" name:"CompanyId"`
}

type DescribeCompaniesRequest struct {
	*tchttp.BaseRequest
	
	// 分页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页每页限制数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 公司ID
	CompanyId *int64 `json:"CompanyId,omitempty" name:"CompanyId"`
}

func (r *DescribeCompaniesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCompaniesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "CompanyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCompaniesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCompaniesResponseParams struct {
	// 公司列表
	Companies []*CompanyInfo `json:"Companies,omitempty" name:"Companies"`

	// 公司总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCompaniesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCompaniesResponseParams `json:"Response"`
}

func (r *DescribeCompaniesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCompaniesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeployedResourcesRequestParams struct {
	// 证书ID
	CertificateIds []*string `json:"CertificateIds,omitempty" name:"CertificateIds"`

	// 资源类型:clb,cdn,live,waf,antiddos
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
}

type DescribeDeployedResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 证书ID
	CertificateIds []*string `json:"CertificateIds,omitempty" name:"CertificateIds"`

	// 资源类型:clb,cdn,live,waf,antiddos
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
}

func (r *DescribeDeployedResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeployedResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateIds")
	delete(f, "ResourceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeployedResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeployedResourcesResponseParams struct {
	// 资源详情
	DeployedResources []*DeployedResources `json:"DeployedResources,omitempty" name:"DeployedResources"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDeployedResourcesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeployedResourcesResponseParams `json:"Response"`
}

func (r *DescribeDeployedResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeployedResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostApiGatewayInstanceListRequestParams struct {
	// 待部署的证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 部署资源类型
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitempty" name:"IsCache"`

	// 过滤参数列表； FilterKey：domainMatch（查询域名是否匹配的实例列表） FilterValue：1，表示查询匹配； 0，表示查询不匹配； 默认查询匹配
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 已部署的证书ID
	OldCertificateId *string `json:"OldCertificateId,omitempty" name:"OldCertificateId"`
}

type DescribeHostApiGatewayInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 待部署的证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 部署资源类型
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitempty" name:"IsCache"`

	// 过滤参数列表； FilterKey：domainMatch（查询域名是否匹配的实例列表） FilterValue：1，表示查询匹配； 0，表示查询不匹配； 默认查询匹配
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 已部署的证书ID
	OldCertificateId *string `json:"OldCertificateId,omitempty" name:"OldCertificateId"`
}

func (r *DescribeHostApiGatewayInstanceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostApiGatewayInstanceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	delete(f, "ResourceType")
	delete(f, "IsCache")
	delete(f, "Filters")
	delete(f, "OldCertificateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostApiGatewayInstanceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostApiGatewayInstanceListResponseParams struct {
	// apiGateway实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceList []*ApiGatewayInstanceDetail `json:"InstanceList,omitempty" name:"InstanceList"`

	// 总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeHostApiGatewayInstanceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHostApiGatewayInstanceListResponseParams `json:"Response"`
}

func (r *DescribeHostApiGatewayInstanceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostApiGatewayInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostCdnInstanceListRequestParams struct {
	// 待部署的证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 部署资源类型
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitempty" name:"IsCache"`

	// 过滤参数列表； FilterKey：domainMatch（查询域名是否匹配的实例列表） FilterValue：1，表示查询匹配； 0，表示查询不匹配； 默认查询匹配
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 原证书ID
	OldCertificateId *string `json:"OldCertificateId,omitempty" name:"OldCertificateId"`

	// 分页偏移量，从0开始。	
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页数量，默认10。	
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 是否异步
	AsyncCache *int64 `json:"AsyncCache,omitempty" name:"AsyncCache"`
}

type DescribeHostCdnInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 待部署的证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 部署资源类型
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitempty" name:"IsCache"`

	// 过滤参数列表； FilterKey：domainMatch（查询域名是否匹配的实例列表） FilterValue：1，表示查询匹配； 0，表示查询不匹配； 默认查询匹配
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 原证书ID
	OldCertificateId *string `json:"OldCertificateId,omitempty" name:"OldCertificateId"`

	// 分页偏移量，从0开始。	
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页数量，默认10。	
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 是否异步
	AsyncCache *int64 `json:"AsyncCache,omitempty" name:"AsyncCache"`
}

func (r *DescribeHostCdnInstanceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostCdnInstanceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	delete(f, "ResourceType")
	delete(f, "IsCache")
	delete(f, "Filters")
	delete(f, "OldCertificateId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "AsyncCache")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostCdnInstanceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostCdnInstanceListResponseParams struct {
	// CDN实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceList []*CdnInstanceDetail `json:"InstanceList,omitempty" name:"InstanceList"`

	// CDN域名总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 异步刷新总数	
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsyncTotalNum *int64 `json:"AsyncTotalNum,omitempty" name:"AsyncTotalNum"`

	// 异步刷新当前执行数	
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsyncOffset *int64 `json:"AsyncOffset,omitempty" name:"AsyncOffset"`

	// 当前缓存读取时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsyncCacheTime *string `json:"AsyncCacheTime,omitempty" name:"AsyncCacheTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeHostCdnInstanceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHostCdnInstanceListResponseParams `json:"Response"`
}

func (r *DescribeHostCdnInstanceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostCdnInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostClbInstanceListRequestParams struct {
	// 待部署的证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 分页偏移量，从0开始。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页数量，默认10。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitempty" name:"IsCache"`

	// 过滤参数列表； FilterKey：domainMatch（查询域名是否匹配的实例列表） FilterValue：1，表示查询匹配； 0，表示查询不匹配； 默认查询匹配
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 是否异步缓存
	AsyncCache *int64 `json:"AsyncCache,omitempty" name:"AsyncCache"`

	// 原证书ID
	OldCertificateId *string `json:"OldCertificateId,omitempty" name:"OldCertificateId"`
}

type DescribeHostClbInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 待部署的证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 分页偏移量，从0开始。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页数量，默认10。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitempty" name:"IsCache"`

	// 过滤参数列表； FilterKey：domainMatch（查询域名是否匹配的实例列表） FilterValue：1，表示查询匹配； 0，表示查询不匹配； 默认查询匹配
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 是否异步缓存
	AsyncCache *int64 `json:"AsyncCache,omitempty" name:"AsyncCache"`

	// 原证书ID
	OldCertificateId *string `json:"OldCertificateId,omitempty" name:"OldCertificateId"`
}

func (r *DescribeHostClbInstanceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostClbInstanceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "IsCache")
	delete(f, "Filters")
	delete(f, "AsyncCache")
	delete(f, "OldCertificateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostClbInstanceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostClbInstanceListResponseParams struct {
	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// CLB实例监听器列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceList []*ClbInstanceDetail `json:"InstanceList,omitempty" name:"InstanceList"`

	// 异步刷新总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsyncTotalNum *int64 `json:"AsyncTotalNum,omitempty" name:"AsyncTotalNum"`

	// 异步刷新当前执行数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsyncOffset *int64 `json:"AsyncOffset,omitempty" name:"AsyncOffset"`

	// 当前缓存读取时间	
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsyncCacheTime *string `json:"AsyncCacheTime,omitempty" name:"AsyncCacheTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeHostClbInstanceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHostClbInstanceListResponseParams `json:"Response"`
}

func (r *DescribeHostClbInstanceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostClbInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostCosInstanceListRequestParams struct {
	// 待部署的证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 部署资源类型 cos
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitempty" name:"IsCache"`

	// 过滤参数列表
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeHostCosInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 待部署的证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 部署资源类型 cos
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitempty" name:"IsCache"`

	// 过滤参数列表
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeHostCosInstanceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostCosInstanceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	delete(f, "ResourceType")
	delete(f, "IsCache")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostCosInstanceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostCosInstanceListResponseParams struct {
	// COS实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceList []*CosInstanceDetail `json:"InstanceList,omitempty" name:"InstanceList"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 异步刷新总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsyncTotalNum *int64 `json:"AsyncTotalNum,omitempty" name:"AsyncTotalNum"`

	// 异步刷新当前执行数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsyncOffset *int64 `json:"AsyncOffset,omitempty" name:"AsyncOffset"`

	// 当前缓存读取时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsyncCacheTime *string `json:"AsyncCacheTime,omitempty" name:"AsyncCacheTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeHostCosInstanceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHostCosInstanceListResponseParams `json:"Response"`
}

func (r *DescribeHostCosInstanceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostCosInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostDdosInstanceListRequestParams struct {
	// 待部署的证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 部署资源类型
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitempty" name:"IsCache"`

	// 过滤参数列表； FilterKey：domainMatch（查询域名是否匹配的实例列表） FilterValue：1，表示查询匹配； 0，表示查询不匹配； 默认查询匹配
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 已部署的证书ID
	OldCertificateId *string `json:"OldCertificateId,omitempty" name:"OldCertificateId"`
}

type DescribeHostDdosInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 待部署的证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 部署资源类型
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitempty" name:"IsCache"`

	// 过滤参数列表； FilterKey：domainMatch（查询域名是否匹配的实例列表） FilterValue：1，表示查询匹配； 0，表示查询不匹配； 默认查询匹配
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 已部署的证书ID
	OldCertificateId *string `json:"OldCertificateId,omitempty" name:"OldCertificateId"`
}

func (r *DescribeHostDdosInstanceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostDdosInstanceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	delete(f, "ResourceType")
	delete(f, "IsCache")
	delete(f, "Filters")
	delete(f, "OldCertificateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostDdosInstanceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostDdosInstanceListResponseParams struct {
	// DDOS实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceList []*DdosInstanceDetail `json:"InstanceList,omitempty" name:"InstanceList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeHostDdosInstanceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHostDdosInstanceListResponseParams `json:"Response"`
}

func (r *DescribeHostDdosInstanceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostDdosInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostDeployRecordDetailRequestParams struct {
	// 部署记录ID
	DeployRecordId *string `json:"DeployRecordId,omitempty" name:"DeployRecordId"`

	// 分页偏移量，从0开始。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页数量，默认10。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeHostDeployRecordDetailRequest struct {
	*tchttp.BaseRequest
	
	// 部署记录ID
	DeployRecordId *string `json:"DeployRecordId,omitempty" name:"DeployRecordId"`

	// 分页偏移量，从0开始。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页数量，默认10。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeHostDeployRecordDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostDeployRecordDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeployRecordId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostDeployRecordDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostDeployRecordDetailResponseParams struct {
	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 证书部署记录列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployRecordDetailList []*DeployRecordDetail `json:"DeployRecordDetailList,omitempty" name:"DeployRecordDetailList"`

	// 成功总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessTotalCount *int64 `json:"SuccessTotalCount,omitempty" name:"SuccessTotalCount"`

	// 失败总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedTotalCount *int64 `json:"FailedTotalCount,omitempty" name:"FailedTotalCount"`

	// 部署中总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunningTotalCount *int64 `json:"RunningTotalCount,omitempty" name:"RunningTotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeHostDeployRecordDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHostDeployRecordDetailResponseParams `json:"Response"`
}

func (r *DescribeHostDeployRecordDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostDeployRecordDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostDeployRecordRequestParams struct {
	// 待部署的证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 分页偏移量，从0开始。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页数量，默认10。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 资源类型
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
}

type DescribeHostDeployRecordRequest struct {
	*tchttp.BaseRequest
	
	// 待部署的证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 分页偏移量，从0开始。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页数量，默认10。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 资源类型
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
}

func (r *DescribeHostDeployRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostDeployRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ResourceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostDeployRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostDeployRecordResponseParams struct {
	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 证书部署记录列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployRecordList []*DeployRecordInfo `json:"DeployRecordList,omitempty" name:"DeployRecordList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeHostDeployRecordResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHostDeployRecordResponseParams `json:"Response"`
}

func (r *DescribeHostDeployRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostDeployRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostLighthouseInstanceListRequestParams struct {
	// 待部署的证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 部署资源类型 lighthouse
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitempty" name:"IsCache"`

	// 过滤参数列表
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeHostLighthouseInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 待部署的证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 部署资源类型 lighthouse
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitempty" name:"IsCache"`

	// 过滤参数列表
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeHostLighthouseInstanceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostLighthouseInstanceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	delete(f, "ResourceType")
	delete(f, "IsCache")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostLighthouseInstanceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostLighthouseInstanceListResponseParams struct {
	// Lighthouse实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceList []*LighthouseInstanceDetail `json:"InstanceList,omitempty" name:"InstanceList"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeHostLighthouseInstanceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHostLighthouseInstanceListResponseParams `json:"Response"`
}

func (r *DescribeHostLighthouseInstanceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostLighthouseInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostLiveInstanceListRequestParams struct {
	// 待部署的证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 部署资源类型
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitempty" name:"IsCache"`

	// 过滤参数列表； FilterKey：domainMatch（查询域名是否匹配的实例列表） FilterValue：1，表示查询匹配； 0，表示查询不匹配； 默认查询匹配
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 已部署的证书ID
	OldCertificateId *string `json:"OldCertificateId,omitempty" name:"OldCertificateId"`
}

type DescribeHostLiveInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 待部署的证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 部署资源类型
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitempty" name:"IsCache"`

	// 过滤参数列表； FilterKey：domainMatch（查询域名是否匹配的实例列表） FilterValue：1，表示查询匹配； 0，表示查询不匹配； 默认查询匹配
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 已部署的证书ID
	OldCertificateId *string `json:"OldCertificateId,omitempty" name:"OldCertificateId"`
}

func (r *DescribeHostLiveInstanceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostLiveInstanceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	delete(f, "ResourceType")
	delete(f, "IsCache")
	delete(f, "Filters")
	delete(f, "OldCertificateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostLiveInstanceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostLiveInstanceListResponseParams struct {
	// live实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceList []*LiveInstanceDetail `json:"InstanceList,omitempty" name:"InstanceList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeHostLiveInstanceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHostLiveInstanceListResponseParams `json:"Response"`
}

func (r *DescribeHostLiveInstanceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostLiveInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostTeoInstanceListRequestParams struct {
	// 待部署的证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 部署资源类型
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitempty" name:"IsCache"`

	// 过滤参数列表； FilterKey：domainMatch（查询域名是否匹配的实例列表） FilterValue：1，表示查询匹配； 0，表示查询不匹配； 默认查询匹配
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 已部署的证书ID
	OldCertificateId *string `json:"OldCertificateId,omitempty" name:"OldCertificateId"`
}

type DescribeHostTeoInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 待部署的证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 部署资源类型
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitempty" name:"IsCache"`

	// 过滤参数列表； FilterKey：domainMatch（查询域名是否匹配的实例列表） FilterValue：1，表示查询匹配； 0，表示查询不匹配； 默认查询匹配
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 已部署的证书ID
	OldCertificateId *string `json:"OldCertificateId,omitempty" name:"OldCertificateId"`
}

func (r *DescribeHostTeoInstanceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostTeoInstanceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	delete(f, "ResourceType")
	delete(f, "IsCache")
	delete(f, "Filters")
	delete(f, "OldCertificateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostTeoInstanceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostTeoInstanceListResponseParams struct {
	// teo实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceList []*TeoInstanceDetail `json:"InstanceList,omitempty" name:"InstanceList"`

	// 总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeHostTeoInstanceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHostTeoInstanceListResponseParams `json:"Response"`
}

func (r *DescribeHostTeoInstanceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostTeoInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostTkeInstanceListRequestParams struct {
	// 待部署的证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 分页偏移量，从0开始。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页数量，默认10。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitempty" name:"IsCache"`

	// 过滤参数列表； FilterKey：domainMatch（查询域名是否匹配的实例列表） FilterValue：1，表示查询匹配； 0，表示查询不匹配； 默认查询匹配
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 是否异步缓存
	AsyncCache *int64 `json:"AsyncCache,omitempty" name:"AsyncCache"`

	// 原证书ID
	OldCertificateId *string `json:"OldCertificateId,omitempty" name:"OldCertificateId"`
}

type DescribeHostTkeInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 待部署的证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 分页偏移量，从0开始。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页数量，默认10。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitempty" name:"IsCache"`

	// 过滤参数列表； FilterKey：domainMatch（查询域名是否匹配的实例列表） FilterValue：1，表示查询匹配； 0，表示查询不匹配； 默认查询匹配
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 是否异步缓存
	AsyncCache *int64 `json:"AsyncCache,omitempty" name:"AsyncCache"`

	// 原证书ID
	OldCertificateId *string `json:"OldCertificateId,omitempty" name:"OldCertificateId"`
}

func (r *DescribeHostTkeInstanceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostTkeInstanceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "IsCache")
	delete(f, "Filters")
	delete(f, "AsyncCache")
	delete(f, "OldCertificateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostTkeInstanceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostTkeInstanceListResponseParams struct {
	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// CLB实例监听器列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceList []*TkeInstanceDetail `json:"InstanceList,omitempty" name:"InstanceList"`

	// 异步刷新总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsyncTotalNum *int64 `json:"AsyncTotalNum,omitempty" name:"AsyncTotalNum"`

	// 异步刷新当前执行数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsyncOffset *int64 `json:"AsyncOffset,omitempty" name:"AsyncOffset"`

	// 当前缓存读取时间	
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsyncCacheTime *string `json:"AsyncCacheTime,omitempty" name:"AsyncCacheTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeHostTkeInstanceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHostTkeInstanceListResponseParams `json:"Response"`
}

func (r *DescribeHostTkeInstanceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostTkeInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostUpdateRecordDetailRequestParams struct {
	// 一键更新记录ID
	DeployRecordId *string `json:"DeployRecordId,omitempty" name:"DeployRecordId"`

	// 每页数量，默认10。
	Limit *string `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量，从0开始。
	Offset *string `json:"Offset,omitempty" name:"Offset"`
}

type DescribeHostUpdateRecordDetailRequest struct {
	*tchttp.BaseRequest
	
	// 一键更新记录ID
	DeployRecordId *string `json:"DeployRecordId,omitempty" name:"DeployRecordId"`

	// 每页数量，默认10。
	Limit *string `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量，从0开始。
	Offset *string `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeHostUpdateRecordDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostUpdateRecordDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeployRecordId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostUpdateRecordDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostUpdateRecordDetailResponseParams struct {
	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 证书部署记录列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordDetailList []*UpdateRecordDetails `json:"RecordDetailList,omitempty" name:"RecordDetailList"`

	// 成功总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessTotalCount *int64 `json:"SuccessTotalCount,omitempty" name:"SuccessTotalCount"`

	// 失败总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedTotalCount *int64 `json:"FailedTotalCount,omitempty" name:"FailedTotalCount"`

	// 部署中总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunningTotalCount *int64 `json:"RunningTotalCount,omitempty" name:"RunningTotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeHostUpdateRecordDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHostUpdateRecordDetailResponseParams `json:"Response"`
}

func (r *DescribeHostUpdateRecordDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostUpdateRecordDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostUpdateRecordRequestParams struct {
	// 分页偏移量，从0开始。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页数量，默认10。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 新证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 原证书ID
	OldCertificateId *string `json:"OldCertificateId,omitempty" name:"OldCertificateId"`
}

type DescribeHostUpdateRecordRequest struct {
	*tchttp.BaseRequest
	
	// 分页偏移量，从0开始。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页数量，默认10。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 新证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 原证书ID
	OldCertificateId *string `json:"OldCertificateId,omitempty" name:"OldCertificateId"`
}

func (r *DescribeHostUpdateRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostUpdateRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "CertificateId")
	delete(f, "OldCertificateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostUpdateRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostUpdateRecordResponseParams struct {
	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 证书部署记录列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployRecordList []*UpdateRecordInfo `json:"DeployRecordList,omitempty" name:"DeployRecordList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeHostUpdateRecordResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHostUpdateRecordResponseParams `json:"Response"`
}

func (r *DescribeHostUpdateRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostUpdateRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostVodInstanceListRequestParams struct {
	// 待部署的证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 部署资源类型 vod
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitempty" name:"IsCache"`

	// 过滤参数列表
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 已部署的证书ID
	OldCertificateId *string `json:"OldCertificateId,omitempty" name:"OldCertificateId"`
}

type DescribeHostVodInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 待部署的证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 部署资源类型 vod
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitempty" name:"IsCache"`

	// 过滤参数列表
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 已部署的证书ID
	OldCertificateId *string `json:"OldCertificateId,omitempty" name:"OldCertificateId"`
}

func (r *DescribeHostVodInstanceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostVodInstanceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	delete(f, "ResourceType")
	delete(f, "IsCache")
	delete(f, "Filters")
	delete(f, "OldCertificateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostVodInstanceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostVodInstanceListResponseParams struct {
	// Vod实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceList []*VodInstanceDetail `json:"InstanceList,omitempty" name:"InstanceList"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeHostVodInstanceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHostVodInstanceListResponseParams `json:"Response"`
}

func (r *DescribeHostVodInstanceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostVodInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostWafInstanceListRequestParams struct {
	// 待部署的证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 部署资源类型
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitempty" name:"IsCache"`

	// 过滤参数列表； FilterKey：domainMatch（查询域名是否匹配的实例列表） FilterValue：1，表示查询匹配； 0，表示查询不匹配； 默认查询匹配
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 已部署的证书ID
	OldCertificateId *string `json:"OldCertificateId,omitempty" name:"OldCertificateId"`
}

type DescribeHostWafInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 待部署的证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 部署资源类型
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitempty" name:"IsCache"`

	// 过滤参数列表； FilterKey：domainMatch（查询域名是否匹配的实例列表） FilterValue：1，表示查询匹配； 0，表示查询不匹配； 默认查询匹配
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 已部署的证书ID
	OldCertificateId *string `json:"OldCertificateId,omitempty" name:"OldCertificateId"`
}

func (r *DescribeHostWafInstanceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostWafInstanceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	delete(f, "ResourceType")
	delete(f, "IsCache")
	delete(f, "Filters")
	delete(f, "OldCertificateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostWafInstanceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostWafInstanceListResponseParams struct {
	// WAF实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceList []*LiveInstanceDetail `json:"InstanceList,omitempty" name:"InstanceList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeHostWafInstanceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHostWafInstanceListResponseParams `json:"Response"`
}

func (r *DescribeHostWafInstanceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostWafInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeManagerDetailRequestParams struct {
	// 管理人ID
	ManagerId *int64 `json:"ManagerId,omitempty" name:"ManagerId"`

	// 分页每页数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeManagerDetailRequest struct {
	*tchttp.BaseRequest
	
	// 管理人ID
	ManagerId *int64 `json:"ManagerId,omitempty" name:"ManagerId"`

	// 分页每页数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeManagerDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeManagerDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ManagerId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeManagerDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeManagerDetailResponseParams struct {
	// 状态: audit: 审核中 ok: 审核通过 invalid: 失效 expiring: 即将过期 expired: 已过期
	Status *string `json:"Status,omitempty" name:"Status"`

	// 管理人姓名
	ManagerFirstName *string `json:"ManagerFirstName,omitempty" name:"ManagerFirstName"`

	// 管理人邮箱
	ManagerMail *string `json:"ManagerMail,omitempty" name:"ManagerMail"`

	// 联系人姓名
	ContactFirstName *string `json:"ContactFirstName,omitempty" name:"ContactFirstName"`

	// 管理人姓名
	ManagerLastName *string `json:"ManagerLastName,omitempty" name:"ManagerLastName"`

	// 联系人职位
	ContactPosition *string `json:"ContactPosition,omitempty" name:"ContactPosition"`

	// 管理人职位
	ManagerPosition *string `json:"ManagerPosition,omitempty" name:"ManagerPosition"`

	// 核验通过时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerifyTime *string `json:"VerifyTime,omitempty" name:"VerifyTime"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 核验过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 联系人姓名
	ContactLastName *string `json:"ContactLastName,omitempty" name:"ContactLastName"`

	// 管理人电话
	ManagerPhone *string `json:"ManagerPhone,omitempty" name:"ManagerPhone"`

	// 联系人电话
	ContactPhone *string `json:"ContactPhone,omitempty" name:"ContactPhone"`

	// 联系人邮箱
	ContactMail *string `json:"ContactMail,omitempty" name:"ContactMail"`

	// 管理人所属部门
	ManagerDepartment *string `json:"ManagerDepartment,omitempty" name:"ManagerDepartment"`

	// 管理人所属公司信息
	CompanyInfo *CompanyInfo `json:"CompanyInfo,omitempty" name:"CompanyInfo"`

	// 管理人公司ID
	CompanyId *int64 `json:"CompanyId,omitempty" name:"CompanyId"`

	// 管理人ID
	ManagerId *int64 `json:"ManagerId,omitempty" name:"ManagerId"`

	// 审核状态详细信息
	StatusInfo []*ManagerStatusInfo `json:"StatusInfo,omitempty" name:"StatusInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeManagerDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeManagerDetailResponseParams `json:"Response"`
}

func (r *DescribeManagerDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeManagerDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeManagersRequestParams struct {
	// 公司ID
	CompanyId *int64 `json:"CompanyId,omitempty" name:"CompanyId"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页每页数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 管理人姓名（将废弃），请使用SearchKey
	ManagerName *string `json:"ManagerName,omitempty" name:"ManagerName"`

	// 模糊查询管理人邮箱（将废弃），请使用SearchKey
	ManagerMail *string `json:"ManagerMail,omitempty" name:"ManagerMail"`

	// 根据管理人状态进行筛选，取值有
	// 'none' 未提交审核
	// 'audit', 亚信审核中
	// 'CAaudit' CA审核中
	// 'ok' 已审核
	// 'invalid'  审核失败
	// 'expiring'  即将过期
	// 'expired' 已过期
	Status *string `json:"Status,omitempty" name:"Status"`

	// 管理人姓/管理人名/邮箱/部门精准匹配
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
}

type DescribeManagersRequest struct {
	*tchttp.BaseRequest
	
	// 公司ID
	CompanyId *int64 `json:"CompanyId,omitempty" name:"CompanyId"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页每页数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 管理人姓名（将废弃），请使用SearchKey
	ManagerName *string `json:"ManagerName,omitempty" name:"ManagerName"`

	// 模糊查询管理人邮箱（将废弃），请使用SearchKey
	ManagerMail *string `json:"ManagerMail,omitempty" name:"ManagerMail"`

	// 根据管理人状态进行筛选，取值有
	// 'none' 未提交审核
	// 'audit', 亚信审核中
	// 'CAaudit' CA审核中
	// 'ok' 已审核
	// 'invalid'  审核失败
	// 'expiring'  即将过期
	// 'expired' 已过期
	Status *string `json:"Status,omitempty" name:"Status"`

	// 管理人姓/管理人名/邮箱/部门精准匹配
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
}

func (r *DescribeManagersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeManagersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CompanyId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ManagerName")
	delete(f, "ManagerMail")
	delete(f, "Status")
	delete(f, "SearchKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeManagersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeManagersResponseParams struct {
	// 公司管理人列表
	Managers []*ManagerInfo `json:"Managers,omitempty" name:"Managers"`

	// 公司管理人总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeManagersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeManagersResponseParams `json:"Response"`
}

func (r *DescribeManagersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeManagersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePackagesRequestParams struct {
	// 偏移量，默认0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目，默认20。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 按状态筛选。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 按过期时间升序或降序排列。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 按权益包ID搜索。
	PackageId *string `json:"PackageId,omitempty" name:"PackageId"`

	// 按权益包类型搜索。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 子产品编号
	Pid *int64 `json:"Pid,omitempty" name:"Pid"`
}

type DescribePackagesRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目，默认20。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 按状态筛选。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 按过期时间升序或降序排列。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 按权益包ID搜索。
	PackageId *string `json:"PackageId,omitempty" name:"PackageId"`

	// 按权益包类型搜索。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 子产品编号
	Pid *int64 `json:"Pid,omitempty" name:"Pid"`
}

func (r *DescribePackagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePackagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Status")
	delete(f, "ExpireTime")
	delete(f, "PackageId")
	delete(f, "Type")
	delete(f, "Pid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePackagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePackagesResponseParams struct {
	// 权益包列表。
	Packages []*PackageInfo `json:"Packages,omitempty" name:"Packages"`

	// 总条数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 权益点总余额。
	TotalBalance *uint64 `json:"TotalBalance,omitempty" name:"TotalBalance"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePackagesResponse struct {
	*tchttp.BaseResponse
	Response *DescribePackagesResponseParams `json:"Response"`
}

func (r *DescribePackagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePackagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadCertificateRequestParams struct {
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

type DownloadCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

func (r *DownloadCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DownloadCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadCertificateResponseParams struct {
	// ZIP base64 编码内容，base64 解码后可保存为 ZIP 文件。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitempty" name:"Content"`

	// MIME 类型：application/zip = ZIP 压缩文件。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContentType *string `json:"ContentType,omitempty" name:"ContentType"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DownloadCertificateResponse struct {
	*tchttp.BaseResponse
	Response *DownloadCertificateResponseParams `json:"Response"`
}

func (r *DownloadCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DvAuthDetail struct {
	// DV 认证密钥。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DvAuthKey *string `json:"DvAuthKey,omitempty" name:"DvAuthKey"`

	// DV 认证值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DvAuthValue *string `json:"DvAuthValue,omitempty" name:"DvAuthValue"`

	// DV 认证值域名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DvAuthDomain *string `json:"DvAuthDomain,omitempty" name:"DvAuthDomain"`

	// DV 认证值路径。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DvAuthPath *string `json:"DvAuthPath,omitempty" name:"DvAuthPath"`

	// DV 认证子域名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DvAuthKeySubDomain *string `json:"DvAuthKeySubDomain,omitempty" name:"DvAuthKeySubDomain"`

	// DV 认证信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DvAuths []*DvAuths `json:"DvAuths,omitempty" name:"DvAuths"`
}

type DvAuths struct {
	// DV 认证密钥。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DvAuthKey *string `json:"DvAuthKey,omitempty" name:"DvAuthKey"`

	// DV 认证值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DvAuthValue *string `json:"DvAuthValue,omitempty" name:"DvAuthValue"`

	// DV 认证值域名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DvAuthDomain *string `json:"DvAuthDomain,omitempty" name:"DvAuthDomain"`

	// DV 认证值路径。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DvAuthPath *string `json:"DvAuthPath,omitempty" name:"DvAuthPath"`

	// DV 认证子域名，
	// 注意：此字段可能返回 null，表示取不到有效值。
	DvAuthSubDomain *string `json:"DvAuthSubDomain,omitempty" name:"DvAuthSubDomain"`

	// DV 认证类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DvAuthVerifyType *string `json:"DvAuthVerifyType,omitempty" name:"DvAuthVerifyType"`
}

type Filter struct {
	// 过滤参数key
	FilterKey *string `json:"FilterKey,omitempty" name:"FilterKey"`

	// 过滤参数值
	FilterValue *string `json:"FilterValue,omitempty" name:"FilterValue"`
}

// Predefined struct for user
type HostCertificateRequestParams struct {
	// 证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 资源类型：目前仅限于CLB,CDN
	ResourceType []*string `json:"ResourceType,omitempty" name:"ResourceType"`
}

type HostCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 资源类型：目前仅限于CLB,CDN
	ResourceType []*string `json:"ResourceType,omitempty" name:"ResourceType"`
}

func (r *HostCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *HostCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	delete(f, "ResourceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "HostCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type HostCertificateResponseParams struct {
	// 云资源配置详情
	CertHostingInfo *CertHostingInfo `json:"CertHostingInfo,omitempty" name:"CertHostingInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type HostCertificateResponse struct {
	*tchttp.BaseResponse
	Response *HostCertificateResponseParams `json:"Response"`
}

func (r *HostCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *HostCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LighthouseInstanceDetail struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// IP地址
	IP []*string `json:"IP,omitempty" name:"IP"`

	// 可选择域名
	Domain []*string `json:"Domain,omitempty" name:"Domain"`
}

type LiveInstanceDetail struct {
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 已绑定的证书ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// -1：域名未关联证书。
	// 1： 域名https已开启。
	// 0： 域名https已关闭。
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type ManagerInfo struct {
	// 状态: audit: 审核中 ok: 审核通过 invalid: 失效 expiring: 即将过期 expired: 已过期
	Status *string `json:"Status,omitempty" name:"Status"`

	// 管理人姓名
	ManagerFirstName *string `json:"ManagerFirstName,omitempty" name:"ManagerFirstName"`

	// 管理人姓名
	ManagerLastName *string `json:"ManagerLastName,omitempty" name:"ManagerLastName"`

	// 管理人职位
	ManagerPosition *string `json:"ManagerPosition,omitempty" name:"ManagerPosition"`

	// 管理人电话
	ManagerPhone *string `json:"ManagerPhone,omitempty" name:"ManagerPhone"`

	// 管理人邮箱
	ManagerMail *string `json:"ManagerMail,omitempty" name:"ManagerMail"`

	// 管理人所属部门
	ManagerDepartment *string `json:"ManagerDepartment,omitempty" name:"ManagerDepartment"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 管理人域名数量
	DomainCount *int64 `json:"DomainCount,omitempty" name:"DomainCount"`

	// 管理人证书数量
	CertCount *int64 `json:"CertCount,omitempty" name:"CertCount"`

	// 管理人ID
	ManagerId *int64 `json:"ManagerId,omitempty" name:"ManagerId"`

	// 审核有效到期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 最近一次提交审核时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubmitAuditTime *string `json:"SubmitAuditTime,omitempty" name:"SubmitAuditTime"`

	// 审核通过时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerifyTime *string `json:"VerifyTime,omitempty" name:"VerifyTime"`

	// 具体审核状态信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusInfo []*ManagerStatusInfo `json:"StatusInfo,omitempty" name:"StatusInfo"`
}

type ManagerStatusInfo struct {

}

// Predefined struct for user
type ModifyCertificateAliasRequestParams struct {
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 备注名称。
	Alias *string `json:"Alias,omitempty" name:"Alias"`
}

type ModifyCertificateAliasRequest struct {
	*tchttp.BaseRequest
	
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 备注名称。
	Alias *string `json:"Alias,omitempty" name:"Alias"`
}

func (r *ModifyCertificateAliasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCertificateAliasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	delete(f, "Alias")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCertificateAliasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCertificateAliasResponseParams struct {
	// 修改成功的证书 ID。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyCertificateAliasResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCertificateAliasResponseParams `json:"Response"`
}

func (r *ModifyCertificateAliasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCertificateAliasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCertificateProjectRequestParams struct {
	// 需要修改所属项目的证书 ID 集合，最多100个证书。
	CertificateIdList []*string `json:"CertificateIdList,omitempty" name:"CertificateIdList"`

	// 项目 ID。
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

type ModifyCertificateProjectRequest struct {
	*tchttp.BaseRequest
	
	// 需要修改所属项目的证书 ID 集合，最多100个证书。
	CertificateIdList []*string `json:"CertificateIdList,omitempty" name:"CertificateIdList"`

	// 项目 ID。
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *ModifyCertificateProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCertificateProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateIdList")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCertificateProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCertificateProjectResponseParams struct {
	// 修改所属项目成功的证书集合。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessCertificates []*string `json:"SuccessCertificates,omitempty" name:"SuccessCertificates"`

	// 修改所属项目失败的证书集合。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailCertificates []*string `json:"FailCertificates,omitempty" name:"FailCertificates"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyCertificateProjectResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCertificateProjectResponseParams `json:"Response"`
}

func (r *ModifyCertificateProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCertificateProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCertificatesExpiringNotificationSwitchRequestParams struct {
	// 证书ID列表。最多50个
	CertificateIds []*string `json:"CertificateIds,omitempty" name:"CertificateIds"`

	// 0:不忽略通知。1:忽略通知
	SwitchStatus *uint64 `json:"SwitchStatus,omitempty" name:"SwitchStatus"`
}

type ModifyCertificatesExpiringNotificationSwitchRequest struct {
	*tchttp.BaseRequest
	
	// 证书ID列表。最多50个
	CertificateIds []*string `json:"CertificateIds,omitempty" name:"CertificateIds"`

	// 0:不忽略通知。1:忽略通知
	SwitchStatus *uint64 `json:"SwitchStatus,omitempty" name:"SwitchStatus"`
}

func (r *ModifyCertificatesExpiringNotificationSwitchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCertificatesExpiringNotificationSwitchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateIds")
	delete(f, "SwitchStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCertificatesExpiringNotificationSwitchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCertificatesExpiringNotificationSwitchResponseParams struct {
	// 证书ID列表
	CertificateIds []*string `json:"CertificateIds,omitempty" name:"CertificateIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyCertificatesExpiringNotificationSwitchResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCertificatesExpiringNotificationSwitchResponseParams `json:"Response"`
}

func (r *ModifyCertificatesExpiringNotificationSwitchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCertificatesExpiringNotificationSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperationLog struct {
	// 操作证书动作。
	Action *string `json:"Action,omitempty" name:"Action"`

	// 操作时间。
	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`
}

type PackageInfo struct {
	// 权益包ID。
	PackageId *string `json:"PackageId,omitempty" name:"PackageId"`

	// 权益包内权益点总量。
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 权益包内权益点余量。
	Balance *uint64 `json:"Balance,omitempty" name:"Balance"`

	// 权益包名称。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 权益点是转入时，来源信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceUin *uint64 `json:"SourceUin,omitempty" name:"SourceUin"`

	// 权益点状态。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 过期时间。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 更新时间。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 生成时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 来源类型。
	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`

	// 转移信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferOutInfos []*PackageTransferOutInfo `json:"TransferOutInfos,omitempty" name:"TransferOutInfos"`
}

type PackageTransferOutInfo struct {
	// 权益包ID。
	PackageId *string `json:"PackageId,omitempty" name:"PackageId"`

	// 转移码。
	TransferCode *string `json:"TransferCode,omitempty" name:"TransferCode"`

	// 本次转移点数。
	TransferCount *uint64 `json:"TransferCount,omitempty" name:"TransferCount"`

	// 转入的PackageID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReceivePackageId *string `json:"ReceivePackageId,omitempty" name:"ReceivePackageId"`

	// 本次转移过期时间。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 本次转移生成时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 本次转移更新时间。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 转移状态。
	TransferStatus *string `json:"TransferStatus,omitempty" name:"TransferStatus"`

	// 接收者uin。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReceiverUin *uint64 `json:"ReceiverUin,omitempty" name:"ReceiverUin"`

	// 接收时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReceiveTime *string `json:"ReceiveTime,omitempty" name:"ReceiveTime"`
}

type PreAuditInfo struct {
	// 证书总年限
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPeriod *int64 `json:"TotalPeriod,omitempty" name:"TotalPeriod"`

	// 证书当前年限
	// 注意：此字段可能返回 null，表示取不到有效值。
	NowPeriod *int64 `json:"NowPeriod,omitempty" name:"NowPeriod"`

	// 证书预审核管理人ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ManagerId *string `json:"ManagerId,omitempty" name:"ManagerId"`
}

type ProjectInfo struct {
	// 项目名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 项目创建用户 UIN。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectCreatorUin *uint64 `json:"ProjectCreatorUin,omitempty" name:"ProjectCreatorUin"`

	// 项目创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectCreateTime *string `json:"ProjectCreateTime,omitempty" name:"ProjectCreateTime"`

	// 项目信息简述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectResume *string `json:"ProjectResume,omitempty" name:"ProjectResume"`

	// 用户 UIN。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUin *uint64 `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 项目 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

// Predefined struct for user
type ReplaceCertificateRequestParams struct {
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 验证类型：DNS_AUTO = 自动DNS验证（仅支持在腾讯云解析且解析状态正常的域名使用该验证类型），DNS = 手动DNS验证，FILE = 文件验证。
	ValidType *string `json:"ValidType,omitempty" name:"ValidType"`

	// 类型，默认 Original。可选项：Original = 原证书 CSR，Upload = 手动上传，Online = 在线生成。
	CsrType *string `json:"CsrType,omitempty" name:"CsrType"`

	// CSR 内容。
	CsrContent *string `json:"CsrContent,omitempty" name:"CsrContent"`

	// KEY 密码。
	CsrkeyPassword *string `json:"CsrkeyPassword,omitempty" name:"CsrkeyPassword"`

	// 重颁发原因。
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// CSR加密方式，可选：RSA、ECC、SM2
	// （CsrType为Online才可选）， 默认为RSA
	CertCSREncryptAlgo *string `json:"CertCSREncryptAlgo,omitempty" name:"CertCSREncryptAlgo"`

	// CSR加密参数，CsrEncryptAlgo为RSA时， 可选2048、4096等默认为2048；CsrEncryptAlgo为ECC时，可选prime256v1，secp384r1等，默认为prime256v1; 
	CertCSRKeyParameter *string `json:"CertCSRKeyParameter,omitempty" name:"CertCSRKeyParameter"`
}

type ReplaceCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 验证类型：DNS_AUTO = 自动DNS验证（仅支持在腾讯云解析且解析状态正常的域名使用该验证类型），DNS = 手动DNS验证，FILE = 文件验证。
	ValidType *string `json:"ValidType,omitempty" name:"ValidType"`

	// 类型，默认 Original。可选项：Original = 原证书 CSR，Upload = 手动上传，Online = 在线生成。
	CsrType *string `json:"CsrType,omitempty" name:"CsrType"`

	// CSR 内容。
	CsrContent *string `json:"CsrContent,omitempty" name:"CsrContent"`

	// KEY 密码。
	CsrkeyPassword *string `json:"CsrkeyPassword,omitempty" name:"CsrkeyPassword"`

	// 重颁发原因。
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// CSR加密方式，可选：RSA、ECC、SM2
	// （CsrType为Online才可选）， 默认为RSA
	CertCSREncryptAlgo *string `json:"CertCSREncryptAlgo,omitempty" name:"CertCSREncryptAlgo"`

	// CSR加密参数，CsrEncryptAlgo为RSA时， 可选2048、4096等默认为2048；CsrEncryptAlgo为ECC时，可选prime256v1，secp384r1等，默认为prime256v1; 
	CertCSRKeyParameter *string `json:"CertCSRKeyParameter,omitempty" name:"CertCSRKeyParameter"`
}

func (r *ReplaceCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	delete(f, "ValidType")
	delete(f, "CsrType")
	delete(f, "CsrContent")
	delete(f, "CsrkeyPassword")
	delete(f, "Reason")
	delete(f, "CertCSREncryptAlgo")
	delete(f, "CertCSRKeyParameter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReplaceCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplaceCertificateResponseParams struct {
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ReplaceCertificateResponse struct {
	*tchttp.BaseResponse
	Response *ReplaceCertificateResponseParams `json:"Response"`
}

func (r *ReplaceCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceTypeRegions struct {
	// 云资源类型
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 地域列表
	Regions []*string `json:"Regions,omitempty" name:"Regions"`
}

// Predefined struct for user
type RevokeCertificateRequestParams struct {
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 吊销证书原因。
	Reason *string `json:"Reason,omitempty" name:"Reason"`
}

type RevokeCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 吊销证书原因。
	Reason *string `json:"Reason,omitempty" name:"Reason"`
}

func (r *RevokeCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RevokeCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	delete(f, "Reason")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RevokeCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RevokeCertificateResponseParams struct {
	// 吊销证书域名验证信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RevokeDomainValidateAuths []*RevokeDomainValidateAuths `json:"RevokeDomainValidateAuths,omitempty" name:"RevokeDomainValidateAuths"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RevokeCertificateResponse struct {
	*tchttp.BaseResponse
	Response *RevokeCertificateResponseParams `json:"Response"`
}

func (r *RevokeCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RevokeCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RevokeDomainValidateAuths struct {
	// DV 认证值路径。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainValidateAuthPath *string `json:"DomainValidateAuthPath,omitempty" name:"DomainValidateAuthPath"`

	// DV 认证 KEY。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainValidateAuthKey *string `json:"DomainValidateAuthKey,omitempty" name:"DomainValidateAuthKey"`

	// DV 认证值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainValidateAuthValue *string `json:"DomainValidateAuthValue,omitempty" name:"DomainValidateAuthValue"`

	// DV 认证域名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainValidateAuthDomain *string `json:"DomainValidateAuthDomain,omitempty" name:"DomainValidateAuthDomain"`
}

type RootCertificates struct {
	// 国密签名证书
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sign *string `json:"Sign,omitempty" name:"Sign"`

	// 国密加密证书
	// 注意：此字段可能返回 null，表示取不到有效值。
	Encrypt *string `json:"Encrypt,omitempty" name:"Encrypt"`

	// 标准证书
	// 注意：此字段可能返回 null，表示取不到有效值。
	Standard *string `json:"Standard,omitempty" name:"Standard"`
}

// Predefined struct for user
type SubmitAuditManagerRequestParams struct {
	// 管理人ID
	ManagerId *int64 `json:"ManagerId,omitempty" name:"ManagerId"`
}

type SubmitAuditManagerRequest struct {
	*tchttp.BaseRequest
	
	// 管理人ID
	ManagerId *int64 `json:"ManagerId,omitempty" name:"ManagerId"`
}

func (r *SubmitAuditManagerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitAuditManagerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ManagerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitAuditManagerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitAuditManagerResponseParams struct {
	// 管理人ID
	ManagerId *int64 `json:"ManagerId,omitempty" name:"ManagerId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SubmitAuditManagerResponse struct {
	*tchttp.BaseResponse
	Response *SubmitAuditManagerResponseParams `json:"Response"`
}

func (r *SubmitAuditManagerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitAuditManagerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitCertificateInformationRequestParams struct {
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// CSR 生成方式：online = 在线生成, parse = 手动上传。
	CsrType *string `json:"CsrType,omitempty" name:"CsrType"`

	// 上传的 CSR 内容。
	CsrContent *string `json:"CsrContent,omitempty" name:"CsrContent"`

	// 绑定证书的域名。
	CertificateDomain *string `json:"CertificateDomain,omitempty" name:"CertificateDomain"`

	// 上传的域名数组（多域名证书可以上传）。
	DomainList []*string `json:"DomainList,omitempty" name:"DomainList"`

	// 私钥密码（非必填）。
	KeyPassword *string `json:"KeyPassword,omitempty" name:"KeyPassword"`

	// 公司名称。
	OrganizationName *string `json:"OrganizationName,omitempty" name:"OrganizationName"`

	// 部门名称。
	OrganizationDivision *string `json:"OrganizationDivision,omitempty" name:"OrganizationDivision"`

	// 公司详细地址。
	OrganizationAddress *string `json:"OrganizationAddress,omitempty" name:"OrganizationAddress"`

	// 国家名称，如中国：CN 。
	OrganizationCountry *string `json:"OrganizationCountry,omitempty" name:"OrganizationCountry"`

	// 公司所在城市。
	OrganizationCity *string `json:"OrganizationCity,omitempty" name:"OrganizationCity"`

	// 公司所在省份。
	OrganizationRegion *string `json:"OrganizationRegion,omitempty" name:"OrganizationRegion"`

	// 公司邮编。
	PostalCode *string `json:"PostalCode,omitempty" name:"PostalCode"`

	// 公司座机区号。
	PhoneAreaCode *string `json:"PhoneAreaCode,omitempty" name:"PhoneAreaCode"`

	// 公司座机号码。
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// 证书验证方式。验证类型：DNS_AUTO = 自动DNS验证（仅支持在腾讯云解析且解析状态正常的域名使用该验证类型），DNS = 手动DNS验证，FILE = 文件验证。
	VerifyType *string `json:"VerifyType,omitempty" name:"VerifyType"`

	// 管理人名。
	AdminFirstName *string `json:"AdminFirstName,omitempty" name:"AdminFirstName"`

	// 管理人姓。
	AdminLastName *string `json:"AdminLastName,omitempty" name:"AdminLastName"`

	// 管理人手机号码。
	AdminPhoneNum *string `json:"AdminPhoneNum,omitempty" name:"AdminPhoneNum"`

	// 管理人邮箱地址。
	AdminEmail *string `json:"AdminEmail,omitempty" name:"AdminEmail"`

	// 管理人职位。
	AdminPosition *string `json:"AdminPosition,omitempty" name:"AdminPosition"`

	// 联系人名。
	ContactFirstName *string `json:"ContactFirstName,omitempty" name:"ContactFirstName"`

	// 联系人姓。
	ContactLastName *string `json:"ContactLastName,omitempty" name:"ContactLastName"`

	// 联系人邮箱地址。
	ContactEmail *string `json:"ContactEmail,omitempty" name:"ContactEmail"`

	// 联系人手机号码。
	ContactNumber *string `json:"ContactNumber,omitempty" name:"ContactNumber"`

	// 联系人职位。
	ContactPosition *string `json:"ContactPosition,omitempty" name:"ContactPosition"`
}

type SubmitCertificateInformationRequest struct {
	*tchttp.BaseRequest
	
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// CSR 生成方式：online = 在线生成, parse = 手动上传。
	CsrType *string `json:"CsrType,omitempty" name:"CsrType"`

	// 上传的 CSR 内容。
	CsrContent *string `json:"CsrContent,omitempty" name:"CsrContent"`

	// 绑定证书的域名。
	CertificateDomain *string `json:"CertificateDomain,omitempty" name:"CertificateDomain"`

	// 上传的域名数组（多域名证书可以上传）。
	DomainList []*string `json:"DomainList,omitempty" name:"DomainList"`

	// 私钥密码（非必填）。
	KeyPassword *string `json:"KeyPassword,omitempty" name:"KeyPassword"`

	// 公司名称。
	OrganizationName *string `json:"OrganizationName,omitempty" name:"OrganizationName"`

	// 部门名称。
	OrganizationDivision *string `json:"OrganizationDivision,omitempty" name:"OrganizationDivision"`

	// 公司详细地址。
	OrganizationAddress *string `json:"OrganizationAddress,omitempty" name:"OrganizationAddress"`

	// 国家名称，如中国：CN 。
	OrganizationCountry *string `json:"OrganizationCountry,omitempty" name:"OrganizationCountry"`

	// 公司所在城市。
	OrganizationCity *string `json:"OrganizationCity,omitempty" name:"OrganizationCity"`

	// 公司所在省份。
	OrganizationRegion *string `json:"OrganizationRegion,omitempty" name:"OrganizationRegion"`

	// 公司邮编。
	PostalCode *string `json:"PostalCode,omitempty" name:"PostalCode"`

	// 公司座机区号。
	PhoneAreaCode *string `json:"PhoneAreaCode,omitempty" name:"PhoneAreaCode"`

	// 公司座机号码。
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// 证书验证方式。验证类型：DNS_AUTO = 自动DNS验证（仅支持在腾讯云解析且解析状态正常的域名使用该验证类型），DNS = 手动DNS验证，FILE = 文件验证。
	VerifyType *string `json:"VerifyType,omitempty" name:"VerifyType"`

	// 管理人名。
	AdminFirstName *string `json:"AdminFirstName,omitempty" name:"AdminFirstName"`

	// 管理人姓。
	AdminLastName *string `json:"AdminLastName,omitempty" name:"AdminLastName"`

	// 管理人手机号码。
	AdminPhoneNum *string `json:"AdminPhoneNum,omitempty" name:"AdminPhoneNum"`

	// 管理人邮箱地址。
	AdminEmail *string `json:"AdminEmail,omitempty" name:"AdminEmail"`

	// 管理人职位。
	AdminPosition *string `json:"AdminPosition,omitempty" name:"AdminPosition"`

	// 联系人名。
	ContactFirstName *string `json:"ContactFirstName,omitempty" name:"ContactFirstName"`

	// 联系人姓。
	ContactLastName *string `json:"ContactLastName,omitempty" name:"ContactLastName"`

	// 联系人邮箱地址。
	ContactEmail *string `json:"ContactEmail,omitempty" name:"ContactEmail"`

	// 联系人手机号码。
	ContactNumber *string `json:"ContactNumber,omitempty" name:"ContactNumber"`

	// 联系人职位。
	ContactPosition *string `json:"ContactPosition,omitempty" name:"ContactPosition"`
}

func (r *SubmitCertificateInformationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitCertificateInformationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	delete(f, "CsrType")
	delete(f, "CsrContent")
	delete(f, "CertificateDomain")
	delete(f, "DomainList")
	delete(f, "KeyPassword")
	delete(f, "OrganizationName")
	delete(f, "OrganizationDivision")
	delete(f, "OrganizationAddress")
	delete(f, "OrganizationCountry")
	delete(f, "OrganizationCity")
	delete(f, "OrganizationRegion")
	delete(f, "PostalCode")
	delete(f, "PhoneAreaCode")
	delete(f, "PhoneNumber")
	delete(f, "VerifyType")
	delete(f, "AdminFirstName")
	delete(f, "AdminLastName")
	delete(f, "AdminPhoneNum")
	delete(f, "AdminEmail")
	delete(f, "AdminPosition")
	delete(f, "ContactFirstName")
	delete(f, "ContactLastName")
	delete(f, "ContactEmail")
	delete(f, "ContactNumber")
	delete(f, "ContactPosition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitCertificateInformationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitCertificateInformationResponseParams struct {
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SubmitCertificateInformationResponse struct {
	*tchttp.BaseResponse
	Response *SubmitCertificateInformationResponseParams `json:"Response"`
}

func (r *SubmitCertificateInformationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitCertificateInformationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubmittedData struct {
	// CSR 类型，（online = 在线生成CSR，parse = 粘贴 CSR）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CsrType *string `json:"CsrType,omitempty" name:"CsrType"`

	// CSR 内容。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CsrContent *string `json:"CsrContent,omitempty" name:"CsrContent"`

	// 域名信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertificateDomain *string `json:"CertificateDomain,omitempty" name:"CertificateDomain"`

	// DNS 信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainList []*string `json:"DomainList,omitempty" name:"DomainList"`

	// 私钥密码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyPassword *string `json:"KeyPassword,omitempty" name:"KeyPassword"`

	// 企业或单位名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationName *string `json:"OrganizationName,omitempty" name:"OrganizationName"`

	// 部门。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationDivision *string `json:"OrganizationDivision,omitempty" name:"OrganizationDivision"`

	// 地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationAddress *string `json:"OrganizationAddress,omitempty" name:"OrganizationAddress"`

	// 国家。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationCountry *string `json:"OrganizationCountry,omitempty" name:"OrganizationCountry"`

	// 市。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationCity *string `json:"OrganizationCity,omitempty" name:"OrganizationCity"`

	// 省。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationRegion *string `json:"OrganizationRegion,omitempty" name:"OrganizationRegion"`

	// 邮政编码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PostalCode *string `json:"PostalCode,omitempty" name:"PostalCode"`

	// 座机区号。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneAreaCode *string `json:"PhoneAreaCode,omitempty" name:"PhoneAreaCode"`

	// 座机号码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// 管理员名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdminFirstName *string `json:"AdminFirstName,omitempty" name:"AdminFirstName"`

	// 管理员姓。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdminLastName *string `json:"AdminLastName,omitempty" name:"AdminLastName"`

	// 管理员电话号码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdminPhoneNum *string `json:"AdminPhoneNum,omitempty" name:"AdminPhoneNum"`

	// 管理员邮箱地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdminEmail *string `json:"AdminEmail,omitempty" name:"AdminEmail"`

	// 管理员职位。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdminPosition *string `json:"AdminPosition,omitempty" name:"AdminPosition"`

	// 联系人名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContactFirstName *string `json:"ContactFirstName,omitempty" name:"ContactFirstName"`

	// 联系人姓。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContactLastName *string `json:"ContactLastName,omitempty" name:"ContactLastName"`

	// 联系人电话号码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContactNumber *string `json:"ContactNumber,omitempty" name:"ContactNumber"`

	// 联系人邮箱地址，
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContactEmail *string `json:"ContactEmail,omitempty" name:"ContactEmail"`

	// 联系人职位。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContactPosition *string `json:"ContactPosition,omitempty" name:"ContactPosition"`

	// 验证类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerifyType *string `json:"VerifyType,omitempty" name:"VerifyType"`
}

type Tags struct {
	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type TeoInstanceDetail struct {
	// 域名
	Host *string `json:"Host,omitempty" name:"Host"`

	// 证书ID
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// 区域ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 域名状态
	Status *string `json:"Status,omitempty" name:"Status"`
}

type TkeIngressDetail struct {
	// ingress名称
	IngressName *string `json:"IngressName,omitempty" name:"IngressName"`

	// tls域名列表
	TlsDomains []*string `json:"TlsDomains,omitempty" name:"TlsDomains"`

	// ingress域名列表
	Domains []*string `json:"Domains,omitempty" name:"Domains"`
}

type TkeInstanceDetail struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 集群命名空间列表
	NamespaceList []*TkeNameSpaceDetail `json:"NamespaceList,omitempty" name:"NamespaceList"`

	// 集群类型
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 集群版本
	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`
}

type TkeNameSpaceDetail struct {
	// namespace名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// secret列表
	SecretList []*TkeSecretDetail `json:"SecretList,omitempty" name:"SecretList"`
}

type TkeSecretDetail struct {
	// secret名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 证书ID
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// ingress列表
	IngressList []*TkeIngressDetail `json:"IngressList,omitempty" name:"IngressList"`

	// 和新证书不匹配的域名列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NoMatchDomains []*string `json:"NoMatchDomains,omitempty" name:"NoMatchDomains"`
}

// Predefined struct for user
type UpdateCertificateInstanceRequestParams struct {
	// 一键更新新证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 一键更新原证书ID
	OldCertificateId *string `json:"OldCertificateId,omitempty" name:"OldCertificateId"`

	// 需要部署的资源类型，参数值可选：clb、cdn、waf、live、ddos、teo、apigateway、vod、tke、tcb
	ResourceTypes []*string `json:"ResourceTypes,omitempty" name:"ResourceTypes"`

	// 需要部署的地域列表（废弃）
	//
	// Deprecated: Regions is deprecated.
	Regions []*string `json:"Regions,omitempty" name:"Regions"`

	// 云资源需要部署的地域列表
	ResourceTypesRegions []*ResourceTypeRegions `json:"ResourceTypesRegions,omitempty" name:"ResourceTypesRegions"`
}

type UpdateCertificateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 一键更新新证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 一键更新原证书ID
	OldCertificateId *string `json:"OldCertificateId,omitempty" name:"OldCertificateId"`

	// 需要部署的资源类型，参数值可选：clb、cdn、waf、live、ddos、teo、apigateway、vod、tke、tcb
	ResourceTypes []*string `json:"ResourceTypes,omitempty" name:"ResourceTypes"`

	// 需要部署的地域列表（废弃）
	Regions []*string `json:"Regions,omitempty" name:"Regions"`

	// 云资源需要部署的地域列表
	ResourceTypesRegions []*ResourceTypeRegions `json:"ResourceTypesRegions,omitempty" name:"ResourceTypesRegions"`
}

func (r *UpdateCertificateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCertificateInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	delete(f, "OldCertificateId")
	delete(f, "ResourceTypes")
	delete(f, "Regions")
	delete(f, "ResourceTypesRegions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCertificateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCertificateInstanceResponseParams struct {
	// 云资源部署任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployRecordId *uint64 `json:"DeployRecordId,omitempty" name:"DeployRecordId"`

	// 部署状态，1表示部署成功，0表示部署失败
	DeployStatus *int64 `json:"DeployStatus,omitempty" name:"DeployStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateCertificateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *UpdateCertificateInstanceResponseParams `json:"Response"`
}

func (r *UpdateCertificateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCertificateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCertificateRecordRetryRequestParams struct {
	// 待重试部署记录ID
	DeployRecordId *int64 `json:"DeployRecordId,omitempty" name:"DeployRecordId"`

	// 待重试部署记录详情ID
	DeployRecordDetailId *int64 `json:"DeployRecordDetailId,omitempty" name:"DeployRecordDetailId"`
}

type UpdateCertificateRecordRetryRequest struct {
	*tchttp.BaseRequest
	
	// 待重试部署记录ID
	DeployRecordId *int64 `json:"DeployRecordId,omitempty" name:"DeployRecordId"`

	// 待重试部署记录详情ID
	DeployRecordDetailId *int64 `json:"DeployRecordDetailId,omitempty" name:"DeployRecordDetailId"`
}

func (r *UpdateCertificateRecordRetryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCertificateRecordRetryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeployRecordId")
	delete(f, "DeployRecordDetailId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCertificateRecordRetryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCertificateRecordRetryResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateCertificateRecordRetryResponse struct {
	*tchttp.BaseResponse
	Response *UpdateCertificateRecordRetryResponseParams `json:"Response"`
}

func (r *UpdateCertificateRecordRetryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCertificateRecordRetryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCertificateRecordRollbackRequestParams struct {
	// 待重试部署记录ID
	DeployRecordId *int64 `json:"DeployRecordId,omitempty" name:"DeployRecordId"`
}

type UpdateCertificateRecordRollbackRequest struct {
	*tchttp.BaseRequest
	
	// 待重试部署记录ID
	DeployRecordId *int64 `json:"DeployRecordId,omitempty" name:"DeployRecordId"`
}

func (r *UpdateCertificateRecordRollbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCertificateRecordRollbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeployRecordId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCertificateRecordRollbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCertificateRecordRollbackResponseParams struct {
	// 回滚部署记录ID
	DeployRecordId *int64 `json:"DeployRecordId,omitempty" name:"DeployRecordId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateCertificateRecordRollbackResponse struct {
	*tchttp.BaseResponse
	Response *UpdateCertificateRecordRollbackResponseParams `json:"Response"`
}

func (r *UpdateCertificateRecordRollbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCertificateRecordRollbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateRecordDetail struct {
	// 详情记录id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 新证书ID
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// 旧证书ID
	OldCertId *string `json:"OldCertId,omitempty" name:"OldCertId"`

	// 部署域名列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 部署资源类型
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 部署地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 部署状态
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 部署错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`

	// 部署时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 最后一次更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 部署实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 部署实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 部署监听器ID（CLB专用）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 部署监听器名称（CLB专用）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 是否开启SNI（CLB专用）
	// 注意：此字段可能返回 null，表示取不到有效值。
	SniSwitch *uint64 `json:"SniSwitch,omitempty" name:"SniSwitch"`

	// bucket名称（COS专用）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// 端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 命名空间（TKE专用）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// secret名称（TKE专用）
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretName *string `json:"SecretName,omitempty" name:"SecretName"`

	// 环境ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// TCB部署类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TCBType *string `json:"TCBType,omitempty" name:"TCBType"`
}

type UpdateRecordDetails struct {
	// 部署资源类型
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 部署资源详情列表
	List []*UpdateRecordDetail `json:"List,omitempty" name:"List"`

	// 该部署资源总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type UpdateRecordInfo struct {
	// 记录ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 新证书ID
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// 原证书ID
	OldCertId *string `json:"OldCertId,omitempty" name:"OldCertId"`

	// 部署资源类型列表
	ResourceTypes []*string `json:"ResourceTypes,omitempty" name:"ResourceTypes"`

	// 部署地域列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Regions []*string `json:"Regions,omitempty" name:"Regions"`

	// 部署状态
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 部署时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 最后一次更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

// Predefined struct for user
type UploadCertificateRequestParams struct {
	// 证书内容。
	CertificatePublicKey *string `json:"CertificatePublicKey,omitempty" name:"CertificatePublicKey"`

	// 私钥内容，证书类型为 SVR 时必填，为 CA 时可不填。
	CertificatePrivateKey *string `json:"CertificatePrivateKey,omitempty" name:"CertificatePrivateKey"`

	// 证书类型，默认 SVR。CA = CA证书，SVR = 服务器证书。
	CertificateType *string `json:"CertificateType,omitempty" name:"CertificateType"`

	// 备注名称。
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 项目 ID。
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 证书用途/证书来源。“CLB，CDN，WAF，LIVE，DDOS”
	CertificateUse *string `json:"CertificateUse,omitempty" name:"CertificateUse"`

	// 标签列表
	Tags []*Tags `json:"Tags,omitempty" name:"Tags"`

	// 相同的证书是否允许重复上传
	Repeatable *bool `json:"Repeatable,omitempty" name:"Repeatable"`
}

type UploadCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 证书内容。
	CertificatePublicKey *string `json:"CertificatePublicKey,omitempty" name:"CertificatePublicKey"`

	// 私钥内容，证书类型为 SVR 时必填，为 CA 时可不填。
	CertificatePrivateKey *string `json:"CertificatePrivateKey,omitempty" name:"CertificatePrivateKey"`

	// 证书类型，默认 SVR。CA = CA证书，SVR = 服务器证书。
	CertificateType *string `json:"CertificateType,omitempty" name:"CertificateType"`

	// 备注名称。
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 项目 ID。
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 证书用途/证书来源。“CLB，CDN，WAF，LIVE，DDOS”
	CertificateUse *string `json:"CertificateUse,omitempty" name:"CertificateUse"`

	// 标签列表
	Tags []*Tags `json:"Tags,omitempty" name:"Tags"`

	// 相同的证书是否允许重复上传
	Repeatable *bool `json:"Repeatable,omitempty" name:"Repeatable"`
}

func (r *UploadCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificatePublicKey")
	delete(f, "CertificatePrivateKey")
	delete(f, "CertificateType")
	delete(f, "Alias")
	delete(f, "ProjectId")
	delete(f, "CertificateUse")
	delete(f, "Tags")
	delete(f, "Repeatable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadCertificateResponseParams struct {
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 重复证书的ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepeatCertId *string `json:"RepeatCertId,omitempty" name:"RepeatCertId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UploadCertificateResponse struct {
	*tchttp.BaseResponse
	Response *UploadCertificateResponseParams `json:"Response"`
}

func (r *UploadCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadConfirmLetterRequestParams struct {
	// 证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// base64编码后的证书确认函文件，格式应为jpg、jpeg、png、pdf，大小应在1kb与1.4M之间。
	ConfirmLetter *string `json:"ConfirmLetter,omitempty" name:"ConfirmLetter"`
}

type UploadConfirmLetterRequest struct {
	*tchttp.BaseRequest
	
	// 证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// base64编码后的证书确认函文件，格式应为jpg、jpeg、png、pdf，大小应在1kb与1.4M之间。
	ConfirmLetter *string `json:"ConfirmLetter,omitempty" name:"ConfirmLetter"`
}

func (r *UploadConfirmLetterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadConfirmLetterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	delete(f, "ConfirmLetter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadConfirmLetterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadConfirmLetterResponseParams struct {
	// 证书ID
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 是否成功
	IsSuccess *bool `json:"IsSuccess,omitempty" name:"IsSuccess"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UploadConfirmLetterResponse struct {
	*tchttp.BaseResponse
	Response *UploadConfirmLetterResponseParams `json:"Response"`
}

func (r *UploadConfirmLetterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadConfirmLetterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadRevokeLetterRequestParams struct {
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// base64编码后的证书确认函文件，格式应为jpg、jpeg、png、pdf，大小应在1kb与1.4M之间。
	RevokeLetter *string `json:"RevokeLetter,omitempty" name:"RevokeLetter"`
}

type UploadRevokeLetterRequest struct {
	*tchttp.BaseRequest
	
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// base64编码后的证书确认函文件，格式应为jpg、jpeg、png、pdf，大小应在1kb与1.4M之间。
	RevokeLetter *string `json:"RevokeLetter,omitempty" name:"RevokeLetter"`
}

func (r *UploadRevokeLetterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadRevokeLetterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	delete(f, "RevokeLetter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadRevokeLetterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadRevokeLetterResponseParams struct {
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

	// 是否成功。
	IsSuccess *bool `json:"IsSuccess,omitempty" name:"IsSuccess"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UploadRevokeLetterResponse struct {
	*tchttp.BaseResponse
	Response *UploadRevokeLetterResponseParams `json:"Response"`
}

func (r *UploadRevokeLetterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadRevokeLetterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyManagerRequestParams struct {
	// 管理人ID
	ManagerId *int64 `json:"ManagerId,omitempty" name:"ManagerId"`
}

type VerifyManagerRequest struct {
	*tchttp.BaseRequest
	
	// 管理人ID
	ManagerId *int64 `json:"ManagerId,omitempty" name:"ManagerId"`
}

func (r *VerifyManagerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyManagerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ManagerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyManagerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyManagerResponseParams struct {
	// 管理人ID
	ManagerId *int64 `json:"ManagerId,omitempty" name:"ManagerId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type VerifyManagerResponse struct {
	*tchttp.BaseResponse
	Response *VerifyManagerResponseParams `json:"Response"`
}

func (r *VerifyManagerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyManagerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VodInstanceDetail struct {
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 证书ID
	CertId *string `json:"CertId,omitempty" name:"CertId"`
}