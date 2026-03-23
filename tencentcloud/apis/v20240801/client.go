// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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

package v20240801

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2024-08-01"

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


func NewCreateAgentAppRequest() (request *CreateAgentAppRequest) {
    request = &CreateAgentAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apis", APIVersion, "CreateAgentApp")
    
    
    return
}

func NewCreateAgentAppResponse() (response *CreateAgentAppResponse) {
    response = &CreateAgentAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAgentApp
// 创建app
func (c *Client) CreateAgentApp(request *CreateAgentAppRequest) (response *CreateAgentAppResponse, err error) {
    return c.CreateAgentAppWithContext(context.Background(), request)
}

// CreateAgentApp
// 创建app
func (c *Client) CreateAgentAppWithContext(ctx context.Context, request *CreateAgentAppRequest) (response *CreateAgentAppResponse, err error) {
    if request == nil {
        request = NewCreateAgentAppRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apis", APIVersion, "CreateAgentApp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAgentApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAgentAppResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAgentAppMcpServersRequest() (request *CreateAgentAppMcpServersRequest) {
    request = &CreateAgentAppMcpServersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apis", APIVersion, "CreateAgentAppMcpServers")
    
    
    return
}

func NewCreateAgentAppMcpServersResponse() (response *CreateAgentAppMcpServersResponse) {
    response = &CreateAgentAppMcpServersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAgentAppMcpServers
// 创建app的mcp server关联
func (c *Client) CreateAgentAppMcpServers(request *CreateAgentAppMcpServersRequest) (response *CreateAgentAppMcpServersResponse, err error) {
    return c.CreateAgentAppMcpServersWithContext(context.Background(), request)
}

// CreateAgentAppMcpServers
// 创建app的mcp server关联
func (c *Client) CreateAgentAppMcpServersWithContext(ctx context.Context, request *CreateAgentAppMcpServersRequest) (response *CreateAgentAppMcpServersResponse, err error) {
    if request == nil {
        request = NewCreateAgentAppMcpServersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apis", APIVersion, "CreateAgentAppMcpServers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAgentAppMcpServers require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAgentAppMcpServersResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAgentAppModelServicesRequest() (request *CreateAgentAppModelServicesRequest) {
    request = &CreateAgentAppModelServicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apis", APIVersion, "CreateAgentAppModelServices")
    
    
    return
}

func NewCreateAgentAppModelServicesResponse() (response *CreateAgentAppModelServicesResponse) {
    response = &CreateAgentAppModelServicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAgentAppModelServices
// 创建app的model service关联
func (c *Client) CreateAgentAppModelServices(request *CreateAgentAppModelServicesRequest) (response *CreateAgentAppModelServicesResponse, err error) {
    return c.CreateAgentAppModelServicesWithContext(context.Background(), request)
}

// CreateAgentAppModelServices
// 创建app的model service关联
func (c *Client) CreateAgentAppModelServicesWithContext(ctx context.Context, request *CreateAgentAppModelServicesRequest) (response *CreateAgentAppModelServicesResponse, err error) {
    if request == nil {
        request = NewCreateAgentAppModelServicesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apis", APIVersion, "CreateAgentAppModelServices")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAgentAppModelServices require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAgentAppModelServicesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAgentCredentialRequest() (request *CreateAgentCredentialRequest) {
    request = &CreateAgentCredentialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apis", APIVersion, "CreateAgentCredential")
    
    
    return
}

func NewCreateAgentCredentialResponse() (response *CreateAgentCredentialResponse) {
    response = &CreateAgentCredentialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAgentCredential
// 创建Credential
func (c *Client) CreateAgentCredential(request *CreateAgentCredentialRequest) (response *CreateAgentCredentialResponse, err error) {
    return c.CreateAgentCredentialWithContext(context.Background(), request)
}

// CreateAgentCredential
// 创建Credential
func (c *Client) CreateAgentCredentialWithContext(ctx context.Context, request *CreateAgentCredentialRequest) (response *CreateAgentCredentialResponse, err error) {
    if request == nil {
        request = NewCreateAgentCredentialRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apis", APIVersion, "CreateAgentCredential")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAgentCredential require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAgentCredentialResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMcpServerRequest() (request *CreateMcpServerRequest) {
    request = &CreateMcpServerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apis", APIVersion, "CreateMcpServer")
    
    
    return
}

func NewCreateMcpServerResponse() (response *CreateMcpServerResponse) {
    response = &CreateMcpServerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateMcpServer
// 创建mcp server
func (c *Client) CreateMcpServer(request *CreateMcpServerRequest) (response *CreateMcpServerResponse, err error) {
    return c.CreateMcpServerWithContext(context.Background(), request)
}

// CreateMcpServer
// 创建mcp server
func (c *Client) CreateMcpServerWithContext(ctx context.Context, request *CreateMcpServerRequest) (response *CreateMcpServerResponse, err error) {
    if request == nil {
        request = NewCreateMcpServerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apis", APIVersion, "CreateMcpServer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMcpServer require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMcpServerResponse()
    err = c.Send(request, response)
    return
}

func NewCreateModelRequest() (request *CreateModelRequest) {
    request = &CreateModelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apis", APIVersion, "CreateModel")
    
    
    return
}

func NewCreateModelResponse() (response *CreateModelResponse) {
    response = &CreateModelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateModel
// 创建模型
func (c *Client) CreateModel(request *CreateModelRequest) (response *CreateModelResponse, err error) {
    return c.CreateModelWithContext(context.Background(), request)
}

// CreateModel
// 创建模型
func (c *Client) CreateModelWithContext(ctx context.Context, request *CreateModelRequest) (response *CreateModelResponse, err error) {
    if request == nil {
        request = NewCreateModelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apis", APIVersion, "CreateModel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateModel require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateModelResponse()
    err = c.Send(request, response)
    return
}

func NewCreateModelServiceRequest() (request *CreateModelServiceRequest) {
    request = &CreateModelServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apis", APIVersion, "CreateModelService")
    
    
    return
}

func NewCreateModelServiceResponse() (response *CreateModelServiceResponse) {
    response = &CreateModelServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateModelService
// 创建模型服务
func (c *Client) CreateModelService(request *CreateModelServiceRequest) (response *CreateModelServiceResponse, err error) {
    return c.CreateModelServiceWithContext(context.Background(), request)
}

// CreateModelService
// 创建模型服务
func (c *Client) CreateModelServiceWithContext(ctx context.Context, request *CreateModelServiceRequest) (response *CreateModelServiceResponse, err error) {
    if request == nil {
        request = NewCreateModelServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apis", APIVersion, "CreateModelService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateModelService require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateModelServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAgentAppRequest() (request *DeleteAgentAppRequest) {
    request = &DeleteAgentAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apis", APIVersion, "DeleteAgentApp")
    
    
    return
}

func NewDeleteAgentAppResponse() (response *DeleteAgentAppResponse) {
    response = &DeleteAgentAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAgentApp
// 删除app
func (c *Client) DeleteAgentApp(request *DeleteAgentAppRequest) (response *DeleteAgentAppResponse, err error) {
    return c.DeleteAgentAppWithContext(context.Background(), request)
}

// DeleteAgentApp
// 删除app
func (c *Client) DeleteAgentAppWithContext(ctx context.Context, request *DeleteAgentAppRequest) (response *DeleteAgentAppResponse, err error) {
    if request == nil {
        request = NewDeleteAgentAppRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apis", APIVersion, "DeleteAgentApp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAgentApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAgentAppResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAgentAppMcpServersRequest() (request *DeleteAgentAppMcpServersRequest) {
    request = &DeleteAgentAppMcpServersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apis", APIVersion, "DeleteAgentAppMcpServers")
    
    
    return
}

func NewDeleteAgentAppMcpServersResponse() (response *DeleteAgentAppMcpServersResponse) {
    response = &DeleteAgentAppMcpServersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAgentAppMcpServers
// 删除app的mcp server
func (c *Client) DeleteAgentAppMcpServers(request *DeleteAgentAppMcpServersRequest) (response *DeleteAgentAppMcpServersResponse, err error) {
    return c.DeleteAgentAppMcpServersWithContext(context.Background(), request)
}

// DeleteAgentAppMcpServers
// 删除app的mcp server
func (c *Client) DeleteAgentAppMcpServersWithContext(ctx context.Context, request *DeleteAgentAppMcpServersRequest) (response *DeleteAgentAppMcpServersResponse, err error) {
    if request == nil {
        request = NewDeleteAgentAppMcpServersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apis", APIVersion, "DeleteAgentAppMcpServers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAgentAppMcpServers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAgentAppMcpServersResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAgentAppModelServicesRequest() (request *DeleteAgentAppModelServicesRequest) {
    request = &DeleteAgentAppModelServicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apis", APIVersion, "DeleteAgentAppModelServices")
    
    
    return
}

func NewDeleteAgentAppModelServicesResponse() (response *DeleteAgentAppModelServicesResponse) {
    response = &DeleteAgentAppModelServicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAgentAppModelServices
// 删除app的model service关联
func (c *Client) DeleteAgentAppModelServices(request *DeleteAgentAppModelServicesRequest) (response *DeleteAgentAppModelServicesResponse, err error) {
    return c.DeleteAgentAppModelServicesWithContext(context.Background(), request)
}

// DeleteAgentAppModelServices
// 删除app的model service关联
func (c *Client) DeleteAgentAppModelServicesWithContext(ctx context.Context, request *DeleteAgentAppModelServicesRequest) (response *DeleteAgentAppModelServicesResponse, err error) {
    if request == nil {
        request = NewDeleteAgentAppModelServicesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apis", APIVersion, "DeleteAgentAppModelServices")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAgentAppModelServices require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAgentAppModelServicesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAgentCredentialRequest() (request *DeleteAgentCredentialRequest) {
    request = &DeleteAgentCredentialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apis", APIVersion, "DeleteAgentCredential")
    
    
    return
}

func NewDeleteAgentCredentialResponse() (response *DeleteAgentCredentialResponse) {
    response = &DeleteAgentCredentialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAgentCredential
// 删除Credential
func (c *Client) DeleteAgentCredential(request *DeleteAgentCredentialRequest) (response *DeleteAgentCredentialResponse, err error) {
    return c.DeleteAgentCredentialWithContext(context.Background(), request)
}

// DeleteAgentCredential
// 删除Credential
func (c *Client) DeleteAgentCredentialWithContext(ctx context.Context, request *DeleteAgentCredentialRequest) (response *DeleteAgentCredentialResponse, err error) {
    if request == nil {
        request = NewDeleteAgentCredentialRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apis", APIVersion, "DeleteAgentCredential")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAgentCredential require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAgentCredentialResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMcpServerRequest() (request *DeleteMcpServerRequest) {
    request = &DeleteMcpServerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apis", APIVersion, "DeleteMcpServer")
    
    
    return
}

func NewDeleteMcpServerResponse() (response *DeleteMcpServerResponse) {
    response = &DeleteMcpServerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteMcpServer
// 删除mcp server
func (c *Client) DeleteMcpServer(request *DeleteMcpServerRequest) (response *DeleteMcpServerResponse, err error) {
    return c.DeleteMcpServerWithContext(context.Background(), request)
}

// DeleteMcpServer
// 删除mcp server
func (c *Client) DeleteMcpServerWithContext(ctx context.Context, request *DeleteMcpServerRequest) (response *DeleteMcpServerResponse, err error) {
    if request == nil {
        request = NewDeleteMcpServerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apis", APIVersion, "DeleteMcpServer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMcpServer require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMcpServerResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteModelRequest() (request *DeleteModelRequest) {
    request = &DeleteModelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apis", APIVersion, "DeleteModel")
    
    
    return
}

func NewDeleteModelResponse() (response *DeleteModelResponse) {
    response = &DeleteModelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteModel
// 删除模型
func (c *Client) DeleteModel(request *DeleteModelRequest) (response *DeleteModelResponse, err error) {
    return c.DeleteModelWithContext(context.Background(), request)
}

// DeleteModel
// 删除模型
func (c *Client) DeleteModelWithContext(ctx context.Context, request *DeleteModelRequest) (response *DeleteModelResponse, err error) {
    if request == nil {
        request = NewDeleteModelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apis", APIVersion, "DeleteModel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteModel require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteModelResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteModelServiceRequest() (request *DeleteModelServiceRequest) {
    request = &DeleteModelServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apis", APIVersion, "DeleteModelService")
    
    
    return
}

func NewDeleteModelServiceResponse() (response *DeleteModelServiceResponse) {
    response = &DeleteModelServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteModelService
// 删除模型服务
func (c *Client) DeleteModelService(request *DeleteModelServiceRequest) (response *DeleteModelServiceResponse, err error) {
    return c.DeleteModelServiceWithContext(context.Background(), request)
}

// DeleteModelService
// 删除模型服务
func (c *Client) DeleteModelServiceWithContext(ctx context.Context, request *DeleteModelServiceRequest) (response *DeleteModelServiceResponse, err error) {
    if request == nil {
        request = NewDeleteModelServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apis", APIVersion, "DeleteModelService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteModelService require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteModelServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentAppRequest() (request *DescribeAgentAppRequest) {
    request = &DescribeAgentAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apis", APIVersion, "DescribeAgentApp")
    
    
    return
}

func NewDescribeAgentAppResponse() (response *DescribeAgentAppResponse) {
    response = &DescribeAgentAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAgentApp
// 查询app详情
func (c *Client) DescribeAgentApp(request *DescribeAgentAppRequest) (response *DescribeAgentAppResponse, err error) {
    return c.DescribeAgentAppWithContext(context.Background(), request)
}

// DescribeAgentApp
// 查询app详情
func (c *Client) DescribeAgentAppWithContext(ctx context.Context, request *DescribeAgentAppRequest) (response *DescribeAgentAppResponse, err error) {
    if request == nil {
        request = NewDescribeAgentAppRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apis", APIVersion, "DescribeAgentApp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgentApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentAppResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentAppMcpServersRequest() (request *DescribeAgentAppMcpServersRequest) {
    request = &DescribeAgentAppMcpServersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apis", APIVersion, "DescribeAgentAppMcpServers")
    
    
    return
}

func NewDescribeAgentAppMcpServersResponse() (response *DescribeAgentAppMcpServersResponse) {
    response = &DescribeAgentAppMcpServersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAgentAppMcpServers
// 查询app mcpServer关联列表
func (c *Client) DescribeAgentAppMcpServers(request *DescribeAgentAppMcpServersRequest) (response *DescribeAgentAppMcpServersResponse, err error) {
    return c.DescribeAgentAppMcpServersWithContext(context.Background(), request)
}

// DescribeAgentAppMcpServers
// 查询app mcpServer关联列表
func (c *Client) DescribeAgentAppMcpServersWithContext(ctx context.Context, request *DescribeAgentAppMcpServersRequest) (response *DescribeAgentAppMcpServersResponse, err error) {
    if request == nil {
        request = NewDescribeAgentAppMcpServersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apis", APIVersion, "DescribeAgentAppMcpServers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgentAppMcpServers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentAppMcpServersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentAppModelServicesRequest() (request *DescribeAgentAppModelServicesRequest) {
    request = &DescribeAgentAppModelServicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apis", APIVersion, "DescribeAgentAppModelServices")
    
    
    return
}

func NewDescribeAgentAppModelServicesResponse() (response *DescribeAgentAppModelServicesResponse) {
    response = &DescribeAgentAppModelServicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAgentAppModelServices
// 查询app modelService关联列表
func (c *Client) DescribeAgentAppModelServices(request *DescribeAgentAppModelServicesRequest) (response *DescribeAgentAppModelServicesResponse, err error) {
    return c.DescribeAgentAppModelServicesWithContext(context.Background(), request)
}

// DescribeAgentAppModelServices
// 查询app modelService关联列表
func (c *Client) DescribeAgentAppModelServicesWithContext(ctx context.Context, request *DescribeAgentAppModelServicesRequest) (response *DescribeAgentAppModelServicesResponse, err error) {
    if request == nil {
        request = NewDescribeAgentAppModelServicesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apis", APIVersion, "DescribeAgentAppModelServices")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgentAppModelServices require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentAppModelServicesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentAppsRequest() (request *DescribeAgentAppsRequest) {
    request = &DescribeAgentAppsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apis", APIVersion, "DescribeAgentApps")
    
    
    return
}

func NewDescribeAgentAppsResponse() (response *DescribeAgentAppsResponse) {
    response = &DescribeAgentAppsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAgentApps
// 查询app列表
func (c *Client) DescribeAgentApps(request *DescribeAgentAppsRequest) (response *DescribeAgentAppsResponse, err error) {
    return c.DescribeAgentAppsWithContext(context.Background(), request)
}

// DescribeAgentApps
// 查询app列表
func (c *Client) DescribeAgentAppsWithContext(ctx context.Context, request *DescribeAgentAppsRequest) (response *DescribeAgentAppsResponse, err error) {
    if request == nil {
        request = NewDescribeAgentAppsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apis", APIVersion, "DescribeAgentApps")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgentApps require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentAppsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentCredentialRequest() (request *DescribeAgentCredentialRequest) {
    request = &DescribeAgentCredentialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apis", APIVersion, "DescribeAgentCredential")
    
    
    return
}

func NewDescribeAgentCredentialResponse() (response *DescribeAgentCredentialResponse) {
    response = &DescribeAgentCredentialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAgentCredential
// 查询Credential详情
func (c *Client) DescribeAgentCredential(request *DescribeAgentCredentialRequest) (response *DescribeAgentCredentialResponse, err error) {
    return c.DescribeAgentCredentialWithContext(context.Background(), request)
}

// DescribeAgentCredential
// 查询Credential详情
func (c *Client) DescribeAgentCredentialWithContext(ctx context.Context, request *DescribeAgentCredentialRequest) (response *DescribeAgentCredentialResponse, err error) {
    if request == nil {
        request = NewDescribeAgentCredentialRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apis", APIVersion, "DescribeAgentCredential")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgentCredential require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentCredentialResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentCredentialsRequest() (request *DescribeAgentCredentialsRequest) {
    request = &DescribeAgentCredentialsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apis", APIVersion, "DescribeAgentCredentials")
    
    
    return
}

func NewDescribeAgentCredentialsResponse() (response *DescribeAgentCredentialsResponse) {
    response = &DescribeAgentCredentialsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAgentCredentials
// 查询Credential列表
func (c *Client) DescribeAgentCredentials(request *DescribeAgentCredentialsRequest) (response *DescribeAgentCredentialsResponse, err error) {
    return c.DescribeAgentCredentialsWithContext(context.Background(), request)
}

// DescribeAgentCredentials
// 查询Credential列表
func (c *Client) DescribeAgentCredentialsWithContext(ctx context.Context, request *DescribeAgentCredentialsRequest) (response *DescribeAgentCredentialsResponse, err error) {
    if request == nil {
        request = NewDescribeAgentCredentialsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apis", APIVersion, "DescribeAgentCredentials")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgentCredentials require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentCredentialsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMcpServerRequest() (request *DescribeMcpServerRequest) {
    request = &DescribeMcpServerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apis", APIVersion, "DescribeMcpServer")
    
    
    return
}

func NewDescribeMcpServerResponse() (response *DescribeMcpServerResponse) {
    response = &DescribeMcpServerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMcpServer
// 查询mcp server详情
func (c *Client) DescribeMcpServer(request *DescribeMcpServerRequest) (response *DescribeMcpServerResponse, err error) {
    return c.DescribeMcpServerWithContext(context.Background(), request)
}

// DescribeMcpServer
// 查询mcp server详情
func (c *Client) DescribeMcpServerWithContext(ctx context.Context, request *DescribeMcpServerRequest) (response *DescribeMcpServerResponse, err error) {
    if request == nil {
        request = NewDescribeMcpServerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apis", APIVersion, "DescribeMcpServer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMcpServer require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMcpServerResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMcpServersRequest() (request *DescribeMcpServersRequest) {
    request = &DescribeMcpServersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apis", APIVersion, "DescribeMcpServers")
    
    
    return
}

func NewDescribeMcpServersResponse() (response *DescribeMcpServersResponse) {
    response = &DescribeMcpServersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMcpServers
// 查询mcp server列表
func (c *Client) DescribeMcpServers(request *DescribeMcpServersRequest) (response *DescribeMcpServersResponse, err error) {
    return c.DescribeMcpServersWithContext(context.Background(), request)
}

// DescribeMcpServers
// 查询mcp server列表
func (c *Client) DescribeMcpServersWithContext(ctx context.Context, request *DescribeMcpServersRequest) (response *DescribeMcpServersResponse, err error) {
    if request == nil {
        request = NewDescribeMcpServersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apis", APIVersion, "DescribeMcpServers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMcpServers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMcpServersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModelRequest() (request *DescribeModelRequest) {
    request = &DescribeModelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apis", APIVersion, "DescribeModel")
    
    
    return
}

func NewDescribeModelResponse() (response *DescribeModelResponse) {
    response = &DescribeModelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeModel
// 查询模型详情
func (c *Client) DescribeModel(request *DescribeModelRequest) (response *DescribeModelResponse, err error) {
    return c.DescribeModelWithContext(context.Background(), request)
}

// DescribeModel
// 查询模型详情
func (c *Client) DescribeModelWithContext(ctx context.Context, request *DescribeModelRequest) (response *DescribeModelResponse, err error) {
    if request == nil {
        request = NewDescribeModelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apis", APIVersion, "DescribeModel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeModel require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeModelResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModelServiceRequest() (request *DescribeModelServiceRequest) {
    request = &DescribeModelServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apis", APIVersion, "DescribeModelService")
    
    
    return
}

func NewDescribeModelServiceResponse() (response *DescribeModelServiceResponse) {
    response = &DescribeModelServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeModelService
// 查询模型服务详情
func (c *Client) DescribeModelService(request *DescribeModelServiceRequest) (response *DescribeModelServiceResponse, err error) {
    return c.DescribeModelServiceWithContext(context.Background(), request)
}

// DescribeModelService
// 查询模型服务详情
func (c *Client) DescribeModelServiceWithContext(ctx context.Context, request *DescribeModelServiceRequest) (response *DescribeModelServiceResponse, err error) {
    if request == nil {
        request = NewDescribeModelServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apis", APIVersion, "DescribeModelService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeModelService require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeModelServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModelServicesRequest() (request *DescribeModelServicesRequest) {
    request = &DescribeModelServicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apis", APIVersion, "DescribeModelServices")
    
    
    return
}

func NewDescribeModelServicesResponse() (response *DescribeModelServicesResponse) {
    response = &DescribeModelServicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeModelServices
// 查询模型服务列表
func (c *Client) DescribeModelServices(request *DescribeModelServicesRequest) (response *DescribeModelServicesResponse, err error) {
    return c.DescribeModelServicesWithContext(context.Background(), request)
}

// DescribeModelServices
// 查询模型服务列表
func (c *Client) DescribeModelServicesWithContext(ctx context.Context, request *DescribeModelServicesRequest) (response *DescribeModelServicesResponse, err error) {
    if request == nil {
        request = NewDescribeModelServicesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apis", APIVersion, "DescribeModelServices")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeModelServices require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeModelServicesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModelsRequest() (request *DescribeModelsRequest) {
    request = &DescribeModelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apis", APIVersion, "DescribeModels")
    
    
    return
}

func NewDescribeModelsResponse() (response *DescribeModelsResponse) {
    response = &DescribeModelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeModels
// 查询模型列表
func (c *Client) DescribeModels(request *DescribeModelsRequest) (response *DescribeModelsResponse, err error) {
    return c.DescribeModelsWithContext(context.Background(), request)
}

// DescribeModels
// 查询模型列表
func (c *Client) DescribeModelsWithContext(ctx context.Context, request *DescribeModelsRequest) (response *DescribeModelsResponse, err error) {
    if request == nil {
        request = NewDescribeModelsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apis", APIVersion, "DescribeModels")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeModels require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeModelsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAgentAppRequest() (request *ModifyAgentAppRequest) {
    request = &ModifyAgentAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apis", APIVersion, "ModifyAgentApp")
    
    
    return
}

func NewModifyAgentAppResponse() (response *ModifyAgentAppResponse) {
    response = &ModifyAgentAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAgentApp
// 修改app
func (c *Client) ModifyAgentApp(request *ModifyAgentAppRequest) (response *ModifyAgentAppResponse, err error) {
    return c.ModifyAgentAppWithContext(context.Background(), request)
}

// ModifyAgentApp
// 修改app
func (c *Client) ModifyAgentAppWithContext(ctx context.Context, request *ModifyAgentAppRequest) (response *ModifyAgentAppResponse, err error) {
    if request == nil {
        request = NewModifyAgentAppRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apis", APIVersion, "ModifyAgentApp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAgentApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAgentAppResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAgentAppModelServicesRequest() (request *ModifyAgentAppModelServicesRequest) {
    request = &ModifyAgentAppModelServicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apis", APIVersion, "ModifyAgentAppModelServices")
    
    
    return
}

func NewModifyAgentAppModelServicesResponse() (response *ModifyAgentAppModelServicesResponse) {
    response = &ModifyAgentAppModelServicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAgentAppModelServices
// 编辑app的model service关联
func (c *Client) ModifyAgentAppModelServices(request *ModifyAgentAppModelServicesRequest) (response *ModifyAgentAppModelServicesResponse, err error) {
    return c.ModifyAgentAppModelServicesWithContext(context.Background(), request)
}

// ModifyAgentAppModelServices
// 编辑app的model service关联
func (c *Client) ModifyAgentAppModelServicesWithContext(ctx context.Context, request *ModifyAgentAppModelServicesRequest) (response *ModifyAgentAppModelServicesResponse, err error) {
    if request == nil {
        request = NewModifyAgentAppModelServicesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apis", APIVersion, "ModifyAgentAppModelServices")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAgentAppModelServices require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAgentAppModelServicesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAgentCredentialRequest() (request *ModifyAgentCredentialRequest) {
    request = &ModifyAgentCredentialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apis", APIVersion, "ModifyAgentCredential")
    
    
    return
}

func NewModifyAgentCredentialResponse() (response *ModifyAgentCredentialResponse) {
    response = &ModifyAgentCredentialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAgentCredential
// 修改Credential
func (c *Client) ModifyAgentCredential(request *ModifyAgentCredentialRequest) (response *ModifyAgentCredentialResponse, err error) {
    return c.ModifyAgentCredentialWithContext(context.Background(), request)
}

// ModifyAgentCredential
// 修改Credential
func (c *Client) ModifyAgentCredentialWithContext(ctx context.Context, request *ModifyAgentCredentialRequest) (response *ModifyAgentCredentialResponse, err error) {
    if request == nil {
        request = NewModifyAgentCredentialRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apis", APIVersion, "ModifyAgentCredential")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAgentCredential require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAgentCredentialResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMcpServerRequest() (request *ModifyMcpServerRequest) {
    request = &ModifyMcpServerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apis", APIVersion, "ModifyMcpServer")
    
    
    return
}

func NewModifyMcpServerResponse() (response *ModifyMcpServerResponse) {
    response = &ModifyMcpServerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMcpServer
// 修改mcp server
func (c *Client) ModifyMcpServer(request *ModifyMcpServerRequest) (response *ModifyMcpServerResponse, err error) {
    return c.ModifyMcpServerWithContext(context.Background(), request)
}

// ModifyMcpServer
// 修改mcp server
func (c *Client) ModifyMcpServerWithContext(ctx context.Context, request *ModifyMcpServerRequest) (response *ModifyMcpServerResponse, err error) {
    if request == nil {
        request = NewModifyMcpServerRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apis", APIVersion, "ModifyMcpServer")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMcpServer require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMcpServerResponse()
    err = c.Send(request, response)
    return
}

func NewModifyModelRequest() (request *ModifyModelRequest) {
    request = &ModifyModelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apis", APIVersion, "ModifyModel")
    
    
    return
}

func NewModifyModelResponse() (response *ModifyModelResponse) {
    response = &ModifyModelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyModel
// 修改模型
func (c *Client) ModifyModel(request *ModifyModelRequest) (response *ModifyModelResponse, err error) {
    return c.ModifyModelWithContext(context.Background(), request)
}

// ModifyModel
// 修改模型
func (c *Client) ModifyModelWithContext(ctx context.Context, request *ModifyModelRequest) (response *ModifyModelResponse, err error) {
    if request == nil {
        request = NewModifyModelRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apis", APIVersion, "ModifyModel")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyModel require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyModelResponse()
    err = c.Send(request, response)
    return
}

func NewModifyModelServiceRequest() (request *ModifyModelServiceRequest) {
    request = &ModifyModelServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("apis", APIVersion, "ModifyModelService")
    
    
    return
}

func NewModifyModelServiceResponse() (response *ModifyModelServiceResponse) {
    response = &ModifyModelServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyModelService
// 修改模型服务
func (c *Client) ModifyModelService(request *ModifyModelServiceRequest) (response *ModifyModelServiceResponse, err error) {
    return c.ModifyModelServiceWithContext(context.Background(), request)
}

// ModifyModelService
// 修改模型服务
func (c *Client) ModifyModelServiceWithContext(ctx context.Context, request *ModifyModelServiceRequest) (response *ModifyModelServiceResponse, err error) {
    if request == nil {
        request = NewModifyModelServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "apis", APIVersion, "ModifyModelService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyModelService require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyModelServiceResponse()
    err = c.Send(request, response)
    return
}
