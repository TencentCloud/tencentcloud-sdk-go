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

// 选中目标表，上传并校验改表文件，返回是否允许修改表结构
func (c *Client) CompareIdlFiles(request *CompareIdlFilesRequest) (response *CompareIdlFilesResponse, err error) {
    if request == nil {
        request = NewCompareIdlFilesRequest()
    }
    response = NewCompareIdlFilesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAppRequest() (request *CreateAppRequest) {
    request = &CreateAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "CreateApp")
    return
}

func NewCreateAppResponse() (response *CreateAppResponse) {
    response = &CreateAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口用于创建TcaplusDB应用
func (c *Client) CreateApp(request *CreateAppRequest) (response *CreateAppResponse, err error) {
    if request == nil {
        request = NewCreateAppRequest()
    }
    response = NewCreateAppResponse()
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

// 根据选择的IDL文件列表，批量创建表
func (c *Client) CreateTables(request *CreateTablesRequest) (response *CreateTablesResponse, err error) {
    if request == nil {
        request = NewCreateTablesRequest()
    }
    response = NewCreateTablesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateZoneRequest() (request *CreateZoneRequest) {
    request = &CreateZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "CreateZone")
    return
}

func NewCreateZoneResponse() (response *CreateZoneResponse) {
    response = &CreateZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 在TcaplusDB应用下创建大区
func (c *Client) CreateZone(request *CreateZoneRequest) (response *CreateZoneResponse, err error) {
    if request == nil {
        request = NewCreateZoneRequest()
    }
    response = NewCreateZoneResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAppRequest() (request *DeleteAppRequest) {
    request = &DeleteAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DeleteApp")
    return
}

func NewDeleteAppResponse() (response *DeleteAppResponse) {
    response = &DeleteAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除TcaplusDB应用实例，必须在应用实例所属所有资源（包括大区，表）都已经释放的情况下才会成功。
func (c *Client) DeleteApp(request *DeleteAppRequest) (response *DeleteAppResponse, err error) {
    if request == nil {
        request = NewDeleteAppRequest()
    }
    response = NewDeleteAppResponse()
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

// 指定应用ID和待删除IDL文件的信息，删除目标文件，如果文件正在被表关联则删除失败。
func (c *Client) DeleteIdlFiles(request *DeleteIdlFilesRequest) (response *DeleteIdlFilesResponse, err error) {
    if request == nil {
        request = NewDeleteIdlFilesRequest()
    }
    response = NewDeleteIdlFilesResponse()
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

func NewDeleteZoneRequest() (request *DeleteZoneRequest) {
    request = &DeleteZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DeleteZone")
    return
}

func NewDeleteZoneResponse() (response *DeleteZoneResponse) {
    response = &DeleteZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除大区
func (c *Client) DeleteZone(request *DeleteZoneRequest) (response *DeleteZoneResponse, err error) {
    if request == nil {
        request = NewDeleteZoneRequest()
    }
    response = NewDeleteZoneResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAppsRequest() (request *DescribeAppsRequest) {
    request = &DescribeAppsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DescribeApps")
    return
}

func NewDescribeAppsResponse() (response *DescribeAppsResponse) {
    response = &DescribeAppsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询TcaplusDB应用列表，包含应用详细信息。
func (c *Client) DescribeApps(request *DescribeAppsRequest) (response *DescribeAppsResponse, err error) {
    if request == nil {
        request = NewDescribeAppsRequest()
    }
    response = NewDescribeAppsResponse()
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

func NewDescribeZonesRequest() (request *DescribeZonesRequest) {
    request = &DescribeZonesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "DescribeZones")
    return
}

func NewDescribeZonesResponse() (response *DescribeZonesResponse) {
    response = &DescribeZonesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询大区列表
func (c *Client) DescribeZones(request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
    if request == nil {
        request = NewDescribeZonesRequest()
    }
    response = NewDescribeZonesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAppNameRequest() (request *ModifyAppNameRequest) {
    request = &ModifyAppNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "ModifyAppName")
    return
}

func NewModifyAppNameResponse() (response *ModifyAppNameResponse) {
    response = &ModifyAppNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改指定的应用名称
func (c *Client) ModifyAppName(request *ModifyAppNameRequest) (response *ModifyAppNameResponse, err error) {
    if request == nil {
        request = NewModifyAppNameRequest()
    }
    response = NewModifyAppNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAppPasswordRequest() (request *ModifyAppPasswordRequest) {
    request = &ModifyAppPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "ModifyAppPassword")
    return
}

func NewModifyAppPasswordResponse() (response *ModifyAppPasswordResponse) {
    response = &ModifyAppPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改指定AppInstanceId的实例密码，后台将在旧密码失效之前同时支持TcaplusDB SDK使用旧密码和新密码访问数据库。在旧密码失效之前不能提交新的密码修改请求，在旧密码失效之后不能提交修改旧密码过期时间的请求。
func (c *Client) ModifyAppPassword(request *ModifyAppPasswordRequest) (response *ModifyAppPasswordResponse, err error) {
    if request == nil {
        request = NewModifyAppPasswordRequest()
    }
    response = NewModifyAppPasswordResponse()
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

// 表扩缩容
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

func NewModifyZoneNameRequest() (request *ModifyZoneNameRequest) {
    request = &ModifyZoneNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcaplusdb", APIVersion, "ModifyZoneName")
    return
}

func NewModifyZoneNameResponse() (response *ModifyZoneNameResponse) {
    response = &ModifyZoneNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改TcaplusDB大区名称
func (c *Client) ModifyZoneName(request *ModifyZoneNameRequest) (response *ModifyZoneNameResponse, err error) {
    if request == nil {
        request = NewModifyZoneNameRequest()
    }
    response = NewModifyZoneNameResponse()
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

// 表数据回档
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

// 上传并校验加表文件，返回校验合法的表定义
func (c *Client) VerifyIdlFiles(request *VerifyIdlFilesRequest) (response *VerifyIdlFilesResponse, err error) {
    if request == nil {
        request = NewVerifyIdlFilesRequest()
    }
    response = NewVerifyIdlFilesResponse()
    err = c.Send(request, response)
    return
}
