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


func NewAcquireDeploymentTokenRequest() (request *AcquireDeploymentTokenRequest) {
    request = &AcquireDeploymentTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "AcquireDeploymentToken")
    
    
    return
}

func NewAcquireDeploymentTokenResponse() (response *AcquireDeploymentTokenResponse) {
    response = &AcquireDeploymentTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AcquireDeploymentToken
// 获取 Deployment 访问 Token
//
// 可能返回的错误码:
//  INVALIDPARAMETER_DEPLOYMENTID = "InvalidParameter.DeploymentId"
//  RESOURCENOTFOUND_DEPLOYMENT = "ResourceNotFound.Deployment"
//  RESOURCEUNAVAILABLE_DEPLOYMENT = "ResourceUnavailable.Deployment"
func (c *Client) AcquireDeploymentToken(request *AcquireDeploymentTokenRequest) (response *AcquireDeploymentTokenResponse, err error) {
    return c.AcquireDeploymentTokenWithContext(context.Background(), request)
}

// AcquireDeploymentToken
// 获取 Deployment 访问 Token
//
// 可能返回的错误码:
//  INVALIDPARAMETER_DEPLOYMENTID = "InvalidParameter.DeploymentId"
//  RESOURCENOTFOUND_DEPLOYMENT = "ResourceNotFound.Deployment"
//  RESOURCEUNAVAILABLE_DEPLOYMENT = "ResourceUnavailable.Deployment"
func (c *Client) AcquireDeploymentTokenWithContext(ctx context.Context, request *AcquireDeploymentTokenRequest) (response *AcquireDeploymentTokenResponse, err error) {
    if request == nil {
        request = NewAcquireDeploymentTokenRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "AcquireDeploymentToken")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AcquireDeploymentToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewAcquireDeploymentTokenResponse()
    err = c.Send(request, response)
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

func NewAppendEventRequest() (request *AppendEventRequest) {
    request = &AppendEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "AppendEvent")
    
    
    return
}

func NewAppendEventResponse() (response *AppendEventResponse) {
    response = &AppendEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AppendEvent
// 追加事件。
//
// 
//
// 向指定会话追加一条事件。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_EVENTIDDUPLICATE = "InvalidParameter.EventIdDuplicate"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND_SESSION = "ResourceNotFound.Session"
//  RESOURCENOTFOUND_SESSIONSPACENOTEXIST = "ResourceNotFound.SessionSpaceNotExist"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) AppendEvent(request *AppendEventRequest) (response *AppendEventResponse, err error) {
    return c.AppendEventWithContext(context.Background(), request)
}

// AppendEvent
// 追加事件。
//
// 
//
// 向指定会话追加一条事件。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_EVENTIDDUPLICATE = "InvalidParameter.EventIdDuplicate"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND_SESSION = "ResourceNotFound.Session"
//  RESOURCENOTFOUND_SESSIONSPACENOTEXIST = "ResourceNotFound.SessionSpaceNotExist"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) AppendEventWithContext(ctx context.Context, request *AppendEventRequest) (response *AppendEventResponse, err error) {
    if request == nil {
        request = NewAppendEventRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "AppendEvent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AppendEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewAppendEventResponse()
    err = c.Send(request, response)
    return
}

func NewApproveRegistryRecordRequest() (request *ApproveRegistryRecordRequest) {
    request = &ApproveRegistryRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "ApproveRegistryRecord")
    
    
    return
}

func NewApproveRegistryRecordResponse() (response *ApproveRegistryRecordResponse) {
    response = &ApproveRegistryRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ApproveRegistryRecord
// 通过 Version 审批：PENDING_APPROVAL → APPROVED。Comment 必填。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_RECORDID = "InvalidParameter.RecordId"
//  INVALIDPARAMETER_REGISTRYID = "InvalidParameter.RegistryId"
//  INVALIDPARAMETER_VERSIONID = "InvalidParameter.VersionId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_REGISTRYRECORD = "ResourceNotFound.RegistryRecord"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ApproveRegistryRecord(request *ApproveRegistryRecordRequest) (response *ApproveRegistryRecordResponse, err error) {
    return c.ApproveRegistryRecordWithContext(context.Background(), request)
}

// ApproveRegistryRecord
// 通过 Version 审批：PENDING_APPROVAL → APPROVED。Comment 必填。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_RECORDID = "InvalidParameter.RecordId"
//  INVALIDPARAMETER_REGISTRYID = "InvalidParameter.RegistryId"
//  INVALIDPARAMETER_VERSIONID = "InvalidParameter.VersionId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_REGISTRYRECORD = "ResourceNotFound.RegistryRecord"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ApproveRegistryRecordWithContext(ctx context.Context, request *ApproveRegistryRecordRequest) (response *ApproveRegistryRecordResponse, err error) {
    if request == nil {
        request = NewApproveRegistryRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "ApproveRegistryRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApproveRegistryRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewApproveRegistryRecordResponse()
    err = c.Send(request, response)
    return
}

func NewCancelRegistryRecordRequest() (request *CancelRegistryRecordRequest) {
    request = &CancelRegistryRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "CancelRegistryRecord")
    
    
    return
}

func NewCancelRegistryRecordResponse() (response *CancelRegistryRecordResponse) {
    response = &CancelRegistryRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CancelRegistryRecord
// PREPARING/PENDING_APPROVAL → CANCELED。Comment 必填。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_RECORDID = "InvalidParameter.RecordId"
//  INVALIDPARAMETER_REGISTRYID = "InvalidParameter.RegistryId"
//  INVALIDPARAMETER_VERSIONID = "InvalidParameter.VersionId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_REGISTRYRECORD = "ResourceNotFound.RegistryRecord"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CancelRegistryRecord(request *CancelRegistryRecordRequest) (response *CancelRegistryRecordResponse, err error) {
    return c.CancelRegistryRecordWithContext(context.Background(), request)
}

// CancelRegistryRecord
// PREPARING/PENDING_APPROVAL → CANCELED。Comment 必填。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_RECORDID = "InvalidParameter.RecordId"
//  INVALIDPARAMETER_REGISTRYID = "InvalidParameter.RegistryId"
//  INVALIDPARAMETER_VERSIONID = "InvalidParameter.VersionId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_REGISTRYRECORD = "ResourceNotFound.RegistryRecord"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CancelRegistryRecordWithContext(ctx context.Context, request *CancelRegistryRecordRequest) (response *CancelRegistryRecordResponse, err error) {
    if request == nil {
        request = NewCancelRegistryRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "CancelRegistryRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelRegistryRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelRegistryRecordResponse()
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

func NewCreateDeploymentRequest() (request *CreateDeploymentRequest) {
    request = &CreateDeploymentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "CreateDeployment")
    
    
    return
}

func NewCreateDeploymentResponse() (response *CreateDeploymentResponse) {
    response = &CreateDeploymentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDeployment
// 创建 Deployment
//
// 可能返回的错误码:
//  INVALIDPARAMETER_AFFINITYCONFIGURATION = "InvalidParameter.AffinityConfiguration"
//  INVALIDPARAMETER_DEPLOYMENTNAME = "InvalidParameter.DeploymentName"
//  INVALIDPARAMETER_LIFECYCLECONFIGURATION = "InvalidParameter.LifecycleConfiguration"
//  INVALIDPARAMETER_SCALINGCONFIGURATION = "InvalidParameter.ScalingConfiguration"
//  INVALIDPARAMETER_TOOLID = "InvalidParameter.ToolId"
//  LIMITEXCEEDED_DEPLOYMENT = "LimitExceeded.Deployment"
//  RESOURCEINUSE_DEPLOYMENTNAME = "ResourceInUse.DeploymentName"
//  RESOURCENOTFOUND_SANDBOXTOOL = "ResourceNotFound.SandboxTool"
func (c *Client) CreateDeployment(request *CreateDeploymentRequest) (response *CreateDeploymentResponse, err error) {
    return c.CreateDeploymentWithContext(context.Background(), request)
}

// CreateDeployment
// 创建 Deployment
//
// 可能返回的错误码:
//  INVALIDPARAMETER_AFFINITYCONFIGURATION = "InvalidParameter.AffinityConfiguration"
//  INVALIDPARAMETER_DEPLOYMENTNAME = "InvalidParameter.DeploymentName"
//  INVALIDPARAMETER_LIFECYCLECONFIGURATION = "InvalidParameter.LifecycleConfiguration"
//  INVALIDPARAMETER_SCALINGCONFIGURATION = "InvalidParameter.ScalingConfiguration"
//  INVALIDPARAMETER_TOOLID = "InvalidParameter.ToolId"
//  LIMITEXCEEDED_DEPLOYMENT = "LimitExceeded.Deployment"
//  RESOURCEINUSE_DEPLOYMENTNAME = "ResourceInUse.DeploymentName"
//  RESOURCENOTFOUND_SANDBOXTOOL = "ResourceNotFound.SandboxTool"
func (c *Client) CreateDeploymentWithContext(ctx context.Context, request *CreateDeploymentRequest) (response *CreateDeploymentResponse, err error) {
    if request == nil {
        request = NewCreateDeploymentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "CreateDeployment")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDeployment require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDeploymentResponse()
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
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreatePreCacheImageTask(request *CreatePreCacheImageTaskRequest) (response *CreatePreCacheImageTaskResponse, err error) {
    return c.CreatePreCacheImageTaskWithContext(context.Background(), request)
}

// CreatePreCacheImageTask
// 创建镜像预热任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
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

func NewCreateRegistryRequest() (request *CreateRegistryRequest) {
    request = &CreateRegistryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "CreateRegistry")
    
    
    return
}

func NewCreateRegistryResponse() (response *CreateRegistryResponse) {
    response = &CreateRegistryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRegistry
// 创建 Agent Registry（注册中心）。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_DESCRIPTION = "InvalidParameter.Description"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_TAGS = "InvalidParameter.Tags"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_APPROVALMODE = "InvalidParameterValue.ApprovalMode"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateRegistry(request *CreateRegistryRequest) (response *CreateRegistryResponse, err error) {
    return c.CreateRegistryWithContext(context.Background(), request)
}

// CreateRegistry
// 创建 Agent Registry（注册中心）。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_DESCRIPTION = "InvalidParameter.Description"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_TAGS = "InvalidParameter.Tags"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_APPROVALMODE = "InvalidParameterValue.ApprovalMode"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateRegistryWithContext(ctx context.Context, request *CreateRegistryRequest) (response *CreateRegistryResponse, err error) {
    if request == nil {
        request = NewCreateRegistryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "CreateRegistry")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRegistry require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRegistryResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRegistryRecordRequest() (request *CreateRegistryRecordRequest) {
    request = &CreateRegistryRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "CreateRegistryRecord")
    
    
    return
}

func NewCreateRegistryRecordResponse() (response *CreateRegistryRecordResponse) {
    response = &CreateRegistryRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRegistryRecord
// 统一创建 Registry Record（含 revision 1）。请求通过 DescriptorType 与严格内容输入 Union 选择底层类型：MCPSource / AgentSource / SkillSource / CustomDescriptors 四选一，必须与 DescriptorType 对应。不接受 RecordId 或 ChangeLog；同名 Record 返回冲突，不隐式追加 Version。追加 Version 请使用 UpdateRegistryRecord。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_AGENTSOURCE = "InvalidParameter.AgentSource"
//  INVALIDPARAMETER_CUSTOMDESCRIPTORS = "InvalidParameter.CustomDescriptors"
//  INVALIDPARAMETER_DESCRIPTION = "InvalidParameter.Description"
//  INVALIDPARAMETER_MCPSOURCE = "InvalidParameter.MCPSource"
//  INVALIDPARAMETER_RECORDSOURCE = "InvalidParameter.RecordSource"
//  INVALIDPARAMETER_REGISTRYID = "InvalidParameter.RegistryId"
//  INVALIDPARAMETER_SKILLSOURCE = "InvalidParameter.SkillSource"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_AGENTSOURCETYPE = "InvalidParameterValue.AgentSourceType"
//  INVALIDPARAMETERVALUE_DESCRIPTORTYPE = "InvalidParameterValue.DescriptorType"
//  INVALIDPARAMETERVALUE_MCPSOURCETYPE = "InvalidParameterValue.MCPSourceType"
//  INVALIDPARAMETERVALUE_SKILLSOURCETYPE = "InvalidParameterValue.SkillSourceType"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_AGENTSOURCEENDPOINTURL = "MissingParameter.AgentSourceEndpointURL"
//  MISSINGPARAMETER_MCPSOURCEDESCRIPTORS = "MissingParameter.MCPSourceDescriptors"
//  MISSINGPARAMETER_RECORDSOURCE = "MissingParameter.RecordSource"
//  MISSINGPARAMETER_SKILLSOURCESKILLMD = "MissingParameter.SkillSourceSkillMd"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateRegistryRecord(request *CreateRegistryRecordRequest) (response *CreateRegistryRecordResponse, err error) {
    return c.CreateRegistryRecordWithContext(context.Background(), request)
}

// CreateRegistryRecord
// 统一创建 Registry Record（含 revision 1）。请求通过 DescriptorType 与严格内容输入 Union 选择底层类型：MCPSource / AgentSource / SkillSource / CustomDescriptors 四选一，必须与 DescriptorType 对应。不接受 RecordId 或 ChangeLog；同名 Record 返回冲突，不隐式追加 Version。追加 Version 请使用 UpdateRegistryRecord。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_AGENTSOURCE = "InvalidParameter.AgentSource"
//  INVALIDPARAMETER_CUSTOMDESCRIPTORS = "InvalidParameter.CustomDescriptors"
//  INVALIDPARAMETER_DESCRIPTION = "InvalidParameter.Description"
//  INVALIDPARAMETER_MCPSOURCE = "InvalidParameter.MCPSource"
//  INVALIDPARAMETER_RECORDSOURCE = "InvalidParameter.RecordSource"
//  INVALIDPARAMETER_REGISTRYID = "InvalidParameter.RegistryId"
//  INVALIDPARAMETER_SKILLSOURCE = "InvalidParameter.SkillSource"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_AGENTSOURCETYPE = "InvalidParameterValue.AgentSourceType"
//  INVALIDPARAMETERVALUE_DESCRIPTORTYPE = "InvalidParameterValue.DescriptorType"
//  INVALIDPARAMETERVALUE_MCPSOURCETYPE = "InvalidParameterValue.MCPSourceType"
//  INVALIDPARAMETERVALUE_SKILLSOURCETYPE = "InvalidParameterValue.SkillSourceType"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_AGENTSOURCEENDPOINTURL = "MissingParameter.AgentSourceEndpointURL"
//  MISSINGPARAMETER_MCPSOURCEDESCRIPTORS = "MissingParameter.MCPSourceDescriptors"
//  MISSINGPARAMETER_RECORDSOURCE = "MissingParameter.RecordSource"
//  MISSINGPARAMETER_SKILLSOURCESKILLMD = "MissingParameter.SkillSourceSkillMd"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateRegistryRecordWithContext(ctx context.Context, request *CreateRegistryRecordRequest) (response *CreateRegistryRecordResponse, err error) {
    if request == nil {
        request = NewCreateRegistryRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "CreateRegistryRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRegistryRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRegistryRecordResponse()
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
//  FAILEDOPERATION_DEPENDENCYUNAVAILABLE = "FailedOperation.DependencyUnavailable"
//  FAILEDOPERATION_DUPLICATEREQUEST = "FailedOperation.DuplicateRequest"
//  FAILEDOPERATION_REQUESTINPROGRESS = "FailedOperation.RequestInProgress"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_NETWORKSETUPFAILED = "InternalError.NetworkSetupFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INTERNALERROR_VPCSERVICEUNAVAILABLE = "InternalError.VPCServiceUnavailable"
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
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SECURITYGROUP = "ResourceNotFound.SecurityGroup"
//  RESOURCENOTFOUND_STORAGEMOUNT = "ResourceNotFound.StorageMount"
//  RESOURCENOTFOUND_SUBNET = "ResourceNotFound.Subnet"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSandboxTool(request *CreateSandboxToolRequest) (response *CreateSandboxToolResponse, err error) {
    return c.CreateSandboxToolWithContext(context.Background(), request)
}

// CreateSandboxTool
// 创建沙箱工具
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DEPENDENCYUNAVAILABLE = "FailedOperation.DependencyUnavailable"
//  FAILEDOPERATION_DUPLICATEREQUEST = "FailedOperation.DuplicateRequest"
//  FAILEDOPERATION_REQUESTINPROGRESS = "FailedOperation.RequestInProgress"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_NETWORKSETUPFAILED = "InternalError.NetworkSetupFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INTERNALERROR_VPCSERVICEUNAVAILABLE = "InternalError.VPCServiceUnavailable"
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
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SECURITYGROUP = "ResourceNotFound.SecurityGroup"
//  RESOURCENOTFOUND_STORAGEMOUNT = "ResourceNotFound.StorageMount"
//  RESOURCENOTFOUND_SUBNET = "ResourceNotFound.Subnet"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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

func NewCreateSessionRequest() (request *CreateSessionRequest) {
    request = &CreateSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "CreateSession")
    
    
    return
}

func NewCreateSessionResponse() (response *CreateSessionResponse) {
    response = &CreateSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSession
// 创建会话。
//
// 
//
// 为指定 Agent 和用户创建会话，创建成功后返回会话信息。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_SESSIONIDDUPLICATE = "InvalidParameter.SessionIdDuplicate"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND_SESSIONSPACENOTEXIST = "ResourceNotFound.SessionSpaceNotExist"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateSession(request *CreateSessionRequest) (response *CreateSessionResponse, err error) {
    return c.CreateSessionWithContext(context.Background(), request)
}

// CreateSession
// 创建会话。
//
// 
//
// 为指定 Agent 和用户创建会话，创建成功后返回会话信息。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_SESSIONIDDUPLICATE = "InvalidParameter.SessionIdDuplicate"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND_SESSIONSPACENOTEXIST = "ResourceNotFound.SessionSpaceNotExist"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateSessionWithContext(ctx context.Context, request *CreateSessionRequest) (response *CreateSessionResponse, err error) {
    if request == nil {
        request = NewCreateSessionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "CreateSession")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSession require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSessionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSessionSpaceRequest() (request *CreateSessionSpaceRequest) {
    request = &CreateSessionSpaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "CreateSessionSpace")
    
    
    return
}

func NewCreateSessionSpaceResponse() (response *CreateSessionSpaceResponse) {
    response = &CreateSessionSpaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSessionSpace
// 创建会话空间。
//
// 为当前应用在指定地域创建会话空间，创建成功后返回会话空间信息。会话空间用于隔离不同业务场景下的用户、会话、事件及状态数据。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) CreateSessionSpace(request *CreateSessionSpaceRequest) (response *CreateSessionSpaceResponse, err error) {
    return c.CreateSessionSpaceWithContext(context.Background(), request)
}

// CreateSessionSpace
// 创建会话空间。
//
// 为当前应用在指定地域创建会话空间，创建成功后返回会话空间信息。会话空间用于隔离不同业务场景下的用户、会话、事件及状态数据。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) CreateSessionSpaceWithContext(ctx context.Context, request *CreateSessionSpaceRequest) (response *CreateSessionSpaceResponse, err error) {
    if request == nil {
        request = NewCreateSessionSpaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "CreateSessionSpace")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSessionSpace require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSessionSpaceResponse()
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

func NewDeleteDeploymentRequest() (request *DeleteDeploymentRequest) {
    request = &DeleteDeploymentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "DeleteDeployment")
    
    
    return
}

func NewDeleteDeploymentResponse() (response *DeleteDeploymentResponse) {
    response = &DeleteDeploymentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDeployment
// 删除 Deployment
//
// 可能返回的错误码:
//  INVALIDPARAMETER_DEPLOYMENTID = "InvalidParameter.DeploymentId"
//  RESOURCEINUSE_DEPLOYMENT = "ResourceInUse.Deployment"
//  RESOURCENOTFOUND_DEPLOYMENT = "ResourceNotFound.Deployment"
func (c *Client) DeleteDeployment(request *DeleteDeploymentRequest) (response *DeleteDeploymentResponse, err error) {
    return c.DeleteDeploymentWithContext(context.Background(), request)
}

// DeleteDeployment
// 删除 Deployment
//
// 可能返回的错误码:
//  INVALIDPARAMETER_DEPLOYMENTID = "InvalidParameter.DeploymentId"
//  RESOURCEINUSE_DEPLOYMENT = "ResourceInUse.Deployment"
//  RESOURCENOTFOUND_DEPLOYMENT = "ResourceNotFound.Deployment"
func (c *Client) DeleteDeploymentWithContext(ctx context.Context, request *DeleteDeploymentRequest) (response *DeleteDeploymentResponse, err error) {
    if request == nil {
        request = NewDeleteDeploymentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "DeleteDeployment")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDeployment require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDeploymentResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRegistryRequest() (request *DeleteRegistryRequest) {
    request = &DeleteRegistryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "DeleteRegistry")
    
    
    return
}

func NewDeleteRegistryResponse() (response *DeleteRegistryResponse) {
    response = &DeleteRegistryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRegistry
// 删除 Registry。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_REGISTRYID = "InvalidParameter.RegistryId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRegistry(request *DeleteRegistryRequest) (response *DeleteRegistryResponse, err error) {
    return c.DeleteRegistryWithContext(context.Background(), request)
}

// DeleteRegistry
// 删除 Registry。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_REGISTRYID = "InvalidParameter.RegistryId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRegistryWithContext(ctx context.Context, request *DeleteRegistryRequest) (response *DeleteRegistryResponse, err error) {
    if request == nil {
        request = NewDeleteRegistryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "DeleteRegistry")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRegistry require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRegistryResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRegistryRecordRequest() (request *DeleteRegistryRecordRequest) {
    request = &DeleteRegistryRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "DeleteRegistryRecord")
    
    
    return
}

func NewDeleteRegistryRecordResponse() (response *DeleteRegistryRecordResponse) {
    response = &DeleteRegistryRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRegistryRecord
// 删除 Registry Record 或指定 Version。省略 VersionId 时对整个 Record 进行软删除；传入 VersionId 时只删除指定 Version（Stable 指向的 Version 不允许删除；仅剩一个 Approved Version 时不允许删除）。取代原 DeleteRegistryRecordVersion。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_RECORDID = "InvalidParameter.RecordId"
//  INVALIDPARAMETER_REGISTRYID = "InvalidParameter.RegistryId"
//  INVALIDPARAMETER_VERSIONID = "InvalidParameter.VersionId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_REASON = "MissingParameter.Reason"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_REGISTRYRECORD = "ResourceNotFound.RegistryRecord"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRegistryRecord(request *DeleteRegistryRecordRequest) (response *DeleteRegistryRecordResponse, err error) {
    return c.DeleteRegistryRecordWithContext(context.Background(), request)
}

// DeleteRegistryRecord
// 删除 Registry Record 或指定 Version。省略 VersionId 时对整个 Record 进行软删除；传入 VersionId 时只删除指定 Version（Stable 指向的 Version 不允许删除；仅剩一个 Approved Version 时不允许删除）。取代原 DeleteRegistryRecordVersion。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_RECORDID = "InvalidParameter.RecordId"
//  INVALIDPARAMETER_REGISTRYID = "InvalidParameter.RegistryId"
//  INVALIDPARAMETER_VERSIONID = "InvalidParameter.VersionId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_REASON = "MissingParameter.Reason"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_REGISTRYRECORD = "ResourceNotFound.RegistryRecord"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRegistryRecordWithContext(ctx context.Context, request *DeleteRegistryRecordRequest) (response *DeleteRegistryRecordResponse, err error) {
    if request == nil {
        request = NewDeleteRegistryRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "DeleteRegistryRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRegistryRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRegistryRecordResponse()
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

func NewDeleteSessionRequest() (request *DeleteSessionRequest) {
    request = &DeleteSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "DeleteSession")
    
    
    return
}

func NewDeleteSessionResponse() (response *DeleteSessionResponse) {
    response = &DeleteSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSession
// 删除会话
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_SESSIONNOTEXIST = "ResourceNotFound.SessionNotExist"
//  RESOURCENOTFOUND_SESSIONSPACENOTEXIST = "ResourceNotFound.SessionSpaceNotExist"
func (c *Client) DeleteSession(request *DeleteSessionRequest) (response *DeleteSessionResponse, err error) {
    return c.DeleteSessionWithContext(context.Background(), request)
}

// DeleteSession
// 删除会话
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_SESSIONNOTEXIST = "ResourceNotFound.SessionNotExist"
//  RESOURCENOTFOUND_SESSIONSPACENOTEXIST = "ResourceNotFound.SessionSpaceNotExist"
func (c *Client) DeleteSessionWithContext(ctx context.Context, request *DeleteSessionRequest) (response *DeleteSessionResponse, err error) {
    if request == nil {
        request = NewDeleteSessionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "DeleteSession")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSession require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSessionResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSessionSpaceRequest() (request *DeleteSessionSpaceRequest) {
    request = &DeleteSessionSpaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "DeleteSessionSpace")
    
    
    return
}

func NewDeleteSessionSpaceResponse() (response *DeleteSessionSpaceResponse) {
    response = &DeleteSessionSpaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSessionSpace
// 删除会话空间。
//
// 删除指定的会话空间。仅允许删除不包含会话、事件或用户状态数据的非默认会话空间；系统默认会话空间不能删除。删除成功后不再返回会话空间信息。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_DEFAULTSESSIONSPACEPROTECTED = "InvalidParameter.DefaultSessionSpaceProtected"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCEINUSE_SESSIONSPACENOTEMPTY = "ResourceInUse.SessionSpaceNotEmpty"
//  RESOURCENOTFOUND_SESSIONSPACENOTEXIST = "ResourceNotFound.SessionSpaceNotExist"
func (c *Client) DeleteSessionSpace(request *DeleteSessionSpaceRequest) (response *DeleteSessionSpaceResponse, err error) {
    return c.DeleteSessionSpaceWithContext(context.Background(), request)
}

// DeleteSessionSpace
// 删除会话空间。
//
// 删除指定的会话空间。仅允许删除不包含会话、事件或用户状态数据的非默认会话空间；系统默认会话空间不能删除。删除成功后不再返回会话空间信息。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_DEFAULTSESSIONSPACEPROTECTED = "InvalidParameter.DefaultSessionSpaceProtected"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCEINUSE_SESSIONSPACENOTEMPTY = "ResourceInUse.SessionSpaceNotEmpty"
//  RESOURCENOTFOUND_SESSIONSPACENOTEXIST = "ResourceNotFound.SessionSpaceNotExist"
func (c *Client) DeleteSessionSpaceWithContext(ctx context.Context, request *DeleteSessionSpaceRequest) (response *DeleteSessionSpaceResponse, err error) {
    if request == nil {
        request = NewDeleteSessionSpaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "DeleteSessionSpace")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSessionSpace require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSessionSpaceResponse()
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

func NewDescribeDeploymentRequest() (request *DescribeDeploymentRequest) {
    request = &DescribeDeploymentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "DescribeDeployment")
    
    
    return
}

func NewDescribeDeploymentResponse() (response *DescribeDeploymentResponse) {
    response = &DescribeDeploymentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeployment
// 查询 Deployment 信息
//
// 可能返回的错误码:
//  INVALIDPARAMETER_DEPLOYMENTID = "InvalidParameter.DeploymentId"
//  RESOURCENOTFOUND_DEPLOYMENT = "ResourceNotFound.Deployment"
func (c *Client) DescribeDeployment(request *DescribeDeploymentRequest) (response *DescribeDeploymentResponse, err error) {
    return c.DescribeDeploymentWithContext(context.Background(), request)
}

// DescribeDeployment
// 查询 Deployment 信息
//
// 可能返回的错误码:
//  INVALIDPARAMETER_DEPLOYMENTID = "InvalidParameter.DeploymentId"
//  RESOURCENOTFOUND_DEPLOYMENT = "ResourceNotFound.Deployment"
func (c *Client) DescribeDeploymentWithContext(ctx context.Context, request *DescribeDeploymentRequest) (response *DescribeDeploymentResponse, err error) {
    if request == nil {
        request = NewDescribeDeploymentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "DescribeDeployment")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeployment require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeploymentResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeploymentListRequest() (request *DescribeDeploymentListRequest) {
    request = &DescribeDeploymentListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "DescribeDeploymentList")
    
    
    return
}

func NewDescribeDeploymentListResponse() (response *DescribeDeploymentListResponse) {
    response = &DescribeDeploymentListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDeploymentList
// 查询 Deployment 列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FILTERS = "InvalidParameter.Filters"
//  INVALIDPARAMETER_LIMIT = "InvalidParameter.Limit"
//  INVALIDPARAMETER_OFFSET = "InvalidParameter.Offset"
func (c *Client) DescribeDeploymentList(request *DescribeDeploymentListRequest) (response *DescribeDeploymentListResponse, err error) {
    return c.DescribeDeploymentListWithContext(context.Background(), request)
}

// DescribeDeploymentList
// 查询 Deployment 列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER_FILTERS = "InvalidParameter.Filters"
//  INVALIDPARAMETER_LIMIT = "InvalidParameter.Limit"
//  INVALIDPARAMETER_OFFSET = "InvalidParameter.Offset"
func (c *Client) DescribeDeploymentListWithContext(ctx context.Context, request *DescribeDeploymentListRequest) (response *DescribeDeploymentListResponse, err error) {
    if request == nil {
        request = NewDescribeDeploymentListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "DescribeDeploymentList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeploymentList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeploymentListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEventsRequest() (request *DescribeEventsRequest) {
    request = &DescribeEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "DescribeEvents")
    
    
    return
}

func NewDescribeEventsResponse() (response *DescribeEventsResponse) {
    response = &DescribeEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEvents
// 查询事件列表。
//
// 
//
// 查询指定会话的事件流，支持按作者和起始时间筛选。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND_SESSIONNOTEXIST = "ResourceNotFound.SessionNotExist"
//  RESOURCENOTFOUND_SESSIONSPACENOTEXIST = "ResourceNotFound.SessionSpaceNotExist"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeEvents(request *DescribeEventsRequest) (response *DescribeEventsResponse, err error) {
    return c.DescribeEventsWithContext(context.Background(), request)
}

// DescribeEvents
// 查询事件列表。
//
// 
//
// 查询指定会话的事件流，支持按作者和起始时间筛选。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND_SESSIONNOTEXIST = "ResourceNotFound.SessionNotExist"
//  RESOURCENOTFOUND_SESSIONSPACENOTEXIST = "ResourceNotFound.SessionSpaceNotExist"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeEventsWithContext(ctx context.Context, request *DescribeEventsRequest) (response *DescribeEventsResponse, err error) {
    if request == nil {
        request = NewDescribeEventsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "DescribeEvents")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEventsResponse()
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
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribePreCacheImageTask(request *DescribePreCacheImageTaskRequest) (response *DescribePreCacheImageTaskResponse, err error) {
    return c.DescribePreCacheImageTaskWithContext(context.Background(), request)
}

// DescribePreCacheImageTask
// 查询镜像预热任务信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
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

func NewDescribeQuotaOverviewRequest() (request *DescribeQuotaOverviewRequest) {
    request = &DescribeQuotaOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "DescribeQuotaOverview")
    
    
    return
}

func NewDescribeQuotaOverviewResponse() (response *DescribeQuotaOverviewResponse) {
    response = &DescribeQuotaOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeQuotaOverview
// 查询当前调用账号的资源配额和当前总用量，以及账号下各配额组的资源配额和当前用量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeQuotaOverview(request *DescribeQuotaOverviewRequest) (response *DescribeQuotaOverviewResponse, err error) {
    return c.DescribeQuotaOverviewWithContext(context.Background(), request)
}

// DescribeQuotaOverview
// 查询当前调用账号的资源配额和当前总用量，以及账号下各配额组的资源配额和当前用量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeQuotaOverviewWithContext(ctx context.Context, request *DescribeQuotaOverviewRequest) (response *DescribeQuotaOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeQuotaOverviewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "DescribeQuotaOverview")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeQuotaOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeQuotaOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegistryRequest() (request *DescribeRegistryRequest) {
    request = &DescribeRegistryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "DescribeRegistry")
    
    
    return
}

func NewDescribeRegistryResponse() (response *DescribeRegistryResponse) {
    response = &DescribeRegistryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRegistry
// 按 RegistryId 查询 Registry 详情。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_REGISTRYID = "InvalidParameter.RegistryId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRegistry(request *DescribeRegistryRequest) (response *DescribeRegistryResponse, err error) {
    return c.DescribeRegistryWithContext(context.Background(), request)
}

// DescribeRegistry
// 按 RegistryId 查询 Registry 详情。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_REGISTRYID = "InvalidParameter.RegistryId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRegistryWithContext(ctx context.Context, request *DescribeRegistryRequest) (response *DescribeRegistryResponse, err error) {
    if request == nil {
        request = NewDescribeRegistryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "DescribeRegistry")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRegistry require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRegistryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegistryAuditLogListRequest() (request *DescribeRegistryAuditLogListRequest) {
    request = &DescribeRegistryAuditLogListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "DescribeRegistryAuditLogList")
    
    
    return
}

func NewDescribeRegistryAuditLogListResponse() (response *DescribeRegistryAuditLogListResponse) {
    response = &DescribeRegistryAuditLogListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRegistryAuditLogList
// 分页查询指定Registry / Record / Version的审计日志。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_LIMIT = "InvalidParameter.Limit"
//  INVALIDPARAMETER_OFFSET = "InvalidParameter.Offset"
//  INVALIDPARAMETER_RECORDID = "InvalidParameter.RecordId"
//  INVALIDPARAMETER_REGISTRYID = "InvalidParameter.RegistryId"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  INVALIDPARAMETER_VERSIONID = "InvalidParameter.VersionId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRegistryAuditLogList(request *DescribeRegistryAuditLogListRequest) (response *DescribeRegistryAuditLogListResponse, err error) {
    return c.DescribeRegistryAuditLogListWithContext(context.Background(), request)
}

// DescribeRegistryAuditLogList
// 分页查询指定Registry / Record / Version的审计日志。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"
//  INVALIDPARAMETER_LIMIT = "InvalidParameter.Limit"
//  INVALIDPARAMETER_OFFSET = "InvalidParameter.Offset"
//  INVALIDPARAMETER_RECORDID = "InvalidParameter.RecordId"
//  INVALIDPARAMETER_REGISTRYID = "InvalidParameter.RegistryId"
//  INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"
//  INVALIDPARAMETER_VERSIONID = "InvalidParameter.VersionId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRegistryAuditLogListWithContext(ctx context.Context, request *DescribeRegistryAuditLogListRequest) (response *DescribeRegistryAuditLogListResponse, err error) {
    if request == nil {
        request = NewDescribeRegistryAuditLogListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "DescribeRegistryAuditLogList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRegistryAuditLogList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRegistryAuditLogListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegistryListRequest() (request *DescribeRegistryListRequest) {
    request = &DescribeRegistryListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "DescribeRegistryList")
    
    
    return
}

func NewDescribeRegistryListResponse() (response *DescribeRegistryListResponse) {
    response = &DescribeRegistryListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRegistryList
// 分页查询当前租户可见的 Registry 列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_FILTERS = "InvalidParameter.Filters"
//  INVALIDPARAMETER_LIMIT = "InvalidParameter.Limit"
//  INVALIDPARAMETER_OFFSET = "InvalidParameter.Offset"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRegistryList(request *DescribeRegistryListRequest) (response *DescribeRegistryListResponse, err error) {
    return c.DescribeRegistryListWithContext(context.Background(), request)
}

// DescribeRegistryList
// 分页查询当前租户可见的 Registry 列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_FILTERS = "InvalidParameter.Filters"
//  INVALIDPARAMETER_LIMIT = "InvalidParameter.Limit"
//  INVALIDPARAMETER_OFFSET = "InvalidParameter.Offset"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRegistryListWithContext(ctx context.Context, request *DescribeRegistryListRequest) (response *DescribeRegistryListResponse, err error) {
    if request == nil {
        request = NewDescribeRegistryListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "DescribeRegistryList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRegistryList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRegistryListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegistryRecordRequest() (request *DescribeRegistryRecordRequest) {
    request = &DescribeRegistryRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "DescribeRegistryRecord")
    
    
    return
}

func NewDescribeRegistryRecordResponse() (response *DescribeRegistryRecordResponse) {
    response = &DescribeRegistryRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRegistryRecord
// 查询 Record 详情和其中一个 Version。请求可通过互斥的 VersionId 或 Label 选择 Version；均省略时默认 Label=stable。取代原 DescribeRegistryRecordVersion。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_RECORDID = "InvalidParameter.RecordId"
//  INVALIDPARAMETER_RECORDSELECTOR = "InvalidParameter.RecordSelector"
//  INVALIDPARAMETER_REGISTRYID = "InvalidParameter.RegistryId"
//  INVALIDPARAMETER_VERSIONID = "InvalidParameter.VersionId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_REGISTRYRECORD = "ResourceNotFound.RegistryRecord"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRegistryRecord(request *DescribeRegistryRecordRequest) (response *DescribeRegistryRecordResponse, err error) {
    return c.DescribeRegistryRecordWithContext(context.Background(), request)
}

// DescribeRegistryRecord
// 查询 Record 详情和其中一个 Version。请求可通过互斥的 VersionId 或 Label 选择 Version；均省略时默认 Label=stable。取代原 DescribeRegistryRecordVersion。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_RECORDID = "InvalidParameter.RecordId"
//  INVALIDPARAMETER_RECORDSELECTOR = "InvalidParameter.RecordSelector"
//  INVALIDPARAMETER_REGISTRYID = "InvalidParameter.RegistryId"
//  INVALIDPARAMETER_VERSIONID = "InvalidParameter.VersionId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_REGISTRYRECORD = "ResourceNotFound.RegistryRecord"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRegistryRecordWithContext(ctx context.Context, request *DescribeRegistryRecordRequest) (response *DescribeRegistryRecordResponse, err error) {
    if request == nil {
        request = NewDescribeRegistryRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "DescribeRegistryRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRegistryRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRegistryRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegistryRecordListRequest() (request *DescribeRegistryRecordListRequest) {
    request = &DescribeRegistryRecordListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "DescribeRegistryRecordList")
    
    
    return
}

func NewDescribeRegistryRecordListResponse() (response *DescribeRegistryRecordListResponse) {
    response = &DescribeRegistryRecordListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRegistryRecordList
// 分页查询 Registry 下的 Record 列表。list 类接口不接入 CAM 转发鉴权；业务侧按 CAM 二次过滤。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_FILTERS = "InvalidParameter.Filters"
//  INVALIDPARAMETER_LIMIT = "InvalidParameter.Limit"
//  INVALIDPARAMETER_OFFSET = "InvalidParameter.Offset"
//  INVALIDPARAMETER_REGISTRYID = "InvalidParameter.RegistryId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRegistryRecordList(request *DescribeRegistryRecordListRequest) (response *DescribeRegistryRecordListResponse, err error) {
    return c.DescribeRegistryRecordListWithContext(context.Background(), request)
}

// DescribeRegistryRecordList
// 分页查询 Registry 下的 Record 列表。list 类接口不接入 CAM 转发鉴权；业务侧按 CAM 二次过滤。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_FILTERS = "InvalidParameter.Filters"
//  INVALIDPARAMETER_LIMIT = "InvalidParameter.Limit"
//  INVALIDPARAMETER_OFFSET = "InvalidParameter.Offset"
//  INVALIDPARAMETER_REGISTRYID = "InvalidParameter.RegistryId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRegistryRecordListWithContext(ctx context.Context, request *DescribeRegistryRecordListRequest) (response *DescribeRegistryRecordListResponse, err error) {
    if request == nil {
        request = NewDescribeRegistryRecordListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "DescribeRegistryRecordList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRegistryRecordList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRegistryRecordListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegistryRecordVersionListRequest() (request *DescribeRegistryRecordVersionListRequest) {
    request = &DescribeRegistryRecordVersionListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "DescribeRegistryRecordVersionList")
    
    
    return
}

func NewDescribeRegistryRecordVersionListResponse() (response *DescribeRegistryRecordVersionListResponse) {
    response = &DescribeRegistryRecordVersionListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRegistryRecordVersionList
// 分页查询 Record 的 Version 列表。list 类接口不接入 CAM 转发鉴权。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_FILTERS = "InvalidParameter.Filters"
//  INVALIDPARAMETER_LIMIT = "InvalidParameter.Limit"
//  INVALIDPARAMETER_OFFSET = "InvalidParameter.Offset"
//  INVALIDPARAMETER_RECORDID = "InvalidParameter.RecordId"
//  INVALIDPARAMETER_REGISTRYID = "InvalidParameter.RegistryId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_REGISTRYRECORD = "ResourceNotFound.RegistryRecord"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRegistryRecordVersionList(request *DescribeRegistryRecordVersionListRequest) (response *DescribeRegistryRecordVersionListResponse, err error) {
    return c.DescribeRegistryRecordVersionListWithContext(context.Background(), request)
}

// DescribeRegistryRecordVersionList
// 分页查询 Record 的 Version 列表。list 类接口不接入 CAM 转发鉴权。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_FILTERS = "InvalidParameter.Filters"
//  INVALIDPARAMETER_LIMIT = "InvalidParameter.Limit"
//  INVALIDPARAMETER_OFFSET = "InvalidParameter.Offset"
//  INVALIDPARAMETER_RECORDID = "InvalidParameter.RecordId"
//  INVALIDPARAMETER_REGISTRYID = "InvalidParameter.RegistryId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_REGISTRYRECORD = "ResourceNotFound.RegistryRecord"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRegistryRecordVersionListWithContext(ctx context.Context, request *DescribeRegistryRecordVersionListRequest) (response *DescribeRegistryRecordVersionListResponse, err error) {
    if request == nil {
        request = NewDescribeRegistryRecordVersionListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "DescribeRegistryRecordVersionList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRegistryRecordVersionList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRegistryRecordVersionListResponse()
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
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETERVALUE_INSTANCEIDS = "InvalidParameterValue.InstanceIds"
//  INVALIDPARAMETERVALUE_OFFSETTOOLARGE = "InvalidParameterValue.OffsetTooLarge"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeSandboxInstanceList(request *DescribeSandboxInstanceListRequest) (response *DescribeSandboxInstanceListResponse, err error) {
    return c.DescribeSandboxInstanceListWithContext(context.Background(), request)
}

// DescribeSandboxInstanceList
// 查询沙箱实例列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETERVALUE_INSTANCEIDS = "InvalidParameterValue.InstanceIds"
//  INVALIDPARAMETERVALUE_OFFSETTOOLARGE = "InvalidParameterValue.OffsetTooLarge"
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

func NewDescribeSessionRequest() (request *DescribeSessionRequest) {
    request = &DescribeSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "DescribeSession")
    
    
    return
}

func NewDescribeSessionResponse() (response *DescribeSessionResponse) {
    response = &DescribeSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSession
// 查询会话。
//
// 
//
// 查询指定会话的信息。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND_SESSION = "ResourceNotFound.Session"
//  RESOURCENOTFOUND_SESSIONSPACENOTEXIST = "ResourceNotFound.SessionSpaceNotExist"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeSession(request *DescribeSessionRequest) (response *DescribeSessionResponse, err error) {
    return c.DescribeSessionWithContext(context.Background(), request)
}

// DescribeSession
// 查询会话。
//
// 
//
// 查询指定会话的信息。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND_SESSION = "ResourceNotFound.Session"
//  RESOURCENOTFOUND_SESSIONSPACENOTEXIST = "ResourceNotFound.SessionSpaceNotExist"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeSessionWithContext(ctx context.Context, request *DescribeSessionRequest) (response *DescribeSessionResponse, err error) {
    if request == nil {
        request = NewDescribeSessionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "DescribeSession")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSession require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSessionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSessionSpaceRequest() (request *DescribeSessionSpaceRequest) {
    request = &DescribeSessionSpaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "DescribeSessionSpace")
    
    
    return
}

func NewDescribeSessionSpaceResponse() (response *DescribeSessionSpaceResponse) {
    response = &DescribeSessionSpaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSessionSpace
// 查询会话空间详情。
//
// 查询指定会话空间的详细信息，查询成功后返回会话空间的名称、描述、状态、所属地域及创建时间等信息。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_SESSIONSPACENOTEXIST = "ResourceNotFound.SessionSpaceNotExist"
func (c *Client) DescribeSessionSpace(request *DescribeSessionSpaceRequest) (response *DescribeSessionSpaceResponse, err error) {
    return c.DescribeSessionSpaceWithContext(context.Background(), request)
}

// DescribeSessionSpace
// 查询会话空间详情。
//
// 查询指定会话空间的详细信息，查询成功后返回会话空间的名称、描述、状态、所属地域及创建时间等信息。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_SESSIONSPACENOTEXIST = "ResourceNotFound.SessionSpaceNotExist"
func (c *Client) DescribeSessionSpaceWithContext(ctx context.Context, request *DescribeSessionSpaceRequest) (response *DescribeSessionSpaceResponse, err error) {
    if request == nil {
        request = NewDescribeSessionSpaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "DescribeSessionSpace")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSessionSpace require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSessionSpaceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSessionSpacesRequest() (request *DescribeSessionSpacesRequest) {
    request = &DescribeSessionSpacesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "DescribeSessionSpaces")
    
    
    return
}

func NewDescribeSessionSpacesResponse() (response *DescribeSessionSpacesResponse) {
    response = &DescribeSessionSpacesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSessionSpaces
// 分页查询当前应用和地域下的会话空间。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) DescribeSessionSpaces(request *DescribeSessionSpacesRequest) (response *DescribeSessionSpacesResponse, err error) {
    return c.DescribeSessionSpacesWithContext(context.Background(), request)
}

// DescribeSessionSpaces
// 分页查询当前应用和地域下的会话空间。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) DescribeSessionSpacesWithContext(ctx context.Context, request *DescribeSessionSpacesRequest) (response *DescribeSessionSpacesResponse, err error) {
    if request == nil {
        request = NewDescribeSessionSpacesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "DescribeSessionSpaces")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSessionSpaces require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSessionSpacesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSessionsRequest() (request *DescribeSessionsRequest) {
    request = &DescribeSessionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "DescribeSessions")
    
    
    return
}

func NewDescribeSessionsResponse() (response *DescribeSessionsResponse) {
    response = &DescribeSessionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSessions
// 查询会话列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_SESSIONSPACENOTEXIST = "ResourceNotFound.SessionSpaceNotExist"
func (c *Client) DescribeSessions(request *DescribeSessionsRequest) (response *DescribeSessionsResponse, err error) {
    return c.DescribeSessionsWithContext(context.Background(), request)
}

// DescribeSessions
// 查询会话列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_SESSIONSPACENOTEXIST = "ResourceNotFound.SessionSpaceNotExist"
func (c *Client) DescribeSessionsWithContext(ctx context.Context, request *DescribeSessionsRequest) (response *DescribeSessionsResponse, err error) {
    if request == nil {
        request = NewDescribeSessionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "DescribeSessions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSessions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSessionsResponse()
    err = c.Send(request, response)
    return
}

func NewGetSkillPackageDownloadURLRequest() (request *GetSkillPackageDownloadURLRequest) {
    request = &GetSkillPackageDownloadURLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "GetSkillPackageDownloadURL")
    
    
    return
}

func NewGetSkillPackageDownloadURLResponse() (response *GetSkillPackageDownloadURLResponse) {
    response = &GetSkillPackageDownloadURLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetSkillPackageDownloadURL
// 获取 Skill 包下载 URL。VersionId 与 Label 互斥；均省略时使用 Stable。响应包含 ResolvedVersionId，便于调用方回填。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_RECORDID = "InvalidParameter.RecordId"
//  INVALIDPARAMETER_RECORDSELECTOR = "InvalidParameter.RecordSelector"
//  INVALIDPARAMETER_REGISTRYID = "InvalidParameter.RegistryId"
//  INVALIDPARAMETER_VERSIONID = "InvalidParameter.VersionId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_REGISTRYRECORD = "ResourceNotFound.RegistryRecord"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetSkillPackageDownloadURL(request *GetSkillPackageDownloadURLRequest) (response *GetSkillPackageDownloadURLResponse, err error) {
    return c.GetSkillPackageDownloadURLWithContext(context.Background(), request)
}

// GetSkillPackageDownloadURL
// 获取 Skill 包下载 URL。VersionId 与 Label 互斥；均省略时使用 Stable。响应包含 ResolvedVersionId，便于调用方回填。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_RECORDID = "InvalidParameter.RecordId"
//  INVALIDPARAMETER_RECORDSELECTOR = "InvalidParameter.RecordSelector"
//  INVALIDPARAMETER_REGISTRYID = "InvalidParameter.RegistryId"
//  INVALIDPARAMETER_VERSIONID = "InvalidParameter.VersionId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_REGISTRYRECORD = "ResourceNotFound.RegistryRecord"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetSkillPackageDownloadURLWithContext(ctx context.Context, request *GetSkillPackageDownloadURLRequest) (response *GetSkillPackageDownloadURLResponse, err error) {
    if request == nil {
        request = NewGetSkillPackageDownloadURLRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "GetSkillPackageDownloadURL")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetSkillPackageDownloadURL require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetSkillPackageDownloadURLResponse()
    err = c.Send(request, response)
    return
}

func NewGetSkillPackageUploadURLRequest() (request *GetSkillPackageUploadURLRequest) {
    request = &GetSkillPackageUploadURLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "GetSkillPackageUploadURL")
    
    
    return
}

func NewGetSkillPackageUploadURLResponse() (response *GetSkillPackageUploadURLResponse) {
    response = &GetSkillPackageUploadURLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetSkillPackageUploadURL
// 为 FAILED / EXPIRED 的 TAR Skill Version 生成新的上传尝试；VersionId 与 Revision 保持不变。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_RECORDID = "InvalidParameter.RecordId"
//  INVALIDPARAMETER_REGISTRYID = "InvalidParameter.RegistryId"
//  INVALIDPARAMETER_VERSIONID = "InvalidParameter.VersionId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_REGISTRYRECORD = "ResourceNotFound.RegistryRecord"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetSkillPackageUploadURL(request *GetSkillPackageUploadURLRequest) (response *GetSkillPackageUploadURLResponse, err error) {
    return c.GetSkillPackageUploadURLWithContext(context.Background(), request)
}

// GetSkillPackageUploadURL
// 为 FAILED / EXPIRED 的 TAR Skill Version 生成新的上传尝试；VersionId 与 Revision 保持不变。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_RECORDID = "InvalidParameter.RecordId"
//  INVALIDPARAMETER_REGISTRYID = "InvalidParameter.RegistryId"
//  INVALIDPARAMETER_VERSIONID = "InvalidParameter.VersionId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_REGISTRYRECORD = "ResourceNotFound.RegistryRecord"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetSkillPackageUploadURLWithContext(ctx context.Context, request *GetSkillPackageUploadURLRequest) (response *GetSkillPackageUploadURLResponse, err error) {
    if request == nil {
        request = NewGetSkillPackageUploadURLRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "GetSkillPackageUploadURL")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetSkillPackageUploadURL require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetSkillPackageUploadURLResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDeploymentRequest() (request *ModifyDeploymentRequest) {
    request = &ModifyDeploymentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "ModifyDeployment")
    
    
    return
}

func NewModifyDeploymentResponse() (response *ModifyDeploymentResponse) {
    response = &ModifyDeploymentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDeployment
// 修改 Deployment
//
// 可能返回的错误码:
//  INVALIDPARAMETER_AFFINITYCONFIGURATION = "InvalidParameter.AffinityConfiguration"
//  INVALIDPARAMETER_DEPLOYMENTID = "InvalidParameter.DeploymentId"
//  INVALIDPARAMETER_LIFECYCLECONFIGURATION = "InvalidParameter.LifecycleConfiguration"
//  INVALIDPARAMETER_SCALINGCONFIGURATION = "InvalidParameter.ScalingConfiguration"
//  RESOURCEINUSE_DEPLOYMENT = "ResourceInUse.Deployment"
//  RESOURCENOTFOUND_DEPLOYMENT = "ResourceNotFound.Deployment"
func (c *Client) ModifyDeployment(request *ModifyDeploymentRequest) (response *ModifyDeploymentResponse, err error) {
    return c.ModifyDeploymentWithContext(context.Background(), request)
}

// ModifyDeployment
// 修改 Deployment
//
// 可能返回的错误码:
//  INVALIDPARAMETER_AFFINITYCONFIGURATION = "InvalidParameter.AffinityConfiguration"
//  INVALIDPARAMETER_DEPLOYMENTID = "InvalidParameter.DeploymentId"
//  INVALIDPARAMETER_LIFECYCLECONFIGURATION = "InvalidParameter.LifecycleConfiguration"
//  INVALIDPARAMETER_SCALINGCONFIGURATION = "InvalidParameter.ScalingConfiguration"
//  RESOURCEINUSE_DEPLOYMENT = "ResourceInUse.Deployment"
//  RESOURCENOTFOUND_DEPLOYMENT = "ResourceNotFound.Deployment"
func (c *Client) ModifyDeploymentWithContext(ctx context.Context, request *ModifyDeploymentRequest) (response *ModifyDeploymentResponse, err error) {
    if request == nil {
        request = NewModifyDeploymentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "ModifyDeployment")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDeployment require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDeploymentResponse()
    err = c.Send(request, response)
    return
}

func NewModifySessionRequest() (request *ModifySessionRequest) {
    request = &ModifySessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "ModifySession")
    
    
    return
}

func NewModifySessionResponse() (response *ModifySessionResponse) {
    response = &ModifySessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySession
// 修改会话信息
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND_SESSIONNOTEXIST = "ResourceNotFound.SessionNotExist"
//  RESOURCENOTFOUND_SESSIONSPACENOTEXIST = "ResourceNotFound.SessionSpaceNotExist"
func (c *Client) ModifySession(request *ModifySessionRequest) (response *ModifySessionResponse, err error) {
    return c.ModifySessionWithContext(context.Background(), request)
}

// ModifySession
// 修改会话信息
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND_SESSIONNOTEXIST = "ResourceNotFound.SessionNotExist"
//  RESOURCENOTFOUND_SESSIONSPACENOTEXIST = "ResourceNotFound.SessionSpaceNotExist"
func (c *Client) ModifySessionWithContext(ctx context.Context, request *ModifySessionRequest) (response *ModifySessionResponse, err error) {
    if request == nil {
        request = NewModifySessionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "ModifySession")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySession require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySessionResponse()
    err = c.Send(request, response)
    return
}

func NewModifySessionSpaceRequest() (request *ModifySessionSpaceRequest) {
    request = &ModifySessionSpaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "ModifySessionSpace")
    
    
    return
}

func NewModifySessionSpaceResponse() (response *ModifySessionSpaceResponse) {
    response = &ModifySessionSpaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySessionSpace
// 修改会话空间。
//
// 修改指定会话空间的名称和描述，修改成功后返回更新后的会话空间信息。默认会话空间允许修改名称和描述。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_SESSIONSPACENOTEXIST = "ResourceNotFound.SessionSpaceNotExist"
func (c *Client) ModifySessionSpace(request *ModifySessionSpaceRequest) (response *ModifySessionSpaceResponse, err error) {
    return c.ModifySessionSpaceWithContext(context.Background(), request)
}

// ModifySessionSpace
// 修改会话空间。
//
// 修改指定会话空间的名称和描述，修改成功后返回更新后的会话空间信息。默认会话空间允许修改名称和描述。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  RESOURCENOTFOUND_SESSIONSPACENOTEXIST = "ResourceNotFound.SessionSpaceNotExist"
func (c *Client) ModifySessionSpaceWithContext(ctx context.Context, request *ModifySessionSpaceRequest) (response *ModifySessionSpaceResponse, err error) {
    if request == nil {
        request = NewModifySessionSpaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "ModifySessionSpace")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySessionSpace require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySessionSpaceResponse()
    err = c.Send(request, response)
    return
}

func NewPauseSandboxInstanceRequest() (request *PauseSandboxInstanceRequest) {
    request = &PauseSandboxInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "PauseSandboxInstance")
    
    
    return
}

func NewPauseSandboxInstanceResponse() (response *PauseSandboxInstanceResponse) {
    response = &PauseSandboxInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PauseSandboxInstance
// 暂停沙箱实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_SANDBOXINSTANCE = "ResourceNotFound.SandboxInstance"
func (c *Client) PauseSandboxInstance(request *PauseSandboxInstanceRequest) (response *PauseSandboxInstanceResponse, err error) {
    return c.PauseSandboxInstanceWithContext(context.Background(), request)
}

// PauseSandboxInstance
// 暂停沙箱实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_SANDBOXINSTANCE = "ResourceNotFound.SandboxInstance"
func (c *Client) PauseSandboxInstanceWithContext(ctx context.Context, request *PauseSandboxInstanceRequest) (response *PauseSandboxInstanceResponse, err error) {
    if request == nil {
        request = NewPauseSandboxInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "PauseSandboxInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("PauseSandboxInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewPauseSandboxInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewPreviewRegistryRecordRequest() (request *PreviewRegistryRecordRequest) {
    request = &PreviewRegistryRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "PreviewRegistryRecord")
    
    
    return
}

func NewPreviewRegistryRecordResponse() (response *PreviewRegistryRecordResponse) {
    response = &PreviewRegistryRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PreviewRegistryRecord
// 对 Record 的指定 Version 或 Label 目标发起一次预览调用。VersionId 与 Label 互斥；均省略时使用 Stable。不创建 Version、不修改 Label。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_RECORDID = "InvalidParameter.RecordId"
//  INVALIDPARAMETER_RECORDSELECTOR = "InvalidParameter.RecordSelector"
//  INVALIDPARAMETER_REGISTRYID = "InvalidParameter.RegistryId"
//  INVALIDPARAMETER_VERSIONID = "InvalidParameter.VersionId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_REGISTRYRECORD = "ResourceNotFound.RegistryRecord"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) PreviewRegistryRecord(request *PreviewRegistryRecordRequest) (response *PreviewRegistryRecordResponse, err error) {
    return c.PreviewRegistryRecordWithContext(context.Background(), request)
}

// PreviewRegistryRecord
// 对 Record 的指定 Version 或 Label 目标发起一次预览调用。VersionId 与 Label 互斥；均省略时使用 Stable。不创建 Version、不修改 Label。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_RECORDID = "InvalidParameter.RecordId"
//  INVALIDPARAMETER_RECORDSELECTOR = "InvalidParameter.RecordSelector"
//  INVALIDPARAMETER_REGISTRYID = "InvalidParameter.RegistryId"
//  INVALIDPARAMETER_VERSIONID = "InvalidParameter.VersionId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_REGISTRYRECORD = "ResourceNotFound.RegistryRecord"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) PreviewRegistryRecordWithContext(ctx context.Context, request *PreviewRegistryRecordRequest) (response *PreviewRegistryRecordResponse, err error) {
    if request == nil {
        request = NewPreviewRegistryRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "PreviewRegistryRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("PreviewRegistryRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewPreviewRegistryRecordResponse()
    err = c.Send(request, response)
    return
}

func NewRejectRegistryRecordRequest() (request *RejectRegistryRecordRequest) {
    request = &RejectRegistryRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "RejectRegistryRecord")
    
    
    return
}

func NewRejectRegistryRecordResponse() (response *RejectRegistryRecordResponse) {
    response = &RejectRegistryRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RejectRegistryRecord
// 驳回 Version 审批：PENDING_APPROVAL → REJECTED。Comment 必填。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_RECORDID = "InvalidParameter.RecordId"
//  INVALIDPARAMETER_REGISTRYID = "InvalidParameter.RegistryId"
//  INVALIDPARAMETER_VERSIONID = "InvalidParameter.VersionId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_REGISTRYRECORD = "ResourceNotFound.RegistryRecord"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RejectRegistryRecord(request *RejectRegistryRecordRequest) (response *RejectRegistryRecordResponse, err error) {
    return c.RejectRegistryRecordWithContext(context.Background(), request)
}

// RejectRegistryRecord
// 驳回 Version 审批：PENDING_APPROVAL → REJECTED。Comment 必填。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_RECORDID = "InvalidParameter.RecordId"
//  INVALIDPARAMETER_REGISTRYID = "InvalidParameter.RegistryId"
//  INVALIDPARAMETER_VERSIONID = "InvalidParameter.VersionId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_REGISTRYRECORD = "ResourceNotFound.RegistryRecord"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RejectRegistryRecordWithContext(ctx context.Context, request *RejectRegistryRecordRequest) (response *RejectRegistryRecordResponse, err error) {
    if request == nil {
        request = NewRejectRegistryRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "RejectRegistryRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RejectRegistryRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewRejectRegistryRecordResponse()
    err = c.Send(request, response)
    return
}

func NewResumeSandboxInstanceRequest() (request *ResumeSandboxInstanceRequest) {
    request = &ResumeSandboxInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "ResumeSandboxInstance")
    
    
    return
}

func NewResumeSandboxInstanceResponse() (response *ResumeSandboxInstanceResponse) {
    response = &ResumeSandboxInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResumeSandboxInstance
// 恢复沙箱实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_SANDBOXINSTANCE = "ResourceNotFound.SandboxInstance"
func (c *Client) ResumeSandboxInstance(request *ResumeSandboxInstanceRequest) (response *ResumeSandboxInstanceResponse, err error) {
    return c.ResumeSandboxInstanceWithContext(context.Background(), request)
}

// ResumeSandboxInstance
// 恢复沙箱实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_SANDBOXINSTANCE = "ResourceNotFound.SandboxInstance"
func (c *Client) ResumeSandboxInstanceWithContext(ctx context.Context, request *ResumeSandboxInstanceRequest) (response *ResumeSandboxInstanceResponse, err error) {
    if request == nil {
        request = NewResumeSandboxInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "ResumeSandboxInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResumeSandboxInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewResumeSandboxInstanceResponse()
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
//  FAILEDOPERATION_DEPENDENCYUNAVAILABLE = "FailedOperation.DependencyUnavailable"
//  FAILEDOPERATION_DUPLICATEREQUEST = "FailedOperation.DuplicateRequest"
//  FAILEDOPERATION_REQUESTINPROGRESS = "FailedOperation.RequestInProgress"
//  FAILEDOPERATION_STORAGEMOUNT = "FailedOperation.StorageMount"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETERVALUE_MOUNTOPTION = "InvalidParameterValue.MountOption"
//  INVALIDPARAMETERVALUE_STORAGEMOUNT = "InvalidParameterValue.StorageMount"
//  INVALIDPARAMETERVALUE_TIMEOUT = "InvalidParameterValue.Timeout"
//  LIMITEXCEEDED_SANDBOXINSTANCE = "LimitExceeded.SandboxInstance"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_ROLEARN = "MissingParameter.RoleArn"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SANDBOXTOOL = "ResourceNotFound.SandboxTool"
//  RESOURCENOTFOUND_STORAGEMOUNT = "ResourceNotFound.StorageMount"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_SANDBOXTOOL = "ResourceUnavailable.SandboxTool"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StartSandboxInstance(request *StartSandboxInstanceRequest) (response *StartSandboxInstanceResponse, err error) {
    return c.StartSandboxInstanceWithContext(context.Background(), request)
}

// StartSandboxInstance
// 启动沙箱实例
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DEPENDENCYUNAVAILABLE = "FailedOperation.DependencyUnavailable"
//  FAILEDOPERATION_DUPLICATEREQUEST = "FailedOperation.DuplicateRequest"
//  FAILEDOPERATION_REQUESTINPROGRESS = "FailedOperation.RequestInProgress"
//  FAILEDOPERATION_STORAGEMOUNT = "FailedOperation.StorageMount"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETERVALUE_MOUNTOPTION = "InvalidParameterValue.MountOption"
//  INVALIDPARAMETERVALUE_STORAGEMOUNT = "InvalidParameterValue.StorageMount"
//  INVALIDPARAMETERVALUE_TIMEOUT = "InvalidParameterValue.Timeout"
//  LIMITEXCEEDED_SANDBOXINSTANCE = "LimitExceeded.SandboxInstance"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_ROLEARN = "MissingParameter.RoleArn"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SANDBOXTOOL = "ResourceNotFound.SandboxTool"
//  RESOURCENOTFOUND_STORAGEMOUNT = "ResourceNotFound.StorageMount"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_SANDBOXTOOL = "ResourceUnavailable.SandboxTool"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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

func NewSyncRegistryRecordRequest() (request *SyncRegistryRecordRequest) {
    request = &SyncRegistryRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "SyncRegistryRecord")
    
    
    return
}

func NewSyncRegistryRecordResponse() (response *SyncRegistryRecordResponse) {
    response = &SyncRegistryRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SyncRegistryRecord
// 触发一次从远端拉取描述符 / 元数据的同步。可通过互斥的 VersionId 或 Label 指定来源 Version，均省略时默认使用 Stable。有变化时创建新 Version 并移动 Latest；来源必须 SourceType=URL_IMPORT，否则返回 UnsupportedOperation.SourceType。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_RECORDID = "InvalidParameter.RecordId"
//  INVALIDPARAMETER_RECORDSELECTOR = "InvalidParameter.RecordSelector"
//  INVALIDPARAMETER_REGISTRYID = "InvalidParameter.RegistryId"
//  INVALIDPARAMETER_VERSIONID = "InvalidParameter.VersionId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_REGISTRYRECORD = "ResourceNotFound.RegistryRecord"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SyncRegistryRecord(request *SyncRegistryRecordRequest) (response *SyncRegistryRecordResponse, err error) {
    return c.SyncRegistryRecordWithContext(context.Background(), request)
}

// SyncRegistryRecord
// 触发一次从远端拉取描述符 / 元数据的同步。可通过互斥的 VersionId 或 Label 指定来源 Version，均省略时默认使用 Stable。有变化时创建新 Version 并移动 Latest；来源必须 SourceType=URL_IMPORT，否则返回 UnsupportedOperation.SourceType。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_RECORDID = "InvalidParameter.RecordId"
//  INVALIDPARAMETER_RECORDSELECTOR = "InvalidParameter.RecordSelector"
//  INVALIDPARAMETER_REGISTRYID = "InvalidParameter.RegistryId"
//  INVALIDPARAMETER_VERSIONID = "InvalidParameter.VersionId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_REGISTRYRECORD = "ResourceNotFound.RegistryRecord"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SyncRegistryRecordWithContext(ctx context.Context, request *SyncRegistryRecordRequest) (response *SyncRegistryRecordResponse, err error) {
    if request == nil {
        request = NewSyncRegistryRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "SyncRegistryRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SyncRegistryRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewSyncRegistryRecordResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateRegistryRequest() (request *UpdateRegistryRequest) {
    request = &UpdateRegistryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "UpdateRegistry")
    
    
    return
}

func NewUpdateRegistryResponse() (response *UpdateRegistryResponse) {
    response = &UpdateRegistryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateRegistry
// 更新 Registry 的可变元数据。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_REGISTRYID = "InvalidParameter.RegistryId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateRegistry(request *UpdateRegistryRequest) (response *UpdateRegistryResponse, err error) {
    return c.UpdateRegistryWithContext(context.Background(), request)
}

// UpdateRegistry
// 更新 Registry 的可变元数据。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_REGISTRYID = "InvalidParameter.RegistryId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateRegistryWithContext(ctx context.Context, request *UpdateRegistryRequest) (response *UpdateRegistryResponse, err error) {
    if request == nil {
        request = NewUpdateRegistryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "UpdateRegistry")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateRegistry require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateRegistryResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateRegistryRecordRequest() (request *UpdateRegistryRecordRequest) {
    request = &UpdateRegistryRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ags", APIVersion, "UpdateRegistryRecord")
    
    
    return
}

func NewUpdateRegistryRecordResponse() (response *UpdateRegistryRecordResponse) {
    response = &UpdateRegistryRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateRegistryRecord
// 更新 Registry Record。两种互斥模式：①Record 更新模式：不提交任何 Source / CustomDescriptors，可通过 Description、LabelMutations 修改元数据与 Label（至少提交一项）；②Version 创建模式：提交且仅提交一种与现有 DescriptorType 匹配的内容输入，可选 VersionName / ChangeLog，禁止 Description / LabelMutations，服务端在 Record 下创建下一个 Revision。取代原 ChangeRegistryRecordStableVersion / RollbackRegistryRecordVersion / Create*RegistryRecordVersion。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_AGENTSOURCE = "InvalidParameter.AgentSource"
//  INVALIDPARAMETER_CUSTOMDESCRIPTORS = "InvalidParameter.CustomDescriptors"
//  INVALIDPARAMETER_MCPSOURCE = "InvalidParameter.MCPSource"
//  INVALIDPARAMETER_RECORDID = "InvalidParameter.RecordId"
//  INVALIDPARAMETER_RECORDSOURCE = "InvalidParameter.RecordSource"
//  INVALIDPARAMETER_REGISTRYID = "InvalidParameter.RegistryId"
//  INVALIDPARAMETER_SKILLSOURCE = "InvalidParameter.SkillSource"
//  INVALIDPARAMETER_VERSIONID = "InvalidParameter.VersionId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_AGENTSOURCETYPE = "InvalidParameterValue.AgentSourceType"
//  INVALIDPARAMETERVALUE_MCPSOURCETYPE = "InvalidParameterValue.MCPSourceType"
//  INVALIDPARAMETERVALUE_SKILLSOURCETYPE = "InvalidParameterValue.SkillSourceType"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_RECORDLABELCOUNT = "LimitExceeded.RecordLabelCount"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_AGENTSOURCEENDPOINTURL = "MissingParameter.AgentSourceEndpointURL"
//  MISSINGPARAMETER_MCPSOURCEDESCRIPTORS = "MissingParameter.MCPSourceDescriptors"
//  MISSINGPARAMETER_RECORDSOURCE = "MissingParameter.RecordSource"
//  MISSINGPARAMETER_SKILLSOURCESKILLMD = "MissingParameter.SkillSourceSkillMd"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_REGISTRYRECORD = "ResourceNotFound.RegistryRecord"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateRegistryRecord(request *UpdateRegistryRecordRequest) (response *UpdateRegistryRecordResponse, err error) {
    return c.UpdateRegistryRecordWithContext(context.Background(), request)
}

// UpdateRegistryRecord
// 更新 Registry Record。两种互斥模式：①Record 更新模式：不提交任何 Source / CustomDescriptors，可通过 Description、LabelMutations 修改元数据与 Label（至少提交一项）；②Version 创建模式：提交且仅提交一种与现有 DescriptorType 匹配的内容输入，可选 VersionName / ChangeLog，禁止 Description / LabelMutations，服务端在 Record 下创建下一个 Revision。取代原 ChangeRegistryRecordStableVersion / RollbackRegistryRecordVersion / Create*RegistryRecordVersion。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_AGENTSOURCE = "InvalidParameter.AgentSource"
//  INVALIDPARAMETER_CUSTOMDESCRIPTORS = "InvalidParameter.CustomDescriptors"
//  INVALIDPARAMETER_MCPSOURCE = "InvalidParameter.MCPSource"
//  INVALIDPARAMETER_RECORDID = "InvalidParameter.RecordId"
//  INVALIDPARAMETER_RECORDSOURCE = "InvalidParameter.RecordSource"
//  INVALIDPARAMETER_REGISTRYID = "InvalidParameter.RegistryId"
//  INVALIDPARAMETER_SKILLSOURCE = "InvalidParameter.SkillSource"
//  INVALIDPARAMETER_VERSIONID = "InvalidParameter.VersionId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_AGENTSOURCETYPE = "InvalidParameterValue.AgentSourceType"
//  INVALIDPARAMETERVALUE_MCPSOURCETYPE = "InvalidParameterValue.MCPSourceType"
//  INVALIDPARAMETERVALUE_SKILLSOURCETYPE = "InvalidParameterValue.SkillSourceType"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_RECORDLABELCOUNT = "LimitExceeded.RecordLabelCount"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_AGENTSOURCEENDPOINTURL = "MissingParameter.AgentSourceEndpointURL"
//  MISSINGPARAMETER_MCPSOURCEDESCRIPTORS = "MissingParameter.MCPSourceDescriptors"
//  MISSINGPARAMETER_RECORDSOURCE = "MissingParameter.RecordSource"
//  MISSINGPARAMETER_SKILLSOURCESKILLMD = "MissingParameter.SkillSourceSkillMd"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_REGISTRYRECORD = "ResourceNotFound.RegistryRecord"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateRegistryRecordWithContext(ctx context.Context, request *UpdateRegistryRecordRequest) (response *UpdateRegistryRecordResponse, err error) {
    if request == nil {
        request = NewUpdateRegistryRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "ags", APIVersion, "UpdateRegistryRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateRegistryRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateRegistryRecordResponse()
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
//  FAILEDOPERATION_DEPENDENCYUNAVAILABLE = "FailedOperation.DependencyUnavailable"
//  FAILEDOPERATION_REQUESTINPROGRESS = "FailedOperation.RequestInProgress"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETERVALUE_TIMEOUT = "InvalidParameterValue.Timeout"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SANDBOXINSTANCE = "ResourceNotFound.SandboxInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION_SANDBOXINSTANCE = "UnsupportedOperation.SandboxInstance"
func (c *Client) UpdateSandboxInstance(request *UpdateSandboxInstanceRequest) (response *UpdateSandboxInstanceResponse, err error) {
    return c.UpdateSandboxInstanceWithContext(context.Background(), request)
}

// UpdateSandboxInstance
// 更新沙箱实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_DEPENDENCYUNAVAILABLE = "FailedOperation.DependencyUnavailable"
//  FAILEDOPERATION_REQUESTINPROGRESS = "FailedOperation.RequestInProgress"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETERVALUE_TIMEOUT = "InvalidParameterValue.Timeout"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SANDBOXINSTANCE = "ResourceNotFound.SandboxInstance"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
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
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DEPENDENCYUNAVAILABLE = "FailedOperation.DependencyUnavailable"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SANDBOXTOOL = "ResourceNotFound.SandboxTool"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) UpdateSandboxTool(request *UpdateSandboxToolRequest) (response *UpdateSandboxToolResponse, err error) {
    return c.UpdateSandboxToolWithContext(context.Background(), request)
}

// UpdateSandboxTool
// 更新沙箱工具
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DEPENDENCYUNAVAILABLE = "FailedOperation.DependencyUnavailable"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SANDBOXTOOL = "ResourceNotFound.SandboxTool"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
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
