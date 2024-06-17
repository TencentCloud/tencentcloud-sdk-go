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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AuctionInfo struct {
	// 竞拍人
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bidder *string `json:"Bidder,omitnil,omitempty" name:"Bidder"`

	// 竞拍时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuctionTime *string `json:"AuctionTime,omitnil,omitempty" name:"AuctionTime"`

	// 竞拍价格
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuctionPrice *float64 `json:"AuctionPrice,omitnil,omitempty" name:"AuctionPrice"`

	// 状态 up: 领先 down: 落后
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type BatchModifyDomainInfoRequestParams struct {
	// 批量修改的域名。
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// 模板ID(可从模板列表接口获取)
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// true： 开启60天内禁止转移注册商锁定
	// false：关闭60天内禁止转移注册商锁定
	// 默认 true
	LockTransfer *bool `json:"LockTransfer,omitnil,omitempty" name:"LockTransfer"`
}

type BatchModifyDomainInfoRequest struct {
	*tchttp.BaseRequest
	
	// 批量修改的域名。
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// 模板ID(可从模板列表接口获取)
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// true： 开启60天内禁止转移注册商锁定
	// false：关闭60天内禁止转移注册商锁定
	// 默认 true
	LockTransfer *bool `json:"LockTransfer,omitnil,omitempty" name:"LockTransfer"`
}

func (r *BatchModifyDomainInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchModifyDomainInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domains")
	delete(f, "TemplateId")
	delete(f, "LockTransfer")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchModifyDomainInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchModifyDomainInfoResponseParams struct {
	// 日志ID
	LogId *int64 `json:"LogId,omitnil,omitempty" name:"LogId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BatchModifyDomainInfoResponse struct {
	*tchttp.BaseResponse
	Response *BatchModifyDomainInfoResponseParams `json:"Response"`
}

func (r *BatchModifyDomainInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchModifyDomainInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchStatus struct {
	// 批量任务id
	LogId *uint64 `json:"LogId,omitnil,omitempty" name:"LogId"`

	// 批量任务状态  doing：进行中  success：成功  failed：失败  partial_success：部分成功
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 批量任务类型
	BatchAction *string `json:"BatchAction,omitnil,omitempty" name:"BatchAction"`
}

// Predefined struct for user
type BidDetailPageRequestParams struct {
	// 业务ID
	BusinessId *string `json:"BusinessId,omitnil,omitempty" name:"BusinessId"`
}

type BidDetailPageRequest struct {
	*tchttp.BaseRequest
	
	// 业务ID
	BusinessId *string `json:"BusinessId,omitnil,omitempty" name:"BusinessId"`
}

func (r *BidDetailPageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BidDetailPageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BusinessId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BidDetailPageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BidDetailPageResponseParams struct {
	// 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 当前域名价格
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentPrice *float64 `json:"CurrentPrice,omitnil,omitempty" name:"CurrentPrice"`

	// 用户上次出价
	// 注意：此字段可能返回 null，表示取不到有效值。
	BidPrice *float64 `json:"BidPrice,omitnil,omitempty" name:"BidPrice"`

	// 当前加价幅度
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentPriceScope *float64 `json:"CurrentPriceScope,omitnil,omitempty" name:"CurrentPriceScope"`

	// 加价幅度区间配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	PriceScope []*PriceScopeConf `json:"PriceScope,omitnil,omitempty" name:"PriceScope"`

	// 用户当前已经支付了的保证金
	// 注意：此字段可能返回 null，表示取不到有效值。
	DepositPrice *float64 `json:"DepositPrice,omitnil,omitempty" name:"DepositPrice"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BidDetailPageResponse struct {
	*tchttp.BaseResponse
	Response *BidDetailPageResponseParams `json:"Response"`
}

func (r *BidDetailPageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BidDetailPageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BidPreDomainsRequestParams struct {
	// 业务ID
	BusinessId *string `json:"BusinessId,omitnil,omitempty" name:"BusinessId"`

	// 价格
	Price *int64 `json:"Price,omitnil,omitempty" name:"Price"`
}

type BidPreDomainsRequest struct {
	*tchttp.BaseRequest
	
	// 业务ID
	BusinessId *string `json:"BusinessId,omitnil,omitempty" name:"BusinessId"`

	// 价格
	Price *int64 `json:"Price,omitnil,omitempty" name:"Price"`
}

func (r *BidPreDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BidPreDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BusinessId")
	delete(f, "Price")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BidPreDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BidPreDomainsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BidPreDomainsResponse struct {
	*tchttp.BaseResponse
	Response *BidPreDomainsResponseParams `json:"Response"`
}

func (r *BidPreDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BidPreDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BiddingAppointResult struct {
	// business_id
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessID *string `json:"BusinessID,omitnil,omitempty" name:"BusinessID"`

	// 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 预定价格
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppointPrice *uint64 `json:"AppointPrice,omitnil,omitempty" name:"AppointPrice"`

	// 预约保证金
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppointBondPrice *uint64 `json:"AppointBondPrice,omitnil,omitempty" name:"AppointBondPrice"`

	// 预约结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppointEndTime *string `json:"AppointEndTime,omitnil,omitempty" name:"AppointEndTime"`

	// 预约人数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppointNum *uint64 `json:"AppointNum,omitnil,omitempty" name:"AppointNum"`

	// 1 已预约，2 竞价中，3 等待出价 4 竞价失败 5 等待支付 6 等待转移，7 转移中 8 交易成功 9 预约持有者赎回 10 竞价持有者赎回 11 其他阶段持有者赎回 12 违约
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type BiddingPreReleaseRequestParams struct {
	// 业务ID
	BusinessId *string `json:"BusinessId,omitnil,omitempty" name:"BusinessId"`

	// 价格
	Price *float64 `json:"Price,omitnil,omitempty" name:"Price"`
}

type BiddingPreReleaseRequest struct {
	*tchttp.BaseRequest
	
	// 业务ID
	BusinessId *string `json:"BusinessId,omitnil,omitempty" name:"BusinessId"`

	// 价格
	Price *float64 `json:"Price,omitnil,omitempty" name:"Price"`
}

func (r *BiddingPreReleaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BiddingPreReleaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BusinessId")
	delete(f, "Price")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BiddingPreReleaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BiddingPreReleaseResponseParams struct {
	// 是否需要额外支付
	IsNeedPay *bool `json:"IsNeedPay,omitnil,omitempty" name:"IsNeedPay"`

	// 计费请求参数，以Json字符串的形式进行返回。
	BillingParam *string `json:"BillingParam,omitnil,omitempty" name:"BillingParam"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BiddingPreReleaseResponse struct {
	*tchttp.BaseResponse
	Response *BiddingPreReleaseResponseParams `json:"Response"`
}

func (r *BiddingPreReleaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BiddingPreReleaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BiddingResult struct {
	// business_id
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessID *string `json:"BusinessID,omitnil,omitempty" name:"BusinessID"`

	// 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 当前价格
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentPrice *uint64 `json:"CurrentPrice,omitnil,omitempty" name:"CurrentPrice"`

	// 当前用户昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentNickname *string `json:"CurrentNickname,omitnil,omitempty" name:"CurrentNickname"`

	// 我的出价
	// 注意：此字段可能返回 null，表示取不到有效值。
	BiddingPrice *uint64 `json:"BiddingPrice,omitnil,omitempty" name:"BiddingPrice"`

	// 竞价保证金
	// 注意：此字段可能返回 null，表示取不到有效值。
	BiddingBondPrice *uint64 `json:"BiddingBondPrice,omitnil,omitempty" name:"BiddingBondPrice"`

	// 竞价结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	BiddingEndTime *string `json:"BiddingEndTime,omitnil,omitempty" name:"BiddingEndTime"`

	// 竞价标识，1 领先，2 落后
	// 注意：此字段可能返回 null，表示取不到有效值。
	BiddingFlag *uint64 `json:"BiddingFlag,omitnil,omitempty" name:"BiddingFlag"`

	// 出价次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	BiddingNum *uint64 `json:"BiddingNum,omitnil,omitempty" name:"BiddingNum"`

	// 2 竞价中  3 等待出价 4 竞价失败 10 竞价持有者赎回
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type BiddingSuccessfulResult struct {
	// 支付结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayEndTime *string `json:"PayEndTime,omitnil,omitempty" name:"PayEndTime"`
}

type CertificateInfo struct {
	// 证件号码。
	CertificateCode *string `json:"CertificateCode,omitnil,omitempty" name:"CertificateCode"`

	// 证件类型。
	// SFZ: 身份证。
	// HZ: 护照。
	// TXZ: 中国港澳居民来往内地通行证。
	// TWSFZ: 中国台湾居民来往大陆通行证。
	// GWSFZ: 外国人永久居留身份证。
	// ORG: 组织机构代码证
	// YYZZ: 工商营业执照。
	// TYDMZ: 统一社会信用代码证书。
	// BDDH: 部队代号
	// JDXKZ: 军队单位对外有偿服务许可证。
	// SYZS: 事业单位法定代表人证书。
	// GWCZDJZ: 外国企业常驻代表机构登记证。
	// STDJZ: 社会团体法定代表人登记证书。
	// ZJDJZ: 宗教活动场所登记证。
	// MBDJZ: 民办非企业单位登记证书。
	// JJDJZ: 基金会法定代表人登记证书。
	// LSXKZ: 律师事务所执业许可证。
	// GWZHDJZ: 外国在华文化中心登记证。
	// GWLYDJZ: 外国政府旅游部门常驻代表机构批准登记证。
	// SFXKZ: 司法鉴定许可证
	// GWJGZJ: 外国机构证件。
	// SHFWJGZ: 社会服务机构登记证书。
	// MBXXXKZ: 民办学校办学许可证。
	// YLJGXKZ: 医疗机构执业许可证。
	// GAJZZ: 中国港澳居住证。
	// TWJZZ: 中国台湾居住证。
	// QTTYDM: 其他-统一社会信用代码证书。
	// GZJGZY: 公证机构执业证。
	CertificateType *string `json:"CertificateType,omitnil,omitempty" name:"CertificateType"`

	// 证件照片地址。
	ImgUrl *string `json:"ImgUrl,omitnil,omitempty" name:"ImgUrl"`

	// 原始照片地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginImgUrl *string `json:"OriginImgUrl,omitnil,omitempty" name:"OriginImgUrl"`

	// 联系人证件号码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegistrantCertificateCode *string `json:"RegistrantCertificateCode,omitnil,omitempty" name:"RegistrantCertificateCode"`

	// 联系人证件类型。
	// SFZ: 身份证。
	// HZ: 护照。
	// TXZ: 中国港澳居民来往内地通行证。
	// TWSFZ: 中国台湾居民来往大陆通行证。
	// GWSFZ: 外国人永久居留身份证。
	// ORG: 组织机构代码证
	// YYZZ: 工商营业执照。
	// TYDMZ: 统一社会信用代码证书。
	// BDDH: 部队代号
	// JDXKZ: 军队单位对外有偿服务许可证。
	// SYZS: 事业单位法定代表人证书。
	// GWCZDJZ: 外国企业常驻代表机构登记证。
	// STDJZ: 社会团体法定代表人登记证书。
	// ZJDJZ: 宗教活动场所登记证。
	// MBDJZ: 民办非企业单位登记证书。
	// JJDJZ: 基金会法定代表人登记证书。
	// LSXKZ: 律师事务所执业许可证。
	// GWZHDJZ: 外国在华文化中心登记证。
	// GWLYDJZ: 外国政府旅游部门常驻代表机构批准登记证。
	// SFXKZ: 司法鉴定许可证
	// GWJGZJ: 外国机构证件。
	// SHFWJGZ: 社会服务机构登记证书。
	// MBXXXKZ: 民办学校办学许可证。
	// YLJGXKZ: 医疗机构执业许可证。
	// GAJZZ: 中国港澳居住证。
	// TWJZZ: 中国台湾居住证。
	// QTTYDM: 其他-统一社会信用代码证书。
	// GZJGZY: 公证机构执业证。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegistrantCertificateType *string `json:"RegistrantCertificateType,omitnil,omitempty" name:"RegistrantCertificateType"`

	// 联系人证件照片地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegistrantImgUrl *string `json:"RegistrantImgUrl,omitnil,omitempty" name:"RegistrantImgUrl"`
}

// Predefined struct for user
type CheckBatchStatusRequestParams struct {
	// 操作日志 ID数组，最多 200 个
	LogIds []*uint64 `json:"LogIds,omitnil,omitempty" name:"LogIds"`
}

type CheckBatchStatusRequest struct {
	*tchttp.BaseRequest
	
	// 操作日志 ID数组，最多 200 个
	LogIds []*uint64 `json:"LogIds,omitnil,omitempty" name:"LogIds"`
}

func (r *CheckBatchStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckBatchStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckBatchStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckBatchStatusResponseParams struct {
	// 批量任务状态集
	StatusSet []*BatchStatus `json:"StatusSet,omitnil,omitempty" name:"StatusSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckBatchStatusResponse struct {
	*tchttp.BaseResponse
	Response *CheckBatchStatusResponseParams `json:"Response"`
}

func (r *CheckBatchStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckBatchStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckDomainRequestParams struct {
	// 所查询域名名称
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 年限。该参数为空时无法查询溢价词域名
	Period *string `json:"Period,omitnil,omitempty" name:"Period"`
}

type CheckDomainRequest struct {
	*tchttp.BaseRequest
	
	// 所查询域名名称
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 年限。该参数为空时无法查询溢价词域名
	Period *string `json:"Period,omitnil,omitempty" name:"Period"`
}

func (r *CheckDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainName")
	delete(f, "Period")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckDomainResponseParams struct {
	// 所查询域名名称
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 是否能够注册
	Available *bool `json:"Available,omitnil,omitempty" name:"Available"`

	// 不能注册原因
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 是否是溢价词
	Premium *bool `json:"Premium,omitnil,omitempty" name:"Premium"`

	// 域名价格
	Price *uint64 `json:"Price,omitnil,omitempty" name:"Price"`

	// 是否是敏感词
	BlackWord *bool `json:"BlackWord,omitnil,omitempty" name:"BlackWord"`

	// 溢价词描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// 溢价词的续费价格
	// 注意：此字段可能返回 null，表示取不到有效值。
	FeeRenew *uint64 `json:"FeeRenew,omitnil,omitempty" name:"FeeRenew"`

	// 域名真实价格, 溢价词时价格跟年限有关，非溢价词时价格为1年的价格
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealPrice *uint64 `json:"RealPrice,omitnil,omitempty" name:"RealPrice"`

	// 溢价词的转入价格
	// 注意：此字段可能返回 null，表示取不到有效值。
	FeeTransfer *uint64 `json:"FeeTransfer,omitnil,omitempty" name:"FeeTransfer"`

	// 溢价词的赎回价格
	FeeRestore *uint64 `json:"FeeRestore,omitnil,omitempty" name:"FeeRestore"`

	// 检测年限
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 是否支持北京备案  true 支持  false 不支持
	RecordSupport *bool `json:"RecordSupport,omitnil,omitempty" name:"RecordSupport"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckDomainResponse struct {
	*tchttp.BaseResponse
	Response *CheckDomainResponseParams `json:"Response"`
}

func (r *CheckDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ContactInfo struct {
	// 注册人（中文）
	OrganizationNameCN *string `json:"OrganizationNameCN,omitnil,omitempty" name:"OrganizationNameCN"`

	// 注册人（英文）
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	// 联系人（中文）
	RegistrantNameCN *string `json:"RegistrantNameCN,omitnil,omitempty" name:"RegistrantNameCN"`

	// 联系人（英文）
	RegistrantName *string `json:"RegistrantName,omitnil,omitempty" name:"RegistrantName"`

	// 省份（中文）
	ProvinceCN *string `json:"ProvinceCN,omitnil,omitempty" name:"ProvinceCN"`

	// 城市（中文）
	CityCN *string `json:"CityCN,omitnil,omitempty" name:"CityCN"`

	// 街道（中文）
	StreetCN *string `json:"StreetCN,omitnil,omitempty" name:"StreetCN"`

	// 街道（英文）
	Street *string `json:"Street,omitnil,omitempty" name:"Street"`

	// 国家（中文）
	CountryCN *string `json:"CountryCN,omitnil,omitempty" name:"CountryCN"`

	// 联系人手机号
	Telephone *string `json:"Telephone,omitnil,omitempty" name:"Telephone"`

	// 联系人邮箱
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 邮编
	ZipCode *string `json:"ZipCode,omitnil,omitempty" name:"ZipCode"`

	// 用户类型 E:组织， I:个人
	RegistrantType *string `json:"RegistrantType,omitnil,omitempty" name:"RegistrantType"`

	// 省份（英文）。作为入参时可以不填
	Province *string `json:"Province,omitnil,omitempty" name:"Province"`

	// 城市（英文）。作为入参时可以不填
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// 国家（英文）。作为入参时可以不填
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`
}

// Predefined struct for user
type CreateCustomDnsHostRequestParams struct {
	// 域名实例ID
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// Dns名称
	DnsName *string `json:"DnsName,omitnil,omitempty" name:"DnsName"`

	// IP地址列表
	IpSet []*string `json:"IpSet,omitnil,omitempty" name:"IpSet"`
}

type CreateCustomDnsHostRequest struct {
	*tchttp.BaseRequest
	
	// 域名实例ID
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// Dns名称
	DnsName *string `json:"DnsName,omitnil,omitempty" name:"DnsName"`

	// IP地址列表
	IpSet []*string `json:"IpSet,omitnil,omitempty" name:"IpSet"`
}

func (r *CreateCustomDnsHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomDnsHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainId")
	delete(f, "DnsName")
	delete(f, "IpSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCustomDnsHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCustomDnsHostResponseParams struct {
	// 异步任务ID
	LogId *uint64 `json:"LogId,omitnil,omitempty" name:"LogId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCustomDnsHostResponse struct {
	*tchttp.BaseResponse
	Response *CreateCustomDnsHostResponseParams `json:"Response"`
}

func (r *CreateCustomDnsHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomDnsHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDomainBatchRequestParams struct {
	// 模板ID。详情请查看：[获取模板列表](https://cloud.tencent.com/document/product/242/48940)
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 购买域名的年限，可选值：[1-10]
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 批量购买的域名,最多为4000个
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// 付费模式 0手动在线付费，1使用余额付费，2使用特惠包
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 自动续费开关。有两个可选值：
	// 0 表示关闭，不自动续费（默认值）
	// 1 表示开启，将自动续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 使用的特惠包ID，PayMode为2时必填
	PackageResourceId *string `json:"PackageResourceId,omitnil,omitempty" name:"PackageResourceId"`

	// 是否开启更新锁：0=默认不开启，1=开启
	UpdateProhibition *int64 `json:"UpdateProhibition,omitnil,omitempty" name:"UpdateProhibition"`

	// 是否开启转移锁：0=默认不开启，1=开启
	TransferProhibition *int64 `json:"TransferProhibition,omitnil,omitempty" name:"TransferProhibition"`

	// 渠道来源，pc/miniprogram/h5等
	ChannelFrom *string `json:"ChannelFrom,omitnil,omitempty" name:"ChannelFrom"`

	// 订单来源，common正常/dianshi_active点石活动等
	OrderFrom *string `json:"OrderFrom,omitnil,omitempty" name:"OrderFrom"`

	// 活动id
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`
}

type CreateDomainBatchRequest struct {
	*tchttp.BaseRequest
	
	// 模板ID。详情请查看：[获取模板列表](https://cloud.tencent.com/document/product/242/48940)
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 购买域名的年限，可选值：[1-10]
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 批量购买的域名,最多为4000个
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// 付费模式 0手动在线付费，1使用余额付费，2使用特惠包
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 自动续费开关。有两个可选值：
	// 0 表示关闭，不自动续费（默认值）
	// 1 表示开启，将自动续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 使用的特惠包ID，PayMode为2时必填
	PackageResourceId *string `json:"PackageResourceId,omitnil,omitempty" name:"PackageResourceId"`

	// 是否开启更新锁：0=默认不开启，1=开启
	UpdateProhibition *int64 `json:"UpdateProhibition,omitnil,omitempty" name:"UpdateProhibition"`

	// 是否开启转移锁：0=默认不开启，1=开启
	TransferProhibition *int64 `json:"TransferProhibition,omitnil,omitempty" name:"TransferProhibition"`

	// 渠道来源，pc/miniprogram/h5等
	ChannelFrom *string `json:"ChannelFrom,omitnil,omitempty" name:"ChannelFrom"`

	// 订单来源，common正常/dianshi_active点石活动等
	OrderFrom *string `json:"OrderFrom,omitnil,omitempty" name:"OrderFrom"`

	// 活动id
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`
}

func (r *CreateDomainBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDomainBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "Period")
	delete(f, "Domains")
	delete(f, "PayMode")
	delete(f, "AutoRenewFlag")
	delete(f, "PackageResourceId")
	delete(f, "UpdateProhibition")
	delete(f, "TransferProhibition")
	delete(f, "ChannelFrom")
	delete(f, "OrderFrom")
	delete(f, "ActivityId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDomainBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDomainBatchResponseParams struct {
	// 批量日志ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogId *int64 `json:"LogId,omitnil,omitempty" name:"LogId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDomainBatchResponse struct {
	*tchttp.BaseResponse
	Response *CreateDomainBatchResponseParams `json:"Response"`
}

func (r *CreateDomainBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDomainBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDomainRedemptionRequestParams struct {
	// 域名 ID
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type CreateDomainRedemptionRequest struct {
	*tchttp.BaseRequest
	
	// 域名 ID
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

func (r *CreateDomainRedemptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDomainRedemptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDomainRedemptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDomainRedemptionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDomainRedemptionResponse struct {
	*tchttp.BaseResponse
	Response *CreateDomainRedemptionResponseParams `json:"Response"`
}

func (r *CreateDomainRedemptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDomainRedemptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePhoneEmailRequestParams struct {
	// 手机号或者邮箱
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 1：手机   2：邮箱
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 验证码(通过SendPhoneEmailCode发送到手机或邮箱的验证码)
	VerifyCode *string `json:"VerifyCode,omitnil,omitempty" name:"VerifyCode"`
}

type CreatePhoneEmailRequest struct {
	*tchttp.BaseRequest
	
	// 手机号或者邮箱
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 1：手机   2：邮箱
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 验证码(通过SendPhoneEmailCode发送到手机或邮箱的验证码)
	VerifyCode *string `json:"VerifyCode,omitnil,omitempty" name:"VerifyCode"`
}

func (r *CreatePhoneEmailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePhoneEmailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Code")
	delete(f, "Type")
	delete(f, "VerifyCode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePhoneEmailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePhoneEmailResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePhoneEmailResponse struct {
	*tchttp.BaseResponse
	Response *CreatePhoneEmailResponseParams `json:"Response"`
}

func (r *CreatePhoneEmailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePhoneEmailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTemplateRequestParams struct {
	// 联系人信息
	ContactInfo *ContactInfo `json:"ContactInfo,omitnil,omitempty" name:"ContactInfo"`

	// 证件信息
	CertificateInfo *CertificateInfo `json:"CertificateInfo,omitnil,omitempty" name:"CertificateInfo"`
}

type CreateTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 联系人信息
	ContactInfo *ContactInfo `json:"ContactInfo,omitnil,omitempty" name:"ContactInfo"`

	// 证件信息
	CertificateInfo *CertificateInfo `json:"CertificateInfo,omitnil,omitempty" name:"CertificateInfo"`
}

func (r *CreateTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ContactInfo")
	delete(f, "CertificateInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTemplateResponseParams struct {
	// 模板信息
	Template *TemplateInfo `json:"Template,omitnil,omitempty" name:"Template"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateTemplateResponseParams `json:"Response"`
}

func (r *CreateTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomDnsHost struct {
	// DNS名称
	DnsName *string `json:"DnsName,omitnil,omitempty" name:"DnsName"`

	// IP地址列表
	IpSet []*string `json:"IpSet,omitnil,omitempty" name:"IpSet"`
}

// Predefined struct for user
type DeleteBiddingRequestParams struct {
	// business_id
	BusinessID *string `json:"BusinessID,omitnil,omitempty" name:"BusinessID"`
}

type DeleteBiddingRequest struct {
	*tchttp.BaseRequest
	
	// business_id
	BusinessID *string `json:"BusinessID,omitnil,omitempty" name:"BusinessID"`
}

func (r *DeleteBiddingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBiddingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BusinessID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBiddingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBiddingResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteBiddingResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBiddingResponseParams `json:"Response"`
}

func (r *DeleteBiddingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBiddingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCustomDnsHostRequestParams struct {
	// 域名实例ID
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// DNS名称
	DnsName *string `json:"DnsName,omitnil,omitempty" name:"DnsName"`
}

type DeleteCustomDnsHostRequest struct {
	*tchttp.BaseRequest
	
	// 域名实例ID
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// DNS名称
	DnsName *string `json:"DnsName,omitnil,omitempty" name:"DnsName"`
}

func (r *DeleteCustomDnsHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCustomDnsHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainId")
	delete(f, "DnsName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCustomDnsHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCustomDnsHostResponseParams struct {
	// 异步任务ID
	LogId *uint64 `json:"LogId,omitnil,omitempty" name:"LogId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCustomDnsHostResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCustomDnsHostResponseParams `json:"Response"`
}

func (r *DeleteCustomDnsHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCustomDnsHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePhoneEmailRequestParams struct {
	// 手机或者邮箱
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 1：手机  2：邮箱
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`
}

type DeletePhoneEmailRequest struct {
	*tchttp.BaseRequest
	
	// 手机或者邮箱
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 1：手机  2：邮箱
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *DeletePhoneEmailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePhoneEmailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Code")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePhoneEmailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePhoneEmailResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeletePhoneEmailResponse struct {
	*tchttp.BaseResponse
	Response *DeletePhoneEmailResponseParams `json:"Response"`
}

func (r *DeletePhoneEmailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePhoneEmailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReservedPreDomainInfoRequestParams struct {
	// 资源ID列表
	ResourceIdList []*string `json:"ResourceIdList,omitnil,omitempty" name:"ResourceIdList"`
}

type DeleteReservedPreDomainInfoRequest struct {
	*tchttp.BaseRequest
	
	// 资源ID列表
	ResourceIdList []*string `json:"ResourceIdList,omitnil,omitempty" name:"ResourceIdList"`
}

func (r *DeleteReservedPreDomainInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReservedPreDomainInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteReservedPreDomainInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReservedPreDomainInfoResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteReservedPreDomainInfoResponse struct {
	*tchttp.BaseResponse
	Response *DeleteReservedPreDomainInfoResponseParams `json:"Response"`
}

func (r *DeleteReservedPreDomainInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReservedPreDomainInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTemplateRequestParams struct {
	// 模板ID(可通过模板信息列表获取)
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DeleteTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板ID(可通过模板信息列表获取)
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *DeleteTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTemplateResponseParams `json:"Response"`
}

func (r *DeleteTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuctionListRequestParams struct {
	// 业务ID
	BusinessId *string `json:"BusinessId,omitnil,omitempty" name:"BusinessId"`

	// 条数，默认10条
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	OffSet *int64 `json:"OffSet,omitnil,omitempty" name:"OffSet"`
}

type DescribeAuctionListRequest struct {
	*tchttp.BaseRequest
	
	// 业务ID
	BusinessId *string `json:"BusinessId,omitnil,omitempty" name:"BusinessId"`

	// 条数，默认10条
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	OffSet *int64 `json:"OffSet,omitnil,omitempty" name:"OffSet"`
}

func (r *DescribeAuctionListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuctionListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BusinessId")
	delete(f, "Limit")
	delete(f, "OffSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuctionListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuctionListResponseParams struct {
	// 竞拍详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuctionList []*AuctionInfo `json:"AuctionList,omitnil,omitempty" name:"AuctionList"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAuctionListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuctionListResponseParams `json:"Response"`
}

func (r *DescribeAuctionListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuctionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBatchOperationLogDetailsRequestParams struct {
	// 日志ID。
	LogId *int64 `json:"LogId,omitnil,omitempty" name:"LogId"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为200。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeBatchOperationLogDetailsRequest struct {
	*tchttp.BaseRequest
	
	// 日志ID。
	LogId *int64 `json:"LogId,omitnil,omitempty" name:"LogId"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为200。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeBatchOperationLogDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBatchOperationLogDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBatchOperationLogDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBatchOperationLogDetailsResponseParams struct {
	// 总数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 日志详情列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainBatchDetailSet []*DomainBatchDetailSet `json:"DomainBatchDetailSet,omitnil,omitempty" name:"DomainBatchDetailSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBatchOperationLogDetailsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBatchOperationLogDetailsResponseParams `json:"Response"`
}

func (r *DescribeBatchOperationLogDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBatchOperationLogDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBatchOperationLogsRequestParams struct {
	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为200。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeBatchOperationLogsRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为200。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeBatchOperationLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBatchOperationLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBatchOperationLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBatchOperationLogsResponseParams struct {
	// 总数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 日志列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainBatchLogSet []*DomainBatchLogSet `json:"DomainBatchLogSet,omitnil,omitempty" name:"DomainBatchLogSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBatchOperationLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBatchOperationLogsResponseParams `json:"Response"`
}

func (r *DescribeBatchOperationLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBatchOperationLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBiddingAppointDetailRequestParams struct {
	// business_id
	BusinessID *string `json:"BusinessID,omitnil,omitempty" name:"BusinessID"`
}

type DescribeBiddingAppointDetailRequest struct {
	*tchttp.BaseRequest
	
	// business_id
	BusinessID *string `json:"BusinessID,omitnil,omitempty" name:"BusinessID"`
}

func (r *DescribeBiddingAppointDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBiddingAppointDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BusinessID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBiddingAppointDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBiddingAppointDetailResponseParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 预约人数
	AppointNum *uint64 `json:"AppointNum,omitnil,omitempty" name:"AppointNum"`

	// 预约开始时间
	AppointStartTime *string `json:"AppointStartTime,omitnil,omitempty" name:"AppointStartTime"`

	// 预约结束时间
	AppointEndTime *string `json:"AppointEndTime,omitnil,omitempty" name:"AppointEndTime"`

	//  注册时间
	RegTime *string `json:"RegTime,omitnil,omitempty" name:"RegTime"`

	// 过期时间
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 删除时间
	DeleteTime *string `json:"DeleteTime,omitnil,omitempty" name:"DeleteTime"`

	// 当前价格
	AppointPrice *uint64 `json:"AppointPrice,omitnil,omitempty" name:"AppointPrice"`

	// 预约保证金
	AppointBondPrice *uint64 `json:"AppointBondPrice,omitnil,omitempty" name:"AppointBondPrice"`

	// 1 已预约，2 竞价中，3 等待出价 4 竞价失败 5 等待支付 6 等待转移，7 转移中 8 交易成功 9 预约持有者赎回 10 竞价持有者赎回 11 其他阶段持有者赎回 12 违约
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 预约保证金是否已经退回
	// yes：退回 no: 未退回
	BiddingBondRefund *string `json:"BiddingBondRefund,omitnil,omitempty" name:"BiddingBondRefund"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBiddingAppointDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBiddingAppointDetailResponseParams `json:"Response"`
}

func (r *DescribeBiddingAppointDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBiddingAppointDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBiddingAppointListRequestParams struct {
	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 状态： 1 已预约 9 预约持有者索回
	Status []*uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 排序字段：AppointEndTime 预约结束时间
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// 排序规则：asc升序，desc降序
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`
}

type DescribeBiddingAppointListRequest struct {
	*tchttp.BaseRequest
	
	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 状态： 1 已预约 9 预约持有者索回
	Status []*uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 排序字段：AppointEndTime 预约结束时间
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// 排序规则：asc升序，desc降序
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`
}

func (r *DescribeBiddingAppointListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBiddingAppointListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Domain")
	delete(f, "Status")
	delete(f, "SortField")
	delete(f, "SortOrder")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBiddingAppointListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBiddingAppointListResponseParams struct {
	// 搜索结果条数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 预约列表
	AppointList []*BiddingAppointResult `json:"AppointList,omitnil,omitempty" name:"AppointList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBiddingAppointListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBiddingAppointListResponseParams `json:"Response"`
}

func (r *DescribeBiddingAppointListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBiddingAppointListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBiddingDetailRequestParams struct {
	// business_id
	BusinessID *string `json:"BusinessID,omitnil,omitempty" name:"BusinessID"`
}

type DescribeBiddingDetailRequest struct {
	*tchttp.BaseRequest
	
	// business_id
	BusinessID *string `json:"BusinessID,omitnil,omitempty" name:"BusinessID"`
}

func (r *DescribeBiddingDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBiddingDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BusinessID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBiddingDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBiddingDetailResponseParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 出价次数
	BiddingNum *uint64 `json:"BiddingNum,omitnil,omitempty" name:"BiddingNum"`

	// 竞价开始时间
	BiddingStartTime *string `json:"BiddingStartTime,omitnil,omitempty" name:"BiddingStartTime"`

	// 竞价结束时间
	BiddingEndTime *string `json:"BiddingEndTime,omitnil,omitempty" name:"BiddingEndTime"`

	//  注册时间
	RegTime *string `json:"RegTime,omitnil,omitempty" name:"RegTime"`

	// 过期时间
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 删除时间
	DeleteTime *string `json:"DeleteTime,omitnil,omitempty" name:"DeleteTime"`

	// 当前价格
	CurrentPrice *uint64 `json:"CurrentPrice,omitnil,omitempty" name:"CurrentPrice"`

	// 当前用户昵称
	CurrentNickname *string `json:"CurrentNickname,omitnil,omitempty" name:"CurrentNickname"`

	// 竞价保证金
	BiddingBondPrice *uint64 `json:"BiddingBondPrice,omitnil,omitempty" name:"BiddingBondPrice"`

	// 2 竞价中  3 等待出价 4 竞价失败 10 竞价持有者赎回
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 竞价标识，1 领先，2 落后
	BiddingFlag *uint64 `json:"BiddingFlag,omitnil,omitempty" name:"BiddingFlag"`

	// 是否退款，yes表示退款，no表示不退款
	BiddingBondRefund *string `json:"BiddingBondRefund,omitnil,omitempty" name:"BiddingBondRefund"`

	// 我的出价
	BiddingPrice *uint64 `json:"BiddingPrice,omitnil,omitempty" name:"BiddingPrice"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBiddingDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBiddingDetailResponseParams `json:"Response"`
}

func (r *DescribeBiddingDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBiddingDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBiddingListRequestParams struct {
	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 2 竞价中  3 等待出价 4 竞价失败 10 竞价持有者赎回
	Status []*uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 排序字段：BiddingEndTime 竞价结束时间	
	// BiddingPrice 我的价格
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// 排序规则：asc升序，desc降序
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`
}

type DescribeBiddingListRequest struct {
	*tchttp.BaseRequest
	
	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 2 竞价中  3 等待出价 4 竞价失败 10 竞价持有者赎回
	Status []*uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 排序字段：BiddingEndTime 竞价结束时间	
	// BiddingPrice 我的价格
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// 排序规则：asc升序，desc降序
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`
}

func (r *DescribeBiddingListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBiddingListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Domain")
	delete(f, "Status")
	delete(f, "SortField")
	delete(f, "SortOrder")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBiddingListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBiddingListResponseParams struct {
	// 搜索结果条数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 竞价列表
	BiddingList []*BiddingResult `json:"BiddingList,omitnil,omitempty" name:"BiddingList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBiddingListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBiddingListResponseParams `json:"Response"`
}

func (r *DescribeBiddingListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBiddingListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBiddingSuccessfulDetailRequestParams struct {
	// business_id
	BusinessID *string `json:"BusinessID,omitnil,omitempty" name:"BusinessID"`
}

type DescribeBiddingSuccessfulDetailRequest struct {
	*tchttp.BaseRequest
	
	// business_id
	BusinessID *string `json:"BusinessID,omitnil,omitempty" name:"BusinessID"`
}

func (r *DescribeBiddingSuccessfulDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBiddingSuccessfulDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BusinessID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBiddingSuccessfulDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBiddingSuccessfulDetailResponseParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 得标时间
	SuccessfulTime *string `json:"SuccessfulTime,omitnil,omitempty" name:"SuccessfulTime"`

	// 得标价格
	SuccessfulPrice *float64 `json:"SuccessfulPrice,omitnil,omitempty" name:"SuccessfulPrice"`

	//  注册时间
	RegTime *string `json:"RegTime,omitnil,omitempty" name:"RegTime"`

	// 过期时间
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 删除时间
	DeleteTime *string `json:"DeleteTime,omitnil,omitempty" name:"DeleteTime"`

	// 付款结束时间
	PayEndTime *string `json:"PayEndTime,omitnil,omitempty" name:"PayEndTime"`

	// 保证金，是否退款，yes表示退款，no表示不退款
	BiddingBondRefund *string `json:"BiddingBondRefund,omitnil,omitempty" name:"BiddingBondRefund"`

	// 保证金
	BiddingBondPrice *float64 `json:"BiddingBondPrice,omitnil,omitempty" name:"BiddingBondPrice"`

	// 状态：5 等待支付 6 等待转移， 7 转移中，8 交易成功，11 尾款阶段持有者索回，12 已违约
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBiddingSuccessfulDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBiddingSuccessfulDetailResponseParams `json:"Response"`
}

func (r *DescribeBiddingSuccessfulDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBiddingSuccessfulDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBiddingSuccessfulListRequestParams struct {
	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 状态：5 等待支付 6 等待转移， 7 转移中，8 交易成功，11 尾款阶段持有者索回，12 已违约
	Status []*uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 排序字段：SuccessfulTime 预约结束时间
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// 排序规则：asc升序，desc降序
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`
}

type DescribeBiddingSuccessfulListRequest struct {
	*tchttp.BaseRequest
	
	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 状态：5 等待支付 6 等待转移， 7 转移中，8 交易成功，11 尾款阶段持有者索回，12 已违约
	Status []*uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 排序字段：SuccessfulTime 预约结束时间
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// 排序规则：asc升序，desc降序
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`
}

func (r *DescribeBiddingSuccessfulListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBiddingSuccessfulListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Domain")
	delete(f, "Status")
	delete(f, "SortField")
	delete(f, "SortOrder")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBiddingSuccessfulListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBiddingSuccessfulListResponseParams struct {
	// 搜索结果条数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 得标列表
	SuccessfulList []*BiddingSuccessfulResult `json:"SuccessfulList,omitnil,omitempty" name:"SuccessfulList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBiddingSuccessfulListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBiddingSuccessfulListResponseParams `json:"Response"`
}

func (r *DescribeBiddingSuccessfulListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBiddingSuccessfulListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomDnsHostSetRequestParams struct {
	// 域名实例ID(域名基本信息或我的域名列表接口可获取)
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 返回数量，默认为20，取值范围[1,100]
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeCustomDnsHostSetRequest struct {
	*tchttp.BaseRequest
	
	// 域名实例ID(域名基本信息或我的域名列表接口可获取)
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 返回数量，默认为20，取值范围[1,100]
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeCustomDnsHostSetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomDnsHostSetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomDnsHostSetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomDnsHostSetResponseParams struct {
	// 自定义DNS Host 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DnsHostSet []*CustomDnsHost `json:"DnsHostSet,omitnil,omitempty" name:"DnsHostSet"`

	// 自定义DNS Host总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCustomDnsHostSetResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomDnsHostSetResponseParams `json:"Response"`
}

func (r *DescribeCustomDnsHostSetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomDnsHostSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainBaseInfoRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

type DescribeDomainBaseInfoRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

func (r *DescribeDomainBaseInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainBaseInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainBaseInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainBaseInfoResponseParams struct {
	// 域名信息
	DomainInfo *DomainBaseInfo `json:"DomainInfo,omitnil,omitempty" name:"DomainInfo"`

	// 用户Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDomainBaseInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainBaseInfoResponseParams `json:"Response"`
}

func (r *DescribeDomainBaseInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainBaseInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainNameListRequestParams struct {
	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，取值范围[1,100]
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeDomainNameListRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，取值范围[1,100]
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeDomainNameListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainNameListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainNameListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainNameListResponseParams struct {
	// 域名信息集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainSet []*DomainList `json:"DomainSet,omitnil,omitempty" name:"DomainSet"`

	// 域名总数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDomainNameListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainNameListResponseParams `json:"Response"`
}

func (r *DescribeDomainNameListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainNameListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainPriceListRequestParams struct {
	// 查询价格的后缀列表。默认则为全部后缀
	TldList []*string `json:"TldList,omitnil,omitempty" name:"TldList"`

	// 查询购买的年份，默认会列出所有年份的价格
	Year []*int64 `json:"Year,omitnil,omitempty" name:"Year"`

	// 域名的购买类型：new  新购，renew 续费，redem 赎回，tran 转入
	Operation []*string `json:"Operation,omitnil,omitempty" name:"Operation"`
}

type DescribeDomainPriceListRequest struct {
	*tchttp.BaseRequest
	
	// 查询价格的后缀列表。默认则为全部后缀
	TldList []*string `json:"TldList,omitnil,omitempty" name:"TldList"`

	// 查询购买的年份，默认会列出所有年份的价格
	Year []*int64 `json:"Year,omitnil,omitempty" name:"Year"`

	// 域名的购买类型：new  新购，renew 续费，redem 赎回，tran 转入
	Operation []*string `json:"Operation,omitnil,omitempty" name:"Operation"`
}

func (r *DescribeDomainPriceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainPriceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TldList")
	delete(f, "Year")
	delete(f, "Operation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainPriceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainPriceListResponseParams struct {
	// 域名价格列表
	PriceList []*PriceInfo `json:"PriceList,omitnil,omitempty" name:"PriceList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDomainPriceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainPriceListResponseParams `json:"Response"`
}

func (r *DescribeDomainPriceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainPriceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainSimpleInfoRequestParams struct {
	// 域名
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`
}

type DescribeDomainSimpleInfoRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`
}

func (r *DescribeDomainSimpleInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainSimpleInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainSimpleInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainSimpleInfoResponseParams struct {
	// 域名信息
	DomainInfo *DomainSimpleInfo `json:"DomainInfo,omitnil,omitempty" name:"DomainInfo"`

	// 账号ID
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDomainSimpleInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainSimpleInfoResponseParams `json:"Response"`
}

func (r *DescribeDomainSimpleInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainSimpleInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePayWaitDetailRequestParams struct {
	// 业务ID
	BusinessId *string `json:"BusinessId,omitnil,omitempty" name:"BusinessId"`
}

type DescribePayWaitDetailRequest struct {
	*tchttp.BaseRequest
	
	// 业务ID
	BusinessId *string `json:"BusinessId,omitnil,omitempty" name:"BusinessId"`
}

func (r *DescribePayWaitDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePayWaitDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BusinessId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePayWaitDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePayWaitDetailResponseParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名类型
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 支付结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 域名注册时间
	RegTime *string `json:"RegTime,omitnil,omitempty" name:"RegTime"`

	// 域名成交价格
	Price *float64 `json:"Price,omitnil,omitempty" name:"Price"`

	// 待退还保证金
	RetDeposit *float64 `json:"RetDeposit,omitnil,omitempty" name:"RetDeposit"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePayWaitDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribePayWaitDetailResponseParams `json:"Response"`
}

func (r *DescribePayWaitDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePayWaitDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePhoneEmailListRequestParams struct {
	// 0：所有类型  1：手机  2：邮箱，默认0
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，取值范围[1,200]
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 手机或者邮箱，用于精确搜索
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`
}

type DescribePhoneEmailListRequest struct {
	*tchttp.BaseRequest
	
	// 0：所有类型  1：手机  2：邮箱，默认0
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，取值范围[1,200]
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 手机或者邮箱，用于精确搜索
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`
}

func (r *DescribePhoneEmailListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePhoneEmailListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Code")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePhoneEmailListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePhoneEmailListResponseParams struct {
	// 手机或者邮箱列表
	PhoneEmailList []*PhoneEmailData `json:"PhoneEmailList,omitnil,omitempty" name:"PhoneEmailList"`

	// 总数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePhoneEmailListResponse struct {
	*tchttp.BaseResponse
	Response *DescribePhoneEmailListResponseParams `json:"Response"`
}

func (r *DescribePhoneEmailListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePhoneEmailListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePreAuctionListRequestParams struct {
	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 条数
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribePreAuctionListRequest struct {
	*tchttp.BaseRequest
	
	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 条数
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribePreAuctionListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePreAuctionListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePreAuctionListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePreAuctionListResponseParams struct {
	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 预释放竞价列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PreAuctionList []*PreAuctionInfo `json:"PreAuctionList,omitnil,omitempty" name:"PreAuctionList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePreAuctionListResponse struct {
	*tchttp.BaseResponse
	Response *DescribePreAuctionListResponseParams `json:"Response"`
}

func (r *DescribePreAuctionListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePreAuctionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePreDomainListRequestParams struct {
	// 页码，默认为1
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 条数，默认为20
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 用于结束时间筛选
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 用户指定上架时间筛选
	UpTime *string `json:"UpTime,omitnil,omitempty" name:"UpTime"`
}

type DescribePreDomainListRequest struct {
	*tchttp.BaseRequest
	
	// 页码，默认为1
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 条数，默认为20
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 用于结束时间筛选
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 用户指定上架时间筛选
	UpTime *string `json:"UpTime,omitnil,omitempty" name:"UpTime"`
}

func (r *DescribePreDomainListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePreDomainListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Page")
	delete(f, "Size")
	delete(f, "EndTime")
	delete(f, "UpTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePreDomainListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePreDomainListResponseParams struct {
	// 预释放预约列表数据
	ReservedDomainList []*ReservedDomainInfo `json:"ReservedDomainList,omitnil,omitempty" name:"ReservedDomainList"`

	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePreDomainListResponse struct {
	*tchttp.BaseResponse
	Response *DescribePreDomainListResponseParams `json:"Response"`
}

func (r *DescribePreDomainListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePreDomainListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePreReleaseListRequestParams struct {
	// 关键词
	Keywords *string `json:"Keywords,omitnil,omitempty" name:"Keywords"`

	// 搜索关键字，开头
	DomainStart *bool `json:"DomainStart,omitnil,omitempty" name:"DomainStart"`

	// 搜索关键字结尾
	DomainEnd *bool `json:"DomainEnd,omitnil,omitempty" name:"DomainEnd"`

	// 排序
	Sort *int64 `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 起始价格
	PriceStart *float64 `json:"PriceStart,omitnil,omitempty" name:"PriceStart"`

	// 结束价格
	PriceEnd *float64 `json:"PriceEnd,omitnil,omitempty" name:"PriceEnd"`

	// 起始域名长度
	LengthStart *int64 `json:"LengthStart,omitnil,omitempty" name:"LengthStart"`

	// 结束域名长度
	LengthEnd *int64 `json:"LengthEnd,omitnil,omitempty" name:"LengthEnd"`

	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页显示数
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 后缀
	Suffix []*int64 `json:"Suffix,omitnil,omitempty" name:"Suffix"`

	// 一级分类
	ClassOne *int64 `json:"ClassOne,omitnil,omitempty" name:"ClassOne"`

	// 二级分类
	ClassTwo []*int64 `json:"ClassTwo,omitnil,omitempty" name:"ClassTwo"`

	// 三级分类
	ClassThree []*int64 `json:"ClassThree,omitnil,omitempty" name:"ClassThree"`

	// 四级分类
	ClassFour []*int64 `json:"ClassFour,omitnil,omitempty" name:"ClassFour"`

	// 排除关键字，开头
	FilterStart *bool `json:"FilterStart,omitnil,omitempty" name:"FilterStart"`

	// 排除关键字，结尾
	FilterEnd *bool `json:"FilterEnd,omitnil,omitempty" name:"FilterEnd"`

	// 排除关键字
	FilterWords *string `json:"FilterWords,omitnil,omitempty" name:"FilterWords"`

	// 交易类型
	TransType *int64 `json:"TransType,omitnil,omitempty" name:"TransType"`

	// 搜索白金域名
	IsTop *bool `json:"IsTop,omitnil,omitempty" name:"IsTop"`

	// 结束时间排序啊 desc:倒序 asc:正序
	EndTimeSort *string `json:"EndTimeSort,omitnil,omitempty" name:"EndTimeSort"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribePreReleaseListRequest struct {
	*tchttp.BaseRequest
	
	// 关键词
	Keywords *string `json:"Keywords,omitnil,omitempty" name:"Keywords"`

	// 搜索关键字，开头
	DomainStart *bool `json:"DomainStart,omitnil,omitempty" name:"DomainStart"`

	// 搜索关键字结尾
	DomainEnd *bool `json:"DomainEnd,omitnil,omitempty" name:"DomainEnd"`

	// 排序
	Sort *int64 `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 起始价格
	PriceStart *float64 `json:"PriceStart,omitnil,omitempty" name:"PriceStart"`

	// 结束价格
	PriceEnd *float64 `json:"PriceEnd,omitnil,omitempty" name:"PriceEnd"`

	// 起始域名长度
	LengthStart *int64 `json:"LengthStart,omitnil,omitempty" name:"LengthStart"`

	// 结束域名长度
	LengthEnd *int64 `json:"LengthEnd,omitnil,omitempty" name:"LengthEnd"`

	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页显示数
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 后缀
	Suffix []*int64 `json:"Suffix,omitnil,omitempty" name:"Suffix"`

	// 一级分类
	ClassOne *int64 `json:"ClassOne,omitnil,omitempty" name:"ClassOne"`

	// 二级分类
	ClassTwo []*int64 `json:"ClassTwo,omitnil,omitempty" name:"ClassTwo"`

	// 三级分类
	ClassThree []*int64 `json:"ClassThree,omitnil,omitempty" name:"ClassThree"`

	// 四级分类
	ClassFour []*int64 `json:"ClassFour,omitnil,omitempty" name:"ClassFour"`

	// 排除关键字，开头
	FilterStart *bool `json:"FilterStart,omitnil,omitempty" name:"FilterStart"`

	// 排除关键字，结尾
	FilterEnd *bool `json:"FilterEnd,omitnil,omitempty" name:"FilterEnd"`

	// 排除关键字
	FilterWords *string `json:"FilterWords,omitnil,omitempty" name:"FilterWords"`

	// 交易类型
	TransType *int64 `json:"TransType,omitnil,omitempty" name:"TransType"`

	// 搜索白金域名
	IsTop *bool `json:"IsTop,omitnil,omitempty" name:"IsTop"`

	// 结束时间排序啊 desc:倒序 asc:正序
	EndTimeSort *string `json:"EndTimeSort,omitnil,omitempty" name:"EndTimeSort"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *DescribePreReleaseListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePreReleaseListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Keywords")
	delete(f, "DomainStart")
	delete(f, "DomainEnd")
	delete(f, "Sort")
	delete(f, "PriceStart")
	delete(f, "PriceEnd")
	delete(f, "LengthStart")
	delete(f, "LengthEnd")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Suffix")
	delete(f, "ClassOne")
	delete(f, "ClassTwo")
	delete(f, "ClassThree")
	delete(f, "ClassFour")
	delete(f, "FilterStart")
	delete(f, "FilterEnd")
	delete(f, "FilterWords")
	delete(f, "TransType")
	delete(f, "IsTop")
	delete(f, "EndTimeSort")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePreReleaseListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePreReleaseListResponseParams struct {
	// 数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 预释放列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PreReleaseList []*PreReleaseInfo `json:"PreReleaseList,omitnil,omitempty" name:"PreReleaseList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePreReleaseListResponse struct {
	*tchttp.BaseResponse
	Response *DescribePreReleaseListResponseParams `json:"Response"`
}

func (r *DescribePreReleaseListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePreReleaseListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReservedBidInfoRequestParams struct {
	// 业务ID
	BusinessId *string `json:"BusinessId,omitnil,omitempty" name:"BusinessId"`
}

type DescribeReservedBidInfoRequest struct {
	*tchttp.BaseRequest
	
	// 业务ID
	BusinessId *string `json:"BusinessId,omitnil,omitempty" name:"BusinessId"`
}

func (r *DescribeReservedBidInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReservedBidInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BusinessId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReservedBidInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReservedBidInfoResponseParams struct {
	// 竞价领先价格
	UpPrice *int64 `json:"UpPrice,omitnil,omitempty" name:"UpPrice"`

	// 请求用户当前价格
	Price *int64 `json:"Price,omitnil,omitempty" name:"Price"`

	// 领先用户
	UpUser *string `json:"UpUser,omitnil,omitempty" name:"UpUser"`

	// 竞价详细数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	BidList []*ReserveBidInfo `json:"BidList,omitnil,omitempty" name:"BidList"`

	// 竞价结束时间
	BidEndTime *string `json:"BidEndTime,omitnil,omitempty" name:"BidEndTime"`

	// 是否领先
	IsUp *bool `json:"IsUp,omitnil,omitempty" name:"IsUp"`

	// 下次出价金额
	NextPrice *int64 `json:"NextPrice,omitnil,omitempty" name:"NextPrice"`

	// 状态：1. 等待竞价 2.竞价中 3.竞价结束
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeReservedBidInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReservedBidInfoResponseParams `json:"Response"`
}

func (r *DescribeReservedBidInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReservedBidInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReservedPreDomainInfoRequestParams struct {
	// 域名,每次最多支持500条域名查询
	DomainList []*string `json:"DomainList,omitnil,omitempty" name:"DomainList"`

	// 状态，用于筛选，可不填写(1. 成功 2. 失败（失败Reason字段将会被赋值）3. 域名交割中 4. 域名交割完成 5. 预约 6. 竞价)
	ReservedStatus *int64 `json:"ReservedStatus,omitnil,omitempty" name:"ReservedStatus"`

	// 根据预约时间排序，仅支持："desc","asc"。
	ReservedTimeSort *string `json:"ReservedTimeSort,omitnil,omitempty" name:"ReservedTimeSort"`

	// 条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeReservedPreDomainInfoRequest struct {
	*tchttp.BaseRequest
	
	// 域名,每次最多支持500条域名查询
	DomainList []*string `json:"DomainList,omitnil,omitempty" name:"DomainList"`

	// 状态，用于筛选，可不填写(1. 成功 2. 失败（失败Reason字段将会被赋值）3. 域名交割中 4. 域名交割完成 5. 预约 6. 竞价)
	ReservedStatus *int64 `json:"ReservedStatus,omitnil,omitempty" name:"ReservedStatus"`

	// 根据预约时间排序，仅支持："desc","asc"。
	ReservedTimeSort *string `json:"ReservedTimeSort,omitnil,omitempty" name:"ReservedTimeSort"`

	// 条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeReservedPreDomainInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReservedPreDomainInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainList")
	delete(f, "ReservedStatus")
	delete(f, "ReservedTimeSort")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReservedPreDomainInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReservedPreDomainInfoResponseParams struct {
	// 预释放预约列表
	ReservedPreDomainInfoList []*ReservedPreDomainInfo `json:"ReservedPreDomainInfoList,omitnil,omitempty" name:"ReservedPreDomainInfoList"`

	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeReservedPreDomainInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReservedPreDomainInfoResponseParams `json:"Response"`
}

func (r *DescribeReservedPreDomainInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReservedPreDomainInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTemplateListRequestParams struct {
	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 用户注册类型，默认:all , 个人：I ,企业: E
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 认证状态：未实名审核:NotUpload, 实名审核中:InAudit，已实名审核:Approved，实名审核失败:Reject，更新手机邮箱:NotVerified。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 关键字，用于域名所有者筛选
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

type DescribeTemplateListRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 用户注册类型，默认:all , 个人：I ,企业: E
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 认证状态：未实名审核:NotUpload, 实名审核中:InAudit，已实名审核:Approved，实名审核失败:Reject，更新手机邮箱:NotVerified。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 关键字，用于域名所有者筛选
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

func (r *DescribeTemplateListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTemplateListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Type")
	delete(f, "Status")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTemplateListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTemplateListResponseParams struct {
	// 模板数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 模板详细信息列表。
	TemplateSet []*TemplateInfo `json:"TemplateSet,omitnil,omitempty" name:"TemplateSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTemplateListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTemplateListResponseParams `json:"Response"`
}

func (r *DescribeTemplateListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTemplateListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTemplateRequestParams struct {
	// 模板ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DescribeTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *DescribeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTemplateResponseParams struct {
	// 模板信息
	Template *TemplateInfo `json:"Template,omitnil,omitempty" name:"Template"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTemplateResponseParams `json:"Response"`
}

func (r *DescribeTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTldListRequestParams struct {

}

type DescribeTldListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeTldListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTldListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTldListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTldListResponseParams struct {
	// 支持的后缀列表
	List []*string `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTldListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTldListResponseParams `json:"Response"`
}

func (r *DescribeTldListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTldListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUnPreDomainDetailRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

type DescribeUnPreDomainDetailRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

func (r *DescribeUnPreDomainDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUnPreDomainDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUnPreDomainDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUnPreDomainDetailResponseParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 预约人数
	PreCount *int64 `json:"PreCount,omitnil,omitempty" name:"PreCount"`

	// 域名注册时间
	RegTime *string `json:"RegTime,omitnil,omitempty" name:"RegTime"`

	// 域名删除时间
	DeleteTime *string `json:"DeleteTime,omitnil,omitempty" name:"DeleteTime"`

	// 到期时间
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 域名状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 域名价格
	CurrentPrice *float64 `json:"CurrentPrice,omitnil,omitempty" name:"CurrentPrice"`

	// 域名保证金
	AppointBondPrice *float64 `json:"AppointBondPrice,omitnil,omitempty" name:"AppointBondPrice"`

	// 是否已经预约
	IsAppoint *bool `json:"IsAppoint,omitnil,omitempty" name:"IsAppoint"`

	// 业务ID
	BusinessId *string `json:"BusinessId,omitnil,omitempty" name:"BusinessId"`

	// 是否为原持有者域名
	IsDomainUser *bool `json:"IsDomainUser,omitnil,omitempty" name:"IsDomainUser"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUnPreDomainDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUnPreDomainDetailResponseParams `json:"Response"`
}

func (r *DescribeUnPreDomainDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUnPreDomainDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DomainBaseInfo struct {
	// 域名资源ID。
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 域名名称。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 域名实名认证状态。
	// NotUpload：未实名认证
	// InAudit：实名审核中
	// Approved：实名审核通过
	// Reject：实名审核失败
	// NoAudit: 无需实名认证
	RealNameAuditStatus *string `json:"RealNameAuditStatus,omitnil,omitempty" name:"RealNameAuditStatus"`

	// 域名实名认证不通过原因。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealNameAuditUnpassReason *string `json:"RealNameAuditUnpassReason,omitnil,omitempty" name:"RealNameAuditUnpassReason"`

	// 域名命名审核状态。
	// NotAudit：命名审核未上传
	// Pending：命名审核待上传
	// Auditing：域名命名审核中
	// Approved：域名命名审核通过
	// Rejected：域名命名审核拒绝
	DomainNameAuditStatus *string `json:"DomainNameAuditStatus,omitnil,omitempty" name:"DomainNameAuditStatus"`

	// 域名命名审核不通过原因。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainNameAuditUnpassReason *string `json:"DomainNameAuditUnpassReason,omitnil,omitempty" name:"DomainNameAuditUnpassReason"`

	// 注册时间。
	CreationDate *string `json:"CreationDate,omitnil,omitempty" name:"CreationDate"`

	// 到期时间
	ExpirationDate *string `json:"ExpirationDate,omitnil,omitempty" name:"ExpirationDate"`

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
	// serverRenewProhibited: 注册局禁止续费
	// clientRenewProhobited: 注册商禁止续费
	DomainStatus []*string `json:"DomainStatus,omitnil,omitempty" name:"DomainStatus"`

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
	BuyStatus *string `json:"BuyStatus,omitnil,omitempty" name:"BuyStatus"`

	// 注册商类型
	// epp: DNSPod, Inc.（烟台帝思普网络科技有限公司）
	// qcloud: Tencent Cloud Computing (Beijing) Limited Liability Company（腾讯云计算（北京）有限责任公司）
	// yunxun: Guangzhou Yunxun Information Technology Co., Ltd.（广州云讯信息科技有限公司）
	// xinnet: Xin Net Technology Corporation（北京新网数码信息技术有限公司）
	RegistrarType *string `json:"RegistrarType,omitnil,omitempty" name:"RegistrarType"`

	// 域名绑定的ns
	NameServer []*string `json:"NameServer,omitnil,omitempty" name:"NameServer"`

	// true：开启锁定
	// false：关闭锁定
	LockTransfer *bool `json:"LockTransfer,omitnil,omitempty" name:"LockTransfer"`

	// 锁定结束时间
	LockEndTime *string `json:"LockEndTime,omitnil,omitempty" name:"LockEndTime"`
}

type DomainBatchDetailSet struct {
	// 详情ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 类型  new: 注册域名 batch_transfer_prohibition_on:开启禁止转移  batch_transfer_prohibition_off:关闭禁止转移 batch_update_prohibition_on:开启禁止更新   batch_update_prohibition_off:关闭禁止更新
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 执行状态：
	// doing 执行中。
	// failed 操作失败。
	// success  操作成功。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 创建时间
	CreatedOn *string `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// 更新时间
	UpdatedOn *string `json:"UpdatedOn,omitnil,omitempty" name:"UpdatedOn"`

	// 订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	BigDealId *string `json:"BigDealId,omitnil,omitempty" name:"BigDealId"`
}

type DomainBatchLogSet struct {
	// 日志ID
	LogId *int64 `json:"LogId,omitnil,omitempty" name:"LogId"`

	// 数量
	Number *int64 `json:"Number,omitnil,omitempty" name:"Number"`

	// 执行状态：
	// doing 执行中。
	// done 执行完成。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 提交时间
	CreatedOn *string `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// 批量操作成功个数
	Success *int64 `json:"Success,omitnil,omitempty" name:"Success"`

	// 批量操作处理中个数
	Doing *int64 `json:"Doing,omitnil,omitempty" name:"Doing"`

	// 批量操作失败个数
	Failed *int64 `json:"Failed,omitnil,omitempty" name:"Failed"`
}

type DomainList struct {
	// 是否是溢价域名：
	// ture 是    
	// false 不是
	IsPremium *bool `json:"IsPremium,omitnil,omitempty" name:"IsPremium"`

	// 域名资源ID。
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 域名名称。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 是否已设置自动续费 。
	// 0：未设置 
	// 1：已设置
	// 2：设置后，关闭
	AutoRenew *uint64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// 注册时间。
	CreationDate *string `json:"CreationDate,omitnil,omitempty" name:"CreationDate"`

	// 到期时间。
	ExpirationDate *string `json:"ExpirationDate,omitnil,omitempty" name:"ExpirationDate"`

	// 域名后缀
	Tld *string `json:"Tld,omitnil,omitempty" name:"Tld"`

	// 编码后的后缀（中文会进行编码）
	CodeTld *string `json:"CodeTld,omitnil,omitempty" name:"CodeTld"`

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
	BuyStatus *string `json:"BuyStatus,omitnil,omitempty" name:"BuyStatus"`
}

type DomainSimpleInfo struct {
	// 域名资源ID。
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 域名名称。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 域名实名认证状态。
	// NotUpload：未实名认证
	// InAudit：实名审核中
	// Approved：实名审核通过
	// Reject：实名审核失败
	// NoAudit: 无需实名认证
	RealNameAuditStatus *string `json:"RealNameAuditStatus,omitnil,omitempty" name:"RealNameAuditStatus"`

	// 域名实名认证不通过原因。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealNameAuditUnpassReason *string `json:"RealNameAuditUnpassReason,omitnil,omitempty" name:"RealNameAuditUnpassReason"`

	// 域名命名审核状态。
	// NotAudit：命名审核未上传
	// Pending：命名审核待上传
	// Auditing：域名命名审核中
	// Approved：域名命名审核通过
	// Rejected：域名命名审核拒绝
	DomainNameAuditStatus *string `json:"DomainNameAuditStatus,omitnil,omitempty" name:"DomainNameAuditStatus"`

	// 域名命名审核不通过原因。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainNameAuditUnpassReason *string `json:"DomainNameAuditUnpassReason,omitnil,omitempty" name:"DomainNameAuditUnpassReason"`

	// 注册时间。
	CreationDate *string `json:"CreationDate,omitnil,omitempty" name:"CreationDate"`

	// 到期时间
	ExpirationDate *string `json:"ExpirationDate,omitnil,omitempty" name:"ExpirationDate"`

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
	// serverRenewProhibited: 注册局禁止续费
	// clientRenewProhobited: 注册商禁止续费
	DomainStatus []*string `json:"DomainStatus,omitnil,omitempty" name:"DomainStatus"`

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
	BuyStatus *string `json:"BuyStatus,omitnil,omitempty" name:"BuyStatus"`

	// 注册商类型
	// epp: DNSPod, Inc.（烟台帝思普网络科技有限公司）
	// qcloud: Tencent Cloud Computing (Beijing) Limited Liability Company（腾讯云计算（北京）有限责任公司）
	// yunxun: Guangzhou Yunxun Information Technology Co., Ltd.（广州云讯信息科技有限公司）
	// xinnet: Xin Net Technology Corporation（北京新网数码信息技术有限公司）
	RegistrarType *string `json:"RegistrarType,omitnil,omitempty" name:"RegistrarType"`

	// 域名绑定的ns
	NameServer []*string `json:"NameServer,omitnil,omitempty" name:"NameServer"`

	// true：开启锁定
	// false：关闭锁定
	LockTransfer *bool `json:"LockTransfer,omitnil,omitempty" name:"LockTransfer"`

	// 锁定结束时间
	LockEndTime *string `json:"LockEndTime,omitnil,omitempty" name:"LockEndTime"`

	// 认证类型：I=个人，E=企业
	RegistrantType *string `json:"RegistrantType,omitnil,omitempty" name:"RegistrantType"`

	// 域名所有者，中文
	OrganizationNameCN *string `json:"OrganizationNameCN,omitnil,omitempty" name:"OrganizationNameCN"`

	// 域名所有者，英文
	OrganizationName *string `json:"OrganizationName,omitnil,omitempty" name:"OrganizationName"`

	// 域名联系人，中文
	RegistrantNameCN *string `json:"RegistrantNameCN,omitnil,omitempty" name:"RegistrantNameCN"`

	// 域名联系人，英文
	RegistrantName *string `json:"RegistrantName,omitnil,omitempty" name:"RegistrantName"`
}

type FailReservedDomainInfo struct {
	// 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 预约失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailReason *string `json:"FailReason,omitnil,omitempty" name:"FailReason"`
}

// Predefined struct for user
type ModifyCustomDnsHostRequestParams struct {
	// 域名实例ID
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// DNS名称
	DnsName *string `json:"DnsName,omitnil,omitempty" name:"DnsName"`

	// IP地址列表
	IpSet []*string `json:"IpSet,omitnil,omitempty" name:"IpSet"`
}

type ModifyCustomDnsHostRequest struct {
	*tchttp.BaseRequest
	
	// 域名实例ID
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// DNS名称
	DnsName *string `json:"DnsName,omitnil,omitempty" name:"DnsName"`

	// IP地址列表
	IpSet []*string `json:"IpSet,omitnil,omitempty" name:"IpSet"`
}

func (r *ModifyCustomDnsHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomDnsHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainId")
	delete(f, "DnsName")
	delete(f, "IpSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCustomDnsHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomDnsHostResponseParams struct {
	// 异步任务ID
	LogId *uint64 `json:"LogId,omitnil,omitempty" name:"LogId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCustomDnsHostResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCustomDnsHostResponseParams `json:"Response"`
}

func (r *ModifyCustomDnsHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomDnsHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDomainDNSBatchRequestParams struct {
	// 批量操作的域名。
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// 域名DNS 数组。
	Dns []*string `json:"Dns,omitnil,omitempty" name:"Dns"`
}

type ModifyDomainDNSBatchRequest struct {
	*tchttp.BaseRequest
	
	// 批量操作的域名。
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// 域名DNS 数组。
	Dns []*string `json:"Dns,omitnil,omitempty" name:"Dns"`
}

func (r *ModifyDomainDNSBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainDNSBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domains")
	delete(f, "Dns")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDomainDNSBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDomainDNSBatchResponseParams struct {
	// 日志ID。
	LogId *uint64 `json:"LogId,omitnil,omitempty" name:"LogId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDomainDNSBatchResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDomainDNSBatchResponseParams `json:"Response"`
}

func (r *ModifyDomainDNSBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainDNSBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDomainOwnerBatchRequestParams struct {
	// 要过户的域名。
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// 转入账户的uin。
	NewOwnerUin *string `json:"NewOwnerUin,omitnil,omitempty" name:"NewOwnerUin"`

	// 是否同时转移对应的 DNS 解析域名，默认false
	TransferDns *bool `json:"TransferDns,omitnil,omitempty" name:"TransferDns"`

	// 转入账户的appid。
	NewOwnerAppId *string `json:"NewOwnerAppId,omitnil,omitempty" name:"NewOwnerAppId"`
}

type ModifyDomainOwnerBatchRequest struct {
	*tchttp.BaseRequest
	
	// 要过户的域名。
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// 转入账户的uin。
	NewOwnerUin *string `json:"NewOwnerUin,omitnil,omitempty" name:"NewOwnerUin"`

	// 是否同时转移对应的 DNS 解析域名，默认false
	TransferDns *bool `json:"TransferDns,omitnil,omitempty" name:"TransferDns"`

	// 转入账户的appid。
	NewOwnerAppId *string `json:"NewOwnerAppId,omitnil,omitempty" name:"NewOwnerAppId"`
}

func (r *ModifyDomainOwnerBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainOwnerBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domains")
	delete(f, "NewOwnerUin")
	delete(f, "TransferDns")
	delete(f, "NewOwnerAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDomainOwnerBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDomainOwnerBatchResponseParams struct {
	// 日志id
	LogId *uint64 `json:"LogId,omitnil,omitempty" name:"LogId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDomainOwnerBatchResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDomainOwnerBatchResponseParams `json:"Response"`
}

func (r *ModifyDomainOwnerBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainOwnerBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIntlCustomDnsHostRequestParams struct {
	// 域名ID
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// DNS Host
	DnsName *string `json:"DnsName,omitnil,omitempty" name:"DnsName"`

	// IP地址
	IpSet []*string `json:"IpSet,omitnil,omitempty" name:"IpSet"`
}

type ModifyIntlCustomDnsHostRequest struct {
	*tchttp.BaseRequest
	
	// 域名ID
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// DNS Host
	DnsName *string `json:"DnsName,omitnil,omitempty" name:"DnsName"`

	// IP地址
	IpSet []*string `json:"IpSet,omitnil,omitempty" name:"IpSet"`
}

func (r *ModifyIntlCustomDnsHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIntlCustomDnsHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainId")
	delete(f, "DnsName")
	delete(f, "IpSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyIntlCustomDnsHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIntlCustomDnsHostResponseParams struct {
	// 任务ID
	LogId *int64 `json:"LogId,omitnil,omitempty" name:"LogId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyIntlCustomDnsHostResponse struct {
	*tchttp.BaseResponse
	Response *ModifyIntlCustomDnsHostResponseParams `json:"Response"`
}

func (r *ModifyIntlCustomDnsHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIntlCustomDnsHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTemplateRequestParams struct {
	// 证件信息
	CertificateInfo *CertificateInfo `json:"CertificateInfo,omitnil,omitempty" name:"CertificateInfo"`

	// 联系人信息
	ContactInfo *ContactInfo `json:"ContactInfo,omitnil,omitempty" name:"ContactInfo"`

	// 模板ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type ModifyTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 证件信息
	CertificateInfo *CertificateInfo `json:"CertificateInfo,omitnil,omitempty" name:"CertificateInfo"`

	// 联系人信息
	ContactInfo *ContactInfo `json:"ContactInfo,omitnil,omitempty" name:"ContactInfo"`

	// 模板ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *ModifyTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateInfo")
	delete(f, "ContactInfo")
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTemplateResponseParams struct {
	// 模板信息
	Template *TemplateInfo `json:"Template,omitnil,omitempty" name:"Template"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTemplateResponseParams `json:"Response"`
}

func (r *ModifyTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PhoneEmailData struct {
	// 手机号或者邮箱
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 1：手机  2：邮箱
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 创建时间
	CreatedOn *string `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// 1=控制台校验，2=第三方校验
	CheckStatus *int64 `json:"CheckStatus,omitnil,omitempty" name:"CheckStatus"`
}

type PreAuctionInfo struct {
	// 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 竞价倒计时
	// 注意：此字段可能返回 null，表示取不到有效值。
	BiddingTime *string `json:"BiddingTime,omitnil,omitempty" name:"BiddingTime"`

	// 出价次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	BidCount *int64 `json:"BidCount,omitnil,omitempty" name:"BidCount"`

	// 当前价格
	// 注意：此字段可能返回 null，表示取不到有效值。
	Price *float64 `json:"Price,omitnil,omitempty" name:"Price"`

	// 用户操作 bid：出价 "noAction"：无法操作
	// 注意：此字段可能返回 null，表示取不到有效值。
	Op *string `json:"Op,omitnil,omitempty" name:"Op"`

	// 业务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessId *string `json:"BusinessId,omitnil,omitempty" name:"BusinessId"`
}

type PreReleaseInfo struct {
	// 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 预订倒计时
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReservationTime *string `json:"ReservationTime,omitnil,omitempty" name:"ReservationTime"`

	// 域名注册时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegTime *string `json:"RegTime,omitnil,omitempty" name:"RegTime"`

	// 域名删除时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DelTime *string `json:"DelTime,omitnil,omitempty" name:"DelTime"`

	// 当前人数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentPeople *int64 `json:"CurrentPeople,omitnil,omitempty" name:"CurrentPeople"`

	// 当前价格
	// 注意：此字段可能返回 null，表示取不到有效值。
	Price *float64 `json:"Price,omitnil,omitempty" name:"Price"`

	// 是否收藏
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsFollow *bool `json:"IsFollow,omitnil,omitempty" name:"IsFollow"`

	// 是否已经预约
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAppoint *bool `json:"IsAppoint,omitnil,omitempty" name:"IsAppoint"`

	// 业务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessId *string `json:"BusinessId,omitnil,omitempty" name:"BusinessId"`

	// 是否为原持有者
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDomainUser *bool `json:"IsDomainUser,omitnil,omitempty" name:"IsDomainUser"`
}

type PriceInfo struct {
	// 域名后缀，例如.com
	Tld *string `json:"Tld,omitnil,omitempty" name:"Tld"`

	// 购买年限，范围[1-10]
	Year *uint64 `json:"Year,omitnil,omitempty" name:"Year"`

	// 域名原价
	Price *uint64 `json:"Price,omitnil,omitempty" name:"Price"`

	// 域名现价
	RealPrice *uint64 `json:"RealPrice,omitnil,omitempty" name:"RealPrice"`

	// 商品的购买类型，新购，续费，赎回，转入，续费并转入
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`
}

type PriceScopeConf struct {
	// 最高价格
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxPrice *float64 `json:"MaxPrice,omitnil,omitempty" name:"MaxPrice"`

	// 最低价格
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinPrice *float64 `json:"MinPrice,omitnil,omitempty" name:"MinPrice"`

	// 价格幅度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Price *float64 `json:"Price,omitnil,omitempty" name:"Price"`

	// 保证金
	// 注意：此字段可能返回 null，表示取不到有效值。
	DepositPrice *float64 `json:"DepositPrice,omitnil,omitempty" name:"DepositPrice"`
}

// Predefined struct for user
type RenewDomainBatchRequestParams struct {
	// 域名续费的年限。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 批量续费的域名。
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// 付费模式 0手动在线付费，1使用余额付费，2使用特惠包。
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 自动续费开关。有三个可选值：
	// 0 表示关闭，不自动续费
	// 1 表示开启，将自动续费
	// 2 表示不处理，保留域名原有状态（默认值）
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 特惠包ID
	PackageResourceId *string `json:"PackageResourceId,omitnil,omitempty" name:"PackageResourceId"`

	// 渠道来源，pc/miniprogram/h5等
	ChannelFrom *string `json:"ChannelFrom,omitnil,omitempty" name:"ChannelFrom"`

	// 订单来源，common正常/dianshi_active点石活动等
	OrderFrom *string `json:"OrderFrom,omitnil,omitempty" name:"OrderFrom"`

	// 活动id
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`
}

type RenewDomainBatchRequest struct {
	*tchttp.BaseRequest
	
	// 域名续费的年限。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 批量续费的域名。
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// 付费模式 0手动在线付费，1使用余额付费，2使用特惠包。
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 自动续费开关。有三个可选值：
	// 0 表示关闭，不自动续费
	// 1 表示开启，将自动续费
	// 2 表示不处理，保留域名原有状态（默认值）
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 特惠包ID
	PackageResourceId *string `json:"PackageResourceId,omitnil,omitempty" name:"PackageResourceId"`

	// 渠道来源，pc/miniprogram/h5等
	ChannelFrom *string `json:"ChannelFrom,omitnil,omitempty" name:"ChannelFrom"`

	// 订单来源，common正常/dianshi_active点石活动等
	OrderFrom *string `json:"OrderFrom,omitnil,omitempty" name:"OrderFrom"`

	// 活动id
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`
}

func (r *RenewDomainBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewDomainBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Period")
	delete(f, "Domains")
	delete(f, "PayMode")
	delete(f, "AutoRenewFlag")
	delete(f, "PackageResourceId")
	delete(f, "ChannelFrom")
	delete(f, "OrderFrom")
	delete(f, "ActivityId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewDomainBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewDomainBatchResponseParams struct {
	// 操作日志ID。
	LogId *int64 `json:"LogId,omitnil,omitempty" name:"LogId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RenewDomainBatchResponse struct {
	*tchttp.BaseResponse
	Response *RenewDomainBatchResponseParams `json:"Response"`
}

func (r *RenewDomainBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewDomainBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReserveBidInfo struct {
	// 用户
	// 注意：此字段可能返回 null，表示取不到有效值。
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 出价
	// 注意：此字段可能返回 null，表示取不到有效值。
	Price *int64 `json:"Price,omitnil,omitempty" name:"Price"`

	// 出价时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	BidTime *string `json:"BidTime,omitnil,omitempty" name:"BidTime"`

	// 当前状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	BidStatus *string `json:"BidStatus,omitnil,omitempty" name:"BidStatus"`
}

type ReservedDomainInfo struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 注册时间
	RegTime *string `json:"RegTime,omitnil,omitempty" name:"RegTime"`

	// 到期时间
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 续费时间结束
	RenewEndTime *string `json:"RenewEndTime,omitnil,omitempty" name:"RenewEndTime"`

	// 赎回结束时间
	RestoreEndTime *string `json:"RestoreEndTime,omitnil,omitempty" name:"RestoreEndTime"`

	// 域名预约结束时间
	ReservedEndTime *string `json:"ReservedEndTime,omitnil,omitempty" name:"ReservedEndTime"`
}

type ReservedPreDomainInfo struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 1. 预定成功 2. 预定失败（预定失败Reason字段将会被赋值）3. 域名交割中 4. 域名交割完成
	ReservedStatus *int64 `json:"ReservedStatus,omitnil,omitempty" name:"ReservedStatus"`

	// 域名预定失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailReason *string `json:"FailReason,omitnil,omitempty" name:"FailReason"`

	// 预计变更所有权时间（仅用于参考，实际时间会存在误差）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChangeOwnerTime *string `json:"ChangeOwnerTime,omitnil,omitempty" name:"ChangeOwnerTime"`

	// 注册时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegTime *string `json:"RegTime,omitnil,omitempty" name:"RegTime"`

	// 到期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 资源ID，用于删除资源信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 业务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessId *string `json:"BusinessId,omitnil,omitempty" name:"BusinessId"`
}

// Predefined struct for user
type ReservedPreDomainsRequestParams struct {
	// 预约预释放域名列表
	DomainList []*string `json:"DomainList,omitnil,omitempty" name:"DomainList"`

	// 模板ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 结束后是否自动支付尾款，默认开启 传入1关闭
	IsAutoPay *int64 `json:"IsAutoPay,omitnil,omitempty" name:"IsAutoPay"`

	// 结束后是否自动进行梯度保证金扣除，默认开启 传入1关闭
	IsBidAutoPay *int64 `json:"IsBidAutoPay,omitnil,omitempty" name:"IsBidAutoPay"`
}

type ReservedPreDomainsRequest struct {
	*tchttp.BaseRequest
	
	// 预约预释放域名列表
	DomainList []*string `json:"DomainList,omitnil,omitempty" name:"DomainList"`

	// 模板ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 结束后是否自动支付尾款，默认开启 传入1关闭
	IsAutoPay *int64 `json:"IsAutoPay,omitnil,omitempty" name:"IsAutoPay"`

	// 结束后是否自动进行梯度保证金扣除，默认开启 传入1关闭
	IsBidAutoPay *int64 `json:"IsBidAutoPay,omitnil,omitempty" name:"IsBidAutoPay"`
}

func (r *ReservedPreDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReservedPreDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainList")
	delete(f, "TemplateId")
	delete(f, "IsAutoPay")
	delete(f, "IsBidAutoPay")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReservedPreDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReservedPreDomainsResponseParams struct {
	// 预定成功域名列表
	SucDomainList []*string `json:"SucDomainList,omitnil,omitempty" name:"SucDomainList"`

	// 预定失败域名列表
	FailDomainList []*FailReservedDomainInfo `json:"FailDomainList,omitnil,omitempty" name:"FailDomainList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReservedPreDomainsResponse struct {
	*tchttp.BaseResponse
	Response *ReservedPreDomainsResponseParams `json:"Response"`
}

func (r *ReservedPreDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReservedPreDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendPhoneEmailCodeRequestParams struct {
	// 手机或者邮箱号。
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 1：手机  2：邮箱。
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`
}

type SendPhoneEmailCodeRequest struct {
	*tchttp.BaseRequest
	
	// 手机或者邮箱号。
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 1：手机  2：邮箱。
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *SendPhoneEmailCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendPhoneEmailCodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Code")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendPhoneEmailCodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendPhoneEmailCodeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SendPhoneEmailCodeResponse struct {
	*tchttp.BaseResponse
	Response *SendPhoneEmailCodeResponseParams `json:"Response"`
}

func (r *SendPhoneEmailCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendPhoneEmailCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetDomainAutoRenewRequestParams struct {
	// 域名ID 例如：domain-123abc
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// AutoRenew 有三个可选值：
	//  0：不设置自动续费
	// 1：设置自动续费
	// 2：设置到期后不续费
	AutoRenew *uint64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`
}

type SetDomainAutoRenewRequest struct {
	*tchttp.BaseRequest
	
	// 域名ID 例如：domain-123abc
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// AutoRenew 有三个可选值：
	//  0：不设置自动续费
	// 1：设置自动续费
	// 2：设置到期后不续费
	AutoRenew *uint64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`
}

func (r *SetDomainAutoRenewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetDomainAutoRenewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainId")
	delete(f, "AutoRenew")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetDomainAutoRenewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetDomainAutoRenewResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetDomainAutoRenewResponse struct {
	*tchttp.BaseResponse
	Response *SetDomainAutoRenewResponseParams `json:"Response"`
}

func (r *SetDomainAutoRenewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetDomainAutoRenewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SyncCustomDnsHostRequestParams struct {
	// 域名实例ID
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type SyncCustomDnsHostRequest struct {
	*tchttp.BaseRequest
	
	// 域名实例ID
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

func (r *SyncCustomDnsHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncCustomDnsHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SyncCustomDnsHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SyncCustomDnsHostResponseParams struct {
	// 异步任务ID
	LogId *uint64 `json:"LogId,omitnil,omitempty" name:"LogId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SyncCustomDnsHostResponse struct {
	*tchttp.BaseResponse
	Response *SyncCustomDnsHostResponseParams `json:"Response"`
}

func (r *SyncCustomDnsHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncCustomDnsHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TemplateInfo struct {
	// 模板ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 认证状态：未实名认证:NotUpload, 实名审核中:InAudit，已实名认证:Approved，实名审核失败:Reject
	AuditStatus *string `json:"AuditStatus,omitnil,omitempty" name:"AuditStatus"`

	// 创建时间
	CreatedOn *string `json:"CreatedOn,omitnil,omitempty" name:"CreatedOn"`

	// 更新时间
	UpdatedOn *string `json:"UpdatedOn,omitnil,omitempty" name:"UpdatedOn"`

	// 用户UIN
	UserUin *string `json:"UserUin,omitnil,omitempty" name:"UserUin"`

	// 是否是默认模板: 是:yes，否:no
	IsDefault *string `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`

	// 认证失败原因
	AuditReason *string `json:"AuditReason,omitnil,omitempty" name:"AuditReason"`

	// 认证信息
	CertificateInfo *CertificateInfo `json:"CertificateInfo,omitnil,omitempty" name:"CertificateInfo"`

	// 联系人信息
	ContactInfo *ContactInfo `json:"ContactInfo,omitnil,omitempty" name:"ContactInfo"`

	// 模板是否符合规范， 1是 0 否
	IsValidTemplate *int64 `json:"IsValidTemplate,omitnil,omitempty" name:"IsValidTemplate"`

	// 不符合规范原因
	InvalidReason *string `json:"InvalidReason,omitnil,omitempty" name:"InvalidReason"`

	// 是包含黑名单手机或邮箱
	IsBlack *bool `json:"IsBlack,omitnil,omitempty" name:"IsBlack"`
}

// Predefined struct for user
type TransferInDomainBatchRequestParams struct {
	// 转入的域名名称数组。
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// 域名转移码数组。
	PassWords []*string `json:"PassWords,omitnil,omitempty" name:"PassWords"`

	// 模板ID。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 付费模式 0手动在线付费，1使用余额付费。
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 自动续费开关。有两个可选值：
	// 0 表示关闭，不自动续费（默认值）
	// 1 表示开启，将自动续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// true： 开启60天内禁止转移注册商锁定
	// false：关闭60天内禁止转移注册商锁定
	// 默认 true
	LockTransfer *bool `json:"LockTransfer,omitnil,omitempty" name:"LockTransfer"`

	// 是否开启更新锁：0=默认不开启，1=开启
	UpdateProhibition *int64 `json:"UpdateProhibition,omitnil,omitempty" name:"UpdateProhibition"`

	// 是否开启转移锁：0=默认不开启，1=开启
	TransferProhibition *int64 `json:"TransferProhibition,omitnil,omitempty" name:"TransferProhibition"`

	// 渠道来源，pc/miniprogram/h5等
	ChannelFrom *string `json:"ChannelFrom,omitnil,omitempty" name:"ChannelFrom"`

	// 订单来源，common正常/dianshi_active点石活动等
	OrderFrom *string `json:"OrderFrom,omitnil,omitempty" name:"OrderFrom"`

	// 活动id
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`
}

type TransferInDomainBatchRequest struct {
	*tchttp.BaseRequest
	
	// 转入的域名名称数组。
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// 域名转移码数组。
	PassWords []*string `json:"PassWords,omitnil,omitempty" name:"PassWords"`

	// 模板ID。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 付费模式 0手动在线付费，1使用余额付费。
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 自动续费开关。有两个可选值：
	// 0 表示关闭，不自动续费（默认值）
	// 1 表示开启，将自动续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// true： 开启60天内禁止转移注册商锁定
	// false：关闭60天内禁止转移注册商锁定
	// 默认 true
	LockTransfer *bool `json:"LockTransfer,omitnil,omitempty" name:"LockTransfer"`

	// 是否开启更新锁：0=默认不开启，1=开启
	UpdateProhibition *int64 `json:"UpdateProhibition,omitnil,omitempty" name:"UpdateProhibition"`

	// 是否开启转移锁：0=默认不开启，1=开启
	TransferProhibition *int64 `json:"TransferProhibition,omitnil,omitempty" name:"TransferProhibition"`

	// 渠道来源，pc/miniprogram/h5等
	ChannelFrom *string `json:"ChannelFrom,omitnil,omitempty" name:"ChannelFrom"`

	// 订单来源，common正常/dianshi_active点石活动等
	OrderFrom *string `json:"OrderFrom,omitnil,omitempty" name:"OrderFrom"`

	// 活动id
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`
}

func (r *TransferInDomainBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TransferInDomainBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domains")
	delete(f, "PassWords")
	delete(f, "TemplateId")
	delete(f, "PayMode")
	delete(f, "AutoRenewFlag")
	delete(f, "LockTransfer")
	delete(f, "UpdateProhibition")
	delete(f, "TransferProhibition")
	delete(f, "ChannelFrom")
	delete(f, "OrderFrom")
	delete(f, "ActivityId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TransferInDomainBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TransferInDomainBatchResponseParams struct {
	// 日志ID
	LogId *uint64 `json:"LogId,omitnil,omitempty" name:"LogId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TransferInDomainBatchResponse struct {
	*tchttp.BaseResponse
	Response *TransferInDomainBatchResponseParams `json:"Response"`
}

func (r *TransferInDomainBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TransferInDomainBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TransferProhibitionBatchRequestParams struct {
	// 批量操作的域名。
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// 是否开启禁止域名转移。
	// True: 开启禁止域名转移状态。
	// False：关闭禁止域名转移状态。
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

type TransferProhibitionBatchRequest struct {
	*tchttp.BaseRequest
	
	// 批量操作的域名。
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// 是否开启禁止域名转移。
	// True: 开启禁止域名转移状态。
	// False：关闭禁止域名转移状态。
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *TransferProhibitionBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TransferProhibitionBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domains")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TransferProhibitionBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TransferProhibitionBatchResponseParams struct {
	// 日志ID
	LogId *uint64 `json:"LogId,omitnil,omitempty" name:"LogId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TransferProhibitionBatchResponse struct {
	*tchttp.BaseResponse
	Response *TransferProhibitionBatchResponseParams `json:"Response"`
}

func (r *TransferProhibitionBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TransferProhibitionBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateProhibitionBatchRequestParams struct {
	// 批量操作的域名。
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// 是否开启禁止域名更新。
	// True:开启禁止域名更新状态。
	// False：关闭禁止域名更新状态。
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

type UpdateProhibitionBatchRequest struct {
	*tchttp.BaseRequest
	
	// 批量操作的域名。
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// 是否开启禁止域名更新。
	// True:开启禁止域名更新状态。
	// False：关闭禁止域名更新状态。
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *UpdateProhibitionBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateProhibitionBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domains")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateProhibitionBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateProhibitionBatchResponseParams struct {
	// 日志ID
	LogId *uint64 `json:"LogId,omitnil,omitempty" name:"LogId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateProhibitionBatchResponse struct {
	*tchttp.BaseResponse
	Response *UpdateProhibitionBatchResponseParams `json:"Response"`
}

func (r *UpdateProhibitionBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateProhibitionBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadImageRequestParams struct {
	// 资质照片，照片的base64编码。
	ImageFile *string `json:"ImageFile,omitnil,omitempty" name:"ImageFile"`
}

type UploadImageRequest struct {
	*tchttp.BaseRequest
	
	// 资质照片，照片的base64编码。
	ImageFile *string `json:"ImageFile,omitnil,omitempty" name:"ImageFile"`
}

func (r *UploadImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageFile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadImageResponseParams struct {
	// 资质照片地址。
	AccessUrl *string `json:"AccessUrl,omitnil,omitempty" name:"AccessUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UploadImageResponse struct {
	*tchttp.BaseResponse
	Response *UploadImageResponseParams `json:"Response"`
}

func (r *UploadImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}