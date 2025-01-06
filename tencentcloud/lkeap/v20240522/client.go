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

package v20240522

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2024-05-22"

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


func NewCreateAttributeLabelRequest() (request *CreateAttributeLabelRequest) {
    request = &CreateAttributeLabelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "CreateAttributeLabel")
    
    
    return
}

func NewCreateAttributeLabelResponse() (response *CreateAttributeLabelResponse) {
    response = &CreateAttributeLabelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAttributeLabel
// 用于为问答对创建属性标签，以便对内容进行分类和管理。 使用场景：当需要为问答对添加分类标签和属性标记时使用，比如为问答对添加“售后”标签。
func (c *Client) CreateAttributeLabel(request *CreateAttributeLabelRequest) (response *CreateAttributeLabelResponse, err error) {
    return c.CreateAttributeLabelWithContext(context.Background(), request)
}

// CreateAttributeLabel
// 用于为问答对创建属性标签，以便对内容进行分类和管理。 使用场景：当需要为问答对添加分类标签和属性标记时使用，比如为问答对添加“售后”标签。
func (c *Client) CreateAttributeLabelWithContext(ctx context.Context, request *CreateAttributeLabelRequest) (response *CreateAttributeLabelResponse, err error) {
    if request == nil {
        request = NewCreateAttributeLabelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAttributeLabel require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAttributeLabelResponse()
    err = c.Send(request, response)
    return
}

func NewCreateKnowledgeBaseRequest() (request *CreateKnowledgeBaseRequest) {
    request = &CreateKnowledgeBaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "CreateKnowledgeBase")
    
    
    return
}

func NewCreateKnowledgeBaseResponse() (response *CreateKnowledgeBaseResponse) {
    response = &CreateKnowledgeBaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateKnowledgeBase
// 用于在系统中创建一个新的知识库。知识库是一个用于存储和管理知识条目的集合，可以包括文档、问答对、属性标签等。创建知识库后，可以向其中添加各种知识条目，以便在后续的知识检索中使用。 使用场景：当需要在系统中建立一个新的知识库以存储和管理特定领域或项目的知识条目时使用。例如，一个用户可能需要创建一个知识库，以存储用户指南、常见问题解答和技术文档。
func (c *Client) CreateKnowledgeBase(request *CreateKnowledgeBaseRequest) (response *CreateKnowledgeBaseResponse, err error) {
    return c.CreateKnowledgeBaseWithContext(context.Background(), request)
}

// CreateKnowledgeBase
// 用于在系统中创建一个新的知识库。知识库是一个用于存储和管理知识条目的集合，可以包括文档、问答对、属性标签等。创建知识库后，可以向其中添加各种知识条目，以便在后续的知识检索中使用。 使用场景：当需要在系统中建立一个新的知识库以存储和管理特定领域或项目的知识条目时使用。例如，一个用户可能需要创建一个知识库，以存储用户指南、常见问题解答和技术文档。
func (c *Client) CreateKnowledgeBaseWithContext(ctx context.Context, request *CreateKnowledgeBaseRequest) (response *CreateKnowledgeBaseResponse, err error) {
    if request == nil {
        request = NewCreateKnowledgeBaseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateKnowledgeBase require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateKnowledgeBaseResponse()
    err = c.Send(request, response)
    return
}

func NewCreateQARequest() (request *CreateQARequest) {
    request = &CreateQARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "CreateQA")
    
    
    return
}

func NewCreateQAResponse() (response *CreateQAResponse) {
    response = &CreateQAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateQA
// 用于创建新的问答对。问答对可以在SearchKnowledge接口知识检索时提供匹配的答案。 使用场景：当需要添加新的知识点和对应的问答对时使用，比如为产品添加新的常见问题解答。
func (c *Client) CreateQA(request *CreateQARequest) (response *CreateQAResponse, err error) {
    return c.CreateQAWithContext(context.Background(), request)
}

// CreateQA
// 用于创建新的问答对。问答对可以在SearchKnowledge接口知识检索时提供匹配的答案。 使用场景：当需要添加新的知识点和对应的问答对时使用，比如为产品添加新的常见问题解答。
func (c *Client) CreateQAWithContext(ctx context.Context, request *CreateQARequest) (response *CreateQAResponse, err error) {
    if request == nil {
        request = NewCreateQARequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateQA require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateQAResponse()
    err = c.Send(request, response)
    return
}

func NewCreateReconstructDocumentFlowRequest() (request *CreateReconstructDocumentFlowRequest) {
    request = &CreateReconstructDocumentFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "CreateReconstructDocumentFlow")
    
    
    return
}

func NewCreateReconstructDocumentFlowResponse() (response *CreateReconstructDocumentFlowResponse) {
    response = &CreateReconstructDocumentFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateReconstructDocumentFlow
// 本接口为异步接口的发起请求接口，用于发起文档解析任务。
//
// 文档解析支持将图片或PDF文件转换成Markdown格式文件，可解析包括表格、公式、图片、标题、段落、页眉、页脚等内容元素，并将内容智能转换成阅读顺序。
//
// 
//
// 体验期间单账号限制qps仅为1，若有正式接入需要请与产研团队沟通开放。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_FILEDECODEFAILED = "FailedOperation.FileDecodeFailed"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNKNOWFILETYPEERROR = "FailedOperation.UnKnowFileTypeError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) CreateReconstructDocumentFlow(request *CreateReconstructDocumentFlowRequest) (response *CreateReconstructDocumentFlowResponse, err error) {
    return c.CreateReconstructDocumentFlowWithContext(context.Background(), request)
}

// CreateReconstructDocumentFlow
// 本接口为异步接口的发起请求接口，用于发起文档解析任务。
//
// 文档解析支持将图片或PDF文件转换成Markdown格式文件，可解析包括表格、公式、图片、标题、段落、页眉、页脚等内容元素，并将内容智能转换成阅读顺序。
//
// 
//
// 体验期间单账号限制qps仅为1，若有正式接入需要请与产研团队沟通开放。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_FILEDECODEFAILED = "FailedOperation.FileDecodeFailed"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNKNOWFILETYPEERROR = "FailedOperation.UnKnowFileTypeError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) CreateReconstructDocumentFlowWithContext(ctx context.Context, request *CreateReconstructDocumentFlowRequest) (response *CreateReconstructDocumentFlowResponse, err error) {
    if request == nil {
        request = NewCreateReconstructDocumentFlowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateReconstructDocumentFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateReconstructDocumentFlowResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSplitDocumentFlowRequest() (request *CreateSplitDocumentFlowRequest) {
    request = &CreateSplitDocumentFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "CreateSplitDocumentFlow")
    
    
    return
}

func NewCreateSplitDocumentFlowResponse() (response *CreateSplitDocumentFlowResponse) {
    response = &CreateSplitDocumentFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSplitDocumentFlow
// 创建文档拆分任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateSplitDocumentFlow(request *CreateSplitDocumentFlowRequest) (response *CreateSplitDocumentFlowResponse, err error) {
    return c.CreateSplitDocumentFlowWithContext(context.Background(), request)
}

// CreateSplitDocumentFlow
// 创建文档拆分任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateSplitDocumentFlowWithContext(ctx context.Context, request *CreateSplitDocumentFlowRequest) (response *CreateSplitDocumentFlowResponse, err error) {
    if request == nil {
        request = NewCreateSplitDocumentFlowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSplitDocumentFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSplitDocumentFlowResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAttributeLabelsRequest() (request *DeleteAttributeLabelsRequest) {
    request = &DeleteAttributeLabelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "DeleteAttributeLabels")
    
    
    return
}

func NewDeleteAttributeLabelsResponse() (response *DeleteAttributeLabelsResponse) {
    response = &DeleteAttributeLabelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAttributeLabels
// 用于删除不再需要的属性标签。 使用场景：当某个标签不再使用时，可以将其删除以保持标签系统的整洁。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteAttributeLabels(request *DeleteAttributeLabelsRequest) (response *DeleteAttributeLabelsResponse, err error) {
    return c.DeleteAttributeLabelsWithContext(context.Background(), request)
}

// DeleteAttributeLabels
// 用于删除不再需要的属性标签。 使用场景：当某个标签不再使用时，可以将其删除以保持标签系统的整洁。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteAttributeLabelsWithContext(ctx context.Context, request *DeleteAttributeLabelsRequest) (response *DeleteAttributeLabelsResponse, err error) {
    if request == nil {
        request = NewDeleteAttributeLabelsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAttributeLabels require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAttributeLabelsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDocsRequest() (request *DeleteDocsRequest) {
    request = &DeleteDocsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "DeleteDocs")
    
    
    return
}

func NewDeleteDocsResponse() (response *DeleteDocsResponse) {
    response = &DeleteDocsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDocs
// 用于删除已有的文档。 使用场景：当某个文档不再需要时，可以将其删除以保持文档库的整洁。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteDocs(request *DeleteDocsRequest) (response *DeleteDocsResponse, err error) {
    return c.DeleteDocsWithContext(context.Background(), request)
}

// DeleteDocs
// 用于删除已有的文档。 使用场景：当某个文档不再需要时，可以将其删除以保持文档库的整洁。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteDocsWithContext(ctx context.Context, request *DeleteDocsRequest) (response *DeleteDocsResponse, err error) {
    if request == nil {
        request = NewDeleteDocsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDocs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDocsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteKnowledgeBaseRequest() (request *DeleteKnowledgeBaseRequest) {
    request = &DeleteKnowledgeBaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "DeleteKnowledgeBase")
    
    
    return
}

func NewDeleteKnowledgeBaseResponse() (response *DeleteKnowledgeBaseResponse) {
    response = &DeleteKnowledgeBaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteKnowledgeBase
// 用于从系统中删除一个现有的知识库。删除知识库将移除该知识库及其所有关联的知识条目（如文档、问答对、属性标签等）。该操作是不可逆的，请在执行前确认是否需要删除。**使用场景**：当某个知识库不再需要时，可以使用此接口将其从系统中删除。例如，一个项目结束后，其相关的知识库可能不再需要存储，可以使用该接口进行删除。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteKnowledgeBase(request *DeleteKnowledgeBaseRequest) (response *DeleteKnowledgeBaseResponse, err error) {
    return c.DeleteKnowledgeBaseWithContext(context.Background(), request)
}

// DeleteKnowledgeBase
// 用于从系统中删除一个现有的知识库。删除知识库将移除该知识库及其所有关联的知识条目（如文档、问答对、属性标签等）。该操作是不可逆的，请在执行前确认是否需要删除。**使用场景**：当某个知识库不再需要时，可以使用此接口将其从系统中删除。例如，一个项目结束后，其相关的知识库可能不再需要存储，可以使用该接口进行删除。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteKnowledgeBaseWithContext(ctx context.Context, request *DeleteKnowledgeBaseRequest) (response *DeleteKnowledgeBaseResponse, err error) {
    if request == nil {
        request = NewDeleteKnowledgeBaseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteKnowledgeBase require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteKnowledgeBaseResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteQAsRequest() (request *DeleteQAsRequest) {
    request = &DeleteQAsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "DeleteQAs")
    
    
    return
}

func NewDeleteQAsResponse() (response *DeleteQAsResponse) {
    response = &DeleteQAsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteQAs
// 用于删除已有的问答对。 使用场景：当某个问答对不再适用或需要移除时使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteQAs(request *DeleteQAsRequest) (response *DeleteQAsResponse, err error) {
    return c.DeleteQAsWithContext(context.Background(), request)
}

// DeleteQAs
// 用于删除已有的问答对。 使用场景：当某个问答对不再适用或需要移除时使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteQAsWithContext(ctx context.Context, request *DeleteQAsRequest) (response *DeleteQAsResponse, err error) {
    if request == nil {
        request = NewDeleteQAsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteQAs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteQAsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDocRequest() (request *DescribeDocRequest) {
    request = &DescribeDocRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "DescribeDoc")
    
    
    return
}

func NewDescribeDocResponse() (response *DescribeDocResponse) {
    response = &DescribeDocResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDoc
// 用于查询特定文档的详细信息。 使用场景：当需要查看某个文档的具体内容和属性时使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeDoc(request *DescribeDocRequest) (response *DescribeDocResponse, err error) {
    return c.DescribeDocWithContext(context.Background(), request)
}

// DescribeDoc
// 用于查询特定文档的详细信息。 使用场景：当需要查看某个文档的具体内容和属性时使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeDocWithContext(ctx context.Context, request *DescribeDocRequest) (response *DescribeDocResponse, err error) {
    if request == nil {
        request = NewDescribeDocRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDoc require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDocResponse()
    err = c.Send(request, response)
    return
}

func NewGetEmbeddingRequest() (request *GetEmbeddingRequest) {
    request = &GetEmbeddingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "GetEmbedding")
    
    
    return
}

func NewGetEmbeddingResponse() (response *GetEmbeddingResponse) {
    response = &GetEmbeddingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetEmbedding
// 本接口（GetEmbedding）调用文本表示模型，将文本转化为用数值表示的向量形式，可用于文本检索、信息推荐、知识挖掘等场景。
//
// 本接口（GetEmbedding）有单账号调用上限控制，如您有提高并发限制的需求请 [联系我们](https://cloud.tencent.com/act/event/Online_service) 。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetEmbedding(request *GetEmbeddingRequest) (response *GetEmbeddingResponse, err error) {
    return c.GetEmbeddingWithContext(context.Background(), request)
}

// GetEmbedding
// 本接口（GetEmbedding）调用文本表示模型，将文本转化为用数值表示的向量形式，可用于文本检索、信息推荐、知识挖掘等场景。
//
// 本接口（GetEmbedding）有单账号调用上限控制，如您有提高并发限制的需求请 [联系我们](https://cloud.tencent.com/act/event/Online_service) 。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetEmbeddingWithContext(ctx context.Context, request *GetEmbeddingRequest) (response *GetEmbeddingResponse, err error) {
    if request == nil {
        request = NewGetEmbeddingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetEmbedding require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetEmbeddingResponse()
    err = c.Send(request, response)
    return
}

func NewGetReconstructDocumentResultRequest() (request *GetReconstructDocumentResultRequest) {
    request = &GetReconstructDocumentResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "GetReconstructDocumentResult")
    
    
    return
}

func NewGetReconstructDocumentResultResponse() (response *GetReconstructDocumentResultResponse) {
    response = &GetReconstructDocumentResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetReconstructDocumentResult
// 本接口为异步接口的查询结果接口，用于获取文档解析处理结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) GetReconstructDocumentResult(request *GetReconstructDocumentResultRequest) (response *GetReconstructDocumentResultResponse, err error) {
    return c.GetReconstructDocumentResultWithContext(context.Background(), request)
}

// GetReconstructDocumentResult
// 本接口为异步接口的查询结果接口，用于获取文档解析处理结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) GetReconstructDocumentResultWithContext(ctx context.Context, request *GetReconstructDocumentResultRequest) (response *GetReconstructDocumentResultResponse, err error) {
    if request == nil {
        request = NewGetReconstructDocumentResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetReconstructDocumentResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetReconstructDocumentResultResponse()
    err = c.Send(request, response)
    return
}

func NewGetSplitDocumentResultRequest() (request *GetSplitDocumentResultRequest) {
    request = &GetSplitDocumentResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "GetSplitDocumentResult")
    
    
    return
}

func NewGetSplitDocumentResultResponse() (response *GetSplitDocumentResultResponse) {
    response = &GetSplitDocumentResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetSplitDocumentResult
// 查询文档拆分任务结果
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) GetSplitDocumentResult(request *GetSplitDocumentResultRequest) (response *GetSplitDocumentResultResponse, err error) {
    return c.GetSplitDocumentResultWithContext(context.Background(), request)
}

// GetSplitDocumentResult
// 查询文档拆分任务结果
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) GetSplitDocumentResultWithContext(ctx context.Context, request *GetSplitDocumentResultRequest) (response *GetSplitDocumentResultResponse, err error) {
    if request == nil {
        request = NewGetSplitDocumentResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetSplitDocumentResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetSplitDocumentResultResponse()
    err = c.Send(request, response)
    return
}

func NewImportQAsRequest() (request *ImportQAsRequest) {
    request = &ImportQAsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "ImportQAs")
    
    
    return
}

func NewImportQAsResponse() (response *ImportQAsResponse) {
    response = &ImportQAsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ImportQAs
// 用于批量导入问答对，最多支持50,000条数据导入。通过此接口，可以将多个问答对一次性导入系统中，以便在后续的知识检索中使用。导入的问答对可以来自外部系统、文件或其他数据源。使用场景：当需要一次性导入大量的问答对时使用。例如，一个公司可能需要将其产品的常见问题解答从一个文档或外部系统批量导入到知识库中，以便用户可以通过知识检索系统进行查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ImportQAs(request *ImportQAsRequest) (response *ImportQAsResponse, err error) {
    return c.ImportQAsWithContext(context.Background(), request)
}

// ImportQAs
// 用于批量导入问答对，最多支持50,000条数据导入。通过此接口，可以将多个问答对一次性导入系统中，以便在后续的知识检索中使用。导入的问答对可以来自外部系统、文件或其他数据源。使用场景：当需要一次性导入大量的问答对时使用。例如，一个公司可能需要将其产品的常见问题解答从一个文档或外部系统批量导入到知识库中，以便用户可以通过知识检索系统进行查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ImportQAsWithContext(ctx context.Context, request *ImportQAsRequest) (response *ImportQAsResponse, err error) {
    if request == nil {
        request = NewImportQAsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImportQAs require credential")
    }

    request.SetContext(ctx)
    
    response = NewImportQAsResponse()
    err = c.Send(request, response)
    return
}

func NewListAttributeLabelsRequest() (request *ListAttributeLabelsRequest) {
    request = &ListAttributeLabelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "ListAttributeLabels")
    
    
    return
}

func NewListAttributeLabelsResponse() (response *ListAttributeLabelsResponse) {
    response = &ListAttributeLabelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListAttributeLabels
// 用于获取所有属性标签的列表。 使用场景：用于查看当前系统中所有已有的属性标签，方便进行管理和维护。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListAttributeLabels(request *ListAttributeLabelsRequest) (response *ListAttributeLabelsResponse, err error) {
    return c.ListAttributeLabelsWithContext(context.Background(), request)
}

// ListAttributeLabels
// 用于获取所有属性标签的列表。 使用场景：用于查看当前系统中所有已有的属性标签，方便进行管理和维护。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListAttributeLabelsWithContext(ctx context.Context, request *ListAttributeLabelsRequest) (response *ListAttributeLabelsResponse, err error) {
    if request == nil {
        request = NewListAttributeLabelsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAttributeLabels require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAttributeLabelsResponse()
    err = c.Send(request, response)
    return
}

func NewListDocsRequest() (request *ListDocsRequest) {
    request = &ListDocsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "ListDocs")
    
    
    return
}

func NewListDocsResponse() (response *ListDocsResponse) {
    response = &ListDocsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListDocs
// 用于获取所有文档的列表。 使用场景：用于查看当前系统中所有已有的文档，方便进行管理和维护。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListDocs(request *ListDocsRequest) (response *ListDocsResponse, err error) {
    return c.ListDocsWithContext(context.Background(), request)
}

// ListDocs
// 用于获取所有文档的列表。 使用场景：用于查看当前系统中所有已有的文档，方便进行管理和维护。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListDocsWithContext(ctx context.Context, request *ListDocsRequest) (response *ListDocsResponse, err error) {
    if request == nil {
        request = NewListDocsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListDocs require credential")
    }

    request.SetContext(ctx)
    
    response = NewListDocsResponse()
    err = c.Send(request, response)
    return
}

func NewListQAsRequest() (request *ListQAsRequest) {
    request = &ListQAsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "ListQAs")
    
    
    return
}

func NewListQAsResponse() (response *ListQAsResponse) {
    response = &ListQAsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListQAs
// 用于获取所有问答对的列表。 使用场景：用于查看当前系统中所有已有的问答对，方便进行管理和维护。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListQAs(request *ListQAsRequest) (response *ListQAsResponse, err error) {
    return c.ListQAsWithContext(context.Background(), request)
}

// ListQAs
// 用于获取所有问答对的列表。 使用场景：用于查看当前系统中所有已有的问答对，方便进行管理和维护。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListQAsWithContext(ctx context.Context, request *ListQAsRequest) (response *ListQAsResponse, err error) {
    if request == nil {
        request = NewListQAsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListQAs require credential")
    }

    request.SetContext(ctx)
    
    response = NewListQAsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAttributeLabelRequest() (request *ModifyAttributeLabelRequest) {
    request = &ModifyAttributeLabelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "ModifyAttributeLabel")
    
    
    return
}

func NewModifyAttributeLabelResponse() (response *ModifyAttributeLabelResponse) {
    response = &ModifyAttributeLabelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAttributeLabel
// 用于修改已有的属性标签。 使用场景：当需要更改属性标签的名称或描述时使用，比如将“售后”标签改为“售前”。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyAttributeLabel(request *ModifyAttributeLabelRequest) (response *ModifyAttributeLabelResponse, err error) {
    return c.ModifyAttributeLabelWithContext(context.Background(), request)
}

// ModifyAttributeLabel
// 用于修改已有的属性标签。 使用场景：当需要更改属性标签的名称或描述时使用，比如将“售后”标签改为“售前”。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyAttributeLabelWithContext(ctx context.Context, request *ModifyAttributeLabelRequest) (response *ModifyAttributeLabelResponse, err error) {
    if request == nil {
        request = NewModifyAttributeLabelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAttributeLabel require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAttributeLabelResponse()
    err = c.Send(request, response)
    return
}

func NewModifyQARequest() (request *ModifyQARequest) {
    request = &ModifyQARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "ModifyQA")
    
    
    return
}

func NewModifyQAResponse() (response *ModifyQAResponse) {
    response = &ModifyQAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyQA
// 用于修改已有的问答对。 使用场景：当需要更新问答对的内容或答案时使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyQA(request *ModifyQARequest) (response *ModifyQAResponse, err error) {
    return c.ModifyQAWithContext(context.Background(), request)
}

// ModifyQA
// 用于修改已有的问答对。 使用场景：当需要更新问答对的内容或答案时使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyQAWithContext(ctx context.Context, request *ModifyQARequest) (response *ModifyQAResponse, err error) {
    if request == nil {
        request = NewModifyQARequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyQA require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyQAResponse()
    err = c.Send(request, response)
    return
}

func NewReconstructDocumentSSERequest() (request *ReconstructDocumentSSERequest) {
    request = &ReconstructDocumentSSERequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "ReconstructDocumentSSE")
    
    
    return
}

func NewReconstructDocumentSSEResponse() (response *ReconstructDocumentSSEResponse) {
    response = &ReconstructDocumentSSEResponse{} 
    return

}

// ReconstructDocumentSSE
// 准实时文档解析接口，使用HTTP SSE 协议通信。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NONSUPPORTPARSE = "FailedOperation.NonsupportParse"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  FAILEDOPERATION_UPLOADRESULTFILEFAILED = "FailedOperation.UploadResultFileFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_EXCEEDEDMAXPAGESERROR = "LimitExceeded.ExceededMaxPagesError"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) ReconstructDocumentSSE(request *ReconstructDocumentSSERequest) (response *ReconstructDocumentSSEResponse, err error) {
    return c.ReconstructDocumentSSEWithContext(context.Background(), request)
}

// ReconstructDocumentSSE
// 准实时文档解析接口，使用HTTP SSE 协议通信。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NONSUPPORTPARSE = "FailedOperation.NonsupportParse"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  FAILEDOPERATION_UPLOADRESULTFILEFAILED = "FailedOperation.UploadResultFileFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_EXCEEDEDMAXPAGESERROR = "LimitExceeded.ExceededMaxPagesError"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) ReconstructDocumentSSEWithContext(ctx context.Context, request *ReconstructDocumentSSERequest) (response *ReconstructDocumentSSEResponse, err error) {
    if request == nil {
        request = NewReconstructDocumentSSERequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReconstructDocumentSSE require credential")
    }

    request.SetContext(ctx)
    
    response = NewReconstructDocumentSSEResponse()
    err = c.Send(request, response)
    return
}

func NewRetrieveKnowledgeRequest() (request *RetrieveKnowledgeRequest) {
    request = &RetrieveKnowledgeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "RetrieveKnowledge")
    
    
    return
}

func NewRetrieveKnowledgeResponse() (response *RetrieveKnowledgeResponse) {
    response = &RetrieveKnowledgeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RetrieveKnowledge
// 用于检索知识库中的文档和问答对内容。 使用场景：适用于查询长期存储在知识库中的文档和问答对，比如产品手册、用户指南等内容的检索。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NONSUPPORTPARSE = "FailedOperation.NonsupportParse"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  FAILEDOPERATION_UPLOADRESULTFILEFAILED = "FailedOperation.UploadResultFileFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_EXCEEDEDMAXPAGESERROR = "LimitExceeded.ExceededMaxPagesError"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RetrieveKnowledge(request *RetrieveKnowledgeRequest) (response *RetrieveKnowledgeResponse, err error) {
    return c.RetrieveKnowledgeWithContext(context.Background(), request)
}

// RetrieveKnowledge
// 用于检索知识库中的文档和问答对内容。 使用场景：适用于查询长期存储在知识库中的文档和问答对，比如产品手册、用户指南等内容的检索。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NONSUPPORTPARSE = "FailedOperation.NonsupportParse"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  FAILEDOPERATION_UPLOADRESULTFILEFAILED = "FailedOperation.UploadResultFileFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_EXCEEDEDMAXPAGESERROR = "LimitExceeded.ExceededMaxPagesError"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RetrieveKnowledgeWithContext(ctx context.Context, request *RetrieveKnowledgeRequest) (response *RetrieveKnowledgeResponse, err error) {
    if request == nil {
        request = NewRetrieveKnowledgeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RetrieveKnowledge require credential")
    }

    request.SetContext(ctx)
    
    response = NewRetrieveKnowledgeResponse()
    err = c.Send(request, response)
    return
}

func NewRunRerankRequest() (request *RunRerankRequest) {
    request = &RunRerankRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "RunRerank")
    
    
    return
}

func NewRunRerankResponse() (response *RunRerankResponse) {
    response = &RunRerankResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RunRerank
// 基于知识引擎精调模型技术的rerank模型，支持对多路召回的结果进行重排序，根据query与切片内容的相关性，按分数由高到低对切片进行排序，并输出对应的打分结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RunRerank(request *RunRerankRequest) (response *RunRerankResponse, err error) {
    return c.RunRerankWithContext(context.Background(), request)
}

// RunRerank
// 基于知识引擎精调模型技术的rerank模型，支持对多路召回的结果进行重排序，根据query与切片内容的相关性，按分数由高到低对切片进行排序，并输出对应的打分结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RunRerankWithContext(ctx context.Context, request *RunRerankRequest) (response *RunRerankResponse, err error) {
    if request == nil {
        request = NewRunRerankRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunRerank require credential")
    }

    request.SetContext(ctx)
    
    response = NewRunRerankResponse()
    err = c.Send(request, response)
    return
}

func NewUploadDocRequest() (request *UploadDocRequest) {
    request = &UploadDocRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "UploadDoc")
    
    
    return
}

func NewUploadDocResponse() (response *UploadDocResponse) {
    response = &UploadDocResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UploadDoc
// 用于上传文档内容。上传的文档将存储在知识库中，可以通过SearchKnowledge知识库内容检索接口进行检索。 
//
// 使用场景：适用于需要长期存储和检索的文档内容，如产品手册、用户指南等。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UploadDoc(request *UploadDocRequest) (response *UploadDocResponse, err error) {
    return c.UploadDocWithContext(context.Background(), request)
}

// UploadDoc
// 用于上传文档内容。上传的文档将存储在知识库中，可以通过SearchKnowledge知识库内容检索接口进行检索。 
//
// 使用场景：适用于需要长期存储和检索的文档内容，如产品手册、用户指南等。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UploadDocWithContext(ctx context.Context, request *UploadDocRequest) (response *UploadDocResponse, err error) {
    if request == nil {
        request = NewUploadDocRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadDoc require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadDocResponse()
    err = c.Send(request, response)
    return
}

func NewUploadDocRealtimeRequest() (request *UploadDocRealtimeRequest) {
    request = &UploadDocRealtimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "UploadDocRealtime")
    
    
    return
}

func NewUploadDocRealtimeResponse() (response *UploadDocRealtimeResponse) {
    response = &UploadDocRealtimeResponse{} 
    return

}

// UploadDocRealtime
// 用于上传实时文档内容。实时文档在上传后可以立即通过SearchRealtime进行实时检索，适用于在会话中对文档进行问答的场景。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UploadDocRealtime(request *UploadDocRealtimeRequest) (response *UploadDocRealtimeResponse, err error) {
    return c.UploadDocRealtimeWithContext(context.Background(), request)
}

// UploadDocRealtime
// 用于上传实时文档内容。实时文档在上传后可以立即通过SearchRealtime进行实时检索，适用于在会话中对文档进行问答的场景。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UploadDocRealtimeWithContext(ctx context.Context, request *UploadDocRealtimeRequest) (response *UploadDocRealtimeResponse, err error) {
    if request == nil {
        request = NewUploadDocRealtimeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadDocRealtime require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadDocRealtimeResponse()
    err = c.Send(request, response)
    return
}
