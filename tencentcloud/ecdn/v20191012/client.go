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
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-10-12"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewAddEcdnDomainRequest() (request *AddEcdnDomainRequest) {
    request = &AddEcdnDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "AddEcdnDomain")
    return
}

func NewAddEcdnDomainResponse() (response *AddEcdnDomainResponse) {
    response = &AddEcdnDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（AddEcdnDomain）用于创建加速域名。
func (c *Client) AddEcdnDomain(request *AddEcdnDomainRequest) (response *AddEcdnDomainResponse, err error) {
    if request == nil {
        request = NewAddEcdnDomainRequest()
    }
    response = NewAddEcdnDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEcdnDomainRequest() (request *DeleteEcdnDomainRequest) {
    request = &DeleteEcdnDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "DeleteEcdnDomain")
    return
}

func NewDeleteEcdnDomainResponse() (response *DeleteEcdnDomainResponse) {
    response = &DeleteEcdnDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeleteEcdnDomain）用于删除指定加速域名。待删除域名必须处于已停用状态。
func (c *Client) DeleteEcdnDomain(request *DeleteEcdnDomainRequest) (response *DeleteEcdnDomainResponse, err error) {
    if request == nil {
        request = NewDeleteEcdnDomainRequest()
    }
    response = NewDeleteEcdnDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainsRequest() (request *DescribeDomainsRequest) {
    request = &DescribeDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "DescribeDomains")
    return
}

func NewDescribeDomainsResponse() (response *DescribeDomainsResponse) {
    response = &DescribeDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeDomains）用于查询CDN域名基本信息，包括项目id，状态，业务类型，创建时间，更新时间等。
func (c *Client) DescribeDomains(request *DescribeDomainsRequest) (response *DescribeDomainsResponse, err error) {
    if request == nil {
        request = NewDescribeDomainsRequest()
    }
    response = NewDescribeDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainsConfigRequest() (request *DescribeDomainsConfigRequest) {
    request = &DescribeDomainsConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "DescribeDomainsConfig")
    return
}

func NewDescribeDomainsConfigResponse() (response *DescribeDomainsConfigResponse) {
    response = &DescribeDomainsConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeDomainsConfig）用于查询CDN加速域名详细配置信息。
func (c *Client) DescribeDomainsConfig(request *DescribeDomainsConfigRequest) (response *DescribeDomainsConfigResponse, err error) {
    if request == nil {
        request = NewDescribeDomainsConfigRequest()
    }
    response = NewDescribeDomainsConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEcdnDomainLogsRequest() (request *DescribeEcdnDomainLogsRequest) {
    request = &DescribeEcdnDomainLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "DescribeEcdnDomainLogs")
    return
}

func NewDescribeEcdnDomainLogsResponse() (response *DescribeEcdnDomainLogsResponse) {
    response = &DescribeEcdnDomainLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeEcdnDomainLogs）用于查询域名的访问日志下载地址。
func (c *Client) DescribeEcdnDomainLogs(request *DescribeEcdnDomainLogsRequest) (response *DescribeEcdnDomainLogsResponse, err error) {
    if request == nil {
        request = NewDescribeEcdnDomainLogsRequest()
    }
    response = NewDescribeEcdnDomainLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEcdnDomainStatisticsRequest() (request *DescribeEcdnDomainStatisticsRequest) {
    request = &DescribeEcdnDomainStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "DescribeEcdnDomainStatistics")
    return
}

func NewDescribeEcdnDomainStatisticsResponse() (response *DescribeEcdnDomainStatisticsResponse) {
    response = &DescribeEcdnDomainStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeEcdnDomainStatistics）用于查询指定时间段内的域名访问统计指标
func (c *Client) DescribeEcdnDomainStatistics(request *DescribeEcdnDomainStatisticsRequest) (response *DescribeEcdnDomainStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeEcdnDomainStatisticsRequest()
    }
    response = NewDescribeEcdnDomainStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEcdnStatisticsRequest() (request *DescribeEcdnStatisticsRequest) {
    request = &DescribeEcdnStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "DescribeEcdnStatistics")
    return
}

func NewDescribeEcdnStatisticsResponse() (response *DescribeEcdnStatisticsResponse) {
    response = &DescribeEcdnStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEcdnStatistics用于查询 ECDN 实时访问监控数据，支持以下指标查询：
// 
// + 流量（单位为 byte）
// + 带宽（单位为 bps）
// + 请求数（单位为 次）
// + 响应时间（单位为ms）
// + 状态码 2xx 汇总及各 2 开头状态码明细（单位为 个）
// + 状态码 3xx 汇总及各 3 开头状态码明细（单位为 个）
// + 状态码 4xx 汇总及各 4 开头状态码明细（单位为 个）
// + 状态码 5xx 汇总及各 5 开头状态码明细（单位为 个）
func (c *Client) DescribeEcdnStatistics(request *DescribeEcdnStatisticsRequest) (response *DescribeEcdnStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeEcdnStatisticsRequest()
    }
    response = NewDescribeEcdnStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePurgeQuotaRequest() (request *DescribePurgeQuotaRequest) {
    request = &DescribePurgeQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "DescribePurgeQuota")
    return
}

func NewDescribePurgeQuotaResponse() (response *DescribePurgeQuotaResponse) {
    response = &DescribePurgeQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询刷新接口的用量配额。
func (c *Client) DescribePurgeQuota(request *DescribePurgeQuotaRequest) (response *DescribePurgeQuotaResponse, err error) {
    if request == nil {
        request = NewDescribePurgeQuotaRequest()
    }
    response = NewDescribePurgeQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePurgeTasksRequest() (request *DescribePurgeTasksRequest) {
    request = &DescribePurgeTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "DescribePurgeTasks")
    return
}

func NewDescribePurgeTasksResponse() (response *DescribePurgeTasksResponse) {
    response = &DescribePurgeTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePurgeTasks 用于查询刷新任务提交历史记录及执行进度。
func (c *Client) DescribePurgeTasks(request *DescribePurgeTasksRequest) (response *DescribePurgeTasksResponse, err error) {
    if request == nil {
        request = NewDescribePurgeTasksRequest()
    }
    response = NewDescribePurgeTasksResponse()
    err = c.Send(request, response)
    return
}

func NewPurgePathCacheRequest() (request *PurgePathCacheRequest) {
    request = &PurgePathCacheRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "PurgePathCache")
    return
}

func NewPurgePathCacheResponse() (response *PurgePathCacheResponse) {
    response = &PurgePathCacheResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PurgeUrlsCache 用于批量刷新目录缓存，一次提交将返回一个刷新任务id。
func (c *Client) PurgePathCache(request *PurgePathCacheRequest) (response *PurgePathCacheResponse, err error) {
    if request == nil {
        request = NewPurgePathCacheRequest()
    }
    response = NewPurgePathCacheResponse()
    err = c.Send(request, response)
    return
}

func NewPurgeUrlsCacheRequest() (request *PurgeUrlsCacheRequest) {
    request = &PurgeUrlsCacheRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "PurgeUrlsCache")
    return
}

func NewPurgeUrlsCacheResponse() (response *PurgeUrlsCacheResponse) {
    response = &PurgeUrlsCacheResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PurgeUrlsCache 用于批量刷新Url，一次提交将返回一个刷新任务id。
func (c *Client) PurgeUrlsCache(request *PurgeUrlsCacheRequest) (response *PurgeUrlsCacheResponse, err error) {
    if request == nil {
        request = NewPurgeUrlsCacheRequest()
    }
    response = NewPurgeUrlsCacheResponse()
    err = c.Send(request, response)
    return
}

func NewStartEcdnDomainRequest() (request *StartEcdnDomainRequest) {
    request = &StartEcdnDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "StartEcdnDomain")
    return
}

func NewStartEcdnDomainResponse() (response *StartEcdnDomainResponse) {
    response = &StartEcdnDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（StartEcdnDomain）用于启用加速域名，待启用域名必须处于已下线状态。
func (c *Client) StartEcdnDomain(request *StartEcdnDomainRequest) (response *StartEcdnDomainResponse, err error) {
    if request == nil {
        request = NewStartEcdnDomainRequest()
    }
    response = NewStartEcdnDomainResponse()
    err = c.Send(request, response)
    return
}

func NewStopEcdnDomainRequest() (request *StopEcdnDomainRequest) {
    request = &StopEcdnDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "StopEcdnDomain")
    return
}

func NewStopEcdnDomainResponse() (response *StopEcdnDomainResponse) {
    response = &StopEcdnDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（StopCdnDomain）用于停止加速域名，待停用加速域名必须处于已上线或部署中状态。
func (c *Client) StopEcdnDomain(request *StopEcdnDomainRequest) (response *StopEcdnDomainResponse, err error) {
    if request == nil {
        request = NewStopEcdnDomainRequest()
    }
    response = NewStopEcdnDomainResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateDomainConfigRequest() (request *UpdateDomainConfigRequest) {
    request = &UpdateDomainConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "UpdateDomainConfig")
    return
}

func NewUpdateDomainConfigResponse() (response *UpdateDomainConfigResponse) {
    response = &UpdateDomainConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（UpdateDomainConfig）用于更新ECDN加速域名配置信息。
// 注意：如果需要更新复杂类型的配置项，必须传递整个对象的所有属性，未传递的属性将使用默认值。建议通过查询接口获取配置属性后，直接修改后传递给本接口。Https配置由于证书的特殊性，更新时不用传递证书和密钥字段。
func (c *Client) UpdateDomainConfig(request *UpdateDomainConfigRequest) (response *UpdateDomainConfigResponse, err error) {
    if request == nil {
        request = NewUpdateDomainConfigRequest()
    }
    response = NewUpdateDomainConfigResponse()
    err = c.Send(request, response)
    return
}
