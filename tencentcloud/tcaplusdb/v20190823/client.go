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

package v20190823

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-08-23"

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

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewClearTablesRequest() (request *ClearTablesRequest) {
    request = &ClearTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "ClearTables")
    return
}

func NewClearTablesResponse() (response *ClearTablesResponse) {
    response = &ClearTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据给定的表信息，清除表数据。
func (c *Client) ClearTables(request *ClearTablesRequest) (response *ClearTablesResponse, err error) {
    if request == nil {
        request = NewClearTablesRequest()
    }
    response = NewClearTablesResponse()
    err = c.Send(request, response)
    return
}

func NewCompareIdlFilesRequest() (request *CompareIdlFilesRequest) {
    request = &CompareIdlFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "CompareIdlFiles")
    return
}

func NewCompareIdlFilesResponse() (response *CompareIdlFilesResponse) {
    response = &CompareIdlFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 选中目标表格，上传并校验改表文件，返回是否允许修改表格结构的结果。
func (c *Client) CompareIdlFiles(request *CompareIdlFilesRequest) (response *CompareIdlFilesResponse, err error) {
    if request == nil {
        request = NewCompareIdlFilesRequest()
    }
    response = NewCompareIdlFilesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBackupRequest() (request *CreateBackupRequest) {
    request = &CreateBackupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "CreateBackup")
    return
}

func NewCreateBackupResponse() (response *CreateBackupResponse) {
    response = &CreateBackupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用户创建备份任务
func (c *Client) CreateBackup(request *CreateBackupRequest) (response *CreateBackupResponse, err error) {
    if request == nil {
        request = NewCreateBackupRequest()
    }
    response = NewCreateBackupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClusterRequest() (request *CreateClusterRequest) {
    request = &CreateClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "CreateCluster")
    return
}

func NewCreateClusterResponse() (response *CreateClusterResponse) {
    response = &CreateClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口用于创建TcaplusDB集群
func (c *Client) CreateCluster(request *CreateClusterRequest) (response *CreateClusterResponse, err error) {
    if request == nil {
        request = NewCreateClusterRequest()
    }
    response = NewCreateClusterResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTableGroupRequest() (request *CreateTableGroupRequest) {
    request = &CreateTableGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "CreateTableGroup")
    return
}

func NewCreateTableGroupResponse() (response *CreateTableGroupResponse) {
    response = &CreateTableGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 在TcaplusDB集群下创建表格组
func (c *Client) CreateTableGroup(request *CreateTableGroupRequest) (response *CreateTableGroupResponse, err error) {
    if request == nil {
        request = NewCreateTableGroupRequest()
    }
    response = NewCreateTableGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTablesRequest() (request *CreateTablesRequest) {
    request = &CreateTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "CreateTables")
    return
}

func NewCreateTablesResponse() (response *CreateTablesResponse) {
    response = &CreateTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据选择的IDL文件列表，批量创建表格
func (c *Client) CreateTables(request *CreateTablesRequest) (response *CreateTablesResponse, err error) {
    if request == nil {
        request = NewCreateTablesRequest()
    }
    response = NewCreateTablesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterRequest() (request *DeleteClusterRequest) {
    request = &DeleteClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DeleteCluster")
    return
}

func NewDeleteClusterResponse() (response *DeleteClusterResponse) {
    response = &DeleteClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除TcaplusDB集群，必须在集群所属所有资源（包括表格组，表）都已经释放的情况下才会成功。
func (c *Client) DeleteCluster(request *DeleteClusterRequest) (response *DeleteClusterResponse, err error) {
    if request == nil {
        request = NewDeleteClusterRequest()
    }
    response = NewDeleteClusterResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteIdlFilesRequest() (request *DeleteIdlFilesRequest) {
    request = &DeleteIdlFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DeleteIdlFiles")
    return
}

func NewDeleteIdlFilesResponse() (response *DeleteIdlFilesResponse) {
    response = &DeleteIdlFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 指定集群ID和待删除IDL文件的信息，删除目标文件，如果文件正在被表关联则删除失败。
func (c *Client) DeleteIdlFiles(request *DeleteIdlFilesRequest) (response *DeleteIdlFilesResponse, err error) {
    if request == nil {
        request = NewDeleteIdlFilesRequest()
    }
    response = NewDeleteIdlFilesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTableGroupRequest() (request *DeleteTableGroupRequest) {
    request = &DeleteTableGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DeleteTableGroup")
    return
}

func NewDeleteTableGroupResponse() (response *DeleteTableGroupResponse) {
    response = &DeleteTableGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除表格组
func (c *Client) DeleteTableGroup(request *DeleteTableGroupRequest) (response *DeleteTableGroupResponse, err error) {
    if request == nil {
        request = NewDeleteTableGroupRequest()
    }
    response = NewDeleteTableGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTablesRequest() (request *DeleteTablesRequest) {
    request = &DeleteTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DeleteTables")
    return
}

func NewDeleteTablesResponse() (response *DeleteTablesResponse) {
    response = &DeleteTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据指定的表信息删除目标表
func (c *Client) DeleteTables(request *DeleteTablesRequest) (response *DeleteTablesResponse, err error) {
    if request == nil {
        request = NewDeleteTablesRequest()
    }
    response = NewDeleteTablesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClustersRequest() (request *DescribeClustersRequest) {
    request = &DescribeClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DescribeClusters")
    return
}

func NewDescribeClustersResponse() (response *DescribeClustersResponse) {
    response = &DescribeClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询TcaplusDB集群列表，包含集群详细信息。
func (c *Client) DescribeClusters(request *DescribeClustersRequest) (response *DescribeClustersResponse, err error) {
    if request == nil {
        request = NewDescribeClustersRequest()
    }
    response = NewDescribeClustersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIdlFileInfosRequest() (request *DescribeIdlFileInfosRequest) {
    request = &DescribeIdlFileInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DescribeIdlFileInfos")
    return
}

func NewDescribeIdlFileInfosResponse() (response *DescribeIdlFileInfosResponse) {
    response = &DescribeIdlFileInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询表描述文件详情
func (c *Client) DescribeIdlFileInfos(request *DescribeIdlFileInfosRequest) (response *DescribeIdlFileInfosResponse, err error) {
    if request == nil {
        request = NewDescribeIdlFileInfosRequest()
    }
    response = NewDescribeIdlFileInfosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegionsRequest() (request *DescribeRegionsRequest) {
    request = &DescribeRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DescribeRegions")
    return
}

func NewDescribeRegionsResponse() (response *DescribeRegionsResponse) {
    response = &DescribeRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询TcaplusDB服务支持的地域列表
func (c *Client) DescribeRegions(request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeRegionsRequest()
    }
    response = NewDescribeRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTableGroupsRequest() (request *DescribeTableGroupsRequest) {
    request = &DescribeTableGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DescribeTableGroups")
    return
}

func NewDescribeTableGroupsResponse() (response *DescribeTableGroupsResponse) {
    response = &DescribeTableGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询表格组列表
func (c *Client) DescribeTableGroups(request *DescribeTableGroupsRequest) (response *DescribeTableGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeTableGroupsRequest()
    }
    response = NewDescribeTableGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTablesRequest() (request *DescribeTablesRequest) {
    request = &DescribeTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DescribeTables")
    return
}

func NewDescribeTablesResponse() (response *DescribeTablesResponse) {
    response = &DescribeTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询表详情
func (c *Client) DescribeTables(request *DescribeTablesRequest) (response *DescribeTablesResponse, err error) {
    if request == nil {
        request = NewDescribeTablesRequest()
    }
    response = NewDescribeTablesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTablesInRecycleRequest() (request *DescribeTablesInRecycleRequest) {
    request = &DescribeTablesInRecycleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DescribeTablesInRecycle")
    return
}

func NewDescribeTablesInRecycleResponse() (response *DescribeTablesInRecycleResponse) {
    response = &DescribeTablesInRecycleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询回收站中的表详情
func (c *Client) DescribeTablesInRecycle(request *DescribeTablesInRecycleRequest) (response *DescribeTablesInRecycleResponse, err error) {
    if request == nil {
        request = NewDescribeTablesInRecycleRequest()
    }
    response = NewDescribeTablesInRecycleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTasksRequest() (request *DescribeTasksRequest) {
    request = &DescribeTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DescribeTasks")
    return
}

func NewDescribeTasksResponse() (response *DescribeTasksResponse) {
    response = &DescribeTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询任务列表
func (c *Client) DescribeTasks(request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
    if request == nil {
        request = NewDescribeTasksRequest()
    }
    response = NewDescribeTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUinInWhitelistRequest() (request *DescribeUinInWhitelistRequest) {
    request = &DescribeUinInWhitelistRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DescribeUinInWhitelist")
    return
}

func NewDescribeUinInWhitelistResponse() (response *DescribeUinInWhitelistResponse) {
    response = &DescribeUinInWhitelistResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询本用户是否在白名单中，控制是否能创建TDR类型的APP或表
func (c *Client) DescribeUinInWhitelist(request *DescribeUinInWhitelistRequest) (response *DescribeUinInWhitelistResponse, err error) {
    if request == nil {
        request = NewDescribeUinInWhitelistRequest()
    }
    response = NewDescribeUinInWhitelistResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterNameRequest() (request *ModifyClusterNameRequest) {
    request = &ModifyClusterNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "ModifyClusterName")
    return
}

func NewModifyClusterNameResponse() (response *ModifyClusterNameResponse) {
    response = &ModifyClusterNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改指定的集群名称
func (c *Client) ModifyClusterName(request *ModifyClusterNameRequest) (response *ModifyClusterNameResponse, err error) {
    if request == nil {
        request = NewModifyClusterNameRequest()
    }
    response = NewModifyClusterNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterPasswordRequest() (request *ModifyClusterPasswordRequest) {
    request = &ModifyClusterPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "ModifyClusterPassword")
    return
}

func NewModifyClusterPasswordResponse() (response *ModifyClusterPasswordResponse) {
    response = &ModifyClusterPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改指定集群的密码，后台将在旧密码失效之前同时支持TcaplusDB SDK使用旧密码和新密码访问数据库。在旧密码失效之前不能提交新的密码修改请求，在旧密码失效之后不能提交修改旧密码过期时间的请求。
func (c *Client) ModifyClusterPassword(request *ModifyClusterPasswordRequest) (response *ModifyClusterPasswordResponse, err error) {
    if request == nil {
        request = NewModifyClusterPasswordRequest()
    }
    response = NewModifyClusterPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTableGroupNameRequest() (request *ModifyTableGroupNameRequest) {
    request = &ModifyTableGroupNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "ModifyTableGroupName")
    return
}

func NewModifyTableGroupNameResponse() (response *ModifyTableGroupNameResponse) {
    response = &ModifyTableGroupNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改TcaplusDB表格组名称
func (c *Client) ModifyTableGroupName(request *ModifyTableGroupNameRequest) (response *ModifyTableGroupNameResponse, err error) {
    if request == nil {
        request = NewModifyTableGroupNameRequest()
    }
    response = NewModifyTableGroupNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTableMemosRequest() (request *ModifyTableMemosRequest) {
    request = &ModifyTableMemosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "ModifyTableMemos")
    return
}

func NewModifyTableMemosResponse() (response *ModifyTableMemosResponse) {
    response = &ModifyTableMemosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改表备注信息
func (c *Client) ModifyTableMemos(request *ModifyTableMemosRequest) (response *ModifyTableMemosResponse, err error) {
    if request == nil {
        request = NewModifyTableMemosRequest()
    }
    response = NewModifyTableMemosResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTableQuotasRequest() (request *ModifyTableQuotasRequest) {
    request = &ModifyTableQuotasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "ModifyTableQuotas")
    return
}

func NewModifyTableQuotasResponse() (response *ModifyTableQuotasResponse) {
    response = &ModifyTableQuotasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 表格扩缩容
func (c *Client) ModifyTableQuotas(request *ModifyTableQuotasRequest) (response *ModifyTableQuotasResponse, err error) {
    if request == nil {
        request = NewModifyTableQuotasRequest()
    }
    response = NewModifyTableQuotasResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTablesRequest() (request *ModifyTablesRequest) {
    request = &ModifyTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "ModifyTables")
    return
}

func NewModifyTablesResponse() (response *ModifyTablesResponse) {
    response = &ModifyTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据用户选定的表定义IDL文件，批量修改指定的表
func (c *Client) ModifyTables(request *ModifyTablesRequest) (response *ModifyTablesResponse, err error) {
    if request == nil {
        request = NewModifyTablesRequest()
    }
    response = NewModifyTablesResponse()
    err = c.Send(request, response)
    return
}

func NewRecoverRecycleTablesRequest() (request *RecoverRecycleTablesRequest) {
    request = &RecoverRecycleTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "RecoverRecycleTables")
    return
}

func NewRecoverRecycleTablesResponse() (response *RecoverRecycleTablesResponse) {
    response = &RecoverRecycleTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 恢复回收站中，用户自行删除的表。对欠费待释放的表无效。
func (c *Client) RecoverRecycleTables(request *RecoverRecycleTablesRequest) (response *RecoverRecycleTablesResponse, err error) {
    if request == nil {
        request = NewRecoverRecycleTablesRequest()
    }
    response = NewRecoverRecycleTablesResponse()
    err = c.Send(request, response)
    return
}

func NewRollbackTablesRequest() (request *RollbackTablesRequest) {
    request = &RollbackTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "RollbackTables")
    return
}

func NewRollbackTablesResponse() (response *RollbackTablesResponse) {
    response = &RollbackTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 表格数据回档
func (c *Client) RollbackTables(request *RollbackTablesRequest) (response *RollbackTablesResponse, err error) {
    if request == nil {
        request = NewRollbackTablesRequest()
    }
    response = NewRollbackTablesResponse()
    err = c.Send(request, response)
    return
}

func NewVerifyIdlFilesRequest() (request *VerifyIdlFilesRequest) {
    request = &VerifyIdlFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "VerifyIdlFiles")
    return
}

func NewVerifyIdlFilesResponse() (response *VerifyIdlFilesResponse) {
    response = &VerifyIdlFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 上传并校验创建表格文件，返回校验合法的表格定义
func (c *Client) VerifyIdlFiles(request *VerifyIdlFilesRequest) (response *VerifyIdlFilesResponse, err error) {
    if request == nil {
        request = NewVerifyIdlFilesRequest()
    }
    response = NewVerifyIdlFilesResponse()
    err = c.Send(request, response)
    return
}
