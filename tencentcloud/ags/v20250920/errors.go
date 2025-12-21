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

package v20250920

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 请求未CAM授权
	AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 重复请求（幂等性检查）
	FAILEDOPERATION_DUPLICATEREQUEST = "FailedOperation.DuplicateRequest"

	// 请求正在处理中（幂等性检查）
	FAILEDOPERATION_REQUESTINPROGRESS = "FailedOperation.RequestInProgress"

	// 沙箱实例存储配置挂载失败
	FAILEDOPERATION_STORAGEMOUNT = "FailedOperation.StorageMount"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// InstanceIds 参数格式错误或 ID 列表超过最大数量限制
	INVALIDPARAMETERVALUE_INSTANCEIDS = "InvalidParameterValue.InstanceIds"

	// MountOption部分参数填写错误
	INVALIDPARAMETERVALUE_MOUNTOPTION = "InvalidParameterValue.MountOption"

	// Role Arn格式不满足腾讯云规范
	INVALIDPARAMETERVALUE_ROLEARN = "InvalidParameterValue.RoleArn"

	// 沙箱工具名称不可用，可能是已经存在
	INVALIDPARAMETERVALUE_SANDBOXTOOL = "InvalidParameterValue.SandboxTool"

	// StorageMount部分参数错误
	INVALIDPARAMETERVALUE_STORAGEMOUNT = "InvalidParameterValue.StorageMount"

	// 超时时间格式错误或超过最大限制
	INVALIDPARAMETERVALUE_TIMEOUT = "InvalidParameterValue.Timeout"

	// ToolIds 参数格式错误或 ID 列表超过最大数量限制
	INVALIDPARAMETERVALUE_TOOLIDS = "InvalidParameterValue.ToolIds"

	// 不支持的沙箱工具类型
	INVALIDPARAMETERVALUE_TOOLTYPE = "InvalidParameterValue.ToolType"

	// 账号下 API 密钥数量达到上限
	LIMITEXCEEDED_APIKEYQUOTA = "LimitExceeded.APIKeyQuota"

	// 沙箱实例配额超限
	LIMITEXCEEDED_SANDBOXINSTANCE = "LimitExceeded.SandboxInstance"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 缺少角色Arn
	MISSINGPARAMETER_ROLEARN = "MissingParameter.RoleArn"

	// 沙箱工具正在使用中
	RESOURCEINUSE_SANDBOXTOOL = "ResourceInUse.SandboxTool"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 沙箱实例不存在
	RESOURCENOTFOUND_SANDBOXINSTANCE = "ResourceNotFound.SandboxInstance"

	// 沙箱工具不存在
	RESOURCENOTFOUND_SANDBOXTOOL = "ResourceNotFound.SandboxTool"

	// 指定的StorageMount不存在
	RESOURCENOTFOUND_STORAGEMOUNT = "ResourceNotFound.StorageMount"

	// 沙箱工具不可用
	RESOURCEUNAVAILABLE_SANDBOXTOOL = "ResourceUnavailable.SandboxTool"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 实例状态不允许修改（只有RUNNING状态的实例可以修改）
	UNSUPPORTEDOPERATION_SANDBOXINSTANCE = "UnsupportedOperation.SandboxInstance"
)
