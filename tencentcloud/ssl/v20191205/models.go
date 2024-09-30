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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type ApiGatewayInstanceDetail struct {
	// 实例ID
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 实例名称
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 证书ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// 使用协议
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`
}

type ApiGatewayInstanceList struct {
	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// apigateway实例详情	
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceList []*ApiGatewayInstanceDetail `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// 该地域下apigateway实例总数	
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 是否查询异常
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *string `json:"Error,omitnil,omitempty" name:"Error"`
}

// Predefined struct for user
type ApplyCertificateRequestParams struct {
	// 验证方式：DNS_AUTO = 自动DNS验证，DNS = 手动DNS验证，FILE = 文件验证。
	DvAuthMethod *string `json:"DvAuthMethod,omitnil,omitempty" name:"DvAuthMethod"`

	// 域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 项目 ID。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 证书类型，目前仅支持类型83。83 = TrustAsia C1 DV Free。
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// 邮箱。
	ContactEmail *string `json:"ContactEmail,omitnil,omitempty" name:"ContactEmail"`

	// 手机。
	ContactPhone *string `json:"ContactPhone,omitnil,omitempty" name:"ContactPhone"`

	// 有效期，默认3个月，目前仅支持3个月。
	ValidityPeriod *string `json:"ValidityPeriod,omitnil,omitempty" name:"ValidityPeriod"`

	// 加密算法，支持 RSA及ECC。
	CsrEncryptAlgo *string `json:"CsrEncryptAlgo,omitnil,omitempty" name:"CsrEncryptAlgo"`

	// 密钥对参数，RSA仅支持2048。ECC仅支持prime256v1。加密算法选择ECC时，此参数必填
	CsrKeyParameter *string `json:"CsrKeyParameter,omitnil,omitempty" name:"CsrKeyParameter"`

	// CSR 的加密密码。
	CsrKeyPassword *string `json:"CsrKeyPassword,omitnil,omitempty" name:"CsrKeyPassword"`

	// 备注名称。
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 原证书 ID，用于重新申请。
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`

	// 权益包ID，用于免费证书扩容包使用
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 签发后是否删除自动域名验证记录， 默认为否；仅域名为DNS_AUTO验证类型支持传参
	DeleteDnsAutoRecord *bool `json:"DeleteDnsAutoRecord,omitnil,omitempty" name:"DeleteDnsAutoRecord"`

	// 域名数组（多域名证书可以上传）。	
	DnsNames []*string `json:"DnsNames,omitnil,omitempty" name:"DnsNames"`
}

type ApplyCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 验证方式：DNS_AUTO = 自动DNS验证，DNS = 手动DNS验证，FILE = 文件验证。
	DvAuthMethod *string `json:"DvAuthMethod,omitnil,omitempty" name:"DvAuthMethod"`

	// 域名。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 项目 ID。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 证书类型，目前仅支持类型83。83 = TrustAsia C1 DV Free。
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// 邮箱。
	ContactEmail *string `json:"ContactEmail,omitnil,omitempty" name:"ContactEmail"`

	// 手机。
	ContactPhone *string `json:"ContactPhone,omitnil,omitempty" name:"ContactPhone"`

	// 有效期，默认3个月，目前仅支持3个月。
	ValidityPeriod *string `json:"ValidityPeriod,omitnil,omitempty" name:"ValidityPeriod"`

	// 加密算法，支持 RSA及ECC。
	CsrEncryptAlgo *string `json:"CsrEncryptAlgo,omitnil,omitempty" name:"CsrEncryptAlgo"`

	// 密钥对参数，RSA仅支持2048。ECC仅支持prime256v1。加密算法选择ECC时，此参数必填
	CsrKeyParameter *string `json:"CsrKeyParameter,omitnil,omitempty" name:"CsrKeyParameter"`

	// CSR 的加密密码。
	CsrKeyPassword *string `json:"CsrKeyPassword,omitnil,omitempty" name:"CsrKeyPassword"`

	// 备注名称。
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 原证书 ID，用于重新申请。
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`

	// 权益包ID，用于免费证书扩容包使用
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 签发后是否删除自动域名验证记录， 默认为否；仅域名为DNS_AUTO验证类型支持传参
	DeleteDnsAutoRecord *bool `json:"DeleteDnsAutoRecord,omitnil,omitempty" name:"DeleteDnsAutoRecord"`

	// 域名数组（多域名证书可以上传）。	
	DnsNames []*string `json:"DnsNames,omitnil,omitempty" name:"DnsNames"`
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
	delete(f, "DnsNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyCertificateResponseParams struct {
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type BatchDeleteFail struct {
	// 失败的证书ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// 失败的原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`
}

type BindResourceRegionResult struct {
	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 关联资源总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 是否查询异常
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *string `json:"Error,omitnil,omitempty" name:"Error"`
}

type BindResourceResult struct {
	// 资源类型：clb、cdn、waf、live、vod、ddos、tke、apigateway、tcb、teo（edgeOne）
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 绑定资源地域结果
	BindResourceRegionResult []*BindResourceRegionResult `json:"BindResourceRegionResult,omitnil,omitempty" name:"BindResourceRegionResult"`
}

type COSInstanceList struct {
	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 实例详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceList []*CosInstanceDetail `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// 地域下总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *string `json:"Error,omitnil,omitempty" name:"Error"`
}

// Predefined struct for user
type CancelAuditCertificateRequestParams struct {
	// 证书ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`
}

type CancelAuditCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 证书ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`
}

func (r *CancelAuditCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelAuditCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelAuditCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelAuditCertificateResponseParams struct {
	// 操作是否成功
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CancelAuditCertificateResponse struct {
	*tchttp.BaseResponse
	Response *CancelAuditCertificateResponseParams `json:"Response"`
}

func (r *CancelAuditCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelAuditCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelCertificateOrderRequestParams struct {
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`
}

type CancelCertificateOrderRequest struct {
	*tchttp.BaseRequest
	
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`
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
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 已部署证书ID
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// 域名状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 域名计费状态
	HttpsBillingSwitch *string `json:"HttpsBillingSwitch,omitnil,omitempty" name:"HttpsBillingSwitch"`
}

type CdnInstanceList struct {
	// 该地域下CDN域名总数	
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// cdn域名详情	
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceList []*CdnInstanceDetail `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// 是否查询异常
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *string `json:"Error,omitnil,omitempty" name:"Error"`
}

type CertTaskId struct {
	// 证书ID
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// 异步任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type Certificate struct {
	// 证书ID
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// 证书绑定的域名
	DnsNames []*string `json:"DnsNames,omitnil,omitempty" name:"DnsNames"`

	// 根证书ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertCaId *string `json:"CertCaId,omitnil,omitempty" name:"CertCaId"`

	// 证书认证模式：UNIDIRECTIONAL单向认证，MUTUAL双向认证
	// 注意：此字段可能返回 null，表示取不到有效值。
	SSLMode *string `json:"SSLMode,omitnil,omitempty" name:"SSLMode"`
}

type CertificateExtra struct {
	// 证书可配置域名数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainNumber *string `json:"DomainNumber,omitnil,omitempty" name:"DomainNumber"`

	// 原始证书 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginCertificateId *string `json:"OriginCertificateId,omitnil,omitempty" name:"OriginCertificateId"`

	// 重颁发证书原始 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReplacedBy *string `json:"ReplacedBy,omitnil,omitempty" name:"ReplacedBy"`

	// 重颁发证书新 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReplacedFor *string `json:"ReplacedFor,omitnil,omitempty" name:"ReplacedFor"`

	// 新订单证书 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewOrder *string `json:"RenewOrder,omitnil,omitempty" name:"RenewOrder"`

	// 是否是国密证书
	// 注意：此字段可能返回 null，表示取不到有效值。
	SMCert *int64 `json:"SMCert,omitnil,omitempty" name:"SMCert"`

	// 公司类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompanyType *int64 `json:"CompanyType,omitnil,omitempty" name:"CompanyType"`
}

type Certificates struct {
	// 用户 UIN。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 项目 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 证书来源。
	// 注意：此字段可能返回 null，表示取不到有效值。
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 证书套餐类型：
	// 2 = TrustAsia TLS RSA CA，3 = SecureSite 增强型企业版（EV Pro）， 4 = SecureSite 增强型（EV）， 5 = SecureSite 企业型专业版（OV Pro）， 6 = SecureSite 企业型（OV）， 7 = SecureSite 企业型（OV）通配符， 8 = Geotrust 增强型（EV）， 9 = Geotrust 企业型（OV）， 10 = Geotrust 企业型（OV）通配符， 11 = TrustAsia 域名型多域名 SSL 证书， 12 = TrustAsia 域名型（DV）通配符， 13 = TrustAsia 企业型通配符（OV）SSL 证书（D3）， 14 = TrustAsia 企业型（OV）SSL 证书（D3）， 15 = TrustAsia 企业型多域名 （OV）SSL 证书（D3）， 16 = TrustAsia 增强型 （EV）SSL 证书（D3）， 17 = TrustAsia 增强型多域名（EV）SSL 证书（D3）， 18 = GlobalSign 企业型（OV）SSL 证书， 19 = GlobalSign 企业型通配符 （OV）SSL 证书， 20 = GlobalSign 增强型 （EV）SSL 证书， 21 = TrustAsia 企业型通配符多域名（OV）SSL 证书（D3）， 22 = GlobalSign 企业型多域名（OV）SSL 证书， 23 = GlobalSign 企业型通配符多域名（OV）SSL 证书， 24 = GlobalSign 增强型多域名（EV）SSL 证书，25 = Wotrus 域名型证书，26 = Wotrus 域名型多域名证书，27 = Wotrus 域名型通配符证书，28 = Wotrus 企业型证书，29 = Wotrus 企业型多域名证书，30 = Wotrus 企业型通配符证书，31 = Wotrus 增强型证书，32 = Wotrus 增强型多域名证书，33 = Wotrus 国密域名型证书，34 = Wotrus 国密域名型多域名证书，35 = Wotrus 国密域名型通配符证书，37 = Wotrus 国密企业型证书，38 = Wotrus 国密企业型多域名证书，39 = Wotrus 国密企业型通配符证书，40 = Wotrus 国密增强型证书，41 = Wotrus 国密增强型多域名证书，42 = TrustAsia 域名型通配符多域名证书，43 = DNSPod-企业型(OV)SSL证书，44 = DNSPod-企业型(OV)通配符SSL证书，45 = DNSPod-企业型(OV)多域名SSL证书， 46 = DNSPod-增强型(EV)SSL证书，47 = DNSPod-增强型(EV)多域名SSL证书，48 = DNSPod-域名型(DV)SSL证书，49 = DNSPod-域名型(DV)通配符SSL证书，50 = DNSPod-域名型(DV)多域名SSL证书，51 = DNSPod（国密）-企业型(OV)SSL证书，52 = DNSPod（国密）-企业型(OV)通配符SSL证书，53 = DNSPod（国密）-企业型(OV)多域名SSL证书，54 = DNSPod（国密）-域名型(DV)SSL证书，55 = DNSPod（国密）-域名型(DV)通配符SSL证书， 56 = DNSPod（国密）-域名型(DV)多域名SSL证书，57 = SecureSite 企业型专业版多域名(OV Pro)，58 = SecureSite 企业型多域名(OV)，59 = SecureSite 增强型专业版多域名(EV Pro)，60 = SecureSite 增强型多域名(EV)，61 = Geotrust 增强型多域名(EV)
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// 证书类型：CA = 客户端证书，SVR = 服务器证书。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertificateType *string `json:"CertificateType,omitnil,omitempty" name:"CertificateType"`

	// 颁发者。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductZhName *string `json:"ProductZhName,omitnil,omitempty" name:"ProductZhName"`

	// 主域名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 备注名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 状态。0：审核中，1：已通过，2：审核失败，3：已过期，4：验证方式为 DNS_AUTO 类型的证书， 已添加DNS记录，5：企业证书，待提交，6：订单取消中，7：已取消，8：已提交资料， 待上传确认函，9：证书吊销中，10：已吊销，11：重颁发中，12：待上传吊销确认函，13：免费证书待提交资料状态，14：已退款，
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 证书扩展信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertificateExtra *CertificateExtra `json:"CertificateExtra,omitnil,omitempty" name:"CertificateExtra"`

	// 漏洞扫描状态：INACTIVE = 未开启，ACTIVE = 已开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulnerabilityStatus *string `json:"VulnerabilityStatus,omitnil,omitempty" name:"VulnerabilityStatus"`

	// 状态信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusMsg *string `json:"StatusMsg,omitnil,omitempty" name:"StatusMsg"`

	// 验证类型：DNS_AUTO = 自动DNS验证，DNS = 手动DNS验证，FILE = 文件验证，EMAIL = 邮件验证。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerifyType *string `json:"VerifyType,omitnil,omitempty" name:"VerifyType"`

	// 证书生效时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertBeginTime *string `json:"CertBeginTime,omitnil,omitempty" name:"CertBeginTime"`

	// 证书过期时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertEndTime *string `json:"CertEndTime,omitnil,omitempty" name:"CertEndTime"`

	// 证书有效期，单位（月）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValidityPeriod *string `json:"ValidityPeriod,omitnil,omitempty" name:"ValidityPeriod"`

	// 创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InsertTime *string `json:"InsertTime,omitnil,omitempty" name:"InsertTime"`

	// 证书 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 证书包含的多个域名（包含主域名）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubjectAltName []*string `json:"SubjectAltName,omitnil,omitempty" name:"SubjectAltName"`

	// 证书类型名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageTypeName *string `json:"PackageTypeName,omitnil,omitempty" name:"PackageTypeName"`

	// 状态名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusName *string `json:"StatusName,omitnil,omitempty" name:"StatusName"`

	// 是否为 VIP 客户。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsVip *bool `json:"IsVip,omitnil,omitempty" name:"IsVip"`

	// 是否为 DV 版证书。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDv *bool `json:"IsDv,omitnil,omitempty" name:"IsDv"`

	// 是否为泛域名证书。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsWildcard *bool `json:"IsWildcard,omitnil,omitempty" name:"IsWildcard"`

	// 是否启用了漏洞扫描功能。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsVulnerability *bool `json:"IsVulnerability,omitnil,omitempty" name:"IsVulnerability"`

	// 是否可续费。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewAble *bool `json:"RenewAble,omitnil,omitempty" name:"RenewAble"`

	// 项目信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectInfo *ProjectInfo `json:"ProjectInfo,omitnil,omitempty" name:"ProjectInfo"`

	// 关联的云资源，暂不可用
	// 注意：此字段可能返回 null，表示取不到有效值。
	BoundResource []*string `json:"BoundResource,omitnil,omitempty" name:"BoundResource"`

	// 是否可部署。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Deployable *bool `json:"Deployable,omitnil,omitempty" name:"Deployable"`

	// 标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tags `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 是否已忽略到期通知
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsIgnore *bool `json:"IsIgnore,omitnil,omitempty" name:"IsIgnore"`

	// 是否国密证书
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsSM *bool `json:"IsSM,omitnil,omitempty" name:"IsSM"`

	// 证书算法
	// 注意：此字段可能返回 null，表示取不到有效值。
	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitnil,omitempty" name:"EncryptAlgorithm"`

	// 上传CA证书的加密算法
	// 注意：此字段可能返回 null，表示取不到有效值。
	CAEncryptAlgorithms []*string `json:"CAEncryptAlgorithms,omitnil,omitempty" name:"CAEncryptAlgorithms"`

	// 上传CA证书的过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CAEndTimes []*string `json:"CAEndTimes,omitnil,omitempty" name:"CAEndTimes"`

	// 上传CA证书的通用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	CACommonNames []*string `json:"CACommonNames,omitnil,omitempty" name:"CACommonNames"`

	// 证书预审核信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PreAuditInfo *PreAuditInfo `json:"PreAuditInfo,omitnil,omitempty" name:"PreAuditInfo"`

	// 是否自动续费
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 托管状态，0，托管中，5，资源替换中， 10， 托管完成， -1未托管 
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostingStatus *int64 `json:"HostingStatus,omitnil,omitempty" name:"HostingStatus"`

	// 托管完成时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostingCompleteTime *string `json:"HostingCompleteTime,omitnil,omitempty" name:"HostingCompleteTime"`

	// 托管新证书ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostingRenewCertId *string `json:"HostingRenewCertId,omitnil,omitempty" name:"HostingRenewCertId"`

	// 存在的续费证书ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	HasRenewOrder *string `json:"HasRenewOrder,omitnil,omitempty" name:"HasRenewOrder"`

	// 重颁发证书原证书是否删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReplaceOriCertIsDelete *bool `json:"ReplaceOriCertIsDelete,omitnil,omitempty" name:"ReplaceOriCertIsDelete"`

	// 是否即将过期， 证书即将到期的30天内为即将过期
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsExpiring *bool `json:"IsExpiring,omitnil,omitempty" name:"IsExpiring"`

	// DV证书添加验证截止时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DVAuthDeadline *string `json:"DVAuthDeadline,omitnil,omitempty" name:"DVAuthDeadline"`

	// 域名验证通过时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValidationPassedTime *string `json:"ValidationPassedTime,omitnil,omitempty" name:"ValidationPassedTime"`

	// 证书关联的多域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertSANs []*string `json:"CertSANs,omitnil,omitempty" name:"CertSANs"`

	// 域名验证驳回信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AwaitingValidationMsg *string `json:"AwaitingValidationMsg,omitnil,omitempty" name:"AwaitingValidationMsg"`

	// 是否允许下载
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllowDownload *bool `json:"AllowDownload,omitnil,omitempty" name:"AllowDownload"`

	// 证书域名是否全部在DNSPOD托管解析
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDNSPODResolve *bool `json:"IsDNSPODResolve,omitnil,omitempty" name:"IsDNSPODResolve"`

	// 是否是权益点购买的证书
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsPackage *bool `json:"IsPackage,omitnil,omitempty" name:"IsPackage"`

	// 是否存在私钥密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyPasswordCustomFlag *bool `json:"KeyPasswordCustomFlag,omitnil,omitempty" name:"KeyPasswordCustomFlag"`

	// 支持下载的WEB服务器类型： nginx、apache、iis、tomcat、jks、root、other
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportDownloadType *SupportDownloadType `json:"SupportDownloadType,omitnil,omitempty" name:"SupportDownloadType"`

	// 证书吊销完成时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertRevokedTime *string `json:"CertRevokedTime,omitnil,omitempty" name:"CertRevokedTime"`

	// 托管资源类型列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostingResourceTypes []*string `json:"HostingResourceTypes,omitnil,omitempty" name:"HostingResourceTypes"`

	// 托管配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostingConfig *HostingConfig `json:"HostingConfig,omitnil,omitempty" name:"HostingConfig"`
}

// Predefined struct for user
type CheckCertificateChainRequestParams struct {
	// 待检查的证书链
	CertificateChain *string `json:"CertificateChain,omitnil,omitempty" name:"CertificateChain"`
}

type CheckCertificateChainRequest struct {
	*tchttp.BaseRequest
	
	// 待检查的证书链
	CertificateChain *string `json:"CertificateChain,omitnil,omitempty" name:"CertificateChain"`
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
	IsValid *bool `json:"IsValid,omitnil,omitempty" name:"IsValid"`

	// true为可信CA，false为不可信CA。
	IsTrustedCA *bool `json:"IsTrustedCA,omitnil,omitempty" name:"IsTrustedCA"`

	// 包含证书链中每一段证书的通用名称。
	Chains []*string `json:"Chains,omitnil,omitempty" name:"Chains"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type CheckCertificateDomainVerificationRequestParams struct {
	// 证书ID。
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`
}

type CheckCertificateDomainVerificationRequest struct {
	*tchttp.BaseRequest
	
	// 证书ID。
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`
}

func (r *CheckCertificateDomainVerificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckCertificateDomainVerificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckCertificateDomainVerificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckCertificateDomainVerificationResponseParams struct {
	// 域名验证结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerificationResults []*DomainValidationResult `json:"VerificationResults,omitnil,omitempty" name:"VerificationResults"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckCertificateDomainVerificationResponse struct {
	*tchttp.BaseResponse
	Response *CheckCertificateDomainVerificationResponseParams `json:"Response"`
}

func (r *CheckCertificateDomainVerificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckCertificateDomainVerificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClbInstanceDetail struct {
	// CLB实例ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// CLB实例名称
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// CLB监听器列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Listeners []*ClbListener `json:"Listeners,omitnil,omitempty" name:"Listeners"`
}

type ClbInstanceList struct {
	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// clb实例详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceList []*ClbInstanceDetail `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// 该地域下Clb实例总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 是否查询异常
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *string `json:"Error,omitnil,omitempty" name:"Error"`
}

type ClbListener struct {
	// 监听器ID
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 监听器名称
	ListenerName *string `json:"ListenerName,omitnil,omitempty" name:"ListenerName"`

	// 是否开启SNI，1为开启，0为关闭
	SniSwitch *uint64 `json:"SniSwitch,omitnil,omitempty" name:"SniSwitch"`

	// 监听器协议类型， HTTPS|TCP_SSL
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 监听器绑定的证书数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Certificate *Certificate `json:"Certificate,omitnil,omitempty" name:"Certificate"`

	// 监听器规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rules []*ClbListenerRule `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 不匹配域名列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NoMatchDomains []*string `json:"NoMatchDomains,omitnil,omitempty" name:"NoMatchDomains"`
}

type ClbListenerRule struct {
	// 规则ID
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// 规则绑定的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 规则是否匹配待绑定证书的域名
	IsMatch *bool `json:"IsMatch,omitnil,omitempty" name:"IsMatch"`

	// 规则已绑定的证书数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Certificate *Certificate `json:"Certificate,omitnil,omitempty" name:"Certificate"`

	// 不匹配域名列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NoMatchDomains []*string `json:"NoMatchDomains,omitnil,omitempty" name:"NoMatchDomains"`

	// 规则绑定的路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

// Predefined struct for user
type CommitCertificateInformationRequestParams struct {
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 域名验证方式
	VerifyType *string `json:"VerifyType,omitnil,omitempty" name:"VerifyType"`
}

type CommitCertificateInformationRequest struct {
	*tchttp.BaseRequest
	
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 域名验证方式
	VerifyType *string `json:"VerifyType,omitnil,omitempty" name:"VerifyType"`
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
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// 证书状态：0 = 审核中，1 = 已通过，2 = 审核失败，3 = 已过期，4 = 已添加DNS记录，5 = 企业证书，待提交，6 = 订单取消中，7 = 已取消，8 = 已提交资料， 待上传确认函，9 = 证书吊销中，10 = 已吊销，11 = 重颁发中，12 = 待上传吊销确认函，13 = 免费证书待提交资料。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	CompanyName *string `json:"CompanyName,omitnil,omitempty" name:"CompanyName"`

	// 公司ID
	CompanyId *int64 `json:"CompanyId,omitnil,omitempty" name:"CompanyId"`

	// 公司所在国家
	CompanyCountry *string `json:"CompanyCountry,omitnil,omitempty" name:"CompanyCountry"`

	// 公司所在省份
	CompanyProvince *string `json:"CompanyProvince,omitnil,omitempty" name:"CompanyProvince"`

	// 公司所在城市
	CompanyCity *string `json:"CompanyCity,omitnil,omitempty" name:"CompanyCity"`

	// 公司所在详细地址
	CompanyAddress *string `json:"CompanyAddress,omitnil,omitempty" name:"CompanyAddress"`

	// 公司电话
	CompanyPhone *string `json:"CompanyPhone,omitnil,omitempty" name:"CompanyPhone"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdType *string `json:"IdType,omitnil,omitempty" name:"IdType"`

	// ID号
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdNumber *string `json:"IdNumber,omitnil,omitempty" name:"IdNumber"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tags `json:"Tags,omitnil,omitempty" name:"Tags"`
}

// Predefined struct for user
type CompleteCertificateRequestParams struct {
	// 证书ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`
}

type CompleteCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 证书ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`
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
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 已绑定的证书ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// ENABLED: 域名上线状态
	// DISABLED:域名下线状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 存储桶名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// 存储桶地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

// Predefined struct for user
type CreateCertificateBindResourceSyncTaskRequestParams struct {
	// 证书ID列表，总数不能超过100
	CertificateIds []*string `json:"CertificateIds,omitnil,omitempty" name:"CertificateIds"`

	// 是否使用缓存， 1使用缓存，0不使用缓存； 默认为1使用缓存； 若当前证书ID存在半小时已完成的任务， 则使用缓存的情况下， 会读取半小时内离当前时间最近的查询结果
	IsCache *uint64 `json:"IsCache,omitnil,omitempty" name:"IsCache"`
}

type CreateCertificateBindResourceSyncTaskRequest struct {
	*tchttp.BaseRequest
	
	// 证书ID列表，总数不能超过100
	CertificateIds []*string `json:"CertificateIds,omitnil,omitempty" name:"CertificateIds"`

	// 是否使用缓存， 1使用缓存，0不使用缓存； 默认为1使用缓存； 若当前证书ID存在半小时已完成的任务， 则使用缓存的情况下， 会读取半小时内离当前时间最近的查询结果
	IsCache *uint64 `json:"IsCache,omitnil,omitempty" name:"IsCache"`
}

func (r *CreateCertificateBindResourceSyncTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCertificateBindResourceSyncTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateIds")
	delete(f, "IsCache")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCertificateBindResourceSyncTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCertificateBindResourceSyncTaskResponseParams struct {
	// 证书关联云资源异步任务ID列表
	CertTaskIds []*CertTaskId `json:"CertTaskIds,omitnil,omitempty" name:"CertTaskIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCertificateBindResourceSyncTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateCertificateBindResourceSyncTaskResponseParams `json:"Response"`
}

func (r *CreateCertificateBindResourceSyncTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCertificateBindResourceSyncTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCertificateByPackageRequestParams struct {
	// 证书产品PID。
	ProductPid *uint64 `json:"ProductPid,omitnil,omitempty" name:"ProductPid"`

	// 要消耗的权益包ID。
	PackageIds []*string `json:"PackageIds,omitnil,omitempty" name:"PackageIds"`

	// 证书域名数量。
	DomainCount *string `json:"DomainCount,omitnil,omitempty" name:"DomainCount"`

	// 多年期证书年限。
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 要续费的原证书ID（续费时填写）。
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`

	// 续费时CSR生成方式（original、upload、online）。
	RenewGenCsrMethod *string `json:"RenewGenCsrMethod,omitnil,omitempty" name:"RenewGenCsrMethod"`

	// 续费时选择上传CSR时填写CSR。
	RenewCsr *string `json:"RenewCsr,omitnil,omitempty" name:"RenewCsr"`

	// 续费证书CSR的算法类型。
	RenewAlgorithmType *string `json:"RenewAlgorithmType,omitnil,omitempty" name:"RenewAlgorithmType"`

	// 续费证书CSR的算法参数。
	RenewAlgorithmParam *string `json:"RenewAlgorithmParam,omitnil,omitempty" name:"RenewAlgorithmParam"`

	// 项目ID。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 标签。
	Tags []*Tags `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 续费证书的私钥密码。
	RenewKeyPass *string `json:"RenewKeyPass,omitnil,omitempty" name:"RenewKeyPass"`

	// 批量购买证书时预填写的域名。
	DomainNames *string `json:"DomainNames,omitnil,omitempty" name:"DomainNames"`

	// 批量购买证书数量。
	CertificateCount *uint64 `json:"CertificateCount,omitnil,omitempty" name:"CertificateCount"`

	// 预填写的管理人ID。
	ManagerId *uint64 `json:"ManagerId,omitnil,omitempty" name:"ManagerId"`

	// 预填写的公司ID。
	CompanyId *uint64 `json:"CompanyId,omitnil,omitempty" name:"CompanyId"`

	// 验证方式
	VerifyType *string `json:"VerifyType,omitnil,omitempty" name:"VerifyType"`
}

type CreateCertificateByPackageRequest struct {
	*tchttp.BaseRequest
	
	// 证书产品PID。
	ProductPid *uint64 `json:"ProductPid,omitnil,omitempty" name:"ProductPid"`

	// 要消耗的权益包ID。
	PackageIds []*string `json:"PackageIds,omitnil,omitempty" name:"PackageIds"`

	// 证书域名数量。
	DomainCount *string `json:"DomainCount,omitnil,omitempty" name:"DomainCount"`

	// 多年期证书年限。
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 要续费的原证书ID（续费时填写）。
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`

	// 续费时CSR生成方式（original、upload、online）。
	RenewGenCsrMethod *string `json:"RenewGenCsrMethod,omitnil,omitempty" name:"RenewGenCsrMethod"`

	// 续费时选择上传CSR时填写CSR。
	RenewCsr *string `json:"RenewCsr,omitnil,omitempty" name:"RenewCsr"`

	// 续费证书CSR的算法类型。
	RenewAlgorithmType *string `json:"RenewAlgorithmType,omitnil,omitempty" name:"RenewAlgorithmType"`

	// 续费证书CSR的算法参数。
	RenewAlgorithmParam *string `json:"RenewAlgorithmParam,omitnil,omitempty" name:"RenewAlgorithmParam"`

	// 项目ID。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 标签。
	Tags []*Tags `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 续费证书的私钥密码。
	RenewKeyPass *string `json:"RenewKeyPass,omitnil,omitempty" name:"RenewKeyPass"`

	// 批量购买证书时预填写的域名。
	DomainNames *string `json:"DomainNames,omitnil,omitempty" name:"DomainNames"`

	// 批量购买证书数量。
	CertificateCount *uint64 `json:"CertificateCount,omitnil,omitempty" name:"CertificateCount"`

	// 预填写的管理人ID。
	ManagerId *uint64 `json:"ManagerId,omitnil,omitempty" name:"ManagerId"`

	// 预填写的公司ID。
	CompanyId *uint64 `json:"CompanyId,omitnil,omitempty" name:"CompanyId"`

	// 验证方式
	VerifyType *string `json:"VerifyType,omitnil,omitempty" name:"VerifyType"`
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
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 批量购买证书时返回多个证书ID。
	CertificateIds []*string `json:"CertificateIds,omitnil,omitempty" name:"CertificateIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ProductId *int64 `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 证书包含的域名数量
	DomainNum *int64 `json:"DomainNum,omitnil,omitempty" name:"DomainNum"`

	// 证书年限，当前只支持 1 年证书的购买
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`
}

type CreateCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 证书商品ID，3 = SecureSite 增强型企业版（EV Pro）， 4 = SecureSite 增强型（EV）， 5 = SecureSite 企业型专业版（OV Pro）， 6 = SecureSite 企业型（OV）， 7 = SecureSite 企业型（OV）通配符， 8 = Geotrust 增强型（EV）， 9 = Geotrust 企业型（OV）， 10 = Geotrust 企业型（OV）通配符， 11 = TrustAsia 域名型多域名 SSL 证书， 12 = TrustAsia 域名型（DV）通配符， 13 = TrustAsia 企业型通配符（OV）SSL 证书（D3）， 14 = TrustAsia 企业型（OV）SSL 证书（D3）， 15 = TrustAsia 企业型多域名 （OV）SSL 证书（D3）， 16 = TrustAsia 增强型 （EV）SSL 证书（D3）， 17 = TrustAsia 增强型多域名（EV）SSL 证书（D3）， 18 = GlobalSign 企业型（OV）SSL 证书， 19 = GlobalSign 企业型通配符 （OV）SSL 证书， 20 = GlobalSign 增强型 （EV）SSL 证书， 21 = TrustAsia 企业型通配符多域名（OV）SSL 证书（D3）， 22 = GlobalSign 企业型多域名（OV）SSL 证书， 23 = GlobalSign 企业型通配符多域名（OV）SSL 证书， 24 = GlobalSign 增强型多域名（EV）SSL 证书，25 = Wotrus 域名型证书，26 = Wotrus 域名型多域名证书，27 = Wotrus 域名型通配符证书，28 = Wotrus 企业型证书，29 = Wotrus 企业型多域名证书，30 = Wotrus 企业型通配符证书，31 = Wotrus 增强型证书，32 = Wotrus 增强型多域名证书，33 = DNSPod 国密域名型证书，34 = DNSPod 国密域名型多域名证书，35 = DNSPod 国密域名型通配符证书，37 = DNSPod 国密企业型证书，38 = DNSPod 国密企业型多域名证书，39 = DNSPod 国密企业型通配符证书，40 = DNSPod 国密增强型证书，41 = DNSPod 国密增强型多域名证书，42 = TrustAsia 域名型通配符多域名证书。
	ProductId *int64 `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 证书包含的域名数量
	DomainNum *int64 `json:"DomainNum,omitnil,omitempty" name:"DomainNum"`

	// 证书年限，当前只支持 1 年证书的购买
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`
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
	CertificateIds []*string `json:"CertificateIds,omitnil,omitempty" name:"CertificateIds"`

	// 订单号列表
	DealIds []*string `json:"DealIds,omitnil,omitempty" name:"DealIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 协议类型
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 证书ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// 转发端口
	VirtualPort *string `json:"VirtualPort,omitnil,omitempty" name:"VirtualPort"`
}

type DdosInstanceList struct {
	// 该地域下ddos域名总数	
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// ddos实例详情	
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceList []*DdosInstanceDetail `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// 是否查询异常
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *string `json:"Error,omitnil,omitempty" name:"Error"`
}

// Predefined struct for user
type DeleteCertificateRequestParams struct {
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 删除时是否检查证书关联了云资源。默认不检查。如选择检查(需要授权服务角色SSL_QCSLinkedRoleInReplaceLoadCertificate)删除将变成异步,接口会返回异步任务ID。需使用DescribeDeleteCertificatesTaskResult接口查询删除是否成功。
	IsCheckResource *bool `json:"IsCheckResource,omitnil,omitempty" name:"IsCheckResource"`
}

type DeleteCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 删除时是否检查证书关联了云资源。默认不检查。如选择检查(需要授权服务角色SSL_QCSLinkedRoleInReplaceLoadCertificate)删除将变成异步,接口会返回异步任务ID。需使用DescribeDeleteCertificatesTaskResult接口查询删除是否成功。
	IsCheckResource *bool `json:"IsCheckResource,omitnil,omitempty" name:"IsCheckResource"`
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
	delete(f, "IsCheckResource")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCertificateResponseParams struct {
	// 删除结果（true：删除成功，false：删除失败）
	DeleteResult *bool `json:"DeleteResult,omitnil,omitempty" name:"DeleteResult"`

	// 异步删除的任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteCertificatesRequestParams struct {
	// 要删除的证书ID。单次最多100个
	CertificateIds []*string `json:"CertificateIds,omitnil,omitempty" name:"CertificateIds"`

	// 删除时是否检查证书关联了云资源。默认不检查。如需要检查关联云资源 (需授权服务角色SSL_QCSLinkedRoleInReplaceLoadCertificate)，完成授权后，删除将变成异步任务，接口会返回异步任务ID。需搭配 DescribeDeleteCertificatesTaskResult接口使用，查询删除任务是否成功。
	IsSync *bool `json:"IsSync,omitnil,omitempty" name:"IsSync"`
}

type DeleteCertificatesRequest struct {
	*tchttp.BaseRequest
	
	// 要删除的证书ID。单次最多100个
	CertificateIds []*string `json:"CertificateIds,omitnil,omitempty" name:"CertificateIds"`

	// 删除时是否检查证书关联了云资源。默认不检查。如需要检查关联云资源 (需授权服务角色SSL_QCSLinkedRoleInReplaceLoadCertificate)，完成授权后，删除将变成异步任务，接口会返回异步任务ID。需搭配 DescribeDeleteCertificatesTaskResult接口使用，查询删除任务是否成功。
	IsSync *bool `json:"IsSync,omitnil,omitempty" name:"IsSync"`
}

func (r *DeleteCertificatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCertificatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateIds")
	delete(f, "IsSync")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCertificatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCertificatesResponseParams struct {
	// 成功的ID
	Success []*string `json:"Success,omitnil,omitempty" name:"Success"`

	// 失败的ID和原因
	Fail []*BatchDeleteFail `json:"Fail,omitnil,omitempty" name:"Fail"`

	// 证书ID和异步任务的ID
	CertTaskIds []*CertTaskId `json:"CertTaskIds,omitnil,omitempty" name:"CertTaskIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCertificatesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCertificatesResponseParams `json:"Response"`
}

func (r *DeleteCertificatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCertificatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteManagerRequestParams struct {
	// 管理人ID
	ManagerId *int64 `json:"ManagerId,omitnil,omitempty" name:"ManagerId"`
}

type DeleteManagerRequest struct {
	*tchttp.BaseRequest
	
	// 管理人ID
	ManagerId *int64 `json:"ManagerId,omitnil,omitempty" name:"ManagerId"`
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
	ManagerId *int64 `json:"ManagerId,omitnil,omitempty" name:"ManagerId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type DeleteTaskResult struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 证书ID
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// 异步查询结果： 0表示任务进行中、 1表示任务成功、 2表示任务失败、3表示未授权服务角色导致任务失败、4表示有未解绑的云资源导致任务失败、5表示查询关联云资源超时导致任务失败
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *string `json:"Error,omitnil,omitempty" name:"Error"`

	// 当前结果缓存时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheTime *string `json:"CacheTime,omitnil,omitempty" name:"CacheTime"`

	// 包含的域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`
}

// Predefined struct for user
type DeployCertificateInstanceRequestParams struct {
	// 待部署的证书ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 需要部署实例列表
	InstanceIdList []*string `json:"InstanceIdList,omitnil,omitempty" name:"InstanceIdList"`

	// 部署的云资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 部署云资源状态：
	// 云直播：
	// -1：域名未关联证书。
	// 1： 域名https已开启。
	// 0： 域名https已关闭。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitnil,omitempty" name:"IsCache"`
}

type DeployCertificateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 待部署的证书ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 需要部署实例列表
	InstanceIdList []*string `json:"InstanceIdList,omitnil,omitempty" name:"InstanceIdList"`

	// 部署的云资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 部署云资源状态：
	// 云直播：
	// -1：域名未关联证书。
	// 1： 域名https已开启。
	// 0： 域名https已关闭。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitnil,omitempty" name:"IsCache"`
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
	delete(f, "IsCache")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeployCertificateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeployCertificateInstanceResponseParams struct {
	// 云资源部署任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployRecordId *uint64 `json:"DeployRecordId,omitnil,omitempty" name:"DeployRecordId"`

	// 部署状态，1表示部署成功，0表示部署失败
	DeployStatus *int64 `json:"DeployStatus,omitnil,omitempty" name:"DeployStatus"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	DeployRecordId *int64 `json:"DeployRecordId,omitnil,omitempty" name:"DeployRecordId"`

	// 待重试部署记录详情ID
	DeployRecordDetailId *int64 `json:"DeployRecordDetailId,omitnil,omitempty" name:"DeployRecordDetailId"`
}

type DeployCertificateRecordRetryRequest struct {
	*tchttp.BaseRequest
	
	// 待重试部署记录ID
	DeployRecordId *int64 `json:"DeployRecordId,omitnil,omitempty" name:"DeployRecordId"`

	// 待重试部署记录详情ID
	DeployRecordDetailId *int64 `json:"DeployRecordDetailId,omitnil,omitempty" name:"DeployRecordDetailId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	DeployRecordId *int64 `json:"DeployRecordId,omitnil,omitempty" name:"DeployRecordId"`
}

type DeployCertificateRecordRollbackRequest struct {
	*tchttp.BaseRequest
	
	// 待重试部署记录ID
	DeployRecordId *int64 `json:"DeployRecordId,omitnil,omitempty" name:"DeployRecordId"`
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
	DeployRecordId *int64 `json:"DeployRecordId,omitnil,omitempty" name:"DeployRecordId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 部署证书ID
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// 原绑定证书ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	OldCertId *string `json:"OldCertId,omitnil,omitempty" name:"OldCertId"`

	// 部署实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 部署实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 部署监听器ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 部署域名列表
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// 部署监听器协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 部署状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 部署错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 部署记录详情创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 部署记录详情最后一次更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 部署监听器名称
	ListenerName *string `json:"ListenerName,omitnil,omitempty" name:"ListenerName"`

	// 是否开启SNI
	SniSwitch *int64 `json:"SniSwitch,omitnil,omitempty" name:"SniSwitch"`

	// COS存储桶名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// secret名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// 端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// TCB环境ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 部署的TCB类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TCBType *string `json:"TCBType,omitnil,omitempty" name:"TCBType"`

	// 部署的TCB地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 部署CLB监听器的Url
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url []*string `json:"Url,omitnil,omitempty" name:"Url"`
}

type DeployRecordInfo struct {
	// 部署记录ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 部署证书ID
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// 部署资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 部署地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 部署状态
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 部署时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最近一次更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type DeployedResources struct {
	// 证书ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 数量
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 资源标识:clb,cdn,live,waf,antiddos
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 不建议使用。字段返回和Resources相同。本字段后续只返回null
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 关联资源ID或关联域名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resources []*string `json:"Resources,omitnil,omitempty" name:"Resources"`
}

// Predefined struct for user
type DescribeCertificateBindResourceTaskDetailRequestParams struct {
	// 任务ID，根据任务ID查询绑定云资源结果
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 每页展示数量， 默认10，最大值100; 分页总数为云资源地域下实例总数， 即第一页会拉群每个云资源的地域下面Limit数量实例
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 当前偏移量
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询资源类型的结果详情， 不传则查询所有
	ResourceTypes []*string `json:"ResourceTypes,omitnil,omitempty" name:"ResourceTypes"`

	// 查询地域列表的数据，CLB、TKE、WAF、APIGATEWAY、TCB支持地域查询， 其他资源类型不支持
	Regions []*string `json:"Regions,omitnil,omitempty" name:"Regions"`
}

type DescribeCertificateBindResourceTaskDetailRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID，根据任务ID查询绑定云资源结果
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 每页展示数量， 默认10，最大值100; 分页总数为云资源地域下实例总数， 即第一页会拉群每个云资源的地域下面Limit数量实例
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 当前偏移量
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询资源类型的结果详情， 不传则查询所有
	ResourceTypes []*string `json:"ResourceTypes,omitnil,omitempty" name:"ResourceTypes"`

	// 查询地域列表的数据，CLB、TKE、WAF、APIGATEWAY、TCB支持地域查询， 其他资源类型不支持
	Regions []*string `json:"Regions,omitnil,omitempty" name:"Regions"`
}

func (r *DescribeCertificateBindResourceTaskDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificateBindResourceTaskDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "ResourceTypes")
	delete(f, "Regions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCertificateBindResourceTaskDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertificateBindResourceTaskDetailResponseParams struct {
	// 关联clb资源详情	
	// 注意：此字段可能返回 null，表示取不到有效值。
	CLB []*ClbInstanceList `json:"CLB,omitnil,omitempty" name:"CLB"`

	// 关联cdn资源详情	
	// 注意：此字段可能返回 null，表示取不到有效值。
	CDN []*CdnInstanceList `json:"CDN,omitnil,omitempty" name:"CDN"`

	// 关联waf资源详情	
	// 注意：此字段可能返回 null，表示取不到有效值。
	WAF []*WafInstanceList `json:"WAF,omitnil,omitempty" name:"WAF"`

	// 关联ddos资源详情	
	// 注意：此字段可能返回 null，表示取不到有效值。
	DDOS []*DdosInstanceList `json:"DDOS,omitnil,omitempty" name:"DDOS"`

	// 关联live资源详情	
	// 注意：此字段可能返回 null，表示取不到有效值。
	LIVE []*LiveInstanceList `json:"LIVE,omitnil,omitempty" name:"LIVE"`

	// 关联vod资源详情	
	// 注意：此字段可能返回 null，表示取不到有效值。
	VOD []*VODInstanceList `json:"VOD,omitnil,omitempty" name:"VOD"`

	// 关联tke资源详情	
	// 注意：此字段可能返回 null，表示取不到有效值。
	TKE []*TkeInstanceList `json:"TKE,omitnil,omitempty" name:"TKE"`

	// 关联apigateway资源详情	
	// 注意：此字段可能返回 null，表示取不到有效值。
	APIGATEWAY []*ApiGatewayInstanceList `json:"APIGATEWAY,omitnil,omitempty" name:"APIGATEWAY"`

	// 关联tcb资源详情	
	// 注意：此字段可能返回 null，表示取不到有效值。
	TCB []*TCBInstanceList `json:"TCB,omitnil,omitempty" name:"TCB"`

	// 关联teo资源详情	
	// 注意：此字段可能返回 null，表示取不到有效值。
	TEO []*TeoInstanceList `json:"TEO,omitnil,omitempty" name:"TEO"`

	// 关联云资源异步查询结果： 0表示查询中， 1表示查询成功。 2表示查询异常； 若状态为1，则查看BindResourceResult结果；若状态为2，则查看Error原因
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 当前结果缓存时间
	CacheTime *string `json:"CacheTime,omitnil,omitempty" name:"CacheTime"`

	// 关联tse资源详情	
	// 注意：此字段可能返回 null，表示取不到有效值。
	TSE []*TSEInstanceList `json:"TSE,omitnil,omitempty" name:"TSE"`

	// 关联的COS资源详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	COS []*COSInstanceList `json:"COS,omitnil,omitempty" name:"COS"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCertificateBindResourceTaskDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCertificateBindResourceTaskDetailResponseParams `json:"Response"`
}

func (r *DescribeCertificateBindResourceTaskDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificateBindResourceTaskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertificateBindResourceTaskResultRequestParams struct {
	// 任务ID，根据任务ID查询绑定云资源结果， 最大支持100个
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`
}

type DescribeCertificateBindResourceTaskResultRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID，根据任务ID查询绑定云资源结果， 最大支持100个
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`
}

func (r *DescribeCertificateBindResourceTaskResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificateBindResourceTaskResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCertificateBindResourceTaskResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertificateBindResourceTaskResultResponseParams struct {
	// 异步任务绑定关联云资源结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SyncTaskBindResourceResult []*SyncTaskBindResourceResult `json:"SyncTaskBindResourceResult,omitnil,omitempty" name:"SyncTaskBindResourceResult"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCertificateBindResourceTaskResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCertificateBindResourceTaskResultResponseParams `json:"Response"`
}

func (r *DescribeCertificateBindResourceTaskResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificateBindResourceTaskResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertificateDetailRequestParams struct {
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`
}

type DescribeCertificateDetailRequest struct {
	*tchttp.BaseRequest
	
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`
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
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 项目 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 证书来源：trustasia = 亚洲诚信，upload = 用户上传。
	// 注意：此字段可能返回 null，表示取不到有效值。
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 证书类型：CA = 客户端证书，SVR = 服务器证书。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertificateType *string `json:"CertificateType,omitnil,omitempty" name:"CertificateType"`

	// 证书套餐类型：null = 用户上传证书（没有套餐类型），1 = GeoTrust DV SSL CA - G3， 2 = TrustAsia TLS RSA CA， 3 = SecureSite 增强型企业版（EV Pro）， 4 = SecureSite 增强型（EV）， 5 = SecureSite 企业型专业版（OV Pro）， 6 = SecureSite 企业型（OV）， 7 = SecureSite 企业型（OV）通配符， 8 = Geotrust 增强型（EV）， 9 = Geotrust 企业型（OV）， 10 = Geotrust 企业型（OV）通配符， 11 = TrustAsia 域名型多域名 SSL 证书， 12 = TrustAsia 域名型（DV）通配符， 13 = TrustAsia 企业型通配符（OV）SSL 证书（D3）， 14 = TrustAsia 企业型（OV）SSL 证书（D3）， 15 = TrustAsia 企业型多域名 （OV）SSL 证书（D3）， 16 = TrustAsia 增强型 （EV）SSL 证书（D3）， 17 = TrustAsia 增强型多域名（EV）SSL 证书（D3）， 18 = GlobalSign 企业型（OV）SSL 证书， 19 = GlobalSign 企业型通配符 （OV）SSL 证书， 20 = GlobalSign 增强型 （EV）SSL 证书， 21 = TrustAsia 企业型通配符多域名（OV）SSL 证书（D3）， 22 = GlobalSign 企业型多域名（OV）SSL 证书， 23 = GlobalSign 企业型通配符多域名（OV）SSL 证书， 24 = GlobalSign 增强型多域名（EV）SSL 证书，25 = Wotrus 域名型证书，26 = Wotrus 域名型多域名证书，27 = Wotrus 域名型通配符证书，28 = Wotrus 企业型证书，29 = Wotrus 企业型多域名证书，30 = Wotrus 企业型通配符证书，31 = Wotrus 增强型证书，32 = Wotrus 增强型多域名证书，33 = DNSPod 国密域名型证书，34 = DNSPod 国密域名型多域名证书，35 = DNSPod 国密域名型通配符证书，37 = DNSPod 国密企业型证书，38 = DNSPod 国密企业型多域名证书，39 = DNSPod 国密企业型通配符证书，40 = DNSPod 国密增强型证书，41 = DNSPod 国密增强型多域名证书，42 = TrustAsia 域名型通配符多域名证书。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// 颁发者。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductZhName *string `json:"ProductZhName,omitnil,omitempty" name:"ProductZhName"`

	// 域名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 备注名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 证书状态：0 = 审核中，1 = 已通过，2 = 审核失败，3 = 已过期，4 = 已添加DNS记录，5 = 企业证书，待提交，6 = 订单取消中，7 = 已取消，8 = 已提交资料， 待上传确认函，9 = 证书吊销中，10 = 已吊销，11 = 重颁发中，12 = 待上传吊销确认函，13 = 免费证书待提交资料。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 状态信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusMsg *string `json:"StatusMsg,omitnil,omitempty" name:"StatusMsg"`

	// 验证类型：DNS_AUTO = 自动DNS验证，DNS = 手动DNS验证，FILE = 文件验证，EMAIL = 邮件验证。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerifyType *string `json:"VerifyType,omitnil,omitempty" name:"VerifyType"`

	// 漏洞扫描状态。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulnerabilityStatus *string `json:"VulnerabilityStatus,omitnil,omitempty" name:"VulnerabilityStatus"`

	// 证书生效时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertBeginTime *string `json:"CertBeginTime,omitnil,omitempty" name:"CertBeginTime"`

	// 证书失效时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertEndTime *string `json:"CertEndTime,omitnil,omitempty" name:"CertEndTime"`

	// 证书有效期：单位（月）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValidityPeriod *string `json:"ValidityPeriod,omitnil,omitempty" name:"ValidityPeriod"`

	// 申请时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InsertTime *string `json:"InsertTime,omitnil,omitempty" name:"InsertTime"`

	// 订单 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// 证书扩展信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertificateExtra *CertificateExtra `json:"CertificateExtra,omitnil,omitempty" name:"CertificateExtra"`

	// 证书私钥
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertificatePrivateKey *string `json:"CertificatePrivateKey,omitnil,omitempty" name:"CertificatePrivateKey"`

	// 证书公钥（即证书内容）
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertificatePublicKey *string `json:"CertificatePublicKey,omitnil,omitempty" name:"CertificatePublicKey"`

	// DV 认证信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DvAuthDetail *DvAuthDetail `json:"DvAuthDetail,omitnil,omitempty" name:"DvAuthDetail"`

	// 漏洞扫描评估报告。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulnerabilityReport *string `json:"VulnerabilityReport,omitnil,omitempty" name:"VulnerabilityReport"`

	// 证书 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 证书类型名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TypeName *string `json:"TypeName,omitnil,omitempty" name:"TypeName"`

	// 状态描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusName *string `json:"StatusName,omitnil,omitempty" name:"StatusName"`

	// 证书包含的多个域名（不包含主域名，主域名使用Domain字段）
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubjectAltName []*string `json:"SubjectAltName,omitnil,omitempty" name:"SubjectAltName"`

	// 是否为付费证书。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsVip *bool `json:"IsVip,omitnil,omitempty" name:"IsVip"`

	// 是否为泛域名证书。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsWildcard *bool `json:"IsWildcard,omitnil,omitempty" name:"IsWildcard"`

	// 是否为 DV 版证书。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDv *bool `json:"IsDv,omitnil,omitempty" name:"IsDv"`

	// 是否启用了漏洞扫描功能。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsVulnerability *bool `json:"IsVulnerability,omitnil,omitempty" name:"IsVulnerability"`

	// 提交的资料信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubmittedData *SubmittedData `json:"SubmittedData,omitnil,omitempty" name:"SubmittedData"`

	// 是否可续费。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewAble *bool `json:"RenewAble,omitnil,omitempty" name:"RenewAble"`

	// 是否可部署。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Deployable *bool `json:"Deployable,omitnil,omitempty" name:"Deployable"`

	// 关联标签列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tags `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 根证书。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RootCert *RootCertificates `json:"RootCert,omitnil,omitempty" name:"RootCert"`

	// 国密加密证书
	// 注意：此字段可能返回 null，表示取不到有效值。
	EncryptCert *string `json:"EncryptCert,omitnil,omitempty" name:"EncryptCert"`

	// 国密加密私钥
	// 注意：此字段可能返回 null，表示取不到有效值。
	EncryptPrivateKey *string `json:"EncryptPrivateKey,omitnil,omitempty" name:"EncryptPrivateKey"`

	// 签名证书 SHA1指纹
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertFingerprint *string `json:"CertFingerprint,omitnil,omitempty" name:"CertFingerprint"`

	// 加密证书 SHA1指纹 （国密证书特有）
	// 注意：此字段可能返回 null，表示取不到有效值。
	EncryptCertFingerprint *string `json:"EncryptCertFingerprint,omitnil,omitempty" name:"EncryptCertFingerprint"`

	// 证书算法
	// 注意：此字段可能返回 null，表示取不到有效值。
	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitnil,omitempty" name:"EncryptAlgorithm"`

	// DV证书吊销验证值
	// 注意：此字段可能返回 null，表示取不到有效值。
	DvRevokeAuthDetail []*DvAuths `json:"DvRevokeAuthDetail,omitnil,omitempty" name:"DvRevokeAuthDetail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 请求日志数量，默认为20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 开始时间，默认15天前。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，默认现在时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeCertificateOperateLogsRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 请求日志数量，默认为20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 开始时间，默认15天前。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，默认现在时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
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
	AllTotal *uint64 `json:"AllTotal,omitnil,omitempty" name:"AllTotal"`

	// 本次请求返回的日志数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 证书操作日志列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperateLogs []*OperationLog `json:"OperateLogs,omitnil,omitempty" name:"OperateLogs"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`
}

type DescribeCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`
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
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 项目 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 证书来源：trustasia = 亚洲诚信，upload = 用户上传。
	// 注意：此字段可能返回 null，表示取不到有效值。
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 证书类型：CA = 客户端证书，SVR = 服务器证书。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertificateType *string `json:"CertificateType,omitnil,omitempty" name:"CertificateType"`

	// 证书套餐类型： 2 = TrustAsia TLS RSA CA，3 = SecureSite 增强型企业版（EV Pro）， 4 = SecureSite 增强型（EV）， 5 = SecureSite 企业型专业版（OV Pro）， 6 = SecureSite 企业型（OV）， 7 = SecureSite 企业型（OV）通配符， 8 = Geotrust 增强型（EV）， 9 = Geotrust 企业型（OV）， 10 = Geotrust 企业型（OV）通配符， 11 = TrustAsia 域名型多域名 SSL 证书， 12 = TrustAsia 域名型（DV）通配符， 13 = TrustAsia 企业型通配符（OV）SSL 证书（D3）， 14 = TrustAsia 企业型（OV）SSL 证书（D3）， 15 = TrustAsia 企业型多域名 （OV）SSL 证书（D3）， 16 = TrustAsia 增强型 （EV）SSL 证书（D3）， 17 = TrustAsia 增强型多域名（EV）SSL 证书（D3）， 18 = GlobalSign 企业型（OV）SSL 证书， 19 = GlobalSign 企业型通配符 （OV）SSL 证书， 20 = GlobalSign 增强型 （EV）SSL 证书， 21 = TrustAsia 企业型通配符多域名（OV）SSL 证书（D3）， 22 = GlobalSign 企业型多域名（OV）SSL 证书， 23 = GlobalSign 企业型通配符多域名（OV）SSL 证书， 24 = GlobalSign 增强型多域名（EV）SSL 证书，25 = Wotrus 域名型证书，26 = Wotrus 域名型多域名证书，27 = Wotrus 域名型通配符证书，28 = Wotrus 企业型证书，29 = Wotrus 企业型多域名证书，30 = Wotrus 企业型通配符证书，31 = Wotrus 增强型证书，32 = Wotrus 增强型多域名证书，33 = Wotrus 国密域名型证书，34 = Wotrus 国密域名型多域名证书，35 = Wotrus 国密域名型通配符证书，37 = Wotrus 国密企业型证书，38 = Wotrus 国密企业型多域名证书，39 = Wotrus 国密企业型通配符证书，40 = Wotrus 国密增强型证书，41 = Wotrus 国密增强型多域名证书，42 = TrustAsia 域名型通配符多域名证书，43 = DNSPod-企业型(OV)SSL证书，44 = DNSPod-企业型(OV)通配符SSL证书，45 = DNSPod-企业型(OV)多域名SSL证书， 46 = DNSPod-增强型(EV)SSL证书，47 = DNSPod-增强型(EV)多域名SSL证书，48 = DNSPod-域名型(DV)SSL证书，49 = DNSPod-域名型(DV)通配符SSL证书，50 = DNSPod-域名型(DV)多域名SSL证书，51 = DNSPod（国密）-企业型(OV)SSL证书，52 = DNSPod（国密）-企业型(OV)通配符SSL证书，53 = DNSPod（国密）-企业型(OV)多域名SSL证书，54 = DNSPod（国密）-域名型(DV)SSL证书，55 = DNSPod（国密）-域名型(DV)通配符SSL证书， 56 = DNSPod（国密）-域名型(DV)多域名SSL证书，57 = SecureSite 企业型专业版多域名(OV Pro)，58 = SecureSite 企业型多域名(OV)，59 = SecureSite 增强型专业版多域名(EV Pro)，60 = SecureSite 增强型多域名(EV)，61 = Geotrust 增强型多域名(EV)
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// 证书颁发者名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductZhName *string `json:"ProductZhName,omitnil,omitempty" name:"ProductZhName"`

	// 域名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 备注名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 证书状态：0 = 审核中，1 = 已通过，2 = 审核失败，3 = 已过期，4 = 已添加DNS记录，5 = 企业证书，待提交，6 = 订单取消中，7 = 已取消，8 = 已提交资料， 待上传确认函，9 = 证书吊销中，10 = 已吊销，11 = 重颁发中，12 = 待上传吊销确认函。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 状态信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusMsg *string `json:"StatusMsg,omitnil,omitempty" name:"StatusMsg"`

	// 验证类型：DNS_AUTO = 自动DNS验证，DNS = 手动DNS验证，FILE = 文件验证，EMAIL = 邮件验证。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerifyType *string `json:"VerifyType,omitnil,omitempty" name:"VerifyType"`

	// 漏洞扫描状态。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulnerabilityStatus *string `json:"VulnerabilityStatus,omitnil,omitempty" name:"VulnerabilityStatus"`

	// 证书生效时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertBeginTime *string `json:"CertBeginTime,omitnil,omitempty" name:"CertBeginTime"`

	// 证书失效时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertEndTime *string `json:"CertEndTime,omitnil,omitempty" name:"CertEndTime"`

	// 证书有效期：单位(月)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValidityPeriod *string `json:"ValidityPeriod,omitnil,omitempty" name:"ValidityPeriod"`

	// 申请时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InsertTime *string `json:"InsertTime,omitnil,omitempty" name:"InsertTime"`

	// 订单 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// 证书扩展信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertificateExtra *CertificateExtra `json:"CertificateExtra,omitnil,omitempty" name:"CertificateExtra"`

	// DV 认证信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DvAuthDetail *DvAuthDetail `json:"DvAuthDetail,omitnil,omitempty" name:"DvAuthDetail"`

	// 漏洞扫描评估报告。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulnerabilityReport *string `json:"VulnerabilityReport,omitnil,omitempty" name:"VulnerabilityReport"`

	// 证书 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 证书类型名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageTypeName *string `json:"PackageTypeName,omitnil,omitempty" name:"PackageTypeName"`

	// 状态描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusName *string `json:"StatusName,omitnil,omitempty" name:"StatusName"`

	// 证书包含的多个域名（包含主域名）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubjectAltName []*string `json:"SubjectAltName,omitnil,omitempty" name:"SubjectAltName"`

	// 是否为 VIP 客户。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsVip *bool `json:"IsVip,omitnil,omitempty" name:"IsVip"`

	// 是否为泛域名证书。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsWildcard *bool `json:"IsWildcard,omitnil,omitempty" name:"IsWildcard"`

	// 是否为 DV 版证书。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDv *bool `json:"IsDv,omitnil,omitempty" name:"IsDv"`

	// 是否启用了漏洞扫描功能。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsVulnerability *bool `json:"IsVulnerability,omitnil,omitempty" name:"IsVulnerability"`

	// 是否可重颁发证书。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewAble *bool `json:"RenewAble,omitnil,omitempty" name:"RenewAble"`

	// 提交的资料信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubmittedData *SubmittedData `json:"SubmittedData,omitnil,omitempty" name:"SubmittedData"`

	// 是否可部署。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Deployable *bool `json:"Deployable,omitnil,omitempty" name:"Deployable"`

	// 标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tags `json:"Tags,omitnil,omitempty" name:"Tags"`

	// CA证书的所有加密方式	
	// 注意：此字段可能返回 null，表示取不到有效值。
	CAEncryptAlgorithms []*string `json:"CAEncryptAlgorithms,omitnil,omitempty" name:"CAEncryptAlgorithms"`

	// CA证书的所有通用名称	
	// 注意：此字段可能返回 null，表示取不到有效值。
	CACommonNames []*string `json:"CACommonNames,omitnil,omitempty" name:"CACommonNames"`

	// CA证书所有的到期时间	
	// 注意：此字段可能返回 null，表示取不到有效值。
	CAEndTimes []*string `json:"CAEndTimes,omitnil,omitempty" name:"CAEndTimes"`

	// DV证书吊销验证值
	// 注意：此字段可能返回 null，表示取不到有效值。
	DvRevokeAuthDetail []*DvAuths `json:"DvRevokeAuthDetail,omitnil,omitempty" name:"DvRevokeAuthDetail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量，默认10。最大1000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 搜索关键词，可搜索证书 ID、备注名称、域名。例如： a8xHcaIs。
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 证书类型：CA = 客户端证书，SVR = 服务器证书。
	CertificateType *string `json:"CertificateType,omitnil,omitempty" name:"CertificateType"`

	// 项目 ID。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 按到期时间排序：DESC = 降序， ASC = 升序。
	ExpirationSort *string `json:"ExpirationSort,omitnil,omitempty" name:"ExpirationSort"`

	// 证书状态：0 = 审核中，1 = 已通过，2 = 审核失败，3 = 已过期，4 = 已添加DNS记录，5 = 企业证书，待提交，6 = 订单取消中，7 = 已取消，8 = 已提交资料， 待上传确认函，9 = 证书吊销中，10 = 已吊销，11 = 重颁发中，12 = 待上传吊销确认函，13 = 免费证书待提交资料。
	CertificateStatus []*uint64 `json:"CertificateStatus,omitnil,omitempty" name:"CertificateStatus"`

	// 是否可部署，可选值：1 = 可部署，0 =  不可部署。
	Deployable *uint64 `json:"Deployable,omitnil,omitempty" name:"Deployable"`

	// 是否筛选上传托管的 1筛选，0不筛选
	Upload *int64 `json:"Upload,omitnil,omitempty" name:"Upload"`

	// 是否筛选可续期证书 1筛选 0不筛选
	Renew *int64 `json:"Renew,omitnil,omitempty" name:"Renew"`

	// 筛选来源， upload：上传证书， buy：腾讯云证书， 不传默认全部
	FilterSource *string `json:"FilterSource,omitnil,omitempty" name:"FilterSource"`

	// 是否筛选国密证书。1:筛选  0:不筛选
	IsSM *int64 `json:"IsSM,omitnil,omitempty" name:"IsSM"`

	// 筛选证书是否即将过期，传1是筛选，0不筛选
	FilterExpiring *uint64 `json:"FilterExpiring,omitnil,omitempty" name:"FilterExpiring"`

	// 是否可托管，可选值：1 = 可托管，0 =  不可托管。
	Hostable *uint64 `json:"Hostable,omitnil,omitempty" name:"Hostable"`

	// 筛选指定标签的证书
	Tags []*Tags `json:"Tags,omitnil,omitempty" name:"Tags"`

	// //是否筛选等待签发的证书，传1是筛选，0和null不筛选
	IsPendingIssue *int64 `json:"IsPendingIssue,omitnil,omitempty" name:"IsPendingIssue"`
}

type DescribeCertificatesRequest struct {
	*tchttp.BaseRequest
	
	// 分页偏移量，从0开始。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量，默认10。最大1000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 搜索关键词，可搜索证书 ID、备注名称、域名。例如： a8xHcaIs。
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 证书类型：CA = 客户端证书，SVR = 服务器证书。
	CertificateType *string `json:"CertificateType,omitnil,omitempty" name:"CertificateType"`

	// 项目 ID。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 按到期时间排序：DESC = 降序， ASC = 升序。
	ExpirationSort *string `json:"ExpirationSort,omitnil,omitempty" name:"ExpirationSort"`

	// 证书状态：0 = 审核中，1 = 已通过，2 = 审核失败，3 = 已过期，4 = 已添加DNS记录，5 = 企业证书，待提交，6 = 订单取消中，7 = 已取消，8 = 已提交资料， 待上传确认函，9 = 证书吊销中，10 = 已吊销，11 = 重颁发中，12 = 待上传吊销确认函，13 = 免费证书待提交资料。
	CertificateStatus []*uint64 `json:"CertificateStatus,omitnil,omitempty" name:"CertificateStatus"`

	// 是否可部署，可选值：1 = 可部署，0 =  不可部署。
	Deployable *uint64 `json:"Deployable,omitnil,omitempty" name:"Deployable"`

	// 是否筛选上传托管的 1筛选，0不筛选
	Upload *int64 `json:"Upload,omitnil,omitempty" name:"Upload"`

	// 是否筛选可续期证书 1筛选 0不筛选
	Renew *int64 `json:"Renew,omitnil,omitempty" name:"Renew"`

	// 筛选来源， upload：上传证书， buy：腾讯云证书， 不传默认全部
	FilterSource *string `json:"FilterSource,omitnil,omitempty" name:"FilterSource"`

	// 是否筛选国密证书。1:筛选  0:不筛选
	IsSM *int64 `json:"IsSM,omitnil,omitempty" name:"IsSM"`

	// 筛选证书是否即将过期，传1是筛选，0不筛选
	FilterExpiring *uint64 `json:"FilterExpiring,omitnil,omitempty" name:"FilterExpiring"`

	// 是否可托管，可选值：1 = 可托管，0 =  不可托管。
	Hostable *uint64 `json:"Hostable,omitnil,omitempty" name:"Hostable"`

	// 筛选指定标签的证书
	Tags []*Tags `json:"Tags,omitnil,omitempty" name:"Tags"`

	// //是否筛选等待签发的证书，传1是筛选，0和null不筛选
	IsPendingIssue *int64 `json:"IsPendingIssue,omitnil,omitempty" name:"IsPendingIssue"`
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
	delete(f, "Tags")
	delete(f, "IsPendingIssue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCertificatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertificatesResponseParams struct {
	// 总数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Certificates []*Certificates `json:"Certificates,omitnil,omitempty" name:"Certificates"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页每页限制数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 公司ID
	CompanyId *int64 `json:"CompanyId,omitnil,omitempty" name:"CompanyId"`
}

type DescribeCompaniesRequest struct {
	*tchttp.BaseRequest
	
	// 分页偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页每页限制数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 公司ID
	CompanyId *int64 `json:"CompanyId,omitnil,omitempty" name:"CompanyId"`
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
	Companies []*CompanyInfo `json:"Companies,omitnil,omitempty" name:"Companies"`

	// 公司总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeDeleteCertificatesTaskResultRequestParams struct {
	// DeleteCertificates接口返回的任务ID， 最大支持100个
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`
}

type DescribeDeleteCertificatesTaskResultRequest struct {
	*tchttp.BaseRequest
	
	// DeleteCertificates接口返回的任务ID， 最大支持100个
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`
}

func (r *DescribeDeleteCertificatesTaskResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeleteCertificatesTaskResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeleteCertificatesTaskResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeleteCertificatesTaskResultResponseParams struct {
	// 批量删除证书异步任务结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteTaskResult []*DeleteTaskResult `json:"DeleteTaskResult,omitnil,omitempty" name:"DeleteTaskResult"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDeleteCertificatesTaskResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeleteCertificatesTaskResultResponseParams `json:"Response"`
}

func (r *DescribeDeleteCertificatesTaskResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeleteCertificatesTaskResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeployedResourcesRequestParams struct {
	// 证书ID
	CertificateIds []*string `json:"CertificateIds,omitnil,omitempty" name:"CertificateIds"`

	// 资源类型:clb,cdn,live,waf,antiddos
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`
}

type DescribeDeployedResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 证书ID
	CertificateIds []*string `json:"CertificateIds,omitnil,omitempty" name:"CertificateIds"`

	// 资源类型:clb,cdn,live,waf,antiddos
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`
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
	DeployedResources []*DeployedResources `json:"DeployedResources,omitnil,omitempty" name:"DeployedResources"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeDownloadCertificateUrlRequestParams struct {
	// 证书ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 下载的服务类型: nginx tomcat apache iis jks other root
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`
}

type DescribeDownloadCertificateUrlRequest struct {
	*tchttp.BaseRequest
	
	// 证书ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 下载的服务类型: nginx tomcat apache iis jks other root
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`
}

func (r *DescribeDownloadCertificateUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDownloadCertificateUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	delete(f, "ServiceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDownloadCertificateUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDownloadCertificateUrlResponseParams struct {
	// 下载链接
	DownloadCertificateUrl *string `json:"DownloadCertificateUrl,omitnil,omitempty" name:"DownloadCertificateUrl"`

	// 下载文件的名称
	DownloadFilename *string `json:"DownloadFilename,omitnil,omitempty" name:"DownloadFilename"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDownloadCertificateUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDownloadCertificateUrlResponseParams `json:"Response"`
}

func (r *DescribeDownloadCertificateUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDownloadCertificateUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostApiGatewayInstanceListRequestParams struct {
	// 待部署的证书ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 部署资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitnil,omitempty" name:"IsCache"`

	// 过滤参数列表； FilterKey：domainMatch（查询域名是否匹配的实例列表） FilterValue：1，表示查询匹配； 0，表示查询不匹配； 默认查询匹配
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 已部署的证书ID
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`
}

type DescribeHostApiGatewayInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 待部署的证书ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 部署资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitnil,omitempty" name:"IsCache"`

	// 过滤参数列表； FilterKey：domainMatch（查询域名是否匹配的实例列表） FilterValue：1，表示查询匹配； 0，表示查询不匹配； 默认查询匹配
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 已部署的证书ID
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`
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
	InstanceList []*ApiGatewayInstanceDetail `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 部署资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitnil,omitempty" name:"IsCache"`

	// 过滤参数列表； FilterKey：domainMatch（查询域名是否匹配的实例列表） FilterValue：1，表示查询匹配； 0，表示查询不匹配； 默认查询匹配
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 原证书ID
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`

	// 分页偏移量，从0开始。	
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量，默认10。	
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 是否异步
	AsyncCache *int64 `json:"AsyncCache,omitnil,omitempty" name:"AsyncCache"`
}

type DescribeHostCdnInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 待部署的证书ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 部署资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitnil,omitempty" name:"IsCache"`

	// 过滤参数列表； FilterKey：domainMatch（查询域名是否匹配的实例列表） FilterValue：1，表示查询匹配； 0，表示查询不匹配； 默认查询匹配
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 原证书ID
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`

	// 分页偏移量，从0开始。	
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量，默认10。	
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 是否异步
	AsyncCache *int64 `json:"AsyncCache,omitnil,omitempty" name:"AsyncCache"`
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
	InstanceList []*CdnInstanceDetail `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// CDN域名总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 异步刷新总数	
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsyncTotalNum *int64 `json:"AsyncTotalNum,omitnil,omitempty" name:"AsyncTotalNum"`

	// 异步刷新当前执行数	
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsyncOffset *int64 `json:"AsyncOffset,omitnil,omitempty" name:"AsyncOffset"`

	// 当前缓存读取时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsyncCacheTime *string `json:"AsyncCacheTime,omitnil,omitempty" name:"AsyncCacheTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 分页偏移量，从0开始。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量，默认10。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitnil,omitempty" name:"IsCache"`

	// 过滤参数列表； FilterKey：domainMatch（查询域名是否匹配的实例列表） FilterValue：1，表示查询匹配； 0，表示查询不匹配； 默认查询匹配
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否异步缓存
	AsyncCache *int64 `json:"AsyncCache,omitnil,omitempty" name:"AsyncCache"`

	// 原证书ID
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`
}

type DescribeHostClbInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 待部署的证书ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 分页偏移量，从0开始。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量，默认10。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitnil,omitempty" name:"IsCache"`

	// 过滤参数列表； FilterKey：domainMatch（查询域名是否匹配的实例列表） FilterValue：1，表示查询匹配； 0，表示查询不匹配； 默认查询匹配
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否异步缓存
	AsyncCache *int64 `json:"AsyncCache,omitnil,omitempty" name:"AsyncCache"`

	// 原证书ID
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// CLB实例监听器列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceList []*ClbInstanceDetail `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// 异步刷新总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsyncTotalNum *int64 `json:"AsyncTotalNum,omitnil,omitempty" name:"AsyncTotalNum"`

	// 异步刷新当前执行数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsyncOffset *int64 `json:"AsyncOffset,omitnil,omitempty" name:"AsyncOffset"`

	// 当前缓存读取时间	
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsyncCacheTime *string `json:"AsyncCacheTime,omitnil,omitempty" name:"AsyncCacheTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitnil,omitempty" name:"IsCache"`

	// 过滤参数列表
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 部署资源类型 cos
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 原证书ID	
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`

	// 分页偏移量，从0开始。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量，默认10。	
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 是否异步	
	AsyncCache *int64 `json:"AsyncCache,omitnil,omitempty" name:"AsyncCache"`
}

type DescribeHostCosInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 待部署的证书ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitnil,omitempty" name:"IsCache"`

	// 过滤参数列表
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 部署资源类型 cos
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 原证书ID	
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`

	// 分页偏移量，从0开始。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量，默认10。	
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 是否异步	
	AsyncCache *int64 `json:"AsyncCache,omitnil,omitempty" name:"AsyncCache"`
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
	delete(f, "IsCache")
	delete(f, "Filters")
	delete(f, "ResourceType")
	delete(f, "OldCertificateId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "AsyncCache")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostCosInstanceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostCosInstanceListResponseParams struct {
	// COS实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceList []*CosInstanceDetail `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 异步刷新总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsyncTotalNum *int64 `json:"AsyncTotalNum,omitnil,omitempty" name:"AsyncTotalNum"`

	// 异步刷新当前执行数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsyncOffset *int64 `json:"AsyncOffset,omitnil,omitempty" name:"AsyncOffset"`

	// 当前缓存读取时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsyncCacheTime *string `json:"AsyncCacheTime,omitnil,omitempty" name:"AsyncCacheTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 部署资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitnil,omitempty" name:"IsCache"`

	// 过滤参数列表； FilterKey：domainMatch（查询域名是否匹配的实例列表） FilterValue：1，表示查询匹配； 0，表示查询不匹配； 默认查询匹配
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 已部署的证书ID
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`
}

type DescribeHostDdosInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 待部署的证书ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 部署资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitnil,omitempty" name:"IsCache"`

	// 过滤参数列表； FilterKey：domainMatch（查询域名是否匹配的实例列表） FilterValue：1，表示查询匹配； 0，表示查询不匹配； 默认查询匹配
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 已部署的证书ID
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`
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
	InstanceList []*DdosInstanceDetail `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	DeployRecordId *string `json:"DeployRecordId,omitnil,omitempty" name:"DeployRecordId"`

	// 分页偏移量，从0开始。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量，默认10。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeHostDeployRecordDetailRequest struct {
	*tchttp.BaseRequest
	
	// 部署记录ID
	DeployRecordId *string `json:"DeployRecordId,omitnil,omitempty" name:"DeployRecordId"`

	// 分页偏移量，从0开始。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量，默认10。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 证书部署记录列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployRecordDetailList []*DeployRecordDetail `json:"DeployRecordDetailList,omitnil,omitempty" name:"DeployRecordDetailList"`

	// 成功总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessTotalCount *int64 `json:"SuccessTotalCount,omitnil,omitempty" name:"SuccessTotalCount"`

	// 失败总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedTotalCount *int64 `json:"FailedTotalCount,omitnil,omitempty" name:"FailedTotalCount"`

	// 部署中总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunningTotalCount *int64 `json:"RunningTotalCount,omitnil,omitempty" name:"RunningTotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 分页偏移量，从0开始。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量，默认10。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`
}

type DescribeHostDeployRecordRequest struct {
	*tchttp.BaseRequest
	
	// 待部署的证书ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 分页偏移量，从0开始。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量，默认10。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 证书部署记录列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployRecordList []*DeployRecordInfo `json:"DeployRecordList,omitnil,omitempty" name:"DeployRecordList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 部署资源类型 lighthouse
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitnil,omitempty" name:"IsCache"`

	// 过滤参数列表
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeHostLighthouseInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 待部署的证书ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 部署资源类型 lighthouse
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitnil,omitempty" name:"IsCache"`

	// 过滤参数列表
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	InstanceList []*LighthouseInstanceDetail `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 部署资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitnil,omitempty" name:"IsCache"`

	// 过滤参数列表； FilterKey：domainMatch（查询域名是否匹配的实例列表） FilterValue：1，表示查询匹配； 0，表示查询不匹配； 默认查询匹配
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 已部署的证书ID
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`
}

type DescribeHostLiveInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 待部署的证书ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 部署资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitnil,omitempty" name:"IsCache"`

	// 过滤参数列表； FilterKey：domainMatch（查询域名是否匹配的实例列表） FilterValue：1，表示查询匹配； 0，表示查询不匹配； 默认查询匹配
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 已部署的证书ID
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`
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
	InstanceList []*LiveInstanceDetail `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 部署资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitnil,omitempty" name:"IsCache"`

	// 过滤参数列表； FilterKey：domainMatch（查询域名是否匹配的实例列表） FilterValue：1，表示查询匹配； 0，表示查询不匹配； 默认查询匹配
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 已部署的证书ID
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`

	// 分页偏移量，从0开始。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量，默认10。	
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 是否异步
	AsyncCache *int64 `json:"AsyncCache,omitnil,omitempty" name:"AsyncCache"`
}

type DescribeHostTeoInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 待部署的证书ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 部署资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitnil,omitempty" name:"IsCache"`

	// 过滤参数列表； FilterKey：domainMatch（查询域名是否匹配的实例列表） FilterValue：1，表示查询匹配； 0，表示查询不匹配； 默认查询匹配
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 已部署的证书ID
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`

	// 分页偏移量，从0开始。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量，默认10。	
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 是否异步
	AsyncCache *int64 `json:"AsyncCache,omitnil,omitempty" name:"AsyncCache"`
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
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "AsyncCache")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostTeoInstanceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostTeoInstanceListResponseParams struct {
	// teo实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceList []*TeoInstanceDetail `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 分页偏移量，从0开始。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量，默认10。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitnil,omitempty" name:"IsCache"`

	// 过滤参数列表； FilterKey：domainMatch（查询域名是否匹配的实例列表） FilterValue：1，表示查询匹配； 0，表示查询不匹配； 默认查询匹配
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否异步缓存
	AsyncCache *int64 `json:"AsyncCache,omitnil,omitempty" name:"AsyncCache"`

	// 原证书ID
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`
}

type DescribeHostTkeInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 待部署的证书ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 分页偏移量，从0开始。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量，默认10。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitnil,omitempty" name:"IsCache"`

	// 过滤参数列表； FilterKey：domainMatch（查询域名是否匹配的实例列表） FilterValue：1，表示查询匹配； 0，表示查询不匹配； 默认查询匹配
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否异步缓存
	AsyncCache *int64 `json:"AsyncCache,omitnil,omitempty" name:"AsyncCache"`

	// 原证书ID
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// CLB实例监听器列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceList []*TkeInstanceDetail `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// 异步刷新总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsyncTotalNum *int64 `json:"AsyncTotalNum,omitnil,omitempty" name:"AsyncTotalNum"`

	// 异步刷新当前执行数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsyncOffset *int64 `json:"AsyncOffset,omitnil,omitempty" name:"AsyncOffset"`

	// 当前缓存读取时间	
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsyncCacheTime *string `json:"AsyncCacheTime,omitnil,omitempty" name:"AsyncCacheTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	DeployRecordId *string `json:"DeployRecordId,omitnil,omitempty" name:"DeployRecordId"`

	// 每页数量，默认10。
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量，从0开始。
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeHostUpdateRecordDetailRequest struct {
	*tchttp.BaseRequest
	
	// 一键更新记录ID
	DeployRecordId *string `json:"DeployRecordId,omitnil,omitempty" name:"DeployRecordId"`

	// 每页数量，默认10。
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量，从0开始。
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 证书部署记录列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordDetailList []*UpdateRecordDetails `json:"RecordDetailList,omitnil,omitempty" name:"RecordDetailList"`

	// 成功总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessTotalCount *int64 `json:"SuccessTotalCount,omitnil,omitempty" name:"SuccessTotalCount"`

	// 失败总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedTotalCount *int64 `json:"FailedTotalCount,omitnil,omitempty" name:"FailedTotalCount"`

	// 部署中总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunningTotalCount *int64 `json:"RunningTotalCount,omitnil,omitempty" name:"RunningTotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量，默认10。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 新证书ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 原证书ID
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`
}

type DescribeHostUpdateRecordRequest struct {
	*tchttp.BaseRequest
	
	// 分页偏移量，从0开始。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量，默认10。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 新证书ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 原证书ID
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 证书部署记录列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployRecordList []*UpdateRecordInfo `json:"DeployRecordList,omitnil,omitempty" name:"DeployRecordList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 部署资源类型 vod
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitnil,omitempty" name:"IsCache"`

	// 过滤参数列表
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 已部署的证书ID
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`
}

type DescribeHostVodInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 待部署的证书ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 部署资源类型 vod
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitnil,omitempty" name:"IsCache"`

	// 过滤参数列表
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 已部署的证书ID
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`
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
	InstanceList []*VodInstanceDetail `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 部署资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitnil,omitempty" name:"IsCache"`

	// 过滤参数列表； FilterKey：domainMatch（查询域名是否匹配的实例列表） FilterValue：1，表示查询匹配； 0，表示查询不匹配； 默认查询匹配
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 已部署的证书ID
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`
}

type DescribeHostWafInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 待部署的证书ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 部署资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 是否查询缓存，1：是； 0：否， 默认为查询缓存，缓存半小时
	IsCache *uint64 `json:"IsCache,omitnil,omitempty" name:"IsCache"`

	// 过滤参数列表； FilterKey：domainMatch（查询域名是否匹配的实例列表） FilterValue：1，表示查询匹配； 0，表示查询不匹配； 默认查询匹配
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 已部署的证书ID
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`
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
	InstanceList []*LiveInstanceDetail `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ManagerId *int64 `json:"ManagerId,omitnil,omitempty" name:"ManagerId"`

	// 分页每页数量
	//
	// Deprecated: Limit is deprecated.
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量
	//
	// Deprecated: Offset is deprecated.
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeManagerDetailRequest struct {
	*tchttp.BaseRequest
	
	// 管理人ID
	ManagerId *int64 `json:"ManagerId,omitnil,omitempty" name:"ManagerId"`

	// 分页每页数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
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
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 管理人姓名
	ManagerFirstName *string `json:"ManagerFirstName,omitnil,omitempty" name:"ManagerFirstName"`

	// 管理人邮箱
	ManagerMail *string `json:"ManagerMail,omitnil,omitempty" name:"ManagerMail"`

	// 联系人姓名
	ContactFirstName *string `json:"ContactFirstName,omitnil,omitempty" name:"ContactFirstName"`

	// 管理人姓名
	ManagerLastName *string `json:"ManagerLastName,omitnil,omitempty" name:"ManagerLastName"`

	// 联系人职位
	ContactPosition *string `json:"ContactPosition,omitnil,omitempty" name:"ContactPosition"`

	// 管理人职位
	ManagerPosition *string `json:"ManagerPosition,omitnil,omitempty" name:"ManagerPosition"`

	// 核验通过时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerifyTime *string `json:"VerifyTime,omitnil,omitempty" name:"VerifyTime"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 核验过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 联系人姓名
	ContactLastName *string `json:"ContactLastName,omitnil,omitempty" name:"ContactLastName"`

	// 管理人电话
	ManagerPhone *string `json:"ManagerPhone,omitnil,omitempty" name:"ManagerPhone"`

	// 联系人电话
	ContactPhone *string `json:"ContactPhone,omitnil,omitempty" name:"ContactPhone"`

	// 联系人邮箱
	ContactMail *string `json:"ContactMail,omitnil,omitempty" name:"ContactMail"`

	// 管理人所属部门
	ManagerDepartment *string `json:"ManagerDepartment,omitnil,omitempty" name:"ManagerDepartment"`

	// 管理人所属公司信息
	CompanyInfo *CompanyInfo `json:"CompanyInfo,omitnil,omitempty" name:"CompanyInfo"`

	// 管理人公司ID
	CompanyId *int64 `json:"CompanyId,omitnil,omitempty" name:"CompanyId"`

	// 管理人ID
	ManagerId *int64 `json:"ManagerId,omitnil,omitempty" name:"ManagerId"`

	// 审核状态详细信息
	StatusInfo []*ManagerStatusInfo `json:"StatusInfo,omitnil,omitempty" name:"StatusInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	CompanyId *int64 `json:"CompanyId,omitnil,omitempty" name:"CompanyId"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页每页数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 管理人姓名（将废弃），请使用SearchKey
	ManagerName *string `json:"ManagerName,omitnil,omitempty" name:"ManagerName"`

	// 模糊查询管理人邮箱（将废弃），请使用SearchKey
	ManagerMail *string `json:"ManagerMail,omitnil,omitempty" name:"ManagerMail"`

	// 根据管理人状态进行筛选，取值有
	// 'none' 未提交审核
	// 'audit', 亚信审核中
	// 'CAaudit' CA审核中
	// 'ok' 已审核
	// 'invalid'  审核失败
	// 'expiring'  即将过期
	// 'expired' 已过期
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 管理人姓/管理人名/邮箱/部门精准匹配
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`
}

type DescribeManagersRequest struct {
	*tchttp.BaseRequest
	
	// 公司ID
	CompanyId *int64 `json:"CompanyId,omitnil,omitempty" name:"CompanyId"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页每页数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 管理人姓名（将废弃），请使用SearchKey
	ManagerName *string `json:"ManagerName,omitnil,omitempty" name:"ManagerName"`

	// 模糊查询管理人邮箱（将废弃），请使用SearchKey
	ManagerMail *string `json:"ManagerMail,omitnil,omitempty" name:"ManagerMail"`

	// 根据管理人状态进行筛选，取值有
	// 'none' 未提交审核
	// 'audit', 亚信审核中
	// 'CAaudit' CA审核中
	// 'ok' 已审核
	// 'invalid'  审核失败
	// 'expiring'  即将过期
	// 'expired' 已过期
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 管理人姓/管理人名/邮箱/部门精准匹配
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`
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
	Managers []*ManagerInfo `json:"Managers,omitnil,omitempty" name:"Managers"`

	// 公司管理人总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目，默认20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 按状态筛选。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 按过期时间升序或降序排列。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 按权益包ID搜索。
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 按权益包类型搜索。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 子产品编号
	Pid *int64 `json:"Pid,omitnil,omitempty" name:"Pid"`
}

type DescribePackagesRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目，默认20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 按状态筛选。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 按过期时间升序或降序排列。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 按权益包ID搜索。
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 按权益包类型搜索。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 子产品编号
	Pid *int64 `json:"Pid,omitnil,omitempty" name:"Pid"`
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
	Packages []*PackageInfo `json:"Packages,omitnil,omitempty" name:"Packages"`

	// 总条数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 权益点总余额。
	TotalBalance *uint64 `json:"TotalBalance,omitnil,omitempty" name:"TotalBalance"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type DomainValidationResult struct {
	// 域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 验证类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerifyType *string `json:"VerifyType,omitnil,omitempty" name:"VerifyType"`

	// 本地检查结果。
	LocalCheck *int64 `json:"LocalCheck,omitnil,omitempty" name:"LocalCheck"`

	// CA检查结果。
	CaCheck *int64 `json:"CaCheck,omitnil,omitempty" name:"CaCheck"`

	// 检查失败原因。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocalCheckFailReason *string `json:"LocalCheckFailReason,omitnil,omitempty" name:"LocalCheckFailReason"`

	// 检查到的值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckValue []*string `json:"CheckValue,omitnil,omitempty" name:"CheckValue"`

	// 是否频繁请求。
	Frequently *bool `json:"Frequently,omitnil,omitempty" name:"Frequently"`

	// 是否已经签发。
	Issued *bool `json:"Issued,omitnil,omitempty" name:"Issued"`
}

// Predefined struct for user
type DownloadCertificateRequestParams struct {
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`
}

type DownloadCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`
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
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// MIME 类型：application/zip = ZIP 压缩文件。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContentType *string `json:"ContentType,omitnil,omitempty" name:"ContentType"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	DvAuthKey *string `json:"DvAuthKey,omitnil,omitempty" name:"DvAuthKey"`

	// DV 认证值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DvAuthValue *string `json:"DvAuthValue,omitnil,omitempty" name:"DvAuthValue"`

	// DV 认证值域名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DvAuthDomain *string `json:"DvAuthDomain,omitnil,omitempty" name:"DvAuthDomain"`

	// DV 认证值路径。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DvAuthPath *string `json:"DvAuthPath,omitnil,omitempty" name:"DvAuthPath"`

	// DV 认证子域名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DvAuthKeySubDomain *string `json:"DvAuthKeySubDomain,omitnil,omitempty" name:"DvAuthKeySubDomain"`

	// DV 认证信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DvAuths []*DvAuths `json:"DvAuths,omitnil,omitempty" name:"DvAuths"`
}

type DvAuths struct {
	// DV 认证密钥。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DvAuthKey *string `json:"DvAuthKey,omitnil,omitempty" name:"DvAuthKey"`

	// DV 认证值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DvAuthValue *string `json:"DvAuthValue,omitnil,omitempty" name:"DvAuthValue"`

	// DV 认证值域名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DvAuthDomain *string `json:"DvAuthDomain,omitnil,omitempty" name:"DvAuthDomain"`

	// DV 认证值路径。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DvAuthPath *string `json:"DvAuthPath,omitnil,omitempty" name:"DvAuthPath"`

	// DV 认证子域名，
	// 注意：此字段可能返回 null，表示取不到有效值。
	DvAuthSubDomain *string `json:"DvAuthSubDomain,omitnil,omitempty" name:"DvAuthSubDomain"`

	// DV 认证类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DvAuthVerifyType *string `json:"DvAuthVerifyType,omitnil,omitempty" name:"DvAuthVerifyType"`
}

type Error struct {
	// 异常错误码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 异常错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

type Filter struct {
	// 过滤参数key
	FilterKey *string `json:"FilterKey,omitnil,omitempty" name:"FilterKey"`

	// 过滤参数值
	FilterValue *string `json:"FilterValue,omitnil,omitempty" name:"FilterValue"`
}

type GatewayCertificate struct {
	// 网关证书ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 网关证书名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 绑定域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindDomains []*string `json:"BindDomains,omitnil,omitempty" name:"BindDomains"`

	// 证书来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertSource *string `json:"CertSource,omitnil,omitempty" name:"CertSource"`

	// 当前绑定的SSL证书ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`
}

type HostingConfig struct {
	// 托管资源替换时间， 默认为证书过期前30天存在续费证书则替换
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReplaceTime *int64 `json:"ReplaceTime,omitnil,omitempty" name:"ReplaceTime"`

	// 托管发送消息类型：0，托管开始前消息提醒（没有续费证书也会收到该提示消息）； 1， 托管开始消息提醒（存在续费证书才会收到消息提醒）； 2， 托管资源替换失败消息提醒； 3 托管资源替换成功消息提醒
	// 注意：此字段可能返回 null，表示取不到有效值。
	MessageTypes []*int64 `json:"MessageTypes,omitnil,omitempty" name:"MessageTypes"`
}

type LighthouseInstanceDetail struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// IP地址
	IP []*string `json:"IP,omitnil,omitempty" name:"IP"`

	// 可选择域名
	Domain []*string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

type LiveInstanceDetail struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 已绑定的证书ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// -1：域名未关联证书。
	// 1： 域名https已开启。
	// 0： 域名https已关闭。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type LiveInstanceList struct {
	// 该地域下live实例总数	
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// live实例详情	
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceList []*LiveInstanceDetail `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// 是否查询异常
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *string `json:"Error,omitnil,omitempty" name:"Error"`
}

type ManagerInfo struct {
	// 状态: audit: 审核中 ok: 审核通过 invalid: 失效 expiring: 即将过期 expired: 已过期
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 管理人姓名
	ManagerFirstName *string `json:"ManagerFirstName,omitnil,omitempty" name:"ManagerFirstName"`

	// 管理人姓名
	ManagerLastName *string `json:"ManagerLastName,omitnil,omitempty" name:"ManagerLastName"`

	// 管理人职位
	ManagerPosition *string `json:"ManagerPosition,omitnil,omitempty" name:"ManagerPosition"`

	// 管理人电话
	ManagerPhone *string `json:"ManagerPhone,omitnil,omitempty" name:"ManagerPhone"`

	// 管理人邮箱
	ManagerMail *string `json:"ManagerMail,omitnil,omitempty" name:"ManagerMail"`

	// 管理人所属部门
	ManagerDepartment *string `json:"ManagerDepartment,omitnil,omitempty" name:"ManagerDepartment"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 管理人域名数量
	DomainCount *int64 `json:"DomainCount,omitnil,omitempty" name:"DomainCount"`

	// 管理人证书数量
	CertCount *int64 `json:"CertCount,omitnil,omitempty" name:"CertCount"`

	// 管理人ID
	ManagerId *int64 `json:"ManagerId,omitnil,omitempty" name:"ManagerId"`

	// 审核有效到期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 最近一次提交审核时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubmitAuditTime *string `json:"SubmitAuditTime,omitnil,omitempty" name:"SubmitAuditTime"`

	// 审核通过时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerifyTime *string `json:"VerifyTime,omitnil,omitempty" name:"VerifyTime"`

	// 具体审核状态信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusInfo []*ManagerStatusInfo `json:"StatusInfo,omitnil,omitempty" name:"StatusInfo"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tags `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type ManagerStatusInfo struct {
	// 审核类型，枚举值：ov,ev,cs,ev_cs
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 审核状态，枚举值：pending,completed,invalid,submitted,expiring,expired
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`
}

// Predefined struct for user
type ModifyCertificateAliasRequestParams struct {
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 备注名称。
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`
}

type ModifyCertificateAliasRequest struct {
	*tchttp.BaseRequest
	
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 备注名称。
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`
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
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	CertificateIdList []*string `json:"CertificateIdList,omitnil,omitempty" name:"CertificateIdList"`

	// 项目 ID。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type ModifyCertificateProjectRequest struct {
	*tchttp.BaseRequest
	
	// 需要修改所属项目的证书 ID 集合，最多100个证书。
	CertificateIdList []*string `json:"CertificateIdList,omitnil,omitempty" name:"CertificateIdList"`

	// 项目 ID。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
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
	SuccessCertificates []*string `json:"SuccessCertificates,omitnil,omitempty" name:"SuccessCertificates"`

	// 修改所属项目失败的证书集合。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailCertificates []*string `json:"FailCertificates,omitnil,omitempty" name:"FailCertificates"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyCertificateResubmitRequestParams struct {
	// 证书ID。
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`
}

type ModifyCertificateResubmitRequest struct {
	*tchttp.BaseRequest
	
	// 证书ID。
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`
}

func (r *ModifyCertificateResubmitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCertificateResubmitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCertificateResubmitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCertificateResubmitResponseParams struct {
	// 证书ID。
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCertificateResubmitResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCertificateResubmitResponseParams `json:"Response"`
}

func (r *ModifyCertificateResubmitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCertificateResubmitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCertificatesExpiringNotificationSwitchRequestParams struct {
	// 证书ID列表。最多50个
	CertificateIds []*string `json:"CertificateIds,omitnil,omitempty" name:"CertificateIds"`

	// 0:不忽略通知。1:忽略通知
	SwitchStatus *uint64 `json:"SwitchStatus,omitnil,omitempty" name:"SwitchStatus"`
}

type ModifyCertificatesExpiringNotificationSwitchRequest struct {
	*tchttp.BaseRequest
	
	// 证书ID列表。最多50个
	CertificateIds []*string `json:"CertificateIds,omitnil,omitempty" name:"CertificateIds"`

	// 0:不忽略通知。1:忽略通知
	SwitchStatus *uint64 `json:"SwitchStatus,omitnil,omitempty" name:"SwitchStatus"`
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
	CertificateIds []*string `json:"CertificateIds,omitnil,omitempty" name:"CertificateIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 操作时间。
	CreatedOn *string `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// 主账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 子账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubAccountUin *string `json:"SubAccountUin,omitnil,omitempty" name:"SubAccountUin"`

	// 证书ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// 操作
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type PackageInfo struct {
	// 权益包ID。
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 权益包内权益点总量。
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 权益包内权益点余量。
	Balance *uint64 `json:"Balance,omitnil,omitempty" name:"Balance"`

	// 权益包名称。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 权益点是转入时，来源信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceUin *uint64 `json:"SourceUin,omitnil,omitempty" name:"SourceUin"`

	// 权益点状态。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 过期时间。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 更新时间。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 生成时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 来源类型。
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 转移信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferOutInfos []*PackageTransferOutInfo `json:"TransferOutInfos,omitnil,omitempty" name:"TransferOutInfos"`
}

type PackageTransferOutInfo struct {
	// 权益包ID。
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 转移码。
	TransferCode *string `json:"TransferCode,omitnil,omitempty" name:"TransferCode"`

	// 本次转移点数。
	TransferCount *uint64 `json:"TransferCount,omitnil,omitempty" name:"TransferCount"`

	// 转入的PackageID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReceivePackageId *string `json:"ReceivePackageId,omitnil,omitempty" name:"ReceivePackageId"`

	// 本次转移过期时间。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 本次转移生成时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 本次转移更新时间。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 转移状态。
	TransferStatus *string `json:"TransferStatus,omitnil,omitempty" name:"TransferStatus"`

	// 接收者uin。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReceiverUin *uint64 `json:"ReceiverUin,omitnil,omitempty" name:"ReceiverUin"`

	// 接收时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReceiveTime *string `json:"ReceiveTime,omitnil,omitempty" name:"ReceiveTime"`
}

type PreAuditInfo struct {
	// 证书总年限
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPeriod *int64 `json:"TotalPeriod,omitnil,omitempty" name:"TotalPeriod"`

	// 证书当前年限
	// 注意：此字段可能返回 null，表示取不到有效值。
	NowPeriod *int64 `json:"NowPeriod,omitnil,omitempty" name:"NowPeriod"`

	// 证书预审核管理人ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ManagerId *string `json:"ManagerId,omitnil,omitempty" name:"ManagerId"`
}

type ProjectInfo struct {
	// 项目名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 项目创建用户 UIN。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectCreatorUin *uint64 `json:"ProjectCreatorUin,omitnil,omitempty" name:"ProjectCreatorUin"`

	// 项目创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectCreateTime *string `json:"ProjectCreateTime,omitnil,omitempty" name:"ProjectCreateTime"`

	// 项目信息简述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectResume *string `json:"ProjectResume,omitnil,omitempty" name:"ProjectResume"`

	// 用户 UIN。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUin *uint64 `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 项目 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

// Predefined struct for user
type ReplaceCertificateRequestParams struct {
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 验证类型：DNS_AUTO = 自动DNS验证（仅支持在腾讯云解析且解析状态正常的域名使用该验证类型），DNS = 手动DNS验证，FILE = 文件验证。
	ValidType *string `json:"ValidType,omitnil,omitempty" name:"ValidType"`

	// 类型，默认 Original。可选项：Original = 原证书 CSR，Upload = 手动上传，Online = 在线生成。
	CsrType *string `json:"CsrType,omitnil,omitempty" name:"CsrType"`

	// CSR 内容。
	CsrContent *string `json:"CsrContent,omitnil,omitempty" name:"CsrContent"`

	// KEY 密码。
	CsrkeyPassword *string `json:"CsrkeyPassword,omitnil,omitempty" name:"CsrkeyPassword"`

	// 重颁发原因。
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// CSR加密方式，可选：RSA、ECC、SM2
	// （CsrType为Online才可选）， 默认为RSA
	CertCSREncryptAlgo *string `json:"CertCSREncryptAlgo,omitnil,omitempty" name:"CertCSREncryptAlgo"`

	// CSR加密参数，CsrEncryptAlgo为RSA时， 可选2048、4096等默认为2048；CsrEncryptAlgo为ECC时，可选prime256v1，secp384r1等，默认为prime256v1; 
	CertCSRKeyParameter *string `json:"CertCSRKeyParameter,omitnil,omitempty" name:"CertCSRKeyParameter"`
}

type ReplaceCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 验证类型：DNS_AUTO = 自动DNS验证（仅支持在腾讯云解析且解析状态正常的域名使用该验证类型），DNS = 手动DNS验证，FILE = 文件验证。
	ValidType *string `json:"ValidType,omitnil,omitempty" name:"ValidType"`

	// 类型，默认 Original。可选项：Original = 原证书 CSR，Upload = 手动上传，Online = 在线生成。
	CsrType *string `json:"CsrType,omitnil,omitempty" name:"CsrType"`

	// CSR 内容。
	CsrContent *string `json:"CsrContent,omitnil,omitempty" name:"CsrContent"`

	// KEY 密码。
	CsrkeyPassword *string `json:"CsrkeyPassword,omitnil,omitempty" name:"CsrkeyPassword"`

	// 重颁发原因。
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// CSR加密方式，可选：RSA、ECC、SM2
	// （CsrType为Online才可选）， 默认为RSA
	CertCSREncryptAlgo *string `json:"CertCSREncryptAlgo,omitnil,omitempty" name:"CertCSREncryptAlgo"`

	// CSR加密参数，CsrEncryptAlgo为RSA时， 可选2048、4096等默认为2048；CsrEncryptAlgo为ECC时，可选prime256v1，secp384r1等，默认为prime256v1; 
	CertCSRKeyParameter *string `json:"CertCSRKeyParameter,omitnil,omitempty" name:"CertCSRKeyParameter"`
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
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 地域列表
	Regions []*string `json:"Regions,omitnil,omitempty" name:"Regions"`
}

// Predefined struct for user
type RevokeCertificateRequestParams struct {
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 吊销证书原因。
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`
}

type RevokeCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 吊销证书原因。
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`
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
	RevokeDomainValidateAuths []*RevokeDomainValidateAuths `json:"RevokeDomainValidateAuths,omitnil,omitempty" name:"RevokeDomainValidateAuths"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	DomainValidateAuthPath *string `json:"DomainValidateAuthPath,omitnil,omitempty" name:"DomainValidateAuthPath"`

	// DV 认证 KEY。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainValidateAuthKey *string `json:"DomainValidateAuthKey,omitnil,omitempty" name:"DomainValidateAuthKey"`

	// DV 认证值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainValidateAuthValue *string `json:"DomainValidateAuthValue,omitnil,omitempty" name:"DomainValidateAuthValue"`

	// DV 认证域名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainValidateAuthDomain *string `json:"DomainValidateAuthDomain,omitnil,omitempty" name:"DomainValidateAuthDomain"`
}

type RootCertificates struct {
	// 国密签名证书
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sign *string `json:"Sign,omitnil,omitempty" name:"Sign"`

	// 国密加密证书
	// 注意：此字段可能返回 null，表示取不到有效值。
	Encrypt *string `json:"Encrypt,omitnil,omitempty" name:"Encrypt"`

	// 标准证书
	// 注意：此字段可能返回 null，表示取不到有效值。
	Standard *string `json:"Standard,omitnil,omitempty" name:"Standard"`
}

// Predefined struct for user
type SubmitAuditManagerRequestParams struct {
	// 管理人ID
	ManagerId *int64 `json:"ManagerId,omitnil,omitempty" name:"ManagerId"`
}

type SubmitAuditManagerRequest struct {
	*tchttp.BaseRequest
	
	// 管理人ID
	ManagerId *int64 `json:"ManagerId,omitnil,omitempty" name:"ManagerId"`
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
	ManagerId *int64 `json:"ManagerId,omitnil,omitempty" name:"ManagerId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// CSR 生成方式：online = 在线生成, parse = 手动上传。
	CsrType *string `json:"CsrType,omitnil,omitempty" name:"CsrType"`

	// 上传的 CSR 内容。
	CsrContent *string `json:"CsrContent,omitnil,omitempty" name:"CsrContent"`

	// 绑定证书的域名。
	CertificateDomain *string `json:"CertificateDomain,omitnil,omitempty" name:"CertificateDomain"`

	// 上传的域名数组（多域名证书可以上传）。
	DomainList []*string `json:"DomainList,omitnil,omitempty" name:"DomainList"`

	// 私钥密码（非必填）。
	KeyPassword *string `json:"KeyPassword,omitnil,omitempty" name:"KeyPassword"`

	// 公司名称。
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	// 部门名称。
	OrganizationDivision *string `json:"OrganizationDivision,omitnil,omitempty" name:"OrganizationDivision"`

	// 公司详细地址。
	OrganizationAddress *string `json:"OrganizationAddress,omitnil,omitempty" name:"OrganizationAddress"`

	// 国家名称，如中国：CN 。
	OrganizationCountry *string `json:"OrganizationCountry,omitnil,omitempty" name:"OrganizationCountry"`

	// 公司所在城市。
	OrganizationCity *string `json:"OrganizationCity,omitnil,omitempty" name:"OrganizationCity"`

	// 公司所在省份。
	OrganizationRegion *string `json:"OrganizationRegion,omitnil,omitempty" name:"OrganizationRegion"`

	// 公司邮编。
	PostalCode *string `json:"PostalCode,omitnil,omitempty" name:"PostalCode"`

	// 公司座机区号。
	PhoneAreaCode *string `json:"PhoneAreaCode,omitnil,omitempty" name:"PhoneAreaCode"`

	// 公司座机号码。
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// 证书验证方式。验证类型：DNS_AUTO = 自动DNS验证（仅支持在腾讯云解析且解析状态正常的域名使用该验证类型），DNS = 手动DNS验证，FILE = 文件验证。
	VerifyType *string `json:"VerifyType,omitnil,omitempty" name:"VerifyType"`

	// 管理人名。
	AdminFirstName *string `json:"AdminFirstName,omitnil,omitempty" name:"AdminFirstName"`

	// 管理人姓。
	AdminLastName *string `json:"AdminLastName,omitnil,omitempty" name:"AdminLastName"`

	// 管理人手机号码。
	AdminPhoneNum *string `json:"AdminPhoneNum,omitnil,omitempty" name:"AdminPhoneNum"`

	// 管理人邮箱地址。
	AdminEmail *string `json:"AdminEmail,omitnil,omitempty" name:"AdminEmail"`

	// 管理人职位。
	AdminPosition *string `json:"AdminPosition,omitnil,omitempty" name:"AdminPosition"`

	// 联系人名。
	ContactFirstName *string `json:"ContactFirstName,omitnil,omitempty" name:"ContactFirstName"`

	// 联系人姓。
	ContactLastName *string `json:"ContactLastName,omitnil,omitempty" name:"ContactLastName"`

	// 联系人邮箱地址。
	ContactEmail *string `json:"ContactEmail,omitnil,omitempty" name:"ContactEmail"`

	// 联系人手机号码。
	ContactNumber *string `json:"ContactNumber,omitnil,omitempty" name:"ContactNumber"`

	// 联系人职位。
	ContactPosition *string `json:"ContactPosition,omitnil,omitempty" name:"ContactPosition"`
}

type SubmitCertificateInformationRequest struct {
	*tchttp.BaseRequest
	
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// CSR 生成方式：online = 在线生成, parse = 手动上传。
	CsrType *string `json:"CsrType,omitnil,omitempty" name:"CsrType"`

	// 上传的 CSR 内容。
	CsrContent *string `json:"CsrContent,omitnil,omitempty" name:"CsrContent"`

	// 绑定证书的域名。
	CertificateDomain *string `json:"CertificateDomain,omitnil,omitempty" name:"CertificateDomain"`

	// 上传的域名数组（多域名证书可以上传）。
	DomainList []*string `json:"DomainList,omitnil,omitempty" name:"DomainList"`

	// 私钥密码（非必填）。
	KeyPassword *string `json:"KeyPassword,omitnil,omitempty" name:"KeyPassword"`

	// 公司名称。
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	// 部门名称。
	OrganizationDivision *string `json:"OrganizationDivision,omitnil,omitempty" name:"OrganizationDivision"`

	// 公司详细地址。
	OrganizationAddress *string `json:"OrganizationAddress,omitnil,omitempty" name:"OrganizationAddress"`

	// 国家名称，如中国：CN 。
	OrganizationCountry *string `json:"OrganizationCountry,omitnil,omitempty" name:"OrganizationCountry"`

	// 公司所在城市。
	OrganizationCity *string `json:"OrganizationCity,omitnil,omitempty" name:"OrganizationCity"`

	// 公司所在省份。
	OrganizationRegion *string `json:"OrganizationRegion,omitnil,omitempty" name:"OrganizationRegion"`

	// 公司邮编。
	PostalCode *string `json:"PostalCode,omitnil,omitempty" name:"PostalCode"`

	// 公司座机区号。
	PhoneAreaCode *string `json:"PhoneAreaCode,omitnil,omitempty" name:"PhoneAreaCode"`

	// 公司座机号码。
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// 证书验证方式。验证类型：DNS_AUTO = 自动DNS验证（仅支持在腾讯云解析且解析状态正常的域名使用该验证类型），DNS = 手动DNS验证，FILE = 文件验证。
	VerifyType *string `json:"VerifyType,omitnil,omitempty" name:"VerifyType"`

	// 管理人名。
	AdminFirstName *string `json:"AdminFirstName,omitnil,omitempty" name:"AdminFirstName"`

	// 管理人姓。
	AdminLastName *string `json:"AdminLastName,omitnil,omitempty" name:"AdminLastName"`

	// 管理人手机号码。
	AdminPhoneNum *string `json:"AdminPhoneNum,omitnil,omitempty" name:"AdminPhoneNum"`

	// 管理人邮箱地址。
	AdminEmail *string `json:"AdminEmail,omitnil,omitempty" name:"AdminEmail"`

	// 管理人职位。
	AdminPosition *string `json:"AdminPosition,omitnil,omitempty" name:"AdminPosition"`

	// 联系人名。
	ContactFirstName *string `json:"ContactFirstName,omitnil,omitempty" name:"ContactFirstName"`

	// 联系人姓。
	ContactLastName *string `json:"ContactLastName,omitnil,omitempty" name:"ContactLastName"`

	// 联系人邮箱地址。
	ContactEmail *string `json:"ContactEmail,omitnil,omitempty" name:"ContactEmail"`

	// 联系人手机号码。
	ContactNumber *string `json:"ContactNumber,omitnil,omitempty" name:"ContactNumber"`

	// 联系人职位。
	ContactPosition *string `json:"ContactPosition,omitnil,omitempty" name:"ContactPosition"`
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
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	CsrType *string `json:"CsrType,omitnil,omitempty" name:"CsrType"`

	// CSR 内容。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CsrContent *string `json:"CsrContent,omitnil,omitempty" name:"CsrContent"`

	// 域名信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertificateDomain *string `json:"CertificateDomain,omitnil,omitempty" name:"CertificateDomain"`

	// DNS 信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainList []*string `json:"DomainList,omitnil,omitempty" name:"DomainList"`

	// 私钥密码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyPassword *string `json:"KeyPassword,omitnil,omitempty" name:"KeyPassword"`

	// 企业或单位名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	// 部门。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationDivision *string `json:"OrganizationDivision,omitnil,omitempty" name:"OrganizationDivision"`

	// 地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationAddress *string `json:"OrganizationAddress,omitnil,omitempty" name:"OrganizationAddress"`

	// 国家。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationCountry *string `json:"OrganizationCountry,omitnil,omitempty" name:"OrganizationCountry"`

	// 市。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationCity *string `json:"OrganizationCity,omitnil,omitempty" name:"OrganizationCity"`

	// 省。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrganizationRegion *string `json:"OrganizationRegion,omitnil,omitempty" name:"OrganizationRegion"`

	// 邮政编码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PostalCode *string `json:"PostalCode,omitnil,omitempty" name:"PostalCode"`

	// 座机区号。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneAreaCode *string `json:"PhoneAreaCode,omitnil,omitempty" name:"PhoneAreaCode"`

	// 座机号码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// 管理员名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdminFirstName *string `json:"AdminFirstName,omitnil,omitempty" name:"AdminFirstName"`

	// 管理员姓。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdminLastName *string `json:"AdminLastName,omitnil,omitempty" name:"AdminLastName"`

	// 管理员电话号码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdminPhoneNum *string `json:"AdminPhoneNum,omitnil,omitempty" name:"AdminPhoneNum"`

	// 管理员邮箱地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdminEmail *string `json:"AdminEmail,omitnil,omitempty" name:"AdminEmail"`

	// 管理员职位。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdminPosition *string `json:"AdminPosition,omitnil,omitempty" name:"AdminPosition"`

	// 联系人名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContactFirstName *string `json:"ContactFirstName,omitnil,omitempty" name:"ContactFirstName"`

	// 联系人姓。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContactLastName *string `json:"ContactLastName,omitnil,omitempty" name:"ContactLastName"`

	// 联系人电话号码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContactNumber *string `json:"ContactNumber,omitnil,omitempty" name:"ContactNumber"`

	// 联系人邮箱地址，
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContactEmail *string `json:"ContactEmail,omitnil,omitempty" name:"ContactEmail"`

	// 联系人职位。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContactPosition *string `json:"ContactPosition,omitnil,omitempty" name:"ContactPosition"`

	// 验证类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerifyType *string `json:"VerifyType,omitnil,omitempty" name:"VerifyType"`
}

type SupportDownloadType struct {
	// 是否可以下载nginx可用格式
	NGINX *bool `json:"NGINX,omitnil,omitempty" name:"NGINX"`

	// 是否可以下载apache可用格式
	APACHE *bool `json:"APACHE,omitnil,omitempty" name:"APACHE"`

	// 是否可以下载tomcat可用格式
	TOMCAT *bool `json:"TOMCAT,omitnil,omitempty" name:"TOMCAT"`

	// 是否可以下载iis可用格式
	IIS *bool `json:"IIS,omitnil,omitempty" name:"IIS"`

	// 是否可以下载JKS可用格式
	JKS *bool `json:"JKS,omitnil,omitempty" name:"JKS"`

	// 是否可以下载其他格式
	OTHER *bool `json:"OTHER,omitnil,omitempty" name:"OTHER"`

	// 是否可以下载根证书
	ROOT *bool `json:"ROOT,omitnil,omitempty" name:"ROOT"`
}

type SyncTaskBindResourceResult struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 关联云资源结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindResourceResult []*BindResourceResult `json:"BindResourceResult,omitnil,omitempty" name:"BindResourceResult"`

	// 关联云资源异步查询结果： 0表示查询中， 1表示查询成功。 2表示查询异常； 若状态为1，则查看BindResourceResult结果；若状态为2，则查看Error原因
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 关联云资源错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *Error `json:"Error,omitnil,omitempty" name:"Error"`

	// 当前结果缓存时间
	CacheTime *string `json:"CacheTime,omitnil,omitempty" name:"CacheTime"`
}

type TCBAccessInstance struct {
	// 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 统一域名状态
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnionStatus *int64 `json:"UnionStatus,omitnil,omitempty" name:"UnionStatus"`

	// 是否被抢占, 被抢占表示域名被其他环境绑定了，需要解绑或者重新绑定。
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsPreempted *bool `json:"IsPreempted,omitnil,omitempty" name:"IsPreempted"`

	// icp黑名单封禁状态，0-未封禁，1-封禁
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	ICPStatus *int64 `json:"ICPStatus,omitnil,omitempty" name:"ICPStatus"`

	// 已绑定证书ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`
}

type TCBAccessService struct {
	// 实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceList []*TCBAccessInstance `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// 数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type TCBEnvironment struct {
	// 唯一ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type TCBEnvironments struct {
	// tcb环境	
	// 注意：此字段可能返回 null，表示取不到有效值。
	Environment *TCBEnvironment `json:"Environment,omitnil,omitempty" name:"Environment"`

	// 访问服务	
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessService *TCBAccessService `json:"AccessService,omitnil,omitempty" name:"AccessService"`

	// 静态托管	
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostService *TCBHostService `json:"HostService,omitnil,omitempty" name:"HostService"`
}

type TCBHostInstance struct {
	// 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 解析状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	DNSStatus *string `json:"DNSStatus,omitnil,omitempty" name:"DNSStatus"`

	// 已绑定证书ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`
}

type TCBHostService struct {
	// 实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceList []*TCBHostInstance `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// 数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type TCBInstanceList struct {
	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// tcb环境实例详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Environments []*TCBEnvironments `json:"Environments,omitnil,omitempty" name:"Environments"`

	// 是否查询异常
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *string `json:"Error,omitnil,omitempty" name:"Error"`
}

type TSEInstanceDetail struct {
	// 网关ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 网关名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GatewayName *string `json:"GatewayName,omitnil,omitempty" name:"GatewayName"`

	// 网关证书列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertificateList []*GatewayCertificate `json:"CertificateList,omitnil,omitempty" name:"CertificateList"`
}

type TSEInstanceList struct {
	// TSE实例详情	
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceList []*TSEInstanceDetail `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// 该地域下TSE实例总数	
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 地域	
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 是否查询异常
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *string `json:"Error,omitnil,omitempty" name:"Error"`
}

type Tags struct {
	// 标签键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TeoInstanceDetail struct {
	// 域名
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 证书ID
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// 区域ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 域名状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type TeoInstanceList struct {
	// edgeone实例详情	
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceList []*TeoInstanceDetail `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// edgeone实例总数	
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 是否查询异常
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *string `json:"Error,omitnil,omitempty" name:"Error"`
}

type TkeIngressDetail struct {
	// ingress名称
	IngressName *string `json:"IngressName,omitnil,omitempty" name:"IngressName"`

	// tls域名列表
	TlsDomains []*string `json:"TlsDomains,omitnil,omitempty" name:"TlsDomains"`

	// ingress域名列表
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`
}

type TkeInstanceDetail struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 集群命名空间列表
	NamespaceList []*TkeNameSpaceDetail `json:"NamespaceList,omitnil,omitempty" name:"NamespaceList"`

	// 集群类型
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 集群版本
	ClusterVersion *string `json:"ClusterVersion,omitnil,omitempty" name:"ClusterVersion"`
}

type TkeInstanceList struct {
	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// tke实例详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceList []*TkeInstanceDetail `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// 该地域下tke实例总数	
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 是否查询异常
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *string `json:"Error,omitnil,omitempty" name:"Error"`
}

type TkeNameSpaceDetail struct {
	// namespace名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// secret列表
	SecretList []*TkeSecretDetail `json:"SecretList,omitnil,omitempty" name:"SecretList"`
}

type TkeSecretDetail struct {
	// secret名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 证书ID
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// ingress列表
	IngressList []*TkeIngressDetail `json:"IngressList,omitnil,omitempty" name:"IngressList"`

	// 和新证书不匹配的域名列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NoMatchDomains []*string `json:"NoMatchDomains,omitnil,omitempty" name:"NoMatchDomains"`
}

// Predefined struct for user
type UpdateCertificateInstanceRequestParams struct {
	// 一键更新原证书ID， 查询绑定该证书的云资源然后进行证书更新
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`

	// 需要部署的资源类型，参数值可选（小写）：clb、cdn、waf、live、ddos、teo、apigateway、vod、tke、tcb、tse、cos
	ResourceTypes []*string `json:"ResourceTypes,omitnil,omitempty" name:"ResourceTypes"`

	// 一键更新新证书ID，不传该则证书公钥和私钥必传
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 需要部署的地域列表（废弃）
	//
	// Deprecated: Regions is deprecated.
	Regions []*string `json:"Regions,omitnil,omitempty" name:"Regions"`

	// 云资源需要部署的地域列表，支持地域的云资源类型必传，如：clb、tke、apigateway、waf、tcb、tse等
	ResourceTypesRegions []*ResourceTypeRegions `json:"ResourceTypesRegions,omitnil,omitempty" name:"ResourceTypesRegions"`

	// 证书公钥， 若上传证书公钥， 则CertificateId不用传
	CertificatePublicKey *string `json:"CertificatePublicKey,omitnil,omitempty" name:"CertificatePublicKey"`

	// 证书私钥，若上传证书公钥， 则CertificateId不用传
	CertificatePrivateKey *string `json:"CertificatePrivateKey,omitnil,omitempty" name:"CertificatePrivateKey"`

	// 旧证书是否忽略到期提醒  0:不忽略通知。1:忽略通知，忽略OldCertificateId到期提醒
	ExpiringNotificationSwitch *uint64 `json:"ExpiringNotificationSwitch,omitnil,omitempty" name:"ExpiringNotificationSwitch"`

	// 相同的证书是否允许重复上传，若选择上传证书， 则可以配置该参数
	Repeatable *bool `json:"Repeatable,omitnil,omitempty" name:"Repeatable"`

	// 是否允许下载，若选择上传证书， 则可以配置该参数
	AllowDownload *bool `json:"AllowDownload,omitnil,omitempty" name:"AllowDownload"`

	// 标签列表，若选择上传证书， 则可以配置该参数
	Tags []*Tags `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 项目 ID，若选择上传证书， 则可以配置该参数
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type UpdateCertificateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 一键更新原证书ID， 查询绑定该证书的云资源然后进行证书更新
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`

	// 需要部署的资源类型，参数值可选（小写）：clb、cdn、waf、live、ddos、teo、apigateway、vod、tke、tcb、tse、cos
	ResourceTypes []*string `json:"ResourceTypes,omitnil,omitempty" name:"ResourceTypes"`

	// 一键更新新证书ID，不传该则证书公钥和私钥必传
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 需要部署的地域列表（废弃）
	Regions []*string `json:"Regions,omitnil,omitempty" name:"Regions"`

	// 云资源需要部署的地域列表，支持地域的云资源类型必传，如：clb、tke、apigateway、waf、tcb、tse等
	ResourceTypesRegions []*ResourceTypeRegions `json:"ResourceTypesRegions,omitnil,omitempty" name:"ResourceTypesRegions"`

	// 证书公钥， 若上传证书公钥， 则CertificateId不用传
	CertificatePublicKey *string `json:"CertificatePublicKey,omitnil,omitempty" name:"CertificatePublicKey"`

	// 证书私钥，若上传证书公钥， 则CertificateId不用传
	CertificatePrivateKey *string `json:"CertificatePrivateKey,omitnil,omitempty" name:"CertificatePrivateKey"`

	// 旧证书是否忽略到期提醒  0:不忽略通知。1:忽略通知，忽略OldCertificateId到期提醒
	ExpiringNotificationSwitch *uint64 `json:"ExpiringNotificationSwitch,omitnil,omitempty" name:"ExpiringNotificationSwitch"`

	// 相同的证书是否允许重复上传，若选择上传证书， 则可以配置该参数
	Repeatable *bool `json:"Repeatable,omitnil,omitempty" name:"Repeatable"`

	// 是否允许下载，若选择上传证书， 则可以配置该参数
	AllowDownload *bool `json:"AllowDownload,omitnil,omitempty" name:"AllowDownload"`

	// 标签列表，若选择上传证书， 则可以配置该参数
	Tags []*Tags `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 项目 ID，若选择上传证书， 则可以配置该参数
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
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
	delete(f, "OldCertificateId")
	delete(f, "ResourceTypes")
	delete(f, "CertificateId")
	delete(f, "Regions")
	delete(f, "ResourceTypesRegions")
	delete(f, "CertificatePublicKey")
	delete(f, "CertificatePrivateKey")
	delete(f, "ExpiringNotificationSwitch")
	delete(f, "Repeatable")
	delete(f, "AllowDownload")
	delete(f, "Tags")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCertificateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCertificateInstanceResponseParams struct {
	// 云资源部署任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployRecordId *uint64 `json:"DeployRecordId,omitnil,omitempty" name:"DeployRecordId"`

	// 部署状态，1表示部署成功，0表示部署失败
	DeployStatus *int64 `json:"DeployStatus,omitnil,omitempty" name:"DeployStatus"`

	// 更新异步创建任务进度详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateSyncProgress []*UpdateSyncProgress `json:"UpdateSyncProgress,omitnil,omitempty" name:"UpdateSyncProgress"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	DeployRecordId *int64 `json:"DeployRecordId,omitnil,omitempty" name:"DeployRecordId"`

	// 待重试部署记录详情ID
	DeployRecordDetailId *int64 `json:"DeployRecordDetailId,omitnil,omitempty" name:"DeployRecordDetailId"`
}

type UpdateCertificateRecordRetryRequest struct {
	*tchttp.BaseRequest
	
	// 待重试部署记录ID
	DeployRecordId *int64 `json:"DeployRecordId,omitnil,omitempty" name:"DeployRecordId"`

	// 待重试部署记录详情ID
	DeployRecordDetailId *int64 `json:"DeployRecordDetailId,omitnil,omitempty" name:"DeployRecordDetailId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	DeployRecordId *int64 `json:"DeployRecordId,omitnil,omitempty" name:"DeployRecordId"`
}

type UpdateCertificateRecordRollbackRequest struct {
	*tchttp.BaseRequest
	
	// 待重试部署记录ID
	DeployRecordId *int64 `json:"DeployRecordId,omitnil,omitempty" name:"DeployRecordId"`
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
	DeployRecordId *int64 `json:"DeployRecordId,omitnil,omitempty" name:"DeployRecordId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 新证书ID
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// 旧证书ID
	OldCertId *string `json:"OldCertId,omitnil,omitempty" name:"OldCertId"`

	// 部署域名列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// 部署资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 部署地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 部署状态
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 部署错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 部署时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最后一次更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 部署实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 部署实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 部署监听器ID（CLB专用）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 部署监听器名称（CLB专用）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListenerName *string `json:"ListenerName,omitnil,omitempty" name:"ListenerName"`

	// 协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 是否开启SNI（CLB专用）
	// 注意：此字段可能返回 null，表示取不到有效值。
	SniSwitch *uint64 `json:"SniSwitch,omitnil,omitempty" name:"SniSwitch"`

	// bucket名称（COS专用）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// 端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 命名空间（TKE专用）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// secret名称（TKE专用）
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretName *string `json:"SecretName,omitnil,omitempty" name:"SecretName"`

	// 环境ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// TCB部署类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TCBType *string `json:"TCBType,omitnil,omitempty" name:"TCBType"`

	// 监听器Url(clb专属)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type UpdateRecordDetails struct {
	// 部署资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 部署资源详情列表
	List []*UpdateRecordDetail `json:"List,omitnil,omitempty" name:"List"`

	// 该部署资源总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type UpdateRecordInfo struct {
	// 记录ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 新证书ID
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// 原证书ID
	OldCertId *string `json:"OldCertId,omitnil,omitempty" name:"OldCertId"`

	// 部署资源类型列表
	ResourceTypes []*string `json:"ResourceTypes,omitnil,omitempty" name:"ResourceTypes"`

	// 部署地域列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Regions []*string `json:"Regions,omitnil,omitempty" name:"Regions"`

	// 部署状态
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 部署时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最后一次更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type UpdateSyncProgress struct {
	// 资源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 地域结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateSyncProgressRegions []*UpdateSyncProgressRegion `json:"UpdateSyncProgressRegions,omitnil,omitempty" name:"UpdateSyncProgressRegions"`

	// 异步更新进度状态：0， 待处理， 1 已处理， 3 处理中
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type UpdateSyncProgressRegion struct {
	// 资源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 执行完成数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	OffsetCount *int64 `json:"OffsetCount,omitnil,omitempty" name:"OffsetCount"`

	// 异步更新进度状态：0， 待处理， 1 已处理， 3 处理中
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type UploadCertificateRequestParams struct {
	// 证书内容。
	CertificatePublicKey *string `json:"CertificatePublicKey,omitnil,omitempty" name:"CertificatePublicKey"`

	// 私钥内容，证书类型为 SVR 时必填，为 CA 时可不填。
	CertificatePrivateKey *string `json:"CertificatePrivateKey,omitnil,omitempty" name:"CertificatePrivateKey"`

	// 证书类型，默认 SVR。CA = CA证书，SVR = 服务器证书。
	CertificateType *string `json:"CertificateType,omitnil,omitempty" name:"CertificateType"`

	// 备注名称。
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 项目 ID。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 证书用途/证书来源。“CLB，CDN，WAF，LIVE，DDOS”
	CertificateUse *string `json:"CertificateUse,omitnil,omitempty" name:"CertificateUse"`

	// 标签列表
	Tags []*Tags `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 相同的证书是否允许重复上传
	Repeatable *bool `json:"Repeatable,omitnil,omitempty" name:"Repeatable"`
}

type UploadCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 证书内容。
	CertificatePublicKey *string `json:"CertificatePublicKey,omitnil,omitempty" name:"CertificatePublicKey"`

	// 私钥内容，证书类型为 SVR 时必填，为 CA 时可不填。
	CertificatePrivateKey *string `json:"CertificatePrivateKey,omitnil,omitempty" name:"CertificatePrivateKey"`

	// 证书类型，默认 SVR。CA = CA证书，SVR = 服务器证书。
	CertificateType *string `json:"CertificateType,omitnil,omitempty" name:"CertificateType"`

	// 备注名称。
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 项目 ID。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 证书用途/证书来源。“CLB，CDN，WAF，LIVE，DDOS”
	CertificateUse *string `json:"CertificateUse,omitnil,omitempty" name:"CertificateUse"`

	// 标签列表
	Tags []*Tags `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 相同的证书是否允许重复上传
	Repeatable *bool `json:"Repeatable,omitnil,omitempty" name:"Repeatable"`
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
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 重复证书的ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepeatCertId *string `json:"RepeatCertId,omitnil,omitempty" name:"RepeatCertId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// base64编码后的证书确认函文件，格式应为jpg、jpeg、png、pdf，大小应在1kb与1.4M之间。
	ConfirmLetter *string `json:"ConfirmLetter,omitnil,omitempty" name:"ConfirmLetter"`
}

type UploadConfirmLetterRequest struct {
	*tchttp.BaseRequest
	
	// 证书ID
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// base64编码后的证书确认函文件，格式应为jpg、jpeg、png、pdf，大小应在1kb与1.4M之间。
	ConfirmLetter *string `json:"ConfirmLetter,omitnil,omitempty" name:"ConfirmLetter"`
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
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 是否成功
	IsSuccess *bool `json:"IsSuccess,omitnil,omitempty" name:"IsSuccess"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// base64编码后的证书确认函文件，格式应为jpg、jpeg、png、pdf，大小应在1kb与1.4M之间。
	RevokeLetter *string `json:"RevokeLetter,omitnil,omitempty" name:"RevokeLetter"`
}

type UploadRevokeLetterRequest struct {
	*tchttp.BaseRequest
	
	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// base64编码后的证书确认函文件，格式应为jpg、jpeg、png、pdf，大小应在1kb与1.4M之间。
	RevokeLetter *string `json:"RevokeLetter,omitnil,omitempty" name:"RevokeLetter"`
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
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 是否成功。
	IsSuccess *bool `json:"IsSuccess,omitnil,omitempty" name:"IsSuccess"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type VODInstanceList struct {
	// vod实例详情	
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceList []*VodInstanceDetail `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// 该地域下vod实例总数	
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 是否查询异常
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *string `json:"Error,omitnil,omitempty" name:"Error"`
}

// Predefined struct for user
type VerifyManagerRequestParams struct {
	// 管理人ID
	ManagerId *int64 `json:"ManagerId,omitnil,omitempty" name:"ManagerId"`
}

type VerifyManagerRequest struct {
	*tchttp.BaseRequest
	
	// 管理人ID
	ManagerId *int64 `json:"ManagerId,omitnil,omitempty" name:"ManagerId"`
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
	ManagerId *int64 `json:"ManagerId,omitnil,omitempty" name:"ManagerId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 证书ID
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`
}

type WafInstanceDetail struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 证书ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// 是否保持长连接，1是，0 否
	// 注意：此字段可能返回 null，表示取不到有效值。
	Keepalive *uint64 `json:"Keepalive,omitnil,omitempty" name:"Keepalive"`
}

type WafInstanceList struct {
	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// waf实例详情	
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceList []*WafInstanceDetail `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// 该地域下waf实例总数	
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 是否查询异常
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *string `json:"Error,omitnil,omitempty" name:"Error"`
}