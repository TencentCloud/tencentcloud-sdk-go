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

	// 证书已存在。
	FAILEDOPERATION_CERTIFICATEEXISTS = "FailedOperation.CertificateExists"

	// 已替换证书，无法进行托管。
	FAILEDOPERATION_CERTIFICATEHASRENEWED = "FailedOperation.CertificateHasRenewed"

	// 证书资源托管数量超过限制。
	FAILEDOPERATION_CERTIFICATEHOSTINGTYPENUMBERLIMIT = "FailedOperation.CertificateHostingTypeNumberLimit"

	// 证书不符合标准。
	FAILEDOPERATION_CERTIFICATEINVALID = "FailedOperation.CertificateInvalid"

	// 证书与私钥不对应。
	FAILEDOPERATION_CERTIFICATEMISMATCH = "FailedOperation.CertificateMismatch"

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

	// 权益点ID列表无效。
	INVALIDPARAMETER_PACKAGEIDSINVALID = "InvalidParameter.PackageIdsInvalid"

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
