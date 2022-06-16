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

package v20180606

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 鉴权错误，请确认后重试。
	AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 域名配置更新操作失败，请重试或联系客服人员解决。
	FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 内部鉴权系统错误。
	INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"

	// 域名配置更新失败。
	INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"

	// 内部数据错误，请联系腾讯云工程师进一步排查。
	INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"

	// 内部查询错误，请重试或联系客服人员解决。
	INTERNALERROR_CDNQUERYPARAMERROR = "InternalError.CdnQueryParamError"

	// 内部查询错误，请重试或联系客服人员解决。
	INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"

	// 系统错误，请联系腾讯云工程师进一步排查。
	INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"

	// 日志服务内部错误。
	INTERNALERROR_CLSINTERNALERROR = "InternalError.ClsInternalError"

	// 计费数据内部查询错误，请重试或联系客服人员解决。
	INTERNALERROR_COSTDATASYSTEMERROR = "InternalError.CostDataSystemError"

	// 数据查询错误，请联系腾讯云工程师进一步排查。
	INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"

	// 内部服务错误，请联系腾讯云工程师进一步排查。
	INTERNALERROR_ERROR = "InternalError.Error"

	// 内部服务错误，请联系腾讯云工程师进一步排查。
	INTERNALERROR_INVALIDERRORCODE = "InternalError.InvalidErrorCode"

	// 内部服务错误，请联系腾讯云工程师进一步排查。
	INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"

	// 内部服务错误，请联系腾讯云工程师进一步排查。
	INTERNALERROR_ROUTEERROR = "InternalError.RouteError"

	// SCDN服务未生效，请购买或续费SCDN套餐后重试。
	INTERNALERROR_SCDNUSERNOPACKAGE = "InternalError.ScdnUserNoPackage"

	// 安全加速服务已停服，请重新购买套餐后开启。
	INTERNALERROR_SCDNUSERSUSPEND = "InternalError.ScdnUserSuspend"

	// 内部数据错误，请重试或联系客服人员解决。
	INTERNALERROR_SYSTEMDBERROR = "InternalError.SystemDBError"

	// 内部服务错误，请联系腾讯云工程师进一步排查。
	INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"

	// 标签内部错误，请重试或联系客服人员解决。
	INTERNALERROR_TAGSYSTEMERROR = "InternalError.TagSystemError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 域名启用 HTTPS 配置需保持访问端口配置-443端口为开启状态。
	INVALIDPARAMETER_ACCESSPORTOPENEDHTTPS = "InvalidParameter.AccessPortOpenedHttps"

	// 请删除域名的限流管理配置后再切换加速区域。
	INVALIDPARAMETER_BANDLIMITREQUIREDMAINLAND = "InvalidParameter.BandLimitRequiredMainland"

	// 源站类型为COS源或第三方对象存储的域名，用量封顶-超出阈值的处理方式仅支持访问返回404，请修改该配置后重试。
	INVALIDPARAMETER_BANDWIDTHALERTCOUNTERMEASURECONFLICTORIGINTYPE = "InvalidParameter.BandwidthAlertCounterMeasureConflictOriginType"

	// 域名状态不合法。
	INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"

	// 同一次请求的资源AppId不一致。
	INVALIDPARAMETER_CAMRESOURCEBELONGTODIFFERENTUSER = "InvalidParameter.CamResourceBelongToDifferentUser"

	// 资源六段式标记参数错误。
	INVALIDPARAMETER_CAMRESOURCESIXSTAGEERROR = "InvalidParameter.CamResourceSixStageError"

	// 域名已与该标签关联，请勿重复操作。
	INVALIDPARAMETER_CAMTAGKEYALREADYATTACHED = "InvalidParameter.CamTagKeyAlreadyAttached"

	// 标签键字符不合法。
	INVALIDPARAMETER_CAMTAGKEYILLEGAL = "InvalidParameter.CamTagKeyIllegal"

	// 标签键不存在。
	INVALIDPARAMETER_CAMTAGKEYNOTEXIST = "InvalidParameter.CamTagKeyNotExist"

	// 标签值字符不合法。
	INVALIDPARAMETER_CAMTAGVALUEILLEGAL = "InvalidParameter.CamTagValueIllegal"

	// 证书信息不存在或非法，请确认后重试。
	INVALIDPARAMETER_CDNCERTINFONOTFOUND = "InvalidParameter.CdnCertInfoNotFound"

	// 证书无效，请确认后重试。
	INVALIDPARAMETER_CDNCERTNOCERTINFO = "InvalidParameter.CdnCertNoCertInfo"

	// HTTPS证书无效。
	INVALIDPARAMETER_CDNCERTNOTPEM = "InvalidParameter.CdnCertNotPem"

	// 存在重复主题。
	INVALIDPARAMETER_CDNCLSDUPLICATETOPIC = "InvalidParameter.CdnClsDuplicateTopic"

	// 主题名字不合法。
	INVALIDPARAMETER_CDNCLSTOPICNAMEINVALID = "InvalidParameter.CdnClsTopicNameInvalid"

	// CLS主题不存在。
	INVALIDPARAMETER_CDNCLSTOPICNOTEXIST = "InvalidParameter.CdnClsTopicNotExist"

	// 缓存配置不合法，请确认后重试。
	INVALIDPARAMETER_CDNCONFIGINVALIDCACHE = "InvalidParameter.CdnConfigInvalidCache"

	// 域名不符合规范，请确认域名是否符合规范。
	INVALIDPARAMETER_CDNCONFIGINVALIDHOST = "InvalidParameter.CdnConfigInvalidHost"

	// 标签配置不合法。
	INVALIDPARAMETER_CDNCONFIGINVALIDTAG = "InvalidParameter.CdnConfigInvalidTag"

	// 域名添加失败，当前域名必须选择标签，请确认后重试。
	INVALIDPARAMETER_CDNCONFIGTAGREQUIRED = "InvalidParameter.CdnConfigTagRequired"

	// 域名拥有特殊配置，需人工处理。
	INVALIDPARAMETER_CDNHOSTHASSPECIALCONFIG = "InvalidParameter.CdnHostHasSpecialConfig"

	// 该域名属于指定账号域名，不允许接入。
	INVALIDPARAMETER_CDNHOSTINTERNALHOST = "InvalidParameter.CdnHostInternalHost"

	// 错误的中间源配置。
	INVALIDPARAMETER_CDNHOSTINVALIDMIDDLECONFIG = "InvalidParameter.CdnHostInvalidMiddleConfig"

	// 域名格式不合法，请确认后重试。
	INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"

	// 域名状态不合法。
	INVALIDPARAMETER_CDNHOSTINVALIDSTATUS = "InvalidParameter.CdnHostInvalidStatus"

	// 该域名为COS访问域名，无法接入，如需启动加速服务，请前往COS控制台启用默认 CDN 加速域。
	INVALIDPARAMETER_CDNHOSTISCOSDEFAULTACCESS = "InvalidParameter.CdnHostIsCosDefaultAccess"

	// 域名太长。
	INVALIDPARAMETER_CDNHOSTTOOLONGHOST = "InvalidParameter.CdnHostTooLongHost"

	// 内部接口错误，请联系腾讯云工程师进一步排查。
	INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"

	// 参数Interval错误，请确认后重试。
	INVALIDPARAMETER_CDNINVALIDPARAMINTERVAL = "InvalidParameter.CdnInvalidParamInterval"

	// 参数Metric错误，请检查后重试。
	INVALIDPARAMETER_CDNINVALIDPARAMMETRIC = "InvalidParameter.CdnInvalidParamMetric"

	// ServiceType字段不合法，请确认后重试。
	INVALIDPARAMETER_CDNINVALIDPARAMSERVICETYPE = "InvalidParameter.CdnInvalidParamServiceType"

	// 配置暂不支持开启该配置。
	INVALIDPARAMETER_CDNKEYRULESEXCLUDECUSTOMREQUIRESFULLLEGO = "InvalidParameter.CdnKeyRulesExcludeCustomRequiresFullLego"

	// QueryString字段不合法，请确认后重试。
	INVALIDPARAMETER_CDNKEYRULESINVALIDQUERYSTRINGVALUE = "InvalidParameter.CdnKeyRulesInvalidQueryStringValue"

	// 参数错误，请参考文档中示例参数填充。
	INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"

	// 刷新不支持泛域名。
	INVALIDPARAMETER_CDNPURGEWILDCARDNOTALLOWED = "InvalidParameter.CdnPurgeWildcardNotAllowed"

	// 预热不支持泛域名。
	INVALIDPARAMETER_CDNPUSHWILDCARDNOTALLOWED = "InvalidParameter.CdnPushWildcardNotAllowed"

	// 日期不合法，请参考文档中日期示例。
	INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"

	// 统计维度不合法，请参考文档中统计分析示例。
	INVALIDPARAMETER_CDNSTATINVALIDFILTER = "InvalidParameter.CdnStatInvalidFilter"

	// 统计类型不合法，请参考文档中统计分析示例。
	INVALIDPARAMETER_CDNSTATINVALIDMETRIC = "InvalidParameter.CdnStatInvalidMetric"

	// 项目ID错误，请确认后重试。
	INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"

	// 查询的域名数量超过限制。
	INVALIDPARAMETER_CDNSTATTOOMANYDOMAINS = "InvalidParameter.CdnStatTooManyDomains"

	// URL 超过限制长度。
	INVALIDPARAMETER_CDNURLEXCEEDLENGTH = "InvalidParameter.CdnUrlExceedLength"

	// 索引冲突。
	INVALIDPARAMETER_CLSINDEXCONFLICT = "InvalidParameter.ClsIndexConflict"

	// 索引规则为空。
	INVALIDPARAMETER_CLSINDEXRULEEMPTY = "InvalidParameter.ClsIndexRuleEmpty"

	// 无效内容。
	INVALIDPARAMETER_CLSINVALIDCONTENT = "InvalidParameter.ClsInvalidContent"

	// 无效的Content-Type。
	INVALIDPARAMETER_CLSINVALIDCONTENTTYPE = "InvalidParameter.ClsInvalidContentType"

	// 参数错误，请检查后重试。
	INVALIDPARAMETER_CLSINVALIDPARAM = "InvalidParameter.ClsInvalidParam"

	// 日志集冲突。
	INVALIDPARAMETER_CLSLOGSETCONFLICT = "InvalidParameter.ClsLogsetConflict"

	// 日志集为空。
	INVALIDPARAMETER_CLSLOGSETEMPTY = "InvalidParameter.ClsLogsetEmpty"

	// 日志集非空。
	INVALIDPARAMETER_CLSLOGSETNOTEMPTY = "InvalidParameter.ClsLogsetNotEmpty"

	// 没有授权信息。
	INVALIDPARAMETER_CLSMISSINGAUTHORIZATION = "InvalidParameter.ClsMissingAuthorization"

	// 丢失内容。
	INVALIDPARAMETER_CLSMISSINGCONTENT = "InvalidParameter.ClsMissingContent"

	// 语法错误。
	INVALIDPARAMETER_CLSSYNTAXERROR = "InvalidParameter.ClsSyntaxError"

	// 主题已经关闭。
	INVALIDPARAMETER_CLSTOPICCLOSED = "InvalidParameter.ClsTopicClosed"

	// 主题冲突。
	INVALIDPARAMETER_CLSTOPICCONFLICT = "InvalidParameter.ClsTopicConflict"

	// 内部接口错误，请重试或联系客服人员解决。
	INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"

	// 参数错误。
	INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"

	// 正则子模式超出上限。
	INVALIDPARAMETER_PATHREGEXTOOMANYSUBPATTERN = "InvalidParameter.PathRegexTooManySubPattern"

	// 域名所在平台不支持远程鉴权。
	INVALIDPARAMETER_REMOTEAUTHINVALIDPLATFORM = "InvalidParameter.RemoteAuthInvalidPlatform"

	// 域名所在平台不支持使用https协议访问远程鉴权地址。
	INVALIDPARAMETER_REMOTEAUTHINVALIDPROTOCOL = "InvalidParameter.RemoteAuthInvalidProtocol"

	// 任务已过期,无法重试。
	INVALIDPARAMETER_SCDNLOGTASKEXPIRED = "InvalidParameter.ScdnLogTaskExpired"

	// 任务不存在或任务未失败。
	INVALIDPARAMETER_SCDNLOGTASKNOTFOUNDORNOTFAIL = "InvalidParameter.ScdnLogTaskNotFoundOrNotFail"

	// 时间范围错误。
	INVALIDPARAMETER_SCDNLOGTASKTIMERANGEINVALID = "InvalidParameter.ScdnLogTaskTimeRangeInvalid"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 资源数组超过最大值。
	LIMITEXCEEDED_CAMRESOURCEARRAYTOOLONG = "LimitExceeded.CamResourceArrayTooLong"

	// 单个资源标签键数不能超过50。
	LIMITEXCEEDED_CAMRESOURCETOOMANYTAGKEY = "LimitExceeded.CamResourceTooManyTagKey"

	// 标签键长度超过最大值。
	LIMITEXCEEDED_CAMTAGKEYTOOLONG = "LimitExceeded.CamTagKeyTooLong"

	// 单个标签键对应标签值不能超过1000。
	LIMITEXCEEDED_CAMTAGKEYTOOMANYTAGVALUE = "LimitExceeded.CamTagKeyTooManyTagValue"

	// 域名绑定标签数量超出限制。
	LIMITEXCEEDED_CAMTAGQUOTAEXCEEDLIMIT = "LimitExceeded.CamTagQuotaExceedLimit"

	// 单个用户最多1000个不同的key。
	LIMITEXCEEDED_CAMUSERTOOMANYTAGKEY = "LimitExceeded.CamUserTooManyTagKey"

	// 查询IP归属操作过于频繁。
	LIMITEXCEEDED_CDNCALLINGQUERYIPTOOOFTEN = "LimitExceeded.CdnCallingQueryIpTooOften"

	// 该账号已经创建了太多主题。
	LIMITEXCEEDED_CDNCLSTOOMANYTOPICS = "LimitExceeded.CdnClsTooManyTopics"

	// 缓存配置规则数超出限制。
	LIMITEXCEEDED_CDNCONFIGTOOMANYCACHERULES = "LimitExceeded.CdnConfigTooManyCacheRules"

	// 域名操作过于频繁。
	LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"

	// 刷新的目录数量超过限制。
	LIMITEXCEEDED_CDNPURGEPATHEXCEEDBATCHLIMIT = "LimitExceeded.CdnPurgePathExceedBatchLimit"

	// 刷新的目录数量超过每日限制。
	LIMITEXCEEDED_CDNPURGEPATHEXCEEDDAYLIMIT = "LimitExceeded.CdnPurgePathExceedDayLimit"

	// 刷新的Url数量超过限制。
	LIMITEXCEEDED_CDNPURGEURLEXCEEDBATCHLIMIT = "LimitExceeded.CdnPurgeUrlExceedBatchLimit"

	// 刷新的Url数量超过每日限额。
	LIMITEXCEEDED_CDNPURGEURLEXCEEDDAYLIMIT = "LimitExceeded.CdnPurgeUrlExceedDayLimit"

	// 预热的Url数量超过单次限制。
	LIMITEXCEEDED_CDNPUSHEXCEEDBATCHLIMIT = "LimitExceeded.CdnPushExceedBatchLimit"

	// 预热的Url数量超过每日限制。
	LIMITEXCEEDED_CDNPUSHEXCEEDDAYLIMIT = "LimitExceeded.CdnPushExceedDayLimit"

	// 批量查询IP归属个数超过限制。
	LIMITEXCEEDED_CDNQUERYIPBATCHTOOMANY = "LimitExceeded.CdnQueryIpBatchTooMany"

	// 用户域名数量已达上限，请联系腾讯云工程师处理。
	LIMITEXCEEDED_CDNUSERTOOMANYHOSTS = "LimitExceeded.CdnUserTooManyHosts"

	// 日志大小超限。
	LIMITEXCEEDED_CLSLOGSIZEEXCEED = "LimitExceeded.ClsLogSizeExceed"

	// 日志集数目超出。
	LIMITEXCEEDED_CLSLOGSETEXCEED = "LimitExceeded.ClsLogsetExceed"

	// 主题超限。
	LIMITEXCEEDED_CLSTOPICEXCEED = "LimitExceeded.ClsTopicExceed"

	// 每日任务数量超出最大值。
	LIMITEXCEEDED_SCDNLOGTASKEXCEEDDAYLIMIT = "LimitExceeded.ScdnLogTaskExceedDayLimit"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 请求的次数超过了频率限制。
	REQUESTLIMITEXCEEDED = "RequestLimitExceeded"

	// 域名与系统中已存在域名存在冲突。
	RESOURCEINUSE_CDNCONFLICTHOSTEXISTS = "ResourceInUse.CdnConflictHostExists"

	// 域名已存在。
	RESOURCEINUSE_CDNHOSTEXISTS = "ResourceInUse.CdnHostExists"

	// CDN资源正在被操作中。
	RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"

	// 域名已存在。
	RESOURCEINUSE_TCBHOSTEXISTS = "ResourceInUse.TcbHostExists"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 标签键不存在。
	RESOURCENOTFOUND_CAMTAGKEYNOTEXIST = "ResourceNotFound.CamTagKeyNotExist"

	// 未查询到该域名，请确认域名是否正确。
	RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"

	// 项目不存在，请确认后重试。
	RESOURCENOTFOUND_CDNPROJECTNOTEXISTS = "ResourceNotFound.CdnProjectNotExists"

	// 未开通CDN服务，请开通后使用此接口。
	RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"

	// 用户域名数量已达上限，请联系腾讯云工程师处理。
	RESOURCENOTFOUND_CDNUSERTOOMANYHOSTS = "ResourceNotFound.CdnUserTooManyHosts"

	// 索引不存在。
	RESOURCENOTFOUND_CLSINDEXNOTEXIST = "ResourceNotFound.ClsIndexNotExist"

	// 日志集不存在。
	RESOURCENOTFOUND_CLSLOGSETNOTEXIST = "ResourceNotFound.ClsLogsetNotExist"

	// 主题不存在。
	RESOURCENOTFOUND_CLSTOPICNOTEXIST = "ResourceNotFound.ClsTopicNotExist"

	// 域名不存在，请确认后重试。
	RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 该域名已在其他处接入中国境内服务地域，如需修改服务地域为全球，需验证取回域名。
	RESOURCEUNAVAILABLE_CDNHOSTBELONGSTOOTHERSINMAINLAND = "ResourceUnavailable.CdnHostBelongsToOthersInMainland"

	// 该域名已在其他处接入中国境外服务地域，如需修改服务地域为全球，需验证取回域名。
	RESOURCEUNAVAILABLE_CDNHOSTBELONGSTOOTHERSINOVERSEAS = "ResourceUnavailable.CdnHostBelongsToOthersInOverseas"

	// 域名已接入DSA功能。
	RESOURCEUNAVAILABLE_CDNHOSTEXISTSINDSA = "ResourceUnavailable.CdnHostExistsInDsa"

	// 域名已经在TCB控制台接入。
	RESOURCEUNAVAILABLE_CDNHOSTEXISTSINTCB = "ResourceUnavailable.CdnHostExistsInTcb"

	// 域名被锁定。
	RESOURCEUNAVAILABLE_CDNHOSTISLOCKED = "ResourceUnavailable.CdnHostIsLocked"

	// 该域名有违法违规风险，不可接入。
	RESOURCEUNAVAILABLE_CDNHOSTISMALICIOUS = "ResourceUnavailable.CdnHostIsMalicious"

	// 域名未下线。
	RESOURCEUNAVAILABLE_CDNHOSTISNOTOFFLINE = "ResourceUnavailable.CdnHostIsNotOffline"

	// 域名已下线，无法提交预热。
	RESOURCEUNAVAILABLE_CDNHOSTISNOTONLINE = "ResourceUnavailable.CdnHostIsNotOnline"

	// 域名未备案，请将域名备案。备案同步周期为2小时，若域名已备案，可稍后重新接入。
	RESOURCEUNAVAILABLE_CDNHOSTNOICP = "ResourceUnavailable.CdnHostNoIcp"

	// 该域名已在云点播内接入，请先在云点播内删除域名后再接入。
	RESOURCEUNAVAILABLE_HOSTEXISTINVOD = "ResourceUnavailable.HostExistInVod"

	// SCDN服务未生效，请购买或续费SCDN套餐后重试。
	RESOURCEUNAVAILABLE_SCDNUSERNOPACKAGE = "ResourceUnavailable.ScdnUserNoPackage"

	// SCDN服务未生效，请购买或续费SCDN套餐后重试。
	RESOURCEUNAVAILABLE_SCDNUSERSUSPEND = "ResourceUnavailable.ScdnUserSuspend"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 子账号禁止查询整体数据。
	UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"

	// 子账号未配置cam策略。
	UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"

	// 该账号未授权开通CLS。
	UNAUTHORIZEDOPERATION_CDNCLSNOTREGISTERED = "UnauthorizedOperation.CdnClsNotRegistered"

	// 域名解析未进行验证。
	UNAUTHORIZEDOPERATION_CDNDOMAINRECORDNOTVERIFIED = "UnauthorizedOperation.CdnDomainRecordNotVerified"

	// 域名在内部系统已存在，请提工单处理。
	UNAUTHORIZEDOPERATION_CDNHOSTEXISTSININTERNAL = "UnauthorizedOperation.CdnHostExistsInInternal"

	// 该域名涉及违法违规风险，不可接入。
	UNAUTHORIZEDOPERATION_CDNHOSTINICPBLACKLIST = "UnauthorizedOperation.CdnHostInIcpBlacklist"

	// 该域名属于其他账号，您没有权限接入。
	UNAUTHORIZEDOPERATION_CDNHOSTISOWNEDBYOTHER = "UnauthorizedOperation.CdnHostIsOwnedByOther"

	// 域名需要提工单申请接入。
	UNAUTHORIZEDOPERATION_CDNHOSTISTOAPPLYHOST = "UnauthorizedOperation.CdnHostIsToApplyHost"

	// 域名已被其他账号接入，更多详情请提交工单联系我们。
	UNAUTHORIZEDOPERATION_CDNHOSTISUSEDBYOTHER = "UnauthorizedOperation.CdnHostIsUsedByOther"

	// CDN子账号加速域名未授权。
	UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"

	// 用户状态不合法，暂时无法使用服务。
	UNAUTHORIZEDOPERATION_CDNINVALIDUSERSTATUS = "UnauthorizedOperation.CdnInvalidUserStatus"

	// 子账号项目未授权。
	UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"

	// 子账号标签未授权。
	UNAUTHORIZEDOPERATION_CDNTAGUNAUTHORIZED = "UnauthorizedOperation.CdnTagUnauthorized"

	// 域名解析记录值验证不通过。
	UNAUTHORIZEDOPERATION_CDNTXTRECORDVALUENOTMATCH = "UnauthorizedOperation.CdnTxtRecordValueNotMatch"

	// CDN用户认证失败。
	UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"

	// CDN用户待认证。
	UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"

	// 内部服务错误，请联系腾讯云工程师进一步排查。
	UNAUTHORIZEDOPERATION_CDNUSERINVALIDCREDENTIAL = "UnauthorizedOperation.CdnUserInvalidCredential"

	// 账号由于欠费被隔离，请冲正后重试。
	UNAUTHORIZEDOPERATION_CDNUSERISISOLATED = "UnauthorizedOperation.CdnUserIsIsolated"

	// 加速服务已停服，请重启加速服务后重试。
	UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"

	// 非内测白名单用户，无该功能使用权限。
	UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"

	// 无效的授权。
	UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"

	// CLS服务未开通，请先在CLS控制台开通服务。
	UNAUTHORIZEDOPERATION_CLSSERVICENOTACTIVATED = "UnauthorizedOperation.ClsServiceNotActivated"

	// 授权未通过。
	UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"

	// 内部服务错误，请联系腾讯云工程师进一步排查。
	UNAUTHORIZEDOPERATION_CSRFERROR = "UnauthorizedOperation.CsrfError"

	// 鉴权域名为空。
	UNAUTHORIZEDOPERATION_DOMAINEMPTY = "UnauthorizedOperation.DomainEmpty"

	// 请前往CDN控制台进行操作。
	UNAUTHORIZEDOPERATION_ECDNMIGRATEDCDN = "UnauthorizedOperation.EcdnMigratedCdn"

	// 未授权的操作。
	UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"

	// 暂不支持此操作，请联系腾讯云工程师处理。
	UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"

	// 操作超出调用频次限制。
	UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"

	// 未授权操作。
	UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 不允许操作。
	UNSUPPORTEDOPERATION_CLSNOTALLOWED = "UnsupportedOperation.ClsNotAllowed"

	// 暂不支持此操作，请联系腾讯云工程师处理。
	UNSUPPORTEDOPERATION_OPNOAUTH = "UnsupportedOperation.OpNoAuth"
)
