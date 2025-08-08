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

package v20190723

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-07-23"

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


func NewAuthorizeDSPAMetaResourcesRequest() (request *AuthorizeDSPAMetaResourcesRequest) {
    request = &AuthorizeDSPAMetaResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "AuthorizeDSPAMetaResources")
    
    
    return
}

func NewAuthorizeDSPAMetaResourcesResponse() (response *AuthorizeDSPAMetaResourcesResponse) {
    response = &AuthorizeDSPAMetaResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AuthorizeDSPAMetaResources
// 授权用户云资源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AuthorizeDSPAMetaResources(request *AuthorizeDSPAMetaResourcesRequest) (response *AuthorizeDSPAMetaResourcesResponse, err error) {
    return c.AuthorizeDSPAMetaResourcesWithContext(context.Background(), request)
}

// AuthorizeDSPAMetaResources
// 授权用户云资源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AuthorizeDSPAMetaResourcesWithContext(ctx context.Context, request *AuthorizeDSPAMetaResourcesRequest) (response *AuthorizeDSPAMetaResourcesResponse, err error) {
    if request == nil {
        request = NewAuthorizeDSPAMetaResourcesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "AuthorizeDSPAMetaResources")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AuthorizeDSPAMetaResources require credential")
    }

    request.SetContext(ctx)
    
    response = NewAuthorizeDSPAMetaResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewBindDSPAResourceCosBucketsRequest() (request *BindDSPAResourceCosBucketsRequest) {
    request = &BindDSPAResourceCosBucketsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "BindDSPAResourceCosBuckets")
    
    
    return
}

func NewBindDSPAResourceCosBucketsResponse() (response *BindDSPAResourceCosBucketsResponse) {
    response = &BindDSPAResourceCosBucketsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindDSPAResourceCosBuckets
// 绑定或解绑COS桶
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) BindDSPAResourceCosBuckets(request *BindDSPAResourceCosBucketsRequest) (response *BindDSPAResourceCosBucketsResponse, err error) {
    return c.BindDSPAResourceCosBucketsWithContext(context.Background(), request)
}

// BindDSPAResourceCosBuckets
// 绑定或解绑COS桶
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) BindDSPAResourceCosBucketsWithContext(ctx context.Context, request *BindDSPAResourceCosBucketsRequest) (response *BindDSPAResourceCosBucketsResponse, err error) {
    if request == nil {
        request = NewBindDSPAResourceCosBucketsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "BindDSPAResourceCosBuckets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindDSPAResourceCosBuckets require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindDSPAResourceCosBucketsResponse()
    err = c.Send(request, response)
    return
}

func NewBindDSPAResourceDatabasesRequest() (request *BindDSPAResourceDatabasesRequest) {
    request = &BindDSPAResourceDatabasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "BindDSPAResourceDatabases")
    
    
    return
}

func NewBindDSPAResourceDatabasesResponse() (response *BindDSPAResourceDatabasesResponse) {
    response = &BindDSPAResourceDatabasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindDSPAResourceDatabases
// 绑定或解绑数据库实例DB
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) BindDSPAResourceDatabases(request *BindDSPAResourceDatabasesRequest) (response *BindDSPAResourceDatabasesResponse, err error) {
    return c.BindDSPAResourceDatabasesWithContext(context.Background(), request)
}

// BindDSPAResourceDatabases
// 绑定或解绑数据库实例DB
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) BindDSPAResourceDatabasesWithContext(ctx context.Context, request *BindDSPAResourceDatabasesRequest) (response *BindDSPAResourceDatabasesResponse, err error) {
    if request == nil {
        request = NewBindDSPAResourceDatabasesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "BindDSPAResourceDatabases")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindDSPAResourceDatabases require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindDSPAResourceDatabasesResponse()
    err = c.Send(request, response)
    return
}

func NewCopyDSPATemplateRequest() (request *CopyDSPATemplateRequest) {
    request = &CopyDSPATemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "CopyDSPATemplate")
    
    
    return
}

func NewCopyDSPATemplateResponse() (response *CopyDSPATemplateResponse) {
    response = &CopyDSPATemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CopyDSPATemplate
// 复制合规组模板
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CopyDSPATemplate(request *CopyDSPATemplateRequest) (response *CopyDSPATemplateResponse, err error) {
    return c.CopyDSPATemplateWithContext(context.Background(), request)
}

// CopyDSPATemplate
// 复制合规组模板
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CopyDSPATemplateWithContext(ctx context.Context, request *CopyDSPATemplateRequest) (response *CopyDSPATemplateResponse, err error) {
    if request == nil {
        request = NewCopyDSPATemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "CopyDSPATemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CopyDSPATemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCopyDSPATemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAssetSortingReportRetryTaskRequest() (request *CreateAssetSortingReportRetryTaskRequest) {
    request = &CreateAssetSortingReportRetryTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "CreateAssetSortingReportRetryTask")
    
    
    return
}

func NewCreateAssetSortingReportRetryTaskResponse() (response *CreateAssetSortingReportRetryTaskResponse) {
    response = &CreateAssetSortingReportRetryTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAssetSortingReportRetryTask
// 创建资产梳理报表导出重试任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateAssetSortingReportRetryTask(request *CreateAssetSortingReportRetryTaskRequest) (response *CreateAssetSortingReportRetryTaskResponse, err error) {
    return c.CreateAssetSortingReportRetryTaskWithContext(context.Background(), request)
}

// CreateAssetSortingReportRetryTask
// 创建资产梳理报表导出重试任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateAssetSortingReportRetryTaskWithContext(ctx context.Context, request *CreateAssetSortingReportRetryTaskRequest) (response *CreateAssetSortingReportRetryTaskResponse, err error) {
    if request == nil {
        request = NewCreateAssetSortingReportRetryTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "CreateAssetSortingReportRetryTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAssetSortingReportRetryTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAssetSortingReportRetryTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAssetSortingReportTaskRequest() (request *CreateAssetSortingReportTaskRequest) {
    request = &CreateAssetSortingReportTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "CreateAssetSortingReportTask")
    
    
    return
}

func NewCreateAssetSortingReportTaskResponse() (response *CreateAssetSortingReportTaskResponse) {
    response = &CreateAssetSortingReportTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAssetSortingReportTask
// 创建资产梳理报告任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateAssetSortingReportTask(request *CreateAssetSortingReportTaskRequest) (response *CreateAssetSortingReportTaskResponse, err error) {
    return c.CreateAssetSortingReportTaskWithContext(context.Background(), request)
}

// CreateAssetSortingReportTask
// 创建资产梳理报告任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateAssetSortingReportTaskWithContext(ctx context.Context, request *CreateAssetSortingReportTaskRequest) (response *CreateAssetSortingReportTaskResponse, err error) {
    if request == nil {
        request = NewCreateAssetSortingReportTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "CreateAssetSortingReportTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAssetSortingReportTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAssetSortingReportTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDSPAAssessmentRiskLevelRequest() (request *CreateDSPAAssessmentRiskLevelRequest) {
    request = &CreateDSPAAssessmentRiskLevelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "CreateDSPAAssessmentRiskLevel")
    
    
    return
}

func NewCreateDSPAAssessmentRiskLevelResponse() (response *CreateDSPAAssessmentRiskLevelResponse) {
    response = &CreateDSPAAssessmentRiskLevelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDSPAAssessmentRiskLevel
// 风险项页面---创建风险等级
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateDSPAAssessmentRiskLevel(request *CreateDSPAAssessmentRiskLevelRequest) (response *CreateDSPAAssessmentRiskLevelResponse, err error) {
    return c.CreateDSPAAssessmentRiskLevelWithContext(context.Background(), request)
}

// CreateDSPAAssessmentRiskLevel
// 风险项页面---创建风险等级
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateDSPAAssessmentRiskLevelWithContext(ctx context.Context, request *CreateDSPAAssessmentRiskLevelRequest) (response *CreateDSPAAssessmentRiskLevelResponse, err error) {
    if request == nil {
        request = NewCreateDSPAAssessmentRiskLevelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "CreateDSPAAssessmentRiskLevel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDSPAAssessmentRiskLevel require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDSPAAssessmentRiskLevelResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDSPAAssessmentRiskTemplateRequest() (request *CreateDSPAAssessmentRiskTemplateRequest) {
    request = &CreateDSPAAssessmentRiskTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "CreateDSPAAssessmentRiskTemplate")
    
    
    return
}

func NewCreateDSPAAssessmentRiskTemplateResponse() (response *CreateDSPAAssessmentRiskTemplateResponse) {
    response = &CreateDSPAAssessmentRiskTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDSPAAssessmentRiskTemplate
// 风险评估模板---创建风险评估模板
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateDSPAAssessmentRiskTemplate(request *CreateDSPAAssessmentRiskTemplateRequest) (response *CreateDSPAAssessmentRiskTemplateResponse, err error) {
    return c.CreateDSPAAssessmentRiskTemplateWithContext(context.Background(), request)
}

// CreateDSPAAssessmentRiskTemplate
// 风险评估模板---创建风险评估模板
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateDSPAAssessmentRiskTemplateWithContext(ctx context.Context, request *CreateDSPAAssessmentRiskTemplateRequest) (response *CreateDSPAAssessmentRiskTemplateResponse, err error) {
    if request == nil {
        request = NewCreateDSPAAssessmentRiskTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "CreateDSPAAssessmentRiskTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDSPAAssessmentRiskTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDSPAAssessmentRiskTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDSPAAssessmentTaskRequest() (request *CreateDSPAAssessmentTaskRequest) {
    request = &CreateDSPAAssessmentTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "CreateDSPAAssessmentTask")
    
    
    return
}

func NewCreateDSPAAssessmentTaskResponse() (response *CreateDSPAAssessmentTaskResponse) {
    response = &CreateDSPAAssessmentTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDSPAAssessmentTask
// 新建DSPA风险评估任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDSPAAssessmentTask(request *CreateDSPAAssessmentTaskRequest) (response *CreateDSPAAssessmentTaskResponse, err error) {
    return c.CreateDSPAAssessmentTaskWithContext(context.Background(), request)
}

// CreateDSPAAssessmentTask
// 新建DSPA风险评估任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDSPAAssessmentTaskWithContext(ctx context.Context, request *CreateDSPAAssessmentTaskRequest) (response *CreateDSPAAssessmentTaskResponse, err error) {
    if request == nil {
        request = NewCreateDSPAAssessmentTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "CreateDSPAAssessmentTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDSPAAssessmentTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDSPAAssessmentTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDSPACOSDiscoveryTaskRequest() (request *CreateDSPACOSDiscoveryTaskRequest) {
    request = &CreateDSPACOSDiscoveryTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "CreateDSPACOSDiscoveryTask")
    
    
    return
}

func NewCreateDSPACOSDiscoveryTaskResponse() (response *CreateDSPACOSDiscoveryTaskResponse) {
    response = &CreateDSPACOSDiscoveryTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDSPACOSDiscoveryTask
// 新增COS分类分级扫描任务，单个用户最多允许创建100个任务。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateDSPACOSDiscoveryTask(request *CreateDSPACOSDiscoveryTaskRequest) (response *CreateDSPACOSDiscoveryTaskResponse, err error) {
    return c.CreateDSPACOSDiscoveryTaskWithContext(context.Background(), request)
}

// CreateDSPACOSDiscoveryTask
// 新增COS分类分级扫描任务，单个用户最多允许创建100个任务。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateDSPACOSDiscoveryTaskWithContext(ctx context.Context, request *CreateDSPACOSDiscoveryTaskRequest) (response *CreateDSPACOSDiscoveryTaskResponse, err error) {
    if request == nil {
        request = NewCreateDSPACOSDiscoveryTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "CreateDSPACOSDiscoveryTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDSPACOSDiscoveryTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDSPACOSDiscoveryTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDSPACategoryRequest() (request *CreateDSPACategoryRequest) {
    request = &CreateDSPACategoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "CreateDSPACategory")
    
    
    return
}

func NewCreateDSPACategoryResponse() (response *CreateDSPACategoryResponse) {
    response = &CreateDSPACategoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDSPACategory
// 新增分类，单个用户最多允许创建100个数据分类。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateDSPACategory(request *CreateDSPACategoryRequest) (response *CreateDSPACategoryResponse, err error) {
    return c.CreateDSPACategoryWithContext(context.Background(), request)
}

// CreateDSPACategory
// 新增分类，单个用户最多允许创建100个数据分类。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateDSPACategoryWithContext(ctx context.Context, request *CreateDSPACategoryRequest) (response *CreateDSPACategoryResponse, err error) {
    if request == nil {
        request = NewCreateDSPACategoryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "CreateDSPACategory")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDSPACategory require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDSPACategoryResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDSPACategoryRelationRequest() (request *CreateDSPACategoryRelationRequest) {
    request = &CreateDSPACategoryRelationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "CreateDSPACategoryRelation")
    
    
    return
}

func NewCreateDSPACategoryRelationResponse() (response *CreateDSPACategoryRelationResponse) {
    response = &CreateDSPACategoryRelationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDSPACategoryRelation
// 创建分类关联关系
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateDSPACategoryRelation(request *CreateDSPACategoryRelationRequest) (response *CreateDSPACategoryRelationResponse, err error) {
    return c.CreateDSPACategoryRelationWithContext(context.Background(), request)
}

// CreateDSPACategoryRelation
// 创建分类关联关系
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateDSPACategoryRelationWithContext(ctx context.Context, request *CreateDSPACategoryRelationRequest) (response *CreateDSPACategoryRelationResponse, err error) {
    if request == nil {
        request = NewCreateDSPACategoryRelationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "CreateDSPACategoryRelation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDSPACategoryRelation require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDSPACategoryRelationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDSPAComplianceGroupRequest() (request *CreateDSPAComplianceGroupRequest) {
    request = &CreateDSPAComplianceGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "CreateDSPAComplianceGroup")
    
    
    return
}

func NewCreateDSPAComplianceGroupResponse() (response *CreateDSPAComplianceGroupResponse) {
    response = &CreateDSPAComplianceGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDSPAComplianceGroup
// 新增分类分级模板，单个用户最多允许创建100个合规组。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateDSPAComplianceGroup(request *CreateDSPAComplianceGroupRequest) (response *CreateDSPAComplianceGroupResponse, err error) {
    return c.CreateDSPAComplianceGroupWithContext(context.Background(), request)
}

// CreateDSPAComplianceGroup
// 新增分类分级模板，单个用户最多允许创建100个合规组。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateDSPAComplianceGroupWithContext(ctx context.Context, request *CreateDSPAComplianceGroupRequest) (response *CreateDSPAComplianceGroupResponse, err error) {
    if request == nil {
        request = NewCreateDSPAComplianceGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "CreateDSPAComplianceGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDSPAComplianceGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDSPAComplianceGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDSPAComplianceRulesRequest() (request *CreateDSPAComplianceRulesRequest) {
    request = &CreateDSPAComplianceRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "CreateDSPAComplianceRules")
    
    
    return
}

func NewCreateDSPAComplianceRulesResponse() (response *CreateDSPAComplianceRulesResponse) {
    response = &CreateDSPAComplianceRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDSPAComplianceRules
// 创建合规组分类规则关联关系
//
// 可能返回的错误码:
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDSPAComplianceRules(request *CreateDSPAComplianceRulesRequest) (response *CreateDSPAComplianceRulesResponse, err error) {
    return c.CreateDSPAComplianceRulesWithContext(context.Background(), request)
}

// CreateDSPAComplianceRules
// 创建合规组分类规则关联关系
//
// 可能返回的错误码:
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDSPAComplianceRulesWithContext(ctx context.Context, request *CreateDSPAComplianceRulesRequest) (response *CreateDSPAComplianceRulesResponse, err error) {
    if request == nil {
        request = NewCreateDSPAComplianceRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "CreateDSPAComplianceRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDSPAComplianceRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDSPAComplianceRulesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDSPACosMetaResourcesRequest() (request *CreateDSPACosMetaResourcesRequest) {
    request = &CreateDSPACosMetaResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "CreateDSPACosMetaResources")
    
    
    return
}

func NewCreateDSPACosMetaResourcesResponse() (response *CreateDSPACosMetaResourcesResponse) {
    response = &CreateDSPACosMetaResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDSPACosMetaResources
// 添加COS元数据
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateDSPACosMetaResources(request *CreateDSPACosMetaResourcesRequest) (response *CreateDSPACosMetaResourcesResponse, err error) {
    return c.CreateDSPACosMetaResourcesWithContext(context.Background(), request)
}

// CreateDSPACosMetaResources
// 添加COS元数据
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateDSPACosMetaResourcesWithContext(ctx context.Context, request *CreateDSPACosMetaResourcesRequest) (response *CreateDSPACosMetaResourcesResponse, err error) {
    if request == nil {
        request = NewCreateDSPACosMetaResourcesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "CreateDSPACosMetaResources")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDSPACosMetaResources require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDSPACosMetaResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDSPADbMetaResourcesRequest() (request *CreateDSPADbMetaResourcesRequest) {
    request = &CreateDSPADbMetaResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "CreateDSPADbMetaResources")
    
    
    return
}

func NewCreateDSPADbMetaResourcesResponse() (response *CreateDSPADbMetaResourcesResponse) {
    response = &CreateDSPADbMetaResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDSPADbMetaResources
// 添加用户云上数据库类型资源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateDSPADbMetaResources(request *CreateDSPADbMetaResourcesRequest) (response *CreateDSPADbMetaResourcesResponse, err error) {
    return c.CreateDSPADbMetaResourcesWithContext(context.Background(), request)
}

// CreateDSPADbMetaResources
// 添加用户云上数据库类型资源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateDSPADbMetaResourcesWithContext(ctx context.Context, request *CreateDSPADbMetaResourcesRequest) (response *CreateDSPADbMetaResourcesResponse, err error) {
    if request == nil {
        request = NewCreateDSPADbMetaResourcesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "CreateDSPADbMetaResources")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDSPADbMetaResources require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDSPADbMetaResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDSPADiscoveryRuleRequest() (request *CreateDSPADiscoveryRuleRequest) {
    request = &CreateDSPADiscoveryRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "CreateDSPADiscoveryRule")
    
    
    return
}

func NewCreateDSPADiscoveryRuleResponse() (response *CreateDSPADiscoveryRuleResponse) {
    response = &CreateDSPADiscoveryRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDSPADiscoveryRule
// 新增分类分级规则，单个用户最多允许创建200个规则。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateDSPADiscoveryRule(request *CreateDSPADiscoveryRuleRequest) (response *CreateDSPADiscoveryRuleResponse, err error) {
    return c.CreateDSPADiscoveryRuleWithContext(context.Background(), request)
}

// CreateDSPADiscoveryRule
// 新增分类分级规则，单个用户最多允许创建200个规则。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateDSPADiscoveryRuleWithContext(ctx context.Context, request *CreateDSPADiscoveryRuleRequest) (response *CreateDSPADiscoveryRuleResponse, err error) {
    if request == nil {
        request = NewCreateDSPADiscoveryRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "CreateDSPADiscoveryRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDSPADiscoveryRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDSPADiscoveryRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDSPADiscoveryTaskRequest() (request *CreateDSPADiscoveryTaskRequest) {
    request = &CreateDSPADiscoveryTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "CreateDSPADiscoveryTask")
    
    
    return
}

func NewCreateDSPADiscoveryTaskResponse() (response *CreateDSPADiscoveryTaskResponse) {
    response = &CreateDSPADiscoveryTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDSPADiscoveryTask
// 新增分类分级任务，单个用户最多允许创建100个任务。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateDSPADiscoveryTask(request *CreateDSPADiscoveryTaskRequest) (response *CreateDSPADiscoveryTaskResponse, err error) {
    return c.CreateDSPADiscoveryTaskWithContext(context.Background(), request)
}

// CreateDSPADiscoveryTask
// 新增分类分级任务，单个用户最多允许创建100个任务。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateDSPADiscoveryTaskWithContext(ctx context.Context, request *CreateDSPADiscoveryTaskRequest) (response *CreateDSPADiscoveryTaskResponse, err error) {
    if request == nil {
        request = NewCreateDSPADiscoveryTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "CreateDSPADiscoveryTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDSPADiscoveryTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDSPADiscoveryTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDSPALevelGroupRequest() (request *CreateDSPALevelGroupRequest) {
    request = &CreateDSPALevelGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "CreateDSPALevelGroup")
    
    
    return
}

func NewCreateDSPALevelGroupResponse() (response *CreateDSPALevelGroupResponse) {
    response = &CreateDSPALevelGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDSPALevelGroup
// 新增分级，单个Casb实例最多允许创建100个数据分级
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateDSPALevelGroup(request *CreateDSPALevelGroupRequest) (response *CreateDSPALevelGroupResponse, err error) {
    return c.CreateDSPALevelGroupWithContext(context.Background(), request)
}

// CreateDSPALevelGroup
// 新增分级，单个Casb实例最多允许创建100个数据分级
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateDSPALevelGroupWithContext(ctx context.Context, request *CreateDSPALevelGroupRequest) (response *CreateDSPALevelGroupResponse, err error) {
    if request == nil {
        request = NewCreateDSPALevelGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "CreateDSPALevelGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDSPALevelGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDSPALevelGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDSPAMetaResourcesRequest() (request *CreateDSPAMetaResourcesRequest) {
    request = &CreateDSPAMetaResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "CreateDSPAMetaResources")
    
    
    return
}

func NewCreateDSPAMetaResourcesResponse() (response *CreateDSPAMetaResourcesResponse) {
    response = &CreateDSPAMetaResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDSPAMetaResources
// 添加用户云上资源列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateDSPAMetaResources(request *CreateDSPAMetaResourcesRequest) (response *CreateDSPAMetaResourcesResponse, err error) {
    return c.CreateDSPAMetaResourcesWithContext(context.Background(), request)
}

// CreateDSPAMetaResources
// 添加用户云上资源列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateDSPAMetaResourcesWithContext(ctx context.Context, request *CreateDSPAMetaResourcesRequest) (response *CreateDSPAMetaResourcesResponse, err error) {
    if request == nil {
        request = NewCreateDSPAMetaResourcesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "CreateDSPAMetaResources")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDSPAMetaResources require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDSPAMetaResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDSPASelfBuildMetaResourceRequest() (request *CreateDSPASelfBuildMetaResourceRequest) {
    request = &CreateDSPASelfBuildMetaResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "CreateDSPASelfBuildMetaResource")
    
    
    return
}

func NewCreateDSPASelfBuildMetaResourceResponse() (response *CreateDSPASelfBuildMetaResourceResponse) {
    response = &CreateDSPASelfBuildMetaResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDSPASelfBuildMetaResource
// 新建用户自建云资源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateDSPASelfBuildMetaResource(request *CreateDSPASelfBuildMetaResourceRequest) (response *CreateDSPASelfBuildMetaResourceResponse, err error) {
    return c.CreateDSPASelfBuildMetaResourceWithContext(context.Background(), request)
}

// CreateDSPASelfBuildMetaResource
// 新建用户自建云资源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateDSPASelfBuildMetaResourceWithContext(ctx context.Context, request *CreateDSPASelfBuildMetaResourceRequest) (response *CreateDSPASelfBuildMetaResourceResponse, err error) {
    if request == nil {
        request = NewCreateDSPASelfBuildMetaResourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "CreateDSPASelfBuildMetaResource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDSPASelfBuildMetaResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDSPASelfBuildMetaResourceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateIdentifyRuleAnotherNameRequest() (request *CreateIdentifyRuleAnotherNameRequest) {
    request = &CreateIdentifyRuleAnotherNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "CreateIdentifyRuleAnotherName")
    
    
    return
}

func NewCreateIdentifyRuleAnotherNameResponse() (response *CreateIdentifyRuleAnotherNameResponse) {
    response = &CreateIdentifyRuleAnotherNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateIdentifyRuleAnotherName
// 创建规则别名
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateIdentifyRuleAnotherName(request *CreateIdentifyRuleAnotherNameRequest) (response *CreateIdentifyRuleAnotherNameResponse, err error) {
    return c.CreateIdentifyRuleAnotherNameWithContext(context.Background(), request)
}

// CreateIdentifyRuleAnotherName
// 创建规则别名
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateIdentifyRuleAnotherNameWithContext(ctx context.Context, request *CreateIdentifyRuleAnotherNameRequest) (response *CreateIdentifyRuleAnotherNameResponse, err error) {
    if request == nil {
        request = NewCreateIdentifyRuleAnotherNameRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "CreateIdentifyRuleAnotherName")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateIdentifyRuleAnotherName require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateIdentifyRuleAnotherNameResponse()
    err = c.Send(request, response)
    return
}

func NewDecribeSuggestRiskLevelMatrixRequest() (request *DecribeSuggestRiskLevelMatrixRequest) {
    request = &DecribeSuggestRiskLevelMatrixRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DecribeSuggestRiskLevelMatrix")
    
    
    return
}

func NewDecribeSuggestRiskLevelMatrixResponse() (response *DecribeSuggestRiskLevelMatrixResponse) {
    response = &DecribeSuggestRiskLevelMatrixResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DecribeSuggestRiskLevelMatrix
// 风险等级的定义页面-创建风险等级的时候生成的一个默认的矩阵
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DecribeSuggestRiskLevelMatrix(request *DecribeSuggestRiskLevelMatrixRequest) (response *DecribeSuggestRiskLevelMatrixResponse, err error) {
    return c.DecribeSuggestRiskLevelMatrixWithContext(context.Background(), request)
}

// DecribeSuggestRiskLevelMatrix
// 风险等级的定义页面-创建风险等级的时候生成的一个默认的矩阵
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DecribeSuggestRiskLevelMatrixWithContext(ctx context.Context, request *DecribeSuggestRiskLevelMatrixRequest) (response *DecribeSuggestRiskLevelMatrixResponse, err error) {
    if request == nil {
        request = NewDecribeSuggestRiskLevelMatrixRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DecribeSuggestRiskLevelMatrix")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DecribeSuggestRiskLevelMatrix require credential")
    }

    request.SetContext(ctx)
    
    response = NewDecribeSuggestRiskLevelMatrixResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCosMetaResourceRequest() (request *DeleteCosMetaResourceRequest) {
    request = &DeleteCosMetaResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DeleteCosMetaResource")
    
    
    return
}

func NewDeleteCosMetaResourceResponse() (response *DeleteCosMetaResourceResponse) {
    response = &DeleteCosMetaResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCosMetaResource
// 本接口（DeleteCOSMetaData）用于删除COS元数据信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteCosMetaResource(request *DeleteCosMetaResourceRequest) (response *DeleteCosMetaResourceResponse, err error) {
    return c.DeleteCosMetaResourceWithContext(context.Background(), request)
}

// DeleteCosMetaResource
// 本接口（DeleteCOSMetaData）用于删除COS元数据信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteCosMetaResourceWithContext(ctx context.Context, request *DeleteCosMetaResourceRequest) (response *DeleteCosMetaResourceResponse, err error) {
    if request == nil {
        request = NewDeleteCosMetaResourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DeleteCosMetaResource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCosMetaResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCosMetaResourceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDSPAAssessmentTaskRequest() (request *DeleteDSPAAssessmentTaskRequest) {
    request = &DeleteDSPAAssessmentTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DeleteDSPAAssessmentTask")
    
    
    return
}

func NewDeleteDSPAAssessmentTaskResponse() (response *DeleteDSPAAssessmentTaskResponse) {
    response = &DeleteDSPAAssessmentTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDSPAAssessmentTask
// 删除DSPA风险评估任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteDSPAAssessmentTask(request *DeleteDSPAAssessmentTaskRequest) (response *DeleteDSPAAssessmentTaskResponse, err error) {
    return c.DeleteDSPAAssessmentTaskWithContext(context.Background(), request)
}

// DeleteDSPAAssessmentTask
// 删除DSPA风险评估任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteDSPAAssessmentTaskWithContext(ctx context.Context, request *DeleteDSPAAssessmentTaskRequest) (response *DeleteDSPAAssessmentTaskResponse, err error) {
    if request == nil {
        request = NewDeleteDSPAAssessmentTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DeleteDSPAAssessmentTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDSPAAssessmentTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDSPAAssessmentTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDSPACOSDiscoveryTaskRequest() (request *DeleteDSPACOSDiscoveryTaskRequest) {
    request = &DeleteDSPACOSDiscoveryTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DeleteDSPACOSDiscoveryTask")
    
    
    return
}

func NewDeleteDSPACOSDiscoveryTaskResponse() (response *DeleteDSPACOSDiscoveryTaskResponse) {
    response = &DeleteDSPACOSDiscoveryTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDSPACOSDiscoveryTask
// 删除COS分类分级任务，该接口只有在任务状态为以下几个状态值时才支持正确删除：
//
// 0 待扫描，
//
// 2 扫描终止， 
//
// 3 扫描成功，
//
// 4 扫描失败
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteDSPACOSDiscoveryTask(request *DeleteDSPACOSDiscoveryTaskRequest) (response *DeleteDSPACOSDiscoveryTaskResponse, err error) {
    return c.DeleteDSPACOSDiscoveryTaskWithContext(context.Background(), request)
}

// DeleteDSPACOSDiscoveryTask
// 删除COS分类分级任务，该接口只有在任务状态为以下几个状态值时才支持正确删除：
//
// 0 待扫描，
//
// 2 扫描终止， 
//
// 3 扫描成功，
//
// 4 扫描失败
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteDSPACOSDiscoveryTaskWithContext(ctx context.Context, request *DeleteDSPACOSDiscoveryTaskRequest) (response *DeleteDSPACOSDiscoveryTaskResponse, err error) {
    if request == nil {
        request = NewDeleteDSPACOSDiscoveryTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DeleteDSPACOSDiscoveryTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDSPACOSDiscoveryTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDSPACOSDiscoveryTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDSPACOSDiscoveryTaskResultRequest() (request *DeleteDSPACOSDiscoveryTaskResultRequest) {
    request = &DeleteDSPACOSDiscoveryTaskResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DeleteDSPACOSDiscoveryTaskResult")
    
    
    return
}

func NewDeleteDSPACOSDiscoveryTaskResultResponse() (response *DeleteDSPACOSDiscoveryTaskResultResponse) {
    response = &DeleteDSPACOSDiscoveryTaskResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDSPACOSDiscoveryTaskResult
// 删除COS分类分级任务结果
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteDSPACOSDiscoveryTaskResult(request *DeleteDSPACOSDiscoveryTaskResultRequest) (response *DeleteDSPACOSDiscoveryTaskResultResponse, err error) {
    return c.DeleteDSPACOSDiscoveryTaskResultWithContext(context.Background(), request)
}

// DeleteDSPACOSDiscoveryTaskResult
// 删除COS分类分级任务结果
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteDSPACOSDiscoveryTaskResultWithContext(ctx context.Context, request *DeleteDSPACOSDiscoveryTaskResultRequest) (response *DeleteDSPACOSDiscoveryTaskResultResponse, err error) {
    if request == nil {
        request = NewDeleteDSPACOSDiscoveryTaskResultRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DeleteDSPACOSDiscoveryTaskResult")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDSPACOSDiscoveryTaskResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDSPACOSDiscoveryTaskResultResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDSPADiscoveryTaskRequest() (request *DeleteDSPADiscoveryTaskRequest) {
    request = &DeleteDSPADiscoveryTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DeleteDSPADiscoveryTask")
    
    
    return
}

func NewDeleteDSPADiscoveryTaskResponse() (response *DeleteDSPADiscoveryTaskResponse) {
    response = &DeleteDSPADiscoveryTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDSPADiscoveryTask
// 删除分类分级识别任务，该接口只有在任务状态为以下几个状态值时才支持正确删除：
//
// 0 待扫描，
//
// 2 扫描终止， 
//
// 3 扫描成功，
//
// 4 扫描失败
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteDSPADiscoveryTask(request *DeleteDSPADiscoveryTaskRequest) (response *DeleteDSPADiscoveryTaskResponse, err error) {
    return c.DeleteDSPADiscoveryTaskWithContext(context.Background(), request)
}

// DeleteDSPADiscoveryTask
// 删除分类分级识别任务，该接口只有在任务状态为以下几个状态值时才支持正确删除：
//
// 0 待扫描，
//
// 2 扫描终止， 
//
// 3 扫描成功，
//
// 4 扫描失败
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteDSPADiscoveryTaskWithContext(ctx context.Context, request *DeleteDSPADiscoveryTaskRequest) (response *DeleteDSPADiscoveryTaskResponse, err error) {
    if request == nil {
        request = NewDeleteDSPADiscoveryTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DeleteDSPADiscoveryTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDSPADiscoveryTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDSPADiscoveryTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDSPADiscoveryTaskResultRequest() (request *DeleteDSPADiscoveryTaskResultRequest) {
    request = &DeleteDSPADiscoveryTaskResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DeleteDSPADiscoveryTaskResult")
    
    
    return
}

func NewDeleteDSPADiscoveryTaskResultResponse() (response *DeleteDSPADiscoveryTaskResultResponse) {
    response = &DeleteDSPADiscoveryTaskResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDSPADiscoveryTaskResult
// 删除分类分级识别任务结果
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteDSPADiscoveryTaskResult(request *DeleteDSPADiscoveryTaskResultRequest) (response *DeleteDSPADiscoveryTaskResultResponse, err error) {
    return c.DeleteDSPADiscoveryTaskResultWithContext(context.Background(), request)
}

// DeleteDSPADiscoveryTaskResult
// 删除分类分级识别任务结果
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteDSPADiscoveryTaskResultWithContext(ctx context.Context, request *DeleteDSPADiscoveryTaskResultRequest) (response *DeleteDSPADiscoveryTaskResultResponse, err error) {
    if request == nil {
        request = NewDeleteDSPADiscoveryTaskResultRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DeleteDSPADiscoveryTaskResult")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDSPADiscoveryTaskResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDSPADiscoveryTaskResultResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDSPAMetaResourceRequest() (request *DeleteDSPAMetaResourceRequest) {
    request = &DeleteDSPAMetaResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DeleteDSPAMetaResource")
    
    
    return
}

func NewDeleteDSPAMetaResourceResponse() (response *DeleteDSPAMetaResourceResponse) {
    response = &DeleteDSPAMetaResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDSPAMetaResource
// 删除用户云资源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteDSPAMetaResource(request *DeleteDSPAMetaResourceRequest) (response *DeleteDSPAMetaResourceResponse, err error) {
    return c.DeleteDSPAMetaResourceWithContext(context.Background(), request)
}

// DeleteDSPAMetaResource
// 删除用户云资源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteDSPAMetaResourceWithContext(ctx context.Context, request *DeleteDSPAMetaResourceRequest) (response *DeleteDSPAMetaResourceResponse, err error) {
    if request == nil {
        request = NewDeleteDSPAMetaResourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DeleteDSPAMetaResource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDSPAMetaResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDSPAMetaResourceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetDetailDataExportResultRequest() (request *DescribeAssetDetailDataExportResultRequest) {
    request = &DescribeAssetDetailDataExportResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeAssetDetailDataExportResult")
    
    
    return
}

func NewDescribeAssetDetailDataExportResultResponse() (response *DescribeAssetDetailDataExportResultResponse) {
    response = &DescribeAssetDetailDataExportResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetDetailDataExportResult
// 查询敏感数据导出结果URL
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetDetailDataExportResult(request *DescribeAssetDetailDataExportResultRequest) (response *DescribeAssetDetailDataExportResultResponse, err error) {
    return c.DescribeAssetDetailDataExportResultWithContext(context.Background(), request)
}

// DescribeAssetDetailDataExportResult
// 查询敏感数据导出结果URL
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetDetailDataExportResultWithContext(ctx context.Context, request *DescribeAssetDetailDataExportResultRequest) (response *DescribeAssetDetailDataExportResultResponse, err error) {
    if request == nil {
        request = NewDescribeAssetDetailDataExportResultRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeAssetDetailDataExportResult")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetDetailDataExportResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetDetailDataExportResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetOverviewRequest() (request *DescribeAssetOverviewRequest) {
    request = &DescribeAssetOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeAssetOverview")
    
    
    return
}

func NewDescribeAssetOverviewResponse() (response *DescribeAssetOverviewResponse) {
    response = &DescribeAssetOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAssetOverview
// 数据资产报告页面-查询数据资产概览接口（包括数据库资产详情和存储资产详情）
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetOverview(request *DescribeAssetOverviewRequest) (response *DescribeAssetOverviewResponse, err error) {
    return c.DescribeAssetOverviewWithContext(context.Background(), request)
}

// DescribeAssetOverview
// 数据资产报告页面-查询数据资产概览接口（包括数据库资产详情和存储资产详情）
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAssetOverviewWithContext(ctx context.Context, request *DescribeAssetOverviewRequest) (response *DescribeAssetOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeAssetOverviewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeAssetOverview")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAssetOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAssetOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBindDBListRequest() (request *DescribeBindDBListRequest) {
    request = &DescribeBindDBListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeBindDBList")
    
    
    return
}

func NewDescribeBindDBListResponse() (response *DescribeBindDBListResponse) {
    response = &DescribeBindDBListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBindDBList
// 查询DB绑定的列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeBindDBList(request *DescribeBindDBListRequest) (response *DescribeBindDBListResponse, err error) {
    return c.DescribeBindDBListWithContext(context.Background(), request)
}

// DescribeBindDBList
// 查询DB绑定的列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeBindDBListWithContext(ctx context.Context, request *DescribeBindDBListRequest) (response *DescribeBindDBListResponse, err error) {
    if request == nil {
        request = NewDescribeBindDBListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeBindDBList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBindDBList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBindDBListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCOSAssetSensitiveDistributionRequest() (request *DescribeCOSAssetSensitiveDistributionRequest) {
    request = &DescribeCOSAssetSensitiveDistributionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeCOSAssetSensitiveDistribution")
    
    
    return
}

func NewDescribeCOSAssetSensitiveDistributionResponse() (response *DescribeCOSAssetSensitiveDistributionResponse) {
    response = &DescribeCOSAssetSensitiveDistributionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCOSAssetSensitiveDistribution
// 数据资产报告-查询cos的资产分布详情接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeCOSAssetSensitiveDistribution(request *DescribeCOSAssetSensitiveDistributionRequest) (response *DescribeCOSAssetSensitiveDistributionResponse, err error) {
    return c.DescribeCOSAssetSensitiveDistributionWithContext(context.Background(), request)
}

// DescribeCOSAssetSensitiveDistribution
// 数据资产报告-查询cos的资产分布详情接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeCOSAssetSensitiveDistributionWithContext(ctx context.Context, request *DescribeCOSAssetSensitiveDistributionRequest) (response *DescribeCOSAssetSensitiveDistributionResponse, err error) {
    if request == nil {
        request = NewDescribeCOSAssetSensitiveDistributionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeCOSAssetSensitiveDistribution")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCOSAssetSensitiveDistribution require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCOSAssetSensitiveDistributionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPAAssessmentHighRiskTop10OverviewRequest() (request *DescribeDSPAAssessmentHighRiskTop10OverviewRequest) {
    request = &DescribeDSPAAssessmentHighRiskTop10OverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPAAssessmentHighRiskTop10Overview")
    
    
    return
}

func NewDescribeDSPAAssessmentHighRiskTop10OverviewResponse() (response *DescribeDSPAAssessmentHighRiskTop10OverviewResponse) {
    response = &DescribeDSPAAssessmentHighRiskTop10OverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPAAssessmentHighRiskTop10Overview
// 查询高风险资产的top10
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDSPAAssessmentHighRiskTop10Overview(request *DescribeDSPAAssessmentHighRiskTop10OverviewRequest) (response *DescribeDSPAAssessmentHighRiskTop10OverviewResponse, err error) {
    return c.DescribeDSPAAssessmentHighRiskTop10OverviewWithContext(context.Background(), request)
}

// DescribeDSPAAssessmentHighRiskTop10Overview
// 查询高风险资产的top10
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDSPAAssessmentHighRiskTop10OverviewWithContext(ctx context.Context, request *DescribeDSPAAssessmentHighRiskTop10OverviewRequest) (response *DescribeDSPAAssessmentHighRiskTop10OverviewResponse, err error) {
    if request == nil {
        request = NewDescribeDSPAAssessmentHighRiskTop10OverviewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPAAssessmentHighRiskTop10Overview")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPAAssessmentHighRiskTop10Overview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPAAssessmentHighRiskTop10OverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPAAssessmentLatestRiskDetailInfoRequest() (request *DescribeDSPAAssessmentLatestRiskDetailInfoRequest) {
    request = &DescribeDSPAAssessmentLatestRiskDetailInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPAAssessmentLatestRiskDetailInfo")
    
    
    return
}

func NewDescribeDSPAAssessmentLatestRiskDetailInfoResponse() (response *DescribeDSPAAssessmentLatestRiskDetailInfoResponse) {
    response = &DescribeDSPAAssessmentLatestRiskDetailInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPAAssessmentLatestRiskDetailInfo
// 查询最新风险项详情数据
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDSPAAssessmentLatestRiskDetailInfo(request *DescribeDSPAAssessmentLatestRiskDetailInfoRequest) (response *DescribeDSPAAssessmentLatestRiskDetailInfoResponse, err error) {
    return c.DescribeDSPAAssessmentLatestRiskDetailInfoWithContext(context.Background(), request)
}

// DescribeDSPAAssessmentLatestRiskDetailInfo
// 查询最新风险项详情数据
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDSPAAssessmentLatestRiskDetailInfoWithContext(ctx context.Context, request *DescribeDSPAAssessmentLatestRiskDetailInfoRequest) (response *DescribeDSPAAssessmentLatestRiskDetailInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDSPAAssessmentLatestRiskDetailInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPAAssessmentLatestRiskDetailInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPAAssessmentLatestRiskDetailInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPAAssessmentLatestRiskDetailInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPAAssessmentLatestRiskListRequest() (request *DescribeDSPAAssessmentLatestRiskListRequest) {
    request = &DescribeDSPAAssessmentLatestRiskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPAAssessmentLatestRiskList")
    
    
    return
}

func NewDescribeDSPAAssessmentLatestRiskListResponse() (response *DescribeDSPAAssessmentLatestRiskListResponse) {
    response = &DescribeDSPAAssessmentLatestRiskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPAAssessmentLatestRiskList
// 查询最新的风险详情列表数据
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDSPAAssessmentLatestRiskList(request *DescribeDSPAAssessmentLatestRiskListRequest) (response *DescribeDSPAAssessmentLatestRiskListResponse, err error) {
    return c.DescribeDSPAAssessmentLatestRiskListWithContext(context.Background(), request)
}

// DescribeDSPAAssessmentLatestRiskList
// 查询最新的风险详情列表数据
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDSPAAssessmentLatestRiskListWithContext(ctx context.Context, request *DescribeDSPAAssessmentLatestRiskListRequest) (response *DescribeDSPAAssessmentLatestRiskListResponse, err error) {
    if request == nil {
        request = NewDescribeDSPAAssessmentLatestRiskListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPAAssessmentLatestRiskList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPAAssessmentLatestRiskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPAAssessmentLatestRiskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPAAssessmentNewDiscoveredRiskOverviewRequest() (request *DescribeDSPAAssessmentNewDiscoveredRiskOverviewRequest) {
    request = &DescribeDSPAAssessmentNewDiscoveredRiskOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPAAssessmentNewDiscoveredRiskOverview")
    
    
    return
}

func NewDescribeDSPAAssessmentNewDiscoveredRiskOverviewResponse() (response *DescribeDSPAAssessmentNewDiscoveredRiskOverviewResponse) {
    response = &DescribeDSPAAssessmentNewDiscoveredRiskOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPAAssessmentNewDiscoveredRiskOverview
// 风险概览-查询新发现风险统计数
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDSPAAssessmentNewDiscoveredRiskOverview(request *DescribeDSPAAssessmentNewDiscoveredRiskOverviewRequest) (response *DescribeDSPAAssessmentNewDiscoveredRiskOverviewResponse, err error) {
    return c.DescribeDSPAAssessmentNewDiscoveredRiskOverviewWithContext(context.Background(), request)
}

// DescribeDSPAAssessmentNewDiscoveredRiskOverview
// 风险概览-查询新发现风险统计数
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDSPAAssessmentNewDiscoveredRiskOverviewWithContext(ctx context.Context, request *DescribeDSPAAssessmentNewDiscoveredRiskOverviewRequest) (response *DescribeDSPAAssessmentNewDiscoveredRiskOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeDSPAAssessmentNewDiscoveredRiskOverviewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPAAssessmentNewDiscoveredRiskOverview")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPAAssessmentNewDiscoveredRiskOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPAAssessmentNewDiscoveredRiskOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPAAssessmentPendingRiskOverviewRequest() (request *DescribeDSPAAssessmentPendingRiskOverviewRequest) {
    request = &DescribeDSPAAssessmentPendingRiskOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPAAssessmentPendingRiskOverview")
    
    
    return
}

func NewDescribeDSPAAssessmentPendingRiskOverviewResponse() (response *DescribeDSPAAssessmentPendingRiskOverviewResponse) {
    response = &DescribeDSPAAssessmentPendingRiskOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPAAssessmentPendingRiskOverview
// 风险概览-查询待处理风险统计数
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDSPAAssessmentPendingRiskOverview(request *DescribeDSPAAssessmentPendingRiskOverviewRequest) (response *DescribeDSPAAssessmentPendingRiskOverviewResponse, err error) {
    return c.DescribeDSPAAssessmentPendingRiskOverviewWithContext(context.Background(), request)
}

// DescribeDSPAAssessmentPendingRiskOverview
// 风险概览-查询待处理风险统计数
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDSPAAssessmentPendingRiskOverviewWithContext(ctx context.Context, request *DescribeDSPAAssessmentPendingRiskOverviewRequest) (response *DescribeDSPAAssessmentPendingRiskOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeDSPAAssessmentPendingRiskOverviewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPAAssessmentPendingRiskOverview")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPAAssessmentPendingRiskOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPAAssessmentPendingRiskOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPAAssessmentProcessingRiskOverviewRequest() (request *DescribeDSPAAssessmentProcessingRiskOverviewRequest) {
    request = &DescribeDSPAAssessmentProcessingRiskOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPAAssessmentProcessingRiskOverview")
    
    
    return
}

func NewDescribeDSPAAssessmentProcessingRiskOverviewResponse() (response *DescribeDSPAAssessmentProcessingRiskOverviewResponse) {
    response = &DescribeDSPAAssessmentProcessingRiskOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPAAssessmentProcessingRiskOverview
// 风险概览-查询处理中风险统计数
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDSPAAssessmentProcessingRiskOverview(request *DescribeDSPAAssessmentProcessingRiskOverviewRequest) (response *DescribeDSPAAssessmentProcessingRiskOverviewResponse, err error) {
    return c.DescribeDSPAAssessmentProcessingRiskOverviewWithContext(context.Background(), request)
}

// DescribeDSPAAssessmentProcessingRiskOverview
// 风险概览-查询处理中风险统计数
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDSPAAssessmentProcessingRiskOverviewWithContext(ctx context.Context, request *DescribeDSPAAssessmentProcessingRiskOverviewRequest) (response *DescribeDSPAAssessmentProcessingRiskOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeDSPAAssessmentProcessingRiskOverviewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPAAssessmentProcessingRiskOverview")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPAAssessmentProcessingRiskOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPAAssessmentProcessingRiskOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPAAssessmentRiskAmountOverviewRequest() (request *DescribeDSPAAssessmentRiskAmountOverviewRequest) {
    request = &DescribeDSPAAssessmentRiskAmountOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPAAssessmentRiskAmountOverview")
    
    
    return
}

func NewDescribeDSPAAssessmentRiskAmountOverviewResponse() (response *DescribeDSPAAssessmentRiskAmountOverviewResponse) {
    response = &DescribeDSPAAssessmentRiskAmountOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPAAssessmentRiskAmountOverview
// 风险概览页风险数量和受影响资产数概览统计
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDSPAAssessmentRiskAmountOverview(request *DescribeDSPAAssessmentRiskAmountOverviewRequest) (response *DescribeDSPAAssessmentRiskAmountOverviewResponse, err error) {
    return c.DescribeDSPAAssessmentRiskAmountOverviewWithContext(context.Background(), request)
}

// DescribeDSPAAssessmentRiskAmountOverview
// 风险概览页风险数量和受影响资产数概览统计
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDSPAAssessmentRiskAmountOverviewWithContext(ctx context.Context, request *DescribeDSPAAssessmentRiskAmountOverviewRequest) (response *DescribeDSPAAssessmentRiskAmountOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeDSPAAssessmentRiskAmountOverviewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPAAssessmentRiskAmountOverview")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPAAssessmentRiskAmountOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPAAssessmentRiskAmountOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPAAssessmentRiskDatasourceTop5Request() (request *DescribeDSPAAssessmentRiskDatasourceTop5Request) {
    request = &DescribeDSPAAssessmentRiskDatasourceTop5Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPAAssessmentRiskDatasourceTop5")
    
    
    return
}

func NewDescribeDSPAAssessmentRiskDatasourceTop5Response() (response *DescribeDSPAAssessmentRiskDatasourceTop5Response) {
    response = &DescribeDSPAAssessmentRiskDatasourceTop5Response{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPAAssessmentRiskDatasourceTop5
// 受影响资产TOP5统计
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDSPAAssessmentRiskDatasourceTop5(request *DescribeDSPAAssessmentRiskDatasourceTop5Request) (response *DescribeDSPAAssessmentRiskDatasourceTop5Response, err error) {
    return c.DescribeDSPAAssessmentRiskDatasourceTop5WithContext(context.Background(), request)
}

// DescribeDSPAAssessmentRiskDatasourceTop5
// 受影响资产TOP5统计
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDSPAAssessmentRiskDatasourceTop5WithContext(ctx context.Context, request *DescribeDSPAAssessmentRiskDatasourceTop5Request) (response *DescribeDSPAAssessmentRiskDatasourceTop5Response, err error) {
    if request == nil {
        request = NewDescribeDSPAAssessmentRiskDatasourceTop5Request()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPAAssessmentRiskDatasourceTop5")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPAAssessmentRiskDatasourceTop5 require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPAAssessmentRiskDatasourceTop5Response()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPAAssessmentRiskDealedOverviewRequest() (request *DescribeDSPAAssessmentRiskDealedOverviewRequest) {
    request = &DescribeDSPAAssessmentRiskDealedOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPAAssessmentRiskDealedOverview")
    
    
    return
}

func NewDescribeDSPAAssessmentRiskDealedOverviewResponse() (response *DescribeDSPAAssessmentRiskDealedOverviewResponse) {
    response = &DescribeDSPAAssessmentRiskDealedOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPAAssessmentRiskDealedOverview
// 遗留待处理&昨日完成风险处置概览统计
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDSPAAssessmentRiskDealedOverview(request *DescribeDSPAAssessmentRiskDealedOverviewRequest) (response *DescribeDSPAAssessmentRiskDealedOverviewResponse, err error) {
    return c.DescribeDSPAAssessmentRiskDealedOverviewWithContext(context.Background(), request)
}

// DescribeDSPAAssessmentRiskDealedOverview
// 遗留待处理&昨日完成风险处置概览统计
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDSPAAssessmentRiskDealedOverviewWithContext(ctx context.Context, request *DescribeDSPAAssessmentRiskDealedOverviewRequest) (response *DescribeDSPAAssessmentRiskDealedOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeDSPAAssessmentRiskDealedOverviewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPAAssessmentRiskDealedOverview")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPAAssessmentRiskDealedOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPAAssessmentRiskDealedOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPAAssessmentRiskDealedTrendRequest() (request *DescribeDSPAAssessmentRiskDealedTrendRequest) {
    request = &DescribeDSPAAssessmentRiskDealedTrendRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPAAssessmentRiskDealedTrend")
    
    
    return
}

func NewDescribeDSPAAssessmentRiskDealedTrendResponse() (response *DescribeDSPAAssessmentRiskDealedTrendResponse) {
    response = &DescribeDSPAAssessmentRiskDealedTrendResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPAAssessmentRiskDealedTrend
// 风险项处理趋势统计
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDSPAAssessmentRiskDealedTrend(request *DescribeDSPAAssessmentRiskDealedTrendRequest) (response *DescribeDSPAAssessmentRiskDealedTrendResponse, err error) {
    return c.DescribeDSPAAssessmentRiskDealedTrendWithContext(context.Background(), request)
}

// DescribeDSPAAssessmentRiskDealedTrend
// 风险项处理趋势统计
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDSPAAssessmentRiskDealedTrendWithContext(ctx context.Context, request *DescribeDSPAAssessmentRiskDealedTrendRequest) (response *DescribeDSPAAssessmentRiskDealedTrendResponse, err error) {
    if request == nil {
        request = NewDescribeDSPAAssessmentRiskDealedTrendRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPAAssessmentRiskDealedTrend")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPAAssessmentRiskDealedTrend require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPAAssessmentRiskDealedTrendResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPAAssessmentRiskDistributionOverviewRequest() (request *DescribeDSPAAssessmentRiskDistributionOverviewRequest) {
    request = &DescribeDSPAAssessmentRiskDistributionOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPAAssessmentRiskDistributionOverview")
    
    
    return
}

func NewDescribeDSPAAssessmentRiskDistributionOverviewResponse() (response *DescribeDSPAAssessmentRiskDistributionOverviewResponse) {
    response = &DescribeDSPAAssessmentRiskDistributionOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPAAssessmentRiskDistributionOverview
// 查询风险分布数据，包含风险类型分布，风险详情分布，风险资产的分布
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDSPAAssessmentRiskDistributionOverview(request *DescribeDSPAAssessmentRiskDistributionOverviewRequest) (response *DescribeDSPAAssessmentRiskDistributionOverviewResponse, err error) {
    return c.DescribeDSPAAssessmentRiskDistributionOverviewWithContext(context.Background(), request)
}

// DescribeDSPAAssessmentRiskDistributionOverview
// 查询风险分布数据，包含风险类型分布，风险详情分布，风险资产的分布
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDSPAAssessmentRiskDistributionOverviewWithContext(ctx context.Context, request *DescribeDSPAAssessmentRiskDistributionOverviewRequest) (response *DescribeDSPAAssessmentRiskDistributionOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeDSPAAssessmentRiskDistributionOverviewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPAAssessmentRiskDistributionOverview")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPAAssessmentRiskDistributionOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPAAssessmentRiskDistributionOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPAAssessmentRiskItemTop5Request() (request *DescribeDSPAAssessmentRiskItemTop5Request) {
    request = &DescribeDSPAAssessmentRiskItemTop5Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPAAssessmentRiskItemTop5")
    
    
    return
}

func NewDescribeDSPAAssessmentRiskItemTop5Response() (response *DescribeDSPAAssessmentRiskItemTop5Response) {
    response = &DescribeDSPAAssessmentRiskItemTop5Response{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPAAssessmentRiskItemTop5
// 风险分类TOP5统计
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDSPAAssessmentRiskItemTop5(request *DescribeDSPAAssessmentRiskItemTop5Request) (response *DescribeDSPAAssessmentRiskItemTop5Response, err error) {
    return c.DescribeDSPAAssessmentRiskItemTop5WithContext(context.Background(), request)
}

// DescribeDSPAAssessmentRiskItemTop5
// 风险分类TOP5统计
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDSPAAssessmentRiskItemTop5WithContext(ctx context.Context, request *DescribeDSPAAssessmentRiskItemTop5Request) (response *DescribeDSPAAssessmentRiskItemTop5Response, err error) {
    if request == nil {
        request = NewDescribeDSPAAssessmentRiskItemTop5Request()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPAAssessmentRiskItemTop5")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPAAssessmentRiskItemTop5 require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPAAssessmentRiskItemTop5Response()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPAAssessmentRiskLevelDetailRequest() (request *DescribeDSPAAssessmentRiskLevelDetailRequest) {
    request = &DescribeDSPAAssessmentRiskLevelDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPAAssessmentRiskLevelDetail")
    
    
    return
}

func NewDescribeDSPAAssessmentRiskLevelDetailResponse() (response *DescribeDSPAAssessmentRiskLevelDetailResponse) {
    response = &DescribeDSPAAssessmentRiskLevelDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPAAssessmentRiskLevelDetail
// 风险项页面----查询风险等级的详情数据
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDSPAAssessmentRiskLevelDetail(request *DescribeDSPAAssessmentRiskLevelDetailRequest) (response *DescribeDSPAAssessmentRiskLevelDetailResponse, err error) {
    return c.DescribeDSPAAssessmentRiskLevelDetailWithContext(context.Background(), request)
}

// DescribeDSPAAssessmentRiskLevelDetail
// 风险项页面----查询风险等级的详情数据
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDSPAAssessmentRiskLevelDetailWithContext(ctx context.Context, request *DescribeDSPAAssessmentRiskLevelDetailRequest) (response *DescribeDSPAAssessmentRiskLevelDetailResponse, err error) {
    if request == nil {
        request = NewDescribeDSPAAssessmentRiskLevelDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPAAssessmentRiskLevelDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPAAssessmentRiskLevelDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPAAssessmentRiskLevelDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPAAssessmentRiskLevelListRequest() (request *DescribeDSPAAssessmentRiskLevelListRequest) {
    request = &DescribeDSPAAssessmentRiskLevelListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPAAssessmentRiskLevelList")
    
    
    return
}

func NewDescribeDSPAAssessmentRiskLevelListResponse() (response *DescribeDSPAAssessmentRiskLevelListResponse) {
    response = &DescribeDSPAAssessmentRiskLevelListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPAAssessmentRiskLevelList
// 风险项页面----查询风险等级的列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDSPAAssessmentRiskLevelList(request *DescribeDSPAAssessmentRiskLevelListRequest) (response *DescribeDSPAAssessmentRiskLevelListResponse, err error) {
    return c.DescribeDSPAAssessmentRiskLevelListWithContext(context.Background(), request)
}

// DescribeDSPAAssessmentRiskLevelList
// 风险项页面----查询风险等级的列表
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDSPAAssessmentRiskLevelListWithContext(ctx context.Context, request *DescribeDSPAAssessmentRiskLevelListRequest) (response *DescribeDSPAAssessmentRiskLevelListResponse, err error) {
    if request == nil {
        request = NewDescribeDSPAAssessmentRiskLevelListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPAAssessmentRiskLevelList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPAAssessmentRiskLevelList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPAAssessmentRiskLevelListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPAAssessmentRiskLevelTrendRequest() (request *DescribeDSPAAssessmentRiskLevelTrendRequest) {
    request = &DescribeDSPAAssessmentRiskLevelTrendRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPAAssessmentRiskLevelTrend")
    
    
    return
}

func NewDescribeDSPAAssessmentRiskLevelTrendResponse() (response *DescribeDSPAAssessmentRiskLevelTrendResponse) {
    response = &DescribeDSPAAssessmentRiskLevelTrendResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPAAssessmentRiskLevelTrend
// 风险级别趋势统计
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDSPAAssessmentRiskLevelTrend(request *DescribeDSPAAssessmentRiskLevelTrendRequest) (response *DescribeDSPAAssessmentRiskLevelTrendResponse, err error) {
    return c.DescribeDSPAAssessmentRiskLevelTrendWithContext(context.Background(), request)
}

// DescribeDSPAAssessmentRiskLevelTrend
// 风险级别趋势统计
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDSPAAssessmentRiskLevelTrendWithContext(ctx context.Context, request *DescribeDSPAAssessmentRiskLevelTrendRequest) (response *DescribeDSPAAssessmentRiskLevelTrendResponse, err error) {
    if request == nil {
        request = NewDescribeDSPAAssessmentRiskLevelTrendRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPAAssessmentRiskLevelTrend")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPAAssessmentRiskLevelTrend require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPAAssessmentRiskLevelTrendResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPAAssessmentRiskOverviewRequest() (request *DescribeDSPAAssessmentRiskOverviewRequest) {
    request = &DescribeDSPAAssessmentRiskOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPAAssessmentRiskOverview")
    
    
    return
}

func NewDescribeDSPAAssessmentRiskOverviewResponse() (response *DescribeDSPAAssessmentRiskOverviewResponse) {
    response = &DescribeDSPAAssessmentRiskOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPAAssessmentRiskOverview
// 风险数量概览统计
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDSPAAssessmentRiskOverview(request *DescribeDSPAAssessmentRiskOverviewRequest) (response *DescribeDSPAAssessmentRiskOverviewResponse, err error) {
    return c.DescribeDSPAAssessmentRiskOverviewWithContext(context.Background(), request)
}

// DescribeDSPAAssessmentRiskOverview
// 风险数量概览统计
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDSPAAssessmentRiskOverviewWithContext(ctx context.Context, request *DescribeDSPAAssessmentRiskOverviewRequest) (response *DescribeDSPAAssessmentRiskOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeDSPAAssessmentRiskOverviewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPAAssessmentRiskOverview")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPAAssessmentRiskOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPAAssessmentRiskOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPAAssessmentRiskProcessHistoryRequest() (request *DescribeDSPAAssessmentRiskProcessHistoryRequest) {
    request = &DescribeDSPAAssessmentRiskProcessHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPAAssessmentRiskProcessHistory")
    
    
    return
}

func NewDescribeDSPAAssessmentRiskProcessHistoryResponse() (response *DescribeDSPAAssessmentRiskProcessHistoryResponse) {
    response = &DescribeDSPAAssessmentRiskProcessHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPAAssessmentRiskProcessHistory
// 查询风险的处理历史
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDSPAAssessmentRiskProcessHistory(request *DescribeDSPAAssessmentRiskProcessHistoryRequest) (response *DescribeDSPAAssessmentRiskProcessHistoryResponse, err error) {
    return c.DescribeDSPAAssessmentRiskProcessHistoryWithContext(context.Background(), request)
}

// DescribeDSPAAssessmentRiskProcessHistory
// 查询风险的处理历史
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDSPAAssessmentRiskProcessHistoryWithContext(ctx context.Context, request *DescribeDSPAAssessmentRiskProcessHistoryRequest) (response *DescribeDSPAAssessmentRiskProcessHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeDSPAAssessmentRiskProcessHistoryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPAAssessmentRiskProcessHistory")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPAAssessmentRiskProcessHistory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPAAssessmentRiskProcessHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPAAssessmentRiskSideDistributedRequest() (request *DescribeDSPAAssessmentRiskSideDistributedRequest) {
    request = &DescribeDSPAAssessmentRiskSideDistributedRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPAAssessmentRiskSideDistributed")
    
    
    return
}

func NewDescribeDSPAAssessmentRiskSideDistributedResponse() (response *DescribeDSPAAssessmentRiskSideDistributedResponse) {
    response = &DescribeDSPAAssessmentRiskSideDistributedResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPAAssessmentRiskSideDistributed
// 风险评估概览页，查询风险面的分布
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDSPAAssessmentRiskSideDistributed(request *DescribeDSPAAssessmentRiskSideDistributedRequest) (response *DescribeDSPAAssessmentRiskSideDistributedResponse, err error) {
    return c.DescribeDSPAAssessmentRiskSideDistributedWithContext(context.Background(), request)
}

// DescribeDSPAAssessmentRiskSideDistributed
// 风险评估概览页，查询风险面的分布
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDSPAAssessmentRiskSideDistributedWithContext(ctx context.Context, request *DescribeDSPAAssessmentRiskSideDistributedRequest) (response *DescribeDSPAAssessmentRiskSideDistributedResponse, err error) {
    if request == nil {
        request = NewDescribeDSPAAssessmentRiskSideDistributedRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPAAssessmentRiskSideDistributed")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPAAssessmentRiskSideDistributed require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPAAssessmentRiskSideDistributedResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPAAssessmentRiskSideListRequest() (request *DescribeDSPAAssessmentRiskSideListRequest) {
    request = &DescribeDSPAAssessmentRiskSideListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPAAssessmentRiskSideList")
    
    
    return
}

func NewDescribeDSPAAssessmentRiskSideListResponse() (response *DescribeDSPAAssessmentRiskSideListResponse) {
    response = &DescribeDSPAAssessmentRiskSideListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPAAssessmentRiskSideList
// 风险评估概览页，查询风险面的分布
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDSPAAssessmentRiskSideList(request *DescribeDSPAAssessmentRiskSideListRequest) (response *DescribeDSPAAssessmentRiskSideListResponse, err error) {
    return c.DescribeDSPAAssessmentRiskSideListWithContext(context.Background(), request)
}

// DescribeDSPAAssessmentRiskSideList
// 风险评估概览页，查询风险面的分布
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDSPAAssessmentRiskSideListWithContext(ctx context.Context, request *DescribeDSPAAssessmentRiskSideListRequest) (response *DescribeDSPAAssessmentRiskSideListResponse, err error) {
    if request == nil {
        request = NewDescribeDSPAAssessmentRiskSideListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPAAssessmentRiskSideList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPAAssessmentRiskSideList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPAAssessmentRiskSideListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPAAssessmentRiskTemplateDetailRequest() (request *DescribeDSPAAssessmentRiskTemplateDetailRequest) {
    request = &DescribeDSPAAssessmentRiskTemplateDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPAAssessmentRiskTemplateDetail")
    
    
    return
}

func NewDescribeDSPAAssessmentRiskTemplateDetailResponse() (response *DescribeDSPAAssessmentRiskTemplateDetailResponse) {
    response = &DescribeDSPAAssessmentRiskTemplateDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPAAssessmentRiskTemplateDetail
// 风险项页面--查看评估模板详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDSPAAssessmentRiskTemplateDetail(request *DescribeDSPAAssessmentRiskTemplateDetailRequest) (response *DescribeDSPAAssessmentRiskTemplateDetailResponse, err error) {
    return c.DescribeDSPAAssessmentRiskTemplateDetailWithContext(context.Background(), request)
}

// DescribeDSPAAssessmentRiskTemplateDetail
// 风险项页面--查看评估模板详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDSPAAssessmentRiskTemplateDetailWithContext(ctx context.Context, request *DescribeDSPAAssessmentRiskTemplateDetailRequest) (response *DescribeDSPAAssessmentRiskTemplateDetailResponse, err error) {
    if request == nil {
        request = NewDescribeDSPAAssessmentRiskTemplateDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPAAssessmentRiskTemplateDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPAAssessmentRiskTemplateDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPAAssessmentRiskTemplateDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPAAssessmentRiskTemplateVulnerableListRequest() (request *DescribeDSPAAssessmentRiskTemplateVulnerableListRequest) {
    request = &DescribeDSPAAssessmentRiskTemplateVulnerableListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPAAssessmentRiskTemplateVulnerableList")
    
    
    return
}

func NewDescribeDSPAAssessmentRiskTemplateVulnerableListResponse() (response *DescribeDSPAAssessmentRiskTemplateVulnerableListResponse) {
    response = &DescribeDSPAAssessmentRiskTemplateVulnerableListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPAAssessmentRiskTemplateVulnerableList
// 风险模板页面--查询风险模板中的脆弱项配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDSPAAssessmentRiskTemplateVulnerableList(request *DescribeDSPAAssessmentRiskTemplateVulnerableListRequest) (response *DescribeDSPAAssessmentRiskTemplateVulnerableListResponse, err error) {
    return c.DescribeDSPAAssessmentRiskTemplateVulnerableListWithContext(context.Background(), request)
}

// DescribeDSPAAssessmentRiskTemplateVulnerableList
// 风险模板页面--查询风险模板中的脆弱项配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDSPAAssessmentRiskTemplateVulnerableListWithContext(ctx context.Context, request *DescribeDSPAAssessmentRiskTemplateVulnerableListRequest) (response *DescribeDSPAAssessmentRiskTemplateVulnerableListResponse, err error) {
    if request == nil {
        request = NewDescribeDSPAAssessmentRiskTemplateVulnerableListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPAAssessmentRiskTemplateVulnerableList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPAAssessmentRiskTemplateVulnerableList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPAAssessmentRiskTemplateVulnerableListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPAAssessmentRisksRequest() (request *DescribeDSPAAssessmentRisksRequest) {
    request = &DescribeDSPAAssessmentRisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPAAssessmentRisks")
    
    
    return
}

func NewDescribeDSPAAssessmentRisksResponse() (response *DescribeDSPAAssessmentRisksResponse) {
    response = &DescribeDSPAAssessmentRisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPAAssessmentRisks
// 获取DSPA评估风险项列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDSPAAssessmentRisks(request *DescribeDSPAAssessmentRisksRequest) (response *DescribeDSPAAssessmentRisksResponse, err error) {
    return c.DescribeDSPAAssessmentRisksWithContext(context.Background(), request)
}

// DescribeDSPAAssessmentRisks
// 获取DSPA评估风险项列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDSPAAssessmentRisksWithContext(ctx context.Context, request *DescribeDSPAAssessmentRisksRequest) (response *DescribeDSPAAssessmentRisksResponse, err error) {
    if request == nil {
        request = NewDescribeDSPAAssessmentRisksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPAAssessmentRisks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPAAssessmentRisks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPAAssessmentRisksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPAAssessmentTasksRequest() (request *DescribeDSPAAssessmentTasksRequest) {
    request = &DescribeDSPAAssessmentTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPAAssessmentTasks")
    
    
    return
}

func NewDescribeDSPAAssessmentTasksResponse() (response *DescribeDSPAAssessmentTasksResponse) {
    response = &DescribeDSPAAssessmentTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPAAssessmentTasks
// 获取DSPA评估任务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDSPAAssessmentTasks(request *DescribeDSPAAssessmentTasksRequest) (response *DescribeDSPAAssessmentTasksResponse, err error) {
    return c.DescribeDSPAAssessmentTasksWithContext(context.Background(), request)
}

// DescribeDSPAAssessmentTasks
// 获取DSPA评估任务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDSPAAssessmentTasksWithContext(ctx context.Context, request *DescribeDSPAAssessmentTasksRequest) (response *DescribeDSPAAssessmentTasksResponse, err error) {
    if request == nil {
        request = NewDescribeDSPAAssessmentTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPAAssessmentTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPAAssessmentTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPAAssessmentTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPAAssessmentTemplateControlItemsRequest() (request *DescribeDSPAAssessmentTemplateControlItemsRequest) {
    request = &DescribeDSPAAssessmentTemplateControlItemsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPAAssessmentTemplateControlItems")
    
    
    return
}

func NewDescribeDSPAAssessmentTemplateControlItemsResponse() (response *DescribeDSPAAssessmentTemplateControlItemsResponse) {
    response = &DescribeDSPAAssessmentTemplateControlItemsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPAAssessmentTemplateControlItems
// 获取DSPA评估模板关联的评估控制项列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDSPAAssessmentTemplateControlItems(request *DescribeDSPAAssessmentTemplateControlItemsRequest) (response *DescribeDSPAAssessmentTemplateControlItemsResponse, err error) {
    return c.DescribeDSPAAssessmentTemplateControlItemsWithContext(context.Background(), request)
}

// DescribeDSPAAssessmentTemplateControlItems
// 获取DSPA评估模板关联的评估控制项列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDSPAAssessmentTemplateControlItemsWithContext(ctx context.Context, request *DescribeDSPAAssessmentTemplateControlItemsRequest) (response *DescribeDSPAAssessmentTemplateControlItemsResponse, err error) {
    if request == nil {
        request = NewDescribeDSPAAssessmentTemplateControlItemsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPAAssessmentTemplateControlItems")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPAAssessmentTemplateControlItems require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPAAssessmentTemplateControlItemsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPAAssessmentTemplatesRequest() (request *DescribeDSPAAssessmentTemplatesRequest) {
    request = &DescribeDSPAAssessmentTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPAAssessmentTemplates")
    
    
    return
}

func NewDescribeDSPAAssessmentTemplatesResponse() (response *DescribeDSPAAssessmentTemplatesResponse) {
    response = &DescribeDSPAAssessmentTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPAAssessmentTemplates
// 获取DSPA评估模板列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDSPAAssessmentTemplates(request *DescribeDSPAAssessmentTemplatesRequest) (response *DescribeDSPAAssessmentTemplatesResponse, err error) {
    return c.DescribeDSPAAssessmentTemplatesWithContext(context.Background(), request)
}

// DescribeDSPAAssessmentTemplates
// 获取DSPA评估模板列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDSPAAssessmentTemplatesWithContext(ctx context.Context, request *DescribeDSPAAssessmentTemplatesRequest) (response *DescribeDSPAAssessmentTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeDSPAAssessmentTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPAAssessmentTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPAAssessmentTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPAAssessmentTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPACOSDataAssetBucketsRequest() (request *DescribeDSPACOSDataAssetBucketsRequest) {
    request = &DescribeDSPACOSDataAssetBucketsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPACOSDataAssetBuckets")
    
    
    return
}

func NewDescribeDSPACOSDataAssetBucketsResponse() (response *DescribeDSPACOSDataAssetBucketsResponse) {
    response = &DescribeDSPACOSDataAssetBucketsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPACOSDataAssetBuckets
// 获取COS敏感数据资产桶列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDSPACOSDataAssetBuckets(request *DescribeDSPACOSDataAssetBucketsRequest) (response *DescribeDSPACOSDataAssetBucketsResponse, err error) {
    return c.DescribeDSPACOSDataAssetBucketsWithContext(context.Background(), request)
}

// DescribeDSPACOSDataAssetBuckets
// 获取COS敏感数据资产桶列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDSPACOSDataAssetBucketsWithContext(ctx context.Context, request *DescribeDSPACOSDataAssetBucketsRequest) (response *DescribeDSPACOSDataAssetBucketsResponse, err error) {
    if request == nil {
        request = NewDescribeDSPACOSDataAssetBucketsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPACOSDataAssetBuckets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPACOSDataAssetBuckets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPACOSDataAssetBucketsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPACOSDataAssetByComplianceIdRequest() (request *DescribeDSPACOSDataAssetByComplianceIdRequest) {
    request = &DescribeDSPACOSDataAssetByComplianceIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPACOSDataAssetByComplianceId")
    
    
    return
}

func NewDescribeDSPACOSDataAssetByComplianceIdResponse() (response *DescribeDSPACOSDataAssetByComplianceIdResponse) {
    response = &DescribeDSPACOSDataAssetByComplianceIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPACOSDataAssetByComplianceId
// 获取COS单个模板下的敏感数据资产
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDSPACOSDataAssetByComplianceId(request *DescribeDSPACOSDataAssetByComplianceIdRequest) (response *DescribeDSPACOSDataAssetByComplianceIdResponse, err error) {
    return c.DescribeDSPACOSDataAssetByComplianceIdWithContext(context.Background(), request)
}

// DescribeDSPACOSDataAssetByComplianceId
// 获取COS单个模板下的敏感数据资产
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDSPACOSDataAssetByComplianceIdWithContext(ctx context.Context, request *DescribeDSPACOSDataAssetByComplianceIdRequest) (response *DescribeDSPACOSDataAssetByComplianceIdResponse, err error) {
    if request == nil {
        request = NewDescribeDSPACOSDataAssetByComplianceIdRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPACOSDataAssetByComplianceId")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPACOSDataAssetByComplianceId require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPACOSDataAssetByComplianceIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPACOSDataAssetDetailRequest() (request *DescribeDSPACOSDataAssetDetailRequest) {
    request = &DescribeDSPACOSDataAssetDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPACOSDataAssetDetail")
    
    
    return
}

func NewDescribeDSPACOSDataAssetDetailResponse() (response *DescribeDSPACOSDataAssetDetailResponse) {
    response = &DescribeDSPACOSDataAssetDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPACOSDataAssetDetail
// 获取COS分类分级数据资产详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDSPACOSDataAssetDetail(request *DescribeDSPACOSDataAssetDetailRequest) (response *DescribeDSPACOSDataAssetDetailResponse, err error) {
    return c.DescribeDSPACOSDataAssetDetailWithContext(context.Background(), request)
}

// DescribeDSPACOSDataAssetDetail
// 获取COS分类分级数据资产详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDSPACOSDataAssetDetailWithContext(ctx context.Context, request *DescribeDSPACOSDataAssetDetailRequest) (response *DescribeDSPACOSDataAssetDetailResponse, err error) {
    if request == nil {
        request = NewDescribeDSPACOSDataAssetDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPACOSDataAssetDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPACOSDataAssetDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPACOSDataAssetDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPACOSDiscoveryTaskDetailRequest() (request *DescribeDSPACOSDiscoveryTaskDetailRequest) {
    request = &DescribeDSPACOSDiscoveryTaskDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPACOSDiscoveryTaskDetail")
    
    
    return
}

func NewDescribeDSPACOSDiscoveryTaskDetailResponse() (response *DescribeDSPACOSDiscoveryTaskDetailResponse) {
    response = &DescribeDSPACOSDiscoveryTaskDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPACOSDiscoveryTaskDetail
// 获取COS分类分级任务详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDSPACOSDiscoveryTaskDetail(request *DescribeDSPACOSDiscoveryTaskDetailRequest) (response *DescribeDSPACOSDiscoveryTaskDetailResponse, err error) {
    return c.DescribeDSPACOSDiscoveryTaskDetailWithContext(context.Background(), request)
}

// DescribeDSPACOSDiscoveryTaskDetail
// 获取COS分类分级任务详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDSPACOSDiscoveryTaskDetailWithContext(ctx context.Context, request *DescribeDSPACOSDiscoveryTaskDetailRequest) (response *DescribeDSPACOSDiscoveryTaskDetailResponse, err error) {
    if request == nil {
        request = NewDescribeDSPACOSDiscoveryTaskDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPACOSDiscoveryTaskDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPACOSDiscoveryTaskDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPACOSDiscoveryTaskDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPACOSDiscoveryTaskFilesRequest() (request *DescribeDSPACOSDiscoveryTaskFilesRequest) {
    request = &DescribeDSPACOSDiscoveryTaskFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPACOSDiscoveryTaskFiles")
    
    
    return
}

func NewDescribeDSPACOSDiscoveryTaskFilesResponse() (response *DescribeDSPACOSDiscoveryTaskFilesResponse) {
    response = &DescribeDSPACOSDiscoveryTaskFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPACOSDiscoveryTaskFiles
// 获取COS分类分级任务结果详情文件列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDSPACOSDiscoveryTaskFiles(request *DescribeDSPACOSDiscoveryTaskFilesRequest) (response *DescribeDSPACOSDiscoveryTaskFilesResponse, err error) {
    return c.DescribeDSPACOSDiscoveryTaskFilesWithContext(context.Background(), request)
}

// DescribeDSPACOSDiscoveryTaskFiles
// 获取COS分类分级任务结果详情文件列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDSPACOSDiscoveryTaskFilesWithContext(ctx context.Context, request *DescribeDSPACOSDiscoveryTaskFilesRequest) (response *DescribeDSPACOSDiscoveryTaskFilesResponse, err error) {
    if request == nil {
        request = NewDescribeDSPACOSDiscoveryTaskFilesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPACOSDiscoveryTaskFiles")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPACOSDiscoveryTaskFiles require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPACOSDiscoveryTaskFilesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPACOSDiscoveryTaskResultRequest() (request *DescribeDSPACOSDiscoveryTaskResultRequest) {
    request = &DescribeDSPACOSDiscoveryTaskResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPACOSDiscoveryTaskResult")
    
    
    return
}

func NewDescribeDSPACOSDiscoveryTaskResultResponse() (response *DescribeDSPACOSDiscoveryTaskResultResponse) {
    response = &DescribeDSPACOSDiscoveryTaskResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPACOSDiscoveryTaskResult
// 获取COS分类分级任务结果，该接口只有在任务状态为以下状态时才支持结果正常查询：
//
// 3 扫描成功，
//
// 4 扫描失败
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDSPACOSDiscoveryTaskResult(request *DescribeDSPACOSDiscoveryTaskResultRequest) (response *DescribeDSPACOSDiscoveryTaskResultResponse, err error) {
    return c.DescribeDSPACOSDiscoveryTaskResultWithContext(context.Background(), request)
}

// DescribeDSPACOSDiscoveryTaskResult
// 获取COS分类分级任务结果，该接口只有在任务状态为以下状态时才支持结果正常查询：
//
// 3 扫描成功，
//
// 4 扫描失败
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDSPACOSDiscoveryTaskResultWithContext(ctx context.Context, request *DescribeDSPACOSDiscoveryTaskResultRequest) (response *DescribeDSPACOSDiscoveryTaskResultResponse, err error) {
    if request == nil {
        request = NewDescribeDSPACOSDiscoveryTaskResultRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPACOSDiscoveryTaskResult")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPACOSDiscoveryTaskResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPACOSDiscoveryTaskResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPACOSDiscoveryTasksRequest() (request *DescribeDSPACOSDiscoveryTasksRequest) {
    request = &DescribeDSPACOSDiscoveryTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPACOSDiscoveryTasks")
    
    
    return
}

func NewDescribeDSPACOSDiscoveryTasksResponse() (response *DescribeDSPACOSDiscoveryTasksResponse) {
    response = &DescribeDSPACOSDiscoveryTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPACOSDiscoveryTasks
// 获取COS分类分级任务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDSPACOSDiscoveryTasks(request *DescribeDSPACOSDiscoveryTasksRequest) (response *DescribeDSPACOSDiscoveryTasksResponse, err error) {
    return c.DescribeDSPACOSDiscoveryTasksWithContext(context.Background(), request)
}

// DescribeDSPACOSDiscoveryTasks
// 获取COS分类分级任务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDSPACOSDiscoveryTasksWithContext(ctx context.Context, request *DescribeDSPACOSDiscoveryTasksRequest) (response *DescribeDSPACOSDiscoveryTasksResponse, err error) {
    if request == nil {
        request = NewDescribeDSPACOSDiscoveryTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPACOSDiscoveryTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPACOSDiscoveryTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPACOSDiscoveryTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPACOSTaskResultDetailRequest() (request *DescribeDSPACOSTaskResultDetailRequest) {
    request = &DescribeDSPACOSTaskResultDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPACOSTaskResultDetail")
    
    
    return
}

func NewDescribeDSPACOSTaskResultDetailResponse() (response *DescribeDSPACOSTaskResultDetailResponse) {
    response = &DescribeDSPACOSTaskResultDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPACOSTaskResultDetail
// 获取COS分类分级任务结果详情，该接口只有在任务状态为时才支持结果正确查询：
//
// 3 扫描成功
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDSPACOSTaskResultDetail(request *DescribeDSPACOSTaskResultDetailRequest) (response *DescribeDSPACOSTaskResultDetailResponse, err error) {
    return c.DescribeDSPACOSTaskResultDetailWithContext(context.Background(), request)
}

// DescribeDSPACOSTaskResultDetail
// 获取COS分类分级任务结果详情，该接口只有在任务状态为时才支持结果正确查询：
//
// 3 扫描成功
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDSPACOSTaskResultDetailWithContext(ctx context.Context, request *DescribeDSPACOSTaskResultDetailRequest) (response *DescribeDSPACOSTaskResultDetailResponse, err error) {
    if request == nil {
        request = NewDescribeDSPACOSTaskResultDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPACOSTaskResultDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPACOSTaskResultDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPACOSTaskResultDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPACategoriesRequest() (request *DescribeDSPACategoriesRequest) {
    request = &DescribeDSPACategoriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPACategories")
    
    
    return
}

func NewDescribeDSPACategoriesResponse() (response *DescribeDSPACategoriesResponse) {
    response = &DescribeDSPACategoriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPACategories
// 获取敏感数据分类列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDSPACategories(request *DescribeDSPACategoriesRequest) (response *DescribeDSPACategoriesResponse, err error) {
    return c.DescribeDSPACategoriesWithContext(context.Background(), request)
}

// DescribeDSPACategories
// 获取敏感数据分类列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDSPACategoriesWithContext(ctx context.Context, request *DescribeDSPACategoriesRequest) (response *DescribeDSPACategoriesResponse, err error) {
    if request == nil {
        request = NewDescribeDSPACategoriesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPACategories")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPACategories require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPACategoriesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPACategoryRuleStatisticRequest() (request *DescribeDSPACategoryRuleStatisticRequest) {
    request = &DescribeDSPACategoryRuleStatisticRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPACategoryRuleStatistic")
    
    
    return
}

func NewDescribeDSPACategoryRuleStatisticResponse() (response *DescribeDSPACategoryRuleStatisticResponse) {
    response = &DescribeDSPACategoryRuleStatisticResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPACategoryRuleStatistic
// 获取分类规则统计信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDSPACategoryRuleStatistic(request *DescribeDSPACategoryRuleStatisticRequest) (response *DescribeDSPACategoryRuleStatisticResponse, err error) {
    return c.DescribeDSPACategoryRuleStatisticWithContext(context.Background(), request)
}

// DescribeDSPACategoryRuleStatistic
// 获取分类规则统计信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDSPACategoryRuleStatisticWithContext(ctx context.Context, request *DescribeDSPACategoryRuleStatisticRequest) (response *DescribeDSPACategoryRuleStatisticResponse, err error) {
    if request == nil {
        request = NewDescribeDSPACategoryRuleStatisticRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPACategoryRuleStatistic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPACategoryRuleStatistic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPACategoryRuleStatisticResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPACategoryRulesRequest() (request *DescribeDSPACategoryRulesRequest) {
    request = &DescribeDSPACategoryRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPACategoryRules")
    
    
    return
}

func NewDescribeDSPACategoryRulesResponse() (response *DescribeDSPACategoryRulesResponse) {
    response = &DescribeDSPACategoryRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPACategoryRules
// 获取合规组分类规则信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDSPACategoryRules(request *DescribeDSPACategoryRulesRequest) (response *DescribeDSPACategoryRulesResponse, err error) {
    return c.DescribeDSPACategoryRulesWithContext(context.Background(), request)
}

// DescribeDSPACategoryRules
// 获取合规组分类规则信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDSPACategoryRulesWithContext(ctx context.Context, request *DescribeDSPACategoryRulesRequest) (response *DescribeDSPACategoryRulesResponse, err error) {
    if request == nil {
        request = NewDescribeDSPACategoryRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPACategoryRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPACategoryRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPACategoryRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPACategoryTreeRequest() (request *DescribeDSPACategoryTreeRequest) {
    request = &DescribeDSPACategoryTreeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPACategoryTree")
    
    
    return
}

func NewDescribeDSPACategoryTreeResponse() (response *DescribeDSPACategoryTreeResponse) {
    response = &DescribeDSPACategoryTreeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPACategoryTree
// 获取分类树信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDSPACategoryTree(request *DescribeDSPACategoryTreeRequest) (response *DescribeDSPACategoryTreeResponse, err error) {
    return c.DescribeDSPACategoryTreeWithContext(context.Background(), request)
}

// DescribeDSPACategoryTree
// 获取分类树信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDSPACategoryTreeWithContext(ctx context.Context, request *DescribeDSPACategoryTreeRequest) (response *DescribeDSPACategoryTreeResponse, err error) {
    if request == nil {
        request = NewDescribeDSPACategoryTreeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPACategoryTree")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPACategoryTree require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPACategoryTreeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPACategoryTreeWithRulesRequest() (request *DescribeDSPACategoryTreeWithRulesRequest) {
    request = &DescribeDSPACategoryTreeWithRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPACategoryTreeWithRules")
    
    
    return
}

func NewDescribeDSPACategoryTreeWithRulesResponse() (response *DescribeDSPACategoryTreeWithRulesResponse) {
    response = &DescribeDSPACategoryTreeWithRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPACategoryTreeWithRules
// 获取分类规则树信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDSPACategoryTreeWithRules(request *DescribeDSPACategoryTreeWithRulesRequest) (response *DescribeDSPACategoryTreeWithRulesResponse, err error) {
    return c.DescribeDSPACategoryTreeWithRulesWithContext(context.Background(), request)
}

// DescribeDSPACategoryTreeWithRules
// 获取分类规则树信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDSPACategoryTreeWithRulesWithContext(ctx context.Context, request *DescribeDSPACategoryTreeWithRulesRequest) (response *DescribeDSPACategoryTreeWithRulesResponse, err error) {
    if request == nil {
        request = NewDescribeDSPACategoryTreeWithRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPACategoryTreeWithRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPACategoryTreeWithRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPACategoryTreeWithRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPAComplianceGroupDetailRequest() (request *DescribeDSPAComplianceGroupDetailRequest) {
    request = &DescribeDSPAComplianceGroupDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPAComplianceGroupDetail")
    
    
    return
}

func NewDescribeDSPAComplianceGroupDetailResponse() (response *DescribeDSPAComplianceGroupDetailResponse) {
    response = &DescribeDSPAComplianceGroupDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPAComplianceGroupDetail
// 获取模板详情信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDSPAComplianceGroupDetail(request *DescribeDSPAComplianceGroupDetailRequest) (response *DescribeDSPAComplianceGroupDetailResponse, err error) {
    return c.DescribeDSPAComplianceGroupDetailWithContext(context.Background(), request)
}

// DescribeDSPAComplianceGroupDetail
// 获取模板详情信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDSPAComplianceGroupDetailWithContext(ctx context.Context, request *DescribeDSPAComplianceGroupDetailRequest) (response *DescribeDSPAComplianceGroupDetailResponse, err error) {
    if request == nil {
        request = NewDescribeDSPAComplianceGroupDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPAComplianceGroupDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPAComplianceGroupDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPAComplianceGroupDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPAComplianceGroupsRequest() (request *DescribeDSPAComplianceGroupsRequest) {
    request = &DescribeDSPAComplianceGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPAComplianceGroups")
    
    
    return
}

func NewDescribeDSPAComplianceGroupsResponse() (response *DescribeDSPAComplianceGroupsResponse) {
    response = &DescribeDSPAComplianceGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPAComplianceGroups
// 获取分类分级模板列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDSPAComplianceGroups(request *DescribeDSPAComplianceGroupsRequest) (response *DescribeDSPAComplianceGroupsResponse, err error) {
    return c.DescribeDSPAComplianceGroupsWithContext(context.Background(), request)
}

// DescribeDSPAComplianceGroups
// 获取分类分级模板列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDSPAComplianceGroupsWithContext(ctx context.Context, request *DescribeDSPAComplianceGroupsRequest) (response *DescribeDSPAComplianceGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeDSPAComplianceGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPAComplianceGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPAComplianceGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPAComplianceGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPAComplianceUpdateNotificationRequest() (request *DescribeDSPAComplianceUpdateNotificationRequest) {
    request = &DescribeDSPAComplianceUpdateNotificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPAComplianceUpdateNotification")
    
    
    return
}

func NewDescribeDSPAComplianceUpdateNotificationResponse() (response *DescribeDSPAComplianceUpdateNotificationResponse) {
    response = &DescribeDSPAComplianceUpdateNotificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPAComplianceUpdateNotification
// 获取模板更新提示信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDSPAComplianceUpdateNotification(request *DescribeDSPAComplianceUpdateNotificationRequest) (response *DescribeDSPAComplianceUpdateNotificationResponse, err error) {
    return c.DescribeDSPAComplianceUpdateNotificationWithContext(context.Background(), request)
}

// DescribeDSPAComplianceUpdateNotification
// 获取模板更新提示信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDSPAComplianceUpdateNotificationWithContext(ctx context.Context, request *DescribeDSPAComplianceUpdateNotificationRequest) (response *DescribeDSPAComplianceUpdateNotificationResponse, err error) {
    if request == nil {
        request = NewDescribeDSPAComplianceUpdateNotificationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPAComplianceUpdateNotification")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPAComplianceUpdateNotification require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPAComplianceUpdateNotificationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPADataSourceDbInfoRequest() (request *DescribeDSPADataSourceDbInfoRequest) {
    request = &DescribeDSPADataSourceDbInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPADataSourceDbInfo")
    
    
    return
}

func NewDescribeDSPADataSourceDbInfoResponse() (response *DescribeDSPADataSourceDbInfoResponse) {
    response = &DescribeDSPADataSourceDbInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPADataSourceDbInfo
// 获取数据源的数据库信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDSPADataSourceDbInfo(request *DescribeDSPADataSourceDbInfoRequest) (response *DescribeDSPADataSourceDbInfoResponse, err error) {
    return c.DescribeDSPADataSourceDbInfoWithContext(context.Background(), request)
}

// DescribeDSPADataSourceDbInfo
// 获取数据源的数据库信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDSPADataSourceDbInfoWithContext(ctx context.Context, request *DescribeDSPADataSourceDbInfoRequest) (response *DescribeDSPADataSourceDbInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDSPADataSourceDbInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPADataSourceDbInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPADataSourceDbInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPADataSourceDbInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPADiscoveryRulesRequest() (request *DescribeDSPADiscoveryRulesRequest) {
    request = &DescribeDSPADiscoveryRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPADiscoveryRules")
    
    
    return
}

func NewDescribeDSPADiscoveryRulesResponse() (response *DescribeDSPADiscoveryRulesResponse) {
    response = &DescribeDSPADiscoveryRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPADiscoveryRules
// 获取分类分级规则列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDSPADiscoveryRules(request *DescribeDSPADiscoveryRulesRequest) (response *DescribeDSPADiscoveryRulesResponse, err error) {
    return c.DescribeDSPADiscoveryRulesWithContext(context.Background(), request)
}

// DescribeDSPADiscoveryRules
// 获取分类分级规则列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDSPADiscoveryRulesWithContext(ctx context.Context, request *DescribeDSPADiscoveryRulesRequest) (response *DescribeDSPADiscoveryRulesResponse, err error) {
    if request == nil {
        request = NewDescribeDSPADiscoveryRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPADiscoveryRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPADiscoveryRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPADiscoveryRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPADiscoveryServiceStatusRequest() (request *DescribeDSPADiscoveryServiceStatusRequest) {
    request = &DescribeDSPADiscoveryServiceStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPADiscoveryServiceStatus")
    
    
    return
}

func NewDescribeDSPADiscoveryServiceStatusResponse() (response *DescribeDSPADiscoveryServiceStatusResponse) {
    response = &DescribeDSPADiscoveryServiceStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPADiscoveryServiceStatus
// 用于查询该用户是否已开通分类分级服务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDSPADiscoveryServiceStatus(request *DescribeDSPADiscoveryServiceStatusRequest) (response *DescribeDSPADiscoveryServiceStatusResponse, err error) {
    return c.DescribeDSPADiscoveryServiceStatusWithContext(context.Background(), request)
}

// DescribeDSPADiscoveryServiceStatus
// 用于查询该用户是否已开通分类分级服务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDSPADiscoveryServiceStatusWithContext(ctx context.Context, request *DescribeDSPADiscoveryServiceStatusRequest) (response *DescribeDSPADiscoveryServiceStatusResponse, err error) {
    if request == nil {
        request = NewDescribeDSPADiscoveryServiceStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPADiscoveryServiceStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPADiscoveryServiceStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPADiscoveryServiceStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPADiscoveryTaskDetailRequest() (request *DescribeDSPADiscoveryTaskDetailRequest) {
    request = &DescribeDSPADiscoveryTaskDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPADiscoveryTaskDetail")
    
    
    return
}

func NewDescribeDSPADiscoveryTaskDetailResponse() (response *DescribeDSPADiscoveryTaskDetailResponse) {
    response = &DescribeDSPADiscoveryTaskDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPADiscoveryTaskDetail
// 获取分类分级任务详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDSPADiscoveryTaskDetail(request *DescribeDSPADiscoveryTaskDetailRequest) (response *DescribeDSPADiscoveryTaskDetailResponse, err error) {
    return c.DescribeDSPADiscoveryTaskDetailWithContext(context.Background(), request)
}

// DescribeDSPADiscoveryTaskDetail
// 获取分类分级任务详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDSPADiscoveryTaskDetailWithContext(ctx context.Context, request *DescribeDSPADiscoveryTaskDetailRequest) (response *DescribeDSPADiscoveryTaskDetailResponse, err error) {
    if request == nil {
        request = NewDescribeDSPADiscoveryTaskDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPADiscoveryTaskDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPADiscoveryTaskDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPADiscoveryTaskDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPADiscoveryTaskResultRequest() (request *DescribeDSPADiscoveryTaskResultRequest) {
    request = &DescribeDSPADiscoveryTaskResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPADiscoveryTaskResult")
    
    
    return
}

func NewDescribeDSPADiscoveryTaskResultResponse() (response *DescribeDSPADiscoveryTaskResultResponse) {
    response = &DescribeDSPADiscoveryTaskResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPADiscoveryTaskResult
// 获取分类分级任务结果，该接口只有在任务状态为以下状态时才支持结果正常查询：3 扫描成功，4 扫描失败
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDSPADiscoveryTaskResult(request *DescribeDSPADiscoveryTaskResultRequest) (response *DescribeDSPADiscoveryTaskResultResponse, err error) {
    return c.DescribeDSPADiscoveryTaskResultWithContext(context.Background(), request)
}

// DescribeDSPADiscoveryTaskResult
// 获取分类分级任务结果，该接口只有在任务状态为以下状态时才支持结果正常查询：3 扫描成功，4 扫描失败
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDSPADiscoveryTaskResultWithContext(ctx context.Context, request *DescribeDSPADiscoveryTaskResultRequest) (response *DescribeDSPADiscoveryTaskResultResponse, err error) {
    if request == nil {
        request = NewDescribeDSPADiscoveryTaskResultRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPADiscoveryTaskResult")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPADiscoveryTaskResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPADiscoveryTaskResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPADiscoveryTaskResultDetailRequest() (request *DescribeDSPADiscoveryTaskResultDetailRequest) {
    request = &DescribeDSPADiscoveryTaskResultDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPADiscoveryTaskResultDetail")
    
    
    return
}

func NewDescribeDSPADiscoveryTaskResultDetailResponse() (response *DescribeDSPADiscoveryTaskResultDetailResponse) {
    response = &DescribeDSPADiscoveryTaskResultDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPADiscoveryTaskResultDetail
// 获取分类分级任务结果详情，该接口只有在任务状态为时才支持结果正确查询：
//
// 3 扫描成功
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDSPADiscoveryTaskResultDetail(request *DescribeDSPADiscoveryTaskResultDetailRequest) (response *DescribeDSPADiscoveryTaskResultDetailResponse, err error) {
    return c.DescribeDSPADiscoveryTaskResultDetailWithContext(context.Background(), request)
}

// DescribeDSPADiscoveryTaskResultDetail
// 获取分类分级任务结果详情，该接口只有在任务状态为时才支持结果正确查询：
//
// 3 扫描成功
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDSPADiscoveryTaskResultDetailWithContext(ctx context.Context, request *DescribeDSPADiscoveryTaskResultDetailRequest) (response *DescribeDSPADiscoveryTaskResultDetailResponse, err error) {
    if request == nil {
        request = NewDescribeDSPADiscoveryTaskResultDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPADiscoveryTaskResultDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPADiscoveryTaskResultDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPADiscoveryTaskResultDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPADiscoveryTaskTablesRequest() (request *DescribeDSPADiscoveryTaskTablesRequest) {
    request = &DescribeDSPADiscoveryTaskTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPADiscoveryTaskTables")
    
    
    return
}

func NewDescribeDSPADiscoveryTaskTablesResponse() (response *DescribeDSPADiscoveryTaskTablesResponse) {
    response = &DescribeDSPADiscoveryTaskTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPADiscoveryTaskTables
// 获取分级分级扫描的表集合
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDSPADiscoveryTaskTables(request *DescribeDSPADiscoveryTaskTablesRequest) (response *DescribeDSPADiscoveryTaskTablesResponse, err error) {
    return c.DescribeDSPADiscoveryTaskTablesWithContext(context.Background(), request)
}

// DescribeDSPADiscoveryTaskTables
// 获取分级分级扫描的表集合
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDSPADiscoveryTaskTablesWithContext(ctx context.Context, request *DescribeDSPADiscoveryTaskTablesRequest) (response *DescribeDSPADiscoveryTaskTablesResponse, err error) {
    if request == nil {
        request = NewDescribeDSPADiscoveryTaskTablesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPADiscoveryTaskTables")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPADiscoveryTaskTables require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPADiscoveryTaskTablesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPADiscoveryTasksRequest() (request *DescribeDSPADiscoveryTasksRequest) {
    request = &DescribeDSPADiscoveryTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPADiscoveryTasks")
    
    
    return
}

func NewDescribeDSPADiscoveryTasksResponse() (response *DescribeDSPADiscoveryTasksResponse) {
    response = &DescribeDSPADiscoveryTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPADiscoveryTasks
// 获取分类分级任务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDSPADiscoveryTasks(request *DescribeDSPADiscoveryTasksRequest) (response *DescribeDSPADiscoveryTasksResponse, err error) {
    return c.DescribeDSPADiscoveryTasksWithContext(context.Background(), request)
}

// DescribeDSPADiscoveryTasks
// 获取分类分级任务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDSPADiscoveryTasksWithContext(ctx context.Context, request *DescribeDSPADiscoveryTasksRequest) (response *DescribeDSPADiscoveryTasksResponse, err error) {
    if request == nil {
        request = NewDescribeDSPADiscoveryTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPADiscoveryTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPADiscoveryTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPADiscoveryTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPAESDataAssetByComplianceIdRequest() (request *DescribeDSPAESDataAssetByComplianceIdRequest) {
    request = &DescribeDSPAESDataAssetByComplianceIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPAESDataAssetByComplianceId")
    
    
    return
}

func NewDescribeDSPAESDataAssetByComplianceIdResponse() (response *DescribeDSPAESDataAssetByComplianceIdResponse) {
    response = &DescribeDSPAESDataAssetByComplianceIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPAESDataAssetByComplianceId
// 根据合规组id，去查询ES的概览页统计数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDSPAESDataAssetByComplianceId(request *DescribeDSPAESDataAssetByComplianceIdRequest) (response *DescribeDSPAESDataAssetByComplianceIdResponse, err error) {
    return c.DescribeDSPAESDataAssetByComplianceIdWithContext(context.Background(), request)
}

// DescribeDSPAESDataAssetByComplianceId
// 根据合规组id，去查询ES的概览页统计数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDSPAESDataAssetByComplianceIdWithContext(ctx context.Context, request *DescribeDSPAESDataAssetByComplianceIdRequest) (response *DescribeDSPAESDataAssetByComplianceIdResponse, err error) {
    if request == nil {
        request = NewDescribeDSPAESDataAssetByComplianceIdRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPAESDataAssetByComplianceId")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPAESDataAssetByComplianceId require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPAESDataAssetByComplianceIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPAESDataAssetDetailRequest() (request *DescribeDSPAESDataAssetDetailRequest) {
    request = &DescribeDSPAESDataAssetDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPAESDataAssetDetail")
    
    
    return
}

func NewDescribeDSPAESDataAssetDetailResponse() (response *DescribeDSPAESDataAssetDetailResponse) {
    response = &DescribeDSPAESDataAssetDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPAESDataAssetDetail
// 根据合规组id，去查询ES的概览页下的统计列表数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDSPAESDataAssetDetail(request *DescribeDSPAESDataAssetDetailRequest) (response *DescribeDSPAESDataAssetDetailResponse, err error) {
    return c.DescribeDSPAESDataAssetDetailWithContext(context.Background(), request)
}

// DescribeDSPAESDataAssetDetail
// 根据合规组id，去查询ES的概览页下的统计列表数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDSPAESDataAssetDetailWithContext(ctx context.Context, request *DescribeDSPAESDataAssetDetailRequest) (response *DescribeDSPAESDataAssetDetailResponse, err error) {
    if request == nil {
        request = NewDescribeDSPAESDataAssetDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPAESDataAssetDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPAESDataAssetDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPAESDataAssetDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPAESDataSampleRequest() (request *DescribeDSPAESDataSampleRequest) {
    request = &DescribeDSPAESDataSampleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPAESDataSample")
    
    
    return
}

func NewDescribeDSPAESDataSampleResponse() (response *DescribeDSPAESDataSampleResponse) {
    response = &DescribeDSPAESDataSampleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPAESDataSample
// 获取ES扫描结果数据样本
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDSPAESDataSample(request *DescribeDSPAESDataSampleRequest) (response *DescribeDSPAESDataSampleResponse, err error) {
    return c.DescribeDSPAESDataSampleWithContext(context.Background(), request)
}

// DescribeDSPAESDataSample
// 获取ES扫描结果数据样本
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDSPAESDataSampleWithContext(ctx context.Context, request *DescribeDSPAESDataSampleRequest) (response *DescribeDSPAESDataSampleResponse, err error) {
    if request == nil {
        request = NewDescribeDSPAESDataSampleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPAESDataSample")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPAESDataSample require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPAESDataSampleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPAESDiscoveryTaskResultDetailRequest() (request *DescribeDSPAESDiscoveryTaskResultDetailRequest) {
    request = &DescribeDSPAESDiscoveryTaskResultDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPAESDiscoveryTaskResultDetail")
    
    
    return
}

func NewDescribeDSPAESDiscoveryTaskResultDetailResponse() (response *DescribeDSPAESDiscoveryTaskResultDetailResponse) {
    response = &DescribeDSPAESDiscoveryTaskResultDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPAESDiscoveryTaskResultDetail
// 获取ES的分类分级任务结果详情，该接口只有在任务状态为时才支持结果正确查询：
//
// 3 扫描成功
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDSPAESDiscoveryTaskResultDetail(request *DescribeDSPAESDiscoveryTaskResultDetailRequest) (response *DescribeDSPAESDiscoveryTaskResultDetailResponse, err error) {
    return c.DescribeDSPAESDiscoveryTaskResultDetailWithContext(context.Background(), request)
}

// DescribeDSPAESDiscoveryTaskResultDetail
// 获取ES的分类分级任务结果详情，该接口只有在任务状态为时才支持结果正确查询：
//
// 3 扫描成功
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDSPAESDiscoveryTaskResultDetailWithContext(ctx context.Context, request *DescribeDSPAESDiscoveryTaskResultDetailRequest) (response *DescribeDSPAESDiscoveryTaskResultDetailResponse, err error) {
    if request == nil {
        request = NewDescribeDSPAESDiscoveryTaskResultDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPAESDiscoveryTaskResultDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPAESDiscoveryTaskResultDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPAESDiscoveryTaskResultDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPALevelDetailRequest() (request *DescribeDSPALevelDetailRequest) {
    request = &DescribeDSPALevelDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPALevelDetail")
    
    
    return
}

func NewDescribeDSPALevelDetailResponse() (response *DescribeDSPALevelDetailResponse) {
    response = &DescribeDSPALevelDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPALevelDetail
// 获取分级详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDSPALevelDetail(request *DescribeDSPALevelDetailRequest) (response *DescribeDSPALevelDetailResponse, err error) {
    return c.DescribeDSPALevelDetailWithContext(context.Background(), request)
}

// DescribeDSPALevelDetail
// 获取分级详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDSPALevelDetailWithContext(ctx context.Context, request *DescribeDSPALevelDetailRequest) (response *DescribeDSPALevelDetailResponse, err error) {
    if request == nil {
        request = NewDescribeDSPALevelDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPALevelDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPALevelDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPALevelDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPALevelGroupsRequest() (request *DescribeDSPALevelGroupsRequest) {
    request = &DescribeDSPALevelGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPALevelGroups")
    
    
    return
}

func NewDescribeDSPALevelGroupsResponse() (response *DescribeDSPALevelGroupsResponse) {
    response = &DescribeDSPALevelGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPALevelGroups
// 获取分级列表，限制100个 不分页返回
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDSPALevelGroups(request *DescribeDSPALevelGroupsRequest) (response *DescribeDSPALevelGroupsResponse, err error) {
    return c.DescribeDSPALevelGroupsWithContext(context.Background(), request)
}

// DescribeDSPALevelGroups
// 获取分级列表，限制100个 不分页返回
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDSPALevelGroupsWithContext(ctx context.Context, request *DescribeDSPALevelGroupsRequest) (response *DescribeDSPALevelGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeDSPALevelGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPALevelGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPALevelGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPALevelGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPARDBDataAssetByComplianceIdRequest() (request *DescribeDSPARDBDataAssetByComplianceIdRequest) {
    request = &DescribeDSPARDBDataAssetByComplianceIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPARDBDataAssetByComplianceId")
    
    
    return
}

func NewDescribeDSPARDBDataAssetByComplianceIdResponse() (response *DescribeDSPARDBDataAssetByComplianceIdResponse) {
    response = &DescribeDSPARDBDataAssetByComplianceIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPARDBDataAssetByComplianceId
// 获取单个合规组下的RDB关系数据库分类分级数据资产
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDSPARDBDataAssetByComplianceId(request *DescribeDSPARDBDataAssetByComplianceIdRequest) (response *DescribeDSPARDBDataAssetByComplianceIdResponse, err error) {
    return c.DescribeDSPARDBDataAssetByComplianceIdWithContext(context.Background(), request)
}

// DescribeDSPARDBDataAssetByComplianceId
// 获取单个合规组下的RDB关系数据库分类分级数据资产
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeDSPARDBDataAssetByComplianceIdWithContext(ctx context.Context, request *DescribeDSPARDBDataAssetByComplianceIdRequest) (response *DescribeDSPARDBDataAssetByComplianceIdResponse, err error) {
    if request == nil {
        request = NewDescribeDSPARDBDataAssetByComplianceIdRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPARDBDataAssetByComplianceId")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPARDBDataAssetByComplianceId require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPARDBDataAssetByComplianceIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPARDBDataAssetDetailRequest() (request *DescribeDSPARDBDataAssetDetailRequest) {
    request = &DescribeDSPARDBDataAssetDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPARDBDataAssetDetail")
    
    
    return
}

func NewDescribeDSPARDBDataAssetDetailResponse() (response *DescribeDSPARDBDataAssetDetailResponse) {
    response = &DescribeDSPARDBDataAssetDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPARDBDataAssetDetail
// 获取RDB关系数据库分类分级资产详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDSPARDBDataAssetDetail(request *DescribeDSPARDBDataAssetDetailRequest) (response *DescribeDSPARDBDataAssetDetailResponse, err error) {
    return c.DescribeDSPARDBDataAssetDetailWithContext(context.Background(), request)
}

// DescribeDSPARDBDataAssetDetail
// 获取RDB关系数据库分类分级资产详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeDSPARDBDataAssetDetailWithContext(ctx context.Context, request *DescribeDSPARDBDataAssetDetailRequest) (response *DescribeDSPARDBDataAssetDetailResponse, err error) {
    if request == nil {
        request = NewDescribeDSPARDBDataAssetDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPARDBDataAssetDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPARDBDataAssetDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPARDBDataAssetDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPASupportedMetasRequest() (request *DescribeDSPASupportedMetasRequest) {
    request = &DescribeDSPASupportedMetasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPASupportedMetas")
    
    
    return
}

func NewDescribeDSPASupportedMetasResponse() (response *DescribeDSPASupportedMetasResponse) {
    response = &DescribeDSPASupportedMetasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPASupportedMetas
// 拉取DSPA支持的Meta元数据类型，返回包括：元数据类型，支持的元数据地域信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDSPASupportedMetas(request *DescribeDSPASupportedMetasRequest) (response *DescribeDSPASupportedMetasResponse, err error) {
    return c.DescribeDSPASupportedMetasWithContext(context.Background(), request)
}

// DescribeDSPASupportedMetas
// 拉取DSPA支持的Meta元数据类型，返回包括：元数据类型，支持的元数据地域信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDSPASupportedMetasWithContext(ctx context.Context, request *DescribeDSPASupportedMetasRequest) (response *DescribeDSPASupportedMetasResponse, err error) {
    if request == nil {
        request = NewDescribeDSPASupportedMetasRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPASupportedMetas")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPASupportedMetas require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPASupportedMetasResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDSPATaskResultDataSampleRequest() (request *DescribeDSPATaskResultDataSampleRequest) {
    request = &DescribeDSPATaskResultDataSampleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeDSPATaskResultDataSample")
    
    
    return
}

func NewDescribeDSPATaskResultDataSampleResponse() (response *DescribeDSPATaskResultDataSampleResponse) {
    response = &DescribeDSPATaskResultDataSampleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDSPATaskResultDataSample
// 获取扫描结果数据样本
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDSPATaskResultDataSample(request *DescribeDSPATaskResultDataSampleRequest) (response *DescribeDSPATaskResultDataSampleResponse, err error) {
    return c.DescribeDSPATaskResultDataSampleWithContext(context.Background(), request)
}

// DescribeDSPATaskResultDataSample
// 获取扫描结果数据样本
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDSPATaskResultDataSampleWithContext(ctx context.Context, request *DescribeDSPATaskResultDataSampleRequest) (response *DescribeDSPATaskResultDataSampleResponse, err error) {
    if request == nil {
        request = NewDescribeDSPATaskResultDataSampleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeDSPATaskResultDataSample")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDSPATaskResultDataSample require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDSPATaskResultDataSampleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeESAssetSensitiveDistributionRequest() (request *DescribeESAssetSensitiveDistributionRequest) {
    request = &DescribeESAssetSensitiveDistributionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeESAssetSensitiveDistribution")
    
    
    return
}

func NewDescribeESAssetSensitiveDistributionResponse() (response *DescribeESAssetSensitiveDistributionResponse) {
    response = &DescribeESAssetSensitiveDistributionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeESAssetSensitiveDistribution
// 数据资产报告-查询es的敏感资产报告，包含（数据库资产，设敏级别数据库top10，资产详情）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeESAssetSensitiveDistribution(request *DescribeESAssetSensitiveDistributionRequest) (response *DescribeESAssetSensitiveDistributionResponse, err error) {
    return c.DescribeESAssetSensitiveDistributionWithContext(context.Background(), request)
}

// DescribeESAssetSensitiveDistribution
// 数据资产报告-查询es的敏感资产报告，包含（数据库资产，设敏级别数据库top10，资产详情）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeESAssetSensitiveDistributionWithContext(ctx context.Context, request *DescribeESAssetSensitiveDistributionRequest) (response *DescribeESAssetSensitiveDistributionResponse, err error) {
    if request == nil {
        request = NewDescribeESAssetSensitiveDistributionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeESAssetSensitiveDistribution")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeESAssetSensitiveDistribution require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeESAssetSensitiveDistributionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExportTaskResultRequest() (request *DescribeExportTaskResultRequest) {
    request = &DescribeExportTaskResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeExportTaskResult")
    
    
    return
}

func NewDescribeExportTaskResultResponse() (response *DescribeExportTaskResultResponse) {
    response = &DescribeExportTaskResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeExportTaskResult
// 获取导出任务结果
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeExportTaskResult(request *DescribeExportTaskResultRequest) (response *DescribeExportTaskResultResponse, err error) {
    return c.DescribeExportTaskResultWithContext(context.Background(), request)
}

// DescribeExportTaskResult
// 获取导出任务结果
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeExportTaskResultWithContext(ctx context.Context, request *DescribeExportTaskResultRequest) (response *DescribeExportTaskResultResponse, err error) {
    if request == nil {
        request = NewDescribeExportTaskResultRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeExportTaskResult")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExportTaskResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExportTaskResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMongoAssetSensitiveDistributionRequest() (request *DescribeMongoAssetSensitiveDistributionRequest) {
    request = &DescribeMongoAssetSensitiveDistributionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeMongoAssetSensitiveDistribution")
    
    
    return
}

func NewDescribeMongoAssetSensitiveDistributionResponse() (response *DescribeMongoAssetSensitiveDistributionResponse) {
    response = &DescribeMongoAssetSensitiveDistributionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMongoAssetSensitiveDistribution
// 数据资产报告-查询mongo 的敏感资产报告，包含（数据库资产，设敏级别数据库top10，资产详情）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeMongoAssetSensitiveDistribution(request *DescribeMongoAssetSensitiveDistributionRequest) (response *DescribeMongoAssetSensitiveDistributionResponse, err error) {
    return c.DescribeMongoAssetSensitiveDistributionWithContext(context.Background(), request)
}

// DescribeMongoAssetSensitiveDistribution
// 数据资产报告-查询mongo 的敏感资产报告，包含（数据库资产，设敏级别数据库top10，资产详情）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeMongoAssetSensitiveDistributionWithContext(ctx context.Context, request *DescribeMongoAssetSensitiveDistributionRequest) (response *DescribeMongoAssetSensitiveDistributionResponse, err error) {
    if request == nil {
        request = NewDescribeMongoAssetSensitiveDistributionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeMongoAssetSensitiveDistribution")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMongoAssetSensitiveDistribution require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMongoAssetSensitiveDistributionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRDBAssetSensitiveDistributionRequest() (request *DescribeRDBAssetSensitiveDistributionRequest) {
    request = &DescribeRDBAssetSensitiveDistributionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeRDBAssetSensitiveDistribution")
    
    
    return
}

func NewDescribeRDBAssetSensitiveDistributionResponse() (response *DescribeRDBAssetSensitiveDistributionResponse) {
    response = &DescribeRDBAssetSensitiveDistributionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRDBAssetSensitiveDistribution
// 数据资产报告-查询rbd 的敏感资产报告，包含（数据库资产，设敏级别数据库top10，资产详情）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeRDBAssetSensitiveDistribution(request *DescribeRDBAssetSensitiveDistributionRequest) (response *DescribeRDBAssetSensitiveDistributionResponse, err error) {
    return c.DescribeRDBAssetSensitiveDistributionWithContext(context.Background(), request)
}

// DescribeRDBAssetSensitiveDistribution
// 数据资产报告-查询rbd 的敏感资产报告，包含（数据库资产，设敏级别数据库top10，资产详情）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeRDBAssetSensitiveDistributionWithContext(ctx context.Context, request *DescribeRDBAssetSensitiveDistributionRequest) (response *DescribeRDBAssetSensitiveDistributionResponse, err error) {
    if request == nil {
        request = NewDescribeRDBAssetSensitiveDistributionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeRDBAssetSensitiveDistribution")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRDBAssetSensitiveDistribution require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRDBAssetSensitiveDistributionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReportTaskDownloadUrlRequest() (request *DescribeReportTaskDownloadUrlRequest) {
    request = &DescribeReportTaskDownloadUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeReportTaskDownloadUrl")
    
    
    return
}

func NewDescribeReportTaskDownloadUrlResponse() (response *DescribeReportTaskDownloadUrlResponse) {
    response = &DescribeReportTaskDownloadUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeReportTaskDownloadUrl
// 获取报表下载链接
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeReportTaskDownloadUrl(request *DescribeReportTaskDownloadUrlRequest) (response *DescribeReportTaskDownloadUrlResponse, err error) {
    return c.DescribeReportTaskDownloadUrlWithContext(context.Background(), request)
}

// DescribeReportTaskDownloadUrl
// 获取报表下载链接
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeReportTaskDownloadUrlWithContext(ctx context.Context, request *DescribeReportTaskDownloadUrlRequest) (response *DescribeReportTaskDownloadUrlResponse, err error) {
    if request == nil {
        request = NewDescribeReportTaskDownloadUrlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeReportTaskDownloadUrl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReportTaskDownloadUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReportTaskDownloadUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReportTasksRequest() (request *DescribeReportTasksRequest) {
    request = &DescribeReportTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeReportTasks")
    
    
    return
}

func NewDescribeReportTasksResponse() (response *DescribeReportTasksResponse) {
    response = &DescribeReportTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeReportTasks
// 获取资产报表任务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeReportTasks(request *DescribeReportTasksRequest) (response *DescribeReportTasksResponse, err error) {
    return c.DescribeReportTasksWithContext(context.Background(), request)
}

// DescribeReportTasks
// 获取资产报表任务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeReportTasksWithContext(ctx context.Context, request *DescribeReportTasksRequest) (response *DescribeReportTasksResponse, err error) {
    if request == nil {
        request = NewDescribeReportTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeReportTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReportTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReportTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSensitiveCOSDataDistributionRequest() (request *DescribeSensitiveCOSDataDistributionRequest) {
    request = &DescribeSensitiveCOSDataDistributionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeSensitiveCOSDataDistribution")
    
    
    return
}

func NewDescribeSensitiveCOSDataDistributionResponse() (response *DescribeSensitiveCOSDataDistributionResponse) {
    response = &DescribeSensitiveCOSDataDistributionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSensitiveCOSDataDistribution
// 数据资产报告-查询cos的敏感数据分布（分级分布 分类分布 敏感规则分布详情列表）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeSensitiveCOSDataDistribution(request *DescribeSensitiveCOSDataDistributionRequest) (response *DescribeSensitiveCOSDataDistributionResponse, err error) {
    return c.DescribeSensitiveCOSDataDistributionWithContext(context.Background(), request)
}

// DescribeSensitiveCOSDataDistribution
// 数据资产报告-查询cos的敏感数据分布（分级分布 分类分布 敏感规则分布详情列表）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeSensitiveCOSDataDistributionWithContext(ctx context.Context, request *DescribeSensitiveCOSDataDistributionRequest) (response *DescribeSensitiveCOSDataDistributionResponse, err error) {
    if request == nil {
        request = NewDescribeSensitiveCOSDataDistributionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeSensitiveCOSDataDistribution")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSensitiveCOSDataDistribution require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSensitiveCOSDataDistributionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSensitiveRDBDataDistributionRequest() (request *DescribeSensitiveRDBDataDistributionRequest) {
    request = &DescribeSensitiveRDBDataDistributionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DescribeSensitiveRDBDataDistribution")
    
    
    return
}

func NewDescribeSensitiveRDBDataDistributionResponse() (response *DescribeSensitiveRDBDataDistributionResponse) {
    response = &DescribeSensitiveRDBDataDistributionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSensitiveRDBDataDistribution
// 数据资产报告-查询rdb的敏感数据分布-敏感规则字段分布-敏感分布详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeSensitiveRDBDataDistribution(request *DescribeSensitiveRDBDataDistributionRequest) (response *DescribeSensitiveRDBDataDistributionResponse, err error) {
    return c.DescribeSensitiveRDBDataDistributionWithContext(context.Background(), request)
}

// DescribeSensitiveRDBDataDistribution
// 数据资产报告-查询rdb的敏感数据分布-敏感规则字段分布-敏感分布详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeSensitiveRDBDataDistributionWithContext(ctx context.Context, request *DescribeSensitiveRDBDataDistributionRequest) (response *DescribeSensitiveRDBDataDistributionResponse, err error) {
    if request == nil {
        request = NewDescribeSensitiveRDBDataDistributionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DescribeSensitiveRDBDataDistribution")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSensitiveRDBDataDistribution require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSensitiveRDBDataDistributionResponse()
    err = c.Send(request, response)
    return
}

func NewDisableDSPAMetaResourceAuthRequest() (request *DisableDSPAMetaResourceAuthRequest) {
    request = &DisableDSPAMetaResourceAuthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "DisableDSPAMetaResourceAuth")
    
    
    return
}

func NewDisableDSPAMetaResourceAuthResponse() (response *DisableDSPAMetaResourceAuthResponse) {
    response = &DisableDSPAMetaResourceAuthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisableDSPAMetaResourceAuth
// 取消用户云资源授权
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DisableDSPAMetaResourceAuth(request *DisableDSPAMetaResourceAuthRequest) (response *DisableDSPAMetaResourceAuthResponse, err error) {
    return c.DisableDSPAMetaResourceAuthWithContext(context.Background(), request)
}

// DisableDSPAMetaResourceAuth
// 取消用户云资源授权
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DisableDSPAMetaResourceAuthWithContext(ctx context.Context, request *DisableDSPAMetaResourceAuthRequest) (response *DisableDSPAMetaResourceAuthResponse, err error) {
    if request == nil {
        request = NewDisableDSPAMetaResourceAuthRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "DisableDSPAMetaResourceAuth")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableDSPAMetaResourceAuth require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableDSPAMetaResourceAuthResponse()
    err = c.Send(request, response)
    return
}

func NewEnableDSPADiscoveryRuleRequest() (request *EnableDSPADiscoveryRuleRequest) {
    request = &EnableDSPADiscoveryRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "EnableDSPADiscoveryRule")
    
    
    return
}

func NewEnableDSPADiscoveryRuleResponse() (response *EnableDSPADiscoveryRuleResponse) {
    response = &EnableDSPADiscoveryRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnableDSPADiscoveryRule
// 打开或者关闭分类分级规则
//
// 注：此API同时对该规则下的RDB跟COS规则操作。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) EnableDSPADiscoveryRule(request *EnableDSPADiscoveryRuleRequest) (response *EnableDSPADiscoveryRuleResponse, err error) {
    return c.EnableDSPADiscoveryRuleWithContext(context.Background(), request)
}

// EnableDSPADiscoveryRule
// 打开或者关闭分类分级规则
//
// 注：此API同时对该规则下的RDB跟COS规则操作。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) EnableDSPADiscoveryRuleWithContext(ctx context.Context, request *EnableDSPADiscoveryRuleRequest) (response *EnableDSPADiscoveryRuleResponse, err error) {
    if request == nil {
        request = NewEnableDSPADiscoveryRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "EnableDSPADiscoveryRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableDSPADiscoveryRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableDSPADiscoveryRuleResponse()
    err = c.Send(request, response)
    return
}

func NewEnableTrialVersionRequest() (request *EnableTrialVersionRequest) {
    request = &EnableTrialVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "EnableTrialVersion")
    
    
    return
}

func NewEnableTrialVersionResponse() (response *EnableTrialVersionResponse) {
    response = &EnableTrialVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnableTrialVersion
// 启用版本体验
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) EnableTrialVersion(request *EnableTrialVersionRequest) (response *EnableTrialVersionResponse, err error) {
    return c.EnableTrialVersionWithContext(context.Background(), request)
}

// EnableTrialVersion
// 启用版本体验
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) EnableTrialVersionWithContext(ctx context.Context, request *EnableTrialVersionRequest) (response *EnableTrialVersionResponse, err error) {
    if request == nil {
        request = NewEnableTrialVersionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "EnableTrialVersion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableTrialVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableTrialVersionResponse()
    err = c.Send(request, response)
    return
}

func NewExportAssetDetailDataRequest() (request *ExportAssetDetailDataRequest) {
    request = &ExportAssetDetailDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "ExportAssetDetailData")
    
    
    return
}

func NewExportAssetDetailDataResponse() (response *ExportAssetDetailDataResponse) {
    response = &ExportAssetDetailDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ExportAssetDetailData
// 创建敏感数据导出任务
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ExportAssetDetailData(request *ExportAssetDetailDataRequest) (response *ExportAssetDetailDataResponse, err error) {
    return c.ExportAssetDetailDataWithContext(context.Background(), request)
}

// ExportAssetDetailData
// 创建敏感数据导出任务
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ExportAssetDetailDataWithContext(ctx context.Context, request *ExportAssetDetailDataRequest) (response *ExportAssetDetailDataResponse, err error) {
    if request == nil {
        request = NewExportAssetDetailDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "ExportAssetDetailData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportAssetDetailData require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportAssetDetailDataResponse()
    err = c.Send(request, response)
    return
}

func NewGetResourceConnectionStatusRequest() (request *GetResourceConnectionStatusRequest) {
    request = &GetResourceConnectionStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "GetResourceConnectionStatus")
    
    
    return
}

func NewGetResourceConnectionStatusResponse() (response *GetResourceConnectionStatusResponse) {
    response = &GetResourceConnectionStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetResourceConnectionStatus
// 获取授权资源的连接状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetResourceConnectionStatus(request *GetResourceConnectionStatusRequest) (response *GetResourceConnectionStatusResponse, err error) {
    return c.GetResourceConnectionStatusWithContext(context.Background(), request)
}

// GetResourceConnectionStatus
// 获取授权资源的连接状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetResourceConnectionStatusWithContext(ctx context.Context, request *GetResourceConnectionStatusRequest) (response *GetResourceConnectionStatusResponse, err error) {
    if request == nil {
        request = NewGetResourceConnectionStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "GetResourceConnectionStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetResourceConnectionStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetResourceConnectionStatusResponse()
    err = c.Send(request, response)
    return
}

func NewGetTrialVersionRequest() (request *GetTrialVersionRequest) {
    request = &GetTrialVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "GetTrialVersion")
    
    
    return
}

func NewGetTrialVersionResponse() (response *GetTrialVersionResponse) {
    response = &GetTrialVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetTrialVersion
// 获取体验版本信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetTrialVersion(request *GetTrialVersionRequest) (response *GetTrialVersionResponse, err error) {
    return c.GetTrialVersionWithContext(context.Background(), request)
}

// GetTrialVersion
// 获取体验版本信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetTrialVersionWithContext(ctx context.Context, request *GetTrialVersionRequest) (response *GetTrialVersionResponse, err error) {
    if request == nil {
        request = NewGetTrialVersionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "GetTrialVersion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTrialVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTrialVersionResponse()
    err = c.Send(request, response)
    return
}

func NewGetUserQuotaInfoRequest() (request *GetUserQuotaInfoRequest) {
    request = &GetUserQuotaInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "GetUserQuotaInfo")
    
    
    return
}

func NewGetUserQuotaInfoResponse() (response *GetUserQuotaInfoResponse) {
    response = &GetUserQuotaInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetUserQuotaInfo
// 获取用户购买配额信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetUserQuotaInfo(request *GetUserQuotaInfoRequest) (response *GetUserQuotaInfoResponse, err error) {
    return c.GetUserQuotaInfoWithContext(context.Background(), request)
}

// GetUserQuotaInfo
// 获取用户购买配额信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetUserQuotaInfoWithContext(ctx context.Context, request *GetUserQuotaInfoRequest) (response *GetUserQuotaInfoResponse, err error) {
    if request == nil {
        request = NewGetUserQuotaInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "GetUserQuotaInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetUserQuotaInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetUserQuotaInfoResponse()
    err = c.Send(request, response)
    return
}

func NewListDSPAClustersRequest() (request *ListDSPAClustersRequest) {
    request = &ListDSPAClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "ListDSPAClusters")
    
    
    return
}

func NewListDSPAClustersResponse() (response *ListDSPAClustersResponse) {
    response = &ListDSPAClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListDSPAClusters
// 拉取DSPA集群列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ListDSPAClusters(request *ListDSPAClustersRequest) (response *ListDSPAClustersResponse, err error) {
    return c.ListDSPAClustersWithContext(context.Background(), request)
}

// ListDSPAClusters
// 拉取DSPA集群列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ListDSPAClustersWithContext(ctx context.Context, request *ListDSPAClustersRequest) (response *ListDSPAClustersResponse, err error) {
    if request == nil {
        request = NewListDSPAClustersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "ListDSPAClusters")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListDSPAClusters require credential")
    }

    request.SetContext(ctx)
    
    response = NewListDSPAClustersResponse()
    err = c.Send(request, response)
    return
}

func NewListDSPACosMetaResourcesRequest() (request *ListDSPACosMetaResourcesRequest) {
    request = &ListDSPACosMetaResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "ListDSPACosMetaResources")
    
    
    return
}

func NewListDSPACosMetaResourcesResponse() (response *ListDSPACosMetaResourcesResponse) {
    response = &ListDSPACosMetaResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListDSPACosMetaResources
// 本接口用于获取COS元数据信息，返回COS元数据信息列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ListDSPACosMetaResources(request *ListDSPACosMetaResourcesRequest) (response *ListDSPACosMetaResourcesResponse, err error) {
    return c.ListDSPACosMetaResourcesWithContext(context.Background(), request)
}

// ListDSPACosMetaResources
// 本接口用于获取COS元数据信息，返回COS元数据信息列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ListDSPACosMetaResourcesWithContext(ctx context.Context, request *ListDSPACosMetaResourcesRequest) (response *ListDSPACosMetaResourcesResponse, err error) {
    if request == nil {
        request = NewListDSPACosMetaResourcesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "ListDSPACosMetaResources")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListDSPACosMetaResources require credential")
    }

    request.SetContext(ctx)
    
    response = NewListDSPACosMetaResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewListDSPAMetaResourcesRequest() (request *ListDSPAMetaResourcesRequest) {
    request = &ListDSPAMetaResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "ListDSPAMetaResources")
    
    
    return
}

func NewListDSPAMetaResourcesResponse() (response *ListDSPAMetaResourcesResponse) {
    response = &ListDSPAMetaResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListDSPAMetaResources
// 拉取用户云资源
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ListDSPAMetaResources(request *ListDSPAMetaResourcesRequest) (response *ListDSPAMetaResourcesResponse, err error) {
    return c.ListDSPAMetaResourcesWithContext(context.Background(), request)
}

// ListDSPAMetaResources
// 拉取用户云资源
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ListDSPAMetaResourcesWithContext(ctx context.Context, request *ListDSPAMetaResourcesRequest) (response *ListDSPAMetaResourcesResponse, err error) {
    if request == nil {
        request = NewListDSPAMetaResourcesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "ListDSPAMetaResources")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListDSPAMetaResources require credential")
    }

    request.SetContext(ctx)
    
    response = NewListDSPAMetaResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDSPAAssessmentRiskRequest() (request *ModifyDSPAAssessmentRiskRequest) {
    request = &ModifyDSPAAssessmentRiskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "ModifyDSPAAssessmentRisk")
    
    
    return
}

func NewModifyDSPAAssessmentRiskResponse() (response *ModifyDSPAAssessmentRiskResponse) {
    response = &ModifyDSPAAssessmentRiskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDSPAAssessmentRisk
// 修改DSPA评估风险项，支持修改Status
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDSPAAssessmentRisk(request *ModifyDSPAAssessmentRiskRequest) (response *ModifyDSPAAssessmentRiskResponse, err error) {
    return c.ModifyDSPAAssessmentRiskWithContext(context.Background(), request)
}

// ModifyDSPAAssessmentRisk
// 修改DSPA评估风险项，支持修改Status
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDSPAAssessmentRiskWithContext(ctx context.Context, request *ModifyDSPAAssessmentRiskRequest) (response *ModifyDSPAAssessmentRiskResponse, err error) {
    if request == nil {
        request = NewModifyDSPAAssessmentRiskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "ModifyDSPAAssessmentRisk")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDSPAAssessmentRisk require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDSPAAssessmentRiskResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDSPAAssessmentRiskLatestRequest() (request *ModifyDSPAAssessmentRiskLatestRequest) {
    request = &ModifyDSPAAssessmentRiskLatestRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "ModifyDSPAAssessmentRiskLatest")
    
    
    return
}

func NewModifyDSPAAssessmentRiskLatestResponse() (response *ModifyDSPAAssessmentRiskLatestResponse) {
    response = &ModifyDSPAAssessmentRiskLatestResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDSPAAssessmentRiskLatest
// 修改最新评估风险项状态
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyDSPAAssessmentRiskLatest(request *ModifyDSPAAssessmentRiskLatestRequest) (response *ModifyDSPAAssessmentRiskLatestResponse, err error) {
    return c.ModifyDSPAAssessmentRiskLatestWithContext(context.Background(), request)
}

// ModifyDSPAAssessmentRiskLatest
// 修改最新评估风险项状态
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyDSPAAssessmentRiskLatestWithContext(ctx context.Context, request *ModifyDSPAAssessmentRiskLatestRequest) (response *ModifyDSPAAssessmentRiskLatestResponse, err error) {
    if request == nil {
        request = NewModifyDSPAAssessmentRiskLatestRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "ModifyDSPAAssessmentRiskLatest")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDSPAAssessmentRiskLatest require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDSPAAssessmentRiskLatestResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDSPAAssessmentRiskLevelRequest() (request *ModifyDSPAAssessmentRiskLevelRequest) {
    request = &ModifyDSPAAssessmentRiskLevelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "ModifyDSPAAssessmentRiskLevel")
    
    
    return
}

func NewModifyDSPAAssessmentRiskLevelResponse() (response *ModifyDSPAAssessmentRiskLevelResponse) {
    response = &ModifyDSPAAssessmentRiskLevelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDSPAAssessmentRiskLevel
// 风险项页面----修改风险等级的详情数据
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyDSPAAssessmentRiskLevel(request *ModifyDSPAAssessmentRiskLevelRequest) (response *ModifyDSPAAssessmentRiskLevelResponse, err error) {
    return c.ModifyDSPAAssessmentRiskLevelWithContext(context.Background(), request)
}

// ModifyDSPAAssessmentRiskLevel
// 风险项页面----修改风险等级的详情数据
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyDSPAAssessmentRiskLevelWithContext(ctx context.Context, request *ModifyDSPAAssessmentRiskLevelRequest) (response *ModifyDSPAAssessmentRiskLevelResponse, err error) {
    if request == nil {
        request = NewModifyDSPAAssessmentRiskLevelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "ModifyDSPAAssessmentRiskLevel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDSPAAssessmentRiskLevel require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDSPAAssessmentRiskLevelResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDSPAAssessmentRiskTemplateRequest() (request *ModifyDSPAAssessmentRiskTemplateRequest) {
    request = &ModifyDSPAAssessmentRiskTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "ModifyDSPAAssessmentRiskTemplate")
    
    
    return
}

func NewModifyDSPAAssessmentRiskTemplateResponse() (response *ModifyDSPAAssessmentRiskTemplateResponse) {
    response = &ModifyDSPAAssessmentRiskTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDSPAAssessmentRiskTemplate
// 风险模板---修改风险模板
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyDSPAAssessmentRiskTemplate(request *ModifyDSPAAssessmentRiskTemplateRequest) (response *ModifyDSPAAssessmentRiskTemplateResponse, err error) {
    return c.ModifyDSPAAssessmentRiskTemplateWithContext(context.Background(), request)
}

// ModifyDSPAAssessmentRiskTemplate
// 风险模板---修改风险模板
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyDSPAAssessmentRiskTemplateWithContext(ctx context.Context, request *ModifyDSPAAssessmentRiskTemplateRequest) (response *ModifyDSPAAssessmentRiskTemplateResponse, err error) {
    if request == nil {
        request = NewModifyDSPAAssessmentRiskTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "ModifyDSPAAssessmentRiskTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDSPAAssessmentRiskTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDSPAAssessmentRiskTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDSPACOSDiscoveryTaskRequest() (request *ModifyDSPACOSDiscoveryTaskRequest) {
    request = &ModifyDSPACOSDiscoveryTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "ModifyDSPACOSDiscoveryTask")
    
    
    return
}

func NewModifyDSPACOSDiscoveryTaskResponse() (response *ModifyDSPACOSDiscoveryTaskResponse) {
    response = &ModifyDSPACOSDiscoveryTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDSPACOSDiscoveryTask
// 修改COS分类分级任务，该接口只有在任务状态为以下状态时才支持正确修改：
//
// 0 待扫描，
//
// 2 扫描终止， 
//
// 3 扫描成功，
//
// 4 扫描失败
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDSPACOSDiscoveryTask(request *ModifyDSPACOSDiscoveryTaskRequest) (response *ModifyDSPACOSDiscoveryTaskResponse, err error) {
    return c.ModifyDSPACOSDiscoveryTaskWithContext(context.Background(), request)
}

// ModifyDSPACOSDiscoveryTask
// 修改COS分类分级任务，该接口只有在任务状态为以下状态时才支持正确修改：
//
// 0 待扫描，
//
// 2 扫描终止， 
//
// 3 扫描成功，
//
// 4 扫描失败
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDSPACOSDiscoveryTaskWithContext(ctx context.Context, request *ModifyDSPACOSDiscoveryTaskRequest) (response *ModifyDSPACOSDiscoveryTaskResponse, err error) {
    if request == nil {
        request = NewModifyDSPACOSDiscoveryTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "ModifyDSPACOSDiscoveryTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDSPACOSDiscoveryTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDSPACOSDiscoveryTaskResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDSPACOSTaskResultRequest() (request *ModifyDSPACOSTaskResultRequest) {
    request = &ModifyDSPACOSTaskResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "ModifyDSPACOSTaskResult")
    
    
    return
}

func NewModifyDSPACOSTaskResultResponse() (response *ModifyDSPACOSTaskResultResponse) {
    response = &ModifyDSPACOSTaskResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDSPACOSTaskResult
// 调整COS任务扫描结果
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyDSPACOSTaskResult(request *ModifyDSPACOSTaskResultRequest) (response *ModifyDSPACOSTaskResultResponse, err error) {
    return c.ModifyDSPACOSTaskResultWithContext(context.Background(), request)
}

// ModifyDSPACOSTaskResult
// 调整COS任务扫描结果
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyDSPACOSTaskResultWithContext(ctx context.Context, request *ModifyDSPACOSTaskResultRequest) (response *ModifyDSPACOSTaskResultResponse, err error) {
    if request == nil {
        request = NewModifyDSPACOSTaskResultRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "ModifyDSPACOSTaskResult")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDSPACOSTaskResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDSPACOSTaskResultResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDSPACategoryRequest() (request *ModifyDSPACategoryRequest) {
    request = &ModifyDSPACategoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "ModifyDSPACategory")
    
    
    return
}

func NewModifyDSPACategoryResponse() (response *ModifyDSPACategoryResponse) {
    response = &ModifyDSPACategoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDSPACategory
// 修改分类，内置分类不支持修改。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDSPACategory(request *ModifyDSPACategoryRequest) (response *ModifyDSPACategoryResponse, err error) {
    return c.ModifyDSPACategoryWithContext(context.Background(), request)
}

// ModifyDSPACategory
// 修改分类，内置分类不支持修改。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDSPACategoryWithContext(ctx context.Context, request *ModifyDSPACategoryRequest) (response *ModifyDSPACategoryResponse, err error) {
    if request == nil {
        request = NewModifyDSPACategoryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "ModifyDSPACategory")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDSPACategory require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDSPACategoryResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDSPACategoryRelationRequest() (request *ModifyDSPACategoryRelationRequest) {
    request = &ModifyDSPACategoryRelationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "ModifyDSPACategoryRelation")
    
    
    return
}

func NewModifyDSPACategoryRelationResponse() (response *ModifyDSPACategoryRelationResponse) {
    response = &ModifyDSPACategoryRelationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDSPACategoryRelation
// 修改分类分级关系
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDSPACategoryRelation(request *ModifyDSPACategoryRelationRequest) (response *ModifyDSPACategoryRelationResponse, err error) {
    return c.ModifyDSPACategoryRelationWithContext(context.Background(), request)
}

// ModifyDSPACategoryRelation
// 修改分类分级关系
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDSPACategoryRelationWithContext(ctx context.Context, request *ModifyDSPACategoryRelationRequest) (response *ModifyDSPACategoryRelationResponse, err error) {
    if request == nil {
        request = NewModifyDSPACategoryRelationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "ModifyDSPACategoryRelation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDSPACategoryRelation require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDSPACategoryRelationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDSPAClusterInfoRequest() (request *ModifyDSPAClusterInfoRequest) {
    request = &ModifyDSPAClusterInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "ModifyDSPAClusterInfo")
    
    
    return
}

func NewModifyDSPAClusterInfoResponse() (response *ModifyDSPAClusterInfoResponse) {
    response = &ModifyDSPAClusterInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDSPAClusterInfo
// 修改DSPA集群信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyDSPAClusterInfo(request *ModifyDSPAClusterInfoRequest) (response *ModifyDSPAClusterInfoResponse, err error) {
    return c.ModifyDSPAClusterInfoWithContext(context.Background(), request)
}

// ModifyDSPAClusterInfo
// 修改DSPA集群信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyDSPAClusterInfoWithContext(ctx context.Context, request *ModifyDSPAClusterInfoRequest) (response *ModifyDSPAClusterInfoResponse, err error) {
    if request == nil {
        request = NewModifyDSPAClusterInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "ModifyDSPAClusterInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDSPAClusterInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDSPAClusterInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDSPAComplianceGroupRequest() (request *ModifyDSPAComplianceGroupRequest) {
    request = &ModifyDSPAComplianceGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "ModifyDSPAComplianceGroup")
    
    
    return
}

func NewModifyDSPAComplianceGroupResponse() (response *ModifyDSPAComplianceGroupResponse) {
    response = &ModifyDSPAComplianceGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDSPAComplianceGroup
// 修改分类分级模板，内置模板不支持修改。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDSPAComplianceGroup(request *ModifyDSPAComplianceGroupRequest) (response *ModifyDSPAComplianceGroupResponse, err error) {
    return c.ModifyDSPAComplianceGroupWithContext(context.Background(), request)
}

// ModifyDSPAComplianceGroup
// 修改分类分级模板，内置模板不支持修改。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDSPAComplianceGroupWithContext(ctx context.Context, request *ModifyDSPAComplianceGroupRequest) (response *ModifyDSPAComplianceGroupResponse, err error) {
    if request == nil {
        request = NewModifyDSPAComplianceGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "ModifyDSPAComplianceGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDSPAComplianceGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDSPAComplianceGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDSPADiscoveryRuleRequest() (request *ModifyDSPADiscoveryRuleRequest) {
    request = &ModifyDSPADiscoveryRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "ModifyDSPADiscoveryRule")
    
    
    return
}

func NewModifyDSPADiscoveryRuleResponse() (response *ModifyDSPADiscoveryRuleResponse) {
    response = &ModifyDSPADiscoveryRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDSPADiscoveryRule
// 修改分类分级规则，单个用户最多允许创建200个规则。
//
// 注：此API同时适用RDB跟COS类型数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyDSPADiscoveryRule(request *ModifyDSPADiscoveryRuleRequest) (response *ModifyDSPADiscoveryRuleResponse, err error) {
    return c.ModifyDSPADiscoveryRuleWithContext(context.Background(), request)
}

// ModifyDSPADiscoveryRule
// 修改分类分级规则，单个用户最多允许创建200个规则。
//
// 注：此API同时适用RDB跟COS类型数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyDSPADiscoveryRuleWithContext(ctx context.Context, request *ModifyDSPADiscoveryRuleRequest) (response *ModifyDSPADiscoveryRuleResponse, err error) {
    if request == nil {
        request = NewModifyDSPADiscoveryRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "ModifyDSPADiscoveryRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDSPADiscoveryRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDSPADiscoveryRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDSPADiscoveryTaskRequest() (request *ModifyDSPADiscoveryTaskRequest) {
    request = &ModifyDSPADiscoveryTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "ModifyDSPADiscoveryTask")
    
    
    return
}

func NewModifyDSPADiscoveryTaskResponse() (response *ModifyDSPADiscoveryTaskResponse) {
    response = &ModifyDSPADiscoveryTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDSPADiscoveryTask
// 修改分类分级任务，该接口只有在任务状态为以下状态时才支持正确修改：
//
// 0 待扫描，
//
// 2 扫描终止， 
//
// 3 扫描成功，
//
// 4 扫描失败
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDSPADiscoveryTask(request *ModifyDSPADiscoveryTaskRequest) (response *ModifyDSPADiscoveryTaskResponse, err error) {
    return c.ModifyDSPADiscoveryTaskWithContext(context.Background(), request)
}

// ModifyDSPADiscoveryTask
// 修改分类分级任务，该接口只有在任务状态为以下状态时才支持正确修改：
//
// 0 待扫描，
//
// 2 扫描终止， 
//
// 3 扫描成功，
//
// 4 扫描失败
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDSPADiscoveryTaskWithContext(ctx context.Context, request *ModifyDSPADiscoveryTaskRequest) (response *ModifyDSPADiscoveryTaskResponse, err error) {
    if request == nil {
        request = NewModifyDSPADiscoveryTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "ModifyDSPADiscoveryTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDSPADiscoveryTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDSPADiscoveryTaskResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDSPAESTaskResultRequest() (request *ModifyDSPAESTaskResultRequest) {
    request = &ModifyDSPAESTaskResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "ModifyDSPAESTaskResult")
    
    
    return
}

func NewModifyDSPAESTaskResultResponse() (response *ModifyDSPAESTaskResultResponse) {
    response = &ModifyDSPAESTaskResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDSPAESTaskResult
// 调整ES任务扫描结果
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyDSPAESTaskResult(request *ModifyDSPAESTaskResultRequest) (response *ModifyDSPAESTaskResultResponse, err error) {
    return c.ModifyDSPAESTaskResultWithContext(context.Background(), request)
}

// ModifyDSPAESTaskResult
// 调整ES任务扫描结果
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyDSPAESTaskResultWithContext(ctx context.Context, request *ModifyDSPAESTaskResultRequest) (response *ModifyDSPAESTaskResultResponse, err error) {
    if request == nil {
        request = NewModifyDSPAESTaskResultRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "ModifyDSPAESTaskResult")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDSPAESTaskResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDSPAESTaskResultResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDSPATaskResultRequest() (request *ModifyDSPATaskResultRequest) {
    request = &ModifyDSPATaskResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "ModifyDSPATaskResult")
    
    
    return
}

func NewModifyDSPATaskResultResponse() (response *ModifyDSPATaskResultResponse) {
    response = &ModifyDSPATaskResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDSPATaskResult
// 调整任务扫描结果
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION_MULTIPLEMARKINGNOTALLOWED = "UnsupportedOperation.MultipleMarkingNotAllowed"
func (c *Client) ModifyDSPATaskResult(request *ModifyDSPATaskResultRequest) (response *ModifyDSPATaskResultResponse, err error) {
    return c.ModifyDSPATaskResultWithContext(context.Background(), request)
}

// ModifyDSPATaskResult
// 调整任务扫描结果
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION_MULTIPLEMARKINGNOTALLOWED = "UnsupportedOperation.MultipleMarkingNotAllowed"
func (c *Client) ModifyDSPATaskResultWithContext(ctx context.Context, request *ModifyDSPATaskResultRequest) (response *ModifyDSPATaskResultResponse, err error) {
    if request == nil {
        request = NewModifyDSPATaskResultRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "ModifyDSPATaskResult")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDSPATaskResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDSPATaskResultResponse()
    err = c.Send(request, response)
    return
}

func NewQueryDSPAMetaResourceDbListRequest() (request *QueryDSPAMetaResourceDbListRequest) {
    request = &QueryDSPAMetaResourceDbListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "QueryDSPAMetaResourceDbList")
    
    
    return
}

func NewQueryDSPAMetaResourceDbListResponse() (response *QueryDSPAMetaResourceDbListResponse) {
    response = &QueryDSPAMetaResourceDbListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryDSPAMetaResourceDbList
// 查询DSPA实例的db列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) QueryDSPAMetaResourceDbList(request *QueryDSPAMetaResourceDbListRequest) (response *QueryDSPAMetaResourceDbListResponse, err error) {
    return c.QueryDSPAMetaResourceDbListWithContext(context.Background(), request)
}

// QueryDSPAMetaResourceDbList
// 查询DSPA实例的db列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) QueryDSPAMetaResourceDbListWithContext(ctx context.Context, request *QueryDSPAMetaResourceDbListRequest) (response *QueryDSPAMetaResourceDbListResponse, err error) {
    if request == nil {
        request = NewQueryDSPAMetaResourceDbListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "QueryDSPAMetaResourceDbList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryDSPAMetaResourceDbList require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryDSPAMetaResourceDbListResponse()
    err = c.Send(request, response)
    return
}

func NewQueryResourceDbBindStatusRequest() (request *QueryResourceDbBindStatusRequest) {
    request = &QueryResourceDbBindStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "QueryResourceDbBindStatus")
    
    
    return
}

func NewQueryResourceDbBindStatusResponse() (response *QueryResourceDbBindStatusResponse) {
    response = &QueryResourceDbBindStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryResourceDbBindStatus
// 获取资源绑定DB状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) QueryResourceDbBindStatus(request *QueryResourceDbBindStatusRequest) (response *QueryResourceDbBindStatusResponse, err error) {
    return c.QueryResourceDbBindStatusWithContext(context.Background(), request)
}

// QueryResourceDbBindStatus
// 获取资源绑定DB状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) QueryResourceDbBindStatusWithContext(ctx context.Context, request *QueryResourceDbBindStatusRequest) (response *QueryResourceDbBindStatusResponse, err error) {
    if request == nil {
        request = NewQueryResourceDbBindStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "QueryResourceDbBindStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryResourceDbBindStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryResourceDbBindStatusResponse()
    err = c.Send(request, response)
    return
}

func NewRestartDSPAAssessmentTaskRequest() (request *RestartDSPAAssessmentTaskRequest) {
    request = &RestartDSPAAssessmentTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "RestartDSPAAssessmentTask")
    
    
    return
}

func NewRestartDSPAAssessmentTaskResponse() (response *RestartDSPAAssessmentTaskResponse) {
    response = &RestartDSPAAssessmentTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RestartDSPAAssessmentTask
// 重新启动DSPA风险评估任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RestartDSPAAssessmentTask(request *RestartDSPAAssessmentTaskRequest) (response *RestartDSPAAssessmentTaskResponse, err error) {
    return c.RestartDSPAAssessmentTaskWithContext(context.Background(), request)
}

// RestartDSPAAssessmentTask
// 重新启动DSPA风险评估任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RestartDSPAAssessmentTaskWithContext(ctx context.Context, request *RestartDSPAAssessmentTaskRequest) (response *RestartDSPAAssessmentTaskResponse, err error) {
    if request == nil {
        request = NewRestartDSPAAssessmentTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "RestartDSPAAssessmentTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestartDSPAAssessmentTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewRestartDSPAAssessmentTaskResponse()
    err = c.Send(request, response)
    return
}

func NewStartDSPADiscoveryTaskRequest() (request *StartDSPADiscoveryTaskRequest) {
    request = &StartDSPADiscoveryTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "StartDSPADiscoveryTask")
    
    
    return
}

func NewStartDSPADiscoveryTaskResponse() (response *StartDSPADiscoveryTaskResponse) {
    response = &StartDSPADiscoveryTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartDSPADiscoveryTask
// 立即启动分类分级任务，该接口只有在任务状态为以下状态时才支持正确执行立即扫描：
//
// 0 待扫描，
//
// 2 扫描终止，
//
// 3 扫描成功，
//
// 4 扫描失败
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_TASKDISABLED = "ResourceUnavailable.TaskDisabled"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StartDSPADiscoveryTask(request *StartDSPADiscoveryTaskRequest) (response *StartDSPADiscoveryTaskResponse, err error) {
    return c.StartDSPADiscoveryTaskWithContext(context.Background(), request)
}

// StartDSPADiscoveryTask
// 立即启动分类分级任务，该接口只有在任务状态为以下状态时才支持正确执行立即扫描：
//
// 0 待扫描，
//
// 2 扫描终止，
//
// 3 扫描成功，
//
// 4 扫描失败
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_TASKDISABLED = "ResourceUnavailable.TaskDisabled"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StartDSPADiscoveryTaskWithContext(ctx context.Context, request *StartDSPADiscoveryTaskRequest) (response *StartDSPADiscoveryTaskResponse, err error) {
    if request == nil {
        request = NewStartDSPADiscoveryTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "StartDSPADiscoveryTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartDSPADiscoveryTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartDSPADiscoveryTaskResponse()
    err = c.Send(request, response)
    return
}

func NewStopDSPADiscoveryTaskRequest() (request *StopDSPADiscoveryTaskRequest) {
    request = &StopDSPADiscoveryTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "StopDSPADiscoveryTask")
    
    
    return
}

func NewStopDSPADiscoveryTaskResponse() (response *StopDSPADiscoveryTaskResponse) {
    response = &StopDSPADiscoveryTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopDSPADiscoveryTask
// 停止分类分级任务，该接口只有在任务状态为以下状态时才支持正确执行停止扫描：
//
// 1 扫描中
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StopDSPADiscoveryTask(request *StopDSPADiscoveryTaskRequest) (response *StopDSPADiscoveryTaskResponse, err error) {
    return c.StopDSPADiscoveryTaskWithContext(context.Background(), request)
}

// StopDSPADiscoveryTask
// 停止分类分级任务，该接口只有在任务状态为以下状态时才支持正确执行停止扫描：
//
// 1 扫描中
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StopDSPADiscoveryTaskWithContext(ctx context.Context, request *StopDSPADiscoveryTaskRequest) (response *StopDSPADiscoveryTaskResponse, err error) {
    if request == nil {
        request = NewStopDSPADiscoveryTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "StopDSPADiscoveryTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopDSPADiscoveryTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopDSPADiscoveryTaskResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateDSPASelfBuildResourceRequest() (request *UpdateDSPASelfBuildResourceRequest) {
    request = &UpdateDSPASelfBuildResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "UpdateDSPASelfBuildResource")
    
    
    return
}

func NewUpdateDSPASelfBuildResourceResponse() (response *UpdateDSPASelfBuildResourceResponse) {
    response = &UpdateDSPASelfBuildResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateDSPASelfBuildResource
// 更新自建资源基础信息，包括：端口、账户名、密码。
//
// 请注意：
//
// 如果资源自身的VPC、VIP信息发生变化，后台会自动更新。
//
// 如果监听的端口发生变化，请显式输入端口。
//
// 如果账户名密码任意一个发生变化，请务必同时显式输入账户名密码。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateDSPASelfBuildResource(request *UpdateDSPASelfBuildResourceRequest) (response *UpdateDSPASelfBuildResourceResponse, err error) {
    return c.UpdateDSPASelfBuildResourceWithContext(context.Background(), request)
}

// UpdateDSPASelfBuildResource
// 更新自建资源基础信息，包括：端口、账户名、密码。
//
// 请注意：
//
// 如果资源自身的VPC、VIP信息发生变化，后台会自动更新。
//
// 如果监听的端口发生变化，请显式输入端口。
//
// 如果账户名密码任意一个发生变化，请务必同时显式输入账户名密码。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateDSPASelfBuildResourceWithContext(ctx context.Context, request *UpdateDSPASelfBuildResourceRequest) (response *UpdateDSPASelfBuildResourceResponse, err error) {
    if request == nil {
        request = NewUpdateDSPASelfBuildResourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "UpdateDSPASelfBuildResource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateDSPASelfBuildResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateDSPASelfBuildResourceResponse()
    err = c.Send(request, response)
    return
}

func NewVerifyDSPACOSRuleRequest() (request *VerifyDSPACOSRuleRequest) {
    request = &VerifyDSPACOSRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "VerifyDSPACOSRule")
    
    
    return
}

func NewVerifyDSPACOSRuleResponse() (response *VerifyDSPACOSRuleResponse) {
    response = &VerifyDSPACOSRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// VerifyDSPACOSRule
// 验证COS分类分级规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) VerifyDSPACOSRule(request *VerifyDSPACOSRuleRequest) (response *VerifyDSPACOSRuleResponse, err error) {
    return c.VerifyDSPACOSRuleWithContext(context.Background(), request)
}

// VerifyDSPACOSRule
// 验证COS分类分级规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) VerifyDSPACOSRuleWithContext(ctx context.Context, request *VerifyDSPACOSRuleRequest) (response *VerifyDSPACOSRuleResponse, err error) {
    if request == nil {
        request = NewVerifyDSPACOSRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "VerifyDSPACOSRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("VerifyDSPACOSRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewVerifyDSPACOSRuleResponse()
    err = c.Send(request, response)
    return
}

func NewVerifyDSPADiscoveryRuleRequest() (request *VerifyDSPADiscoveryRuleRequest) {
    request = &VerifyDSPADiscoveryRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "VerifyDSPADiscoveryRule")
    
    
    return
}

func NewVerifyDSPADiscoveryRuleResponse() (response *VerifyDSPADiscoveryRuleResponse) {
    response = &VerifyDSPADiscoveryRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// VerifyDSPADiscoveryRule
// 验证分类分级规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) VerifyDSPADiscoveryRule(request *VerifyDSPADiscoveryRuleRequest) (response *VerifyDSPADiscoveryRuleResponse, err error) {
    return c.VerifyDSPADiscoveryRuleWithContext(context.Background(), request)
}

// VerifyDSPADiscoveryRule
// 验证分类分级规则
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) VerifyDSPADiscoveryRuleWithContext(ctx context.Context, request *VerifyDSPADiscoveryRuleRequest) (response *VerifyDSPADiscoveryRuleResponse, err error) {
    if request == nil {
        request = NewVerifyDSPADiscoveryRuleRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dsgc", APIVersion, "VerifyDSPADiscoveryRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("VerifyDSPADiscoveryRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewVerifyDSPADiscoveryRuleResponse()
    err = c.Send(request, response)
    return
}
