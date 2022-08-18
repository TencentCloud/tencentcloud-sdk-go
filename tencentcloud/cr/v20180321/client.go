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
    "context"
    "errors"
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

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
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

// ApplyBlackList
// 提交黑名单后，黑名单中有效期内的号码将停止拨打，适用于到期/逾期提醒、回访场景。
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLYBLACKLISTERROR = "FailedOperation.ApplyBlackListError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
func (c *Client) ApplyBlackList(request *ApplyBlackListRequest) (response *ApplyBlackListResponse, err error) {
    return c.ApplyBlackListWithContext(context.Background(), request)
}

// ApplyBlackList
// 提交黑名单后，黑名单中有效期内的号码将停止拨打，适用于到期/逾期提醒、回访场景。
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLYBLACKLISTERROR = "FailedOperation.ApplyBlackListError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
func (c *Client) ApplyBlackListWithContext(ctx context.Context, request *ApplyBlackListRequest) (response *ApplyBlackListResponse, err error) {
    if request == nil {
        request = NewApplyBlackListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyBlackList require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyBlackListResponse()
    err = c.Send(request, response)
    return
}

func NewApplyBlackListDataRequest() (request *ApplyBlackListDataRequest) {
    request = &ApplyBlackListDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cr", APIVersion, "ApplyBlackListData")
    
    
    return
}

func NewApplyBlackListDataResponse() (response *ApplyBlackListDataResponse) {
    response = &ApplyBlackListDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ApplyBlackListData
// 提交机器人黑名单申请
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLYBLACKLISTDATAERROR = "FailedOperation.ApplyBlackListDataError"
//  FAILEDOPERATION_APPLYBLACKLISTERROR = "FailedOperation.ApplyBlackListError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
func (c *Client) ApplyBlackListData(request *ApplyBlackListDataRequest) (response *ApplyBlackListDataResponse, err error) {
    return c.ApplyBlackListDataWithContext(context.Background(), request)
}

// ApplyBlackListData
// 提交机器人黑名单申请
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLYBLACKLISTDATAERROR = "FailedOperation.ApplyBlackListDataError"
//  FAILEDOPERATION_APPLYBLACKLISTERROR = "FailedOperation.ApplyBlackListError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
func (c *Client) ApplyBlackListDataWithContext(ctx context.Context, request *ApplyBlackListDataRequest) (response *ApplyBlackListDataResponse, err error) {
    if request == nil {
        request = NewApplyBlackListDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyBlackListData require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyBlackListDataResponse()
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

// ApplyCreditAudit
// 提交信审外呼申请，返回当次请求日期。
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLYCREDITAUDITERROR = "FailedOperation.ApplyCreditAuditError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  UNAUTHORIZEDOPERATION_ACCOUNTNOTFOUND = "UnauthorizedOperation.AccountNotFound"
func (c *Client) ApplyCreditAudit(request *ApplyCreditAuditRequest) (response *ApplyCreditAuditResponse, err error) {
    return c.ApplyCreditAuditWithContext(context.Background(), request)
}

// ApplyCreditAudit
// 提交信审外呼申请，返回当次请求日期。
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLYCREDITAUDITERROR = "FailedOperation.ApplyCreditAuditError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  UNAUTHORIZEDOPERATION_ACCOUNTNOTFOUND = "UnauthorizedOperation.AccountNotFound"
func (c *Client) ApplyCreditAuditWithContext(ctx context.Context, request *ApplyCreditAuditRequest) (response *ApplyCreditAuditResponse, err error) {
    if request == nil {
        request = NewApplyCreditAuditRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyCreditAudit require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyCreditAuditResponse()
    err = c.Send(request, response)
    return
}

func NewChangeBotCallStatusRequest() (request *ChangeBotCallStatusRequest) {
    request = &ChangeBotCallStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cr", APIVersion, "ChangeBotCallStatus")
    
    
    return
}

func NewChangeBotCallStatusResponse() (response *ChangeBotCallStatusResponse) {
    response = &ChangeBotCallStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ChangeBotCallStatus
// 更新机器人任务作业状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_CHANGEBOTCALLSTATUSERROR = "FailedOperation.ChangeBotCallStatusError"
//  FAILEDOPERATION_CHANGEBOTSTATUSERROR = "FailedOperation.ChangeBotStatusError"
//  FAILEDOPERATION_QUERYCALLLISTERROR = "FailedOperation.QueryCallListError"
//  FAILEDOPERATION_UPLOADBOTFILE = "FailedOperation.UploadBotFile"
//  FAILEDOPERATION_UPLOADDATAERROR = "FailedOperation.UploadDataError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) ChangeBotCallStatus(request *ChangeBotCallStatusRequest) (response *ChangeBotCallStatusResponse, err error) {
    return c.ChangeBotCallStatusWithContext(context.Background(), request)
}

// ChangeBotCallStatus
// 更新机器人任务作业状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_CHANGEBOTCALLSTATUSERROR = "FailedOperation.ChangeBotCallStatusError"
//  FAILEDOPERATION_CHANGEBOTSTATUSERROR = "FailedOperation.ChangeBotStatusError"
//  FAILEDOPERATION_QUERYCALLLISTERROR = "FailedOperation.QueryCallListError"
//  FAILEDOPERATION_UPLOADBOTFILE = "FailedOperation.UploadBotFile"
//  FAILEDOPERATION_UPLOADDATAERROR = "FailedOperation.UploadDataError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) ChangeBotCallStatusWithContext(ctx context.Context, request *ChangeBotCallStatusRequest) (response *ChangeBotCallStatusResponse, err error) {
    if request == nil {
        request = NewChangeBotCallStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChangeBotCallStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewChangeBotCallStatusResponse()
    err = c.Send(request, response)
    return
}

func NewChangeBotTaskStatusRequest() (request *ChangeBotTaskStatusRequest) {
    request = &ChangeBotTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cr", APIVersion, "ChangeBotTaskStatus")
    
    
    return
}

func NewChangeBotTaskStatusResponse() (response *ChangeBotTaskStatusResponse) {
    response = &ChangeBotTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ChangeBotTaskStatus
// 更新机器人任务状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_CHANGEBOTSTATUSERROR = "FailedOperation.ChangeBotStatusError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) ChangeBotTaskStatus(request *ChangeBotTaskStatusRequest) (response *ChangeBotTaskStatusResponse, err error) {
    return c.ChangeBotTaskStatusWithContext(context.Background(), request)
}

// ChangeBotTaskStatus
// 更新机器人任务状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_CHANGEBOTSTATUSERROR = "FailedOperation.ChangeBotStatusError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) ChangeBotTaskStatusWithContext(ctx context.Context, request *ChangeBotTaskStatusRequest) (response *ChangeBotTaskStatusResponse, err error) {
    if request == nil {
        request = NewChangeBotTaskStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChangeBotTaskStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewChangeBotTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBotTaskRequest() (request *CreateBotTaskRequest) {
    request = &CreateBotTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cr", APIVersion, "CreateBotTask")
    
    
    return
}

func NewCreateBotTaskResponse() (response *CreateBotTaskResponse) {
    response = &CreateBotTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBotTask
// 创建机器人任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEBOTTASKERROR = "FailedOperation.CreateBotTaskError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
func (c *Client) CreateBotTask(request *CreateBotTaskRequest) (response *CreateBotTaskResponse, err error) {
    return c.CreateBotTaskWithContext(context.Background(), request)
}

// CreateBotTask
// 创建机器人任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEBOTTASKERROR = "FailedOperation.CreateBotTaskError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
func (c *Client) CreateBotTaskWithContext(ctx context.Context, request *CreateBotTaskRequest) (response *CreateBotTaskResponse, err error) {
    if request == nil {
        request = NewCreateBotTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBotTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBotTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBotFlowRequest() (request *DescribeBotFlowRequest) {
    request = &DescribeBotFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cr", APIVersion, "DescribeBotFlow")
    
    
    return
}

func NewDescribeBotFlowResponse() (response *DescribeBotFlowResponse) {
    response = &DescribeBotFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBotFlow
// 查询机器人对话流
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEBOTFLOWERROR = "FailedOperation.DescribeBotFlowError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
func (c *Client) DescribeBotFlow(request *DescribeBotFlowRequest) (response *DescribeBotFlowResponse, err error) {
    return c.DescribeBotFlowWithContext(context.Background(), request)
}

// DescribeBotFlow
// 查询机器人对话流
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEBOTFLOWERROR = "FailedOperation.DescribeBotFlowError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
func (c *Client) DescribeBotFlowWithContext(ctx context.Context, request *DescribeBotFlowRequest) (response *DescribeBotFlowResponse, err error) {
    if request == nil {
        request = NewDescribeBotFlowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBotFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBotFlowResponse()
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

// DescribeCreditResult
// 根据信审任务ID和请求日期，获取相关信审结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION_GETCREDITAUDITERROR = "FailedOperation.GetCreditAuditError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  UNAUTHORIZEDOPERATION_ACCOUNTNOTFOUND = "UnauthorizedOperation.AccountNotFound"
func (c *Client) DescribeCreditResult(request *DescribeCreditResultRequest) (response *DescribeCreditResultResponse, err error) {
    return c.DescribeCreditResultWithContext(context.Background(), request)
}

// DescribeCreditResult
// 根据信审任务ID和请求日期，获取相关信审结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION_GETCREDITAUDITERROR = "FailedOperation.GetCreditAuditError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  UNAUTHORIZEDOPERATION_ACCOUNTNOTFOUND = "UnauthorizedOperation.AccountNotFound"
func (c *Client) DescribeCreditResultWithContext(ctx context.Context, request *DescribeCreditResultRequest) (response *DescribeCreditResultResponse, err error) {
    if request == nil {
        request = NewDescribeCreditResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCreditResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCreditResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFileModelRequest() (request *DescribeFileModelRequest) {
    request = &DescribeFileModelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cr", APIVersion, "DescribeFileModel")
    
    
    return
}

func NewDescribeFileModelResponse() (response *DescribeFileModelResponse) {
    response = &DescribeFileModelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFileModel
// 查询机器人文件模板
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEFILEMODELERROR = "FailedOperation.DescribeFileModelError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DescribeFileModel(request *DescribeFileModelRequest) (response *DescribeFileModelResponse, err error) {
    return c.DescribeFileModelWithContext(context.Background(), request)
}

// DescribeFileModel
// 查询机器人文件模板
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEFILEMODELERROR = "FailedOperation.DescribeFileModelError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DescribeFileModelWithContext(ctx context.Context, request *DescribeFileModelRequest) (response *DescribeFileModelResponse, err error) {
    if request == nil {
        request = NewDescribeFileModelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFileModel require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFileModelResponse()
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

// DescribeRecords
// 用于获取指定案件的录音地址，次日早上8:00后可查询前日录音。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBERECORDSERROR = "FailedOperation.DescribeRecordsError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
func (c *Client) DescribeRecords(request *DescribeRecordsRequest) (response *DescribeRecordsResponse, err error) {
    return c.DescribeRecordsWithContext(context.Background(), request)
}

// DescribeRecords
// 用于获取指定案件的录音地址，次日早上8:00后可查询前日录音。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBERECORDSERROR = "FailedOperation.DescribeRecordsError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
func (c *Client) DescribeRecordsWithContext(ctx context.Context, request *DescribeRecordsRequest) (response *DescribeRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeRecordsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecords require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeTaskStatus
// 根据上传文件接口的输出参数DataResId，获取相关上传结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBETASKSTATUSERROR = "FailedOperation.DescribeTaskStatusError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
func (c *Client) DescribeTaskStatus(request *DescribeTaskStatusRequest) (response *DescribeTaskStatusResponse, err error) {
    return c.DescribeTaskStatusWithContext(context.Background(), request)
}

// DescribeTaskStatus
// 根据上传文件接口的输出参数DataResId，获取相关上传结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBETASKSTATUSERROR = "FailedOperation.DescribeTaskStatusError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
func (c *Client) DescribeTaskStatusWithContext(ctx context.Context, request *DescribeTaskStatusRequest) (response *DescribeTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeTaskStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDownloadBotRecordRequest() (request *DownloadBotRecordRequest) {
    request = &DownloadBotRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cr", APIVersion, "DownloadBotRecord")
    
    
    return
}

func NewDownloadBotRecordResponse() (response *DownloadBotRecordResponse) {
    response = &DownloadBotRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DownloadBotRecord
// 下载任务录音与文本，第二天12点后可使用此接口获取对应的录音与文本
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBETASKSTATUSERROR = "FailedOperation.DescribeTaskStatusError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
func (c *Client) DownloadBotRecord(request *DownloadBotRecordRequest) (response *DownloadBotRecordResponse, err error) {
    return c.DownloadBotRecordWithContext(context.Background(), request)
}

// DownloadBotRecord
// 下载任务录音与文本，第二天12点后可使用此接口获取对应的录音与文本
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBETASKSTATUSERROR = "FailedOperation.DescribeTaskStatusError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
func (c *Client) DownloadBotRecordWithContext(ctx context.Context, request *DownloadBotRecordRequest) (response *DownloadBotRecordResponse, err error) {
    if request == nil {
        request = NewDownloadBotRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DownloadBotRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDownloadBotRecordResponse()
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

// DownloadDialogueText
// 用于获取指定案件的对话文本内容，次日早上8:00后可查询前日对话文本内容。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADREPORTERROR = "FailedOperation.DownloadReportError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  RESOURCENOTFOUND_REPORTNOTFOUND = "ResourceNotFound.ReportNotFound"
func (c *Client) DownloadDialogueText(request *DownloadDialogueTextRequest) (response *DownloadDialogueTextResponse, err error) {
    return c.DownloadDialogueTextWithContext(context.Background(), request)
}

// DownloadDialogueText
// 用于获取指定案件的对话文本内容，次日早上8:00后可查询前日对话文本内容。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADREPORTERROR = "FailedOperation.DownloadReportError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  RESOURCENOTFOUND_REPORTNOTFOUND = "ResourceNotFound.ReportNotFound"
func (c *Client) DownloadDialogueTextWithContext(ctx context.Context, request *DownloadDialogueTextRequest) (response *DownloadDialogueTextResponse, err error) {
    if request == nil {
        request = NewDownloadDialogueTextRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DownloadDialogueText require credential")
    }

    request.SetContext(ctx)
    
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

// DownloadRecordList
// <p>用于获取录音下载链接清单，次日早上8:00后可查询前日录音清单。</p>
//
// <p>注意：录音清单中的录音下载链接仅次日20:00之前有效，请及时下载。</p>
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADRECORDLISTERROR = "FailedOperation.DownloadRecordListError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  RESOURCENOTFOUND_RECORDLISTNOTFOUND = "ResourceNotFound.RecordListNotFound"
func (c *Client) DownloadRecordList(request *DownloadRecordListRequest) (response *DownloadRecordListResponse, err error) {
    return c.DownloadRecordListWithContext(context.Background(), request)
}

// DownloadRecordList
// <p>用于获取录音下载链接清单，次日早上8:00后可查询前日录音清单。</p>
//
// <p>注意：录音清单中的录音下载链接仅次日20:00之前有效，请及时下载。</p>
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADRECORDLISTERROR = "FailedOperation.DownloadRecordListError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  RESOURCENOTFOUND_RECORDLISTNOTFOUND = "ResourceNotFound.RecordListNotFound"
func (c *Client) DownloadRecordListWithContext(ctx context.Context, request *DownloadRecordListRequest) (response *DownloadRecordListResponse, err error) {
    if request == nil {
        request = NewDownloadRecordListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DownloadRecordList require credential")
    }

    request.SetContext(ctx)
    
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

// DownloadReport
// 用于下载结果报表。当日23:00后，可获取当日到期/逾期提醒结果，次日00:30后，可获取昨日回访结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADREPORTERROR = "FailedOperation.DownloadReportError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  RESOURCENOTFOUND_COMPANYNOTFOUND = "ResourceNotFound.CompanyNotFound"
//  RESOURCENOTFOUND_REPORTNOTFOUND = "ResourceNotFound.ReportNotFound"
func (c *Client) DownloadReport(request *DownloadReportRequest) (response *DownloadReportResponse, err error) {
    return c.DownloadReportWithContext(context.Background(), request)
}

// DownloadReport
// 用于下载结果报表。当日23:00后，可获取当日到期/逾期提醒结果，次日00:30后，可获取昨日回访结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADREPORTERROR = "FailedOperation.DownloadReportError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  RESOURCENOTFOUND_COMPANYNOTFOUND = "ResourceNotFound.CompanyNotFound"
//  RESOURCENOTFOUND_REPORTNOTFOUND = "ResourceNotFound.ReportNotFound"
func (c *Client) DownloadReportWithContext(ctx context.Context, request *DownloadReportRequest) (response *DownloadReportResponse, err error) {
    if request == nil {
        request = NewDownloadReportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DownloadReport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDownloadReportResponse()
    err = c.Send(request, response)
    return
}

func NewExportBotDataRequest() (request *ExportBotDataRequest) {
    request = &ExportBotDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cr", APIVersion, "ExportBotData")
    
    
    return
}

func NewExportBotDataResponse() (response *ExportBotDataResponse) {
    response = &ExportBotDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExportBotData
// 导出机器人数据
//
// 可能返回的错误码:
//  FAILEDOPERATION_EXPORTBOTDATAERROR = "FailedOperation.ExportBotDataError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) ExportBotData(request *ExportBotDataRequest) (response *ExportBotDataResponse, err error) {
    return c.ExportBotDataWithContext(context.Background(), request)
}

// ExportBotData
// 导出机器人数据
//
// 可能返回的错误码:
//  FAILEDOPERATION_EXPORTBOTDATAERROR = "FailedOperation.ExportBotDataError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) ExportBotDataWithContext(ctx context.Context, request *ExportBotDataRequest) (response *ExportBotDataResponse, err error) {
    if request == nil {
        request = NewExportBotDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportBotData require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportBotDataResponse()
    err = c.Send(request, response)
    return
}

func NewQueryBlackListDataRequest() (request *QueryBlackListDataRequest) {
    request = &QueryBlackListDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cr", APIVersion, "QueryBlackListData")
    
    
    return
}

func NewQueryBlackListDataResponse() (response *QueryBlackListDataResponse) {
    response = &QueryBlackListDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryBlackListData
// 查看黑名单数据列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEBLACKLISTDATAERROR = "FailedOperation.DescribeBlacklistDataError"
//  FAILEDOPERATION_QUERYBLACKLISTDATAERROR = "FailedOperation.QueryBlackListDataError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  RESOURCENOTFOUND_BLACKLISTDATANOTFOUND = "ResourceNotFound.BlacklistDataNotFound"
func (c *Client) QueryBlackListData(request *QueryBlackListDataRequest) (response *QueryBlackListDataResponse, err error) {
    return c.QueryBlackListDataWithContext(context.Background(), request)
}

// QueryBlackListData
// 查看黑名单数据列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEBLACKLISTDATAERROR = "FailedOperation.DescribeBlacklistDataError"
//  FAILEDOPERATION_QUERYBLACKLISTDATAERROR = "FailedOperation.QueryBlackListDataError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  RESOURCENOTFOUND_BLACKLISTDATANOTFOUND = "ResourceNotFound.BlacklistDataNotFound"
func (c *Client) QueryBlackListDataWithContext(ctx context.Context, request *QueryBlackListDataRequest) (response *QueryBlackListDataResponse, err error) {
    if request == nil {
        request = NewQueryBlackListDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryBlackListData require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryBlackListDataResponse()
    err = c.Send(request, response)
    return
}

func NewQueryBotListRequest() (request *QueryBotListRequest) {
    request = &QueryBotListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cr", APIVersion, "QueryBotList")
    
    
    return
}

func NewQueryBotListResponse() (response *QueryBotListResponse) {
    response = &QueryBotListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryBotList
// 查询机器人任务状态列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYBOTLISTERROR = "FailedOperation.QueryBotListError"
func (c *Client) QueryBotList(request *QueryBotListRequest) (response *QueryBotListResponse, err error) {
    return c.QueryBotListWithContext(context.Background(), request)
}

// QueryBotList
// 查询机器人任务状态列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYBOTLISTERROR = "FailedOperation.QueryBotListError"
func (c *Client) QueryBotListWithContext(ctx context.Context, request *QueryBotListRequest) (response *QueryBotListResponse, err error) {
    if request == nil {
        request = NewQueryBotListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryBotList require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryBotListResponse()
    err = c.Send(request, response)
    return
}

func NewQueryCallListRequest() (request *QueryCallListRequest) {
    request = &QueryCallListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cr", APIVersion, "QueryCallList")
    
    
    return
}

func NewQueryCallListResponse() (response *QueryCallListResponse) {
    response = &QueryCallListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryCallList
// 机器人任务查询
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYCALLLISTERROR = "FailedOperation.QueryCallListError"
//  FAILEDOPERATION_UPLOADBOTFILE = "FailedOperation.UploadBotFile"
//  FAILEDOPERATION_UPLOADDATAERROR = "FailedOperation.UploadDataError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) QueryCallList(request *QueryCallListRequest) (response *QueryCallListResponse, err error) {
    return c.QueryCallListWithContext(context.Background(), request)
}

// QueryCallList
// 机器人任务查询
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYCALLLISTERROR = "FailedOperation.QueryCallListError"
//  FAILEDOPERATION_UPLOADBOTFILE = "FailedOperation.UploadBotFile"
//  FAILEDOPERATION_UPLOADDATAERROR = "FailedOperation.UploadDataError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) QueryCallListWithContext(ctx context.Context, request *QueryCallListRequest) (response *QueryCallListResponse, err error) {
    if request == nil {
        request = NewQueryCallListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryCallList require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryCallListResponse()
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

// QueryInstantData
// 实时数据查询
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
func (c *Client) QueryInstantData(request *QueryInstantDataRequest) (response *QueryInstantDataResponse, err error) {
    return c.QueryInstantDataWithContext(context.Background(), request)
}

// QueryInstantData
// 实时数据查询
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
func (c *Client) QueryInstantDataWithContext(ctx context.Context, request *QueryInstantDataRequest) (response *QueryInstantDataResponse, err error) {
    if request == nil {
        request = NewQueryInstantDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryInstantData require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryInstantDataResponse()
    err = c.Send(request, response)
    return
}

func NewQueryProductsRequest() (request *QueryProductsRequest) {
    request = &QueryProductsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cr", APIVersion, "QueryProducts")
    
    
    return
}

func NewQueryProductsResponse() (response *QueryProductsResponse) {
    response = &QueryProductsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryProducts
// 查询产品列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEPRODUCTSERROR = "FailedOperation.DescribeProductsError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
func (c *Client) QueryProducts(request *QueryProductsRequest) (response *QueryProductsResponse, err error) {
    return c.QueryProductsWithContext(context.Background(), request)
}

// QueryProducts
// 查询产品列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEPRODUCTSERROR = "FailedOperation.DescribeProductsError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
func (c *Client) QueryProductsWithContext(ctx context.Context, request *QueryProductsRequest) (response *QueryProductsResponse, err error) {
    if request == nil {
        request = NewQueryProductsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryProducts require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryProductsResponse()
    err = c.Send(request, response)
    return
}

func NewQueryRecordListRequest() (request *QueryRecordListRequest) {
    request = &QueryRecordListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cr", APIVersion, "QueryRecordList")
    
    
    return
}

func NewQueryRecordListResponse() (response *QueryRecordListResponse) {
    response = &QueryRecordListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryRecordList
// 查询录音列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYRECORDLISTERROR = "FailedOperation.QueryRecordListError"
func (c *Client) QueryRecordList(request *QueryRecordListRequest) (response *QueryRecordListResponse, err error) {
    return c.QueryRecordListWithContext(context.Background(), request)
}

// QueryRecordList
// 查询录音列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYRECORDLISTERROR = "FailedOperation.QueryRecordListError"
func (c *Client) QueryRecordListWithContext(ctx context.Context, request *QueryRecordListRequest) (response *QueryRecordListResponse, err error) {
    if request == nil {
        request = NewQueryRecordListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryRecordList require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryRecordListResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateBotTaskRequest() (request *UpdateBotTaskRequest) {
    request = &UpdateBotTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cr", APIVersion, "UpdateBotTask")
    
    
    return
}

func NewUpdateBotTaskResponse() (response *UpdateBotTaskResponse) {
    response = &UpdateBotTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateBotTask
// 更新机器人任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_UPDATEBOTTASKERROR = "FailedOperation.UpdateBotTaskError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
func (c *Client) UpdateBotTask(request *UpdateBotTaskRequest) (response *UpdateBotTaskResponse, err error) {
    return c.UpdateBotTaskWithContext(context.Background(), request)
}

// UpdateBotTask
// 更新机器人任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_UPDATEBOTTASKERROR = "FailedOperation.UpdateBotTaskError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
func (c *Client) UpdateBotTaskWithContext(ctx context.Context, request *UpdateBotTaskRequest) (response *UpdateBotTaskResponse, err error) {
    if request == nil {
        request = NewUpdateBotTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateBotTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateBotTaskResponse()
    err = c.Send(request, response)
    return
}

func NewUploadBotDataRequest() (request *UploadBotDataRequest) {
    request = &UploadBotDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cr", APIVersion, "UploadBotData")
    
    
    return
}

func NewUploadBotDataResponse() (response *UploadBotDataResponse) {
    response = &UploadBotDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UploadBotData
// 上传机器人任务数据
//
// 可能返回的错误码:
//  FAILEDOPERATION_UPLOADBOTFILE = "FailedOperation.UploadBotFile"
//  FAILEDOPERATION_UPLOADDATAERROR = "FailedOperation.UploadDataError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
func (c *Client) UploadBotData(request *UploadBotDataRequest) (response *UploadBotDataResponse, err error) {
    return c.UploadBotDataWithContext(context.Background(), request)
}

// UploadBotData
// 上传机器人任务数据
//
// 可能返回的错误码:
//  FAILEDOPERATION_UPLOADBOTFILE = "FailedOperation.UploadBotFile"
//  FAILEDOPERATION_UPLOADDATAERROR = "FailedOperation.UploadDataError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
func (c *Client) UploadBotDataWithContext(ctx context.Context, request *UploadBotDataRequest) (response *UploadBotDataResponse, err error) {
    if request == nil {
        request = NewUploadBotDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadBotData require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadBotDataResponse()
    err = c.Send(request, response)
    return
}

func NewUploadBotFileRequest() (request *UploadBotFileRequest) {
    request = &UploadBotFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cr", APIVersion, "UploadBotFile")
    
    
    return
}

func NewUploadBotFileResponse() (response *UploadBotFileResponse) {
    response = &UploadBotFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UploadBotFile
// 上传机器人文件
//
// 可能返回的错误码:
//  FAILEDOPERATION_UPLOADBOTFILE = "FailedOperation.UploadBotFile"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
func (c *Client) UploadBotFile(request *UploadBotFileRequest) (response *UploadBotFileResponse, err error) {
    return c.UploadBotFileWithContext(context.Background(), request)
}

// UploadBotFile
// 上传机器人文件
//
// 可能返回的错误码:
//  FAILEDOPERATION_UPLOADBOTFILE = "FailedOperation.UploadBotFile"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
func (c *Client) UploadBotFileWithContext(ctx context.Context, request *UploadBotFileRequest) (response *UploadBotFileResponse, err error) {
    if request == nil {
        request = NewUploadBotFileRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadBotFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadBotFileResponse()
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

// UploadDataFile
// 上传文件，接口返回数据任务ID，支持xlsx、xls、csv、zip格式。
//
// 可能返回的错误码:
//  FAILEDOPERATION_UPLOADDATAERROR = "FailedOperation.UploadDataError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  UNAUTHORIZEDOPERATION_ACCOUNTNOTFOUND = "UnauthorizedOperation.AccountNotFound"
func (c *Client) UploadDataFile(request *UploadDataFileRequest) (response *UploadDataFileResponse, err error) {
    return c.UploadDataFileWithContext(context.Background(), request)
}

// UploadDataFile
// 上传文件，接口返回数据任务ID，支持xlsx、xls、csv、zip格式。
//
// 可能返回的错误码:
//  FAILEDOPERATION_UPLOADDATAERROR = "FailedOperation.UploadDataError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  UNAUTHORIZEDOPERATION_ACCOUNTNOTFOUND = "UnauthorizedOperation.AccountNotFound"
func (c *Client) UploadDataFileWithContext(ctx context.Context, request *UploadDataFileRequest) (response *UploadDataFileResponse, err error) {
    if request == nil {
        request = NewUploadDataFileRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadDataFile require credential")
    }

    request.SetContext(ctx)
    
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

// UploadDataJson
// 上传Json格式数据，接口返回数据任务ID
//
// 可能返回的错误码:
//  FAILEDOPERATION_UPLOADDATAERROR = "FailedOperation.UploadDataError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
func (c *Client) UploadDataJson(request *UploadDataJsonRequest) (response *UploadDataJsonResponse, err error) {
    return c.UploadDataJsonWithContext(context.Background(), request)
}

// UploadDataJson
// 上传Json格式数据，接口返回数据任务ID
//
// 可能返回的错误码:
//  FAILEDOPERATION_UPLOADDATAERROR = "FailedOperation.UploadDataError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
func (c *Client) UploadDataJsonWithContext(ctx context.Context, request *UploadDataJsonRequest) (response *UploadDataJsonResponse, err error) {
    if request == nil {
        request = NewUploadDataJsonRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadDataJson require credential")
    }

    request.SetContext(ctx)
    
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

// UploadFile
// 客户通过调用该接口上传需催收文档，格式需为excel格式。接口返回任务ID。
//
// 可能返回的错误码:
//  FAILEDOPERATION_UPLOADDATAERROR = "FailedOperation.UploadDataError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  RESOURCENOTFOUND_COMPANYNOTFOUND = "ResourceNotFound.CompanyNotFound"
//  RESOURCEUNAVAILABLE_COMPANYUNAVAILABLE = "ResourceUnavailable.CompanyUnavailable"
func (c *Client) UploadFile(request *UploadFileRequest) (response *UploadFileResponse, err error) {
    return c.UploadFileWithContext(context.Background(), request)
}

// UploadFile
// 客户通过调用该接口上传需催收文档，格式需为excel格式。接口返回任务ID。
//
// 可能返回的错误码:
//  FAILEDOPERATION_UPLOADDATAERROR = "FailedOperation.UploadDataError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"
//  RESOURCENOTFOUND_COMPANYNOTFOUND = "ResourceNotFound.CompanyNotFound"
//  RESOURCEUNAVAILABLE_COMPANYUNAVAILABLE = "ResourceUnavailable.CompanyUnavailable"
func (c *Client) UploadFileWithContext(ctx context.Context, request *UploadFileRequest) (response *UploadFileResponse, err error) {
    if request == nil {
        request = NewUploadFileRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadFileResponse()
    err = c.Send(request, response)
    return
}
