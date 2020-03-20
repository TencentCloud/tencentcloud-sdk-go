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

package v20180330

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-03-30"

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


func NewActivateSubscribeRequest() (request *ActivateSubscribeRequest) {
    request = &ActivateSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "ActivateSubscribe")
    return
}

func NewActivateSubscribeResponse() (response *ActivateSubscribeResponse) {
    response = &ActivateSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口用于配置数据订阅，只有在未配置状态的订阅实例才能调用此接口。
func (c *Client) ActivateSubscribe(request *ActivateSubscribeRequest) (response *ActivateSubscribeResponse, err error) {
    if request == nil {
        request = NewActivateSubscribeRequest()
    }
    response = NewActivateSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewCompleteMigrateJobRequest() (request *CompleteMigrateJobRequest) {
    request = &CompleteMigrateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "CompleteMigrateJob")
    return
}

func NewCompleteMigrateJobResponse() (response *CompleteMigrateJobResponse) {
    response = &CompleteMigrateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CompleteMigrateJob）用于完成数据迁移任务。
// 选择采用增量迁移方式的任务, 需要在迁移进度进入准备完成阶段后, 调用本接口, 停止迁移增量数据。
// 通过DescribeMigrateJobs接口查询到任务的状态为准备完成（status=8）时，此时可以调用本接口完成迁移任务。
func (c *Client) CompleteMigrateJob(request *CompleteMigrateJobRequest) (response *CompleteMigrateJobResponse, err error) {
    if request == nil {
        request = NewCompleteMigrateJobRequest()
    }
    response = NewCompleteMigrateJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMigrateCheckJobRequest() (request *CreateMigrateCheckJobRequest) {
    request = &CreateMigrateCheckJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "CreateMigrateCheckJob")
    return
}

func NewCreateMigrateCheckJobResponse() (response *CreateMigrateCheckJobResponse) {
    response = &CreateMigrateCheckJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建校验迁移任务
// 在开始迁移前, 必须调用本接口创建校验, 且校验成功后才能开始迁移. 校验的结果可以通过DescribeMigrateCheckJob查看.
// 校验成功后,迁移任务若有修改, 则必须重新创建校验并通过后, 才能开始迁移.
func (c *Client) CreateMigrateCheckJob(request *CreateMigrateCheckJobRequest) (response *CreateMigrateCheckJobResponse, err error) {
    if request == nil {
        request = NewCreateMigrateCheckJobRequest()
    }
    response = NewCreateMigrateCheckJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMigrateJobRequest() (request *CreateMigrateJobRequest) {
    request = &CreateMigrateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "CreateMigrateJob")
    return
}

func NewCreateMigrateJobResponse() (response *CreateMigrateJobResponse) {
    response = &CreateMigrateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateMigrateJob）用于创建数据迁移任务。
// 
// 如果是金融区链路, 请使用域名: dts.ap-shenzhen-fsi.tencentcloudapi.com
func (c *Client) CreateMigrateJob(request *CreateMigrateJobRequest) (response *CreateMigrateJobResponse, err error) {
    if request == nil {
        request = NewCreateMigrateJobRequest()
    }
    response = NewCreateMigrateJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSubscribeRequest() (request *CreateSubscribeRequest) {
    request = &CreateSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "CreateSubscribe")
    return
}

func NewCreateSubscribeResponse() (response *CreateSubscribeResponse) {
    response = &CreateSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(CreateSubscribe)用于创建一个数据订阅实例。
func (c *Client) CreateSubscribe(request *CreateSubscribeRequest) (response *CreateSubscribeResponse, err error) {
    if request == nil {
        request = NewCreateSubscribeRequest()
    }
    response = NewCreateSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSyncCheckJobRequest() (request *CreateSyncCheckJobRequest) {
    request = &CreateSyncCheckJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "CreateSyncCheckJob")
    return
}

func NewCreateSyncCheckJobResponse() (response *CreateSyncCheckJobResponse) {
    response = &CreateSyncCheckJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 在调用 StartSyncJob 接口启动灾备同步前, 必须调用本接口创建校验, 且校验成功后才能开始同步数据. 校验的结果可以通过 DescribeSyncCheckJob 查看.
// 校验成功后才能启动同步.
func (c *Client) CreateSyncCheckJob(request *CreateSyncCheckJobRequest) (response *CreateSyncCheckJobResponse, err error) {
    if request == nil {
        request = NewCreateSyncCheckJobRequest()
    }
    response = NewCreateSyncCheckJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSyncJobRequest() (request *CreateSyncJobRequest) {
    request = &CreateSyncJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "CreateSyncJob")
    return
}

func NewCreateSyncJobResponse() (response *CreateSyncJobResponse) {
    response = &CreateSyncJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(CreateSyncJob)用于创建灾备同步任务。
// 创建同步任务后，可以通过 CreateSyncCheckJob 接口发起校验任务。校验成功后才可以通过 StartSyncJob 接口启动同步任务。
func (c *Client) CreateSyncJob(request *CreateSyncJobRequest) (response *CreateSyncJobResponse, err error) {
    if request == nil {
        request = NewCreateSyncJobRequest()
    }
    response = NewCreateSyncJobResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMigrateJobRequest() (request *DeleteMigrateJobRequest) {
    request = &DeleteMigrateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "DeleteMigrateJob")
    return
}

func NewDeleteMigrateJobResponse() (response *DeleteMigrateJobResponse) {
    response = &DeleteMigrateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeleteMigrationJob）用于删除数据迁移任务。当通过DescribeMigrateJobs接口查询到任务的状态为：检验中（status=3）、运行中（status=7）、准备完成（status=8）、撤销中（status=11）或者完成中（status=12）时，不允许删除任务。
func (c *Client) DeleteMigrateJob(request *DeleteMigrateJobRequest) (response *DeleteMigrateJobResponse, err error) {
    if request == nil {
        request = NewDeleteMigrateJobRequest()
    }
    response = NewDeleteMigrateJobResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSyncJobRequest() (request *DeleteSyncJobRequest) {
    request = &DeleteSyncJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "DeleteSyncJob")
    return
}

func NewDeleteSyncJobResponse() (response *DeleteSyncJobResponse) {
    response = &DeleteSyncJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除灾备同步任务 （运行中的同步任务不能删除）。
func (c *Client) DeleteSyncJob(request *DeleteSyncJobRequest) (response *DeleteSyncJobResponse, err error) {
    if request == nil {
        request = NewDeleteSyncJobRequest()
    }
    response = NewDeleteSyncJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAsyncRequestInfoRequest() (request *DescribeAsyncRequestInfoRequest) {
    request = &DescribeAsyncRequestInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "DescribeAsyncRequestInfo")
    return
}

func NewDescribeAsyncRequestInfoResponse() (response *DescribeAsyncRequestInfoResponse) {
    response = &DescribeAsyncRequestInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeAsyncRequestInfo）用于查询任务执行结果
func (c *Client) DescribeAsyncRequestInfo(request *DescribeAsyncRequestInfoRequest) (response *DescribeAsyncRequestInfoResponse, err error) {
    if request == nil {
        request = NewDescribeAsyncRequestInfoRequest()
    }
    response = NewDescribeAsyncRequestInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMigrateCheckJobRequest() (request *DescribeMigrateCheckJobRequest) {
    request = &DescribeMigrateCheckJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "DescribeMigrateCheckJob")
    return
}

func NewDescribeMigrateCheckJobResponse() (response *DescribeMigrateCheckJobResponse) {
    response = &DescribeMigrateCheckJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口用于创建校验后,获取校验的结果. 能查询到当前校验的状态和进度. 
// 若通过校验, 则可调用'StartMigrateJob' 开始迁移.
// 若未通过校验, 则能查询到校验失败的原因. 请按照报错, 通过'ModifyMigrateJob'修改迁移配置或是调整源/目标实例的相关参数.
func (c *Client) DescribeMigrateCheckJob(request *DescribeMigrateCheckJobRequest) (response *DescribeMigrateCheckJobResponse, err error) {
    if request == nil {
        request = NewDescribeMigrateCheckJobRequest()
    }
    response = NewDescribeMigrateCheckJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMigrateJobsRequest() (request *DescribeMigrateJobsRequest) {
    request = &DescribeMigrateJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "DescribeMigrateJobs")
    return
}

func NewDescribeMigrateJobsResponse() (response *DescribeMigrateJobsResponse) {
    response = &DescribeMigrateJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询数据迁移任务.
// 如果是金融区链路, 请使用域名: https://dts.ap-shenzhen-fsi.tencentcloudapi.com
func (c *Client) DescribeMigrateJobs(request *DescribeMigrateJobsRequest) (response *DescribeMigrateJobsResponse, err error) {
    if request == nil {
        request = NewDescribeMigrateJobsRequest()
    }
    response = NewDescribeMigrateJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegionConfRequest() (request *DescribeRegionConfRequest) {
    request = &DescribeRegionConfRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "DescribeRegionConf")
    return
}

func NewDescribeRegionConfResponse() (response *DescribeRegionConfResponse) {
    response = &DescribeRegionConfResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeRegionConf）用于查询可售卖订阅实例的地域
func (c *Client) DescribeRegionConf(request *DescribeRegionConfRequest) (response *DescribeRegionConfResponse, err error) {
    if request == nil {
        request = NewDescribeRegionConfRequest()
    }
    response = NewDescribeRegionConfResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubscribeConfRequest() (request *DescribeSubscribeConfRequest) {
    request = &DescribeSubscribeConfRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "DescribeSubscribeConf")
    return
}

func NewDescribeSubscribeConfResponse() (response *DescribeSubscribeConfResponse) {
    response = &DescribeSubscribeConfResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeSubscribeConf）用于查询订阅实例配置
func (c *Client) DescribeSubscribeConf(request *DescribeSubscribeConfRequest) (response *DescribeSubscribeConfResponse, err error) {
    if request == nil {
        request = NewDescribeSubscribeConfRequest()
    }
    response = NewDescribeSubscribeConfResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubscribesRequest() (request *DescribeSubscribesRequest) {
    request = &DescribeSubscribesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "DescribeSubscribes")
    return
}

func NewDescribeSubscribesResponse() (response *DescribeSubscribesResponse) {
    response = &DescribeSubscribesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(DescribeSubscribes)获取数据订阅实例信息列表，默认分页，每次返回20条
func (c *Client) DescribeSubscribes(request *DescribeSubscribesRequest) (response *DescribeSubscribesResponse, err error) {
    if request == nil {
        request = NewDescribeSubscribesRequest()
    }
    response = NewDescribeSubscribesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSyncCheckJobRequest() (request *DescribeSyncCheckJobRequest) {
    request = &DescribeSyncCheckJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "DescribeSyncCheckJob")
    return
}

func NewDescribeSyncCheckJobResponse() (response *DescribeSyncCheckJobResponse) {
    response = &DescribeSyncCheckJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口用于在通过 CreateSyncCheckJob 接口创建灾备同步校验任务后，获取校验的结果。能查询到当前校验的状态和进度。
// 若通过校验, 则可调用 StartSyncJob 启动同步任务。
// 若未通过校验, 则会返回校验失败的原因。 可通过 ModifySyncJob 修改配置，然后再次发起校验。
// 校验任务需要大概约30秒，当返回的 Status 不为 finished 时表示尚未校验完成，需要轮询该接口。
// 如果 Status=finished 且 CheckFlag=1 时表示校验成功。
// 如果 Status=finished 且 CheckFlag !=1 时表示校验失败。
func (c *Client) DescribeSyncCheckJob(request *DescribeSyncCheckJobRequest) (response *DescribeSyncCheckJobResponse, err error) {
    if request == nil {
        request = NewDescribeSyncCheckJobRequest()
    }
    response = NewDescribeSyncCheckJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSyncJobsRequest() (request *DescribeSyncJobsRequest) {
    request = &DescribeSyncJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "DescribeSyncJobs")
    return
}

func NewDescribeSyncJobsResponse() (response *DescribeSyncJobsResponse) {
    response = &DescribeSyncJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询在迁移平台发起的灾备同步任务
func (c *Client) DescribeSyncJobs(request *DescribeSyncJobsRequest) (response *DescribeSyncJobsResponse, err error) {
    if request == nil {
        request = NewDescribeSyncJobsRequest()
    }
    response = NewDescribeSyncJobsResponse()
    err = c.Send(request, response)
    return
}

func NewIsolateSubscribeRequest() (request *IsolateSubscribeRequest) {
    request = &IsolateSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "IsolateSubscribe")
    return
}

func NewIsolateSubscribeResponse() (response *IsolateSubscribeResponse) {
    response = &IsolateSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（IsolateSubscribe）用于隔离小时计费的订阅实例。调用后，订阅实例将不能使用，同时停止计费。
func (c *Client) IsolateSubscribe(request *IsolateSubscribeRequest) (response *IsolateSubscribeResponse, err error) {
    if request == nil {
        request = NewIsolateSubscribeRequest()
    }
    response = NewIsolateSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMigrateJobRequest() (request *ModifyMigrateJobRequest) {
    request = &ModifyMigrateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "ModifyMigrateJob")
    return
}

func NewModifyMigrateJobResponse() (response *ModifyMigrateJobResponse) {
    response = &ModifyMigrateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyMigrateJob）用于修改数据迁移任务。
// 当迁移任务处于下述状态时，允许调用本接口修改迁移任务：迁移创建中（status=1）、 校验成功(status=4)、校验失败(status=5)、迁移失败(status=10)。但源实例、目标实例类型和目标实例地域不允许修改。
// 
// 如果是金融区链路, 请使用域名: dts.ap-shenzhen-fsi.tencentcloudapi.com
func (c *Client) ModifyMigrateJob(request *ModifyMigrateJobRequest) (response *ModifyMigrateJobResponse, err error) {
    if request == nil {
        request = NewModifyMigrateJobRequest()
    }
    response = NewModifyMigrateJobResponse()
    err = c.Send(request, response)
    return
}

func NewModifySubscribeConsumeTimeRequest() (request *ModifySubscribeConsumeTimeRequest) {
    request = &ModifySubscribeConsumeTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "ModifySubscribeConsumeTime")
    return
}

func NewModifySubscribeConsumeTimeResponse() (response *ModifySubscribeConsumeTimeResponse) {
    response = &ModifySubscribeConsumeTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(ModifySubscribeConsumeTime)用于修改数据订阅通道的消费时间点
func (c *Client) ModifySubscribeConsumeTime(request *ModifySubscribeConsumeTimeRequest) (response *ModifySubscribeConsumeTimeResponse, err error) {
    if request == nil {
        request = NewModifySubscribeConsumeTimeRequest()
    }
    response = NewModifySubscribeConsumeTimeResponse()
    err = c.Send(request, response)
    return
}

func NewModifySubscribeNameRequest() (request *ModifySubscribeNameRequest) {
    request = &ModifySubscribeNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "ModifySubscribeName")
    return
}

func NewModifySubscribeNameResponse() (response *ModifySubscribeNameResponse) {
    response = &ModifySubscribeNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(ModifySubscribeName)用于修改数据订阅实例的名称
func (c *Client) ModifySubscribeName(request *ModifySubscribeNameRequest) (response *ModifySubscribeNameResponse, err error) {
    if request == nil {
        request = NewModifySubscribeNameRequest()
    }
    response = NewModifySubscribeNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifySubscribeObjectsRequest() (request *ModifySubscribeObjectsRequest) {
    request = &ModifySubscribeObjectsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "ModifySubscribeObjects")
    return
}

func NewModifySubscribeObjectsResponse() (response *ModifySubscribeObjectsResponse) {
    response = &ModifySubscribeObjectsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(ModifySubscribeObjects)用于修改数据订阅通道的订阅规则
func (c *Client) ModifySubscribeObjects(request *ModifySubscribeObjectsRequest) (response *ModifySubscribeObjectsResponse, err error) {
    if request == nil {
        request = NewModifySubscribeObjectsRequest()
    }
    response = NewModifySubscribeObjectsResponse()
    err = c.Send(request, response)
    return
}

func NewModifySubscribeVipVportRequest() (request *ModifySubscribeVipVportRequest) {
    request = &ModifySubscribeVipVportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "ModifySubscribeVipVport")
    return
}

func NewModifySubscribeVipVportResponse() (response *ModifySubscribeVipVportResponse) {
    response = &ModifySubscribeVipVportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(ModifySubscribeVipVport)用于修改数据订阅实例的IP和端口号
func (c *Client) ModifySubscribeVipVport(request *ModifySubscribeVipVportRequest) (response *ModifySubscribeVipVportResponse, err error) {
    if request == nil {
        request = NewModifySubscribeVipVportRequest()
    }
    response = NewModifySubscribeVipVportResponse()
    err = c.Send(request, response)
    return
}

func NewModifySyncJobRequest() (request *ModifySyncJobRequest) {
    request = &ModifySyncJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "ModifySyncJob")
    return
}

func NewModifySyncJobResponse() (response *ModifySyncJobResponse) {
    response = &ModifySyncJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改灾备同步任务. 
// 当同步任务处于下述状态时, 允许调用本接口: 同步任务创建中, 创建完成, 校验成功, 校验失败. 
// 源实例和目标实例信息不允许修改，可以修改任务名、需要同步的库表。
func (c *Client) ModifySyncJob(request *ModifySyncJobRequest) (response *ModifySyncJobResponse, err error) {
    if request == nil {
        request = NewModifySyncJobRequest()
    }
    response = NewModifySyncJobResponse()
    err = c.Send(request, response)
    return
}

func NewOfflineIsolatedSubscribeRequest() (request *OfflineIsolatedSubscribeRequest) {
    request = &OfflineIsolatedSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "OfflineIsolatedSubscribe")
    return
}

func NewOfflineIsolatedSubscribeResponse() (response *OfflineIsolatedSubscribeResponse) {
    response = &OfflineIsolatedSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（OfflineIsolatedSubscribe）用于下线已隔离的数据订阅实例
func (c *Client) OfflineIsolatedSubscribe(request *OfflineIsolatedSubscribeRequest) (response *OfflineIsolatedSubscribeResponse, err error) {
    if request == nil {
        request = NewOfflineIsolatedSubscribeRequest()
    }
    response = NewOfflineIsolatedSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewResetSubscribeRequest() (request *ResetSubscribeRequest) {
    request = &ResetSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "ResetSubscribe")
    return
}

func NewResetSubscribeResponse() (response *ResetSubscribeResponse) {
    response = &ResetSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口(ResetSubscribe)用于重置数据订阅实例，已经激活的数据订阅实例，重置后可以使用ActivateSubscribe接口绑定其他的数据库实例
func (c *Client) ResetSubscribe(request *ResetSubscribeRequest) (response *ResetSubscribeResponse, err error) {
    if request == nil {
        request = NewResetSubscribeRequest()
    }
    response = NewResetSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewStartMigrateJobRequest() (request *StartMigrateJobRequest) {
    request = &StartMigrateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "StartMigrateJob")
    return
}

func NewStartMigrateJobResponse() (response *StartMigrateJobResponse) {
    response = &StartMigrateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（StartMigrationJob）用于启动迁移任务。非定时迁移任务会在调用后立即开始迁移，定时任务则会开始倒计时。
// 调用此接口前，请务必先使用CreateMigrateCheckJob校验数据迁移任务，并通过DescribeMigrateJobs接口查询到任务状态为校验通过（status=4）时，才能启动数据迁移任务。
func (c *Client) StartMigrateJob(request *StartMigrateJobRequest) (response *StartMigrateJobResponse, err error) {
    if request == nil {
        request = NewStartMigrateJobRequest()
    }
    response = NewStartMigrateJobResponse()
    err = c.Send(request, response)
    return
}

func NewStartSyncJobRequest() (request *StartSyncJobRequest) {
    request = &StartSyncJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "StartSyncJob")
    return
}

func NewStartSyncJobResponse() (response *StartSyncJobResponse) {
    response = &StartSyncJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建的灾备同步任务在通过 CreateSyncCheckJob 和 DescribeSyncCheckJob 确定校验成功后，可以调用该接口启动同步
func (c *Client) StartSyncJob(request *StartSyncJobRequest) (response *StartSyncJobResponse, err error) {
    if request == nil {
        request = NewStartSyncJobRequest()
    }
    response = NewStartSyncJobResponse()
    err = c.Send(request, response)
    return
}

func NewStopMigrateJobRequest() (request *StopMigrateJobRequest) {
    request = &StopMigrateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "StopMigrateJob")
    return
}

func NewStopMigrateJobResponse() (response *StopMigrateJobResponse) {
    response = &StopMigrateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（StopMigrateJob）用于撤销数据迁移任务。
// 在迁移过程中允许调用该接口撤销迁移, 撤销迁移的任务会失败。通过DescribeMigrateJobs接口查询到任务状态为运行中（status=7）或准备完成（status=8）时，才能撤销数据迁移任务。
func (c *Client) StopMigrateJob(request *StopMigrateJobRequest) (response *StopMigrateJobResponse, err error) {
    if request == nil {
        request = NewStopMigrateJobRequest()
    }
    response = NewStopMigrateJobResponse()
    err = c.Send(request, response)
    return
}

func NewSwitchDrToMasterRequest() (request *SwitchDrToMasterRequest) {
    request = &SwitchDrToMasterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dts", APIVersion, "SwitchDrToMaster")
    return
}

func NewSwitchDrToMasterResponse() (response *SwitchDrToMasterResponse) {
    response = &SwitchDrToMasterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 将灾备升级为主实例，停止从原来所属主实例的同步，断开主备关系。
func (c *Client) SwitchDrToMaster(request *SwitchDrToMasterRequest) (response *SwitchDrToMasterResponse, err error) {
    if request == nil {
        request = NewSwitchDrToMasterRequest()
    }
    response = NewSwitchDrToMasterResponse()
    err = c.Send(request, response)
    return
}
