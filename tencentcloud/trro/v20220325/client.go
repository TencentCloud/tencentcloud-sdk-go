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

package v20220325

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2022-03-25"

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


func NewBatchDeleteDevicesRequest() (request *BatchDeleteDevicesRequest) {
    request = &BatchDeleteDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trro", APIVersion, "BatchDeleteDevices")
    
    
    return
}

func NewBatchDeleteDevicesResponse() (response *BatchDeleteDevicesResponse) {
    response = &BatchDeleteDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchDeleteDevices
// 用于批量删除设备
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) BatchDeleteDevices(request *BatchDeleteDevicesRequest) (response *BatchDeleteDevicesResponse, err error) {
    return c.BatchDeleteDevicesWithContext(context.Background(), request)
}

// BatchDeleteDevices
// 用于批量删除设备
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) BatchDeleteDevicesWithContext(ctx context.Context, request *BatchDeleteDevicesRequest) (response *BatchDeleteDevicesResponse, err error) {
    if request == nil {
        request = NewBatchDeleteDevicesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchDeleteDevices require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchDeleteDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewBatchDeletePolicyRequest() (request *BatchDeletePolicyRequest) {
    request = &BatchDeletePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trro", APIVersion, "BatchDeletePolicy")
    
    
    return
}

func NewBatchDeletePolicyResponse() (response *BatchDeletePolicyResponse) {
    response = &BatchDeletePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchDeletePolicy
// 用于批量删除修改权限配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) BatchDeletePolicy(request *BatchDeletePolicyRequest) (response *BatchDeletePolicyResponse, err error) {
    return c.BatchDeletePolicyWithContext(context.Background(), request)
}

// BatchDeletePolicy
// 用于批量删除修改权限配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) BatchDeletePolicyWithContext(ctx context.Context, request *BatchDeletePolicyRequest) (response *BatchDeletePolicyResponse, err error) {
    if request == nil {
        request = NewBatchDeletePolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchDeletePolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchDeletePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewBoundLicensesRequest() (request *BoundLicensesRequest) {
    request = &BoundLicensesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trro", APIVersion, "BoundLicenses")
    
    
    return
}

func NewBoundLicensesResponse() (response *BoundLicensesResponse) {
    response = &BoundLicensesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BoundLicenses
// 为推流设备绑定license，优先绑定到期时间最近的，到期时间相同优先绑定月包
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) BoundLicenses(request *BoundLicensesRequest) (response *BoundLicensesResponse, err error) {
    return c.BoundLicensesWithContext(context.Background(), request)
}

// BoundLicenses
// 为推流设备绑定license，优先绑定到期时间最近的，到期时间相同优先绑定月包
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) BoundLicensesWithContext(ctx context.Context, request *BoundLicensesRequest) (response *BoundLicensesResponse, err error) {
    if request == nil {
        request = NewBoundLicensesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BoundLicenses require credential")
    }

    request.SetContext(ctx)
    
    response = NewBoundLicensesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDeviceRequest() (request *CreateDeviceRequest) {
    request = &CreateDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trro", APIVersion, "CreateDevice")
    
    
    return
}

func NewCreateDeviceResponse() (response *CreateDeviceResponse) {
    response = &CreateDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDevice
// 用于创建设备
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateDevice(request *CreateDeviceRequest) (response *CreateDeviceResponse, err error) {
    return c.CreateDeviceWithContext(context.Background(), request)
}

// CreateDevice
// 用于创建设备
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateDeviceWithContext(ctx context.Context, request *CreateDeviceRequest) (response *CreateDeviceResponse, err error) {
    if request == nil {
        request = NewCreateDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProjectRequest() (request *CreateProjectRequest) {
    request = &CreateProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trro", APIVersion, "CreateProject")
    
    
    return
}

func NewCreateProjectResponse() (response *CreateProjectResponse) {
    response = &CreateProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateProject
// 用于创建项目
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateProject(request *CreateProjectRequest) (response *CreateProjectResponse, err error) {
    return c.CreateProjectWithContext(context.Background(), request)
}

// CreateProject
// 用于创建项目
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateProjectWithContext(ctx context.Context, request *CreateProjectRequest) (response *CreateProjectResponse, err error) {
    if request == nil {
        request = NewCreateProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateProjectResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteProjectRequest() (request *DeleteProjectRequest) {
    request = &DeleteProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trro", APIVersion, "DeleteProject")
    
    
    return
}

func NewDeleteProjectResponse() (response *DeleteProjectResponse) {
    response = &DeleteProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteProject
// 用于删除项目
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DeleteProject(request *DeleteProjectRequest) (response *DeleteProjectResponse, err error) {
    return c.DeleteProjectWithContext(context.Background(), request)
}

// DeleteProject
// 用于删除项目
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DeleteProjectWithContext(ctx context.Context, request *DeleteProjectRequest) (response *DeleteProjectResponse, err error) {
    if request == nil {
        request = NewDeleteProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteProjectResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceInfoRequest() (request *DescribeDeviceInfoRequest) {
    request = &DescribeDeviceInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trro", APIVersion, "DescribeDeviceInfo")
    
    
    return
}

func NewDescribeDeviceInfoResponse() (response *DescribeDeviceInfoResponse) {
    response = &DescribeDeviceInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeviceInfo
// 用于获取指定设备信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeDeviceInfo(request *DescribeDeviceInfoRequest) (response *DescribeDeviceInfoResponse, err error) {
    return c.DescribeDeviceInfoWithContext(context.Background(), request)
}

// DescribeDeviceInfo
// 用于获取指定设备信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeDeviceInfoWithContext(ctx context.Context, request *DescribeDeviceInfoRequest) (response *DescribeDeviceInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceListRequest() (request *DescribeDeviceListRequest) {
    request = &DescribeDeviceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trro", APIVersion, "DescribeDeviceList")
    
    
    return
}

func NewDescribeDeviceListResponse() (response *DescribeDeviceListResponse) {
    response = &DescribeDeviceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeviceList
// 用于获取设备信息列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeDeviceList(request *DescribeDeviceListRequest) (response *DescribeDeviceListResponse, err error) {
    return c.DescribeDeviceListWithContext(context.Background(), request)
}

// DescribeDeviceList
// 用于获取设备信息列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeDeviceListWithContext(ctx context.Context, request *DescribeDeviceListRequest) (response *DescribeDeviceListResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceSessionDetailsRequest() (request *DescribeDeviceSessionDetailsRequest) {
    request = &DescribeDeviceSessionDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trro", APIVersion, "DescribeDeviceSessionDetails")
    
    
    return
}

func NewDescribeDeviceSessionDetailsResponse() (response *DescribeDeviceSessionDetailsResponse) {
    response = &DescribeDeviceSessionDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeviceSessionDetails
// 获取设备会话数据详单
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeDeviceSessionDetails(request *DescribeDeviceSessionDetailsRequest) (response *DescribeDeviceSessionDetailsResponse, err error) {
    return c.DescribeDeviceSessionDetailsWithContext(context.Background(), request)
}

// DescribeDeviceSessionDetails
// 获取设备会话数据详单
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeDeviceSessionDetailsWithContext(ctx context.Context, request *DescribeDeviceSessionDetailsRequest) (response *DescribeDeviceSessionDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceSessionDetailsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceSessionDetails require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceSessionDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceSessionListRequest() (request *DescribeDeviceSessionListRequest) {
    request = &DescribeDeviceSessionListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trro", APIVersion, "DescribeDeviceSessionList")
    
    
    return
}

func NewDescribeDeviceSessionListResponse() (response *DescribeDeviceSessionListResponse) {
    response = &DescribeDeviceSessionListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeviceSessionList
// 获取设备会话列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeDeviceSessionList(request *DescribeDeviceSessionListRequest) (response *DescribeDeviceSessionListResponse, err error) {
    return c.DescribeDeviceSessionListWithContext(context.Background(), request)
}

// DescribeDeviceSessionList
// 获取设备会话列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeDeviceSessionListWithContext(ctx context.Context, request *DescribeDeviceSessionListRequest) (response *DescribeDeviceSessionListResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceSessionListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceSessionList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceSessionListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePolicyRequest() (request *DescribePolicyRequest) {
    request = &DescribePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trro", APIVersion, "DescribePolicy")
    
    
    return
}

func NewDescribePolicyResponse() (response *DescribePolicyResponse) {
    response = &DescribePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePolicy
// 用于查看权限配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribePolicy(request *DescribePolicyRequest) (response *DescribePolicyResponse, err error) {
    return c.DescribePolicyWithContext(context.Background(), request)
}

// DescribePolicy
// 用于查看权限配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribePolicyWithContext(ctx context.Context, request *DescribePolicyRequest) (response *DescribePolicyResponse, err error) {
    if request == nil {
        request = NewDescribePolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectInfoRequest() (request *DescribeProjectInfoRequest) {
    request = &DescribeProjectInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trro", APIVersion, "DescribeProjectInfo")
    
    
    return
}

func NewDescribeProjectInfoResponse() (response *DescribeProjectInfoResponse) {
    response = &DescribeProjectInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProjectInfo
// 用于获取项目信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeProjectInfo(request *DescribeProjectInfoRequest) (response *DescribeProjectInfoResponse, err error) {
    return c.DescribeProjectInfoWithContext(context.Background(), request)
}

// DescribeProjectInfo
// 用于获取项目信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeProjectInfoWithContext(ctx context.Context, request *DescribeProjectInfoRequest) (response *DescribeProjectInfoResponse, err error) {
    if request == nil {
        request = NewDescribeProjectInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProjectInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProjectInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectListRequest() (request *DescribeProjectListRequest) {
    request = &DescribeProjectListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trro", APIVersion, "DescribeProjectList")
    
    
    return
}

func NewDescribeProjectListResponse() (response *DescribeProjectListResponse) {
    response = &DescribeProjectListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProjectList
// 用于获取项目列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeProjectList(request *DescribeProjectListRequest) (response *DescribeProjectListResponse, err error) {
    return c.DescribeProjectListWithContext(context.Background(), request)
}

// DescribeProjectList
// 用于获取项目列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeProjectListWithContext(ctx context.Context, request *DescribeProjectListRequest) (response *DescribeProjectListResponse, err error) {
    if request == nil {
        request = NewDescribeProjectListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProjectList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProjectListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecentSessionListRequest() (request *DescribeRecentSessionListRequest) {
    request = &DescribeRecentSessionListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trro", APIVersion, "DescribeRecentSessionList")
    
    
    return
}

func NewDescribeRecentSessionListResponse() (response *DescribeRecentSessionListResponse) {
    response = &DescribeRecentSessionListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRecentSessionList
// 获取最新设备会话列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeRecentSessionList(request *DescribeRecentSessionListRequest) (response *DescribeRecentSessionListResponse, err error) {
    return c.DescribeRecentSessionListWithContext(context.Background(), request)
}

// DescribeRecentSessionList
// 获取最新设备会话列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeRecentSessionListWithContext(ctx context.Context, request *DescribeRecentSessionListRequest) (response *DescribeRecentSessionListResponse, err error) {
    if request == nil {
        request = NewDescribeRecentSessionListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecentSessionList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecentSessionListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSessionStatisticsRequest() (request *DescribeSessionStatisticsRequest) {
    request = &DescribeSessionStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trro", APIVersion, "DescribeSessionStatistics")
    
    
    return
}

func NewDescribeSessionStatisticsResponse() (response *DescribeSessionStatisticsResponse) {
    response = &DescribeSessionStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSessionStatistics
// 获取会话统计值
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSessionStatistics(request *DescribeSessionStatisticsRequest) (response *DescribeSessionStatisticsResponse, err error) {
    return c.DescribeSessionStatisticsWithContext(context.Background(), request)
}

// DescribeSessionStatistics
// 获取会话统计值
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSessionStatisticsWithContext(ctx context.Context, request *DescribeSessionStatisticsRequest) (response *DescribeSessionStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeSessionStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSessionStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSessionStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSessionStatisticsByIntervalRequest() (request *DescribeSessionStatisticsByIntervalRequest) {
    request = &DescribeSessionStatisticsByIntervalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trro", APIVersion, "DescribeSessionStatisticsByInterval")
    
    
    return
}

func NewDescribeSessionStatisticsByIntervalResponse() (response *DescribeSessionStatisticsByIntervalResponse) {
    response = &DescribeSessionStatisticsByIntervalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSessionStatisticsByInterval
// 获取各时间段的会话统计值
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSessionStatisticsByInterval(request *DescribeSessionStatisticsByIntervalRequest) (response *DescribeSessionStatisticsByIntervalResponse, err error) {
    return c.DescribeSessionStatisticsByIntervalWithContext(context.Background(), request)
}

// DescribeSessionStatisticsByInterval
// 获取各时间段的会话统计值
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSessionStatisticsByIntervalWithContext(ctx context.Context, request *DescribeSessionStatisticsByIntervalRequest) (response *DescribeSessionStatisticsByIntervalResponse, err error) {
    if request == nil {
        request = NewDescribeSessionStatisticsByIntervalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSessionStatisticsByInterval require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSessionStatisticsByIntervalResponse()
    err = c.Send(request, response)
    return
}

func NewGetDeviceLicenseRequest() (request *GetDeviceLicenseRequest) {
    request = &GetDeviceLicenseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trro", APIVersion, "GetDeviceLicense")
    
    
    return
}

func NewGetDeviceLicenseResponse() (response *GetDeviceLicenseResponse) {
    response = &GetDeviceLicenseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetDeviceLicense
// 获取设备已经绑定的可用授权数量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) GetDeviceLicense(request *GetDeviceLicenseRequest) (response *GetDeviceLicenseResponse, err error) {
    return c.GetDeviceLicenseWithContext(context.Background(), request)
}

// GetDeviceLicense
// 获取设备已经绑定的可用授权数量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) GetDeviceLicenseWithContext(ctx context.Context, request *GetDeviceLicenseRequest) (response *GetDeviceLicenseResponse, err error) {
    if request == nil {
        request = NewGetDeviceLicenseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDeviceLicense require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDeviceLicenseResponse()
    err = c.Send(request, response)
    return
}

func NewGetDevicesRequest() (request *GetDevicesRequest) {
    request = &GetDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trro", APIVersion, "GetDevices")
    
    
    return
}

func NewGetDevicesResponse() (response *GetDevicesResponse) {
    response = &GetDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetDevices
// 查询用户设备的授权绑定情况
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetDevices(request *GetDevicesRequest) (response *GetDevicesResponse, err error) {
    return c.GetDevicesWithContext(context.Background(), request)
}

// GetDevices
// 查询用户设备的授权绑定情况
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetDevicesWithContext(ctx context.Context, request *GetDevicesRequest) (response *GetDevicesResponse, err error) {
    if request == nil {
        request = NewGetDevicesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDevices require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewGetLicenseStatRequest() (request *GetLicenseStatRequest) {
    request = &GetLicenseStatRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trro", APIVersion, "GetLicenseStat")
    
    
    return
}

func NewGetLicenseStatResponse() (response *GetLicenseStatResponse) {
    response = &GetLicenseStatResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetLicenseStat
// 统计license类型数量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) GetLicenseStat(request *GetLicenseStatRequest) (response *GetLicenseStatResponse, err error) {
    return c.GetLicenseStatWithContext(context.Background(), request)
}

// GetLicenseStat
// 统计license类型数量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) GetLicenseStatWithContext(ctx context.Context, request *GetLicenseStatRequest) (response *GetLicenseStatResponse, err error) {
    if request == nil {
        request = NewGetLicenseStatRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetLicenseStat require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetLicenseStatResponse()
    err = c.Send(request, response)
    return
}

func NewGetLicensesRequest() (request *GetLicensesRequest) {
    request = &GetLicensesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trro", APIVersion, "GetLicenses")
    
    
    return
}

func NewGetLicensesResponse() (response *GetLicensesResponse) {
    response = &GetLicensesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetLicenses
// 按授权查看license列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetLicenses(request *GetLicensesRequest) (response *GetLicensesResponse, err error) {
    return c.GetLicensesWithContext(context.Background(), request)
}

// GetLicenses
// 按授权查看license列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) GetLicensesWithContext(ctx context.Context, request *GetLicensesRequest) (response *GetLicensesResponse, err error) {
    if request == nil {
        request = NewGetLicensesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetLicenses require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetLicensesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDeviceRequest() (request *ModifyDeviceRequest) {
    request = &ModifyDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trro", APIVersion, "ModifyDevice")
    
    
    return
}

func NewModifyDeviceResponse() (response *ModifyDeviceResponse) {
    response = &ModifyDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDevice
// 用于修改设备信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyDevice(request *ModifyDeviceRequest) (response *ModifyDeviceResponse, err error) {
    return c.ModifyDeviceWithContext(context.Background(), request)
}

// ModifyDevice
// 用于修改设备信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyDeviceWithContext(ctx context.Context, request *ModifyDeviceRequest) (response *ModifyDeviceResponse, err error) {
    if request == nil {
        request = NewModifyDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPolicyRequest() (request *ModifyPolicyRequest) {
    request = &ModifyPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trro", APIVersion, "ModifyPolicy")
    
    
    return
}

func NewModifyPolicyResponse() (response *ModifyPolicyResponse) {
    response = &ModifyPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPolicy
// 用于修改权限配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyPolicy(request *ModifyPolicyRequest) (response *ModifyPolicyResponse, err error) {
    return c.ModifyPolicyWithContext(context.Background(), request)
}

// ModifyPolicy
// 用于修改权限配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyPolicyWithContext(ctx context.Context, request *ModifyPolicyRequest) (response *ModifyPolicyResponse, err error) {
    if request == nil {
        request = NewModifyPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyProjectRequest() (request *ModifyProjectRequest) {
    request = &ModifyProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trro", APIVersion, "ModifyProject")
    
    
    return
}

func NewModifyProjectResponse() (response *ModifyProjectResponse) {
    response = &ModifyProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyProject
// 用于修改项目信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyProject(request *ModifyProjectRequest) (response *ModifyProjectResponse, err error) {
    return c.ModifyProjectWithContext(context.Background(), request)
}

// ModifyProject
// 用于修改项目信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyProjectWithContext(ctx context.Context, request *ModifyProjectRequest) (response *ModifyProjectResponse, err error) {
    if request == nil {
        request = NewModifyProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyProjectResponse()
    err = c.Send(request, response)
    return
}
