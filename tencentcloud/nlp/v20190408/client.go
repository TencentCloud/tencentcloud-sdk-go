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


func NewAutoSummarizationRequest() (request *AutoSummarizationRequest) {
    request = &AutoSummarizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("nlp", APIVersion, "AutoSummarization")
    
    
    return
}

func NewAutoSummarizationResponse() (response *AutoSummarizationResponse) {
    response = &AutoSummarizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AutoSummarization
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 利用人工智能算法，自动抽取文本中的关键信息并生成指定长度的文本摘要。可用于新闻标题生成、科技文献摘要生成和商品评论摘要等。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_ILLEGALTEXTERROR = "FailedOperation.IllegalTextError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
//  INVALIDPARAMETERVALUE_TEXTENCODEERROR = "InvalidParameterValue.TextEncodeError"
//  INVALIDPARAMETERVALUE_TEXTFORMATERROR = "InvalidParameterValue.TextFormatError"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  INVALIDPARAMETERVALUE_VALUERANGEERROR = "InvalidParameterValue.ValueRangeError"
//  RESOURCEINSUFFICIENT_QUOTARUNOUT = "ResourceInsufficient.QuotaRunOut"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
//  UNAUTHORIZEDOPERATION_AUTHENTICATEFAILED = "UnauthorizedOperation.AuthenticateFailed"
func (c *Client) AutoSummarization(request *AutoSummarizationRequest) (response *AutoSummarizationResponse, err error) {
    return c.AutoSummarizationWithContext(context.Background(), request)
}

// AutoSummarization
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 利用人工智能算法，自动抽取文本中的关键信息并生成指定长度的文本摘要。可用于新闻标题生成、科技文献摘要生成和商品评论摘要等。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_ILLEGALTEXTERROR = "FailedOperation.IllegalTextError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
//  INVALIDPARAMETERVALUE_TEXTENCODEERROR = "InvalidParameterValue.TextEncodeError"
//  INVALIDPARAMETERVALUE_TEXTFORMATERROR = "InvalidParameterValue.TextFormatError"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  INVALIDPARAMETERVALUE_VALUERANGEERROR = "InvalidParameterValue.ValueRangeError"
//  RESOURCEINSUFFICIENT_QUOTARUNOUT = "ResourceInsufficient.QuotaRunOut"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
//  UNAUTHORIZEDOPERATION_AUTHENTICATEFAILED = "UnauthorizedOperation.AuthenticateFailed"
func (c *Client) AutoSummarizationWithContext(ctx context.Context, request *AutoSummarizationRequest) (response *AutoSummarizationResponse, err error) {
    if request == nil {
        request = NewAutoSummarizationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AutoSummarization require credential")
    }

    request.SetContext(ctx)
    
    response = NewAutoSummarizationResponse()
    err = c.Send(request, response)
    return
}

func NewChatBotRequest() (request *ChatBotRequest) {
    request = &ChatBotRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("nlp", APIVersion, "ChatBot")
    
    
    return
}

func NewChatBotResponse() (response *ChatBotResponse) {
    response = &ChatBotResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ChatBot
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 闲聊服务基于腾讯领先的NLP引擎能力、数据运算能力和千亿级互联网语料数据的支持，同时集成了广泛的知识问答能力，可实现上百种自定义属性配置，以及儿童语言风格及说话方式，从而让聊天变得更睿智、简单和有趣。
//
// 
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  RESOURCEINSUFFICIENT_QUOTARUNOUT = "ResourceInsufficient.QuotaRunOut"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
//  UNAUTHORIZEDOPERATION_AUTHENTICATEFAILED = "UnauthorizedOperation.AuthenticateFailed"
func (c *Client) ChatBot(request *ChatBotRequest) (response *ChatBotResponse, err error) {
    return c.ChatBotWithContext(context.Background(), request)
}

// ChatBot
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 闲聊服务基于腾讯领先的NLP引擎能力、数据运算能力和千亿级互联网语料数据的支持，同时集成了广泛的知识问答能力，可实现上百种自定义属性配置，以及儿童语言风格及说话方式，从而让聊天变得更睿智、简单和有趣。
//
// 
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  RESOURCEINSUFFICIENT_QUOTARUNOUT = "ResourceInsufficient.QuotaRunOut"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
//  UNAUTHORIZEDOPERATION_AUTHENTICATEFAILED = "UnauthorizedOperation.AuthenticateFailed"
func (c *Client) ChatBotWithContext(ctx context.Context, request *ChatBotRequest) (response *ChatBotResponse, err error) {
    if request == nil {
        request = NewChatBotRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChatBot require credential")
    }

    request.SetContext(ctx)
    
    response = NewChatBotResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDictRequest() (request *CreateDictRequest) {
    request = &CreateDictRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("nlp", APIVersion, "CreateDict")
    
    
    return
}

func NewCreateDictResponse() (response *CreateDictResponse) {
    response = &CreateDictResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDict
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 根据指定的名称、描述创建自定义词库。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ILLEGALTEXTERROR = "FailedOperation.IllegalTextError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_RESOURCEREQUESTERROR = "InternalError.ResourceRequestError"
//  INTERNALERROR_SERVICECALLERROR = "InternalError.ServiceCallError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
//  INVALIDPARAMETERVALUE_TEXTFORMATERROR = "InvalidParameterValue.TextFormatError"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  INVALIDPARAMETERVALUE_VALUERANGEERROR = "InvalidParameterValue.ValueRangeError"
//  LIMITEXCEEDED_RESOURCEREACHEDLIMIT = "LimitExceeded.ResourceReachedLimit"
//  RESOURCEINUSE_NAMEEXISTS = "ResourceInUse.NameExists"
//  RESOURCEINUSE_RESOURCEOPERATING = "ResourceInUse.ResourceOperating"
//  RESOURCEINSUFFICIENT_QUOTARUNOUT = "ResourceInsufficient.QuotaRunOut"
//  RESOURCENOTFOUND_DATANOTFOUND = "ResourceNotFound.DataNotFound"
//  RESOURCENOTFOUND_FILENOTFOUND = "ResourceNotFound.FileNotFound"
//  RESOURCEUNAVAILABLE_FILEUNAVAILABLE = "ResourceUnavailable.FileUnavailable"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
//  UNAUTHORIZEDOPERATION_AUTHENTICATEFAILED = "UnauthorizedOperation.AuthenticateFailed"
func (c *Client) CreateDict(request *CreateDictRequest) (response *CreateDictResponse, err error) {
    return c.CreateDictWithContext(context.Background(), request)
}

// CreateDict
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 根据指定的名称、描述创建自定义词库。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ILLEGALTEXTERROR = "FailedOperation.IllegalTextError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_RESOURCEREQUESTERROR = "InternalError.ResourceRequestError"
//  INTERNALERROR_SERVICECALLERROR = "InternalError.ServiceCallError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
//  INVALIDPARAMETERVALUE_TEXTFORMATERROR = "InvalidParameterValue.TextFormatError"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  INVALIDPARAMETERVALUE_VALUERANGEERROR = "InvalidParameterValue.ValueRangeError"
//  LIMITEXCEEDED_RESOURCEREACHEDLIMIT = "LimitExceeded.ResourceReachedLimit"
//  RESOURCEINUSE_NAMEEXISTS = "ResourceInUse.NameExists"
//  RESOURCEINUSE_RESOURCEOPERATING = "ResourceInUse.ResourceOperating"
//  RESOURCEINSUFFICIENT_QUOTARUNOUT = "ResourceInsufficient.QuotaRunOut"
//  RESOURCENOTFOUND_DATANOTFOUND = "ResourceNotFound.DataNotFound"
//  RESOURCENOTFOUND_FILENOTFOUND = "ResourceNotFound.FileNotFound"
//  RESOURCEUNAVAILABLE_FILEUNAVAILABLE = "ResourceUnavailable.FileUnavailable"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
//  UNAUTHORIZEDOPERATION_AUTHENTICATEFAILED = "UnauthorizedOperation.AuthenticateFailed"
func (c *Client) CreateDictWithContext(ctx context.Context, request *CreateDictRequest) (response *CreateDictResponse, err error) {
    if request == nil {
        request = NewCreateDictRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDict require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDictResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWordItemsRequest() (request *CreateWordItemsRequest) {
    request = &CreateWordItemsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("nlp", APIVersion, "CreateWordItems")
    
    
    return
}

func NewCreateWordItemsResponse() (response *CreateWordItemsResponse) {
    response = &CreateWordItemsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateWordItems
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 向指定的词库中添加词条。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
func (c *Client) CreateWordItems(request *CreateWordItemsRequest) (response *CreateWordItemsResponse, err error) {
    return c.CreateWordItemsWithContext(context.Background(), request)
}

// CreateWordItems
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 向指定的词库中添加词条。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
func (c *Client) CreateWordItemsWithContext(ctx context.Context, request *CreateWordItemsRequest) (response *CreateWordItemsResponse, err error) {
    if request == nil {
        request = NewCreateWordItemsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWordItems require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWordItemsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDictRequest() (request *DeleteDictRequest) {
    request = &DeleteDictRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("nlp", APIVersion, "DeleteDict")
    
    
    return
}

func NewDeleteDictResponse() (response *DeleteDictResponse) {
    response = &DeleteDictResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDict
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 删除自定义词库，会附带相应删除词库包含的所有词条。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
func (c *Client) DeleteDict(request *DeleteDictRequest) (response *DeleteDictResponse, err error) {
    return c.DeleteDictWithContext(context.Background(), request)
}

// DeleteDict
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 删除自定义词库，会附带相应删除词库包含的所有词条。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
func (c *Client) DeleteDictWithContext(ctx context.Context, request *DeleteDictRequest) (response *DeleteDictResponse, err error) {
    if request == nil {
        request = NewDeleteDictRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDict require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDictResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteWordItemsRequest() (request *DeleteWordItemsRequest) {
    request = &DeleteWordItemsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("nlp", APIVersion, "DeleteWordItems")
    
    
    return
}

func NewDeleteWordItemsResponse() (response *DeleteWordItemsResponse) {
    response = &DeleteWordItemsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteWordItems
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 用于删除自定义词库中的词条。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
func (c *Client) DeleteWordItems(request *DeleteWordItemsRequest) (response *DeleteWordItemsResponse, err error) {
    return c.DeleteWordItemsWithContext(context.Background(), request)
}

// DeleteWordItems
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 用于删除自定义词库中的词条。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
func (c *Client) DeleteWordItemsWithContext(ctx context.Context, request *DeleteWordItemsRequest) (response *DeleteWordItemsResponse, err error) {
    if request == nil {
        request = NewDeleteWordItemsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWordItems require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteWordItemsResponse()
    err = c.Send(request, response)
    return
}

func NewDependencyParsingRequest() (request *DependencyParsingRequest) {
    request = &DependencyParsingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("nlp", APIVersion, "DependencyParsing")
    
    
    return
}

func NewDependencyParsingResponse() (response *DependencyParsingResponse) {
    response = &DependencyParsingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DependencyParsing
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 句法依存分析接口能够分析出句子中词与词之间的相互依存关系，并揭示其句法结构，包括主谓关系、动宾关系、核心关系等等，可用于提取句子主干、提取句子核心词等，在机器翻译、自动问答、知识抽取等领域都有很好的应用。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
//  INVALIDPARAMETERVALUE_TEXTENCODEERROR = "InvalidParameterValue.TextEncodeError"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  RESOURCEINSUFFICIENT_QUOTARUNOUT = "ResourceInsufficient.QuotaRunOut"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
//  UNAUTHORIZEDOPERATION_AUTHENTICATEFAILED = "UnauthorizedOperation.AuthenticateFailed"
func (c *Client) DependencyParsing(request *DependencyParsingRequest) (response *DependencyParsingResponse, err error) {
    return c.DependencyParsingWithContext(context.Background(), request)
}

// DependencyParsing
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 句法依存分析接口能够分析出句子中词与词之间的相互依存关系，并揭示其句法结构，包括主谓关系、动宾关系、核心关系等等，可用于提取句子主干、提取句子核心词等，在机器翻译、自动问答、知识抽取等领域都有很好的应用。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
//  INVALIDPARAMETERVALUE_TEXTENCODEERROR = "InvalidParameterValue.TextEncodeError"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  RESOURCEINSUFFICIENT_QUOTARUNOUT = "ResourceInsufficient.QuotaRunOut"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
//  UNAUTHORIZEDOPERATION_AUTHENTICATEFAILED = "UnauthorizedOperation.AuthenticateFailed"
func (c *Client) DependencyParsingWithContext(ctx context.Context, request *DependencyParsingRequest) (response *DependencyParsingResponse, err error) {
    if request == nil {
        request = NewDependencyParsingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DependencyParsing require credential")
    }

    request.SetContext(ctx)
    
    response = NewDependencyParsingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDictRequest() (request *DescribeDictRequest) {
    request = &DescribeDictRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("nlp", APIVersion, "DescribeDict")
    
    
    return
}

func NewDescribeDictResponse() (response *DescribeDictResponse) {
    response = &DescribeDictResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDict
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 根据id或名称查询自定义词库信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
//  UNAUTHORIZEDOPERATION_AUTHENTICATEFAILED = "UnauthorizedOperation.AuthenticateFailed"
func (c *Client) DescribeDict(request *DescribeDictRequest) (response *DescribeDictResponse, err error) {
    return c.DescribeDictWithContext(context.Background(), request)
}

// DescribeDict
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 根据id或名称查询自定义词库信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
//  UNAUTHORIZEDOPERATION_AUTHENTICATEFAILED = "UnauthorizedOperation.AuthenticateFailed"
func (c *Client) DescribeDictWithContext(ctx context.Context, request *DescribeDictRequest) (response *DescribeDictResponse, err error) {
    if request == nil {
        request = NewDescribeDictRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDict require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDictResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDictsRequest() (request *DescribeDictsRequest) {
    request = &DescribeDictsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("nlp", APIVersion, "DescribeDicts")
    
    
    return
}

func NewDescribeDictsResponse() (response *DescribeDictsResponse) {
    response = &DescribeDictsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDicts
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 返回属于当前用户的所有自定义词库列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_VALUERANGEERROR = "InvalidParameterValue.ValueRangeError"
//  RESOURCEINSUFFICIENT_QUOTARUNOUT = "ResourceInsufficient.QuotaRunOut"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
//  UNAUTHORIZEDOPERATION_AUTHENTICATEFAILED = "UnauthorizedOperation.AuthenticateFailed"
func (c *Client) DescribeDicts(request *DescribeDictsRequest) (response *DescribeDictsResponse, err error) {
    return c.DescribeDictsWithContext(context.Background(), request)
}

// DescribeDicts
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 返回属于当前用户的所有自定义词库列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_VALUERANGEERROR = "InvalidParameterValue.ValueRangeError"
//  RESOURCEINSUFFICIENT_QUOTARUNOUT = "ResourceInsufficient.QuotaRunOut"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
//  UNAUTHORIZEDOPERATION_AUTHENTICATEFAILED = "UnauthorizedOperation.AuthenticateFailed"
func (c *Client) DescribeDictsWithContext(ctx context.Context, request *DescribeDictsRequest) (response *DescribeDictsResponse, err error) {
    if request == nil {
        request = NewDescribeDictsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDicts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDictsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWordItemsRequest() (request *DescribeWordItemsRequest) {
    request = &DescribeWordItemsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("nlp", APIVersion, "DescribeWordItems")
    
    
    return
}

func NewDescribeWordItemsResponse() (response *DescribeWordItemsResponse) {
    response = &DescribeWordItemsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWordItems
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 依据自定义词库的ID，查询对应的词条信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  INVALIDPARAMETERVALUE_VALUERANGEERROR = "InvalidParameterValue.ValueRangeError"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
//  UNAUTHORIZEDOPERATION_AUTHENTICATEFAILED = "UnauthorizedOperation.AuthenticateFailed"
func (c *Client) DescribeWordItems(request *DescribeWordItemsRequest) (response *DescribeWordItemsResponse, err error) {
    return c.DescribeWordItemsWithContext(context.Background(), request)
}

// DescribeWordItems
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 依据自定义词库的ID，查询对应的词条信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  INVALIDPARAMETERVALUE_VALUERANGEERROR = "InvalidParameterValue.ValueRangeError"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
//  UNAUTHORIZEDOPERATION_AUTHENTICATEFAILED = "UnauthorizedOperation.AuthenticateFailed"
func (c *Client) DescribeWordItemsWithContext(ctx context.Context, request *DescribeWordItemsRequest) (response *DescribeWordItemsResponse, err error) {
    if request == nil {
        request = NewDescribeWordItemsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWordItems require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWordItemsResponse()
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
func (c *Client) EvaluateSentenceSimilarity(request *EvaluateSentenceSimilarityRequest) (response *EvaluateSentenceSimilarityResponse, err error) {
    return c.EvaluateSentenceSimilarityWithContext(context.Background(), request)
}

// EvaluateSentenceSimilarity
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
func (c *Client) EvaluateSentenceSimilarityWithContext(ctx context.Context, request *EvaluateSentenceSimilarityRequest) (response *EvaluateSentenceSimilarityResponse, err error) {
    if request == nil {
        request = NewEvaluateSentenceSimilarityRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EvaluateSentenceSimilarity require credential")
    }

    request.SetContext(ctx)
    
    response = NewEvaluateSentenceSimilarityResponse()
    err = c.Send(request, response)
    return
}

func NewEvaluateWordSimilarityRequest() (request *EvaluateWordSimilarityRequest) {
    request = &EvaluateWordSimilarityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("nlp", APIVersion, "EvaluateWordSimilarity")
    
    
    return
}

func NewEvaluateWordSimilarityResponse() (response *EvaluateWordSimilarityResponse) {
    response = &EvaluateWordSimilarityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EvaluateWordSimilarity
// 评估两个词语在语义空间的相似程度，为您的场景应用提供有力支持，如关键词过滤、热门话题挖掘等。（目前仅支持中文）
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
func (c *Client) EvaluateWordSimilarity(request *EvaluateWordSimilarityRequest) (response *EvaluateWordSimilarityResponse, err error) {
    return c.EvaluateWordSimilarityWithContext(context.Background(), request)
}

// EvaluateWordSimilarity
// 评估两个词语在语义空间的相似程度，为您的场景应用提供有力支持，如关键词过滤、热门话题挖掘等。（目前仅支持中文）
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
func (c *Client) EvaluateWordSimilarityWithContext(ctx context.Context, request *EvaluateWordSimilarityRequest) (response *EvaluateWordSimilarityResponse, err error) {
    if request == nil {
        request = NewEvaluateWordSimilarityRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EvaluateWordSimilarity require credential")
    }

    request.SetContext(ctx)
    
    response = NewEvaluateWordSimilarityResponse()
    err = c.Send(request, response)
    return
}

func NewGenerateCoupletRequest() (request *GenerateCoupletRequest) {
    request = &GenerateCoupletRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("nlp", APIVersion, "GenerateCouplet")
    
    
    return
}

func NewGenerateCoupletResponse() (response *GenerateCoupletResponse) {
    response = &GenerateCoupletResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GenerateCouplet
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 根据用户输入的命题关键词自动生成一副春联，包括上联、下联和横批。（如需开通请联系商务）
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
func (c *Client) GenerateCouplet(request *GenerateCoupletRequest) (response *GenerateCoupletResponse, err error) {
    return c.GenerateCoupletWithContext(context.Background(), request)
}

// GenerateCouplet
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 根据用户输入的命题关键词自动生成一副春联，包括上联、下联和横批。（如需开通请联系商务）
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
func (c *Client) GenerateCoupletWithContext(ctx context.Context, request *GenerateCoupletRequest) (response *GenerateCoupletResponse, err error) {
    if request == nil {
        request = NewGenerateCoupletRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GenerateCouplet require credential")
    }

    request.SetContext(ctx)
    
    response = NewGenerateCoupletResponse()
    err = c.Send(request, response)
    return
}

func NewGenerateKeywordSentenceRequest() (request *GenerateKeywordSentenceRequest) {
    request = &GenerateKeywordSentenceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("nlp", APIVersion, "GenerateKeywordSentence")
    
    
    return
}

func NewGenerateKeywordSentenceResponse() (response *GenerateKeywordSentenceResponse) {
    response = &GenerateKeywordSentenceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GenerateKeywordSentence
// 提取文本中的关键信息，生成简洁明了的关键句子，便于用户快速获取核心观点。
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
//  INVALIDPARAMETER_SERVICEERROR = "InvalidParameter.ServiceError"
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
func (c *Client) GenerateKeywordSentence(request *GenerateKeywordSentenceRequest) (response *GenerateKeywordSentenceResponse, err error) {
    return c.GenerateKeywordSentenceWithContext(context.Background(), request)
}

// GenerateKeywordSentence
// 提取文本中的关键信息，生成简洁明了的关键句子，便于用户快速获取核心观点。
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
//  INVALIDPARAMETER_SERVICEERROR = "InvalidParameter.ServiceError"
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
func (c *Client) GenerateKeywordSentenceWithContext(ctx context.Context, request *GenerateKeywordSentenceRequest) (response *GenerateKeywordSentenceResponse, err error) {
    if request == nil {
        request = NewGenerateKeywordSentenceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GenerateKeywordSentence require credential")
    }

    request.SetContext(ctx)
    
    response = NewGenerateKeywordSentenceResponse()
    err = c.Send(request, response)
    return
}

func NewGeneratePoetryRequest() (request *GeneratePoetryRequest) {
    request = &GeneratePoetryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("nlp", APIVersion, "GeneratePoetry")
    
    
    return
}

func NewGeneratePoetryResponse() (response *GeneratePoetryResponse) {
    response = &GeneratePoetryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GeneratePoetry
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 根据用户输入的命题关键词自动生成一首七言律诗或五言律诗。（如需开通请联系商务）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  FAILEDOPERATION_NOPOETRY = "FailedOperation.NoPoetry"
//  FAILEDOPERATION_NOTFOUNDDATA = "FailedOperation.NotFoundData"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVICECALLERROR = "InternalError.ServiceCallError"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INVALIDPARAMETERVALUE_GENRE = "InvalidParameterValue.Genre"
//  INVALIDPARAMETERVALUE_POETRYTYPE = "InvalidParameterValue.PoetryType"
//  LIMITEXCEEDED_RESOURCEREACHEDLIMIT = "LimitExceeded.ResourceReachedLimit"
//  RESOURCEINSUFFICIENT_QUOTARUNOUT = "ResourceInsufficient.QuotaRunOut"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_ISOPENING = "ResourceUnavailable.IsOpening"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
func (c *Client) GeneratePoetry(request *GeneratePoetryRequest) (response *GeneratePoetryResponse, err error) {
    return c.GeneratePoetryWithContext(context.Background(), request)
}

// GeneratePoetry
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 根据用户输入的命题关键词自动生成一首七言律诗或五言律诗。（如需开通请联系商务）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  FAILEDOPERATION_NOPOETRY = "FailedOperation.NoPoetry"
//  FAILEDOPERATION_NOTFOUNDDATA = "FailedOperation.NotFoundData"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVICECALLERROR = "InternalError.ServiceCallError"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INVALIDPARAMETERVALUE_GENRE = "InvalidParameterValue.Genre"
//  INVALIDPARAMETERVALUE_POETRYTYPE = "InvalidParameterValue.PoetryType"
//  LIMITEXCEEDED_RESOURCEREACHEDLIMIT = "LimitExceeded.ResourceReachedLimit"
//  RESOURCEINSUFFICIENT_QUOTARUNOUT = "ResourceInsufficient.QuotaRunOut"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_ISOPENING = "ResourceUnavailable.IsOpening"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
func (c *Client) GeneratePoetryWithContext(ctx context.Context, request *GeneratePoetryRequest) (response *GeneratePoetryResponse, err error) {
    if request == nil {
        request = NewGeneratePoetryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GeneratePoetry require credential")
    }

    request.SetContext(ctx)
    
    response = NewGeneratePoetryResponse()
    err = c.Send(request, response)
    return
}

func NewKeywordsExtractionRequest() (request *KeywordsExtractionRequest) {
    request = &KeywordsExtractionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("nlp", APIVersion, "KeywordsExtraction")
    
    
    return
}

func NewKeywordsExtractionResponse() (response *KeywordsExtractionResponse) {
    response = &KeywordsExtractionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// KeywordsExtraction
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 基于关键词提取平台，通过对文本内容进行深度分析，提取出文本内容中的关键信息，为用户实现诸如新闻内容关键词自动提取、评论关键词提取等提供基础服务。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_ILLEGALTEXTERROR = "FailedOperation.IllegalTextError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
//  INVALIDPARAMETERVALUE_TEXTENCODEERROR = "InvalidParameterValue.TextEncodeError"
//  INVALIDPARAMETERVALUE_TEXTFORMATERROR = "InvalidParameterValue.TextFormatError"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  INVALIDPARAMETERVALUE_VALUERANGEERROR = "InvalidParameterValue.ValueRangeError"
//  RESOURCEINSUFFICIENT_QUOTARUNOUT = "ResourceInsufficient.QuotaRunOut"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
//  UNAUTHORIZEDOPERATION_AUTHENTICATEFAILED = "UnauthorizedOperation.AuthenticateFailed"
func (c *Client) KeywordsExtraction(request *KeywordsExtractionRequest) (response *KeywordsExtractionResponse, err error) {
    return c.KeywordsExtractionWithContext(context.Background(), request)
}

// KeywordsExtraction
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 基于关键词提取平台，通过对文本内容进行深度分析，提取出文本内容中的关键信息，为用户实现诸如新闻内容关键词自动提取、评论关键词提取等提供基础服务。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_ILLEGALTEXTERROR = "FailedOperation.IllegalTextError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
//  INVALIDPARAMETERVALUE_TEXTENCODEERROR = "InvalidParameterValue.TextEncodeError"
//  INVALIDPARAMETERVALUE_TEXTFORMATERROR = "InvalidParameterValue.TextFormatError"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  INVALIDPARAMETERVALUE_VALUERANGEERROR = "InvalidParameterValue.ValueRangeError"
//  RESOURCEINSUFFICIENT_QUOTARUNOUT = "ResourceInsufficient.QuotaRunOut"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
//  UNAUTHORIZEDOPERATION_AUTHENTICATEFAILED = "UnauthorizedOperation.AuthenticateFailed"
func (c *Client) KeywordsExtractionWithContext(ctx context.Context, request *KeywordsExtractionRequest) (response *KeywordsExtractionResponse, err error) {
    if request == nil {
        request = NewKeywordsExtractionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("KeywordsExtraction require credential")
    }

    request.SetContext(ctx)
    
    response = NewKeywordsExtractionResponse()
    err = c.Send(request, response)
    return
}

func NewLexicalAnalysisRequest() (request *LexicalAnalysisRequest) {
    request = &LexicalAnalysisRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("nlp", APIVersion, "LexicalAnalysis")
    
    
    return
}

func NewLexicalAnalysisResponse() (response *LexicalAnalysisResponse) {
    response = &LexicalAnalysisResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// LexicalAnalysis
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 词法分析接口提供以下三个功能：
//
// 
//
// 1、智能分词：将连续的自然语言文本，切分成具有语义合理性和完整性的词汇序列；
//
// 
//
// 2、词性标注：为每一个词附上对应的词性，例如名词、代词、形容词、动词等；
//
// 
//
// 3、命名实体识别：快速识别文本中的实体，例如人名、地名、机构名等。
//
// 
//
// 所有的功能均基于千亿级大规模互联网语料进行持续迭代更新，以保证效果不断提升，用户无需担心新词发现、歧义消除、调用性能等问题。目前词法分析已经在泛互联网、金融、政务等不同垂直领域提供业务支持，并取得良好的效果。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ILLEGALTEXTERROR = "FailedOperation.IllegalTextError"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
//  INVALIDPARAMETERVALUE_TEXTENCODEERROR = "InvalidParameterValue.TextEncodeError"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  INVALIDPARAMETERVALUE_VALUERANGEERROR = "InvalidParameterValue.ValueRangeError"
//  RESOURCEINSUFFICIENT_QUOTARUNOUT = "ResourceInsufficient.QuotaRunOut"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
//  UNAUTHORIZEDOPERATION_AUTHENTICATEFAILED = "UnauthorizedOperation.AuthenticateFailed"
func (c *Client) LexicalAnalysis(request *LexicalAnalysisRequest) (response *LexicalAnalysisResponse, err error) {
    return c.LexicalAnalysisWithContext(context.Background(), request)
}

// LexicalAnalysis
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 词法分析接口提供以下三个功能：
//
// 
//
// 1、智能分词：将连续的自然语言文本，切分成具有语义合理性和完整性的词汇序列；
//
// 
//
// 2、词性标注：为每一个词附上对应的词性，例如名词、代词、形容词、动词等；
//
// 
//
// 3、命名实体识别：快速识别文本中的实体，例如人名、地名、机构名等。
//
// 
//
// 所有的功能均基于千亿级大规模互联网语料进行持续迭代更新，以保证效果不断提升，用户无需担心新词发现、歧义消除、调用性能等问题。目前词法分析已经在泛互联网、金融、政务等不同垂直领域提供业务支持，并取得良好的效果。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ILLEGALTEXTERROR = "FailedOperation.IllegalTextError"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
//  INVALIDPARAMETERVALUE_TEXTENCODEERROR = "InvalidParameterValue.TextEncodeError"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  INVALIDPARAMETERVALUE_VALUERANGEERROR = "InvalidParameterValue.ValueRangeError"
//  RESOURCEINSUFFICIENT_QUOTARUNOUT = "ResourceInsufficient.QuotaRunOut"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
//  UNAUTHORIZEDOPERATION_AUTHENTICATEFAILED = "UnauthorizedOperation.AuthenticateFailed"
func (c *Client) LexicalAnalysisWithContext(ctx context.Context, request *LexicalAnalysisRequest) (response *LexicalAnalysisResponse, err error) {
    if request == nil {
        request = NewLexicalAnalysisRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("LexicalAnalysis require credential")
    }

    request.SetContext(ctx)
    
    response = NewLexicalAnalysisResponse()
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
func (c *Client) ParseWordsWithContext(ctx context.Context, request *ParseWordsRequest) (response *ParseWordsResponse, err error) {
    if request == nil {
        request = NewParseWordsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ParseWords require credential")
    }

    request.SetContext(ctx)
    
    response = NewParseWordsResponse()
    err = c.Send(request, response)
    return
}

func NewRetrieveSimilarWordsRequest() (request *RetrieveSimilarWordsRequest) {
    request = &RetrieveSimilarWordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("nlp", APIVersion, "RetrieveSimilarWords")
    
    
    return
}

func NewRetrieveSimilarWordsResponse() (response *RetrieveSimilarWordsResponse) {
    response = &RetrieveSimilarWordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RetrieveSimilarWords
// 基于大数据和深度学习技术，可以快速地找到与给定词语高度相似的其他词语，有助于提高搜索和推荐的准确性。（目前仅支持中文）
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
func (c *Client) RetrieveSimilarWords(request *RetrieveSimilarWordsRequest) (response *RetrieveSimilarWordsResponse, err error) {
    return c.RetrieveSimilarWordsWithContext(context.Background(), request)
}

// RetrieveSimilarWords
// 基于大数据和深度学习技术，可以快速地找到与给定词语高度相似的其他词语，有助于提高搜索和推荐的准确性。（目前仅支持中文）
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
func (c *Client) RetrieveSimilarWordsWithContext(ctx context.Context, request *RetrieveSimilarWordsRequest) (response *RetrieveSimilarWordsResponse, err error) {
    if request == nil {
        request = NewRetrieveSimilarWordsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RetrieveSimilarWords require credential")
    }

    request.SetContext(ctx)
    
    response = NewRetrieveSimilarWordsResponse()
    err = c.Send(request, response)
    return
}

func NewSearchWordItemsRequest() (request *SearchWordItemsRequest) {
    request = &SearchWordItemsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("nlp", APIVersion, "SearchWordItems")
    
    
    return
}

func NewSearchWordItemsResponse() (response *SearchWordItemsResponse) {
    response = &SearchWordItemsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SearchWordItems
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 查询指定自定义词库中的词条是否存在。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
//  INVALIDPARAMETERVALUE_VALUERANGEERROR = "InvalidParameterValue.ValueRangeError"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
func (c *Client) SearchWordItems(request *SearchWordItemsRequest) (response *SearchWordItemsResponse, err error) {
    return c.SearchWordItemsWithContext(context.Background(), request)
}

// SearchWordItems
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 查询指定自定义词库中的词条是否存在。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
//  INVALIDPARAMETERVALUE_VALUERANGEERROR = "InvalidParameterValue.ValueRangeError"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
func (c *Client) SearchWordItemsWithContext(ctx context.Context, request *SearchWordItemsRequest) (response *SearchWordItemsResponse, err error) {
    if request == nil {
        request = NewSearchWordItemsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchWordItems require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchWordItemsResponse()
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
func (c *Client) SentenceCorrectionWithContext(ctx context.Context, request *SentenceCorrectionRequest) (response *SentenceCorrectionResponse, err error) {
    if request == nil {
        request = NewSentenceCorrectionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SentenceCorrection require credential")
    }

    request.SetContext(ctx)
    
    response = NewSentenceCorrectionResponse()
    err = c.Send(request, response)
    return
}

func NewSentenceEmbeddingRequest() (request *SentenceEmbeddingRequest) {
    request = &SentenceEmbeddingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("nlp", APIVersion, "SentenceEmbedding")
    
    
    return
}

func NewSentenceEmbeddingResponse() (response *SentenceEmbeddingResponse) {
    response = &SentenceEmbeddingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SentenceEmbedding
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 句向量接口能够将输入的句子映射成一个固定维度的向量，用来表示这个句子的语义特征，可用于文本聚类、文本相似度、文本分类等任务，能够显著提高它们的效果。
//
// 
//
// 该句向量服务由腾讯云自然语言处理团队联合微信智言团队共同打造，基于千亿级大规模互联网语料并采用Bert等领先的深度神经网络模型训练而成，在腾讯内部诸多业务的NLP任务上实测效果显著。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ILLEGALTEXTERROR = "FailedOperation.IllegalTextError"
//  FAILEDOPERATION_TEXTEMBEDDINGFAILED = "FailedOperation.TextEmbeddingFailed"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
//  INVALIDPARAMETERVALUE_TEXTENCODEERROR = "InvalidParameterValue.TextEncodeError"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  INVALIDPARAMETERVALUE_VALUERANGEERROR = "InvalidParameterValue.ValueRangeError"
//  RESOURCEINSUFFICIENT_QUOTARUNOUT = "ResourceInsufficient.QuotaRunOut"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
func (c *Client) SentenceEmbedding(request *SentenceEmbeddingRequest) (response *SentenceEmbeddingResponse, err error) {
    return c.SentenceEmbeddingWithContext(context.Background(), request)
}

// SentenceEmbedding
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 句向量接口能够将输入的句子映射成一个固定维度的向量，用来表示这个句子的语义特征，可用于文本聚类、文本相似度、文本分类等任务，能够显著提高它们的效果。
//
// 
//
// 该句向量服务由腾讯云自然语言处理团队联合微信智言团队共同打造，基于千亿级大规模互联网语料并采用Bert等领先的深度神经网络模型训练而成，在腾讯内部诸多业务的NLP任务上实测效果显著。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ILLEGALTEXTERROR = "FailedOperation.IllegalTextError"
//  FAILEDOPERATION_TEXTEMBEDDINGFAILED = "FailedOperation.TextEmbeddingFailed"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
//  INVALIDPARAMETERVALUE_TEXTENCODEERROR = "InvalidParameterValue.TextEncodeError"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  INVALIDPARAMETERVALUE_VALUERANGEERROR = "InvalidParameterValue.ValueRangeError"
//  RESOURCEINSUFFICIENT_QUOTARUNOUT = "ResourceInsufficient.QuotaRunOut"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
func (c *Client) SentenceEmbeddingWithContext(ctx context.Context, request *SentenceEmbeddingRequest) (response *SentenceEmbeddingResponse, err error) {
    if request == nil {
        request = NewSentenceEmbeddingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SentenceEmbedding require credential")
    }

    request.SetContext(ctx)
    
    response = NewSentenceEmbeddingResponse()
    err = c.Send(request, response)
    return
}

func NewSentimentAnalysisRequest() (request *SentimentAnalysisRequest) {
    request = &SentimentAnalysisRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("nlp", APIVersion, "SentimentAnalysis")
    
    
    return
}

func NewSentimentAnalysisResponse() (response *SentimentAnalysisResponse) {
    response = &SentimentAnalysisResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SentimentAnalysis
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 情感分析接口能够对带有情感色彩的主观性文本进行分析、处理、归纳和推理，识别出用户的情感倾向，是积极还是消极，并且提供各自概率。
//
// 
//
// 该功能基于千亿级大规模互联网语料和LSTM、BERT等深度神经网络模型进行训练，并持续迭代更新，以保证效果不断提升。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ILLEGALTEXTERROR = "FailedOperation.IllegalTextError"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
//  INVALIDPARAMETERVALUE_TEXTENCODEERROR = "InvalidParameterValue.TextEncodeError"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  INVALIDPARAMETERVALUE_VALUERANGEERROR = "InvalidParameterValue.ValueRangeError"
//  REQUESTLIMITEXCEEDED_UINLIMITEXCEEDED = "RequestLimitExceeded.UinLimitExceeded"
//  RESOURCEINSUFFICIENT_QUOTARUNOUT = "ResourceInsufficient.QuotaRunOut"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
//  UNAUTHORIZEDOPERATION_AUTHENTICATEFAILED = "UnauthorizedOperation.AuthenticateFailed"
func (c *Client) SentimentAnalysis(request *SentimentAnalysisRequest) (response *SentimentAnalysisResponse, err error) {
    return c.SentimentAnalysisWithContext(context.Background(), request)
}

// SentimentAnalysis
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 情感分析接口能够对带有情感色彩的主观性文本进行分析、处理、归纳和推理，识别出用户的情感倾向，是积极还是消极，并且提供各自概率。
//
// 
//
// 该功能基于千亿级大规模互联网语料和LSTM、BERT等深度神经网络模型进行训练，并持续迭代更新，以保证效果不断提升。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ILLEGALTEXTERROR = "FailedOperation.IllegalTextError"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
//  INVALIDPARAMETERVALUE_TEXTENCODEERROR = "InvalidParameterValue.TextEncodeError"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  INVALIDPARAMETERVALUE_VALUERANGEERROR = "InvalidParameterValue.ValueRangeError"
//  REQUESTLIMITEXCEEDED_UINLIMITEXCEEDED = "RequestLimitExceeded.UinLimitExceeded"
//  RESOURCEINSUFFICIENT_QUOTARUNOUT = "ResourceInsufficient.QuotaRunOut"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
//  UNAUTHORIZEDOPERATION_AUTHENTICATEFAILED = "UnauthorizedOperation.AuthenticateFailed"
func (c *Client) SentimentAnalysisWithContext(ctx context.Context, request *SentimentAnalysisRequest) (response *SentimentAnalysisResponse, err error) {
    if request == nil {
        request = NewSentimentAnalysisRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SentimentAnalysis require credential")
    }

    request.SetContext(ctx)
    
    response = NewSentimentAnalysisResponse()
    err = c.Send(request, response)
    return
}

func NewSimilarWordsRequest() (request *SimilarWordsRequest) {
    request = &SimilarWordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("nlp", APIVersion, "SimilarWords")
    
    
    return
}

func NewSimilarWordsResponse() (response *SimilarWordsResponse) {
    response = &SimilarWordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SimilarWords
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 相似词接口能够基于同义词库及词向量技术，检索出与输入词语在语义上最相似的若干个词语，可广泛用于检索系统、问答系统、文档归档等场景。
//
// 可能返回的错误码:
//  FAILEDOPERATION_WORDNOTFOUND = "FailedOperation.WordNotFound"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
//  INVALIDPARAMETERVALUE_TEXTENCODEERROR = "InvalidParameterValue.TextEncodeError"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  INVALIDPARAMETERVALUE_VALUERANGEERROR = "InvalidParameterValue.ValueRangeError"
//  RESOURCEINSUFFICIENT_QUOTARUNOUT = "ResourceInsufficient.QuotaRunOut"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
//  UNAUTHORIZEDOPERATION_AUTHENTICATEFAILED = "UnauthorizedOperation.AuthenticateFailed"
func (c *Client) SimilarWords(request *SimilarWordsRequest) (response *SimilarWordsResponse, err error) {
    return c.SimilarWordsWithContext(context.Background(), request)
}

// SimilarWords
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 相似词接口能够基于同义词库及词向量技术，检索出与输入词语在语义上最相似的若干个词语，可广泛用于检索系统、问答系统、文档归档等场景。
//
// 可能返回的错误码:
//  FAILEDOPERATION_WORDNOTFOUND = "FailedOperation.WordNotFound"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
//  INVALIDPARAMETERVALUE_TEXTENCODEERROR = "InvalidParameterValue.TextEncodeError"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  INVALIDPARAMETERVALUE_VALUERANGEERROR = "InvalidParameterValue.ValueRangeError"
//  RESOURCEINSUFFICIENT_QUOTARUNOUT = "ResourceInsufficient.QuotaRunOut"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
//  UNAUTHORIZEDOPERATION_AUTHENTICATEFAILED = "UnauthorizedOperation.AuthenticateFailed"
func (c *Client) SimilarWordsWithContext(ctx context.Context, request *SimilarWordsRequest) (response *SimilarWordsResponse, err error) {
    if request == nil {
        request = NewSimilarWordsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SimilarWords require credential")
    }

    request.SetContext(ctx)
    
    response = NewSimilarWordsResponse()
    err = c.Send(request, response)
    return
}

func NewTextClassificationRequest() (request *TextClassificationRequest) {
    request = &TextClassificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("nlp", APIVersion, "TextClassification")
    
    
    return
}

func NewTextClassificationResponse() (response *TextClassificationResponse) {
    response = &TextClassificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TextClassification
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 文本分类接口能够对用户输入的文本进行自动分类，将其映射到具体的类目上，用户只需要提供待分类的文本，而无需关注具体实现。
//
// 
//
// 该功能基于千亿级大规模互联网语料和LSTM、BERT等深度神经网络模型进行训练，并持续迭代更新，以保证效果不断提升。
//
// 
//
// 目前已提供：
//
// 
//
// - 通用领域分类体系，二级分类，包括14个分类类目，分别是汽车、科技、健康、体育、旅行、教育、职业、文化、房产、娱乐、女性、奥运、财经以及其他，适用于通用的场景。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_ILLEGALTEXTERROR = "FailedOperation.IllegalTextError"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
//  INVALIDPARAMETERVALUE_TEXTENCODEERROR = "InvalidParameterValue.TextEncodeError"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  INVALIDPARAMETERVALUE_VALUERANGEERROR = "InvalidParameterValue.ValueRangeError"
//  RESOURCEINSUFFICIENT_QUOTARUNOUT = "ResourceInsufficient.QuotaRunOut"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
//  UNAUTHORIZEDOPERATION_AUTHENTICATEFAILED = "UnauthorizedOperation.AuthenticateFailed"
func (c *Client) TextClassification(request *TextClassificationRequest) (response *TextClassificationResponse, err error) {
    return c.TextClassificationWithContext(context.Background(), request)
}

// TextClassification
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 文本分类接口能够对用户输入的文本进行自动分类，将其映射到具体的类目上，用户只需要提供待分类的文本，而无需关注具体实现。
//
// 
//
// 该功能基于千亿级大规模互联网语料和LSTM、BERT等深度神经网络模型进行训练，并持续迭代更新，以保证效果不断提升。
//
// 
//
// 目前已提供：
//
// 
//
// - 通用领域分类体系，二级分类，包括14个分类类目，分别是汽车、科技、健康、体育、旅行、教育、职业、文化、房产、娱乐、女性、奥运、财经以及其他，适用于通用的场景。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_ILLEGALTEXTERROR = "FailedOperation.IllegalTextError"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
//  INVALIDPARAMETERVALUE_TEXTENCODEERROR = "InvalidParameterValue.TextEncodeError"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  INVALIDPARAMETERVALUE_VALUERANGEERROR = "InvalidParameterValue.ValueRangeError"
//  RESOURCEINSUFFICIENT_QUOTARUNOUT = "ResourceInsufficient.QuotaRunOut"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
//  UNAUTHORIZEDOPERATION_AUTHENTICATEFAILED = "UnauthorizedOperation.AuthenticateFailed"
func (c *Client) TextClassificationWithContext(ctx context.Context, request *TextClassificationRequest) (response *TextClassificationResponse, err error) {
    if request == nil {
        request = NewTextClassificationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TextClassification require credential")
    }

    request.SetContext(ctx)
    
    response = NewTextClassificationResponse()
    err = c.Send(request, response)
    return
}

func NewTextCorrectionRequest() (request *TextCorrectionRequest) {
    request = &TextCorrectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("nlp", APIVersion, "TextCorrection")
    
    
    return
}

func NewTextCorrectionResponse() (response *TextCorrectionResponse) {
    response = &TextCorrectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TextCorrection
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 提供对中文文本的自动纠错功能，能够识别输入文本中的错误片段，定位错误并给出正确的文本结果；支持长度不超过2000字符（含标点符号）的长文本纠错。
//
// 
//
// 此功能是基于千亿级大规模互联网语料和LSTM、BERT等深度神经网络模型进行训练，并持续迭代更新，以保证效果不断提升，是搜索引擎、语音识别、内容审核等功能更好运行的基础之一。 
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
//  INVALIDPARAMETERVALUE_TEXTENCODEERROR = "InvalidParameterValue.TextEncodeError"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  RESOURCEINSUFFICIENT_QUOTARUNOUT = "ResourceInsufficient.QuotaRunOut"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
//  UNAUTHORIZEDOPERATION_AUTHENTICATEFAILED = "UnauthorizedOperation.AuthenticateFailed"
func (c *Client) TextCorrection(request *TextCorrectionRequest) (response *TextCorrectionResponse, err error) {
    return c.TextCorrectionWithContext(context.Background(), request)
}

// TextCorrection
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 提供对中文文本的自动纠错功能，能够识别输入文本中的错误片段，定位错误并给出正确的文本结果；支持长度不超过2000字符（含标点符号）的长文本纠错。
//
// 
//
// 此功能是基于千亿级大规模互联网语料和LSTM、BERT等深度神经网络模型进行训练，并持续迭代更新，以保证效果不断提升，是搜索引擎、语音识别、内容审核等功能更好运行的基础之一。 
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
//  INVALIDPARAMETERVALUE_TEXTENCODEERROR = "InvalidParameterValue.TextEncodeError"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  RESOURCEINSUFFICIENT_QUOTARUNOUT = "ResourceInsufficient.QuotaRunOut"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
//  UNAUTHORIZEDOPERATION_AUTHENTICATEFAILED = "UnauthorizedOperation.AuthenticateFailed"
func (c *Client) TextCorrectionWithContext(ctx context.Context, request *TextCorrectionRequest) (response *TextCorrectionResponse, err error) {
    if request == nil {
        request = NewTextCorrectionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TextCorrection require credential")
    }

    request.SetContext(ctx)
    
    response = NewTextCorrectionResponse()
    err = c.Send(request, response)
    return
}

func NewTextCorrectionProRequest() (request *TextCorrectionProRequest) {
    request = &TextCorrectionProRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("nlp", APIVersion, "TextCorrectionPro")
    
    
    return
}

func NewTextCorrectionProResponse() (response *TextCorrectionProResponse) {
    response = &TextCorrectionProResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TextCorrectionPro
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 提供对中文文本的自动纠错功能，能够识别输入文本中的错误片段，定位错误并给出正确的文本结果；支持长度不超过128字符（含标点符号）的长文本纠错。
//
// 
//
// 此功能是基于千亿级大规模互联网语料和LSTM、BERT等深度神经网络模型进行训练，并持续迭代更新，以保证效果不断提升，是搜索引擎、语音识别、内容审核等功能更好运行的基础之一。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ILLEGALTEXTERROR = "FailedOperation.IllegalTextError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_TEXTEMBEDDINGFAILED = "FailedOperation.TextEmbeddingFailed"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
//  INVALIDPARAMETERVALUE_TEXTENCODEERROR = "InvalidParameterValue.TextEncodeError"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  RESOURCEINSUFFICIENT_QUOTARUNOUT = "ResourceInsufficient.QuotaRunOut"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
//  UNAUTHORIZEDOPERATION_AUTHENTICATEFAILED = "UnauthorizedOperation.AuthenticateFailed"
func (c *Client) TextCorrectionPro(request *TextCorrectionProRequest) (response *TextCorrectionProResponse, err error) {
    return c.TextCorrectionProWithContext(context.Background(), request)
}

// TextCorrectionPro
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 提供对中文文本的自动纠错功能，能够识别输入文本中的错误片段，定位错误并给出正确的文本结果；支持长度不超过128字符（含标点符号）的长文本纠错。
//
// 
//
// 此功能是基于千亿级大规模互联网语料和LSTM、BERT等深度神经网络模型进行训练，并持续迭代更新，以保证效果不断提升，是搜索引擎、语音识别、内容审核等功能更好运行的基础之一。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ILLEGALTEXTERROR = "FailedOperation.IllegalTextError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_TEXTEMBEDDINGFAILED = "FailedOperation.TextEmbeddingFailed"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
//  INVALIDPARAMETERVALUE_TEXTENCODEERROR = "InvalidParameterValue.TextEncodeError"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  RESOURCEINSUFFICIENT_QUOTARUNOUT = "ResourceInsufficient.QuotaRunOut"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
//  UNAUTHORIZEDOPERATION_AUTHENTICATEFAILED = "UnauthorizedOperation.AuthenticateFailed"
func (c *Client) TextCorrectionProWithContext(ctx context.Context, request *TextCorrectionProRequest) (response *TextCorrectionProResponse, err error) {
    if request == nil {
        request = NewTextCorrectionProRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TextCorrectionPro require credential")
    }

    request.SetContext(ctx)
    
    response = NewTextCorrectionProResponse()
    err = c.Send(request, response)
    return
}

func NewTextEmbellishRequest() (request *TextEmbellishRequest) {
    request = &TextEmbellishRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("nlp", APIVersion, "TextEmbellish")
    
    
    return
}

func NewTextEmbellishResponse() (response *TextEmbellishResponse) {
    response = &TextEmbellishResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TextEmbellish
// 运用先进的自然语言处理技术，对原始文本进行优化润色，提升文本的通顺性、表达力和语言质量。
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
//  INVALIDPARAMETERVALUE_INVALIDPARAMETER = "InvalidParameterValue.InvalidParameter"
//  INVALIDPARAMETERVALUE_SENSITIVETEXT = "InvalidParameterValue.SensitiveText"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  LIMITEXCEEDED_RESOURCEREACHEDLIMIT = "LimitExceeded.ResourceReachedLimit"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_ISOPENING = "ResourceUnavailable.IsOpening"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
func (c *Client) TextEmbellish(request *TextEmbellishRequest) (response *TextEmbellishResponse, err error) {
    return c.TextEmbellishWithContext(context.Background(), request)
}

// TextEmbellish
// 运用先进的自然语言处理技术，对原始文本进行优化润色，提升文本的通顺性、表达力和语言质量。
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
//  INVALIDPARAMETERVALUE_INVALIDPARAMETER = "InvalidParameterValue.InvalidParameter"
//  INVALIDPARAMETERVALUE_SENSITIVETEXT = "InvalidParameterValue.SensitiveText"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  LIMITEXCEEDED_RESOURCEREACHEDLIMIT = "LimitExceeded.ResourceReachedLimit"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_ISOPENING = "ResourceUnavailable.IsOpening"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
func (c *Client) TextEmbellishWithContext(ctx context.Context, request *TextEmbellishRequest) (response *TextEmbellishResponse, err error) {
    if request == nil {
        request = NewTextEmbellishRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TextEmbellish require credential")
    }

    request.SetContext(ctx)
    
    response = NewTextEmbellishResponse()
    err = c.Send(request, response)
    return
}

func NewTextSimilarityRequest() (request *TextSimilarityRequest) {
    request = &TextSimilarityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("nlp", APIVersion, "TextSimilarity")
    
    
    return
}

func NewTextSimilarityResponse() (response *TextSimilarityResponse) {
    response = &TextSimilarityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TextSimilarity
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 句子相似度接口能够基于深度学习技术来计算一个源句子和多个目标句子的相似度，相似度分值越大的两个句子在语义上越相似。目前仅支持短文本（不超过500字符）的相似度计算，长文本的相似度计算也即将推出。
//
// 
//
// 鉴于句子相似度是一个应用非常广泛的功能，腾讯云自然语言处理团队在Bert等领先的深度神经网络模型的基础上，专门针对文本相似任务进行了优化，并持续迭代更新。基于句子相似度，可以轻松实现诸如文本去重、相似推荐等功能。
//
// 
//
// 接口将以句子数量为单位消耗资源包，而不是调用接口次数为单位。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ILLEGALTEXTERROR = "FailedOperation.IllegalTextError"
//  FAILEDOPERATION_TEXTEMBEDDINGFAILED = "FailedOperation.TextEmbeddingFailed"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
//  INVALIDPARAMETERVALUE_TEXTENCODEERROR = "InvalidParameterValue.TextEncodeError"
//  INVALIDPARAMETERVALUE_TEXTNUMTOOMUCH = "InvalidParameterValue.TextNumTooMuch"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  INVALIDPARAMETERVALUE_VALUERANGEERROR = "InvalidParameterValue.ValueRangeError"
//  RESOURCEINSUFFICIENT_QUOTARUNOUT = "ResourceInsufficient.QuotaRunOut"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
//  UNAUTHORIZEDOPERATION_AUTHENTICATEFAILED = "UnauthorizedOperation.AuthenticateFailed"
func (c *Client) TextSimilarity(request *TextSimilarityRequest) (response *TextSimilarityResponse, err error) {
    return c.TextSimilarityWithContext(context.Background(), request)
}

// TextSimilarity
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 句子相似度接口能够基于深度学习技术来计算一个源句子和多个目标句子的相似度，相似度分值越大的两个句子在语义上越相似。目前仅支持短文本（不超过500字符）的相似度计算，长文本的相似度计算也即将推出。
//
// 
//
// 鉴于句子相似度是一个应用非常广泛的功能，腾讯云自然语言处理团队在Bert等领先的深度神经网络模型的基础上，专门针对文本相似任务进行了优化，并持续迭代更新。基于句子相似度，可以轻松实现诸如文本去重、相似推荐等功能。
//
// 
//
// 接口将以句子数量为单位消耗资源包，而不是调用接口次数为单位。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ILLEGALTEXTERROR = "FailedOperation.IllegalTextError"
//  FAILEDOPERATION_TEXTEMBEDDINGFAILED = "FailedOperation.TextEmbeddingFailed"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
//  INVALIDPARAMETERVALUE_TEXTENCODEERROR = "InvalidParameterValue.TextEncodeError"
//  INVALIDPARAMETERVALUE_TEXTNUMTOOMUCH = "InvalidParameterValue.TextNumTooMuch"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  INVALIDPARAMETERVALUE_VALUERANGEERROR = "InvalidParameterValue.ValueRangeError"
//  RESOURCEINSUFFICIENT_QUOTARUNOUT = "ResourceInsufficient.QuotaRunOut"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
//  UNAUTHORIZEDOPERATION_AUTHENTICATEFAILED = "UnauthorizedOperation.AuthenticateFailed"
func (c *Client) TextSimilarityWithContext(ctx context.Context, request *TextSimilarityRequest) (response *TextSimilarityResponse, err error) {
    if request == nil {
        request = NewTextSimilarityRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TextSimilarity require credential")
    }

    request.SetContext(ctx)
    
    response = NewTextSimilarityResponse()
    err = c.Send(request, response)
    return
}

func NewTextSimilarityProRequest() (request *TextSimilarityProRequest) {
    request = &TextSimilarityProRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("nlp", APIVersion, "TextSimilarityPro")
    
    
    return
}

func NewTextSimilarityProResponse() (response *TextSimilarityProResponse) {
    response = &TextSimilarityProResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TextSimilarityPro
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 句子相似度接口能够基于深度学习技术来计算一个源句子和多个目标句子的相似度，相似度分值越大的两个句子在语义上越相似。目前仅支持短文本（不超过128字符）的相似度计算，长文本的相似度计算也即将推出。
//
// 
//
// 鉴于句子相似度是一个应用非常广泛的功能，腾讯云自然语言处理团队在Bert等领先的深度神经网络模型的基础上，专门针对文本相似任务进行了优化，并持续迭代更新。基于句子相似度，可以轻松实现诸如文本去重、相似推荐等功能。
//
// 
//
// 接口将以句子数量为单位消耗资源包，而不是调用接口次数为单位。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ILLEGALTEXTERROR = "FailedOperation.IllegalTextError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_TEXTEMBEDDINGFAILED = "FailedOperation.TextEmbeddingFailed"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
//  INVALIDPARAMETERVALUE_TEXTENCODEERROR = "InvalidParameterValue.TextEncodeError"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  INVALIDPARAMETERVALUE_VALUERANGEERROR = "InvalidParameterValue.ValueRangeError"
//  RESOURCEINSUFFICIENT_QUOTARUNOUT = "ResourceInsufficient.QuotaRunOut"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
func (c *Client) TextSimilarityPro(request *TextSimilarityProRequest) (response *TextSimilarityProResponse, err error) {
    return c.TextSimilarityProWithContext(context.Background(), request)
}

// TextSimilarityPro
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 句子相似度接口能够基于深度学习技术来计算一个源句子和多个目标句子的相似度，相似度分值越大的两个句子在语义上越相似。目前仅支持短文本（不超过128字符）的相似度计算，长文本的相似度计算也即将推出。
//
// 
//
// 鉴于句子相似度是一个应用非常广泛的功能，腾讯云自然语言处理团队在Bert等领先的深度神经网络模型的基础上，专门针对文本相似任务进行了优化，并持续迭代更新。基于句子相似度，可以轻松实现诸如文本去重、相似推荐等功能。
//
// 
//
// 接口将以句子数量为单位消耗资源包，而不是调用接口次数为单位。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ILLEGALTEXTERROR = "FailedOperation.IllegalTextError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_TEXTEMBEDDINGFAILED = "FailedOperation.TextEmbeddingFailed"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
//  INVALIDPARAMETERVALUE_TEXTENCODEERROR = "InvalidParameterValue.TextEncodeError"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  INVALIDPARAMETERVALUE_VALUERANGEERROR = "InvalidParameterValue.ValueRangeError"
//  RESOURCEINSUFFICIENT_QUOTARUNOUT = "ResourceInsufficient.QuotaRunOut"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
func (c *Client) TextSimilarityProWithContext(ctx context.Context, request *TextSimilarityProRequest) (response *TextSimilarityProResponse, err error) {
    if request == nil {
        request = NewTextSimilarityProRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TextSimilarityPro require credential")
    }

    request.SetContext(ctx)
    
    response = NewTextSimilarityProResponse()
    err = c.Send(request, response)
    return
}

func NewTextWritingRequest() (request *TextWritingRequest) {
    request = &TextWritingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("nlp", APIVersion, "TextWriting")
    
    
    return
}

func NewTextWritingResponse() (response *TextWritingResponse) {
    response = &TextWritingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TextWriting
// 通过自动补全文本片段，帮助用户快速生成高质量、连贯的完整文本，提高创作效率。
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
func (c *Client) TextWriting(request *TextWritingRequest) (response *TextWritingResponse, err error) {
    return c.TextWritingWithContext(context.Background(), request)
}

// TextWriting
// 通过自动补全文本片段，帮助用户快速生成高质量、连贯的完整文本，提高创作效率。
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
func (c *Client) TextWritingWithContext(ctx context.Context, request *TextWritingRequest) (response *TextWritingResponse, err error) {
    if request == nil {
        request = NewTextWritingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TextWriting require credential")
    }

    request.SetContext(ctx)
    
    response = NewTextWritingResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateDictRequest() (request *UpdateDictRequest) {
    request = &UpdateDictRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("nlp", APIVersion, "UpdateDict")
    
    
    return
}

func NewUpdateDictResponse() (response *UpdateDictResponse) {
    response = &UpdateDictResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateDict
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 修改自定义词库元数据信息，包括名称、描述。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
//  UNAUTHORIZEDOPERATION_AUTHENTICATEFAILED = "UnauthorizedOperation.AuthenticateFailed"
func (c *Client) UpdateDict(request *UpdateDictRequest) (response *UpdateDictResponse, err error) {
    return c.UpdateDictWithContext(context.Background(), request)
}

// UpdateDict
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 修改自定义词库元数据信息，包括名称、描述。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
//  UNAUTHORIZEDOPERATION_AUTHENTICATEFAILED = "UnauthorizedOperation.AuthenticateFailed"
func (c *Client) UpdateDictWithContext(ctx context.Context, request *UpdateDictRequest) (response *UpdateDictResponse, err error) {
    if request == nil {
        request = NewUpdateDictRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateDict require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateDictResponse()
    err = c.Send(request, response)
    return
}

func NewWordEmbeddingRequest() (request *WordEmbeddingRequest) {
    request = &WordEmbeddingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("nlp", APIVersion, "WordEmbedding")
    
    
    return
}

func NewWordEmbeddingResponse() (response *WordEmbeddingResponse) {
    response = &WordEmbeddingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// WordEmbedding
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 词向量接口能够将输入的词语映射成一个固定维度的词向量，用来表示这个词语的语义特征。词向量是很多自然语言处理技术的基础，能够显著提高它们的效果。
//
// 
//
// 该词向量服务由腾讯知文自然语言处理团队联合腾讯AI Lab共同打造。使用的词向量基于千亿级大规模互联网语料并采用AI Lab自研的DSG算法训练而成，开源的词向量包含800多万中文词汇，在覆盖率、新鲜度及准确性等三方面性能突出。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ILLEGALTEXTERROR = "FailedOperation.IllegalTextError"
//  FAILEDOPERATION_WORDNOTFOUND = "FailedOperation.WordNotFound"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
//  INVALIDPARAMETERVALUE_TEXTENCODEERROR = "InvalidParameterValue.TextEncodeError"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  INVALIDPARAMETERVALUE_VALUERANGEERROR = "InvalidParameterValue.ValueRangeError"
//  RESOURCEINSUFFICIENT_QUOTARUNOUT = "ResourceInsufficient.QuotaRunOut"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
func (c *Client) WordEmbedding(request *WordEmbeddingRequest) (response *WordEmbeddingResponse, err error) {
    return c.WordEmbeddingWithContext(context.Background(), request)
}

// WordEmbedding
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 词向量接口能够将输入的词语映射成一个固定维度的词向量，用来表示这个词语的语义特征。词向量是很多自然语言处理技术的基础，能够显著提高它们的效果。
//
// 
//
// 该词向量服务由腾讯知文自然语言处理团队联合腾讯AI Lab共同打造。使用的词向量基于千亿级大规模互联网语料并采用AI Lab自研的DSG算法训练而成，开源的词向量包含800多万中文词汇，在覆盖率、新鲜度及准确性等三方面性能突出。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ILLEGALTEXTERROR = "FailedOperation.IllegalTextError"
//  FAILEDOPERATION_WORDNOTFOUND = "FailedOperation.WordNotFound"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
//  INVALIDPARAMETERVALUE_TEXTENCODEERROR = "InvalidParameterValue.TextEncodeError"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  INVALIDPARAMETERVALUE_VALUERANGEERROR = "InvalidParameterValue.ValueRangeError"
//  RESOURCEINSUFFICIENT_QUOTARUNOUT = "ResourceInsufficient.QuotaRunOut"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
func (c *Client) WordEmbeddingWithContext(ctx context.Context, request *WordEmbeddingRequest) (response *WordEmbeddingResponse, err error) {
    if request == nil {
        request = NewWordEmbeddingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("WordEmbedding require credential")
    }

    request.SetContext(ctx)
    
    response = NewWordEmbeddingResponse()
    err = c.Send(request, response)
    return
}

func NewWordSimilarityRequest() (request *WordSimilarityRequest) {
    request = &WordSimilarityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("nlp", APIVersion, "WordSimilarity")
    
    
    return
}

func NewWordSimilarityResponse() (response *WordSimilarityResponse) {
    response = &WordSimilarityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// WordSimilarity
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 词相似度接口能够基于词向量技术来计算两个输入词语的余弦相似度，相似度数值越大的两个词语在语义上越相似。
//
// 可能返回的错误码:
//  FAILEDOPERATION_WORDNOTFOUND = "FailedOperation.WordNotFound"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
//  INVALIDPARAMETERVALUE_TEXTENCODEERROR = "InvalidParameterValue.TextEncodeError"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  RESOURCEINSUFFICIENT_QUOTARUNOUT = "ResourceInsufficient.QuotaRunOut"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
func (c *Client) WordSimilarity(request *WordSimilarityRequest) (response *WordSimilarityResponse, err error) {
    return c.WordSimilarityWithContext(context.Background(), request)
}

// WordSimilarity
// 因业务调整该接口将于北京时间2023年8月1日0点下线，届时该产品功能将无法正常使用，为了避免对您的业务造成影响，请您尽快做好相关业务调整。详见：https://cloud.tencent.com/document/product/271/90711
//
// 
//
// 词相似度接口能够基于词向量技术来计算两个输入词语的余弦相似度，相似度数值越大的两个词语在语义上越相似。
//
// 可能返回的错误码:
//  FAILEDOPERATION_WORDNOTFOUND = "FailedOperation.WordNotFound"
//  INVALIDPARAMETERVALUE_EMPTYVALUEERROR = "InvalidParameterValue.EmptyValueError"
//  INVALIDPARAMETERVALUE_TEXTENCODEERROR = "InvalidParameterValue.TextEncodeError"
//  INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"
//  RESOURCEINSUFFICIENT_QUOTARUNOUT = "ResourceInsufficient.QuotaRunOut"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_SERVICENOTOPENEDERROR = "ResourceUnavailable.ServiceNotOpenedError"
func (c *Client) WordSimilarityWithContext(ctx context.Context, request *WordSimilarityRequest) (response *WordSimilarityResponse, err error) {
    if request == nil {
        request = NewWordSimilarityRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("WordSimilarity require credential")
    }

    request.SetContext(ctx)
    
    response = NewWordSimilarityResponse()
    err = c.Send(request, response)
    return
}
