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

package v20220601

const (
	// 此产品的特有错误码

	// 查询数据失败。
	FAILEDOPERATION_QUERYDATA = "FailedOperation.QueryData"

	// RPC服务调用失败。
	FAILEDOPERATION_RPCSERVICECALLFAILED = "FailedOperation.RPCServiceCallFailed"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 内部错误，数据库异常。
	INTERNALERROR_DATABASEEXCEPTION = "InternalError.DatabaseException"

	// 数据库查询失败。
	INTERNALERROR_DATABASEQUERYFAILED = "InternalError.DatabaseQueryFailed"

	// 数据库写入失败。
	INTERNALERROR_DATABASEWRITEFAILED = "InternalError.DatabaseWriteFailed"

	// 没有零信任网关数据。
	INTERNALERROR_NOTZEROTRUSTGATEWAY = "InternalError.NotZeroTrustGateway"

	// 后台服务系统类型错误。
	INTERNALERROR_SERVICETYPEERROR = "InternalError.ServiceTypeError"

	// 内部未知错误。
	INTERNALERROR_UNKNOWN = "InternalError.Unknown"

	// 参数错误：规则相关字段传参或者字段内容错误
	INVALIDPARAMETER_AUTORULEPARAMETERERROR = "InvalidParameter.AutoRuleParameterError"

	// 数据库错误
	INVALIDPARAMETER_DATABASEEXCEPTION = "InvalidParameter.DatabaseException"

	// 重复的终端自定义分组名称。
	INVALIDPARAMETER_DUPLICATEDEVICEVIRTUALGROUPNAME = "InvalidParameter.DuplicateDeviceVirtualGroupName"

	// 黑白名单选中有相同的终端设备
	INVALIDPARAMETER_DUPLICATEIDINBLACKWHITELIST = "InvalidParameter.DuplicateIdInBlackWhiteList"

	// 存在重复的资源数据。
	INVALIDPARAMETER_DUPLICATERESOURCESEXIST = "InvalidParameter.DuplicateResourcesExist"

	// 存在重复业务资源名称。
	INVALIDPARAMETER_DUPLICATERESOURCESNAME = "InvalidParameter.DuplicateResourcesName"

	// 参数错误有重复数据。
	INVALIDPARAMETER_IDENTICALNAME = "InvalidParameter.IdenticalName"

	// 请求参数错误。
	INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"

	// 请求参数错误。
	INVALIDPARAMETER_REQUESTPARAM = "InvalidParameter.RequestParam"

	// 业务资源访问方式错误。
	INVALIDPARAMETER_RESOURCEACCESSMETHOD = "InvalidParameter.ResourceAccessMethod"

	// 业务资源域名不合法。
	INVALIDPARAMETER_RESOURCEDOMAINERR = "InvalidParameter.ResourceDomainErr"

	// 业务资源IP数据错误。
	INVALIDPARAMETER_RESOURCEIPERROR = "InvalidParameter.ResourceIPError"

	// 业务资源IP段数据错误。
	INVALIDPARAMETER_RESOURCEIPSEGMENTERR = "InvalidParameter.ResourceIPSegmentErr"

	// 业务资源优先级取值范围[1-65535]。
	INVALIDPARAMETER_RESOURCELEVELSPARAMETER = "InvalidParameter.ResourceLevelsParameter"

	// 业务资源端口错误。
	INVALIDPARAMETER_RESOURCEPORTERR = "InvalidParameter.ResourcePortErr"

	// 业务资源协议类型错误。
	INVALIDPARAMETER_RESOURCEPROTOCOLTYPEERR = "InvalidParameter.ResourceProtocolTypeErr"

	// 业务资源类型错误。
	INVALIDPARAMETER_RESOURCETYPEPARAMETER = "InvalidParameter.ResourceTypeParameter"

	// 存在同名用户组
	INVALIDPARAMETER_SAMEACCOUNTGROUPNAME = "InvalidParameter.SameAccountGroupName"

	// 没有找到对应的终端自定义分组
	INVALIDPARAMETERVALUE_VIRTUALDEVICEGROUPNOTFOUND = "InvalidParameterValue.VirtualDeviceGroupNotFound"

	// 可创建目录总数超过了上限
	LIMITEXCEEDED_COMPANYDIRECTORYMAXLIMIT = "LimitExceeded.CompanyDirectoryMaxLimit"

	// 缺少公共参数。
	MISSINGPARAMETER_COMMONPARAM = "MissingParameter.CommonParam"

	// 终端导出任务同一时间仅允许一个，请稍后再试。
	REQUESTLIMITEXCEEDED_DEVICEDOWNLOADTASK = "RequestLimitExceeded.DeviceDownloadTask"

	// 网关限流
	REQUESTLIMITEXCEEDED_GATEWAY = "RequestLimitExceeded.Gateway"

	// 上次操作还未完成，请稍后重试。
	REQUESTLIMITEXCEEDED_WAITFORTHELASTOPERATIONTOCOMPLETE = "RequestLimitExceeded.WaitForTheLastOperationToComplete"

	// 没有模块数据。
	RESOURCENOTFOUND_NOMODULEDATA = "ResourceNotFound.NoModuleData"

	// 没有业务资源数据。
	RESOURCENOTFOUND_NORESOURCEDATA = "ResourceNotFound.NoResourceData"

	// 没有连接器分组信息。
	RESOURCENOTFOUND_NOTCONNECTORGROUPINFO = "ResourceNotFound.NotConnectorGroupInfo"

	// 资源不存在。
	RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"

	// license无效
	RESOURCEUNAVAILABLE_LICENSEINVALID = "ResourceUnavailable.LicenseInvalid"

	// 没有模块权限。
	UNAUTHORIZEDOPERATION_NOMODULEPERMISSIONS = "UnauthorizedOperation.NoModulePermissions"

	// 没有业务资源权限。
	UNAUTHORIZEDOPERATION_NORESOURCEPERMISSIONS = "UnauthorizedOperation.NoResourcePermissions"

	// 未授权的操作。
	UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"

	// 业务资源操作异常。
	UNSUPPORTEDOPERATION_DATABASEEXCEPTION = "UnsupportedOperation.DatabaseException"
)
