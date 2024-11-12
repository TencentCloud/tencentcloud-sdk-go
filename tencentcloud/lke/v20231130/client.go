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

package v20231130

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2023-11-30"

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


func NewCheckAttributeLabelExistRequest() (request *CheckAttributeLabelExistRequest) {
    request = &CheckAttributeLabelExistRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "CheckAttributeLabelExist")
    
    
    return
}

func NewCheckAttributeLabelExistResponse() (response *CheckAttributeLabelExistResponse) {
    response = &CheckAttributeLabelExistResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckAttributeLabelExist
// 检查属性下的标签名是否存在
func (c *Client) CheckAttributeLabelExist(request *CheckAttributeLabelExistRequest) (response *CheckAttributeLabelExistResponse, err error) {
    return c.CheckAttributeLabelExistWithContext(context.Background(), request)
}

// CheckAttributeLabelExist
// 检查属性下的标签名是否存在
func (c *Client) CheckAttributeLabelExistWithContext(ctx context.Context, request *CheckAttributeLabelExistRequest) (response *CheckAttributeLabelExistResponse, err error) {
    if request == nil {
        request = NewCheckAttributeLabelExistRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckAttributeLabelExist require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckAttributeLabelExistResponse()
    err = c.Send(request, response)
    return
}

func NewCheckAttributeLabelReferRequest() (request *CheckAttributeLabelReferRequest) {
    request = &CheckAttributeLabelReferRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "CheckAttributeLabelRefer")
    
    
    return
}

func NewCheckAttributeLabelReferResponse() (response *CheckAttributeLabelReferResponse) {
    response = &CheckAttributeLabelReferResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckAttributeLabelRefer
// 检查属性标签引用
func (c *Client) CheckAttributeLabelRefer(request *CheckAttributeLabelReferRequest) (response *CheckAttributeLabelReferResponse, err error) {
    return c.CheckAttributeLabelReferWithContext(context.Background(), request)
}

// CheckAttributeLabelRefer
// 检查属性标签引用
func (c *Client) CheckAttributeLabelReferWithContext(ctx context.Context, request *CheckAttributeLabelReferRequest) (response *CheckAttributeLabelReferResponse, err error) {
    if request == nil {
        request = NewCheckAttributeLabelReferRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckAttributeLabelRefer require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckAttributeLabelReferResponse()
    err = c.Send(request, response)
    return
}

func NewConvertDocumentRequest() (request *ConvertDocumentRequest) {
    request = &ConvertDocumentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ConvertDocument")
    
    
    return
}

func NewConvertDocumentResponse() (response *ConvertDocumentResponse) {
    response = &ConvertDocumentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ConvertDocument
// 产品规划
//
// 
//
// 接口支持图片和PDF转可编辑word格式文件，将文件中的图片、文本、表格等元素识别，并根据位置进行还原。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_FILEDECODEFAILED = "FailedOperation.FileDecodeFailed"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) ConvertDocument(request *ConvertDocumentRequest) (response *ConvertDocumentResponse, err error) {
    return c.ConvertDocumentWithContext(context.Background(), request)
}

// ConvertDocument
// 产品规划
//
// 
//
// 接口支持图片和PDF转可编辑word格式文件，将文件中的图片、文本、表格等元素识别，并根据位置进行还原。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_FILEDECODEFAILED = "FailedOperation.FileDecodeFailed"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) ConvertDocumentWithContext(ctx context.Context, request *ConvertDocumentRequest) (response *ConvertDocumentResponse, err error) {
    if request == nil {
        request = NewConvertDocumentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ConvertDocument require credential")
    }

    request.SetContext(ctx)
    
    response = NewConvertDocumentResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAppRequest() (request *CreateAppRequest) {
    request = &CreateAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "CreateApp")
    
    
    return
}

func NewCreateAppResponse() (response *CreateAppResponse) {
    response = &CreateAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateApp
// 创建知识引擎应用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_FILEDECODEFAILED = "FailedOperation.FileDecodeFailed"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) CreateApp(request *CreateAppRequest) (response *CreateAppResponse, err error) {
    return c.CreateAppWithContext(context.Background(), request)
}

// CreateApp
// 创建知识引擎应用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_FILEDECODEFAILED = "FailedOperation.FileDecodeFailed"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) CreateAppWithContext(ctx context.Context, request *CreateAppRequest) (response *CreateAppResponse, err error) {
    if request == nil {
        request = NewCreateAppRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAppResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAttributeLabelRequest() (request *CreateAttributeLabelRequest) {
    request = &CreateAttributeLabelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "CreateAttributeLabel")
    
    
    return
}

func NewCreateAttributeLabelResponse() (response *CreateAttributeLabelResponse) {
    response = &CreateAttributeLabelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAttributeLabel
// 创建属性
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_FILEDECODEFAILED = "FailedOperation.FileDecodeFailed"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) CreateAttributeLabel(request *CreateAttributeLabelRequest) (response *CreateAttributeLabelResponse, err error) {
    return c.CreateAttributeLabelWithContext(context.Background(), request)
}

// CreateAttributeLabel
// 创建属性
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_FILEDECODEFAILED = "FailedOperation.FileDecodeFailed"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
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

func NewCreateCorpRequest() (request *CreateCorpRequest) {
    request = &CreateCorpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "CreateCorp")
    
    
    return
}

func NewCreateCorpResponse() (response *CreateCorpResponse) {
    response = &CreateCorpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCorp
// 创建企业
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_FILEDECODEFAILED = "FailedOperation.FileDecodeFailed"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) CreateCorp(request *CreateCorpRequest) (response *CreateCorpResponse, err error) {
    return c.CreateCorpWithContext(context.Background(), request)
}

// CreateCorp
// 创建企业
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_FILEDECODEFAILED = "FailedOperation.FileDecodeFailed"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) CreateCorpWithContext(ctx context.Context, request *CreateCorpRequest) (response *CreateCorpResponse, err error) {
    if request == nil {
        request = NewCreateCorpRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCorp require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCorpResponse()
    err = c.Send(request, response)
    return
}

func NewCreateQARequest() (request *CreateQARequest) {
    request = &CreateQARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "CreateQA")
    
    
    return
}

func NewCreateQAResponse() (response *CreateQAResponse) {
    response = &CreateQAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateQA
// 录入问答
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateQA(request *CreateQARequest) (response *CreateQAResponse, err error) {
    return c.CreateQAWithContext(context.Background(), request)
}

// CreateQA
// 录入问答
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
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

func NewCreateQACateRequest() (request *CreateQACateRequest) {
    request = &CreateQACateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "CreateQACate")
    
    
    return
}

func NewCreateQACateResponse() (response *CreateQACateResponse) {
    response = &CreateQACateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateQACate
// 创建QA分类
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateQACate(request *CreateQACateRequest) (response *CreateQACateResponse, err error) {
    return c.CreateQACateWithContext(context.Background(), request)
}

// CreateQACate
// 创建QA分类
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateQACateWithContext(ctx context.Context, request *CreateQACateRequest) (response *CreateQACateResponse, err error) {
    if request == nil {
        request = NewCreateQACateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateQACate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateQACateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateReconstructDocumentFlowRequest() (request *CreateReconstructDocumentFlowRequest) {
    request = &CreateReconstructDocumentFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "CreateReconstructDocumentFlow")
    
    
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

func NewCreateRejectedQuestionRequest() (request *CreateRejectedQuestionRequest) {
    request = &CreateRejectedQuestionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "CreateRejectedQuestion")
    
    
    return
}

func NewCreateRejectedQuestionResponse() (response *CreateRejectedQuestionResponse) {
    response = &CreateRejectedQuestionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRejectedQuestion
// 创建拒答问题
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
func (c *Client) CreateRejectedQuestion(request *CreateRejectedQuestionRequest) (response *CreateRejectedQuestionResponse, err error) {
    return c.CreateRejectedQuestionWithContext(context.Background(), request)
}

// CreateRejectedQuestion
// 创建拒答问题
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
func (c *Client) CreateRejectedQuestionWithContext(ctx context.Context, request *CreateRejectedQuestionRequest) (response *CreateRejectedQuestionResponse, err error) {
    if request == nil {
        request = NewCreateRejectedQuestionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRejectedQuestion require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRejectedQuestionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateReleaseRequest() (request *CreateReleaseRequest) {
    request = &CreateReleaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "CreateRelease")
    
    
    return
}

func NewCreateReleaseResponse() (response *CreateReleaseResponse) {
    response = &CreateReleaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRelease
// 创建发布
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateRelease(request *CreateReleaseRequest) (response *CreateReleaseResponse, err error) {
    return c.CreateReleaseWithContext(context.Background(), request)
}

// CreateRelease
// 创建发布
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateReleaseWithContext(ctx context.Context, request *CreateReleaseRequest) (response *CreateReleaseResponse, err error) {
    if request == nil {
        request = NewCreateReleaseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRelease require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateReleaseResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAppRequest() (request *DeleteAppRequest) {
    request = &DeleteAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DeleteApp")
    
    
    return
}

func NewDeleteAppResponse() (response *DeleteAppResponse) {
    response = &DeleteAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteApp
// 删除应用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteApp(request *DeleteAppRequest) (response *DeleteAppResponse, err error) {
    return c.DeleteAppWithContext(context.Background(), request)
}

// DeleteApp
// 删除应用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteAppWithContext(ctx context.Context, request *DeleteAppRequest) (response *DeleteAppResponse, err error) {
    if request == nil {
        request = NewDeleteAppRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAppResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAttributeLabelRequest() (request *DeleteAttributeLabelRequest) {
    request = &DeleteAttributeLabelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DeleteAttributeLabel")
    
    
    return
}

func NewDeleteAttributeLabelResponse() (response *DeleteAttributeLabelResponse) {
    response = &DeleteAttributeLabelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAttributeLabel
// 删除属性标签
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteAttributeLabel(request *DeleteAttributeLabelRequest) (response *DeleteAttributeLabelResponse, err error) {
    return c.DeleteAttributeLabelWithContext(context.Background(), request)
}

// DeleteAttributeLabel
// 删除属性标签
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteAttributeLabelWithContext(ctx context.Context, request *DeleteAttributeLabelRequest) (response *DeleteAttributeLabelResponse, err error) {
    if request == nil {
        request = NewDeleteAttributeLabelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAttributeLabel require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAttributeLabelResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDocRequest() (request *DeleteDocRequest) {
    request = &DeleteDocRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DeleteDoc")
    
    
    return
}

func NewDeleteDocResponse() (response *DeleteDocResponse) {
    response = &DeleteDocResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDoc
// 删除文档
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteDoc(request *DeleteDocRequest) (response *DeleteDocResponse, err error) {
    return c.DeleteDocWithContext(context.Background(), request)
}

// DeleteDoc
// 删除文档
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteDocWithContext(ctx context.Context, request *DeleteDocRequest) (response *DeleteDocResponse, err error) {
    if request == nil {
        request = NewDeleteDocRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDoc require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDocResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteQARequest() (request *DeleteQARequest) {
    request = &DeleteQARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DeleteQA")
    
    
    return
}

func NewDeleteQAResponse() (response *DeleteQAResponse) {
    response = &DeleteQAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteQA
// 删除问答
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteQA(request *DeleteQARequest) (response *DeleteQAResponse, err error) {
    return c.DeleteQAWithContext(context.Background(), request)
}

// DeleteQA
// 删除问答
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteQAWithContext(ctx context.Context, request *DeleteQARequest) (response *DeleteQAResponse, err error) {
    if request == nil {
        request = NewDeleteQARequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteQA require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteQAResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteQACateRequest() (request *DeleteQACateRequest) {
    request = &DeleteQACateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DeleteQACate")
    
    
    return
}

func NewDeleteQACateResponse() (response *DeleteQACateResponse) {
    response = &DeleteQACateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteQACate
// 分类删除
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteQACate(request *DeleteQACateRequest) (response *DeleteQACateResponse, err error) {
    return c.DeleteQACateWithContext(context.Background(), request)
}

// DeleteQACate
// 分类删除
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteQACateWithContext(ctx context.Context, request *DeleteQACateRequest) (response *DeleteQACateResponse, err error) {
    if request == nil {
        request = NewDeleteQACateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteQACate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteQACateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRejectedQuestionRequest() (request *DeleteRejectedQuestionRequest) {
    request = &DeleteRejectedQuestionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DeleteRejectedQuestion")
    
    
    return
}

func NewDeleteRejectedQuestionResponse() (response *DeleteRejectedQuestionResponse) {
    response = &DeleteRejectedQuestionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRejectedQuestion
// 删除拒答问题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteRejectedQuestion(request *DeleteRejectedQuestionRequest) (response *DeleteRejectedQuestionResponse, err error) {
    return c.DeleteRejectedQuestionWithContext(context.Background(), request)
}

// DeleteRejectedQuestion
// 删除拒答问题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteRejectedQuestionWithContext(ctx context.Context, request *DeleteRejectedQuestionRequest) (response *DeleteRejectedQuestionResponse, err error) {
    if request == nil {
        request = NewDeleteRejectedQuestionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRejectedQuestion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRejectedQuestionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAppRequest() (request *DescribeAppRequest) {
    request = &DescribeAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeApp")
    
    
    return
}

func NewDescribeAppResponse() (response *DescribeAppResponse) {
    response = &DescribeAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApp
// 获取企业下应用详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeApp(request *DescribeAppRequest) (response *DescribeAppResponse, err error) {
    return c.DescribeAppWithContext(context.Background(), request)
}

// DescribeApp
// 获取企业下应用详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeAppWithContext(ctx context.Context, request *DescribeAppRequest) (response *DescribeAppResponse, err error) {
    if request == nil {
        request = NewDescribeAppRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAppResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAttributeLabelRequest() (request *DescribeAttributeLabelRequest) {
    request = &DescribeAttributeLabelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeAttributeLabel")
    
    
    return
}

func NewDescribeAttributeLabelResponse() (response *DescribeAttributeLabelResponse) {
    response = &DescribeAttributeLabelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAttributeLabel
// 查询属性标签详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeAttributeLabel(request *DescribeAttributeLabelRequest) (response *DescribeAttributeLabelResponse, err error) {
    return c.DescribeAttributeLabelWithContext(context.Background(), request)
}

// DescribeAttributeLabel
// 查询属性标签详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeAttributeLabelWithContext(ctx context.Context, request *DescribeAttributeLabelRequest) (response *DescribeAttributeLabelResponse, err error) {
    if request == nil {
        request = NewDescribeAttributeLabelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAttributeLabel require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAttributeLabelResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCallStatsGraphRequest() (request *DescribeCallStatsGraphRequest) {
    request = &DescribeCallStatsGraphRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeCallStatsGraph")
    
    
    return
}

func NewDescribeCallStatsGraphResponse() (response *DescribeCallStatsGraphResponse) {
    response = &DescribeCallStatsGraphResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCallStatsGraph
// 接口调用折线图
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeCallStatsGraph(request *DescribeCallStatsGraphRequest) (response *DescribeCallStatsGraphResponse, err error) {
    return c.DescribeCallStatsGraphWithContext(context.Background(), request)
}

// DescribeCallStatsGraph
// 接口调用折线图
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeCallStatsGraphWithContext(ctx context.Context, request *DescribeCallStatsGraphRequest) (response *DescribeCallStatsGraphResponse, err error) {
    if request == nil {
        request = NewDescribeCallStatsGraphRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCallStatsGraph require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCallStatsGraphResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConcurrencyUsageRequest() (request *DescribeConcurrencyUsageRequest) {
    request = &DescribeConcurrencyUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeConcurrencyUsage")
    
    
    return
}

func NewDescribeConcurrencyUsageResponse() (response *DescribeConcurrencyUsageResponse) {
    response = &DescribeConcurrencyUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConcurrencyUsage
// 并发调用响应
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeConcurrencyUsage(request *DescribeConcurrencyUsageRequest) (response *DescribeConcurrencyUsageResponse, err error) {
    return c.DescribeConcurrencyUsageWithContext(context.Background(), request)
}

// DescribeConcurrencyUsage
// 并发调用响应
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeConcurrencyUsageWithContext(ctx context.Context, request *DescribeConcurrencyUsageRequest) (response *DescribeConcurrencyUsageResponse, err error) {
    if request == nil {
        request = NewDescribeConcurrencyUsageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConcurrencyUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConcurrencyUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConcurrencyUsageGraphRequest() (request *DescribeConcurrencyUsageGraphRequest) {
    request = &DescribeConcurrencyUsageGraphRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeConcurrencyUsageGraph")
    
    
    return
}

func NewDescribeConcurrencyUsageGraphResponse() (response *DescribeConcurrencyUsageGraphResponse) {
    response = &DescribeConcurrencyUsageGraphResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConcurrencyUsageGraph
// 并发调用折线图
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeConcurrencyUsageGraph(request *DescribeConcurrencyUsageGraphRequest) (response *DescribeConcurrencyUsageGraphResponse, err error) {
    return c.DescribeConcurrencyUsageGraphWithContext(context.Background(), request)
}

// DescribeConcurrencyUsageGraph
// 并发调用折线图
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeConcurrencyUsageGraphWithContext(ctx context.Context, request *DescribeConcurrencyUsageGraphRequest) (response *DescribeConcurrencyUsageGraphResponse, err error) {
    if request == nil {
        request = NewDescribeConcurrencyUsageGraphRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConcurrencyUsageGraph require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConcurrencyUsageGraphResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCorpRequest() (request *DescribeCorpRequest) {
    request = &DescribeCorpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeCorp")
    
    
    return
}

func NewDescribeCorpResponse() (response *DescribeCorpResponse) {
    response = &DescribeCorpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCorp
// 企业详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeCorp(request *DescribeCorpRequest) (response *DescribeCorpResponse, err error) {
    return c.DescribeCorpWithContext(context.Background(), request)
}

// DescribeCorp
// 企业详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeCorpWithContext(ctx context.Context, request *DescribeCorpRequest) (response *DescribeCorpResponse, err error) {
    if request == nil {
        request = NewDescribeCorpRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCorp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCorpResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDocRequest() (request *DescribeDocRequest) {
    request = &DescribeDocRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeDoc")
    
    
    return
}

func NewDescribeDocResponse() (response *DescribeDocResponse) {
    response = &DescribeDocResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDoc
// 文档详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDoc(request *DescribeDocRequest) (response *DescribeDocResponse, err error) {
    return c.DescribeDocWithContext(context.Background(), request)
}

// DescribeDoc
// 文档详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
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

func NewDescribeKnowledgeUsageRequest() (request *DescribeKnowledgeUsageRequest) {
    request = &DescribeKnowledgeUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeKnowledgeUsage")
    
    
    return
}

func NewDescribeKnowledgeUsageResponse() (response *DescribeKnowledgeUsageResponse) {
    response = &DescribeKnowledgeUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeKnowledgeUsage
// 查询知识库用量
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeKnowledgeUsage(request *DescribeKnowledgeUsageRequest) (response *DescribeKnowledgeUsageResponse, err error) {
    return c.DescribeKnowledgeUsageWithContext(context.Background(), request)
}

// DescribeKnowledgeUsage
// 查询知识库用量
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeKnowledgeUsageWithContext(ctx context.Context, request *DescribeKnowledgeUsageRequest) (response *DescribeKnowledgeUsageResponse, err error) {
    if request == nil {
        request = NewDescribeKnowledgeUsageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKnowledgeUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKnowledgeUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKnowledgeUsagePieGraphRequest() (request *DescribeKnowledgeUsagePieGraphRequest) {
    request = &DescribeKnowledgeUsagePieGraphRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeKnowledgeUsagePieGraph")
    
    
    return
}

func NewDescribeKnowledgeUsagePieGraphResponse() (response *DescribeKnowledgeUsagePieGraphResponse) {
    response = &DescribeKnowledgeUsagePieGraphResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeKnowledgeUsagePieGraph
// 查询企业知识库容量饼图
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeKnowledgeUsagePieGraph(request *DescribeKnowledgeUsagePieGraphRequest) (response *DescribeKnowledgeUsagePieGraphResponse, err error) {
    return c.DescribeKnowledgeUsagePieGraphWithContext(context.Background(), request)
}

// DescribeKnowledgeUsagePieGraph
// 查询企业知识库容量饼图
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeKnowledgeUsagePieGraphWithContext(ctx context.Context, request *DescribeKnowledgeUsagePieGraphRequest) (response *DescribeKnowledgeUsagePieGraphResponse, err error) {
    if request == nil {
        request = NewDescribeKnowledgeUsagePieGraphRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKnowledgeUsagePieGraph require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKnowledgeUsagePieGraphResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeQARequest() (request *DescribeQARequest) {
    request = &DescribeQARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeQA")
    
    
    return
}

func NewDescribeQAResponse() (response *DescribeQAResponse) {
    response = &DescribeQAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeQA
// 问答详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeQA(request *DescribeQARequest) (response *DescribeQAResponse, err error) {
    return c.DescribeQAWithContext(context.Background(), request)
}

// DescribeQA
// 问答详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeQAWithContext(ctx context.Context, request *DescribeQARequest) (response *DescribeQAResponse, err error) {
    if request == nil {
        request = NewDescribeQARequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeQA require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeQAResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReferRequest() (request *DescribeReferRequest) {
    request = &DescribeReferRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeRefer")
    
    
    return
}

func NewDescribeReferResponse() (response *DescribeReferResponse) {
    response = &DescribeReferResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRefer
// 获取来源详情列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeRefer(request *DescribeReferRequest) (response *DescribeReferResponse, err error) {
    return c.DescribeReferWithContext(context.Background(), request)
}

// DescribeRefer
// 获取来源详情列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeReferWithContext(ctx context.Context, request *DescribeReferRequest) (response *DescribeReferResponse, err error) {
    if request == nil {
        request = NewDescribeReferRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRefer require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReferResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReleaseRequest() (request *DescribeReleaseRequest) {
    request = &DescribeReleaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeRelease")
    
    
    return
}

func NewDescribeReleaseResponse() (response *DescribeReleaseResponse) {
    response = &DescribeReleaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRelease
// 发布详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeRelease(request *DescribeReleaseRequest) (response *DescribeReleaseResponse, err error) {
    return c.DescribeReleaseWithContext(context.Background(), request)
}

// DescribeRelease
// 发布详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeReleaseWithContext(ctx context.Context, request *DescribeReleaseRequest) (response *DescribeReleaseResponse, err error) {
    if request == nil {
        request = NewDescribeReleaseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRelease require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReleaseResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReleaseInfoRequest() (request *DescribeReleaseInfoRequest) {
    request = &DescribeReleaseInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeReleaseInfo")
    
    
    return
}

func NewDescribeReleaseInfoResponse() (response *DescribeReleaseInfoResponse) {
    response = &DescribeReleaseInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeReleaseInfo
// 拉取发布按钮状态、最后发布时间
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeReleaseInfo(request *DescribeReleaseInfoRequest) (response *DescribeReleaseInfoResponse, err error) {
    return c.DescribeReleaseInfoWithContext(context.Background(), request)
}

// DescribeReleaseInfo
// 拉取发布按钮状态、最后发布时间
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeReleaseInfoWithContext(ctx context.Context, request *DescribeReleaseInfoRequest) (response *DescribeReleaseInfoResponse, err error) {
    if request == nil {
        request = NewDescribeReleaseInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReleaseInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReleaseInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRobotBizIDByAppKeyRequest() (request *DescribeRobotBizIDByAppKeyRequest) {
    request = &DescribeRobotBizIDByAppKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeRobotBizIDByAppKey")
    
    
    return
}

func NewDescribeRobotBizIDByAppKeyResponse() (response *DescribeRobotBizIDByAppKeyResponse) {
    response = &DescribeRobotBizIDByAppKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRobotBizIDByAppKey
// 通过appKey获取应用业务ID
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeRobotBizIDByAppKey(request *DescribeRobotBizIDByAppKeyRequest) (response *DescribeRobotBizIDByAppKeyResponse, err error) {
    return c.DescribeRobotBizIDByAppKeyWithContext(context.Background(), request)
}

// DescribeRobotBizIDByAppKey
// 通过appKey获取应用业务ID
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeRobotBizIDByAppKeyWithContext(ctx context.Context, request *DescribeRobotBizIDByAppKeyRequest) (response *DescribeRobotBizIDByAppKeyResponse, err error) {
    if request == nil {
        request = NewDescribeRobotBizIDByAppKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRobotBizIDByAppKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRobotBizIDByAppKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSearchStatsGraphRequest() (request *DescribeSearchStatsGraphRequest) {
    request = &DescribeSearchStatsGraphRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeSearchStatsGraph")
    
    
    return
}

func NewDescribeSearchStatsGraphResponse() (response *DescribeSearchStatsGraphResponse) {
    response = &DescribeSearchStatsGraphResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSearchStatsGraph
// 查询搜索服务调用折线图
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeSearchStatsGraph(request *DescribeSearchStatsGraphRequest) (response *DescribeSearchStatsGraphResponse, err error) {
    return c.DescribeSearchStatsGraphWithContext(context.Background(), request)
}

// DescribeSearchStatsGraph
// 查询搜索服务调用折线图
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeSearchStatsGraphWithContext(ctx context.Context, request *DescribeSearchStatsGraphRequest) (response *DescribeSearchStatsGraphResponse, err error) {
    if request == nil {
        request = NewDescribeSearchStatsGraphRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSearchStatsGraph require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSearchStatsGraphResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSegmentsRequest() (request *DescribeSegmentsRequest) {
    request = &DescribeSegmentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeSegments")
    
    
    return
}

func NewDescribeSegmentsResponse() (response *DescribeSegmentsResponse) {
    response = &DescribeSegmentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSegments
// 获取片段详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeSegments(request *DescribeSegmentsRequest) (response *DescribeSegmentsResponse, err error) {
    return c.DescribeSegmentsWithContext(context.Background(), request)
}

// DescribeSegments
// 获取片段详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeSegmentsWithContext(ctx context.Context, request *DescribeSegmentsRequest) (response *DescribeSegmentsResponse, err error) {
    if request == nil {
        request = NewDescribeSegmentsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSegments require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSegmentsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStorageCredentialRequest() (request *DescribeStorageCredentialRequest) {
    request = &DescribeStorageCredentialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeStorageCredential")
    
    
    return
}

func NewDescribeStorageCredentialResponse() (response *DescribeStorageCredentialResponse) {
    response = &DescribeStorageCredentialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStorageCredential
// 获取文件上传临时密钥
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeStorageCredential(request *DescribeStorageCredentialRequest) (response *DescribeStorageCredentialResponse, err error) {
    return c.DescribeStorageCredentialWithContext(context.Background(), request)
}

// DescribeStorageCredential
// 获取文件上传临时密钥
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeStorageCredentialWithContext(ctx context.Context, request *DescribeStorageCredentialRequest) (response *DescribeStorageCredentialResponse, err error) {
    if request == nil {
        request = NewDescribeStorageCredentialRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStorageCredential require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStorageCredentialResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTokenUsageRequest() (request *DescribeTokenUsageRequest) {
    request = &DescribeTokenUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeTokenUsage")
    
    
    return
}

func NewDescribeTokenUsageResponse() (response *DescribeTokenUsageResponse) {
    response = &DescribeTokenUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTokenUsage
// 接口调用token详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTokenUsage(request *DescribeTokenUsageRequest) (response *DescribeTokenUsageResponse, err error) {
    return c.DescribeTokenUsageWithContext(context.Background(), request)
}

// DescribeTokenUsage
// 接口调用token详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTokenUsageWithContext(ctx context.Context, request *DescribeTokenUsageRequest) (response *DescribeTokenUsageResponse, err error) {
    if request == nil {
        request = NewDescribeTokenUsageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTokenUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTokenUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTokenUsageGraphRequest() (request *DescribeTokenUsageGraphRequest) {
    request = &DescribeTokenUsageGraphRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeTokenUsageGraph")
    
    
    return
}

func NewDescribeTokenUsageGraphResponse() (response *DescribeTokenUsageGraphResponse) {
    response = &DescribeTokenUsageGraphResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTokenUsageGraph
// 接口调用token折线图
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTokenUsageGraph(request *DescribeTokenUsageGraphRequest) (response *DescribeTokenUsageGraphResponse, err error) {
    return c.DescribeTokenUsageGraphWithContext(context.Background(), request)
}

// DescribeTokenUsageGraph
// 接口调用token折线图
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTokenUsageGraphWithContext(ctx context.Context, request *DescribeTokenUsageGraphRequest) (response *DescribeTokenUsageGraphResponse, err error) {
    if request == nil {
        request = NewDescribeTokenUsageGraphRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTokenUsageGraph require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTokenUsageGraphResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUnsatisfiedReplyContextRequest() (request *DescribeUnsatisfiedReplyContextRequest) {
    request = &DescribeUnsatisfiedReplyContextRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeUnsatisfiedReplyContext")
    
    
    return
}

func NewDescribeUnsatisfiedReplyContextResponse() (response *DescribeUnsatisfiedReplyContextResponse) {
    response = &DescribeUnsatisfiedReplyContextResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUnsatisfiedReplyContext
// 获取不满意回复上下文
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeUnsatisfiedReplyContext(request *DescribeUnsatisfiedReplyContextRequest) (response *DescribeUnsatisfiedReplyContextResponse, err error) {
    return c.DescribeUnsatisfiedReplyContextWithContext(context.Background(), request)
}

// DescribeUnsatisfiedReplyContext
// 获取不满意回复上下文
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeUnsatisfiedReplyContextWithContext(ctx context.Context, request *DescribeUnsatisfiedReplyContextRequest) (response *DescribeUnsatisfiedReplyContextResponse, err error) {
    if request == nil {
        request = NewDescribeUnsatisfiedReplyContextRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUnsatisfiedReplyContext require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUnsatisfiedReplyContextResponse()
    err = c.Send(request, response)
    return
}

func NewExportAttributeLabelRequest() (request *ExportAttributeLabelRequest) {
    request = &ExportAttributeLabelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ExportAttributeLabel")
    
    
    return
}

func NewExportAttributeLabelResponse() (response *ExportAttributeLabelResponse) {
    response = &ExportAttributeLabelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ExportAttributeLabel
// 导出属性标签
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ExportAttributeLabel(request *ExportAttributeLabelRequest) (response *ExportAttributeLabelResponse, err error) {
    return c.ExportAttributeLabelWithContext(context.Background(), request)
}

// ExportAttributeLabel
// 导出属性标签
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ExportAttributeLabelWithContext(ctx context.Context, request *ExportAttributeLabelRequest) (response *ExportAttributeLabelResponse, err error) {
    if request == nil {
        request = NewExportAttributeLabelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportAttributeLabel require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportAttributeLabelResponse()
    err = c.Send(request, response)
    return
}

func NewExportQAListRequest() (request *ExportQAListRequest) {
    request = &ExportQAListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ExportQAList")
    
    
    return
}

func NewExportQAListResponse() (response *ExportQAListResponse) {
    response = &ExportQAListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ExportQAList
// 导出QA列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ExportQAList(request *ExportQAListRequest) (response *ExportQAListResponse, err error) {
    return c.ExportQAListWithContext(context.Background(), request)
}

// ExportQAList
// 导出QA列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ExportQAListWithContext(ctx context.Context, request *ExportQAListRequest) (response *ExportQAListResponse, err error) {
    if request == nil {
        request = NewExportQAListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportQAList require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportQAListResponse()
    err = c.Send(request, response)
    return
}

func NewExportUnsatisfiedReplyRequest() (request *ExportUnsatisfiedReplyRequest) {
    request = &ExportUnsatisfiedReplyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ExportUnsatisfiedReply")
    
    
    return
}

func NewExportUnsatisfiedReplyResponse() (response *ExportUnsatisfiedReplyResponse) {
    response = &ExportUnsatisfiedReplyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ExportUnsatisfiedReply
// 导出不满意回复
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ExportUnsatisfiedReply(request *ExportUnsatisfiedReplyRequest) (response *ExportUnsatisfiedReplyResponse, err error) {
    return c.ExportUnsatisfiedReplyWithContext(context.Background(), request)
}

// ExportUnsatisfiedReply
// 导出不满意回复
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ExportUnsatisfiedReplyWithContext(ctx context.Context, request *ExportUnsatisfiedReplyRequest) (response *ExportUnsatisfiedReplyResponse, err error) {
    if request == nil {
        request = NewExportUnsatisfiedReplyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportUnsatisfiedReply require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportUnsatisfiedReplyResponse()
    err = c.Send(request, response)
    return
}

func NewGenerateQARequest() (request *GenerateQARequest) {
    request = &GenerateQARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "GenerateQA")
    
    
    return
}

func NewGenerateQAResponse() (response *GenerateQAResponse) {
    response = &GenerateQAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GenerateQA
// 文档生成问答
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GenerateQA(request *GenerateQARequest) (response *GenerateQAResponse, err error) {
    return c.GenerateQAWithContext(context.Background(), request)
}

// GenerateQA
// 文档生成问答
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GenerateQAWithContext(ctx context.Context, request *GenerateQARequest) (response *GenerateQAResponse, err error) {
    if request == nil {
        request = NewGenerateQARequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GenerateQA require credential")
    }

    request.SetContext(ctx)
    
    response = NewGenerateQAResponse()
    err = c.Send(request, response)
    return
}

func NewGetAnswerTypeDataCountRequest() (request *GetAnswerTypeDataCountRequest) {
    request = &GetAnswerTypeDataCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "GetAnswerTypeDataCount")
    
    
    return
}

func NewGetAnswerTypeDataCountResponse() (response *GetAnswerTypeDataCountResponse) {
    response = &GetAnswerTypeDataCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetAnswerTypeDataCount
// 回答类型数据统计
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetAnswerTypeDataCount(request *GetAnswerTypeDataCountRequest) (response *GetAnswerTypeDataCountResponse, err error) {
    return c.GetAnswerTypeDataCountWithContext(context.Background(), request)
}

// GetAnswerTypeDataCount
// 回答类型数据统计
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetAnswerTypeDataCountWithContext(ctx context.Context, request *GetAnswerTypeDataCountRequest) (response *GetAnswerTypeDataCountResponse, err error) {
    if request == nil {
        request = NewGetAnswerTypeDataCountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetAnswerTypeDataCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetAnswerTypeDataCountResponse()
    err = c.Send(request, response)
    return
}

func NewGetAppKnowledgeCountRequest() (request *GetAppKnowledgeCountRequest) {
    request = &GetAppKnowledgeCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "GetAppKnowledgeCount")
    
    
    return
}

func NewGetAppKnowledgeCountResponse() (response *GetAppKnowledgeCountResponse) {
    response = &GetAppKnowledgeCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetAppKnowledgeCount
// 获取模型列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetAppKnowledgeCount(request *GetAppKnowledgeCountRequest) (response *GetAppKnowledgeCountResponse, err error) {
    return c.GetAppKnowledgeCountWithContext(context.Background(), request)
}

// GetAppKnowledgeCount
// 获取模型列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetAppKnowledgeCountWithContext(ctx context.Context, request *GetAppKnowledgeCountRequest) (response *GetAppKnowledgeCountResponse, err error) {
    if request == nil {
        request = NewGetAppKnowledgeCountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetAppKnowledgeCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetAppKnowledgeCountResponse()
    err = c.Send(request, response)
    return
}

func NewGetAppSecretRequest() (request *GetAppSecretRequest) {
    request = &GetAppSecretRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "GetAppSecret")
    
    
    return
}

func NewGetAppSecretResponse() (response *GetAppSecretResponse) {
    response = &GetAppSecretResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetAppSecret
// 获取应用密钥
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetAppSecret(request *GetAppSecretRequest) (response *GetAppSecretResponse, err error) {
    return c.GetAppSecretWithContext(context.Background(), request)
}

// GetAppSecret
// 获取应用密钥
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetAppSecretWithContext(ctx context.Context, request *GetAppSecretRequest) (response *GetAppSecretResponse, err error) {
    if request == nil {
        request = NewGetAppSecretRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetAppSecret require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetAppSecretResponse()
    err = c.Send(request, response)
    return
}

func NewGetDocPreviewRequest() (request *GetDocPreviewRequest) {
    request = &GetDocPreviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "GetDocPreview")
    
    
    return
}

func NewGetDocPreviewResponse() (response *GetDocPreviewResponse) {
    response = &GetDocPreviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetDocPreview
// 获取文档预览信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetDocPreview(request *GetDocPreviewRequest) (response *GetDocPreviewResponse, err error) {
    return c.GetDocPreviewWithContext(context.Background(), request)
}

// GetDocPreview
// 获取文档预览信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetDocPreviewWithContext(ctx context.Context, request *GetDocPreviewRequest) (response *GetDocPreviewResponse, err error) {
    if request == nil {
        request = NewGetDocPreviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDocPreview require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDocPreviewResponse()
    err = c.Send(request, response)
    return
}

func NewGetEmbeddingRequest() (request *GetEmbeddingRequest) {
    request = &GetEmbeddingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "GetEmbedding")
    
    
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

func NewGetLikeDataCountRequest() (request *GetLikeDataCountRequest) {
    request = &GetLikeDataCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "GetLikeDataCount")
    
    
    return
}

func NewGetLikeDataCountResponse() (response *GetLikeDataCountResponse) {
    response = &GetLikeDataCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetLikeDataCount
// 点赞点踩数据统计
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetLikeDataCount(request *GetLikeDataCountRequest) (response *GetLikeDataCountResponse, err error) {
    return c.GetLikeDataCountWithContext(context.Background(), request)
}

// GetLikeDataCount
// 点赞点踩数据统计
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetLikeDataCountWithContext(ctx context.Context, request *GetLikeDataCountRequest) (response *GetLikeDataCountResponse, err error) {
    if request == nil {
        request = NewGetLikeDataCountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetLikeDataCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetLikeDataCountResponse()
    err = c.Send(request, response)
    return
}

func NewGetMsgRecordRequest() (request *GetMsgRecordRequest) {
    request = &GetMsgRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "GetMsgRecord")
    
    
    return
}

func NewGetMsgRecordResponse() (response *GetMsgRecordResponse) {
    response = &GetMsgRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetMsgRecord
// 获取聊天历史请求
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetMsgRecord(request *GetMsgRecordRequest) (response *GetMsgRecordResponse, err error) {
    return c.GetMsgRecordWithContext(context.Background(), request)
}

// GetMsgRecord
// 获取聊天历史请求
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetMsgRecordWithContext(ctx context.Context, request *GetMsgRecordRequest) (response *GetMsgRecordResponse, err error) {
    if request == nil {
        request = NewGetMsgRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetMsgRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetMsgRecordResponse()
    err = c.Send(request, response)
    return
}

func NewGetReconstructDocumentResultRequest() (request *GetReconstructDocumentResultRequest) {
    request = &GetReconstructDocumentResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "GetReconstructDocumentResult")
    
    
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

func NewGetTaskStatusRequest() (request *GetTaskStatusRequest) {
    request = &GetTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "GetTaskStatus")
    
    
    return
}

func NewGetTaskStatusResponse() (response *GetTaskStatusResponse) {
    response = &GetTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetTaskStatus
// 获取任务状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) GetTaskStatus(request *GetTaskStatusRequest) (response *GetTaskStatusResponse, err error) {
    return c.GetTaskStatusWithContext(context.Background(), request)
}

// GetTaskStatus
// 获取任务状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) GetTaskStatusWithContext(ctx context.Context, request *GetTaskStatusRequest) (response *GetTaskStatusResponse, err error) {
    if request == nil {
        request = NewGetTaskStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTaskStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewGetWsTokenRequest() (request *GetWsTokenRequest) {
    request = &GetWsTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "GetWsToken")
    
    
    return
}

func NewGetWsTokenResponse() (response *GetWsTokenResponse) {
    response = &GetWsTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetWsToken
// 获取ws token
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetWsToken(request *GetWsTokenRequest) (response *GetWsTokenResponse, err error) {
    return c.GetWsTokenWithContext(context.Background(), request)
}

// GetWsToken
// 获取ws token
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetWsTokenWithContext(ctx context.Context, request *GetWsTokenRequest) (response *GetWsTokenResponse, err error) {
    if request == nil {
        request = NewGetWsTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetWsToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetWsTokenResponse()
    err = c.Send(request, response)
    return
}

func NewGroupQARequest() (request *GroupQARequest) {
    request = &GroupQARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "GroupQA")
    
    
    return
}

func NewGroupQAResponse() (response *GroupQAResponse) {
    response = &GroupQAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GroupQA
// QA分组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GroupQA(request *GroupQARequest) (response *GroupQAResponse, err error) {
    return c.GroupQAWithContext(context.Background(), request)
}

// GroupQA
// QA分组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GroupQAWithContext(ctx context.Context, request *GroupQARequest) (response *GroupQAResponse, err error) {
    if request == nil {
        request = NewGroupQARequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GroupQA require credential")
    }

    request.SetContext(ctx)
    
    response = NewGroupQAResponse()
    err = c.Send(request, response)
    return
}

func NewIgnoreUnsatisfiedReplyRequest() (request *IgnoreUnsatisfiedReplyRequest) {
    request = &IgnoreUnsatisfiedReplyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "IgnoreUnsatisfiedReply")
    
    
    return
}

func NewIgnoreUnsatisfiedReplyResponse() (response *IgnoreUnsatisfiedReplyResponse) {
    response = &IgnoreUnsatisfiedReplyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// IgnoreUnsatisfiedReply
// 忽略不满意回复
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) IgnoreUnsatisfiedReply(request *IgnoreUnsatisfiedReplyRequest) (response *IgnoreUnsatisfiedReplyResponse, err error) {
    return c.IgnoreUnsatisfiedReplyWithContext(context.Background(), request)
}

// IgnoreUnsatisfiedReply
// 忽略不满意回复
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) IgnoreUnsatisfiedReplyWithContext(ctx context.Context, request *IgnoreUnsatisfiedReplyRequest) (response *IgnoreUnsatisfiedReplyResponse, err error) {
    if request == nil {
        request = NewIgnoreUnsatisfiedReplyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("IgnoreUnsatisfiedReply require credential")
    }

    request.SetContext(ctx)
    
    response = NewIgnoreUnsatisfiedReplyResponse()
    err = c.Send(request, response)
    return
}

func NewIsTransferIntentRequest() (request *IsTransferIntentRequest) {
    request = &IsTransferIntentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "IsTransferIntent")
    
    
    return
}

func NewIsTransferIntentResponse() (response *IsTransferIntentResponse) {
    response = &IsTransferIntentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// IsTransferIntent
// 是否意图转人工
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) IsTransferIntent(request *IsTransferIntentRequest) (response *IsTransferIntentResponse, err error) {
    return c.IsTransferIntentWithContext(context.Background(), request)
}

// IsTransferIntent
// 是否意图转人工
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) IsTransferIntentWithContext(ctx context.Context, request *IsTransferIntentRequest) (response *IsTransferIntentResponse, err error) {
    if request == nil {
        request = NewIsTransferIntentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("IsTransferIntent require credential")
    }

    request.SetContext(ctx)
    
    response = NewIsTransferIntentResponse()
    err = c.Send(request, response)
    return
}

func NewListAppRequest() (request *ListAppRequest) {
    request = &ListAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListApp")
    
    
    return
}

func NewListAppResponse() (response *ListAppResponse) {
    response = &ListAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListApp
// 获取企业下应用列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListApp(request *ListAppRequest) (response *ListAppResponse, err error) {
    return c.ListAppWithContext(context.Background(), request)
}

// ListApp
// 获取企业下应用列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListAppWithContext(ctx context.Context, request *ListAppRequest) (response *ListAppResponse, err error) {
    if request == nil {
        request = NewListAppRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAppResponse()
    err = c.Send(request, response)
    return
}

func NewListAppCategoryRequest() (request *ListAppCategoryRequest) {
    request = &ListAppCategoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListAppCategory")
    
    
    return
}

func NewListAppCategoryResponse() (response *ListAppCategoryResponse) {
    response = &ListAppCategoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListAppCategory
// 应用类型列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListAppCategory(request *ListAppCategoryRequest) (response *ListAppCategoryResponse, err error) {
    return c.ListAppCategoryWithContext(context.Background(), request)
}

// ListAppCategory
// 应用类型列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListAppCategoryWithContext(ctx context.Context, request *ListAppCategoryRequest) (response *ListAppCategoryResponse, err error) {
    if request == nil {
        request = NewListAppCategoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAppCategory require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAppCategoryResponse()
    err = c.Send(request, response)
    return
}

func NewListAttributeLabelRequest() (request *ListAttributeLabelRequest) {
    request = &ListAttributeLabelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListAttributeLabel")
    
    
    return
}

func NewListAttributeLabelResponse() (response *ListAttributeLabelResponse) {
    response = &ListAttributeLabelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListAttributeLabel
// 查询属性标签列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListAttributeLabel(request *ListAttributeLabelRequest) (response *ListAttributeLabelResponse, err error) {
    return c.ListAttributeLabelWithContext(context.Background(), request)
}

// ListAttributeLabel
// 查询属性标签列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListAttributeLabelWithContext(ctx context.Context, request *ListAttributeLabelRequest) (response *ListAttributeLabelResponse, err error) {
    if request == nil {
        request = NewListAttributeLabelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAttributeLabel require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAttributeLabelResponse()
    err = c.Send(request, response)
    return
}

func NewListDocRequest() (request *ListDocRequest) {
    request = &ListDocRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListDoc")
    
    
    return
}

func NewListDocResponse() (response *ListDocResponse) {
    response = &ListDocResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListDoc
// 文档列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListDoc(request *ListDocRequest) (response *ListDocResponse, err error) {
    return c.ListDocWithContext(context.Background(), request)
}

// ListDoc
// 文档列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListDocWithContext(ctx context.Context, request *ListDocRequest) (response *ListDocResponse, err error) {
    if request == nil {
        request = NewListDocRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListDoc require credential")
    }

    request.SetContext(ctx)
    
    response = NewListDocResponse()
    err = c.Send(request, response)
    return
}

func NewListModelRequest() (request *ListModelRequest) {
    request = &ListModelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListModel")
    
    
    return
}

func NewListModelResponse() (response *ListModelResponse) {
    response = &ListModelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListModel
// 获取模型列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListModel(request *ListModelRequest) (response *ListModelResponse, err error) {
    return c.ListModelWithContext(context.Background(), request)
}

// ListModel
// 获取模型列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListModelWithContext(ctx context.Context, request *ListModelRequest) (response *ListModelResponse, err error) {
    if request == nil {
        request = NewListModelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListModel require credential")
    }

    request.SetContext(ctx)
    
    response = NewListModelResponse()
    err = c.Send(request, response)
    return
}

func NewListQARequest() (request *ListQARequest) {
    request = &ListQARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListQA")
    
    
    return
}

func NewListQAResponse() (response *ListQAResponse) {
    response = &ListQAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListQA
// 问答列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListQA(request *ListQARequest) (response *ListQAResponse, err error) {
    return c.ListQAWithContext(context.Background(), request)
}

// ListQA
// 问答列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListQAWithContext(ctx context.Context, request *ListQARequest) (response *ListQAResponse, err error) {
    if request == nil {
        request = NewListQARequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListQA require credential")
    }

    request.SetContext(ctx)
    
    response = NewListQAResponse()
    err = c.Send(request, response)
    return
}

func NewListQACateRequest() (request *ListQACateRequest) {
    request = &ListQACateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListQACate")
    
    
    return
}

func NewListQACateResponse() (response *ListQACateResponse) {
    response = &ListQACateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListQACate
// 获取QA分类
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListQACate(request *ListQACateRequest) (response *ListQACateResponse, err error) {
    return c.ListQACateWithContext(context.Background(), request)
}

// ListQACate
// 获取QA分类
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListQACateWithContext(ctx context.Context, request *ListQACateRequest) (response *ListQACateResponse, err error) {
    if request == nil {
        request = NewListQACateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListQACate require credential")
    }

    request.SetContext(ctx)
    
    response = NewListQACateResponse()
    err = c.Send(request, response)
    return
}

func NewListRejectedQuestionRequest() (request *ListRejectedQuestionRequest) {
    request = &ListRejectedQuestionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListRejectedQuestion")
    
    
    return
}

func NewListRejectedQuestionResponse() (response *ListRejectedQuestionResponse) {
    response = &ListRejectedQuestionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListRejectedQuestion
// 获取拒答问题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListRejectedQuestion(request *ListRejectedQuestionRequest) (response *ListRejectedQuestionResponse, err error) {
    return c.ListRejectedQuestionWithContext(context.Background(), request)
}

// ListRejectedQuestion
// 获取拒答问题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListRejectedQuestionWithContext(ctx context.Context, request *ListRejectedQuestionRequest) (response *ListRejectedQuestionResponse, err error) {
    if request == nil {
        request = NewListRejectedQuestionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListRejectedQuestion require credential")
    }

    request.SetContext(ctx)
    
    response = NewListRejectedQuestionResponse()
    err = c.Send(request, response)
    return
}

func NewListRejectedQuestionPreviewRequest() (request *ListRejectedQuestionPreviewRequest) {
    request = &ListRejectedQuestionPreviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListRejectedQuestionPreview")
    
    
    return
}

func NewListRejectedQuestionPreviewResponse() (response *ListRejectedQuestionPreviewResponse) {
    response = &ListRejectedQuestionPreviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListRejectedQuestionPreview
// 发布拒答问题预览
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListRejectedQuestionPreview(request *ListRejectedQuestionPreviewRequest) (response *ListRejectedQuestionPreviewResponse, err error) {
    return c.ListRejectedQuestionPreviewWithContext(context.Background(), request)
}

// ListRejectedQuestionPreview
// 发布拒答问题预览
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListRejectedQuestionPreviewWithContext(ctx context.Context, request *ListRejectedQuestionPreviewRequest) (response *ListRejectedQuestionPreviewResponse, err error) {
    if request == nil {
        request = NewListRejectedQuestionPreviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListRejectedQuestionPreview require credential")
    }

    request.SetContext(ctx)
    
    response = NewListRejectedQuestionPreviewResponse()
    err = c.Send(request, response)
    return
}

func NewListReleaseRequest() (request *ListReleaseRequest) {
    request = &ListReleaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListRelease")
    
    
    return
}

func NewListReleaseResponse() (response *ListReleaseResponse) {
    response = &ListReleaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListRelease
// 发布列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListRelease(request *ListReleaseRequest) (response *ListReleaseResponse, err error) {
    return c.ListReleaseWithContext(context.Background(), request)
}

// ListRelease
// 发布列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListReleaseWithContext(ctx context.Context, request *ListReleaseRequest) (response *ListReleaseResponse, err error) {
    if request == nil {
        request = NewListReleaseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListRelease require credential")
    }

    request.SetContext(ctx)
    
    response = NewListReleaseResponse()
    err = c.Send(request, response)
    return
}

func NewListReleaseConfigPreviewRequest() (request *ListReleaseConfigPreviewRequest) {
    request = &ListReleaseConfigPreviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListReleaseConfigPreview")
    
    
    return
}

func NewListReleaseConfigPreviewResponse() (response *ListReleaseConfigPreviewResponse) {
    response = &ListReleaseConfigPreviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListReleaseConfigPreview
// 发布配置项预览
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListReleaseConfigPreview(request *ListReleaseConfigPreviewRequest) (response *ListReleaseConfigPreviewResponse, err error) {
    return c.ListReleaseConfigPreviewWithContext(context.Background(), request)
}

// ListReleaseConfigPreview
// 发布配置项预览
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListReleaseConfigPreviewWithContext(ctx context.Context, request *ListReleaseConfigPreviewRequest) (response *ListReleaseConfigPreviewResponse, err error) {
    if request == nil {
        request = NewListReleaseConfigPreviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListReleaseConfigPreview require credential")
    }

    request.SetContext(ctx)
    
    response = NewListReleaseConfigPreviewResponse()
    err = c.Send(request, response)
    return
}

func NewListReleaseDocPreviewRequest() (request *ListReleaseDocPreviewRequest) {
    request = &ListReleaseDocPreviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListReleaseDocPreview")
    
    
    return
}

func NewListReleaseDocPreviewResponse() (response *ListReleaseDocPreviewResponse) {
    response = &ListReleaseDocPreviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListReleaseDocPreview
// 发布文档预览
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListReleaseDocPreview(request *ListReleaseDocPreviewRequest) (response *ListReleaseDocPreviewResponse, err error) {
    return c.ListReleaseDocPreviewWithContext(context.Background(), request)
}

// ListReleaseDocPreview
// 发布文档预览
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListReleaseDocPreviewWithContext(ctx context.Context, request *ListReleaseDocPreviewRequest) (response *ListReleaseDocPreviewResponse, err error) {
    if request == nil {
        request = NewListReleaseDocPreviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListReleaseDocPreview require credential")
    }

    request.SetContext(ctx)
    
    response = NewListReleaseDocPreviewResponse()
    err = c.Send(request, response)
    return
}

func NewListReleaseQAPreviewRequest() (request *ListReleaseQAPreviewRequest) {
    request = &ListReleaseQAPreviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListReleaseQAPreview")
    
    
    return
}

func NewListReleaseQAPreviewResponse() (response *ListReleaseQAPreviewResponse) {
    response = &ListReleaseQAPreviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListReleaseQAPreview
// 文档列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListReleaseQAPreview(request *ListReleaseQAPreviewRequest) (response *ListReleaseQAPreviewResponse, err error) {
    return c.ListReleaseQAPreviewWithContext(context.Background(), request)
}

// ListReleaseQAPreview
// 文档列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListReleaseQAPreviewWithContext(ctx context.Context, request *ListReleaseQAPreviewRequest) (response *ListReleaseQAPreviewResponse, err error) {
    if request == nil {
        request = NewListReleaseQAPreviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListReleaseQAPreview require credential")
    }

    request.SetContext(ctx)
    
    response = NewListReleaseQAPreviewResponse()
    err = c.Send(request, response)
    return
}

func NewListSelectDocRequest() (request *ListSelectDocRequest) {
    request = &ListSelectDocRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListSelectDoc")
    
    
    return
}

func NewListSelectDocResponse() (response *ListSelectDocResponse) {
    response = &ListSelectDocResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListSelectDoc
// 获取账户信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListSelectDoc(request *ListSelectDocRequest) (response *ListSelectDocResponse, err error) {
    return c.ListSelectDocWithContext(context.Background(), request)
}

// ListSelectDoc
// 获取账户信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListSelectDocWithContext(ctx context.Context, request *ListSelectDocRequest) (response *ListSelectDocResponse, err error) {
    if request == nil {
        request = NewListSelectDocRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListSelectDoc require credential")
    }

    request.SetContext(ctx)
    
    response = NewListSelectDocResponse()
    err = c.Send(request, response)
    return
}

func NewListUnsatisfiedReplyRequest() (request *ListUnsatisfiedReplyRequest) {
    request = &ListUnsatisfiedReplyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListUnsatisfiedReply")
    
    
    return
}

func NewListUnsatisfiedReplyResponse() (response *ListUnsatisfiedReplyResponse) {
    response = &ListUnsatisfiedReplyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListUnsatisfiedReply
// 查询不满意回复列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListUnsatisfiedReply(request *ListUnsatisfiedReplyRequest) (response *ListUnsatisfiedReplyResponse, err error) {
    return c.ListUnsatisfiedReplyWithContext(context.Background(), request)
}

// ListUnsatisfiedReply
// 查询不满意回复列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListUnsatisfiedReplyWithContext(ctx context.Context, request *ListUnsatisfiedReplyRequest) (response *ListUnsatisfiedReplyResponse, err error) {
    if request == nil {
        request = NewListUnsatisfiedReplyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListUnsatisfiedReply require credential")
    }

    request.SetContext(ctx)
    
    response = NewListUnsatisfiedReplyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAppRequest() (request *ModifyAppRequest) {
    request = &ModifyAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ModifyApp")
    
    
    return
}

func NewModifyAppResponse() (response *ModifyAppResponse) {
    response = &ModifyAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApp
// 修改应用请求结构体
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyApp(request *ModifyAppRequest) (response *ModifyAppResponse, err error) {
    return c.ModifyAppWithContext(context.Background(), request)
}

// ModifyApp
// 修改应用请求结构体
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyAppWithContext(ctx context.Context, request *ModifyAppRequest) (response *ModifyAppResponse, err error) {
    if request == nil {
        request = NewModifyAppRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAppResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAttributeLabelRequest() (request *ModifyAttributeLabelRequest) {
    request = &ModifyAttributeLabelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ModifyAttributeLabel")
    
    
    return
}

func NewModifyAttributeLabelResponse() (response *ModifyAttributeLabelResponse) {
    response = &ModifyAttributeLabelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAttributeLabel
// 编辑属性标签
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyAttributeLabel(request *ModifyAttributeLabelRequest) (response *ModifyAttributeLabelResponse, err error) {
    return c.ModifyAttributeLabelWithContext(context.Background(), request)
}

// ModifyAttributeLabel
// 编辑属性标签
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
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

func NewModifyDocRequest() (request *ModifyDocRequest) {
    request = &ModifyDocRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ModifyDoc")
    
    
    return
}

func NewModifyDocResponse() (response *ModifyDocResponse) {
    response = &ModifyDocResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDoc
// 修改文档
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyDoc(request *ModifyDocRequest) (response *ModifyDocResponse, err error) {
    return c.ModifyDocWithContext(context.Background(), request)
}

// ModifyDoc
// 修改文档
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyDocWithContext(ctx context.Context, request *ModifyDocRequest) (response *ModifyDocResponse, err error) {
    if request == nil {
        request = NewModifyDocRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDoc require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDocResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDocAttrRangeRequest() (request *ModifyDocAttrRangeRequest) {
    request = &ModifyDocAttrRangeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ModifyDocAttrRange")
    
    
    return
}

func NewModifyDocAttrRangeResponse() (response *ModifyDocAttrRangeResponse) {
    response = &ModifyDocAttrRangeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDocAttrRange
// 批量修改文档适用范围
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyDocAttrRange(request *ModifyDocAttrRangeRequest) (response *ModifyDocAttrRangeResponse, err error) {
    return c.ModifyDocAttrRangeWithContext(context.Background(), request)
}

// ModifyDocAttrRange
// 批量修改文档适用范围
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyDocAttrRangeWithContext(ctx context.Context, request *ModifyDocAttrRangeRequest) (response *ModifyDocAttrRangeResponse, err error) {
    if request == nil {
        request = NewModifyDocAttrRangeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDocAttrRange require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDocAttrRangeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyQARequest() (request *ModifyQARequest) {
    request = &ModifyQARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ModifyQA")
    
    
    return
}

func NewModifyQAResponse() (response *ModifyQAResponse) {
    response = &ModifyQAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyQA
// 更新问答
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyQA(request *ModifyQARequest) (response *ModifyQAResponse, err error) {
    return c.ModifyQAWithContext(context.Background(), request)
}

// ModifyQA
// 更新问答
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
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

func NewModifyQAAttrRangeRequest() (request *ModifyQAAttrRangeRequest) {
    request = &ModifyQAAttrRangeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ModifyQAAttrRange")
    
    
    return
}

func NewModifyQAAttrRangeResponse() (response *ModifyQAAttrRangeResponse) {
    response = &ModifyQAAttrRangeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyQAAttrRange
// 批量修改问答适用范围
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyQAAttrRange(request *ModifyQAAttrRangeRequest) (response *ModifyQAAttrRangeResponse, err error) {
    return c.ModifyQAAttrRangeWithContext(context.Background(), request)
}

// ModifyQAAttrRange
// 批量修改问答适用范围
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyQAAttrRangeWithContext(ctx context.Context, request *ModifyQAAttrRangeRequest) (response *ModifyQAAttrRangeResponse, err error) {
    if request == nil {
        request = NewModifyQAAttrRangeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyQAAttrRange require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyQAAttrRangeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyQACateRequest() (request *ModifyQACateRequest) {
    request = &ModifyQACateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ModifyQACate")
    
    
    return
}

func NewModifyQACateResponse() (response *ModifyQACateResponse) {
    response = &ModifyQACateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyQACate
// 更新QA分类
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyQACate(request *ModifyQACateRequest) (response *ModifyQACateResponse, err error) {
    return c.ModifyQACateWithContext(context.Background(), request)
}

// ModifyQACate
// 更新QA分类
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyQACateWithContext(ctx context.Context, request *ModifyQACateRequest) (response *ModifyQACateResponse, err error) {
    if request == nil {
        request = NewModifyQACateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyQACate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyQACateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRejectedQuestionRequest() (request *ModifyRejectedQuestionRequest) {
    request = &ModifyRejectedQuestionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ModifyRejectedQuestion")
    
    
    return
}

func NewModifyRejectedQuestionResponse() (response *ModifyRejectedQuestionResponse) {
    response = &ModifyRejectedQuestionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRejectedQuestion
// 修改拒答问题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyRejectedQuestion(request *ModifyRejectedQuestionRequest) (response *ModifyRejectedQuestionResponse, err error) {
    return c.ModifyRejectedQuestionWithContext(context.Background(), request)
}

// ModifyRejectedQuestion
// 修改拒答问题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyRejectedQuestionWithContext(ctx context.Context, request *ModifyRejectedQuestionRequest) (response *ModifyRejectedQuestionResponse, err error) {
    if request == nil {
        request = NewModifyRejectedQuestionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRejectedQuestion require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRejectedQuestionResponse()
    err = c.Send(request, response)
    return
}

func NewParseDocRequest() (request *ParseDocRequest) {
    request = &ParseDocRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ParseDoc")
    
    
    return
}

func NewParseDocResponse() (response *ParseDocResponse) {
    response = &ParseDocResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ParseDoc
// 接口即将下线，请切换使用新接口：[文档解析](https://cloud.tencent.com/document/product/1759/107504)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ParseDoc(request *ParseDocRequest) (response *ParseDocResponse, err error) {
    return c.ParseDocWithContext(context.Background(), request)
}

// ParseDoc
// 接口即将下线，请切换使用新接口：[文档解析](https://cloud.tencent.com/document/product/1759/107504)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) ParseDocWithContext(ctx context.Context, request *ParseDocRequest) (response *ParseDocResponse, err error) {
    if request == nil {
        request = NewParseDocRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ParseDoc require credential")
    }

    request.SetContext(ctx)
    
    response = NewParseDocResponse()
    err = c.Send(request, response)
    return
}

func NewQueryParseDocResultRequest() (request *QueryParseDocResultRequest) {
    request = &QueryParseDocResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "QueryParseDocResult")
    
    
    return
}

func NewQueryParseDocResultResponse() (response *QueryParseDocResultResponse) {
    response = &QueryParseDocResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryParseDocResult
// 查询文档解析结果。该接口需开通文档解析原子能力后调用。文档解析原子能力内测中，如有需要请联系架构师或[联系客服](https://cloud.tencent.com/act/event/Online_service) 。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) QueryParseDocResult(request *QueryParseDocResultRequest) (response *QueryParseDocResultResponse, err error) {
    return c.QueryParseDocResultWithContext(context.Background(), request)
}

// QueryParseDocResult
// 查询文档解析结果。该接口需开通文档解析原子能力后调用。文档解析原子能力内测中，如有需要请联系架构师或[联系客服](https://cloud.tencent.com/act/event/Online_service) 。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) QueryParseDocResultWithContext(ctx context.Context, request *QueryParseDocResultRequest) (response *QueryParseDocResultResponse, err error) {
    if request == nil {
        request = NewQueryParseDocResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryParseDocResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryParseDocResultResponse()
    err = c.Send(request, response)
    return
}

func NewQueryRewriteRequest() (request *QueryRewriteRequest) {
    request = &QueryRewriteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "QueryRewrite")
    
    
    return
}

func NewQueryRewriteResponse() (response *QueryRewriteResponse) {
    response = &QueryRewriteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryRewrite
// 多轮改写（QueryRewrite）主要用于多轮对话中，进行指代消解和省略补全。使用本接口，无需输入prompt描述，根据对话历史即可生成更精确的用户查询。在应用场景上，本接口可应用于智能问答、对话式搜索等多种场景。
//
// 本接口（QueryRewrite）有单账号调用上限控制，如您有提高并发限制的需求请 [联系我们](https://cloud.tencent.com/act/event/Online_service) 。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) QueryRewrite(request *QueryRewriteRequest) (response *QueryRewriteResponse, err error) {
    return c.QueryRewriteWithContext(context.Background(), request)
}

// QueryRewrite
// 多轮改写（QueryRewrite）主要用于多轮对话中，进行指代消解和省略补全。使用本接口，无需输入prompt描述，根据对话历史即可生成更精确的用户查询。在应用场景上，本接口可应用于智能问答、对话式搜索等多种场景。
//
// 本接口（QueryRewrite）有单账号调用上限控制，如您有提高并发限制的需求请 [联系我们](https://cloud.tencent.com/act/event/Online_service) 。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) QueryRewriteWithContext(ctx context.Context, request *QueryRewriteRequest) (response *QueryRewriteResponse, err error) {
    if request == nil {
        request = NewQueryRewriteRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryRewrite require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryRewriteResponse()
    err = c.Send(request, response)
    return
}

func NewRateMsgRecordRequest() (request *RateMsgRecordRequest) {
    request = &RateMsgRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "RateMsgRecord")
    
    
    return
}

func NewRateMsgRecordResponse() (response *RateMsgRecordResponse) {
    response = &RateMsgRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RateMsgRecord
// 点赞点踩消息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) RateMsgRecord(request *RateMsgRecordRequest) (response *RateMsgRecordResponse, err error) {
    return c.RateMsgRecordWithContext(context.Background(), request)
}

// RateMsgRecord
// 点赞点踩消息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) RateMsgRecordWithContext(ctx context.Context, request *RateMsgRecordRequest) (response *RateMsgRecordResponse, err error) {
    if request == nil {
        request = NewRateMsgRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RateMsgRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewRateMsgRecordResponse()
    err = c.Send(request, response)
    return
}

func NewReconstructDocumentRequest() (request *ReconstructDocumentRequest) {
    request = &ReconstructDocumentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ReconstructDocument")
    
    
    return
}

func NewReconstructDocumentResponse() (response *ReconstructDocumentResponse) {
    response = &ReconstructDocumentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReconstructDocument
// 支持将图片或PDF文件转换成Markdown格式文件，可解析包括表格、公式、图片、标题、段落、页眉、页脚等内容元素，并将内容智能转换成阅读顺序。
//
// 
//
// 体验期间单账号限制qps仅为1，若有正式接入需要请与产研团队沟通开放。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FILEDECODEFAILED = "FailedOperation.FileDecodeFailed"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) ReconstructDocument(request *ReconstructDocumentRequest) (response *ReconstructDocumentResponse, err error) {
    return c.ReconstructDocumentWithContext(context.Background(), request)
}

// ReconstructDocument
// 支持将图片或PDF文件转换成Markdown格式文件，可解析包括表格、公式、图片、标题、段落、页眉、页脚等内容元素，并将内容智能转换成阅读顺序。
//
// 
//
// 体验期间单账号限制qps仅为1，若有正式接入需要请与产研团队沟通开放。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FILEDECODEFAILED = "FailedOperation.FileDecodeFailed"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) ReconstructDocumentWithContext(ctx context.Context, request *ReconstructDocumentRequest) (response *ReconstructDocumentResponse, err error) {
    if request == nil {
        request = NewReconstructDocumentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReconstructDocument require credential")
    }

    request.SetContext(ctx)
    
    response = NewReconstructDocumentResponse()
    err = c.Send(request, response)
    return
}

func NewResetSessionRequest() (request *ResetSessionRequest) {
    request = &ResetSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ResetSession")
    
    
    return
}

func NewResetSessionResponse() (response *ResetSessionResponse) {
    response = &ResetSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetSession
// 重置会话
//
// 可能返回的错误码:
//  FAILEDOPERATION_FILEDECODEFAILED = "FailedOperation.FileDecodeFailed"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) ResetSession(request *ResetSessionRequest) (response *ResetSessionResponse, err error) {
    return c.ResetSessionWithContext(context.Background(), request)
}

// ResetSession
// 重置会话
//
// 可能返回的错误码:
//  FAILEDOPERATION_FILEDECODEFAILED = "FailedOperation.FileDecodeFailed"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) ResetSessionWithContext(ctx context.Context, request *ResetSessionRequest) (response *ResetSessionResponse, err error) {
    if request == nil {
        request = NewResetSessionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetSession require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetSessionResponse()
    err = c.Send(request, response)
    return
}

func NewRetryDocAuditRequest() (request *RetryDocAuditRequest) {
    request = &RetryDocAuditRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "RetryDocAudit")
    
    
    return
}

func NewRetryDocAuditResponse() (response *RetryDocAuditResponse) {
    response = &RetryDocAuditResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RetryDocAudit
// 文档解析重试
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RetryDocAudit(request *RetryDocAuditRequest) (response *RetryDocAuditResponse, err error) {
    return c.RetryDocAuditWithContext(context.Background(), request)
}

// RetryDocAudit
// 文档解析重试
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RetryDocAuditWithContext(ctx context.Context, request *RetryDocAuditRequest) (response *RetryDocAuditResponse, err error) {
    if request == nil {
        request = NewRetryDocAuditRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RetryDocAudit require credential")
    }

    request.SetContext(ctx)
    
    response = NewRetryDocAuditResponse()
    err = c.Send(request, response)
    return
}

func NewRetryDocParseRequest() (request *RetryDocParseRequest) {
    request = &RetryDocParseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "RetryDocParse")
    
    
    return
}

func NewRetryDocParseResponse() (response *RetryDocParseResponse) {
    response = &RetryDocParseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RetryDocParse
// 文档解析重试
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RetryDocParse(request *RetryDocParseRequest) (response *RetryDocParseResponse, err error) {
    return c.RetryDocParseWithContext(context.Background(), request)
}

// RetryDocParse
// 文档解析重试
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RetryDocParseWithContext(ctx context.Context, request *RetryDocParseRequest) (response *RetryDocParseResponse, err error) {
    if request == nil {
        request = NewRetryDocParseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RetryDocParse require credential")
    }

    request.SetContext(ctx)
    
    response = NewRetryDocParseResponse()
    err = c.Send(request, response)
    return
}

func NewRetryReleaseRequest() (request *RetryReleaseRequest) {
    request = &RetryReleaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "RetryRelease")
    
    
    return
}

func NewRetryReleaseResponse() (response *RetryReleaseResponse) {
    response = &RetryReleaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RetryRelease
// 发布暂停后重试
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RetryRelease(request *RetryReleaseRequest) (response *RetryReleaseResponse, err error) {
    return c.RetryReleaseWithContext(context.Background(), request)
}

// RetryRelease
// 发布暂停后重试
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RetryReleaseWithContext(ctx context.Context, request *RetryReleaseRequest) (response *RetryReleaseResponse, err error) {
    if request == nil {
        request = NewRetryReleaseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RetryRelease require credential")
    }

    request.SetContext(ctx)
    
    response = NewRetryReleaseResponse()
    err = c.Send(request, response)
    return
}

func NewSaveDocRequest() (request *SaveDocRequest) {
    request = &SaveDocRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "SaveDoc")
    
    
    return
}

func NewSaveDocResponse() (response *SaveDocResponse) {
    response = &SaveDocResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SaveDoc
// 保存文档
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) SaveDoc(request *SaveDocRequest) (response *SaveDocResponse, err error) {
    return c.SaveDocWithContext(context.Background(), request)
}

// SaveDoc
// 保存文档
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) SaveDocWithContext(ctx context.Context, request *SaveDocRequest) (response *SaveDocResponse, err error) {
    if request == nil {
        request = NewSaveDocRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SaveDoc require credential")
    }

    request.SetContext(ctx)
    
    response = NewSaveDocResponse()
    err = c.Send(request, response)
    return
}

func NewStopDocParseRequest() (request *StopDocParseRequest) {
    request = &StopDocParseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "StopDocParse")
    
    
    return
}

func NewStopDocParseResponse() (response *StopDocParseResponse) {
    response = &StopDocParseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopDocParse
// 终止文档解析
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) StopDocParse(request *StopDocParseRequest) (response *StopDocParseResponse, err error) {
    return c.StopDocParseWithContext(context.Background(), request)
}

// StopDocParse
// 终止文档解析
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) StopDocParseWithContext(ctx context.Context, request *StopDocParseRequest) (response *StopDocParseResponse, err error) {
    if request == nil {
        request = NewStopDocParseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopDocParse require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopDocParseResponse()
    err = c.Send(request, response)
    return
}

func NewUploadAttributeLabelRequest() (request *UploadAttributeLabelRequest) {
    request = &UploadAttributeLabelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "UploadAttributeLabel")
    
    
    return
}

func NewUploadAttributeLabelResponse() (response *UploadAttributeLabelResponse) {
    response = &UploadAttributeLabelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UploadAttributeLabel
// 上传导入属性标签
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UploadAttributeLabel(request *UploadAttributeLabelRequest) (response *UploadAttributeLabelResponse, err error) {
    return c.UploadAttributeLabelWithContext(context.Background(), request)
}

// UploadAttributeLabel
// 上传导入属性标签
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UploadAttributeLabelWithContext(ctx context.Context, request *UploadAttributeLabelRequest) (response *UploadAttributeLabelResponse, err error) {
    if request == nil {
        request = NewUploadAttributeLabelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadAttributeLabel require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadAttributeLabelResponse()
    err = c.Send(request, response)
    return
}

func NewVerifyQARequest() (request *VerifyQARequest) {
    request = &VerifyQARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "VerifyQA")
    
    
    return
}

func NewVerifyQAResponse() (response *VerifyQAResponse) {
    response = &VerifyQAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// VerifyQA
// 校验问答
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) VerifyQA(request *VerifyQARequest) (response *VerifyQAResponse, err error) {
    return c.VerifyQAWithContext(context.Background(), request)
}

// VerifyQA
// 校验问答
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) VerifyQAWithContext(ctx context.Context, request *VerifyQARequest) (response *VerifyQAResponse, err error) {
    if request == nil {
        request = NewVerifyQARequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("VerifyQA require credential")
    }

    request.SetContext(ctx)
    
    response = NewVerifyQAResponse()
    err = c.Send(request, response)
    return
}
