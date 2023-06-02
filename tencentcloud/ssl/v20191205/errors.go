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

package v20191205

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 操作未授权。
	AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 请检查是否有权限。
	FAILEDOPERATION_AUTHERROR = "FailedOperation.AuthError"

	// CAM鉴权出现错误。
	FAILEDOPERATION_CAMAUTHORIZEDFAIL = "FailedOperation.CAMAuthorizedFail"

	// 取消订单失败。
	FAILEDOPERATION_CANCELORDERFAILED = "FailedOperation.CancelOrderFailed"

	// 该证书已颁发， 不能删除。
	FAILEDOPERATION_CANNOTBEDELETEDISSUED = "FailedOperation.CannotBeDeletedIssued"

	// 免费证书申请1小时内不允许删除。
	FAILEDOPERATION_CANNOTBEDELETEDWITHINHOUR = "FailedOperation.CannotBeDeletedWithinHour"

	// 获取订单信息失败，请稍后重试。
	FAILEDOPERATION_CANNOTGETORDER = "FailedOperation.CannotGetOrder"

	// 记录状态必须完结才可以执行该操作。
	FAILEDOPERATION_CERTIFICATEDEPLOYDETAILROLLBACKSTATUSINVALID = "FailedOperation.CertificateDeployDetailRollbackStatusInvalid"

	// 证书部署有正在进行中的任务，请部署完成后再重试。
	FAILEDOPERATION_CERTIFICATEDEPLOYHASPENDINGRECORD = "FailedOperation.CertificateDeployHasPendingRecord"

	// 已选择的云资源无实例，无法更新，请重新核对后重试。
	FAILEDOPERATION_CERTIFICATEDEPLOYINSTANCEEMPTY = "FailedOperation.CertificateDeployInstanceEmpty"

	// 证书部署记录不存在。
	FAILEDOPERATION_CERTIFICATEDEPLOYNOTEXIST = "FailedOperation.CertificateDeployNotExist"

	// 记录状态必须失败才可以执行该操作。
	FAILEDOPERATION_CERTIFICATEDEPLOYRETRYSTATUSINVALID = "FailedOperation.CertificateDeployRetryStatusInvalid"

	// 必须有部署成功的记录才可以回滚。
	FAILEDOPERATION_CERTIFICATEDEPLOYROLLBACKSTATUSINVALID = "FailedOperation.CertificateDeployRollbackStatusInvalid"

	// 证书已存在。
	FAILEDOPERATION_CERTIFICATEEXISTS = "FailedOperation.CertificateExists"

	// 已替换证书，无法进行托管。
	FAILEDOPERATION_CERTIFICATEHASRENEWED = "FailedOperation.CertificateHasRenewed"

	// 当前证书不允许使用一键更新的功能。
	FAILEDOPERATION_CERTIFICATEHOSTDEPLOYCANNOTALLOW = "FailedOperation.CertificateHostDeployCanNotAllow"

	// 当前为内部账号，账号涉及实例资源较多，无法使用部署功能，请联系SSL证书特殊处理。
	FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINNERINTERRUPT = "FailedOperation.CertificateHostResourceInnerInterrupt"

	// 当前账号下实例数量过多，无法正常加载，请您切换加载方式。切换后点击“刷新列表”等待一段时间后即可全部加载。
	FAILEDOPERATION_CERTIFICATEHOSTRESOURCEINSTANCEHUGELIMIT = "FailedOperation.CertificateHostResourceInstanceHugeLimit"

	// 云资源类型无效。
	FAILEDOPERATION_CERTIFICATEHOSTRESOURCETYPEINVALID = "FailedOperation.CertificateHostResourceTypeInvalid"

	// 证书资源托管数量超过限制。
	FAILEDOPERATION_CERTIFICATEHOSTINGTYPENUMBERLIMIT = "FailedOperation.CertificateHostingTypeNumberLimit"

	// 证书不符合标准。
	FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"

	// 证书与私钥不对应。
	FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"

	// 证书不可用，请检查后再试。
	FAILEDOPERATION_CERTIFICATENOTAVAILABLE = "FailedOperation.CertificateNotAvailable"

	// 证书不可以部署到实例列表下。
	FAILEDOPERATION_CERTIFICATENOTDEPLOYINSTANCE = "FailedOperation.CertificateNotDeployInstance"

	// 证书不存在。
	FAILEDOPERATION_CERTIFICATENOTFOUND = "FailedOperation.CertificateNotFound"

	// 证书确认函文件过大（需小于1.4M）。
	FAILEDOPERATION_CONFIRMLETTERTOOLARGE = "FailedOperation.ConfirmLetterTooLarge"

	// 证书确认函文件过小（需大于1KB）。
	FAILEDOPERATION_CONFIRMLETTERTOOSMALL = "FailedOperation.ConfirmLetterTooSmall"

	// 证书已关联云资源，无法删除。
	FAILEDOPERATION_DELETERESOURCEFAILED = "FailedOperation.DeleteResourceFailed"

	// 免费证书数量超出限制。
	FAILEDOPERATION_EXCEEDSFREELIMIT = "FailedOperation.ExceedsFreeLimit"

	// 文件尺寸太大（需小于1.4MB）。
	FAILEDOPERATION_FILETOOLARGE = "FailedOperation.FileTooLarge"

	// 文件尺寸太小，请上传清晰图片。
	FAILEDOPERATION_FILETOOSMALL = "FailedOperation.FileTooSmall"

	// 公司管理人状态错误。
	FAILEDOPERATION_ILLEGALMANAGERSTATUS = "FailedOperation.IllegalManagerStatus"

	// 证书来源错误。
	FAILEDOPERATION_INVALIDCERTIFICATESOURCE = "FailedOperation.InvalidCertificateSource"

	// 证书状态不正确。
	FAILEDOPERATION_INVALIDCERTIFICATESTATUSCODE = "FailedOperation.InvalidCertificateStatusCode"

	// 证书确认函格式错误（支持格式为jpg、jpeg、png、pdf）。
	FAILEDOPERATION_INVALIDCONFIRMLETTERFORMAT = "FailedOperation.InvalidConfirmLetterFormat"

	// 证书确认函格式错误（支持格式为jpg、pdf、gif）。
	FAILEDOPERATION_INVALIDCONFIRMLETTERFORMATWOSIGN = "FailedOperation.InvalidConfirmLetterFormatWosign"

	// 文件格式错误，请重新上传。
	FAILEDOPERATION_INVALIDFILETYPE = "FailedOperation.InvalidFileType"

	// 参数有误。
	FAILEDOPERATION_INVALIDPARAM = "FailedOperation.InvalidParam"

	// 该主域（%s）下申请的免费证书数量已达到%s个上限，请购买付费证书。
	FAILEDOPERATION_MAINDOMAINCERTIFICATECOUNTLIMIT = "FailedOperation.MainDomainCertificateCountLimit"

	// 管理人信息已提交CA，不可以删除。
	FAILEDOPERATION_MANAGERCANNOTDELETECA = "FailedOperation.ManagerCanNotDeleteCa"

	// 管理人信息已关联证书，不可以删除。
	FAILEDOPERATION_MANAGERCANNOTDELETECERT = "FailedOperation.ManagerCanNotDeleteCert"

	// 当前 CA 机构访问繁忙，请稍后重试。
	FAILEDOPERATION_NETWORKERROR = "FailedOperation.NetworkError"

	// 您没有该项目的操作权限。
	FAILEDOPERATION_NOPROJECTPERMISSION = "FailedOperation.NoProjectPermission"

	// 尚未通过实名认证。
	FAILEDOPERATION_NOREALNAMEAUTH = "FailedOperation.NoRealNameAuth"

	// 该订单已重签发。
	FAILEDOPERATION_ORDERALREADYREPLACED = "FailedOperation.OrderAlreadyReplaced"

	// 重颁发失败。
	FAILEDOPERATION_ORDERREPLACEFAILED = "FailedOperation.OrderReplaceFailed"

	// 剩余权益点不足。
	FAILEDOPERATION_PACKAGECOUNTLIMIT = "FailedOperation.PackageCountLimit"

	// 权益包已过期。
	FAILEDOPERATION_PACKAGEEXPIRED = "FailedOperation.PackageExpired"

	// 证书吊销失败。
	FAILEDOPERATION_REVOKEFAILED = "FailedOperation.RevokeFailed"

	// 证书已关联云资源，无法吊销。
	FAILEDOPERATION_REVOKERESOURCEFAILED = "FailedOperation.RevokeResourceFailed"

	// 角色不存在，请前往授权。
	FAILEDOPERATION_ROLENOTFOUNDAUTHORIZATION = "FailedOperation.RoleNotFoundAuthorization"

	// 系统错误。
	FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"

	// 计费中心错误。
	FAILEDOPERATION_TRADEERROR = "FailedOperation.TradeError"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 后端服务响应为空。
	INTERNALERROR_BACKENDRESPONSEEMPTY = "InternalError.BackendResponseEmpty"

	// 后端服务响应错误。
	INTERNALERROR_BACKENDRESPONSEERROR = "InternalError.BackendResponseError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 查询的证书ID不能超过50个。
	INVALIDPARAMETER_CERTIFICATEIDNUMBERLIMIT = "InvalidParameter.CertificateIdNumberLimit"

	// 证书数量超出限制。
	INVALIDPARAMETER_CERTIFICATESNUMBEREXCEEDED = "InvalidParameter.CertificatesNumberExceeded"

	// 包含无效的证书ID。
	INVALIDPARAMETER_CONTAINSINVALIDCERTIFICATEID = "InvalidParameter.ContainsInvalidCertificateId"

	// 域名数量无效。
	INVALIDPARAMETER_DOMAINCOUNTINVALID = "InvalidParameter.DomainCountInvalid"

	// 域名无效，请重新输入。
	INVALIDPARAMETER_DOMAININVALID = "InvalidParameter.DomainInvalid"

	// 权益点ID列表无效。
	INVALIDPARAMETER_PACKAGEIDSINVALID = "InvalidParameter.PackageIdsInvalid"

	// 证书期限无效。
	INVALIDPARAMETER_PERIODINVALID = "InvalidParameter.PeriodInvalid"

	// 产品PID无效。
	INVALIDPARAMETER_PRODUCTPIDINVALID = "InvalidParameter.ProductPidInvalid"

	// 算法无效。
	INVALIDPARAMETER_RENEWALGORITHMINVALID = "InvalidParameter.RenewAlgorithmInvalid"

	// CSR无效。
	INVALIDPARAMETER_RENEWCSRINVALID = "InvalidParameter.RenewCsrInvalid"

	// 生成CSR方式无效。
	INVALIDPARAMETER_RENEWGENCSRMETHODINVALID = "InvalidParameter.RenewGenCsrMethodInvalid"

	// 参数有误。
	INVALIDPARAMETER_WITHDETAILREASON = "InvalidParameter.WithDetailReason"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 单位时间内接口请求频率达到限制。
	LIMITEXCEEDED_RATELIMITEXCEEDED = "LimitExceeded.RateLimitExceeded"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 公司管理人不存在。
	RESOURCENOTFOUND_MANAGER = "ResourceNotFound.Manager"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"
)
