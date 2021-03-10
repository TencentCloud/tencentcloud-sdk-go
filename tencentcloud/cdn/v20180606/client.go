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
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-06-06"

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


func NewAddCdnDomainRequest() (request *AddCdnDomainRequest) {
    request = &AddCdnDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "AddCdnDomain")
    return
}

func NewAddCdnDomainResponse() (response *AddCdnDomainResponse) {
    response = &AddCdnDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddCdnDomain 用于新增内容分发网络加速域名。
func (c *Client) AddCdnDomain(request *AddCdnDomainRequest) (response *AddCdnDomainResponse, err error) {
    if request == nil {
        request = NewAddCdnDomainRequest()
    }
    response = NewAddCdnDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClsLogTopicRequest() (request *CreateClsLogTopicRequest) {
    request = &CreateClsLogTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "CreateClsLogTopic")
    return
}

func NewCreateClsLogTopicResponse() (response *CreateClsLogTopicResponse) {
    response = &CreateClsLogTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateClsLogTopic 用于创建日志主题。注意：一个日志集下至多可创建10个日志主题。
func (c *Client) CreateClsLogTopic(request *CreateClsLogTopicRequest) (response *CreateClsLogTopicResponse, err error) {
    if request == nil {
        request = NewCreateClsLogTopicRequest()
    }
    response = NewCreateClsLogTopicResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDiagnoseUrlRequest() (request *CreateDiagnoseUrlRequest) {
    request = &CreateDiagnoseUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "CreateDiagnoseUrl")
    return
}

func NewCreateDiagnoseUrlResponse() (response *CreateDiagnoseUrlResponse) {
    response = &CreateDiagnoseUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDiagnoseUrl 用于添加域名诊断任务URL
func (c *Client) CreateDiagnoseUrl(request *CreateDiagnoseUrlRequest) (response *CreateDiagnoseUrlResponse, err error) {
    if request == nil {
        request = NewCreateDiagnoseUrlRequest()
    }
    response = NewCreateDiagnoseUrlResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEdgePackTaskRequest() (request *CreateEdgePackTaskRequest) {
    request = &CreateEdgePackTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "CreateEdgePackTask")
    return
}

func NewCreateEdgePackTaskResponse() (response *CreateEdgePackTaskResponse) {
    response = &CreateEdgePackTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 动态打包任务提交接口
func (c *Client) CreateEdgePackTask(request *CreateEdgePackTaskRequest) (response *CreateEdgePackTaskResponse, err error) {
    if request == nil {
        request = NewCreateEdgePackTaskRequest()
    }
    response = NewCreateEdgePackTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateScdnFailedLogTaskRequest() (request *CreateScdnFailedLogTaskRequest) {
    request = &CreateScdnFailedLogTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "CreateScdnFailedLogTask")
    return
}

func NewCreateScdnFailedLogTaskResponse() (response *CreateScdnFailedLogTaskResponse) {
    response = &CreateScdnFailedLogTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateScdnFailedLogTask 用于重试创建失败的事件日志任务
func (c *Client) CreateScdnFailedLogTask(request *CreateScdnFailedLogTaskRequest) (response *CreateScdnFailedLogTaskResponse, err error) {
    if request == nil {
        request = NewCreateScdnFailedLogTaskRequest()
    }
    response = NewCreateScdnFailedLogTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateScdnLogTaskRequest() (request *CreateScdnLogTaskRequest) {
    request = &CreateScdnLogTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "CreateScdnLogTask")
    return
}

func NewCreateScdnLogTaskResponse() (response *CreateScdnLogTaskResponse) {
    response = &CreateScdnLogTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateScdnLogTask 用于创建事件日志任务
func (c *Client) CreateScdnLogTask(request *CreateScdnLogTaskRequest) (response *CreateScdnLogTaskResponse, err error) {
    if request == nil {
        request = NewCreateScdnLogTaskRequest()
    }
    response = NewCreateScdnLogTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVerifyRecordRequest() (request *CreateVerifyRecordRequest) {
    request = &CreateVerifyRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "CreateVerifyRecord")
    return
}

func NewCreateVerifyRecordResponse() (response *CreateVerifyRecordResponse) {
    response = &CreateVerifyRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 生成一条子域名解析，提示客户添加到域名解析上，用于泛域名及域名取回校验归属权
func (c *Client) CreateVerifyRecord(request *CreateVerifyRecordRequest) (response *CreateVerifyRecordResponse, err error) {
    if request == nil {
        request = NewCreateVerifyRecordRequest()
    }
    response = NewCreateVerifyRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCdnDomainRequest() (request *DeleteCdnDomainRequest) {
    request = &DeleteCdnDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "DeleteCdnDomain")
    return
}

func NewDeleteCdnDomainResponse() (response *DeleteCdnDomainResponse) {
    response = &DeleteCdnDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCdnDomain 用于删除指定加速域名
func (c *Client) DeleteCdnDomain(request *DeleteCdnDomainRequest) (response *DeleteCdnDomainResponse, err error) {
    if request == nil {
        request = NewDeleteCdnDomainRequest()
    }
    response = NewDeleteCdnDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClsLogTopicRequest() (request *DeleteClsLogTopicRequest) {
    request = &DeleteClsLogTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "DeleteClsLogTopic")
    return
}

func NewDeleteClsLogTopicResponse() (response *DeleteClsLogTopicResponse) {
    response = &DeleteClsLogTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteClsLogTopic 用于删除日志主题。注意：删除后，所有该日志主题下绑定域名的日志将不再继续投递至该主题，已经投递的日志将会被全部清空。生效时间约为 5~15 分钟。
func (c *Client) DeleteClsLogTopic(request *DeleteClsLogTopicRequest) (response *DeleteClsLogTopicResponse, err error) {
    if request == nil {
        request = NewDeleteClsLogTopicRequest()
    }
    response = NewDeleteClsLogTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteScdnDomainRequest() (request *DeleteScdnDomainRequest) {
    request = &DeleteScdnDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "DeleteScdnDomain")
    return
}

func NewDeleteScdnDomainResponse() (response *DeleteScdnDomainResponse) {
    response = &DeleteScdnDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除SCDN域名
func (c *Client) DeleteScdnDomain(request *DeleteScdnDomainRequest) (response *DeleteScdnDomainResponse, err error) {
    if request == nil {
        request = NewDeleteScdnDomainRequest()
    }
    response = NewDeleteScdnDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillingDataRequest() (request *DescribeBillingDataRequest) {
    request = &DescribeBillingDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeBillingData")
    return
}

func NewDescribeBillingDataResponse() (response *DescribeBillingDataResponse) {
    response = &DescribeBillingDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBillingData 用于查询实际计费数据明细。
func (c *Client) DescribeBillingData(request *DescribeBillingDataRequest) (response *DescribeBillingDataResponse, err error) {
    if request == nil {
        request = NewDescribeBillingDataRequest()
    }
    response = NewDescribeBillingDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCdnDataRequest() (request *DescribeCdnDataRequest) {
    request = &DescribeCdnDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeCdnData")
    return
}

func NewDescribeCdnDataResponse() (response *DescribeCdnDataResponse) {
    response = &DescribeCdnDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCdnData 用于查询 CDN 实时访问监控数据，支持以下指标查询：
// 
// + 流量（单位为 byte）
// + 带宽（单位为 bps）
// + 请求数（单位为 次）
// + 流量命中率（单位为 %，小数点后保留两位）
// + 状态码 2xx 汇总及各 2 开头状态码明细（单位为 个）
// + 状态码 3xx 汇总及各 3 开头状态码明细（单位为 个）
// + 状态码 4xx 汇总及各 4 开头状态码明细（单位为 个）
// + 状态码 5xx 汇总及各 5 开头状态码明细（单位为 个）
func (c *Client) DescribeCdnData(request *DescribeCdnDataRequest) (response *DescribeCdnDataResponse, err error) {
    if request == nil {
        request = NewDescribeCdnDataRequest()
    }
    response = NewDescribeCdnDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCdnDomainLogsRequest() (request *DescribeCdnDomainLogsRequest) {
    request = &DescribeCdnDomainLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeCdnDomainLogs")
    return
}

func NewDescribeCdnDomainLogsResponse() (response *DescribeCdnDomainLogsResponse) {
    response = &DescribeCdnDomainLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCdnDomainLogs 用于查询访问日志下载地址，仅支持 30 天以内的境内、境外访问日志下载链接查询。
func (c *Client) DescribeCdnDomainLogs(request *DescribeCdnDomainLogsRequest) (response *DescribeCdnDomainLogsResponse, err error) {
    if request == nil {
        request = NewDescribeCdnDomainLogsRequest()
    }
    response = NewDescribeCdnDomainLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCdnIpRequest() (request *DescribeCdnIpRequest) {
    request = &DescribeCdnIpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeCdnIp")
    return
}

func NewDescribeCdnIpResponse() (response *DescribeCdnIpResponse) {
    response = &DescribeCdnIpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCdnIp 用于查询 CDN IP 归属。
// （注意：此接口请求频率限制以 CDN 侧限制为准：200次/10分钟）  
func (c *Client) DescribeCdnIp(request *DescribeCdnIpRequest) (response *DescribeCdnIpResponse, err error) {
    if request == nil {
        request = NewDescribeCdnIpRequest()
    }
    response = NewDescribeCdnIpResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCdnOriginIpRequest() (request *DescribeCdnOriginIpRequest) {
    request = &DescribeCdnOriginIpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeCdnOriginIp")
    return
}

func NewDescribeCdnOriginIpResponse() (response *DescribeCdnOriginIpResponse) {
    response = &DescribeCdnOriginIpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeCdnOriginIp）用于查询 CDN 回源节点的IP信息。（注：此接口即将下线，不再进行维护，请通过DescribeIpStatus 接口进行查询）
func (c *Client) DescribeCdnOriginIp(request *DescribeCdnOriginIpRequest) (response *DescribeCdnOriginIpResponse, err error) {
    if request == nil {
        request = NewDescribeCdnOriginIpRequest()
    }
    response = NewDescribeCdnOriginIpResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCertDomainsRequest() (request *DescribeCertDomainsRequest) {
    request = &DescribeCertDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeCertDomains")
    return
}

func NewDescribeCertDomainsResponse() (response *DescribeCertDomainsResponse) {
    response = &DescribeCertDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCertDomains 用于校验SSL证书并提取证书中包含的域名。
func (c *Client) DescribeCertDomains(request *DescribeCertDomainsRequest) (response *DescribeCertDomainsResponse, err error) {
    if request == nil {
        request = NewDescribeCertDomainsRequest()
    }
    response = NewDescribeCertDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDiagnoseReportRequest() (request *DescribeDiagnoseReportRequest) {
    request = &DescribeDiagnoseReportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeDiagnoseReport")
    return
}

func NewDescribeDiagnoseReportResponse() (response *DescribeDiagnoseReportResponse) {
    response = &DescribeDiagnoseReportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDiagnoseReport 用于获取指定报告id的内容
func (c *Client) DescribeDiagnoseReport(request *DescribeDiagnoseReportRequest) (response *DescribeDiagnoseReportResponse, err error) {
    if request == nil {
        request = NewDescribeDiagnoseReportRequest()
    }
    response = NewDescribeDiagnoseReportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDistrictIspDataRequest() (request *DescribeDistrictIspDataRequest) {
    request = &DescribeDistrictIspDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeDistrictIspData")
    return
}

func NewDescribeDistrictIspDataResponse() (response *DescribeDistrictIspDataResponse) {
    response = &DescribeDistrictIspDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询指定域名的区域、运营商明细数据
// 注意事项：接口尚未全量开放，未在内测名单中的账号不支持调用
func (c *Client) DescribeDistrictIspData(request *DescribeDistrictIspDataRequest) (response *DescribeDistrictIspDataResponse, err error) {
    if request == nil {
        request = NewDescribeDistrictIspDataRequest()
    }
    response = NewDescribeDistrictIspDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainsRequest() (request *DescribeDomainsRequest) {
    request = &DescribeDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeDomains")
    return
}

func NewDescribeDomainsResponse() (response *DescribeDomainsResponse) {
    response = &DescribeDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDomains 用于查询内容分发网络加速域名（含境内、境外）基本配置信息，包括项目ID、服务状态，业务类型、创建时间、更新时间等信息。
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
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeDomainsConfig")
    return
}

func NewDescribeDomainsConfigResponse() (response *DescribeDomainsConfigResponse) {
    response = &DescribeDomainsConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDomainsConfig 用于查询内容分发网络加速域名（含境内、境外）的所有配置信息。
func (c *Client) DescribeDomainsConfig(request *DescribeDomainsConfigRequest) (response *DescribeDomainsConfigResponse, err error) {
    if request == nil {
        request = NewDescribeDomainsConfigRequest()
    }
    response = NewDescribeDomainsConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageConfigRequest() (request *DescribeImageConfigRequest) {
    request = &DescribeImageConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeImageConfig")
    return
}

func NewDescribeImageConfigResponse() (response *DescribeImageConfigResponse) {
    response = &DescribeImageConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeImageConfig 用于获取域名图片优化的当前配置，支持Webp、TPG 和 Guetzli。 
func (c *Client) DescribeImageConfig(request *DescribeImageConfigRequest) (response *DescribeImageConfigResponse, err error) {
    if request == nil {
        request = NewDescribeImageConfigRequest()
    }
    response = NewDescribeImageConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIpStatusRequest() (request *DescribeIpStatusRequest) {
    request = &DescribeIpStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeIpStatus")
    return
}

func NewDescribeIpStatusResponse() (response *DescribeIpStatusResponse) {
    response = &DescribeIpStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIpStatus 用于查询域名所在加速平台的边缘节点、回源节点明细。注意事项：边缘节点（edge）尚未全量开放，未在内测名单中的账号不支持调用
func (c *Client) DescribeIpStatus(request *DescribeIpStatusRequest) (response *DescribeIpStatusResponse, err error) {
    if request == nil {
        request = NewDescribeIpStatusRequest()
    }
    response = NewDescribeIpStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIpVisitRequest() (request *DescribeIpVisitRequest) {
    request = &DescribeIpVisitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeIpVisit")
    return
}

func NewDescribeIpVisitResponse() (response *DescribeIpVisitResponse) {
    response = &DescribeIpVisitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIpVisit 用于查询 5 分钟活跃用户数，及日活跃用户数明细
// 
// + 5 分钟活跃用户数：根据日志中客户端 IP，5 分钟粒度去重统计
// + 日活跃用户数：根据日志中客户端 IP，按天粒度去重统计
func (c *Client) DescribeIpVisit(request *DescribeIpVisitRequest) (response *DescribeIpVisitResponse, err error) {
    if request == nil {
        request = NewDescribeIpVisitRequest()
    }
    response = NewDescribeIpVisitResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMapInfoRequest() (request *DescribeMapInfoRequest) {
    request = &DescribeMapInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeMapInfo")
    return
}

func NewDescribeMapInfoResponse() (response *DescribeMapInfoResponse) {
    response = &DescribeMapInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMapInfo 用于查询省份对应的 ID，运营商对应的 ID 信息。
func (c *Client) DescribeMapInfo(request *DescribeMapInfoRequest) (response *DescribeMapInfoResponse, err error) {
    if request == nil {
        request = NewDescribeMapInfoRequest()
    }
    response = NewDescribeMapInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOriginDataRequest() (request *DescribeOriginDataRequest) {
    request = &DescribeOriginDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeOriginData")
    return
}

func NewDescribeOriginDataResponse() (response *DescribeOriginDataResponse) {
    response = &DescribeOriginDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOriginData 用于查询 CDN 实时回源监控数据，支持以下指标查询：
// 
// + 回源流量（单位为 byte）
// + 回源带宽（单位为 bps）
// + 回源请求数（单位为 次）
// + 回源失败请求数（单位为 次）
// + 回源失败率（单位为 %，小数点后保留两位）
// + 回源状态码 2xx 汇总及各 2 开头回源状态码明细（单位为 个）
// + 回源状态码 3xx 汇总及各 3 开头回源状态码明细（单位为 个）
// + 回源状态码 4xx 汇总及各 4 开头回源状态码明细（单位为 个）
// + 回源状态码 5xx 汇总及各 5 开头回源状态码明细（单位为 个）
func (c *Client) DescribeOriginData(request *DescribeOriginDataRequest) (response *DescribeOriginDataResponse, err error) {
    if request == nil {
        request = NewDescribeOriginDataRequest()
    }
    response = NewDescribeOriginDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePayTypeRequest() (request *DescribePayTypeRequest) {
    request = &DescribePayTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "DescribePayType")
    return
}

func NewDescribePayTypeResponse() (response *DescribePayTypeResponse) {
    response = &DescribePayTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePayType 用于查询用户的计费类型，计费周期等信息。
func (c *Client) DescribePayType(request *DescribePayTypeRequest) (response *DescribePayTypeResponse, err error) {
    if request == nil {
        request = NewDescribePayTypeRequest()
    }
    response = NewDescribePayTypeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePurgeQuotaRequest() (request *DescribePurgeQuotaRequest) {
    request = &DescribePurgeQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "DescribePurgeQuota")
    return
}

func NewDescribePurgeQuotaResponse() (response *DescribePurgeQuotaResponse) {
    response = &DescribePurgeQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePurgeQuota 用于查询账户刷新配额和每日可用量。
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
    request.Init().WithApiInfo("cdn", APIVersion, "DescribePurgeTasks")
    return
}

func NewDescribePurgeTasksResponse() (response *DescribePurgeTasksResponse) {
    response = &DescribePurgeTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePurgeTasks 用于查询提交的 URL 刷新、目录刷新记录及执行进度，通过 PurgePathCache 与 PurgeUrlsCache 接口提交的任务均可通过此接口进行查询。
func (c *Client) DescribePurgeTasks(request *DescribePurgeTasksRequest) (response *DescribePurgeTasksResponse, err error) {
    if request == nil {
        request = NewDescribePurgeTasksRequest()
    }
    response = NewDescribePurgeTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePushQuotaRequest() (request *DescribePushQuotaRequest) {
    request = &DescribePushQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "DescribePushQuota")
    return
}

func NewDescribePushQuotaResponse() (response *DescribePushQuotaResponse) {
    response = &DescribePushQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePushQuota  用于查询预热配额和每日可用量。
func (c *Client) DescribePushQuota(request *DescribePushQuotaRequest) (response *DescribePushQuotaResponse, err error) {
    if request == nil {
        request = NewDescribePushQuotaRequest()
    }
    response = NewDescribePushQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePushTasksRequest() (request *DescribePushTasksRequest) {
    request = &DescribePushTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "DescribePushTasks")
    return
}

func NewDescribePushTasksResponse() (response *DescribePushTasksResponse) {
    response = &DescribePushTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePushTasks  用于查询预热任务提交历史记录及执行进度。
// 接口灰度中，暂未全量开放，敬请期待。
func (c *Client) DescribePushTasks(request *DescribePushTasksRequest) (response *DescribePushTasksResponse, err error) {
    if request == nil {
        request = NewDescribePushTasksRequest()
    }
    response = NewDescribePushTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReportDataRequest() (request *DescribeReportDataRequest) {
    request = &DescribeReportDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeReportData")
    return
}

func NewDescribeReportDataResponse() (response *DescribeReportDataResponse) {
    response = &DescribeReportDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeReportData 用于查询域名/项目维度的日/周/月报表数据。
func (c *Client) DescribeReportData(request *DescribeReportDataRequest) (response *DescribeReportDataResponse, err error) {
    if request == nil {
        request = NewDescribeReportDataRequest()
    }
    response = NewDescribeReportDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScdnConfigRequest() (request *DescribeScdnConfigRequest) {
    request = &DescribeScdnConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeScdnConfig")
    return
}

func NewDescribeScdnConfigResponse() (response *DescribeScdnConfigResponse) {
    response = &DescribeScdnConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScdnConfig 用于查询指定 SCDN 加速域名的安全相关配置
func (c *Client) DescribeScdnConfig(request *DescribeScdnConfigRequest) (response *DescribeScdnConfigResponse, err error) {
    if request == nil {
        request = NewDescribeScdnConfigRequest()
    }
    response = NewDescribeScdnConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScdnTopDataRequest() (request *DescribeScdnTopDataRequest) {
    request = &DescribeScdnTopDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeScdnTopData")
    return
}

func NewDescribeScdnTopDataResponse() (response *DescribeScdnTopDataResponse) {
    response = &DescribeScdnTopDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取SCDN的Top数据
func (c *Client) DescribeScdnTopData(request *DescribeScdnTopDataRequest) (response *DescribeScdnTopDataResponse, err error) {
    if request == nil {
        request = NewDescribeScdnTopDataRequest()
    }
    response = NewDescribeScdnTopDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTrafficPackagesRequest() (request *DescribeTrafficPackagesRequest) {
    request = &DescribeTrafficPackagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeTrafficPackages")
    return
}

func NewDescribeTrafficPackagesResponse() (response *DescribeTrafficPackagesResponse) {
    response = &DescribeTrafficPackagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTrafficPackages 用于查询中国境内 CDN 流量包详情。
func (c *Client) DescribeTrafficPackages(request *DescribeTrafficPackagesRequest) (response *DescribeTrafficPackagesResponse, err error) {
    if request == nil {
        request = NewDescribeTrafficPackagesRequest()
    }
    response = NewDescribeTrafficPackagesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUrlViolationsRequest() (request *DescribeUrlViolationsRequest) {
    request = &DescribeUrlViolationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeUrlViolations")
    return
}

func NewDescribeUrlViolationsResponse() (response *DescribeUrlViolationsResponse) {
    response = &DescribeUrlViolationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUrlViolations 用于查询被 CDN 系统扫描到的域名违规 URL 列表及当前状态。
// 对应内容分发网络控制台【图片鉴黄】页面。
func (c *Client) DescribeUrlViolations(request *DescribeUrlViolationsRequest) (response *DescribeUrlViolationsResponse, err error) {
    if request == nil {
        request = NewDescribeUrlViolationsRequest()
    }
    response = NewDescribeUrlViolationsResponse()
    err = c.Send(request, response)
    return
}

func NewDisableCachesRequest() (request *DisableCachesRequest) {
    request = &DisableCachesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "DisableCaches")
    return
}

func NewDisableCachesResponse() (response *DisableCachesResponse) {
    response = &DisableCachesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisableCaches 用于禁用 CDN 上指定 URL 的访问，禁用完成后，中国境内访问会直接返回 403。（接口尚在内测中，暂未全量开放使用）
func (c *Client) DisableCaches(request *DisableCachesRequest) (response *DisableCachesResponse, err error) {
    if request == nil {
        request = NewDisableCachesRequest()
    }
    response = NewDisableCachesResponse()
    err = c.Send(request, response)
    return
}

func NewDisableClsLogTopicRequest() (request *DisableClsLogTopicRequest) {
    request = &DisableClsLogTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "DisableClsLogTopic")
    return
}

func NewDisableClsLogTopicResponse() (response *DisableClsLogTopicResponse) {
    response = &DisableClsLogTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisableClsLogTopic 用于停止日志主题投递。注意：停止后，所有绑定该日志主题域名的日志将不再继续投递至该主题，已经投递的日志将会继续保留。生效时间约为 5~15 分钟。
func (c *Client) DisableClsLogTopic(request *DisableClsLogTopicRequest) (response *DisableClsLogTopicResponse, err error) {
    if request == nil {
        request = NewDisableClsLogTopicRequest()
    }
    response = NewDisableClsLogTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDuplicateDomainConfigRequest() (request *DuplicateDomainConfigRequest) {
    request = &DuplicateDomainConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "DuplicateDomainConfig")
    return
}

func NewDuplicateDomainConfigResponse() (response *DuplicateDomainConfigResponse) {
    response = &DuplicateDomainConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 拷贝参考域名的配置至新域名。暂不支持自有证书以及定制化配置
func (c *Client) DuplicateDomainConfig(request *DuplicateDomainConfigRequest) (response *DuplicateDomainConfigResponse, err error) {
    if request == nil {
        request = NewDuplicateDomainConfigRequest()
    }
    response = NewDuplicateDomainConfigResponse()
    err = c.Send(request, response)
    return
}

func NewEnableCachesRequest() (request *EnableCachesRequest) {
    request = &EnableCachesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "EnableCaches")
    return
}

func NewEnableCachesResponse() (response *EnableCachesResponse) {
    response = &EnableCachesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EnableCaches 用于解禁手工封禁的 URL，解禁成功后，全网生效时间约 5~10 分钟。（接口尚在内测中，暂未全量开放使用）
func (c *Client) EnableCaches(request *EnableCachesRequest) (response *EnableCachesResponse, err error) {
    if request == nil {
        request = NewEnableCachesRequest()
    }
    response = NewEnableCachesResponse()
    err = c.Send(request, response)
    return
}

func NewEnableClsLogTopicRequest() (request *EnableClsLogTopicRequest) {
    request = &EnableClsLogTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "EnableClsLogTopic")
    return
}

func NewEnableClsLogTopicResponse() (response *EnableClsLogTopicResponse) {
    response = &EnableClsLogTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EnableClsLogTopic 用于启动日志主题投递。注意：启动后，所有绑定该日志主题域名的日志将继续投递至该主题。生效时间约为 5~15 分钟。
func (c *Client) EnableClsLogTopic(request *EnableClsLogTopicRequest) (response *EnableClsLogTopicResponse, err error) {
    if request == nil {
        request = NewEnableClsLogTopicRequest()
    }
    response = NewEnableClsLogTopicResponse()
    err = c.Send(request, response)
    return
}

func NewGetDisableRecordsRequest() (request *GetDisableRecordsRequest) {
    request = &GetDisableRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "GetDisableRecords")
    return
}

func NewGetDisableRecordsResponse() (response *GetDisableRecordsResponse) {
    response = &GetDisableRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetDisableRecords 用于查询资源禁用历史，及 URL 当前状态。（接口尚在内测中，暂未全量开放使用）
func (c *Client) GetDisableRecords(request *GetDisableRecordsRequest) (response *GetDisableRecordsResponse, err error) {
    if request == nil {
        request = NewGetDisableRecordsRequest()
    }
    response = NewGetDisableRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewListClsLogTopicsRequest() (request *ListClsLogTopicsRequest) {
    request = &ListClsLogTopicsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "ListClsLogTopics")
    return
}

func NewListClsLogTopicsResponse() (response *ListClsLogTopicsResponse) {
    response = &ListClsLogTopicsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListClsLogTopics 用于显示日志主题列表。注意：一个日志集下至多含10个日志主题。
func (c *Client) ListClsLogTopics(request *ListClsLogTopicsRequest) (response *ListClsLogTopicsResponse, err error) {
    if request == nil {
        request = NewListClsLogTopicsRequest()
    }
    response = NewListClsLogTopicsResponse()
    err = c.Send(request, response)
    return
}

func NewListClsTopicDomainsRequest() (request *ListClsTopicDomainsRequest) {
    request = &ListClsTopicDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "ListClsTopicDomains")
    return
}

func NewListClsTopicDomainsResponse() (response *ListClsTopicDomainsResponse) {
    response = &ListClsTopicDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListClsTopicDomains 用于获取某日志主题下绑定的域名列表。
func (c *Client) ListClsTopicDomains(request *ListClsTopicDomainsRequest) (response *ListClsTopicDomainsResponse, err error) {
    if request == nil {
        request = NewListClsTopicDomainsRequest()
    }
    response = NewListClsTopicDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewListDiagnoseReportRequest() (request *ListDiagnoseReportRequest) {
    request = &ListDiagnoseReportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "ListDiagnoseReport")
    return
}

func NewListDiagnoseReportResponse() (response *ListDiagnoseReportResponse) {
    response = &ListDiagnoseReportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListDiagnoseReport 用于获取用户诊断URL访问后各个子任务的简要详情。
func (c *Client) ListDiagnoseReport(request *ListDiagnoseReportRequest) (response *ListDiagnoseReportResponse, err error) {
    if request == nil {
        request = NewListDiagnoseReportRequest()
    }
    response = NewListDiagnoseReportResponse()
    err = c.Send(request, response)
    return
}

func NewListScdnDomainsRequest() (request *ListScdnDomainsRequest) {
    request = &ListScdnDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "ListScdnDomains")
    return
}

func NewListScdnDomainsResponse() (response *ListScdnDomainsResponse) {
    response = &ListScdnDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListScdnDomains 用于查询 SCDN 安全加速域名列表，及域名基本配置信息
func (c *Client) ListScdnDomains(request *ListScdnDomainsRequest) (response *ListScdnDomainsResponse, err error) {
    if request == nil {
        request = NewListScdnDomainsRequest()
    }
    response = NewListScdnDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewListScdnLogTasksRequest() (request *ListScdnLogTasksRequest) {
    request = &ListScdnLogTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "ListScdnLogTasks")
    return
}

func NewListScdnLogTasksResponse() (response *ListScdnLogTasksResponse) {
    response = &ListScdnLogTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListScdnLogTasks 用于查询SCDN日志下载任务列表,以及展示下载任务基本信息
func (c *Client) ListScdnLogTasks(request *ListScdnLogTasksRequest) (response *ListScdnLogTasksResponse, err error) {
    if request == nil {
        request = NewListScdnLogTasksRequest()
    }
    response = NewListScdnLogTasksResponse()
    err = c.Send(request, response)
    return
}

func NewListTopDataRequest() (request *ListTopDataRequest) {
    request = &ListTopDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "ListTopData")
    return
}

func NewListTopDataResponse() (response *ListTopDataResponse) {
    response = &ListTopDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListTopData 通过入参 Metric 和 Filter 组合不同，可以查询以下排序数据：
// 
// + 依据总流量、总请求数对访问 URL 排序，从大至小返回 TOP 1000 URL
// + 依据总流量、总请求数对客户端省份排序，从大至小返回省份列表
// + 依据总流量、总请求数对客户端运营商排序，从大至小返回运营商列表
// + 依据总流量、峰值带宽、总请求数、平均命中率、2XX/3XX/4XX/5XX 状态码对域名排序，从大至小返回域名列表
// + 依据总回源流量、回源峰值带宽、总回源请求数、平均回源失败率、2XX/3XX/4XX/5XX 回源状态码对域名排序，从大至小返回域名列表
// 
// 注意：仅支持 90 天内数据查询
func (c *Client) ListTopData(request *ListTopDataRequest) (response *ListTopDataResponse, err error) {
    if request == nil {
        request = NewListTopDataRequest()
    }
    response = NewListTopDataResponse()
    err = c.Send(request, response)
    return
}

func NewManageClsTopicDomainsRequest() (request *ManageClsTopicDomainsRequest) {
    request = &ManageClsTopicDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "ManageClsTopicDomains")
    return
}

func NewManageClsTopicDomainsResponse() (response *ManageClsTopicDomainsResponse) {
    response = &ManageClsTopicDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ManageClsTopicDomains 用于管理某日志主题下绑定的域名列表。
func (c *Client) ManageClsTopicDomains(request *ManageClsTopicDomainsRequest) (response *ManageClsTopicDomainsResponse, err error) {
    if request == nil {
        request = NewManageClsTopicDomainsRequest()
    }
    response = NewManageClsTopicDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewPurgePathCacheRequest() (request *PurgePathCacheRequest) {
    request = &PurgePathCacheRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "PurgePathCache")
    return
}

func NewPurgePathCacheResponse() (response *PurgePathCacheResponse) {
    response = &PurgePathCacheResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PurgePathCache 用于批量提交目录刷新，根据域名的加速区域进行对应区域的刷新。
// 默认情况下境内、境外加速区域每日目录刷新额度为各 100 条，每次最多可提交 20 条。
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
    request.Init().WithApiInfo("cdn", APIVersion, "PurgeUrlsCache")
    return
}

func NewPurgeUrlsCacheResponse() (response *PurgeUrlsCacheResponse) {
    response = &PurgeUrlsCacheResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PurgeUrlsCache 用于批量提交 URL 进行刷新，根据 URL 中域名的当前加速区域进行对应区域的刷新。
// 默认情况下境内、境外加速区域每日 URL 刷新额度各为 10000 条，每次最多可提交 1000 条。
func (c *Client) PurgeUrlsCache(request *PurgeUrlsCacheRequest) (response *PurgeUrlsCacheResponse, err error) {
    if request == nil {
        request = NewPurgeUrlsCacheRequest()
    }
    response = NewPurgeUrlsCacheResponse()
    err = c.Send(request, response)
    return
}

func NewPushUrlsCacheRequest() (request *PushUrlsCacheRequest) {
    request = &PushUrlsCacheRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "PushUrlsCache")
    return
}

func NewPushUrlsCacheResponse() (response *PushUrlsCacheResponse) {
    response = &PushUrlsCacheResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PushUrlsCache 用于将指定 URL 资源列表加载至 CDN 节点，支持指定加速区域预热。
// 默认情况下境内、境外每日预热 URL 限额为各 1000 条，每次最多可提交 20 条。
func (c *Client) PushUrlsCache(request *PushUrlsCacheRequest) (response *PushUrlsCacheResponse, err error) {
    if request == nil {
        request = NewPushUrlsCacheRequest()
    }
    response = NewPushUrlsCacheResponse()
    err = c.Send(request, response)
    return
}

func NewSearchClsLogRequest() (request *SearchClsLogRequest) {
    request = &SearchClsLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "SearchClsLog")
    return
}

func NewSearchClsLogResponse() (response *SearchClsLogResponse) {
    response = &SearchClsLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SearchClsLog 用于 CLS 日志检索。支持检索今天，24小时（可选近7中的某一天），近7天的日志数据。
func (c *Client) SearchClsLog(request *SearchClsLogRequest) (response *SearchClsLogResponse, err error) {
    if request == nil {
        request = NewSearchClsLogRequest()
    }
    response = NewSearchClsLogResponse()
    err = c.Send(request, response)
    return
}

func NewStartCdnDomainRequest() (request *StartCdnDomainRequest) {
    request = &StartCdnDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "StartCdnDomain")
    return
}

func NewStartCdnDomainResponse() (response *StartCdnDomainResponse) {
    response = &StartCdnDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartCdnDomain 用于启用已停用域名的加速服务
func (c *Client) StartCdnDomain(request *StartCdnDomainRequest) (response *StartCdnDomainResponse, err error) {
    if request == nil {
        request = NewStartCdnDomainRequest()
    }
    response = NewStartCdnDomainResponse()
    err = c.Send(request, response)
    return
}

func NewStartScdnDomainRequest() (request *StartScdnDomainRequest) {
    request = &StartScdnDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "StartScdnDomain")
    return
}

func NewStartScdnDomainResponse() (response *StartScdnDomainResponse) {
    response = &StartScdnDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartScdnDomain 用于开启域名的安全防护配置
func (c *Client) StartScdnDomain(request *StartScdnDomainRequest) (response *StartScdnDomainResponse, err error) {
    if request == nil {
        request = NewStartScdnDomainRequest()
    }
    response = NewStartScdnDomainResponse()
    err = c.Send(request, response)
    return
}

func NewStopCdnDomainRequest() (request *StopCdnDomainRequest) {
    request = &StopCdnDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "StopCdnDomain")
    return
}

func NewStopCdnDomainResponse() (response *StopCdnDomainResponse) {
    response = &StopCdnDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopCdnDomain 用于停止域名的加速服务。
// 注意：停止加速服务后，访问至加速节点的请求将会直接返回 404。为避免对您的业务造成影响，请在停止加速服务前将解析切走。
func (c *Client) StopCdnDomain(request *StopCdnDomainRequest) (response *StopCdnDomainResponse, err error) {
    if request == nil {
        request = NewStopCdnDomainRequest()
    }
    response = NewStopCdnDomainResponse()
    err = c.Send(request, response)
    return
}

func NewStopScdnDomainRequest() (request *StopScdnDomainRequest) {
    request = &StopScdnDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "StopScdnDomain")
    return
}

func NewStopScdnDomainResponse() (response *StopScdnDomainResponse) {
    response = &StopScdnDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopScdnDomain 用于关闭域名的安全防护配置
func (c *Client) StopScdnDomain(request *StopScdnDomainRequest) (response *StopScdnDomainResponse, err error) {
    if request == nil {
        request = NewStopScdnDomainRequest()
    }
    response = NewStopScdnDomainResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateDomainConfigRequest() (request *UpdateDomainConfigRequest) {
    request = &UpdateDomainConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "UpdateDomainConfig")
    return
}

func NewUpdateDomainConfigResponse() (response *UpdateDomainConfigResponse) {
    response = &UpdateDomainConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateDomainConfig 用于修改内容分发网络加速域名配置信息
// 注意：如果需要更新复杂类型的配置项，必须传递整个对象的所有属性，未传递的属性将使用默认值，建议通过查询接口获取配置属性后，直接修改后传递给本接口。Https配置由于证书的特殊性，更新时不用传递证书和密钥字段。
func (c *Client) UpdateDomainConfig(request *UpdateDomainConfigRequest) (response *UpdateDomainConfigResponse, err error) {
    if request == nil {
        request = NewUpdateDomainConfigRequest()
    }
    response = NewUpdateDomainConfigResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateImageConfigRequest() (request *UpdateImageConfigRequest) {
    request = &UpdateImageConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "UpdateImageConfig")
    return
}

func NewUpdateImageConfigResponse() (response *UpdateImageConfigResponse) {
    response = &UpdateImageConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateImageConfig 用于更新控制台图片优化的相关配置，支持Webp、TPG 和 Guetzli。 
func (c *Client) UpdateImageConfig(request *UpdateImageConfigRequest) (response *UpdateImageConfigResponse, err error) {
    if request == nil {
        request = NewUpdateImageConfigRequest()
    }
    response = NewUpdateImageConfigResponse()
    err = c.Send(request, response)
    return
}

func NewUpdatePayTypeRequest() (request *UpdatePayTypeRequest) {
    request = &UpdatePayTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "UpdatePayType")
    return
}

func NewUpdatePayTypeResponse() (response *UpdatePayTypeResponse) {
    response = &UpdatePayTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(UpdatePayType)用于修改账号计费类型，暂不支持月结用户或子账号修改。
func (c *Client) UpdatePayType(request *UpdatePayTypeRequest) (response *UpdatePayTypeResponse, err error) {
    if request == nil {
        request = NewUpdatePayTypeRequest()
    }
    response = NewUpdatePayTypeResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateScdnDomainRequest() (request *UpdateScdnDomainRequest) {
    request = &UpdateScdnDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "UpdateScdnDomain")
    return
}

func NewUpdateScdnDomainResponse() (response *UpdateScdnDomainResponse) {
    response = &UpdateScdnDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateScdnDomain 用于修改 SCDN 加速域名安全相关配置
func (c *Client) UpdateScdnDomain(request *UpdateScdnDomainRequest) (response *UpdateScdnDomainResponse, err error) {
    if request == nil {
        request = NewUpdateScdnDomainRequest()
    }
    response = NewUpdateScdnDomainResponse()
    err = c.Send(request, response)
    return
}

func NewVerifyDomainRecordRequest() (request *VerifyDomainRecordRequest) {
    request = &VerifyDomainRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cdn", APIVersion, "VerifyDomainRecord")
    return
}

func NewVerifyDomainRecordResponse() (response *VerifyDomainRecordResponse) {
    response = &VerifyDomainRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 验证域名解析值
func (c *Client) VerifyDomainRecord(request *VerifyDomainRecordRequest) (response *VerifyDomainRecordResponse, err error) {
    if request == nil {
        request = NewVerifyDomainRecordRequest()
    }
    response = NewVerifyDomainRecordResponse()
    err = c.Send(request, response)
    return
}
