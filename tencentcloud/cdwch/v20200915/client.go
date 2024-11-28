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

package v20200915

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-09-15"

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


func NewActionAlterCkUserRequest() (request *ActionAlterCkUserRequest) {
    request = &ActionAlterCkUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwch", APIVersion, "ActionAlterCkUser")
    
    
    return
}

func NewActionAlterCkUserResponse() (response *ActionAlterCkUserResponse) {
    response = &ActionAlterCkUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ActionAlterCkUser
// 新增和修改用户接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ActionAlterCkUser(request *ActionAlterCkUserRequest) (response *ActionAlterCkUserResponse, err error) {
    return c.ActionAlterCkUserWithContext(context.Background(), request)
}

// ActionAlterCkUser
// 新增和修改用户接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ActionAlterCkUserWithContext(ctx context.Context, request *ActionAlterCkUserRequest) (response *ActionAlterCkUserResponse, err error) {
    if request == nil {
        request = NewActionAlterCkUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ActionAlterCkUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewActionAlterCkUserResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBackUpScheduleRequest() (request *CreateBackUpScheduleRequest) {
    request = &CreateBackUpScheduleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwch", APIVersion, "CreateBackUpSchedule")
    
    
    return
}

func NewCreateBackUpScheduleResponse() (response *CreateBackUpScheduleResponse) {
    response = &CreateBackUpScheduleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateBackUpSchedule
// 创建或者修改备份策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) CreateBackUpSchedule(request *CreateBackUpScheduleRequest) (response *CreateBackUpScheduleResponse, err error) {
    return c.CreateBackUpScheduleWithContext(context.Background(), request)
}

// CreateBackUpSchedule
// 创建或者修改备份策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) CreateBackUpScheduleWithContext(ctx context.Context, request *CreateBackUpScheduleRequest) (response *CreateBackUpScheduleResponse, err error) {
    if request == nil {
        request = NewCreateBackUpScheduleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBackUpSchedule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBackUpScheduleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstanceNewRequest() (request *CreateInstanceNewRequest) {
    request = &CreateInstanceNewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwch", APIVersion, "CreateInstanceNew")
    
    
    return
}

func NewCreateInstanceNewResponse() (response *CreateInstanceNewResponse) {
    response = &CreateInstanceNewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateInstanceNew
// 创建集群
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) CreateInstanceNew(request *CreateInstanceNewRequest) (response *CreateInstanceNewResponse, err error) {
    return c.CreateInstanceNewWithContext(context.Background(), request)
}

// CreateInstanceNew
// 创建集群
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) CreateInstanceNewWithContext(ctx context.Context, request *CreateInstanceNewRequest) (response *CreateInstanceNewResponse, err error) {
    if request == nil {
        request = NewCreateInstanceNewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInstanceNew require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInstanceNewResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBackUpDataRequest() (request *DeleteBackUpDataRequest) {
    request = &DeleteBackUpDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwch", APIVersion, "DeleteBackUpData")
    
    
    return
}

func NewDeleteBackUpDataResponse() (response *DeleteBackUpDataResponse) {
    response = &DeleteBackUpDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteBackUpData
// 删除备份数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DeleteBackUpData(request *DeleteBackUpDataRequest) (response *DeleteBackUpDataResponse, err error) {
    return c.DeleteBackUpDataWithContext(context.Background(), request)
}

// DeleteBackUpData
// 删除备份数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DeleteBackUpDataWithContext(ctx context.Context, request *DeleteBackUpDataRequest) (response *DeleteBackUpDataResponse, err error) {
    if request == nil {
        request = NewDeleteBackUpDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteBackUpData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteBackUpDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackUpJobRequest() (request *DescribeBackUpJobRequest) {
    request = &DescribeBackUpJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwch", APIVersion, "DescribeBackUpJob")
    
    
    return
}

func NewDescribeBackUpJobResponse() (response *DescribeBackUpJobResponse) {
    response = &DescribeBackUpJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackUpJob
// 查询备份任务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeBackUpJob(request *DescribeBackUpJobRequest) (response *DescribeBackUpJobResponse, err error) {
    return c.DescribeBackUpJobWithContext(context.Background(), request)
}

// DescribeBackUpJob
// 查询备份任务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeBackUpJobWithContext(ctx context.Context, request *DescribeBackUpJobRequest) (response *DescribeBackUpJobResponse, err error) {
    if request == nil {
        request = NewDescribeBackUpJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackUpJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackUpJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackUpJobDetailRequest() (request *DescribeBackUpJobDetailRequest) {
    request = &DescribeBackUpJobDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwch", APIVersion, "DescribeBackUpJobDetail")
    
    
    return
}

func NewDescribeBackUpJobDetailResponse() (response *DescribeBackUpJobDetailResponse) {
    response = &DescribeBackUpJobDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackUpJobDetail
// 查询备份任务详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeBackUpJobDetail(request *DescribeBackUpJobDetailRequest) (response *DescribeBackUpJobDetailResponse, err error) {
    return c.DescribeBackUpJobDetailWithContext(context.Background(), request)
}

// DescribeBackUpJobDetail
// 查询备份任务详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeBackUpJobDetailWithContext(ctx context.Context, request *DescribeBackUpJobDetailRequest) (response *DescribeBackUpJobDetailResponse, err error) {
    if request == nil {
        request = NewDescribeBackUpJobDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackUpJobDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackUpJobDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackUpScheduleRequest() (request *DescribeBackUpScheduleRequest) {
    request = &DescribeBackUpScheduleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwch", APIVersion, "DescribeBackUpSchedule")
    
    
    return
}

func NewDescribeBackUpScheduleResponse() (response *DescribeBackUpScheduleResponse) {
    response = &DescribeBackUpScheduleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackUpSchedule
// 查询备份策略信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeBackUpSchedule(request *DescribeBackUpScheduleRequest) (response *DescribeBackUpScheduleResponse, err error) {
    return c.DescribeBackUpScheduleWithContext(context.Background(), request)
}

// DescribeBackUpSchedule
// 查询备份策略信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeBackUpScheduleWithContext(ctx context.Context, request *DescribeBackUpScheduleRequest) (response *DescribeBackUpScheduleResponse, err error) {
    if request == nil {
        request = NewDescribeBackUpScheduleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackUpSchedule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackUpScheduleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackUpTablesRequest() (request *DescribeBackUpTablesRequest) {
    request = &DescribeBackUpTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwch", APIVersion, "DescribeBackUpTables")
    
    
    return
}

func NewDescribeBackUpTablesResponse() (response *DescribeBackUpTablesResponse) {
    response = &DescribeBackUpTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackUpTables
// 获取可备份表信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeBackUpTables(request *DescribeBackUpTablesRequest) (response *DescribeBackUpTablesResponse, err error) {
    return c.DescribeBackUpTablesWithContext(context.Background(), request)
}

// DescribeBackUpTables
// 获取可备份表信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeBackUpTablesWithContext(ctx context.Context, request *DescribeBackUpTablesRequest) (response *DescribeBackUpTablesResponse, err error) {
    if request == nil {
        request = NewDescribeBackUpTablesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackUpTables require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackUpTablesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCkSqlApisRequest() (request *DescribeCkSqlApisRequest) {
    request = &DescribeCkSqlApisRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwch", APIVersion, "DescribeCkSqlApis")
    
    
    return
}

func NewDescribeCkSqlApisResponse() (response *DescribeCkSqlApisResponse) {
    response = &DescribeCkSqlApisResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCkSqlApis
// 查询集群用户、集群表，数据库等相关信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeCkSqlApis(request *DescribeCkSqlApisRequest) (response *DescribeCkSqlApisResponse, err error) {
    return c.DescribeCkSqlApisWithContext(context.Background(), request)
}

// DescribeCkSqlApis
// 查询集群用户、集群表，数据库等相关信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeCkSqlApisWithContext(ctx context.Context, request *DescribeCkSqlApisRequest) (response *DescribeCkSqlApisResponse, err error) {
    if request == nil {
        request = NewDescribeCkSqlApisRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCkSqlApis require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCkSqlApisResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterConfigsRequest() (request *DescribeClusterConfigsRequest) {
    request = &DescribeClusterConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwch", APIVersion, "DescribeClusterConfigs")
    
    
    return
}

func NewDescribeClusterConfigsResponse() (response *DescribeClusterConfigsResponse) {
    response = &DescribeClusterConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterConfigs
// 获取集群的最新的几个配置文件（config.xml、metrika.xml、user.xml）的内容，显示给用户
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeClusterConfigs(request *DescribeClusterConfigsRequest) (response *DescribeClusterConfigsResponse, err error) {
    return c.DescribeClusterConfigsWithContext(context.Background(), request)
}

// DescribeClusterConfigs
// 获取集群的最新的几个配置文件（config.xml、metrika.xml、user.xml）的内容，显示给用户
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeClusterConfigsWithContext(ctx context.Context, request *DescribeClusterConfigsRequest) (response *DescribeClusterConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterConfigsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterConfigs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceRequest() (request *DescribeInstanceRequest) {
    request = &DescribeInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwch", APIVersion, "DescribeInstance")
    
    
    return
}

func NewDescribeInstanceResponse() (response *DescribeInstanceResponse) {
    response = &DescribeInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstance
// 根据实例ID查询某个实例的具体信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeInstance(request *DescribeInstanceRequest) (response *DescribeInstanceResponse, err error) {
    return c.DescribeInstanceWithContext(context.Background(), request)
}

// DescribeInstance
// 根据实例ID查询某个实例的具体信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeInstanceWithContext(ctx context.Context, request *DescribeInstanceRequest) (response *DescribeInstanceResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceClustersRequest() (request *DescribeInstanceClustersRequest) {
    request = &DescribeInstanceClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwch", APIVersion, "DescribeInstanceClusters")
    
    
    return
}

func NewDescribeInstanceClustersResponse() (response *DescribeInstanceClustersResponse) {
    response = &DescribeInstanceClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceClusters
// 集群vcluster列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeInstanceClusters(request *DescribeInstanceClustersRequest) (response *DescribeInstanceClustersResponse, err error) {
    return c.DescribeInstanceClustersWithContext(context.Background(), request)
}

// DescribeInstanceClusters
// 集群vcluster列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeInstanceClustersWithContext(ctx context.Context, request *DescribeInstanceClustersRequest) (response *DescribeInstanceClustersResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceClustersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceClusters require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceClustersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceKeyValConfigsRequest() (request *DescribeInstanceKeyValConfigsRequest) {
    request = &DescribeInstanceKeyValConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwch", APIVersion, "DescribeInstanceKeyValConfigs")
    
    
    return
}

func NewDescribeInstanceKeyValConfigsResponse() (response *DescribeInstanceKeyValConfigsResponse) {
    response = &DescribeInstanceKeyValConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceKeyValConfigs
// 在集群详情页面获取所有参数列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeInstanceKeyValConfigs(request *DescribeInstanceKeyValConfigsRequest) (response *DescribeInstanceKeyValConfigsResponse, err error) {
    return c.DescribeInstanceKeyValConfigsWithContext(context.Background(), request)
}

// DescribeInstanceKeyValConfigs
// 在集群详情页面获取所有参数列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeInstanceKeyValConfigsWithContext(ctx context.Context, request *DescribeInstanceKeyValConfigsRequest) (response *DescribeInstanceKeyValConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceKeyValConfigsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceKeyValConfigs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceKeyValConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceNodesRequest() (request *DescribeInstanceNodesRequest) {
    request = &DescribeInstanceNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwch", APIVersion, "DescribeInstanceNodes")
    
    
    return
}

func NewDescribeInstanceNodesResponse() (response *DescribeInstanceNodesResponse) {
    response = &DescribeInstanceNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceNodes
// 获取实例节点信息列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeInstanceNodes(request *DescribeInstanceNodesRequest) (response *DescribeInstanceNodesResponse, err error) {
    return c.DescribeInstanceNodesWithContext(context.Background(), request)
}

// DescribeInstanceNodes
// 获取实例节点信息列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeInstanceNodesWithContext(ctx context.Context, request *DescribeInstanceNodesRequest) (response *DescribeInstanceNodesResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceNodesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceNodesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceShardsRequest() (request *DescribeInstanceShardsRequest) {
    request = &DescribeInstanceShardsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwch", APIVersion, "DescribeInstanceShards")
    
    
    return
}

func NewDescribeInstanceShardsResponse() (response *DescribeInstanceShardsResponse) {
    response = &DescribeInstanceShardsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceShards
// 获取实例shard信息列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeInstanceShards(request *DescribeInstanceShardsRequest) (response *DescribeInstanceShardsResponse, err error) {
    return c.DescribeInstanceShardsWithContext(context.Background(), request)
}

// DescribeInstanceShards
// 获取实例shard信息列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeInstanceShardsWithContext(ctx context.Context, request *DescribeInstanceShardsRequest) (response *DescribeInstanceShardsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceShardsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceShards require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceShardsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceStateRequest() (request *DescribeInstanceStateRequest) {
    request = &DescribeInstanceStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwch", APIVersion, "DescribeInstanceState")
    
    
    return
}

func NewDescribeInstanceStateResponse() (response *DescribeInstanceStateResponse) {
    response = &DescribeInstanceStateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceState
// 集群详情页中显示集群状态、流程进度等
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeInstanceState(request *DescribeInstanceStateRequest) (response *DescribeInstanceStateResponse, err error) {
    return c.DescribeInstanceStateWithContext(context.Background(), request)
}

// DescribeInstanceState
// 集群详情页中显示集群状态、流程进度等
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeInstanceStateWithContext(ctx context.Context, request *DescribeInstanceStateRequest) (response *DescribeInstanceStateResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceStateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceState require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceStateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesNewRequest() (request *DescribeInstancesNewRequest) {
    request = &DescribeInstancesNewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwch", APIVersion, "DescribeInstancesNew")
    
    
    return
}

func NewDescribeInstancesNewResponse() (response *DescribeInstancesNewResponse) {
    response = &DescribeInstancesNewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstancesNew
// 获取实例列表，供外部sdk使用
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeInstancesNew(request *DescribeInstancesNewRequest) (response *DescribeInstancesNewResponse, err error) {
    return c.DescribeInstancesNewWithContext(context.Background(), request)
}

// DescribeInstancesNew
// 获取实例列表，供外部sdk使用
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeInstancesNewWithContext(ctx context.Context, request *DescribeInstancesNewRequest) (response *DescribeInstancesNewResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesNewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstancesNew require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesNewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSpecRequest() (request *DescribeSpecRequest) {
    request = &DescribeSpecRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwch", APIVersion, "DescribeSpec")
    
    
    return
}

func NewDescribeSpecResponse() (response *DescribeSpecResponse) {
    response = &DescribeSpecResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSpec
// 购买页拉取集群的数据节点和zookeeper节点的规格列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeSpec(request *DescribeSpecRequest) (response *DescribeSpecResponse, err error) {
    return c.DescribeSpecWithContext(context.Background(), request)
}

// DescribeSpec
// 购买页拉取集群的数据节点和zookeeper节点的规格列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeSpecWithContext(ctx context.Context, request *DescribeSpecRequest) (response *DescribeSpecResponse, err error) {
    if request == nil {
        request = NewDescribeSpecRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSpec require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSpecResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyInstanceRequest() (request *DestroyInstanceRequest) {
    request = &DestroyInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwch", APIVersion, "DestroyInstance")
    
    
    return
}

func NewDestroyInstanceResponse() (response *DestroyInstanceResponse) {
    response = &DestroyInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DestroyInstance
// 销毁集群 open api
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DestroyInstance(request *DestroyInstanceRequest) (response *DestroyInstanceResponse, err error) {
    return c.DestroyInstanceWithContext(context.Background(), request)
}

// DestroyInstance
// 销毁集群 open api
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DestroyInstanceWithContext(ctx context.Context, request *DestroyInstanceRequest) (response *DestroyInstanceResponse, err error) {
    if request == nil {
        request = NewDestroyInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DestroyInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDestroyInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterConfigsRequest() (request *ModifyClusterConfigsRequest) {
    request = &ModifyClusterConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwch", APIVersion, "ModifyClusterConfigs")
    
    
    return
}

func NewModifyClusterConfigsResponse() (response *ModifyClusterConfigsResponse) {
    response = &ModifyClusterConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyClusterConfigs
// 在集群配置页面修改集群配置文件接口，xml模式
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyClusterConfigs(request *ModifyClusterConfigsRequest) (response *ModifyClusterConfigsResponse, err error) {
    return c.ModifyClusterConfigsWithContext(context.Background(), request)
}

// ModifyClusterConfigs
// 在集群配置页面修改集群配置文件接口，xml模式
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyClusterConfigsWithContext(ctx context.Context, request *ModifyClusterConfigsRequest) (response *ModifyClusterConfigsResponse, err error) {
    if request == nil {
        request = NewModifyClusterConfigsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClusterConfigs require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClusterConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceKeyValConfigsRequest() (request *ModifyInstanceKeyValConfigsRequest) {
    request = &ModifyInstanceKeyValConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwch", APIVersion, "ModifyInstanceKeyValConfigs")
    
    
    return
}

func NewModifyInstanceKeyValConfigsResponse() (response *ModifyInstanceKeyValConfigsResponse) {
    response = &ModifyInstanceKeyValConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstanceKeyValConfigs
// KV模式修改配置接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyInstanceKeyValConfigs(request *ModifyInstanceKeyValConfigsRequest) (response *ModifyInstanceKeyValConfigsResponse, err error) {
    return c.ModifyInstanceKeyValConfigsWithContext(context.Background(), request)
}

// ModifyInstanceKeyValConfigs
// KV模式修改配置接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyInstanceKeyValConfigsWithContext(ctx context.Context, request *ModifyInstanceKeyValConfigsRequest) (response *ModifyInstanceKeyValConfigsResponse, err error) {
    if request == nil {
        request = NewModifyInstanceKeyValConfigsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceKeyValConfigs require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceKeyValConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUserNewPrivilegeRequest() (request *ModifyUserNewPrivilegeRequest) {
    request = &ModifyUserNewPrivilegeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwch", APIVersion, "ModifyUserNewPrivilege")
    
    
    return
}

func NewModifyUserNewPrivilegeResponse() (response *ModifyUserNewPrivilegeResponse) {
    response = &ModifyUserNewPrivilegeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyUserNewPrivilege
// 针对集群账号的权限做管控（新版）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyUserNewPrivilege(request *ModifyUserNewPrivilegeRequest) (response *ModifyUserNewPrivilegeResponse, err error) {
    return c.ModifyUserNewPrivilegeWithContext(context.Background(), request)
}

// ModifyUserNewPrivilege
// 针对集群账号的权限做管控（新版）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyUserNewPrivilegeWithContext(ctx context.Context, request *ModifyUserNewPrivilegeRequest) (response *ModifyUserNewPrivilegeResponse, err error) {
    if request == nil {
        request = NewModifyUserNewPrivilegeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUserNewPrivilege require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUserNewPrivilegeResponse()
    err = c.Send(request, response)
    return
}

func NewOpenBackUpRequest() (request *OpenBackUpRequest) {
    request = &OpenBackUpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwch", APIVersion, "OpenBackUp")
    
    
    return
}

func NewOpenBackUpResponse() (response *OpenBackUpResponse) {
    response = &OpenBackUpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OpenBackUp
// 开启或者关闭策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) OpenBackUp(request *OpenBackUpRequest) (response *OpenBackUpResponse, err error) {
    return c.OpenBackUpWithContext(context.Background(), request)
}

// OpenBackUp
// 开启或者关闭策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) OpenBackUpWithContext(ctx context.Context, request *OpenBackUpRequest) (response *OpenBackUpResponse, err error) {
    if request == nil {
        request = NewOpenBackUpRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenBackUp require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenBackUpResponse()
    err = c.Send(request, response)
    return
}

func NewRecoverBackUpJobRequest() (request *RecoverBackUpJobRequest) {
    request = &RecoverBackUpJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwch", APIVersion, "RecoverBackUpJob")
    
    
    return
}

func NewRecoverBackUpJobResponse() (response *RecoverBackUpJobResponse) {
    response = &RecoverBackUpJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RecoverBackUpJob
// 备份恢复
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) RecoverBackUpJob(request *RecoverBackUpJobRequest) (response *RecoverBackUpJobResponse, err error) {
    return c.RecoverBackUpJobWithContext(context.Background(), request)
}

// RecoverBackUpJob
// 备份恢复
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) RecoverBackUpJobWithContext(ctx context.Context, request *RecoverBackUpJobRequest) (response *RecoverBackUpJobResponse, err error) {
    if request == nil {
        request = NewRecoverBackUpJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecoverBackUpJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecoverBackUpJobResponse()
    err = c.Send(request, response)
    return
}

func NewResizeDiskRequest() (request *ResizeDiskRequest) {
    request = &ResizeDiskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwch", APIVersion, "ResizeDisk")
    
    
    return
}

func NewResizeDiskResponse() (response *ResizeDiskResponse) {
    response = &ResizeDiskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResizeDisk
// 扩容磁盘，包含扩容数据节点，zk节点
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ResizeDisk(request *ResizeDiskRequest) (response *ResizeDiskResponse, err error) {
    return c.ResizeDiskWithContext(context.Background(), request)
}

// ResizeDisk
// 扩容磁盘，包含扩容数据节点，zk节点
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ResizeDiskWithContext(ctx context.Context, request *ResizeDiskRequest) (response *ResizeDiskResponse, err error) {
    if request == nil {
        request = NewResizeDiskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResizeDisk require credential")
    }

    request.SetContext(ctx)
    
    response = NewResizeDiskResponse()
    err = c.Send(request, response)
    return
}

func NewScaleCNOutUpInstanceRequest() (request *ScaleCNOutUpInstanceRequest) {
    request = &ScaleCNOutUpInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwch", APIVersion, "ScaleCNOutUpInstance")
    
    
    return
}

func NewScaleCNOutUpInstanceResponse() (response *ScaleCNOutUpInstanceResponse) {
    response = &ScaleCNOutUpInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ScaleCNOutUpInstance
// open-api接口提供弹性伸缩云原生集群能力
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ScaleCNOutUpInstance(request *ScaleCNOutUpInstanceRequest) (response *ScaleCNOutUpInstanceResponse, err error) {
    return c.ScaleCNOutUpInstanceWithContext(context.Background(), request)
}

// ScaleCNOutUpInstance
// open-api接口提供弹性伸缩云原生集群能力
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ScaleCNOutUpInstanceWithContext(ctx context.Context, request *ScaleCNOutUpInstanceRequest) (response *ScaleCNOutUpInstanceResponse, err error) {
    if request == nil {
        request = NewScaleCNOutUpInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ScaleCNOutUpInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewScaleCNOutUpInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewScaleOutInstanceRequest() (request *ScaleOutInstanceRequest) {
    request = &ScaleOutInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwch", APIVersion, "ScaleOutInstance")
    
    
    return
}

func NewScaleOutInstanceResponse() (response *ScaleOutInstanceResponse) {
    response = &ScaleOutInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ScaleOutInstance
// 调整clickhouse节点数量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ScaleOutInstance(request *ScaleOutInstanceRequest) (response *ScaleOutInstanceResponse, err error) {
    return c.ScaleOutInstanceWithContext(context.Background(), request)
}

// ScaleOutInstance
// 调整clickhouse节点数量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ScaleOutInstanceWithContext(ctx context.Context, request *ScaleOutInstanceRequest) (response *ScaleOutInstanceResponse, err error) {
    if request == nil {
        request = NewScaleOutInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ScaleOutInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewScaleOutInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewScaleUpInstanceRequest() (request *ScaleUpInstanceRequest) {
    request = &ScaleUpInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwch", APIVersion, "ScaleUpInstance")
    
    
    return
}

func NewScaleUpInstanceResponse() (response *ScaleUpInstanceResponse) {
    response = &ScaleUpInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ScaleUpInstance
// 垂直扩缩容节点规格，修改节点cvm的规格cpu，内存。 规格变化阶段，服务不可用。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ScaleUpInstance(request *ScaleUpInstanceRequest) (response *ScaleUpInstanceResponse, err error) {
    return c.ScaleUpInstanceWithContext(context.Background(), request)
}

// ScaleUpInstance
// 垂直扩缩容节点规格，修改节点cvm的规格cpu，内存。 规格变化阶段，服务不可用。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ScaleUpInstanceWithContext(ctx context.Context, request *ScaleUpInstanceRequest) (response *ScaleUpInstanceResponse, err error) {
    if request == nil {
        request = NewScaleUpInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ScaleUpInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewScaleUpInstanceResponse()
    err = c.Send(request, response)
    return
}
