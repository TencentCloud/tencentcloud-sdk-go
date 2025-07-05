// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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

package v20210519

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 机构名称要求输入中英文或数字，不超过32个字符
	AUTHORITY_AUTHORITYNAMELENGTHLONG = "Authority.AuthorityNameLengthLong"

	// 认证备注要求输入中英文或数字，不超过64个字符
	AUTHORITY_AUTHORITYREMARKLENGTHLONG = "Authority.AuthorityRemarkLengthLong"

	// 该权威机构已被认证:<%ExtMsg%>
	AUTHORITY_ERRCODEAUTHORITYREGISTERED = "Authority.ErrCodeAuthorityRegistered"

	// 该DID已经认证为权威机构:<%ExtMsg%>
	AUTHORITY_ERRCODEDIDAUTHORITYREGISTERED = "Authority.ErrCodeDidAuthorityRegistered"

	// CAM鉴权失败
	CAM_INVALIDAUTH = "Cam.InvalidAuth"

	// 创建凭证模板失败，请重新操作
	CPT_CREATECPTFAILED = "Cpt.CreateCPTFailed"

	// 模板内容为无效的json格式
	CPT_INVALIDCPTJSON = "Cpt.InvalidCPTJson"

	// 凭证模板不存在
	CREDENTIAL_CPTNOTEXISTED = "Credential.CPTNotExisted"

	// 凭证过期:<%ExtMsg%>
	CREDENTIAL_CREDENTIALEXPIRED = "Credential.CredentialExpired"

	// 创建凭证表达失败
	CREDENTIAL_FAILEDCREATEPRESENTATION = "Credential.FailedCreatePresentation"

	// 生成证明承诺失败
	CREDENTIAL_GENERATECOMMITMENTFAILED = "Credential.GenerateCommitmentFailed"

	// 生成范围证明失败
	CREDENTIAL_GENERATERANGEPROOF = "Credential.GenerateRangeProof"

	// 凭证持有人DID不匹配，请确认后重新操作:<%ExtMsg%>
	CREDENTIAL_HOLDERDIDNOTMATCH = "Credential.HolderDidNotMatch"

	// 无效的凭证ID，请确认后重新操作:<%ExtMsg%>
	CREDENTIAL_INVALIDCRDLID = "Credential.InvalidCRDLId"

	// 无效的凭证颁发者，请确认后重新操作
	CREDENTIAL_INVALIDCRDLISSUER = "Credential.InvalidCRDLIssuer"

	// 无效的凭证声明，请确认声明格式:<%ExtMsg%>
	CREDENTIAL_INVALIDCLAIM = "Credential.InvalidClaim"

	// 无效的选择性披露策略，请确认后重新操作:<%ExtMsg%>
	CREDENTIAL_INVALIDDISCLOSUREPOLICY = "Credential.InvalidDisclosurePolicy"

	// 无效的操作类型声明，请确认后重新操作:<%ExtMsg%>
	CREDENTIAL_INVALIDOPERATECLAIM = "Credential.InvalidOperateClaim"

	// 无效请求参数，请确认参数格式:<%ExtMsg%>
	CREDENTIAL_INVALIDPARAS = "Credential.InvalidParas"

	// 无效的证明承诺
	CREDENTIAL_INVALIDPROOFCOMMITMENT = "Credential.InvalidProofCommitment"

	// 无效的证明值，证明值要求是无符号整形
	CREDENTIAL_INVALIDPROOFVALUE = "Credential.InvalidProofValue"

	// 无效的零知识证明内容
	CREDENTIAL_INVALIDZEROPROOF = "Credential.InvalidZeroProof"

	// 凭证颁发者DID不匹配，请确认后重新操作:<%ExtMsg%>
	CREDENTIAL_ISSUERDIDNOTMATCH = "Credential.IssuerDidNotMatch"

	// 凭证声明缺少ID字段，请确认后重新操作
	CREDENTIAL_NOIDINCLAIM = "Credential.NoIdInClaim"

	// 证明数值不存在
	CREDENTIAL_PROOFVALUENOTEXISTED = "Credential.ProofValueNotExisted"

	// 凭证验证失败:<%ExtMsg%>
	CREDENTIAL_VERIFYCRDLFAILED = "Credential.VerifyCRDLFailed"

	// 零知识证明验证失败
	CREDENTIAL_VERIFYZEROPROOFFAILED = "Credential.VerifyZeroProofFailed"

	// 数据库操作失败，请重新操作
	DATABASE_FAILEDOPERATION = "DataBase.FailedOperation"

	// 您无操作权限
	DID_PERMISSIONDENIED = "Did.PermissionDenied"

	// DID标识已存在
	DIDFAILEDOPERATION_DIDEXISTED = "DidFailedOperation.DidExisted"

	// DID标识不存在:<%ExtMsg%>
	DIDFAILEDOPERATION_DIDNOTEXISTED = "DidFailedOperation.DidNotExisted"

	// DID属性objectId已存在
	DIDFAILEDOPERATION_DIDOBJECTIDEXISTED = "DidFailedOperation.DidObjectIdExisted"

	// DID状态无效
	DIDFAILEDOPERATION_DIDSTATUSINVALID = "DidFailedOperation.DidStatusInvalid"

	// 用户尚未创建DID服务，请确认后重新操作
	DIDFAILEDOPERATION_DIDSVCNOTEXISTED = "DidFailedOperation.DidSvcNotExisted"

	// 获取DID文档失败，请重新操作
	DIDFAILEDOPERATION_GETDIDDOCFILED = "DidFailedOperation.GetDidDocFiled"

	// 非权威机构DID:<%ExtMsg%>
	DIDFAILEDOPERATION_NOTAUTHORITY = "DidFailedOperation.NotAuthority"

	// 非DID创建者，无法更新DID状态
	DIDFAILEDOPERATION_NOTDIDCREATOR = "DidFailedOperation.NotDidCreator"

	// DID私钥格式错误，请填写正确的DID私钥
	DIDFAILEDOPERATION_PRIVATEKEYINVALID = "DidFailedOperation.PrivateKeyInvalid"

	// DID公钥格式错误，请填写正确的DID公钥
	DIDFAILEDOPERATION_PUBLICKEYINVALID = "DidFailedOperation.PublicKeyInvalid"

	// DID服务请求错误，请稍后重试或联系客服
	DIDSDK_UNKNOWNERROR = "DidSdk.UnknownError"

	// 用户接口鉴权失败
	FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"

	// 无效的请求参数
	FAILEDOPERATION_INVALIDPARAMETERVALUE = "FailedOperation.InvalidParameterValue"

	// 操作失败。
	FAILEDOPERATION_OPERATIONEXCEPTION = "FailedOperation.OperationException"

	// 用户无接口访问权限
	FAILEDOPERATION_PERMISSIONDENIED = "FailedOperation.PermissionDenied"

	// HTTP请求创建失败，请稍后重试，或联系客服
	HTTPINVOKER_NEWQUESTERROR = "HttpInvoker.NewQuestError"

	// HTTP请求发送失败，请稍后重试，或联系客服
	HTTPINVOKER_SENDQUESTERROR = "HttpInvoker.SendQuestError"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// API服务出错
	INTERNALERROR_APIUNKNOWNERROR = "InternalError.ApiUnknownError"

	// 服务器异常。
	INTERNALERROR_SERVEREXCEPTION = "InternalError.ServerException"

	// 服务已停止，请您检查当前是否有计费套餐包
	INTERNALERROR_SERVICEDISABLED = "InternalError.ServiceDisabled"

	// 服务错误，请稍后重试，或联系客服
	INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"

	// 服务异常。
	INTERNALERROR_SERVICEPANIC = "InternalError.ServicePanic"

	// 服务错误，请稍后重试，或联系客服
	INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 无效的时间格式，请确认后重新操作
	INVALIDPARAMETERVALUE_ILLEGALDATETIME = "InvalidParameterValue.IllegalDateTime"

	// 请求参数格式错误，请按照格式要求重新填写
	INVALIDPARAMETERVALUE_ILLEGALVALUE = "InvalidParameterValue.IllegalValue"

	// 用户请求资源未存在:<%ExtMsg%>
	INVALIDPARAMETERVALUE_RESOURCENOTEXISTED = "InvalidParameterValue.ResourceNotExisted"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"
)
