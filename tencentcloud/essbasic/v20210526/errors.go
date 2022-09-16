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

package v20210526

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 签署人未达到合法年龄。
	FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"

	// 鉴权失败。
	FAILEDOPERATION_AUTHFAIL = "FailedOperation.AuthFail"

	// 存在同名印章。
	FAILEDOPERATION_EXISTSAMESEALNAME = "FailedOperation.ExistSameSealName"

	// 合同数目超出。
	FAILEDOPERATION_FLOWNUMEXCEED = "FailedOperation.FlowNumExceed"

	// 已授权。
	FAILEDOPERATION_HASAUTHORIZED = "FailedOperation.HasAuthorized"

	// 发起签署存在填写控件。
	FAILEDOPERATION_QRCODECREATORSIGNCOMPONENTS = "FailedOperation.QrCodeCreatorSignComponents"

	// 模板签署人不存在。
	FAILEDOPERATION_QRCODESIGNUSERS = "FailedOperation.QrCodeSignUsers"

	// 签署二维码模板信息有误，请检查参数后重试。
	FAILEDOPERATION_QRCODETEMPLATEID = "FailedOperation.QrCodeTemplateId"

	// 员工已实名。
	FAILEDOPERATION_STAFFALREADYVERIFY = "FailedOperation.StaffAlreadyVerify"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 其他API错误。
	INTERNALERROR_API = "InternalError.Api"

	// 数据库错误。
	INTERNALERROR_DB = "InternalError.Db"

	// 数据库连接出错。
	INTERNALERROR_DBCONNECTION = "InternalError.DbConnection"

	// 数据库新增记录出错。
	INTERNALERROR_DBINSERT = "InternalError.DbInsert"

	// 数据库读取失败。
	INTERNALERROR_DBREAD = "InternalError.DbRead"

	// 数据库更新记录出错。
	INTERNALERROR_DBUPDATE = "InternalError.DbUpdate"

	// 解密错误。
	INTERNALERROR_DECRYPTION = "InternalError.Decryption"

	// 依赖的其他api出错。
	INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"

	// 加密错误。
	INTERNALERROR_ENCRYPTION = "InternalError.Encryption"

	// 生成唯一ID错误。
	INTERNALERROR_GENERATEID = "InternalError.GenerateId"

	// 上传印章失败。
	INTERNALERROR_SEALUPLOAD = "InternalError.SealUpload"

	// 序列化错误。
	INTERNALERROR_SERIALIZE = "InternalError.Serialize"

	// 系统错误。
	INTERNALERROR_SYSTEM = "InternalError.System"

	// 第三方错误。
	INTERNALERROR_THIRDPARTY = "InternalError.ThirdParty"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 应用号不存在。
	INVALIDPARAMETER_APPLICATION = "InvalidParameter.Application"

	// 参数错误，不合法的签署人类型，请修改后重试。
	INVALIDPARAMETER_APPROVERTYPE = "InvalidParameter.ApproverType"

	// 重复添加签署人。
	INVALIDPARAMETER_BIZAPPROVERALREADYEXISTS = "InvalidParameter.BizApproverAlreadyExists"

	// 营业执照格式不合法。
	INVALIDPARAMETER_BUSINESSLICENSE = "InvalidParameter.BusinessLicense"

	// 证件类型错误。
	INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"

	// 查询内容参数有误。
	INVALIDPARAMETER_CONTENTTYPE = "InvalidParameter.ContentType"

	// 参数错误，无效的自定义页卡模板，仅支持{合同名称}{发起方姓名}{发起方企业}{签署方N姓名}{签署方N企业}，请修改后重试。
	INVALIDPARAMETER_CUSTOMSHOWMAP = "InvalidParameter.CustomShowMap"

	// 参数错误，UserData长度非法，请修改后重试。
	INVALIDPARAMETER_CUSTOMERDATA = "InvalidParameter.CustomerData"

	// 数据不存在。
	INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"

	// 参数错误，不合法的日期，请检查后重试。
	INVALIDPARAMETER_DATE = "InvalidParameter.Date"

	// 参数为空。
	INVALIDPARAMETER_EMPTYPARAMS = "InvalidParameter.EmptyParams"

	// 不合法的EndPoint，请检查修改后重试。
	INVALIDPARAMETER_ENDPOINT = "InvalidParameter.EndPoint"

	// 文件类型不合法。
	INVALIDPARAMETER_FILETYPE = "InvalidParameter.FileType"

	// 参数错误，不合法的备选签署人数量，请检查后重试。
	INVALIDPARAMETER_FLOWAPPROVERINFOS = "InvalidParameter.FlowApproverInfos"

	// 参数错误，参与者数量不能为空且不能超过数量限制，请修改后重试。
	INVALIDPARAMETER_FLOWAPPROVERS = "InvalidParameter.FlowApprovers"

	// 参数错误，不合法的签署流程回调链接，请修改后重试。
	INVALIDPARAMETER_FLOWCALLBACKURL = "InvalidParameter.FlowCallbackUrl"

	// 参数错误，不合法的签署流程截止日期，请修改后重试。
	INVALIDPARAMETER_FLOWDEADLINE = "InvalidParameter.FlowDeadLine"

	// 参数错误，不合法的签署流程描述，请修改后重试。
	INVALIDPARAMETER_FLOWDESCRIPTION = "InvalidParameter.FlowDescription"

	// 参数错误，目前仅支持单个文件发起，请修改后重试。
	INVALIDPARAMETER_FLOWFILEIDS = "InvalidParameter.FlowFileIds"

	// 参数错误，合同id列表长度非法，请检查后重试。
	INVALIDPARAMETER_FLOWIDS = "InvalidParameter.FlowIds"

	// 参数错误，FlowInfos非法，请修改后重试。
	INVALIDPARAMETER_FLOWINFOS = "InvalidParameter.FlowInfos"

	// 参数错误，不合法的签署流程名称，请修改后重试。
	INVALIDPARAMETER_FLOWNAME = "InvalidParameter.FlowName"

	// 参数错误，不合法的FlowType，请修改后重试。
	INVALIDPARAMETER_FLOWTYPE = "InvalidParameter.FlowType"

	// 参数错误，不受支持的GenerateType，请检查后重试。
	INVALIDPARAMETER_GENERATETYPE = "InvalidParameter.GenerateType"

	// 图片不正确。
	INVALIDPARAMETER_IMAGE = "InvalidParameter.Image"

	// 印章名称长度超出。
	INVALIDPARAMETER_LIMITSEALNAME = "InvalidParameter.LimitSealName"

	// 菜单栏状态非法。
	INVALIDPARAMETER_MENUSTATUS = "InvalidParameter.MenuStatus"

	// 缺少必填参数的值。
	INVALIDPARAMETER_MISSINGREQUIREDPARAMETERVALUE = "InvalidParameter.MissingRequiredParameterValue"

	// OpenId不合法。
	INVALIDPARAMETER_OPENID = "InvalidParameter.OpenId"

	// OrganizationId不合法。
	INVALIDPARAMETER_ORGANIZATIONID = "InvalidParameter.OrganizationId"

	// 企业名称不合法。
	INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"

	// 参数错误。
	INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"

	// 文件内容敏感信息。
	INVALIDPARAMETER_SENSITIVEFILECONTENT = "InvalidParameter.SensitiveFileContent"

	// 参数错误，不合法的签署控件类型，请修改后重试。
	INVALIDPARAMETER_SIGNCOMPONENTTYPE = "InvalidParameter.SignComponentType"

	// 状态异常。
	INVALIDPARAMETER_STATUS = "InvalidParameter.Status"

	// 参数错误，不合法的签署顺序，请检查后重试。
	INVALIDPARAMETER_UNORDERED = "InvalidParameter.Unordered"

	// 参数错误，不支持的控件类型，请检查后重试。
	INVALIDPARAMETER_UNSUPPORTEDCOMPONENTTYPE = "InvalidParameter.UnsupportedComponentType"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 超出调用次数限制。
	LIMITEXCEEDED_CALLTIMES = "LimitExceeded.CallTimes"

	// 超出流程创建数量限制。
	LIMITEXCEEDED_CREATEFLOWNUM = "LimitExceeded.CreateFlowNum"

	// 文件大小限制。
	LIMITEXCEEDED_FILESIZE = "LimitExceeded.FileSize"

	// 合同id数超出限制。
	LIMITEXCEEDED_FLOWIDS = "LimitExceeded.FlowIds"

	// 超出流程数量限制。
	LIMITEXCEEDED_FLOWINFOS = "LimitExceeded.FlowInfos"

	// 经办人数据超出。
	LIMITEXCEEDED_PROXYORGANIZATIONOPERATOR = "LimitExceeded.ProxyOrganizationOperator"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 合作企业激活信息不存在。
	MISSINGPARAMETER_COMPANYACTIVEINFO = "MissingParameter.CompanyActiveInfo"

	// 缺少时间参数，请检查后重试。
	MISSINGPARAMETER_DATE = "MissingParameter.Date"

	// 缺少流程id，请检查后重试。
	MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"

	// 缺少流程id，请检查后重试。
	MISSINGPARAMETER_FLOWIDS = "MissingParameter.FlowIds"

	// 请传入需要查询的合同或合同组的ID。
	MISSINGPARAMETER_FLOWIDSORFLOWGROUPID = "MissingParameter.FlowIdsOrFlowGroupId"

	// 未指定流程合同信息。
	MISSINGPARAMETER_FLOWINFO = "MissingParameter.FlowInfo"

	// 缺少控件名称参数，请检查后重试。
	MISSINGPARAMETER_MISSCOMPONENTNAME = "MissingParameter.MissComponentName"

	// 企业OpenId不存在。
	MISSINGPARAMETER_ORGOPENID = "MissingParameter.OrgOpenId"

	// 企业信息为空。
	MISSINGPARAMETER_ORGANIZATIONID = "MissingParameter.OrganizationId"

	// ProxyOperatorOpenId不存在。
	MISSINGPARAMETER_PROXYOPERATOROPENID = "MissingParameter.ProxyOperatorOpenId"

	// 印章图片为空。
	MISSINGPARAMETER_SEALIMAGE = "MissingParameter.SealImage"

	// 印章名称为空。
	MISSINGPARAMETER_SEALNAME = "MissingParameter.SealName"

	// 缺少模板参数。
	MISSINGPARAMETER_TEMPLATES = "MissingParameter.Templates"

	// OpenId不存在。
	MISSINGPARAMETER_USEROPENID = "MissingParameter.UserOpenId"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 不支持的可见性标识。
	OPERATIONDENIED_AUTHTAG = "OperationDenied.AuthTag"

	// 应用号已被禁止。
	OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"

	// 文件发起静默签未开通白名单。
	OPERATIONDENIED_BYFILESSERVERSIGNFORBID = "OperationDenied.ByFilesServerSignForbid"

	// 只支持下载单个。
	OPERATIONDENIED_DOWNLOADMORETHANONE = "OperationDenied.DownLoadMoreThanOne"

	// 无资源访问权限。
	OPERATIONDENIED_ERRNORESOURCEACCESS = "OperationDenied.ErrNoResourceAccess"

	// 文件已删除。
	OPERATIONDENIED_FILEDELETED = "OperationDenied.FileDeleted"

	// 签署流程状态不正确，请检查后重试。
	OPERATIONDENIED_FLOWSTATUSFORBID = "OperationDenied.FlowStatusForbid"

	// 禁止操作。
	OPERATIONDENIED_FORBID = "OperationDenied.Forbid"

	// 没有API权限。
	OPERATIONDENIED_NOAPIAUTH = "OperationDenied.NoApiAuth"

	// 未通过个人实名。
	OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"

	// 流程配额不足。
	OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"

	// 不属于企业超管或者法人。
	OPERATIONDENIED_NOTBELONGSUPERADMINORLEGALPERSON = "OperationDenied.NotBelongSuperAdminOrLegalPerson"

	// 操作类型不支持。
	OPERATIONDENIED_OPERATETYPE = "OperationDenied.OperateType"

	// 超出查询上限。
	OPERATIONDENIED_OUTQUERYLIMIT = "OperationDenied.OutQueryLimit"

	// 用户不归属于当前企业，无法操作，请检查后重试。
	OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"

	// 未开通功能白名单，请联系客服处理。
	OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 应用号不存在。
	RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"

	// 应用授权记录未找到。
	RESOURCENOTFOUND_APPLICATIONAUTH = "ResourceNotFound.ApplicationAuth"

	// ApplicationId不存在。
	RESOURCENOTFOUND_APPLICATIONID = "ResourceNotFound.ApplicationId"

	// 文件不存在。
	RESOURCENOTFOUND_FILE = "ResourceNotFound.File"

	// 未找到对应流程。
	RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"

	// 签署流程的签署人不存在，请检查后重试。
	RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"

	// 签署人信息不存在，请检查后重试。
	RESOURCENOTFOUND_FLOWAPPROVERS = "ResourceNotFound.FlowApprovers"

	// 机构不存在。
	RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"

	// 资源不存在。
	RESOURCENOTFOUND_RESOURCE = "ResourceNotFound.Resource"

	// 印章不存在，请检查后重试。
	RESOURCENOTFOUND_SEAL = "ResourceNotFound.Seal"

	// 合作企业不存在。
	RESOURCENOTFOUND_TEAMWORKORGANIZATION = "ResourceNotFound.TeamWorkOrganization"

	// 模板不存在。
	RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"

	// Url不存在。
	RESOURCENOTFOUND_URL = "ResourceNotFound.URL"

	// 用户信息不存在。
	RESOURCENOTFOUND_USER = "ResourceNotFound.User"

	// 实名用户信息不存在。
	RESOURCENOTFOUND_VERIFYUSER = "ResourceNotFound.VerifyUser"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"
)
