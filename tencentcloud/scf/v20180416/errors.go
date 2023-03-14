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

package v20180416

const (
	// 此产品的特有错误码

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// API网关触发器创建失败。
	FAILEDOPERATION_APIGATEWAY = "FailedOperation.ApiGateway"

	// 创建触发器失败。
	FAILEDOPERATION_APIGW = "FailedOperation.Apigw"

	// 获取Apm InstanceId失败。
	FAILEDOPERATION_APMCONFIGINSTANCEID = "FailedOperation.ApmConfigInstanceId"

	// 当前异步事件状态不支持此操作，请稍后重试。
	FAILEDOPERATION_ASYNCEVENTSTATUS = "FailedOperation.AsyncEventStatus"

	// 登录信息验证失败，token 验证失败。
	FAILEDOPERATION_AUTHFAILURE = "FailedOperation.AuthFailure"

	// 调用 NetDeploy 失败。
	FAILEDOPERATION_CALLNETDEPLOYFAILED = "FailedOperation.CallNetDeployFailed"

	// 请求role信息失败。
	FAILEDOPERATION_CALLROLEFAILED = "FailedOperation.CallRoleFailed"

	// CLS服务未注册。
	FAILEDOPERATION_CLSSERVICEUNREGISTERED = "FailedOperation.ClsServiceUnregistered"

	// CopyAsyncRun 传参异常。
	FAILEDOPERATION_COPYASYNCRUN = "FailedOperation.CopyAsyncRun"

	// 复制函数失败。
	FAILEDOPERATION_COPYFAILED = "FailedOperation.CopyFailed"

	// 不支持复制到该地域。
	FAILEDOPERATION_COPYFUNCTION = "FailedOperation.CopyFunction"

	// 操作COS资源失败。
	FAILEDOPERATION_COS = "FailedOperation.Cos"

	// 创建别名失败。
	FAILEDOPERATION_CREATEALIAS = "FailedOperation.CreateAlias"

	// 操作失败。
	FAILEDOPERATION_CREATEFUNCTION = "FailedOperation.CreateFunction"

	// 创建命名空间失败。
	FAILEDOPERATION_CREATENAMESPACE = "FailedOperation.CreateNamespace"

	// 当前函数状态无法进行此操作。
	FAILEDOPERATION_CREATETRIGGER = "FailedOperation.CreateTrigger"

	// 当前调试状态无法执行此操作。
	FAILEDOPERATION_DEBUGMODESTATUS = "FailedOperation.DebugModeStatus"

	// 调试状态下无法更新执行超时时间。
	FAILEDOPERATION_DEBUGMODEUPDATETIMEOUTFAIL = "FailedOperation.DebugModeUpdateTimeOutFail"

	// 删除别名失败。
	FAILEDOPERATION_DELETEALIAS = "FailedOperation.DeleteAlias"

	// 当前函数状态无法进行此操作，请在函数状态正常时重试。
	FAILEDOPERATION_DELETEFUNCTION = "FailedOperation.DeleteFunction"

	// 删除layer版本失败。
	FAILEDOPERATION_DELETELAYERVERSION = "FailedOperation.DeleteLayerVersion"

	// 无法删除默认Namespace。
	FAILEDOPERATION_DELETENAMESPACE = "FailedOperation.DeleteNamespace"

	// 删除触发器失败。
	FAILEDOPERATION_DELETETRIGGER = "FailedOperation.DeleteTrigger"

	// 当前函数状态无法更新代码，请在状态为正常时更新。
	FAILEDOPERATION_FUNCTIONNAMESTATUSERROR = "FailedOperation.FunctionNameStatusError"

	// 函数在部署中,无法做此操作。
	FAILEDOPERATION_FUNCTIONSTATUSERROR = "FailedOperation.FunctionStatusError"

	// 当前函数版本状态无法进行此操作，请在版本状态为正常时重试。
	FAILEDOPERATION_FUNCTIONVERSIONSTATUSNOTACTIVE = "FailedOperation.FunctionVersionStatusNotActive"

	// 获取别名信息失败。
	FAILEDOPERATION_GETALIAS = "FailedOperation.GetAlias"

	// 获取函数代码地址失败。
	FAILEDOPERATION_GETFUNCTIONADDRESS = "FailedOperation.GetFunctionAddress"

	// InstanceNotFound 实例不存在。
	FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"

	// 当前账号或命名空间处于欠费状态，请在可用时重试。
	FAILEDOPERATION_INSUFFICIENTBALANCE = "FailedOperation.InsufficientBalance"

	// 调用函数失败。
	FAILEDOPERATION_INVOKEFUNCTION = "FailedOperation.InvokeFunction"

	// 命名空间已存在，请勿重复创建。
	FAILEDOPERATION_NAMESPACE = "FailedOperation.Namespace"

	// 服务开通失败。
	FAILEDOPERATION_OPENSERVICE = "FailedOperation.OpenService"

	// 操作冲突。
	FAILEDOPERATION_OPERATIONCONFLICT = "FailedOperation.OperationConflict"

	// 创建定时预置任务失败。
	FAILEDOPERATION_PROVISIONCREATETIMER = "FailedOperation.ProvisionCreateTimer"

	// 删除定时预置任务失败。
	FAILEDOPERATION_PROVISIONDELETETIMER = "FailedOperation.ProvisionDeleteTimer"

	// 预置超过可用。
	FAILEDOPERATION_PROVISIONEDEXCEEDAVAILABLE = "FailedOperation.ProvisionedExceedAvailable"

	// 预置超限。
	FAILEDOPERATION_PROVISIONEDEXCEEDRESERVED = "FailedOperation.ProvisionedExceedReserved"

	// 当前函数版本已有预置任务处于进行中，请稍后重试。
	FAILEDOPERATION_PROVISIONEDINPROGRESS = "FailedOperation.ProvisionedInProgress"

	// 发布layer版本失败。
	FAILEDOPERATION_PUBLISHLAYERVERSION = "FailedOperation.PublishLayerVersion"

	// 当前函数状态无法发布版本，请在状态为正常时发布。
	FAILEDOPERATION_PUBLISHVERSION = "FailedOperation.PublishVersion"

	// 角色不存在。
	FAILEDOPERATION_QCSROLENOTFOUND = "FailedOperation.QcsRoleNotFound"

	// ReservedExceedTotal 总保留超限。
	FAILEDOPERATION_RESERVEDEXCEEDTOTAL = "FailedOperation.ReservedExceedTotal"

	// 当前函数已有保留并发设置任务处于进行中，请稍后重试。
	FAILEDOPERATION_RESERVEDINPROGRESS = "FailedOperation.ReservedInProgress"

	// ServiceClosed 请确认后再操作。
	FAILEDOPERATION_SERVICECLOSED = "FailedOperation.ServiceClosed"

	// Topic不存在。
	FAILEDOPERATION_TOPICNOTEXIST = "FailedOperation.TopicNotExist"

	// 用户并发内存配额设置任务处于进行中，请稍后重试。
	FAILEDOPERATION_TOTALCONCURRENCYMEMORYINPROGRESS = "FailedOperation.TotalConcurrencyMemoryInProgress"

	// 指定的服务未开通，可以提交工单申请开通服务。
	FAILEDOPERATION_UNOPENEDSERVICE = "FailedOperation.UnOpenedService"

	// 更新别名失败。
	FAILEDOPERATION_UPDATEALIAS = "FailedOperation.UpdateAlias"

	// 当前函数状态无法更新代码，请在状态为正常时更新。
	FAILEDOPERATION_UPDATEFUNCTIONCODE = "FailedOperation.UpdateFunctionCode"

	// UpdateFunctionConfiguration操作失败。
	FAILEDOPERATION_UPDATEFUNCTIONCONFIGURATION = "FailedOperation.UpdateFunctionConfiguration"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 创建apigw触发器内部错误。
	INTERNALERROR_APIGATEWAY = "InternalError.ApiGateway"

	// ckafka接口失败。
	INTERNALERROR_CKAFKA = "InternalError.Ckafka"

	// 删除cmq触发器失败。
	INTERNALERROR_CMQ = "InternalError.Cmq"

	// 更新触发器失败。
	INTERNALERROR_COS = "InternalError.Cos"

	// ES错误。
	INTERNALERROR_ES = "InternalError.ES"

	// 内部服务异常。
	INTERNALERROR_EXCEPTION = "InternalError.Exception"

	// 内部服务错误。
	INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"

	// 获取sts票据信息失败。
	INTERNALERROR_GETSTSTOKENFAILED = "InternalError.GetStsTokenFailed"

	// 内部系统错误。
	INTERNALERROR_SYSTEM = "InternalError.System"

	// 内部服务错误。
	INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"

	// 日志参数有误。
	INVALIDPARAMETER_CLS = "InvalidParameter.Cls"

	// FunctionName取值与规范不符，请修正后再试。可参考：https://tencentcs.com/5jXKFnBW。
	INVALIDPARAMETER_FUNCTIONNAME = "InvalidParameter.FunctionName"

	// 创建函数传参异常。
	INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"

	// 请求参数不合法。
	INVALIDPARAMETER_PAYLOAD = "InvalidParameter.Payload"

	// 请求大小超限。
	INVALIDPARAMETER_REQUESTTOOLARGE = "InvalidParameter.RequestTooLarge"

	// RoleCheck 传参有误。
	INVALIDPARAMETER_ROLECHECK = "InvalidParameter.RoleCheck"

	// RoutingConfig参数传入错误。
	INVALIDPARAMETER_ROUTINGCONFIG = "InvalidParameter.RoutingConfig"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// Action取值与规范不符，请修正后再试。可参考：https://tencentcs.com/5jXKFnBW。
	INVALIDPARAMETERVALUE_ACTION = "InvalidParameterValue.Action"

	// AdditionalVersionWeights参数传入错误。
	INVALIDPARAMETERVALUE_ADDITIONALVERSIONWEIGHTS = "InvalidParameterValue.AdditionalVersionWeights"

	// 不支持删除默认别名，请修正后重试。
	INVALIDPARAMETERVALUE_ALIAS = "InvalidParameterValue.Alias"

	// ApiGateway参数错误。
	INVALIDPARAMETERVALUE_APIGATEWAY = "InvalidParameterValue.ApiGateway"

	// ApmConfig参数传入错误。
	INVALIDPARAMETERVALUE_APMCONFIG = "InvalidParameterValue.ApmConfig"

	// ApmConfigInstanceId参数传入错误。
	INVALIDPARAMETERVALUE_APMCONFIGINSTANCEID = "InvalidParameterValue.ApmConfigInstanceId"

	// ApmConfigRegion参数传入错误。
	INVALIDPARAMETERVALUE_APMCONFIGREGION = "InvalidParameterValue.ApmConfigRegion"

	// Args 参数值有误。
	INVALIDPARAMETERVALUE_ARGS = "InvalidParameterValue.Args"

	// AsyncRunEnable 取值不正确。
	INVALIDPARAMETERVALUE_ASYNCRUNENABLE = "InvalidParameterValue.AsyncRunEnable"

	// 函数异步重试配置参数无效。
	INVALIDPARAMETERVALUE_ASYNCTRIGGERCONFIG = "InvalidParameterValue.AsyncTriggerConfig"

	// Cdn传入错误。
	INVALIDPARAMETERVALUE_CDN = "InvalidParameterValue.Cdn"

	// cfs配置项重复。
	INVALIDPARAMETERVALUE_CFSPARAMETERDUPLICATE = "InvalidParameterValue.CfsParameterDuplicate"

	// cfs配置项取值与规范不符。
	INVALIDPARAMETERVALUE_CFSPARAMETERERROR = "InvalidParameterValue.CfsParameterError"

	// cfs参数格式与规范不符。
	INVALIDPARAMETERVALUE_CFSSTRUCTIONERROR = "InvalidParameterValue.CfsStructionError"

	// Ckafka传入错误。
	INVALIDPARAMETERVALUE_CKAFKA = "InvalidParameterValue.Ckafka"

	// 运行函数时的参数传入有误。
	INVALIDPARAMETERVALUE_CLIENTCONTEXT = "InvalidParameterValue.ClientContext"

	// Cls传入错误。
	INVALIDPARAMETERVALUE_CLS = "InvalidParameterValue.Cls"

	// 修改Cls配置需要传入Role参数，请修正后重试。
	INVALIDPARAMETERVALUE_CLSROLE = "InvalidParameterValue.ClsRole"

	// Cmq传入错误。
	INVALIDPARAMETERVALUE_CMQ = "InvalidParameterValue.Cmq"

	// Code传入错误。
	INVALIDPARAMETERVALUE_CODE = "InvalidParameterValue.Code"

	// CodeSecret传入错误。
	INVALIDPARAMETERVALUE_CODESECRET = "InvalidParameterValue.CodeSecret"

	// CodeSource传入错误。
	INVALIDPARAMETERVALUE_CODESOURCE = "InvalidParameterValue.CodeSource"

	// Command[Entrypoint] 参数值有误。
	INVALIDPARAMETERVALUE_COMMAND = "InvalidParameterValue.Command"

	// CompatibleRuntimes参数传入错误。
	INVALIDPARAMETERVALUE_COMPATIBLERUNTIMES = "InvalidParameterValue.CompatibleRuntimes"

	// Content参数传入错误。
	INVALIDPARAMETERVALUE_CONTENT = "InvalidParameterValue.Content"

	// Cos传入错误。
	INVALIDPARAMETERVALUE_COS = "InvalidParameterValue.Cos"

	// CosBucketName不符合规范。
	INVALIDPARAMETERVALUE_COSBUCKETNAME = "InvalidParameterValue.CosBucketName"

	// CosBucketRegion取值与规范不符，请修正后再试。可参考：https://tencentcs.com/5jXKFnBW。
	INVALIDPARAMETERVALUE_COSBUCKETREGION = "InvalidParameterValue.CosBucketRegion"

	// COS通知规则冲突。
	INVALIDPARAMETERVALUE_COSNOTIFYRULECONFLICT = "InvalidParameterValue.CosNotifyRuleConflict"

	// CosObjectName不符合规范。
	INVALIDPARAMETERVALUE_COSOBJECTNAME = "InvalidParameterValue.CosObjectName"

	// CustomArgument参数长度超限。
	INVALIDPARAMETERVALUE_CUSTOMARGUMENT = "InvalidParameterValue.CustomArgument"

	// DateTime传入错误。
	INVALIDPARAMETERVALUE_DATETIME = "InvalidParameterValue.DateTime"

	// DeadLetterConfig取值与规范不符，请修正后再试。可参考：https://tencentcs.com/5jXKFnBW。
	INVALIDPARAMETERVALUE_DEADLETTERCONFIG = "InvalidParameterValue.DeadLetterConfig"

	// 默认Namespace无法创建。
	INVALIDPARAMETERVALUE_DEFAULTNAMESPACE = "InvalidParameterValue.DefaultNamespace"

	// DemoID  对应的函数模板 ，code 参数值有误，请确认后重试。
	INVALIDPARAMETERVALUE_DEMO = "InvalidParameterValue.Demo"

	// DemoId 不存在。
	INVALIDPARAMETERVALUE_DEMOID = "InvalidParameterValue.DemoId"

	// Description传入错误。
	INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"

	// 环境变量DNS[OS_NAMESERVER]配置有误。
	INVALIDPARAMETERVALUE_DNSINFO = "InvalidParameterValue.DnsInfo"

	// DynamicEnabled 参数传入错误。
	INVALIDPARAMETERVALUE_DYNAMICENABLED = "InvalidParameterValue.DynamicEnabled"

	// EipConfig参数错误。
	INVALIDPARAMETERVALUE_EIPCONFIG = "InvalidParameterValue.EipConfig"

	// Enable取值与规范不符，请修正后再试。可参考：https://tencentcs.com/5jXKFnBW。
	INVALIDPARAMETERVALUE_ENABLE = "InvalidParameterValue.Enable"

	// Environment传入错误。
	INVALIDPARAMETERVALUE_ENVIRONMENT = "InvalidParameterValue.Environment"

	// 环境变量大小超限，请保持在 4KB 以内。
	INVALIDPARAMETERVALUE_ENVIRONMENTEXCEEDEDLIMIT = "InvalidParameterValue.EnvironmentExceededLimit"

	// 不支持修改函数系统环境变量和运行环境变量。
	INVALIDPARAMETERVALUE_ENVIRONMENTSYSTEMPROTECT = "InvalidParameterValue.EnvironmentSystemProtect"

	// Filters参数错误。
	INVALIDPARAMETERVALUE_FILTERS = "InvalidParameterValue.Filters"

	// Function取值与规范不符，请修正后再试。可参考：https://tencentcs.com/5jXKFnBW。
	INVALIDPARAMETERVALUE_FUNCTION = "InvalidParameterValue.Function"

	// 函数不存在。
	INVALIDPARAMETERVALUE_FUNCTIONNAME = "InvalidParameterValue.FunctionName"

	// 请求 id 传参错误。
	INVALIDPARAMETERVALUE_FUNCTIONREQUESTID = "InvalidParameterValue.FunctionRequestId"

	// FunctionType参数错误。
	INVALIDPARAMETERVALUE_FUNCTIONTYPE = "InvalidParameterValue.FunctionType"

	// GitBranch不符合规范。
	INVALIDPARAMETERVALUE_GITBRANCH = "InvalidParameterValue.GitBranch"

	// GitCommitId取值与规范不符，请修正后再试。可参考：https://tencentcs.com/5jXKFnBW。
	INVALIDPARAMETERVALUE_GITCOMMITID = "InvalidParameterValue.GitCommitId"

	// GitDirectory不符合规范。
	INVALIDPARAMETERVALUE_GITDIRECTORY = "InvalidParameterValue.GitDirectory"

	// GitPassword不符合规范。
	INVALIDPARAMETERVALUE_GITPASSWORD = "InvalidParameterValue.GitPassword"

	// GitPasswordSecret 传参有误。
	INVALIDPARAMETERVALUE_GITPASSWORDSECRET = "InvalidParameterValue.GitPasswordSecret"

	// GitUrl不符合规范。
	INVALIDPARAMETERVALUE_GITURL = "InvalidParameterValue.GitUrl"

	// GitUserName不符合规范。
	INVALIDPARAMETERVALUE_GITUSERNAME = "InvalidParameterValue.GitUserName"

	// GitUserNameSecret 传参有误。
	INVALIDPARAMETERVALUE_GITUSERNAMESECRET = "InvalidParameterValue.GitUserNameSecret"

	// Handler传入错误。
	INVALIDPARAMETERVALUE_HANDLER = "InvalidParameterValue.Handler"

	// IdleTimeOut参数传入错误。
	INVALIDPARAMETERVALUE_IDLETIMEOUT = "InvalidParameterValue.IdleTimeOut"

	// ImageType 参数值有误。
	INVALIDPARAMETERVALUE_IMAGETYPE = "InvalidParameterValue.ImageType"

	// imageUri 传入有误。
	INVALIDPARAMETERVALUE_IMAGEURI = "InvalidParameterValue.ImageUri"

	// InlineZipFile非法。
	INVALIDPARAMETERVALUE_INLINEZIPFILE = "InvalidParameterValue.InlineZipFile"

	// InstanceConcurrencyConfig 参数传入错误。
	INVALIDPARAMETERVALUE_INSTANCECONCURRENCYCONFIG = "InvalidParameterValue.InstanceConcurrencyConfig"

	// InvokeType取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_INVOKETYPE = "InvalidParameterValue.InvokeType"

	// L5Enable取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_L5ENABLE = "InvalidParameterValue.L5Enable"

	// LayerName参数传入错误。
	INVALIDPARAMETERVALUE_LAYERNAME = "InvalidParameterValue.LayerName"

	// Layers参数传入错误。
	INVALIDPARAMETERVALUE_LAYERS = "InvalidParameterValue.Layers"

	// Limit传入错误。
	INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"

	// 参数超出长度限制。
	INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"

	// MaxConcurrency 参数传入错误。
	INVALIDPARAMETERVALUE_MAXCONCURRENCY = "InvalidParameterValue.MaxConcurrency"

	// Memory取值与规范不符，请修正后再试。可参考：https://tencentcs.com/5jXKFnBW。
	INVALIDPARAMETERVALUE_MEMORY = "InvalidParameterValue.Memory"

	// MemorySize错误。
	INVALIDPARAMETERVALUE_MEMORYSIZE = "InvalidParameterValue.MemorySize"

	// MinCapacity 参数传入错误。
	INVALIDPARAMETERVALUE_MINCAPACITY = "InvalidParameterValue.MinCapacity"

	// Name参数传入错误。
	INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"

	// Namespace参数传入错误。
	INVALIDPARAMETERVALUE_NAMESPACE = "InvalidParameterValue.Namespace"

	// 规则不正确，Namespace为英文字母、数字、-_ 符号组成，长度30。
	INVALIDPARAMETERVALUE_NAMESPACEINVALID = "InvalidParameterValue.NamespaceInvalid"

	// NodeSpec 参数传入错误。
	INVALIDPARAMETERVALUE_NODESPEC = "InvalidParameterValue.NodeSpec"

	// NodeType 参数传入错误。
	INVALIDPARAMETERVALUE_NODETYPE = "InvalidParameterValue.NodeType"

	// 偏移量不合法。
	INVALIDPARAMETERVALUE_OFFSET = "InvalidParameterValue.Offset"

	// Order传入错误。
	INVALIDPARAMETERVALUE_ORDER = "InvalidParameterValue.Order"

	// OrderBy取值与规范不符，请修正后再试。可参考：https://tencentcs.com/5jXKFnBW。
	INVALIDPARAMETERVALUE_ORDERBY = "InvalidParameterValue.OrderBy"

	// 入参不是标准的json。
	INVALIDPARAMETERVALUE_PARAM = "InvalidParameterValue.Param"

	// ProtocolType参数传入错误。
	INVALIDPARAMETERVALUE_PROTOCOLTYPE = "InvalidParameterValue.ProtocolType"

	// 定时预置的cron配置重复。
	INVALIDPARAMETERVALUE_PROVISIONTRIGGERCRONCONFIGDUPLICATE = "InvalidParameterValue.ProvisionTriggerCronConfigDuplicate"

	// TriggerName参数传入错误。
	INVALIDPARAMETERVALUE_PROVISIONTRIGGERNAME = "InvalidParameterValue.ProvisionTriggerName"

	// TriggerName重复。
	INVALIDPARAMETERVALUE_PROVISIONTRIGGERNAMEDUPLICATE = "InvalidParameterValue.ProvisionTriggerNameDuplicate"

	// ProvisionType 参数传入错误。
	INVALIDPARAMETERVALUE_PROVISIONTYPE = "InvalidParameterValue.ProvisionType"

	// PublicNetConfig参数错误。
	INVALIDPARAMETERVALUE_PUBLICNETCONFIG = "InvalidParameterValue.PublicNetConfig"

	// 不支持的函数版本。
	INVALIDPARAMETERVALUE_QUALIFIER = "InvalidParameterValue.Qualifier"

	// 查询版本详情，版本参数传入错误。
	INVALIDPARAMETERVALUE_QUERYVERSION = "InvalidParameterValue.QueryVersion"

	// 企业版镜像实例ID[RegistryId]传值错误。
	INVALIDPARAMETERVALUE_REGISTRYID = "InvalidParameterValue.RegistryId"

	// RetCode不合法。
	INVALIDPARAMETERVALUE_RETCODE = "InvalidParameterValue.RetCode"

	// RoutingConfig取值与规范不符，请修正后再试。可参考：https://tencentcs.com/5jXKFnBW。
	INVALIDPARAMETERVALUE_ROUTINGCONFIG = "InvalidParameterValue.RoutingConfig"

	// Runtime传入错误。
	INVALIDPARAMETERVALUE_RUNTIME = "InvalidParameterValue.Runtime"

	// searchkey 不是 Keyword,Tag 或者 Runtime。
	INVALIDPARAMETERVALUE_SEARCHKEY = "InvalidParameterValue.SearchKey"

	// SecretInfo错误。
	INVALIDPARAMETERVALUE_SECRETINFO = "InvalidParameterValue.SecretInfo"

	// ServiceName命名不规范。
	INVALIDPARAMETERVALUE_SERVICENAME = "InvalidParameterValue.ServiceName"

	// Stamp取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_STAMP = "InvalidParameterValue.Stamp"

	// 起始时间传入错误。
	INVALIDPARAMETERVALUE_STARTTIME = "InvalidParameterValue.StartTime"

	// 需要同时指定开始日期与结束日期。
	INVALIDPARAMETERVALUE_STARTTIMEORENDTIME = "InvalidParameterValue.StartTimeOrEndTime"

	// Status取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_STATUS = "InvalidParameterValue.Status"

	// 系统环境变量错误。
	INVALIDPARAMETERVALUE_SYSTEMENVIRONMENT = "InvalidParameterValue.SystemEnvironment"

	// 非法的TempCosObjectName。
	INVALIDPARAMETERVALUE_TEMPCOSOBJECTNAME = "InvalidParameterValue.TempCosObjectName"

	// TraceEnable取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_TRACEENABLE = "InvalidParameterValue.TraceEnable"

	// TrackingTarget 参数输入错误。
	INVALIDPARAMETERVALUE_TRACKINGTARGET = "InvalidParameterValue.TrackingTarget"

	// TriggerCronConfig参数传入错误。
	INVALIDPARAMETERVALUE_TRIGGERCRONCONFIG = "InvalidParameterValue.TriggerCronConfig"

	// TriggerCronConfig参数定时触发间隔小于指定值。
	INVALIDPARAMETERVALUE_TRIGGERCRONCONFIGTIMEINTERVAL = "InvalidParameterValue.TriggerCronConfigTimeInterval"

	// TriggerDesc传入参数错误。
	INVALIDPARAMETERVALUE_TRIGGERDESC = "InvalidParameterValue.TriggerDesc"

	// TriggerName传入错误。
	INVALIDPARAMETERVALUE_TRIGGERNAME = "InvalidParameterValue.TriggerName"

	// TriggerProvisionedConcurrencyNum参数传入错误。
	INVALIDPARAMETERVALUE_TRIGGERPROVISIONEDCONCURRENCYNUM = "InvalidParameterValue.TriggerProvisionedConcurrencyNum"

	// Type传入错误。
	INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"

	// 开启cfs配置的同时必须开启vpc。
	INVALIDPARAMETERVALUE_VPCNOTSETWHENOPENCFS = "InvalidParameterValue.VpcNotSetWhenOpenCfs"

	// WebSocketsParams参数传入错误。
	INVALIDPARAMETERVALUE_WEBSOCKETSPARAMS = "InvalidParameterValue.WebSocketsParams"

	// 检测到不是标准的zip文件，请重新压缩后再试。
	INVALIDPARAMETERVALUE_ZIPFILE = "InvalidParameterValue.ZipFile"

	// 压缩文件base64解码失败: `Incorrect padding`，请修正后再试。
	INVALIDPARAMETERVALUE_ZIPFILEBASE64BINASCIIERROR = "InvalidParameterValue.ZipFileBase64BinasciiError"

	// 别名个数超过最大限制。
	LIMITEXCEEDED_ALIAS = "LimitExceeded.Alias"

	// Cdn使用超过最大限制。
	LIMITEXCEEDED_CDN = "LimitExceeded.Cdn"

	// 用户开启镜像加速函数版本超限。
	LIMITEXCEEDED_CONTAINERIMAGEACCELERATE = "LimitExceeded.ContainerImageAccelerate"

	// 用户开启镜像加速函数版本超限。
	LIMITEXCEEDED_CONTAINERIMAGEACCELERATEQUOTA = "LimitExceeded.ContainerImageAccelerateQuota"

	// eip资源超限。
	LIMITEXCEEDED_EIP = "LimitExceeded.Eip"

	// 函数数量超出最大限制 ，可通过[提交工单](https://cloud.tencent.com/act/event/Online_service?from=scf%7Cindex)申请提升限制。
	LIMITEXCEEDED_FUNCTION = "LimitExceeded.Function"

	// 同一个主题下的函数超过最大限制。
	LIMITEXCEEDED_FUNCTIONONTOPIC = "LimitExceeded.FunctionOnTopic"

	// FunctionProvisionedConcurrencyMemory数量达到限制，可提交工单申请提升限制：https://tencentcs.com/7Fixwt63。
	LIMITEXCEEDED_FUNCTIONPROVISIONEDCONCURRENCYMEMORY = "LimitExceeded.FunctionProvisionedConcurrencyMemory"

	// 函数保留并发内存超限。
	LIMITEXCEEDED_FUNCTIONRESERVEDCONCURRENCYMEMORY = "LimitExceeded.FunctionReservedConcurrencyMemory"

	// FunctionTotalProvisionedConcurrencyMemory达到限制，可提交工单申请提升限制：https://tencentcs.com/7Fixwt63。
	LIMITEXCEEDED_FUNCTIONTOTALPROVISIONEDCONCURRENCYMEMORY = "LimitExceeded.FunctionTotalProvisionedConcurrencyMemory"

	// 函数预置并发总数达到限制。
	LIMITEXCEEDED_FUNCTIONTOTALPROVISIONEDCONCURRENCYNUM = "LimitExceeded.FunctionTotalProvisionedConcurrencyNum"

	// InitTimeout达到限制，可提交工单申请提升限制：https://tencentcs.com/7Fixwt63。
	LIMITEXCEEDED_INITTIMEOUT = "LimitExceeded.InitTimeout"

	// 内网固定IP个数超限。
	LIMITEXCEEDED_INTRAIP = "LimitExceeded.IntraIp"

	// layer版本数量超出最大限制。
	LIMITEXCEEDED_LAYERVERSIONS = "LimitExceeded.LayerVersions"

	// layer数量超出最大限制。
	LIMITEXCEEDED_LAYERS = "LimitExceeded.Layers"

	// 动态扩容最大值超限。
	LIMITEXCEEDED_MAXCAPACITY = "LimitExceeded.MaxCapacity"

	// 内存超出最大限制。
	LIMITEXCEEDED_MEMORY = "LimitExceeded.Memory"

	// 函数异步重试配置消息保留时间超过限制。
	LIMITEXCEEDED_MSGTTL = "LimitExceeded.MsgTTL"

	// 命名空间数量超过最大限制，可通过[提交工单](https://cloud.tencent.com/act/event/Online_service?from=scf%7Cindex)申请提升限制。
	LIMITEXCEEDED_NAMESPACE = "LimitExceeded.Namespace"

	// Offset超出限制。
	LIMITEXCEEDED_OFFSET = "LimitExceeded.Offset"

	// 定时预置数量超过最大限制。
	LIMITEXCEEDED_PROVISIONTRIGGERACTION = "LimitExceeded.ProvisionTriggerAction"

	// 定时触发间隔小于最大限制。
	LIMITEXCEEDED_PROVISIONTRIGGERINTERVAL = "LimitExceeded.ProvisionTriggerInterval"

	// 配额超限。
	LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"

	// 函数异步重试配置异步重试次数超过限制。
	LIMITEXCEEDED_RETRYNUM = "LimitExceeded.RetryNum"

	// Timeout超出最大限制。
	LIMITEXCEEDED_TIMEOUT = "LimitExceeded.Timeout"

	// 用户并发内存配额超限。
	LIMITEXCEEDED_TOTALCONCURRENCYMEMORY = "LimitExceeded.TotalConcurrencyMemory"

	// 触发器数量超出最大限制，可通过[提交工单](https://cloud.tencent.com/act/event/Online_service?from=scf%7Cindex)申请提升限制。
	LIMITEXCEEDED_TRIGGER = "LimitExceeded.Trigger"

	// UserTotalConcurrencyMemory达到限制，可提交工单申请提升限制：https://tencentcs.com/7Fixwt63。
	LIMITEXCEEDED_USERTOTALCONCURRENCYMEMORY = "LimitExceeded.UserTotalConcurrencyMemory"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// Code没有传入。
	MISSINGPARAMETER_CODE = "MissingParameter.Code"

	// 缺失 Runtime 字段。
	MISSINGPARAMETER_RUNTIME = "MissingParameter.Runtime"

	// 账号已被隔离。
	OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// Alias已被占用。
	RESOURCEINUSE_ALIAS = "ResourceInUse.Alias"

	// Cdn已被占用。
	RESOURCEINUSE_CDN = "ResourceInUse.Cdn"

	// Cmq已被占用。
	RESOURCEINUSE_CMQ = "ResourceInUse.Cmq"

	// Cos已被占用。
	RESOURCEINUSE_COS = "ResourceInUse.Cos"

	// 函数已存在。
	RESOURCEINUSE_FUNCTION = "ResourceInUse.Function"

	// FunctionName已存在。
	RESOURCEINUSE_FUNCTIONNAME = "ResourceInUse.FunctionName"

	// Layer版本正在使用中。
	RESOURCEINUSE_LAYERVERSION = "ResourceInUse.LayerVersion"

	// Namespace已存在。
	RESOURCEINUSE_NAMESPACE = "ResourceInUse.Namespace"

	// TriggerName已存在。
	RESOURCEINUSE_TRIGGER = "ResourceInUse.Trigger"

	// TriggerName已存在。
	RESOURCEINUSE_TRIGGERNAME = "ResourceInUse.TriggerName"

	// COS资源不足。
	RESOURCEINSUFFICIENT_COS = "ResourceInsufficient.COS"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 别名不存在。
	RESOURCENOTFOUND_ALIAS = "ResourceNotFound.Alias"

	// 未找到指定的AsyncEvent，请创建后再试。
	RESOURCENOTFOUND_ASYNCEVENT = "ResourceNotFound.AsyncEvent"

	// 函数需要关联的标签Key不存在。
	RESOURCENOTFOUND_ATTACHEDTAGKEYNOTFOUND = "ResourceNotFound.AttachedTagKeyNotFound"

	// Cdn不存在。
	RESOURCENOTFOUND_CDN = "ResourceNotFound.Cdn"

	// 指定的cfs下未找到您所指定的挂载点。
	RESOURCENOTFOUND_CFSMOUNTINSNOTMATCH = "ResourceNotFound.CfsMountInsNotMatch"

	// CfsProtocolError 参数异常。
	RESOURCENOTFOUND_CFSPROTOCOLERROR = "ResourceNotFound.CfsProtocolError"

	// 检测cfs状态为不可用。
	RESOURCENOTFOUND_CFSSTATUSERROR = "ResourceNotFound.CfsStatusError"

	// cfs与云函数所处vpc不一致。
	RESOURCENOTFOUND_CFSVPCNOTMATCH = "ResourceNotFound.CfsVpcNotMatch"

	// Ckafka不存在。
	RESOURCENOTFOUND_CKAFKA = "ResourceNotFound.Ckafka"

	// Cmq不存在。
	RESOURCENOTFOUND_CMQ = "ResourceNotFound.Cmq"

	// Cos不存在。
	RESOURCENOTFOUND_COS = "ResourceNotFound.Cos"

	// 不存在的Demo。
	RESOURCENOTFOUND_DEMO = "ResourceNotFound.Demo"

	// 函数不存在。
	RESOURCENOTFOUND_FUNCTION = "ResourceNotFound.Function"

	// 函数不存在。
	RESOURCENOTFOUND_FUNCTIONNAME = "ResourceNotFound.FunctionName"

	// 函数版本不存在。
	RESOURCENOTFOUND_FUNCTIONVERSION = "ResourceNotFound.FunctionVersion"

	// 获取cfs挂载点信息错误。
	RESOURCENOTFOUND_GETCFSMOUNTINSERROR = "ResourceNotFound.GetCfsMountInsError"

	// 获取cfs信息错误。
	RESOURCENOTFOUND_GETCFSNOTMATCH = "ResourceNotFound.GetCfsNotMatch"

	// 未找到指定的ImageConfig，请创建后再试。
	RESOURCENOTFOUND_IMAGECONFIG = "ResourceNotFound.ImageConfig"

	// layer不存在。
	RESOURCENOTFOUND_LAYER = "ResourceNotFound.Layer"

	// Layer版本不存在。
	RESOURCENOTFOUND_LAYERVERSION = "ResourceNotFound.LayerVersion"

	// Namespace不存在。
	RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"

	// 版本不存在。
	RESOURCENOTFOUND_QUALIFIER = "ResourceNotFound.Qualifier"

	// 角色不存在。
	RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"

	// Role不存在。
	RESOURCENOTFOUND_ROLECHECK = "ResourceNotFound.RoleCheck"

	// Timer不存在。
	RESOURCENOTFOUND_TIMER = "ResourceNotFound.Timer"

	// 并发内存配额资源未找到。
	RESOURCENOTFOUND_TOTALCONCURRENCYMEMORY = "ResourceNotFound.TotalConcurrencyMemory"

	// 触发器不存在。
	RESOURCENOTFOUND_TRIGGER = "ResourceNotFound.Trigger"

	// 版本不存在。
	RESOURCENOTFOUND_VERSION = "ResourceNotFound.Version"

	// VPC或子网不存在。
	RESOURCENOTFOUND_VPC = "ResourceNotFound.Vpc"

	// 余额不足，请先充值。
	RESOURCEUNAVAILABLE_INSUFFICIENTBALANCE = "ResourceUnavailable.InsufficientBalance"

	// Namespace不可用。
	RESOURCEUNAVAILABLE_NAMESPACE = "ResourceUnavailable.Namespace"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// CAM鉴权失败。
	UNAUTHORIZEDOPERATION_CAM = "UnauthorizedOperation.CAM"

	// 无访问代码权限。
	UNAUTHORIZEDOPERATION_CODESECRET = "UnauthorizedOperation.CodeSecret"

	// 没有权限。
	UNAUTHORIZEDOPERATION_CREATETRIGGER = "UnauthorizedOperation.CreateTrigger"

	// 没有权限的操作。
	UNAUTHORIZEDOPERATION_DELETEFUNCTION = "UnauthorizedOperation.DeleteFunction"

	// 没有权限。
	UNAUTHORIZEDOPERATION_DELETETRIGGER = "UnauthorizedOperation.DeleteTrigger"

	// 不是从控制台调用的该接口。
	UNAUTHORIZEDOPERATION_NOTMC = "UnauthorizedOperation.NotMC"

	// Region错误。
	UNAUTHORIZEDOPERATION_REGION = "UnauthorizedOperation.Region"

	// 没有权限访问您的Cos资源。
	UNAUTHORIZEDOPERATION_ROLE = "UnauthorizedOperation.Role"

	// TempCos的Appid和请求账户的APPID不一致。
	UNAUTHORIZEDOPERATION_TEMPCOSAPPID = "UnauthorizedOperation.TempCosAppid"

	// 无法进行此操作。
	UNAUTHORIZEDOPERATION_UPDATEFUNCTIONCODE = "UnauthorizedOperation.UpdateFunctionCode"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 资源还有别名绑定，不支持当前操作，请解绑别名后重试。
	UNSUPPORTEDOPERATION_ALIASBIND = "UnsupportedOperation.AliasBind"

	// 指定的配置AsyncRunEnable暂不支持，请修正后再试。
	UNSUPPORTEDOPERATION_ASYNCRUNENABLE = "UnsupportedOperation.AsyncRunEnable"

	// Cdn不支持。
	UNSUPPORTEDOPERATION_CDN = "UnsupportedOperation.Cdn"

	// Cos操作不支持。
	UNSUPPORTEDOPERATION_COS = "UnsupportedOperation.Cos"

	// 指定的配置EipFixed暂不支持。
	UNSUPPORTEDOPERATION_EIPFIXED = "UnsupportedOperation.EipFixed"

	// 不支持的地域。
	UNSUPPORTEDOPERATION_NOTSUPPORTREGION = "UnsupportedOperation.NotSupportRegion"

	// 不支持此地域。
	UNSUPPORTEDOPERATION_REGION = "UnsupportedOperation.Region"

	// Trigger操作不支持。
	UNSUPPORTEDOPERATION_TRIGGER = "UnsupportedOperation.Trigger"

	// 指定的配置暂不支持，请修正后再试。
	UNSUPPORTEDOPERATION_UPDATEFUNCTIONEVENTINVOKECONFIG = "UnsupportedOperation.UpdateFunctionEventInvokeConfig"

	// 指定的配置VpcConfig暂不支持。
	UNSUPPORTEDOPERATION_VPCCONFIG = "UnsupportedOperation.VpcConfig"
)
