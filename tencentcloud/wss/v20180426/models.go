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

package v20180426

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type DeleteCertRequestParams struct {
	// 证书 ID，即通过 GetList 拿到的证书列表的 ID 字段。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 模块名称，应填 ssl。
	ModuleType *string `json:"ModuleType,omitempty" name:"ModuleType"`
}

type DeleteCertRequest struct {
	*tchttp.BaseRequest
	
	// 证书 ID，即通过 GetList 拿到的证书列表的 ID 字段。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 模块名称，应填 ssl。
	ModuleType *string `json:"ModuleType,omitempty" name:"ModuleType"`
}

func (r *DeleteCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCertRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "ModuleType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCertRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCertResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteCertResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCertResponseParams `json:"Response"`
}

func (r *DeleteCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCertResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertListRequestParams struct {
	// 模块名称，应填 ssl。
	ModuleType *string `json:"ModuleType,omitempty" name:"ModuleType"`

	// 页数，默认第一页。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数，默认每页20条。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索关键字。
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// 证书类型（目前支持:CA=客户端证书,SVR=服务器证书）。
	CertType *string `json:"CertType,omitempty" name:"CertType"`

	// 证书ID。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 是否同时获取证书内容。
	WithCert *string `json:"WithCert,omitempty" name:"WithCert"`

	// 如传，则只返回可以给该域名使用的证书。
	AltDomain *string `json:"AltDomain,omitempty" name:"AltDomain"`
}

type DescribeCertListRequest struct {
	*tchttp.BaseRequest
	
	// 模块名称，应填 ssl。
	ModuleType *string `json:"ModuleType,omitempty" name:"ModuleType"`

	// 页数，默认第一页。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数，默认每页20条。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索关键字。
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// 证书类型（目前支持:CA=客户端证书,SVR=服务器证书）。
	CertType *string `json:"CertType,omitempty" name:"CertType"`

	// 证书ID。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 是否同时获取证书内容。
	WithCert *string `json:"WithCert,omitempty" name:"WithCert"`

	// 如传，则只返回可以给该域名使用的证书。
	AltDomain *string `json:"AltDomain,omitempty" name:"AltDomain"`
}

func (r *DescribeCertListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModuleType")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchKey")
	delete(f, "CertType")
	delete(f, "Id")
	delete(f, "WithCert")
	delete(f, "AltDomain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCertListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertListResponseParams struct {
	// 总数量。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 列表。
	CertificateSet []*SSLCertificate `json:"CertificateSet,omitempty" name:"CertificateSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCertListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCertListResponseParams `json:"Response"`
}

func (r *DescribeCertListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SSLCertificate struct {
	// 所属账户
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 项目ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 证书来源：trustasia = 亚洲诚信， upload = 用户上传
	// 注意：此字段可能返回 null，表示取不到有效值。
	From *string `json:"From,omitempty" name:"From"`

	// 证书类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 证书类型（目前支持：CA = 客户端证书，SVR = 服务器证书）
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertType *string `json:"CertType,omitempty" name:"CertType"`

	// 证书办法者名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductZhName *string `json:"ProductZhName,omitempty" name:"ProductZhName"`

	// 主域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 状态值 0：审核中，1：已通过，2：审核失败，3：已过期，4：已添加云解析记录，5：OV/EV 证书，待提交资料，6：订单取消中，7：已取消，8：已提交资料， 待上传确认函
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 漏洞扫描状态：INACTIVE = 未开启，ACTIVE = 已开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulnerabilityStatus *string `json:"VulnerabilityStatus,omitempty" name:"VulnerabilityStatus"`

	// 状态信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusMsg *string `json:"StatusMsg,omitempty" name:"StatusMsg"`

	// 验证类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerifyType *string `json:"VerifyType,omitempty" name:"VerifyType"`

	// 证书生效时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertBeginTime *string `json:"CertBeginTime,omitempty" name:"CertBeginTime"`

	// 证书过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertEndTime *string `json:"CertEndTime,omitempty" name:"CertEndTime"`

	// 证书过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValidityPeriod *string `json:"ValidityPeriod,omitempty" name:"ValidityPeriod"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	InsertTime *string `json:"InsertTime,omitempty" name:"InsertTime"`

	// 项目信息，ProjectId：项目ID，OwnerUin：项目所属的 uin（默认项目为0），Name：项目名称，CreatorUin：创建项目的 uin，CreateTime：项目创建时间，Info：项目说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectInfo *SSLProjectInfo `json:"ProjectInfo,omitempty" name:"ProjectInfo"`

	// 证书ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 证书包含的多个域名（包含主域名）
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubjectAltName []*string `json:"SubjectAltName,omitempty" name:"SubjectAltName"`

	// 证书类型名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TypeName *string `json:"TypeName,omitempty" name:"TypeName"`

	// 状态名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusName *string `json:"StatusName,omitempty" name:"StatusName"`

	// 是否为 VIP 客户
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsVip *bool `json:"IsVip,omitempty" name:"IsVip"`

	// 是否我 DV 版证书
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDv *bool `json:"IsDv,omitempty" name:"IsDv"`

	// 是否为泛域名证书
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsWildcard *bool `json:"IsWildcard,omitempty" name:"IsWildcard"`

	// 是否启用了漏洞扫描功能
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsVulnerability *bool `json:"IsVulnerability,omitempty" name:"IsVulnerability"`

	// 证书
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cert *string `json:"Cert,omitempty" name:"Cert"`
}

type SSLProjectInfo struct {
	// 项目ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 项目所属的 uin（默认项目为0）
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUin *int64 `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 创建项目的 uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatorUin *int64 `json:"CreatorUin,omitempty" name:"CreatorUin"`

	// 项目创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 项目说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	Info *string `json:"Info,omitempty" name:"Info"`
}

// Predefined struct for user
type UploadCertRequestParams struct {
	// 证书内容。
	Cert *string `json:"Cert,omitempty" name:"Cert"`

	// 证书类型（目前支持：CA 为客户端证书，SVR 为服务器证书）。
	CertType *string `json:"CertType,omitempty" name:"CertType"`

	// 项目ID，详见用户指南的 [项目与标签](https://cloud.tencent.com/document/product/598/32738)。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 模块名称，应填 ssl。
	ModuleType *string `json:"ModuleType,omitempty" name:"ModuleType"`

	// 证书私钥，certType=SVR 时必填。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 证书备注。
	Alias *string `json:"Alias,omitempty" name:"Alias"`
}

type UploadCertRequest struct {
	*tchttp.BaseRequest
	
	// 证书内容。
	Cert *string `json:"Cert,omitempty" name:"Cert"`

	// 证书类型（目前支持：CA 为客户端证书，SVR 为服务器证书）。
	CertType *string `json:"CertType,omitempty" name:"CertType"`

	// 项目ID，详见用户指南的 [项目与标签](https://cloud.tencent.com/document/product/598/32738)。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 模块名称，应填 ssl。
	ModuleType *string `json:"ModuleType,omitempty" name:"ModuleType"`

	// 证书私钥，certType=SVR 时必填。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 证书备注。
	Alias *string `json:"Alias,omitempty" name:"Alias"`
}

func (r *UploadCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadCertRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Cert")
	delete(f, "CertType")
	delete(f, "ProjectId")
	delete(f, "ModuleType")
	delete(f, "Key")
	delete(f, "Alias")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadCertRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadCertResponseParams struct {
	// 证书ID。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UploadCertResponse struct {
	*tchttp.BaseResponse
	Response *UploadCertResponseParams `json:"Response"`
}

func (r *UploadCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadCertResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}