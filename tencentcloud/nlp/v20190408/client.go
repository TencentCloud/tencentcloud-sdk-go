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

package v20190408

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-04-08"

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


func NewAnalyzeSentimentRequest() (request *AnalyzeSentimentRequest) {
    request = &AnalyzeSentimentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("nlp", APIVersion, "AnalyzeSentiment")
    
    
    return
}

func NewAnalyzeSentimentResponse() (response *AnalyzeSentimentResponse) {
    response = &AnalyzeSentimentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AnalyzeSentiment
// 情感分析接口能够对带有情感色彩的主观性文本进行分析、处理、归纳和推理，识别出用户的情感倾向，是积极、中性还是消极，并且提供各自概率。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_NORESULTS = "FailedOperation.NoResults"
//  FAILEDOPERATION_NOTFOUNDDATA = "FailedOperation.NotFoundData"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RESOURCEBUSY = "FailedOperation.ResourceBusy"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVICECALLERROR = "InternalError.ServiceCallError"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_TEXTCLASSIFYERROR = "InternalError.TextClassifyError"
//  INTERNALERROR_TEXTPARSINGERROR = "InternalError.TextParsingError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETER = "InvalidParameterValue.InvalidParameter"
//  INVALIDPARAMETERVALUE_SENSITIVETEXT = "InvalidParameterValue.SensitiveText"
//  INVALIDPARAMETERVALUE_TEXT = "InvalidParameterValue.Text"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  LIMITEXCEEDED_RESOURCEREACHEDLIMIT = "LimitExceeded.ResourceReachedLimit"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_ISOPENING = "ResourceUnavailable.IsOpening"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  UNSUPPORTEDOPERATION_PKGEXHAUSTED = "UnsupportedOperation.PkgExhausted"
func (c *Client) AnalyzeSentiment(request *AnalyzeSentimentRequest) (response *AnalyzeSentimentResponse, err error) {
    return c.AnalyzeSentimentWithContext(context.Background(), request)
}

// AnalyzeSentiment
// 情感分析接口能够对带有情感色彩的主观性文本进行分析、处理、归纳和推理，识别出用户的情感倾向，是积极、中性还是消极，并且提供各自概率。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_NORESULTS = "FailedOperation.NoResults"
//  FAILEDOPERATION_NOTFOUNDDATA = "FailedOperation.NotFoundData"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RESOURCEBUSY = "FailedOperation.ResourceBusy"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVICECALLERROR = "InternalError.ServiceCallError"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_TEXTCLASSIFYERROR = "InternalError.TextClassifyError"
//  INTERNALERROR_TEXTPARSINGERROR = "InternalError.TextParsingError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETER = "InvalidParameterValue.InvalidParameter"
//  INVALIDPARAMETERVALUE_SENSITIVETEXT = "InvalidParameterValue.SensitiveText"
//  INVALIDPARAMETERVALUE_TEXT = "InvalidParameterValue.Text"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  LIMITEXCEEDED_RESOURCEREACHEDLIMIT = "LimitExceeded.ResourceReachedLimit"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_ISOPENING = "ResourceUnavailable.IsOpening"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  UNSUPPORTEDOPERATION_PKGEXHAUSTED = "UnsupportedOperation.PkgExhausted"
func (c *Client) AnalyzeSentimentWithContext(ctx context.Context, request *AnalyzeSentimentRequest) (response *AnalyzeSentimentResponse, err error) {
    if request == nil {
        request = NewAnalyzeSentimentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "nlp", APIVersion, "AnalyzeSentiment")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AnalyzeSentiment require credential")
    }

    request.SetContext(ctx)
    
    response = NewAnalyzeSentimentResponse()
    err = c.Send(request, response)
    return
}

func NewParseWordsRequest() (request *ParseWordsRequest) {
    request = &ParseWordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("nlp", APIVersion, "ParseWords")
    
    
    return
}

func NewParseWordsResponse() (response *ParseWordsResponse) {
    response = &ParseWordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ParseWords
// 通过精准地对文本进行分词、词性标注、命名实体识别等功能，助您更好地理解文本内容，挖掘出潜在的价值信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_NORESULTS = "FailedOperation.NoResults"
//  FAILEDOPERATION_NOTFOUNDDATA = "FailedOperation.NotFoundData"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RESOURCEBUSY = "FailedOperation.ResourceBusy"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVICECALLERROR = "InternalError.ServiceCallError"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETER = "InvalidParameterValue.InvalidParameter"
//  INVALIDPARAMETERVALUE_SENSITIVETEXT = "InvalidParameterValue.SensitiveText"
//  INVALIDPARAMETERVALUE_TEXTNUMTOOMUCH = "InvalidParameterValue.TextNumTooMuch"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  LIMITEXCEEDED_RESOURCEREACHEDLIMIT = "LimitExceeded.ResourceReachedLimit"
//  RESOURCEINSUFFICIENT_QUOTARUNOUT = "ResourceInsufficient.QuotaRunOut"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_ISOPENING = "ResourceUnavailable.IsOpening"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  UNSUPPORTEDOPERATION_PKGEXHAUSTED = "UnsupportedOperation.PkgExhausted"
func (c *Client) ParseWords(request *ParseWordsRequest) (response *ParseWordsResponse, err error) {
    return c.ParseWordsWithContext(context.Background(), request)
}

// ParseWords
// 通过精准地对文本进行分词、词性标注、命名实体识别等功能，助您更好地理解文本内容，挖掘出潜在的价值信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_NORESULTS = "FailedOperation.NoResults"
//  FAILEDOPERATION_NOTFOUNDDATA = "FailedOperation.NotFoundData"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RESOURCEBUSY = "FailedOperation.ResourceBusy"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVICECALLERROR = "InternalError.ServiceCallError"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETER = "InvalidParameterValue.InvalidParameter"
//  INVALIDPARAMETERVALUE_SENSITIVETEXT = "InvalidParameterValue.SensitiveText"
//  INVALIDPARAMETERVALUE_TEXTNUMTOOMUCH = "InvalidParameterValue.TextNumTooMuch"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  LIMITEXCEEDED_RESOURCEREACHEDLIMIT = "LimitExceeded.ResourceReachedLimit"
//  RESOURCEINSUFFICIENT_QUOTARUNOUT = "ResourceInsufficient.QuotaRunOut"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_ISOPENING = "ResourceUnavailable.IsOpening"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  UNSUPPORTEDOPERATION_PKGEXHAUSTED = "UnsupportedOperation.PkgExhausted"
func (c *Client) ParseWordsWithContext(ctx context.Context, request *ParseWordsRequest) (response *ParseWordsResponse, err error) {
    if request == nil {
        request = NewParseWordsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "nlp", APIVersion, "ParseWords")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ParseWords require credential")
    }

    request.SetContext(ctx)
    
    response = NewParseWordsResponse()
    err = c.Send(request, response)
    return
}

func NewSentenceCorrectionRequest() (request *SentenceCorrectionRequest) {
    request = &SentenceCorrectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("nlp", APIVersion, "SentenceCorrection")
    
    
    return
}

func NewSentenceCorrectionResponse() (response *SentenceCorrectionResponse) {
    response = &SentenceCorrectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SentenceCorrection
// 智能识别并纠正句子中的语法、拼写、用词等错误，确保文本的准确性和可读性。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_NORESULTS = "FailedOperation.NoResults"
//  FAILEDOPERATION_NOTFOUNDDATA = "FailedOperation.NotFoundData"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RESOURCEBUSY = "FailedOperation.ResourceBusy"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVICECALLERROR = "InternalError.ServiceCallError"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INVALIDPARAMETER_TEXTTOOLONGCODE = "InvalidParameter.TextTooLongCode"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETER = "InvalidParameterValue.InvalidParameter"
//  INVALIDPARAMETERVALUE_SENSITIVETEXT = "InvalidParameterValue.SensitiveText"
//  INVALIDPARAMETERVALUE_TEXTNUMTOOMUCH = "InvalidParameterValue.TextNumTooMuch"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  LIMITEXCEEDED_RESOURCEREACHEDLIMIT = "LimitExceeded.ResourceReachedLimit"
//  REQUESTLIMITEXCEEDED_UINLIMITEXCEEDED = "RequestLimitExceeded.UinLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_ISOPENING = "ResourceUnavailable.IsOpening"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  UNSUPPORTEDOPERATION_PKGEXHAUSTED = "UnsupportedOperation.PkgExhausted"
func (c *Client) SentenceCorrection(request *SentenceCorrectionRequest) (response *SentenceCorrectionResponse, err error) {
    return c.SentenceCorrectionWithContext(context.Background(), request)
}

// SentenceCorrection
// 智能识别并纠正句子中的语法、拼写、用词等错误，确保文本的准确性和可读性。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_NORESULTS = "FailedOperation.NoResults"
//  FAILEDOPERATION_NOTFOUNDDATA = "FailedOperation.NotFoundData"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RESOURCEBUSY = "FailedOperation.ResourceBusy"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVICECALLERROR = "InternalError.ServiceCallError"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INVALIDPARAMETER_TEXTTOOLONGCODE = "InvalidParameter.TextTooLongCode"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETER = "InvalidParameterValue.InvalidParameter"
//  INVALIDPARAMETERVALUE_SENSITIVETEXT = "InvalidParameterValue.SensitiveText"
//  INVALIDPARAMETERVALUE_TEXTNUMTOOMUCH = "InvalidParameterValue.TextNumTooMuch"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  LIMITEXCEEDED_RESOURCEREACHEDLIMIT = "LimitExceeded.ResourceReachedLimit"
//  REQUESTLIMITEXCEEDED_UINLIMITEXCEEDED = "RequestLimitExceeded.UinLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_ISOPENING = "ResourceUnavailable.IsOpening"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  UNSUPPORTEDOPERATION_PKGEXHAUSTED = "UnsupportedOperation.PkgExhausted"
func (c *Client) SentenceCorrectionWithContext(ctx context.Context, request *SentenceCorrectionRequest) (response *SentenceCorrectionResponse, err error) {
    if request == nil {
        request = NewSentenceCorrectionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "nlp", APIVersion, "SentenceCorrection")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SentenceCorrection require credential")
    }

    request.SetContext(ctx)
    
    response = NewSentenceCorrectionResponse()
    err = c.Send(request, response)
    return
}
