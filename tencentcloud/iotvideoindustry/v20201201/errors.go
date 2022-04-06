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

package v20201201

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// DryRun 操作，代表请求将会是成功的，只是多传了 DryRun 参数。
	DRYRUNOPERATION = "DryRunOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// DB操作错误。
	INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"

	// 内部错误。
	INTERNALERROR_INTERNALSERVEREXCEPTION = "InternalError.InternalServerException"

	// DB操作错误。
	INTERNALERROR_INTERNALSERVEREXCEPTIONDB = "InternalError.InternalServerExceptionDB"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 通道状态在线。
	INVALIDPARAMETER_DEVICEONLINE = "InvalidParameter.DeviceOnline"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 账户余额不足，请检查。
	INVALIDPARAMETERVALUE_BALANCENOTENOUGH = "InvalidParameterValue.BalanceNotEnough"

	// 无法通过该设备ID查找设备。
	INVALIDPARAMETERVALUE_DEVICEDATAMAPERROR = "InvalidParameterValue.DeviceDataMapError"

	// 设备Id不合法。
	INVALIDPARAMETERVALUE_DEVICEID = "InvalidParameterValue.DeviceId"

	// 设备通道不在线，请检查通道配置。
	INVALIDPARAMETERVALUE_DEVICEINFONOTEXIST = "InvalidParameterValue.DeviceInfoNotExist"

	// 设备离线或未注册。
	INVALIDPARAMETERVALUE_DEVICEOFFLINE = "InvalidParameterValue.DeviceOffline"

	// 设备或通道处于在线状态。
	INVALIDPARAMETERVALUE_DEVICEONLINE = "InvalidParameterValue.DeviceOnline"

	// 设备类型不支持。
	INVALIDPARAMETERVALUE_DEVICETYPENOTSUPPORT = "InvalidParameterValue.DeviceTypeNotSupport"

	// 国标域/设备不存在或未注册。
	INVALIDPARAMETERVALUE_DOMAINID = "InvalidParameterValue.DomainId"

	// 超时时间错误。
	INVALIDPARAMETERVALUE_EXPIRETIME = "InvalidParameterValue.ExpireTime"

	// domainid分组不允许修改ExtraInformation。
	INVALIDPARAMETERVALUE_GROUPDOMAINIDEXTRAINFORMATION = "InvalidParameterValue.GroupDomainidExtraInformation"

	// 国标平台下级分组不允许修改。
	INVALIDPARAMETERVALUE_GROUPDOMAINIDNOTUPDATE = "InvalidParameterValue.GroupDomainidNotUpdate"

	// 分组参数错误。
	INVALIDPARAMETERVALUE_GROUPPARMSERROR = "InvalidParameterValue.GroupParmsError"

	// 回看流不存在。
	INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"

	// 最多创建100个录制计划。
	INVALIDPARAMETERVALUE_RECORDPLANBEYONDLIMIT = "InvalidParameterValue.RecordPlanBeyondLimit"

	// 超过创建条目限制。
	INVALIDPARAMETERVALUE_RULELIMIT = "InvalidParameterValue.RuleLimit"

	// 规则不存在。
	INVALIDPARAMETERVALUE_RULENOTEXIST = "InvalidParameterValue.RuleNotExist"

	// 流Id不合法。
	INVALIDPARAMETERVALUE_STREAMID = "InvalidParameterValue.StreamId"

	// 流不存在，请检查设备配置。
	INVALIDPARAMETERVALUE_STREAMINFONOTEXIST = "InvalidParameterValue.StreamInfoNotExist"

	// 最多创建100个时间模板。
	INVALIDPARAMETERVALUE_TEMPLATEBEYONDLIMIT = "InvalidParameterValue.TemplateBeyondLimit"

	// 模板时间片段为空。
	INVALIDPARAMETERVALUE_TEMPLATESPECEMPTY = "InvalidParameterValue.TemplateSpecEmpty"

	// 请求时间粒度不支持。
	INVALIDPARAMETERVALUE_TIMESPECNOTSUPPORT = "InvalidParameterValue.TimeSpecNotSupport"

	// 请求类型不支持。
	INVALIDPARAMETERVALUE_TYPENOTSUPPORT = "InvalidParameterValue.TypeNotSupport"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 请求的次数超过了频率限制。
	REQUESTLIMITEXCEEDED = "RequestLimitExceeded"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 设备未创建或是已删除。
	RESOURCENOTFOUND_DEVICENOTEXIST = "ResourceNotFound.DeviceNotExist"

	// PTZ控制的资源处于离线状态。
	RESOURCENOTFOUND_DEVICESIPPTZERROR = "ResourceNotFound.DeviceSipPTZError"

	// 未找到分组或分组已删除。
	RESOURCENOTFOUND_GROUPNOTEXIST = "ResourceNotFound.GroupNotExist"

	// 未找到录制计划或录制计划已删除。
	RESOURCENOTFOUND_PLANNOTEXIST = "ResourceNotFound.PlanNotExist"

	// 未找到时间模板或时间模板已删除。
	RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"

	// 统计数据不存在。
	RESOURCENOTFOUND_STATDATANOTEXIST = "ResourceNotFound.StatDataNotExist"

	// 未找到时间模板或时间模板已删除。
	RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 国标信令异常。
	RESOURCEUNAVAILABLE_GBPROTOCOLEXECEXCEPTION = "ResourceUnavailable.GBProtocolExecException"

	// 流信息异常。
	RESOURCEUNAVAILABLE_STREAMINFOEXCEPTION = "ResourceUnavailable.StreamInfoException"

	// 资源售罄。
	RESOURCESSOLDOUT = "ResourcesSoldOut"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 该名字分组已创建，请修改其他分组名字。
	UNAUTHORIZEDOPERATION_GROUPEXIST = "UnauthorizedOperation.GroupExist"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 分组下还存在设备或子分组。
	UNSUPPORTEDOPERATION_BINDEXIST = "UnsupportedOperation.BindExist"

	// 绑定的设备列表中有已绑定录制计划的设备。
	UNSUPPORTEDOPERATION_DEVICEBINDEXIST = "UnsupportedOperation.DeviceBindExist"

	// 设备没有操作权限。
	UNSUPPORTEDOPERATION_DEVICECANNOTDO = "UnsupportedOperation.DeviceCanNotDo"

	// 存在重复设备。
	UNSUPPORTEDOPERATION_DEVICEDUPKEYEXIST = "UnsupportedOperation.DeviceDupKeyExist"

	// 设备和通道不匹配，请检查。
	UNSUPPORTEDOPERATION_DEVICEMISSMATCH = "UnsupportedOperation.DeviceMissMatch"

	// 绑定的设备列表中有不存在的设备。
	UNSUPPORTEDOPERATION_DEVICENOTFOUND = "UnsupportedOperation.DeviceNotFound"

	// 设备信令不通。
	UNSUPPORTEDOPERATION_DEVICESIPCOMMANDFAIL = "UnsupportedOperation.DeviceSipCommandFail"

	// 该分组为平台下级分组，不允许删除。
	UNSUPPORTEDOPERATION_DOMAINGROUP = "UnsupportedOperation.DomainGroup"

	// 该名字分组已创建，请修改其他分组名字。
	UNSUPPORTEDOPERATION_GROUPEXIST = "UnsupportedOperation.GroupExist"

	// 分组层级超过最大值。
	UNSUPPORTEDOPERATION_GROUPLAYERISMAX = "UnsupportedOperation.GroupLayerIsMax"

	// 父分组不存在。
	UNSUPPORTEDOPERATION_GROUPPARENTIDNOTEXIST = "UnsupportedOperation.GroupParentidNotExist"

	// 存在直播频道与当前直播录制计划绑定。
	UNSUPPORTEDOPERATION_LIVECHANNELBINDEXIST = "UnsupportedOperation.LiveChannelBindExist"

	// 没有该设备的权限。
	UNSUPPORTEDOPERATION_NOPERMISSION = "UnsupportedOperation.NoPermission"

	// 时间模板下存在未删除的录制计划时不允许删除或修改。
	UNSUPPORTEDOPERATION_PLANEXISTUNDERTEMPLATE = "UnsupportedOperation.PlanExistUnderTemplate"

	// 该名字录制计划已创建，请修改其他计划名字。
	UNSUPPORTEDOPERATION_RECORDPLANEXIST = "UnsupportedOperation.RecordPlanExist"

	// 存在重复规则。
	UNSUPPORTEDOPERATION_RULEDUPKEYEXIST = "UnsupportedOperation.RuleDupKeyExist"

	// 已存在同名场景。
	UNSUPPORTEDOPERATION_SCENEEXIST = "UnsupportedOperation.SceneExist"

	// 子分组数量超过最大值。
	UNSUPPORTEDOPERATION_SUBGROUPISMAX = "UnsupportedOperation.SubGroupIsMax"

	// 分组下还存在设备或子分组。
	UNSUPPORTEDOPERATION_SUBGRPEXIST = "UnsupportedOperation.SubgrpExist"

	// 该名字时间模板已创建，请修改其他模板名字。
	UNSUPPORTEDOPERATION_TEMPLATEEXIST = "UnsupportedOperation.TemplateExist"

	// 预置时间模板不允许删除或修改。
	UNSUPPORTEDOPERATION_TEMPLATEPRESET = "UnsupportedOperation.TemplatePreset"

	// 处于隔离中的资源不能进行此操作。
	UNSUPPORTEDOPERATION_USERISISOLATE = "UnsupportedOperation.UserIsIsolate"
)
