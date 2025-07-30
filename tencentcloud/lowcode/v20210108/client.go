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

package v20210108

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-01-08"

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


func NewCheckDeployAppRequest() (request *CheckDeployAppRequest) {
    request = &CheckDeployAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lowcode", APIVersion, "CheckDeployApp")
    
    
    return
}

func NewCheckDeployAppResponse() (response *CheckDeployAppResponse) {
    response = &CheckDeployAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckDeployApp
// 检查应用发布状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CheckDeployApp(request *CheckDeployAppRequest) (response *CheckDeployAppResponse, err error) {
    return c.CheckDeployAppWithContext(context.Background(), request)
}

// CheckDeployApp
// 检查应用发布状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CheckDeployAppWithContext(ctx context.Context, request *CheckDeployAppRequest) (response *CheckDeployAppResponse, err error) {
    if request == nil {
        request = NewCheckDeployAppRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lowcode", APIVersion, "CheckDeployApp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckDeployApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckDeployAppResponse()
    err = c.Send(request, response)
    return
}

func NewCreateKnowledgeSetRequest() (request *CreateKnowledgeSetRequest) {
    request = &CreateKnowledgeSetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lowcode", APIVersion, "CreateKnowledgeSet")
    
    
    return
}

func NewCreateKnowledgeSetResponse() (response *CreateKnowledgeSetResponse) {
    response = &CreateKnowledgeSetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateKnowledgeSet
// 创建知识库
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateKnowledgeSet(request *CreateKnowledgeSetRequest) (response *CreateKnowledgeSetResponse, err error) {
    return c.CreateKnowledgeSetWithContext(context.Background(), request)
}

// CreateKnowledgeSet
// 创建知识库
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateKnowledgeSetWithContext(ctx context.Context, request *CreateKnowledgeSetRequest) (response *CreateKnowledgeSetResponse, err error) {
    if request == nil {
        request = NewCreateKnowledgeSetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lowcode", APIVersion, "CreateKnowledgeSet")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateKnowledgeSet require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateKnowledgeSetResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAppBindWxAppRequest() (request *DeleteAppBindWxAppRequest) {
    request = &DeleteAppBindWxAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lowcode", APIVersion, "DeleteAppBindWxApp")
    
    
    return
}

func NewDeleteAppBindWxAppResponse() (response *DeleteAppBindWxAppResponse) {
    response = &DeleteAppBindWxAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAppBindWxApp
// 删除应用绑定小程序
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) DeleteAppBindWxApp(request *DeleteAppBindWxAppRequest) (response *DeleteAppBindWxAppResponse, err error) {
    return c.DeleteAppBindWxAppWithContext(context.Background(), request)
}

// DeleteAppBindWxApp
// 删除应用绑定小程序
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) DeleteAppBindWxAppWithContext(ctx context.Context, request *DeleteAppBindWxAppRequest) (response *DeleteAppBindWxAppResponse, err error) {
    if request == nil {
        request = NewDeleteAppBindWxAppRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lowcode", APIVersion, "DeleteAppBindWxApp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAppBindWxApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAppBindWxAppResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteKnowledgeDocumentSetRequest() (request *DeleteKnowledgeDocumentSetRequest) {
    request = &DeleteKnowledgeDocumentSetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lowcode", APIVersion, "DeleteKnowledgeDocumentSet")
    
    
    return
}

func NewDeleteKnowledgeDocumentSetResponse() (response *DeleteKnowledgeDocumentSetResponse) {
    response = &DeleteKnowledgeDocumentSetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteKnowledgeDocumentSet
// 删除知识库下文档
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOCPARSINGNOTDELETE = "FailedOperation.DocParsingNotDelete"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteKnowledgeDocumentSet(request *DeleteKnowledgeDocumentSetRequest) (response *DeleteKnowledgeDocumentSetResponse, err error) {
    return c.DeleteKnowledgeDocumentSetWithContext(context.Background(), request)
}

// DeleteKnowledgeDocumentSet
// 删除知识库下文档
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOCPARSINGNOTDELETE = "FailedOperation.DocParsingNotDelete"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteKnowledgeDocumentSetWithContext(ctx context.Context, request *DeleteKnowledgeDocumentSetRequest) (response *DeleteKnowledgeDocumentSetResponse, err error) {
    if request == nil {
        request = NewDeleteKnowledgeDocumentSetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lowcode", APIVersion, "DeleteKnowledgeDocumentSet")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteKnowledgeDocumentSet require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteKnowledgeDocumentSetResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteKnowledgeSetRequest() (request *DeleteKnowledgeSetRequest) {
    request = &DeleteKnowledgeSetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lowcode", APIVersion, "DeleteKnowledgeSet")
    
    
    return
}

func NewDeleteKnowledgeSetResponse() (response *DeleteKnowledgeSetResponse) {
    response = &DeleteKnowledgeSetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteKnowledgeSet
// 删除知识库
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteKnowledgeSet(request *DeleteKnowledgeSetRequest) (response *DeleteKnowledgeSetResponse, err error) {
    return c.DeleteKnowledgeSetWithContext(context.Background(), request)
}

// DeleteKnowledgeSet
// 删除知识库
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteKnowledgeSetWithContext(ctx context.Context, request *DeleteKnowledgeSetRequest) (response *DeleteKnowledgeSetResponse, err error) {
    if request == nil {
        request = NewDeleteKnowledgeSetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lowcode", APIVersion, "DeleteKnowledgeSet")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteKnowledgeSet require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteKnowledgeSetResponse()
    err = c.Send(request, response)
    return
}

func NewDeployAppRequest() (request *DeployAppRequest) {
    request = &DeployAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lowcode", APIVersion, "DeployApp")
    
    
    return
}

func NewDeployAppResponse() (response *DeployAppResponse) {
    response = &DeployAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeployApp
// 发布应用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeployApp(request *DeployAppRequest) (response *DeployAppResponse, err error) {
    return c.DeployAppWithContext(context.Background(), request)
}

// DeployApp
// 发布应用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeployAppWithContext(ctx context.Context, request *DeployAppRequest) (response *DeployAppResponse, err error) {
    if request == nil {
        request = NewDeployAppRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lowcode", APIVersion, "DeployApp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeployApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeployAppResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataSourceListRequest() (request *DescribeDataSourceListRequest) {
    request = &DescribeDataSourceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lowcode", APIVersion, "DescribeDataSourceList")
    
    
    return
}

func NewDescribeDataSourceListResponse() (response *DescribeDataSourceListResponse) {
    response = &DescribeDataSourceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDataSourceList
// 获取数据源详情列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOWCODEUSERNOTEXIST = "FailedOperation.LowcodeUserNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDataSourceList(request *DescribeDataSourceListRequest) (response *DescribeDataSourceListResponse, err error) {
    return c.DescribeDataSourceListWithContext(context.Background(), request)
}

// DescribeDataSourceList
// 获取数据源详情列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOWCODEUSERNOTEXIST = "FailedOperation.LowcodeUserNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDataSourceListWithContext(ctx context.Context, request *DescribeDataSourceListRequest) (response *DescribeDataSourceListResponse, err error) {
    if request == nil {
        request = NewDescribeDataSourceListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lowcode", APIVersion, "DescribeDataSourceList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataSourceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataSourceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKnowledgeDocumentSetDetailRequest() (request *DescribeKnowledgeDocumentSetDetailRequest) {
    request = &DescribeKnowledgeDocumentSetDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lowcode", APIVersion, "DescribeKnowledgeDocumentSetDetail")
    
    
    return
}

func NewDescribeKnowledgeDocumentSetDetailResponse() (response *DescribeKnowledgeDocumentSetDetailResponse) {
    response = &DescribeKnowledgeDocumentSetDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeKnowledgeDocumentSetDetail
// 获取知识库下文档详情
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeKnowledgeDocumentSetDetail(request *DescribeKnowledgeDocumentSetDetailRequest) (response *DescribeKnowledgeDocumentSetDetailResponse, err error) {
    return c.DescribeKnowledgeDocumentSetDetailWithContext(context.Background(), request)
}

// DescribeKnowledgeDocumentSetDetail
// 获取知识库下文档详情
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeKnowledgeDocumentSetDetailWithContext(ctx context.Context, request *DescribeKnowledgeDocumentSetDetailRequest) (response *DescribeKnowledgeDocumentSetDetailResponse, err error) {
    if request == nil {
        request = NewDescribeKnowledgeDocumentSetDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lowcode", APIVersion, "DescribeKnowledgeDocumentSetDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKnowledgeDocumentSetDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKnowledgeDocumentSetDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKnowledgeDocumentSetListRequest() (request *DescribeKnowledgeDocumentSetListRequest) {
    request = &DescribeKnowledgeDocumentSetListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lowcode", APIVersion, "DescribeKnowledgeDocumentSetList")
    
    
    return
}

func NewDescribeKnowledgeDocumentSetListResponse() (response *DescribeKnowledgeDocumentSetListResponse) {
    response = &DescribeKnowledgeDocumentSetListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeKnowledgeDocumentSetList
// 查询知识库下文件集合
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeKnowledgeDocumentSetList(request *DescribeKnowledgeDocumentSetListRequest) (response *DescribeKnowledgeDocumentSetListResponse, err error) {
    return c.DescribeKnowledgeDocumentSetListWithContext(context.Background(), request)
}

// DescribeKnowledgeDocumentSetList
// 查询知识库下文件集合
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeKnowledgeDocumentSetListWithContext(ctx context.Context, request *DescribeKnowledgeDocumentSetListRequest) (response *DescribeKnowledgeDocumentSetListResponse, err error) {
    if request == nil {
        request = NewDescribeKnowledgeDocumentSetListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lowcode", APIVersion, "DescribeKnowledgeDocumentSetList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKnowledgeDocumentSetList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKnowledgeDocumentSetListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKnowledgeSetListRequest() (request *DescribeKnowledgeSetListRequest) {
    request = &DescribeKnowledgeSetListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lowcode", APIVersion, "DescribeKnowledgeSetList")
    
    
    return
}

func NewDescribeKnowledgeSetListResponse() (response *DescribeKnowledgeSetListResponse) {
    response = &DescribeKnowledgeSetListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeKnowledgeSetList
// 查询知识库
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeKnowledgeSetList(request *DescribeKnowledgeSetListRequest) (response *DescribeKnowledgeSetListResponse, err error) {
    return c.DescribeKnowledgeSetListWithContext(context.Background(), request)
}

// DescribeKnowledgeSetList
// 查询知识库
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeKnowledgeSetListWithContext(ctx context.Context, request *DescribeKnowledgeSetListRequest) (response *DescribeKnowledgeSetListResponse, err error) {
    if request == nil {
        request = NewDescribeKnowledgeSetListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lowcode", APIVersion, "DescribeKnowledgeSetList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKnowledgeSetList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKnowledgeSetListResponse()
    err = c.Send(request, response)
    return
}

func NewPutWxAppIdToWeAppRequest() (request *PutWxAppIdToWeAppRequest) {
    request = &PutWxAppIdToWeAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lowcode", APIVersion, "PutWxAppIdToWeApp")
    
    
    return
}

func NewPutWxAppIdToWeAppResponse() (response *PutWxAppIdToWeAppResponse) {
    response = &PutWxAppIdToWeAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PutWxAppIdToWeApp
// 接口提供应用绑定微信ID功能。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) PutWxAppIdToWeApp(request *PutWxAppIdToWeAppRequest) (response *PutWxAppIdToWeAppResponse, err error) {
    return c.PutWxAppIdToWeAppWithContext(context.Background(), request)
}

// PutWxAppIdToWeApp
// 接口提供应用绑定微信ID功能。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) PutWxAppIdToWeAppWithContext(ctx context.Context, request *PutWxAppIdToWeAppRequest) (response *PutWxAppIdToWeAppResponse, err error) {
    if request == nil {
        request = NewPutWxAppIdToWeAppRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lowcode", APIVersion, "PutWxAppIdToWeApp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("PutWxAppIdToWeApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewPutWxAppIdToWeAppResponse()
    err = c.Send(request, response)
    return
}

func NewSearchDocListRequest() (request *SearchDocListRequest) {
    request = &SearchDocListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lowcode", APIVersion, "SearchDocList")
    
    
    return
}

func NewSearchDocListResponse() (response *SearchDocListResponse) {
    response = &SearchDocListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SearchDocList
// 知识库文档搜索接口
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) SearchDocList(request *SearchDocListRequest) (response *SearchDocListResponse, err error) {
    return c.SearchDocListWithContext(context.Background(), request)
}

// SearchDocList
// 知识库文档搜索接口
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) SearchDocListWithContext(ctx context.Context, request *SearchDocListRequest) (response *SearchDocListResponse, err error) {
    if request == nil {
        request = NewSearchDocListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lowcode", APIVersion, "SearchDocList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchDocList require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchDocListResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateKnowledgeSetRequest() (request *UpdateKnowledgeSetRequest) {
    request = &UpdateKnowledgeSetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lowcode", APIVersion, "UpdateKnowledgeSet")
    
    
    return
}

func NewUpdateKnowledgeSetResponse() (response *UpdateKnowledgeSetResponse) {
    response = &UpdateKnowledgeSetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateKnowledgeSet
// 更新知识库
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UpdateKnowledgeSet(request *UpdateKnowledgeSetRequest) (response *UpdateKnowledgeSetResponse, err error) {
    return c.UpdateKnowledgeSetWithContext(context.Background(), request)
}

// UpdateKnowledgeSet
// 更新知识库
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UpdateKnowledgeSetWithContext(ctx context.Context, request *UpdateKnowledgeSetRequest) (response *UpdateKnowledgeSetResponse, err error) {
    if request == nil {
        request = NewUpdateKnowledgeSetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lowcode", APIVersion, "UpdateKnowledgeSet")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateKnowledgeSet require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateKnowledgeSetResponse()
    err = c.Send(request, response)
    return
}

func NewUploadKnowledgeDocumentSetRequest() (request *UploadKnowledgeDocumentSetRequest) {
    request = &UploadKnowledgeDocumentSetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lowcode", APIVersion, "UploadKnowledgeDocumentSet")
    
    
    return
}

func NewUploadKnowledgeDocumentSetResponse() (response *UploadKnowledgeDocumentSetResponse) {
    response = &UploadKnowledgeDocumentSetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UploadKnowledgeDocumentSet
// 更新知识库
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UploadKnowledgeDocumentSet(request *UploadKnowledgeDocumentSetRequest) (response *UploadKnowledgeDocumentSetResponse, err error) {
    return c.UploadKnowledgeDocumentSetWithContext(context.Background(), request)
}

// UploadKnowledgeDocumentSet
// 更新知识库
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UploadKnowledgeDocumentSetWithContext(ctx context.Context, request *UploadKnowledgeDocumentSetRequest) (response *UploadKnowledgeDocumentSetResponse, err error) {
    if request == nil {
        request = NewUploadKnowledgeDocumentSetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lowcode", APIVersion, "UploadKnowledgeDocumentSet")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadKnowledgeDocumentSet require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadKnowledgeDocumentSetResponse()
    err = c.Send(request, response)
    return
}
