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

package v20240516

const (
	// 此产品的特有错误码

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 交易异常
	FAILEDOPERATION_CALLTRADE = "FailedOperation.CallTrade"

	// FailedOperation.CertificateVerificationFailed
	FAILEDOPERATION_CERTIFICATEVERIFICATIONFAILED = "FailedOperation.CertificateVerificationFailed"

	// 策略优先级重复
	FAILEDOPERATION_DUPLICATEAUTHORIZATIONIDORPRIORITY = "FailedOperation.DuplicateAuthorizationIdOrPriority"

	// DuplicatePolicy
	FAILEDOPERATION_DUPLICATEPOLICY = "FailedOperation.DuplicatePolicy"

	// DuplicatePriority
	FAILEDOPERATION_DUPLICATEPRIORITY = "FailedOperation.DuplicatePriority"

	// FailedOperation.InstanceNotReady
	FAILEDOPERATION_INSTANCENOTREADY = "FailedOperation.InstanceNotReady"

	// FailedOperation.InstanceRegistrationCodeEmpty
	FAILEDOPERATION_INSTANCEREGISTRATIONCODEEMPTY = "FailedOperation.InstanceRegistrationCodeEmpty"

	// 授权策略不支持关闭。
	FAILEDOPERATION_NOTSUPPORTDISABLEAUTHORIZATIONPOLICY = "FailedOperation.NotSupportDisableAuthorizationPolicy"

	// FailedOperation.PublicKeyVerifyFailed
	FAILEDOPERATION_PUBLICKEYVERIFYFAILED = "FailedOperation.PublicKeyVerifyFailed"

	// FailedOperation.RegistrationCodeVerifyFailed
	FAILEDOPERATION_REGISTRATIONCODEVERIFYFAILED = "FailedOperation.RegistrationCodeVerifyFailed"

	// RelatedDeviceCertificateExists
	FAILEDOPERATION_RELATEDDEVICECERTIFICATEEXISTS = "FailedOperation.RelatedDeviceCertificateExists"

	// InstanceTypeNotMatch
	INVALIDPARAMETER_INSTANCETYPENOTMATCH = "InvalidParameter.InstanceTypeNotMatch"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// PublicNetworkInvalidParameterValue
	INVALIDPARAMETERVALUE_PUBLICNETWORKINVALIDPARAMETERVALUE = "InvalidParameterValue.PublicNetworkInvalidParameterValue"

	// LimitExceeded.TopicNum
	LIMITEXCEEDED_TOPICNUM = "LimitExceeded.TopicNum"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 授权策略未找到
	RESOURCENOTFOUND_AUTHORIZATIONPOLICY = "ResourceNotFound.AuthorizationPolicy"

	// ResourceNotFound.Ca
	RESOURCENOTFOUND_CA = "ResourceNotFound.Ca"

	// ResourceNotFound.Certificate
	RESOURCENOTFOUND_CERTIFICATE = "ResourceNotFound.Certificate"

	// ResourceNotFound.Instance
	RESOURCENOTFOUND_INSTANCE = "ResourceNotFound.Instance"

	// ResourceNotFound.Message
	RESOURCENOTFOUND_MESSAGE = "ResourceNotFound.Message"

	// NoAuthenticator
	RESOURCENOTFOUND_NOAUTHENTICATOR = "ResourceNotFound.NoAuthenticator"

	// ResourceNotFound.Role
	RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"

	// ResourceNotFound.Topic
	RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"

	// 用户名未找到
	RESOURCENOTFOUND_USERNAME = "ResourceNotFound.Username"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// UnsupportedOperation.ResourceAlreadyExists
	UNSUPPORTEDOPERATION_RESOURCEALREADYEXISTS = "UnsupportedOperation.ResourceAlreadyExists"
)
