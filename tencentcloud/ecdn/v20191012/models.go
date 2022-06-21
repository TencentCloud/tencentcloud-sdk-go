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

package v20191012

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type AddEcdnDomainRequestParams struct {
	// 域名。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 源站配置。
	Origin *Origin `json:"Origin,omitempty" name:"Origin"`

	// 域名加速区域，mainland，overseas或global，分别表示中国境内加速，海外加速或全球加速。
	Area *string `json:"Area,omitempty" name:"Area"`

	// 项目id，默认0。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// IP黑白名单配置。
	IpFilter *IpFilter `json:"IpFilter,omitempty" name:"IpFilter"`

	// IP限频配置。
	IpFreqLimit *IpFreqLimit `json:"IpFreqLimit,omitempty" name:"IpFreqLimit"`

	// 源站响应头部配置。
	ResponseHeader *ResponseHeader `json:"ResponseHeader,omitempty" name:"ResponseHeader"`

	// 节点缓存配置。
	CacheKey *CacheKey `json:"CacheKey,omitempty" name:"CacheKey"`

	// 缓存规则配置。
	Cache *Cache `json:"Cache,omitempty" name:"Cache"`

	// Https配置。
	Https *Https `json:"Https,omitempty" name:"Https"`

	// 访问协议强制跳转配置。
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitempty" name:"ForceRedirect"`

	// 域名绑定的标签
	Tag []*Tag `json:"Tag,omitempty" name:"Tag"`

	// WebSocket配置
	WebSocket *WebSocket `json:"WebSocket,omitempty" name:"WebSocket"`
}

type AddEcdnDomainRequest struct {
	*tchttp.BaseRequest
	
	// 域名。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 源站配置。
	Origin *Origin `json:"Origin,omitempty" name:"Origin"`

	// 域名加速区域，mainland，overseas或global，分别表示中国境内加速，海外加速或全球加速。
	Area *string `json:"Area,omitempty" name:"Area"`

	// 项目id，默认0。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// IP黑白名单配置。
	IpFilter *IpFilter `json:"IpFilter,omitempty" name:"IpFilter"`

	// IP限频配置。
	IpFreqLimit *IpFreqLimit `json:"IpFreqLimit,omitempty" name:"IpFreqLimit"`

	// 源站响应头部配置。
	ResponseHeader *ResponseHeader `json:"ResponseHeader,omitempty" name:"ResponseHeader"`

	// 节点缓存配置。
	CacheKey *CacheKey `json:"CacheKey,omitempty" name:"CacheKey"`

	// 缓存规则配置。
	Cache *Cache `json:"Cache,omitempty" name:"Cache"`

	// Https配置。
	Https *Https `json:"Https,omitempty" name:"Https"`

	// 访问协议强制跳转配置。
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitempty" name:"ForceRedirect"`

	// 域名绑定的标签
	Tag []*Tag `json:"Tag,omitempty" name:"Tag"`

	// WebSocket配置
	WebSocket *WebSocket `json:"WebSocket,omitempty" name:"WebSocket"`
}

func (r *AddEcdnDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddEcdnDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Origin")
	delete(f, "Area")
	delete(f, "ProjectId")
	delete(f, "IpFilter")
	delete(f, "IpFreqLimit")
	delete(f, "ResponseHeader")
	delete(f, "CacheKey")
	delete(f, "Cache")
	delete(f, "Https")
	delete(f, "ForceRedirect")
	delete(f, "Tag")
	delete(f, "WebSocket")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddEcdnDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddEcdnDomainResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddEcdnDomainResponse struct {
	*tchttp.BaseResponse
	Response *AddEcdnDomainResponseParams `json:"Response"`
}

func (r *AddEcdnDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddEcdnDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AdvanceHttps struct {
	// 自定义Tls数据开关
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomTlsStatus *string `json:"CustomTlsStatus,omitempty" name:"CustomTlsStatus"`

	// Tls版本列表，支持设置 TLSv1, TLSV1.1, TLSV1.2, TLSv1.3，修改时必须开启连续的版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	TlsVersion []*string `json:"TlsVersion,omitempty" name:"TlsVersion"`

	// 自定义加密套件
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cipher *string `json:"Cipher,omitempty" name:"Cipher"`

	// 回源双向校验开启状态
	// off - 关闭校验
	// oneWay - 校验源站
	// twoWay - 双向校验
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerifyOriginType *string `json:"VerifyOriginType,omitempty" name:"VerifyOriginType"`

	// 回源层证书配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertInfo *ServerCert `json:"CertInfo,omitempty" name:"CertInfo"`

	// 源站证书配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginCertInfo *ClientCert `json:"OriginCertInfo,omitempty" name:"OriginCertInfo"`
}

type Cache struct {
	// 缓存配置规则数组。
	CacheRules []*CacheRule `json:"CacheRules,omitempty" name:"CacheRules"`

	// 遵循源站 Cache-Control: max-age 配置，白名单功能。
	// on：开启
	// off：关闭
	// 开启后，未能匹配 CacheRules 规则的资源将根据源站返回的 max-age 值进行节点缓存；匹配了 CacheRules 规则的资源将按照 CacheRules 中设置的缓存过期时间在节点进行缓存
	// 注意：此字段可能返回 null，表示取不到有效值。
	FollowOrigin *string `json:"FollowOrigin,omitempty" name:"FollowOrigin"`
}

type CacheKey struct {
	// 是否开启全路径缓存，on或off。
	FullUrlCache *string `json:"FullUrlCache,omitempty" name:"FullUrlCache"`
}

type CacheRule struct {
	// 缓存类型，支持all，file，directory，path，index，分别表示全部文件，后缀类型，目录，完整路径，首页。
	CacheType *string `json:"CacheType,omitempty" name:"CacheType"`

	// 缓存内容列表。
	CacheContents []*string `json:"CacheContents,omitempty" name:"CacheContents"`

	// 缓存时间，单位秒。
	CacheTime *int64 `json:"CacheTime,omitempty" name:"CacheTime"`
}

type ClientCert struct {
	// 客户端证书，pem格式。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Certificate *string `json:"Certificate,omitempty" name:"Certificate"`

	// 客户端证书名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// 证书过期时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 证书颁发时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployTime *string `json:"DeployTime,omitempty" name:"DeployTime"`
}

// Predefined struct for user
type CreateVerifyRecordRequestParams struct {
	// 要取回的域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

type CreateVerifyRecordRequest struct {
	*tchttp.BaseRequest
	
	// 要取回的域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *CreateVerifyRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVerifyRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVerifyRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVerifyRecordResponseParams struct {
	// 子解析
	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`

	// 解析值
	Record *string `json:"Record,omitempty" name:"Record"`

	// 解析类型
	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateVerifyRecordResponse struct {
	*tchttp.BaseResponse
	Response *CreateVerifyRecordResponseParams `json:"Response"`
}

func (r *CreateVerifyRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVerifyRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEcdnDomainRequestParams struct {
	// 待删除域名。
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

type DeleteEcdnDomainRequest struct {
	*tchttp.BaseRequest
	
	// 待删除域名。
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *DeleteEcdnDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEcdnDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteEcdnDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEcdnDomainResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteEcdnDomainResponse struct {
	*tchttp.BaseResponse
	Response *DeleteEcdnDomainResponseParams `json:"Response"`
}

func (r *DeleteEcdnDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEcdnDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainsConfigRequestParams struct {
	// 分页查询的偏移地址，默认0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询的域名个数，默认100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 查询条件过滤器。
	Filters []*DomainFilter `json:"Filters,omitempty" name:"Filters"`

	// 查询结果排序规则。
	Sort *Sort `json:"Sort,omitempty" name:"Sort"`
}

type DescribeDomainsConfigRequest struct {
	*tchttp.BaseRequest
	
	// 分页查询的偏移地址，默认0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询的域名个数，默认100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 查询条件过滤器。
	Filters []*DomainFilter `json:"Filters,omitempty" name:"Filters"`

	// 查询结果排序规则。
	Sort *Sort `json:"Sort,omitempty" name:"Sort"`
}

func (r *DescribeDomainsConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainsConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "Sort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainsConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainsConfigResponseParams struct {
	// 域名列表。
	Domains []*DomainDetailInfo `json:"Domains,omitempty" name:"Domains"`

	// 符合查询条件的域名总数，用于分页查询。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDomainsConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainsConfigResponseParams `json:"Response"`
}

func (r *DescribeDomainsConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainsConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainsRequestParams struct {
	// 分页查询的偏移地址，默认0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询的域名个数，默认100，最大支持1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 查询条件过滤器。
	Filters []*DomainFilter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeDomainsRequest struct {
	*tchttp.BaseRequest
	
	// 分页查询的偏移地址，默认0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询的域名个数，默认100，最大支持1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 查询条件过滤器。
	Filters []*DomainFilter `json:"Filters,omitempty" name:"Filters"`
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
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainsResponseParams struct {
	// 域名信息列表。
	Domains []*DomainBriefInfo `json:"Domains,omitempty" name:"Domains"`

	// 域名总个数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

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
type DescribeEcdnDomainLogsRequestParams struct {
	// 待查询域名。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 日志起始时间。如：2019-10-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 日志结束时间，只支持最近30天内日志查询。2019-10-02 00:00:00
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 日志链接列表分页起始地址，默认0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 日志链接列表分页记录条数，默认100，最大1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeEcdnDomainLogsRequest struct {
	*tchttp.BaseRequest
	
	// 待查询域名。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 日志起始时间。如：2019-10-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 日志结束时间，只支持最近30天内日志查询。2019-10-02 00:00:00
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 日志链接列表分页起始地址，默认0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 日志链接列表分页记录条数，默认100，最大1000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeEcdnDomainLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEcdnDomainLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEcdnDomainLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEcdnDomainLogsResponseParams struct {
	// 日志链接列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainLogs []*DomainLogs `json:"DomainLogs,omitempty" name:"DomainLogs"`

	// 日志链接总条数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEcdnDomainLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEcdnDomainLogsResponseParams `json:"Response"`
}

func (r *DescribeEcdnDomainLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEcdnDomainLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEcdnDomainStatisticsRequestParams struct {
	// 查询起始时间，如：2019-12-13 00:00:00。
	// 起止时间不超过90天。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，如：2019-12-13 23:59:59。
	// 起止时间不超过90天。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 统计指标名称:
	// flux：流量，单位为 byte
	// bandwidth：带宽，单位为 bps
	// request：请求数，单位为 次
	Metrics []*string `json:"Metrics,omitempty" name:"Metrics"`

	// 指定查询域名列表
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 指定要查询的项目 ID，[前往查看项目 ID](https://console.cloud.tencent.com/project)
	// 未填充域名情况下，指定项目查询，若填充了具体域名信息，以域名为主
	Projects []*int64 `json:"Projects,omitempty" name:"Projects"`

	// 列表分页起始地址，默认0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 列表分页记录条数，默认1000，最大3000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 统计区域:
	// mainland: 境内
	// oversea: 境外
	// global: 全部
	// 默认 global
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeEcdnDomainStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 查询起始时间，如：2019-12-13 00:00:00。
	// 起止时间不超过90天。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，如：2019-12-13 23:59:59。
	// 起止时间不超过90天。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 统计指标名称:
	// flux：流量，单位为 byte
	// bandwidth：带宽，单位为 bps
	// request：请求数，单位为 次
	Metrics []*string `json:"Metrics,omitempty" name:"Metrics"`

	// 指定查询域名列表
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 指定要查询的项目 ID，[前往查看项目 ID](https://console.cloud.tencent.com/project)
	// 未填充域名情况下，指定项目查询，若填充了具体域名信息，以域名为主
	Projects []*int64 `json:"Projects,omitempty" name:"Projects"`

	// 列表分页起始地址，默认0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 列表分页记录条数，默认1000，最大3000。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 统计区域:
	// mainland: 境内
	// oversea: 境外
	// global: 全部
	// 默认 global
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeEcdnDomainStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEcdnDomainStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Metrics")
	delete(f, "Domains")
	delete(f, "Projects")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEcdnDomainStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEcdnDomainStatisticsResponseParams struct {
	// 域名数据
	Data []*DomainData `json:"Data,omitempty" name:"Data"`

	// 数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEcdnDomainStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEcdnDomainStatisticsResponseParams `json:"Response"`
}

func (r *DescribeEcdnDomainStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEcdnDomainStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEcdnStatisticsRequestParams struct {
	// 查询起始时间，如：2019-12-13 00:00:00
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，如：2019-12-13 23:59:59
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 指定查询指标，支持的类型有：
	// flux：流量，单位为 byte
	// bandwidth：带宽，单位为 bps
	// request：请求数，单位为 次
	// 2xx：返回 2xx 状态码汇总或者 2 开头状态码数据，单位为 个
	// 3xx：返回 3xx 状态码汇总或者 3 开头状态码数据，单位为 个
	// 4xx：返回 4xx 状态码汇总或者 4 开头状态码数据，单位为 个
	// 5xx：返回 5xx 状态码汇总或者 5 开头状态码数据，单位为 个
	Metrics []*string `json:"Metrics,omitempty" name:"Metrics"`

	// 时间粒度，支持以下几种模式：
	// 1 天	 1，5，15，30，60，120，240，1440 
	// 2 ~ 3 天	15，30，60，120，240，1440
	// 4 ~ 7 天	30，60，120，240，1440
	// 8 ~ 31 天	 60，120，240，1440
	Interval *int64 `json:"Interval,omitempty" name:"Interval"`

	// 指定查询域名列表
	// 
	// 最多可一次性查询30个加速域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 指定要查询的项目 ID，[前往查看项目 ID](https://console.cloud.tencent.com/project)
	// 未填充域名情况下，指定项目查询，若填充了具体域名信息，以域名为主
	Projects []*int64 `json:"Projects,omitempty" name:"Projects"`

	// 统计区域:
	// mainland: 境内
	// oversea: 境外
	// global: 全部
	// 默认 global
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeEcdnStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 查询起始时间，如：2019-12-13 00:00:00
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，如：2019-12-13 23:59:59
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 指定查询指标，支持的类型有：
	// flux：流量，单位为 byte
	// bandwidth：带宽，单位为 bps
	// request：请求数，单位为 次
	// 2xx：返回 2xx 状态码汇总或者 2 开头状态码数据，单位为 个
	// 3xx：返回 3xx 状态码汇总或者 3 开头状态码数据，单位为 个
	// 4xx：返回 4xx 状态码汇总或者 4 开头状态码数据，单位为 个
	// 5xx：返回 5xx 状态码汇总或者 5 开头状态码数据，单位为 个
	Metrics []*string `json:"Metrics,omitempty" name:"Metrics"`

	// 时间粒度，支持以下几种模式：
	// 1 天	 1，5，15，30，60，120，240，1440 
	// 2 ~ 3 天	15，30，60，120，240，1440
	// 4 ~ 7 天	30，60，120，240，1440
	// 8 ~ 31 天	 60，120，240，1440
	Interval *int64 `json:"Interval,omitempty" name:"Interval"`

	// 指定查询域名列表
	// 
	// 最多可一次性查询30个加速域名。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 指定要查询的项目 ID，[前往查看项目 ID](https://console.cloud.tencent.com/project)
	// 未填充域名情况下，指定项目查询，若填充了具体域名信息，以域名为主
	Projects []*int64 `json:"Projects,omitempty" name:"Projects"`

	// 统计区域:
	// mainland: 境内
	// oversea: 境外
	// global: 全部
	// 默认 global
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeEcdnStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEcdnStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Metrics")
	delete(f, "Interval")
	delete(f, "Domains")
	delete(f, "Projects")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEcdnStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEcdnStatisticsResponseParams struct {
	// 指定条件查询得到的数据明细
	Data []*ResourceData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEcdnStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEcdnStatisticsResponseParams `json:"Response"`
}

func (r *DescribeEcdnStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEcdnStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIpStatusRequestParams struct {
	// 加速域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 查询区域：
	// mainland: 国内节点
	// overseas: 海外节点
	// global: 全球节点
	Area *string `json:"Area,omitempty" name:"Area"`
}

type DescribeIpStatusRequest struct {
	*tchttp.BaseRequest
	
	// 加速域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 查询区域：
	// mainland: 国内节点
	// overseas: 海外节点
	// global: 全球节点
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeIpStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIpStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIpStatusResponseParams struct {
	// 节点列表
	Ips []*IpStatus `json:"Ips,omitempty" name:"Ips"`

	// 节点总个数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIpStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIpStatusResponseParams `json:"Response"`
}

func (r *DescribeIpStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePurgeQuotaRequestParams struct {

}

type DescribePurgeQuotaRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribePurgeQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePurgeQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePurgeQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePurgeQuotaResponseParams struct {
	// Url刷新用量及配额。
	UrlPurge *Quota `json:"UrlPurge,omitempty" name:"UrlPurge"`

	// 目录刷新用量及配额。
	PathPurge *Quota `json:"PathPurge,omitempty" name:"PathPurge"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePurgeQuotaResponse struct {
	*tchttp.BaseResponse
	Response *DescribePurgeQuotaResponseParams `json:"Response"`
}

func (r *DescribePurgeQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePurgeQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePurgeTasksRequestParams struct {
	// 查询刷新类型。url：查询 url 刷新记录；path：查询目录刷新记录。
	PurgeType *string `json:"PurgeType,omitempty" name:"PurgeType"`

	// 开始时间，如2018-08-08 00:00:00。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，如2018-08-08 23:59:59。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 提交时返回的任务 Id，查询时 TaskId 和起始时间必须指定一项。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 分页查询偏移量，默认为0（从第0条开始）。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询限制数目，默认为20。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 查询关键字，请输入域名或 http(s):// 开头完整 URL。
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// 查询指定任务状态，fail表示失败，done表示成功，process表示刷新中。
	Status *string `json:"Status,omitempty" name:"Status"`
}

type DescribePurgeTasksRequest struct {
	*tchttp.BaseRequest
	
	// 查询刷新类型。url：查询 url 刷新记录；path：查询目录刷新记录。
	PurgeType *string `json:"PurgeType,omitempty" name:"PurgeType"`

	// 开始时间，如2018-08-08 00:00:00。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，如2018-08-08 23:59:59。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 提交时返回的任务 Id，查询时 TaskId 和起始时间必须指定一项。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 分页查询偏移量，默认为0（从第0条开始）。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询限制数目，默认为20。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 查询关键字，请输入域名或 http(s):// 开头完整 URL。
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// 查询指定任务状态，fail表示失败，done表示成功，process表示刷新中。
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *DescribePurgeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePurgeTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PurgeType")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "TaskId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Keyword")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePurgeTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePurgeTasksResponseParams struct {
	// 刷新历史记录。
	PurgeLogs []*PurgeTask `json:"PurgeLogs,omitempty" name:"PurgeLogs"`

	// 任务总数，用于分页。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePurgeTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribePurgeTasksResponseParams `json:"Response"`
}

func (r *DescribePurgeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePurgeTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetailData struct {
	// 数据类型的名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 数据值
	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type DomainBriefInfo struct {
	// 域名ID。
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 腾讯云账号ID。
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// CDN加速域名。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名CName。
	Cname *string `json:"Cname,omitempty" name:"Cname"`

	// 域名状态，pending，rejected，processing， online，offline，deleted分别表示审核中，审核未通过，审核通过部署中，已开启，已关闭，已删除。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 项目ID。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 域名创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 域名更新时间。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 源站配置详情。
	Origin *Origin `json:"Origin,omitempty" name:"Origin"`

	// 域名封禁状态，normal，overdue，quota，malicious，ddos，idle，unlicensed，capping，readonly分别表示 正常，欠费停服，试用客户流量包耗尽，恶意用户，ddos攻击，无流量域名，未备案，带宽封顶，只读
	Disable *string `json:"Disable,omitempty" name:"Disable"`

	// 加速区域，mainland，oversea或global。
	Area *string `json:"Area,omitempty" name:"Area"`

	// 域名锁定状态，normal、global，分别表示未被锁定、全球锁定。
	Readonly *string `json:"Readonly,omitempty" name:"Readonly"`

	// 域名标签。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag []*Tag `json:"Tag,omitempty" name:"Tag"`
}

type DomainData struct {
	// 域名
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// 结果详情
	DetailData []*DetailData `json:"DetailData,omitempty" name:"DetailData"`
}

type DomainDetailInfo struct {
	// 域名ID。
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 腾讯云账号ID。
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 加速域名。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 域名CName。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cname *string `json:"Cname,omitempty" name:"Cname"`

	// 域名状态，pending，rejected，processing， online，offline，deleted分别表示审核中，审核未通过，审核通过部署中，已开启，已关闭，已删除。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 项目ID。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 域名创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 域名更新时间。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 源站配置。
	Origin *Origin `json:"Origin,omitempty" name:"Origin"`

	// IP黑白名单配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpFilter *IpFilter `json:"IpFilter,omitempty" name:"IpFilter"`

	// IP限频配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpFreqLimit *IpFreqLimit `json:"IpFreqLimit,omitempty" name:"IpFreqLimit"`

	// 源站响应头部配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResponseHeader *ResponseHeader `json:"ResponseHeader,omitempty" name:"ResponseHeader"`

	// 节点缓存配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheKey *CacheKey `json:"CacheKey,omitempty" name:"CacheKey"`

	// 缓存规则配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cache *Cache `json:"Cache,omitempty" name:"Cache"`

	// Https配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Https *Https `json:"Https,omitempty" name:"Https"`

	// 域名封禁状态，normal，overdue，quota，malicious，ddos，idle，unlicensed，capping，readonly分别表示 正常，欠费停服，试用客户流量包耗尽，恶意用户，ddos攻击，无流量域名，未备案，带宽封顶，只读。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Disable *string `json:"Disable,omitempty" name:"Disable"`

	// 访问协议强制跳转配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitempty" name:"ForceRedirect"`

	// 加速区域，mainland，overseas或global。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Area *string `json:"Area,omitempty" name:"Area"`

	// 域名锁定状态，normal、global 分别表示未被锁定，全球锁定。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Readonly *string `json:"Readonly,omitempty" name:"Readonly"`

	// 域名标签。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag []*Tag `json:"Tag,omitempty" name:"Tag"`

	// WebSocket配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebSocket *WebSocket `json:"WebSocket,omitempty" name:"WebSocket"`
}

type DomainFilter struct {
	// 过滤字段名，支持的列表如下：
	// - origin：主源站。
	// - domain：域名。
	// - resourceId：域名id。
	// - status：域名状态，online，offline，processing。
	// - disable：域名封禁状态，normal，unlicensed。
	// - projectId：项目ID。
	// - fullUrlCache：全路径缓存，on或off。
	// - https：是否配置https，on，off或processing。
	// - originPullProtocol：回源协议类型，支持http，follow或https。
	// - area：加速区域，支持mainland，overseas或global。
	// - tagKey：标签键。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 过滤字段值。
	Value []*string `json:"Value,omitempty" name:"Value"`

	// 是否启用模糊查询，仅支持过滤字段名为origin，domain。
	Fuzzy *bool `json:"Fuzzy,omitempty" name:"Fuzzy"`
}

type DomainLogs struct {
	// 日志起始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 日志结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 日志下载路径。
	LogPath *string `json:"LogPath,omitempty" name:"LogPath"`
}

type EcdnData struct {
	// 查询指定的指标名称：Bandwidth，Flux，Request，Delay，状态码，LogBandwidth，LogFlux，LogRequest
	Metrics []*string `json:"Metrics,omitempty" name:"Metrics"`

	// 明细数据组合
	DetailData []*TimestampData `json:"DetailData,omitempty" name:"DetailData"`
}

type ForceRedirect struct {
	// 访问协议强制跳转配置开关，on或off。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 强制跳转访问协议类型，支持http，https，分别表示请求强制跳转http协议，请求强制跳转https协议。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RedirectType *string `json:"RedirectType,omitempty" name:"RedirectType"`

	// 强制跳转开启时返回的http状态码，支持301或302。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RedirectStatusCode *int64 `json:"RedirectStatusCode,omitempty" name:"RedirectStatusCode"`
}

type Hsts struct {
	// 是否开启，on或off。
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// MaxAge数值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxAge *int64 `json:"MaxAge,omitempty" name:"MaxAge"`

	// 是否包含子域名，on或off。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncludeSubDomains *string `json:"IncludeSubDomains,omitempty" name:"IncludeSubDomains"`
}

type HttpHeaderPathRule struct {
	// http头部设置方式，支持add，set或del，分别表示新增，设置或删除头部。
	// 请求头部暂不支持set。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeaderMode *string `json:"HeaderMode,omitempty" name:"HeaderMode"`

	// http头部名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeaderName *string `json:"HeaderName,omitempty" name:"HeaderName"`

	// http头部值。del时可不填写该字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeaderValue *string `json:"HeaderValue,omitempty" name:"HeaderValue"`

	// 生效的url路径规则类型，支持all，file，directory或path，分别表示全部路径，文件后缀类型，目录或绝对路径生效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`

	// url路径或文件类型列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RulePaths []*string `json:"RulePaths,omitempty" name:"RulePaths"`
}

type Https struct {
	// https配置开关，on或off。开启https配置的域名在部署中状态，开关保持off。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 是否开启http2，on或off。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Http2 *string `json:"Http2,omitempty" name:"Http2"`

	// 是否开启OCSP功能，on或off。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcspStapling *string `json:"OcspStapling,omitempty" name:"OcspStapling"`

	// 是否开启客户端证书校验功能，on或off，开启时必选上传客户端证书信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VerifyClient *string `json:"VerifyClient,omitempty" name:"VerifyClient"`

	// 服务器证书配置信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertInfo *ServerCert `json:"CertInfo,omitempty" name:"CertInfo"`

	// 客户端证书配置信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientCertInfo *ClientCert `json:"ClientCertInfo,omitempty" name:"ClientCertInfo"`

	// 是否开启Spdy，on或off。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Spdy *string `json:"Spdy,omitempty" name:"Spdy"`

	// https证书部署状态，closed，deploying，deployed，failed分别表示已关闭，部署中，部署成功，部署失败。不可作为入参使用。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SslStatus *string `json:"SslStatus,omitempty" name:"SslStatus"`

	// Hsts配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Hsts *Hsts `json:"Hsts,omitempty" name:"Hsts"`
}

type IpFilter struct {
	// IP黑白名单开关，on或off。
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// IP黑白名单类型，whitelist或blacklist。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilterType *string `json:"FilterType,omitempty" name:"FilterType"`

	// IP黑白名单列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Filters []*string `json:"Filters,omitempty" name:"Filters"`
}

type IpFreqLimit struct {
	// IP限频配置开关，on或off。
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 每秒请求数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Qps *int64 `json:"Qps,omitempty" name:"Qps"`
}

type IpStatus struct {
	// 节点 IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 节点所属区域
	District *string `json:"District,omitempty" name:"District"`

	// 节点所属运营商
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// 节点所在城市
	City *string `json:"City,omitempty" name:"City"`

	// 节点状态
	// online：上线状态，正常调度服务中
	// offline：下线状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 节点 IP 添加时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type Origin struct {
	// 主源站列表，IP与域名源站不可混填。配置源站端口["origin1:port1", "origin2:port2"]，配置回源权重["origin1::weight1", "origin2::weight2"]，同时配置端口与权重 ["origin1:port1:weight1", "origin2:port2:weight2"]，权重值有效范围为0-100。
	Origins []*string `json:"Origins,omitempty" name:"Origins"`

	// 主源站类型，支持domain，ip，分别表示域名源站，ip源站。
	// 设置Origins时必须填写。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginType *string `json:"OriginType,omitempty" name:"OriginType"`

	// 回源时Host头部值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`

	// 回源协议类型，支持http，follow，https，分别表示强制http回源，协议跟随回源，https回源。
	// 不传入的情况下默认为http回源.
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginPullProtocol *string `json:"OriginPullProtocol,omitempty" name:"OriginPullProtocol"`

	// 备份源站列表。
	BackupOrigins []*string `json:"BackupOrigins,omitempty" name:"BackupOrigins"`

	// 备份源站类型，同OriginType。
	// 设置BackupOrigins时必须填写。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupOriginType *string `json:"BackupOriginType,omitempty" name:"BackupOriginType"`

	// HTTPS回源高级配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdvanceHttps *AdvanceHttps `json:"AdvanceHttps,omitempty" name:"AdvanceHttps"`
}

// Predefined struct for user
type PurgePathCacheRequestParams struct {
	// 要刷新的目录列表，必须包含协议头部。
	Paths []*string `json:"Paths,omitempty" name:"Paths"`

	// 刷新类型，flush 代表刷新有更新的资源，delete 表示刷新全部资源。
	FlushType *string `json:"FlushType,omitempty" name:"FlushType"`
}

type PurgePathCacheRequest struct {
	*tchttp.BaseRequest
	
	// 要刷新的目录列表，必须包含协议头部。
	Paths []*string `json:"Paths,omitempty" name:"Paths"`

	// 刷新类型，flush 代表刷新有更新的资源，delete 表示刷新全部资源。
	FlushType *string `json:"FlushType,omitempty" name:"FlushType"`
}

func (r *PurgePathCacheRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PurgePathCacheRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Paths")
	delete(f, "FlushType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PurgePathCacheRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PurgePathCacheResponseParams struct {
	// 刷新任务Id。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type PurgePathCacheResponse struct {
	*tchttp.BaseResponse
	Response *PurgePathCacheResponseParams `json:"Response"`
}

func (r *PurgePathCacheResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PurgePathCacheResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PurgeTask struct {
	// 刷新任务ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 刷新Url。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 刷新任务状态，fail表示失败，done表示成功，process表示刷新中。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 刷新类型，url表示url刷新，path表示目录刷新。
	PurgeType *string `json:"PurgeType,omitempty" name:"PurgeType"`

	// 刷新资源方式，flush代表刷新更新资源，delete代表刷新全部资源。
	FlushType *string `json:"FlushType,omitempty" name:"FlushType"`

	// 刷新任务提交时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

// Predefined struct for user
type PurgeUrlsCacheRequestParams struct {
	// 要刷新的Url列表，必须包含协议头部。
	Urls []*string `json:"Urls,omitempty" name:"Urls"`
}

type PurgeUrlsCacheRequest struct {
	*tchttp.BaseRequest
	
	// 要刷新的Url列表，必须包含协议头部。
	Urls []*string `json:"Urls,omitempty" name:"Urls"`
}

func (r *PurgeUrlsCacheRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PurgeUrlsCacheRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Urls")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PurgeUrlsCacheRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PurgeUrlsCacheResponseParams struct {
	// 刷新任务Id。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type PurgeUrlsCacheResponse struct {
	*tchttp.BaseResponse
	Response *PurgeUrlsCacheResponseParams `json:"Response"`
}

func (r *PurgeUrlsCacheResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PurgeUrlsCacheResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Quota struct {
	// 单次批量提交配额上限。
	Batch *int64 `json:"Batch,omitempty" name:"Batch"`

	// 每日提交配额上限。
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 每日剩余的可提交配额。
	Available *int64 `json:"Available,omitempty" name:"Available"`
}

type ResourceData struct {
	// 资源名称，根据查询条件不同分为以下几类：
	// 具体域名：表示该域名明细数据
	// multiDomains：表示多域名汇总明细数据
	// 项目 ID：指定项目查询时，显示为项目 ID
	// all：账号维度明细数据
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// 资源对应的数据明细
	EcdnData *EcdnData `json:"EcdnData,omitempty" name:"EcdnData"`
}

type ResponseHeader struct {
	// 自定义响应头开关，on或off。
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 自定义响应头规则数组。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeaderRules []*HttpHeaderPathRule `json:"HeaderRules,omitempty" name:"HeaderRules"`
}

type ServerCert struct {
	// 服务器证书id，当证书为腾讯云托管证书时必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// 服务器证书名称，当证书为腾讯云托管证书时必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// 服务器证书信息，上传自有证书时必填，必须包含完整的证书链信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Certificate *string `json:"Certificate,omitempty" name:"Certificate"`

	// 服务器密钥信息，上传自有证书时必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`

	// 证书过期时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 证书颁发时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployTime *string `json:"DeployTime,omitempty" name:"DeployTime"`

	// 证书备注信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`
}

type Sort struct {
	// 排序字段，当前支持：
	// createTime，域名创建时间
	// certExpireTime，证书过期时间
	Key *string `json:"Key,omitempty" name:"Key"`

	// asc/desc，默认desc。
	Sequence *string `json:"Sequence,omitempty" name:"Sequence"`
}

// Predefined struct for user
type StartEcdnDomainRequestParams struct {
	// 待启用域名。
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

type StartEcdnDomainRequest struct {
	*tchttp.BaseRequest
	
	// 待启用域名。
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *StartEcdnDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartEcdnDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartEcdnDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartEcdnDomainResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StartEcdnDomainResponse struct {
	*tchttp.BaseResponse
	Response *StartEcdnDomainResponseParams `json:"Response"`
}

func (r *StartEcdnDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartEcdnDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopEcdnDomainRequestParams struct {
	// 待停用域名。
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

type StopEcdnDomainRequest struct {
	*tchttp.BaseRequest
	
	// 待停用域名。
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

func (r *StopEcdnDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopEcdnDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopEcdnDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopEcdnDomainResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StopEcdnDomainResponse struct {
	*tchttp.BaseResponse
	Response *StopEcdnDomainResponseParams `json:"Response"`
}

func (r *StopEcdnDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopEcdnDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {
	// 标签键
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type TimestampData struct {
	// 数据统计时间点，采用向前汇总模式
	// 以 5 分钟粒度为例，13:35:00 时间点代表的统计数据区间为 13:35:00 至 13:39:59
	Time *string `json:"Time,omitempty" name:"Time"`

	// 数据值
	Value []*float64 `json:"Value,omitempty" name:"Value"`
}

// Predefined struct for user
type UpdateDomainConfigRequestParams struct {
	// 域名。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 源站配置。
	Origin *Origin `json:"Origin,omitempty" name:"Origin"`

	// 项目id。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// IP黑白名单配置。
	IpFilter *IpFilter `json:"IpFilter,omitempty" name:"IpFilter"`

	// IP限频配置。
	IpFreqLimit *IpFreqLimit `json:"IpFreqLimit,omitempty" name:"IpFreqLimit"`

	// 源站响应头部配置。
	ResponseHeader *ResponseHeader `json:"ResponseHeader,omitempty" name:"ResponseHeader"`

	// 节点缓存配置。
	CacheKey *CacheKey `json:"CacheKey,omitempty" name:"CacheKey"`

	// 缓存规则配置。
	Cache *Cache `json:"Cache,omitempty" name:"Cache"`

	// Https配置。
	Https *Https `json:"Https,omitempty" name:"Https"`

	// 访问协议强制跳转配置。
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitempty" name:"ForceRedirect"`

	// 域名加速区域，mainland，overseas或global，分别表示中国境内加速，海外加速或全球加速。
	Area *string `json:"Area,omitempty" name:"Area"`

	// WebSocket配置
	WebSocket *WebSocket `json:"WebSocket,omitempty" name:"WebSocket"`
}

type UpdateDomainConfigRequest struct {
	*tchttp.BaseRequest
	
	// 域名。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 源站配置。
	Origin *Origin `json:"Origin,omitempty" name:"Origin"`

	// 项目id。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// IP黑白名单配置。
	IpFilter *IpFilter `json:"IpFilter,omitempty" name:"IpFilter"`

	// IP限频配置。
	IpFreqLimit *IpFreqLimit `json:"IpFreqLimit,omitempty" name:"IpFreqLimit"`

	// 源站响应头部配置。
	ResponseHeader *ResponseHeader `json:"ResponseHeader,omitempty" name:"ResponseHeader"`

	// 节点缓存配置。
	CacheKey *CacheKey `json:"CacheKey,omitempty" name:"CacheKey"`

	// 缓存规则配置。
	Cache *Cache `json:"Cache,omitempty" name:"Cache"`

	// Https配置。
	Https *Https `json:"Https,omitempty" name:"Https"`

	// 访问协议强制跳转配置。
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitempty" name:"ForceRedirect"`

	// 域名加速区域，mainland，overseas或global，分别表示中国境内加速，海外加速或全球加速。
	Area *string `json:"Area,omitempty" name:"Area"`

	// WebSocket配置
	WebSocket *WebSocket `json:"WebSocket,omitempty" name:"WebSocket"`
}

func (r *UpdateDomainConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDomainConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Origin")
	delete(f, "ProjectId")
	delete(f, "IpFilter")
	delete(f, "IpFreqLimit")
	delete(f, "ResponseHeader")
	delete(f, "CacheKey")
	delete(f, "Cache")
	delete(f, "Https")
	delete(f, "ForceRedirect")
	delete(f, "Area")
	delete(f, "WebSocket")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateDomainConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDomainConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateDomainConfigResponse struct {
	*tchttp.BaseResponse
	Response *UpdateDomainConfigResponseParams `json:"Response"`
}

func (r *UpdateDomainConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDomainConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WebSocket struct {
	// WebSocket 超时配置开关, 开关为off时，平台仍支持WebSocket连接，此时超时时间默认为15秒，若需要调整超时时间，将开关置为on.
	// 
	// * WebSocket 为内测功能,如需使用,请联系腾讯云工程师开白.
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 设置超时时间，单位为秒，最大超时时间65秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`
}