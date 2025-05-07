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

package v20191118

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-11-18"

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


func NewBackUpAndroidInstanceToStorageRequest() (request *BackUpAndroidInstanceToStorageRequest) {
    request = &BackUpAndroidInstanceToStorageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "BackUpAndroidInstanceToStorage")
    
    
    return
}

func NewBackUpAndroidInstanceToStorageResponse() (response *BackUpAndroidInstanceToStorageResponse) {
    response = &BackUpAndroidInstanceToStorageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BackUpAndroidInstanceToStorage
// 备份云手机到指定存储
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BackUpAndroidInstanceToStorage(request *BackUpAndroidInstanceToStorageRequest) (response *BackUpAndroidInstanceToStorageResponse, err error) {
    return c.BackUpAndroidInstanceToStorageWithContext(context.Background(), request)
}

// BackUpAndroidInstanceToStorage
// 备份云手机到指定存储
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BackUpAndroidInstanceToStorageWithContext(ctx context.Context, request *BackUpAndroidInstanceToStorageRequest) (response *BackUpAndroidInstanceToStorageResponse, err error) {
    if request == nil {
        request = NewBackUpAndroidInstanceToStorageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BackUpAndroidInstanceToStorage require credential")
    }

    request.SetContext(ctx)
    
    response = NewBackUpAndroidInstanceToStorageResponse()
    err = c.Send(request, response)
    return
}

func NewConnectAndroidInstanceRequest() (request *ConnectAndroidInstanceRequest) {
    request = &ConnectAndroidInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "ConnectAndroidInstance")
    
    
    return
}

func NewConnectAndroidInstanceResponse() (response *ConnectAndroidInstanceResponse) {
    response = &ConnectAndroidInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ConnectAndroidInstance
// 连接安卓实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOIDLE = "ResourceNotFound.NoIdle"
func (c *Client) ConnectAndroidInstance(request *ConnectAndroidInstanceRequest) (response *ConnectAndroidInstanceResponse, err error) {
    return c.ConnectAndroidInstanceWithContext(context.Background(), request)
}

// ConnectAndroidInstance
// 连接安卓实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOIDLE = "ResourceNotFound.NoIdle"
func (c *Client) ConnectAndroidInstanceWithContext(ctx context.Context, request *ConnectAndroidInstanceRequest) (response *ConnectAndroidInstanceResponse, err error) {
    if request == nil {
        request = NewConnectAndroidInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ConnectAndroidInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewConnectAndroidInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCopyAndroidInstanceRequest() (request *CopyAndroidInstanceRequest) {
    request = &CopyAndroidInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "CopyAndroidInstance")
    
    
    return
}

func NewCopyAndroidInstanceResponse() (response *CopyAndroidInstanceResponse) {
    response = &CopyAndroidInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CopyAndroidInstance
// 复制安卓实例：
//
// 1. 排除和包含文件只能指定/data下的文件，不指定时复制整个/data目录
//
// 2. 源实例和目的实例必须在同一区域
//
// 3. 复制时，源实例和目的实例都会停机，复制完后实例会自动启动
//
// 4. 复制时会产生大量内网流量，请限制并发
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CopyAndroidInstance(request *CopyAndroidInstanceRequest) (response *CopyAndroidInstanceResponse, err error) {
    return c.CopyAndroidInstanceWithContext(context.Background(), request)
}

// CopyAndroidInstance
// 复制安卓实例：
//
// 1. 排除和包含文件只能指定/data下的文件，不指定时复制整个/data目录
//
// 2. 源实例和目的实例必须在同一区域
//
// 3. 复制时，源实例和目的实例都会停机，复制完后实例会自动启动
//
// 4. 复制时会产生大量内网流量，请限制并发
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CopyAndroidInstanceWithContext(ctx context.Context, request *CopyAndroidInstanceRequest) (response *CopyAndroidInstanceResponse, err error) {
    if request == nil {
        request = NewCopyAndroidInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CopyAndroidInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCopyAndroidInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAndroidAppRequest() (request *CreateAndroidAppRequest) {
    request = &CreateAndroidAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "CreateAndroidApp")
    
    
    return
}

func NewCreateAndroidAppResponse() (response *CreateAndroidAppResponse) {
    response = &CreateAndroidAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAndroidApp
// 创建安卓应用
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONCREATEFAIL = "FailedOperation.ApplicationCreateFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_APPLICATIONLIMITEXCEEDED = "OperationDenied.ApplicationLimitExceeded"
func (c *Client) CreateAndroidApp(request *CreateAndroidAppRequest) (response *CreateAndroidAppResponse, err error) {
    return c.CreateAndroidAppWithContext(context.Background(), request)
}

// CreateAndroidApp
// 创建安卓应用
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONCREATEFAIL = "FailedOperation.ApplicationCreateFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_APPLICATIONLIMITEXCEEDED = "OperationDenied.ApplicationLimitExceeded"
func (c *Client) CreateAndroidAppWithContext(ctx context.Context, request *CreateAndroidAppRequest) (response *CreateAndroidAppResponse, err error) {
    if request == nil {
        request = NewCreateAndroidAppRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAndroidApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAndroidAppResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAndroidAppVersionRequest() (request *CreateAndroidAppVersionRequest) {
    request = &CreateAndroidAppVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "CreateAndroidAppVersion")
    
    
    return
}

func NewCreateAndroidAppVersionResponse() (response *CreateAndroidAppVersionResponse) {
    response = &CreateAndroidAppVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAndroidAppVersion
// 创建安卓应用版本
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONLOCKFAIL = "FailedOperation.ApplicationLockFail"
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_VERSIONCREATING = "OperationDenied.VersionCreating"
//  OPERATIONDENIED_VERSIONLIMITEXCEEDED = "OperationDenied.VersionLimitExceeded"
func (c *Client) CreateAndroidAppVersion(request *CreateAndroidAppVersionRequest) (response *CreateAndroidAppVersionResponse, err error) {
    return c.CreateAndroidAppVersionWithContext(context.Background(), request)
}

// CreateAndroidAppVersion
// 创建安卓应用版本
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONLOCKFAIL = "FailedOperation.ApplicationLockFail"
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_VERSIONCREATING = "OperationDenied.VersionCreating"
//  OPERATIONDENIED_VERSIONLIMITEXCEEDED = "OperationDenied.VersionLimitExceeded"
func (c *Client) CreateAndroidAppVersionWithContext(ctx context.Context, request *CreateAndroidAppVersionRequest) (response *CreateAndroidAppVersionResponse, err error) {
    if request == nil {
        request = NewCreateAndroidAppVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAndroidAppVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAndroidAppVersionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAndroidInstanceImageRequest() (request *CreateAndroidInstanceImageRequest) {
    request = &CreateAndroidInstanceImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "CreateAndroidInstanceImage")
    
    
    return
}

func NewCreateAndroidInstanceImageResponse() (response *CreateAndroidInstanceImageResponse) {
    response = &CreateAndroidInstanceImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAndroidInstanceImage
// 创建安卓实例镜像
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAndroidInstanceImage(request *CreateAndroidInstanceImageRequest) (response *CreateAndroidInstanceImageResponse, err error) {
    return c.CreateAndroidInstanceImageWithContext(context.Background(), request)
}

// CreateAndroidInstanceImage
// 创建安卓实例镜像
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAndroidInstanceImageWithContext(ctx context.Context, request *CreateAndroidInstanceImageRequest) (response *CreateAndroidInstanceImageResponse, err error) {
    if request == nil {
        request = NewCreateAndroidInstanceImageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAndroidInstanceImage require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAndroidInstanceImageResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAndroidInstanceLabelRequest() (request *CreateAndroidInstanceLabelRequest) {
    request = &CreateAndroidInstanceLabelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "CreateAndroidInstanceLabel")
    
    
    return
}

func NewCreateAndroidInstanceLabelResponse() (response *CreateAndroidInstanceLabelResponse) {
    response = &CreateAndroidInstanceLabelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAndroidInstanceLabel
// 创建安卓实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAndroidInstanceLabel(request *CreateAndroidInstanceLabelRequest) (response *CreateAndroidInstanceLabelResponse, err error) {
    return c.CreateAndroidInstanceLabelWithContext(context.Background(), request)
}

// CreateAndroidInstanceLabel
// 创建安卓实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAndroidInstanceLabelWithContext(ctx context.Context, request *CreateAndroidInstanceLabelRequest) (response *CreateAndroidInstanceLabelResponse, err error) {
    if request == nil {
        request = NewCreateAndroidInstanceLabelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAndroidInstanceLabel require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAndroidInstanceLabelResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAndroidInstanceSSHRequest() (request *CreateAndroidInstanceSSHRequest) {
    request = &CreateAndroidInstanceSSHRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "CreateAndroidInstanceSSH")
    
    
    return
}

func NewCreateAndroidInstanceSSHResponse() (response *CreateAndroidInstanceSSHResponse) {
    response = &CreateAndroidInstanceSSHResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAndroidInstanceSSH
// 创建安卓实例 SSH 连接
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateAndroidInstanceSSH(request *CreateAndroidInstanceSSHRequest) (response *CreateAndroidInstanceSSHResponse, err error) {
    return c.CreateAndroidInstanceSSHWithContext(context.Background(), request)
}

// CreateAndroidInstanceSSH
// 创建安卓实例 SSH 连接
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateAndroidInstanceSSHWithContext(ctx context.Context, request *CreateAndroidInstanceSSHRequest) (response *CreateAndroidInstanceSSHResponse, err error) {
    if request == nil {
        request = NewCreateAndroidInstanceSSHRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAndroidInstanceSSH require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAndroidInstanceSSHResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAndroidInstanceWebShellRequest() (request *CreateAndroidInstanceWebShellRequest) {
    request = &CreateAndroidInstanceWebShellRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "CreateAndroidInstanceWebShell")
    
    
    return
}

func NewCreateAndroidInstanceWebShellResponse() (response *CreateAndroidInstanceWebShellResponse) {
    response = &CreateAndroidInstanceWebShellResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAndroidInstanceWebShell
// 创建安卓实例 WebShell 连接
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateAndroidInstanceWebShell(request *CreateAndroidInstanceWebShellRequest) (response *CreateAndroidInstanceWebShellResponse, err error) {
    return c.CreateAndroidInstanceWebShellWithContext(context.Background(), request)
}

// CreateAndroidInstanceWebShell
// 创建安卓实例 WebShell 连接
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateAndroidInstanceWebShellWithContext(ctx context.Context, request *CreateAndroidInstanceWebShellRequest) (response *CreateAndroidInstanceWebShellResponse, err error) {
    if request == nil {
        request = NewCreateAndroidInstanceWebShellRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAndroidInstanceWebShell require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAndroidInstanceWebShellResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAndroidInstancesRequest() (request *CreateAndroidInstancesRequest) {
    request = &CreateAndroidInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "CreateAndroidInstances")
    
    
    return
}

func NewCreateAndroidInstancesResponse() (response *CreateAndroidInstancesResponse) {
    response = &CreateAndroidInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAndroidInstances
// 创建安卓实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateAndroidInstances(request *CreateAndroidInstancesRequest) (response *CreateAndroidInstancesResponse, err error) {
    return c.CreateAndroidInstancesWithContext(context.Background(), request)
}

// CreateAndroidInstances
// 创建安卓实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateAndroidInstancesWithContext(ctx context.Context, request *CreateAndroidInstancesRequest) (response *CreateAndroidInstancesResponse, err error) {
    if request == nil {
        request = NewCreateAndroidInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAndroidInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAndroidInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAndroidInstancesScreenshotRequest() (request *CreateAndroidInstancesScreenshotRequest) {
    request = &CreateAndroidInstancesScreenshotRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "CreateAndroidInstancesScreenshot")
    
    
    return
}

func NewCreateAndroidInstancesScreenshotResponse() (response *CreateAndroidInstancesScreenshotResponse) {
    response = &CreateAndroidInstancesScreenshotResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAndroidInstancesScreenshot
// 安卓实例截图
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAndroidInstancesScreenshot(request *CreateAndroidInstancesScreenshotRequest) (response *CreateAndroidInstancesScreenshotResponse, err error) {
    return c.CreateAndroidInstancesScreenshotWithContext(context.Background(), request)
}

// CreateAndroidInstancesScreenshot
// 安卓实例截图
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAndroidInstancesScreenshotWithContext(ctx context.Context, request *CreateAndroidInstancesScreenshotRequest) (response *CreateAndroidInstancesScreenshotResponse, err error) {
    if request == nil {
        request = NewCreateAndroidInstancesScreenshotRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAndroidInstancesScreenshot require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAndroidInstancesScreenshotResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCosCredentialRequest() (request *CreateCosCredentialRequest) {
    request = &CreateCosCredentialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "CreateCosCredential")
    
    
    return
}

func NewCreateCosCredentialResponse() (response *CreateCosCredentialResponse) {
    response = &CreateCosCredentialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCosCredential
// 用于创建 Cos 临时密钥
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPLICATIONLOCKFAIL = "FailedOperation.ApplicationLockFail"
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  FAILEDOPERATION_GAMELOCKFAIL = "FailedOperation.GameLockFail"
//  FAILEDOPERATION_GAMENOTFIND = "FailedOperation.GameNotFind"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_VERSIONCREATING = "OperationDenied.VersionCreating"
//  OPERATIONDENIED_VERSIONLIMITEXCEEDED = "OperationDenied.VersionLimitExceeded"
func (c *Client) CreateCosCredential(request *CreateCosCredentialRequest) (response *CreateCosCredentialResponse, err error) {
    return c.CreateCosCredentialWithContext(context.Background(), request)
}

// CreateCosCredential
// 用于创建 Cos 临时密钥
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPLICATIONLOCKFAIL = "FailedOperation.ApplicationLockFail"
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  FAILEDOPERATION_GAMELOCKFAIL = "FailedOperation.GameLockFail"
//  FAILEDOPERATION_GAMENOTFIND = "FailedOperation.GameNotFind"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_VERSIONCREATING = "OperationDenied.VersionCreating"
//  OPERATIONDENIED_VERSIONLIMITEXCEEDED = "OperationDenied.VersionLimitExceeded"
func (c *Client) CreateCosCredentialWithContext(ctx context.Context, request *CreateCosCredentialRequest) (response *CreateCosCredentialResponse, err error) {
    if request == nil {
        request = NewCreateCosCredentialRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCosCredential require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCosCredentialResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSessionRequest() (request *CreateSessionRequest) {
    request = &CreateSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "CreateSession")
    
    
    return
}

func NewCreateSessionResponse() (response *CreateSessionResponse) {
    response = &CreateSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSession
// 创建会话
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOCKTIMEOUT = "FailedOperation.LockTimeout"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  FAILEDOPERATION_SLOWDOWN = "FailedOperation.SlowDown"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  FAILEDOPERATION_TOOFREQUENTLY = "FailedOperation.TooFrequently"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_ROLE = "LimitExceeded.Role"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_NOIDLE = "ResourceNotFound.NoIdle"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
//  RESOURCEUNAVAILABLE_INITIALIZATION = "ResourceUnavailable.Initialization"
//  UNSUPPORTEDOPERATION_STOPPING = "UnsupportedOperation.Stopping"
func (c *Client) CreateSession(request *CreateSessionRequest) (response *CreateSessionResponse, err error) {
    return c.CreateSessionWithContext(context.Background(), request)
}

// CreateSession
// 创建会话
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOCKTIMEOUT = "FailedOperation.LockTimeout"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  FAILEDOPERATION_SLOWDOWN = "FailedOperation.SlowDown"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  FAILEDOPERATION_TOOFREQUENTLY = "FailedOperation.TooFrequently"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_ROLE = "LimitExceeded.Role"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_NOIDLE = "ResourceNotFound.NoIdle"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
//  RESOURCEUNAVAILABLE_INITIALIZATION = "ResourceUnavailable.Initialization"
//  UNSUPPORTEDOPERATION_STOPPING = "UnsupportedOperation.Stopping"
func (c *Client) CreateSessionWithContext(ctx context.Context, request *CreateSessionRequest) (response *CreateSessionResponse, err error) {
    if request == nil {
        request = NewCreateSessionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSession require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSessionResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAndroidAppRequest() (request *DeleteAndroidAppRequest) {
    request = &DeleteAndroidAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "DeleteAndroidApp")
    
    
    return
}

func NewDeleteAndroidAppResponse() (response *DeleteAndroidAppResponse) {
    response = &DeleteAndroidAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAndroidApp
// 删除安卓应用
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONLOCKFAIL = "FailedOperation.ApplicationLockFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DeleteAndroidApp(request *DeleteAndroidAppRequest) (response *DeleteAndroidAppResponse, err error) {
    return c.DeleteAndroidAppWithContext(context.Background(), request)
}

// DeleteAndroidApp
// 删除安卓应用
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONLOCKFAIL = "FailedOperation.ApplicationLockFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DeleteAndroidAppWithContext(ctx context.Context, request *DeleteAndroidAppRequest) (response *DeleteAndroidAppResponse, err error) {
    if request == nil {
        request = NewDeleteAndroidAppRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAndroidApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAndroidAppResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAndroidAppVersionRequest() (request *DeleteAndroidAppVersionRequest) {
    request = &DeleteAndroidAppVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "DeleteAndroidAppVersion")
    
    
    return
}

func NewDeleteAndroidAppVersionResponse() (response *DeleteAndroidAppVersionResponse) {
    response = &DeleteAndroidAppVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAndroidAppVersion
// 删除安卓应用版本
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONLOCKFAIL = "FailedOperation.ApplicationLockFail"
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_VERSIONCREATING = "OperationDenied.VersionCreating"
func (c *Client) DeleteAndroidAppVersion(request *DeleteAndroidAppVersionRequest) (response *DeleteAndroidAppVersionResponse, err error) {
    return c.DeleteAndroidAppVersionWithContext(context.Background(), request)
}

// DeleteAndroidAppVersion
// 删除安卓应用版本
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONLOCKFAIL = "FailedOperation.ApplicationLockFail"
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_VERSIONCREATING = "OperationDenied.VersionCreating"
func (c *Client) DeleteAndroidAppVersionWithContext(ctx context.Context, request *DeleteAndroidAppVersionRequest) (response *DeleteAndroidAppVersionResponse, err error) {
    if request == nil {
        request = NewDeleteAndroidAppVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAndroidAppVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAndroidAppVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAndroidInstanceImagesRequest() (request *DeleteAndroidInstanceImagesRequest) {
    request = &DeleteAndroidInstanceImagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "DeleteAndroidInstanceImages")
    
    
    return
}

func NewDeleteAndroidInstanceImagesResponse() (response *DeleteAndroidInstanceImagesResponse) {
    response = &DeleteAndroidInstanceImagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAndroidInstanceImages
// 删除安卓实例镜像
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteAndroidInstanceImages(request *DeleteAndroidInstanceImagesRequest) (response *DeleteAndroidInstanceImagesResponse, err error) {
    return c.DeleteAndroidInstanceImagesWithContext(context.Background(), request)
}

// DeleteAndroidInstanceImages
// 删除安卓实例镜像
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteAndroidInstanceImagesWithContext(ctx context.Context, request *DeleteAndroidInstanceImagesRequest) (response *DeleteAndroidInstanceImagesResponse, err error) {
    if request == nil {
        request = NewDeleteAndroidInstanceImagesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAndroidInstanceImages require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAndroidInstanceImagesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAndroidInstanceLabelRequest() (request *DeleteAndroidInstanceLabelRequest) {
    request = &DeleteAndroidInstanceLabelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "DeleteAndroidInstanceLabel")
    
    
    return
}

func NewDeleteAndroidInstanceLabelResponse() (response *DeleteAndroidInstanceLabelResponse) {
    response = &DeleteAndroidInstanceLabelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAndroidInstanceLabel
// 创建安卓实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteAndroidInstanceLabel(request *DeleteAndroidInstanceLabelRequest) (response *DeleteAndroidInstanceLabelResponse, err error) {
    return c.DeleteAndroidInstanceLabelWithContext(context.Background(), request)
}

// DeleteAndroidInstanceLabel
// 创建安卓实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteAndroidInstanceLabelWithContext(ctx context.Context, request *DeleteAndroidInstanceLabelRequest) (response *DeleteAndroidInstanceLabelResponse, err error) {
    if request == nil {
        request = NewDeleteAndroidInstanceLabelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAndroidInstanceLabel require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAndroidInstanceLabelResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAndroidAppsRequest() (request *DescribeAndroidAppsRequest) {
    request = &DescribeAndroidAppsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "DescribeAndroidApps")
    
    
    return
}

func NewDescribeAndroidAppsResponse() (response *DescribeAndroidAppsResponse) {
    response = &DescribeAndroidAppsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAndroidApps
// 查询安卓应用信息
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAndroidApps(request *DescribeAndroidAppsRequest) (response *DescribeAndroidAppsResponse, err error) {
    return c.DescribeAndroidAppsWithContext(context.Background(), request)
}

// DescribeAndroidApps
// 查询安卓应用信息
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeAndroidAppsWithContext(ctx context.Context, request *DescribeAndroidAppsRequest) (response *DescribeAndroidAppsResponse, err error) {
    if request == nil {
        request = NewDescribeAndroidAppsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAndroidApps require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAndroidAppsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAndroidInstanceAppsRequest() (request *DescribeAndroidInstanceAppsRequest) {
    request = &DescribeAndroidInstanceAppsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "DescribeAndroidInstanceApps")
    
    
    return
}

func NewDescribeAndroidInstanceAppsResponse() (response *DescribeAndroidInstanceAppsResponse) {
    response = &DescribeAndroidInstanceAppsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAndroidInstanceApps
// 查询安卓实例应用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAndroidInstanceApps(request *DescribeAndroidInstanceAppsRequest) (response *DescribeAndroidInstanceAppsResponse, err error) {
    return c.DescribeAndroidInstanceAppsWithContext(context.Background(), request)
}

// DescribeAndroidInstanceApps
// 查询安卓实例应用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAndroidInstanceAppsWithContext(ctx context.Context, request *DescribeAndroidInstanceAppsRequest) (response *DescribeAndroidInstanceAppsResponse, err error) {
    if request == nil {
        request = NewDescribeAndroidInstanceAppsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAndroidInstanceApps require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAndroidInstanceAppsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAndroidInstanceImagesRequest() (request *DescribeAndroidInstanceImagesRequest) {
    request = &DescribeAndroidInstanceImagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "DescribeAndroidInstanceImages")
    
    
    return
}

func NewDescribeAndroidInstanceImagesResponse() (response *DescribeAndroidInstanceImagesResponse) {
    response = &DescribeAndroidInstanceImagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAndroidInstanceImages
// 查询安卓实例镜像
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAndroidInstanceImages(request *DescribeAndroidInstanceImagesRequest) (response *DescribeAndroidInstanceImagesResponse, err error) {
    return c.DescribeAndroidInstanceImagesWithContext(context.Background(), request)
}

// DescribeAndroidInstanceImages
// 查询安卓实例镜像
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAndroidInstanceImagesWithContext(ctx context.Context, request *DescribeAndroidInstanceImagesRequest) (response *DescribeAndroidInstanceImagesResponse, err error) {
    if request == nil {
        request = NewDescribeAndroidInstanceImagesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAndroidInstanceImages require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAndroidInstanceImagesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAndroidInstanceLabelsRequest() (request *DescribeAndroidInstanceLabelsRequest) {
    request = &DescribeAndroidInstanceLabelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "DescribeAndroidInstanceLabels")
    
    
    return
}

func NewDescribeAndroidInstanceLabelsResponse() (response *DescribeAndroidInstanceLabelsResponse) {
    response = &DescribeAndroidInstanceLabelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAndroidInstanceLabels
// 创建安卓实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAndroidInstanceLabels(request *DescribeAndroidInstanceLabelsRequest) (response *DescribeAndroidInstanceLabelsResponse, err error) {
    return c.DescribeAndroidInstanceLabelsWithContext(context.Background(), request)
}

// DescribeAndroidInstanceLabels
// 创建安卓实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAndroidInstanceLabelsWithContext(ctx context.Context, request *DescribeAndroidInstanceLabelsRequest) (response *DescribeAndroidInstanceLabelsResponse, err error) {
    if request == nil {
        request = NewDescribeAndroidInstanceLabelsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAndroidInstanceLabels require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAndroidInstanceLabelsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAndroidInstanceTasksStatusRequest() (request *DescribeAndroidInstanceTasksStatusRequest) {
    request = &DescribeAndroidInstanceTasksStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "DescribeAndroidInstanceTasksStatus")
    
    
    return
}

func NewDescribeAndroidInstanceTasksStatusResponse() (response *DescribeAndroidInstanceTasksStatusResponse) {
    response = &DescribeAndroidInstanceTasksStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAndroidInstanceTasksStatus
// 查询安卓实例任务状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAndroidInstanceTasksStatus(request *DescribeAndroidInstanceTasksStatusRequest) (response *DescribeAndroidInstanceTasksStatusResponse, err error) {
    return c.DescribeAndroidInstanceTasksStatusWithContext(context.Background(), request)
}

// DescribeAndroidInstanceTasksStatus
// 查询安卓实例任务状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAndroidInstanceTasksStatusWithContext(ctx context.Context, request *DescribeAndroidInstanceTasksStatusRequest) (response *DescribeAndroidInstanceTasksStatusResponse, err error) {
    if request == nil {
        request = NewDescribeAndroidInstanceTasksStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAndroidInstanceTasksStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAndroidInstanceTasksStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAndroidInstancesRequest() (request *DescribeAndroidInstancesRequest) {
    request = &DescribeAndroidInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "DescribeAndroidInstances")
    
    
    return
}

func NewDescribeAndroidInstancesResponse() (response *DescribeAndroidInstancesResponse) {
    response = &DescribeAndroidInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAndroidInstances
// 查询安卓实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAndroidInstances(request *DescribeAndroidInstancesRequest) (response *DescribeAndroidInstancesResponse, err error) {
    return c.DescribeAndroidInstancesWithContext(context.Background(), request)
}

// DescribeAndroidInstances
// 查询安卓实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAndroidInstancesWithContext(ctx context.Context, request *DescribeAndroidInstancesRequest) (response *DescribeAndroidInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeAndroidInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAndroidInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAndroidInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesCountRequest() (request *DescribeInstancesCountRequest) {
    request = &DescribeInstancesCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "DescribeInstancesCount")
    
    
    return
}

func NewDescribeInstancesCountResponse() (response *DescribeInstancesCountResponse) {
    response = &DescribeInstancesCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstancesCount
// 获取并发总数和运行数
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeInstancesCount(request *DescribeInstancesCountRequest) (response *DescribeInstancesCountResponse, err error) {
    return c.DescribeInstancesCountWithContext(context.Background(), request)
}

// DescribeInstancesCount
// 获取并发总数和运行数
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeInstancesCountWithContext(ctx context.Context, request *DescribeInstancesCountRequest) (response *DescribeInstancesCountResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesCountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstancesCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesCountResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyAndroidInstancesRequest() (request *DestroyAndroidInstancesRequest) {
    request = &DestroyAndroidInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "DestroyAndroidInstances")
    
    
    return
}

func NewDestroyAndroidInstancesResponse() (response *DestroyAndroidInstancesResponse) {
    response = &DestroyAndroidInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DestroyAndroidInstances
// 销毁安卓实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DestroyAndroidInstances(request *DestroyAndroidInstancesRequest) (response *DestroyAndroidInstancesResponse, err error) {
    return c.DestroyAndroidInstancesWithContext(context.Background(), request)
}

// DestroyAndroidInstances
// 销毁安卓实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DestroyAndroidInstancesWithContext(ctx context.Context, request *DestroyAndroidInstancesRequest) (response *DestroyAndroidInstancesResponse, err error) {
    if request == nil {
        request = NewDestroyAndroidInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DestroyAndroidInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDestroyAndroidInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewExecuteCommandOnAndroidInstancesRequest() (request *ExecuteCommandOnAndroidInstancesRequest) {
    request = &ExecuteCommandOnAndroidInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "ExecuteCommandOnAndroidInstances")
    
    
    return
}

func NewExecuteCommandOnAndroidInstancesResponse() (response *ExecuteCommandOnAndroidInstancesResponse) {
    response = &ExecuteCommandOnAndroidInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ExecuteCommandOnAndroidInstances
// 在安卓实例上异步执行命令，命令输出结果如果内容过长会被截断
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ExecuteCommandOnAndroidInstances(request *ExecuteCommandOnAndroidInstancesRequest) (response *ExecuteCommandOnAndroidInstancesResponse, err error) {
    return c.ExecuteCommandOnAndroidInstancesWithContext(context.Background(), request)
}

// ExecuteCommandOnAndroidInstances
// 在安卓实例上异步执行命令，命令输出结果如果内容过长会被截断
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ExecuteCommandOnAndroidInstancesWithContext(ctx context.Context, request *ExecuteCommandOnAndroidInstancesRequest) (response *ExecuteCommandOnAndroidInstancesResponse, err error) {
    if request == nil {
        request = NewExecuteCommandOnAndroidInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExecuteCommandOnAndroidInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewExecuteCommandOnAndroidInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewFetchAndroidInstancesLogsRequest() (request *FetchAndroidInstancesLogsRequest) {
    request = &FetchAndroidInstancesLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "FetchAndroidInstancesLogs")
    
    
    return
}

func NewFetchAndroidInstancesLogsResponse() (response *FetchAndroidInstancesLogsResponse) {
    response = &FetchAndroidInstancesLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// FetchAndroidInstancesLogs
// 批量获取安卓实例日志
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) FetchAndroidInstancesLogs(request *FetchAndroidInstancesLogsRequest) (response *FetchAndroidInstancesLogsResponse, err error) {
    return c.FetchAndroidInstancesLogsWithContext(context.Background(), request)
}

// FetchAndroidInstancesLogs
// 批量获取安卓实例日志
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) FetchAndroidInstancesLogsWithContext(ctx context.Context, request *FetchAndroidInstancesLogsRequest) (response *FetchAndroidInstancesLogsResponse, err error) {
    if request == nil {
        request = NewFetchAndroidInstancesLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("FetchAndroidInstancesLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewFetchAndroidInstancesLogsResponse()
    err = c.Send(request, response)
    return
}

func NewInstallAndroidInstancesAppRequest() (request *InstallAndroidInstancesAppRequest) {
    request = &InstallAndroidInstancesAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "InstallAndroidInstancesApp")
    
    
    return
}

func NewInstallAndroidInstancesAppResponse() (response *InstallAndroidInstancesAppResponse) {
    response = &InstallAndroidInstancesAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InstallAndroidInstancesApp
// 安装安卓实例应用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) InstallAndroidInstancesApp(request *InstallAndroidInstancesAppRequest) (response *InstallAndroidInstancesAppResponse, err error) {
    return c.InstallAndroidInstancesAppWithContext(context.Background(), request)
}

// InstallAndroidInstancesApp
// 安装安卓实例应用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) InstallAndroidInstancesAppWithContext(ctx context.Context, request *InstallAndroidInstancesAppRequest) (response *InstallAndroidInstancesAppResponse, err error) {
    if request == nil {
        request = NewInstallAndroidInstancesAppRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InstallAndroidInstancesApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewInstallAndroidInstancesAppResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAndroidAppRequest() (request *ModifyAndroidAppRequest) {
    request = &ModifyAndroidAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "ModifyAndroidApp")
    
    
    return
}

func NewModifyAndroidAppResponse() (response *ModifyAndroidAppResponse) {
    response = &ModifyAndroidAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAndroidApp
// 修改安卓应用信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONLOCKFAIL = "FailedOperation.ApplicationLockFail"
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyAndroidApp(request *ModifyAndroidAppRequest) (response *ModifyAndroidAppResponse, err error) {
    return c.ModifyAndroidAppWithContext(context.Background(), request)
}

// ModifyAndroidApp
// 修改安卓应用信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLICATIONLOCKFAIL = "FailedOperation.ApplicationLockFail"
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyAndroidAppWithContext(ctx context.Context, request *ModifyAndroidAppRequest) (response *ModifyAndroidAppResponse, err error) {
    if request == nil {
        request = NewModifyAndroidAppRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAndroidApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAndroidAppResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAndroidAppVersionRequest() (request *ModifyAndroidAppVersionRequest) {
    request = &ModifyAndroidAppVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "ModifyAndroidAppVersion")
    
    
    return
}

func NewModifyAndroidAppVersionResponse() (response *ModifyAndroidAppVersionResponse) {
    response = &ModifyAndroidAppVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAndroidAppVersion
// 修改安卓应用版本
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPLICATIONLOCKFAIL = "FailedOperation.ApplicationLockFail"
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyAndroidAppVersion(request *ModifyAndroidAppVersionRequest) (response *ModifyAndroidAppVersionResponse, err error) {
    return c.ModifyAndroidAppVersionWithContext(context.Background(), request)
}

// ModifyAndroidAppVersion
// 修改安卓应用版本
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPLICATIONLOCKFAIL = "FailedOperation.ApplicationLockFail"
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyAndroidAppVersionWithContext(ctx context.Context, request *ModifyAndroidAppVersionRequest) (response *ModifyAndroidAppVersionResponse, err error) {
    if request == nil {
        request = NewModifyAndroidAppVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAndroidAppVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAndroidAppVersionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAndroidInstanceInformationRequest() (request *ModifyAndroidInstanceInformationRequest) {
    request = &ModifyAndroidInstanceInformationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "ModifyAndroidInstanceInformation")
    
    
    return
}

func NewModifyAndroidInstanceInformationResponse() (response *ModifyAndroidInstanceInformationResponse) {
    response = &ModifyAndroidInstanceInformationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAndroidInstanceInformation
// 修改安卓实例的信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPLICATIONLOCKFAIL = "FailedOperation.ApplicationLockFail"
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyAndroidInstanceInformation(request *ModifyAndroidInstanceInformationRequest) (response *ModifyAndroidInstanceInformationResponse, err error) {
    return c.ModifyAndroidInstanceInformationWithContext(context.Background(), request)
}

// ModifyAndroidInstanceInformation
// 修改安卓实例的信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPLICATIONLOCKFAIL = "FailedOperation.ApplicationLockFail"
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyAndroidInstanceInformationWithContext(ctx context.Context, request *ModifyAndroidInstanceInformationRequest) (response *ModifyAndroidInstanceInformationResponse, err error) {
    if request == nil {
        request = NewModifyAndroidInstanceInformationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAndroidInstanceInformation require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAndroidInstanceInformationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAndroidInstanceResolutionRequest() (request *ModifyAndroidInstanceResolutionRequest) {
    request = &ModifyAndroidInstanceResolutionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "ModifyAndroidInstanceResolution")
    
    
    return
}

func NewModifyAndroidInstanceResolutionResponse() (response *ModifyAndroidInstanceResolutionResponse) {
    response = &ModifyAndroidInstanceResolutionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAndroidInstanceResolution
// 修改安卓实例分辨率。需要注意的是该接口可能导致正在运行的应用出现闪退，所以建议在实例维护时期才进行调用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPLICATIONLOCKFAIL = "FailedOperation.ApplicationLockFail"
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyAndroidInstanceResolution(request *ModifyAndroidInstanceResolutionRequest) (response *ModifyAndroidInstanceResolutionResponse, err error) {
    return c.ModifyAndroidInstanceResolutionWithContext(context.Background(), request)
}

// ModifyAndroidInstanceResolution
// 修改安卓实例分辨率。需要注意的是该接口可能导致正在运行的应用出现闪退，所以建议在实例维护时期才进行调用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APPLICATIONLOCKFAIL = "FailedOperation.ApplicationLockFail"
//  FAILEDOPERATION_APPLICATIONNOTFIND = "FailedOperation.ApplicationNotFind"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) ModifyAndroidInstanceResolutionWithContext(ctx context.Context, request *ModifyAndroidInstanceResolutionRequest) (response *ModifyAndroidInstanceResolutionResponse, err error) {
    if request == nil {
        request = NewModifyAndroidInstanceResolutionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAndroidInstanceResolution require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAndroidInstanceResolutionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAndroidInstancesInformationRequest() (request *ModifyAndroidInstancesInformationRequest) {
    request = &ModifyAndroidInstancesInformationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "ModifyAndroidInstancesInformation")
    
    
    return
}

func NewModifyAndroidInstancesInformationResponse() (response *ModifyAndroidInstancesInformationResponse) {
    response = &ModifyAndroidInstancesInformationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAndroidInstancesInformation
// 批量修改安卓实例信息
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyAndroidInstancesInformation(request *ModifyAndroidInstancesInformationRequest) (response *ModifyAndroidInstancesInformationResponse, err error) {
    return c.ModifyAndroidInstancesInformationWithContext(context.Background(), request)
}

// ModifyAndroidInstancesInformation
// 批量修改安卓实例信息
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyAndroidInstancesInformationWithContext(ctx context.Context, request *ModifyAndroidInstancesInformationRequest) (response *ModifyAndroidInstancesInformationResponse, err error) {
    if request == nil {
        request = NewModifyAndroidInstancesInformationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAndroidInstancesInformation require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAndroidInstancesInformationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAndroidInstancesLabelsRequest() (request *ModifyAndroidInstancesLabelsRequest) {
    request = &ModifyAndroidInstancesLabelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "ModifyAndroidInstancesLabels")
    
    
    return
}

func NewModifyAndroidInstancesLabelsResponse() (response *ModifyAndroidInstancesLabelsResponse) {
    response = &ModifyAndroidInstancesLabelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAndroidInstancesLabels
// 修改安卓实例分辨率。需要注意的是该接口可能导致正在运行的应用出现闪退，所以建议在实例维护时期才进行调用。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyAndroidInstancesLabels(request *ModifyAndroidInstancesLabelsRequest) (response *ModifyAndroidInstancesLabelsResponse, err error) {
    return c.ModifyAndroidInstancesLabelsWithContext(context.Background(), request)
}

// ModifyAndroidInstancesLabels
// 修改安卓实例分辨率。需要注意的是该接口可能导致正在运行的应用出现闪退，所以建议在实例维护时期才进行调用。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyAndroidInstancesLabelsWithContext(ctx context.Context, request *ModifyAndroidInstancesLabelsRequest) (response *ModifyAndroidInstancesLabelsResponse, err error) {
    if request == nil {
        request = NewModifyAndroidInstancesLabelsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAndroidInstancesLabels require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAndroidInstancesLabelsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAndroidInstancesResolutionRequest() (request *ModifyAndroidInstancesResolutionRequest) {
    request = &ModifyAndroidInstancesResolutionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "ModifyAndroidInstancesResolution")
    
    
    return
}

func NewModifyAndroidInstancesResolutionResponse() (response *ModifyAndroidInstancesResolutionResponse) {
    response = &ModifyAndroidInstancesResolutionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAndroidInstancesResolution
// 修改安卓实例分辨率。需要注意的是该接口需要重启才能生效。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyAndroidInstancesResolution(request *ModifyAndroidInstancesResolutionRequest) (response *ModifyAndroidInstancesResolutionResponse, err error) {
    return c.ModifyAndroidInstancesResolutionWithContext(context.Background(), request)
}

// ModifyAndroidInstancesResolution
// 修改安卓实例分辨率。需要注意的是该接口需要重启才能生效。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyAndroidInstancesResolutionWithContext(ctx context.Context, request *ModifyAndroidInstancesResolutionRequest) (response *ModifyAndroidInstancesResolutionResponse, err error) {
    if request == nil {
        request = NewModifyAndroidInstancesResolutionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAndroidInstancesResolution require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAndroidInstancesResolutionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAndroidInstancesUserIdRequest() (request *ModifyAndroidInstancesUserIdRequest) {
    request = &ModifyAndroidInstancesUserIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "ModifyAndroidInstancesUserId")
    
    
    return
}

func NewModifyAndroidInstancesUserIdResponse() (response *ModifyAndroidInstancesUserIdResponse) {
    response = &ModifyAndroidInstancesUserIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAndroidInstancesUserId
// 批量修改安卓实例的用户ID
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyAndroidInstancesUserId(request *ModifyAndroidInstancesUserIdRequest) (response *ModifyAndroidInstancesUserIdResponse, err error) {
    return c.ModifyAndroidInstancesUserIdWithContext(context.Background(), request)
}

// ModifyAndroidInstancesUserId
// 批量修改安卓实例的用户ID
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) ModifyAndroidInstancesUserIdWithContext(ctx context.Context, request *ModifyAndroidInstancesUserIdRequest) (response *ModifyAndroidInstancesUserIdResponse, err error) {
    if request == nil {
        request = NewModifyAndroidInstancesUserIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAndroidInstancesUserId require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAndroidInstancesUserIdResponse()
    err = c.Send(request, response)
    return
}

func NewRebootAndroidInstanceHostsRequest() (request *RebootAndroidInstanceHostsRequest) {
    request = &RebootAndroidInstanceHostsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "RebootAndroidInstanceHosts")
    
    
    return
}

func NewRebootAndroidInstanceHostsResponse() (response *RebootAndroidInstanceHostsResponse) {
    response = &RebootAndroidInstanceHostsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RebootAndroidInstanceHosts
// 重启安卓实例宿主机。请注意：
//
// 
//
// - 当前每 15 分钟只能重启一次
//
// - 一个宿主机可能有多个云手机实例，重启宿主机会影响运行在上面的所有实例，请确保该宿主机上的所有云手机实例未投入业务使用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RebootAndroidInstanceHosts(request *RebootAndroidInstanceHostsRequest) (response *RebootAndroidInstanceHostsResponse, err error) {
    return c.RebootAndroidInstanceHostsWithContext(context.Background(), request)
}

// RebootAndroidInstanceHosts
// 重启安卓实例宿主机。请注意：
//
// 
//
// - 当前每 15 分钟只能重启一次
//
// - 一个宿主机可能有多个云手机实例，重启宿主机会影响运行在上面的所有实例，请确保该宿主机上的所有云手机实例未投入业务使用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RebootAndroidInstanceHostsWithContext(ctx context.Context, request *RebootAndroidInstanceHostsRequest) (response *RebootAndroidInstanceHostsResponse, err error) {
    if request == nil {
        request = NewRebootAndroidInstanceHostsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RebootAndroidInstanceHosts require credential")
    }

    request.SetContext(ctx)
    
    response = NewRebootAndroidInstanceHostsResponse()
    err = c.Send(request, response)
    return
}

func NewRebootAndroidInstancesRequest() (request *RebootAndroidInstancesRequest) {
    request = &RebootAndroidInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "RebootAndroidInstances")
    
    
    return
}

func NewRebootAndroidInstancesResponse() (response *RebootAndroidInstancesResponse) {
    response = &RebootAndroidInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RebootAndroidInstances
// 重启安卓实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RebootAndroidInstances(request *RebootAndroidInstancesRequest) (response *RebootAndroidInstancesResponse, err error) {
    return c.RebootAndroidInstancesWithContext(context.Background(), request)
}

// RebootAndroidInstances
// 重启安卓实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RebootAndroidInstancesWithContext(ctx context.Context, request *RebootAndroidInstancesRequest) (response *RebootAndroidInstancesResponse, err error) {
    if request == nil {
        request = NewRebootAndroidInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RebootAndroidInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewRebootAndroidInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewResetAndroidInstancesRequest() (request *ResetAndroidInstancesRequest) {
    request = &ResetAndroidInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "ResetAndroidInstances")
    
    
    return
}

func NewResetAndroidInstancesResponse() (response *ResetAndroidInstancesResponse) {
    response = &ResetAndroidInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetAndroidInstances
// 重置安卓实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ResetAndroidInstances(request *ResetAndroidInstancesRequest) (response *ResetAndroidInstancesResponse, err error) {
    return c.ResetAndroidInstancesWithContext(context.Background(), request)
}

// ResetAndroidInstances
// 重置安卓实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ResetAndroidInstancesWithContext(ctx context.Context, request *ResetAndroidInstancesRequest) (response *ResetAndroidInstancesResponse, err error) {
    if request == nil {
        request = NewResetAndroidInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetAndroidInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetAndroidInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewRestartAndroidInstancesAppRequest() (request *RestartAndroidInstancesAppRequest) {
    request = &RestartAndroidInstancesAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "RestartAndroidInstancesApp")
    
    
    return
}

func NewRestartAndroidInstancesAppResponse() (response *RestartAndroidInstancesAppResponse) {
    response = &RestartAndroidInstancesAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RestartAndroidInstancesApp
// 启动安卓实例应用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RestartAndroidInstancesApp(request *RestartAndroidInstancesAppRequest) (response *RestartAndroidInstancesAppResponse, err error) {
    return c.RestartAndroidInstancesAppWithContext(context.Background(), request)
}

// RestartAndroidInstancesApp
// 启动安卓实例应用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RestartAndroidInstancesAppWithContext(ctx context.Context, request *RestartAndroidInstancesAppRequest) (response *RestartAndroidInstancesAppResponse, err error) {
    if request == nil {
        request = NewRestartAndroidInstancesAppRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestartAndroidInstancesApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewRestartAndroidInstancesAppResponse()
    err = c.Send(request, response)
    return
}

func NewRestoreAndroidInstanceFromStorageRequest() (request *RestoreAndroidInstanceFromStorageRequest) {
    request = &RestoreAndroidInstanceFromStorageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "RestoreAndroidInstanceFromStorage")
    
    
    return
}

func NewRestoreAndroidInstanceFromStorageResponse() (response *RestoreAndroidInstanceFromStorageResponse) {
    response = &RestoreAndroidInstanceFromStorageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RestoreAndroidInstanceFromStorage
// 指定存储还原云手机
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RestoreAndroidInstanceFromStorage(request *RestoreAndroidInstanceFromStorageRequest) (response *RestoreAndroidInstanceFromStorageResponse, err error) {
    return c.RestoreAndroidInstanceFromStorageWithContext(context.Background(), request)
}

// RestoreAndroidInstanceFromStorage
// 指定存储还原云手机
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RestoreAndroidInstanceFromStorageWithContext(ctx context.Context, request *RestoreAndroidInstanceFromStorageRequest) (response *RestoreAndroidInstanceFromStorageResponse, err error) {
    if request == nil {
        request = NewRestoreAndroidInstanceFromStorageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestoreAndroidInstanceFromStorage require credential")
    }

    request.SetContext(ctx)
    
    response = NewRestoreAndroidInstanceFromStorageResponse()
    err = c.Send(request, response)
    return
}

func NewSaveGameArchiveRequest() (request *SaveGameArchiveRequest) {
    request = &SaveGameArchiveRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "SaveGameArchive")
    
    
    return
}

func NewSaveGameArchiveResponse() (response *SaveGameArchiveResponse) {
    response = &SaveGameArchiveResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SaveGameArchive
// 保存游戏存档
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SLOWDOWN = "FailedOperation.SlowDown"
//  FAILEDOPERATION_TOOFREQUENTLY = "FailedOperation.TooFrequently"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
func (c *Client) SaveGameArchive(request *SaveGameArchiveRequest) (response *SaveGameArchiveResponse, err error) {
    return c.SaveGameArchiveWithContext(context.Background(), request)
}

// SaveGameArchive
// 保存游戏存档
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SLOWDOWN = "FailedOperation.SlowDown"
//  FAILEDOPERATION_TOOFREQUENTLY = "FailedOperation.TooFrequently"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
func (c *Client) SaveGameArchiveWithContext(ctx context.Context, request *SaveGameArchiveRequest) (response *SaveGameArchiveResponse, err error) {
    if request == nil {
        request = NewSaveGameArchiveRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SaveGameArchive require credential")
    }

    request.SetContext(ctx)
    
    response = NewSaveGameArchiveResponse()
    err = c.Send(request, response)
    return
}

func NewStartAndroidInstancesRequest() (request *StartAndroidInstancesRequest) {
    request = &StartAndroidInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "StartAndroidInstances")
    
    
    return
}

func NewStartAndroidInstancesResponse() (response *StartAndroidInstancesResponse) {
    response = &StartAndroidInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartAndroidInstances
// 重启安卓实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StartAndroidInstances(request *StartAndroidInstancesRequest) (response *StartAndroidInstancesResponse, err error) {
    return c.StartAndroidInstancesWithContext(context.Background(), request)
}

// StartAndroidInstances
// 重启安卓实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StartAndroidInstancesWithContext(ctx context.Context, request *StartAndroidInstancesRequest) (response *StartAndroidInstancesResponse, err error) {
    if request == nil {
        request = NewStartAndroidInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartAndroidInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartAndroidInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewStartAndroidInstancesAppRequest() (request *StartAndroidInstancesAppRequest) {
    request = &StartAndroidInstancesAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "StartAndroidInstancesApp")
    
    
    return
}

func NewStartAndroidInstancesAppResponse() (response *StartAndroidInstancesAppResponse) {
    response = &StartAndroidInstancesAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartAndroidInstancesApp
// 启动安卓实例应用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StartAndroidInstancesApp(request *StartAndroidInstancesAppRequest) (response *StartAndroidInstancesAppResponse, err error) {
    return c.StartAndroidInstancesAppWithContext(context.Background(), request)
}

// StartAndroidInstancesApp
// 启动安卓实例应用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StartAndroidInstancesAppWithContext(ctx context.Context, request *StartAndroidInstancesAppRequest) (response *StartAndroidInstancesAppResponse, err error) {
    if request == nil {
        request = NewStartAndroidInstancesAppRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartAndroidInstancesApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartAndroidInstancesAppResponse()
    err = c.Send(request, response)
    return
}

func NewStartPublishStreamRequest() (request *StartPublishStreamRequest) {
    request = &StartPublishStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "StartPublishStream")
    
    
    return
}

func NewStartPublishStreamResponse() (response *StartPublishStreamResponse) {
    response = &StartPublishStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartPublishStream
// 开始云端推流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
//  UNSUPPORTEDOPERATION_NOTRUNNING = "UnsupportedOperation.NotRunning"
func (c *Client) StartPublishStream(request *StartPublishStreamRequest) (response *StartPublishStreamResponse, err error) {
    return c.StartPublishStreamWithContext(context.Background(), request)
}

// StartPublishStream
// 开始云端推流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
//  UNSUPPORTEDOPERATION_NOTRUNNING = "UnsupportedOperation.NotRunning"
func (c *Client) StartPublishStreamWithContext(ctx context.Context, request *StartPublishStreamRequest) (response *StartPublishStreamResponse, err error) {
    if request == nil {
        request = NewStartPublishStreamRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartPublishStream require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartPublishStreamResponse()
    err = c.Send(request, response)
    return
}

func NewStartPublishStreamToCSSRequest() (request *StartPublishStreamToCSSRequest) {
    request = &StartPublishStreamToCSSRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "StartPublishStreamToCSS")
    
    
    return
}

func NewStartPublishStreamToCSSResponse() (response *StartPublishStreamToCSSResponse) {
    response = &StartPublishStreamToCSSResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartPublishStreamToCSS
// 开始云端推流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
//  UNSUPPORTEDOPERATION_NOTRUNNING = "UnsupportedOperation.NotRunning"
func (c *Client) StartPublishStreamToCSS(request *StartPublishStreamToCSSRequest) (response *StartPublishStreamToCSSResponse, err error) {
    return c.StartPublishStreamToCSSWithContext(context.Background(), request)
}

// StartPublishStreamToCSS
// 开始云端推流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
//  UNSUPPORTEDOPERATION_NOTRUNNING = "UnsupportedOperation.NotRunning"
func (c *Client) StartPublishStreamToCSSWithContext(ctx context.Context, request *StartPublishStreamToCSSRequest) (response *StartPublishStreamToCSSResponse, err error) {
    if request == nil {
        request = NewStartPublishStreamToCSSRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartPublishStreamToCSS require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartPublishStreamToCSSResponse()
    err = c.Send(request, response)
    return
}

func NewStopAndroidInstancesRequest() (request *StopAndroidInstancesRequest) {
    request = &StopAndroidInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "StopAndroidInstances")
    
    
    return
}

func NewStopAndroidInstancesResponse() (response *StopAndroidInstancesResponse) {
    response = &StopAndroidInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopAndroidInstances
// 重启安卓实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StopAndroidInstances(request *StopAndroidInstancesRequest) (response *StopAndroidInstancesResponse, err error) {
    return c.StopAndroidInstancesWithContext(context.Background(), request)
}

// StopAndroidInstances
// 重启安卓实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StopAndroidInstancesWithContext(ctx context.Context, request *StopAndroidInstancesRequest) (response *StopAndroidInstancesResponse, err error) {
    if request == nil {
        request = NewStopAndroidInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopAndroidInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopAndroidInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewStopAndroidInstancesAppRequest() (request *StopAndroidInstancesAppRequest) {
    request = &StopAndroidInstancesAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "StopAndroidInstancesApp")
    
    
    return
}

func NewStopAndroidInstancesAppResponse() (response *StopAndroidInstancesAppResponse) {
    response = &StopAndroidInstancesAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopAndroidInstancesApp
// 停止安卓实例应用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StopAndroidInstancesApp(request *StopAndroidInstancesAppRequest) (response *StopAndroidInstancesAppResponse, err error) {
    return c.StopAndroidInstancesAppWithContext(context.Background(), request)
}

// StopAndroidInstancesApp
// 停止安卓实例应用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StopAndroidInstancesAppWithContext(ctx context.Context, request *StopAndroidInstancesAppRequest) (response *StopAndroidInstancesAppResponse, err error) {
    if request == nil {
        request = NewStopAndroidInstancesAppRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopAndroidInstancesApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopAndroidInstancesAppResponse()
    err = c.Send(request, response)
    return
}

func NewStopGameRequest() (request *StopGameRequest) {
    request = &StopGameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "StopGame")
    
    
    return
}

func NewStopGameResponse() (response *StopGameResponse) {
    response = &StopGameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopGame
// 强制退出游戏
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  RESOURCEUNAVAILABLE_ACCESSFAILED = "ResourceUnavailable.AccessFailed"
func (c *Client) StopGame(request *StopGameRequest) (response *StopGameResponse, err error) {
    return c.StopGameWithContext(context.Background(), request)
}

// StopGame
// 强制退出游戏
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  RESOURCEUNAVAILABLE_ACCESSFAILED = "ResourceUnavailable.AccessFailed"
func (c *Client) StopGameWithContext(ctx context.Context, request *StopGameRequest) (response *StopGameResponse, err error) {
    if request == nil {
        request = NewStopGameRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopGame require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopGameResponse()
    err = c.Send(request, response)
    return
}

func NewStopPublishStreamRequest() (request *StopPublishStreamRequest) {
    request = &StopPublishStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "StopPublishStream")
    
    
    return
}

func NewStopPublishStreamResponse() (response *StopPublishStreamResponse) {
    response = &StopPublishStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopPublishStream
// 停止云端推流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
//  UNSUPPORTEDOPERATION_NOTRUNNING = "UnsupportedOperation.NotRunning"
func (c *Client) StopPublishStream(request *StopPublishStreamRequest) (response *StopPublishStreamResponse, err error) {
    return c.StopPublishStreamWithContext(context.Background(), request)
}

// StopPublishStream
// 停止云端推流
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
//  UNSUPPORTEDOPERATION_NOTRUNNING = "UnsupportedOperation.NotRunning"
func (c *Client) StopPublishStreamWithContext(ctx context.Context, request *StopPublishStreamRequest) (response *StopPublishStreamResponse, err error) {
    if request == nil {
        request = NewStopPublishStreamRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopPublishStream require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopPublishStreamResponse()
    err = c.Send(request, response)
    return
}

func NewSwitchGameArchiveRequest() (request *SwitchGameArchiveRequest) {
    request = &SwitchGameArchiveRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "SwitchGameArchive")
    
    
    return
}

func NewSwitchGameArchiveResponse() (response *SwitchGameArchiveResponse) {
    response = &SwitchGameArchiveResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SwitchGameArchive
// 切换游戏存档
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SLOWDOWN = "FailedOperation.SlowDown"
//  FAILEDOPERATION_TOOFREQUENTLY = "FailedOperation.TooFrequently"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
func (c *Client) SwitchGameArchive(request *SwitchGameArchiveRequest) (response *SwitchGameArchiveResponse, err error) {
    return c.SwitchGameArchiveWithContext(context.Background(), request)
}

// SwitchGameArchive
// 切换游戏存档
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SLOWDOWN = "FailedOperation.SlowDown"
//  FAILEDOPERATION_TOOFREQUENTLY = "FailedOperation.TooFrequently"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
func (c *Client) SwitchGameArchiveWithContext(ctx context.Context, request *SwitchGameArchiveRequest) (response *SwitchGameArchiveResponse, err error) {
    if request == nil {
        request = NewSwitchGameArchiveRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SwitchGameArchive require credential")
    }

    request.SetContext(ctx)
    
    response = NewSwitchGameArchiveResponse()
    err = c.Send(request, response)
    return
}

func NewSyncAndroidInstanceImageRequest() (request *SyncAndroidInstanceImageRequest) {
    request = &SyncAndroidInstanceImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "SyncAndroidInstanceImage")
    
    
    return
}

func NewSyncAndroidInstanceImageResponse() (response *SyncAndroidInstanceImageResponse) {
    response = &SyncAndroidInstanceImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SyncAndroidInstanceImage
// 同步安卓实例镜像
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SLOWDOWN = "FailedOperation.SlowDown"
//  FAILEDOPERATION_TOOFREQUENTLY = "FailedOperation.TooFrequently"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
func (c *Client) SyncAndroidInstanceImage(request *SyncAndroidInstanceImageRequest) (response *SyncAndroidInstanceImageResponse, err error) {
    return c.SyncAndroidInstanceImageWithContext(context.Background(), request)
}

// SyncAndroidInstanceImage
// 同步安卓实例镜像
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SLOWDOWN = "FailedOperation.SlowDown"
//  FAILEDOPERATION_TOOFREQUENTLY = "FailedOperation.TooFrequently"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
func (c *Client) SyncAndroidInstanceImageWithContext(ctx context.Context, request *SyncAndroidInstanceImageRequest) (response *SyncAndroidInstanceImageResponse, err error) {
    if request == nil {
        request = NewSyncAndroidInstanceImageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SyncAndroidInstanceImage require credential")
    }

    request.SetContext(ctx)
    
    response = NewSyncAndroidInstanceImageResponse()
    err = c.Send(request, response)
    return
}

func NewSyncExecuteCommandOnAndroidInstancesRequest() (request *SyncExecuteCommandOnAndroidInstancesRequest) {
    request = &SyncExecuteCommandOnAndroidInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "SyncExecuteCommandOnAndroidInstances")
    
    
    return
}

func NewSyncExecuteCommandOnAndroidInstancesResponse() (response *SyncExecuteCommandOnAndroidInstancesResponse) {
    response = &SyncExecuteCommandOnAndroidInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SyncExecuteCommandOnAndroidInstances
// 在安卓实例上同步执行命令，仅支持1秒内可以返回结果的命令，例如：ls、cd。同时执行的实例数量不能过多，否则可能云api返回超时。不支持超过1秒无法返回或无法自主结束的命令，例如：top、vim，执行结果最大1KB
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SyncExecuteCommandOnAndroidInstances(request *SyncExecuteCommandOnAndroidInstancesRequest) (response *SyncExecuteCommandOnAndroidInstancesResponse, err error) {
    return c.SyncExecuteCommandOnAndroidInstancesWithContext(context.Background(), request)
}

// SyncExecuteCommandOnAndroidInstances
// 在安卓实例上同步执行命令，仅支持1秒内可以返回结果的命令，例如：ls、cd。同时执行的实例数量不能过多，否则可能云api返回超时。不支持超过1秒无法返回或无法自主结束的命令，例如：top、vim，执行结果最大1KB
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SyncExecuteCommandOnAndroidInstancesWithContext(ctx context.Context, request *SyncExecuteCommandOnAndroidInstancesRequest) (response *SyncExecuteCommandOnAndroidInstancesResponse, err error) {
    if request == nil {
        request = NewSyncExecuteCommandOnAndroidInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SyncExecuteCommandOnAndroidInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewSyncExecuteCommandOnAndroidInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewTrylockWorkerRequest() (request *TrylockWorkerRequest) {
    request = &TrylockWorkerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "TrylockWorker")
    
    
    return
}

func NewTrylockWorkerResponse() (response *TrylockWorkerResponse) {
    response = &TrylockWorkerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TrylockWorker
// 尝试锁定机器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_NOIDLE = "ResourceNotFound.NoIdle"
//  RESOURCEUNAVAILABLE_INITIALIZATION = "ResourceUnavailable.Initialization"
//  UNSUPPORTEDOPERATION_STOPPING = "UnsupportedOperation.Stopping"
func (c *Client) TrylockWorker(request *TrylockWorkerRequest) (response *TrylockWorkerResponse, err error) {
    return c.TrylockWorkerWithContext(context.Background(), request)
}

// TrylockWorker
// 尝试锁定机器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JSONPARSEERROR = "InvalidParameter.JsonParseError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_NOIDLE = "ResourceNotFound.NoIdle"
//  RESOURCEUNAVAILABLE_INITIALIZATION = "ResourceUnavailable.Initialization"
//  UNSUPPORTEDOPERATION_STOPPING = "UnsupportedOperation.Stopping"
func (c *Client) TrylockWorkerWithContext(ctx context.Context, request *TrylockWorkerRequest) (response *TrylockWorkerResponse, err error) {
    if request == nil {
        request = NewTrylockWorkerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TrylockWorker require credential")
    }

    request.SetContext(ctx)
    
    response = NewTrylockWorkerResponse()
    err = c.Send(request, response)
    return
}

func NewUninstallAndroidInstancesAppRequest() (request *UninstallAndroidInstancesAppRequest) {
    request = &UninstallAndroidInstancesAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "UninstallAndroidInstancesApp")
    
    
    return
}

func NewUninstallAndroidInstancesAppResponse() (response *UninstallAndroidInstancesAppResponse) {
    response = &UninstallAndroidInstancesAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UninstallAndroidInstancesApp
// 卸载安卓实例应用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UninstallAndroidInstancesApp(request *UninstallAndroidInstancesAppRequest) (response *UninstallAndroidInstancesAppResponse, err error) {
    return c.UninstallAndroidInstancesAppWithContext(context.Background(), request)
}

// UninstallAndroidInstancesApp
// 卸载安卓实例应用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UninstallAndroidInstancesAppWithContext(ctx context.Context, request *UninstallAndroidInstancesAppRequest) (response *UninstallAndroidInstancesAppResponse, err error) {
    if request == nil {
        request = NewUninstallAndroidInstancesAppRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UninstallAndroidInstancesApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewUninstallAndroidInstancesAppResponse()
    err = c.Send(request, response)
    return
}

func NewUploadFileToAndroidInstancesRequest() (request *UploadFileToAndroidInstancesRequest) {
    request = &UploadFileToAndroidInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("gs", APIVersion, "UploadFileToAndroidInstances")
    
    
    return
}

func NewUploadFileToAndroidInstancesResponse() (response *UploadFileToAndroidInstancesResponse) {
    response = &UploadFileToAndroidInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UploadFileToAndroidInstances
// 上传文件到安卓实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UploadFileToAndroidInstances(request *UploadFileToAndroidInstancesRequest) (response *UploadFileToAndroidInstancesResponse, err error) {
    return c.UploadFileToAndroidInstancesWithContext(context.Background(), request)
}

// UploadFileToAndroidInstances
// 上传文件到安卓实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PROCESSTIMEOUT = "FailedOperation.ProcessTimeout"
//  INTERNALERROR = "InternalError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UploadFileToAndroidInstancesWithContext(ctx context.Context, request *UploadFileToAndroidInstancesRequest) (response *UploadFileToAndroidInstancesResponse, err error) {
    if request == nil {
        request = NewUploadFileToAndroidInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadFileToAndroidInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadFileToAndroidInstancesResponse()
    err = c.Send(request, response)
    return
}
