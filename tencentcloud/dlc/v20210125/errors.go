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

package v20210125

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 另一个请求正在处理中，请稍后再试。
	FAILEDOPERATION_ANOTHERREQUESTPROCESSING = "FailedOperation.AnotherRequestProcessing"

	// 获取鉴权策略失败。
	FAILEDOPERATION_GETPOLICYFAILED = "FailedOperation.GetPolicyFailed"

	// 获取用户信息失败。
	FAILEDOPERATION_GETUSERINFOFAILED = "FailedOperation.GetUserInfoFailed"

	// 获取工作组信息失败。
	FAILEDOPERATION_GETWORKGROUPINFOFAILED = "FailedOperation.GetWorkGroupInfoFailed"

	// 授权失败。
	FAILEDOPERATION_GRANTPOLICYFAILED = "FailedOperation.GrantPolicyFailed"

	// HTTP客户端请求失败。
	FAILEDOPERATION_HTTPCLIENTDOREQUESTFAILED = "FailedOperation.HttpClientDoRequestFailed"

	// 取消授权失败。
	FAILEDOPERATION_REVOKEPOLICYFAILED = "FailedOperation.RevokePolicyFailed"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 数据库错误。
	INTERNALERROR_DBERROR = "InternalError.DBError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 重复的工作组名称。
	INVALIDPARAMETER_DUPLICATEGROUPNAME = "InvalidParameter.DuplicateGroupName"

	// 重复的用户名。
	INVALIDPARAMETER_DUPLICATEUSERNAME = "InvalidParameter.DuplicateUserName"

	// 无效的访问策略。
	INVALIDPARAMETER_INVALIDACCESSPOLICY = "InvalidParameter.InvalidAccessPolicy"

	// 无效的数据引擎名。
	INVALIDPARAMETER_INVALIDDATAENGINENAME = "InvalidParameter.InvalidDataEngineName"

	// 无效的描述信息。
	INVALIDPARAMETER_INVALIDDESCRIPTION = "InvalidParameter.InvalidDescription"

	// 无效的容错策略。
	INVALIDPARAMETER_INVALIDFAILURETOLERANCE = "InvalidParameter.InvalidFailureTolerance"

	// 无效的过滤条件。
	INVALIDPARAMETER_INVALIDFILTERKEY = "InvalidParameter.InvalidFilterKey"

	// 无效的工作组Id。
	INVALIDPARAMETER_INVALIDGROUPID = "InvalidParameter.InvalidGroupId"

	// 无效的最大结果数。
	INVALIDPARAMETER_INVALIDMAXRESULTS = "InvalidParameter.InvalidMaxResults"

	// 无效的Offset值。
	INVALIDPARAMETER_INVALIDOFFSET = "InvalidParameter.InvalidOffset"

	// 无效的CAM role arn。
	INVALIDPARAMETER_INVALIDROLEARN = "InvalidParameter.InvalidRoleArn"

	// SQL解析失败。
	INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"

	// SQL数量不符合规范。
	INVALIDPARAMETER_INVALIDSQLNUM = "InvalidParameter.InvalidSQLNum"

	// 不支持的排序类型。
	INVALIDPARAMETER_INVALIDSORTBYTYPE = "InvalidParameter.InvalidSortByType"

	// SparkAppParam无效。
	INVALIDPARAMETER_INVALIDSPARKAPPPARAM = "InvalidParameter.InvalidSparkAppParam"

	// 存储位置错误。
	INVALIDPARAMETER_INVALIDSTORELOCATION = "InvalidParameter.InvalidStoreLocation"

	// 无效的taskid。
	INVALIDPARAMETER_INVALIDTASKID = "InvalidParameter.InvalidTaskId"

	// 无效的任务类型。
	INVALIDPARAMETER_INVALIDTASKTYPE = "InvalidParameter.InvalidTaskType"

	// 无效用户名称。
	INVALIDPARAMETER_INVALIDUSERALIAS = "InvalidParameter.InvalidUserAlias"

	// 无效的用户名。
	INVALIDPARAMETER_INVALIDUSERNAME = "InvalidParameter.InvalidUserName"

	// 无效的用户类型。
	INVALIDPARAMETER_INVALIDUSERTYPE = "InvalidParameter.InvalidUserType"

	// 无效的工作组名。
	INVALIDPARAMETER_INVALIDWORKGROUPNAME = "InvalidParameter.InvalidWorkGroupName"

	// 任务已经结束，不能取消。
	INVALIDPARAMETER_TASKALREADYFINISHED = "InvalidParameter.TaskAlreadyFinished"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 指定数据源连接没有找到。
	RESOURCENOTFOUND_DATASOURCENOTFOUND = "ResourceNotFound.DatasourceNotFound"

	// 未找到结果路径。
	RESOURCENOTFOUND_RESULTOUTPUTPATHNOTFOUND = "ResourceNotFound.ResultOutputPathNotFound"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 账号余额不足，无法执行SQL任务。
	RESOURCEUNAVAILABLE_BALANCEINSUFFICIENT = "ResourceUnavailable.BalanceInsufficient"

	// 资源售罄。
	RESOURCESSOLDOUT = "ResourcesSoldOut"

	// 无权限授权策略。
	RESOURCESSOLDOUT_UNAUTHORIZEDGRANTPOLICY = "ResourcesSoldOut.UnauthorizedGrantPolicy"

	// 无权限操作。
	RESOURCESSOLDOUT_UNAUTHORIZEDOPERATION = "ResourcesSoldOut.UnauthorizedOperation"

	// 无权限回收权限。
	RESOURCESSOLDOUT_UNAUTHORIZEDREVOKEPOLICY = "ResourcesSoldOut.UnauthorizedRevokePolicy"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 子用户不是管理员，无权将用户添加到工作组。
	UNAUTHORIZEDOPERATION_ADDUSERSTOWORKGROUP = "UnauthorizedOperation.AddUsersToWorkgroup"

	// 子用户不是管理员，无权绑定工作组到用户。
	UNAUTHORIZEDOPERATION_BINDWORKGROUPSTOUSER = "UnauthorizedOperation.BindWorkgroupsToUser"

	// 子用户不是管理员，无权创建工作组。
	UNAUTHORIZEDOPERATION_CREATEWORKGROUP = "UnauthorizedOperation.CreateWorkgroup"

	// 子用户不是管理员，无权删除用户。
	UNAUTHORIZEDOPERATION_DELETEUSER = "UnauthorizedOperation.DeleteUser"

	// 子用户不是管理员，无权将用户从工作组解绑。
	UNAUTHORIZEDOPERATION_DELETEUSERSFROMWORKGROUP = "UnauthorizedOperation.DeleteUsersFromWorkgroup"

	// 子用户不是管理员，无权删除工作组。
	UNAUTHORIZEDOPERATION_DELETEWORKGROUP = "UnauthorizedOperation.DeleteWorkgroup"

	// 子用户无权授予特定权限。
	UNAUTHORIZEDOPERATION_GRANTPOLICY = "UnauthorizedOperation.GrantPolicy"

	// 子用户不是管理员，无权修改用户信息。
	UNAUTHORIZEDOPERATION_MODIFYUSERINFO = "UnauthorizedOperation.ModifyUserInfo"

	// 子用户不是管理员，无权修改工作组信息。
	UNAUTHORIZEDOPERATION_MODIFYWORKGROUPINFO = "UnauthorizedOperation.ModifyWorkgroupInfo"

	// 子用户无权取消特定权限。
	UNAUTHORIZEDOPERATION_REVOKEPOLICY = "UnauthorizedOperation.RevokePolicy"

	// 子用户不是管理员，无权将工作组和用户解绑。
	UNAUTHORIZEDOPERATION_UNBINDWORKGROUPSFROMUSER = "UnauthorizedOperation.UnbindWorkgroupsFromUser"

	// 子用户无权使用计算引擎。
	UNAUTHORIZEDOPERATION_USECOMPUTINGENGINE = "UnauthorizedOperation.UseComputingEngine"

	// 子用户不存在。
	UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 无法修改主账号。
	UNSUPPORTEDOPERATION_MODIFYOWNERUNSUPPORTED = "UnsupportedOperation.ModifyOwnerUnsupported"
)
