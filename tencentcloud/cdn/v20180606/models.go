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

package v20180606

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type CacheOptResult struct {

	// 成功的url列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessUrls []*string `json:"SuccessUrls,omitempty" name:"SuccessUrls" list`

	// 失败的url列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailUrls []*string `json:"FailUrls,omitempty" name:"FailUrls" list`
}

type CdnData struct {

	// 查询指定的指标名称：
	// flux：流量，单位为 byte
	// bandwidth：带宽，单位为 bps
	// request：请求数，单位为 次
	// fluxHitRate：流量命中率，单位为 %
	// statusCode：状态码，返回 2XX、3XX、4XX、5XX 汇总数据，单位为 个
	// 2XX：返回 2XX 状态码汇总及各 2 开头状态码数据，单位为 个
	// 3XX：返回 3XX 状态码汇总及各 3 开头状态码数据，单位为 个
	// 4XX：返回 4XX 状态码汇总及各 4 开头状态码数据，单位为 个
	// 5XX：返回 5XX 状态码汇总及各 5 开头状态码数据，单位为 个
	// 或指定查询的某一具体状态码
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// 明细数据组合
	DetailData []*TimestampData `json:"DetailData,omitempty" name:"DetailData" list`

	// 汇总数据组合
	SummarizedData *SummarizedData `json:"SummarizedData,omitempty" name:"SummarizedData"`
}

type CdnIp struct {

	// 节点 ip。
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 是否为腾讯云 CDN 加速节点。yes 表示该节点为腾讯云 CDN 节点，no 表示该节点不是腾讯云 CDN 节点。
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 表示该节点所处的省份/国家。unknown 表示节点位置未知。
	Location *string `json:"Location,omitempty" name:"Location"`

	// 节点上下线历史记录。
	History []*CdnIpHistory `json:"History,omitempty" name:"History" list`

	// 节点的服务地域。mainland 表示服务地域为中国境内，overseas 表示服务地域为中国境外， unknown 表示服务地域未知。
	Area *string `json:"Area,omitempty" name:"Area"`
}

type CdnIpHistory struct {

	// 上下线状态。online 为上线，offline 为下线。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 操作时间。当该值为 null 时表示无历史状态变更记录。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Datetime *string `json:"Datetime,omitempty" name:"Datetime"`
}

type DescribeCdnDataRequest struct {
	*tchttp.BaseRequest

	// 查询起始时间，如：2018-09-04 10:40:00，返回结果大于等于指定时间
	// 根据指定时间粒度不同，会进行向前归整，如 2018-09-04 10:40:00 在按 1 小时的时间粒度查询时，返回的第一个数据对应时间点为 2018-09-04 10:00:00
	// 起始时间与结束时间间隔小于等于 90 天
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，如：2018-09-04 10:40:00，返回结果小于等于指定时间
	// 根据指定时间粒度不同，会进行向前归整，如 2018-09-04 10:40:00 在按 1 小时的时间粒度查询时，返回的最后一个数据对应时间点为 2018-09-04 10:00:00
	// 起始时间与结束时间间隔小于等于 90 天
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 指定查询指标，支持的类型有：
	// flux：流量，单位为 byte
	// bandwidth：带宽，单位为 bps
	// request：请求数，单位为 次
	// fluxHitRate：流量命中率，单位为 %
	// statusCode：状态码，返回 2xx、3xx、4xx、5xx 汇总数据，单位为 个
	// 2xx：返回 2xx 状态码汇总及各 2 开头状态码数据，单位为 个
	// 3xx：返回 3xx 状态码汇总及各 3 开头状态码数据，单位为 个
	// 4xx：返回 4xx 状态码汇总及各 4 开头状态码数据，单位为 个
	// 5xx：返回 5xx 状态码汇总及各 5 开头状态码数据，单位为 个
	// 支持指定具体状态码查询，若未产生过，则返回为空
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// 指定查询域名列表
	// 最多可一次性查询 30 个加速域名明细
	Domains []*string `json:"Domains,omitempty" name:"Domains" list`

	// 指定要查询的项目 ID，[前往查看项目 ID](https://console.cloud.tencent.com/project)
	// 未填充域名情况下，指定项目查询，若填充了具体域名信息，以域名为主
	Project *int64 `json:"Project,omitempty" name:"Project"`

	// 时间粒度，支持以下几种模式：
	// min：1 分钟粒度，指定查询区间 24 小时内（含 24 小时），可返回 1 分钟粒度明细数据（指定查询服务地域为中国境外时不支持 1 分钟粒度）
	// 5min：5 分钟粒度，指定查询区间 31 天内（含 31 天），可返回 5 分钟粒度明细数据
	// hour：1 小时粒度，指定查询区间 31 天内（含 31 天），可返回 1 小时粒度明细数据
	// day：天粒度，指定查询区间大于 31 天，可返回天粒度明细数据
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// 多域名查询时，默认（false)返回多个域名的汇总数据
	// 可按需指定为 true，返回每一个 Domain 的明细数据（statusCode 指标暂不支持）
	Detail *bool `json:"Detail,omitempty" name:"Detail"`

	// 查询中国境内CDN数据时，指定运营商查询，不填充表示查询所有运营商
	// 运营商编码可以查看 [运营商编码映射](https://cloud.tencent.com/document/product/228/6316#.E8.BF.90.E8.90.A5.E5.95.86.E6.98.A0.E5.B0.84)
	// 指定运营商查询时，不可同时指定省份、IP协议查询
	Isp *int64 `json:"Isp,omitempty" name:"Isp"`

	// 查询中国境内CDN数据时，指定省份查询，不填充表示查询所有省份
	// 查询中国境外CDN数据时，指定国家/地区查询，不填充表示查询所有国家/地区
	// 省份、国家/地区编码可以查看 [省份编码映射](https://cloud.tencent.com/document/product/228/6316#.E7.9C.81.E4.BB.BD.E6.98.A0.E5.B0.84)
	// 指定（中国境内）省份查询时，不可同时指定运营商、IP协议查询
	District *int64 `json:"District,omitempty" name:"District"`

	// 指定协议查询，不填充表示查询所有协议
	// all：所有协议
	// http：指定查询 HTTP 对应指标
	// https：指定查询 HTTPS 对应指标
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 指定数据源查询，白名单功能
	DataSource *string `json:"DataSource,omitempty" name:"DataSource"`

	// 指定IP协议查询，不填充表示查询所有协议
	// all：所有协议
	// ipv4：指定查询 ipv4 对应指标
	// ipv6：指定查询 ipv6 对应指标
	// 指定IP协议查询时，不可同时指定省份、运营商查询
	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`

	// 指定服务地域查询，不填充表示查询中国境内CDN数据
	// mainland：指定查询中国境内 CDN 数据
	// overseas：指定查询中国境外 CDN 数据
	Area *string `json:"Area,omitempty" name:"Area"`

	// 查询中国境外CDN数据时，可指定地区类型查询，不填充表示查询服务地区数据（仅在 Area 为 overseas 时可用）
	// server：指定查询服务地区（腾讯云 CDN 节点服务器所在地区）数据
	// client：指定查询客户端地区（用户请求终端所在地区）数据
	AreaType *string `json:"AreaType,omitempty" name:"AreaType"`
}

func (r *DescribeCdnDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCdnDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCdnDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回数据的时间粒度，查询时指定：
	// min：1 分钟粒度
	// 5min：5 分钟粒度
	// hour：1 小时粒度
	// day：天粒度
		Interval *string `json:"Interval,omitempty" name:"Interval"`

		// 指定条件查询得到的数据明细
		Data []*ResourceData `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCdnDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCdnDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCdnIpRequest struct {
	*tchttp.BaseRequest

	// 需要查询的 IP 列表
	Ips []*string `json:"Ips,omitempty" name:"Ips" list`
}

func (r *DescribeCdnIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCdnIpRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCdnIpResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询的节点归属详情。
		Ips []*CdnIp `json:"Ips,omitempty" name:"Ips" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCdnIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCdnIpResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeIpVisitRequest struct {
	*tchttp.BaseRequest

	// 查询起始时间，如：2018-09-04 10:40:10，返回结果大于等于指定时间
	// 根据指定时间粒度不同，会进行向前归整，如 2018-09-04 10:40:10 在按 5 分钟的时间粒度查询时，返回的第一个数据对应时间点为 2018-09-04 10:40:00
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，如：2018-09-04 10:40:10，返回结果小于等于指定时间
	// 根据指定时间粒度不同，会进行向前归整，如 2018-09-04 10:40:10 在按 5 分钟的时间粒度查询时，返回的最后一个数据对应时间点为 2018-09-04 10:40:00
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 指定查询域名列表，最多可一次性查询 30 个加速域名明细
	Domains []*string `json:"Domains,omitempty" name:"Domains" list`

	// 指定要查询的项目 ID，[前往查看项目 ID](https://console.cloud.tencent.com/project)
	// 未填充域名情况下，指定项目查询，若填充了具体域名信息，以域名为主
	Project *int64 `json:"Project,omitempty" name:"Project"`

	// 时间粒度，支持以下几种模式：
	// 5min：5 分钟粒度，查询时间区间 24 小时内，默认返回 5 分钟粒度活跃用户数
	// day：天粒度，查询时间区间大于 1 天时，默认返回天粒度活跃用户数
	Interval *string `json:"Interval,omitempty" name:"Interval"`
}

func (r *DescribeIpVisitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeIpVisitRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeIpVisitResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据统计的时间粒度，支持5min,  day，分别表示5分钟，1天的时间粒度。
		Interval *string `json:"Interval,omitempty" name:"Interval"`

		// 各个资源的回源数据详情。
		Data []*ResourceData `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIpVisitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeIpVisitResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMapInfoRequest struct {
	*tchttp.BaseRequest

	// 映射查询类别：
	// isp：运营商映射查询
	// district：省份（中国境内）、国家/地区（中国境外）映射查询
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeMapInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMapInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMapInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 映射关系数组。
		MapInfoList []*MapInfo `json:"MapInfoList,omitempty" name:"MapInfoList" list`

		// 服务端区域id和子区域id的映射关系。
	// 注意：此字段可能返回 null，表示取不到有效值。
		ServerRegionRelation []*RegionMapRelation `json:"ServerRegionRelation,omitempty" name:"ServerRegionRelation" list`

		// 客户端区域id和子区域id的映射关系。
	// 注意：此字段可能返回 null，表示取不到有效值。
		ClientRegionRelation []*RegionMapRelation `json:"ClientRegionRelation,omitempty" name:"ClientRegionRelation" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMapInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMapInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeOriginDataRequest struct {
	*tchttp.BaseRequest

	// 查询起始时间，如：2018-09-04 10:40:00，返回结果大于等于指定时间
	// 根据指定时间粒度不同，会进行向前归整，如 2018-09-04 10:40:00 在按 1 小时的时间粒度查询时，返回的第一个数据对应时间点为 2018-09-04 10:00:00
	// 起始时间与结束时间间隔小于等于 90 天
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间，如：2018-09-04 10:40:00，返回结果小于等于指定时间
	// 根据指定时间粒度不同，会进行向前归整，如 2018-09-04 10:40:00 在按 1 小时的时间粒度查询时，返回的最后一个数据对应时间点为 2018-09-04 10:00:00
	// 起始时间与结束时间间隔小于等于 90 天
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 指定查询指标，支持的类型有：
	// flux：回源流量，单位为 byte
	// bandwidth：回源带宽，单位为 bps
	// request：回源请求数，单位为 次
	// failRequest：回源失败请求数，单位为 次
	// failRate：回源失败率，单位为 %
	// statusCode：回源状态码，返回 2xx、3xx、4xx、5xx 汇总数据，单位为 个
	// 2xx：返回 2xx 回源状态码汇总及各 2 开头回源状态码数据，单位为 个
	// 3xx：返回 3xx 回源状态码汇总及各 3 开头回源状态码数据，单位为 个
	// 4xx：返回 4xx 回源状态码汇总及各 4 开头回源状态码数据，单位为 个
	// 5xx：返回 5xx 回源状态码汇总及各 5 开头回源状态码数据，单位为 个
	// 支持指定具体状态码查询，若未产生过，则返回为空
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// 指定查询域名列表，最多可一次性查询 30 个加速域名明细
	Domains []*string `json:"Domains,omitempty" name:"Domains" list`

	// 指定要查询的项目 ID，[前往查看项目 ID](https://console.cloud.tencent.com/project)
	// 未填充域名情况下，指定项目查询，若填充了具体域名信息，以域名为主
	Project *int64 `json:"Project,omitempty" name:"Project"`

	// 时间粒度，支持以下几种模式：
	// min：1 分钟粒度，指定查询区间 24 小时内（含 24 小时），可返回 1 分钟粒度明细数据（指定查询服务地域为中国境外时不支持 1 分钟粒度）
	// 5min：5 分钟粒度，指定查询区间 31 天内（含 31 天），可返回 5 分钟粒度明细数据
	// hour：1 小时粒度，指定查询区间 31 天内（含 31 天），可返回 1 小时粒度明细数据
	// day：天粒度，指定查询区间大于 31 天，可返回天粒度明细数据
	Interval *string `json:"Interval,omitempty" name:"Interval"`

	// Domains 传入多个时，默认（false)返回多个域名的汇总数据
	// 可按需指定为 true，返回每一个 Domain 的明细数据（statusCode 指标暂不支持）
	Detail *bool `json:"Detail,omitempty" name:"Detail"`

	// 指定服务地域查询，不填充表示查询中国境内 CDN 数据
	// mainland：指定查询中国境内 CDN 数据
	// overseas：指定查询中国境外 CDN 数据
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeOriginDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeOriginDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeOriginDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据统计的时间粒度，支持min, 5min, hour, day，分别表示1分钟，5分钟，1小时和1天的时间粒度。
		Interval *string `json:"Interval,omitempty" name:"Interval"`

		// 各个资源的回源数据详情。
		Data []*ResourceOriginData `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOriginDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeOriginDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePayTypeRequest struct {
	*tchttp.BaseRequest

	// 指定服务地域查询，不填充表示查询中国境内 CDN 计费方式
	// mainland：指定查询中国境内 CDN 计费方式
	// overseas：指定查询中国境外 CDN 计费方式
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribePayTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePayTypeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePayTypeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 计费类型：
	// flux：流量计费
	// bandwidth：带宽计费
	// 如果修改过计费方式，表示下次生效的计费类型，否则表示当前计费类型。
		PayType *string `json:"PayType,omitempty" name:"PayType"`

		// 计费周期：
	// day：日结计费
	// month：月结计费
		BillingCycle *string `json:"BillingCycle,omitempty" name:"BillingCycle"`

		// 计费方式：
	// monthMax：日峰值月平均计费，月结模式
	// day95：日 95 带宽计费，月结模式
	// month95：月95带宽计费，月结模式
	// sum：总流量计费，日结与月结均有流量计费模式
	// max：峰值带宽计费，日结模式
		StatType *string `json:"StatType,omitempty" name:"StatType"`

		// 地区计费方式，仅在查询中国境外 CDN 计费方式时可用
	// all：表示全地区统一计费
	// multiple：表示分地区计费。
		RegionType *string `json:"RegionType,omitempty" name:"RegionType"`

		// 当前计费类型：
	// flux：流量计费
	// bandwidth：带宽计费
		CurrentPayType *string `json:"CurrentPayType,omitempty" name:"CurrentPayType"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePayTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePayTypeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

	// 分页查询偏移量，默认为 0 （第一页）。
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

func (r *DescribePurgeTasksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePurgeTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 刷新历史记录
	// 注意：此字段可能返回 null，表示取不到有效值。
		PurgeLogs []*PurgeTask `json:"PurgeLogs,omitempty" name:"PurgeLogs" list`

		// 任务总数，用于分页
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePurgeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePurgeTasksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePushTasksRequest struct {
	*tchttp.BaseRequest

	// 开始时间，如2018-08-08 00:00:00。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，如2018-08-08 23:59:59。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 提交时返回的任务 Id，查询时 TaskId 和起始时间必须指定一项。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 查询关键字，请输入域名或 http(s):// 开头完整 URL。
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// 分页查询偏移量，默认为 0 （第一页）。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询限制数目，默认为20。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 查询刷新记录指定地区。mainland：中国大陆。
	Area *string `json:"Area,omitempty" name:"Area"`

	// 查询指定任务状态，fail表示失败，done表示成功，process表示刷新中。
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *DescribePushTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePushTasksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePushTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 预热历史记录
	// 注意：此字段可能返回 null，表示取不到有效值。
		PushLogs []*PushTask `json:"PushLogs,omitempty" name:"PushLogs" list`

		// 任务总数，用于分页
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePushTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePushTasksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableCachesRequest struct {
	*tchttp.BaseRequest

	// 需要禁用的 URL 列表
	// 每次最多可提交 100 条，每日最多可提交 3000 条
	Urls []*string `json:"Urls,omitempty" name:"Urls" list`
}

func (r *DisableCachesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableCachesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableCachesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 提交结果
	// 注意：此字段可能返回 null，表示取不到有效值。
		CacheOptResult *CacheOptResult `json:"CacheOptResult,omitempty" name:"CacheOptResult"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableCachesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableCachesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableCachesRequest struct {
	*tchttp.BaseRequest

	// 解封 URL 列表
	Urls []*string `json:"Urls,omitempty" name:"Urls" list`
}

func (r *EnableCachesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableCachesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableCachesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		CacheOptResult *CacheOptResult `json:"CacheOptResult,omitempty" name:"CacheOptResult"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableCachesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableCachesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetDisableRecordsRequest struct {
	*tchttp.BaseRequest

	// 开始时间，如：2018-12-12 10:24:00。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间，如：2018-12-14 10:24:00。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 指定 URL 查询
	Url *string `json:"Url,omitempty" name:"Url"`

	// URL 当前状态
	// disable：当前仍为禁用状态，访问返回 403
	// enable：当前为可用状态，已解禁，可正常访问
	Status *string `json:"Status,omitempty" name:"Status"`

	// 分页查询偏移量，默认为 0 （第一页）。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页查询限制数目，默认为20。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetDisableRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDisableRecordsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetDisableRecordsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 封禁历史记录
	// 注意：此字段可能返回 null，表示取不到有效值。
		UrlRecordList []*UrlRecord `json:"UrlRecordList,omitempty" name:"UrlRecordList" list`

		// 任务总数，用于分页
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDisableRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDisableRecordsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListTopDataRequest struct {
	*tchttp.BaseRequest

	// 查询起始日期，如：2018-09-09 00:00:00。目前只支持按天粒度的数据查询，只取入参中的天数信息。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束日期，如：2018-09-10 00:00:00。目前只支持按天粒度的数据查询，只取入参中的天数信息。例如，要查询2018-09-10的数据，输入StartTime=2018-09-10 00:00:00，EndTime=2018-09-10 00:00:00即可。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 排序对象，支持以下几种形式：
	// url：访问 URL 排序，带参数统计，支持的 Filter 为 flux、request
	// path：访问 URL 排序，不带参数统计，支持的 Filter 为 flux、request（白名单功能）
	// district：省份、国家/地区排序，支持的 Filter 为 flux、request
	// isp：运营商排序，支持的 Filter 为 flux、request
	// host：域名访问数据排序，支持的 Filter 为：flux, request, bandwidth, fluxHitRate, 2XX, 3XX, 4XX, 5XX，具体状态码统计
	// originHost：域名回源数据排序，支持的 Filter 为 flux， request，bandwidth，origin_2XX，origin_3XX，oringin_4XX，origin_5XX，具体回源状态码统计
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// 排序使用的指标名称：
	// flux：Metric 为 host 时指代访问流量，originHost 时指代回源流量
	// bandwidth：Metric 为 host 时指代访问带宽，originHost 时指代回源带宽
	// request：Metric 为 host 时指代访问请求数，originHost 时指代回源请求数
	// fluxHitRate：平均流量命中率
	// 2XX：访问 2XX 状态码
	// 3XX：访问 3XX 状态码
	// 4XX：访问 4XX 状态码
	// 5XX：访问 5XX 状态码
	// origin_2XX：回源 2XX 状态码
	// origin_3XX：回源 3XX 状态码
	// origin_4XX：回源 4XX 状态码
	// origin_5XX：回源 5XX 状态码
	// statusCode：指定访问状态码统计，在 Code 参数中填充指定状态码
	// OriginStatusCode：指定回源状态码统计，在 Code 参数中填充指定状态码
	Filter *string `json:"Filter,omitempty" name:"Filter"`

	// 指定查询域名列表，最多可一次性查询 30 个加速域名明细
	Domains []*string `json:"Domains,omitempty" name:"Domains" list`

	// 指定要查询的项目 ID，[前往查看项目 ID](https://console.cloud.tencent.com/project)
	// 未填充域名情况下，指定项目查询，若填充了具体域名信息，以域名为主
	Project *int64 `json:"Project,omitempty" name:"Project"`

	// 多域名查询时，默认（false)返回所有域名汇总排序结果
	// Metric 为 Url、Path、District、Isp，Filter 为 flux、reqeust 时，可设置为 true，返回每一个 Domain 的排序数据
	Detail *bool `json:"Detail,omitempty" name:"Detail"`

	// Filter 为 statusCode、OriginStatusCode 时，填充指定状态码查询排序结果
	Code *string `json:"Code,omitempty" name:"Code"`

	// 指定服务地域查询，不填充表示查询中国境内 CDN 数据
	// mainland：指定查询中国境内 CDN 数据
	// overseas：指定查询中国境外 CDN 数据，支持的 Metric 为 url、district、host、originHost，当 Metric 为 originHost 时仅支持 flux、request、bandwidth Filter
	Area *string `json:"Area,omitempty" name:"Area"`

	// 查询中国境外CDN数据，且仅当 Metric 为 District 或 Host 时，可指定地区类型查询，不填充表示查询服务地区数据（仅在 Area 为 overseas，且 Metric 是 District 或 Host 时可用）
	// server：指定查询服务地区（腾讯云 CDN 节点服务器所在地区）数据
	// client：指定查询客户端地区（用户请求终端所在地区）数据，当 Metric 为 host 时仅支持 flux、request、bandwidth Filter
	AreaType *string `json:"AreaType,omitempty" name:"AreaType"`
}

func (r *ListTopDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListTopDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListTopDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 各个资源的Top 访问数据详情。
		Data []*TopData `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListTopDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListTopDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MapInfo struct {

	// 对象 Id
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 对象名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

type PurgePathCacheRequest struct {
	*tchttp.BaseRequest

	// 要刷新的目录列表，必须包含协议头部。
	Paths []*string `json:"Paths,omitempty" name:"Paths" list`

	// 刷新类型，flush 代表刷新有更新的资源，delete 表示刷新全部资源。
	FlushType *string `json:"FlushType,omitempty" name:"FlushType"`
}

func (r *PurgePathCacheRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PurgePathCacheRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PurgePathCacheResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 刷新任务Id，前十位为提交任务时的UTC时间。
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PurgePathCacheResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

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

type PurgeUrlsCacheRequest struct {
	*tchttp.BaseRequest

	// 要刷新的Url列表，必须包含协议头部。
	Urls []*string `json:"Urls,omitempty" name:"Urls" list`
}

func (r *PurgeUrlsCacheRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PurgeUrlsCacheRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PurgeUrlsCacheResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 刷新任务Id，前十位为提交任务时的UTC时间。
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PurgeUrlsCacheResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PurgeUrlsCacheResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PushTask struct {

	// 预热任务Id，前十位为时间戳。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 预热Url。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 预热任务状态，fail表示失败，done表示成功，process表示预热中。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 预热百分比。
	Percent *int64 `json:"Percent,omitempty" name:"Percent"`

	// 预热任务提交时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 预热区域，mainland，overseas或global。
	Area *string `json:"Area,omitempty" name:"Area"`
}

type PushUrlsCacheRequest struct {
	*tchttp.BaseRequest

	// URL 列表，提交时需要包含协议头部（http:// 或 https://）
	Urls []*string `json:"Urls,omitempty" name:"Urls" list`

	// 预热请求回源时 HTTP 请求的 User-Agent 头部，默认为 TencentCdn
	UserAgent *string `json:"UserAgent,omitempty" name:"UserAgent"`
}

func (r *PushUrlsCacheRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PushUrlsCacheRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PushUrlsCacheResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 此批次提交任务对应的 Id，值唯一
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PushUrlsCacheResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PushUrlsCacheResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RegionMapRelation struct {

	// 区域ID。
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// 子区域ID列表
	SubRegionIdList []*int64 `json:"SubRegionIdList,omitempty" name:"SubRegionIdList" list`
}

type ResourceData struct {

	// 资源名称，根据查询条件不同分为以下几类：
	// 具体域名：表示该域名明细数据
	// multiDomains：表示多域名汇总明细数据
	// 项目 ID：指定项目查询时，显示为项目 ID
	// all：账号维度明细数据
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// 资源对应的数据明细
	CdnData []*CdnData `json:"CdnData,omitempty" name:"CdnData" list`
}

type ResourceOriginData struct {

	// 资源名称，根据查询条件不同分为以下几类：
	// 具体域名：表示该域名明细数据
	// multiDomains：表示多域名汇总明细数据
	// 项目 ID：指定项目查询时，显示为项目 ID
	// all：账号维度明细数据
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// 回源数据详情
	OriginData []*CdnData `json:"OriginData,omitempty" name:"OriginData" list`
}

type SummarizedData struct {

	// 汇总方式，存在以下几种：
	// sum：累加求和
	// max：最大值，带宽模式下，采用 5 分钟粒度汇总数据，计算峰值带宽
	// avg：平均值
	Name *string `json:"Name,omitempty" name:"Name"`

	// 汇总后的数据值
	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type TimestampData struct {

	// 数据统计时间点，采用向前汇总模式
	// 以 5 分钟粒度为例，13:35:00 时间点代表的统计数据区间为 13:35:00 至 13:39:59
	Time *string `json:"Time,omitempty" name:"Time"`

	// 数据值
	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type TopData struct {

	// 资源名称，根据查询条件不同分为以下几类：
	// 具体域名：表示该域名明细数据
	// multiDomains：表示多域名汇总明细数据
	// 项目 ID：指定项目查询时，显示为项目 ID
	// all：账号维度明细数据
	Resource *string `json:"Resource,omitempty" name:"Resource"`

	// 排序结果详情
	DetailData []*TopDetailData `json:"DetailData,omitempty" name:"DetailData" list`
}

type TopDetailData struct {

	// 数据类型的名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 数据值
	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type UrlRecord struct {

	// 状态(disable表示封禁，enable表示解封)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 对应的url
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealUrl *string `json:"RealUrl,omitempty" name:"RealUrl"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}
