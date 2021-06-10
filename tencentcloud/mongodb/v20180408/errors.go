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

package v20180408

const (
	// 此产品的特有错误码

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 查询异步任务错误。
	INTERNALERROR_ASYNCREQUESTERROR = "InternalError.AsyncRequestError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 当前子帐号无权执行该操作。
	INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"

	// 非法的实例名。
	INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"

	// 非法的实例状态。
	INVALIDPARAMETERVALUE_ILLEGALINSTANCESTATUS = "InvalidParameterValue.IllegalInstanceStatus"

	// 实例已删除。
	INVALIDPARAMETERVALUE_INSTANCEHASBEENDELETED = "InvalidParameterValue.InstanceHasBeenDeleted"

	// 实例已隔离。
	INVALIDPARAMETERVALUE_INSTANCEHASBEENISOLATED = "InvalidParameterValue.InstanceHasBeenIsolated"

	// limit取值范围[1,100]。
	INVALIDPARAMETERVALUE_LIMITPARAOUTOFRANGE = "InvalidParameterValue.LimitParaOutOfRange"

	// 实例锁定失败。
	INVALIDPARAMETERVALUE_LOCKFAILED = "InvalidParameterValue.LockFailed"

	// 实例版本不支持查询客户端信息。
	INVALIDPARAMETERVALUE_MONGOVERSIONNOTSUPPORTQUERYCLIENT = "InvalidParameterValue.MongoVersionNotSupportQueryClient"

	// 未找到实例。
	INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"

	// offset取值范围[0, 10000]。
	INVALIDPARAMETERVALUE_OFFSETPARAOUTOFRANGE = "InvalidParameterValue.OffsetParaOutOfRange"

	// 密码不符合规范。
	INVALIDPARAMETERVALUE_PASSWORDRULEFAILED = "InvalidParameterValue.PasswordRuleFailed"

	// 按量计费实例不允许续费。
	INVALIDPARAMETERVALUE_POSTPAYRENEWERROR = "InvalidParameterValue.PostPayRenewError"

	// 项目不存在。
	INVALIDPARAMETERVALUE_PROJECTNOTFOUND = "InvalidParameterValue.ProjectNotFound"

	// proxy版本不支持查询客户端信息，请联系工作人员进行升级。
	INVALIDPARAMETERVALUE_PROXYNOTSUPPORTQUERYCLIENT = "InvalidParameterValue.ProxyNotSupportQueryClient"

	// 只能查询7天内的慢日志。
	INVALIDPARAMETERVALUE_QUERYTIMEOUTOFRANGE = "InvalidParameterValue.QueryTimeOutOfRange"

	// 只能查询7天内的慢日志。
	INVALIDPARAMETERVALUE_QUERYTIMERANGEBEYONDLIMIT = "InvalidParameterValue.QueryTimeRangeBeyondLimit"

	// 无效的地域。
	INVALIDPARAMETERVALUE_REGIONERROR = "InvalidParameterValue.RegionError"

	// 地域尚不支持查询客户端信息。
	INVALIDPARAMETERVALUE_REGIONNOTSUPPORTQUERYCLIENT = "InvalidParameterValue.RegionNotSupportQueryClient"

	// slowMS参数取值范围[100, 。
	INVALIDPARAMETERVALUE_SLOWMSBELOWLIMIT = "InvalidParameterValue.SlowMSBelowLimit"

	// 起始时间晚于结束时间。
	INVALIDPARAMETERVALUE_STARTTIMENOTBEFORETHANENDTIME = "InvalidParameterValue.StartTimeNotBeforeThanEndTime"

	// 实例处于不允许操作的状态。
	INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"

	// 非法的时间格式。
	INVALIDPARAMETERVALUE_TIMEFORMATERR = "InvalidParameterValue.TimeFormatErr"

	// 用户账户不存在。
	INVALIDPARAMETERVALUE_USERNOTFOUND = "InvalidParameterValue.UserNotFound"

	// 未找到虚拟网络（子网）。
	INVALIDPARAMETERVALUE_VPCIDORSUBNETIDNOTFOUND = "InvalidParameterValue.VpcIdOrSubnetIdNotFound"

	// 无效的可用区。
	INVALIDPARAMETERVALUE_ZONEERROR = "InvalidParameterValue.ZoneError"
)
