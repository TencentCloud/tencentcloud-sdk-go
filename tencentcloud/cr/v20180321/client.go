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

// 加入黑名单的客户，将停止拨打。用于：
// 将客户进行黑名单的增加和移除，用于对某些客户阶段性停催。
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

// 用于下载当日催收和回访结果报表。当日23:00后，可获取当日催收结果，次日00:30后，可获取昨日回访结果。
func (c *Client) DownloadReport(request *DownloadReportRequest) (response *DownloadReportResponse, err error) {
    if request == nil {
        request = NewDownloadReportRequest()
    }
    response = NewDownloadReportResponse()
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

// <p>该接口包含上传下列文件：</p>
// <ol style="margin-bottom:10px;">
//   <li>入催文件：用于每天入催文件的上传</li>
//   <li>回访文件：用于每天贷中回访文件的上传</li>
//   <li>还款文件：实时上传当前已还款客户，用于还款客户的实时停催</li>
// </ol>
// 接口返回数据任务ID，支持xlsx、xls、csv、zip格式，文档大小不超过50MB。
func (c *Client) UploadDataFile(request *UploadDataFileRequest) (response *UploadDataFileResponse, err error) {
    if request == nil {
        request = NewUploadDataFileRequest()
    }
    response = NewUploadDataFileResponse()
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
