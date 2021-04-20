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

package v20180419

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-04-19"

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


func NewAttachInstancesRequest() (request *AttachInstancesRequest) {
    request = &AttachInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "AttachInstances")
    return
}

func NewAttachInstancesResponse() (response *AttachInstancesResponse) {
    response = &AttachInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（AttachInstances）用于将 CVM 实例添加到伸缩组。
func (c *Client) AttachInstances(request *AttachInstancesRequest) (response *AttachInstancesResponse, err error) {
    if request == nil {
        request = NewAttachInstancesRequest()
    }
    response = NewAttachInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewClearLaunchConfigurationAttributesRequest() (request *ClearLaunchConfigurationAttributesRequest) {
    request = &ClearLaunchConfigurationAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "ClearLaunchConfigurationAttributes")
    return
}

func NewClearLaunchConfigurationAttributesResponse() (response *ClearLaunchConfigurationAttributesResponse) {
    response = &ClearLaunchConfigurationAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ClearLaunchConfigurationAttributes）用于将启动配置内的特定属性完全清空。
func (c *Client) ClearLaunchConfigurationAttributes(request *ClearLaunchConfigurationAttributesRequest) (response *ClearLaunchConfigurationAttributesResponse, err error) {
    if request == nil {
        request = NewClearLaunchConfigurationAttributesRequest()
    }
    response = NewClearLaunchConfigurationAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewCompleteLifecycleActionRequest() (request *CompleteLifecycleActionRequest) {
    request = &CompleteLifecycleActionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "CompleteLifecycleAction")
    return
}

func NewCompleteLifecycleActionResponse() (response *CompleteLifecycleActionResponse) {
    response = &CompleteLifecycleActionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CompleteLifecycleAction）用于完成生命周期动作。
// 
// * 用户通过调用本接口，指定一个具体的生命周期挂钩的结果（“CONITNUE”或者“ABANDON”）。如果一直不调用本接口，则生命周期挂钩会在超时后按照“DefaultResult”进行处理。
func (c *Client) CompleteLifecycleAction(request *CompleteLifecycleActionRequest) (response *CompleteLifecycleActionResponse, err error) {
    if request == nil {
        request = NewCompleteLifecycleActionRequest()
    }
    response = NewCompleteLifecycleActionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAutoScalingGroupRequest() (request *CreateAutoScalingGroupRequest) {
    request = &CreateAutoScalingGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "CreateAutoScalingGroup")
    return
}

func NewCreateAutoScalingGroupResponse() (response *CreateAutoScalingGroupResponse) {
    response = &CreateAutoScalingGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateAutoScalingGroup）用于创建伸缩组
func (c *Client) CreateAutoScalingGroup(request *CreateAutoScalingGroupRequest) (response *CreateAutoScalingGroupResponse, err error) {
    if request == nil {
        request = NewCreateAutoScalingGroupRequest()
    }
    response = NewCreateAutoScalingGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAutoScalingGroupFromInstanceRequest() (request *CreateAutoScalingGroupFromInstanceRequest) {
    request = &CreateAutoScalingGroupFromInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "CreateAutoScalingGroupFromInstance")
    return
}

func NewCreateAutoScalingGroupFromInstanceResponse() (response *CreateAutoScalingGroupFromInstanceResponse) {
    response = &CreateAutoScalingGroupFromInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateAutoScalingGroupFromInstance）用于根据实例创建启动配置及伸缩组。
// 
// 说明：根据按包年包月计费的实例所创建的伸缩组，其扩容的实例为按量计费实例。
func (c *Client) CreateAutoScalingGroupFromInstance(request *CreateAutoScalingGroupFromInstanceRequest) (response *CreateAutoScalingGroupFromInstanceResponse, err error) {
    if request == nil {
        request = NewCreateAutoScalingGroupFromInstanceRequest()
    }
    response = NewCreateAutoScalingGroupFromInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLaunchConfigurationRequest() (request *CreateLaunchConfigurationRequest) {
    request = &CreateLaunchConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "CreateLaunchConfiguration")
    return
}

func NewCreateLaunchConfigurationResponse() (response *CreateLaunchConfigurationResponse) {
    response = &CreateLaunchConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateLaunchConfiguration）用于创建新的启动配置。
// 
// * 启动配置，可以通过 `ModifyLaunchConfigurationAttributes` 修改少量字段。如需使用新的启动配置，建议重新创建启动配置。
// 
// * 每个项目最多只能创建20个启动配置，详见[使用限制](https://cloud.tencent.com/document/product/377/3120)。
func (c *Client) CreateLaunchConfiguration(request *CreateLaunchConfigurationRequest) (response *CreateLaunchConfigurationResponse, err error) {
    if request == nil {
        request = NewCreateLaunchConfigurationRequest()
    }
    response = NewCreateLaunchConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLifecycleHookRequest() (request *CreateLifecycleHookRequest) {
    request = &CreateLifecycleHookRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "CreateLifecycleHook")
    return
}

func NewCreateLifecycleHookResponse() (response *CreateLifecycleHookResponse) {
    response = &CreateLifecycleHookResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateLifecycleHook）用于创建生命周期挂钩。
// 
// * 您可以为生命周期挂钩配置消息通知，弹性伸缩会通知您的CMQ消息队列，通知内容形如：
// 
// ```
// {
// 	"Service": "Tencent Cloud Auto Scaling",
// 	"Time": "2019-03-14T10:15:11Z",
// 	"AppId": "1251783334",
// 	"ActivityId": "asa-fznnvrja",
// 	"AutoScalingGroupId": "asg-rrrrtttt",
// 	"LifecycleHookId": "ash-xxxxyyyy",
// 	"LifecycleHookName": "my-hook",
// 	"LifecycleActionToken": "3080e1c9-0efe-4dd7-ad3b-90cd6618298f",
// 	"InstanceId": "ins-aaaabbbb",
// 	"LifecycleTransition": "INSTANCE_LAUNCHING",
// 	"NotificationMetadata": ""
// }
// ```
func (c *Client) CreateLifecycleHook(request *CreateLifecycleHookRequest) (response *CreateLifecycleHookResponse, err error) {
    if request == nil {
        request = NewCreateLifecycleHookRequest()
    }
    response = NewCreateLifecycleHookResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNotificationConfigurationRequest() (request *CreateNotificationConfigurationRequest) {
    request = &CreateNotificationConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "CreateNotificationConfiguration")
    return
}

func NewCreateNotificationConfigurationResponse() (response *CreateNotificationConfigurationResponse) {
    response = &CreateNotificationConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateNotificationConfiguration）用于创建通知。
func (c *Client) CreateNotificationConfiguration(request *CreateNotificationConfigurationRequest) (response *CreateNotificationConfigurationResponse, err error) {
    if request == nil {
        request = NewCreateNotificationConfigurationRequest()
    }
    response = NewCreateNotificationConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePaiInstanceRequest() (request *CreatePaiInstanceRequest) {
    request = &CreatePaiInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "CreatePaiInstance")
    return
}

func NewCreatePaiInstanceResponse() (response *CreatePaiInstanceResponse) {
    response = &CreatePaiInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (CreatePaiInstance) 用于创建一个指定配置的PAI实例。
func (c *Client) CreatePaiInstance(request *CreatePaiInstanceRequest) (response *CreatePaiInstanceResponse, err error) {
    if request == nil {
        request = NewCreatePaiInstanceRequest()
    }
    response = NewCreatePaiInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateScalingPolicyRequest() (request *CreateScalingPolicyRequest) {
    request = &CreateScalingPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "CreateScalingPolicy")
    return
}

func NewCreateScalingPolicyResponse() (response *CreateScalingPolicyResponse) {
    response = &CreateScalingPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateScalingPolicy）用于创建告警触发策略。
func (c *Client) CreateScalingPolicy(request *CreateScalingPolicyRequest) (response *CreateScalingPolicyResponse, err error) {
    if request == nil {
        request = NewCreateScalingPolicyRequest()
    }
    response = NewCreateScalingPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateScheduledActionRequest() (request *CreateScheduledActionRequest) {
    request = &CreateScheduledActionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "CreateScheduledAction")
    return
}

func NewCreateScheduledActionResponse() (response *CreateScheduledActionResponse) {
    response = &CreateScheduledActionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateScheduledAction）用于创建定时任务。
func (c *Client) CreateScheduledAction(request *CreateScheduledActionRequest) (response *CreateScheduledActionResponse, err error) {
    if request == nil {
        request = NewCreateScheduledActionRequest()
    }
    response = NewCreateScheduledActionResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAutoScalingGroupRequest() (request *DeleteAutoScalingGroupRequest) {
    request = &DeleteAutoScalingGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "DeleteAutoScalingGroup")
    return
}

func NewDeleteAutoScalingGroupResponse() (response *DeleteAutoScalingGroupResponse) {
    response = &DeleteAutoScalingGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeleteAutoScalingGroup）用于删除指定伸缩组，删除前提是伸缩组内无实例且当前未在执行伸缩活动。
func (c *Client) DeleteAutoScalingGroup(request *DeleteAutoScalingGroupRequest) (response *DeleteAutoScalingGroupResponse, err error) {
    if request == nil {
        request = NewDeleteAutoScalingGroupRequest()
    }
    response = NewDeleteAutoScalingGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLaunchConfigurationRequest() (request *DeleteLaunchConfigurationRequest) {
    request = &DeleteLaunchConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "DeleteLaunchConfiguration")
    return
}

func NewDeleteLaunchConfigurationResponse() (response *DeleteLaunchConfigurationResponse) {
    response = &DeleteLaunchConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeleteLaunchConfiguration）用于删除启动配置。
// 
// * 若启动配置在伸缩组中属于生效状态，则该启动配置不允许删除。
func (c *Client) DeleteLaunchConfiguration(request *DeleteLaunchConfigurationRequest) (response *DeleteLaunchConfigurationResponse, err error) {
    if request == nil {
        request = NewDeleteLaunchConfigurationRequest()
    }
    response = NewDeleteLaunchConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLifecycleHookRequest() (request *DeleteLifecycleHookRequest) {
    request = &DeleteLifecycleHookRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "DeleteLifecycleHook")
    return
}

func NewDeleteLifecycleHookResponse() (response *DeleteLifecycleHookResponse) {
    response = &DeleteLifecycleHookResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeleteLifecycleHook）用于删除生命周期挂钩。
func (c *Client) DeleteLifecycleHook(request *DeleteLifecycleHookRequest) (response *DeleteLifecycleHookResponse, err error) {
    if request == nil {
        request = NewDeleteLifecycleHookRequest()
    }
    response = NewDeleteLifecycleHookResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNotificationConfigurationRequest() (request *DeleteNotificationConfigurationRequest) {
    request = &DeleteNotificationConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "DeleteNotificationConfiguration")
    return
}

func NewDeleteNotificationConfigurationResponse() (response *DeleteNotificationConfigurationResponse) {
    response = &DeleteNotificationConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeleteNotificationConfiguration）用于删除特定的通知。
func (c *Client) DeleteNotificationConfiguration(request *DeleteNotificationConfigurationRequest) (response *DeleteNotificationConfigurationResponse, err error) {
    if request == nil {
        request = NewDeleteNotificationConfigurationRequest()
    }
    response = NewDeleteNotificationConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteScalingPolicyRequest() (request *DeleteScalingPolicyRequest) {
    request = &DeleteScalingPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "DeleteScalingPolicy")
    return
}

func NewDeleteScalingPolicyResponse() (response *DeleteScalingPolicyResponse) {
    response = &DeleteScalingPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeleteScalingPolicy）用于删除告警触发策略。
func (c *Client) DeleteScalingPolicy(request *DeleteScalingPolicyRequest) (response *DeleteScalingPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteScalingPolicyRequest()
    }
    response = NewDeleteScalingPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteScheduledActionRequest() (request *DeleteScheduledActionRequest) {
    request = &DeleteScheduledActionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "DeleteScheduledAction")
    return
}

func NewDeleteScheduledActionResponse() (response *DeleteScheduledActionResponse) {
    response = &DeleteScheduledActionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeleteScheduledAction）用于删除特定的定时任务。
func (c *Client) DeleteScheduledAction(request *DeleteScheduledActionRequest) (response *DeleteScheduledActionResponse, err error) {
    if request == nil {
        request = NewDeleteScheduledActionRequest()
    }
    response = NewDeleteScheduledActionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountLimitsRequest() (request *DescribeAccountLimitsRequest) {
    request = &DescribeAccountLimitsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "DescribeAccountLimits")
    return
}

func NewDescribeAccountLimitsResponse() (response *DescribeAccountLimitsResponse) {
    response = &DescribeAccountLimitsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeAccountLimits）用于查询用户账户在弹性伸缩中的资源限制。
func (c *Client) DescribeAccountLimits(request *DescribeAccountLimitsRequest) (response *DescribeAccountLimitsResponse, err error) {
    if request == nil {
        request = NewDescribeAccountLimitsRequest()
    }
    response = NewDescribeAccountLimitsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAutoScalingActivitiesRequest() (request *DescribeAutoScalingActivitiesRequest) {
    request = &DescribeAutoScalingActivitiesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "DescribeAutoScalingActivities")
    return
}

func NewDescribeAutoScalingActivitiesResponse() (response *DescribeAutoScalingActivitiesResponse) {
    response = &DescribeAutoScalingActivitiesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeAutoScalingActivities）用于查询伸缩组的伸缩活动记录。
func (c *Client) DescribeAutoScalingActivities(request *DescribeAutoScalingActivitiesRequest) (response *DescribeAutoScalingActivitiesResponse, err error) {
    if request == nil {
        request = NewDescribeAutoScalingActivitiesRequest()
    }
    response = NewDescribeAutoScalingActivitiesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAutoScalingGroupLastActivitiesRequest() (request *DescribeAutoScalingGroupLastActivitiesRequest) {
    request = &DescribeAutoScalingGroupLastActivitiesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "DescribeAutoScalingGroupLastActivities")
    return
}

func NewDescribeAutoScalingGroupLastActivitiesResponse() (response *DescribeAutoScalingGroupLastActivitiesResponse) {
    response = &DescribeAutoScalingGroupLastActivitiesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeAutoScalingGroupLastActivities）用于查询伸缩组的最新一次伸缩活动记录。
func (c *Client) DescribeAutoScalingGroupLastActivities(request *DescribeAutoScalingGroupLastActivitiesRequest) (response *DescribeAutoScalingGroupLastActivitiesResponse, err error) {
    if request == nil {
        request = NewDescribeAutoScalingGroupLastActivitiesRequest()
    }
    response = NewDescribeAutoScalingGroupLastActivitiesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAutoScalingGroupsRequest() (request *DescribeAutoScalingGroupsRequest) {
    request = &DescribeAutoScalingGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "DescribeAutoScalingGroups")
    return
}

func NewDescribeAutoScalingGroupsResponse() (response *DescribeAutoScalingGroupsResponse) {
    response = &DescribeAutoScalingGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeAutoScalingGroups）用于查询伸缩组信息。
// 
// * 可以根据伸缩组ID、伸缩组名称或者启动配置ID等信息来查询伸缩组的详细信息。过滤信息详细请见过滤器`Filter`。
// * 如果参数为空，返回当前用户一定数量（`Limit`所指定的数量，默认为20）的伸缩组。
func (c *Client) DescribeAutoScalingGroups(request *DescribeAutoScalingGroupsRequest) (response *DescribeAutoScalingGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeAutoScalingGroupsRequest()
    }
    response = NewDescribeAutoScalingGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAutoScalingInstancesRequest() (request *DescribeAutoScalingInstancesRequest) {
    request = &DescribeAutoScalingInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "DescribeAutoScalingInstances")
    return
}

func NewDescribeAutoScalingInstancesResponse() (response *DescribeAutoScalingInstancesResponse) {
    response = &DescribeAutoScalingInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeAutoScalingInstances）用于查询弹性伸缩关联实例的信息。
// 
// * 可以根据实例ID、伸缩组ID等信息来查询实例的详细信息。过滤信息详细请见过滤器`Filter`。
// * 如果参数为空，返回当前用户一定数量（`Limit`所指定的数量，默认为20）的实例。
func (c *Client) DescribeAutoScalingInstances(request *DescribeAutoScalingInstancesRequest) (response *DescribeAutoScalingInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeAutoScalingInstancesRequest()
    }
    response = NewDescribeAutoScalingInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLaunchConfigurationsRequest() (request *DescribeLaunchConfigurationsRequest) {
    request = &DescribeLaunchConfigurationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "DescribeLaunchConfigurations")
    return
}

func NewDescribeLaunchConfigurationsResponse() (response *DescribeLaunchConfigurationsResponse) {
    response = &DescribeLaunchConfigurationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeLaunchConfigurations）用于查询启动配置的信息。
// 
// * 可以根据启动配置ID、启动配置名称等信息来查询启动配置的详细信息。过滤信息详细请见过滤器`Filter`。
// * 如果参数为空，返回当前用户一定数量（`Limit`所指定的数量，默认为20）的启动配置。
func (c *Client) DescribeLaunchConfigurations(request *DescribeLaunchConfigurationsRequest) (response *DescribeLaunchConfigurationsResponse, err error) {
    if request == nil {
        request = NewDescribeLaunchConfigurationsRequest()
    }
    response = NewDescribeLaunchConfigurationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLifecycleHooksRequest() (request *DescribeLifecycleHooksRequest) {
    request = &DescribeLifecycleHooksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "DescribeLifecycleHooks")
    return
}

func NewDescribeLifecycleHooksResponse() (response *DescribeLifecycleHooksResponse) {
    response = &DescribeLifecycleHooksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeLifecycleHooks）用于查询生命周期挂钩信息。
// 
// * 可以根据伸缩组ID、生命周期挂钩ID或者生命周期挂钩名称等信息来查询生命周期挂钩的详细信息。过滤信息详细请见过滤器`Filter`。
// * 如果参数为空，返回当前用户一定数量（`Limit`所指定的数量，默认为20）的生命周期挂钩。
func (c *Client) DescribeLifecycleHooks(request *DescribeLifecycleHooksRequest) (response *DescribeLifecycleHooksResponse, err error) {
    if request == nil {
        request = NewDescribeLifecycleHooksRequest()
    }
    response = NewDescribeLifecycleHooksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNotificationConfigurationsRequest() (request *DescribeNotificationConfigurationsRequest) {
    request = &DescribeNotificationConfigurationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "DescribeNotificationConfigurations")
    return
}

func NewDescribeNotificationConfigurationsResponse() (response *DescribeNotificationConfigurationsResponse) {
    response = &DescribeNotificationConfigurationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeNotificationConfigurations) 用于查询一个或多个通知的详细信息。
// 
// 可以根据通知ID、伸缩组ID等信息来查询通知的详细信息。过滤信息详细请见过滤器`Filter`。
// 如果参数为空，返回当前用户一定数量（Limit所指定的数量，默认为20）的通知。
func (c *Client) DescribeNotificationConfigurations(request *DescribeNotificationConfigurationsRequest) (response *DescribeNotificationConfigurationsResponse, err error) {
    if request == nil {
        request = NewDescribeNotificationConfigurationsRequest()
    }
    response = NewDescribeNotificationConfigurationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePaiInstancesRequest() (request *DescribePaiInstancesRequest) {
    request = &DescribePaiInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "DescribePaiInstances")
    return
}

func NewDescribePaiInstancesResponse() (response *DescribePaiInstancesResponse) {
    response = &DescribePaiInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribePaiInstances）用于查询PAI实例信息。
// 
// * 可以根据实例ID、实例域名等信息来查询PAI实例的详细信息。过滤信息详细请见过滤器`Filter`。
// * 如果参数为空，返回当前用户一定数量（`Limit`所指定的数量，默认为20）的PAI实例。
func (c *Client) DescribePaiInstances(request *DescribePaiInstancesRequest) (response *DescribePaiInstancesResponse, err error) {
    if request == nil {
        request = NewDescribePaiInstancesRequest()
    }
    response = NewDescribePaiInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScalingPoliciesRequest() (request *DescribeScalingPoliciesRequest) {
    request = &DescribeScalingPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "DescribeScalingPolicies")
    return
}

func NewDescribeScalingPoliciesResponse() (response *DescribeScalingPoliciesResponse) {
    response = &DescribeScalingPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeScalingPolicies）用于查询告警触发策略。
func (c *Client) DescribeScalingPolicies(request *DescribeScalingPoliciesRequest) (response *DescribeScalingPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeScalingPoliciesRequest()
    }
    response = NewDescribeScalingPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScheduledActionsRequest() (request *DescribeScheduledActionsRequest) {
    request = &DescribeScheduledActionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "DescribeScheduledActions")
    return
}

func NewDescribeScheduledActionsResponse() (response *DescribeScheduledActionsResponse) {
    response = &DescribeScheduledActionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeScheduledActions) 用于查询一个或多个定时任务的详细信息。
// 
// * 可以根据定时任务ID、定时任务名称或者伸缩组ID等信息来查询定时任务的详细信息。过滤信息详细请见过滤器`Filter`。
// * 如果参数为空，返回当前用户一定数量（Limit所指定的数量，默认为20）的定时任务。
func (c *Client) DescribeScheduledActions(request *DescribeScheduledActionsRequest) (response *DescribeScheduledActionsResponse, err error) {
    if request == nil {
        request = NewDescribeScheduledActionsRequest()
    }
    response = NewDescribeScheduledActionsResponse()
    err = c.Send(request, response)
    return
}

func NewDetachInstancesRequest() (request *DetachInstancesRequest) {
    request = &DetachInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "DetachInstances")
    return
}

func NewDetachInstancesResponse() (response *DetachInstancesResponse) {
    response = &DetachInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DetachInstances）用于从伸缩组移出 CVM 实例，本接口不会销毁实例。
// * 如果移出指定实例后，伸缩组内处于`IN_SERVICE`状态的实例数量小于伸缩组最小值，接口将报错
// * 如果伸缩组处于`DISABLED`状态，移出操作不校验`IN_SERVICE`实例数量和最小值的关系
// 
// 如果伸缩组关联过 CLB，伸缩组在移出 CVM 实例时，AS 会将 CVM 与该 CLB 进行解挂载。
// 具体来说，该动作是对称的：如果该 CVM 关联进伸缩组时，AS 进行过 CLB 的挂载动作，那么该 CVM 在离开伸缩组时，AS 会进行解挂载动作。
func (c *Client) DetachInstances(request *DetachInstancesRequest) (response *DetachInstancesResponse, err error) {
    if request == nil {
        request = NewDetachInstancesRequest()
    }
    response = NewDetachInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDisableAutoScalingGroupRequest() (request *DisableAutoScalingGroupRequest) {
    request = &DisableAutoScalingGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "DisableAutoScalingGroup")
    return
}

func NewDisableAutoScalingGroupResponse() (response *DisableAutoScalingGroupResponse) {
    response = &DisableAutoScalingGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DisableAutoScalingGroup）用于停用指定伸缩组。
func (c *Client) DisableAutoScalingGroup(request *DisableAutoScalingGroupRequest) (response *DisableAutoScalingGroupResponse, err error) {
    if request == nil {
        request = NewDisableAutoScalingGroupRequest()
    }
    response = NewDisableAutoScalingGroupResponse()
    err = c.Send(request, response)
    return
}

func NewEnableAutoScalingGroupRequest() (request *EnableAutoScalingGroupRequest) {
    request = &EnableAutoScalingGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "EnableAutoScalingGroup")
    return
}

func NewEnableAutoScalingGroupResponse() (response *EnableAutoScalingGroupResponse) {
    response = &EnableAutoScalingGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（EnableAutoScalingGroup）用于启用指定伸缩组。
func (c *Client) EnableAutoScalingGroup(request *EnableAutoScalingGroupRequest) (response *EnableAutoScalingGroupResponse, err error) {
    if request == nil {
        request = NewEnableAutoScalingGroupRequest()
    }
    response = NewEnableAutoScalingGroupResponse()
    err = c.Send(request, response)
    return
}

func NewExecuteScalingPolicyRequest() (request *ExecuteScalingPolicyRequest) {
    request = &ExecuteScalingPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "ExecuteScalingPolicy")
    return
}

func NewExecuteScalingPolicyResponse() (response *ExecuteScalingPolicyResponse) {
    response = &ExecuteScalingPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ExecuteScalingPolicy）用于执行伸缩策略。
// 
// * 可以根据伸缩策略ID执行伸缩策略。
// * 伸缩策略所属伸缩组处于伸缩活动时，会拒绝执行伸缩策略。
func (c *Client) ExecuteScalingPolicy(request *ExecuteScalingPolicyRequest) (response *ExecuteScalingPolicyResponse, err error) {
    if request == nil {
        request = NewExecuteScalingPolicyRequest()
    }
    response = NewExecuteScalingPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAutoScalingGroupRequest() (request *ModifyAutoScalingGroupRequest) {
    request = &ModifyAutoScalingGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "ModifyAutoScalingGroup")
    return
}

func NewModifyAutoScalingGroupResponse() (response *ModifyAutoScalingGroupResponse) {
    response = &ModifyAutoScalingGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyAutoScalingGroup）用于修改伸缩组。
func (c *Client) ModifyAutoScalingGroup(request *ModifyAutoScalingGroupRequest) (response *ModifyAutoScalingGroupResponse, err error) {
    if request == nil {
        request = NewModifyAutoScalingGroupRequest()
    }
    response = NewModifyAutoScalingGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDesiredCapacityRequest() (request *ModifyDesiredCapacityRequest) {
    request = &ModifyDesiredCapacityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "ModifyDesiredCapacity")
    return
}

func NewModifyDesiredCapacityResponse() (response *ModifyDesiredCapacityResponse) {
    response = &ModifyDesiredCapacityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyDesiredCapacity）用于修改指定伸缩组的期望实例数
func (c *Client) ModifyDesiredCapacity(request *ModifyDesiredCapacityRequest) (response *ModifyDesiredCapacityResponse, err error) {
    if request == nil {
        request = NewModifyDesiredCapacityRequest()
    }
    response = NewModifyDesiredCapacityResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLaunchConfigurationAttributesRequest() (request *ModifyLaunchConfigurationAttributesRequest) {
    request = &ModifyLaunchConfigurationAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "ModifyLaunchConfigurationAttributes")
    return
}

func NewModifyLaunchConfigurationAttributesResponse() (response *ModifyLaunchConfigurationAttributesResponse) {
    response = &ModifyLaunchConfigurationAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyLaunchConfigurationAttributes）用于修改启动配置部分属性。
// 
// * 修改启动配置后，已经使用该启动配置扩容的存量实例不会发生变更，此后使用该启动配置的新增实例会按照新的配置进行扩容。
// * 本接口支持修改部分简单类型。
func (c *Client) ModifyLaunchConfigurationAttributes(request *ModifyLaunchConfigurationAttributesRequest) (response *ModifyLaunchConfigurationAttributesResponse, err error) {
    if request == nil {
        request = NewModifyLaunchConfigurationAttributesRequest()
    }
    response = NewModifyLaunchConfigurationAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLoadBalancersRequest() (request *ModifyLoadBalancersRequest) {
    request = &ModifyLoadBalancersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "ModifyLoadBalancers")
    return
}

func NewModifyLoadBalancersResponse() (response *ModifyLoadBalancersResponse) {
    response = &ModifyLoadBalancersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyLoadBalancers）用于修改伸缩组的负载均衡器。
// 
// * 本接口用于为伸缩组指定新的负载均衡器配置，采用“完全覆盖”风格，无论之前配置如何，统一按照接口参数配置为新的负载均衡器。
// * 如果要为伸缩组清空负载均衡器，则在调用本接口时仅指定伸缩组ID，不指定具体负载均衡器。
// * 本接口会立即修改伸缩组的负载均衡器，并生成一个伸缩活动，异步修改存量实例的负载均衡器。
func (c *Client) ModifyLoadBalancers(request *ModifyLoadBalancersRequest) (response *ModifyLoadBalancersResponse, err error) {
    if request == nil {
        request = NewModifyLoadBalancersRequest()
    }
    response = NewModifyLoadBalancersResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNotificationConfigurationRequest() (request *ModifyNotificationConfigurationRequest) {
    request = &ModifyNotificationConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "ModifyNotificationConfiguration")
    return
}

func NewModifyNotificationConfigurationResponse() (response *ModifyNotificationConfigurationResponse) {
    response = &ModifyNotificationConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyNotificationConfiguration）用于修改通知。
func (c *Client) ModifyNotificationConfiguration(request *ModifyNotificationConfigurationRequest) (response *ModifyNotificationConfigurationResponse, err error) {
    if request == nil {
        request = NewModifyNotificationConfigurationRequest()
    }
    response = NewModifyNotificationConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyScalingPolicyRequest() (request *ModifyScalingPolicyRequest) {
    request = &ModifyScalingPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "ModifyScalingPolicy")
    return
}

func NewModifyScalingPolicyResponse() (response *ModifyScalingPolicyResponse) {
    response = &ModifyScalingPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyScalingPolicy）用于修改告警触发策略。
func (c *Client) ModifyScalingPolicy(request *ModifyScalingPolicyRequest) (response *ModifyScalingPolicyResponse, err error) {
    if request == nil {
        request = NewModifyScalingPolicyRequest()
    }
    response = NewModifyScalingPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyScheduledActionRequest() (request *ModifyScheduledActionRequest) {
    request = &ModifyScheduledActionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "ModifyScheduledAction")
    return
}

func NewModifyScheduledActionResponse() (response *ModifyScheduledActionResponse) {
    response = &ModifyScheduledActionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifyScheduledAction）用于修改定时任务。
func (c *Client) ModifyScheduledAction(request *ModifyScheduledActionRequest) (response *ModifyScheduledActionResponse, err error) {
    if request == nil {
        request = NewModifyScheduledActionRequest()
    }
    response = NewModifyScheduledActionResponse()
    err = c.Send(request, response)
    return
}

func NewPreviewPaiDomainNameRequest() (request *PreviewPaiDomainNameRequest) {
    request = &PreviewPaiDomainNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "PreviewPaiDomainName")
    return
}

func NewPreviewPaiDomainNameResponse() (response *PreviewPaiDomainNameResponse) {
    response = &PreviewPaiDomainNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（PreviewPaiDomainName）用于预览PAI实例域名。
func (c *Client) PreviewPaiDomainName(request *PreviewPaiDomainNameRequest) (response *PreviewPaiDomainNameResponse, err error) {
    if request == nil {
        request = NewPreviewPaiDomainNameRequest()
    }
    response = NewPreviewPaiDomainNameResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveInstancesRequest() (request *RemoveInstancesRequest) {
    request = &RemoveInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "RemoveInstances")
    return
}

func NewRemoveInstancesResponse() (response *RemoveInstancesResponse) {
    response = &RemoveInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（RemoveInstances）用于从伸缩组删除 CVM 实例。根据当前的产品逻辑，如果实例由弹性伸缩自动创建，则实例会被销毁；如果实例系创建后加入伸缩组的，则会从伸缩组中移除，保留实例。
// * 如果删除指定实例后，伸缩组内处于`IN_SERVICE`状态的实例数量小于伸缩组最小值，接口将报错
// * 如果伸缩组处于`DISABLED`状态，删除操作不校验`IN_SERVICE`实例数量和最小值的关系
// 
// 如果伸缩组关联过 CLB，伸缩组在删除 CVM 实例时，AS 会将 CVM 与该 CLB 进行解挂载。
// 具体来说，该动作是对称的：如果该 CVM 关联进伸缩组时，AS 进行过 CLB 的挂载动作，那么该 CVM 在离开伸缩组时，AS 会进行解挂载动作。
func (c *Client) RemoveInstances(request *RemoveInstancesRequest) (response *RemoveInstancesResponse, err error) {
    if request == nil {
        request = NewRemoveInstancesRequest()
    }
    response = NewRemoveInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewSetInstancesProtectionRequest() (request *SetInstancesProtectionRequest) {
    request = &SetInstancesProtectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "SetInstancesProtection")
    return
}

func NewSetInstancesProtectionResponse() (response *SetInstancesProtectionResponse) {
    response = &SetInstancesProtectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（SetInstancesProtection）用于设置实例移除保护。
// 子机设置为移除保护之后，当发生不健康替换、报警策略、期望值变更等触发缩容时，将不对此子机缩容操作。
func (c *Client) SetInstancesProtection(request *SetInstancesProtectionRequest) (response *SetInstancesProtectionResponse, err error) {
    if request == nil {
        request = NewSetInstancesProtectionRequest()
    }
    response = NewSetInstancesProtectionResponse()
    err = c.Send(request, response)
    return
}

func NewStartAutoScalingInstancesRequest() (request *StartAutoScalingInstancesRequest) {
    request = &StartAutoScalingInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "StartAutoScalingInstances")
    return
}

func NewStartAutoScalingInstancesResponse() (response *StartAutoScalingInstancesResponse) {
    response = &StartAutoScalingInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（StartAutoScalingInstances）用于开启伸缩组内 CVM 实例。
// * 开机成功，实例转为`IN_SERVICE`状态后，会增加期望实例数，期望实例数不可超过设置的最大值
// * 本接口支持批量操作，每次请求开机实例的上限为100
func (c *Client) StartAutoScalingInstances(request *StartAutoScalingInstancesRequest) (response *StartAutoScalingInstancesResponse, err error) {
    if request == nil {
        request = NewStartAutoScalingInstancesRequest()
    }
    response = NewStartAutoScalingInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewStopAutoScalingInstancesRequest() (request *StopAutoScalingInstancesRequest) {
    request = &StopAutoScalingInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "StopAutoScalingInstances")
    return
}

func NewStopAutoScalingInstancesResponse() (response *StopAutoScalingInstancesResponse) {
    response = &StopAutoScalingInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（StopAutoScalingInstances）用于关闭伸缩组内 CVM 实例。
// * 关机方式采用`SOFT_FIRST`方式，表示在正常关闭失败后进行强制关闭
// * 关闭`IN_SERVICE`状态的实例，会减少期望实例数，期望实例数不可低于设置的最小值
// * 使用`STOP_CHARGING`选项关机，待关机的实例需要满足[关机不收费条件](https://cloud.tencent.com/document/product/213/19918)
// * 本接口支持批量操作，每次请求关机实例的上限为100
func (c *Client) StopAutoScalingInstances(request *StopAutoScalingInstancesRequest) (response *StopAutoScalingInstancesResponse, err error) {
    if request == nil {
        request = NewStopAutoScalingInstancesRequest()
    }
    response = NewStopAutoScalingInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeLaunchConfigurationRequest() (request *UpgradeLaunchConfigurationRequest) {
    request = &UpgradeLaunchConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "UpgradeLaunchConfiguration")
    return
}

func NewUpgradeLaunchConfigurationResponse() (response *UpgradeLaunchConfigurationResponse) {
    response = &UpgradeLaunchConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（UpgradeLaunchConfiguration）用于升级启动配置。
// 
// * 本接口用于升级启动配置，采用“完全覆盖”风格，无论之前参数如何，统一按照接口参数设置为新的配置。对于非必填字段，不填写则按照默认值赋值。
// * 升级修改启动配置后，已经使用该启动配置扩容的存量实例不会发生变更，此后使用该启动配置的新增实例会按照新的配置进行扩容。
func (c *Client) UpgradeLaunchConfiguration(request *UpgradeLaunchConfigurationRequest) (response *UpgradeLaunchConfigurationResponse, err error) {
    if request == nil {
        request = NewUpgradeLaunchConfigurationRequest()
    }
    response = NewUpgradeLaunchConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeLifecycleHookRequest() (request *UpgradeLifecycleHookRequest) {
    request = &UpgradeLifecycleHookRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("as", APIVersion, "UpgradeLifecycleHook")
    return
}

func NewUpgradeLifecycleHookResponse() (response *UpgradeLifecycleHookResponse) {
    response = &UpgradeLifecycleHookResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（UpgradeLifecycleHook）用于升级生命周期挂钩。
// 
// * 本接口用于升级生命周期挂钩，采用“完全覆盖”风格，无论之前参数如何，统一按照接口参数设置为新的配置。对于非必填字段，不填写则按照默认值赋值。
func (c *Client) UpgradeLifecycleHook(request *UpgradeLifecycleHookRequest) (response *UpgradeLifecycleHookResponse, err error) {
    if request == nil {
        request = NewUpgradeLifecycleHookRequest()
    }
    response = NewUpgradeLifecycleHookResponse()
    err = c.Send(request, response)
    return
}
