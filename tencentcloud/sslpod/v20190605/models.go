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

package v20190605

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type CertInfo struct {
	// 证书sha1
	Hash *string `json:"Hash,omitempty" name:"Hash"`

	// 证书通用名称
	CN *string `json:"CN,omitempty" name:"CN"`

	// 备用名称
	SANs *string `json:"SANs,omitempty" name:"SANs"`

	// 公钥算法
	KeyAlgo *string `json:"KeyAlgo,omitempty" name:"KeyAlgo"`

	// 颁发者
	Issuer *string `json:"Issuer,omitempty" name:"Issuer"`

	// 有效期开始
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 有效期结束
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 剩余天数
	Days *int64 `json:"Days,omitempty" name:"Days"`

	// 品牌
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// 信任状态
	TrustStatus *string `json:"TrustStatus,omitempty" name:"TrustStatus"`

	// 证书类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertType *string `json:"CertType,omitempty" name:"CertType"`
}

type ChartHistogram struct {
	// 项目名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 项目值
	Children []*ChartNameValue `json:"Children,omitempty" name:"Children"`
}

type ChartNameValue struct {
	// 图表项名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 图表项值
	Value *int64 `json:"Value,omitempty" name:"Value"`
}

// Predefined struct for user
type CreateDomainRequestParams struct {
	// 监控的服务器类型（0：web，1：smtp，2：imap，3：pops）
	ServerType *int64 `json:"ServerType,omitempty" name:"ServerType"`

	// 添加的域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 添加的端口
	Port *string `json:"Port,omitempty" name:"Port"`

	// 指定域名的IP
	IP *string `json:"IP,omitempty" name:"IP"`

	// 是否开启通知告警；true：开启通知告警，false：关闭通知告警
	Notice *bool `json:"Notice,omitempty" name:"Notice"`

	// 给域名添加标签，多个以逗号隔开
	Tags *string `json:"Tags,omitempty" name:"Tags"`
}

type CreateDomainRequest struct {
	*tchttp.BaseRequest
	
	// 监控的服务器类型（0：web，1：smtp，2：imap，3：pops）
	ServerType *int64 `json:"ServerType,omitempty" name:"ServerType"`

	// 添加的域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 添加的端口
	Port *string `json:"Port,omitempty" name:"Port"`

	// 指定域名的IP
	IP *string `json:"IP,omitempty" name:"IP"`

	// 是否开启通知告警；true：开启通知告警，false：关闭通知告警
	Notice *bool `json:"Notice,omitempty" name:"Notice"`

	// 给域名添加标签，多个以逗号隔开
	Tags *string `json:"Tags,omitempty" name:"Tags"`
}

func (r *CreateDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServerType")
	delete(f, "Domain")
	delete(f, "Port")
	delete(f, "IP")
	delete(f, "Notice")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDomainResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDomainResponse struct {
	*tchttp.BaseResponse
	Response *CreateDomainResponseParams `json:"Response"`
}

func (r *CreateDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DashboardResult struct {
	// 安全等级图表
	SecurityLevelPie []*ChartNameValue `json:"SecurityLevelPie,omitempty" name:"SecurityLevelPie"`

	// 证书品牌图表
	CertBrandsPie []*ChartNameValue `json:"CertBrandsPie,omitempty" name:"CertBrandsPie"`

	// 证书有效时间图表
	CertValidTimePie []*ChartNameValue `json:"CertValidTimePie,omitempty" name:"CertValidTimePie"`

	// 证书类型图表
	CertTypePie []*ChartNameValue `json:"CertTypePie,omitempty" name:"CertTypePie"`

	// ssl bugs图表
	SSLBugsLoopholeHistogram []*ChartHistogram `json:"SSLBugsLoopholeHistogram,omitempty" name:"SSLBugsLoopholeHistogram"`

	// 合规图表
	ComplianceHistogram []*ChartHistogram `json:"ComplianceHistogram,omitempty" name:"ComplianceHistogram"`
}

// Predefined struct for user
type DeleteDomainRequestParams struct {
	// 域名ID，可通过<a href="https://cloud.tencent.com/document/api/1084/49339">搜索域名</a>接口获得
	DomainId *int64 `json:"DomainId,omitempty" name:"DomainId"`
}

type DeleteDomainRequest struct {
	*tchttp.BaseRequest
	
	// 域名ID，可通过<a href="https://cloud.tencent.com/document/api/1084/49339">搜索域名</a>接口获得
	DomainId *int64 `json:"DomainId,omitempty" name:"DomainId"`
}

func (r *DeleteDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDomainResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteDomainResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDomainResponseParams `json:"Response"`
}

func (r *DeleteDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDashboardRequestParams struct {

}

type DescribeDashboardRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeDashboardRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDashboardRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDashboardRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDashboardResponseParams struct {
	// dashboard面板数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *DashboardResult `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDashboardResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDashboardResponseParams `json:"Response"`
}

func (r *DescribeDashboardResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDashboardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainCertsRequestParams struct {
	// 域名ID，可通过搜索域名接口获得
	DomainId *int64 `json:"DomainId,omitempty" name:"DomainId"`
}

type DescribeDomainCertsRequest struct {
	*tchttp.BaseRequest
	
	// 域名ID，可通过搜索域名接口获得
	DomainId *int64 `json:"DomainId,omitempty" name:"DomainId"`
}

func (r *DescribeDomainCertsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainCertsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainCertsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainCertsResponseParams struct {
	// 证书信息
	Data []*CertInfo `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDomainCertsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainCertsResponseParams `json:"Response"`
}

func (r *DescribeDomainCertsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainCertsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainTagsRequestParams struct {

}

type DescribeDomainTagsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeDomainTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainTagsResponseParams struct {
	// Tag数组
	Data []*string `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDomainTagsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainTagsResponseParams `json:"Response"`
}

func (r *DescribeDomainTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDomains struct {
	// 列表数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result []*DomainSiteInfo `json:"Result,omitempty" name:"Result"`

	// 搜索出来的数量
	SearchTotal *int64 `json:"SearchTotal,omitempty" name:"SearchTotal"`

	// 总数
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 允许的监控数量
	AllowMonitoringCount *int64 `json:"AllowMonitoringCount,omitempty" name:"AllowMonitoringCount"`

	// 当前监控的数量
	CurrentMonitoringCount *int64 `json:"CurrentMonitoringCount,omitempty" name:"CurrentMonitoringCount"`

	// 允许添加域名总数
	AllowMaxAddDomain *int64 `json:"AllowMaxAddDomain,omitempty" name:"AllowMaxAddDomain"`
}

// Predefined struct for user
type DescribeDomainsRequestParams struct {
	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 获取数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索的类型有：none，tags，grade，brand，code，hash，limit，domain。
	// 选tags，入参请填Tag，
	// 选grade，入参请填Grade，
	// 选brand，入参请填Brand，
	// 选code，入参请填Code，
	// 选hash，入参请填Hash
	// 选limit，标识只返回数量信息
	// 选domain，入参请填Domain
	SearchType *string `json:"SearchType,omitempty" name:"SearchType"`

	// 标签，多个标签用逗号分隔
	Tag *string `json:"Tag,omitempty" name:"Tag"`

	// 等级
	Grade *string `json:"Grade,omitempty" name:"Grade"`

	// 品牌
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// 混合搜索
	Code *string `json:"Code,omitempty" name:"Code"`

	// 证书指纹
	Hash *string `json:"Hash,omitempty" name:"Hash"`

	// 搜索图标类型
	Item *string `json:"Item,omitempty" name:"Item"`

	// 搜索图标值
	Status *string `json:"Status,omitempty" name:"Status"`

	// 搜索域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

type DescribeDomainsRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 获取数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索的类型有：none，tags，grade，brand，code，hash，limit，domain。
	// 选tags，入参请填Tag，
	// 选grade，入参请填Grade，
	// 选brand，入参请填Brand，
	// 选code，入参请填Code，
	// 选hash，入参请填Hash
	// 选limit，标识只返回数量信息
	// 选domain，入参请填Domain
	SearchType *string `json:"SearchType,omitempty" name:"SearchType"`

	// 标签，多个标签用逗号分隔
	Tag *string `json:"Tag,omitempty" name:"Tag"`

	// 等级
	Grade *string `json:"Grade,omitempty" name:"Grade"`

	// 品牌
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// 混合搜索
	Code *string `json:"Code,omitempty" name:"Code"`

	// 证书指纹
	Hash *string `json:"Hash,omitempty" name:"Hash"`

	// 搜索图标类型
	Item *string `json:"Item,omitempty" name:"Item"`

	// 搜索图标值
	Status *string `json:"Status,omitempty" name:"Status"`

	// 搜索域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DescribeDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchType")
	delete(f, "Tag")
	delete(f, "Grade")
	delete(f, "Brand")
	delete(f, "Code")
	delete(f, "Hash")
	delete(f, "Item")
	delete(f, "Status")
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainsResponseParams struct {
	// 列表数据
	Data *DescribeDomains `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDomainsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainsResponseParams `json:"Response"`
}

func (r *DescribeDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNoticeInfoRequestParams struct {

}

type DescribeNoticeInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeNoticeInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNoticeInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNoticeInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNoticeInfoResponseParams struct {
	// 通知信息结果
	Data *NoticeInfoResult `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNoticeInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNoticeInfoResponseParams `json:"Response"`
}

func (r *DescribeNoticeInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNoticeInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DomainSiteInfo struct {
	// ID标识
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// IP地址
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 是否自动获取IP：true：是，false:否
	AutoIP *bool `json:"AutoIP,omitempty" name:"AutoIP"`

	// 评级
	// "A+"，
	//  "A"，
	// "A-"，
	// "B"，
	// "C"，
	// "D"，
	//  "E"，
	//  "F"，
	// "T"，
	Grade *string `json:"Grade,omitempty" name:"Grade"`

	// 证书品牌
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// 监控服务类型
	// 0 :Web
	// 1: SMTP
	// 2: IMAP
	// 3: POP3
	ServerType *int64 `json:"ServerType,omitempty" name:"ServerType"`

	// 评级Code
	// 0："unknown"，
	// 1："A+"，
	// 2： "A"，
	// 3："A-"，
	// 4："B"，
	// 5："C"，
	// 6："D"，
	// 7： "E"，
	// 8： "F"，
	// 9："T"，
	GradeCode *int64 `json:"GradeCode,omitempty" name:"GradeCode"`

	// 是否监控告警；true：是，false:否
	Notice *bool `json:"Notice,omitempty" name:"Notice"`

	// 账号域名关系ID
	AccountDomainId *int64 `json:"AccountDomainId,omitempty" name:"AccountDomainId"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*string `json:"Tags,omitempty" name:"Tags"`

	// 域名状态:
	// 连接异常，
	// 证书已过期，
	// 证书已吊销，
	// 证书黑名单，
	// 证书域名不匹配，
	// 证书不可信，
	// 证书密钥弱，
	// 证书即将过期，少于7天，
	// 证书即将过期，少于30天，
	// 正常，
	// 部分异常
	Status *string `json:"Status,omitempty" name:"Status"`

	// 域名端口
	Port *string `json:"Port,omitempty" name:"Port"`
}

type LimitInfo struct {
	// 通知类型：
	// limit_emai：邮件
	// limit_wechat：微信
	// limit_phone：手机
	Type *string `json:"Type,omitempty" name:"Type"`

	// 总量
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 已发送
	Sent *int64 `json:"Sent,omitempty" name:"Sent"`
}

// Predefined struct for user
type ModifyDomainTagsRequestParams struct {
	// 账号下域名ID
	AccountDomainId *int64 `json:"AccountDomainId,omitempty" name:"AccountDomainId"`

	// 更新后的tag，多个以逗号隔开
	Tags *string `json:"Tags,omitempty" name:"Tags"`
}

type ModifyDomainTagsRequest struct {
	*tchttp.BaseRequest
	
	// 账号下域名ID
	AccountDomainId *int64 `json:"AccountDomainId,omitempty" name:"AccountDomainId"`

	// 更新后的tag，多个以逗号隔开
	Tags *string `json:"Tags,omitempty" name:"Tags"`
}

func (r *ModifyDomainTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccountDomainId")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDomainTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDomainTagsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDomainTagsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDomainTagsResponseParams `json:"Response"`
}

func (r *ModifyDomainTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NoticeInfoResult struct {
	// 通知ID
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 通知开关信息；0：关闭；15开启
	NoticeType *int64 `json:"NoticeType,omitempty" name:"NoticeType"`

	// 额度信息
	LimitInfos []*LimitInfo `json:"LimitInfos,omitempty" name:"LimitInfos"`
}

// Predefined struct for user
type RefreshDomainRequestParams struct {
	// 域名列表中的ID，可通过搜索域名接口获得
	DomainId *int64 `json:"DomainId,omitempty" name:"DomainId"`
}

type RefreshDomainRequest struct {
	*tchttp.BaseRequest
	
	// 域名列表中的ID，可通过搜索域名接口获得
	DomainId *int64 `json:"DomainId,omitempty" name:"DomainId"`
}

func (r *RefreshDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefreshDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RefreshDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RefreshDomainResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RefreshDomainResponse struct {
	*tchttp.BaseResponse
	Response *RefreshDomainResponseParams `json:"Response"`
}

func (r *RefreshDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefreshDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResolveDomainRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

type ResolveDomainRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *ResolveDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResolveDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResolveDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResolveDomainResponseParams struct {
	// 响应数据
	Data []*string `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ResolveDomainResponse struct {
	*tchttp.BaseResponse
	Response *ResolveDomainResponseParams `json:"Response"`
}

func (r *ResolveDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResolveDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}