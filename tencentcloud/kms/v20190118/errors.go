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

package v20190118

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// CMK正在被云产品使用。
	FAILEDOPERATION_CMKUSEDBYCLOUDPRODUCT = "FailedOperation.CmkUsedByCloudProduct"

	// 解密失败。
	FAILEDOPERATION_DECRYPTERROR = "FailedOperation.DecryptError"

	// 加密操作失败。
	FAILEDOPERATION_ENCRYPTIONERROR = "FailedOperation.EncryptionError"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 解密EncryptedKeyMaterial失败。
	INVALIDPARAMETER_DECRYPTMATERIALERROR = "InvalidParameter.DecryptMaterialError"

	// 计划删除时间参数非法。
	INVALIDPARAMETER_INVALIDPENDINGWINDOWINDAYS = "InvalidParameter.InvalidPendingWindowInDays"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 别名已经存在。
	INVALIDPARAMETERVALUE_ALIASALREADYEXISTS = "InvalidParameterValue.AliasAlreadyExists"

	// KeyId重复。
	INVALIDPARAMETERVALUE_DUPLICATEDKEYID = "InvalidParameterValue.DuplicatedKeyId"

	// 别名格式错误。
	INVALIDPARAMETERVALUE_INVALIDALIAS = "InvalidParameterValue.InvalidAlias"

	// 密文格式错误。
	INVALIDPARAMETERVALUE_INVALIDCIPHERTEXT = "InvalidParameterValue.InvalidCiphertext"

	// KeyId不合法。
	INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"

	// KeyUsage参数错误。
	INVALIDPARAMETERVALUE_INVALIDKEYUSAGE = "InvalidParameterValue.InvalidKeyUsage"

	// Plaintext不合法。
	INVALIDPARAMETERVALUE_INVALIDPLAINTEXT = "InvalidParameterValue.InvalidPlaintext"

	// Type参数错误。
	INVALIDPARAMETERVALUE_INVALIDTYPE = "InvalidParameterValue.InvalidType"

	// 导入的密钥材料和历史导入不同。
	INVALIDPARAMETERVALUE_MATERIALNOTMATCH = "InvalidParameterValue.MaterialNotMatch"

	// 标签键重复。
	INVALIDPARAMETERVALUE_TAGKEYSDUPLICATED = "InvalidParameterValue.TagKeysDuplicated"

	// 标签键或标签值不存在。
	INVALIDPARAMETERVALUE_TAGSNOTEXISTED = "InvalidParameterValue.TagsNotExisted"

	// CMK数量已达上限。
	LIMITEXCEEDED_CMKLIMITEXCEEDED = "LimitExceeded.CmkLimitExceeded"

	// 设备指纹个数超过限制。
	LIMITEXCEEDED_FINGERPRINTSLIMITEXCEEDED = "LimitExceeded.FingerprintsLimitExceeded"

	// 创建的密钥个数超过限制。
	LIMITEXCEEDED_KEYLIMITEXCEEDED = "LimitExceeded.KeyLimitExceeded"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// cmk和云资源的绑定关系不存在。
	RESOURCEUNAVAILABLE_CLOUDRESOURCEBINDINGNOTFOUND = "ResourceUnavailable.CloudResourceBindingNotFound"

	// CMK已存档。
	RESOURCEUNAVAILABLE_CMKARCHIVED = "ResourceUnavailable.CmkArchived"

	// CMK已被禁用。
	RESOURCEUNAVAILABLE_CMKDISABLED = "ResourceUnavailable.CmkDisabled"

	// CMK不存在。
	RESOURCEUNAVAILABLE_CMKNOTFOUND = "ResourceUnavailable.CmkNotFound"

	// CMK不是计划删除状态不能被执行取消计划删除。
	RESOURCEUNAVAILABLE_CMKNOTPENDINGDELETE = "ResourceUnavailable.CmkNotPendingDelete"

	// 未被禁用的CMK不能被计划删除。
	RESOURCEUNAVAILABLE_CMKSHOULDBEDISABLED = "ResourceUnavailable.CmkShouldBeDisabled"

	// CMK 状态不支持该操作。
	RESOURCEUNAVAILABLE_CMKSTATENOTSUPPORT = "ResourceUnavailable.CmkStateNotSupport"

	// 密钥已被禁用。
	RESOURCEUNAVAILABLE_KEYDISABLED = "ResourceUnavailable.KeyDisabled"

	// 不可用密钥：密钥待删除。
	RESOURCEUNAVAILABLE_KEYPENDINGDELETE = "ResourceUnavailable.KeyPendingDelete"

	// 白盒密钥服务尚未开通。
	RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"

	// Token已过期。
	RESOURCEUNAVAILABLE_TOKENEXPIRED = "ResourceUnavailable.TokenExpired"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 用户导入类型的CMK禁止轮换。
	UNSUPPORTEDOPERATION_EXTERNALCMKCANNOTROTATE = "UnsupportedOperation.ExternalCmkCanNotRotate"

	// CMK类型错误，仅支持External CMK。
	UNSUPPORTEDOPERATION_NOTEXTERNALCMK = "UnsupportedOperation.NotExternalCmk"

	// 仅支持对用户自己创建的CMK做更新。
	UNSUPPORTEDOPERATION_NOTUSERCREATEDCMK = "UnsupportedOperation.NotUserCreatedCmk"

	// 服务暂时不可用。
	UNSUPPORTEDOPERATION_SERVICETEMPORARYUNAVAILABLE = "UnsupportedOperation.ServiceTemporaryUnavailable"

	// 加密方式在当前region不支持。
	UNSUPPORTEDOPERATION_UNSUPPORTEDKEYUSAGEINCURRENTREGION = "UnsupportedOperation.UnsupportedKeyUsageInCurrentRegion"
)
