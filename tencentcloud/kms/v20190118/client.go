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
    "context"
    "errors"
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
    return c.ArchiveKeyWithContext(context.Background(), request)
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
func (c *Client) ArchiveKeyWithContext(ctx context.Context, request *ArchiveKeyRequest) (response *ArchiveKeyResponse, err error) {
    if request == nil {
        request = NewArchiveKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ArchiveKey require credential")
    }

    request.SetContext(ctx)
    
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
    return c.AsymmetricRsaDecryptWithContext(context.Background(), request)
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
func (c *Client) AsymmetricRsaDecryptWithContext(ctx context.Context, request *AsymmetricRsaDecryptRequest) (response *AsymmetricRsaDecryptResponse, err error) {
    if request == nil {
        request = NewAsymmetricRsaDecryptRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AsymmetricRsaDecrypt require credential")
    }

    request.SetContext(ctx)
    
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
    return c.AsymmetricSm2DecryptWithContext(context.Background(), request)
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
func (c *Client) AsymmetricSm2DecryptWithContext(ctx context.Context, request *AsymmetricSm2DecryptRequest) (response *AsymmetricSm2DecryptResponse, err error) {
    if request == nil {
        request = NewAsymmetricSm2DecryptRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AsymmetricSm2Decrypt require credential")
    }

    request.SetContext(ctx)
    
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
    return c.BindCloudResourceWithContext(context.Background(), request)
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
func (c *Client) BindCloudResourceWithContext(ctx context.Context, request *BindCloudResourceRequest) (response *BindCloudResourceResponse, err error) {
    if request == nil {
        request = NewBindCloudResourceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindCloudResource require credential")
    }

    request.SetContext(ctx)
    
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
    return c.CancelKeyArchiveWithContext(context.Background(), request)
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
func (c *Client) CancelKeyArchiveWithContext(ctx context.Context, request *CancelKeyArchiveRequest) (response *CancelKeyArchiveResponse, err error) {
    if request == nil {
        request = NewCancelKeyArchiveRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelKeyArchive require credential")
    }

    request.SetContext(ctx)
    
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
    return c.CancelKeyDeletionWithContext(context.Background(), request)
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
func (c *Client) CancelKeyDeletionWithContext(ctx context.Context, request *CancelKeyDeletionRequest) (response *CancelKeyDeletionResponse, err error) {
    if request == nil {
        request = NewCancelKeyDeletionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelKeyDeletion require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION_TAGGINGERROR = "FailedOperation.TaggingError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ALIASALREADYEXISTS = "InvalidParameterValue.AliasAlreadyExists"
//  INVALIDPARAMETERVALUE_INVALIDALIAS = "InvalidParameterValue.InvalidAlias"
//  INVALIDPARAMETERVALUE_INVALIDHSMCLUSTERID = "InvalidParameterValue.InvalidHsmClusterId"
//  INVALIDPARAMETERVALUE_INVALIDKEYUSAGE = "InvalidParameterValue.InvalidKeyUsage"
//  INVALIDPARAMETERVALUE_INVALIDTYPE = "InvalidParameterValue.InvalidType"
//  INVALIDPARAMETERVALUE_TAGKEYSDUPLICATED = "InvalidParameterValue.TagKeysDuplicated"
//  INVALIDPARAMETERVALUE_TAGSNOTEXISTED = "InvalidParameterValue.TagsNotExisted"
//  LIMITEXCEEDED_CMKLIMITEXCEEDED = "LimitExceeded.CmkLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_SERVICETEMPORARYUNAVAILABLE = "UnsupportedOperation.ServiceTemporaryUnavailable"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDKEYUSAGEINCURRENTREGION = "UnsupportedOperation.UnsupportedKeyUsageInCurrentRegion"
func (c *Client) CreateKey(request *CreateKeyRequest) (response *CreateKeyResponse, err error) {
    return c.CreateKeyWithContext(context.Background(), request)
}

// CreateKey
// 创建用户管理数据密钥的主密钥CMK（Custom Master Key）。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TAGGINGERROR = "FailedOperation.TaggingError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ALIASALREADYEXISTS = "InvalidParameterValue.AliasAlreadyExists"
//  INVALIDPARAMETERVALUE_INVALIDALIAS = "InvalidParameterValue.InvalidAlias"
//  INVALIDPARAMETERVALUE_INVALIDHSMCLUSTERID = "InvalidParameterValue.InvalidHsmClusterId"
//  INVALIDPARAMETERVALUE_INVALIDKEYUSAGE = "InvalidParameterValue.InvalidKeyUsage"
//  INVALIDPARAMETERVALUE_INVALIDTYPE = "InvalidParameterValue.InvalidType"
//  INVALIDPARAMETERVALUE_TAGKEYSDUPLICATED = "InvalidParameterValue.TagKeysDuplicated"
//  INVALIDPARAMETERVALUE_TAGSNOTEXISTED = "InvalidParameterValue.TagsNotExisted"
//  LIMITEXCEEDED_CMKLIMITEXCEEDED = "LimitExceeded.CmkLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_SERVICETEMPORARYUNAVAILABLE = "UnsupportedOperation.ServiceTemporaryUnavailable"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDKEYUSAGEINCURRENTREGION = "UnsupportedOperation.UnsupportedKeyUsageInCurrentRegion"
func (c *Client) CreateKeyWithContext(ctx context.Context, request *CreateKeyRequest) (response *CreateKeyResponse, err error) {
    if request == nil {
        request = NewCreateKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateKey require credential")
    }

    request.SetContext(ctx)
    
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
    return c.CreateWhiteBoxKeyWithContext(context.Background(), request)
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
func (c *Client) CreateWhiteBoxKeyWithContext(ctx context.Context, request *CreateWhiteBoxKeyRequest) (response *CreateWhiteBoxKeyResponse, err error) {
    if request == nil {
        request = NewCreateWhiteBoxKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWhiteBoxKey require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION_ENCRYPTIONERROR = "FailedOperation.EncryptionError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDCIPHERTEXT = "InvalidParameterValue.InvalidCiphertext"
//  RESOURCEUNAVAILABLE_CMKDISABLED = "ResourceUnavailable.CmkDisabled"
//  RESOURCEUNAVAILABLE_CMKNOTFOUND = "ResourceUnavailable.CmkNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) Decrypt(request *DecryptRequest) (response *DecryptResponse, err error) {
    return c.DecryptWithContext(context.Background(), request)
}

// Decrypt
// 本接口用于解密密文，得到明文数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ENCRYPTIONERROR = "FailedOperation.EncryptionError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDCIPHERTEXT = "InvalidParameterValue.InvalidCiphertext"
//  RESOURCEUNAVAILABLE_CMKDISABLED = "ResourceUnavailable.CmkDisabled"
//  RESOURCEUNAVAILABLE_CMKNOTFOUND = "ResourceUnavailable.CmkNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DecryptWithContext(ctx context.Context, request *DecryptRequest) (response *DecryptResponse, err error) {
    if request == nil {
        request = NewDecryptRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("Decrypt require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeleteImportedKeyMaterialWithContext(context.Background(), request)
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
func (c *Client) DeleteImportedKeyMaterialWithContext(ctx context.Context, request *DeleteImportedKeyMaterialRequest) (response *DeleteImportedKeyMaterialResponse, err error) {
    if request == nil {
        request = NewDeleteImportedKeyMaterialRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteImportedKeyMaterial require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeleteWhiteBoxKeyWithContext(context.Background(), request)
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
func (c *Client) DeleteWhiteBoxKeyWithContext(ctx context.Context, request *DeleteWhiteBoxKeyRequest) (response *DeleteWhiteBoxKeyResponse, err error) {
    if request == nil {
        request = NewDeleteWhiteBoxKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWhiteBoxKey require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeKeyWithContext(context.Background(), request)
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
func (c *Client) DescribeKeyWithContext(ctx context.Context, request *DescribeKeyRequest) (response *DescribeKeyResponse, err error) {
    if request == nil {
        request = NewDescribeKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKey require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeKeysWithContext(context.Background(), request)
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
func (c *Client) DescribeKeysWithContext(ctx context.Context, request *DescribeKeysRequest) (response *DescribeKeysResponse, err error) {
    if request == nil {
        request = NewDescribeKeysRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKeys require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeWhiteBoxDecryptKeyWithContext(context.Background(), request)
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
func (c *Client) DescribeWhiteBoxDecryptKeyWithContext(ctx context.Context, request *DescribeWhiteBoxDecryptKeyRequest) (response *DescribeWhiteBoxDecryptKeyResponse, err error) {
    if request == nil {
        request = NewDescribeWhiteBoxDecryptKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWhiteBoxDecryptKey require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeWhiteBoxDeviceFingerprints(request *DescribeWhiteBoxDeviceFingerprintsRequest) (response *DescribeWhiteBoxDeviceFingerprintsResponse, err error) {
    return c.DescribeWhiteBoxDeviceFingerprintsWithContext(context.Background(), request)
}

// DescribeWhiteBoxDeviceFingerprints
// 获取指定密钥的设备指纹列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeWhiteBoxDeviceFingerprintsWithContext(ctx context.Context, request *DescribeWhiteBoxDeviceFingerprintsRequest) (response *DescribeWhiteBoxDeviceFingerprintsResponse, err error) {
    if request == nil {
        request = NewDescribeWhiteBoxDeviceFingerprintsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWhiteBoxDeviceFingerprints require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeWhiteBoxKeyWithContext(context.Background(), request)
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
func (c *Client) DescribeWhiteBoxKeyWithContext(ctx context.Context, request *DescribeWhiteBoxKeyRequest) (response *DescribeWhiteBoxKeyResponse, err error) {
    if request == nil {
        request = NewDescribeWhiteBoxKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWhiteBoxKey require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeWhiteBoxKeyDetailsWithContext(context.Background(), request)
}

// DescribeWhiteBoxKeyDetails
// 获取白盒密钥列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeWhiteBoxKeyDetailsWithContext(ctx context.Context, request *DescribeWhiteBoxKeyDetailsRequest) (response *DescribeWhiteBoxKeyDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeWhiteBoxKeyDetailsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWhiteBoxKeyDetails require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeWhiteBoxServiceStatusWithContext(context.Background(), request)
}

// DescribeWhiteBoxServiceStatus
// 获取白盒密钥服务状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  RESOURCEUNAVAILABLE_NOTPURCHASED = "ResourceUnavailable.NotPurchased"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeWhiteBoxServiceStatusWithContext(ctx context.Context, request *DescribeWhiteBoxServiceStatusRequest) (response *DescribeWhiteBoxServiceStatusResponse, err error) {
    if request == nil {
        request = NewDescribeWhiteBoxServiceStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWhiteBoxServiceStatus require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DisableKeyWithContext(context.Background(), request)
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
func (c *Client) DisableKeyWithContext(ctx context.Context, request *DisableKeyRequest) (response *DisableKeyResponse, err error) {
    if request == nil {
        request = NewDisableKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableKey require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DisableKeyRotationWithContext(context.Background(), request)
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
func (c *Client) DisableKeyRotationWithContext(ctx context.Context, request *DisableKeyRotationRequest) (response *DisableKeyRotationResponse, err error) {
    if request == nil {
        request = NewDisableKeyRotationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableKeyRotation require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DisableKeysWithContext(context.Background(), request)
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
func (c *Client) DisableKeysWithContext(ctx context.Context, request *DisableKeysRequest) (response *DisableKeysResponse, err error) {
    if request == nil {
        request = NewDisableKeysRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableKeys require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DisableWhiteBoxKeyWithContext(context.Background(), request)
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
func (c *Client) DisableWhiteBoxKeyWithContext(ctx context.Context, request *DisableWhiteBoxKeyRequest) (response *DisableWhiteBoxKeyResponse, err error) {
    if request == nil {
        request = NewDisableWhiteBoxKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableWhiteBoxKey require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DisableWhiteBoxKeysWithContext(context.Background(), request)
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
func (c *Client) DisableWhiteBoxKeysWithContext(ctx context.Context, request *DisableWhiteBoxKeysRequest) (response *DisableWhiteBoxKeysResponse, err error) {
    if request == nil {
        request = NewDisableWhiteBoxKeysRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableWhiteBoxKeys require credential")
    }

    request.SetContext(ctx)
    
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
    return c.EnableKeyWithContext(context.Background(), request)
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
func (c *Client) EnableKeyWithContext(ctx context.Context, request *EnableKeyRequest) (response *EnableKeyResponse, err error) {
    if request == nil {
        request = NewEnableKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableKey require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  RESOURCEUNAVAILABLE_CMKNOTFOUND = "ResourceUnavailable.CmkNotFound"
//  RESOURCEUNAVAILABLE_CMKSTATENOTSUPPORT = "ResourceUnavailable.CmkStateNotSupport"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_EXTERNALCMKCANNOTROTATE = "UnsupportedOperation.ExternalCmkCanNotRotate"
//  UNSUPPORTEDOPERATION_SERVICETEMPORARYUNAVAILABLE = "UnsupportedOperation.ServiceTemporaryUnavailable"
func (c *Client) EnableKeyRotation(request *EnableKeyRotationRequest) (response *EnableKeyRotationResponse, err error) {
    return c.EnableKeyRotationWithContext(context.Background(), request)
}

// EnableKeyRotation
// 对指定的CMK开启密钥轮换功能。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  RESOURCEUNAVAILABLE_CMKNOTFOUND = "ResourceUnavailable.CmkNotFound"
//  RESOURCEUNAVAILABLE_CMKSTATENOTSUPPORT = "ResourceUnavailable.CmkStateNotSupport"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_EXTERNALCMKCANNOTROTATE = "UnsupportedOperation.ExternalCmkCanNotRotate"
//  UNSUPPORTEDOPERATION_SERVICETEMPORARYUNAVAILABLE = "UnsupportedOperation.ServiceTemporaryUnavailable"
func (c *Client) EnableKeyRotationWithContext(ctx context.Context, request *EnableKeyRotationRequest) (response *EnableKeyRotationResponse, err error) {
    if request == nil {
        request = NewEnableKeyRotationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableKeyRotation require credential")
    }

    request.SetContext(ctx)
    
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
    return c.EnableKeysWithContext(context.Background(), request)
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
func (c *Client) EnableKeysWithContext(ctx context.Context, request *EnableKeysRequest) (response *EnableKeysResponse, err error) {
    if request == nil {
        request = NewEnableKeysRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableKeys require credential")
    }

    request.SetContext(ctx)
    
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
    return c.EnableWhiteBoxKeyWithContext(context.Background(), request)
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
func (c *Client) EnableWhiteBoxKeyWithContext(ctx context.Context, request *EnableWhiteBoxKeyRequest) (response *EnableWhiteBoxKeyResponse, err error) {
    if request == nil {
        request = NewEnableWhiteBoxKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableWhiteBoxKey require credential")
    }

    request.SetContext(ctx)
    
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
    return c.EnableWhiteBoxKeysWithContext(context.Background(), request)
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
func (c *Client) EnableWhiteBoxKeysWithContext(ctx context.Context, request *EnableWhiteBoxKeysRequest) (response *EnableWhiteBoxKeysResponse, err error) {
    if request == nil {
        request = NewEnableWhiteBoxKeysRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableWhiteBoxKeys require credential")
    }

    request.SetContext(ctx)
    
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
    return c.EncryptWithContext(context.Background(), request)
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
func (c *Client) EncryptWithContext(ctx context.Context, request *EncryptRequest) (response *EncryptResponse, err error) {
    if request == nil {
        request = NewEncryptRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("Encrypt require credential")
    }

    request.SetContext(ctx)
    
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
    return c.EncryptByWhiteBoxWithContext(context.Background(), request)
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
func (c *Client) EncryptByWhiteBoxWithContext(ctx context.Context, request *EncryptByWhiteBoxRequest) (response *EncryptByWhiteBoxResponse, err error) {
    if request == nil {
        request = NewEncryptByWhiteBoxRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EncryptByWhiteBox require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION_ENCRYPTIONERROR = "FailedOperation.EncryptionError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  RESOURCEUNAVAILABLE_CMKDISABLED = "ResourceUnavailable.CmkDisabled"
//  RESOURCEUNAVAILABLE_CMKNOTFOUND = "ResourceUnavailable.CmkNotFound"
//  RESOURCEUNAVAILABLE_KEYPENDINGDELETE = "ResourceUnavailable.KeyPendingDelete"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) GenerateDataKey(request *GenerateDataKeyRequest) (response *GenerateDataKeyResponse, err error) {
    return c.GenerateDataKeyWithContext(context.Background(), request)
}

// GenerateDataKey
// 本接口生成一个数据密钥，您可以用这个密钥进行本地数据的加密。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ENCRYPTIONERROR = "FailedOperation.EncryptionError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  RESOURCEUNAVAILABLE_CMKDISABLED = "ResourceUnavailable.CmkDisabled"
//  RESOURCEUNAVAILABLE_CMKNOTFOUND = "ResourceUnavailable.CmkNotFound"
//  RESOURCEUNAVAILABLE_KEYPENDINGDELETE = "ResourceUnavailable.KeyPendingDelete"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) GenerateDataKeyWithContext(ctx context.Context, request *GenerateDataKeyRequest) (response *GenerateDataKeyResponse, err error) {
    if request == nil {
        request = NewGenerateDataKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GenerateDataKey require credential")
    }

    request.SetContext(ctx)
    
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
    return c.GenerateRandomWithContext(context.Background(), request)
}

// GenerateRandom
// 随机数生成接口。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) GenerateRandomWithContext(ctx context.Context, request *GenerateRandomRequest) (response *GenerateRandomResponse, err error) {
    if request == nil {
        request = NewGenerateRandomRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GenerateRandom require credential")
    }

    request.SetContext(ctx)
    
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
    return c.GetKeyRotationStatusWithContext(context.Background(), request)
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
func (c *Client) GetKeyRotationStatusWithContext(ctx context.Context, request *GetKeyRotationStatusRequest) (response *GetKeyRotationStatusResponse, err error) {
    if request == nil {
        request = NewGetKeyRotationStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetKeyRotationStatus require credential")
    }

    request.SetContext(ctx)
    
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
    return c.GetParametersForImportWithContext(context.Background(), request)
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
func (c *Client) GetParametersForImportWithContext(ctx context.Context, request *GetParametersForImportRequest) (response *GetParametersForImportResponse, err error) {
    if request == nil {
        request = NewGetParametersForImportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetParametersForImport require credential")
    }

    request.SetContext(ctx)
    
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
    return c.GetPublicKeyWithContext(context.Background(), request)
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
func (c *Client) GetPublicKeyWithContext(ctx context.Context, request *GetPublicKeyRequest) (response *GetPublicKeyResponse, err error) {
    if request == nil {
        request = NewGetPublicKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetPublicKey require credential")
    }

    request.SetContext(ctx)
    
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
// 获取可以提供KMS服务的地域列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) GetRegions(request *GetRegionsRequest) (response *GetRegionsResponse, err error) {
    return c.GetRegionsWithContext(context.Background(), request)
}

// GetRegions
// 获取可以提供KMS服务的地域列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) GetRegionsWithContext(ctx context.Context, request *GetRegionsRequest) (response *GetRegionsResponse, err error) {
    if request == nil {
        request = NewGetRegionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetRegions require credential")
    }

    request.SetContext(ctx)
    
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetServiceStatus(request *GetServiceStatusRequest) (response *GetServiceStatusResponse, err error) {
    return c.GetServiceStatusWithContext(context.Background(), request)
}

// GetServiceStatus
// 用于查询该用户是否已开通KMS服务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetServiceStatusWithContext(ctx context.Context, request *GetServiceStatusRequest) (response *GetServiceStatusResponse, err error) {
    if request == nil {
        request = NewGetServiceStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetServiceStatus require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  INVALIDPARAMETERVALUE_MATERIALNOTMATCH = "InvalidParameterValue.MaterialNotMatch"
//  RESOURCEUNAVAILABLE_CMKNOTFOUND = "ResourceUnavailable.CmkNotFound"
//  RESOURCEUNAVAILABLE_CMKSTATENOTSUPPORT = "ResourceUnavailable.CmkStateNotSupport"
//  RESOURCEUNAVAILABLE_TOKENEXPIRED = "ResourceUnavailable.TokenExpired"
//  UNSUPPORTEDOPERATION_NOTEXTERNALCMK = "UnsupportedOperation.NotExternalCmk"
//  UNSUPPORTEDOPERATION_SERVICETEMPORARYUNAVAILABLE = "UnsupportedOperation.ServiceTemporaryUnavailable"
func (c *Client) ImportKeyMaterial(request *ImportKeyMaterialRequest) (response *ImportKeyMaterialResponse, err error) {
    return c.ImportKeyMaterialWithContext(context.Background(), request)
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
//  INVALIDPARAMETERVALUE_INVALIDKEYID = "InvalidParameterValue.InvalidKeyId"
//  INVALIDPARAMETERVALUE_MATERIALNOTMATCH = "InvalidParameterValue.MaterialNotMatch"
//  RESOURCEUNAVAILABLE_CMKNOTFOUND = "ResourceUnavailable.CmkNotFound"
//  RESOURCEUNAVAILABLE_CMKSTATENOTSUPPORT = "ResourceUnavailable.CmkStateNotSupport"
//  RESOURCEUNAVAILABLE_TOKENEXPIRED = "ResourceUnavailable.TokenExpired"
//  UNSUPPORTEDOPERATION_NOTEXTERNALCMK = "UnsupportedOperation.NotExternalCmk"
//  UNSUPPORTEDOPERATION_SERVICETEMPORARYUNAVAILABLE = "UnsupportedOperation.ServiceTemporaryUnavailable"
func (c *Client) ImportKeyMaterialWithContext(ctx context.Context, request *ImportKeyMaterialRequest) (response *ImportKeyMaterialResponse, err error) {
    if request == nil {
        request = NewImportKeyMaterialRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImportKeyMaterial require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ListAlgorithmsWithContext(context.Background(), request)
}

// ListAlgorithms
// 列出当前Region支持的加密方式
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ListAlgorithmsWithContext(ctx context.Context, request *ListAlgorithmsRequest) (response *ListAlgorithmsResponse, err error) {
    if request == nil {
        request = NewListAlgorithmsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAlgorithms require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_INVALIDHSMCLUSTERID = "InvalidParameterValue.InvalidHsmClusterId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ListKeyDetail(request *ListKeyDetailRequest) (response *ListKeyDetailResponse, err error) {
    return c.ListKeyDetailWithContext(context.Background(), request)
}

// ListKeyDetail
// 根据指定Offset和Limit获取主密钥列表详情。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDHSMCLUSTERID = "InvalidParameterValue.InvalidHsmClusterId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ListKeyDetailWithContext(ctx context.Context, request *ListKeyDetailRequest) (response *ListKeyDetailResponse, err error) {
    if request == nil {
        request = NewListKeyDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListKeyDetail require credential")
    }

    request.SetContext(ctx)
    
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
//  INVALIDPARAMETERVALUE_INVALIDHSMCLUSTERID = "InvalidParameterValue.InvalidHsmClusterId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ListKeys(request *ListKeysRequest) (response *ListKeysResponse, err error) {
    return c.ListKeysWithContext(context.Background(), request)
}

// ListKeys
// 列出账号下面状态为Enabled， Disabled 和 PendingImport 的CMK KeyId 列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDHSMCLUSTERID = "InvalidParameterValue.InvalidHsmClusterId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ListKeysWithContext(ctx context.Context, request *ListKeysRequest) (response *ListKeysResponse, err error) {
    if request == nil {
        request = NewListKeysRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListKeys require credential")
    }

    request.SetContext(ctx)
    
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
    return c.OverwriteWhiteBoxDeviceFingerprintsWithContext(context.Background(), request)
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
func (c *Client) OverwriteWhiteBoxDeviceFingerprintsWithContext(ctx context.Context, request *OverwriteWhiteBoxDeviceFingerprintsRequest) (response *OverwriteWhiteBoxDeviceFingerprintsResponse, err error) {
    if request == nil {
        request = NewOverwriteWhiteBoxDeviceFingerprintsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("OverwriteWhiteBoxDeviceFingerprints require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ReEncryptWithContext(context.Background(), request)
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
func (c *Client) ReEncryptWithContext(ctx context.Context, request *ReEncryptRequest) (response *ReEncryptResponse, err error) {
    if request == nil {
        request = NewReEncryptRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReEncrypt require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ScheduleKeyDeletionWithContext(context.Background(), request)
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
func (c *Client) ScheduleKeyDeletionWithContext(ctx context.Context, request *ScheduleKeyDeletionRequest) (response *ScheduleKeyDeletionResponse, err error) {
    if request == nil {
        request = NewScheduleKeyDeletionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ScheduleKeyDeletion require credential")
    }

    request.SetContext(ctx)
    
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
    return c.SignByAsymmetricKeyWithContext(context.Background(), request)
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
func (c *Client) SignByAsymmetricKeyWithContext(ctx context.Context, request *SignByAsymmetricKeyRequest) (response *SignByAsymmetricKeyResponse, err error) {
    if request == nil {
        request = NewSignByAsymmetricKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SignByAsymmetricKey require credential")
    }

    request.SetContext(ctx)
    
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
    return c.UnbindCloudResourceWithContext(context.Background(), request)
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
func (c *Client) UnbindCloudResourceWithContext(ctx context.Context, request *UnbindCloudResourceRequest) (response *UnbindCloudResourceResponse, err error) {
    if request == nil {
        request = NewUnbindCloudResourceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindCloudResource require credential")
    }

    request.SetContext(ctx)
    
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
    return c.UpdateAliasWithContext(context.Background(), request)
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
func (c *Client) UpdateAliasWithContext(ctx context.Context, request *UpdateAliasRequest) (response *UpdateAliasResponse, err error) {
    if request == nil {
        request = NewUpdateAliasRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAlias require credential")
    }

    request.SetContext(ctx)
    
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
    return c.UpdateKeyDescriptionWithContext(context.Background(), request)
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
func (c *Client) UpdateKeyDescriptionWithContext(ctx context.Context, request *UpdateKeyDescriptionRequest) (response *UpdateKeyDescriptionResponse, err error) {
    if request == nil {
        request = NewUpdateKeyDescriptionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateKeyDescription require credential")
    }

    request.SetContext(ctx)
    
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
    return c.VerifyByAsymmetricKeyWithContext(context.Background(), request)
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
func (c *Client) VerifyByAsymmetricKeyWithContext(ctx context.Context, request *VerifyByAsymmetricKeyRequest) (response *VerifyByAsymmetricKeyResponse, err error) {
    if request == nil {
        request = NewVerifyByAsymmetricKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("VerifyByAsymmetricKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewVerifyByAsymmetricKeyResponse()
    err = c.Send(request, response)
    return
}
