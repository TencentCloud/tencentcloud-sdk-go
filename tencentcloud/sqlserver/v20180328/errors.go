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

package v20180328

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 查询对象存储失败。
	FAILEDOPERATION_COSERROR = "FailedOperation.CosError"

	// 获取上传配置信息错误。
	FAILEDOPERATION_COSPROPERTIESERROR = "FailedOperation.CosPropertiesError"

	// 创建订单失败。
	FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"

	// 数据库错误。
	FAILEDOPERATION_DBERROR = "FailedOperation.DBError"

	// 操作失败或者网络超时。
	FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"

	// 获取VPC网络信息失败。
	FAILEDOPERATION_GETVPCFAILED = "FailedOperation.GetVpcFailed"

	// 备份导入任务锁定失败。
	FAILEDOPERATION_MIGRATIONLOCKERROR = "FailedOperation.MigrationLockError"

	// 查询订单失败。
	FAILEDOPERATION_QUERYORDERFAILED = "FailedOperation.QueryOrderFailed"

	// 计费相关错误，查询价格失败。
	FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"

	// 安全组操作失败。
	FAILEDOPERATION_SECURITYGROUPOPERATIONERROR = "FailedOperation.SecurityGroupOperationError"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// CAM鉴权请求失败。
	INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"

	// COS接口错误。
	INTERNALERROR_COSERROR = "InternalError.CosError"

	// 流程创建失败。
	INTERNALERROR_CREATEFLOWFAILED = "InternalError.CreateFlowFailed"

	// 数据库连接错误。
	INTERNALERROR_DBCONNECTERROR = "InternalError.DBConnectError"

	// 数据库错误。
	INTERNALERROR_DBERROR = "InternalError.DBError"

	// GCS接口错误。
	INTERNALERROR_GCSERROR = "InternalError.GcsError"

	// 获取临时密钥错误。
	INTERNALERROR_STSERROR = "InternalError.StsError"

	// 系统错误。
	INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"

	// 未知错误。
	INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"

	// VPC错误。
	INTERNALERROR_VPCERROR = "InternalError.VPCError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 输入错误。
	INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"

	// 接口不存在。
	INVALIDPARAMETER_INTERFACENAMENOTFOUND = "InvalidParameter.InterfaceNameNotFound"

	// 参数断言转换错误。
	INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"

	// 支付订单失败。
	INVALIDPARAMETER_PAYORDERFAILED = "InvalidParameter.PayOrderFailed"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 账号已经存在。
	INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"

	// 账户名称不合法。
	INVALIDPARAMETERVALUE_ACCOUNTNAMEISILLEGAL = "InvalidParameterValue.AccountNameIsIllegal"

	// 账号名称不允许是保留字。
	INVALIDPARAMETERVALUE_ACCOUNTNAMEISKEYWORDS = "InvalidParameterValue.AccountNameIsKeyWords"

	// 账户备注内容不合法。
	INVALIDPARAMETERVALUE_ACCOUNTREMARKISILLEGAL = "InvalidParameterValue.AccountRemarkIsIllegal"

	// 管理员账号只能申请一个。
	INVALIDPARAMETERVALUE_ADMINACCOUNTNOTUNIQUE = "InvalidParameterValue.AdminAccountNotUnique"

	// 备份名称存在非法字符。
	INVALIDPARAMETERVALUE_BACKUPNAMEISILLEGAL = "InvalidParameterValue.BackupNameIsIllegal"

	// 购买实例数量错误。
	INVALIDPARAMETERVALUE_BADGOODSNUM = "InvalidParameterValue.BadGoodsNum"

	// 数据库字符集设置错误。
	INVALIDPARAMETERVALUE_CHARSETISILLEGAL = "InvalidParameterValue.CharsetIsIllegal"

	// 上传路径错误。
	INVALIDPARAMETERVALUE_COSPATHERROR = "InvalidParameterValue.CosPathError"

	// 计费类型错误。
	INVALIDPARAMETERVALUE_COSTTYPENOTSUPPORTED = "InvalidParameterValue.CostTypeNotSupported"

	// 数据库名称包含非法字符。
	INVALIDPARAMETERVALUE_DBCHARILLEGAL = "InvalidParameterValue.DBCharIllegal"

	// 数据库已经存在。
	INVALIDPARAMETERVALUE_DBEXIST = "InvalidParameterValue.DBExist"

	// 数据库名称不允许是保留字。
	INVALIDPARAMETERVALUE_DBNAMEISKEYWRODS = "InvalidParameterValue.DBNameIsKeyWrods"

	// 数据库名称不能为空。
	INVALIDPARAMETERVALUE_DBNAMENOTNULL = "InvalidParameterValue.DBNameNotNull"

	// 数据库重命名名称一样。
	INVALIDPARAMETERVALUE_DBNAMESAME = "InvalidParameterValue.DBNameSame"

	// 数据库备注内容不合法。
	INVALIDPARAMETERVALUE_DATABASEREMARKISILLEGAL = "InvalidParameterValue.DataBaseRemarkIsIllegal"

	// 参数错误，GRANT 值非法。
	INVALIDPARAMETERVALUE_GRANTISILLEGAL = "InvalidParameterValue.GrantIsIllegal"

	// 地域错误。
	INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"

	// 实例规格信息错误。
	INVALIDPARAMETERVALUE_ILLEGALSPEC = "InvalidParameterValue.IllegalSpec"

	// 可用区ID错误。
	INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"

	// 实例扩容容量低于当前容量。
	INVALIDPARAMETERVALUE_INSTANCEEXPANDVOLUMELOW = "InvalidParameterValue.InstanceExpandVolumeLow"

	// 实例名称存在非法字符。
	INVALIDPARAMETERVALUE_INSTANCENAMEISILLEGAL = "InvalidParameterValue.InstanceNameIsIllegal"

	// 迁移名称包含非法字符。
	INVALIDPARAMETERVALUE_MIGRATIONNAMEISILLEGAL = "InvalidParameterValue.MigrationNameIsIllegal"

	// 可选参数[enable, disable]。
	INVALIDPARAMETERVALUE_MODIFYTYPEVALUEINVALID = "InvalidParameterValue.ModifyTypeValueInvalid"

	// 是CVM类型。
	INVALIDPARAMETERVALUE_ONCVMTYPENOTSUPPORTED = "InvalidParameterValue.OnCvmTypeNotSupported"

	// 参数类型错误。
	INVALIDPARAMETERVALUE_PARAMETERTYPEERROR = "InvalidParameterValue.ParameterTypeError"

	// 账号密码非法。
	INVALIDPARAMETERVALUE_PASSWORDISILLEGAL = "InvalidParameterValue.PasswordIsIllegal"

	// 数据库权限设置错误。
	INVALIDPARAMETERVALUE_PRIVILEGEISILLEGAL = "InvalidParameterValue.PrivilegeIsIllegal"

	// 发布订阅名称不符合规范。
	INVALIDPARAMETERVALUE_PUBSUBNAMEISILLEGAL = "InvalidParameterValue.PubSubNameIsIllegal"

	// 只读组名称包含非法字符。
	INVALIDPARAMETERVALUE_ROGROUPNAMEISILLEGAL = "InvalidParameterValue.RoGroupNameIsIllegal"

	// 只读组状态不正确。
	INVALIDPARAMETERVALUE_ROGROUPSTATUSISILLEGAL = "InvalidParameterValue.RoGroupStatusIsIllegal"

	// 安全组ID非法。
	INVALIDPARAMETERVALUE_SECURITYGROUPIDISILLEGAL = "InvalidParameterValue.SecurityGroupIdIsIllegal"

	// 数据库超过限制。
	LIMITEXCEEDED_TOOMANYDB = "LimitExceeded.TooManyDB"

	// 已经存在一个准备启动的增量导入任务。
	RESOURCEINUSE_INCREMENTALMIGRATIONEXIST = "ResourceInUse.IncrementalMigrationExist"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 账号不存在。
	RESOURCENOTFOUND_ACCOUNTNOTEXIST = "ResourceNotFound.AccountNotExist"

	// 备份不存在。
	RESOURCENOTFOUND_BACKUPNOTFOUND = "ResourceNotFound.BackupNotFound"

	// 数据库不存在。
	RESOURCENOTFOUND_DBNOTEXIT = "ResourceNotFound.DBNotExit"

	// 数据库不存在。
	RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"

	// 全量备份导入任务不存在。
	RESOURCENOTFOUND_FULLBACKUPMIGRATIONNOTEXIST = "ResourceNotFound.FullBackupMigrationNotExist"

	// 增量备份导入任务不存在。
	RESOURCENOTFOUND_INCREBACKUPMIGRATIONNOTEXIST = "ResourceNotFound.IncreBackupMigrationNotExist"

	// 实例不存在。
	RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"

	// VPC网络不存在。
	RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"

	// 账号状态无效。
	RESOURCEUNAVAILABLE_ACCOUNTINVALIDSTATUS = "ResourceUnavailable.AccountInvalidStatus"

	// 恢复类型不支持增量备份导入。
	RESOURCEUNAVAILABLE_BACKUPMIGRATIONRECOVERYTYPEERR = "ResourceUnavailable.BackupMigrationRecoveryTypeErr"

	// 离线恢复任务状态错误。
	RESOURCEUNAVAILABLE_COSSTATUSERR = "ResourceUnavailable.CosStatusErr"

	// 数据库状态无效。
	RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"

	// 实例迁移地域非法。
	RESOURCEUNAVAILABLE_INSTANCEMIGRATEREGIONILLEGAL = "ResourceUnavailable.InstanceMigrateRegionIllegal"

	// 实例迁移状态无效。
	RESOURCEUNAVAILABLE_INSTANCEMIGRATESTATUSINVALID = "ResourceUnavailable.InstanceMigrateStatusInvalid"

	// 实例状态无效。
	RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"

	// 不支持只读实例。
	RESOURCEUNAVAILABLE_NOTSUPPORTROINSTANCE = "ResourceUnavailable.NotSupportRoInstance"

	// VPC不存在。
	RESOURCEUNAVAILABLE_VPCNOTEXIST = "ResourceUnavailable.VpcNotExist"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// cam鉴权错误。
	UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 不支持重复操作。
	UNSUPPORTEDOPERATION_NOTSUPPORTREPEAT = "UnsupportedOperation.NotSupportRepeat"

	// 上传类型错误。
	UNSUPPORTEDOPERATION_UPLOADTYPEERROR = "UnsupportedOperation.UploadTypeError"
)
