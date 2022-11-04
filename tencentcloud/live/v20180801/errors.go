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

package v20180801

const (
	// 此产品的特有错误码

	// DryRun 操作，代表请求将会是成功的，只是多传了 DryRun 参数。
	DRYRUNOPERATION = "DryRunOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 操作 AI 接口失败。
	FAILEDOPERATION_AITRANSCODEOPTIONFAIL = "FailedOperation.AiTranscodeOptionFail"

	// 改变任务状态失败。
	FAILEDOPERATION_ALTERTASKSTATE = "FailedOperation.AlterTaskState"

	// 请检查是否有权限。
	FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"

	// 调用第三方服务失败。
	FAILEDOPERATION_CALLOTHERSVRERROR = "FailedOperation.CallOtherSvrError"

	// 调用内部服务失败。
	FAILEDOPERATION_CALLOTHERSVRFAILED = "FailedOperation.CallOtherSvrFailed"

	// 取消的混流 session 不存在。
	FAILEDOPERATION_CANCELSESSIONNOTEXIST = "FailedOperation.CancelSessionNotExist"

	// 该证书已颁发，不能删除。
	FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"

	// 免费证书申请1小时内不允许删除。
	FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"

	// 证书已存在。
	FAILEDOPERATION_CERTIFICATEEXISTS = "FailedOperation.CertificateExists"

	// 证书不符合标准。
	FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"

	// 证书与私钥不对应。
	FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"

	// 证书不存在。
	FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"

	// 模版使用中。
	FAILEDOPERATION_CONFINUSED = "FailedOperation.ConfInUsed"

	// 下发CDN平台失败。
	FAILEDOPERATION_CONFIGCDNFAILED = "FailedOperation.ConfigCDNFailed"

	// Cos bucket 不存在。
	FAILEDOPERATION_COSBUCKETNOTEXIST = "FailedOperation.CosBucketNotExist"

	// 无权限访问 Cos bucket。
	FAILEDOPERATION_COSBUCKETNOTPERMISSION = "FailedOperation.CosBucketNotPermission"

	// Cos 角色不存在，请前往 控制台 -> 功能配置 -> 直播截图&鉴黄 页面进行授权。
	FAILEDOPERATION_COSROLENOTEXISTS = "FailedOperation.CosRoleNotExists"

	// 2天内有产生流量，域名处于锁定期间，2天内无流量产生才允许删除域名。
	FAILEDOPERATION_DELETEDOMAININLOCKEDTIME = "FailedOperation.DeleteDomainInLockedTime"

	// 域名已添加。
	FAILEDOPERATION_DOMAINADDED = "FailedOperation.DomainAdded"

	// 配置域名规则失败。
	FAILEDOPERATION_DOMAINGSLBFAIL = "FailedOperation.DomainGslbFail"

	// 域名需要实名认证。
	FAILEDOPERATION_DOMAINNEEDREALNAME = "FailedOperation.DomainNeedRealName"

	// 域名归属待验证。
	FAILEDOPERATION_DOMAINNEEDVERIFYOWNER = "FailedOperation.DomainNeedVerifyOwner"

	// 免费证书数量超出限制。
	FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"

	// 无法获取水印 url。
	FAILEDOPERATION_GETPICTUREURLERROR = "FailedOperation.GetPictureUrlError"

	// 获取输入流长宽失败。
	FAILEDOPERATION_GETSTREAMRESOLUTIONERROR = "FailedOperation.GetStreamResolutionError"

	// 无在线流。
	FAILEDOPERATION_HASNOTLIVINGSTREAM = "FailedOperation.HasNotLivingStream"

	// 域名数量超过限制(100个）。
	FAILEDOPERATION_HOSTOUTLIMIT = "FailedOperation.HostOutLimit"

	// 证书状态不正确。
	FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"

	// 参数有误。
	FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"

	// 操作 CDN 接口失败。
	FAILEDOPERATION_INVOKECDNAPIFAIL = "FailedOperation.InvokeCdnApiFail"

	// 操作点播接口异常。
	FAILEDOPERATION_INVOKEVIDEOAPIFAIL = "FailedOperation.InvokeVideoApiFail"

	// 计费平台返回余额不足。
	FAILEDOPERATION_JIFEINOENOUGHFUND = "FailedOperation.JiFeiNoEnoughFund"

	// 当前 CA 机构访问繁忙，请稍后重试。
	FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"

	// 您没有该项目的操作权限。
	FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"

	// 尚未通过实名认证。
	FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"

	// 找不到记录。
	FAILEDOPERATION_NOTFOUND = "FailedOperation.NotFound"

	// 父域名已添加。
	FAILEDOPERATION_PARENTDOMAINADDED = "FailedOperation.ParentDomainAdded"

	// 启动混流失败。
	FAILEDOPERATION_PROCESSMIXERROR = "FailedOperation.ProcessMixError"

	// 查询 upload 信息失败。
	FAILEDOPERATION_QUERYUPLOADINFOFAILED = "FailedOperation.QueryUploadInfoFailed"

	// 规则已经存在。
	FAILEDOPERATION_RULEALREADYEXIST = "FailedOperation.RuleAlreadyExist"

	// 用户没有有效的流量套餐包。
	FAILEDOPERATION_SDKNOPACKAGE = "FailedOperation.SdkNoPackage"

	// 流不存在。
	FAILEDOPERATION_STREAMNOTEXIST = "FailedOperation.StreamNotExist"

	// 子域名已添加。
	FAILEDOPERATION_SUBDOMAINADDED = "FailedOperation.SubDomainAdded"

	// 解绑Tag失败，请尝试手动解绑。
	FAILEDOPERATION_TAGUNBINDERROR = "FailedOperation.TagUnbindError"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 针对添加转码模版的接口。
	INTERNALERROR_ARGSNOTMATCH = "InternalError.ArgsNotMatch"

	// 调用内部服务错误。
	INTERNALERROR_CALLOTHERSVRERROR = "InternalError.CallOtherSvrError"

	// 暂不支持添加中文域名，请核对域名格式。
	INTERNALERROR_CHINESECHARACTERDETECTED = "InternalError.ChineseCharacterDetected"

	// 模版使用中。
	INTERNALERROR_CONFINUSED = "InternalError.ConfInUsed"

	// 模版不存在。
	INTERNALERROR_CONFNOTFOUND = "InternalError.ConfNotFound"

	// 模版数量超过限制。
	INTERNALERROR_CONFOUTLIMIT = "InternalError.ConfOutLimit"

	// 配置不存在。
	INTERNALERROR_CONFIGNOTEXIST = "InternalError.ConfigNotExist"

	// DB 连接错误。
	INTERNALERROR_CONNECTDBERROR = "InternalError.ConnectDbError"

	// 证书使用中。
	INTERNALERROR_CRTDATEINUSING = "InternalError.CrtDateInUsing"

	// 证书不存在。
	INTERNALERROR_CRTDATENOTFOUND = "InternalError.CrtDateNotFound"

	// 证书不合法。
	INTERNALERROR_CRTDATENOTLEGAL = "InternalError.CrtDateNotLegal"

	// 证书过期。
	INTERNALERROR_CRTDATEOVERDUE = "InternalError.CrtDateOverdue"

	// 没有相关域名。
	INTERNALERROR_CRTDOMAINNOTFOUND = "InternalError.CrtDomainNotFound"

	// 证书 Key 不匹配。
	INTERNALERROR_CRTKEYNOTMATCH = "InternalError.CrtKeyNotMatch"

	// DB执行错误。
	INTERNALERROR_DBERROR = "InternalError.DBError"

	// 该域名已在其他处接入，请检查域名是否输入正确。  如域名正确，您可通过验证后重新添加域名。
	INTERNALERROR_DOMAINALREADYEXIST = "InternalError.DomainAlreadyExist"

	// 域名格式错误，请输入合法格式域名。
	INTERNALERROR_DOMAINFORMATERROR = "InternalError.DomainFormatError"

	// 添加 GSLB 规则失败。
	INTERNALERROR_DOMAINGSLBFAIL = "InternalError.DomainGslbFail"

	// 该域名已在其他处接入，请检查域名是否输入正确。  如域名正确，您可通过验证后重新添加域名。
	INTERNALERROR_DOMAINISFAMOUS = "InternalError.DomainIsFamous"

	// 您的域名不可用，请输入正确的域名。
	INTERNALERROR_DOMAINISLIMITED = "InternalError.DomainIsLimited"

	// 域名没有备案。
	INTERNALERROR_DOMAINNORECORD = "InternalError.DomainNoRecord"

	// 域名不存在。
	INTERNALERROR_DOMAINNOTEXIST = "InternalError.DomainNotExist"

	// 域名长度超过限制。
	INTERNALERROR_DOMAINTOOLONG = "InternalError.DomainTooLong"

	// 获取用户账号错误。
	INTERNALERROR_GETBIZIDERROR = "InternalError.GetBizidError"

	// 获取配置错误。
	INTERNALERROR_GETCONFIGERROR = "InternalError.GetConfigError"

	// 获取流信息失败。
	INTERNALERROR_GETSTREAMINFOERROR = "InternalError.GetStreamInfoError"

	// 获取直播源信息错误。
	INTERNALERROR_GETUPSTREAMINFOERROR = "InternalError.GetUpstreamInfoError"

	// 获取水印错误。
	INTERNALERROR_GETWATERMARKERROR = "InternalError.GetWatermarkError"

	// 无在线流。
	INTERNALERROR_HASNOTLIVINGSTREAM = "InternalError.HasNotLivingStream"

	// 参数检校不通过。
	INTERNALERROR_INVALIDINPUT = "InternalError.InvalidInput"

	// 无效的请求。
	INTERNALERROR_INVALIDREQUEST = "InternalError.InvalidRequest"

	// 账号信息错误。
	INTERNALERROR_INVALIDUSER = "InternalError.InvalidUser"

	// 计费平台返回其他错误。
	INTERNALERROR_JIFEIOTHERERROR = "InternalError.JiFeiOtherError"

	// 内部网络错误。
	INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"

	// 记录不存在。
	INTERNALERROR_NOTFOUND = "InternalError.NotFound"

	// 无权限操作。
	INTERNALERROR_NOTPERMMITOPERAT = "InternalError.NotPermmitOperat"

	// 播放域名不存在。
	INTERNALERROR_PLAYDOMAINNORECORD = "InternalError.PlayDomainNoRecord"

	// 转码模板名称已经存在。
	INTERNALERROR_PROCESSORALREADYEXIST = "InternalError.ProcessorAlreadyExist"

	// 推流域名不存在。
	INTERNALERROR_PUSHDOMAINNORECORD = "InternalError.PushDomainNoRecord"

	// 按省份运营商查询播放信息失败。
	INTERNALERROR_QUERYPROISPPLAYINFOERROR = "InternalError.QueryProIspPlayInfoError"

	// 查询 upload 信息失败。
	INTERNALERROR_QUERYUPLOADINFOFAILED = "InternalError.QueryUploadInfoFailed"

	// 规则已经配置。
	INTERNALERROR_RULEALREADYEXIST = "InternalError.RuleAlreadyExist"

	// 规则使用中。
	INTERNALERROR_RULEINUSING = "InternalError.RuleInUsing"

	// 规则不存在。
	INTERNALERROR_RULENOTFOUND = "InternalError.RuleNotFound"

	// 规则超过限制。
	INTERNALERROR_RULEOUTLIMIT = "InternalError.RuleOutLimit"

	// 流状态异常。
	INTERNALERROR_STREAMSTATUSERROR = "InternalError.StreamStatusError"

	// 系统内部错误。
	INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"

	// 更新数据失败。
	INTERNALERROR_UPDATEDATAERROR = "InternalError.UpdateDataError"

	// 添加直播水印失败。
	INTERNALERROR_WATERMARKADDERROR = "InternalError.WatermarkAddError"

	// 水印修改内部错误。
	INTERNALERROR_WATERMARKEDITERROR = "InternalError.WatermarkEditError"

	// 水印不存在。
	INTERNALERROR_WATERMARKNOTEXIST = "InternalError.WatermarkNotExist"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 错误的模板名。
	INVALIDPARAMETER_ARGSNOTMATCH = "InvalidParameter.ArgsNotMatch"

	// Cos自定义文件名错误。
	INVALIDPARAMETER_COSCUSTOMFILENAMEERROR = "InvalidParameter.COSCustomFileNameError"

	// 取消不存在的会话。
	INVALIDPARAMETER_CANCELSESSIONNOTEXIST = "InvalidParameter.CancelSessionNotExist"

	// 腾讯云证书托管 ID 错误。
	INVALIDPARAMETER_CLOUDCRTIDERROR = "InvalidParameter.CloudCrtIdError"

	// 赠送的腾讯云域名已过期。
	INVALIDPARAMETER_CLOUDDOMAINISSTOP = "InvalidParameter.CloudDomainIsStop"

	// 证书使用中。
	INVALIDPARAMETER_CRTDATEINUSING = "InvalidParameter.CrtDateInUsing"

	// 证书不存在。
	INVALIDPARAMETER_CRTDATENOTFOUND = "InvalidParameter.CrtDateNotFound"

	// 证书内容不合法。
	INVALIDPARAMETER_CRTDATENOTLEGAL = "InvalidParameter.CrtDateNotLegal"

	// 证书过期。
	INVALIDPARAMETER_CRTDATEOVERDUE = "InvalidParameter.CrtDateOverdue"

	// 证书没有相关域名。
	INVALIDPARAMETER_CRTDOMAINNOTFOUND = "InvalidParameter.CrtDomainNotFound"

	// 证书 Key 不匹配。
	INVALIDPARAMETER_CRTKEYNOTMATCH = "InvalidParameter.CrtKeyNotMatch"

	// 证书内容或者私钥未提供。
	INVALIDPARAMETER_CRTORKEYNOTEXIST = "InvalidParameter.CrtOrKeyNotExist"

	// 域名已经存在。
	INVALIDPARAMETER_DOMAINALREADYEXIST = "InvalidParameter.DomainAlreadyExist"

	// 域名格式错误，请输入合法格式域名。
	INVALIDPARAMETER_DOMAINFORMATERROR = "InvalidParameter.DomainFormatError"

	// 当前域名在黑名单中。
	INVALIDPARAMETER_DOMAINHITBLACKLIST = "InvalidParameter.DomainHitBlackList"

	// 使用黑名单域名。
	INVALIDPARAMETER_DOMAINISFAMOUS = "InvalidParameter.DomainIsFamous"

	// 域名受限，请提交工单，申请解除限制。
	INVALIDPARAMETER_DOMAINISLIMITED = "InvalidParameter.DomainIsLimited"

	// 域名长度超过限制。
	INVALIDPARAMETER_DOMAINTOOLONG = "InvalidParameter.DomainTooLong"

	// 自适应码率子模板gop值必须存在且相等。
	INVALIDPARAMETER_GOPMUSTEQUALANDEXISTS = "InvalidParameter.GopMustEqualAndExists"

	// 输入数目超出限制。
	INVALIDPARAMETER_INPUTNUMLIMITEXCEEDED = "InvalidParameter.InputNumLimitExceeded"

	// 非法的背景长宽。
	INVALIDPARAMETER_INVALIDBACKGROUDRESOLUTION = "InvalidParameter.InvalidBackgroudResolution"

	// 非法的输出码率。
	INVALIDPARAMETER_INVALIDBITRATE = "InvalidParameter.InvalidBitrate"

	// 回调地址不规范。
	INVALIDPARAMETER_INVALIDCALLBACKURL = "InvalidParameter.InvalidCallbackUrl"

	// 裁剪区域溢出原始图片。
	INVALIDPARAMETER_INVALIDCROPPARAM = "InvalidParameter.InvalidCropParam"

	// 图层参数错误。
	INVALIDPARAMETER_INVALIDLAYERPARAM = "InvalidParameter.InvalidLayerParam"

	// 输出流 ID 被占用。
	INVALIDPARAMETER_INVALIDOUTPUTSTREAMID = "InvalidParameter.InvalidOutputStreamID"

	// 非法输出类型，检查 OutputPram-StreamId 与 OutputType 是否匹配。
	INVALIDPARAMETER_INVALIDOUTPUTTYPE = "InvalidParameter.InvalidOutputType"

	// 水印 ID 未设置。
	INVALIDPARAMETER_INVALIDPICTUREID = "InvalidParameter.InvalidPictureID"

	// 非法的圆角矩形圆角半径。
	INVALIDPARAMETER_INVALIDROUNDRECTRADIUS = "InvalidParameter.InvalidRoundRectRadius"

	// 源地址不规范。
	INVALIDPARAMETER_INVALIDSOURCEURL = "InvalidParameter.InvalidSourceUrl"

	// 任务时间超出限制。
	INVALIDPARAMETER_INVALIDTASKTIME = "InvalidParameter.InvalidTaskTime"

	// 目标地址不规范。
	INVALIDPARAMETER_INVALIDTOURL = "InvalidParameter.InvalidToUrl"

	// 错误的VodFileName。
	INVALIDPARAMETER_INVALIDVODFILENAME = "InvalidParameter.InvalidVodFileName"

	// 水印参数有误。
	INVALIDPARAMETER_INVALIDWATERMARK = "InvalidParameter.InvalidWatermark"

	// 当月不允许添加已删除的小程序域名。
	INVALIDPARAMETER_MPHOSTDELETE = "InvalidParameter.MpHostDelete"

	// 小程序插件没有授权。
	INVALIDPARAMETER_MPPLUGINNOUSE = "InvalidParameter.MpPluginNoUse"

	// 其他错误。
	INVALIDPARAMETER_OTHERERROR = "InvalidParameter.OtherError"

	// 同一会话输出流发生变化。
	INVALIDPARAMETER_SESSIONOUTPUTSTREAMCHANGED = "InvalidParameter.SessionOutputStreamChanged"

	// 任务不存在。
	INVALIDPARAMETER_TASKNOTEXIST = "InvalidParameter.TaskNotExist"

	// 任务数超过限制。
	INVALIDPARAMETER_TASKNUMMORETHANLIMIT = "InvalidParameter.TaskNumMoreThanLimit"

	// 模板与输入流条数不匹配。
	INVALIDPARAMETER_TEMPLATENOTMATCHINPUTNUM = "InvalidParameter.TemplateNotMatchInputNum"

	// 外部地址无权限。
	INVALIDPARAMETER_TOURLNOPERMISSION = "InvalidParameter.ToUrlNoPermission"

	// 域名解析不正确。
	INVALIDPARAMETER_URLNOTSAFE = "InvalidParameter.UrlNotSafe"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 当前并发任务数超限制。
	LIMITEXCEEDED_MAXIMUMRUNNINGTASK = "LimitExceeded.MaximumRunningTask"

	// 当天已创建任务数超限制。
	LIMITEXCEEDED_MAXIMUMTASK = "LimitExceeded.MaximumTask"

	// 单位时间内接口请求频率达到限制。
	LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 日志下载 URL 不存在。
	RESOURCENOTFOUND_CDNLOGEMPTY = "ResourceNotFound.CdnLogEmpty"

	// 日志主题不存在。
	RESOURCENOTFOUND_CDNTHEMEEMPTY = "ResourceNotFound.CdnThemeEmpty"

	// 频道不存在。
	RESOURCENOTFOUND_CHANNELNOTEXIST = "ResourceNotFound.ChannelNotExist"

	// 证书不存在。
	RESOURCENOTFOUND_CRTDATENOTFOUND = "ResourceNotFound.CrtDateNotFound"

	// 未找到证书信息。
	RESOURCENOTFOUND_CRTDOMAINNOTFOUND = "ResourceNotFound.CrtDomainNotFound"

	// 域名没有备案。
	RESOURCENOTFOUND_DOMAINNORECORD = "ResourceNotFound.DomainNoRecord"

	// 域名不存在或不匹配。
	RESOURCENOTFOUND_DOMAINNOTEXIST = "ResourceNotFound.DomainNotExist"

	// 用户被禁用。
	RESOURCENOTFOUND_FORBIDSERVICE = "ResourceNotFound.ForbidService"

	// 用户服务被冻结。
	RESOURCENOTFOUND_FREEZESERVICE = "ResourceNotFound.FreezeService"

	// 用户不支持此接口。
	RESOURCENOTFOUND_INVALIDUSER = "ResourceNotFound.InvalidUser"

	// 播放域名不存在。
	RESOURCENOTFOUND_PLAYDOMAINNORECORD = "ResourceNotFound.PlayDomainNoRecord"

	// 推流域名不存在。
	RESOURCENOTFOUND_PUSHDOMAINNORECORD = "ResourceNotFound.PushDomainNoRecord"

	// 账号停服，请先冲正开通服务后再操作。
	RESOURCENOTFOUND_STOPSERVICE = "ResourceNotFound.StopService"

	// TaskId 不存在。
	RESOURCENOTFOUND_TASKID = "ResourceNotFound.TaskId"

	// 用户主动停服。
	RESOURCENOTFOUND_USERDISABLESERVICE = "ResourceNotFound.UserDisableService"

	// 用户未注册直播。
	RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"

	// 用户不存在。
	RESOURCENOTFOUND_USERNOTFOUNT = "ResourceNotFound.UserNotFount"

	// 水印不存在。
	RESOURCENOTFOUND_WATERMARKNOTEXIST = "ResourceNotFound.WatermarkNotExist"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 点播未开服。
	RESOURCEUNAVAILABLE_INVALIDVODSTATUS = "ResourceUnavailable.InvalidVodStatus"

	// 推流不存在。
	RESOURCEUNAVAILABLE_STREAMNOTEXIST = "ResourceUnavailable.StreamNotExist"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 非直播码/新版控制台模式。
	UNSUPPORTEDOPERATION_NOTLVBCODEMODE = "UnsupportedOperation.NotLVBCodeMode"
)
