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

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

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

	// 加密算法，仅支持 RSA。
	CsrEncryptAlgo *string `json:"CsrEncryptAlgo,omitempty" name:"CsrEncryptAlgo"`

	// 密钥对参数，仅支持2048。
	CsrKeyParameter *string `json:"CsrKeyParameter,omitempty" name:"CsrKeyParameter"`

	// CSR 的加密密码。
	CsrKeyPassword *string `json:"CsrKeyPassword,omitempty" name:"CsrKeyPassword"`

	// 备注名称。
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 原证书 ID，用于重新申请。
	OldCertificateId *string `json:"OldCertificateId,omitempty" name:"OldCertificateId"`
}

func (r *ApplyCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ApplyCertificateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ApplyCertificateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 证书 ID。
		CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApplyCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ApplyCertificateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *CancelCertificateOrderRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CancelCertificateOrderResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 取消订单成功的证书 ID。
		CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelCertificateOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CancelCertificateOrderResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

	// 证书套餐类型：1 = GeoTrust DV SSL CA - G3， 2 = TrustAsia TLS RSA CA， 3 = SecureSite 增强型企业版（EV Pro）， 4 = SecureSite 增强型（EV）， 5 = SecureSite 企业型专业版（OV Pro）， 6 = SecureSite 企业型（OV）， 7 = SecureSite 企业型（OV）通配符， 8 = Geotrust 增强型（EV）， 9 = Geotrust 企业型（OV）， 10 = Geotrust 企业型（OV）通配符， 11 = TrustAsia 域名型多域名 SSL 证书， 12 = TrustAsia 域名型（DV）通配符， 13 = TrustAsia 企业型通配符（OV）SSL 证书（D3）， 14 = TrustAsia 企业型（OV）SSL 证书（D3）， 15 = TrustAsia 企业型多域名 （OV）SSL 证书（D3）， 16 = TrustAsia 增强型 （EV）SSL 证书（D3）， 17 = TrustAsia 增强型多域名（EV）SSL 证书（D3）， 18 = GlobalSign 企业型（OV）SSL 证书， 19 = GlobalSign 企业型通配符 （OV）SSL 证书， 20 = GlobalSign 增强型 （EV）SSL 证书， 21 = TrustAsia 企业型通配符多域名（OV）SSL 证书（D3）， 22 = GlobalSign 企业型多域名（OV）SSL 证书， 23 = GlobalSign 企业型通配符多域名（OV）SSL 证书， 24 = GlobalSign 增强型多域名（EV）SSL 证书。
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

	// 状态值 0：审核中，1：已通过，2：审核失败，3：已过期，4：已添加 DNS 解析记录，5：OV/EV 证书，待提交资料，6：订单取消中，7：已取消，8：已提交资料， 待上传确认函。
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
	SubjectAltName []*string `json:"SubjectAltName,omitempty" name:"SubjectAltName" list`

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
	BoundResource []*string `json:"BoundResource,omitempty" name:"BoundResource" list`

	// 是否可部署。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Deployable *bool `json:"Deployable,omitempty" name:"Deployable"`
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

func (r *CheckCertificateChainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckCertificateChainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true为通过检查，false为未通过检查。
		IsValid *bool `json:"IsValid,omitempty" name:"IsValid"`

		// true为可信CA，false为不可信CA。
		IsTrustedCA *bool `json:"IsTrustedCA,omitempty" name:"IsTrustedCA"`

		// 包含证书链中每一段证书的通用名称。
		Chains []*string `json:"Chains,omitempty" name:"Chains" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckCertificateChainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckCertificateChainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CommitCertificateInformationRequest struct {
	*tchttp.BaseRequest

	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`
}

func (r *CommitCertificateInformationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CommitCertificateInformationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CommitCertificateInformationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// CA机构侧订单号。
		OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

		// 证书状态：0 = 审核中，1 = 已通过，2 = 审核失败，3 = 已过期，4 = 已添加DNS记录，5 = 企业证书，待提交，6 = 订单取消中，7 = 已取消，8 = 已提交资料， 待上传确认函，9 = 证书吊销中，10 = 已吊销，11 = 重颁发中，12 = 待上传吊销确认函，13 = 免费证书待提交资料。
		Status *uint64 `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CommitCertificateInformationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CommitCertificateInformationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *CompleteCertificateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CompleteCertificateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 证书ID
		CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CompleteCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CompleteCertificateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *CreateCertificateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateCertificateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 证书ID列表
		CertificateIds []*string `json:"CertificateIds,omitempty" name:"CertificateIds" list`

		// 订单号列表
		DealIds []*string `json:"DealIds,omitempty" name:"DealIds" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateCertificateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DeleteCertificateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteCertificateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 删除结果（true：删除成功，false：删除失败）
		DeleteResult *bool `json:"DeleteResult,omitempty" name:"DeleteResult"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteCertificateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeCertificateDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCertificateDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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

		// 证书套餐类型：1 = GeoTrust DV SSL CA - G3， 2 = TrustAsia TLS RSA CA， 3 = SecureSite 增强型企业版（EV Pro）， 4 = SecureSite 增强型（EV）， 5 = SecureSite 企业型专业版（OV Pro）， 6 = SecureSite 企业型（OV）， 7 = SecureSite 企业型（OV）通配符， 8 = Geotrust 增强型（EV）， 9 = Geotrust 企业型（OV）， 10 = Geotrust 企业型（OV）通配符， 11 = TrustAsia 域名型多域名 SSL 证书， 12 = TrustAsia 域名型（DV）通配符， 13 = TrustAsia 企业型通配符（OV）SSL 证书（D3）， 14 = TrustAsia 企业型（OV）SSL 证书（D3）， 15 = TrustAsia 企业型多域名 （OV）SSL 证书（D3）， 16 = TrustAsia 增强型 （EV）SSL 证书（D3）， 17 = TrustAsia 增强型多域名（EV）SSL 证书（D3）， 18 = GlobalSign 企业型（OV）SSL 证书， 19 = GlobalSign 企业型通配符 （OV）SSL 证书， 20 = GlobalSign 增强型 （EV）SSL 证书， 21 = TrustAsia 企业型通配符多域名（OV）SSL 证书（D3）， 22 = GlobalSign 企业型多域名（OV）SSL 证书， 23 = GlobalSign 企业型通配符多域名（OV）SSL 证书， 24 = GlobalSign 增强型多域名（EV）SSL 证书，25 = Wotrus 域名型证书，26 = Wotrus 域名型多域名证书，27 = Wotrus 域名型通配符证书，28 = Wotrus 企业型证书，29 = Wotrus 企业型多域名证书，30 = Wotrus 企业型通配符证书，31 = Wotrus 增强型证书，32 = Wotrus 增强型多域名证书，33 = DNSPod 国密域名型证书，34 = DNSPod 国密域名型多域名证书，35 = DNSPod 国密域名型通配符证书，37 = DNSPod 国密企业型证书，38 = DNSPod 国密企业型多域名证书，39 = DNSPod 国密企业型通配符证书，40 = DNSPod 国密增强型证书，41 = DNSPod 国密增强型多域名证书，42 = TrustAsia 域名型通配符多域名证书。
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

		// 证书包含的多个域名（包含主域名）
	// 注意：此字段可能返回 null，表示取不到有效值。
		SubjectAltName []*string `json:"SubjectAltName,omitempty" name:"SubjectAltName" list`

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

		// 提交的资料信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
		SubmittedData *SubmittedData `json:"SubmittedData,omitempty" name:"SubmittedData"`

		// 是否可重颁发证书。
	// 注意：此字段可能返回 null，表示取不到有效值。
		RenewAble *bool `json:"RenewAble,omitempty" name:"RenewAble"`

		// 是否可部署。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Deployable *bool `json:"Deployable,omitempty" name:"Deployable"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCertificateDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCertificateDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeCertificateOperateLogsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCertificateOperateLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 当前查询条件日志总数。
		AllTotal *uint64 `json:"AllTotal,omitempty" name:"AllTotal"`

		// 本次请求返回的日志数量。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 证书操作日志列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		OperateLogs []*OperationLog `json:"OperateLogs,omitempty" name:"OperateLogs" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCertificateOperateLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCertificateOperateLogsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeCertificateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCertificateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
		SubjectAltName []*string `json:"SubjectAltName,omitempty" name:"SubjectAltName" list`

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

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCertificateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCertificatesRequest struct {
	*tchttp.BaseRequest

	// 分页偏移量，从0开始。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页数量，默认20。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索关键词，可搜索证书 ID、备注名称、域名。例如： a8xHcaIs。
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// 证书类型：CA = 客户端证书，SVR = 服务器证书。
	CertificateType *string `json:"CertificateType,omitempty" name:"CertificateType"`

	// 项目 ID。
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 按到期时间排序：DESC = 降序， ASC = 升序。
	ExpirationSort *string `json:"ExpirationSort,omitempty" name:"ExpirationSort"`

	// 证书状态。
	CertificateStatus []*uint64 `json:"CertificateStatus,omitempty" name:"CertificateStatus" list`

	// 是否可部署，可选值：1 = 可部署，0 =  不可部署。
	Deployable *uint64 `json:"Deployable,omitempty" name:"Deployable"`
}

func (r *DescribeCertificatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCertificatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCertificatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Certificates []*Certificates `json:"Certificates,omitempty" name:"Certificates" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCertificatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCertificatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *DownloadCertificateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DownloadCertificateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// ZIP base64 编码内容，base64 解码后可保存为 ZIP 文件。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Content *string `json:"Content,omitempty" name:"Content"`

		// MIME 类型：application/zip = ZIP 压缩文件。
	// 注意：此字段可能返回 null，表示取不到有效值。
		ContentType *string `json:"ContentType,omitempty" name:"ContentType"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

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
	DvAuths []*DvAuths `json:"DvAuths,omitempty" name:"DvAuths" list`
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

func (r *ModifyCertificateAliasRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyCertificateAliasResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 修改成功的证书 ID。
		CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCertificateAliasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyCertificateAliasResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyCertificateProjectRequest struct {
	*tchttp.BaseRequest

	// 需要修改所属项目的证书 ID 集合，最多100个证书。
	CertificateIdList []*string `json:"CertificateIdList,omitempty" name:"CertificateIdList" list`

	// 项目 ID。
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *ModifyCertificateProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyCertificateProjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyCertificateProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 修改所属项目成功的证书集合。
	// 注意：此字段可能返回 null，表示取不到有效值。
		SuccessCertificates []*string `json:"SuccessCertificates,omitempty" name:"SuccessCertificates" list`

		// 修改所属项目失败的证书集合。
	// 注意：此字段可能返回 null，表示取不到有效值。
		FailCertificates []*string `json:"FailCertificates,omitempty" name:"FailCertificates" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCertificateProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyCertificateProjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OperationLog struct {

	// 操作证书动作。
	Action *string `json:"Action,omitempty" name:"Action"`

	// 操作时间。
	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`
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
}

func (r *ReplaceCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReplaceCertificateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReplaceCertificateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 证书 ID。
		CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReplaceCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReplaceCertificateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *RevokeCertificateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RevokeCertificateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 吊销证书域名验证信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
		RevokeDomainValidateAuths []*RevokeDomainValidateAuths `json:"RevokeDomainValidateAuths,omitempty" name:"RevokeDomainValidateAuths" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RevokeCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

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
	DomainList []*string `json:"DomainList,omitempty" name:"DomainList" list`

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

func (r *SubmitCertificateInformationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SubmitCertificateInformationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 证书 ID。
		CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SubmitCertificateInformationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

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
	DomainList []*string `json:"DomainList,omitempty" name:"DomainList" list`

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

type UploadCertificateRequest struct {
	*tchttp.BaseRequest

	// 证书内容。
	CertificatePublicKey *string `json:"CertificatePublicKey,omitempty" name:"CertificatePublicKey"`

	// 私钥内容，证书类型为 SVR 时必填，为 CA 时可不填。
	CertificatePrivateKey *string `json:"CertificatePrivateKey,omitempty" name:"CertificatePrivateKey"`

	// 证书类型，默认 SVR。CA = 客户端证书，SVR = 服务器证书。
	CertificateType *string `json:"CertificateType,omitempty" name:"CertificateType"`

	// 备注名称。
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 项目 ID。
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *UploadCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UploadCertificateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UploadCertificateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 证书 ID。
		CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UploadCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UploadCertificateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *UploadConfirmLetterRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UploadConfirmLetterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 证书ID
		CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

		// 是否成功
		IsSuccess *bool `json:"IsSuccess,omitempty" name:"IsSuccess"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UploadConfirmLetterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UploadConfirmLetterResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *UploadRevokeLetterRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UploadRevokeLetterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 证书 ID。
		CertificateId *string `json:"CertificateId,omitempty" name:"CertificateId"`

		// 是否成功。
		IsSuccess *bool `json:"IsSuccess,omitempty" name:"IsSuccess"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UploadRevokeLetterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UploadRevokeLetterResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
