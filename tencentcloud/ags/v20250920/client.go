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

package v20250920

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2025-09-20"

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


func NewAcquireSandboxInstanceTokenRequest() (request *AcquireSandboxInstanceTokenRequest) {
    request = &AcquireSandboxInstanceTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "AcquireSandboxInstanceToken")
    
    
    return
}

func NewAcquireSandboxInstanceTokenResponse() (response *AcquireSandboxInstanceTokenResponse) {
    response = &AcquireSandboxInstanceTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AcquireSandboxInstanceToken
// 获取访问沙箱工具时所需要使用的访问Token，创建沙箱实例后需调用此接口获取沙箱实例访问Token。
//
// 此Token可用于调用代码沙箱实例执行代码，或浏览器沙箱实例进行浏览器操作等。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AcquireSandboxInstanceToken(request *AcquireSandboxInstanceTokenRequest) (response *AcquireSandboxInstanceTokenResponse, err error) {
    return c.AcquireSandboxInstanceTokenWithContext(context.Background(), request)
}

// AcquireSandboxInstanceToken
// 获取访问沙箱工具时所需要使用的访问Token，创建沙箱实例后需调用此接口获取沙箱实例访问Token。
//
// 此Token可用于调用代码沙箱实例执行代码，或浏览器沙箱实例进行浏览器操作等。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AcquireSandboxInstanceTokenWithContext(ctx context.Context, request *AcquireSandboxInstanceTokenRequest) (response *AcquireSandboxInstanceTokenResponse, err error) {
    if request == nil {
        request = NewAcquireSandboxInstanceTokenRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "AcquireSandboxInstanceToken")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AcquireSandboxInstanceToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewAcquireSandboxInstanceTokenResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAPIKeyRequest() (request *CreateAPIKeyRequest) {
    request = &CreateAPIKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "CreateAPIKey")
    
    
    return
}

func NewCreateAPIKeyResponse() (response *CreateAPIKeyResponse) {
    response = &CreateAPIKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAPIKey
// 创建新的API密钥，用于调用Agent Sandbox接口。相较于腾讯云Secret ID Secret Key支持调用所有接口使用，仅有部分接口支持使用API密钥调用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_APIKEYQUOTA = "LimitExceeded.APIKeyQuota"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAPIKey(request *CreateAPIKeyRequest) (response *CreateAPIKeyResponse, err error) {
    return c.CreateAPIKeyWithContext(context.Background(), request)
}

// CreateAPIKey
// 创建新的API密钥，用于调用Agent Sandbox接口。相较于腾讯云Secret ID Secret Key支持调用所有接口使用，仅有部分接口支持使用API密钥调用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_APIKEYQUOTA = "LimitExceeded.APIKeyQuota"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAPIKeyWithContext(ctx context.Context, request *CreateAPIKeyRequest) (response *CreateAPIKeyResponse, err error) {
    if request == nil {
        request = NewCreateAPIKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "CreateAPIKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAPIKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAPIKeyResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePreCacheImageTaskRequest() (request *CreatePreCacheImageTaskRequest) {
    request = &CreatePreCacheImageTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "CreatePreCacheImageTask")
    
    
    return
}

func NewCreatePreCacheImageTaskResponse() (response *CreatePreCacheImageTaskResponse) {
    response = &CreatePreCacheImageTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePreCacheImageTask
// 创建镜像预热任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreatePreCacheImageTask(request *CreatePreCacheImageTaskRequest) (response *CreatePreCacheImageTaskResponse, err error) {
    return c.CreatePreCacheImageTaskWithContext(context.Background(), request)
}

// CreatePreCacheImageTask
// 创建镜像预热任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreatePreCacheImageTaskWithContext(ctx context.Context, request *CreatePreCacheImageTaskRequest) (response *CreatePreCacheImageTaskResponse, err error) {
    if request == nil {
        request = NewCreatePreCacheImageTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "CreatePreCacheImageTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePreCacheImageTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePreCacheImageTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSandboxToolRequest() (request *CreateSandboxToolRequest) {
    request = &CreateSandboxToolRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "CreateSandboxTool")
    
    
    return
}

func NewCreateSandboxToolResponse() (response *CreateSandboxToolResponse) {
    response = &CreateSandboxToolResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSandboxTool
// 创建沙箱工具
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DUPLICATEREQUEST = "FailedOperation.DuplicateRequest"
//  FAILEDOPERATION_REQUESTINPROGRESS = "FailedOperation.RequestInProgress"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_NETWORKSETUPFAILED = "InternalError.NetworkSetupFailed"
//  INTERNALERROR_VPCSERVICEUNAVAILABLE = "InternalError.VPCServiceUnavailable"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ROLEARN = "InvalidParameterValue.RoleArn"
//  INVALIDPARAMETERVALUE_SANDBOXTOOL = "InvalidParameterValue.SandboxTool"
//  INVALIDPARAMETERVALUE_SECURITYGROUPID = "InvalidParameterValue.SecurityGroupId"
//  INVALIDPARAMETERVALUE_STORAGEMOUNT = "InvalidParameterValue.StorageMount"
//  INVALIDPARAMETERVALUE_SUBNETID = "InvalidParameterValue.SubnetId"
//  INVALIDPARAMETERVALUE_TIMEOUT = "InvalidParameterValue.Timeout"
//  INVALIDPARAMETERVALUE_TOOLTYPE = "InvalidParameterValue.ToolType"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_ROLEARN = "MissingParameter.RoleArn"
//  MISSINGPARAMETER_VPCPARAMETERS = "MissingParameter.VPCParameters"
//  RESOURCENOTFOUND_SECURITYGROUP = "ResourceNotFound.SecurityGroup"
//  RESOURCENOTFOUND_STORAGEMOUNT = "ResourceNotFound.StorageMount"
//  RESOURCENOTFOUND_SUBNET = "ResourceNotFound.Subnet"
func (c *Client) CreateSandboxTool(request *CreateSandboxToolRequest) (response *CreateSandboxToolResponse, err error) {
    return c.CreateSandboxToolWithContext(context.Background(), request)
}

// CreateSandboxTool
// 创建沙箱工具
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DUPLICATEREQUEST = "FailedOperation.DuplicateRequest"
//  FAILEDOPERATION_REQUESTINPROGRESS = "FailedOperation.RequestInProgress"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_NETWORKSETUPFAILED = "InternalError.NetworkSetupFailed"
//  INTERNALERROR_VPCSERVICEUNAVAILABLE = "InternalError.VPCServiceUnavailable"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ROLEARN = "InvalidParameterValue.RoleArn"
//  INVALIDPARAMETERVALUE_SANDBOXTOOL = "InvalidParameterValue.SandboxTool"
//  INVALIDPARAMETERVALUE_SECURITYGROUPID = "InvalidParameterValue.SecurityGroupId"
//  INVALIDPARAMETERVALUE_STORAGEMOUNT = "InvalidParameterValue.StorageMount"
//  INVALIDPARAMETERVALUE_SUBNETID = "InvalidParameterValue.SubnetId"
//  INVALIDPARAMETERVALUE_TIMEOUT = "InvalidParameterValue.Timeout"
//  INVALIDPARAMETERVALUE_TOOLTYPE = "InvalidParameterValue.ToolType"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_ROLEARN = "MissingParameter.RoleArn"
//  MISSINGPARAMETER_VPCPARAMETERS = "MissingParameter.VPCParameters"
//  RESOURCENOTFOUND_SECURITYGROUP = "ResourceNotFound.SecurityGroup"
//  RESOURCENOTFOUND_STORAGEMOUNT = "ResourceNotFound.StorageMount"
//  RESOURCENOTFOUND_SUBNET = "ResourceNotFound.Subnet"
func (c *Client) CreateSandboxToolWithContext(ctx context.Context, request *CreateSandboxToolRequest) (response *CreateSandboxToolResponse, err error) {
    if request == nil {
        request = NewCreateSandboxToolRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "CreateSandboxTool")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSandboxTool require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSandboxToolResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAPIKeyRequest() (request *DeleteAPIKeyRequest) {
    request = &DeleteAPIKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "DeleteAPIKey")
    
    
    return
}

func NewDeleteAPIKeyResponse() (response *DeleteAPIKeyResponse) {
    response = &DeleteAPIKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAPIKey
// 删除API密钥。注意区别于腾讯云Secret ID Secret Key，本接口删除的是Agent Sandbox专用API key。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteAPIKey(request *DeleteAPIKeyRequest) (response *DeleteAPIKeyResponse, err error) {
    return c.DeleteAPIKeyWithContext(context.Background(), request)
}

// DeleteAPIKey
// 删除API密钥。注意区别于腾讯云Secret ID Secret Key，本接口删除的是Agent Sandbox专用API key。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteAPIKeyWithContext(ctx context.Context, request *DeleteAPIKeyRequest) (response *DeleteAPIKeyResponse, err error) {
    if request == nil {
        request = NewDeleteAPIKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "DeleteAPIKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAPIKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAPIKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSandboxToolRequest() (request *DeleteSandboxToolRequest) {
    request = &DeleteSandboxToolRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "DeleteSandboxTool")
    
    
    return
}

func NewDeleteSandboxToolResponse() (response *DeleteSandboxToolResponse) {
    response = &DeleteSandboxToolResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSandboxTool
// 删除沙箱工具
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE_SANDBOXTOOL = "ResourceInUse.SandboxTool"
//  RESOURCENOTFOUND_SANDBOXTOOL = "ResourceNotFound.SandboxTool"
func (c *Client) DeleteSandboxTool(request *DeleteSandboxToolRequest) (response *DeleteSandboxToolResponse, err error) {
    return c.DeleteSandboxToolWithContext(context.Background(), request)
}

// DeleteSandboxTool
// 删除沙箱工具
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE_SANDBOXTOOL = "ResourceInUse.SandboxTool"
//  RESOURCENOTFOUND_SANDBOXTOOL = "ResourceNotFound.SandboxTool"
func (c *Client) DeleteSandboxToolWithContext(ctx context.Context, request *DeleteSandboxToolRequest) (response *DeleteSandboxToolResponse, err error) {
    if request == nil {
        request = NewDeleteSandboxToolRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "DeleteSandboxTool")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSandboxTool require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSandboxToolResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAPIKeyListRequest() (request *DescribeAPIKeyListRequest) {
    request = &DescribeAPIKeyListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "DescribeAPIKeyList")
    
    
    return
}

func NewDescribeAPIKeyListResponse() (response *DescribeAPIKeyListResponse) {
    response = &DescribeAPIKeyListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAPIKeyList
// 获取API密钥列表，包含API密钥简略信息，包含名称、创建时间等。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAPIKeyList(request *DescribeAPIKeyListRequest) (response *DescribeAPIKeyListResponse, err error) {
    return c.DescribeAPIKeyListWithContext(context.Background(), request)
}

// DescribeAPIKeyList
// 获取API密钥列表，包含API密钥简略信息，包含名称、创建时间等。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAPIKeyListWithContext(ctx context.Context, request *DescribeAPIKeyListRequest) (response *DescribeAPIKeyListResponse, err error) {
    if request == nil {
        request = NewDescribeAPIKeyListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "DescribeAPIKeyList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAPIKeyList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAPIKeyListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePreCacheImageTaskRequest() (request *DescribePreCacheImageTaskRequest) {
    request = &DescribePreCacheImageTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "DescribePreCacheImageTask")
    
    
    return
}

func NewDescribePreCacheImageTaskResponse() (response *DescribePreCacheImageTaskResponse) {
    response = &DescribePreCacheImageTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePreCacheImageTask
// 查询镜像预热任务信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribePreCacheImageTask(request *DescribePreCacheImageTaskRequest) (response *DescribePreCacheImageTaskResponse, err error) {
    return c.DescribePreCacheImageTaskWithContext(context.Background(), request)
}

// DescribePreCacheImageTask
// 查询镜像预热任务信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribePreCacheImageTaskWithContext(ctx context.Context, request *DescribePreCacheImageTaskRequest) (response *DescribePreCacheImageTaskResponse, err error) {
    if request == nil {
        request = NewDescribePreCacheImageTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "DescribePreCacheImageTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePreCacheImageTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePreCacheImageTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSandboxInstanceListRequest() (request *DescribeSandboxInstanceListRequest) {
    request = &DescribeSandboxInstanceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "DescribeSandboxInstanceList")
    
    
    return
}

func NewDescribeSandboxInstanceListResponse() (response *DescribeSandboxInstanceListResponse) {
    response = &DescribeSandboxInstanceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSandboxInstanceList
// 查询沙箱实例列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCEIDS = "InvalidParameterValue.InstanceIds"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeSandboxInstanceList(request *DescribeSandboxInstanceListRequest) (response *DescribeSandboxInstanceListResponse, err error) {
    return c.DescribeSandboxInstanceListWithContext(context.Background(), request)
}

// DescribeSandboxInstanceList
// 查询沙箱实例列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCEIDS = "InvalidParameterValue.InstanceIds"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeSandboxInstanceListWithContext(ctx context.Context, request *DescribeSandboxInstanceListRequest) (response *DescribeSandboxInstanceListResponse, err error) {
    if request == nil {
        request = NewDescribeSandboxInstanceListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "DescribeSandboxInstanceList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSandboxInstanceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSandboxInstanceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSandboxToolListRequest() (request *DescribeSandboxToolListRequest) {
    request = &DescribeSandboxToolListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "DescribeSandboxToolList")
    
    
    return
}

func NewDescribeSandboxToolListResponse() (response *DescribeSandboxToolListResponse) {
    response = &DescribeSandboxToolListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSandboxToolList
// 查询沙箱工具列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_TOOLIDS = "InvalidParameterValue.ToolIds"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeSandboxToolList(request *DescribeSandboxToolListRequest) (response *DescribeSandboxToolListResponse, err error) {
    return c.DescribeSandboxToolListWithContext(context.Background(), request)
}

// DescribeSandboxToolList
// 查询沙箱工具列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_TOOLIDS = "InvalidParameterValue.ToolIds"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeSandboxToolListWithContext(ctx context.Context, request *DescribeSandboxToolListRequest) (response *DescribeSandboxToolListResponse, err error) {
    if request == nil {
        request = NewDescribeSandboxToolListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "DescribeSandboxToolList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSandboxToolList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSandboxToolListResponse()
    err = c.Send(request, response)
    return
}

func NewStartSandboxInstanceRequest() (request *StartSandboxInstanceRequest) {
    request = &StartSandboxInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "StartSandboxInstance")
    
    
    return
}

func NewStartSandboxInstanceResponse() (response *StartSandboxInstanceResponse) {
    response = &StartSandboxInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartSandboxInstance
// 启动沙箱实例
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DUPLICATEREQUEST = "FailedOperation.DuplicateRequest"
//  FAILEDOPERATION_REQUESTINPROGRESS = "FailedOperation.RequestInProgress"
//  FAILEDOPERATION_STORAGEMOUNT = "FailedOperation.StorageMount"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_MOUNTOPTION = "InvalidParameterValue.MountOption"
//  INVALIDPARAMETERVALUE_STORAGEMOUNT = "InvalidParameterValue.StorageMount"
//  INVALIDPARAMETERVALUE_TIMEOUT = "InvalidParameterValue.Timeout"
//  LIMITEXCEEDED_SANDBOXINSTANCE = "LimitExceeded.SandboxInstance"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_SANDBOXTOOL = "ResourceNotFound.SandboxTool"
//  RESOURCENOTFOUND_STORAGEMOUNT = "ResourceNotFound.StorageMount"
//  RESOURCEUNAVAILABLE_SANDBOXTOOL = "ResourceUnavailable.SandboxTool"
func (c *Client) StartSandboxInstance(request *StartSandboxInstanceRequest) (response *StartSandboxInstanceResponse, err error) {
    return c.StartSandboxInstanceWithContext(context.Background(), request)
}

// StartSandboxInstance
// 启动沙箱实例
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DUPLICATEREQUEST = "FailedOperation.DuplicateRequest"
//  FAILEDOPERATION_REQUESTINPROGRESS = "FailedOperation.RequestInProgress"
//  FAILEDOPERATION_STORAGEMOUNT = "FailedOperation.StorageMount"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_MOUNTOPTION = "InvalidParameterValue.MountOption"
//  INVALIDPARAMETERVALUE_STORAGEMOUNT = "InvalidParameterValue.StorageMount"
//  INVALIDPARAMETERVALUE_TIMEOUT = "InvalidParameterValue.Timeout"
//  LIMITEXCEEDED_SANDBOXINSTANCE = "LimitExceeded.SandboxInstance"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_SANDBOXTOOL = "ResourceNotFound.SandboxTool"
//  RESOURCENOTFOUND_STORAGEMOUNT = "ResourceNotFound.StorageMount"
//  RESOURCEUNAVAILABLE_SANDBOXTOOL = "ResourceUnavailable.SandboxTool"
func (c *Client) StartSandboxInstanceWithContext(ctx context.Context, request *StartSandboxInstanceRequest) (response *StartSandboxInstanceResponse, err error) {
    if request == nil {
        request = NewStartSandboxInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "StartSandboxInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartSandboxInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartSandboxInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewStopSandboxInstanceRequest() (request *StopSandboxInstanceRequest) {
    request = &StopSandboxInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "StopSandboxInstance")
    
    
    return
}

func NewStopSandboxInstanceResponse() (response *StopSandboxInstanceResponse) {
    response = &StopSandboxInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopSandboxInstance
// 停止沙箱实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_SANDBOXINSTANCE = "ResourceNotFound.SandboxInstance"
func (c *Client) StopSandboxInstance(request *StopSandboxInstanceRequest) (response *StopSandboxInstanceResponse, err error) {
    return c.StopSandboxInstanceWithContext(context.Background(), request)
}

// StopSandboxInstance
// 停止沙箱实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_SANDBOXINSTANCE = "ResourceNotFound.SandboxInstance"
func (c *Client) StopSandboxInstanceWithContext(ctx context.Context, request *StopSandboxInstanceRequest) (response *StopSandboxInstanceResponse, err error) {
    if request == nil {
        request = NewStopSandboxInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "StopSandboxInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopSandboxInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopSandboxInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateSandboxInstanceRequest() (request *UpdateSandboxInstanceRequest) {
    request = &UpdateSandboxInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "UpdateSandboxInstance")
    
    
    return
}

func NewUpdateSandboxInstanceResponse() (response *UpdateSandboxInstanceResponse) {
    response = &UpdateSandboxInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateSandboxInstance
// 更新沙箱实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_TIMEOUT = "InvalidParameterValue.Timeout"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_SANDBOXINSTANCE = "ResourceNotFound.SandboxInstance"
//  UNSUPPORTEDOPERATION_SANDBOXINSTANCE = "UnsupportedOperation.SandboxInstance"
func (c *Client) UpdateSandboxInstance(request *UpdateSandboxInstanceRequest) (response *UpdateSandboxInstanceResponse, err error) {
    return c.UpdateSandboxInstanceWithContext(context.Background(), request)
}

// UpdateSandboxInstance
// 更新沙箱实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_TIMEOUT = "InvalidParameterValue.Timeout"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_SANDBOXINSTANCE = "ResourceNotFound.SandboxInstance"
//  UNSUPPORTEDOPERATION_SANDBOXINSTANCE = "UnsupportedOperation.SandboxInstance"
func (c *Client) UpdateSandboxInstanceWithContext(ctx context.Context, request *UpdateSandboxInstanceRequest) (response *UpdateSandboxInstanceResponse, err error) {
    if request == nil {
        request = NewUpdateSandboxInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "UpdateSandboxInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateSandboxInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateSandboxInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateSandboxToolRequest() (request *UpdateSandboxToolRequest) {
    request = &UpdateSandboxToolRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "UpdateSandboxTool")
    
    
    return
}

func NewUpdateSandboxToolResponse() (response *UpdateSandboxToolResponse) {
    response = &UpdateSandboxToolResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateSandboxTool
// 更新沙箱工具
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_SANDBOXTOOL = "ResourceNotFound.SandboxTool"
func (c *Client) UpdateSandboxTool(request *UpdateSandboxToolRequest) (response *UpdateSandboxToolResponse, err error) {
    return c.UpdateSandboxToolWithContext(context.Background(), request)
}

// UpdateSandboxTool
// 更新沙箱工具
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_SANDBOXTOOL = "ResourceNotFound.SandboxTool"
func (c *Client) UpdateSandboxToolWithContext(ctx context.Context, request *UpdateSandboxToolRequest) (response *UpdateSandboxToolResponse, err error) {
    if request == nil {
        request = NewUpdateSandboxToolRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "UpdateSandboxTool")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateSandboxTool require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateSandboxToolResponse()
    err = c.Send(request, response)
    return
}
