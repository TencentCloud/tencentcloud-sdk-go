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
