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

func NewClassifyContentRequest() (request *ClassifyContentRequest) {
    request = &ClassifyContentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("nlp", APIVersion, "ClassifyContent")
    
    
    return
}

func NewClassifyContentResponse() (response *ClassifyContentResponse) {
    response = &ClassifyContentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ClassifyContent
// NLP技术的句子相似度、相似词召回、文本分类、对联生成、诗词生成、词相似度、文本润色、句子生成和文本补全API接口将于2025年10月31日下线，届时将无法正常调用。为了避免对您的业务造成影响，请您尽快做好相关业务调整。如果您有NLP技术的产品需求，推荐您调用腾讯混元大模型（https://cloud.tencent.com/product/tclm）。
//
// 
//
// 文本分类接口能够对用户输入的文章进行自动分类，将其映射到具体的类目上，用户只需要提供待分类的文本，而无需关注具体实现。该功能定义了一套较为完备的[三级分类体系](https://cloud.tencent.com/document/product/271/94286)，积累了数百万的语料，经过多轮迭代优化打造了较先进的深度学习模型，以保证效果不断提升。
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
func (c *Client) ClassifyContent(request *ClassifyContentRequest) (response *ClassifyContentResponse, err error) {
    return c.ClassifyContentWithContext(context.Background(), request)
}

// ClassifyContent
// NLP技术的句子相似度、相似词召回、文本分类、对联生成、诗词生成、词相似度、文本润色、句子生成和文本补全API接口将于2025年10月31日下线，届时将无法正常调用。为了避免对您的业务造成影响，请您尽快做好相关业务调整。如果您有NLP技术的产品需求，推荐您调用腾讯混元大模型（https://cloud.tencent.com/product/tclm）。
//
// 
//
// 文本分类接口能够对用户输入的文章进行自动分类，将其映射到具体的类目上，用户只需要提供待分类的文本，而无需关注具体实现。该功能定义了一套较为完备的[三级分类体系](https://cloud.tencent.com/document/product/271/94286)，积累了数百万的语料，经过多轮迭代优化打造了较先进的深度学习模型，以保证效果不断提升。
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
func (c *Client) ClassifyContentWithContext(ctx context.Context, request *ClassifyContentRequest) (response *ClassifyContentResponse, err error) {
    if request == nil {
        request = NewClassifyContentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "nlp", APIVersion, "ClassifyContent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ClassifyContent require credential")
    }

    request.SetContext(ctx)
    
    response = NewClassifyContentResponse()
    err = c.Send(request, response)
    return
}

func NewComposeCoupletRequest() (request *ComposeCoupletRequest) {
    request = &ComposeCoupletRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("nlp", APIVersion, "ComposeCouplet")
    
    
    return
}

func NewComposeCoupletResponse() (response *ComposeCoupletResponse) {
    response = &ComposeCoupletResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ComposeCouplet
// NLP技术的句子相似度、相似词召回、文本分类、对联生成、诗词生成、词相似度、文本润色、句子生成和文本补全API接口将于2025年10月31日下线，届时将无法正常调用。为了避免对您的业务造成影响，请您尽快做好相关业务调整。如果您有NLP技术的产品需求，推荐您调用腾讯混元大模型（https://cloud.tencent.com/product/tclm）。
//
// 
//
// 对联生成接口根据用户输入的命题关键词，智能生成一副完整的春联，包括上联、下联和横批。该接口利用先进的自然语言处理技术，确保生成的春联既符合传统对仗、对韵、对义的要求，又具有新意和创意，为用户提供独特的春节祝福。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  FAILEDOPERATION_NOCOUPLETS = "FailedOperation.NoCouplets"
//  FAILEDOPERATION_NOTFOUNDDATA = "FailedOperation.NotFoundData"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVICECALLERROR = "InternalError.ServiceCallError"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETER = "InvalidParameterValue.InvalidParameter"
//  INVALIDPARAMETERVALUE_SENSITIVETEXT = "InvalidParameterValue.SensitiveText"
//  INVALIDPARAMETERVALUE_TARGETTYPE = "InvalidParameterValue.TargetType"
//  INVALIDPARAMETERVALUE_TEXT = "InvalidParameterValue.Text"
//  LIMITEXCEEDED_RESOURCEREACHEDLIMIT = "LimitExceeded.ResourceReachedLimit"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_ISOPENING = "ResourceUnavailable.IsOpening"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  UNSUPPORTEDOPERATION_PKGEXHAUSTED = "UnsupportedOperation.PkgExhausted"
func (c *Client) ComposeCouplet(request *ComposeCoupletRequest) (response *ComposeCoupletResponse, err error) {
    return c.ComposeCoupletWithContext(context.Background(), request)
}

// ComposeCouplet
// NLP技术的句子相似度、相似词召回、文本分类、对联生成、诗词生成、词相似度、文本润色、句子生成和文本补全API接口将于2025年10月31日下线，届时将无法正常调用。为了避免对您的业务造成影响，请您尽快做好相关业务调整。如果您有NLP技术的产品需求，推荐您调用腾讯混元大模型（https://cloud.tencent.com/product/tclm）。
//
// 
//
// 对联生成接口根据用户输入的命题关键词，智能生成一副完整的春联，包括上联、下联和横批。该接口利用先进的自然语言处理技术，确保生成的春联既符合传统对仗、对韵、对义的要求，又具有新意和创意，为用户提供独特的春节祝福。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  FAILEDOPERATION_NOCOUPLETS = "FailedOperation.NoCouplets"
//  FAILEDOPERATION_NOTFOUNDDATA = "FailedOperation.NotFoundData"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVICECALLERROR = "InternalError.ServiceCallError"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETER = "InvalidParameterValue.InvalidParameter"
//  INVALIDPARAMETERVALUE_SENSITIVETEXT = "InvalidParameterValue.SensitiveText"
//  INVALIDPARAMETERVALUE_TARGETTYPE = "InvalidParameterValue.TargetType"
//  INVALIDPARAMETERVALUE_TEXT = "InvalidParameterValue.Text"
//  LIMITEXCEEDED_RESOURCEREACHEDLIMIT = "LimitExceeded.ResourceReachedLimit"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_ISOPENING = "ResourceUnavailable.IsOpening"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  UNSUPPORTEDOPERATION_PKGEXHAUSTED = "UnsupportedOperation.PkgExhausted"
func (c *Client) ComposeCoupletWithContext(ctx context.Context, request *ComposeCoupletRequest) (response *ComposeCoupletResponse, err error) {
    if request == nil {
        request = NewComposeCoupletRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "nlp", APIVersion, "ComposeCouplet")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ComposeCouplet require credential")
    }

    request.SetContext(ctx)
    
    response = NewComposeCoupletResponse()
    err = c.Send(request, response)
    return
}

func NewEvaluateSentenceSimilarityRequest() (request *EvaluateSentenceSimilarityRequest) {
    request = &EvaluateSentenceSimilarityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("nlp", APIVersion, "EvaluateSentenceSimilarity")
    
    
    return
}

func NewEvaluateSentenceSimilarityResponse() (response *EvaluateSentenceSimilarityResponse) {
    response = &EvaluateSentenceSimilarityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EvaluateSentenceSimilarity
// NLP技术的句子相似度、相似词召回、文本分类、对联生成、诗词生成、词相似度、文本润色、句子生成和文本补全API接口将于2025年10月31日下线，届时将无法正常调用。为了避免对您的业务造成影响，请您尽快做好相关业务调整。如果您有NLP技术的产品需求，推荐您调用腾讯混元大模型（https://cloud.tencent.com/product/tclm）。
//
// 
//
// 通过计算句子间的语义相似性，帮助您快速找到文本中重复或相似的句子，用于文本聚类、相似问题检索等应用场景。
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
//  INTERNALERROR_TEXTPARSINGERROR = "InternalError.TextParsingError"
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
func (c *Client) EvaluateSentenceSimilarity(request *EvaluateSentenceSimilarityRequest) (response *EvaluateSentenceSimilarityResponse, err error) {
    return c.EvaluateSentenceSimilarityWithContext(context.Background(), request)
}

// EvaluateSentenceSimilarity
// NLP技术的句子相似度、相似词召回、文本分类、对联生成、诗词生成、词相似度、文本润色、句子生成和文本补全API接口将于2025年10月31日下线，届时将无法正常调用。为了避免对您的业务造成影响，请您尽快做好相关业务调整。如果您有NLP技术的产品需求，推荐您调用腾讯混元大模型（https://cloud.tencent.com/product/tclm）。
//
// 
//
// 通过计算句子间的语义相似性，帮助您快速找到文本中重复或相似的句子，用于文本聚类、相似问题检索等应用场景。
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
//  INTERNALERROR_TEXTPARSINGERROR = "InternalError.TextParsingError"
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
func (c *Client) EvaluateSentenceSimilarityWithContext(ctx context.Context, request *EvaluateSentenceSimilarityRequest) (response *EvaluateSentenceSimilarityResponse, err error) {
    if request == nil {
        request = NewEvaluateSentenceSimilarityRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "nlp", APIVersion, "EvaluateSentenceSimilarity")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EvaluateSentenceSimilarity require credential")
    }

    request.SetContext(ctx)
    
    response = NewEvaluateSentenceSimilarityResponse()
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
