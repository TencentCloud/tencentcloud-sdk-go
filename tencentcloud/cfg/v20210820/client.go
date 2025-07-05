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

package v20210820

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-08-20"

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


func NewCreateTaskFromActionRequest() (request *CreateTaskFromActionRequest) {
    request = &CreateTaskFromActionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfg", APIVersion, "CreateTaskFromAction")
    
    
    return
}

func NewCreateTaskFromActionResponse() (response *CreateTaskFromActionResponse) {
    response = &CreateTaskFromActionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTaskFromAction
// 从动作创建演练
func (c *Client) CreateTaskFromAction(request *CreateTaskFromActionRequest) (response *CreateTaskFromActionResponse, err error) {
    return c.CreateTaskFromActionWithContext(context.Background(), request)
}

// CreateTaskFromAction
// 从动作创建演练
func (c *Client) CreateTaskFromActionWithContext(ctx context.Context, request *CreateTaskFromActionRequest) (response *CreateTaskFromActionResponse, err error) {
    if request == nil {
        request = NewCreateTaskFromActionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTaskFromAction require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTaskFromActionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTaskFromMultiActionRequest() (request *CreateTaskFromMultiActionRequest) {
    request = &CreateTaskFromMultiActionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfg", APIVersion, "CreateTaskFromMultiAction")
    
    
    return
}

func NewCreateTaskFromMultiActionResponse() (response *CreateTaskFromMultiActionResponse) {
    response = &CreateTaskFromMultiActionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTaskFromMultiAction
// 以多个动作创建演练
func (c *Client) CreateTaskFromMultiAction(request *CreateTaskFromMultiActionRequest) (response *CreateTaskFromMultiActionResponse, err error) {
    return c.CreateTaskFromMultiActionWithContext(context.Background(), request)
}

// CreateTaskFromMultiAction
// 以多个动作创建演练
func (c *Client) CreateTaskFromMultiActionWithContext(ctx context.Context, request *CreateTaskFromMultiActionRequest) (response *CreateTaskFromMultiActionResponse, err error) {
    if request == nil {
        request = NewCreateTaskFromMultiActionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTaskFromMultiAction require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTaskFromMultiActionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTaskFromTemplateRequest() (request *CreateTaskFromTemplateRequest) {
    request = &CreateTaskFromTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfg", APIVersion, "CreateTaskFromTemplate")
    
    
    return
}

func NewCreateTaskFromTemplateResponse() (response *CreateTaskFromTemplateResponse) {
    response = &CreateTaskFromTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTaskFromTemplate
// 从经验库创建演练
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateTaskFromTemplate(request *CreateTaskFromTemplateRequest) (response *CreateTaskFromTemplateResponse, err error) {
    return c.CreateTaskFromTemplateWithContext(context.Background(), request)
}

// CreateTaskFromTemplate
// 从经验库创建演练
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateTaskFromTemplateWithContext(ctx context.Context, request *CreateTaskFromTemplateRequest) (response *CreateTaskFromTemplateResponse, err error) {
    if request == nil {
        request = NewCreateTaskFromTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTaskFromTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTaskFromTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTaskRequest() (request *DeleteTaskRequest) {
    request = &DeleteTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfg", APIVersion, "DeleteTask")
    
    
    return
}

func NewDeleteTaskResponse() (response *DeleteTaskResponse) {
    response = &DeleteTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTask
// 删除任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteTask(request *DeleteTaskRequest) (response *DeleteTaskResponse, err error) {
    return c.DeleteTaskWithContext(context.Background(), request)
}

// DeleteTask
// 删除任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DeleteTaskWithContext(ctx context.Context, request *DeleteTaskRequest) (response *DeleteTaskResponse, err error) {
    if request == nil {
        request = NewDeleteTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeActionFieldConfigListRequest() (request *DescribeActionFieldConfigListRequest) {
    request = &DescribeActionFieldConfigListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfg", APIVersion, "DescribeActionFieldConfigList")
    
    
    return
}

func NewDescribeActionFieldConfigListResponse() (response *DescribeActionFieldConfigListResponse) {
    response = &DescribeActionFieldConfigListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeActionFieldConfigList
// 根据动作ID获取动作栏位动态配置参数信息，里面包含动作自有和通用两部分参数。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeActionFieldConfigList(request *DescribeActionFieldConfigListRequest) (response *DescribeActionFieldConfigListResponse, err error) {
    return c.DescribeActionFieldConfigListWithContext(context.Background(), request)
}

// DescribeActionFieldConfigList
// 根据动作ID获取动作栏位动态配置参数信息，里面包含动作自有和通用两部分参数。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeActionFieldConfigListWithContext(ctx context.Context, request *DescribeActionFieldConfigListRequest) (response *DescribeActionFieldConfigListResponse, err error) {
    if request == nil {
        request = NewDescribeActionFieldConfigListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeActionFieldConfigList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeActionFieldConfigListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeActionLibraryListRequest() (request *DescribeActionLibraryListRequest) {
    request = &DescribeActionLibraryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfg", APIVersion, "DescribeActionLibraryList")
    
    
    return
}

func NewDescribeActionLibraryListResponse() (response *DescribeActionLibraryListResponse) {
    response = &DescribeActionLibraryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeActionLibraryList
// 获取混沌演练平台的动作库列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeActionLibraryList(request *DescribeActionLibraryListRequest) (response *DescribeActionLibraryListResponse, err error) {
    return c.DescribeActionLibraryListWithContext(context.Background(), request)
}

// DescribeActionLibraryList
// 获取混沌演练平台的动作库列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeActionLibraryListWithContext(ctx context.Context, request *DescribeActionLibraryListRequest) (response *DescribeActionLibraryListResponse, err error) {
    if request == nil {
        request = NewDescribeActionLibraryListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeActionLibraryList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeActionLibraryListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeObjectTypeListRequest() (request *DescribeObjectTypeListRequest) {
    request = &DescribeObjectTypeListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfg", APIVersion, "DescribeObjectTypeList")
    
    
    return
}

func NewDescribeObjectTypeListResponse() (response *DescribeObjectTypeListResponse) {
    response = &DescribeObjectTypeListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeObjectTypeList
// 查询对象类型列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeObjectTypeList(request *DescribeObjectTypeListRequest) (response *DescribeObjectTypeListResponse, err error) {
    return c.DescribeObjectTypeListWithContext(context.Background(), request)
}

// DescribeObjectTypeList
// 查询对象类型列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeObjectTypeListWithContext(ctx context.Context, request *DescribeObjectTypeListRequest) (response *DescribeObjectTypeListResponse, err error) {
    if request == nil {
        request = NewDescribeObjectTypeListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeObjectTypeList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeObjectTypeListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskRequest() (request *DescribeTaskRequest) {
    request = &DescribeTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfg", APIVersion, "DescribeTask")
    
    
    return
}

func NewDescribeTaskResponse() (response *DescribeTaskResponse) {
    response = &DescribeTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTask
// 查询任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTask(request *DescribeTaskRequest) (response *DescribeTaskResponse, err error) {
    return c.DescribeTaskWithContext(context.Background(), request)
}

// DescribeTask
// 查询任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTaskWithContext(ctx context.Context, request *DescribeTaskRequest) (response *DescribeTaskResponse, err error) {
    if request == nil {
        request = NewDescribeTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskExecuteLogsRequest() (request *DescribeTaskExecuteLogsRequest) {
    request = &DescribeTaskExecuteLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfg", APIVersion, "DescribeTaskExecuteLogs")
    
    
    return
}

func NewDescribeTaskExecuteLogsResponse() (response *DescribeTaskExecuteLogsResponse) {
    response = &DescribeTaskExecuteLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTaskExecuteLogs
// 获取演练过程中的所有日志
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeTaskExecuteLogs(request *DescribeTaskExecuteLogsRequest) (response *DescribeTaskExecuteLogsResponse, err error) {
    return c.DescribeTaskExecuteLogsWithContext(context.Background(), request)
}

// DescribeTaskExecuteLogs
// 获取演练过程中的所有日志
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeTaskExecuteLogsWithContext(ctx context.Context, request *DescribeTaskExecuteLogsRequest) (response *DescribeTaskExecuteLogsResponse, err error) {
    if request == nil {
        request = NewDescribeTaskExecuteLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskExecuteLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskExecuteLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskListRequest() (request *DescribeTaskListRequest) {
    request = &DescribeTaskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfg", APIVersion, "DescribeTaskList")
    
    
    return
}

func NewDescribeTaskListResponse() (response *DescribeTaskListResponse) {
    response = &DescribeTaskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTaskList
// 查询任务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTaskList(request *DescribeTaskListRequest) (response *DescribeTaskListResponse, err error) {
    return c.DescribeTaskListWithContext(context.Background(), request)
}

// DescribeTaskList
// 查询任务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTaskListWithContext(ctx context.Context, request *DescribeTaskListRequest) (response *DescribeTaskListResponse, err error) {
    if request == nil {
        request = NewDescribeTaskListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskPolicyTriggerLogRequest() (request *DescribeTaskPolicyTriggerLogRequest) {
    request = &DescribeTaskPolicyTriggerLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfg", APIVersion, "DescribeTaskPolicyTriggerLog")
    
    
    return
}

func NewDescribeTaskPolicyTriggerLogResponse() (response *DescribeTaskPolicyTriggerLogResponse) {
    response = &DescribeTaskPolicyTriggerLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTaskPolicyTriggerLog
// 获取护栏触发日志
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTaskPolicyTriggerLog(request *DescribeTaskPolicyTriggerLogRequest) (response *DescribeTaskPolicyTriggerLogResponse, err error) {
    return c.DescribeTaskPolicyTriggerLogWithContext(context.Background(), request)
}

// DescribeTaskPolicyTriggerLog
// 获取护栏触发日志
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTaskPolicyTriggerLogWithContext(ctx context.Context, request *DescribeTaskPolicyTriggerLogRequest) (response *DescribeTaskPolicyTriggerLogResponse, err error) {
    if request == nil {
        request = NewDescribeTaskPolicyTriggerLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskPolicyTriggerLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskPolicyTriggerLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTemplateRequest() (request *DescribeTemplateRequest) {
    request = &DescribeTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfg", APIVersion, "DescribeTemplate")
    
    
    return
}

func NewDescribeTemplateResponse() (response *DescribeTemplateResponse) {
    response = &DescribeTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTemplate
// 查询经验库
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTemplate(request *DescribeTemplateRequest) (response *DescribeTemplateResponse, err error) {
    return c.DescribeTemplateWithContext(context.Background(), request)
}

// DescribeTemplate
// 查询经验库
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTemplateWithContext(ctx context.Context, request *DescribeTemplateRequest) (response *DescribeTemplateResponse, err error) {
    if request == nil {
        request = NewDescribeTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTemplateListRequest() (request *DescribeTemplateListRequest) {
    request = &DescribeTemplateListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfg", APIVersion, "DescribeTemplateList")
    
    
    return
}

func NewDescribeTemplateListResponse() (response *DescribeTemplateListResponse) {
    response = &DescribeTemplateListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTemplateList
// 查询经验库列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTemplateList(request *DescribeTemplateListRequest) (response *DescribeTemplateListResponse, err error) {
    return c.DescribeTemplateListWithContext(context.Background(), request)
}

// DescribeTemplateList
// 查询经验库列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTemplateListWithContext(ctx context.Context, request *DescribeTemplateListRequest) (response *DescribeTemplateListResponse, err error) {
    if request == nil {
        request = NewDescribeTemplateListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTemplateList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTemplateListResponse()
    err = c.Send(request, response)
    return
}

func NewExecuteTaskRequest() (request *ExecuteTaskRequest) {
    request = &ExecuteTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfg", APIVersion, "ExecuteTask")
    
    
    return
}

func NewExecuteTaskResponse() (response *ExecuteTaskResponse) {
    response = &ExecuteTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ExecuteTask
// 执行任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ExecuteTask(request *ExecuteTaskRequest) (response *ExecuteTaskResponse, err error) {
    return c.ExecuteTaskWithContext(context.Background(), request)
}

// ExecuteTask
// 执行任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ExecuteTaskWithContext(ctx context.Context, request *ExecuteTaskRequest) (response *ExecuteTaskResponse, err error) {
    if request == nil {
        request = NewExecuteTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExecuteTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewExecuteTaskResponse()
    err = c.Send(request, response)
    return
}

func NewExecuteTaskInstanceRequest() (request *ExecuteTaskInstanceRequest) {
    request = &ExecuteTaskInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfg", APIVersion, "ExecuteTaskInstance")
    
    
    return
}

func NewExecuteTaskInstanceResponse() (response *ExecuteTaskInstanceResponse) {
    response = &ExecuteTaskInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ExecuteTaskInstance
// 触发混沌演练任务的动作，对于实例进行演练操作
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ExecuteTaskInstance(request *ExecuteTaskInstanceRequest) (response *ExecuteTaskInstanceResponse, err error) {
    return c.ExecuteTaskInstanceWithContext(context.Background(), request)
}

// ExecuteTaskInstance
// 触发混沌演练任务的动作，对于实例进行演练操作
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ExecuteTaskInstanceWithContext(ctx context.Context, request *ExecuteTaskInstanceRequest) (response *ExecuteTaskInstanceResponse, err error) {
    if request == nil {
        request = NewExecuteTaskInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExecuteTaskInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewExecuteTaskInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTaskRunStatusRequest() (request *ModifyTaskRunStatusRequest) {
    request = &ModifyTaskRunStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfg", APIVersion, "ModifyTaskRunStatus")
    
    
    return
}

func NewModifyTaskRunStatusResponse() (response *ModifyTaskRunStatusResponse) {
    response = &ModifyTaskRunStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTaskRunStatus
// 修改任务运行状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyTaskRunStatus(request *ModifyTaskRunStatusRequest) (response *ModifyTaskRunStatusResponse, err error) {
    return c.ModifyTaskRunStatusWithContext(context.Background(), request)
}

// ModifyTaskRunStatus
// 修改任务运行状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyTaskRunStatusWithContext(ctx context.Context, request *ModifyTaskRunStatusRequest) (response *ModifyTaskRunStatusResponse, err error) {
    if request == nil {
        request = NewModifyTaskRunStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTaskRunStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTaskRunStatusResponse()
    err = c.Send(request, response)
    return
}

func NewTriggerPolicyRequest() (request *TriggerPolicyRequest) {
    request = &TriggerPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cfg", APIVersion, "TriggerPolicy")
    
    
    return
}

func NewTriggerPolicyResponse() (response *TriggerPolicyResponse) {
    response = &TriggerPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TriggerPolicy
// 用于触发混沌演练护栏（类型为触发和恢复2种）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) TriggerPolicy(request *TriggerPolicyRequest) (response *TriggerPolicyResponse, err error) {
    return c.TriggerPolicyWithContext(context.Background(), request)
}

// TriggerPolicy
// 用于触发混沌演练护栏（类型为触发和恢复2种）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) TriggerPolicyWithContext(ctx context.Context, request *TriggerPolicyRequest) (response *TriggerPolicyResponse, err error) {
    if request == nil {
        request = NewTriggerPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TriggerPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewTriggerPolicyResponse()
    err = c.Send(request, response)
    return
}
