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

	// DryRun 操作，代表请求将会是成功的，只是多传了 DryRun 参数。
	DRYRUNOPERATION = "DryRunOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 实名认证失败。
	FAILEDOPERATION_ACCOUNTVERIFYFAIL = "FailedOperation.AccountVerifyFail"

	// 16岁以下不提供电子签服务。
	FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"

	// 流程已关联文档。
	FAILEDOPERATION_FLOWHASDOCUMENT = "FailedOperation.FlowHasDocument"

	// 模版无资源信息。
	FAILEDOPERATION_TEMPLATEHASNORESOURCE = "FailedOperation.TemplateHasNoResource"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 缓存错误。
	INTERNALERROR_CACHE = "InternalError.Cache"

	// 数据库异常。
	INTERNALERROR_DB = "InternalError.Db"

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

	// Pdf错误。
	INTERNALERROR_PDF = "InternalError.Pdf"

	// 系统错误。
	INTERNALERROR_SYSTEM = "InternalError.System"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 数据已存在。
	INVALIDPARAMETER_DATAEXISTS = "InvalidParameter.DataExists"

	// 数据不存在。
	INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"

	// 参数为空。
	INVALIDPARAMETER_EMPTYPARAMS = "InvalidParameter.EmptyParams"

	// Channel不正确。
	INVALIDPARAMETER_INVALIDCHANNEL = "InvalidParameter.InvalidChannel"

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

	// 参数错误。
	INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"

	// 参数Status不正确。
	INVALIDPARAMETER_STATUS = "InvalidParameter.Status"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 需要屏蔽的告警。
	INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 禁止此项操作。
	OPERATIONDENIED_FORBID = "OperationDenied.Forbid"

	// 未通过个人实名认证。
	OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"

	// 用户未登录,请先登录后再操作。
	OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"

	// 流程配额不足。
	OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"

	// 企业未激活。
	OPERATIONDENIED_ORGANIZATIONNOTACTIVATED = "OperationDenied.OrganizationNotActivated"

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

	// 流程不存在。
	RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"

	// 电子文档不存在。
	RESOURCENOTFOUND_NOTEXISTDOCUMENT = "ResourceNotFound.NotExistDocument"

	// 流程不存在。
	RESOURCENOTFOUND_NOTEXISTFLOW = "ResourceNotFound.NotExistFlow"

	// 资源不存在。
	RESOURCENOTFOUND_RESOURCE = "ResourceNotFound.Resource"

	// 模版不存在。
	RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"
)
