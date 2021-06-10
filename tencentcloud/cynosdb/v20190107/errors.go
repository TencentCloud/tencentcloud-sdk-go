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

package v20190107

const (
	// 此产品的特有错误码

	// 创建并支付订单失败。
	FAILEDOPERATION_CREATEORDER = "FailedOperation.CreateOrder"

	// 账号余额不足。
	FAILEDOPERATION_INSUFFICIENTBALANCE = "FailedOperation.InsufficientBalance"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 查询数据库失败。
	INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"

	// 获取安全组信息失败。
	INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"

	// 获取子网失败。
	INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"

	// 获取VPC失败。
	INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"

	// 安全组查询实例失败。
	INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"

	// 操作不支持。
	INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"

	// 查询数据库失败。
	INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"

	// 系统内部错误。
	INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"

	// 当前实例不可隔离。
	INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"

	// 未查询到集群。
	INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"

	// 不支持的实例类型。
	INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"

	// 未查询到订单id。
	INVALIDPARAMETERVALUE_DEALNAMENOTFOUND = "InvalidParameterValue.DealNameNotFound"

	// 实例名称字符非法。
	INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"

	// 无效的排序字段。
	INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"

	// 密码不符合要求。
	INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"

	// 实例不存在。
	INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"

	// 实例版本非法。
	INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"

	// 实例规格非法。
	INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"

	// 参数错误。
	INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"

	// 所选地域和可用区不可用。
	INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"

	// 未找到相关存储pool。
	INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"

	// 找不到所选子网。
	INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"

	// 找不到所选VPC网络。
	INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"

	// 用户实例个数超出限制。
	LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"

	// 锁定实例失败，暂时不可操作。
	RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"

	// 实例状态异常，暂时不可操作。
	RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"

	// 非实名用户禁止购买。
	UNAUTHORIZEDOPERATION_NOTREALNAMEACCOUNT = "UnauthorizedOperation.NotRealNameAccount"

	// CAM鉴权不通过。
	UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
)
