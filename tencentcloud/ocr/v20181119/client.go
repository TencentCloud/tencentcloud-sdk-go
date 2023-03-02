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

package v20181119

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-11-19"

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


func NewAdvertiseOCRRequest() (request *AdvertiseOCRRequest) {
    request = &AdvertiseOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "AdvertiseOCR")
    
    
    return
}

func NewAdvertiseOCRResponse() (response *AdvertiseOCRResponse) {
    response = &AdvertiseOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AdvertiseOCR
// 本接口支持广告商品图片内文字的检测和识别，返回文本框位置与文字内容。
//
// 
//
// 产品优势：针对广告商品图片普遍存在较多繁体字、艺术字的特点，进行了识别能力的增强。支持中英文、横排、竖排以及倾斜场景文字识别。文字识别的召回率和准确率能达到96%以上。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_ENGINERECOGNIZETIMEOUT = "FailedOperation.EngineRecognizeTimeout"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_LANGUAGENOTSUPPORT = "FailedOperation.LanguageNotSupport"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) AdvertiseOCR(request *AdvertiseOCRRequest) (response *AdvertiseOCRResponse, err error) {
    return c.AdvertiseOCRWithContext(context.Background(), request)
}

// AdvertiseOCR
// 本接口支持广告商品图片内文字的检测和识别，返回文本框位置与文字内容。
//
// 
//
// 产品优势：针对广告商品图片普遍存在较多繁体字、艺术字的特点，进行了识别能力的增强。支持中英文、横排、竖排以及倾斜场景文字识别。文字识别的召回率和准确率能达到96%以上。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_ENGINERECOGNIZETIMEOUT = "FailedOperation.EngineRecognizeTimeout"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_LANGUAGENOTSUPPORT = "FailedOperation.LanguageNotSupport"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) AdvertiseOCRWithContext(ctx context.Context, request *AdvertiseOCRRequest) (response *AdvertiseOCRResponse, err error) {
    if request == nil {
        request = NewAdvertiseOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AdvertiseOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewAdvertiseOCRResponse()
    err = c.Send(request, response)
    return
}

func NewArithmeticOCRRequest() (request *ArithmeticOCRRequest) {
    request = &ArithmeticOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "ArithmeticOCR")
    
    
    return
}

func NewArithmeticOCRResponse() (response *ArithmeticOCRResponse) {
    response = &ArithmeticOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ArithmeticOCR
// 本接口支持作业算式题目的自动识别和判分，目前覆盖 K12 学力范围内的 11 种题型，包括加减乘除四则、加减乘除已知结果求运算因子、判断大小、约等于估算、带余数除法、分数四则运算、单位换算、竖式加减法、竖式乘除法、脱式计算和解方程，平均识别精度达到93%以上。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_LANGUAGENOTSUPPORT = "FailedOperation.LanguageNotSupport"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) ArithmeticOCR(request *ArithmeticOCRRequest) (response *ArithmeticOCRResponse, err error) {
    return c.ArithmeticOCRWithContext(context.Background(), request)
}

// ArithmeticOCR
// 本接口支持作业算式题目的自动识别和判分，目前覆盖 K12 学力范围内的 11 种题型，包括加减乘除四则、加减乘除已知结果求运算因子、判断大小、约等于估算、带余数除法、分数四则运算、单位换算、竖式加减法、竖式乘除法、脱式计算和解方程，平均识别精度达到93%以上。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_LANGUAGENOTSUPPORT = "FailedOperation.LanguageNotSupport"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) ArithmeticOCRWithContext(ctx context.Context, request *ArithmeticOCRRequest) (response *ArithmeticOCRResponse, err error) {
    if request == nil {
        request = NewArithmeticOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ArithmeticOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewArithmeticOCRResponse()
    err = c.Send(request, response)
    return
}

func NewBankCardOCRRequest() (request *BankCardOCRRequest) {
    request = &BankCardOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "BankCardOCR")
    
    
    return
}

func NewBankCardOCRResponse() (response *BankCardOCRResponse) {
    response = &BankCardOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BankCardOCR
// 本接口支持对中国大陆主流银行卡正反面关键字段的检测与识别，包括卡号、卡类型、卡名字、银行信息、有效期。支持竖排异形卡识别、多角度旋转图片识别。支持对复印件、翻拍件、边框遮挡的银行卡进行告警，可应用于各种银行卡信息有效性校验场景，如金融行业身份认证、第三方支付绑卡等场景。
//
// 
//
// 默认接口请求频率限制：10次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_ILLEGALBANKCARDERROR = "FailedOperation.IllegalBankCardError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_NOBANKCARDERROR = "FailedOperation.NoBankCardError"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) BankCardOCR(request *BankCardOCRRequest) (response *BankCardOCRResponse, err error) {
    return c.BankCardOCRWithContext(context.Background(), request)
}

// BankCardOCR
// 本接口支持对中国大陆主流银行卡正反面关键字段的检测与识别，包括卡号、卡类型、卡名字、银行信息、有效期。支持竖排异形卡识别、多角度旋转图片识别。支持对复印件、翻拍件、边框遮挡的银行卡进行告警，可应用于各种银行卡信息有效性校验场景，如金融行业身份认证、第三方支付绑卡等场景。
//
// 
//
// 默认接口请求频率限制：10次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_ILLEGALBANKCARDERROR = "FailedOperation.IllegalBankCardError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_NOBANKCARDERROR = "FailedOperation.NoBankCardError"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) BankCardOCRWithContext(ctx context.Context, request *BankCardOCRRequest) (response *BankCardOCRResponse, err error) {
    if request == nil {
        request = NewBankCardOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BankCardOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewBankCardOCRResponse()
    err = c.Send(request, response)
    return
}

func NewBankSlipOCRRequest() (request *BankSlipOCRRequest) {
    request = &BankSlipOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "BankSlipOCR")
    
    
    return
}

func NewBankSlipOCRResponse() (response *BankSlipOCRResponse) {
    response = &BankSlipOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BankSlipOCR
// 本接口支持银行回单全字段的识别，包括付款开户行、收款开户行、付款账号、收款账号、回单类型、回单编号、币种、流水号、凭证号码、交易机构、交易金额、手续费、日期等字段信息。
//
// 
//
// 默认接口请求频率限制：10次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) BankSlipOCR(request *BankSlipOCRRequest) (response *BankSlipOCRResponse, err error) {
    return c.BankSlipOCRWithContext(context.Background(), request)
}

// BankSlipOCR
// 本接口支持银行回单全字段的识别，包括付款开户行、收款开户行、付款账号、收款账号、回单类型、回单编号、币种、流水号、凭证号码、交易机构、交易金额、手续费、日期等字段信息。
//
// 
//
// 默认接口请求频率限制：10次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) BankSlipOCRWithContext(ctx context.Context, request *BankSlipOCRRequest) (response *BankSlipOCRResponse, err error) {
    if request == nil {
        request = NewBankSlipOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BankSlipOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewBankSlipOCRResponse()
    err = c.Send(request, response)
    return
}

func NewBizLicenseOCRRequest() (request *BizLicenseOCRRequest) {
    request = &BizLicenseOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "BizLicenseOCR")
    
    
    return
}

func NewBizLicenseOCRResponse() (response *BizLicenseOCRResponse) {
    response = &BizLicenseOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BizLicenseOCR
// 本接口支持快速精准识别营业执照上的字段，包括统一社会信用代码、公司名称、经营场所、主体类型、法定代表人、注册资金、组成形式、成立日期、营业期限和经营范围等字段。
//
// 
//
// 默认接口请求频率限制：10次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) BizLicenseOCR(request *BizLicenseOCRRequest) (response *BizLicenseOCRResponse, err error) {
    return c.BizLicenseOCRWithContext(context.Background(), request)
}

// BizLicenseOCR
// 本接口支持快速精准识别营业执照上的字段，包括统一社会信用代码、公司名称、经营场所、主体类型、法定代表人、注册资金、组成形式、成立日期、营业期限和经营范围等字段。
//
// 
//
// 默认接口请求频率限制：10次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) BizLicenseOCRWithContext(ctx context.Context, request *BizLicenseOCRRequest) (response *BizLicenseOCRResponse, err error) {
    if request == nil {
        request = NewBizLicenseOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BizLicenseOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewBizLicenseOCRResponse()
    err = c.Send(request, response)
    return
}

func NewBusInvoiceOCRRequest() (request *BusInvoiceOCRRequest) {
    request = &BusInvoiceOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "BusInvoiceOCR")
    
    
    return
}

func NewBusInvoiceOCRResponse() (response *BusInvoiceOCRResponse) {
    response = &BusInvoiceOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BusInvoiceOCR
// 本接口支持识别公路汽车客票的发票代码、发票号码、日期、姓名、票价等字段。
//
// 
//
// 默认接口请求频率限制：5次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) BusInvoiceOCR(request *BusInvoiceOCRRequest) (response *BusInvoiceOCRResponse, err error) {
    return c.BusInvoiceOCRWithContext(context.Background(), request)
}

// BusInvoiceOCR
// 本接口支持识别公路汽车客票的发票代码、发票号码、日期、姓名、票价等字段。
//
// 
//
// 默认接口请求频率限制：5次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) BusInvoiceOCRWithContext(ctx context.Context, request *BusInvoiceOCRRequest) (response *BusInvoiceOCRResponse, err error) {
    if request == nil {
        request = NewBusInvoiceOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BusInvoiceOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewBusInvoiceOCRResponse()
    err = c.Send(request, response)
    return
}

func NewBusinessCardOCRRequest() (request *BusinessCardOCRRequest) {
    request = &BusinessCardOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "BusinessCardOCR")
    
    
    return
}

func NewBusinessCardOCRResponse() (response *BusinessCardOCRResponse) {
    response = &BusinessCardOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BusinessCardOCR
// 本接口支持名片各字段的自动定位与识别，包含姓名、电话、手机号、邮箱、公司、部门、职位、网址、地址、QQ、微信、MSN等。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOBUSINESSCARD = "FailedOperation.ImageNoBusinessCard"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_CONFIGFORMATERROR = "InvalidParameter.ConfigFormatError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) BusinessCardOCR(request *BusinessCardOCRRequest) (response *BusinessCardOCRResponse, err error) {
    return c.BusinessCardOCRWithContext(context.Background(), request)
}

// BusinessCardOCR
// 本接口支持名片各字段的自动定位与识别，包含姓名、电话、手机号、邮箱、公司、部门、职位、网址、地址、QQ、微信、MSN等。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOBUSINESSCARD = "FailedOperation.ImageNoBusinessCard"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_CONFIGFORMATERROR = "InvalidParameter.ConfigFormatError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) BusinessCardOCRWithContext(ctx context.Context, request *BusinessCardOCRRequest) (response *BusinessCardOCRResponse, err error) {
    if request == nil {
        request = NewBusinessCardOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BusinessCardOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewBusinessCardOCRResponse()
    err = c.Send(request, response)
    return
}

func NewCarInvoiceOCRRequest() (request *CarInvoiceOCRRequest) {
    request = &CarInvoiceOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "CarInvoiceOCR")
    
    
    return
}

func NewCarInvoiceOCRResponse() (response *CarInvoiceOCRResponse) {
    response = &CarInvoiceOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CarInvoiceOCR
// 本接口支持机动车销售统一发票和二手车销售统一发票的识别，包括发票号码、发票代码、合计金额、合计税额等二十多个字段。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) CarInvoiceOCR(request *CarInvoiceOCRRequest) (response *CarInvoiceOCRResponse, err error) {
    return c.CarInvoiceOCRWithContext(context.Background(), request)
}

// CarInvoiceOCR
// 本接口支持机动车销售统一发票和二手车销售统一发票的识别，包括发票号码、发票代码、合计金额、合计税额等二十多个字段。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) CarInvoiceOCRWithContext(ctx context.Context, request *CarInvoiceOCRRequest) (response *CarInvoiceOCRResponse, err error) {
    if request == nil {
        request = NewCarInvoiceOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CarInvoiceOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewCarInvoiceOCRResponse()
    err = c.Send(request, response)
    return
}

func NewClassifyDetectOCRRequest() (request *ClassifyDetectOCRRequest) {
    request = &ClassifyDetectOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "ClassifyDetectOCR")
    
    
    return
}

func NewClassifyDetectOCRResponse() (response *ClassifyDetectOCRResponse) {
    response = &ClassifyDetectOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ClassifyDetectOCR
// 支持身份证、护照、名片、银行卡、行驶证、驾驶证、港澳台通行证、户口本、港澳台来往内地通行证、港澳台居住证、不动产证、营业执照的智能分类。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEBLUR = "FailedOperation.ImageBlur"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) ClassifyDetectOCR(request *ClassifyDetectOCRRequest) (response *ClassifyDetectOCRResponse, err error) {
    return c.ClassifyDetectOCRWithContext(context.Background(), request)
}

// ClassifyDetectOCR
// 支持身份证、护照、名片、银行卡、行驶证、驾驶证、港澳台通行证、户口本、港澳台来往内地通行证、港澳台居住证、不动产证、营业执照的智能分类。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEBLUR = "FailedOperation.ImageBlur"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) ClassifyDetectOCRWithContext(ctx context.Context, request *ClassifyDetectOCRRequest) (response *ClassifyDetectOCRResponse, err error) {
    if request == nil {
        request = NewClassifyDetectOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ClassifyDetectOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewClassifyDetectOCRResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAIFormTaskRequest() (request *CreateAIFormTaskRequest) {
    request = &CreateAIFormTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "CreateAIFormTask")
    
    
    return
}

func NewCreateAIFormTaskResponse() (response *CreateAIFormTaskResponse) {
    response = &CreateAIFormTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAIFormTask
// 本接口可创建智能表单录入任务，支持多个识别图片和PDF的URL上传，返回含有识别内容的操作页面URL。
//
// 
//
// 智能表单录入产品提供高准确率的表单识别技术和人工核对工具，支持自定义字段，将识别结果自动填入到自定义条目中，并提供人工操作工具，完成整个表单识别过程。适用性强，可对票据、合同、货单等文件的识别，适用于金融、货代、保险、档案等领域。本产品免费公测中，您可以点击demo（超连接：https://ocr.smartform.cloud.tencent.com/）试用，如需购买请与商务团队联系。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DbError"
//  FAILEDOPERATION_USERQUOTAERROR = "FailedOperation.UserQuotaError"
//  INVALIDPARAMETERVALUE_FILEURLILLEGALERROR = "InvalidParameterValue.FileUrlIllegalError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
func (c *Client) CreateAIFormTask(request *CreateAIFormTaskRequest) (response *CreateAIFormTaskResponse, err error) {
    return c.CreateAIFormTaskWithContext(context.Background(), request)
}

// CreateAIFormTask
// 本接口可创建智能表单录入任务，支持多个识别图片和PDF的URL上传，返回含有识别内容的操作页面URL。
//
// 
//
// 智能表单录入产品提供高准确率的表单识别技术和人工核对工具，支持自定义字段，将识别结果自动填入到自定义条目中，并提供人工操作工具，完成整个表单识别过程。适用性强，可对票据、合同、货单等文件的识别，适用于金融、货代、保险、档案等领域。本产品免费公测中，您可以点击demo（超连接：https://ocr.smartform.cloud.tencent.com/）试用，如需购买请与商务团队联系。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DbError"
//  FAILEDOPERATION_USERQUOTAERROR = "FailedOperation.UserQuotaError"
//  INVALIDPARAMETERVALUE_FILEURLILLEGALERROR = "InvalidParameterValue.FileUrlIllegalError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
func (c *Client) CreateAIFormTaskWithContext(ctx context.Context, request *CreateAIFormTaskRequest) (response *CreateAIFormTaskResponse, err error) {
    if request == nil {
        request = NewCreateAIFormTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAIFormTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAIFormTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDriverLicenseOCRRequest() (request *DriverLicenseOCRRequest) {
    request = &DriverLicenseOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "DriverLicenseOCR")
    
    
    return
}

func NewDriverLicenseOCRResponse() (response *DriverLicenseOCRResponse) {
    response = &DriverLicenseOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DriverLicenseOCR
// 本接口支持驾驶证主页和副页所有字段的自动定位与识别，重点字段的识别准确度达到99%以上。
//
// 
//
// 驾驶证主页：包括证号、姓名、性别、国籍、住址、出生日期、初次领证日期、准驾车型、有效期限、发证单位
//
// 
//
// 驾驶证副页：包括证号、姓名、档案编号、记录。
//
// 
//
// 另外，本接口还支持复印件、翻拍和PS告警功能。同时支持识别交管12123APP发放的电子驾驶证正页。
//
// 
//
// 电子驾驶证正页：包括证号、姓名、性别、国籍、出生日期、初次领证日期、准驾车型、有效期开始时间、有效期截止时间、档案编号、状态、累积记分。
//
// 
//
// 默认接口请求频率限制：10次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEBLUR = "FailedOperation.ImageBlur"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) DriverLicenseOCR(request *DriverLicenseOCRRequest) (response *DriverLicenseOCRResponse, err error) {
    return c.DriverLicenseOCRWithContext(context.Background(), request)
}

// DriverLicenseOCR
// 本接口支持驾驶证主页和副页所有字段的自动定位与识别，重点字段的识别准确度达到99%以上。
//
// 
//
// 驾驶证主页：包括证号、姓名、性别、国籍、住址、出生日期、初次领证日期、准驾车型、有效期限、发证单位
//
// 
//
// 驾驶证副页：包括证号、姓名、档案编号、记录。
//
// 
//
// 另外，本接口还支持复印件、翻拍和PS告警功能。同时支持识别交管12123APP发放的电子驾驶证正页。
//
// 
//
// 电子驾驶证正页：包括证号、姓名、性别、国籍、出生日期、初次领证日期、准驾车型、有效期开始时间、有效期截止时间、档案编号、状态、累积记分。
//
// 
//
// 默认接口请求频率限制：10次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEBLUR = "FailedOperation.ImageBlur"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) DriverLicenseOCRWithContext(ctx context.Context, request *DriverLicenseOCRRequest) (response *DriverLicenseOCRResponse, err error) {
    if request == nil {
        request = NewDriverLicenseOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DriverLicenseOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewDriverLicenseOCRResponse()
    err = c.Send(request, response)
    return
}

func NewDutyPaidProofOCRRequest() (request *DutyPaidProofOCRRequest) {
    request = &DutyPaidProofOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "DutyPaidProofOCR")
    
    
    return
}

func NewDutyPaidProofOCRResponse() (response *DutyPaidProofOCRResponse) {
    response = &DutyPaidProofOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DutyPaidProofOCR
// 本接口支持对完税证明的税号、纳税人识别号、纳税人名称、金额合计大写、金额合计小写、填发日期、税务机关、填票人等关键字段的识别。
//
// 
//
// 默认接口请求频率限制：5次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) DutyPaidProofOCR(request *DutyPaidProofOCRRequest) (response *DutyPaidProofOCRResponse, err error) {
    return c.DutyPaidProofOCRWithContext(context.Background(), request)
}

// DutyPaidProofOCR
// 本接口支持对完税证明的税号、纳税人识别号、纳税人名称、金额合计大写、金额合计小写、填发日期、税务机关、填票人等关键字段的识别。
//
// 
//
// 默认接口请求频率限制：5次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) DutyPaidProofOCRWithContext(ctx context.Context, request *DutyPaidProofOCRRequest) (response *DutyPaidProofOCRResponse, err error) {
    if request == nil {
        request = NewDutyPaidProofOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DutyPaidProofOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewDutyPaidProofOCRResponse()
    err = c.Send(request, response)
    return
}

func NewEduPaperOCRRequest() (request *EduPaperOCRRequest) {
    request = &EduPaperOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "EduPaperOCR")
    
    
    return
}

func NewEduPaperOCRResponse() (response *EduPaperOCRResponse) {
    response = &EduPaperOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EduPaperOCR
// 本接口支持数学试题内容的识别和结构化输出，包括通用文本解析和小学/初中/高中数学公式解析能力（包括91种题型，180种符号），公式返回格式为 Latex 格式文本。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) EduPaperOCR(request *EduPaperOCRRequest) (response *EduPaperOCRResponse, err error) {
    return c.EduPaperOCRWithContext(context.Background(), request)
}

// EduPaperOCR
// 本接口支持数学试题内容的识别和结构化输出，包括通用文本解析和小学/初中/高中数学公式解析能力（包括91种题型，180种符号），公式返回格式为 Latex 格式文本。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) EduPaperOCRWithContext(ctx context.Context, request *EduPaperOCRRequest) (response *EduPaperOCRResponse, err error) {
    if request == nil {
        request = NewEduPaperOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EduPaperOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewEduPaperOCRResponse()
    err = c.Send(request, response)
    return
}

func NewEnglishOCRRequest() (request *EnglishOCRRequest) {
    request = &EnglishOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "EnglishOCR")
    
    
    return
}

func NewEnglishOCRResponse() (response *EnglishOCRResponse) {
    response = &EnglishOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EnglishOCR
// 本接口支持图像英文文字的检测和识别，返回文字框位置与文字内容。支持多场景、任意版面下的英文、字母、数字和常见字符的识别，同时覆盖英文印刷体和英文手写体识别。
//
// 
//
// 默认接口请求频率限制：10次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_LANGUAGENOTSUPPORT = "FailedOperation.LanguageNotSupport"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) EnglishOCR(request *EnglishOCRRequest) (response *EnglishOCRResponse, err error) {
    return c.EnglishOCRWithContext(context.Background(), request)
}

// EnglishOCR
// 本接口支持图像英文文字的检测和识别，返回文字框位置与文字内容。支持多场景、任意版面下的英文、字母、数字和常见字符的识别，同时覆盖英文印刷体和英文手写体识别。
//
// 
//
// 默认接口请求频率限制：10次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_LANGUAGENOTSUPPORT = "FailedOperation.LanguageNotSupport"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) EnglishOCRWithContext(ctx context.Context, request *EnglishOCRRequest) (response *EnglishOCRResponse, err error) {
    if request == nil {
        request = NewEnglishOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnglishOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnglishOCRResponse()
    err = c.Send(request, response)
    return
}

func NewEnterpriseLicenseOCRRequest() (request *EnterpriseLicenseOCRRequest) {
    request = &EnterpriseLicenseOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "EnterpriseLicenseOCR")
    
    
    return
}

func NewEnterpriseLicenseOCRResponse() (response *EnterpriseLicenseOCRResponse) {
    response = &EnterpriseLicenseOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EnterpriseLicenseOCR
// 本接口支持智能化识别各类企业登记证书、许可证书、企业执照、三证合一类证书，结构化输出统一社会信用代码、公司名称、法定代表人、公司地址、注册资金、企业类型、经营范围等关键字段。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) EnterpriseLicenseOCR(request *EnterpriseLicenseOCRRequest) (response *EnterpriseLicenseOCRResponse, err error) {
    return c.EnterpriseLicenseOCRWithContext(context.Background(), request)
}

// EnterpriseLicenseOCR
// 本接口支持智能化识别各类企业登记证书、许可证书、企业执照、三证合一类证书，结构化输出统一社会信用代码、公司名称、法定代表人、公司地址、注册资金、企业类型、经营范围等关键字段。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) EnterpriseLicenseOCRWithContext(ctx context.Context, request *EnterpriseLicenseOCRRequest) (response *EnterpriseLicenseOCRResponse, err error) {
    if request == nil {
        request = NewEnterpriseLicenseOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnterpriseLicenseOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnterpriseLicenseOCRResponse()
    err = c.Send(request, response)
    return
}

func NewEstateCertOCRRequest() (request *EstateCertOCRRequest) {
    request = &EstateCertOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "EstateCertOCR")
    
    
    return
}

func NewEstateCertOCRResponse() (response *EstateCertOCRResponse) {
    response = &EstateCertOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EstateCertOCR
// 本接口支持不动产权证关键字段的识别，包括使用期限、面积、用途、权利性质、权利类型、坐落、共有情况、权利人、权利其他状况等。
//
// 
//
// 
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) EstateCertOCR(request *EstateCertOCRRequest) (response *EstateCertOCRResponse, err error) {
    return c.EstateCertOCRWithContext(context.Background(), request)
}

// EstateCertOCR
// 本接口支持不动产权证关键字段的识别，包括使用期限、面积、用途、权利性质、权利类型、坐落、共有情况、权利人、权利其他状况等。
//
// 
//
// 
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) EstateCertOCRWithContext(ctx context.Context, request *EstateCertOCRRequest) (response *EstateCertOCRResponse, err error) {
    if request == nil {
        request = NewEstateCertOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EstateCertOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewEstateCertOCRResponse()
    err = c.Send(request, response)
    return
}

func NewFinanBillOCRRequest() (request *FinanBillOCRRequest) {
    request = &FinanBillOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "FinanBillOCR")
    
    
    return
}

func NewFinanBillOCRResponse() (response *FinanBillOCRResponse) {
    response = &FinanBillOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// FinanBillOCR
// 本接口支持常见银行票据的自动分类和识别。整单识别包括支票（含现金支票、普通支票、转账支票），承兑汇票（含银行承兑汇票、商业承兑汇票）以及进账单等，适用于中国人民银行印发的 2010 版银行票据凭证版式（银发[2010]299 号）。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) FinanBillOCR(request *FinanBillOCRRequest) (response *FinanBillOCRResponse, err error) {
    return c.FinanBillOCRWithContext(context.Background(), request)
}

// FinanBillOCR
// 本接口支持常见银行票据的自动分类和识别。整单识别包括支票（含现金支票、普通支票、转账支票），承兑汇票（含银行承兑汇票、商业承兑汇票）以及进账单等，适用于中国人民银行印发的 2010 版银行票据凭证版式（银发[2010]299 号）。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) FinanBillOCRWithContext(ctx context.Context, request *FinanBillOCRRequest) (response *FinanBillOCRResponse, err error) {
    if request == nil {
        request = NewFinanBillOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("FinanBillOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewFinanBillOCRResponse()
    err = c.Send(request, response)
    return
}

func NewFinanBillSliceOCRRequest() (request *FinanBillSliceOCRRequest) {
    request = &FinanBillSliceOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "FinanBillSliceOCR")
    
    
    return
}

func NewFinanBillSliceOCRResponse() (response *FinanBillSliceOCRResponse) {
    response = &FinanBillSliceOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// FinanBillSliceOCR
// 本接口支持常见银行票据的自动分类和识别。切片识别包括金融行业常见票据的重要切片字段识别，包括金额、账号、日期、凭证号码等。（金融票据切片：金融票据中待识别字段及其周围局部区域的裁剪图像。）
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) FinanBillSliceOCR(request *FinanBillSliceOCRRequest) (response *FinanBillSliceOCRResponse, err error) {
    return c.FinanBillSliceOCRWithContext(context.Background(), request)
}

// FinanBillSliceOCR
// 本接口支持常见银行票据的自动分类和识别。切片识别包括金融行业常见票据的重要切片字段识别，包括金额、账号、日期、凭证号码等。（金融票据切片：金融票据中待识别字段及其周围局部区域的裁剪图像。）
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) FinanBillSliceOCRWithContext(ctx context.Context, request *FinanBillSliceOCRRequest) (response *FinanBillSliceOCRResponse, err error) {
    if request == nil {
        request = NewFinanBillSliceOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("FinanBillSliceOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewFinanBillSliceOCRResponse()
    err = c.Send(request, response)
    return
}

func NewFlightInvoiceOCRRequest() (request *FlightInvoiceOCRRequest) {
    request = &FlightInvoiceOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "FlightInvoiceOCR")
    
    
    return
}

func NewFlightInvoiceOCRResponse() (response *FlightInvoiceOCRResponse) {
    response = &FlightInvoiceOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// FlightInvoiceOCR
// 本接口支持机票行程单关键字段的识别，包括旅客姓名、有效身份证件号码、电子客票号码、验证码、填开单位、其他税费、燃油附加费、民航发展基金、保险费、销售单位代号、始发地、目的地、航班号、时间、日期、座位等级、承运人、发票消费类型、票价、合计金额、填开日期、国内国际标签、印刷序号、客票级别/类别、客票生效日期、有效期截止日期、免费行李等字段，支持航班信息多行明细输出。
//
// 
//
// 默认接口请求频率限制：5次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) FlightInvoiceOCR(request *FlightInvoiceOCRRequest) (response *FlightInvoiceOCRResponse, err error) {
    return c.FlightInvoiceOCRWithContext(context.Background(), request)
}

// FlightInvoiceOCR
// 本接口支持机票行程单关键字段的识别，包括旅客姓名、有效身份证件号码、电子客票号码、验证码、填开单位、其他税费、燃油附加费、民航发展基金、保险费、销售单位代号、始发地、目的地、航班号、时间、日期、座位等级、承运人、发票消费类型、票价、合计金额、填开日期、国内国际标签、印刷序号、客票级别/类别、客票生效日期、有效期截止日期、免费行李等字段，支持航班信息多行明细输出。
//
// 
//
// 默认接口请求频率限制：5次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) FlightInvoiceOCRWithContext(ctx context.Context, request *FlightInvoiceOCRRequest) (response *FlightInvoiceOCRResponse, err error) {
    if request == nil {
        request = NewFlightInvoiceOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("FlightInvoiceOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewFlightInvoiceOCRResponse()
    err = c.Send(request, response)
    return
}

func NewFormulaOCRRequest() (request *FormulaOCRRequest) {
    request = &FormulaOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "FormulaOCR")
    
    
    return
}

func NewFormulaOCRResponse() (response *FormulaOCRResponse) {
    response = &FormulaOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// FormulaOCR
// 本接口支持识别主流初高中数学符号和公式，返回公式的 Latex 格式文本。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) FormulaOCR(request *FormulaOCRRequest) (response *FormulaOCRResponse, err error) {
    return c.FormulaOCRWithContext(context.Background(), request)
}

// FormulaOCR
// 本接口支持识别主流初高中数学符号和公式，返回公式的 Latex 格式文本。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) FormulaOCRWithContext(ctx context.Context, request *FormulaOCRRequest) (response *FormulaOCRResponse, err error) {
    if request == nil {
        request = NewFormulaOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("FormulaOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewFormulaOCRResponse()
    err = c.Send(request, response)
    return
}

func NewGeneralAccurateOCRRequest() (request *GeneralAccurateOCRRequest) {
    request = &GeneralAccurateOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "GeneralAccurateOCR")
    
    
    return
}

func NewGeneralAccurateOCRResponse() (response *GeneralAccurateOCRResponse) {
    response = &GeneralAccurateOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GeneralAccurateOCR
// 本接口支持图像整体文字的检测和识别。支持中文、英文、中英文、数字和特殊字符号的识别，并返回文字框位置和文字内容。
//
// 
//
// 适用于文字较多、版式复杂、对识别准召率要求较高的场景，如试卷试题、网络图片、街景店招牌、法律卷宗等场景。
//
// 
//
// 产品优势：与通用印刷体识别相比，提供更高精度的文字识别服务，在文字较多、长串数字、小字、模糊字、倾斜文本等困难场景下，高精度版的准确率和召回率更高。
//
// 
//
// 通用印刷体识别不同版本的差异如下：
//
// <table style="width:715px">
//
//       <thead>
//
//         <tr>
//
//           <th style="width:150px"></th>
//
//           <th >【荐】通用印刷体识别（高精度版）</th>
//
//           <th style="width:200px"><a href="https://cloud.tencent.com/document/product/866/33526">【荐】通用印刷体识别</a></th>
//
//           <th><a href="https://cloud.tencent.com/document/product/866/37831">通用印刷体识别（精简版）</a></th>
//
//         </tr>
//
//       </thead>
//
//       <tbody>
//
//         <tr>
//
//           <td> 适用场景</td>
//
//           <td>适用于文字较多、长串数字、小字、模糊字、倾斜文本等困难场景</td>
//
//           <td>适用于所有通用场景的印刷体识别</td>
//
//           <td>适用于快速文本识别场景，准召率有一定损失，价格更优惠</td>
//
//         </tr>
//
//         <tr>
//
//           <td>识别准确率</td>
//
//           <td>99%</td>
//
//           <td>96%</td>
//
//           <td>91%</td>
//
//         </tr>
//
//         <tr>
//
//           <td>价格</td>
//
//           <td>高</td>
//
//           <td>中</td>
//
//           <td>低</td>
//
//         </tr>
//
//         <tr>
//
//           <td>支持的语言</td>
//
//           <td>中文、英文、中英文</td>
//
//           <td>中文、英文、中英文、日语、韩语、西班牙语、法语、德语、葡萄牙语、越南语、马来语、俄语、意大利语、荷兰语、瑞典语、芬兰语、丹麦语、挪威语、匈牙利语、泰语</td>  
//
//           <td>中文、英文、中英文</td>
//
//         </tr>
//
//         <tr>
//
//           <td>自动语言检测</td>
//
//           <td>支持</td>
//
//           <td>支持</td>  
//
//           <td>支持</td>
//
//         </tr>
//
//         <tr>
//
//           <td>返回文本行坐标</td>
//
//           <td>支持</td>
//
//           <td>支持</td>
//
//           <td>支持</td>
//
//         </tr>
//
//         <tr>
//
//           <td>自动旋转纠正</td>
//
//           <td>支持旋转识别，返回角度信息</td>
//
//           <td>支持旋转识别，返回角度信息</td>
//
//           <td>支持旋转识别，返回角度信息</td>
//
//         </tr>
//
//       </tbody>
//
//     </table>
//
// 
//
// 默认接口请求频率限制：10次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_ENGINERECOGNIZETIMEOUT = "FailedOperation.EngineRecognizeTimeout"
//  FAILEDOPERATION_IMAGEBLUR = "FailedOperation.ImageBlur"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) GeneralAccurateOCR(request *GeneralAccurateOCRRequest) (response *GeneralAccurateOCRResponse, err error) {
    return c.GeneralAccurateOCRWithContext(context.Background(), request)
}

// GeneralAccurateOCR
// 本接口支持图像整体文字的检测和识别。支持中文、英文、中英文、数字和特殊字符号的识别，并返回文字框位置和文字内容。
//
// 
//
// 适用于文字较多、版式复杂、对识别准召率要求较高的场景，如试卷试题、网络图片、街景店招牌、法律卷宗等场景。
//
// 
//
// 产品优势：与通用印刷体识别相比，提供更高精度的文字识别服务，在文字较多、长串数字、小字、模糊字、倾斜文本等困难场景下，高精度版的准确率和召回率更高。
//
// 
//
// 通用印刷体识别不同版本的差异如下：
//
// <table style="width:715px">
//
//       <thead>
//
//         <tr>
//
//           <th style="width:150px"></th>
//
//           <th >【荐】通用印刷体识别（高精度版）</th>
//
//           <th style="width:200px"><a href="https://cloud.tencent.com/document/product/866/33526">【荐】通用印刷体识别</a></th>
//
//           <th><a href="https://cloud.tencent.com/document/product/866/37831">通用印刷体识别（精简版）</a></th>
//
//         </tr>
//
//       </thead>
//
//       <tbody>
//
//         <tr>
//
//           <td> 适用场景</td>
//
//           <td>适用于文字较多、长串数字、小字、模糊字、倾斜文本等困难场景</td>
//
//           <td>适用于所有通用场景的印刷体识别</td>
//
//           <td>适用于快速文本识别场景，准召率有一定损失，价格更优惠</td>
//
//         </tr>
//
//         <tr>
//
//           <td>识别准确率</td>
//
//           <td>99%</td>
//
//           <td>96%</td>
//
//           <td>91%</td>
//
//         </tr>
//
//         <tr>
//
//           <td>价格</td>
//
//           <td>高</td>
//
//           <td>中</td>
//
//           <td>低</td>
//
//         </tr>
//
//         <tr>
//
//           <td>支持的语言</td>
//
//           <td>中文、英文、中英文</td>
//
//           <td>中文、英文、中英文、日语、韩语、西班牙语、法语、德语、葡萄牙语、越南语、马来语、俄语、意大利语、荷兰语、瑞典语、芬兰语、丹麦语、挪威语、匈牙利语、泰语</td>  
//
//           <td>中文、英文、中英文</td>
//
//         </tr>
//
//         <tr>
//
//           <td>自动语言检测</td>
//
//           <td>支持</td>
//
//           <td>支持</td>  
//
//           <td>支持</td>
//
//         </tr>
//
//         <tr>
//
//           <td>返回文本行坐标</td>
//
//           <td>支持</td>
//
//           <td>支持</td>
//
//           <td>支持</td>
//
//         </tr>
//
//         <tr>
//
//           <td>自动旋转纠正</td>
//
//           <td>支持旋转识别，返回角度信息</td>
//
//           <td>支持旋转识别，返回角度信息</td>
//
//           <td>支持旋转识别，返回角度信息</td>
//
//         </tr>
//
//       </tbody>
//
//     </table>
//
// 
//
// 默认接口请求频率限制：10次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_ENGINERECOGNIZETIMEOUT = "FailedOperation.EngineRecognizeTimeout"
//  FAILEDOPERATION_IMAGEBLUR = "FailedOperation.ImageBlur"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) GeneralAccurateOCRWithContext(ctx context.Context, request *GeneralAccurateOCRRequest) (response *GeneralAccurateOCRResponse, err error) {
    if request == nil {
        request = NewGeneralAccurateOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GeneralAccurateOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewGeneralAccurateOCRResponse()
    err = c.Send(request, response)
    return
}

func NewGeneralBasicOCRRequest() (request *GeneralBasicOCRRequest) {
    request = &GeneralBasicOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "GeneralBasicOCR")
    
    
    return
}

func NewGeneralBasicOCRResponse() (response *GeneralBasicOCRResponse) {
    response = &GeneralBasicOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GeneralBasicOCR
// 本接口支持图像整体文字的检测和识别。可以识别中文、英文、中英文、日语、韩语、西班牙语、法语、德语、葡萄牙语、越南语、马来语、俄语、意大利语、荷兰语、瑞典语、芬兰语、丹麦语、挪威语、匈牙利语、泰语，阿拉伯语20种语言，且各种语言均支持与英文混合的文字识别。
//
// 
//
// 适用于印刷文档识别、网络图片识别、广告图文字识别、街景店招牌识别、菜单识别、视频标题识别、头像文字识别等场景。
//
// 
//
// 产品优势：支持自动识别语言类型，可返回文本框坐标信息，对于倾斜文本支持自动旋转纠正。
//
// 
//
// 通用印刷体识别不同版本的差异如下：
//
// <table style="width:715px">
//
//       <thead>
//
//         <tr>
//
//           <th style="width:150px"></th>
//
//           <th style="width:200px">【荐】通用印刷体识别</th>
//
//           <th ><a href="https://cloud.tencent.com/document/product/866/34937">【荐】通用印刷体识别（高精度版）</a></th>
//
//           <th><a href="https://cloud.tencent.com/document/product/866/37831">通用印刷体识别（精简版）</a></th>
//
//         </tr>
//
//       </thead>
//
//       <tbody>
//
//         <tr>
//
//           <td> 适用场景</td>
//
//           <td>适用于所有通用场景的印刷体识别</td>
//
//           <td>适用于文字较多、长串数字、小字、模糊字、倾斜文本等困难场景</td>
//
//           <td>适用于快速文本识别场景，准召率有一定损失，价格更优惠</td>
//
//         </tr>
//
//         <tr>
//
//           <td>识别准确率</td>
//
//           <td>96%</td>
//
//           <td>99%</td>
//
//           <td>91%</td>
//
//         </tr>
//
//         <tr>
//
//           <td>价格</td>
//
//           <td>中</td>
//
//           <td>高</td>
//
//           <td>低</td>
//
//         </tr>
//
//         <tr>
//
//           <td>支持的语言</td>
//
//           <td>中文、英文、中英文、日语、韩语、西班牙语、法语、德语、葡萄牙语、越南语、马来语、俄语、意大利语、荷兰语、瑞典语、芬兰语、丹麦语、挪威语、匈牙利语、泰语</td>
//
//           <td>中文、英文、中英文</td>
//
//           <td>中文、英文、中英文</td>
//
//         </tr>
//
//         <tr>
//
//           <td>自动语言检测</td>
//
//           <td>支持</td>
//
//           <td>支持</td>
//
//           <td>支持</td>
//
//         </tr>
//
//         <tr>
//
//           <td>返回文本行坐标</td>
//
//           <td>支持</td>
//
//           <td>支持</td>
//
//           <td>支持</td>
//
//         </tr>
//
//         <tr>
//
//           <td>自动旋转纠正</td>
//
//           <td>支持旋转识别，返回角度信息</td>
//
//           <td>支持旋转识别，返回角度信息</td>
//
//           <td>支持旋转识别，返回角度信息</td>
//
//         </tr>
//
//       </tbody>
//
//     </table>
//
// 
//
// 默认接口请求频率限制：20次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_ENGINERECOGNIZETIMEOUT = "FailedOperation.EngineRecognizeTimeout"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_LANGUAGENOTSUPPORT = "FailedOperation.LanguageNotSupport"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) GeneralBasicOCR(request *GeneralBasicOCRRequest) (response *GeneralBasicOCRResponse, err error) {
    return c.GeneralBasicOCRWithContext(context.Background(), request)
}

// GeneralBasicOCR
// 本接口支持图像整体文字的检测和识别。可以识别中文、英文、中英文、日语、韩语、西班牙语、法语、德语、葡萄牙语、越南语、马来语、俄语、意大利语、荷兰语、瑞典语、芬兰语、丹麦语、挪威语、匈牙利语、泰语，阿拉伯语20种语言，且各种语言均支持与英文混合的文字识别。
//
// 
//
// 适用于印刷文档识别、网络图片识别、广告图文字识别、街景店招牌识别、菜单识别、视频标题识别、头像文字识别等场景。
//
// 
//
// 产品优势：支持自动识别语言类型，可返回文本框坐标信息，对于倾斜文本支持自动旋转纠正。
//
// 
//
// 通用印刷体识别不同版本的差异如下：
//
// <table style="width:715px">
//
//       <thead>
//
//         <tr>
//
//           <th style="width:150px"></th>
//
//           <th style="width:200px">【荐】通用印刷体识别</th>
//
//           <th ><a href="https://cloud.tencent.com/document/product/866/34937">【荐】通用印刷体识别（高精度版）</a></th>
//
//           <th><a href="https://cloud.tencent.com/document/product/866/37831">通用印刷体识别（精简版）</a></th>
//
//         </tr>
//
//       </thead>
//
//       <tbody>
//
//         <tr>
//
//           <td> 适用场景</td>
//
//           <td>适用于所有通用场景的印刷体识别</td>
//
//           <td>适用于文字较多、长串数字、小字、模糊字、倾斜文本等困难场景</td>
//
//           <td>适用于快速文本识别场景，准召率有一定损失，价格更优惠</td>
//
//         </tr>
//
//         <tr>
//
//           <td>识别准确率</td>
//
//           <td>96%</td>
//
//           <td>99%</td>
//
//           <td>91%</td>
//
//         </tr>
//
//         <tr>
//
//           <td>价格</td>
//
//           <td>中</td>
//
//           <td>高</td>
//
//           <td>低</td>
//
//         </tr>
//
//         <tr>
//
//           <td>支持的语言</td>
//
//           <td>中文、英文、中英文、日语、韩语、西班牙语、法语、德语、葡萄牙语、越南语、马来语、俄语、意大利语、荷兰语、瑞典语、芬兰语、丹麦语、挪威语、匈牙利语、泰语</td>
//
//           <td>中文、英文、中英文</td>
//
//           <td>中文、英文、中英文</td>
//
//         </tr>
//
//         <tr>
//
//           <td>自动语言检测</td>
//
//           <td>支持</td>
//
//           <td>支持</td>
//
//           <td>支持</td>
//
//         </tr>
//
//         <tr>
//
//           <td>返回文本行坐标</td>
//
//           <td>支持</td>
//
//           <td>支持</td>
//
//           <td>支持</td>
//
//         </tr>
//
//         <tr>
//
//           <td>自动旋转纠正</td>
//
//           <td>支持旋转识别，返回角度信息</td>
//
//           <td>支持旋转识别，返回角度信息</td>
//
//           <td>支持旋转识别，返回角度信息</td>
//
//         </tr>
//
//       </tbody>
//
//     </table>
//
// 
//
// 默认接口请求频率限制：20次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_ENGINERECOGNIZETIMEOUT = "FailedOperation.EngineRecognizeTimeout"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_LANGUAGENOTSUPPORT = "FailedOperation.LanguageNotSupport"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) GeneralBasicOCRWithContext(ctx context.Context, request *GeneralBasicOCRRequest) (response *GeneralBasicOCRResponse, err error) {
    if request == nil {
        request = NewGeneralBasicOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GeneralBasicOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewGeneralBasicOCRResponse()
    err = c.Send(request, response)
    return
}

func NewGeneralEfficientOCRRequest() (request *GeneralEfficientOCRRequest) {
    request = &GeneralEfficientOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "GeneralEfficientOCR")
    
    
    return
}

func NewGeneralEfficientOCRResponse() (response *GeneralEfficientOCRResponse) {
    response = &GeneralEfficientOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GeneralEfficientOCR
// 本接口支持图像整体文字的检测和识别。支持中文、英文、中英文、数字和特殊字符号的识别，并返回文字框位置和文字内容。
//
// 
//
// 适用于快速文本识别场景。
//
// 
//
// 产品优势：与通用印刷体识别接口相比，精简版虽然在准确率和召回率上有一定损失，但价格更加优惠。
//
// 
//
// 通用印刷体识别不同版本的差异如下：
//
// <table style="width:715px">
//
//       <thead>
//
//         <tr>
//
//           <th style="width:150px"></th>
//
//           <th >通用印刷体识别（精简版）</th>
//
//           <th style="width:200px"><a href="https://cloud.tencent.com/document/product/866/33526">【荐】通用印刷体识别</a></th>
//
//           <th><a href="https://cloud.tencent.com/document/product/866/34937">【荐】通用印刷体识别（高精度版）</a></th>
//
//         </tr>
//
//       </thead>
//
//       <tbody>
//
//         <tr>
//
//           <td> 适用场景</td>
//
//           <td>适用于快速文本识别场景，准召率有一定损失，价格更优惠</td>
//
//           <td>适用于所有通用场景的印刷体识别</td>
//
//           <td>适用于文字较多、长串数字、小字、模糊字、倾斜文本等困难场景</td>
//
//         </tr>
//
//         <tr>
//
//           <td>识别准确率</td>
//
//           <td>91%</td>
//
//           <td>96%</td>
//
//           <td>99%</td>
//
//         </tr>
//
//         <tr>
//
//           <td>价格</td>
//
//           <td>低</td>
//
//           <td>中</td>
//
//           <td>高</td>
//
//         </tr>
//
//         <tr>
//
//           <td>支持的语言</td>
//
//           <td>中文、英文、中英文</td>
//
//           <td>中文、英文、中英文、日语、韩语、西班牙语、法语、德语、葡萄牙语、越南语、马来语、俄语、意大利语、荷兰语、瑞典语、芬兰语、丹麦语、挪威语、匈牙利语、泰语</td>  
//
//           <td>中文、英文、中英文</td>
//
//         </tr>
//
//         <tr>
//
//           <td>自动语言检测</td>
//
//           <td>支持</td>
//
//           <td>支持</td>  
//
//           <td>支持</td>
//
//         </tr>
//
//         <tr>
//
//           <td>返回文本行坐标</td>
//
//           <td>支持</td>
//
//           <td>支持</td>
//
//           <td>支持</td>
//
//         </tr>
//
//         <tr>
//
//           <td>自动旋转纠正</td>
//
//           <td>支持旋转识别，返回角度信息</td>
//
//           <td>支持旋转识别，返回角度信息</td>
//
//           <td>支持旋转识别，返回角度信息</td>
//
//         </tr>
//
//       </tbody>
//
//     </table>
//
// 
//
// 默认接口请求频率限制：10次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) GeneralEfficientOCR(request *GeneralEfficientOCRRequest) (response *GeneralEfficientOCRResponse, err error) {
    return c.GeneralEfficientOCRWithContext(context.Background(), request)
}

// GeneralEfficientOCR
// 本接口支持图像整体文字的检测和识别。支持中文、英文、中英文、数字和特殊字符号的识别，并返回文字框位置和文字内容。
//
// 
//
// 适用于快速文本识别场景。
//
// 
//
// 产品优势：与通用印刷体识别接口相比，精简版虽然在准确率和召回率上有一定损失，但价格更加优惠。
//
// 
//
// 通用印刷体识别不同版本的差异如下：
//
// <table style="width:715px">
//
//       <thead>
//
//         <tr>
//
//           <th style="width:150px"></th>
//
//           <th >通用印刷体识别（精简版）</th>
//
//           <th style="width:200px"><a href="https://cloud.tencent.com/document/product/866/33526">【荐】通用印刷体识别</a></th>
//
//           <th><a href="https://cloud.tencent.com/document/product/866/34937">【荐】通用印刷体识别（高精度版）</a></th>
//
//         </tr>
//
//       </thead>
//
//       <tbody>
//
//         <tr>
//
//           <td> 适用场景</td>
//
//           <td>适用于快速文本识别场景，准召率有一定损失，价格更优惠</td>
//
//           <td>适用于所有通用场景的印刷体识别</td>
//
//           <td>适用于文字较多、长串数字、小字、模糊字、倾斜文本等困难场景</td>
//
//         </tr>
//
//         <tr>
//
//           <td>识别准确率</td>
//
//           <td>91%</td>
//
//           <td>96%</td>
//
//           <td>99%</td>
//
//         </tr>
//
//         <tr>
//
//           <td>价格</td>
//
//           <td>低</td>
//
//           <td>中</td>
//
//           <td>高</td>
//
//         </tr>
//
//         <tr>
//
//           <td>支持的语言</td>
//
//           <td>中文、英文、中英文</td>
//
//           <td>中文、英文、中英文、日语、韩语、西班牙语、法语、德语、葡萄牙语、越南语、马来语、俄语、意大利语、荷兰语、瑞典语、芬兰语、丹麦语、挪威语、匈牙利语、泰语</td>  
//
//           <td>中文、英文、中英文</td>
//
//         </tr>
//
//         <tr>
//
//           <td>自动语言检测</td>
//
//           <td>支持</td>
//
//           <td>支持</td>  
//
//           <td>支持</td>
//
//         </tr>
//
//         <tr>
//
//           <td>返回文本行坐标</td>
//
//           <td>支持</td>
//
//           <td>支持</td>
//
//           <td>支持</td>
//
//         </tr>
//
//         <tr>
//
//           <td>自动旋转纠正</td>
//
//           <td>支持旋转识别，返回角度信息</td>
//
//           <td>支持旋转识别，返回角度信息</td>
//
//           <td>支持旋转识别，返回角度信息</td>
//
//         </tr>
//
//       </tbody>
//
//     </table>
//
// 
//
// 默认接口请求频率限制：10次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) GeneralEfficientOCRWithContext(ctx context.Context, request *GeneralEfficientOCRRequest) (response *GeneralEfficientOCRResponse, err error) {
    if request == nil {
        request = NewGeneralEfficientOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GeneralEfficientOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewGeneralEfficientOCRResponse()
    err = c.Send(request, response)
    return
}

func NewGeneralFastOCRRequest() (request *GeneralFastOCRRequest) {
    request = &GeneralFastOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "GeneralFastOCR")
    
    
    return
}

func NewGeneralFastOCRResponse() (response *GeneralFastOCRResponse) {
    response = &GeneralFastOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GeneralFastOCR
// 本接口支持图片中整体文字的检测和识别，返回文字框位置与文字内容。相比通用印刷体识别接口，识别速度更快。
//
// 
//
// 默认接口请求频率限制：10次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_LANGUAGENOTSUPPORT = "FailedOperation.LanguageNotSupport"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) GeneralFastOCR(request *GeneralFastOCRRequest) (response *GeneralFastOCRResponse, err error) {
    return c.GeneralFastOCRWithContext(context.Background(), request)
}

// GeneralFastOCR
// 本接口支持图片中整体文字的检测和识别，返回文字框位置与文字内容。相比通用印刷体识别接口，识别速度更快。
//
// 
//
// 默认接口请求频率限制：10次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_LANGUAGENOTSUPPORT = "FailedOperation.LanguageNotSupport"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) GeneralFastOCRWithContext(ctx context.Context, request *GeneralFastOCRRequest) (response *GeneralFastOCRResponse, err error) {
    if request == nil {
        request = NewGeneralFastOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GeneralFastOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewGeneralFastOCRResponse()
    err = c.Send(request, response)
    return
}

func NewGeneralHandwritingOCRRequest() (request *GeneralHandwritingOCRRequest) {
    request = &GeneralHandwritingOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "GeneralHandwritingOCR")
    
    
    return
}

func NewGeneralHandwritingOCRResponse() (response *GeneralHandwritingOCRResponse) {
    response = &GeneralHandwritingOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GeneralHandwritingOCR
// 本接口支持图片内手写体文字的检测和识别，针对手写字体无规则、字迹潦草、模糊等特点进行了识别能力的增强。
//
// 
//
// 默认接口请求频率限制：10次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEBLUR = "FailedOperation.ImageBlur"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) GeneralHandwritingOCR(request *GeneralHandwritingOCRRequest) (response *GeneralHandwritingOCRResponse, err error) {
    return c.GeneralHandwritingOCRWithContext(context.Background(), request)
}

// GeneralHandwritingOCR
// 本接口支持图片内手写体文字的检测和识别，针对手写字体无规则、字迹潦草、模糊等特点进行了识别能力的增强。
//
// 
//
// 默认接口请求频率限制：10次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEBLUR = "FailedOperation.ImageBlur"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) GeneralHandwritingOCRWithContext(ctx context.Context, request *GeneralHandwritingOCRRequest) (response *GeneralHandwritingOCRResponse, err error) {
    if request == nil {
        request = NewGeneralHandwritingOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GeneralHandwritingOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewGeneralHandwritingOCRResponse()
    err = c.Send(request, response)
    return
}

func NewGetTaskStateRequest() (request *GetTaskStateRequest) {
    request = &GetTaskStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "GetTaskState")
    
    
    return
}

func NewGetTaskStateResponse() (response *GetTaskStateResponse) {
    response = &GetTaskStateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetTaskState
// 支持查询智能表单录入任务的状态。本产品免费公测中，您可以点击demo（超连接：https://ocr.smartform.cloud.tencent.com/）试用，如需购买请与商务团队联系。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DbError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
func (c *Client) GetTaskState(request *GetTaskStateRequest) (response *GetTaskStateResponse, err error) {
    return c.GetTaskStateWithContext(context.Background(), request)
}

// GetTaskState
// 支持查询智能表单录入任务的状态。本产品免费公测中，您可以点击demo（超连接：https://ocr.smartform.cloud.tencent.com/）试用，如需购买请与商务团队联系。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DbError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
func (c *Client) GetTaskStateWithContext(ctx context.Context, request *GetTaskStateRequest) (response *GetTaskStateResponse, err error) {
    if request == nil {
        request = NewGetTaskStateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTaskState require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTaskStateResponse()
    err = c.Send(request, response)
    return
}

func NewHKIDCardOCRRequest() (request *HKIDCardOCRRequest) {
    request = &HKIDCardOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "HKIDCardOCR")
    
    
    return
}

func NewHKIDCardOCRResponse() (response *HKIDCardOCRResponse) {
    response = &HKIDCardOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// HKIDCardOCR
// 本接口支持中国香港身份证人像面中关键字段的识别，包括中文姓名、英文姓名、姓名电码、出生日期、性别、证件符号、首次签发日期、最近领用日期、身份证号、是否是永久性居民身份证；具备防伪识别、人像照片裁剪等扩展功能。
//
// 
//
// 默认接口请求频率限制：5次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_NOHKIDCARD = "FailedOperation.NoHKIDCard"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) HKIDCardOCR(request *HKIDCardOCRRequest) (response *HKIDCardOCRResponse, err error) {
    return c.HKIDCardOCRWithContext(context.Background(), request)
}

// HKIDCardOCR
// 本接口支持中国香港身份证人像面中关键字段的识别，包括中文姓名、英文姓名、姓名电码、出生日期、性别、证件符号、首次签发日期、最近领用日期、身份证号、是否是永久性居民身份证；具备防伪识别、人像照片裁剪等扩展功能。
//
// 
//
// 默认接口请求频率限制：5次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_NOHKIDCARD = "FailedOperation.NoHKIDCard"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) HKIDCardOCRWithContext(ctx context.Context, request *HKIDCardOCRRequest) (response *HKIDCardOCRResponse, err error) {
    if request == nil {
        request = NewHKIDCardOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("HKIDCardOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewHKIDCardOCRResponse()
    err = c.Send(request, response)
    return
}

func NewHmtResidentPermitOCRRequest() (request *HmtResidentPermitOCRRequest) {
    request = &HmtResidentPermitOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "HmtResidentPermitOCR")
    
    
    return
}

func NewHmtResidentPermitOCRResponse() (response *HmtResidentPermitOCRResponse) {
    response = &HmtResidentPermitOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// HmtResidentPermitOCR
// 港澳台居住证OCR支持港澳台居住证正反面全字段内容检测识别功能，包括姓名、性别、出生日期、地址、身份证ID、签发机关、有效期限、签发次数、通行证号码关键字段识别。可以应用于港澳台居住证信息有效性校验场景，例如银行开户、用户注册等场景。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) HmtResidentPermitOCR(request *HmtResidentPermitOCRRequest) (response *HmtResidentPermitOCRResponse, err error) {
    return c.HmtResidentPermitOCRWithContext(context.Background(), request)
}

// HmtResidentPermitOCR
// 港澳台居住证OCR支持港澳台居住证正反面全字段内容检测识别功能，包括姓名、性别、出生日期、地址、身份证ID、签发机关、有效期限、签发次数、通行证号码关键字段识别。可以应用于港澳台居住证信息有效性校验场景，例如银行开户、用户注册等场景。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) HmtResidentPermitOCRWithContext(ctx context.Context, request *HmtResidentPermitOCRRequest) (response *HmtResidentPermitOCRResponse, err error) {
    if request == nil {
        request = NewHmtResidentPermitOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("HmtResidentPermitOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewHmtResidentPermitOCRResponse()
    err = c.Send(request, response)
    return
}

func NewIDCardOCRRequest() (request *IDCardOCRRequest) {
    request = &IDCardOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "IDCardOCR")
    
    
    return
}

func NewIDCardOCRResponse() (response *IDCardOCRResponse) {
    response = &IDCardOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// IDCardOCR
// 本接口支持中国大陆居民二代身份证正反面所有字段的识别，包括姓名、性别、民族、出生日期、住址、公民身份证号、签发机关、有效期限，识别准确度达到99%以上。
//
// 
//
// 另外，本接口还支持多种增值能力，满足不同场景的需求。如身份证照片、人像照片的裁剪功能，同时具备9种告警功能，如下表所示。
//
// 
//
// <table style="width:650px">
//
//       <thead>
//
//         <tr>
//
//        <th width="150">增值能力</th>
//
//           <th width="500">能力项</th>
//
//         </tr>
//
//       </thead>
//
//       <tbody>
//
//         <tr>
//
//           <td rowspan="2">裁剪功能</td>
//
//           <td>身份证照片裁剪（去掉证件外多余的边缘、自动矫正拍摄角度）</td>
//
//         </tr>
//
//         <tr>
//
//           <td>人像照片裁剪（自动抠取身份证头像区域）</td>
//
//         </tr>
//
//         <tr>
//
//           <td rowspan="9">告警功能</td>
//
//           <td>身份证有效日期不合法告警</td>
//
//         </tr>
//
//         <tr>
//
//           <td>身份证边框不完整告警</td>
//
//         </tr>
//
//         <tr>
//
//           <td>身份证复印件告警</td>
//
//         </tr>
//
//         <tr>
//
//           <td>身份证翻拍告警</td>
//
//         </tr>
//
//           <tr>
//
//           <td>身份证框内遮挡告警</td>
//
//         </tr>
//
//          <tr>
//
//           <td>临时身份证告警</td>
//
//         </tr>
//
//           <tr>
//
//           <td>身份证 PS 告警</td>
//
//         </tr>
//
//           <tr>
//
//           <td>图片模糊告警（可根据图片质量分数判断）</td>
//
//         </tr>
//
//       </tbody>
//
//     </table>
//
// 
//
// 默认接口请求频率限制：20次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IDCARDINFOILLEGAL = "FailedOperation.IdCardInfoIllegal"
//  FAILEDOPERATION_IMAGEBLUR = "FailedOperation.ImageBlur"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOIDCARD = "FailedOperation.ImageNoIdCard"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_MULTICARDERROR = "FailedOperation.MultiCardError"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_CONFIGFORMATERROR = "InvalidParameter.ConfigFormatError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) IDCardOCR(request *IDCardOCRRequest) (response *IDCardOCRResponse, err error) {
    return c.IDCardOCRWithContext(context.Background(), request)
}

// IDCardOCR
// 本接口支持中国大陆居民二代身份证正反面所有字段的识别，包括姓名、性别、民族、出生日期、住址、公民身份证号、签发机关、有效期限，识别准确度达到99%以上。
//
// 
//
// 另外，本接口还支持多种增值能力，满足不同场景的需求。如身份证照片、人像照片的裁剪功能，同时具备9种告警功能，如下表所示。
//
// 
//
// <table style="width:650px">
//
//       <thead>
//
//         <tr>
//
//        <th width="150">增值能力</th>
//
//           <th width="500">能力项</th>
//
//         </tr>
//
//       </thead>
//
//       <tbody>
//
//         <tr>
//
//           <td rowspan="2">裁剪功能</td>
//
//           <td>身份证照片裁剪（去掉证件外多余的边缘、自动矫正拍摄角度）</td>
//
//         </tr>
//
//         <tr>
//
//           <td>人像照片裁剪（自动抠取身份证头像区域）</td>
//
//         </tr>
//
//         <tr>
//
//           <td rowspan="9">告警功能</td>
//
//           <td>身份证有效日期不合法告警</td>
//
//         </tr>
//
//         <tr>
//
//           <td>身份证边框不完整告警</td>
//
//         </tr>
//
//         <tr>
//
//           <td>身份证复印件告警</td>
//
//         </tr>
//
//         <tr>
//
//           <td>身份证翻拍告警</td>
//
//         </tr>
//
//           <tr>
//
//           <td>身份证框内遮挡告警</td>
//
//         </tr>
//
//          <tr>
//
//           <td>临时身份证告警</td>
//
//         </tr>
//
//           <tr>
//
//           <td>身份证 PS 告警</td>
//
//         </tr>
//
//           <tr>
//
//           <td>图片模糊告警（可根据图片质量分数判断）</td>
//
//         </tr>
//
//       </tbody>
//
//     </table>
//
// 
//
// 默认接口请求频率限制：20次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IDCARDINFOILLEGAL = "FailedOperation.IdCardInfoIllegal"
//  FAILEDOPERATION_IMAGEBLUR = "FailedOperation.ImageBlur"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOIDCARD = "FailedOperation.ImageNoIdCard"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_MULTICARDERROR = "FailedOperation.MultiCardError"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_CONFIGFORMATERROR = "InvalidParameter.ConfigFormatError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) IDCardOCRWithContext(ctx context.Context, request *IDCardOCRRequest) (response *IDCardOCRResponse, err error) {
    if request == nil {
        request = NewIDCardOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("IDCardOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewIDCardOCRResponse()
    err = c.Send(request, response)
    return
}

func NewImageEnhancementRequest() (request *ImageEnhancementRequest) {
    request = &ImageEnhancementRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "ImageEnhancement")
    
    
    return
}

func NewImageEnhancementResponse() (response *ImageEnhancementResponse) {
    response = &ImageEnhancementResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ImageEnhancement
// 文本图像增强是面向文档类图片提供的图像增强处理能力，包括切边增强、图像矫正、阴影去除、摩尔纹去除等；可以有效优化文档类的图片质量，提升文字的清晰度。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) ImageEnhancement(request *ImageEnhancementRequest) (response *ImageEnhancementResponse, err error) {
    return c.ImageEnhancementWithContext(context.Background(), request)
}

// ImageEnhancement
// 文本图像增强是面向文档类图片提供的图像增强处理能力，包括切边增强、图像矫正、阴影去除、摩尔纹去除等；可以有效优化文档类的图片质量，提升文字的清晰度。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) ImageEnhancementWithContext(ctx context.Context, request *ImageEnhancementRequest) (response *ImageEnhancementResponse, err error) {
    if request == nil {
        request = NewImageEnhancementRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImageEnhancement require credential")
    }

    request.SetContext(ctx)
    
    response = NewImageEnhancementResponse()
    err = c.Send(request, response)
    return
}

func NewInstitutionOCRRequest() (request *InstitutionOCRRequest) {
    request = &InstitutionOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "InstitutionOCR")
    
    
    return
}

func NewInstitutionOCRResponse() (response *InstitutionOCRResponse) {
    response = &InstitutionOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InstitutionOCR
// 本接口支持事业单位法人证书关键字段识别，包括注册号、有效期、住所、名称、法定代表人等。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) InstitutionOCR(request *InstitutionOCRRequest) (response *InstitutionOCRResponse, err error) {
    return c.InstitutionOCRWithContext(context.Background(), request)
}

// InstitutionOCR
// 本接口支持事业单位法人证书关键字段识别，包括注册号、有效期、住所、名称、法定代表人等。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) InstitutionOCRWithContext(ctx context.Context, request *InstitutionOCRRequest) (response *InstitutionOCRResponse, err error) {
    if request == nil {
        request = NewInstitutionOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InstitutionOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewInstitutionOCRResponse()
    err = c.Send(request, response)
    return
}

func NewInsuranceBillOCRRequest() (request *InsuranceBillOCRRequest) {
    request = &InsuranceBillOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "InsuranceBillOCR")
    
    
    return
}

func NewInsuranceBillOCRResponse() (response *InsuranceBillOCRResponse) {
    response = &InsuranceBillOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InsuranceBillOCR
// 本接口支持病案首页、费用清单、结算单、医疗发票四种保险理赔单据的文本识别和结构化输出。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) InsuranceBillOCR(request *InsuranceBillOCRRequest) (response *InsuranceBillOCRResponse, err error) {
    return c.InsuranceBillOCRWithContext(context.Background(), request)
}

// InsuranceBillOCR
// 本接口支持病案首页、费用清单、结算单、医疗发票四种保险理赔单据的文本识别和结构化输出。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) InsuranceBillOCRWithContext(ctx context.Context, request *InsuranceBillOCRRequest) (response *InsuranceBillOCRResponse, err error) {
    if request == nil {
        request = NewInsuranceBillOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InsuranceBillOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewInsuranceBillOCRResponse()
    err = c.Send(request, response)
    return
}

func NewInvoiceGeneralOCRRequest() (request *InvoiceGeneralOCRRequest) {
    request = &InvoiceGeneralOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "InvoiceGeneralOCR")
    
    
    return
}

func NewInvoiceGeneralOCRResponse() (response *InvoiceGeneralOCRResponse) {
    response = &InvoiceGeneralOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InvoiceGeneralOCR
// 本接口支持对通用机打发票的发票代码、发票号码、日期、购买方识别号、销售方识别号、校验码、小写金额等关键字段的识别。
//
// 
//
// 默认接口请求频率限制：5次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) InvoiceGeneralOCR(request *InvoiceGeneralOCRRequest) (response *InvoiceGeneralOCRResponse, err error) {
    return c.InvoiceGeneralOCRWithContext(context.Background(), request)
}

// InvoiceGeneralOCR
// 本接口支持对通用机打发票的发票代码、发票号码、日期、购买方识别号、销售方识别号、校验码、小写金额等关键字段的识别。
//
// 
//
// 默认接口请求频率限制：5次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) InvoiceGeneralOCRWithContext(ctx context.Context, request *InvoiceGeneralOCRRequest) (response *InvoiceGeneralOCRResponse, err error) {
    if request == nil {
        request = NewInvoiceGeneralOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InvoiceGeneralOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewInvoiceGeneralOCRResponse()
    err = c.Send(request, response)
    return
}

func NewLicensePlateOCRRequest() (request *LicensePlateOCRRequest) {
    request = &LicensePlateOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "LicensePlateOCR")
    
    
    return
}

func NewLicensePlateOCRResponse() (response *LicensePlateOCRResponse) {
    response = &LicensePlateOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// LicensePlateOCR
// 本接口支持对中国大陆机动车车牌的自动定位和识别，返回地域编号和车牌号码与车牌颜色信息。
//
// 
//
// 默认接口请求频率限制：10次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) LicensePlateOCR(request *LicensePlateOCRRequest) (response *LicensePlateOCRResponse, err error) {
    return c.LicensePlateOCRWithContext(context.Background(), request)
}

// LicensePlateOCR
// 本接口支持对中国大陆机动车车牌的自动定位和识别，返回地域编号和车牌号码与车牌颜色信息。
//
// 
//
// 默认接口请求频率限制：10次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) LicensePlateOCRWithContext(ctx context.Context, request *LicensePlateOCRRequest) (response *LicensePlateOCRResponse, err error) {
    if request == nil {
        request = NewLicensePlateOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("LicensePlateOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewLicensePlateOCRResponse()
    err = c.Send(request, response)
    return
}

func NewMLIDCardOCRRequest() (request *MLIDCardOCRRequest) {
    request = &MLIDCardOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "MLIDCardOCR")
    
    
    return
}

func NewMLIDCardOCRResponse() (response *MLIDCardOCRResponse) {
    response = &MLIDCardOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// MLIDCardOCR
// 本接口支持马来西亚身份证识别，识别字段包括身份证号、姓名、性别、地址；具备身份证人像照片的裁剪功能和翻拍、复印件告警功能。
//
// 本接口暂未完全对外开放，如需咨询，请[联系商务](https://cloud.tencent.com/about/connect)
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_NOMASIDCARD = "FailedOperation.NoMASIDCard"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) MLIDCardOCR(request *MLIDCardOCRRequest) (response *MLIDCardOCRResponse, err error) {
    return c.MLIDCardOCRWithContext(context.Background(), request)
}

// MLIDCardOCR
// 本接口支持马来西亚身份证识别，识别字段包括身份证号、姓名、性别、地址；具备身份证人像照片的裁剪功能和翻拍、复印件告警功能。
//
// 本接口暂未完全对外开放，如需咨询，请[联系商务](https://cloud.tencent.com/about/connect)
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_NOMASIDCARD = "FailedOperation.NoMASIDCard"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) MLIDCardOCRWithContext(ctx context.Context, request *MLIDCardOCRRequest) (response *MLIDCardOCRResponse, err error) {
    if request == nil {
        request = NewMLIDCardOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("MLIDCardOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewMLIDCardOCRResponse()
    err = c.Send(request, response)
    return
}

func NewMLIDPassportOCRRequest() (request *MLIDPassportOCRRequest) {
    request = &MLIDPassportOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "MLIDPassportOCR")
    
    
    return
}

func NewMLIDPassportOCRResponse() (response *MLIDPassportOCRResponse) {
    response = &MLIDPassportOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// MLIDPassportOCR
// 本接口支持中国港澳台地区以及其他国家、地区的护照识别。识别字段包括护照ID、姓名、出生日期、性别、有效期、发行国、国籍，具备护照人像照片的裁剪功能和翻拍、复印件告警功能。
//
// 
//
// 默认接口请求频率限制：5次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_NOPASSPORT = "FailedOperation.NoPassport"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) MLIDPassportOCR(request *MLIDPassportOCRRequest) (response *MLIDPassportOCRResponse, err error) {
    return c.MLIDPassportOCRWithContext(context.Background(), request)
}

// MLIDPassportOCR
// 本接口支持中国港澳台地区以及其他国家、地区的护照识别。识别字段包括护照ID、姓名、出生日期、性别、有效期、发行国、国籍，具备护照人像照片的裁剪功能和翻拍、复印件告警功能。
//
// 
//
// 默认接口请求频率限制：5次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_NOPASSPORT = "FailedOperation.NoPassport"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) MLIDPassportOCRWithContext(ctx context.Context, request *MLIDPassportOCRRequest) (response *MLIDPassportOCRResponse, err error) {
    if request == nil {
        request = NewMLIDPassportOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("MLIDPassportOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewMLIDPassportOCRResponse()
    err = c.Send(request, response)
    return
}

func NewMainlandPermitOCRRequest() (request *MainlandPermitOCRRequest) {
    request = &MainlandPermitOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "MainlandPermitOCR")
    
    
    return
}

func NewMainlandPermitOCRResponse() (response *MainlandPermitOCRResponse) {
    response = &MainlandPermitOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// MainlandPermitOCR
// 智能识别并结构化港澳台居民来往内地通行证正面全部字段，包含中文姓名、英文姓名、性别、出生日期、签发机关、有效期限、证件号、签发地点、签发次数、证件类别。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) MainlandPermitOCR(request *MainlandPermitOCRRequest) (response *MainlandPermitOCRResponse, err error) {
    return c.MainlandPermitOCRWithContext(context.Background(), request)
}

// MainlandPermitOCR
// 智能识别并结构化港澳台居民来往内地通行证正面全部字段，包含中文姓名、英文姓名、性别、出生日期、签发机关、有效期限、证件号、签发地点、签发次数、证件类别。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) MainlandPermitOCRWithContext(ctx context.Context, request *MainlandPermitOCRRequest) (response *MainlandPermitOCRResponse, err error) {
    if request == nil {
        request = NewMainlandPermitOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("MainlandPermitOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewMainlandPermitOCRResponse()
    err = c.Send(request, response)
    return
}

func NewMixedInvoiceDetectRequest() (request *MixedInvoiceDetectRequest) {
    request = &MixedInvoiceDetectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "MixedInvoiceDetect")
    
    
    return
}

func NewMixedInvoiceDetectResponse() (response *MixedInvoiceDetectResponse) {
    response = &MixedInvoiceDetectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// MixedInvoiceDetect
// 本接口支持多张、多类型票据的混合检测和自动分类，返回对应票据类型。目前已支持增值税发票、增值税发票（卷票）、定额发票、通用机打发票、购车发票、火车票、出租车发票、机票行程单、汽车票、轮船票、过路过桥费发票、酒店账单、客运限额发票、购物小票、完税证明共15种票据。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) MixedInvoiceDetect(request *MixedInvoiceDetectRequest) (response *MixedInvoiceDetectResponse, err error) {
    return c.MixedInvoiceDetectWithContext(context.Background(), request)
}

// MixedInvoiceDetect
// 本接口支持多张、多类型票据的混合检测和自动分类，返回对应票据类型。目前已支持增值税发票、增值税发票（卷票）、定额发票、通用机打发票、购车发票、火车票、出租车发票、机票行程单、汽车票、轮船票、过路过桥费发票、酒店账单、客运限额发票、购物小票、完税证明共15种票据。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) MixedInvoiceDetectWithContext(ctx context.Context, request *MixedInvoiceDetectRequest) (response *MixedInvoiceDetectResponse, err error) {
    if request == nil {
        request = NewMixedInvoiceDetectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("MixedInvoiceDetect require credential")
    }

    request.SetContext(ctx)
    
    response = NewMixedInvoiceDetectResponse()
    err = c.Send(request, response)
    return
}

func NewMixedInvoiceOCRRequest() (request *MixedInvoiceOCRRequest) {
    request = &MixedInvoiceOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "MixedInvoiceOCR")
    
    
    return
}

func NewMixedInvoiceOCRResponse() (response *MixedInvoiceOCRResponse) {
    response = &MixedInvoiceOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// MixedInvoiceOCR
// 本接口支持 单张、多张、多类型 票据的混合识别，同时支持自选需要识别的票据类型，已支持票种包括：增值税发票（专票、普票、卷票）、全电发票、非税发票、定额发票、通用机打发票、购车发票、火车票、出租车发票、机票行程单、汽车票、轮船票、过路过桥费发票共14种标准报销发票，并支持其他类发票的识别。
//
// 
//
// 默认接口请求频率限制：5次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEBLUR = "FailedOperation.ImageBlur"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) MixedInvoiceOCR(request *MixedInvoiceOCRRequest) (response *MixedInvoiceOCRResponse, err error) {
    return c.MixedInvoiceOCRWithContext(context.Background(), request)
}

// MixedInvoiceOCR
// 本接口支持 单张、多张、多类型 票据的混合识别，同时支持自选需要识别的票据类型，已支持票种包括：增值税发票（专票、普票、卷票）、全电发票、非税发票、定额发票、通用机打发票、购车发票、火车票、出租车发票、机票行程单、汽车票、轮船票、过路过桥费发票共14种标准报销发票，并支持其他类发票的识别。
//
// 
//
// 默认接口请求频率限制：5次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEBLUR = "FailedOperation.ImageBlur"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) MixedInvoiceOCRWithContext(ctx context.Context, request *MixedInvoiceOCRRequest) (response *MixedInvoiceOCRResponse, err error) {
    if request == nil {
        request = NewMixedInvoiceOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("MixedInvoiceOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewMixedInvoiceOCRResponse()
    err = c.Send(request, response)
    return
}

func NewOrgCodeCertOCRRequest() (request *OrgCodeCertOCRRequest) {
    request = &OrgCodeCertOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "OrgCodeCertOCR")
    
    
    return
}

func NewOrgCodeCertOCRResponse() (response *OrgCodeCertOCRResponse) {
    response = &OrgCodeCertOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// OrgCodeCertOCR
// 本接口支持组织机构代码证关键字段的识别，包括代码、有效期、地址、机构名称等。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) OrgCodeCertOCR(request *OrgCodeCertOCRRequest) (response *OrgCodeCertOCRResponse, err error) {
    return c.OrgCodeCertOCRWithContext(context.Background(), request)
}

// OrgCodeCertOCR
// 本接口支持组织机构代码证关键字段的识别，包括代码、有效期、地址、机构名称等。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) OrgCodeCertOCRWithContext(ctx context.Context, request *OrgCodeCertOCRRequest) (response *OrgCodeCertOCRResponse, err error) {
    if request == nil {
        request = NewOrgCodeCertOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("OrgCodeCertOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewOrgCodeCertOCRResponse()
    err = c.Send(request, response)
    return
}

func NewPassportOCRRequest() (request *PassportOCRRequest) {
    request = &PassportOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "PassportOCR")
    
    
    return
}

func NewPassportOCRResponse() (response *PassportOCRResponse) {
    response = &PassportOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PassportOCR
// 本接口支持中国大陆地区护照个人资料页多个字段的检测与识别。已支持字段包括英文姓名、中文姓名、国家码、护照号、出生地、出生日期、国籍英文、性别英文、有效期、签发地点英文、签发日期、持证人签名、护照机读码（MRZ码）等。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) PassportOCR(request *PassportOCRRequest) (response *PassportOCRResponse, err error) {
    return c.PassportOCRWithContext(context.Background(), request)
}

// PassportOCR
// 本接口支持中国大陆地区护照个人资料页多个字段的检测与识别。已支持字段包括英文姓名、中文姓名、国家码、护照号、出生地、出生日期、国籍英文、性别英文、有效期、签发地点英文、签发日期、持证人签名、护照机读码（MRZ码）等。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) PassportOCRWithContext(ctx context.Context, request *PassportOCRRequest) (response *PassportOCRResponse, err error) {
    if request == nil {
        request = NewPassportOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PassportOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewPassportOCRResponse()
    err = c.Send(request, response)
    return
}

func NewPermitOCRRequest() (request *PermitOCRRequest) {
    request = &PermitOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "PermitOCR")
    
    
    return
}

func NewPermitOCRResponse() (response *PermitOCRResponse) {
    response = &PermitOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PermitOCR
// 本接口支持对卡式港澳台通行证的识别，包括签发地点、签发机关、有效期限、性别、出生日期、英文姓名、姓名、证件号等字段。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) PermitOCR(request *PermitOCRRequest) (response *PermitOCRResponse, err error) {
    return c.PermitOCRWithContext(context.Background(), request)
}

// PermitOCR
// 本接口支持对卡式港澳台通行证的识别，包括签发地点、签发机关、有效期限、性别、出生日期、英文姓名、姓名、证件号等字段。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) PermitOCRWithContext(ctx context.Context, request *PermitOCRRequest) (response *PermitOCRResponse, err error) {
    if request == nil {
        request = NewPermitOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PermitOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewPermitOCRResponse()
    err = c.Send(request, response)
    return
}

func NewPropOwnerCertOCRRequest() (request *PropOwnerCertOCRRequest) {
    request = &PropOwnerCertOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "PropOwnerCertOCR")
    
    
    return
}

func NewPropOwnerCertOCRResponse() (response *PropOwnerCertOCRResponse) {
    response = &PropOwnerCertOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PropOwnerCertOCR
// 本接口支持房产证关键字段的识别，包括房地产权利人、共有情况、登记时间、规划用途、房屋性质、房屋坐落等。
//
// 目前接口对合肥、成都、佛山三个城市的房产证版式识别较好。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_LANGUAGENOTSUPPORT = "FailedOperation.LanguageNotSupport"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) PropOwnerCertOCR(request *PropOwnerCertOCRRequest) (response *PropOwnerCertOCRResponse, err error) {
    return c.PropOwnerCertOCRWithContext(context.Background(), request)
}

// PropOwnerCertOCR
// 本接口支持房产证关键字段的识别，包括房地产权利人、共有情况、登记时间、规划用途、房屋性质、房屋坐落等。
//
// 目前接口对合肥、成都、佛山三个城市的房产证版式识别较好。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_LANGUAGENOTSUPPORT = "FailedOperation.LanguageNotSupport"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) PropOwnerCertOCRWithContext(ctx context.Context, request *PropOwnerCertOCRRequest) (response *PropOwnerCertOCRResponse, err error) {
    if request == nil {
        request = NewPropOwnerCertOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PropOwnerCertOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewPropOwnerCertOCRResponse()
    err = c.Send(request, response)
    return
}

func NewQrcodeOCRRequest() (request *QrcodeOCRRequest) {
    request = &QrcodeOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "QrcodeOCR")
    
    
    return
}

func NewQrcodeOCRResponse() (response *QrcodeOCRResponse) {
    response = &QrcodeOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QrcodeOCR
// 本接口支持条形码和二维码的识别（包括 DataMatrix 和 PDF417）。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEBLUR = "FailedOperation.ImageBlur"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) QrcodeOCR(request *QrcodeOCRRequest) (response *QrcodeOCRResponse, err error) {
    return c.QrcodeOCRWithContext(context.Background(), request)
}

// QrcodeOCR
// 本接口支持条形码和二维码的识别（包括 DataMatrix 和 PDF417）。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEBLUR = "FailedOperation.ImageBlur"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) QrcodeOCRWithContext(ctx context.Context, request *QrcodeOCRRequest) (response *QrcodeOCRResponse, err error) {
    if request == nil {
        request = NewQrcodeOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QrcodeOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewQrcodeOCRResponse()
    err = c.Send(request, response)
    return
}

func NewQueryBarCodeRequest() (request *QueryBarCodeRequest) {
    request = &QueryBarCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "QueryBarCode")
    
    
    return
}

func NewQueryBarCodeResponse() (response *QueryBarCodeResponse) {
    response = &QueryBarCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryBarCode
// 本接口支持条形码备案信息查询，返回条形码查询结果的相关信息，包括产品名称、产品英文名称、品牌名称、规格型号、宽度、高度、深度、关键字、产品描述、厂家名称、厂家地址、企业社会信用代码13个字段信息。
//
// 
//
// 产品优势：直联中国物品编码中心，查询结果更加准确、可靠。
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_INVALIDGTINERROR = "InvalidParameter.InvalidGTINError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) QueryBarCode(request *QueryBarCodeRequest) (response *QueryBarCodeResponse, err error) {
    return c.QueryBarCodeWithContext(context.Background(), request)
}

// QueryBarCode
// 本接口支持条形码备案信息查询，返回条形码查询结果的相关信息，包括产品名称、产品英文名称、品牌名称、规格型号、宽度、高度、深度、关键字、产品描述、厂家名称、厂家地址、企业社会信用代码13个字段信息。
//
// 
//
// 产品优势：直联中国物品编码中心，查询结果更加准确、可靠。
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_INVALIDGTINERROR = "InvalidParameter.InvalidGTINError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) QueryBarCodeWithContext(ctx context.Context, request *QueryBarCodeRequest) (response *QueryBarCodeResponse, err error) {
    if request == nil {
        request = NewQueryBarCodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryBarCode require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryBarCodeResponse()
    err = c.Send(request, response)
    return
}

func NewQuotaInvoiceOCRRequest() (request *QuotaInvoiceOCRRequest) {
    request = &QuotaInvoiceOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "QuotaInvoiceOCR")
    
    
    return
}

func NewQuotaInvoiceOCRResponse() (response *QuotaInvoiceOCRResponse) {
    response = &QuotaInvoiceOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QuotaInvoiceOCR
// 本接口支持定额发票的发票号码、发票代码、金额(大小写)、发票消费类型、地区及是否有公司印章等关键字段的识别。
//
// 
//
// 默认接口请求频率限制：5次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) QuotaInvoiceOCR(request *QuotaInvoiceOCRRequest) (response *QuotaInvoiceOCRResponse, err error) {
    return c.QuotaInvoiceOCRWithContext(context.Background(), request)
}

// QuotaInvoiceOCR
// 本接口支持定额发票的发票号码、发票代码、金额(大小写)、发票消费类型、地区及是否有公司印章等关键字段的识别。
//
// 
//
// 默认接口请求频率限制：5次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) QuotaInvoiceOCRWithContext(ctx context.Context, request *QuotaInvoiceOCRRequest) (response *QuotaInvoiceOCRResponse, err error) {
    if request == nil {
        request = NewQuotaInvoiceOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QuotaInvoiceOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewQuotaInvoiceOCRResponse()
    err = c.Send(request, response)
    return
}

func NewRecognizeContainerOCRRequest() (request *RecognizeContainerOCRRequest) {
    request = &RecognizeContainerOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "RecognizeContainerOCR")
    
    
    return
}

func NewRecognizeContainerOCRResponse() (response *RecognizeContainerOCRResponse) {
    response = &RecognizeContainerOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RecognizeContainerOCR
// 本接口支持集装箱箱门信息识别，识别字段包括集装箱箱号、类型、总重量、有效承重、容量、自身重量，具备集装箱箱号、类型不完整或者不清晰的告警功能。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizeContainerOCR(request *RecognizeContainerOCRRequest) (response *RecognizeContainerOCRResponse, err error) {
    return c.RecognizeContainerOCRWithContext(context.Background(), request)
}

// RecognizeContainerOCR
// 本接口支持集装箱箱门信息识别，识别字段包括集装箱箱号、类型、总重量、有效承重、容量、自身重量，具备集装箱箱号、类型不完整或者不清晰的告警功能。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizeContainerOCRWithContext(ctx context.Context, request *RecognizeContainerOCRRequest) (response *RecognizeContainerOCRResponse, err error) {
    if request == nil {
        request = NewRecognizeContainerOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecognizeContainerOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecognizeContainerOCRResponse()
    err = c.Send(request, response)
    return
}

func NewRecognizeHealthCodeOCRRequest() (request *RecognizeHealthCodeOCRRequest) {
    request = &RecognizeHealthCodeOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "RecognizeHealthCodeOCR")
    
    
    return
}

func NewRecognizeHealthCodeOCRResponse() (response *RecognizeHealthCodeOCRResponse) {
    response = &RecognizeHealthCodeOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RecognizeHealthCodeOCR
// 本接口支持北京、上海、广东、江苏、吉林、黑龙江、天津、辽宁、浙江、河南、四川、贵州、山东、安徽、福建、江西、湖北、湖南等省份健康码的识别，包括持码人姓名、持码人身份证号、健康码更新时间、健康码颜色、核酸检测结果、核酸检测间隔时长、核酸检测时间，疫苗接种信息，八个字段的识别结果输出。不同省市健康码显示的字段信息有所不同，上述字段的识别结果可能为空，以图片上具体展示的信息为准。
//
// 
//
// 默认接口请求频率限制：10次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizeHealthCodeOCR(request *RecognizeHealthCodeOCRRequest) (response *RecognizeHealthCodeOCRResponse, err error) {
    return c.RecognizeHealthCodeOCRWithContext(context.Background(), request)
}

// RecognizeHealthCodeOCR
// 本接口支持北京、上海、广东、江苏、吉林、黑龙江、天津、辽宁、浙江、河南、四川、贵州、山东、安徽、福建、江西、湖北、湖南等省份健康码的识别，包括持码人姓名、持码人身份证号、健康码更新时间、健康码颜色、核酸检测结果、核酸检测间隔时长、核酸检测时间，疫苗接种信息，八个字段的识别结果输出。不同省市健康码显示的字段信息有所不同，上述字段的识别结果可能为空，以图片上具体展示的信息为准。
//
// 
//
// 默认接口请求频率限制：10次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizeHealthCodeOCRWithContext(ctx context.Context, request *RecognizeHealthCodeOCRRequest) (response *RecognizeHealthCodeOCRResponse, err error) {
    if request == nil {
        request = NewRecognizeHealthCodeOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecognizeHealthCodeOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecognizeHealthCodeOCRResponse()
    err = c.Send(request, response)
    return
}

func NewRecognizeIndonesiaIDCardOCRRequest() (request *RecognizeIndonesiaIDCardOCRRequest) {
    request = &RecognizeIndonesiaIDCardOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "RecognizeIndonesiaIDCardOCR")
    
    
    return
}

func NewRecognizeIndonesiaIDCardOCRResponse() (response *RecognizeIndonesiaIDCardOCRResponse) {
    response = &RecognizeIndonesiaIDCardOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RecognizeIndonesiaIDCardOCR
// 印尼身份证识别
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizeIndonesiaIDCardOCR(request *RecognizeIndonesiaIDCardOCRRequest) (response *RecognizeIndonesiaIDCardOCRResponse, err error) {
    return c.RecognizeIndonesiaIDCardOCRWithContext(context.Background(), request)
}

// RecognizeIndonesiaIDCardOCR
// 印尼身份证识别
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizeIndonesiaIDCardOCRWithContext(ctx context.Context, request *RecognizeIndonesiaIDCardOCRRequest) (response *RecognizeIndonesiaIDCardOCRResponse, err error) {
    if request == nil {
        request = NewRecognizeIndonesiaIDCardOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecognizeIndonesiaIDCardOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecognizeIndonesiaIDCardOCRResponse()
    err = c.Send(request, response)
    return
}

func NewRecognizeMedicalInvoiceOCRRequest() (request *RecognizeMedicalInvoiceOCRRequest) {
    request = &RecognizeMedicalInvoiceOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "RecognizeMedicalInvoiceOCR")
    
    
    return
}

func NewRecognizeMedicalInvoiceOCRResponse() (response *RecognizeMedicalInvoiceOCRResponse) {
    response = &RecognizeMedicalInvoiceOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RecognizeMedicalInvoiceOCR
// 医疗发票识别目前支持全国统一门诊发票、全国统一住院发票、以及部分地方的门诊和住院发票的识别。
//
// 
//
// 默认接口请求频率限制：5次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_ENGINERECOGNIZETIMEOUT = "FailedOperation.EngineRecognizeTimeout"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizeMedicalInvoiceOCR(request *RecognizeMedicalInvoiceOCRRequest) (response *RecognizeMedicalInvoiceOCRResponse, err error) {
    return c.RecognizeMedicalInvoiceOCRWithContext(context.Background(), request)
}

// RecognizeMedicalInvoiceOCR
// 医疗发票识别目前支持全国统一门诊发票、全国统一住院发票、以及部分地方的门诊和住院发票的识别。
//
// 
//
// 默认接口请求频率限制：5次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_ENGINERECOGNIZETIMEOUT = "FailedOperation.EngineRecognizeTimeout"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizeMedicalInvoiceOCRWithContext(ctx context.Context, request *RecognizeMedicalInvoiceOCRRequest) (response *RecognizeMedicalInvoiceOCRResponse, err error) {
    if request == nil {
        request = NewRecognizeMedicalInvoiceOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecognizeMedicalInvoiceOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecognizeMedicalInvoiceOCRResponse()
    err = c.Send(request, response)
    return
}

func NewRecognizeOnlineTaxiItineraryOCRRequest() (request *RecognizeOnlineTaxiItineraryOCRRequest) {
    request = &RecognizeOnlineTaxiItineraryOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "RecognizeOnlineTaxiItineraryOCR")
    
    
    return
}

func NewRecognizeOnlineTaxiItineraryOCRResponse() (response *RecognizeOnlineTaxiItineraryOCRResponse) {
    response = &RecognizeOnlineTaxiItineraryOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RecognizeOnlineTaxiItineraryOCR
// 本接口支持网约车行程单关键字段的识别，包括行程起止日期、上车时间、起点、终点、里程、金额等字段。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizeOnlineTaxiItineraryOCR(request *RecognizeOnlineTaxiItineraryOCRRequest) (response *RecognizeOnlineTaxiItineraryOCRResponse, err error) {
    return c.RecognizeOnlineTaxiItineraryOCRWithContext(context.Background(), request)
}

// RecognizeOnlineTaxiItineraryOCR
// 本接口支持网约车行程单关键字段的识别，包括行程起止日期、上车时间、起点、终点、里程、金额等字段。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizeOnlineTaxiItineraryOCRWithContext(ctx context.Context, request *RecognizeOnlineTaxiItineraryOCRRequest) (response *RecognizeOnlineTaxiItineraryOCRResponse, err error) {
    if request == nil {
        request = NewRecognizeOnlineTaxiItineraryOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecognizeOnlineTaxiItineraryOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecognizeOnlineTaxiItineraryOCRResponse()
    err = c.Send(request, response)
    return
}

func NewRecognizePhilippinesDrivingLicenseOCRRequest() (request *RecognizePhilippinesDrivingLicenseOCRRequest) {
    request = &RecognizePhilippinesDrivingLicenseOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "RecognizePhilippinesDrivingLicenseOCR")
    
    
    return
}

func NewRecognizePhilippinesDrivingLicenseOCRResponse() (response *RecognizePhilippinesDrivingLicenseOCRResponse) {
    response = &RecognizePhilippinesDrivingLicenseOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RecognizePhilippinesDrivingLicenseOCR
// 菲律宾驾驶证识别
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizePhilippinesDrivingLicenseOCR(request *RecognizePhilippinesDrivingLicenseOCRRequest) (response *RecognizePhilippinesDrivingLicenseOCRResponse, err error) {
    return c.RecognizePhilippinesDrivingLicenseOCRWithContext(context.Background(), request)
}

// RecognizePhilippinesDrivingLicenseOCR
// 菲律宾驾驶证识别
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizePhilippinesDrivingLicenseOCRWithContext(ctx context.Context, request *RecognizePhilippinesDrivingLicenseOCRRequest) (response *RecognizePhilippinesDrivingLicenseOCRResponse, err error) {
    if request == nil {
        request = NewRecognizePhilippinesDrivingLicenseOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecognizePhilippinesDrivingLicenseOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecognizePhilippinesDrivingLicenseOCRResponse()
    err = c.Send(request, response)
    return
}

func NewRecognizePhilippinesVoteIDOCRRequest() (request *RecognizePhilippinesVoteIDOCRRequest) {
    request = &RecognizePhilippinesVoteIDOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "RecognizePhilippinesVoteIDOCR")
    
    
    return
}

func NewRecognizePhilippinesVoteIDOCRResponse() (response *RecognizePhilippinesVoteIDOCRResponse) {
    response = &RecognizePhilippinesVoteIDOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RecognizePhilippinesVoteIDOCR
// 菲律宾VoteID识别
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizePhilippinesVoteIDOCR(request *RecognizePhilippinesVoteIDOCRRequest) (response *RecognizePhilippinesVoteIDOCRResponse, err error) {
    return c.RecognizePhilippinesVoteIDOCRWithContext(context.Background(), request)
}

// RecognizePhilippinesVoteIDOCR
// 菲律宾VoteID识别
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizePhilippinesVoteIDOCRWithContext(ctx context.Context, request *RecognizePhilippinesVoteIDOCRRequest) (response *RecognizePhilippinesVoteIDOCRResponse, err error) {
    if request == nil {
        request = NewRecognizePhilippinesVoteIDOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecognizePhilippinesVoteIDOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecognizePhilippinesVoteIDOCRResponse()
    err = c.Send(request, response)
    return
}

func NewRecognizeTableAccurateOCRRequest() (request *RecognizeTableAccurateOCRRequest) {
    request = &RecognizeTableAccurateOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "RecognizeTableAccurateOCR")
    
    
    return
}

func NewRecognizeTableAccurateOCRResponse() (response *RecognizeTableAccurateOCRResponse) {
    response = &RecognizeTableAccurateOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RecognizeTableAccurateOCR
// 本接口支持中英文图片/PDF内常规表格、无线表格、多表格的检测和识别，返回每个单元格的文字内容，支持旋转的表格图片识别，且支持将识别结果保存为 Excel 格式。识别效果比表格识别V2更好，覆盖场景更加广泛，对表格难例场景，如无线表格、嵌套表格（有线表格中包含无线表格）的识别效果均优于表格识别V2。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizeTableAccurateOCR(request *RecognizeTableAccurateOCRRequest) (response *RecognizeTableAccurateOCRResponse, err error) {
    return c.RecognizeTableAccurateOCRWithContext(context.Background(), request)
}

// RecognizeTableAccurateOCR
// 本接口支持中英文图片/PDF内常规表格、无线表格、多表格的检测和识别，返回每个单元格的文字内容，支持旋转的表格图片识别，且支持将识别结果保存为 Excel 格式。识别效果比表格识别V2更好，覆盖场景更加广泛，对表格难例场景，如无线表格、嵌套表格（有线表格中包含无线表格）的识别效果均优于表格识别V2。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizeTableAccurateOCRWithContext(ctx context.Context, request *RecognizeTableAccurateOCRRequest) (response *RecognizeTableAccurateOCRResponse, err error) {
    if request == nil {
        request = NewRecognizeTableAccurateOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecognizeTableAccurateOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecognizeTableAccurateOCRResponse()
    err = c.Send(request, response)
    return
}

func NewRecognizeTableOCRRequest() (request *RecognizeTableOCRRequest) {
    request = &RecognizeTableOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "RecognizeTableOCR")
    
    
    return
}

func NewRecognizeTableOCRResponse() (response *RecognizeTableOCRResponse) {
    response = &RecognizeTableOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RecognizeTableOCR
// 本接口支持中英文图片/ PDF内常规表格、无线表格、多表格的检测和识别，支持日文有线表格识别，返回每个单元格的文字内容，支持旋转的表格图片识别，且支持将识别结果保存为 Excel 格式。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizeTableOCR(request *RecognizeTableOCRRequest) (response *RecognizeTableOCRResponse, err error) {
    return c.RecognizeTableOCRWithContext(context.Background(), request)
}

// RecognizeTableOCR
// 本接口支持中英文图片/ PDF内常规表格、无线表格、多表格的检测和识别，支持日文有线表格识别，返回每个单元格的文字内容，支持旋转的表格图片识别，且支持将识别结果保存为 Excel 格式。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizeTableOCRWithContext(ctx context.Context, request *RecognizeTableOCRRequest) (response *RecognizeTableOCRResponse, err error) {
    if request == nil {
        request = NewRecognizeTableOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecognizeTableOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecognizeTableOCRResponse()
    err = c.Send(request, response)
    return
}

func NewRecognizeThaiIDCardOCRRequest() (request *RecognizeThaiIDCardOCRRequest) {
    request = &RecognizeThaiIDCardOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "RecognizeThaiIDCardOCR")
    
    
    return
}

func NewRecognizeThaiIDCardOCRResponse() (response *RecognizeThaiIDCardOCRResponse) {
    response = &RecognizeThaiIDCardOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RecognizeThaiIDCardOCR
// 本接口支持泰国身份证识别，识别字段包括泰文姓名、英文姓名、地址、出生日期、身份证号码。
//
// 本接口暂未完全对外开放，如需咨询，请[联系商务](https://cloud.tencent.com/about/connect)
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEBLUR = "FailedOperation.ImageBlur"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOSPECIFIEDCARD = "FailedOperation.ImageNoSpecifiedCard"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizeThaiIDCardOCR(request *RecognizeThaiIDCardOCRRequest) (response *RecognizeThaiIDCardOCRResponse, err error) {
    return c.RecognizeThaiIDCardOCRWithContext(context.Background(), request)
}

// RecognizeThaiIDCardOCR
// 本接口支持泰国身份证识别，识别字段包括泰文姓名、英文姓名、地址、出生日期、身份证号码。
//
// 本接口暂未完全对外开放，如需咨询，请[联系商务](https://cloud.tencent.com/about/connect)
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEBLUR = "FailedOperation.ImageBlur"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOSPECIFIEDCARD = "FailedOperation.ImageNoSpecifiedCard"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizeThaiIDCardOCRWithContext(ctx context.Context, request *RecognizeThaiIDCardOCRRequest) (response *RecognizeThaiIDCardOCRResponse, err error) {
    if request == nil {
        request = NewRecognizeThaiIDCardOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecognizeThaiIDCardOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecognizeThaiIDCardOCRResponse()
    err = c.Send(request, response)
    return
}

func NewRecognizeTravelCardOCRRequest() (request *RecognizeTravelCardOCRRequest) {
    request = &RecognizeTravelCardOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "RecognizeTravelCardOCR")
    
    
    return
}

func NewRecognizeTravelCardOCRResponse() (response *RecognizeTravelCardOCRResponse) {
    response = &RecognizeTravelCardOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RecognizeTravelCardOCR
// 本接口支持通信大数据行程卡识别，包括行程卡颜色、更新时间、途经地、存在中高风险地区的城市、电话号码，五个字段的识别结果输出。
//
// 
//
// 默认接口请求频率限制：20次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizeTravelCardOCR(request *RecognizeTravelCardOCRRequest) (response *RecognizeTravelCardOCRResponse, err error) {
    return c.RecognizeTravelCardOCRWithContext(context.Background(), request)
}

// RecognizeTravelCardOCR
// 本接口支持通信大数据行程卡识别，包括行程卡颜色、更新时间、途经地、存在中高风险地区的城市、电话号码，五个字段的识别结果输出。
//
// 
//
// 默认接口请求频率限制：20次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizeTravelCardOCRWithContext(ctx context.Context, request *RecognizeTravelCardOCRRequest) (response *RecognizeTravelCardOCRResponse, err error) {
    if request == nil {
        request = NewRecognizeTravelCardOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecognizeTravelCardOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecognizeTravelCardOCRResponse()
    err = c.Send(request, response)
    return
}

func NewResidenceBookletOCRRequest() (request *ResidenceBookletOCRRequest) {
    request = &ResidenceBookletOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "ResidenceBookletOCR")
    
    
    return
}

func NewResidenceBookletOCRResponse() (response *ResidenceBookletOCRResponse) {
    response = &ResidenceBookletOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResidenceBookletOCR
// 本接口支持居民户口簿户主页及成员页关键字段的识别，包括姓名、户别、地址、籍贯、身份证号码等。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_LANGUAGENOTSUPPORT = "FailedOperation.LanguageNotSupport"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) ResidenceBookletOCR(request *ResidenceBookletOCRRequest) (response *ResidenceBookletOCRResponse, err error) {
    return c.ResidenceBookletOCRWithContext(context.Background(), request)
}

// ResidenceBookletOCR
// 本接口支持居民户口簿户主页及成员页关键字段的识别，包括姓名、户别、地址、籍贯、身份证号码等。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_LANGUAGENOTSUPPORT = "FailedOperation.LanguageNotSupport"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) ResidenceBookletOCRWithContext(ctx context.Context, request *ResidenceBookletOCRRequest) (response *ResidenceBookletOCRResponse, err error) {
    if request == nil {
        request = NewResidenceBookletOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResidenceBookletOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewResidenceBookletOCRResponse()
    err = c.Send(request, response)
    return
}

func NewRideHailingDriverLicenseOCRRequest() (request *RideHailingDriverLicenseOCRRequest) {
    request = &RideHailingDriverLicenseOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "RideHailingDriverLicenseOCR")
    
    
    return
}

func NewRideHailingDriverLicenseOCRResponse() (response *RideHailingDriverLicenseOCRResponse) {
    response = &RideHailingDriverLicenseOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RideHailingDriverLicenseOCR
// 本接口支持网约车驾驶证关键字段的识别，包括姓名、证号、起始日期、截止日期、发证日期。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RideHailingDriverLicenseOCR(request *RideHailingDriverLicenseOCRRequest) (response *RideHailingDriverLicenseOCRResponse, err error) {
    return c.RideHailingDriverLicenseOCRWithContext(context.Background(), request)
}

// RideHailingDriverLicenseOCR
// 本接口支持网约车驾驶证关键字段的识别，包括姓名、证号、起始日期、截止日期、发证日期。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RideHailingDriverLicenseOCRWithContext(ctx context.Context, request *RideHailingDriverLicenseOCRRequest) (response *RideHailingDriverLicenseOCRResponse, err error) {
    if request == nil {
        request = NewRideHailingDriverLicenseOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RideHailingDriverLicenseOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewRideHailingDriverLicenseOCRResponse()
    err = c.Send(request, response)
    return
}

func NewRideHailingTransportLicenseOCRRequest() (request *RideHailingTransportLicenseOCRRequest) {
    request = &RideHailingTransportLicenseOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "RideHailingTransportLicenseOCR")
    
    
    return
}

func NewRideHailingTransportLicenseOCRResponse() (response *RideHailingTransportLicenseOCRResponse) {
    response = &RideHailingTransportLicenseOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RideHailingTransportLicenseOCR
// 本接口支持网约车运输证关键字段的识别，包括交运管许可字号、车辆所有人、车辆号牌、起始日期、截止日期、发证日期。
//
//            
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RideHailingTransportLicenseOCR(request *RideHailingTransportLicenseOCRRequest) (response *RideHailingTransportLicenseOCRResponse, err error) {
    return c.RideHailingTransportLicenseOCRWithContext(context.Background(), request)
}

// RideHailingTransportLicenseOCR
// 本接口支持网约车运输证关键字段的识别，包括交运管许可字号、车辆所有人、车辆号牌、起始日期、截止日期、发证日期。
//
//            
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RideHailingTransportLicenseOCRWithContext(ctx context.Context, request *RideHailingTransportLicenseOCRRequest) (response *RideHailingTransportLicenseOCRResponse, err error) {
    if request == nil {
        request = NewRideHailingTransportLicenseOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RideHailingTransportLicenseOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewRideHailingTransportLicenseOCRResponse()
    err = c.Send(request, response)
    return
}

func NewSealOCRRequest() (request *SealOCRRequest) {
    request = &SealOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "SealOCR")
    
    
    return
}

func NewSealOCRResponse() (response *SealOCRResponse) {
    response = &SealOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SealOCR
// 本接口支持各类印章识别，包括发票章，财务章等，适用于公文，票据等场景。
//
// 
//
// 默认接口请求频率限制：5次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) SealOCR(request *SealOCRRequest) (response *SealOCRResponse, err error) {
    return c.SealOCRWithContext(context.Background(), request)
}

// SealOCR
// 本接口支持各类印章识别，包括发票章，财务章等，适用于公文，票据等场景。
//
// 
//
// 默认接口请求频率限制：5次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) SealOCRWithContext(ctx context.Context, request *SealOCRRequest) (response *SealOCRResponse, err error) {
    if request == nil {
        request = NewSealOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SealOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewSealOCRResponse()
    err = c.Send(request, response)
    return
}

func NewShipInvoiceOCRRequest() (request *ShipInvoiceOCRRequest) {
    request = &ShipInvoiceOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "ShipInvoiceOCR")
    
    
    return
}

func NewShipInvoiceOCRResponse() (response *ShipInvoiceOCRResponse) {
    response = &ShipInvoiceOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ShipInvoiceOCR
// 本接口支持识别轮船票的发票代码、发票号码、日期、姓名、票价、始发地、目的地、姓名、时间、发票消费类型、省、市、币种字段。
//
// 
//
// 默认接口请求频率限制：5次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) ShipInvoiceOCR(request *ShipInvoiceOCRRequest) (response *ShipInvoiceOCRResponse, err error) {
    return c.ShipInvoiceOCRWithContext(context.Background(), request)
}

// ShipInvoiceOCR
// 本接口支持识别轮船票的发票代码、发票号码、日期、姓名、票价、始发地、目的地、姓名、时间、发票消费类型、省、市、币种字段。
//
// 
//
// 默认接口请求频率限制：5次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) ShipInvoiceOCRWithContext(ctx context.Context, request *ShipInvoiceOCRRequest) (response *ShipInvoiceOCRResponse, err error) {
    if request == nil {
        request = NewShipInvoiceOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ShipInvoiceOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewShipInvoiceOCRResponse()
    err = c.Send(request, response)
    return
}

func NewSmartStructuralOCRRequest() (request *SmartStructuralOCRRequest) {
    request = &SmartStructuralOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "SmartStructuralOCR")
    
    
    return
}

func NewSmartStructuralOCRResponse() (response *SmartStructuralOCRResponse) {
    response = &SmartStructuralOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SmartStructuralOCR
// 本接口支持识别并提取各类证照、票据、表单、合同等结构化场景的字段信息。无需任何配置，灵活高效。适用于各类结构化信息录入场景。
//
// 
//
// 默认接口请求频率限制：5次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) SmartStructuralOCR(request *SmartStructuralOCRRequest) (response *SmartStructuralOCRResponse, err error) {
    return c.SmartStructuralOCRWithContext(context.Background(), request)
}

// SmartStructuralOCR
// 本接口支持识别并提取各类证照、票据、表单、合同等结构化场景的字段信息。无需任何配置，灵活高效。适用于各类结构化信息录入场景。
//
// 
//
// 默认接口请求频率限制：5次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) SmartStructuralOCRWithContext(ctx context.Context, request *SmartStructuralOCRRequest) (response *SmartStructuralOCRResponse, err error) {
    if request == nil {
        request = NewSmartStructuralOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SmartStructuralOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewSmartStructuralOCRResponse()
    err = c.Send(request, response)
    return
}

func NewTableOCRRequest() (request *TableOCRRequest) {
    request = &TableOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "TableOCR")
    
    
    return
}

func NewTableOCRResponse() (response *TableOCRResponse) {
    response = &TableOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TableOCR
// <b>此接口为表格识别的旧版本服务，不再进行服务升级，建议您使用识别能力更强、服务性能更优的<a href="https://cloud.tencent.com/document/product/866/49525">新版表格识别</a>。</b>
//
// 
//
// 本接口支持图片内表格文档的检测和识别，返回每个单元格的文字内容，支持将识别结果保存为 Excel 格式。
//
// 
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_LANGUAGENOTSUPPORT = "FailedOperation.LanguageNotSupport"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) TableOCR(request *TableOCRRequest) (response *TableOCRResponse, err error) {
    return c.TableOCRWithContext(context.Background(), request)
}

// TableOCR
// <b>此接口为表格识别的旧版本服务，不再进行服务升级，建议您使用识别能力更强、服务性能更优的<a href="https://cloud.tencent.com/document/product/866/49525">新版表格识别</a>。</b>
//
// 
//
// 本接口支持图片内表格文档的检测和识别，返回每个单元格的文字内容，支持将识别结果保存为 Excel 格式。
//
// 
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_LANGUAGENOTSUPPORT = "FailedOperation.LanguageNotSupport"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) TableOCRWithContext(ctx context.Context, request *TableOCRRequest) (response *TableOCRResponse, err error) {
    if request == nil {
        request = NewTableOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TableOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewTableOCRResponse()
    err = c.Send(request, response)
    return
}

func NewTaxiInvoiceOCRRequest() (request *TaxiInvoiceOCRRequest) {
    request = &TaxiInvoiceOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "TaxiInvoiceOCR")
    
    
    return
}

func NewTaxiInvoiceOCRResponse() (response *TaxiInvoiceOCRResponse) {
    response = &TaxiInvoiceOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TaxiInvoiceOCR
// 本接口支持出租车发票关键字段的识别，包括发票号码、发票代码、金额、日期、上下车时间、里程、车牌号、发票类型及所属地区等字段。
//
// 
//
// 默认接口请求频率限制：5次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) TaxiInvoiceOCR(request *TaxiInvoiceOCRRequest) (response *TaxiInvoiceOCRResponse, err error) {
    return c.TaxiInvoiceOCRWithContext(context.Background(), request)
}

// TaxiInvoiceOCR
// 本接口支持出租车发票关键字段的识别，包括发票号码、发票代码、金额、日期、上下车时间、里程、车牌号、发票类型及所属地区等字段。
//
// 
//
// 默认接口请求频率限制：5次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) TaxiInvoiceOCRWithContext(ctx context.Context, request *TaxiInvoiceOCRRequest) (response *TaxiInvoiceOCRResponse, err error) {
    if request == nil {
        request = NewTaxiInvoiceOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TaxiInvoiceOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewTaxiInvoiceOCRResponse()
    err = c.Send(request, response)
    return
}

func NewTextDetectRequest() (request *TextDetectRequest) {
    request = &TextDetectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "TextDetect")
    
    
    return
}

func NewTextDetectResponse() (response *TextDetectResponse) {
    response = &TextDetectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TextDetect
// 本接口通过检测图片中的文字信息特征，快速判断图片中有无文字并返回判断结果，帮助用户过滤无文字的图片。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DETECTFAILED = "FailedOperation.DetectFailed"
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) TextDetect(request *TextDetectRequest) (response *TextDetectResponse, err error) {
    return c.TextDetectWithContext(context.Background(), request)
}

// TextDetect
// 本接口通过检测图片中的文字信息特征，快速判断图片中有无文字并返回判断结果，帮助用户过滤无文字的图片。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DETECTFAILED = "FailedOperation.DetectFailed"
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) TextDetectWithContext(ctx context.Context, request *TextDetectRequest) (response *TextDetectResponse, err error) {
    if request == nil {
        request = NewTextDetectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TextDetect require credential")
    }

    request.SetContext(ctx)
    
    response = NewTextDetectResponse()
    err = c.Send(request, response)
    return
}

func NewTollInvoiceOCRRequest() (request *TollInvoiceOCRRequest) {
    request = &TollInvoiceOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "TollInvoiceOCR")
    
    
    return
}

func NewTollInvoiceOCRResponse() (response *TollInvoiceOCRResponse) {
    response = &TollInvoiceOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TollInvoiceOCR
// 本接口支持对过路过桥费发票的发票代码、发票号码、日期、小写金额等关键字段的识别。
//
// 
//
// 默认接口请求频率限制：5次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) TollInvoiceOCR(request *TollInvoiceOCRRequest) (response *TollInvoiceOCRResponse, err error) {
    return c.TollInvoiceOCRWithContext(context.Background(), request)
}

// TollInvoiceOCR
// 本接口支持对过路过桥费发票的发票代码、发票号码、日期、小写金额等关键字段的识别。
//
// 
//
// 默认接口请求频率限制：5次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) TollInvoiceOCRWithContext(ctx context.Context, request *TollInvoiceOCRRequest) (response *TollInvoiceOCRResponse, err error) {
    if request == nil {
        request = NewTollInvoiceOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TollInvoiceOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewTollInvoiceOCRResponse()
    err = c.Send(request, response)
    return
}

func NewTrainTicketOCRRequest() (request *TrainTicketOCRRequest) {
    request = &TrainTicketOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "TrainTicketOCR")
    
    
    return
}

func NewTrainTicketOCRResponse() (response *TrainTicketOCRResponse) {
    response = &TrainTicketOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TrainTicketOCR
// 本接口支持火车票全字段的识别，包括编号、票价、姓名、座位号、出发时间、出发站、到达站、车次、席别、发票类型及序列号等。
//
// 
//
// 默认接口请求频率限制：5次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) TrainTicketOCR(request *TrainTicketOCRRequest) (response *TrainTicketOCRResponse, err error) {
    return c.TrainTicketOCRWithContext(context.Background(), request)
}

// TrainTicketOCR
// 本接口支持火车票全字段的识别，包括编号、票价、姓名、座位号、出发时间、出发站、到达站、车次、席别、发票类型及序列号等。
//
// 
//
// 默认接口请求频率限制：5次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) TrainTicketOCRWithContext(ctx context.Context, request *TrainTicketOCRRequest) (response *TrainTicketOCRResponse, err error) {
    if request == nil {
        request = NewTrainTicketOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TrainTicketOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewTrainTicketOCRResponse()
    err = c.Send(request, response)
    return
}

func NewVatInvoiceOCRRequest() (request *VatInvoiceOCRRequest) {
    request = &VatInvoiceOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "VatInvoiceOCR")
    
    
    return
}

func NewVatInvoiceOCRResponse() (response *VatInvoiceOCRResponse) {
    response = &VatInvoiceOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// VatInvoiceOCR
// 本接口支持增值税专用发票、增值税普通发票、增值税电子发票全字段的内容检测和识别，包括发票代码、发票号码、打印发票代码、打印发票号码、开票日期、合计金额、校验码、税率、合计税额、价税合计、购买方识别号、复核、销售方识别号、开票人、密码区1、密码区2、密码区3、密码区4、发票名称、购买方名称、销售方名称、服务名称、备注、规格型号、数量、单价、金额、税额、收款人等字段。
//
// 
//
// 默认接口请求频率限制：10次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEBLUR = "FailedOperation.ImageBlur"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) VatInvoiceOCR(request *VatInvoiceOCRRequest) (response *VatInvoiceOCRResponse, err error) {
    return c.VatInvoiceOCRWithContext(context.Background(), request)
}

// VatInvoiceOCR
// 本接口支持增值税专用发票、增值税普通发票、增值税电子发票全字段的内容检测和识别，包括发票代码、发票号码、打印发票代码、打印发票号码、开票日期、合计金额、校验码、税率、合计税额、价税合计、购买方识别号、复核、销售方识别号、开票人、密码区1、密码区2、密码区3、密码区4、发票名称、购买方名称、销售方名称、服务名称、备注、规格型号、数量、单价、金额、税额、收款人等字段。
//
// 
//
// 默认接口请求频率限制：10次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEBLUR = "FailedOperation.ImageBlur"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) VatInvoiceOCRWithContext(ctx context.Context, request *VatInvoiceOCRRequest) (response *VatInvoiceOCRResponse, err error) {
    if request == nil {
        request = NewVatInvoiceOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("VatInvoiceOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewVatInvoiceOCRResponse()
    err = c.Send(request, response)
    return
}

func NewVatInvoiceVerifyRequest() (request *VatInvoiceVerifyRequest) {
    request = &VatInvoiceVerifyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "VatInvoiceVerify")
    
    
    return
}

func NewVatInvoiceVerifyResponse() (response *VatInvoiceVerifyResponse) {
    response = &VatInvoiceVerifyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// VatInvoiceVerify
// 本接口支持增值税发票的准确性核验，您可以通过输入增值税发票的关键字段提供所需的验证信息，接口返回真实的票面相关信息，包括发票代码、发票号码、开票日期、金额、消费类型、购方名称、购方税号、销方名称、销方税号等多个常用字段。支持多种发票类型核验，包括增值税专用发票、增值税普通发票（含电子普通发票、卷式发票、通行费发票）、全电发票、机动车销售统一发票、货物运输业增值税专用发票、二手车销售统一发票。
//
// 
//
// 默认接口请求频率限制：20次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ARREARSERROR = "FailedOperation.ArrearsError"
//  FAILEDOPERATION_COUNTLIMITERROR = "FailedOperation.CountLimitError"
//  FAILEDOPERATION_INVOICEMISMATCH = "FailedOperation.InvoiceMismatch"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  INVALIDPARAMETERVALUE_PRICEORVERIFICATIONPARAMETERVALUELIMIT = "InvalidParameterValue.PriceOrVerificationParameterValueLimit"
//  INVALIDPARAMETERVALUE_TICKETDATEPARAMETERVALUELIMIT = "InvalidParameterValue.TicketDateParameterValueLimit"
//  RESOURCENOTFOUND_NOINVOICE = "ResourceNotFound.NoInvoice"
//  RESOURCENOTFOUND_NOTSUPPORTCURRENTINVOICEQUERY = "ResourceNotFound.NotSupportCurrentInvoiceQuery"
//  RESOURCEUNAVAILABLE_TAXNETWORKERROR = "ResourceUnavailable.TaxNetworkError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) VatInvoiceVerify(request *VatInvoiceVerifyRequest) (response *VatInvoiceVerifyResponse, err error) {
    return c.VatInvoiceVerifyWithContext(context.Background(), request)
}

// VatInvoiceVerify
// 本接口支持增值税发票的准确性核验，您可以通过输入增值税发票的关键字段提供所需的验证信息，接口返回真实的票面相关信息，包括发票代码、发票号码、开票日期、金额、消费类型、购方名称、购方税号、销方名称、销方税号等多个常用字段。支持多种发票类型核验，包括增值税专用发票、增值税普通发票（含电子普通发票、卷式发票、通行费发票）、全电发票、机动车销售统一发票、货物运输业增值税专用发票、二手车销售统一发票。
//
// 
//
// 默认接口请求频率限制：20次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ARREARSERROR = "FailedOperation.ArrearsError"
//  FAILEDOPERATION_COUNTLIMITERROR = "FailedOperation.CountLimitError"
//  FAILEDOPERATION_INVOICEMISMATCH = "FailedOperation.InvoiceMismatch"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  INVALIDPARAMETERVALUE_PRICEORVERIFICATIONPARAMETERVALUELIMIT = "InvalidParameterValue.PriceOrVerificationParameterValueLimit"
//  INVALIDPARAMETERVALUE_TICKETDATEPARAMETERVALUELIMIT = "InvalidParameterValue.TicketDateParameterValueLimit"
//  RESOURCENOTFOUND_NOINVOICE = "ResourceNotFound.NoInvoice"
//  RESOURCENOTFOUND_NOTSUPPORTCURRENTINVOICEQUERY = "ResourceNotFound.NotSupportCurrentInvoiceQuery"
//  RESOURCEUNAVAILABLE_TAXNETWORKERROR = "ResourceUnavailable.TaxNetworkError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) VatInvoiceVerifyWithContext(ctx context.Context, request *VatInvoiceVerifyRequest) (response *VatInvoiceVerifyResponse, err error) {
    if request == nil {
        request = NewVatInvoiceVerifyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("VatInvoiceVerify require credential")
    }

    request.SetContext(ctx)
    
    response = NewVatInvoiceVerifyResponse()
    err = c.Send(request, response)
    return
}

func NewVatInvoiceVerifyNewRequest() (request *VatInvoiceVerifyNewRequest) {
    request = &VatInvoiceVerifyNewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "VatInvoiceVerifyNew")
    
    
    return
}

func NewVatInvoiceVerifyNewResponse() (response *VatInvoiceVerifyNewResponse) {
    response = &VatInvoiceVerifyNewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// VatInvoiceVerifyNew
// 本接口支持增值税发票的准确性核验，您可以通过输入增值税发票的关键字段提供所需的验证信息，接口返回真实的票面相关信息，包括发票代码、发票号码、开票日期、金额、消费类型、购方名称、购方税号、销方名称、销方税号等多个常用字段。支持多种发票类型核验，包括增值税专用发票、增值税普通发票（含电子普通发票、卷式发票、通行费发票）、全电发票、机动车销售统一发票、货物运输业增值税专用发票、二手车销售统一发票、通用机打电子发票（广东和浙江）。
//
// 
//
// 默认接口请求频率限制：20次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ARREARSERROR = "FailedOperation.ArrearsError"
//  FAILEDOPERATION_COUNTLIMITERROR = "FailedOperation.CountLimitError"
//  FAILEDOPERATION_INVOICEMISMATCH = "FailedOperation.InvoiceMismatch"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  INVALIDPARAMETERVALUE_PRICEORVERIFICATIONPARAMETERVALUELIMIT = "InvalidParameterValue.PriceOrVerificationParameterValueLimit"
//  INVALIDPARAMETERVALUE_TICKETCODEPARAMETERVALUELIMIT = "InvalidParameterValue.TicketCodeParameterValueLimit"
//  INVALIDPARAMETERVALUE_TICKETDATEPARAMETERVALUELIMIT = "InvalidParameterValue.TicketDateParameterValueLimit"
//  INVALIDPARAMETERVALUE_TICKETSNPARAMETERVALUELIMIT = "InvalidParameterValue.TicketSnParameterValueLimit"
//  RESOURCENOTFOUND_NOAREACODE = "ResourceNotFound.NoAreaCode"
//  RESOURCENOTFOUND_NOINVOICE = "ResourceNotFound.NoInvoice"
//  RESOURCENOTFOUND_NOTSUPPORTCURRENTINVOICEQUERY = "ResourceNotFound.NotSupportCurrentInvoiceQuery"
//  RESOURCEUNAVAILABLE_TAXNETWORKERROR = "ResourceUnavailable.TaxNetworkError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) VatInvoiceVerifyNew(request *VatInvoiceVerifyNewRequest) (response *VatInvoiceVerifyNewResponse, err error) {
    return c.VatInvoiceVerifyNewWithContext(context.Background(), request)
}

// VatInvoiceVerifyNew
// 本接口支持增值税发票的准确性核验，您可以通过输入增值税发票的关键字段提供所需的验证信息，接口返回真实的票面相关信息，包括发票代码、发票号码、开票日期、金额、消费类型、购方名称、购方税号、销方名称、销方税号等多个常用字段。支持多种发票类型核验，包括增值税专用发票、增值税普通发票（含电子普通发票、卷式发票、通行费发票）、全电发票、机动车销售统一发票、货物运输业增值税专用发票、二手车销售统一发票、通用机打电子发票（广东和浙江）。
//
// 
//
// 默认接口请求频率限制：20次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ARREARSERROR = "FailedOperation.ArrearsError"
//  FAILEDOPERATION_COUNTLIMITERROR = "FailedOperation.CountLimitError"
//  FAILEDOPERATION_INVOICEMISMATCH = "FailedOperation.InvoiceMismatch"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  INVALIDPARAMETERVALUE_PRICEORVERIFICATIONPARAMETERVALUELIMIT = "InvalidParameterValue.PriceOrVerificationParameterValueLimit"
//  INVALIDPARAMETERVALUE_TICKETCODEPARAMETERVALUELIMIT = "InvalidParameterValue.TicketCodeParameterValueLimit"
//  INVALIDPARAMETERVALUE_TICKETDATEPARAMETERVALUELIMIT = "InvalidParameterValue.TicketDateParameterValueLimit"
//  INVALIDPARAMETERVALUE_TICKETSNPARAMETERVALUELIMIT = "InvalidParameterValue.TicketSnParameterValueLimit"
//  RESOURCENOTFOUND_NOAREACODE = "ResourceNotFound.NoAreaCode"
//  RESOURCENOTFOUND_NOINVOICE = "ResourceNotFound.NoInvoice"
//  RESOURCENOTFOUND_NOTSUPPORTCURRENTINVOICEQUERY = "ResourceNotFound.NotSupportCurrentInvoiceQuery"
//  RESOURCEUNAVAILABLE_TAXNETWORKERROR = "ResourceUnavailable.TaxNetworkError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) VatInvoiceVerifyNewWithContext(ctx context.Context, request *VatInvoiceVerifyNewRequest) (response *VatInvoiceVerifyNewResponse, err error) {
    if request == nil {
        request = NewVatInvoiceVerifyNewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("VatInvoiceVerifyNew require credential")
    }

    request.SetContext(ctx)
    
    response = NewVatInvoiceVerifyNewResponse()
    err = c.Send(request, response)
    return
}

func NewVatRollInvoiceOCRRequest() (request *VatRollInvoiceOCRRequest) {
    request = &VatRollInvoiceOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "VatRollInvoiceOCR")
    
    
    return
}

func NewVatRollInvoiceOCRResponse() (response *VatRollInvoiceOCRResponse) {
    response = &VatRollInvoiceOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// VatRollInvoiceOCR
// 本接口支持对增值税发票（卷票）的发票代码、发票号码、日期、校验码、合计金额（小写）等关键字段的识别。
//
// 
//
// 默认接口请求频率限制：5次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) VatRollInvoiceOCR(request *VatRollInvoiceOCRRequest) (response *VatRollInvoiceOCRResponse, err error) {
    return c.VatRollInvoiceOCRWithContext(context.Background(), request)
}

// VatRollInvoiceOCR
// 本接口支持对增值税发票（卷票）的发票代码、发票号码、日期、校验码、合计金额（小写）等关键字段的识别。
//
// 
//
// 默认接口请求频率限制：5次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) VatRollInvoiceOCRWithContext(ctx context.Context, request *VatRollInvoiceOCRRequest) (response *VatRollInvoiceOCRResponse, err error) {
    if request == nil {
        request = NewVatRollInvoiceOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("VatRollInvoiceOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewVatRollInvoiceOCRResponse()
    err = c.Send(request, response)
    return
}

func NewVehicleLicenseOCRRequest() (request *VehicleLicenseOCRRequest) {
    request = &VehicleLicenseOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "VehicleLicenseOCR")
    
    
    return
}

func NewVehicleLicenseOCRResponse() (response *VehicleLicenseOCRResponse) {
    response = &VehicleLicenseOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// VehicleLicenseOCR
// 本接口支持行驶证主页和副页所有字段的自动定位与识别。
//
// 
//
// 行驶证主页：车牌号码、车辆类型、所有人、住址、使用性质、品牌型号、识别代码、发动机号、注册日期、发证日期、发证单位。
//
// 
//
// 行驶证副页：号牌号码、档案编号、核定载人数、总质量、整备质量、核定载质量、外廓尺寸、准牵引总质量、备注、检验记录。
//
// 
//
// 另外，本接口还支持复印件、翻拍和PS告警功能。
//
// 
//
// 默认接口请求频率限制：10次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEBLUR = "FailedOperation.ImageBlur"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) VehicleLicenseOCR(request *VehicleLicenseOCRRequest) (response *VehicleLicenseOCRResponse, err error) {
    return c.VehicleLicenseOCRWithContext(context.Background(), request)
}

// VehicleLicenseOCR
// 本接口支持行驶证主页和副页所有字段的自动定位与识别。
//
// 
//
// 行驶证主页：车牌号码、车辆类型、所有人、住址、使用性质、品牌型号、识别代码、发动机号、注册日期、发证日期、发证单位。
//
// 
//
// 行驶证副页：号牌号码、档案编号、核定载人数、总质量、整备质量、核定载质量、外廓尺寸、准牵引总质量、备注、检验记录。
//
// 
//
// 另外，本接口还支持复印件、翻拍和PS告警功能。
//
// 
//
// 默认接口请求频率限制：10次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEBLUR = "FailedOperation.ImageBlur"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) VehicleLicenseOCRWithContext(ctx context.Context, request *VehicleLicenseOCRRequest) (response *VehicleLicenseOCRResponse, err error) {
    if request == nil {
        request = NewVehicleLicenseOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("VehicleLicenseOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewVehicleLicenseOCRResponse()
    err = c.Send(request, response)
    return
}

func NewVehicleRegCertOCRRequest() (request *VehicleRegCertOCRRequest) {
    request = &VehicleRegCertOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "VehicleRegCertOCR")
    
    
    return
}

func NewVehicleRegCertOCRResponse() (response *VehicleRegCertOCRResponse) {
    response = &VehicleRegCertOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// VehicleRegCertOCR
// 本接口支持国内机动车登记证书主要字段的结构化识别，包括机动车所有人、身份证明名称、号码、车辆型号、车辆识别代号、发动机号、制造厂名称等。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) VehicleRegCertOCR(request *VehicleRegCertOCRRequest) (response *VehicleRegCertOCRResponse, err error) {
    return c.VehicleRegCertOCRWithContext(context.Background(), request)
}

// VehicleRegCertOCR
// 本接口支持国内机动车登记证书主要字段的结构化识别，包括机动车所有人、身份证明名称、号码、车辆型号、车辆识别代号、发动机号、制造厂名称等。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) VehicleRegCertOCRWithContext(ctx context.Context, request *VehicleRegCertOCRRequest) (response *VehicleRegCertOCRResponse, err error) {
    if request == nil {
        request = NewVehicleRegCertOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("VehicleRegCertOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewVehicleRegCertOCRResponse()
    err = c.Send(request, response)
    return
}

func NewVerifyBasicBizLicenseRequest() (request *VerifyBasicBizLicenseRequest) {
    request = &VerifyBasicBizLicenseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "VerifyBasicBizLicense")
    
    
    return
}

func NewVerifyBasicBizLicenseResponse() (response *VerifyBasicBizLicenseResponse) {
    response = &VerifyBasicBizLicenseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// VerifyBasicBizLicense
// 本接口支持营业执照信息的识别与准确性核验。
//
// 
//
// 您可以通过输入营业执照注册号或营业执照图片（若两者都输入则只用注册号做查询）进行核验，接口返回查询到的工商照面信息，并比对要校验的字段与查询结果的一致性。查询到工商信息包括：统一社会信用代码、经营期限、法人姓名、经营状态、经营业务范围、注册资本等。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATASOURCEQUERYFAILED = "FailedOperation.DataSourceQueryFailed"
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_NOBIZLICENSE = "FailedOperation.NoBizLicense"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_QUERYNORECORD = "FailedOperation.QueryNoRecord"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) VerifyBasicBizLicense(request *VerifyBasicBizLicenseRequest) (response *VerifyBasicBizLicenseResponse, err error) {
    return c.VerifyBasicBizLicenseWithContext(context.Background(), request)
}

// VerifyBasicBizLicense
// 本接口支持营业执照信息的识别与准确性核验。
//
// 
//
// 您可以通过输入营业执照注册号或营业执照图片（若两者都输入则只用注册号做查询）进行核验，接口返回查询到的工商照面信息，并比对要校验的字段与查询结果的一致性。查询到工商信息包括：统一社会信用代码、经营期限、法人姓名、经营状态、经营业务范围、注册资本等。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATASOURCEQUERYFAILED = "FailedOperation.DataSourceQueryFailed"
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_NOBIZLICENSE = "FailedOperation.NoBizLicense"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_QUERYNORECORD = "FailedOperation.QueryNoRecord"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) VerifyBasicBizLicenseWithContext(ctx context.Context, request *VerifyBasicBizLicenseRequest) (response *VerifyBasicBizLicenseResponse, err error) {
    if request == nil {
        request = NewVerifyBasicBizLicenseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("VerifyBasicBizLicense require credential")
    }

    request.SetContext(ctx)
    
    response = NewVerifyBasicBizLicenseResponse()
    err = c.Send(request, response)
    return
}

func NewVerifyBizLicenseRequest() (request *VerifyBizLicenseRequest) {
    request = &VerifyBizLicenseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "VerifyBizLicense")
    
    
    return
}

func NewVerifyBizLicenseResponse() (response *VerifyBizLicenseResponse) {
    response = &VerifyBizLicenseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// VerifyBizLicense
// 本接口支持营业执照信息的识别与准确性核验，返回的真实工商照面信息比营业执照识别及核验（基础版）接口更详细。
//
// 
//
// 您可以输入营业执照注册号或营业执照图片（若两者都输入则只用注册号做查询），接口返回查询到的工商照面信息，并比对要校验的字段与查询结果的一致性。
//
// 
//
// 查询到工商信息包括：统一社会信用代码、组织机构代码、经营期限、法人姓名、经营状态、经营业务范围及方式、注册资金、注册币种、登记机关、开业日期、企业（机构）类型、注销日期、吊销日期、许可经营项目、一般经营项目、核准时间、省、地级市、区/县、住所所在行政区划代码、行业门类代码、行业门类名称、国民经济行业代码、国民经济行业名称、经营（业务）范围等。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATASOURCEQUERYFAILED = "FailedOperation.DataSourceQueryFailed"
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_NOBIZLICENSE = "FailedOperation.NoBizLicense"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_QUERYNORECORD = "FailedOperation.QueryNoRecord"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) VerifyBizLicense(request *VerifyBizLicenseRequest) (response *VerifyBizLicenseResponse, err error) {
    return c.VerifyBizLicenseWithContext(context.Background(), request)
}

// VerifyBizLicense
// 本接口支持营业执照信息的识别与准确性核验，返回的真实工商照面信息比营业执照识别及核验（基础版）接口更详细。
//
// 
//
// 您可以输入营业执照注册号或营业执照图片（若两者都输入则只用注册号做查询），接口返回查询到的工商照面信息，并比对要校验的字段与查询结果的一致性。
//
// 
//
// 查询到工商信息包括：统一社会信用代码、组织机构代码、经营期限、法人姓名、经营状态、经营业务范围及方式、注册资金、注册币种、登记机关、开业日期、企业（机构）类型、注销日期、吊销日期、许可经营项目、一般经营项目、核准时间、省、地级市、区/县、住所所在行政区划代码、行业门类代码、行业门类名称、国民经济行业代码、国民经济行业名称、经营（业务）范围等。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATASOURCEQUERYFAILED = "FailedOperation.DataSourceQueryFailed"
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_NOBIZLICENSE = "FailedOperation.NoBizLicense"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_QUERYNORECORD = "FailedOperation.QueryNoRecord"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) VerifyBizLicenseWithContext(ctx context.Context, request *VerifyBizLicenseRequest) (response *VerifyBizLicenseResponse, err error) {
    if request == nil {
        request = NewVerifyBizLicenseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("VerifyBizLicense require credential")
    }

    request.SetContext(ctx)
    
    response = NewVerifyBizLicenseResponse()
    err = c.Send(request, response)
    return
}

func NewVerifyEnterpriseFourFactorsRequest() (request *VerifyEnterpriseFourFactorsRequest) {
    request = &VerifyEnterpriseFourFactorsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "VerifyEnterpriseFourFactors")
    
    
    return
}

func NewVerifyEnterpriseFourFactorsResponse() (response *VerifyEnterpriseFourFactorsResponse) {
    response = &VerifyEnterpriseFourFactorsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// VerifyEnterpriseFourFactors
// 此接口基于企业四要素授权“姓名、证件号码、企业标识、企业全称”，验证企业信息是否一致。
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) VerifyEnterpriseFourFactors(request *VerifyEnterpriseFourFactorsRequest) (response *VerifyEnterpriseFourFactorsResponse, err error) {
    return c.VerifyEnterpriseFourFactorsWithContext(context.Background(), request)
}

// VerifyEnterpriseFourFactors
// 此接口基于企业四要素授权“姓名、证件号码、企业标识、企业全称”，验证企业信息是否一致。
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) VerifyEnterpriseFourFactorsWithContext(ctx context.Context, request *VerifyEnterpriseFourFactorsRequest) (response *VerifyEnterpriseFourFactorsResponse, err error) {
    if request == nil {
        request = NewVerifyEnterpriseFourFactorsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("VerifyEnterpriseFourFactors require credential")
    }

    request.SetContext(ctx)
    
    response = NewVerifyEnterpriseFourFactorsResponse()
    err = c.Send(request, response)
    return
}

func NewVerifyOfdVatInvoiceOCRRequest() (request *VerifyOfdVatInvoiceOCRRequest) {
    request = &VerifyOfdVatInvoiceOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "VerifyOfdVatInvoiceOCR")
    
    
    return
}

func NewVerifyOfdVatInvoiceOCRResponse() (response *VerifyOfdVatInvoiceOCRResponse) {
    response = &VerifyOfdVatInvoiceOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// VerifyOfdVatInvoiceOCR
// 本接口支持OFD格式的增值税电子普通发票和增值税电子专用发票的识别，返回发票代码、发票号码、开票日期、验证码、机器编号、密码区，购买方和销售方信息，包括名称、纳税人识别号、地址电话、开户行及账号，以及价税合计、开票人、收款人、复核人、税额、不含税金额等字段信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) VerifyOfdVatInvoiceOCR(request *VerifyOfdVatInvoiceOCRRequest) (response *VerifyOfdVatInvoiceOCRResponse, err error) {
    return c.VerifyOfdVatInvoiceOCRWithContext(context.Background(), request)
}

// VerifyOfdVatInvoiceOCR
// 本接口支持OFD格式的增值税电子普通发票和增值税电子专用发票的识别，返回发票代码、发票号码、开票日期、验证码、机器编号、密码区，购买方和销售方信息，包括名称、纳税人识别号、地址电话、开户行及账号，以及价税合计、开票人、收款人、复核人、税额、不含税金额等字段信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) VerifyOfdVatInvoiceOCRWithContext(ctx context.Context, request *VerifyOfdVatInvoiceOCRRequest) (response *VerifyOfdVatInvoiceOCRResponse, err error) {
    if request == nil {
        request = NewVerifyOfdVatInvoiceOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("VerifyOfdVatInvoiceOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewVerifyOfdVatInvoiceOCRResponse()
    err = c.Send(request, response)
    return
}

func NewVinOCRRequest() (request *VinOCRRequest) {
    request = &VinOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "VinOCR")
    
    
    return
}

func NewVinOCRResponse() (response *VinOCRResponse) {
    response = &VinOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// VinOCR
// 本接口支持图片内车辆识别代号（VIN）的检测和识别。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) VinOCR(request *VinOCRRequest) (response *VinOCRResponse, err error) {
    return c.VinOCRWithContext(context.Background(), request)
}

// VinOCR
// 本接口支持图片内车辆识别代号（VIN）的检测和识别。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) VinOCRWithContext(ctx context.Context, request *VinOCRRequest) (response *VinOCRResponse, err error) {
    if request == nil {
        request = NewVinOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("VinOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewVinOCRResponse()
    err = c.Send(request, response)
    return
}

func NewWaybillOCRRequest() (request *WaybillOCRRequest) {
    request = &WaybillOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "WaybillOCR")
    
    
    return
}

func NewWaybillOCRResponse() (response *WaybillOCRResponse) {
    response = &WaybillOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// WaybillOCR
// 本接口支持市面上主流版式电子运单的识别，包括收件人和寄件人的姓名、电话、地址以及运单号等字段，精度均处于业界领先水平，识别准确率达到99%以上。
//
// 
//
// 默认接口请求频率限制：10次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_LANGUAGENOTSUPPORT = "FailedOperation.LanguageNotSupport"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) WaybillOCR(request *WaybillOCRRequest) (response *WaybillOCRResponse, err error) {
    return c.WaybillOCRWithContext(context.Background(), request)
}

// WaybillOCR
// 本接口支持市面上主流版式电子运单的识别，包括收件人和寄件人的姓名、电话、地址以及运单号等字段，精度均处于业界领先水平，识别准确率达到99%以上。
//
// 
//
// 默认接口请求频率限制：10次/秒。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_LANGUAGENOTSUPPORT = "FailedOperation.LanguageNotSupport"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) WaybillOCRWithContext(ctx context.Context, request *WaybillOCRRequest) (response *WaybillOCRResponse, err error) {
    if request == nil {
        request = NewWaybillOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("WaybillOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewWaybillOCRResponse()
    err = c.Send(request, response)
    return
}
