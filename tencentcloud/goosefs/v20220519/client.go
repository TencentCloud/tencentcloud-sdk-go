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
    
    if c.GetCredential() == nil {
        return nil, errors.New("BuildClientNodeMountCommand require credential")
    }

    request.SetContext(ctx)
    
    response = NewBuildClientNodeMountCommandResponse()
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
func (c *Client) CreateDataRepositoryTask(request *CreateDataRepositoryTaskRequest) (response *CreateDataRepositoryTaskResponse, err error) {
    return c.CreateDataRepositoryTaskWithContext(context.Background(), request)
}

// CreateDataRepositoryTask
// 创建数据流通任务,包括从将文件系统的数据上传到存储桶下, 以及从存储桶下载到文件系统里。
func (c *Client) CreateDataRepositoryTaskWithContext(ctx context.Context, request *CreateDataRepositoryTaskRequest) (response *CreateDataRepositoryTaskResponse, err error) {
    if request == nil {
        request = NewCreateDataRepositoryTaskRequest()
    }
    
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
func (c *Client) CreateFileSystem(request *CreateFileSystemRequest) (response *CreateFileSystemResponse, err error) {
    return c.CreateFileSystemWithContext(context.Background(), request)
}

// CreateFileSystem
// 创建文件系统
func (c *Client) CreateFileSystemWithContext(ctx context.Context, request *CreateFileSystemRequest) (response *CreateFileSystemResponse, err error) {
    if request == nil {
        request = NewCreateFileSystemRequest()
    }
    
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
func (c *Client) CreateFileset(request *CreateFilesetRequest) (response *CreateFilesetResponse, err error) {
    return c.CreateFilesetWithContext(context.Background(), request)
}

// CreateFileset
// 创建Fileset
func (c *Client) CreateFilesetWithContext(ctx context.Context, request *CreateFilesetRequest) (response *CreateFilesetResponse, err error) {
    if request == nil {
        request = NewCreateFilesetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFileset require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFilesetResponse()
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
func (c *Client) DeleteCrossVpcSubnetSupportForClientNode(request *DeleteCrossVpcSubnetSupportForClientNodeRequest) (response *DeleteCrossVpcSubnetSupportForClientNodeResponse, err error) {
    return c.DeleteCrossVpcSubnetSupportForClientNodeWithContext(context.Background(), request)
}

// DeleteCrossVpcSubnetSupportForClientNode
// 为客户端节点删除跨vpc子网访问能力
func (c *Client) DeleteCrossVpcSubnetSupportForClientNodeWithContext(ctx context.Context, request *DeleteCrossVpcSubnetSupportForClientNodeRequest) (response *DeleteCrossVpcSubnetSupportForClientNodeResponse, err error) {
    if request == nil {
        request = NewDeleteCrossVpcSubnetSupportForClientNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCrossVpcSubnetSupportForClientNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCrossVpcSubnetSupportForClientNodeResponse()
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
func (c *Client) DeleteFileSystem(request *DeleteFileSystemRequest) (response *DeleteFileSystemResponse, err error) {
    return c.DeleteFileSystemWithContext(context.Background(), request)
}

// DeleteFileSystem
// 删除文件系统
func (c *Client) DeleteFileSystemWithContext(ctx context.Context, request *DeleteFileSystemRequest) (response *DeleteFileSystemResponse, err error) {
    if request == nil {
        request = NewDeleteFileSystemRequest()
    }
    
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
func (c *Client) DeleteFileset(request *DeleteFilesetRequest) (response *DeleteFilesetResponse, err error) {
    return c.DeleteFilesetWithContext(context.Background(), request)
}

// DeleteFileset
// 删除Fileset
func (c *Client) DeleteFilesetWithContext(ctx context.Context, request *DeleteFilesetRequest) (response *DeleteFilesetResponse, err error) {
    if request == nil {
        request = NewDeleteFilesetRequest()
    }
    
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
func (c *Client) DescribeClientNodes(request *DescribeClientNodesRequest) (response *DescribeClientNodesResponse, err error) {
    return c.DescribeClientNodesWithContext(context.Background(), request)
}

// DescribeClientNodes
// 列出集群中所有的客户端节点
func (c *Client) DescribeClientNodesWithContext(ctx context.Context, request *DescribeClientNodesRequest) (response *DescribeClientNodesResponse, err error) {
    if request == nil {
        request = NewDescribeClientNodesRequest()
    }
    
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
func (c *Client) DescribeClusterClientToken(request *DescribeClusterClientTokenRequest) (response *DescribeClusterClientTokenResponse, err error) {
    return c.DescribeClusterClientTokenWithContext(context.Background(), request)
}

// DescribeClusterClientToken
// 查询GooseFS集群客户端凭证
func (c *Client) DescribeClusterClientTokenWithContext(ctx context.Context, request *DescribeClusterClientTokenRequest) (response *DescribeClusterClientTokenResponse, err error) {
    if request == nil {
        request = NewDescribeClusterClientTokenRequest()
    }
    
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
func (c *Client) DescribeClusterRoleToken(request *DescribeClusterRoleTokenRequest) (response *DescribeClusterRoleTokenResponse, err error) {
    return c.DescribeClusterRoleTokenWithContext(context.Background(), request)
}

// DescribeClusterRoleToken
// 查询GooseFS集群角色凭证
func (c *Client) DescribeClusterRoleTokenWithContext(ctx context.Context, request *DescribeClusterRoleTokenRequest) (response *DescribeClusterRoleTokenResponse, err error) {
    if request == nil {
        request = NewDescribeClusterRoleTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterRoleToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterRoleTokenResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterRolesRequest() (request *DescribeClusterRolesRequest) {
    request = &DescribeClusterRolesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("goosefs", APIVersion, "DescribeClusterRoles")
    
    
    return
}

func NewDescribeClusterRolesResponse() (response *DescribeClusterRolesResponse) {
    response = &DescribeClusterRolesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterRoles
// 接口废弃
//
// 
//
// 查询GooseFS集群角色
func (c *Client) DescribeClusterRoles(request *DescribeClusterRolesRequest) (response *DescribeClusterRolesResponse, err error) {
    return c.DescribeClusterRolesWithContext(context.Background(), request)
}

// DescribeClusterRoles
// 接口废弃
//
// 
//
// 查询GooseFS集群角色
func (c *Client) DescribeClusterRolesWithContext(ctx context.Context, request *DescribeClusterRolesRequest) (response *DescribeClusterRolesResponse, err error) {
    if request == nil {
        request = NewDescribeClusterRolesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterRoles require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterRolesResponse()
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
func (c *Client) DescribeDataRepositoryTaskStatus(request *DescribeDataRepositoryTaskStatusRequest) (response *DescribeDataRepositoryTaskStatusResponse, err error) {
    return c.DescribeDataRepositoryTaskStatusWithContext(context.Background(), request)
}

// DescribeDataRepositoryTaskStatus
// 获取数据流通任务实时状态，用作客户端控制
func (c *Client) DescribeDataRepositoryTaskStatusWithContext(ctx context.Context, request *DescribeDataRepositoryTaskStatusRequest) (response *DescribeDataRepositoryTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeDataRepositoryTaskStatusRequest()
    }
    
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
func (c *Client) DescribeFileSystemBuckets(request *DescribeFileSystemBucketsRequest) (response *DescribeFileSystemBucketsResponse, err error) {
    return c.DescribeFileSystemBucketsWithContext(context.Background(), request)
}

// DescribeFileSystemBuckets
// 罗列文件系统关联的Bucket映射
func (c *Client) DescribeFileSystemBucketsWithContext(ctx context.Context, request *DescribeFileSystemBucketsRequest) (response *DescribeFileSystemBucketsResponse, err error) {
    if request == nil {
        request = NewDescribeFileSystemBucketsRequest()
    }
    
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
func (c *Client) DescribeFileSystems(request *DescribeFileSystemsRequest) (response *DescribeFileSystemsResponse, err error) {
    return c.DescribeFileSystemsWithContext(context.Background(), request)
}

// DescribeFileSystems
// 列出所有的文件系统
func (c *Client) DescribeFileSystemsWithContext(ctx context.Context, request *DescribeFileSystemsRequest) (response *DescribeFileSystemsResponse, err error) {
    if request == nil {
        request = NewDescribeFileSystemsRequest()
    }
    
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
func (c *Client) DescribeFilesetGeneralConfig(request *DescribeFilesetGeneralConfigRequest) (response *DescribeFilesetGeneralConfigResponse, err error) {
    return c.DescribeFilesetGeneralConfigWithContext(context.Background(), request)
}

// DescribeFilesetGeneralConfig
// 查询Fileset通用配置
func (c *Client) DescribeFilesetGeneralConfigWithContext(ctx context.Context, request *DescribeFilesetGeneralConfigRequest) (response *DescribeFilesetGeneralConfigResponse, err error) {
    if request == nil {
        request = NewDescribeFilesetGeneralConfigRequest()
    }
    
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
func (c *Client) DescribeFilesets(request *DescribeFilesetsRequest) (response *DescribeFilesetsResponse, err error) {
    return c.DescribeFilesetsWithContext(context.Background(), request)
}

// DescribeFilesets
// 查询Fileset列表
func (c *Client) DescribeFilesetsWithContext(ctx context.Context, request *DescribeFilesetsRequest) (response *DescribeFilesetsResponse, err error) {
    if request == nil {
        request = NewDescribeFilesetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFilesets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFilesetsResponse()
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
func (c *Client) DetachFileSystemBucket(request *DetachFileSystemBucketRequest) (response *DetachFileSystemBucketResponse, err error) {
    return c.DetachFileSystemBucketWithContext(context.Background(), request)
}

// DetachFileSystemBucket
// 解绑文件系统与Bucket的映射
func (c *Client) DetachFileSystemBucketWithContext(ctx context.Context, request *DetachFileSystemBucketRequest) (response *DetachFileSystemBucketResponse, err error) {
    if request == nil {
        request = NewDetachFileSystemBucketRequest()
    }
    
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
func (c *Client) ExpandCapacity(request *ExpandCapacityRequest) (response *ExpandCapacityResponse, err error) {
    return c.ExpandCapacityWithContext(context.Background(), request)
}

// ExpandCapacity
// 扩展文件系统容量
func (c *Client) ExpandCapacityWithContext(ctx context.Context, request *ExpandCapacityRequest) (response *ExpandCapacityResponse, err error) {
    if request == nil {
        request = NewExpandCapacityRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExpandCapacity require credential")
    }

    request.SetContext(ctx)
    
    response = NewExpandCapacityResponse()
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
func (c *Client) ModifyDataRepositoryBandwidth(request *ModifyDataRepositoryBandwidthRequest) (response *ModifyDataRepositoryBandwidthResponse, err error) {
    return c.ModifyDataRepositoryBandwidthWithContext(context.Background(), request)
}

// ModifyDataRepositoryBandwidth
// 修改数据流动带宽
func (c *Client) ModifyDataRepositoryBandwidthWithContext(ctx context.Context, request *ModifyDataRepositoryBandwidthRequest) (response *ModifyDataRepositoryBandwidthResponse, err error) {
    if request == nil {
        request = NewModifyDataRepositoryBandwidthRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDataRepositoryBandwidth require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDataRepositoryBandwidthResponse()
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
func (c *Client) QueryCrossVpcSubnetSupportForClientNode(request *QueryCrossVpcSubnetSupportForClientNodeRequest) (response *QueryCrossVpcSubnetSupportForClientNodeResponse, err error) {
    return c.QueryCrossVpcSubnetSupportForClientNodeWithContext(context.Background(), request)
}

// QueryCrossVpcSubnetSupportForClientNode
// 查询客户端节点跨vpc子网访问能力
func (c *Client) QueryCrossVpcSubnetSupportForClientNodeWithContext(ctx context.Context, request *QueryCrossVpcSubnetSupportForClientNodeRequest) (response *QueryCrossVpcSubnetSupportForClientNodeResponse, err error) {
    if request == nil {
        request = NewQueryCrossVpcSubnetSupportForClientNodeRequest()
    }
    
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
func (c *Client) QueryDataRepositoryBandwidth(request *QueryDataRepositoryBandwidthRequest) (response *QueryDataRepositoryBandwidthResponse, err error) {
    return c.QueryDataRepositoryBandwidthWithContext(context.Background(), request)
}

// QueryDataRepositoryBandwidth
// 查询数据流动带宽
func (c *Client) QueryDataRepositoryBandwidthWithContext(ctx context.Context, request *QueryDataRepositoryBandwidthRequest) (response *QueryDataRepositoryBandwidthResponse, err error) {
    if request == nil {
        request = NewQueryDataRepositoryBandwidthRequest()
    }
    
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
func (c *Client) UpdateFileset(request *UpdateFilesetRequest) (response *UpdateFilesetResponse, err error) {
    return c.UpdateFilesetWithContext(context.Background(), request)
}

// UpdateFileset
// 修改FIleset
func (c *Client) UpdateFilesetWithContext(ctx context.Context, request *UpdateFilesetRequest) (response *UpdateFilesetResponse, err error) {
    if request == nil {
        request = NewUpdateFilesetRequest()
    }
    
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
func (c *Client) UpdateFilesetGeneralConfig(request *UpdateFilesetGeneralConfigRequest) (response *UpdateFilesetGeneralConfigResponse, err error) {
    return c.UpdateFilesetGeneralConfigWithContext(context.Background(), request)
}

// UpdateFilesetGeneralConfig
// 修改Fileset通用配置
func (c *Client) UpdateFilesetGeneralConfigWithContext(ctx context.Context, request *UpdateFilesetGeneralConfigRequest) (response *UpdateFilesetGeneralConfigResponse, err error) {
    if request == nil {
        request = NewUpdateFilesetGeneralConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateFilesetGeneralConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateFilesetGeneralConfigResponse()
    err = c.Send(request, response)
    return
}
