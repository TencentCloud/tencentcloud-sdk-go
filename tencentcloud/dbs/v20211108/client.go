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

package v20211108

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-11-08"

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


func NewConfigureBackupPlanRequest() (request *ConfigureBackupPlanRequest) {
    request = &ConfigureBackupPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbs", APIVersion, "ConfigureBackupPlan")
    
    
    return
}

func NewConfigureBackupPlanResponse() (response *ConfigureBackupPlanResponse) {
    response = &ConfigureBackupPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ConfigureBackupPlan
// 本接口（ConfigureBackupPlan）用于配置备份计划。包括配置备份源实例信息、备份对象以及备份策略等。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ConfigureBackupPlan(request *ConfigureBackupPlanRequest) (response *ConfigureBackupPlanResponse, err error) {
    return c.ConfigureBackupPlanWithContext(context.Background(), request)
}

// ConfigureBackupPlan
// 本接口（ConfigureBackupPlan）用于配置备份计划。包括配置备份源实例信息、备份对象以及备份策略等。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ConfigureBackupPlanWithContext(ctx context.Context, request *ConfigureBackupPlanRequest) (response *ConfigureBackupPlanResponse, err error) {
    if request == nil {
        request = NewConfigureBackupPlanRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbs", APIVersion, "ConfigureBackupPlan")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ConfigureBackupPlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewConfigureBackupPlanResponse()
    err = c.Send(request, response)
    return
}

func NewCreateConnectTestJobRequest() (request *CreateConnectTestJobRequest) {
    request = &CreateConnectTestJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbs", APIVersion, "CreateConnectTestJob")
    
    
    return
}

func NewCreateConnectTestJobResponse() (response *CreateConnectTestJobResponse) {
    response = &CreateConnectTestJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateConnectTestJob
// 该接口用于创建连通性检测任务，请在创建备份计划前，通过该接口来检测你的源端实例是否连通性正常。
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateConnectTestJob(request *CreateConnectTestJobRequest) (response *CreateConnectTestJobResponse, err error) {
    return c.CreateConnectTestJobWithContext(context.Background(), request)
}

// CreateConnectTestJob
// 该接口用于创建连通性检测任务，请在创建备份计划前，通过该接口来检测你的源端实例是否连通性正常。
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateConnectTestJobWithContext(ctx context.Context, request *CreateConnectTestJobRequest) (response *CreateConnectTestJobResponse, err error) {
    if request == nil {
        request = NewCreateConnectTestJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbs", APIVersion, "CreateConnectTestJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateConnectTestJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateConnectTestJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupCheckJobRequest() (request *DescribeBackupCheckJobRequest) {
    request = &DescribeBackupCheckJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbs", APIVersion, "DescribeBackupCheckJob")
    
    
    return
}

func NewDescribeBackupCheckJobResponse() (response *DescribeBackupCheckJobResponse) {
    response = &DescribeBackupCheckJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackupCheckJob
// 本接口（DescribeBackupCheckJob）用于查询备份计划预校验任务的结果。仅对于预校验通过的任务，才能启动备份计划。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeBackupCheckJob(request *DescribeBackupCheckJobRequest) (response *DescribeBackupCheckJobResponse, err error) {
    return c.DescribeBackupCheckJobWithContext(context.Background(), request)
}

// DescribeBackupCheckJob
// 本接口（DescribeBackupCheckJob）用于查询备份计划预校验任务的结果。仅对于预校验通过的任务，才能启动备份计划。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeBackupCheckJobWithContext(ctx context.Context, request *DescribeBackupCheckJobRequest) (response *DescribeBackupCheckJobResponse, err error) {
    if request == nil {
        request = NewDescribeBackupCheckJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbs", APIVersion, "DescribeBackupCheckJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupCheckJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupCheckJobResponse()
    err = c.Send(request, response)
    return
}

func NewStartBackupCheckJobRequest() (request *StartBackupCheckJobRequest) {
    request = &StartBackupCheckJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbs", APIVersion, "StartBackupCheckJob")
    
    
    return
}

func NewStartBackupCheckJobResponse() (response *StartBackupCheckJobResponse) {
    response = &StartBackupCheckJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartBackupCheckJob
// 本接口（StartBackupCheckJob）用于创建备份计划预校验任务。
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) StartBackupCheckJob(request *StartBackupCheckJobRequest) (response *StartBackupCheckJobResponse, err error) {
    return c.StartBackupCheckJobWithContext(context.Background(), request)
}

// StartBackupCheckJob
// 本接口（StartBackupCheckJob）用于创建备份计划预校验任务。
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) StartBackupCheckJobWithContext(ctx context.Context, request *StartBackupCheckJobRequest) (response *StartBackupCheckJobResponse, err error) {
    if request == nil {
        request = NewStartBackupCheckJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbs", APIVersion, "StartBackupCheckJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartBackupCheckJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartBackupCheckJobResponse()
    err = c.Send(request, response)
    return
}

func NewStartBackupPlanRequest() (request *StartBackupPlanRequest) {
    request = &StartBackupPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbs", APIVersion, "StartBackupPlan")
    
    
    return
}

func NewStartBackupPlanResponse() (response *StartBackupPlanResponse) {
    response = &StartBackupPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartBackupPlan
// 本接口（StartBackupPlan）用于启动备份计划。调用此接口前，请务必先使用 StartBackupCheckJob 创建备份计划预校验任务，并通过 DescribeBackupCheckJob 接口查询到任务状态为校验通过时，才能启动备份计划。
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) StartBackupPlan(request *StartBackupPlanRequest) (response *StartBackupPlanResponse, err error) {
    return c.StartBackupPlanWithContext(context.Background(), request)
}

// StartBackupPlan
// 本接口（StartBackupPlan）用于启动备份计划。调用此接口前，请务必先使用 StartBackupCheckJob 创建备份计划预校验任务，并通过 DescribeBackupCheckJob 接口查询到任务状态为校验通过时，才能启动备份计划。
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) StartBackupPlanWithContext(ctx context.Context, request *StartBackupPlanRequest) (response *StartBackupPlanResponse, err error) {
    if request == nil {
        request = NewStartBackupPlanRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbs", APIVersion, "StartBackupPlan")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartBackupPlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartBackupPlanResponse()
    err = c.Send(request, response)
    return
}
