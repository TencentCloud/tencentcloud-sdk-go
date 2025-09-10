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

package v20221128

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 表格表头重复
	FAILEDOPERATION_DUPLICATETABLEHEADER = "FailedOperation.DuplicateTableHeader"

	// 邮件数量超出限制
	FAILEDOPERATION_EMAILCOUNTLIMITEXCEEDED = "FailedOperation.EmailCountLimitExceeded"

	// 表格表头为空
	FAILEDOPERATION_EMPTYTABLEHEADER = "FailedOperation.EmptyTableHeader"

	// 表格表头错误
	FAILEDOPERATION_INVALIDTABLEHEADER = "FailedOperation.InvalidTableHeader"

	// 表格行数错误
	FAILEDOPERATION_INVALIDTABLELENGTH = "FailedOperation.InvalidTableLength"

	// 重试次数超过上限
	FAILEDOPERATION_RETRYLIMITEXCEEDED = "FailedOperation.RetryLimitExceeded"

	// 状态不支持
	FAILEDOPERATION_STATUSNOTSUPPORTED = "FailedOperation.StatusNotSupported"

	// 表格数据与类型不匹配
	FAILEDOPERATION_TABLEDATATYPEMISMATCH = "FailedOperation.TableDataTypeMismatch"

	// 版本未发布
	FAILEDOPERATION_VERSIONNOTRELEASED = "FailedOperation.VersionNotReleased"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

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

	// 邮箱格式错误
	INVALIDPARAMETERVALUE_INVALIDEMAILFORMAT = "InvalidParameterValue.InvalidEmailFormat"

	// 输入JSON格式错误
	INVALIDPARAMETERVALUE_INVALIDINPUTJSONFORMAT = "InvalidParameterValue.InvalidInputJsonFormat"

	// 输入占位符错误
	INVALIDPARAMETERVALUE_INVALIDINPUTPLACEHOLDER = "InvalidParameterValue.InvalidInputPlaceholder"

	// 名称错误
	INVALIDPARAMETERVALUE_INVALIDNAME = "InvalidParameterValue.InvalidName"

	// 运行参数错误
	INVALIDPARAMETERVALUE_INVALIDRUNOPTION = "InvalidParameterValue.InvalidRunOption"

	// 超时分钟数无效
	INVALIDPARAMETERVALUE_INVALIDTIMEOUTMINUTES = "InvalidParameterValue.InvalidTimeoutMinutes"

	// 表格数据与类型长度不匹配
	INVALIDPARAMETERVALUE_TABLEDATATYPELENGTHMISMATCH = "InvalidParameterValue.TableDataTypeLengthMismatch"

	// 表格数据类型不支持
	INVALIDPARAMETERVALUE_UNSUPPORTEDTABLEDATATYPE = "InvalidParameterValue.UnsupportedTableDataType"

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
