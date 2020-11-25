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

package v20180328

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-03-28"

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


func NewAssociateSecurityGroupsRequest() (request *AssociateSecurityGroupsRequest) {
    request = &AssociateSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "AssociateSecurityGroups")
    return
}

func NewAssociateSecurityGroupsResponse() (response *AssociateSecurityGroupsResponse) {
    response = &AssociateSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(AssociateSecurityGroups)用于安全组批量绑定实例。
func (c *Client) AssociateSecurityGroups(request *AssociateSecurityGroupsRequest) (response *AssociateSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewAssociateSecurityGroupsRequest()
    }
    response = NewAssociateSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewCompleteExpansionRequest() (request *CompleteExpansionRequest) {
    request = &CompleteExpansionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "CompleteExpansion")
    return
}

func NewCompleteExpansionResponse() (response *CompleteExpansionResponse) {
    response = &CompleteExpansionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CompleteExpansion）在实例发起扩容后，实例状态处于“升级待切换”时，可立即完成实例升级切换操作，无需等待可维护时间窗。本接口需要在实例低峰时调用，在完全切换成功前，存在部分库不可访问的风险。
func (c *Client) CompleteExpansion(request *CompleteExpansionRequest) (response *CompleteExpansionResponse, err error) {
    if request == nil {
        request = NewCompleteExpansionRequest()
    }
    response = NewCompleteExpansionResponse()
    err = c.Send(request, response)
    return
}

func NewCompleteMigrationRequest() (request *CompleteMigrationRequest) {
    request = &CompleteMigrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "CompleteMigration")
    return
}

func NewCompleteMigrationResponse() (response *CompleteMigrationResponse) {
    response = &CompleteMigrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CompleteMigration）作用是完成一个迁移任务
func (c *Client) CompleteMigration(request *CompleteMigrationRequest) (response *CompleteMigrationResponse, err error) {
    if request == nil {
        request = NewCompleteMigrationRequest()
    }
    response = NewCompleteMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAccountRequest() (request *CreateAccountRequest) {
    request = &CreateAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "CreateAccount")
    return
}

func NewCreateAccountResponse() (response *CreateAccountResponse) {
    response = &CreateAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateAccount）用于创建实例账号
func (c *Client) CreateAccount(request *CreateAccountRequest) (response *CreateAccountResponse, err error) {
    if request == nil {
        request = NewCreateAccountRequest()
    }
    response = NewCreateAccountResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBackupRequest() (request *CreateBackupRequest) {
    request = &CreateBackupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "CreateBackup")
    return
}

func NewCreateBackupResponse() (response *CreateBackupResponse) {
    response = &CreateBackupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(CreateBackup)用于创建备份。
func (c *Client) CreateBackup(request *CreateBackupRequest) (response *CreateBackupResponse, err error) {
    if request == nil {
        request = NewCreateBackupRequest()
    }
    response = NewCreateBackupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBasicDBInstancesRequest() (request *CreateBasicDBInstancesRequest) {
    request = &CreateBasicDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "CreateBasicDBInstances")
    return
}

func NewCreateBasicDBInstancesResponse() (response *CreateBasicDBInstancesResponse) {
    response = &CreateBasicDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateBasicDBInstances）用于创建SQL server基础版实例。
func (c *Client) CreateBasicDBInstances(request *CreateBasicDBInstancesRequest) (response *CreateBasicDBInstancesResponse, err error) {
    if request == nil {
        request = NewCreateBasicDBInstancesRequest()
    }
    response = NewCreateBasicDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDBRequest() (request *CreateDBRequest) {
    request = &CreateDBRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "CreateDB")
    return
}

func NewCreateDBResponse() (response *CreateDBResponse) {
    response = &CreateDBResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateDB）用于创建数据库。
func (c *Client) CreateDB(request *CreateDBRequest) (response *CreateDBResponse, err error) {
    if request == nil {
        request = NewCreateDBRequest()
    }
    response = NewCreateDBResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDBInstancesRequest() (request *CreateDBInstancesRequest) {
    request = &CreateDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "CreateDBInstances")
    return
}

func NewCreateDBInstancesResponse() (response *CreateDBInstancesResponse) {
    response = &CreateDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateDBInstances）用于创建实例。
func (c *Client) CreateDBInstances(request *CreateDBInstancesRequest) (response *CreateDBInstancesResponse, err error) {
    if request == nil {
        request = NewCreateDBInstancesRequest()
    }
    response = NewCreateDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMigrationRequest() (request *CreateMigrationRequest) {
    request = &CreateMigrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "CreateMigration")
    return
}

func NewCreateMigrationResponse() (response *CreateMigrationResponse) {
    response = &CreateMigrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateMigration）作用是创建一个迁移任务
func (c *Client) CreateMigration(request *CreateMigrationRequest) (response *CreateMigrationResponse, err error) {
    if request == nil {
        request = NewCreateMigrationRequest()
    }
    response = NewCreateMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePublishSubscribeRequest() (request *CreatePublishSubscribeRequest) {
    request = &CreatePublishSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "CreatePublishSubscribe")
    return
}

func NewCreatePublishSubscribeResponse() (response *CreatePublishSubscribeResponse) {
    response = &CreatePublishSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreatePublishSubscribe）用于创建两个数据库之间的发布订阅关系。作为订阅者，不能再充当发布者，作为发布者可以有多个订阅者实例。
func (c *Client) CreatePublishSubscribe(request *CreatePublishSubscribeRequest) (response *CreatePublishSubscribeResponse, err error) {
    if request == nil {
        request = NewCreatePublishSubscribeRequest()
    }
    response = NewCreatePublishSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateReadOnlyDBInstancesRequest() (request *CreateReadOnlyDBInstancesRequest) {
    request = &CreateReadOnlyDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "CreateReadOnlyDBInstances")
    return
}

func NewCreateReadOnlyDBInstancesResponse() (response *CreateReadOnlyDBInstancesResponse) {
    response = &CreateReadOnlyDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateReadOnlyDBInstances）用于添加只读副本实例。
func (c *Client) CreateReadOnlyDBInstances(request *CreateReadOnlyDBInstancesRequest) (response *CreateReadOnlyDBInstancesResponse, err error) {
    if request == nil {
        request = NewCreateReadOnlyDBInstancesRequest()
    }
    response = NewCreateReadOnlyDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAccountRequest() (request *DeleteAccountRequest) {
    request = &DeleteAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DeleteAccount")
    return
}

func NewDeleteAccountResponse() (response *DeleteAccountResponse) {
    response = &DeleteAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeleteAccount）用于删除实例账号。
func (c *Client) DeleteAccount(request *DeleteAccountRequest) (response *DeleteAccountResponse, err error) {
    if request == nil {
        request = NewDeleteAccountRequest()
    }
    response = NewDeleteAccountResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDBRequest() (request *DeleteDBRequest) {
    request = &DeleteDBRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DeleteDB")
    return
}

func NewDeleteDBResponse() (response *DeleteDBResponse) {
    response = &DeleteDBResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DeleteDB)用于删除数据库。
func (c *Client) DeleteDB(request *DeleteDBRequest) (response *DeleteDBResponse, err error) {
    if request == nil {
        request = NewDeleteDBRequest()
    }
    response = NewDeleteDBResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDBInstanceRequest() (request *DeleteDBInstanceRequest) {
    request = &DeleteDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DeleteDBInstance")
    return
}

func NewDeleteDBInstanceResponse() (response *DeleteDBInstanceResponse) {
    response = &DeleteDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeleteDBInstance）用于释放回收站中的SQL server实例。释放后的实例将保存一段时间后物理销毁。其发布订阅将自动解除，其ro副本将自动释放。
func (c *Client) DeleteDBInstance(request *DeleteDBInstanceRequest) (response *DeleteDBInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteDBInstanceRequest()
    }
    response = NewDeleteDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMigrationRequest() (request *DeleteMigrationRequest) {
    request = &DeleteMigrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DeleteMigration")
    return
}

func NewDeleteMigrationResponse() (response *DeleteMigrationResponse) {
    response = &DeleteMigrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeleteMigration）用于删除迁移任务
func (c *Client) DeleteMigration(request *DeleteMigrationRequest) (response *DeleteMigrationResponse, err error) {
    if request == nil {
        request = NewDeleteMigrationRequest()
    }
    response = NewDeleteMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePublishSubscribeRequest() (request *DeletePublishSubscribeRequest) {
    request = &DeletePublishSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DeletePublishSubscribe")
    return
}

func NewDeletePublishSubscribeResponse() (response *DeletePublishSubscribeResponse) {
    response = &DeletePublishSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeletePublishSubscribe）用于删除两个数据库间的发布订阅关系。
func (c *Client) DeletePublishSubscribe(request *DeletePublishSubscribeRequest) (response *DeletePublishSubscribeResponse, err error) {
    if request == nil {
        request = NewDeletePublishSubscribeRequest()
    }
    response = NewDeletePublishSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountsRequest() (request *DescribeAccountsRequest) {
    request = &DescribeAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeAccounts")
    return
}

func NewDescribeAccountsResponse() (response *DescribeAccountsResponse) {
    response = &DescribeAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeAccounts）用于拉取实例账户列表。
func (c *Client) DescribeAccounts(request *DescribeAccountsRequest) (response *DescribeAccountsResponse, err error) {
    if request == nil {
        request = NewDescribeAccountsRequest()
    }
    response = NewDescribeAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupByFlowIdRequest() (request *DescribeBackupByFlowIdRequest) {
    request = &DescribeBackupByFlowIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeBackupByFlowId")
    return
}

func NewDescribeBackupByFlowIdResponse() (response *DescribeBackupByFlowIdResponse) {
    response = &DescribeBackupByFlowIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeBackupByFlowId)用于通过备份创建流程的ID查询创建的备份详情，流程ID可从接口CreateBackup中获得。
func (c *Client) DescribeBackupByFlowId(request *DescribeBackupByFlowIdRequest) (response *DescribeBackupByFlowIdResponse, err error) {
    if request == nil {
        request = NewDescribeBackupByFlowIdRequest()
    }
    response = NewDescribeBackupByFlowIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupsRequest() (request *DescribeBackupsRequest) {
    request = &DescribeBackupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeBackups")
    return
}

func NewDescribeBackupsResponse() (response *DescribeBackupsResponse) {
    response = &DescribeBackupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeBackups)用于查询备份列表。
func (c *Client) DescribeBackups(request *DescribeBackupsRequest) (response *DescribeBackupsResponse, err error) {
    if request == nil {
        request = NewDescribeBackupsRequest()
    }
    response = NewDescribeBackupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCrossRegionZoneRequest() (request *DescribeCrossRegionZoneRequest) {
    request = &DescribeCrossRegionZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeCrossRegionZone")
    return
}

func NewDescribeCrossRegionZoneResponse() (response *DescribeCrossRegionZoneResponse) {
    response = &DescribeCrossRegionZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeCrossRegionZone)根据主实例查询备机的容灾地域和可用区。
func (c *Client) DescribeCrossRegionZone(request *DescribeCrossRegionZoneRequest) (response *DescribeCrossRegionZoneResponse, err error) {
    if request == nil {
        request = NewDescribeCrossRegionZoneRequest()
    }
    response = NewDescribeCrossRegionZoneResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstancesRequest() (request *DescribeDBInstancesRequest) {
    request = &DescribeDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeDBInstances")
    return
}

func NewDescribeDBInstancesResponse() (response *DescribeDBInstancesResponse) {
    response = &DescribeDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeDBInstances)用于查询实例列表。
func (c *Client) DescribeDBInstances(request *DescribeDBInstancesRequest) (response *DescribeDBInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstancesRequest()
    }
    response = NewDescribeDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBSecurityGroupsRequest() (request *DescribeDBSecurityGroupsRequest) {
    request = &DescribeDBSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeDBSecurityGroups")
    return
}

func NewDescribeDBSecurityGroupsResponse() (response *DescribeDBSecurityGroupsResponse) {
    response = &DescribeDBSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeDBSecurityGroups)用于查询实例的安全组详情。
func (c *Client) DescribeDBSecurityGroups(request *DescribeDBSecurityGroupsRequest) (response *DescribeDBSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeDBSecurityGroupsRequest()
    }
    response = NewDescribeDBSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBsRequest() (request *DescribeDBsRequest) {
    request = &DescribeDBsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeDBs")
    return
}

func NewDescribeDBsResponse() (response *DescribeDBsResponse) {
    response = &DescribeDBsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeDBs）用于查询数据库列表。
func (c *Client) DescribeDBs(request *DescribeDBsRequest) (response *DescribeDBsResponse, err error) {
    if request == nil {
        request = NewDescribeDBsRequest()
    }
    response = NewDescribeDBsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFlowStatusRequest() (request *DescribeFlowStatusRequest) {
    request = &DescribeFlowStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeFlowStatus")
    return
}

func NewDescribeFlowStatusResponse() (response *DescribeFlowStatusResponse) {
    response = &DescribeFlowStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeFlowStatus)用于查询流程状态。
func (c *Client) DescribeFlowStatus(request *DescribeFlowStatusRequest) (response *DescribeFlowStatusResponse, err error) {
    if request == nil {
        request = NewDescribeFlowStatusRequest()
    }
    response = NewDescribeFlowStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMaintenanceSpanRequest() (request *DescribeMaintenanceSpanRequest) {
    request = &DescribeMaintenanceSpanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeMaintenanceSpan")
    return
}

func NewDescribeMaintenanceSpanResponse() (response *DescribeMaintenanceSpanResponse) {
    response = &DescribeMaintenanceSpanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeMaintenanceSpan）根据实例ID查询该实例的可维护时间窗。
func (c *Client) DescribeMaintenanceSpan(request *DescribeMaintenanceSpanRequest) (response *DescribeMaintenanceSpanResponse, err error) {
    if request == nil {
        request = NewDescribeMaintenanceSpanRequest()
    }
    response = NewDescribeMaintenanceSpanResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMigrationDatabasesRequest() (request *DescribeMigrationDatabasesRequest) {
    request = &DescribeMigrationDatabasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeMigrationDatabases")
    return
}

func NewDescribeMigrationDatabasesResponse() (response *DescribeMigrationDatabasesResponse) {
    response = &DescribeMigrationDatabasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeMigrationDatabases）的作用是查询待迁移数据库列表
func (c *Client) DescribeMigrationDatabases(request *DescribeMigrationDatabasesRequest) (response *DescribeMigrationDatabasesResponse, err error) {
    if request == nil {
        request = NewDescribeMigrationDatabasesRequest()
    }
    response = NewDescribeMigrationDatabasesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMigrationDetailRequest() (request *DescribeMigrationDetailRequest) {
    request = &DescribeMigrationDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeMigrationDetail")
    return
}

func NewDescribeMigrationDetailResponse() (response *DescribeMigrationDetailResponse) {
    response = &DescribeMigrationDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeMigrationDetail）用于查询迁移任务的详细情况
func (c *Client) DescribeMigrationDetail(request *DescribeMigrationDetailRequest) (response *DescribeMigrationDetailResponse, err error) {
    if request == nil {
        request = NewDescribeMigrationDetailRequest()
    }
    response = NewDescribeMigrationDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMigrationsRequest() (request *DescribeMigrationsRequest) {
    request = &DescribeMigrationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeMigrations")
    return
}

func NewDescribeMigrationsResponse() (response *DescribeMigrationsResponse) {
    response = &DescribeMigrationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeMigrations）根据输入的限定条件，查询符合条件的迁移任务列表
func (c *Client) DescribeMigrations(request *DescribeMigrationsRequest) (response *DescribeMigrationsResponse, err error) {
    if request == nil {
        request = NewDescribeMigrationsRequest()
    }
    response = NewDescribeMigrationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrdersRequest() (request *DescribeOrdersRequest) {
    request = &DescribeOrdersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeOrders")
    return
}

func NewDescribeOrdersResponse() (response *DescribeOrdersResponse) {
    response = &DescribeOrdersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeOrders）用于查询订单信息
func (c *Client) DescribeOrders(request *DescribeOrdersRequest) (response *DescribeOrdersResponse, err error) {
    if request == nil {
        request = NewDescribeOrdersRequest()
    }
    response = NewDescribeOrdersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProductConfigRequest() (request *DescribeProductConfigRequest) {
    request = &DescribeProductConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeProductConfig")
    return
}

func NewDescribeProductConfigResponse() (response *DescribeProductConfigResponse) {
    response = &DescribeProductConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeProductConfig) 用于查询售卖规格配置。
func (c *Client) DescribeProductConfig(request *DescribeProductConfigRequest) (response *DescribeProductConfigResponse, err error) {
    if request == nil {
        request = NewDescribeProductConfigRequest()
    }
    response = NewDescribeProductConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectSecurityGroupsRequest() (request *DescribeProjectSecurityGroupsRequest) {
    request = &DescribeProjectSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeProjectSecurityGroups")
    return
}

func NewDescribeProjectSecurityGroupsResponse() (response *DescribeProjectSecurityGroupsResponse) {
    response = &DescribeProjectSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeProjectSecurityGroups)用于查询项目的安全组详情。
func (c *Client) DescribeProjectSecurityGroups(request *DescribeProjectSecurityGroupsRequest) (response *DescribeProjectSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeProjectSecurityGroupsRequest()
    }
    response = NewDescribeProjectSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePublishSubscribeRequest() (request *DescribePublishSubscribeRequest) {
    request = &DescribePublishSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribePublishSubscribe")
    return
}

func NewDescribePublishSubscribeResponse() (response *DescribePublishSubscribeResponse) {
    response = &DescribePublishSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribePublishSubscribe）用于查询发布订阅关系列表。
func (c *Client) DescribePublishSubscribe(request *DescribePublishSubscribeRequest) (response *DescribePublishSubscribeResponse, err error) {
    if request == nil {
        request = NewDescribePublishSubscribeRequest()
    }
    response = NewDescribePublishSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReadOnlyGroupByReadOnlyInstanceRequest() (request *DescribeReadOnlyGroupByReadOnlyInstanceRequest) {
    request = &DescribeReadOnlyGroupByReadOnlyInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeReadOnlyGroupByReadOnlyInstance")
    return
}

func NewDescribeReadOnlyGroupByReadOnlyInstanceResponse() (response *DescribeReadOnlyGroupByReadOnlyInstanceResponse) {
    response = &DescribeReadOnlyGroupByReadOnlyInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeReadOnlyGroupByReadOnlyInstance）用于通过只读副本实例ID查询其所在的只读组。
func (c *Client) DescribeReadOnlyGroupByReadOnlyInstance(request *DescribeReadOnlyGroupByReadOnlyInstanceRequest) (response *DescribeReadOnlyGroupByReadOnlyInstanceResponse, err error) {
    if request == nil {
        request = NewDescribeReadOnlyGroupByReadOnlyInstanceRequest()
    }
    response = NewDescribeReadOnlyGroupByReadOnlyInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReadOnlyGroupDetailsRequest() (request *DescribeReadOnlyGroupDetailsRequest) {
    request = &DescribeReadOnlyGroupDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeReadOnlyGroupDetails")
    return
}

func NewDescribeReadOnlyGroupDetailsResponse() (response *DescribeReadOnlyGroupDetailsResponse) {
    response = &DescribeReadOnlyGroupDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeReadOnlyGroupDetails）用于查询只读组详情。
func (c *Client) DescribeReadOnlyGroupDetails(request *DescribeReadOnlyGroupDetailsRequest) (response *DescribeReadOnlyGroupDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeReadOnlyGroupDetailsRequest()
    }
    response = NewDescribeReadOnlyGroupDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReadOnlyGroupListRequest() (request *DescribeReadOnlyGroupListRequest) {
    request = &DescribeReadOnlyGroupListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeReadOnlyGroupList")
    return
}

func NewDescribeReadOnlyGroupListResponse() (response *DescribeReadOnlyGroupListResponse) {
    response = &DescribeReadOnlyGroupListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeReadOnlyGroupList）用于查询只读组列表。
func (c *Client) DescribeReadOnlyGroupList(request *DescribeReadOnlyGroupListRequest) (response *DescribeReadOnlyGroupListResponse, err error) {
    if request == nil {
        request = NewDescribeReadOnlyGroupListRequest()
    }
    response = NewDescribeReadOnlyGroupListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegionsRequest() (request *DescribeRegionsRequest) {
    request = &DescribeRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeRegions")
    return
}

func NewDescribeRegionsResponse() (response *DescribeRegionsResponse) {
    response = &DescribeRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeRegions) 用于查询售卖地域信息。
func (c *Client) DescribeRegions(request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeRegionsRequest()
    }
    response = NewDescribeRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRollbackTimeRequest() (request *DescribeRollbackTimeRequest) {
    request = &DescribeRollbackTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeRollbackTime")
    return
}

func NewDescribeRollbackTimeResponse() (response *DescribeRollbackTimeResponse) {
    response = &DescribeRollbackTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeRollbackTime）用于查询实例可回档时间范围
func (c *Client) DescribeRollbackTime(request *DescribeRollbackTimeRequest) (response *DescribeRollbackTimeResponse, err error) {
    if request == nil {
        request = NewDescribeRollbackTimeRequest()
    }
    response = NewDescribeRollbackTimeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSlowlogsRequest() (request *DescribeSlowlogsRequest) {
    request = &DescribeSlowlogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeSlowlogs")
    return
}

func NewDescribeSlowlogsResponse() (response *DescribeSlowlogsResponse) {
    response = &DescribeSlowlogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeSlowlogs）用于获取慢查询日志文件信息
func (c *Client) DescribeSlowlogs(request *DescribeSlowlogsRequest) (response *DescribeSlowlogsResponse, err error) {
    if request == nil {
        request = NewDescribeSlowlogsRequest()
    }
    response = NewDescribeSlowlogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZonesRequest() (request *DescribeZonesRequest) {
    request = &DescribeZonesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeZones")
    return
}

func NewDescribeZonesResponse() (response *DescribeZonesResponse) {
    response = &DescribeZonesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeZones) 用于查询当前可售卖的可用区信息。
func (c *Client) DescribeZones(request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
    if request == nil {
        request = NewDescribeZonesRequest()
    }
    response = NewDescribeZonesResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateSecurityGroupsRequest() (request *DisassociateSecurityGroupsRequest) {
    request = &DisassociateSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "DisassociateSecurityGroups")
    return
}

func NewDisassociateSecurityGroupsResponse() (response *DisassociateSecurityGroupsResponse) {
    response = &DisassociateSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DisassociateSecurityGroups)用于安全组批量解绑实例。
func (c *Client) DisassociateSecurityGroups(request *DisassociateSecurityGroupsRequest) (response *DisassociateSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDisassociateSecurityGroupsRequest()
    }
    response = NewDisassociateSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceCreateDBInstancesRequest() (request *InquiryPriceCreateDBInstancesRequest) {
    request = &InquiryPriceCreateDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "InquiryPriceCreateDBInstances")
    return
}

func NewInquiryPriceCreateDBInstancesResponse() (response *InquiryPriceCreateDBInstancesResponse) {
    response = &InquiryPriceCreateDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（InquiryPriceCreateDBInstances）用于查询申请实例价格。
func (c *Client) InquiryPriceCreateDBInstances(request *InquiryPriceCreateDBInstancesRequest) (response *InquiryPriceCreateDBInstancesResponse, err error) {
    if request == nil {
        request = NewInquiryPriceCreateDBInstancesRequest()
    }
    response = NewInquiryPriceCreateDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceRenewDBInstanceRequest() (request *InquiryPriceRenewDBInstanceRequest) {
    request = &InquiryPriceRenewDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "InquiryPriceRenewDBInstance")
    return
}

func NewInquiryPriceRenewDBInstanceResponse() (response *InquiryPriceRenewDBInstanceResponse) {
    response = &InquiryPriceRenewDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（InquiryPriceRenewDBInstance）用于查询续费实例的价格。
func (c *Client) InquiryPriceRenewDBInstance(request *InquiryPriceRenewDBInstanceRequest) (response *InquiryPriceRenewDBInstanceResponse, err error) {
    if request == nil {
        request = NewInquiryPriceRenewDBInstanceRequest()
    }
    response = NewInquiryPriceRenewDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceUpgradeDBInstanceRequest() (request *InquiryPriceUpgradeDBInstanceRequest) {
    request = &InquiryPriceUpgradeDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "InquiryPriceUpgradeDBInstance")
    return
}

func NewInquiryPriceUpgradeDBInstanceResponse() (response *InquiryPriceUpgradeDBInstanceResponse) {
    response = &InquiryPriceUpgradeDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（InquiryPriceUpgradeDBInstance）用于查询升级实例的价格。
func (c *Client) InquiryPriceUpgradeDBInstance(request *InquiryPriceUpgradeDBInstanceRequest) (response *InquiryPriceUpgradeDBInstanceResponse, err error) {
    if request == nil {
        request = NewInquiryPriceUpgradeDBInstanceRequest()
    }
    response = NewInquiryPriceUpgradeDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccountPrivilegeRequest() (request *ModifyAccountPrivilegeRequest) {
    request = &ModifyAccountPrivilegeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyAccountPrivilege")
    return
}

func NewModifyAccountPrivilegeResponse() (response *ModifyAccountPrivilegeResponse) {
    response = &ModifyAccountPrivilegeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyAccountPrivilege）用于修改实例账户权限。
func (c *Client) ModifyAccountPrivilege(request *ModifyAccountPrivilegeRequest) (response *ModifyAccountPrivilegeResponse, err error) {
    if request == nil {
        request = NewModifyAccountPrivilegeRequest()
    }
    response = NewModifyAccountPrivilegeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccountRemarkRequest() (request *ModifyAccountRemarkRequest) {
    request = &ModifyAccountRemarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyAccountRemark")
    return
}

func NewModifyAccountRemarkResponse() (response *ModifyAccountRemarkResponse) {
    response = &ModifyAccountRemarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyAccountRemark）用于修改账户备注。
func (c *Client) ModifyAccountRemark(request *ModifyAccountRemarkRequest) (response *ModifyAccountRemarkResponse, err error) {
    if request == nil {
        request = NewModifyAccountRemarkRequest()
    }
    response = NewModifyAccountRemarkResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBackupNameRequest() (request *ModifyBackupNameRequest) {
    request = &ModifyBackupNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyBackupName")
    return
}

func NewModifyBackupNameResponse() (response *ModifyBackupNameResponse) {
    response = &ModifyBackupNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(ModifyBackupName)用于修改备份名称。
func (c *Client) ModifyBackupName(request *ModifyBackupNameRequest) (response *ModifyBackupNameResponse, err error) {
    if request == nil {
        request = NewModifyBackupNameRequest()
    }
    response = NewModifyBackupNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBackupStrategyRequest() (request *ModifyBackupStrategyRequest) {
    request = &ModifyBackupStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyBackupStrategy")
    return
}

func NewModifyBackupStrategyResponse() (response *ModifyBackupStrategyResponse) {
    response = &ModifyBackupStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyBackupStrategy）用于修改备份策略
func (c *Client) ModifyBackupStrategy(request *ModifyBackupStrategyRequest) (response *ModifyBackupStrategyResponse, err error) {
    if request == nil {
        request = NewModifyBackupStrategyRequest()
    }
    response = NewModifyBackupStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceNameRequest() (request *ModifyDBInstanceNameRequest) {
    request = &ModifyDBInstanceNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyDBInstanceName")
    return
}

func NewModifyDBInstanceNameResponse() (response *ModifyDBInstanceNameResponse) {
    response = &ModifyDBInstanceNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyDBInstanceName）用于修改实例名字。
func (c *Client) ModifyDBInstanceName(request *ModifyDBInstanceNameRequest) (response *ModifyDBInstanceNameResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceNameRequest()
    }
    response = NewModifyDBInstanceNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceProjectRequest() (request *ModifyDBInstanceProjectRequest) {
    request = &ModifyDBInstanceProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyDBInstanceProject")
    return
}

func NewModifyDBInstanceProjectResponse() (response *ModifyDBInstanceProjectResponse) {
    response = &ModifyDBInstanceProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyDBInstanceProject）用于修改数据库实例所属项目。
func (c *Client) ModifyDBInstanceProject(request *ModifyDBInstanceProjectRequest) (response *ModifyDBInstanceProjectResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceProjectRequest()
    }
    response = NewModifyDBInstanceProjectResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceRenewFlagRequest() (request *ModifyDBInstanceRenewFlagRequest) {
    request = &ModifyDBInstanceRenewFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyDBInstanceRenewFlag")
    return
}

func NewModifyDBInstanceRenewFlagResponse() (response *ModifyDBInstanceRenewFlagResponse) {
    response = &ModifyDBInstanceRenewFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyDBInstanceRenewFlag）用于修改实例续费标记
func (c *Client) ModifyDBInstanceRenewFlag(request *ModifyDBInstanceRenewFlagRequest) (response *ModifyDBInstanceRenewFlagResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceRenewFlagRequest()
    }
    response = NewModifyDBInstanceRenewFlagResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceSecurityGroupsRequest() (request *ModifyDBInstanceSecurityGroupsRequest) {
    request = &ModifyDBInstanceSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyDBInstanceSecurityGroups")
    return
}

func NewModifyDBInstanceSecurityGroupsResponse() (response *ModifyDBInstanceSecurityGroupsResponse) {
    response = &ModifyDBInstanceSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(ModifyDBInstanceSecurityGroups)用于修改实例绑定的安全组。
func (c *Client) ModifyDBInstanceSecurityGroups(request *ModifyDBInstanceSecurityGroupsRequest) (response *ModifyDBInstanceSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceSecurityGroupsRequest()
    }
    response = NewModifyDBInstanceSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBNameRequest() (request *ModifyDBNameRequest) {
    request = &ModifyDBNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyDBName")
    return
}

func NewModifyDBNameResponse() (response *ModifyDBNameResponse) {
    response = &ModifyDBNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyDBName）用于更新数据库名。
func (c *Client) ModifyDBName(request *ModifyDBNameRequest) (response *ModifyDBNameResponse, err error) {
    if request == nil {
        request = NewModifyDBNameRequest()
    }
    response = NewModifyDBNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBRemarkRequest() (request *ModifyDBRemarkRequest) {
    request = &ModifyDBRemarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyDBRemark")
    return
}

func NewModifyDBRemarkResponse() (response *ModifyDBRemarkResponse) {
    response = &ModifyDBRemarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyDBRemark）用于修改数据库备注。
func (c *Client) ModifyDBRemark(request *ModifyDBRemarkRequest) (response *ModifyDBRemarkResponse, err error) {
    if request == nil {
        request = NewModifyDBRemarkRequest()
    }
    response = NewModifyDBRemarkResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMaintenanceSpanRequest() (request *ModifyMaintenanceSpanRequest) {
    request = &ModifyMaintenanceSpanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyMaintenanceSpan")
    return
}

func NewModifyMaintenanceSpanResponse() (response *ModifyMaintenanceSpanResponse) {
    response = &ModifyMaintenanceSpanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyMaintenanceSpan）用于修改实例的可维护时间窗
func (c *Client) ModifyMaintenanceSpan(request *ModifyMaintenanceSpanRequest) (response *ModifyMaintenanceSpanResponse, err error) {
    if request == nil {
        request = NewModifyMaintenanceSpanRequest()
    }
    response = NewModifyMaintenanceSpanResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMigrationRequest() (request *ModifyMigrationRequest) {
    request = &ModifyMigrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyMigration")
    return
}

func NewModifyMigrationResponse() (response *ModifyMigrationResponse) {
    response = &ModifyMigrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyMigration）可以修改已有的迁移任务信息
func (c *Client) ModifyMigration(request *ModifyMigrationRequest) (response *ModifyMigrationResponse, err error) {
    if request == nil {
        request = NewModifyMigrationRequest()
    }
    response = NewModifyMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPublishSubscribeNameRequest() (request *ModifyPublishSubscribeNameRequest) {
    request = &ModifyPublishSubscribeNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyPublishSubscribeName")
    return
}

func NewModifyPublishSubscribeNameResponse() (response *ModifyPublishSubscribeNameResponse) {
    response = &ModifyPublishSubscribeNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyPublishSubscribeName）修改发布订阅的名称。
func (c *Client) ModifyPublishSubscribeName(request *ModifyPublishSubscribeNameRequest) (response *ModifyPublishSubscribeNameResponse, err error) {
    if request == nil {
        request = NewModifyPublishSubscribeNameRequest()
    }
    response = NewModifyPublishSubscribeNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyReadOnlyGroupDetailsRequest() (request *ModifyReadOnlyGroupDetailsRequest) {
    request = &ModifyReadOnlyGroupDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyReadOnlyGroupDetails")
    return
}

func NewModifyReadOnlyGroupDetailsResponse() (response *ModifyReadOnlyGroupDetailsResponse) {
    response = &ModifyReadOnlyGroupDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyReadOnlyGroupDetails）用于修改只读组详情。
func (c *Client) ModifyReadOnlyGroupDetails(request *ModifyReadOnlyGroupDetailsRequest) (response *ModifyReadOnlyGroupDetailsResponse, err error) {
    if request == nil {
        request = NewModifyReadOnlyGroupDetailsRequest()
    }
    response = NewModifyReadOnlyGroupDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewQueryMigrationCheckProcessRequest() (request *QueryMigrationCheckProcessRequest) {
    request = &QueryMigrationCheckProcessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "QueryMigrationCheckProcess")
    return
}

func NewQueryMigrationCheckProcessResponse() (response *QueryMigrationCheckProcessResponse) {
    response = &QueryMigrationCheckProcessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（QueryMigrationCheckProcess）的作用是查询迁移检查任务的进度，适用于迁移源的类型为TencentDB for SQLServer 的迁移方式
func (c *Client) QueryMigrationCheckProcess(request *QueryMigrationCheckProcessRequest) (response *QueryMigrationCheckProcessResponse, err error) {
    if request == nil {
        request = NewQueryMigrationCheckProcessRequest()
    }
    response = NewQueryMigrationCheckProcessResponse()
    err = c.Send(request, response)
    return
}

func NewRecycleDBInstanceRequest() (request *RecycleDBInstanceRequest) {
    request = &RecycleDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "RecycleDBInstance")
    return
}

func NewRecycleDBInstanceResponse() (response *RecycleDBInstanceResponse) {
    response = &RecycleDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（RecycleDBInstance）用于主动回收已下线的SQLSERVER实例
func (c *Client) RecycleDBInstance(request *RecycleDBInstanceRequest) (response *RecycleDBInstanceResponse, err error) {
    if request == nil {
        request = NewRecycleDBInstanceRequest()
    }
    response = NewRecycleDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewRecycleReadOnlyGroupRequest() (request *RecycleReadOnlyGroupRequest) {
    request = &RecycleReadOnlyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "RecycleReadOnlyGroup")
    return
}

func NewRecycleReadOnlyGroupResponse() (response *RecycleReadOnlyGroupResponse) {
    response = &RecycleReadOnlyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（RecycleReadOnlyGroup）立即回收只读组的资源，只读组占用的vip等资源将立即释放且不可找回。
func (c *Client) RecycleReadOnlyGroup(request *RecycleReadOnlyGroupRequest) (response *RecycleReadOnlyGroupResponse, err error) {
    if request == nil {
        request = NewRecycleReadOnlyGroupRequest()
    }
    response = NewRecycleReadOnlyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveBackupsRequest() (request *RemoveBackupsRequest) {
    request = &RemoveBackupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "RemoveBackups")
    return
}

func NewRemoveBackupsResponse() (response *RemoveBackupsResponse) {
    response = &RemoveBackupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（RemoveBackups）可以删除用户手动创建的备份文件。待删除的备份策略可以是实例备份，也可以是多库备份。
func (c *Client) RemoveBackups(request *RemoveBackupsRequest) (response *RemoveBackupsResponse, err error) {
    if request == nil {
        request = NewRemoveBackupsRequest()
    }
    response = NewRemoveBackupsResponse()
    err = c.Send(request, response)
    return
}

func NewRenewDBInstanceRequest() (request *RenewDBInstanceRequest) {
    request = &RenewDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "RenewDBInstance")
    return
}

func NewRenewDBInstanceResponse() (response *RenewDBInstanceResponse) {
    response = &RenewDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（RenewDBInstance）用于续费实例。
func (c *Client) RenewDBInstance(request *RenewDBInstanceRequest) (response *RenewDBInstanceResponse, err error) {
    if request == nil {
        request = NewRenewDBInstanceRequest()
    }
    response = NewRenewDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewRenewPostpaidDBInstanceRequest() (request *RenewPostpaidDBInstanceRequest) {
    request = &RenewPostpaidDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "RenewPostpaidDBInstance")
    return
}

func NewRenewPostpaidDBInstanceResponse() (response *RenewPostpaidDBInstanceResponse) {
    response = &RenewPostpaidDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（RenewPostpaidDBInstance）用于将通过接口TerminateDBInstance手动隔离的按量计费实例从回收站中恢复。
func (c *Client) RenewPostpaidDBInstance(request *RenewPostpaidDBInstanceRequest) (response *RenewPostpaidDBInstanceResponse, err error) {
    if request == nil {
        request = NewRenewPostpaidDBInstanceRequest()
    }
    response = NewRenewPostpaidDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewResetAccountPasswordRequest() (request *ResetAccountPasswordRequest) {
    request = &ResetAccountPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "ResetAccountPassword")
    return
}

func NewResetAccountPasswordResponse() (response *ResetAccountPasswordResponse) {
    response = &ResetAccountPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ResetAccountPassword）用于重置实例的账户密码。
func (c *Client) ResetAccountPassword(request *ResetAccountPasswordRequest) (response *ResetAccountPasswordResponse, err error) {
    if request == nil {
        request = NewResetAccountPasswordRequest()
    }
    response = NewResetAccountPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewRestartDBInstanceRequest() (request *RestartDBInstanceRequest) {
    request = &RestartDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "RestartDBInstance")
    return
}

func NewRestartDBInstanceResponse() (response *RestartDBInstanceResponse) {
    response = &RestartDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（RestartDBInstance）用于重启数据库实例。
func (c *Client) RestartDBInstance(request *RestartDBInstanceRequest) (response *RestartDBInstanceResponse, err error) {
    if request == nil {
        request = NewRestartDBInstanceRequest()
    }
    response = NewRestartDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewRestoreInstanceRequest() (request *RestoreInstanceRequest) {
    request = &RestoreInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "RestoreInstance")
    return
}

func NewRestoreInstanceResponse() (response *RestoreInstanceResponse) {
    response = &RestoreInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（RestoreInstance）用于根据备份文件恢复实例。
func (c *Client) RestoreInstance(request *RestoreInstanceRequest) (response *RestoreInstanceResponse, err error) {
    if request == nil {
        request = NewRestoreInstanceRequest()
    }
    response = NewRestoreInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewRollbackInstanceRequest() (request *RollbackInstanceRequest) {
    request = &RollbackInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "RollbackInstance")
    return
}

func NewRollbackInstanceResponse() (response *RollbackInstanceResponse) {
    response = &RollbackInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（RollbackInstance）用于回档实例
func (c *Client) RollbackInstance(request *RollbackInstanceRequest) (response *RollbackInstanceResponse, err error) {
    if request == nil {
        request = NewRollbackInstanceRequest()
    }
    response = NewRollbackInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewRunMigrationRequest() (request *RunMigrationRequest) {
    request = &RunMigrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "RunMigration")
    return
}

func NewRunMigrationResponse() (response *RunMigrationResponse) {
    response = &RunMigrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（RunMigration）用于启动迁移任务，开始迁移
func (c *Client) RunMigration(request *RunMigrationRequest) (response *RunMigrationResponse, err error) {
    if request == nil {
        request = NewRunMigrationRequest()
    }
    response = NewRunMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewStartMigrationCheckRequest() (request *StartMigrationCheckRequest) {
    request = &StartMigrationCheckRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "StartMigrationCheck")
    return
}

func NewStartMigrationCheckResponse() (response *StartMigrationCheckResponse) {
    response = &StartMigrationCheckResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（StartMigrationCheck）的作用是启动一个迁移前的校验任务，适用于迁移源的类型为TencentDB for SQLServer 的迁移方式
func (c *Client) StartMigrationCheck(request *StartMigrationCheckRequest) (response *StartMigrationCheckResponse, err error) {
    if request == nil {
        request = NewStartMigrationCheckRequest()
    }
    response = NewStartMigrationCheckResponse()
    err = c.Send(request, response)
    return
}

func NewStopMigrationRequest() (request *StopMigrationRequest) {
    request = &StopMigrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "StopMigration")
    return
}

func NewStopMigrationResponse() (response *StopMigrationResponse) {
    response = &StopMigrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（StopMigration）作用是中止一个迁移任务
func (c *Client) StopMigration(request *StopMigrationRequest) (response *StopMigrationResponse, err error) {
    if request == nil {
        request = NewStopMigrationRequest()
    }
    response = NewStopMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateDBInstanceRequest() (request *TerminateDBInstanceRequest) {
    request = &TerminateDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "TerminateDBInstance")
    return
}

func NewTerminateDBInstanceResponse() (response *TerminateDBInstanceResponse) {
    response = &TerminateDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(TerminateDBInstance)用于主动隔离实例，使得实例进入回收站。
func (c *Client) TerminateDBInstance(request *TerminateDBInstanceRequest) (response *TerminateDBInstanceResponse, err error) {
    if request == nil {
        request = NewTerminateDBInstanceRequest()
    }
    response = NewTerminateDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeDBInstanceRequest() (request *UpgradeDBInstanceRequest) {
    request = &UpgradeDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("sqlserver", APIVersion, "UpgradeDBInstance")
    return
}

func NewUpgradeDBInstanceResponse() (response *UpgradeDBInstanceResponse) {
    response = &UpgradeDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（UpgradeDBInstance）用于升级实例
func (c *Client) UpgradeDBInstance(request *UpgradeDBInstanceRequest) (response *UpgradeDBInstanceResponse, err error) {
    if request == nil {
        request = NewUpgradeDBInstanceRequest()
    }
    response = NewUpgradeDBInstanceResponse()
    err = c.Send(request, response)
    return
}
