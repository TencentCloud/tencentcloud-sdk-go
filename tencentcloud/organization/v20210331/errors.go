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

	// 用户未实名。
	FAILEDOPERATION_AUTHINFOEMPTY = "FailedOperation.AuthInfoEmpty"

	// 用户非企业实名。
	FAILEDOPERATION_AUTHNOTENTERPRISE = "FailedOperation.AuthNotEnterprise"

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

	// 不能退出自己创建的企业组织。
	FAILEDOPERATION_DISABLEQUITSELFCREATEDORGANIZATION = "FailedOperation.DisableQuitSelfCreatedOrganization"

	// 查询实名信息出错。
	FAILEDOPERATION_GETAUTHINFO = "FailedOperation.GetAuthInfo"

	// 成员存在代付者,不允许删除。
	FAILEDOPERATION_MEMBEREXISTDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberExistDelegatePayerNotAllowDelete"

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

	// 组织权限不合法。
	FAILEDOPERATION_ORGANIZATIONPERMISSIONILLEGAL = "FailedOperation.OrganizationPermissionIllegal"

	// 组织策略不合法。
	FAILEDOPERATION_ORGANIZATIONPOLICYILLEGAL = "FailedOperation.OrganizationPolicyIllegal"

	// 代付者不合法。
	FAILEDOPERATION_PAYUINILLEGAL = "FailedOperation.PayUinIllegal"

	// 退出共享单元错误。
	FAILEDOPERATION_QUITSHAREUINTERROR = "FailedOperation.QuitShareUintError"

	// 成员正在使用共享资源。
	FAILEDOPERATION_SHARERESOURCEMEMBERINUSE = "FailedOperation.ShareResourceMemberInUse"

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

	// 创建成员超过上限。
	LIMITEXCEEDED_CREATEMEMBEROVERLIMIT = "LimitExceeded.CreateMemberOverLimit"

	// 企业组织单元层级太多。
	LIMITEXCEEDED_NODEDEPTHEXCEEDLIMIT = "LimitExceeded.NodeDepthExceedLimit"

	// 组织单元数量超过上限。
	LIMITEXCEEDED_NODEEXCEEDLIMIT = "LimitExceeded.NodeExceedLimit"

	// 成员超过上限。
	LIMITEXCEEDED_ORGANIZATIONMEMBEROVERLIMIT = "LimitExceeded.OrganizationMemberOverLimit"

	// 成员可授权身份不存在。
	RESOURCENOTFOUND_MEMBERIDENTITYNOTEXIST = "ResourceNotFound.MemberIdentityNotExist"

	// 成员不存在。
	RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"

	// 组织成员策略不存在。
	RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"

	// 组织成员不存在。
	RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"

	// 组织节点不在。
	RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"

	// 企业组织不存在。
	RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"

	// 集团服务不存在。
	RESOURCENOTFOUND_ORGANIZATIONSERVICENOTEXIST = "ResourceNotFound.OrganizationServiceNotExist"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 不允许添加代付关系。
	UNSUPPORTEDOPERATION_ADDDELEGATEPAYERNOTALLOW = "UnsupportedOperation.AddDelegatePayerNotAllow"

	// 不允许添加优惠继承关系。
	UNSUPPORTEDOPERATION_ADDDISCOUNTINHERITNOTALLOW = "UnsupportedOperation.AddDiscountInheritNotAllow"

	// 创建的成员不允许删除。
	UNSUPPORTEDOPERATION_CREATEMEMBERNOTALLOWDELETE = "UnsupportedOperation.CreateMemberNotAllowDelete"

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

	// 成员是代理商或代客。
	UNSUPPORTEDOPERATION_MEMBERISAGENT = "UnsupportedOperation.MemberIsAgent"

	// 成员没有绑卡。
	UNSUPPORTEDOPERATION_MEMBERNOPAYMENT = "UnsupportedOperation.MemberNoPayment"

	// 存在在途订单。
	UNSUPPORTEDOPERATION_ORDERINPROGRESSEXISTED = "UnsupportedOperation.OrderInProgressExisted"

	// 管理员存在优惠继承。
	UNSUPPORTEDOPERATION_OWNERDISCOUNTINHERITEXISTED = "UnsupportedOperation.OwnerDiscountInheritExisted"

	// 代付者欠费且未开通信用账户。
	UNSUPPORTEDOPERATION_PAYERARREARSANDNOCREDITACCOUNT = "UnsupportedOperation.PayerArrearsAndNoCreditAccount"

	// 代付者存在账户级优惠。
	UNSUPPORTEDOPERATION_PAYEREXISTACCOUNTLEVELDISCOUNTINHERIT = "UnsupportedOperation.PayerExistAccountLevelDiscountInherit"
)
