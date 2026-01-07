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

package v20231024

const (
	// 此产品的特有错误码

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 地址池中地址数量超出限额
	FAILEDOPERATION_ADDREXCEEDEDLIMIT = "FailedOperation.AddrExceededLimit"

	// 自动续费设置失败
	FAILEDOPERATION_AUTORENEWFAILED = "FailedOperation.AutoRenewFailed"

	// 必须存在一条默认解析线路。
	FAILEDOPERATION_CANNOTDELETEDEFAULTLINEERROR = "FailedOperation.CanNotDeleteDefaultLineError"

	// 创建失败。
	FAILEDOPERATION_CREATEERROR = "FailedOperation.CreateError"

	// 创建实例失败
	FAILEDOPERATION_CREATEINSTANCEFAILED = "FailedOperation.CreateInstanceFailed"

	// 创建监控器失败
	FAILEDOPERATION_CREATEMONITORFAILED = "FailedOperation.CreateMonitorFailed"

	// 创建地址池失败
	FAILEDOPERATION_CREATEPOOLFAILED = "FailedOperation.CreatePoolFailed"

	// 创建策略失败
	FAILEDOPERATION_CREATESTRATEGYFAILED = "FailedOperation.CreateStrategyFailed"

	// 下单与支付失败
	FAILEDOPERATION_DEALANDPAYFAILED = "FailedOperation.DealAndPayFailed"

	// 操作删除失败。
	FAILEDOPERATION_DELETEERROR = "FailedOperation.DeleteError"

	// 删除监控器失败
	FAILEDOPERATION_DELETEMONITORFAILED = "FailedOperation.DeleteMonitorFailed"

	// 删除地址池失败
	FAILEDOPERATION_DELETEPOOLFAILED = "FailedOperation.DeletePoolFailed"

	// 删除策略失败
	FAILEDOPERATION_DELETESTRATEGYFAILED = "FailedOperation.DeleteStrategyFailed"

	// 超出限额。
	FAILEDOPERATION_EXCEEDEDLIMIT = "FailedOperation.ExceededLimit"

	// 探点数配额不足。
	FAILEDOPERATION_EXCEEDEDMONITORNUMLIMIT = "FailedOperation.ExceededMonitorNumLimit"

	// 修改错误。
	FAILEDOPERATION_MODIFYERROR = "FailedOperation.ModifyError"

	// 修改实例失败
	FAILEDOPERATION_MODIFYINSTANCEFAILED = "FailedOperation.ModifyInstanceFailed"

	// 修改监控器失败
	FAILEDOPERATION_MODIFYMONITORFAILED = "FailedOperation.ModifyMonitorFailed"

	// 修改地址池失败
	FAILEDOPERATION_MODIFYPOOLFAILED = "FailedOperation.ModifyPoolFailed"

	// 修改策略失败
	FAILEDOPERATION_MODIFYSTRATEGYFAILED = "FailedOperation.ModifyStrategyFailed"

	// 地址池数量超出限额
	FAILEDOPERATION_POOLEXCEEDEDLIMIT = "FailedOperation.PoolExceededLimit"

	// 地址池正在使用。
	FAILEDOPERATION_POOLWASUSINGERROR = "FailedOperation.PoolWasUsingError"

	// 查询失败
	FAILEDOPERATION_QUERYERR = "FailedOperation.QueryErr"

	// 探测任务数量超出限额
	FAILEDOPERATION_TASKEXCEEDEDLIMIT = "FailedOperation.TaskExceededLimit"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 数据库读写错误。
	INTERNALERROR_DBERR = "InternalError.DbErr"

	// 参数格式错误
	INTERNALERROR_PARAMETERERROR = "InternalError.ParameterError"

	// 错误未定义
	INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 实例接入类型不合法
	INVALIDPARAMETER_ACCESSTYPEINVALID = "InvalidParameter.AccessTypeInvalid"

	// 标准版只支持变配到旗舰版
	INVALIDPARAMETER_DEALSTANDARDMODIFYFAILED = "InvalidParameter.DealStandardModifyFailed"

	// 此套餐非标准版，不支持变配到旗舰版
	INVALIDPARAMETER_DEALSTANDARDMODIFYINVALID = "InvalidParameter.DealStandardModifyInvalid"

	// resourceId非探测任务，请确认参数
	INVALIDPARAMETER_DEALTASKMODIFYINVALID = "InvalidParameter.DealTaskModifyInvalid"

	// 变配的探测任务配额小于已使用数，请重新输入
	INVALIDPARAMETER_DEALTASKMODIFYNUMLESS = "InvalidParameter.DealTaskModifyNumLess"

	// 变配的探测任务配额无变化，请重新输入
	INVALIDPARAMETER_DEALTASKMODIFYNUMSAME = "InvalidParameter.DealTaskModifyNumSame"

	// 旗舰版不支持变配
	INVALIDPARAMETER_DEALULTIMATEMODIFYFAILED = "InvalidParameter.DealUltimateModifyFailed"

	// 全局接入域名冲突，当前域名被其他实例使用，请解决占用后重试
	INVALIDPARAMETER_GLOBALACCESSDOMAINCONFLICT = "InvalidParameter.GlobalAccessDomainConflict"

	// 实例不可用，请检查实例状态
	INVALIDPARAMETER_INSTANCEDISABLED = "InvalidParameter.InstanceDisabled"

	// 续费或变配时，资源ID不能为空
	INVALIDPARAMETER_RESOURCEIDEMPTY = "InvalidParameter.ResourceIdEmpty"

	// 创建或续费时，下单时间长度不能为空
	INVALIDPARAMETER_TIMESPANEMPTY = "InvalidParameter.TimeSpanEmpty"

	// 下单时间范围非法，只支持1到120
	INVALIDPARAMETER_TIMESPANINVALID = "InvalidParameter.TimeSpanInvalid"

	// 流量策略不合法
	INVALIDPARAMETER_TRAFFICSTRATEGYINVALID = "InvalidParameter.TrafficStrategyInvalid"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 地址池地址重复
	INVALIDPARAMETERVALUE_POOLADDRDUPLICATED = "InvalidParameterValue.PoolAddrDuplicated"

	// 地址池地址不能为空
	INVALIDPARAMETERVALUE_POOLADDREMPTY = "InvalidParameterValue.PoolAddrEmpty"

	// 地址池地址非法
	INVALIDPARAMETERVALUE_POOLADDRINVALID = "InvalidParameterValue.PoolAddrInvalid"

	// 地址池地址不能混用
	INVALIDPARAMETERVALUE_POOLADDRMIXED = "InvalidParameterValue.PoolAddrMixed"

	// 地址池地址权重不能为空
	INVALIDPARAMETERVALUE_POOLADDRWEIGHTEMPTY = "InvalidParameterValue.PoolAddrWeightEmpty"

	// 流量策略不合法
	INVALIDPARAMETERVALUE_TRAFFICSTRATEGYINVALID = "InvalidParameterValue.TrafficStrategyInvalid"

	// 流量策略为负载均衡且主力集合内有多个地址池时，权重不能为空
	INVALIDPARAMETERVALUE_TRAFFICSTRATEGYWEIGHTEMPTY = "InvalidParameterValue.TrafficStrategyWeightEmpty"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 当前请求过于频繁，请稍后重试
	REQUESTLIMITEXCEEDED_INNERREQUESTLIMITEXCEEDED = "RequestLimitExceeded.InnerRequestLimitExceeded"

	// 没有实例,请先购买实例
	RESOURCEINSUFFICIENT_NOINSTANCE = "ResourceInsufficient.NoInstance"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 查询错误，未查询到相关资源。
	RESOURCENOTFOUND_SEARCHINFO = "ResourceNotFound.SearchInfo"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 未授权,不可操作该实例
	UNAUTHORIZEDOPERATION_INSTANCEUNAUTHORIZED = "UnauthorizedOperation.InstanceUnauthorized"

	// 未授权,主力集合中包含未授权地址池
	UNAUTHORIZEDOPERATION_MAINBINDPOOLUNAUTHORIZED = "UnauthorizedOperation.MainBindPoolUnauthorized"

	// 未授权,不可操作该主力地址池
	UNAUTHORIZEDOPERATION_MAINPOOLUNAUTHORIZED = "UnauthorizedOperation.MainPoolUnauthorized"

	// 未授权,不可操作该地址池地址
	UNAUTHORIZEDOPERATION_POOLADDRESSUNAUTHORIZED = "UnauthorizedOperation.PoolAddressUnauthorized"

	// 未授权,不可操作该地址池
	UNAUTHORIZEDOPERATION_POOLUNAUTHORIZED = "UnauthorizedOperation.PoolUnauthorized"

	// 资源Id与传入的套餐类型不匹配
	UNAUTHORIZEDOPERATION_RESOURCEIDNOTMATCH = "UnauthorizedOperation.ResourceIdNotMatch"

	// 资源Id未授权或已过期，请重新输入
	UNAUTHORIZEDOPERATION_RESOURCEIDUNAUTHORIZED = "UnauthorizedOperation.ResourceIdUnauthorized"

	// 未授权,不可操作该策略
	UNAUTHORIZEDOPERATION_STRATEGYUNAUTHORIZED = "UnauthorizedOperation.StrategyUnauthorized"

	// 资源未授权。
	UNAUTHORIZEDOPERATION_UNAUTHORIZEDERROR = "UnauthorizedOperation.UnauthorizedError"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 接入域名未使用腾讯云 DNSPod 解析，请重新输入
	UNSUPPORTEDOPERATION_ACCESSDOMAINUNAUTHORIZED = "UnsupportedOperation.AccessDomainUnauthorized"

	// 接入域名的套餐不允许此TTL值
	UNSUPPORTEDOPERATION_DOMAINTTLINVALID = "UnsupportedOperation.DomainTtlInvalid"

	// 兜底地址池格式不合法
	UNSUPPORTEDOPERATION_FALLBACKPOOLINVALID = "UnsupportedOperation.FallBackPoolInvalid"

	// 实例名称不能重复
	UNSUPPORTEDOPERATION_INSTANCENAMEDUPLICATE = "UnsupportedOperation.InstanceNameDuplicate"

	// 主力地址池重复
	UNSUPPORTEDOPERATION_MAINBINDPOOLDUPLICATE = "UnsupportedOperation.MainBindPoolDuplicate"

	// 主力地址池为空
	UNSUPPORTEDOPERATION_MAINBINDPOOLEMPTY = "UnsupportedOperation.MainBindPoolEmpty"

	// 主力地址集合数量超过限制
	UNSUPPORTEDOPERATION_MAINBINDPOOLLEVELEXCEED = "UnsupportedOperation.MainBindPoolLevelExceed"

	// 主力地址集合内地址类型不能混用
	UNSUPPORTEDOPERATION_MAINBINDPOOLMIXED = "UnsupportedOperation.MainBindPoolMixed"

	// 主力地址集合数量不合法
	UNSUPPORTEDOPERATION_MAINBINDPOOLNUMINVALID = "UnsupportedOperation.MainBindPoolNumInvalid"

	// 主力地址集合重复
	UNSUPPORTEDOPERATION_MAINPOOLDUPLICATE = "UnsupportedOperation.MainPoolDuplicate"

	// 主力集合不能为空
	UNSUPPORTEDOPERATION_MAINPOOLEMPTY = "UnsupportedOperation.MainPoolEmpty"

	// 主力地址池阈值不合法
	UNSUPPORTEDOPERATION_MAINPOOLSURVIVENUMINVALID = "UnsupportedOperation.MainPoolSurviveNumInvalid"

	// 监控器名已存在
	UNSUPPORTEDOPERATION_MONITORNAMEEXIST = "UnsupportedOperation.MonitorNameExist"

	// 套餐无效，请选择正确套餐
	UNSUPPORTEDOPERATION_PACKAGEUNAUTHORIZED = "UnsupportedOperation.PackageUnauthorized"

	// 地址池名称不能重复
	UNSUPPORTEDOPERATION_POOLNAMEDUPLICATE = "UnsupportedOperation.PoolNameDuplicate"

	// 不支持的查询
	UNSUPPORTEDOPERATION_QUERYNOTALLOW = "UnsupportedOperation.QueryNotAllow"

	// 资源名不能重复。
	UNSUPPORTEDOPERATION_RESOURCENAMEDUPLICATED = "UnsupportedOperation.ResourceNameDuplicated"

	// 主备资源组中资源不存在。
	UNSUPPORTEDOPERATION_RESOURCEPOOLSOURCENOTEXIST = "UnsupportedOperation.ResourcePoolSourceNotExist"

	// 存在重复的解析来源。
	UNSUPPORTEDOPERATION_RESOURCESOURCEDUPLICATED = "UnsupportedOperation.ResourceSourceDuplicated"

	// 策略名称不能重复
	UNSUPPORTEDOPERATION_STRATEGYNAMEDUPLICATE = "UnsupportedOperation.StrategyNameDuplicate"

	// 策略线路重复
	UNSUPPORTEDOPERATION_STRATEGYSOURCEDUPLICATE = "UnsupportedOperation.StrategySourceDuplicate"

	// 策略线路已使用
	UNSUPPORTEDOPERATION_STRATEGYSOURCEINUSED = "UnsupportedOperation.StrategySourceInUsed"

	// 策略线路不在实例套餐中
	UNSUPPORTEDOPERATION_STRATEGYSOURCEUNAUTHORIZED = "UnsupportedOperation.StrategySourceUnauthorized"

	// 策略切换类型不合法
	UNSUPPORTEDOPERATION_STRATEGYSWITCHTYPEINVALID = "UnsupportedOperation.StrategySwitchTypeInvalid"

	// TTL无效,请设置正确的TTL
	UNSUPPORTEDOPERATION_TTLINVALID = "UnsupportedOperation.TTLInvalid"

	// 用户未登录，缺少Uin参数。
	UNSUPPORTEDOPERATION_USERNOTFOUND = "UnsupportedOperation.UserNotFound"
)
