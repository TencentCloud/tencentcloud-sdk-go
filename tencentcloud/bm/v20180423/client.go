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

package v20180423

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-04-23"

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


func NewAttachCamRoleRequest() (request *AttachCamRoleRequest) {
    request = &AttachCamRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "AttachCamRole")
    
    
    return
}

func NewAttachCamRoleResponse() (response *AttachCamRoleResponse) {
    response = &AttachCamRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AttachCamRole
// 服务器绑定CAM角色，该角色授权访问黑石物理服务器服务，为黑石物理服务器提供了访问资源的权限，如请求服务器的临时证书
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AttachCamRole(request *AttachCamRoleRequest) (response *AttachCamRoleResponse, err error) {
    return c.AttachCamRoleWithContext(context.Background(), request)
}

// AttachCamRole
// 服务器绑定CAM角色，该角色授权访问黑石物理服务器服务，为黑石物理服务器提供了访问资源的权限，如请求服务器的临时证书
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AttachCamRoleWithContext(ctx context.Context, request *AttachCamRoleRequest) (response *AttachCamRoleResponse, err error) {
    if request == nil {
        request = NewAttachCamRoleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AttachCamRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewAttachCamRoleResponse()
    err = c.Send(request, response)
    return
}

func NewBindPsaTagRequest() (request *BindPsaTagRequest) {
    request = &BindPsaTagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "BindPsaTag")
    
    
    return
}

func NewBindPsaTagResponse() (response *BindPsaTagResponse) {
    response = &BindPsaTagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindPsaTag
// 为预授权规则绑定标签
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETER = "InvalidParameterValue.InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) BindPsaTag(request *BindPsaTagRequest) (response *BindPsaTagResponse, err error) {
    return c.BindPsaTagWithContext(context.Background(), request)
}

// BindPsaTag
// 为预授权规则绑定标签
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETER = "InvalidParameterValue.InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) BindPsaTagWithContext(ctx context.Context, request *BindPsaTagRequest) (response *BindPsaTagResponse, err error) {
    if request == nil {
        request = NewBindPsaTagRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindPsaTag require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindPsaTagResponse()
    err = c.Send(request, response)
    return
}

func NewBuyDevicesRequest() (request *BuyDevicesRequest) {
    request = &BuyDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "BuyDevices")
    
    
    return
}

func NewBuyDevicesResponse() (response *BuyDevicesResponse) {
    response = &BuyDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BuyDevices
// 购买黑石物理机
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE_FLOWBUSY = "ResourceInUse.FlowBusy"
//  RESOURCEINSUFFICIENT_DEVICEINSUFFICIENT = "ResourceInsufficient.DeviceInsufficient"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_FUNDINSUFFICIENT = "UnsupportedOperation.FundInsufficient"
func (c *Client) BuyDevices(request *BuyDevicesRequest) (response *BuyDevicesResponse, err error) {
    return c.BuyDevicesWithContext(context.Background(), request)
}

// BuyDevices
// 购买黑石物理机
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE_FLOWBUSY = "ResourceInUse.FlowBusy"
//  RESOURCEINSUFFICIENT_DEVICEINSUFFICIENT = "ResourceInsufficient.DeviceInsufficient"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_FUNDINSUFFICIENT = "UnsupportedOperation.FundInsufficient"
func (c *Client) BuyDevicesWithContext(ctx context.Context, request *BuyDevicesRequest) (response *BuyDevicesResponse, err error) {
    if request == nil {
        request = NewBuyDevicesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BuyDevices require credential")
    }

    request.SetContext(ctx)
    
    response = NewBuyDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCustomImageRequest() (request *CreateCustomImageRequest) {
    request = &CreateCustomImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "CreateCustomImage")
    
    
    return
}

func NewCreateCustomImageResponse() (response *CreateCustomImageResponse) {
    response = &CreateCustomImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCustomImage
// 创建自定义镜像<br>
//
// 每个AppId在每个可用区最多保留20个自定义镜像
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCustomImage(request *CreateCustomImageRequest) (response *CreateCustomImageResponse, err error) {
    return c.CreateCustomImageWithContext(context.Background(), request)
}

// CreateCustomImage
// 创建自定义镜像<br>
//
// 每个AppId在每个可用区最多保留20个自定义镜像
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCustomImageWithContext(ctx context.Context, request *CreateCustomImageRequest) (response *CreateCustomImageResponse, err error) {
    if request == nil {
        request = NewCreateCustomImageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCustomImage require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCustomImageResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePsaRegulationRequest() (request *CreatePsaRegulationRequest) {
    request = &CreatePsaRegulationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "CreatePsaRegulation")
    
    
    return
}

func NewCreatePsaRegulationResponse() (response *CreatePsaRegulationResponse) {
    response = &CreatePsaRegulationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePsaRegulation
// 创建预授权规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreatePsaRegulation(request *CreatePsaRegulationRequest) (response *CreatePsaRegulationResponse, err error) {
    return c.CreatePsaRegulationWithContext(context.Background(), request)
}

// CreatePsaRegulation
// 创建预授权规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreatePsaRegulationWithContext(ctx context.Context, request *CreatePsaRegulationRequest) (response *CreatePsaRegulationResponse, err error) {
    if request == nil {
        request = NewCreatePsaRegulationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePsaRegulation require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePsaRegulationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSpotDeviceRequest() (request *CreateSpotDeviceRequest) {
    request = &CreateSpotDeviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "CreateSpotDevice")
    
    
    return
}

func NewCreateSpotDeviceResponse() (response *CreateSpotDeviceResponse) {
    response = &CreateSpotDeviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSpotDevice
// 创建黑石竞价实例
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) CreateSpotDevice(request *CreateSpotDeviceRequest) (response *CreateSpotDeviceResponse, err error) {
    return c.CreateSpotDeviceWithContext(context.Background(), request)
}

// CreateSpotDevice
// 创建黑石竞价实例
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) CreateSpotDeviceWithContext(ctx context.Context, request *CreateSpotDeviceRequest) (response *CreateSpotDeviceResponse, err error) {
    if request == nil {
        request = NewCreateSpotDeviceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSpotDevice require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSpotDeviceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserCmdRequest() (request *CreateUserCmdRequest) {
    request = &CreateUserCmdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "CreateUserCmd")
    
    
    return
}

func NewCreateUserCmdResponse() (response *CreateUserCmdResponse) {
    response = &CreateUserCmdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateUserCmd
// 创建自定义脚本
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_USERCMDCOUNT = "LimitExceeded.UserCmdCount"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateUserCmd(request *CreateUserCmdRequest) (response *CreateUserCmdResponse, err error) {
    return c.CreateUserCmdWithContext(context.Background(), request)
}

// CreateUserCmd
// 创建自定义脚本
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_USERCMDCOUNT = "LimitExceeded.UserCmdCount"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateUserCmdWithContext(ctx context.Context, request *CreateUserCmdRequest) (response *CreateUserCmdResponse, err error) {
    if request == nil {
        request = NewCreateUserCmdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUserCmd require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUserCmdResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCustomImagesRequest() (request *DeleteCustomImagesRequest) {
    request = &DeleteCustomImagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "DeleteCustomImages")
    
    
    return
}

func NewDeleteCustomImagesResponse() (response *DeleteCustomImagesResponse) {
    response = &DeleteCustomImagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCustomImages
// 删除自定义镜像<br>
//
// 正用于部署或重装中的镜像被删除后，镜像文件将保留一段时间，直到部署或重装结束
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteCustomImages(request *DeleteCustomImagesRequest) (response *DeleteCustomImagesResponse, err error) {
    return c.DeleteCustomImagesWithContext(context.Background(), request)
}

// DeleteCustomImages
// 删除自定义镜像<br>
//
// 正用于部署或重装中的镜像被删除后，镜像文件将保留一段时间，直到部署或重装结束
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteCustomImagesWithContext(ctx context.Context, request *DeleteCustomImagesRequest) (response *DeleteCustomImagesResponse, err error) {
    if request == nil {
        request = NewDeleteCustomImagesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCustomImages require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCustomImagesResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePsaRegulationRequest() (request *DeletePsaRegulationRequest) {
    request = &DeletePsaRegulationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "DeletePsaRegulation")
    
    
    return
}

func NewDeletePsaRegulationResponse() (response *DeletePsaRegulationResponse) {
    response = &DeletePsaRegulationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeletePsaRegulation
// 删除预授权规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeletePsaRegulation(request *DeletePsaRegulationRequest) (response *DeletePsaRegulationResponse, err error) {
    return c.DeletePsaRegulationWithContext(context.Background(), request)
}

// DeletePsaRegulation
// 删除预授权规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeletePsaRegulationWithContext(ctx context.Context, request *DeletePsaRegulationRequest) (response *DeletePsaRegulationResponse, err error) {
    if request == nil {
        request = NewDeletePsaRegulationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePsaRegulation require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePsaRegulationResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUserCmdsRequest() (request *DeleteUserCmdsRequest) {
    request = &DeleteUserCmdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "DeleteUserCmds")
    
    
    return
}

func NewDeleteUserCmdsResponse() (response *DeleteUserCmdsResponse) {
    response = &DeleteUserCmdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteUserCmds
// 删除自定义脚本
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteUserCmds(request *DeleteUserCmdsRequest) (response *DeleteUserCmdsResponse, err error) {
    return c.DeleteUserCmdsWithContext(context.Background(), request)
}

// DeleteUserCmds
// 删除自定义脚本
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteUserCmdsWithContext(ctx context.Context, request *DeleteUserCmdsRequest) (response *DeleteUserCmdsResponse, err error) {
    if request == nil {
        request = NewDeleteUserCmdsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUserCmds require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUserCmdsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomImageProcessRequest() (request *DescribeCustomImageProcessRequest) {
    request = &DescribeCustomImageProcessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "DescribeCustomImageProcess")
    
    
    return
}

func NewDescribeCustomImageProcessResponse() (response *DescribeCustomImageProcessResponse) {
    response = &DescribeCustomImageProcessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCustomImageProcess
// 查询自定义镜像制作进度
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeCustomImageProcess(request *DescribeCustomImageProcessRequest) (response *DescribeCustomImageProcessResponse, err error) {
    return c.DescribeCustomImageProcessWithContext(context.Background(), request)
}

// DescribeCustomImageProcess
// 查询自定义镜像制作进度
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeCustomImageProcessWithContext(ctx context.Context, request *DescribeCustomImageProcessRequest) (response *DescribeCustomImageProcessResponse, err error) {
    if request == nil {
        request = NewDescribeCustomImageProcessRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCustomImageProcess require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCustomImageProcessResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCustomImagesRequest() (request *DescribeCustomImagesRequest) {
    request = &DescribeCustomImagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "DescribeCustomImages")
    
    
    return
}

func NewDescribeCustomImagesResponse() (response *DescribeCustomImagesResponse) {
    response = &DescribeCustomImagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCustomImages
// 查看自定义镜像列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeCustomImages(request *DescribeCustomImagesRequest) (response *DescribeCustomImagesResponse, err error) {
    return c.DescribeCustomImagesWithContext(context.Background(), request)
}

// DescribeCustomImages
// 查看自定义镜像列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeCustomImagesWithContext(ctx context.Context, request *DescribeCustomImagesRequest) (response *DescribeCustomImagesResponse, err error) {
    if request == nil {
        request = NewDescribeCustomImagesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCustomImages require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCustomImagesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceClassRequest() (request *DescribeDeviceClassRequest) {
    request = &DescribeDeviceClassRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "DescribeDeviceClass")
    
    
    return
}

func NewDescribeDeviceClassResponse() (response *DescribeDeviceClassResponse) {
    response = &DescribeDeviceClassResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDeviceClass
// 获取设备类型
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDeviceClass(request *DescribeDeviceClassRequest) (response *DescribeDeviceClassResponse, err error) {
    return c.DescribeDeviceClassWithContext(context.Background(), request)
}

// DescribeDeviceClass
// 获取设备类型
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDeviceClassWithContext(ctx context.Context, request *DescribeDeviceClassRequest) (response *DescribeDeviceClassResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceClassRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceClass require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceClassResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceClassPartitionRequest() (request *DescribeDeviceClassPartitionRequest) {
    request = &DescribeDeviceClassPartitionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "DescribeDeviceClassPartition")
    
    
    return
}

func NewDescribeDeviceClassPartitionResponse() (response *DescribeDeviceClassPartitionResponse) {
    response = &DescribeDeviceClassPartitionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDeviceClassPartition
// 查询机型支持的RAID方式， 并返回系统盘的分区和逻辑盘的列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeDeviceClassPartition(request *DescribeDeviceClassPartitionRequest) (response *DescribeDeviceClassPartitionResponse, err error) {
    return c.DescribeDeviceClassPartitionWithContext(context.Background(), request)
}

// DescribeDeviceClassPartition
// 查询机型支持的RAID方式， 并返回系统盘的分区和逻辑盘的列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeDeviceClassPartitionWithContext(ctx context.Context, request *DescribeDeviceClassPartitionRequest) (response *DescribeDeviceClassPartitionResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceClassPartitionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceClassPartition require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceClassPartitionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceHardwareInfoRequest() (request *DescribeDeviceHardwareInfoRequest) {
    request = &DescribeDeviceHardwareInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "DescribeDeviceHardwareInfo")
    
    
    return
}

func NewDescribeDeviceHardwareInfoResponse() (response *DescribeDeviceHardwareInfoResponse) {
    response = &DescribeDeviceHardwareInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDeviceHardwareInfo
// 查询设备硬件配置信息，如 CPU 型号，内存大小，磁盘大小和数量
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDeviceHardwareInfo(request *DescribeDeviceHardwareInfoRequest) (response *DescribeDeviceHardwareInfoResponse, err error) {
    return c.DescribeDeviceHardwareInfoWithContext(context.Background(), request)
}

// DescribeDeviceHardwareInfo
// 查询设备硬件配置信息，如 CPU 型号，内存大小，磁盘大小和数量
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDeviceHardwareInfoWithContext(ctx context.Context, request *DescribeDeviceHardwareInfoRequest) (response *DescribeDeviceHardwareInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceHardwareInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceHardwareInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceHardwareInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceInventoryRequest() (request *DescribeDeviceInventoryRequest) {
    request = &DescribeDeviceInventoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "DescribeDeviceInventory")
    
    
    return
}

func NewDescribeDeviceInventoryResponse() (response *DescribeDeviceInventoryResponse) {
    response = &DescribeDeviceInventoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDeviceInventory
// 查询设备库存
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeDeviceInventory(request *DescribeDeviceInventoryRequest) (response *DescribeDeviceInventoryResponse, err error) {
    return c.DescribeDeviceInventoryWithContext(context.Background(), request)
}

// DescribeDeviceInventory
// 查询设备库存
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeDeviceInventoryWithContext(ctx context.Context, request *DescribeDeviceInventoryRequest) (response *DescribeDeviceInventoryResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceInventoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceInventory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceInventoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceOperationLogRequest() (request *DescribeDeviceOperationLogRequest) {
    request = &DescribeDeviceOperationLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "DescribeDeviceOperationLog")
    
    
    return
}

func NewDescribeDeviceOperationLogResponse() (response *DescribeDeviceOperationLogResponse) {
    response = &DescribeDeviceOperationLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDeviceOperationLog
// 查询设备操作日志， 如设备重启，重装，设置密码等操作
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDeviceOperationLog(request *DescribeDeviceOperationLogRequest) (response *DescribeDeviceOperationLogResponse, err error) {
    return c.DescribeDeviceOperationLogWithContext(context.Background(), request)
}

// DescribeDeviceOperationLog
// 查询设备操作日志， 如设备重启，重装，设置密码等操作
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDeviceOperationLogWithContext(ctx context.Context, request *DescribeDeviceOperationLogRequest) (response *DescribeDeviceOperationLogResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceOperationLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceOperationLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceOperationLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDevicePartitionRequest() (request *DescribeDevicePartitionRequest) {
    request = &DescribeDevicePartitionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "DescribeDevicePartition")
    
    
    return
}

func NewDescribeDevicePartitionResponse() (response *DescribeDevicePartitionResponse) {
    response = &DescribeDevicePartitionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDevicePartition
// 获取物理机的分区格式
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDevicePartition(request *DescribeDevicePartitionRequest) (response *DescribeDevicePartitionResponse, err error) {
    return c.DescribeDevicePartitionWithContext(context.Background(), request)
}

// DescribeDevicePartition
// 获取物理机的分区格式
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDevicePartitionWithContext(ctx context.Context, request *DescribeDevicePartitionRequest) (response *DescribeDevicePartitionResponse, err error) {
    if request == nil {
        request = NewDescribeDevicePartitionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDevicePartition require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDevicePartitionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDevicePositionRequest() (request *DescribeDevicePositionRequest) {
    request = &DescribeDevicePositionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "DescribeDevicePosition")
    
    
    return
}

func NewDescribeDevicePositionResponse() (response *DescribeDevicePositionResponse) {
    response = &DescribeDevicePositionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDevicePosition
// 查询服务器所在的位置，如机架，上联交换机等信息
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDevicePosition(request *DescribeDevicePositionRequest) (response *DescribeDevicePositionResponse, err error) {
    return c.DescribeDevicePositionWithContext(context.Background(), request)
}

// DescribeDevicePosition
// 查询服务器所在的位置，如机架，上联交换机等信息
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDevicePositionWithContext(ctx context.Context, request *DescribeDevicePositionRequest) (response *DescribeDevicePositionResponse, err error) {
    if request == nil {
        request = NewDescribeDevicePositionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDevicePosition require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDevicePositionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDevicePriceInfoRequest() (request *DescribeDevicePriceInfoRequest) {
    request = &DescribeDevicePriceInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "DescribeDevicePriceInfo")
    
    
    return
}

func NewDescribeDevicePriceInfoResponse() (response *DescribeDevicePriceInfoResponse) {
    response = &DescribeDevicePriceInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDevicePriceInfo
// 查询服务器价格信息，支持设备的批量查找，支持标准机型和弹性机型的混合查找
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDevicePriceInfo(request *DescribeDevicePriceInfoRequest) (response *DescribeDevicePriceInfoResponse, err error) {
    return c.DescribeDevicePriceInfoWithContext(context.Background(), request)
}

// DescribeDevicePriceInfo
// 查询服务器价格信息，支持设备的批量查找，支持标准机型和弹性机型的混合查找
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDevicePriceInfoWithContext(ctx context.Context, request *DescribeDevicePriceInfoRequest) (response *DescribeDevicePriceInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDevicePriceInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDevicePriceInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDevicePriceInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDevicesRequest() (request *DescribeDevicesRequest) {
    request = &DescribeDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "DescribeDevices")
    
    
    return
}

func NewDescribeDevicesResponse() (response *DescribeDevicesResponse) {
    response = &DescribeDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDevices
// 查询物理服务器，可以按照实例，业务IP等过滤
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDevices(request *DescribeDevicesRequest) (response *DescribeDevicesResponse, err error) {
    return c.DescribeDevicesWithContext(context.Background(), request)
}

// DescribeDevices
// 查询物理服务器，可以按照实例，业务IP等过滤
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDevicesWithContext(ctx context.Context, request *DescribeDevicesRequest) (response *DescribeDevicesResponse, err error) {
    if request == nil {
        request = NewDescribeDevicesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDevices require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHardwareSpecificationRequest() (request *DescribeHardwareSpecificationRequest) {
    request = &DescribeHardwareSpecificationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "DescribeHardwareSpecification")
    
    
    return
}

func NewDescribeHardwareSpecificationResponse() (response *DescribeHardwareSpecificationResponse) {
    response = &DescribeHardwareSpecificationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeHardwareSpecification
// 查询自定义机型部件信息，包括CpuId对应的型号，DiskTypeId对应的磁盘类型
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeHardwareSpecification(request *DescribeHardwareSpecificationRequest) (response *DescribeHardwareSpecificationResponse, err error) {
    return c.DescribeHardwareSpecificationWithContext(context.Background(), request)
}

// DescribeHardwareSpecification
// 查询自定义机型部件信息，包括CpuId对应的型号，DiskTypeId对应的磁盘类型
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeHardwareSpecificationWithContext(ctx context.Context, request *DescribeHardwareSpecificationRequest) (response *DescribeHardwareSpecificationResponse, err error) {
    if request == nil {
        request = NewDescribeHardwareSpecificationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHardwareSpecification require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHardwareSpecificationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHostedDeviceOutBandInfoRequest() (request *DescribeHostedDeviceOutBandInfoRequest) {
    request = &DescribeHostedDeviceOutBandInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "DescribeHostedDeviceOutBandInfo")
    
    
    return
}

func NewDescribeHostedDeviceOutBandInfoResponse() (response *DescribeHostedDeviceOutBandInfoResponse) {
    response = &DescribeHostedDeviceOutBandInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeHostedDeviceOutBandInfo
// 查询托管设备带外信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeHostedDeviceOutBandInfo(request *DescribeHostedDeviceOutBandInfoRequest) (response *DescribeHostedDeviceOutBandInfoResponse, err error) {
    return c.DescribeHostedDeviceOutBandInfoWithContext(context.Background(), request)
}

// DescribeHostedDeviceOutBandInfo
// 查询托管设备带外信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeHostedDeviceOutBandInfoWithContext(ctx context.Context, request *DescribeHostedDeviceOutBandInfoRequest) (response *DescribeHostedDeviceOutBandInfoResponse, err error) {
    if request == nil {
        request = NewDescribeHostedDeviceOutBandInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHostedDeviceOutBandInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHostedDeviceOutBandInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOperationResultRequest() (request *DescribeOperationResultRequest) {
    request = &DescribeOperationResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "DescribeOperationResult")
    
    
    return
}

func NewDescribeOperationResultResponse() (response *DescribeOperationResultResponse) {
    response = &DescribeOperationResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOperationResult
// 获取异步操作状态的完成状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeOperationResult(request *DescribeOperationResultRequest) (response *DescribeOperationResultResponse, err error) {
    return c.DescribeOperationResultWithContext(context.Background(), request)
}

// DescribeOperationResult
// 获取异步操作状态的完成状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeOperationResultWithContext(ctx context.Context, request *DescribeOperationResultRequest) (response *DescribeOperationResultResponse, err error) {
    if request == nil {
        request = NewDescribeOperationResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOperationResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOperationResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOsInfoRequest() (request *DescribeOsInfoRequest) {
    request = &DescribeOsInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "DescribeOsInfo")
    
    
    return
}

func NewDescribeOsInfoResponse() (response *DescribeOsInfoResponse) {
    response = &DescribeOsInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOsInfo
// 查询指定机型所支持的操作系统
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeOsInfo(request *DescribeOsInfoRequest) (response *DescribeOsInfoResponse, err error) {
    return c.DescribeOsInfoWithContext(context.Background(), request)
}

// DescribeOsInfo
// 查询指定机型所支持的操作系统
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeOsInfoWithContext(ctx context.Context, request *DescribeOsInfoRequest) (response *DescribeOsInfoResponse, err error) {
    if request == nil {
        request = NewDescribeOsInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOsInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOsInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePsaRegulationsRequest() (request *DescribePsaRegulationsRequest) {
    request = &DescribePsaRegulationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "DescribePsaRegulations")
    
    
    return
}

func NewDescribePsaRegulationsResponse() (response *DescribePsaRegulationsResponse) {
    response = &DescribePsaRegulationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePsaRegulations
// 获取预授权规则列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribePsaRegulations(request *DescribePsaRegulationsRequest) (response *DescribePsaRegulationsResponse, err error) {
    return c.DescribePsaRegulationsWithContext(context.Background(), request)
}

// DescribePsaRegulations
// 获取预授权规则列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribePsaRegulationsWithContext(ctx context.Context, request *DescribePsaRegulationsRequest) (response *DescribePsaRegulationsResponse, err error) {
    if request == nil {
        request = NewDescribePsaRegulationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePsaRegulations require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePsaRegulationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegionsRequest() (request *DescribeRegionsRequest) {
    request = &DescribeRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "DescribeRegions")
    
    
    return
}

func NewDescribeRegionsResponse() (response *DescribeRegionsResponse) {
    response = &DescribeRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRegions
// 查询地域以及可用区
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeRegions(request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
    return c.DescribeRegionsWithContext(context.Background(), request)
}

// DescribeRegions
// 查询地域以及可用区
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeRegionsWithContext(ctx context.Context, request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeRegionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRegions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRepairTaskConstantRequest() (request *DescribeRepairTaskConstantRequest) {
    request = &DescribeRepairTaskConstantRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "DescribeRepairTaskConstant")
    
    
    return
}

func NewDescribeRepairTaskConstantResponse() (response *DescribeRepairTaskConstantResponse) {
    response = &DescribeRepairTaskConstantResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRepairTaskConstant
// 维修任务配置获取
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeRepairTaskConstant(request *DescribeRepairTaskConstantRequest) (response *DescribeRepairTaskConstantResponse, err error) {
    return c.DescribeRepairTaskConstantWithContext(context.Background(), request)
}

// DescribeRepairTaskConstant
// 维修任务配置获取
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeRepairTaskConstantWithContext(ctx context.Context, request *DescribeRepairTaskConstantRequest) (response *DescribeRepairTaskConstantResponse, err error) {
    if request == nil {
        request = NewDescribeRepairTaskConstantRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRepairTaskConstant require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRepairTaskConstantResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskInfoRequest() (request *DescribeTaskInfoRequest) {
    request = &DescribeTaskInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "DescribeTaskInfo")
    
    
    return
}

func NewDescribeTaskInfoResponse() (response *DescribeTaskInfoResponse) {
    response = &DescribeTaskInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTaskInfo
// 获取用户维修任务列表及详细信息<br>
//
// <br>
//
// TaskStatus（任务状态ID）与状态中文名的对应关系如下：<br>
//
// 1：未授权<br>
//
// 2：处理中<br>
//
// 3：待确认<br>
//
// 4：未授权-暂不处理<br>
//
// 5：已恢复<br>
//
// 6：待确认-未恢复<br>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTaskInfo(request *DescribeTaskInfoRequest) (response *DescribeTaskInfoResponse, err error) {
    return c.DescribeTaskInfoWithContext(context.Background(), request)
}

// DescribeTaskInfo
// 获取用户维修任务列表及详细信息<br>
//
// <br>
//
// TaskStatus（任务状态ID）与状态中文名的对应关系如下：<br>
//
// 1：未授权<br>
//
// 2：处理中<br>
//
// 3：待确认<br>
//
// 4：未授权-暂不处理<br>
//
// 5：已恢复<br>
//
// 6：待确认-未恢复<br>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTaskInfoWithContext(ctx context.Context, request *DescribeTaskInfoRequest) (response *DescribeTaskInfoResponse, err error) {
    if request == nil {
        request = NewDescribeTaskInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskOperationLogRequest() (request *DescribeTaskOperationLogRequest) {
    request = &DescribeTaskOperationLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "DescribeTaskOperationLog")
    
    
    return
}

func NewDescribeTaskOperationLogResponse() (response *DescribeTaskOperationLogResponse) {
    response = &DescribeTaskOperationLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTaskOperationLog
// 获取维修任务操作日志
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTaskOperationLog(request *DescribeTaskOperationLogRequest) (response *DescribeTaskOperationLogResponse, err error) {
    return c.DescribeTaskOperationLogWithContext(context.Background(), request)
}

// DescribeTaskOperationLog
// 获取维修任务操作日志
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTaskOperationLogWithContext(ctx context.Context, request *DescribeTaskOperationLogRequest) (response *DescribeTaskOperationLogResponse, err error) {
    if request == nil {
        request = NewDescribeTaskOperationLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskOperationLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskOperationLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserCmdTaskInfoRequest() (request *DescribeUserCmdTaskInfoRequest) {
    request = &DescribeUserCmdTaskInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "DescribeUserCmdTaskInfo")
    
    
    return
}

func NewDescribeUserCmdTaskInfoResponse() (response *DescribeUserCmdTaskInfoResponse) {
    response = &DescribeUserCmdTaskInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUserCmdTaskInfo
// 获取自定义脚本任务详细信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeUserCmdTaskInfo(request *DescribeUserCmdTaskInfoRequest) (response *DescribeUserCmdTaskInfoResponse, err error) {
    return c.DescribeUserCmdTaskInfoWithContext(context.Background(), request)
}

// DescribeUserCmdTaskInfo
// 获取自定义脚本任务详细信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeUserCmdTaskInfoWithContext(ctx context.Context, request *DescribeUserCmdTaskInfoRequest) (response *DescribeUserCmdTaskInfoResponse, err error) {
    if request == nil {
        request = NewDescribeUserCmdTaskInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserCmdTaskInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserCmdTaskInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserCmdTasksRequest() (request *DescribeUserCmdTasksRequest) {
    request = &DescribeUserCmdTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "DescribeUserCmdTasks")
    
    
    return
}

func NewDescribeUserCmdTasksResponse() (response *DescribeUserCmdTasksResponse) {
    response = &DescribeUserCmdTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUserCmdTasks
// 获取自定义脚本任务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeUserCmdTasks(request *DescribeUserCmdTasksRequest) (response *DescribeUserCmdTasksResponse, err error) {
    return c.DescribeUserCmdTasksWithContext(context.Background(), request)
}

// DescribeUserCmdTasks
// 获取自定义脚本任务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeUserCmdTasksWithContext(ctx context.Context, request *DescribeUserCmdTasksRequest) (response *DescribeUserCmdTasksResponse, err error) {
    if request == nil {
        request = NewDescribeUserCmdTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserCmdTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserCmdTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserCmdsRequest() (request *DescribeUserCmdsRequest) {
    request = &DescribeUserCmdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "DescribeUserCmds")
    
    
    return
}

func NewDescribeUserCmdsResponse() (response *DescribeUserCmdsResponse) {
    response = &DescribeUserCmdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUserCmds
// 获取自定义脚本信息列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUserCmds(request *DescribeUserCmdsRequest) (response *DescribeUserCmdsResponse, err error) {
    return c.DescribeUserCmdsWithContext(context.Background(), request)
}

// DescribeUserCmds
// 获取自定义脚本信息列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUserCmdsWithContext(ctx context.Context, request *DescribeUserCmdsRequest) (response *DescribeUserCmdsResponse, err error) {
    if request == nil {
        request = NewDescribeUserCmdsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserCmds require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserCmdsResponse()
    err = c.Send(request, response)
    return
}

func NewDetachCamRoleRequest() (request *DetachCamRoleRequest) {
    request = &DetachCamRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "DetachCamRole")
    
    
    return
}

func NewDetachCamRoleResponse() (response *DetachCamRoleResponse) {
    response = &DetachCamRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetachCamRole
// 服务器绑定CAM角色
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DetachCamRole(request *DetachCamRoleRequest) (response *DetachCamRoleResponse, err error) {
    return c.DetachCamRoleWithContext(context.Background(), request)
}

// DetachCamRole
// 服务器绑定CAM角色
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DetachCamRoleWithContext(ctx context.Context, request *DetachCamRoleRequest) (response *DetachCamRoleResponse, err error) {
    if request == nil {
        request = NewDetachCamRoleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetachCamRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetachCamRoleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCustomImageAttributeRequest() (request *ModifyCustomImageAttributeRequest) {
    request = &ModifyCustomImageAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "ModifyCustomImageAttribute")
    
    
    return
}

func NewModifyCustomImageAttributeResponse() (response *ModifyCustomImageAttributeResponse) {
    response = &ModifyCustomImageAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCustomImageAttribute
// 用于修改自定义镜像名或描述
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyCustomImageAttribute(request *ModifyCustomImageAttributeRequest) (response *ModifyCustomImageAttributeResponse, err error) {
    return c.ModifyCustomImageAttributeWithContext(context.Background(), request)
}

// ModifyCustomImageAttribute
// 用于修改自定义镜像名或描述
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyCustomImageAttributeWithContext(ctx context.Context, request *ModifyCustomImageAttributeRequest) (response *ModifyCustomImageAttributeResponse, err error) {
    if request == nil {
        request = NewModifyCustomImageAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCustomImageAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCustomImageAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDeviceAliasesRequest() (request *ModifyDeviceAliasesRequest) {
    request = &ModifyDeviceAliasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "ModifyDeviceAliases")
    
    
    return
}

func NewModifyDeviceAliasesResponse() (response *ModifyDeviceAliasesResponse) {
    response = &ModifyDeviceAliasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDeviceAliases
// 修改服务器名称
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyDeviceAliases(request *ModifyDeviceAliasesRequest) (response *ModifyDeviceAliasesResponse, err error) {
    return c.ModifyDeviceAliasesWithContext(context.Background(), request)
}

// ModifyDeviceAliases
// 修改服务器名称
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyDeviceAliasesWithContext(ctx context.Context, request *ModifyDeviceAliasesRequest) (response *ModifyDeviceAliasesResponse, err error) {
    if request == nil {
        request = NewModifyDeviceAliasesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDeviceAliases require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDeviceAliasesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDeviceAutoRenewFlagRequest() (request *ModifyDeviceAutoRenewFlagRequest) {
    request = &ModifyDeviceAutoRenewFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "ModifyDeviceAutoRenewFlag")
    
    
    return
}

func NewModifyDeviceAutoRenewFlagResponse() (response *ModifyDeviceAutoRenewFlagResponse) {
    response = &ModifyDeviceAutoRenewFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDeviceAutoRenewFlag
// 修改物理机服务器自动续费标志
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyDeviceAutoRenewFlag(request *ModifyDeviceAutoRenewFlagRequest) (response *ModifyDeviceAutoRenewFlagResponse, err error) {
    return c.ModifyDeviceAutoRenewFlagWithContext(context.Background(), request)
}

// ModifyDeviceAutoRenewFlag
// 修改物理机服务器自动续费标志
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyDeviceAutoRenewFlagWithContext(ctx context.Context, request *ModifyDeviceAutoRenewFlagRequest) (response *ModifyDeviceAutoRenewFlagResponse, err error) {
    if request == nil {
        request = NewModifyDeviceAutoRenewFlagRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDeviceAutoRenewFlag require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDeviceAutoRenewFlagResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLanIpRequest() (request *ModifyLanIpRequest) {
    request = &ModifyLanIpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "ModifyLanIp")
    
    
    return
}

func NewModifyLanIpResponse() (response *ModifyLanIpResponse) {
    response = &ModifyLanIpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLanIp
// 修改物理机内网IP（不重装系统）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TSCAGENTOFFLINE = "FailedOperation.TscAgentOffline"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDOPERATION = "UnsupportedOperation.InvalidOperation"
func (c *Client) ModifyLanIp(request *ModifyLanIpRequest) (response *ModifyLanIpResponse, err error) {
    return c.ModifyLanIpWithContext(context.Background(), request)
}

// ModifyLanIp
// 修改物理机内网IP（不重装系统）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TSCAGENTOFFLINE = "FailedOperation.TscAgentOffline"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INVALIDOPERATION = "UnsupportedOperation.InvalidOperation"
func (c *Client) ModifyLanIpWithContext(ctx context.Context, request *ModifyLanIpRequest) (response *ModifyLanIpResponse, err error) {
    if request == nil {
        request = NewModifyLanIpRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLanIp require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLanIpResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPayModePre2PostRequest() (request *ModifyPayModePre2PostRequest) {
    request = &ModifyPayModePre2PostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "ModifyPayModePre2Post")
    
    
    return
}

func NewModifyPayModePre2PostResponse() (response *ModifyPayModePre2PostResponse) {
    response = &ModifyPayModePre2PostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyPayModePre2Post
// 将设备的预付费模式修改为后付费计费模式，支持批量转换。（前提是客户要加入黑石物理机后付费计费的白名单，申请黑石物理机后付费可以联系腾讯云客服）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyPayModePre2Post(request *ModifyPayModePre2PostRequest) (response *ModifyPayModePre2PostResponse, err error) {
    return c.ModifyPayModePre2PostWithContext(context.Background(), request)
}

// ModifyPayModePre2Post
// 将设备的预付费模式修改为后付费计费模式，支持批量转换。（前提是客户要加入黑石物理机后付费计费的白名单，申请黑石物理机后付费可以联系腾讯云客服）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyPayModePre2PostWithContext(ctx context.Context, request *ModifyPayModePre2PostRequest) (response *ModifyPayModePre2PostResponse, err error) {
    if request == nil {
        request = NewModifyPayModePre2PostRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPayModePre2Post require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPayModePre2PostResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPsaRegulationRequest() (request *ModifyPsaRegulationRequest) {
    request = &ModifyPsaRegulationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "ModifyPsaRegulation")
    
    
    return
}

func NewModifyPsaRegulationResponse() (response *ModifyPsaRegulationResponse) {
    response = &ModifyPsaRegulationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyPsaRegulation
// 允许修改规则信息及关联故障类型
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyPsaRegulation(request *ModifyPsaRegulationRequest) (response *ModifyPsaRegulationResponse, err error) {
    return c.ModifyPsaRegulationWithContext(context.Background(), request)
}

// ModifyPsaRegulation
// 允许修改规则信息及关联故障类型
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyPsaRegulationWithContext(ctx context.Context, request *ModifyPsaRegulationRequest) (response *ModifyPsaRegulationResponse, err error) {
    if request == nil {
        request = NewModifyPsaRegulationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPsaRegulation require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPsaRegulationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUserCmdRequest() (request *ModifyUserCmdRequest) {
    request = &ModifyUserCmdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "ModifyUserCmd")
    
    
    return
}

func NewModifyUserCmdResponse() (response *ModifyUserCmdResponse) {
    response = &ModifyUserCmdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyUserCmd
// 修改自定义脚本
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyUserCmd(request *ModifyUserCmdRequest) (response *ModifyUserCmdResponse, err error) {
    return c.ModifyUserCmdWithContext(context.Background(), request)
}

// ModifyUserCmd
// 修改自定义脚本
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyUserCmdWithContext(ctx context.Context, request *ModifyUserCmdRequest) (response *ModifyUserCmdResponse, err error) {
    if request == nil {
        request = NewModifyUserCmdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUserCmd require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUserCmdResponse()
    err = c.Send(request, response)
    return
}

func NewOfflineDevicesRequest() (request *OfflineDevicesRequest) {
    request = &OfflineDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "OfflineDevices")
    
    
    return
}

func NewOfflineDevicesResponse() (response *OfflineDevicesResponse) {
    response = &OfflineDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// OfflineDevices
// 销毁黑石物理机实例：可以销毁物理机列表中的竞价实例，或回收站列表中所有计费模式的实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) OfflineDevices(request *OfflineDevicesRequest) (response *OfflineDevicesResponse, err error) {
    return c.OfflineDevicesWithContext(context.Background(), request)
}

// OfflineDevices
// 销毁黑石物理机实例：可以销毁物理机列表中的竞价实例，或回收站列表中所有计费模式的实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) OfflineDevicesWithContext(ctx context.Context, request *OfflineDevicesRequest) (response *OfflineDevicesResponse, err error) {
    if request == nil {
        request = NewOfflineDevicesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("OfflineDevices require credential")
    }

    request.SetContext(ctx)
    
    response = NewOfflineDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewRebootDevicesRequest() (request *RebootDevicesRequest) {
    request = &RebootDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "RebootDevices")
    
    
    return
}

func NewRebootDevicesResponse() (response *RebootDevicesResponse) {
    response = &RebootDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RebootDevices
// 重启机器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) RebootDevices(request *RebootDevicesRequest) (response *RebootDevicesResponse, err error) {
    return c.RebootDevicesWithContext(context.Background(), request)
}

// RebootDevices
// 重启机器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) RebootDevicesWithContext(ctx context.Context, request *RebootDevicesRequest) (response *RebootDevicesResponse, err error) {
    if request == nil {
        request = NewRebootDevicesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RebootDevices require credential")
    }

    request.SetContext(ctx)
    
    response = NewRebootDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewRecoverDevicesRequest() (request *RecoverDevicesRequest) {
    request = &RecoverDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "RecoverDevices")
    
    
    return
}

func NewRecoverDevicesResponse() (response *RecoverDevicesResponse) {
    response = &RecoverDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RecoverDevices
// 恢复回收站中的物理机（仅限后付费的物理机）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) RecoverDevices(request *RecoverDevicesRequest) (response *RecoverDevicesResponse, err error) {
    return c.RecoverDevicesWithContext(context.Background(), request)
}

// RecoverDevices
// 恢复回收站中的物理机（仅限后付费的物理机）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) RecoverDevicesWithContext(ctx context.Context, request *RecoverDevicesRequest) (response *RecoverDevicesResponse, err error) {
    if request == nil {
        request = NewRecoverDevicesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecoverDevices require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecoverDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewReloadDeviceOsRequest() (request *ReloadDeviceOsRequest) {
    request = &ReloadDeviceOsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "ReloadDeviceOs")
    
    
    return
}

func NewReloadDeviceOsResponse() (response *ReloadDeviceOsResponse) {
    response = &ReloadDeviceOsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReloadDeviceOs
// 重装操作系统
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXISTREPAIRTASK = "FailedOperation.ExistRepairTask"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ReloadDeviceOs(request *ReloadDeviceOsRequest) (response *ReloadDeviceOsResponse, err error) {
    return c.ReloadDeviceOsWithContext(context.Background(), request)
}

// ReloadDeviceOs
// 重装操作系统
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXISTREPAIRTASK = "FailedOperation.ExistRepairTask"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ReloadDeviceOsWithContext(ctx context.Context, request *ReloadDeviceOsRequest) (response *ReloadDeviceOsResponse, err error) {
    if request == nil {
        request = NewReloadDeviceOsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReloadDeviceOs require credential")
    }

    request.SetContext(ctx)
    
    response = NewReloadDeviceOsResponse()
    err = c.Send(request, response)
    return
}

func NewRepairTaskControlRequest() (request *RepairTaskControlRequest) {
    request = &RepairTaskControlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "RepairTaskControl")
    
    
    return
}

func NewRepairTaskControlResponse() (response *RepairTaskControlResponse) {
    response = &RepairTaskControlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RepairTaskControl
// 此接口用于操作维修任务<br>
//
// 入参TaskId为维修任务ID<br>
//
// 入参Operate表示对维修任务的操作，支持如下取值：<br>
//
// AuthorizeRepair（授权维修）<br>
//
// Ignore（暂不提醒）<br>
//
// ConfirmRecovered（维修完成后，确认故障恢复）<br>
//
// ConfirmUnRecovered（维修完成后，确认故障未恢复，该操作已不推荐用）<br>
//
// NeedRepairAgain（维修完成后，故障未恢复，需要重新维修，推荐用此操作打回）<br>
//
// 入参OperateRemark仅在Operate为NeedRepairAgain时有效，表示打回重修原因，建议给出打回的具体原因。<br>
//
// <br>
//
// 操作约束（当前任务状态(TaskStatus)->对应可执行的操作）：<br>
//
// 未授权(1)->授权维修；暂不处理<br>
//
// 暂不处理(4)->授权维修<br>
//
// 待确认(3)->确认故障恢复；确认故障未恢复；需要重新维修<br>
//
// 未恢复(6)->确认故障恢复<br>
//
// <br>
//
// 对于Ping不可达故障的任务，还允许：<br>
//
// 未授权->确认故障恢复<br>
//
// 暂不处理->确认故障恢复<br>
//
// <br>
//
// 处理中与已恢复状态的任务不允许进行操作。<br>
//
// <br>
//
// 详细信息请访问：https://cloud.tencent.com/document/product/386/18190
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) RepairTaskControl(request *RepairTaskControlRequest) (response *RepairTaskControlResponse, err error) {
    return c.RepairTaskControlWithContext(context.Background(), request)
}

// RepairTaskControl
// 此接口用于操作维修任务<br>
//
// 入参TaskId为维修任务ID<br>
//
// 入参Operate表示对维修任务的操作，支持如下取值：<br>
//
// AuthorizeRepair（授权维修）<br>
//
// Ignore（暂不提醒）<br>
//
// ConfirmRecovered（维修完成后，确认故障恢复）<br>
//
// ConfirmUnRecovered（维修完成后，确认故障未恢复，该操作已不推荐用）<br>
//
// NeedRepairAgain（维修完成后，故障未恢复，需要重新维修，推荐用此操作打回）<br>
//
// 入参OperateRemark仅在Operate为NeedRepairAgain时有效，表示打回重修原因，建议给出打回的具体原因。<br>
//
// <br>
//
// 操作约束（当前任务状态(TaskStatus)->对应可执行的操作）：<br>
//
// 未授权(1)->授权维修；暂不处理<br>
//
// 暂不处理(4)->授权维修<br>
//
// 待确认(3)->确认故障恢复；确认故障未恢复；需要重新维修<br>
//
// 未恢复(6)->确认故障恢复<br>
//
// <br>
//
// 对于Ping不可达故障的任务，还允许：<br>
//
// 未授权->确认故障恢复<br>
//
// 暂不处理->确认故障恢复<br>
//
// <br>
//
// 处理中与已恢复状态的任务不允许进行操作。<br>
//
// <br>
//
// 详细信息请访问：https://cloud.tencent.com/document/product/386/18190
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) RepairTaskControlWithContext(ctx context.Context, request *RepairTaskControlRequest) (response *RepairTaskControlResponse, err error) {
    if request == nil {
        request = NewRepairTaskControlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RepairTaskControl require credential")
    }

    request.SetContext(ctx)
    
    response = NewRepairTaskControlResponse()
    err = c.Send(request, response)
    return
}

func NewResetDevicePasswordRequest() (request *ResetDevicePasswordRequest) {
    request = &ResetDevicePasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "ResetDevicePassword")
    
    
    return
}

func NewResetDevicePasswordResponse() (response *ResetDevicePasswordResponse) {
    response = &ResetDevicePasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResetDevicePassword
// 重置服务器密码
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ResetDevicePassword(request *ResetDevicePasswordRequest) (response *ResetDevicePasswordResponse, err error) {
    return c.ResetDevicePasswordWithContext(context.Background(), request)
}

// ResetDevicePassword
// 重置服务器密码
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ResetDevicePasswordWithContext(ctx context.Context, request *ResetDevicePasswordRequest) (response *ResetDevicePasswordResponse, err error) {
    if request == nil {
        request = NewResetDevicePasswordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetDevicePassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetDevicePasswordResponse()
    err = c.Send(request, response)
    return
}

func NewReturnDevicesRequest() (request *ReturnDevicesRequest) {
    request = &ReturnDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "ReturnDevices")
    
    
    return
}

func NewReturnDevicesResponse() (response *ReturnDevicesResponse) {
    response = &ReturnDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReturnDevices
// 退回物理机至回收站，支持批量退还不同计费模式的物理机（包括预付费、后付费、预付费转后付费）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ReturnDevices(request *ReturnDevicesRequest) (response *ReturnDevicesResponse, err error) {
    return c.ReturnDevicesWithContext(context.Background(), request)
}

// ReturnDevices
// 退回物理机至回收站，支持批量退还不同计费模式的物理机（包括预付费、后付费、预付费转后付费）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ReturnDevicesWithContext(ctx context.Context, request *ReturnDevicesRequest) (response *ReturnDevicesResponse, err error) {
    if request == nil {
        request = NewReturnDevicesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReturnDevices require credential")
    }

    request.SetContext(ctx)
    
    response = NewReturnDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewRunUserCmdRequest() (request *RunUserCmdRequest) {
    request = &RunUserCmdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "RunUserCmd")
    
    
    return
}

func NewRunUserCmdResponse() (response *RunUserCmdResponse) {
    response = &RunUserCmdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RunUserCmd
// 运行自定义脚本
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RunUserCmd(request *RunUserCmdRequest) (response *RunUserCmdResponse, err error) {
    return c.RunUserCmdWithContext(context.Background(), request)
}

// RunUserCmd
// 运行自定义脚本
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RunUserCmdWithContext(ctx context.Context, request *RunUserCmdRequest) (response *RunUserCmdResponse, err error) {
    if request == nil {
        request = NewRunUserCmdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunUserCmd require credential")
    }

    request.SetContext(ctx)
    
    response = NewRunUserCmdResponse()
    err = c.Send(request, response)
    return
}

func NewSetOutBandVpnAuthPasswordRequest() (request *SetOutBandVpnAuthPasswordRequest) {
    request = &SetOutBandVpnAuthPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "SetOutBandVpnAuthPassword")
    
    
    return
}

func NewSetOutBandVpnAuthPasswordResponse() (response *SetOutBandVpnAuthPasswordResponse) {
    response = &SetOutBandVpnAuthPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetOutBandVpnAuthPassword
// 设置带外VPN认证用户密码
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SetOutBandVpnAuthPassword(request *SetOutBandVpnAuthPasswordRequest) (response *SetOutBandVpnAuthPasswordResponse, err error) {
    return c.SetOutBandVpnAuthPasswordWithContext(context.Background(), request)
}

// SetOutBandVpnAuthPassword
// 设置带外VPN认证用户密码
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SetOutBandVpnAuthPasswordWithContext(ctx context.Context, request *SetOutBandVpnAuthPasswordRequest) (response *SetOutBandVpnAuthPasswordResponse, err error) {
    if request == nil {
        request = NewSetOutBandVpnAuthPasswordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetOutBandVpnAuthPassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetOutBandVpnAuthPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewShutdownDevicesRequest() (request *ShutdownDevicesRequest) {
    request = &ShutdownDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "ShutdownDevices")
    
    
    return
}

func NewShutdownDevicesResponse() (response *ShutdownDevicesResponse) {
    response = &ShutdownDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ShutdownDevices
// 关闭服务器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ShutdownDevices(request *ShutdownDevicesRequest) (response *ShutdownDevicesResponse, err error) {
    return c.ShutdownDevicesWithContext(context.Background(), request)
}

// ShutdownDevices
// 关闭服务器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ShutdownDevicesWithContext(ctx context.Context, request *ShutdownDevicesRequest) (response *ShutdownDevicesResponse, err error) {
    if request == nil {
        request = NewShutdownDevicesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ShutdownDevices require credential")
    }

    request.SetContext(ctx)
    
    response = NewShutdownDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewStartDevicesRequest() (request *StartDevicesRequest) {
    request = &StartDevicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "StartDevices")
    
    
    return
}

func NewStartDevicesResponse() (response *StartDevicesResponse) {
    response = &StartDevicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartDevices
// 开启服务器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) StartDevices(request *StartDevicesRequest) (response *StartDevicesResponse, err error) {
    return c.StartDevicesWithContext(context.Background(), request)
}

// StartDevices
// 开启服务器
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) StartDevicesWithContext(ctx context.Context, request *StartDevicesRequest) (response *StartDevicesResponse, err error) {
    if request == nil {
        request = NewStartDevicesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartDevices require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartDevicesResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindPsaTagRequest() (request *UnbindPsaTagRequest) {
    request = &UnbindPsaTagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bm", APIVersion, "UnbindPsaTag")
    
    
    return
}

func NewUnbindPsaTagResponse() (response *UnbindPsaTagResponse) {
    response = &UnbindPsaTagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnbindPsaTag
// 解除标签与预授权规则的绑定
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) UnbindPsaTag(request *UnbindPsaTagRequest) (response *UnbindPsaTagResponse, err error) {
    return c.UnbindPsaTagWithContext(context.Background(), request)
}

// UnbindPsaTag
// 解除标签与预授权规则的绑定
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) UnbindPsaTagWithContext(ctx context.Context, request *UnbindPsaTagRequest) (response *UnbindPsaTagResponse, err error) {
    if request == nil {
        request = NewUnbindPsaTagRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindPsaTag require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindPsaTagResponse()
    err = c.Send(request, response)
    return
}
