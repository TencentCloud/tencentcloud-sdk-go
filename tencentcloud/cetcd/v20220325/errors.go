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

	// 资源已存在。
	FAILEDOPERATION_RESOURCEEXISTALREADY = "FailedOperation.ResourceExistAlready"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 记录未找到。
	INTERNALERROR_DBRECORDNOTFOUND = "InternalError.DbRecordNotFound"

	// 内部错误。
	INTERNALERROR_UNEXPECTEDINTERNAL = "InternalError.UnexpectedInternal"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数错误。
	INVALIDPARAMETER_PARAM = "InvalidParameter.Param"

	// etcd配额达到限制。
	LIMITEXCEEDED_QUOTAETCDLIMIT = "LimitExceeded.QuotaEtcdLimit"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// Prometheus状态异常。
	RESOURCEUNAVAILABLE_PROMETHEUSSTATUSERROR = "ResourceUnavailable.PrometheusStatusError"

	// CAM权限校验失败。
	UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"
)
