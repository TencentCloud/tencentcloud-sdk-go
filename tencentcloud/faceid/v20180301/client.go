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

package v20180301

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-03-01"

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


func NewBankCard2EVerificationRequest() (request *BankCard2EVerificationRequest) {
    request = &BankCard2EVerificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "BankCard2EVerification")
    
    
    return
}

func NewBankCard2EVerificationResponse() (response *BankCard2EVerificationResponse) {
    response = &BankCard2EVerificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BankCard2EVerification
// 本接口用于校验姓名和银行卡号的真实性和一致性。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTENCRYPTFIELD = "InvalidParameter.UnsupportEncryptField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_CHARGESTATUSEXCEPTION = "UnauthorizedOperation.ChargeStatusException"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) BankCard2EVerification(request *BankCard2EVerificationRequest) (response *BankCard2EVerificationResponse, err error) {
    return c.BankCard2EVerificationWithContext(context.Background(), request)
}

// BankCard2EVerification
// 本接口用于校验姓名和银行卡号的真实性和一致性。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTENCRYPTFIELD = "InvalidParameter.UnsupportEncryptField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_CHARGESTATUSEXCEPTION = "UnauthorizedOperation.ChargeStatusException"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) BankCard2EVerificationWithContext(ctx context.Context, request *BankCard2EVerificationRequest) (response *BankCard2EVerificationResponse, err error) {
    if request == nil {
        request = NewBankCard2EVerificationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BankCard2EVerification require credential")
    }

    request.SetContext(ctx)
    
    response = NewBankCard2EVerificationResponse()
    err = c.Send(request, response)
    return
}

func NewBankCard4EVerificationRequest() (request *BankCard4EVerificationRequest) {
    request = &BankCard4EVerificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "BankCard4EVerification")
    
    
    return
}

func NewBankCard4EVerificationResponse() (response *BankCard4EVerificationResponse) {
    response = &BankCard4EVerificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BankCard4EVerification
// 本接口用于输入银行卡号、姓名、开户证件号、开户手机号，校验信息的真实性和一致性。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTENCRYPTFIELD = "InvalidParameter.UnsupportEncryptField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_CHARGESTATUSEXCEPTION = "UnauthorizedOperation.ChargeStatusException"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) BankCard4EVerification(request *BankCard4EVerificationRequest) (response *BankCard4EVerificationResponse, err error) {
    return c.BankCard4EVerificationWithContext(context.Background(), request)
}

// BankCard4EVerification
// 本接口用于输入银行卡号、姓名、开户证件号、开户手机号，校验信息的真实性和一致性。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTENCRYPTFIELD = "InvalidParameter.UnsupportEncryptField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_CHARGESTATUSEXCEPTION = "UnauthorizedOperation.ChargeStatusException"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) BankCard4EVerificationWithContext(ctx context.Context, request *BankCard4EVerificationRequest) (response *BankCard4EVerificationResponse, err error) {
    if request == nil {
        request = NewBankCard4EVerificationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BankCard4EVerification require credential")
    }

    request.SetContext(ctx)
    
    response = NewBankCard4EVerificationResponse()
    err = c.Send(request, response)
    return
}

func NewBankCardVerificationRequest() (request *BankCardVerificationRequest) {
    request = &BankCardVerificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "BankCardVerification")
    
    
    return
}

func NewBankCardVerificationResponse() (response *BankCardVerificationResponse) {
    response = &BankCardVerificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BankCardVerification
// 本接口用于银行卡号、姓名、开户证件号信息的真实性和一致性。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTENCRYPTFIELD = "InvalidParameter.UnsupportEncryptField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_CHARGESTATUSEXCEPTION = "UnauthorizedOperation.ChargeStatusException"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) BankCardVerification(request *BankCardVerificationRequest) (response *BankCardVerificationResponse, err error) {
    return c.BankCardVerificationWithContext(context.Background(), request)
}

// BankCardVerification
// 本接口用于银行卡号、姓名、开户证件号信息的真实性和一致性。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTENCRYPTFIELD = "InvalidParameter.UnsupportEncryptField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_CHARGESTATUSEXCEPTION = "UnauthorizedOperation.ChargeStatusException"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) BankCardVerificationWithContext(ctx context.Context, request *BankCardVerificationRequest) (response *BankCardVerificationResponse, err error) {
    if request == nil {
        request = NewBankCardVerificationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BankCardVerification require credential")
    }

    request.SetContext(ctx)
    
    response = NewBankCardVerificationResponse()
    err = c.Send(request, response)
    return
}

func NewCheckBankCardInformationRequest() (request *CheckBankCardInformationRequest) {
    request = &CheckBankCardInformationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "CheckBankCardInformation")
    
    
    return
}

func NewCheckBankCardInformationResponse() (response *CheckBankCardInformationResponse) {
    response = &CheckBankCardInformationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckBankCardInformation
// 银行卡基础信息查询
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_UNSUPPORTENCRYPTFIELD = "InvalidParameter.UnsupportEncryptField"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_CHARGESTATUSEXCEPTION = "UnauthorizedOperation.ChargeStatusException"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) CheckBankCardInformation(request *CheckBankCardInformationRequest) (response *CheckBankCardInformationResponse, err error) {
    return c.CheckBankCardInformationWithContext(context.Background(), request)
}

// CheckBankCardInformation
// 银行卡基础信息查询
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_UNSUPPORTENCRYPTFIELD = "InvalidParameter.UnsupportEncryptField"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_CHARGESTATUSEXCEPTION = "UnauthorizedOperation.ChargeStatusException"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) CheckBankCardInformationWithContext(ctx context.Context, request *CheckBankCardInformationRequest) (response *CheckBankCardInformationResponse, err error) {
    if request == nil {
        request = NewCheckBankCardInformationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckBankCardInformation require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckBankCardInformationResponse()
    err = c.Send(request, response)
    return
}

func NewCheckEidTokenStatusRequest() (request *CheckEidTokenStatusRequest) {
    request = &CheckEidTokenStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "CheckEidTokenStatus")
    
    
    return
}

func NewCheckEidTokenStatusResponse() (response *CheckEidTokenStatusResponse) {
    response = &CheckEidTokenStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckEidTokenStatus
// 用于轮询E证通H5场景EidToken验证状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DbError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BIZTOKENEXPIRED = "InvalidParameterValue.BizTokenExpired"
//  INVALIDPARAMETERVALUE_BIZTOKENILLEGAL = "InvalidParameterValue.BizTokenIllegal"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) CheckEidTokenStatus(request *CheckEidTokenStatusRequest) (response *CheckEidTokenStatusResponse, err error) {
    return c.CheckEidTokenStatusWithContext(context.Background(), request)
}

// CheckEidTokenStatus
// 用于轮询E证通H5场景EidToken验证状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DbError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BIZTOKENEXPIRED = "InvalidParameterValue.BizTokenExpired"
//  INVALIDPARAMETERVALUE_BIZTOKENILLEGAL = "InvalidParameterValue.BizTokenIllegal"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) CheckEidTokenStatusWithContext(ctx context.Context, request *CheckEidTokenStatusRequest) (response *CheckEidTokenStatusResponse, err error) {
    if request == nil {
        request = NewCheckEidTokenStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckEidTokenStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckEidTokenStatusResponse()
    err = c.Send(request, response)
    return
}

func NewCheckIdCardInformationRequest() (request *CheckIdCardInformationRequest) {
    request = &CheckIdCardInformationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "CheckIdCardInformation")
    
    
    return
}

func NewCheckIdCardInformationResponse() (response *CheckIdCardInformationResponse) {
    response = &CheckIdCardInformationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckIdCardInformation
// 传入身份证人像面照片，识别身份证照片上的信息，并将姓名、身份证号、身份证人像照片与权威库的证件照进行比对，是否属于同一个人，从而验证身份证信息的真实性。
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPAREFAIL = "FailedOperation.CompareFail"
//  FAILEDOPERATION_COMPARELOWSIMILARITY = "FailedOperation.CompareLowSimilarity"
//  FAILEDOPERATION_COMPARESYSTEMERROR = "FailedOperation.CompareSystemError"
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_ENCRYPTSYSTEMERROR = "FailedOperation.EncryptSystemError"
//  FAILEDOPERATION_FILESAVEERROR = "FailedOperation.FileSaveError"
//  FAILEDOPERATION_IDFORMATERROR = "FailedOperation.IdFormatError"
//  FAILEDOPERATION_IDNAMEMISMATCH = "FailedOperation.IdNameMisMatch"
//  FAILEDOPERATION_IDNOEXISTSYSTEM = "FailedOperation.IdNoExistSystem"
//  FAILEDOPERATION_IDPHOTONOEXIST = "FailedOperation.IdPhotoNoExist"
//  FAILEDOPERATION_IDPHOTOPOORQUALITY = "FailedOperation.IdPhotoPoorQuality"
//  FAILEDOPERATION_IDPHOTOSYSTEMNOANSWER = "FailedOperation.IdPhotoSystemNoanswer"
//  FAILEDOPERATION_IMAGEBLUR = "FailedOperation.ImageBlur"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOIDCARD = "FailedOperation.ImageNoIdCard"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_LIFEPHOTODETECTFACES = "FailedOperation.LifePhotoDetectFaces"
//  FAILEDOPERATION_LIFEPHOTODETECTFAKE = "FailedOperation.LifePhotoDetectFake"
//  FAILEDOPERATION_LIFEPHOTODETECTNOFACES = "FailedOperation.LifePhotoDetectNoFaces"
//  FAILEDOPERATION_LIFEPHOTOPOORQUALITY = "FailedOperation.LifePhotoPoorQuality"
//  FAILEDOPERATION_LIFEPHOTOSIZEERROR = "FailedOperation.LifePhotoSizeError"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INTERNALERROR_ENCRYPTSYSTEMERROR = "InternalError.EncryptSystemError"
//  INTERNALERROR_UNKNOWN = "InternalError.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_CHARGESTATUSEXCEPTION = "UnauthorizedOperation.ChargeStatusException"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) CheckIdCardInformation(request *CheckIdCardInformationRequest) (response *CheckIdCardInformationResponse, err error) {
    return c.CheckIdCardInformationWithContext(context.Background(), request)
}

// CheckIdCardInformation
// 传入身份证人像面照片，识别身份证照片上的信息，并将姓名、身份证号、身份证人像照片与权威库的证件照进行比对，是否属于同一个人，从而验证身份证信息的真实性。
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPAREFAIL = "FailedOperation.CompareFail"
//  FAILEDOPERATION_COMPARELOWSIMILARITY = "FailedOperation.CompareLowSimilarity"
//  FAILEDOPERATION_COMPARESYSTEMERROR = "FailedOperation.CompareSystemError"
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_ENCRYPTSYSTEMERROR = "FailedOperation.EncryptSystemError"
//  FAILEDOPERATION_FILESAVEERROR = "FailedOperation.FileSaveError"
//  FAILEDOPERATION_IDFORMATERROR = "FailedOperation.IdFormatError"
//  FAILEDOPERATION_IDNAMEMISMATCH = "FailedOperation.IdNameMisMatch"
//  FAILEDOPERATION_IDNOEXISTSYSTEM = "FailedOperation.IdNoExistSystem"
//  FAILEDOPERATION_IDPHOTONOEXIST = "FailedOperation.IdPhotoNoExist"
//  FAILEDOPERATION_IDPHOTOPOORQUALITY = "FailedOperation.IdPhotoPoorQuality"
//  FAILEDOPERATION_IDPHOTOSYSTEMNOANSWER = "FailedOperation.IdPhotoSystemNoanswer"
//  FAILEDOPERATION_IMAGEBLUR = "FailedOperation.ImageBlur"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOIDCARD = "FailedOperation.ImageNoIdCard"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_LIFEPHOTODETECTFACES = "FailedOperation.LifePhotoDetectFaces"
//  FAILEDOPERATION_LIFEPHOTODETECTFAKE = "FailedOperation.LifePhotoDetectFake"
//  FAILEDOPERATION_LIFEPHOTODETECTNOFACES = "FailedOperation.LifePhotoDetectNoFaces"
//  FAILEDOPERATION_LIFEPHOTOPOORQUALITY = "FailedOperation.LifePhotoPoorQuality"
//  FAILEDOPERATION_LIFEPHOTOSIZEERROR = "FailedOperation.LifePhotoSizeError"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INTERNALERROR_ENCRYPTSYSTEMERROR = "InternalError.EncryptSystemError"
//  INTERNALERROR_UNKNOWN = "InternalError.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_CHARGESTATUSEXCEPTION = "UnauthorizedOperation.ChargeStatusException"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) CheckIdCardInformationWithContext(ctx context.Context, request *CheckIdCardInformationRequest) (response *CheckIdCardInformationResponse, err error) {
    if request == nil {
        request = NewCheckIdCardInformationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckIdCardInformation require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckIdCardInformationResponse()
    err = c.Send(request, response)
    return
}

func NewCheckIdNameDateRequest() (request *CheckIdNameDateRequest) {
    request = &CheckIdNameDateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "CheckIdNameDate")
    
    
    return
}

func NewCheckIdNameDateResponse() (response *CheckIdNameDateResponse) {
    response = &CheckIdNameDateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckIdNameDate
// 本接口用于校验姓名、身份证号、身份证有效期的真实性和一致性。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_UNSUPPORTENCRYPTFIELD = "InvalidParameter.UnsupportEncryptField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_CHARGESTATUSEXCEPTION = "UnauthorizedOperation.ChargeStatusException"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) CheckIdNameDate(request *CheckIdNameDateRequest) (response *CheckIdNameDateResponse, err error) {
    return c.CheckIdNameDateWithContext(context.Background(), request)
}

// CheckIdNameDate
// 本接口用于校验姓名、身份证号、身份证有效期的真实性和一致性。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_UNSUPPORTENCRYPTFIELD = "InvalidParameter.UnsupportEncryptField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_CHARGESTATUSEXCEPTION = "UnauthorizedOperation.ChargeStatusException"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) CheckIdNameDateWithContext(ctx context.Context, request *CheckIdNameDateRequest) (response *CheckIdNameDateResponse, err error) {
    if request == nil {
        request = NewCheckIdNameDateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckIdNameDate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckIdNameDateResponse()
    err = c.Send(request, response)
    return
}

func NewCheckPhoneAndNameRequest() (request *CheckPhoneAndNameRequest) {
    request = &CheckPhoneAndNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "CheckPhoneAndName")
    
    
    return
}

func NewCheckPhoneAndNameResponse() (response *CheckPhoneAndNameResponse) {
    response = &CheckPhoneAndNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckPhoneAndName
// 手机号二要素核验接口用于校验手机号和姓名的真实性和一致性，支持的手机号段详情请查阅<a href="https://cloud.tencent.com/document/product/1007/46063">运营商类</a>文档。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_UNSUPPORTENCRYPTFIELD = "InvalidParameter.UnsupportEncryptField"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_CHARGESTATUSEXCEPTION = "UnauthorizedOperation.ChargeStatusException"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) CheckPhoneAndName(request *CheckPhoneAndNameRequest) (response *CheckPhoneAndNameResponse, err error) {
    return c.CheckPhoneAndNameWithContext(context.Background(), request)
}

// CheckPhoneAndName
// 手机号二要素核验接口用于校验手机号和姓名的真实性和一致性，支持的手机号段详情请查阅<a href="https://cloud.tencent.com/document/product/1007/46063">运营商类</a>文档。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_UNSUPPORTENCRYPTFIELD = "InvalidParameter.UnsupportEncryptField"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_CHARGESTATUSEXCEPTION = "UnauthorizedOperation.ChargeStatusException"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) CheckPhoneAndNameWithContext(ctx context.Context, request *CheckPhoneAndNameRequest) (response *CheckPhoneAndNameResponse, err error) {
    if request == nil {
        request = NewCheckPhoneAndNameRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckPhoneAndName require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckPhoneAndNameResponse()
    err = c.Send(request, response)
    return
}

func NewDetectAuthRequest() (request *DetectAuthRequest) {
    request = &DetectAuthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "DetectAuth")
    
    
    return
}

func NewDetectAuthResponse() (response *DetectAuthResponse) {
    response = &DetectAuthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetectAuth
// 每次调用人脸核身SaaS化服务前，需先调用本接口获取BizToken，用来串联核身流程，在验证完成后，用于获取验证结果信息。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTENCRYPTFIELD = "InvalidParameter.UnsupportEncryptField"
//  INVALIDPARAMETERVALUE_RULEIDDISABLED = "InvalidParameterValue.RuleIdDisabled"
//  INVALIDPARAMETERVALUE_RULEIDNOTEXIST = "InvalidParameterValue.RuleIdNotExist"
//  UNAUTHORIZEDOPERATION_ACTIVATEERROR = "UnauthorizedOperation.ActivateError"
//  UNAUTHORIZEDOPERATION_ACTIVATING = "UnauthorizedOperation.Activating"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) DetectAuth(request *DetectAuthRequest) (response *DetectAuthResponse, err error) {
    return c.DetectAuthWithContext(context.Background(), request)
}

// DetectAuth
// 每次调用人脸核身SaaS化服务前，需先调用本接口获取BizToken，用来串联核身流程，在验证完成后，用于获取验证结果信息。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTENCRYPTFIELD = "InvalidParameter.UnsupportEncryptField"
//  INVALIDPARAMETERVALUE_RULEIDDISABLED = "InvalidParameterValue.RuleIdDisabled"
//  INVALIDPARAMETERVALUE_RULEIDNOTEXIST = "InvalidParameterValue.RuleIdNotExist"
//  UNAUTHORIZEDOPERATION_ACTIVATEERROR = "UnauthorizedOperation.ActivateError"
//  UNAUTHORIZEDOPERATION_ACTIVATING = "UnauthorizedOperation.Activating"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) DetectAuthWithContext(ctx context.Context, request *DetectAuthRequest) (response *DetectAuthResponse, err error) {
    if request == nil {
        request = NewDetectAuthRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetectAuth require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetectAuthResponse()
    err = c.Send(request, response)
    return
}

func NewEncryptedPhoneVerificationRequest() (request *EncryptedPhoneVerificationRequest) {
    request = &EncryptedPhoneVerificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "EncryptedPhoneVerification")
    
    
    return
}

func NewEncryptedPhoneVerificationResponse() (response *EncryptedPhoneVerificationResponse) {
    response = &EncryptedPhoneVerificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EncryptedPhoneVerification
// 本接口用于校验手机号、姓名和身份证号的真实性和一致性，入参支持明文、MD5和SHA256加密传输。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTENCRYPTFIELD = "InvalidParameter.UnsupportEncryptField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_CHARGESTATUSEXCEPTION = "UnauthorizedOperation.ChargeStatusException"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) EncryptedPhoneVerification(request *EncryptedPhoneVerificationRequest) (response *EncryptedPhoneVerificationResponse, err error) {
    return c.EncryptedPhoneVerificationWithContext(context.Background(), request)
}

// EncryptedPhoneVerification
// 本接口用于校验手机号、姓名和身份证号的真实性和一致性，入参支持明文、MD5和SHA256加密传输。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTENCRYPTFIELD = "InvalidParameter.UnsupportEncryptField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_CHARGESTATUSEXCEPTION = "UnauthorizedOperation.ChargeStatusException"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) EncryptedPhoneVerificationWithContext(ctx context.Context, request *EncryptedPhoneVerificationRequest) (response *EncryptedPhoneVerificationResponse, err error) {
    if request == nil {
        request = NewEncryptedPhoneVerificationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EncryptedPhoneVerification require credential")
    }

    request.SetContext(ctx)
    
    response = NewEncryptedPhoneVerificationResponse()
    err = c.Send(request, response)
    return
}

func NewGetActionSequenceRequest() (request *GetActionSequenceRequest) {
    request = &GetActionSequenceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "GetActionSequence")
    
    
    return
}

func NewGetActionSequenceResponse() (response *GetActionSequenceResponse) {
    response = &GetActionSequenceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetActionSequence
// 使用动作活体检测模式前，需调用本接口获取动作顺序。
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetActionSequence(request *GetActionSequenceRequest) (response *GetActionSequenceResponse, err error) {
    return c.GetActionSequenceWithContext(context.Background(), request)
}

// GetActionSequence
// 使用动作活体检测模式前，需调用本接口获取动作顺序。
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetActionSequenceWithContext(ctx context.Context, request *GetActionSequenceRequest) (response *GetActionSequenceResponse, err error) {
    if request == nil {
        request = NewGetActionSequenceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetActionSequence require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetActionSequenceResponse()
    err = c.Send(request, response)
    return
}

func NewGetDetectInfoRequest() (request *GetDetectInfoRequest) {
    request = &GetDetectInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "GetDetectInfo")
    
    
    return
}

func NewGetDetectInfoResponse() (response *GetDetectInfoResponse) {
    response = &GetDetectInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetDetectInfo
// 完成验证后，用BizToken调用本接口获取结果信息，BizToken生成后三天内（3\*24\*3,600秒）可多次拉取。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BIZTOKENEXPIRED = "InvalidParameterValue.BizTokenExpired"
//  INVALIDPARAMETERVALUE_RULEIDNOTEXIST = "InvalidParameterValue.RuleIdNotExist"
//  UNAUTHORIZEDOPERATION_ACTIVATEERROR = "UnauthorizedOperation.ActivateError"
//  UNAUTHORIZEDOPERATION_ACTIVATING = "UnauthorizedOperation.Activating"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) GetDetectInfo(request *GetDetectInfoRequest) (response *GetDetectInfoResponse, err error) {
    return c.GetDetectInfoWithContext(context.Background(), request)
}

// GetDetectInfo
// 完成验证后，用BizToken调用本接口获取结果信息，BizToken生成后三天内（3\*24\*3,600秒）可多次拉取。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BIZTOKENEXPIRED = "InvalidParameterValue.BizTokenExpired"
//  INVALIDPARAMETERVALUE_RULEIDNOTEXIST = "InvalidParameterValue.RuleIdNotExist"
//  UNAUTHORIZEDOPERATION_ACTIVATEERROR = "UnauthorizedOperation.ActivateError"
//  UNAUTHORIZEDOPERATION_ACTIVATING = "UnauthorizedOperation.Activating"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) GetDetectInfoWithContext(ctx context.Context, request *GetDetectInfoRequest) (response *GetDetectInfoResponse, err error) {
    if request == nil {
        request = NewGetDetectInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDetectInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDetectInfoResponse()
    err = c.Send(request, response)
    return
}

func NewGetDetectInfoEnhancedRequest() (request *GetDetectInfoEnhancedRequest) {
    request = &GetDetectInfoEnhancedRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "GetDetectInfoEnhanced")
    
    
    return
}

func NewGetDetectInfoEnhancedResponse() (response *GetDetectInfoEnhancedResponse) {
    response = &GetDetectInfoEnhancedResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetDetectInfoEnhanced
// 完成验证后，用BizToken调用本接口获取结果信息，BizToken生成后三天内（3\*24\*3,600秒）可多次拉取。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_ENCRYPTSYSTEMERROR = "FailedOperation.EncryptSystemError"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BIZTOKENEXPIRED = "InvalidParameterValue.BizTokenExpired"
//  INVALIDPARAMETERVALUE_RULEIDNOTEXIST = "InvalidParameterValue.RuleIdNotExist"
//  UNAUTHORIZEDOPERATION_ACTIVATEERROR = "UnauthorizedOperation.ActivateError"
//  UNAUTHORIZEDOPERATION_ACTIVATING = "UnauthorizedOperation.Activating"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) GetDetectInfoEnhanced(request *GetDetectInfoEnhancedRequest) (response *GetDetectInfoEnhancedResponse, err error) {
    return c.GetDetectInfoEnhancedWithContext(context.Background(), request)
}

// GetDetectInfoEnhanced
// 完成验证后，用BizToken调用本接口获取结果信息，BizToken生成后三天内（3\*24\*3,600秒）可多次拉取。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_ENCRYPTSYSTEMERROR = "FailedOperation.EncryptSystemError"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BIZTOKENEXPIRED = "InvalidParameterValue.BizTokenExpired"
//  INVALIDPARAMETERVALUE_RULEIDNOTEXIST = "InvalidParameterValue.RuleIdNotExist"
//  UNAUTHORIZEDOPERATION_ACTIVATEERROR = "UnauthorizedOperation.ActivateError"
//  UNAUTHORIZEDOPERATION_ACTIVATING = "UnauthorizedOperation.Activating"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) GetDetectInfoEnhancedWithContext(ctx context.Context, request *GetDetectInfoEnhancedRequest) (response *GetDetectInfoEnhancedResponse, err error) {
    if request == nil {
        request = NewGetDetectInfoEnhancedRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDetectInfoEnhanced require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDetectInfoEnhancedResponse()
    err = c.Send(request, response)
    return
}

func NewGetEidResultRequest() (request *GetEidResultRequest) {
    request = &GetEidResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "GetEidResult")
    
    
    return
}

func NewGetEidResultResponse() (response *GetEidResultResponse) {
    response = &GetEidResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetEidResult
// 完成验证后，用EidToken调用本接口获取结果信息，EidToken生成后三天内（3\*24\*3,600秒）可多次拉取。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ENCRYPTSYSTEMERROR = "InternalError.EncryptSystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BIZTOKENEXPIRED = "InvalidParameterValue.BizTokenExpired"
//  INVALIDPARAMETERVALUE_BIZTOKENILLEGAL = "InvalidParameterValue.BizTokenIllegal"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) GetEidResult(request *GetEidResultRequest) (response *GetEidResultResponse, err error) {
    return c.GetEidResultWithContext(context.Background(), request)
}

// GetEidResult
// 完成验证后，用EidToken调用本接口获取结果信息，EidToken生成后三天内（3\*24\*3,600秒）可多次拉取。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ENCRYPTSYSTEMERROR = "InternalError.EncryptSystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BIZTOKENEXPIRED = "InvalidParameterValue.BizTokenExpired"
//  INVALIDPARAMETERVALUE_BIZTOKENILLEGAL = "InvalidParameterValue.BizTokenIllegal"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) GetEidResultWithContext(ctx context.Context, request *GetEidResultRequest) (response *GetEidResultResponse, err error) {
    if request == nil {
        request = NewGetEidResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetEidResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetEidResultResponse()
    err = c.Send(request, response)
    return
}

func NewGetEidTokenRequest() (request *GetEidTokenRequest) {
    request = &GetEidTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "GetEidToken")
    
    
    return
}

func NewGetEidTokenResponse() (response *GetEidTokenResponse) {
    response = &GetEidTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetEidToken
// 每次调用E证通服务前，需先调用本接口获取EidToken，用来串联E证通流程，在验证完成后，用于获取E证通结果信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  FAILEDOPERATION_UNREGISTEREDEID = "FailedOperation.UnregisteredEid"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTENCRYPTFIELD = "InvalidParameter.UnsupportEncryptField"
//  INVALIDPARAMETERVALUE_RULEIDDISABLED = "InvalidParameterValue.RuleIdDisabled"
//  INVALIDPARAMETERVALUE_RULEIDNOTEXIST = "InvalidParameterValue.RuleIdNotExist"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) GetEidToken(request *GetEidTokenRequest) (response *GetEidTokenResponse, err error) {
    return c.GetEidTokenWithContext(context.Background(), request)
}

// GetEidToken
// 每次调用E证通服务前，需先调用本接口获取EidToken，用来串联E证通流程，在验证完成后，用于获取E证通结果信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  FAILEDOPERATION_UNREGISTEREDEID = "FailedOperation.UnregisteredEid"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTENCRYPTFIELD = "InvalidParameter.UnsupportEncryptField"
//  INVALIDPARAMETERVALUE_RULEIDDISABLED = "InvalidParameterValue.RuleIdDisabled"
//  INVALIDPARAMETERVALUE_RULEIDNOTEXIST = "InvalidParameterValue.RuleIdNotExist"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) GetEidTokenWithContext(ctx context.Context, request *GetEidTokenRequest) (response *GetEidTokenResponse, err error) {
    if request == nil {
        request = NewGetEidTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetEidToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetEidTokenResponse()
    err = c.Send(request, response)
    return
}

func NewGetFaceIdResultRequest() (request *GetFaceIdResultRequest) {
    request = &GetFaceIdResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "GetFaceIdResult")
    
    
    return
}

func NewGetFaceIdResultResponse() (response *GetFaceIdResultResponse) {
    response = &GetFaceIdResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetFaceIdResult
// 完成验证后，用FaceIdToken调用本接口获取结果信息，FaceIdToken生成后三天内（3\*24\*3,600秒）可多次拉取。
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BIZTOKENEXPIRED = "InvalidParameterValue.BizTokenExpired"
//  INVALIDPARAMETERVALUE_BIZTOKENILLEGAL = "InvalidParameterValue.BizTokenIllegal"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) GetFaceIdResult(request *GetFaceIdResultRequest) (response *GetFaceIdResultResponse, err error) {
    return c.GetFaceIdResultWithContext(context.Background(), request)
}

// GetFaceIdResult
// 完成验证后，用FaceIdToken调用本接口获取结果信息，FaceIdToken生成后三天内（3\*24\*3,600秒）可多次拉取。
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BIZTOKENEXPIRED = "InvalidParameterValue.BizTokenExpired"
//  INVALIDPARAMETERVALUE_BIZTOKENILLEGAL = "InvalidParameterValue.BizTokenIllegal"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) GetFaceIdResultWithContext(ctx context.Context, request *GetFaceIdResultRequest) (response *GetFaceIdResultResponse, err error) {
    if request == nil {
        request = NewGetFaceIdResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetFaceIdResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetFaceIdResultResponse()
    err = c.Send(request, response)
    return
}

func NewGetFaceIdTokenRequest() (request *GetFaceIdTokenRequest) {
    request = &GetFaceIdTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "GetFaceIdToken")
    
    
    return
}

func NewGetFaceIdTokenResponse() (response *GetFaceIdTokenResponse) {
    response = &GetFaceIdTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetFaceIdToken
// 每次调用人脸核身SDK服务前，需先调用本接口获取SDKToken，用来串联核身流程，在验证完成后，用于获取验证结果信息，该token仅能核身一次。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) GetFaceIdToken(request *GetFaceIdTokenRequest) (response *GetFaceIdTokenResponse, err error) {
    return c.GetFaceIdTokenWithContext(context.Background(), request)
}

// GetFaceIdToken
// 每次调用人脸核身SDK服务前，需先调用本接口获取SDKToken，用来串联核身流程，在验证完成后，用于获取验证结果信息，该token仅能核身一次。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) GetFaceIdTokenWithContext(ctx context.Context, request *GetFaceIdTokenRequest) (response *GetFaceIdTokenResponse, err error) {
    if request == nil {
        request = NewGetFaceIdTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetFaceIdToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetFaceIdTokenResponse()
    err = c.Send(request, response)
    return
}

func NewGetLiveCodeRequest() (request *GetLiveCodeRequest) {
    request = &GetLiveCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "GetLiveCode")
    
    
    return
}

func NewGetLiveCodeResponse() (response *GetLiveCodeResponse) {
    response = &GetLiveCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetLiveCode
// 使用数字活体检测模式前，需调用本接口获取数字验证码。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetLiveCode(request *GetLiveCodeRequest) (response *GetLiveCodeResponse, err error) {
    return c.GetLiveCodeWithContext(context.Background(), request)
}

// GetLiveCode
// 使用数字活体检测模式前，需调用本接口获取数字验证码。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetLiveCodeWithContext(ctx context.Context, request *GetLiveCodeRequest) (response *GetLiveCodeResponse, err error) {
    if request == nil {
        request = NewGetLiveCodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetLiveCode require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetLiveCodeResponse()
    err = c.Send(request, response)
    return
}

func NewGetWeChatBillDetailsRequest() (request *GetWeChatBillDetailsRequest) {
    request = &GetWeChatBillDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "GetWeChatBillDetails")
    
    
    return
}

func NewGetWeChatBillDetailsResponse() (response *GetWeChatBillDetailsResponse) {
    response = &GetWeChatBillDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetWeChatBillDetails
// 查询微信渠道服务（微信小程序、微信原生H5、微信普通H5）的账单明细及计费状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_RULEIDNOTEXIST = "InvalidParameterValue.RuleIdNotExist"
func (c *Client) GetWeChatBillDetails(request *GetWeChatBillDetailsRequest) (response *GetWeChatBillDetailsResponse, err error) {
    return c.GetWeChatBillDetailsWithContext(context.Background(), request)
}

// GetWeChatBillDetails
// 查询微信渠道服务（微信小程序、微信原生H5、微信普通H5）的账单明细及计费状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_RULEIDNOTEXIST = "InvalidParameterValue.RuleIdNotExist"
func (c *Client) GetWeChatBillDetailsWithContext(ctx context.Context, request *GetWeChatBillDetailsRequest) (response *GetWeChatBillDetailsResponse, err error) {
    if request == nil {
        request = NewGetWeChatBillDetailsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetWeChatBillDetails require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetWeChatBillDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewIdCardOCRVerificationRequest() (request *IdCardOCRVerificationRequest) {
    request = &IdCardOCRVerificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "IdCardOCRVerification")
    
    
    return
}

func NewIdCardOCRVerificationResponse() (response *IdCardOCRVerificationResponse) {
    response = &IdCardOCRVerificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// IdCardOCRVerification
// 本接口用于校验姓名和身份证号的真实性和一致性，您可以通过输入姓名和身份证号或传入身份证人像面照片提供所需验证信息。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEBLUR = "FailedOperation.ImageBlur"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOIDCARD = "FailedOperation.ImageNoIdCard"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTENCRYPTFIELD = "InvalidParameter.UnsupportEncryptField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_CHARGESTATUSEXCEPTION = "UnauthorizedOperation.ChargeStatusException"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) IdCardOCRVerification(request *IdCardOCRVerificationRequest) (response *IdCardOCRVerificationResponse, err error) {
    return c.IdCardOCRVerificationWithContext(context.Background(), request)
}

// IdCardOCRVerification
// 本接口用于校验姓名和身份证号的真实性和一致性，您可以通过输入姓名和身份证号或传入身份证人像面照片提供所需验证信息。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEBLUR = "FailedOperation.ImageBlur"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOIDCARD = "FailedOperation.ImageNoIdCard"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTENCRYPTFIELD = "InvalidParameter.UnsupportEncryptField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_CHARGESTATUSEXCEPTION = "UnauthorizedOperation.ChargeStatusException"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) IdCardOCRVerificationWithContext(ctx context.Context, request *IdCardOCRVerificationRequest) (response *IdCardOCRVerificationResponse, err error) {
    if request == nil {
        request = NewIdCardOCRVerificationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("IdCardOCRVerification require credential")
    }

    request.SetContext(ctx)
    
    response = NewIdCardOCRVerificationResponse()
    err = c.Send(request, response)
    return
}

func NewIdCardVerificationRequest() (request *IdCardVerificationRequest) {
    request = &IdCardVerificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "IdCardVerification")
    
    
    return
}

func NewIdCardVerificationResponse() (response *IdCardVerificationResponse) {
    response = &IdCardVerificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// IdCardVerification
// 传入姓名和身份证号，校验两者的真实性和一致性。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTENCRYPTFIELD = "InvalidParameter.UnsupportEncryptField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_CHARGESTATUSEXCEPTION = "UnauthorizedOperation.ChargeStatusException"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) IdCardVerification(request *IdCardVerificationRequest) (response *IdCardVerificationResponse, err error) {
    return c.IdCardVerificationWithContext(context.Background(), request)
}

// IdCardVerification
// 传入姓名和身份证号，校验两者的真实性和一致性。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTENCRYPTFIELD = "InvalidParameter.UnsupportEncryptField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_CHARGESTATUSEXCEPTION = "UnauthorizedOperation.ChargeStatusException"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) IdCardVerificationWithContext(ctx context.Context, request *IdCardVerificationRequest) (response *IdCardVerificationResponse, err error) {
    if request == nil {
        request = NewIdCardVerificationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("IdCardVerification require credential")
    }

    request.SetContext(ctx)
    
    response = NewIdCardVerificationResponse()
    err = c.Send(request, response)
    return
}

func NewImageRecognitionRequest() (request *ImageRecognitionRequest) {
    request = &ImageRecognitionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "ImageRecognition")
    
    
    return
}

func NewImageRecognitionResponse() (response *ImageRecognitionResponse) {
    response = &ImageRecognitionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ImageRecognition
// 传入照片和身份信息，判断该照片与权威库的证件照是否属于同一个人。
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPAREFAIL = "FailedOperation.CompareFail"
//  FAILEDOPERATION_COMPARELOWSIMILARITY = "FailedOperation.CompareLowSimilarity"
//  FAILEDOPERATION_COMPARESYSTEMERROR = "FailedOperation.CompareSystemError"
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_IDFORMATERROR = "FailedOperation.IdFormatError"
//  FAILEDOPERATION_IDNAMEMISMATCH = "FailedOperation.IdNameMisMatch"
//  FAILEDOPERATION_IDNOEXISTSYSTEM = "FailedOperation.IdNoExistSystem"
//  FAILEDOPERATION_IDPHOTONOEXIST = "FailedOperation.IdPhotoNoExist"
//  FAILEDOPERATION_IDPHOTOPOORQUALITY = "FailedOperation.IdPhotoPoorQuality"
//  FAILEDOPERATION_IDPHOTOSYSTEMNOANSWER = "FailedOperation.IdPhotoSystemNoanswer"
//  FAILEDOPERATION_LIFEPHOTODETECTFACES = "FailedOperation.LifePhotoDetectFaces"
//  FAILEDOPERATION_LIFEPHOTODETECTFAKE = "FailedOperation.LifePhotoDetectFake"
//  FAILEDOPERATION_LIFEPHOTODETECTNOFACES = "FailedOperation.LifePhotoDetectNoFaces"
//  FAILEDOPERATION_LIFEPHOTOPOORQUALITY = "FailedOperation.LifePhotoPoorQuality"
//  FAILEDOPERATION_LIFEPHOTOSIZEERROR = "FailedOperation.LifePhotoSizeError"
//  FAILEDOPERATION_NAMEFORMATERROR = "FailedOperation.NameFormatError"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTENCRYPTFIELD = "InvalidParameter.UnsupportEncryptField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ImageRecognition(request *ImageRecognitionRequest) (response *ImageRecognitionResponse, err error) {
    return c.ImageRecognitionWithContext(context.Background(), request)
}

// ImageRecognition
// 传入照片和身份信息，判断该照片与权威库的证件照是否属于同一个人。
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMPAREFAIL = "FailedOperation.CompareFail"
//  FAILEDOPERATION_COMPARELOWSIMILARITY = "FailedOperation.CompareLowSimilarity"
//  FAILEDOPERATION_COMPARESYSTEMERROR = "FailedOperation.CompareSystemError"
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_IDFORMATERROR = "FailedOperation.IdFormatError"
//  FAILEDOPERATION_IDNAMEMISMATCH = "FailedOperation.IdNameMisMatch"
//  FAILEDOPERATION_IDNOEXISTSYSTEM = "FailedOperation.IdNoExistSystem"
//  FAILEDOPERATION_IDPHOTONOEXIST = "FailedOperation.IdPhotoNoExist"
//  FAILEDOPERATION_IDPHOTOPOORQUALITY = "FailedOperation.IdPhotoPoorQuality"
//  FAILEDOPERATION_IDPHOTOSYSTEMNOANSWER = "FailedOperation.IdPhotoSystemNoanswer"
//  FAILEDOPERATION_LIFEPHOTODETECTFACES = "FailedOperation.LifePhotoDetectFaces"
//  FAILEDOPERATION_LIFEPHOTODETECTFAKE = "FailedOperation.LifePhotoDetectFake"
//  FAILEDOPERATION_LIFEPHOTODETECTNOFACES = "FailedOperation.LifePhotoDetectNoFaces"
//  FAILEDOPERATION_LIFEPHOTOPOORQUALITY = "FailedOperation.LifePhotoPoorQuality"
//  FAILEDOPERATION_LIFEPHOTOSIZEERROR = "FailedOperation.LifePhotoSizeError"
//  FAILEDOPERATION_NAMEFORMATERROR = "FailedOperation.NameFormatError"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTENCRYPTFIELD = "InvalidParameter.UnsupportEncryptField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ImageRecognitionWithContext(ctx context.Context, request *ImageRecognitionRequest) (response *ImageRecognitionResponse, err error) {
    if request == nil {
        request = NewImageRecognitionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImageRecognition require credential")
    }

    request.SetContext(ctx)
    
    response = NewImageRecognitionResponse()
    err = c.Send(request, response)
    return
}

func NewLivenessRequest() (request *LivenessRequest) {
    request = &LivenessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "Liveness")
    
    
    return
}

func NewLivenessResponse() (response *LivenessResponse) {
    response = &LivenessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Liveness
// 活体检测
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTIONCLOSEEYE = "FailedOperation.ActionCloseEye"
//  FAILEDOPERATION_ACTIONFACECLOSE = "FailedOperation.ActionFaceClose"
//  FAILEDOPERATION_ACTIONFACEFAR = "FailedOperation.ActionFaceFar"
//  FAILEDOPERATION_ACTIONFACELEFT = "FailedOperation.ActionFaceLeft"
//  FAILEDOPERATION_ACTIONFACERIGHT = "FailedOperation.ActionFaceRight"
//  FAILEDOPERATION_ACTIONFIRSTACTION = "FailedOperation.ActionFirstAction"
//  FAILEDOPERATION_ACTIONLIGHTDARK = "FailedOperation.ActionLightDark"
//  FAILEDOPERATION_ACTIONLIGHTSTRONG = "FailedOperation.ActionLightStrong"
//  FAILEDOPERATION_ACTIONNODETECTFACE = "FailedOperation.ActionNodetectFace"
//  FAILEDOPERATION_ACTIONOPENMOUTH = "FailedOperation.ActionOpenMouth"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DbError"
//  FAILEDOPERATION_FILESAVEERROR = "FailedOperation.FileSaveError"
//  FAILEDOPERATION_LIPFACEINCOMPLETE = "FailedOperation.LipFaceIncomplete"
//  FAILEDOPERATION_LIPMOVESMALL = "FailedOperation.LipMoveSmall"
//  FAILEDOPERATION_LIPNETFAILED = "FailedOperation.LipNetFailed"
//  FAILEDOPERATION_LIPSIZEERROR = "FailedOperation.LipSizeError"
//  FAILEDOPERATION_LIPVIDEOINVALID = "FailedOperation.LipVideoInvalid"
//  FAILEDOPERATION_LIPVIDEOQUAILITY = "FailedOperation.LipVideoQuaility"
//  FAILEDOPERATION_LIPVOICEDETECT = "FailedOperation.LipVoiceDetect"
//  FAILEDOPERATION_LIPVOICELOW = "FailedOperation.LipVoiceLow"
//  FAILEDOPERATION_LIPVOICERECOGNIZE = "FailedOperation.LipVoiceRecognize"
//  FAILEDOPERATION_LIVESSBESTFRAMEERROR = "FailedOperation.LivessBestFrameError"
//  FAILEDOPERATION_LIVESSDETECTFAIL = "FailedOperation.LivessDetectFail"
//  FAILEDOPERATION_LIVESSDETECTFAKE = "FailedOperation.LivessDetectFake"
//  FAILEDOPERATION_LIVESSSYSTEMERROR = "FailedOperation.LivessSystemError"
//  FAILEDOPERATION_LIVESSUNKNOWNERROR = "FailedOperation.LivessUnknownError"
//  FAILEDOPERATION_SILENTDETECTFAIL = "FailedOperation.SilentDetectFail"
//  FAILEDOPERATION_SILENTTHRESHOLD = "FailedOperation.SilentThreshold"
//  FAILEDOPERATION_SILENTTOOSHORT = "FailedOperation.SilentTooShort"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  FAILEDOPERATION_VERIFICATIONFAIL = "FailedOperation.VerificationFail"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWN = "InternalError.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RULEID = "InvalidParameter.RuleId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BIZTOKENEXPIRED = "InvalidParameterValue.BizTokenExpired"
//  INVALIDPARAMETERVALUE_RULEIDNOTEXIST = "InvalidParameterValue.RuleIdNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) Liveness(request *LivenessRequest) (response *LivenessResponse, err error) {
    return c.LivenessWithContext(context.Background(), request)
}

// Liveness
// 活体检测
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACTIONCLOSEEYE = "FailedOperation.ActionCloseEye"
//  FAILEDOPERATION_ACTIONFACECLOSE = "FailedOperation.ActionFaceClose"
//  FAILEDOPERATION_ACTIONFACEFAR = "FailedOperation.ActionFaceFar"
//  FAILEDOPERATION_ACTIONFACELEFT = "FailedOperation.ActionFaceLeft"
//  FAILEDOPERATION_ACTIONFACERIGHT = "FailedOperation.ActionFaceRight"
//  FAILEDOPERATION_ACTIONFIRSTACTION = "FailedOperation.ActionFirstAction"
//  FAILEDOPERATION_ACTIONLIGHTDARK = "FailedOperation.ActionLightDark"
//  FAILEDOPERATION_ACTIONLIGHTSTRONG = "FailedOperation.ActionLightStrong"
//  FAILEDOPERATION_ACTIONNODETECTFACE = "FailedOperation.ActionNodetectFace"
//  FAILEDOPERATION_ACTIONOPENMOUTH = "FailedOperation.ActionOpenMouth"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DbError"
//  FAILEDOPERATION_FILESAVEERROR = "FailedOperation.FileSaveError"
//  FAILEDOPERATION_LIPFACEINCOMPLETE = "FailedOperation.LipFaceIncomplete"
//  FAILEDOPERATION_LIPMOVESMALL = "FailedOperation.LipMoveSmall"
//  FAILEDOPERATION_LIPNETFAILED = "FailedOperation.LipNetFailed"
//  FAILEDOPERATION_LIPSIZEERROR = "FailedOperation.LipSizeError"
//  FAILEDOPERATION_LIPVIDEOINVALID = "FailedOperation.LipVideoInvalid"
//  FAILEDOPERATION_LIPVIDEOQUAILITY = "FailedOperation.LipVideoQuaility"
//  FAILEDOPERATION_LIPVOICEDETECT = "FailedOperation.LipVoiceDetect"
//  FAILEDOPERATION_LIPVOICELOW = "FailedOperation.LipVoiceLow"
//  FAILEDOPERATION_LIPVOICERECOGNIZE = "FailedOperation.LipVoiceRecognize"
//  FAILEDOPERATION_LIVESSBESTFRAMEERROR = "FailedOperation.LivessBestFrameError"
//  FAILEDOPERATION_LIVESSDETECTFAIL = "FailedOperation.LivessDetectFail"
//  FAILEDOPERATION_LIVESSDETECTFAKE = "FailedOperation.LivessDetectFake"
//  FAILEDOPERATION_LIVESSSYSTEMERROR = "FailedOperation.LivessSystemError"
//  FAILEDOPERATION_LIVESSUNKNOWNERROR = "FailedOperation.LivessUnknownError"
//  FAILEDOPERATION_SILENTDETECTFAIL = "FailedOperation.SilentDetectFail"
//  FAILEDOPERATION_SILENTTHRESHOLD = "FailedOperation.SilentThreshold"
//  FAILEDOPERATION_SILENTTOOSHORT = "FailedOperation.SilentTooShort"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  FAILEDOPERATION_VERIFICATIONFAIL = "FailedOperation.VerificationFail"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWN = "InternalError.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RULEID = "InvalidParameter.RuleId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BIZTOKENEXPIRED = "InvalidParameterValue.BizTokenExpired"
//  INVALIDPARAMETERVALUE_RULEIDNOTEXIST = "InvalidParameterValue.RuleIdNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) LivenessWithContext(ctx context.Context, request *LivenessRequest) (response *LivenessResponse, err error) {
    if request == nil {
        request = NewLivenessRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("Liveness require credential")
    }

    request.SetContext(ctx)
    
    response = NewLivenessResponse()
    err = c.Send(request, response)
    return
}

func NewLivenessCompareRequest() (request *LivenessCompareRequest) {
    request = &LivenessCompareRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "LivenessCompare")
    
    
    return
}

func NewLivenessCompareResponse() (response *LivenessCompareResponse) {
    response = &LivenessCompareResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// LivenessCompare
// 传入视频和照片，先判断视频中是否为真人，判断为真人后，再判断该视频中的人与上传照片是否属于同一个人。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_ACTIONCLOSEEYE = "FailedOperation.ActionCloseEye"
//  FAILEDOPERATION_ACTIONFACECLOSE = "FailedOperation.ActionFaceClose"
//  FAILEDOPERATION_ACTIONFACEFAR = "FailedOperation.ActionFaceFar"
//  FAILEDOPERATION_ACTIONFACELEFT = "FailedOperation.ActionFaceLeft"
//  FAILEDOPERATION_ACTIONFACERIGHT = "FailedOperation.ActionFaceRight"
//  FAILEDOPERATION_ACTIONFIRSTACTION = "FailedOperation.ActionFirstAction"
//  FAILEDOPERATION_ACTIONLIGHTDARK = "FailedOperation.ActionLightDark"
//  FAILEDOPERATION_ACTIONLIGHTSTRONG = "FailedOperation.ActionLightStrong"
//  FAILEDOPERATION_ACTIONNODETECTFACE = "FailedOperation.ActionNodetectFace"
//  FAILEDOPERATION_ACTIONOPENMOUTH = "FailedOperation.ActionOpenMouth"
//  FAILEDOPERATION_COMPAREFAIL = "FailedOperation.CompareFail"
//  FAILEDOPERATION_COMPARELOWSIMILARITY = "FailedOperation.CompareLowSimilarity"
//  FAILEDOPERATION_COMPARESYSTEMERROR = "FailedOperation.CompareSystemError"
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_DOWNLOADTIMEOUTERROR = "FailedOperation.DownLoadTimeoutError"
//  FAILEDOPERATION_LIFEPHOTODETECTFACES = "FailedOperation.LifePhotoDetectFaces"
//  FAILEDOPERATION_LIFEPHOTODETECTFAKE = "FailedOperation.LifePhotoDetectFake"
//  FAILEDOPERATION_LIFEPHOTODETECTNOFACES = "FailedOperation.LifePhotoDetectNoFaces"
//  FAILEDOPERATION_LIFEPHOTOPOORQUALITY = "FailedOperation.LifePhotoPoorQuality"
//  FAILEDOPERATION_LIFEPHOTOSIZEERROR = "FailedOperation.LifePhotoSizeError"
//  FAILEDOPERATION_LIPFACEINCOMPLETE = "FailedOperation.LipFaceIncomplete"
//  FAILEDOPERATION_LIPMOVESMALL = "FailedOperation.LipMoveSmall"
//  FAILEDOPERATION_LIPNETFAILED = "FailedOperation.LipNetFailed"
//  FAILEDOPERATION_LIPSIZEERROR = "FailedOperation.LipSizeError"
//  FAILEDOPERATION_LIPVIDEOINVALID = "FailedOperation.LipVideoInvalid"
//  FAILEDOPERATION_LIPVIDEOQUAILITY = "FailedOperation.LipVideoQuaility"
//  FAILEDOPERATION_LIPVOICEDETECT = "FailedOperation.LipVoiceDetect"
//  FAILEDOPERATION_LIPVOICELOW = "FailedOperation.LipVoiceLow"
//  FAILEDOPERATION_LIPVOICERECOGNIZE = "FailedOperation.LipVoiceRecognize"
//  FAILEDOPERATION_LIVESSBESTFRAMEERROR = "FailedOperation.LivessBestFrameError"
//  FAILEDOPERATION_LIVESSDETECTFAIL = "FailedOperation.LivessDetectFail"
//  FAILEDOPERATION_LIVESSDETECTFAKE = "FailedOperation.LivessDetectFake"
//  FAILEDOPERATION_LIVESSSYSTEMERROR = "FailedOperation.LivessSystemError"
//  FAILEDOPERATION_LIVESSUNKNOWNERROR = "FailedOperation.LivessUnknownError"
//  FAILEDOPERATION_SILENTDETECTFAIL = "FailedOperation.SilentDetectFail"
//  FAILEDOPERATION_SILENTEYELIVEFAIL = "FailedOperation.SilentEyeLiveFail"
//  FAILEDOPERATION_SILENTFACEDETECTFAIL = "FailedOperation.SilentFaceDetectFail"
//  FAILEDOPERATION_SILENTFACEQUALITYFAIL = "FailedOperation.SilentFaceQualityFail"
//  FAILEDOPERATION_SILENTFACEWITHMASKFAIL = "FailedOperation.SilentFaceWithMaskFail"
//  FAILEDOPERATION_SILENTMOUTHLIVEFAIL = "FailedOperation.SilentMouthLiveFail"
//  FAILEDOPERATION_SILENTMULTIFACEFAIL = "FailedOperation.SilentMultiFaceFail"
//  FAILEDOPERATION_SILENTPICTURELIVEFAIL = "FailedOperation.SilentPictureLiveFail"
//  FAILEDOPERATION_SILENTTHRESHOLD = "FailedOperation.SilentThreshold"
//  FAILEDOPERATION_SILENTTOOSHORT = "FailedOperation.SilentTooShort"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) LivenessCompare(request *LivenessCompareRequest) (response *LivenessCompareResponse, err error) {
    return c.LivenessCompareWithContext(context.Background(), request)
}

// LivenessCompare
// 传入视频和照片，先判断视频中是否为真人，判断为真人后，再判断该视频中的人与上传照片是否属于同一个人。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_ACTIONCLOSEEYE = "FailedOperation.ActionCloseEye"
//  FAILEDOPERATION_ACTIONFACECLOSE = "FailedOperation.ActionFaceClose"
//  FAILEDOPERATION_ACTIONFACEFAR = "FailedOperation.ActionFaceFar"
//  FAILEDOPERATION_ACTIONFACELEFT = "FailedOperation.ActionFaceLeft"
//  FAILEDOPERATION_ACTIONFACERIGHT = "FailedOperation.ActionFaceRight"
//  FAILEDOPERATION_ACTIONFIRSTACTION = "FailedOperation.ActionFirstAction"
//  FAILEDOPERATION_ACTIONLIGHTDARK = "FailedOperation.ActionLightDark"
//  FAILEDOPERATION_ACTIONLIGHTSTRONG = "FailedOperation.ActionLightStrong"
//  FAILEDOPERATION_ACTIONNODETECTFACE = "FailedOperation.ActionNodetectFace"
//  FAILEDOPERATION_ACTIONOPENMOUTH = "FailedOperation.ActionOpenMouth"
//  FAILEDOPERATION_COMPAREFAIL = "FailedOperation.CompareFail"
//  FAILEDOPERATION_COMPARELOWSIMILARITY = "FailedOperation.CompareLowSimilarity"
//  FAILEDOPERATION_COMPARESYSTEMERROR = "FailedOperation.CompareSystemError"
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_DOWNLOADTIMEOUTERROR = "FailedOperation.DownLoadTimeoutError"
//  FAILEDOPERATION_LIFEPHOTODETECTFACES = "FailedOperation.LifePhotoDetectFaces"
//  FAILEDOPERATION_LIFEPHOTODETECTFAKE = "FailedOperation.LifePhotoDetectFake"
//  FAILEDOPERATION_LIFEPHOTODETECTNOFACES = "FailedOperation.LifePhotoDetectNoFaces"
//  FAILEDOPERATION_LIFEPHOTOPOORQUALITY = "FailedOperation.LifePhotoPoorQuality"
//  FAILEDOPERATION_LIFEPHOTOSIZEERROR = "FailedOperation.LifePhotoSizeError"
//  FAILEDOPERATION_LIPFACEINCOMPLETE = "FailedOperation.LipFaceIncomplete"
//  FAILEDOPERATION_LIPMOVESMALL = "FailedOperation.LipMoveSmall"
//  FAILEDOPERATION_LIPNETFAILED = "FailedOperation.LipNetFailed"
//  FAILEDOPERATION_LIPSIZEERROR = "FailedOperation.LipSizeError"
//  FAILEDOPERATION_LIPVIDEOINVALID = "FailedOperation.LipVideoInvalid"
//  FAILEDOPERATION_LIPVIDEOQUAILITY = "FailedOperation.LipVideoQuaility"
//  FAILEDOPERATION_LIPVOICEDETECT = "FailedOperation.LipVoiceDetect"
//  FAILEDOPERATION_LIPVOICELOW = "FailedOperation.LipVoiceLow"
//  FAILEDOPERATION_LIPVOICERECOGNIZE = "FailedOperation.LipVoiceRecognize"
//  FAILEDOPERATION_LIVESSBESTFRAMEERROR = "FailedOperation.LivessBestFrameError"
//  FAILEDOPERATION_LIVESSDETECTFAIL = "FailedOperation.LivessDetectFail"
//  FAILEDOPERATION_LIVESSDETECTFAKE = "FailedOperation.LivessDetectFake"
//  FAILEDOPERATION_LIVESSSYSTEMERROR = "FailedOperation.LivessSystemError"
//  FAILEDOPERATION_LIVESSUNKNOWNERROR = "FailedOperation.LivessUnknownError"
//  FAILEDOPERATION_SILENTDETECTFAIL = "FailedOperation.SilentDetectFail"
//  FAILEDOPERATION_SILENTEYELIVEFAIL = "FailedOperation.SilentEyeLiveFail"
//  FAILEDOPERATION_SILENTFACEDETECTFAIL = "FailedOperation.SilentFaceDetectFail"
//  FAILEDOPERATION_SILENTFACEQUALITYFAIL = "FailedOperation.SilentFaceQualityFail"
//  FAILEDOPERATION_SILENTFACEWITHMASKFAIL = "FailedOperation.SilentFaceWithMaskFail"
//  FAILEDOPERATION_SILENTMOUTHLIVEFAIL = "FailedOperation.SilentMouthLiveFail"
//  FAILEDOPERATION_SILENTMULTIFACEFAIL = "FailedOperation.SilentMultiFaceFail"
//  FAILEDOPERATION_SILENTPICTURELIVEFAIL = "FailedOperation.SilentPictureLiveFail"
//  FAILEDOPERATION_SILENTTHRESHOLD = "FailedOperation.SilentThreshold"
//  FAILEDOPERATION_SILENTTOOSHORT = "FailedOperation.SilentTooShort"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) LivenessCompareWithContext(ctx context.Context, request *LivenessCompareRequest) (response *LivenessCompareResponse, err error) {
    if request == nil {
        request = NewLivenessCompareRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("LivenessCompare require credential")
    }

    request.SetContext(ctx)
    
    response = NewLivenessCompareResponse()
    err = c.Send(request, response)
    return
}

func NewLivenessRecognitionRequest() (request *LivenessRecognitionRequest) {
    request = &LivenessRecognitionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "LivenessRecognition")
    
    
    return
}

func NewLivenessRecognitionResponse() (response *LivenessRecognitionResponse) {
    response = &LivenessRecognitionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// LivenessRecognition
// 传入视频和身份信息，先判断视频中是否为真人，判断为真人后，再判断该视频中的人与权威库的证件照是否属于同一个人。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_ACTIONCLOSEEYE = "FailedOperation.ActionCloseEye"
//  FAILEDOPERATION_ACTIONFACECLOSE = "FailedOperation.ActionFaceClose"
//  FAILEDOPERATION_ACTIONFACEFAR = "FailedOperation.ActionFaceFar"
//  FAILEDOPERATION_ACTIONFACELEFT = "FailedOperation.ActionFaceLeft"
//  FAILEDOPERATION_ACTIONFACERIGHT = "FailedOperation.ActionFaceRight"
//  FAILEDOPERATION_ACTIONFIRSTACTION = "FailedOperation.ActionFirstAction"
//  FAILEDOPERATION_ACTIONLIGHTDARK = "FailedOperation.ActionLightDark"
//  FAILEDOPERATION_ACTIONLIGHTSTRONG = "FailedOperation.ActionLightStrong"
//  FAILEDOPERATION_ACTIONNODETECTFACE = "FailedOperation.ActionNodetectFace"
//  FAILEDOPERATION_ACTIONOPENMOUTH = "FailedOperation.ActionOpenMouth"
//  FAILEDOPERATION_COMPAREFAIL = "FailedOperation.CompareFail"
//  FAILEDOPERATION_COMPARELOWSIMILARITY = "FailedOperation.CompareLowSimilarity"
//  FAILEDOPERATION_COMPARESYSTEMERROR = "FailedOperation.CompareSystemError"
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_DOWNLOADTIMEOUTERROR = "FailedOperation.DownLoadTimeoutError"
//  FAILEDOPERATION_IDFORMATERROR = "FailedOperation.IdFormatError"
//  FAILEDOPERATION_IDNAMEMISMATCH = "FailedOperation.IdNameMisMatch"
//  FAILEDOPERATION_IDNOEXISTSYSTEM = "FailedOperation.IdNoExistSystem"
//  FAILEDOPERATION_IDPHOTONOEXIST = "FailedOperation.IdPhotoNoExist"
//  FAILEDOPERATION_IDPHOTOPOORQUALITY = "FailedOperation.IdPhotoPoorQuality"
//  FAILEDOPERATION_LIFEPHOTODETECTFACES = "FailedOperation.LifePhotoDetectFaces"
//  FAILEDOPERATION_LIFEPHOTODETECTFAKE = "FailedOperation.LifePhotoDetectFake"
//  FAILEDOPERATION_LIFEPHOTODETECTNOFACES = "FailedOperation.LifePhotoDetectNoFaces"
//  FAILEDOPERATION_LIFEPHOTOSIZEERROR = "FailedOperation.LifePhotoSizeError"
//  FAILEDOPERATION_LIPFACEINCOMPLETE = "FailedOperation.LipFaceIncomplete"
//  FAILEDOPERATION_LIPMOVESMALL = "FailedOperation.LipMoveSmall"
//  FAILEDOPERATION_LIPNETFAILED = "FailedOperation.LipNetFailed"
//  FAILEDOPERATION_LIPSIZEERROR = "FailedOperation.LipSizeError"
//  FAILEDOPERATION_LIPVIDEOINVALID = "FailedOperation.LipVideoInvalid"
//  FAILEDOPERATION_LIPVIDEOQUAILITY = "FailedOperation.LipVideoQuaility"
//  FAILEDOPERATION_LIPVOICEDETECT = "FailedOperation.LipVoiceDetect"
//  FAILEDOPERATION_LIPVOICELOW = "FailedOperation.LipVoiceLow"
//  FAILEDOPERATION_LIPVOICERECOGNIZE = "FailedOperation.LipVoiceRecognize"
//  FAILEDOPERATION_LIVESSBESTFRAMEERROR = "FailedOperation.LivessBestFrameError"
//  FAILEDOPERATION_LIVESSDETECTFAIL = "FailedOperation.LivessDetectFail"
//  FAILEDOPERATION_LIVESSDETECTFAKE = "FailedOperation.LivessDetectFake"
//  FAILEDOPERATION_LIVESSSYSTEMERROR = "FailedOperation.LivessSystemError"
//  FAILEDOPERATION_LIVESSUNKNOWNERROR = "FailedOperation.LivessUnknownError"
//  FAILEDOPERATION_NAMEFORMATERROR = "FailedOperation.NameFormatError"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_SILENTDETECTFAIL = "FailedOperation.SilentDetectFail"
//  FAILEDOPERATION_SILENTTHRESHOLD = "FailedOperation.SilentThreshold"
//  FAILEDOPERATION_SILENTTOOSHORT = "FailedOperation.SilentTooShort"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTENCRYPTFIELD = "InvalidParameter.UnsupportEncryptField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) LivenessRecognition(request *LivenessRecognitionRequest) (response *LivenessRecognitionResponse, err error) {
    return c.LivenessRecognitionWithContext(context.Background(), request)
}

// LivenessRecognition
// 传入视频和身份信息，先判断视频中是否为真人，判断为真人后，再判断该视频中的人与权威库的证件照是否属于同一个人。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_ACTIONCLOSEEYE = "FailedOperation.ActionCloseEye"
//  FAILEDOPERATION_ACTIONFACECLOSE = "FailedOperation.ActionFaceClose"
//  FAILEDOPERATION_ACTIONFACEFAR = "FailedOperation.ActionFaceFar"
//  FAILEDOPERATION_ACTIONFACELEFT = "FailedOperation.ActionFaceLeft"
//  FAILEDOPERATION_ACTIONFACERIGHT = "FailedOperation.ActionFaceRight"
//  FAILEDOPERATION_ACTIONFIRSTACTION = "FailedOperation.ActionFirstAction"
//  FAILEDOPERATION_ACTIONLIGHTDARK = "FailedOperation.ActionLightDark"
//  FAILEDOPERATION_ACTIONLIGHTSTRONG = "FailedOperation.ActionLightStrong"
//  FAILEDOPERATION_ACTIONNODETECTFACE = "FailedOperation.ActionNodetectFace"
//  FAILEDOPERATION_ACTIONOPENMOUTH = "FailedOperation.ActionOpenMouth"
//  FAILEDOPERATION_COMPAREFAIL = "FailedOperation.CompareFail"
//  FAILEDOPERATION_COMPARELOWSIMILARITY = "FailedOperation.CompareLowSimilarity"
//  FAILEDOPERATION_COMPARESYSTEMERROR = "FailedOperation.CompareSystemError"
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_DOWNLOADTIMEOUTERROR = "FailedOperation.DownLoadTimeoutError"
//  FAILEDOPERATION_IDFORMATERROR = "FailedOperation.IdFormatError"
//  FAILEDOPERATION_IDNAMEMISMATCH = "FailedOperation.IdNameMisMatch"
//  FAILEDOPERATION_IDNOEXISTSYSTEM = "FailedOperation.IdNoExistSystem"
//  FAILEDOPERATION_IDPHOTONOEXIST = "FailedOperation.IdPhotoNoExist"
//  FAILEDOPERATION_IDPHOTOPOORQUALITY = "FailedOperation.IdPhotoPoorQuality"
//  FAILEDOPERATION_LIFEPHOTODETECTFACES = "FailedOperation.LifePhotoDetectFaces"
//  FAILEDOPERATION_LIFEPHOTODETECTFAKE = "FailedOperation.LifePhotoDetectFake"
//  FAILEDOPERATION_LIFEPHOTODETECTNOFACES = "FailedOperation.LifePhotoDetectNoFaces"
//  FAILEDOPERATION_LIFEPHOTOSIZEERROR = "FailedOperation.LifePhotoSizeError"
//  FAILEDOPERATION_LIPFACEINCOMPLETE = "FailedOperation.LipFaceIncomplete"
//  FAILEDOPERATION_LIPMOVESMALL = "FailedOperation.LipMoveSmall"
//  FAILEDOPERATION_LIPNETFAILED = "FailedOperation.LipNetFailed"
//  FAILEDOPERATION_LIPSIZEERROR = "FailedOperation.LipSizeError"
//  FAILEDOPERATION_LIPVIDEOINVALID = "FailedOperation.LipVideoInvalid"
//  FAILEDOPERATION_LIPVIDEOQUAILITY = "FailedOperation.LipVideoQuaility"
//  FAILEDOPERATION_LIPVOICEDETECT = "FailedOperation.LipVoiceDetect"
//  FAILEDOPERATION_LIPVOICELOW = "FailedOperation.LipVoiceLow"
//  FAILEDOPERATION_LIPVOICERECOGNIZE = "FailedOperation.LipVoiceRecognize"
//  FAILEDOPERATION_LIVESSBESTFRAMEERROR = "FailedOperation.LivessBestFrameError"
//  FAILEDOPERATION_LIVESSDETECTFAIL = "FailedOperation.LivessDetectFail"
//  FAILEDOPERATION_LIVESSDETECTFAKE = "FailedOperation.LivessDetectFake"
//  FAILEDOPERATION_LIVESSSYSTEMERROR = "FailedOperation.LivessSystemError"
//  FAILEDOPERATION_LIVESSUNKNOWNERROR = "FailedOperation.LivessUnknownError"
//  FAILEDOPERATION_NAMEFORMATERROR = "FailedOperation.NameFormatError"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_SILENTDETECTFAIL = "FailedOperation.SilentDetectFail"
//  FAILEDOPERATION_SILENTTHRESHOLD = "FailedOperation.SilentThreshold"
//  FAILEDOPERATION_SILENTTOOSHORT = "FailedOperation.SilentTooShort"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTENCRYPTFIELD = "InvalidParameter.UnsupportEncryptField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) LivenessRecognitionWithContext(ctx context.Context, request *LivenessRecognitionRequest) (response *LivenessRecognitionResponse, err error) {
    if request == nil {
        request = NewLivenessRecognitionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("LivenessRecognition require credential")
    }

    request.SetContext(ctx)
    
    response = NewLivenessRecognitionResponse()
    err = c.Send(request, response)
    return
}

func NewMinorsVerificationRequest() (request *MinorsVerificationRequest) {
    request = &MinorsVerificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "MinorsVerification")
    
    
    return
}

func NewMinorsVerificationResponse() (response *MinorsVerificationResponse) {
    response = &MinorsVerificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// MinorsVerification
// 通过传入手机号或姓名和身份证号，结合权威数据源和腾讯健康守护可信模型，判断该信息是否真实且年满18周岁。腾讯健康守护可信模型覆盖了上十亿手机库源，覆盖率高、准确率高，如果不在库中的手机号，还可以通过姓名+身份证进行兜底验证。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTENCRYPTFIELD = "InvalidParameter.UnsupportEncryptField"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_CHARGESTATUSEXCEPTION = "UnauthorizedOperation.ChargeStatusException"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) MinorsVerification(request *MinorsVerificationRequest) (response *MinorsVerificationResponse, err error) {
    return c.MinorsVerificationWithContext(context.Background(), request)
}

// MinorsVerification
// 通过传入手机号或姓名和身份证号，结合权威数据源和腾讯健康守护可信模型，判断该信息是否真实且年满18周岁。腾讯健康守护可信模型覆盖了上十亿手机库源，覆盖率高、准确率高，如果不在库中的手机号，还可以通过姓名+身份证进行兜底验证。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTENCRYPTFIELD = "InvalidParameter.UnsupportEncryptField"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_CHARGESTATUSEXCEPTION = "UnauthorizedOperation.ChargeStatusException"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) MinorsVerificationWithContext(ctx context.Context, request *MinorsVerificationRequest) (response *MinorsVerificationResponse, err error) {
    if request == nil {
        request = NewMinorsVerificationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("MinorsVerification require credential")
    }

    request.SetContext(ctx)
    
    response = NewMinorsVerificationResponse()
    err = c.Send(request, response)
    return
}

func NewMobileNetworkTimeVerificationRequest() (request *MobileNetworkTimeVerificationRequest) {
    request = &MobileNetworkTimeVerificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "MobileNetworkTimeVerification")
    
    
    return
}

func NewMobileNetworkTimeVerificationResponse() (response *MobileNetworkTimeVerificationResponse) {
    response = &MobileNetworkTimeVerificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// MobileNetworkTimeVerification
// 本接口用于查询手机号在网时长，输入手机号进行查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTENCRYPTFIELD = "InvalidParameter.UnsupportEncryptField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BIZTOKENEXPIRED = "InvalidParameterValue.BizTokenExpired"
//  INVALIDPARAMETERVALUE_RULEIDNOTEXIST = "InvalidParameterValue.RuleIdNotExist"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_CHARGESTATUSEXCEPTION = "UnauthorizedOperation.ChargeStatusException"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) MobileNetworkTimeVerification(request *MobileNetworkTimeVerificationRequest) (response *MobileNetworkTimeVerificationResponse, err error) {
    return c.MobileNetworkTimeVerificationWithContext(context.Background(), request)
}

// MobileNetworkTimeVerification
// 本接口用于查询手机号在网时长，输入手机号进行查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTENCRYPTFIELD = "InvalidParameter.UnsupportEncryptField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BIZTOKENEXPIRED = "InvalidParameterValue.BizTokenExpired"
//  INVALIDPARAMETERVALUE_RULEIDNOTEXIST = "InvalidParameterValue.RuleIdNotExist"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_CHARGESTATUSEXCEPTION = "UnauthorizedOperation.ChargeStatusException"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) MobileNetworkTimeVerificationWithContext(ctx context.Context, request *MobileNetworkTimeVerificationRequest) (response *MobileNetworkTimeVerificationResponse, err error) {
    if request == nil {
        request = NewMobileNetworkTimeVerificationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("MobileNetworkTimeVerification require credential")
    }

    request.SetContext(ctx)
    
    response = NewMobileNetworkTimeVerificationResponse()
    err = c.Send(request, response)
    return
}

func NewMobileStatusRequest() (request *MobileStatusRequest) {
    request = &MobileStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "MobileStatus")
    
    
    return
}

func NewMobileStatusResponse() (response *MobileStatusResponse) {
    response = &MobileStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// MobileStatus
// 本接口用于验证手机号的状态，您可以输入手机号进行查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTENCRYPTFIELD = "InvalidParameter.UnsupportEncryptField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_CHARGESTATUSEXCEPTION = "UnauthorizedOperation.ChargeStatusException"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) MobileStatus(request *MobileStatusRequest) (response *MobileStatusResponse, err error) {
    return c.MobileStatusWithContext(context.Background(), request)
}

// MobileStatus
// 本接口用于验证手机号的状态，您可以输入手机号进行查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTENCRYPTFIELD = "InvalidParameter.UnsupportEncryptField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_CHARGESTATUSEXCEPTION = "UnauthorizedOperation.ChargeStatusException"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) MobileStatusWithContext(ctx context.Context, request *MobileStatusRequest) (response *MobileStatusResponse, err error) {
    if request == nil {
        request = NewMobileStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("MobileStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewMobileStatusResponse()
    err = c.Send(request, response)
    return
}

func NewParseNfcDataRequest() (request *ParseNfcDataRequest) {
    request = &ParseNfcDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "ParseNfcData")
    
    
    return
}

func NewParseNfcDataResponse() (response *ParseNfcDataResponse) {
    response = &ParseNfcDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ParseNfcData
// 解析SDK获取到的证件NFC数据，接口传入SDK返回的ReqId，返回证件信息（个别字段为特定证件类型特有）。SDK生成的ReqId五分钟内有效，重复查询仅收一次费。支持身份证类证件（二代身份证、港澳居住证、台湾居住证、外国人永居证）以及旅行类证件（港澳通行证、台湾通行证、台胞证、回乡证）的NFC识别及核验。
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.UnKnown"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) ParseNfcData(request *ParseNfcDataRequest) (response *ParseNfcDataResponse, err error) {
    return c.ParseNfcDataWithContext(context.Background(), request)
}

// ParseNfcData
// 解析SDK获取到的证件NFC数据，接口传入SDK返回的ReqId，返回证件信息（个别字段为特定证件类型特有）。SDK生成的ReqId五分钟内有效，重复查询仅收一次费。支持身份证类证件（二代身份证、港澳居住证、台湾居住证、外国人永居证）以及旅行类证件（港澳通行证、台湾通行证、台胞证、回乡证）的NFC识别及核验。
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWN = "InternalError.UnKnown"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) ParseNfcDataWithContext(ctx context.Context, request *ParseNfcDataRequest) (response *ParseNfcDataResponse, err error) {
    if request == nil {
        request = NewParseNfcDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ParseNfcData require credential")
    }

    request.SetContext(ctx)
    
    response = NewParseNfcDataResponse()
    err = c.Send(request, response)
    return
}

func NewPhoneVerificationRequest() (request *PhoneVerificationRequest) {
    request = &PhoneVerificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "PhoneVerification")
    
    
    return
}

func NewPhoneVerificationResponse() (response *PhoneVerificationResponse) {
    response = &PhoneVerificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PhoneVerification
// 本接口用于校验手机号、姓名和身份证号的真实性和一致性。支持的手机号段详情请查阅<a href="https://cloud.tencent.com/document/product/1007/46063">运营商类</a>文档。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTENCRYPTFIELD = "InvalidParameter.UnsupportEncryptField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_CHARGESTATUSEXCEPTION = "UnauthorizedOperation.ChargeStatusException"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) PhoneVerification(request *PhoneVerificationRequest) (response *PhoneVerificationResponse, err error) {
    return c.PhoneVerificationWithContext(context.Background(), request)
}

// PhoneVerification
// 本接口用于校验手机号、姓名和身份证号的真实性和一致性。支持的手机号段详情请查阅<a href="https://cloud.tencent.com/document/product/1007/46063">运营商类</a>文档。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTENCRYPTFIELD = "InvalidParameter.UnsupportEncryptField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_CHARGESTATUSEXCEPTION = "UnauthorizedOperation.ChargeStatusException"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) PhoneVerificationWithContext(ctx context.Context, request *PhoneVerificationRequest) (response *PhoneVerificationResponse, err error) {
    if request == nil {
        request = NewPhoneVerificationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PhoneVerification require credential")
    }

    request.SetContext(ctx)
    
    response = NewPhoneVerificationResponse()
    err = c.Send(request, response)
    return
}

func NewPhoneVerificationCMCCRequest() (request *PhoneVerificationCMCCRequest) {
    request = &PhoneVerificationCMCCRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "PhoneVerificationCMCC")
    
    
    return
}

func NewPhoneVerificationCMCCResponse() (response *PhoneVerificationCMCCResponse) {
    response = &PhoneVerificationCMCCResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PhoneVerificationCMCC
// 本接口用于校验中国移动手机号、姓名和身份证号的真实性和一致性。中国移动支持的手机号段详情请查阅<a href="https://cloud.tencent.com/document/product/1007/46063">运营商类</a>文档。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTENCRYPTFIELD = "InvalidParameter.UnsupportEncryptField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_CHARGESTATUSEXCEPTION = "UnauthorizedOperation.ChargeStatusException"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) PhoneVerificationCMCC(request *PhoneVerificationCMCCRequest) (response *PhoneVerificationCMCCResponse, err error) {
    return c.PhoneVerificationCMCCWithContext(context.Background(), request)
}

// PhoneVerificationCMCC
// 本接口用于校验中国移动手机号、姓名和身份证号的真实性和一致性。中国移动支持的手机号段详情请查阅<a href="https://cloud.tencent.com/document/product/1007/46063">运营商类</a>文档。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTENCRYPTFIELD = "InvalidParameter.UnsupportEncryptField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_CHARGESTATUSEXCEPTION = "UnauthorizedOperation.ChargeStatusException"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) PhoneVerificationCMCCWithContext(ctx context.Context, request *PhoneVerificationCMCCRequest) (response *PhoneVerificationCMCCResponse, err error) {
    if request == nil {
        request = NewPhoneVerificationCMCCRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PhoneVerificationCMCC require credential")
    }

    request.SetContext(ctx)
    
    response = NewPhoneVerificationCMCCResponse()
    err = c.Send(request, response)
    return
}

func NewPhoneVerificationCTCCRequest() (request *PhoneVerificationCTCCRequest) {
    request = &PhoneVerificationCTCCRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "PhoneVerificationCTCC")
    
    
    return
}

func NewPhoneVerificationCTCCResponse() (response *PhoneVerificationCTCCResponse) {
    response = &PhoneVerificationCTCCResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PhoneVerificationCTCC
// 本接口用于校验中国电信手机号、姓名和身份证号的真实性和一致性。中国电信支持的手机号段详情请查阅<a href="https://cloud.tencent.com/document/product/1007/46063">运营商类</a>文档。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTENCRYPTFIELD = "InvalidParameter.UnsupportEncryptField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_CHARGESTATUSEXCEPTION = "UnauthorizedOperation.ChargeStatusException"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) PhoneVerificationCTCC(request *PhoneVerificationCTCCRequest) (response *PhoneVerificationCTCCResponse, err error) {
    return c.PhoneVerificationCTCCWithContext(context.Background(), request)
}

// PhoneVerificationCTCC
// 本接口用于校验中国电信手机号、姓名和身份证号的真实性和一致性。中国电信支持的手机号段详情请查阅<a href="https://cloud.tencent.com/document/product/1007/46063">运营商类</a>文档。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTENCRYPTFIELD = "InvalidParameter.UnsupportEncryptField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_CHARGESTATUSEXCEPTION = "UnauthorizedOperation.ChargeStatusException"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) PhoneVerificationCTCCWithContext(ctx context.Context, request *PhoneVerificationCTCCRequest) (response *PhoneVerificationCTCCResponse, err error) {
    if request == nil {
        request = NewPhoneVerificationCTCCRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PhoneVerificationCTCC require credential")
    }

    request.SetContext(ctx)
    
    response = NewPhoneVerificationCTCCResponse()
    err = c.Send(request, response)
    return
}

func NewPhoneVerificationCUCCRequest() (request *PhoneVerificationCUCCRequest) {
    request = &PhoneVerificationCUCCRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("faceid", APIVersion, "PhoneVerificationCUCC")
    
    
    return
}

func NewPhoneVerificationCUCCResponse() (response *PhoneVerificationCUCCResponse) {
    response = &PhoneVerificationCUCCResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PhoneVerificationCUCC
// 本接口用于校验中国联通手机号、姓名和身份证号的真实性和一致性。中国联通支持的手机号段详情请查阅<a href="https://cloud.tencent.com/document/product/1007/46063">运营商类</a>文档。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTENCRYPTFIELD = "InvalidParameter.UnsupportEncryptField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_CHARGESTATUSEXCEPTION = "UnauthorizedOperation.ChargeStatusException"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) PhoneVerificationCUCC(request *PhoneVerificationCUCCRequest) (response *PhoneVerificationCUCCResponse, err error) {
    return c.PhoneVerificationCUCCWithContext(context.Background(), request)
}

// PhoneVerificationCUCC
// 本接口用于校验中国联通手机号、姓名和身份证号的真实性和一致性。中国联通支持的手机号段详情请查阅<a href="https://cloud.tencent.com/document/product/1007/46063">运营商类</a>文档。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DECRYPTSYSTEMERROR = "FailedOperation.DecryptSystemError"
//  FAILEDOPERATION_STSUNAUTHERRERROR = "FailedOperation.StsUnAuthErrError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.UnKnown"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTENCRYPTFIELD = "InvalidParameter.UnsupportEncryptField"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_ARREARS = "UnauthorizedOperation.Arrears"
//  UNAUTHORIZEDOPERATION_CHARGESTATUSEXCEPTION = "UnauthorizedOperation.ChargeStatusException"
//  UNAUTHORIZEDOPERATION_NONAUTHORIZE = "UnauthorizedOperation.NonAuthorize"
//  UNAUTHORIZEDOPERATION_NONACTIVATED = "UnauthorizedOperation.Nonactivated"
func (c *Client) PhoneVerificationCUCCWithContext(ctx context.Context, request *PhoneVerificationCUCCRequest) (response *PhoneVerificationCUCCResponse, err error) {
    if request == nil {
        request = NewPhoneVerificationCUCCRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PhoneVerificationCUCC require credential")
    }

    request.SetContext(ctx)
    
    response = NewPhoneVerificationCUCCResponse()
    err = c.Send(request, response)
    return
}
