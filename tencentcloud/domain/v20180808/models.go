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

package v20180808

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type BatchModifyDomainInfoRequest struct {
	*tchttp.BaseRequest

	// 批量修改的域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains" list`

	// 模板ID。
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *BatchModifyDomainInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BatchModifyDomainInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BatchModifyDomainInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志ID
		LogId *int64 `json:"LogId,omitempty" name:"LogId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BatchModifyDomainInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BatchModifyDomainInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BatchStatus struct {

	// 批量任务id
	LogId *uint64 `json:"LogId,omitempty" name:"LogId"`

	// 批量任务状态  doing：进行中  success：成功  failed：失败  partial_success：部分成功
	Status *string `json:"Status,omitempty" name:"Status"`

	// 批量任务类型
	BatchAction *string `json:"BatchAction,omitempty" name:"BatchAction"`
}

type CertificateInfo struct {

	// 证件号码。
	CertificateCode *string `json:"CertificateCode,omitempty" name:"CertificateCode"`

	// 证件类型。
	CertificateType *string `json:"CertificateType,omitempty" name:"CertificateType"`

	// 证件照片地址。
	ImgUrl *string `json:"ImgUrl,omitempty" name:"ImgUrl"`
}

type CheckBatchStatusRequest struct {
	*tchttp.BaseRequest

	// 批量任务id数组，最多 200 个
	LogIds []*uint64 `json:"LogIds,omitempty" name:"LogIds" list`
}

func (r *CheckBatchStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckBatchStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckBatchStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 批量任务状态集
		StatusSet []*BatchStatus `json:"StatusSet,omitempty" name:"StatusSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckBatchStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckBatchStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckDomainRequest struct {
	*tchttp.BaseRequest

	// 所查询域名名称
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 年限
	Period *string `json:"Period,omitempty" name:"Period"`
}

func (r *CheckDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckDomainRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CheckDomainResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 所查询域名名称
		DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

		// 是否能够注册
		Available *bool `json:"Available,omitempty" name:"Available"`

		// 不能注册原因
		Reason *string `json:"Reason,omitempty" name:"Reason"`

		// 是否是溢价词
		Premium *bool `json:"Premium,omitempty" name:"Premium"`

		// 域名价格
		Price *uint64 `json:"Price,omitempty" name:"Price"`

		// 是否是敏感词
		BlackWord *bool `json:"BlackWord,omitempty" name:"BlackWord"`

		// 溢价词描述
	// 注意：此字段可能返回 null，表示取不到有效值。
		Describe *string `json:"Describe,omitempty" name:"Describe"`

		// 溢价词的续费价格
	// 注意：此字段可能返回 null，表示取不到有效值。
		FeeRenew *uint64 `json:"FeeRenew,omitempty" name:"FeeRenew"`

		// 域名真实价格
	// 注意：此字段可能返回 null，表示取不到有效值。
		RealPrice *uint64 `json:"RealPrice,omitempty" name:"RealPrice"`

		// 溢价词的转入价格
	// 注意：此字段可能返回 null，表示取不到有效值。
		FeeTransfer *uint64 `json:"FeeTransfer,omitempty" name:"FeeTransfer"`

		// 溢价词的赎回价格
		FeeRestore *uint64 `json:"FeeRestore,omitempty" name:"FeeRestore"`

		// 检测年限
		Period *uint64 `json:"Period,omitempty" name:"Period"`

		// 是否支持北京备案  true 支持  false 不支持
		RecordSupport *bool `json:"RecordSupport,omitempty" name:"RecordSupport"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CheckDomainResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ContactInfo struct {

	// 注册人（中文）
	OrganizationNameCN *string `json:"OrganizationNameCN,omitempty" name:"OrganizationNameCN"`

	// 注册人（英文）
	OrganizationName *string `json:"OrganizationName,omitempty" name:"OrganizationName"`

	// 联系人（中文）
	RegistrantNameCN *string `json:"RegistrantNameCN,omitempty" name:"RegistrantNameCN"`

	// 联系人（英文）
	RegistrantName *string `json:"RegistrantName,omitempty" name:"RegistrantName"`

	// 省份（中文）
	ProvinceCN *string `json:"ProvinceCN,omitempty" name:"ProvinceCN"`

	// 城市（中文）
	CityCN *string `json:"CityCN,omitempty" name:"CityCN"`

	// 街道（中文）
	StreetCN *string `json:"StreetCN,omitempty" name:"StreetCN"`

	// 街道（英文）
	Street *string `json:"Street,omitempty" name:"Street"`

	// 国家（中文）
	CountryCN *string `json:"CountryCN,omitempty" name:"CountryCN"`

	// 联系人手机号
	Telephone *string `json:"Telephone,omitempty" name:"Telephone"`

	// 联系人邮箱
	Email *string `json:"Email,omitempty" name:"Email"`

	// 邮编
	ZipCode *string `json:"ZipCode,omitempty" name:"ZipCode"`

	// 用户类型 E:组织， I:个人
	RegistrantType *string `json:"RegistrantType,omitempty" name:"RegistrantType"`

	// 省份（英文）。作为入参时可以不填
	Province *string `json:"Province,omitempty" name:"Province"`

	// 城市（英文）。作为入参时可以不填
	City *string `json:"City,omitempty" name:"City"`

	// 国家（英文）。作为入参时可以不填
	Country *string `json:"Country,omitempty" name:"Country"`
}

type CreateDomainBatchRequest struct {
	*tchttp.BaseRequest

	// 模板ID
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 购买域名的年限，可选值：[1-10]
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 批量购买的域名,最多为4000个
	Domains []*string `json:"Domains,omitempty" name:"Domains" list`

	// 付费模式 0手动在线付费，1使用余额付费
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`
}

func (r *CreateDomainBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDomainBatchRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDomainBatchResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 批量日志ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		LogId *int64 `json:"LogId,omitempty" name:"LogId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDomainBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDomainBatchResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainBaseInfoRequest struct {
	*tchttp.BaseRequest

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeDomainBaseInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDomainBaseInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainBaseInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 域名信息
		DomainInfo *DomainBaseInfo `json:"DomainInfo,omitempty" name:"DomainInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDomainBaseInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDomainBaseInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainNameListRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，取值范围[1,100]
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDomainNameListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDomainNameListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainNameListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 域名信息集合
	// 注意：此字段可能返回 null，表示取不到有效值。
		DomainSet []*DomainList `json:"DomainSet,omitempty" name:"DomainSet" list`

		// 域名总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDomainNameListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDomainNameListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainPriceListRequest struct {
	*tchttp.BaseRequest

	// 查询价格的后缀列表。默认则为全部后缀
	TldList []*string `json:"TldList,omitempty" name:"TldList" list`

	// 查询购买的年份，默认会列出所有年份的价格
	Year []*int64 `json:"Year,omitempty" name:"Year" list`

	// 域名的购买类型：new  新购，renew 续费，redem 赎回，tran 转入
	Operation []*string `json:"Operation,omitempty" name:"Operation" list`
}

func (r *DescribeDomainPriceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDomainPriceListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDomainPriceListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 域名价格列表
		PriceList []*PriceInfo `json:"PriceList,omitempty" name:"PriceList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDomainPriceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDomainPriceListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTemplateListRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 用户注册类型，默认:all , 个人：I ,企业: E
	Type *string `json:"Type,omitempty" name:"Type"`

	// 认证状态：未实名认证:NotUpload, 实名审核中:InAudit，已实名认证:Approved，实名审核失败:Reject
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *DescribeTemplateListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTemplateListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTemplateListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 模板数量。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 模板详细信息列表。
		TemplateSet []*TemplateInfo `json:"TemplateSet,omitempty" name:"TemplateSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTemplateListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTemplateListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DomainBaseInfo struct {

	// 域名资源ID。
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// 域名名称。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 域名实名认证状态。
	// NotUpload：未实名认证
	// InAudit：实名审核中
	// Approved：实名审核通过
	// Reject：实名审核失败
	RealNameAuditStatus *string `json:"RealNameAuditStatus,omitempty" name:"RealNameAuditStatus"`

	// 域名实名认证不通过原因。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealNameAuditUnpassReason *string `json:"RealNameAuditUnpassReason,omitempty" name:"RealNameAuditUnpassReason"`

	// 域名命名审核状态。
	// NotAudit：命名审核未上传
	// Pending：命名审核待上传
	// Auditing：域名命名审核中
	// Approved：域名命名审核通过
	// Rejected：域名命名审核拒绝
	DomainNameAuditStatus *string `json:"DomainNameAuditStatus,omitempty" name:"DomainNameAuditStatus"`

	// 域名命名审核不通过原因。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainNameAuditUnpassReason *string `json:"DomainNameAuditUnpassReason,omitempty" name:"DomainNameAuditUnpassReason"`

	// 注册时间。
	CreationDate *string `json:"CreationDate,omitempty" name:"CreationDate"`

	// 到期时间
	ExpirationDate *string `json:"ExpirationDate,omitempty" name:"ExpirationDate"`

	// 域名状态。
	// ok：正常
	// serverHold：注册局暂停解析 
	// clientHold：注册商暂停解析
	// pendingTransfer：转移中
	// renewingPeriod：续费期
	// redemptionPeriod：偿还期
	// pendingDelete：删除期
	// serverTransferProhibited：注册局禁止转移
	// serverUpdateProhibited：注册局禁止更新
	// serverDeleteProhibited：注册局禁止删除
	// clientTransferProhibited：注册商禁止转移
	// clientUpdateProhibited：注册商禁止更新
	// clientDeleteProhibited：注册商禁止删除
	DomainStatus []*string `json:"DomainStatus,omitempty" name:"DomainStatus" list`

	// 域名购买状态。
	// ok：正常
	// RegisterPending：待注册
	// RegisterDoing：注册中
	// RegisterFailed：注册失败
	// AboutToExpire: 即将过期
	// RenewPending：已进入续费期，需要进行续费
	// RenewDoing：续费中
	// RedemptionPending：已进入赎回期，需要进行续费
	// RedemptionDoing：赎回中
	// TransferPending：待转入中
	// TransferTransing：转入中
	// TransferFailed：转入失败
	BuyStatus *string `json:"BuyStatus,omitempty" name:"BuyStatus"`

	// 注册类型
	// epp （腾讯云自有域名）
	// xinnet (新网域名)
	RegistrarType *string `json:"RegistrarType,omitempty" name:"RegistrarType"`
}

type DomainList struct {

	// 是否是溢价域名：
	// ture 是    
	// false 不是
	IsPremium *bool `json:"IsPremium,omitempty" name:"IsPremium"`

	// 域名资源ID。
	DomainId *string `json:"DomainId,omitempty" name:"DomainId"`

	// 域名名称。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 是否已设置自动续费 。
	// 0：未设置 
	// 1：已设置
	AutoRenew *uint64 `json:"AutoRenew,omitempty" name:"AutoRenew"`

	// 注册时间。
	CreationDate *string `json:"CreationDate,omitempty" name:"CreationDate"`

	// 到期时间。
	ExpirationDate *string `json:"ExpirationDate,omitempty" name:"ExpirationDate"`

	// 域名后缀
	Tld *string `json:"Tld,omitempty" name:"Tld"`

	// 编码后的后缀（中文会进行编码）
	CodeTld *string `json:"CodeTld,omitempty" name:"CodeTld"`

	// 域名购买状态。
	// ok：正常
	// AboutToExpire: 即将到期
	// RegisterPending：注册中
	// RegisterDoing：注册中
	// RegisterFailed：注册失败
	// RenewPending：续费期
	// RenewDoing：续费中
	// RedemptionPending：赎回期
	// RedemptionDoing：赎回中
	// TransferPending：转入中
	// TransferTransing：转入中
	// TransferFailed：转入失败
	BuyStatus *string `json:"BuyStatus,omitempty" name:"BuyStatus"`
}

type ModifyDomainOwnerBatchRequest struct {
	*tchttp.BaseRequest

	// 要过户的域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains" list`

	// 转入账户的uin。
	NewOwnerUin *string `json:"NewOwnerUin,omitempty" name:"NewOwnerUin"`
}

func (r *ModifyDomainOwnerBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDomainOwnerBatchRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDomainOwnerBatchResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志id
		LogId *uint64 `json:"LogId,omitempty" name:"LogId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDomainOwnerBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDomainOwnerBatchResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PriceInfo struct {

	// 域名后缀，例如.com
	Tld *string `json:"Tld,omitempty" name:"Tld"`

	// 购买年限，范围[1-10]
	Year *uint64 `json:"Year,omitempty" name:"Year"`

	// 域名原价
	Price *uint64 `json:"Price,omitempty" name:"Price"`

	// 域名现价
	RealPrice *uint64 `json:"RealPrice,omitempty" name:"RealPrice"`

	// 商品的购买类型，新购，续费，赎回，转入，续费并转入
	Operation *string `json:"Operation,omitempty" name:"Operation"`
}

type TemplateInfo struct {

	// 模板ID
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 认证状态
	AuditStatus *string `json:"AuditStatus,omitempty" name:"AuditStatus"`

	// 创建时间
	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// 更新时间
	UpdatedOn *string `json:"UpdatedOn,omitempty" name:"UpdatedOn"`

	// 用户UIN
	UserUin *string `json:"UserUin,omitempty" name:"UserUin"`

	// 是否是默认模板
	IsDefault *string `json:"IsDefault,omitempty" name:"IsDefault"`

	// 认证失败原因
	AuditReason *string `json:"AuditReason,omitempty" name:"AuditReason"`

	// 认证信息
	CertificateInfo *CertificateInfo `json:"CertificateInfo,omitempty" name:"CertificateInfo"`

	// 联系人信息
	ContactInfo *ContactInfo `json:"ContactInfo,omitempty" name:"ContactInfo"`

	// 模板是否符合规范， 1是 0 否
	IsValidTemplate *int64 `json:"IsValidTemplate,omitempty" name:"IsValidTemplate"`

	// 不符合规范原因
	InvalidReason *string `json:"InvalidReason,omitempty" name:"InvalidReason"`
}

type TransferInDomainBatchRequest struct {
	*tchttp.BaseRequest

	// 转入的域名名称数组。
	Domains []*string `json:"Domains,omitempty" name:"Domains" list`

	// 域名转移码数组。
	PassWords []*string `json:"PassWords,omitempty" name:"PassWords" list`

	// 模板ID。
	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`

	// 付费模式 0手动在线付费，1使用余额付费。
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`
}

func (r *TransferInDomainBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TransferInDomainBatchRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TransferInDomainBatchResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志ID
		LogId *uint64 `json:"LogId,omitempty" name:"LogId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TransferInDomainBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TransferInDomainBatchResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TransferProhibitionBatchRequest struct {
	*tchttp.BaseRequest

	// 批量操作的域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains" list`

	// 是否开启禁止域名转移。
	// True: 开启禁止域名转移状态。
	// False：关闭禁止域名转移状态。
	Status *bool `json:"Status,omitempty" name:"Status"`
}

func (r *TransferProhibitionBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TransferProhibitionBatchRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TransferProhibitionBatchResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志ID
		LogId *uint64 `json:"LogId,omitempty" name:"LogId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TransferProhibitionBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TransferProhibitionBatchResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateProhibitionBatchRequest struct {
	*tchttp.BaseRequest

	// 批量操作的域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains" list`

	// 是否开启禁止域名更新。
	// True:开启禁止域名更新状态。
	// False：关闭禁止域名更新状态。
	Status *bool `json:"Status,omitempty" name:"Status"`
}

func (r *UpdateProhibitionBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateProhibitionBatchRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateProhibitionBatchResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志ID
		LogId *uint64 `json:"LogId,omitempty" name:"LogId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateProhibitionBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateProhibitionBatchResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
