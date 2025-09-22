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

package v20210622

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// CAM鉴权失败
	AUTHFAILURE_ACCESSCAMFAIL = "AuthFailure.AccessCAMFail"

	// AuthFailure.UnauthorizedOperation
	AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"

	// CAM鉴权解析失败
	AUTHFAILURE_UNMARSHALRESPONSE = "AuthFailure.UnmarshalResponse"

	// DryRun 操作，代表请求将会是成功的，只是多传了 DryRun 参数。
	DRYRUNOPERATION = "DryRunOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 访问标签失败。
	FAILEDOPERATION_ACCESSTAGFAIL = "FailedOperation.AccessTagFail"

	// 未检测到探针在线
	FAILEDOPERATION_AGENTNOTONLINEERROR = "FailedOperation.AgentNotOnlineError"

	// 当前探针版本不支持此功能
	FAILEDOPERATION_AGENTVERSIONNOTSUPPORTERROR = "FailedOperation.AgentVersionNotSupportError"

	// token信息不存在。
	FAILEDOPERATION_APMCREDENTIALNOTEXIST = "FailedOperation.ApmCredentialNotExist"

	// AppID 和业务系统信息不匹配。
	FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"

	// 未命中白名单且业务系统 ID 为官方 Demo 业务系统 ID 时，不允许修改接口。
	FAILEDOPERATION_DEMOINSTANCENOTALLOWMODIFIED = "FailedOperation.DemoInstanceNotAllowModified"

	// 请勿填写重复的应用名
	FAILEDOPERATION_DUPLICATESERVICE = "FailedOperation.DuplicateService"

	// 请勿填写重复的标签名
	FAILEDOPERATION_DUPLICATETAGFIELD = "FailedOperation.DuplicateTagField"

	// 该业务系统不允许修改
	FAILEDOPERATION_INSTANCECANNOTMODIFY = "FailedOperation.InstanceCannotModify"

	// 该业务系统不允许销毁
	FAILEDOPERATION_INSTANCECANNOTTERMINATE = "FailedOperation.InstanceCannotTerminate"

	// 业务系统 ID 为空。
	FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"

	// APM 业务系统不存在。
	FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"

	// 非法业务系统 ID。
	FAILEDOPERATION_INVALIDINSTANCEID = "FailedOperation.InvalidInstanceID"

	// 不合法入参
	FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"

	// 非法的正则表达式。
	FAILEDOPERATION_INVALIDREGEX = "FailedOperation.InvalidRegex"

	// 不合法请求
	FAILEDOPERATION_INVALIDREQUEST = "FailedOperation.InvalidRequest"

	// 业务系统和应用名不匹配。
	FAILEDOPERATION_INVALIDSERVICENAME = "FailedOperation.InvalidServiceName"

	// Tag中指定了无效的Key
	FAILEDOPERATION_INVALIDTAGFIELD = "FailedOperation.InvalidTagField"

	// 不合法token
	FAILEDOPERATION_INVALIDTOKEN = "FailedOperation.InvalidToken"

	// 查询指标类数据查询条件缺少过滤参数。
	FAILEDOPERATION_METRICFILTERSLACKPARAMS = "FailedOperation.MetricFiltersLackParams"

	// 非内网 VPC。
	FAILEDOPERATION_NOTINNERVPC = "FailedOperation.NotInnerVPC"

	// 查询时间区间不支持。
	FAILEDOPERATION_QUERYTIMEINTERVALISNOTSUPPORTED = "FailedOperation.QueryTimeIntervalIsNotSupported"

	// 不支持该地域。
	FAILEDOPERATION_REGIONNOTSUPPORT = "FailedOperation.RegionNotSupport"

	// 发送查询请求失败。
	FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"

	// 应用数超过10个
	FAILEDOPERATION_SERVICELISTEXCEEDINGLIMITNUMBER = "FailedOperation.ServiceListExceedingLimitNumber"

	// 应用列表为空
	FAILEDOPERATION_SERVICELISTNULL = "FailedOperation.ServiceListNull"

	// 没有找到应用资源
	FAILEDOPERATION_SERVICENOTFOUND = "FailedOperation.ServiceNotFound"

	// serviceId 与 appId 不匹配
	FAILEDOPERATION_SERVICENOTMATCHAPPIDERR = "FailedOperation.ServiceNotMatchAppIdErr"

	// 视图名不存在或非法。
	FAILEDOPERATION_VIEWNAMENOTEXISTORILLEGAL = "FailedOperation.ViewNameNotExistOrIllegal"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// Filters 中的字段不存在或非法。
	INVALIDPARAMETER_FILTERSFIELDSNOTEXISTORILLEGAL = "InvalidParameter.FiltersFieldsNotExistOrIllegal"

	// GroupBy 中的字段不存在或非法。
	INVALIDPARAMETER_GROUPBYFIELDSNOTEXISTORILLEGAL = "InvalidParameter.GroupByFieldsNotExistOrIllegal"

	// Filters 中必须存在 service.name 字段，否则会报错。
	INVALIDPARAMETER_METRICFILTERSLACKPARAMS = "InvalidParameter.MetricFiltersLackParams"

	// Metrics 中的字段不存在或非法。
	INVALIDPARAMETER_METRICSFIELDNOTEXISTORILLEGAL = "InvalidParameter.MetricsFieldNotExistOrIllegal"

	// Metrics 中不允许为空。
	INVALIDPARAMETER_METRICSFIELDSNOTALLOWEMPTY = "InvalidParameter.MetricsFieldsNotAllowEmpty"

	// Period不为空，0或60。
	INVALIDPARAMETER_PERIODISILLEGAL = "InvalidParameter.PeriodIsIllegal"

	// 查询时间不支持，最多只能查询最近30天的数据。
	INVALIDPARAMETER_QUERYTIMEINTERVALISNOTSUPPORTED = "InvalidParameter.QueryTimeIntervalIsNotSupported"

	// 视图名称不存在或非法。
	INVALIDPARAMETER_VIEWNAMENOTEXISTORILLEGAL = "InvalidParameter.ViewNameNotExistOrIllegal"
)
