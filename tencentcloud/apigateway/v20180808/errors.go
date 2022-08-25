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

package v20180808

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// DryRun 操作，代表请求将会是成功的，只是多传了 DryRun 参数。
	DRYRUNOPERATION = "DryRunOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 自定义密钥已存在。
	FAILEDOPERATION_ACCESSKEYEXIST = "FailedOperation.AccessKeyExist"

	// 该api绑定了使用计划。
	FAILEDOPERATION_APIBINDENVIRONMEN = "FailedOperation.ApiBindEnvironmen"

	// Apis 绑定了使用计划，请先解绑再试。
	FAILEDOPERATION_APIBINDENVIRONMENT = "FailedOperation.ApiBindEnvironment"

	// api类错误。
	FAILEDOPERATION_APIERROR = "FailedOperation.ApiError"

	// 当前API正在操作中，请稍后再试。
	FAILEDOPERATION_APIINOPERATION = "FailedOperation.ApiInOperation"

	// 证书绑定错误。
	FAILEDOPERATION_CERTIFICATEIDBINDERROR = "FailedOperation.CertificateIdBindError"

	// 该证书为企业证书，待提交。
	FAILEDOPERATION_CERTIFICATEIDENTERPRISEWAITSUBMIT = "FailedOperation.CertificateIdEnterpriseWaitSubmit"

	// 证书不存在或者您没有权限查看，请在ssl证书管理平台上传有效证书。
	FAILEDOPERATION_CERTIFICATEIDERROR = "FailedOperation.CertificateIdError"

	// 证书已过期，请上传其他有效证书。
	FAILEDOPERATION_CERTIFICATEIDEXPIRED = "FailedOperation.CertificateIdExpired"

	// 该证书信息内容为空，请联系证书同事。
	FAILEDOPERATION_CERTIFICATEIDINFOERROR = "FailedOperation.CertificateIdInfoError"

	// 该证书处于审核中状态，请上传其他有效证书。
	FAILEDOPERATION_CERTIFICATEIDUNDERVERIFY = "FailedOperation.CertificateIdUnderVerify"

	// 证书处于未知状态，请联系证书同事查看具体状态。
	FAILEDOPERATION_CERTIFICATEIDUNKNOWNERROR = "FailedOperation.CertificateIdUnknownError"

	// 证书处于审核失败状态，请上传其他有效证书。
	FAILEDOPERATION_CERTIFICATEIDVERIFYFAIL = "FailedOperation.CertificateIdVerifyFail"

	// Apis 证书为空，请上传证书。
	FAILEDOPERATION_CERTIFICATEISNULL = "FailedOperation.CertificateIsNull"

	// 调用CLS服务失败。
	FAILEDOPERATION_CLSERROR = "FailedOperation.ClsError"

	// API文档操作失败。
	FAILEDOPERATION_CODINGERROR = "FailedOperation.CodingError"

	// Apis 环境配置错误。
	FAILEDOPERATION_DEFINEMAPPINGENVIRONMENTERROR = "FailedOperation.DefineMappingEnvironmentError"

	// Apis自定义路径映射不能为空。
	FAILEDOPERATION_DEFINEMAPPINGNOTNULL = "FailedOperation.DefineMappingNotNull"

	// Apis 参数重复。
	FAILEDOPERATION_DEFINEMAPPINGPARAMREPEAT = "FailedOperation.DefineMappingParamRepeat"

	// Apis自定义路径配置错误。
	FAILEDOPERATION_DEFINEMAPPINGPATHERROR = "FailedOperation.DefineMappingPathError"

	// 删除自定义域名错误。
	FAILEDOPERATION_DELETEPATHMAPPINGSETERROR = "FailedOperation.DeletePathMappingSetError"

	// 服务自定义域名错误。
	FAILEDOPERATION_DESCRIBESERVICESUBDOMAINSERROR = "FailedOperation.DescribeServiceSubDomainsError"

	// 该域名已经绑定其他服务。
	FAILEDOPERATION_DOMAINALREADYBINDOTHERSERVICE = "FailedOperation.DomainAlreadyBindOtherService"

	// 该域名已经绑定该服务。
	FAILEDOPERATION_DOMAINALREADYBINDSERVICE = "FailedOperation.DomainAlreadyBindService"

	// 当前域名不满足合规要求，无法进行接入。
	FAILEDOPERATION_DOMAININBLACKLIST = "FailedOperation.DomainInBlackList"

	// 该域名还未在腾讯云备案，请备案之后再绑定域名。
	FAILEDOPERATION_DOMAINNEEDBEIAN = "FailedOperation.DomainNeedBeian"

	// 自定义域名解绑失败，域名未绑定该服务。
	FAILEDOPERATION_DOMAINNOTBINDSERVICE = "FailedOperation.DomainNotBindService"

	// 该域名未配置cname到默认域名，或者解析未生效。
	FAILEDOPERATION_DOMAINRESOLVEERROR = "FailedOperation.DomainResolveError"

	// EIAM返回错误。
	FAILEDOPERATION_EIAMERROR = "FailedOperation.EIAMError"

	// 事件总线绑定/解绑失败。
	FAILEDOPERATION_EBERROR = "FailedOperation.EbError"

	// 域名格式错误。
	FAILEDOPERATION_FORMATERROR = "FailedOperation.FormatError"

	// API文档生成失败。
	FAILEDOPERATION_GENERATEAPIDOCUMENTERROR = "FailedOperation.GenerateApiDocumentError"

	// 获取角色失败，请确认完成API网关相关服务接口授权。
	FAILEDOPERATION_GETROLEERROR = "FailedOperation.GetRoleError"

	// 实例不存在，或者无效状态。
	FAILEDOPERATION_INSTANCENOTEXIST = "FailedOperation.InstanceNotExist"

	// Apis 自定义路径与默认路径冲突。
	FAILEDOPERATION_ISDEFAULTMAPPING = "FailedOperation.IsDefaultMapping"

	// 默认的二级域名输入错误。
	FAILEDOPERATION_NETSUBDOMAINERROR = "FailedOperation.NetSubDomainError"

	// 操作后端通道失败。
	FAILEDOPERATION_OPERATEUPSTREAM = "FailedOperation.OperateUpstream"

	// 设置路径映射错误。
	FAILEDOPERATION_PATHMAPPINGSETERROR = "FailedOperation.PathMappingSetError"

	// SCF触发器绑定/解绑操作失败。
	FAILEDOPERATION_SCFERROR = "FailedOperation.ScfError"

	// 服务相关的错误。
	FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"

	// 当前Service正在操作中, 请稍后再试。
	FAILEDOPERATION_SERVICEINOPERATION = "FailedOperation.ServiceInOperation"

	// 服务不存在。
	FAILEDOPERATION_SERVICENOTEXIST = "FailedOperation.ServiceNotExist"

	// 自定义域名已使用默认路径映射，不支持设置自定义路径。
	FAILEDOPERATION_SETCUSTOMPATHMAPPINGERROR = "FailedOperation.SetCustomPathMappingError"

	// 自定义域名格式错误。
	FAILEDOPERATION_SUBDOMAINFORMATERROR = "FailedOperation.SubDomainFormatError"

	// 标签绑定服务失败。
	FAILEDOPERATION_TAGBINDSERVICEERROR = "FailedOperation.TagBindServiceError"

	// 协议类型错误。
	FAILEDOPERATION_UNKNOWNPROTOCOLTYPEERROR = "FailedOperation.UnknownProtocolTypeError"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// API网关内部请求错误，请稍后重试。若无法解决，请联系智能客服或提交工单。
	INTERNALERROR_APIGWEXCEPTION = "InternalError.ApigwException"

	// CAuth内部请求错误，请稍后重试。若无法解决，请联系智能客服或提交工单。
	INTERNALERROR_CAUTHEXCEPTION = "InternalError.CauthException"

	// CLB内部请求错误，请稍后重试。若无法解决，请联系智能客服或提交工单。
	INTERNALERROR_CLBEXCEPTION = "InternalError.ClbException"

	// oss内部请求错误，请稍后重试。若无法解决，请联系智能客服或提交工单。
	INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"

	// SCF内部请求错误，请稍后重试。若无法解决，请联系智能客服或提交工单。
	INTERNALERROR_SCFEXCEPTION = "InternalError.ScfException"

	// TSF内部请求错误，请稍后重试。若无法解决，请联系智能客服或提交工单。
	INTERNALERROR_TSFEXCEPTION = "InternalError.TsfException"

	// vpc内部请求错误，请稍后重试。若无法解决，请联系智能客服或提交工单。
	INTERNALERROR_VPCEXCEPTION = "InternalError.VpcException"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 基础版服务不能绑定插件。
	INVALIDPARAMETER_BASICSERVICENOTALLOWATTACHPLUGIN = "InvalidParameter.BasicServiceNotAllowAttachPlugin"

	// 当前绑定插件和已有插件冲突。
	INVALIDPARAMETER_DUPLICATEPLUGINCONFIG = "InvalidParameter.DuplicatePluginConfig"

	// 参数格式错误。
	INVALIDPARAMETER_FORMATERROR = "InvalidParameter.FormatError"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 插件定义参数取值重复，请修改后重新操作。
	INVALIDPARAMETERVALUE_DUPLICATEPLUGINCONFIG = "InvalidParameterValue.DuplicatePluginConfig"

	// 非法的后端ip地址。
	INVALIDPARAMETERVALUE_ILLEGALPROXYIP = "InvalidParameterValue.IllegalProxyIp"

	// 密钥错误。
	INVALIDPARAMETERVALUE_INVALIDACCESSKEYIDS = "InvalidParameterValue.InvalidAccessKeyIds"

	// 传入的Api业务类型必须为OAUTH。
	INVALIDPARAMETERVALUE_INVALIDAPIBUSINESSTYPE = "InvalidParameterValue.InvalidApiBusinessType"

	// API Id错误。
	INVALIDPARAMETERVALUE_INVALIDAPIIDS = "InvalidParameterValue.InvalidApiIds"

	// 无效的API配置。
	INVALIDPARAMETERVALUE_INVALIDAPIREQUESTCONFIG = "InvalidParameterValue.InvalidApiRequestConfig"

	// API类型错误，微服务API只支持TSF后端服务类型。
	INVALIDPARAMETERVALUE_INVALIDAPITYPE = "InvalidParameterValue.InvalidApiType"

	// 后端服务路径配置错误。
	INVALIDPARAMETERVALUE_INVALIDBACKENDPATH = "InvalidParameterValue.InvalidBackendPath"

	// 不合法的clb。
	INVALIDPARAMETERVALUE_INVALIDCLB = "InvalidParameterValue.InvalidClb"

	// 策略错误。
	INVALIDPARAMETERVALUE_INVALIDCONDITION = "InvalidParameterValue.InvalidCondition"

	// 不合法的常量参数。
	INVALIDPARAMETERVALUE_INVALIDCONSTANTPARAMETERS = "InvalidParameterValue.InvalidConstantParameters"

	// 参数Env取值错误，取值范围为（release, prepub, test）。
	INVALIDPARAMETERVALUE_INVALIDENV = "InvalidParameterValue.InvalidEnv"

	// 服务当前环境状态，不支持此操作。
	INVALIDPARAMETERVALUE_INVALIDENVSTATUS = "InvalidParameterValue.InvalidEnvStatus"

	// 参数取值错误。
	INVALIDPARAMETERVALUE_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameterValue.InvalidFilterNotSupportedName"

	// 参数GenLanguage取值错误。
	INVALIDPARAMETERVALUE_INVALIDGENLANGUAGE = "InvalidParameterValue.InvalidGenLanguage"

	// 参数后端地址取值错误。
	INVALIDPARAMETERVALUE_INVALIDIPADDRESS = "InvalidParameterValue.InvalidIPAddress"

	// 参数请求配额总数取值错误。
	INVALIDPARAMETERVALUE_INVALIDMAXREQUESTNUM = "InvalidParameterValue.InvalidMaxRequestNum"

	// 方法错误。仅支持 ANY, BEGIN, GET, POST, DELETE, HEAD, PUT, OPTIONS, TRACE, PATCH，请修改后重新操作。
	INVALIDPARAMETERVALUE_INVALIDMETHOD = "InvalidParameterValue.InvalidMethod"

	// 插件定义参数取值错误，请修改后重新操作。
	INVALIDPARAMETERVALUE_INVALIDPLUGINCONFIG = "InvalidParameterValue.InvalidPluginConfig"

	// 后端服务端口错误。
	INVALIDPARAMETERVALUE_INVALIDPORT = "InvalidParameterValue.InvalidPort"

	// 无效的协议类型参数。
	INVALIDPARAMETERVALUE_INVALIDPROCOTOL = "InvalidParameterValue.InvalidProcotol"

	// OAUTH2.0 API 公钥参数格式错误。
	INVALIDPARAMETERVALUE_INVALIDPUBLICKEY = "InvalidParameterValue.InvalidPublicKey"

	// 地域错误。
	INVALIDPARAMETERVALUE_INVALIDREGION = "InvalidParameterValue.InvalidRegion"

	// 不合法的请求参数。
	INVALIDPARAMETERVALUE_INVALIDREQUESTPARAMETERS = "InvalidParameterValue.InvalidRequestParameters"

	// SCF类型API参数错误。
	INVALIDPARAMETERVALUE_INVALIDSCFCONFIG = "InvalidParameterValue.InvalidScfConfig"

	// 不合法的服务配置。
	INVALIDPARAMETERVALUE_INVALIDSERVICECONFIG = "InvalidParameterValue.InvalidServiceConfig"

	// 参数ServiceMockReturnMessage错误。
	INVALIDPARAMETERVALUE_INVALIDSERVICEMOCKRETURNMESSAGE = "InvalidParameterValue.InvalidServiceMockReturnMessage"

	// 参数配置错误，未配置在后端参数配置列表中。
	INVALIDPARAMETERVALUE_INVALIDSERVICEPARAM = "InvalidParameterValue.InvalidServiceParam"

	// 参数ServiceParameters错误。
	INVALIDPARAMETERVALUE_INVALIDSERVICEPARAMETERS = "InvalidParameterValue.InvalidServiceParameters"

	// 后端服务类型错误。WEBSOCKET类型的后端服务，前端协议需要配置为WEBSOCKET。
	INVALIDPARAMETERVALUE_INVALIDSERVICETYPE = "InvalidParameterValue.InvalidServiceType"

	// 标签参数错误。
	INVALIDPARAMETERVALUE_INVALIDTAGVALUES = "InvalidParameterValue.InvalidTagValues"

	// TSF类型API配置错误。
	INVALIDPARAMETERVALUE_INVALIDTSFCONFIG = "InvalidParameterValue.InvalidTsfConfig"

	// 当前后端通道不存在或者是无效状态。
	INVALIDPARAMETERVALUE_INVALIDUPSTREAM = "InvalidParameterValue.InvalidUpstream"

	// URL参数错误。
	INVALIDPARAMETERVALUE_INVALIDURL = "InvalidParameterValue.InvalidUrl"

	// 无效vpc信息。
	INVALIDPARAMETERVALUE_INVALIDVPCCONFIG = "InvalidParameterValue.InvalidVpcConfig"

	// 方法错误。WEBSOCKET类型的后端服务，前端方法需要配置为GET。
	INVALIDPARAMETERVALUE_INVALIDWSMETHOD = "InvalidParameterValue.InvalidWSMethod"

	// 参数的长度超出限制。
	INVALIDPARAMETERVALUE_LENGTHEXCEEDED = "InvalidParameterValue.LengthExceeded"

	// 参数个数超出限制，上限为100。
	INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"

	// 参数取值错误。当前值不在可选范围内。
	INVALIDPARAMETERVALUE_NOTINOPTIONS = "InvalidParameterValue.NotInOptions"

	// OAuth接口无可用修改项。
	INVALIDPARAMETERVALUE_NOTHINGMODIFYFOROAUTH = "InvalidParameterValue.NothingModifyForOauth"

	// 后端参数配置错误，未配置相关前端参数。
	INVALIDPARAMETERVALUE_PARAMETERNOTMATCH = "InvalidParameterValue.ParameterNotMatch"

	// 参数个数超过限制。
	INVALIDPARAMETERVALUE_PARAMETERVALUELIMITEXCEEDED = "InvalidParameterValue.ParameterValueLimitExceeded"

	// 参数取值错误，不在参数范围内。
	INVALIDPARAMETERVALUE_RANGEEXCEEDED = "InvalidParameterValue.RangeExceeded"

	// 参数取值错误。
	INVALIDPARAMETERVALUE_UNSUPPORTEDPARAMETER = "InvalidParameterValue.UnsupportedParameter"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// API文档数量超出限制。
	LIMITEXCEEDED_APIDOCLIMITEXCEEDED = "LimitExceeded.APIDocLimitExceeded"

	// 访问密钥数量超出限制。
	LIMITEXCEEDED_ACCESSKEYCOUNTINUSAGEPLANLIMITEXCEEDED = "LimitExceeded.AccessKeyCountInUsagePlanLimitExceeded"

	// 服务绑定的应用数量超过限制。
	LIMITEXCEEDED_APIAPPCOUNTLIMITEXCEEDED = "LimitExceeded.ApiAppCountLimitExceeded"

	// API数量超过限制。
	LIMITEXCEEDED_APICOUNTLIMITEXCEEDED = "LimitExceeded.ApiCountLimitExceeded"

	// 密钥数量超出限制。
	LIMITEXCEEDED_APIKEYCOUNTLIMITEXCEEDED = "LimitExceeded.ApiKeyCountLimitExceeded"

	// 自定义路径映射，最多三组。
	LIMITEXCEEDED_EXCEEDEDDEFINEMAPPINGLIMIT = "LimitExceeded.ExceededDefineMappingLimit"

	// 绑定域名数量超出限制，默认最多5个。
	LIMITEXCEEDED_EXCEEDEDDOMAINLIMIT = "LimitExceeded.ExceededDomainLimit"

	// IP策略数量超出限制。
	LIMITEXCEEDED_IPSTRATEGYLIMITEXCEEDED = "LimitExceeded.IpStrategyLimitExceeded"

	// 请求频率超限制，请稍等再试。
	LIMITEXCEEDED_REQUESTLIMITEXCEEDED = "LimitExceeded.RequestLimitExceeded"

	// 插件绑定的服务数量超出限制，请解绑服务或提升限额后重试。
	LIMITEXCEEDED_SERVICECOUNTFORPLUGINLIMITEXCEEDED = "LimitExceeded.ServiceCountForPluginLimitExceeded"

	// 服务数量超过限制，请删除服务或提升限额后重试。
	LIMITEXCEEDED_SERVICECOUNTLIMITEXCEEDED = "LimitExceeded.ServiceCountLimitExceeded"

	// 使用计划数量超出限制。
	LIMITEXCEEDED_USAGEPLANLIMITEXCEEDED = "LimitExceeded.UsagePlanLimitExceeded"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 缺少后端服务参数。
	MISSINGPARAMETER_BACKENDSPECIFICPARAM = "MissingParameter.BackendSpecificParam"

	// 插件缺少参数，请按规则重新操作。
	MISSINGPARAMETER_PLUGINCONFIG = "MissingParameter.PluginConfig"

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

	// 密钥不存在。
	RESOURCENOTFOUND_INVALIDACCESSKEYID = "ResourceNotFound.InvalidAccessKeyId"

	// ApiId错误。
	RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"

	// 应用ID错误。
	RESOURCENOTFOUND_INVALIDAPIAPP = "ResourceNotFound.InvalidApiApp"

	// API文档不存在。
	RESOURCENOTFOUND_INVALIDAPIDOC = "ResourceNotFound.InvalidApiDoc"

	// IP策略不存在。
	RESOURCENOTFOUND_INVALIDIPSTRATEGY = "ResourceNotFound.InvalidIPStrategy"

	// OAuth业务API错误。
	RESOURCENOTFOUND_INVALIDOAUTHAPI = "ResourceNotFound.InvalidOauthApi"

	// 插件不存在。
	RESOURCENOTFOUND_INVALIDPLUGIN = "ResourceNotFound.InvalidPlugin"

	// 对应服务不可见。
	RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"

	// 使用计划不存在。
	RESOURCENOTFOUND_INVALIDUSAGEPLAN = "ResourceNotFound.InvalidUsagePlan"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 资源售罄。
	RESOURCESSOLDOUT = "ResourcesSoldOut"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 您没有权限访问资源。
	UNAUTHORIZEDOPERATION_ACCESSRESOURCE = "UnauthorizedOperation.AccessResource"

	// 用户未实名制。
	UNAUTHORIZEDOPERATION_UNCERTIFIEDUSER = "UnauthorizedOperation.UncertifiedUser"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 账号余额不足
	UNSUPPORTEDOPERATION_ACCOUNTARREARS = "UnsupportedOperation.AccountArrears"

	// 密钥已绑定使用计划。
	UNSUPPORTEDOPERATION_ALREADYBINDUSAGEPLAN = "UnsupportedOperation.AlreadyBindUsagePlan"

	// 当前插件不支持绑定。
	UNSUPPORTEDOPERATION_ATTACHPLUGIN = "UnsupportedOperation.AttachPlugin"

	// 基础版服务不能创建超过一个API。
	UNSUPPORTEDOPERATION_BASICSERVICENOMOREAPI = "UnsupportedOperation.BasicServiceNoMoreApi"

	// 日志检索起始时间间隔。
	UNSUPPORTEDOPERATION_CLSSEARCHTIME = "UnsupportedOperation.ClsSearchTime"

	// 协议为HTTP时，不支持强制Https。
	UNSUPPORTEDOPERATION_FORCEHTTPS = "UnsupportedOperation.ForceHttps"

	// 接口错误。
	UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"

	// 不支持后端类型。
	UNSUPPORTEDOPERATION_INVALIDENDPOINTTYPE = "UnsupportedOperation.InvalidEndpointType"

	// 当前实例状态，不支持当前操作。
	UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"

	// 当前服务处于隔离中，暂不支持此操作。
	UNSUPPORTEDOPERATION_INVALIDSERVICETRADE = "UnsupportedOperation.InvalidServiceTrade"

	// 当前密钥状态不支持此操作。
	UNSUPPORTEDOPERATION_INVALIDSTATUS = "UnsupportedOperation.InvalidStatus"

	// 不支持修改EIAM类型授权API。
	UNSUPPORTEDOPERATION_MODIFYEIAMAUTHAPI = "UnsupportedOperation.ModifyEIAMAuthApi"

	// 网络类型不支持修改。
	UNSUPPORTEDOPERATION_MODIFYNETTYPE = "UnsupportedOperation.ModifyNetType"

	// 前端协议类型不支持修改。
	UNSUPPORTEDOPERATION_MODIFYPROTOCOL = "UnsupportedOperation.ModifyProtocol"

	// 当前使用计划未绑定环境。
	UNSUPPORTEDOPERATION_NOUSAGEPLANENV = "UnsupportedOperation.NoUsagePlanEnv"

	// 不支持减少网络类型的操作。
	UNSUPPORTEDOPERATION_REDUCENETTYPES = "UnsupportedOperation.ReduceNetTypes"

	// 资源已关联，请先解除。
	UNSUPPORTEDOPERATION_RESOURCEASSOCIATED = "UnsupportedOperation.ResourceAssociated"

	// 密钥已绑定使用计划，请先解绑再试。
	UNSUPPORTEDOPERATION_RESOURCEISINUSE = "UnsupportedOperation.ResourceIsInUse"

	// 资源未关联。
	UNSUPPORTEDOPERATION_RESOURCEUNASSOCIATED = "UnsupportedOperation.ResourceUnassociated"

	// 当前Uin未在手工密钥白名单列表内。
	UNSUPPORTEDOPERATION_UINNOTINWHITELIST = "UnsupportedOperation.UinNotInWhiteList"

	// 密钥已绑定使用计划。
	UNSUPPORTEDOPERATION_UNSUPPORTEDBINDAPIKEY = "UnsupportedOperation.UnsupportedBindApiKey"

	// 不支持绑定环境。
	UNSUPPORTEDOPERATION_UNSUPPORTEDBINDENVIRONMENT = "UnsupportedOperation.UnsupportedBindEnvironment"

	// 当前API已绑定业务API，请先解绑再试。
	UNSUPPORTEDOPERATION_UNSUPPORTEDDELETEAPI = "UnsupportedOperation.UnsupportedDeleteApi"

	// 服务使用中，不能被删除。
	UNSUPPORTEDOPERATION_UNSUPPORTEDDELETESERVICE = "UnsupportedOperation.UnsupportedDeleteService"

	// 当前VPC通道存在绑定关系，不支持当前操作。
	UNSUPPORTEDOPERATION_UNSUPPORTEDDELETEUPSTREAM = "UnsupportedOperation.UnsupportedDeleteUpstream"

	// 不支持网络类型。
	UNSUPPORTEDOPERATION_UNSUPPORTEDNETTYPE = "UnsupportedOperation.UnsupportedNetType"

	// 不支持解绑环境，存在未绑定使用计划的API。
	UNSUPPORTEDOPERATION_UNSUPPORTEDUNBINDENVIRONMENT = "UnsupportedOperation.UnsupportedUnBindEnvironment"

	// 当前密钥状态不支持此操作。
	UNSUPPORTEDOPERATION_UNSUPPORTEDUPDATEAPIKEY = "UnsupportedOperation.UnsupportedUpdateApiKey"

	// 当前使用计划在使用中，请先解绑再试。
	UNSUPPORTEDOPERATION_USAGEPLANINUSE = "UnsupportedOperation.UsagePlanInUse"
)
