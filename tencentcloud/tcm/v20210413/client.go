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

package v20210413

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-04-13"

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


func NewCreateMeshRequest() (request *CreateMeshRequest) {
    request = &CreateMeshRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcm", APIVersion, "CreateMesh")
    
    
    return
}

func NewCreateMeshResponse() (response *CreateMeshResponse) {
    response = &CreateMeshResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateMesh
// 创建网格
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RBACFORBIDDEN = "FailedOperation.RBACForbidden"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateMesh(request *CreateMeshRequest) (response *CreateMeshResponse, err error) {
    return c.CreateMeshWithContext(context.Background(), request)
}

// CreateMesh
// 创建网格
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RBACFORBIDDEN = "FailedOperation.RBACForbidden"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateMeshWithContext(ctx context.Context, request *CreateMeshRequest) (response *CreateMeshResponse, err error) {
    if request == nil {
        request = NewCreateMeshRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMesh require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMeshResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMeshRequest() (request *DeleteMeshRequest) {
    request = &DeleteMeshRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcm", APIVersion, "DeleteMesh")
    
    
    return
}

func NewDeleteMeshResponse() (response *DeleteMeshResponse) {
    response = &DeleteMeshResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteMesh
// 删除网格
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RBACFORBIDDEN = "FailedOperation.RBACForbidden"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteMesh(request *DeleteMeshRequest) (response *DeleteMeshResponse, err error) {
    return c.DeleteMeshWithContext(context.Background(), request)
}

// DeleteMesh
// 删除网格
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RBACFORBIDDEN = "FailedOperation.RBACForbidden"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteMeshWithContext(ctx context.Context, request *DeleteMeshRequest) (response *DeleteMeshResponse, err error) {
    if request == nil {
        request = NewDeleteMeshRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMesh require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMeshResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessLogConfigRequest() (request *DescribeAccessLogConfigRequest) {
    request = &DescribeAccessLogConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcm", APIVersion, "DescribeAccessLogConfig")
    
    
    return
}

func NewDescribeAccessLogConfigResponse() (response *DescribeAccessLogConfigResponse) {
    response = &DescribeAccessLogConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAccessLogConfig
// 获取AccessLog配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RBACFORBIDDEN = "FailedOperation.RBACForbidden"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAccessLogConfig(request *DescribeAccessLogConfigRequest) (response *DescribeAccessLogConfigResponse, err error) {
    return c.DescribeAccessLogConfigWithContext(context.Background(), request)
}

// DescribeAccessLogConfig
// 获取AccessLog配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RBACFORBIDDEN = "FailedOperation.RBACForbidden"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAccessLogConfigWithContext(ctx context.Context, request *DescribeAccessLogConfigRequest) (response *DescribeAccessLogConfigResponse, err error) {
    if request == nil {
        request = NewDescribeAccessLogConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccessLogConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccessLogConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMeshRequest() (request *DescribeMeshRequest) {
    request = &DescribeMeshRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcm", APIVersion, "DescribeMesh")
    
    
    return
}

func NewDescribeMeshResponse() (response *DescribeMeshResponse) {
    response = &DescribeMeshResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMesh
// 查询网格详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeMesh(request *DescribeMeshRequest) (response *DescribeMeshResponse, err error) {
    return c.DescribeMeshWithContext(context.Background(), request)
}

// DescribeMesh
// 查询网格详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeMeshWithContext(ctx context.Context, request *DescribeMeshRequest) (response *DescribeMeshResponse, err error) {
    if request == nil {
        request = NewDescribeMeshRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMesh require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMeshResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMeshListRequest() (request *DescribeMeshListRequest) {
    request = &DescribeMeshListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcm", APIVersion, "DescribeMeshList")
    
    
    return
}

func NewDescribeMeshListResponse() (response *DescribeMeshListResponse) {
    response = &DescribeMeshListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMeshList
// 查询网格列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeMeshList(request *DescribeMeshListRequest) (response *DescribeMeshListResponse, err error) {
    return c.DescribeMeshListWithContext(context.Background(), request)
}

// DescribeMeshList
// 查询网格列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeMeshListWithContext(ctx context.Context, request *DescribeMeshListRequest) (response *DescribeMeshListResponse, err error) {
    if request == nil {
        request = NewDescribeMeshListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMeshList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMeshListResponse()
    err = c.Send(request, response)
    return
}

func NewLinkClusterListRequest() (request *LinkClusterListRequest) {
    request = &LinkClusterListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcm", APIVersion, "LinkClusterList")
    
    
    return
}

func NewLinkClusterListResponse() (response *LinkClusterListResponse) {
    response = &LinkClusterListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// LinkClusterList
// 关联集群
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RBACFORBIDDEN = "FailedOperation.RBACForbidden"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) LinkClusterList(request *LinkClusterListRequest) (response *LinkClusterListResponse, err error) {
    return c.LinkClusterListWithContext(context.Background(), request)
}

// LinkClusterList
// 关联集群
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RBACFORBIDDEN = "FailedOperation.RBACForbidden"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) LinkClusterListWithContext(ctx context.Context, request *LinkClusterListRequest) (response *LinkClusterListResponse, err error) {
    if request == nil {
        request = NewLinkClusterListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("LinkClusterList require credential")
    }

    request.SetContext(ctx)
    
    response = NewLinkClusterListResponse()
    err = c.Send(request, response)
    return
}

func NewLinkPrometheusRequest() (request *LinkPrometheusRequest) {
    request = &LinkPrometheusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcm", APIVersion, "LinkPrometheus")
    
    
    return
}

func NewLinkPrometheusResponse() (response *LinkPrometheusResponse) {
    response = &LinkPrometheusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// LinkPrometheus
// 关联Prometheus
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLUSTERNOENOUGHRESOURCE = "FailedOperation.ClusterNoEnoughResource"
//  FAILEDOPERATION_RBACFORBIDDEN = "FailedOperation.RBACForbidden"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) LinkPrometheus(request *LinkPrometheusRequest) (response *LinkPrometheusResponse, err error) {
    return c.LinkPrometheusWithContext(context.Background(), request)
}

// LinkPrometheus
// 关联Prometheus
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLUSTERNOENOUGHRESOURCE = "FailedOperation.ClusterNoEnoughResource"
//  FAILEDOPERATION_RBACFORBIDDEN = "FailedOperation.RBACForbidden"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) LinkPrometheusWithContext(ctx context.Context, request *LinkPrometheusRequest) (response *LinkPrometheusResponse, err error) {
    if request == nil {
        request = NewLinkPrometheusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("LinkPrometheus require credential")
    }

    request.SetContext(ctx)
    
    response = NewLinkPrometheusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccessLogConfigRequest() (request *ModifyAccessLogConfigRequest) {
    request = &ModifyAccessLogConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcm", APIVersion, "ModifyAccessLogConfig")
    
    
    return
}

func NewModifyAccessLogConfigResponse() (response *ModifyAccessLogConfigResponse) {
    response = &ModifyAccessLogConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAccessLogConfig
// 修改访问日志配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RBACFORBIDDEN = "FailedOperation.RBACForbidden"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAccessLogConfig(request *ModifyAccessLogConfigRequest) (response *ModifyAccessLogConfigResponse, err error) {
    return c.ModifyAccessLogConfigWithContext(context.Background(), request)
}

// ModifyAccessLogConfig
// 修改访问日志配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RBACFORBIDDEN = "FailedOperation.RBACForbidden"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAccessLogConfigWithContext(ctx context.Context, request *ModifyAccessLogConfigRequest) (response *ModifyAccessLogConfigResponse, err error) {
    if request == nil {
        request = NewModifyAccessLogConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccessLogConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccessLogConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMeshRequest() (request *ModifyMeshRequest) {
    request = &ModifyMeshRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcm", APIVersion, "ModifyMesh")
    
    
    return
}

func NewModifyMeshResponse() (response *ModifyMeshResponse) {
    response = &ModifyMeshResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyMesh
// 修改网格
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RBACFORBIDDEN = "FailedOperation.RBACForbidden"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyMesh(request *ModifyMeshRequest) (response *ModifyMeshResponse, err error) {
    return c.ModifyMeshWithContext(context.Background(), request)
}

// ModifyMesh
// 修改网格
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RBACFORBIDDEN = "FailedOperation.RBACForbidden"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyMeshWithContext(ctx context.Context, request *ModifyMeshRequest) (response *ModifyMeshResponse, err error) {
    if request == nil {
        request = NewModifyMeshRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMesh require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMeshResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTracingConfigRequest() (request *ModifyTracingConfigRequest) {
    request = &ModifyTracingConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcm", APIVersion, "ModifyTracingConfig")
    
    
    return
}

func NewModifyTracingConfigResponse() (response *ModifyTracingConfigResponse) {
    response = &ModifyTracingConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyTracingConfig
// 修改 Tracing 配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RBACFORBIDDEN = "FailedOperation.RBACForbidden"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyTracingConfig(request *ModifyTracingConfigRequest) (response *ModifyTracingConfigResponse, err error) {
    return c.ModifyTracingConfigWithContext(context.Background(), request)
}

// ModifyTracingConfig
// 修改 Tracing 配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RBACFORBIDDEN = "FailedOperation.RBACForbidden"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyTracingConfigWithContext(ctx context.Context, request *ModifyTracingConfigRequest) (response *ModifyTracingConfigResponse, err error) {
    if request == nil {
        request = NewModifyTracingConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTracingConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTracingConfigResponse()
    err = c.Send(request, response)
    return
}

func NewUnlinkClusterRequest() (request *UnlinkClusterRequest) {
    request = &UnlinkClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcm", APIVersion, "UnlinkCluster")
    
    
    return
}

func NewUnlinkClusterResponse() (response *UnlinkClusterResponse) {
    response = &UnlinkClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnlinkCluster
// 解关联集群
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RBACFORBIDDEN = "FailedOperation.RBACForbidden"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UnlinkCluster(request *UnlinkClusterRequest) (response *UnlinkClusterResponse, err error) {
    return c.UnlinkClusterWithContext(context.Background(), request)
}

// UnlinkCluster
// 解关联集群
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RBACFORBIDDEN = "FailedOperation.RBACForbidden"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UnlinkClusterWithContext(ctx context.Context, request *UnlinkClusterRequest) (response *UnlinkClusterResponse, err error) {
    if request == nil {
        request = NewUnlinkClusterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnlinkCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnlinkClusterResponse()
    err = c.Send(request, response)
    return
}

func NewUnlinkPrometheusRequest() (request *UnlinkPrometheusRequest) {
    request = &UnlinkPrometheusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcm", APIVersion, "UnlinkPrometheus")
    
    
    return
}

func NewUnlinkPrometheusResponse() (response *UnlinkPrometheusResponse) {
    response = &UnlinkPrometheusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnlinkPrometheus
// 解除关联Prometheus
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UnlinkPrometheus(request *UnlinkPrometheusRequest) (response *UnlinkPrometheusResponse, err error) {
    return c.UnlinkPrometheusWithContext(context.Background(), request)
}

// UnlinkPrometheus
// 解除关联Prometheus
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UnlinkPrometheusWithContext(ctx context.Context, request *UnlinkPrometheusRequest) (response *UnlinkPrometheusResponse, err error) {
    if request == nil {
        request = NewUnlinkPrometheusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnlinkPrometheus require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnlinkPrometheusResponse()
    err = c.Send(request, response)
    return
}
