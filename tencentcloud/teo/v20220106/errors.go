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

package v20220106

const (
	// 此产品的特有错误码

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 后台处理出错。
	INTERNALERROR_BACKENDERROR = "InternalError.BackendError"

	// 获取配置失败。
	INTERNALERROR_DOMAINCONFIG = "InternalError.DomainConfig"

	// 上传链接生成失败。
	INTERNALERROR_FAILEDTOGENERATEURL = "InternalError.FailedToGenerateUrl"

	// 配额系统处理失败。
	INTERNALERROR_QUOTASYSTEM = "InternalError.QuotaSystem"

	// 域名不存在或不属于该账号。
	INVALIDPARAMETER_DOMAINNOTFOUND = "InvalidParameter.DomainNotFound"

	// 参数错误。
	INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"

	// 资源存在错误。
	INVALIDPARAMETER_TARGET = "InvalidParameter.Target"

	// 任务无法生成。
	INVALIDPARAMETER_TASKNOTGENERATED = "InvalidParameter.TaskNotGenerated"

	// 文件上传链接存在问题。
	INVALIDPARAMETER_UPLOADURL = "InvalidParameter.UploadUrl"

	// 本次提交的资源数超过上限。
	LIMITEXCEEDED_BATCHQUOTA = "LimitExceeded.BatchQuota"

	// 当天提交的资源数超过上限。
	LIMITEXCEEDED_DAILYQUOTA = "LimitExceeded.DailyQuota"

	// Cam 未授权。
	UNAUTHORIZEDOPERATION_CAMUNAUTHORIZED = "UnauthorizedOperation.CamUnauthorized"

	// 子账户没有操作权限，请添加权限后继续操作。
	UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
)
