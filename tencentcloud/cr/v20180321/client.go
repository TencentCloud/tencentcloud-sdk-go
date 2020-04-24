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

package v20180321

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-03-21"

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


func NewApplyBlackListRequest() (request *ApplyBlackListRequest) {
    request = &ApplyBlackListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cr", APIVersion, "ApplyBlackList")
    return
}

func NewApplyBlackListResponse() (response *ApplyBlackListResponse) {
    response = &ApplyBlackListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 提交黑名单后，黑名单中有效期内的号码将停止拨打，适用于到期/逾期提醒、回访场景。
func (c *Client) ApplyBlackList(request *ApplyBlackListRequest) (response *ApplyBlackListResponse, err error) {
    if request == nil {
        request = NewApplyBlackListRequest()
    }
    response = NewApplyBlackListResponse()
    err = c.Send(request, response)
    return
}

func NewApplyCreditAuditRequest() (request *ApplyCreditAuditRequest) {
    request = &ApplyCreditAuditRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cr", APIVersion, "ApplyCreditAudit")
    return
}

func NewApplyCreditAuditResponse() (response *ApplyCreditAuditResponse) {
    response = &ApplyCreditAuditResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 提交信审外呼申请，返回当次请求日期。
func (c *Client) ApplyCreditAudit(request *ApplyCreditAuditRequest) (response *ApplyCreditAuditResponse, err error) {
    if request == nil {
        request = NewApplyCreditAuditRequest()
    }
    response = NewApplyCreditAuditResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCreditResultRequest() (request *DescribeCreditResultRequest) {
    request = &DescribeCreditResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cr", APIVersion, "DescribeCreditResult")
    return
}

func NewDescribeCreditResultResponse() (response *DescribeCreditResultResponse) {
    response = &DescribeCreditResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据信审任务ID和请求日期，获取相关信审结果。
func (c *Client) DescribeCreditResult(request *DescribeCreditResultRequest) (response *DescribeCreditResultResponse, err error) {
    if request == nil {
        request = NewDescribeCreditResultRequest()
    }
    response = NewDescribeCreditResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordsRequest() (request *DescribeRecordsRequest) {
    request = &DescribeRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cr", APIVersion, "DescribeRecords")
    return
}

func NewDescribeRecordsResponse() (response *DescribeRecordsResponse) {
    response = &DescribeRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于获取指定案件的录音地址，次日早上8:00后可查询前日录音。
func (c *Client) DescribeRecords(request *DescribeRecordsRequest) (response *DescribeRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeRecordsRequest()
    }
    response = NewDescribeRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskStatusRequest() (request *DescribeTaskStatusRequest) {
    request = &DescribeTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cr", APIVersion, "DescribeTaskStatus")
    return
}

func NewDescribeTaskStatusResponse() (response *DescribeTaskStatusResponse) {
    response = &DescribeTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据上传文件接口的输出参数DataResId，获取相关上传结果。
func (c *Client) DescribeTaskStatus(request *DescribeTaskStatusRequest) (response *DescribeTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeTaskStatusRequest()
    }
    response = NewDescribeTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDownloadDialogueTextRequest() (request *DownloadDialogueTextRequest) {
    request = &DownloadDialogueTextRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cr", APIVersion, "DownloadDialogueText")
    return
}

func NewDownloadDialogueTextResponse() (response *DownloadDialogueTextResponse) {
    response = &DownloadDialogueTextResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于获取指定案件的对话文本内容，次日早上8:00后可查询前日对话文本内容。
func (c *Client) DownloadDialogueText(request *DownloadDialogueTextRequest) (response *DownloadDialogueTextResponse, err error) {
    if request == nil {
        request = NewDownloadDialogueTextRequest()
    }
    response = NewDownloadDialogueTextResponse()
    err = c.Send(request, response)
    return
}

func NewDownloadRecordListRequest() (request *DownloadRecordListRequest) {
    request = &DownloadRecordListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cr", APIVersion, "DownloadRecordList")
    return
}

func NewDownloadRecordListResponse() (response *DownloadRecordListResponse) {
    response = &DownloadRecordListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// <p>用于获取录音下载链接清单，次日早上8:00后可查询前日录音清单。</p>
// <p>注意：录音清单中的录音下载链接仅次日20:00之前有效，请及时下载。</p>
func (c *Client) DownloadRecordList(request *DownloadRecordListRequest) (response *DownloadRecordListResponse, err error) {
    if request == nil {
        request = NewDownloadRecordListRequest()
    }
    response = NewDownloadRecordListResponse()
    err = c.Send(request, response)
    return
}

func NewDownloadReportRequest() (request *DownloadReportRequest) {
    request = &DownloadReportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cr", APIVersion, "DownloadReport")
    return
}

func NewDownloadReportResponse() (response *DownloadReportResponse) {
    response = &DownloadReportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于下载结果报表。当日23:00后，可获取当日到期/逾期提醒结果，次日00:30后，可获取昨日回访结果。
func (c *Client) DownloadReport(request *DownloadReportRequest) (response *DownloadReportResponse, err error) {
    if request == nil {
        request = NewDownloadReportRequest()
    }
    response = NewDownloadReportResponse()
    err = c.Send(request, response)
    return
}

func NewQueryInstantDataRequest() (request *QueryInstantDataRequest) {
    request = &QueryInstantDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cr", APIVersion, "QueryInstantData")
    return
}

func NewQueryInstantDataResponse() (response *QueryInstantDataResponse) {
    response = &QueryInstantDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 实时数据查询
func (c *Client) QueryInstantData(request *QueryInstantDataRequest) (response *QueryInstantDataResponse, err error) {
    if request == nil {
        request = NewQueryInstantDataRequest()
    }
    response = NewQueryInstantDataResponse()
    err = c.Send(request, response)
    return
}

func NewUploadDataFileRequest() (request *UploadDataFileRequest) {
    request = &UploadDataFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cr", APIVersion, "UploadDataFile")
    return
}

func NewUploadDataFileResponse() (response *UploadDataFileResponse) {
    response = &UploadDataFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 上传文件，接口返回数据任务ID，支持xlsx、xls、csv、zip格式。
func (c *Client) UploadDataFile(request *UploadDataFileRequest) (response *UploadDataFileResponse, err error) {
    if request == nil {
        request = NewUploadDataFileRequest()
    }
    response = NewUploadDataFileResponse()
    err = c.Send(request, response)
    return
}

func NewUploadDataJsonRequest() (request *UploadDataJsonRequest) {
    request = &UploadDataJsonRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cr", APIVersion, "UploadDataJson")
    return
}

func NewUploadDataJsonResponse() (response *UploadDataJsonResponse) {
    response = &UploadDataJsonResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 上传Json格式数据，接口返回数据任务ID
func (c *Client) UploadDataJson(request *UploadDataJsonRequest) (response *UploadDataJsonResponse, err error) {
    if request == nil {
        request = NewUploadDataJsonRequest()
    }
    response = NewUploadDataJsonResponse()
    err = c.Send(request, response)
    return
}

func NewUploadFileRequest() (request *UploadFileRequest) {
    request = &UploadFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cr", APIVersion, "UploadFile")
    return
}

func NewUploadFileResponse() (response *UploadFileResponse) {
    response = &UploadFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 客户通过调用该接口上传需催收文档，格式需为excel格式。接口返回任务ID。
func (c *Client) UploadFile(request *UploadFileRequest) (response *UploadFileResponse, err error) {
    if request == nil {
        request = NewUploadFileRequest()
    }
    response = NewUploadFileResponse()
    err = c.Send(request, response)
    return
}
