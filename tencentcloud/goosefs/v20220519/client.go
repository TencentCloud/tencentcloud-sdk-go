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

package v20220519

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2022-05-19"

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


func NewAddCrossVpcSubnetSupportForClientNodeRequest() (request *AddCrossVpcSubnetSupportForClientNodeRequest) {
    request = &AddCrossVpcSubnetSupportForClientNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("goosefs", APIVersion, "AddCrossVpcSubnetSupportForClientNode")
    
    
    return
}

func NewAddCrossVpcSubnetSupportForClientNodeResponse() (response *AddCrossVpcSubnetSupportForClientNodeResponse) {
    response = &AddCrossVpcSubnetSupportForClientNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddCrossVpcSubnetSupportForClientNode
// 为客户端节点添加跨vpc或子网访问能力
func (c *Client) AddCrossVpcSubnetSupportForClientNode(request *AddCrossVpcSubnetSupportForClientNodeRequest) (response *AddCrossVpcSubnetSupportForClientNodeResponse, err error) {
    return c.AddCrossVpcSubnetSupportForClientNodeWithContext(context.Background(), request)
}

// AddCrossVpcSubnetSupportForClientNode
// 为客户端节点添加跨vpc或子网访问能力
func (c *Client) AddCrossVpcSubnetSupportForClientNodeWithContext(ctx context.Context, request *AddCrossVpcSubnetSupportForClientNodeRequest) (response *AddCrossVpcSubnetSupportForClientNodeResponse, err error) {
    if request == nil {
        request = NewAddCrossVpcSubnetSupportForClientNodeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "goosefs", APIVersion, "AddCrossVpcSubnetSupportForClientNode")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddCrossVpcSubnetSupportForClientNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddCrossVpcSubnetSupportForClientNodeResponse()
    err = c.Send(request, response)
    return
}

func NewAttachFileSystemBucketRequest() (request *AttachFileSystemBucketRequest) {
    request = &AttachFileSystemBucketRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("goosefs", APIVersion, "AttachFileSystemBucket")
    
    
    return
}

func NewAttachFileSystemBucketResponse() (response *AttachFileSystemBucketResponse) {
    response = &AttachFileSystemBucketResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AttachFileSystemBucket
// 为文件系统关联Bucket
func (c *Client) AttachFileSystemBucket(request *AttachFileSystemBucketRequest) (response *AttachFileSystemBucketResponse, err error) {
    return c.AttachFileSystemBucketWithContext(context.Background(), request)
}

// AttachFileSystemBucket
// 为文件系统关联Bucket
func (c *Client) AttachFileSystemBucketWithContext(ctx context.Context, request *AttachFileSystemBucketRequest) (response *AttachFileSystemBucketResponse, err error) {
    if request == nil {
        request = NewAttachFileSystemBucketRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "goosefs", APIVersion, "AttachFileSystemBucket")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AttachFileSystemBucket require credential")
    }

    request.SetContext(ctx)
    
    response = NewAttachFileSystemBucketResponse()
    err = c.Send(request, response)
    return
}

func NewBatchAddClientNodesRequest() (request *BatchAddClientNodesRequest) {
    request = &BatchAddClientNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("goosefs", APIVersion, "BatchAddClientNodes")
    
    
    return
}

func NewBatchAddClientNodesResponse() (response *BatchAddClientNodesResponse) {
    response = &BatchAddClientNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchAddClientNodes
// 批量添加客户端节点
func (c *Client) BatchAddClientNodes(request *BatchAddClientNodesRequest) (response *BatchAddClientNodesResponse, err error) {
    return c.BatchAddClientNodesWithContext(context.Background(), request)
}

// BatchAddClientNodes
// 批量添加客户端节点
func (c *Client) BatchAddClientNodesWithContext(ctx context.Context, request *BatchAddClientNodesRequest) (response *BatchAddClientNodesResponse, err error) {
    if request == nil {
        request = NewBatchAddClientNodesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "goosefs", APIVersion, "BatchAddClientNodes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchAddClientNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchAddClientNodesResponse()
    err = c.Send(request, response)
    return
}

func NewBatchDeleteClientNodesRequest() (request *BatchDeleteClientNodesRequest) {
    request = &BatchDeleteClientNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("goosefs", APIVersion, "BatchDeleteClientNodes")
    
    
    return
}

func NewBatchDeleteClientNodesResponse() (response *BatchDeleteClientNodesResponse) {
    response = &BatchDeleteClientNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BatchDeleteClientNodes
// 批量删除客户端节点
func (c *Client) BatchDeleteClientNodes(request *BatchDeleteClientNodesRequest) (response *BatchDeleteClientNodesResponse, err error) {
    return c.BatchDeleteClientNodesWithContext(context.Background(), request)
}

// BatchDeleteClientNodes
// 批量删除客户端节点
func (c *Client) BatchDeleteClientNodesWithContext(ctx context.Context, request *BatchDeleteClientNodesRequest) (response *BatchDeleteClientNodesResponse, err error) {
    if request == nil {
        request = NewBatchDeleteClientNodesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "goosefs", APIVersion, "BatchDeleteClientNodes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchDeleteClientNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchDeleteClientNodesResponse()
    err = c.Send(request, response)
    return
}

func NewBuildClientNodeMountCommandRequest() (request *BuildClientNodeMountCommandRequest) {
    request = &BuildClientNodeMountCommandRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("goosefs", APIVersion, "BuildClientNodeMountCommand")
    
    
    return
}

func NewBuildClientNodeMountCommandResponse() (response *BuildClientNodeMountCommandResponse) {
    response = &BuildClientNodeMountCommandResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BuildClientNodeMountCommand
// 生成客户端的挂载命令
func (c *Client) BuildClientNodeMountCommand(request *BuildClientNodeMountCommandRequest) (response *BuildClientNodeMountCommandResponse, err error) {
    return c.BuildClientNodeMountCommandWithContext(context.Background(), request)
}

// BuildClientNodeMountCommand
// 生成客户端的挂载命令
func (c *Client) BuildClientNodeMountCommandWithContext(ctx context.Context, request *BuildClientNodeMountCommandRequest) (response *BuildClientNodeMountCommandResponse, err error) {
    if request == nil {
        request = NewBuildClientNodeMountCommandRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "goosefs", APIVersion, "BuildClientNodeMountCommand")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BuildClientNodeMountCommand require credential")
    }

    request.SetContext(ctx)
    
    response = NewBuildClientNodeMountCommandResponse()
    err = c.Send(request, response)
    return
}

func NewBuildCustomerClusterRequest() (request *BuildCustomerClusterRequest) {
    request = &BuildCustomerClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("goosefs", APIVersion, "BuildCustomerCluster")
    
    
    return
}

func NewBuildCustomerClusterResponse() (response *BuildCustomerClusterResponse) {
    response = &BuildCustomerClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BuildCustomerCluster
// 构建客户端集群
func (c *Client) BuildCustomerCluster(request *BuildCustomerClusterRequest) (response *BuildCustomerClusterResponse, err error) {
    return c.BuildCustomerClusterWithContext(context.Background(), request)
}

// BuildCustomerCluster
// 构建客户端集群
func (c *Client) BuildCustomerClusterWithContext(ctx context.Context, request *BuildCustomerClusterRequest) (response *BuildCustomerClusterResponse, err error) {
    if request == nil {
        request = NewBuildCustomerClusterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "goosefs", APIVersion, "BuildCustomerCluster")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BuildCustomerCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewBuildCustomerClusterResponse()
    err = c.Send(request, response)
    return
}

func NewCancelLoadTaskRequest() (request *CancelLoadTaskRequest) {
    request = &CancelLoadTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("goosefs", APIVersion, "CancelLoadTask")
    
    
    return
}

func NewCancelLoadTaskResponse() (response *CancelLoadTaskResponse) {
    response = &CancelLoadTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CancelLoadTask
// 取消单个预热任务，仅任务在 waiting、running 状态时可以调用此接口。注意，该接口需要 GooseFS 集群版本 ≥ 1.5.1。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CancelLoadTask(request *CancelLoadTaskRequest) (response *CancelLoadTaskResponse, err error) {
    return c.CancelLoadTaskWithContext(context.Background(), request)
}

// CancelLoadTask
// 取消单个预热任务，仅任务在 waiting、running 状态时可以调用此接口。注意，该接口需要 GooseFS 集群版本 ≥ 1.5.1。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CancelLoadTaskWithContext(ctx context.Context, request *CancelLoadTaskRequest) (response *CancelLoadTaskResponse, err error) {
    if request == nil {
        request = NewCancelLoadTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "goosefs", APIVersion, "CancelLoadTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelLoadTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelLoadTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDataRepositoryTaskRequest() (request *CreateDataRepositoryTaskRequest) {
    request = &CreateDataRepositoryTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("goosefs", APIVersion, "CreateDataRepositoryTask")
    
    
    return
}

func NewCreateDataRepositoryTaskResponse() (response *CreateDataRepositoryTaskResponse) {
    response = &CreateDataRepositoryTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDataRepositoryTask
// 创建数据流通任务,包括从将文件系统的数据上传到存储桶下, 以及从存储桶下载到文件系统里。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDataRepositoryTask(request *CreateDataRepositoryTaskRequest) (response *CreateDataRepositoryTaskResponse, err error) {
    return c.CreateDataRepositoryTaskWithContext(context.Background(), request)
}

// CreateDataRepositoryTask
// 创建数据流通任务,包括从将文件系统的数据上传到存储桶下, 以及从存储桶下载到文件系统里。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDataRepositoryTaskWithContext(ctx context.Context, request *CreateDataRepositoryTaskRequest) (response *CreateDataRepositoryTaskResponse, err error) {
    if request == nil {
        request = NewCreateDataRepositoryTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "goosefs", APIVersion, "CreateDataRepositoryTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDataRepositoryTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDataRepositoryTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFileSystemRequest() (request *CreateFileSystemRequest) {
    request = &CreateFileSystemRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("goosefs", APIVersion, "CreateFileSystem")
    
    
    return
}

func NewCreateFileSystemResponse() (response *CreateFileSystemResponse) {
    response = &CreateFileSystemResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateFileSystem
// 创建文件系统
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateFileSystem(request *CreateFileSystemRequest) (response *CreateFileSystemResponse, err error) {
    return c.CreateFileSystemWithContext(context.Background(), request)
}

// CreateFileSystem
// 创建文件系统
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateFileSystemWithContext(ctx context.Context, request *CreateFileSystemRequest) (response *CreateFileSystemResponse, err error) {
    if request == nil {
        request = NewCreateFileSystemRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "goosefs", APIVersion, "CreateFileSystem")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFileSystem require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFileSystemResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFilesetRequest() (request *CreateFilesetRequest) {
    request = &CreateFilesetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("goosefs", APIVersion, "CreateFileset")
    
    
    return
}

func NewCreateFilesetResponse() (response *CreateFilesetResponse) {
    response = &CreateFilesetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateFileset
// 创建Fileset
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateFileset(request *CreateFilesetRequest) (response *CreateFilesetResponse, err error) {
    return c.CreateFilesetWithContext(context.Background(), request)
}

// CreateFileset
// 创建Fileset
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateFilesetWithContext(ctx context.Context, request *CreateFilesetRequest) (response *CreateFilesetResponse, err error) {
    if request == nil {
        request = NewCreateFilesetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "goosefs", APIVersion, "CreateFileset")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFileset require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFilesetResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLoadTaskRequest() (request *CreateLoadTaskRequest) {
    request = &CreateLoadTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("goosefs", APIVersion, "CreateLoadTask")
    
    
    return
}

func NewCreateLoadTaskResponse() (response *CreateLoadTaskResponse) {
    response = &CreateLoadTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLoadTask
// GooseFS 预热相关接口，用于下发，列出，查询，修改预热任务。用于元数据预热、数据预热场景。 注意，该接口需要 GooseFS 集群版本 ≥ 1.5.1。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateLoadTask(request *CreateLoadTaskRequest) (response *CreateLoadTaskResponse, err error) {
    return c.CreateLoadTaskWithContext(context.Background(), request)
}

// CreateLoadTask
// GooseFS 预热相关接口，用于下发，列出，查询，修改预热任务。用于元数据预热、数据预热场景。 注意，该接口需要 GooseFS 集群版本 ≥ 1.5.1。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateLoadTaskWithContext(ctx context.Context, request *CreateLoadTaskRequest) (response *CreateLoadTaskResponse, err error) {
    if request == nil {
        request = NewCreateLoadTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "goosefs", APIVersion, "CreateLoadTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLoadTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLoadTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCrossVpcSubnetSupportForClientNodeRequest() (request *DeleteCrossVpcSubnetSupportForClientNodeRequest) {
    request = &DeleteCrossVpcSubnetSupportForClientNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("goosefs", APIVersion, "DeleteCrossVpcSubnetSupportForClientNode")
    
    
    return
}

func NewDeleteCrossVpcSubnetSupportForClientNodeResponse() (response *DeleteCrossVpcSubnetSupportForClientNodeResponse) {
    response = &DeleteCrossVpcSubnetSupportForClientNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCrossVpcSubnetSupportForClientNode
// 为客户端节点删除跨vpc子网访问能力
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCrossVpcSubnetSupportForClientNode(request *DeleteCrossVpcSubnetSupportForClientNodeRequest) (response *DeleteCrossVpcSubnetSupportForClientNodeResponse, err error) {
    return c.DeleteCrossVpcSubnetSupportForClientNodeWithContext(context.Background(), request)
}

// DeleteCrossVpcSubnetSupportForClientNode
// 为客户端节点删除跨vpc子网访问能力
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCrossVpcSubnetSupportForClientNodeWithContext(ctx context.Context, request *DeleteCrossVpcSubnetSupportForClientNodeRequest) (response *DeleteCrossVpcSubnetSupportForClientNodeResponse, err error) {
    if request == nil {
        request = NewDeleteCrossVpcSubnetSupportForClientNodeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "goosefs", APIVersion, "DeleteCrossVpcSubnetSupportForClientNode")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCrossVpcSubnetSupportForClientNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCrossVpcSubnetSupportForClientNodeResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCustomerClusterRequest() (request *DeleteCustomerClusterRequest) {
    request = &DeleteCustomerClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("goosefs", APIVersion, "DeleteCustomerCluster")
    
    
    return
}

func NewDeleteCustomerClusterResponse() (response *DeleteCustomerClusterResponse) {
    response = &DeleteCustomerClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCustomerCluster
// 删除客户端集群
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCustomerCluster(request *DeleteCustomerClusterRequest) (response *DeleteCustomerClusterResponse, err error) {
    return c.DeleteCustomerClusterWithContext(context.Background(), request)
}

// DeleteCustomerCluster
// 删除客户端集群
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteCustomerClusterWithContext(ctx context.Context, request *DeleteCustomerClusterRequest) (response *DeleteCustomerClusterResponse, err error) {
    if request == nil {
        request = NewDeleteCustomerClusterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "goosefs", APIVersion, "DeleteCustomerCluster")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCustomerCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCustomerClusterResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteFileSystemRequest() (request *DeleteFileSystemRequest) {
    request = &DeleteFileSystemRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("goosefs", APIVersion, "DeleteFileSystem")
    
    
    return
}

func NewDeleteFileSystemResponse() (response *DeleteFileSystemResponse) {
    response = &DeleteFileSystemResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteFileSystem
// 删除文件系统
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteFileSystem(request *DeleteFileSystemRequest) (response *DeleteFileSystemResponse, err error) {
    return c.DeleteFileSystemWithContext(context.Background(), request)
}

// DeleteFileSystem
// 删除文件系统
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteFileSystemWithContext(ctx context.Context, request *DeleteFileSystemRequest) (response *DeleteFileSystemResponse, err error) {
    if request == nil {
        request = NewDeleteFileSystemRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "goosefs", APIVersion, "DeleteFileSystem")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteFileSystem require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteFileSystemResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteFilesetRequest() (request *DeleteFilesetRequest) {
    request = &DeleteFilesetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("goosefs", APIVersion, "DeleteFileset")
    
    
    return
}

func NewDeleteFilesetResponse() (response *DeleteFilesetResponse) {
    response = &DeleteFilesetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteFileset
// 删除Fileset
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteFileset(request *DeleteFilesetRequest) (response *DeleteFilesetResponse, err error) {
    return c.DeleteFilesetWithContext(context.Background(), request)
}

// DeleteFileset
// 删除Fileset
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteFilesetWithContext(ctx context.Context, request *DeleteFilesetRequest) (response *DeleteFilesetResponse, err error) {
    if request == nil {
        request = NewDeleteFilesetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "goosefs", APIVersion, "DeleteFileset")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteFileset require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteFilesetResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClientNodesRequest() (request *DescribeClientNodesRequest) {
    request = &DescribeClientNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("goosefs", APIVersion, "DescribeClientNodes")
    
    
    return
}

func NewDescribeClientNodesResponse() (response *DescribeClientNodesResponse) {
    response = &DescribeClientNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClientNodes
// 列出集群中所有的客户端节点
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeClientNodes(request *DescribeClientNodesRequest) (response *DescribeClientNodesResponse, err error) {
    return c.DescribeClientNodesWithContext(context.Background(), request)
}

// DescribeClientNodes
// 列出集群中所有的客户端节点
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeClientNodesWithContext(ctx context.Context, request *DescribeClientNodesRequest) (response *DescribeClientNodesResponse, err error) {
    if request == nil {
        request = NewDescribeClientNodesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "goosefs", APIVersion, "DescribeClientNodes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClientNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClientNodesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterClientTokenRequest() (request *DescribeClusterClientTokenRequest) {
    request = &DescribeClusterClientTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("goosefs", APIVersion, "DescribeClusterClientToken")
    
    
    return
}

func NewDescribeClusterClientTokenResponse() (response *DescribeClusterClientTokenResponse) {
    response = &DescribeClusterClientTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterClientToken
// 查询GooseFS集群客户端凭证
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeClusterClientToken(request *DescribeClusterClientTokenRequest) (response *DescribeClusterClientTokenResponse, err error) {
    return c.DescribeClusterClientTokenWithContext(context.Background(), request)
}

// DescribeClusterClientToken
// 查询GooseFS集群客户端凭证
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeClusterClientTokenWithContext(ctx context.Context, request *DescribeClusterClientTokenRequest) (response *DescribeClusterClientTokenResponse, err error) {
    if request == nil {
        request = NewDescribeClusterClientTokenRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "goosefs", APIVersion, "DescribeClusterClientToken")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterClientToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterClientTokenResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterRoleTokenRequest() (request *DescribeClusterRoleTokenRequest) {
    request = &DescribeClusterRoleTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("goosefs", APIVersion, "DescribeClusterRoleToken")
    
    
    return
}

func NewDescribeClusterRoleTokenResponse() (response *DescribeClusterRoleTokenResponse) {
    response = &DescribeClusterRoleTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterRoleToken
// 查询GooseFS集群角色凭证
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeClusterRoleToken(request *DescribeClusterRoleTokenRequest) (response *DescribeClusterRoleTokenResponse, err error) {
    return c.DescribeClusterRoleTokenWithContext(context.Background(), request)
}

// DescribeClusterRoleToken
// 查询GooseFS集群角色凭证
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeClusterRoleTokenWithContext(ctx context.Context, request *DescribeClusterRoleTokenRequest) (response *DescribeClusterRoleTokenResponse, err error) {
    if request == nil {
        request = NewDescribeClusterRoleTokenRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "goosefs", APIVersion, "DescribeClusterRoleToken")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterRoleToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterRoleTokenResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomerClusterRequest() (request *DescribeCustomerClusterRequest) {
    request = &DescribeCustomerClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("goosefs", APIVersion, "DescribeCustomerCluster")
    
    
    return
}

func NewDescribeCustomerClusterResponse() (response *DescribeCustomerClusterResponse) {
    response = &DescribeCustomerClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCustomerCluster
// 查询客户端集群
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCustomerCluster(request *DescribeCustomerClusterRequest) (response *DescribeCustomerClusterResponse, err error) {
    return c.DescribeCustomerClusterWithContext(context.Background(), request)
}

// DescribeCustomerCluster
// 查询客户端集群
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCustomerClusterWithContext(ctx context.Context, request *DescribeCustomerClusterRequest) (response *DescribeCustomerClusterResponse, err error) {
    if request == nil {
        request = NewDescribeCustomerClusterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "goosefs", APIVersion, "DescribeCustomerCluster")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCustomerCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCustomerClusterResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataRepositoryTaskStatusRequest() (request *DescribeDataRepositoryTaskStatusRequest) {
    request = &DescribeDataRepositoryTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("goosefs", APIVersion, "DescribeDataRepositoryTaskStatus")
    
    
    return
}

func NewDescribeDataRepositoryTaskStatusResponse() (response *DescribeDataRepositoryTaskStatusResponse) {
    response = &DescribeDataRepositoryTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDataRepositoryTaskStatus
// 获取数据流通任务实时状态，用作客户端控制
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataRepositoryTaskStatus(request *DescribeDataRepositoryTaskStatusRequest) (response *DescribeDataRepositoryTaskStatusResponse, err error) {
    return c.DescribeDataRepositoryTaskStatusWithContext(context.Background(), request)
}

// DescribeDataRepositoryTaskStatus
// 获取数据流通任务实时状态，用作客户端控制
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDataRepositoryTaskStatusWithContext(ctx context.Context, request *DescribeDataRepositoryTaskStatusRequest) (response *DescribeDataRepositoryTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeDataRepositoryTaskStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "goosefs", APIVersion, "DescribeDataRepositoryTaskStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataRepositoryTaskStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataRepositoryTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFileSystemBucketsRequest() (request *DescribeFileSystemBucketsRequest) {
    request = &DescribeFileSystemBucketsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("goosefs", APIVersion, "DescribeFileSystemBuckets")
    
    
    return
}

func NewDescribeFileSystemBucketsResponse() (response *DescribeFileSystemBucketsResponse) {
    response = &DescribeFileSystemBucketsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFileSystemBuckets
// 罗列文件系统关联的Bucket映射
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeFileSystemBuckets(request *DescribeFileSystemBucketsRequest) (response *DescribeFileSystemBucketsResponse, err error) {
    return c.DescribeFileSystemBucketsWithContext(context.Background(), request)
}

// DescribeFileSystemBuckets
// 罗列文件系统关联的Bucket映射
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeFileSystemBucketsWithContext(ctx context.Context, request *DescribeFileSystemBucketsRequest) (response *DescribeFileSystemBucketsResponse, err error) {
    if request == nil {
        request = NewDescribeFileSystemBucketsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "goosefs", APIVersion, "DescribeFileSystemBuckets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFileSystemBuckets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFileSystemBucketsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFileSystemsRequest() (request *DescribeFileSystemsRequest) {
    request = &DescribeFileSystemsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("goosefs", APIVersion, "DescribeFileSystems")
    
    
    return
}

func NewDescribeFileSystemsResponse() (response *DescribeFileSystemsResponse) {
    response = &DescribeFileSystemsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFileSystems
// 列出所有的文件系统
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeFileSystems(request *DescribeFileSystemsRequest) (response *DescribeFileSystemsResponse, err error) {
    return c.DescribeFileSystemsWithContext(context.Background(), request)
}

// DescribeFileSystems
// 列出所有的文件系统
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeFileSystemsWithContext(ctx context.Context, request *DescribeFileSystemsRequest) (response *DescribeFileSystemsResponse, err error) {
    if request == nil {
        request = NewDescribeFileSystemsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "goosefs", APIVersion, "DescribeFileSystems")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFileSystems require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFileSystemsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFilesetGeneralConfigRequest() (request *DescribeFilesetGeneralConfigRequest) {
    request = &DescribeFilesetGeneralConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("goosefs", APIVersion, "DescribeFilesetGeneralConfig")
    
    
    return
}

func NewDescribeFilesetGeneralConfigResponse() (response *DescribeFilesetGeneralConfigResponse) {
    response = &DescribeFilesetGeneralConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFilesetGeneralConfig
// 查询Fileset通用配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeFilesetGeneralConfig(request *DescribeFilesetGeneralConfigRequest) (response *DescribeFilesetGeneralConfigResponse, err error) {
    return c.DescribeFilesetGeneralConfigWithContext(context.Background(), request)
}

// DescribeFilesetGeneralConfig
// 查询Fileset通用配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeFilesetGeneralConfigWithContext(ctx context.Context, request *DescribeFilesetGeneralConfigRequest) (response *DescribeFilesetGeneralConfigResponse, err error) {
    if request == nil {
        request = NewDescribeFilesetGeneralConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "goosefs", APIVersion, "DescribeFilesetGeneralConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFilesetGeneralConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFilesetGeneralConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFilesetsRequest() (request *DescribeFilesetsRequest) {
    request = &DescribeFilesetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("goosefs", APIVersion, "DescribeFilesets")
    
    
    return
}

func NewDescribeFilesetsResponse() (response *DescribeFilesetsResponse) {
    response = &DescribeFilesetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFilesets
// 查询Fileset列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeFilesets(request *DescribeFilesetsRequest) (response *DescribeFilesetsResponse, err error) {
    return c.DescribeFilesetsWithContext(context.Background(), request)
}

// DescribeFilesets
// 查询Fileset列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeFilesetsWithContext(ctx context.Context, request *DescribeFilesetsRequest) (response *DescribeFilesetsResponse, err error) {
    if request == nil {
        request = NewDescribeFilesetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "goosefs", APIVersion, "DescribeFilesets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFilesets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFilesetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLoadTaskRequest() (request *DescribeLoadTaskRequest) {
    request = &DescribeLoadTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("goosefs", APIVersion, "DescribeLoadTask")
    
    
    return
}

func NewDescribeLoadTaskResponse() (response *DescribeLoadTaskResponse) {
    response = &DescribeLoadTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLoadTask
// 查询单个预热任务执行情况。注意，该接口需要 GooseFS 集群版本 ≥ 1.5.1。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeLoadTask(request *DescribeLoadTaskRequest) (response *DescribeLoadTaskResponse, err error) {
    return c.DescribeLoadTaskWithContext(context.Background(), request)
}

// DescribeLoadTask
// 查询单个预热任务执行情况。注意，该接口需要 GooseFS 集群版本 ≥ 1.5.1。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeLoadTaskWithContext(ctx context.Context, request *DescribeLoadTaskRequest) (response *DescribeLoadTaskResponse, err error) {
    if request == nil {
        request = NewDescribeLoadTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "goosefs", APIVersion, "DescribeLoadTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLoadTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLoadTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDetachFileSystemBucketRequest() (request *DetachFileSystemBucketRequest) {
    request = &DetachFileSystemBucketRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("goosefs", APIVersion, "DetachFileSystemBucket")
    
    
    return
}

func NewDetachFileSystemBucketResponse() (response *DetachFileSystemBucketResponse) {
    response = &DetachFileSystemBucketResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DetachFileSystemBucket
// 解绑文件系统与Bucket的映射
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DetachFileSystemBucket(request *DetachFileSystemBucketRequest) (response *DetachFileSystemBucketResponse, err error) {
    return c.DetachFileSystemBucketWithContext(context.Background(), request)
}

// DetachFileSystemBucket
// 解绑文件系统与Bucket的映射
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DetachFileSystemBucketWithContext(ctx context.Context, request *DetachFileSystemBucketRequest) (response *DetachFileSystemBucketResponse, err error) {
    if request == nil {
        request = NewDetachFileSystemBucketRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "goosefs", APIVersion, "DetachFileSystemBucket")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetachFileSystemBucket require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetachFileSystemBucketResponse()
    err = c.Send(request, response)
    return
}

func NewExpandCapacityRequest() (request *ExpandCapacityRequest) {
    request = &ExpandCapacityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("goosefs", APIVersion, "ExpandCapacity")
    
    
    return
}

func NewExpandCapacityResponse() (response *ExpandCapacityResponse) {
    response = &ExpandCapacityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ExpandCapacity
// 扩展文件系统容量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ExpandCapacity(request *ExpandCapacityRequest) (response *ExpandCapacityResponse, err error) {
    return c.ExpandCapacityWithContext(context.Background(), request)
}

// ExpandCapacity
// 扩展文件系统容量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ExpandCapacityWithContext(ctx context.Context, request *ExpandCapacityRequest) (response *ExpandCapacityResponse, err error) {
    if request == nil {
        request = NewExpandCapacityRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "goosefs", APIVersion, "ExpandCapacity")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExpandCapacity require credential")
    }

    request.SetContext(ctx)
    
    response = NewExpandCapacityResponse()
    err = c.Send(request, response)
    return
}

func NewListLoadTasksRequest() (request *ListLoadTasksRequest) {
    request = &ListLoadTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("goosefs", APIVersion, "ListLoadTasks")
    
    
    return
}

func NewListLoadTasksResponse() (response *ListLoadTasksResponse) {
    response = &ListLoadTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListLoadTasks
// 列出该集群下所有预热任务。注意，该接口需要 GooseFS 集群版本 ≥ 1.5.1。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ListLoadTasks(request *ListLoadTasksRequest) (response *ListLoadTasksResponse, err error) {
    return c.ListLoadTasksWithContext(context.Background(), request)
}

// ListLoadTasks
// 列出该集群下所有预热任务。注意，该接口需要 GooseFS 集群版本 ≥ 1.5.1。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ListLoadTasksWithContext(ctx context.Context, request *ListLoadTasksRequest) (response *ListLoadTasksResponse, err error) {
    if request == nil {
        request = NewListLoadTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "goosefs", APIVersion, "ListLoadTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListLoadTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewListLoadTasksResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDataRepositoryBandwidthRequest() (request *ModifyDataRepositoryBandwidthRequest) {
    request = &ModifyDataRepositoryBandwidthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("goosefs", APIVersion, "ModifyDataRepositoryBandwidth")
    
    
    return
}

func NewModifyDataRepositoryBandwidthResponse() (response *ModifyDataRepositoryBandwidthResponse) {
    response = &ModifyDataRepositoryBandwidthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDataRepositoryBandwidth
// 修改数据流动带宽
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDataRepositoryBandwidth(request *ModifyDataRepositoryBandwidthRequest) (response *ModifyDataRepositoryBandwidthResponse, err error) {
    return c.ModifyDataRepositoryBandwidthWithContext(context.Background(), request)
}

// ModifyDataRepositoryBandwidth
// 修改数据流动带宽
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDataRepositoryBandwidthWithContext(ctx context.Context, request *ModifyDataRepositoryBandwidthRequest) (response *ModifyDataRepositoryBandwidthResponse, err error) {
    if request == nil {
        request = NewModifyDataRepositoryBandwidthRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "goosefs", APIVersion, "ModifyDataRepositoryBandwidth")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDataRepositoryBandwidth require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDataRepositoryBandwidthResponse()
    err = c.Send(request, response)
    return
}

func NewMountMultipleStorageFileSystemRequest() (request *MountMultipleStorageFileSystemRequest) {
    request = &MountMultipleStorageFileSystemRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("goosefs", APIVersion, "MountMultipleStorageFileSystem")
    
    
    return
}

func NewMountMultipleStorageFileSystemResponse() (response *MountMultipleStorageFileSystemResponse) {
    response = &MountMultipleStorageFileSystemResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// MountMultipleStorageFileSystem
// 客户端集群挂载存储集群
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) MountMultipleStorageFileSystem(request *MountMultipleStorageFileSystemRequest) (response *MountMultipleStorageFileSystemResponse, err error) {
    return c.MountMultipleStorageFileSystemWithContext(context.Background(), request)
}

// MountMultipleStorageFileSystem
// 客户端集群挂载存储集群
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) MountMultipleStorageFileSystemWithContext(ctx context.Context, request *MountMultipleStorageFileSystemRequest) (response *MountMultipleStorageFileSystemResponse, err error) {
    if request == nil {
        request = NewMountMultipleStorageFileSystemRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "goosefs", APIVersion, "MountMultipleStorageFileSystem")
    
    if c.GetCredential() == nil {
        return nil, errors.New("MountMultipleStorageFileSystem require credential")
    }

    request.SetContext(ctx)
    
    response = NewMountMultipleStorageFileSystemResponse()
    err = c.Send(request, response)
    return
}

func NewQueryClientNodeMountCommandRequest() (request *QueryClientNodeMountCommandRequest) {
    request = &QueryClientNodeMountCommandRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("goosefs", APIVersion, "QueryClientNodeMountCommand")
    
    
    return
}

func NewQueryClientNodeMountCommandResponse() (response *QueryClientNodeMountCommandResponse) {
    response = &QueryClientNodeMountCommandResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryClientNodeMountCommand
// 生成客户端的挂载命令
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) QueryClientNodeMountCommand(request *QueryClientNodeMountCommandRequest) (response *QueryClientNodeMountCommandResponse, err error) {
    return c.QueryClientNodeMountCommandWithContext(context.Background(), request)
}

// QueryClientNodeMountCommand
// 生成客户端的挂载命令
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) QueryClientNodeMountCommandWithContext(ctx context.Context, request *QueryClientNodeMountCommandRequest) (response *QueryClientNodeMountCommandResponse, err error) {
    if request == nil {
        request = NewQueryClientNodeMountCommandRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "goosefs", APIVersion, "QueryClientNodeMountCommand")
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryClientNodeMountCommand require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryClientNodeMountCommandResponse()
    err = c.Send(request, response)
    return
}

func NewQueryCrossVpcSubnetSupportForClientNodeRequest() (request *QueryCrossVpcSubnetSupportForClientNodeRequest) {
    request = &QueryCrossVpcSubnetSupportForClientNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("goosefs", APIVersion, "QueryCrossVpcSubnetSupportForClientNode")
    
    
    return
}

func NewQueryCrossVpcSubnetSupportForClientNodeResponse() (response *QueryCrossVpcSubnetSupportForClientNodeResponse) {
    response = &QueryCrossVpcSubnetSupportForClientNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryCrossVpcSubnetSupportForClientNode
// 查询客户端节点跨vpc子网访问能力
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) QueryCrossVpcSubnetSupportForClientNode(request *QueryCrossVpcSubnetSupportForClientNodeRequest) (response *QueryCrossVpcSubnetSupportForClientNodeResponse, err error) {
    return c.QueryCrossVpcSubnetSupportForClientNodeWithContext(context.Background(), request)
}

// QueryCrossVpcSubnetSupportForClientNode
// 查询客户端节点跨vpc子网访问能力
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) QueryCrossVpcSubnetSupportForClientNodeWithContext(ctx context.Context, request *QueryCrossVpcSubnetSupportForClientNodeRequest) (response *QueryCrossVpcSubnetSupportForClientNodeResponse, err error) {
    if request == nil {
        request = NewQueryCrossVpcSubnetSupportForClientNodeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "goosefs", APIVersion, "QueryCrossVpcSubnetSupportForClientNode")
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryCrossVpcSubnetSupportForClientNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryCrossVpcSubnetSupportForClientNodeResponse()
    err = c.Send(request, response)
    return
}

func NewQueryDataRepositoryBandwidthRequest() (request *QueryDataRepositoryBandwidthRequest) {
    request = &QueryDataRepositoryBandwidthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("goosefs", APIVersion, "QueryDataRepositoryBandwidth")
    
    
    return
}

func NewQueryDataRepositoryBandwidthResponse() (response *QueryDataRepositoryBandwidthResponse) {
    response = &QueryDataRepositoryBandwidthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryDataRepositoryBandwidth
// 查询数据流动带宽
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) QueryDataRepositoryBandwidth(request *QueryDataRepositoryBandwidthRequest) (response *QueryDataRepositoryBandwidthResponse, err error) {
    return c.QueryDataRepositoryBandwidthWithContext(context.Background(), request)
}

// QueryDataRepositoryBandwidth
// 查询数据流动带宽
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) QueryDataRepositoryBandwidthWithContext(ctx context.Context, request *QueryDataRepositoryBandwidthRequest) (response *QueryDataRepositoryBandwidthResponse, err error) {
    if request == nil {
        request = NewQueryDataRepositoryBandwidthRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "goosefs", APIVersion, "QueryDataRepositoryBandwidth")
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryDataRepositoryBandwidth require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryDataRepositoryBandwidthResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateFilesetRequest() (request *UpdateFilesetRequest) {
    request = &UpdateFilesetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("goosefs", APIVersion, "UpdateFileset")
    
    
    return
}

func NewUpdateFilesetResponse() (response *UpdateFilesetResponse) {
    response = &UpdateFilesetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateFileset
// 修改FIleset
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateFileset(request *UpdateFilesetRequest) (response *UpdateFilesetResponse, err error) {
    return c.UpdateFilesetWithContext(context.Background(), request)
}

// UpdateFileset
// 修改FIleset
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateFilesetWithContext(ctx context.Context, request *UpdateFilesetRequest) (response *UpdateFilesetResponse, err error) {
    if request == nil {
        request = NewUpdateFilesetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "goosefs", APIVersion, "UpdateFileset")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateFileset require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateFilesetResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateFilesetGeneralConfigRequest() (request *UpdateFilesetGeneralConfigRequest) {
    request = &UpdateFilesetGeneralConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("goosefs", APIVersion, "UpdateFilesetGeneralConfig")
    
    
    return
}

func NewUpdateFilesetGeneralConfigResponse() (response *UpdateFilesetGeneralConfigResponse) {
    response = &UpdateFilesetGeneralConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateFilesetGeneralConfig
// 修改Fileset通用配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateFilesetGeneralConfig(request *UpdateFilesetGeneralConfigRequest) (response *UpdateFilesetGeneralConfigResponse, err error) {
    return c.UpdateFilesetGeneralConfigWithContext(context.Background(), request)
}

// UpdateFilesetGeneralConfig
// 修改Fileset通用配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateFilesetGeneralConfigWithContext(ctx context.Context, request *UpdateFilesetGeneralConfigRequest) (response *UpdateFilesetGeneralConfigResponse, err error) {
    if request == nil {
        request = NewUpdateFilesetGeneralConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "goosefs", APIVersion, "UpdateFilesetGeneralConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateFilesetGeneralConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateFilesetGeneralConfigResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateLoadTaskPriorityRequest() (request *UpdateLoadTaskPriorityRequest) {
    request = &UpdateLoadTaskPriorityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("goosefs", APIVersion, "UpdateLoadTaskPriority")
    
    
    return
}

func NewUpdateLoadTaskPriorityResponse() (response *UpdateLoadTaskPriorityResponse) {
    response = &UpdateLoadTaskPriorityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateLoadTaskPriority
// 变更已有 GooseFS 预热任务配置，仅任务状态为 waiting 时可调用该接口。注意，该接口需要 GooseFS 集群版本 ≥ 1.5.1。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateLoadTaskPriority(request *UpdateLoadTaskPriorityRequest) (response *UpdateLoadTaskPriorityResponse, err error) {
    return c.UpdateLoadTaskPriorityWithContext(context.Background(), request)
}

// UpdateLoadTaskPriority
// 变更已有 GooseFS 预热任务配置，仅任务状态为 waiting 时可调用该接口。注意，该接口需要 GooseFS 集群版本 ≥ 1.5.1。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REGIONERROR = "RegionError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateLoadTaskPriorityWithContext(ctx context.Context, request *UpdateLoadTaskPriorityRequest) (response *UpdateLoadTaskPriorityResponse, err error) {
    if request == nil {
        request = NewUpdateLoadTaskPriorityRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "goosefs", APIVersion, "UpdateLoadTaskPriority")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateLoadTaskPriority require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateLoadTaskPriorityResponse()
    err = c.Send(request, response)
    return
}
