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

package v20210420

const (
	// 此产品的特有错误码

	// 账号已存在于该账号组。
	FAILEDOPERATION_ACCOUNTALREADYEXISTEDINACCOUNTGROUP = "FailedOperation.AccountAlreadyExistedInAccountGroup"

	// 账号组名称已存在。
	FAILEDOPERATION_ACCOUNTGROUPNAMEALREADYEXISTS = "FailedOperation.AccountGroupNameAlreadyExists"

	// 账号组不存在。
	FAILEDOPERATION_ACCOUNTGROUPNOTEXIST = "FailedOperation.AccountGroupNotExist"

	// 账号名称已存在。
	FAILEDOPERATION_ACCOUNTNAMEALREADYEXISTS = "FailedOperation.AccountNameAlreadyExists"

	// 账号不存在。
	FAILEDOPERATION_ACCOUNTNOTEXIST = "FailedOperation.AccountNotExist"

	// 账号不存在。
	FAILEDOPERATION_ACCOUNTNOTEXISTED = "FailedOperation.AccountNotExisted"

	// 应用名称已存在。
	FAILEDOPERATION_APPDISPLAYNAMEALREADYEXISTS = "FailedOperation.AppDisplayNameAlreadyExists"

	// 应用不存在。
	FAILEDOPERATION_APPNOTEXIST = "FailedOperation.AppNotExist"

	// 同一组织下相同名称的子组织机构已存在。
	FAILEDOPERATION_CHILDORGNODENAMEALREADYEXISTS = "FailedOperation.ChildOrgNodeNameAlreadyExists"

	// 其子机构中存在用户的机构节点不能被删除。
	FAILEDOPERATION_CHILDORGNODEWITHUSERSCANNOTBEDELETED = "FailedOperation.ChildOrgNodeWithUsersCanNotBeDeleted"

	// 创建机构节点失败。
	FAILEDOPERATION_CREATEORGNODEERROR = "FailedOperation.CreateOrgNodeError"

	// 创建用户异常。
	FAILEDOPERATION_CREATEUSERFAILURE = "FailedOperation.CreateUserFailure"

	// 创建用户组失败。
	FAILEDOPERATION_CREATEUSERGROUPERROR = "FailedOperation.CreateUserGroupError"

	// 自定义机构节点对外ID已存在。
	FAILEDOPERATION_CUSTOMIZEPARENTORGNODEIDALREADYEXISTS = "FailedOperation.CustomizeParentOrgNodeIdAlreadyExists"

	// 自定义机构节点对外ID不可填为空字符。
	FAILEDOPERATION_CUSTOMIZEDORGNODEIDCANNOTBEEMPTY = "FailedOperation.CustomizedOrgNodeIdCanNotBeEmpty"

	// 默认导入机构不能被删除。
	FAILEDOPERATION_DEFAULTORGNODECANNOTBEDELETED = "FailedOperation.DefaultOrgNodeCanNotBeDeleted"

	// 删除机构节点失败。
	FAILEDOPERATION_DELETEORGNODEERROR = "FailedOperation.DeleteOrgNodeError"

	// 删除用户失败。
	FAILEDOPERATION_DELETEUSERERROR = "FailedOperation.DeleteUserError"

	// 用户为管理员，不能删除。
	FAILEDOPERATION_DELETEUSEREXISTSADMINISTRATOR = "FailedOperation.DeleteUserExistsAdministrator"

	// 删除用户组失败。
	FAILEDOPERATION_DELETEUSERGROUPERROR = "FailedOperation.DeleteUserGroupError"

	// 读取机构节点信息失败。
	FAILEDOPERATION_DESCRIBEORGNODEERROR = "FailedOperation.DescribeOrgNodeError"

	// 读取根机构节点信息失败。
	FAILEDOPERATION_DESCRIBEORGNODEROOTERROR = "FailedOperation.DescribeOrgNodeRootError"

	// 期望返回的用户属性类型不存在。
	FAILEDOPERATION_EXPECTFIELDSNOTFOUND = "FailedOperation.ExpectFieldsNotFound"

	// 用户组ID未找到。
	FAILEDOPERATION_GROUPIDNOTFOUND = "FailedOperation.GroupIdNotFound"

	// ID转换编码失败。
	FAILEDOPERATION_IDTOCODEERROR = "FailedOperation.IdToCodeError"

	// 同一个应用，相同的用户不能添加超过 %s 个映射关系。
	FAILEDOPERATION_LIMITQUOTANOTENOUGH = "FailedOperation.LimitQuotaNotEnough"

	// 获取用户所在的用户组列表失败。
	FAILEDOPERATION_LISTUSERGROUPSOFUSERERROR = "FailedOperation.ListUserGroupsOfUserError"

	// 读取节点下用户失败。
	FAILEDOPERATION_LISTUSERSINORGNODEERROR = "FailedOperation.ListUsersInOrgNodeError"

	// 获取用户组中的用户列表失败。
	FAILEDOPERATION_LISTUSERSINUSERGROUPERROR = "FailedOperation.ListUsersInUserGroupError"

	// 主组织机构不存在。
	FAILEDOPERATION_MAINORGNODENOTEXIST = "FailedOperation.MainOrgNodeNotExist"

	// 操作失败。
	FAILEDOPERATION_OPERATIONERROR = "FailedOperation.OperationError"

	// 操作失败。
	FAILEDOPERATION_OPERATIONFAILURE = "FailedOperation.OperationFailure"

	// 机构节点ID不存在。
	FAILEDOPERATION_ORGNODEIDNOTEXIST = "FailedOperation.OrgNodeIdNotExist"

	// 机构节点不存在。
	FAILEDOPERATION_ORGNODENOTEXIST = "FailedOperation.OrgNodeNotExist"

	// 机构根节点不能被删除。
	FAILEDOPERATION_ORGNODEROOTCANNOTBEDELETED = "FailedOperation.OrgNodeRootCanNotBeDeleted"

	// 组织机构设置失败，组织机构不可同时被选择为用户所属的主组织和次要机构。
	FAILEDOPERATION_ORGNODESETTINGERROR = "FailedOperation.OrgNodeSettingError"

	// 有用户存在的机构节点不能被删除。
	FAILEDOPERATION_ORGNODEWITHUSERSCANNOTBEDELETED = "FailedOperation.OrgNodeWithUsersCanNotBeDeleted"

	// 父机构节点ID未找到。
	FAILEDOPERATION_PARENTORGNODEIDNOTFOUND = "FailedOperation.ParentOrgNodeIdNotFound"

	// 身份主体未找到。
	FAILEDOPERATION_PERSONNOTFOUND = "FailedOperation.PersonNotFound"

	// 从用户组中移除用户失败。
	FAILEDOPERATION_REMOVEUSERSFROMUSERGROUPERROR = "FailedOperation.RemoveUsersFromUserGroupError"

	// 次要机构元素重复。
	FAILEDOPERATION_SECONDARYORGNODEDUPLICATES = "FailedOperation.SecondaryOrgNodeDuplicates"

	// 不能变更LDAP导入的用户到目录外的组织机构下。
	FAILEDOPERATION_UPDATELDAPUSERORGEXCEEDSRANGE = "FailedOperation.UpdateLDAPUserOrgExceedsRange"

	// 不能变更LDAP导入的用户到其他目录的企业机构下。
	FAILEDOPERATION_UPDATELDAPUSERORGNOTINSAMECROP = "FailedOperation.UpdateLDAPUserOrgNotInSameCrop"

	// 不能变更自建用户到有同步操作类型的组织机构下。
	FAILEDOPERATION_UPDATEUSEREXCEEDSRANGE = "FailedOperation.UpdateUserExceedsRange"

	// 不能变更企业微信用户到企业外的组织机构。
	FAILEDOPERATION_UPDATEWECOMUSERORGEXCEEDSRANGE = "FailedOperation.UpdateWeComUserOrgExceedsRange"

	// 不能变更企业微信用户到非相同的企业机构下。
	FAILEDOPERATION_UPDATEWECOMUSERORGNOTINSAMECROP = "FailedOperation.UpdateWeComUserOrgNotInSameCrop"

	// 获取用户被授权访问的应用列表失败。
	FAILEDOPERATION_USERAUTHLISTERROR = "FailedOperation.UserAuthListError"

	// 用户邮箱已存在。
	FAILEDOPERATION_USEREMAILALREADYEXISTS = "FailedOperation.UserEmailAlreadyExists"

	// 用户组不存在。
	FAILEDOPERATION_USERGROUPNOTEXIST = "FailedOperation.UserGroupNotExist"

	// 用户名字已存在。
	FAILEDOPERATION_USERNAMEALREADYEXISTS = "FailedOperation.UserNameAlreadyExists"

	// 用户不存在于该用户组。
	FAILEDOPERATION_USERNOTEXISTINUSERGROUP = "FailedOperation.UserNotExistInUserGroup"

	// 用户不存在。
	FAILEDOPERATION_USERNOTFOUND = "FailedOperation.UserNotFound"

	// 用户手机已存在。
	FAILEDOPERATION_USERPHONEALREADYEXISTS = "FailedOperation.UserPhoneAlreadyExists"

	// 属性校验失败。
	INVALIDPARAMETER_ATTRIBUTEVALUEVALIDERROR = "InvalidParameter.AttributeValueValidError"

	// 组织机构代码不合法。
	INVALIDPARAMETER_ORGCODEILLEGAL = "InvalidParameter.OrgCodeIllegal"

	// 参数不合法。
	INVALIDPARAMETER_PARAMETERILLEGAL = "InvalidParameter.ParameterIllegal"

	// 参数不合法。
	INVALIDPARAMETER_PARAMETERLLLEGAL = "InvalidParameter.Parameterlllegal"

	// 列表搜索条件不合法。
	INVALIDPARAMETER_SEARCHCRITERIAILLEGAL = "InvalidParameter.SearchCriteriaIllegal"

	// 时间格式不合法。
	INVALIDPARAMETER_TIMEFORMATILLEGAL = "InvalidParameter.TimeFormatIllegal"

	// 设定的用户过期时间不合法。
	INVALIDPARAMETER_USEREXPIRATIONTIMEILLEGAL = "InvalidParameter.UserExpirationTimeIllegal"

	// 用户ID参数为空。
	INVALIDPARAMETER_USERIDISNULL = "InvalidParameter.UserIDIsNull"

	// 用户名参数为空。
	INVALIDPARAMETER_USERNAMEISNULL = "InvalidParameter.UserNameIsNull"

	// 账号ID列表不能为空。
	INVALIDPARAMETERVALUE_ACCOUNTIDSCANNOTBEEMPTY = "InvalidParameterValue.AccountIdsCanNotBeEmpty"

	// 应用ID不能为空。
	INVALIDPARAMETERVALUE_APPIDCANNOTBEEMPTY = "InvalidParameterValue.AppIdCanNotBeEmpty"

	// 应用信息分页展示的排列属性不合法。
	INVALIDPARAMETERVALUE_APPLICATIONINFOSORTKEYILLEGAL = "InvalidParameterValue.ApplicationInfoSortKeyIllegal"

	// EntityType不匹配。
	INVALIDPARAMETERVALUE_ENTITYTYPENOTMATCH = "InvalidParameterValue.EntityTypeNotMatch"

	// 参数的长度超出限制。
	INVALIDPARAMETERVALUE_LENGTHEXCEEDED = "InvalidParameterValue.LengthExceeded"

	// 新密码不能为空。
	INVALIDPARAMETERVALUE_NEWPASSWORDCANNOTBEEMPTY = "InvalidParameterValue.NewPasswordCanNotBeEmpty"

	// 参数不合法。
	INVALIDPARAMETERVALUE_PARAMETERILLEGAL = "InvalidParameterValue.ParameterIllegal"

	// 分页展示的排序属性不合法。
	INVALIDPARAMETERVALUE_SORTKEYILLEGAL = "InvalidParameterValue.SortKeyIllegal"

	// 用户ID不能为空。
	INVALIDPARAMETERVALUE_USERIDCANNOTBEEMPTY = "InvalidParameterValue.UserIdCanNotBeEmpty"

	// 用户名不能为空。
	INVALIDPARAMETERVALUE_USERNAMECANNOTBEEMPTY = "InvalidParameterValue.UserNameCanNotBeEmpty"

	// 用户手机不能为空。
	INVALIDPARAMETERVALUE_USERPHONECANNOTBEEMPTY = "InvalidParameterValue.UserPhoneCanNotBeEmpty"

	// 数据条目数超出最大限制。
	LIMITEXCEEDED_ITEMSNUMBERLIMITEXCEEDED = "LimitExceeded.ItemsNumberLimitExceeded"

	// 输入参数超出长度限制。
	LIMITEXCEEDED_PARAMETERLENGTHLIMITEXCEEDED = "LimitExceeded.ParameterLengthLimitExceeded"

	// 用户的次要机构超出数量上限。
	LIMITEXCEEDED_SECONDARYNODECOUNTLIMITEXCEEDED = "LimitExceeded.SecondaryNodeCountLimitExceeded"

	// 当前用户缺乏访问该操作的权限。
	OPERATIONDENIED_ACTIONPERMISSIONDENY = "OperationDenied.ActionPermissionDeny"
)
