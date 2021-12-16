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

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-01-18"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewArchiveKeyRequest() (request *ArchiveKeyRequest) {
    request = &ArchiveKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "ArchiveKey")
    
    
    return
}

func NewArchiveKeyResponse() (response *ArchiveKeyResponse) {
    response = &ArchiveKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ArchiveKey
// 对密钥进行归档，被归档的密钥只能用于解密，不能加密
//
// 可能返回的错误码:
//  FAILEDOPERATION_CMKUSEDBYCLOUDPRODUCT = "FailedOperation.CmkUsedByCloudProduct"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  RESOURCEUNAVAILABLE_CMKNOTFOUND = "ResourceUnavailable.CmkNotFound"
//  RESOURCEUNAVAILABLE_CMKSTATENOTSUPPORT = "ResourceUnavailable.CmkStateNotSupport"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_NOTUSERCREATEDCMK = "UnsupportedOperation.NotUserCreatedCmk"
//  UNSUPPORTEDOPERATION_SERVICETEMPORARYUNAVAILABLE = "UnsupportedOperation.ServiceTemporaryUnavailable"
func (c *Client) ArchiveKey(request *ArchiveKeyRequest) (response *ArchiveKeyResponse, err error) {
    if request == nil {
        request = NewArchiveKeyRequest()
    }
    
    response = NewArchiveKeyResponse()
    err = c.Send(request, response)
    return
}

func NewAsymmetricRsaDecryptRequest() (request *AsymmetricRsaDecryptRequest) {
    request = &AsymmetricRsaDecryptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "AsymmetricRsaDecrypt")
    
    
    return
}

func NewAsymmetricRsaDecryptResponse() (response *AsymmetricRsaDecryptResponse) {
    response = &AsymmetricRsaDecryptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AsymmetricRsaDecrypt
// 使用指定的RSA非对称密钥的私钥进行数据解密，密文必须是使用对应公钥加密的。处于Enabled 状态的非对称密钥才能进行解密操作。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DECRYPTERROR = "FailedOperation.DecryptError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  INVALIDPARAMETERVALUE_INVALIDKEYUSAGE = "InvalidParameterValue.InvalidKeyUsage"
//  RESOURCEUNAVAILABLE_CMKNOTFOUND = "ResourceUnavailable.CmkNotFound"
//  RESOURCEUNAVAILABLE_CMKSTATENOTSUPPORT = "ResourceUnavailable.CmkStateNotSupport"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AsymmetricRsaDecrypt(request *AsymmetricRsaDecryptRequest) (response *AsymmetricRsaDecryptResponse, err error) {
    if request == nil {
        request = NewAsymmetricRsaDecryptRequest()
    }
    
    response = NewAsymmetricRsaDecryptResponse()
    err = c.Send(request, response)
    return
}

func NewAsymmetricSm2DecryptRequest() (request *AsymmetricSm2DecryptRequest) {
    request = &AsymmetricSm2DecryptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "AsymmetricSm2Decrypt")
    
    
    return
}

func NewAsymmetricSm2DecryptResponse() (response *AsymmetricSm2DecryptResponse) {
    response = &AsymmetricSm2DecryptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AsymmetricSm2Decrypt
// 使用指定的SM2非对称密钥的私钥进行数据解密，密文必须是使用对应公钥加密的。处于Enabled 状态的非对称密钥才能进行解密操作。传入的密文的长度不能超过256字节。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DECRYPTERROR = "FailedOperation.DecryptError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  INVALIDPARAMETERVALUE_INVALIDKEYUSAGE = "InvalidParameterValue.InvalidKeyUsage"
//  RESOURCEUNAVAILABLE_CMKNOTFOUND = "ResourceUnavailable.CmkNotFound"
//  RESOURCEUNAVAILABLE_CMKSTATENOTSUPPORT = "ResourceUnavailable.CmkStateNotSupport"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDKEYUSAGEINCURRENTREGION = "UnsupportedOperation.UnsupportedKeyUsageInCurrentRegion"
func (c *Client) AsymmetricSm2Decrypt(request *AsymmetricSm2DecryptRequest) (response *AsymmetricSm2DecryptResponse, err error) {
    if request == nil {
        request = NewAsymmetricSm2DecryptRequest()
    }
    
    response = NewAsymmetricSm2DecryptResponse()
    err = c.Send(request, response)
    return
}

func NewBindCloudResourceRequest() (request *BindCloudResourceRequest) {
    request = &BindCloudResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "BindCloudResource")
    
    
    return
}

func NewBindCloudResourceResponse() (response *BindCloudResourceResponse) {
    response = &BindCloudResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindCloudResource
// 记录当前key被哪个云产品的那个资源所使用。如果当前key设置了自动过期，则取消该设置，确保当前key不会自动失效。如果当前关联关系已经创建，也返回成功。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CMKUSEDBYCLOUDPRODUCT = "FailedOperation.CmkUsedByCloudProduct"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  RESOURCEUNAVAILABLE_CMKNOTFOUND = "ResourceUnavailable.CmkNotFound"
//  RESOURCEUNAVAILABLE_CMKSTATENOTSUPPORT = "ResourceUnavailable.CmkStateNotSupport"
//  UNSUPPORTEDOPERATION_SERVICETEMPORARYUNAVAILABLE = "UnsupportedOperation.ServiceTemporaryUnavailable"
func (c *Client) BindCloudResource(request *BindCloudResourceRequest) (response *BindCloudResourceResponse, err error) {
    if request == nil {
        request = NewBindCloudResourceRequest()
    }
    
    response = NewBindCloudResourceResponse()
    err = c.Send(request, response)
    return
}

func NewCancelKeyArchiveRequest() (request *CancelKeyArchiveRequest) {
    request = &CancelKeyArchiveRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "CancelKeyArchive")
    
    
    return
}

func NewCancelKeyArchiveResponse() (response *CancelKeyArchiveResponse) {
    response = &CancelKeyArchiveResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CancelKeyArchive
// 取消密钥归档，取消后密钥的状态变为Enabled。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  RESOURCEUNAVAILABLE_CMKNOTFOUND = "ResourceUnavailable.CmkNotFound"
//  RESOURCEUNAVAILABLE_CMKSTATENOTSUPPORT = "ResourceUnavailable.CmkStateNotSupport"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_NOTUSERCREATEDCMK = "UnsupportedOperation.NotUserCreatedCmk"
//  UNSUPPORTEDOPERATION_SERVICETEMPORARYUNAVAILABLE = "UnsupportedOperation.ServiceTemporaryUnavailable"
func (c *Client) CancelKeyArchive(request *CancelKeyArchiveRequest) (response *CancelKeyArchiveResponse, err error) {
    if request == nil {
        request = NewCancelKeyArchiveRequest()
    }
    
    response = NewCancelKeyArchiveResponse()
    err = c.Send(request, response)
    return
}

func NewCancelKeyDeletionRequest() (request *CancelKeyDeletionRequest) {
    request = &CancelKeyDeletionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "CancelKeyDeletion")
    
    
    return
}

func NewCancelKeyDeletionResponse() (response *CancelKeyDeletionResponse) {
    response = &CancelKeyDeletionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CancelKeyDeletion
// 取消CMK的计划删除操作
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  RESOURCEUNAVAILABLE_CMKNOTFOUND = "ResourceUnavailable.CmkNotFound"
//  RESOURCEUNAVAILABLE_CMKNOTPENDINGDELETE = "ResourceUnavailable.CmkNotPendingDelete"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_SERVICETEMPORARYUNAVAILABLE = "UnsupportedOperation.ServiceTemporaryUnavailable"
func (c *Client) CancelKeyDeletion(request *CancelKeyDeletionRequest) (response *CancelKeyDeletionResponse, err error) {
    if request == nil {
        request = NewCancelKeyDeletionRequest()
    }
    
    response = NewCancelKeyDeletionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateKeyRequest() (request *CreateKeyRequest) {
    request = &CreateKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "CreateKey")
    
    
    return
}

func NewCreateKeyResponse() (response *CreateKeyResponse) {
    response = &CreateKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateKey
// 创建用户管理数据密钥的主密钥CMK（Custom Master Key）。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ALIASALREADYEXISTS = "InvalidParameterValue.AliasAlreadyExists"
//  INVALIDPARAMETERVALUE_INVALIDALIAS = "InvalidParameterValue.InvalidAlias"
//  INVALIDPARAMETERVALUE_INVALIDKEYUSAGE = "InvalidParameterValue.InvalidKeyUsage"
//  INVALIDPARAMETERVALUE_INVALIDTYPE = "InvalidParameterValue.InvalidType"
//  INVALIDPARAMETERVALUE_TAGKEYSDUPLICATED = "InvalidParameterValue.TagKeysDuplicated"
//  INVALIDPARAMETERVALUE_TAGSNOTEXISTED = "InvalidParameterValue.TagsNotExisted"
//  LIMITEXCEEDED_CMKLIMITEXCEEDED = "LimitExceeded.CmkLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_SERVICETEMPORARYUNAVAILABLE = "UnsupportedOperation.ServiceTemporaryUnavailable"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDKEYUSAGEINCURRENTREGION = "UnsupportedOperation.UnsupportedKeyUsageInCurrentRegion"
func (c *Client) CreateKey(request *CreateKeyRequest) (response *CreateKeyResponse, err error) {
    if request == nil {
        request = NewCreateKeyRequest()
    }
    
    response = NewCreateKeyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWhiteBoxKeyRequest() (request *CreateWhiteBoxKeyRequest) {
    request = &CreateWhiteBoxKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "CreateWhiteBoxKey")
    
    
    return
}

func NewCreateWhiteBoxKeyResponse() (response *CreateWhiteBoxKeyResponse) {
    response = &CreateWhiteBoxKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateWhiteBoxKey
// 创建白盒密钥。 密钥个数的上限为 50。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ALIASALREADYEXISTS = "InvalidParameterValue.AliasAlreadyExists"
//  INVALIDPARAMETERVALUE_INVALIDALIAS = "InvalidParameterValue.InvalidAlias"
//  INVALIDPARAMETERVALUE_TAGKEYSDUPLICATED = "InvalidParameterValue.TagKeysDuplicated"
//  INVALIDPARAMETERVALUE_TAGSNOTEXISTED = "InvalidParameterValue.TagsNotExisted"
//  LIMITEXCEEDED_KEYLIMITEXCEEDED = "LimitExceeded.KeyLimitExceeded"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateWhiteBoxKey(request *CreateWhiteBoxKeyRequest) (response *CreateWhiteBoxKeyResponse, err error) {
    if request == nil {
        request = NewCreateWhiteBoxKeyRequest()
    }
    
    response = NewCreateWhiteBoxKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDecryptRequest() (request *DecryptRequest) {
    request = &DecryptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "Decrypt")
    
    
    return
}

func NewDecryptResponse() (response *DecryptResponse) {
    response = &DecryptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Decrypt
// 本接口用于解密密文，得到明文数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDCIPHERTEXT = "InvalidParameterValue.InvalidCiphertext"
//  RESOURCEUNAVAILABLE_CMKDISABLED = "ResourceUnavailable.CmkDisabled"
//  RESOURCEUNAVAILABLE_CMKNOTFOUND = "ResourceUnavailable.CmkNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) Decrypt(request *DecryptRequest) (response *DecryptResponse, err error) {
    if request == nil {
        request = NewDecryptRequest()
    }
    
    response = NewDecryptResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteImportedKeyMaterialRequest() (request *DeleteImportedKeyMaterialRequest) {
    request = &DeleteImportedKeyMaterialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "DeleteImportedKeyMaterial")
    
    
    return
}

func NewDeleteImportedKeyMaterialResponse() (response *DeleteImportedKeyMaterialResponse) {
    response = &DeleteImportedKeyMaterialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteImportedKeyMaterial
// 用于删除导入的密钥材料，仅对EXTERNAL类型的CMK有效，该接口将CMK设置为PendingImport 状态，并不会删除CMK，在重新进行密钥导入后可继续使用。彻底删除CMK请使用 ScheduleKeyDeletion 接口。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CMKUSEDBYCLOUDPRODUCT = "FailedOperation.CmkUsedByCloudProduct"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  RESOURCEUNAVAILABLE_CMKNOTFOUND = "ResourceUnavailable.CmkNotFound"
//  RESOURCEUNAVAILABLE_CMKSTATENOTSUPPORT = "ResourceUnavailable.CmkStateNotSupport"
//  UNSUPPORTEDOPERATION_NOTEXTERNALCMK = "UnsupportedOperation.NotExternalCmk"
//  UNSUPPORTEDOPERATION_SERVICETEMPORARYUNAVAILABLE = "UnsupportedOperation.ServiceTemporaryUnavailable"
func (c *Client) DeleteImportedKeyMaterial(request *DeleteImportedKeyMaterialRequest) (response *DeleteImportedKeyMaterialResponse, err error) {
    if request == nil {
        request = NewDeleteImportedKeyMaterialRequest()
    }
    
    response = NewDeleteImportedKeyMaterialResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWhiteBoxKeyRequest() (request *DeleteWhiteBoxKeyRequest) {
    request = &DeleteWhiteBoxKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "DeleteWhiteBoxKey")
    
    
    return
}

func NewDeleteWhiteBoxKeyResponse() (response *DeleteWhiteBoxKeyResponse) {
    response = &DeleteWhiteBoxKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteWhiteBoxKey
// 删除白盒密钥, 注意：必须先禁用后，才可以删除。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteWhiteBoxKey(request *DeleteWhiteBoxKeyRequest) (response *DeleteWhiteBoxKeyResponse, err error) {
    if request == nil {
        request = NewDeleteWhiteBoxKeyRequest()
    }
    
    response = NewDeleteWhiteBoxKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKeyRequest() (request *DescribeKeyRequest) {
    request = &DescribeKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "DescribeKey")
    
    
    return
}

func NewDescribeKeyResponse() (response *DescribeKeyResponse) {
    response = &DescribeKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeKey
// 用于获取指定KeyId的主密钥属性详情信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  RESOURCEUNAVAILABLE_CMKNOTFOUND = "ResourceUnavailable.CmkNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeKey(request *DescribeKeyRequest) (response *DescribeKeyResponse, err error) {
    if request == nil {
        request = NewDescribeKeyRequest()
    }
    
    response = NewDescribeKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKeysRequest() (request *DescribeKeysRequest) {
    request = &DescribeKeysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "DescribeKeys")
    
    
    return
}

func NewDescribeKeysResponse() (response *DescribeKeysResponse) {
    response = &DescribeKeysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeKeys
// 该接口用于批量获取主密钥属性信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DUPLICATEDKEYID = "InvalidParameterValue.DuplicatedKeyId"
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  RESOURCEUNAVAILABLE_CMKNOTFOUND = "ResourceUnavailable.CmkNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeKeys(request *DescribeKeysRequest) (response *DescribeKeysResponse, err error) {
    if request == nil {
        request = NewDescribeKeysRequest()
    }
    
    response = NewDescribeKeysResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWhiteBoxDecryptKeyRequest() (request *DescribeWhiteBoxDecryptKeyRequest) {
    request = &DescribeWhiteBoxDecryptKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "DescribeWhiteBoxDecryptKey")
    
    
    return
}

func NewDescribeWhiteBoxDecryptKeyResponse() (response *DescribeWhiteBoxDecryptKeyResponse) {
    response = &DescribeWhiteBoxDecryptKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWhiteBoxDecryptKey
// 获取白盒解密密钥
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeWhiteBoxDecryptKey(request *DescribeWhiteBoxDecryptKeyRequest) (response *DescribeWhiteBoxDecryptKeyResponse, err error) {
    if request == nil {
        request = NewDescribeWhiteBoxDecryptKeyRequest()
    }
    
    response = NewDescribeWhiteBoxDecryptKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWhiteBoxDeviceFingerprintsRequest() (request *DescribeWhiteBoxDeviceFingerprintsRequest) {
    request = &DescribeWhiteBoxDeviceFingerprintsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "DescribeWhiteBoxDeviceFingerprints")
    
    
    return
}

func NewDescribeWhiteBoxDeviceFingerprintsResponse() (response *DescribeWhiteBoxDeviceFingerprintsResponse) {
    response = &DescribeWhiteBoxDeviceFingerprintsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWhiteBoxDeviceFingerprints
// 获取指定密钥的设备指纹列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeWhiteBoxDeviceFingerprints(request *DescribeWhiteBoxDeviceFingerprintsRequest) (response *DescribeWhiteBoxDeviceFingerprintsResponse, err error) {
    if request == nil {
        request = NewDescribeWhiteBoxDeviceFingerprintsRequest()
    }
    
    response = NewDescribeWhiteBoxDeviceFingerprintsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWhiteBoxKeyRequest() (request *DescribeWhiteBoxKeyRequest) {
    request = &DescribeWhiteBoxKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "DescribeWhiteBoxKey")
    
    
    return
}

func NewDescribeWhiteBoxKeyResponse() (response *DescribeWhiteBoxKeyResponse) {
    response = &DescribeWhiteBoxKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWhiteBoxKey
// 展示白盒密钥的信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeWhiteBoxKey(request *DescribeWhiteBoxKeyRequest) (response *DescribeWhiteBoxKeyResponse, err error) {
    if request == nil {
        request = NewDescribeWhiteBoxKeyRequest()
    }
    
    response = NewDescribeWhiteBoxKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWhiteBoxKeyDetailsRequest() (request *DescribeWhiteBoxKeyDetailsRequest) {
    request = &DescribeWhiteBoxKeyDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "DescribeWhiteBoxKeyDetails")
    
    
    return
}

func NewDescribeWhiteBoxKeyDetailsResponse() (response *DescribeWhiteBoxKeyDetailsResponse) {
    response = &DescribeWhiteBoxKeyDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWhiteBoxKeyDetails
// 获取白盒密钥列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeWhiteBoxKeyDetails(request *DescribeWhiteBoxKeyDetailsRequest) (response *DescribeWhiteBoxKeyDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeWhiteBoxKeyDetailsRequest()
    }
    
    response = NewDescribeWhiteBoxKeyDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWhiteBoxServiceStatusRequest() (request *DescribeWhiteBoxServiceStatusRequest) {
    request = &DescribeWhiteBoxServiceStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "DescribeWhiteBoxServiceStatus")
    
    
    return
}

func NewDescribeWhiteBoxServiceStatusResponse() (response *DescribeWhiteBoxServiceStatusResponse) {
    response = &DescribeWhiteBoxServiceStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWhiteBoxServiceStatus
// 获取白盒密钥服务状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeWhiteBoxServiceStatus(request *DescribeWhiteBoxServiceStatusRequest) (response *DescribeWhiteBoxServiceStatusResponse, err error) {
    if request == nil {
        request = NewDescribeWhiteBoxServiceStatusRequest()
    }
    
    response = NewDescribeWhiteBoxServiceStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDisableKeyRequest() (request *DisableKeyRequest) {
    request = &DisableKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "DisableKey")
    
    
    return
}

func NewDisableKeyResponse() (response *DisableKeyResponse) {
    response = &DisableKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisableKey
// 本接口用于禁用一个主密钥，处于禁用状态的Key无法用于加密、解密操作。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CMKUSEDBYCLOUDPRODUCT = "FailedOperation.CmkUsedByCloudProduct"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  RESOURCEUNAVAILABLE_CMKNOTFOUND = "ResourceUnavailable.CmkNotFound"
//  RESOURCEUNAVAILABLE_CMKSTATENOTSUPPORT = "ResourceUnavailable.CmkStateNotSupport"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_SERVICETEMPORARYUNAVAILABLE = "UnsupportedOperation.ServiceTemporaryUnavailable"
func (c *Client) DisableKey(request *DisableKeyRequest) (response *DisableKeyResponse, err error) {
    if request == nil {
        request = NewDisableKeyRequest()
    }
    
    response = NewDisableKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDisableKeyRotationRequest() (request *DisableKeyRotationRequest) {
    request = &DisableKeyRotationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "DisableKeyRotation")
    
    
    return
}

func NewDisableKeyRotationResponse() (response *DisableKeyRotationResponse) {
    response = &DisableKeyRotationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisableKeyRotation
// 对指定的CMK禁止密钥轮换功能。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  RESOURCEUNAVAILABLE_CMKNOTFOUND = "ResourceUnavailable.CmkNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DisableKeyRotation(request *DisableKeyRotationRequest) (response *DisableKeyRotationResponse, err error) {
    if request == nil {
        request = NewDisableKeyRotationRequest()
    }
    
    response = NewDisableKeyRotationResponse()
    err = c.Send(request, response)
    return
}

func NewDisableKeysRequest() (request *DisableKeysRequest) {
    request = &DisableKeysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "DisableKeys")
    
    
    return
}

func NewDisableKeysResponse() (response *DisableKeysResponse) {
    response = &DisableKeysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisableKeys
// 该接口用于批量禁止CMK的使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CMKUSEDBYCLOUDPRODUCT = "FailedOperation.CmkUsedByCloudProduct"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DUPLICATEDKEYID = "InvalidParameterValue.DuplicatedKeyId"
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  RESOURCEUNAVAILABLE_CMKNOTFOUND = "ResourceUnavailable.CmkNotFound"
//  RESOURCEUNAVAILABLE_CMKSTATENOTSUPPORT = "ResourceUnavailable.CmkStateNotSupport"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_SERVICETEMPORARYUNAVAILABLE = "UnsupportedOperation.ServiceTemporaryUnavailable"
func (c *Client) DisableKeys(request *DisableKeysRequest) (response *DisableKeysResponse, err error) {
    if request == nil {
        request = NewDisableKeysRequest()
    }
    
    response = NewDisableKeysResponse()
    err = c.Send(request, response)
    return
}

func NewDisableWhiteBoxKeyRequest() (request *DisableWhiteBoxKeyRequest) {
    request = &DisableWhiteBoxKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "DisableWhiteBoxKey")
    
    
    return
}

func NewDisableWhiteBoxKeyResponse() (response *DisableWhiteBoxKeyResponse) {
    response = &DisableWhiteBoxKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisableWhiteBoxKey
// 禁用白盒密钥
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DisableWhiteBoxKey(request *DisableWhiteBoxKeyRequest) (response *DisableWhiteBoxKeyResponse, err error) {
    if request == nil {
        request = NewDisableWhiteBoxKeyRequest()
    }
    
    response = NewDisableWhiteBoxKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDisableWhiteBoxKeysRequest() (request *DisableWhiteBoxKeysRequest) {
    request = &DisableWhiteBoxKeysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "DisableWhiteBoxKeys")
    
    
    return
}

func NewDisableWhiteBoxKeysResponse() (response *DisableWhiteBoxKeysResponse) {
    response = &DisableWhiteBoxKeysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisableWhiteBoxKeys
// 批量禁用白盒密钥
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DUPLICATEDKEYID = "InvalidParameterValue.DuplicatedKeyId"
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DisableWhiteBoxKeys(request *DisableWhiteBoxKeysRequest) (response *DisableWhiteBoxKeysResponse, err error) {
    if request == nil {
        request = NewDisableWhiteBoxKeysRequest()
    }
    
    response = NewDisableWhiteBoxKeysResponse()
    err = c.Send(request, response)
    return
}

func NewEnableKeyRequest() (request *EnableKeyRequest) {
    request = &EnableKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "EnableKey")
    
    
    return
}

func NewEnableKeyResponse() (response *EnableKeyResponse) {
    response = &EnableKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EnableKey
// 用于启用一个指定的CMK。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  RESOURCEUNAVAILABLE_CMKNOTFOUND = "ResourceUnavailable.CmkNotFound"
//  RESOURCEUNAVAILABLE_CMKSTATENOTSUPPORT = "ResourceUnavailable.CmkStateNotSupport"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_SERVICETEMPORARYUNAVAILABLE = "UnsupportedOperation.ServiceTemporaryUnavailable"
func (c *Client) EnableKey(request *EnableKeyRequest) (response *EnableKeyResponse, err error) {
    if request == nil {
        request = NewEnableKeyRequest()
    }
    
    response = NewEnableKeyResponse()
    err = c.Send(request, response)
    return
}

func NewEnableKeyRotationRequest() (request *EnableKeyRotationRequest) {
    request = &EnableKeyRotationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "EnableKeyRotation")
    
    
    return
}

func NewEnableKeyRotationResponse() (response *EnableKeyRotationResponse) {
    response = &EnableKeyRotationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EnableKeyRotation
// 对指定的CMK开启密钥轮换功能。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  RESOURCEUNAVAILABLE_CMKNOTFOUND = "ResourceUnavailable.CmkNotFound"
//  RESOURCEUNAVAILABLE_CMKSTATENOTSUPPORT = "ResourceUnavailable.CmkStateNotSupport"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_EXTERNALCMKCANNOTROTATE = "UnsupportedOperation.ExternalCmkCanNotRotate"
//  UNSUPPORTEDOPERATION_SERVICETEMPORARYUNAVAILABLE = "UnsupportedOperation.ServiceTemporaryUnavailable"
func (c *Client) EnableKeyRotation(request *EnableKeyRotationRequest) (response *EnableKeyRotationResponse, err error) {
    if request == nil {
        request = NewEnableKeyRotationRequest()
    }
    
    response = NewEnableKeyRotationResponse()
    err = c.Send(request, response)
    return
}

func NewEnableKeysRequest() (request *EnableKeysRequest) {
    request = &EnableKeysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "EnableKeys")
    
    
    return
}

func NewEnableKeysResponse() (response *EnableKeysResponse) {
    response = &EnableKeysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EnableKeys
// 该接口用于批量启用CMK。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DUPLICATEDKEYID = "InvalidParameterValue.DuplicatedKeyId"
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  RESOURCEUNAVAILABLE_CMKNOTFOUND = "ResourceUnavailable.CmkNotFound"
//  RESOURCEUNAVAILABLE_CMKSTATENOTSUPPORT = "ResourceUnavailable.CmkStateNotSupport"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_SERVICETEMPORARYUNAVAILABLE = "UnsupportedOperation.ServiceTemporaryUnavailable"
func (c *Client) EnableKeys(request *EnableKeysRequest) (response *EnableKeysResponse, err error) {
    if request == nil {
        request = NewEnableKeysRequest()
    }
    
    response = NewEnableKeysResponse()
    err = c.Send(request, response)
    return
}

func NewEnableWhiteBoxKeyRequest() (request *EnableWhiteBoxKeyRequest) {
    request = &EnableWhiteBoxKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "EnableWhiteBoxKey")
    
    
    return
}

func NewEnableWhiteBoxKeyResponse() (response *EnableWhiteBoxKeyResponse) {
    response = &EnableWhiteBoxKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EnableWhiteBoxKey
// 启用白盒密钥
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) EnableWhiteBoxKey(request *EnableWhiteBoxKeyRequest) (response *EnableWhiteBoxKeyResponse, err error) {
    if request == nil {
        request = NewEnableWhiteBoxKeyRequest()
    }
    
    response = NewEnableWhiteBoxKeyResponse()
    err = c.Send(request, response)
    return
}

func NewEnableWhiteBoxKeysRequest() (request *EnableWhiteBoxKeysRequest) {
    request = &EnableWhiteBoxKeysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "EnableWhiteBoxKeys")
    
    
    return
}

func NewEnableWhiteBoxKeysResponse() (response *EnableWhiteBoxKeysResponse) {
    response = &EnableWhiteBoxKeysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EnableWhiteBoxKeys
// 批量启用白盒密钥
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DUPLICATEDKEYID = "InvalidParameterValue.DuplicatedKeyId"
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) EnableWhiteBoxKeys(request *EnableWhiteBoxKeysRequest) (response *EnableWhiteBoxKeysResponse, err error) {
    if request == nil {
        request = NewEnableWhiteBoxKeysRequest()
    }
    
    response = NewEnableWhiteBoxKeysResponse()
    err = c.Send(request, response)
    return
}

func NewEncryptRequest() (request *EncryptRequest) {
    request = &EncryptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "Encrypt")
    
    
    return
}

func NewEncryptResponse() (response *EncryptResponse) {
    response = &EncryptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Encrypt
// 本接口用于加密最多为4KB任意数据，可用于加密数据库密码，RSA Key，或其它较小的敏感信息。对于应用的数据加密，使用GenerateDataKey生成的DataKey进行本地数据的加解密操作
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  INVALIDPARAMETERVALUE_INVALIDPLAINTEXT = "InvalidParameterValue.InvalidPlaintext"
//  RESOURCEUNAVAILABLE_CMKARCHIVED = "ResourceUnavailable.CmkArchived"
//  RESOURCEUNAVAILABLE_CMKDISABLED = "ResourceUnavailable.CmkDisabled"
//  RESOURCEUNAVAILABLE_CMKNOTFOUND = "ResourceUnavailable.CmkNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) Encrypt(request *EncryptRequest) (response *EncryptResponse, err error) {
    if request == nil {
        request = NewEncryptRequest()
    }
    
    response = NewEncryptResponse()
    err = c.Send(request, response)
    return
}

func NewEncryptByWhiteBoxRequest() (request *EncryptByWhiteBoxRequest) {
    request = &EncryptByWhiteBoxRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "EncryptByWhiteBox")
    
    
    return
}

func NewEncryptByWhiteBoxResponse() (response *EncryptByWhiteBoxResponse) {
    response = &EncryptByWhiteBoxResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EncryptByWhiteBox
// 使用白盒密钥进行加密
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_KEYDISABLED = "ResourceUnavailable.KeyDisabled"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) EncryptByWhiteBox(request *EncryptByWhiteBoxRequest) (response *EncryptByWhiteBoxResponse, err error) {
    if request == nil {
        request = NewEncryptByWhiteBoxRequest()
    }
    
    response = NewEncryptByWhiteBoxResponse()
    err = c.Send(request, response)
    return
}

func NewGenerateDataKeyRequest() (request *GenerateDataKeyRequest) {
    request = &GenerateDataKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "GenerateDataKey")
    
    
    return
}

func NewGenerateDataKeyResponse() (response *GenerateDataKeyResponse) {
    response = &GenerateDataKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GenerateDataKey
// 本接口生成一个数据密钥，您可以用这个密钥进行本地数据的加密。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  RESOURCEUNAVAILABLE_CMKDISABLED = "ResourceUnavailable.CmkDisabled"
//  RESOURCEUNAVAILABLE_CMKNOTFOUND = "ResourceUnavailable.CmkNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) GenerateDataKey(request *GenerateDataKeyRequest) (response *GenerateDataKeyResponse, err error) {
    if request == nil {
        request = NewGenerateDataKeyRequest()
    }
    
    response = NewGenerateDataKeyResponse()
    err = c.Send(request, response)
    return
}

func NewGenerateRandomRequest() (request *GenerateRandomRequest) {
    request = &GenerateRandomRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "GenerateRandom")
    
    
    return
}

func NewGenerateRandomResponse() (response *GenerateRandomResponse) {
    response = &GenerateRandomResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GenerateRandom
// 随机数生成接口。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) GenerateRandom(request *GenerateRandomRequest) (response *GenerateRandomResponse, err error) {
    if request == nil {
        request = NewGenerateRandomRequest()
    }
    
    response = NewGenerateRandomResponse()
    err = c.Send(request, response)
    return
}

func NewGetKeyRotationStatusRequest() (request *GetKeyRotationStatusRequest) {
    request = &GetKeyRotationStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "GetKeyRotationStatus")
    
    
    return
}

func NewGetKeyRotationStatusResponse() (response *GetKeyRotationStatusResponse) {
    response = &GetKeyRotationStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetKeyRotationStatus
// 查询指定的CMK是否开启了密钥轮换功能。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  RESOURCEUNAVAILABLE_CMKNOTFOUND = "ResourceUnavailable.CmkNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) GetKeyRotationStatus(request *GetKeyRotationStatusRequest) (response *GetKeyRotationStatusResponse, err error) {
    if request == nil {
        request = NewGetKeyRotationStatusRequest()
    }
    
    response = NewGetKeyRotationStatusResponse()
    err = c.Send(request, response)
    return
}

func NewGetParametersForImportRequest() (request *GetParametersForImportRequest) {
    request = &GetParametersForImportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "GetParametersForImport")
    
    
    return
}

func NewGetParametersForImportResponse() (response *GetParametersForImportResponse) {
    response = &GetParametersForImportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetParametersForImport
// 获取导入主密钥（CMK）材料的参数，返回的Token作为执行ImportKeyMaterial的参数之一，返回的PublicKey用于对自主导入密钥材料进行加密。返回的Token和PublicKey 24小时后失效，失效后如需重新导入，需要再次调用该接口获取新的Token和PublicKey。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  RESOURCEUNAVAILABLE_CMKNOTFOUND = "ResourceUnavailable.CmkNotFound"
//  UNSUPPORTEDOPERATION_NOTEXTERNALCMK = "UnsupportedOperation.NotExternalCmk"
//  UNSUPPORTEDOPERATION_SERVICETEMPORARYUNAVAILABLE = "UnsupportedOperation.ServiceTemporaryUnavailable"
func (c *Client) GetParametersForImport(request *GetParametersForImportRequest) (response *GetParametersForImportResponse, err error) {
    if request == nil {
        request = NewGetParametersForImportRequest()
    }
    
    response = NewGetParametersForImportResponse()
    err = c.Send(request, response)
    return
}

func NewGetPublicKeyRequest() (request *GetPublicKeyRequest) {
    request = &GetPublicKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "GetPublicKey")
    
    
    return
}

func NewGetPublicKeyResponse() (response *GetPublicKeyResponse) {
    response = &GetPublicKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetPublicKey
// 该接口用于获取非对称密钥的公钥信息，可用于本地数据加密或验签。只有处于Enabled状态的非对称密钥才可能获取公钥。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  RESOURCEUNAVAILABLE_CMKNOTFOUND = "ResourceUnavailable.CmkNotFound"
//  RESOURCEUNAVAILABLE_CMKSTATENOTSUPPORT = "ResourceUnavailable.CmkStateNotSupport"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) GetPublicKey(request *GetPublicKeyRequest) (response *GetPublicKeyResponse, err error) {
    if request == nil {
        request = NewGetPublicKeyRequest()
    }
    
    response = NewGetPublicKeyResponse()
    err = c.Send(request, response)
    return
}

func NewGetRegionsRequest() (request *GetRegionsRequest) {
    request = &GetRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "GetRegions")
    
    
    return
}

func NewGetRegionsResponse() (response *GetRegionsResponse) {
    response = &GetRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetRegions
// 获取支持的地域列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) GetRegions(request *GetRegionsRequest) (response *GetRegionsResponse, err error) {
    if request == nil {
        request = NewGetRegionsRequest()
    }
    
    response = NewGetRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewGetServiceStatusRequest() (request *GetServiceStatusRequest) {
    request = &GetServiceStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "GetServiceStatus")
    
    
    return
}

func NewGetServiceStatusResponse() (response *GetServiceStatusResponse) {
    response = &GetServiceStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetServiceStatus
// 用于查询该用户是否已开通KMS服务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) GetServiceStatus(request *GetServiceStatusRequest) (response *GetServiceStatusResponse, err error) {
    if request == nil {
        request = NewGetServiceStatusRequest()
    }
    
    response = NewGetServiceStatusResponse()
    err = c.Send(request, response)
    return
}

func NewImportKeyMaterialRequest() (request *ImportKeyMaterialRequest) {
    request = &ImportKeyMaterialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "ImportKeyMaterial")
    
    
    return
}

func NewImportKeyMaterialResponse() (response *ImportKeyMaterialResponse) {
    response = &ImportKeyMaterialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ImportKeyMaterial
// 用于导入密钥材料。只有类型为EXTERNAL 的CMK 才可以导入，导入的密钥材料使用 GetParametersForImport 获取的密钥进行加密。可以为指定的 CMK 重新导入密钥材料，并重新指定过期时间，但必须导入相同的密钥材料。CMK 密钥材料导入后不可以更换密钥材料。导入的密钥材料过期或者被删除后，指定的CMK将无法使用，需要再次导入相同的密钥材料才能正常使用。CMK是独立的，同样的密钥材料可导入不同的 CMK 中，但使用其中一个 CMK 加密的数据无法使用另一个 CMK解密。
//
// 只有Enabled 和 PendingImport状态的CMK可以导入密钥材料。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DECRYPTMATERIALERROR = "InvalidParameter.DecryptMaterialError"
//  INVALIDPARAMETERVALUE_MATERIALNOTMATCH = "InvalidParameterValue.MaterialNotMatch"
//  RESOURCEUNAVAILABLE_CMKNOTFOUND = "ResourceUnavailable.CmkNotFound"
//  RESOURCEUNAVAILABLE_CMKSTATENOTSUPPORT = "ResourceUnavailable.CmkStateNotSupport"
//  RESOURCEUNAVAILABLE_TOKENEXPIRED = "ResourceUnavailable.TokenExpired"
//  UNSUPPORTEDOPERATION_NOTEXTERNALCMK = "UnsupportedOperation.NotExternalCmk"
//  UNSUPPORTEDOPERATION_SERVICETEMPORARYUNAVAILABLE = "UnsupportedOperation.ServiceTemporaryUnavailable"
func (c *Client) ImportKeyMaterial(request *ImportKeyMaterialRequest) (response *ImportKeyMaterialResponse, err error) {
    if request == nil {
        request = NewImportKeyMaterialRequest()
    }
    
    response = NewImportKeyMaterialResponse()
    err = c.Send(request, response)
    return
}

func NewListAlgorithmsRequest() (request *ListAlgorithmsRequest) {
    request = &ListAlgorithmsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "ListAlgorithms")
    
    
    return
}

func NewListAlgorithmsResponse() (response *ListAlgorithmsResponse) {
    response = &ListAlgorithmsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListAlgorithms
// 列出当前Region支持的加密方式
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ListAlgorithms(request *ListAlgorithmsRequest) (response *ListAlgorithmsResponse, err error) {
    if request == nil {
        request = NewListAlgorithmsRequest()
    }
    
    response = NewListAlgorithmsResponse()
    err = c.Send(request, response)
    return
}

func NewListKeyDetailRequest() (request *ListKeyDetailRequest) {
    request = &ListKeyDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "ListKeyDetail")
    
    
    return
}

func NewListKeyDetailResponse() (response *ListKeyDetailResponse) {
    response = &ListKeyDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListKeyDetail
// 根据指定Offset和Limit获取主密钥列表详情。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ListKeyDetail(request *ListKeyDetailRequest) (response *ListKeyDetailResponse, err error) {
    if request == nil {
        request = NewListKeyDetailRequest()
    }
    
    response = NewListKeyDetailResponse()
    err = c.Send(request, response)
    return
}

func NewListKeysRequest() (request *ListKeysRequest) {
    request = &ListKeysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "ListKeys")
    
    
    return
}

func NewListKeysResponse() (response *ListKeysResponse) {
    response = &ListKeysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListKeys
// 列出账号下面状态为Enabled， Disabled 和 PendingImport 的CMK KeyId 列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ListKeys(request *ListKeysRequest) (response *ListKeysResponse, err error) {
    if request == nil {
        request = NewListKeysRequest()
    }
    
    response = NewListKeysResponse()
    err = c.Send(request, response)
    return
}

func NewOverwriteWhiteBoxDeviceFingerprintsRequest() (request *OverwriteWhiteBoxDeviceFingerprintsRequest) {
    request = &OverwriteWhiteBoxDeviceFingerprintsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "OverwriteWhiteBoxDeviceFingerprints")
    
    
    return
}

func NewOverwriteWhiteBoxDeviceFingerprintsResponse() (response *OverwriteWhiteBoxDeviceFingerprintsResponse) {
    response = &OverwriteWhiteBoxDeviceFingerprintsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// OverwriteWhiteBoxDeviceFingerprints
// 覆盖指定密钥的设备指纹信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  LIMITEXCEEDED_FINGERPRINTSLIMITEXCEEDED = "LimitExceeded.FingerprintsLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) OverwriteWhiteBoxDeviceFingerprints(request *OverwriteWhiteBoxDeviceFingerprintsRequest) (response *OverwriteWhiteBoxDeviceFingerprintsResponse, err error) {
    if request == nil {
        request = NewOverwriteWhiteBoxDeviceFingerprintsRequest()
    }
    
    response = NewOverwriteWhiteBoxDeviceFingerprintsResponse()
    err = c.Send(request, response)
    return
}

func NewReEncryptRequest() (request *ReEncryptRequest) {
    request = &ReEncryptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "ReEncrypt")
    
    
    return
}

func NewReEncryptResponse() (response *ReEncryptResponse) {
    response = &ReEncryptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReEncrypt
// 使用指定CMK对密文重新加密。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDCIPHERTEXT = "InvalidParameterValue.InvalidCiphertext"
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  RESOURCEUNAVAILABLE_CMKDISABLED = "ResourceUnavailable.CmkDisabled"
//  RESOURCEUNAVAILABLE_CMKNOTFOUND = "ResourceUnavailable.CmkNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ReEncrypt(request *ReEncryptRequest) (response *ReEncryptResponse, err error) {
    if request == nil {
        request = NewReEncryptRequest()
    }
    
    response = NewReEncryptResponse()
    err = c.Send(request, response)
    return
}

func NewScheduleKeyDeletionRequest() (request *ScheduleKeyDeletionRequest) {
    request = &ScheduleKeyDeletionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "ScheduleKeyDeletion")
    
    
    return
}

func NewScheduleKeyDeletionResponse() (response *ScheduleKeyDeletionResponse) {
    response = &ScheduleKeyDeletionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ScheduleKeyDeletion
// CMK计划删除接口，用于指定CMK删除的时间，可选时间区间为[7,30]天
//
// 可能返回的错误码:
//  FAILEDOPERATION_CMKUSEDBYCLOUDPRODUCT = "FailedOperation.CmkUsedByCloudProduct"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPENDINGWINDOWINDAYS = "InvalidParameter.InvalidPendingWindowInDays"
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  RESOURCEUNAVAILABLE_CMKNOTFOUND = "ResourceUnavailable.CmkNotFound"
//  RESOURCEUNAVAILABLE_CMKSHOULDBEDISABLED = "ResourceUnavailable.CmkShouldBeDisabled"
//  RESOURCEUNAVAILABLE_CMKSTATENOTSUPPORT = "ResourceUnavailable.CmkStateNotSupport"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_SERVICETEMPORARYUNAVAILABLE = "UnsupportedOperation.ServiceTemporaryUnavailable"
func (c *Client) ScheduleKeyDeletion(request *ScheduleKeyDeletionRequest) (response *ScheduleKeyDeletionResponse, err error) {
    if request == nil {
        request = NewScheduleKeyDeletionRequest()
    }
    
    response = NewScheduleKeyDeletionResponse()
    err = c.Send(request, response)
    return
}

func NewSignByAsymmetricKeyRequest() (request *SignByAsymmetricKeyRequest) {
    request = &SignByAsymmetricKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "SignByAsymmetricKey")
    
    
    return
}

func NewSignByAsymmetricKeyResponse() (response *SignByAsymmetricKeyResponse) {
    response = &SignByAsymmetricKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SignByAsymmetricKey
// 非对称密钥签名。
//
// 注意：只有 KeyUsage 为 ASYMMETRIC_SIGN_VERIFY_SM2、ASYMMETRIC_SIGN_VERIFY_ECC 或其他支持的 ASYMMETRIC_SIGN_VERIFY_${ALGORITHM} 的密钥才可以使用签名功能。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEUNAVAILABLE_CMKNOTFOUND = "ResourceUnavailable.CmkNotFound"
//  RESOURCEUNAVAILABLE_CMKSTATENOTSUPPORT = "ResourceUnavailable.CmkStateNotSupport"
func (c *Client) SignByAsymmetricKey(request *SignByAsymmetricKeyRequest) (response *SignByAsymmetricKeyResponse, err error) {
    if request == nil {
        request = NewSignByAsymmetricKeyRequest()
    }
    
    response = NewSignByAsymmetricKeyResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindCloudResourceRequest() (request *UnbindCloudResourceRequest) {
    request = &UnbindCloudResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "UnbindCloudResource")
    
    
    return
}

func NewUnbindCloudResourceResponse() (response *UnbindCloudResourceResponse) {
    response = &UnbindCloudResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnbindCloudResource
// 删除指定（key, 资源，云产品）的记录，以表明：指定的云产品的资源已不再使用当前的key。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  RESOURCEUNAVAILABLE_CLOUDRESOURCEBINDINGNOTFOUND = "ResourceUnavailable.CloudResourceBindingNotFound"
//  RESOURCEUNAVAILABLE_CMKNOTFOUND = "ResourceUnavailable.CmkNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_SERVICETEMPORARYUNAVAILABLE = "UnsupportedOperation.ServiceTemporaryUnavailable"
func (c *Client) UnbindCloudResource(request *UnbindCloudResourceRequest) (response *UnbindCloudResourceResponse, err error) {
    if request == nil {
        request = NewUnbindCloudResourceRequest()
    }
    
    response = NewUnbindCloudResourceResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAliasRequest() (request *UpdateAliasRequest) {
    request = &UpdateAliasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "UpdateAlias")
    
    
    return
}

func NewUpdateAliasResponse() (response *UpdateAliasResponse) {
    response = &UpdateAliasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateAlias
// 用于修改CMK的别名。对于处于PendingDelete状态的CMK禁止修改。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ALIASALREADYEXISTS = "InvalidParameterValue.AliasAlreadyExists"
//  INVALIDPARAMETERVALUE_INVALIDALIAS = "InvalidParameterValue.InvalidAlias"
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  RESOURCEUNAVAILABLE_CMKNOTFOUND = "ResourceUnavailable.CmkNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_SERVICETEMPORARYUNAVAILABLE = "UnsupportedOperation.ServiceTemporaryUnavailable"
func (c *Client) UpdateAlias(request *UpdateAliasRequest) (response *UpdateAliasResponse, err error) {
    if request == nil {
        request = NewUpdateAliasRequest()
    }
    
    response = NewUpdateAliasResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateKeyDescriptionRequest() (request *UpdateKeyDescriptionRequest) {
    request = &UpdateKeyDescriptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "UpdateKeyDescription")
    
    
    return
}

func NewUpdateKeyDescriptionResponse() (response *UpdateKeyDescriptionResponse) {
    response = &UpdateKeyDescriptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateKeyDescription
// 该接口用于对指定的cmk修改描述信息。对于处于PendingDelete状态的CMK禁止修改。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  RESOURCEUNAVAILABLE_CMKNOTFOUND = "ResourceUnavailable.CmkNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_SERVICETEMPORARYUNAVAILABLE = "UnsupportedOperation.ServiceTemporaryUnavailable"
func (c *Client) UpdateKeyDescription(request *UpdateKeyDescriptionRequest) (response *UpdateKeyDescriptionResponse, err error) {
    if request == nil {
        request = NewUpdateKeyDescriptionRequest()
    }
    
    response = NewUpdateKeyDescriptionResponse()
    err = c.Send(request, response)
    return
}

func NewVerifyByAsymmetricKeyRequest() (request *VerifyByAsymmetricKeyRequest) {
    request = &VerifyByAsymmetricKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("kms", APIVersion, "VerifyByAsymmetricKey")
    
    
    return
}

func NewVerifyByAsymmetricKeyResponse() (response *VerifyByAsymmetricKeyResponse) {
    response = &VerifyByAsymmetricKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// VerifyByAsymmetricKey
// 使用非对称密钥验签
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  RESOURCEUNAVAILABLE_CMKNOTFOUND = "ResourceUnavailable.CmkNotFound"
//  RESOURCEUNAVAILABLE_CMKSTATENOTSUPPORT = "ResourceUnavailable.CmkStateNotSupport"
func (c *Client) VerifyByAsymmetricKey(request *VerifyByAsymmetricKeyRequest) (response *VerifyByAsymmetricKeyResponse, err error) {
    if request == nil {
        request = NewVerifyByAsymmetricKeyRequest()
    }
    
    response = NewVerifyByAsymmetricKeyResponse()
    err = c.Send(request, response)
    return
}
