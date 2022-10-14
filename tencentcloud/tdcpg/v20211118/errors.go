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

package v20211118

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 元数据库访问失败，请稍后重试。如果持续不成功，请联系客服进行处理。
	FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"

	// 创建流程失败,请稍后重试。如果持续不成功，请联系客服进行处理。
	FAILEDOPERATION_FLOWERROR = "FailedOperation.FlowError"

	// 获取VPC信息异常，请稍后重试。如果持续不成功，请联系客服进行处理。
	FAILEDOPERATION_GETVPCFAILED = "FailedOperation.GetVpcFailed"

	// 内部服务访问失败，请稍后重试。如果持续不成功，请联系客服。
	FAILEDOPERATION_INTERNALSERVICEACCESSERROR = "FailedOperation.InternalServiceAccessError"

	// 当前的付费模式不支持此操作。
	FAILEDOPERATION_PAYMODEINVALID = "FailedOperation.PayModeInvalid"

	// 规格没有变化，CPU/Memory至少升级一项。
	FAILEDOPERATION_SPECNOTCHANGE = "FailedOperation.SpecNotChange"

	// 对应规格的存储上限不够，无法满足当前集群的数据存储。
	FAILEDOPERATION_SPECSTORAGELACK = "FailedOperation.SpecStorageLack"

	// 状态异常，不允许操作。
	FAILEDOPERATION_STATUSERROR = "FailedOperation.StatusError"

	// 当前存储的付费模式不支持此操作。
	FAILEDOPERATION_STORAGEPAYMODEINVALID = "FailedOperation.StoragePayModeInvalid"

	// 任务冲突，请稍后重试。如果持续不成功，请联系客服。
	FAILEDOPERATION_TASKCONFLICT = "FailedOperation.TaskConflict"

	// 请求计费服务异常，请稍后重试。如果持续不成功，请联系客服进行处理。
	FAILEDOPERATION_TRADEACCESSERROR = "FailedOperation.TradeAccessError"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 账号不存在。
	INVALIDPARAMETERVALUE_ACCOUNTNOTFOUND = "InvalidParameterValue.AccountNotFound"

	// 未找到可回档的时间,请确认指定时间点是否在(创建集群时间点,当前时间点)之间。
	INVALIDPARAMETERVALUE_BACKUPDATAPOINTINVALID = "InvalidParameterValue.BackupDataPointInvalid"

	// 集群不存在。
	INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"

	// 数据库版本号相关参数数量错误。
	INVALIDPARAMETERVALUE_DATABASEVERSIONPARAMCOUNTERROR = "InvalidParameterValue.DatabaseVersionParamCountError"

	// 订单未找到。
	INVALIDPARAMETERVALUE_DEALNAMENOTFOUND = "InvalidParameterValue.DealNameNotFound"

	// 接入点不存在。
	INVALIDPARAMETERVALUE_ENDPOINTNOTFOUND = "InvalidParameterValue.EndpointNotFound"

	// 集群/实例名字非法，需要满足：长度在1-60个字符，只能由中文、字母、数字、'-'或'.'或'_'组成，区分大小写。
	INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"

	// 密码设置无效，需要满足： 8-64个字符，至少包含 大写字母、小写字母、数字和符号~!@#$%^&*_-+=`|(){}[]:;'<>,.?/中的任意三种
	INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"

	// 实例不存在。
	INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"

	// 数据库版本信息无法被识别。
	INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"

	// 数据库版本号相关参数不合法。
	INVALIDPARAMETERVALUE_INVALIDDATABASEVERSION = "InvalidParameterValue.InvalidDatabaseVersion"

	// 传入的参数非法。
	INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"

	// 规格信息(CPU/Memory)信息无法被识别。
	INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"

	// 参数不合法。
	INVALIDPARAMETERVALUE_PARAMETEROUTRANGEERROR = "InvalidParameterValue.ParameterOutRangeError"

	// 不支持当前地域/可用区的售卖。
	INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"

	// 使用的备份数据来源集群ID非法。
	INVALIDPARAMETERVALUE_SOURCEBACKUPCLUSTERIDINVALID = "InvalidParameterValue.SourceBackupClusterIdInvalid"

	// 集群对应的存储已经被删除。
	INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"

	// 您没有权限操作该VPC网络。
	INVALIDPARAMETERVALUE_VPCDENIEDERROR = "InvalidParameterValue.VpcDeniedError"

	// 未获取到VPC信息，请检查输入的VPC相关参数。
	INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"

	// VPC子网中IP数量不够。
	INVALIDPARAMETERVALUE_VPCSUBNETIPLACK = "InvalidParameterValue.VpcSubnetIpLack"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 当前集群下实例数量超过限制。
	LIMITEXCEEDED_CLUSTERINSTANCELIMIT = "LimitExceeded.ClusterInstanceLimit"

	// 用户集群数量超过限制。
	LIMITEXCEEDED_USERCLUSTERLIMIT = "LimitExceeded.UserClusterLimit"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 当前集群/实例状态不允许操作。
	RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 未经授权的操作。
	UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"
)
