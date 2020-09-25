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

package v20191112

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-11-12"

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


func NewAttachCcnInstancesRequest() (request *AttachCcnInstancesRequest) {
    request = &AttachCcnInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "AttachCcnInstances")
    return
}

func NewAttachCcnInstancesResponse() (response *AttachCcnInstancesResponse) {
    response = &AttachCcnInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（AttachCcnInstances）用于关联云联网实例
func (c *Client) AttachCcnInstances(request *AttachCcnInstancesRequest) (response *AttachCcnInstancesResponse, err error) {
    if request == nil {
        request = NewAttachCcnInstancesRequest()
    }
    response = NewAttachCcnInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAliasRequest() (request *CreateAliasRequest) {
    request = &CreateAliasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "CreateAlias")
    return
}

func NewCreateAliasResponse() (response *CreateAliasResponse) {
    response = &CreateAliasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateAlias）用于创建别名
func (c *Client) CreateAlias(request *CreateAliasRequest) (response *CreateAliasResponse, err error) {
    if request == nil {
        request = NewCreateAliasRequest()
    }
    response = NewCreateAliasResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAssetRequest() (request *CreateAssetRequest) {
    request = &CreateAssetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "CreateAsset")
    return
}

func NewCreateAssetResponse() (response *CreateAssetResponse) {
    response = &CreateAssetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateAsset）用于创建生成包。
// 通过获取上传cos的临时秘钥，将文件上传至cos，然后将生成包的zip名称下发给[CreateAsset](https://tcloud-dev.oa.com/document/product/1139/46582?!preview&!document=1)完成接口创建。上传文件至 cos支持俩种方式：
// 
//  a.获取预签名， cos 调用上传 （小的文件 5G以内， 前端 1G 以内）  
//   1). [GetUploadCredentials](https://tcloud-dev.oa.com/document/product/1139/39889?!preview&!document=1)  
//   2). 使用 cos API 上传 （cos sdk ）  
//  b.新的方式，适用场景，（大文件）  
//   1). [GetUploadCredentials](https://tcloud-dev.oa.com/document/product/1139/39889?!preview&!document=1) （获取上传 bucket  第一次调用需要，后续可以不用调用 ）  
//   2). GetUploadFederationToken（获取临时密钥）  
//   3). 分块上传 API （cos sdk 有集成 upload_file）  
func (c *Client) CreateAsset(request *CreateAssetRequest) (response *CreateAssetResponse, err error) {
    if request == nil {
        request = NewCreateAssetRequest()
    }
    response = NewCreateAssetResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGameServerSessionRequest() (request *CreateGameServerSessionRequest) {
    request = &CreateGameServerSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "CreateGameServerSession")
    return
}

func NewCreateGameServerSessionResponse() (response *CreateGameServerSessionResponse) {
    response = &CreateGameServerSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateGameServerSession）用于创建游戏服务会话
func (c *Client) CreateGameServerSession(request *CreateGameServerSessionRequest) (response *CreateGameServerSessionResponse, err error) {
    if request == nil {
        request = NewCreateGameServerSessionRequest()
    }
    response = NewCreateGameServerSessionResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAliasRequest() (request *DeleteAliasRequest) {
    request = &DeleteAliasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DeleteAlias")
    return
}

func NewDeleteAliasResponse() (response *DeleteAliasResponse) {
    response = &DeleteAliasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeleteAlias）用于删除别名
func (c *Client) DeleteAlias(request *DeleteAliasRequest) (response *DeleteAliasResponse, err error) {
    if request == nil {
        request = NewDeleteAliasRequest()
    }
    response = NewDeleteAliasResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAssetRequest() (request *DeleteAssetRequest) {
    request = &DeleteAssetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DeleteAsset")
    return
}

func NewDeleteAssetResponse() (response *DeleteAssetResponse) {
    response = &DeleteAssetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeleteAsset）用于删除生成包
func (c *Client) DeleteAsset(request *DeleteAssetRequest) (response *DeleteAssetResponse, err error) {
    if request == nil {
        request = NewDeleteAssetRequest()
    }
    response = NewDeleteAssetResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteFleetRequest() (request *DeleteFleetRequest) {
    request = &DeleteFleetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DeleteFleet")
    return
}

func NewDeleteFleetResponse() (response *DeleteFleetResponse) {
    response = &DeleteFleetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeleteFleet）用于删除服务器舰队
func (c *Client) DeleteFleet(request *DeleteFleetRequest) (response *DeleteFleetResponse, err error) {
    if request == nil {
        request = NewDeleteFleetRequest()
    }
    response = NewDeleteFleetResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteScalingPolicyRequest() (request *DeleteScalingPolicyRequest) {
    request = &DeleteScalingPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DeleteScalingPolicy")
    return
}

func NewDeleteScalingPolicyResponse() (response *DeleteScalingPolicyResponse) {
    response = &DeleteScalingPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeleteScalingPolicy）用于删除扩缩容配置
func (c *Client) DeleteScalingPolicy(request *DeleteScalingPolicyRequest) (response *DeleteScalingPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteScalingPolicyRequest()
    }
    response = NewDeleteScalingPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAliasRequest() (request *DescribeAliasRequest) {
    request = &DescribeAliasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeAlias")
    return
}

func NewDescribeAliasResponse() (response *DescribeAliasResponse) {
    response = &DescribeAliasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeAlias）用于获取别名详情
func (c *Client) DescribeAlias(request *DescribeAliasRequest) (response *DescribeAliasResponse, err error) {
    if request == nil {
        request = NewDescribeAliasRequest()
    }
    response = NewDescribeAliasResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetRequest() (request *DescribeAssetRequest) {
    request = &DescribeAssetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeAsset")
    return
}

func NewDescribeAssetResponse() (response *DescribeAssetResponse) {
    response = &DescribeAssetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeAsset）获取生成包信息
func (c *Client) DescribeAsset(request *DescribeAssetRequest) (response *DescribeAssetResponse, err error) {
    if request == nil {
        request = NewDescribeAssetRequest()
    }
    response = NewDescribeAssetResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetsRequest() (request *DescribeAssetsRequest) {
    request = &DescribeAssetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeAssets")
    return
}

func NewDescribeAssetsResponse() (response *DescribeAssetsResponse) {
    response = &DescribeAssetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeAssets）用于获取生成包列表
func (c *Client) DescribeAssets(request *DescribeAssetsRequest) (response *DescribeAssetsResponse, err error) {
    if request == nil {
        request = NewDescribeAssetsRequest()
    }
    response = NewDescribeAssetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCcnInstancesRequest() (request *DescribeCcnInstancesRequest) {
    request = &DescribeCcnInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeCcnInstances")
    return
}

func NewDescribeCcnInstancesResponse() (response *DescribeCcnInstancesResponse) {
    response = &DescribeCcnInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeCcnInstances）用于查询云联网实例
func (c *Client) DescribeCcnInstances(request *DescribeCcnInstancesRequest) (response *DescribeCcnInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeCcnInstancesRequest()
    }
    response = NewDescribeCcnInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFleetAttributesRequest() (request *DescribeFleetAttributesRequest) {
    request = &DescribeFleetAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeFleetAttributes")
    return
}

func NewDescribeFleetAttributesResponse() (response *DescribeFleetAttributesResponse) {
    response = &DescribeFleetAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeFleetAttributes）用于查询服务器舰队属性
func (c *Client) DescribeFleetAttributes(request *DescribeFleetAttributesRequest) (response *DescribeFleetAttributesResponse, err error) {
    if request == nil {
        request = NewDescribeFleetAttributesRequest()
    }
    response = NewDescribeFleetAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFleetEventsRequest() (request *DescribeFleetEventsRequest) {
    request = &DescribeFleetEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeFleetEvents")
    return
}

func NewDescribeFleetEventsResponse() (response *DescribeFleetEventsResponse) {
    response = &DescribeFleetEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeFleetEvents）用于查询部署服务器舰队相关的事件列表
func (c *Client) DescribeFleetEvents(request *DescribeFleetEventsRequest) (response *DescribeFleetEventsResponse, err error) {
    if request == nil {
        request = NewDescribeFleetEventsRequest()
    }
    response = NewDescribeFleetEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFleetPortSettingsRequest() (request *DescribeFleetPortSettingsRequest) {
    request = &DescribeFleetPortSettingsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeFleetPortSettings")
    return
}

func NewDescribeFleetPortSettingsResponse() (response *DescribeFleetPortSettingsResponse) {
    response = &DescribeFleetPortSettingsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeFleetPortSettings）用于获取服务器舰队安全组信息
func (c *Client) DescribeFleetPortSettings(request *DescribeFleetPortSettingsRequest) (response *DescribeFleetPortSettingsResponse, err error) {
    if request == nil {
        request = NewDescribeFleetPortSettingsRequest()
    }
    response = NewDescribeFleetPortSettingsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFleetUtilizationRequest() (request *DescribeFleetUtilizationRequest) {
    request = &DescribeFleetUtilizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeFleetUtilization")
    return
}

func NewDescribeFleetUtilizationResponse() (response *DescribeFleetUtilizationResponse) {
    response = &DescribeFleetUtilizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeFleetUtilization）用于查询服务器舰队的利用率信息
func (c *Client) DescribeFleetUtilization(request *DescribeFleetUtilizationRequest) (response *DescribeFleetUtilizationResponse, err error) {
    if request == nil {
        request = NewDescribeFleetUtilizationRequest()
    }
    response = NewDescribeFleetUtilizationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGameServerSessionDetailsRequest() (request *DescribeGameServerSessionDetailsRequest) {
    request = &DescribeGameServerSessionDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeGameServerSessionDetails")
    return
}

func NewDescribeGameServerSessionDetailsResponse() (response *DescribeGameServerSessionDetailsResponse) {
    response = &DescribeGameServerSessionDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeGameServerSessionDetails）用于查询游戏服务器会话详情列表
func (c *Client) DescribeGameServerSessionDetails(request *DescribeGameServerSessionDetailsRequest) (response *DescribeGameServerSessionDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeGameServerSessionDetailsRequest()
    }
    response = NewDescribeGameServerSessionDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGameServerSessionPlacementRequest() (request *DescribeGameServerSessionPlacementRequest) {
    request = &DescribeGameServerSessionPlacementRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeGameServerSessionPlacement")
    return
}

func NewDescribeGameServerSessionPlacementResponse() (response *DescribeGameServerSessionPlacementResponse) {
    response = &DescribeGameServerSessionPlacementResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeGameServerSessionPlacement）用于查询游戏服务器会话的放置
func (c *Client) DescribeGameServerSessionPlacement(request *DescribeGameServerSessionPlacementRequest) (response *DescribeGameServerSessionPlacementResponse, err error) {
    if request == nil {
        request = NewDescribeGameServerSessionPlacementRequest()
    }
    response = NewDescribeGameServerSessionPlacementResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGameServerSessionQueuesRequest() (request *DescribeGameServerSessionQueuesRequest) {
    request = &DescribeGameServerSessionQueuesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeGameServerSessionQueues")
    return
}

func NewDescribeGameServerSessionQueuesResponse() (response *DescribeGameServerSessionQueuesResponse) {
    response = &DescribeGameServerSessionQueuesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeGameServerSessionQueues）用于查询游戏服务器会话队列
func (c *Client) DescribeGameServerSessionQueues(request *DescribeGameServerSessionQueuesRequest) (response *DescribeGameServerSessionQueuesResponse, err error) {
    if request == nil {
        request = NewDescribeGameServerSessionQueuesRequest()
    }
    response = NewDescribeGameServerSessionQueuesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGameServerSessionsRequest() (request *DescribeGameServerSessionsRequest) {
    request = &DescribeGameServerSessionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeGameServerSessions")
    return
}

func NewDescribeGameServerSessionsResponse() (response *DescribeGameServerSessionsResponse) {
    response = &DescribeGameServerSessionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeGameServerSessions）用于查询游戏服务器会话列表
func (c *Client) DescribeGameServerSessions(request *DescribeGameServerSessionsRequest) (response *DescribeGameServerSessionsResponse, err error) {
    if request == nil {
        request = NewDescribeGameServerSessionsRequest()
    }
    response = NewDescribeGameServerSessionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceTypesRequest() (request *DescribeInstanceTypesRequest) {
    request = &DescribeInstanceTypesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeInstanceTypes")
    return
}

func NewDescribeInstanceTypesResponse() (response *DescribeInstanceTypesResponse) {
    response = &DescribeInstanceTypesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeInstanceTypes）用于获取服务器实例类型列表
func (c *Client) DescribeInstanceTypes(request *DescribeInstanceTypesRequest) (response *DescribeInstanceTypesResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceTypesRequest()
    }
    response = NewDescribeInstanceTypesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeInstances")
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeInstances）用于查询服务器实例列表
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePlayerSessionsRequest() (request *DescribePlayerSessionsRequest) {
    request = &DescribePlayerSessionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribePlayerSessions")
    return
}

func NewDescribePlayerSessionsResponse() (response *DescribePlayerSessionsResponse) {
    response = &DescribePlayerSessionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribePlayerSessions）用于获取玩家会话列表
func (c *Client) DescribePlayerSessions(request *DescribePlayerSessionsRequest) (response *DescribePlayerSessionsResponse, err error) {
    if request == nil {
        request = NewDescribePlayerSessionsRequest()
    }
    response = NewDescribePlayerSessionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRuntimeConfigurationRequest() (request *DescribeRuntimeConfigurationRequest) {
    request = &DescribeRuntimeConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeRuntimeConfiguration")
    return
}

func NewDescribeRuntimeConfigurationResponse() (response *DescribeRuntimeConfigurationResponse) {
    response = &DescribeRuntimeConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeRuntimeConfiguration）用于获取服务器舰队运行配置
func (c *Client) DescribeRuntimeConfiguration(request *DescribeRuntimeConfigurationRequest) (response *DescribeRuntimeConfigurationResponse, err error) {
    if request == nil {
        request = NewDescribeRuntimeConfigurationRequest()
    }
    response = NewDescribeRuntimeConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScalingPoliciesRequest() (request *DescribeScalingPoliciesRequest) {
    request = &DescribeScalingPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeScalingPolicies")
    return
}

func NewDescribeScalingPoliciesResponse() (response *DescribeScalingPoliciesResponse) {
    response = &DescribeScalingPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeScalingPolicies）用于查询服务部署的动态扩缩容配置
func (c *Client) DescribeScalingPolicies(request *DescribeScalingPoliciesRequest) (response *DescribeScalingPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeScalingPoliciesRequest()
    }
    response = NewDescribeScalingPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserQuotaRequest() (request *DescribeUserQuotaRequest) {
    request = &DescribeUserQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeUserQuota")
    return
}

func NewDescribeUserQuotaResponse() (response *DescribeUserQuotaResponse) {
    response = &DescribeUserQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeUserQuota）获取用户单个模块配额
func (c *Client) DescribeUserQuota(request *DescribeUserQuotaRequest) (response *DescribeUserQuotaResponse, err error) {
    if request == nil {
        request = NewDescribeUserQuotaRequest()
    }
    response = NewDescribeUserQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserQuotasRequest() (request *DescribeUserQuotasRequest) {
    request = &DescribeUserQuotasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeUserQuotas")
    return
}

func NewDescribeUserQuotasResponse() (response *DescribeUserQuotasResponse) {
    response = &DescribeUserQuotasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeUserQuotas）用于获取用户配额
func (c *Client) DescribeUserQuotas(request *DescribeUserQuotasRequest) (response *DescribeUserQuotasResponse, err error) {
    if request == nil {
        request = NewDescribeUserQuotasRequest()
    }
    response = NewDescribeUserQuotasResponse()
    err = c.Send(request, response)
    return
}

func NewDetachCcnInstancesRequest() (request *DetachCcnInstancesRequest) {
    request = &DetachCcnInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DetachCcnInstances")
    return
}

func NewDetachCcnInstancesResponse() (response *DetachCcnInstancesResponse) {
    response = &DetachCcnInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DetachCcnInstances）用于解关联云联网实例
func (c *Client) DetachCcnInstances(request *DetachCcnInstancesRequest) (response *DetachCcnInstancesResponse, err error) {
    if request == nil {
        request = NewDetachCcnInstancesRequest()
    }
    response = NewDetachCcnInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewGetGameServerSessionLogUrlRequest() (request *GetGameServerSessionLogUrlRequest) {
    request = &GetGameServerSessionLogUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "GetGameServerSessionLogUrl")
    return
}

func NewGetGameServerSessionLogUrlResponse() (response *GetGameServerSessionLogUrlResponse) {
    response = &GetGameServerSessionLogUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（GetGameServerSessionLogUrl）用于获取游戏服务器会话的日志URL
func (c *Client) GetGameServerSessionLogUrl(request *GetGameServerSessionLogUrlRequest) (response *GetGameServerSessionLogUrlResponse, err error) {
    if request == nil {
        request = NewGetGameServerSessionLogUrlRequest()
    }
    response = NewGetGameServerSessionLogUrlResponse()
    err = c.Send(request, response)
    return
}

func NewGetInstanceAccessRequest() (request *GetInstanceAccessRequest) {
    request = &GetInstanceAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "GetInstanceAccess")
    return
}

func NewGetInstanceAccessResponse() (response *GetInstanceAccessResponse) {
    response = &GetInstanceAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（GetInstanceAccess）用于获取实例登录所需要的凭据
func (c *Client) GetInstanceAccess(request *GetInstanceAccessRequest) (response *GetInstanceAccessResponse, err error) {
    if request == nil {
        request = NewGetInstanceAccessRequest()
    }
    response = NewGetInstanceAccessResponse()
    err = c.Send(request, response)
    return
}

func NewGetUploadCredentialsRequest() (request *GetUploadCredentialsRequest) {
    request = &GetUploadCredentialsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "GetUploadCredentials")
    return
}

func NewGetUploadCredentialsResponse() (response *GetUploadCredentialsResponse) {
    response = &GetUploadCredentialsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（GetUploadCredentials）获取上传文件授权信息。
// 详细描述：通过[GetUploadCredentials](https://tcloud-dev.oa.com/document/product/1139/39889?!preview&!document=1)接口获取临时秘钥后，使用http put请求将数据上传至cos，然后将生成的生成包zip名称下发给[CreateAsset](https://tcloud-dev.oa.com/document/product/1139/46582?!preview&!document=1)接口进行asset创建
func (c *Client) GetUploadCredentials(request *GetUploadCredentialsRequest) (response *GetUploadCredentialsResponse, err error) {
    if request == nil {
        request = NewGetUploadCredentialsRequest()
    }
    response = NewGetUploadCredentialsResponse()
    err = c.Send(request, response)
    return
}

func NewGetUploadFederationTokenRequest() (request *GetUploadFederationTokenRequest) {
    request = &GetUploadFederationTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "GetUploadFederationToken")
    return
}

func NewGetUploadFederationTokenResponse() (response *GetUploadFederationTokenResponse) {
    response = &GetUploadFederationTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（GetUploadFederationToken）用于 获取生成包上传所需要的临时密钥
func (c *Client) GetUploadFederationToken(request *GetUploadFederationTokenRequest) (response *GetUploadFederationTokenResponse, err error) {
    if request == nil {
        request = NewGetUploadFederationTokenRequest()
    }
    response = NewGetUploadFederationTokenResponse()
    err = c.Send(request, response)
    return
}

func NewJoinGameServerSessionRequest() (request *JoinGameServerSessionRequest) {
    request = &JoinGameServerSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "JoinGameServerSession")
    return
}

func NewJoinGameServerSessionResponse() (response *JoinGameServerSessionResponse) {
    response = &JoinGameServerSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（JoinGameServerSession）用于加入游戏服务器会话
func (c *Client) JoinGameServerSession(request *JoinGameServerSessionRequest) (response *JoinGameServerSessionResponse, err error) {
    if request == nil {
        request = NewJoinGameServerSessionRequest()
    }
    response = NewJoinGameServerSessionResponse()
    err = c.Send(request, response)
    return
}

func NewListAliasesRequest() (request *ListAliasesRequest) {
    request = &ListAliasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "ListAliases")
    return
}

func NewListAliasesResponse() (response *ListAliasesResponse) {
    response = &ListAliasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ListAliases）用于检索帐户下的所有别名
func (c *Client) ListAliases(request *ListAliasesRequest) (response *ListAliasesResponse, err error) {
    if request == nil {
        request = NewListAliasesRequest()
    }
    response = NewListAliasesResponse()
    err = c.Send(request, response)
    return
}

func NewListFleetsRequest() (request *ListFleetsRequest) {
    request = &ListFleetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "ListFleets")
    return
}

func NewListFleetsResponse() (response *ListFleetsResponse) {
    response = &ListFleetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ListFleets）用于获取服务器舰队列表
func (c *Client) ListFleets(request *ListFleetsRequest) (response *ListFleetsResponse, err error) {
    if request == nil {
        request = NewListFleetsRequest()
    }
    response = NewListFleetsResponse()
    err = c.Send(request, response)
    return
}

func NewPutScalingPolicyRequest() (request *PutScalingPolicyRequest) {
    request = &PutScalingPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "PutScalingPolicy")
    return
}

func NewPutScalingPolicyResponse() (response *PutScalingPolicyResponse) {
    response = &PutScalingPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（PutScalingPolicy）用于设置动态扩缩容配置
func (c *Client) PutScalingPolicy(request *PutScalingPolicyRequest) (response *PutScalingPolicyResponse, err error) {
    if request == nil {
        request = NewPutScalingPolicyRequest()
    }
    response = NewPutScalingPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewResolveAliasRequest() (request *ResolveAliasRequest) {
    request = &ResolveAliasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "ResolveAlias")
    return
}

func NewResolveAliasResponse() (response *ResolveAliasResponse) {
    response = &ResolveAliasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ResolveAlias）用于获取别名当前指向的fleetId
func (c *Client) ResolveAlias(request *ResolveAliasRequest) (response *ResolveAliasResponse, err error) {
    if request == nil {
        request = NewResolveAliasRequest()
    }
    response = NewResolveAliasResponse()
    err = c.Send(request, response)
    return
}

func NewSearchGameServerSessionsRequest() (request *SearchGameServerSessionsRequest) {
    request = &SearchGameServerSessionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "SearchGameServerSessions")
    return
}

func NewSearchGameServerSessionsResponse() (response *SearchGameServerSessionsResponse) {
    response = &SearchGameServerSessionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（SearchGameServerSessions）用于搜索游戏服务器会话列表
func (c *Client) SearchGameServerSessions(request *SearchGameServerSessionsRequest) (response *SearchGameServerSessionsResponse, err error) {
    if request == nil {
        request = NewSearchGameServerSessionsRequest()
    }
    response = NewSearchGameServerSessionsResponse()
    err = c.Send(request, response)
    return
}

func NewSetServerWeightRequest() (request *SetServerWeightRequest) {
    request = &SetServerWeightRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "SetServerWeight")
    return
}

func NewSetServerWeightResponse() (response *SetServerWeightResponse) {
    response = &SetServerWeightResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 设置服务器权重
func (c *Client) SetServerWeight(request *SetServerWeightRequest) (response *SetServerWeightResponse, err error) {
    if request == nil {
        request = NewSetServerWeightRequest()
    }
    response = NewSetServerWeightResponse()
    err = c.Send(request, response)
    return
}

func NewStartFleetActionsRequest() (request *StartFleetActionsRequest) {
    request = &StartFleetActionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "StartFleetActions")
    return
}

func NewStartFleetActionsResponse() (response *StartFleetActionsResponse) {
    response = &StartFleetActionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（StartFleetActions）用于启用服务器舰队自动扩缩容
func (c *Client) StartFleetActions(request *StartFleetActionsRequest) (response *StartFleetActionsResponse, err error) {
    if request == nil {
        request = NewStartFleetActionsRequest()
    }
    response = NewStartFleetActionsResponse()
    err = c.Send(request, response)
    return
}

func NewStartGameServerSessionPlacementRequest() (request *StartGameServerSessionPlacementRequest) {
    request = &StartGameServerSessionPlacementRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "StartGameServerSessionPlacement")
    return
}

func NewStartGameServerSessionPlacementResponse() (response *StartGameServerSessionPlacementResponse) {
    response = &StartGameServerSessionPlacementResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（StartGameServerSessionPlacement）用于开始放置游戏服务器会话
func (c *Client) StartGameServerSessionPlacement(request *StartGameServerSessionPlacementRequest) (response *StartGameServerSessionPlacementResponse, err error) {
    if request == nil {
        request = NewStartGameServerSessionPlacementRequest()
    }
    response = NewStartGameServerSessionPlacementResponse()
    err = c.Send(request, response)
    return
}

func NewStartMatchPlacementRequest() (request *StartMatchPlacementRequest) {
    request = &StartMatchPlacementRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "StartMatchPlacement")
    return
}

func NewStartMatchPlacementResponse() (response *StartMatchPlacementResponse) {
    response = &StartMatchPlacementResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（StartMatchPlacement）用于开始匹配放置游戏服务器会话
func (c *Client) StartMatchPlacement(request *StartMatchPlacementRequest) (response *StartMatchPlacementResponse, err error) {
    if request == nil {
        request = NewStartMatchPlacementRequest()
    }
    response = NewStartMatchPlacementResponse()
    err = c.Send(request, response)
    return
}

func NewStopFleetActionsRequest() (request *StopFleetActionsRequest) {
    request = &StopFleetActionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "StopFleetActions")
    return
}

func NewStopFleetActionsResponse() (response *StopFleetActionsResponse) {
    response = &StopFleetActionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（StopFleetActions）用于停止服务器舰队自动扩缩容，改为手动扩缩容
func (c *Client) StopFleetActions(request *StopFleetActionsRequest) (response *StopFleetActionsResponse, err error) {
    if request == nil {
        request = NewStopFleetActionsRequest()
    }
    response = NewStopFleetActionsResponse()
    err = c.Send(request, response)
    return
}

func NewStopGameServerSessionPlacementRequest() (request *StopGameServerSessionPlacementRequest) {
    request = &StopGameServerSessionPlacementRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "StopGameServerSessionPlacement")
    return
}

func NewStopGameServerSessionPlacementResponse() (response *StopGameServerSessionPlacementResponse) {
    response = &StopGameServerSessionPlacementResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（StopGameServerSessionPlacement）用于停止放置游戏服务器会话
func (c *Client) StopGameServerSessionPlacement(request *StopGameServerSessionPlacementRequest) (response *StopGameServerSessionPlacementResponse, err error) {
    if request == nil {
        request = NewStopGameServerSessionPlacementRequest()
    }
    response = NewStopGameServerSessionPlacementResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAliasRequest() (request *UpdateAliasRequest) {
    request = &UpdateAliasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "UpdateAlias")
    return
}

func NewUpdateAliasResponse() (response *UpdateAliasResponse) {
    response = &UpdateAliasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（UpdateAlias）用于更新别名的属性
func (c *Client) UpdateAlias(request *UpdateAliasRequest) (response *UpdateAliasResponse, err error) {
    if request == nil {
        request = NewUpdateAliasRequest()
    }
    response = NewUpdateAliasResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAssetRequest() (request *UpdateAssetRequest) {
    request = &UpdateAssetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "UpdateAsset")
    return
}

func NewUpdateAssetResponse() (response *UpdateAssetResponse) {
    response = &UpdateAssetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（UpdateAsset）用于修改生成包信息
func (c *Client) UpdateAsset(request *UpdateAssetRequest) (response *UpdateAssetResponse, err error) {
    if request == nil {
        request = NewUpdateAssetRequest()
    }
    response = NewUpdateAssetResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateFleetAttributesRequest() (request *UpdateFleetAttributesRequest) {
    request = &UpdateFleetAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "UpdateFleetAttributes")
    return
}

func NewUpdateFleetAttributesResponse() (response *UpdateFleetAttributesResponse) {
    response = &UpdateFleetAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（UpdateFleetAttributes）用于更新服务器舰队属性
func (c *Client) UpdateFleetAttributes(request *UpdateFleetAttributesRequest) (response *UpdateFleetAttributesResponse, err error) {
    if request == nil {
        request = NewUpdateFleetAttributesRequest()
    }
    response = NewUpdateFleetAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateFleetPortSettingsRequest() (request *UpdateFleetPortSettingsRequest) {
    request = &UpdateFleetPortSettingsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "UpdateFleetPortSettings")
    return
}

func NewUpdateFleetPortSettingsResponse() (response *UpdateFleetPortSettingsResponse) {
    response = &UpdateFleetPortSettingsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（UpdateFleetPortSettings）用于更新服务器舰队安全组
func (c *Client) UpdateFleetPortSettings(request *UpdateFleetPortSettingsRequest) (response *UpdateFleetPortSettingsResponse, err error) {
    if request == nil {
        request = NewUpdateFleetPortSettingsRequest()
    }
    response = NewUpdateFleetPortSettingsResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateGameServerSessionRequest() (request *UpdateGameServerSessionRequest) {
    request = &UpdateGameServerSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "UpdateGameServerSession")
    return
}

func NewUpdateGameServerSessionResponse() (response *UpdateGameServerSessionResponse) {
    response = &UpdateGameServerSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（UpdateGameServerSession）用于更新游戏服务器会话
func (c *Client) UpdateGameServerSession(request *UpdateGameServerSessionRequest) (response *UpdateGameServerSessionResponse, err error) {
    if request == nil {
        request = NewUpdateGameServerSessionRequest()
    }
    response = NewUpdateGameServerSessionResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateRuntimeConfigurationRequest() (request *UpdateRuntimeConfigurationRequest) {
    request = &UpdateRuntimeConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "UpdateRuntimeConfiguration")
    return
}

func NewUpdateRuntimeConfigurationResponse() (response *UpdateRuntimeConfigurationResponse) {
    response = &UpdateRuntimeConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（UpdateRuntimeConfiguration）用于更新服务器舰队配置
func (c *Client) UpdateRuntimeConfiguration(request *UpdateRuntimeConfigurationRequest) (response *UpdateRuntimeConfigurationResponse, err error) {
    if request == nil {
        request = NewUpdateRuntimeConfigurationRequest()
    }
    response = NewUpdateRuntimeConfigurationResponse()
    err = c.Send(request, response)
    return
}
