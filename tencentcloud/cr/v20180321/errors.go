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

const (
	// 此产品的特有错误码

	// 黑名单更新出错。
	FAILEDOPERATION_APPLYBLACKLISTDATAERROR = "FailedOperation.ApplyBlackListDataError"

	// 提交黑名单出错。
	FAILEDOPERATION_APPLYBLACKLISTERROR = "FailedOperation.ApplyBlackListError"

	// 提交信审申请出错。
	FAILEDOPERATION_APPLYCREDITAUDITERROR = "FailedOperation.ApplyCreditAuditError"

	// 机器人任务作业状态更新出错。
	FAILEDOPERATION_CHANGEBOTCALLSTATUSERROR = "FailedOperation.ChangeBotCallStatusError"

	// 机器人任务状态更新出错。
	FAILEDOPERATION_CHANGEBOTSTATUSERROR = "FailedOperation.ChangeBotStatusError"

	// 创建机器人任务出错。
	FAILEDOPERATION_CREATEBOTTASKERROR = "FailedOperation.CreateBotTaskError"

	// 获取黑名单数据列表出错。
	FAILEDOPERATION_DESCRIBEBLACKLISTDATAERROR = "FailedOperation.DescribeBlacklistDataError"

	// 查询机器人对话流出错。
	FAILEDOPERATION_DESCRIBEBOTFLOWERROR = "FailedOperation.DescribeBotFlowError"

	// 查询机器人文件模板出错。
	FAILEDOPERATION_DESCRIBEFILEMODELERROR = "FailedOperation.DescribeFileModelError"

	// 获取产品列表出错。
	FAILEDOPERATION_DESCRIBEPRODUCTSERROR = "FailedOperation.DescribeProductsError"

	// 获取录音列表出错。
	FAILEDOPERATION_DESCRIBERECORDSERROR = "FailedOperation.DescribeRecordsError"

	// 查询任务状态出错。
	FAILEDOPERATION_DESCRIBETASKSTATUSERROR = "FailedOperation.DescribeTaskStatusError"

	// 录音列表下载出错。
	FAILEDOPERATION_DOWNLOADRECORDLISTERROR = "FailedOperation.DownloadRecordListError"

	// 报告下载出错。
	FAILEDOPERATION_DOWNLOADREPORTERROR = "FailedOperation.DownloadReportError"

	// 导出机器人数据失败。
	FAILEDOPERATION_EXPORTBOTDATAERROR = "FailedOperation.ExportBotDataError"

	// 获取信审结果出错。
	FAILEDOPERATION_GETCREDITAUDITERROR = "FailedOperation.GetCreditAuditError"

	// 查询黑名单出错。
	FAILEDOPERATION_QUERYBLACKLISTDATAERROR = "FailedOperation.QueryBlackListDataError"

	// 查询机器人列表出错。
	FAILEDOPERATION_QUERYBOTLISTERROR = "FailedOperation.QueryBotListError"

	// 查询任务失败。
	FAILEDOPERATION_QUERYCALLLISTERROR = "FailedOperation.QueryCallListError"

	// 查询录音列表出错。
	FAILEDOPERATION_QUERYRECORDLISTERROR = "FailedOperation.QueryRecordListError"

	// 修改机器人任务出错。
	FAILEDOPERATION_UPDATEBOTTASKERROR = "FailedOperation.UpdateBotTaskError"

	// 上传机器人文件出错。
	FAILEDOPERATION_UPLOADBOTFILE = "FailedOperation.UploadBotFile"

	// 上传数据出错。
	FAILEDOPERATION_UPLOADDATAERROR = "FailedOperation.UploadDataError"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 内部未知错误。
	INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"

	// 找不到Module 或 Operation。
	MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"

	// 黑名单数据不存在。
	RESOURCENOTFOUND_BLACKLISTDATANOTFOUND = "ResourceNotFound.BlacklistDataNotFound"

	// 企业不存在。
	RESOURCENOTFOUND_COMPANYNOTFOUND = "ResourceNotFound.CompanyNotFound"

	// 录音列表不存在。
	RESOURCENOTFOUND_RECORDLISTNOTFOUND = "ResourceNotFound.RecordListNotFound"

	// 报告不存在。
	RESOURCENOTFOUND_REPORTNOTFOUND = "ResourceNotFound.ReportNotFound"

	// 任务不存在。
	RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"

	// 企业已欠费或已被回收。
	RESOURCEUNAVAILABLE_COMPANYUNAVAILABLE = "ResourceUnavailable.CompanyUnavailable"

	// 账户不存在或未开通金融联络机器人。
	UNAUTHORIZEDOPERATION_ACCOUNTNOTFOUND = "UnauthorizedOperation.AccountNotFound"
)
