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

package v20180523

const (
	// 此产品的特有错误码

	// 该账号不具有该合同。
	FAILEDOPERATION_ACCOUNTNOTOWNCONTRACTERROR = "FailedOperation.AccountNotOwnContractError"

	// 授权时间格式错误。
	FAILEDOPERATION_AUTHORIZATIONTIMEERROR = "FailedOperation.AuthorizationTimeError"

	// 后端接口调用失败。
	FAILEDOPERATION_BACKENDINTERFACEERROR = "FailedOperation.BackendInterfaceError"

	// 后端接口返回内容异常，请检查请求参数。
	FAILEDOPERATION_BACKENDINTERFACERESPONSECONTENTERROR = "FailedOperation.BackendInterfaceResponseContentError"

	// 后端接口返回头异常，请检查请求参数。
	FAILEDOPERATION_BACKENDINTERFACERESPONSEHEADERERROR = "FailedOperation.BackendInterfaceResponseHeaderError"

	// 证书类型错误。
	FAILEDOPERATION_CERTTYPEERROR = "FailedOperation.CertTypeError"

	// 验证签署验证码失败。
	FAILEDOPERATION_CHECKVCODEERROR = "FailedOperation.CheckVcodeError"

	// 合同已过期。
	FAILEDOPERATION_CONTRACTEXPIRED = "FailedOperation.ContractExpired"

	// 该合同已被签署过。
	FAILEDOPERATION_CONTRACTSIGNEDERROR = "FailedOperation.ContractSignedError"

	// 坐标值错误，必须保证坐标值为数字。
	FAILEDOPERATION_COORDINATEERROR = "FailedOperation.CoordinateError"

	// 坐标值超出PDF尺寸范围。
	FAILEDOPERATION_COORDINATEOUTSIDEPDF = "FailedOperation.CoordinateOutsidePDF"

	// 计费账户问题。
	FAILEDOPERATION_COSTACCOUNTERROR = "FailedOperation.CostAccountError"

	// 合同创建失败。
	FAILEDOPERATION_CREATECONTRACTERROR = "FailedOperation.CreateContractError"

	// 创建企业账户失败，请检查请求参数。
	FAILEDOPERATION_CREATEENTERPRISEACCOUNTERROR = "FailedOperation.CreateEnterpriseAccountError"

	// 创建个人账户失败。
	FAILEDOPERATION_CREATEPERSONALACCOUNTERROR = "FailedOperation.CreatePersonalAccountError"

	// 新增印章失败。
	FAILEDOPERATION_CREATESEALERROR = "FailedOperation.CreateSealError"

	// 删除账户失败。
	FAILEDOPERATION_DELETEACCOUNTERROR = "FailedOperation.DeleteAccountError"

	// 删除印章失败。
	FAILEDOPERATION_DELETESEALERROR = "FailedOperation.DeleteSealError"

	// 获取子平台信息失败。
	FAILEDOPERATION_DESCRIBESUBPLATERROR = "FailedOperation.DescribeSubplatError"

	// 查询任务结果失败。
	FAILEDOPERATION_DESCRIBETASKSTATUSERROR = "FailedOperation.DescribeTaskStatusError"

	// 下载签章文件失败，请确保路径正确。
	FAILEDOPERATION_DOWNLOADSEALERROR = "FailedOperation.DownloadSealError"

	// 企业名称格式错误，不能包含数字。
	FAILEDOPERATION_ENTERPRISENAMEFORMATERROR = "FailedOperation.EnterpriseNameFormatError"

	// 不能删除第一个有效的企业账户。
	FAILEDOPERATION_FIRSTENTERPRISEACCOUNTDELETEERROR = "FailedOperation.FirstEnterpriseAccountDeleteError"

	// 签署关键字信息中存在格式错误(偏移坐标，签章图片宽/高度都必须为数字)。
	FAILEDOPERATION_FORMATERROR = "FailedOperation.FormatError"

	// 获得PDF页面尺寸失败。
	FAILEDOPERATION_GETPDFSIZEFAILED = "FailedOperation.GetPDFSizeFailed"

	// 授权IP地址格式错误。
	FAILEDOPERATION_IPFORMATERROR = "FailedOperation.IPFormatError"

	// 证件号码格式错误，请输入正确的个人/企业证件号码。
	FAILEDOPERATION_IDENTNOFORMATERROR = "FailedOperation.IdentNoFormatError"

	// 证件类型错误，企业证件类型值只允许[0-8]，个人证件类型值为0。
	FAILEDOPERATION_IDENTTYPEERROR = "FailedOperation.IdentTypeError"

	// 图片宽度/高度超过最大像素限制(300px)。
	FAILEDOPERATION_IMAGEMEASUREMENTOVERLIMITERROR = "FailedOperation.ImageMeasurementOverLimitError"

	// 图片数据应为BASE64编码格式。
	FAILEDOPERATION_IMAGENOTBASE = "FailedOperation.ImageNotBase"

	// 图片数据应为PNG格式。
	FAILEDOPERATION_IMAGENOTPNG = "FailedOperation.ImageNotPNG"

	// 报文不可包含以下字符 '%&<>。
	FAILEDOPERATION_MESSAGEDATAILLEGAL = "FailedOperation.MessageDataIllegal"

	// 报文数据长度过长。
	FAILEDOPERATION_MESSAGEDATAOVERSIZE = "FailedOperation.MessageDataOverSize"

	// 用户名称不能包含数字。
	FAILEDOPERATION_NAMECONTAINSNUMBER = "FailedOperation.NameContainsNumber"

	// 企业名称不能为纯数字。
	FAILEDOPERATION_NAMEISPURENUMBER = "FailedOperation.NameIsPureNumber"

	// 用户没有开通代签权限。
	FAILEDOPERATION_NOPERMISSIONTOSIGN = "FailedOperation.NoPermissionToSign"

	// 未经短信验证。
	FAILEDOPERATION_NOVERIFYERROR = "FailedOperation.NoVerifyError"

	// 偏移坐标超过最大偏移量，X轴/Y轴最大偏移量为50。
	FAILEDOPERATION_OFFSETCOORDOVERLIMITERROR = "FailedOperation.OffsetCoordOverLimitError"

	// 其它。
	FAILEDOPERATION_OTHER = "FailedOperation.Other"

	// 签名域对角坐标值不能重复。
	FAILEDOPERATION_REPEATEDCOORDINATE = "FailedOperation.RepeatedCoordinate"

	// 短信验证码已超时。
	FAILEDOPERATION_SMSCODEEXPIRED = "FailedOperation.SMSCodeExpired"

	// 短信校验码长度错误。
	FAILEDOPERATION_SMSCODELENGTHWRONG = "FailedOperation.SMSCodeLengthWrong"

	// 印章不匹配。
	FAILEDOPERATION_SEALMISMATCHED = "FailedOperation.SealMismatched"

	// 印章数超过上限。
	FAILEDOPERATION_SEALNUMOVERLIMIT = "FailedOperation.SealNumOverLimit"

	// 印章数量超出限制。
	FAILEDOPERATION_SEALSEXCEED = "FailedOperation.SealsExceed"

	// 发送签署验证码失败。
	FAILEDOPERATION_SENDVCODEERROR = "FailedOperation.SendVcodeError"

	// 按坐标签署合同失败。
	FAILEDOPERATION_SIGNCONTRACTBYCOORDINATEERROR = "FailedOperation.SignContractByCoordinateError"

	// 按关键字签署合同失败。
	FAILEDOPERATION_SIGNCONTRACTBYKEYWORDERROR = "FailedOperation.SignContractByKeywordError"

	// 根据关键字获得PDF签名域失败。
	FAILEDOPERATION_SIGNFIELDNOTFOUND = "FailedOperation.SignFieldNotFound"

	// 页数必须是自然数。
	FAILEDOPERATION_SIGNPAGEERROR = "FailedOperation.SignPageError"

	// 已经开通代签权限。
	FAILEDOPERATION_SIGNPERMISSIONEXISTED = "FailedOperation.SignPermissionExisted"

	// 经办人姓名格式错误，不能包含数字。
	FAILEDOPERATION_TRANSACTORNAMEFORMATERROR = "FailedOperation.TransactorNameFormatError"

	// 经办人手机号格式错误，必须为全数字。
	FAILEDOPERATION_TRANSACTORPHONEFORMATERROR = "FailedOperation.TransactorPhoneFormatError"

	// 更新计费状态错误。
	FAILEDOPERATION_UPDATEFEESTATUSERROR = "FailedOperation.UpdateFeeStatusError"

	// 已经开通代签权限。
	FAILEDOPERATION_VCODECHECKED = "FailedOperation.VcodeChecked"

	// 证书类型为不支持的类型。
	FAILEDOPERATION_WRONGCERTTYPE = "FailedOperation.WrongCertType"

	// 身份证号格式错误。
	FAILEDOPERATION_WRONGIDENTNOFORMAT = "FailedOperation.WrongIdentNoFormat"

	// 身份证号只能为15或18位字符。
	FAILEDOPERATION_WRONGIDENTNOSIZE = "FailedOperation.WrongIdentNoSize"

	// 短信验证码不正确。
	FAILEDOPERATION_WRONGSMSCODE = "FailedOperation.WrongSMSCode"

	// 内部错误。
	INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"

	// 合同名称不能为空。
	MISSINGPARAMETER_CONTRACTFILENAMEERROR = "MissingParameter.ContractFileNameError"

	// 合同文件路径不能为空。
	MISSINGPARAMETER_CONTRACTFILEPATHERROR = "MissingParameter.ContractFilePathError"

	// 签署关键字信息-签章图片宽度/高度不能为空。
	MISSINGPARAMETER_IMAGEMEASUREMENTNULLERROR = "MissingParameter.ImageMeasurementNullError"

	// 签署关键字信息-关键字不能为空。
	MISSINGPARAMETER_KEYWORDNULLERROR = "MissingParameter.KeywordNullError"

	// 授权IP地址不能为空。
	MISSINGPARAMETER_LOCATIONNULLERROR = "MissingParameter.LocationNullError"

	// 找不到 module 或 operation。
	MISSINGPARAMETER_MOERROR = "MissingParameter.MOError"

	// 签署关键字信息-偏移坐标不能为空。
	MISSINGPARAMETER_OFFSETCOORDNULLERROR = "MissingParameter.OffsetCoordNullError"

	// 签署人不能为空。
	MISSINGPARAMETER_SIGNERNULLERROR = "MissingParameter.SignerNullError"

	// 已经开过该账户。
	RESOURCEINUSE_ACCOUNTEXIST = "ResourceInUse.AccountExist"

	// 企业账户已存在。
	RESOURCEINUSE_ENTERPRISEACCOUNTALREADYEXIST = "ResourceInUse.EnterpriseAccountAlreadyExist"

	// 个人账户已存在。
	RESOURCEINUSE_PERSONACCOUNTALREADYEXIST = "ResourceInUse.PersonAccountAlreadyExist"

	// 账户不存在或已删除。
	RESOURCENOTFOUND_ACCOUNTNOTFOUND = "ResourceNotFound.AccountNotFound"

	// 合同ID不存在。
	RESOURCENOTFOUND_CONTRACTNOTFOUND = "ResourceNotFound.ContractNotFound"

	// 合同project_code不存在。
	RESOURCENOTFOUND_CONTRACTPROJECTCODENOTFOUND = "ResourceNotFound.ContractProjectCodeNotFound"

	// 合同发起方账号不存在。
	RESOURCENOTFOUND_INITIATORNOTFOUNDERROR = "ResourceNotFound.InitiatorNotFoundError"

	// 印章不存在。
	RESOURCENOTFOUND_SEALNOTFOUND = "ResourceNotFound.SealNotFound"

	// 子平台ID不存在。
	RESOURCENOTFOUND_SUBPLATIDNOTFOUND = "ResourceNotFound.SubplatIdNotFound"

	// 任务不存在。
	RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"

	// 账户状态已冻结或已注销。
	RESOURCEUNAVAILABLE_ACCOUNTUNAVAILABLE = "ResourceUnavailable.AccountUnavailable"

	// 合同签署人不存在或状态异常，不能进行合同创建操作。
	RESOURCEUNAVAILABLE_CONTRACTSIGNERUNAVAILABLE = "ResourceUnavailable.ContractSignerUnavailable"

	// 下载签章文件失败，请确保路径正确。
	RESOURCEUNAVAILABLE_DOWNLOADSEALERROR = "ResourceUnavailable.DownloadSealError"

	// 子平台欠费。
	RESOURCEUNAVAILABLE_SUBPLATUNAVAILABLE = "ResourceUnavailable.SubplatUnavailable"
)
