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

package v20211228

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-12-28"

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


func NewCancelBackupJobRequest() (request *CancelBackupJobRequest) {
    request = &CancelBackupJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "CancelBackupJob")
    
    
    return
}

func NewCancelBackupJobResponse() (response *CancelBackupJobResponse) {
    response = &CancelBackupJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CancelBackupJob
// 取消对应的备份实例任务
func (c *Client) CancelBackupJob(request *CancelBackupJobRequest) (response *CancelBackupJobResponse, err error) {
    return c.CancelBackupJobWithContext(context.Background(), request)
}

// CancelBackupJob
// 取消对应的备份实例任务
func (c *Client) CancelBackupJobWithContext(ctx context.Context, request *CancelBackupJobRequest) (response *CancelBackupJobResponse, err error) {
    if request == nil {
        request = NewCancelBackupJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelBackupJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelBackupJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBackUpScheduleRequest() (request *CreateBackUpScheduleRequest) {
    request = &CreateBackUpScheduleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "CreateBackUpSchedule")
    
    
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
func (c *Client) CreateBackUpSchedule(request *CreateBackUpScheduleRequest) (response *CreateBackUpScheduleResponse, err error) {
    return c.CreateBackUpScheduleWithContext(context.Background(), request)
}

// CreateBackUpSchedule
// 创建或者修改备份策略
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
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "CreateInstanceNew")
    
    
    return
}

func NewCreateInstanceNewResponse() (response *CreateInstanceNewResponse) {
    response = &CreateInstanceNewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateInstanceNew
// 通过API创建集群
func (c *Client) CreateInstanceNew(request *CreateInstanceNewRequest) (response *CreateInstanceNewResponse, err error) {
    return c.CreateInstanceNewWithContext(context.Background(), request)
}

// CreateInstanceNew
// 通过API创建集群
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

func NewCreateWorkloadGroupRequest() (request *CreateWorkloadGroupRequest) {
    request = &CreateWorkloadGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "CreateWorkloadGroup")
    
    
    return
}

func NewCreateWorkloadGroupResponse() (response *CreateWorkloadGroupResponse) {
    response = &CreateWorkloadGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateWorkloadGroup
// 创建资源组
func (c *Client) CreateWorkloadGroup(request *CreateWorkloadGroupRequest) (response *CreateWorkloadGroupResponse, err error) {
    return c.CreateWorkloadGroupWithContext(context.Background(), request)
}

// CreateWorkloadGroup
// 创建资源组
func (c *Client) CreateWorkloadGroupWithContext(ctx context.Context, request *CreateWorkloadGroupRequest) (response *CreateWorkloadGroupResponse, err error) {
    if request == nil {
        request = NewCreateWorkloadGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWorkloadGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWorkloadGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBackUpDataRequest() (request *DeleteBackUpDataRequest) {
    request = &DeleteBackUpDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DeleteBackUpData")
    
    
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
func (c *Client) DeleteBackUpData(request *DeleteBackUpDataRequest) (response *DeleteBackUpDataResponse, err error) {
    return c.DeleteBackUpDataWithContext(context.Background(), request)
}

// DeleteBackUpData
// 删除备份数据
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

func NewDeleteWorkloadGroupRequest() (request *DeleteWorkloadGroupRequest) {
    request = &DeleteWorkloadGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DeleteWorkloadGroup")
    
    
    return
}

func NewDeleteWorkloadGroupResponse() (response *DeleteWorkloadGroupResponse) {
    response = &DeleteWorkloadGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteWorkloadGroup
// 删除资源组
func (c *Client) DeleteWorkloadGroup(request *DeleteWorkloadGroupRequest) (response *DeleteWorkloadGroupResponse, err error) {
    return c.DeleteWorkloadGroupWithContext(context.Background(), request)
}

// DeleteWorkloadGroup
// 删除资源组
func (c *Client) DeleteWorkloadGroupWithContext(ctx context.Context, request *DeleteWorkloadGroupRequest) (response *DeleteWorkloadGroupResponse, err error) {
    if request == nil {
        request = NewDeleteWorkloadGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteWorkloadGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteWorkloadGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAreaRegionRequest() (request *DescribeAreaRegionRequest) {
    request = &DescribeAreaRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeAreaRegion")
    
    
    return
}

func NewDescribeAreaRegionResponse() (response *DescribeAreaRegionResponse) {
    response = &DescribeAreaRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAreaRegion
// 集群列表页上显示地域信息及各个地域的集群总数
func (c *Client) DescribeAreaRegion(request *DescribeAreaRegionRequest) (response *DescribeAreaRegionResponse, err error) {
    return c.DescribeAreaRegionWithContext(context.Background(), request)
}

// DescribeAreaRegion
// 集群列表页上显示地域信息及各个地域的集群总数
func (c *Client) DescribeAreaRegionWithContext(ctx context.Context, request *DescribeAreaRegionRequest) (response *DescribeAreaRegionResponse, err error) {
    if request == nil {
        request = NewDescribeAreaRegionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAreaRegion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAreaRegionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackUpJobRequest() (request *DescribeBackUpJobRequest) {
    request = &DescribeBackUpJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeBackUpJob")
    
    
    return
}

func NewDescribeBackUpJobResponse() (response *DescribeBackUpJobResponse) {
    response = &DescribeBackUpJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackUpJob
// 查询备份实例列表
func (c *Client) DescribeBackUpJob(request *DescribeBackUpJobRequest) (response *DescribeBackUpJobResponse, err error) {
    return c.DescribeBackUpJobWithContext(context.Background(), request)
}

// DescribeBackUpJob
// 查询备份实例列表
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
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeBackUpJobDetail")
    
    
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
func (c *Client) DescribeBackUpJobDetail(request *DescribeBackUpJobDetailRequest) (response *DescribeBackUpJobDetailResponse, err error) {
    return c.DescribeBackUpJobDetailWithContext(context.Background(), request)
}

// DescribeBackUpJobDetail
// 查询备份任务详情
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

func NewDescribeBackUpSchedulesRequest() (request *DescribeBackUpSchedulesRequest) {
    request = &DescribeBackUpSchedulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeBackUpSchedules")
    
    
    return
}

func NewDescribeBackUpSchedulesResponse() (response *DescribeBackUpSchedulesResponse) {
    response = &DescribeBackUpSchedulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackUpSchedules
// 获取备份、迁移的调度任务信息
func (c *Client) DescribeBackUpSchedules(request *DescribeBackUpSchedulesRequest) (response *DescribeBackUpSchedulesResponse, err error) {
    return c.DescribeBackUpSchedulesWithContext(context.Background(), request)
}

// DescribeBackUpSchedules
// 获取备份、迁移的调度任务信息
func (c *Client) DescribeBackUpSchedulesWithContext(ctx context.Context, request *DescribeBackUpSchedulesRequest) (response *DescribeBackUpSchedulesResponse, err error) {
    if request == nil {
        request = NewDescribeBackUpSchedulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackUpSchedules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackUpSchedulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackUpTablesRequest() (request *DescribeBackUpTablesRequest) {
    request = &DescribeBackUpTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeBackUpTables")
    
    
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
func (c *Client) DescribeBackUpTables(request *DescribeBackUpTablesRequest) (response *DescribeBackUpTablesResponse, err error) {
    return c.DescribeBackUpTablesWithContext(context.Background(), request)
}

// DescribeBackUpTables
// 获取可备份表信息
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

func NewDescribeBackUpTaskDetailRequest() (request *DescribeBackUpTaskDetailRequest) {
    request = &DescribeBackUpTaskDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeBackUpTaskDetail")
    
    
    return
}

func NewDescribeBackUpTaskDetailResponse() (response *DescribeBackUpTaskDetailResponse) {
    response = &DescribeBackUpTaskDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackUpTaskDetail
// 查询备份任务进度详情
func (c *Client) DescribeBackUpTaskDetail(request *DescribeBackUpTaskDetailRequest) (response *DescribeBackUpTaskDetailResponse, err error) {
    return c.DescribeBackUpTaskDetailWithContext(context.Background(), request)
}

// DescribeBackUpTaskDetail
// 查询备份任务进度详情
func (c *Client) DescribeBackUpTaskDetailWithContext(ctx context.Context, request *DescribeBackUpTaskDetailRequest) (response *DescribeBackUpTaskDetailResponse, err error) {
    if request == nil {
        request = NewDescribeBackUpTaskDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackUpTaskDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackUpTaskDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterConfigsRequest() (request *DescribeClusterConfigsRequest) {
    request = &DescribeClusterConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeClusterConfigs")
    
    
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
func (c *Client) DescribeClusterConfigs(request *DescribeClusterConfigsRequest) (response *DescribeClusterConfigsResponse, err error) {
    return c.DescribeClusterConfigsWithContext(context.Background(), request)
}

// DescribeClusterConfigs
// 获取集群的最新的几个配置文件（config.xml、metrika.xml、user.xml）的内容，显示给用户
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

func NewDescribeClusterConfigsHistoryRequest() (request *DescribeClusterConfigsHistoryRequest) {
    request = &DescribeClusterConfigsHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeClusterConfigsHistory")
    
    
    return
}

func NewDescribeClusterConfigsHistoryResponse() (response *DescribeClusterConfigsHistoryResponse) {
    response = &DescribeClusterConfigsHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterConfigsHistory
// 获取集群配置文件修改历史
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeClusterConfigsHistory(request *DescribeClusterConfigsHistoryRequest) (response *DescribeClusterConfigsHistoryResponse, err error) {
    return c.DescribeClusterConfigsHistoryWithContext(context.Background(), request)
}

// DescribeClusterConfigsHistory
// 获取集群配置文件修改历史
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeClusterConfigsHistoryWithContext(ctx context.Context, request *DescribeClusterConfigsHistoryRequest) (response *DescribeClusterConfigsHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeClusterConfigsHistoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterConfigsHistory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterConfigsHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatabaseAuditDownloadRequest() (request *DescribeDatabaseAuditDownloadRequest) {
    request = &DescribeDatabaseAuditDownloadRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeDatabaseAuditDownload")
    
    
    return
}

func NewDescribeDatabaseAuditDownloadResponse() (response *DescribeDatabaseAuditDownloadResponse) {
    response = &DescribeDatabaseAuditDownloadResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDatabaseAuditDownload
// 下载数据库审计日志
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeDatabaseAuditDownload(request *DescribeDatabaseAuditDownloadRequest) (response *DescribeDatabaseAuditDownloadResponse, err error) {
    return c.DescribeDatabaseAuditDownloadWithContext(context.Background(), request)
}

// DescribeDatabaseAuditDownload
// 下载数据库审计日志
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeDatabaseAuditDownloadWithContext(ctx context.Context, request *DescribeDatabaseAuditDownloadRequest) (response *DescribeDatabaseAuditDownloadResponse, err error) {
    if request == nil {
        request = NewDescribeDatabaseAuditDownloadRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDatabaseAuditDownload require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDatabaseAuditDownloadResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatabaseAuditRecordsRequest() (request *DescribeDatabaseAuditRecordsRequest) {
    request = &DescribeDatabaseAuditRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeDatabaseAuditRecords")
    
    
    return
}

func NewDescribeDatabaseAuditRecordsResponse() (response *DescribeDatabaseAuditRecordsResponse) {
    response = &DescribeDatabaseAuditRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDatabaseAuditRecords
// 获取数据库审计记录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeDatabaseAuditRecords(request *DescribeDatabaseAuditRecordsRequest) (response *DescribeDatabaseAuditRecordsResponse, err error) {
    return c.DescribeDatabaseAuditRecordsWithContext(context.Background(), request)
}

// DescribeDatabaseAuditRecords
// 获取数据库审计记录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeDatabaseAuditRecordsWithContext(ctx context.Context, request *DescribeDatabaseAuditRecordsRequest) (response *DescribeDatabaseAuditRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeDatabaseAuditRecordsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDatabaseAuditRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDatabaseAuditRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatabaseAuditResourceRequest() (request *DescribeDatabaseAuditResourceRequest) {
    request = &DescribeDatabaseAuditResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeDatabaseAuditResource")
    
    
    return
}

func NewDescribeDatabaseAuditResourceResponse() (response *DescribeDatabaseAuditResourceResponse) {
    response = &DescribeDatabaseAuditResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDatabaseAuditResource
// 数据库审计数据库、用户等
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeDatabaseAuditResource(request *DescribeDatabaseAuditResourceRequest) (response *DescribeDatabaseAuditResourceResponse, err error) {
    return c.DescribeDatabaseAuditResourceWithContext(context.Background(), request)
}

// DescribeDatabaseAuditResource
// 数据库审计数据库、用户等
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeDatabaseAuditResourceWithContext(ctx context.Context, request *DescribeDatabaseAuditResourceRequest) (response *DescribeDatabaseAuditResourceResponse, err error) {
    if request == nil {
        request = NewDescribeDatabaseAuditResourceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDatabaseAuditResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDatabaseAuditResourceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDmsSqlHistoryRequest() (request *DescribeDmsSqlHistoryRequest) {
    request = &DescribeDmsSqlHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeDmsSqlHistory")
    
    
    return
}

func NewDescribeDmsSqlHistoryResponse() (response *DescribeDmsSqlHistoryResponse) {
    response = &DescribeDmsSqlHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDmsSqlHistory
// 查询sql工作区历史运行记录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeDmsSqlHistory(request *DescribeDmsSqlHistoryRequest) (response *DescribeDmsSqlHistoryResponse, err error) {
    return c.DescribeDmsSqlHistoryWithContext(context.Background(), request)
}

// DescribeDmsSqlHistory
// 查询sql工作区历史运行记录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeDmsSqlHistoryWithContext(ctx context.Context, request *DescribeDmsSqlHistoryRequest) (response *DescribeDmsSqlHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeDmsSqlHistoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDmsSqlHistory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDmsSqlHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDorisMetricFilesRequest() (request *DescribeDorisMetricFilesRequest) {
    request = &DescribeDorisMetricFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeDorisMetricFiles")
    
    
    return
}

func NewDescribeDorisMetricFilesResponse() (response *DescribeDorisMetricFilesResponse) {
    response = &DescribeDorisMetricFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDorisMetricFiles
// 展示监控指标文件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeDorisMetricFiles(request *DescribeDorisMetricFilesRequest) (response *DescribeDorisMetricFilesResponse, err error) {
    return c.DescribeDorisMetricFilesWithContext(context.Background(), request)
}

// DescribeDorisMetricFiles
// 展示监控指标文件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeDorisMetricFilesWithContext(ctx context.Context, request *DescribeDorisMetricFilesRequest) (response *DescribeDorisMetricFilesResponse, err error) {
    if request == nil {
        request = NewDescribeDorisMetricFilesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDorisMetricFiles require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDorisMetricFilesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFederationTokenRequest() (request *DescribeFederationTokenRequest) {
    request = &DescribeFederationTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeFederationToken")
    
    
    return
}

func NewDescribeFederationTokenResponse() (response *DescribeFederationTokenResponse) {
    response = &DescribeFederationTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFederationToken
// 获取联合身份临时访问凭证
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeFederationToken(request *DescribeFederationTokenRequest) (response *DescribeFederationTokenResponse, err error) {
    return c.DescribeFederationTokenWithContext(context.Background(), request)
}

// DescribeFederationToken
// 获取联合身份临时访问凭证
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeFederationTokenWithContext(ctx context.Context, request *DescribeFederationTokenRequest) (response *DescribeFederationTokenResponse, err error) {
    if request == nil {
        request = NewDescribeFederationTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFederationToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFederationTokenResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFrontEndRequest() (request *DescribeFrontEndRequest) {
    request = &DescribeFrontEndRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeFrontEnd")
    
    
    return
}

func NewDescribeFrontEndResponse() (response *DescribeFrontEndResponse) {
    response = &DescribeFrontEndResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFrontEnd
// 查询前端内容
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeFrontEnd(request *DescribeFrontEndRequest) (response *DescribeFrontEndResponse, err error) {
    return c.DescribeFrontEndWithContext(context.Background(), request)
}

// DescribeFrontEnd
// 查询前端内容
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeFrontEndWithContext(ctx context.Context, request *DescribeFrontEndRequest) (response *DescribeFrontEndResponse, err error) {
    if request == nil {
        request = NewDescribeFrontEndRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFrontEnd require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFrontEndResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGoodsDetailRequest() (request *DescribeGoodsDetailRequest) {
    request = &DescribeGoodsDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeGoodsDetail")
    
    
    return
}

func NewDescribeGoodsDetailResponse() (response *DescribeGoodsDetailResponse) {
    response = &DescribeGoodsDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGoodsDetail
// 生成计费相关接口的GoodsDetail结构
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeGoodsDetail(request *DescribeGoodsDetailRequest) (response *DescribeGoodsDetailResponse, err error) {
    return c.DescribeGoodsDetailWithContext(context.Background(), request)
}

// DescribeGoodsDetail
// 生成计费相关接口的GoodsDetail结构
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeGoodsDetailWithContext(ctx context.Context, request *DescribeGoodsDetailRequest) (response *DescribeGoodsDetailResponse, err error) {
    if request == nil {
        request = NewDescribeGoodsDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGoodsDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGoodsDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceRequest() (request *DescribeInstanceRequest) {
    request = &DescribeInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeInstance")
    
    
    return
}

func NewDescribeInstanceResponse() (response *DescribeInstanceResponse) {
    response = &DescribeInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstance
// 根据集群ID查询某个集群的具体信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeInstance(request *DescribeInstanceRequest) (response *DescribeInstanceResponse, err error) {
    return c.DescribeInstanceWithContext(context.Background(), request)
}

// DescribeInstance
// 根据集群ID查询某个集群的具体信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
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

func NewDescribeInstanceNodesRequest() (request *DescribeInstanceNodesRequest) {
    request = &DescribeInstanceNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeInstanceNodes")
    
    
    return
}

func NewDescribeInstanceNodesResponse() (response *DescribeInstanceNodesResponse) {
    response = &DescribeInstanceNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceNodes
// 获取集群节点信息列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeInstanceNodes(request *DescribeInstanceNodesRequest) (response *DescribeInstanceNodesResponse, err error) {
    return c.DescribeInstanceNodesWithContext(context.Background(), request)
}

// DescribeInstanceNodes
// 获取集群节点信息列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
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

func NewDescribeInstanceNodesInfoRequest() (request *DescribeInstanceNodesInfoRequest) {
    request = &DescribeInstanceNodesInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeInstanceNodesInfo")
    
    
    return
}

func NewDescribeInstanceNodesInfoResponse() (response *DescribeInstanceNodesInfoResponse) {
    response = &DescribeInstanceNodesInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceNodesInfo
// 获取BE/FE节点角色
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeInstanceNodesInfo(request *DescribeInstanceNodesInfoRequest) (response *DescribeInstanceNodesInfoResponse, err error) {
    return c.DescribeInstanceNodesInfoWithContext(context.Background(), request)
}

// DescribeInstanceNodesInfo
// 获取BE/FE节点角色
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeInstanceNodesInfoWithContext(ctx context.Context, request *DescribeInstanceNodesInfoRequest) (response *DescribeInstanceNodesInfoResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceNodesInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceNodesInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceNodesInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceNodesRoleRequest() (request *DescribeInstanceNodesRoleRequest) {
    request = &DescribeInstanceNodesRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeInstanceNodesRole")
    
    
    return
}

func NewDescribeInstanceNodesRoleResponse() (response *DescribeInstanceNodesRoleResponse) {
    response = &DescribeInstanceNodesRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceNodesRole
// 获取集群节点角色
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeInstanceNodesRole(request *DescribeInstanceNodesRoleRequest) (response *DescribeInstanceNodesRoleResponse, err error) {
    return c.DescribeInstanceNodesRoleWithContext(context.Background(), request)
}

// DescribeInstanceNodesRole
// 获取集群节点角色
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeInstanceNodesRoleWithContext(ctx context.Context, request *DescribeInstanceNodesRoleRequest) (response *DescribeInstanceNodesRoleResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceNodesRoleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceNodesRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceNodesRoleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceOperationsRequest() (request *DescribeInstanceOperationsRequest) {
    request = &DescribeInstanceOperationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeInstanceOperations")
    
    
    return
}

func NewDescribeInstanceOperationsResponse() (response *DescribeInstanceOperationsResponse) {
    response = &DescribeInstanceOperationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceOperations
// 在集群详情页面，拉取该集群的操作
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeInstanceOperations(request *DescribeInstanceOperationsRequest) (response *DescribeInstanceOperationsResponse, err error) {
    return c.DescribeInstanceOperationsWithContext(context.Background(), request)
}

// DescribeInstanceOperations
// 在集群详情页面，拉取该集群的操作
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeInstanceOperationsWithContext(ctx context.Context, request *DescribeInstanceOperationsRequest) (response *DescribeInstanceOperationsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceOperationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceOperations require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceOperationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceStateRequest() (request *DescribeInstanceStateRequest) {
    request = &DescribeInstanceStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeInstanceState")
    
    
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
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeInstanceState(request *DescribeInstanceStateRequest) (response *DescribeInstanceStateResponse, err error) {
    return c.DescribeInstanceStateWithContext(context.Background(), request)
}

// DescribeInstanceState
// 集群详情页中显示集群状态、流程进度等
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
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

func NewDescribeInstanceUsedSubnetsRequest() (request *DescribeInstanceUsedSubnetsRequest) {
    request = &DescribeInstanceUsedSubnetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeInstanceUsedSubnets")
    
    
    return
}

func NewDescribeInstanceUsedSubnetsResponse() (response *DescribeInstanceUsedSubnetsResponse) {
    response = &DescribeInstanceUsedSubnetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceUsedSubnets
// 获取集群已使用子网信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeInstanceUsedSubnets(request *DescribeInstanceUsedSubnetsRequest) (response *DescribeInstanceUsedSubnetsResponse, err error) {
    return c.DescribeInstanceUsedSubnetsWithContext(context.Background(), request)
}

// DescribeInstanceUsedSubnets
// 获取集群已使用子网信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeInstanceUsedSubnetsWithContext(ctx context.Context, request *DescribeInstanceUsedSubnetsRequest) (response *DescribeInstanceUsedSubnetsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceUsedSubnetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceUsedSubnets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceUsedSubnetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeInstances")
    
    
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstances
// 获取集群列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    return c.DescribeInstancesWithContext(context.Background(), request)
}

// DescribeInstances
// 获取集群列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeInstancesWithContext(ctx context.Context, request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesHealthStateRequest() (request *DescribeInstancesHealthStateRequest) {
    request = &DescribeInstancesHealthStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeInstancesHealthState")
    
    
    return
}

func NewDescribeInstancesHealthStateResponse() (response *DescribeInstancesHealthStateResponse) {
    response = &DescribeInstancesHealthStateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstancesHealthState
// 集群健康检查
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeInstancesHealthState(request *DescribeInstancesHealthStateRequest) (response *DescribeInstancesHealthStateResponse, err error) {
    return c.DescribeInstancesHealthStateWithContext(context.Background(), request)
}

// DescribeInstancesHealthState
// 集群健康检查
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeInstancesHealthStateWithContext(ctx context.Context, request *DescribeInstancesHealthStateRequest) (response *DescribeInstancesHealthStateResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesHealthStateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstancesHealthState require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesHealthStateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegionZoneRequest() (request *DescribeRegionZoneRequest) {
    request = &DescribeRegionZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeRegionZone")
    
    
    return
}

func NewDescribeRegionZoneResponse() (response *DescribeRegionZoneResponse) {
    response = &DescribeRegionZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRegionZone
// 购买页获取地域及可用区列表、内核版本、网络规则等
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeRegionZone(request *DescribeRegionZoneRequest) (response *DescribeRegionZoneResponse, err error) {
    return c.DescribeRegionZoneWithContext(context.Background(), request)
}

// DescribeRegionZone
// 购买页获取地域及可用区列表、内核版本、网络规则等
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeRegionZoneWithContext(ctx context.Context, request *DescribeRegionZoneRequest) (response *DescribeRegionZoneResponse, err error) {
    if request == nil {
        request = NewDescribeRegionZoneRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRegionZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRegionZoneResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReplicaVersionRequest() (request *DescribeReplicaVersionRequest) {
    request = &DescribeReplicaVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeReplicaVersion")
    
    
    return
}

func NewDescribeReplicaVersionResponse() (response *DescribeReplicaVersionResponse) {
    response = &DescribeReplicaVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeReplicaVersion
// 检查内核版本是否支持新的备份恢复语法
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeReplicaVersion(request *DescribeReplicaVersionRequest) (response *DescribeReplicaVersionResponse, err error) {
    return c.DescribeReplicaVersionWithContext(context.Background(), request)
}

// DescribeReplicaVersion
// 检查内核版本是否支持新的备份恢复语法
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeReplicaVersionWithContext(ctx context.Context, request *DescribeReplicaVersionRequest) (response *DescribeReplicaVersionResponse, err error) {
    if request == nil {
        request = NewDescribeReplicaVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReplicaVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReplicaVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRestoreTaskDetailRequest() (request *DescribeRestoreTaskDetailRequest) {
    request = &DescribeRestoreTaskDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeRestoreTaskDetail")
    
    
    return
}

func NewDescribeRestoreTaskDetailResponse() (response *DescribeRestoreTaskDetailResponse) {
    response = &DescribeRestoreTaskDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRestoreTaskDetail
// 查询恢复任务进度详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeRestoreTaskDetail(request *DescribeRestoreTaskDetailRequest) (response *DescribeRestoreTaskDetailResponse, err error) {
    return c.DescribeRestoreTaskDetailWithContext(context.Background(), request)
}

// DescribeRestoreTaskDetail
// 查询恢复任务进度详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeRestoreTaskDetailWithContext(ctx context.Context, request *DescribeRestoreTaskDetailRequest) (response *DescribeRestoreTaskDetailResponse, err error) {
    if request == nil {
        request = NewDescribeRestoreTaskDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRestoreTaskDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRestoreTaskDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSlowQueryRecordsRequest() (request *DescribeSlowQueryRecordsRequest) {
    request = &DescribeSlowQueryRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeSlowQueryRecords")
    
    
    return
}

func NewDescribeSlowQueryRecordsResponse() (response *DescribeSlowQueryRecordsResponse) {
    response = &DescribeSlowQueryRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSlowQueryRecords
// 获取慢查询列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeSlowQueryRecords(request *DescribeSlowQueryRecordsRequest) (response *DescribeSlowQueryRecordsResponse, err error) {
    return c.DescribeSlowQueryRecordsWithContext(context.Background(), request)
}

// DescribeSlowQueryRecords
// 获取慢查询列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeSlowQueryRecordsWithContext(ctx context.Context, request *DescribeSlowQueryRecordsRequest) (response *DescribeSlowQueryRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeSlowQueryRecordsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSlowQueryRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSlowQueryRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSlowQueryRecordsDownloadRequest() (request *DescribeSlowQueryRecordsDownloadRequest) {
    request = &DescribeSlowQueryRecordsDownloadRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeSlowQueryRecordsDownload")
    
    
    return
}

func NewDescribeSlowQueryRecordsDownloadResponse() (response *DescribeSlowQueryRecordsDownloadResponse) {
    response = &DescribeSlowQueryRecordsDownloadResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSlowQueryRecordsDownload
// 下载慢查询文件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeSlowQueryRecordsDownload(request *DescribeSlowQueryRecordsDownloadRequest) (response *DescribeSlowQueryRecordsDownloadResponse, err error) {
    return c.DescribeSlowQueryRecordsDownloadWithContext(context.Background(), request)
}

// DescribeSlowQueryRecordsDownload
// 下载慢查询文件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeSlowQueryRecordsDownloadWithContext(ctx context.Context, request *DescribeSlowQueryRecordsDownloadRequest) (response *DescribeSlowQueryRecordsDownloadResponse, err error) {
    if request == nil {
        request = NewDescribeSlowQueryRecordsDownloadRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSlowQueryRecordsDownload require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSlowQueryRecordsDownloadResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSpecRequest() (request *DescribeSpecRequest) {
    request = &DescribeSpecRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeSpec")
    
    
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

func NewDescribeSqlApisRequest() (request *DescribeSqlApisRequest) {
    request = &DescribeSqlApisRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeSqlApis")
    
    
    return
}

func NewDescribeSqlApisResponse() (response *DescribeSqlApisResponse) {
    response = &DescribeSqlApisResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSqlApis
// 针对驱动sql命令查询ck集群接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeSqlApis(request *DescribeSqlApisRequest) (response *DescribeSqlApisResponse, err error) {
    return c.DescribeSqlApisWithContext(context.Background(), request)
}

// DescribeSqlApis
// 针对驱动sql命令查询ck集群接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeSqlApisWithContext(ctx context.Context, request *DescribeSqlApisRequest) (response *DescribeSqlApisResponse, err error) {
    if request == nil {
        request = NewDescribeSqlApisRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSqlApis require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSqlApisResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserBindWorkloadGroupRequest() (request *DescribeUserBindWorkloadGroupRequest) {
    request = &DescribeUserBindWorkloadGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeUserBindWorkloadGroup")
    
    
    return
}

func NewDescribeUserBindWorkloadGroupResponse() (response *DescribeUserBindWorkloadGroupResponse) {
    response = &DescribeUserBindWorkloadGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserBindWorkloadGroup
// 获取当前集群各用户绑定的资源信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeUserBindWorkloadGroup(request *DescribeUserBindWorkloadGroupRequest) (response *DescribeUserBindWorkloadGroupResponse, err error) {
    return c.DescribeUserBindWorkloadGroupWithContext(context.Background(), request)
}

// DescribeUserBindWorkloadGroup
// 获取当前集群各用户绑定的资源信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeUserBindWorkloadGroupWithContext(ctx context.Context, request *DescribeUserBindWorkloadGroupRequest) (response *DescribeUserBindWorkloadGroupResponse, err error) {
    if request == nil {
        request = NewDescribeUserBindWorkloadGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserBindWorkloadGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserBindWorkloadGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWorkloadGroupRequest() (request *DescribeWorkloadGroupRequest) {
    request = &DescribeWorkloadGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DescribeWorkloadGroup")
    
    
    return
}

func NewDescribeWorkloadGroupResponse() (response *DescribeWorkloadGroupResponse) {
    response = &DescribeWorkloadGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWorkloadGroup
// 获取资源组信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeWorkloadGroup(request *DescribeWorkloadGroupRequest) (response *DescribeWorkloadGroupResponse, err error) {
    return c.DescribeWorkloadGroupWithContext(context.Background(), request)
}

// DescribeWorkloadGroup
// 获取资源组信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeWorkloadGroupWithContext(ctx context.Context, request *DescribeWorkloadGroupRequest) (response *DescribeWorkloadGroupResponse, err error) {
    if request == nil {
        request = NewDescribeWorkloadGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWorkloadGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWorkloadGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyInstanceRequest() (request *DestroyInstanceRequest) {
    request = &DestroyInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "DestroyInstance")
    
    
    return
}

func NewDestroyInstanceResponse() (response *DestroyInstanceResponse) {
    response = &DestroyInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DestroyInstance
// 销毁集群
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DestroyInstance(request *DestroyInstanceRequest) (response *DestroyInstanceResponse, err error) {
    return c.DestroyInstanceWithContext(context.Background(), request)
}

// DestroyInstance
// 销毁集群
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

func NewFitClsLogRequest() (request *FitClsLogRequest) {
    request = &FitClsLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "FitClsLog")
    
    
    return
}

func NewFitClsLogResponse() (response *FitClsLogResponse) {
    response = &FitClsLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// FitClsLog
// 给已存在集群，配置日志服务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) FitClsLog(request *FitClsLogRequest) (response *FitClsLogResponse, err error) {
    return c.FitClsLogWithContext(context.Background(), request)
}

// FitClsLog
// 给已存在集群，配置日志服务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) FitClsLogWithContext(ctx context.Context, request *FitClsLogRequest) (response *FitClsLogResponse, err error) {
    if request == nil {
        request = NewFitClsLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("FitClsLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewFitClsLogResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceRequest() (request *ModifyInstanceRequest) {
    request = &ModifyInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "ModifyInstance")
    
    
    return
}

func NewModifyInstanceResponse() (response *ModifyInstanceResponse) {
    response = &ModifyInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstance
// 修改集群名称
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyInstance(request *ModifyInstanceRequest) (response *ModifyInstanceResponse, err error) {
    return c.ModifyInstanceWithContext(context.Background(), request)
}

// ModifyInstance
// 修改集群名称
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyInstanceWithContext(ctx context.Context, request *ModifyInstanceRequest) (response *ModifyInstanceResponse, err error) {
    if request == nil {
        request = NewModifyInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceKeyValConfigsRequest() (request *ModifyInstanceKeyValConfigsRequest) {
    request = &ModifyInstanceKeyValConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "ModifyInstanceKeyValConfigs")
    
    
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

func NewModifyNodeStatusRequest() (request *ModifyNodeStatusRequest) {
    request = &ModifyNodeStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "ModifyNodeStatus")
    
    
    return
}

func NewModifyNodeStatusResponse() (response *ModifyNodeStatusResponse) {
    response = &ModifyNodeStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyNodeStatus
// 修改节点状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyNodeStatus(request *ModifyNodeStatusRequest) (response *ModifyNodeStatusResponse, err error) {
    return c.ModifyNodeStatusWithContext(context.Background(), request)
}

// ModifyNodeStatus
// 修改节点状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyNodeStatusWithContext(ctx context.Context, request *ModifyNodeStatusRequest) (response *ModifyNodeStatusResponse, err error) {
    if request == nil {
        request = NewModifyNodeStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNodeStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNodeStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifySecurityGroupsRequest() (request *ModifySecurityGroupsRequest) {
    request = &ModifySecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "ModifySecurityGroups")
    
    
    return
}

func NewModifySecurityGroupsResponse() (response *ModifySecurityGroupsResponse) {
    response = &ModifySecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySecurityGroups
// 更改安全组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifySecurityGroups(request *ModifySecurityGroupsRequest) (response *ModifySecurityGroupsResponse, err error) {
    return c.ModifySecurityGroupsWithContext(context.Background(), request)
}

// ModifySecurityGroups
// 更改安全组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifySecurityGroupsWithContext(ctx context.Context, request *ModifySecurityGroupsRequest) (response *ModifySecurityGroupsResponse, err error) {
    if request == nil {
        request = NewModifySecurityGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySecurityGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUserBindWorkloadGroupRequest() (request *ModifyUserBindWorkloadGroupRequest) {
    request = &ModifyUserBindWorkloadGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "ModifyUserBindWorkloadGroup")
    
    
    return
}

func NewModifyUserBindWorkloadGroupResponse() (response *ModifyUserBindWorkloadGroupResponse) {
    response = &ModifyUserBindWorkloadGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyUserBindWorkloadGroup
// 修改用户绑定的资源组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyUserBindWorkloadGroup(request *ModifyUserBindWorkloadGroupRequest) (response *ModifyUserBindWorkloadGroupResponse, err error) {
    return c.ModifyUserBindWorkloadGroupWithContext(context.Background(), request)
}

// ModifyUserBindWorkloadGroup
// 修改用户绑定的资源组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyUserBindWorkloadGroupWithContext(ctx context.Context, request *ModifyUserBindWorkloadGroupRequest) (response *ModifyUserBindWorkloadGroupResponse, err error) {
    if request == nil {
        request = NewModifyUserBindWorkloadGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUserBindWorkloadGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUserBindWorkloadGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUserPrivilegesV3Request() (request *ModifyUserPrivilegesV3Request) {
    request = &ModifyUserPrivilegesV3Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "ModifyUserPrivilegesV3")
    
    
    return
}

func NewModifyUserPrivilegesV3Response() (response *ModifyUserPrivilegesV3Response) {
    response = &ModifyUserPrivilegesV3Response{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyUserPrivilegesV3
// 修改用户权限,支持catalog，全部db，部分db表三种权限设置类别
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyUserPrivilegesV3(request *ModifyUserPrivilegesV3Request) (response *ModifyUserPrivilegesV3Response, err error) {
    return c.ModifyUserPrivilegesV3WithContext(context.Background(), request)
}

// ModifyUserPrivilegesV3
// 修改用户权限,支持catalog，全部db，部分db表三种权限设置类别
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyUserPrivilegesV3WithContext(ctx context.Context, request *ModifyUserPrivilegesV3Request) (response *ModifyUserPrivilegesV3Response, err error) {
    if request == nil {
        request = NewModifyUserPrivilegesV3Request()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUserPrivilegesV3 require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUserPrivilegesV3Response()
    err = c.Send(request, response)
    return
}

func NewModifyWorkloadGroupRequest() (request *ModifyWorkloadGroupRequest) {
    request = &ModifyWorkloadGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "ModifyWorkloadGroup")
    
    
    return
}

func NewModifyWorkloadGroupResponse() (response *ModifyWorkloadGroupResponse) {
    response = &ModifyWorkloadGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyWorkloadGroup
// 修改资源组信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyWorkloadGroup(request *ModifyWorkloadGroupRequest) (response *ModifyWorkloadGroupResponse, err error) {
    return c.ModifyWorkloadGroupWithContext(context.Background(), request)
}

// ModifyWorkloadGroup
// 修改资源组信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyWorkloadGroupWithContext(ctx context.Context, request *ModifyWorkloadGroupRequest) (response *ModifyWorkloadGroupResponse, err error) {
    if request == nil {
        request = NewModifyWorkloadGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyWorkloadGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyWorkloadGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWorkloadGroupStatusRequest() (request *ModifyWorkloadGroupStatusRequest) {
    request = &ModifyWorkloadGroupStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "ModifyWorkloadGroupStatus")
    
    
    return
}

func NewModifyWorkloadGroupStatusResponse() (response *ModifyWorkloadGroupStatusResponse) {
    response = &ModifyWorkloadGroupStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyWorkloadGroupStatus
// 开启或关闭资源组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyWorkloadGroupStatus(request *ModifyWorkloadGroupStatusRequest) (response *ModifyWorkloadGroupStatusResponse, err error) {
    return c.ModifyWorkloadGroupStatusWithContext(context.Background(), request)
}

// ModifyWorkloadGroupStatus
// 开启或关闭资源组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ModifyWorkloadGroupStatusWithContext(ctx context.Context, request *ModifyWorkloadGroupStatusRequest) (response *ModifyWorkloadGroupStatusResponse, err error) {
    if request == nil {
        request = NewModifyWorkloadGroupStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyWorkloadGroupStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyWorkloadGroupStatusResponse()
    err = c.Send(request, response)
    return
}

func NewOpenBackUpRequest() (request *OpenBackUpRequest) {
    request = &OpenBackUpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "OpenBackUp")
    
    
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
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "RecoverBackUpJob")
    
    
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

func NewReduceInstanceRequest() (request *ReduceInstanceRequest) {
    request = &ReduceInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "ReduceInstance")
    
    
    return
}

func NewReduceInstanceResponse() (response *ReduceInstanceResponse) {
    response = &ReduceInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReduceInstance
// 集群缩容
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ReduceInstance(request *ReduceInstanceRequest) (response *ReduceInstanceResponse, err error) {
    return c.ReduceInstanceWithContext(context.Background(), request)
}

// ReduceInstance
// 集群缩容
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ReduceInstanceWithContext(ctx context.Context, request *ReduceInstanceRequest) (response *ReduceInstanceResponse, err error) {
    if request == nil {
        request = NewReduceInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReduceInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewReduceInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewResizeDiskRequest() (request *ResizeDiskRequest) {
    request = &ResizeDiskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "ResizeDisk")
    
    
    return
}

func NewResizeDiskResponse() (response *ResizeDiskResponse) {
    response = &ResizeDiskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResizeDisk
// 扩容云盘
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ResizeDisk(request *ResizeDiskRequest) (response *ResizeDiskResponse, err error) {
    return c.ResizeDiskWithContext(context.Background(), request)
}

// ResizeDisk
// 扩容云盘
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

func NewRestartClusterForConfigsRequest() (request *RestartClusterForConfigsRequest) {
    request = &RestartClusterForConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "RestartClusterForConfigs")
    
    
    return
}

func NewRestartClusterForConfigsResponse() (response *RestartClusterForConfigsResponse) {
    response = &RestartClusterForConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RestartClusterForConfigs
// 重启集群让配置文件生效
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) RestartClusterForConfigs(request *RestartClusterForConfigsRequest) (response *RestartClusterForConfigsResponse, err error) {
    return c.RestartClusterForConfigsWithContext(context.Background(), request)
}

// RestartClusterForConfigs
// 重启集群让配置文件生效
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) RestartClusterForConfigsWithContext(ctx context.Context, request *RestartClusterForConfigsRequest) (response *RestartClusterForConfigsResponse, err error) {
    if request == nil {
        request = NewRestartClusterForConfigsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestartClusterForConfigs require credential")
    }

    request.SetContext(ctx)
    
    response = NewRestartClusterForConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewRestartClusterForNodeRequest() (request *RestartClusterForNodeRequest) {
    request = &RestartClusterForNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "RestartClusterForNode")
    
    
    return
}

func NewRestartClusterForNodeResponse() (response *RestartClusterForNodeResponse) {
    response = &RestartClusterForNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RestartClusterForNode
// 集群滚动重启
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) RestartClusterForNode(request *RestartClusterForNodeRequest) (response *RestartClusterForNodeResponse, err error) {
    return c.RestartClusterForNodeWithContext(context.Background(), request)
}

// RestartClusterForNode
// 集群滚动重启
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) RestartClusterForNodeWithContext(ctx context.Context, request *RestartClusterForNodeRequest) (response *RestartClusterForNodeResponse, err error) {
    if request == nil {
        request = NewRestartClusterForNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestartClusterForNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewRestartClusterForNodeResponse()
    err = c.Send(request, response)
    return
}

func NewScaleOutInstanceRequest() (request *ScaleOutInstanceRequest) {
    request = &ScaleOutInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "ScaleOutInstance")
    
    
    return
}

func NewScaleOutInstanceResponse() (response *ScaleOutInstanceResponse) {
    response = &ScaleOutInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ScaleOutInstance
// 水平扩容节点
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ScaleOutInstance(request *ScaleOutInstanceRequest) (response *ScaleOutInstanceResponse, err error) {
    return c.ScaleOutInstanceWithContext(context.Background(), request)
}

// ScaleOutInstance
// 水平扩容节点
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
    
    request.Init().WithApiInfo("cdwdoris", APIVersion, "ScaleUpInstance")
    
    
    return
}

func NewScaleUpInstanceResponse() (response *ScaleUpInstanceResponse) {
    response = &ScaleUpInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ScaleUpInstance
// 计算资源垂直变配
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) ScaleUpInstance(request *ScaleUpInstanceRequest) (response *ScaleUpInstanceResponse, err error) {
    return c.ScaleUpInstanceWithContext(context.Background(), request)
}

// ScaleUpInstance
// 计算资源垂直变配
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
