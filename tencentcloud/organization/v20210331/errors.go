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

package v20210331

const (
	// 此产品的特有错误码

	// 该帐号已被注册。
	FAILEDOPERATION_ACCOUNTALREADYREGISTER = "FailedOperation.AccountAlreadyRegister"

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

	// 删除授权策略异常。
	FAILEDOPERATION_DELETEPOLICY = "FailedOperation.DeletePolicy"

	// 不能退出自己创建的企业组织。
	FAILEDOPERATION_DISABLEQUITSELFCREATEDORGANIZATION = "FailedOperation.DisableQuitSelfCreatedOrganization"

	// 邮箱已经被使用。
	FAILEDOPERATION_EMAILALREADYUSED = "FailedOperation.EmailAlreadyUsed"

	// 邮箱绑定已经失效。
	FAILEDOPERATION_EMAILBINDRECORDINVALID = "FailedOperation.EmailBindRecordInvalid"

	// 查询实名信息出错。
	FAILEDOPERATION_GETAUTHINFO = "FailedOperation.GetAuthInfo"

	// 查询策略失败。
	FAILEDOPERATION_GETPOLICYDETAIL = "FailedOperation.GetPolicyDetail"

	// 邮箱绑定失败。
	FAILEDOPERATION_MEMBERBINDEMAILERROR = "FailedOperation.MemberBindEmailError"

	// 安全手机绑定失败。
	FAILEDOPERATION_MEMBERBINDPHONEERROR = "FailedOperation.MemberBindPhoneError"

	// 成员存在邮箱。
	FAILEDOPERATION_MEMBEREMAILEXIST = "FailedOperation.MemberEmailExist"

	// 成员存在代付者,不允许删除。
	FAILEDOPERATION_MEMBEREXISTDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberExistDelegatePayerNotAllowDelete"

	// 成员授权在使用。
	FAILEDOPERATION_MEMBERIDENTITYAUTHUSED = "FailedOperation.MemberIdentityAuthUsed"

	// 成员是代付者，不允许删除。
	FAILEDOPERATION_MEMBERISDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberIsDelegatePayerNotAllowDelete"

	// 名字已经被使用。
	FAILEDOPERATION_MEMBERNAMEUSED = "FailedOperation.MemberNameUsed"

	// 成员授权策略名已存在。
	FAILEDOPERATION_MEMBERPOLICYNAMEEXIST = "FailedOperation.MemberPolicyNameExist"

	// 成员正在共享资源。
	FAILEDOPERATION_MEMBERSHARERESOURCE = "FailedOperation.MemberShareResource"

	// 企业组织单元成员不为空。
	FAILEDOPERATION_NODENOTEMPTY = "FailedOperation.NodeNotEmpty"

	// 操作计费侧成员权限错误。
	FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"

	// 操作策略失败。
	FAILEDOPERATION_OPERATEPOLICY = "FailedOperation.OperatePolicy"

	// 成员是主体管理账号不允许删除。
	FAILEDOPERATION_ORGANIZATIONAUTHMANAGENOTALLOWDELETE = "FailedOperation.OrganizationAuthManageNotAllowDelete"

	// 企业组织已经存在。
	FAILEDOPERATION_ORGANIZATIONEXISTALREADY = "FailedOperation.OrganizationExistAlready"

	// 组织身份在使用中。
	FAILEDOPERATION_ORGANIZATIONIDENTITYINUSED = "FailedOperation.OrganizationIdentityInUsed"

	// 组织身份名称被使用。
	FAILEDOPERATION_ORGANIZATIONIDENTITYNAMEUSED = "FailedOperation.OrganizationIdentityNameUsed"

	// 组织身份策略不合法。
	FAILEDOPERATION_ORGANIZATIONIDENTITYPOLICYERROR = "FailedOperation.OrganizationIdentityPolicyError"

	// 成员名已存在。
	FAILEDOPERATION_ORGANIZATIONMEMBERNAMEUSED = "FailedOperation.OrganizationMemberNameUsed"

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

	// 集团策略不是禁用状态。
	FAILEDOPERATION_ORGANIZATIONPOLICYISNOTDISABLED = "FailedOperation.OrganizationPolicyIsNotDisabled"

	// 代付者不合法。
	FAILEDOPERATION_PAYUINILLEGAL = "FailedOperation.PayUinIllegal"

	// 退出共享单元失败。
	FAILEDOPERATION_QUITSHAREUINT = "FailedOperation.QuitShareUint"

	// 退出共享单元错误。
	FAILEDOPERATION_QUITSHAREUINTERROR = "FailedOperation.QuitShareUintError"

	// 退出共享单元失败。
	FAILEDOPERATION_QUITESHAREUNIT = "FailedOperation.QuiteShareUnit"

	// 成员正在使用共享资源。
	FAILEDOPERATION_SHARERESOURCEMEMBERINUSE = "FailedOperation.ShareResourceMemberInUse"

	// 共享单元不为空。
	FAILEDOPERATION_SHAREUNITNOTEMPTY = "FailedOperation.ShareUnitNotEmpty"

	// 存在不属于当前组织的uin。
	FAILEDOPERATION_SOMEUINSNOTINORGANIZATION = "FailedOperation.SomeUinsNotInOrganization"

	// 子账号存在身份。
	FAILEDOPERATION_SUBACCOUNTIDENTITYEXIST = "FailedOperation.SubAccountIdentityExist"

	// 子账号不存在。
	FAILEDOPERATION_SUBACCOUNTNOTEXIST = "FailedOperation.SubAccountNotExist"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 成员代付费模式，不允许主动退出组织。
	INVALIDPARAMETER_ALLOWQUITILLEGAL = "InvalidParameter.AllowQuitIllegal"

	// 手机验证码错误。
	INVALIDPARAMETER_CODEERROR = "InvalidParameter.CodeError"

	// 手机验证码已过期。
	INVALIDPARAMETER_CODEEXPIRED = "InvalidParameter.CodeExpired"

	// 无效的邮箱。
	INVALIDPARAMETER_INVALIDEMAIL = "InvalidParameter.InvalidEmail"

	// 组织成员不存在。
	INVALIDPARAMETER_ORGANIZATIONMEMBERNOTEXIST = "InvalidParameter.OrganizationMemberNotExist"

	// 组织节点不存在。
	INVALIDPARAMETER_ORGANIZATIONNODENOTEXIST = "InvalidParameter.OrganizationNodeNotExist"

	// 密码不合法。
	INVALIDPARAMETER_PASSWORDILLEGAL = "InvalidParameter.PasswordIllegal"

	// 创建成员超过上限。
	LIMITEXCEEDED_CREATEMEMBEROVERLIMIT = "LimitExceeded.CreateMemberOverLimit"

	// 配置邮箱超过当日上限。
	LIMITEXCEEDED_EMAILBINDOVERLIMIT = "LimitExceeded.EmailBindOverLimit"

	// 组织身份超过最大限制。
	LIMITEXCEEDED_IDENTITYEXCEEDLIMIT = "LimitExceeded.IdentityExceedLimit"

	// 企业组织单元层级太多。
	LIMITEXCEEDED_NODEDEPTHEXCEEDLIMIT = "LimitExceeded.NodeDepthExceedLimit"

	// 组织单元数量超过上限。
	LIMITEXCEEDED_NODEEXCEEDLIMIT = "LimitExceeded.NodeExceedLimit"

	// 成员超过上限。
	LIMITEXCEEDED_ORGANIZATIONMEMBEROVERLIMIT = "LimitExceeded.OrganizationMemberOverLimit"

	// 手机超过绑定上限。
	LIMITEXCEEDED_PHONENUMBOUND = "LimitExceeded.PhoneNumBound"

	// 修改成员绑定信息超过限制。
	LIMITEXCEEDED_UPDATEEMAILBINDOVERLIMIT = "LimitExceeded.UpdateEmailBindOverLimit"

	// 邮箱绑定记录不存在。
	RESOURCENOTFOUND_EMAILBINDRECORDNOTEXIST = "ResourceNotFound.EmailBindRecordNotExist"

	// 成员可授权身份不存在。
	RESOURCENOTFOUND_MEMBERIDENTITYNOTEXIST = "ResourceNotFound.MemberIdentityNotExist"

	// 成员不存在。
	RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"

	// 组织成员策略不存在。
	RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"

	// 组织身份不存在。
	RESOURCENOTFOUND_ORGANIZATIONIDENTITYNOTEXIST = "ResourceNotFound.OrganizationIdentityNotExist"

	// 组织成员不存在。
	RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"

	// 组织节点不在。
	RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"

	// 企业组织不存在。
	RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"

	// 集团服务不存在。
	RESOURCENOTFOUND_ORGANIZATIONSERVICENOTEXIST = "ResourceNotFound.OrganizationServiceNotExist"

	// 策略不存在。
	RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"

	// 用户不存在。
	RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 不允许添加代付关系。
	UNSUPPORTEDOPERATION_ADDDELEGATEPAYERNOTALLOW = "UnsupportedOperation.AddDelegatePayerNotAllow"

	// 不允许添加优惠继承关系。
	UNSUPPORTEDOPERATION_ADDDISCOUNTINHERITNOTALLOW = "UnsupportedOperation.AddDiscountInheritNotAllow"

	// 创建的成员不允许创建组织。
	UNSUPPORTEDOPERATION_CREATEMEMBERNOTALLOWCREATEORGANIZATION = "UnsupportedOperation.CreateMemberNotAllowCreateOrganization"

	// 创建的成员不允许删除。
	UNSUPPORTEDOPERATION_CREATEMEMBERNOTALLOWDELETE = "UnsupportedOperation.CreateMemberNotAllowDelete"

	// 创建的成员不允许退出。
	UNSUPPORTEDOPERATION_CREATEMEMBERNOTALLOWQUIT = "UnsupportedOperation.CreateMemberNotAllowQuit"

	// 不允许删除代付关系。
	UNSUPPORTEDOPERATION_DELETEDELEGATEPAYERNOTALLOW = "UnsupportedOperation.DeleteDelegatePayerNotAllow"

	// 成员或者代付者存在经销商。
	UNSUPPORTEDOPERATION_EXISTEDAGENT = "UnsupportedOperation.ExistedAgent"

	// 成员或者代付者存在经销商子客。
	UNSUPPORTEDOPERATION_EXISTEDCLIENT = "UnsupportedOperation.ExistedClient"

	// 用户类型不一致。
	UNSUPPORTEDOPERATION_INCONSISTENTUSERTYPES = "UnsupportedOperation.InconsistentUserTypes"

	// 调用经管系统报错。
	UNSUPPORTEDOPERATION_MANAGEMENTSYSTEMERROR = "UnsupportedOperation.ManagementSystemError"

	// 成员账户欠费。
	UNSUPPORTEDOPERATION_MEMBERACCOUNTARREARS = "UnsupportedOperation.MemberAccountArrears"

	// 成员存在优惠继承。
	UNSUPPORTEDOPERATION_MEMBERDISCOUNTINHERITEXISTED = "UnsupportedOperation.MemberDiscountInheritExisted"

	// 成员存在账户级优惠。
	UNSUPPORTEDOPERATION_MEMBEREXISTACCOUNTLEVELDISCOUNTINHERIT = "UnsupportedOperation.MemberExistAccountLevelDiscountInherit"

	// 成员设置了操作审批,不允许删除。
	UNSUPPORTEDOPERATION_MEMBEREXISTOPERATEPROCESSNOTALLOWDELETE = "UnsupportedOperation.MemberExistOperateProcessNotAllowDelete"

	// 组织成员被委派集团服务，不允许退出。
	UNSUPPORTEDOPERATION_MEMBEREXISTSERVICENOTALLOWDELETE = "UnsupportedOperation.MemberExistServiceNotAllowDelete"

	// 账号是代理商或代客。
	UNSUPPORTEDOPERATION_MEMBERISAGENT = "UnsupportedOperation.MemberIsAgent"

	// 成员没有绑卡。
	UNSUPPORTEDOPERATION_MEMBERNOPAYMENT = "UnsupportedOperation.MemberNoPayment"

	// 成员不支持主动退出。
	UNSUPPORTEDOPERATION_MEMBERNOTALLOWQUIT = "UnsupportedOperation.MemberNotAllowQuit"

	// 存在在途订单。
	UNSUPPORTEDOPERATION_ORDERINPROGRESSEXISTED = "UnsupportedOperation.OrderInProgressExisted"

	// 管理员存在优惠继承。
	UNSUPPORTEDOPERATION_OWNERDISCOUNTINHERITEXISTED = "UnsupportedOperation.OwnerDiscountInheritExisted"

	// 代付者欠费且未开通信用账户。
	UNSUPPORTEDOPERATION_PAYERARREARSANDNOCREDITACCOUNT = "UnsupportedOperation.PayerArrearsAndNoCreditAccount"

	// 代付者存在账户级优惠。
	UNSUPPORTEDOPERATION_PAYEREXISTACCOUNTLEVELDISCOUNTINHERIT = "UnsupportedOperation.PayerExistAccountLevelDiscountInherit"
)
