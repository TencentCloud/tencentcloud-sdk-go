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

package v20170312

const (
	// 此产品的特有错误码

	// 当前账号不存在。
	ACCOUNTNOTEXIST = "AccountNotExist"

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 鉴权失败。
	AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"

	// 后台DB执行错误。
	DBERROR = "DBError"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 申请资源标签的配额失败。
	FAILEDOPERATION_ALLOCATEQUOTASERROR = "FailedOperation.AllocateQuotasError"

	// 访问基础网络服务失败，请稍后重试。如果持续不成功，请联系客服进行处理。
	FAILEDOPERATION_BASENETWORKACCESSERROR = "FailedOperation.BaseNetworkAccessError"

	// 访问CMQ失败。
	FAILEDOPERATION_CMQRESPONSEERROR = "FailedOperation.CMQResponseError"

	// CAM鉴权失败。
	FAILEDOPERATION_CAMAUTHFAILED = "FailedOperation.CamAuthFailed"

	// 获取权限失败，请稍后重试。如果持续不成功，请联系客服进行处理。。
	FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"

	// CAM资源检查失败。
	FAILEDOPERATION_CAMCHECKRESOURCEFAILED = "FailedOperation.CamCheckResourceFailed"

	// 鉴权失败，请稍后重试。如果持续不成功，请联系客服进行处理。
	FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"

	// 获取项目信息失败，请稍后重试。如果持续不成功，请联系客服进行处理。
	FAILEDOPERATION_CDBCGWCONNECTERROR = "FailedOperation.CdbCgwConnectError"

	// 不支持新增基础网络。
	FAILEDOPERATION_CREATEBASICNETWORKDENIEDERROR = "FailedOperation.CreateBasicNetworkDeniedError"

	// 创建续费订单失败。
	FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"

	// 管控元数据库访问失败，请稍后重试。如果持续不成功，请联系客服进行处理。
	FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"

	// 数据操作失败，请联系客服进行处理。
	FAILEDOPERATION_DATABASEAFFECTEDERROR = "FailedOperation.DatabaseAffectedError"

	// VPC删除路由失败。
	FAILEDOPERATION_DELETEALLROUTE = "FailedOperation.DeleteAllRoute"

	// 资源解绑项目失败。
	FAILEDOPERATION_DELETERESOURCEPROJECTTAGERROR = "FailedOperation.DeleteResourceProjectTagError"

	// 资源解绑标签失败。
	FAILEDOPERATION_DELETERESOURCESTOTAGERROR = "FailedOperation.DeleteResourcesToTagError"

	// ES访问失败，请稍后重试。如果持续不成功，请联系客服进行处理。
	FAILEDOPERATION_ESCONNECTERROR = "FailedOperation.ESConnectError"

	// ES查询失败，请稍后重试。如果持续不成功，请联系客服进行处理。
	FAILEDOPERATION_ESQUERYERROR = "FailedOperation.ESQueryError"

	// 操作失败，请稍后重试。
	FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"

	// 创建流程失败，请稍后重试。如果持续不成功，请联系客服进行处理。
	FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"

	// 查询实例信息失败，请稍后重试。
	FAILEDOPERATION_GETINSTANCEBYRESOURCEIDERROR = "FailedOperation.GetInstanceByResourceIdError"

	// 获取购买记录失败。
	FAILEDOPERATION_GETPURCHASERECORDFAILED = "FailedOperation.GetPurchaseRecordFailed"

	// 获取VPC子网失败。
	FAILEDOPERATION_GETSUBNETERROR = "FailedOperation.GetSubnetError"

	// 获取VPC信息失败，请稍后重试。如果持续不成功，请联系客服进行处理。
	FAILEDOPERATION_GETVPCINFOERROR = "FailedOperation.GetVpcInfoError"

	// 只读实例数目非法。
	FAILEDOPERATION_ILLEGALROINSTANCENUM = "FailedOperation.IllegalROInstanceNum"

	// 当前账号状态不正确，不允许操作。
	FAILEDOPERATION_INVALIDACCOUNTSTATUS = "FailedOperation.InvalidAccountStatus"

	// 计费相关错误，不允许对当前实例进行对应的新购/续费/配置变更操作。
	FAILEDOPERATION_INVALIDTRADEOPERATE = "FailedOperation.InvalidTradeOperate"

	// 实例被限制操作。
	FAILEDOPERATION_LIMITOPERATION = "FailedOperation.LimitOperation"

	// 查询主实例信息失败。
	FAILEDOPERATION_MASTERINSTANCEQUERYERROR = "FailedOperation.MasterInstanceQueryError"

	// 修改只读组配置失败。
	FAILEDOPERATION_MODIFYROGROUPERROR = "FailedOperation.ModifyROGroupError"

	// 不符合资源所拥有的网络数量要求。
	FAILEDOPERATION_NETWORKNUMLIMITERROR = "FailedOperation.NetworkNumLimitError"

	// 操作超过频率限制，请稍后重试。如果持续不成功，请联系客服进行处理。
	FAILEDOPERATION_OPERATEFREQUENCYLIMITEDERROR = "FailedOperation.OperateFrequencyLimitedError"

	// 访问管控服务失败，请稍后重试。如果持续不成功，请联系客服进行处理。
	FAILEDOPERATION_OSSACCESSERROR = "FailedOperation.OssAccessError"

	// 请求后端接口失败。
	FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"

	// 支付订单失败。
	FAILEDOPERATION_PAYORDERFAILED = "FailedOperation.PayOrderFailed"

	// 按量计费实例账户解冻结失败，请稍后重试。如果持续不成功，请联系客服进行处理。
	FAILEDOPERATION_POSTPAIDUNBLOCKERROR = "FailedOperation.PostPaidUnblockError"

	// 获取预签名授权URL错误。
	FAILEDOPERATION_PRESIGNEDURLGETERROR = "FailedOperation.PresignedURLGetError"

	// 查询订单信息失败。
	FAILEDOPERATION_QUERYDEALFAILED = "FailedOperation.QueryDealFailed"

	// 批量获取实例信息失败，请稍后重试。如果持续不成功，请联系客服进行处理。
	FAILEDOPERATION_QUERYINSTANCEBATCHERROR = "FailedOperation.QueryInstanceBatchError"

	// 查询价格失败。
	FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"

	// 查询规格信息失败。
	FAILEDOPERATION_QUERYSPECBYSPECCODEERROR = "FailedOperation.QuerySpecBySpecCodeError"

	// 查询规格信息失败，请稍后重试。如果持续不成功，请联系客服进行处理。
	FAILEDOPERATION_QUERYSPECERROR = "FailedOperation.QuerySpecError"

	// 查询订单交易状态失败。
	FAILEDOPERATION_QUERYTRADESTATUSERROR = "FailedOperation.QueryTradeStatusError"

	// 查询VPC失败。
	FAILEDOPERATION_QUERYVPCFAILED = "FailedOperation.QueryVpcFailed"

	// 获取VPC详情失败。
	FAILEDOPERATION_QUERYVPCFALIED = "FailedOperation.QueryVpcFalied"

	// 只读组还有实例，该只读组不允许删除。
	FAILEDOPERATION_ROGROUPCANNOTBEDELETEDERROR = "FailedOperation.ROGroupCannotBeDeletedError"

	// 只读组主实例信息不匹配。
	FAILEDOPERATION_ROGROUPMASTERINSTANCENOTRIGHT = "FailedOperation.ROGroupMasterInstanceNotRight"

	// 只读组不存在。
	FAILEDOPERATION_ROGROUPNOTFOUNDERROR = "FailedOperation.ROGroupNotFoundError"

	// 只读组数量已经达到上限值。
	FAILEDOPERATION_ROGROUPNUMEXCEEDERROR = "FailedOperation.ROGroupNumExceedError"

	// 只读实例已在其他只读组。
	FAILEDOPERATION_ROINSTANCEHASINROGROUPERROR = "FailedOperation.ROInstanceHasInROGroupError"

	// Serverless实例不支持此操作。
	FAILEDOPERATION_SERVERLESSNOTSUPPORTEDERROR = "FailedOperation.ServerlessNotSupportedError"

	// 访问内部服务失败，请稍后重试。如果持续不成功，请联系客服进行处理。
	FAILEDOPERATION_SERVICEACCESSERROR = "FailedOperation.ServiceAccessError"

	// 设置自动续费标记失败。
	FAILEDOPERATION_SETAUTORENEWFLAGERROR = "FailedOperation.SetAutoRenewFlagError"

	// 实例升配时，存储或内存需要高于原实例规格。
	FAILEDOPERATION_STORAGEMEMORYCHECKERROR = "FailedOperation.StorageMemoryCheckError"

	// 存储路径格式不正确。
	FAILEDOPERATION_STOREPATHSPLITERROR = "FailedOperation.StorePathSplitError"

	// 访问计费平台失败，请稍后重试。如果持续不成功，请联系客服进行处理。
	FAILEDOPERATION_TRADEACCESSERROR = "FailedOperation.TradeAccessError"

	// 请求计费平台创建订单失败。
	FAILEDOPERATION_TRADECREATEERROR = "FailedOperation.TradeCreateError"

	// 请求支付订单失败。
	FAILEDOPERATION_TRADEPAYORDERSERROR = "FailedOperation.TradePayOrdersError"

	// 查询订单信息失败，请稍后重试。如果持续不成功，请联系客服进行处理。
	FAILEDOPERATION_TRADEQUERYORDERERROR = "FailedOperation.TradeQueryOrderError"

	// 获取价格信息失败，请稍后重试。如果持续不成功，请联系客服进行处理。
	FAILEDOPERATION_TRADEQUERYPRICEERROR = "FailedOperation.TradeQueryPriceError"

	// 资源修改标签失败。
	FAILEDOPERATION_UPDATERESOURCEPROJECTTAGVALUEERROR = "FailedOperation.UpdateResourceProjectTagValueError"

	// VPC设置失败。
	FAILEDOPERATION_VPCRESETSERVICEERROR = "FailedOperation.VPCResetServiceError"

	// VPC更新路由失败。
	FAILEDOPERATION_VPCUPDATEROUTEERROR = "FailedOperation.VPCUpdateRouteError"

	// 流程创建失败。
	FLOWERROR = "FlowError"

	// 当前实例不存在。
	INSTANCENOTEXIST = "InstanceNotExist"

	// ACTION输入错误。
	INTERFACENAMENOTFOUND = "InterfaceNameNotFound"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// CGW相关错误。
	INTERNALERROR_CGWERROR = "InternalError.CgwError"

	// 基础网络相关错误。
	INTERNALERROR_CNSERROR = "InternalError.CnsError"

	// 后台DB执行错误。
	INTERNALERROR_DBERROR = "InternalError.DBError"

	// DFW相关错误。
	INTERNALERROR_DFWERROR = "InternalError.DfwError"

	// 流程创建失败。
	INTERNALERROR_FLOWERROR = "InternalError.FlowError"

	// 管控系统元数据访问异常，请联系客服处理。
	INTERNALERROR_INSTANCEDATAERROR = "InternalError.InstanceDataError"

	// 请求执行异常。
	INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"

	// 后台数据解析失败，请联系客服进行处理。
	INTERNALERROR_MARSHALERROR = "InternalError.MarshalError"

	// 系统错误。出现这种错误时，请联系客服支持。
	INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"

	// 开启事务失败。
	INTERNALERROR_TRANSACTIOBEGINERROR = "InternalError.TransactioBeginError"

	// 其他未知错误。出现这种错误时，请联系客服支持。
	INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"

	// VPC相关错误。
	INTERNALERROR_VPCERROR = "InternalError.VpcError"

	// 账号密码不正确。
	INVALIDACCOUNTPASSWORD = "InvalidAccountPassword"

	// 当前账号状态不正确，不允许操作。
	INVALIDACCOUNTSTATUS = "InvalidAccountStatus"

	// 实例状态错误。
	INVALIDINSTANCESTATUS = "InvalidInstanceStatus"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 实例名已存在。
	INVALIDPARAMETER_INSTANCENAMEEXIST = "InvalidParameter.InstanceNameExist"

	// 参数检查失败。
	INVALIDPARAMETER_PARAMETERCHECKERROR = "InvalidParameter.ParameterCheckError"

	// pid错误。
	INVALIDPARAMETER_TRADEACCESSDENIEDERROR = "InvalidParameter.TradeAccessDeniedError"

	// 未获取到VPC信息。
	INVALIDPARAMETER_VPCNOTFOUNDERROR = "InvalidParameter.VpcNotFoundError"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 当前账号已存在。
	INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"

	// 当前账号不存在。
	INVALIDPARAMETERVALUE_ACCOUNTNOTEXIST = "InvalidParameterValue.AccountNotExist"

	// 账号不存在。
	INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"

	// 当前实例所要扩容的规格目前不售卖。
	INVALIDPARAMETERVALUE_BADSPEC = "InvalidParameterValue.BadSpec"

	// 数据库字符集错误。
	INVALIDPARAMETERVALUE_CHARSETNOTFOUNDERROR = "InvalidParameterValue.CharsetNotFoundError"

	// 数据格式转换失败，请联系客服处理。
	INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"

	// 不支持的计费类型。
	INVALIDPARAMETERVALUE_ILLEGALCHARGETYPE = "InvalidParameterValue.IllegalChargeType"

	// 计费模式错误。
	INVALIDPARAMETERVALUE_ILLEGALINSTANCECHARGETYPE = "InvalidParameterValue.IllegalInstanceChargeType"

	// 非法ProjectId。
	INVALIDPARAMETERVALUE_ILLEGALPROJECTID = "InvalidParameterValue.IllegalProjectId"

	// 非法的Region参数。
	INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"

	// 非法的Zone参数。
	INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"

	// 当前账号已存在。
	INVALIDPARAMETERVALUE_INSTANCENAMEEXIST = "InvalidParameterValue.InstanceNameExist"

	// 当前实例不存在。
	INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"

	// ACTION输入错误。
	INVALIDPARAMETERVALUE_INTERFACENAMENOTFOUND = "InvalidParameterValue.InterfaceNameNotFound"

	// 账号设置无效，请遵循账号命名规则：账号名需要1-16个字符，只能由字母、数字或下划线组成；不能为postgres；不能由数字和pg_开头；所有规则均不区分大小写。
	INVALIDPARAMETERVALUE_INVALIDACCOUNTERROR = "InvalidParameterValue.InvalidAccountError"

	// 账号格式不正确。
	INVALIDPARAMETERVALUE_INVALIDACCOUNTFORMAT = "InvalidParameterValue.InvalidAccountFormat"

	// 当前账号名称不允许是保留字符。
	INVALIDPARAMETERVALUE_INVALIDACCOUNTNAME = "InvalidParameterValue.InvalidAccountName"

	// 数据库字符集错误，当前只支持UTF8、LATIN1。
	INVALIDPARAMETERVALUE_INVALIDCHARSET = "InvalidParameterValue.InvalidCharset"

	// 购买实例数超过限制。
	INVALIDPARAMETERVALUE_INVALIDINSTANCENUM = "InvalidParameterValue.InvalidInstanceNum"

	// 扩容后的实例规格不能小于当前实例规格。
	INVALIDPARAMETERVALUE_INVALIDINSTANCEVOLUME = "InvalidParameterValue.InvalidInstanceVolume"

	// 计费相关错误，订单类型ID无效。
	INVALIDPARAMETERVALUE_INVALIDORDERNUM = "InvalidParameterValue.InvalidOrderNum"

	// 参数值有误。
	INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"

	// 密码格式不正确。
	INVALIDPARAMETERVALUE_INVALIDPASSWORDFORMAT = "InvalidParameterValue.InvalidPasswordFormat"

	// 密码设置无效，长度不满足规则。
	INVALIDPARAMETERVALUE_INVALIDPASSWORDLENGTHERROR = "InvalidParameterValue.InvalidPasswordLengthError"

	// 密码设置无效，不能以“/”开头，必须包含大写字母、小写字母、符号()`~!@#$%^&*-+=_|{}[]:;'<>,.?/和数字。
	INVALIDPARAMETERVALUE_INVALIDPASSWORDVALUEERROR = "InvalidParameterValue.InvalidPasswordValueError"

	// PID参数输入错误。
	INVALIDPARAMETERVALUE_INVALIDPID = "InvalidParameterValue.InvalidPid"

	// 只读实例组状态错误。
	INVALIDPARAMETERVALUE_INVALIDREADONLYGROUPSTATUS = "InvalidParameterValue.InvalidReadOnlyGroupStatus"

	// 无效地域。
	INVALIDPARAMETERVALUE_INVALIDREGIONIDERROR = "InvalidParameterValue.InvalidRegionIdError"

	// 无效的可用区。
	INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"

	// 输入订单名为空。
	INVALIDPARAMETERVALUE_NULLDEALNAMES = "InvalidParameterValue.NullDealNames"

	// 参数无效，只允许英文字母、数字、下划线、中划线，以及全体汉字。
	INVALIDPARAMETERVALUE_PARAMETERCHARACTERERROR = "InvalidParameterValue.ParameterCharacterError"

	// 参数不满足规则，请修改后重试。
	INVALIDPARAMETERVALUE_PARAMETERCHARACTERLIMITERROR = "InvalidParameterValue.ParameterCharacterLimitError"

	// 参数前缀设置不符合规则要求，请修改后重试。
	INVALIDPARAMETERVALUE_PARAMETERCHARACTERPRELIMITERROR = "InvalidParameterValue.ParameterCharacterPreLimitError"

	// 参数处理失败，请检参数值设置是否有效。
	INVALIDPARAMETERVALUE_PARAMETERHANDLEERROR = "InvalidParameterValue.ParameterHandleError"

	// 参数长度超过限制。
	INVALIDPARAMETERVALUE_PARAMETERLENGTHLIMITERROR = "InvalidParameterValue.ParameterLengthLimitError"

	// 存在无效的参数值。
	INVALIDPARAMETERVALUE_PARAMETEROUTRANGEERROR = "InvalidParameterValue.ParameterOutRangeError"

	// 参数值超过上限。
	INVALIDPARAMETERVALUE_PARAMETERVALUEEXCEEDERROR = "InvalidParameterValue.ParameterValueExceedError"

	// 只读实例组不存在。
	INVALIDPARAMETERVALUE_READONLYGROUPNOTEXIST = "InvalidParameterValue.ReadOnlyGroupNotExist"

	// 不支持当前地域。
	INVALIDPARAMETERVALUE_REGIONNOTSUPPORTED = "InvalidParameterValue.RegionNotSupported"

	// 规格信息{{1}}无法被识别。
	INVALIDPARAMETERVALUE_SPECNOTRECOGNIZEDERROR = "InvalidParameterValue.SpecNotRecognizedError"

	// 解析参数出错。
	INVALIDPARAMETERVALUE_STRUCTPARSEFAILED = "InvalidParameterValue.StructParseFailed"

	// PID参数输入错误。
	INVALIDPID = "InvalidPid"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 实例被限制操作。
	LIMITOPERATION = "LimitOperation"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 当前操作被限制。
	OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"

	// 目标实例状态检查不通过。
	OPERATIONDENIED_DTSINSTANCESTATUSERROR = "OperationDenied.DTSInstanceStatusError"

	// 您没有权限操作当前资源。
	OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"

	// 不支持ipv6。
	OPERATIONDENIED_INSTANCEIPV6NOTSUPPORTEDERROR = "OperationDenied.InstanceIpv6NotSupportedError"

	// 实例状态限制当前操作。
	OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"

	// 实例当前状态限制本次操作。
	OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"

	// 实例状态限制当前操作。
	OPERATIONDENIED_INSTANCESTATUSLIMITOPERROR = "OperationDenied.InstanceStatusLimitOpError"

	// Serverless不支持当前可用区。
	OPERATIONDENIED_NOTSUPPORTZONEERROR = "OperationDenied.NotSupportZoneError"

	// 不支持的付费类型。
	OPERATIONDENIED_PAYMODEERROR = "OperationDenied.PayModeError"

	// 按量付费实例不支持续费。
	OPERATIONDENIED_POSTPAIDPAYMODEERROR = "OperationDenied.PostPaidPayModeError"

	// 只读组状态限制当前操作。
	OPERATIONDENIED_ROGROUPSTATUSERROR = "OperationDenied.ROGroupStatusError"

	// 只读实例不支持ipv6。
	OPERATIONDENIED_ROINSTANCEIPV6NOTSUPPORTEDERROR = "OperationDenied.ROInstanceIpv6NotSupportedError"

	// 只读实例状态限制当前操作。
	OPERATIONDENIED_ROINSTANCESTATUSLIMITOPERROR = "OperationDenied.ROInstanceStatusLimitOpError"

	// 只读节点总数不能超过上限值。
	OPERATIONDENIED_ROINSTANCECOUNTEXEEDERROR = "OperationDenied.RoInstanceCountExeedError"

	// 用户未进行实名认证，请先进行实名认证才可购买。
	OPERATIONDENIED_USERNOTAUTHENTICATEDERROR = "OperationDenied.UserNotAuthenticatedError"

	// Serverless不支持该版本。
	OPERATIONDENIED_VERSIONNOTSUPPORTERROR = "OperationDenied.VersionNotSupportError"

	// 您没有权限操作该VPC网络。
	OPERATIONDENIED_VPCDENIEDERROR = "OperationDenied.VpcDeniedError"

	// 不支持当前地域。
	REGIONNOTSUPPORTED = "RegionNotSupported"

	// 请求的次数超过了频率限制。
	REQUESTLIMITEXCEEDED = "RequestLimitExceeded"

	// 当前地域购买此规格的实例没有足够的资源。
	RESOURCEINSUFFICIENT_RESOURCENOTENOUGH = "ResourceInsufficient.ResourceNotEnough"

	// 实例不存在。
	RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"

	// 实例状态错误。
	RESOURCEUNAVAILABLE_INVALIDINSTANCESTATUS = "ResourceUnavailable.InvalidInstanceStatus"

	// 没有该VPC网络权限。
	RESOURCEUNAVAILABLE_RESOURCENOPERMISSION = "ResourceUnavailable.ResourceNoPermission"

	// 没有找到实例所属VPC信息。
	RESOURCEUNAVAILABLE_VPCRESOURCENOTFOUND = "ResourceUnavailable.VpcResourceNotFound"

	// 解析参数出错。
	STRUCTPARSEFAILED = "StructParseFailed"

	// 系统错误。出现这种错误时，请联系客服支持。
	SYSTEMERROR = "SystemError"

	// 用户未进行实名认证。
	UNAUTHORIZEDOPERATION_USERHASNOREALNAMEAUTHENTICATION = "UnauthorizedOperation.UserHasNoRealnameAuthentication"

	// 其他未知错误。出现这种错误时，请联系客服支持。
	UNKNOWNERROR = "UnknownError"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// VPC相关错误。
	VPCERROR = "VpcError"
)
