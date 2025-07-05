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

package v20180408

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// CAM签名/鉴权错误。
	AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// Apk检测服务端无法响应。
	INTERNALERROR_APKSERVERERROR = "InternalError.ApkServerError"

	// 服务端无法响应。
	INTERNALERROR_SERVERERROR = "InternalError.ServerError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// ServiceInfo结构体参数缺失。
	INVALIDPARAMETER_MISSINGSERVICEINFO = "InvalidParameter.MissingServiceInfo"

	// 参数格式错误。
	INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"

	// 不能找到指定的加固策略。
	INVALIDPARAMETER_PLANIDNOTFOUND = "InvalidParameter.PlanIdNotFound"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// Android在线加固任务,入参字段AppMd5必输。
	INVALIDPARAMETERVALUE_ANDROIDAPPMD5ERROR = "InvalidParameterValue.AndroidAppMd5Error"

	// Android应用包名必输,请求参数中的AppPkgName值要相等。
	INVALIDPARAMETERVALUE_ANDROIDAPPPKGNAMEERROR = "InvalidParameterValue.AndroidAppPkgNameError"

	// AppPkgNameList字段检查不通过，必输字段，值可为单个包年或多个包名时。值为多个包名时，包名需用逗号隔开，其数量不能超过10个。
	INVALIDPARAMETERVALUE_ANDROIDAPPPKGNAMELISTERROR = "InvalidParameterValue.AndroidAppPkgNameListError"

	// Android应用类型必输,请求参数中的AppType值要相等，其值需等于“apk”或“aab”。
	INVALIDPARAMETERVALUE_ANDROIDAPPTYPEERROR = "InvalidParameterValue.AndroidAppTypeError"

	// Android在线加固任务,入参字段AppUrl必输。
	INVALIDPARAMETERVALUE_ANDROIDAPPURLERROR = "InvalidParameterValue.AndroidAppUrlError"

	// Android加固配置错误,assets资源文件路径格式错误。
	INVALIDPARAMETERVALUE_ANDROIDASSETSERROR = "InvalidParameterValue.AndroidAssetsError"

	// Android加固配置参数不正确
	INVALIDPARAMETERVALUE_ANDROIDENCRYPTPARAMERROR = "InvalidParameterValue.AndroidEncryptParamError"

	// Android加固配置错误,res资源文件路径格式错误。
	INVALIDPARAMETERVALUE_ANDROIDRESERROR = "InvalidParameterValue.AndroidResError"

	// Android加固配置错误,so库文件路径格式错误。
	INVALIDPARAMETERVALUE_ANDROIDSOERROR = "InvalidParameterValue.AndroidSoError"

	// Android加固配置错误,vmp加固profile文件必输
	INVALIDPARAMETERVALUE_ANDROIDVMPERROR = "InvalidParameterValue.AndroidVMPError"

	// 小程序加固配置错误,请检查相关配置。
	INVALIDPARAMETERVALUE_APPLETENCRYPTPARAMERROR = "InvalidParameterValue.AppletEncryptParamError"

	// 请求参数中值为Url下载链接的相关字段检查不通过。
	INVALIDPARAMETERVALUE_CHECKURLERROR = "InvalidParameterValue.CheckUrlError"

	// Android加固配置错误,入参字段AppType必输。
	INVALIDPARAMETERVALUE_ENCRYPTPARAMAPPTYPEERROR = "InvalidParameterValue.EncryptParamAppTypeError"

	// 不能同时指定ItemIds和Filters。
	INVALIDPARAMETERVALUE_INVALIDCOEXISTITEMIDSFILTERS = "InvalidParameterValue.InvalidCoexistItemIdsFilters"

	// 无效的过滤器。
	INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"

	// ItemIds不合法。
	INVALIDPARAMETERVALUE_INVALIDITEMIDS = "InvalidParameterValue.InvalidItemIds"

	// Limit不是有效的数字。
	INVALIDPARAMETERVALUE_INVALIDLIMIT = "InvalidParameterValue.InvalidLimit"

	// Offset不是有效的数字。
	INVALIDPARAMETERVALUE_INVALIDOFFSET = "InvalidParameterValue.InvalidOffset"

	// OrderDirection参数。
	INVALIDPARAMETERVALUE_INVALIDORDERDIRECTION = "InvalidParameterValue.InvalidOrderDirection"

	// OrderField取值不合法。
	INVALIDPARAMETERVALUE_INVALIDORDERFIELD = "InvalidParameterValue.InvalidOrderField"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 加固任务超过配额限制。
	LIMITEXCEEDED_ENCRYPTTASKLIMITEXCEEDED = "LimitExceeded.EncryptTaskLimitExceeded"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// AppInfo结构体参数缺失。
	MISSINGPARAMETER_MISSINGAPPINFO = "MissingParameter.MissingAppInfo"

	// 缺少ItemId字段。
	MISSINGPARAMETER_MISSINGITEMID = "MissingParameter.MissingItemId"

	// 缺少ItemIds字段。
	MISSINGPARAMETER_MISSINGITEMIDS = "MissingParameter.MissingItemIds"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// ItemId不存在。
	RESOURCENOTFOUND_ITEMIDNOTFOUND = "ResourceNotFound.ItemIdNotFound"

	// 无法找到指定的加固策略。
	RESOURCENOTFOUND_PLANIDNOTFOUND = "ResourceNotFound.PlanIdNotFound"

	// 加固任务id不存在。
	RESOURCENOTFOUND_RESULTIDNOTFOUND = "ResourceNotFound.ResultIdNotFound"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 资源未绑定应用包名。
	RESOURCEUNAVAILABLE_NOTBIND = "ResourceUnavailable.NotBind"

	// 找不到该资源。
	RESOURCEUNAVAILABLE_NOTFOUND = "ResourceUnavailable.NotFound"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 鉴权失败。
	UNAUTHORIZEDOPERATION_AUTHFAILED = "UnauthorizedOperation.AuthFailed"

	// 不是白名单用户。
	UNAUTHORIZEDOPERATION_NOTWHITEUSER = "UnauthorizedOperation.NotWhiteUser"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"
)
