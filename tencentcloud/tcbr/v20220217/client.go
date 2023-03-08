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

package v20220217

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2022-02-17"

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


func NewCreateCloudRunEnvRequest() (request *CreateCloudRunEnvRequest) {
    request = &CreateCloudRunEnvRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcbr", APIVersion, "CreateCloudRunEnv")
    
    
    return
}

func NewCreateCloudRunEnvResponse() (response *CreateCloudRunEnvResponse) {
    response = &CreateCloudRunEnvResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCloudRunEnv
// 创建云托管环境，并开通资源。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) CreateCloudRunEnv(request *CreateCloudRunEnvRequest) (response *CreateCloudRunEnvResponse, err error) {
    return c.CreateCloudRunEnvWithContext(context.Background(), request)
}

// CreateCloudRunEnv
// 创建云托管环境，并开通资源。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) CreateCloudRunEnvWithContext(ctx context.Context, request *CreateCloudRunEnvRequest) (response *CreateCloudRunEnvResponse, err error) {
    if request == nil {
        request = NewCreateCloudRunEnvRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudRunEnv require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudRunEnvResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCloudRunServerRequest() (request *CreateCloudRunServerRequest) {
    request = &CreateCloudRunServerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcbr", APIVersion, "CreateCloudRunServer")
    
    
    return
}

func NewCreateCloudRunServerResponse() (response *CreateCloudRunServerResponse) {
    response = &CreateCloudRunServerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCloudRunServer
// 创建云托管服务接口
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateCloudRunServer(request *CreateCloudRunServerRequest) (response *CreateCloudRunServerResponse, err error) {
    return c.CreateCloudRunServerWithContext(context.Background(), request)
}

// CreateCloudRunServer
// 创建云托管服务接口
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateCloudRunServerWithContext(ctx context.Context, request *CreateCloudRunServerRequest) (response *CreateCloudRunServerResponse, err error) {
    if request == nil {
        request = NewCreateCloudRunServerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudRunServer require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudRunServerResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudRunEnvsRequest() (request *DescribeCloudRunEnvsRequest) {
    request = &DescribeCloudRunEnvsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcbr", APIVersion, "DescribeCloudRunEnvs")
    
    
    return
}

func NewDescribeCloudRunEnvsResponse() (response *DescribeCloudRunEnvsResponse) {
    response = &DescribeCloudRunEnvsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCloudRunEnvs
// 获取环境列表，含环境下的各个资源信息。尤其是各资源的唯一标识，是请求各资源的关键参数
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeCloudRunEnvs(request *DescribeCloudRunEnvsRequest) (response *DescribeCloudRunEnvsResponse, err error) {
    return c.DescribeCloudRunEnvsWithContext(context.Background(), request)
}

// DescribeCloudRunEnvs
// 获取环境列表，含环境下的各个资源信息。尤其是各资源的唯一标识，是请求各资源的关键参数
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeCloudRunEnvsWithContext(ctx context.Context, request *DescribeCloudRunEnvsRequest) (response *DescribeCloudRunEnvsResponse, err error) {
    if request == nil {
        request = NewDescribeCloudRunEnvsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudRunEnvs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudRunEnvsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudRunServerDetailRequest() (request *DescribeCloudRunServerDetailRequest) {
    request = &DescribeCloudRunServerDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcbr", APIVersion, "DescribeCloudRunServerDetail")
    
    
    return
}

func NewDescribeCloudRunServerDetailResponse() (response *DescribeCloudRunServerDetailResponse) {
    response = &DescribeCloudRunServerDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCloudRunServerDetail
// 查询云托管服务详情
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCloudRunServerDetail(request *DescribeCloudRunServerDetailRequest) (response *DescribeCloudRunServerDetailResponse, err error) {
    return c.DescribeCloudRunServerDetailWithContext(context.Background(), request)
}

// DescribeCloudRunServerDetail
// 查询云托管服务详情
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCloudRunServerDetailWithContext(ctx context.Context, request *DescribeCloudRunServerDetailRequest) (response *DescribeCloudRunServerDetailResponse, err error) {
    if request == nil {
        request = NewDescribeCloudRunServerDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudRunServerDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudRunServerDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudRunServersRequest() (request *DescribeCloudRunServersRequest) {
    request = &DescribeCloudRunServersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcbr", APIVersion, "DescribeCloudRunServers")
    
    
    return
}

func NewDescribeCloudRunServersResponse() (response *DescribeCloudRunServersResponse) {
    response = &DescribeCloudRunServersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCloudRunServers
// 查询云托管服务列表接口
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCloudRunServers(request *DescribeCloudRunServersRequest) (response *DescribeCloudRunServersResponse, err error) {
    return c.DescribeCloudRunServersWithContext(context.Background(), request)
}

// DescribeCloudRunServers
// 查询云托管服务列表接口
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCloudRunServersWithContext(ctx context.Context, request *DescribeCloudRunServersRequest) (response *DescribeCloudRunServersResponse, err error) {
    if request == nil {
        request = NewDescribeCloudRunServersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudRunServers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudRunServersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnvBaseInfoRequest() (request *DescribeEnvBaseInfoRequest) {
    request = &DescribeEnvBaseInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcbr", APIVersion, "DescribeEnvBaseInfo")
    
    
    return
}

func NewDescribeEnvBaseInfoResponse() (response *DescribeEnvBaseInfoResponse) {
    response = &DescribeEnvBaseInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEnvBaseInfo
// 查询环境基础信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeEnvBaseInfo(request *DescribeEnvBaseInfoRequest) (response *DescribeEnvBaseInfoResponse, err error) {
    return c.DescribeEnvBaseInfoWithContext(context.Background(), request)
}

// DescribeEnvBaseInfo
// 查询环境基础信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeEnvBaseInfoWithContext(ctx context.Context, request *DescribeEnvBaseInfoRequest) (response *DescribeEnvBaseInfoResponse, err error) {
    if request == nil {
        request = NewDescribeEnvBaseInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEnvBaseInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEnvBaseInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServerManageTaskRequest() (request *DescribeServerManageTaskRequest) {
    request = &DescribeServerManageTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcbr", APIVersion, "DescribeServerManageTask")
    
    
    return
}

func NewDescribeServerManageTaskResponse() (response *DescribeServerManageTaskResponse) {
    response = &DescribeServerManageTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeServerManageTask
// 查询服务管理任务信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeServerManageTask(request *DescribeServerManageTaskRequest) (response *DescribeServerManageTaskResponse, err error) {
    return c.DescribeServerManageTaskWithContext(context.Background(), request)
}

// DescribeServerManageTask
// 查询服务管理任务信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeServerManageTaskWithContext(ctx context.Context, request *DescribeServerManageTaskRequest) (response *DescribeServerManageTaskResponse, err error) {
    if request == nil {
        request = NewDescribeServerManageTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServerManageTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServerManageTaskResponse()
    err = c.Send(request, response)
    return
}

func NewOperateServerManageRequest() (request *OperateServerManageRequest) {
    request = &OperateServerManageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcbr", APIVersion, "OperateServerManage")
    
    
    return
}

func NewOperateServerManageResponse() (response *OperateServerManageResponse) {
    response = &OperateServerManageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// OperateServerManage
// 操作发布单
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) OperateServerManage(request *OperateServerManageRequest) (response *OperateServerManageResponse, err error) {
    return c.OperateServerManageWithContext(context.Background(), request)
}

// OperateServerManage
// 操作发布单
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) OperateServerManageWithContext(ctx context.Context, request *OperateServerManageRequest) (response *OperateServerManageResponse, err error) {
    if request == nil {
        request = NewOperateServerManageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("OperateServerManage require credential")
    }

    request.SetContext(ctx)
    
    response = NewOperateServerManageResponse()
    err = c.Send(request, response)
    return
}

func NewReleaseGrayRequest() (request *ReleaseGrayRequest) {
    request = &ReleaseGrayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcbr", APIVersion, "ReleaseGray")
    
    
    return
}

func NewReleaseGrayResponse() (response *ReleaseGrayResponse) {
    response = &ReleaseGrayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReleaseGray
// 灰度发布
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ReleaseGray(request *ReleaseGrayRequest) (response *ReleaseGrayResponse, err error) {
    return c.ReleaseGrayWithContext(context.Background(), request)
}

// ReleaseGray
// 灰度发布
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ReleaseGrayWithContext(ctx context.Context, request *ReleaseGrayRequest) (response *ReleaseGrayResponse, err error) {
    if request == nil {
        request = NewReleaseGrayRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReleaseGray require credential")
    }

    request.SetContext(ctx)
    
    response = NewReleaseGrayResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCloudRunServerRequest() (request *UpdateCloudRunServerRequest) {
    request = &UpdateCloudRunServerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tcbr", APIVersion, "UpdateCloudRunServer")
    
    
    return
}

func NewUpdateCloudRunServerResponse() (response *UpdateCloudRunServerResponse) {
    response = &UpdateCloudRunServerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateCloudRunServer
// 更新云托管服务
//
// 可能返回的错误码:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) UpdateCloudRunServer(request *UpdateCloudRunServerRequest) (response *UpdateCloudRunServerResponse, err error) {
    return c.UpdateCloudRunServerWithContext(context.Background(), request)
}

// UpdateCloudRunServer
// 更新云托管服务
//
// 可能返回的错误码:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) UpdateCloudRunServerWithContext(ctx context.Context, request *UpdateCloudRunServerRequest) (response *UpdateCloudRunServerResponse, err error) {
    if request == nil {
        request = NewUpdateCloudRunServerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateCloudRunServer require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateCloudRunServerResponse()
    err = c.Send(request, response)
    return
}
