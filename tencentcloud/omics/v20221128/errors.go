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

package v20221128

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 重试次数超过上限
	FAILEDOPERATION_RETRYLIMITEXCEEDED = "FailedOperation.RetryLimitExceeded"

	// 状态不支持
	FAILEDOPERATION_STATUSNOTSUPPORTED = "FailedOperation.StatusNotSupported"

	// 版本未发布
	FAILEDOPERATION_VERSIONNOTRELEASED = "FailedOperation.VersionNotReleased"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 表格列重复
	INVALIDPARAMETERVALUE_DUPLICATEHEADER = "InvalidParameterValue.DuplicateHeader"

	// 名称重复
	INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"

	// 入口文件未指定
	INVALIDPARAMETERVALUE_ENTRYPOINTNOTSET = "InvalidParameterValue.EntrypointNotSet"

	// 环境不可用
	INVALIDPARAMETERVALUE_ENVIRONMENTNOTAVAILABLE = "InvalidParameterValue.EnvironmentNotAvailable"

	// Base64编码错误
	INVALIDPARAMETERVALUE_INVALIDBASE64ENCODE = "InvalidParameterValue.InvalidBase64Encode"

	// COS路径错误
	INVALIDPARAMETERVALUE_INVALIDCOSKEY = "InvalidParameterValue.InvalidCosKey"

	// CSV文件格式错误
	INVALIDPARAMETERVALUE_INVALIDCSVFORMAT = "InvalidParameterValue.InvalidCsvFormat"

	// 描述错误
	INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"

	// 表格列错误
	INVALIDPARAMETERVALUE_INVALIDHEADER = "InvalidParameterValue.InvalidHeader"

	// 输入JSON格式错误
	INVALIDPARAMETERVALUE_INVALIDINPUTJSONFORMAT = "InvalidParameterValue.InvalidInputJsonFormat"

	// 输入占位符错误
	INVALIDPARAMETERVALUE_INVALIDINPUTPLACEHOLDER = "InvalidParameterValue.InvalidInputPlaceholder"

	// 名称错误
	INVALIDPARAMETERVALUE_INVALIDNAME = "InvalidParameterValue.InvalidName"

	// 运行参数错误
	INVALIDPARAMETERVALUE_INVALIDRUNOPTION = "InvalidParameterValue.InvalidRunOption"

	// 表格行错误
	INVALIDPARAMETERVALUE_INVALIDTABLEROW = "InvalidParameterValue.InvalidTableRow"

	// 长度超过限制
	INVALIDPARAMETERVALUE_LENGTHLIMITEXCEEDED = "InvalidParameterValue.LengthLimitExceeded"

	// 数据类型不支持
	INVALIDPARAMETERVALUE_UNSUPPORTEDDATATYPE = "InvalidParameterValue.UnsupportedDataType"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 应用不存在
	RESOURCENOTFOUND_APPLICATIONNOTEXIST = "ResourceNotFound.ApplicationNotExist"

	// 应用版本不存在
	RESOURCENOTFOUND_APPLICATIONVERSIONNOTEXIST = "ResourceNotFound.ApplicationVersionNotExist"

	// 存储桶不存在
	RESOURCENOTFOUND_COSBUCKETNOTEXIST = "ResourceNotFound.CosBucketNotExist"

	// 存储对象不存在
	RESOURCENOTFOUND_COSOBJECTNOTEXIST = "ResourceNotFound.CosObjectNotExist"

	// 环境不存在
	RESOURCENOTFOUND_ENVIRONMENTNOTEXIST = "ResourceNotFound.EnvironmentNotExist"

	// 项目不存在
	RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"

	// 任务批次不存在
	RESOURCENOTFOUND_RUNGROUPNOTEXIST = "ResourceNotFound.RunGroupNotExist"

	// 任务不存在
	RESOURCENOTFOUND_RUNNOTEXIST = "ResourceNotFound.RunNotExist"

	// 表格不存在
	RESOURCENOTFOUND_TABLENOTEXIST = "ResourceNotFound.TableNotExist"

	// 表格行不存在
	RESOURCENOTFOUND_TABLEROWNOTEXIST = "ResourceNotFound.TableRowNotExist"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"
)
