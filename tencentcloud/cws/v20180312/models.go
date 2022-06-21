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

package v20180312

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type CreateMonitorsRequestParams struct {
	// 站点的url列表
	Urls []*string `json:"Urls,omitempty" name:"Urls"`

	// 任务名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 扫描模式，normal-正常扫描；deep-深度扫描
	ScannerType *string `json:"ScannerType,omitempty" name:"ScannerType"`

	// 扫描周期，单位小时，每X小时执行一次
	Crontab *uint64 `json:"Crontab,omitempty" name:"Crontab"`

	// 扫描速率限制，每秒发送X个HTTP请求
	RateLimit *uint64 `json:"RateLimit,omitempty" name:"RateLimit"`

	// 首次扫描开始时间
	FirstScanStartTime *string `json:"FirstScanStartTime,omitempty" name:"FirstScanStartTime"`
}

type CreateMonitorsRequest struct {
	*tchttp.BaseRequest
	
	// 站点的url列表
	Urls []*string `json:"Urls,omitempty" name:"Urls"`

	// 任务名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 扫描模式，normal-正常扫描；deep-深度扫描
	ScannerType *string `json:"ScannerType,omitempty" name:"ScannerType"`

	// 扫描周期，单位小时，每X小时执行一次
	Crontab *uint64 `json:"Crontab,omitempty" name:"Crontab"`

	// 扫描速率限制，每秒发送X个HTTP请求
	RateLimit *uint64 `json:"RateLimit,omitempty" name:"RateLimit"`

	// 首次扫描开始时间
	FirstScanStartTime *string `json:"FirstScanStartTime,omitempty" name:"FirstScanStartTime"`
}

func (r *CreateMonitorsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMonitorsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Urls")
	delete(f, "Name")
	delete(f, "ScannerType")
	delete(f, "Crontab")
	delete(f, "RateLimit")
	delete(f, "FirstScanStartTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMonitorsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMonitorsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateMonitorsResponse struct {
	*tchttp.BaseResponse
	Response *CreateMonitorsResponseParams `json:"Response"`
}

func (r *CreateMonitorsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMonitorsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSitesRequestParams struct {
	// 站点的url列表
	Urls []*string `json:"Urls,omitempty" name:"Urls"`

	// 访问网站的客户端标识
	UserAgent *string `json:"UserAgent,omitempty" name:"UserAgent"`
}

type CreateSitesRequest struct {
	*tchttp.BaseRequest
	
	// 站点的url列表
	Urls []*string `json:"Urls,omitempty" name:"Urls"`

	// 访问网站的客户端标识
	UserAgent *string `json:"UserAgent,omitempty" name:"UserAgent"`
}

func (r *CreateSitesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSitesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Urls")
	delete(f, "UserAgent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSitesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSitesResponseParams struct {
	// 新增站点数。
	Number *uint64 `json:"Number,omitempty" name:"Number"`

	// 站点数组
	Sites []*MiniSite `json:"Sites,omitempty" name:"Sites"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSitesResponse struct {
	*tchttp.BaseResponse
	Response *CreateSitesResponseParams `json:"Response"`
}

func (r *CreateSitesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSitesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSitesScansRequestParams struct {
	// 站点的ID列表
	SiteIds []*uint64 `json:"SiteIds,omitempty" name:"SiteIds"`

	// 扫描模式，normal-正常扫描；deep-深度扫描
	ScannerType *string `json:"ScannerType,omitempty" name:"ScannerType"`

	// 扫描速率限制，每秒发送X个HTTP请求
	RateLimit *uint64 `json:"RateLimit,omitempty" name:"RateLimit"`
}

type CreateSitesScansRequest struct {
	*tchttp.BaseRequest
	
	// 站点的ID列表
	SiteIds []*uint64 `json:"SiteIds,omitempty" name:"SiteIds"`

	// 扫描模式，normal-正常扫描；deep-深度扫描
	ScannerType *string `json:"ScannerType,omitempty" name:"ScannerType"`

	// 扫描速率限制，每秒发送X个HTTP请求
	RateLimit *uint64 `json:"RateLimit,omitempty" name:"RateLimit"`
}

func (r *CreateSitesScansRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSitesScansRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SiteIds")
	delete(f, "ScannerType")
	delete(f, "RateLimit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSitesScansRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSitesScansResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSitesScansResponse struct {
	*tchttp.BaseResponse
	Response *CreateSitesScansResponseParams `json:"Response"`
}

func (r *CreateSitesScansResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSitesScansResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVulsMisinformationRequestParams struct {
	// 漏洞ID列表
	VulIds []*uint64 `json:"VulIds,omitempty" name:"VulIds"`
}

type CreateVulsMisinformationRequest struct {
	*tchttp.BaseRequest
	
	// 漏洞ID列表
	VulIds []*uint64 `json:"VulIds,omitempty" name:"VulIds"`
}

func (r *CreateVulsMisinformationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVulsMisinformationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VulIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVulsMisinformationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVulsMisinformationResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateVulsMisinformationResponse struct {
	*tchttp.BaseResponse
	Response *CreateVulsMisinformationResponseParams `json:"Response"`
}

func (r *CreateVulsMisinformationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVulsMisinformationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVulsReportRequestParams struct {
	// 站点ID
	SiteId *uint64 `json:"SiteId,omitempty" name:"SiteId"`

	// 监控任务ID
	MonitorId *uint64 `json:"MonitorId,omitempty" name:"MonitorId"`
}

type CreateVulsReportRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID
	SiteId *uint64 `json:"SiteId,omitempty" name:"SiteId"`

	// 监控任务ID
	MonitorId *uint64 `json:"MonitorId,omitempty" name:"MonitorId"`
}

func (r *CreateVulsReportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVulsReportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SiteId")
	delete(f, "MonitorId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVulsReportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVulsReportResponseParams struct {
	// 报告下载地址
	ReportFileUrl *string `json:"ReportFileUrl,omitempty" name:"ReportFileUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateVulsReportResponse struct {
	*tchttp.BaseResponse
	Response *CreateVulsReportResponseParams `json:"Response"`
}

func (r *CreateVulsReportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVulsReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMonitorsRequestParams struct {
	// 监控任务ID列表
	MonitorIds []*uint64 `json:"MonitorIds,omitempty" name:"MonitorIds"`
}

type DeleteMonitorsRequest struct {
	*tchttp.BaseRequest
	
	// 监控任务ID列表
	MonitorIds []*uint64 `json:"MonitorIds,omitempty" name:"MonitorIds"`
}

func (r *DeleteMonitorsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMonitorsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MonitorIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMonitorsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMonitorsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteMonitorsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMonitorsResponseParams `json:"Response"`
}

func (r *DeleteMonitorsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMonitorsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSitesRequestParams struct {
	// 站点ID列表
	SiteIds []*uint64 `json:"SiteIds,omitempty" name:"SiteIds"`
}

type DeleteSitesRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID列表
	SiteIds []*uint64 `json:"SiteIds,omitempty" name:"SiteIds"`
}

func (r *DeleteSitesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSitesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SiteIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSitesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSitesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteSitesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSitesResponseParams `json:"Response"`
}

func (r *DeleteSitesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSitesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigRequestParams struct {

}

type DescribeConfigRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigResponseParams struct {
	// 漏洞告警通知等级，4位分别代表：高危、中危、低危、提示。
	NoticeLevel *string `json:"NoticeLevel,omitempty" name:"NoticeLevel"`

	// 配置ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 记录创建时间。
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 记录更新新建。
	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`

	// 云用户appid。
	Appid *uint64 `json:"Appid,omitempty" name:"Appid"`

	// 内容检测通知等级-1:通知,0-不通知
	ContentLevel *uint64 `json:"ContentLevel,omitempty" name:"ContentLevel"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConfigResponseParams `json:"Response"`
}

func (r *DescribeConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMonitorsRequestParams struct {
	// 监控任务ID列表
	MonitorIds []*uint64 `json:"MonitorIds,omitempty" name:"MonitorIds"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeMonitorsRequest struct {
	*tchttp.BaseRequest
	
	// 监控任务ID列表
	MonitorIds []*uint64 `json:"MonitorIds,omitempty" name:"MonitorIds"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeMonitorsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMonitorsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MonitorIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMonitorsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMonitorsResponseParams struct {
	// 监控任务列表。
	Monitors []*MonitorsDetail `json:"Monitors,omitempty" name:"Monitors"`

	// 监控任务数量。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMonitorsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMonitorsResponseParams `json:"Response"`
}

func (r *DescribeMonitorsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMonitorsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSiteQuotaRequestParams struct {

}

type DescribeSiteQuotaRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeSiteQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSiteQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSiteQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSiteQuotaResponseParams struct {
	// 已购买的扫描次数。
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 已使用的扫描次数。
	Used *uint64 `json:"Used,omitempty" name:"Used"`

	// 剩余可用的扫描次数。
	Available *uint64 `json:"Available,omitempty" name:"Available"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSiteQuotaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSiteQuotaResponseParams `json:"Response"`
}

func (r *DescribeSiteQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSiteQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSitesRequestParams struct {
	// 站点ID列表
	SiteIds []*uint64 `json:"SiteIds,omitempty" name:"SiteIds"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeSitesRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID列表
	SiteIds []*uint64 `json:"SiteIds,omitempty" name:"SiteIds"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSitesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSitesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SiteIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSitesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSitesResponseParams struct {
	// 站点数量。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 站点信息列表。
	Sites []*Site `json:"Sites,omitempty" name:"Sites"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSitesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSitesResponseParams `json:"Response"`
}

func (r *DescribeSitesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSitesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSitesVerificationRequestParams struct {
	// 站点的url列表
	Urls []*string `json:"Urls,omitempty" name:"Urls"`
}

type DescribeSitesVerificationRequest struct {
	*tchttp.BaseRequest
	
	// 站点的url列表
	Urls []*string `json:"Urls,omitempty" name:"Urls"`
}

func (r *DescribeSitesVerificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSitesVerificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Urls")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSitesVerificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSitesVerificationResponseParams struct {
	// 验证信息数量。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 验证信息列表。
	SitesVerification []*SitesVerification `json:"SitesVerification,omitempty" name:"SitesVerification"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSitesVerificationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSitesVerificationResponseParams `json:"Response"`
}

func (r *DescribeSitesVerificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSitesVerificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulsNumberRequestParams struct {

}

type DescribeVulsNumberRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeVulsNumberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulsNumberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulsNumberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulsNumberResponseParams struct {
	// 受影响的网站总数。
	ImpactSiteNumber *uint64 `json:"ImpactSiteNumber,omitempty" name:"ImpactSiteNumber"`

	// 已验证的网站总数。
	SiteNumber *uint64 `json:"SiteNumber,omitempty" name:"SiteNumber"`

	// 高风险漏洞总数。
	VulsHighNumber *uint64 `json:"VulsHighNumber,omitempty" name:"VulsHighNumber"`

	// 中风险漏洞总数。
	VulsMiddleNumber *uint64 `json:"VulsMiddleNumber,omitempty" name:"VulsMiddleNumber"`

	// 低高风险漏洞总数。
	VulsLowNumber *uint64 `json:"VulsLowNumber,omitempty" name:"VulsLowNumber"`

	// 风险提示总数。
	VulsNoticeNumber *uint64 `json:"VulsNoticeNumber,omitempty" name:"VulsNoticeNumber"`

	// 扫描页面总数。
	PageCount *uint64 `json:"PageCount,omitempty" name:"PageCount"`

	// 已验证的网站列表。
	Sites []*MonitorMiniSite `json:"Sites,omitempty" name:"Sites"`

	// 受影响的网站列表。
	ImpactSites []*MonitorMiniSite `json:"ImpactSites,omitempty" name:"ImpactSites"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVulsNumberResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulsNumberResponseParams `json:"Response"`
}

func (r *DescribeVulsNumberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulsNumberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulsNumberTimelineRequestParams struct {

}

type DescribeVulsNumberTimelineRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeVulsNumberTimelineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulsNumberTimelineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulsNumberTimelineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulsNumberTimelineResponseParams struct {
	// 统计数据记录数量。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 用户漏洞数随时间变化统计数据。
	VulsTimeline []*VulsTimeline `json:"VulsTimeline,omitempty" name:"VulsTimeline"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVulsNumberTimelineResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulsNumberTimelineResponseParams `json:"Response"`
}

func (r *DescribeVulsNumberTimelineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulsNumberTimelineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulsRequestParams struct {
	// 站点ID
	SiteId *uint64 `json:"SiteId,omitempty" name:"SiteId"`

	// 监控任务ID
	MonitorId *uint64 `json:"MonitorId,omitempty" name:"MonitorId"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeVulsRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID
	SiteId *uint64 `json:"SiteId,omitempty" name:"SiteId"`

	// 监控任务ID
	MonitorId *uint64 `json:"MonitorId,omitempty" name:"MonitorId"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeVulsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SiteId")
	delete(f, "MonitorId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulsResponseParams struct {
	// 漏洞数量。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 漏洞信息列表。
	Vuls []*Vul `json:"Vuls,omitempty" name:"Vuls"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeVulsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulsResponseParams `json:"Response"`
}

func (r *DescribeVulsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 过滤键的名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 一个或者多个过滤值。
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type MiniSite struct {
	// 站点ID。
	SiteId *uint64 `json:"SiteId,omitempty" name:"SiteId"`

	// 站点Url。
	Url *string `json:"Url,omitempty" name:"Url"`
}

// Predefined struct for user
type ModifyConfigAttributeRequestParams struct {
	// 漏洞告警通知等级，4位分别代表：高危、中危、低危、提示
	NoticeLevel *string `json:"NoticeLevel,omitempty" name:"NoticeLevel"`
}

type ModifyConfigAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 漏洞告警通知等级，4位分别代表：高危、中危、低危、提示
	NoticeLevel *string `json:"NoticeLevel,omitempty" name:"NoticeLevel"`
}

func (r *ModifyConfigAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConfigAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NoticeLevel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyConfigAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyConfigAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyConfigAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyConfigAttributeResponseParams `json:"Response"`
}

func (r *ModifyConfigAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConfigAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMonitorAttributeRequestParams struct {
	// 监测任务ID
	MonitorId *uint64 `json:"MonitorId,omitempty" name:"MonitorId"`

	// 站点的url列表
	Urls []*string `json:"Urls,omitempty" name:"Urls"`

	// 任务名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 扫描模式，normal-正常扫描；deep-深度扫描
	ScannerType *string `json:"ScannerType,omitempty" name:"ScannerType"`

	// 扫描周期，单位小时，每X小时执行一次
	Crontab *uint64 `json:"Crontab,omitempty" name:"Crontab"`

	// 扫描速率限制，每秒发送X个HTTP请求
	RateLimit *uint64 `json:"RateLimit,omitempty" name:"RateLimit"`

	// 首次扫描开始时间
	FirstScanStartTime *string `json:"FirstScanStartTime,omitempty" name:"FirstScanStartTime"`

	// 监测状态：1-监测中；2-暂停监测
	MonitorStatus *uint64 `json:"MonitorStatus,omitempty" name:"MonitorStatus"`
}

type ModifyMonitorAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 监测任务ID
	MonitorId *uint64 `json:"MonitorId,omitempty" name:"MonitorId"`

	// 站点的url列表
	Urls []*string `json:"Urls,omitempty" name:"Urls"`

	// 任务名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 扫描模式，normal-正常扫描；deep-深度扫描
	ScannerType *string `json:"ScannerType,omitempty" name:"ScannerType"`

	// 扫描周期，单位小时，每X小时执行一次
	Crontab *uint64 `json:"Crontab,omitempty" name:"Crontab"`

	// 扫描速率限制，每秒发送X个HTTP请求
	RateLimit *uint64 `json:"RateLimit,omitempty" name:"RateLimit"`

	// 首次扫描开始时间
	FirstScanStartTime *string `json:"FirstScanStartTime,omitempty" name:"FirstScanStartTime"`

	// 监测状态：1-监测中；2-暂停监测
	MonitorStatus *uint64 `json:"MonitorStatus,omitempty" name:"MonitorStatus"`
}

func (r *ModifyMonitorAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMonitorAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MonitorId")
	delete(f, "Urls")
	delete(f, "Name")
	delete(f, "ScannerType")
	delete(f, "Crontab")
	delete(f, "RateLimit")
	delete(f, "FirstScanStartTime")
	delete(f, "MonitorStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMonitorAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMonitorAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyMonitorAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMonitorAttributeResponseParams `json:"Response"`
}

func (r *ModifyMonitorAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMonitorAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySiteAttributeRequestParams struct {
	// 站点ID
	SiteId *uint64 `json:"SiteId,omitempty" name:"SiteId"`

	// 站点名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 网站是否需要登录扫描：0-未知；-1-不需要；1-需要
	NeedLogin *int64 `json:"NeedLogin,omitempty" name:"NeedLogin"`

	// 登录后的cookie
	LoginCookie *string `json:"LoginCookie,omitempty" name:"LoginCookie"`

	// 用于测试cookie是否有效的URL
	LoginCheckUrl *string `json:"LoginCheckUrl,omitempty" name:"LoginCheckUrl"`

	// 用于测试cookie是否有效的关键字
	LoginCheckKw *string `json:"LoginCheckKw,omitempty" name:"LoginCheckKw"`

	// 禁止扫描器扫描的目录关键字
	ScanDisallow *string `json:"ScanDisallow,omitempty" name:"ScanDisallow"`
}

type ModifySiteAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID
	SiteId *uint64 `json:"SiteId,omitempty" name:"SiteId"`

	// 站点名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 网站是否需要登录扫描：0-未知；-1-不需要；1-需要
	NeedLogin *int64 `json:"NeedLogin,omitempty" name:"NeedLogin"`

	// 登录后的cookie
	LoginCookie *string `json:"LoginCookie,omitempty" name:"LoginCookie"`

	// 用于测试cookie是否有效的URL
	LoginCheckUrl *string `json:"LoginCheckUrl,omitempty" name:"LoginCheckUrl"`

	// 用于测试cookie是否有效的关键字
	LoginCheckKw *string `json:"LoginCheckKw,omitempty" name:"LoginCheckKw"`

	// 禁止扫描器扫描的目录关键字
	ScanDisallow *string `json:"ScanDisallow,omitempty" name:"ScanDisallow"`
}

func (r *ModifySiteAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySiteAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SiteId")
	delete(f, "Name")
	delete(f, "NeedLogin")
	delete(f, "LoginCookie")
	delete(f, "LoginCheckUrl")
	delete(f, "LoginCheckKw")
	delete(f, "ScanDisallow")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySiteAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySiteAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySiteAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifySiteAttributeResponseParams `json:"Response"`
}

func (r *ModifySiteAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySiteAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Monitor struct {
	// 监控任务ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 监控名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 监测状态：1-监测中；2-暂停监测。
	MonitorStatus *uint64 `json:"MonitorStatus,omitempty" name:"MonitorStatus"`

	// 监测模式，normal-正常扫描；deep-深度扫描。
	ScannerType *string `json:"ScannerType,omitempty" name:"ScannerType"`

	// 扫描周期，单位小时，每X小时执行一次。
	Crontab *uint64 `json:"Crontab,omitempty" name:"Crontab"`

	// 指定扫描类型，3位数每位依次表示：扫描Web漏洞、扫描系统漏洞、扫描系统端口。
	IncludedVulsTypes *string `json:"IncludedVulsTypes,omitempty" name:"IncludedVulsTypes"`

	// 速率限制，每秒发送X个HTTP请求。
	RateLimit *uint64 `json:"RateLimit,omitempty" name:"RateLimit"`

	// 首次扫描开始时间。
	FirstScanStartTime *string `json:"FirstScanStartTime,omitempty" name:"FirstScanStartTime"`

	// 扫描状态：0-待扫描（无任何扫描结果）；1-扫描中（正在进行扫描）；2-已扫描（有扫描结果且不正在扫描）；3-扫描完成待同步结果。
	ScanStatus *uint64 `json:"ScanStatus,omitempty" name:"ScanStatus"`

	// 上一次扫描完成时间。
	LastScanFinishTime *string `json:"LastScanFinishTime,omitempty" name:"LastScanFinishTime"`

	// 当前扫描开始时间，如扫描完成则为上一次扫描的开始时间。
	CurrentScanStartTime *string `json:"CurrentScanStartTime,omitempty" name:"CurrentScanStartTime"`

	// CreatedAt。
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// UpdatedAt。
	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`

	// 云用户appid。
	Appid *uint64 `json:"Appid,omitempty" name:"Appid"`

	// 扫描状态：0-待检测；1-检测完成
	ContentScanStatus *uint64 `json:"ContentScanStatus,omitempty" name:"ContentScanStatus"`
}

type MonitorMiniSite struct {
	// 站点ID。
	SiteId *uint64 `json:"SiteId,omitempty" name:"SiteId"`

	// 站点Url。
	Url *string `json:"Url,omitempty" name:"Url"`
}

type MonitorsDetail struct {
	// 监控任务基础信息。
	Basic *Monitor `json:"Basic,omitempty" name:"Basic"`

	// 监控任务包含的站点列表。
	Sites []*MonitorMiniSite `json:"Sites,omitempty" name:"Sites"`

	// 监控任务包含的站点列表数量。
	SiteNumber *uint64 `json:"SiteNumber,omitempty" name:"SiteNumber"`

	// 监控任务包含的受漏洞威胁的站点列表。
	ImpactSites []*MonitorMiniSite `json:"ImpactSites,omitempty" name:"ImpactSites"`

	// 监控任务包含的受漏洞威胁的站点列表数量。
	ImpactSiteNumber *uint64 `json:"ImpactSiteNumber,omitempty" name:"ImpactSiteNumber"`

	// 高风险漏洞数量。
	VulsHighNumber *uint64 `json:"VulsHighNumber,omitempty" name:"VulsHighNumber"`

	// 中风险漏洞数量。
	VulsMiddleNumber *uint64 `json:"VulsMiddleNumber,omitempty" name:"VulsMiddleNumber"`

	// 低风险漏洞数量。
	VulsLowNumber *uint64 `json:"VulsLowNumber,omitempty" name:"VulsLowNumber"`

	// 提示数量。
	VulsNoticeNumber *uint64 `json:"VulsNoticeNumber,omitempty" name:"VulsNoticeNumber"`

	// 监控任务包含的站点列表的平均扫描进度。
	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`

	// 扫描页面总数。
	PageCount *uint64 `json:"PageCount,omitempty" name:"PageCount"`

	// 内容检测数量。
	ContentNumber *uint64 `json:"ContentNumber,omitempty" name:"ContentNumber"`
}

type Site struct {
	// 站点ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 监控任务ID，为0时表示未加入监控任务。
	MonitorId *uint64 `json:"MonitorId,omitempty" name:"MonitorId"`

	// 站点url。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 站点名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 验证状态：0-未验证；1-已验证；2-验证失效，待重新验证。
	VerifyStatus *uint64 `json:"VerifyStatus,omitempty" name:"VerifyStatus"`

	// 监测状态：0-未监测；1-监测中；2-暂停监测。
	MonitorStatus *uint64 `json:"MonitorStatus,omitempty" name:"MonitorStatus"`

	// 扫描状态：0-待扫描（无任何扫描结果）；1-扫描中（正在进行扫描）；2-已扫描（有扫描结果且不正在扫描）；3-扫描完成待同步结果。
	ScanStatus *uint64 `json:"ScanStatus,omitempty" name:"ScanStatus"`

	// 最近一次的AIScanner的扫描任务id，注意取消的情况。
	LastScanTaskId *uint64 `json:"LastScanTaskId,omitempty" name:"LastScanTaskId"`

	// 最近一次扫描开始时间。
	LastScanStartTime *string `json:"LastScanStartTime,omitempty" name:"LastScanStartTime"`

	// 最近一次扫描完成时间。
	LastScanFinishTime *string `json:"LastScanFinishTime,omitempty" name:"LastScanFinishTime"`

	// 最近一次取消时间，取消即使用上一次扫描结果。
	LastScanCancelTime *string `json:"LastScanCancelTime,omitempty" name:"LastScanCancelTime"`

	// 最近一次扫描扫描的页面数。
	LastScanPageCount *uint64 `json:"LastScanPageCount,omitempty" name:"LastScanPageCount"`

	// normal-正常扫描；deep-深度扫描。
	LastScanScannerType *string `json:"LastScanScannerType,omitempty" name:"LastScanScannerType"`

	// 最近一次扫描高风险漏洞数量。
	LastScanVulsHighNum *uint64 `json:"LastScanVulsHighNum,omitempty" name:"LastScanVulsHighNum"`

	// 最近一次扫描中风险漏洞数量。
	LastScanVulsMiddleNum *uint64 `json:"LastScanVulsMiddleNum,omitempty" name:"LastScanVulsMiddleNum"`

	// 最近一次扫描低风险漏洞数量。
	LastScanVulsLowNum *uint64 `json:"LastScanVulsLowNum,omitempty" name:"LastScanVulsLowNum"`

	// 最近一次扫描提示信息数量。
	LastScanVulsNoticeNum *uint64 `json:"LastScanVulsNoticeNum,omitempty" name:"LastScanVulsNoticeNum"`

	// 记录添加时间。
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 记录最近修改时间。
	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`

	// 速率限制，每秒发送X个HTTP请求。
	LastScanRateLimit *uint64 `json:"LastScanRateLimit,omitempty" name:"LastScanRateLimit"`

	// 最近一次扫描漏洞总数量。
	LastScanVulsNum *uint64 `json:"LastScanVulsNum,omitempty" name:"LastScanVulsNum"`

	// 最近一次扫描提示总数量
	LastScanNoticeNum *uint64 `json:"LastScanNoticeNum,omitempty" name:"LastScanNoticeNum"`

	// 扫描进度，百分比整数
	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`

	// 云用户appid。
	Appid *uint64 `json:"Appid,omitempty" name:"Appid"`

	// 云用户标识。
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 网站是否需要登录扫描：0-未知；-1-不需要；1-需要。
	NeedLogin *int64 `json:"NeedLogin,omitempty" name:"NeedLogin"`

	// 登录后的cookie。
	LoginCookie *string `json:"LoginCookie,omitempty" name:"LoginCookie"`

	// 登录后的cookie是否有效：0-无效；1-有效。
	LoginCookieValid *uint64 `json:"LoginCookieValid,omitempty" name:"LoginCookieValid"`

	// 用于测试cookie是否有效的URL。
	LoginCheckUrl *string `json:"LoginCheckUrl,omitempty" name:"LoginCheckUrl"`

	// 用于测试cookie是否有效的关键字。
	LoginCheckKw *string `json:"LoginCheckKw,omitempty" name:"LoginCheckKw"`

	// 禁止扫描器扫描的目录关键字。
	ScanDisallow *string `json:"ScanDisallow,omitempty" name:"ScanDisallow"`

	// 访问网站的客户端标识。
	UserAgent *string `json:"UserAgent,omitempty" name:"UserAgent"`

	// 内容检测状态：0-未检测；1-已检测；
	ContentStatus *uint64 `json:"ContentStatus,omitempty" name:"ContentStatus"`

	// 最近一次扫描内容检测数量
	LastScanContentNum *uint64 `json:"LastScanContentNum,omitempty" name:"LastScanContentNum"`
}

type SitesVerification struct {
	// 根域名。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// txt解析域名验证的name。
	TxtName *string `json:"TxtName,omitempty" name:"TxtName"`

	// txt解析域名验证的text。
	TxtText *string `json:"TxtText,omitempty" name:"TxtText"`

	// 验证有效期，在此之前有效。
	ValidTo *string `json:"ValidTo,omitempty" name:"ValidTo"`

	// 验证状态：0-未验证；1-已验证；2-验证失效，待重新验证。
	VerifyStatus *uint64 `json:"VerifyStatus,omitempty" name:"VerifyStatus"`

	// CreatedAt。
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// UpdatedAt。
	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`

	// ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 云用户appid
	Appid *uint64 `json:"Appid,omitempty" name:"Appid"`

	// 用于验证站点的url，即访问该url获取验证数据。
	VerifyUrl *string `json:"VerifyUrl,omitempty" name:"VerifyUrl"`

	// 获取验证验证文件的url。
	VerifyFileUrl *string `json:"VerifyFileUrl,omitempty" name:"VerifyFileUrl"`
}

// Predefined struct for user
type VerifySitesRequestParams struct {
	// 站点的url列表
	Urls []*string `json:"Urls,omitempty" name:"Urls"`
}

type VerifySitesRequest struct {
	*tchttp.BaseRequest
	
	// 站点的url列表
	Urls []*string `json:"Urls,omitempty" name:"Urls"`
}

func (r *VerifySitesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifySitesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Urls")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifySitesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifySitesResponseParams struct {
	// 验证成功的根域名数量。
	SuccessNumber *uint64 `json:"SuccessNumber,omitempty" name:"SuccessNumber"`

	// 验证失败的根域名数量。
	FailNumber *uint64 `json:"FailNumber,omitempty" name:"FailNumber"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type VerifySitesResponse struct {
	*tchttp.BaseResponse
	Response *VerifySitesResponseParams `json:"Response"`
}

func (r *VerifySitesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifySitesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Vul struct {
	// 漏洞ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 站点ID。
	SiteId *uint64 `json:"SiteId,omitempty" name:"SiteId"`

	// 扫描引擎的扫描任务ID。
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`

	// 漏洞级别：high、middle、low、notice。
	Level *string `json:"Level,omitempty" name:"Level"`

	// 漏洞名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 出现漏洞的url。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 网址/细节。
	Html *string `json:"Html,omitempty" name:"Html"`

	// 漏洞类型。
	Nickname *string `json:"Nickname,omitempty" name:"Nickname"`

	// 危害说明。
	Harm *string `json:"Harm,omitempty" name:"Harm"`

	// 漏洞描述。
	Describe *string `json:"Describe,omitempty" name:"Describe"`

	// 解决方案。
	Solution *string `json:"Solution,omitempty" name:"Solution"`

	// 漏洞参考。
	From *string `json:"From,omitempty" name:"From"`

	// 漏洞通过该参数攻击。
	Parameter *string `json:"Parameter,omitempty" name:"Parameter"`

	// CreatedAt。
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// UpdatedAt。
	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`

	// 是否已经添加误报，0-否，1-是。
	IsReported *uint64 `json:"IsReported,omitempty" name:"IsReported"`

	// 云用户appid。
	Appid *uint64 `json:"Appid,omitempty" name:"Appid"`

	// 云用户标识。
	Uin *string `json:"Uin,omitempty" name:"Uin"`
}

type VulsTimeline struct {
	// ID。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 云用户appid。
	Appid *uint64 `json:"Appid,omitempty" name:"Appid"`

	// 日期。
	Date *string `json:"Date,omitempty" name:"Date"`

	// 扫描页面总数量。
	PageCount *uint64 `json:"PageCount,omitempty" name:"PageCount"`

	// 已验证网站总数量。
	SiteNum *uint64 `json:"SiteNum,omitempty" name:"SiteNum"`

	// 受影响的网站总数量。
	ImpactSiteNum *uint64 `json:"ImpactSiteNum,omitempty" name:"ImpactSiteNum"`

	// 高危漏洞总数量。
	VulsHighNum *uint64 `json:"VulsHighNum,omitempty" name:"VulsHighNum"`

	// 中危漏洞总数量。
	VulsMiddleNum *uint64 `json:"VulsMiddleNum,omitempty" name:"VulsMiddleNum"`

	// 低危漏洞总数量。
	VulsLowNum *uint64 `json:"VulsLowNum,omitempty" name:"VulsLowNum"`

	// 风险提示总数量
	VulsNoticeNum *uint64 `json:"VulsNoticeNum,omitempty" name:"VulsNoticeNum"`

	// 记录添加时间。
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 记录最近修改时间。
	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}