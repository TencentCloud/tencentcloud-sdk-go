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

package v20190107

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-01-07"

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


func NewAddInstancesRequest() (request *AddInstancesRequest) {
    request = &AddInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cynosdb", APIVersion, "AddInstances")
    return
}

func NewAddInstancesResponse() (response *AddInstancesResponse) {
    response = &AddInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（AddInstances）用于集群添加实例
func (c *Client) AddInstances(request *AddInstancesRequest) (response *AddInstancesResponse, err error) {
    if request == nil {
        request = NewAddInstancesRequest()
    }
    response = NewAddInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClustersRequest() (request *CreateClustersRequest) {
    request = &CreateClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cynosdb", APIVersion, "CreateClusters")
    return
}

func NewCreateClustersResponse() (response *CreateClustersResponse) {
    response = &CreateClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建集群
func (c *Client) CreateClusters(request *CreateClustersRequest) (response *CreateClustersResponse, err error) {
    if request == nil {
        request = NewCreateClustersRequest()
    }
    response = NewCreateClustersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountsRequest() (request *DescribeAccountsRequest) {
    request = &DescribeAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeAccounts")
    return
}

func NewDescribeAccountsResponse() (response *DescribeAccountsResponse) {
    response = &DescribeAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeAccounts)用于查询数据库管理账号。
func (c *Client) DescribeAccounts(request *DescribeAccountsRequest) (response *DescribeAccountsResponse, err error) {
    if request == nil {
        request = NewDescribeAccountsRequest()
    }
    response = NewDescribeAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupConfigRequest() (request *DescribeBackupConfigRequest) {
    request = &DescribeBackupConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeBackupConfig")
    return
}

func NewDescribeBackupConfigResponse() (response *DescribeBackupConfigResponse) {
    response = &DescribeBackupConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取指定集群的备份配置信息，包括全量备份时间段，备份文件保留时间
func (c *Client) DescribeBackupConfig(request *DescribeBackupConfigRequest) (response *DescribeBackupConfigResponse, err error) {
    if request == nil {
        request = NewDescribeBackupConfigRequest()
    }
    response = NewDescribeBackupConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupListRequest() (request *DescribeBackupListRequest) {
    request = &DescribeBackupListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeBackupList")
    return
}

func NewDescribeBackupListResponse() (response *DescribeBackupListResponse) {
    response = &DescribeBackupListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询备份文件列表
func (c *Client) DescribeBackupList(request *DescribeBackupListRequest) (response *DescribeBackupListResponse, err error) {
    if request == nil {
        request = NewDescribeBackupListRequest()
    }
    response = NewDescribeBackupListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterDetailRequest() (request *DescribeClusterDetailRequest) {
    request = &DescribeClusterDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeClusterDetail")
    return
}

func NewDescribeClusterDetailResponse() (response *DescribeClusterDetailResponse) {
    response = &DescribeClusterDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 显示集群详情
func (c *Client) DescribeClusterDetail(request *DescribeClusterDetailRequest) (response *DescribeClusterDetailResponse, err error) {
    if request == nil {
        request = NewDescribeClusterDetailRequest()
    }
    response = NewDescribeClusterDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterInstanceGrpsRequest() (request *DescribeClusterInstanceGrpsRequest) {
    request = &DescribeClusterInstanceGrpsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeClusterInstanceGrps")
    return
}

func NewDescribeClusterInstanceGrpsResponse() (response *DescribeClusterInstanceGrpsResponse) {
    response = &DescribeClusterInstanceGrpsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeClusterInstanceGrps）用于查询实例组
func (c *Client) DescribeClusterInstanceGrps(request *DescribeClusterInstanceGrpsRequest) (response *DescribeClusterInstanceGrpsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterInstanceGrpsRequest()
    }
    response = NewDescribeClusterInstanceGrpsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClustersRequest() (request *DescribeClustersRequest) {
    request = &DescribeClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeClusters")
    return
}

func NewDescribeClustersResponse() (response *DescribeClustersResponse) {
    response = &DescribeClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询集群列表
func (c *Client) DescribeClusters(request *DescribeClustersRequest) (response *DescribeClustersResponse, err error) {
    if request == nil {
        request = NewDescribeClustersRequest()
    }
    response = NewDescribeClustersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBSecurityGroupsRequest() (request *DescribeDBSecurityGroupsRequest) {
    request = &DescribeDBSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeDBSecurityGroups")
    return
}

func NewDescribeDBSecurityGroupsResponse() (response *DescribeDBSecurityGroupsResponse) {
    response = &DescribeDBSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询实例安全组信息
func (c *Client) DescribeDBSecurityGroups(request *DescribeDBSecurityGroupsRequest) (response *DescribeDBSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeDBSecurityGroupsRequest()
    }
    response = NewDescribeDBSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceDetailRequest() (request *DescribeInstanceDetailRequest) {
    request = &DescribeInstanceDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeInstanceDetail")
    return
}

func NewDescribeInstanceDetailResponse() (response *DescribeInstanceDetailResponse) {
    response = &DescribeInstanceDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeInstanceDetail)用于查询实例详情。
func (c *Client) DescribeInstanceDetail(request *DescribeInstanceDetailRequest) (response *DescribeInstanceDetailResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceDetailRequest()
    }
    response = NewDescribeInstanceDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceSpecsRequest() (request *DescribeInstanceSpecsRequest) {
    request = &DescribeInstanceSpecsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeInstanceSpecs")
    return
}

func NewDescribeInstanceSpecsResponse() (response *DescribeInstanceSpecsResponse) {
    response = &DescribeInstanceSpecsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeInstanceSpecs）用于查询实例规格
func (c *Client) DescribeInstanceSpecs(request *DescribeInstanceSpecsRequest) (response *DescribeInstanceSpecsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceSpecsRequest()
    }
    response = NewDescribeInstanceSpecsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeInstances")
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeInstances)用于查询实例列表。
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMaintainPeriodRequest() (request *DescribeMaintainPeriodRequest) {
    request = &DescribeMaintainPeriodRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeMaintainPeriod")
    return
}

func NewDescribeMaintainPeriodResponse() (response *DescribeMaintainPeriodResponse) {
    response = &DescribeMaintainPeriodResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询实例维护时间窗
func (c *Client) DescribeMaintainPeriod(request *DescribeMaintainPeriodRequest) (response *DescribeMaintainPeriodResponse, err error) {
    if request == nil {
        request = NewDescribeMaintainPeriodRequest()
    }
    response = NewDescribeMaintainPeriodResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectSecurityGroupsRequest() (request *DescribeProjectSecurityGroupsRequest) {
    request = &DescribeProjectSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeProjectSecurityGroups")
    return
}

func NewDescribeProjectSecurityGroupsResponse() (response *DescribeProjectSecurityGroupsResponse) {
    response = &DescribeProjectSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询项目安全组信息
func (c *Client) DescribeProjectSecurityGroups(request *DescribeProjectSecurityGroupsRequest) (response *DescribeProjectSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeProjectSecurityGroupsRequest()
    }
    response = NewDescribeProjectSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRollbackTimeRangeRequest() (request *DescribeRollbackTimeRangeRequest) {
    request = &DescribeRollbackTimeRangeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeRollbackTimeRange")
    return
}

func NewDescribeRollbackTimeRangeResponse() (response *DescribeRollbackTimeRangeResponse) {
    response = &DescribeRollbackTimeRangeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询指定集群有效回滚时间范围
func (c *Client) DescribeRollbackTimeRange(request *DescribeRollbackTimeRangeRequest) (response *DescribeRollbackTimeRangeResponse, err error) {
    if request == nil {
        request = NewDescribeRollbackTimeRangeRequest()
    }
    response = NewDescribeRollbackTimeRangeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRollbackTimeValidityRequest() (request *DescribeRollbackTimeValidityRequest) {
    request = &DescribeRollbackTimeValidityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeRollbackTimeValidity")
    return
}

func NewDescribeRollbackTimeValidityResponse() (response *DescribeRollbackTimeValidityResponse) {
    response = &DescribeRollbackTimeValidityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 指定时间和集群查询是否可回滚
func (c *Client) DescribeRollbackTimeValidity(request *DescribeRollbackTimeValidityRequest) (response *DescribeRollbackTimeValidityResponse, err error) {
    if request == nil {
        request = NewDescribeRollbackTimeValidityRequest()
    }
    response = NewDescribeRollbackTimeValidityResponse()
    err = c.Send(request, response)
    return
}

func NewIsolateClusterRequest() (request *IsolateClusterRequest) {
    request = &IsolateClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cynosdb", APIVersion, "IsolateCluster")
    return
}

func NewIsolateClusterResponse() (response *IsolateClusterResponse) {
    response = &IsolateClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 隔离集群
func (c *Client) IsolateCluster(request *IsolateClusterRequest) (response *IsolateClusterResponse, err error) {
    if request == nil {
        request = NewIsolateClusterRequest()
    }
    response = NewIsolateClusterResponse()
    err = c.Send(request, response)
    return
}

func NewIsolateInstanceRequest() (request *IsolateInstanceRequest) {
    request = &IsolateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cynosdb", APIVersion, "IsolateInstance")
    return
}

func NewIsolateInstanceResponse() (response *IsolateInstanceResponse) {
    response = &IsolateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(IsolateInstance)用于隔离实例。
func (c *Client) IsolateInstance(request *IsolateInstanceRequest) (response *IsolateInstanceResponse, err error) {
    if request == nil {
        request = NewIsolateInstanceRequest()
    }
    response = NewIsolateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBackupConfigRequest() (request *ModifyBackupConfigRequest) {
    request = &ModifyBackupConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyBackupConfig")
    return
}

func NewModifyBackupConfigResponse() (response *ModifyBackupConfigResponse) {
    response = &ModifyBackupConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改指定集群的备份配置
func (c *Client) ModifyBackupConfig(request *ModifyBackupConfigRequest) (response *ModifyBackupConfigResponse, err error) {
    if request == nil {
        request = NewModifyBackupConfigRequest()
    }
    response = NewModifyBackupConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceSecurityGroupsRequest() (request *ModifyDBInstanceSecurityGroupsRequest) {
    request = &ModifyDBInstanceSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyDBInstanceSecurityGroups")
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

func NewModifyMaintainPeriodConfigRequest() (request *ModifyMaintainPeriodConfigRequest) {
    request = &ModifyMaintainPeriodConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyMaintainPeriodConfig")
    return
}

func NewModifyMaintainPeriodConfigResponse() (response *ModifyMaintainPeriodConfigResponse) {
    response = &ModifyMaintainPeriodConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改维护时间配置
func (c *Client) ModifyMaintainPeriodConfig(request *ModifyMaintainPeriodConfigRequest) (response *ModifyMaintainPeriodConfigResponse, err error) {
    if request == nil {
        request = NewModifyMaintainPeriodConfigRequest()
    }
    response = NewModifyMaintainPeriodConfigResponse()
    err = c.Send(request, response)
    return
}

func NewOfflineClusterRequest() (request *OfflineClusterRequest) {
    request = &OfflineClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cynosdb", APIVersion, "OfflineCluster")
    return
}

func NewOfflineClusterResponse() (response *OfflineClusterResponse) {
    response = &OfflineClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 下线集群
func (c *Client) OfflineCluster(request *OfflineClusterRequest) (response *OfflineClusterResponse, err error) {
    if request == nil {
        request = NewOfflineClusterRequest()
    }
    response = NewOfflineClusterResponse()
    err = c.Send(request, response)
    return
}

func NewOfflineInstanceRequest() (request *OfflineInstanceRequest) {
    request = &OfflineInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cynosdb", APIVersion, "OfflineInstance")
    return
}

func NewOfflineInstanceResponse() (response *OfflineInstanceResponse) {
    response = &OfflineInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 下线实例
func (c *Client) OfflineInstance(request *OfflineInstanceRequest) (response *OfflineInstanceResponse, err error) {
    if request == nil {
        request = NewOfflineInstanceRequest()
    }
    response = NewOfflineInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewSetRenewFlagRequest() (request *SetRenewFlagRequest) {
    request = &SetRenewFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cynosdb", APIVersion, "SetRenewFlag")
    return
}

func NewSetRenewFlagResponse() (response *SetRenewFlagResponse) {
    response = &SetRenewFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetRenewFlag设置实例的自动续费功能
func (c *Client) SetRenewFlag(request *SetRenewFlagRequest) (response *SetRenewFlagResponse, err error) {
    if request == nil {
        request = NewSetRenewFlagRequest()
    }
    response = NewSetRenewFlagResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeInstanceRequest() (request *UpgradeInstanceRequest) {
    request = &UpgradeInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cynosdb", APIVersion, "UpgradeInstance")
    return
}

func NewUpgradeInstanceResponse() (response *UpgradeInstanceResponse) {
    response = &UpgradeInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 升级实例
func (c *Client) UpgradeInstance(request *UpgradeInstanceRequest) (response *UpgradeInstanceResponse, err error) {
    if request == nil {
        request = NewUpgradeInstanceRequest()
    }
    response = NewUpgradeInstanceResponse()
    err = c.Send(request, response)
    return
}
