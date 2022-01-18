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

package v20190923

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// DryRun 操作，代表请求将会是成功的，只是多传了 DryRun 参数。
	DRYRUNOPERATION = "DryRunOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// KMS操作失败。
	FAILEDOPERATION_ACCESSKMSERROR = "FailedOperation.AccessKmsError"

	// 凭据被禁止轮转。
	FAILEDOPERATION_ROTATIONFORBIDDEN = "FailedOperation.RotationForbidden"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 标签键重复。
	INVALIDPARAMETERVALUE_TAGKEYSDUPLICATED = "InvalidParameterValue.TagKeysDuplicated"

	// 标签键或标签值不存在。
	INVALIDPARAMETERVALUE_TAGSNOTEXISTED = "InvalidParameterValue.TagsNotExisted"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// AccessKey已经达到上限。
	OPERATIONDENIED_ACCESSKEYOVERLIMIT = "OperationDenied.AccessKeyOverLimit"

	// 不允许手动更新具有自动轮换功能的凭据。
	OPERATIONDENIED_AUTOROTATEDRESOURCE = "OperationDenied.AutoRotatedResource"

	// 角色不存在。
	OPERATIONDENIED_ROLENOTEXIST = "OperationDenied.RoleNotExist"

	// 请求的次数超过了频率限制。
	REQUESTLIMITEXCEEDED = "RequestLimitExceeded"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 凭据名已存在。
	RESOURCEINUSE_SECRETEXISTS = "ResourceInUse.SecretExists"

	// 版本号已存在。
	RESOURCEINUSE_VERSIONIDEXISTS = "ResourceInUse.VersionIdExists"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 凭据不存在。
	RESOURCENOTFOUND_SECRETNOTEXIST = "ResourceNotFound.SecretNotExist"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 服务未购买。
	RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"

	// 凭据被禁用。
	RESOURCEUNAVAILABLE_RESOURCEDISABLED = "ResourceUnavailable.ResourceDisabled"

	// 凭据处于计划删除状态。
	RESOURCEUNAVAILABLE_RESOURCEPENDINGDELETED = "ResourceUnavailable.ResourcePendingDeleted"

	// 凭据未完成初始化。
	RESOURCEUNAVAILABLE_RESOURCEUNINITIALIZED = "ResourceUnavailable.ResourceUninitialized"

	// 资源售罄。
	RESOURCESSOLDOUT = "ResourcesSoldOut"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 访问KMS失败。
	UNAUTHORIZEDOPERATION_ACCESSKMSERROR = "UnauthorizedOperation.AccessKmsError"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"
)
