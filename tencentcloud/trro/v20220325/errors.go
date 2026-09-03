// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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

package v20220325

const (
	// 此产品的特有错误码

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 存储连通性测试失败。
	FAILEDOPERATION_STORAGECHECKFAILED = "FailedOperation.StorageCheckFailed"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 批量列举 Filter 正则表达式非法。
	INVALIDPARAMETERVALUE_INVALIDFILTERREGEX = "InvalidParameterValue.InvalidFilterRegex"

	// 视频 URL 不可达或非法。
	INVALIDPARAMETERVALUE_INVALIDHTTPURL = "InvalidParameterValue.InvalidHttpUrl"

	// 存储区域与桶实际地域不匹配。
	INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"

	// 指定前缀下没有视频文件。
	INVALIDPARAMETERVALUE_NOVIDEOFILES = "InvalidParameterValue.NoVideoFiles"

	// 存储服务地址不可达。
	INVALIDPARAMETERVALUE_STORAGEUNREACHABLE = "InvalidParameterValue.StorageUnreachable"

	// 输入文件不是支持的视频格式。
	INVALIDPARAMETERVALUE_UNSUPPORTEDVIDEOFORMAT = "InvalidParameterValue.UnsupportedVideoFormat"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 标注模式未开放。
	OPERATIONDENIED_ANNOTATIONTYPENOTENABLED = "OperationDenied.AnnotationTypeNotEnabled"

	// license数量不足
	OPERATIONDENIED_LICENSESNOTENOUGHERROR = "OperationDenied.LicensesNotEnoughError"

	// 处理项当前状态不可重试。
	OPERATIONDENIED_TASKNOTRETRYABLE = "OperationDenied.TaskNotRetryable"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 存储密钥无效或无权限。
	UNAUTHORIZEDOPERATION_STORAGEAUTHFAILED = "UnauthorizedOperation.StorageAuthFailed"
)
