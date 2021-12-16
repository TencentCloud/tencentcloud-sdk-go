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

package v20201222

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-12-22"

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


func NewArchiveFlowRequest() (request *ArchiveFlowRequest) {
    request = &ArchiveFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "ArchiveFlow")
    
    
    return
}

func NewArchiveFlowResponse() (response *ArchiveFlowResponse) {
    response = &ArchiveFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ArchiveFlow
// 此接口（ArchiveFlow）用于流程的归档。
//
// 
//
// 注意：归档后的流程不可再进行发送、签署、拒签、撤回等一系列操作。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ArchiveFlow(request *ArchiveFlowRequest) (response *ArchiveFlowResponse, err error) {
    if request == nil {
        request = NewArchiveFlowRequest()
    }
    
    response = NewArchiveFlowResponse()
    err = c.Send(request, response)
    return
}

func NewCancelFlowRequest() (request *CancelFlowRequest) {
    request = &CancelFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "CancelFlow")
    
    
    return
}

func NewCancelFlowResponse() (response *CancelFlowResponse) {
    response = &CancelFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CancelFlow
// 此接口（CancelFlow）用于撤销正在进行中的流程。
//
// 
//
// 注：已归档流程不可完成撤销动作。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CancelFlow(request *CancelFlowRequest) (response *CancelFlowResponse, err error) {
    if request == nil {
        request = NewCancelFlowRequest()
    }
    
    response = NewCancelFlowResponse()
    err = c.Send(request, response)
    return
}

func NewCheckBankCard2EVerificationRequest() (request *CheckBankCard2EVerificationRequest) {
    request = &CheckBankCard2EVerificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "CheckBankCard2EVerification")
    
    
    return
}

func NewCheckBankCard2EVerificationResponse() (response *CheckBankCard2EVerificationResponse) {
    response = &CheckBankCard2EVerificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckBankCard2EVerification
// 该接口为第三方平台向电子签平台验证银行卡二要素
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CheckBankCard2EVerification(request *CheckBankCard2EVerificationRequest) (response *CheckBankCard2EVerificationResponse, err error) {
    if request == nil {
        request = NewCheckBankCard2EVerificationRequest()
    }
    
    response = NewCheckBankCard2EVerificationResponse()
    err = c.Send(request, response)
    return
}

func NewCheckBankCard3EVerificationRequest() (request *CheckBankCard3EVerificationRequest) {
    request = &CheckBankCard3EVerificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "CheckBankCard3EVerification")
    
    
    return
}

func NewCheckBankCard3EVerificationResponse() (response *CheckBankCard3EVerificationResponse) {
    response = &CheckBankCard3EVerificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckBankCard3EVerification
// 该接口为第三方平台向电子签平台验证银行卡三要素
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CheckBankCard3EVerification(request *CheckBankCard3EVerificationRequest) (response *CheckBankCard3EVerificationResponse, err error) {
    if request == nil {
        request = NewCheckBankCard3EVerificationRequest()
    }
    
    response = NewCheckBankCard3EVerificationResponse()
    err = c.Send(request, response)
    return
}

func NewCheckBankCard4EVerificationRequest() (request *CheckBankCard4EVerificationRequest) {
    request = &CheckBankCard4EVerificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "CheckBankCard4EVerification")
    
    
    return
}

func NewCheckBankCard4EVerificationResponse() (response *CheckBankCard4EVerificationResponse) {
    response = &CheckBankCard4EVerificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckBankCard4EVerification
// 该接口为第三方平台向电子签平台验证银行卡四要素
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CheckBankCard4EVerification(request *CheckBankCard4EVerificationRequest) (response *CheckBankCard4EVerificationResponse, err error) {
    if request == nil {
        request = NewCheckBankCard4EVerificationRequest()
    }
    
    response = NewCheckBankCard4EVerificationResponse()
    err = c.Send(request, response)
    return
}

func NewCheckBankCardVerificationRequest() (request *CheckBankCardVerificationRequest) {
    request = &CheckBankCardVerificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "CheckBankCardVerification")
    
    
    return
}

func NewCheckBankCardVerificationResponse() (response *CheckBankCardVerificationResponse) {
    response = &CheckBankCardVerificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckBankCardVerification
// 该接口为第三方平台向电子签平台验证银行卡二/三/四要素
//
// 银行卡二要素(同CheckBankCard2EVerification): bank_card + name
//
// 银行卡三要素(同CheckBankCard3EVerification): bank_card + name + id_card_number
//
// 银行卡四要素(同CheckBankCard4EVerification): bank_card + name + id_card_number + mobile
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CheckBankCardVerification(request *CheckBankCardVerificationRequest) (response *CheckBankCardVerificationResponse, err error) {
    if request == nil {
        request = NewCheckBankCardVerificationRequest()
    }
    
    response = NewCheckBankCardVerificationResponse()
    err = c.Send(request, response)
    return
}

func NewCheckFaceIdentifyRequest() (request *CheckFaceIdentifyRequest) {
    request = &CheckFaceIdentifyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "CheckFaceIdentify")
    
    
    return
}

func NewCheckFaceIdentifyResponse() (response *CheckFaceIdentifyResponse) {
    response = &CheckFaceIdentifyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckFaceIdentify
// 该接口为第三方平台向电子签平台检测慧眼或腾讯电子签小程序人脸核身结果
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CheckFaceIdentify(request *CheckFaceIdentifyRequest) (response *CheckFaceIdentifyResponse, err error) {
    if request == nil {
        request = NewCheckFaceIdentifyRequest()
    }
    
    response = NewCheckFaceIdentifyResponse()
    err = c.Send(request, response)
    return
}

func NewCheckIdCardVerificationRequest() (request *CheckIdCardVerificationRequest) {
    request = &CheckIdCardVerificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "CheckIdCardVerification")
    
    
    return
}

func NewCheckIdCardVerificationResponse() (response *CheckIdCardVerificationResponse) {
    response = &CheckIdCardVerificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckIdCardVerification
// 该接口为第三方平台向电子签平台验证姓名和身份证信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CheckIdCardVerification(request *CheckIdCardVerificationRequest) (response *CheckIdCardVerificationResponse, err error) {
    if request == nil {
        request = NewCheckIdCardVerificationRequest()
    }
    
    response = NewCheckIdCardVerificationResponse()
    err = c.Send(request, response)
    return
}

func NewCheckMobileAndNameRequest() (request *CheckMobileAndNameRequest) {
    request = &CheckMobileAndNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "CheckMobileAndName")
    
    
    return
}

func NewCheckMobileAndNameResponse() (response *CheckMobileAndNameResponse) {
    response = &CheckMobileAndNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckMobileAndName
// 该接口为第三方平台向电子签平台验证手机号二要素
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CheckMobileAndName(request *CheckMobileAndNameRequest) (response *CheckMobileAndNameResponse, err error) {
    if request == nil {
        request = NewCheckMobileAndNameRequest()
    }
    
    response = NewCheckMobileAndNameResponse()
    err = c.Send(request, response)
    return
}

func NewCheckMobileVerificationRequest() (request *CheckMobileVerificationRequest) {
    request = &CheckMobileVerificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "CheckMobileVerification")
    
    
    return
}

func NewCheckMobileVerificationResponse() (response *CheckMobileVerificationResponse) {
    response = &CheckMobileVerificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckMobileVerification
// 该接口为第三方平台向电子签平台验证手机号三要素
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CheckMobileVerification(request *CheckMobileVerificationRequest) (response *CheckMobileVerificationResponse, err error) {
    if request == nil {
        request = NewCheckMobileVerificationRequest()
    }
    
    response = NewCheckMobileVerificationResponse()
    err = c.Send(request, response)
    return
}

func NewCheckVerifyCodeMatchFlowIdRequest() (request *CheckVerifyCodeMatchFlowIdRequest) {
    request = &CheckVerifyCodeMatchFlowIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "CheckVerifyCodeMatchFlowId")
    
    
    return
}

func NewCheckVerifyCodeMatchFlowIdResponse() (response *CheckVerifyCodeMatchFlowIdResponse) {
    response = &CheckVerifyCodeMatchFlowIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckVerifyCodeMatchFlowId
// 此接口用于确认验证码是否正确
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CACHE = "InternalError.Cache"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_SERIALIZE = "InternalError.Serialize"
//  INTERNALERROR_UNSERIALIZE = "InternalError.UnSerialize"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CheckVerifyCodeMatchFlowId(request *CheckVerifyCodeMatchFlowIdRequest) (response *CheckVerifyCodeMatchFlowIdResponse, err error) {
    if request == nil {
        request = NewCheckVerifyCodeMatchFlowIdRequest()
    }
    
    response = NewCheckVerifyCodeMatchFlowIdResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFaceIdSignRequest() (request *CreateFaceIdSignRequest) {
    request = &CreateFaceIdSignRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "CreateFaceIdSign")
    
    
    return
}

func NewCreateFaceIdSignResponse() (response *CreateFaceIdSignResponse) {
    response = &CreateFaceIdSignResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateFaceIdSign
// 该接口为第三方平台向电子签平台获取慧眼慧眼API签名
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateFaceIdSign(request *CreateFaceIdSignRequest) (response *CreateFaceIdSignResponse, err error) {
    if request == nil {
        request = NewCreateFaceIdSignRequest()
    }
    
    response = NewCreateFaceIdSignResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFlowByFilesRequest() (request *CreateFlowByFilesRequest) {
    request = &CreateFlowByFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "CreateFlowByFiles")
    
    
    return
}

func NewCreateFlowByFilesResponse() (response *CreateFlowByFilesResponse) {
    response = &CreateFlowByFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateFlowByFiles
// 此接口（CreateFlowByFiles）用于通过PDF文件创建签署流程。
//
// 
//
// 注意：调用此接口前，请先调用多文件上传接口 (UploadFiles)，提前上传合同文件。
//
// 可能返回的错误码:
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateFlowByFiles(request *CreateFlowByFilesRequest) (response *CreateFlowByFilesResponse, err error) {
    if request == nil {
        request = NewCreateFlowByFilesRequest()
    }
    
    response = NewCreateFlowByFilesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateH5FaceIdUrlRequest() (request *CreateH5FaceIdUrlRequest) {
    request = &CreateH5FaceIdUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "CreateH5FaceIdUrl")
    
    
    return
}

func NewCreateH5FaceIdUrlResponse() (response *CreateH5FaceIdUrlResponse) {
    response = &CreateH5FaceIdUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateH5FaceIdUrl
// 该接口为第三方平台向电子签平台获取慧眼H5人脸核身Url
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateH5FaceIdUrl(request *CreateH5FaceIdUrlRequest) (response *CreateH5FaceIdUrlResponse, err error) {
    if request == nil {
        request = NewCreateH5FaceIdUrlRequest()
    }
    
    response = NewCreateH5FaceIdUrlResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePreviewSignUrlRequest() (request *CreatePreviewSignUrlRequest) {
    request = &CreatePreviewSignUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "CreatePreviewSignUrl")
    
    
    return
}

func NewCreatePreviewSignUrlResponse() (response *CreatePreviewSignUrlResponse) {
    response = &CreatePreviewSignUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePreviewSignUrl
// 此接口（CreatePreviewSignUrl）用于生成生成预览签署URL。
//
// 
//
// 注：调用此接口前，请确保您已提前调用了发送流程接口（SendFlow）指定相关签署方。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTALREADYEXISTS = "FailedOperation.AccountAlreadyExists"
//  FAILEDOPERATION_ACCOUNTVERIFYFAIL = "FailedOperation.AccountVerifyFail"
//  FAILEDOPERATION_AUTHFAIL = "FailedOperation.AuthFail"
//  FAILEDOPERATION_DLOCKFAILED = "FailedOperation.DLockFailed"
//  FAILEDOPERATION_ERRBINDINGREPEATED = "FailedOperation.ErrBindingRepeated"
//  FAILEDOPERATION_GENERATEUSERSEAL = "FailedOperation.GenerateUserSeal"
//  FAILEDOPERATION_NOROLE = "FailedOperation.NoRole"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_CACHE = "InternalError.Cache"
//  INTERNALERROR_CALLBACK = "InternalError.Callback"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_GENERATEID = "InternalError.GenerateId"
//  INTERNALERROR_MQ = "InternalError.Mq"
//  INTERNALERROR_PDF = "InternalError.Pdf"
//  INTERNALERROR_SERIALIZE = "InternalError.Serialize"
//  INTERNALERROR_STORAGE = "InternalError.Storage"
//  INTERNALERROR_THIRDPARTY = "InternalError.ThirdParty"
//  INTERNALERROR_UNSERIALIZE = "InternalError.UnSerialize"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"
//  OPERATIONDENIED_NOAPIAUTH = "OperationDenied.NoApiAuth"
//  OPERATIONDENIED_NOEXTRAVERIFY = "OperationDenied.NoExtraVerify"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOPAYMENTVERIFY = "OperationDenied.NoPaymentVerify"
//  OPERATIONDENIED_NOSESSION = "OperationDenied.NoSession"
//  OPERATIONDENIED_NOVERIFY = "OperationDenied.NoVerify"
//  OPERATIONDENIED_NOTOWNERUIN = "OperationDenied.NotOwnerUin"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreatePreviewSignUrl(request *CreatePreviewSignUrlRequest) (response *CreatePreviewSignUrlResponse, err error) {
    if request == nil {
        request = NewCreatePreviewSignUrlRequest()
    }
    
    response = NewCreatePreviewSignUrlResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSealRequest() (request *CreateSealRequest) {
    request = &CreateSealRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "CreateSeal")
    
    
    return
}

func NewCreateSealResponse() (response *CreateSealResponse) {
    response = &CreateSealResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSeal
// 此接口（CreateSeal）用于创建个人/企业印章。
//
// 
//
// 注意：使用FileId参数指定印章，需先调用多文件上传 (UploadFiles) 上传印章图片。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_GENERATEID = "InternalError.GenerateId"
//  INTERNALERROR_STORAGE = "InternalError.Storage"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateSeal(request *CreateSealRequest) (response *CreateSealResponse, err error) {
    if request == nil {
        request = NewCreateSealRequest()
    }
    
    response = NewCreateSealResponse()
    err = c.Send(request, response)
    return
}

func NewCreateServerFlowSignRequest() (request *CreateServerFlowSignRequest) {
    request = &CreateServerFlowSignRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "CreateServerFlowSign")
    
    
    return
}

func NewCreateServerFlowSignResponse() (response *CreateServerFlowSignResponse) {
    response = &CreateServerFlowSignResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateServerFlowSign
// 此接口（CreateServerFlowSign）用于静默签署文件。
//
// 
//
// 注：
//
// 1、此接口为白名单接口，调用前请提前与客服经理或邮件至e-contract@tencent.com进行联系。
//
// 2、仅合同发起者可使用流程静默签署能力。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTVERIFYFAIL = "FailedOperation.AccountVerifyFail"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_PDF = "InternalError.Pdf"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateServerFlowSign(request *CreateServerFlowSignRequest) (response *CreateServerFlowSignResponse, err error) {
    if request == nil {
        request = NewCreateServerFlowSignRequest()
    }
    
    response = NewCreateServerFlowSignResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSignUrlRequest() (request *CreateSignUrlRequest) {
    request = &CreateSignUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "CreateSignUrl")
    
    
    return
}

func NewCreateSignUrlResponse() (response *CreateSignUrlResponse) {
    response = &CreateSignUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSignUrl
// 此接口（CreateSignUrl）用于生成指定用户的签署URL。
//
// 
//
// 注：调用此接口前，请确保您已提前调用了发送流程接口（SendFlow）指定相关签署方。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTALREADYEXISTS = "FailedOperation.AccountAlreadyExists"
//  FAILEDOPERATION_ACCOUNTVERIFYFAIL = "FailedOperation.AccountVerifyFail"
//  FAILEDOPERATION_AUTHFAIL = "FailedOperation.AuthFail"
//  FAILEDOPERATION_DLOCKFAILED = "FailedOperation.DLockFailed"
//  FAILEDOPERATION_ERRBINDINGREPEATED = "FailedOperation.ErrBindingRepeated"
//  FAILEDOPERATION_GENERATEUSERSEAL = "FailedOperation.GenerateUserSeal"
//  FAILEDOPERATION_NOROLE = "FailedOperation.NoRole"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_CACHE = "InternalError.Cache"
//  INTERNALERROR_CALLBACK = "InternalError.Callback"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_GENERATEID = "InternalError.GenerateId"
//  INTERNALERROR_MQ = "InternalError.Mq"
//  INTERNALERROR_PDF = "InternalError.Pdf"
//  INTERNALERROR_SERIALIZE = "InternalError.Serialize"
//  INTERNALERROR_STORAGE = "InternalError.Storage"
//  INTERNALERROR_THIRDPARTY = "InternalError.ThirdParty"
//  INTERNALERROR_UNSERIALIZE = "InternalError.UnSerialize"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"
//  OPERATIONDENIED_NOAPIAUTH = "OperationDenied.NoApiAuth"
//  OPERATIONDENIED_NOEXTRAVERIFY = "OperationDenied.NoExtraVerify"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOPAYMENTVERIFY = "OperationDenied.NoPaymentVerify"
//  OPERATIONDENIED_NOSESSION = "OperationDenied.NoSession"
//  OPERATIONDENIED_NOVERIFY = "OperationDenied.NoVerify"
//  OPERATIONDENIED_NOTOWNERUIN = "OperationDenied.NotOwnerUin"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSignUrl(request *CreateSignUrlRequest) (response *CreateSignUrlResponse, err error) {
    if request == nil {
        request = NewCreateSignUrlRequest()
    }
    
    response = NewCreateSignUrlResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSubOrganizationRequest() (request *CreateSubOrganizationRequest) {
    request = &CreateSubOrganizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "CreateSubOrganization")
    
    
    return
}

func NewCreateSubOrganizationResponse() (response *CreateSubOrganizationResponse) {
    response = &CreateSubOrganizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSubOrganization
// 此接口（CreateSubOrganization）用于在腾讯电子签内注册子机构。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTALREADYEXISTS = "FailedOperation.AccountAlreadyExists"
//  FAILEDOPERATION_ACCOUNTVERIFYFAIL = "FailedOperation.AccountVerifyFail"
//  FAILEDOPERATION_AUTHFAIL = "FailedOperation.AuthFail"
//  FAILEDOPERATION_DLOCKFAILED = "FailedOperation.DLockFailed"
//  FAILEDOPERATION_ERRBINDINGREPEATED = "FailedOperation.ErrBindingRepeated"
//  FAILEDOPERATION_GENERATEUSERSEAL = "FailedOperation.GenerateUserSeal"
//  FAILEDOPERATION_NOROLE = "FailedOperation.NoRole"
//  FAILEDOPERATION_OPENIDALREADYEXISTS = "FailedOperation.OpenIdAlreadyExists"
//  FAILEDOPERATION_ORGIDCARDNUMBERALREADYEXISTS = "FailedOperation.OrgIdCardNumberAlreadyExists"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_CACHE = "InternalError.Cache"
//  INTERNALERROR_CALLBACK = "InternalError.Callback"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_GENERATEID = "InternalError.GenerateId"
//  INTERNALERROR_MQ = "InternalError.Mq"
//  INTERNALERROR_PDF = "InternalError.Pdf"
//  INTERNALERROR_SERIALIZE = "InternalError.Serialize"
//  INTERNALERROR_STORAGE = "InternalError.Storage"
//  INTERNALERROR_THIRDPARTY = "InternalError.ThirdParty"
//  INTERNALERROR_UNSERIALIZE = "InternalError.UnSerialize"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"
//  OPERATIONDENIED_NOAPIAUTH = "OperationDenied.NoApiAuth"
//  OPERATIONDENIED_NOEXTRAVERIFY = "OperationDenied.NoExtraVerify"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOPAYMENTVERIFY = "OperationDenied.NoPaymentVerify"
//  OPERATIONDENIED_NOSESSION = "OperationDenied.NoSession"
//  OPERATIONDENIED_NOVERIFY = "OperationDenied.NoVerify"
//  OPERATIONDENIED_NOTOWNERUIN = "OperationDenied.NotOwnerUin"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSubOrganization(request *CreateSubOrganizationRequest) (response *CreateSubOrganizationResponse, err error) {
    if request == nil {
        request = NewCreateSubOrganizationRequest()
    }
    
    response = NewCreateSubOrganizationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSubOrganizationAndSealRequest() (request *CreateSubOrganizationAndSealRequest) {
    request = &CreateSubOrganizationAndSealRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "CreateSubOrganizationAndSeal")
    
    
    return
}

func NewCreateSubOrganizationAndSealResponse() (response *CreateSubOrganizationAndSealResponse) {
    response = &CreateSubOrganizationAndSealResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSubOrganizationAndSeal
// 此接口（CreateSubOrganizationAndSeal）用于注册子机构，同时系统将为该子企业自动生成一个默认电子印章图片。
//
// 
//
// 注意：
//
// 1. 在后续的签署流程中，若未指定签署使用的印章ID，则默认调用自动生成的印章图片进行签署。
//
// 2. 此接口为白名单接口，如您需要使用此能力，请提前与客户经理沟通或邮件至e-contract@tencent.com与我们联系。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTALREADYEXISTS = "FailedOperation.AccountAlreadyExists"
//  FAILEDOPERATION_ACCOUNTVERIFYFAIL = "FailedOperation.AccountVerifyFail"
//  FAILEDOPERATION_AUTHFAIL = "FailedOperation.AuthFail"
//  FAILEDOPERATION_DLOCKFAILED = "FailedOperation.DLockFailed"
//  FAILEDOPERATION_ERRBINDINGREPEATED = "FailedOperation.ErrBindingRepeated"
//  FAILEDOPERATION_GENERATEUSERSEAL = "FailedOperation.GenerateUserSeal"
//  FAILEDOPERATION_NOROLE = "FailedOperation.NoRole"
//  FAILEDOPERATION_ORGIDCARDNUMBERALREADYEXISTS = "FailedOperation.OrgIdCardNumberAlreadyExists"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_CACHE = "InternalError.Cache"
//  INTERNALERROR_CALLBACK = "InternalError.Callback"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_GENERATEID = "InternalError.GenerateId"
//  INTERNALERROR_MQ = "InternalError.Mq"
//  INTERNALERROR_PDF = "InternalError.Pdf"
//  INTERNALERROR_SERIALIZE = "InternalError.Serialize"
//  INTERNALERROR_STORAGE = "InternalError.Storage"
//  INTERNALERROR_THIRDPARTY = "InternalError.ThirdParty"
//  INTERNALERROR_UNSERIALIZE = "InternalError.UnSerialize"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"
//  OPERATIONDENIED_NOAPIAUTH = "OperationDenied.NoApiAuth"
//  OPERATIONDENIED_NOEXTRAVERIFY = "OperationDenied.NoExtraVerify"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOPAYMENTVERIFY = "OperationDenied.NoPaymentVerify"
//  OPERATIONDENIED_NOSESSION = "OperationDenied.NoSession"
//  OPERATIONDENIED_NOVERIFY = "OperationDenied.NoVerify"
//  OPERATIONDENIED_NOTOWNERUIN = "OperationDenied.NotOwnerUin"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSubOrganizationAndSeal(request *CreateSubOrganizationAndSealRequest) (response *CreateSubOrganizationAndSealResponse, err error) {
    if request == nil {
        request = NewCreateSubOrganizationAndSealRequest()
    }
    
    response = NewCreateSubOrganizationAndSealResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserRequest() (request *CreateUserRequest) {
    request = &CreateUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "CreateUser")
    
    
    return
}

func NewCreateUserResponse() (response *CreateUserResponse) {
    response = &CreateUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateUser
// 此接口（CreateUser）用于注册腾讯电子签个人用户。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTALREADYEXISTS = "FailedOperation.AccountAlreadyExists"
//  FAILEDOPERATION_ACCOUNTVERIFYFAIL = "FailedOperation.AccountVerifyFail"
//  FAILEDOPERATION_AUTHFAIL = "FailedOperation.AuthFail"
//  FAILEDOPERATION_DLOCKFAILED = "FailedOperation.DLockFailed"
//  FAILEDOPERATION_ERRBINDINGREPEATED = "FailedOperation.ErrBindingRepeated"
//  FAILEDOPERATION_GENERATEUSERSEAL = "FailedOperation.GenerateUserSeal"
//  FAILEDOPERATION_NOROLE = "FailedOperation.NoRole"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED1D = "FailedOperation.RequestLimitExceeded1D"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED1H = "FailedOperation.RequestLimitExceeded1H"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED30S = "FailedOperation.RequestLimitExceeded30S"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_CACHE = "InternalError.Cache"
//  INTERNALERROR_CALLBACK = "InternalError.Callback"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_GENERATEID = "InternalError.GenerateId"
//  INTERNALERROR_MQ = "InternalError.Mq"
//  INTERNALERROR_PDF = "InternalError.Pdf"
//  INTERNALERROR_SERIALIZE = "InternalError.Serialize"
//  INTERNALERROR_STORAGE = "InternalError.Storage"
//  INTERNALERROR_THIRDPARTY = "InternalError.ThirdParty"
//  INTERNALERROR_UNSERIALIZE = "InternalError.UnSerialize"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDVERIFYRESULT = "InvalidParameter.InvalidVerifyResult"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"
//  OPERATIONDENIED_NOAPIAUTH = "OperationDenied.NoApiAuth"
//  OPERATIONDENIED_NOEXTRAVERIFY = "OperationDenied.NoExtraVerify"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOPAYMENTVERIFY = "OperationDenied.NoPaymentVerify"
//  OPERATIONDENIED_NOSESSION = "OperationDenied.NoSession"
//  OPERATIONDENIED_NOVERIFY = "OperationDenied.NoVerify"
//  OPERATIONDENIED_NOTOWNERUIN = "OperationDenied.NotOwnerUin"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateUser(request *CreateUserRequest) (response *CreateUserResponse, err error) {
    if request == nil {
        request = NewCreateUserRequest()
    }
    
    response = NewCreateUserResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserAndSealRequest() (request *CreateUserAndSealRequest) {
    request = &CreateUserAndSealRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "CreateUserAndSeal")
    
    
    return
}

func NewCreateUserAndSealResponse() (response *CreateUserAndSealResponse) {
    response = &CreateUserAndSealResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateUserAndSeal
// 第三方应用可通过此接口（CreateUserAndSeal）注册腾讯电子签实名个人用户，同时系统将为该用户自动生成一个默认电子签名图片。
//
// 
//
// 注意：
//
// 1. 在后续的签署流程中，若未指定签署使用的印章ID，则默认调用自动生成的签名图片进行签署。
//
// 2. 此接口为白名单接口，如您需要使用此能力，请提前与客户经理沟通或邮件至e-contract@tencent.com与我们联系。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTALREADYEXISTS = "FailedOperation.AccountAlreadyExists"
//  FAILEDOPERATION_ACCOUNTVERIFYFAIL = "FailedOperation.AccountVerifyFail"
//  FAILEDOPERATION_AUTHFAIL = "FailedOperation.AuthFail"
//  FAILEDOPERATION_DLOCKFAILED = "FailedOperation.DLockFailed"
//  FAILEDOPERATION_ERRBINDINGREPEATED = "FailedOperation.ErrBindingRepeated"
//  FAILEDOPERATION_GENERATEUSERSEAL = "FailedOperation.GenerateUserSeal"
//  FAILEDOPERATION_NOROLE = "FailedOperation.NoRole"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED1D = "FailedOperation.RequestLimitExceeded1D"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED1H = "FailedOperation.RequestLimitExceeded1H"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED30S = "FailedOperation.RequestLimitExceeded30S"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_CACHE = "InternalError.Cache"
//  INTERNALERROR_CALLBACK = "InternalError.Callback"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_GENERATEID = "InternalError.GenerateId"
//  INTERNALERROR_MQ = "InternalError.Mq"
//  INTERNALERROR_PDF = "InternalError.Pdf"
//  INTERNALERROR_SERIALIZE = "InternalError.Serialize"
//  INTERNALERROR_STORAGE = "InternalError.Storage"
//  INTERNALERROR_THIRDPARTY = "InternalError.ThirdParty"
//  INTERNALERROR_UNSERIALIZE = "InternalError.UnSerialize"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDVERIFYRESULT = "InvalidParameter.InvalidVerifyResult"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"
//  OPERATIONDENIED_NOAPIAUTH = "OperationDenied.NoApiAuth"
//  OPERATIONDENIED_NOEXTRAVERIFY = "OperationDenied.NoExtraVerify"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOPAYMENTVERIFY = "OperationDenied.NoPaymentVerify"
//  OPERATIONDENIED_NOSESSION = "OperationDenied.NoSession"
//  OPERATIONDENIED_NOVERIFY = "OperationDenied.NoVerify"
//  OPERATIONDENIED_NOTOWNERUIN = "OperationDenied.NotOwnerUin"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateUserAndSeal(request *CreateUserAndSealRequest) (response *CreateUserAndSealResponse, err error) {
    if request == nil {
        request = NewCreateUserAndSealRequest()
    }
    
    response = NewCreateUserAndSealResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSealRequest() (request *DeleteSealRequest) {
    request = &DeleteSealRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "DeleteSeal")
    
    
    return
}

func NewDeleteSealResponse() (response *DeleteSealResponse) {
    response = &DeleteSealResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteSeal
// 此接口 (DeleteSeal) 用于删除指定ID的印章。
//
// 
//
// 注意：默认印章不支持删除
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteSeal(request *DeleteSealRequest) (response *DeleteSealResponse, err error) {
    if request == nil {
        request = NewDeleteSealRequest()
    }
    
    response = NewDeleteSealResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCatalogApproversRequest() (request *DescribeCatalogApproversRequest) {
    request = &DescribeCatalogApproversRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "DescribeCatalogApprovers")
    
    
    return
}

func NewDescribeCatalogApproversResponse() (response *DescribeCatalogApproversResponse) {
    response = &DescribeCatalogApproversResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCatalogApprovers
// 第三方应用可通过此接口（DescribeCatalogApprovers）查询指定目录的参与者列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_UNSERIALIZE = "InternalError.UnSerialize"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeCatalogApprovers(request *DescribeCatalogApproversRequest) (response *DescribeCatalogApproversResponse, err error) {
    if request == nil {
        request = NewDescribeCatalogApproversRequest()
    }
    
    response = NewDescribeCatalogApproversResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCatalogSignComponentsRequest() (request *DescribeCatalogSignComponentsRequest) {
    request = &DescribeCatalogSignComponentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "DescribeCatalogSignComponents")
    
    
    return
}

func NewDescribeCatalogSignComponentsResponse() (response *DescribeCatalogSignComponentsResponse) {
    response = &DescribeCatalogSignComponentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCatalogSignComponents
// 第三方应用可通过此接口（DescribeCatalogSignComponents）拉取目录签署区
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTALREADYEXISTS = "FailedOperation.AccountAlreadyExists"
//  FAILEDOPERATION_ACCOUNTVERIFYFAIL = "FailedOperation.AccountVerifyFail"
//  FAILEDOPERATION_AUTHFAIL = "FailedOperation.AuthFail"
//  FAILEDOPERATION_DLOCKFAILED = "FailedOperation.DLockFailed"
//  FAILEDOPERATION_ERRBINDINGREPEATED = "FailedOperation.ErrBindingRepeated"
//  FAILEDOPERATION_GENERATEUSERSEAL = "FailedOperation.GenerateUserSeal"
//  FAILEDOPERATION_NOROLE = "FailedOperation.NoRole"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED1D = "FailedOperation.RequestLimitExceeded1D"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED1H = "FailedOperation.RequestLimitExceeded1H"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED30S = "FailedOperation.RequestLimitExceeded30S"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_CACHE = "InternalError.Cache"
//  INTERNALERROR_CALLBACK = "InternalError.Callback"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_GENERATEID = "InternalError.GenerateId"
//  INTERNALERROR_MQ = "InternalError.Mq"
//  INTERNALERROR_PDF = "InternalError.Pdf"
//  INTERNALERROR_SERIALIZE = "InternalError.Serialize"
//  INTERNALERROR_STORAGE = "InternalError.Storage"
//  INTERNALERROR_THIRDPARTY = "InternalError.ThirdParty"
//  INTERNALERROR_UNSERIALIZE = "InternalError.UnSerialize"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDVERIFYRESULT = "InvalidParameter.InvalidVerifyResult"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"
//  OPERATIONDENIED_NOAPIAUTH = "OperationDenied.NoApiAuth"
//  OPERATIONDENIED_NOEXTRAVERIFY = "OperationDenied.NoExtraVerify"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOPAYMENTVERIFY = "OperationDenied.NoPaymentVerify"
//  OPERATIONDENIED_NOSESSION = "OperationDenied.NoSession"
//  OPERATIONDENIED_NOVERIFY = "OperationDenied.NoVerify"
//  OPERATIONDENIED_NOTOWNERUIN = "OperationDenied.NotOwnerUin"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCatalogSignComponents(request *DescribeCatalogSignComponentsRequest) (response *DescribeCatalogSignComponentsResponse, err error) {
    if request == nil {
        request = NewDescribeCatalogSignComponentsRequest()
    }
    
    response = NewDescribeCatalogSignComponentsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomFlowIdsRequest() (request *DescribeCustomFlowIdsRequest) {
    request = &DescribeCustomFlowIdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "DescribeCustomFlowIds")
    
    
    return
}

func NewDescribeCustomFlowIdsResponse() (response *DescribeCustomFlowIdsResponse) {
    response = &DescribeCustomFlowIdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCustomFlowIds
// 此接口（DescribeCustomFlowIds）用于通过自定义流程id来查询对应的电子签流程id
//
// 可能返回的错误码:
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeCustomFlowIds(request *DescribeCustomFlowIdsRequest) (response *DescribeCustomFlowIdsResponse, err error) {
    if request == nil {
        request = NewDescribeCustomFlowIdsRequest()
    }
    
    response = NewDescribeCustomFlowIdsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomFlowIdsByFlowIdRequest() (request *DescribeCustomFlowIdsByFlowIdRequest) {
    request = &DescribeCustomFlowIdsByFlowIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "DescribeCustomFlowIdsByFlowId")
    
    
    return
}

func NewDescribeCustomFlowIdsByFlowIdResponse() (response *DescribeCustomFlowIdsByFlowIdResponse) {
    response = &DescribeCustomFlowIdsByFlowIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCustomFlowIdsByFlowId
// 此接口（DescribeCustomFlowIdsByFlowId）用于根据流程id反查自定义流程id
//
// 可能返回的错误码:
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeCustomFlowIdsByFlowId(request *DescribeCustomFlowIdsByFlowIdRequest) (response *DescribeCustomFlowIdsByFlowIdResponse, err error) {
    if request == nil {
        request = NewDescribeCustomFlowIdsByFlowIdRequest()
    }
    
    response = NewDescribeCustomFlowIdsByFlowIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFaceIdPhotosRequest() (request *DescribeFaceIdPhotosRequest) {
    request = &DescribeFaceIdPhotosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "DescribeFaceIdPhotos")
    
    
    return
}

func NewDescribeFaceIdPhotosResponse() (response *DescribeFaceIdPhotosResponse) {
    response = &DescribeFaceIdPhotosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFaceIdPhotos
// 该接口为第三方平台向电子签平台获取慧眼人脸核身照片
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeFaceIdPhotos(request *DescribeFaceIdPhotosRequest) (response *DescribeFaceIdPhotosResponse, err error) {
    if request == nil {
        request = NewDescribeFaceIdPhotosRequest()
    }
    
    response = NewDescribeFaceIdPhotosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFaceIdResultsRequest() (request *DescribeFaceIdResultsRequest) {
    request = &DescribeFaceIdResultsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "DescribeFaceIdResults")
    
    
    return
}

func NewDescribeFaceIdResultsResponse() (response *DescribeFaceIdResultsResponse) {
    response = &DescribeFaceIdResultsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFaceIdResults
// 该接口为第三方平台向电子签平台获取慧眼人脸核身结果
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeFaceIdResults(request *DescribeFaceIdResultsRequest) (response *DescribeFaceIdResultsResponse, err error) {
    if request == nil {
        request = NewDescribeFaceIdResultsRequest()
    }
    
    response = NewDescribeFaceIdResultsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFileIdsByCustomIdsRequest() (request *DescribeFileIdsByCustomIdsRequest) {
    request = &DescribeFileIdsByCustomIdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "DescribeFileIdsByCustomIds")
    
    
    return
}

func NewDescribeFileIdsByCustomIdsResponse() (response *DescribeFileIdsByCustomIdsResponse) {
    response = &DescribeFileIdsByCustomIdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFileIdsByCustomIds
// 根据用户自定义id查询文件id
//
// 可能返回的错误码:
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeFileIdsByCustomIds(request *DescribeFileIdsByCustomIdsRequest) (response *DescribeFileIdsByCustomIdsResponse, err error) {
    if request == nil {
        request = NewDescribeFileIdsByCustomIdsRequest()
    }
    
    response = NewDescribeFileIdsByCustomIdsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFileUrlsRequest() (request *DescribeFileUrlsRequest) {
    request = &DescribeFileUrlsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "DescribeFileUrls")
    
    
    return
}

func NewDescribeFileUrlsResponse() (response *DescribeFileUrlsResponse) {
    response = &DescribeFileUrlsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFileUrls
// 此接口（DescribeFileUrls）用于获取签署文件下载的URL。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeFileUrls(request *DescribeFileUrlsRequest) (response *DescribeFileUrlsResponse, err error) {
    if request == nil {
        request = NewDescribeFileUrlsRequest()
    }
    
    response = NewDescribeFileUrlsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFlowRequest() (request *DescribeFlowRequest) {
    request = &DescribeFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "DescribeFlow")
    
    
    return
}

func NewDescribeFlowResponse() (response *DescribeFlowResponse) {
    response = &DescribeFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFlow
// 通过此接口（DescribeFlow）可查询签署流程的详细信息。
//
// 可能返回的错误码:
//  INTERNALERROR_DB = "InternalError.Db"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeFlow(request *DescribeFlowRequest) (response *DescribeFlowResponse, err error) {
    if request == nil {
        request = NewDescribeFlowRequest()
    }
    
    response = NewDescribeFlowResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFlowApproversRequest() (request *DescribeFlowApproversRequest) {
    request = &DescribeFlowApproversRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "DescribeFlowApprovers")
    
    
    return
}

func NewDescribeFlowApproversResponse() (response *DescribeFlowApproversResponse) {
    response = &DescribeFlowApproversResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFlowApprovers
// 第三方应用可通过此接口（DescribeFlowApprovers）查询流程参与者信息。
//
// 可能返回的错误码:
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_UNSERIALIZE = "InternalError.UnSerialize"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeFlowApprovers(request *DescribeFlowApproversRequest) (response *DescribeFlowApproversResponse, err error) {
    if request == nil {
        request = NewDescribeFlowApproversRequest()
    }
    
    response = NewDescribeFlowApproversResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFlowFilesRequest() (request *DescribeFlowFilesRequest) {
    request = &DescribeFlowFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "DescribeFlowFiles")
    
    
    return
}

func NewDescribeFlowFilesResponse() (response *DescribeFlowFilesResponse) {
    response = &DescribeFlowFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFlowFiles
// 查询流程文件
//
// 可能返回的错误码:
//  INTERNALERROR_DB = "InternalError.Db"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeFlowFiles(request *DescribeFlowFilesRequest) (response *DescribeFlowFilesResponse, err error) {
    if request == nil {
        request = NewDescribeFlowFilesRequest()
    }
    
    response = NewDescribeFlowFilesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSealsRequest() (request *DescribeSealsRequest) {
    request = &DescribeSealsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "DescribeSeals")
    
    
    return
}

func NewDescribeSealsResponse() (response *DescribeSealsResponse) {
    response = &DescribeSealsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSeals
// 此接口（DescribeSeals）用于查询指定ID的印章信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSeals(request *DescribeSealsRequest) (response *DescribeSealsResponse, err error) {
    if request == nil {
        request = NewDescribeSealsRequest()
    }
    
    response = NewDescribeSealsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubOrganizationsRequest() (request *DescribeSubOrganizationsRequest) {
    request = &DescribeSubOrganizationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "DescribeSubOrganizations")
    
    
    return
}

func NewDescribeSubOrganizationsResponse() (response *DescribeSubOrganizationsResponse) {
    response = &DescribeSubOrganizationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSubOrganizations
// 此接口（DescribeSubOrganizations）用于查询子机构信息。
//
// 
//
// 注：此接口仅可查询您所属机构应用号创建的子机构信息，不可跨应用/跨机构查询。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTALREADYEXISTS = "FailedOperation.AccountAlreadyExists"
//  FAILEDOPERATION_ACCOUNTVERIFYFAIL = "FailedOperation.AccountVerifyFail"
//  FAILEDOPERATION_AUTHFAIL = "FailedOperation.AuthFail"
//  FAILEDOPERATION_DLOCKFAILED = "FailedOperation.DLockFailed"
//  FAILEDOPERATION_ERRBINDINGREPEATED = "FailedOperation.ErrBindingRepeated"
//  FAILEDOPERATION_GENERATEUSERSEAL = "FailedOperation.GenerateUserSeal"
//  FAILEDOPERATION_NOROLE = "FailedOperation.NoRole"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_CACHE = "InternalError.Cache"
//  INTERNALERROR_CALLBACK = "InternalError.Callback"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_GENERATEID = "InternalError.GenerateId"
//  INTERNALERROR_MQ = "InternalError.Mq"
//  INTERNALERROR_PDF = "InternalError.Pdf"
//  INTERNALERROR_SERIALIZE = "InternalError.Serialize"
//  INTERNALERROR_STORAGE = "InternalError.Storage"
//  INTERNALERROR_THIRDPARTY = "InternalError.ThirdParty"
//  INTERNALERROR_UNSERIALIZE = "InternalError.UnSerialize"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"
//  OPERATIONDENIED_NOAPIAUTH = "OperationDenied.NoApiAuth"
//  OPERATIONDENIED_NOEXTRAVERIFY = "OperationDenied.NoExtraVerify"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOPAYMENTVERIFY = "OperationDenied.NoPaymentVerify"
//  OPERATIONDENIED_NOSESSION = "OperationDenied.NoSession"
//  OPERATIONDENIED_NOVERIFY = "OperationDenied.NoVerify"
//  OPERATIONDENIED_NOTOWNERUIN = "OperationDenied.NotOwnerUin"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSubOrganizations(request *DescribeSubOrganizationsRequest) (response *DescribeSubOrganizationsResponse, err error) {
    if request == nil {
        request = NewDescribeSubOrganizationsRequest()
    }
    
    response = NewDescribeSubOrganizationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsersRequest() (request *DescribeUsersRequest) {
    request = &DescribeUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "DescribeUsers")
    
    
    return
}

func NewDescribeUsersResponse() (response *DescribeUsersResponse) {
    response = &DescribeUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUsers
// 此接口（DescribeUsers）用于查询应用号下的个人用户信息。
//
// 
//
// 注：此接口仅可查询您所属机构应用号创建的个人用户信息，不可跨应用/跨机构查询。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeUsers(request *DescribeUsersRequest) (response *DescribeUsersResponse, err error) {
    if request == nil {
        request = NewDescribeUsersRequest()
    }
    
    response = NewDescribeUsersResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyFlowFileRequest() (request *DestroyFlowFileRequest) {
    request = &DestroyFlowFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "DestroyFlowFile")
    
    
    return
}

func NewDestroyFlowFileResponse() (response *DestroyFlowFileResponse) {
    response = &DestroyFlowFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DestroyFlowFile
// 通过此接口（DestroyFlowFile）可删除指定流程中的合同文件。
//
// 
//
// 注：调用此接口前，请确保此流程已属于归档状态。您可通过查询流程信息接口（DescribeFlow）进行查询。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DestroyFlowFile(request *DestroyFlowFileRequest) (response *DestroyFlowFileResponse, err error) {
    if request == nil {
        request = NewDestroyFlowFileRequest()
    }
    
    response = NewDestroyFlowFileResponse()
    err = c.Send(request, response)
    return
}

func NewGenerateOrganizationSealRequest() (request *GenerateOrganizationSealRequest) {
    request = &GenerateOrganizationSealRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "GenerateOrganizationSeal")
    
    
    return
}

func NewGenerateOrganizationSealResponse() (response *GenerateOrganizationSealResponse) {
    response = &GenerateOrganizationSealResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GenerateOrganizationSeal
// 生成企业电子印章
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GENERATEORGSEAL = "FailedOperation.GenerateOrgSeal"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_GENERATEID = "InternalError.GenerateId"
//  INTERNALERROR_STORAGE = "InternalError.Storage"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GenerateOrganizationSeal(request *GenerateOrganizationSealRequest) (response *GenerateOrganizationSealResponse, err error) {
    if request == nil {
        request = NewGenerateOrganizationSealRequest()
    }
    
    response = NewGenerateOrganizationSealResponse()
    err = c.Send(request, response)
    return
}

func NewGenerateUserSealRequest() (request *GenerateUserSealRequest) {
    request = &GenerateUserSealRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "GenerateUserSeal")
    
    
    return
}

func NewGenerateUserSealResponse() (response *GenerateUserSealResponse) {
    response = &GenerateUserSealResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GenerateUserSeal
// 此接口（GenerateUserSeal）用于生成个人签名图片。
//
// 
//
// 注意：
//
// 1. 个人签名由用户注册时预留的姓名信息生成，不支持自定义签名内容。
//
// 2. 个人用户仅支持拥有一个系统生成的电子签名。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GENERATEUSERSEAL = "FailedOperation.GenerateUserSeal"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_GENERATEID = "InternalError.GenerateId"
//  INTERNALERROR_STORAGE = "InternalError.Storage"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GenerateUserSeal(request *GenerateUserSealRequest) (response *GenerateUserSealResponse, err error) {
    if request == nil {
        request = NewGenerateUserSealRequest()
    }
    
    response = NewGenerateUserSealResponse()
    err = c.Send(request, response)
    return
}

func NewModifyOrganizationDefaultSealRequest() (request *ModifyOrganizationDefaultSealRequest) {
    request = &ModifyOrganizationDefaultSealRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "ModifyOrganizationDefaultSeal")
    
    
    return
}

func NewModifyOrganizationDefaultSealResponse() (response *ModifyOrganizationDefaultSealResponse) {
    response = &ModifyOrganizationDefaultSealResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyOrganizationDefaultSeal
// 此接口 (ModifyOrganizationDefaultSeal) 用于重新指定企业默认印章。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyOrganizationDefaultSeal(request *ModifyOrganizationDefaultSealRequest) (response *ModifyOrganizationDefaultSealResponse, err error) {
    if request == nil {
        request = NewModifyOrganizationDefaultSealRequest()
    }
    
    response = NewModifyOrganizationDefaultSealResponse()
    err = c.Send(request, response)
    return
}

func NewModifySealRequest() (request *ModifySealRequest) {
    request = &ModifySealRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "ModifySeal")
    
    
    return
}

func NewModifySealResponse() (response *ModifySealResponse) {
    response = &ModifySealResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySeal
// 此接口（ModifySeal）用于修改指定印章ID的印章图片和名称。
//
// 
//
// 注：印章类型暂不支持修改，如需调整，请联系客服经理或通过创建印章接口（CreateSeal）进行创建新印章。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_GENERATEID = "InternalError.GenerateId"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifySeal(request *ModifySealRequest) (response *ModifySealResponse, err error) {
    if request == nil {
        request = NewModifySealRequest()
    }
    
    response = NewModifySealResponse()
    err = c.Send(request, response)
    return
}

func NewModifySubOrganizationInfoRequest() (request *ModifySubOrganizationInfoRequest) {
    request = &ModifySubOrganizationInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "ModifySubOrganizationInfo")
    
    
    return
}

func NewModifySubOrganizationInfoResponse() (response *ModifySubOrganizationInfoResponse) {
    response = &ModifySubOrganizationInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySubOrganizationInfo
// 此接口（ModifySubOrganizationInfo）用于更新子机构信息。
//
// 
//
// 注：若修改子机构名称或更新机构证件照片，需要重新通过子机构实名接口（VerifySubOrganization）进行重新实名。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTALREADYEXISTS = "FailedOperation.AccountAlreadyExists"
//  FAILEDOPERATION_ACCOUNTVERIFYFAIL = "FailedOperation.AccountVerifyFail"
//  FAILEDOPERATION_AUTHFAIL = "FailedOperation.AuthFail"
//  FAILEDOPERATION_DLOCKFAILED = "FailedOperation.DLockFailed"
//  FAILEDOPERATION_ERRBINDINGREPEATED = "FailedOperation.ErrBindingRepeated"
//  FAILEDOPERATION_GENERATEUSERSEAL = "FailedOperation.GenerateUserSeal"
//  FAILEDOPERATION_NOROLE = "FailedOperation.NoRole"
//  FAILEDOPERATION_OPENIDALREADYEXISTS = "FailedOperation.OpenIdAlreadyExists"
//  FAILEDOPERATION_ORGIDCARDNUMBERALREADYEXISTS = "FailedOperation.OrgIdCardNumberAlreadyExists"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_CACHE = "InternalError.Cache"
//  INTERNALERROR_CALLBACK = "InternalError.Callback"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_GENERATEID = "InternalError.GenerateId"
//  INTERNALERROR_MQ = "InternalError.Mq"
//  INTERNALERROR_PDF = "InternalError.Pdf"
//  INTERNALERROR_SERIALIZE = "InternalError.Serialize"
//  INTERNALERROR_STORAGE = "InternalError.Storage"
//  INTERNALERROR_THIRDPARTY = "InternalError.ThirdParty"
//  INTERNALERROR_UNSERIALIZE = "InternalError.UnSerialize"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"
//  OPERATIONDENIED_NOAPIAUTH = "OperationDenied.NoApiAuth"
//  OPERATIONDENIED_NOEXTRAVERIFY = "OperationDenied.NoExtraVerify"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOPAYMENTVERIFY = "OperationDenied.NoPaymentVerify"
//  OPERATIONDENIED_NOSESSION = "OperationDenied.NoSession"
//  OPERATIONDENIED_NOVERIFY = "OperationDenied.NoVerify"
//  OPERATIONDENIED_NOTOWNERUIN = "OperationDenied.NotOwnerUin"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifySubOrganizationInfo(request *ModifySubOrganizationInfoRequest) (response *ModifySubOrganizationInfoResponse, err error) {
    if request == nil {
        request = NewModifySubOrganizationInfoRequest()
    }
    
    response = NewModifySubOrganizationInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUserRequest() (request *ModifyUserRequest) {
    request = &ModifyUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "ModifyUser")
    
    
    return
}

func NewModifyUserResponse() (response *ModifyUserResponse) {
    response = &ModifyUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyUser
// 此接口（ModifyUser）用于更新个人用户信息。
//
// 
//
// 注：若修改用户姓名，需要重新通过个人用户实名接口（VerifyUser）进行重新实名。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTALREADYEXISTS = "FailedOperation.AccountAlreadyExists"
//  FAILEDOPERATION_ACCOUNTVERIFYFAIL = "FailedOperation.AccountVerifyFail"
//  FAILEDOPERATION_AUTHFAIL = "FailedOperation.AuthFail"
//  FAILEDOPERATION_DLOCKFAILED = "FailedOperation.DLockFailed"
//  FAILEDOPERATION_ERRBINDINGREPEATED = "FailedOperation.ErrBindingRepeated"
//  FAILEDOPERATION_GENERATEUSERSEAL = "FailedOperation.GenerateUserSeal"
//  FAILEDOPERATION_NOROLE = "FailedOperation.NoRole"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED1D = "FailedOperation.RequestLimitExceeded1D"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED1H = "FailedOperation.RequestLimitExceeded1H"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED30S = "FailedOperation.RequestLimitExceeded30S"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_CACHE = "InternalError.Cache"
//  INTERNALERROR_CALLBACK = "InternalError.Callback"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_GENERATEID = "InternalError.GenerateId"
//  INTERNALERROR_MQ = "InternalError.Mq"
//  INTERNALERROR_PDF = "InternalError.Pdf"
//  INTERNALERROR_SERIALIZE = "InternalError.Serialize"
//  INTERNALERROR_STORAGE = "InternalError.Storage"
//  INTERNALERROR_THIRDPARTY = "InternalError.ThirdParty"
//  INTERNALERROR_UNSERIALIZE = "InternalError.UnSerialize"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDVERIFYRESULT = "InvalidParameter.InvalidVerifyResult"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"
//  OPERATIONDENIED_NOAPIAUTH = "OperationDenied.NoApiAuth"
//  OPERATIONDENIED_NOEXTRAVERIFY = "OperationDenied.NoExtraVerify"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOPAYMENTVERIFY = "OperationDenied.NoPaymentVerify"
//  OPERATIONDENIED_NOSESSION = "OperationDenied.NoSession"
//  OPERATIONDENIED_NOVERIFY = "OperationDenied.NoVerify"
//  OPERATIONDENIED_NOTOWNERUIN = "OperationDenied.NotOwnerUin"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyUser(request *ModifyUserRequest) (response *ModifyUserResponse, err error) {
    if request == nil {
        request = NewModifyUserRequest()
    }
    
    response = NewModifyUserResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUserDefaultSealRequest() (request *ModifyUserDefaultSealRequest) {
    request = &ModifyUserDefaultSealRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "ModifyUserDefaultSeal")
    
    
    return
}

func NewModifyUserDefaultSealResponse() (response *ModifyUserDefaultSealResponse) {
    response = &ModifyUserDefaultSealResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyUserDefaultSeal
// 此接口 (ModifyUserDefaultSeal) 用于重新指定个人默认印章。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyUserDefaultSeal(request *ModifyUserDefaultSealRequest) (response *ModifyUserDefaultSealResponse, err error) {
    if request == nil {
        request = NewModifyUserDefaultSealRequest()
    }
    
    response = NewModifyUserDefaultSealResponse()
    err = c.Send(request, response)
    return
}

func NewRejectFlowRequest() (request *RejectFlowRequest) {
    request = &RejectFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "RejectFlow")
    
    
    return
}

func NewRejectFlowResponse() (response *RejectFlowResponse) {
    response = &RejectFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RejectFlow
// 此接口（RejectFlow）用于用户拒绝签署合同流程。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTALREADYEXISTS = "FailedOperation.AccountAlreadyExists"
//  FAILEDOPERATION_ACCOUNTVERIFYFAIL = "FailedOperation.AccountVerifyFail"
//  FAILEDOPERATION_AUTHFAIL = "FailedOperation.AuthFail"
//  FAILEDOPERATION_DLOCKFAILED = "FailedOperation.DLockFailed"
//  FAILEDOPERATION_ERRBINDINGREPEATED = "FailedOperation.ErrBindingRepeated"
//  FAILEDOPERATION_GENERATEUSERSEAL = "FailedOperation.GenerateUserSeal"
//  FAILEDOPERATION_NOROLE = "FailedOperation.NoRole"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED1D = "FailedOperation.RequestLimitExceeded1D"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED1H = "FailedOperation.RequestLimitExceeded1H"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED30S = "FailedOperation.RequestLimitExceeded30S"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_CACHE = "InternalError.Cache"
//  INTERNALERROR_CALLBACK = "InternalError.Callback"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_GENERATEID = "InternalError.GenerateId"
//  INTERNALERROR_MQ = "InternalError.Mq"
//  INTERNALERROR_PDF = "InternalError.Pdf"
//  INTERNALERROR_SERIALIZE = "InternalError.Serialize"
//  INTERNALERROR_STORAGE = "InternalError.Storage"
//  INTERNALERROR_THIRDPARTY = "InternalError.ThirdParty"
//  INTERNALERROR_UNSERIALIZE = "InternalError.UnSerialize"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDVERIFYRESULT = "InvalidParameter.InvalidVerifyResult"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"
//  OPERATIONDENIED_NOAPIAUTH = "OperationDenied.NoApiAuth"
//  OPERATIONDENIED_NOEXTRAVERIFY = "OperationDenied.NoExtraVerify"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOPAYMENTVERIFY = "OperationDenied.NoPaymentVerify"
//  OPERATIONDENIED_NOSESSION = "OperationDenied.NoSession"
//  OPERATIONDENIED_NOVERIFY = "OperationDenied.NoVerify"
//  OPERATIONDENIED_NOTOWNERUIN = "OperationDenied.NotOwnerUin"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RejectFlow(request *RejectFlowRequest) (response *RejectFlowResponse, err error) {
    if request == nil {
        request = NewRejectFlowRequest()
    }
    
    response = NewRejectFlowResponse()
    err = c.Send(request, response)
    return
}

func NewSendFlowRequest() (request *SendFlowRequest) {
    request = &SendFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "SendFlow")
    
    
    return
}

func NewSendFlowResponse() (response *SendFlowResponse) {
    response = &SendFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SendFlow
// 此接口（SendFlow）用于指定签署者及签署内容，后续可通过生成签署接口（CreateSignUrl）获取签署url。
//
// 可能返回的错误码:
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_GENERATEID = "InternalError.GenerateId"
//  INTERNALERROR_SERIALIZE = "InternalError.Serialize"
//  INTERNALERROR_UNSERIALIZE = "InternalError.UnSerialize"
//  INVALIDPARAMETER_BIZAPPROVERALREADYEXISTS = "InvalidParameter.BizApproverAlreadyExists"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) SendFlow(request *SendFlowRequest) (response *SendFlowResponse, err error) {
    if request == nil {
        request = NewSendFlowRequest()
    }
    
    response = NewSendFlowResponse()
    err = c.Send(request, response)
    return
}

func NewSendFlowUrlRequest() (request *SendFlowUrlRequest) {
    request = &SendFlowUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "SendFlowUrl")
    
    
    return
}

func NewSendFlowUrlResponse() (response *SendFlowUrlResponse) {
    response = &SendFlowUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SendFlowUrl
// 发送流程并获取签署URL
//
// 可能返回的错误码:
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_GENERATEID = "InternalError.GenerateId"
//  INTERNALERROR_PDF = "InternalError.Pdf"
//  INTERNALERROR_SERIALIZE = "InternalError.Serialize"
//  INTERNALERROR_UNSERIALIZE = "InternalError.UnSerialize"
//  INVALIDPARAMETER_BIZAPPROVERALREADYEXISTS = "InvalidParameter.BizApproverAlreadyExists"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SendFlowUrl(request *SendFlowUrlRequest) (response *SendFlowUrlResponse, err error) {
    if request == nil {
        request = NewSendFlowUrlRequest()
    }
    
    response = NewSendFlowUrlResponse()
    err = c.Send(request, response)
    return
}

func NewSendSignInnerVerifyCodeRequest() (request *SendSignInnerVerifyCodeRequest) {
    request = &SendSignInnerVerifyCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "SendSignInnerVerifyCode")
    
    
    return
}

func NewSendSignInnerVerifyCodeResponse() (response *SendSignInnerVerifyCodeResponse) {
    response = &SendSignInnerVerifyCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SendSignInnerVerifyCode
// 此接口用于发送签署验证码
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED1D = "FailedOperation.RequestLimitExceeded1D"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED1H = "FailedOperation.RequestLimitExceeded1H"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED30S = "FailedOperation.RequestLimitExceeded30S"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_CACHE = "InternalError.Cache"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_SERIALIZE = "InternalError.Serialize"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SendSignInnerVerifyCode(request *SendSignInnerVerifyCodeRequest) (response *SendSignInnerVerifyCodeResponse, err error) {
    if request == nil {
        request = NewSendSignInnerVerifyCodeRequest()
    }
    
    response = NewSendSignInnerVerifyCodeResponse()
    err = c.Send(request, response)
    return
}

func NewSignFlowRequest() (request *SignFlowRequest) {
    request = &SignFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "SignFlow")
    
    
    return
}

func NewSignFlowResponse() (response *SignFlowResponse) {
    response = &SignFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SignFlow
// 此接口（SignFlow）可用于对流程文件进行签署。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDVERIFYRESULT = "InvalidParameter.InvalidVerifyResult"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SignFlow(request *SignFlowRequest) (response *SignFlowResponse, err error) {
    if request == nil {
        request = NewSignFlowRequest()
    }
    
    response = NewSignFlowResponse()
    err = c.Send(request, response)
    return
}

func NewUploadFilesRequest() (request *UploadFilesRequest) {
    request = &UploadFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "UploadFiles")
    
    
    return
}

func NewUploadFilesResponse() (response *UploadFilesResponse) {
    response = &UploadFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UploadFiles
// 此接口（UploadFiles）用于文件上传。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SERIALIZE = "InternalError.Serialize"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UploadFiles(request *UploadFilesRequest) (response *UploadFilesResponse, err error) {
    if request == nil {
        request = NewUploadFilesRequest()
    }
    
    response = NewUploadFilesResponse()
    err = c.Send(request, response)
    return
}

func NewVerifySubOrganizationRequest() (request *VerifySubOrganizationRequest) {
    request = &VerifySubOrganizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "VerifySubOrganization")
    
    
    return
}

func NewVerifySubOrganizationResponse() (response *VerifySubOrganizationResponse) {
    response = &VerifySubOrganizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// VerifySubOrganization
// 此接口（VerifySubOrganization）用于通过子机构的实名认证。
//
// 
//
// 注：此接口为白名单接口，如您需要使用此能力，请提前与客户经理沟通或邮件至e-contract@tencent.com与我们联系。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ACCOUNTALREADYEXISTS = "FailedOperation.AccountAlreadyExists"
//  FAILEDOPERATION_ACCOUNTVERIFYFAIL = "FailedOperation.AccountVerifyFail"
//  FAILEDOPERATION_AUTHFAIL = "FailedOperation.AuthFail"
//  FAILEDOPERATION_DLOCKFAILED = "FailedOperation.DLockFailed"
//  FAILEDOPERATION_ERRBINDINGREPEATED = "FailedOperation.ErrBindingRepeated"
//  FAILEDOPERATION_GENERATEUSERSEAL = "FailedOperation.GenerateUserSeal"
//  FAILEDOPERATION_NOROLE = "FailedOperation.NoRole"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_CACHE = "InternalError.Cache"
//  INTERNALERROR_CALLBACK = "InternalError.Callback"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_GENERATEID = "InternalError.GenerateId"
//  INTERNALERROR_MQ = "InternalError.Mq"
//  INTERNALERROR_PDF = "InternalError.Pdf"
//  INTERNALERROR_SERIALIZE = "InternalError.Serialize"
//  INTERNALERROR_STORAGE = "InternalError.Storage"
//  INTERNALERROR_THIRDPARTY = "InternalError.ThirdParty"
//  INTERNALERROR_UNSERIALIZE = "InternalError.UnSerialize"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BANNEDAPPLICATION = "OperationDenied.BannedApplication"
//  OPERATIONDENIED_NOAPIAUTH = "OperationDenied.NoApiAuth"
//  OPERATIONDENIED_NOEXTRAVERIFY = "OperationDenied.NoExtraVerify"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOPAYMENTVERIFY = "OperationDenied.NoPaymentVerify"
//  OPERATIONDENIED_NOSESSION = "OperationDenied.NoSession"
//  OPERATIONDENIED_NOVERIFY = "OperationDenied.NoVerify"
//  OPERATIONDENIED_NOTOWNERUIN = "OperationDenied.NotOwnerUin"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) VerifySubOrganization(request *VerifySubOrganizationRequest) (response *VerifySubOrganizationResponse, err error) {
    if request == nil {
        request = NewVerifySubOrganizationRequest()
    }
    
    response = NewVerifySubOrganizationResponse()
    err = c.Send(request, response)
    return
}

func NewVerifyUserRequest() (request *VerifyUserRequest) {
    request = &VerifyUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("essbasic", APIVersion, "VerifyUser")
    
    
    return
}

func NewVerifyUserResponse() (response *VerifyUserResponse) {
    response = &VerifyUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// VerifyUser
// 第三方应用可通过此接口（VerifyUser）将腾讯电子签个人用户的实名认证状态设为通过。
//
// 
//
// 注：此接口为白名单接口，如您需要使用此能力，请提前与客户经理沟通或邮件至e-contract@tencent.com与我们联系。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) VerifyUser(request *VerifyUserRequest) (response *VerifyUserResponse, err error) {
    if request == nil {
        request = NewVerifyUserRequest()
    }
    
    response = NewVerifyUserResponse()
    err = c.Send(request, response)
    return
}
