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

package v20221128

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2022-11-28"

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


func NewCreateEnvironmentRequest() (request *CreateEnvironmentRequest) {
    request = &CreateEnvironmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("omics", APIVersion, "CreateEnvironment")
    
    
    return
}

func NewCreateEnvironmentResponse() (response *CreateEnvironmentResponse) {
    response = &CreateEnvironmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateEnvironment
// 创建组学平台计算环境。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) CreateEnvironment(request *CreateEnvironmentRequest) (response *CreateEnvironmentResponse, err error) {
    return c.CreateEnvironmentWithContext(context.Background(), request)
}

// CreateEnvironment
// 创建组学平台计算环境。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) CreateEnvironmentWithContext(ctx context.Context, request *CreateEnvironmentRequest) (response *CreateEnvironmentResponse, err error) {
    if request == nil {
        request = NewCreateEnvironmentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEnvironment require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEnvironmentResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEnvironmentRequest() (request *DeleteEnvironmentRequest) {
    request = &DeleteEnvironmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("omics", APIVersion, "DeleteEnvironment")
    
    
    return
}

func NewDeleteEnvironmentResponse() (response *DeleteEnvironmentResponse) {
    response = &DeleteEnvironmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteEnvironment
// 删除环境。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteEnvironment(request *DeleteEnvironmentRequest) (response *DeleteEnvironmentResponse, err error) {
    return c.DeleteEnvironmentWithContext(context.Background(), request)
}

// DeleteEnvironment
// 删除环境。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteEnvironmentWithContext(ctx context.Context, request *DeleteEnvironmentRequest) (response *DeleteEnvironmentResponse, err error) {
    if request == nil {
        request = NewDeleteEnvironmentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteEnvironment require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteEnvironmentResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnvironmentsRequest() (request *DescribeEnvironmentsRequest) {
    request = &DescribeEnvironmentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("omics", APIVersion, "DescribeEnvironments")
    
    
    return
}

func NewDescribeEnvironmentsResponse() (response *DescribeEnvironmentsResponse) {
    response = &DescribeEnvironmentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEnvironments
// 查询环境列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeEnvironments(request *DescribeEnvironmentsRequest) (response *DescribeEnvironmentsResponse, err error) {
    return c.DescribeEnvironmentsWithContext(context.Background(), request)
}

// DescribeEnvironments
// 查询环境列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeEnvironmentsWithContext(ctx context.Context, request *DescribeEnvironmentsRequest) (response *DescribeEnvironmentsResponse, err error) {
    if request == nil {
        request = NewDescribeEnvironmentsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEnvironments require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEnvironmentsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRunGroupsRequest() (request *DescribeRunGroupsRequest) {
    request = &DescribeRunGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("omics", APIVersion, "DescribeRunGroups")
    
    
    return
}

func NewDescribeRunGroupsResponse() (response *DescribeRunGroupsResponse) {
    response = &DescribeRunGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRunGroups
// 查询任务批次列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
func (c *Client) DescribeRunGroups(request *DescribeRunGroupsRequest) (response *DescribeRunGroupsResponse, err error) {
    return c.DescribeRunGroupsWithContext(context.Background(), request)
}

// DescribeRunGroups
// 查询任务批次列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
func (c *Client) DescribeRunGroupsWithContext(ctx context.Context, request *DescribeRunGroupsRequest) (response *DescribeRunGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeRunGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRunGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRunGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRunsRequest() (request *DescribeRunsRequest) {
    request = &DescribeRunsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("omics", APIVersion, "DescribeRuns")
    
    
    return
}

func NewDescribeRunsResponse() (response *DescribeRunsResponse) {
    response = &DescribeRunsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRuns
// 查询任务列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
func (c *Client) DescribeRuns(request *DescribeRunsRequest) (response *DescribeRunsResponse, err error) {
    return c.DescribeRunsWithContext(context.Background(), request)
}

// DescribeRuns
// 查询任务列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
func (c *Client) DescribeRunsWithContext(ctx context.Context, request *DescribeRunsRequest) (response *DescribeRunsResponse, err error) {
    if request == nil {
        request = NewDescribeRunsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRuns require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRunsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTablesRequest() (request *DescribeTablesRequest) {
    request = &DescribeTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("omics", APIVersion, "DescribeTables")
    
    
    return
}

func NewDescribeTablesResponse() (response *DescribeTablesResponse) {
    response = &DescribeTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTables
// 查询表格。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
func (c *Client) DescribeTables(request *DescribeTablesRequest) (response *DescribeTablesResponse, err error) {
    return c.DescribeTablesWithContext(context.Background(), request)
}

// DescribeTables
// 查询表格。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
func (c *Client) DescribeTablesWithContext(ctx context.Context, request *DescribeTablesRequest) (response *DescribeTablesResponse, err error) {
    if request == nil {
        request = NewDescribeTablesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTables require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTablesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTablesRowsRequest() (request *DescribeTablesRowsRequest) {
    request = &DescribeTablesRowsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("omics", APIVersion, "DescribeTablesRows")
    
    
    return
}

func NewDescribeTablesRowsResponse() (response *DescribeTablesRowsResponse) {
    response = &DescribeTablesRowsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTablesRows
// 查询表格行数据。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  RESOURCENOTFOUND_TABLENOTEXIST = "ResourceNotFound.TableNotExist"
func (c *Client) DescribeTablesRows(request *DescribeTablesRowsRequest) (response *DescribeTablesRowsResponse, err error) {
    return c.DescribeTablesRowsWithContext(context.Background(), request)
}

// DescribeTablesRows
// 查询表格行数据。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  RESOURCENOTFOUND_TABLENOTEXIST = "ResourceNotFound.TableNotExist"
func (c *Client) DescribeTablesRowsWithContext(ctx context.Context, request *DescribeTablesRowsRequest) (response *DescribeTablesRowsResponse, err error) {
    if request == nil {
        request = NewDescribeTablesRowsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTablesRows require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTablesRowsResponse()
    err = c.Send(request, response)
    return
}

func NewGetRunCallsRequest() (request *GetRunCallsRequest) {
    request = &GetRunCallsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("omics", APIVersion, "GetRunCalls")
    
    
    return
}

func NewGetRunCallsResponse() (response *GetRunCallsResponse) {
    response = &GetRunCallsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetRunCalls
// 查询作业详情。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_ENVIRONMENTNOTEXIST = "ResourceNotFound.EnvironmentNotExist"
//  RESOURCENOTFOUND_RUNNOTEXIST = "ResourceNotFound.RunNotExist"
func (c *Client) GetRunCalls(request *GetRunCallsRequest) (response *GetRunCallsResponse, err error) {
    return c.GetRunCallsWithContext(context.Background(), request)
}

// GetRunCalls
// 查询作业详情。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_ENVIRONMENTNOTEXIST = "ResourceNotFound.EnvironmentNotExist"
//  RESOURCENOTFOUND_RUNNOTEXIST = "ResourceNotFound.RunNotExist"
func (c *Client) GetRunCallsWithContext(ctx context.Context, request *GetRunCallsRequest) (response *GetRunCallsResponse, err error) {
    if request == nil {
        request = NewGetRunCallsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetRunCalls require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetRunCallsResponse()
    err = c.Send(request, response)
    return
}

func NewGetRunMetadataFileRequest() (request *GetRunMetadataFileRequest) {
    request = &GetRunMetadataFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("omics", APIVersion, "GetRunMetadataFile")
    
    
    return
}

func NewGetRunMetadataFileResponse() (response *GetRunMetadataFileResponse) {
    response = &GetRunMetadataFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetRunMetadataFile
// 查询任务详情文件。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_COSOBJECTNOTEXIST = "ResourceNotFound.CosObjectNotExist"
//  RESOURCENOTFOUND_ENVIRONMENTNOTEXIST = "ResourceNotFound.EnvironmentNotExist"
//  RESOURCENOTFOUND_RUNNOTEXIST = "ResourceNotFound.RunNotExist"
func (c *Client) GetRunMetadataFile(request *GetRunMetadataFileRequest) (response *GetRunMetadataFileResponse, err error) {
    return c.GetRunMetadataFileWithContext(context.Background(), request)
}

// GetRunMetadataFile
// 查询任务详情文件。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_COSOBJECTNOTEXIST = "ResourceNotFound.CosObjectNotExist"
//  RESOURCENOTFOUND_ENVIRONMENTNOTEXIST = "ResourceNotFound.EnvironmentNotExist"
//  RESOURCENOTFOUND_RUNNOTEXIST = "ResourceNotFound.RunNotExist"
func (c *Client) GetRunMetadataFileWithContext(ctx context.Context, request *GetRunMetadataFileRequest) (response *GetRunMetadataFileResponse, err error) {
    if request == nil {
        request = NewGetRunMetadataFileRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetRunMetadataFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetRunMetadataFileResponse()
    err = c.Send(request, response)
    return
}

func NewGetRunStatusRequest() (request *GetRunStatusRequest) {
    request = &GetRunStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("omics", APIVersion, "GetRunStatus")
    
    
    return
}

func NewGetRunStatusResponse() (response *GetRunStatusResponse) {
    response = &GetRunStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetRunStatus
// 查询任务详情。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_ENVIRONMENTNOTEXIST = "ResourceNotFound.EnvironmentNotExist"
//  RESOURCENOTFOUND_RUNNOTEXIST = "ResourceNotFound.RunNotExist"
func (c *Client) GetRunStatus(request *GetRunStatusRequest) (response *GetRunStatusResponse, err error) {
    return c.GetRunStatusWithContext(context.Background(), request)
}

// GetRunStatus
// 查询任务详情。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_ENVIRONMENTNOTEXIST = "ResourceNotFound.EnvironmentNotExist"
//  RESOURCENOTFOUND_RUNNOTEXIST = "ResourceNotFound.RunNotExist"
func (c *Client) GetRunStatusWithContext(ctx context.Context, request *GetRunStatusRequest) (response *GetRunStatusResponse, err error) {
    if request == nil {
        request = NewGetRunStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetRunStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetRunStatusResponse()
    err = c.Send(request, response)
    return
}

func NewImportTableFileRequest() (request *ImportTableFileRequest) {
    request = &ImportTableFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("omics", APIVersion, "ImportTableFile")
    
    
    return
}

func NewImportTableFileResponse() (response *ImportTableFileResponse) {
    response = &ImportTableFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ImportTableFile
// 导入表格文件。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DUPLICATETABLEHEADER = "FailedOperation.DuplicateTableHeader"
//  FAILEDOPERATION_EMPTYTABLEHEADER = "FailedOperation.EmptyTableHeader"
//  FAILEDOPERATION_INVALIDTABLEHEADER = "FailedOperation.InvalidTableHeader"
//  FAILEDOPERATION_INVALIDTABLELENGTH = "FailedOperation.InvalidTableLength"
//  FAILEDOPERATION_TABLEDATATYPEMISMATCH = "FailedOperation.TableDataTypeMismatch"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  INVALIDPARAMETERVALUE_INVALIDCOSKEY = "InvalidParameterValue.InvalidCosKey"
//  INVALIDPARAMETERVALUE_INVALIDCSVFORMAT = "InvalidParameterValue.InvalidCsvFormat"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"
//  INVALIDPARAMETERVALUE_INVALIDNAME = "InvalidParameterValue.InvalidName"
//  INVALIDPARAMETERVALUE_TABLEDATATYPELENGTHMISMATCH = "InvalidParameterValue.TableDataTypeLengthMismatch"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTABLEDATATYPE = "InvalidParameterValue.UnsupportedTableDataType"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_COSBUCKETNOTEXIST = "ResourceNotFound.CosBucketNotExist"
//  RESOURCENOTFOUND_COSOBJECTNOTEXIST = "ResourceNotFound.CosObjectNotExist"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  RESOURCENOTFOUND_TABLENOTEXIST = "ResourceNotFound.TableNotExist"
func (c *Client) ImportTableFile(request *ImportTableFileRequest) (response *ImportTableFileResponse, err error) {
    return c.ImportTableFileWithContext(context.Background(), request)
}

// ImportTableFile
// 导入表格文件。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DUPLICATETABLEHEADER = "FailedOperation.DuplicateTableHeader"
//  FAILEDOPERATION_EMPTYTABLEHEADER = "FailedOperation.EmptyTableHeader"
//  FAILEDOPERATION_INVALIDTABLEHEADER = "FailedOperation.InvalidTableHeader"
//  FAILEDOPERATION_INVALIDTABLELENGTH = "FailedOperation.InvalidTableLength"
//  FAILEDOPERATION_TABLEDATATYPEMISMATCH = "FailedOperation.TableDataTypeMismatch"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  INVALIDPARAMETERVALUE_INVALIDCOSKEY = "InvalidParameterValue.InvalidCosKey"
//  INVALIDPARAMETERVALUE_INVALIDCSVFORMAT = "InvalidParameterValue.InvalidCsvFormat"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"
//  INVALIDPARAMETERVALUE_INVALIDNAME = "InvalidParameterValue.InvalidName"
//  INVALIDPARAMETERVALUE_TABLEDATATYPELENGTHMISMATCH = "InvalidParameterValue.TableDataTypeLengthMismatch"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTABLEDATATYPE = "InvalidParameterValue.UnsupportedTableDataType"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_COSBUCKETNOTEXIST = "ResourceNotFound.CosBucketNotExist"
//  RESOURCENOTFOUND_COSOBJECTNOTEXIST = "ResourceNotFound.CosObjectNotExist"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  RESOURCENOTFOUND_TABLENOTEXIST = "ResourceNotFound.TableNotExist"
func (c *Client) ImportTableFileWithContext(ctx context.Context, request *ImportTableFileRequest) (response *ImportTableFileResponse, err error) {
    if request == nil {
        request = NewImportTableFileRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImportTableFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewImportTableFileResponse()
    err = c.Send(request, response)
    return
}

func NewRetryRunsRequest() (request *RetryRunsRequest) {
    request = &RetryRunsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("omics", APIVersion, "RetryRuns")
    
    
    return
}

func NewRetryRunsResponse() (response *RetryRunsResponse) {
    response = &RetryRunsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RetryRuns
// 重试任务。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_RETRYLIMITEXCEEDED = "FailedOperation.RetryLimitExceeded"
//  FAILEDOPERATION_STATUSNOTSUPPORTED = "FailedOperation.StatusNotSupported"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ENVIRONMENTNOTAVAILABLE = "InvalidParameterValue.EnvironmentNotAvailable"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_ENVIRONMENTNOTEXIST = "ResourceNotFound.EnvironmentNotExist"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  RESOURCENOTFOUND_RUNGROUPNOTEXIST = "ResourceNotFound.RunGroupNotExist"
//  RESOURCENOTFOUND_RUNNOTEXIST = "ResourceNotFound.RunNotExist"
func (c *Client) RetryRuns(request *RetryRunsRequest) (response *RetryRunsResponse, err error) {
    return c.RetryRunsWithContext(context.Background(), request)
}

// RetryRuns
// 重试任务。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_RETRYLIMITEXCEEDED = "FailedOperation.RetryLimitExceeded"
//  FAILEDOPERATION_STATUSNOTSUPPORTED = "FailedOperation.StatusNotSupported"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ENVIRONMENTNOTAVAILABLE = "InvalidParameterValue.EnvironmentNotAvailable"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_ENVIRONMENTNOTEXIST = "ResourceNotFound.EnvironmentNotExist"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  RESOURCENOTFOUND_RUNGROUPNOTEXIST = "ResourceNotFound.RunGroupNotExist"
//  RESOURCENOTFOUND_RUNNOTEXIST = "ResourceNotFound.RunNotExist"
func (c *Client) RetryRunsWithContext(ctx context.Context, request *RetryRunsRequest) (response *RetryRunsResponse, err error) {
    if request == nil {
        request = NewRetryRunsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RetryRuns require credential")
    }

    request.SetContext(ctx)
    
    response = NewRetryRunsResponse()
    err = c.Send(request, response)
    return
}

func NewRunApplicationRequest() (request *RunApplicationRequest) {
    request = &RunApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("omics", APIVersion, "RunApplication")
    
    
    return
}

func NewRunApplicationResponse() (response *RunApplicationResponse) {
    response = &RunApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RunApplication
// 运行应用。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_VERSIONNOTRELEASED = "FailedOperation.VersionNotReleased"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  INVALIDPARAMETERVALUE_ENTRYPOINTNOTSET = "InvalidParameterValue.EntrypointNotSet"
//  INVALIDPARAMETERVALUE_ENVIRONMENTNOTAVAILABLE = "InvalidParameterValue.EnvironmentNotAvailable"
//  INVALIDPARAMETERVALUE_INVALIDBASE64ENCODE = "InvalidParameterValue.InvalidBase64Encode"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"
//  INVALIDPARAMETERVALUE_INVALIDINPUTJSONFORMAT = "InvalidParameterValue.InvalidInputJsonFormat"
//  INVALIDPARAMETERVALUE_INVALIDINPUTPLACEHOLDER = "InvalidParameterValue.InvalidInputPlaceholder"
//  INVALIDPARAMETERVALUE_INVALIDNAME = "InvalidParameterValue.InvalidName"
//  INVALIDPARAMETERVALUE_INVALIDRUNOPTION = "InvalidParameterValue.InvalidRunOption"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATIONNOTEXIST = "ResourceNotFound.ApplicationNotExist"
//  RESOURCENOTFOUND_APPLICATIONVERSIONNOTEXIST = "ResourceNotFound.ApplicationVersionNotExist"
//  RESOURCENOTFOUND_ENVIRONMENTNOTEXIST = "ResourceNotFound.EnvironmentNotExist"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  RESOURCENOTFOUND_TABLENOTEXIST = "ResourceNotFound.TableNotExist"
//  RESOURCENOTFOUND_TABLEROWNOTEXIST = "ResourceNotFound.TableRowNotExist"
func (c *Client) RunApplication(request *RunApplicationRequest) (response *RunApplicationResponse, err error) {
    return c.RunApplicationWithContext(context.Background(), request)
}

// RunApplication
// 运行应用。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_VERSIONNOTRELEASED = "FailedOperation.VersionNotReleased"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  INVALIDPARAMETERVALUE_ENTRYPOINTNOTSET = "InvalidParameterValue.EntrypointNotSet"
//  INVALIDPARAMETERVALUE_ENVIRONMENTNOTAVAILABLE = "InvalidParameterValue.EnvironmentNotAvailable"
//  INVALIDPARAMETERVALUE_INVALIDBASE64ENCODE = "InvalidParameterValue.InvalidBase64Encode"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"
//  INVALIDPARAMETERVALUE_INVALIDINPUTJSONFORMAT = "InvalidParameterValue.InvalidInputJsonFormat"
//  INVALIDPARAMETERVALUE_INVALIDINPUTPLACEHOLDER = "InvalidParameterValue.InvalidInputPlaceholder"
//  INVALIDPARAMETERVALUE_INVALIDNAME = "InvalidParameterValue.InvalidName"
//  INVALIDPARAMETERVALUE_INVALIDRUNOPTION = "InvalidParameterValue.InvalidRunOption"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATIONNOTEXIST = "ResourceNotFound.ApplicationNotExist"
//  RESOURCENOTFOUND_APPLICATIONVERSIONNOTEXIST = "ResourceNotFound.ApplicationVersionNotExist"
//  RESOURCENOTFOUND_ENVIRONMENTNOTEXIST = "ResourceNotFound.EnvironmentNotExist"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  RESOURCENOTFOUND_TABLENOTEXIST = "ResourceNotFound.TableNotExist"
//  RESOURCENOTFOUND_TABLEROWNOTEXIST = "ResourceNotFound.TableRowNotExist"
func (c *Client) RunApplicationWithContext(ctx context.Context, request *RunApplicationRequest) (response *RunApplicationResponse, err error) {
    if request == nil {
        request = NewRunApplicationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewRunApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewRunWorkflowRequest() (request *RunWorkflowRequest) {
    request = &RunWorkflowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("omics", APIVersion, "RunWorkflow")
    
    
    return
}

func NewRunWorkflowResponse() (response *RunWorkflowResponse) {
    response = &RunWorkflowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RunWorkflow
// 运行工作流。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_VERSIONNOTRELEASED = "FailedOperation.VersionNotReleased"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  INVALIDPARAMETERVALUE_ENTRYPOINTNOTSET = "InvalidParameterValue.EntrypointNotSet"
//  INVALIDPARAMETERVALUE_ENVIRONMENTNOTAVAILABLE = "InvalidParameterValue.EnvironmentNotAvailable"
//  INVALIDPARAMETERVALUE_INVALIDBASE64ENCODE = "InvalidParameterValue.InvalidBase64Encode"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"
//  INVALIDPARAMETERVALUE_INVALIDINPUTJSONFORMAT = "InvalidParameterValue.InvalidInputJsonFormat"
//  INVALIDPARAMETERVALUE_INVALIDINPUTPLACEHOLDER = "InvalidParameterValue.InvalidInputPlaceholder"
//  INVALIDPARAMETERVALUE_INVALIDNAME = "InvalidParameterValue.InvalidName"
//  INVALIDPARAMETERVALUE_INVALIDRUNOPTION = "InvalidParameterValue.InvalidRunOption"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATIONNOTEXIST = "ResourceNotFound.ApplicationNotExist"
//  RESOURCENOTFOUND_APPLICATIONVERSIONNOTEXIST = "ResourceNotFound.ApplicationVersionNotExist"
//  RESOURCENOTFOUND_COSBUCKETNOTEXIST = "ResourceNotFound.CosBucketNotExist"
//  RESOURCENOTFOUND_ENVIRONMENTNOTEXIST = "ResourceNotFound.EnvironmentNotExist"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  RESOURCENOTFOUND_TABLENOTEXIST = "ResourceNotFound.TableNotExist"
//  RESOURCENOTFOUND_TABLEROWNOTEXIST = "ResourceNotFound.TableRowNotExist"
func (c *Client) RunWorkflow(request *RunWorkflowRequest) (response *RunWorkflowResponse, err error) {
    return c.RunWorkflowWithContext(context.Background(), request)
}

// RunWorkflow
// 运行工作流。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_VERSIONNOTRELEASED = "FailedOperation.VersionNotReleased"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  INVALIDPARAMETERVALUE_ENTRYPOINTNOTSET = "InvalidParameterValue.EntrypointNotSet"
//  INVALIDPARAMETERVALUE_ENVIRONMENTNOTAVAILABLE = "InvalidParameterValue.EnvironmentNotAvailable"
//  INVALIDPARAMETERVALUE_INVALIDBASE64ENCODE = "InvalidParameterValue.InvalidBase64Encode"
//  INVALIDPARAMETERVALUE_INVALIDDESCRIPTION = "InvalidParameterValue.InvalidDescription"
//  INVALIDPARAMETERVALUE_INVALIDINPUTJSONFORMAT = "InvalidParameterValue.InvalidInputJsonFormat"
//  INVALIDPARAMETERVALUE_INVALIDINPUTPLACEHOLDER = "InvalidParameterValue.InvalidInputPlaceholder"
//  INVALIDPARAMETERVALUE_INVALIDNAME = "InvalidParameterValue.InvalidName"
//  INVALIDPARAMETERVALUE_INVALIDRUNOPTION = "InvalidParameterValue.InvalidRunOption"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATIONNOTEXIST = "ResourceNotFound.ApplicationNotExist"
//  RESOURCENOTFOUND_APPLICATIONVERSIONNOTEXIST = "ResourceNotFound.ApplicationVersionNotExist"
//  RESOURCENOTFOUND_COSBUCKETNOTEXIST = "ResourceNotFound.CosBucketNotExist"
//  RESOURCENOTFOUND_ENVIRONMENTNOTEXIST = "ResourceNotFound.EnvironmentNotExist"
//  RESOURCENOTFOUND_PROJECTNOTEXIST = "ResourceNotFound.ProjectNotExist"
//  RESOURCENOTFOUND_TABLENOTEXIST = "ResourceNotFound.TableNotExist"
//  RESOURCENOTFOUND_TABLEROWNOTEXIST = "ResourceNotFound.TableRowNotExist"
func (c *Client) RunWorkflowWithContext(ctx context.Context, request *RunWorkflowRequest) (response *RunWorkflowResponse, err error) {
    if request == nil {
        request = NewRunWorkflowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunWorkflow require credential")
    }

    request.SetContext(ctx)
    
    response = NewRunWorkflowResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateRunGroupRequest() (request *TerminateRunGroupRequest) {
    request = &TerminateRunGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("omics", APIVersion, "TerminateRunGroup")
    
    
    return
}

func NewTerminateRunGroupResponse() (response *TerminateRunGroupResponse) {
    response = &TerminateRunGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TerminateRunGroup
// 终止任务批次。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_ENVIRONMENTNOTEXIST = "ResourceNotFound.EnvironmentNotExist"
//  RESOURCENOTFOUND_RUNNOTEXIST = "ResourceNotFound.RunNotExist"
func (c *Client) TerminateRunGroup(request *TerminateRunGroupRequest) (response *TerminateRunGroupResponse, err error) {
    return c.TerminateRunGroupWithContext(context.Background(), request)
}

// TerminateRunGroup
// 终止任务批次。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_ENVIRONMENTNOTEXIST = "ResourceNotFound.EnvironmentNotExist"
//  RESOURCENOTFOUND_RUNNOTEXIST = "ResourceNotFound.RunNotExist"
func (c *Client) TerminateRunGroupWithContext(ctx context.Context, request *TerminateRunGroupRequest) (response *TerminateRunGroupResponse, err error) {
    if request == nil {
        request = NewTerminateRunGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateRunGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminateRunGroupResponse()
    err = c.Send(request, response)
    return
}
