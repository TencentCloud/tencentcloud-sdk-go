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

package v20201111

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// DryRun 操作，代表请求将会是成功的，只是多传了 DryRun 参数。
	DRYRUNOPERATION = "DryRunOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 实名认证失败。
	FAILEDOPERATION_ACCOUNTVERIFYFAIL = "FailedOperation.AccountVerifyFail"

	// 年龄限制无法使用电子签服务，请联系客服咨询处理。
	FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"

	// 签署流程已有关联文档，请检查参数修改后重试。
	FAILEDOPERATION_FLOWHASDOCUMENT = "FailedOperation.FlowHasDocument"

	// 流程未找到关联的电子文件信息，请检查操作步骤，检查参数，并在修改后重试。
	FAILEDOPERATION_FLOWHASNODOCUMENT = "FailedOperation.FlowHasNoDocument"

	// 签署审核未通过，请先完成审核。
	FAILEDOPERATION_NOSIGNREVIEWPASS = "FailedOperation.NoSignReviewPass"

	// 企业经营状态与工商局信息不符。
	FAILEDOPERATION_ORGANIZATIONEXPERIENCECHANGE = "FailedOperation.OrganizationExperienceChange"

	// 企业名称与工商局信息不符。
	FAILEDOPERATION_ORGANIZATIONNAMECHANGED = "FailedOperation.OrganizationNameChanged"

	// 企业名称与工商局信息不符,需要超管修改。
	FAILEDOPERATION_ORGANIZATIONNAMENEEDCHANGE = "FailedOperation.OrganizationNameNeedChange"

	// 创建签署流程预览链接失败，请稍后重试。
	FAILEDOPERATION_PREVIEWURLFAIL = "FailedOperation.PreViewUrlFail"

	// 签署二维码模板发起方签署人存在签署控件，请检查模板后重试。
	FAILEDOPERATION_QRCODECREATORSIGNCOMPONENTS = "FailedOperation.QrCodeCreatorSignComponents"

	// 签署二维码模板签署人不存在，请检查模板后重试。
	FAILEDOPERATION_QRCODESIGNUSERS = "FailedOperation.QrCodeSignUsers"

	// 一码多人二维码模板有误。
	FAILEDOPERATION_QRCODETEMPLATEID = "FailedOperation.QrCodeTemplateId"

	// 请求的次数超过了频率限制，请联系客服处理。
	FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"

	// 模板无资源信息。
	FAILEDOPERATION_TEMPLATEHASNORESOURCE = "FailedOperation.TemplateHasNoResource"

	// 用户信息不匹配，请检查后重试。
	FAILEDOPERATION_USERINFONOMATCH = "FailedOperation.UserInfoNoMatch"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 第三方接口失败。
	INTERNALERROR_API = "InternalError.Api"

	// 缓存错误。
	INTERNALERROR_CACHE = "InternalError.Cache"

	// 调用OpenCloud云api内部错误。
	INTERNALERROR_CALLOPENCLOUDAPIERROR = "InternalError.CallOpenCloudApiError"

	// 回调错误。
	INTERNALERROR_CALLBACK = "InternalError.Callback"

	// 数据库异常。
	INTERNALERROR_DB = "InternalError.Db"

	// 数据库连接出错。
	INTERNALERROR_DBCONNECTION = "InternalError.DbConnection"

	// 数据库新增记录出错。
	INTERNALERROR_DBINSERT = "InternalError.DbInsert"

	// 内部错误,数据库查询失败,请稍后重试。
	INTERNALERROR_DBREAD = "InternalError.DbRead"

	// 解密失败。
	INTERNALERROR_DECRYPTION = "InternalError.Decryption"

	// 依赖的第三方API返回错误。
	INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"

	// 数据库执行错误。
	INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"

	// 加密失败。
	INTERNALERROR_ENCRYPTION = "InternalError.Encryption"

	// Json序列化失败。
	INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"

	// Pdf错误。
	INTERNALERROR_PDF = "InternalError.Pdf"

	// 序列化错误。
	INTERNALERROR_SERIALIZE = "InternalError.Serialize"

	// 系统错误，请稍后重试。
	INTERNALERROR_SYSTEM = "InternalError.System"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 不合法的签署人类型，请检查后重试。
	INVALIDPARAMETER_APPROVERTYPE = "InvalidParameter.ApproverType"

	// 不合法的业务id，请检查是否传递，检查是否超过接口上限数量，并在修改后重试。
	INVALIDPARAMETER_BUSINESSID = "InvalidParameter.BusinessId"

	// 不合法的业务类型，请检查后重试。
	INVALIDPARAMETER_BUSINESSTYPE = "InvalidParameter.BusinessType"

	// 不合法的撤销取消理由，将检查长度，内容，并在修改后重试。
	INVALIDPARAMETER_CANCELREASON = "InvalidParameter.CancelReason"

	// 不合法的证件信息，请检查证件号证件类型是否正确。
	INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"

	// 不合法的证件信息，请检查证件号证件类型是否正确。
	INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"

	// 不合法的抄送方设置，请联系客服了解抄送设置规则，修改后重试。
	INVALIDPARAMETER_CCNUM = "InvalidParameter.CcNum"

	// ClientToken不合法请检查。
	INVALIDPARAMETER_CLIENTTOKEN = "InvalidParameter.ClientToken"

	// 不合法的控件所在文件序号，请检查文件资源与控件的对应关系，并在修改后重试。
	INVALIDPARAMETER_COMPONENTFILEINDEX = "InvalidParameter.ComponentFileIndex"

	// 不合法的控件页码，请与文件资源检查，并在修改后重试。
	INVALIDPARAMETER_COMPONENTPAGE = "InvalidParameter.ComponentPage"

	// 不合法的控件大小或坐标，请检查控件坐标，大小是否合理，并在修改后重试。
	INVALIDPARAMETER_COMPONENTPOSITION = "InvalidParameter.ComponentPosition"

	// 控件的类型与输入值不匹配，请检查参数修改后重试。
	INVALIDPARAMETER_COMPONENTTYPENOMATCHVALUE = "InvalidParameter.ComponentTypeNoMatchValue"

	// 不合法的控件内容，请检查控件是否必填，检查控件内容是否正确设置，并在修改后重试。
	INVALIDPARAMETER_COMPONENTVALUE = "InvalidParameter.ComponentValue"

	// 不合法的模板查询类型，请检查后重试。
	INVALIDPARAMETER_CONTENTTYPE = "InvalidParameter.ContentType"

	// 无效的自定义页卡模板，请检查后重试。
	INVALIDPARAMETER_CUSTOMSHOWMAP = "InvalidParameter.CustomShowMap"

	// 数据已存在。
	INVALIDPARAMETER_DATAEXISTS = "InvalidParameter.DataExists"

	// 数据不存在。
	INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"

	// 参数为空，请检查参数修改后重试。
	INVALIDPARAMETER_EMPTYPARAMS = "InvalidParameter.EmptyParams"

	// 不合法的EndPoint，请检查修改后重试。
	INVALIDPARAMETER_ENDPOINT = "InvalidParameter.EndPoint"

	// 不合法的签署流程回调链接，请修改后重试。
	INVALIDPARAMETER_FLOWCALLBACKURL = "InvalidParameter.FlowCallbackUrl"

	// 不合法的签署流程截止日期，请修改后重试。
	INVALIDPARAMETER_FLOWDEADLINE = "InvalidParameter.FlowDeadLine"

	// 不合法的签署流程描述，请修改后重试。
	INVALIDPARAMETER_FLOWDESCRIPTION = "InvalidParameter.FlowDescription"

	// 不合法的签署流程名称，请修改后重试。
	INVALIDPARAMETER_FLOWNAME = "InvalidParameter.FlowName"

	// 不合法的签署流程类型，请修改后重试。
	INVALIDPARAMETER_FLOWTYPE = "InvalidParameter.FlowType"

	// 不合法的签署流程用户自定义数据，请修改后重试。
	INVALIDPARAMETER_FLOWUSERDATA = "InvalidParameter.FlowUserData"

	// 不合法的FromSource，请联系开发，检查后重试。
	INVALIDPARAMETER_FROMSOURCE = "InvalidParameter.FromSource"

	// 用户个人证件已过期。
	INVALIDPARAMETER_IDCARDVALIDITYOVERLIMIT = "InvalidParameter.IdCardValidityOverLimit"

	// Channel不正确。
	INVALIDPARAMETER_INVALIDCHANNEL = "InvalidParameter.InvalidChannel"

	// id类型参数不合法，请检查后重试。
	INVALIDPARAMETER_INVALIDID = "InvalidParameter.InvalidId"

	// 参数Limit不正确。
	INVALIDPARAMETER_INVALIDLIMIT = "InvalidParameter.InvalidLimit"

	// 手机号码不正确。
	INVALIDPARAMETER_INVALIDMOBILE = "InvalidParameter.InvalidMobile"

	// 姓名不正确。
	INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"

	// 参数Offset不正确。
	INVALIDPARAMETER_INVALIDOFFSET = "InvalidParameter.InvalidOffset"

	// OpenId不正确。
	INVALIDPARAMETER_INVALIDOPENID = "InvalidParameter.InvalidOpenId"

	// 操作人ID不正确。
	INVALIDPARAMETER_INVALIDOPERATORID = "InvalidParameter.InvalidOperatorId"

	// 机构ID不正确。
	INVALIDPARAMETER_INVALIDORGANIZATIONID = "InvalidParameter.InvalidOrganizationId"

	// 组织机构名称不正确。
	INVALIDPARAMETER_INVALIDORGANIZATIONNAME = "InvalidParameter.InvalidOrganizationName"

	// 角色ID不正确。
	INVALIDPARAMETER_INVALIDROLEID = "InvalidParameter.InvalidRoleId"

	// 角色名称不正确。
	INVALIDPARAMETER_INVALIDROLENAME = "InvalidParameter.InvalidRoleName"

	// 实名认证渠道不正确。
	INVALIDPARAMETER_INVALIDVERIFYCHANNEL = "InvalidParameter.InvalidVerifyChannel"

	// 验证码不正确。
	INVALIDPARAMETER_INVALIDVERIFYCODE = "InvalidParameter.InvalidVerifyCode"

	// 缺少必填控件的值。
	INVALIDPARAMETER_MISSINGREQUIREDCOMPONENTVALUE = "InvalidParameter.MissingRequiredComponentValue"

	// 不合法的手机号，请检查后重试。
	INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"

	// 不合法的用户名称，请修改后重试。
	INVALIDPARAMETER_NAME = "InvalidParameter.Name"

	// 不支持的通知类型，请检查并联系客服处理。
	INVALIDPARAMETER_NOTIFYTYPE = "InvalidParameter.NotifyType"

	// 不合法的企业名称，请修改后重试。
	INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"

	// 参数错误。
	INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"

	// 个人静默签Tag未设置，请检查后重试。
	INVALIDPARAMETER_PERSONAUTOSIGNTAG = "InvalidParameter.PersonAutoSignTag"

	// 不合法的阅读时长限制，请联系客服了解阅读时长设置规则，修改后重试。
	INVALIDPARAMETER_PREREADTIME = "InvalidParameter.PreReadTime"

	// 签署二维码的有效期不合法，请联系客服了解规则，并修稿后重试。
	INVALIDPARAMETER_QREFFECTDAY = "InvalidParameter.QrEffectDay"

	// 二维码合同的有效期不合法，请联系客服了解规则，并修稿后重试。
	INVALIDPARAMETER_QRFLOWEFFECTDAY = "InvalidParameter.QrFlowEffectDay"

	// 不合法的资源类型，请联系客服了解，并在修改后重试。
	INVALIDPARAMETER_RESOURCETYPE = "InvalidParameter.ResourceType"

	// 不合法的印章id，请检查印章id是够正确，并在修改后重试。
	INVALIDPARAMETER_SEALID = "InvalidParameter.SealId"

	// 参数包含有敏感词
	INVALIDPARAMETER_SENSITIVE = "InvalidParameter.Sensitive"

	// 签署控件参数不合法，请检查后重试。
	INVALIDPARAMETER_SIGNCOMPONENTS = "InvalidParameter.SignComponents"

	// 参数Status不正确。
	INVALIDPARAMETER_STATUS = "InvalidParameter.Status"

	// 不合法的签署二维码模板id，请检查修改后重试。
	INVALIDPARAMETER_TEMPLATEID = "InvalidParameter.TemplateId"

	// 不合法的认证渠道，请联系客户了解电子签支持的认证渠道，并在修改后重试。
	INVALIDPARAMETER_VERIFYCHANNEL = "InvalidParameter.VerifyChannel"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 需要屏蔽的告警。
	INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 缺少签署人手机号，请检查后重试。
	MISSINGPARAMETER_APPROVERMOBILE = "MissingParameter.ApproverMobile"

	// 缺少签署人姓名，请检查后重试。
	MISSINGPARAMETER_APPROVERNAME = "MissingParameter.ApproverName"

	// 缺少签署人企业信息，请检查后重试。
	MISSINGPARAMETER_APPROVERORGANIZATIONINFO = "MissingParameter.ApproverOrganizationInfo"

	// 缺少签署流程签署人角色信息，请检查修改后重试。
	MISSINGPARAMETER_APPROVERROLE = "MissingParameter.ApproverRole"

	// 缺少签署人签署控件配置，请联系客服了解控件传递规则，并在修改后重试。
	MISSINGPARAMETER_APPROVERSIGNCOMPONENT = "MissingParameter.ApproverSignComponent"

	// 授权码为空，请检查后重试。
	MISSINGPARAMETER_AUTHCODE = "MissingParameter.AuthCode"

	// 缺少撤销取消理由，请检查修改后重试。
	MISSINGPARAMETER_CANCELREASON = "MissingParameter.CancelReason"

	// 未找到表单域配置信息，请检查控件传参，检查文件资源配置，并在修改后重试。
	MISSINGPARAMETER_FIELD = "MissingParameter.Field"

	// 文件名为空，请检查参数修改后重试。
	MISSINGPARAMETER_FILENAMES = "MissingParameter.FileNames"

	// 缺少签署流程签署人信息，请指定签署人信息后重试。
	MISSINGPARAMETER_FLOWAPPROVER = "MissingParameter.FlowApprover"

	// 缺少签署流程id，请检查修改后重试。
	MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"

	// 未找到关键字配置信息，请检查控件传参，检查文件资源配置，并在修改后重试。
	MISSINGPARAMETER_KEYWORD = "MissingParameter.KeyWord"

	// 缺少签署二维码id，请检查后重试。
	MISSINGPARAMETER_QRCODEID = "MissingParameter.QrCodeId"

	// 缺少模板签署人信息，请检查后重试。
	MISSINGPARAMETER_RECIPIENT = "MissingParameter.Recipient"

	// 缺少文件资源ID，请检查后重试。
	MISSINGPARAMETER_RESOURCEID = "MissingParameter.ResourceId"

	// 缺少资源名称，请检查修改后重试。
	MISSINGPARAMETER_RESOURCENAME = "MissingParameter.ResourceName"

	// 缺少静默签印章id，请检查修改后重试。
	MISSINGPARAMETER_SERVERSIGNSEALID = "MissingParameter.ServerSignSealId"

	// 缺少签署控件参数。
	MISSINGPARAMETER_SIGNCOMPONENTS = "MissingParameter.SignComponents"

	// 缺少用户id，请检查后重试。
	MISSINGPARAMETER_USERID = "MissingParameter.UserId"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 签署人设置与模板中签署人配置信息不一致，请检查模板修改参数后重试。
	OPERATIONDENIED_APPROVERNOMATCHTEMPLATE = "OperationDenied.ApproverNoMatchTemplate"

	// 签署人重复，请联系客服了解发起签署流程签署人规则，修改后重试。
	OPERATIONDENIED_APPROVERREPEAT = "OperationDenied.ApproverRepeat"

	// 授权码已失效，请检查是否传递正确，是否已经过期，并在修改后重试。
	OPERATIONDENIED_AUTHCODEINVALID = "OperationDenied.AuthCodeInvalid"

	// 不允许批量撤销签署流程，请检查批量撤销合同信息。
	OPERATIONDENIED_BATCHCANCELFORBID = "OperationDenied.BatchCancelForbid"

	// 子公司不能发起本方母体公司的合同。
	OPERATIONDENIED_BRANCHSENDFLOWTOPARENTNOTALLOW = "OperationDenied.BranchSendFlowToParentNotAllow"

	// 当前不支持抄送，请联系客服咨询处理。
	OPERATIONDENIED_CCFORBID = "OperationDenied.CcForbid"

	// 抄送方存在相同抄送人，请检查修改后重试。
	OPERATIONDENIED_CCUSERREPEAT = "OperationDenied.CcUserRepeat"

	// 电子文档不可用，请稍后重试。
	OPERATIONDENIED_DOCUMENTNOAVAILABLE = "OperationDenied.DocumentNoAvailable"

	// 此企业无该资源使用权限。
	OPERATIONDENIED_ERRNORESOURCEACCESS = "OperationDenied.ErrNoResourceAccess"

	// 个人签署方不支持设置企业名称，请确认签署人类型后重试。
	OPERATIONDENIED_ERRNOSUPPORTINDIVIDUALHASORGANIZATIONNAME = "OperationDenied.ErrNoSupportIndividualHasOrganizationName"

	// 文件已被删除，请联系客服处理。
	OPERATIONDENIED_FILEDELETED = "OperationDenied.FileDeleted"

	// 文件与资源不匹配，请检查文件名与资源id数量以及对应关系，并在修改后重试。
	OPERATIONDENIED_FILENOMATCHRESOURCE = "OperationDenied.FileNoMatchResource"

	// 签署流程无法撤销，请检查签署流程状态，检查签署流程归属企业，检查当前操作人是否有权限，并在修改后重试。
	OPERATIONDENIED_FLOWCANCELFORBID = "OperationDenied.FlowCancelForbid"

	// 签署流程已经被发起，请检查。
	OPERATIONDENIED_FLOWHASSTARTED = "OperationDenied.FlowHasStarted"

	// 签署流程已经被终止，请检查。
	OPERATIONDENIED_FLOWHASTERMINATED = "OperationDenied.FlowHasTerminated"

	// 签署流程不需要进行审核,请修改后重试。
	OPERATIONDENIED_FLOWNONEEDREVIEW = "OperationDenied.FlowNoNeedReview"

	// 签署流程状态不正确，请检查后重试。
	OPERATIONDENIED_FLOWSTATUSFORBID = "OperationDenied.FlowStatusForbid"

	// 禁止此项操作。
	OPERATIONDENIED_FORBID = "OperationDenied.Forbid"

	// 签署人年龄限制无法使用电子签服务。
	OPERATIONDENIED_INVALIDAPPROVERAGE = "OperationDenied.InvalidApproverAge"

	// 资源id超过使用上限限制，请联系客服了解规则，并在修改后重试。
	OPERATIONDENIED_MANYRESOURCEID = "OperationDenied.ManyResourceId"

	// 无权限操作签署流程，请联系客服了解权限，并在修改后重试。
	OPERATIONDENIED_NOFLOWPERMISSION = "OperationDenied.NoFlowPermission"

	// 未通过个人实名认证。
	OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"

	// 用户未登录,请先登录后再操作。
	OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"

	// 未开通静默签功能，请联系签署方企业开通后重试。
	OPERATIONDENIED_NOOPENSERVERSIGN = "OperationDenied.NoOpenServerSign"

	// 无权限使用文件资源，请检查资源有效性以及资源归属，并在修改后重试。
	OPERATIONDENIED_NOPERMISSIONUSERESOURCE = "OperationDenied.NoPermissionUseResource"

	// 无权限使用印章做静默签，请检查印章是否有效，是否有使用权限，并在修改后重试。
	OPERATIONDENIED_NOPERMISSIONUSESERVERSIGNSEAL = "OperationDenied.NoPermissionUseServerSignSeal"

	// 企业额度不足，请检查企业额度后处理。
	OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"

	// 不支持的控件类型，请联系客服了解支持的控件类型，并在修改后重试。
	OPERATIONDENIED_NOSUPPORTCOMPONENTTYPE = "OperationDenied.NoSupportComponentType"

	// 不支持的跳转页，请联系客服了解跳转配置规则，修改后重试。
	OPERATIONDENIED_NOSUPPORTJUMPPAGE = "OperationDenied.NoSupportJumpPage"

	// 未完成实名认证，请检查后重试。
	OPERATIONDENIED_NOVERIFY = "OperationDenied.NoVerify"

	// 不属于企业超管或者法人。
	OPERATIONDENIED_NOTBELONGSUPERADMINORLEGALPERSON = "OperationDenied.NotBelongSuperAdminOrLegalPerson"

	// 操作者权限不足。
	OPERATIONDENIED_OPERATORHASNOPERMISSION = "OperationDenied.OperatorHasNoPermission"

	// 此社会信用编码未查询到结果，请检查后重试。
	OPERATIONDENIED_ORGUNIFORMSOCIALCREDITCODEERR = "OperationDenied.OrgUniformSocialCreditCodeErr"

	// 企业未激活。
	OPERATIONDENIED_ORGANIZATIONNOTACTIVATED = "OperationDenied.OrganizationNotActivated"

	// 查询限频，请先联系客服了解限频策略，稍后重试。
	OPERATIONDENIED_OUTQUERYLIMIT = "OperationDenied.OutQueryLimit"

	// 当前不支持境外用户，请联系客服咨询处理。
	OPERATIONDENIED_OVERSEAFORBID = "OperationDenied.OverSeaForbid"

	// 个人名下没用可使用的签名，请联系个人配置签名后重试。
	OPERATIONDENIED_PERSONHASNOSIGNATURE = "OperationDenied.PersonHasNoSignature"

	// 该用户已关闭或者未开启自动签服务，请检查后重试。
	OPERATIONDENIED_PERSONNOOPENSERVERSIGN = "OperationDenied.PersonNoOpenServerSign"

	// 拒绝个人静默签，请检查个人静默签签署人，并在修改后重试。
	OPERATIONDENIED_PERSONSERVERSIGNFORBID = "OperationDenied.PersonServerSignForbid"

	// 签署二维码已过期，请检查后重试。
	OPERATIONDENIED_QRHASEXPIRE = "OperationDenied.QrHasExpire"

	// 签署二维码不可用，请检查后重试。
	OPERATIONDENIED_QRINVALID = "OperationDenied.QrInvalid"

	// 必填控件未设置填写内容，将检查修改后重试。
	OPERATIONDENIED_REQUIREDCOMPONENTNOTFILL = "OperationDenied.RequiredComponentNotFill"

	// 静默签署方不允许有填写控件，请修改模板，修改参数后重试。
	OPERATIONDENIED_SERVERSIGNNOALLOWCOMPONENT = "OperationDenied.ServerSignNoAllowComponent"

	// 静默签署不支持手写签名，请配置印章并使用印章重试。
	OPERATIONDENIED_SERVERSIGNNOSUPPORTSIGNATURE = "OperationDenied.ServerSignNoSupportSignature"

	// 用户不归属于当前企业，无法操作，请检查后重试。
	OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"

	// 未开通功能白名单，请联系客服处理。
	OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"

	// 请求的次数超过了频率限制。
	REQUESTLIMITEXCEEDED = "RequestLimitExceeded"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 应用号不存在或已删除。
	RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"

	// 机构未完成认证激活，请检查并联系客服处理。
	RESOURCENOTFOUND_AUTHACTIVEORGANIZATION = "ResourceNotFound.AuthActiveOrganization"

	// 未找到电子文档信息，请检查后重试。
	RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"

	// 未获取到用户角色信息，请检查员工角色配置。
	RESOURCENOTFOUND_ERRNOTEXISTROLE = "ResourceNotFound.ErrNotExistRole"

	// 签署流程未找到，请检查参数。
	RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"

	// 签署流程的签署人不存在，请检查后重试。
	RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"

	// 电子文档不存在。
	RESOURCENOTFOUND_NOTEXISTDOCUMENT = "ResourceNotFound.NotExistDocument"

	// 流程不存在。
	RESOURCENOTFOUND_NOTEXISTFLOW = "ResourceNotFound.NotExistFlow"

	// 指定的资源不存在。
	RESOURCENOTFOUND_NOTEXISTRESOURCE = "ResourceNotFound.NotExistResource"

	// 模板不存在。
	RESOURCENOTFOUND_NOTEXISTTEMPLATE = "ResourceNotFound.NotExistTemplate"

	// 机构不存在或者未完成认证，请检查机构信息。
	RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"

	// 签署二维码信息不存在，请检查后重试。
	RESOURCENOTFOUND_QRINFO = "ResourceNotFound.QrInfo"

	// 未获取到文件资源，请检查资源是否完成上传，是否有效，并在修改后重试。
	RESOURCENOTFOUND_RESOURCE = "ResourceNotFound.Resource"

	// 超管信息不存在，请检查企业认证信息。
	RESOURCENOTFOUND_SUPERADMIN = "ResourceNotFound.SuperAdmin"

	// 模板不存在，请检查模板参数，模板配置，并稍后重试。
	RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"

	// Url不存在。
	RESOURCENOTFOUND_URL = "ResourceNotFound.Url"

	// 用户或者员工信息不存在，请检查参数后重试。
	RESOURCENOTFOUND_USER = "ResourceNotFound.User"

	// 用户或者员工未完成实名认证，请检查参数后重试。
	RESOURCENOTFOUND_VERIFYUSER = "ResourceNotFound.VerifyUser"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 请升级到对应版本后即可使用该接口。
	UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"
)
