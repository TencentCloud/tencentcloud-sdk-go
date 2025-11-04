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

package v20210331

const (
	// 此产品的特有错误码

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 该账号已被注册。
	FAILEDOPERATION_ACCOUNTALREADYREGISTER = "FailedOperation.AccountAlreadyRegister"

	// 申请已存在。
	FAILEDOPERATION_APPLYEXIST = "FailedOperation.ApplyExist"

	// 用户未实名。
	FAILEDOPERATION_AUTHINFOEMPTY = "FailedOperation.AuthInfoEmpty"

	// 企业实名不一样。
	FAILEDOPERATION_AUTHINFONOTSAME = "FailedOperation.AuthInfoNotSame"

	// 用户非企业实名。
	FAILEDOPERATION_AUTHNOTENTERPRISE = "FailedOperation.AuthNotEnterprise"

	// 绑定邮箱链接过期。
	FAILEDOPERATION_BINDEMAILLINKEXPIRED = "FailedOperation.BindEmailLinkExpired"

	// 绑定邮箱链接无效。
	FAILEDOPERATION_BINDEMAILLINKINVALID = "FailedOperation.BindEmailLinkInvalid"

	// 成员权限变更记录存在。
	FAILEDOPERATION_CHANGEPERMISSIONRECORDEXIST = "FailedOperation.ChangePermissionRecordExist"

	// 检查手机绑定上限失败。
	FAILEDOPERATION_CHECKACCOUNTPHONEBINDLIMIT = "FailedOperation.CheckAccountPhoneBindLimit"

	// 检查邮箱绑定状态失败。
	FAILEDOPERATION_CHECKMAILACCOUNT = "FailedOperation.CheckMailAccount"

	// 一个目录的权限配置创建超过上限
	FAILEDOPERATION_CONFIGURATIONOVERUPPERLIMIT = "FailedOperation.ConfigurationOverUpperLimit"

	// 创建成员异常。
	FAILEDOPERATION_CREATEACCOUNT = "FailedOperation.CreateAccount"

	// 添加计费权限失败。
	FAILEDOPERATION_CREATEBILLINGPERMISSIONERR = "FailedOperation.CreateBillingPermissionErr"

	// 账号实名认证超过上限。
	FAILEDOPERATION_CREATEMEMBERAUTHOVERLIMIT = "FailedOperation.CreateMemberAuthOverLimit"

	// 创建策略失败。
	FAILEDOPERATION_CREATEPOLICY = "FailedOperation.CreatePolicy"

	// 创建已成功。
	FAILEDOPERATION_CREATERECORDALREADYSUCCESS = "FailedOperation.CreateRecordAlreadySuccess"

	// 创建记录不存在。
	FAILEDOPERATION_CREATERECORDNOTEXIST = "FailedOperation.CreateRecordNotExist"

	// 创建角色异常。
	FAILEDOPERATION_CREATEROLE = "FailedOperation.CreateRole"

	// 权限配置绑定自定义策略超过上限
	FAILEDOPERATION_CUSTOMPOLICYOVERUPPERLIMIT = "FailedOperation.CustomPolicyOverUpperLimit"

	// SQL执行报错
	FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"

	// 解码元数据文档失败。
	FAILEDOPERATION_DECODEMETADATADOCUMENTFAILED = "FailedOperation.DecodeMetadataDocumentFailed"

	// 用户组还存在用户，不允许删除用户
	FAILEDOPERATION_DELETEGROUPNOTALLOWEXISTUSER = "FailedOperation.DeleteGroupNotAllowExistUser"

	// 删除授权策略异常。
	FAILEDOPERATION_DELETEPOLICY = "FailedOperation.DeletePolicy"

	// 查询集团服务使用状态出错。
	FAILEDOPERATION_DESCRIBEORGSERVICEUSAGESTATUSERR = "FailedOperation.DescribeOrgServiceUsageStatusErr"

	// 不能退出自己创建的企业组织。
	FAILEDOPERATION_DISABLEQUITSELFCREATEDORGANIZATION = "FailedOperation.DisableQuitSelfCreatedOrganization"

	// 邮箱已经被使用。
	FAILEDOPERATION_EMAILALREADYUSED = "FailedOperation.EmailAlreadyUsed"

	// 邮箱绑定已经失效。
	FAILEDOPERATION_EMAILBINDRECORDINVALID = "FailedOperation.EmailBindRecordInvalid"

	// 存在共享资源给其他组织成员或被其他组织成员共享资源。
	FAILEDOPERATION_EXISTOTHERORGANIZATIONMEMBERSHARED = "FailedOperation.ExistOtherOrganizationMemberShared"

	// 存在不在组织内的共享成员。
	FAILEDOPERATION_EXISTSHAREMEMBERNOTINORGANIZATION = "FailedOperation.ExistShareMemberNotInOrganization"

	// 获取账号地域属性错误。
	FAILEDOPERATION_GETACCOUNTREGION = "FailedOperation.GetAccountRegion"

	// 查询实名信息出错。
	FAILEDOPERATION_GETAUTHINFO = "FailedOperation.GetAuthInfo"

	// 查询策略失败。
	FAILEDOPERATION_GETPOLICYDETAIL = "FailedOperation.GetPolicyDetail"

	// 用户组创建超过限制
	FAILEDOPERATION_GROUPOVERUPPERLIMIT = "FailedOperation.GroupOverUpperLimit"

	// 用户组类型和用户类型不匹配
	FAILEDOPERATION_GROUPTYPEUSERTYPENOTMATCH = "FailedOperation.GroupTypeUserTypeNotMatch"

	// CIC的用户组组员超出限制。
	FAILEDOPERATION_GROUPUSERCOUNTOVERUPPERLIMIT = "FailedOperation.GroupUserCountOverUpperLimit"

	// 共享单元存在不同的共享资源类型
	FAILEDOPERATION_HASDIFFERENTRESOURCETYPE = "FailedOperation.HasDifferentResourceType"

	// CIC服务已经被开通
	FAILEDOPERATION_IDENTITYCENTERALREADYOPEN = "FailedOperation.IdentityCenterAlreadyOpen"

	// 用户没有实名认证，无法开通CIC服务
	FAILEDOPERATION_IDENTITYCENTERNOTAUTH = "FailedOperation.IdentityCenterNotAuth"

	// 用户没有企业实名认证，无法开通CIC服务
	FAILEDOPERATION_IDENTITYCENTERNOTENTERPRISEAUTH = "FailedOperation.IdentityCenterNotEnterpriseAuth"

	// CIC服务没有开通
	FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"

	// 用户不是集团账号管理员，无法开通CIC服务
	FAILEDOPERATION_IDENTITYCENTERNOTORGANIZATIONMANAGER = "FailedOperation.IdentityCenterNotOrganizationManager"

	// 用户没有开通集团账号，无法开通CIC服务
	FAILEDOPERATION_IDENTITYCENTERORGANIZATIONNOTOPEN = "FailedOperation.IdentityCenterOrganizationNotOpen"

	// 上传文件不合法。
	FAILEDOPERATION_IMPORTFILEILLEGAL = "FailedOperation.ImportFileIllegal"

	// 邀请已存在。
	FAILEDOPERATION_INVITATIONEXIST = "FailedOperation.InvitationExist"

	// 手动用户组不允许删除
	FAILEDOPERATION_MANUALGROUPNOTDELETE = "FailedOperation.ManualGroupNotDelete"

	// 手动用户组不允许更新
	FAILEDOPERATION_MANUALGROUPNOTUPDATE = "FailedOperation.ManualGroupNotUpdate"

	// 手动用户不允许删除
	FAILEDOPERATION_MANUALUSERNOTDELETE = "FailedOperation.ManualUserNotDelete"

	// 手动用户不允许更新
	FAILEDOPERATION_MANUALUSERNOTUPDATE = "FailedOperation.ManualUserNotUpdate"

	// 成员账号删除审核中。
	FAILEDOPERATION_MEMBERACCOUNTDEREGISTERPENDING = "FailedOperation.MemberAccountDeregisterPending"

	// 邮箱绑定失败。
	FAILEDOPERATION_MEMBERBINDEMAILERROR = "FailedOperation.MemberBindEmailError"

	// 安全手机绑定失败。
	FAILEDOPERATION_MEMBERBINDPHONEERROR = "FailedOperation.MemberBindPhoneError"

	// 成员存在邮箱。
	FAILEDOPERATION_MEMBEREMAILEXIST = "FailedOperation.MemberEmailExist"

	// 成员存在代付者,不允许删除。
	FAILEDOPERATION_MEMBEREXISTDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberExistDelegatePayerNotAllowDelete"

	// 成员已存其他组织中。
	FAILEDOPERATION_MEMBEREXISTINOTHERORGANIZATION = "FailedOperation.MemberExistInOtherOrganization"

	// 成员授权在使用。
	FAILEDOPERATION_MEMBERIDENTITYAUTHUSED = "FailedOperation.MemberIdentityAuthUsed"

	// 成员是代付者，不允许删除。
	FAILEDOPERATION_MEMBERISDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberIsDelegatePayerNotAllowDelete"

	// 名字已经被使用。
	FAILEDOPERATION_MEMBERNAMEUSED = "FailedOperation.MemberNameUsed"

	// 成员授权策略名已存在。
	FAILEDOPERATION_MEMBERPOLICYNAMEEXIST = "FailedOperation.MemberPolicyNameExist"

	// 成员正在共享资源，不允许退出组织。
	FAILEDOPERATION_MEMBERSHARERESOURCE = "FailedOperation.MemberShareResource"

	// idp原数据解析失败
	FAILEDOPERATION_METADATACOSPARSINGFAILED = "FailedOperation.MetadataCosParsingFailed"

	// 企业组织单元成员不为空。
	FAILEDOPERATION_NODENOTEMPTY = "FailedOperation.NodeNotEmpty"

	// 只能邀请同一个站点内的账号。
	FAILEDOPERATION_NOTSAMEREGION = "FailedOperation.NotSameRegion"

	// 操作计费侧成员权限错误。
	FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"

	// 操作策略失败。
	FAILEDOPERATION_OPERATEPOLICY = "FailedOperation.OperatePolicy"

	// 存在组织成员访问策略。
	FAILEDOPERATION_ORGMEMBERPOLICYEXIST = "FailedOperation.OrgMemberPolicyExist"

	// 成员是主体管理账号，不允许退出组织。
	FAILEDOPERATION_ORGANIZATIONAUTHMANAGENOTALLOWDELETE = "FailedOperation.OrganizationAuthManageNotAllowDelete"

	// 实名认证关系已经存在。
	FAILEDOPERATION_ORGANIZATIONAUTHRELATIONEXIST = "FailedOperation.OrganizationAuthRelationExist"

	// 解绑最后一个策略失败。
	FAILEDOPERATION_ORGANIZATIONDETACHLASTPOLICYERROR = "FailedOperation.OrganizationDetachLastPolicyError"

	// 解绑策略失败。
	FAILEDOPERATION_ORGANIZATIONDETACHPOLICYERROR = "FailedOperation.OrganizationDetachPolicyError"

	// 企业组织已经存在。
	FAILEDOPERATION_ORGANIZATIONEXISTALREADY = "FailedOperation.OrganizationExistAlready"

	// 组织身份在使用中。
	FAILEDOPERATION_ORGANIZATIONIDENTITYINUSED = "FailedOperation.OrganizationIdentityInUsed"

	// 组织身份名称被使用。
	FAILEDOPERATION_ORGANIZATIONIDENTITYNAMEUSED = "FailedOperation.OrganizationIdentityNameUsed"

	// 组织身份策略不合法。
	FAILEDOPERATION_ORGANIZATIONIDENTITYPOLICYERROR = "FailedOperation.OrganizationIdentityPolicyError"

	// 成员已存在。
	FAILEDOPERATION_ORGANIZATIONMEMBEREXIST = "FailedOperation.OrganizationMemberExist"

	// 成员名已存在。
	FAILEDOPERATION_ORGANIZATIONMEMBERNAMEUSED = "FailedOperation.OrganizationMemberNameUsed"

	// 成员账号不存在。
	FAILEDOPERATION_ORGANIZATIONMEMBERNOTEXIST = "FailedOperation.OrganizationMemberNotExist"

	// 删除节点超过上限。
	FAILEDOPERATION_ORGANIZATIONNODEDELETEOVERLIMIT = "FailedOperation.OrganizationNodeDeleteOverLimit"

	// 节点名已存在。
	FAILEDOPERATION_ORGANIZATIONNODENAMEUSED = "FailedOperation.OrganizationNodeNameUsed"

	// 组织节点不为空。
	FAILEDOPERATION_ORGANIZATIONNODENOTEMPTY = "FailedOperation.OrganizationNodeNotEmpty"

	// 组织节点不存在。
	FAILEDOPERATION_ORGANIZATIONNODENOTEXIST = "FailedOperation.OrganizationNodeNotExist"

	// 企业组织成员不为空。
	FAILEDOPERATION_ORGANIZATIONNOTEMPTY = "FailedOperation.OrganizationNotEmpty"

	// 组织权限不合法。
	FAILEDOPERATION_ORGANIZATIONPERMISSIONILLEGAL = "FailedOperation.OrganizationPermissionIllegal"

	// 组织策略不合法。
	FAILEDOPERATION_ORGANIZATIONPOLICYILLEGAL = "FailedOperation.OrganizationPolicyIllegal"

	// 集团策略正在被使用。
	FAILEDOPERATION_ORGANIZATIONPOLICYINUSED = "FailedOperation.OrganizationPolicyInUsed"

	// 集团策略不是禁用状态。
	FAILEDOPERATION_ORGANIZATIONPOLICYISNOTDISABLED = "FailedOperation.OrganizationPolicyIsNotDisabled"

	// 集团策略不是启用状态。
	FAILEDOPERATION_ORGANIZATIONPOLICYISNOTENABLED = "FailedOperation.OrganizationPolicyIsNotEnabled"

	// 集团服务委派在使用中。
	FAILEDOPERATION_ORGANIZATIONSERVICEASSIGNISUSE = "FailedOperation.OrganizationServiceAssignIsUse"

	// 调用经销系统报错。
	FAILEDOPERATION_PARTNERMANAGEMENTERR = "FailedOperation.PartnerManagementErr"

	// 代付者不合法。
	FAILEDOPERATION_PAYUINILLEGAL = "FailedOperation.PayUinIllegal"

	// 存在成员账号已经开启标签策略，不支持开启集团标签策略
	FAILEDOPERATION_POLICYENABLEINVALID = "FailedOperation.PolicyEnableInvalid"

	// 用户策略数超过上限。
	FAILEDOPERATION_POLICYFULL = "FailedOperation.PolicyFull"

	// PolicyName字段指定的策略名已存在。
	FAILEDOPERATION_POLICYNAMEINUSE = "FailedOperation.PolicyNameInUse"

	// 退出共享单元失败。
	FAILEDOPERATION_QUITSHAREUINT = "FailedOperation.QuitShareUint"

	// 退出共享单元错误。
	FAILEDOPERATION_QUITSHAREUINTERROR = "FailedOperation.QuitShareUintError"

	// 退出共享单元失败。
	FAILEDOPERATION_QUITESHAREUNIT = "FailedOperation.QuiteShareUnit"

	// 重复发送邀请。
	FAILEDOPERATION_RESENTINVITATION = "FailedOperation.ReSentInvitation"

	// 资源超过最大上限。
	FAILEDOPERATION_RESOURCEOVERLIMIT = "FailedOperation.ResourceOverLimit"

	// 权限授权已经存在了。
	FAILEDOPERATION_ROLECONFIGURATIONAUTHORIZATIONALREADYEXIST = "FailedOperation.RoleConfigurationAuthorizationAlreadyExist"

	// 存在权限配置授权。
	FAILEDOPERATION_ROLECONFIGURATIONAUTHORIZATIONEXIST = "FailedOperation.RoleConfigurationAuthorizationExist"

	// 配置权限授权数量超出限制。
	FAILEDOPERATION_ROLECONFIGURATIONAUTHORIZATIONOVERLIMIT = "FailedOperation.RoleConfigurationAuthorizationOverLimit"

	// 访问配置已经部署到成员账号。
	FAILEDOPERATION_ROLECONFIGURATIONPROVISIONINGALREADYDEPLOYED = "FailedOperation.RoleConfigurationProvisioningAlreadyDeployed"

	// 权限配置同步状态错误。
	FAILEDOPERATION_ROLECONFIGURATIONPROVISIONINGSTATUSERROR = "FailedOperation.RoleConfigurationProvisioningStatusError"

	// saml元数据文档获取失败。
	FAILEDOPERATION_SAMLMETADATADOCUMENTGETFAILED = "FailedOperation.SAMLMetadataDocumentGetFailed"

	// SAML服务提供商创建失败
	FAILEDOPERATION_SAMLSERVICEPROVIDERCREATEFAILED = "FailedOperation.SAMLServiceProviderCreateFailed"

	// 当 SSO 登录处于开启状态时不能清空 SAML 身份提供商信息。
	FAILEDOPERATION_SSOSTATUSENABLENOTCLEARIDENTITYPROVIDER = "FailedOperation.SSoStatusEnableNotClearIdentityProvider"

	// 生成SCIM密钥失败
	FAILEDOPERATION_SCIMCREDENTIALGENERATEERROR = "FailedOperation.ScimCredentialGenerateError"

	// 共享地域不存在。
	FAILEDOPERATION_SHAREAREANOTEXIST = "FailedOperation.ShareAreaNotExist"

	// 共享成员不存在。
	FAILEDOPERATION_SHAREMEMBERNOTEXIST = "FailedOperation.ShareMemberNotExist"

	// 成员正在使用共享资源。
	FAILEDOPERATION_SHARERESOURCEMEMBERINUSE = "FailedOperation.ShareResourceMemberInUse"

	// 共享资源不存在。
	FAILEDOPERATION_SHARERESOURCENOTEXIST = "FailedOperation.ShareResourceNotExist"

	// 共享资源类型不存在。
	FAILEDOPERATION_SHARERESOURCETYPENOTEXIST = "FailedOperation.ShareResourceTypeNotExist"

	// 共享单元不为空。
	FAILEDOPERATION_SHAREUNITNOTEMPTY = "FailedOperation.ShareUnitNotEmpty"

	// 共享单元不存在。
	FAILEDOPERATION_SHAREUNITNOTEXIST = "FailedOperation.ShareUnitNotExist"

	// 存在不属于当前组织的uin。
	FAILEDOPERATION_SOMEUINSNOTINORGANIZATION = "FailedOperation.SomeUinsNotInOrganization"

	// 子账号存在身份。
	FAILEDOPERATION_SUBACCOUNTIDENTITYEXIST = "FailedOperation.SubAccountIdentityExist"

	// 子账号不存在。
	FAILEDOPERATION_SUBACCOUNTNOTEXIST = "FailedOperation.SubAccountNotExist"

	// 同步的用户组不允许添加用户
	FAILEDOPERATION_SYNCHRONIZEDGROUPNOTADDUSER = "FailedOperation.SynchronizedGroupNotAddUser"

	// 同步的用户组不允许删除
	FAILEDOPERATION_SYNCHRONIZEDGROUPNOTDELETE = "FailedOperation.SynchronizedGroupNotDelete"

	// 同步的用户组不允许删除用户
	FAILEDOPERATION_SYNCHRONIZEDGROUPNOTREMOVEUSER = "FailedOperation.SynchronizedGroupNotRemoveUser"

	// 同步的用户组不允许更新
	FAILEDOPERATION_SYNCHRONIZEDGROUPNOTUPDATE = "FailedOperation.SynchronizedGroupNotUpdate"

	// 同步的用户不允许删除
	FAILEDOPERATION_SYNCHRONIZEDUSERNOTDELETE = "FailedOperation.SynchronizedUserNotDelete"

	// 同步的用户不允许修改
	FAILEDOPERATION_SYNCHRONIZEDUSERNOTUPDATE = "FailedOperation.SynchronizedUserNotUpdate"

	// 权限配置绑定系统策略超过上限
	FAILEDOPERATION_SYSTEMPOLICYOVERUPPERLIMIT = "FailedOperation.SystemPolicyOverUpperLimit"

	// 打标签异常。
	FAILEDOPERATION_TAGRESOURCESERROR = "FailedOperation.TagResourcesError"

	// 上传数据文件失败。
	FAILEDOPERATION_UPLOADMETADATAFAILED = "FailedOperation.UploadMetadataFailed"

	// 用户加入用户组超出限制。
	FAILEDOPERATION_USERADDGROUPCOUNTOVERUPPERLIMIT = "FailedOperation.UserAddGroupCountOverUpperLimit"

	// 用户创建超过上限
	FAILEDOPERATION_USEROVERUPPERLIMIT = "FailedOperation.UserOverUpperLimit"

	// 用户同步已经存在了。
	FAILEDOPERATION_USERPROVISIONINGALREADYEXISTS = "FailedOperation.UserProvisioningAlreadyExists"

	// 存在用户同步。
	FAILEDOPERATION_USERPROVISIONINGEXISTS = "FailedOperation.UserProvisioningExists"

	// 用户同步失败。
	FAILEDOPERATION_USERPROVISIONINGFAILED = "FailedOperation.UserProvisioningFailed"

	// 用户同步数量超出限制。
	FAILEDOPERATION_USERPROVISIONINGOVERLIMIT = "FailedOperation.UserProvisioningOverLimit"

	// 该x509证书已经存在了。
	FAILEDOPERATION_X509CERTIFICATEALREADYEXIST = "FailedOperation.X509CertificateAlreadyExist"

	// x509证书数量超出限制
	FAILEDOPERATION_X509CERTIFICATELIMITEXCEEDED = "FailedOperation.X509CertificateLimitExceeded"

	// 需要最少一个x509证书
	FAILEDOPERATION_X509CERTIFICATEMINIMUMREQUIRED = "FailedOperation.X509CertificateMinimumRequired"

	// X509证书解析失败。
	FAILEDOPERATION_X509CERTIFICATEPARSINGFAILED = "FailedOperation.X509CertificateParsingFailed"

	// 元数据文档解析XML失败。
	FAILEDOPERATION_XMLDATAUNMARSHALFAILED = "FailedOperation.XMLDataUnmarshalFailed"

	// CIC服务的用户zoneId不存在
	FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 策略文档的Action字段不合法。
	INVALIDPARAMETER_ACTIONERROR = "InvalidParameter.ActionError"

	// 策略语法中操作不存在。
	INVALIDPARAMETER_ACTIONMISS = "InvalidParameter.ActionMiss"

	// 策略语法中操作不存在。
	INVALIDPARAMETER_ACTIONNOTEXIST = "InvalidParameter.ActionNotExist"

	// 策略语法中操作服务不存在。
	INVALIDPARAMETER_ACTIONSERVICENOTEXIST = "InvalidParameter.ActionServiceNotExist"

	// 成员代付费模式，不允许主动退出组织。
	INVALIDPARAMETER_ALLOWQUITILLEGAL = "InvalidParameter.AllowQuitIllegal"

	// 授权对象策略已达上限。
	INVALIDPARAMETER_ATTACHMENTFULL = "InvalidParameter.AttachmentFull"

	// 权限配置不允许绑定该策略。
	INVALIDPARAMETER_BINDPOLICYNAMENOTALLOWED = "InvalidParameter.BindPolicyNameNotAllowed"

	// 手机验证码错误。
	INVALIDPARAMETER_CODEERROR = "InvalidParameter.CodeError"

	// 手机验证码已过期。
	INVALIDPARAMETER_CODEEXPIRED = "InvalidParameter.CodeExpired"

	// 策略语法中条件内容不正确。
	INVALIDPARAMETER_CONDITIONCONTENTERROR = "InvalidParameter.ConditionContentError"

	// 策略文档的condition字段不合法。
	INVALIDPARAMETER_CONDITIONERROR = "InvalidParameter.ConditionError"

	// 策略语法中条件操作符不正确。
	INVALIDPARAMETER_CONDITIONTYPEERROR = "InvalidParameter.ConditionTypeError"

	// 权限配置名称已经存在
	INVALIDPARAMETER_CONFIGURATIONNAMEALREADYEXISTS = "InvalidParameter.ConfigurationNameAlreadyExists"

	// 权限配置名称格式错误
	INVALIDPARAMETER_CONFIGURATIONNAMEFORMATERROR = "InvalidParameter.ConfigurationNameFormatError"

	// Description入参长度不能大于300字节。
	INVALIDPARAMETER_DESCRIPTIONLENGTHOVERLIMIT = "InvalidParameter.DescriptionLengthOverlimit"

	// 策略文档的Effect字段不合法。
	INVALIDPARAMETER_EFFECTERROR = "InvalidParameter.EffectError"

	// 邮箱已经存在
	INVALIDPARAMETER_EMAILALREADYEXISTS = "InvalidParameter.EmailAlreadyExists"

	// 用户组名称已经存在
	INVALIDPARAMETER_GROUPNAMEALREADYEXISTS = "InvalidParameter.GroupNameAlreadyExists"

	// 用户组名称格式错误
	INVALIDPARAMETER_GROUPNAMEFORMATERROR = "InvalidParameter.GroupNameFormatError"

	// 用户组不存在
	INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"

	// 用户组已经存在该用户
	INVALIDPARAMETER_GROUPUSERALREADYEXISTS = "InvalidParameter.GroupUserAlreadyExists"

	// 用户组不存在该用户
	INVALIDPARAMETER_GROUPUSERNOTEXIST = "InvalidParameter.GroupUserNotExist"

	// 接口不存在。
	INVALIDPARAMETER_INTERFACENOTEXIST = "InvalidParameter.InterfaceNotExist"

	// 无效的邮箱。
	INVALIDPARAMETER_INVALIDEMAIL = "InvalidParameter.InvalidEmail"

	// 无效的nextToken
	INVALIDPARAMETER_NEXTTOKENINVALID = "InvalidParameter.NextTokenInvalid"

	// CAM不支持策略文档中所指定的资源类型。
	INVALIDPARAMETER_NOTSUPPORTPRODUCT = "InvalidParameter.NotSupportProduct"

	// 组织成员不存在。
	INVALIDPARAMETER_ORGANIZATIONMEMBERNOTEXIST = "InvalidParameter.OrganizationMemberNotExist"

	// 该成员不是企业管理员。
	INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"

	// 组织节点不存在。
	INVALIDPARAMETER_ORGANIZATIONNODENOTEXIST = "InvalidParameter.OrganizationNodeNotExist"

	// 企业组织不存在。
	INVALIDPARAMETER_ORGANIZATIONNOTEXIST = "InvalidParameter.OrganizationNotExist"

	// 参数错误。
	INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"

	// 密码不合法。
	INVALIDPARAMETER_PASSWORDILLEGAL = "InvalidParameter.PasswordIllegal"

	// 自定义策略内容不能为空
	INVALIDPARAMETER_POLICYDOCUMENTEMPTY = "InvalidParameter.PolicyDocumentEmpty"

	// PolicyDocument字段不合法。
	INVALIDPARAMETER_POLICYDOCUMENTERROR = "InvalidParameter.PolicyDocumentError"

	// PolicyDocument字段超过长度限制。
	INVALIDPARAMETER_POLICYDOCUMENTLENGTHOVERLIMIT = "InvalidParameter.PolicyDocumentLengthOverLimit"

	// 策略ID不存在。
	INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"

	// 标签策略内容的策略键有重复。
	INVALIDPARAMETER_POLICYKEYDUPLICATED = "InvalidParameter.PolicyKeyDuplicated"

	// 策略名称已经存在
	INVALIDPARAMETER_POLICYNAMEALREADYEXISTS = "InvalidParameter.PolicyNameAlreadyExists"

	// PolicyName字段不合法。
	INVALIDPARAMETER_POLICYNAMEERROR = "InvalidParameter.PolicyNameError"

	// 策略名称已经存在。
	INVALIDPARAMETER_POLICYNAMEEXISTED = "InvalidParameter.PolicyNameExisted"

	// 策略名称长度超限。
	INVALIDPARAMETER_POLICYNAMESIZEOVERUPPERLIMIT = "InvalidParameter.PolicyNameSizeOverUpperLimit"

	// 修改策略类型错误，只允许修改自定义策略
	INVALIDPARAMETER_POLICYTYPEERROR = "InvalidParameter.PolicyTypeError"

	// 策略文档的principal字段不合法。
	INVALIDPARAMETER_PRINCIPALERROR = "InvalidParameter.PrincipalError"

	// PrincipalQcs错误。
	INVALIDPARAMETER_PRINCIPALQCSERROR = "InvalidParameter.PrincipalQcsError"

	// PrincipalQcs不存在。
	INVALIDPARAMETER_PRINCIPALQCSNOTEXIST = "InvalidParameter.PrincipalQcsNotExist"

	// PrincipalService不存在。
	INVALIDPARAMETER_PRINCIPALSERVICENOTEXIST = "InvalidParameter.PrincipalServiceNotExist"

	// 系统预留标签键 qcloud、tencent和project 禁止创建。
	INVALIDPARAMETER_RESERVEDTAGKEY = "InvalidParameter.ReservedTagKey"

	// 策略语法中资源内容不正确。
	INVALIDPARAMETER_RESOURCECONTENTERROR = "InvalidParameter.ResourceContentError"

	// 策略文档的Resource字段不合法。
	INVALIDPARAMETER_RESOURCEERROR = "InvalidParameter.ResourceError"

	// 策略语法中资源项目不正确。
	INVALIDPARAMETER_RESOURCEPROJECTERROR = "InvalidParameter.ResourceProjectError"

	// 资源QCS错误。
	INVALIDPARAMETER_RESOURCEQCSERROR = "InvalidParameter.ResourceQcsError"

	// 策略语法中资源地域不正确。
	INVALIDPARAMETER_RESOURCEREGIONERROR = "InvalidParameter.ResourceRegionError"

	// 策略语法中资源服务不存在。
	INVALIDPARAMETER_RESOURCESERVICENOTEXIST = "InvalidParameter.ResourceServiceNotExist"

	// 策略语法中资源所属主账号不正确。
	INVALIDPARAMETER_RESOURCEUINERROR = "InvalidParameter.ResourceUinError"

	// 权限配置不存在
	INVALIDPARAMETER_ROLECONFIGURATIONNOTEXIST = "InvalidParameter.RoleConfigurationNotExist"

	// 权限配置还存在策略
	INVALIDPARAMETER_ROLEPOLICYALREADYEXIST = "InvalidParameter.RolePolicyAlreadyExist"

	// 策略不存在
	INVALIDPARAMETER_ROLEPOLICYNOTEXIST = "InvalidParameter.RolePolicyNotExist"

	// SCIM密钥不存在
	INVALIDPARAMETER_SCIMCREDENTIALNOTFOUND = "InvalidParameter.ScimCredentialNotFound"

	// SCIM同步状态错误
	INVALIDPARAMETER_SCIMSYNCSTATUSERROR = "InvalidParameter.ScimSyncStatusError"

	// 策略文档的Statement字段不合法。
	INVALIDPARAMETER_STATEMENTERROR = "InvalidParameter.StatementError"

	// 标签值错误。
	INVALIDPARAMETER_TAGERROR = "InvalidParameter.TagError"

	// 当前业务不支持标签操作。
	INVALIDPARAMETER_UNSUPPORTEDSERVICE = "InvalidParameter.UnsupportedService"

	// 用户还存在用户组里。
	INVALIDPARAMETER_USERALREADYEXISTSGROUP = "InvalidParameter.UserAlreadyExistsGroup"

	// SCIM密钥状态错误
	INVALIDPARAMETER_USERSCIMCREDENTIALSTATUSERROR = "InvalidParameter.UserScimCredentialStatusError"

	// 用户类型错误。
	INVALIDPARAMETER_USERTYPEERROR = "InvalidParameter.UserTypeError"

	// 用户名已经存在
	INVALIDPARAMETER_USERNAMEALREADYEXISTS = "InvalidParameter.UsernameAlreadyExists"

	// 用户名格式错误
	INVALIDPARAMETER_USERNAMEFORMATERROR = "InvalidParameter.UsernameFormatError"

	// 策略文档的Version字段不合法。
	INVALIDPARAMETER_VERSIONERROR = "InvalidParameter.VersionError"

	// CIC空间名称已经存在，需要更换空间名称。
	INVALIDPARAMETERVALUE_IDENTITYCENTERZONENAMEALREADYEXIST = "InvalidParameterValue.IdentityCenterZoneNameAlreadyExist"

	// 策略内容不合法。
	INVALIDPARAMETERVALUE_POLICYCONTENTINVALID = "InvalidParameterValue.PolicyContentInvalid"

	// sso登录启用状态值非法
	INVALIDPARAMETERVALUE_SSOSTATUSINVALID = "InvalidParameterValue.SSoStatusInvalid"

	// X509证书格式错误。
	INVALIDPARAMETERVALUE_X509CERTIFICATEFORMATERROR = "InvalidParameterValue.X509CertificateFormatError"

	// 空间名不符合规范
	INVALIDPARAMETERVALUE_ZONENAMEFORMATERROR = "InvalidParameterValue.ZoneNameFormatError"

	// 请求加入用户到用户组超出限制。
	LIMITEXCEEDED_ADDUSERTOGROUPLIMITEXCEEDED = "LimitExceeded.AddUserToGroupLimitExceeded"

	// 创建成员超过上限。
	LIMITEXCEEDED_CREATEMEMBEROVERLIMIT = "LimitExceeded.CreateMemberOverLimit"

	// 服务委派管理员超过上限。
	LIMITEXCEEDED_CREATEORGSERVICEASSIGNOVERLIMIT = "LimitExceeded.CreateOrgServiceAssignOverLimit"

	// 请求创建权限配置超出限制。
	LIMITEXCEEDED_CREATEROLEASSIGNMENTLIMITEXCEEDED = "LimitExceeded.CreateRoleAssignmentLimitExceeded"

	// 请求创建用户超出限制。
	LIMITEXCEEDED_CREATEUSERLIMITEXCEEDED = "LimitExceeded.CreateUserLimitExceeded"

	// 请求创建用户同步超出限制。
	LIMITEXCEEDED_CREATEUSERSYNCPROVISIONINGLIMITEXCEEDED = "LimitExceeded.CreateUserSyncProvisioningLimitExceeded"

	// 配置邮箱超过当日上限。
	LIMITEXCEEDED_EMAILBINDOVERLIMIT = "LimitExceeded.EmailBindOverLimit"

	// 组织身份超过最大限制。
	LIMITEXCEEDED_IDENTITYEXCEEDLIMIT = "LimitExceeded.IdentityExceedLimit"

	// 邀请超过上限。
	LIMITEXCEEDED_INVITATIONOVERLIMIT = "LimitExceeded.InvitationOverLimit"

	// 企业组织单元层级太多。
	LIMITEXCEEDED_NODEDEPTHEXCEEDLIMIT = "LimitExceeded.NodeDepthExceedLimit"

	// 组织单元数量超过上限。
	LIMITEXCEEDED_NODEEXCEEDLIMIT = "LimitExceeded.NodeExceedLimit"

	// 成员超过上限。
	LIMITEXCEEDED_ORGANIZATIONMEMBEROVERLIMIT = "LimitExceeded.OrganizationMemberOverLimit"

	// 手机超过绑定上限。
	LIMITEXCEEDED_PHONENUMBOUND = "LimitExceeded.PhoneNumBound"

	// 请求用户组移除用户超出限制。
	LIMITEXCEEDED_REMOVEUSERFROMGROUPLIMITEXCEEDED = "LimitExceeded.RemoveUserFromGroupLimitExceeded"

	// SCIM密钥数量超过限制
	LIMITEXCEEDED_SCIMCREDENTIALLIMITEXCEEDED = "LimitExceeded.ScimCredentialLimitExceeded"

	// 重新发送激活邮件次数超过限制。
	LIMITEXCEEDED_SENDEMAILLIMIT = "LimitExceeded.SendEmailLimit"

	// 一小时内重新发送激活邮件次数超过限制。
	LIMITEXCEEDED_SENDEMAILWITHINONEHOURLIMIT = "LimitExceeded.SendEmailWithinOneHourLimit"

	// 当次操作的共享成员超过上限。
	LIMITEXCEEDED_SHAREUNITMEMBEROVERLIMIT = "LimitExceeded.ShareUnitMemberOverLimit"

	// 当次操作的共享资源超过上限。
	LIMITEXCEEDED_SHAREUNITRESOURCEOVERLIMIT = "LimitExceeded.ShareUnitResourceOverLimit"

	// 标签策略数量超出限制。
	LIMITEXCEEDED_TAGPOLICY = "LimitExceeded.TagPolicy"

	// 修改成员绑定信息超过限制。
	LIMITEXCEEDED_UPDATEEMAILBINDOVERLIMIT = "LimitExceeded.UpdateEmailBindOverLimit"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 申请不存在。
	RESOURCENOTFOUND_APPLYNOTEXIST = "ResourceNotFound.ApplyNotExist"

	// 成员权限变更记录不存在。
	RESOURCENOTFOUND_CHANGEPERMISSIONNOTEXIST = "ResourceNotFound.ChangePermissionNotExist"

	// 有效策略不存在。
	RESOURCENOTFOUND_EFFECTIVEPOLICYNOTFOUND = "ResourceNotFound.EffectivePolicyNotFound"

	// 邮箱绑定记录不存在。
	RESOURCENOTFOUND_EMAILBINDRECORDNOTEXIST = "ResourceNotFound.EmailBindRecordNotExist"

	// 邀请信息不存在。
	RESOURCENOTFOUND_INVITATIONNOTEXIST = "ResourceNotFound.InvitationNotExist"

	// 操作事件不存在。
	RESOURCENOTFOUND_MEMBEREVENTNOTEXIST = "ResourceNotFound.MemberEventNotExist"

	// 成员可授权身份不存在。
	RESOURCENOTFOUND_MEMBERIDENTITYNOTEXIST = "ResourceNotFound.MemberIdentityNotExist"

	// 成员不存在。
	RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"

	// 成员操作审批不存在。
	RESOURCENOTFOUND_MEMBEROPERATEPROCESSNOTEXIST = "ResourceNotFound.MemberOperateProcessNotExist"

	// 组织成员策略不存在。
	RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"

	// 企业组织单元不存在。
	RESOURCENOTFOUND_NODENOTEXIST = "ResourceNotFound.NodeNotExist"

	// 资源不存在。
	RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"

	// 认证主体不存在。
	RESOURCENOTFOUND_ORGANIZATIONAUTHRELATIONNOTEXIST = "ResourceNotFound.OrganizationAuthRelationNotExist"

	// 组织身份不存在。
	RESOURCENOTFOUND_ORGANIZATIONIDENTITYNOTEXIST = "ResourceNotFound.OrganizationIdentityNotExist"

	// 组织成员不存在。
	RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"

	// 组织节点不在。
	RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"

	// 企业组织不存在。
	RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"

	// 集团服务委派不存在。
	RESOURCENOTFOUND_ORGANIZATIONSERVICEASSIGNNOTEXIST = "ResourceNotFound.OrganizationServiceAssignNotExist"

	// 集团服务不存在。
	RESOURCENOTFOUND_ORGANIZATIONSERVICENOTEXIST = "ResourceNotFound.OrganizationServiceNotExist"

	// 策略不存在。
	RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"

	// 策略不存在。
	RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"

	// 可共享资源类型不存在。
	RESOURCENOTFOUND_RESOURCETYPENOTEXIST = "ResourceNotFound.ResourceTypeNotExist"

	// 权限授权记录不存在。
	RESOURCENOTFOUND_ROLECONFIGURATIONAUTHORIZATIONNOTFOUND = "ResourceNotFound.RoleConfigurationAuthorizationNotFound"

	// 权限配置同步不存在。
	RESOURCENOTFOUND_ROLECONFIGURATIONPROVISIONINGNOTFOUND = "ResourceNotFound.RoleConfigurationProvisioningNotFound"

	// 权限配置任务不存在。
	RESOURCENOTFOUND_ROLECONFIGURATIONTASKNOTFOUND = "ResourceNotFound.RoleConfigurationTaskNotFound"

	// saml 身份提供商配置不存在。
	RESOURCENOTFOUND_SAMLIDENTITYPROVIDERNOTFOUND = "ResourceNotFound.SAMLIdentityProviderNotFound"

	// 服务提供商信息不存在。
	RESOURCENOTFOUND_SAMLSERVICEPROVIDERNOTFOUND = "ResourceNotFound.SAMLServiceProviderNotFound"

	// 组织服务角色不存在。
	RESOURCENOTFOUND_SERVICEROLENOTEXIST = "ResourceNotFound.ServiceRoleNotExist"

	// 共享资源成员信息不存在。
	RESOURCENOTFOUND_SHARERESOURCEMEMBERNOTEXIST = "ResourceNotFound.ShareResourceMemberNotExist"

	// 用户不存在。
	RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"

	// 用户同步不存在。
	RESOURCENOTFOUND_USERPROVISIONINGNOTFOUND = "ResourceNotFound.UserProvisioningNotFound"

	// 用户同步任务不存在。
	RESOURCENOTFOUND_USERPROVISIONINGTASKNOTFOUND = "ResourceNotFound.UserProvisioningTaskNotFound"

	// 该x509证书不存在。
	RESOURCENOTFOUND_X509CERTIFICATENOTFOUND = "ResourceNotFound.X509CertificateNotFound"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 代付者财务状态异常，不支持代付费。
	UNSUPPORTEDOPERATION_ABNORMALFINANCIALSTATUSOFADMIN = "UnsupportedOperation.AbnormalFinancialStatusOfAdmin"

	// 成员财务状态异常，不支持代付费。
	UNSUPPORTEDOPERATION_ABNORMALFINANCIALSTATUSOFMEMBER = "UnsupportedOperation.AbnormalFinancialStatusOfMember"

	// 不允许添加代付关系。
	UNSUPPORTEDOPERATION_ADDDELEGATEPAYERNOTALLOW = "UnsupportedOperation.AddDelegatePayerNotAllow"

	// 不允许添加优惠继承关系。
	UNSUPPORTEDOPERATION_ADDDISCOUNTINHERITNOTALLOW = "UnsupportedOperation.AddDiscountInheritNotAllow"

	// 管理员和成员的经销商不一致。
	UNSUPPORTEDOPERATION_AGENTNOTSAME = "UnsupportedOperation.AgentNotSame"

	// 创建的成员不允许创建组织。
	UNSUPPORTEDOPERATION_CREATEMEMBERNOTALLOWCREATEORGANIZATION = "UnsupportedOperation.CreateMemberNotAllowCreateOrganization"

	// 创建的成员不允许移除组织。
	UNSUPPORTEDOPERATION_CREATEMEMBERNOTALLOWDELETE = "UnsupportedOperation.CreateMemberNotAllowDelete"

	// 创建的成员不允许退出。
	UNSUPPORTEDOPERATION_CREATEMEMBERNOTALLOWQUIT = "UnsupportedOperation.CreateMemberNotAllowQuit"

	// 成员删除许可未开启。
	UNSUPPORTEDOPERATION_DELETEACCOUNTDISABLED = "UnsupportedOperation.DeleteAccountDisabled"

	// 不允许删除代付关系。
	UNSUPPORTEDOPERATION_DELETEDELEGATEPAYERNOTALLOW = "UnsupportedOperation.DeleteDelegatePayerNotAllow"

	// 成员或者代付者存在经销商。
	UNSUPPORTEDOPERATION_EXISTEDAGENT = "UnsupportedOperation.ExistedAgent"

	// 成员或者代付者存在经销商子客。
	UNSUPPORTEDOPERATION_EXISTEDCLIENT = "UnsupportedOperation.ExistedClient"

	// 用户类型不一致。
	UNSUPPORTEDOPERATION_INCONSISTENTUSERTYPES = "UnsupportedOperation.InconsistentUserTypes"

	// 邀请的成员不允许删除。
	UNSUPPORTEDOPERATION_INVITEACCOUNTNOTALLOWDELETE = "UnsupportedOperation.InviteAccountNotAllowDelete"

	// 调用经管系统报错。
	UNSUPPORTEDOPERATION_MANAGEMENTSYSTEMERROR = "UnsupportedOperation.ManagementSystemError"

	// 成员账户欠费。
	UNSUPPORTEDOPERATION_MEMBERACCOUNTARREARS = "UnsupportedOperation.MemberAccountArrears"

	// 成员账号存在资源，不允许删除。
	UNSUPPORTEDOPERATION_MEMBERACCOUNTEXISTRESOURCE = "UnsupportedOperation.MemberAccountExistResource"

	// 成员存在优惠继承。
	UNSUPPORTEDOPERATION_MEMBERDISCOUNTINHERITEXISTED = "UnsupportedOperation.MemberDiscountInheritExisted"

	// 成员存在账户级优惠。
	UNSUPPORTEDOPERATION_MEMBEREXISTACCOUNTLEVELDISCOUNTINHERIT = "UnsupportedOperation.MemberExistAccountLevelDiscountInherit"

	// 成员设置了操作审批,不允许删除。
	UNSUPPORTEDOPERATION_MEMBEREXISTOPERATEPROCESSNOTALLOWDELETE = "UnsupportedOperation.MemberExistOperateProcessNotAllowDelete"

	// 成员是集团服务委派管理员，不允许退出组织。
	UNSUPPORTEDOPERATION_MEMBEREXISTSERVICENOTALLOWDELETE = "UnsupportedOperation.MemberExistServiceNotAllowDelete"

	// 成员存在代金劵，不支持代付费。
	UNSUPPORTEDOPERATION_MEMBERHASVOUCHER = "UnsupportedOperation.MemberHasVoucher"

	// 账号是代理商或代客。
	UNSUPPORTEDOPERATION_MEMBERISAGENT = "UnsupportedOperation.MemberIsAgent"

	// 成员不是经销子客。
	UNSUPPORTEDOPERATION_MEMBERISNOTCLIENT = "UnsupportedOperation.MemberIsNotClient"

	// 成员没有绑卡。
	UNSUPPORTEDOPERATION_MEMBERNOPAYMENT = "UnsupportedOperation.MemberNoPayment"

	// 成员不支持主动退出。
	UNSUPPORTEDOPERATION_MEMBERNOTALLOWQUIT = "UnsupportedOperation.MemberNotAllowQuit"

	// 成员不支持操作。
	UNSUPPORTEDOPERATION_MEMBERUNSUPPORTEDOPERATION = "UnsupportedOperation.MemberUnsupportedOperation"

	// 存在在途订单。
	UNSUPPORTEDOPERATION_ORDERINPROGRESSEXISTED = "UnsupportedOperation.OrderInProgressExisted"

	// 管理员存在优惠继承。
	UNSUPPORTEDOPERATION_OWNERDISCOUNTINHERITEXISTED = "UnsupportedOperation.OwnerDiscountInheritExisted"

	// 代付者欠费且未开通信用账户。
	UNSUPPORTEDOPERATION_PAYERARREARSANDNOCREDITACCOUNT = "UnsupportedOperation.PayerArrearsAndNoCreditAccount"

	// 代付者存在账户级优惠。
	UNSUPPORTEDOPERATION_PAYEREXISTACCOUNTLEVELDISCOUNTINHERIT = "UnsupportedOperation.PayerExistAccountLevelDiscountInherit"

	// 存在二级经销商子客，不支持代付费。
	UNSUPPORTEDOPERATION_SECONDARYDISTRIBUTORSUBCLIENTEXISTED = "UnsupportedOperation.SecondaryDistributorSubClientExisted"

	// 不支持共享给其他组织成员。
	UNSUPPORTEDOPERATION_SHARINGTOOTHERORGANIZATIONMEMBER = "UnsupportedOperation.SharingToOtherOrganizationMember"
)
