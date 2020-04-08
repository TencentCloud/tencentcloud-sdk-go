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

package v20190719

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-07-19"

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


func NewCreateModuleRequest() (request *CreateModuleRequest) {
    request = &CreateModuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "CreateModule")
    return
}

func NewCreateModuleResponse() (response *CreateModuleResponse) {
    response = &CreateModuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建模块
func (c *Client) CreateModule(request *CreateModuleRequest) (response *CreateModuleResponse, err error) {
    if request == nil {
        request = NewCreateModuleRequest()
    }
    response = NewCreateModuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteImageRequest() (request *DeleteImageRequest) {
    request = &DeleteImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DeleteImage")
    return
}

func NewDeleteImageResponse() (response *DeleteImageResponse) {
    response = &DeleteImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除镜像
func (c *Client) DeleteImage(request *DeleteImageRequest) (response *DeleteImageResponse, err error) {
    if request == nil {
        request = NewDeleteImageRequest()
    }
    response = NewDeleteImageResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteModuleRequest() (request *DeleteModuleRequest) {
    request = &DeleteModuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DeleteModule")
    return
}

func NewDeleteModuleResponse() (response *DeleteModuleResponse) {
    response = &DeleteModuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除业务模块
func (c *Client) DeleteModule(request *DeleteModuleRequest) (response *DeleteModuleResponse, err error) {
    if request == nil {
        request = NewDeleteModuleRequest()
    }
    response = NewDeleteModuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBaseOverviewRequest() (request *DescribeBaseOverviewRequest) {
    request = &DescribeBaseOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeBaseOverview")
    return
}

func NewDescribeBaseOverviewResponse() (response *DescribeBaseOverviewResponse) {
    response = &DescribeBaseOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取概览页统计的基本数据
func (c *Client) DescribeBaseOverview(request *DescribeBaseOverviewRequest) (response *DescribeBaseOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeBaseOverviewRequest()
    }
    response = NewDescribeBaseOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConfigRequest() (request *DescribeConfigRequest) {
    request = &DescribeConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeConfig")
    return
}

func NewDescribeConfigResponse() (response *DescribeConfigResponse) {
    response = &DescribeConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取带宽硬盘等数据的限制
func (c *Client) DescribeConfig(request *DescribeConfigRequest) (response *DescribeConfigResponse, err error) {
    if request == nil {
        request = NewDescribeConfigRequest()
    }
    response = NewDescribeConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageRequest() (request *DescribeImageRequest) {
    request = &DescribeImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeImage")
    return
}

func NewDescribeImageResponse() (response *DescribeImageResponse) {
    response = &DescribeImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 展示镜像列表
func (c *Client) DescribeImage(request *DescribeImageRequest) (response *DescribeImageResponse, err error) {
    if request == nil {
        request = NewDescribeImageRequest()
    }
    response = NewDescribeImageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceTypeConfigRequest() (request *DescribeInstanceTypeConfigRequest) {
    request = &DescribeInstanceTypeConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeInstanceTypeConfig")
    return
}

func NewDescribeInstanceTypeConfigResponse() (response *DescribeInstanceTypeConfigResponse) {
    response = &DescribeInstanceTypeConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取机型配置列表
func (c *Client) DescribeInstanceTypeConfig(request *DescribeInstanceTypeConfigRequest) (response *DescribeInstanceTypeConfigResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceTypeConfigRequest()
    }
    response = NewDescribeInstanceTypeConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeInstances")
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取实例的相关信息。
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesDeniedActionsRequest() (request *DescribeInstancesDeniedActionsRequest) {
    request = &DescribeInstancesDeniedActionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeInstancesDeniedActions")
    return
}

func NewDescribeInstancesDeniedActionsResponse() (response *DescribeInstancesDeniedActionsResponse) {
    response = &DescribeInstancesDeniedActionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 通过实例id获取当前禁止的操作
func (c *Client) DescribeInstancesDeniedActions(request *DescribeInstancesDeniedActionsRequest) (response *DescribeInstancesDeniedActionsResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesDeniedActionsRequest()
    }
    response = NewDescribeInstancesDeniedActionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModuleRequest() (request *DescribeModuleRequest) {
    request = &DescribeModuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeModule")
    return
}

func NewDescribeModuleResponse() (response *DescribeModuleResponse) {
    response = &DescribeModuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取模块列表
func (c *Client) DescribeModule(request *DescribeModuleRequest) (response *DescribeModuleResponse, err error) {
    if request == nil {
        request = NewDescribeModuleRequest()
    }
    response = NewDescribeModuleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModuleDetailRequest() (request *DescribeModuleDetailRequest) {
    request = &DescribeModuleDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeModuleDetail")
    return
}

func NewDescribeModuleDetailResponse() (response *DescribeModuleDetailResponse) {
    response = &DescribeModuleDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 展示模块详细信息
func (c *Client) DescribeModuleDetail(request *DescribeModuleDetailRequest) (response *DescribeModuleDetailResponse, err error) {
    if request == nil {
        request = NewDescribeModuleDetailRequest()
    }
    response = NewDescribeModuleDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNodeRequest() (request *DescribeNodeRequest) {
    request = &DescribeNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribeNode")
    return
}

func NewDescribeNodeResponse() (response *DescribeNodeResponse) {
    response = &DescribeNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取节点列表
func (c *Client) DescribeNode(request *DescribeNodeRequest) (response *DescribeNodeResponse, err error) {
    if request == nil {
        request = NewDescribeNodeRequest()
    }
    response = NewDescribeNodeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePeakBaseOverviewRequest() (request *DescribePeakBaseOverviewRequest) {
    request = &DescribePeakBaseOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribePeakBaseOverview")
    return
}

func NewDescribePeakBaseOverviewResponse() (response *DescribePeakBaseOverviewResponse) {
    response = &DescribePeakBaseOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CPU 内存 硬盘等基础信息峰值数据
func (c *Client) DescribePeakBaseOverview(request *DescribePeakBaseOverviewRequest) (response *DescribePeakBaseOverviewResponse, err error) {
    if request == nil {
        request = NewDescribePeakBaseOverviewRequest()
    }
    response = NewDescribePeakBaseOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePeakNetworkOverviewRequest() (request *DescribePeakNetworkOverviewRequest) {
    request = &DescribePeakNetworkOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "DescribePeakNetworkOverview")
    return
}

func NewDescribePeakNetworkOverviewResponse() (response *DescribePeakNetworkOverviewResponse) {
    response = &DescribePeakNetworkOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取网络峰值数据
func (c *Client) DescribePeakNetworkOverview(request *DescribePeakNetworkOverviewRequest) (response *DescribePeakNetworkOverviewResponse, err error) {
    if request == nil {
        request = NewDescribePeakNetworkOverviewRequest()
    }
    response = NewDescribePeakNetworkOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewImportImageRequest() (request *ImportImageRequest) {
    request = &ImportImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "ImportImage")
    return
}

func NewImportImageResponse() (response *ImportImageResponse) {
    response = &ImportImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 从CVM产品导入镜像到ECM
func (c *Client) ImportImage(request *ImportImageRequest) (response *ImportImageResponse, err error) {
    if request == nil {
        request = NewImportImageRequest()
    }
    response = NewImportImageResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstancesAttributeRequest() (request *ModifyInstancesAttributeRequest) {
    request = &ModifyInstancesAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyInstancesAttribute")
    return
}

func NewModifyInstancesAttributeResponse() (response *ModifyInstancesAttributeResponse) {
    response = &ModifyInstancesAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改实例的属性。
func (c *Client) ModifyInstancesAttribute(request *ModifyInstancesAttributeRequest) (response *ModifyInstancesAttributeResponse, err error) {
    if request == nil {
        request = NewModifyInstancesAttributeRequest()
    }
    response = NewModifyInstancesAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyModuleImageRequest() (request *ModifyModuleImageRequest) {
    request = &ModifyModuleImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyModuleImage")
    return
}

func NewModifyModuleImageResponse() (response *ModifyModuleImageResponse) {
    response = &ModifyModuleImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyModuleImage
func (c *Client) ModifyModuleImage(request *ModifyModuleImageRequest) (response *ModifyModuleImageResponse, err error) {
    if request == nil {
        request = NewModifyModuleImageRequest()
    }
    response = NewModifyModuleImageResponse()
    err = c.Send(request, response)
    return
}

func NewModifyModuleNameRequest() (request *ModifyModuleNameRequest) {
    request = &ModifyModuleNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyModuleName")
    return
}

func NewModifyModuleNameResponse() (response *ModifyModuleNameResponse) {
    response = &ModifyModuleNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改模块名称
func (c *Client) ModifyModuleName(request *ModifyModuleNameRequest) (response *ModifyModuleNameResponse, err error) {
    if request == nil {
        request = NewModifyModuleNameRequest()
    }
    response = NewModifyModuleNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyModuleNetworkRequest() (request *ModifyModuleNetworkRequest) {
    request = &ModifyModuleNetworkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "ModifyModuleNetwork")
    return
}

func NewModifyModuleNetworkResponse() (response *ModifyModuleNetworkResponse) {
    response = &ModifyModuleNetworkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改模块默认带宽上限
func (c *Client) ModifyModuleNetwork(request *ModifyModuleNetworkRequest) (response *ModifyModuleNetworkResponse, err error) {
    if request == nil {
        request = NewModifyModuleNetworkRequest()
    }
    response = NewModifyModuleNetworkResponse()
    err = c.Send(request, response)
    return
}

func NewRebootInstancesRequest() (request *RebootInstancesRequest) {
    request = &RebootInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "RebootInstances")
    return
}

func NewRebootInstancesResponse() (response *RebootInstancesResponse) {
    response = &RebootInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 只有状态为RUNNING的实例才可以进行此操作；接口调用成功时，实例会进入REBOOTING状态；重启实例成功时，实例会进入RUNNING状态；支持强制重启，强制重启的效果等同于关闭物理计算机的电源开关再重新启动。强制重启可能会导致数据丢失或文件系统损坏，请仅在服务器不能正常重启时使用。
func (c *Client) RebootInstances(request *RebootInstancesRequest) (response *RebootInstancesResponse, err error) {
    if request == nil {
        request = NewRebootInstancesRequest()
    }
    response = NewRebootInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewResetInstancesRequest() (request *ResetInstancesRequest) {
    request = &ResetInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "ResetInstances")
    return
}

func NewResetInstancesResponse() (response *ResetInstancesResponse) {
    response = &ResetInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 重装实例，若指定了ImageId参数，则使用指定的镜像重装；否则按照当前实例使用的镜像进行重装；若未指定密码，则密码通过站内信形式随后发送。
func (c *Client) ResetInstances(request *ResetInstancesRequest) (response *ResetInstancesResponse, err error) {
    if request == nil {
        request = NewResetInstancesRequest()
    }
    response = NewResetInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewResetInstancesMaxBandwidthRequest() (request *ResetInstancesMaxBandwidthRequest) {
    request = &ResetInstancesMaxBandwidthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "ResetInstancesMaxBandwidth")
    return
}

func NewResetInstancesMaxBandwidthResponse() (response *ResetInstancesMaxBandwidthResponse) {
    response = &ResetInstancesMaxBandwidthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 重置实例的最大带宽上限。
func (c *Client) ResetInstancesMaxBandwidth(request *ResetInstancesMaxBandwidthRequest) (response *ResetInstancesMaxBandwidthResponse, err error) {
    if request == nil {
        request = NewResetInstancesMaxBandwidthRequest()
    }
    response = NewResetInstancesMaxBandwidthResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateInstancesRequest() (request *TerminateInstancesRequest) {
    request = &TerminateInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecm", APIVersion, "TerminateInstances")
    return
}

func NewTerminateInstancesResponse() (response *TerminateInstancesResponse) {
    response = &TerminateInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 销毁实例
func (c *Client) TerminateInstances(request *TerminateInstancesRequest) (response *TerminateInstancesResponse, err error) {
    if request == nil {
        request = NewTerminateInstancesRequest()
    }
    response = NewTerminateInstancesResponse()
    err = c.Send(request, response)
    return
}
