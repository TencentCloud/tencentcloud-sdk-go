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

package v20181225

const (
	// 此产品的特有错误码

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 不能从Root单元删除成员。
	FAILEDOPERATION_DISABLEDELETEMEMBERFROMROOTNODE = "FailedOperation.DisableDeleteMemberFromRootNode"

	// 不能退出自己创业的企业组织。
	FAILEDOPERATION_DISABLEQUITSELFCREATEDORGANIZATION = "FailedOperation.DisableQuitSelfCreatedOrganization"

	// 已经加入企业组织。
	FAILEDOPERATION_INORGANIZATIONALREADY = "FailedOperation.InOrganizationAlready"

	// 名字已经被使用。
	FAILEDOPERATION_MEMBERNAMEUSED = "FailedOperation.MemberNameUsed"

	// 企业组织单元成员不为空。
	FAILEDOPERATION_NODENOTEMPTY = "FailedOperation.NodeNotEmpty"

	// 只能邀请同一个站点内的账号。
	FAILEDOPERATION_NOTSAMEREGION = "FailedOperation.NotSameRegion"

	// 企业组织已经存在。
	FAILEDOPERATION_ORGANIZATIONEXISTALREADY = "FailedOperation.OrganizationExistAlready"

	// 企业组织成员不为空。
	FAILEDOPERATION_ORGANIZATIONNOTEMPTY = "FailedOperation.OrganizationNotEmpty"

	// 退出共享单元错误。
	FAILEDOPERATION_QUITSHAREUINTERROR = "FailedOperation.QuitShareUintError"

	// 重复发送邀请。
	FAILEDOPERATION_RESENTINVITATION = "FailedOperation.ReSentInvitation"

	// 共享单元不为空。
	FAILEDOPERATION_SHAREUNITNOTEMPTY = "FailedOperation.ShareUnitNotEmpty"

	// 存在不属于当前组织的uin。
	FAILEDOPERATION_SOMEUINSNOTINORGANIZATION = "FailedOperation.SomeUinsNotInOrganization"

	// 用户已经加入企业组织。
	FAILEDOPERATION_USERINORGANIZATION = "FailedOperation.UserInOrganization"

	// 邀请账户不是主账号。
	FAILEDOPERATION_USERNOTREGISTER = "FailedOperation.UserNotRegister"

	// 企业组织成员数量超过上限。
	LIMITEXCEEDED_MEMBERS = "LimitExceeded.Members"

	// 企业组织单元层级太多。
	LIMITEXCEEDED_NODEDEPTHEXCEEDLIMIT = "LimitExceeded.NodeDepthExceedLimit"

	// 组织单元数量超过上限。
	LIMITEXCEEDED_NODEEXCEEDLIMIT = "LimitExceeded.NodeExceedLimit"

	// 组织单元名字已被使用。
	RESOURCEINUSE_NODENAME = "ResourceInUse.NodeName"

	// 名字已被使用。
	RESOURCEINUSE_NODENAMEUSED = "ResourceInUse.NodeNameUsed"

	// 邀请信息不存在。
	RESOURCENOTFOUND_INVITATIONNOTEXIST = "ResourceNotFound.InvitationNotExist"

	// 成员不存在。
	RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"

	// 企业组织单元不存在。
	RESOURCENOTFOUND_NODENOTEXIST = "ResourceNotFound.NodeNotExist"

	// 企业组织不存在。
	RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"

	// 用户不存在。
	RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
)
