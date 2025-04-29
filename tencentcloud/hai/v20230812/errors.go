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

package v20230812

const (
	// 此产品的特有错误码

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 欠费账户不能创建实例
	FAILEDOPERATION_ARREARSACCOUNTCANNOTRUNINSTANCES = "FailedOperation.ArrearsAccountCannotRunInstances"

	// 资源不允许退还
	FAILEDOPERATION_CANNOTRETURN = "FailedOperation.CannotReturn"

	// 并发操作实例失败
	FAILEDOPERATION_OPERATEINSTANCEREPEATEDLY = "FailedOperation.OperateInstanceRepeatedly"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 两个指定的入参每次只能使用其中一个
	INVALIDPARAMETER_ATMOSTONE = "InvalidParameter.AtMostOne"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 指定的应用不存在
	INVALIDPARAMETERVALUE_APPLICATIONIDNOTFOUND = "InvalidParameterValue.ApplicationIdNotFound"

	// 算力套餐类型不存在
	INVALIDPARAMETERVALUE_BUNDLETYPENOTFOUND = "InvalidParameterValue.BundleTypeNotFound"

	// 云盘大小需要大于应用所需云盘最小大小
	INVALIDPARAMETERVALUE_DISKSIZEMUSTLARGERTHANAPPLICATIONMINREQUIREDSIZE = "InvalidParameterValue.DiskSizeMustLargerThanApplicationMinRequiredSize"

	// 列表入参中存在重复值
	INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"

	// 指定的实例不存在
	INVALIDPARAMETERVALUE_INSTANCEIDNOTFOUND = "InvalidParameterValue.InstanceIdNotFound"

	// 实例名称过长
	INVALIDPARAMETERVALUE_INSTANCENAMETOOLONG = "InvalidParameterValue.InstanceNameTooLong"

	// 应用ID格式非法
	INVALIDPARAMETERVALUE_INVALIDAPPLICATIONIDMALFORMED = "InvalidParameterValue.InvalidApplicationIdMalformed"

	// 自定义应用名称重复
	INVALIDPARAMETERVALUE_INVALIDAPPLICATIONNAMEDUPLICATE = "InvalidParameterValue.InvalidApplicationNameDuplicate"

	// 应用类型application-type仅可以选择PUBLIC_APPLICATION,PRIVATE_APPLICATION或COMMUNITY_APPLICATION
	INVALIDPARAMETERVALUE_INVALIDAPPLICATIONTYPE = "InvalidParameterValue.InvalidApplicationType"

	// 付费方式字段不符合规范
	INVALIDPARAMETERVALUE_INVALIDCHARGETYPE = "InvalidParameterValue.InvalidChargeType"

	// 付费方式与时间单位不匹配
	INVALIDPARAMETERVALUE_INVALIDCHARGETYPEANDTIMEUNIT = "InvalidParameterValue.InvalidChargeTypeAndTimeUnit"

	// 每次购买的实例数目不在合理范围内
	INVALIDPARAMETERVALUE_INVALIDINSTANCECOUNT = "InvalidParameterValue.InvalidInstanceCount"

	// 实例ID格式非法
	INVALIDPARAMETERVALUE_INVALIDINSTANCEIDMALFORMED = "InvalidParameterValue.InvalidInstanceIdMalformed"

	// 无效的排练顺序
	INVALIDPARAMETERVALUE_INVALIDORDER = "InvalidParameterValue.InvalidOrder"

	// 无效的排序字段
	INVALIDPARAMETERVALUE_INVALIDORDERFIELD = "InvalidParameterValue.InvalidOrderField"

	// 场景ID格式非法
	INVALIDPARAMETERVALUE_INVALIDSCENEIDMALFORMED = "InvalidParameterValue.InvalidSceneIdMalformed"

	// 地域非法
	INVALIDPARAMETERVALUE_REGIONINVALID = "InvalidParameterValue.RegionInvalid"

	// 无法找到指定的安全组
	INVALIDPARAMETERVALUE_SECURITYGROUPNOTFOUND = "InvalidParameterValue.SecurityGroupNotFound"

	// 入参中若存在子网ID，则必须同时存在VPC ID。
	INVALIDPARAMETERVALUE_SUBNETMUSTCOEXISTWITHVPC = "InvalidParameterValue.SubnetMustCoexistWithVpc"

	// 无法找到指定的子网，或该子网不属于指定的VPC
	INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"

	// 无法找到指定的VPC
	INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"

	// 没有配额去为HAI创建新的安全组
	LIMITEXCEEDED_SECURITYGROUPLIMITEXCEEDED = "LimitExceeded.SecurityGroupLimitExceeded"

	// 没有配额去为HAI创建新的VPC
	LIMITEXCEEDED_VPCLIMITEXCEEDED = "LimitExceeded.VpcLimitExceeded"

	// 指定实例有正在执行的操作，不能执行新的操作
	OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"

	// 指定的算力套餐库存不足
	RESOURCEINSUFFICIENT_BUNDLEINVENTORYSHORTAGE = "ResourceInsufficient.BundleInventoryShortage"

	// 高性能应用服务相关角色获取失败，无法完成操作
	UNAUTHORIZEDOPERATION_GETROLEERROR = "UnauthorizedOperation.GetRoleError"

	// 用户账号的网络类型是传统型，不允许使用HAI
	UNAUTHORIZEDOPERATION_UNAUTHORIZEDNETWORKUSER = "UnauthorizedOperation.UnauthorizedNetworkUser"

	// 客户未授权使用本产品
	UNAUTHORIZEDOPERATION_UNAUTHORIZEDUSER = "UnauthorizedOperation.UnauthorizedUser"

	// 不允许同时操作预付费过期实例和未过期实例
	UNSUPPORTEDOPERATION_CANNOTTERMINATEEXPIREDANDNOTEXPIREDINSTANCES = "UnsupportedOperation.CannotTerminateExpiredAndNotExpiredInstances"

	// 实例不支持创建自定义应用
	UNSUPPORTEDOPERATION_INSTANCECREATEAPPLICATIONNOTSUPPORT = "UnsupportedOperation.InstanceCreateApplicationNotSupport"

	// 不支持对处于`Arrears`状态的实例执行当前操作。
	UNSUPPORTEDOPERATION_INSTANCESTATEARREARS = "UnsupportedOperation.InstanceStateArrears"

	// 不支持对处于`LaunchFailed`状态的实例执行当前操作。
	UNSUPPORTEDOPERATION_INSTANCESTATELAUNCHFAILED = "UnsupportedOperation.InstanceStateLaunchFailed"

	// 不支持对处于`Pending`状态的实例执行当前操作。
	UNSUPPORTEDOPERATION_INSTANCESTATEPENDING = "UnsupportedOperation.InstanceStatePending"

	// 不支持对处于`Running`状态的实例执行当前操作。
	UNSUPPORTEDOPERATION_INSTANCESTATERUNNING = "UnsupportedOperation.InstanceStateRunning"

	// 不支持对处于`StoppedNoCharge`状态的实例执行当前操作。
	UNSUPPORTEDOPERATION_INSTANCESTATESTOPPEDNOCHARGE = "UnsupportedOperation.InstanceStateStoppedNoCharge"

	// 不支持对处于`Terminated`状态的实例执行当前操作。
	UNSUPPORTEDOPERATION_INSTANCESTATETERMINATED = "UnsupportedOperation.InstanceStateTerminated"

	// 不支持同时操作预付费和后付费实例
	UNSUPPORTEDOPERATION_NOTSUPPORTTERMINATEPREPAIDANDPOSTPAID = "UnsupportedOperation.NotSupportTerminatePrePaidAndPostPaid"
)
