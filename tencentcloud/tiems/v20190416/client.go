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

package v20190416

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-04-16"

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


func NewCreateJobRequest() (request *CreateJobRequest) {
    request = &CreateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiems", APIVersion, "CreateJob")
    
    
    return
}

func NewCreateJobResponse() (response *CreateJobResponse) {
    response = &CreateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateJob
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 创建任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALREADYEXISTS = "FailedOperation.AlreadyExists"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateJob(request *CreateJobRequest) (response *CreateJobResponse, err error) {
    return c.CreateJobWithContext(context.Background(), request)
}

// CreateJob
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 创建任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALREADYEXISTS = "FailedOperation.AlreadyExists"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateJobWithContext(ctx context.Context, request *CreateJobRequest) (response *CreateJobResponse, err error) {
    if request == nil {
        request = NewCreateJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRsgAsGroupRequest() (request *CreateRsgAsGroupRequest) {
    request = &CreateRsgAsGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiems", APIVersion, "CreateRsgAsGroup")
    
    
    return
}

func NewCreateRsgAsGroupResponse() (response *CreateRsgAsGroupResponse) {
    response = &CreateRsgAsGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRsgAsGroup
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 创建资源组的伸缩组。当前一个资源组仅允许创建一个伸缩组。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALREADYEXISTS = "FailedOperation.AlreadyExists"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateRsgAsGroup(request *CreateRsgAsGroupRequest) (response *CreateRsgAsGroupResponse, err error) {
    return c.CreateRsgAsGroupWithContext(context.Background(), request)
}

// CreateRsgAsGroup
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 创建资源组的伸缩组。当前一个资源组仅允许创建一个伸缩组。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALREADYEXISTS = "FailedOperation.AlreadyExists"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateRsgAsGroupWithContext(ctx context.Context, request *CreateRsgAsGroupRequest) (response *CreateRsgAsGroupResponse, err error) {
    if request == nil {
        request = NewCreateRsgAsGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRsgAsGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRsgAsGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRuntimeRequest() (request *CreateRuntimeRequest) {
    request = &CreateRuntimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiems", APIVersion, "CreateRuntime")
    
    
    return
}

func NewCreateRuntimeResponse() (response *CreateRuntimeResponse) {
    response = &CreateRuntimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRuntime
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 创建运行环境
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALREADYEXISTS = "FailedOperation.AlreadyExists"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateRuntime(request *CreateRuntimeRequest) (response *CreateRuntimeResponse, err error) {
    return c.CreateRuntimeWithContext(context.Background(), request)
}

// CreateRuntime
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 创建运行环境
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALREADYEXISTS = "FailedOperation.AlreadyExists"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateRuntimeWithContext(ctx context.Context, request *CreateRuntimeRequest) (response *CreateRuntimeResponse, err error) {
    if request == nil {
        request = NewCreateRuntimeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRuntime require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRuntimeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateServiceRequest() (request *CreateServiceRequest) {
    request = &CreateServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiems", APIVersion, "CreateService")
    
    
    return
}

func NewCreateServiceResponse() (response *CreateServiceResponse) {
    response = &CreateServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateService
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 创建服务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALREADYEXISTS = "FailedOperation.AlreadyExists"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateService(request *CreateServiceRequest) (response *CreateServiceResponse, err error) {
    return c.CreateServiceWithContext(context.Background(), request)
}

// CreateService
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 创建服务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALREADYEXISTS = "FailedOperation.AlreadyExists"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateServiceWithContext(ctx context.Context, request *CreateServiceRequest) (response *CreateServiceResponse, err error) {
    if request == nil {
        request = NewCreateServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateService require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateServiceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateServiceConfigRequest() (request *CreateServiceConfigRequest) {
    request = &CreateServiceConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiems", APIVersion, "CreateServiceConfig")
    
    
    return
}

func NewCreateServiceConfigResponse() (response *CreateServiceConfigResponse) {
    response = &CreateServiceConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateServiceConfig
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 创建服务配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALREADYEXISTS = "FailedOperation.AlreadyExists"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateServiceConfig(request *CreateServiceConfigRequest) (response *CreateServiceConfigResponse, err error) {
    return c.CreateServiceConfigWithContext(context.Background(), request)
}

// CreateServiceConfig
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 创建服务配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALREADYEXISTS = "FailedOperation.AlreadyExists"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateServiceConfigWithContext(ctx context.Context, request *CreateServiceConfigRequest) (response *CreateServiceConfigResponse, err error) {
    if request == nil {
        request = NewCreateServiceConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateServiceConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateServiceConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteInstanceRequest() (request *DeleteInstanceRequest) {
    request = &DeleteInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiems", APIVersion, "DeleteInstance")
    
    
    return
}

func NewDeleteInstanceResponse() (response *DeleteInstanceResponse) {
    response = &DeleteInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteInstance
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 删除资源组中的节点。目前仅支持删除已经到期的预付费节点，和按量付费节点。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteInstance(request *DeleteInstanceRequest) (response *DeleteInstanceResponse, err error) {
    return c.DeleteInstanceWithContext(context.Background(), request)
}

// DeleteInstance
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 删除资源组中的节点。目前仅支持删除已经到期的预付费节点，和按量付费节点。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteInstanceWithContext(ctx context.Context, request *DeleteInstanceRequest) (response *DeleteInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteJobRequest() (request *DeleteJobRequest) {
    request = &DeleteJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiems", APIVersion, "DeleteJob")
    
    
    return
}

func NewDeleteJobResponse() (response *DeleteJobResponse) {
    response = &DeleteJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteJob
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 删除任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteJob(request *DeleteJobRequest) (response *DeleteJobResponse, err error) {
    return c.DeleteJobWithContext(context.Background(), request)
}

// DeleteJob
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 删除任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteJobWithContext(ctx context.Context, request *DeleteJobRequest) (response *DeleteJobResponse, err error) {
    if request == nil {
        request = NewDeleteJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteJobResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteResourceGroupRequest() (request *DeleteResourceGroupRequest) {
    request = &DeleteResourceGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiems", APIVersion, "DeleteResourceGroup")
    
    
    return
}

func NewDeleteResourceGroupResponse() (response *DeleteResourceGroupResponse) {
    response = &DeleteResourceGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteResourceGroup
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 删除资源组
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteResourceGroup(request *DeleteResourceGroupRequest) (response *DeleteResourceGroupResponse, err error) {
    return c.DeleteResourceGroupWithContext(context.Background(), request)
}

// DeleteResourceGroup
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 删除资源组
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteResourceGroupWithContext(ctx context.Context, request *DeleteResourceGroupRequest) (response *DeleteResourceGroupResponse, err error) {
    if request == nil {
        request = NewDeleteResourceGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteResourceGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteResourceGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRsgAsGroupRequest() (request *DeleteRsgAsGroupRequest) {
    request = &DeleteRsgAsGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiems", APIVersion, "DeleteRsgAsGroup")
    
    
    return
}

func NewDeleteRsgAsGroupResponse() (response *DeleteRsgAsGroupResponse) {
    response = &DeleteRsgAsGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRsgAsGroup
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 伸缩
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteRsgAsGroup(request *DeleteRsgAsGroupRequest) (response *DeleteRsgAsGroupResponse, err error) {
    return c.DeleteRsgAsGroupWithContext(context.Background(), request)
}

// DeleteRsgAsGroup
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 伸缩
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteRsgAsGroupWithContext(ctx context.Context, request *DeleteRsgAsGroupRequest) (response *DeleteRsgAsGroupResponse, err error) {
    if request == nil {
        request = NewDeleteRsgAsGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRsgAsGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRsgAsGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRuntimeRequest() (request *DeleteRuntimeRequest) {
    request = &DeleteRuntimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiems", APIVersion, "DeleteRuntime")
    
    
    return
}

func NewDeleteRuntimeResponse() (response *DeleteRuntimeResponse) {
    response = &DeleteRuntimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRuntime
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 删除运行环境
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRuntime(request *DeleteRuntimeRequest) (response *DeleteRuntimeResponse, err error) {
    return c.DeleteRuntimeWithContext(context.Background(), request)
}

// DeleteRuntime
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 删除运行环境
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRuntimeWithContext(ctx context.Context, request *DeleteRuntimeRequest) (response *DeleteRuntimeResponse, err error) {
    if request == nil {
        request = NewDeleteRuntimeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRuntime require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRuntimeResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteServiceRequest() (request *DeleteServiceRequest) {
    request = &DeleteServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiems", APIVersion, "DeleteService")
    
    
    return
}

func NewDeleteServiceResponse() (response *DeleteServiceResponse) {
    response = &DeleteServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteService
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 删除服务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteService(request *DeleteServiceRequest) (response *DeleteServiceResponse, err error) {
    return c.DeleteServiceWithContext(context.Background(), request)
}

// DeleteService
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 删除服务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteServiceWithContext(ctx context.Context, request *DeleteServiceRequest) (response *DeleteServiceResponse, err error) {
    if request == nil {
        request = NewDeleteServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteService require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteServiceConfigRequest() (request *DeleteServiceConfigRequest) {
    request = &DeleteServiceConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiems", APIVersion, "DeleteServiceConfig")
    
    
    return
}

func NewDeleteServiceConfigResponse() (response *DeleteServiceConfigResponse) {
    response = &DeleteServiceConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteServiceConfig
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 删除服务配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteServiceConfig(request *DeleteServiceConfigRequest) (response *DeleteServiceConfigResponse, err error) {
    return c.DeleteServiceConfigWithContext(context.Background(), request)
}

// DeleteServiceConfig
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 删除服务配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteServiceConfigWithContext(ctx context.Context, request *DeleteServiceConfigRequest) (response *DeleteServiceConfigResponse, err error) {
    if request == nil {
        request = NewDeleteServiceConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteServiceConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteServiceConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiems", APIVersion, "DescribeInstances")
    
    
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstances
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 获取节点列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    return c.DescribeInstancesWithContext(context.Background(), request)
}

// DescribeInstances
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 获取节点列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
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

func NewDescribeResourceGroupsRequest() (request *DescribeResourceGroupsRequest) {
    request = &DescribeResourceGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiems", APIVersion, "DescribeResourceGroups")
    
    
    return
}

func NewDescribeResourceGroupsResponse() (response *DescribeResourceGroupsResponse) {
    response = &DescribeResourceGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeResourceGroups
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 获取资源组列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeResourceGroups(request *DescribeResourceGroupsRequest) (response *DescribeResourceGroupsResponse, err error) {
    return c.DescribeResourceGroupsWithContext(context.Background(), request)
}

// DescribeResourceGroups
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 获取资源组列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeResourceGroupsWithContext(ctx context.Context, request *DescribeResourceGroupsRequest) (response *DescribeResourceGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeResourceGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResourceGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourceGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRsgAsGroupActivitiesRequest() (request *DescribeRsgAsGroupActivitiesRequest) {
    request = &DescribeRsgAsGroupActivitiesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiems", APIVersion, "DescribeRsgAsGroupActivities")
    
    
    return
}

func NewDescribeRsgAsGroupActivitiesResponse() (response *DescribeRsgAsGroupActivitiesResponse) {
    response = &DescribeRsgAsGroupActivitiesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRsgAsGroupActivities
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 查询伸缩组活动
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRsgAsGroupActivities(request *DescribeRsgAsGroupActivitiesRequest) (response *DescribeRsgAsGroupActivitiesResponse, err error) {
    return c.DescribeRsgAsGroupActivitiesWithContext(context.Background(), request)
}

// DescribeRsgAsGroupActivities
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 查询伸缩组活动
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRsgAsGroupActivitiesWithContext(ctx context.Context, request *DescribeRsgAsGroupActivitiesRequest) (response *DescribeRsgAsGroupActivitiesResponse, err error) {
    if request == nil {
        request = NewDescribeRsgAsGroupActivitiesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRsgAsGroupActivities require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRsgAsGroupActivitiesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRsgAsGroupsRequest() (request *DescribeRsgAsGroupsRequest) {
    request = &DescribeRsgAsGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiems", APIVersion, "DescribeRsgAsGroups")
    
    
    return
}

func NewDescribeRsgAsGroupsResponse() (response *DescribeRsgAsGroupsResponse) {
    response = &DescribeRsgAsGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRsgAsGroups
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 查询资源组的伸缩组信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRsgAsGroups(request *DescribeRsgAsGroupsRequest) (response *DescribeRsgAsGroupsResponse, err error) {
    return c.DescribeRsgAsGroupsWithContext(context.Background(), request)
}

// DescribeRsgAsGroups
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 查询资源组的伸缩组信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRsgAsGroupsWithContext(ctx context.Context, request *DescribeRsgAsGroupsRequest) (response *DescribeRsgAsGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeRsgAsGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRsgAsGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRsgAsGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRuntimesRequest() (request *DescribeRuntimesRequest) {
    request = &DescribeRuntimesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiems", APIVersion, "DescribeRuntimes")
    
    
    return
}

func NewDescribeRuntimesResponse() (response *DescribeRuntimesResponse) {
    response = &DescribeRuntimesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRuntimes
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 描述服务运行环境
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeRuntimes(request *DescribeRuntimesRequest) (response *DescribeRuntimesResponse, err error) {
    return c.DescribeRuntimesWithContext(context.Background(), request)
}

// DescribeRuntimes
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 描述服务运行环境
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeRuntimesWithContext(ctx context.Context, request *DescribeRuntimesRequest) (response *DescribeRuntimesResponse, err error) {
    if request == nil {
        request = NewDescribeRuntimesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRuntimes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRuntimesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServiceConfigsRequest() (request *DescribeServiceConfigsRequest) {
    request = &DescribeServiceConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiems", APIVersion, "DescribeServiceConfigs")
    
    
    return
}

func NewDescribeServiceConfigsResponse() (response *DescribeServiceConfigsResponse) {
    response = &DescribeServiceConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeServiceConfigs
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 描述服务配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeServiceConfigs(request *DescribeServiceConfigsRequest) (response *DescribeServiceConfigsResponse, err error) {
    return c.DescribeServiceConfigsWithContext(context.Background(), request)
}

// DescribeServiceConfigs
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 描述服务配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeServiceConfigsWithContext(ctx context.Context, request *DescribeServiceConfigsRequest) (response *DescribeServiceConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeServiceConfigsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServiceConfigs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServiceConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServicesRequest() (request *DescribeServicesRequest) {
    request = &DescribeServicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiems", APIVersion, "DescribeServices")
    
    
    return
}

func NewDescribeServicesResponse() (response *DescribeServicesResponse) {
    response = &DescribeServicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeServices
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 描述服务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALREADYEXISTS = "FailedOperation.AlreadyExists"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeServices(request *DescribeServicesRequest) (response *DescribeServicesResponse, err error) {
    return c.DescribeServicesWithContext(context.Background(), request)
}

// DescribeServices
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 描述服务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALREADYEXISTS = "FailedOperation.AlreadyExists"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeServicesWithContext(ctx context.Context, request *DescribeServicesRequest) (response *DescribeServicesResponse, err error) {
    if request == nil {
        request = NewDescribeServicesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServices require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServicesResponse()
    err = c.Send(request, response)
    return
}

func NewDisableRsgAsGroupRequest() (request *DisableRsgAsGroupRequest) {
    request = &DisableRsgAsGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiems", APIVersion, "DisableRsgAsGroup")
    
    
    return
}

func NewDisableRsgAsGroupResponse() (response *DisableRsgAsGroupResponse) {
    response = &DisableRsgAsGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisableRsgAsGroup
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 停用资源组的伸缩组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DisableRsgAsGroup(request *DisableRsgAsGroupRequest) (response *DisableRsgAsGroupResponse, err error) {
    return c.DisableRsgAsGroupWithContext(context.Background(), request)
}

// DisableRsgAsGroup
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 停用资源组的伸缩组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DisableRsgAsGroupWithContext(ctx context.Context, request *DisableRsgAsGroupRequest) (response *DisableRsgAsGroupResponse, err error) {
    if request == nil {
        request = NewDisableRsgAsGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableRsgAsGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableRsgAsGroupResponse()
    err = c.Send(request, response)
    return
}

func NewEnableRsgAsGroupRequest() (request *EnableRsgAsGroupRequest) {
    request = &EnableRsgAsGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiems", APIVersion, "EnableRsgAsGroup")
    
    
    return
}

func NewEnableRsgAsGroupResponse() (response *EnableRsgAsGroupResponse) {
    response = &EnableRsgAsGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EnableRsgAsGroup
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 启用资源组的伸缩组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) EnableRsgAsGroup(request *EnableRsgAsGroupRequest) (response *EnableRsgAsGroupResponse, err error) {
    return c.EnableRsgAsGroupWithContext(context.Background(), request)
}

// EnableRsgAsGroup
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 启用资源组的伸缩组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) EnableRsgAsGroupWithContext(ctx context.Context, request *EnableRsgAsGroupRequest) (response *EnableRsgAsGroupResponse, err error) {
    if request == nil {
        request = NewEnableRsgAsGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableRsgAsGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableRsgAsGroupResponse()
    err = c.Send(request, response)
    return
}

func NewExposeServiceRequest() (request *ExposeServiceRequest) {
    request = &ExposeServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiems", APIVersion, "ExposeService")
    
    
    return
}

func NewExposeServiceResponse() (response *ExposeServiceResponse) {
    response = &ExposeServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ExposeService
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 暴露服务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALREADYEXISTS = "FailedOperation.AlreadyExists"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ExposeService(request *ExposeServiceRequest) (response *ExposeServiceResponse, err error) {
    return c.ExposeServiceWithContext(context.Background(), request)
}

// ExposeService
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 暴露服务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALREADYEXISTS = "FailedOperation.AlreadyExists"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ExposeServiceWithContext(ctx context.Context, request *ExposeServiceRequest) (response *ExposeServiceResponse, err error) {
    if request == nil {
        request = NewExposeServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExposeService require credential")
    }

    request.SetContext(ctx)
    
    response = NewExposeServiceResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateJobRequest() (request *UpdateJobRequest) {
    request = &UpdateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiems", APIVersion, "UpdateJob")
    
    
    return
}

func NewUpdateJobResponse() (response *UpdateJobResponse) {
    response = &UpdateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateJob
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 更新任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALREADYEXISTS = "FailedOperation.AlreadyExists"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateJob(request *UpdateJobRequest) (response *UpdateJobResponse, err error) {
    return c.UpdateJobWithContext(context.Background(), request)
}

// UpdateJob
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 更新任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ALREADYEXISTS = "FailedOperation.AlreadyExists"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateJobWithContext(ctx context.Context, request *UpdateJobRequest) (response *UpdateJobResponse, err error) {
    if request == nil {
        request = NewUpdateJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateJobResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateRsgAsGroupRequest() (request *UpdateRsgAsGroupRequest) {
    request = &UpdateRsgAsGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiems", APIVersion, "UpdateRsgAsGroup")
    
    
    return
}

func NewUpdateRsgAsGroupResponse() (response *UpdateRsgAsGroupResponse) {
    response = &UpdateRsgAsGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateRsgAsGroup
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 更新资源组的伸缩组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateRsgAsGroup(request *UpdateRsgAsGroupRequest) (response *UpdateRsgAsGroupResponse, err error) {
    return c.UpdateRsgAsGroupWithContext(context.Background(), request)
}

// UpdateRsgAsGroup
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 更新资源组的伸缩组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateRsgAsGroupWithContext(ctx context.Context, request *UpdateRsgAsGroupRequest) (response *UpdateRsgAsGroupResponse, err error) {
    if request == nil {
        request = NewUpdateRsgAsGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateRsgAsGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateRsgAsGroupResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateServiceRequest() (request *UpdateServiceRequest) {
    request = &UpdateServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiems", APIVersion, "UpdateService")
    
    
    return
}

func NewUpdateServiceResponse() (response *UpdateServiceResponse) {
    response = &UpdateServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateService
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 更新服务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) UpdateService(request *UpdateServiceRequest) (response *UpdateServiceResponse, err error) {
    return c.UpdateServiceWithContext(context.Background(), request)
}

// UpdateService
// 因业务策略调整，腾讯云TI平台TI-EMS已经于2022年6月30日下线并停止提供服务。若您有新增的业务需求，可前往TI-ONE(https://cloud.tencent.com/document/product/851)使用。
//
// 
//
// 更新服务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) UpdateServiceWithContext(ctx context.Context, request *UpdateServiceRequest) (response *UpdateServiceResponse, err error) {
    if request == nil {
        request = NewUpdateServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateService require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateServiceResponse()
    err = c.Send(request, response)
    return
}
